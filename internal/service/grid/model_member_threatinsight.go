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

type MemberThreatinsightModel struct {
	Ref           types.String `tfsdk:"ref"`
	Uuid          types.String `tfsdk:"uuid"`
	Comment       types.String `tfsdk:"comment"`
	EnableService types.Bool   `tfsdk:"enable_service"`
	HostName      types.String `tfsdk:"host_name"`
	Ipv4Address   types.String `tfsdk:"ipv4_address"`
	Ipv6Address   types.String `tfsdk:"ipv6_address"`
	Status        types.String `tfsdk:"status"`
}

var MemberThreatinsightAttrTypes = map[string]attr.Type{
	"ref":            types.StringType,
	"uuid":           types.StringType,
	"comment":        types.StringType,
	"enable_service": types.BoolType,
	"host_name":      types.StringType,
	"ipv4_address":   types.StringType,
	"ipv6_address":   types.StringType,
	"status":         types.StringType,
}

var MemberThreatinsightResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The UUID of the object.",
	},
	"comment": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Grid member descriptive comment.",
	},
	"enable_service": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the threat insight service is enabled.",
	},
	"host_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Grid member host name.",
	},
	"ipv4_address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IPv4 Address address of the Grid member.",
	},
	"ipv6_address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IPv6 Address address of the Grid member.",
	},
	"status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Grid member threat insight status.",
	},
}

func ExpandMemberThreatinsight(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberThreatinsight {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberThreatinsightModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberThreatinsightModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberThreatinsight {
	if m == nil {
		return nil
	}
	to := &grid.MemberThreatinsight{
		Ref:           flex.ExpandStringPointer(m.Ref),
		Uuid:          flex.ExpandStringPointer(m.Uuid),
		EnableService: flex.ExpandBoolPointer(m.EnableService),
	}
	return to
}

func FlattenMemberThreatinsight(ctx context.Context, from *grid.MemberThreatinsight, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberThreatinsightAttrTypes)
	}
	m := MemberThreatinsightModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberThreatinsightAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberThreatinsightModel) Flatten(ctx context.Context, from *grid.MemberThreatinsight, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberThreatinsightModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.EnableService = types.BoolPointerValue(from.EnableService)
	m.HostName = flex.FlattenStringPointer(from.HostName)
	m.Ipv4Address = flex.FlattenStringPointer(from.Ipv4Address)
	m.Ipv6Address = flex.FlattenStringPointer(from.Ipv6Address)
	m.Status = flex.FlattenStringPointer(from.Status)
}
