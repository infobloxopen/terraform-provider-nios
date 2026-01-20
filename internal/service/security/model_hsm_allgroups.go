package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type HsmAllgroupsModel struct {
	Ref    types.String `tfsdk:"ref"`
	Uuid   types.String `tfsdk:"uuid"`
	Groups types.List   `tfsdk:"groups"`
}

var HsmAllgroupsAttrTypes = map[string]attr.Type{
	"ref":    types.StringType,
	"uuid":   types.StringType,
	"groups": types.ListType{ElemType: types.StringType},
}

var HsmAllgroupsResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"groups": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of HSM groups configured on the appliance.",
	},
}

func ExpandHsmAllgroups(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.HsmAllgroups {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m HsmAllgroupsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *HsmAllgroupsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.HsmAllgroups {
	if m == nil {
		return nil
	}
	to := &security.HsmAllgroups{
		Ref:    flex.ExpandStringPointer(m.Ref),
		Uuid:   flex.ExpandStringPointer(m.Uuid),
		Groups: flex.ExpandFrameworkListString(ctx, m.Groups, diags),
	}
	return to
}

func FlattenHsmAllgroups(ctx context.Context, from *security.HsmAllgroups, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(HsmAllgroupsAttrTypes)
	}
	m := HsmAllgroupsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, HsmAllgroupsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *HsmAllgroupsModel) Flatten(ctx context.Context, from *security.HsmAllgroups, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = HsmAllgroupsModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Groups = flex.FlattenFrameworkListString(ctx, from.Groups, diags)
}
