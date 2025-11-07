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

type MemberNtpSettingModel struct {
	EnableNtp                  types.Bool   `tfsdk:"enable_ntp"`
	NtpServers                 types.List   `tfsdk:"ntp_servers"`
	NtpKeys                    types.List   `tfsdk:"ntp_keys"`
	NtpAcl                     types.Object `tfsdk:"ntp_acl"`
	NtpKod                     types.Bool   `tfsdk:"ntp_kod"`
	EnableExternalNtpServers   types.Bool   `tfsdk:"enable_external_ntp_servers"`
	ExcludeGridMasterNtpServer types.Bool   `tfsdk:"exclude_grid_master_ntp_server"`
	UseLocalNtpStratum         types.Bool   `tfsdk:"use_local_ntp_stratum"`
	LocalNtpStratum            types.Int64  `tfsdk:"local_ntp_stratum"`
	UseDefaultStratum          types.Bool   `tfsdk:"use_default_stratum"`
	UseNtpServers              types.Bool   `tfsdk:"use_ntp_servers"`
	UseNtpKeys                 types.Bool   `tfsdk:"use_ntp_keys"`
	UseNtpAcl                  types.Bool   `tfsdk:"use_ntp_acl"`
	UseNtpKod                  types.Bool   `tfsdk:"use_ntp_kod"`
}

var MemberNtpSettingAttrTypes = map[string]attr.Type{
	"enable_ntp":                     types.BoolType,
	"ntp_servers":                    types.ListType{ElemType: types.ObjectType{AttrTypes: MemberntpsettingNtpServersAttrTypes}},
	"ntp_keys":                       types.ListType{ElemType: types.ObjectType{AttrTypes: MemberntpsettingNtpKeysAttrTypes}},
	"ntp_acl":                        types.ObjectType{AttrTypes: MemberntpsettingNtpAclAttrTypes},
	"ntp_kod":                        types.BoolType,
	"enable_external_ntp_servers":    types.BoolType,
	"exclude_grid_master_ntp_server": types.BoolType,
	"use_local_ntp_stratum":          types.BoolType,
	"local_ntp_stratum":              types.Int64Type,
	"use_default_stratum":            types.BoolType,
	"use_ntp_servers":                types.BoolType,
	"use_ntp_keys":                   types.BoolType,
	"use_ntp_acl":                    types.BoolType,
	"use_ntp_kod":                    types.BoolType,
}

var MemberNtpSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"enable_ntp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the NTP service is enabled on the member.",
	},
	"ntp_servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberntpsettingNtpServersResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of NTP servers configured on a member.",
	},
	"ntp_keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: MemberntpsettingNtpKeysResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of NTP authentication keys used to authenticate NTP clients.",
	},
	"ntp_acl": schema.SingleNestedAttribute{
		Attributes: MemberntpsettingNtpAclResourceSchemaAttributes,
		Optional:   true,
	},
	"ntp_kod": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the Kiss-o'-Death packets are enabled or disabled.",
	},
	"enable_external_ntp_servers": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the use of the external NTP servers is enabled for the member.",
	},
	"exclude_grid_master_ntp_server": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the Grid Master is excluded as an NTP server.",
	},
	"use_local_ntp_stratum": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Override Grid level NTP stratum.",
	},
	"local_ntp_stratum": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Vnode level local NTP stratum.",
	},
	"use_default_stratum": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Vnode level default stratum.",
	},
	"use_ntp_servers": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ntp_servers",
	},
	"use_ntp_keys": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ntp_keys",
	},
	"use_ntp_acl": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ntp_acl",
	},
	"use_ntp_kod": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ntp_kod",
	},
}

func ExpandMemberNtpSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberNtpSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberNtpSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberNtpSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberNtpSetting {
	if m == nil {
		return nil
	}
	to := &grid.MemberNtpSetting{
		EnableNtp:                  flex.ExpandBoolPointer(m.EnableNtp),
		NtpServers:                 flex.ExpandFrameworkListNestedBlock(ctx, m.NtpServers, diags, ExpandMemberntpsettingNtpServers),
		NtpKeys:                    flex.ExpandFrameworkListNestedBlock(ctx, m.NtpKeys, diags, ExpandMemberntpsettingNtpKeys),
		NtpAcl:                     ExpandMemberntpsettingNtpAcl(ctx, m.NtpAcl, diags),
		NtpKod:                     flex.ExpandBoolPointer(m.NtpKod),
		EnableExternalNtpServers:   flex.ExpandBoolPointer(m.EnableExternalNtpServers),
		ExcludeGridMasterNtpServer: flex.ExpandBoolPointer(m.ExcludeGridMasterNtpServer),
		UseLocalNtpStratum:         flex.ExpandBoolPointer(m.UseLocalNtpStratum),
		LocalNtpStratum:            flex.ExpandInt64Pointer(m.LocalNtpStratum),
		UseDefaultStratum:          flex.ExpandBoolPointer(m.UseDefaultStratum),
		UseNtpServers:              flex.ExpandBoolPointer(m.UseNtpServers),
		UseNtpKeys:                 flex.ExpandBoolPointer(m.UseNtpKeys),
		UseNtpAcl:                  flex.ExpandBoolPointer(m.UseNtpAcl),
		UseNtpKod:                  flex.ExpandBoolPointer(m.UseNtpKod),
	}
	return to
}

func FlattenMemberNtpSetting(ctx context.Context, from *grid.MemberNtpSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberNtpSettingAttrTypes)
	}
	m := MemberNtpSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberNtpSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberNtpSettingModel) Flatten(ctx context.Context, from *grid.MemberNtpSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberNtpSettingModel{}
	}
	m.EnableNtp = types.BoolPointerValue(from.EnableNtp)
	m.NtpServers = flex.FlattenFrameworkListNestedBlock(ctx, from.NtpServers, MemberntpsettingNtpServersAttrTypes, diags, FlattenMemberntpsettingNtpServers)
	m.NtpKeys = flex.FlattenFrameworkListNestedBlock(ctx, from.NtpKeys, MemberntpsettingNtpKeysAttrTypes, diags, FlattenMemberntpsettingNtpKeys)
	m.NtpAcl = FlattenMemberntpsettingNtpAcl(ctx, from.NtpAcl, diags)
	m.NtpKod = types.BoolPointerValue(from.NtpKod)
	m.EnableExternalNtpServers = types.BoolPointerValue(from.EnableExternalNtpServers)
	m.ExcludeGridMasterNtpServer = types.BoolPointerValue(from.ExcludeGridMasterNtpServer)
	m.UseLocalNtpStratum = types.BoolPointerValue(from.UseLocalNtpStratum)
	m.LocalNtpStratum = flex.FlattenInt64Pointer(from.LocalNtpStratum)
	m.UseDefaultStratum = types.BoolPointerValue(from.UseDefaultStratum)
	m.UseNtpServers = types.BoolPointerValue(from.UseNtpServers)
	m.UseNtpKeys = types.BoolPointerValue(from.UseNtpKeys)
	m.UseNtpAcl = types.BoolPointerValue(from.UseNtpAcl)
	m.UseNtpKod = types.BoolPointerValue(from.UseNtpKod)
}
