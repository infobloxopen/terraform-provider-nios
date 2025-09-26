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

var readableAttributesForFtpuser = "extattrs,home_dir,permission,username"

func TestAccFtpuserResource_basic(t *testing.T) {
	var resourceName = "nios_security_ftpuser.test"
	var v security.Ftpuser
	username := acctest.RandomNameWithPrefix("ftf-test-user-")
	password := acctest.RandomAlphaNumeric(12)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFtpuserBasicConfig(username, password),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					// Test fields with required value
					resource.TestCheckResourceAttr(resourceName, "username", username),
					resource.TestCheckResourceAttr(resourceName, "password", password),
					resource.TestCheckResourceAttr(resourceName, "permission", "RO"),
					resource.TestCheckResourceAttr(resourceName, "create_home_dir", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFtpuserResource_disappears(t *testing.T) {
	resourceName := "nios_security_ftpuser.test"
	var v security.Ftpuser
	username := acctest.RandomNameWithPrefix("ftf-test-user-")
	password := acctest.RandomAlphaNumeric(12)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFtpuserDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFtpuserBasicConfig(username, password),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					testAccCheckFtpuserDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccFtpuserResource_CreateHomeDir(t *testing.T) {
	var resourceName = "nios_security_ftpuser.test_create_home_dir"
	var v security.Ftpuser
	username := acctest.RandomNameWithPrefix("ftf-test-user-")
	password := acctest.RandomAlphaNumeric(12)
	homeDir := "/ftpusers"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFtpuserCreateHomeDir(username, password, "false", homeDir),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "create_home_dir", "false"),
					resource.TestCheckResourceAttr(resourceName, "home_dir", homeDir),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFtpuserResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_security_ftpuser.test_extattrs"
	var v security.Ftpuser
	username := acctest.RandomNameWithPrefix("ftf-test-user-")
	password := acctest.RandomAlphaNumeric(12)
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFtpuserExtAttrs(username, password, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccFtpuserExtAttrs(username, password, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFtpuserResource_HomeDir(t *testing.T) {
	var resourceName = "nios_security_ftpuser.test_home_dir"
	var v security.Ftpuser
	username := acctest.RandomNameWithPrefix("ftf-test-user-")
	password := acctest.RandomAlphaNumeric(12)
	homeDir := "/ftpusers"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFtpuserHomeDir(username, password, homeDir, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "home_dir", homeDir),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFtpuserResource_Permission(t *testing.T) {
	var resourceName = "nios_security_ftpuser.test_permission"
	var v security.Ftpuser
	username := acctest.RandomNameWithPrefix("ftf-test-user-")
	password := acctest.RandomAlphaNumeric(12)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFtpuserPermission(username, password, "RW"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "permission", "RW"),
				),
			},
			// Update and Read
			{
				Config: testAccFtpuserPermission(username, password, "RO"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "permission", "RO"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFtpuserResource_Username(t *testing.T) {
	var resourceName = "nios_security_ftpuser.test_username"
	var v security.Ftpuser
	username := acctest.RandomNameWithPrefix("ftf-test-user-")
	password := acctest.RandomAlphaNumeric(12)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFtpuserUsername(username, password),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "username", username),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckFtpuserExists(ctx context.Context, resourceName string, v *security.Ftpuser) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SecurityAPI.
			FtpuserAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForFtpuser).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetFtpuserResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetFtpuserResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckFtpuserDestroy(ctx context.Context, v *security.Ftpuser) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SecurityAPI.
			FtpuserAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForFtpuser).
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

func testAccCheckFtpuserDisappears(ctx context.Context, v *security.Ftpuser) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SecurityAPI.
			FtpuserAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccFtpuserBasicConfig(username, password string) string {

	return fmt.Sprintf(`
resource "nios_security_ftpuser" "test" {
    username = %q
    password = %q
}
`, username, password)
}

func testAccFtpuserCreateHomeDir(username, password, createHomeDir, homeDir string) string {
	return fmt.Sprintf(`
resource "nios_security_ftpuser" "test_create_home_dir" {
    username      = %q
    password      = %q
    create_home_dir = %q
    home_dir     = %q
}
`, username, password, createHomeDir, homeDir)
}

func testAccFtpuserExtAttrs(username, password string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_security_ftpuser" "test_extattrs" {
	username      = %q
    password      = %q
    extattrs      = %s
}
	`, username, password, extattrsStr)
}

func testAccFtpuserHomeDir(username, password, homeDir string, createHomeDir bool) string {
	return fmt.Sprintf(`
resource "nios_security_ftpuser" "test_home_dir" {
    username      = %q
    password      = %q
    home_dir      = %q
    create_home_dir = %t
}
`, username, password, homeDir, createHomeDir)
}

func testAccFtpuserPermission(username, password, permission string) string {
	return fmt.Sprintf(`
resource "nios_security_ftpuser" "test_permission" {
    username = %q
    password = %q
    permission = %q
}
`, username, password, permission)
}

func testAccFtpuserUsername(username, password string) string {
	return fmt.Sprintf(`
resource "nios_security_ftpuser" "test_username" {
    username = %q
    password = %q
}
`, username, password)
}
