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

type ZoneAuthDnssecKeysModel struct {
	Tag           types.Int64  `tfsdk:"tag"`
	Status        types.String `tfsdk:"status"`
	NextEventDate types.Int64  `tfsdk:"next_event_date"`
	Type          types.String `tfsdk:"type"`
	Algorithm     types.String `tfsdk:"algorithm"`
	PublicKey     types.String `tfsdk:"public_key"`
}

var ZoneAuthDnssecKeysAttrTypes = map[string]attr.Type{
	"tag":             types.Int64Type,
	"status":          types.StringType,
	"next_event_date": types.Int64Type,
	"type":            types.StringType,
	"algorithm":       types.StringType,
	"public_key":      types.StringType,
}

var ZoneAuthDnssecKeysResourceSchemaAttributes = map[string]schema.Attribute{
	"tag": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The tag of the key for the zone.",
	},
	"status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The status of the key for the zone.",
	},
	"next_event_date": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The next event date for the key, the rollover date for an active key or the removal date for an already rolled one.",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The key type.",
	},
	"algorithm": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The public-key encryption algorithm. Values 1, 3 and 6 are deprecated from NIOS 9.0.",
	},
	"public_key": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Base-64 encoding of the public key.",
	},
}

func ExpandZoneAuthDnssecKeys(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneAuthDnssecKeys {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneAuthDnssecKeysModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneAuthDnssecKeysModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZoneAuthDnssecKeys {
	if m == nil {
		return nil
	}
	to := &dns.ZoneAuthDnssecKeys{
		Tag: flex.ExpandInt64Pointer(m.Tag),
	}
	return to
}

func FlattenZoneAuthDnssecKeys(ctx context.Context, from *dns.ZoneAuthDnssecKeys, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneAuthDnssecKeysAttrTypes)
	}
	m := ZoneAuthDnssecKeysModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneAuthDnssecKeysAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneAuthDnssecKeysModel) Flatten(ctx context.Context, from *dns.ZoneAuthDnssecKeys, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneAuthDnssecKeysModel{}
	}
	m.Tag = flex.FlattenInt64Pointer(from.Tag)
	m.Status = flex.FlattenStringPointer(from.Status)
	m.NextEventDate = flex.FlattenInt64Pointer(from.NextEventDate)
	m.Type = flex.FlattenStringPointer(from.Type)
	m.Algorithm = flex.FlattenStringPointer(from.Algorithm)
	m.PublicKey = flex.FlattenStringPointer(from.PublicKey)
}
