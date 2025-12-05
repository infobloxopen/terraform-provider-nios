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

type DeletedObjectsModel struct {
	Ref        types.String `tfsdk:"ref"`
	ObjectType types.String `tfsdk:"object_type"`
}

var DeletedObjectsAttrTypes = map[string]attr.Type{
	"ref":         types.StringType,
	"object_type": types.StringType,
}

var DeletedObjectsResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"object_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The object type of the deleted object. This is undefined if the object is not supported.",
	},
}

func ExpandDeletedObjects(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.DeletedObjects {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DeletedObjectsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DeletedObjectsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.DeletedObjects {
	if m == nil {
		return nil
	}
	to := &misc.DeletedObjects{
		Ref: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenDeletedObjects(ctx context.Context, from *misc.DeletedObjects, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DeletedObjectsAttrTypes)
	}
	m := DeletedObjectsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DeletedObjectsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DeletedObjectsModel) Flatten(ctx context.Context, from *misc.DeletedObjects, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DeletedObjectsModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.ObjectType = flex.FlattenStringPointer(from.ObjectType)
}
