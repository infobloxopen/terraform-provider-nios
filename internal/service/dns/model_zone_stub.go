package dns

import (
	"context"

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

type ZoneStubModel struct {
	Ref                  types.String `tfsdk:"ref"`
	Address              types.String `tfsdk:"address"`
	Comment              types.String `tfsdk:"comment"`
	Disable              types.Bool   `tfsdk:"disable"`
	DisableForwarding    types.Bool   `tfsdk:"disable_forwarding"`
	DisplayDomain        types.String `tfsdk:"display_domain"`
	DnsFqdn              types.String `tfsdk:"dns_fqdn"`
	ExtAttrs             types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll          types.Map    `tfsdk:"extattrs_all"`
	ExternalNsGroup      types.String `tfsdk:"external_ns_group"`
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
	SoaEmail             types.String `tfsdk:"soa_email"`
	SoaExpire            types.Int64  `tfsdk:"soa_expire"`
	SoaMname             types.String `tfsdk:"soa_mname"`
	SoaNegativeTtl       types.Int64  `tfsdk:"soa_negative_ttl"`
	SoaRefresh           types.Int64  `tfsdk:"soa_refresh"`
	SoaRetry             types.Int64  `tfsdk:"soa_retry"`
	SoaSerialNumber      types.Int64  `tfsdk:"soa_serial_number"`
	StubFrom             types.List   `tfsdk:"stub_from"`
	StubMembers          types.List   `tfsdk:"stub_members"`
	StubMsservers        types.List   `tfsdk:"stub_msservers"`
	UsingSrgAssociations types.Bool   `tfsdk:"using_srg_associations"`
	View                 types.String `tfsdk:"view"`
	ZoneFormat           types.String `tfsdk:"zone_format"`
}

var ZoneStubAttrTypes = map[string]attr.Type{
	"ref":                    types.StringType,
	"address":                types.StringType,
	"comment":                types.StringType,
	"disable":                types.BoolType,
	"disable_forwarding":     types.BoolType,
	"display_domain":         types.StringType,
	"dns_fqdn":               types.StringType,
	"extattrs":               types.MapType{ElemType: types.StringType},
	"extattrs_all":           types.MapType{ElemType: types.StringType},
	"external_ns_group":      types.StringType,
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
	"soa_email":              types.StringType,
	"soa_expire":             types.Int64Type,
	"soa_mname":              types.StringType,
	"soa_negative_ttl":       types.Int64Type,
	"soa_refresh":            types.Int64Type,
	"soa_retry":              types.Int64Type,
	"soa_serial_number":      types.Int64Type,
	"stub_from":              types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneStubStubFromAttrTypes}},
	"stub_members":           types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneStubStubMembersAttrTypes}},
	"stub_msservers":         types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneStubStubMsserversAttrTypes}},
	"using_srg_associations": types.BoolType,
	"view":                   types.StringType,
	"zone_format":            types.StringType,
}

var ZoneStubResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IP address of the server that is serving this zone.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateNoLeadingOrTrailingWhitespace(),
		},
		MarkdownDescription: "Comment for the zone; maximum 256 characters.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a zone is disabled or not. When this is set to False, the zone is enabled.",
	},
	"disable_forwarding": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the name servers that host the zone should not forward queries that end with the domain name of the zone to any configured forwarders.",
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
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
	},
	"extattrs_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
	},
	"external_ns_group": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateNoLeadingOrTrailingWhitespace(),
		},
		MarkdownDescription: "A forward stub server name server group.",
	},
	"fqdn": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.Any(
				customvalidator.IsValidFQDN(),
				customvalidator.IsValidIPCIDR(),
			),
			stringvalidator.AlsoRequires(path.MatchRoot("stub_from")),
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
		Default:  stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("ANY", "NONE", "SECURE"),
		},
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
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateNoLeadingOrTrailingWhitespace(),
		},
		MarkdownDescription: "A stub member name server group.",
	},
	"parent": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The parent zone of this zone. Note that when searching for reverse zones, the \"in-addr.arpa\" notation should be used.",
	},
	"prefix": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateNoLeadingOrTrailingWhitespace(),
		},
		MarkdownDescription: "The RFC2317 prefix value of this DNS zone. Use this field only when the netmask is greater than 24 bits; that is, for a mask between 25 and 31 bits. Enter a prefix, such as the name of the allocated address block. The prefix can be alphanumeric characters, such as 128/26 , 128-189 , or sub-B.",
	},
	"soa_email": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The SOA email for the zone. This value can be in unicode format.",
	},
	"soa_expire": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "This setting defines the amount of time, in seconds, after which the secondary server stops giving out answers about the zone because the zone data is too old to be useful.",
	},
	"soa_mname": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The SOA mname value for this zone. The Infoblox appliance allows you to change the name of the primary server on the SOA record that is automatically created when you initially configure a zone. Use this method to change the name of the primary server on the SOA record. For example, you may want to hide the primary server for a zone. If your device is named dns1.zone.tld, and for security reasons, you want to show a secondary server called dns2.zone.tld as the primary server. To do so, you would go to dns1.zone.tld zone (being the true primary) and change the primary server on the SOA to dns2.zone.tld to hide the true identity of the real primary server. This value can be in unicode format.",
	},
	"soa_negative_ttl": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The negative Time to Live (TTL) value of the SOA of the zone indicates how long a secondary server can cache data for \"Does Not Respond\" responses.",
	},
	"soa_refresh": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "This indicates the interval at which a secondary server sends a message to the primary server for a zone to check that its data is current, and retrieve fresh data if it is not.",
	},
	"soa_retry": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "This indicates how long a secondary server must wait before attempting to recontact the primary server after a connection failure between the two servers occurs.",
	},
	"soa_serial_number": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The serial number in the SOA record incrementally changes every time the record is modified. The Infoblox appliance allows you to change the serial number (in the SOA record) for the primary server so it is higher than the secondary server, thereby ensuring zone transfers come from the primary server.",
	},
	"stub_from": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneStubStubFromResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The primary servers (masters) of this stub zone.Note that the stealth/tsig_key/tsig_key_alg/tsig_key_name/use_tsig_key_name fields of the struct will be ignored when set in this field.",
	},
	"stub_members": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneStubStubMembersResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Grid member servers of this stub zone. Note that the lead/stealth/grid_replicate/ preferred_primaries/enable_preferred_primaries fields of the struct will be ignored when set in this field.",
	},
	"stub_msservers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneStubStubMsserversResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Microsoft DNS servers of this stub zone. Note that the stealth field of the struct will be ignored when set in this field.",
	},
	"using_srg_associations": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "This is true if the zone is associated with a shared record group.",
	},
	"view": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateNoLeadingOrTrailingWhitespace(),
		},
		Default:             stringdefault.StaticString("default"),
		MarkdownDescription: "The name of the DNS view in which the zone resides. Example \"external\".",
	},
	"zone_format": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("FORWARD"),
		Validators: []validator.String{
			stringvalidator.OneOf("FORWARD", "IPV4", "IPV6"),
		},
		MarkdownDescription: "Determines the format of this zone.",
	},
}

