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

var readableAttributesForAdminuser = "admin_groups,auth_method,auth_type,ca_certificate_issuer,client_certificate_serial_number,comment,disable,email,enable_certificate_authentication,extattrs,name,ssh_keys,status,time_zone,use_ssh_keys,use_time_zone"

func TestAccAdminuserResource_basic(t *testing.T) {
	var resourceName = "nios_security_admin_user.test"
	var v security.Adminuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_disappears(t *testing.T) {
	resourceName := "nios_security_admin_user.test"
	var v security.Adminuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAdminuserDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAdminuserBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					testAccCheckAdminuserDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAdminuserResource_Ref(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_ref"
	var v security.Adminuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_AdminGroups(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_admin_groups"
	var v security.Adminuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserAdminGroups("ADMIN_GROUPS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "admin_groups", "ADMIN_GROUPS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserAdminGroups("ADMIN_GROUPS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "admin_groups", "ADMIN_GROUPS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_AuthMethod(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_auth_method"
	var v security.Adminuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserAuthMethod("AUTH_METHOD_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_method", "AUTH_METHOD_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserAuthMethod("AUTH_METHOD_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_method", "AUTH_METHOD_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_AuthType(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_auth_type"
	var v security.Adminuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserAuthType("AUTH_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_type", "AUTH_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserAuthType("AUTH_TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_type", "AUTH_TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_CaCertificateIssuer(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_ca_certificate_issuer"
	var v security.Adminuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserCaCertificateIssuer("CA_CERTIFICATE_ISSUER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_certificate_issuer", "CA_CERTIFICATE_ISSUER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserCaCertificateIssuer("CA_CERTIFICATE_ISSUER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_certificate_issuer", "CA_CERTIFICATE_ISSUER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_ClientCertificateSerialNumber(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_client_certificate_serial_number"
	var v security.Adminuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserClientCertificateSerialNumber("CLIENT_CERTIFICATE_SERIAL_NUMBER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_certificate_serial_number", "CLIENT_CERTIFICATE_SERIAL_NUMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserClientCertificateSerialNumber("CLIENT_CERTIFICATE_SERIAL_NUMBER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_certificate_serial_number", "CLIENT_CERTIFICATE_SERIAL_NUMBER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_Comment(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_comment"
	var v security.Adminuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_Disable(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_disable"
	var v security.Adminuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_Email(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_email"
	var v security.Adminuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserEmail("EMAIL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email", "EMAIL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserEmail("EMAIL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email", "EMAIL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_EnableCertificateAuthentication(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_enable_certificate_authentication"
	var v security.Adminuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserEnableCertificateAuthentication("ENABLE_CERTIFICATE_AUTHENTICATION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_certificate_authentication", "ENABLE_CERTIFICATE_AUTHENTICATION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserEnableCertificateAuthentication("ENABLE_CERTIFICATE_AUTHENTICATION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_certificate_authentication", "ENABLE_CERTIFICATE_AUTHENTICATION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_extattrs"
	var v security.Adminuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_Name(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_name"
	var v security.Adminuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_Password(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_password"
	var v security.Adminuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserPassword("PASSWORD_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "password", "PASSWORD_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserPassword("PASSWORD_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "password", "PASSWORD_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_SshKeys(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_ssh_keys"
	var v security.Adminuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserSshKeys("SSH_KEYS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ssh_keys", "SSH_KEYS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserSshKeys("SSH_KEYS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ssh_keys", "SSH_KEYS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_TimeZone(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_time_zone"
	var v security.Adminuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserTimeZone("TIME_ZONE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "time_zone", "TIME_ZONE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserTimeZone("TIME_ZONE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "time_zone", "TIME_ZONE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_UseSshKeys(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_use_ssh_keys"
	var v security.Adminuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserUseSshKeys("USE_SSH_KEYS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ssh_keys", "USE_SSH_KEYS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserUseSshKeys("USE_SSH_KEYS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ssh_keys", "USE_SSH_KEYS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_UseTimeZone(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_use_time_zone"
	var v security.Adminuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserUseTimeZone("USE_TIME_ZONE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_time_zone", "USE_TIME_ZONE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserUseTimeZone("USE_TIME_ZONE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_time_zone", "USE_TIME_ZONE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckAdminuserExists(ctx context.Context, resourceName string, v *security.Adminuser) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SecurityAPI.
			AdminuserAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForAdminuser).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetAdminuserResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetAdminuserResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckAdminuserDestroy(ctx context.Context, v *security.Adminuser) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SecurityAPI.
			AdminuserAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForAdminuser).
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

func testAccCheckAdminuserDisappears(ctx context.Context, v *security.Adminuser) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SecurityAPI.
			AdminuserAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccAdminuserBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test" {
}
`)
}

func testAccAdminuserRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccAdminuserAdminGroups(adminGroups string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_admin_groups" {
    admin_groups = %q
}
`, adminGroups)
}

func testAccAdminuserAuthMethod(authMethod string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_auth_method" {
    auth_method = %q
}
`, authMethod)
}

func testAccAdminuserAuthType(authType string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_auth_type" {
    auth_type = %q
}
`, authType)
}

func testAccAdminuserCaCertificateIssuer(caCertificateIssuer string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_ca_certificate_issuer" {
    ca_certificate_issuer = %q
}
`, caCertificateIssuer)
}

func testAccAdminuserClientCertificateSerialNumber(clientCertificateSerialNumber string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_client_certificate_serial_number" {
    client_certificate_serial_number = %q
}
`, clientCertificateSerialNumber)
}

func testAccAdminuserComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccAdminuserDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccAdminuserEmail(email string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_email" {
    email = %q
}
`, email)
}

func testAccAdminuserEnableCertificateAuthentication(enableCertificateAuthentication string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_enable_certificate_authentication" {
    enable_certificate_authentication = %q
}
`, enableCertificateAuthentication)
}

func testAccAdminuserExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccAdminuserName(name string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_name" {
    name = %q
}
`, name)
}

func testAccAdminuserPassword(password string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_password" {
    password = %q
}
`, password)
}

func testAccAdminuserSshKeys(sshKeys string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_ssh_keys" {
    ssh_keys = %q
}
`, sshKeys)
}

func testAccAdminuserTimeZone(timeZone string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_time_zone" {
    time_zone = %q
}
`, timeZone)
}

func testAccAdminuserUseSshKeys(useSshKeys string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_use_ssh_keys" {
    use_ssh_keys = %q
}
`, useSshKeys)
}

func testAccAdminuserUseTimeZone(useTimeZone string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_use_time_zone" {
    use_time_zone = %q
}
`, useTimeZone)
}
