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

var readableAttributesForAdminrole = "comment,disable,extattrs,name"

func TestAccAdminroleResource_basic(t *testing.T) {
	var resourceName = "nios_security_admin_role.test"
	var v security.Adminrole
	name := acctest.RandomNameWithPrefix("admin-role")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminroleBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminroleExists(context.Background(), resourceName, &v),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminroleResource_disappears(t *testing.T) {
	resourceName := "nios_security_admin_role.test"
	var v security.Adminrole

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAdminroleDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAdminroleBasicConfig(acctest.RandomNameWithPrefix("admin-role")),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminroleExists(context.Background(), resourceName, &v),
					testAccCheckAdminroleDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAdminroleResource_Comment(t *testing.T) {
	var resourceName = "nios_security_admin_role.test_comment"
	var v security.Adminrole
	name := acctest.RandomNameWithPrefix("admin-role")
	comment1 := "test comment"
	comment2 := "test comment updated"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminroleComment(name, comment1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminroleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment1),
				),
			},
			// Update and Read
			{
				Config: testAccAdminroleComment(name, comment2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminroleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminroleResource_Disable(t *testing.T) {
	var resourceName = "nios_security_admin_role.test_disable"
	var v security.Adminrole
	name := acctest.RandomNameWithPrefix("admin-role")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminroleDisable(name, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminroleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminroleDisable(name, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminroleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminroleResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_security_admin_role.test_extattrs"
	var v security.Adminrole
	name := acctest.RandomNameWithPrefix("admin-role")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminroleExtAttrs(name, map[string]string{"Site": extAttrValue1}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminroleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccAdminroleExtAttrs(name, map[string]string{"Site": extAttrValue2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminroleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminroleResource_Name(t *testing.T) {
	var resourceName = "nios_security_admin_role.test_name"
	var v security.Adminrole
	name1 := acctest.RandomNameWithPrefix("admin-role")
	name2 := acctest.RandomNameWithPrefix("admin-role")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminroleName(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminroleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccAdminroleName(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminroleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckAdminroleExists(ctx context.Context, resourceName string, v *security.Adminrole) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SecurityAPI.
			AdminroleAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForAdminrole).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetAdminroleResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetAdminroleResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckAdminroleDestroy(ctx context.Context, v *security.Adminrole) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SecurityAPI.
			AdminroleAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForAdminrole).
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

func testAccCheckAdminroleDisappears(ctx context.Context, v *security.Adminrole) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SecurityAPI.
			AdminroleAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccAdminroleBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_role" "test" {
	name = %q
}
`, name)
}

func testAccAdminroleComment(name, comment string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_role" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccAdminroleDisable(name string, disable bool) string {
	return fmt.Sprintf(`
resource "nios_security_admin_role" "test_disable" {
    name = %q
    disable = %t
}
`, name, disable)
}

func testAccAdminroleExtAttrs(name string, extAttrs map[string]string) string {
	extattrsStr := "{"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`%s = %q`, k, v)
	}
	extattrsStr += "}"
	return fmt.Sprintf(`
resource "nios_security_admin_role" "test_extattrs" {
    name = %q
    extattrs = %s
}
`, name, extattrsStr)
}

func testAccAdminroleName(name string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_role" "test_name" {
    name = %q
}
`, name)
}
