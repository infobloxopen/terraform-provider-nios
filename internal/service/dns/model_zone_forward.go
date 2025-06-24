package dns

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type ZoneForwardModel struct {
	Ref                  types.String `tfsdk:"ref"`
	Address              types.String `tfsdk:"address"`
	Comment              types.String `tfsdk:"comment"`
	Disable              types.Bool   `tfsdk:"disable"`
	DisableNsGeneration  types.Bool   `tfsdk:"disable_ns_generation"`
	DisplayDomain        types.String `tfsdk:"display_domain"`
	DnsFqdn              types.String `tfsdk:"dns_fqdn"`
	ExtAttrs             types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll          types.Map    `tfsdk:"extattrs_all"`
	ExternalNsGroup      types.String `tfsdk:"external_ns_group"`
	ForwardTo            types.List   `tfsdk:"forward_to"`
	ForwardersOnly       types.Bool   `tfsdk:"forwarders_only"`
	ForwardingServers    types.List   `tfsdk:"forwarding_servers"`
	Fqdn                 types.String `tfsdk:"fqdn"`
	Locked               types.Bool   `tfsdk:"locked"`
	LockedBy             types.String `tfsdk:"locked_by"`
	MaskPrefix           types.String `tfsdk:"mask_prefix"`
	MsAdIntegrated       types.Bool   `tfsdk:"ms_ad_integrated"`
	MsDdnsMode           types.String `tfsdk:"ms_ddns_mode"`
	MsManaged            types.String `tfsdk:"ms_managed"`
	MsReadOnly           types.Bool   `tfsdk:"ms_read_only"`
	MsSyncMasterName     types.String `tfsdk:"ms_sync_master_name"`
	NsGroup              types.String `tfsdk:"ns_group"`
	Parent               types.String `tfsdk:"parent"`
	Prefix               types.String `tfsdk:"prefix"`
	UsingSrgAssociations types.Bool   `tfsdk:"using_srg_associations"`
	View                 types.String `tfsdk:"view"`
	ZoneFormat           types.String `tfsdk:"zone_format"`
}

var ZoneForwardAttrTypes = map[string]attr.Type{
	"ref":                    types.StringType,
	"address":                types.StringType,
	"comment":                types.StringType,
	"disable":                types.BoolType,
	"disable_ns_generation":  types.BoolType,
	"display_domain":         types.StringType,
	"dns_fqdn":               types.StringType,
	"extattrs":               types.MapType{ElemType: types.StringType},
	"extattrs_all":           types.MapType{ElemType: types.StringType},
	"external_ns_group":      types.StringType,
	"forward_to":             types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneForwardForwardToAttrTypes}},
	"forwarders_only":        types.BoolType,
	"forwarding_servers":     types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneForwardForwardingServersAttrTypes}},
	"fqdn":                   types.StringType,
	"locked":                 types.BoolType,
	"locked_by":              types.StringType,
	"mask_prefix":            types.StringType,
	"ms_ad_integrated":       types.BoolType,
	"ms_ddns_mode":           types.StringType,
	"ms_managed":             types.StringType,
	"ms_read_only":           types.BoolType,
	"ms_sync_master_name":    types.StringType,
	"ns_group":               types.StringType,
	"parent":                 types.StringType,
	"prefix":                 types.StringType,
	"using_srg_associations": types.BoolType,
	"view":                   types.StringType,
	"zone_format":            types.StringType,
}

var ZoneForwardResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IP address of the server that is serving this zone.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Comment for the zone; maximum 256 characters.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a zone is disabled or not. When this is set to False, the zone is enabled.",
	},
	"disable_ns_generation": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a auto-generation of NS records in parent zone is disabled or not. When this is set to False, the auto-generation is enabled.",
	},
	"display_domain": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The displayed name of the DNS zone.",
	},
	"dns_fqdn": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of this DNS zone in punycode format. For a reverse zone, this is in \"address/cidr\" format. For other zones, this is in FQDN format in punycode format.",
	},
	"extattrs": schema.MapAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object.",
		ElementType:         types.StringType,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
	},
	"external_ns_group": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "A forward stub server name server group.",
	},
	"forward_to": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneForwardForwardToResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The information for the remote name servers to which you want the Infoblox appliance to forward queries for a specified domain name.",
	},
	"forwarders_only": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the appliance sends queries to forwarders only, and not to other internal or Internet root servers.",
	},
	"forwarding_servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneForwardForwardingServersResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The information for the Grid members to which you want the Infoblox appliance to forward queries for a specified domain name.",
	},
	"fqdn": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of this DNS zone. For a reverse zone, this is in \"address/cidr\" format. For other zones, this is in FQDN format. This value can be in unicode format. Note that for a reverse zone, the corresponding zone_format value should be set.",
	},
	"locked": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If you enable this flag, other administrators cannot make conflicting changes. This is for administration purposes only. The zone will continue to serve DNS data even when it is locked.",
	},
	"locked_by": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of a superuser or the administrator who locked this zone.",
	},
	"mask_prefix": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "IPv4 Netmask or IPv6 prefix for this zone.",
	},
	"ms_ad_integrated": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that determines whether Active Directory is integrated or not. This field is valid only when ms_managed is \"STUB\", \"AUTH_PRIMARY\", or \"AUTH_BOTH\".",
	},
	"ms_ddns_mode": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("NONE"),
		MarkdownDescription: "Determines whether an Active Directory-integrated zone with a Microsoft DNS server as primary allows dynamic updates. Valid values are: \"SECURE\" if the zone allows secure updates only. \"NONE\" if the zone forbids dynamic updates. \"ANY\" if the zone accepts both secure and nonsecure updates. This field is valid only if ms_managed is either \"AUTH_PRIMARY\" or \"AUTH_BOTH\". If the flag ms_ad_integrated is false, the value \"SECURE\" is not allowed.",
	},
	"ms_managed": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The flag that indicates whether the zone is assigned to a Microsoft DNS server. This flag returns the authoritative name server type of the Microsoft DNS server. Valid values are: \"NONE\" if the zone is not assigned to any Microsoft DNS server. \"STUB\" if the zone is assigned to a Microsoft DNS server as a stub zone. \"AUTH_PRIMARY\" if only the primary server of the zone is a Microsoft DNS server. \"AUTH_SECONDARY\" if only the secondary server of the zone is a Microsoft DNS server. \"AUTH_BOTH\" if both the primary and secondary servers of the zone are Microsoft DNS servers.",
	},
	"ms_read_only": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if a Grid member manages the zone served by a Microsoft DNS server in read-only mode. This flag is true when a Grid member manages the zone in read-only mode, false otherwise. When the zone has the ms_read_only flag set to True, no changes can be made to this zone.",
	},
	"ms_sync_master_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of MS synchronization master for this zone.",
	},
	"ns_group": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "A forwarding member name server group.",
	},
	"parent": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The parent zone of this zone. Note that when searching for reverse zones, the \"in-addr.arpa\" notation should be used.",
	},
	"prefix": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The RFC2317 prefix value of this DNS zone. Use this field only when the netmask is greater than 24 bits; that is, for a mask between 25 and 31 bits. Enter a prefix, such as the name of the allocated address block. The prefix can be alphanumeric characters, such as 128/26 , 128-189 , or sub-B.",
	},
	"using_srg_associations": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "This is true if the zone is associated with a shared record group.",
	},
	"view": schema.StringAttribute{
		Optional: true,
		Computed: true,
		//Default:             stringdefault.StaticString("default"),
		MarkdownDescription: "The name of the DNS view in which the zone resides. Example \"external\".",
	},
	"zone_format": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("FORWARD"),
		MarkdownDescription: "Determines the format of this zone.",
	},
}

func ExpandZoneForward(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneForward {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneForwardModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags, true)
}

