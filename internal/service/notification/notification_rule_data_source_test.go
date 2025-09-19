package notification_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/notification"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

func TestAccNotificationRuleDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_notification_rule.test"
	resourceName := "nios_notification_rule.test"
	var v notification.NotificationRule
	name := acctest.RandomNameWithPrefix("example-notification-rule")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNotificationRuleDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNotificationRuleDataSourceConfigFilters(eventType, expressionList, name, notificationAction, notificationTarget, templateInstance),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckNotificationRuleExists(context.Background(), resourceName, &v),
					}, testAccCheckNotificationRuleResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions
func testAccCheckNotificationRuleResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "all_members", dataSourceName, "result.0.all_members"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_event_deduplication", dataSourceName, "result.0.enable_event_deduplication"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_event_deduplication_log", dataSourceName, "result.0.enable_event_deduplication_log"),
		resource.TestCheckResourceAttrPair(resourceName, "event_deduplication_fields", dataSourceName, "result.0.event_deduplication_fields"),
		resource.TestCheckResourceAttrPair(resourceName, "event_deduplication_lookback_period", dataSourceName, "result.0.event_deduplication_lookback_period"),
		resource.TestCheckResourceAttrPair(resourceName, "event_priority", dataSourceName, "result.0.event_priority"),
		resource.TestCheckResourceAttrPair(resourceName, "event_type", dataSourceName, "result.0.event_type"),
		resource.TestCheckResourceAttrPair(resourceName, "expression_list", dataSourceName, "result.0.expression_list"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "notification_action", dataSourceName, "result.0.notification_action"),
		resource.TestCheckResourceAttrPair(resourceName, "notification_target", dataSourceName, "result.0.notification_target"),
		resource.TestCheckResourceAttrPair(resourceName, "publish_settings", dataSourceName, "result.0.publish_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "scheduled_event", dataSourceName, "result.0.scheduled_event"),
		resource.TestCheckResourceAttrPair(resourceName, "selected_members", dataSourceName, "result.0.selected_members"),
		resource.TestCheckResourceAttrPair(resourceName, "template_instance", dataSourceName, "result.0.template_instance"),
		resource.TestCheckResourceAttrPair(resourceName, "use_publish_settings", dataSourceName, "result.0.use_publish_settings"),
	}
}

func testAccNotificationRuleDataSourceConfigFilters(eventType string, expressionList []map[string]any, name, notificationAction, notificationTarget string, templateInstance map[string]any) string {
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

data "nios_notification_rule" "test" {
	filters = {
		name = nios_notification_rule.test.name
	}
}
`, eventType, expressionListHCL, name, notificationAction, notificationTarget, templateInstanceHCL)
}
