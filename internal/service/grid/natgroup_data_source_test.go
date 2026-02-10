package grid_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccNatgroupDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_grid_natgroup.test"
	resourceName := "nios_grid_natgroup.test"
	var v grid.Natgroup
	name := acctest.RandomNameWithPrefix("natgroup")
	comment := "This is a test natgroup"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNatgroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNatgroupDataSourceConfigFilters(name, comment),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckNatgroupExists(context.Background(), resourceName, &v),
					}, testAccCheckNatgroupResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckNatgroupResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
	}
}

func testAccNatgroupDataSourceConfigFilters(name, comment string) string {
	return fmt.Sprintf(`
resource "nios_grid_natgroup" "test" {
	name = %q
	comment = %q
}

data "nios_grid_natgroup" "test" {
  filters = {
	name = nios_grid_natgroup.test.name
  }
}
`, name, comment)
}
