package misc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccDatacollectionclusterDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_misc_datacollectioncluster.test"
	resourceName := "nios_misc_datacollectioncluster.test"
	var v misc.Datacollectioncluster

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDatacollectionclusterDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDatacollectionclusterDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDatacollectionclusterExists(context.Background(), resourceName, &v),
					}, testAccCheckDatacollectionclusterResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDatacollectionclusterResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_registration", dataSourceName, "result.0.enable_registration"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
	}
}

func testAccDatacollectionclusterDataSourceConfigFilters() string {
	return `
resource "nios_misc_datacollectioncluster" "test" {
}

data "nios_misc_datacollectioncluster" "test" {
  filters = {
	 = nios_misc_datacollectioncluster.test.
  }
}
`
}
