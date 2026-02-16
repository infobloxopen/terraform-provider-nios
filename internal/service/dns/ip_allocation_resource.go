package dns

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"maps"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForIPAllocation = "aliases,allow_telnet,cli_credentials,cloud_info,comment,configure_for_dns,creation_time,ddns_protected,device_description,device_location,device_type,device_vendor,disable,disable_discovery,dns_aliases,dns_name,extattrs,ipv4addrs,ipv6addrs,last_queried,ms_ad_user_data,name,network_view,rrset_order,snmp3_credential,snmp_credential,ttl,use_cli_credentials,use_dns_ea_inheritance,use_snmp3_credential,use_snmp_credential,use_ttl,view,zone"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &IPAllocationResource{}
var _ resource.ResourceWithImportState = &IPAllocationResource{}
var _ resource.ResourceWithValidateConfig = &IPAllocationResource{}
var _ resource.ResourceWithModifyPlan = &IPAllocationResource{}

func NewIPAllocationResource() resource.Resource {
	return &IPAllocationResource{}
}

// IPAllocationResource defines the resource implementation.
type IPAllocationResource struct {
	client *niosclient.APIClient
}

func (r *IPAllocationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "ip_allocation"
}

func (r *IPAllocationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "",
		Attributes:          IPAllocationResourceSchemaAttributes,
	}
}

func (r *IPAllocationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *IPAllocationResource) ModifyPlan(
	ctx context.Context,
	req resource.ModifyPlanRequest,
	resp *resource.ModifyPlanResponse,
) {
	// If the whole resource is being destroyed, do nothing with revision.
	if req.Plan.Raw.IsNull() {
		return
	}

	// 1) Read write-only values from CONFIG (never from plan/state)
	var authPwd, privPwd types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(
		ctx, path.Root("snmp3_credential").AtName("authentication_password"), &authPwd,
	)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(
		ctx, path.Root("snmp3_credential").AtName("privacy_password"), &privPwd,
	)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 2) Read plan/state objects and current revision from STATE
	var planSnmp3, stateSnmp3 types.Object
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("snmp3_credential"), &planSnmp3)...)
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("snmp3_credential"), &stateSnmp3)...)

	var stateRev types.Int64
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("snmp3_secret_revision"), &stateRev)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Normalize stateRev if null (e.g., first apply)
	curRev := int64(0)
	if !stateRev.IsNull() && !stateRev.IsUnknown() {
		curRev = stateRev.ValueInt64()
	}

	planHas := !planSnmp3.IsNull()
	stateHas := !stateSnmp3.IsNull()

	// 3) Read previous secrets hash from PRIVATE STATE
	var prev struct {
		Algo string `json:"algo"`
		Hash string `json:"hash"`
	}
	if b, diags := req.Private.GetKey(ctx, "snmp3_secrets_hash"); diags != nil {
		resp.Diagnostics.Append(diags...)
	} else if b != nil {
		if err := json.Unmarshal(b, &prev); err != nil {
			// Older buggy format: ignore and treat as different
			prev.Hash = ""
		}
	}

	// 4) Compute planned secrets hash WHEN we have actual values
	// We only compute a new hash when at least one secret is explicitly set in config
	// and is not unknown. If neither secret is supplied, we treat the planned hash
	// as "same as previous" (prev.Hash) to avoid false bumps.
	computeNewHash := (!authPwd.IsNull() && !authPwd.IsUnknown()) ||
		(!privPwd.IsNull() && !privPwd.IsUnknown())

	plannedHash := prev.Hash // default: assume "unchanged"
	if computeNewHash {
		ap := ""
		pp := ""
		if !authPwd.IsNull() && !authPwd.IsUnknown() {
			ap = authPwd.ValueString()
		}
		if !privPwd.IsNull() && !privPwd.IsUnknown() {
			pp = privPwd.ValueString()
		}
		h := sha256.New()
		h.Write([]byte(ap))
		h.Write([]byte{0})
		h.Write([]byte(pp))
		plannedHash = hex.EncodeToString(h.Sum(nil))
	}

	// 5) Decide whether to bump revision and what to store in PRIVATE STATE
	// Cases:
	//  A) Added (stateHas=false, planHas=true)      -> bump
	//  B) Removed (stateHas=true, planHas=false)    -> bump
	//  C) Both present and secrets changed          -> bump
	//  D) Both absent                               -> retain
	//  E) Both present and secrets unchanged        -> retain
	bump := false
	newHashToStore := prev.Hash // default retain

	switch {
	case planHas && !stateHas:
		// Added
		bump = true
		newHashToStore = plannedHash // if user didn't supply secrets, equals prev.Hash (empty on first)

	case !planHas && stateHas:
		// Removed
		bump = true
		newHashToStore = "" // clear private marker on removal

	case planHas && stateHas:
		// Compare only if we know the planned hash (i.e., user supplied at least one secret)
		if computeNewHash && plannedHash != prev.Hash {
			bump = true
			newHashToStore = plannedHash
		}
		// If computeNewHash=false, we retain (no bump) since secrets are unknown/unspecified

	default:
		// Both absent -> retain
	}

	// 6) Apply the decision to the PLAN and PRIVATE STATE
	if bump {
		newRev := types.Int64Value(curRev + 1)
		resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("snmp3_secret_revision"), newRev)...)

		val := map[string]string{"algo": "sha256", "hash": newHashToStore}
		b, err := json.Marshal(val)
		if err != nil {
			resp.Diagnostics.AddError("Private State Marshal Error", err.Error())
			return
		}
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "snmp3_secrets_hash", b)...)
	} else {
		// Retain current revision
		resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("snmp3_secret_revision"), types.Int64Value(curRev))...)
	}
}

