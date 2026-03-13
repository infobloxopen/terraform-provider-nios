package smartfolder_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/smartfolder"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccSmartfolderGlobalDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_smartfolder_global.test"
	resourceName := "nios_smartfolder_global.test"
	var v smartfolder.SmartfolderGlobal

	name := acctest.RandomNameWithPrefix("example-smartfolder-global")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSmartfolderGlobalDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSmartfolderGlobalDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSmartfolderGlobalExists(context.Background(), resourceName, &v),
					}, testAccCheckSmartfolderGlobalResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckSmartfolderGlobalResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "group_bys", dataSourceName, "result.0.group_bys"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "query_items", dataSourceName, "result.0.query_items"),
	}
}

func testAccSmartfolderGlobalDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_smartfolder_global" "test" {
  name = %q
}

data "nios_smartfolder_global" "test" {
  filters = {
	name = nios_smartfolder_global.test.name
  }
}
`, name)
}
