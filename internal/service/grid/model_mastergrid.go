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

type MastergridModel struct {
	Ref                 types.String `tfsdk:"ref"`
	Uuid                types.String `tfsdk:"uuid"`
	Address             types.String `tfsdk:"address"`
	ConnectionDisabled  types.Bool   `tfsdk:"connection_disabled"`
	ConnectionTimestamp types.Int64  `tfsdk:"connection_timestamp"`
	Detached            types.Bool   `tfsdk:"detached"`
	Enable              types.Bool   `tfsdk:"enable"`
	Joined              types.Bool   `tfsdk:"joined"`
	LastEvent           types.String `tfsdk:"last_event"`
	LastEventDetails    types.String `tfsdk:"last_event_details"`
	LastSyncTimestamp   types.Int64  `tfsdk:"last_sync_timestamp"`
	Port                types.Int64  `tfsdk:"port"`
	Status              types.String `tfsdk:"status"`
	UseMgmtPort         types.Bool   `tfsdk:"use_mgmt_port"`
}

var MastergridAttrTypes = map[string]attr.Type{
	"ref":                  types.StringType,
	"uuid":                 types.StringType,
	"address":              types.StringType,
	"connection_disabled":  types.BoolType,
	"connection_timestamp": types.Int64Type,
	"detached":             types.BoolType,
	"enable":               types.BoolType,
	"joined":               types.BoolType,
	"last_event":           types.StringType,
	"last_event_details":   types.StringType,
	"last_sync_timestamp":  types.Int64Type,
	"port":                 types.Int64Type,
	"status":               types.StringType,
	"use_mgmt_port":        types.BoolType,
}

var MastergridResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The unique identifier for the object.",
	},
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The domain name or IP address for the Master Grid.",
	},
	"connection_disabled": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines if the sub-grid is currently disabled.",
	},
	"connection_timestamp": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The timestamp that indicates when the connection to the Master Grid was established.",
	},
	"detached": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "The detached flag for the Master Grid.",
	},
	"enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the Master Grid is enabled.",
	},
	"joined": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "The flag shows if the Grid has joined the Master Grid.",
	},
	"last_event": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Master Grid's last event.",
	},
	"last_event_details": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The details of the Master Grid's last event.",
	},
	"last_sync_timestamp": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The timestamp or the last synchronization operation with the Master Grid.",
	},
	"port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The Master Grid port to which the Grid connects.",
	},
	"status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Master Grid's status.",
	},
	"use_mgmt_port": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "The flag shows if the MGMT port was used to join the Grid.",
	},
}

func ExpandMastergrid(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.Mastergrid {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MastergridModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MastergridModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.Mastergrid {
	if m == nil {
		return nil
	}
	to := &grid.Mastergrid{
		Ref:     flex.ExpandStringPointer(m.Ref),
		Uuid:    flex.ExpandStringPointer(m.Uuid),
		Address: flex.ExpandStringPointer(m.Address),
		Enable:  flex.ExpandBoolPointer(m.Enable),
		Port:    flex.ExpandInt64Pointer(m.Port),
	}
	return to
}

func FlattenMastergrid(ctx context.Context, from *grid.Mastergrid, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MastergridAttrTypes)
	}
	m := MastergridModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MastergridAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MastergridModel) Flatten(ctx context.Context, from *grid.Mastergrid, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MastergridModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.ConnectionDisabled = types.BoolPointerValue(from.ConnectionDisabled)
	m.ConnectionTimestamp = flex.FlattenInt64Pointer(from.ConnectionTimestamp)
	m.Detached = types.BoolPointerValue(from.Detached)
	m.Enable = types.BoolPointerValue(from.Enable)
	m.Joined = types.BoolPointerValue(from.Joined)
	m.LastEvent = flex.FlattenStringPointer(from.LastEvent)
	m.LastEventDetails = flex.FlattenStringPointer(from.LastEventDetails)
	m.LastSyncTimestamp = flex.FlattenInt64Pointer(from.LastSyncTimestamp)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	//m.Status = flex.FlattenStringPointer(from.Status) -> TO DO: absent in 9.1.0
	m.UseMgmtPort = types.BoolPointerValue(from.UseMgmtPort)
}
