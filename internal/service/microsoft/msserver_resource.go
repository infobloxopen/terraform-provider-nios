package microsoft

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/microsoft"

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/retry"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForMsserver = "ad_domain,ad_sites,ad_user,address,comment,connection_status,connection_status_detail,dhcp_server,disabled,dns_server,dns_view,extattrs,grid_member,last_seen,log_destination,log_level,login_name,managing_member,ms_max_connection,ms_rpc_timeout_in_seconds,network_view,read_only,root_ad_domain,server_name,synchronization_min_delay,synchronization_status,synchronization_status_detail,use_log_destination,use_ms_max_connection,use_ms_rpc_timeout_in_seconds,version"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MsserverResource{}
var _ resource.ResourceWithImportState = &MsserverResource{}
var _ resource.ResourceWithValidateConfig = &MsserverResource{}
var _ resource.ResourceWithModifyPlan = &MsserverResource{}

func NewMsserverResource() resource.Resource {
	return &MsserverResource{}
}

// MsserverResource defines the resource implementation.
type MsserverResource struct {
	client *niosclient.APIClient
}

func (r *MsserverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "microsoft_msserver"
}

func (r *MsserverResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Microsoft Server.",
		Attributes:          MsserverResourceSchemaAttributes,
	}
}

func (r *MsserverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type msServerHashState struct {
	AdSitePwd     string `json:"ad_site_pwd_hash"`
	AdUserPwd     string `json:"ad_user_pwd_hash"`
	DhcpServerPwd string `json:"dhcp_server_pwd_hash"`
	DnsServerPwd  string `json:"dns_server_pwd_hash"`
	LoginPwd      string `json:"login_pwd_hash"`
}

func hasPasswords(state msServerHashState) bool {
	return state.AdSitePwd != "" || state.AdUserPwd != "" || state.DhcpServerPwd != "" || state.DnsServerPwd != "" || state.LoginPwd != ""
}

func marshalMsServerPasswordState(state msServerHashState, diags *diag.Diagnostics) string {
	data, err := json.Marshal(state)
	if err != nil {
		diags.AddError("error marshalling password hash", err.Error())
		return ""
	}
	return string(data)
}

func (r *MsserverResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}

	var statePwdVersion types.Int64
	var adSitePwd, adUserPwd, dhcpServerPwd, dnsServerPwd, loginPwd types.String

	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("ad_sites").AtName("login_password"), &adSitePwd)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("ad_user").AtName("login_password"), &adUserPwd)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("dhcp_server").AtName("login_password"), &dhcpServerPwd)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("dns_server").AtName("login_password"), &dnsServerPwd)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("login_password"), &loginPwd)...)
	if resp.Diagnostics.HasError() {
		return
	}

	curRev := int64(0)
	if !statePwdVersion.IsNull() && !statePwdVersion.IsUnknown() {
		curRev = statePwdVersion.ValueInt64()
	}

	var prevEnvelope struct {
		Algo string `json:"algo"`
		Hash string `json:"hash"`
	}
	if b, diags := req.Private.GetKey(ctx, "password_hash"); diags != nil {
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else if b != nil {
		if err := json.Unmarshal(b, &prevEnvelope); err != nil {
			prevEnvelope.Hash = ""
		}
	}

	prevHashes := msServerHashState{}
	if prevEnvelope.Hash != "" {
		if err := json.Unmarshal([]byte(prevEnvelope.Hash), &prevHashes); err != nil {
			prevHashes = msServerHashState{}
		}
	}

	plannedHashes := prevHashes

	isPasswordPresent(loginPwd, &plannedHashes)
	isPasswordPresent(adSitePwd, &plannedHashes)
	isPasswordPresent(adUserPwd, &plannedHashes)
	isPasswordPresent(dhcpServerPwd, &plannedHashes)
	isPasswordPresent(dnsServerPwd, &plannedHashes)

	prevHasPasswords := hasPasswords(prevHashes)
	plannedHasPasswords := hasPasswords(plannedHashes)
	passwordsChanged := plannedHashes != prevHashes

	bump := false
	newHashToStore := prevEnvelope.Hash

	switch {

	case !plannedHasPasswords && prevHasPasswords:
		bump = true
		newHashToStore = ""

	case plannedHasPasswords && (prevHashes == msServerHashState{} || passwordsChanged):
		bump = true
		newHashToStore = marshalMsServerPasswordState(plannedHashes, &resp.Diagnostics)
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
			resp.Diagnostics.AddError("error marshalling password hash", err.Error())
			return
		}
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "password_hash", b)...)
		return
	}

	resp.Diagnostics.Append(
		resp.Plan.SetAttribute(ctx, path.Root("password_version"), types.Int64Value(curRev))...,
	)
}

