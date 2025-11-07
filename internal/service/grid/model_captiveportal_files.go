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

type CaptiveportalFilesModel struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

var CaptiveportalFilesAttrTypes = map[string]attr.Type{
	"name": types.StringType,
	"type": types.StringType,
}

var CaptiveportalFilesResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the uploaded file.",
	},
	"type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The type of the uploaded file.",
	},
}

func ExpandCaptiveportalFiles(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.CaptiveportalFiles {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m CaptiveportalFilesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *CaptiveportalFilesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.CaptiveportalFiles {
	if m == nil {
		return nil
	}
	to := &grid.CaptiveportalFiles{
		Name: flex.ExpandStringPointer(m.Name),
		Type: flex.ExpandStringPointer(m.Type),
	}
	return to
}

func FlattenCaptiveportalFiles(ctx context.Context, from *grid.CaptiveportalFiles, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(CaptiveportalFilesAttrTypes)
	}
	m := CaptiveportalFilesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, CaptiveportalFilesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *CaptiveportalFilesModel) Flatten(ctx context.Context, from *grid.CaptiveportalFiles, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = CaptiveportalFilesModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Type = flex.FlattenStringPointer(from.Type)
}
