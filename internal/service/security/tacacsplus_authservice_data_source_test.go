package security_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccTacacsplusAuthserviceDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_security_tacacsplus_authservice.test"
	resourceName := "nios_security_tacacsplus_authservice.test"
	var v security.TacacsplusAuthservice

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTacacsplusAuthserviceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccTacacsplusAuthserviceDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					}, testAccCheckTacacsplusAuthserviceResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckTacacsplusAuthserviceResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "acct_retries", dataSourceName, "result.0.acct_retries"),
		resource.TestCheckResourceAttrPair(resourceName, "acct_timeout", dataSourceName, "result.0.acct_timeout"),
		resource.TestCheckResourceAttrPair(resourceName, "auth_retries", dataSourceName, "result.0.auth_retries"),
		resource.TestCheckResourceAttrPair(resourceName, "auth_timeout", dataSourceName, "result.0.auth_timeout"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "servers", dataSourceName, "result.0.servers"),
	}
}

func testAccTacacsplusAuthserviceDataSourceConfigFilters() string {
	return `
resource "nios_security_tacacsplus_authservice" "test" {
}

data "nios_security_tacacsplus_authservice" "test" {
  filters = {
	 = nios_security_tacacsplus_authservice.test.
  }
}
`
}
