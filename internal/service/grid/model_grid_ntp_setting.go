package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type GridNtpSettingModel struct {
	EnableNtp         types.Bool   `tfsdk:"enable_ntp"`
	NtpServers        types.List   `tfsdk:"ntp_servers"`
	NtpKeys           types.List   `tfsdk:"ntp_keys"`
	NtpAcl            types.Object `tfsdk:"ntp_acl"`
	NtpKod            types.Bool   `tfsdk:"ntp_kod"`
	GmLocalNtpStratum types.Int64  `tfsdk:"gm_local_ntp_stratum"`
	LocalNtpStratum   types.Int64  `tfsdk:"local_ntp_stratum"`
	UseDefaultStratum types.Bool   `tfsdk:"use_default_stratum"`
}

var GridNtpSettingAttrTypes = map[string]attr.Type{
	"enable_ntp":           types.BoolType,
	"ntp_servers":          types.ListType{ElemType: types.ObjectType{AttrTypes: GridntpsettingNtpServersAttrTypes}},
	"ntp_keys":             types.ListType{ElemType: types.ObjectType{AttrTypes: GridntpsettingNtpKeysAttrTypes}},
	"ntp_acl":              types.ObjectType{AttrTypes: GridntpsettingNtpAclAttrTypes},
	"ntp_kod":              types.BoolType,
	"gm_local_ntp_stratum": types.Int64Type,
	"local_ntp_stratum":    types.Int64Type,
	"use_default_stratum":  types.BoolType,
}

var GridNtpSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"enable_ntp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether NTP is enabled on the Grid.",
	},
	"ntp_servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridntpsettingNtpServersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of NTP servers configured on a Grid.",
	},
	"ntp_keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: GridntpsettingNtpKeysResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of NTP authentication keys used to authenticate NTP clients.",
	},
	"ntp_acl": schema.SingleNestedAttribute{
		Attributes: GridntpsettingNtpAclResourceSchemaAttributes,
		Optional:   true,
	},
	"ntp_kod": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the Kiss-o'-Death packets are enabled.",
	},
	"gm_local_ntp_stratum": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Grid level GM local NTP stratum.",
	},
	"local_ntp_stratum": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Local NTP stratum for non-GM members.",
	},
	"use_default_stratum": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This flag controls whether gm_local_ntp_stratum value be set to a default value",
	},
}

func ExpandGridNtpSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridNtpSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridNtpSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridNtpSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridNtpSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridNtpSetting{
		EnableNtp:         flex.ExpandBoolPointer(m.EnableNtp),
		NtpServers:        flex.ExpandFrameworkListNestedBlock(ctx, m.NtpServers, diags, ExpandGridntpsettingNtpServers),
		NtpKeys:           flex.ExpandFrameworkListNestedBlock(ctx, m.NtpKeys, diags, ExpandGridntpsettingNtpKeys),
		NtpAcl:            ExpandGridntpsettingNtpAcl(ctx, m.NtpAcl, diags),
		NtpKod:            flex.ExpandBoolPointer(m.NtpKod),
		GmLocalNtpStratum: flex.ExpandInt64Pointer(m.GmLocalNtpStratum),
		LocalNtpStratum:   flex.ExpandInt64Pointer(m.LocalNtpStratum),
		UseDefaultStratum: flex.ExpandBoolPointer(m.UseDefaultStratum),
	}
	return to
}

func FlattenGridNtpSetting(ctx context.Context, from *grid.GridNtpSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridNtpSettingAttrTypes)
	}
	m := GridNtpSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridNtpSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridNtpSettingModel) Flatten(ctx context.Context, from *grid.GridNtpSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridNtpSettingModel{}
	}
	m.EnableNtp = types.BoolPointerValue(from.EnableNtp)
	m.NtpServers = flex.FlattenFrameworkListNestedBlock(ctx, from.NtpServers, GridntpsettingNtpServersAttrTypes, diags, FlattenGridntpsettingNtpServers)
	m.NtpKeys = flex.FlattenFrameworkListNestedBlock(ctx, from.NtpKeys, GridntpsettingNtpKeysAttrTypes, diags, FlattenGridntpsettingNtpKeys)
	m.NtpAcl = FlattenGridntpsettingNtpAcl(ctx, from.NtpAcl, diags)
	m.NtpKod = types.BoolPointerValue(from.NtpKod)
	m.GmLocalNtpStratum = flex.FlattenInt64Pointer(from.GmLocalNtpStratum)
	m.LocalNtpStratum = flex.FlattenInt64Pointer(from.LocalNtpStratum)
	m.UseDefaultStratum = types.BoolPointerValue(from.UseDefaultStratum)
}
