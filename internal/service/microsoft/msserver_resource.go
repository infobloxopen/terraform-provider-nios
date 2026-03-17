package microsoft

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForMsserver = "ad_domain,ad_sites,ad_user,address,comment,connection_status,connection_status_detail,dhcp_server,disabled,dns_server,dns_view,extattrs,grid_member,last_seen,log_destination,log_level,login_name,managing_member,ms_max_connection,ms_rpc_timeout_in_seconds,network_view,read_only,root_ad_domain,server_name,synchronization_min_delay,synchronization_status,synchronization_status_detail,use_log_destination,use_ms_max_connection,use_ms_rpc_timeout_in_seconds,version"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MsserverResource{}
var _ resource.ResourceWithImportState = &MsserverResource{}
var _ resource.ResourceWithValidateConfig = &MsserverResource{}

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

	// Add internal ID exists in the Extensible Attributes if not already present
	data.ExtAttrs, diags = AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	apiRes, _, err := r.client.MicrosoftAPI.
		MsserverAPI.
		Create(ctx).
		Msserver(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForMsserver).
		ReturnAsObject(1).
		Execute()
	if err != nil {
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

	apiRes, httpRes, err := r.client.MicrosoftAPI.
		MsserverAPI.
		Read(ctx, utils.ResolveIdentifier(data.Uuid, data.Ref)).
		ReturnFieldsPlus(readableAttributesForMsserver).
		ReturnAsObject(1).
		ProxySearch(config.GetProxySearch()).
		Execute()

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

	planExtAttrs := data.ExtAttrs
	diags = req.State.GetAttribute(ctx, path.Root("ref"), &data.Ref)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	diags = req.State.GetAttribute(ctx, path.Root("uuid"), &data.Uuid)
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

	apiRes, _, err := r.client.MicrosoftAPI.
		MsserverAPI.
		Update(ctx, utils.ResolveIdentifier(data.Uuid, data.Ref)).
		Msserver(*data.Expand(ctx, &resp.Diagnostics)).
		ReturnFieldsPlus(readableAttributesForMsserver).
		ReturnAsObject(1).
		Execute()
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

	httpRes, err := r.client.MicrosoftAPI.
		MsserverAPI.
		Delete(ctx, utils.ResolveIdentifier(data.Uuid, data.Ref)).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Msserver, got error: %s", err))
		return
	}
}

func (r *MsserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("uuid"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}
