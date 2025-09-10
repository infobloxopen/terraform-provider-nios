package security

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/security"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type AdmingroupAdminToplevelCommandsModel struct {
	Ps             types.Bool `tfsdk:"ps"`
	Iostat         types.Bool `tfsdk:"iostat"`
	Netstat        types.Bool `tfsdk:"netstat"`
	Vmstat         types.Bool `tfsdk:"vmstat"`
	Tcpdump        types.Bool `tfsdk:"tcpdump"`
	Rndc           types.Bool `tfsdk:"rndc"`
	Sar            types.Bool `tfsdk:"sar"`
	Resilver       types.Bool `tfsdk:"resilver"`
	RestartProduct types.Bool `tfsdk:"restart_product"`
	Scrape         types.Bool `tfsdk:"scrape"`
	SamlRestart    types.Bool `tfsdk:"saml_restart"`
	Synctime       types.Bool `tfsdk:"synctime"`
	EnableAll      types.Bool `tfsdk:"enable_all"`
	DisableAll     types.Bool `tfsdk:"disable_all"`
}

var AdmingroupAdminToplevelCommandsAttrTypes = map[string]attr.Type{
	"ps":              types.BoolType,
	"iostat":          types.BoolType,
	"netstat":         types.BoolType,
	"vmstat":          types.BoolType,
	"tcpdump":         types.BoolType,
	"rndc":            types.BoolType,
	"sar":             types.BoolType,
	"resilver":        types.BoolType,
	"restart_product": types.BoolType,
	"scrape":          types.BoolType,
	"saml_restart":    types.BoolType,
	"synctime":        types.BoolType,
	"enable_all":      types.BoolType,
	"disable_all":     types.BoolType,
}

var AdmingroupAdminToplevelCommandsResourceSchemaAttributes = map[string]schema.Attribute{
	"ps": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		//Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"iostat": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		//Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"netstat": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		//Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"vmstat": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		//Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"tcpdump": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		//Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"rndc": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		//Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"sar": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		//Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"resilver": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		//Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"restart_product": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		//Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"scrape": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		//Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"saml_restart": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		//Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"synctime": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		//Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"enable_all": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		//Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then enable all fields",
	},
	"disable_all": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		//Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True then disable all fields",
	},
}

func ExpandAdmingroupAdminToplevelCommands(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupAdminToplevelCommands {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupAdminToplevelCommandsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupAdminToplevelCommandsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupAdminToplevelCommands {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupAdminToplevelCommands{
		Ps:             flex.ExpandBoolPointer(m.Ps),
		Iostat:         flex.ExpandBoolPointer(m.Iostat),
		Netstat:        flex.ExpandBoolPointer(m.Netstat),
		Vmstat:         flex.ExpandBoolPointer(m.Vmstat),
		Tcpdump:        flex.ExpandBoolPointer(m.Tcpdump),
		Rndc:           flex.ExpandBoolPointer(m.Rndc),
		Sar:            flex.ExpandBoolPointer(m.Sar),
		Resilver:       flex.ExpandBoolPointer(m.Resilver),
		RestartProduct: flex.ExpandBoolPointer(m.RestartProduct),
		Scrape:         flex.ExpandBoolPointer(m.Scrape),
		SamlRestart:    flex.ExpandBoolPointer(m.SamlRestart),
		Synctime:       flex.ExpandBoolPointer(m.Synctime),
		EnableAll:      flex.ExpandBoolPointer(m.EnableAll),
		DisableAll:     flex.ExpandBoolPointer(m.DisableAll),
	}
	return to
}

func FlattenAdmingroupAdminToplevelCommands(ctx context.Context, from *security.AdmingroupAdminToplevelCommands, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupAdminToplevelCommandsAttrTypes)
	}
	m := AdmingroupAdminToplevelCommandsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupAdminToplevelCommandsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupAdminToplevelCommandsModel) Flatten(ctx context.Context, from *security.AdmingroupAdminToplevelCommands, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupAdminToplevelCommandsModel{}
	}
	m.Ps = types.BoolPointerValue(from.Ps)
	m.Iostat = types.BoolPointerValue(from.Iostat)
	m.Netstat = types.BoolPointerValue(from.Netstat)
	m.Vmstat = types.BoolPointerValue(from.Vmstat)
	m.Tcpdump = types.BoolPointerValue(from.Tcpdump)
	m.Rndc = types.BoolPointerValue(from.Rndc)
	m.Sar = types.BoolPointerValue(from.Sar)
	m.Resilver = types.BoolPointerValue(from.Resilver)
	m.RestartProduct = types.BoolPointerValue(from.RestartProduct)
	m.Scrape = types.BoolPointerValue(from.Scrape)
	m.SamlRestart = types.BoolPointerValue(from.SamlRestart)
	m.Synctime = types.BoolPointerValue(from.Synctime)
	m.EnableAll = types.BoolPointerValue(from.EnableAll)
	m.DisableAll = types.BoolPointerValue(from.DisableAll)
}
