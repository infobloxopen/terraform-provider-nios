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

type GridntpsettingNtpKeysModel struct {
	Number types.Int64  `tfsdk:"number"`
	String types.String `tfsdk:"string"`
	Type   types.String `tfsdk:"type"`
}

var GridntpsettingNtpKeysAttrTypes = map[string]attr.Type{
	"number": types.Int64Type,
	"string": types.StringType,
	"type":   types.StringType,
}

var GridntpsettingNtpKeysResourceSchemaAttributes = map[string]schema.Attribute{
	"number": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The NTP authentication key identifier.",
	},
	"string": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The NTP authentication key string.",
	},
	"type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The NTP authentication key type.",
	},
}

func ExpandGridntpsettingNtpKeys(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridntpsettingNtpKeys {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridntpsettingNtpKeysModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridntpsettingNtpKeysModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridntpsettingNtpKeys {
	if m == nil {
		return nil
	}
	to := &grid.GridntpsettingNtpKeys{
		Number: flex.ExpandInt64Pointer(m.Number),
		String: flex.ExpandStringPointer(m.String),
		Type:   flex.ExpandStringPointer(m.Type),
	}
	return to
}

func FlattenGridntpsettingNtpKeys(ctx context.Context, from *grid.GridntpsettingNtpKeys, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridntpsettingNtpKeysAttrTypes)
	}
	m := GridntpsettingNtpKeysModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridntpsettingNtpKeysAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridntpsettingNtpKeysModel) Flatten(ctx context.Context, from *grid.GridntpsettingNtpKeys, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridntpsettingNtpKeysModel{}
	}
	m.Number = flex.FlattenInt64Pointer(from.Number)
	m.String = flex.FlattenStringPointer(from.String)
	m.Type = flex.FlattenStringPointer(from.Type)
}
