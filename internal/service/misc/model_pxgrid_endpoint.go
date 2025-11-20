package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
)

type PxgridEndpointModel struct {
	Ref                        types.String `tfsdk:"ref"`
    Uuid        types.String `tfsdk:"uuid"`
	Address                    types.String `tfsdk:"address"`
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
	NetworkView                types.String `tfsdk:"network_view"`
	OutboundMemberType         types.String `tfsdk:"outbound_member_type"`
	OutboundMembers            types.List   `tfsdk:"outbound_members"`
	PublishSettings            types.Object `tfsdk:"publish_settings"`
	SubscribeSettings          types.Object `tfsdk:"subscribe_settings"`
	TemplateInstance           types.Object `tfsdk:"template_instance"`
	Timeout                    types.Int64  `tfsdk:"timeout"`
	VendorIdentifier           types.String `tfsdk:"vendor_identifier"`
	WapiUserName               types.String `tfsdk:"wapi_user_name"`
	WapiUserPassword           types.String `tfsdk:"wapi_user_password"`
}

var PxgridEndpointAttrTypes = map[string]attr.Type{
	"ref":                           types.StringType,
    "uuid":        types.StringType,
	"address":                       types.StringType,
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
	"network_view":                  types.StringType,
	"outbound_member_type":          types.StringType,
	"outbound_members":              types.ListType{ElemType: types.StringType},
	"publish_settings":              types.ObjectType{AttrTypes: PxgridEndpointPublishSettingsAttrTypes},
	"subscribe_settings":            types.ObjectType{AttrTypes: PxgridEndpointSubscribeSettingsAttrTypes},
	"template_instance":             types.ObjectType{AttrTypes: PxgridEndpointTemplateInstanceAttrTypes},
	"timeout":                       types.Int64Type,
	"vendor_identifier":             types.StringType,
	"wapi_user_name":                types.StringType,
	"wapi_user_password":            types.StringType,
}

var PxgridEndpointResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
    "uuid": schema.StringAttribute{
        Computed:            true,
        MarkdownDescription: "The uuid to the object.",
    },
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The pxgrid endpoint IPv4 Address or IPv6 Address or Fully-Qualified Domain Name (FQDN)",
	},
	"client_certificate_subject": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Cisco ISE client certificate subject.",
	},
	"client_certificate_token": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The token returned by the uploadinit function call in object fileop for Cisco ISE client certificate.",
	},
	"client_certificate_valid_from": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The pxgrid endpoint client certificate valid from.",
	},
	"client_certificate_valid_to": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The pxgrid endpoint client certificate valid to.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Cisco ISE endpoint descriptive comment.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether a Cisco ISE endpoint is disabled or not. When this is set to False, the Cisco ISE endpoint is enabled.",
	},
	"extattrs": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"extattrs_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The log level for a notification pxgrid endpoint.",
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the pxgrid endpoint.",
	},
	"network_view": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The pxgrid network view name.",
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
	"publish_settings": schema.SingleNestedAttribute{
		Attributes: PxgridEndpointPublishSettingsResourceSchemaAttributes,
		Optional:   true,
	},
	"subscribe_settings": schema.SingleNestedAttribute{
		Attributes: PxgridEndpointSubscribeSettingsResourceSchemaAttributes,
		Optional:   true,
	},
	"template_instance": schema.SingleNestedAttribute{
		Attributes: PxgridEndpointTemplateInstanceResourceSchemaAttributes,
		Optional:   true,
	},
	"timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The timeout of session management (in seconds).",
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

func ExpandPxgridEndpoint(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.PxgridEndpoint {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m PxgridEndpointModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *PxgridEndpointModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.PxgridEndpoint {
	if m == nil {
		return nil
	}
	to := &misc.PxgridEndpoint{
		Ref:                    flex.ExpandStringPointer(m.Ref),
		Address:                flex.ExpandStringPointer(m.Address),
		ClientCertificateToken: flex.ExpandStringPointer(m.ClientCertificateToken),
		Comment:                flex.ExpandStringPointer(m.Comment),
		Disable:                flex.ExpandBoolPointer(m.Disable),
		ExtAttrs:               ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		LogLevel:               flex.ExpandStringPointer(m.LogLevel),
		Name:                   flex.ExpandStringPointer(m.Name),
		NetworkView:            flex.ExpandStringPointer(m.NetworkView),
		OutboundMemberType:     flex.ExpandStringPointer(m.OutboundMemberType),
		OutboundMembers:        flex.ExpandFrameworkListString(ctx, m.OutboundMembers, diags),
		PublishSettings:        ExpandPxgridEndpointPublishSettings(ctx, m.PublishSettings, diags),
		SubscribeSettings:      ExpandPxgridEndpointSubscribeSettings(ctx, m.SubscribeSettings, diags),
		TemplateInstance:       ExpandPxgridEndpointTemplateInstance(ctx, m.TemplateInstance, diags),
		Timeout:                flex.ExpandInt64Pointer(m.Timeout),
		VendorIdentifier:       flex.ExpandStringPointer(m.VendorIdentifier),
		WapiUserName:           flex.ExpandStringPointer(m.WapiUserName),
		WapiUserPassword:       flex.ExpandStringPointer(m.WapiUserPassword),
	}
	return to
}

func FlattenPxgridEndpoint(ctx context.Context, from *misc.PxgridEndpoint, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(PxgridEndpointAttrTypes)
	}
	m := PxgridEndpointModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, PxgridEndpointAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *PxgridEndpointModel) Flatten(ctx context.Context, from *misc.PxgridEndpoint, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = PxgridEndpointModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
    m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.ClientCertificateSubject = flex.FlattenStringPointer(from.ClientCertificateSubject)
	m.ClientCertificateToken = flex.FlattenStringPointer(from.ClientCertificateToken)
	m.ClientCertificateValidFrom = flex.FlattenInt64Pointer(from.ClientCertificateValidFrom)
	m.ClientCertificateValidTo = flex.FlattenInt64Pointer(from.ClientCertificateValidTo)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.LogLevel = flex.FlattenStringPointer(from.LogLevel)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.OutboundMemberType = flex.FlattenStringPointer(from.OutboundMemberType)
	m.OutboundMembers = flex.FlattenFrameworkListString(ctx, from.OutboundMembers, diags)
	m.PublishSettings = FlattenPxgridEndpointPublishSettings(ctx, from.PublishSettings, diags)
	m.SubscribeSettings = FlattenPxgridEndpointSubscribeSettings(ctx, from.SubscribeSettings, diags)
	m.TemplateInstance = FlattenPxgridEndpointTemplateInstance(ctx, from.TemplateInstance, diags)
	m.Timeout = flex.FlattenInt64Pointer(from.Timeout)
	m.VendorIdentifier = flex.FlattenStringPointer(from.VendorIdentifier)
	m.WapiUserName = flex.FlattenStringPointer(from.WapiUserName)
	m.WapiUserPassword = flex.FlattenStringPointer(from.WapiUserPassword)
}
