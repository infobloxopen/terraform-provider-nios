package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type NsgroupGridPrimaryModel struct {
	Name                     types.String `tfsdk:"name"`
	Stealth                  types.Bool   `tfsdk:"stealth"`
	GridReplicate            types.Bool   `tfsdk:"grid_replicate"`
	Lead                     types.Bool   `tfsdk:"lead"`
	PreferredPrimaries       types.List   `tfsdk:"preferred_primaries"`
	EnablePreferredPrimaries types.Bool   `tfsdk:"enable_preferred_primaries"`
}

var NsgroupGridPrimaryAttrTypes = map[string]attr.Type{
	"name":                       types.StringType,
	"stealth":                    types.BoolType,
	"grid_replicate":             types.BoolType,
	"lead":                       types.BoolType,
	"preferred_primaries":        types.ListType{ElemType: types.ObjectType{AttrTypes: NsgroupgridprimaryPreferredPrimariesAttrTypes}},
	"enable_preferred_primaries": types.BoolType,
}

var NsgroupGridPrimaryResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The grid member name.",
	},
	"stealth": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This flag governs whether the specified Grid member is in stealth mode or not. If set to True, the member is in stealth mode. This flag is ignored if the struct is specified as part of a stub zone.",
	},
	"grid_replicate": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The flag represents DNS zone transfers if set to False, and ID Grid Replication if set to True. This flag is ignored if the struct is specified as part of a stub zone or if it is set as grid_member in an authoritative zone.",
	},
	"lead": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This flag controls whether the Grid lead secondary server performs zone transfers to non lead secondaries. This flag is ignored if the struct is specified as grid_member in an authoritative zone.",
	},
	"preferred_primaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NsgroupgridprimaryPreferredPrimariesResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The primary preference list with Grid member names and\\or External Server extserver structs for this member.",
	},
	"enable_preferred_primaries": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This flag represents whether the preferred_primaries field values of this member are used.",
	},
}

func ExpandNsgroupGridPrimary(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.NsgroupGridPrimary {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NsgroupGridPrimaryModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NsgroupGridPrimaryModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.NsgroupGridPrimary {
	if m == nil {
		return nil
	}
	to := &dns.NsgroupGridPrimary{
		Name:                     flex.ExpandStringPointer(m.Name),
		Stealth:                  flex.ExpandBoolPointer(m.Stealth),
		GridReplicate:            flex.ExpandBoolPointer(m.GridReplicate),
		Lead:                     flex.ExpandBoolPointer(m.Lead),
		PreferredPrimaries:       flex.ExpandFrameworkListNestedBlock(ctx, m.PreferredPrimaries, diags, ExpandNsgroupgridprimaryPreferredPrimaries),
		EnablePreferredPrimaries: flex.ExpandBoolPointer(m.EnablePreferredPrimaries),
	}
	return to
}

func FlattenNsgroupGridPrimary(ctx context.Context, from *dns.NsgroupGridPrimary, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NsgroupGridPrimaryAttrTypes)
	}
	m := NsgroupGridPrimaryModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NsgroupGridPrimaryAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NsgroupGridPrimaryModel) Flatten(ctx context.Context, from *dns.NsgroupGridPrimary, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NsgroupGridPrimaryModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Stealth = types.BoolPointerValue(from.Stealth)
	m.GridReplicate = types.BoolPointerValue(from.GridReplicate)
	m.Lead = types.BoolPointerValue(from.Lead)
	m.PreferredPrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.PreferredPrimaries, NsgroupgridprimaryPreferredPrimariesAttrTypes, diags, FlattenNsgroupgridprimaryPreferredPrimaries)
	m.EnablePreferredPrimaries = types.BoolPointerValue(from.EnablePreferredPrimaries)
}
