package grid

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/retry"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForMember = "active_position,additional_ip_list,automated_traffic_capture_setting,bgp_as,comment,config_addr_type,csp_access_key,csp_member_setting,dns_resolver_setting,dscp,email_setting,enable_ha,enable_lom,enable_member_redirect,enable_ro_api_access,extattrs,external_syslog_backup_servers,external_syslog_server_enable,ha_cloud_platform,ha_on_cloud,host_name,ipv6_setting,ipv6_static_routes,is_dscp_capable,lan2_enabled,lan2_port_setting,lom_network_config,lom_users,master_candidate,member_service_communication,mgmt_port_setting,mmdb_ea_build_time,mmdb_geoip_build_time,nat_setting,node_info,ntp_setting,ospf_list,passive_ha_arp_enabled,platform,pre_provisioning,preserve_if_owns_delegation,remote_console_access_enable,router_id,service_status,service_type_configuration,snmp_setting,static_routes,support_access_enable,support_access_info,syslog_proxy_setting,syslog_servers,syslog_size,threshold_traps,time_zone,traffic_capture_auth_dns_setting,traffic_capture_chr_setting,traffic_capture_qps_setting,traffic_capture_rec_dns_setting,traffic_capture_rec_queries_setting,trap_notifications,upgrade_group,use_automated_traffic_capture,use_dns_resolver_setting,use_dscp,use_email_setting,use_enable_lom,use_enable_member_redirect,use_external_syslog_backup_servers,use_remote_console_access_enable,use_snmp_setting,use_support_access_enable,use_syslog_proxy_setting,use_threshold_traps,use_time_zone,use_traffic_capture_auth_dns,use_traffic_capture_chr,use_traffic_capture_qps,use_traffic_capture_rec_dns,use_traffic_capture_rec_queries,use_trap_notifications,use_v4_vrrp,vip_setting,vpn_mtu"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MemberResource{}
var _ resource.ResourceWithImportState = &MemberResource{}
var _ resource.ResourceWithValidateConfig = &MemberResource{}

var _ resource.ResourceWithModifyPlan = &MemberResource{}

func NewMemberResource() resource.Resource {
	return &MemberResource{}
}

// MemberResource defines the resource implementation.
type MemberResource struct {
	client *niosclient.APIClient
}

func (r *MemberResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "grid_member"
}

func (r *MemberResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Member resource object.",
		Attributes:          MemberResourceSchemaAttributes,
	}
}

func (r *MemberResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*niosclient.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *niosclient.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

type memberPasswordsHashState struct {
	BkpServersPwd string `json:"bkp_servers_password_hash"`
	LomUsersPwd   string `json:"lom_users_password_hash"`
}

func hasPasswordHashes(state memberPasswordsHashState) bool {
	return state.BkpServersPwd != "" || state.LomUsersPwd != ""
}

func hashPasswords[T any](ctx context.Context, passwordList types.List, diags *diag.Diagnostics, passwordOf func(T) types.String) string {
	if passwordList.IsNull() || passwordList.IsUnknown() {
		return ""
	}

	var pwdModels []T
	diags.Append(passwordList.ElementsAs(ctx, &pwdModels, false)...)
	if diags.HasError() {
		return ""
	}

	passwordHashes := make([]string, 0, len(pwdModels))
	hasAnyPassword := false

	for _, pwdModel := range pwdModels {
		password := passwordOf(pwdModel)
		switch {
		case password.IsUnknown(), password.IsNull():
			passwordHashes = append(passwordHashes, "")
		default:
			hasAnyPassword = true
			sum := sha256.Sum256([]byte(password.ValueString()))
			passwordHashes = append(passwordHashes, hex.EncodeToString(sum[:]))
		}
	}

	if !hasAnyPassword {
		return ""
	}
	// Uses config order
	data, err := json.Marshal(passwordHashes)
	if err != nil {
		diags.AddError("CLI Secrets Hash Error", err.Error())
		return ""
	}

	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func marshalSecretsHashState(state memberPasswordsHashState, diags *diag.Diagnostics) string {
	data, err := json.Marshal(state)
	if err != nil {
		diags.AddError("error marshalling password hash state", err.Error())
		return ""
	}
	return string(data)
}

func (r *MemberResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}

	var (
		statePwdVersion types.Int64
		bkpServers      types.List
		lomUsers        types.List
	)

	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("external_syslog_backup_servers"), &bkpServers)...)
	if !req.State.Raw.IsNull() && req.State.Raw.IsKnown() {
		resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("password_version"), &statePwdVersion)...)
	}
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("lom_users"), &lomUsers)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// var passwordVersion types.Int64
	curRev := int64(0)

	if !statePwdVersion.IsNull() && !statePwdVersion.IsUnknown() {
		curRev = statePwdVersion.ValueInt64()
	}

	var prevEnvelope struct {
		Algo string `json:"algo"`
		Hash string `json:"hash"`
	}
	if b, diags := req.Private.GetKey(ctx, "secrets_hash"); diags != nil {
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else if b != nil {
		if err := json.Unmarshal(b, &prevEnvelope); err != nil {
			prevEnvelope.Hash = ""
		}
	}

	var data MemberModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var configData MemberModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &configData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	bkpServersHash := hashPasswords(ctx, configData.ExternalSyslogBackupServers, &resp.Diagnostics,
		func(m MemberExternalSyslogBackupServersModel) types.String { return m.Password })
	if resp.Diagnostics.HasError() {
		return
	}

	lomUsersHash := hashPasswords(ctx, configData.LomUsers, &resp.Diagnostics,
		func(m MemberLomUsersModel) types.String { return m.Password })
	if resp.Diagnostics.HasError() {
		return
	}

	plannedHashes := memberPasswordsHashState{BkpServersPwd: bkpServersHash, LomUsersPwd: lomUsersHash}

	var prev struct {
		Algo string `json:"algo"`
		Hash string `json:"hash"`
	}
	prevHashes := memberPasswordsHashState{}

	if b, diags := req.Private.GetKey(ctx, "password_hash"); diags != nil {
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else if b != nil {
		if err := json.Unmarshal(b, &prev); err != nil {
			prev.Hash = ""
		}
		if prev.Hash != "" {
			if err := json.Unmarshal([]byte(prev.Hash), &prevHashes); err != nil {
				prevHashes = memberPasswordsHashState{}
			}
		}
	}

	prevHasPasswords := hasPasswordHashes(prevHashes)
	plannedHasPasswords := hasPasswordHashes(plannedHashes)
	passwordsChanged := plannedHashes != prevHashes
	bump := false
	newHashToStore := prev.Hash

	switch {
	case !plannedHasPasswords && prevHasPasswords:
		bump = true
		newHashToStore = ""
	case plannedHasPasswords && (!prevHasPasswords || passwordsChanged):
		bump = true
		newHashToStore = marshalSecretsHashState(plannedHashes, &resp.Diagnostics)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	if bump {
		newRev := types.Int64Value(curRev + 1)
		resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("password_version"), newRev)...)

		val := map[string]string{
			"algo": "sha256",
			"hash": newHashToStore,
		}
		b, err := json.Marshal(val)
		if err != nil {
			resp.Diagnostics.AddError("error marshalling secrets hash", err.Error())
			return
		}
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "password_hash", b)...)
		return
	}

	resp.Diagnostics.Append(
		resp.Plan.SetAttribute(ctx, path.Root("password_version"), types.Int64Value(curRev))...,
	)
}

