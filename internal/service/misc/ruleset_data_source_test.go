package misc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccRulesetDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_misc_ruleset.test"
	resourceName := "nios_misc_ruleset.test"
	var v misc.Ruleset

	name := acctest.RandomNameWithPrefix("example_ruleset")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRulesetDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRulesetDataSourceConfigFilters(name, "NXDOMAIN"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRulesetExists(context.Background(), resourceName, &v),
					}, testAccCheckRulesetResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// func TestAccRulesetDataSource_ExtAttrFilters(t *testing.T) {
// 	dataSourceName := "data.nios_misc_ruleset.test"
// 	resourceName := "nios_misc_ruleset.test"
// 	var v misc.Ruleset
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		CheckDestroy:             testAccCheckRulesetDestroy(context.Background(), &v),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccRulesetDataSourceConfigExtAttrFilters("value1"),
// 				Check: resource.ComposeTestCheckFunc(
// 					append([]resource.TestCheckFunc{
// 						testAccCheckRulesetExists(context.Background(), resourceName, &v),
// 					}, testAccCheckRulesetResourceAttrPair(resourceName, dataSourceName)...)...,
// 				),
// 			},
// 		},
// 	})
// }

// below all TestAcc functions

func testAccCheckRulesetResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disabled", dataSourceName, "result.0.disabled"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "nxdomain_rules", dataSourceName, "result.0.nxdomain_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "type", dataSourceName, "result.0.type"),
	}
}

func testAccRulesetDataSourceConfigFilters(name, ruleset_type string) string {
	return fmt.Sprintf(`
resource "nios_misc_ruleset" "test" {
  name = %q
  type = %q
}

data "nios_misc_ruleset" "test" {
  filters = {
	name = nios_misc_ruleset.test.name
  }
}
`, name, ruleset_type)
}

// func testAccRulesetDataSourceConfigExtAttrFilters(extAttrsValue string) string {
// 	return fmt.Sprintf(`
// resource "nios_misc_ruleset" "test" {
//   extattrs = {
//     Site = %q
//   }
// }

// data "nios_misc_ruleset" "test" {
//   extattrfilters = {
// 	Site = nios_misc_ruleset.test.extattrs.Site
//   }
// }
// `, extAttrsValue)
// }
