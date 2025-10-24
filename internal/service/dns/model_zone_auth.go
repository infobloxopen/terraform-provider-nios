package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	planmodifiers "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/immutable"
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

// TODO : Function call support for ms_dc_ns_record_creation

type ZoneAuthModel struct {
	Ref                                     types.String                     `tfsdk:"ref"`
	Address                                 iptypes.IPAddress                `tfsdk:"address"`
	AllowActiveDir                          types.List                       `tfsdk:"allow_active_dir"`
	AllowFixedRrsetOrder                    types.Bool                       `tfsdk:"allow_fixed_rrset_order"`
	AllowGssTsigForUnderscoreZone           types.Bool                       `tfsdk:"allow_gss_tsig_for_underscore_zone"`
	AllowGssTsigZoneUpdates                 types.Bool                       `tfsdk:"allow_gss_tsig_zone_updates"`
	AllowQuery                              types.List                       `tfsdk:"allow_query"`
	AllowTransfer                           types.List                       `tfsdk:"allow_transfer"`
	AllowUpdate                             types.List                       `tfsdk:"allow_update"`
	AllowUpdateForwarding                   types.Bool                       `tfsdk:"allow_update_forwarding"`
	AwsRte53ZoneInfo                        types.Object                     `tfsdk:"aws_rte53_zone_info"`
	CloudInfo                               types.Object                     `tfsdk:"cloud_info"`
	Comment                                 types.String                     `tfsdk:"comment"`
	CopyXferToNotify                        types.Bool                       `tfsdk:"copy_xfer_to_notify"`
	CreatePtrForBulkHosts                   types.Bool                       `tfsdk:"create_ptr_for_bulk_hosts"`
	CreatePtrForHosts                       types.Bool                       `tfsdk:"create_ptr_for_hosts"`
	CreateUnderscoreZones                   types.Bool                       `tfsdk:"create_underscore_zones"`
	DdnsForceCreationTimestampUpdate        types.Bool                       `tfsdk:"ddns_force_creation_timestamp_update"`
	DdnsPrincipalGroup                      types.String                     `tfsdk:"ddns_principal_group"`
	DdnsPrincipalTracking                   types.Bool                       `tfsdk:"ddns_principal_tracking"`
	DdnsRestrictPatterns                    types.Bool                       `tfsdk:"ddns_restrict_patterns"`
	DdnsRestrictPatternsList                internaltypes.UnorderedListValue `tfsdk:"ddns_restrict_patterns_list"`
	DdnsRestrictProtected                   types.Bool                       `tfsdk:"ddns_restrict_protected"`
	DdnsRestrictSecure                      types.Bool                       `tfsdk:"ddns_restrict_secure"`
	DdnsRestrictStatic                      types.Bool                       `tfsdk:"ddns_restrict_static"`
	Disable                                 types.Bool                       `tfsdk:"disable"`
	DisableForwarding                       types.Bool                       `tfsdk:"disable_forwarding"`
	DisplayDomain                           types.String                     `tfsdk:"display_domain"`
	DnsFqdn                                 types.String                     `tfsdk:"dns_fqdn"`
	DnsIntegrityEnable                      types.Bool                       `tfsdk:"dns_integrity_enable"`
	DnsIntegrityFrequency                   types.Int64                      `tfsdk:"dns_integrity_frequency"`
	DnsIntegrityMember                      types.String                     `tfsdk:"dns_integrity_member"`
	DnsIntegrityVerboseLogging              types.Bool                       `tfsdk:"dns_integrity_verbose_logging"`
	DnsSoaEmail                             types.String                     `tfsdk:"dns_soa_email"`
	DnssecKeyParams                         types.Object                     `tfsdk:"dnssec_key_params"`
	DnssecKeys                              types.List                       `tfsdk:"dnssec_keys"`
	DnssecKskRolloverDate                   types.Int64                      `tfsdk:"dnssec_ksk_rollover_date"`
	DnssecZskRolloverDate                   types.Int64                      `tfsdk:"dnssec_zsk_rollover_date"`
	DoHostAbstraction                       types.Bool                       `tfsdk:"do_host_abstraction"`
	EffectiveCheckNamesPolicy               types.String                     `tfsdk:"effective_check_names_policy"`
	EffectiveRecordNamePolicy               types.String                     `tfsdk:"effective_record_name_policy"`
	ExtAttrs                                types.Map                        `tfsdk:"extattrs"`
	ExtAttrsAll                             types.Map                        `tfsdk:"extattrs_all"`
	ExternalPrimaries                       types.List                       `tfsdk:"external_primaries"`
	ExternalSecondaries                     types.List                       `tfsdk:"external_secondaries"`
	Fqdn                                    types.String                     `tfsdk:"fqdn"`
	GridPrimary                             types.List                       `tfsdk:"grid_primary"`
	GridPrimarySharedWithMsParentDelegation types.Bool                       `tfsdk:"grid_primary_shared_with_ms_parent_delegation"`
	GridSecondaries                         types.List                       `tfsdk:"grid_secondaries"`
	ImportFrom                              iptypes.IPAddress                `tfsdk:"import_from"`
	IsDnssecEnabled                         types.Bool                       `tfsdk:"is_dnssec_enabled"`
	IsDnssecSigned                          types.Bool                       `tfsdk:"is_dnssec_signed"`
	IsMultimaster                           types.Bool                       `tfsdk:"is_multimaster"`
	LastQueried                             types.Int64                      `tfsdk:"last_queried"`
	LastQueriedAcl                          types.List                       `tfsdk:"last_queried_acl"`
	Locked                                  types.Bool                       `tfsdk:"locked"`
	LockedBy                                types.String                     `tfsdk:"locked_by"`
	MaskPrefix                              types.String                     `tfsdk:"mask_prefix"`
	MemberSoaMnames                         types.List                       `tfsdk:"member_soa_mnames"`
	MemberSoaSerials                        types.List                       `tfsdk:"member_soa_serials"`
	MsAdIntegrated                          types.Bool                       `tfsdk:"ms_ad_integrated"`
	MsAllowTransfer                         types.List                       `tfsdk:"ms_allow_transfer"`
	MsAllowTransferMode                     types.String                     `tfsdk:"ms_allow_transfer_mode"`
	MsDcNsRecordCreation                    types.List                       `tfsdk:"ms_dc_ns_record_creation"`
	MsDdnsMode                              types.String                     `tfsdk:"ms_ddns_mode"`
	MsManaged                               types.String                     `tfsdk:"ms_managed"`
	MsPrimaries                             types.List                       `tfsdk:"ms_primaries"`
	MsReadOnly                              types.Bool                       `tfsdk:"ms_read_only"`
	MsSecondaries                           types.List                       `tfsdk:"ms_secondaries"`
	MsSyncDisabled                          types.Bool                       `tfsdk:"ms_sync_disabled"`
	MsSyncMasterName                        types.String                     `tfsdk:"ms_sync_master_name"`
	NetworkAssociations                     types.List                       `tfsdk:"network_associations"`
	NetworkView                             types.String                     `tfsdk:"network_view"`
	NotifyDelay                             types.Int64                      `tfsdk:"notify_delay"`
	NsGroup                                 types.String                     `tfsdk:"ns_group"`
	Parent                                  types.String                     `tfsdk:"parent"`
	Prefix                                  types.String                     `tfsdk:"prefix"`
	PrimaryType                             types.String                     `tfsdk:"primary_type"`
	RecordNamePolicy                        types.String                     `tfsdk:"record_name_policy"`
	RecordsMonitored                        types.Bool                       `tfsdk:"records_monitored"`
	RemoveSubzones                          types.Bool                       `tfsdk:"remove_subzones"`
	RestartIfNeeded                         types.Bool                       `tfsdk:"restart_if_needed"`
	RrNotQueriedEnabledTime                 types.Int64                      `tfsdk:"rr_not_queried_enabled_time"`
	ScavengingSettings                      types.Object                     `tfsdk:"scavenging_settings"`
	SetSoaSerialNumber                      types.Bool                       `tfsdk:"set_soa_serial_number"`
	SoaDefaultTtl                           types.Int64                      `tfsdk:"soa_default_ttl"`
	SoaEmail                                types.String                     `tfsdk:"soa_email"`
	SoaExpire                               types.Int64                      `tfsdk:"soa_expire"`
	SoaNegativeTtl                          types.Int64                      `tfsdk:"soa_negative_ttl"`
	SoaRefresh                              types.Int64                      `tfsdk:"soa_refresh"`
	SoaRetry                                types.Int64                      `tfsdk:"soa_retry"`
	SoaSerialNumber                         types.Int64                      `tfsdk:"soa_serial_number"`
	Srgs                                    types.List                       `tfsdk:"srgs"`
	UpdateForwarding                        types.List                       `tfsdk:"update_forwarding"`
	UseAllowActiveDir                       types.Bool                       `tfsdk:"use_allow_active_dir"`
	UseAllowQuery                           types.Bool                       `tfsdk:"use_allow_query"`
	UseAllowTransfer                        types.Bool                       `tfsdk:"use_allow_transfer"`
	UseAllowUpdate                          types.Bool                       `tfsdk:"use_allow_update"`
	UseAllowUpdateForwarding                types.Bool                       `tfsdk:"use_allow_update_forwarding"`
	UseCheckNamesPolicy                     types.Bool                       `tfsdk:"use_check_names_policy"`
	UseCopyXferToNotify                     types.Bool                       `tfsdk:"use_copy_xfer_to_notify"`
	UseDdnsForceCreationTimestampUpdate     types.Bool                       `tfsdk:"use_ddns_force_creation_timestamp_update"`
	UseDdnsPatternsRestriction              types.Bool                       `tfsdk:"use_ddns_patterns_restriction"`
	UseDdnsPrincipalSecurity                types.Bool                       `tfsdk:"use_ddns_principal_security"`
	UseDdnsRestrictProtected                types.Bool                       `tfsdk:"use_ddns_restrict_protected"`
	UseDdnsRestrictStatic                   types.Bool                       `tfsdk:"use_ddns_restrict_static"`
	UseDnssecKeyParams                      types.Bool                       `tfsdk:"use_dnssec_key_params"`
	UseExternalPrimary                      types.Bool                       `tfsdk:"use_external_primary"`
	UseGridZoneTimer                        types.Bool                       `tfsdk:"use_grid_zone_timer"`
	UseImportFrom                           types.Bool                       `tfsdk:"use_import_from"`
	UseNotifyDelay                          types.Bool                       `tfsdk:"use_notify_delay"`
	UseRecordNamePolicy                     types.Bool                       `tfsdk:"use_record_name_policy"`
	UseScavengingSettings                   types.Bool                       `tfsdk:"use_scavenging_settings"`
	UseSoaEmail                             types.Bool                       `tfsdk:"use_soa_email"`
	UsingSrgAssociations                    types.Bool                       `tfsdk:"using_srg_associations"`
	View                                    types.String                     `tfsdk:"view"`
	ZoneFormat                              types.String                     `tfsdk:"zone_format"`
	ZoneNotQueriedEnabledTime               types.Int64                      `tfsdk:"zone_not_queried_enabled_time"`
}

