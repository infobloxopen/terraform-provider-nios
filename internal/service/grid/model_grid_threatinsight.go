package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type GridThreatinsightModel struct {
	Ref                              types.String `tfsdk:"ref"`
	Uuid                             types.String `tfsdk:"uuid"`
	AllowlistUpdatePolicy            types.String `tfsdk:"allowlist_update_policy"`
	ConfigureDomainCollapsing        types.Bool   `tfsdk:"configure_domain_collapsing"`
	CurrentAllowlist                 types.String `tfsdk:"current_allowlist"`
	CurrentModuleset                 types.String `tfsdk:"current_moduleset"`
	DnsTunnelBlockListRpzZones       types.List   `tfsdk:"dns_tunnel_block_list_rpz_zones"`
	DomainCollapsingLevel            types.Int64  `tfsdk:"domain_collapsing_level"`
	EnableAllowlistAutoDownload      types.Bool   `tfsdk:"enable_allowlist_auto_download"`
	EnableAllowlistScheduledDownload types.Bool   `tfsdk:"enable_allowlist_scheduled_download"`
	EnableAutoDownload               types.Bool   `tfsdk:"enable_auto_download"`
	EnableScheduledDownload          types.Bool   `tfsdk:"enable_scheduled_download"`
	LastAllowlistUpdateTime          types.Int64  `tfsdk:"last_allowlist_update_time"`
	LastAllowlistUpdateVersion       types.String `tfsdk:"last_allowlist_update_version"`
	LastCheckedForAllowlistUpdate    types.Int64  `tfsdk:"last_checked_for_allowlist_update"`
	LastCheckedForPackageUpdate      types.Int64  `tfsdk:"last_checked_for_package_update"`
	LastCheckedForUpdate             types.Int64  `tfsdk:"last_checked_for_update"`
	LastModuleUpdateTime             types.Int64  `tfsdk:"last_module_update_time"`
	LastModuleUpdateVersion          types.String `tfsdk:"last_module_update_version"`
	LastUpdatedPackageVersion        types.String `tfsdk:"last_updated_package_version"`
	ModuleUpdatePolicy               types.String `tfsdk:"module_update_policy"`
	Name                             types.String `tfsdk:"name"`
	ScheduledAllowlistDownload       types.Object `tfsdk:"scheduled_allowlist_download"`
	ScheduledDownload                types.Object `tfsdk:"scheduled_download"`
}

var GridThreatinsightAttrTypes = map[string]attr.Type{
	"ref":                                 types.StringType,
	"uuid":                                types.StringType,
	"allowlist_update_policy":             types.StringType,
	"configure_domain_collapsing":         types.BoolType,
	"current_allowlist":                   types.StringType,
	"current_moduleset":                   types.StringType,
	"dns_tunnel_block_list_rpz_zones":     types.ListType{ElemType: types.StringType},
	"domain_collapsing_level":             types.Int64Type,
	"enable_allowlist_auto_download":      types.BoolType,
	"enable_allowlist_scheduled_download": types.BoolType,
	"enable_auto_download":                types.BoolType,
	"enable_scheduled_download":           types.BoolType,
	"last_allowlist_update_time":          types.Int64Type,
	"last_allowlist_update_version":       types.StringType,
	"last_checked_for_allowlist_update":   types.Int64Type,
	"last_checked_for_package_update":     types.Int64Type,
	"last_checked_for_update":             types.Int64Type,
	"last_module_update_time":             types.Int64Type,
	"last_module_update_version":          types.StringType,
	"last_updated_package_version":        types.StringType,
	"module_update_policy":                types.StringType,
	"name":                                types.StringType,
	"scheduled_allowlist_download":        types.ObjectType{AttrTypes: GridThreatinsightScheduledAllowlistDownloadAttrTypes},
	"scheduled_download":                  types.ObjectType{AttrTypes: GridThreatinsightScheduledDownloadAttrTypes},
}

var GridThreatinsightResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The UUID of the object.",
	},
	"allowlist_update_policy": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "allowlist update policy (manual or automatic)",
	},
	"configure_domain_collapsing": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Disable domain collapsing at grid level",
	},
	"current_allowlist": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Grid allowlist.",
	},
	"current_moduleset": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The current threat insight module set.",
	},
	"dns_tunnel_block_list_rpz_zones": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The list of response policy zones for DNS tunnelling requests.",
	},
	"domain_collapsing_level": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Level of domain collapsing",
	},
	"enable_allowlist_auto_download": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Indicates whether auto download service is enabled",
	},
	"enable_allowlist_scheduled_download": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Indicates whether the custom scheduled settings for auto download is enabled. If false then default frequency is once per 24 hours",
	},
	"enable_auto_download": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the automatic threat insight module set download is enabled.",
	},
	"enable_scheduled_download": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the scheduled download of the threat insight module set is enabled.",
	},
	"last_allowlist_update_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The last update time for the threat insight allowlist.",
	},
	"last_allowlist_update_version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The version number of the last updated threat insight allowlist.",
	},
	"last_checked_for_allowlist_update": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Timestamp of last checked allowlist",
	},
	"last_checked_for_package_update": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The last update time for the threat analytics moduleset package.",
	},
	"last_checked_for_update": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The last time when the threat insight module set was checked for the update.",
	},
	"last_module_update_time": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The last update time for the threat insight module set.",
	},
	"last_module_update_version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The version number of the last updated threat insight module set.",
	},
	"last_updated_package_version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The version number of the last updated Moduleset package.",
	},
	"module_update_policy": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The update policy for the threat insight module set.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Grid name.",
	},
	"scheduled_allowlist_download": schema.SingleNestedAttribute{
		Attributes: GridThreatinsightScheduledAllowlistDownloadResourceSchemaAttributes,
		Optional:   true,
	},
	"scheduled_download": schema.SingleNestedAttribute{
		Attributes: GridThreatinsightScheduledDownloadResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandGridThreatinsight(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.GridThreatinsight {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridThreatinsightModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *GridThreatinsightModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.GridThreatinsight {
	if m == nil {
		return nil
	}
	to := &grid.GridThreatinsight{
		Ref:                              flex.ExpandStringPointer(m.Ref),
		Uuid:                             flex.ExpandStringPointer(m.Uuid),
		AllowlistUpdatePolicy:            flex.ExpandStringPointer(m.AllowlistUpdatePolicy),
		ConfigureDomainCollapsing:        flex.ExpandBoolPointer(m.ConfigureDomainCollapsing),
		DnsTunnelBlockListRpzZones:       flex.ExpandFrameworkListString(ctx, m.DnsTunnelBlockListRpzZones, diags),
		DomainCollapsingLevel:            flex.ExpandInt64Pointer(m.DomainCollapsingLevel),
		EnableAllowlistAutoDownload:      flex.ExpandBoolPointer(m.EnableAllowlistAutoDownload),
		EnableAllowlistScheduledDownload: flex.ExpandBoolPointer(m.EnableAllowlistScheduledDownload),
		EnableAutoDownload:               flex.ExpandBoolPointer(m.EnableAutoDownload),
		EnableScheduledDownload:          flex.ExpandBoolPointer(m.EnableScheduledDownload),
		ModuleUpdatePolicy:               flex.ExpandStringPointer(m.ModuleUpdatePolicy),
		ScheduledAllowlistDownload:       ExpandGridThreatinsightScheduledAllowlistDownload(ctx, m.ScheduledAllowlistDownload, diags),
		ScheduledDownload:                ExpandGridThreatinsightScheduledDownload(ctx, m.ScheduledDownload, diags),
	}
	return to
}

func FlattenGridThreatinsight(ctx context.Context, from *grid.GridThreatinsight, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridThreatinsightAttrTypes)
	}
	m := GridThreatinsightModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridThreatinsightAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *GridThreatinsightModel) Flatten(ctx context.Context, from *grid.GridThreatinsight, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = GridThreatinsightModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.AllowlistUpdatePolicy = flex.FlattenStringPointer(from.AllowlistUpdatePolicy)
	m.ConfigureDomainCollapsing = types.BoolPointerValue(from.ConfigureDomainCollapsing)
	m.CurrentAllowlist = flex.FlattenStringPointer(from.CurrentAllowlist)
	m.CurrentModuleset = flex.FlattenStringPointer(from.CurrentModuleset)
	m.DnsTunnelBlockListRpzZones = flex.FlattenFrameworkListString(ctx, from.DnsTunnelBlockListRpzZones, diags)
	m.DomainCollapsingLevel = flex.FlattenInt64Pointer(from.DomainCollapsingLevel)
	m.EnableAllowlistAutoDownload = types.BoolPointerValue(from.EnableAllowlistAutoDownload)
	m.EnableAllowlistScheduledDownload = types.BoolPointerValue(from.EnableAllowlistScheduledDownload)
	m.EnableAutoDownload = types.BoolPointerValue(from.EnableAutoDownload)
	m.EnableScheduledDownload = types.BoolPointerValue(from.EnableScheduledDownload)
	m.LastAllowlistUpdateTime = flex.FlattenInt64Pointer(from.LastAllowlistUpdateTime)
	m.LastAllowlistUpdateVersion = flex.FlattenStringPointer(from.LastAllowlistUpdateVersion)
	m.LastCheckedForAllowlistUpdate = flex.FlattenInt64Pointer(from.LastCheckedForAllowlistUpdate)
	m.LastCheckedForPackageUpdate = flex.FlattenInt64Pointer(from.LastCheckedForPackageUpdate)
	m.LastCheckedForUpdate = flex.FlattenInt64Pointer(from.LastCheckedForUpdate)
	m.LastModuleUpdateTime = flex.FlattenInt64Pointer(from.LastModuleUpdateTime)
	m.LastModuleUpdateVersion = flex.FlattenStringPointer(from.LastModuleUpdateVersion)
	m.LastUpdatedPackageVersion = flex.FlattenStringPointer(from.LastUpdatedPackageVersion)
	m.ModuleUpdatePolicy = flex.FlattenStringPointer(from.ModuleUpdatePolicy)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.ScheduledAllowlistDownload = FlattenGridThreatinsightScheduledAllowlistDownload(ctx, from.ScheduledAllowlistDownload, diags)
	m.ScheduledDownload = FlattenGridThreatinsightScheduledDownload(ctx, from.ScheduledDownload, diags)
}
