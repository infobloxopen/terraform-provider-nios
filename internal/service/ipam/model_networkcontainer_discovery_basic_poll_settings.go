package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type NetworkcontainerDiscoveryBasicPollSettingsModel struct {
	PortScanning                            types.Bool   `tfsdk:"port_scanning"`
	DeviceProfile                           types.Bool   `tfsdk:"device_profile"`
	SnmpCollection                          types.Bool   `tfsdk:"snmp_collection"`
	CliCollection                           types.Bool   `tfsdk:"cli_collection"`
	NetbiosScanning                         types.Bool   `tfsdk:"netbios_scanning"`
	CompletePingSweep                       types.Bool   `tfsdk:"complete_ping_sweep"`
	SmartSubnetPingSweep                    types.Bool   `tfsdk:"smart_subnet_ping_sweep"`
	AutoArpRefreshBeforeSwitchPortPolling   types.Bool   `tfsdk:"auto_arp_refresh_before_switch_port_polling"`
	SwitchPortDataCollectionPolling         types.String `tfsdk:"switch_port_data_collection_polling"`
	SwitchPortDataCollectionPollingSchedule types.Object `tfsdk:"switch_port_data_collection_polling_schedule"`
	SwitchPortDataCollectionPollingInterval types.Int64  `tfsdk:"switch_port_data_collection_polling_interval"`
	CredentialGroup                         types.String `tfsdk:"credential_group"`
	PollingFrequencyModifier                types.String `tfsdk:"polling_frequency_modifier"`
	UseGlobalPollingFrequencyModifier       types.Bool   `tfsdk:"use_global_polling_frequency_modifier"`
}

var NetworkcontainerDiscoveryBasicPollSettingsAttrTypes = map[string]attr.Type{
	"port_scanning":                                types.BoolType,
	"device_profile":                               types.BoolType,
	"snmp_collection":                              types.BoolType,
	"cli_collection":                               types.BoolType,
	"netbios_scanning":                             types.BoolType,
	"complete_ping_sweep":                          types.BoolType,
	"smart_subnet_ping_sweep":                      types.BoolType,
	"auto_arp_refresh_before_switch_port_polling":  types.BoolType,
	"switch_port_data_collection_polling":          types.StringType,
	"switch_port_data_collection_polling_schedule": types.ObjectType{AttrTypes: NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingScheduleAttrTypes},
	"switch_port_data_collection_polling_interval": types.Int64Type,
	"credential_group":                             types.StringType,
	"polling_frequency_modifier":                   types.StringType,
	"use_global_polling_frequency_modifier":        types.BoolType,
}

var NetworkcontainerDiscoveryBasicPollSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"port_scanning": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether port scanning is enabled or not.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"device_profile": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether device profile is enabled or not.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"snmp_collection": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether SNMP collection is enabled or not.",
		Computed:            true,
		Default:             booldefault.StaticBool(true),
	},
	"cli_collection": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether CLI collection is enabled or not.",
		Computed:            true,
		Default:             booldefault.StaticBool(true),
	},
	"netbios_scanning": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether netbios scanning is enabled or not.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"complete_ping_sweep": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether complete ping sweep is enabled or not.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"smart_subnet_ping_sweep": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether smart subnet ping sweep is enabled or not.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"auto_arp_refresh_before_switch_port_polling": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether auto ARP refresh before switch port polling is enabled or not.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"switch_port_data_collection_polling": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "A switch port data collection polling mode.",
		Computed:            true,
		Default:             stringdefault.StaticString("PERIODIC"),
	},
	"switch_port_data_collection_polling_schedule": schema.SingleNestedAttribute{
		Attributes: NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingScheduleResourceSchemaAttributes,
		Optional:   true,
	},
	"switch_port_data_collection_polling_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Indicates the interval for switch port data collection polling.",
		Computed:            true,
		Default:             int64default.StaticInt64(3600),
	},
	"credential_group": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Credential group.",
		Computed:            true,
		Default:             stringdefault.StaticString("default"),
	},
	"polling_frequency_modifier": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Polling Frequency Modifier.",
		Computed:            true,
		Default:             stringdefault.StaticString("1"),
	},
	"use_global_polling_frequency_modifier": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use Global Polling Frequency Modifier.",
		Computed:            true,
		Default:             booldefault.StaticBool(true),
	},
}

