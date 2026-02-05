package dns

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/config"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForIPAssociation = "aliases,allow_telnet,cli_credentials,cloud_info,comment,configure_for_dns,creation_time,ddns_protected,device_description,device_location,device_type,device_vendor,disable,disable_discovery,dns_aliases,dns_name,extattrs,ipv4addrs,ipv6addrs,last_queried,ms_ad_user_data,name,network_view,rrset_order,snmp3_credential,snmp_credential,ttl,use_cli_credentials,use_dns_ea_inheritance,use_snmp3_credential,use_snmp_credential,use_ttl,view,zone"

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &IPAssociationResource{}
var _ resource.ResourceWithImportState = &IPAssociationResource{}
var _ resource.ResourceWithValidateConfig = &IPAssociationResource{}

func NewIPAssociationResource() resource.Resource {
	return &IPAssociationResource{}
}

// IPAssociationResource defines the resource implementation.
type IPAssociationResource struct {
	client *niosclient.APIClient
}

func (r *IPAssociationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "ip_association"
}

func (r *IPAssociationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "",
		Attributes:          IpAssociationResourceSchemaAttributes,
	}
}

func (r *IPAssociationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *IPAssociationResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data IPAssociationModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if both mac and duid are null or empty and dhcp is enabled
	macEmpty := data.MacAddr.IsNull() || data.MacAddr.ValueString() == ""
	duidEmpty := data.Duid.IsNull() || data.Duid.ValueString() == ""
	configure_for_dhcp := data.ConfigureForDhcp.ValueBool()

	if configure_for_dhcp && macEmpty && duidEmpty {
		resp.Diagnostics.AddError(
			"Invalid Configuration",
			"At least one of 'mac' or 'duid' must be configured.",
		)
	}
}

func (r *IPAssociationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data IPAssociationModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	hostRecord, ref, internalID, _, err := r.getOrFindHostRecord(ctx, &data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to locate host record, got error: %s", err))
		return
	}

	data.Ref = types.StringValue(ref)
	data.InternalID = types.StringValue(internalID)

	// Update the host record with DHCP settings
	updatedHostRec, err := r.updateHostRecord(ctx, hostRecord, &data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to associate host record with DHCP settings, got error: %s", err))
		return
	}

	// Extract DHCP-specific data from response
	data = r.flattenDHCPData(updatedHostRec, data)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IPAssociationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data IPAssociationModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	hostRecord, ref, internalID, _, err := r.getOrFindHostRecord(ctx, &data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to locate host record. Please ensure the allocation exists. If you are importing resources, import the allocation first. Original error: %s", err))
		return
	}

	data.Ref = types.StringValue(ref)
	data.InternalID = types.StringValue(internalID)

	// Update data with current DHCP settings
	data = r.flattenDHCPData(hostRecord, data)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IPAssociationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var diags diag.Diagnostics
	var data IPAssociationModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	diags = req.State.GetAttribute(ctx, path.Root("uuid"), &data.Uuid)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	diags = req.State.GetAttribute(ctx, path.Root("internal_id"), &data.InternalID)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	hostRecord, ref, internalID, _, err := r.getOrFindHostRecord(ctx, &data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to locate host record, got error: %s", err))
		return
	}

	data.Ref = types.StringValue(ref)
	data.InternalID = types.StringValue(internalID)

	updatedHostRec, err := r.updateHostRecord(ctx, hostRecord, &data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update DHCP settings, got error: %s", err))
		return
	}

	// Update data with response
	data = r.flattenDHCPData(updatedHostRec, data)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IPAssociationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data IPAssociationModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	hostRecord, _, _, notFound, err := r.getOrFindHostRecord(ctx, &data)
	if err != nil {
		if notFound {
			// If the host record is already gone, consider it deleted
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to locate host record, got error: %s", err))
		return
	}

	// Clear DHCP settings (reset to defaults) but keep the host record
	clearData := IPAssociationModel{
		MacAddr:          internaltypes.NewMACAddressValue(""),
		Duid:             internaltypes.NewDUIDValue(""),
		ConfigureForDhcp: types.BoolValue(false),
	}

	_, err = r.updateHostRecord(ctx, hostRecord, &clearData)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear DHCP settings, got error: %s", err))
		return
	}
}

