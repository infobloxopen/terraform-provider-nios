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

type GmcgroupModel struct {
	Ref                types.String `tfsdk:"ref"`
    Uuid        types.String `tfsdk:"uuid"`
	Comment            types.String `tfsdk:"comment"`
	GmcPromotionPolicy types.String `tfsdk:"gmc_promotion_policy"`
	Members            types.List   `tfsdk:"members"`
	Name               types.String `tfsdk:"name"`
	ScheduledTime      types.Int64  `tfsdk:"scheduled_time"`
	TimeZone           types.String `tfsdk:"time_zone"`
}

var GmcgroupAttrTypes = map[string]attr.Type{
	"ref":                  types.StringType,
    "uuid":        types.StringType,
	"comment":              types.StringType,
	"gmc_promotion_policy": types.StringType,
	"members":              types.ListType{ElemType: types.ObjectType{AttrTypes: GmcgroupMembersAttrTypes}},
	"name":                 types.StringType,
	"scheduled_time":       types.Int64Type,
	"time_zone":            types.StringType,
}

var GmcgroupResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
    "uuid": schema.StringAttribute{
        Computed:            true,
        MarkdownDescription: "The uuid to the object.",
    },
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Description of the group",
	},
	"gmc_promotion_policy": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "This field decides whether the members join back at the same time or sequentially with time gap of 30 seconds.",
	},
	"members": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GmcgroupMembersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "gmcgroup members",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Group name",
	},
	"scheduled_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Absolute time at which the reconnect starts",
	},
	"time_zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The time zone for scheduling operations.",
	},
}

func ExpandGmcgroup(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.Gmcgroup {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GmcgroupModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GmcgroupModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.Gmcgroup {
	if m == nil {
		return nil
	}
	to := &grid.Gmcgroup{
		Ref:                flex.ExpandStringPointer(m.Ref),
		Comment:            flex.ExpandStringPointer(m.Comment),
		GmcPromotionPolicy: flex.ExpandStringPointer(m.GmcPromotionPolicy),
		Members:            flex.ExpandFrameworkListNestedBlock(ctx, m.Members, diags, ExpandGmcgroupMembers),
		Name:               flex.ExpandStringPointer(m.Name),
		ScheduledTime:      flex.ExpandInt64Pointer(m.ScheduledTime),
	}
	return to
}

func FlattenGmcgroup(ctx context.Context, from *grid.Gmcgroup, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GmcgroupAttrTypes)
	}
	m := GmcgroupModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GmcgroupAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GmcgroupModel) Flatten(ctx context.Context, from *grid.Gmcgroup, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GmcgroupModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
    m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.GmcPromotionPolicy = flex.FlattenStringPointer(from.GmcPromotionPolicy)
	m.Members = flex.FlattenFrameworkListNestedBlock(ctx, from.Members, GmcgroupMembersAttrTypes, diags, FlattenGmcgroupMembers)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.ScheduledTime = flex.FlattenInt64Pointer(from.ScheduledTime)
	m.TimeZone = flex.FlattenStringPointer(from.TimeZone)
}
