package security_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccLdapAuthServiceDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_security_ldap_auth_service.test"
	resourceName := "nios_security_ldap_auth_service.test"
	var v security.LdapAuthService

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLdapAuthServiceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccLdapAuthServiceDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					}, testAccCheckLdapAuthServiceResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckLdapAuthServiceResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "ea_mapping", dataSourceName, "result.0.ea_mapping"),
		resource.TestCheckResourceAttrPair(resourceName, "ldap_group_attribute", dataSourceName, "result.0.ldap_group_attribute"),
		resource.TestCheckResourceAttrPair(resourceName, "ldap_group_authentication_type", dataSourceName, "result.0.ldap_group_authentication_type"),
		resource.TestCheckResourceAttrPair(resourceName, "ldap_user_attribute", dataSourceName, "result.0.ldap_user_attribute"),
		resource.TestCheckResourceAttrPair(resourceName, "mode", dataSourceName, "result.0.mode"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "recovery_interval", dataSourceName, "result.0.recovery_interval"),
		resource.TestCheckResourceAttrPair(resourceName, "retries", dataSourceName, "result.0.retries"),
		resource.TestCheckResourceAttrPair(resourceName, "search_scope", dataSourceName, "result.0.search_scope"),
		resource.TestCheckResourceAttrPair(resourceName, "servers", dataSourceName, "result.0.servers"),
		resource.TestCheckResourceAttrPair(resourceName, "timeout", dataSourceName, "result.0.timeout"),
	}
}

func testAccLdapAuthServiceDataSourceConfigFilters() string {
	return `
resource "nios_security_ldap_auth_service" "test" {
}

data "nios_security_ldap_auth_service" "test" {
  filters = {
	 = nios_security_ldap_auth_service.test.
  }
}
`
}
