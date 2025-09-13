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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFtpuserBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFtpuserResource_disappears(t *testing.T) {
	resourceName := "nios_security_ftpuser.test"
	var v security.Ftpuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFtpuserDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFtpuserBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					testAccCheckFtpuserDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccFtpuserResource_Ref(t *testing.T) {
	var resourceName = "nios_security_ftpuser.test_ref"
	var v security.Ftpuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFtpuserRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFtpuserRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFtpuserResource_CreateHomeDir(t *testing.T) {
	var resourceName = "nios_security_ftpuser.test_create_home_dir"
	var v security.Ftpuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFtpuserCreateHomeDir("CREATE_HOME_DIR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "create_home_dir", "CREATE_HOME_DIR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFtpuserCreateHomeDir("CREATE_HOME_DIR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "create_home_dir", "CREATE_HOME_DIR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFtpuserResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_security_ftpuser.test_extattrs"
	var v security.Ftpuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFtpuserExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFtpuserExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFtpuserResource_HomeDir(t *testing.T) {
	var resourceName = "nios_security_ftpuser.test_home_dir"
	var v security.Ftpuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFtpuserHomeDir("HOME_DIR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "home_dir", "HOME_DIR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFtpuserHomeDir("HOME_DIR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "home_dir", "HOME_DIR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFtpuserResource_Password(t *testing.T) {
	var resourceName = "nios_security_ftpuser.test_password"
	var v security.Ftpuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFtpuserPassword("PASSWORD_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "password", "PASSWORD_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFtpuserPassword("PASSWORD_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "password", "PASSWORD_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFtpuserResource_Permission(t *testing.T) {
	var resourceName = "nios_security_ftpuser.test_permission"
	var v security.Ftpuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFtpuserPermission("PERMISSION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "permission", "PERMISSION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFtpuserPermission("PERMISSION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "permission", "PERMISSION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFtpuserResource_Username(t *testing.T) {
	var resourceName = "nios_security_ftpuser.test_username"
	var v security.Ftpuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFtpuserUsername("USERNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "username", "USERNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFtpuserUsername("USERNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "username", "USERNAME_UPDATE_REPLACE_ME"),
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

func testAccFtpuserBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_security_ftpuser" "test" {
}
`)
}

func testAccFtpuserRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_security_ftpuser" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccFtpuserCreateHomeDir(createHomeDir string) string {
	return fmt.Sprintf(`
resource "nios_security_ftpuser" "test_create_home_dir" {
    create_home_dir = %q
}
`, createHomeDir)
}

func testAccFtpuserExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_security_ftpuser" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccFtpuserHomeDir(homeDir string) string {
	return fmt.Sprintf(`
resource "nios_security_ftpuser" "test_home_dir" {
    home_dir = %q
}
`, homeDir)
}

func testAccFtpuserPassword(password string) string {
	return fmt.Sprintf(`
resource "nios_security_ftpuser" "test_password" {
    password = %q
}
`, password)
}

func testAccFtpuserPermission(permission string) string {
	return fmt.Sprintf(`
resource "nios_security_ftpuser" "test_permission" {
    permission = %q
}
`, permission)
}

func testAccFtpuserUsername(username string) string {
	return fmt.Sprintf(`
resource "nios_security_ftpuser" "test_username" {
    username = %q
}
`, username)
}
