package notification

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/infoblox-nios-go-client/notification"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
	planmodifiers "github.com/infobloxopen/terraform-provider-nios/internal/planmodifiers/immutable"
	internaltypes "github.com/infobloxopen/terraform-provider-nios/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-nios/internal/validator"
)

type NotificationRuleModel struct {
	Ref                              types.String                             `tfsdk:"ref"`
	AllMembers                       types.Bool                               `tfsdk:"all_members"`
	Comment                          types.String                             `tfsdk:"comment"`
	Disable                          types.Bool                               `tfsdk:"disable"`
	EnableEventDeduplication         types.Bool                               `tfsdk:"enable_event_deduplication"`
	EnableEventDeduplicationLog      types.Bool                               `tfsdk:"enable_event_deduplication_log"`
	EventDeduplicationFields         types.List                               `tfsdk:"event_deduplication_fields"`
	EventDeduplicationLookbackPeriod types.Int64                              `tfsdk:"event_deduplication_lookback_period"`
	EventPriority                    types.String                             `tfsdk:"event_priority"`
	EventType                        types.String                             `tfsdk:"event_type"`
	ExpressionList                   types.List                               `tfsdk:"expression_list"`
	Name                             types.String                             `tfsdk:"name"`
	NotificationAction               types.String                             `tfsdk:"notification_action"`
	NotificationTarget               internaltypes.CaseInsensitiveStringValue `tfsdk:"notification_target"`
	PublishSettings                  types.Object                             `tfsdk:"publish_settings"`
	ScheduledEvent                   types.Object                             `tfsdk:"scheduled_event"`
	SelectedMembers                  types.List                               `tfsdk:"selected_members"`
	TemplateInstance                 types.Object                             `tfsdk:"template_instance"`
	UsePublishSettings               types.Bool                               `tfsdk:"use_publish_settings"`
}

var NotificationRuleAttrTypes = map[string]attr.Type{
	"ref":                                 types.StringType,
	"all_members":                         types.BoolType,
	"comment":                             types.StringType,
	"disable":                             types.BoolType,
	"enable_event_deduplication":          types.BoolType,
	"enable_event_deduplication_log":      types.BoolType,
	"event_deduplication_fields":          types.ListType{ElemType: types.StringType},
	"event_deduplication_lookback_period": types.Int64Type,
	"event_priority":                      types.StringType,
	"event_type":                          types.StringType,
	"expression_list":                     types.ListType{ElemType: types.ObjectType{AttrTypes: NotificationRuleExpressionListAttrTypes}},
	"name":                                types.StringType,
	"notification_action":                 types.StringType,
	"notification_target":                 internaltypes.CaseInsensitiveString{},
	"publish_settings":                    types.ObjectType{AttrTypes: NotificationRulePublishSettingsAttrTypes},
	"scheduled_event":                     types.ObjectType{AttrTypes: NotificationRuleScheduledEventAttrTypes},
	"selected_members":                    types.ListType{ElemType: types.StringType},
	"template_instance":                   types.ObjectType{AttrTypes: NotificationRuleTemplateInstanceAttrTypes},
	"use_publish_settings":                types.BoolType,
}

var NotificationRuleResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"all_members": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Determines whether the notification rule is applied on all members or not. When this is set to False, the notification rule is applied only on selected_members.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "The notification rule descriptive comment.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a notification rule is disabled or not. When this is set to False, the notification rule is enabled.",
	},
	"enable_event_deduplication": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the notification rule for event deduplication is enabled. Note that to enable event deduplication, you must set at least one deduplication field.",
	},
	"enable_event_deduplication_log": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the notification rule for the event deduplication syslog is enabled.",
	},
	"event_deduplication_fields": schema.ListAttribute{
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.ValueStringsAre(
				stringvalidator.OneOf(
					"DISCOVERER", "DUID", "DXL_TOPIC", "IP_ADDRESS", "MAC_ADDRESS", "NETWORK", "NETWORK_VIEW", "OPERATION_TYPE", "QUERY_FQDN",
					"QUERY_NAME", "QUERY_TYPE", "RPZ_POLICY", "RPZ_TYPE", "RULE_ACTION", "RULE_CATEGORY", "RULE_SEVERITY", "RULE_SID",
					"SOURCE_IP", "SOURCE_PORT",
				),
			),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The list of fields that must be used in the notification rule for event deduplication.",
	},
	"event_deduplication_lookback_period": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(600),
		MarkdownDescription: "The lookback period for the notification rule for event deduplication.",
	},
	"event_priority": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("NORMAL"),
		MarkdownDescription: "Event priority.",
	},
	"event_type": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf(
				"ANALYTICS_DNS_TUNNEL", "DB_CHANGE_DHCP_FIXED_ADDRESS_IPV4", "DB_CHANGE_DHCP_FIXED_ADDRESS_IPV6", "DB_CHANGE_DHCP_NETWORK_IPV4",
				"DB_CHANGE_DHCP_NETWORK_IPV6", "DB_CHANGE_DHCP_RANGE_IPV4", "DB_CHANGE_DHCP_RANGE_IPV6", "DB_CHANGE_DNS_DISCOVERY_DATA",
				"DB_CHANGE_DNS_HOST_ADDRESS_IPV4", "DB_CHANGE_DNS_HOST_ADDRESS_IPV6", "DB_CHANGE_DNS_RECORD", "DB_CHANGE_DNS_ZONE",
				"DHCP_LEASES", "DNS_RPZ", "DXL_EVENT_SUBSCRIBER", "IPAM", "SCHEDULE", "SECURITY_ADP",
			),
		},
		MarkdownDescription: "The notification rule event type.",
	},
	"expression_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NotificationRuleExpressionListResourceSchemaAttributes,
		},
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The notification rule expression list.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The notification rule name.",
		PlanModifiers: []planmodifier.String{
			planmodifiers.ImmutableString(),
		},
	},
	"notification_action": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("CISCOISE_PUBLISH", "CISCOISE_QUARANTINE", "RESTAPI_TEMPLATE_INSTANCE"),
		},
		MarkdownDescription: "The notification rule action is applied if expression list evaluates to True.",
	},
	"notification_target": schema.StringAttribute{
		CustomType:          internaltypes.CaseInsensitiveString{},
		Required:            true,
		MarkdownDescription: "The notification target.",
	},
	"publish_settings": schema.SingleNestedAttribute{
		Attributes: NotificationRulePublishSettingsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_publish_settings")),
		},
		MarkdownDescription: "The CISCO ISE publish settings.",
	},
	"scheduled_event": schema.SingleNestedAttribute{
		Attributes:          NotificationRuleScheduledEventResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Schedule setting that must be specified if event_type is SCHEDULE",
	},
	"selected_members": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The list of the members on which the notification rule is applied. This field is deprecated.",
	},
	"template_instance": schema.SingleNestedAttribute{
		Attributes:          NotificationRuleTemplateInstanceResourceSchemaAttributes,
		Required:            true,
		MarkdownDescription: "The notification REST template instance.",
	},
	"use_publish_settings": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use flag for: publish_settings",
	},
}

