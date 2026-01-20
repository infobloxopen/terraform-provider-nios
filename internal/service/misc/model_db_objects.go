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

type DbObjectsModel struct {
	Ref            types.String `tfsdk:"ref"`
	Uuid           types.String `tfsdk:"uuid"`
	LastSequenceId types.String `tfsdk:"last_sequence_id"`
	Object         types.String `tfsdk:"object"`
	ObjectType     types.String `tfsdk:"object_type"`
	UniqueId       types.String `tfsdk:"unique_id"`
}

var DbObjectsAttrTypes = map[string]attr.Type{
	"ref":              types.StringType,
	"uuid":             types.StringType,	
	"last_sequence_id": types.StringType,
	"object":           types.StringType,
	"object_type":      types.StringType,
	"unique_id":        types.StringType,
}

var DbObjectsResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The UUID of the object.",
	},
	"last_sequence_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The last returned sequence ID.",
	},
	"object": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The record object when supported by WAPI. Otherwise, the value is \"None\".",
	},
	"object_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The object type. This is undefined if the object is not supported.",
	},
	"unique_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The unique ID of the requested object.",
	},
}

func ExpandDbObjects(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.DbObjects {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DbObjectsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DbObjectsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.DbObjects {
	if m == nil {
		return nil
	}
	to := &misc.DbObjects{
		Ref:  flex.ExpandStringPointer(m.Ref),
		Uuid: flex.ExpandStringPointer(m.Uuid),
	}
	return to
}

func FlattenDbObjects(ctx context.Context, from *misc.DbObjects, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DbObjectsAttrTypes)
	}
	m := DbObjectsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DbObjectsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DbObjectsModel) Flatten(ctx context.Context, from *misc.DbObjects, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DbObjectsModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.LastSequenceId = flex.FlattenStringPointer(from.LastSequenceId)
	m.Object = flex.FlattenStringPointer(from.Object)
	m.ObjectType = flex.FlattenStringPointer(from.ObjectType)
	m.UniqueId = flex.FlattenStringPointer(from.UniqueId)
}
