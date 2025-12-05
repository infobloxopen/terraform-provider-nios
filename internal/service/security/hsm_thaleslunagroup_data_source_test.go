package security_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccHsmThaleslunagroupDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_security_hsm_thaleslunagroup.test"
	resourceName := "nios_security_hsm_thaleslunagroup.test"
	var v security.HsmThaleslunagroup

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHsmThaleslunagroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccHsmThaleslunagroupDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					}, testAccCheckHsmThaleslunagroupResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckHsmThaleslunagroupResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "group_sn", dataSourceName, "result.0.group_sn"),
		resource.TestCheckResourceAttrPair(resourceName, "hsm_version", dataSourceName, "result.0.hsm_version"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "pass_phrase", dataSourceName, "result.0.pass_phrase"),
		resource.TestCheckResourceAttrPair(resourceName, "status", dataSourceName, "result.0.status"),
		resource.TestCheckResourceAttrPair(resourceName, "thalesluna", dataSourceName, "result.0.thalesluna"),
	}
}

func testAccHsmThaleslunagroupDataSourceConfigFilters() string {
	return `
resource "nios_security_hsm_thaleslunagroup" "test" {
}

data "nios_security_hsm_thaleslunagroup" "test" {
  filters = {
	 = nios_security_hsm_thaleslunagroup.test.
  }
}
`
}
