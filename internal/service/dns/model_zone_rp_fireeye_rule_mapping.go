package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ZoneRpFireeyeRuleMappingModel struct {
	AptOverride           types.String `tfsdk:"apt_override"`
	FireeyeAlertMapping   types.List   `tfsdk:"fireeye_alert_mapping"`
	SubstitutedDomainName types.String `tfsdk:"substituted_domain_name"`
}

var ZoneRpFireeyeRuleMappingAttrTypes = map[string]attr.Type{
	"apt_override":            types.StringType,
	"fireeye_alert_mapping":   types.ListType{ElemType: types.ObjectType{AttrTypes: ZonerpfireeyerulemappingFireeyeAlertMappingAttrTypes}},
	"substituted_domain_name": types.StringType,
}

var ZoneRpFireeyeRuleMappingResourceSchemaAttributes = map[string]schema.Attribute{
	"apt_override": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The override setting for APT alerts.",
	},
	"fireeye_alert_mapping": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ZonerpfireeyerulemappingFireeyeAlertMappingResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The FireEye alert mapping.",
	},
	"substituted_domain_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The domain name to be substituted, this is applicable only when apt_override is set to \"SUBSTITUTE\".",
	},
}

func ExpandZoneRpFireeyeRuleMapping(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneRpFireeyeRuleMapping {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneRpFireeyeRuleMappingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneRpFireeyeRuleMappingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZoneRpFireeyeRuleMapping {
	if m == nil {
		return nil
	}
	to := &dns.ZoneRpFireeyeRuleMapping{
		AptOverride:           flex.ExpandStringPointer(m.AptOverride),
		FireeyeAlertMapping:   flex.ExpandFrameworkListNestedBlock(ctx, m.FireeyeAlertMapping, diags, ExpandZonerpfireeyerulemappingFireeyeAlertMapping),
		SubstitutedDomainName: flex.ExpandStringPointer(m.SubstitutedDomainName),
	}
	return to
}

func FlattenZoneRpFireeyeRuleMapping(ctx context.Context, from *dns.ZoneRpFireeyeRuleMapping, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneRpFireeyeRuleMappingAttrTypes)
	}
	m := ZoneRpFireeyeRuleMappingModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrsAll = types.MapNull(types.StringType)
	t, d := types.ObjectValueFrom(ctx, ZoneRpFireeyeRuleMappingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneRpFireeyeRuleMappingModel) Flatten(ctx context.Context, from *dns.ZoneRpFireeyeRuleMapping, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneRpFireeyeRuleMappingModel{}
	}
	m.AptOverride = flex.FlattenStringPointer(from.AptOverride)
	m.FireeyeAlertMapping = flex.FlattenFrameworkListNestedBlock(ctx, from.FireeyeAlertMapping, ZonerpfireeyerulemappingFireeyeAlertMappingAttrTypes, diags, FlattenZonerpfireeyerulemappingFireeyeAlertMapping)
	m.SubstitutedDomainName = flex.FlattenStringPointer(from.SubstitutedDomainName)
}
