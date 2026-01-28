package dns

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	planmodifiers "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/immutable"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type IPAllocationModel struct {
	Ref                      types.String                     `tfsdk:"ref"`
	Aliases                  internaltypes.UnorderedListValue `tfsdk:"aliases"`
	AllowTelnet              types.Bool                       `tfsdk:"allow_telnet"`
	CliCredentials           types.List                       `tfsdk:"cli_credentials"`
	CloudInfo                types.Object                     `tfsdk:"cloud_info"`
	Comment                  types.String                     `tfsdk:"comment"`
	ConfigureForDns          types.Bool                       `tfsdk:"configure_for_dns"`
	CreationTime             types.Int64                      `tfsdk:"creation_time"`
	DdnsProtected            types.Bool                       `tfsdk:"ddns_protected"`
	DeviceDescription        types.String                     `tfsdk:"device_description"`
	DeviceLocation           types.String                     `tfsdk:"device_location"`
	DeviceType               types.String                     `tfsdk:"device_type"`
	DeviceVendor             types.String                     `tfsdk:"device_vendor"`
	Disable                  types.Bool                       `tfsdk:"disable"`
	DisableDiscovery         types.Bool                       `tfsdk:"disable_discovery"`
	DnsAliases               types.List                       `tfsdk:"dns_aliases"`
	DnsName                  types.String                     `tfsdk:"dns_name"`
	EnableImmediateDiscovery types.Bool                       `tfsdk:"enable_immediate_discovery"`
	ExtAttrs                 types.Map                        `tfsdk:"extattrs"`
	ExtAttrsAll              types.Map                        `tfsdk:"extattrs_all"`
	InternalID               types.String                     `tfsdk:"internal_id"`
	Ipv4addrs                types.List                       `tfsdk:"ipv4addrs"`
	Ipv6addrs                types.List                       `tfsdk:"ipv6addrs"`
	LastQueried              types.Int64                      `tfsdk:"last_queried"`
	MsAdUserData             types.Object                     `tfsdk:"ms_ad_user_data"`
	Name                     types.String                     `tfsdk:"name"`
	NetworkView              types.String                     `tfsdk:"network_view"`
	RestartIfNeeded          types.Bool                       `tfsdk:"restart_if_needed"`
	RrsetOrder               types.String                     `tfsdk:"rrset_order"`
	Snmp3Credential          types.Object                     `tfsdk:"snmp3_credential"`
	SnmpCredential           types.Object                     `tfsdk:"snmp_credential"`
	Ttl                      types.Int64                      `tfsdk:"ttl"`
	UseCliCredentials        types.Bool                       `tfsdk:"use_cli_credentials"`
	UseDnsEaInheritance      types.Bool                       `tfsdk:"use_dns_ea_inheritance"`
	UseSnmp3Credential       types.Bool                       `tfsdk:"use_snmp3_credential"`
	UseSnmpCredential        types.Bool                       `tfsdk:"use_snmp_credential"`
	UseTtl                   types.Bool                       `tfsdk:"use_ttl"`
	View                     types.String                     `tfsdk:"view"`
	Zone                     types.String                     `tfsdk:"zone"`
}

