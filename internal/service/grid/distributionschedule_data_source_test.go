package grid_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccDistributionscheduleDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_grid_distributionschedule.test"
	resourceName := "nios_grid_distributionschedule.test"
	var v grid.Distributionschedule
	active := true
	start_time := time.Now().Add(24 * time.Hour).Unix()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDistributionscheduleDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDistributionscheduleResourceConfigFilters(active, start_time),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDistributionscheduleExists(context.Background(), resourceName, &v),
				),
			},
			{
				Config: testAccDistributionscheduleDataSourceConfigFilters(active, start_time),
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
		// resource.TestCheckResourceAttrPair(resourceName, "upgrade_groups", dataSourceName, "result.0.upgrade_groups"),
	}
}

func testAccDistributionscheduleResourceConfigFilters(active bool, start_time int64) string {
	return fmt.Sprintf(`
resource "nios_grid_distributionschedule" "test" {
	  active     = %t
	  start_time = %d
}
`, active, start_time)
}

func testAccDistributionscheduleDataSourceConfigFilters(active bool, start_time int64) string {
	return fmt.Sprintf(`
resource "nios_grid_distributionschedule" "test" {
	  active     = %t
	  start_time = %d
}

data "nios_grid_distributionschedule" "test" {}
`, active, start_time)
}
