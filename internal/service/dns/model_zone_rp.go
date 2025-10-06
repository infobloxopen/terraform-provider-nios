package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type ZoneRpModel struct {
	Ref                              types.String                             `tfsdk:"ref"`
	Address                          iptypes.IPAddress                        `tfsdk:"address"`
	Comment                          types.String                             `tfsdk:"comment"`
	Disable                          types.Bool                               `tfsdk:"disable"`
	DisplayDomain                    types.String                             `tfsdk:"display_domain"`
	DnsSoaEmail                      types.String                             `tfsdk:"dns_soa_email"`
	ExtAttrs                         types.Map                                `tfsdk:"extattrs"`
	ExtAttrsAll                      types.Map                                `tfsdk:"extattrs_all"`
	ExternalPrimaries                types.List                               `tfsdk:"external_primaries"`
	ExternalSecondaries              types.List                               `tfsdk:"external_secondaries"`
	FireeyeRuleMapping               types.Object                             `tfsdk:"fireeye_rule_mapping"`
	Fqdn                             types.String                             `tfsdk:"fqdn"`
	GridPrimary                      types.List                               `tfsdk:"grid_primary"`
	GridSecondaries                  types.List                               `tfsdk:"grid_secondaries"`
	Locked                           types.Bool                               `tfsdk:"locked"`
	LockedBy                         types.String                             `tfsdk:"locked_by"`
	LogRpz                           types.Bool                               `tfsdk:"log_rpz"`
	MaskPrefix                       types.String                             `tfsdk:"mask_prefix"`
	MemberSoaMnames                  types.List                               `tfsdk:"member_soa_mnames"`
	MemberSoaSerials                 types.List                               `tfsdk:"member_soa_serials"`
	NetworkView                      types.String                             `tfsdk:"network_view"`
	NsGroup                          types.String                             `tfsdk:"ns_group"`
	Parent                           types.String                             `tfsdk:"parent"`
	Prefix                           internaltypes.CaseInsensitiveStringValue `tfsdk:"prefix"`
	PrimaryType                      types.String                             `tfsdk:"primary_type"`
	RecordNamePolicy                 types.String                             `tfsdk:"record_name_policy"`
	RpzDropIpRuleEnabled             types.Bool                               `tfsdk:"rpz_drop_ip_rule_enabled"`
	RpzDropIpRuleMinPrefixLengthIpv4 types.Int64                              `tfsdk:"rpz_drop_ip_rule_min_prefix_length_ipv4"`
	RpzDropIpRuleMinPrefixLengthIpv6 types.Int64                              `tfsdk:"rpz_drop_ip_rule_min_prefix_length_ipv6"`
	RpzLastUpdatedTime               types.Int64                              `tfsdk:"rpz_last_updated_time"`
	RpzPolicy                        types.String                             `tfsdk:"rpz_policy"`
	RpzPriority                      types.Int64                              `tfsdk:"rpz_priority"`
	RpzPriorityEnd                   types.Int64                              `tfsdk:"rpz_priority_end"`
	RpzSeverity                      types.String                             `tfsdk:"rpz_severity"`
	RpzType                          types.String                             `tfsdk:"rpz_type"`
	SetSoaSerialNumber               types.Bool                               `tfsdk:"set_soa_serial_number"`
	SoaDefaultTtl                    types.Int64                              `tfsdk:"soa_default_ttl"`
	SoaEmail                         types.String                             `tfsdk:"soa_email"`
	SoaExpire                        types.Int64                              `tfsdk:"soa_expire"`
	SoaNegativeTtl                   types.Int64                              `tfsdk:"soa_negative_ttl"`
	SoaRefresh                       types.Int64                              `tfsdk:"soa_refresh"`
	SoaRetry                         types.Int64                              `tfsdk:"soa_retry"`
	SoaSerialNumber                  types.Int64                              `tfsdk:"soa_serial_number"`
	SubstituteName                   types.String                             `tfsdk:"substitute_name"`
	UseExternalPrimary               types.Bool                               `tfsdk:"use_external_primary"`
	UseGridZoneTimer                 types.Bool                               `tfsdk:"use_grid_zone_timer"`
	UseLogRpz                        types.Bool                               `tfsdk:"use_log_rpz"`
	UseRecordNamePolicy              types.Bool                               `tfsdk:"use_record_name_policy"`
	UseRpzDropIpRule                 types.Bool                               `tfsdk:"use_rpz_drop_ip_rule"`
	UseSoaEmail                      types.Bool                               `tfsdk:"use_soa_email"`
	View                             types.String                             `tfsdk:"view"`
}

