package notification_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/notification"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// OBJECTS TO BE PRESENT IN GRID FOR TESTS
// Notification Rest Endpoint - rest_api, syslog, cisco
// Notification Template - DHCP_Lease, syslog_action_template, IPAM_PxgridEvent

// TODO
// TestAccNotificationRuleResource_EventPriority
// TestAccNotificationRuleResource_ScheduledEvent
// TestAccNotificationRuleResource_NotificationAction

var readableAttributesForNotificationRule = "all_members,comment,disable,enable_event_deduplication,enable_event_deduplication_log,event_deduplication_fields,event_deduplication_lookback_period,event_priority,event_type,expression_list,name,notification_action,notification_target,publish_settings,scheduled_event,selected_members,template_instance,use_publish_settings"

var (
	notificationTarget = "notification:rest:endpoint/b25lLmVuZHBvaW50JDI:rest_api"
	eventType          = "DHCP_LEASES"
	expressionList     = []map[string]any{
		{
			"op":       "AND",
			"op1_type": "LIST",
		},
		{
			"op":       "EQ",
			"op1":      "DHCP_LEASE_STATE",
			"op1_type": "FIELD",
			"op2":      "DHCP_LEASE_STATE_ACTIVE",
			"op2_type": "STRING",
		},
		{
			"op": "ENDLIST",
		},
	}
	templateInstance = map[string]any{
		"template": "DHCP_Lease",
	}
	notificationAction = "RESTAPI_TEMPLATE_INSTANCE"
)

func TestAccNotificationRuleResource_basic(t *testing.T) {
	var resourceName = "nios_notification_rule.test"
	var v notification.NotificationRule
	name := acctest.RandomNameWithPrefix("example-notification-rule")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRuleBasicConfig(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "event_type", eventType),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_action", notificationAction),
					resource.TestCheckResourceAttr(resourceName, "notification_target", notificationTarget),
					resource.TestCheckResourceAttr(resourceName, "template_instance.template", "DHCP_Lease"),
					resource.TestCheckResourceAttr(resourceName, "expression_list.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "expression_list.0.op", "AND"),
					resource.TestCheckResourceAttr(resourceName, "expression_list.0.op1_type", "LIST"),
					resource.TestCheckResourceAttr(resourceName, "expression_list.1.op", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "expression_list.1.op1", "DHCP_LEASE_STATE"),
					resource.TestCheckResourceAttr(resourceName, "expression_list.1.op1_type", "FIELD"),
					resource.TestCheckResourceAttr(resourceName, "expression_list.1.op2", "DHCP_LEASE_STATE_ACTIVE"),
					resource.TestCheckResourceAttr(resourceName, "expression_list.1.op2_type", "STRING"),
					resource.TestCheckResourceAttr(resourceName, "expression_list.2.op", "ENDLIST"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "all_members", "true"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_event_deduplication", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_event_deduplication_log", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_publish_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRuleResource_disappears(t *testing.T) {
	resourceName := "nios_notification_rule.test"
	var v notification.NotificationRule
	name := acctest.RandomNameWithPrefix("example-notification-rule")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNotificationRuleDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNotificationRuleBasicConfig(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					testAccCheckNotificationRuleDisappears(context.Background(), &v),
				),
			},
		},
	})
}

