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

func TestAccDistributionscheduleDataSource_Read(t *testing.T) {
	dataSourceName := "data.nios_grid_distributionschedule.test"
	resourceName := "nios_grid_distributionschedule.test"
	var v grid.Distributionschedule
	active := true

	now := time.Now()
	start_time := now.Add(12 * time.Hour).Format(utils.NaiveDatetimeLayout)
	distribution_time := now.Add(24 * time.Hour).Format(utils.NaiveDatetimeLayout)

	upgrade_groups := []map[string]any{
		{
			"distribution_time": distribution_time,
			"name":              "Default",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDistributionscheduleDataSourceConfig(active, start_time, upgrade_groups),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDistributionscheduleExists(context.Background(), resourceName, &v),
					}, testAccCheckDistributionscheduleResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDistributionscheduleResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "active", dataSourceName, "result.0.active"),
		resource.TestCheckResourceAttrPair(resourceName, "start_time", dataSourceName, "result.0.start_time"),
		resource.TestCheckResourceAttrPair(resourceName, "time_zone", dataSourceName, "result.0.time_zone"),
		resource.TestCheckResourceAttrPair(resourceName, "upgrade_groups", dataSourceName, "result.0.upgrade_groups"),
	}
}

func testAccDistributionscheduleDataSourceConfig(active bool, start_time string, upgradeGroups []map[string]any) string {
	upgradeGroupsHCL := utils.ConvertSliceOfMapsToHCL(upgradeGroups)

	return fmt.Sprintf(`
resource "nios_grid_distributionschedule" "test" {
	active     = %t
	start_time = %q
	upgrade_groups = %s
}

data "nios_grid_distributionschedule" "test" {
	depends_on = [nios_grid_distributionschedule.test]
}
`, active, start_time, upgradeGroupsHCL)
}
