package grid

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

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
)

type GridMemberCloudapiModel struct {
	Ref              types.String `tfsdk:"ref"`
	Uuid             types.String `tfsdk:"uuid"`
	AllowApiAdmins   types.String `tfsdk:"allow_api_admins"`
	AllowedApiAdmins types.List   `tfsdk:"allowed_api_admins"`
	EnableService    types.Bool   `tfsdk:"enable_service"`
	ExtAttrs         types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll      types.Map    `tfsdk:"extattrs_all"`
	GatewayConfig    types.Object `tfsdk:"gateway_config"`
	Member           types.Object `tfsdk:"member"`
	Status           types.String `tfsdk:"status"`
}

var GridMemberCloudapiAttrTypes = map[string]attr.Type{
	"ref":                types.StringType,
	"uuid":               types.StringType,
	"allow_api_admins":   types.StringType,
	"allowed_api_admins": types.ListType{ElemType: types.ObjectType{AttrTypes: GridMemberCloudapiAllowedApiAdminsAttrTypes}},
	"enable_service":     types.BoolType,
	"extattrs":           types.MapType{ElemType: types.StringType},
	"extattrs_all":       types.MapType{ElemType: types.StringType},
	"gateway_config":     types.ObjectType{AttrTypes: GridMemberCloudapiGatewayConfigAttrTypes},
	"member":             types.ObjectType{AttrTypes: GridMemberCloudapiMemberAttrTypes},
	"status":             types.StringType,
}

var GridMemberCloudapiResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The UUID of the object.",
	},
	"allow_api_admins": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Defines which administrators are allowed to perform Cloud API request on the Grid Member: no administrators (NONE), any administrators (ALL) or administrators in the ACL list (LIST). Default is ALL.",
	},
	"allowed_api_admins": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridMemberCloudapiAllowedApiAdminsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "List of administrators allowed to perform Cloud Platform API requests on that member.",
	},
	"enable_service": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Controls whether the Cloud API service runs on the member or not.",
	},
	"extattrs": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"gateway_config": schema.SingleNestedAttribute{
		Attributes: GridMemberCloudapiGatewayConfigResourceSchemaAttributes,
		Optional:   true,
	},
	"member": schema.SingleNestedAttribute{
		Attributes: GridMemberCloudapiMemberResourceSchemaAttributes,
		Optional:   true,
	},
	"status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Status of Cloud API service on the member.",
	},
}

func ExpandGridMemberCloudapi(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridMemberCloudapi {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridMemberCloudapiModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridMemberCloudapiModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridMemberCloudapi {
	if m == nil {
		return nil
	}
	to := &grid.GridMemberCloudapi{
		Ref:              flex.ExpandStringPointer(m.Ref),
		Uuid:             flex.ExpandStringPointer(m.Uuid),
		AllowApiAdmins:   flex.ExpandStringPointer(m.AllowApiAdmins),
		AllowedApiAdmins: flex.ExpandFrameworkListNestedBlock(ctx, m.AllowedApiAdmins, diags, ExpandGridMemberCloudapiAllowedApiAdmins),
		EnableService:    flex.ExpandBoolPointer(m.EnableService),
		GatewayConfig:    ExpandGridMemberCloudapiGatewayConfig(ctx, m.GatewayConfig, diags),
		Member:           ExpandGridMemberCloudapiMember(ctx, m.Member, diags),
	}
	return to
}

func FlattenGridMemberCloudapi(ctx context.Context, from *grid.GridMemberCloudapi, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridMemberCloudapiAttrTypes)
	}
	m := GridMemberCloudapiModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, GridMemberCloudapiAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridMemberCloudapiModel) Flatten(ctx context.Context, from *grid.GridMemberCloudapi, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridMemberCloudapiModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.AllowApiAdmins = flex.FlattenStringPointer(from.AllowApiAdmins)
	m.AllowedApiAdmins = flex.FlattenFrameworkListNestedBlock(ctx, from.AllowedApiAdmins, GridMemberCloudapiAllowedApiAdminsAttrTypes, diags, FlattenGridMemberCloudapiAllowedApiAdmins)
	m.EnableService = types.BoolPointerValue(from.EnableService)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.GatewayConfig = FlattenGridMemberCloudapiGatewayConfig(ctx, from.GatewayConfig, diags)
	m.Member = FlattenGridMemberCloudapiMember(ctx, from.Member, diags)
	m.Status = flex.FlattenStringPointer(from.Status)
}
