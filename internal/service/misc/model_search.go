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

type SearchModel struct {
	Ref  types.String `tfsdk:"ref"`
	Uuid types.String `tfsdk:"uuid"`
}

var SearchAttrTypes = map[string]attr.Type{
	"ref":  types.StringType,
	"uuid": types.StringType,
}

var SearchResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The UUID of the object.",
	},
}

func ExpandSearch(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.Search {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SearchModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SearchModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.Search {
	if m == nil {
		return nil
	}
	to := &misc.Search{
		Ref:  flex.ExpandStringPointer(m.Ref),
		Uuid: flex.ExpandStringPointer(m.Uuid),
	}
	return to
}

func FlattenSearch(ctx context.Context, from *misc.Search, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SearchAttrTypes)
	}
	m := SearchModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SearchAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SearchModel) Flatten(ctx context.Context, from *misc.Search, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SearchModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
}
