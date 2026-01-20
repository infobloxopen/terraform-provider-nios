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

type DatacollectionclusterModel struct {
	Ref                types.String `tfsdk:"ref"`
    Uuid        types.String `tfsdk:"uuid"`
	EnableRegistration types.Bool   `tfsdk:"enable_registration"`
	Name               types.String `tfsdk:"name"`
}

var DatacollectionclusterAttrTypes = map[string]attr.Type{
	"ref":                 types.StringType,
    "uuid":        types.StringType,
	"enable_registration": types.BoolType,
	"name":                types.StringType,
}

var DatacollectionclusterResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
    "uuid": schema.StringAttribute{
        Computed:            true,
        MarkdownDescription: "The uuid to the object.",
    },
	"enable_registration": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable/disable new registration requests",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Display name for cluster",
	},
}

func ExpandDatacollectioncluster(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.Datacollectioncluster {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DatacollectionclusterModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DatacollectionclusterModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.Datacollectioncluster {
	if m == nil {
		return nil
	}
	to := &misc.Datacollectioncluster{
		Ref:                flex.ExpandStringPointer(m.Ref),
		EnableRegistration: flex.ExpandBoolPointer(m.EnableRegistration),
	}
	return to
}

func FlattenDatacollectioncluster(ctx context.Context, from *misc.Datacollectioncluster, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DatacollectionclusterAttrTypes)
	}
	m := DatacollectionclusterModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DatacollectionclusterAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DatacollectionclusterModel) Flatten(ctx context.Context, from *misc.Datacollectioncluster, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DatacollectionclusterModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
    m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.EnableRegistration = types.BoolPointerValue(from.EnableRegistration)
	m.Name = flex.FlattenStringPointer(from.Name)
}
