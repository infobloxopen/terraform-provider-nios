package dns_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForNsgroupStubmember = "comment,extattrs,name,stub_members"

func TestAccNsgroupStubmemberResource_basic(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_stubmember.test"
	var v dns.NsgroupStubmember
	name := acctest.RandomNameWithPrefix("test-nsgroup-stubmember")
	stubMember := []map[string]any{
		{
			"name": "member.com",
		},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupStubmemberBasicConfig(name,stubMember),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupStubmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "stub_members.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "stub_members.0.name", "member.com"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupStubmemberResource_disappears(t *testing.T) {
	resourceName := "nios_dns_nsgroup_stubmember.test"
	var v dns.NsgroupStubmember
	name := acctest.RandomNameWithPrefix("test-nsgroup-stubmember")
	stubMember := []map[string]any{
		{
			"name": "member.com",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupStubmemberDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupStubmemberBasicConfig(name,stubMember),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupStubmemberExists(context.Background(), resourceName, &v),
					testAccCheckNsgroupStubmemberDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNsgroupStubmemberResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_stubmember.test_comment"
	var v dns.NsgroupStubmember
	name := acctest.RandomNameWithPrefix("test-nsgroup-stubmember")
	stubMember := []map[string]any{
		{
			"name": "member.com",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupStubmemberComment(name , stubMember , "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupStubmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupStubmemberComment(name, stubMember, "This comment is updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupStubmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This comment is updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupStubmemberResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_stubmember.test_extattrs"
	var v dns.NsgroupStubmember
	name := acctest.RandomNameWithPrefix("test-nsgroup-stubmember")
	stubMember := []map[string]any{
		{
			"name": "member.com",
		},
	}
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupStubmemberExtAttrs(name , stubMember ,  map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupStubmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupStubmemberExtAttrs(name , stubMember , map[string]string{
					"Site":    extAttrValue2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupStubmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupStubmemberResource_Name(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_stubmember.test_name"
	var v dns.NsgroupStubmember
	name := acctest.RandomNameWithPrefix("test-nsgroup-stubmember")
	nameUpdate := acctest.RandomNameWithPrefix("test-nsgroup-stubmember")
	stubMember := []map[string]any{
		{
			"name": "member.com",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupStubmemberName(name , stubMember),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupStubmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupStubmemberName(nameUpdate,stubMember),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupStubmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupStubmemberResource_StubMembers(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_stubmember.test_stub_members"
	var v dns.NsgroupStubmember
	name := acctest.RandomNameWithPrefix("test-nsgroup-stubmember")
	stubMember := []map[string]any{
		{
			"name": "member.com",
		},
	}
	stubMemberUpdate := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupStubmemberStubMembers(name , stubMember ),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupStubmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "stub_members.0.name", "member.com"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupStubmemberStubMembers(name , stubMemberUpdate ),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupStubmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "stub_members.0.name", "infoblox.localdomain"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckNsgroupStubmemberExists(ctx context.Context, resourceName string, v *dns.NsgroupStubmember) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			NsgroupStubmemberAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForNsgroupStubmember).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetNsgroupStubmemberResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetNsgroupStubmemberResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckNsgroupStubmemberDestroy(ctx context.Context, v *dns.NsgroupStubmember) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			NsgroupStubmemberAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForNsgroupStubmember).
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

func testAccCheckNsgroupStubmemberDisappears(ctx context.Context, v *dns.NsgroupStubmember) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			NsgroupStubmemberAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccNsgroupStubmemberBasicConfig(name string , stubMember []map[string]any) string {
	stubMemberStr := utils.ConvertSliceOfMapsToHCL(stubMember)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_stubmember" "test" {
    name = %q
    stub_members = %s
}
`, name, stubMemberStr)
}

func testAccNsgroupStubmemberComment(name string , stubMember []map[string]any, comment string) string {
	stubMemberStr := utils.ConvertSliceOfMapsToHCL(stubMember)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_stubmember" "test_comment" {
    name = %q
    stub_members = %s
    comment = %q
}
`, name, stubMemberStr, comment)
}

func testAccNsgroupStubmemberExtAttrs(name string , stubMember []map[string]any,extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf("    %q = %q\n", k, v)
	}
	extattrsStr += "  }"
	stubMemberStr := utils.ConvertSliceOfMapsToHCL(stubMember)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_stubmember" "test_extattrs" {
	name = %q
    extattrs = %s
    stub_members = %s
}
`, name, extattrsStr, stubMemberStr)
}

func testAccNsgroupStubmemberName(name string,stubMember []map[string]any) string {
	stubMemberStr := utils.ConvertSliceOfMapsToHCL(stubMember)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_stubmember" "test_name" {
    name = %q
    stub_members = %s
}
`, name, stubMemberStr)
}

func testAccNsgroupStubmemberStubMembers(name string,stubMember []map[string]any) string {
	stubMemberStr := utils.ConvertSliceOfMapsToHCL(stubMember)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_stubmember" "test_stub_members" {
    name = %q
    stub_members = %s
}
`, name, stubMemberStr)
}