var ZoneAuthAttrTypes = map[string]attr.Type{
	"ref":                                  types.StringType,
	"address":                              iptypes.IPAddressType{},
	"allow_active_dir":                     types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthAllowActiveDirAttrTypes}},
	"allow_fixed_rrset_order":              types.BoolType,
	"allow_gss_tsig_for_underscore_zone":   types.BoolType,
	"allow_gss_tsig_zone_updates":          types.BoolType,
	"allow_query":                          types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthAllowQueryAttrTypes}},
	"allow_transfer":                       types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthAllowTransferAttrTypes}},
	"allow_update":                         types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthAllowUpdateAttrTypes}},
	"allow_update_forwarding":              types.BoolType,
	"aws_rte53_zone_info":                  types.ObjectType{AttrTypes: ZoneAuthAwsRte53ZoneInfoAttrTypes},
	"cloud_info":                           types.ObjectType{AttrTypes: ZoneAuthCloudInfoAttrTypes},
	"comment":                              types.StringType,
	"copy_xfer_to_notify":                  types.BoolType,
	"create_ptr_for_bulk_hosts":            types.BoolType,
	"create_ptr_for_hosts":                 types.BoolType,
	"create_underscore_zones":              types.BoolType,
	"ddns_force_creation_timestamp_update": types.BoolType,
	"ddns_principal_group":                 types.StringType,
	"ddns_principal_tracking":              types.BoolType,
	"ddns_restrict_patterns":               types.BoolType,
	"ddns_restrict_patterns_list":          internaltypes.UnorderedListOfStringType,
	"ddns_restrict_protected":              types.BoolType,
	"ddns_restrict_secure":                 types.BoolType,
	"ddns_restrict_static":                 types.BoolType,
	"disable":                              types.BoolType,
	"disable_forwarding":                   types.BoolType,
	"display_domain":                       types.StringType,
	"dns_fqdn":                             types.StringType,
	"dns_integrity_enable":                 types.BoolType,
	"dns_integrity_frequency":              types.Int64Type,
	"dns_integrity_member":                 types.StringType,
	"dns_integrity_verbose_logging":        types.BoolType,
	"dns_soa_email":                        types.StringType,
	"dnssec_key_params":                    types.ObjectType{AttrTypes: ZoneAuthDnssecKeyParamsAttrTypes},
	"dnssec_keys":                          types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthDnssecKeysAttrTypes}},
	"dnssec_ksk_rollover_date":             types.Int64Type,
	"dnssec_zsk_rollover_date":             types.Int64Type,
	"do_host_abstraction":                  types.BoolType,
	"effective_check_names_policy":         types.StringType,
	"effective_record_name_policy":         types.StringType,
	"extattrs":                             types.MapType{ElemType: types.StringType},
	"extattrs_all":                         types.MapType{ElemType: types.StringType},
	"external_primaries":                   types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthExternalPrimariesAttrTypes}},
	"external_secondaries":                 types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthExternalSecondariesAttrTypes}},
	"fqdn":                                 types.StringType,
	"grid_primary":                         types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthGridPrimaryAttrTypes}},
	"grid_primary_shared_with_ms_parent_delegation": types.BoolType,
	"grid_secondaries":            types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthGridSecondariesAttrTypes}},
	"import_from":                 iptypes.IPAddressType{},
	"is_dnssec_enabled":           types.BoolType,
	"is_dnssec_signed":            types.BoolType,
	"is_multimaster":              types.BoolType,
	"last_queried":                types.Int64Type,
	"last_queried_acl":            types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthLastQueriedAclAttrTypes}},
	"locked":                      types.BoolType,
	"locked_by":                   types.StringType,
	"mask_prefix":                 types.StringType,
	"member_soa_mnames":           types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthMemberSoaMnamesAttrTypes}},
	"member_soa_serials":          types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthMemberSoaSerialsAttrTypes}},
	"ms_ad_integrated":            types.BoolType,
	"ms_allow_transfer":           types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthMsAllowTransferAttrTypes}},
	"ms_allow_transfer_mode":      types.StringType,
	"ms_dc_ns_record_creation":    types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthMsDcNsRecordCreationAttrTypes}},
	"ms_ddns_mode":                types.StringType,
	"ms_managed":                  types.StringType,
	"ms_primaries":                types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthMsPrimariesAttrTypes}},
	"ms_read_only":                types.BoolType,
	"ms_secondaries":              types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthMsSecondariesAttrTypes}},
	"ms_sync_disabled":            types.BoolType,
	"ms_sync_master_name":         types.StringType,
	"network_associations":        types.ListType{ElemType: types.StringType},
	"network_view":                types.StringType,
	"notify_delay":                types.Int64Type,
	"ns_group":                    types.StringType,
	"parent":                      types.StringType,
	"prefix":                      types.StringType,
	"primary_type":                types.StringType,
	"record_name_policy":          types.StringType,
	"records_monitored":           types.BoolType,
	"remove_subzones":             types.BoolType,
	"restart_if_needed":           types.BoolType,
	"rr_not_queried_enabled_time": types.Int64Type,
	"scavenging_settings":         types.ObjectType{AttrTypes: ZoneAuthScavengingSettingsAttrTypes},
	"set_soa_serial_number":       types.BoolType,
	"soa_default_ttl":             types.Int64Type,
	"soa_email":                   types.StringType,
	"soa_expire":                  types.Int64Type,
	"soa_negative_ttl":            types.Int64Type,
	"soa_refresh":                 types.Int64Type,
	"soa_retry":                   types.Int64Type,
	"soa_serial_number":           types.Int64Type,
	"srgs":                        types.ListType{ElemType: types.StringType},
	"update_forwarding":           types.ListType{ElemType: types.ObjectType{AttrTypes: ZoneAuthUpdateForwardingAttrTypes}},
	"use_allow_active_dir":        types.BoolType,
	"use_allow_query":             types.BoolType,
	"use_allow_transfer":          types.BoolType,
	"use_allow_update":            types.BoolType,
	"use_allow_update_forwarding": types.BoolType,
	"use_check_names_policy":      types.BoolType,
	"use_copy_xfer_to_notify":     types.BoolType,
	"use_ddns_force_creation_timestamp_update": types.BoolType,
	"use_ddns_patterns_restriction":            types.BoolType,
	"use_ddns_principal_security":              types.BoolType,
	"use_ddns_restrict_protected":              types.BoolType,
	"use_ddns_restrict_static":                 types.BoolType,
	"use_dnssec_key_params":                    types.BoolType,
	"use_external_primary":                     types.BoolType,
	"use_grid_zone_timer":                      types.BoolType,
	"use_import_from":                          types.BoolType,
	"use_notify_delay":                         types.BoolType,
	"use_record_name_policy":                   types.BoolType,
	"use_scavenging_settings":                  types.BoolType,
	"use_soa_email":                            types.BoolType,
	"using_srg_associations":                   types.BoolType,
	"view":                                     types.StringType,
	"zone_format":                              types.StringType,
	"zone_not_queried_enabled_time":            types.Int64Type,
}

var ZoneAuthResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"address": schema.StringAttribute{
		CustomType:          iptypes.IPAddressType{},
		Computed:            true,
		MarkdownDescription: "The IP address of the server that is serving this zone.",
	},
	"allow_active_dir": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthAllowActiveDirResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "This field allows the zone to receive GSS-TSIG authenticated DDNS updates from DHCP clients and servers in an AD domain. Note that addresses specified in this field ignore the permission set in the struct which will be set to 'ALLOW'.",
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_allow_active_dir")),
			listvalidator.SizeAtLeast(1),
		},
	},
	"allow_fixed_rrset_order": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that allows to enable or disable fixed RRset ordering for authoritative forward-mapping zones.",
	},
	"allow_gss_tsig_for_underscore_zone": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that allows DHCP clients to perform GSS-TSIG signed updates for underscore zones.",
	},
	"allow_gss_tsig_zone_updates": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that enables or disables the zone for GSS-TSIG updates.",
	},
	"allow_query": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthAllowQueryResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default: listdefault.StaticValue(
			types.ListNull(types.ObjectType{AttrTypes: ZoneAuthAllowQueryAttrTypes}),
		),
		MarkdownDescription: "Determines whether DNS queries are allowed from a named ACL, or from a list of IPv4/IPv6 addresses, networks, and TSIG keys for the hosts.",
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_allow_query")),
			listvalidator.SizeAtLeast(1),
		},
	},
	"allow_transfer": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthAllowTransferResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default: listdefault.StaticValue(
			types.ListNull(types.ObjectType{AttrTypes: ZoneAuthAllowTransferAttrTypes}),
		),
		MarkdownDescription: "Determines whether zone transfers are allowed from a named ACL, or from a list of IPv4/IPv6 addresses, networks, and TSIG keys for the hosts.",
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_allow_transfer")),
			listvalidator.SizeAtLeast(1),
		},
	},
	"allow_update": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthAllowUpdateResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default: listdefault.StaticValue(
			types.ListNull(types.ObjectType{AttrTypes: ZoneAuthAllowUpdateAttrTypes}),
		),
		MarkdownDescription: "Determines whether dynamic DNS updates are allowed from a named ACL, or from a list of IPv4/IPv6 addresses, networks, and TSIG keys for the hosts.",
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_allow_update")),
			listvalidator.SizeAtLeast(1),
		},
	},
	"allow_update_forwarding": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The list with IP addresses, networks or TSIG keys for clients, from which forwarded dynamic updates are allowed.",
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_allow_update_forwarding")),
		},
	},
	"aws_rte53_zone_info": schema.SingleNestedAttribute{
		Attributes:          ZoneAuthAwsRte53ZoneInfoResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The AWS Route 53 zone information associated with the zone.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          ZoneAuthCloudInfoResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The cloud information associated with the zone.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Comment for the zone; maximum 256 characters.",
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
	},
	"copy_xfer_to_notify": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If this flag is set to True then copy allowed IPs from Allow Transfer to Also Notify.",
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_copy_xfer_to_notify")),
		},
	},
	"create_ptr_for_bulk_hosts": schema.BoolAttribute{
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if PTR records are created for hosts automatically, if necessary, when the zone data is imported. This field is meaningful only when import_from is set.",
	},
	"create_ptr_for_hosts": schema.BoolAttribute{
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if PTR records are created for hosts automatically, if necessary, when the zone data is imported. This field is meaningful only when import_from is set.",
	},
	"create_underscore_zones": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether automatic creation of subzones is enabled or not.",
	},
	"ddns_force_creation_timestamp_update": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Defines whether creation timestamp of RR should be updated ' when DDNS update happens even if there is no change to ' the RR.",
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_force_creation_timestamp_update")),
		},
	},
	"ddns_principal_group": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The DDNS Principal cluster group name.",
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_ddns_principal_security")),
		},
	},
	"ddns_principal_tracking": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that indicates whether the DDNS principal track is enabled or disabled.",
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_principal_security")),
		},
	},
	"ddns_restrict_patterns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that indicates whether an option to restrict DDNS update request based on FQDN patterns is enabled or disabled.",
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_patterns_restriction")),
		},
	},
	"ddns_restrict_patterns_list": schema.ListAttribute{
		CustomType:          internaltypes.UnorderedListOfStringType,
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The unordered list of restriction patterns for an option of to restrict DDNS updates based on FQDN patterns.",
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_ddns_patterns_restriction")),
			listvalidator.SizeAtLeast(1),
		},
	},
	"ddns_restrict_protected": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that indicates whether an option to restrict DDNS update request to protected resource records is enabled or disabled.",
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_restrict_protected")),
		},
	},
	"ddns_restrict_secure": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that indicates whether DDNS update request for principal other than target resource record's principal is restricted.",
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_principal_security")),
		},
	},
	"ddns_restrict_static": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that indicates whether an option to restrict DDNS update request to resource records which are marked as 'STATIC' is enabled or disabled.",
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_restrict_static")),
		},
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
		MarkdownDescription: "Determines whether the name servers that host the zone should forward queries (ended with the domain name of the zone) to any configured forwarders.",
	},
	"display_domain": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The displayed name of the DNS zone.",
	},
	"dns_fqdn": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of this DNS zone in punycode format. For a reverse zone, this is in \"address/cidr\" format. For other zones, this is in FQDN format in punycode format.",
	},
	"dns_integrity_enable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If this is set to True, DNS integrity check is enabled for this zone.",
	},
	"dns_integrity_frequency": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(3600),
		Validators: []validator.Int64{
			int64validator.Between(0, 4294967295),
		},
		MarkdownDescription: "The frequency, in seconds, of DNS integrity checks for this zone.",
	},
	"dns_integrity_member": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Grid member that performs DNS integrity checks for this zone.",
	},
	"dns_integrity_verbose_logging": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If this is set to True, more information is logged for DNS integrity checks for this zone.",
	},
	"dns_soa_email": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The SOA email for the zone in punycode format.",
	},
	"dnssec_key_params": schema.SingleNestedAttribute{
		Attributes:          ZoneAuthDnssecKeyParamsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The DNSSEC key parameters for the zone.",
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_dnssec_key_params")),
		},
	},
	"dnssec_keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthDnssecKeysResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: "A list of DNSSEC keys for the zone.",
	},
	"dnssec_ksk_rollover_date": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The rollover date for the Key Signing Key.",
	},
	"dnssec_zsk_rollover_date": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The rollover date for the Zone Signing Key.",
	},
	"do_host_abstraction": schema.BoolAttribute{
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if hosts and bulk hosts are automatically created when the zone data is imported. This field is meaningful only when import_from is set.",
	},
	"effective_check_names_policy": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("FAIL", "WARN"),
		},
		Default:             stringdefault.StaticString("WARN"),
		MarkdownDescription: "The value of the check names policy, which indicates the action the appliance takes when it encounters host names that do not comply with the Strict Hostname Checking policy. This value applies only if the host name restriction policy is set to \"Strict Hostname Checking\".",
	},
	"effective_record_name_policy": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The selected hostname policy for records under this zone.",
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
			Attributes: ZoneAuthExternalPrimariesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.ConflictsWith(
				path.MatchRoot("ns_group"),
			),
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The list of external primary servers.",
	},
	"external_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthExternalSecondariesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.ConflictsWith(path.MatchRoot("ns_group")),
		},
		MarkdownDescription: "The list of external secondary servers.",
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
		MarkdownDescription: "The name of this DNS zone. For a reverse zone, this is in \"address/cidr\" format. For other zones, this is in FQDN format. This value can be in unicode format. Note that for a reverse zone, the corresponding zone_format value should be set.",
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
	},
	"grid_primary": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthGridPrimaryResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.ConflictsWith(path.MatchRoot("ns_group")),
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The grid primary servers for this zone.",
	},
	"grid_primary_shared_with_ms_parent_delegation": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if the server is duplicated with parent delegation.",
	},
	"grid_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthGridSecondariesResourceSchemaAttributes,
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
	"import_from": schema.StringAttribute{
		CustomType:          iptypes.IPAddressType{},
		Computed:            true,
		MarkdownDescription: "The IP address of the Infoblox appliance from which zone data is imported. Setting this address to '255.255.255.255' and do_host_abstraction to 'true' will create Host records from A records in this zone without importing zone data.",
	},
	"is_dnssec_enabled": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "This flag is set to True if DNSSEC is enabled for the zone.",
	},
	"is_dnssec_signed": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if the zone is DNSSEC signed.",
	},
	"is_multimaster": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if multi-master DNS is enabled for the zone.",
	},
	"last_queried": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time the zone was last queried on.",
	},
	"last_queried_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthLastQueriedAclResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.AlsoRequires(path.MatchRoot("use_scavenging_settings")),
		},
		MarkdownDescription: "Determines last queried ACL for the specified IPv4 or IPv6 addresses and networks in scavenging settings.",
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
	"member_soa_mnames": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthMemberSoaMnamesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The list of per-member SOA MNAME information.",
	},
	"member_soa_serials": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthMemberSoaSerialsResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: "The list of per-member SOA serial information.",
	},
	"ms_ad_integrated": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that determines whether Active Directory is integrated or not. This field is valid only when ms_managed is \"STUB\", \"AUTH_PRIMARY\", or \"AUTH_BOTH\".",
	},
	"ms_allow_transfer": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthMsAllowTransferResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The list of DNS clients that are allowed to perform zone transfers from a Microsoft DNS server. This setting applies only to zones with Microsoft DNS servers that are either primary or secondary servers. This setting does not inherit any value from the Grid or from any member that defines an allow_transfer value. This setting does not apply to any grid member. Use the allow_transfer field to control which DNS clients are allowed to perform zone transfers on Grid members.",
	},
	"ms_allow_transfer_mode": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("ADDRESS_AC", "ANY", "ANY_NS", "NONE"),
		},
		MarkdownDescription: "Determines which DNS clients are allowed to perform zone transfers from a Microsoft DNS server. Valid values are: \"ADDRESS_AC\", to use ms_allow_transfer field for specifying IP addresses, networks and Transaction Signature (TSIG) keys for clients that are allowed to do zone transfers. \"ANY\", to allow any client. \"ANY_NS\", to allow only the nameservers listed in this zone. \"NONE\", to deny all zone transfer requests.",
	},
	"ms_dc_ns_record_creation": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthMsDcNsRecordCreationResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The list of domain controllers that are allowed to create NS records for authoritative zones.",
	},
	"ms_ddns_mode": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("NONE"),
		MarkdownDescription: "Determines whether an Active Directory-integrated zone with a Microsoft DNS server as primary allows dynamic updates. Valid values are: \"SECURE\" if the zone allows secure updates only. \"NONE\" if the zone forbids dynamic updates. \"ANY\" if the zone accepts both secure and nonsecure updates. This field is valid only if ms_managed is either \"AUTH_PRIMARY\" or \"AUTH_BOTH\". If the flag ms_ad_integrated is false, the value \"SECURE\" is not allowed.",
		Validators: []validator.String{
			stringvalidator.OneOf("SECURE", "NONE", "ANY"),
		},
	},
	"ms_managed": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The flag that indicates whether the zone is assigned to a Microsoft DNS server. This flag returns the authoritative name server type of the Microsoft DNS server. Valid values are: \"NONE\" if the zone is not assigned to any Microsoft DNS server. \"STUB\" if the zone is assigned to a Microsoft DNS server as a stub zone. \"AUTH_PRIMARY\" if only the primary server of the zone is a Microsoft DNS server. \"AUTH_SECONDARY\" if only the secondary server of the zone is a Microsoft DNS server. \"AUTH_BOTH\" if both the primary and secondary servers of the zone are Microsoft DNS servers.",
	},
	"ms_primaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthMsPrimariesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.ConflictsWith(path.MatchRoot("ns_group")),
		},
		MarkdownDescription: "The list with the Microsoft DNS servers that are primary servers for the zone. Although a zone typically has just one primary name server, you can specify up to ten independent servers for a single zone.",
	},
	"ms_read_only": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if a Grid member manages the zone served by a Microsoft DNS server in read-only mode. This flag is true when a Grid member manages the zone in read-only mode, false otherwise. When the zone has the ms_read_only flag set to True, no changes can be made to this zone.",
	},
	"ms_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthMsSecondariesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.ConflictsWith(path.MatchRoot("ns_group")),
		},
		MarkdownDescription: "The list with the Microsoft DNS servers that are secondary servers for the zone.",
	},
	"ms_sync_disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag controls whether this zone is synchronized with Microsoft DNS servers.",
	},
	"ms_sync_master_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of MS synchronization master for this zone.",
	},
	"network_associations": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The list with the associated network/network container information.",
	},
	"network_view": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the network view in which this zone resides.",
	},
	"notify_delay": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(5),
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_notify_delay")),
			int64validator.Between(5, 86400),
		},
		MarkdownDescription: "The number of seconds in delay with which notify messages are sent to secondaries.",
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
		Optional: true,
		Computed: true,
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
	"records_monitored": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if this zone is also monitoring resource records.",
	},
	"remove_subzones": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Remove subzones delete option. Determines whether all child objects should be removed alongside with the parent zone or child objects should be assigned to another parental zone. By default child objects are deleted with the parent zone.",
	},
	"restart_if_needed": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Restarts the member service.",
	},
	"rr_not_queried_enabled_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time data collection for Not Queried Resource Record was enabled for this zone.",
	},
	"scavenging_settings": schema.SingleNestedAttribute{
		Attributes: ZoneAuthScavengingSettingsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_scavenging_settings")),
		},
	},
	"set_soa_serial_number": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The serial number in the SOA record incrementally changes every time the record is modified. The Infoblox appliance allows you to change the serial number (in the SOA record) for the primary server so it is higher than the secondary server, thereby ensuring zone transfers come from the primary server (as they should). To change the serial number you need to set a new value at \"soa_serial_number\" and pass \"set_soa_serial_number\" as True.",
	},
	"soa_default_ttl": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(
				path.MatchRoot("use_grid_zone_timer"),
				path.MatchRoot("soa_expire"),
				path.MatchRoot("soa_refresh"),
				path.MatchRoot("soa_retry"),
				path.MatchRoot("soa_negative_ttl"),
				path.MatchRoot("grid_primary"),
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
				path.MatchRoot("soa_refresh"),
				path.MatchRoot("soa_retry"),
				path.MatchRoot("soa_negative_ttl"),
				path.MatchRoot("grid_primary"),
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
				path.MatchRoot("soa_refresh"),
				path.MatchRoot("soa_retry"),
				path.MatchRoot("grid_primary"),
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
				path.MatchRoot("soa_refresh"),
				path.MatchRoot("soa_negative_ttl"),
				path.MatchRoot("grid_primary"),
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
	"srgs": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The associated shared record groups of a DNS zone. If a shared record group is associated with a zone, then all shared records in a shared record group will be shared in the zone.",
	},
	"update_forwarding": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZoneAuthUpdateForwardingResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("allow_update_forwarding")),
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Use this field to allow or deny dynamic DNS updates that are forwarded from specific IPv4/IPv6 addresses, networks, or a named ACL. You can also provide TSIG keys for clients that are allowed or denied to perform zone updates. This setting overrides the member-level setting.",
	},
	"use_allow_active_dir": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: allow_active_dir",
	},
	"use_allow_query": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: allow_query",
	},
	"use_allow_transfer": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: allow_transfer",
	},
	"use_allow_update": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: allow_update",
	},
	"use_allow_update_forwarding": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: allow_update_forwarding",
	},
	"use_check_names_policy": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Apply policy to dynamic updates and inbound zone transfers (This value applies only if the host name restriction policy is set to \"Strict Hostname Checking\".)",
	},
	"use_copy_xfer_to_notify": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: copy_xfer_to_notify",
	},
	"use_ddns_force_creation_timestamp_update": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ddns_force_creation_timestamp_update",
	},
	"use_ddns_patterns_restriction": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ddns_restrict_patterns_list , ddns_restrict_patterns",
	},
	"use_ddns_principal_security": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ddns_restrict_secure , ddns_principal_tracking, ddns_principal_group",
	},
	"use_ddns_restrict_protected": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ddns_restrict_protected",
	},
	"use_ddns_restrict_static": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: ddns_restrict_static",
	},
	"use_dnssec_key_params": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: dnssec_key_params",
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
	"use_import_from": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Use flag for: import_from",
	},
	"use_notify_delay": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: notify_delay",
	},
	"use_record_name_policy": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: record_name_policy",
	},
	"use_scavenging_settings": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: scavenging_settings , last_queried_acl",
	},
	"use_soa_email": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Use flag for: soa_email",
	},
	"using_srg_associations": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "This is true if the zone is associated with a shared record group.",
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
	"zone_format": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("FORWARD"),
		Validators: []validator.String{
			stringvalidator.OneOf("FORWARD", "IPV4", "IPV6"),
		},
		MarkdownDescription: "Determines the format of this zone.",
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
	},
	"zone_not_queried_enabled_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time when \"DNS Zones Last Queried\" was turned on for this zone.",
	},
}