var ZoneRpAttrTypes = map[string]attr.Type{
	"ref":                      types.StringType,
	"address":                  iptypes.IPAddressType{},
	"comment":                  types.StringType,
	"disable":                  types.BoolType,
	"display_domain":           types.StringType,
	"dns_soa_email":            types.StringType,
	"extattrs":                 types.MapType{ElemType: types.StringType},
	"extattrs_all":             types.MapType{ElemType: types.StringType},
	"external_primaries":       types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneRpExternalPrimariesAttrTypes}},
	"external_secondaries":     types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneRpExternalSecondariesAttrTypes}},
	"fireeye_rule_mapping":     types.ObjectType{AttrTypes: ZoneRpFireeyeRuleMappingAttrTypes},
	"fqdn":                     types.StringType,
	"grid_primary":             types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneRpGridPrimaryAttrTypes}},
	"grid_secondaries":         types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneRpGridSecondariesAttrTypes}},
	"locked":                   types.BoolType,
	"locked_by":                types.StringType,
	"log_rpz":                  types.BoolType,
	"mask_prefix":              types.StringType,
	"member_soa_mnames":        types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneRpMemberSoaMnamesAttrTypes}},
	"member_soa_serials":       types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneRpMemberSoaSerialsAttrTypes}},
	"network_view":             types.StringType,
	"ns_group":                 types.StringType,
	"parent":                   types.StringType,
	"prefix":                   internaltypes.CaseInsensitiveString{},
	"primary_type":             types.StringType,
	"record_name_policy":       types.StringType,
	"rpz_drop_ip_rule_enabled": types.BoolType,
	"rpz_drop_ip_rule_min_prefix_length_ipv4": types.Int64Type,
	"rpz_drop_ip_rule_min_prefix_length_ipv6": types.Int64Type,
	"rpz_last_updated_time":                   types.Int64Type,
	"rpz_policy":                              types.StringType,
	"rpz_priority":                            types.Int64Type,
	"rpz_priority_end":                        types.Int64Type,
	"rpz_severity":                            types.StringType,
	"rpz_type":                                types.StringType,
	"set_soa_serial_number":                   types.BoolType,
	"soa_default_ttl":                         types.Int64Type,
	"soa_email":                               types.StringType,
	"soa_expire":                              types.Int64Type,
	"soa_negative_ttl":                        types.Int64Type,
	"soa_refresh":                             types.Int64Type,
	"soa_retry":                               types.Int64Type,
	"soa_serial_number":                       types.Int64Type,
	"substitute_name":                         types.StringType,
	"use_external_primary":                    types.BoolType,
	"use_grid_zone_timer":                     types.BoolType,
	"use_log_rpz":                             types.BoolType,
	"use_record_name_policy":                  types.BoolType,
	"use_rpz_drop_ip_rule":                    types.BoolType,
	"use_soa_email":                           types.BoolType,
	"view":                                    types.StringType,
}

