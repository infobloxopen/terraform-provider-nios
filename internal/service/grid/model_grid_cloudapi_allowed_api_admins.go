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

type GridCloudapiAllowedApiAdminsModel struct {
	IsRemote    types.Bool   `tfsdk:"is_remote"`
	RemoteAdmin types.String `tfsdk:"remote_admin"`
	LocalAdmin  types.String `tfsdk:"local_admin"`
}

var GridCloudapiAllowedApiAdminsAttrTypes = map[string]attr.Type{
	"is_remote":    types.BoolType,
	"remote_admin": types.StringType,
	"local_admin":  types.StringType,
}

var GridCloudapiAllowedApiAdminsResourceSchemaAttributes = map[string]schema.Attribute{
	"is_remote": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether this is a remote admin user.",
	},
	"remote_admin": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Username that matches a remote administrator who can perform cloud API requests on the Cloud Platform Appliance.",
	},
	"local_admin": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Local administrator who can perform cloud API requests on the Cloud Platform Appliance.",
	},
}

func ExpandGridCloudapiAllowedApiAdmins(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridCloudapiAllowedApiAdmins {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridCloudapiAllowedApiAdminsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridCloudapiAllowedApiAdminsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridCloudapiAllowedApiAdmins {
	if m == nil {
		return nil
	}
	to := &grid.GridCloudapiAllowedApiAdmins{
		IsRemote:    flex.ExpandBoolPointer(m.IsRemote),
		RemoteAdmin: flex.ExpandStringPointer(m.RemoteAdmin),
		LocalAdmin:  flex.ExpandStringPointer(m.LocalAdmin),
	}
	return to
}

func FlattenGridCloudapiAllowedApiAdmins(ctx context.Context, from *grid.GridCloudapiAllowedApiAdmins, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridCloudapiAllowedApiAdminsAttrTypes)
	}
	m := GridCloudapiAllowedApiAdminsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridCloudapiAllowedApiAdminsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridCloudapiAllowedApiAdminsModel) Flatten(ctx context.Context, from *grid.GridCloudapiAllowedApiAdmins, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridCloudapiAllowedApiAdminsModel{}
	}
	m.IsRemote = types.BoolPointerValue(from.IsRemote)
	m.RemoteAdmin = flex.FlattenStringPointer(from.RemoteAdmin)
	m.LocalAdmin = flex.FlattenStringPointer(from.LocalAdmin)
}