func (r *IPAllocationResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data IPAllocationModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if both ipv4addrs and ipv6addrs are null or empty
	ipv4Empty := data.Ipv4addrs.IsNull() || len(data.Ipv4addrs.Elements()) == 0
	ipv6Empty := data.Ipv6addrs.IsNull() || len(data.Ipv6addrs.Elements()) == 0

	if ipv4Empty && ipv6Empty {
		resp.Diagnostics.AddError(
			"Invalid Configuration",
			"At least one of 'ipv4addrs' or 'ipv6addrs' must be configured.",
		)
	}

	if len(data.Ipv4addrs.Elements()) > 1 {
		resp.Diagnostics.AddError(
			"Invalid Configuration",
			"'ipv4addrs' can contain at most one element.",
		)
	}

	if len(data.Ipv6addrs.Elements()) > 1 {
		resp.Diagnostics.AddError(
			"Invalid Configuration",
			"'ipv6addrs' can contain at most one element.",
		)
	}

	// Validate FQDN if configure_for_dns is true
	if data.ConfigureForDns.ValueBool() {
		name := data.Name.ValueString()
		if strings.TrimSpace(name) != name {
			resp.Diagnostics.AddAttributeError(
				path.Root("name"),
				"Invalid FQDN",
				"must not contain leading or trailing whitespaces",
			)
		}
		if strings.HasSuffix(name, ".") {
			resp.Diagnostics.AddAttributeError(
				path.Root("name"),
				"Invalid FQDN",
				"must not end with a dot",
			)
		}
		if !strings.Contains(name, ".") {
			resp.Diagnostics.AddAttributeError(
				path.Root("name"),
				"Invalid FQDN",
				"Parent not found - ensure the zone or network exists before creating the record.",
			)
		}
	}
}