var ZoneRpResourceSchemaAttributes = map[string]schema.Attribute{
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
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Comment for the zone; maximum 256 characters.",
		Default:             stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
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
	"dns_soa_email": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The SOA email for the zone in punycode format.",
	},
	"extattrs": schema.MapAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
	},
	"external_primaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneRpExternalPrimariesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.ConflictsWith(
				path.MatchRoot("ns_group"),
				path.MatchRoot("grid_primary"),
			),
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The list of external primary servers.",
	},
	"external_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneRpExternalSecondariesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.ConflictsWith(path.MatchRoot("ns_group")),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The list of external secondary servers.",
	},
	"fireeye_rule_mapping": schema.SingleNestedAttribute{
		Attributes:          ZoneRpFireeyeRuleMappingResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Rules to map fireeye alerts",
	},
	"fqdn": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.Any(
				customvalidator.IsValidFQDN(),
				customvalidator.IsValidIPCIDR(),
			),
			customvalidator.IsNotArpa(),
		},
		MarkdownDescription: "The name of this DNS zone in FQDN format.",
	},
	"grid_primary": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneRpGridPrimaryResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.ConflictsWith(path.MatchRoot("ns_group"),
				path.MatchRoot("external_primaries"),
			),
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The grid primary servers for this zone.",
	},
	"grid_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneRpGridSecondariesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.ConflictsWith(
				path.MatchRoot("ns_group"),
			),
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The list with Grid members that are secondary servers for this zone.",
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
	"log_rpz": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(true),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_log_rpz")),
		},
		MarkdownDescription: "Determines whether RPZ logging enabled or not at zone level. When this is set to False, the logging is disabled.",
	},
	"mask_prefix": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "IPv4 Netmask or IPv6 prefix for this zone.",
	},
	"member_soa_mnames": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneRpMemberSoaMnamesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The list of per-member SOA MNAME information.",
	},
	"member_soa_serials": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneRpMemberSoaSerialsResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: "The list of per-member SOA serial information.",
	},
	"network_view": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the network view in which this zone resides.",
	},
	"ns_group": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.ConflictsWith(path.MatchRoot("grid_primary")),
		},
		MarkdownDescription: "The name server group that serves DNS for this zone.",
	},
	"parent": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The parent zone of this zone. Note that when searching for reverse zones, the \"in-addr.arpa\" notation should be used.",
	},
	"prefix": schema.StringAttribute{
		CustomType: internaltypes.CaseInsensitiveString{},
		Optional:   true,
		Computed:   true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The RFC2317 prefix value of this DNS zone. Use this field only when the netmask is greater than 24 bits; that is, for a mask between 25 and 31 bits. Enter a prefix, such as the name of the allocated address block. The prefix can be alphanumeric characters, such as 128/26 , 128-189 , or sub-B.",
	},
	"primary_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of the primary server.",
	},
	"record_name_policy": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_record_name_policy")),
		},
		MarkdownDescription: "The hostname policy for records under this zone.",
	},
	"rpz_drop_ip_rule_enabled": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_rpz_drop_ip_rule")),
		},
		MarkdownDescription: "Enables the appliance to ignore RPZ-IP triggers with prefix lengths less than the specified minimum prefix length.",
	},
	"rpz_drop_ip_rule_min_prefix_length_ipv4": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(29),
		Validators: []validator.Int64{
			int64validator.Between(0, 4294967295),
		},
		MarkdownDescription: "The minimum prefix length for IPv4 RPZ-IP triggers. The appliance ignores RPZ-IP triggers with prefix lengths less than the specified minimum IPv4 prefix length.",
	},
	"rpz_drop_ip_rule_min_prefix_length_ipv6": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(112),
		MarkdownDescription: "The minimum prefix length for IPv6 RPZ-IP triggers. The appliance ignores RPZ-IP triggers with prefix lengths less than the specified minimum IPv6 prefix length.",
	},
	"rpz_last_updated_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The timestamp of the last update for zone data.",
	},
	"rpz_policy": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("DISABLED", "GIVEN", "NODATA", "PASSTHRU", "SUBSTITUTE"),
		},
		Default:             stringdefault.StaticString("GIVEN"),
		MarkdownDescription: "The response policy zone override policy.",
	},
	"rpz_priority": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The priority of this response policy zone.",
	},
	"rpz_priority_end": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "This number is for UI to identify the end of qualified zone list.",
	},
	"rpz_severity": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("CRITICAL", "INFORMATIONAL", "MAJOR", "WARNING"),
		},
		Default:             stringdefault.StaticString("MAJOR"),
		MarkdownDescription: "The severity of this response policy zone.",
	},
	"rpz_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("FEED", "FIREEYE", "LOCAL"),
		},
		MarkdownDescription: "The type of rpz zone.",
	},
	"set_soa_serial_number": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		//Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The serial number in the SOA record incrementally changes every time the record is modified. The Infoblox appliance allows you to change the serial number (in the SOA record) for the primary server so it is higher than the secondary server, thereby ensuring zone transfers come from the primary server (as they should). To change the serial number you need to set a new value at \"soa_serial_number\" and pass \"set_soa_serial_number\" as True.",
	},
	"soa_default_ttl": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(
				path.MatchRoot("use_grid_zone_timer"),
				path.MatchRoot("soa_expire"),
				path.MatchRoot("soa_retry"),
				path.MatchRoot("soa_negative_ttl"),
				path.MatchRoot("grid_primary"),
				path.MatchRoot("soa_refresh"),
			),
		},
		MarkdownDescription: "The Time to Live (TTL) value of the SOA record of this zone. This value is the number of seconds that data is cached.",
	},
	"soa_email": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_soa_email")),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The SOA email value for this zone. This value can be in unicode format.",
	},
	"soa_expire": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(
				path.MatchRoot("use_grid_zone_timer"),
				path.MatchRoot("soa_default_ttl"),
				path.MatchRoot("soa_retry"),
				path.MatchRoot("soa_negative_ttl"),
				path.MatchRoot("grid_primary"),
				path.MatchRoot("soa_refresh"),
			),
		},
		MarkdownDescription: "This setting defines the amount of time, in seconds, after which the secondary server stops giving out answers about the zone because the zone data is too old to be useful. The default is one week.",
	},
	"soa_negative_ttl": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(
				path.MatchRoot("use_grid_zone_timer"),
				path.MatchRoot("soa_default_ttl"),
				path.MatchRoot("soa_expire"),
				path.MatchRoot("soa_retry"),
				path.MatchRoot("grid_primary"),
				path.MatchRoot("soa_refresh"),
			),
		},
		MarkdownDescription: "The negative Time to Live (TTL) value of the SOA of the zone indicates how long a secondary server can cache data for \"Does Not Respond\" responses.",
	},
	"soa_refresh": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(
				path.MatchRoot("use_grid_zone_timer"),
				path.MatchRoot("soa_default_ttl"),
				path.MatchRoot("soa_expire"),
				path.MatchRoot("soa_retry"),
				path.MatchRoot("soa_negative_ttl"),
				path.MatchRoot("grid_primary"),
			),
		},
		MarkdownDescription: "This indicates the interval at which a secondary server sends a message to the primary server for a zone to check that its data is current, and retrieve fresh data if it is not.",
	},
	"soa_retry": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(
				path.MatchRoot("use_grid_zone_timer"),
				path.MatchRoot("soa_default_ttl"),
				path.MatchRoot("soa_expire"),
				path.MatchRoot("soa_negative_ttl"),
				path.MatchRoot("grid_primary"),
				path.MatchRoot("soa_refresh"),
			),
		},
		MarkdownDescription: "This indicates how long a secondary server must wait before attempting to recontact the primary server after a connection failure between the two servers occurs.",
	},
	"soa_serial_number": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("set_soa_serial_number")),
		},
		MarkdownDescription: "The serial number in the SOA record incrementally changes every time the record is modified. The Infoblox appliance allows you to change the serial number (in the SOA record) for the primary server so it is higher than the secondary server, thereby ensuring zone transfers come from the primary server (as they should). To change the serial number you need to set a new value at \"soa_serial_number\" and pass \"set_soa_serial_number\" as True.",
	},
	"substitute_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The canonical name of redirect target in substitute policy of response policy zone.",
	},
	"use_external_primary": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag controls whether the zone is using an external primary.",
	},
	"use_grid_zone_timer": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(
				path.MatchRoot("grid_primary"),
			),
		},
		MarkdownDescription: "Use flag for: soa_default_ttl , soa_expire, soa_negative_ttl, soa_refresh, soa_retry",
	},
	"use_log_rpz": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: log_rpz",
	},
	"use_record_name_policy": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: record_name_policy",
	},
	"use_rpz_drop_ip_rule": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: rpz_drop_ip_rule_enabled , rpz_drop_ip_rule_min_prefix_length_ipv4, rpz_drop_ip_rule_min_prefix_length_ipv6",
	},
	"use_soa_email": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Use flag for: soa_email",
	},
	"view": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		Default:             stringdefault.StaticString("default"),
		MarkdownDescription: "The name of the DNS view in which the zone resides. Example \"external\".",
	},
}

