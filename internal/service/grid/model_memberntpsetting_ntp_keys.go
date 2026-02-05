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

type MemberntpsettingNtpKeysModel struct {
	Number types.Int64  `tfsdk:"number"`
	String types.String `tfsdk:"string"`
	Type   types.String `tfsdk:"type"`
}

var MemberntpsettingNtpKeysAttrTypes = map[string]attr.Type{
	"number": types.Int64Type,
	"string": types.StringType,
	"type":   types.StringType,
}

var MemberntpsettingNtpKeysResourceSchemaAttributes = map[string]schema.Attribute{
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

func ExpandMemberntpsettingNtpKeys(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberntpsettingNtpKeys {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberntpsettingNtpKeysModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberntpsettingNtpKeysModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberntpsettingNtpKeys {
	if m == nil {
		return nil
	}
	to := &grid.MemberntpsettingNtpKeys{
		Number: flex.ExpandInt64Pointer(m.Number),
		String: flex.ExpandStringPointer(m.String),
		Type:   flex.ExpandStringPointer(m.Type),
	}
	return to
}

func FlattenMemberntpsettingNtpKeys(ctx context.Context, from *grid.MemberntpsettingNtpKeys, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberntpsettingNtpKeysAttrTypes)
	}
	m := MemberntpsettingNtpKeysModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberntpsettingNtpKeysAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberntpsettingNtpKeysModel) Flatten(ctx context.Context, from *grid.MemberntpsettingNtpKeys, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberntpsettingNtpKeysModel{}
	}
	m.Number = flex.FlattenInt64Pointer(from.Number)
	m.String = flex.FlattenStringPointer(from.String)
	m.Type = flex.FlattenStringPointer(from.Type)
}
