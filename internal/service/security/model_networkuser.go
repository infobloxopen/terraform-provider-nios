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

type NetworkuserModel struct {
	Ref             types.String `tfsdk:"ref"`
    Uuid        types.String `tfsdk:"uuid"`
	Address         types.String `tfsdk:"address"`
	AddressObject   types.String `tfsdk:"address_object"`
	DataSource      types.String `tfsdk:"data_source"`
	DataSourceIp    types.String `tfsdk:"data_source_ip"`
	Domainname      types.String `tfsdk:"domainname"`
	FirstSeenTime   types.Int64  `tfsdk:"first_seen_time"`
	Guid            types.String `tfsdk:"guid"`
	LastSeenTime    types.Int64  `tfsdk:"last_seen_time"`
	LastUpdatedTime types.Int64  `tfsdk:"last_updated_time"`
	LogonId         types.String `tfsdk:"logon_id"`
	LogoutTime      types.Int64  `tfsdk:"logout_time"`
	Name            types.String `tfsdk:"name"`
	Network         types.String `tfsdk:"network"`
	NetworkView     types.String `tfsdk:"network_view"`
	UserStatus      types.String `tfsdk:"user_status"`
}

var NetworkuserAttrTypes = map[string]attr.Type{
	"ref":               types.StringType,
    "uuid":        types.StringType,
	"address":           types.StringType,
	"address_object":    types.StringType,
	"data_source":       types.StringType,
	"data_source_ip":    types.StringType,
	"domainname":        types.StringType,
	"first_seen_time":   types.Int64Type,
	"guid":              types.StringType,
	"last_seen_time":    types.Int64Type,
	"last_updated_time": types.Int64Type,
	"logon_id":          types.StringType,
	"logout_time":       types.Int64Type,
	"name":              types.StringType,
	"network":           types.StringType,
	"network_view":      types.StringType,
	"user_status":       types.StringType,
}

var NetworkuserResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
    "uuid": schema.StringAttribute{
        Computed:            true,
        MarkdownDescription: "The uuid to the object.",
    },
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IPv4 Address or IPv6 Address of the Network User.",
	},
	"address_object": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference of the IPAM IPv4Address or IPv6Address object describing the address of the Network User.",
	},
	"data_source": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Network User data source.",
	},
	"data_source_ip": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Network User data source IPv4 Address or IPv6 Address or FQDN address.",
	},
	"domainname": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The domain name of the Network User.",
	},
	"first_seen_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The first seen timestamp of the Network User.",
	},
	"guid": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The group identifier of the Network User.",
	},
	"last_seen_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The last seen timestamp of the Network User.",
	},
	"last_updated_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The last updated timestamp of the Network User.",
	},
	"logon_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The logon identifier of the Network User.",
	},
	"logout_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The logout timestamp of the Network User.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the Network User.",
	},
	"network": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the network to which the Network User belongs.",
	},
	"network_view": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the network view in which this Network User resides.",
	},
	"user_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The status of the Network User.",
	},
}

func ExpandNetworkuser(ctx context.Context, o types.Object, diags *diag.Diagnostics) *security.Networkuser {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkuserModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkuserModel) Expand(ctx context.Context, diags *diag.Diagnostics) *security.Networkuser {
	if m == nil {
		return nil
	}
	to := &security.Networkuser{
		Ref:             flex.ExpandStringPointer(m.Ref),
		Address:         flex.ExpandStringPointer(m.Address),
		Domainname:      flex.ExpandStringPointer(m.Domainname),
		FirstSeenTime:   flex.ExpandInt64Pointer(m.FirstSeenTime),
		Guid:            flex.ExpandStringPointer(m.Guid),
		LastSeenTime:    flex.ExpandInt64Pointer(m.LastSeenTime),
		LastUpdatedTime: flex.ExpandInt64Pointer(m.LastUpdatedTime),
		LogonId:         flex.ExpandStringPointer(m.LogonId),
		LogoutTime:      flex.ExpandInt64Pointer(m.LogoutTime),
		Name:            flex.ExpandStringPointer(m.Name),
		NetworkView:     flex.ExpandStringPointer(m.NetworkView),
	}
	return to
}

func FlattenNetworkuser(ctx context.Context, from *security.Networkuser, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkuserAttrTypes)
	}
	m := NetworkuserModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkuserAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkuserModel) Flatten(ctx context.Context, from *security.Networkuser, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkuserModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
    m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.AddressObject = flex.FlattenStringPointer(from.AddressObject)
	m.DataSource = flex.FlattenStringPointer(from.DataSource)
	m.DataSourceIp = flex.FlattenStringPointer(from.DataSourceIp)
	m.Domainname = flex.FlattenStringPointer(from.Domainname)
	m.FirstSeenTime = flex.FlattenInt64Pointer(from.FirstSeenTime)
	m.Guid = flex.FlattenStringPointer(from.Guid)
	m.LastSeenTime = flex.FlattenInt64Pointer(from.LastSeenTime)
	m.LastUpdatedTime = flex.FlattenInt64Pointer(from.LastUpdatedTime)
	m.LogonId = flex.FlattenStringPointer(from.LogonId)
	m.LogoutTime = flex.FlattenInt64Pointer(from.LogoutTime)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Network = flex.FlattenStringPointer(from.Network)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.UserStatus = flex.FlattenStringPointer(from.UserStatus)
}
