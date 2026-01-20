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

type AllendpointsModel struct {
	Ref               types.String `tfsdk:"ref"`
	Uuid              types.String `tfsdk:"uuid"`
	Address           types.String `tfsdk:"address"`
	Comment           types.String `tfsdk:"comment"`
	Disable           types.Bool   `tfsdk:"disable"`
	SubscribingMember types.String `tfsdk:"subscribing_member"`
	Type              types.String `tfsdk:"type"`
	Version           types.String `tfsdk:"version"`
}

var AllendpointsAttrTypes = map[string]attr.Type{
	"ref":                types.StringType,
	"uuid":               types.StringType,
	"address":            types.StringType,
	"comment":            types.StringType,
	"disable":            types.BoolType,
	"subscribing_member": types.StringType,
	"type":               types.StringType,
	"version":            types.StringType,
}

var AllendpointsResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The UUID of the object.",
	},
	"address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Grid endpoint IPv4 Address or IPv6 Address or Fully-Qualified Domain Name (FQDN).",
	},
	"comment": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Grid endpoint descriptive comment.",
	},
	"disable": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines whether a Grid endpoint is disabled or not. When this is set to False, the Grid endpoint is enabled.",
	},
	"subscribing_member": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the Grid Member object that is serving Grid endpoint.",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Grid endpoint type.",
	},
	"version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Grid endpoint version.",
	},
}

func ExpandAllendpoints(ctx context.Context, o types.Object, diags *diag.Diagnostics) *misc.Allendpoints {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AllendpointsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AllendpointsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *misc.Allendpoints {
	if m == nil {
		return nil
	}
	to := &misc.Allendpoints{
		Ref:  flex.ExpandStringPointer(m.Ref),
		Uuid: flex.ExpandStringPointer(m.Uuid),
	}
	return to
}

func FlattenAllendpoints(ctx context.Context, from *misc.Allendpoints, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AllendpointsAttrTypes)
	}
	m := AllendpointsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AllendpointsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AllendpointsModel) Flatten(ctx context.Context, from *misc.Allendpoints, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AllendpointsModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.SubscribingMember = flex.FlattenStringPointer(from.SubscribingMember)
	m.Type = flex.FlattenStringPointer(from.Type)
	m.Version = flex.FlattenStringPointer(from.Version)
}
