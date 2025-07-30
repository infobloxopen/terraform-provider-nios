package dns

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type ZoneDelegatedModel struct {
	Ref                    types.String      `tfsdk:"ref"`
	Address                iptypes.IPAddress `tfsdk:"address"`
	Comment                types.String      `tfsdk:"comment"`
	DelegateTo             types.List        `tfsdk:"delegate_to"`
	DelegatedTtl           types.Int64       `tfsdk:"delegated_ttl"`
	Disable                types.Bool        `tfsdk:"disable"`
	DisplayDomain          types.String      `tfsdk:"display_domain"`
	DnsFqdn                types.String      `tfsdk:"dns_fqdn"`
	EnableRfc2317Exclusion types.Bool        `tfsdk:"enable_rfc2317_exclusion"`
	ExtAttrs               types.Map         `tfsdk:"extattrs"`
	ExtAttrsAll            types.Map         `tfsdk:"extattrs_all"`
	Fqdn                   types.String      `tfsdk:"fqdn"`
	Locked                 types.Bool        `tfsdk:"locked"`
	LockedBy               types.String      `tfsdk:"locked_by"`
	MaskPrefix             types.String      `tfsdk:"mask_prefix"`
	MsAdIntegrated         types.Bool        `tfsdk:"ms_ad_integrated"`
	MsDdnsMode             types.String      `tfsdk:"ms_ddns_mode"`
	MsManaged              types.String      `tfsdk:"ms_managed"`
	MsReadOnly             types.Bool        `tfsdk:"ms_read_only"`
	MsSyncMasterName       types.String      `tfsdk:"ms_sync_master_name"`
	NsGroup                types.String      `tfsdk:"ns_group"`
	Parent                 types.String      `tfsdk:"parent"`
	Prefix                 types.String      `tfsdk:"prefix"`
	UseDelegatedTtl        types.Bool        `tfsdk:"use_delegated_ttl"`
	UsingSrgAssociations   types.Bool        `tfsdk:"using_srg_associations"`
	View                   types.String      `tfsdk:"view"`
	ZoneFormat             types.String      `tfsdk:"zone_format"`
}

var ZoneDelegatedAttrTypes = map[string]attr.Type{
	"ref":                      types.StringType,
	"address":                  iptypes.IPAddressType{},
	"comment":                  types.StringType,
	"delegate_to":              types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneDelegatedDelegateToAttrTypes}},
	"delegated_ttl":            types.Int64Type,
	"disable":                  types.BoolType,
	"display_domain":           types.StringType,
	"dns_fqdn":                 types.StringType,
	"enable_rfc2317_exclusion": types.BoolType,
	"extattrs":                 types.MapType{ElemType: types.StringType},
	"extattrs_all":             types.MapType{ElemType: types.StringType},
	"fqdn":                     types.StringType,
	"locked":                   types.BoolType,
	"locked_by":                types.StringType,
	"mask_prefix":              types.StringType,
	"ms_ad_integrated":         types.BoolType,
	"ms_ddns_mode":             types.StringType,
	"ms_managed":               types.StringType,
	"ms_read_only":             types.BoolType,
	"ms_sync_master_name":      types.StringType,
	"ns_group":                 types.StringType,
	"parent":                   types.StringType,
	"prefix":                   types.StringType,
	"use_delegated_ttl":        types.BoolType,
	"using_srg_associations":   types.BoolType,
	"view":                     types.StringType,
	"zone_format":              types.StringType,
}

var ZoneDelegatedResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"address": schema.StringAttribute{
		CustomType:          iptypes.IPAddressType{},
		Computed:            true,
		MarkdownDescription: "The IP address of the server that is serving this zone.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "Comment for the zone; maximum 256 characters.",
	},
	"delegate_to": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneDelegatedDelegateToResourceSchemaAttributes,
		},
		Required:            true,
		MarkdownDescription: "This provides information for the remote name server that maintains data for the delegated zone. The Infoblox appliance redirects queries for data for the delegated zone to this remote name server.",
	},
	"delegated_ttl": schema.Int64Attribute{
		Optional: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_delegated_ttl")),
		},
		MarkdownDescription: "You can specify the Time to Live (TTL) values of auto-generated NS and glue records for a delegated zone. This value is the number of seconds that data is cached.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a zone is disabled or not. When this is set to False, the zone is enabled.",
	},
	"display_domain": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The displayed name of the DNS zone.",
	},
	"dns_fqdn": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of this DNS zone in punycode format. For a reverse zone, this is in \"address/cidr\" format. For other zones, this is in FQDN format in punycode format.",
	},
	"enable_rfc2317_exclusion": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag controls whether automatic generation of RFC 2317 CNAMEs for delegated reverse zones overwrite existing PTR records. The default behavior is to overwrite all the existing records in the range; this corresponds to \"allow_ptr_creation_in_parent\" set to False. However, when this flag is set to True the existing PTR records are not overwritten.",
	},
	"extattrs": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Extensible attributes associated with the object.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object, including default attributes.",
		ElementType:         types.StringType,
	},
	"fqdn": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.IsValidFQDN(),
		},
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
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("ANY", "NONE", "SECURE"),
		},
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
		Computed:            true,
		MarkdownDescription: "The delegation NS group bound with delegated zone.",
	},
	"parent": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The parent zone of this zone. Note that when searching for reverse zones, the \"in-addr.arpa\" notation should be used.",
	},
	"prefix": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[a-z0-9_\-]+$`),
				"Must be lowercase and cannot contain spaces or uppercase characters",
			),
		},
		MarkdownDescription: "The RFC2317 prefix value of this DNS zone. Use this field only when the netmask is greater than 24 bits; that is, for a mask between 25 and 31 bits. Enter a prefix, such as the name of the allocated address block. The prefix can be alphanumeric characters, such as 128/26 , 128-189 , or sub-B.",
	},
	"use_delegated_ttl": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: delegated_ttl",
	},
	"using_srg_associations": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "This is true if the zone is associated with a shared record group.",
	},
	"view": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
		Default:             stringdefault.StaticString("default"),
		MarkdownDescription: "The name of the DNS view in which the zone resides. Example \"external\".",
	},
	"zone_format": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("FORWARD", "IPV4", "IPV6"),
		},
		Default:             stringdefault.StaticString("FORWARD"),
		MarkdownDescription: "Determines the format of this zone.",
	},
}

func (m *ZoneDelegatedModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dns.ZoneDelegated {
	if m == nil {
		return nil
	}
	to := &dns.ZoneDelegated{
		Comment:                flex.ExpandStringPointer(m.Comment),
		DelegateTo:             flex.ExpandFrameworkListNestedBlock(ctx, m.DelegateTo, diags, ExpandZoneDelegatedDelegateTo),
		DelegatedTtl:           flex.ExpandInt64Pointer(m.DelegatedTtl),
		Disable:                flex.ExpandBoolPointer(m.Disable),
		EnableRfc2317Exclusion: flex.ExpandBoolPointer(m.EnableRfc2317Exclusion),
		ExtAttrs:               ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Locked:                 flex.ExpandBoolPointer(m.Locked),
		MsAdIntegrated:         flex.ExpandBoolPointer(m.MsAdIntegrated),
		MsDdnsMode:             flex.ExpandStringPointer(m.MsDdnsMode),
		NsGroup:                flex.ExpandStringPointer(m.NsGroup),
		Prefix:                 flex.ExpandStringPointer(m.Prefix),
		UseDelegatedTtl:        flex.ExpandBoolPointer(m.UseDelegatedTtl),
	}
	if isCreate {
		to.View = flex.ExpandStringPointer(m.View)
		to.Fqdn = flex.ExpandStringPointer(m.Fqdn)
		to.ZoneFormat = flex.ExpandStringPointer(m.ZoneFormat)
	}
	return to
}

func FlattenZoneDelegated(ctx context.Context, from *dns.ZoneDelegated, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneDelegatedAttrTypes)
	}
	m := ZoneDelegatedModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, ZoneDelegatedAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneDelegatedModel) Flatten(ctx context.Context, from *dns.ZoneDelegated, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneDelegatedModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Address = flex.FlattenIPAddress(from.Address)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DelegateTo = flex.FlattenFrameworkListNestedBlock(ctx, from.DelegateTo, ZoneDelegatedDelegateToAttrTypes, diags, FlattenZoneDelegatedDelegateTo)
	m.DelegatedTtl = flex.FlattenInt64Pointer(from.DelegatedTtl)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DisplayDomain = flex.FlattenStringPointer(from.DisplayDomain)
	m.DnsFqdn = flex.FlattenStringPointer(from.DnsFqdn)
	m.EnableRfc2317Exclusion = types.BoolPointerValue(from.EnableRfc2317Exclusion)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
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
	m.UseDelegatedTtl = types.BoolPointerValue(from.UseDelegatedTtl)
	m.UsingSrgAssociations = types.BoolPointerValue(from.UsingSrgAssociations)
	m.View = flex.FlattenStringPointer(from.View)
	m.ZoneFormat = flex.FlattenStringPointer(from.ZoneFormat)
}
