package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

type FiltermacModel struct {
	Ref                         types.String `tfsdk:"ref"`
	Comment                     types.String `tfsdk:"comment"`
	DefaultMacAddressExpiration types.Int64  `tfsdk:"default_mac_address_expiration"`
	Disable                     types.Bool   `tfsdk:"disable"`
	EnforceExpirationTimes      types.Bool   `tfsdk:"enforce_expiration_times"`
	ExtAttrs                    types.Map    `tfsdk:"extattrs"`
	LeaseTime                   types.Int64  `tfsdk:"lease_time"`
	Name                        types.String `tfsdk:"name"`
	NeverExpires                types.Bool   `tfsdk:"never_expires"`
	Options                     types.List   `tfsdk:"options"`
	ReservedForInfoblox         types.String `tfsdk:"reserved_for_infoblox"`
	ExtAttrsAll                 types.Map    `tfsdk:"extattrs_all"`
}

var FiltermacAttrTypes = map[string]attr.Type{
	"ref":                            types.StringType,
	"comment":                        types.StringType,
	"default_mac_address_expiration": types.Int64Type,
	"disable":                        types.BoolType,
	"enforce_expiration_times":       types.BoolType,
	"extattrs":                       types.MapType{ElemType: types.StringType},
	"lease_time":                     types.Int64Type,
	"name":                           types.StringType,
	"never_expires":                  types.BoolType,
	"options":                        types.ListType{ElemType: types.ObjectType{AttrTypes: FiltermacOptionsAttrTypes}},
	"reserved_for_infoblox":          types.StringType,
	"extattrs_all":                   types.MapType{ElemType: types.StringType},
}

var FiltermacResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The descriptive comment of a DHCP MAC Filter object.",
	},
	"default_mac_address_expiration": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.Between(60, 4294967295),
		},
		MarkdownDescription: "The default MAC expiration time of the DHCP MAC Address Filter object. By default, the MAC address filter never expires; otherwise, it is the absolute interval when the MAC address filter expires. The maximum value can extend up to 4294967295 secs. The minimum value is 60 secs (1 min).",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the DHCP Fingerprint object is disabled or not.",
	},
	"enforce_expiration_times": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "The flag to enforce MAC address expiration of the DHCP MAC Address Filter object.",
	},
	"extattrs": schema.MapAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object.",
		ElementType:         types.StringType,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
	},
	"lease_time": schema.Int64Attribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The length of time the DHCP server leases an IP address to a client. The lease time applies to hosts that meet the filter criteria.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of a DHCP MAC Filter object.",
	},
	"never_expires": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines if DHCP MAC Filter never expires or automatically expires.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: FiltermacOptionsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"reserved_for_infoblox": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "This is reserved for writing comments related to the particular MAC address filter. The length of comment cannot exceed 1024 bytes.",
	},
}

func (m *FiltermacModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dhcp.Filtermac {
	if m == nil {
		return nil
	}
	to := &dhcp.Filtermac{
		Comment:                     flex.ExpandStringPointer(m.Comment),
		DefaultMacAddressExpiration: flex.ExpandInt64Pointer(m.DefaultMacAddressExpiration),
		Disable:                     flex.ExpandBoolPointer(m.Disable),
		EnforceExpirationTimes:      flex.ExpandBoolPointer(m.EnforceExpirationTimes),
		ExtAttrs:                    ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		LeaseTime:                   flex.ExpandInt64Pointer(m.LeaseTime),
		Name:                        flex.ExpandStringPointer(m.Name),
		NeverExpires:                flex.ExpandBoolPointer(m.NeverExpires),
		Options:                     flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandFiltermacOptions),
		ReservedForInfoblox:         flex.ExpandStringPointer(m.ReservedForInfoblox),
	}
	return to
}

func FlattenFiltermac(ctx context.Context, from *dhcp.Filtermac, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FiltermacAttrTypes)
	}
	m := FiltermacModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, FiltermacAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FiltermacModel) Flatten(ctx context.Context, from *dhcp.Filtermac, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FiltermacModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DefaultMacAddressExpiration = flex.FlattenInt64Pointer(from.DefaultMacAddressExpiration)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.EnforceExpirationTimes = types.BoolPointerValue(from.EnforceExpirationTimes)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.LeaseTime = flex.FlattenInt64Pointer(from.LeaseTime)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NeverExpires = types.BoolPointerValue(from.NeverExpires)
	planOptions := m.Options
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, FiltermacOptionsAttrTypes, diags, FlattenFiltermacOptions)
	//m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, Ipv6fixedaddresstemplateOptionsAttrTypes, diags, Flatten)
	if !planOptions.IsUnknown() {
		reOrderedOptions, diags := utils.ReorderAndFilterDHCPOptions(ctx, planOptions, m.Options)
		if !diags.HasError() {
			m.Options = reOrderedOptions.(basetypes.ListValue)
		}
	}
	m.ReservedForInfoblox = flex.FlattenStringPointer(from.ReservedForInfoblox)
}
