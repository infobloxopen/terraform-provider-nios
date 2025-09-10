package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type UpgradegroupMembersModel struct {
	Member   types.String `tfsdk:"member"`
	TimeZone types.String `tfsdk:"time_zone"`
}

var UpgradegroupMembersAttrTypes = map[string]attr.Type{
	"member":    types.StringType,
	"time_zone": types.StringType,
}

var UpgradegroupMembersResourceSchemaAttributes = map[string]schema.Attribute{
	"member": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The upgrade group member name.",
	},
	"time_zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The upgrade group member time zone.",
	},
}

func ExpandUpgradegroupMembers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.UpgradegroupMembers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UpgradegroupMembersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *UpgradegroupMembersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.UpgradegroupMembers {
	if m == nil {
		return nil
	}
	to := &grid.UpgradegroupMembers{
		Member: flex.ExpandStringPointer(m.Member),
	}
	return to
}

func FlattenUpgradegroupMembers(ctx context.Context, from *grid.UpgradegroupMembers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UpgradegroupMembersAttrTypes)
	}
	m := UpgradegroupMembersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, UpgradegroupMembersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *UpgradegroupMembersModel) Flatten(ctx context.Context, from *grid.UpgradegroupMembers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = UpgradegroupMembersModel{}
	}
	m.Member = flex.FlattenStringPointer(from.Member)
	m.TimeZone = flex.FlattenStringPointer(from.TimeZone)
}
