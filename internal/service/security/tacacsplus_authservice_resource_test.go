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

var readableAttributesForTacacsplusAuthservice = "acct_retries,acct_timeout,auth_retries,auth_timeout,comment,disable,name,servers"

func TestAccTacacsplusAuthserviceResource_basic(t *testing.T) {
	var resourceName = "nios_security_tacacsplus_authservice.test"
	var v security.TacacsplusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTacacsplusAuthserviceBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTacacsplusAuthserviceResource_disappears(t *testing.T) {
	resourceName := "nios_security_tacacsplus_authservice.test"
	var v security.TacacsplusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTacacsplusAuthserviceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccTacacsplusAuthserviceBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					testAccCheckTacacsplusAuthserviceDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccTacacsplusAuthserviceResource_Ref(t *testing.T) {
	var resourceName = "nios_security_tacacsplus_authservice.test_ref"
	var v security.TacacsplusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTacacsplusAuthserviceRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccTacacsplusAuthserviceRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTacacsplusAuthserviceResource_AcctRetries(t *testing.T) {
	var resourceName = "nios_security_tacacsplus_authservice.test_acct_retries"
	var v security.TacacsplusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTacacsplusAuthserviceAcctRetries("ACCT_RETRIES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "acct_retries", "ACCT_RETRIES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccTacacsplusAuthserviceAcctRetries("ACCT_RETRIES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "acct_retries", "ACCT_RETRIES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTacacsplusAuthserviceResource_AcctTimeout(t *testing.T) {
	var resourceName = "nios_security_tacacsplus_authservice.test_acct_timeout"
	var v security.TacacsplusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTacacsplusAuthserviceAcctTimeout("ACCT_TIMEOUT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "acct_timeout", "ACCT_TIMEOUT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccTacacsplusAuthserviceAcctTimeout("ACCT_TIMEOUT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "acct_timeout", "ACCT_TIMEOUT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTacacsplusAuthserviceResource_AuthRetries(t *testing.T) {
	var resourceName = "nios_security_tacacsplus_authservice.test_auth_retries"
	var v security.TacacsplusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTacacsplusAuthserviceAuthRetries("AUTH_RETRIES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_retries", "AUTH_RETRIES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccTacacsplusAuthserviceAuthRetries("AUTH_RETRIES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_retries", "AUTH_RETRIES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTacacsplusAuthserviceResource_AuthTimeout(t *testing.T) {
	var resourceName = "nios_security_tacacsplus_authservice.test_auth_timeout"
	var v security.TacacsplusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTacacsplusAuthserviceAuthTimeout("AUTH_TIMEOUT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_timeout", "AUTH_TIMEOUT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccTacacsplusAuthserviceAuthTimeout("AUTH_TIMEOUT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_timeout", "AUTH_TIMEOUT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTacacsplusAuthserviceResource_Comment(t *testing.T) {
	var resourceName = "nios_security_tacacsplus_authservice.test_comment"
	var v security.TacacsplusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTacacsplusAuthserviceComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccTacacsplusAuthserviceComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTacacsplusAuthserviceResource_Disable(t *testing.T) {
	var resourceName = "nios_security_tacacsplus_authservice.test_disable"
	var v security.TacacsplusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTacacsplusAuthserviceDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccTacacsplusAuthserviceDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTacacsplusAuthserviceResource_Name(t *testing.T) {
	var resourceName = "nios_security_tacacsplus_authservice.test_name"
	var v security.TacacsplusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTacacsplusAuthserviceName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccTacacsplusAuthserviceName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTacacsplusAuthserviceResource_Servers(t *testing.T) {
	var resourceName = "nios_security_tacacsplus_authservice.test_servers"
	var v security.TacacsplusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTacacsplusAuthserviceServers("SERVERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "servers", "SERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccTacacsplusAuthserviceServers("SERVERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "servers", "SERVERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckTacacsplusAuthserviceExists(ctx context.Context, resourceName string, v *security.TacacsplusAuthservice) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SecurityAPI.
			TacacsplusAuthserviceAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForTacacsplusAuthservice).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetTacacsplusAuthserviceResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetTacacsplusAuthserviceResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckTacacsplusAuthserviceDestroy(ctx context.Context, v *security.TacacsplusAuthservice) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SecurityAPI.
			TacacsplusAuthserviceAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForTacacsplusAuthservice).
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

func testAccCheckTacacsplusAuthserviceDisappears(ctx context.Context, v *security.TacacsplusAuthservice) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SecurityAPI.
			TacacsplusAuthserviceAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccTacacsplusAuthserviceBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return `
resource "nios_security_tacacsplus_authservice" "test" {
}
`
}

func testAccTacacsplusAuthserviceRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_security_tacacsplus_authservice" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccTacacsplusAuthserviceAcctRetries(acctRetries string) string {
	return fmt.Sprintf(`
resource "nios_security_tacacsplus_authservice" "test_acct_retries" {
    acct_retries = %q
}
`, acctRetries)
}

func testAccTacacsplusAuthserviceAcctTimeout(acctTimeout string) string {
	return fmt.Sprintf(`
resource "nios_security_tacacsplus_authservice" "test_acct_timeout" {
    acct_timeout = %q
}
`, acctTimeout)
}

func testAccTacacsplusAuthserviceAuthRetries(authRetries string) string {
	return fmt.Sprintf(`
resource "nios_security_tacacsplus_authservice" "test_auth_retries" {
    auth_retries = %q
}
`, authRetries)
}

func testAccTacacsplusAuthserviceAuthTimeout(authTimeout string) string {
	return fmt.Sprintf(`
resource "nios_security_tacacsplus_authservice" "test_auth_timeout" {
    auth_timeout = %q
}
`, authTimeout)
}

func testAccTacacsplusAuthserviceComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_security_tacacsplus_authservice" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccTacacsplusAuthserviceDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_security_tacacsplus_authservice" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccTacacsplusAuthserviceName(name string) string {
	return fmt.Sprintf(`
resource "nios_security_tacacsplus_authservice" "test_name" {
    name = %q
}
`, name)
}

func testAccTacacsplusAuthserviceServers(servers string) string {
	return fmt.Sprintf(`
resource "nios_security_tacacsplus_authservice" "test_servers" {
    servers = %q
}
`, servers)
}
