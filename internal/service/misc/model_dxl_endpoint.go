package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type DxlEndpointModel struct {
	Ref                        types.String `tfsdk:"ref"`
	Brokers                    types.List   `tfsdk:"brokers"`
	BrokersImportToken         types.String `tfsdk:"brokers_import_token"`
	ClientCertificateSubject   types.String `tfsdk:"client_certificate_subject"`
	ClientCertificateToken     types.String `tfsdk:"client_certificate_token"`
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
	"brokers":                       types.ListType{ElemType: types.ObjectType{AttrTypes: DxlEndpointBrokersAttrTypes}},
	"brokers_import_token":          types.StringType,
	"client_certificate_subject":    types.StringType,
	"client_certificate_token":      types.StringType,
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
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"brokers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: DxlEndpointBrokersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of DXL endpoint brokers. Note that you cannot specify brokers and brokers_import_token at the same time.",
	},
	"brokers_import_token": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The token returned by the uploadinit function call in object fileop for a DXL broker configuration file. Note that you cannot specify brokers and brokers_import_token at the same time.",
	},
	"client_certificate_subject": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The client certificate subject of a DXL endpoint.",
	},
	"client_certificate_token": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The token returned by the uploadinit function call in object fileop for a DXL endpoint client certificate.",
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
		Optional:            true,
		MarkdownDescription: "The comment of a DXL endpoint.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether a DXL endpoint is disabled.",
	},
	"extattrs": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"extattrs_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The log level for a DXL endpoint.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of a DXL endpoint.",
	},
	"outbound_member_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The outbound member that will generate events.",
	},
	"outbound_members": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of members for outbound events.",
	},
	"template_instance": schema.SingleNestedAttribute{
		Attributes: DxlEndpointTemplateInstanceResourceSchemaAttributes,
		Optional:   true,
	},
	"timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The timeout of session management (in seconds).",
	},
	"topics": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "DXL topics",
	},
	"vendor_identifier": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The vendor identifier.",
	},
	"wapi_user_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The user name for WAPI integration.",
	},
	"wapi_user_password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The user password for WAPI integration.",
	},
}

func ExpandDxlEndpoint(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.DxlEndpoint {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DxlEndpointModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DxlEndpointModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.DxlEndpoint {
	if m == nil {
		return nil
	}
	to := &misc.DxlEndpoint{
		Ref:                    flex.ExpandStringPointer(m.Ref),
		Brokers:                flex.ExpandFrameworkListNestedBlock(ctx, m.Brokers, diags, ExpandDxlEndpointBrokers),
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
		Topics:                 flex.ExpandFrameworkListString(ctx, m.Topics, diags),
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
	m.WapiUserPassword = flex.FlattenStringPointer(from.WapiUserPassword)
}
