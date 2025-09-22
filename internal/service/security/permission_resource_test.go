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

var readableAttributesForPermission = "group,object,permission,resource_type,role"

// TODO : OBJECTS TO BE PRESENT IN GRID FOR TESTS
// Create CustomRole1 and CustomRole2 in the NIOS Grid before running these tests.
func TestAccPermissionResource_basic(t *testing.T) {
	var resourceName = "nios_security_permission.test"
	var v security.Permission
	view := acctest.RandomNameWithPrefix("tf-test-view-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPermissionBasicConfig(view, "cloud-api-only", "WRITE", "ZONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPermissionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "group", "cloud-api-only"),
					resource.TestCheckResourceAttr(resourceName, "permission", "WRITE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPermissionResource_disappears(t *testing.T) {
	resourceName := "nios_security_permission.test"
	var v security.Permission
	view := acctest.RandomNameWithPrefix("tf-test-view-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPermissionDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccPermissionBasicConfig(view, "cloud-api-only", "WRITE", "ZONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPermissionExists(context.Background(), resourceName, &v),
					testAccCheckPermissionDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccPermissionResource_Group(t *testing.T) {
	var resourceName = "nios_security_permission.test_group"
	var v security.Permission
	view := acctest.RandomNameWithPrefix("tf-test-view-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPermissionGroup(view, "cloud-api-only", "READ", "ZONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPermissionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "group", "cloud-api-only"),
					resource.TestCheckResourceAttr(resourceName, "permission", "READ"),
					resource.TestCheckResourceAttr(resourceName, "resource_type", "ZONE"),
				),
			},
			// Update and Read
			{
				Config: testAccPermissionGroup(view, "saml-group", "READ", "ZONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPermissionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "group", "saml-group"),
					resource.TestCheckResourceAttr(resourceName, "permission", "READ"),
					resource.TestCheckResourceAttr(resourceName, "resource_type", "ZONE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPermissionResource_Object(t *testing.T) {
	var resourceName = "nios_security_permission.test_object"
	var v security.Permission
	view := acctest.RandomNameWithPrefix("tf-test-view-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPermissionObject(view, "cloud-api-only", "READ", "ZONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPermissionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "object", "nios_dns_view.test_view", "ref"),
					resource.TestCheckResourceAttr(resourceName, "permission", "READ"),
					resource.TestCheckResourceAttr(resourceName, "resource_type", "ZONE"),
				),
			},
			// Update and Read
			{
				Config: testAccPermissionObject(view, "cloud-api-only", "WRITE", "ZONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPermissionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "object", "nios_dns_view.test_view", "ref"),
					resource.TestCheckResourceAttr(resourceName, "permission", "WRITE"),
					resource.TestCheckResourceAttr(resourceName, "resource_type", "ZONE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPermissionResource_Permission(t *testing.T) {
	var resourceName = "nios_security_permission.test_permission"
	var v security.Permission
	view := acctest.RandomNameWithPrefix("tf-test-view-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPermissionPermission(view, "cloud-api-only", "DENY", "ZONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPermissionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "permission", "DENY"),
				),
			},
			// Update and Read
			{
				Config: testAccPermissionPermission(view, "cloud-api-only", "READ", "ZONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPermissionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "permission", "READ"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPermissionResource_ResourceType(t *testing.T) {
	var resourceName = "nios_security_permission.test_resource_type"
	var v security.Permission
	view := acctest.RandomNameWithPrefix("tf-test-view-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{ //view, fqdn, group, object, permission, resourceType
				Config: testAccPermissionResourceType(view, "example.com", "cloud-api-only", "READ", "ZONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPermissionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "resource_type", "ZONE"),
				),
			},
			// Update and Read
			{
				Config: testAccPermissionResourceTypeUpdate(view, "example.com", "cloud-api-only", "READ", "HOST"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPermissionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "resource_type", "HOST"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPermissionResource_Role(t *testing.T) {
	var resourceName = "nios_security_permission.test_role"
	var v security.Permission
	view := acctest.RandomNameWithPrefix("tf-test-view-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPermissionRole(view, "CustomRole1", "WRITE", "ZONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPermissionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "role", "CustomRole1"),
				),
			},
			// Update and Read
			{
				Config: testAccPermissionRole(view, "CustomRole2", "WRITE", "ZONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPermissionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "role", "CustomRole2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckPermissionExists(ctx context.Context, resourceName string, v *security.Permission) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SecurityAPI.
			PermissionAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForPermission).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetPermissionResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetPermissionResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckPermissionDestroy(ctx context.Context, v *security.Permission) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SecurityAPI.
			PermissionAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForPermission).
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

func testAccCheckPermissionDisappears(ctx context.Context, v *security.Permission) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SecurityAPI.
			PermissionAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccPermissionBasicConfig(view, group, permission, resourceType string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_view" {
    name = %q
}

resource "nios_security_permission" "test" {
    group = %q
    permission = %q
    resource_type = %q
	object = nios_dns_view.test_view.ref
}
`, view, group, permission, resourceType)
}

func testAccPermissionGroup(view, group, permission, resourceType string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_view" {
    name = %q
}

resource "nios_security_permission" "test_group" {
    group = %q
    permission = %q
    resource_type = %q
	object = nios_dns_view.test_view.ref
}
`, view, group, permission, resourceType)
}

func testAccPermissionObject(viewName, group, permission, resourceType string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_view" {
    name = %q
}

resource "nios_security_permission" "test_object" {
	group = %q
    object = nios_dns_view.test_view.ref
    permission = %q
    resource_type = %q
}
`, viewName, group, permission, resourceType)
}

func testAccPermissionPermission(view, group, permission, resourceType string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_view" {
    name = %q
}
resource "nios_security_permission" "test_permission" {
    group = %q
    permission = %q
    resource_type = %q
	object = nios_dns_view.test_view.ref
}
`, view, group, permission, resourceType)
}

func testAccPermissionResourceType(view, fqdn, group, permission, resourceType string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_view" {
    name = %q
}

resource "nios_dns_zone_auth" "example_zone" {
  fqdn = %q
  view = nios_dns_view.test_view.name
}

resource "nios_security_permission" "test_resource_type" {
    group = %q
    object = nios_dns_view.test_view.ref
    permission = %q
    resource_type = %q
}
`, view, fqdn, group, permission, resourceType)
}

func testAccPermissionResourceTypeUpdate(view, fqdn, group, permission, resourceType string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_view" {
    name = %q
}

resource "nios_dns_zone_auth" "example_zone" {
  fqdn = %q
  view = nios_dns_view.test_view.name
}

resource "nios_security_permission" "test_resource_type" {
    group = %q
    object = nios_dns_zone_auth.example_zone.ref
    permission = %q
    resource_type = %q
}
`, view, fqdn, group, permission, resourceType)
}

func testAccPermissionRole(view, role, permission, resourceType string) string {
	return fmt.Sprintf(`

resource "nios_dns_view" "test_view" {
    name = %q
}
resource "nios_security_admin_role" "test_admin_role" {
  	name = %q
}

resource "nios_security_permission" "test_role" {
    role = nios_security_admin_role.test_admin_role.name
    permission = %q
    resource_type = %q
	object = nios_dns_view.test_view.ref
}
`, view, role, permission, resourceType)
}