func isPasswordPresent(password types.String, plannedHashes *msServerHashState) {
	if password.IsNull() && !password.IsUnknown() {
		plannedHashes.LoginPwd = password.ValueString()
	} else {
		plannedHashes.LoginPwd = ""
	}
}

func (r *MsserverResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data MsserverModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !data.AdSites.IsNull() && !data.AdSites.IsUnknown() {
		var obj MsserverAdSitesModel
		resp.Diagnostics.Append(data.AdSites.As(ctx, &obj, basetypes.ObjectAsOptions{})...)
		if resp.Diagnostics.HasError() {
			return
		}
		validateSubConfig(resp, obj.LoginName, obj.UseLogin, obj.SynchronizationMinDelay, obj.UseSynchronizationMinDelay, "adsites")
	}

	if !data.AdUser.IsNull() && !data.AdUser.IsUnknown() {
		var obj MsserverAdUserModel
		resp.Diagnostics.Append(data.AdUser.As(ctx, &obj, basetypes.ObjectAsOptions{})...)
		if resp.Diagnostics.HasError() {
			return
		}
		validateSubConfig(resp, obj.LoginName, obj.UseLogin, obj.SynchronizationInterval, obj.UseSynchronizationMinDelay, "aduser")
	}

	if !data.DhcpServer.IsNull() && !data.DhcpServer.IsUnknown() {
		var obj MsserverDhcpServerModel
		resp.Diagnostics.Append(data.DhcpServer.As(ctx, &obj, basetypes.ObjectAsOptions{})...)
		if resp.Diagnostics.HasError() {
			return
		}
		validateSubConfig(resp, obj.LoginName, obj.UseLogin, obj.SynchronizationMinDelay, obj.UseSynchronizationMinDelay, "dhcpserver")
	}

	if !data.DnsServer.IsNull() && !data.DnsServer.IsUnknown() {
		var obj MsserverDnsServerModel
		resp.Diagnostics.Append(data.DnsServer.As(ctx, &obj, basetypes.ObjectAsOptions{})...)
		if resp.Diagnostics.HasError() {
			return
		}
		validateSubConfig(resp, obj.LoginName, obj.UseLogin, obj.SynchronizationMinDelay, obj.UseSynchronizationMinDelay, "dnsserver")
	}
}