func (m *ZoneAuthModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dns.ZoneAuth {
	if m == nil {
		return nil
	}
	to := &dns.ZoneAuth{
		AllowActiveDir:                      flex.ExpandFrameworkListNestedBlock(ctx, m.AllowActiveDir, diags, ExpandZoneAuthAllowActiveDir),
		AllowFixedRrsetOrder:                flex.ExpandBoolPointer(m.AllowFixedRrsetOrder),
		AllowGssTsigForUnderscoreZone:       flex.ExpandBoolPointer(m.AllowGssTsigForUnderscoreZone),
		AllowGssTsigZoneUpdates:             flex.ExpandBoolPointer(m.AllowGssTsigZoneUpdates),
		AllowQuery:                          flex.ExpandFrameworkListNestedBlock(ctx, m.AllowQuery, diags, ExpandZoneAuthAllowQuery),
		AllowTransfer:                       flex.ExpandFrameworkListNestedBlock(ctx, m.AllowTransfer, diags, ExpandZoneAuthAllowTransfer),
		AllowUpdate:                         flex.ExpandFrameworkListNestedBlock(ctx, m.AllowUpdate, diags, ExpandZoneAuthAllowUpdate),
		AllowUpdateForwarding:               flex.ExpandBoolPointer(m.AllowUpdateForwarding),
		Comment:                             flex.ExpandStringPointer(m.Comment),
		CopyXferToNotify:                    flex.ExpandBoolPointer(m.CopyXferToNotify),
		CreatePtrForBulkHosts:               flex.ExpandBoolPointer(m.CreatePtrForBulkHosts),
		CreatePtrForHosts:                   flex.ExpandBoolPointer(m.CreatePtrForHosts),
		CreateUnderscoreZones:               flex.ExpandBoolPointer(m.CreateUnderscoreZones),
		DdnsForceCreationTimestampUpdate:    flex.ExpandBoolPointer(m.DdnsForceCreationTimestampUpdate),
		DdnsPrincipalGroup:                  flex.ExpandStringPointer(m.DdnsPrincipalGroup),
		DdnsPrincipalTracking:               flex.ExpandBoolPointer(m.DdnsPrincipalTracking),
		DdnsRestrictPatterns:                flex.ExpandBoolPointer(m.DdnsRestrictPatterns),
		DdnsRestrictPatternsList:            flex.ExpandFrameworkListString(ctx, m.DdnsRestrictPatternsList, diags),
		DdnsRestrictProtected:               flex.ExpandBoolPointer(m.DdnsRestrictProtected),
		DdnsRestrictSecure:                  flex.ExpandBoolPointer(m.DdnsRestrictSecure),
		DdnsRestrictStatic:                  flex.ExpandBoolPointer(m.DdnsRestrictStatic),
		Disable:                             flex.ExpandBoolPointer(m.Disable),
		DisableForwarding:                   flex.ExpandBoolPointer(m.DisableForwarding),
		DnsIntegrityEnable:                  flex.ExpandBoolPointer(m.DnsIntegrityEnable),
		DnsIntegrityFrequency:               flex.ExpandInt64Pointer(m.DnsIntegrityFrequency),
		DnsIntegrityMember:                  flex.ExpandStringPointer(m.DnsIntegrityMember),
		DnsIntegrityVerboseLogging:          flex.ExpandBoolPointer(m.DnsIntegrityVerboseLogging),
		DnssecKeyParams:                     ExpandZoneAuthDnssecKeyParams(ctx, m.DnssecKeyParams, diags),
		DnssecKeys:                          flex.ExpandFrameworkListNestedBlock(ctx, m.DnssecKeys, diags, ExpandZoneAuthDnssecKeys),
		DoHostAbstraction:                   flex.ExpandBoolPointer(m.DoHostAbstraction),
		EffectiveCheckNamesPolicy:           flex.ExpandStringPointer(m.EffectiveCheckNamesPolicy),
		ExtAttrs:                            ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		ExternalPrimaries:                   flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalPrimaries, diags, ExpandZoneAuthExternalPrimaries),
		ExternalSecondaries:                 flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalSecondaries, diags, ExpandZoneAuthExternalSecondaries),
		GridPrimary:                         flex.ExpandFrameworkListNestedBlock(ctx, m.GridPrimary, diags, ExpandZoneAuthGridPrimary),
		GridSecondaries:                     flex.ExpandFrameworkListNestedBlock(ctx, m.GridSecondaries, diags, ExpandZoneAuthGridSecondaries),
		LastQueriedAcl:                      flex.ExpandFrameworkListNestedBlock(ctx, m.LastQueriedAcl, diags, ExpandZoneAuthLastQueriedAcl),
		Locked:                              flex.ExpandBoolPointer(m.Locked),
		MemberSoaMnames:                     flex.ExpandFrameworkListNestedBlock(ctx, m.MemberSoaMnames, diags, ExpandZoneAuthMemberSoaMnames),
		MsAdIntegrated:                      flex.ExpandBoolPointer(m.MsAdIntegrated),
		MsAllowTransfer:                     flex.ExpandFrameworkListNestedBlock(ctx, m.MsAllowTransfer, diags, ExpandZoneAuthMsAllowTransfer),
		MsAllowTransferMode:                 flex.ExpandStringPointer(m.MsAllowTransferMode),
		MsDcNsRecordCreation:                flex.ExpandFrameworkListNestedBlock(ctx, m.MsDcNsRecordCreation, diags, ExpandZoneAuthMsDcNsRecordCreation),
		MsDdnsMode:                          flex.ExpandStringPointer(m.MsDdnsMode),
		MsPrimaries:                         flex.ExpandFrameworkListNestedBlock(ctx, m.MsPrimaries, diags, ExpandZoneAuthMsPrimaries),
		MsSecondaries:                       flex.ExpandFrameworkListNestedBlock(ctx, m.MsSecondaries, diags, ExpandZoneAuthMsSecondaries),
		MsSyncDisabled:                      flex.ExpandBoolPointer(m.MsSyncDisabled),
		NotifyDelay:                         flex.ExpandInt64Pointer(m.NotifyDelay),
		NsGroup:                             flex.ExpandStringPointer(m.NsGroup),
		Prefix:                              flex.ExpandStringPointer(m.Prefix),
		RecordNamePolicy:                    flex.ExpandStringPointer(m.RecordNamePolicy),
		RemoveSubzones:                      flex.ExpandBoolPointer(m.RemoveSubzones),
		RestartIfNeeded:                     flex.ExpandBoolPointer(m.RestartIfNeeded),
		ScavengingSettings:                  ExpandZoneAuthScavengingSettings(ctx, m.ScavengingSettings, diags),
		SetSoaSerialNumber:                  flex.ExpandBoolPointer(m.SetSoaSerialNumber),
		SoaDefaultTtl:                       flex.ExpandInt64Pointer(m.SoaDefaultTtl),
		SoaEmail:                            flex.ExpandStringPointer(m.SoaEmail),
		SoaExpire:                           flex.ExpandInt64Pointer(m.SoaExpire),
		SoaNegativeTtl:                      flex.ExpandInt64Pointer(m.SoaNegativeTtl),
		SoaRefresh:                          flex.ExpandInt64Pointer(m.SoaRefresh),
		SoaRetry:                            flex.ExpandInt64Pointer(m.SoaRetry),
		SoaSerial:                           flex.ExpandInt64Pointer(m.SoaSerialNumber),
		Srgs:                                flex.ExpandFrameworkListString(ctx, m.Srgs, diags),
		UpdateForwarding:                    flex.ExpandFrameworkListNestedBlock(ctx, m.UpdateForwarding, diags, ExpandZoneAuthUpdateForwarding),
		UseAllowActiveDir:                   flex.ExpandBoolPointer(m.UseAllowActiveDir),
		UseAllowQuery:                       flex.ExpandBoolPointer(m.UseAllowQuery),
		UseAllowTransfer:                    flex.ExpandBoolPointer(m.UseAllowTransfer),
		UseAllowUpdate:                      flex.ExpandBoolPointer(m.UseAllowUpdate),
		UseAllowUpdateForwarding:            flex.ExpandBoolPointer(m.UseAllowUpdateForwarding),
		UseCheckNamesPolicy:                 flex.ExpandBoolPointer(m.UseCheckNamesPolicy),
		UseCopyXferToNotify:                 flex.ExpandBoolPointer(m.UseCopyXferToNotify),
		UseDdnsForceCreationTimestampUpdate: flex.ExpandBoolPointer(m.UseDdnsForceCreationTimestampUpdate),
		UseDdnsPatternsRestriction:          flex.ExpandBoolPointer(m.UseDdnsPatternsRestriction),
		UseDdnsPrincipalSecurity:            flex.ExpandBoolPointer(m.UseDdnsPrincipalSecurity),
		UseDdnsRestrictProtected:            flex.ExpandBoolPointer(m.UseDdnsRestrictProtected),
		UseDdnsRestrictStatic:               flex.ExpandBoolPointer(m.UseDdnsRestrictStatic),
		UseDnssecKeyParams:                  flex.ExpandBoolPointer(m.UseDnssecKeyParams),
		UseExternalPrimary:                  flex.ExpandBoolPointer(m.UseExternalPrimary),
		UseGridZoneTimer:                    flex.ExpandBoolPointer(m.UseGridZoneTimer),
		UseNotifyDelay:                      flex.ExpandBoolPointer(m.UseNotifyDelay),
		UseRecordNamePolicy:                 flex.ExpandBoolPointer(m.UseRecordNamePolicy),
		UseScavengingSettings:               flex.ExpandBoolPointer(m.UseScavengingSettings),
		UseSoaEmail:                         flex.ExpandBoolPointer(m.UseSoaEmail),
		View:                                flex.ExpandStringPointer(m.View),
	}
	if isCreate {
		to.Fqdn = flex.ExpandStringPointer(m.Fqdn)
		to.ZoneFormat = flex.ExpandStringPointer(m.ZoneFormat)
	}

	return to
}