var IPAllocationAttrTypes = map[string]attr.Type{
	"ref":                        types.StringType,
	"aliases":                    internaltypes.UnorderedListOfStringType,
	"allow_telnet":               types.BoolType,
	"cli_credentials":            types.ListType{ElemType: types.ObjectType{AttrTypes: RecordHostCliCredentialsAttrTypes}},
	"cloud_info":                 types.ObjectType{AttrTypes: RecordHostCloudInfoAttrTypes},
	"comment":                    types.StringType,
	"configure_for_dns":          types.BoolType,
	"creation_time":              types.Int64Type,
	"ddns_protected":             types.BoolType,
	"device_description":         types.StringType,
	"device_location":            types.StringType,
	"device_type":                types.StringType,
	"device_vendor":              types.StringType,
	"disable":                    types.BoolType,
	"disable_discovery":          types.BoolType,
	"dns_aliases":                types.ListType{ElemType: types.StringType},
	"dns_name":                   types.StringType,
	"enable_immediate_discovery": types.BoolType,
	"extattrs":                   types.MapType{ElemType: types.StringType},
	"extattrs_all":               types.MapType{ElemType: types.StringType},
	"internal_id":                types.StringType,
	"ipv4addrs":                  types.ListType{ElemType: types.ObjectType{AttrTypes: RecordHostIpv4addrAttrTypes}},
	"ipv6addrs":                  types.ListType{ElemType: types.ObjectType{AttrTypes: RecordHostIpv6addrAttrTypes}},
	"last_queried":               types.Int64Type,
	"ms_ad_user_data":            types.ObjectType{AttrTypes: RecordHostMsAdUserDataAttrTypes},
	"name":                       types.StringType,
	"network_view":               types.StringType,
	"restart_if_needed":          types.BoolType,
	"rrset_order":                types.StringType,
	"snmp3_credential":           types.ObjectType{AttrTypes: RecordHostSnmp3CredentialAttrTypes},
	"snmp_credential":            types.ObjectType{AttrTypes: RecordHostSnmpCredentialAttrTypes},
	"ttl":                        types.Int64Type,
	"use_cli_credentials":        types.BoolType,
	"use_dns_ea_inheritance":     types.BoolType,
	"use_snmp3_credential":       types.BoolType,
	"use_snmp_credential":        types.BoolType,
	"use_ttl":                    types.BoolType,
	"view":                       types.StringType,
	"zone":                       types.StringType,
}

var IPAllocationResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"aliases": schema.ListAttribute{
		CustomType:          internaltypes.UnorderedListOfStringType,
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "This is a list of aliases for the host. The aliases must be in FQDN format. This value can be in unicode format.",
		Default:             listdefault.StaticValue(types.ListNull(types.StringType)),
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
	},
	"allow_telnet": schema.BoolAttribute{
		Computed:            true, // Setting this as computed only as backend is not setting the value correctly, needs to be fixed in future(temporary workaround)
		MarkdownDescription: "This field controls whether the credential is used for both the Telnet and SSH credentials. If set to False, the credential is used only for SSH.",
	},
	"cli_credentials": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RecordHostCliCredentialsResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The CLI credentials for the host record.",
		Default:             listdefault.StaticValue(types.ListNull(types.ObjectType{AttrTypes: RecordHostCliCredentialsAttrTypes})),
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.AlsoRequires(path.MatchRoot("use_snmp3_credential")),
		},
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          RecordHostCloudInfoResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "Structure containing all cloud API related information for this object.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Comment for the record; maximum 256 characters.",
		Default:             stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
	},
	"configure_for_dns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "When configure_for_dns is false, the host does not have parent zone information.",
		Default:             booldefault.StaticBool(true),
	},
	"creation_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time of the record creation in Epoch seconds format.",
	},
	"ddns_protected": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines if the DDNS updates for this record are allowed or not.",
		Default:             booldefault.StaticBool(false),
	},
	"device_description": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The description of the device.",
		Default:             stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
	},
	"device_location": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The location of the device.",
		Default:             stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
	},
	"device_type": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The type of the device.",
		Default:             stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
	},
	"device_vendor": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The vendor of the device.",
		Default:             stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines if the record is disabled or not. False means that the record is enabled.",
		Default:             booldefault.StaticBool(false),
	},
	"disable_discovery": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines if the discovery for the record is disabled or not. False means that the discovery is enabled.",
		Default:             booldefault.StaticBool(false),
	},
	"dns_aliases": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The list of aliases for the host in punycode format.",
	},
	"dns_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name for a host record in punycode format.",
	},
	"enable_immediate_discovery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the discovery for the record should be immediately enabled.",
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
		MarkdownDescription: "Extensible attributes associated with the object, including default attributes.",
		ElementType:         types.StringType,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"internal_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Internal ID of the object.",
	},
	"ipv4addrs": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RecordHostIpv4addrResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "This is a list of IPv4 Addresses for the host.",
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
	},
	"ipv6addrs": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RecordHostIpv6addrResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "This is a list of IPv6 Addresses for the host.",
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
	},
	"last_queried": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time of the last DNS query in Epoch seconds format.",
	},
	"ms_ad_user_data": schema.SingleNestedAttribute{
		Attributes:          RecordHostMsAdUserDataResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The Microsoft Active Directory user related information.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The host name in FQDN format This value can be in unicode format. Regular expression search is not supported for unicode values.",
	},
	"network_view": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The name of the network view in which the host record resides.",
		Default:             stringdefault.StaticString("default"),
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
	},
	"restart_if_needed": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Restarts the member service.",
	},
	"rrset_order": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The value of this field specifies the order in which resource record sets are returned. The possible values are \"cyclic\", \"random\" and \"fixed\".",
		Default:             stringdefault.StaticString("cyclic"),
		Validators:          []validator.String{stringvalidator.OneOf("cyclic", "random", "fixed")},
	},
	"snmp3_credential": schema.SingleNestedAttribute{
		Attributes:          RecordHostSnmp3CredentialResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "The SNMPv3 credential for this host record.",
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_snmp3_credential")),
		},
	},
	"snmp_credential": schema.SingleNestedAttribute{
		Attributes:          RecordHostSnmpCredentialResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "The SNMP credential for this host record. If set to true, the SNMP credential will override member-level settings.",
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_snmp_credential")),
		},
	},
	"ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Time To Live (TTL) value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached.",
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_ttl")),
		},
	},
	"use_cli_credentials": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If set to true, the CLI credential will override member-level settings.",
		Default:             booldefault.StaticBool(false),
	},
	"use_dns_ea_inheritance": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "When use_dns_ea_inheritance is True, the EA is inherited from associated zone.",
		Default:             booldefault.StaticBool(false),
	},
	"use_snmp3_credential": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines if the SNMPv3 credential should be used for the record.",
		Default:             booldefault.StaticBool(false),
	},
	"use_snmp_credential": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If set to true, the SNMP credential will override member-level settings.",
		Default:             booldefault.StaticBool(false),
	},
	"use_ttl": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Use flag for: ttl",
		Default:             booldefault.StaticBool(false),
	},
	"view": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The name of the DNS view in which the record resides. Example: \"external\".",
		Default:             stringdefault.StaticString("default"),
	},
	"zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the zone in which the record resides. Example: \"zone.com\". If a view is not specified when searching by zone, the default view is used.",
	},
}

