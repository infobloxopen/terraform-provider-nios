package misc

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

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type DxlEndpointModel struct {
	Ref                        types.String `tfsdk:"ref"`
	Uuid                       types.String `tfsdk:"uuid"`
	UUID                       types.String `tfsdk:"uuid"`
	Brokers                    types.List   `tfsdk:"brokers"`
	BrokersImportToken         types.String `tfsdk:"brokers_import_token"`
	BrokersImportFile          types.String `tfsdk:"brokers_import_file"`
	ClientCertificateSubject   types.String `tfsdk:"client_certificate_subject"`
	ClientCertificateToken     types.String `tfsdk:"client_certificate_token"`
	ClientCertificateFile      types.String `tfsdk:"client_certificate_file"`
	ClientCertificateValidFrom types.Int64  `tfsdk:"client_certificate_valid_from"`
	ClientCertificateValidTo   types.Int64  `tfsdk:"client_certificate_valid_to"`
	Comment                    types.String `tfsdk:"comment"`
	Disable                    types.Bool   `tfsdk:"disable"`
	ExtAttrs                   types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll                types.Map    `tfsdk:"extattrs_all"`
	LogLevel                   types.String `tfsdk:"log_level"`
	Name                       types.String `tfsdk:"name"`
	OutboundMemberType         types.String `tfsdk:"outbound_member_type"`
	OutboundMembers            types.List   `tfsdk:"outbound_members"`
	TemplateInstance           types.Object `tfsdk:"template_instance"`
	Timeout                    types.Int64  `tfsdk:"timeout"`
	Topics                     types.List   `tfsdk:"topics"`
	VendorIdentifier           types.String `tfsdk:"vendor_identifier"`
	WapiUserName               types.String `tfsdk:"wapi_user_name"`
	WapiUserPassword           types.String `tfsdk:"wapi_user_password"`
}

var DxlEndpointAttrTypes = map[string]attr.Type{
	"ref":                           types.StringType,
	"uuid":                          types.StringType,
	"brokers":                       types.ListType{ElemType: types.ObjectType{AttrTypes: DxlEndpointBrokersAttrTypes}},
	"brokers_import_token":          types.StringType,
	"brokers_import_file":           types.StringType,
	"client_certificate_subject":    types.StringType,
	"client_certificate_token":      types.StringType,
	"client_certificate_file":       types.StringType,
	"client_certificate_valid_from": types.Int64Type,
	"client_certificate_valid_to":   types.Int64Type,
	"comment":                       types.StringType,
	"disable":                       types.BoolType,
	"extattrs":                      types.MapType{ElemType: types.StringType},
	"extattrs_all":                  types.MapType{ElemType: types.StringType},
	"log_level":                     types.StringType,
	"name":                          types.StringType,
	"outbound_member_type":          types.StringType,
	"outbound_members":              types.ListType{ElemType: types.StringType},
	"template_instance":             types.ObjectType{AttrTypes: DxlEndpointTemplateInstanceAttrTypes},
	"timeout":                       types.Int64Type,
	"topics":                        types.ListType{ElemType: types.StringType},
	"vendor_identifier":             types.StringType,
	"wapi_user_name":                types.StringType,
	"wapi_user_password":            types.StringType,
}

var DxlEndpointResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Universally Unique ID assigned for this object.",
	},
	"brokers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: DxlEndpointBrokersResourceSchemaAttributes,
		},
		Computed: true,
		Optional: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The list of DXL endpoint brokers. Note that you cannot specify brokers and brokers_import_token at the same time.",
	},
	"brokers_import_token": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The token returned by the uploadinit function call in object fileop for a DXL broker configuration file. Note that you cannot specify brokers and brokers_import_token at the same time.",
	},
	"brokers_import_file": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The file path for the DXL broker configuration file. When specified, the file is uploaded and the resulting configuration is stored in the brokers field.",
	},
	"client_certificate_subject": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The client certificate subject of a DXL endpoint.",
	},
	"client_certificate_token": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The token returned by the uploadinit function call in object fileop for a DXL endpoint client certificate.",
	},
	"client_certificate_file": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The file path for the DXL endpoint client certificate.",
	},
	"client_certificate_valid_from": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The timestamp when client certificate for a DXL endpoint was created.",
	},
	"client_certificate_valid_to": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The timestamp when the client certificate for a DXL endpoint expires.",
	},
	"comment": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The comment of a DXL endpoint.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a DXL endpoint is disabled.",
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
		Computed: true,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
		MarkdownDescription: "Extensible attributes associated with the object, including default and internal attributes.",
		ElementType:         types.StringType,
	},
	"log_level": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Default:  stringdefault.StaticString("WARNING"),
		Validators: []validator.String{
			stringvalidator.OneOf("DEBUG", "ERROR", "INFO", "WARNING"),
		},
		MarkdownDescription: "The log level for a DXL endpoint.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of a DXL endpoint.",
	},
	"outbound_member_type": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("GM", "MEMBER"),
		},
		MarkdownDescription: "The outbound member that will generate events.",
	},
	"outbound_members": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The list of members for outbound events.",
	},
	"template_instance": schema.SingleNestedAttribute{
		Attributes:          DxlEndpointTemplateInstanceResourceSchemaAttributes,
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The DXL template instance. You cannot change the parameters of the DXL endpoint template instance.",
	},
	"timeout": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(30),
		MarkdownDescription: "The timeout of session management (in seconds).",
	},
	"topics": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "DXL topics",
	},
	"vendor_identifier": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The vendor identifier.",
	},
	"wapi_user_name": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("wapi_user_password")),
		},
		MarkdownDescription: "The user name for WAPI integration.",
	},
	"wapi_user_password": schema.StringAttribute{
		Sensitive: true,
		Optional:  true,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("wapi_user_name")),
		},
		MarkdownDescription: "The user password for WAPI integration.",
	},
}

