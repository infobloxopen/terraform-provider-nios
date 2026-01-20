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

type MemberDnsLoggingCategoriesModel struct {
	LogDtcGslb        types.Bool `tfsdk:"log_dtc_gslb"`
	LogDtcHealth      types.Bool `tfsdk:"log_dtc_health"`
	LogGeneral        types.Bool `tfsdk:"log_general"`
	LogClient         types.Bool `tfsdk:"log_client"`
	LogConfig         types.Bool `tfsdk:"log_config"`
	LogDatabase       types.Bool `tfsdk:"log_database"`
	LogDnssec         types.Bool `tfsdk:"log_dnssec"`
	LogLameServers    types.Bool `tfsdk:"log_lame_servers"`
	LogNetwork        types.Bool `tfsdk:"log_network"`
	LogNotify         types.Bool `tfsdk:"log_notify"`
	LogQueries        types.Bool `tfsdk:"log_queries"`
	LogQueryRewrite   types.Bool `tfsdk:"log_query_rewrite"`
	LogResponses      types.Bool `tfsdk:"log_responses"`
	LogResolver       types.Bool `tfsdk:"log_resolver"`
	LogSecurity       types.Bool `tfsdk:"log_security"`
	LogUpdate         types.Bool `tfsdk:"log_update"`
	LogXferIn         types.Bool `tfsdk:"log_xfer_in"`
	LogXferOut        types.Bool `tfsdk:"log_xfer_out"`
	LogUpdateSecurity types.Bool `tfsdk:"log_update_security"`
	LogRateLimit      types.Bool `tfsdk:"log_rate_limit"`
	LogRpz            types.Bool `tfsdk:"log_rpz"`
}

var MemberDnsLoggingCategoriesAttrTypes = map[string]attr.Type{
	"log_dtc_gslb":        types.BoolType,
	"log_dtc_health":      types.BoolType,
	"log_general":         types.BoolType,
	"log_client":          types.BoolType,
	"log_config":          types.BoolType,
	"log_database":        types.BoolType,
	"log_dnssec":          types.BoolType,
	"log_lame_servers":    types.BoolType,
	"log_network":         types.BoolType,
	"log_notify":          types.BoolType,
	"log_queries":         types.BoolType,
	"log_query_rewrite":   types.BoolType,
	"log_responses":       types.BoolType,
	"log_resolver":        types.BoolType,
	"log_security":        types.BoolType,
	"log_update":          types.BoolType,
	"log_xfer_in":         types.BoolType,
	"log_xfer_out":        types.BoolType,
	"log_update_security": types.BoolType,
	"log_rate_limit":      types.BoolType,
	"log_rpz":             types.BoolType,
}

var MemberDnsLoggingCategoriesResourceSchemaAttributes = map[string]schema.Attribute{
	"log_dtc_gslb": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the DTC GSLB activity is captured or not.",
	},
	"log_dtc_health": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the DTC health monitoring information is captured or not.",
	},
	"log_general": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the BIND messages that are not specifically classified are captured or not.",
	},
	"log_client": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the client requests are captured or not.",
	},
	"log_config": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the configuration file parsing is captured or not.",
	},
	"log_database": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the BIND's internal database processes are captured or not.",
	},
	"log_dnssec": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the DNSSEC-signed responses are captured or not.",
	},
	"log_lame_servers": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the bad delegation instances are captured or not.",
	},
	"log_network": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the network operation messages are captured or not.",
	},
	"log_notify": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the asynchronous zone change notification messages are captured or not.",
	},
	"log_queries": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the query messages are captured or not.",
	},
	"log_query_rewrite": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the query rewrite messages are captured or not.",
	},
	"log_responses": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the response messages are captured or not.",
	},
	"log_resolver": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the DNS resolution instances, including recursive queries from resolvers are captured or not.",
	},
	"log_security": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the approved and denied requests are captured or not.",
	},
	"log_update": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the dynamic update instances are captured or not.",
	},
	"log_xfer_in": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the zone transfer messages from the remote name servers to the appliance are captured or not.",
	},
	"log_xfer_out": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the zone transfer messages from the Infoblox appliance to remote name servers are captured or not.",
	},
	"log_update_security": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the security update messages are captured or not.",
	},
	"log_rate_limit": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the rate limit messages are captured or not.",
	},
	"log_rpz": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the Response Policy Zone messages are captured or not.",
	},
}

