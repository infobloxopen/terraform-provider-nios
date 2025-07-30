package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type RecordSrvAwsRte53RecordInfoModel struct {
	AliasTargetDnsName              types.String `tfsdk:"alias_target_dns_name"`
	AliasTargetHostedZoneId         types.String `tfsdk:"alias_target_hosted_zone_id"`
	AliasTargetEvaluateTargetHealth types.Bool   `tfsdk:"alias_target_evaluate_target_health"`
	Failover                        types.String `tfsdk:"failover"`
	GeolocationContinentCode        types.String `tfsdk:"geolocation_continent_code"`
	GeolocationCountryCode          types.String `tfsdk:"geolocation_country_code"`
	GeolocationSubdivisionCode      types.String `tfsdk:"geolocation_subdivision_code"`
	HealthCheckId                   types.String `tfsdk:"health_check_id"`
	Region                          types.String `tfsdk:"region"`
	SetIdentifier                   types.String `tfsdk:"set_identifier"`
	Type                            types.String `tfsdk:"type"`
	Weight                          types.Int64  `tfsdk:"weight"`
}

var RecordSrvAwsRte53RecordInfoAttrTypes = map[string]attr.Type{
	"alias_target_dns_name":               types.StringType,
	"alias_target_hosted_zone_id":         types.StringType,
	"alias_target_evaluate_target_health": types.BoolType,
	"failover":                            types.StringType,
	"geolocation_continent_code":          types.StringType,
	"geolocation_country_code":            types.StringType,
	"geolocation_subdivision_code":        types.StringType,
	"health_check_id":                     types.StringType,
	"region":                              types.StringType,
	"set_identifier":                      types.StringType,
	"type":                                types.StringType,
	"weight":                              types.Int64Type,
}

var RecordSrvAwsRte53RecordInfoResourceSchemaAttributes = map[string]schema.Attribute{
	"alias_target_dns_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "DNS name of the alias target.",
	},
	"alias_target_hosted_zone_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Hosted zone ID of the alias target.",
	},
	"alias_target_evaluate_target_health": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates if Amazon Route 53 evaluates the health of the alias target.",
	},
	"failover": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates whether this is the primary or secondary resource record for Amazon Route 53 failover routing.",
	},
	"geolocation_continent_code": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Continent code for Amazon Route 53 geolocation routing.",
	},
	"geolocation_country_code": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Country code for Amazon Route 53 geolocation routing.",
	},
	"geolocation_subdivision_code": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Subdivision code for Amazon Route 53 geolocation routing.",
	},
	"health_check_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "ID of the health check that Amazon Route 53 performs for this resource record.",
	},
	"region": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Amazon EC2 region where this resource record resides for latency routing.",
	},
	"set_identifier": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "An identifier that differentiates records with the same DNS name and type for weighted, latency, geolocation, and failover routing.",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Type of Amazon Route 53 resource record.",
	},
	"weight": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Value that determines the portion of traffic for this record in weighted routing. The range is from 0 to 255.",
	},
}

func (m *RecordSrvAwsRte53RecordInfoModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordSrvAwsRte53RecordInfo {
	if m == nil {
		return nil
	}
	to := &dns.RecordSrvAwsRte53RecordInfo{}
	return to
}

func FlattenRecordSrvAwsRte53RecordInfo(ctx context.Context, from *dns.RecordSrvAwsRte53RecordInfo, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordSrvAwsRte53RecordInfoAttrTypes)
	}
	m := RecordSrvAwsRte53RecordInfoModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordSrvAwsRte53RecordInfoAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordSrvAwsRte53RecordInfoModel) Flatten(ctx context.Context, from *dns.RecordSrvAwsRte53RecordInfo, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordSrvAwsRte53RecordInfoModel{}
	}
	m.AliasTargetDnsName = flex.FlattenStringPointer(from.AliasTargetDnsName)
	m.AliasTargetHostedZoneId = flex.FlattenStringPointer(from.AliasTargetHostedZoneId)
	m.AliasTargetEvaluateTargetHealth = types.BoolPointerValue(from.AliasTargetEvaluateTargetHealth)
	m.Failover = flex.FlattenStringPointer(from.Failover)
	m.GeolocationContinentCode = flex.FlattenStringPointer(from.GeolocationContinentCode)
	m.GeolocationCountryCode = flex.FlattenStringPointer(from.GeolocationCountryCode)
	m.GeolocationSubdivisionCode = flex.FlattenStringPointer(from.GeolocationSubdivisionCode)
	m.HealthCheckId = flex.FlattenStringPointer(from.HealthCheckId)
	m.Region = flex.FlattenStringPointer(from.Region)
	m.SetIdentifier = flex.FlattenStringPointer(from.SetIdentifier)
	m.Type = flex.FlattenStringPointer(from.Type)
	m.Weight = flex.FlattenInt64Pointer(from.Weight)
}