func (r *IPAllocationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var diags diag.Diagnostics
	var data IPAllocationModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Add internal ID exists in the Extensible Attributes if not already present
	data.ExtAttrs, diags = AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
	if diags.HasError() {
		return
	}

	// Populate the internal ID field from the extattrs map
	extAttrsMap := data.ExtAttrs.Elements()
	if internalIDValue, exists := extAttrsMap[terraformInternalIDEA]; exists {
		if stringVal, ok := internalIDValue.(types.String); ok {
			data.InternalID = stringVal
		} else {
			resp.Diagnostics.AddError("Type Error", "Internal ID in ExtAttrs is not a string")
			return
		}
	} else {
		resp.Diagnostics.AddError("Missing Internal ID", "Internal ID was not found in ExtAttrs after generation")
		return
	}

	// Save original IPv4 function call attributes
	savedIPv4FuncCalls := r.saveNestedFuncCallAttrs(data.Ipv4addrs)

	// Save original IPv6 function call attributes
	savedIPv6FuncCalls := r.saveNestedFuncCallAttrs(data.Ipv6addrs)

	var plan IPAllocationModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Build API payload: include write-only values from plan/config if set
	// (Plan includes them as null; read from config directly if needed)
	var cfg IPAllocationModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &cfg)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read write-only from CONFIG
	var authPwd, privPwd types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx,
		path.Root("snmp3_credential").AtName("authentication_password"), &authPwd)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx,
		path.Root("snmp3_credential").AtName("privacy_password"), &privPwd)...)

	// Persist private hash so future plans can read it
	createRequest := data.Expand(ctx, &resp.Diagnostics)
	var ap, pp string
	if !authPwd.IsNull() || !privPwd.IsNull() {
		if !authPwd.IsNull() && !authPwd.IsUnknown() {
			ap = authPwd.ValueString()
		}
		if !privPwd.IsNull() && !privPwd.IsUnknown() {
			pp = privPwd.ValueString()
		}
		h := sha256.New()
		h.Write([]byte(ap))
		h.Write([]byte{0})
		h.Write([]byte(pp))
		val := map[string]string{"algo": "sha256", "hash": hex.EncodeToString(h.Sum(nil))}
		b, err := json.Marshal(val)
		if err != nil {
			resp.Diagnostics.AddError("Private State Marshal Error", err.Error())
			return
		}
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "snmp3_secrets_hash", b)...)
		createRequest.Snmp3Credential.PrivacyPassword = &pp
		createRequest.Snmp3Credential.AuthenticationPassword = &ap
	}

	apiRes, _, err := r.client.DNSAPI.
		RecordHostAPI.
		Create(ctx).
		RecordHost(*createRequest).
		ReturnFieldsPlus(readableAttributesForIPAllocation).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create RecordHost, got error: %s", err))
		return
	}

	res := apiRes.CreateRecordHostResponseAsObject.GetResult()
	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, data.ExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while create RecordHost due inherited Extensible attributes, got error: %s", err))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Restore original IPv4 function call attributes
	if savedIPv4FuncCalls != nil {
		data.Ipv4addrs = r.restoreNestedFuncCallAttrs(ctx, data.Ipv4addrs, savedIPv4FuncCalls)
	}

	// Restore original IPv6 function call attributes
	if savedIPv6FuncCalls != nil {
		data.Ipv6addrs = r.restoreNestedFuncCallAttrs(ctx, data.Ipv6addrs, savedIPv6FuncCalls)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IPAllocationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var diags diag.Diagnostics
	var data IPAllocationModel

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

	// Save original IPv4 function call attributes
	savedIPv4FuncCalls := r.saveNestedFuncCallAttrs(data.Ipv4addrs)

	// Save original IPv6 function call attributes
	savedIPv6FuncCalls := r.saveNestedFuncCallAttrs(data.Ipv6addrs)

	apiRes, httpRes, err := r.client.DNSAPI.
		RecordHostAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForIPAllocation).
		ReturnAsObject(1).
		Execute()

	// If the resource is not found, try searching using Extensible Attributes
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound && r.ReadByExtAttrs(ctx, &data, resp) {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read RecordHost, got error: %s", err))
		return
	}

	res := apiRes.GetRecordHostResponseObjectAsResult.GetResult()

	apiTerraformId, ok := (*res.ExtAttrs)[terraformInternalIDEA]
	if !ok {
		apiTerraformId.Value = ""
	}

	stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)

	if associateInternalId == nil {
		if stateExtAttrs == nil {
			resp.Diagnostics.AddError(
				"Missing Internal ID",
				"Unable to read RecordHost because the internal ID (from extattrs_all) is missing or invalid.",
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while reading RecordHost due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Restore original IPv4 function call attributes
	if savedIPv4FuncCalls != nil {
		data.Ipv4addrs = r.restoreNestedFuncCallAttrs(ctx, data.Ipv4addrs, savedIPv4FuncCalls)
	}

	// Restore original IPv6 function call attributes
	if savedIPv6FuncCalls != nil {
		data.Ipv6addrs = r.restoreNestedFuncCallAttrs(ctx, data.Ipv6addrs, savedIPv6FuncCalls)
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IPAllocationResource) ReadByExtAttrs(ctx context.Context, data *IPAllocationModel, resp *resource.ReadResponse) bool {
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

	apiRes, _, err := r.client.DNSAPI.
		RecordHostAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForIPAllocation).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read RecordHost by extattrs, got error: %s", err))
		return true
	}

	results := apiRes.ListRecordHostResponseObject.GetResult()

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

func (r *IPAllocationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data IPAllocationModel

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

	diags = req.State.GetAttribute(ctx, path.Root("extattrs_all"), &data.ExtAttrsAll)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	diags = req.State.GetAttribute(ctx, path.Root("internal_id"), &data.InternalID)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	associateInternalId, diags := req.Private.GetKey(ctx, "associate_internal_id")
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}
	if associateInternalId != nil {
		data.ExtAttrs, diags = AddInternalIDToExtAttrs(ctx, data.ExtAttrs, diags)
		if diags.HasError() {
			return
		}
	}

	// Add Inherited Extensible Attributes
	data.ExtAttrs, diags = AddInheritedExtAttrs(ctx, data.ExtAttrs, data.ExtAttrsAll)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	// Read current state from backend to preserve DHCP settings
	currentApiRes, httpRes, err := r.client.DNSAPI.
		RecordHostAPI.
		Read(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		ReturnFieldsPlus(readableAttributesForIPAllocation).
		ReturnAsObject(1).
		Execute()

	var currentHost dns.RecordHost
	if err != nil {
		// If ref not found, fallback to searching by internal ID
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			foundHost, foundRef, _, errFound := r.findHostByInternalID(ctx, &data)
			if errFound != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to locate RecordHost by internal id after ref not found: %s", errFound))
				return
			}
			if foundHost == nil {
				resp.Diagnostics.AddError("Not Found", "RecordHost not found by ref and no object found with stored internal id.")
				return
			}
			// Update data.Ref to the found ref so subsequent update targets the correct object
			if foundRef != "" {
				data.Ref = types.StringValue(foundRef)
			}
			currentHost = *foundHost
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read current RecordHost for update, got error: %s", err))
			return
		}
	} else {
		// Successfully read by ref
		currentHost = currentApiRes.GetRecordHostResponseObjectAsResult.GetResult()
	}

	updateReq := data.Expand(ctx, &resp.Diagnostics)
	// Preserve other settings
	preserveDHCPSettings(updateReq, &currentHost)

	var authPwd, privPwd types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx,
		path.Root("snmp3_credential").AtName("authentication_password"), &authPwd)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx,
		path.Root("snmp3_credential").AtName("privacy_password"), &privPwd)...)

	// Set the API fields only when provided
	if !authPwd.IsNull() && !authPwd.IsUnknown() {
		ap := authPwd.ValueString()
		updateReq.Snmp3Credential.AuthenticationPassword = &ap
	}
	if !privPwd.IsNull() && !privPwd.IsUnknown() {
		pp := privPwd.ValueString()
		updateReq.Snmp3Credential.PrivacyPassword = &pp
	}

	// Clear fields not allowed in update call
	updateReq.NetworkView = nil

	apiRes, _, err := r.client.DNSAPI.
		RecordHostAPI.
		Update(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		RecordHost(*updateReq).
		ReturnFieldsPlus(readableAttributesForIPAllocation).
		ReturnAsObject(1).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update RecordHost, got error: %s", err))
		return
	}

	res := apiRes.UpdateRecordHostResponseAsObject.GetResult()

	res.ExtAttrs, data.ExtAttrsAll, diags = RemoveInheritedExtAttrs(ctx, planExtAttrs, *res.ExtAttrs)
	if diags.HasError() {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error while update RecordHost due inherited Extensible attributes, got error: %s", diags))
		return
	}

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	if associateInternalId != nil {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", nil)...)
	}
}