func ExpandMemberDnsLoggingCategories(ctx context.Context, o types.Object, diags *diag.Diagnostics) *grid.MemberDnsLoggingCategories {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m MemberDnsLoggingCategoriesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *MemberDnsLoggingCategoriesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *grid.MemberDnsLoggingCategories {
	if m == nil {
		return nil
	}
	to := &grid.MemberDnsLoggingCategories{
		LogDtcGslb:        flex.ExpandBoolPointer(m.LogDtcGslb),
		LogDtcHealth:      flex.ExpandBoolPointer(m.LogDtcHealth),
		LogGeneral:        flex.ExpandBoolPointer(m.LogGeneral),
		LogClient:         flex.ExpandBoolPointer(m.LogClient),
		LogConfig:         flex.ExpandBoolPointer(m.LogConfig),
		LogDatabase:       flex.ExpandBoolPointer(m.LogDatabase),
		LogDnssec:         flex.ExpandBoolPointer(m.LogDnssec),
		LogLameServers:    flex.ExpandBoolPointer(m.LogLameServers),
		LogNetwork:        flex.ExpandBoolPointer(m.LogNetwork),
		LogNotify:         flex.ExpandBoolPointer(m.LogNotify),
		LogQueries:        flex.ExpandBoolPointer(m.LogQueries),
		LogQueryRewrite:   flex.ExpandBoolPointer(m.LogQueryRewrite),
		LogResponses:      flex.ExpandBoolPointer(m.LogResponses),
		LogResolver:       flex.ExpandBoolPointer(m.LogResolver),
		LogSecurity:       flex.ExpandBoolPointer(m.LogSecurity),
		LogUpdate:         flex.ExpandBoolPointer(m.LogUpdate),
		LogXferIn:         flex.ExpandBoolPointer(m.LogXferIn),
		LogXferOut:        flex.ExpandBoolPointer(m.LogXferOut),
		LogUpdateSecurity: flex.ExpandBoolPointer(m.LogUpdateSecurity),
		LogRateLimit:      flex.ExpandBoolPointer(m.LogRateLimit),
		LogRpz:            flex.ExpandBoolPointer(m.LogRpz),
	}
	return to
}

func FlattenMemberDnsLoggingCategories(ctx context.Context, from *grid.MemberDnsLoggingCategories, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(MemberDnsLoggingCategoriesAttrTypes)
	}
	m := MemberDnsLoggingCategoriesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, MemberDnsLoggingCategoriesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *MemberDnsLoggingCategoriesModel) Flatten(ctx context.Context, from *grid.MemberDnsLoggingCategories, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = MemberDnsLoggingCategoriesModel{}
	}
	m.LogDtcGslb = types.BoolPointerValue(from.LogDtcGslb)
	m.LogDtcHealth = types.BoolPointerValue(from.LogDtcHealth)
	m.LogGeneral = types.BoolPointerValue(from.LogGeneral)
	m.LogClient = types.BoolPointerValue(from.LogClient)
	m.LogConfig = types.BoolPointerValue(from.LogConfig)
	m.LogDatabase = types.BoolPointerValue(from.LogDatabase)
	m.LogDnssec = types.BoolPointerValue(from.LogDnssec)
	m.LogLameServers = types.BoolPointerValue(from.LogLameServers)
	m.LogNetwork = types.BoolPointerValue(from.LogNetwork)
	m.LogNotify = types.BoolPointerValue(from.LogNotify)
	m.LogQueries = types.BoolPointerValue(from.LogQueries)
	m.LogQueryRewrite = types.BoolPointerValue(from.LogQueryRewrite)
	m.LogResponses = types.BoolPointerValue(from.LogResponses)
	m.LogResolver = types.BoolPointerValue(from.LogResolver)
	m.LogSecurity = types.BoolPointerValue(from.LogSecurity)
	m.LogUpdate = types.BoolPointerValue(from.LogUpdate)
	m.LogXferIn = types.BoolPointerValue(from.LogXferIn)
	m.LogXferOut = types.BoolPointerValue(from.LogXferOut)
	m.LogUpdateSecurity = types.BoolPointerValue(from.LogUpdateSecurity)
	m.LogRateLimit = types.BoolPointerValue(from.LogRateLimit)
	m.LogRpz = types.BoolPointerValue(from.LogRpz)
}
