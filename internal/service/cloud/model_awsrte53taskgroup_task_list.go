package cloud

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/cloud"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type Awsrte53taskgroupTaskListModel struct {
	Name             types.String `tfsdk:"name"`
	Disabled         types.Bool   `tfsdk:"disabled"`
	State            types.String `tfsdk:"state"`
	StateMsg         types.String `tfsdk:"state_msg"`
	Filter           types.String `tfsdk:"filter"`
	ScheduleInterval types.Int64  `tfsdk:"schedule_interval"`
	ScheduleUnits    types.String `tfsdk:"schedule_units"`
	AwsUser          types.String `tfsdk:"aws_user"`
	StatusTimestamp  types.Int64  `tfsdk:"status_timestamp"`
	LastRun          types.Int64  `tfsdk:"last_run"`
	SyncPublicZones  types.Bool   `tfsdk:"sync_public_zones"`
	SyncPrivateZones types.Bool   `tfsdk:"sync_private_zones"`
	ZoneCount        types.Int64  `tfsdk:"zone_count"`
	CredentialsType  types.String `tfsdk:"credentials_type"`
}

var Awsrte53taskgroupTaskListAttrTypes = map[string]attr.Type{
	"name":               types.StringType,
	"disabled":           types.BoolType,
	"state":              types.StringType,
	"state_msg":          types.StringType,
	"filter":             types.StringType,
	"schedule_interval":  types.Int64Type,
	"schedule_units":     types.StringType,
	"aws_user":           types.StringType,
	"status_timestamp":   types.Int64Type,
	"last_run":           types.Int64Type,
	"sync_public_zones":  types.BoolType,
	"sync_private_zones": types.BoolType,
	"zone_count":         types.Int64Type,
	"credentials_type":   types.StringType,
}

var Awsrte53taskgroupTaskListResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile(`^[^\s]([^\s]|.*[^\s])?$`),
				"Should not have leading or trailing whitespace",
			),
		},
		MarkdownDescription: "The name of this task.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Indicates if the task is enabled or disabled.",
	},
	"state": schema.StringAttribute{
		Computed: true,
		Validators: []validator.String{
			stringvalidator.OneOf("CANCELED", "COMPLETED", "FAILED", "IDLE", "PARTIALLY_COMPLETED", "READY", "RUNNING"),
		},
		MarkdownDescription: "Indicate the sync status of this task.",
	},
	"state_msg": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "State message for the task.",
	},
	"filter": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("*"),
		MarkdownDescription: "Filter for this task.",
	},
	"schedule_interval": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(60),
		MarkdownDescription: "Periodic interval for this task.",
	},
	"schedule_units": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("MINS"),
		Validators: []validator.String{
			stringvalidator.OneOf("DAYS", "HOURS", "MINS"),
		},
		MarkdownDescription: "Units for the schedule interval.",
	},
	"aws_user": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Reference to associated AWS user whose credentials are to be used for this task.",
	},
	"status_timestamp": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The timestamp when the last state was logged.",
	},
	"last_run": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The timestamp when the task was started last.",
	},
	"sync_public_zones": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Indicates whether public zones are synchronized.",
	},
	"sync_private_zones": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Indicates whether private zones are synchronized.",
	},
	"zone_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of zones synchronized by this task.",
	},
	"credentials_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("DIRECT"),
		Validators: []validator.String{
			stringvalidator.OneOf("DIRECT", "INDIRECT"),
		},
		MarkdownDescription: "Credentials type used for connecting to the cloud management platform.",
	},
}

func ExpandAwsrte53taskgroupTaskList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *cloud.Awsrte53taskgroupTaskList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Awsrte53taskgroupTaskListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Awsrte53taskgroupTaskListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *cloud.Awsrte53taskgroupTaskList {
	if m == nil {
		return nil
	}
	to := &cloud.Awsrte53taskgroupTaskList{
		Name:             flex.ExpandStringPointer(m.Name),
		Disabled:         flex.ExpandBoolPointer(m.Disabled),
		Filter:           flex.ExpandStringPointer(m.Filter),
		ScheduleInterval: flex.ExpandInt64Pointer(m.ScheduleInterval),
		ScheduleUnits:    flex.ExpandStringPointer(m.ScheduleUnits),
		AwsUser:          flex.ExpandStringPointer(m.AwsUser),
		SyncPublicZones:  flex.ExpandBoolPointer(m.SyncPublicZones),
		SyncPrivateZones: flex.ExpandBoolPointer(m.SyncPrivateZones),
		CredentialsType:  flex.ExpandStringPointer(m.CredentialsType),
	}
	return to
}

func FlattenAwsrte53taskgroupTaskList(ctx context.Context, from *cloud.Awsrte53taskgroupTaskList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Awsrte53taskgroupTaskListAttrTypes)
	}
	m := Awsrte53taskgroupTaskListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Awsrte53taskgroupTaskListAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Awsrte53taskgroupTaskListModel) Flatten(ctx context.Context, from *cloud.Awsrte53taskgroupTaskList, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Awsrte53taskgroupTaskListModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Disabled = types.BoolPointerValue(from.Disabled)
	m.State = flex.FlattenStringPointer(from.State)
	m.StateMsg = flex.FlattenStringPointer(from.StateMsg)
	m.Filter = flex.FlattenStringPointer(from.Filter)
	m.ScheduleInterval = flex.FlattenInt64Pointer(from.ScheduleInterval)
	m.ScheduleUnits = flex.FlattenStringPointer(from.ScheduleUnits)
	m.AwsUser = flex.FlattenStringPointer(from.AwsUser)
	m.StatusTimestamp = flex.FlattenInt64Pointer(from.StatusTimestamp)
	m.LastRun = flex.FlattenInt64Pointer(from.LastRun)
	m.SyncPublicZones = types.BoolPointerValue(from.SyncPublicZones)
	m.SyncPrivateZones = types.BoolPointerValue(from.SyncPrivateZones)
	m.ZoneCount = flex.FlattenInt64Pointer(from.ZoneCount)
	m.CredentialsType = flex.FlattenStringPointer(from.CredentialsType)
}
