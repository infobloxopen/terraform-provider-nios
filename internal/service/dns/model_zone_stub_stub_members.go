package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	"github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/plancontrol"
)

// ZoneStubStubMembersModel defines the model for members in a stub zone.
// All attributes of this model except 'name' are Computed, as input values for stub members are ignored in stub zones.
type ZoneStubStubMembersModel struct {
	Name                     types.String `tfsdk:"name"`
	Stealth                  types.Bool   `tfsdk:"stealth"`
	GridReplicate            types.Bool   `tfsdk:"grid_replicate"`
	Lead                     types.Bool   `tfsdk:"lead"`
	PreferredPrimaries       types.List   `tfsdk:"preferred_primaries"`
	EnablePreferredPrimaries types.Bool   `tfsdk:"enable_preferred_primaries"`
}

var ZoneStubStubMembersAttrTypes = map[string]attr.Type{
	"name":                       types.StringType,
	"stealth":                    types.BoolType,
	"grid_replicate":             types.BoolType,
	"lead":                       types.BoolType,
	"preferred_primaries":        types.ListType{ElemType: types.ObjectType{AttrTypes: ZonestubstubmembersPreferredPrimariesAttrTypes}},
	"enable_preferred_primaries": types.BoolType,
}

var ZoneStubStubMembersResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The grid member name.",
	},
	"stealth": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "This flag governs whether the specified Grid member is in stealth mode or not. If set to True, the member is in stealth mode. This flag is ignored if the struct is specified as part of a stub zone.",
		PlanModifiers: []planmodifier.Bool{
			plancontrol.UseStateForUnknownBool(),
		},
	},
	"grid_replicate": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The flag represents DNS zone transfers if set to False, and ID Grid Replication if set to True. This flag is ignored if the struct is specified as part of a stub zone or if it is set as grid_member in an authoritative zone.",
	},
	"lead": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "This flag controls whether the Grid lead secondary server performs zone transfers to non lead secondaries. This flag is ignored if the struct is specified as grid_member in an authoritative zone.",
		PlanModifiers: []planmodifier.Bool{
			plancontrol.UseStateForUnknownBool(),
		},
	},
	"preferred_primaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZonestubstubmembersPreferredPrimariesResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: "The primary preference list with Grid member names and\\or External Server extserver structs for this member.",
		PlanModifiers: []planmodifier.List{
			plancontrol.UseStateForUnknownList(),
		},
	},
	"enable_preferred_primaries": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "This flag represents whether the preferred_primaries field values of this member are used.",
		PlanModifiers: []planmodifier.Bool{
			plancontrol.UseStateForUnknownBool(),
		},
	},
}

func ExpandZoneStubStubMembers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneStubStubMembers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneStubStubMembersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneStubStubMembersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZoneStubStubMembers {
	if m == nil {
		return nil
	}
	to := &dns.ZoneStubStubMembers{
		Name: flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenZoneStubStubMembers(ctx context.Context, from *dns.ZoneStubStubMembers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneStubStubMembersAttrTypes)
	}
	m := ZoneStubStubMembersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneStubStubMembersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneStubStubMembersModel) Flatten(ctx context.Context, from *dns.ZoneStubStubMembers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneStubStubMembersModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Stealth = types.BoolPointerValue(from.Stealth)
	m.GridReplicate = types.BoolPointerValue(from.GridReplicate)
	m.Lead = types.BoolPointerValue(from.Lead)
	m.PreferredPrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.PreferredPrimaries, ZonestubstubmembersPreferredPrimariesAttrTypes, diags, FlattenZonestubstubmembersPreferredPrimaries)
	m.EnablePreferredPrimaries = types.BoolPointerValue(from.EnablePreferredPrimaries)
}
