package security_test

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

func TestAccSamlAuthserviceDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_security_saml_authservice.test"
	resourceName := "nios_security_saml_authservice.test"
	var v security.SamlAuthservice
	name := acctest.RandomNameWithPrefix("saml_authservice")
	testDataPath := getSamlTestDataPath()
	idp := map[string]any{
		"idp_type":           "AZURE_SSO",
		"metadata_file_path": filepath.Join(testDataPath, "metadata.xml"),
		"sso_redirect_url":   "2.2.2.2",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSamlAuthserviceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSamlAuthserviceDataSourceConfigFilters(name, idp),
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

func testAccSamlAuthserviceDataSourceConfigFilters(name string, idp map[string]any) string {
	idpString := utils.ConvertMapToHCL(idp)
	return fmt.Sprintf(`
resource "nios_security_saml_authservice" "test" {
	name = %q
	idp = %s
}

data "nios_security_saml_authservice" "test" {
  filters = {
	name = nios_security_saml_authservice.test.name
  }
}
`, name, idpString)
}
