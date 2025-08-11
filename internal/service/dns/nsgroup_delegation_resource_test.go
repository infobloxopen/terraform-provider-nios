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

var readableAttributesForNsgroupDelegation = "comment,delegate_to,extattrs,name"

func TestAccNsgroupDelegationResource_basic(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_delegation.test"
	var v dns.NsgroupDelegation

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupDelegationBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupDelegationResource_disappears(t *testing.T) {
	resourceName := "nios_dns_nsgroup_delegation.test"
	var v dns.NsgroupDelegation

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupDelegationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupDelegationBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					testAccCheckNsgroupDelegationDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNsgroupDelegationResource_Ref(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_delegation.test_ref"
	var v dns.NsgroupDelegation

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupDelegationRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupDelegationRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupDelegationResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_delegation.test_comment"
	var v dns.NsgroupDelegation

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupDelegationComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupDelegationComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupDelegationResource_DelegateTo(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_delegation.test_delegate_to"
	var v dns.NsgroupDelegation

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupDelegationDelegateTo("DELEGATE_TO_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegate_to", "DELEGATE_TO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupDelegationDelegateTo("DELEGATE_TO_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegate_to", "DELEGATE_TO_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupDelegationResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_delegation.test_extattrs"
	var v dns.NsgroupDelegation

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupDelegationExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupDelegationExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupDelegationResource_Name(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_delegation.test_name"
	var v dns.NsgroupDelegation

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupDelegationName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupDelegationName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckNsgroupDelegationExists(ctx context.Context, resourceName string, v *dns.NsgroupDelegation) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			NsgroupDelegationAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForNsgroupDelegation).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetNsgroupDelegationResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetNsgroupDelegationResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckNsgroupDelegationDestroy(ctx context.Context, v *dns.NsgroupDelegation) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			NsgroupDelegationAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForNsgroupDelegation).
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

func testAccCheckNsgroupDelegationDisappears(ctx context.Context, v *dns.NsgroupDelegation) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			NsgroupDelegationAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccNsgroupDelegationBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_delegation" "test" {
}
`)
}

func testAccNsgroupDelegationRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_delegation" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccNsgroupDelegationComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_delegation" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccNsgroupDelegationDelegateTo(delegateTo string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_delegation" "test_delegate_to" {
    delegate_to = %q
}
`, delegateTo)
}

func testAccNsgroupDelegationExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_delegation" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccNsgroupDelegationName(name string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_delegation" "test_name" {
    name = %q
}
`, name)
}