func TestAccNotificationRuleResource_Comment(t *testing.T) {
	var resourceName = "nios_notification_rule.test_comment"
	var v notification.NotificationRule
	name := acctest.RandomNameWithPrefix("example-notification-rule")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRuleComment(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRuleComment(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRuleResource_Disable(t *testing.T) {
	var resourceName = "nios_notification_rule.test_disable"
	var v notification.NotificationRule
	name := acctest.RandomNameWithPrefix("example-notification-rule")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRuleDisable(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRuleDisable(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// Deduplication events are supported only for DNS_RPZ, SECURITY_ADP, DB_CHANGE_DNS_DISCOVERY_DATA, DXL_EVENT_SUBSCRIBER event types
func TestAccNotificationRuleResource_EnableEventDeduplication(t *testing.T) {
	var resourceName = "nios_notification_rule.test_enable_event_deduplication"
	var v notification.NotificationRule
	name := acctest.RandomNameWithPrefix("example-notification-rule")
	eventType := "DNS_RPZ"
	expressionList := []map[string]any{
		{
			"op":       "AND",
			"op1_type": "LIST",
		},
		{
			"op":       "EQ",
			"op1":      "DNS_RPZ_TYPE",
			"op1_type": "FIELD",
			"op2":      "DNS_RPZ_TYPE_IP",
			"op2_type": "STRING",
		},
		{
			"op": "ENDLIST",
		},
	}
	notificationTarget := "syslog:endpoint/b25lLmVuZHBvaW50JDM:syslog"
	templateInstance := map[string]any{
		"template": "syslog_action_template",
	}
	eventDeduplicationFields := []string{
		"SOURCE_IP",
		"QUERY_NAME",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRuleEnableEventDeduplication(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance, "false", eventDeduplicationFields),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_event_deduplication", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRuleEnableEventDeduplication(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance, "true", eventDeduplicationFields),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_event_deduplication", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRuleResource_EnableEventDeduplicationLog(t *testing.T) {
	var resourceName = "nios_notification_rule.test_enable_event_deduplication_log"
	var v notification.NotificationRule
	name := acctest.RandomNameWithPrefix("example-notification-rule")
	eventType := "DNS_RPZ"
	expressionList := []map[string]any{
		{
			"op":       "AND",
			"op1_type": "LIST",
		},
		{
			"op":       "EQ",
			"op1":      "DNS_RPZ_TYPE",
			"op1_type": "FIELD",
			"op2":      "DNS_RPZ_TYPE_IP",
			"op2_type": "STRING",
		},
		{
			"op": "ENDLIST",
		},
	}
	notificationTarget := "syslog:endpoint/b25lLmVuZHBvaW50JDM:syslog"
	templateInstance := map[string]any{
		"template": "syslog_action_template",
	}
	eventDeduplicationFields := []string{
		"SOURCE_IP",
		"QUERY_NAME",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRuleEnableEventDeduplicationLog(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance, "false", eventDeduplicationFields),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_event_deduplication_log", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRuleEnableEventDeduplicationLog(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance, "true", eventDeduplicationFields),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_event_deduplication_log", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRuleResource_EventDeduplicationFields(t *testing.T) {
	var resourceName = "nios_notification_rule.test_event_deduplication_fields"
	var v notification.NotificationRule
	name := acctest.RandomNameWithPrefix("example-notification-rule")
	eventType := "DNS_RPZ"
	expressionList := []map[string]any{
		{
			"op":       "AND",
			"op1_type": "LIST",
		},
		{
			"op":       "EQ",
			"op1":      "DNS_RPZ_TYPE",
			"op1_type": "FIELD",
			"op2":      "DNS_RPZ_TYPE_IP",
			"op2_type": "STRING",
		},
		{
			"op": "ENDLIST",
		},
	}
	notificationTarget := "syslog:endpoint/b25lLmVuZHBvaW50JDM:syslog"
	templateInstance := map[string]any{
		"template": "syslog_action_template",
	}
	eventDeduplicationFields := []string{
		"SOURCE_IP",
	}
	eventDeduplicationFieldsUpdate := []string{
		"SOURCE_IP",
		"QUERY_NAME",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRuleEventDeduplicationFields(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance, eventDeduplicationFields),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "event_deduplication_fields.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "event_deduplication_fields.0", "SOURCE_IP"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRuleEventDeduplicationFields(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance, eventDeduplicationFieldsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "event_deduplication_fields.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "event_deduplication_fields.0", "SOURCE_IP"),
					resource.TestCheckResourceAttr(resourceName, "event_deduplication_fields.1", "QUERY_NAME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRuleResource_EventDeduplicationLookbackPeriod(t *testing.T) {
	var resourceName = "nios_notification_rule.test_event_deduplication_lookback_period"
	var v notification.NotificationRule
	name := acctest.RandomNameWithPrefix("example-notification-rule")
	eventType := "DNS_RPZ"
	expressionList := []map[string]any{
		{
			"op":       "AND",
			"op1_type": "LIST",
		},
		{
			"op":       "EQ",
			"op1":      "DNS_RPZ_TYPE",
			"op1_type": "FIELD",
			"op2":      "DNS_RPZ_TYPE_IP",
			"op2_type": "STRING",
		},
		{
			"op": "ENDLIST",
		},
	}
	notificationTarget := "syslog:endpoint/b25lLmVuZHBvaW50JDM:syslog"
	templateInstance := map[string]any{
		"template": "syslog_action_template",
	}
	eventDeduplicationFields := []string{
		"SOURCE_IP",
		"QUERY_NAME",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRuleEventDeduplicationLookbackPeriod(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance, eventDeduplicationFields, "500"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "event_deduplication_lookback_period", "500"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRuleEventDeduplicationLookbackPeriod(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance, eventDeduplicationFields, "600"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "event_deduplication_lookback_period", "600"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// The event priority can be configured only for outbound notification rules that contain the scheduled event type
func TestAccNotificationRuleResource_EventPriority(t *testing.T) {
	t.Skip("Additional config is required for test")
	var resourceName = "nios_notification_rule.test_event_priority"
	var v notification.NotificationRule
	name := acctest.RandomNameWithPrefix("example-notification-rule")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRuleEventPriority(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance, "NORMAL"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "event_priority", "NORMAL"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRuleEventPriority(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance, "HIGH"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "event_priority", "HIGH"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRuleResource_EventType(t *testing.T) {
	var resourceName = "nios_notification_rule.test_event_type"
	var v notification.NotificationRule
	name := acctest.RandomNameWithPrefix("example-notification-rule")
	updatedEventType := "DNS_RPZ"
	updatedExpressionList := []map[string]any{
		{
			"op":       "AND",
			"op1_type": "LIST",
		},
		{
			"op":       "EQ",
			"op1":      "DNS_RPZ_TYPE",
			"op1_type": "FIELD",
			"op2":      "DNS_RPZ_TYPE_IP",
			"op2_type": "STRING",
		},
		{
			"op": "ENDLIST",
		},
	}
	updatedNotificationTarget := "syslog:endpoint/b25lLmVuZHBvaW50JDM:syslog"
	updatedTemplateInstance := map[string]any{
		"template": "syslog_action_template",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRuleEventType(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "event_type", "DHCP_LEASES"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRuleEventType(updatedEventType, updatedExpressionList, name, notificationAction, updatedNotificationTarget, updatedTemplateInstance),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "event_type", "DNS_RPZ"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRuleResource_ExpressionList(t *testing.T) {
	var resourceName = "nios_notification_rule.test_expression_list"
	var v notification.NotificationRule
	name := acctest.RandomNameWithPrefix("example-notification-rule")
	updatedEventType := "DNS_RPZ"
	updatedExpressionList := []map[string]any{
		{
			"op":       "AND",
			"op1_type": "LIST",
		},
		{
			"op":       "EQ",
			"op1":      "DNS_RPZ_TYPE",
			"op1_type": "FIELD",
			"op2":      "DNS_RPZ_TYPE_IP",
			"op2_type": "STRING",
		},
		{
			"op": "ENDLIST",
		},
	}
	updatedNotificationTarget := "syslog:endpoint/b25lLmVuZHBvaW50JDM:syslog"
	updatedTemplateInstance := map[string]any{
		"template": "syslog_action_template",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRuleExpressionList(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					// Check expression list length
					resource.TestCheckResourceAttr(resourceName, "expression_list.#", "3"),
					// Check first expression (AND)
					resource.TestCheckResourceAttr(resourceName, "expression_list.0.op", "AND"),
					resource.TestCheckResourceAttr(resourceName, "expression_list.0.op1_type", "LIST"),
					// Check second expression (EQ)
					resource.TestCheckResourceAttr(resourceName, "expression_list.1.op", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "expression_list.1.op1", "DHCP_LEASE_STATE"),
					resource.TestCheckResourceAttr(resourceName, "expression_list.1.op1_type", "FIELD"),
					resource.TestCheckResourceAttr(resourceName, "expression_list.1.op2", "DHCP_LEASE_STATE_ACTIVE"),
					resource.TestCheckResourceAttr(resourceName, "expression_list.1.op2_type", "STRING"),
					// Check third expression (ENDLIST)
					resource.TestCheckResourceAttr(resourceName, "expression_list.2.op", "ENDLIST"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRuleExpressionList(updatedEventType, updatedExpressionList, name, notificationAction, updatedNotificationTarget, updatedTemplateInstance),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					// Check updated expression list length
					resource.TestCheckResourceAttr(resourceName, "expression_list.#", "3"),
					// Check first updated expression (AND)
					resource.TestCheckResourceAttr(resourceName, "expression_list.0.op", "AND"),
					resource.TestCheckResourceAttr(resourceName, "expression_list.0.op1_type", "LIST"),
					// Check second updated expression (EQ)
					resource.TestCheckResourceAttr(resourceName, "expression_list.1.op", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "expression_list.1.op1", "DNS_RPZ_TYPE"),
					resource.TestCheckResourceAttr(resourceName, "expression_list.1.op1_type", "FIELD"),
					resource.TestCheckResourceAttr(resourceName, "expression_list.1.op2", "DNS_RPZ_TYPE_IP"),
					resource.TestCheckResourceAttr(resourceName, "expression_list.1.op2_type", "STRING"),
					// Check third updated expression (ENDLIST)
					resource.TestCheckResourceAttr(resourceName, "expression_list.2.op", "ENDLIST"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRuleResource_Name(t *testing.T) {
	var resourceName = "nios_notification_rule.test_name"
	var v notification.NotificationRule
	name := acctest.RandomNameWithPrefix("example-notification-rule")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRuleName(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRuleResource_NotificationAction(t *testing.T) {
	t.Skip("Additional config is required for test")
	var resourceName = "nios_notification_rule.test_notification_action"
	var v notification.NotificationRule
	name := acctest.RandomNameWithPrefix("example-notification-rule")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRuleNotificationAction(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "notification_action", "RESTAPI_TEMPLATE_INSTANCE"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRuleNotificationAction(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "notification_action", "NOTIFICATION_ACTION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRuleResource_NotificationTarget(t *testing.T) {
	var resourceName = "nios_notification_rule.test_notification_target"
	var v notification.NotificationRule
	name := acctest.RandomNameWithPrefix("example-notification-rule")
	updatedEventType := "DNS_RPZ"
	updatedExpressionList := []map[string]any{
		{
			"op":       "AND",
			"op1_type": "LIST",
		},
		{
			"op":       "EQ",
			"op1":      "DNS_RPZ_TYPE",
			"op1_type": "FIELD",
			"op2":      "DNS_RPZ_TYPE_IP",
			"op2_type": "STRING",
		},
		{
			"op": "ENDLIST",
		},
	}
	updatedNotificationTarget := "syslog:endpoint/b25lLmVuZHBvaW50JDM:syslog"
	updatedTemplateInstance := map[string]any{
		"template": "syslog_action_template",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRuleNotificationTarget(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "notification_target", notificationTarget),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRuleNotificationTarget(updatedEventType, updatedExpressionList, name, notificationAction, updatedNotificationTarget, updatedTemplateInstance),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "notification_target", updatedNotificationTarget),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRuleResource_PublishSettings(t *testing.T) {
	var resourceName = "nios_notification_rule.test_publish_settings"
	var v notification.NotificationRule
	name := acctest.RandomNameWithPrefix("example-notification-rule")
	notificationTarget := "pxgrid:endpoint/b25lLmVuZHBvaW50JDU:cisco"
	templateInstance := map[string]any{
		"template": "IPAM_PxgridEvent",
	}
	publishSettings := map[string]any{
		"enabled_attributes": []string{"CLIENT_ID", "IPADDRESS"},
	}
	updatedPublishSettings := map[string]any{
		"enabled_attributes": []string{"CLIENT_ID", "IPADDRESS", "LEASE_STATE"},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRulePublishSettings(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance, publishSettings, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "publish_settings.enabled_attributes.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "publish_settings.enabled_attributes.0", "CLIENT_ID"),
					resource.TestCheckResourceAttr(resourceName, "publish_settings.enabled_attributes.1", "IPADDRESS"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRulePublishSettings(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance, updatedPublishSettings, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "publish_settings.enabled_attributes.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "publish_settings.enabled_attributes.0", "CLIENT_ID"),
					resource.TestCheckResourceAttr(resourceName, "publish_settings.enabled_attributes.1", "IPADDRESS"),
					resource.TestCheckResourceAttr(resourceName, "publish_settings.enabled_attributes.2", "LEASE_STATE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRuleResource_ScheduledEvent(t *testing.T) {
	t.Skip("Additional config is required for test")
	var resourceName = "nios_notification_rule.test_scheduled_event"
	var v notification.NotificationRule
	name := acctest.RandomNameWithPrefix("example-notification-rule")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRuleScheduledEvent(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance, "SCHEDULED_EVENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "scheduled_event", "SCHEDULED_EVENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRuleScheduledEvent(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance, "SCHEDULED_EVENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "scheduled_event", "SCHEDULED_EVENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRuleResource_TemplateInstance(t *testing.T) {
	var resourceName = "nios_notification_rule.test_template_instance"
	var v notification.NotificationRule
	name := acctest.RandomNameWithPrefix("example-notification-rule")
	updatedEventType := "DNS_RPZ"
	updatedExpressionList := []map[string]any{
		{
			"op":       "AND",
			"op1_type": "LIST",
		},
		{
			"op":       "EQ",
			"op1":      "DNS_RPZ_TYPE",
			"op1_type": "FIELD",
			"op2":      "DNS_RPZ_TYPE_IP",
			"op2_type": "STRING",
		},
		{
			"op": "ENDLIST",
		},
	}
	updatedNotificationTarget := "syslog:endpoint/b25lLmVuZHBvaW50JDM:syslog"
	updatedTemplateInstance := map[string]any{
		"template": "syslog_action_template",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRuleTemplateInstance(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_instance.template", "DHCP_Lease"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRuleTemplateInstance(updatedEventType, updatedExpressionList, name, notificationAction, updatedNotificationTarget, updatedTemplateInstance),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_instance.template", "syslog_action_template"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRuleResource_UsePublishSettings(t *testing.T) {
	var resourceName = "nios_notification_rule.test_use_publish_settings"
	var v notification.NotificationRule
	name := acctest.RandomNameWithPrefix("example-notification-rule")
	updatedNotificationTarget := "pxgrid:endpoint/b25lLmVuZHBvaW50JDU:cisco"
	updatedTemplateInstance := map[string]any{
		"template": "IPAM_PxgridEvent",
	}
	publishSettings := map[string]any{
		"enabled_attributes": []string{"CLIENT_ID", "IPADDRESS"},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRuleUsePublishSettings(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_publish_settings", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRuleUsePublishSettingsUpdate(eventType, expressionList, name, notificationAction, updatedNotificationTarget, updatedTemplateInstance, publishSettings, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_publish_settings", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckNotificationRuleExists(ctx context.Context, resourceName string, v *notification.NotificationRule) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.NotificationAPI.
			NotificationRuleAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForNotificationRule).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetNotificationRuleResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetNotificationRuleResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckNotificationRuleDestroy(ctx context.Context, v *notification.NotificationRule) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.NotificationAPI.
			NotificationRuleAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForNotificationRule).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				// resource was deleted
				return nil
			}
			return err
		}
		return errors.New("expected to be deleted")
	}
}

func testAccCheckNotificationRuleDisappears(ctx context.Context, v *notification.NotificationRule) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.NotificationAPI.
			NotificationRuleAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccNotificationRuleBasicConfig(eventType string, expressionList []map[string]any, name, notificationAction, notificationTarget string, templateInstance map[string]any) string {
	expressionListHCL := utils.ConvertSliceOfMapsToHCL(expressionList)
	templateInstanceHCL := utils.ConvertMapToHCL(templateInstance)
	return fmt.Sprintf(`
resource "nios_notification_rule" "test" {
    event_type = %q
    expression_list = %s
    name = %q
    notification_action = %q
    notification_target = %q
    template_instance = %s
}
`, eventType, expressionListHCL, name, notificationAction, notificationTarget, templateInstanceHCL)
}

func testAccNotificationRuleComment(eventType string, expressionList []map[string]any, name, notificationAction, notificationTarget string, templateInstance map[string]any, comment string) string {
	expressionListHCL := utils.ConvertSliceOfMapsToHCL(expressionList)
	templateInstanceHCL := utils.ConvertMapToHCL(templateInstance)
	return fmt.Sprintf(`
resource "nios_notification_rule" "test_comment" {
    event_type = %q
    expression_list = %s
    name = %q
    notification_action = %q
    notification_target = %q
    template_instance = %s
    comment = %q
}
`, eventType, expressionListHCL, name, notificationAction, notificationTarget, templateInstanceHCL, comment)
}

func testAccNotificationRuleDisable(eventType string, expressionList []map[string]any, name, notificationAction, notificationTarget string, templateInstance map[string]any, disable string) string {
	expressionListHCL := utils.ConvertSliceOfMapsToHCL(expressionList)
	templateInstanceHCL := utils.ConvertMapToHCL(templateInstance)
	return fmt.Sprintf(`
resource "nios_notification_rule" "test_disable" {
    event_type = %q
    expression_list = %s
    name = %q
    notification_action = %q
    notification_target = %q
    template_instance = %s
    disable = %q
}
`, eventType, expressionListHCL, name, notificationAction, notificationTarget, templateInstanceHCL, disable)
}

func testAccNotificationRuleEnableEventDeduplication(eventType string, expressionList []map[string]any, name, notificationAction, notificationTarget string, templateInstance map[string]any, enableEventDeduplication string, eventDeduplicationFields []string) string {
	expressionListHCL := utils.ConvertSliceOfMapsToHCL(expressionList)
	templateInstanceHCL := utils.ConvertMapToHCL(templateInstance)
	eventDeduplicationFieldsHCL := utils.ConvertStringSliceToHCL(eventDeduplicationFields)
	return fmt.Sprintf(`
resource "nios_notification_rule" "test_enable_event_deduplication" {
    event_type = %q
    expression_list = %s
    name = %q
    notification_action = %q
    notification_target = %q
    template_instance = %s
    enable_event_deduplication = %q
	event_deduplication_fields = %s
}
`, eventType, expressionListHCL, name, notificationAction, notificationTarget, templateInstanceHCL, enableEventDeduplication, eventDeduplicationFieldsHCL)
}

func testAccNotificationRuleEnableEventDeduplicationLog(eventType string, expressionList []map[string]any, name, notificationAction, notificationTarget string, templateInstance map[string]any, enableEventDeduplicationLog string, eventDeduplicationFields []string) string {
	expressionListHCL := utils.ConvertSliceOfMapsToHCL(expressionList)
	templateInstanceHCL := utils.ConvertMapToHCL(templateInstance)
	eventDeduplicationFieldsHCL := utils.ConvertStringSliceToHCL(eventDeduplicationFields)
	return fmt.Sprintf(`
resource "nios_notification_rule" "test_enable_event_deduplication_log" {
    event_type = %q
    expression_list = %s
    name = %q
    notification_action = %q
    notification_target = %q
    template_instance = %s
    enable_event_deduplication_log = %q
	event_deduplication_fields = %s
}
`, eventType, expressionListHCL, name, notificationAction, notificationTarget, templateInstanceHCL, enableEventDeduplicationLog, eventDeduplicationFieldsHCL)
}

func testAccNotificationRuleEventDeduplicationFields(eventType string, expressionList []map[string]any, name, notificationAction, notificationTarget string, templateInstance map[string]any, eventDeduplicationFields []string) string {
	expressionListHCL := utils.ConvertSliceOfMapsToHCL(expressionList)
	templateInstanceHCL := utils.ConvertMapToHCL(templateInstance)
	eventDeduplicationFieldsHCL := utils.ConvertStringSliceToHCL(eventDeduplicationFields)
	return fmt.Sprintf(`
resource "nios_notification_rule" "test_event_deduplication_fields" {
    event_type = %q
    expression_list = %s
    name = %q
    notification_action = %q
    notification_target = %q
    template_instance = %s
    event_deduplication_fields = %s
}
`, eventType, expressionListHCL, name, notificationAction, notificationTarget, templateInstanceHCL, eventDeduplicationFieldsHCL)
}

func testAccNotificationRuleEventDeduplicationLookbackPeriod(eventType string, expressionList []map[string]any, name, notificationAction, notificationTarget string, templateInstance map[string]any, eventDeduplicationFields []string, eventDeduplicationLookbackPeriod string) string {
	expressionListHCL := utils.ConvertSliceOfMapsToHCL(expressionList)
	templateInstanceHCL := utils.ConvertMapToHCL(templateInstance)
	eventDeduplicationFieldsHCL := utils.ConvertStringSliceToHCL(eventDeduplicationFields)
	return fmt.Sprintf(`
resource "nios_notification_rule" "test_event_deduplication_lookback_period" {
    event_type = %q
    expression_list = %s
    name = %q
    notification_action = %q
    notification_target = %q
    template_instance = %s
    event_deduplication_lookback_period = %q
	event_deduplication_fields = %s
}
`, eventType, expressionListHCL, name, notificationAction, notificationTarget, templateInstanceHCL, eventDeduplicationLookbackPeriod, eventDeduplicationFieldsHCL)
}

func testAccNotificationRuleEventPriority(eventType string, expressionList []map[string]any, name, notificationAction, notificationTarget string, templateInstance map[string]any, eventPriority string) string {
	expressionListHCL := utils.ConvertSliceOfMapsToHCL(expressionList)
	templateInstanceHCL := utils.ConvertMapToHCL(templateInstance)
	return fmt.Sprintf(`
resource "nios_notification_rule" "test_event_priority" {
    event_type = %q
    expression_list = %s
    name = %q
    notification_action = %q
    notification_target = %q
    template_instance = %s
    event_priority = %q
}
`, eventType, expressionListHCL, name, notificationAction, notificationTarget, templateInstanceHCL, eventPriority)
}

func testAccNotificationRuleEventType(eventType string, expressionList []map[string]any, name, notificationAction, notificationTarget string, templateInstance map[string]any) string {
	expressionListHCL := utils.ConvertSliceOfMapsToHCL(expressionList)
	templateInstanceHCL := utils.ConvertMapToHCL(templateInstance)
	return fmt.Sprintf(`
resource "nios_notification_rule" "test_event_type" {
    event_type = %q
    expression_list = %s
    name = %q
    notification_action = %q
    notification_target = %q
    template_instance = %s
}
`, eventType, expressionListHCL, name, notificationAction, notificationTarget, templateInstanceHCL)
}

func testAccNotificationRuleExpressionList(eventType string, expressionList []map[string]any, name, notificationAction, notificationTarget string, templateInstance map[string]any) string {
	expressionListHCL := utils.ConvertSliceOfMapsToHCL(expressionList)
	templateInstanceHCL := utils.ConvertMapToHCL(templateInstance)
	return fmt.Sprintf(`
resource "nios_notification_rule" "test_expression_list" {
    event_type = %q
    expression_list = %s
    name = %q
    notification_action = %q
    notification_target = %q
    template_instance = %s
}
`, eventType, expressionListHCL, name, notificationAction, notificationTarget, templateInstanceHCL)
}

func testAccNotificationRuleName(eventType string, expressionList []map[string]any, name, notificationAction, notificationTarget string, templateInstance map[string]any) string {
	expressionListHCL := utils.ConvertSliceOfMapsToHCL(expressionList)
	templateInstanceHCL := utils.ConvertMapToHCL(templateInstance)
	return fmt.Sprintf(`
resource "nios_notification_rule" "test_name" {
    event_type = %q
    expression_list = %s
    name = %q
    notification_action = %q
    notification_target = %q
    template_instance = %s
}
`, eventType, expressionListHCL, name, notificationAction, notificationTarget, templateInstanceHCL)
}

func testAccNotificationRuleNotificationAction(eventType string, expressionList []map[string]any, name, notificationAction, notificationTarget string, templateInstance map[string]any) string {
	expressionListHCL := utils.ConvertSliceOfMapsToHCL(expressionList)
	templateInstanceHCL := utils.ConvertMapToHCL(templateInstance)
	return fmt.Sprintf(`
resource "nios_notification_rule" "test_notification_action" {
    event_type = %q
    expression_list = %s
    name = %q
    notification_action = %q
    notification_target = %q
    template_instance = %s
}
`, eventType, expressionListHCL, name, notificationAction, notificationTarget, templateInstanceHCL)
}

func testAccNotificationRuleNotificationTarget(eventType string, expressionList []map[string]any, name, notificationAction, notificationTarget string, templateInstance map[string]any) string {
	expressionListHCL := utils.ConvertSliceOfMapsToHCL(expressionList)
	templateInstanceHCL := utils.ConvertMapToHCL(templateInstance)
	return fmt.Sprintf(`
resource "nios_notification_rule" "test_notification_target" {
    event_type = %q
    expression_list = %s
    name = %q
    notification_action = %q
    notification_target = %q
    template_instance = %s
}
`, eventType, expressionListHCL, name, notificationAction, notificationTarget, templateInstanceHCL)
}

func testAccNotificationRulePublishSettings(eventType string, expressionList []map[string]any, name, notificationAction, notificationTarget string, templateInstance map[string]any, publishSettings map[string]any, usePublishSettings string) string {
	expressionListHCL := utils.ConvertSliceOfMapsToHCL(expressionList)
	templateInstanceHCL := utils.ConvertMapToHCL(templateInstance)
	publishSettingsHCL := utils.ConvertMapToHCL(publishSettings)
	return fmt.Sprintf(`
resource "nios_notification_rule" "test_publish_settings" {
    event_type = %q
    expression_list = %s
    name = %q
    notification_action = %q
    notification_target = %q
    template_instance = %s
    publish_settings = %s
	use_publish_settings = %q
}
`, eventType, expressionListHCL, name, notificationAction, notificationTarget, templateInstanceHCL, publishSettingsHCL, usePublishSettings)
}

func testAccNotificationRuleScheduledEvent(eventType string, expressionList []map[string]any, name, notificationAction, notificationTarget string, templateInstance map[string]any, scheduledEvent string) string {
	return fmt.Sprintf(`
resource "nios_notification_rule" "test_scheduled_event" {
    event_type = %q
    expression_list = %q
    name = %q
    notification_action = %q
    notification_target = %q
    template_instance = %q
    scheduled_event = %q
}
`, eventType, expressionList, name, notificationAction, notificationTarget, templateInstance, scheduledEvent)
}

func testAccNotificationRuleTemplateInstance(eventType string, expressionList []map[string]any, name, notificationAction, notificationTarget string, templateInstance map[string]any) string {
	expressionListHCL := utils.ConvertSliceOfMapsToHCL(expressionList)
	templateInstanceHCL := utils.ConvertMapToHCL(templateInstance)
	return fmt.Sprintf(`
resource "nios_notification_rule" "test_template_instance" {
    event_type = %q
    expression_list = %s
    name = %q
    notification_action = %q
    notification_target = %q
    template_instance = %s
}
`, eventType, expressionListHCL, name, notificationAction, notificationTarget, templateInstanceHCL)
}

func testAccNotificationRuleUsePublishSettings(eventType string, expressionList []map[string]any, name, notificationAction, notificationTarget string, templateInstance map[string]any, usePublishSettings string) string {
	expressionListHCL := utils.ConvertSliceOfMapsToHCL(expressionList)
	templateInstanceHCL := utils.ConvertMapToHCL(templateInstance)
	return fmt.Sprintf(`
resource "nios_notification_rule" "test_use_publish_settings" {
    event_type = %q
    expression_list = %s
    name = %q
    notification_action = %q
    notification_target = %q
    template_instance = %s
    use_publish_settings = %q
}
`, eventType, expressionListHCL, name, notificationAction, notificationTarget, templateInstanceHCL, usePublishSettings)
}

func testAccNotificationRuleUsePublishSettingsUpdate(eventType string, expressionList []map[string]any, name, notificationAction, notificationTarget string, templateInstance map[string]any, publishSettings map[string]any, usePublishSettings string) string {
	expressionListHCL := utils.ConvertSliceOfMapsToHCL(expressionList)
	templateInstanceHCL := utils.ConvertMapToHCL(templateInstance)
	publishSettingsHCL := utils.ConvertMapToHCL(publishSettings)
	return fmt.Sprintf(`
resource "nios_notification_rule" "test_use_publish_settings" {
    event_type = %q
    expression_list = %s
    name = %q
    notification_action = %q
    notification_target = %q
    template_instance = %s
	publish_settings = %s
    use_publish_settings = %q
}
`, eventType, expressionListHCL, name, notificationAction, notificationTarget, templateInstanceHCL, publishSettingsHCL, usePublishSettings)
}
