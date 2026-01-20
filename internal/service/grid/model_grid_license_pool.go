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

type GridLicensePoolModel struct {
	Ref              types.String `tfsdk:"ref"`
	Uuid             types.String `tfsdk:"uuid"`
	Assigned         types.Int64  `tfsdk:"assigned"`
	ExpirationStatus types.String `tfsdk:"expiration_status"`
	ExpiryDate       types.Int64  `tfsdk:"expiry_date"`
	Installed        types.Int64  `tfsdk:"installed"`
	Key              types.String `tfsdk:"key"`
	Limit            types.String `tfsdk:"limit"`
	LimitContext     types.String `tfsdk:"limit_context"`
	Model            types.String `tfsdk:"model"`
	Subpools         types.List   `tfsdk:"subpools"`
	TempAssigned     types.Int64  `tfsdk:"temp_assigned"`
	Type             types.String `tfsdk:"type"`
}

var GridLicensePoolAttrTypes = map[string]attr.Type{
	"ref":               types.StringType,
	"uuid":              types.StringType,
	"assigned":          types.Int64Type,
	"expiration_status": types.StringType,
	"expiry_date":       types.Int64Type,
	"installed":         types.Int64Type,
	"key":               types.StringType,
	"limit":             types.StringType,
	"limit_context":     types.StringType,
	"model":             types.StringType,
	"subpools":          types.ListType{ElemType: types.ObjectType{AttrTypes: GridLicensePoolSubpoolsAttrTypes}},
	"temp_assigned":     types.Int64Type,
	"type":              types.StringType,
}

var GridLicensePoolResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The UUID of the object.",
	},
	"assigned": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of dynamic licenses allocated to vNIOS appliances.",
	},
	"expiration_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The license expiration status. * DELETED - The temporary license has been deleted. * EXPIRED - License/Pool has reached the expiry date. * EXPIRING_SOON - License/Pool expires in 31-90 days. * EXPIRING_VERY_SOON - License/Pool expires in 30 days or earlier. * NOT_EXPIRED - License/Pool has not expired. * PERMANENT - License/Pool does not expire.",
	},
	"expiry_date": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The expiration timestamp of the license.",
	},
	"installed": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The total number of dynamic licenses allowed for this license pool.",
	},
	"key": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The license string for the license pool.",
	},
	"limit": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The limitation of dynamic license that can be allocated from the license pool.",
	},
	"limit_context": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The license limit context.",
	},
	"model": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The supported vNIOS virtual appliance model.",
	},
	"subpools": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridLicensePoolSubpoolsResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Computed:            true,
		MarkdownDescription: "The license pool subpools.",
	},
	"temp_assigned": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The total number of temporary dynamic licenses allocated to vNIOS appliances.",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The license type.",
	},
}

func ExpandGridLicensePool(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridLicensePool {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridLicensePoolModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridLicensePoolModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridLicensePool {
	if m == nil {
		return nil
	}
	to := &grid.GridLicensePool{
		Ref: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenGridLicensePool(ctx context.Context, from *grid.GridLicensePool, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridLicensePoolAttrTypes)
	}
	m := GridLicensePoolModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridLicensePoolAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridLicensePoolModel) Flatten(ctx context.Context, from *grid.GridLicensePool, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridLicensePoolModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Assigned = flex.FlattenInt64Pointer(from.Assigned)
	m.ExpirationStatus = flex.FlattenStringPointer(from.ExpirationStatus)
	m.ExpiryDate = flex.FlattenInt64Pointer(from.ExpiryDate)
	m.Installed = flex.FlattenInt64Pointer(from.Installed)
	m.Key = flex.FlattenStringPointer(from.Key)
	m.Limit = flex.FlattenStringPointer(from.Limit)
	m.LimitContext = flex.FlattenStringPointer(from.LimitContext)
	m.Model = flex.FlattenStringPointer(from.Model)
	m.Subpools = flex.FlattenFrameworkListNestedBlock(ctx, from.Subpools, GridLicensePoolSubpoolsAttrTypes, diags, FlattenGridLicensePoolSubpools)
	m.TempAssigned = flex.FlattenInt64Pointer(from.TempAssigned)
	m.Type = flex.FlattenStringPointer(from.Type)
}