func validateSubConfig(
	resp *resource.ValidateConfigResponse,
	login types.String,
	useLogin types.Bool,
	syncDelay types.Int64,
	useSyncDelay types.Bool,
	blockName string,
) {

	// login validation
	loginSet := !login.IsNull() && !login.IsUnknown()
	useLoginSet := !useLogin.IsNull() && !useLogin.IsUnknown()

	if loginSet {
		if !useLoginSet || !useLogin.ValueBool() {
			resp.Diagnostics.AddAttributeError(
				path.Root(blockName).AtName("uselogin"),
				"Invalid Login Configuration",
				fmt.Sprintf("`%s.uselogin` must be set to true when `%s.login_name` is provided.", blockName, blockName),
			)
		}
	}

	if useLoginSet && useLogin.ValueBool() && !loginSet {
		resp.Diagnostics.AddAttributeError(
			path.Root(blockName).AtName("login_name"),
			"Missing Login Name",
			fmt.Sprintf("`%s.login_name` must be provided when `%s.uselogin` is set to true.", blockName, blockName),
		)
	}

	// synchronization validation
	syncDelaySet := !syncDelay.IsNull() && !syncDelay.IsUnknown()
	useSyncDelaySet := !useSyncDelay.IsNull() && !useSyncDelay.IsUnknown()

	if syncDelaySet {
		if !useSyncDelaySet || !useSyncDelay.ValueBool() {
			resp.Diagnostics.AddAttributeError(
				path.Root(blockName).AtName("use_synchronization_min_delay"),
				"Invalid Synchronization Configuration",
				fmt.Sprintf("`%s.use_synchronization_min_delay` must be set to true when `%s.synchronization_min_delay` is provided.", blockName, blockName),
			)
		}
	}

	if useSyncDelaySet && useSyncDelay.ValueBool() && !syncDelaySet {
		resp.Diagnostics.AddAttributeError(
			path.Root(blockName).AtName("synchronization_min_delay"),
			"Missing Synchronization Delay",
			fmt.Sprintf("`%s.synchronization_min_delay` must be provided when `%s.use_synchronization_min_delay` is set to true.", blockName, blockName),
		)
	}
}

