package acl_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/acl"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForNamedacl = "access_list,comment,exploded_access_list,extattrs,name"

func TestAccNamedaclResource_basic(t *testing.T) {
	var resourceName = "nios_acl_namedacl.test"
	var v acl.Namedacl
	name := acctest.RandomNameWithPrefix("namedacl")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNamedaclBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedaclExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNamedaclResource_disappears(t *testing.T) {
	resourceName := "nios_acl_namedacl.test"
	var v acl.Namedacl
	name := acctest.RandomNameWithPrefix("namedacl")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNamedaclDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNamedaclBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedaclExists(context.Background(), resourceName, &v),
					testAccCheckNamedaclDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNamedaclResource_AccessList(t *testing.T) {
	var resourceName = "nios_acl_namedacl.test_access_list"
	var v acl.Namedacl
	name := acctest.RandomNameWithPrefix("namedacl")
	acl1 := []map[string]any{
		{
			"struct":     "addressac",
			"address":    "10.0.0.5",
			"permission": "DENY",
		},
	}

	acl2 := []map[string]any{
		{
			"struct":        "tsigac",
			"tsig_key":      "X4oRe92t54I+T98NdQpV2w==",
			"tsig_key_name": "example-tsig-key",
			"tsig_key_alg":  "HMAC-SHA256",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNamedaclAccessList(name, acl1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedaclExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "access_list.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "access_list.0.address", "10.0.0.5"),
					resource.TestCheckResourceAttr(resourceName, "access_list.0.permission", "DENY"),
				),
			},
			// Update and Read
			{
				Config: testAccNamedaclAccessList(name, acl2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedaclExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "access_list.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "access_list.0.tsig_key", "X4oRe92t54I+T98NdQpV2w=="),
					resource.TestCheckResourceAttr(resourceName, "access_list.0.tsig_key_name", "example-tsig-key"),
					resource.TestCheckResourceAttr(resourceName, "access_list.0.tsig_key_alg", "HMAC-SHA256"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNamedaclResource_Comment(t *testing.T) {
	var resourceName = "nios_acl_namedacl.test_comment"
	var v acl.Namedacl
	name := acctest.RandomNameWithPrefix("namedacl")
	comment1 := "This is a new named acl"
	comment2 := "This is a updated named acl"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNamedaclComment(name, comment1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedaclExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment1),
				),
			},
			// Update and Read
			{
				Config: testAccNamedaclComment(name, comment2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedaclExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNamedaclResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_acl_namedacl.test_extattrs"
	var v acl.Namedacl
	name := acctest.RandomNameWithPrefix("namedacl")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNamedaclExtAttrs(name, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedaclExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccNamedaclExtAttrs(name, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedaclExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNamedaclResource_Name(t *testing.T) {
	var resourceName = "nios_acl_namedacl.test_name"
	var v acl.Namedacl
	name1 := acctest.RandomNameWithPrefix("namedacl")
	name2 := acctest.RandomNameWithPrefix("namedacl")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNamedaclName(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedaclExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccNamedaclName(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNamedaclExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckNamedaclExists(ctx context.Context, resourceName string, v *acl.Namedacl) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.ACLAPI.
			NamedaclAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForNamedacl).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetNamedaclResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetNamedaclResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckNamedaclDestroy(ctx context.Context, v *acl.Namedacl) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.ACLAPI.
			NamedaclAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForNamedacl).
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

func testAccCheckNamedaclDisappears(ctx context.Context, v *acl.Namedacl) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.ACLAPI.
			NamedaclAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccNamedaclBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_acl_namedacl" "test" {
	name = %q
}
`, name)
}

func testAccNamedaclAccessList(name string, acl []map[string]any) string {
	aclHCL := utils.ConvertSliceOfMapsToHCL(acl)
	return fmt.Sprintf(`
resource "nios_acl_namedacl" "test_access_list" {
	name = %q
    access_list = %s
}
`, name, aclHCL)
}

func testAccNamedaclComment(name, comment string) string {
	return fmt.Sprintf(`
resource "nios_acl_namedacl" "test_comment" {
	name = %q
    comment = %q
}
`, name, comment)
}

func testAccNamedaclExtAttrs(name string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_acl_namedacl" "test_extattrs" {
	name = %q
    extattrs = %s
}
`, name, extattrsStr)
}

func testAccNamedaclName(name string) string {
	return fmt.Sprintf(`
resource "nios_acl_namedacl" "test_name" {
    name = %q
}
`, name)
}
