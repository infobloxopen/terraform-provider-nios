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

type DbsnapshotModel struct {
	Ref       types.String `tfsdk:"ref"`
	Comment   types.String `tfsdk:"comment"`
	Timestamp types.Int64  `tfsdk:"timestamp"`
}

var DbsnapshotAttrTypes = map[string]attr.Type{
	"ref":       types.StringType,
	"comment":   types.StringType,
	"timestamp": types.Int64Type,
}

var DbsnapshotResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The descriptive comment.",
	},
	"timestamp": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The time when the latest OneDB snapshot was taken in Epoch seconds format.",
	},
}

func ExpandDbsnapshot(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.Dbsnapshot {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DbsnapshotModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DbsnapshotModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.Dbsnapshot {
	if m == nil {
		return nil
	}
	to := &misc.Dbsnapshot{
		Ref: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenDbsnapshot(ctx context.Context, from *misc.Dbsnapshot, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DbsnapshotAttrTypes)
	}
	m := DbsnapshotModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DbsnapshotAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DbsnapshotModel) Flatten(ctx context.Context, from *misc.Dbsnapshot, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DbsnapshotModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Timestamp = flex.FlattenInt64Pointer(from.Timestamp)
}
