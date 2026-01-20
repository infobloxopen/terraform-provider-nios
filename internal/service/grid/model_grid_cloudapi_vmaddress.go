package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type GridCloudapiVmaddressModel struct {
	Ref                   types.String `tfsdk:"ref"`
	Uuid                  types.String `tfsdk:"uuid"`
	Address               types.String `tfsdk:"address"`
	AddressType           types.String `tfsdk:"address_type"`
	AssociatedIp          types.String `tfsdk:"associated_ip"`
	AssociatedObjectTypes types.List   `tfsdk:"associated_object_types"`
	AssociatedObjects     types.List   `tfsdk:"associated_objects"`
	CloudInfo             types.Object `tfsdk:"cloud_info"`
	DnsNames              types.List   `tfsdk:"dns_names"`
	ElasticAddress        types.String `tfsdk:"elastic_address"`
	InterfaceName         types.String `tfsdk:"interface_name"`
	IsIpv4                types.Bool   `tfsdk:"is_ipv4"`
	MacAddress            types.String `tfsdk:"mac_address"`
	MsAdUserData          types.Object `tfsdk:"ms_ad_user_data"`
	Network               types.String `tfsdk:"network"`
	NetworkView           types.String `tfsdk:"network_view"`
	PortId                types.Int64  `tfsdk:"port_id"`
	PrivateAddress        types.String `tfsdk:"private_address"`
	PrivateHostname       types.String `tfsdk:"private_hostname"`
	PublicAddress         types.String `tfsdk:"public_address"`
	PublicHostname        types.String `tfsdk:"public_hostname"`
	SubnetAddress         types.String `tfsdk:"subnet_address"`
	SubnetCidr            types.Int64  `tfsdk:"subnet_cidr"`
	SubnetId              types.String `tfsdk:"subnet_id"`
	Tenant                types.String `tfsdk:"tenant"`
	VmAvailabilityZone    types.String `tfsdk:"vm_availability_zone"`
	VmComment             types.String `tfsdk:"vm_comment"`
	VmCreationTime        types.Int64  `tfsdk:"vm_creation_time"`
	VmHostname            types.String `tfsdk:"vm_hostname"`
	VmId                  types.String `tfsdk:"vm_id"`
	VmKernelId            types.String `tfsdk:"vm_kernel_id"`
	VmLastUpdateTime      types.Int64  `tfsdk:"vm_last_update_time"`
	VmName                types.String `tfsdk:"vm_name"`
	VmNetworkCount        types.Int64  `tfsdk:"vm_network_count"`
	VmOperatingSystem     types.String `tfsdk:"vm_operating_system"`
	VmType                types.String `tfsdk:"vm_type"`
	VmVpcAddress          types.String `tfsdk:"vm_vpc_address"`
	VmVpcCidr             types.Int64  `tfsdk:"vm_vpc_cidr"`
	VmVpcId               types.String `tfsdk:"vm_vpc_id"`
	VmVpcName             types.String `tfsdk:"vm_vpc_name"`
	VmVpcRef              types.String `tfsdk:"vm_vpcref"`
}

var GridCloudapiVmaddressAttrTypes = map[string]attr.Type{
	"ref":                     types.StringType,
	"uuid":                    types.StringType,
	"address":                 types.StringType,
	"address_type":            types.StringType,
	"associated_ip":           types.StringType,
	"associated_object_types": types.ListType{ElemType: types.StringType},
	"associated_objects":      types.ListType{ElemType: types.StringType},
	"cloud_info":              types.ObjectType{AttrTypes: GridCloudapiVmaddressCloudInfoAttrTypes},
	"dns_names":               types.ListType{ElemType: types.StringType},
	"elastic_address":         types.StringType,
	"interface_name":          types.StringType,
	"is_ipv4":                 types.BoolType,
	"mac_address":             types.StringType,
	"ms_ad_user_data":         types.ObjectType{AttrTypes: GridCloudapiVmaddressMsAdUserDataAttrTypes},
	"network":                 types.StringType,
	"network_view":            types.StringType,
	"port_id":                 types.Int64Type,
	"private_address":         types.StringType,
	"private_hostname":        types.StringType,
	"public_address":          types.StringType,
	"public_hostname":         types.StringType,
	"subnet_address":          types.StringType,
	"subnet_cidr":             types.Int64Type,
	"subnet_id":               types.StringType,
	"tenant":                  types.StringType,
	"vm_availability_zone":    types.StringType,
	"vm_comment":              types.StringType,
	"vm_creation_time":        types.Int64Type,
	"vm_hostname":             types.StringType,
	"vm_id":                   types.StringType,
	"vm_kernel_id":            types.StringType,
	"vm_last_update_time":     types.Int64Type,
	"vm_name":                 types.StringType,
	"vm_network_count":        types.Int64Type,
	"vm_operating_system":     types.StringType,
	"vm_type":                 types.StringType,
	"vm_vpc_address":          types.StringType,
	"vm_vpc_cidr":             types.Int64Type,
	"vm_vpc_id":               types.StringType,
	"vm_vpc_name":             types.StringType,
	"vm_vpcref":               types.StringType,
}

var GridCloudapiVmaddressResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IP address of the interface.",
	},
	"address_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "IP address type (Public, Private, Elastic, Floating, ...).",
	},
	"associated_ip": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Reference to associated IPv4 or IPv6 address.",
	},
	"associated_object_types": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "Array of string denoting the types of underlying objects IPv4/IPv6 - \"A\", \"AAAA\", \"PTR\", \"HOST\", \"FA\", \"RESERVATION\", \"UNMANAGED\" + (\"BULKHOST\", \"DHCP_RANGE\", \"RESERVED_RANGE\", \"LEASE\", \"NETWORK\", \"BROADCAST\", \"PENDING\"),",
	},
	"associated_objects": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "The list of references to the object (Host, Fixed Address, RR, ...) that defines this IP.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes: GridCloudapiVmaddressCloudInfoResourceSchemaAttributes,
		Optional:   true,
	},
	"dns_names": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "The list of all FQDNs associated with the IP address.",
	},
	"elastic_address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Elastic IP address associated with this private address, if this address is a private address; otherwise empty.",
	},
	"interface_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Name of the interface associated with this IP address.",
	},
	"is_ipv4": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates whether the address is IPv4 or IPv6.",
	},
	"mac_address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The MAC address of the interface.",
	},
	"ms_ad_user_data": schema.SingleNestedAttribute{
		Attributes: GridCloudapiVmaddressMsAdUserDataResourceSchemaAttributes,
		Optional:   true,
	},
	"network": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The network to which this address belongs, in IPv4 Address/CIDR format.",
	},
	"network_view": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Network view name of the delegated object.",
	},
	"port_id": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Port identifier of the interface.",
	},
	"private_address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Private IP address associated with this public (or elastic or floating) address, if this address is a public address; otherwise empty.",
	},
	"private_hostname": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host part of the FQDN of this address if this address is a private address; otherwise empty",
	},
	"public_address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Public IP address associated with this private address, if this address is a private address; otherwise empty.",
	},
	"public_hostname": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host part of the FQDN of this address if this address is a public (or elastic or floating) address; otherwise empty",
	},
	"subnet_address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Network address of the subnet that is the container of this address.",
	},
	"subnet_cidr": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "CIDR of the subnet that is the container of this address.",
	},
	"subnet_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Subnet ID that is the container of this address.",
	},
	"tenant": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Cloud API Tenant object.",
	},
	"vm_availability_zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Availability zone of the VM.",
	},
	"vm_comment": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "VM comment.",
	},
	"vm_creation_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Date/time the VM was first created as NIOS object.",
	},
	"vm_hostname": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host part of the FQDN of the address attached to the primary interface.",
	},
	"vm_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The UUID of the Virtual Machine.",
	},
	"vm_kernel_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Kernel ID of the VM that this address is associated with.",
	},
	"vm_last_update_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Last time the VM was updated.",
	},
	"vm_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the Virtual Machine.",
	},
	"vm_network_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Count of networks containing all the addresses of the VM.",
	},
	"vm_operating_system": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Operating system that the VM is running.",
	},
	"vm_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Type of the VM this address is associated with.",
	},
	"vm_vpc_address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Network address of the VPC of the VM that this address is associated with.",
	},
	"vm_vpc_cidr": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "CIDR of the VPC of the VM that this address is associated with.",
	},
	"vm_vpc_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Identifier of the VPC where the VM is defined.",
	},
	"vm_vpc_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Name of the VPC where the VM is defined.",
	},
	"vm_vpcref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Reference to the VPC where the VM is defined.",
	},
}