func (r *IPAssociationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("ref"), req, resp)
}

func (r *IPAssociationResource) getOrFindHostRecord(ctx context.Context, data *IPAssociationModel) (*dns.RecordHost, string, string, bool, error) {
	// Always try ref first if it exists
	if !data.Ref.IsNull() && !data.Ref.IsUnknown() && data.Ref.ValueString() != "" {
		hostRecord, notFound, err := r.getHostRecordByRef(ctx, data.Ref.ValueString())
		if err == nil {
			internalID, err := r.extractInternalIDFromExtAttrs(hostRecord)
			if err != nil {
				return nil, "", "", notFound, fmt.Errorf("failed to extract internal_id from extensible attributes: %w", err)
			}
			return hostRecord, data.Ref.ValueString(), internalID, notFound, nil
		}
	}

	// Fallback to internal_id search
	if !data.InternalID.IsNull() && !data.InternalID.IsUnknown() && data.InternalID.ValueString() != "" {
		hostRecord, notFound, err := r.getHostRecordByInternalID(ctx, data.InternalID.ValueString())
		if err != nil {
			return nil, "", "", notFound, fmt.Errorf("host record not found by ref or internal_id: %w", err)
		}
		if hostRecord != nil && hostRecord.Ref == nil {
			return nil, "", "", notFound, fmt.Errorf("nil ref found on host record located by internal_id")
		}
		return hostRecord, *hostRecord.Ref, data.InternalID.ValueString(), notFound, nil
	}

	return nil, "", "", false, fmt.Errorf("both ref and internal_id are empty or null")
}

func (r *IPAssociationResource) extractInternalIDFromExtAttrs(hostRecord *dns.RecordHost) (string, error) {
	if hostRecord.ExtAttrs == nil {
		return "", fmt.Errorf("no extensible attributes found")
	}

	extAttrs := *hostRecord.ExtAttrs
	if internalAttr, exists := extAttrs[terraformInternalIDEA]; exists {
		if stringVal, ok := internalAttr.Value.(string); ok && stringVal != "" {
			return stringVal, nil
		}
	}

	return "", fmt.Errorf("terraform internal ID not found in extensible attributes")
}

func (r *IPAssociationResource) getHostRecordByRef(ctx context.Context, ref string) (*dns.RecordHost, bool, error) {
	apiRes, httpRes, err := r.client.DNSAPI.
		RecordHostAPI.
		Read(ctx, data.Uuid.ValueString()).
		ReturnFieldsPlus(readableAttributesForIPAssociation).
		ReturnAsObject(1).
		ProxySearch(config.GetProxySearch()).
		Execute()

	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return nil, true, fmt.Errorf("host record not found with ref: %s", ref)
		}
		return nil, false, fmt.Errorf("failed to read host record by ref %s: %w", ref, err)
	}

	hostRecord := apiRes.GetRecordHostResponseObjectAsResult.GetResult()
	return &hostRecord, false, nil
}

func (r *IPAssociationResource) getHostRecordByInternalID(ctx context.Context, internalID string) (*dns.RecordHost, bool, error) {
	searchFilter := map[string]any{
		terraformInternalIDEA: internalID,
	}

	apiRes, httpRes, err := r.client.DNSAPI.
		RecordHostAPI.
		List(ctx).
		Extattrfilter(searchFilter).
		ReturnAsObject(1).
		ReturnFieldsPlus(readableAttributesForIPAssociation).
		Execute()

	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return nil, true, fmt.Errorf("host record not found with internal_id: %s", internalID)
		}
		return nil, false, fmt.Errorf("failed to search host record by internal_id %s: %w", internalID, err)
	}

	results := apiRes.ListRecordHostResponseObject.GetResult()
	if len(results) == 0 {
		return nil, false, fmt.Errorf("no host record found with internal_id: %s", internalID)
	}

	return &results[0], false, nil
}

