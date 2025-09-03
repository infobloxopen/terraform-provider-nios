package smartfolder_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/smartfolder"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccSmartfolderPersonalDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_smartfolder_personal.test"
	resourceName := "nios_smartfolder_personal.test"
	var v smartfolder.SmartfolderPersonal

	name := acctest.RandomNameWithPrefix("example-smartfolder-personal-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSmartfolderPersonalDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSmartfolderPersonalDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSmartfolderPersonalExists(context.Background(), resourceName, &v),
					}, testAccCheckSmartfolderPersonalResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckSmartfolderPersonalResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "group_bys", dataSourceName, "result.0.group_bys"),
		resource.TestCheckResourceAttrPair(resourceName, "is_shortcut", dataSourceName, "result.0.is_shortcut"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "query_items", dataSourceName, "result.0.query_items"),
	}
}

func testAccSmartfolderPersonalDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_smartfolder_personal" "test" {
	name = %q
}

data "nios_smartfolder_personal" "test" {
  filters = {
	name = nios_smartfolder_personal.test.name
  }
}
`, name)
}
