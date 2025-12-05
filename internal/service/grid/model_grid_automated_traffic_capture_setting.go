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

type GridAutomatedTrafficCaptureSettingModel struct {
	TrafficCaptureEnable    types.Bool   `tfsdk:"traffic_capture_enable"`
	Destination             types.String `tfsdk:"destination"`
	Duration                types.Int64  `tfsdk:"duration"`
	IncludeSupportBundle    types.Bool   `tfsdk:"include_support_bundle"`
	KeepLocalCopy           types.Bool   `tfsdk:"keep_local_copy"`
	DestinationHost         types.String `tfsdk:"destination_host"`
	TrafficCaptureDirectory types.String `tfsdk:"traffic_capture_directory"`
	SupportBundleDirectory  types.String `tfsdk:"support_bundle_directory"`
	Username                types.String `tfsdk:"username"`
	Password                types.String `tfsdk:"password"`
}

var GridAutomatedTrafficCaptureSettingAttrTypes = map[string]attr.Type{
	"traffic_capture_enable":    types.BoolType,
	"destination":               types.StringType,
	"duration":                  types.Int64Type,
	"include_support_bundle":    types.BoolType,
	"keep_local_copy":           types.BoolType,
	"destination_host":          types.StringType,
	"traffic_capture_directory": types.StringType,
	"support_bundle_directory":  types.StringType,
	"username":                  types.StringType,
	"password":                  types.StringType,
}

var GridAutomatedTrafficCaptureSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"traffic_capture_enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable automated traffic capture based on monitoring thresholds.",
	},
	"destination": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Destination of traffic capture files. Save traffic capture locally or upload to remote server using FTP or SCP.",
	},
	"duration": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The time interval on which traffic will be captured(in sec).",
	},
	"include_support_bundle": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable automatic download for support bundle.",
	},
	"keep_local_copy": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Save traffic capture files locally.",
	},
	"destination_host": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "IP Address of the destination host.",
	},
	"traffic_capture_directory": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Directory to store the traffic capture files on the remote server.",
	},
	"support_bundle_directory": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Directory to store the support bundle on the remote server.",
	},
	"username": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "User name for accessing the FTP/SCP server.",
	},
	"password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Password for accessing the FTP/SCP server. This field is not readable.",
	},
}

func ExpandGridAutomatedTrafficCaptureSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridAutomatedTrafficCaptureSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridAutomatedTrafficCaptureSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridAutomatedTrafficCaptureSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridAutomatedTrafficCaptureSetting {
	if m == nil {
		return nil
	}
	to := &grid.GridAutomatedTrafficCaptureSetting{
		TrafficCaptureEnable:    flex.ExpandBoolPointer(m.TrafficCaptureEnable),
		Destination:             flex.ExpandStringPointer(m.Destination),
		Duration:                flex.ExpandInt64Pointer(m.Duration),
		IncludeSupportBundle:    flex.ExpandBoolPointer(m.IncludeSupportBundle),
		KeepLocalCopy:           flex.ExpandBoolPointer(m.KeepLocalCopy),
		DestinationHost:         flex.ExpandStringPointer(m.DestinationHost),
		TrafficCaptureDirectory: flex.ExpandStringPointer(m.TrafficCaptureDirectory),
		SupportBundleDirectory:  flex.ExpandStringPointer(m.SupportBundleDirectory),
		Username:                flex.ExpandStringPointer(m.Username),
		Password:                flex.ExpandStringPointer(m.Password),
	}
	return to
}

func FlattenGridAutomatedTrafficCaptureSetting(ctx context.Context, from *grid.GridAutomatedTrafficCaptureSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridAutomatedTrafficCaptureSettingAttrTypes)
	}
	m := GridAutomatedTrafficCaptureSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridAutomatedTrafficCaptureSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridAutomatedTrafficCaptureSettingModel) Flatten(ctx context.Context, from *grid.GridAutomatedTrafficCaptureSetting, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridAutomatedTrafficCaptureSettingModel{}
	}
	m.TrafficCaptureEnable = types.BoolPointerValue(from.TrafficCaptureEnable)
	m.Destination = flex.FlattenStringPointer(from.Destination)
	m.Duration = flex.FlattenInt64Pointer(from.Duration)
	m.IncludeSupportBundle = types.BoolPointerValue(from.IncludeSupportBundle)
	m.KeepLocalCopy = types.BoolPointerValue(from.KeepLocalCopy)
	m.DestinationHost = flex.FlattenStringPointer(from.DestinationHost)
	m.TrafficCaptureDirectory = flex.FlattenStringPointer(from.TrafficCaptureDirectory)
	m.SupportBundleDirectory = flex.FlattenStringPointer(from.SupportBundleDirectory)
	m.Username = flex.FlattenStringPointer(from.Username)
	m.Password = flex.FlattenStringPointer(from.Password)
}