func (r *MemberResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data MemberModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Add internal ID exists in the Extensible Attributes if not already present
	data.ExtAttrs, diags = AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	if !data.SyslogServers.IsNull() && !data.SyslogServers.IsUnknown() {
		processedList, ok := r.processSyslogServers(ctx, data.SyslogServers, &resp.Diagnostics)
		if !ok {
			return
		}
		data.SyslogServers = processedList
	}

	if !data.GridLevelDnsResolverSetting.IsNull() && !data.GridLevelDnsResolverSetting.IsUnknown() {
		dnsResolverSetting := ExpandMemberDnsResolverSetting(ctx, data.GridLevelDnsResolverSetting, &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}

		err := utils.ConfigureGridDNSResolver(ctx, r.client.GridAPI, dnsResolverSetting)
		if err != nil {
			resp.Diagnostics.AddError("Grid DNS Resolver Configuration Error", fmt.Sprintf("Unable to configure grid-level DNS resolver: %s", err))
			return
		}
	}

	payload := data.Expand(ctx, &resp.Diagnostics, true)
	if resp.Diagnostics.HasError() {
		return
	}

	var passwordVersion types.Int64
	var bkpServers []MemberExternalSyslogBackupServersModel
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("external_syslog_backup_servers"), &bkpServers)...)
	if resp.Diagnostics.HasError() {
		return
	}

	for i, server := range bkpServers {
		if i >= len(payload.ExternalSyslogBackupServers) {
			break
		}
		if !server.Password.IsNull() && !server.Password.IsUnknown() {
			password := server.Password.ValueString()
			payload.ExternalSyslogBackupServers[i].Password = &password
		}
	}

	var lomUsers []MemberLomUsersModel
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("lom_users"), &lomUsers)...)
	if resp.Diagnostics.HasError() {
		return
	}

	for i, user := range lomUsers {
		if i >= len(payload.LomUsers) {
			break
		}
		if !user.Password.IsNull() && !user.Password.IsUnknown() {
			password := user.Password.ValueString()
			payload.LomUsers[i].Password = &password
		}
	}

	var apiRes *grid.CreateMemberResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.GridAPI.
			MemberAPI.
			Create(ctx).
			Member(*payload).
			ReturnFieldsPlus(readableAttributesForMember).
			ReturnAsObject(1).
			Execute()

		if httpRes != nil {
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

	if err != nil {
		if retry.IsAlreadyExistsErr(err) {
			// Resource already exists, import required
			resp.Diagnostics.AddError(
				"Resource Already Exists",
				fmt.Sprintf("Resource already exists, error: %s.\nPlease import the existing resource into terraform state.", err.Error()),
			)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Member, got error: %s", err))
		return
	}

	res := apiRes.CreateMemberResponseAsObject.GetResult()

	if !data.PreProvisioning.IsUnknown() && !data.PreProvisioning.IsNull() || (!data.TrafficCaptureAuthDnsSetting.IsUnknown() && !data.TrafficCaptureAuthDnsSetting.IsNull()) || (!data.MemberServiceCommunication.IsUnknown() && !data.MemberServiceCommunication.IsNull()) {
		apiRes2, _, err2 := r.client.GridAPI.
			MemberAPI.
			Update(ctx, utils.ExtractResourceRef(*res.Ref)).
			Member(*data.Expand(ctx, &resp.Diagnostics, false)).
			ReturnFieldsPlus(readableAttributesForMember).
			ReturnAsObject(1).
			Execute()
		if err2 != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to Create Member with pre-provisioning or syslog proxy settings, got error: %s", err2))
			return
		}
		res = apiRes2.UpdateMemberResponseAsObject.GetResult()
	}

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.AddError("Client Error", "Error while creating Member due to inherited Extensible attributes")
		return
	}
	var configData MemberModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &configData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	bkpServersHash := hashPasswords(ctx, configData.ExternalSyslogBackupServers, &resp.Diagnostics,
		func(m MemberExternalSyslogBackupServersModel) types.String { return m.Password })

	lomUsersHash := hashPasswords(ctx, configData.LomUsers, &resp.Diagnostics,
		func(m MemberLomUsersModel) types.String { return m.Password })
	plannedHashes := memberPasswordsHashState{BkpServersPwd: bkpServersHash, LomUsersPwd: lomUsersHash}

	if hasPasswordHashes(plannedHashes) {
		passwordVersion = types.Int64Value(1)

		newHashToStore := marshalSecretsHashState(plannedHashes, &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}

		val := map[string]string{
			"algo": "sha256",
			"hash": newHashToStore,
		}
		b, err := json.Marshal(val)
		if err != nil {
			resp.Diagnostics.AddError("error marshalling secrets hash", err.Error())
			return
		}
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "password_hash", b)...)
	} else {
		passwordVersion = types.Int64Value(0)
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "password_hash", nil)...)
	}
	data.PasswordVersion = passwordVersion

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MemberResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data MemberModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	associateInternalId, diags := req.Private.GetKey(ctx, "associate_internal_id")
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	var (
		httpRes *http.Response
		apiRes  *grid.GetMemberResponse
	)

	err := retry.Do(ctx, nil, func(ctx context.Context) (int, error) {
		var callErr error
		apiRes, httpRes, callErr = r.client.GridAPI.
			MemberAPI.
			Read(ctx, resourceRef).
			ReturnFieldsPlus(readableAttributesForMember).
			ReturnAsObject(1).
			ProxySearch(config.GetProxySearch()).
			Execute()

		if httpRes != nil {
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Member, got error: %s", err))
		return
	}

	res := apiRes.GetMemberResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	if associateInternalId == nil {
		stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
		if stateExtAttrs == nil {
			resp.Diagnostics.AddError(
				"Missing Internal ID",
				"Unable to read Member because the internal ID (from extattrs_all) is missing or invalid.",
			)
			return
		}

		stateTerraformId := (*stateExtAttrs)[terraformInternalIDEA]
		if apiTerraformId.Value != stateTerraformId.Value {
			if r.ReadByExtAttrs(ctx, &data, resp) {
				return
			}
		}
	}

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.AddError("Client Error", "Error while reading Member due to inherited Extensible attributes")
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MemberResource) ReadByExtAttrs(ctx context.Context, data *MemberModel, resp *resource.ReadResponse) bool {
	var diags diag.Diagnostics

	if data.ExtAttrsAll.IsNull() {
		return false
	}

	internalIdExtAttr := *ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
	if diags.HasError() {
		return false
	}

	internalId := internalIdExtAttr[terraformInternalIDEA].Value
	if internalId == "" {
		return false
	}

	idMap := map[string]interface{}{
		terraformInternalIDEA: internalId,
	}

	apiRes, _, err := r.client.GridAPI.
		MemberAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForMember).
		ProxySearch(config.GetProxySearch()).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Member by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListMemberResponseObject.GetResult()

	// If the list is empty, the resource no longer exists so remove it from state
	if len(results) == 0 {
		resp.State.RemoveResource(ctx)
		return true
	}

	res := results[0]

	// Remove inherited external attributes from extattrs
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return true
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)

	return true
}