func (m *ZoneForwardModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dns.ZoneForward {
	if m == nil {
		return nil
	}
	to := &dns.ZoneForward{
		Ref:                 flex.ExpandStringPointer(m.Ref),
		Comment:             flex.ExpandStringPointer(m.Comment),
		Disable:             flex.ExpandBoolPointer(m.Disable),
		DisableNsGeneration: flex.ExpandBoolPointer(m.DisableNsGeneration),
		Extattrs:            ExpandExtAttr(ctx, m.ExtAttrs, diags),
		ExternalNsGroup:     flex.ExpandStringPointer(m.ExternalNsGroup),
		ForwardTo:           flex.ExpandFrameworkListNestedBlock(ctx, m.ForwardTo, diags, ExpandZoneForwardForwardTo),
		ForwardersOnly:      flex.ExpandBoolPointer(m.ForwardersOnly),
		ForwardingServers:   flex.ExpandFrameworkListNestedBlock(ctx, m.ForwardingServers, diags, ExpandZoneForwardForwardingServers),
		//Fqdn:                flex.ExpandStringPointer(m.Fqdn),
		Locked:         flex.ExpandBoolPointer(m.Locked),
		MsAdIntegrated: flex.ExpandBoolPointer(m.MsAdIntegrated),
		MsDdnsMode:     flex.ExpandStringPointer(m.MsDdnsMode),
		NsGroup:        flex.ExpandStringPointer(m.NsGroup),
		Prefix:         flex.ExpandStringPointer(m.Prefix),
		//View:                flex.ExpandStringPointer(m.View),
		//ZoneFormat:          flex.ExpandStringPointer(m.ZoneFormat),
	}
	if isCreate {
		to.Fqdn = flex.ExpandStringPointer(m.Fqdn)
		to.View = flex.ExpandStringPointer(m.View)
		to.ZoneFormat = flex.ExpandStringPointer(m.ZoneFormat)
	}
	return to
}

func FlattenZoneForward(ctx context.Context, from *dns.ZoneForward, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneForwardAttrTypes)
	}
	m := ZoneForwardModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrs = m.ExtAttrsAll
	t, d := types.ObjectValueFrom(ctx, ZoneForwardAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneForwardModel) Flatten(ctx context.Context, from *dns.ZoneForward, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneForwardModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DisableNsGeneration = types.BoolPointerValue(from.DisableNsGeneration)
	m.DisplayDomain = flex.FlattenStringPointer(from.DisplayDomain)
	m.DnsFqdn = flex.FlattenStringPointer(from.DnsFqdn)
	m.ExtAttrsAll = FlattenExtAttr(ctx, *from.Extattrs, diags)
	m.ExternalNsGroup = flex.FlattenStringPointer(from.ExternalNsGroup)
	m.ForwardTo = flex.FlattenFrameworkListNestedBlock(ctx, from.ForwardTo, ZoneForwardForwardToAttrTypes, diags, FlattenZoneForwardForwardTo)
	m.ForwardersOnly = types.BoolPointerValue(from.ForwardersOnly)
	m.ForwardingServers = flex.FlattenFrameworkListNestedBlock(ctx, from.ForwardingServers, ZoneForwardForwardingServersAttrTypes, diags, FlattenZoneForwardForwardingServers)
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.Locked = types.BoolPointerValue(from.Locked)
	m.LockedBy = flex.FlattenStringPointer(from.LockedBy)
	m.MaskPrefix = flex.FlattenStringPointer(from.MaskPrefix)
	m.MsAdIntegrated = types.BoolPointerValue(from.MsAdIntegrated)
	m.MsDdnsMode = flex.FlattenStringPointer(from.MsDdnsMode)
	m.MsManaged = flex.FlattenStringPointer(from.MsManaged)
	m.MsReadOnly = types.BoolPointerValue(from.MsReadOnly)
	m.MsSyncMasterName = flex.FlattenStringPointer(from.MsSyncMasterName)
	m.NsGroup = flex.FlattenStringPointer(from.NsGroup)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.Prefix = flex.FlattenStringPointer(from.Prefix)
	m.UsingSrgAssociations = types.BoolPointerValue(from.UsingSrgAssociations)
	m.View = flex.FlattenStringPointer(from.View)
	m.ZoneFormat = flex.FlattenStringPointer(from.ZoneFormat)
}