func (r *IPAssociationResource) updateHostRecord(ctx context.Context, hostRec *dns.RecordHost, data *IPAssociationModel) (*dns.RecordHost, error) {
	// Build update request preserving all existing settings except DHCP
	updateReq := *hostRec

	// Update IPv4 DHCP settings
	if len(updateReq.Ipv4addrs) > 0 {
		updateReq.Ipv4addrs[0].Mac = flex.ExpandMACAddr(data.MacAddr)
		updateReq.Ipv4addrs[0].ConfigureForDhcp = flex.ExpandBoolPointer(data.ConfigureForDhcp)
	}

	// Update IPv6 DHCP settings
	if len(updateReq.Ipv6addrs) > 0 {
		ipv6 := &updateReq.Ipv6addrs[0]

		match_client := data.MatchClient.ValueString()
		switch match_client {
		case "DUID":
			ipv6.Duid = flex.ExpandDUID(data.Duid)
			ipv6.Mac = nil
		case "MAC_ADDRESS":
			ipv6.Mac = flex.ExpandMACAddr(data.MacAddr)
			ipv6.Duid = nil
		}
		if data.ConfigureForDhcp.ValueBool() && match_client != "" {
			ipv6.MatchClient = &match_client
		}

		ipv6.ConfigureForDhcp = flex.ExpandBoolPointer(data.ConfigureForDhcp)
	}

	// Clear out read-only fields that should not be sent in update
	updateReq.CloudInfo = nil
	updateReq.CreationTime = nil
	updateReq.DnsAliases = nil
	updateReq.DnsName = nil
	updateReq.LastQueried = nil
	updateReq.NetworkView = nil
	updateReq.Zone = nil
	if len(updateReq.Ipv4addrs) > 0 {
		updateReq.Ipv4addrs[0].Host = nil
	}
	if len(updateReq.Ipv6addrs) > 0 {
		updateReq.Ipv6addrs[0].Host = nil
	}
	if updateReq.ConfigureForDns != nil && !*updateReq.ConfigureForDns {
		updateReq.Name = nil
		updateReq.View = nil
	}

	apiRes, _, err := r.client.DNSAPI.
		RecordHostAPI.
		Update(ctx, data.Uuid.ValueString()).
		RecordHost(updateReq).
		ReturnFieldsPlus(readableAttributesForIPAssociation).
		ReturnAsObject(1).
		Execute()

	if err != nil {
		return nil, err
	}

	result := apiRes.UpdateRecordHostResponseAsObject.GetResult()
	return &result, nil
}

func (r *IPAssociationResource) flattenDHCPData(hostRec *dns.RecordHost, data IPAssociationModel) IPAssociationModel {
	// Extract DHCP settings from IPv4 addresses
	if hostRec != nil && len(hostRec.Ipv4addrs) > 0 {
		ipv4 := hostRec.Ipv4addrs[0]
		if ipv4.Mac != nil {
			data.MacAddr = flex.FlattenMACAddr(ipv4.Mac)
		}
		if ipv4.ConfigureForDhcp != nil {
			data.ConfigureForDhcp = types.BoolValue(*ipv4.ConfigureForDhcp)
		}
		if ipv4.MatchClient != nil {
			data.MatchClient = types.StringValue(*ipv4.MatchClient)
		}
	}

	// Extract DHCP settings from IPv6 addresses
	if hostRec != nil && len(hostRec.Ipv6addrs) > 0 {
		ipv6 := hostRec.Ipv6addrs[0]
		if ipv6.Duid != nil {
			data.Duid = flex.FlattenDUID(ipv6.Duid)
		}
		if ipv6.ConfigureForDhcp != nil {
			data.ConfigureForDhcp = types.BoolValue(*ipv6.ConfigureForDhcp)
		}
		if ipv6.MatchClient != nil {
			data.MatchClient = types.StringValue(*ipv6.MatchClient)
		}
	}

	return data
}
