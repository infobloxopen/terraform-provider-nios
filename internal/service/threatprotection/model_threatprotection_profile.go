package threatprotection

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/threatprotection"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/import"
)

type ThreatprotectionProfileModel struct {
	Ref                             types.String `tfsdk:"ref"`
	Comment                         types.String `tfsdk:"comment"`
	CurrentRuleset                  types.String `tfsdk:"current_ruleset"`
	DisableMultipleDnsTcpRequest    types.Bool   `tfsdk:"disable_multiple_dns_tcp_request"`
	EventsPerSecondPerRule          types.Int64  `tfsdk:"events_per_second_per_rule"`
	ExtAttrs                        types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll                     types.Map    `tfsdk:"extattrs_all"`
	Members                         types.List   `tfsdk:"members"`
	Name                            types.String `tfsdk:"name"`
	SourceMember                    types.String `tfsdk:"source_member"`
	SourceProfile                   types.String `tfsdk:"source_profile"`
	UseCurrentRuleset               types.Bool   `tfsdk:"use_current_ruleset"`
	UseDisableMultipleDnsTcpRequest types.Bool   `tfsdk:"use_disable_multiple_dns_tcp_request"`
	UseEventsPerSecondPerRule       types.Bool   `tfsdk:"use_events_per_second_per_rule"`
}

var ThreatprotectionProfileAttrTypes = map[string]attr.Type{
	"ref":                                  types.StringType,
	"comment":                              types.StringType,
	"current_ruleset":                      types.StringType,
	"disable_multiple_dns_tcp_request":     types.BoolType,
	"events_per_second_per_rule":           types.Int64Type,
	"extattrs":                             types.MapType{ElemType: types.StringType},
	"extattrs_all":                         types.MapType{ElemType: types.StringType},
	"members":                              types.ListType{ElemType: types.StringType},
	"name":                                 types.StringType,
	"source_member":                        types.StringType,
	"source_profile":                       types.StringType,
	"use_current_ruleset":                  types.BoolType,
	"use_disable_multiple_dns_tcp_request": types.BoolType,
	"use_events_per_second_per_rule":       types.BoolType,
}

var ThreatprotectionProfileResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The comment for the Threat Protection profile.",
	},
	"current_ruleset": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The current Threat Protection profile ruleset.",
	},
	"disable_multiple_dns_tcp_request": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if multiple BIND responses via TCP connection are disabled.",
	},
	"events_per_second_per_rule": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of events logged per second per rule.",
	},
	"extattrs": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"extattrs_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"members": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of members that are associated with the profile.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the Threat Protection profile.",
	},
	"source_member": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The source member. It can be used only during the create operation for cloning a profile from an existing member.",
	},
	"source_profile": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The source profile. It can be used only during the create operation for cloning a profile from an existing profile.",
	},
	"use_current_ruleset": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: current_ruleset",
	},
	"use_disable_multiple_dns_tcp_request": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: disable_multiple_dns_tcp_request",
	},
	"use_events_per_second_per_rule": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: events_per_second_per_rule",
	},
}

func ExpandThreatprotectionProfile(ctx context.Context, o types.Object, diags *diag.Diagnostics) *threatprotection.ThreatprotectionProfile {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ThreatprotectionProfileModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ThreatprotectionProfileModel) Expand(ctx context.Context, diags *diag.Diagnostics) *threatprotection.ThreatprotectionProfile {
	if m == nil {
		return nil
	}
	to := &threatprotection.ThreatprotectionProfile{
		Ref:                             flex.ExpandStringPointer(m.Ref),
		Comment:                         flex.ExpandStringPointer(m.Comment),
		CurrentRuleset:                  flex.ExpandStringPointer(m.CurrentRuleset),
		DisableMultipleDnsTcpRequest:    flex.ExpandBoolPointer(m.DisableMultipleDnsTcpRequest),
		EventsPerSecondPerRule:          flex.ExpandInt64Pointer(m.EventsPerSecondPerRule),
		ExtAttrs:                        ExpandExtAttrs(ctx, m.ExtAttrs, diags),
		Members:                         flex.ExpandFrameworkListString(ctx, m.Members, diags),
		Name:                            flex.ExpandStringPointer(m.Name),
		SourceMember:                    flex.ExpandStringPointer(m.SourceMember),
		SourceProfile:                   flex.ExpandStringPointer(m.SourceProfile),
		UseCurrentRuleset:               flex.ExpandBoolPointer(m.UseCurrentRuleset),
		UseDisableMultipleDnsTcpRequest: flex.ExpandBoolPointer(m.UseDisableMultipleDnsTcpRequest),
		UseEventsPerSecondPerRule:       flex.ExpandBoolPointer(m.UseEventsPerSecondPerRule),
	}
	return to
}

func FlattenThreatprotectionProfile(ctx context.Context, from *threatprotection.ThreatprotectionProfile, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ThreatprotectionProfileAttrTypes)
	}
	m := ThreatprotectionProfileModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, ThreatprotectionProfileAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ThreatprotectionProfileModel) Flatten(ctx context.Context, from *threatprotection.ThreatprotectionProfile, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ThreatprotectionProfileModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CurrentRuleset = flex.FlattenStringPointer(from.CurrentRuleset)
	m.DisableMultipleDnsTcpRequest = types.BoolPointerValue(from.DisableMultipleDnsTcpRequest)
	m.EventsPerSecondPerRule = flex.FlattenInt64Pointer(from.EventsPerSecondPerRule)
	m.ExtAttrs = FlattenExtAttrs(ctx, m.ExtAttrs, from.ExtAttrs, diags)
	m.Members = flex.FlattenFrameworkListString(ctx, from.Members, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.SourceMember = flex.FlattenStringPointer(from.SourceMember)
	m.SourceProfile = flex.FlattenStringPointer(from.SourceProfile)
	m.UseCurrentRuleset = types.BoolPointerValue(from.UseCurrentRuleset)
	m.UseDisableMultipleDnsTcpRequest = types.BoolPointerValue(from.UseDisableMultipleDnsTcpRequest)
	m.UseEventsPerSecondPerRule = types.BoolPointerValue(from.UseEventsPerSecondPerRule)
}
