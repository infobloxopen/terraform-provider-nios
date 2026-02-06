package misc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccBfdtemplateDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_misc_bfdtemplate.test"
	resourceName := "nios_misc_bfdtemplate.test"
	var v misc.Bfdtemplate
	name := acctest.RandomNameWithPrefix("tf-test-bfdtemplate-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBfdtemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccBfdtemplateDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckBfdtemplateExists(context.Background(), resourceName, &v),
					}, testAccCheckBfdtemplateResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckBfdtemplateResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "authentication_key_id", dataSourceName, "result.0.authentication_key_id"),
		resource.TestCheckResourceAttrPair(resourceName, "authentication_type", dataSourceName, "result.0.authentication_type"),
		resource.TestCheckResourceAttrPair(resourceName, "detection_multiplier", dataSourceName, "result.0.detection_multiplier"),
		resource.TestCheckResourceAttrPair(resourceName, "min_rx_interval", dataSourceName, "result.0.min_rx_interval"),
		resource.TestCheckResourceAttrPair(resourceName, "min_tx_interval", dataSourceName, "result.0.min_tx_interval"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
	}
}

func testAccBfdtemplateDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_misc_bfdtemplate" "test" {
	name = %q
}

data "nios_misc_bfdtemplate" "test" {
  filters = {
	name = nios_misc_bfdtemplate.test.name
  }
}
`, name)
}
