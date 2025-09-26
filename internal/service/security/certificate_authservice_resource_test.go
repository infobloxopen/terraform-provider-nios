package security_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// TODO: Objects required to be set in the grid
// cacertificate to be uploaded to NIOS Grid
// Active Directory Service - active_dir , active_dir_test

var readableAttributesForCertificateAuthservice = "auto_populate_login,ca_certificates,comment,disabled,enable_password_request,enable_remote_lookup,max_retries,name,ocsp_check,ocsp_responders,recovery_interval,remote_lookup_service,remote_lookup_username,response_timeout,trust_model,user_match_type"

func TestAccCertificateAuthserviceResource_basic(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test"
	var v security.CertificateAuthservice
	name := acctest.RandomNameWithPrefix("certificate_authservice")
	caCertificate := []string{
		"cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceBasicConfig(name, caCertificate, "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "ca_certificates.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ca_certificates.0", caCertificate[0]),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_password_request", "true"),
					resource.TestCheckResourceAttr(resourceName, "enable_remote_lookup", "false"),
					resource.TestCheckResourceAttr(resourceName, "auto_populate_login", "S_DN_CN"),
					resource.TestCheckResourceAttr(resourceName, "max_retries", "0"),
					resource.TestCheckResourceAttr(resourceName, "recovery_interval", "30"),
					resource.TestCheckResourceAttr(resourceName, "response_timeout", "1000"),
					resource.TestCheckResourceAttr(resourceName, "trust_model", "DIRECT"),
					resource.TestCheckResourceAttr(resourceName, "user_match_type", "AUTO_MATCH"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_disappears(t *testing.T) {
	t.Skip("skipping ")
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
				Config: testAccCertificateAuthserviceBasicConfig(name, caCertificate, "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					testAccCheckCertificateAuthserviceDisappears(context.Background(), &v),
				),
			},
		},
	})
}

