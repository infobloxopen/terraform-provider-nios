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

type CaptiveportalModel struct {
	Ref                       types.String `tfsdk:"ref"`
	AuthnServerGroup          types.String `tfsdk:"authn_server_group"`
	CompanyName               types.String `tfsdk:"company_name"`
	EnableSyslogAuthFailure   types.Bool   `tfsdk:"enable_syslog_auth_failure"`
	EnableSyslogAuthSuccess   types.Bool   `tfsdk:"enable_syslog_auth_success"`
	EnableUserType            types.String `tfsdk:"enable_user_type"`
	Encryption                types.String `tfsdk:"encryption"`
	Files                     types.List   `tfsdk:"files"`
	GuestCustomField1Name     types.String `tfsdk:"guest_custom_field1_name"`
	GuestCustomField1Required types.Bool   `tfsdk:"guest_custom_field1_required"`
	GuestCustomField2Name     types.String `tfsdk:"guest_custom_field2_name"`
	GuestCustomField2Required types.Bool   `tfsdk:"guest_custom_field2_required"`
	GuestCustomField3Name     types.String `tfsdk:"guest_custom_field3_name"`
	GuestCustomField3Required types.Bool   `tfsdk:"guest_custom_field3_required"`
	GuestCustomField4Name     types.String `tfsdk:"guest_custom_field4_name"`
	GuestCustomField4Required types.Bool   `tfsdk:"guest_custom_field4_required"`
	GuestEmailRequired        types.Bool   `tfsdk:"guest_email_required"`
	GuestFirstNameRequired    types.Bool   `tfsdk:"guest_first_name_required"`
	GuestLastNameRequired     types.Bool   `tfsdk:"guest_last_name_required"`
	GuestMiddleNameRequired   types.Bool   `tfsdk:"guest_middle_name_required"`
	GuestPhoneRequired        types.Bool   `tfsdk:"guest_phone_required"`
	HelpdeskMessage           types.String `tfsdk:"helpdesk_message"`
	ListenAddressIp           types.String `tfsdk:"listen_address_ip"`
	ListenAddressType         types.String `tfsdk:"listen_address_type"`
	Name                      types.String `tfsdk:"name"`
	NetworkView               types.String `tfsdk:"network_view"`
	Port                      types.Int64  `tfsdk:"port"`
	ServiceEnabled            types.Bool   `tfsdk:"service_enabled"`
	SyslogAuthFailureLevel    types.String `tfsdk:"syslog_auth_failure_level"`
	SyslogAuthSuccessLevel    types.String `tfsdk:"syslog_auth_success_level"`
	WelcomeMessage            types.String `tfsdk:"welcome_message"`
}

var CaptiveportalAttrTypes = map[string]attr.Type{
	"ref":                          types.StringType,
	"authn_server_group":           types.StringType,
	"company_name":                 types.StringType,
	"enable_syslog_auth_failure":   types.BoolType,
	"enable_syslog_auth_success":   types.BoolType,
	"enable_user_type":             types.StringType,
	"encryption":                   types.StringType,
	"files":                        types.ListType{ElemType: types.ObjectType{AttrTypes: CaptiveportalFilesAttrTypes}},
	"guest_custom_field1_name":     types.StringType,
	"guest_custom_field1_required": types.BoolType,
	"guest_custom_field2_name":     types.StringType,
	"guest_custom_field2_required": types.BoolType,
	"guest_custom_field3_name":     types.StringType,
	"guest_custom_field3_required": types.BoolType,
	"guest_custom_field4_name":     types.StringType,
	"guest_custom_field4_required": types.BoolType,
	"guest_email_required":         types.BoolType,
	"guest_first_name_required":    types.BoolType,
	"guest_last_name_required":     types.BoolType,
	"guest_middle_name_required":   types.BoolType,
	"guest_phone_required":         types.BoolType,
	"helpdesk_message":             types.StringType,
	"listen_address_ip":            types.StringType,
	"listen_address_type":          types.StringType,
	"name":                         types.StringType,
	"network_view":                 types.StringType,
	"port":                         types.Int64Type,
	"service_enabled":              types.BoolType,
	"syslog_auth_failure_level":    types.StringType,
	"syslog_auth_success_level":    types.StringType,
	"welcome_message":              types.StringType,
}

var CaptiveportalResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"authn_server_group": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The authentication server group assigned to this captive portal.",
	},
	"company_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The company name that appears in the guest registration page.",
	},
	"enable_syslog_auth_failure": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if authentication failures are logged to syslog or not.",
	},
	"enable_syslog_auth_success": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if successful authentications are logged to syslog or not.",
	},
	"enable_user_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The type of user to be enabled for the captive portal.",
	},
	"encryption": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The encryption the captive portal uses.",
	},
	"files": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: CaptiveportalFilesResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of files associated with the captive portal.",
	},
	"guest_custom_field1_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the custom field that you are adding to the guest registration page.",
	},
	"guest_custom_field1_required": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the custom field is required or not.",
	},
	"guest_custom_field2_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the custom field that you are adding to the guest registration page.",
	},
	"guest_custom_field2_required": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the custom field is required or not.",
	},
	"guest_custom_field3_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the custom field that you are adding to the guest registration page.",
	},
	"guest_custom_field3_required": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the custom field is required or not.",
	},
	"guest_custom_field4_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the custom field that you are adding to the guest registration page.",
	},
	"guest_custom_field4_required": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the custom field is required or not.",
	},
	"guest_email_required": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the email address of the guest is required or not.",
	},
	"guest_first_name_required": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the first name of the guest is required or not.",
	},
	"guest_last_name_required": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the last name of the guest is required or not.",
	},
	"guest_middle_name_required": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the middle name of the guest is required or not.",
	},
	"guest_phone_required": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the phone number of the guest is required or not.",
	},
	"helpdesk_message": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The helpdesk message that appears in the guest registration page.",
	},
	"listen_address_ip": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Determines the IP address on which the captive portal listens. Valid if listen address type is 'IP'.",
	},
	"listen_address_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Determines the type of the IP address on which the captive portal listens.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The hostname of the Grid member that hosts the captive portal.",
	},
	"network_view": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The network view of the captive portal.",
	},
	"port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The TCP port used by the Captive Portal service. The port is required when the Captive Portal service is enabled. Valid values are between 1 and 63999. Please note that setting the port number to 80 or 443 might impact performance.",
	},
	"service_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the captive portal service is enabled or not.",
	},
	"syslog_auth_failure_level": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The syslog level at which authentication failures are logged.",
	},
	"syslog_auth_success_level": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The syslog level at which successful authentications are logged.",
	},
	"welcome_message": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The welcome message that appears in the guest registration page.",
	},
}

