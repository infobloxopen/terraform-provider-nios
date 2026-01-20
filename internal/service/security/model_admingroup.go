package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
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
	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type AdmingroupModel struct {
	Ref                               types.String                     `tfsdk:"ref"`
	Uuid                              types.String                     `tfsdk:"uuid"`
	AccessMethod                      internaltypes.UnorderedListValue `tfsdk:"access_method"`
	AdminSetCommands                  types.Object                     `tfsdk:"admin_set_commands"`
	AdminShowCommands                 types.Object                     `tfsdk:"admin_show_commands"`
	AdminToplevelCommands             types.Object                     `tfsdk:"admin_toplevel_commands"`
	CloudSetCommands                  types.Object                     `tfsdk:"cloud_set_commands"`
	CloudShowCommands                 types.Object                     `tfsdk:"cloud_show_commands"`
	Comment                           types.String                     `tfsdk:"comment"`
	DatabaseSetCommands               types.Object                     `tfsdk:"database_set_commands"`
	DatabaseShowCommands              types.Object                     `tfsdk:"database_show_commands"`
	DhcpSetCommands                   types.Object                     `tfsdk:"dhcp_set_commands"`
	DhcpShowCommands                  types.Object                     `tfsdk:"dhcp_show_commands"`
	Disable                           types.Bool                       `tfsdk:"disable"`
	DisableConcurrentLogin            types.Bool                       `tfsdk:"disable_concurrent_login"`
	DnsSetCommands                    types.Object                     `tfsdk:"dns_set_commands"`
	DnsShowCommands                   types.Object                     `tfsdk:"dns_show_commands"`
	DnsToplevelCommands               types.Object                     `tfsdk:"dns_toplevel_commands"`
	DockerSetCommands                 types.Object                     `tfsdk:"docker_set_commands"`
	DockerShowCommands                types.Object                     `tfsdk:"docker_show_commands"`
	EmailAddresses                    internaltypes.UnorderedListValue `tfsdk:"email_addresses"`
	EnableRestrictedUserAccess        types.Bool                       `tfsdk:"enable_restricted_user_access"`
	ExtAttrs                          types.Map                        `tfsdk:"extattrs"`
	ExtAttrsAll                       types.Map                        `tfsdk:"extattrs_all"`
	GridSetCommands                   types.Object                     `tfsdk:"grid_set_commands"`
	GridShowCommands                  types.Object                     `tfsdk:"grid_show_commands"`
	InactivityLockoutSetting          types.Object                     `tfsdk:"inactivity_lockout_setting"`
	LicensingSetCommands              types.Object                     `tfsdk:"licensing_set_commands"`
	LicensingShowCommands             types.Object                     `tfsdk:"licensing_show_commands"`
	LockoutSetting                    types.Object                     `tfsdk:"lockout_setting"`
	MachineControlToplevelCommands    types.Object                     `tfsdk:"machine_control_toplevel_commands"`
	Name                              types.String                     `tfsdk:"name"`
	NetworkingSetCommands             types.Object                     `tfsdk:"networking_set_commands"`
	NetworkingShowCommands            types.Object                     `tfsdk:"networking_show_commands"`
	PasswordSetting                   types.Object                     `tfsdk:"password_setting"`
	Roles                             types.List                       `tfsdk:"roles"`
	SamlSetting                       types.Object                     `tfsdk:"saml_setting"`
	SecuritySetCommands               types.Object                     `tfsdk:"security_set_commands"`
	SecurityShowCommands              types.Object                     `tfsdk:"security_show_commands"`
	Superuser                         types.Bool                       `tfsdk:"superuser"`
	TroubleShootingToplevelCommands   types.Object                     `tfsdk:"trouble_shooting_toplevel_commands"`
	UseAccountInactivityLockoutEnable types.Bool                       `tfsdk:"use_account_inactivity_lockout_enable"`
	UseDisableConcurrentLogin         types.Bool                       `tfsdk:"use_disable_concurrent_login"`
	UseLockoutSetting                 types.Bool                       `tfsdk:"use_lockout_setting"`
	UsePasswordSetting                types.Bool                       `tfsdk:"use_password_setting"`
	UserAccess                        types.List                       `tfsdk:"user_access"`
}

