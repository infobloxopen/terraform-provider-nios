package grid_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

func TestAccUpgradescheduleDataSource_Read(t *testing.T) {
	dataSourceName := "data.nios_grid_upgradeschedule.test"
	resourceName := "nios_grid_upgradeschedule.test"
	var v grid.Upgradeschedule

	active := true

	now := time.Now()
	start_time := now.Add(12 * time.Hour).Format(utils.NaiveDatetimeLayout)
	upgrade_time := now.Add(24 * time.Hour).Format(utils.NaiveDatetimeLayout)

	upgrade_groups := []map[string]any{
		{
			"upgrade_time": upgrade_time,
			"name":         "Default",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckUpgradescheduleDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccUpgradescheduleDataSourceConfig(active, start_time, upgrade_groups),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckUpgradescheduleExists(context.Background(), resourceName, &v),
					}, testAccCheckUpgradescheduleResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckUpgradescheduleResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "active", dataSourceName, "result.0.active"),
		resource.TestCheckResourceAttrPair(resourceName, "start_time", dataSourceName, "result.0.start_time"),
		resource.TestCheckResourceAttrPair(resourceName, "time_zone", dataSourceName, "result.0.time_zone"),
		resource.TestCheckResourceAttrPair(resourceName, "upgrade_groups", dataSourceName, "result.0.upgrade_groups"),
	}
}

func testAccUpgradescheduleDataSourceConfig(active bool, start_time string, upgradeGroups []map[string]any) string {
	upgradeGroupsHCL := utils.ConvertSliceOfMapsToHCL(upgradeGroups)

	return fmt.Sprintf(`
resource "nios_grid_upgradeschedule" "test" {
	active     = %t
	start_time = %q
	upgrade_groups = %s
}

data "nios_grid_upgradeschedule" "test" {
  filters = {
	depends_on = [nios_grid_upgradeschedule.test]
  }
}
`, active, start_time, upgradeGroupsHCL)
}
