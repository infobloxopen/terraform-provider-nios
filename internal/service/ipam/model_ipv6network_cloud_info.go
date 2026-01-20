package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	"github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/plancontrol"
)

type Ipv6networkCloudInfoModel struct {
	DelegatedMember types.Object `tfsdk:"delegated_member"`
	DelegatedScope  types.String `tfsdk:"delegated_scope"`
	DelegatedRoot   types.String `tfsdk:"delegated_root"`
	OwnedByAdaptor  types.Bool   `tfsdk:"owned_by_adaptor"`
	Usage           types.String `tfsdk:"usage"`
	Tenant          types.String `tfsdk:"tenant"`
	MgmtPlatform    types.String `tfsdk:"mgmt_platform"`
	AuthorityType   types.String `tfsdk:"authority_type"`
}

var Ipv6networkCloudInfoAttrTypes = map[string]attr.Type{
	"delegated_member": types.ObjectType{AttrTypes: Ipv6networkcloudinfoDelegatedMemberAttrTypes},
	"delegated_scope":  types.StringType,
	"delegated_root":   types.StringType,
	"owned_by_adaptor": types.BoolType,
	"usage":            types.StringType,
	"tenant":           types.StringType,
	"mgmt_platform":    types.StringType,
	"authority_type":   types.StringType,
}

var Ipv6networkCloudInfoResourceSchemaAttributes = map[string]schema.Attribute{
	"delegated_member": schema.SingleNestedAttribute{
		Attributes:          Ipv6networkcloudinfoDelegatedMemberResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "Contains information about the delegated member of the object. This is not set if the object is not delegated.",
		Optional:            true,
	},
	"delegated_scope": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates the scope of delegation for the object. This can be one of the following: NONE (outside any delegation), ROOT (the delegation point), SUBTREE (within the scope of a delegation), RECLAIMING (within the scope of a delegation being reclaimed, either as the delegation point or in the subtree).",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"delegated_root": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates the root of the delegation if delegated_scope is SUBTREE or RECLAIMING. This is not set otherwise.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"owned_by_adaptor": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines whether the object was created by the cloud adapter or not.",
		PlanModifiers: []planmodifier.Bool{
			plancontrol.UseStateForUnknownBool(),
		},
	},
	"usage": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates the cloud origin of the object.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"tenant": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Reference to the tenant object associated with the object, if any.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"mgmt_platform": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates the specified cloud management platform.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
	"authority_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Type of authority over the object.",
		PlanModifiers: []planmodifier.String{
			plancontrol.UseStateForUnknownString(),
		},
	},
}

func ExpandIpv6networkCloudInfo(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.Ipv6networkCloudInfo {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkCloudInfoModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Ipv6networkCloudInfoModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Ipv6networkCloudInfo {
	if m == nil {
		return nil
	}
	to := &ipam.Ipv6networkCloudInfo{
		DelegatedMember: ExpandIpv6networkcloudinfoDelegatedMember(ctx, m.DelegatedMember, diags),
	}
	return to
}

func FlattenIpv6networkCloudInfo(ctx context.Context, from *ipam.Ipv6networkCloudInfo, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkCloudInfoAttrTypes)
	}
	m := Ipv6networkCloudInfoModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkCloudInfoAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Ipv6networkCloudInfoModel) Flatten(ctx context.Context, from *ipam.Ipv6networkCloudInfo, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Ipv6networkCloudInfoModel{}
	}
	m.DelegatedMember = FlattenIpv6networkcloudinfoDelegatedMember(ctx, from.DelegatedMember, diags)
	m.DelegatedScope = flex.FlattenStringPointer(from.DelegatedScope)
	m.DelegatedRoot = flex.FlattenStringPointer(from.DelegatedRoot)
	m.OwnedByAdaptor = types.BoolPointerValue(from.OwnedByAdaptor)
	m.Usage = flex.FlattenStringPointer(from.Usage)
	m.Tenant = flex.FlattenStringPointer(from.Tenant)
	m.MgmtPlatform = flex.FlattenStringPointer(from.MgmtPlatform)
	m.AuthorityType = flex.FlattenStringPointer(from.AuthorityType)
}
