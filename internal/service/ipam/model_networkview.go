package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type NetworkviewModel struct {
	Ref                  types.String `tfsdk:"ref"`
	AssociatedDnsViews   types.List   `tfsdk:"associated_dns_views"`
	AssociatedMembers    types.List   `tfsdk:"associated_members"`
	CloudInfo            types.Object `tfsdk:"cloud_info"`
	Comment              types.String `tfsdk:"comment"`
	DdnsDnsView          types.String `tfsdk:"ddns_dns_view"`
	DdnsZonePrimaries    types.List   `tfsdk:"ddns_zone_primaries"`
	ExtAttrs             types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll          types.Map    `tfsdk:"extattrs_all"`
	FederatedRealms      types.List   `tfsdk:"federated_realms"`
	InternalForwardZones types.List   `tfsdk:"internal_forward_zones"`
	IsDefault            types.Bool   `tfsdk:"is_default"`
	MgmPrivate           types.Bool   `tfsdk:"mgm_private"`
	MsAdUserData         types.Object `tfsdk:"ms_ad_user_data"`
	Name                 types.String `tfsdk:"name"`
	RemoteForwardZones   types.List   `tfsdk:"remote_forward_zones"`
	RemoteReverseZones   types.List   `tfsdk:"remote_reverse_zones"`
}

var NetworkviewAttrTypes = map[string]attr.Type{
	"ref":                    types.StringType,
	"associated_dns_views":   types.ListType{ElemType: types.StringType},
	"associated_members":     types.ListType{ElemType: types.ObjectType{AttrTypes: NetworkviewAssociatedMembersAttrTypes}},
	"cloud_info":             types.ObjectType{AttrTypes: NetworkviewCloudInfoAttrTypes},
	"comment":                types.StringType,
	"ddns_dns_view":          types.StringType,
	"ddns_zone_primaries":    types.ListType{ElemType: types.ObjectType{AttrTypes: NetworkviewDdnsZonePrimariesAttrTypes}},
	"extattrs":               types.MapType{ElemType: types.StringType},
	"extattrs_all":           types.MapType{ElemType: types.StringType},
	"federated_realms":       types.ListType{ElemType: types.ObjectType{AttrTypes: NetworkviewFederatedRealmsAttrTypes}},
	"internal_forward_zones": types.ListType{ElemType: types.StringType},
	"is_default":             types.BoolType,
	"mgm_private":            types.BoolType,
	"ms_ad_user_data":        types.ObjectType{AttrTypes: NetworkviewMsAdUserDataAttrTypes},
	"name":                   types.StringType,
	"remote_forward_zones":   types.ListType{ElemType: types.ObjectType{AttrTypes: NetworkviewRemoteForwardZonesAttrTypes}},
	"remote_reverse_zones":   types.ListType{ElemType: types.ObjectType{AttrTypes: NetworkviewRemoteReverseZonesAttrTypes}},
}

var NetworkviewResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"associated_dns_views": schema.ListAttribute{
		ElementType: types.StringType,
		Computed:    true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The list of DNS views associated with this network view.",
	},
	"associated_members": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetworkviewAssociatedMembersResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: "The list of members associated with a network view.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          NetworkviewCloudInfoResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The cloud information associated with the view.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the network view; maximum 256 characters.",
	},
	"ddns_dns_view": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "DNS views that will receive the updates if you enable the appliance to send updates to Grid members.",
	},
	"ddns_zone_primaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetworkviewDdnsZonePrimariesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "An array of Ddns Zone Primary dhcpddns structs that lists the information of primary zone to wich DDNS updates should be sent.",
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
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"federated_realms": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetworkviewFederatedRealmsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "This field contains the federated realms associated to this network view",
	},
	"internal_forward_zones": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Default:     listdefault.StaticValue(types.ListNull(types.StringType)),
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The list of linked authoritative DNS zones.",
	},
	"is_default": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "The NIOS appliance provides one default network view. You can rename the default view and change its settings, but you cannot delete it. There must always be at least one network view in the appliance.",
	},
	"mgm_private": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This field controls whether this object is synchronized with the Multi-Grid Master. If this field is set to True, objects are not synchronized.",
	},
	"ms_ad_user_data": schema.SingleNestedAttribute{
		Attributes:          NetworkviewMsAdUserDataResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The Microsoft Active Directory user related information.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Name of the network view.",
	},
	"remote_forward_zones": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetworkviewRemoteForwardZonesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The list of forward-mapping zones to which the DHCP server sends the updates.",
	},
	"remote_reverse_zones": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetworkviewRemoteReverseZonesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The list of reverse-mapping zones to which the DHCP server sends the updates.",
	},
}