func (m *NotificationRuleModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *notification.NotificationRule {
	if m == nil {
		return nil
	}
	to := &notification.NotificationRule{
		Comment:                          flex.ExpandStringPointer(m.Comment),
		Disable:                          flex.ExpandBoolPointer(m.Disable),
		EnableEventDeduplication:         flex.ExpandBoolPointer(m.EnableEventDeduplication),
		EnableEventDeduplicationLog:      flex.ExpandBoolPointer(m.EnableEventDeduplicationLog),
		EventDeduplicationFields:         flex.ExpandFrameworkListString(ctx, m.EventDeduplicationFields, diags),
		EventDeduplicationLookbackPeriod: flex.ExpandInt64Pointer(m.EventDeduplicationLookbackPeriod),
		EventPriority:                    flex.ExpandStringPointer(m.EventPriority),
		EventType:                        flex.ExpandStringPointer(m.EventType),
		ExpressionList:                   flex.ExpandFrameworkListNestedBlock(ctx, m.ExpressionList, diags, ExpandNotificationRuleExpressionList),
		NotificationAction:               flex.ExpandStringPointer(m.NotificationAction),
		NotificationTarget:               flex.ExpandStringPointer(m.NotificationTarget.StringValue),
		PublishSettings:                  ExpandNotificationRulePublishSettings(ctx, m.PublishSettings, diags),
		ScheduledEvent:                   ExpandNotificationRuleScheduledEvent(ctx, m.ScheduledEvent, diags),
		TemplateInstance:                 ExpandNotificationRuleTemplateInstance(ctx, m.TemplateInstance, diags),
		UsePublishSettings:               flex.ExpandBoolPointer(m.UsePublishSettings),
	}
	if isCreate {
		to.Name = flex.ExpandStringPointer(m.Name)
	}
	return to
}

func FlattenNotificationRule(ctx context.Context, from *notification.NotificationRule, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NotificationRuleAttrTypes)
	}
	m := NotificationRuleModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NotificationRuleAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NotificationRuleModel) Flatten(ctx context.Context, from *notification.NotificationRule, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NotificationRuleModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.AllMembers = types.BoolPointerValue(from.AllMembers)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disable = types.BoolPointerValue(from.Disable)
	m.EnableEventDeduplication = types.BoolPointerValue(from.EnableEventDeduplication)
	m.EnableEventDeduplicationLog = types.BoolPointerValue(from.EnableEventDeduplicationLog)
	m.EventDeduplicationFields = flex.FlattenFrameworkListString(ctx, from.EventDeduplicationFields, diags)
	m.EventDeduplicationLookbackPeriod = flex.FlattenInt64Pointer(from.EventDeduplicationLookbackPeriod)
	m.EventPriority = flex.FlattenStringPointer(from.EventPriority)
	m.EventType = flex.FlattenStringPointer(from.EventType)
	m.ExpressionList = flex.FlattenFrameworkListNestedBlock(ctx, from.ExpressionList, NotificationRuleExpressionListAttrTypes, diags, FlattenNotificationRuleExpressionList)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NotificationAction = flex.FlattenStringPointer(from.NotificationAction)
	m.NotificationTarget.StringValue = flex.FlattenStringPointer(from.NotificationTarget)
	m.PublishSettings = FlattenNotificationRulePublishSettings(ctx, from.PublishSettings, diags)
	m.ScheduledEvent = FlattenNotificationRuleScheduledEvent(ctx, from.ScheduledEvent, diags)
	m.SelectedMembers = flex.FlattenFrameworkListString(ctx, from.SelectedMembers, diags)
	m.TemplateInstance = FlattenNotificationRuleTemplateInstance(ctx, from.TemplateInstance, diags)
	m.UsePublishSettings = types.BoolPointerValue(from.UsePublishSettings)
}
