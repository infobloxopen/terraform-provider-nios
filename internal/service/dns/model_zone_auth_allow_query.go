package dns

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ZoneAuthAllowQueryModel struct {
	Address        iptypes.IPAddress `tfsdk:"address"`
	Struct         types.String      `tfsdk:"struct"`
	Permission     types.String      `tfsdk:"permission"`
	TsigKey        types.String      `tfsdk:"tsig_key"`
	TsigKeyAlg     types.String      `tfsdk:"tsig_key_alg"`
	TsigKeyName    types.String      `tfsdk:"tsig_key_name"`
	UseTsigKeyName types.Bool        `tfsdk:"use_tsig_key_name"`
}

var ZoneAuthAllowQueryAttrTypes = map[string]attr.Type{
	"address":           iptypes.IPAddressType{},
	"struct":            types.StringType,
	"permission":        types.StringType,
	"tsig_key":          types.StringType,
	"tsig_key_alg":      types.StringType,
	"tsig_key_name":     types.StringType,
	"use_tsig_key_name": types.BoolType,
}

var ZoneAuthAllowQueryResourceSchemaAttributes = map[string]schema.Attribute{
	"struct": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("addressac", "tsigac"),
		},
		MarkdownDescription: "The struct type of the object. The value must be one of 'addressac' and 'tsigac'.",
	},
	"address": schema.StringAttribute{
		CustomType:          iptypes.IPAddressType{},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The address this rule applies to or \"Any\".",
		Validators: []validator.String{
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("tsig_key"),
				path.MatchRelative().AtParent().AtName("tsig_key_alg"),
				path.MatchRelative().AtParent().AtName("use_tsig_key_name"),
			),
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
	},
	"permission": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The permission to use for this address.",
		Validators: []validator.String{
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("tsig_key"),
				path.MatchRelative().AtParent().AtName("tsig_key_alg"),
				path.MatchRelative().AtParent().AtName("use_tsig_key_name"),
			),
			stringvalidator.OneOf("ALLOW", "DENY"),
		},
		Default: stringdefault.StaticString("ALLOW"),
	},
	"tsig_key": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "A generated TSIG key. If the external primary server is a NIOS appliance running DNS One 2.x code, this can be set to :2xCOMPAT.",
		Validators: []validator.String{
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("address"),
				path.MatchRelative().AtParent().AtName("permission"),
			),
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
	},
	"tsig_key_alg": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The TSIG key algorithm.",
		Validators: []validator.String{
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("address"),
				path.MatchRelative().AtParent().AtName("permission"),
			),
		},
	},
	"tsig_key_name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The name of the TSIG key. If 2.x TSIG compatibility is used, this is set to 'tsig_xfer' on retrieval, and ignored on insert or update.",
		Validators: []validator.String{
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("address"),
				path.MatchRelative().AtParent().AtName("permission"),
			),
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s].*[^\s]$`),
				"Should not have leading or trailing whitespace",
			),
		},
	},
	"use_tsig_key_name": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Use flag for: tsig_key_name",
		Validators: []validator.Bool{
			boolvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("address"),
				path.MatchRelative().AtParent().AtName("permission"),
			),
		},
	},
}

func ExpandZoneAuthAllowQuery(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneAuthAllowQuery {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneAuthAllowQueryModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneAuthAllowQueryModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZoneAuthAllowQuery {
	if m == nil {
		return nil
	}
	to := &dns.ZoneAuthAllowQuery{
		Address:        flex.ExpandIPAddress(m.Address),
		Struct:         flex.ExpandStringPointer(m.Struct),
		Permission:     flex.ExpandStringPointer(m.Permission),
		TsigKey:        flex.ExpandStringPointer(m.TsigKey),
		TsigKeyAlg:     flex.ExpandStringPointer(m.TsigKeyAlg),
		TsigKeyName:    flex.ExpandStringPointer(m.TsigKeyName),
		UseTsigKeyName: flex.ExpandBoolPointer(m.UseTsigKeyName),
	}
	return to
}

func FlattenZoneAuthAllowQuery(ctx context.Context, from *dns.ZoneAuthAllowQuery, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneAuthAllowQueryAttrTypes)
	}
	m := ZoneAuthAllowQueryModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneAuthAllowQueryAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneAuthAllowQueryModel) Flatten(ctx context.Context, from *dns.ZoneAuthAllowQuery, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneAuthAllowQueryModel{}
	}
	m.Address = flex.FlattenIPAddress(from.Address)
	m.Struct = flex.FlattenStringPointer(from.Struct)
	m.Permission = flex.FlattenStringPointer(from.Permission)
	m.TsigKey = flex.FlattenStringPointer(from.TsigKey)
	m.TsigKeyAlg = flex.FlattenStringPointer(from.TsigKeyAlg)
	m.TsigKeyName = flex.FlattenStringPointer(from.TsigKeyName)
	m.UseTsigKeyName = types.BoolPointerValue(from.UseTsigKeyName)
}
