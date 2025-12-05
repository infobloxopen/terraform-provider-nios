package grid_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccGmcgroupDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_grid_gmcgroup.test"
	resourceName := "nios_grid_gmcgroup.test"
	var v grid.Gmcgroup

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGmcgroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccGmcgroupDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					}, testAccCheckGmcgroupResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckGmcgroupResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "gmc_promotion_policy", dataSourceName, "result.0.gmc_promotion_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "members", dataSourceName, "result.0.members"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "scheduled_time", dataSourceName, "result.0.scheduled_time"),
		resource.TestCheckResourceAttrPair(resourceName, "time_zone", dataSourceName, "result.0.time_zone"),
	}
}

func testAccGmcgroupDataSourceConfigFilters() string {
	return `
resource "nios_grid_gmcgroup" "test" {
}

data "nios_grid_gmcgroup" "test" {
  filters = {
	 = nios_grid_gmcgroup.test.
  }
}
`
}