func preserveDHCPSettings(updateReq *dns.RecordHost, currentHost *dns.RecordHost) {
	if currentHost == nil || updateReq == nil {
		return
	}

	// Preserve IPv4 DHCP settings
	if len(currentHost.Ipv4addrs) > 0 && len(updateReq.Ipv4addrs) > 0 {
		currentIPv4 := &currentHost.Ipv4addrs[0]
		updateIPv4 := &updateReq.Ipv4addrs[0]

		if currentIPv4.Mac != nil {
			updateIPv4.Mac = currentIPv4.Mac
		}
		if currentIPv4.ConfigureForDhcp != nil {
			updateIPv4.ConfigureForDhcp = currentIPv4.ConfigureForDhcp
		}
		if currentIPv4.MatchClient != nil {
			updateIPv4.MatchClient = currentIPv4.MatchClient
		}
	}

	// Preserve IPv6 DHCP settings
	if len(currentHost.Ipv6addrs) > 0 && len(updateReq.Ipv6addrs) > 0 {
		currentIPv6 := &currentHost.Ipv6addrs[0]
		updateIPv6 := &updateReq.Ipv6addrs[0]

		if currentIPv6.Duid != nil {
			updateIPv6.Duid = currentIPv6.Duid
		} else if currentIPv6.Mac != nil {
			updateIPv6.Mac = currentIPv6.Mac
		}
		if currentIPv6.ConfigureForDhcp != nil {
			updateIPv6.ConfigureForDhcp = currentIPv6.ConfigureForDhcp
		}
		if currentIPv6.MatchClient != nil {
			updateIPv6.MatchClient = currentIPv6.MatchClient
		}
	}
}