func (m *DxlEndpointModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.DxlEndpoint {
	if m == nil {
		return nil
	}
	to := &misc.DxlEndpoint{
		Brokers:                flex.ExpandFrameworkListNestedBlockEmptyAsNil(ctx, m.Brokers, diags, ExpandDxlEndpointBrokers),
		BrokersImportToken:     flex.ExpandStringPointer(m.BrokersImportToken),
		ClientCertificateToken: flex.ExpandStringPointer(m.ClientCertificateToken),
		Comment:                flex.ExpandStringPointer(m.Comment),
		Disable:                flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:               ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		LogLevel:               flex.ExpandStringPointer(m.LogLevel),
		Name:                   flex.ExpandStringPointer(m.Name),
		OutboundMemberType:     flex.ExpandStringPointer(m.OutboundMemberType),
		OutboundMembers:        flex.ExpandFrameworkListString(ctx, m.OutboundMembers, diags),
		TemplateInstance:       ExpandDxlEndpointTemplateInstance(ctx, m.TemplateInstance, diags),
		Timeout:                flex.ExpandInt64Pointer(m.Timeout),
		Topics:                 flex.ExpandFrameworkListStringEmptyAsNil(ctx, m.Topics, diags),
		VendorIdentifier:       flex.ExpandStringPointer(m.VendorIdentifier),
		WapiUserName:           flex.ExpandStringPointer(m.WapiUserName),
		WapiUserPassword:       flex.ExpandStringPointer(m.WapiUserPassword),
	}
	return to
}

func FlattenDxlEndpoint(ctx context.Context, from *misc.DxlEndpoint, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DxlEndpointAttrTypes)
	}
	m := DxlEndpointModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, DxlEndpointAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DxlEndpointModel) Flatten(ctx context.Context, from *misc.DxlEndpoint, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DxlEndpointModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Brokers = flex.FlattenFrameworkListNestedBlock(ctx, from.Brokers, DxlEndpointBrokersAttrTypes, diags, FlattenDxlEndpointBrokers)
	m.BrokersImportToken = flex.FlattenStringPointer(from.BrokersImportToken)
	m.ClientCertificateSubject = flex.FlattenStringPointer(from.ClientCertificateSubject)
	m.ClientCertificateToken = flex.FlattenStringPointer(from.ClientCertificateToken)
	m.ClientCertificateValidFrom = flex.FlattenInt64Pointer(from.ClientCertificateValidFrom)
	m.ClientCertificateValidTo = flex.FlattenInt64Pointer(from.ClientCertificateValidTo)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.LogLevel = flex.FlattenStringPointer(from.LogLevel)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.OutboundMemberType = flex.FlattenStringPointer(from.OutboundMemberType)
	m.OutboundMembers = flex.FlattenFrameworkListString(ctx, from.OutboundMembers, diags)
	m.TemplateInstance = FlattenDxlEndpointTemplateInstance(ctx, from.TemplateInstance, diags)
	m.Timeout = flex.FlattenInt64Pointer(from.Timeout)
	m.Topics = flex.FlattenFrameworkListString(ctx, from.Topics, diags)
	m.VendorIdentifier = flex.FlattenStringPointer(from.VendorIdentifier)
	m.WapiUserName = flex.FlattenStringPointer(from.WapiUserName)
}
