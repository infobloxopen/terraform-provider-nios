package acl

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/acl"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	"github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/plancontrol"
)

type NamedaclExplodedAccessListModel struct {
	Struct         types.String `tfsdk:"struct"`
	Address        types.String `tfsdk:"address"`
	Permission     types.String `tfsdk:"permission"`
	TsigKey        types.String `tfsdk:"tsig_key"`
	TsigKeyAlg     types.String `tfsdk:"tsig_key_alg"`
	TsigKeyName    types.String `tfsdk:"tsig_key_name"`
	UseTsigKeyName types.Bool   `tfsdk:"use_tsig_key_name"`
}

var NamedaclExplodedAccessListAttrTypes = map[string]attr.Type{
	"struct":            types.StringType,
	"address":           types.StringType,
	"permission":        types.StringType,
	"tsig_key":          types.StringType,
	"tsig_key_alg":      types.StringType,
	"tsig_key_name":     types.StringType,
	"use_tsig_key_name": types.BoolType,
}

var NamedaclExplodedAccessListResourceSchemaAttributes = map[string]schema.Attribute{
	"struct": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The struct type of the object. The value must be one of 'addressac' and 'tsigac'.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The address this rule applies to or \"Any\".",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"permission": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The permission to use for this address.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"tsig_key": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "A generated TSIG key. If the external primary server is a NIOS appliance running DNS One 2.x code, this can be set to :2xCOMPAT.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"tsig_key_alg": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The TSIG key algorithm.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"tsig_key_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the TSIG key. If 2.x TSIG compatibility is used, this is set to 'tsig_xfer' on retrieval, and ignored on insert or update.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"use_tsig_key_name": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Use flag for: tsig_key_name",
		PlanModifiers: []planmodifier.Bool{
			plancontrol.UseStateForUnknownBool(),
		},
	},
}

func ExpandNamedaclExplodedAccessList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *acl.NamedaclExplodedAccessList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NamedaclExplodedAccessListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NamedaclExplodedAccessListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *acl.NamedaclExplodedAccessList {
	if m == nil {
		return nil
	}
	to := &acl.NamedaclExplodedAccessList{}
	return to
}

func FlattenNamedaclExplodedAccessList(ctx context.Context, from *acl.NamedaclExplodedAccessList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NamedaclExplodedAccessListAttrTypes)
	}
	m := NamedaclExplodedAccessListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NamedaclExplodedAccessListAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NamedaclExplodedAccessListModel) Flatten(ctx context.Context, from *acl.NamedaclExplodedAccessList, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NamedaclExplodedAccessListModel{}
	}
	m.Struct = flex.FlattenStringPointer(from.Struct)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Permission = flex.FlattenStringPointer(from.Permission)
	m.TsigKey = flex.FlattenStringPointer(from.TsigKey)
	m.TsigKeyAlg = flex.FlattenStringPointer(from.TsigKeyAlg)
	m.TsigKeyName = flex.FlattenStringPointer(from.TsigKeyName)
	m.UseTsigKeyName = types.BoolPointerValue(from.UseTsigKeyName)
}