func ExpandNetworkcontainerDiscoveryBasicPollSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkcontainerDiscoveryBasicPollSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainerDiscoveryBasicPollSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkcontainerDiscoveryBasicPollSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkcontainerDiscoveryBasicPollSettings {
	if m == nil {
		return nil
	}
	to := &ipam.NetworkcontainerDiscoveryBasicPollSettings{
		PortScanning:                            flex.ExpandBoolPointer(m.PortScanning),
		DeviceProfile:                           flex.ExpandBoolPointer(m.DeviceProfile),
		SnmpCollection:                          flex.ExpandBoolPointer(m.SnmpCollection),
		CliCollection:                           flex.ExpandBoolPointer(m.CliCollection),
		NetbiosScanning:                         flex.ExpandBoolPointer(m.NetbiosScanning),
		CompletePingSweep:                       flex.ExpandBoolPointer(m.CompletePingSweep),
		SmartSubnetPingSweep:                    flex.ExpandBoolPointer(m.SmartSubnetPingSweep),
		AutoArpRefreshBeforeSwitchPortPolling:   flex.ExpandBoolPointer(m.AutoArpRefreshBeforeSwitchPortPolling),
		SwitchPortDataCollectionPolling:         flex.ExpandStringPointer(m.SwitchPortDataCollectionPolling),
		SwitchPortDataCollectionPollingSchedule: ExpandNetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule(ctx, m.SwitchPortDataCollectionPollingSchedule, diags),
		SwitchPortDataCollectionPollingInterval: flex.ExpandInt64Pointer(m.SwitchPortDataCollectionPollingInterval),
		CredentialGroup:                         flex.ExpandStringPointer(m.CredentialGroup),
		PollingFrequencyModifier:                flex.ExpandStringPointer(m.PollingFrequencyModifier),
		UseGlobalPollingFrequencyModifier:       flex.ExpandBoolPointer(m.UseGlobalPollingFrequencyModifier),
	}
	return to
}

func FlattenNetworkcontainerDiscoveryBasicPollSettings(ctx context.Context, from *ipam.NetworkcontainerDiscoveryBasicPollSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerDiscoveryBasicPollSettingsAttrTypes)
	}
	m := NetworkcontainerDiscoveryBasicPollSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkcontainerDiscoveryBasicPollSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkcontainerDiscoveryBasicPollSettingsModel) Flatten(ctx context.Context, from *ipam.NetworkcontainerDiscoveryBasicPollSettings, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkcontainerDiscoveryBasicPollSettingsModel{}
	}
	m.PortScanning = types.BoolPointerValue(from.PortScanning)
	m.DeviceProfile = types.BoolPointerValue(from.DeviceProfile)
	m.SnmpCollection = types.BoolPointerValue(from.SnmpCollection)
	m.CliCollection = types.BoolPointerValue(from.CliCollection)
	m.NetbiosScanning = types.BoolPointerValue(from.NetbiosScanning)
	m.CompletePingSweep = types.BoolPointerValue(from.CompletePingSweep)
	m.SmartSubnetPingSweep = types.BoolPointerValue(from.SmartSubnetPingSweep)
	m.AutoArpRefreshBeforeSwitchPortPolling = types.BoolPointerValue(from.AutoArpRefreshBeforeSwitchPortPolling)
	m.SwitchPortDataCollectionPolling = flex.FlattenStringPointer(from.SwitchPortDataCollectionPolling)
	m.SwitchPortDataCollectionPollingSchedule = FlattenNetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule(ctx, from.SwitchPortDataCollectionPollingSchedule, diags)
	m.SwitchPortDataCollectionPollingInterval = flex.FlattenInt64Pointer(from.SwitchPortDataCollectionPollingInterval)
	m.CredentialGroup = flex.FlattenStringPointer(from.CredentialGroup)
	m.PollingFrequencyModifier = flex.FlattenStringPointer(from.PollingFrequencyModifier)
	m.UseGlobalPollingFrequencyModifier = types.BoolPointerValue(from.UseGlobalPollingFrequencyModifier)
}
