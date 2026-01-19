package security_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

func TestAccCertificateAuthserviceDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_security_certificate_authservice.test"
	resourceName := "nios_security_certificate_authservice.test"
	var v security.CertificateAuthservice
	name := acctest.RandomNameWithPrefix("certificate_authservice")
	caCertificate := []string{
		"cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCertificateAuthserviceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccCertificateAuthserviceDataSourceConfigFilters(name, caCertificate, "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					}, testAccCheckCertificateAuthserviceResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckCertificateAuthserviceResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "auto_populate_login", dataSourceName, "result.0.auto_populate_login"),
		resource.TestCheckResourceAttrPair(resourceName, "ca_certificates", dataSourceName, "result.0.ca_certificates"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disabled", dataSourceName, "result.0.disabled"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_password_request", dataSourceName, "result.0.enable_password_request"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_remote_lookup", dataSourceName, "result.0.enable_remote_lookup"),
		resource.TestCheckResourceAttrPair(resourceName, "max_retries", dataSourceName, "result.0.max_retries"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "ocsp_check", dataSourceName, "result.0.ocsp_check"),
		resource.TestCheckResourceAttrPair(resourceName, "ocsp_responders", dataSourceName, "result.0.ocsp_responders"),
		resource.TestCheckResourceAttrPair(resourceName, "recovery_interval", dataSourceName, "result.0.recovery_interval"),
		resource.TestCheckResourceAttrPair(resourceName, "remote_lookup_password", dataSourceName, "result.0.remote_lookup_password"),
		resource.TestCheckResourceAttrPair(resourceName, "remote_lookup_service", dataSourceName, "result.0.remote_lookup_service"),
		resource.TestCheckResourceAttrPair(resourceName, "remote_lookup_username", dataSourceName, "result.0.remote_lookup_username"),
		resource.TestCheckResourceAttrPair(resourceName, "response_timeout", dataSourceName, "result.0.response_timeout"),
		resource.TestCheckResourceAttrPair(resourceName, "trust_model", dataSourceName, "result.0.trust_model"),
		resource.TestCheckResourceAttrPair(resourceName, "user_match_type", dataSourceName, "result.0.user_match_type"),
	}
}

func testAccCertificateAuthserviceDataSourceConfigFilters(name string, caCertificate []string, ocspCheck string) string {
	caCertificateStr := utils.ConvertStringSliceToHCL(caCertificate)
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test" {
  name            = %q
  ca_certificates = %s
  ocsp_check     = %q
}

data "nios_security_certificate_authservice" "test" {
  filters = {
    name = nios_security_certificate_authservice.test.name
  }
}
`, name, caCertificateStr, ocspCheck)
}
