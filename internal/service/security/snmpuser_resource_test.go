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

var readableAttributesForSnmpuser = "authentication_protocol,comment,disable,extattrs,name,privacy_protocol"

func TestAccSnmpuserResource_basic(t *testing.T) {
	var resourceName = "nios_security_snmpuser.test"
	var v security.Snmpuser

	name := acctest.RandomNameWithPrefix("example-snmpuser-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSnmpuserBasicConfig(name, "SHA", "abcd1234", "AES", "efgh5678"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "authentication_protocol", "SHA"),
					resource.TestCheckResourceAttr(resourceName, "authentication_password", "abcd1234"),
					resource.TestCheckResourceAttr(resourceName, "privacy_protocol", "AES"),
					resource.TestCheckResourceAttr(resourceName, "privacy_password", "efgh5678"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSnmpuserResource_disappears(t *testing.T) {
	resourceName := "nios_security_snmpuser.test"
	var v security.Snmpuser

	name := acctest.RandomNameWithPrefix("example-snmpuser-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpuserDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpuserBasicConfig(name, "SHA", "abcd1234", "AES", "efgh5678"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					testAccCheckSnmpuserDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSnmpuserResource_AuthenticationPassword(t *testing.T) {
	var resourceName = "nios_security_snmpuser.test_authentication_password"
	var v security.Snmpuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSnmpuserAuthenticationPassword("AUTHENTICATION_PASSWORD_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "authentication_password", "AUTHENTICATION_PASSWORD_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSnmpuserAuthenticationPassword("AUTHENTICATION_PASSWORD_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "authentication_password", "AUTHENTICATION_PASSWORD_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSnmpuserResource_AuthenticationProtocol(t *testing.T) {
	var resourceName = "nios_security_snmpuser.test_authentication_protocol"
	var v security.Snmpuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSnmpuserAuthenticationProtocol("AUTHENTICATION_PROTOCOL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "authentication_protocol", "AUTHENTICATION_PROTOCOL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSnmpuserAuthenticationProtocol("AUTHENTICATION_PROTOCOL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "authentication_protocol", "AUTHENTICATION_PROTOCOL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSnmpuserResource_Comment(t *testing.T) {
	var resourceName = "nios_security_snmpuser.test_comment"
	var v security.Snmpuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSnmpuserComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSnmpuserComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSnmpuserResource_Disable(t *testing.T) {
	var resourceName = "nios_security_snmpuser.test_disable"
	var v security.Snmpuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSnmpuserDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSnmpuserDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSnmpuserResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_security_snmpuser.test_extattrs"
	var v security.Snmpuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSnmpuserExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSnmpuserExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSnmpuserResource_Name(t *testing.T) {
	var resourceName = "nios_security_snmpuser.test_name"
	var v security.Snmpuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSnmpuserName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSnmpuserName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSnmpuserResource_PrivacyPassword(t *testing.T) {
	var resourceName = "nios_security_snmpuser.test_privacy_password"
	var v security.Snmpuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSnmpuserPrivacyPassword("PRIVACY_PASSWORD_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "privacy_password", "PRIVACY_PASSWORD_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSnmpuserPrivacyPassword("PRIVACY_PASSWORD_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "privacy_password", "PRIVACY_PASSWORD_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSnmpuserResource_PrivacyProtocol(t *testing.T) {
	var resourceName = "nios_security_snmpuser.test_privacy_protocol"
	var v security.Snmpuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSnmpuserPrivacyProtocol("PRIVACY_PROTOCOL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "privacy_protocol", "PRIVACY_PROTOCOL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSnmpuserPrivacyProtocol("PRIVACY_PROTOCOL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "privacy_protocol", "PRIVACY_PROTOCOL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckSnmpuserExists(ctx context.Context, resourceName string, v *security.Snmpuser) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SecurityAPI.
			SnmpuserAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForSnmpuser).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetSnmpuserResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetSnmpuserResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckSnmpuserDestroy(ctx context.Context, v *security.Snmpuser) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SecurityAPI.
			SnmpuserAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForSnmpuser).
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

func testAccCheckSnmpuserDisappears(ctx context.Context, v *security.Snmpuser) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SecurityAPI.
			SnmpuserAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccSnmpuserBasicConfig(name, authentication_protocol, authentication_password, privacy_protocol, privacy_password string) string {
	return fmt.Sprintf(`
resource "nios_security_snmpuser" "test" {
    name                 	= %q
    authentication_protocol = %q
	authentication_password = %q
    privacy_protocol     	= %q
    privacy_password     	= %q
}
`, name, authentication_protocol, authentication_password, privacy_protocol, privacy_password)
}

func testAccSnmpuserAuthenticationPassword(authenticationPassword string) string {
	return fmt.Sprintf(`
resource "nios_security_snmpuser" "test_authentication_password" {
    authentication_password = %q
}
`, authenticationPassword)
}

func testAccSnmpuserAuthenticationProtocol(authenticationProtocol string) string {
	return fmt.Sprintf(`
resource "nios_security_snmpuser" "test_authentication_protocol" {
    authentication_protocol = %q
}
`, authenticationProtocol)
}

func testAccSnmpuserComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_security_snmpuser" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccSnmpuserDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_security_snmpuser" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccSnmpuserExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_security_snmpuser" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccSnmpuserName(name string) string {
	return fmt.Sprintf(`
resource "nios_security_snmpuser" "test_name" {
    name = %q
}
`, name)
}

func testAccSnmpuserPrivacyPassword(privacyPassword string) string {
	return fmt.Sprintf(`
resource "nios_security_snmpuser" "test_privacy_password" {
    privacy_password = %q
}
`, privacyPassword)
}

func testAccSnmpuserPrivacyProtocol(privacyProtocol string) string {
	return fmt.Sprintf(`
resource "nios_security_snmpuser" "test_privacy_protocol" {
    privacy_protocol = %q
}
`, privacyProtocol)
}
