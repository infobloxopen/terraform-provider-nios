package threatinsight

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/threatinsight"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type ThreatinsightCloudclientModel struct {
	Ref              types.String `tfsdk:"ref"`
	Uuid             types.String `tfsdk:"uuid"`
	BlacklistRpzList types.List   `tfsdk:"blacklist_rpz_list"`
	Enable           types.Bool   `tfsdk:"enable"`
	ForceRefresh     types.Bool   `tfsdk:"forcerefresh"`
	Interval         types.Int64  `tfsdk:"interval"`
}

var ThreatinsightCloudclientAttrTypes = map[string]attr.Type{
	"ref":                types.StringType,
	"uuid":               types.StringType,
	"blacklist_rpz_list": types.ListType{ElemType: types.StringType},
	"enable":             types.BoolType,
	"forcerefresh":       types.BoolType,
	"interval":           types.Int64Type,
}

var ThreatinsightCloudclientResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"uuid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The uuid to the object.",
	},
	"blacklist_rpz_list": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		MarkdownDescription: "The RPZs to which you apply newly detected domains through the Infoblox Threat Insight Cloud Client.",
	},
	"enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the Threat Insight in Cloud Client is enabled.",
	},
	"forcerefresh": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Force a refresh if at least one RPZ is configured.",
	},
	"interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The time interval (in seconds) for requesting newly detected domains by the Infoblox Threat Insight Cloud Client and applying them to the list of configured RPZs.",
	},
}

func ExpandThreatinsightCloudclient(ctx context.Context, o types.Object, diags *diag.Diagnostics) *threatinsight.ThreatinsightCloudclient {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ThreatinsightCloudclientModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ThreatinsightCloudclientModel) Expand(ctx context.Context, diags *diag.Diagnostics) *threatinsight.ThreatinsightCloudclient {
	if m == nil {
		return nil
	}
	to := &threatinsight.ThreatinsightCloudclient{
		Ref:              flex.ExpandStringPointer(m.Ref),
		Uuid:             flex.ExpandStringPointer(m.Uuid),
		BlacklistRpzList: flex.ExpandFrameworkListString(ctx, m.BlacklistRpzList, diags),
		Enable:           flex.ExpandBoolPointer(m.Enable),
		ForceRefresh:     flex.ExpandBoolPointer(m.ForceRefresh),
		Interval:         flex.ExpandInt64Pointer(m.Interval),
	}
	return to
}

func FlattenThreatinsightCloudclient(ctx context.Context, from *threatinsight.ThreatinsightCloudclient, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ThreatinsightCloudclientAttrTypes)
	}
	m := ThreatinsightCloudclientModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ThreatinsightCloudclientAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ThreatinsightCloudclientModel) Flatten(ctx context.Context, from *threatinsight.ThreatinsightCloudclient, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ThreatinsightCloudclientModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Uuid = flex.FlattenStringPointer(from.Uuid)
	m.BlacklistRpzList = flex.FlattenFrameworkListString(ctx, from.BlacklistRpzList, diags)
	m.Enable = types.BoolPointerValue(from.Enable)
	m.ForceRefresh = types.BoolPointerValue(from.ForceRefresh)
	m.Interval = flex.FlattenInt64Pointer(from.Interval)
}
