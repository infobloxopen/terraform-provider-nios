package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type RecordACloudInfoModel struct {
	AuthorityType  types.String `tfsdk:"authority_type"`
	DelegatedScope types.String `tfsdk:"delegated_scope"`
	MgmtPlatform   types.String `tfsdk:"mgmt_platform"`
	OwnedByAdaptor types.Bool   `tfsdk:"owned_by_adaptor"`
}

var RecordACloudInfoAttrTypes = map[string]attr.Type{
	"authority_type":   types.StringType,
	"delegated_scope":  types.StringType,
	"mgmt_platform":    types.StringType,
	"owned_by_adaptor": types.BoolType,
}

var RecordACloudInfoResourceSchemaAttributes = map[string]schema.Attribute{
	"authority_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Type of authority over the object.",
	},
	"delegated_scope": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Indicates the root of the delegation if delegated_scope is SUBTREE or RECLAIMING. This is not set otherwise.",
	},
	"mgmt_platform": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Indicates the specified cloud management platform.",
	},
	"owned_by_adaptor": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the object was created by the cloud adapter or not.",
	},
}

func ExpandRecordACloudInfo(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.RecordACloudInfo {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordACloudInfoModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RecordACloudInfoModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordACloudInfo {
	if m == nil {
		return nil
	}
	to := &dns.RecordACloudInfo{
		AuthorityType:  flex.ExpandStringPointer(m.AuthorityType),
		DelegatedScope: flex.ExpandStringPointer(m.DelegatedScope),
		MgmtPlatform:   flex.ExpandStringPointer(m.MgmtPlatform),
		OwnedByAdaptor: flex.ExpandBoolPointer(m.OwnedByAdaptor),
	}
	return to
}

func FlattenRecordACloudInfo(ctx context.Context, from *dns.RecordACloudInfo, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordACloudInfoAttrTypes)
	}
	m := RecordACloudInfoModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordACloudInfoAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordACloudInfoModel) Flatten(ctx context.Context, from *dns.RecordACloudInfo, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordACloudInfoModel{}
	}
	m.AuthorityType = flex.FlattenStringPointer(from.AuthorityType)
	m.DelegatedScope = flex.FlattenStringPointer(from.DelegatedScope)
	m.MgmtPlatform = flex.FlattenStringPointer(from.MgmtPlatform)
	m.OwnedByAdaptor = types.BoolPointerValue(from.OwnedByAdaptor)
}
