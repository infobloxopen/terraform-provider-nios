package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type AdAuthServiceModel struct {
	Ref                 types.String `tfsdk:"ref"`
    Uuid        types.String `tfsdk:"uuid"`
	AdDomain            types.String `tfsdk:"ad_domain"`
	Comment             types.String `tfsdk:"comment"`
	Disabled            types.Bool   `tfsdk:"disabled"`
	DomainControllers   types.List   `tfsdk:"domain_controllers"`
	Name                types.String `tfsdk:"name"`
	NestedGroupQuerying types.Bool   `tfsdk:"nested_group_querying"`
	Timeout             types.Int64  `tfsdk:"timeout"`
}

var AdAuthServiceAttrTypes = map[string]attr.Type{
	"ref":                   types.StringType,
    "uuid":        types.StringType,
	"ad_domain":             types.StringType,
	"comment":               types.StringType,
	"disabled":              types.BoolType,
	"domain_controllers":    types.ListType{ElemType: types.ObjectType{AttrTypes: AdAuthServiceDomainControllersAttrTypes}},
	"name":                  types.StringType,
	"nested_group_querying": types.BoolType,
	"timeout":               types.Int64Type,
}

var AdAuthServiceResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
    "uuid": schema.StringAttribute{
        Computed:            true,
        MarkdownDescription: "The uuid to the object.",
    },
	"ad_domain": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Active Directory domain to which this server belongs.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The descriptive comment for the AD authentication service.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if Active Directory Authentication Service is disabled.",
	},
	"domain_controllers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AdAuthServiceDomainControllersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The AD authentication server list.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The AD authentication service name.",
	},
	"nested_group_querying": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the nested group querying is enabled.",
	},
	"timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of seconds that the appliance waits for a response from the AD server.",
	},
}

func ExpandAdAuthService(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdAuthService {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdAuthServiceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdAuthServiceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdAuthService {
	if m == nil {
		return nil
	}
	to := &security.AdAuthService{
		Ref:                 flex.ExpandStringPointer(m.Ref),
		AdDomain:            flex.ExpandStringPointer(m.AdDomain),
		Comment:             flex.ExpandStringPointer(m.Comment),
		Disabled:            flex.ExpandBoolPointer(m.Disabled),
		DomainControllers:   flex.ExpandFrameworkListNestedBlock(ctx, m.DomainControllers, diags, ExpandAdAuthServiceDomainControllers),
		Name:                flex.ExpandStringPointer(m.Name),
		NestedGroupQuerying: flex.ExpandBoolPointer(m.NestedGroupQuerying),
		Timeout:             flex.ExpandInt64Pointer(m.Timeout),
	}
	return to
}

func FlattenAdAuthService(ctx context.Context, from *security.AdAuthService, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdAuthServiceAttrTypes)
	}
	m := AdAuthServiceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdAuthServiceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdAuthServiceModel) Flatten(ctx context.Context, from *security.AdAuthService, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdAuthServiceModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
    m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.AdDomain = flex.FlattenStringPointer(from.AdDomain)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disabled = types.BoolPointerValue(from.Disabled)
	m.DomainControllers = flex.FlattenFrameworkListNestedBlock(ctx, from.DomainControllers, AdAuthServiceDomainControllersAttrTypes, diags, FlattenAdAuthServiceDomainControllers)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NestedGroupQuerying = types.BoolPointerValue(from.NestedGroupQuerying)
	m.Timeout = flex.FlattenInt64Pointer(from.Timeout)
}
