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

type GridCloudapiModel struct {
	Ref              types.String `tfsdk:"ref"`
	AllowApiAdmins   types.String `tfsdk:"allow_api_admins"`
	AllowedApiAdmins types.List   `tfsdk:"allowed_api_admins"`
	EnableRecycleBin types.Bool   `tfsdk:"enable_recycle_bin"`
	GatewayConfig    types.Object `tfsdk:"gateway_config"`
}

var GridCloudapiAttrTypes = map[string]attr.Type{
	"ref":                types.StringType,
	"allow_api_admins":   types.StringType,
	"allowed_api_admins": types.ListType{ElemType: types.ObjectType{AttrTypes: GridCloudapiAllowedApiAdminsAttrTypes}},
	"enable_recycle_bin": types.BoolType,
	"gateway_config":     types.ObjectType{AttrTypes: GridCloudapiGatewayConfigAttrTypes},
}

var GridCloudapiResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"allow_api_admins": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Defines administrators who can perform cloud API requests on the Grid Master. The valid value is NONE (no administrator), ALL (all administrators), or LIST (administrators on the ACL).",
	},
	"allowed_api_admins": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridCloudapiAllowedApiAdminsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of administrators who can perform cloud API requests on the Cloud Platform Appliance.",
	},
	"enable_recycle_bin": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the recycle bin for deleted cloud objects is enabled or not on the Grid Master.",
	},
	"gateway_config": schema.SingleNestedAttribute{
		Attributes: GridCloudapiGatewayConfigResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandGridCloudapi(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridCloudapi {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridCloudapiModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridCloudapiModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridCloudapi {
	if m == nil {
		return nil
	}
	to := &grid.GridCloudapi{
		Ref:              flex.ExpandStringPointer(m.Ref),
		AllowApiAdmins:   flex.ExpandStringPointer(m.AllowApiAdmins),
		AllowedApiAdmins: flex.ExpandFrameworkListNestedBlock(ctx, m.AllowedApiAdmins, diags, ExpandGridCloudapiAllowedApiAdmins),
		EnableRecycleBin: flex.ExpandBoolPointer(m.EnableRecycleBin),
		GatewayConfig:    ExpandGridCloudapiGatewayConfig(ctx, m.GatewayConfig, diags),
	}
	return to
}

func FlattenGridCloudapi(ctx context.Context, from *grid.GridCloudapi, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridCloudapiAttrTypes)
	}
	m := GridCloudapiModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridCloudapiAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridCloudapiModel) Flatten(ctx context.Context, from *grid.GridCloudapi, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridCloudapiModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AllowApiAdmins = flex.FlattenStringPointer(from.AllowApiAdmins)
	m.AllowedApiAdmins = flex.FlattenFrameworkListNestedBlock(ctx, from.AllowedApiAdmins, GridCloudapiAllowedApiAdminsAttrTypes, diags, FlattenGridCloudapiAllowedApiAdmins)
	m.EnableRecycleBin = types.BoolPointerValue(from.EnableRecycleBin)
	m.GatewayConfig = FlattenGridCloudapiGatewayConfig(ctx, from.GatewayConfig, diags)
}
