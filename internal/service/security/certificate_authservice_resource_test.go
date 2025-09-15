package security_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForCertificateAuthservice = "auto_populate_login,ca_certificates,comment,disabled,enable_password_request,enable_remote_lookup,max_retries,name,ocsp_check,ocsp_responders,recovery_interval,remote_lookup_service,remote_lookup_username,response_timeout,trust_model,user_match_type"

func TestAccCertificateAuthserviceResource_basic(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test"
	var v security.CertificateAuthservice
	name := acctest.RandomNameWithPrefix("certificate_authservice")
	caCertificate := []string{
		"cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceBasicConfig(name,caCertificate),
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
	resourceName := "nios_security_certificate_authservice.test"
	var v security.CertificateAuthservice
	name := acctest.RandomNameWithPrefix("certificate_authservice")
	caCertificate := []string{
		"cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCertificateAuthserviceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccCertificateAuthserviceBasicConfig(name, caCertificate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					testAccCheckCertificateAuthserviceDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceAutoPopulateLogin(name , caCertificate, "SAN_EMAIL"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_populate_login", "SAN_EMAIL"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceAutoPopulateLogin(name , caCertificate, "S_DN_CN"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_populate_login", "S_DN_EMAIL"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_CaCertificates(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_ca_certificates"
	var v security.CertificateAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceCaCertificates("CA_CERTIFICATES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_certificates", "CA_CERTIFICATES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceCaCertificates("CA_CERTIFICATES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_certificates", "CA_CERTIFICATES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_Comment(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_comment"
	var v security.CertificateAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_Disabled(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_disabled"
	var v security.CertificateAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceDisabled("DISABLED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "DISABLED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceDisabled("DISABLED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "DISABLED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_EnablePasswordRequest(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_enable_password_request"
	var v security.CertificateAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceEnablePasswordRequest("ENABLE_PASSWORD_REQUEST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_password_request", "ENABLE_PASSWORD_REQUEST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceEnablePasswordRequest("ENABLE_PASSWORD_REQUEST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_password_request", "ENABLE_PASSWORD_REQUEST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_EnableRemoteLookup(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_enable_remote_lookup"
	var v security.CertificateAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceEnableRemoteLookup("ENABLE_REMOTE_LOOKUP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_remote_lookup", "ENABLE_REMOTE_LOOKUP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceEnableRemoteLookup("ENABLE_REMOTE_LOOKUP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_remote_lookup", "ENABLE_REMOTE_LOOKUP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_MaxRetries(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_max_retries"
	var v security.CertificateAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceMaxRetries("MAX_RETRIES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_retries", "MAX_RETRIES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceMaxRetries("MAX_RETRIES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_retries", "MAX_RETRIES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_Name(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_name"
	var v security.CertificateAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_OcspCheck(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_ocsp_check"
	var v security.CertificateAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceOcspCheck("OCSP_CHECK_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ocsp_check", "OCSP_CHECK_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceOcspCheck("OCSP_CHECK_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ocsp_check", "OCSP_CHECK_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_OcspResponders(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_ocsp_responders"
	var v security.CertificateAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceOcspResponders("OCSP_RESPONDERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ocsp_responders", "OCSP_RESPONDERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceOcspResponders("OCSP_RESPONDERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ocsp_responders", "OCSP_RESPONDERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_RecoveryInterval(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_recovery_interval"
	var v security.CertificateAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceRecoveryInterval("RECOVERY_INTERVAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recovery_interval", "RECOVERY_INTERVAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceRecoveryInterval("RECOVERY_INTERVAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recovery_interval", "RECOVERY_INTERVAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_RemoteLookupPassword(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_remote_lookup_password"
	var v security.CertificateAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceRemoteLookupPassword("REMOTE_LOOKUP_PASSWORD_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_lookup_password", "REMOTE_LOOKUP_PASSWORD_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceRemoteLookupPassword("REMOTE_LOOKUP_PASSWORD_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_lookup_password", "REMOTE_LOOKUP_PASSWORD_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_RemoteLookupService(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_remote_lookup_service"
	var v security.CertificateAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceRemoteLookupService("REMOTE_LOOKUP_SERVICE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_lookup_service", "REMOTE_LOOKUP_SERVICE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceRemoteLookupService("REMOTE_LOOKUP_SERVICE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_lookup_service", "REMOTE_LOOKUP_SERVICE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_RemoteLookupUsername(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_remote_lookup_username"
	var v security.CertificateAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceRemoteLookupUsername("REMOTE_LOOKUP_USERNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_lookup_username", "REMOTE_LOOKUP_USERNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceRemoteLookupUsername("REMOTE_LOOKUP_USERNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_lookup_username", "REMOTE_LOOKUP_USERNAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_ResponseTimeout(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_response_timeout"
	var v security.CertificateAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceResponseTimeout("RESPONSE_TIMEOUT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "response_timeout", "RESPONSE_TIMEOUT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceResponseTimeout("RESPONSE_TIMEOUT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "response_timeout", "RESPONSE_TIMEOUT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_TrustModel(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_trust_model"
	var v security.CertificateAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceTrustModel("TRUST_MODEL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "trust_model", "TRUST_MODEL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceTrustModel("TRUST_MODEL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "trust_model", "TRUST_MODEL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCertificateAuthserviceResource_UserMatchType(t *testing.T) {
	var resourceName = "nios_security_certificate_authservice.test_user_match_type"
	var v security.CertificateAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCertificateAuthserviceUserMatchType("USER_MATCH_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "user_match_type", "USER_MATCH_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccCertificateAuthserviceUserMatchType("USER_MATCH_TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCertificateAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "user_match_type", "USER_MATCH_TYPE_UPDATE_REPLACE_ME"),
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

func testAccCertificateAuthserviceBasicConfig(name string, caCertificate []string) string {
	caCertificateStr := utils.ConvertStringSliceToHCL(caCertificate)
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test" {
    name = %q
    ca_certificates = %s
}
`, name, caCertificateStr)
}

func testAccCertificateAuthserviceRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccCertificateAuthserviceAutoPopulateLogin(name string, caCertificate []string ,autoPopulateLogin string ) string {
	caCertificateStr := utils.ConvertStringSliceToHCL(caCertificate)
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_auto_populate_login" {
	name = %q
	ca_certificates = %s
	auto_populate_login = %q
}
`, name, caCertificateStr, autoPopulateLogin)
}

func testAccCertificateAuthserviceCaCertificates(caCertificates string) string {
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_ca_certificates" {
    ca_certificates = %q
}
`, caCertificates)
}

func testAccCertificateAuthserviceComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccCertificateAuthserviceDisabled(disabled string) string {
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_disabled" {
    disabled = %q
}
`, disabled)
}

func testAccCertificateAuthserviceEnablePasswordRequest(enablePasswordRequest string) string {
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_enable_password_request" {
    enable_password_request = %q
}
`, enablePasswordRequest)
}

func testAccCertificateAuthserviceEnableRemoteLookup(enableRemoteLookup string) string {
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_enable_remote_lookup" {
    enable_remote_lookup = %q
}
`, enableRemoteLookup)
}

func testAccCertificateAuthserviceMaxRetries(maxRetries string) string {
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_max_retries" {
    max_retries = %q
}
`, maxRetries)
}

func testAccCertificateAuthserviceName(name string) string {
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_name" {
    name = %q
}
`, name)
}

func testAccCertificateAuthserviceOcspCheck(ocspCheck string) string {
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_ocsp_check" {
    ocsp_check = %q
}
`, ocspCheck)
}

func testAccCertificateAuthserviceOcspResponders(ocspResponders string) string {
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_ocsp_responders" {
    ocsp_responders = %q
}
`, ocspResponders)
}

func testAccCertificateAuthserviceRecoveryInterval(recoveryInterval string) string {
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_recovery_interval" {
    recovery_interval = %q
}
`, recoveryInterval)
}

func testAccCertificateAuthserviceRemoteLookupPassword(remoteLookupPassword string) string {
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_remote_lookup_password" {
    remote_lookup_password = %q
}
`, remoteLookupPassword)
}

func testAccCertificateAuthserviceRemoteLookupService(remoteLookupService string) string {
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_remote_lookup_service" {
    remote_lookup_service = %q
}
`, remoteLookupService)
}

func testAccCertificateAuthserviceRemoteLookupUsername(remoteLookupUsername string) string {
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_remote_lookup_username" {
    remote_lookup_username = %q
}
`, remoteLookupUsername)
}

func testAccCertificateAuthserviceResponseTimeout(responseTimeout string) string {
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_response_timeout" {
    response_timeout = %q
}
`, responseTimeout)
}

func testAccCertificateAuthserviceTrustModel(trustModel string) string {
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_trust_model" {
    trust_model = %q
}
`, trustModel)
}

func testAccCertificateAuthserviceUserMatchType(userMatchType string) string {
	return fmt.Sprintf(`
resource "nios_security_certificate_authservice" "test_user_match_type" {
    user_match_type = %q
}
`, userMatchType)
}