func ExpandCaptiveportal(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.Captiveportal {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m CaptiveportalModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *CaptiveportalModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.Captiveportal {
	if m == nil {
		return nil
	}
	to := &grid.Captiveportal{
		Ref:                       flex.ExpandStringPointer(m.Ref),
		AuthnServerGroup:          flex.ExpandStringPointer(m.AuthnServerGroup),
		CompanyName:               flex.ExpandStringPointer(m.CompanyName),
		EnableSyslogAuthFailure:   flex.ExpandBoolPointer(m.EnableSyslogAuthFailure),
		EnableSyslogAuthSuccess:   flex.ExpandBoolPointer(m.EnableSyslogAuthSuccess),
		EnableUserType:            flex.ExpandStringPointer(m.EnableUserType),
		Encryption:                flex.ExpandStringPointer(m.Encryption),
		Files:                     flex.ExpandFrameworkListNestedBlock(ctx, m.Files, diags, ExpandCaptiveportalFiles),
		GuestCustomField1Name:     flex.ExpandStringPointer(m.GuestCustomField1Name),
		GuestCustomField1Required: flex.ExpandBoolPointer(m.GuestCustomField1Required),
		GuestCustomField2Name:     flex.ExpandStringPointer(m.GuestCustomField2Name),
		GuestCustomField2Required: flex.ExpandBoolPointer(m.GuestCustomField2Required),
		GuestCustomField3Name:     flex.ExpandStringPointer(m.GuestCustomField3Name),
		GuestCustomField3Required: flex.ExpandBoolPointer(m.GuestCustomField3Required),
		GuestCustomField4Name:     flex.ExpandStringPointer(m.GuestCustomField4Name),
		GuestCustomField4Required: flex.ExpandBoolPointer(m.GuestCustomField4Required),
		GuestEmailRequired:        flex.ExpandBoolPointer(m.GuestEmailRequired),
		GuestFirstNameRequired:    flex.ExpandBoolPointer(m.GuestFirstNameRequired),
		GuestLastNameRequired:     flex.ExpandBoolPointer(m.GuestLastNameRequired),
		GuestMiddleNameRequired:   flex.ExpandBoolPointer(m.GuestMiddleNameRequired),
		GuestPhoneRequired:        flex.ExpandBoolPointer(m.GuestPhoneRequired),
		HelpdeskMessage:           flex.ExpandStringPointer(m.HelpdeskMessage),
		ListenAddressIp:           flex.ExpandStringPointer(m.ListenAddressIp),
		ListenAddressType:         flex.ExpandStringPointer(m.ListenAddressType),
		NetworkView:               flex.ExpandStringPointer(m.NetworkView),
		Port:                      flex.ExpandInt64Pointer(m.Port),
		ServiceEnabled:            flex.ExpandBoolPointer(m.ServiceEnabled),
		SyslogAuthFailureLevel:    flex.ExpandStringPointer(m.SyslogAuthFailureLevel),
		SyslogAuthSuccessLevel:    flex.ExpandStringPointer(m.SyslogAuthSuccessLevel),
		WelcomeMessage:            flex.ExpandStringPointer(m.WelcomeMessage),
	}
	return to
}

func FlattenCaptiveportal(ctx context.Context, from *grid.Captiveportal, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(CaptiveportalAttrTypes)
	}
	m := CaptiveportalModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, CaptiveportalAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *CaptiveportalModel) Flatten(ctx context.Context, from *grid.Captiveportal, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = CaptiveportalModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AuthnServerGroup = flex.FlattenStringPointer(from.AuthnServerGroup)
	m.CompanyName = flex.FlattenStringPointer(from.CompanyName)
	m.EnableSyslogAuthFailure = types.BoolPointerValue(from.EnableSyslogAuthFailure)
	m.EnableSyslogAuthSuccess = types.BoolPointerValue(from.EnableSyslogAuthSuccess)
	m.EnableUserType = flex.FlattenStringPointer(from.EnableUserType)
	m.Encryption = flex.FlattenStringPointer(from.Encryption)
	m.Files = flex.FlattenFrameworkListNestedBlock(ctx, from.Files, CaptiveportalFilesAttrTypes, diags, FlattenCaptiveportalFiles)
	m.GuestCustomField1Name = flex.FlattenStringPointer(from.GuestCustomField1Name)
	m.GuestCustomField1Required = types.BoolPointerValue(from.GuestCustomField1Required)
	m.GuestCustomField2Name = flex.FlattenStringPointer(from.GuestCustomField2Name)
	m.GuestCustomField2Required = types.BoolPointerValue(from.GuestCustomField2Required)
	m.GuestCustomField3Name = flex.FlattenStringPointer(from.GuestCustomField3Name)
	m.GuestCustomField3Required = types.BoolPointerValue(from.GuestCustomField3Required)
	m.GuestCustomField4Name = flex.FlattenStringPointer(from.GuestCustomField4Name)
	m.GuestCustomField4Required = types.BoolPointerValue(from.GuestCustomField4Required)
	m.GuestEmailRequired = types.BoolPointerValue(from.GuestEmailRequired)
	m.GuestFirstNameRequired = types.BoolPointerValue(from.GuestFirstNameRequired)
	m.GuestLastNameRequired = types.BoolPointerValue(from.GuestLastNameRequired)
	m.GuestMiddleNameRequired = types.BoolPointerValue(from.GuestMiddleNameRequired)
	m.GuestPhoneRequired = types.BoolPointerValue(from.GuestPhoneRequired)
	m.HelpdeskMessage = flex.FlattenStringPointer(from.HelpdeskMessage)
	m.ListenAddressIp = flex.FlattenStringPointer(from.ListenAddressIp)
	m.ListenAddressType = flex.FlattenStringPointer(from.ListenAddressType)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.ServiceEnabled = types.BoolPointerValue(from.ServiceEnabled)
	m.SyslogAuthFailureLevel = flex.FlattenStringPointer(from.SyslogAuthFailureLevel)
	m.SyslogAuthSuccessLevel = flex.FlattenStringPointer(from.SyslogAuthSuccessLevel)
	m.WelcomeMessage = flex.FlattenStringPointer(from.WelcomeMessage)
}
