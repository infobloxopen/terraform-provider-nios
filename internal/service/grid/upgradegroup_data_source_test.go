package grid_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccUpgradegroupDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_grid_upgradegroup.test"
	resourceName := "nios_grid_upgradegroup.test"
	var v grid.Upgradegroup

	name := acctest.RandomNameWithPrefix("example-upgradegroup-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckUpgradegroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccUpgradegroupDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckUpgradegroupExists(context.Background(), resourceName, &v),
					}, testAccCheckUpgradegroupResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckUpgradegroupResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "distribution_dependent_group", dataSourceName, "result.0.distribution_dependent_group"),
		resource.TestCheckResourceAttrPair(resourceName, "distribution_policy", dataSourceName, "result.0.distribution_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "distribution_time", dataSourceName, "result.0.distribution_time"),
		resource.TestCheckResourceAttrPair(resourceName, "members", dataSourceName, "result.0.members"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "time_zone", dataSourceName, "result.0.time_zone"),
		resource.TestCheckResourceAttrPair(resourceName, "upgrade_dependent_group", dataSourceName, "result.0.upgrade_dependent_group"),
		resource.TestCheckResourceAttrPair(resourceName, "upgrade_policy", dataSourceName, "result.0.upgrade_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "upgrade_time", dataSourceName, "result.0.upgrade_time"),
	}
}

func testAccUpgradegroupDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_grid_upgradegroup" "test" {
  name = "%s"
}

data "nios_grid_upgradegroup" "test" {
  filters = {
	name = nios_grid_upgradegroup.test.name
  }
}
`, name)
}