var AdmingroupAttrTypes = map[string]attr.Type{
	"ref":                                   types.StringType,
	"uuid":                                  types.StringType,
	"access_method":                         internaltypes.UnorderedListOfStringType,
	"admin_set_commands":                    types.ObjectType{AttrTypes: AdmingroupAdminSetCommandsAttrTypes},
	"admin_show_commands":                   types.ObjectType{AttrTypes: AdmingroupAdminShowCommandsAttrTypes},
	"admin_toplevel_commands":               types.ObjectType{AttrTypes: AdmingroupAdminToplevelCommandsAttrTypes},
	"cloud_set_commands":                    types.ObjectType{AttrTypes: AdmingroupCloudSetCommandsAttrTypes},
	"cloud_show_commands":                   types.ObjectType{AttrTypes: AdmingroupCloudShowCommandsAttrTypes},
	"comment":                               types.StringType,
	"database_set_commands":                 types.ObjectType{AttrTypes: AdmingroupDatabaseSetCommandsAttrTypes},
	"database_show_commands":                types.ObjectType{AttrTypes: AdmingroupDatabaseShowCommandsAttrTypes},
	"dhcp_set_commands":                     types.ObjectType{AttrTypes: AdmingroupDhcpSetCommandsAttrTypes},
	"dhcp_show_commands":                    types.ObjectType{AttrTypes: AdmingroupDhcpShowCommandsAttrTypes},
	"disable":                               types.BoolType,
	"disable_concurrent_login":              types.BoolType,
	"dns_set_commands":                      types.ObjectType{AttrTypes: AdmingroupDnsSetCommandsAttrTypes},
	"dns_show_commands":                     types.ObjectType{AttrTypes: AdmingroupDnsShowCommandsAttrTypes},
	"dns_toplevel_commands":                 types.ObjectType{AttrTypes: AdmingroupDnsToplevelCommandsAttrTypes},
	"docker_set_commands":                   types.ObjectType{AttrTypes: AdmingroupDockerSetCommandsAttrTypes},
	"docker_show_commands":                  types.ObjectType{AttrTypes: AdmingroupDockerShowCommandsAttrTypes},
	"email_addresses":                       internaltypes.UnorderedListOfStringType,
	"enable_restricted_user_access":         types.BoolType,
	"extattrs":                              types.MapType{ElemType: types.StringType},
	"extattrs_all":                          types.MapType{ElemType: types.StringType},
	"grid_set_commands":                     types.ObjectType{AttrTypes: AdmingroupGridSetCommandsAttrTypes},
	"grid_show_commands":                    types.ObjectType{AttrTypes: AdmingroupGridShowCommandsAttrTypes},
	"inactivity_lockout_setting":            types.ObjectType{AttrTypes: AdmingroupInactivityLockoutSettingAttrTypes},
	"licensing_set_commands":                types.ObjectType{AttrTypes: AdmingroupLicensingSetCommandsAttrTypes},
	"licensing_show_commands":               types.ObjectType{AttrTypes: AdmingroupLicensingShowCommandsAttrTypes},
	"lockout_setting":                       types.ObjectType{AttrTypes: AdmingroupLockoutSettingAttrTypes},
	"machine_control_toplevel_commands":     types.ObjectType{AttrTypes: AdmingroupMachineControlToplevelCommandsAttrTypes},
	"name":                                  types.StringType,
	"networking_set_commands":               types.ObjectType{AttrTypes: AdmingroupNetworkingSetCommandsAttrTypes},
	"networking_show_commands":              types.ObjectType{AttrTypes: AdmingroupNetworkingShowCommandsAttrTypes},
	"password_setting":                      types.ObjectType{AttrTypes: AdmingroupPasswordSettingAttrTypes},
	"roles":                                 types.ListType{ElemType: types.StringType},
	"saml_setting":                          types.ObjectType{AttrTypes: AdmingroupSamlSettingAttrTypes},
	"security_set_commands":                 types.ObjectType{AttrTypes: AdmingroupSecuritySetCommandsAttrTypes},
	"security_show_commands":                types.ObjectType{AttrTypes: AdmingroupSecurityShowCommandsAttrTypes},
	"superuser":                             types.BoolType,
	"trouble_shooting_toplevel_commands":    types.ObjectType{AttrTypes: AdmingroupTroubleShootingToplevelCommandsAttrTypes},
	"use_account_inactivity_lockout_enable": types.BoolType,
	"use_disable_concurrent_login":          types.BoolType,
	"use_lockout_setting":                   types.BoolType,
	"use_password_setting":                  types.BoolType,
	"user_access":                           types.ListType{ElemType: types.ObjectType{AttrTypes: AdmingroupUserAccessAttrTypes}},
}

var AdmingroupResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"access_method": schema.ListAttribute{
		CustomType:  internaltypes.UnorderedListOfStringType,
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Default: listdefault.StaticValue(
			types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("GUI"),
				types.StringValue("API"),
				types.StringValue("TAXII"),
				types.StringValue("CLI"),
			}),
		),
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			customvalidator.StringsInSlice([]string{"API", "CLI", "CLOUD_API", "GUI", "TAXII"}),
		},
		MarkdownDescription: "Access methods specify whether an admin group can use the GUI and the API to access the appliance or to send Taxii messages to the appliance. Note that API includes both the Perl API and RESTful API.",
	},
	"admin_set_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupAdminSetCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Admin set commands for the admin command group.",
	},
	"admin_show_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupAdminShowCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Admin show commands for the admin command group.",
	},
	"admin_toplevel_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupAdminToplevelCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Admin toplevel commands for the admin command group",
	},
	"cloud_set_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupCloudSetCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Cloud set commands for the cloud command group.",
	},
	"cloud_show_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupCloudShowCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Cloud show commands for admin group.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "Comment for the Admin Group; maximum 256 characters.",
	},
	"database_set_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupDatabaseSetCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Database show commands for admin group.",
	},
	"database_show_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupDatabaseShowCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Database show commands for the database command.",
	},
	"dhcp_set_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupDhcpSetCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Dhcp set commands for the dhcp command group.",
	},
	"dhcp_show_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupDhcpShowCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Dhcp show commands for the dhcp command group.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the Admin Group is disabled or not. When this is set to False, the Admin Group is enabled.",
	},
	"disable_concurrent_login": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_disable_concurrent_login")),
		},
		MarkdownDescription: "Disable concurrent login feature",
	},
	"dns_set_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupDnsSetCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Dns set commands for the dns command group.",
	},
	"dns_show_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupDnsShowCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Dns show commands for the dns command group.",
	},
	"dns_toplevel_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupDnsToplevelCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Dns toplevel commands for the dns command group.",
	},
	"docker_set_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupDockerSetCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Docker set commands for the docker command group.",
	},
	"docker_show_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupDockerShowCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Docker show commands for the docker command group.",
	},
	"email_addresses": schema.ListAttribute{
		CustomType:  internaltypes.UnorderedListOfStringType,
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The e-mail addresses for the Admin Group.",
	},
	"enable_restricted_user_access": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines whether the restrictions will be applied to the admin connector level for users of this Admin Group.",
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
	"grid_set_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupGridSetCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Grid set commands for the grid command group.",
	},
	"grid_show_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupGridShowCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Grid show commands for the grid command group.",
	},
	"inactivity_lockout_setting": schema.SingleNestedAttribute{
		Attributes: AdmingroupInactivityLockoutSettingResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_account_inactivity_lockout_enable")),
		},
		MarkdownDescription: "The Admin group inactivity lockout settings.",
	},
	"licensing_set_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupLicensingSetCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Set commands for the licensing command group.",
	},
	"licensing_show_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupLicensingShowCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Show commands for the licensing command group.",
	},
	"lockout_setting": schema.SingleNestedAttribute{
		Attributes: AdmingroupLockoutSettingResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_lockout_setting")),
		},
		MarkdownDescription: "This struct specifies security policy settings in admin group.",
	},
	"machine_control_toplevel_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupMachineControlToplevelCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Machine control toplevel commands for the machine control command group.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the Admin Group.",
	},
	"networking_set_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupNetworkingSetCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Set commands for the networking command group.",
	},
	"networking_show_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupNetworkingShowCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Show commands for the networking command group.",
	},
	"password_setting": schema.SingleNestedAttribute{
		Attributes: AdmingroupPasswordSettingResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_password_setting")),
		},
		MarkdownDescription: "The Admin Group password settings.",
	},
	"roles": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The names of roles this Admin Group applies to.",
	},
	"saml_setting": schema.SingleNestedAttribute{
		Attributes:          AdmingroupSamlSettingResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Admin Group SAML settings.",
	},
	"security_set_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupSecuritySetCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Set commands for the security command group.",
	},
	"security_show_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupSecurityShowCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Show commands for the security command group.",
	},
	"superuser": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether this Admin Group is a superuser group. A superuser group can perform all operations on the appliance, and can view and configure all types of data.",
	},
	"trouble_shooting_toplevel_commands": schema.SingleNestedAttribute{
		Attributes:          AdmingroupTroubleShootingToplevelCommandsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Toplevel commands for the troubleshooting command group.",
	},
	"use_account_inactivity_lockout_enable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This is the use flag for account inactivity lockout settings.",
	},
	"use_disable_concurrent_login": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Whether to override grid concurrent login",
	},
	"use_lockout_setting": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Whether to override grid sequential lockout setting",
	},
	"use_password_setting": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Whether grid password expiry setting should be override.",
	},
	"user_access": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AdmingroupUserAccessResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The access control items for this Admin Group.",
	},
}

