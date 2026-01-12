package notification

import (
	"context"

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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/notification"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	"github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/plancontrol"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type NotificationRestEndpointModel struct {
	Ref                        types.String `tfsdk:"ref"`
	ClientCertificateSubject   types.String `tfsdk:"client_certificate_subject"`
	ClientCertificateToken     types.String `tfsdk:"client_certificate_token"`
	ClientCertificateFile      types.String `tfsdk:"client_certificate_file"`
	ClientCertificateValidFrom types.Int64  `tfsdk:"client_certificate_valid_from"`
	ClientCertificateValidTo   types.Int64  `tfsdk:"client_certificate_valid_to"`
	Comment                    types.String `tfsdk:"comment"`
	ExtAttrs                   types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll                types.Map    `tfsdk:"extattrs_all"`
	LogLevel                   types.String `tfsdk:"log_level"`
	Name                       types.String `tfsdk:"name"`
	OutboundMemberType         types.String `tfsdk:"outbound_member_type"`
	OutboundMembers            types.List   `tfsdk:"outbound_members"`
	Password                   types.String `tfsdk:"password"`
	ServerCertValidation       types.String `tfsdk:"server_cert_validation"`
	SyncDisabled               types.Bool   `tfsdk:"sync_disabled"`
	TemplateInstance           types.Object `tfsdk:"template_instance"`
	Timeout                    types.Int64  `tfsdk:"timeout"`
	Uri                        types.String `tfsdk:"uri"`
	Username                   types.String `tfsdk:"username"`
	VendorIdentifier           types.String `tfsdk:"vendor_identifier"`
	WapiUserName               types.String `tfsdk:"wapi_user_name"`
	WapiUserPassword           types.String `tfsdk:"wapi_user_password"`
}

var NotificationRestEndpointAttrTypes = map[string]attr.Type{
	"ref":                           types.StringType,
	"client_certificate_subject":    types.StringType,
	"client_certificate_token":      types.StringType,
	"client_certificate_file":       types.StringType,
	"client_certificate_valid_from": types.Int64Type,
	"client_certificate_valid_to":   types.Int64Type,
	"comment":                       types.StringType,
	"extattrs":                      types.MapType{ElemType: types.StringType},
	"extattrs_all":                  types.MapType{ElemType: types.StringType},
	"log_level":                     types.StringType,
	"name":                          types.StringType,
	"outbound_member_type":          types.StringType,
	"outbound_members":              types.ListType{ElemType: types.StringType},
	"password":                      types.StringType,
	"server_cert_validation":        types.StringType,
	"sync_disabled":                 types.BoolType,
	"template_instance":             types.ObjectType{AttrTypes: NotificationRestEndpointTemplateInstanceAttrTypes},
	"timeout":                       types.Int64Type,
	"uri":                           types.StringType,
	"username":                      types.StringType,
	"vendor_identifier":             types.StringType,
	"wapi_user_name":                types.StringType,
	"wapi_user_password":            types.StringType,
}

var NotificationRestEndpointResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"client_certificate_subject": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The client certificate subject of a notification REST endpoint.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"client_certificate_token": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The token returned by the uploadinit function call in object fileop for a notification REST endpoit client certificate.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"client_certificate_file": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The token returned by the uploadinit function call in object fileop for a notification REST endpoit client certificate.",
	},
	"client_certificate_valid_from": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The timestamp when client certificate for a notification REST endpoint was created.",
		PlanModifiers: []planmodifier.Int64{
			plancontrol.UseStateForUnknownInt64(),
		},
	},
	"client_certificate_valid_to": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The timestamp when client certificate for a notification REST endpoint expires.",
		PlanModifiers: []planmodifier.Int64{
			plancontrol.UseStateForUnknownInt64(),
		},
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "The comment of a notification REST endpoint.",
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
	"log_level": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("WARNING"),
		Validators: []validator.String{
			stringvalidator.OneOf("DEBUG", "INFO", "WARNING", "ERROR"),
		},
		MarkdownDescription: "The log level for a notification REST endpoint.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of a notification REST endpoint.",
	},
	"outbound_member_type": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("MEMBER", "GM"),
		},
		MarkdownDescription: "The outbound member which will generate an event.",
	},
	"outbound_members": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.SizeAtMost(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of members for outbound events.",
	},
	"password": schema.StringAttribute{
		Optional:  true,
		Sensitive: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("username")),
		},
		MarkdownDescription: "The password of the user that can log into a notification REST endpoint.",
	},
	"server_cert_validation": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("CA_CERT"),
		Validators: []validator.String{
			stringvalidator.OneOf("CA_CERT", "CA_CERT_NO_HOSTNAME", "NO_VALIDATION"),
		},
		MarkdownDescription: "The server certificate validation type.",
	},
	"sync_disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the sync process is disabled for a notification REST endpoint.",
	},
	"template_instance": schema.SingleNestedAttribute{
		Attributes:          NotificationRestEndpointTemplateInstanceResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The notification REST template instance.",
	},
	"timeout": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(30),
		MarkdownDescription: "The timeout of session management (in seconds).",
	},
	"uri": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The URI of a notification REST endpoint.",
	},
	"username": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
			stringvalidator.AlsoRequires(path.MatchRoot("password")),
		},
		MarkdownDescription: "The username of the user that can log into a notification REST endpoint.",
	},
	"vendor_identifier": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The vendor identifier.",
	},
	"wapi_user_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("wapi_user_password")),
		},
		MarkdownDescription: "The user name for WAPI integration.",
	},
	"wapi_user_password": schema.StringAttribute{
		Optional:  true,
		Sensitive: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("wapi_user_name")),
		},
		MarkdownDescription: "The user password for WAPI integration.",
	},
}

