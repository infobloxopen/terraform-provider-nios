package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type FileopModel struct {
	Ref types.String `tfsdk:"ref"`
}

var FileopAttrTypes = map[string]attr.Type{
	"ref": types.StringType,
}

var FileopResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
}

func ExpandFileop(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.Fileop {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FileopModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *FileopModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.Fileop {
	if m == nil {
		return nil
	}
	to := &misc.Fileop{
		Ref: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenFileop(ctx context.Context, from *misc.Fileop, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FileopAttrTypes)
	}
	m := FileopModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, FileopAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FileopModel) Flatten(ctx context.Context, from *misc.Fileop, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FileopModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
}
