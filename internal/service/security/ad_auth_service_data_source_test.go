package security_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccAdAuthServiceDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_security_ad_auth_service.test"
	resourceName := "nios_security_ad_auth_service.test"
	var v security.AdAuthService

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAdAuthServiceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAdAuthServiceDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckAdAuthServiceExists(context.Background(), resourceName, &v),
					}, testAccCheckAdAuthServiceResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckAdAuthServiceResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "ad_domain", dataSourceName, "result.0.ad_domain"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disabled", dataSourceName, "result.0.disabled"),
		resource.TestCheckResourceAttrPair(resourceName, "domain_controllers", dataSourceName, "result.0.domain_controllers"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "nested_group_querying", dataSourceName, "result.0.nested_group_querying"),
		resource.TestCheckResourceAttrPair(resourceName, "timeout", dataSourceName, "result.0.timeout"),
	}
}

func testAccAdAuthServiceDataSourceConfigFilters() string {
	return `
resource "nios_security_ad_auth_service" "test" {
}

data "nios_security_ad_auth_service" "test" {
  filters = {
	 = nios_security_ad_auth_service.test.
  }
}
`
}