func (r *MemberResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data MemberModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	if !data.SyslogServers.IsNull() && !data.SyslogServers.IsUnknown() {
		processedList, ok := r.processSyslogServers(ctx, data.SyslogServers, &resp.Diagnostics)
		if !ok {
			return
		}
		data.SyslogServers = processedList
	}

	planExtAttrs := data.ExtAttrs
	diags = req.State.GetAttribute(ctx, path.Root("ref"), &data.Ref)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	diags = req.State.GetAttribute(ctx, path.Root("extattrs_all"), &data.ExtAttrsAll)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	associateInternalId, diags := req.Private.GetKey(ctx, "associate_internal_id")
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	if associateInternalId != nil {
		data.ExtAttrs, diags = AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}
	}

	// Add Inherited Extensible Attributes
	data.ExtAttrs, diags = AddInheritedExtAttrs(ctx, data.ExtAttrs, data.ExtAttrsAll)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	if !data.GridLevelDnsResolverSetting.IsNull() && !data.GridLevelDnsResolverSetting.IsUnknown() {
		var stateGridLevelDnsResolverSetting types.Object
		// Check if grid-level DNS resolver setting has changed.
		diags = req.State.GetAttribute(ctx, path.Root("grid_level_dns_resolver_setting"), &stateGridLevelDnsResolverSetting)
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}
		if !data.GridLevelDnsResolverSetting.Equal(stateGridLevelDnsResolverSetting) {
			dnsResolverSetting := ExpandMemberDnsResolverSetting(ctx, data.GridLevelDnsResolverSetting, &resp.Diagnostics)
			if resp.Diagnostics.HasError() {
				return
			}

			err := utils.ConfigureGridDNSResolver(ctx, r.client.GridAPI, dnsResolverSetting)
			if err != nil {
				resp.Diagnostics.AddError("Grid DNS Resolver Configuration Error", fmt.Sprintf("Unable to configure grid-level DNS resolver: %s", err))
				return
			}
		}
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	payload := data.Expand(ctx, &resp.Diagnostics, false)
	if resp.Diagnostics.HasError() {
		return
	}

	var passwordVersion types.Int64
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("password_version"), &passwordVersion)...)
	var bkpServers []MemberExternalSyslogBackupServersModel
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("external_syslog_backup_servers"), &bkpServers)...)
	if resp.Diagnostics.HasError() {
		return
	}

	for i, server := range bkpServers {
		if i >= len(payload.ExternalSyslogBackupServers) {
			break
		}
		if !server.Password.IsNull() && !server.Password.IsUnknown() {
			password := server.Password.ValueString()
			payload.ExternalSyslogBackupServers[i].Password = &password
		}
	}

	var lomUsers []MemberLomUsersModel
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("lom_users"), &lomUsers)...)
	if resp.Diagnostics.HasError() {
		return
	}

	for i, user := range lomUsers {
		if i >= len(payload.LomUsers) {
			break
		}
		if !user.Password.IsNull() && !user.Password.IsUnknown() {
			password := user.Password.ValueString()
			payload.LomUsers[i].Password = &password
		}
	}

	var apiRes *grid.UpdateMemberResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.GridAPI.
			MemberAPI.
			Update(ctx, resourceRef).
			Member(*payload).
			ReturnFieldsPlus(readableAttributesForMember).
			ReturnAsObject(1).
			Execute()

		if httpRes != nil {
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Member, got error: %s", err))
		return
	}

	res := apiRes.UpdateMemberResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.AddError("Client Error", "Error while updating Member due to inherited Extensible attributes")
		return
	}

	data.PasswordVersion = passwordVersion

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if associateInternalId != nil {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", nil)...)
	}
}

