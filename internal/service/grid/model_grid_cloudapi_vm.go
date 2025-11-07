package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type GridCloudapiVmModel struct {
	Ref               types.String `tfsdk:"ref"`
	AvailabilityZone  types.String `tfsdk:"availability_zone"`
	CloudInfo         types.Object `tfsdk:"cloud_info"`
	Comment           types.String `tfsdk:"comment"`
	ElasticIpAddress  types.String `tfsdk:"elastic_ip_address"`
	ExtAttrs          types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll       types.Map    `tfsdk:"extattrs_all"`
	FirstSeen         types.Int64  `tfsdk:"first_seen"`
	Hostname          types.String `tfsdk:"hostname"`
	Id                types.String `tfsdk:"id"`
	KernelId          types.String `tfsdk:"kernel_id"`
	LastSeen          types.Int64  `tfsdk:"last_seen"`
	Name              types.String `tfsdk:"name"`
	NetworkCount      types.Int64  `tfsdk:"network_count"`
	OperatingSystem   types.String `tfsdk:"operating_system"`
	PrimaryMacAddress types.String `tfsdk:"primary_mac_address"`
	SubnetAddress     types.String `tfsdk:"subnet_address"`
	SubnetCidr        types.Int64  `tfsdk:"subnet_cidr"`
	SubnetId          types.String `tfsdk:"subnet_id"`
	TenantName        types.String `tfsdk:"tenant_name"`
	VmType            types.String `tfsdk:"vm_type"`
	VpcAddress        types.String `tfsdk:"vpc_address"`
	VpcCidr           types.Int64  `tfsdk:"vpc_cidr"`
	VpcId             types.String `tfsdk:"vpc_id"`
	VpcName           types.String `tfsdk:"vpc_name"`
}

var GridCloudapiVmAttrTypes = map[string]attr.Type{
	"ref":                 types.StringType,
	"availability_zone":   types.StringType,
	"cloud_info":          types.ObjectType{AttrTypes: GridCloudapiVmCloudInfoAttrTypes},
	"comment":             types.StringType,
	"elastic_ip_address":  types.StringType,
	"extattrs":            types.MapType{ElemType: types.StringType},
	"extattrs_all":        types.MapType{ElemType: types.StringType},
	"first_seen":          types.Int64Type,
	"hostname":            types.StringType,
	"id":                  types.StringType,
	"kernel_id":           types.StringType,
	"last_seen":           types.Int64Type,
	"name":                types.StringType,
	"network_count":       types.Int64Type,
	"operating_system":    types.StringType,
	"primary_mac_address": types.StringType,
	"subnet_address":      types.StringType,
	"subnet_cidr":         types.Int64Type,
	"subnet_id":           types.StringType,
	"tenant_name":         types.StringType,
	"vm_type":             types.StringType,
	"vpc_address":         types.StringType,
	"vpc_cidr":            types.Int64Type,
	"vpc_id":              types.StringType,
	"vpc_name":            types.StringType,
}

var GridCloudapiVmResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"availability_zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Availability zone of the VM.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes: GridCloudapiVmCloudInfoResourceSchemaAttributes,
		Optional:   true,
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Comment for the vm object; maximum 1024 characters.",
	},
	"elastic_ip_address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Elastic IP address associated with the VM's primary interface.",
	},
	"extattrs": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
	},
	"first_seen": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The timestamp when the VM was first seen in the system.",
	},
	"hostname": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Hostname part of the FQDN for the address associated with the VM's primary interface.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Unique ID associated with the VM. This is set only when the VM is first created.",
	},
	"kernel_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Identifier of the kernel that this VM is running; maximum 128 characters.",
	},
	"last_seen": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The timestamp when the last event associated with the VM happened.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Name of the VM.",
	},
	"network_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Number of Networks containing any address associated with this VM.",
	},
	"operating_system": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Guest Operating system that this VM is running; maximum 128 characters.",
	},
	"primary_mac_address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "MAC address associated with the VM's primary interface.",
	},
	"subnet_address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Address of the network that is the container of the address associated with the VM's primary interface.",
	},
	"subnet_cidr": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "CIDR of the network that is the container of the address associated with the VM's primary interface.",
	},
	"subnet_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Subnet ID of the network that is the container of the address associated with the VM's primary interface.",
	},
	"tenant_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Name of the tenant associated with the VM.",
	},
	"vm_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "VM type; maximum 64 characters.",
	},
	"vpc_address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Network address of the parent VPC.",
	},
	"vpc_cidr": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Network CIDR of the parent VPC.",
	},
	"vpc_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Identifier of the parent VPC.",
	},
	"vpc_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Name of the parent VPC.",
	},
}

func ExpandGridCloudapiVm(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridCloudapiVm {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridCloudapiVmModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridCloudapiVmModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridCloudapiVm {
	if m == nil {
		return nil
	}
	to := &grid.GridCloudapiVm{
		Ref:             flex.ExpandStringPointer(m.Ref),
		CloudInfo:       ExpandGridCloudapiVmCloudInfo(ctx, m.CloudInfo, diags),
		Comment:         flex.ExpandStringPointer(m.Comment),
		ExtAttrs:        ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		KernelId:        flex.ExpandStringPointer(m.KernelId),
		Name:            flex.ExpandStringPointer(m.Name),
		OperatingSystem: flex.ExpandStringPointer(m.OperatingSystem),
		VmType:          flex.ExpandStringPointer(m.VmType),
	}
	return to
}

func FlattenGridCloudapiVm(ctx context.Context, from *grid.GridCloudapiVm, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridCloudapiVmAttrTypes)
	}
	m := GridCloudapiVmModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, GridCloudapiVmAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridCloudapiVmModel) Flatten(ctx context.Context, from *grid.GridCloudapiVm, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridCloudapiVmModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AvailabilityZone = flex.FlattenStringPointer(from.AvailabilityZone)
	m.CloudInfo = FlattenGridCloudapiVmCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ElasticIpAddress = flex.FlattenStringPointer(from.ElasticIpAddress)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.FirstSeen = flex.FlattenInt64Pointer(from.FirstSeen)
	m.Hostname = flex.FlattenStringPointer(from.Hostname)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.KernelId = flex.FlattenStringPointer(from.KernelId)
	m.LastSeen = flex.FlattenInt64Pointer(from.LastSeen)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NetworkCount = flex.FlattenInt64Pointer(from.NetworkCount)
	m.OperatingSystem = flex.FlattenStringPointer(from.OperatingSystem)
	m.PrimaryMacAddress = flex.FlattenStringPointer(from.PrimaryMacAddress)
	m.SubnetAddress = flex.FlattenStringPointer(from.SubnetAddress)
	m.SubnetCidr = flex.FlattenInt64Pointer(from.SubnetCidr)
	m.SubnetId = flex.FlattenStringPointer(from.SubnetId)
	m.TenantName = flex.FlattenStringPointer(from.TenantName)
	m.VmType = flex.FlattenStringPointer(from.VmType)
	m.VpcAddress = flex.FlattenStringPointer(from.VpcAddress)
	m.VpcCidr = flex.FlattenInt64Pointer(from.VpcCidr)
	m.VpcId = flex.FlattenStringPointer(from.VpcId)
	m.VpcName = flex.FlattenStringPointer(from.VpcName)
}
