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

type MembercloudsyncModel struct {
	Ref              types.String `tfsdk:"ref"`
	CloudSyncEnabled types.Bool   `tfsdk:"cloud_sync_enabled"`
	HostName         types.String `tfsdk:"host_name"`
}

var MembercloudsyncAttrTypes = map[string]attr.Type{
	"ref":                types.StringType,
	"cloud_sync_enabled": types.BoolType,
	"host_name":          types.StringType,
}

var MembercloudsyncResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"cloud_sync_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Option to enable/disable Cloud Sync.",
	},
	"host_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host name of the parent Member",
	},
}

func ExpandMembercloudsync(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.Membercloudsync {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MembercloudsyncModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MembercloudsyncModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.Membercloudsync {
	if m == nil {
		return nil
	}
	to := &grid.Membercloudsync{
		Ref:              flex.ExpandStringPointer(m.Ref),
		CloudSyncEnabled: flex.ExpandBoolPointer(m.CloudSyncEnabled),
	}
	return to
}

func FlattenMembercloudsync(ctx context.Context, from *grid.Membercloudsync, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MembercloudsyncAttrTypes)
	}
	m := MembercloudsyncModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MembercloudsyncAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MembercloudsyncModel) Flatten(ctx context.Context, from *grid.Membercloudsync, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MembercloudsyncModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.CloudSyncEnabled = types.BoolPointerValue(from.CloudSyncEnabled)
	m.HostName = flex.FlattenStringPointer(from.HostName)
}
