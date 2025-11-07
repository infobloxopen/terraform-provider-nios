package security_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccHsmEntrustnshieldgroupDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_security_hsm_entrustnshieldgroup.test"
	resourceName := "nios_security_hsm_entrustnshieldgroup.test"
	var v security.HsmEntrustnshieldgroup

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHsmEntrustnshieldgroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccHsmEntrustnshieldgroupDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					}, testAccCheckHsmEntrustnshieldgroupResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckHsmEntrustnshieldgroupResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "card_name", dataSourceName, "result.0.card_name"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "entrustnshield_hsm", dataSourceName, "result.0.entrustnshield_hsm"),
		resource.TestCheckResourceAttrPair(resourceName, "key_server_ip", dataSourceName, "result.0.key_server_ip"),
		resource.TestCheckResourceAttrPair(resourceName, "key_server_port", dataSourceName, "result.0.key_server_port"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "pass_phrase", dataSourceName, "result.0.pass_phrase"),
		resource.TestCheckResourceAttrPair(resourceName, "protection", dataSourceName, "result.0.protection"),
		resource.TestCheckResourceAttrPair(resourceName, "status", dataSourceName, "result.0.status"),
	}
}

func testAccHsmEntrustnshieldgroupDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_security_hsm_entrustnshieldgroup" "test" {
}

data "nios_security_hsm_entrustnshieldgroup" "test" {
  filters = {
	 = nios_security_hsm_entrustnshieldgroup.test.
  }
}
`)
}
