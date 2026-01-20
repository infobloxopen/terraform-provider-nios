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

type TaxiiModel struct {
	Ref            types.String `tfsdk:"ref"`
	Uuid           types.String `tfsdk:"uuid"`
	EnableService  types.Bool   `tfsdk:"enable_service"`
	Ipv4addr       types.String `tfsdk:"ipv4addr"`
	Ipv6addr       types.String `tfsdk:"ipv6addr"`
	Name           types.String `tfsdk:"name"`
	TaxiiRpzConfig types.List   `tfsdk:"taxii_rpz_config"`
}

var TaxiiAttrTypes = map[string]attr.Type{
	"ref":              types.StringType,
	"uuid":             types.StringType,
	"enable_service":   types.BoolType,
	"ipv4addr":         types.StringType,
	"ipv6addr":         types.StringType,
	"name":             types.StringType,
	"taxii_rpz_config": types.ListType{ElemType: types.ObjectType{AttrTypes: TaxiiTaxiiRpzConfigAttrTypes}},
}

var TaxiiResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The UUID of the object.",
	},
	"enable_service": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Indicates whether the Taxii service is running on the given member or not.",
	},
	"ipv4addr": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IPv4 Address of the Grid member.",
	},
	"ipv6addr": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IPv6 Address of the Grid member.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the Taxii Member.",
	},
	"taxii_rpz_config": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: TaxiiTaxiiRpzConfigResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "Taxii service RPZ configuration list.",
	},
}

func ExpandTaxii(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.Taxii {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m TaxiiModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *TaxiiModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.Taxii {
	if m == nil {
		return nil
	}
	to := &misc.Taxii{
		Ref:            flex.ExpandStringPointer(m.Ref),
		Uuid:           flex.ExpandStringPointer(m.Uuid),
		EnableService:  flex.ExpandBoolPointer(m.EnableService),
		TaxiiRpzConfig: flex.ExpandFrameworkListNestedBlock(ctx, m.TaxiiRpzConfig, diags, ExpandTaxiiTaxiiRpzConfig),
	}
	return to
}

func FlattenTaxii(ctx context.Context, from *misc.Taxii, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(TaxiiAttrTypes)
	}
	m := TaxiiModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, TaxiiAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *TaxiiModel) Flatten(ctx context.Context, from *misc.Taxii, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = TaxiiModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.EnableService = types.BoolPointerValue(from.EnableService)
	m.Ipv4addr = flex.FlattenStringPointer(from.Ipv4addr)
	m.Ipv6addr = flex.FlattenStringPointer(from.Ipv6addr)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.TaxiiRpzConfig = flex.FlattenFrameworkListNestedBlock(ctx, from.TaxiiRpzConfig, TaxiiTaxiiRpzConfigAttrTypes, diags, FlattenTaxiiTaxiiRpzConfig)
}
