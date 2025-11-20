package misc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccTftpfiledirDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_misc_tftpfiledir.test"
	resourceName := "nios_misc_tftpfiledir.test"
	var v misc.Tftpfiledir

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTftpfiledirDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccTftpfiledirDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
					}, testAccCheckTftpfiledirResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckTftpfiledirResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "directory", dataSourceName, "result.0.directory"),
		resource.TestCheckResourceAttrPair(resourceName, "is_synced_to_gm", dataSourceName, "result.0.is_synced_to_gm"),
		resource.TestCheckResourceAttrPair(resourceName, "last_modify", dataSourceName, "result.0.last_modify"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "type", dataSourceName, "result.0.type"),
		resource.TestCheckResourceAttrPair(resourceName, "vtftp_dir_members", dataSourceName, "result.0.vtftp_dir_members"),
	}
}

func testAccTftpfiledirDataSourceConfigFilters() string {
	return `
resource "nios_misc_tftpfiledir" "test" {
}

data "nios_misc_tftpfiledir" "test" {
  filters = {
	 = nios_misc_tftpfiledir.test.
  }
}
`
}
