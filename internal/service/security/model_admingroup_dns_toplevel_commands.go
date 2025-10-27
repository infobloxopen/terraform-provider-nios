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

type AdmingroupDnsToplevelCommandsModel struct {
	DdnsAdd          types.Bool `tfsdk:"ddns_add"`
	DdnsDelete       types.Bool `tfsdk:"ddns_delete"`
	Delete           types.Bool `tfsdk:"delete"`
	DnsARecordDelete types.Bool `tfsdk:"dns_a_record_delete"`
	EnableAll        types.Bool `tfsdk:"enable_all"`
	DisableAll       types.Bool `tfsdk:"disable_all"`
}

var AdmingroupDnsToplevelCommandsAttrTypes = map[string]attr.Type{
	"ddns_add":            types.BoolType,
	"ddns_delete":         types.BoolType,
	"delete":              types.BoolType,
	"dns_a_record_delete": types.BoolType,
	"enable_all":          types.BoolType,
	"disable_all":         types.BoolType,
}

var AdmingroupDnsToplevelCommandsResourceSchemaAttributes = map[string]schema.Attribute{
	"ddns_add": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"ddns_delete": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"delete": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"dns_a_record_delete": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If True then CLI user has permission to run the command",
	},
	"enable_all": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "If True then enable all fields",
	},
	"disable_all": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "If True then disable all fields",
	},
}

func ExpandAdmingroupDnsToplevelCommands(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.AdmingroupDnsToplevelCommands {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdmingroupDnsToplevelCommandsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdmingroupDnsToplevelCommandsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.AdmingroupDnsToplevelCommands {
	if m == nil {
		return nil
	}
	to := &security.AdmingroupDnsToplevelCommands{
		DdnsAdd:          flex.ExpandBoolPointer(m.DdnsAdd),
		DdnsDelete:       flex.ExpandBoolPointer(m.DdnsDelete),
		Delete:           flex.ExpandBoolPointer(m.Delete),
		DnsARecordDelete: flex.ExpandBoolPointer(m.DnsARecordDelete),
		EnableAll:        flex.ExpandBoolPointer(m.EnableAll),
		DisableAll:       flex.ExpandBoolPointer(m.DisableAll),
	}
	return to
}

func FlattenAdmingroupDnsToplevelCommands(ctx context.Context, from *security.AdmingroupDnsToplevelCommands, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdmingroupDnsToplevelCommandsAttrTypes)
	}
	m := AdmingroupDnsToplevelCommandsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdmingroupDnsToplevelCommandsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdmingroupDnsToplevelCommandsModel) Flatten(ctx context.Context, from *security.AdmingroupDnsToplevelCommands, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdmingroupDnsToplevelCommandsModel{}
	}
	m.DdnsAdd = types.BoolPointerValue(from.DdnsAdd)
	m.DdnsDelete = types.BoolPointerValue(from.DdnsDelete)
	m.Delete = types.BoolPointerValue(from.Delete)
	m.DnsARecordDelete = types.BoolPointerValue(from.DnsARecordDelete)
	m.EnableAll = types.BoolPointerValue(from.EnableAll)
	m.DisableAll = types.BoolPointerValue(from.DisableAll)
}
