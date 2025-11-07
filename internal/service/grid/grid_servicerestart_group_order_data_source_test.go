package grid_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccGridServicerestartGroupOrderDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_grid_servicerestart_group_order.test"
	resourceName := "nios_grid_servicerestart_group_order.test"
	var v grid.GridServicerestartGroupOrder

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGridServicerestartGroupOrderDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccGridServicerestartGroupOrderDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckGridServicerestartGroupOrderExists(context.Background(), resourceName, &v),
					}, testAccCheckGridServicerestartGroupOrderResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckGridServicerestartGroupOrderResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "groups", dataSourceName, "result.0.groups"),
	}
}

func testAccGridServicerestartGroupOrderDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_grid_servicerestart_group_order" "test" {
}

data "nios_grid_servicerestart_group_order" "test" {
  filters = {
	 = nios_grid_servicerestart_group_order.test.
  }
}
`)
}
