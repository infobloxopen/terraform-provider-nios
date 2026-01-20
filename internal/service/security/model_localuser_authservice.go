package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type LocaluserAuthserviceModel struct {
	Ref      types.String `tfsdk:"ref"`
	Uuid     types.String `tfsdk:"uuid"`
	Comment  types.String `tfsdk:"comment"`
	Disabled types.Bool   `tfsdk:"disabled"`
	Name     types.String `tfsdk:"name"`
}

var LocaluserAuthserviceAttrTypes = map[string]attr.Type{
	"ref":      types.StringType,
	"uuid":     types.StringType,
	"comment":  types.StringType,
	"disabled": types.BoolType,
	"name":     types.StringType,
}

var LocaluserAuthserviceResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"comment": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The local user authentication service comment.",
	},
	"disabled": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Flag that indicates whether the local user authentication service is enabled or not.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the local user authentication service.",
	},
}

func ExpandLocaluserAuthservice(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.LocaluserAuthservice {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m LocaluserAuthserviceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *LocaluserAuthserviceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.LocaluserAuthservice {
	if m == nil {
		return nil
	}
	to := &security.LocaluserAuthservice{
		Ref:  flex.ExpandStringPointer(m.Ref),
		Uuid: flex.ExpandStringPointer(m.Uuid),
	}
	return to
}

func FlattenLocaluserAuthservice(ctx context.Context, from *security.LocaluserAuthservice, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(LocaluserAuthserviceAttrTypes)
	}
	m := LocaluserAuthserviceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, LocaluserAuthserviceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *LocaluserAuthserviceModel) Flatten(ctx context.Context, from *security.LocaluserAuthservice, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = LocaluserAuthserviceModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disabled = types.BoolPointerValue(from.Disabled)
	m.Name = flex.FlattenStringPointer(from.Name)
}