func (r *MsserverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data MsserverModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	var passwordVersion types.Int64
	var adSitePwd, adUserPwd, dhcpServerPwd, dnsServerPwd, loginPwd types.String

	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("ad_sites").AtName("login_password"), &adSitePwd)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("ad_user").AtName("login_password"), &adUserPwd)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("dhcp_server").AtName("login_password"), &dhcpServerPwd)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("dns_server").AtName("login_password"), &dnsServerPwd)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("login_password"), &loginPwd)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Add internal ID exists in the Extensible Attributes if not already present
	data.ExtAttrs, diags = AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	payload := data.Expand(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	plannedHashes := msServerHashState{}
	isPasswordPresent(adSitePwd, &plannedHashes)
	isPasswordPresent(adUserPwd, &plannedHashes)
	isPasswordPresent(dhcpServerPwd, &plannedHashes)
	isPasswordPresent(dnsServerPwd, &plannedHashes)
	isPasswordPresent(loginPwd, &plannedHashes)

	*payload.AdSites.LoginPassword = adSitePwd.ValueString()
	*payload.AdUser.LoginPassword = adUserPwd.ValueString()
	*payload.DhcpServer.LoginPassword = dhcpServerPwd.ValueString()
	*payload.DnsServer.LoginPassword = dnsServerPwd.ValueString()
	*payload.LoginPassword = loginPwd.ValueString()

	var apiRes *microsoft.CreateMsserverResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.MicrosoftAPI.
			MsserverAPI.
			Create(ctx).
			Msserver(*payload).
			ReturnFieldsPlus(readableAttributesForMsserver).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Msserver, got error: %s", err))
		return
	}

	res := apiRes.CreateMsserverResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.AddError("Client Error", "Error while creating Msserver due to inherited Extensible attributes")
		return
	}

	if hasPasswords(plannedHashes) {
		passwordVersion = types.Int64Value(1)

		passwordJson := marshalMsServerPasswordState(plannedHashes, &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}
		val := map[string]string{
			"algo": "sha256",
			"hash": passwordJson,
		}
		b, err := json.Marshal(val)
		if err != nil {
			resp.Diagnostics.AddError("error marshalling password hash", err.Error())
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

func (r *MsserverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data MsserverModel

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
		apiRes  *microsoft.GetMsserverResponse
	)

	err := retry.Do(ctx, nil, func(ctx context.Context) (int, error) {
		var callErr error
		apiRes, httpRes, callErr = r.client.MicrosoftAPI.
			MsserverAPI.
			Read(ctx, resourceRef).
			ReturnFieldsPlus(readableAttributesForMsserver).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Msserver, got error: %s", err))
		return
	}

	res := apiRes.GetMsserverResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	if associateInternalId == nil {
		stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
		if stateExtAttrs == nil {
			resp.Diagnostics.AddError(
				"Missing Internal ID",
				"Unable to read Msserver because the internal ID (from extattrs_all) is missing or invalid.",
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
		resp.Diagnostics.AddError("Client Error", "Error while reading Msserver due to inherited Extensible attributes")
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MsserverResource) ReadByExtAttrs(ctx context.Context, data *MsserverModel, resp *resource.ReadResponse) bool {
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

	apiRes, _, err := r.client.MicrosoftAPI.
		MsserverAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForMsserver).
		ProxySearch(config.GetProxySearch()).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Msserver by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListMsserverResponseObject.GetResult()

	// If the list is empty, the resource no longer exists so remove it from state
	if len(results) == 0 {
		resp.State.RemoveResource(ctx)
		return true
	}

	res := results[0]

	// Remove inherited external attributes from extattrs
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		return true
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)

	return true
}

func (r *MsserverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data MsserverModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	var adSitePwd, adUserPwd, dhcpServerPwd, dnsServerPwd, loginPwd types.String

	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("ad_sites").AtName("login_password"), &adSitePwd)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("ad_user").AtName("login_password"), &adUserPwd)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("dhcp_server").AtName("login_password"), &dhcpServerPwd)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("dns_server").AtName("login_password"), &dnsServerPwd)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("login_password"), &loginPwd)...)
	if resp.Diagnostics.HasError() {
		return
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

	payload := data.Expand(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	plannedHashes := msServerHashState{}
	isPasswordPresent(adSitePwd, &plannedHashes)
	isPasswordPresent(adUserPwd, &plannedHashes)
	isPasswordPresent(dhcpServerPwd, &plannedHashes)
	isPasswordPresent(dnsServerPwd, &plannedHashes)
	isPasswordPresent(loginPwd, &plannedHashes)

	*payload.AdSites.LoginPassword = adSitePwd.ValueString()
	*payload.AdUser.LoginPassword = adUserPwd.ValueString()
	*payload.DhcpServer.LoginPassword = dhcpServerPwd.ValueString()
	*payload.DnsServer.LoginPassword = dnsServerPwd.ValueString()
	*payload.LoginPassword = loginPwd.ValueString()

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	var apiRes *microsoft.UpdateMsserverResponse

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		var (
			httpRes *http.Response
			callErr error
		)
		apiRes, httpRes, callErr = r.client.MicrosoftAPI.
			MsserverAPI.
			Update(ctx, resourceRef).
			Msserver(*payload).
			ReturnFieldsPlus(readableAttributesForMsserver).
			ReturnAsObject(1).
			Execute()

		if httpRes != nil {
			return httpRes.StatusCode, callErr
		}
		return 0, callErr
	})

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Msserver, got error: %s", err))
		return
	}

	res := apiRes.UpdateMsserverResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.AddError("Client Error", "Error while updating Msserver due to inherited Extensible attributes")
		return
	}

	if hasPasswords(plannedHashes) {
		passwordJson := marshalMsServerPasswordState(plannedHashes, &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}
		val := map[string]string{
			"algo": "sha256",
			"hash": passwordJson,
		}
		b, err := json.Marshal(val)
		if err != nil {
			resp.Diagnostics.AddError("error marshalling password hash", err.Error())
			return
		}
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "password_hash", b)...)
	} else {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "password_hash", nil)...)
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if associateInternalId != nil {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", nil)...)
	}
}

func (r *MsserverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data MsserverModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resourceRef := utils.ExtractResourceRef(data.Ref.ValueString())

	err := retry.Do(ctx, retry.TransientErrors, func(ctx context.Context) (int, error) {
		httpRes, callErr := r.client.MicrosoftAPI.
			MsserverAPI.
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Msserver, got error: %s", err))
		return
	}
}

func (r *MsserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ref"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}
