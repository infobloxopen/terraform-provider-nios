package notification

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/notification"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

type NotificationRulePublishSettingsModel struct {
	EnabledAttributes types.List `tfsdk:"enabled_attributes"`
}

var NotificationRulePublishSettingsAttrTypes = map[string]attr.Type{
	"enabled_attributes": types.ListType{ElemType: types.StringType},
}

var NotificationRulePublishSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled_attributes": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.ValueStringsAre(
				stringvalidator.OneOf(
					"CLIENT_ID", "FINGERPRINT", "HOSTNAME", "INFOBLOX_MEMBER", "IPADDRESS",
					"LEASE_END_TIME", "LEASE_START_TIME", "LEASE_STATE", "MAC_OR_DUID", "NETBIOS_NAME",
				),
			),
		},
		Required:            true,
		MarkdownDescription: "The list of NIOS extensible attributes enalbed for publishsing to Cisco ISE endpoint.",
	},
}

func ExpandNotificationRulePublishSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *notification.NotificationRulePublishSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NotificationRulePublishSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NotificationRulePublishSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *notification.NotificationRulePublishSettings {
	if m == nil {
		return nil
	}
	to := &notification.NotificationRulePublishSettings{
		EnabledAttributes: flex.ExpandFrameworkListString(ctx, m.EnabledAttributes, diags),
	}
	return to
}

func FlattenNotificationRulePublishSettings(ctx context.Context, from *notification.NotificationRulePublishSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NotificationRulePublishSettingsAttrTypes)
	}
	m := NotificationRulePublishSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NotificationRulePublishSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NotificationRulePublishSettingsModel) Flatten(ctx context.Context, from *notification.NotificationRulePublishSettings, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NotificationRulePublishSettingsModel{}
	}
	m.EnabledAttributes = flex.FlattenFrameworkListString(ctx, from.EnabledAttributes, diags)
}