func (r *MemberResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data MemberModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		httpRes, callErr := r.client.GridAPI.
			MemberAPI.
			Delete(ctx, resourceRef).
			Execute()

		if httpRes != nil {
			if httpRes.StatusCode == http.StatusNotFound {
				return 0, nil
			}
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Member, got error: %s", err))
		return
	}
}

func (r *MemberResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data MemberModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !data.EmailSetting.IsNull() && !data.EmailSetting.IsUnknown() {
		emailSetting := data.EmailSetting.Attributes()
		if emailSetting["use_authentication"].String() == "true" {
			if emailSetting["password"].IsNull() || emailSetting["password"].IsUnknown() {
				resp.Diagnostics.AddError("Validation Error", "Password must be provided when use_authentication is true")
			}
			if emailSetting["from_address"].IsNull() || emailSetting["from_address"].IsUnknown() {
				resp.Diagnostics.AddError("Validation Error", "From address must be provided when use_authentication is true")
			}
			if emailSetting["address"].IsNull() || emailSetting["address"].IsUnknown() {
				resp.Diagnostics.AddError("Validation Error", "Address must be provided when use_authentication is true")
			}
			if emailSetting["smtps"].String() == "true" {
				if emailSetting["port_number"].IsNull() || emailSetting["port_number"].IsUnknown() {
					resp.Diagnostics.AddError("Validation Error", "Port must be provided when email_settings.smtps is true")
				} else {
					port := emailSetting["port_number"].String()
					if port != "587" && port != "2525" {
						resp.Diagnostics.AddError("Validation Error", "Port must be either 587 or 2525 when email_settings.smtps is true")
					}
				}
			}
		}
	}

	// CSP Member Setting Validation
	if !data.CspMemberSetting.IsNull() && !data.CspMemberSetting.IsUnknown() {
		if data.ConfigureCspMemberSetting.IsNull() || data.ConfigureCspMemberSetting.IsUnknown() || !data.ConfigureCspMemberSetting.ValueBool() {
			resp.Diagnostics.AddError("Validation Error", "configure_csp_member_setting must be set to true when csp_member_setting is provided")
	// HA Cloud Platform Validations - bidirectional checks
	if !data.HaCloudPlatform.IsNull() && !data.HaCloudPlatform.IsUnknown() {
		// When ha_cloud_platform is provided, ha_on_cloud must be true
		if data.HaOnCloud.IsNull() || data.HaOnCloud.IsUnknown() || !data.HaOnCloud.ValueBool() {
			resp.Diagnostics.AddError("Validation Error", "ha_on_cloud must be set to true when ha_cloud_platform is provided")
		}
	}

	if !data.HaOnCloud.IsNull() && !data.HaOnCloud.IsUnknown() && data.HaOnCloud.ValueBool() {
		// When ha_on_cloud is true, ha_cloud_platform must be provided
		if data.HaCloudPlatform.IsNull() || data.HaCloudPlatform.IsUnknown() {
			resp.Diagnostics.AddError("Validation Error", "ha_cloud_platform must be set when ha_on_cloud is true")
		}
		// When ha_on_cloud is true, enable_ha must be true
		if data.EnableHa.IsNull() || data.EnableHa.IsUnknown() || !data.EnableHa.ValueBool() {
			resp.Diagnostics.AddError("Validation Error", "enable_ha must be true when ha_on_cloud is true")
		}
		// When ha_on_cloud is true, platform must be VNIOS
		if data.Platform.IsNull() || data.Platform.IsUnknown() || data.Platform.ValueString() != "VNIOS" {
			resp.Diagnostics.AddError("Validation Error", "platform must be set to VNIOS when ha_on_cloud is true")
		}
	}

	// Validation: enable_ha requires node_info to be provided
	if !data.EnableHa.IsNull() && !data.EnableHa.IsUnknown() && data.EnableHa.ValueBool() {
		if data.NodeInfo.IsNull() || data.NodeInfo.IsUnknown() {
			resp.Diagnostics.AddError("Validation Error", "node_info must be provided when enable_ha is true")
		}
	}

	mgmtCheckComplete := false
	if !data.NodeInfo.IsNull() && !data.NodeInfo.IsUnknown() {
		var nodeInfo []MemberNodeInfoModel
		diags := data.NodeInfo.ElementsAs(ctx, &nodeInfo, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		// Validation: If enable_ha is true, node_info must have exactly 2 elements
		if !data.EnableHa.IsNull() && !data.EnableHa.IsUnknown() && data.EnableHa.ValueBool() {
			if len(nodeInfo) != 2 {
				resp.Diagnostics.AddError("Validation Error", "node_info must contain exactly 2 elements when enable_ha is true")
			}
		}

		for i, node := range nodeInfo {
			// Validation: lan_ha_port_setting requires enable_ha to be true
			if !node.LanHaPortSetting.IsNull() && !node.LanHaPortSetting.IsUnknown() {
				lanHaPortSetting := node.LanHaPortSetting.Attributes()

				if !lanHaPortSetting["mgmt_lan"].IsNull() && !lanHaPortSetting["mgmt_lan"].IsUnknown() {
					if data.EnableHa.IsNull() || data.EnableHa.IsUnknown() || !data.EnableHa.ValueBool() {
						resp.Diagnostics.AddError("Validation Error",
							fmt.Sprintf("enable_ha must be set to true when node_info[%d].lan_ha_port_setting.mgmt_lan is provided", i))
					}
				}

				// Validation: When enable_ha is true, both mgmt_lan and ha_ip_address must be set for all nodes
				if !data.EnableHa.IsNull() && !data.EnableHa.IsUnknown() && data.EnableHa.ValueBool() {
					if lanHaPortSetting["mgmt_lan"].IsNull() || lanHaPortSetting["mgmt_lan"].IsUnknown() {
						resp.Diagnostics.AddError("Validation Error",
							fmt.Sprintf("node_info[%d].lan_ha_port_setting.mgmt_lan must be set when enable_ha is true", i))
					}
					if lanHaPortSetting["ha_ip_address"].IsNull() || lanHaPortSetting["ha_ip_address"].IsUnknown() {
						resp.Diagnostics.AddError("Validation Error",
							fmt.Sprintf("node_info[%d].lan_ha_port_setting.ha_ip_address must be set when enable_ha is true", i))
					}
				}
			} else if !data.EnableHa.IsNull() && !data.EnableHa.IsUnknown() && data.EnableHa.ValueBool() {
				// Validation: When enable_ha is true, lan_ha_port_setting must be provided for all nodes
				resp.Diagnostics.AddError("Validation Error",
					fmt.Sprintf("node_info[%d].lan_ha_port_setting must be provided when enable_ha is true", i))
			}

			if !node.V6MgmtNetworkSetting.IsNull() && !node.V6MgmtNetworkSetting.IsUnknown() {
				v6MgmtNetworkSetting := node.V6MgmtNetworkSetting.Attributes()

				// Validation: When v6_mgmt_network_setting is set, mgmt_port_setting.enabled must be true
				if data.MgmtPortSetting.IsNull() || data.MgmtPortSetting.IsUnknown() || data.MgmtPortSetting.Attributes()["enabled"].String() != "true" {
					resp.Diagnostics.AddError("Validation Error", "mgmt_port_setting.enabled must be set to true when node_info.v6_mgmt_network_setting is provided")
				}

				// Validation: When v6_mgmt_network_setting is set, virtual_ip, gateway, cidr_prefix, and enabled must be set
				if v6MgmtNetworkSetting["virtual_ip"].IsNull() || v6MgmtNetworkSetting["virtual_ip"].IsUnknown() {
					resp.Diagnostics.AddError("Validation Error", fmt.Sprintf("node_info[%d].v6_mgmt_network_setting.virtual_ip must be set when node_info.v6_mgmt_network_setting is provided", i))
				}
				if v6MgmtNetworkSetting["gateway"].IsNull() || v6MgmtNetworkSetting["gateway"].IsUnknown() {
					resp.Diagnostics.AddError("Validation Error", fmt.Sprintf("node_info[%d].v6_mgmt_network_setting.gateway must be set when node_info.v6_mgmt_network_setting is provided", i))
				}
				if v6MgmtNetworkSetting["cidr_prefix"].IsNull() || v6MgmtNetworkSetting["cidr_prefix"].IsUnknown() {
					resp.Diagnostics.AddError("Validation Error", fmt.Sprintf("node_info[%d].v6_mgmt_network_setting.cidr_prefix must be set when node_info.v6_mgmt_network_setting is provided", i))
				}
				if v6MgmtNetworkSetting["enabled"].IsNull() || v6MgmtNetworkSetting["enabled"].IsUnknown() || v6MgmtNetworkSetting["enabled"].String() != "true" {
					resp.Diagnostics.AddError("Validation Error", fmt.Sprintf("node_info[%d].v6_mgmt_network_setting.enabled must be set to true when node_info.v6_mgmt_network_setting is provided", i))
				}

				if v6MgmtNetworkSetting["auto_router_config_enabled"].String() == "true" {
					if !v6MgmtNetworkSetting["gateway"].IsNull() && !v6MgmtNetworkSetting["gateway"].IsUnknown() {
						resp.Diagnostics.AddError("Validation Error", "node_info.v6_mgmt_network_setting.gateway cannot be set when node_info.v6_mgmt_network_setting.auto_router_config_enabled is true")
					}
				}
			}
			if !node.MgmtPhysicalSetting.IsNull() && !node.MgmtPhysicalSetting.IsUnknown() {
				mgmtPhysicalSetting := node.MgmtPhysicalSetting.Attributes()
				if mgmtPhysicalSetting["auto_port_setting_enabled"].String() == "true" {
					if (!mgmtPhysicalSetting["speed"].IsNull() && !mgmtPhysicalSetting["speed"].IsUnknown()) || (!mgmtPhysicalSetting["duplex"].IsNull() && !mgmtPhysicalSetting["duplex"].IsUnknown()) {
						resp.Diagnostics.AddError("Validation Error", "node_info.mgmt_physical_setting.speed and node_info.mgmt_physical_setting.duplex cannot be set when node_info.mgmt_physical_setting.auto_port_setting_enabled is true")
					}
				} else {
					if mgmtPhysicalSetting["speed"].IsNull() || mgmtPhysicalSetting["speed"].IsUnknown() {
						resp.Diagnostics.AddError("Validation Error", "node_info.mgmt_physical_setting.speed must be set when node_info.mgmt_physical_setting.auto_port_setting_enabled is false")
					}
					if mgmtPhysicalSetting["duplex"].IsNull() || mgmtPhysicalSetting["duplex"].IsUnknown() {
						resp.Diagnostics.AddError("Validation Error", "node_info.mgmt_physical_setting.duplex must be set when node_info.mgmt_physical_setting.auto_port_setting_enabled is false")
					}
				}
			}
		}

		if len(nodeInfo) > 0 && (!nodeInfo[0].MgmtNetworkSetting.IsNull() && !nodeInfo[0].MgmtNetworkSetting.IsUnknown()) {
			if data.MgmtPortSetting.IsNull() || data.MgmtPortSetting.IsUnknown() || data.MgmtPortSetting.Attributes()["enabled"].String() != "true" {
				resp.Diagnostics.AddError("Validation Error", "node_info.mgmt_network_setting can be set only when mgmt_port_setting.enabled is set to true")
			} else {
				mgmtCheckComplete = true
			}
		}
		// enableHaFalse: true when enable_ha is null (defaults to false) or explicitly false.
		// Skipped when enable_ha is unknown (value not yet determined at plan time).
		enableHaFalse := data.EnableHa.IsNull() || (!data.EnableHa.IsUnknown() && !data.EnableHa.ValueBool())
		// enableHaTrue: true only when enable_ha is explicitly set to true.
		enableHaTrue := !data.EnableHa.IsNull() && !data.EnableHa.IsUnknown() && data.EnableHa.ValueBool()

		nodeCount := len(nodeInfo)

		// Condition 1: len(nodeInfo) == 2 requires enable_ha to be true
		if nodeCount == 2 && enableHaFalse {
			resp.Diagnostics.AddError("Validation Error", "enable_ha must be true when node_info has 2 nodes")
		}

		// Condition 2: enable_ha true requires exactly 2 nodes (not more)
		if enableHaTrue && nodeCount > 2 {
			resp.Diagnostics.AddError("Validation Error", "node_info must have exactly 2 nodes when enable_ha is true")
		}

		// Condition 3a: len(nodeInfo) > 2 with enable_ha false (or not set) is not allowed
		if nodeCount > 2 && enableHaFalse {
			resp.Diagnostics.AddError("Validation Error", "node_info cannot have more than 2 nodes when enable_ha is false")
		}

		// Condition 3b: len(nodeInfo) == 1 with enable_ha true is not allowed
		if nodeCount == 1 && enableHaTrue {
			resp.Diagnostics.AddError("Validation Error", "node_info must have exactly 2 nodes when enable_ha is true; a single node_info entry is not valid")
		}
	}

	if !data.MgmtPortSetting.IsNull() && !data.MgmtPortSetting.IsUnknown() {
		mgmtPortSetting := data.MgmtPortSetting.Attributes()

		if mgmtPortSetting["enabled"].String() == "true" {

			if !mgmtCheckComplete {

				var nodeInfo []MemberNodeInfoModel
				diags := data.NodeInfo.ElementsAs(ctx, &nodeInfo, false)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				if data.NodeInfo.IsNull() || len(nodeInfo) == 0 ||
					(len(nodeInfo) > 0 && (nodeInfo[0].MgmtNetworkSetting.IsNull() || nodeInfo[0].MgmtNetworkSetting.IsUnknown()) &&
						(nodeInfo[0].V6MgmtNetworkSetting.IsNull() || nodeInfo[0].V6MgmtNetworkSetting.IsUnknown())) {
					resp.Diagnostics.AddError("Validation Error", "Either node_info.mgmt_network_setting or node_info.v6_mgmt_network_setting must be set when mgmt_port_setting.enabled is true")
				}
			}
		}

		if mgmtPortSetting["vpn_enabled"].String() == "true" {
			if mgmtPortSetting["enabled"].IsNull() || mgmtPortSetting["enabled"].IsUnknown() || mgmtPortSetting["enabled"].String() != "true" {
				resp.Diagnostics.AddError("Validation Error", "enabled must be true when vpn_enabled is true in mgmt_port_setting")
			}
		}
	}

	if !data.SyslogProxySetting.IsNull() && !data.SyslogProxySetting.IsUnknown() && data.SyslogProxySetting.Attributes()["enable"].String() == "true" {
		if data.ExternalSyslogServerEnable.IsNull() || data.ExternalSyslogServerEnable.IsUnknown() || !data.ExternalSyslogServerEnable.ValueBool() {
			resp.Diagnostics.AddError("Validation Error", "external_syslog_server_enable must be true when syslog_proxy_setting.enabled is true")
		}
	}

	if !data.Lan2PortSetting.IsNull() && !data.Lan2PortSetting.IsUnknown() {
		if data.Lan2PortSetting.Attributes()["enabled"].String() == "true" {
			if !data.Lan2Enabled.IsNull() && !data.Lan2Enabled.IsUnknown() && !data.Lan2Enabled.ValueBool() {
				resp.Diagnostics.AddError("Validation Error", "lan2_enabled must be true when lan2_port_setting.enabled is true")
			}
		} else if !data.Lan2PortSetting.Attributes()["network_setting"].IsNull() && !data.Lan2PortSetting.Attributes()["network_setting"].IsUnknown() {
			resp.Diagnostics.AddError("Validation Error", "lan2_port_setting.network_setting cannot be set when lan2_port_setting.enabled is false or not set")
		} else if !data.Lan2PortSetting.Attributes()["v6_network_setting"].IsNull() && !data.Lan2PortSetting.Attributes()["v6_network_setting"].IsUnknown() {
			resp.Diagnostics.AddError("Validation Error", "lan2_port_setting.v6_network_setting cannot be set when lan2_port_setting.enabled is false or not set")
		} else if !data.Lan2Enabled.IsNull() && !data.Lan2Enabled.IsUnknown() && data.Lan2Enabled.ValueBool() {
			resp.Diagnostics.AddError("Validation Error", "lan2_enabled can be set to true only when lan2_port_setting.enabled is true")
		}
	} else {
		if !data.Lan2Enabled.IsNull() && !data.Lan2Enabled.IsUnknown() && data.Lan2Enabled.ValueBool() {
			resp.Diagnostics.AddError("Validation Error", "lan2_enabled can be set to true only when lan2_port_setting is set and lan2_port_setting.enabled is true")
		}
	}

	if !data.ConfigAddrType.IsNull() && !data.ConfigAddrType.IsUnknown() && (data.ConfigAddrType.ValueString() == "IPV6" || data.ConfigAddrType.ValueString() == "BOTH") {
		if data.Ipv6Setting.IsNull() || data.Ipv6Setting.IsUnknown() {
			resp.Diagnostics.AddError("Validation Error", "ipv6_setting must be provided when config_addr_type is set to IPV6 or BOTH")
		} else {
			if data.Ipv6Setting.Attributes()["enabled"].IsNull() || data.Ipv6Setting.Attributes()["enabled"].IsUnknown() || data.Ipv6Setting.Attributes()["enabled"].String() != "true" {
				resp.Diagnostics.AddError("Validation Error", "ipv6_setting.enabled must be true when config_addr_type is set to IPV6 or BOTH")
			}
		}
	}

	if !data.Ipv6Setting.IsNull() && !data.Ipv6Setting.IsUnknown() {
		ipv6Attrs := data.Ipv6Setting.Attributes()
		hasVirtualIP := !ipv6Attrs["virtual_ip"].IsNull() && !ipv6Attrs["virtual_ip"].IsUnknown()
		hasCidrPrefix := !ipv6Attrs["cidr_prefix"].IsNull() && !ipv6Attrs["cidr_prefix"].IsUnknown()
		hasGateway := !ipv6Attrs["gateway"].IsNull() && !ipv6Attrs["gateway"].IsUnknown()
		if (hasVirtualIP && hasCidrPrefix && hasGateway) &&
			(data.ConfigAddrType.IsNull() || data.ConfigAddrType.IsUnknown() ||
				(data.ConfigAddrType.ValueString() != "IPV6" && data.ConfigAddrType.ValueString() != "BOTH")) {
			resp.Diagnostics.AddError("Validation Error", "config_addr_type must be set to IPV6 or BOTH when ipv6_setting.virtual_ip, ipv6_setting.cidr_prefix, and ipv6_setting.gateway are provided")
		}
	}

	if !data.ConfigAddrType.IsNull() && !data.ConfigAddrType.IsUnknown() && data.ConfigAddrType.ValueString() == "IPV4" {
		if !data.Ipv6Setting.IsNull() && !data.Ipv6Setting.IsUnknown() {
			if !data.Ipv6Setting.Attributes()["virtual_ip"].IsNull() && !data.Ipv6Setting.Attributes()["virtual_ip"].IsUnknown() {
				resp.Diagnostics.AddError("Validation Error", "ipv6_setting.virtual_ip cannot be set when config_addr_type is set to IPV4")
			}
			if !data.Ipv6Setting.Attributes()["gateway"].IsNull() && !data.Ipv6Setting.Attributes()["gateway"].IsUnknown() {
				resp.Diagnostics.AddError("Validation Error", "ipv6_setting.gateway cannot be set when config_addr_type is set to IPV4")
			}
			if !data.Ipv6Setting.Attributes()["cidr_prefix"].IsNull() && !data.Ipv6Setting.Attributes()["cidr_prefix"].IsUnknown() {
				resp.Diagnostics.AddError("Validation Error", "ipv6_setting.cidr_prefix cannot be set when config_addr_type is set to IPV4")
			}
			if !data.Ipv6Setting.Attributes()["enabled"].IsNull() && !data.Ipv6Setting.Attributes()["enabled"].IsUnknown() && data.Ipv6Setting.Attributes()["enabled"].String() != "false" {
				resp.Diagnostics.AddError("Validation Error", "ipv6_setting.enabled must be false when config_addr_type is set to IPV4")
			}
		}
	}

	if !data.ConfigAddrType.IsNull() && !data.ConfigAddrType.IsUnknown() && data.ConfigAddrType.ValueString() == "IPV6" {
		if !data.VipSetting.IsNull() && !data.VipSetting.IsUnknown() {
			if !data.VipSetting.Attributes()["subnet_mask"].IsNull() && !data.VipSetting.Attributes()["subnet_mask"].IsUnknown() {
				resp.Diagnostics.AddError("Validation Error", "vip_setting.subnet_mask cannot be set when config_addr_type is set to IPV6")
			}
			if !data.VipSetting.Attributes()["gateway"].IsNull() && !data.VipSetting.Attributes()["gateway"].IsUnknown() {
				resp.Diagnostics.AddError("Validation Error", "vip_setting.gateway cannot be set when config_addr_type is set to IPV6")
			}
			if !data.VipSetting.Attributes()["address"].IsNull() && !data.VipSetting.Attributes()["address"].IsUnknown() {
				resp.Diagnostics.AddError("Validation Error", "vip_setting.address cannot be set when config_addr_type is set to IPV6")
			}
		}
		if !data.ServiceTypeConfiguration.IsUnknown() && (data.ServiceTypeConfiguration.IsNull() || data.ServiceTypeConfiguration.ValueString() == "ALL_V4") {
			resp.Diagnostics.AddError("Validation Error", "service_type_configuration must be ALL_V6 when the config_addr_type is IPV6")
		}
	}

	if !data.Ipv6Setting.IsNull() && !data.Ipv6Setting.IsUnknown() {
		ipv6Setting := data.Ipv6Setting.Attributes()
		if ipv6Setting["auto_router_config_enabled"].String() == "true" {
			if !ipv6Setting["gateway"].IsNull() && !ipv6Setting["gateway"].IsUnknown() {
				resp.Diagnostics.AddError("Validation Error", "gateway cannot be set when ipv6_setting.auto_router_config_enabled is true")
			}
		}
	}

	if !data.EnableHa.IsNull() && !data.EnableHa.IsUnknown() {
		if data.EnableHa.ValueBool() {
			if data.RouterId.IsNull() || data.RouterId.IsUnknown() {
				resp.Diagnostics.AddError("Validation Error", "router_id must be provided when enable_ha is true")
			}
		} else {
			if !data.RouterId.IsNull() && !data.RouterId.IsUnknown() {
				resp.Diagnostics.AddError("Validation Error", "router_id cannot not be set when enable_ha is false")
			}
		}
	}

	if !data.ThresholdTraps.IsNull() && !data.ThresholdTraps.IsUnknown() {
		var thresholdTraps []MemberThresholdTrapsModel
		diags := data.ThresholdTraps.ElementsAs(ctx, &thresholdTraps, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		allowedTrapTypes := []string{"CpuUsage", "DBObjects", "Disk", "ExtStorage", "FDUsage", "FastpathDroppedTraffic",
			"Fastpathmbuffdepletion", "IPAMUtilization", "Memory", "NetworkCapacity", "RPZHitRate",
			"RecursiveClients", "Reporting", "ReportingVolume", "Rootfs", "SwapUsage", "TcpUdpFloodAlertRate",
			"TcpUdpFloodDropRate", "ThreatProtectionDroppedTraffic", "ThreatProtectionTotalTraffic", "Tmpfs"}

		additionalTrapsforGM := []string{"DBWrites"}
		allowedTrapTypesMap := make(map[string]bool)
		additionalTrapsforGMMap := make(map[string]bool)
		for _, trapType := range allowedTrapTypes {
			allowedTrapTypesMap[trapType] = true
		}
		for _, trapType := range additionalTrapsforGM {
			additionalTrapsforGMMap[trapType] = true
		}

		for i, trap := range thresholdTraps {
			if trap.TrapType.IsNull() || trap.TrapType.IsUnknown() {
				resp.Diagnostics.AddError("Validation Error", fmt.Sprintf("trap_type must be provided for threshold_traps[%d]", i))
			} else if !allowedTrapTypesMap[trap.TrapType.ValueString()] {
				if !data.UpgradeGroup.IsNull() && !data.UpgradeGroup.IsUnknown() && data.UpgradeGroup.ValueString() == "Grid Master" && additionalTrapsforGMMap[trap.TrapType.ValueString()] {
					continue
				}
				resp.Diagnostics.AddError("Validation Error", fmt.Sprintf("trap_type value '%s' is not valid for threshold_traps[%d]. Valid Values are: %v", trap.TrapType.ValueString(), i, allowedTrapTypes))
			}
		}
	}

	if !data.TrapNotifications.IsNull() && !data.TrapNotifications.IsUnknown() {
		var trapNotifications []MemberTrapNotificationsModel
		diags := data.TrapNotifications.ElementsAs(ctx, &trapNotifications, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		allowedTrapTypes := []string{"AnalyticsRPZ", "WATCHFRR", "DFP", "AutomatedTrafficCapture", "BFD",
			"BGP", "Backup", "CPU", "CaptivePortal", "CiscoISEServer", "Clear", "CloudAPI", "CloudDNSsync",
			"Cluster", "Controld", "DHCP", "DNS", "DNSAttack", "DNSIntegrityCheck", "DNSIntegrityCheckConnection",
			"Database", "DisconnectedGrid", "Discovery", "DiscoveryConflict", "DiscoveryUnmanaged", "Disk",
			"DuplicateIP", "ENAT", "FDUsage", "FTP", "Fan", "HA", "HAOnCloud", "HSM", "HTTP",
			"IFMAP", "IMC", "IMCGRPCServer", "IPAMUtilization", "IPMIDevice", "LCD", "LDAPServers",
			"License", "Login", "MGM", "MSServer", "Memory", "NTP", "Network", "OCSPResponders", "OSPF",
			"OSPF6", "Outbound", "PowerSupply", "RAID", "RIRSWIP", "RPZHitRate", "RecursiveClients",
			"Reporting", "RootFS", "SNMP", "SSH", "SerialConsole", "SwapUsage", "Syslog", "System",
			"TFTP", "Taxii", "ThreatInsight", "ThreatProtection", "TmpFS"}

		additionalTrapsforGM := []string{"DBActivity", "WATCHFRR"}
		allowedTrapTypesMap := make(map[string]bool)
		additionalTrapsforGMMap := make(map[string]bool)
		for _, trapType := range allowedTrapTypes {
			allowedTrapTypesMap[trapType] = true
		}
		for _, trapType := range additionalTrapsforGM {
			additionalTrapsforGMMap[trapType] = true
		}

		for i, trap := range trapNotifications {
			if trap.TrapType.IsNull() || trap.TrapType.IsUnknown() {
				resp.Diagnostics.AddError("Validation Error", fmt.Sprintf("trap_type must be provided for trap_notifications[%d]", i))
			} else if !allowedTrapTypesMap[trap.TrapType.ValueString()] {
				if !data.UpgradeGroup.IsNull() && !data.UpgradeGroup.IsUnknown() && data.UpgradeGroup.ValueString() == "Grid Master" && additionalTrapsforGMMap[trap.TrapType.ValueString()] {
					continue
				}
				resp.Diagnostics.AddError("Validation Error", fmt.Sprintf("trap_type value '%s' is not valid for trap_notifications[%d]. Valid Values are: %v", trap.TrapType.ValueString(), i, allowedTrapTypes))
			}
		}
	}
}

func (r *MemberResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ref"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}
