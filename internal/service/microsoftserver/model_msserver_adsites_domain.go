package microsoftserver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/microsoftserver"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type MsserverAdsitesDomainModel struct {
	Ref              types.String `tfsdk:"ref"`
	EaDefinition     types.String `tfsdk:"ea_definition"`
	MsSyncMasterName types.String `tfsdk:"ms_sync_master_name"`
	Name             types.String `tfsdk:"name"`
	Netbios          types.String `tfsdk:"netbios"`
	NetworkView      types.String `tfsdk:"network_view"`
	ReadOnly         types.Bool   `tfsdk:"read_only"`
}

var MsserverAdsitesDomainAttrTypes = map[string]attr.Type{
	"ref":                 types.StringType,
	"ea_definition":       types.StringType,
	"ms_sync_master_name": types.StringType,
	"name":                types.StringType,
	"netbios":             types.StringType,
	"network_view":        types.StringType,
	"read_only":           types.BoolType,
}

var MsserverAdsitesDomainResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"ea_definition": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the Extensible Attribute Definition object that is linked to the Active Directory Sites Domain.",
	},
	"ms_sync_master_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IP address or FQDN of the managing master for the MS server, if applicable.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the Active Directory Domain properties object.",
	},
	"netbios": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The NetBIOS name of the Active Directory Domain properties object.",
	},
	"network_view": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the network view in which the Active Directory Domain resides.",
	},
	"read_only": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines whether the Active Directory Domain properties object is a read-only object.",
	},
}

func ExpandMsserverAdsitesDomain(ctx context.Context, o types.Object, diags *diag.Diagnostics) *microsoftserver.MsserverAdsitesDomain {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MsserverAdsitesDomainModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MsserverAdsitesDomainModel) Expand(ctx context.Context, diags *diag.Diagnostics) *microsoftserver.MsserverAdsitesDomain {
	if m == nil {
		return nil
	}
	to := &microsoftserver.MsserverAdsitesDomain{
		Ref: flex.ExpandStringPointer(m.Ref),
	}
	return to
}

func FlattenMsserverAdsitesDomain(ctx context.Context, from *microsoftserver.MsserverAdsitesDomain, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MsserverAdsitesDomainAttrTypes)
	}
	m := MsserverAdsitesDomainModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MsserverAdsitesDomainAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MsserverAdsitesDomainModel) Flatten(ctx context.Context, from *microsoftserver.MsserverAdsitesDomain, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MsserverAdsitesDomainModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.EaDefinition = flex.FlattenStringPointer(from.EaDefinition)
	m.MsSyncMasterName = flex.FlattenStringPointer(from.MsSyncMasterName)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Netbios = flex.FlattenStringPointer(from.Netbios)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.ReadOnly = types.BoolPointerValue(from.ReadOnly)
}
