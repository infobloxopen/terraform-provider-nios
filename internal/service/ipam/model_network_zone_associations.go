package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/ipam"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type NetworkZoneAssociationsModel struct {
	Fqdn      types.String `tfsdk:"fqdn"`
	IsDefault types.Bool   `tfsdk:"is_default"`
	View      types.String `tfsdk:"view"`
}

var NetworkZoneAssociationsAttrTypes = map[string]attr.Type{
	"fqdn":       types.StringType,
	"is_default": types.BoolType,
	"view":       types.StringType,
}

var NetworkZoneAssociationsResourceSchemaAttributes = map[string]schema.Attribute{
	"fqdn": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The FQDN of the authoritative forward zone.",
	},
	"is_default": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "True if this is the default zone.",
	},
	"view": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The view to which the zone belongs. If a view is not specified, the default view is used.",
	},
}

func ExpandNetworkZoneAssociations(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkZoneAssociations {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkZoneAssociationsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkZoneAssociationsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkZoneAssociations {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkZoneAssociations{
		Fqdn:      flex.ExpandStringPointer(m.Fqdn),
		IsDefault: flex.ExpandBoolPointer(m.IsDefault),
		View:      flex.ExpandStringPointer(m.View),
	}
	return to
}

func FlattenNetworkZoneAssociations(ctx context.Context, from *ipam.NetworkZoneAssociations, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkZoneAssociationsAttrTypes)
	}
	m := NetworkZoneAssociationsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkZoneAssociationsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkZoneAssociationsModel) Flatten(ctx context.Context, from *ipam.NetworkZoneAssociations, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkZoneAssociationsModel{}
	}
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.IsDefault = types.BoolPointerValue(from.IsDefault)
	m.View = flex.FlattenStringPointer(from.View)
}