func (m *ZoneStubModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dns.ZoneStub {
	if m == nil {
		return nil
	}
	to := &dns.ZoneStub{
		Comment:           flex.ExpandStringPointer(m.Comment),
		Disable:           flex.ExpandBoolPointer(m.Disable),
		DisableForwarding: flex.ExpandBoolPointer(m.DisableForwarding),
		ExtAttrs:          ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		ExternalNsGroup:   flex.ExpandStringPointer(m.ExternalNsGroup),
		Locked:            flex.ExpandBoolPointer(m.Locked),
		MsAdIntegrated:    flex.ExpandBoolPointer(m.MsAdIntegrated),
		MsDdnsMode:        flex.ExpandStringPointer(m.MsDdnsMode),
		NsGroup:           flex.ExpandStringPointer(m.NsGroup),
		Prefix:            flex.ExpandStringPointer(m.Prefix),
		StubFrom:          flex.ExpandFrameworkListNestedBlock(ctx, m.StubFrom, diags, ExpandZoneStubStubFrom),
		StubMembers:       flex.ExpandFrameworkListNestedBlock(ctx, m.StubMembers, diags, ExpandZoneStubStubMembers),
		StubMsservers:     flex.ExpandFrameworkListNestedBlock(ctx, m.StubMsservers, diags, ExpandZoneStubStubMsservers),
	}

	if isCreate {
		to.Fqdn = flex.ExpandStringPointer(m.Fqdn)
		to.ZoneFormat = flex.ExpandStringPointer(m.ZoneFormat)
		to.View = flex.ExpandStringPointer(m.View)
	}

	return to
}

func FlattenZoneStub(ctx context.Context, from *dns.ZoneStub, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneStubAttrTypes)
	}
	m := ZoneStubModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, ZoneStubAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneStubModel) Flatten(ctx context.Context, from *dns.ZoneStub, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneStubModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DisableForwarding = types.BoolPointerValue(from.DisableForwarding)
	m.DisplayDomain = flex.FlattenStringPointer(from.DisplayDomain)
	m.DnsFqdn = flex.FlattenStringPointer(from.DnsFqdn)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.ExternalNsGroup = flex.FlattenStringPointer(from.ExternalNsGroup)
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
	m.SoaEmail = flex.FlattenStringPointer(from.SoaEmail)
	m.SoaExpire = flex.FlattenInt64Pointer(from.SoaExpire)
	m.SoaMname = flex.FlattenStringPointer(from.SoaMname)
	m.SoaNegativeTtl = flex.FlattenInt64Pointer(from.SoaNegativeTtl)
	m.SoaRefresh = flex.FlattenInt64Pointer(from.SoaRefresh)
	m.SoaRetry = flex.FlattenInt64Pointer(from.SoaRetry)
	m.SoaSerialNumber = flex.FlattenInt64Pointer(from.SoaSerial)
	m.StubFrom = flex.FlattenFrameworkListNestedBlock(ctx, from.StubFrom, ZoneStubStubFromAttrTypes, diags, FlattenZoneStubStubFrom)
	m.StubMembers = flex.FlattenFrameworkListNestedBlock(ctx, from.StubMembers, ZoneStubStubMembersAttrTypes, diags, FlattenZoneStubStubMembers)
	m.StubMsservers = flex.FlattenFrameworkListNestedBlock(ctx, from.StubMsservers, ZoneStubStubMsserversAttrTypes, diags, FlattenZoneStubStubMsservers)
	m.UsingSrgAssociations = types.BoolPointerValue(from.UsingSrgAssociations)
	m.View = flex.FlattenStringPointer(from.View)
	m.ZoneFormat = flex.FlattenStringPointer(from.ZoneFormat)
}
