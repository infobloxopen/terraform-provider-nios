package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MemberCspMemberSettingModel struct {
	UseCspJoinToken   types.Bool   `tfsdk:"use_csp_join_token"`
	UseCspDnsResolver types.Bool   `tfsdk:"use_csp_dns_resolver"`
	UseCspHttpsProxy  types.Bool   `tfsdk:"use_csp_https_proxy"`
	CspJoinToken      types.String `tfsdk:"csp_join_token"`
	CspDnsResolver    types.String `tfsdk:"csp_dns_resolver"`
	CspHttpsProxy     types.String `tfsdk:"csp_https_proxy"`
}

var MemberCspMemberSettingAttrTypes = map[string]attr.Type{
	"use_csp_join_token":   types.BoolType,
	"use_csp_dns_resolver": types.BoolType,
	"use_csp_https_proxy":  types.BoolType,
	"csp_join_token":       types.StringType,
	"csp_dns_resolver":     types.StringType,
	"csp_https_proxy":      types.StringType,
}

var MemberCspMemberSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"use_csp_join_token": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Overrides grid join token",
	},
	"use_csp_dns_resolver": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Overrides CSP DNS Resolver",
	},
	"use_csp_https_proxy": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Overrides grid https proxy",
	},
	"csp_join_token": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Join token required to connect to a cluster",
	},
	"csp_dns_resolver": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "IP address of DNS resolver in DFP",
	},
	"csp_https_proxy": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "HTTP Proxy IP address of CSP Portal",
	},
}

func ExpandMemberCspMemberSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberCspMemberSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberCspMemberSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberCspMemberSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberCspMemberSetting {
	if m == nil {
		return nil
	}
	to := &grid.MemberCspMemberSetting{
		UseCspJoinToken:   flex.ExpandBoolPointer(m.UseCspJoinToken),
		UseCspDnsResolver: flex.ExpandBoolPointer(m.UseCspDnsResolver),
		UseCspHttpsProxy:  flex.ExpandBoolPointer(m.UseCspHttpsProxy),
		CspJoinToken:      flex.ExpandStringPointer(m.CspJoinToken),
		CspDnsResolver:    flex.ExpandStringPointer(m.CspDnsResolver),
		CspHttpsProxy:     flex.ExpandStringPointer(m.CspHttpsProxy),
	}
	return to
}

func FlattenMemberCspMemberSetting(ctx context.Context, from *grid.MemberCspMemberSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberCspMemberSettingAttrTypes)
	}
	m := MemberCspMemberSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberCspMemberSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberCspMemberSettingModel) Flatten(ctx context.Context, from *grid.MemberCspMemberSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberCspMemberSettingModel{}
	}
	m.UseCspJoinToken = types.BoolPointerValue(from.UseCspJoinToken)
	m.UseCspDnsResolver = types.BoolPointerValue(from.UseCspDnsResolver)
	m.UseCspHttpsProxy = types.BoolPointerValue(from.UseCspHttpsProxy)
	m.CspJoinToken = flex.FlattenStringPointer(from.CspJoinToken)
	m.CspDnsResolver = flex.FlattenStringPointer(from.CspDnsResolver)
	m.CspHttpsProxy = flex.FlattenStringPointer(from.CspHttpsProxy)
}