//func ExpandRecordHost(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.RecordHost {
//	if o.IsNull() || o.IsUnknown() {
//		return nil
//	}
//	var m IPAllocationModel
//	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
//	if diags.HasError() {
//		return nil
//	}
//	return m.Expand(ctx, diags, nil)
//}

func (m *IPAllocationModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordHost {
	if m == nil {
		return nil
	}
	to := &dns.RecordHost{
		Ref:                      flex.ExpandStringPointer(m.Ref),
		Aliases:                  flex.ExpandFrameworkListString(ctx, m.Aliases, diags),
		AllowTelnet:              flex.ExpandBoolPointer(m.AllowTelnet),
		CliCredentials:           flex.ExpandFrameworkListNestedBlock(ctx, m.CliCredentials, diags, ExpandRecordHostCliCredentials),
		Comment:                  flex.ExpandStringPointer(m.Comment),
		ConfigureForDns:          flex.ExpandBoolPointer(m.ConfigureForDns),
		DdnsProtected:            flex.ExpandBoolPointer(m.DdnsProtected),
		DeviceDescription:        flex.ExpandStringPointer(m.DeviceDescription),
		DeviceLocation:           flex.ExpandStringPointer(m.DeviceLocation),
		DeviceType:               flex.ExpandStringPointer(m.DeviceType),
		DeviceVendor:             flex.ExpandStringPointer(m.DeviceVendor),
		Disable:                  flex.ExpandBoolPointer(m.Disable),
		DisableDiscovery:         flex.ExpandBoolPointer(m.DisableDiscovery),
		EnableImmediateDiscovery: flex.ExpandBoolPointer(m.EnableImmediateDiscovery),
		ExtAttrs:                 ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Ipv4addrs:                flex.ExpandFrameworkListNestedBlock(ctx, m.Ipv4addrs, diags, ExpandRecordHostIpv4addr),
		Ipv6addrs:                flex.ExpandFrameworkListNestedBlock(ctx, m.Ipv6addrs, diags, ExpandRecordHostIpv6addr),
		MsAdUserData:             ExpandRecordHostMsAdUserData(ctx, m.MsAdUserData, diags),
		RestartIfNeeded:          flex.ExpandBoolPointer(m.RestartIfNeeded),
		RrsetOrder:               flex.ExpandStringPointer(m.RrsetOrder),
		Snmp3Credential:          ExpandRecordHostSnmp3Credential(ctx, m.Snmp3Credential, diags),
		SnmpCredential:           ExpandRecordHostSnmpCredential(ctx, m.SnmpCredential, diags),
		Ttl:                      flex.ExpandInt64Pointer(m.Ttl),
		UseCliCredentials:        flex.ExpandBoolPointer(m.UseCliCredentials),
		UseDnsEaInheritance:      flex.ExpandBoolPointer(m.UseDnsEaInheritance),
		UseSnmp3Credential:       flex.ExpandBoolPointer(m.UseSnmp3Credential),
		UseSnmpCredential:        flex.ExpandBoolPointer(m.UseSnmpCredential),
		UseTtl:                   flex.ExpandBoolPointer(m.UseTtl),
		NetworkView:              flex.ExpandStringPointer(m.NetworkView),
	}

	if m.ConfigureForDns.IsUnknown() || m.ConfigureForDns.IsNull() || m.ConfigureForDns.ValueBool() {
		to.View = flex.ExpandStringPointer(m.View)
		to.Name = flex.ExpandStringPointer(m.Name)
	}
	return to
}

func FlattenRecordHost(ctx context.Context, from *dns.RecordHost, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IPAllocationAttrTypes)
	}
	m := IPAllocationModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, IPAllocationAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IPAllocationModel) Flatten(ctx context.Context, from *dns.RecordHost, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IPAllocationModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Aliases = flex.FlattenFrameworkUnorderedList(ctx, types.StringType, from.Aliases, diags)
	m.AllowTelnet = types.BoolPointerValue(from.AllowTelnet)
	m.CliCredentials = flex.FlattenFrameworkListNestedBlock(ctx, from.CliCredentials, RecordHostCliCredentialsAttrTypes, diags, FlattenRecordHostCliCredentials)
	m.CloudInfo = FlattenRecordHostCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ConfigureForDns = types.BoolPointerValue(from.ConfigureForDns)
	m.CreationTime = flex.FlattenInt64Pointer(from.CreationTime)
	m.DdnsProtected = types.BoolPointerValue(from.DdnsProtected)
	m.DeviceDescription = flex.FlattenStringPointer(from.DeviceDescription)
	m.DeviceLocation = flex.FlattenStringPointer(from.DeviceLocation)
	m.DeviceType = flex.FlattenStringPointer(from.DeviceType)
	m.DeviceVendor = flex.FlattenStringPointer(from.DeviceVendor)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DisableDiscovery = types.BoolPointerValue(from.DisableDiscovery)
	m.DnsAliases = flex.FlattenFrameworkListString(ctx, from.DnsAliases, diags)
	m.DnsName = flex.FlattenStringPointer(from.DnsName)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Ipv4addrs = flex.FlattenFrameworkListNestedBlock(ctx, from.Ipv4addrs, RecordHostIpv4addrAttrTypes, diags, FlattenRecordHostIpv4addr)
	m.Ipv6addrs = flex.FlattenFrameworkListNestedBlock(ctx, from.Ipv6addrs, RecordHostIpv6addrAttrTypes, diags, FlattenRecordHostIpv6addr)
	m.LastQueried = flex.FlattenInt64Pointer(from.LastQueried)
	m.MsAdUserData = FlattenRecordHostMsAdUserData(ctx, from.MsAdUserData, diags)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.RrsetOrder = flex.FlattenStringPointer(from.RrsetOrder)
	m.Snmp3Credential = FlattenRecordHostSnmp3Credential(ctx, from.Snmp3Credential, diags)
	m.SnmpCredential = FlattenRecordHostSnmpCredential(ctx, from.SnmpCredential, diags)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.UseCliCredentials = types.BoolPointerValue(from.UseCliCredentials)
	m.UseDnsEaInheritance = types.BoolPointerValue(from.UseDnsEaInheritance)
	m.UseSnmp3Credential = types.BoolPointerValue(from.UseSnmp3Credential)
	m.UseSnmpCredential = types.BoolPointerValue(from.UseSnmpCredential)
	m.UseTtl = types.BoolPointerValue(from.UseTtl)
	m.Zone = flex.FlattenStringPointer(from.Zone)

	if from.ConfigureForDns != nil && *from.ConfigureForDns {
		m.Name = flex.FlattenStringPointer(from.Name)
		m.View = flex.FlattenStringPointer(from.View)
	}
}
