package dns

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ZoneAuthMsAllowTransferModel struct {
	Address    iptypes.IPAddress `tfsdk:"address"`
	Permission types.String      `tfsdk:"permission"`
}

var ZoneAuthMsAllowTransferAttrTypes = map[string]attr.Type{
	"address":    iptypes.IPAddressType{},
	"permission": types.StringType,
}

var ZoneAuthMsAllowTransferResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		CustomType:          iptypes.IPAddressType{},
		Required:            true,
		MarkdownDescription: "The address this rule applies to or \"Any\".",
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
	},
	"permission": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("ALLOW"),
		Validators: []validator.String{
			stringvalidator.OneOf("ALLOW", "DENY"),
		},
		MarkdownDescription: "The permission to use for this address.",
	},
}

func ExpandZoneAuthMsAllowTransfer(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneAuthMsAllowTransfer {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneAuthMsAllowTransferModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneAuthMsAllowTransferModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZoneAuthMsAllowTransfer {
	if m == nil {
		return nil
	}
	to := &dns.ZoneAuthMsAllowTransfer{
		Address:    flex.ExpandIPAddress(m.Address),
		Permission: flex.ExpandStringPointer(m.Permission),
	}
	return to
}

func FlattenZoneAuthMsAllowTransfer(ctx context.Context, from *dns.ZoneAuthMsAllowTransfer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneAuthMsAllowTransferAttrTypes)
	}
	m := ZoneAuthMsAllowTransferModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneAuthMsAllowTransferAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneAuthMsAllowTransferModel) Flatten(ctx context.Context, from *dns.ZoneAuthMsAllowTransfer, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneAuthMsAllowTransferModel{}
	}
	m.Address = flex.FlattenIPAddress(from.Address)
	m.Permission = flex.FlattenStringPointer(from.Permission)
}