func (m *AdmingroupModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.Admingroup {
	if m == nil {
		return nil
	}
	to := &security.Admingroup{
		AccessMethod:                      flex.ExpandFrameworkListString(ctx, m.AccessMethod, diags),
		AdminSetCommands:                  ExpandAdmingroupAdminSetCommands(ctx, m.AdminSetCommands, diags),
		AdminShowCommands:                 ExpandAdmingroupAdminShowCommands(ctx, m.AdminShowCommands, diags),
		AdminToplevelCommands:             ExpandAdmingroupAdminToplevelCommands(ctx, m.AdminToplevelCommands, diags),
		CloudSetCommands:                  ExpandAdmingroupCloudSetCommands(ctx, m.CloudSetCommands, diags),
		CloudShowCommands:                 ExpandAdmingroupCloudShowCommands(ctx, m.CloudShowCommands, diags),
		Comment:                           flex.ExpandStringPointer(m.Comment),
		DatabaseSetCommands:               ExpandAdmingroupDatabaseSetCommands(ctx, m.DatabaseSetCommands, diags),
		DatabaseShowCommands:              ExpandAdmingroupDatabaseShowCommands(ctx, m.DatabaseShowCommands, diags),
		DhcpSetCommands:                   ExpandAdmingroupDhcpSetCommands(ctx, m.DhcpSetCommands, diags),
		DhcpShowCommands:                  ExpandAdmingroupDhcpShowCommands(ctx, m.DhcpShowCommands, diags),
		Disable:                           flex.ExpandBoolPointer(m.Disable),
		DisableConcurrentLogin:            flex.ExpandBoolPointer(m.DisableConcurrentLogin),
		DnsSetCommands:                    ExpandAdmingroupDnsSetCommands(ctx, m.DnsSetCommands, diags),
		DnsShowCommands:                   ExpandAdmingroupDnsShowCommands(ctx, m.DnsShowCommands, diags),
		DnsToplevelCommands:               ExpandAdmingroupDnsToplevelCommands(ctx, m.DnsToplevelCommands, diags),
		DockerSetCommands:                 ExpandAdmingroupDockerSetCommands(ctx, m.DockerSetCommands, diags),
		DockerShowCommands:                ExpandAdmingroupDockerShowCommands(ctx, m.DockerShowCommands, diags),
		EmailAddresses:                    flex.ExpandFrameworkListString(ctx, m.EmailAddresses, diags),
		ExtAttrs:                          ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		GridSetCommands:                   ExpandAdmingroupGridSetCommands(ctx, m.GridSetCommands, diags),
		GridShowCommands:                  ExpandAdmingroupGridShowCommands(ctx, m.GridShowCommands, diags),
		InactivityLockoutSetting:          ExpandAdmingroupInactivityLockoutSetting(ctx, m.InactivityLockoutSetting, diags),
		LicensingSetCommands:              ExpandAdmingroupLicensingSetCommands(ctx, m.LicensingSetCommands, diags),
		LicensingShowCommands:             ExpandAdmingroupLicensingShowCommands(ctx, m.LicensingShowCommands, diags),
		LockoutSetting:                    ExpandAdmingroupLockoutSetting(ctx, m.LockoutSetting, diags),
		MachineControlToplevelCommands:    ExpandAdmingroupMachineControlToplevelCommands(ctx, m.MachineControlToplevelCommands, diags),
		Name:                              flex.ExpandStringPointer(m.Name),
		NetworkingSetCommands:             ExpandAdmingroupNetworkingSetCommands(ctx, m.NetworkingSetCommands, diags),
		NetworkingShowCommands:            ExpandAdmingroupNetworkingShowCommands(ctx, m.NetworkingShowCommands, diags),
		PasswordSetting:                   ExpandAdmingroupPasswordSetting(ctx, m.PasswordSetting, diags),
		Roles:                             flex.ExpandFrameworkListString(ctx, m.Roles, diags),
		SamlSetting:                       ExpandAdmingroupSamlSetting(ctx, m.SamlSetting, diags),
		SecuritySetCommands:               ExpandAdmingroupSecuritySetCommands(ctx, m.SecuritySetCommands, diags),
		SecurityShowCommands:              ExpandAdmingroupSecurityShowCommands(ctx, m.SecurityShowCommands, diags),
		Superuser:                         flex.ExpandBoolPointer(m.Superuser),
		TroubleShootingToplevelCommands:   ExpandAdmingroupTroubleShootingToplevelCommands(ctx, m.TroubleShootingToplevelCommands, diags),
		UseAccountInactivityLockoutEnable: flex.ExpandBoolPointer(m.UseAccountInactivityLockoutEnable),
		UseDisableConcurrentLogin:         flex.ExpandBoolPointer(m.UseDisableConcurrentLogin),
		UseLockoutSetting:                 flex.ExpandBoolPointer(m.UseLockoutSetting),
		UsePasswordSetting:                flex.ExpandBoolPointer(m.UsePasswordSetting),
		UserAccess:                        flex.ExpandFrameworkListNestedBlock(ctx, m.UserAccess, diags, ExpandAdmingroupUserAccess),
	}
	return to
}

func FlattenAdmingroup(ctx context.Context, from *security.Admingroup, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupAttrTypes)
	}
	m := AdmingroupModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, AdmingroupAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupModel) Flatten(ctx context.Context, from *security.Admingroup, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.AccessMethod = flex.FlattenFrameworkUnorderedList(ctx, types.StringType, from.AccessMethod, diags)
	m.AdminSetCommands = FlattenAdmingroupAdminSetCommands(ctx, from.AdminSetCommands, diags)
	m.AdminShowCommands = FlattenAdmingroupAdminShowCommands(ctx, from.AdminShowCommands, diags)
	m.AdminToplevelCommands = FlattenAdmingroupAdminToplevelCommands(ctx, from.AdminToplevelCommands, diags)
	m.CloudSetCommands = FlattenAdmingroupCloudSetCommands(ctx, from.CloudSetCommands, diags)
	m.CloudShowCommands = FlattenAdmingroupCloudShowCommands(ctx, from.CloudShowCommands, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DatabaseSetCommands = FlattenAdmingroupDatabaseSetCommands(ctx, from.DatabaseSetCommands, diags)
	m.DatabaseShowCommands = FlattenAdmingroupDatabaseShowCommands(ctx, from.DatabaseShowCommands, diags)
	m.DhcpSetCommands = FlattenAdmingroupDhcpSetCommands(ctx, from.DhcpSetCommands, diags)
	m.DhcpShowCommands = FlattenAdmingroupDhcpShowCommands(ctx, from.DhcpShowCommands, diags)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.DisableConcurrentLogin = types.BoolPointerValue(from.DisableConcurrentLogin)
	m.DnsSetCommands = FlattenAdmingroupDnsSetCommands(ctx, from.DnsSetCommands, diags)
	m.DnsShowCommands = FlattenAdmingroupDnsShowCommands(ctx, from.DnsShowCommands, diags)
	m.DnsToplevelCommands = FlattenAdmingroupDnsToplevelCommands(ctx, from.DnsToplevelCommands, diags)
	m.DockerSetCommands = FlattenAdmingroupDockerSetCommands(ctx, from.DockerSetCommands, diags)
	m.DockerShowCommands = FlattenAdmingroupDockerShowCommands(ctx, from.DockerShowCommands, diags)
	m.EmailAddresses = flex.FlattenFrameworkUnorderedList(ctx, types.StringType, from.EmailAddresses, diags)
	m.EnableRestrictedUserAccess = types.BoolPointerValue(from.EnableRestrictedUserAccess)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.GridSetCommands = FlattenAdmingroupGridSetCommands(ctx, from.GridSetCommands, diags)
	m.GridShowCommands = FlattenAdmingroupGridShowCommands(ctx, from.GridShowCommands, diags)
	m.InactivityLockoutSetting = FlattenAdmingroupInactivityLockoutSetting(ctx, from.InactivityLockoutSetting, diags)
	m.LicensingSetCommands = FlattenAdmingroupLicensingSetCommands(ctx, from.LicensingSetCommands, diags)
	m.LicensingShowCommands = FlattenAdmingroupLicensingShowCommands(ctx, from.LicensingShowCommands, diags)
	m.LockoutSetting = FlattenAdmingroupLockoutSetting(ctx, from.LockoutSetting, diags)
	m.MachineControlToplevelCommands = FlattenAdmingroupMachineControlToplevelCommands(ctx, from.MachineControlToplevelCommands, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NetworkingSetCommands = FlattenAdmingroupNetworkingSetCommands(ctx, from.NetworkingSetCommands, diags)
	m.NetworkingShowCommands = FlattenAdmingroupNetworkingShowCommands(ctx, from.NetworkingShowCommands, diags)
	m.PasswordSetting = FlattenAdmingroupPasswordSetting(ctx, from.PasswordSetting, diags)
	m.Roles = flex.FlattenFrameworkListString(ctx, from.Roles, diags)
	m.SamlSetting = FlattenAdmingroupSamlSetting(ctx, from.SamlSetting, diags)
	m.SecuritySetCommands = FlattenAdmingroupSecuritySetCommands(ctx, from.SecuritySetCommands, diags)
	m.SecurityShowCommands = FlattenAdmingroupSecurityShowCommands(ctx, from.SecurityShowCommands, diags)
	m.Superuser = types.BoolPointerValue(from.Superuser)
	m.TroubleShootingToplevelCommands = FlattenAdmingroupTroubleShootingToplevelCommands(ctx, from.TroubleShootingToplevelCommands, diags)
	m.UseAccountInactivityLockoutEnable = types.BoolPointerValue(from.UseAccountInactivityLockoutEnable)
	m.UseDisableConcurrentLogin = types.BoolPointerValue(from.UseDisableConcurrentLogin)
	m.UseLockoutSetting = types.BoolPointerValue(from.UseLockoutSetting)
	m.UsePasswordSetting = types.BoolPointerValue(from.UsePasswordSetting)
	m.UserAccess = flex.FlattenFrameworkListNestedBlock(ctx, from.UserAccess, AdmingroupUserAccessAttrTypes, diags, FlattenAdmingroupUserAccess)
}
