
package security_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccCertificateAuthserviceDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_security_certificate_authservice.test"
	resourceName := "nios_security_certificate_authservice.test"
	var v security.CertificateAuthservice

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCertificateAuthserviceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccCertificateAuthserviceDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
						}, testAccCheckCertificateAuthserviceResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccCertificateAuthserviceDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_security_certificate_authservice.test"
	resourceName := "nios_security_certificate_authservice.test"
	var v security.CertificateAuthservice
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCertificateAuthserviceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccCertificateAuthserviceDataSourceConfigExtAttrFilters( acctest.RandomName()),
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

func testAccCheckCertificateAuthserviceResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
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

func testAccCertificateAuthserviceDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test" {
}

data "nios_security_certificate_authservice" "test" {
  filters = {
	 = nios_security_certificate_authservice.test.
  }
}
`)
}

func testAccCertificateAuthserviceDataSourceConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test" {
  extattrs = {
    Site = %q
  } 
}

data "nios_security_certificate_authservice" "test" {
  extattrfilters = {
	Site = nios_security_certificate_authservice.test.extattrs.Site
  }
}
`,extAttrsValue)
}