func (m *NotificationRestEndpointModel) Expand(ctx context.Context, diags *diag.Diagnostics) *notification.NotificationRestEndpoint {
	if m == nil {
		return nil
	}
	to := &notification.NotificationRestEndpoint{
		Comment:              flex.ExpandStringPointer(m.Comment),
		ExtAttrs:             ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		LogLevel:             flex.ExpandStringPointer(m.LogLevel),
		Name:                 flex.ExpandStringPointer(m.Name),
		OutboundMemberType:   flex.ExpandStringPointer(m.OutboundMemberType),
		OutboundMembers:      flex.ExpandFrameworkListString(ctx, m.OutboundMembers, diags),
		Password:             flex.ExpandStringPointer(m.Password),
		ServerCertValidation: flex.ExpandStringPointer(m.ServerCertValidation),
		SyncDisabled:         flex.ExpandBoolPointer(m.SyncDisabled),
		TemplateInstance:     ExpandNotificationRestEndpointTemplateInstance(ctx, m.TemplateInstance, diags),
		Timeout:              flex.ExpandInt64Pointer(m.Timeout),
		Uri:                  flex.ExpandStringPointer(m.Uri),
		Username:             flex.ExpandStringPointer(m.Username),
		VendorIdentifier:     flex.ExpandStringPointer(m.VendorIdentifier),
		WapiUserName:         flex.ExpandStringPointer(m.WapiUserName),
		WapiUserPassword:     flex.ExpandStringPointer(m.WapiUserPassword),
	}
	return to
}

func FlattenNotificationRestEndpoint(ctx context.Context, from *notification.NotificationRestEndpoint, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NotificationRestEndpointAttrTypes)
	}
	m := NotificationRestEndpointModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, NotificationRestEndpointAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NotificationRestEndpointModel) Flatten(ctx context.Context, from *notification.NotificationRestEndpoint, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NotificationRestEndpointModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.ClientCertificateSubject = flex.FlattenStringPointer(from.ClientCertificateSubject)
	m.ClientCertificateToken = flex.FlattenStringPointer(from.ClientCertificateToken)
	m.ClientCertificateValidFrom = flex.FlattenInt64Pointer(from.ClientCertificateValidFrom)
	m.ClientCertificateValidTo = flex.FlattenInt64Pointer(from.ClientCertificateValidTo)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.LogLevel = flex.FlattenStringPointer(from.LogLevel)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.OutboundMemberType = flex.FlattenStringPointer(from.OutboundMemberType)
	m.OutboundMembers = flex.FlattenFrameworkListString(ctx, from.OutboundMembers, diags)
	m.ServerCertValidation = flex.FlattenStringPointer(from.ServerCertValidation)
	m.SyncDisabled = types.BoolPointerValue(from.SyncDisabled)
	m.TemplateInstance = FlattenNotificationRestEndpointTemplateInstance(ctx, from.TemplateInstance, diags)
	m.Timeout = flex.FlattenInt64Pointer(from.Timeout)
	m.Uri = flex.FlattenStringPointer(from.Uri)
	m.Username = flex.FlattenStringPointer(from.Username)
	m.VendorIdentifier = flex.FlattenStringPointer(from.VendorIdentifier)
	m.WapiUserName = flex.FlattenStringPointer(from.WapiUserName)
}
