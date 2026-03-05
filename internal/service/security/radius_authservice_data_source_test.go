package security_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccRadiusAuthserviceDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_security_radius_authservice.test"
	resourceName := "nios_security_radius_authservice.test"
	var v security.RadiusAuthservice

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRadiusAuthserviceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRadiusAuthserviceDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					}, testAccCheckRadiusAuthserviceResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRadiusAuthserviceResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "acct_retries", dataSourceName, "result.0.acct_retries"),
		resource.TestCheckResourceAttrPair(resourceName, "acct_timeout", dataSourceName, "result.0.acct_timeout"),
		resource.TestCheckResourceAttrPair(resourceName, "auth_retries", dataSourceName, "result.0.auth_retries"),
		resource.TestCheckResourceAttrPair(resourceName, "auth_timeout", dataSourceName, "result.0.auth_timeout"),
		resource.TestCheckResourceAttrPair(resourceName, "cache_ttl", dataSourceName, "result.0.cache_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_cache", dataSourceName, "result.0.enable_cache"),
		resource.TestCheckResourceAttrPair(resourceName, "mode", dataSourceName, "result.0.mode"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "recovery_interval", dataSourceName, "result.0.recovery_interval"),
		resource.TestCheckResourceAttrPair(resourceName, "servers", dataSourceName, "result.0.servers"),
	}
}

func testAccRadiusAuthserviceDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test" {
}

data "nios_security_radius_authservice" "test" {
  filters = {
	 = nios_security_radius_authservice.test.
  }
}
`)
}
