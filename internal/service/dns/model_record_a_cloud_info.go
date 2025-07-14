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

type RecordACloudInfoModel struct {
	DelegatedMember types.Object `tfsdk:"delegated_member"`
	DelegatedScope  types.String `tfsdk:"delegated_scope"`
	DelegatedRoot   types.String `tfsdk:"delegated_root"`
	OwnedByAdaptor  types.Bool   `tfsdk:"owned_by_adaptor"`
	Usage           types.String `tfsdk:"usage"`
	Tenant          types.String `tfsdk:"tenant"`
	MgmtPlatform    types.String `tfsdk:"mgmt_platform"`
	AuthorityType   types.String `tfsdk:"authority_type"`
}

var RecordACloudInfoAttrTypes = map[string]attr.Type{
	"delegated_member": types.ObjectType{AttrTypes: RecordacloudinfoDelegatedMemberAttrTypes},
	"delegated_scope":  types.StringType,
	"delegated_root":   types.StringType,
	"owned_by_adaptor": types.BoolType,
	"usage":            types.StringType,
	"tenant":           types.StringType,
	"mgmt_platform":    types.StringType,
	"authority_type":   types.StringType,
}

var RecordACloudInfoResourceSchemaAttributes = map[string]schema.Attribute{
	"delegated_member": schema.SingleNestedAttribute{
		Attributes: RecordacloudinfoDelegatedMemberResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"delegated_scope": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates the scope of delegation for the object. This can be one of the following: NONE (outside any delegation), ROOT (the delegation point), SUBTREE (within the scope of a delegation), RECLAIMING (within the scope of a delegation being reclaimed, either as the delegation point or in the subtree).",
	},
	"delegated_root": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates the root of the delegation if delegated_scope is SUBTREE or RECLAIMING. This is not set otherwise.",
	},
	"owned_by_adaptor": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines whether the object was created by the cloud adapter or not.",
	},
	"usage": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates the cloud origin of the object.",
	},
	"tenant": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Reference to the tenant object associated with the object, if any.",
	},
	"mgmt_platform": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates the specified cloud management platform.",
	},
	"authority_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Type of authority over the object.",
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
		DelegatedMember: ExpandRecordacloudinfoDelegatedMember(ctx, m.DelegatedMember, diags),
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
	m.DelegatedMember = FlattenRecordacloudinfoDelegatedMember(ctx, from.DelegatedMember, diags)
	m.DelegatedScope = flex.FlattenStringPointer(from.DelegatedScope)
	m.DelegatedRoot = flex.FlattenStringPointer(from.DelegatedRoot)
	m.OwnedByAdaptor = types.BoolPointerValue(from.OwnedByAdaptor)
	m.Usage = flex.FlattenStringPointer(from.Usage)
	m.Tenant = flex.FlattenStringPointer(from.Tenant)
	m.MgmtPlatform = flex.FlattenStringPointer(from.MgmtPlatform)
	m.AuthorityType = flex.FlattenStringPointer(from.AuthorityType)
}
