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
				Config: testAccSnmpuserBasicConfig(name, "NONE", "NONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "authentication_protocol", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "privacy_protocol", "NONE"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
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
				Config: testAccSnmpuserBasicConfig(name, "NONE", "NONE"),
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

	name := acctest.RandomNameWithPrefix("example-snmpuser-")
	auth_password := "AuthPassword@123"
	updated_auth_password := "UpdatedAuthPassword@123!"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSnmpuserAuthenticationPassword(name, "SHA", auth_password, "NONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "authentication_password", auth_password),
				),
			},
			// Update and Read
			{
				Config: testAccSnmpuserAuthenticationPassword(name, "SHA", updated_auth_password, "NONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "authentication_password", updated_auth_password),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSnmpuserResource_AuthenticationProtocol(t *testing.T) {
	var resourceName = "nios_security_snmpuser.test_authentication_protocol"
	var v security.Snmpuser

	name := acctest.RandomNameWithPrefix("example-snmpuser-")
	auth_password := "AuthPassword@123"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSnmpuserAuthenticationProtocol(name, "SHA", auth_password, "NONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "authentication_protocol", "SHA"),
				),
			},
			// Update and Read
			{
				Config: testAccSnmpuserAuthenticationProtocol(name, "MD5", auth_password, "NONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "authentication_protocol", "MD5"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSnmpuserResource_Comment(t *testing.T) {
	var resourceName = "nios_security_snmpuser.test_comment"
	var v security.Snmpuser

	name := acctest.RandomNameWithPrefix("example-snmpuser-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSnmpuserComment(name, "NONE", "NONE", "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccSnmpuserComment(name, "NONE", "NONE", "Updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSnmpuserResource_Disable(t *testing.T) {
	var resourceName = "nios_security_snmpuser.test_disable"
	var v security.Snmpuser

	name := acctest.RandomNameWithPrefix("example-snmpuser-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSnmpuserDisable(name, "NONE", "NONE", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSnmpuserDisable(name, "NONE", "NONE", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSnmpuserResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_security_snmpuser.test_extattrs"
	var v security.Snmpuser

	name := acctest.RandomNameWithPrefix("example-snmpuser-")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSnmpuserExtAttrs(name, "NONE", "NONE", map[string]string{"Site": extAttrValue1}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccSnmpuserExtAttrs(name, "NONE", "NONE", map[string]string{"Site": extAttrValue2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSnmpuserResource_Name(t *testing.T) {
	var resourceName = "nios_security_snmpuser.test_name"
	var v security.Snmpuser

	name1 := acctest.RandomNameWithPrefix("example-snmpuser-")
	name2 := acctest.RandomNameWithPrefix("updated-snmpuser-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSnmpuserName(name1, "NONE", "NONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccSnmpuserName(name2, "NONE", "NONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSnmpuserResource_PrivacyPassword(t *testing.T) {
	var resourceName = "nios_security_snmpuser.test_privacy_password"
	var v security.Snmpuser

	name := acctest.RandomNameWithPrefix("example-snmpuser-")
	auth_password := "AuthPassword@123"
	privacy_password := "PrivacyPassword@123"
	updated_privacy_password := "UpdatedPassword@123"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSnmpuserPrivacyPassword(name, "MD5", auth_password, "DES", privacy_password),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "privacy_password", privacy_password),
				),
			},
			// Update and Read
			{
				Config: testAccSnmpuserPrivacyPassword(name, "MD5", auth_password, "DES", updated_privacy_password),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "privacy_password", updated_privacy_password),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSnmpuserResource_PrivacyProtocol(t *testing.T) {
	var resourceName = "nios_security_snmpuser.test_privacy_protocol"
	var v security.Snmpuser

	name := acctest.RandomNameWithPrefix("example-snmpuser-")
	auth_password := "AuthPassword@123"
	privacy_password := "PrivacyPassword@123"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSnmpuserPrivacyProtocol(name, "MD5", auth_password, "AES", privacy_password),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "privacy_protocol", "AES"),
				),
			},
			// Update and Read
			{
				Config: testAccSnmpuserPrivacyProtocol(name, "MD5", auth_password, "DES", privacy_password),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "privacy_protocol", "DES"),
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

func testAccSnmpuserBasicConfig(name, authentication_protocol, privacy_protocol string) string {
	return fmt.Sprintf(`
resource "nios_security_snmpuser" "test" {
    name                 	= %q
    authentication_protocol = %q
    privacy_protocol     	= %q
}
`, name, authentication_protocol, privacy_protocol)
}

func testAccSnmpuserAuthenticationPassword(name, authentication_protocol, authenticationPassword, privacy_protocol string) string {
	return fmt.Sprintf(`
resource "nios_security_snmpuser" "test_authentication_password" {
	name                    = %q
	authentication_protocol = %q
    authentication_password = %q
    privacy_protocol        = %q
}
`, name, authentication_protocol, authenticationPassword, privacy_protocol)
}

func testAccSnmpuserAuthenticationProtocol(name, authenticationProtocol, authentication_password, privacy_protocol string) string {
	return fmt.Sprintf(`
resource "nios_security_snmpuser" "test_authentication_protocol" {
    name                    = %q
    authentication_protocol = %q
    authentication_password = %q
    privacy_protocol        = %q
}
`, name, authenticationProtocol, authentication_password, privacy_protocol)
}

func testAccSnmpuserComment(name, authentication_protocol, privacy_protocol, comment string) string {
	return fmt.Sprintf(`
resource "nios_security_snmpuser" "test_comment" {
    name                  	= %q
    authentication_protocol = %q
    privacy_protocol      	= %q
    comment = %q
}
`, name, authentication_protocol, privacy_protocol, comment)
}

func testAccSnmpuserDisable(name, authentication_protocol, privacy_protocol string, disable bool) string {
	return fmt.Sprintf(`
resource "nios_security_snmpuser" "test_disable" {
    name                  	= %q
    authentication_protocol = %q
    privacy_protocol        = %q
    disable = %t
}
`, name, authentication_protocol, privacy_protocol, disable)
}

func testAccSnmpuserExtAttrs(name, authentication_protocol, privacy_protocol string, extAttrs map[string]string) string {
	extattrsStr := "{"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`%s = %q`, k, v)
	}
	extattrsStr += "}"
	return fmt.Sprintf(`
resource "nios_security_snmpuser" "test_extattrs" {
	name                  	= %q
	authentication_protocol = %q
	privacy_protocol        = %q
    extattrs 				= %s
}
`, name, authentication_protocol, privacy_protocol, extattrsStr)
}

func testAccSnmpuserName(name, authentication_protocol, privacy_protocol string) string {
	return fmt.Sprintf(`
resource "nios_security_snmpuser" "test_name" {
    name                  	= %q
    authentication_protocol = %q
    privacy_protocol      	= %q
}
`, name, authentication_protocol, privacy_protocol)
}

func testAccSnmpuserPrivacyPassword(name, authentication_protocol, authentication_password, privacy_protocol, privacyPassword string) string {
	return fmt.Sprintf(`
resource "nios_security_snmpuser" "test_privacy_password" {
	name                   	= %q
	authentication_protocol = %q
	authentication_password = %q
	privacy_protocol        = %q
    privacy_password        = %q
}
`, name, authentication_protocol, authentication_password, privacy_protocol, privacyPassword)
}

func testAccSnmpuserPrivacyProtocol(name, authentication_protocol, authentication_password, privacy_protocol, privacy_password string) string {
	return fmt.Sprintf(`
resource "nios_security_snmpuser" "test_privacy_protocol" {
    name                    = %q
    authentication_protocol = %q
    authentication_password = %q
    privacy_protocol        = %q
	privacy_password        = %q
}
`, name, authentication_protocol, authentication_password, privacy_protocol, privacy_password)
}