func (m *ZoneRpModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dns.ZoneRp {
	if m == nil {
		return nil
	}
	to := &dns.ZoneRp{
		Comment:                          flex.ExpandStringPointer(m.Comment),
		Disable:                          flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:                         ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		ExternalPrimaries:                flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalPrimaries, diags, ExpandZoneRpExternalPrimaries),
		ExternalSecondaries:              flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalSecondaries, diags, ExpandZoneRpExternalSecondaries),
		FireeyeRuleMapping:               ExpandZoneRpFireeyeRuleMapping(ctx, m.FireeyeRuleMapping, diags),
		GridPrimary:                      flex.ExpandFrameworkListNestedBlock(ctx, m.GridPrimary, diags, ExpandZoneRpGridPrimary),
		GridSecondaries:                  flex.ExpandFrameworkListNestedBlock(ctx, m.GridSecondaries, diags, ExpandZoneRpGridSecondaries),
		Locked:                           flex.ExpandBoolPointer(m.Locked),
		LogRpz:                           flex.ExpandBoolPointer(m.LogRpz),
		MemberSoaMnames:                  flex.ExpandFrameworkListNestedBlock(ctx, m.MemberSoaMnames, diags, ExpandZoneRpMemberSoaMnames),
		NsGroup:                          flex.ExpandStringPointer(m.NsGroup),
		Prefix:                           flex.ExpandStringPointer(m.Prefix.StringValue),
		RecordNamePolicy:                 flex.ExpandStringPointer(m.RecordNamePolicy),
		RpzDropIpRuleEnabled:             flex.ExpandBoolPointer(m.RpzDropIpRuleEnabled),
		RpzDropIpRuleMinPrefixLengthIpv4: flex.ExpandInt64Pointer(m.RpzDropIpRuleMinPrefixLengthIpv4),
		RpzDropIpRuleMinPrefixLengthIpv6: flex.ExpandInt64Pointer(m.RpzDropIpRuleMinPrefixLengthIpv6),
		RpzPolicy:                        flex.ExpandStringPointer(m.RpzPolicy),
		RpzSeverity:                      flex.ExpandStringPointer(m.RpzSeverity),
		SetSoaSerialNumber:               flex.ExpandBoolPointer(m.SetSoaSerialNumber),
		SoaDefaultTtl:                    flex.ExpandInt64Pointer(m.SoaDefaultTtl),
		SoaEmail:                         flex.ExpandStringPointer(m.SoaEmail),
		SoaExpire:                        flex.ExpandInt64Pointer(m.SoaExpire),
		SoaNegativeTtl:                   flex.ExpandInt64Pointer(m.SoaNegativeTtl),
		SoaRefresh:                       flex.ExpandInt64Pointer(m.SoaRefresh),
		SoaRetry:                         flex.ExpandInt64Pointer(m.SoaRetry),
		SoaSerial:                        flex.ExpandInt64Pointer(m.SoaSerialNumber),
		SubstituteName:                   flex.ExpandStringPointer(m.SubstituteName),
		UseExternalPrimary:               flex.ExpandBoolPointer(m.UseExternalPrimary),
		UseGridZoneTimer:                 flex.ExpandBoolPointer(m.UseGridZoneTimer),
		UseLogRpz:                        flex.ExpandBoolPointer(m.UseLogRpz),
		UseRecordNamePolicy:              flex.ExpandBoolPointer(m.UseRecordNamePolicy),
		UseRpzDropIpRule:                 flex.ExpandBoolPointer(m.UseRpzDropIpRule),
		UseSoaEmail:                      flex.ExpandBoolPointer(m.UseSoaEmail),
	}

	if isCreate {
		to.Fqdn = flex.ExpandStringPointer(m.Fqdn)
		to.RpzType = flex.ExpandStringPointer(m.RpzType)
		// Zone cannot be moved across views
		to.View = flex.ExpandStringPointer(m.View)
	}

	return to
}

