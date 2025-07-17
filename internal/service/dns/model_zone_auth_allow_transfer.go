package dns

import (
	"context"
	"regexp"

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

type ZoneAuthAllowTransferModel struct {
	Address        types.String `tfsdk:"address"`
	Struct         types.String `tfsdk:"struct"`
	Permission     types.String `tfsdk:"permission"`
	TsigKey        types.String `tfsdk:"tsig_key"`
	TsigKeyAlg     types.String `tfsdk:"tsig_key_alg"`
	TsigKeyName    types.String `tfsdk:"tsig_key_name"`
	UseTsigKeyName types.Bool   `tfsdk:"use_tsig_key_name"`
}

var ZoneAuthAllowTransferAttrTypes = map[string]attr.Type{
	"address":           types.StringType,
	"struct":            types.StringType,
	"permission":        types.StringType,
	"tsig_key":          types.StringType,
	"tsig_key_alg":      types.StringType,
	"tsig_key_name":     types.StringType,
	"use_tsig_key_name": types.BoolType,
}

var ZoneAuthAllowTransferResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The address this rule applies to or \"Any\".",
	},
	"struct": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The struct type of the object.",
		Validators: []validator.String{
			stringvalidator.OneOf("addressac", "tsigac"),
		},
	},
	"permission": schema.StringAttribute{
		Optional: true,
		Computed: true,
		// Default:  stringdefault.StaticString("ALLOW"),
		// Validators: []validator.String{
		// 	stringvalidator.OneOf("ALLOW", "DENY"),
		// },
		MarkdownDescription: "The permission to use for this address.",
	},
	"tsig_key": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Address should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "A generated TSIG key. If the external primary server is a NIOS appliance running DNS One 2.x code, this can be set to :2xCOMPAT.",
	},
	"tsig_key_alg": schema.StringAttribute{
		Optional: true,
		Computed: true,
		// Default:  stringdefault.StaticString("HMAC-MD5"),
		// Validators: []validator.String{
		// 	stringvalidator.OneOf("HMAC-MD5", "HMAC-SHA256"),
		// 	stringvalidator.ConflictsWith(path.MatchRoot("address"), path.MatchRoot("permission")),
		// },
		MarkdownDescription: "The TSIG key algorithm.",
	},
	"tsig_key_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		// Validators: []validator.String{
		// 	stringvalidator.RegexMatches(
		// 		regexp.MustCompile(`^[^\s].*[^\s]$`),
		// 		"Address should not have leading or trailing whitespace",
		// 	),
		// 	stringvalidator.AlsoRequires(path.MatchRoot("use_tsig_key_name")),
		// },
		MarkdownDescription: "The name of the TSIG key. If 2.x TSIG compatibility is used, this is set to 'tsig_xfer' on retrieval, and ignored on insert or update.",
	},
	"use_tsig_key_name": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		// Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: tsig_key_name",
	},
}

func ExpandZoneAuthAllowTransfer(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneAuthAllowTransfer {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneAuthAllowTransferModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneAuthAllowTransferModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZoneAuthAllowTransfer {
	if m == nil {
		return nil
	}
	structValue := m.Struct.ValueString()
	to := &dns.ZoneAuthAllowTransfer{
		Struct: flex.ExpandStringPointer(m.Struct),
	}
	if structValue == "addressac" {
		to.Address = flex.ExpandStringPointer(m.Address)
		to.Permission = flex.ExpandStringPointer(m.Permission)
	} else if structValue == "tsigac" {
		to.TsigKey = flex.ExpandStringPointer(m.TsigKey)
		to.TsigKeyAlg = flex.ExpandStringPointer(m.TsigKeyAlg)
		to.TsigKeyName = flex.ExpandStringPointer(m.TsigKeyName)
	}
	return to
}

func FlattenZoneAuthAllowTransfer(ctx context.Context, from *dns.ZoneAuthAllowTransfer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneAuthAllowTransferAttrTypes)
	}
	m := ZoneAuthAllowTransferModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneAuthAllowTransferAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneAuthAllowTransferModel) Flatten(ctx context.Context, from *dns.ZoneAuthAllowTransfer, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneAuthAllowTransferModel{}
	}

	m.Struct = flex.FlattenStringPointer(from.Struct)

	if m.Struct.ValueString() == "addressac" {
		m.Address = flex.FlattenStringPointer(from.Address)
		m.Permission = flex.FlattenStringPointer(from.Permission)
	} else if m.Struct.ValueString() == "tsigac" {
		m.TsigKey = flex.FlattenStringPointer(from.TsigKey)
		m.TsigKeyAlg = flex.FlattenStringPointer(from.TsigKeyAlg)
		m.TsigKeyName = flex.FlattenStringPointer(from.TsigKeyName)
		m.UseTsigKeyName = types.BoolPointerValue(from.UseTsigKeyName)
	}
}
