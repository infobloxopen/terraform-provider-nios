package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ZoneAuthAwsRte53ZoneInfoModel struct {
	AssociatedVpcs  types.List   `tfsdk:"associated_vpcs"`
	CallerReference types.String `tfsdk:"callerreference"`
	DelegationSetId types.String `tfsdk:"delegation_set_id"`
	HostedZoneId    types.String `tfsdk:"hosted_zone_id"`
	NameServers     types.List   `tfsdk:"name_servers"`
	RecordSetCount  types.Int64  `tfsdk:"record_set_count"`
	Type            types.String `tfsdk:"type"`
}

var ZoneAuthAwsRte53ZoneInfoAttrTypes = map[string]attr.Type{
	"associated_vpcs":   types.ListType{ElemType: types.StringType},
	"callerreference":   types.StringType,
	"delegation_set_id": types.StringType,
	"hosted_zone_id":    types.StringType,
	"name_servers":      types.ListType{ElemType: types.StringType},
	"record_set_count":  types.Int64Type,
	"type":              types.StringType,
}

var ZoneAuthAwsRte53ZoneInfoResourceSchemaAttributes = map[string]schema.Attribute{
	"associated_vpcs": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "List of AWS VPC strings that are associated with this zone.",
	},
	"callerreference": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "User specified caller reference when zone was created.",
	},
	"delegation_set_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "ID of delegation set associated with this zone.",
	},
	"hosted_zone_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "AWS route 53 assigned ID for this zone.",
	},
	"name_servers": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "List of AWS name servers that are authoritative for this domain name.",
	},
	"record_set_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Number of resource record sets in the hosted zone.",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates whether private or public zone.",
	},
}

func ExpandZoneAuthAwsRte53ZoneInfo(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.ZoneAuthAwsRte53ZoneInfo {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneAuthAwsRte53ZoneInfoModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ZoneAuthAwsRte53ZoneInfoModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.ZoneAuthAwsRte53ZoneInfo {
	if m == nil {
		return nil
	}
	to := &dns.ZoneAuthAwsRte53ZoneInfo{}
	return to
}

func FlattenZoneAuthAwsRte53ZoneInfo(ctx context.Context, from *dns.ZoneAuthAwsRte53ZoneInfo, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneAuthAwsRte53ZoneInfoAttrTypes)
	}
	m := ZoneAuthAwsRte53ZoneInfoModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneAuthAwsRte53ZoneInfoAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ZoneAuthAwsRte53ZoneInfoModel) Flatten(ctx context.Context, from *dns.ZoneAuthAwsRte53ZoneInfo, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ZoneAuthAwsRte53ZoneInfoModel{}
	}
	m.AssociatedVpcs = flex.FlattenFrameworkListString(ctx, from.AssociatedVpcs, diags)
	m.CallerReference = flex.FlattenStringPointer(from.CallerReference)
	m.DelegationSetId = flex.FlattenStringPointer(from.DelegationSetId)
	m.HostedZoneId = flex.FlattenStringPointer(from.HostedZoneId)
	m.NameServers = flex.FlattenFrameworkListString(ctx, from.NameServers, diags)
	m.RecordSetCount = flex.FlattenInt64Pointer(from.RecordSetCount)
	m.Type = flex.FlattenStringPointer(from.Type)
}