func FlattenZoneAuth(ctx context.Context, from *dns.ZoneAuth, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneAuthAttrTypes)
	}
	m := ZoneAuthModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, ZoneAuthAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneAuthModel) Flatten(ctx context.Context, from *dns.ZoneAuth, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneAuthModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Address = flex.FlattenIPAddress(from.Address)
	m.AllowActiveDir = flex.FlattenFrameworkListNestedBlock(ctx, from.AllowActiveDir, ZoneAuthAllowActiveDirAttrTypes, diags, FlattenZoneAuthAllowActiveDir)
	m.AllowFixedRrsetOrder = types.BoolPointerValue(from.AllowFixedRrsetOrder)
	m.AllowGssTsigForUnderscoreZone = types.BoolPointerValue(from.AllowGssTsigForUnderscoreZone)
	m.AllowGssTsigZoneUpdates = types.BoolPointerValue(from.AllowGssTsigZoneUpdates)
	planAllowQuery := m.AllowQuery
	m.AllowQuery = flex.FlattenFrameworkListNestedBlock(ctx, from.AllowQuery, ZoneAuthAllowQueryAttrTypes, diags, FlattenZoneAuthAllowQuery)
	if !planAllowQuery.IsNull() {
		result, diags := utils.CopyFieldFromPlanToRespList(ctx, planAllowQuery, m.AllowQuery, "use_tsig_key_name")
		if !diags.HasError() {
			m.AllowQuery = result.(basetypes.ListValue)
		}
	}
	planAllowTransfer := m.AllowTransfer
	m.AllowTransfer = flex.FlattenFrameworkListNestedBlock(ctx, from.AllowTransfer, ZoneAuthAllowTransferAttrTypes, diags, FlattenZoneAuthAllowTransfer)
	if !planAllowTransfer.IsNull() {
		result, diags := utils.CopyFieldFromPlanToRespList(ctx, planAllowTransfer, m.AllowTransfer, "use_tsig_key_name")
		if !diags.HasError() {
			m.AllowTransfer = result.(basetypes.ListValue)
		}
	}
	planAllowUpdate := m.AllowUpdate
	m.AllowUpdate = flex.FlattenFrameworkListNestedBlock(ctx, from.AllowUpdate, ZoneAuthAllowUpdateAttrTypes, diags, FlattenZoneAuthAllowUpdate)
	if !planAllowUpdate.IsNull() {
		result, diags := utils.CopyFieldFromPlanToRespList(ctx, planAllowUpdate, m.AllowUpdate, "use_tsig_key_name")
		if !diags.HasError() {
			m.AllowUpdate = result.(basetypes.ListValue)
		}
	}
	m.AllowUpdateForwarding = types.BoolPointerValue(from.AllowUpdateForwarding)
	m.AwsRte53ZoneInfo = FlattenZoneAuthAwsRte53ZoneInfo(ctx, from.AwsRte53ZoneInfo, diags)
	m.CloudInfo = FlattenZoneAuthCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CopyXferToNotify = types.BoolPointerValue(from.CopyXferToNotify)
	m.CreateUnderscoreZones = types.BoolPointerValue(from.CreateUnderscoreZones)
	m.DdnsForceCreationTimestampUpdate = types.BoolPointerValue(from.DdnsForceCreationTimestampUpdate)
	m.DdnsPrincipalGroup = flex.FlattenStringPointer(from.DdnsPrincipalGroup)
	m.DdnsPrincipalTracking = types.BoolPointerValue(from.DdnsPrincipalTracking)
	m.DdnsRestrictPatterns = types.BoolPointerValue(from.DdnsRestrictPatterns)
	m.DdnsRestrictPatternsList = flex.FlattenFrameworkUnorderedList(ctx, types.StringType, from.DdnsRestrictPatternsList, diags)
	m.DdnsRestrictProtected = types.BoolPointerValue(from.DdnsRestrictProtected)
	m.DdnsRestrictSecure = types.BoolPointerValue(from.DdnsRestrictSecure)
	m.DdnsRestrictStatic = types.BoolPointerValue(from.DdnsRestrictStatic)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DisableForwarding = types.BoolPointerValue(from.DisableForwarding)
	m.DisplayDomain = flex.FlattenStringPointer(from.DisplayDomain)
	m.DnsFqdn = flex.FlattenStringPointer(from.DnsFqdn)
	m.DnsIntegrityEnable = types.BoolPointerValue(from.DnsIntegrityEnable)
	m.DnsIntegrityFrequency = flex.FlattenInt64Pointer(from.DnsIntegrityFrequency)
	m.DnsIntegrityMember = flex.FlattenStringPointer(from.DnsIntegrityMember)
	m.DnsIntegrityVerboseLogging = types.BoolPointerValue(from.DnsIntegrityVerboseLogging)
	m.DnsSoaEmail = flex.FlattenStringPointer(from.DnsSoaEmail)
	m.DnssecKeyParams = FlattenZoneAuthDnssecKeyParams(ctx, from.DnssecKeyParams, diags)
	m.DnssecKeys = flex.FlattenFrameworkListNestedBlock(ctx, from.DnssecKeys, ZoneAuthDnssecKeysAttrTypes, diags, FlattenZoneAuthDnssecKeys)
	m.DnssecKskRolloverDate = flex.FlattenInt64Pointer(from.DnssecKskRolloverDate)
	m.DnssecZskRolloverDate = flex.FlattenInt64Pointer(from.DnssecZskRolloverDate)
	m.EffectiveCheckNamesPolicy = flex.FlattenStringPointer(from.EffectiveCheckNamesPolicy)
	m.EffectiveRecordNamePolicy = flex.FlattenStringPointer(from.EffectiveRecordNamePolicy)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	planExternalPrimaries := m.ExternalPrimaries
	m.ExternalPrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalPrimaries, ZoneAuthExternalPrimariesAttrTypes, diags, FlattenZoneAuthExternalPrimaries)
	if !planExternalPrimaries.IsNull() {
		result, diags := utils.CopyFieldFromPlanToRespList(ctx, planExternalPrimaries, m.ExternalPrimaries, "tsig_key_name")
		if !diags.HasError() {
			m.ExternalPrimaries = result.(basetypes.ListValue)
		}
	}
	planExternalSecondaries := m.ExternalSecondaries
	m.ExternalSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalSecondaries, ZoneAuthExternalSecondariesAttrTypes, diags, FlattenZoneAuthExternalSecondaries)
	if !planExternalSecondaries.IsNull() {
		result, diags := utils.CopyFieldFromPlanToRespList(ctx, planExternalSecondaries, m.ExternalSecondaries, "tsig_key_name")
		if !diags.HasError() {
			m.ExternalSecondaries = result.(basetypes.ListValue)
		}
	}
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	planGridPrimary := m.GridPrimary
	m.GridPrimary = flex.FlattenFrameworkListNestedBlock(ctx, from.GridPrimary, ZoneAuthGridPrimaryAttrTypes, diags, FlattenZoneAuthGridPrimary)
	if !planGridPrimary.IsUnknown() {
		reOrderedList, diags := utils.ReorderAndFilterNestedListResponse(ctx, planGridPrimary, m.GridPrimary, "name")
		if !diags.HasError() {
			m.GridPrimary = reOrderedList.(basetypes.ListValue)
		}
	}
	m.GridPrimarySharedWithMsParentDelegation = types.BoolPointerValue(from.GridPrimarySharedWithMsParentDelegation)
	planGridSecondary := m.GridSecondaries
	m.GridSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.GridSecondaries, ZoneAuthGridSecondariesAttrTypes, diags, FlattenZoneAuthGridSecondaries)
	if !planGridSecondary.IsUnknown() {
		reOrderedList, diags := utils.ReorderAndFilterNestedListResponse(ctx, planGridSecondary, m.GridSecondaries, "name")
		if !diags.HasError() {
			m.GridSecondaries = reOrderedList.(basetypes.ListValue)
		}
	}
	m.ImportFrom = flex.FlattenIPAddress(from.ImportFrom)
	m.IsDnssecEnabled = types.BoolPointerValue(from.IsDnssecEnabled)
	m.IsDnssecSigned = types.BoolPointerValue(from.IsDnssecSigned)
	m.IsMultimaster = types.BoolPointerValue(from.IsMultimaster)
	m.LastQueried = flex.FlattenInt64Pointer(from.LastQueried)
	m.LastQueriedAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.LastQueriedAcl, ZoneAuthLastQueriedAclAttrTypes, diags, FlattenZoneAuthLastQueriedAcl)
	m.Locked = types.BoolPointerValue(from.Locked)
	m.LockedBy = flex.FlattenStringPointer(from.LockedBy)
	m.MaskPrefix = flex.FlattenStringPointer(from.MaskPrefix)
	m.MemberSoaMnames = flex.FlattenFrameworkListNestedBlock(ctx, from.MemberSoaMnames, ZoneAuthMemberSoaMnamesAttrTypes, diags, FlattenZoneAuthMemberSoaMnames)
	m.MemberSoaSerials = flex.FlattenFrameworkListNestedBlock(ctx, from.MemberSoaSerials, ZoneAuthMemberSoaSerialsAttrTypes, diags, FlattenZoneAuthMemberSoaSerials)
	m.MsAdIntegrated = types.BoolPointerValue(from.MsAdIntegrated)
	m.MsAllowTransfer = flex.FlattenFrameworkListNestedBlock(ctx, from.MsAllowTransfer, ZoneAuthMsAllowTransferAttrTypes, diags, FlattenZoneAuthMsAllowTransfer)
	m.MsAllowTransferMode = flex.FlattenStringPointer(from.MsAllowTransferMode)
	m.MsDcNsRecordCreation = flex.FlattenFrameworkListNestedBlock(ctx, from.MsDcNsRecordCreation, ZoneAuthMsDcNsRecordCreationAttrTypes, diags, FlattenZoneAuthMsDcNsRecordCreation)
	m.MsDdnsMode = flex.FlattenStringPointer(from.MsDdnsMode)
	m.MsManaged = flex.FlattenStringPointer(from.MsManaged)
	m.MsPrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.MsPrimaries, ZoneAuthMsPrimariesAttrTypes, diags, FlattenZoneAuthMsPrimaries)
	m.MsReadOnly = types.BoolPointerValue(from.MsReadOnly)
	m.MsSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.MsSecondaries, ZoneAuthMsSecondariesAttrTypes, diags, FlattenZoneAuthMsSecondaries)
	m.MsSyncDisabled = types.BoolPointerValue(from.MsSyncDisabled)
	m.MsSyncMasterName = flex.FlattenStringPointer(from.MsSyncMasterName)
	m.NetworkAssociations = flex.FlattenFrameworkListString(ctx, from.NetworkAssociations, diags)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.NotifyDelay = flex.FlattenInt64Pointer(from.NotifyDelay)
	m.NsGroup = flex.FlattenStringPointer(from.NsGroup)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.Prefix = flex.FlattenStringPointer(from.Prefix)
	m.PrimaryType = flex.FlattenStringPointer(from.PrimaryType)
	m.RecordNamePolicy = flex.FlattenStringPointer(from.RecordNamePolicy)
	m.RecordsMonitored = types.BoolPointerValue(from.RecordsMonitored)
	m.RemoveSubzones = types.BoolPointerValue(from.RemoveSubzones)
	m.RrNotQueriedEnabledTime = flex.FlattenInt64Pointer(from.RrNotQueriedEnabledTime)
	m.ScavengingSettings = FlattenZoneAuthScavengingSettings(ctx, from.ScavengingSettings, diags)
	m.SoaDefaultTtl = flex.FlattenInt64Pointer(from.SoaDefaultTtl)
	m.SoaEmail = flex.FlattenStringPointer(from.SoaEmail)
	m.SoaExpire = flex.FlattenInt64Pointer(from.SoaExpire)
	m.SoaNegativeTtl = flex.FlattenInt64Pointer(from.SoaNegativeTtl)
	m.SoaRefresh = flex.FlattenInt64Pointer(from.SoaRefresh)
	m.SoaRetry = flex.FlattenInt64Pointer(from.SoaRetry)
	m.SoaSerialNumber = flex.FlattenInt64Pointer(from.SoaSerial)
	m.Srgs = flex.FlattenFrameworkListString(ctx, from.Srgs, diags)
	m.UpdateForwarding = flex.FlattenFrameworkListNestedBlock(ctx, from.UpdateForwarding, ZoneAuthUpdateForwardingAttrTypes, diags, FlattenZoneAuthUpdateForwarding)
	m.UseAllowActiveDir = types.BoolPointerValue(from.UseAllowActiveDir)
	m.UseAllowQuery = types.BoolPointerValue(from.UseAllowQuery)
	m.UseAllowTransfer = types.BoolPointerValue(from.UseAllowTransfer)
	m.UseAllowUpdate = types.BoolPointerValue(from.UseAllowUpdate)
	m.UseAllowUpdateForwarding = types.BoolPointerValue(from.UseAllowUpdateForwarding)
	m.UseCheckNamesPolicy = types.BoolPointerValue(from.UseCheckNamesPolicy)
	m.UseCopyXferToNotify = types.BoolPointerValue(from.UseCopyXferToNotify)
	m.UseDdnsForceCreationTimestampUpdate = types.BoolPointerValue(from.UseDdnsForceCreationTimestampUpdate)
	m.UseDdnsPatternsRestriction = types.BoolPointerValue(from.UseDdnsPatternsRestriction)
	m.UseDdnsPrincipalSecurity = types.BoolPointerValue(from.UseDdnsPrincipalSecurity)
	m.UseDdnsRestrictProtected = types.BoolPointerValue(from.UseDdnsRestrictProtected)
	m.UseDdnsRestrictStatic = types.BoolPointerValue(from.UseDdnsRestrictStatic)
	m.UseDnssecKeyParams = types.BoolPointerValue(from.UseDnssecKeyParams)
	m.UseExternalPrimary = types.BoolPointerValue(from.UseExternalPrimary)
	m.UseGridZoneTimer = types.BoolPointerValue(from.UseGridZoneTimer)
	m.UseImportFrom = types.BoolPointerValue(from.UseImportFrom)
	m.UseNotifyDelay = types.BoolPointerValue(from.UseNotifyDelay)
	m.UseRecordNamePolicy = types.BoolPointerValue(from.UseRecordNamePolicy)
	m.UseScavengingSettings = types.BoolPointerValue(from.UseScavengingSettings)
	m.UseSoaEmail = types.BoolPointerValue(from.UseSoaEmail)
	m.UsingSrgAssociations = types.BoolPointerValue(from.UsingSrgAssociations)
	m.View = flex.FlattenStringPointer(from.View)
	m.ZoneFormat = flex.FlattenStringPointer(from.ZoneFormat)
	m.ZoneNotQueriedEnabledTime = flex.FlattenInt64Pointer(from.ZoneNotQueriedEnabledTime)
}