func ExpandNetworkview(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.Networkview {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkviewModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkviewModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Networkview {
	if m == nil {
		return nil
	}
	to := &ipam.Networkview{
		CloudInfo:            ExpandNetworkviewCloudInfo(ctx, m.CloudInfo, diags),
		Comment:              flex.ExpandStringPointer(m.Comment),
		DdnsDnsView:          flex.ExpandStringPointer(m.DdnsDnsView),
		DdnsZonePrimaries:    flex.ExpandFrameworkListNestedBlock(ctx, m.DdnsZonePrimaries, diags, ExpandNetworkviewDdnsZonePrimaries),
		ExtAttrs:             ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		FederatedRealms:      flex.ExpandFrameworkListNestedBlock(ctx, m.FederatedRealms, diags, ExpandNetworkviewFederatedRealms),
		InternalForwardZones: flex.ExpandFrameworkListString(ctx, m.InternalForwardZones, diags),
		MgmPrivate:           flex.ExpandBoolPointer(m.MgmPrivate),
		Name:                 flex.ExpandStringPointer(m.Name),
		RemoteForwardZones:   flex.ExpandFrameworkListNestedBlock(ctx, m.RemoteForwardZones, diags, ExpandNetworkviewRemoteForwardZones),
		RemoteReverseZones:   flex.ExpandFrameworkListNestedBlock(ctx, m.RemoteReverseZones, diags, ExpandNetworkviewRemoteReverseZones),
	}
	return to
}

func FlattenNetworkview(ctx context.Context, from *ipam.Networkview, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkviewAttrTypes)
	}
	m := NetworkviewModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, NetworkviewAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkviewModel) Flatten(ctx context.Context, from *ipam.Networkview, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkviewModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AssociatedDnsViews = flex.FlattenFrameworkListString(ctx, from.AssociatedDnsViews, diags)
	m.AssociatedMembers = flex.FlattenFrameworkListNestedBlock(ctx, from.AssociatedMembers, NetworkviewAssociatedMembersAttrTypes, diags, FlattenNetworkviewAssociatedMembers)
	m.CloudInfo = FlattenNetworkviewCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DdnsDnsView = flex.FlattenStringPointer(from.DdnsDnsView)
	m.DdnsZonePrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.DdnsZonePrimaries, NetworkviewDdnsZonePrimariesAttrTypes, diags, FlattenNetworkviewDdnsZonePrimaries)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.FederatedRealms = flex.FlattenFrameworkListNestedBlock(ctx, from.FederatedRealms, NetworkviewFederatedRealmsAttrTypes, diags, FlattenNetworkviewFederatedRealms)
	m.InternalForwardZones = flex.FlattenFrameworkListString(ctx, from.InternalForwardZones, diags)
	m.IsDefault = types.BoolPointerValue(from.IsDefault)
	m.MgmPrivate = types.BoolPointerValue(from.MgmPrivate)
	m.MsAdUserData = FlattenNetworkviewMsAdUserData(ctx, from.MsAdUserData, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.RemoteForwardZones = flex.FlattenFrameworkListNestedBlock(ctx, from.RemoteForwardZones, NetworkviewRemoteForwardZonesAttrTypes, diags, FlattenNetworkviewRemoteForwardZones)
	m.RemoteReverseZones = flex.FlattenFrameworkListNestedBlock(ctx, from.RemoteReverseZones, NetworkviewRemoteReverseZonesAttrTypes, diags, FlattenNetworkviewRemoteReverseZones)
}