func TestAccCertificateAuthserviceResource_AutoPopulateLogin(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_auto_populate_login"
	var v security.CertificateAuthservice
	name := acctest.RandomNameWithPrefix("certificate_authservice")
	caCertificate := []string{
		"cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceAutoPopulateLogin(name, caCertificate, "SAN_EMAIL", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_populate_login", "SAN_EMAIL"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceAutoPopulateLogin(name, caCertificate, "S_DN_CN", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_populate_login", "S_DN_CN"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_CaCertificates(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_ca_certificates"
	var v security.CertificateAuthservice
	name := acctest.RandomNameWithPrefix("certificate_authservice")
	caCertificate := []string{
		"cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22",
	}
	caCertificateUpdate := []string{
		"cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuZGM2MTlhMWYyYmI0NGYwYjUzMWFiNzcwZjk1ZDQ0MDRhNWY2ODQxZGQxOTQ3Y2Q0YjcxMjU1YWU1MjY5MzM1MTRhMDljNWI5OTMwNmNhYzRiMjczY2JhN2NhODYwOWQ5ODY2YWYxYzU3NDdkNTVmNTFjZjM0ZGY4NzRmYTFjZWU:CN%3D%22ib-root-ca%22",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceCaCertificates(name, caCertificate, "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_certificates.0", caCertificate[0]),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceCaCertificates(name, caCertificateUpdate, "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_certificates.0", caCertificateUpdate[0]),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_Comment(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_comment"
	var v security.CertificateAuthservice
	name := acctest.RandomNameWithPrefix("certificate_authservice")
	caCertificate := []string{
		"cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceComment(name, caCertificate, "This is a comment", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceComment(name, caCertificate, "This is an updated comment", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_Disabled(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_disabled"
	var v security.CertificateAuthservice
	name := acctest.RandomNameWithPrefix("certificate_authservice")
	caCertificate := []string{
		"cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceDisabled(name, caCertificate, "true", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceDisabled(name, caCertificate, "false", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_EnablePasswordRequest(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_enable_password_request"
	var v security.CertificateAuthservice
	name := acctest.RandomNameWithPrefix("certificate_authservice")
	caCertificate := []string{
		"cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceEnablePasswordRequest(name, caCertificate, "true", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_password_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceEnablePasswordRequest(name, caCertificate, "false", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_password_request", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_EnableRemoteLookup(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_enable_remote_lookup"
	var v security.CertificateAuthservice
	name := acctest.RandomNameWithPrefix("certificate_authservice")
	caCertificate := []string{
		"cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22",
	}
	remoteLookupService := "ad_auth_service/b25lLmFkX2F1dGhfc2VydmljZSRhY3RpdmVfZGly:active_dir"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceEnableRemoteLookup(name, caCertificate, remoteLookupService, "false", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_remote_lookup", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceEnableRemoteLookupUpdate(name, caCertificate, remoteLookupService, "true", "admin", "infoblox", "false", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_remote_lookup", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_MaxRetries(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_max_retries"
	var v security.CertificateAuthservice
	name := acctest.RandomNameWithPrefix("certificate_authservice")
	caCertificate := []string{
		"cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceMaxRetries(name, caCertificate, "4", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_retries", "4"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceMaxRetries(name, caCertificate, "5", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_retries", "5"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_Name(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_name"
	var v security.CertificateAuthservice
	name := acctest.RandomNameWithPrefix("certificate_authservice")
	nameUpdate := acctest.RandomNameWithPrefix("certificate_authservice")
	caCertificate := []string{
		"cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceName(name, caCertificate, "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceName(nameUpdate, caCertificate, "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_OcspCheck(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_ocsp_check"
	var v security.CertificateAuthservice
	name := acctest.RandomNameWithPrefix("certificate_authservice")
	caCertificate := []string{
		"cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22",
	}
	testDataPath := getTestDataPath()
	ocspResponders := []map[string]any{
		{
			"fqdn_or_ip":            "2.2.2.2",
			"certificate_file_path": filepath.Join(testDataPath, "cert.pem"),
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceOcspCheck(name, caCertificate, "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ocsp_check", "DISABLED"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceOcspCheckUpdate(name, caCertificate, "MANUAL", ocspResponders),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ocsp_check", "MANUAL"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_OcspResponders(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_ocsp_responders"
	var v security.CertificateAuthservice
	name := acctest.RandomNameWithPrefix("certificate_authservice")
	caCertificate := []string{
		"cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22",
	}

	testDataPath := getTestDataPath()
	ocspResponders := []map[string]any{
		{
			"fqdn_or_ip":            "3.3.3.3",
			"certificate_file_path": filepath.Join(testDataPath, "cert.pem"),
		},
		{
			"fqdn_or_ip":            "3.3.32.3",
			"certificate_file_path": filepath.Join(testDataPath, "client.cert.pem"),
		},
	}

	ocspRespondersUpdate := []map[string]any{
		{
			"fqdn_or_ip":            "3.3.32.3",
			"certificate_file_path": "/Users/chaithra/go/src/github.com/infobloxopen/terraform-provider-nios/internal/testdata/nios_security_certificate_authservice/client.cert.pem",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceOcspResponders(name, caCertificate, ocspResponders),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ocsp_responders.0.fqdn_or_ip", "3.3.3.3"),
					resource.TestCheckResourceAttr(resourceName, "ocsp_responders.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "ocsp_responders.0.disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "ocsp_responders.0.port", "80"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceOcspResponders(name, caCertificate, ocspRespondersUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ocsp_responders.0.fqdn_or_ip", "3.3.32.3"),
					resource.TestCheckResourceAttr(resourceName, "ocsp_responders.0.disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "ocsp_responders.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ocsp_responders.0.port", "80"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_RecoveryInterval(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_recovery_interval"
	var v security.CertificateAuthservice
	name := acctest.RandomNameWithPrefix("certificate_authservice")
	caCertificate := []string{
		"cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceRecoveryInterval(name, caCertificate, "3", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recovery_interval", "3"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceRecoveryInterval(name, caCertificate, "5", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recovery_interval", "5"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_RemoteLookupService(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_remote_lookup_service"
	var v security.CertificateAuthservice
	name := acctest.RandomNameWithPrefix("certificate_authservice")
	caCertificate := []string{
		"cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22",
	}
	remoteLookupService := "ad_auth_service/b25lLmFkX2F1dGhfc2VydmljZSRhY3RpdmVfZGly:active_dir"
	remoteLookupServiceUpdate := "ad_auth_service/b25lLmFkX2F1dGhfc2VydmljZSRhY3RpdmVfZGlyX3Rlc3Q:active_dir_test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceRemoteLookupService(name, caCertificate, remoteLookupService, "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_lookup_service", remoteLookupService),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceRemoteLookupService(name, caCertificate, remoteLookupServiceUpdate, "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_lookup_service", remoteLookupServiceUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_RemoteLookupUsername(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_remote_lookup_username"
	var v security.CertificateAuthservice
	name := acctest.RandomNameWithPrefix("certificate_authservice")
	caCertificate := []string{
		"cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceRemoteLookupUsername(name, caCertificate, "username1", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_lookup_username", "username1"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceRemoteLookupUsername(name, caCertificate, "username2", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_lookup_username", "username2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_ResponseTimeout(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_response_timeout"
	var v security.CertificateAuthservice
	name := acctest.RandomNameWithPrefix("certificate_authservice")
	caCertificate := []string{
		"cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceResponseTimeout(name, caCertificate, "3000", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "response_timeout", "3000"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceResponseTimeout(name, caCertificate, "5000", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "response_timeout", "5000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_TrustModel(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_trust_model"
	var v security.CertificateAuthservice
	name := acctest.RandomNameWithPrefix("certificate_authservice")
	caCertificate := []string{
		"cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceTrustModel(name, caCertificate, "DELEGATED", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "trust_model", "DELEGATED"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceTrustModel(name, caCertificate, "DIRECT", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "trust_model", "DIRECT"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_UserMatchType(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_user_match_type"
	var v security.CertificateAuthservice
	name := acctest.RandomNameWithPrefix("certificate_authservice")
	caCertificate := []string{
		"cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceUserMatchType(name, caCertificate, "DIRECT_MATCH", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "user_match_type", "DIRECT_MATCH"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceUserMatchType(name, caCertificate, "AUTO_MATCH", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "user_match_type", "AUTO_MATCH"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckCertificateAuthserviceExists(ctx context.Context, resourceName string, v *security.CertificateAuthservice) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SecurityAPI.
			CertificateAuthserviceAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForCertificateAuthservice).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetCertificateAuthserviceResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetCertificateAuthserviceResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckCertificateAuthserviceDestroy(ctx context.Context, v *security.CertificateAuthservice) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SecurityAPI.
			CertificateAuthserviceAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForCertificateAuthservice).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				// resource was deleted
				return nil
			}
			return err
		}
		return errors.New("expected to be deleted")
	}
}

func testAccCheckCertificateAuthserviceDisappears(ctx context.Context, v *security.CertificateAuthservice) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SecurityAPI.
			CertificateAuthserviceAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCertificateAuthserviceBasicConfig(name string, caCertificate []string, ocspCheck string) string {
	caCertificateStr := utils.ConvertStringSliceToHCL(caCertificate)
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test" {
	ocsp_check = %q
    name = %q
    ca_certificates = %s
}
`, ocspCheck, name, caCertificateStr)
}

func testAccCertificateAuthserviceAutoPopulateLogin(name string, caCertificate []string, autoPopulateLogin, ocspCheck string) string {
	caCertificateStr := utils.ConvertStringSliceToHCL(caCertificate)
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_auto_populate_login" {
	ocsp_check = %q	
	name = %q
	ca_certificates = %s
	auto_populate_login = %q
}
`, ocspCheck, name, caCertificateStr, autoPopulateLogin)
}

func testAccCertificateAuthserviceCaCertificates(name string, caCertificates []string, ocspCheck string) string {
	caCertificatesStr := utils.ConvertStringSliceToHCL(caCertificates)
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_ca_certificates" {
	ocsp_check = %q       
    name = %q
    ca_certificates = %s
}
`, ocspCheck, name, caCertificatesStr)
}

func testAccCertificateAuthserviceComment(name string, cacertificate []string, comment, ocspCheck string) string {
	caCertificateStr := utils.ConvertStringSliceToHCL(cacertificate)
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_comment" {
	ocsp_check = %q
    name = %q
    ca_certificates = %s
    comment = %q
}
`, ocspCheck, name, caCertificateStr, comment)
}

func testAccCertificateAuthserviceDisabled(name string, caCertificate []string, disabled, ocspCheck string) string {
	caCertificateStr := utils.ConvertStringSliceToHCL(caCertificate)
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_disabled" {
	ocsp_check = %q
    name = %q
    ca_certificates = %s
    disabled = %q
}
`, ocspCheck, name, caCertificateStr, disabled)
}

func testAccCertificateAuthserviceEnablePasswordRequest(name string, caCertificate []string, enablePasswordRequest string, ocspCheck string) string {
	caCertificateStr := utils.ConvertStringSliceToHCL(caCertificate)
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_enable_password_request" {
	ocsp_check = %q
    enable_password_request = %q
    name = %q
    ca_certificates = %s
}
`, ocspCheck, enablePasswordRequest, name, caCertificateStr)
}

func testAccCertificateAuthserviceEnableRemoteLookup(name string, caCertificate []string, enableLookupService, enableRemoteLookup, ocspCheck string) string {
	caCertificateStr := utils.ConvertStringSliceToHCL(caCertificate)
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_enable_remote_lookup" {
	ocsp_check = %q
	remote_lookup_service = %q
    enable_remote_lookup = %q
    name = %q
    ca_certificates = %s
}
`, ocspCheck, enableLookupService, enableRemoteLookup, name, caCertificateStr)
}

func testAccCertificateAuthserviceEnableRemoteLookupUpdate(name string, caCertificate []string, enableLookupService, enableRemoteLookup string, remoteLookupUsername, remoteLookupPassword, enablePasswordRequest, ocspCheck string) string {
	caCertificateStr := utils.ConvertStringSliceToHCL(caCertificate)
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_enable_remote_lookup" {
	ocsp_check = "DISABLED"
	remote_lookup_service = %q
    enable_remote_lookup = %q
    name = %q
    ca_certificates = %s
    remote_lookup_username = %q
    remote_lookup_password = %q
	enable_password_request = %q
}
`, enableLookupService, enableRemoteLookup, name, caCertificateStr, remoteLookupUsername, remoteLookupPassword, enablePasswordRequest)
}

func testAccCertificateAuthserviceMaxRetries(name string, caCertificate []string, maxRetries, ocspCheck string) string {
	caCertificateStr := utils.ConvertStringSliceToHCL(caCertificate)
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_max_retries" {
	ocsp_check = %q
    max_retries = %q
    name = %q
    ca_certificates = %s
}
`, ocspCheck, maxRetries, name, caCertificateStr)
}

func testAccCertificateAuthserviceName(name string, caCertificate []string, ocspCheck string) string {
	caCertificateStr := utils.ConvertStringSliceToHCL(caCertificate)
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_name" {
    name = %q
    ca_certificates = %s
	ocsp_check = %q
}
`, name, caCertificateStr, ocspCheck)
}

func testAccCertificateAuthserviceOcspCheck(name string, caCertificate []string, ocspCheck string) string {
	caCertificateStr := utils.ConvertStringSliceToHCL(caCertificate)
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_ocsp_check" {
    ocsp_check = %q
    name = %q
    ca_certificates = %s
}
`, ocspCheck, name, caCertificateStr)
}

func testAccCertificateAuthserviceOcspCheckUpdate(name string, caCertificate []string, ocspCheck string, ocspResponders []map[string]any) string {
	ocspRespondersStr := utils.ConvertSliceOfMapsToHCL(ocspResponders)
	caCertificateStr := utils.ConvertStringSliceToHCL(caCertificate)
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_ocsp_check" {
	name = %q
	ca_certificates = %s
    ocsp_check = %q
    ocsp_responders = %s
}
`, name, caCertificateStr, ocspCheck, ocspRespondersStr)
}

func testAccCertificateAuthserviceOcspResponders(name string, caCertificates []string, ocspResponders []map[string]any) string {
	caCertificatesStr := utils.ConvertStringSliceToHCL(caCertificates)
	ocspRespondersStr := utils.ConvertSliceOfMapsToHCL(ocspResponders)
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_ocsp_responders" {
	name = %q
    ocsp_responders = %s
    ca_certificates = %s
}
`, name, ocspRespondersStr, caCertificatesStr)
}

func testAccCertificateAuthserviceRecoveryInterval(name string, caCertificate []string, recoveryInterval, ocspCheck string) string {
	caCertificateStr := utils.ConvertStringSliceToHCL(caCertificate)
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_recovery_interval" {
	ocsp_check = %q
	name = %q
	ca_certificates = %s
	recovery_interval = %q
}
`, ocspCheck, name, caCertificateStr, recoveryInterval)
}

func testAccCertificateAuthserviceRemoteLookupService(name string, caCertificate []string, remoteLookupService, ocspCheck string) string {
	caCertificateStr := utils.ConvertStringSliceToHCL(caCertificate)
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_remote_lookup_service" {
	ocsp_check = %q
	name = %q
    remote_lookup_service = %q
    ca_certificates = %s
}
`, ocspCheck, name, remoteLookupService, caCertificateStr)
}

func testAccCertificateAuthserviceRemoteLookupUsername(name string, caCertificate []string, remoteLookupUsername, ocspCheck string) string {
	caCertificateStr := utils.ConvertStringSliceToHCL(caCertificate)
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_remote_lookup_username" {
	ocsp_check = %q
	name = %q
    remote_lookup_username = %q
    ca_certificates = %s
}
`, ocspCheck, name, remoteLookupUsername, caCertificateStr)
}

func testAccCertificateAuthserviceResponseTimeout(name string, caCertificate []string, responseTimeout, ocspCheck string) string {
	caCertificateStr := utils.ConvertStringSliceToHCL(caCertificate)
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_response_timeout" {
	ocsp_check = %q
    response_timeout = %q
    name = %q
    ca_certificates = %s
}
`, ocspCheck, responseTimeout, name, caCertificateStr)
}

func testAccCertificateAuthserviceTrustModel(name string, caCertificate []string, trustModel , ocspCheck string) string {
	caCertificateStr := utils.ConvertStringSliceToHCL(caCertificate)
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_trust_model" {
	ocsp_check = %q
    name = %q
    ca_certificates = %s
    trust_model = %q
}
`, ocspCheck, name, caCertificateStr, trustModel)
}

func testAccCertificateAuthserviceUserMatchType(name string, caCertificate []string, userMatchType ,ocspCheck string) string {
	caCertificateStr := utils.ConvertStringSliceToHCL(caCertificate)
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_user_match_type" {
	ocsp_check = %q
    user_match_type = %q
    ca_certificates = %s
    name = %q
}
`, ocspCheck, userMatchType, caCertificateStr, name)
}

func getTestDataPath() string {
	wd, err := os.Getwd()
	if err != nil {
		return "../../testdata/nios_security_certificate_authservice"
	}
	return filepath.Join(wd, "../../testdata/nios_security_certificate_authservice")
}
