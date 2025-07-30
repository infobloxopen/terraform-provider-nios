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
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ZoneAuthAllowUpdateModel struct {
	Address        iptypes.IPAddress `tfsdk:"address"`
	Struct         types.String      `tfsdk:"struct"`
	Permission     types.String      `tfsdk:"permission"`
	TsigKey        types.String      `tfsdk:"tsig_key"`
	TsigKeyAlg     types.String      `tfsdk:"tsig_key_alg"`
	TsigKeyName    types.String      `tfsdk:"tsig_key_name"`
	UseTsigKeyName types.Bool        `tfsdk:"use_tsig_key_name"`
}

var ZoneAuthAllowUpdateAttrTypes = map[string]attr.Type{
	"address":           iptypes.IPAddressType{},
	"struct":            types.StringType,
	"permission":        types.StringType,
	"tsig_key":          types.StringType,
	"tsig_key_alg":      types.StringType,
	"tsig_key_name":     types.StringType,
	"use_tsig_key_name": types.BoolType,
}

var ZoneAuthAllowUpdateResourceSchemaAttributes = map[string]schema.Attribute{
	"struct": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The struct type of the object. The value must be one of 'addressac' and 'tsigac'.",
		Validators: []validator.String{
			stringvalidator.OneOf("addressac", "tsigac"),
		},
	},
	"address": schema.StringAttribute{
		CustomType: iptypes.IPAddressType{},
		Optional:   true,
		Computed:   true,
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
		MarkdownDescription: "The address this rule applies to or \"Any\".",
	},
	"permission": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("tsig_key"),
				path.MatchRelative().AtParent().AtName("tsig_key_alg"),
				path.MatchRelative().AtParent().AtName("use_tsig_key_name"),
			),
		},
		MarkdownDescription: "The permission to use for this address.",
	},
	"tsig_key": schema.StringAttribute{
		Optional: true,
		Computed: true,
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
		MarkdownDescription: "A generated TSIG key. If the external primary server is a NIOS appliance running DNS One 2.x code, this can be set to :2xCOMPAT.",
	},
	"tsig_key_alg": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("address"),
				path.MatchRelative().AtParent().AtName("permission"),
			),
		},
		MarkdownDescription: "The TSIG key algorithm.",
	},
	"tsig_key_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
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
		MarkdownDescription: "The name of the TSIG key. If 2.x TSIG compatibility is used, this is set to 'tsig_xfer' on retrieval, and ignored on insert or update.",
	},
	"use_tsig_key_name": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Bool{
			boolvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("address"),
				path.MatchRelative().AtParent().AtName("permission"),
			),
		},
		MarkdownDescription: "Use flag for: tsig_key_name",
	},
}

func ExpandZoneAuthAllowUpdate(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneAuthAllowUpdate {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneAuthAllowUpdateModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneAuthAllowUpdateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZoneAuthAllowUpdate {
	if m == nil {
		return nil
	}
	to := &dns.ZoneAuthAllowUpdate{
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

func FlattenZoneAuthAllowUpdate(ctx context.Context, from *dns.ZoneAuthAllowUpdate, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneAuthAllowUpdateAttrTypes)
	}
	m := ZoneAuthAllowUpdateModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneAuthAllowUpdateAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneAuthAllowUpdateModel) Flatten(ctx context.Context, from *dns.ZoneAuthAllowUpdate, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneAuthAllowUpdateModel{}
	}
	m.Address = flex.FlattenIPAddress(from.Address)
	m.Struct = flex.FlattenStringPointer(from.Struct)
	m.Permission = flex.FlattenStringPointer(from.Permission)
	m.TsigKey = flex.FlattenStringPointer(from.TsigKey)
	m.TsigKeyAlg = flex.FlattenStringPointer(from.TsigKeyAlg)
	m.TsigKeyName = flex.FlattenStringPointer(from.TsigKeyName)
	m.UseTsigKeyName = types.BoolPointerValue(from.UseTsigKeyName)
}