func (r *IPAllocationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data IPAllocationModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DNSAPI.
		RecordHostAPI.
		Delete(ctx, utils.ExtractResourceRef(data.Ref.ValueString())).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			// If ref not found, try to locate by internal id and delete using the found ref
			foundRecord, foundRef, _, errFound := r.findHostByInternalID(ctx, &data)
			if errFound != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to locate RecordHost by internal id after delete ref not found: %s", errFound))
				return
			}
			if foundRecord == nil || foundRef == "" {
				// Nothing to delete
				return
			}

			// Attempt delete using the foundRef
			httpResDel, errDel := r.client.DNSAPI.
				RecordHostAPI.
				Delete(ctx, utils.ExtractResourceRef(foundRef)).
				Execute()
			if errDel != nil {
				if httpResDel != nil && httpResDel.StatusCode == http.StatusNotFound {
					return
				}
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete RecordHost (found by internal id), got error: %s", errDel))
				return
			}
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete RecordHost, got error: %s", err))
		return
	}
}

func (r *IPAllocationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("ref"), req.ID)...)
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "associate_internal_id", []byte("true"))...)
}

func (r *IPAllocationResource) findHostByInternalID(ctx context.Context, data *IPAllocationModel) (*dns.RecordHost, string, *http.Response, error) {
	var diags diag.Diagnostics

	if data.ExtAttrsAll.IsNull() {
		// nothing to search by
		return nil, "", nil, nil
	}

	stateExtAttrs := ExpandExtAttrs(ctx, data.ExtAttrsAll, &diags)
	if diags.HasError() {
		return nil, "", nil, fmt.Errorf("error expanding extattrs: %v", diags)
	}

	internalAttr, ok := (*stateExtAttrs)[terraformInternalIDEA]
	if !ok || internalAttr.Value == "" {
		return nil, "", nil, nil
	}

	idMap := map[string]interface{}{
		terraformInternalIDEA: internalAttr.Value,
	}

	apiRes, httpRes, err := r.client.DNSAPI.
		RecordHostAPI.
		List(ctx).
		Extattrfilter(idMap).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForIPAllocation).
		Execute()
	if err != nil {
		return nil, "", httpRes, err
	}

	results := apiRes.ListRecordHostResponseObject.GetResult()
	if len(results) == 0 {
		// not found
		return nil, "", httpRes, nil
	}

	// pick the first match (optionally you can warn if len>1)
	found := results[0]

	var refStr string
	if found.Ref != nil {
		refStr = *found.Ref
	}

	return &found, refStr, httpRes, nil
}

func (r *IPAllocationResource) saveNestedFuncCallAttrs(ipList types.List) []map[string]attr.Value {
	if ipList.IsNull() || ipList.IsUnknown() {
		return nil
	}

	elements := ipList.Elements()
	if len(elements) == 0 {
		return nil
	}

	savedAttrs := make([]map[string]attr.Value, len(elements))

	for i, element := range elements {
		if element.IsNull() || element.IsUnknown() {
			continue
		}

		elementObj := element.(types.Object)
		elementAttrs := elementObj.Attributes()

		if funcCallAttr, exists := elementAttrs["func_call"]; exists && !funcCallAttr.IsNull() && !funcCallAttr.IsUnknown() {
			funcCallObj := funcCallAttr.(types.Object)
			// Save a copy of the original attributes
			savedAttrs[i] = make(map[string]attr.Value)
			maps.Copy(savedAttrs[i], funcCallObj.Attributes())
		}
	}

	return savedAttrs
}

func (r *IPAllocationResource) restoreNestedFuncCallAttrs(ctx context.Context, ipList types.List, savedAttrs []map[string]attr.Value) types.List {
	if ipList.IsNull() || ipList.IsUnknown() || savedAttrs == nil {
		return ipList
	}

	elements := ipList.Elements()
	if len(elements) == 0 || len(savedAttrs) != len(elements) {
		return ipList
	}

	updatedElements := make([]attr.Value, len(elements))
	hasUpdates := false

	for i, element := range elements {
		if element.IsNull() || element.IsUnknown() || savedAttrs[i] == nil {
			updatedElements[i] = element
			continue
		}

		elementObj := element.(types.Object)
		elementAttrs := elementObj.Attributes()

		if _, exists := elementAttrs["func_call"]; exists && len(savedAttrs[i]) > 0 {
			// Restore original FuncCall attributes
			elementAttrs["func_call"] = types.ObjectValueMust(FuncCallAttrTypes, savedAttrs[i])
			updatedElements[i] = types.ObjectValueMust(elementObj.Type(ctx).(types.ObjectType).AttrTypes, elementAttrs)
			hasUpdates = true
		} else {
			updatedElements[i] = element
		}
	}

	if hasUpdates {
		updatedList, _ := types.ListValue(elements[0].(types.Object).Type(ctx), updatedElements)
		return updatedList
	}

	return ipList
}
