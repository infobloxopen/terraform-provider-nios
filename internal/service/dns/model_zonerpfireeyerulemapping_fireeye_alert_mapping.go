package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ZonerpfireeyerulemappingFireeyeAlertMappingModel struct {
	AlertType types.String `tfsdk:"alert_type"`
	RpzRule   types.String `tfsdk:"rpz_rule"`
	Lifetime  types.Int64  `tfsdk:"lifetime"`
}

var ZonerpfireeyerulemappingFireeyeAlertMappingAttrTypes = map[string]attr.Type{
	"alert_type": types.StringType,
	"rpz_rule":   types.StringType,
	"lifetime":   types.Int64Type,
}

var ZonerpfireeyerulemappingFireeyeAlertMappingResourceSchemaAttributes = map[string]schema.Attribute{
	"alert_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("DOMAIN_MATCH", "INFECTION_MATCH", "MALWARE_CALLBACK", "MALWARE_OBJECT", "WEB_INFECTION"),
		},
		MarkdownDescription: "The type of Fireeye Alert.",
	},
	"rpz_rule": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("NODATA", "NOOVERRIDE", "NXDOMAIN", "PASSTHRU", "SUBSTITUTE"),
		},
		MarkdownDescription: "The RPZ rule for the alert.",
	},
	"lifetime": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The expiration Lifetime of alert type. The 32-bit unsigned integer represents the amount of seconds this alert type will live for. 0 means the alert will never expire.",
	},
}

func ExpandZonerpfireeyerulemappingFireeyeAlertMapping(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZonerpfireeyerulemappingFireeyeAlertMapping {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZonerpfireeyerulemappingFireeyeAlertMappingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZonerpfireeyerulemappingFireeyeAlertMappingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZonerpfireeyerulemappingFireeyeAlertMapping {
	if m == nil {
		return nil
	}
	to := &dns.ZonerpfireeyerulemappingFireeyeAlertMapping{
		AlertType: flex.ExpandStringPointer(m.AlertType),
		RpzRule:   flex.ExpandStringPointer(m.RpzRule),
		Lifetime:  flex.ExpandInt64Pointer(m.Lifetime),
	}
	return to
}

func FlattenZonerpfireeyerulemappingFireeyeAlertMapping(ctx context.Context, from *dns.ZonerpfireeyerulemappingFireeyeAlertMapping, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZonerpfireeyerulemappingFireeyeAlertMappingAttrTypes)
	}
	m := ZonerpfireeyerulemappingFireeyeAlertMappingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZonerpfireeyerulemappingFireeyeAlertMappingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZonerpfireeyerulemappingFireeyeAlertMappingModel) Flatten(ctx context.Context, from *dns.ZonerpfireeyerulemappingFireeyeAlertMapping, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZonerpfireeyerulemappingFireeyeAlertMappingModel{}
	}
	m.AlertType = flex.FlattenStringPointer(from.AlertType)
	m.RpzRule = flex.FlattenStringPointer(from.RpzRule)
	m.Lifetime = flex.FlattenInt64Pointer(from.Lifetime)
}
