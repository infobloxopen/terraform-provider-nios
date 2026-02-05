package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type TftpfiledirVtftpDirMembersModel struct {
	Member       types.String `tfsdk:"member"`
	IpType       types.String `tfsdk:"ip_type"`
	Address      types.String `tfsdk:"address"`
	StartAddress types.String `tfsdk:"start_address"`
	EndAddress   types.String `tfsdk:"end_address"`
	Network      types.String `tfsdk:"network"`
	Cidr         types.Int64  `tfsdk:"cidr"`
}

var TftpfiledirVtftpDirMembersAttrTypes = map[string]attr.Type{
	"member":        types.StringType,
	"ip_type":       types.StringType,
	"address":       types.StringType,
	"start_address": types.StringType,
	"end_address":   types.StringType,
	"network":       types.StringType,
	"cidr":          types.Int64Type,
}

var TftpfiledirVtftpDirMembersResourceSchemaAttributes = map[string]schema.Attribute{
	"member": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The Grid member on which to create the virtual TFTP directory.",
	},
	"ip_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IP type of the virtual TFTP root directory.",
	},
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IP address of the clients which will see the virtual TFTP directory as the root directory.",
	},
	"start_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The start IP address of the range within which the clients will see the virtual TFTP directory as the root directory.",
	},
	"end_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The end IP address of the range within which the clients will see the virtual TFTP directory as the root directory.",
	},
	"network": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IP address of network the clients from which will see the virtual TFTP directory as the root directory.",
	},
	"cidr": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The CIDR of network the clients from which will see the virtual TFTP directory as the root directory.",
	},
}

func ExpandTftpfiledirVtftpDirMembers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.TftpfiledirVtftpDirMembers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m TftpfiledirVtftpDirMembersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *TftpfiledirVtftpDirMembersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.TftpfiledirVtftpDirMembers {
	if m == nil {
		return nil
	}
	to := &misc.TftpfiledirVtftpDirMembers{
		Member:       flex.ExpandStringPointer(m.Member),
		IpType:       flex.ExpandStringPointer(m.IpType),
		Address:      flex.ExpandStringPointer(m.Address),
		StartAddress: flex.ExpandStringPointer(m.StartAddress),
		EndAddress:   flex.ExpandStringPointer(m.EndAddress),
		Network:      flex.ExpandStringPointer(m.Network),
		Cidr:         flex.ExpandInt64Pointer(m.Cidr),
	}
	return to
}

func FlattenTftpfiledirVtftpDirMembers(ctx context.Context, from *misc.TftpfiledirVtftpDirMembers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(TftpfiledirVtftpDirMembersAttrTypes)
	}
	m := TftpfiledirVtftpDirMembersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, TftpfiledirVtftpDirMembersAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *TftpfiledirVtftpDirMembersModel) Flatten(ctx context.Context, from *misc.TftpfiledirVtftpDirMembers, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = TftpfiledirVtftpDirMembersModel{}
	}
	m.Member = flex.FlattenStringPointer(from.Member)
	m.IpType = flex.FlattenStringPointer(from.IpType)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.StartAddress = flex.FlattenStringPointer(from.StartAddress)
	m.EndAddress = flex.FlattenStringPointer(from.EndAddress)
	m.Network = flex.FlattenStringPointer(from.Network)
	m.Cidr = flex.FlattenInt64Pointer(from.Cidr)
}
