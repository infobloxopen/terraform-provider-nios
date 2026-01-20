package security_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccSamlAuthserviceDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_security_saml_authservice.test"
	resourceName := "nios_security_saml_authservice.test"
	var v security.SamlAuthservice

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSamlAuthserviceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSamlAuthserviceDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					}, testAccCheckSamlAuthserviceResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckSamlAuthserviceResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "idp", dataSourceName, "result.0.idp"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "session_timeout", dataSourceName, "result.0.session_timeout"),
	}
}

func testAccSamlAuthserviceDataSourceConfigFilters() string {
	return `
resource "nios_security_saml_authservice" "test" {
}

data "nios_security_saml_authservice" "test" {
  filters = {
	 = nios_security_saml_authservice.test.
  }
}
`
}