func FlattenZoneRp(ctx context.Context, from *dns.ZoneRp, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneRpAttrTypes)
	}
	m := ZoneRpModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, ZoneRpAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneRpModel) Flatten(ctx context.Context, from *dns.ZoneRp, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneRpModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Address = flex.FlattenIPAddress(from.Address)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DisplayDomain = flex.FlattenStringPointer(from.DisplayDomain)
	m.DnsSoaEmail = flex.FlattenStringPointer(from.DnsSoaEmail)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	planExternalPrimaries := m.ExternalPrimaries
	m.ExternalPrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalPrimaries, ZoneRpExternalPrimariesAttrTypes, diags, FlattenZoneRpExternalPrimaries)
	if !planExternalPrimaries.IsNull() {
		result, diags := utils.CopyFieldFromPlanToRespList(ctx, planExternalPrimaries, m.ExternalPrimaries, "tsig_key_name")
		if !diags.HasError() {
			m.ExternalPrimaries = result.(basetypes.ListValue)
		}
	}
	planExternalSecondaries := m.ExternalSecondaries
	m.ExternalSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalSecondaries, ZoneRpExternalSecondariesAttrTypes, diags, FlattenZoneRpExternalSecondaries)
	if !planExternalSecondaries.IsNull() {
		result, diags := utils.CopyFieldFromPlanToRespList(ctx, planExternalSecondaries, m.ExternalSecondaries, "tsig_key_name")
		if !diags.HasError() {
			m.ExternalSecondaries = result.(basetypes.ListValue)
		}
	}
	m.FireeyeRuleMapping = FlattenZoneRpFireeyeRuleMapping(ctx, from.FireeyeRuleMapping, diags)
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.GridPrimary = flex.FlattenFrameworkListNestedBlock(ctx, from.GridPrimary, ZoneRpGridPrimaryAttrTypes, diags, FlattenZoneRpGridPrimary)
	m.GridSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.GridSecondaries, ZoneRpGridSecondariesAttrTypes, diags, FlattenZoneRpGridSecondaries)
	m.Locked = types.BoolPointerValue(from.Locked)
	m.LockedBy = flex.FlattenStringPointer(from.LockedBy)
	m.LogRpz = types.BoolPointerValue(from.LogRpz)
	m.MaskPrefix = flex.FlattenStringPointer(from.MaskPrefix)
	m.MemberSoaMnames = flex.FlattenFrameworkListNestedBlock(ctx, from.MemberSoaMnames, ZoneRpMemberSoaMnamesAttrTypes, diags, FlattenZoneRpMemberSoaMnames)
	m.MemberSoaSerials = flex.FlattenFrameworkListNestedBlock(ctx, from.MemberSoaSerials, ZoneRpMemberSoaSerialsAttrTypes, diags, FlattenZoneRpMemberSoaSerials)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.NsGroup = flex.FlattenStringPointer(from.NsGroup)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.Prefix.StringValue = flex.FlattenStringPointer(from.Prefix)
	m.PrimaryType = flex.FlattenStringPointer(from.PrimaryType)
	m.RecordNamePolicy = flex.FlattenStringPointer(from.RecordNamePolicy)
	m.RpzDropIpRuleEnabled = types.BoolPointerValue(from.RpzDropIpRuleEnabled)
	m.RpzDropIpRuleMinPrefixLengthIpv4 = flex.FlattenInt64Pointer(from.RpzDropIpRuleMinPrefixLengthIpv4)
	m.RpzDropIpRuleMinPrefixLengthIpv6 = flex.FlattenInt64Pointer(from.RpzDropIpRuleMinPrefixLengthIpv6)
	m.RpzLastUpdatedTime = flex.FlattenInt64Pointer(from.RpzLastUpdatedTime)
	m.RpzPolicy = flex.FlattenStringPointer(from.RpzPolicy)
	m.RpzPriority = flex.FlattenInt64Pointer(from.RpzPriority)
	m.RpzPriorityEnd = flex.FlattenInt64Pointer(from.RpzPriorityEnd)
	m.RpzSeverity = flex.FlattenStringPointer(from.RpzSeverity)
	m.RpzType = flex.FlattenStringPointer(from.RpzType)
	m.SetSoaSerialNumber = types.BoolPointerValue(from.SetSoaSerialNumber)
	m.SoaDefaultTtl = flex.FlattenInt64Pointer(from.SoaDefaultTtl)
	m.SoaEmail = flex.FlattenStringPointer(from.SoaEmail)
	m.SoaExpire = flex.FlattenInt64Pointer(from.SoaExpire)
	m.SoaNegativeTtl = flex.FlattenInt64Pointer(from.SoaNegativeTtl)
	m.SoaRefresh = flex.FlattenInt64Pointer(from.SoaRefresh)
	m.SoaRetry = flex.FlattenInt64Pointer(from.SoaRetry)
	m.SoaSerialNumber = flex.FlattenInt64Pointer(from.SoaSerial)
	m.SubstituteName = flex.FlattenStringPointer(from.SubstituteName)
	m.UseExternalPrimary = types.BoolPointerValue(from.UseExternalPrimary)
	m.UseGridZoneTimer = types.BoolPointerValue(from.UseGridZoneTimer)
	m.UseLogRpz = types.BoolPointerValue(from.UseLogRpz)
	m.UseRecordNamePolicy = types.BoolPointerValue(from.UseRecordNamePolicy)
	m.UseRpzDropIpRule = types.BoolPointerValue(from.UseRpzDropIpRule)
	m.UseSoaEmail = types.BoolPointerValue(from.UseSoaEmail)
	m.View = flex.FlattenStringPointer(from.View)
}
