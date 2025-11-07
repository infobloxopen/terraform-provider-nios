package threatprotection_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/threatprotection"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccThreatprotectionGridRuleDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_threatprotection_grid_rule.test"
	resourceName := "nios_threatprotection_grid_rule.test"
	var v threatprotection.ThreatprotectionGridRule

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckThreatprotectionGridRuleDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccThreatprotectionGridRuleDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckThreatprotectionGridRuleExists(context.Background(), resourceName, &v),
					}, testAccCheckThreatprotectionGridRuleResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckThreatprotectionGridRuleResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "allowed_actions", dataSourceName, "result.0.allowed_actions"),
		resource.TestCheckResourceAttrPair(resourceName, "category", dataSourceName, "result.0.category"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "config", dataSourceName, "result.0.config"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "result.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "disabled", dataSourceName, "result.0.disabled"),
		resource.TestCheckResourceAttrPair(resourceName, "is_factory_reset_enabled", dataSourceName, "result.0.is_factory_reset_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "ruleset", dataSourceName, "result.0.ruleset"),
		resource.TestCheckResourceAttrPair(resourceName, "sid", dataSourceName, "result.0.sid"),
		resource.TestCheckResourceAttrPair(resourceName, "template", dataSourceName, "result.0.template"),
		resource.TestCheckResourceAttrPair(resourceName, "type", dataSourceName, "result.0.type"),
	}
}

func testAccThreatprotectionGridRuleDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_threatprotection_grid_rule" "test" {
}

data "nios_threatprotection_grid_rule" "test" {
  filters = {
	 = nios_threatprotection_grid_rule.test.
  }
}
`)
}