func ExpandGridCloudapiVmaddress(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridCloudapiVmaddress {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridCloudapiVmaddressModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridCloudapiVmaddressModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridCloudapiVmaddress {
	if m == nil {
		return nil
	}
	to := &grid.GridCloudapiVmaddress{
		Ref:          flex.ExpandStringPointer(m.Ref),
		Uuid:         flex.ExpandStringPointer(m.Uuid),
		CloudInfo:    ExpandGridCloudapiVmaddressCloudInfo(ctx, m.CloudInfo, diags),
		MsAdUserData: ExpandGridCloudapiVmaddressMsAdUserData(ctx, m.MsAdUserData, diags),
	}
	return to
}

func FlattenGridCloudapiVmaddress(ctx context.Context, from *grid.GridCloudapiVmaddress, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridCloudapiVmaddressAttrTypes)
	}
	m := GridCloudapiVmaddressModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridCloudapiVmaddressAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridCloudapiVmaddressModel) Flatten(ctx context.Context, from *grid.GridCloudapiVmaddress, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridCloudapiVmaddressModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.AddressType = flex.FlattenStringPointer(from.AddressType)
	m.AssociatedIp = flex.FlattenStringPointer(from.AssociatedIp)
	m.AssociatedObjectTypes = flex.FlattenFrameworkListString(ctx, from.AssociatedObjectTypes, diags)
	m.AssociatedObjects = flex.FlattenFrameworkListString(ctx, from.AssociatedObjects, diags)
	m.CloudInfo = FlattenGridCloudapiVmaddressCloudInfo(ctx, from.CloudInfo, diags)
	m.DnsNames = flex.FlattenFrameworkListString(ctx, from.DnsNames, diags)
	m.ElasticAddress = flex.FlattenStringPointer(from.ElasticAddress)
	m.InterfaceName = flex.FlattenStringPointer(from.InterfaceName)
	m.IsIpv4 = types.BoolPointerValue(from.IsIpv4)
	m.MacAddress = flex.FlattenStringPointer(from.MacAddress)
	m.MsAdUserData = FlattenGridCloudapiVmaddressMsAdUserData(ctx, from.MsAdUserData, diags)
	m.Network = flex.FlattenStringPointer(from.Network)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.PortId = flex.FlattenInt64Pointer(from.PortId)
	m.PrivateAddress = flex.FlattenStringPointer(from.PrivateAddress)
	m.PrivateHostname = flex.FlattenStringPointer(from.PrivateHostname)
	m.PublicAddress = flex.FlattenStringPointer(from.PublicAddress)
	m.PublicHostname = flex.FlattenStringPointer(from.PublicHostname)
	m.SubnetAddress = flex.FlattenStringPointer(from.SubnetAddress)
	m.SubnetCidr = flex.FlattenInt64Pointer(from.SubnetCidr)
	m.SubnetId = flex.FlattenStringPointer(from.SubnetId)
	m.Tenant = flex.FlattenStringPointer(from.Tenant)
	m.VmAvailabilityZone = flex.FlattenStringPointer(from.VmAvailabilityZone)
	m.VmComment = flex.FlattenStringPointer(from.VmComment)
	m.VmCreationTime = flex.FlattenInt64Pointer(from.VmCreationTime)
	m.VmHostname = flex.FlattenStringPointer(from.VmHostname)
	m.VmId = flex.FlattenStringPointer(from.VmId)
	m.VmKernelId = flex.FlattenStringPointer(from.VmKernelId)
	m.VmLastUpdateTime = flex.FlattenInt64Pointer(from.VmLastUpdateTime)
	m.VmName = flex.FlattenStringPointer(from.VmName)
	m.VmNetworkCount = flex.FlattenInt64Pointer(from.VmNetworkCount)
	m.VmOperatingSystem = flex.FlattenStringPointer(from.VmOperatingSystem)
	m.VmType = flex.FlattenStringPointer(from.VmType)
	m.VmVpcAddress = flex.FlattenStringPointer(from.VmVpcAddress)
	m.VmVpcCidr = flex.FlattenInt64Pointer(from.VmVpcCidr)
	m.VmVpcId = flex.FlattenStringPointer(from.VmVpcId)
	m.VmVpcName = flex.FlattenStringPointer(from.VmVpcName)
	m.VmVpcRef = flex.FlattenStringPointer(from.VmVpcRef)
}
