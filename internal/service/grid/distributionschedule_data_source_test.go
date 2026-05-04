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

	now := time.Now()
	startTime := now.Add(12 * time.Hour).Format(utils.NaiveDatetimeLayout)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDistributionscheduleDataSourceConfig(true, startTime),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDistributionscheduleExists(context.Background(), resourceName, &v),
					}, testAccCheckDistributionscheduleResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
			// Deactivate schedule for Integration Testing
			{
				Config: testAccDistributionscheduleDeactivate(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nios_grid_distributionschedule.deactivate_schedule", "active", "false"),
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

func testAccDistributionscheduleDataSourceConfig(active bool, start_time string) string {

	return fmt.Sprintf(`
resource "nios_grid_distributionschedule" "test" {
	active     = %t
	start_time = %q
}

data "nios_grid_distributionschedule" "test" {
	depends_on = [nios_grid_distributionschedule.test]
}
`, active, start_time)
}
