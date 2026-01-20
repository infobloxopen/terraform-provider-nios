package parentalcontrol

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/parentalcontrol"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ParentalcontrolBlockingpolicyModel struct {
	Ref   types.String `tfsdk:"ref"`
	Uuid  types.String `tfsdk:"uuid"`
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

var ParentalcontrolBlockingpolicyAttrTypes = map[string]attr.Type{
	"ref":   types.StringType,
	"uuid":  types.StringType,
	"name":  types.StringType,
	"value": types.StringType,
}

var ParentalcontrolBlockingpolicyResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the blocking policy.",
	},
	"value": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The 32 bit hex value of the blocking policy.",
	},
}

func ExpandParentalcontrolBlockingpolicy(ctx context.Context, o types.Object, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolBlockingpolicy {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ParentalcontrolBlockingpolicyModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ParentalcontrolBlockingpolicyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *parentalcontrol.ParentalcontrolBlockingpolicy {
	if m == nil {
		return nil
	}
	to := &parentalcontrol.ParentalcontrolBlockingpolicy{
		Ref:   flex.ExpandStringPointer(m.Ref),
		Name:  flex.ExpandStringPointer(m.Name),
		Value: flex.ExpandStringPointer(m.Value),
	}
	return to
}

func FlattenParentalcontrolBlockingpolicy(ctx context.Context, from *parentalcontrol.ParentalcontrolBlockingpolicy, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ParentalcontrolBlockingpolicyAttrTypes)
	}
	m := ParentalcontrolBlockingpolicyModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ParentalcontrolBlockingpolicyAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ParentalcontrolBlockingpolicyModel) Flatten(ctx context.Context, from *parentalcontrol.ParentalcontrolBlockingpolicy, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ParentalcontrolBlockingpolicyModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Value = flex.FlattenStringPointer(from.Value)
}
