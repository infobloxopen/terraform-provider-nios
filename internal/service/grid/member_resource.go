package grid

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForMember = "active_position,additional_ip_list,automated_traffic_capture_setting,bgp_as,comment,config_addr_type,csp_access_key,csp_member_setting,dns_resolver_setting,dscp,email_setting,enable_ha,enable_lom,enable_member_redirect,enable_ro_api_access,extattrs,external_syslog_backup_servers,external_syslog_server_enable,ha_cloud_platform,ha_on_cloud,host_name,ipv6_setting,ipv6_static_routes,is_dscp_capable,lan2_enabled,lan2_port_setting,lom_network_config,lom_users,master_candidate,member_service_communication,mgmt_port_setting,mmdb_ea_build_time,mmdb_geoip_build_time,nat_setting,node_info,ntp_setting,ospf_list,passive_ha_arp_enabled,platform,pre_provisioning,preserve_if_owns_delegation,remote_console_access_enable,router_id,service_status,service_type_configuration,snmp_setting,static_routes,support_access_enable,support_access_info,syslog_proxy_setting,syslog_servers,syslog_size,threshold_traps,time_zone,traffic_capture_auth_dns_setting,traffic_capture_chr_setting,traffic_capture_qps_setting,traffic_capture_rec_dns_setting,traffic_capture_rec_queries_setting,trap_notifications,upgrade_group,use_automated_traffic_capture,use_dns_resolver_setting,use_dscp,use_email_setting,use_enable_lom,use_enable_member_redirect,use_external_syslog_backup_servers,use_remote_console_access_enable,use_snmp_setting,use_support_access_enable,use_syslog_proxy_setting,use_threshold_traps,use_time_zone,use_traffic_capture_auth_dns,use_traffic_capture_chr,use_traffic_capture_qps,use_traffic_capture_rec_dns,use_traffic_capture_rec_queries,use_trap_notifications,use_v4_vrrp,vip_setting,vpn_mtu"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MemberResource{}
var _ resource.ResourceWithImportState = &MemberResource{}
var _ resource.ResourceWithValidateConfig = &MemberResource{}

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

	apiRes, _, err := r.client.GridAPI.
		MemberAPI.
		Create(ctx).
		Member(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForMember).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Member, got error: %s", err))
		return
	}

	res := apiRes.CreateMemberResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		resp.Diagnostics.AddError("Client Error", "Error while creating Member due to inherited Extensible attributes")
		return
	}

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

	apiRes, httpRes, err := r.client.GridAPI.
		MemberAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForMember).
		ReturnAsObject(1).
		ProxySearch(config.GetProxySearch()).
		Execute()

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

	apiRes, _, err := r.client.GridAPI.
		MemberAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Member(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForMember).
		ReturnAsObject(1).
		Execute()
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

	httpRes, err := r.client.GridAPI.
		MemberAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
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

	if !data.HaOnCloud.IsNull() && !data.HaOnCloud.IsUnknown() {
		if data.EnableHa.IsNull() || data.EnableHa.IsUnknown() || data.EnableHa.ValueBool() == false {
			resp.Diagnostics.AddError("Validation Error", "enable_ha must be true when ha_on_cloud is provided")
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

		if len(nodeInfo) > 0 && (!nodeInfo[0].MgmtNetworkSetting.IsNull() && !nodeInfo[0].MgmtNetworkSetting.IsUnknown()) {
			if data.MgmtPortSetting.IsNull() || data.MgmtPortSetting.IsUnknown() || data.MgmtPortSetting.Attributes()["enabled"].String() != "true" {
				resp.Diagnostics.AddError("Validation Error", "node_info.mgmt_network_setting is set but mgmt_port_setting.enabled is not true")
			} else {
				mgmtCheckComplete = true
			}
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
					resp.Diagnostics.AddError("Validation Error", "Either node_info.mgmt_network_setting or node.v6_mgmt_network_setting must be set when mgmt_port_setting.enabled is true")
				}
			}
		}

		if mgmtPortSetting["vpn_enabled"].String() == "true" {
			if mgmtPortSetting["enabled"].IsNull() || mgmtPortSetting["enabled"].IsUnknown() || mgmtPortSetting["enabled"].String() != "true" {
				resp.Diagnostics.AddError("Validation Error", "enabled must be true when vpn_enabled is true in mgmt_port_setting")
			}
		}
	}

}

func (r *MemberResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ref"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}
