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

type MemberbgpasNeighborsModel struct {
	Interface          types.String `tfsdk:"interface"`
	NeighborIp         types.String `tfsdk:"neighbor_ip"`
	RemoteAs           types.Int64  `tfsdk:"remote_as"`
	AuthenticationMode types.String `tfsdk:"authentication_mode"`
	BgpNeighborPass    types.String `tfsdk:"bgp_neighbor_pass"`
	Comment            types.String `tfsdk:"comment"`
	Multihop           types.Bool   `tfsdk:"multihop"`
	MultihopTtl        types.Int64  `tfsdk:"multihop_ttl"`
	BfdTemplate        types.String `tfsdk:"bfd_template"`
	EnableBfd          types.Bool   `tfsdk:"enable_bfd"`
	EnableBfdDnscheck  types.Bool   `tfsdk:"enable_bfd_dnscheck"`
}

var MemberbgpasNeighborsAttrTypes = map[string]attr.Type{
	"interface":           types.StringType,
	"neighbor_ip":         types.StringType,
	"remote_as":           types.Int64Type,
	"authentication_mode": types.StringType,
	"bgp_neighbor_pass":   types.StringType,
	"comment":             types.StringType,
	"multihop":            types.BoolType,
	"multihop_ttl":        types.Int64Type,
	"bfd_template":        types.StringType,
	"enable_bfd":          types.BoolType,
	"enable_bfd_dnscheck": types.BoolType,
}

var MemberbgpasNeighborsResourceSchemaAttributes = map[string]schema.Attribute{
	"interface": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The interface that sends BGP advertisement information.",
	},
	"neighbor_ip": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IP address of the BGP neighbor.",
	},
	"remote_as": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The remote AS number of the BGP neighbor.",
	},
	"authentication_mode": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The BGP authentication mode.",
	},
	"bgp_neighbor_pass": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The password for a BGP neighbor. This is required only if authentication_mode is set to \"MD5\". When the password is entered, the value is preserved even if authentication_mode is changed to \"NONE\". This is a write-only attribute.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "User comments for this BGP neighbor.",
	},
	"multihop": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the multi-hop support is enabled or not.",
	},
	"multihop_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Time To Live (TTL) value for multi-hop. Valid values are between 1 and 255.",
	},
	"bfd_template": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The BFD template name.",
	},
	"enable_bfd": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if BFD is enabled or not.",
	},
	"enable_bfd_dnscheck": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if BFD internal DNS monitor is enabled or not.",
	},
}

func ExpandMemberbgpasNeighbors(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberbgpasNeighbors {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberbgpasNeighborsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberbgpasNeighborsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberbgpasNeighbors {
	if m == nil {
		return nil
	}
	to := &grid.MemberbgpasNeighbors{
		Interface:          flex.ExpandStringPointer(m.Interface),
		NeighborIp:         flex.ExpandStringPointer(m.NeighborIp),
		RemoteAs:           flex.ExpandInt64Pointer(m.RemoteAs),
		AuthenticationMode: flex.ExpandStringPointer(m.AuthenticationMode),
		BgpNeighborPass:    flex.ExpandStringPointer(m.BgpNeighborPass),
		Comment:            flex.ExpandStringPointer(m.Comment),
		Multihop:           flex.ExpandBoolPointer(m.Multihop),
		MultihopTtl:        flex.ExpandInt64Pointer(m.MultihopTtl),
		BfdTemplate:        flex.ExpandStringPointer(m.BfdTemplate),
		EnableBfd:          flex.ExpandBoolPointer(m.EnableBfd),
		EnableBfdDnscheck:  flex.ExpandBoolPointer(m.EnableBfdDnscheck),
	}
	return to
}

func FlattenMemberbgpasNeighbors(ctx context.Context, from *grid.MemberbgpasNeighbors, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberbgpasNeighborsAttrTypes)
	}
	m := MemberbgpasNeighborsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberbgpasNeighborsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberbgpasNeighborsModel) Flatten(ctx context.Context, from *grid.MemberbgpasNeighbors, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberbgpasNeighborsModel{}
	}
	m.Interface = flex.FlattenStringPointer(from.Interface)
	m.NeighborIp = flex.FlattenStringPointer(from.NeighborIp)
	m.RemoteAs = flex.FlattenInt64Pointer(from.RemoteAs)
	m.AuthenticationMode = flex.FlattenStringPointer(from.AuthenticationMode)
	m.BgpNeighborPass = flex.FlattenStringPointer(from.BgpNeighborPass)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Multihop = types.BoolPointerValue(from.Multihop)
	m.MultihopTtl = flex.FlattenInt64Pointer(from.MultihopTtl)
	m.BfdTemplate = flex.FlattenStringPointer(from.BfdTemplate)
	m.EnableBfd = types.BoolPointerValue(from.EnableBfd)
	m.EnableBfdDnscheck = types.BoolPointerValue(from.EnableBfdDnscheck)
}
