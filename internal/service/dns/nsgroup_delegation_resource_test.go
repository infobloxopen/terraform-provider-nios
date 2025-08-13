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
	name := acctest.RandomName()
	delegateTo := []map[string]interface{}{
		{
			"name":"delegate_to_ns_group",
			"address":"2.3.4.5",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupDelegationBasicConfig(name , delegateTo),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "delegate_to.0.address", "2.3.4.5"), 
					resource.TestCheckResourceAttr(resourceName, "delegate_to.0.name", "delegate_to_ns_group"),
					// Test fields with default value	
					resource.TestCheckResourceAttr(resourceName, "comment", ""),

				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupDelegationResource_disappears(t *testing.T) {
	resourceName := "nios_dns_nsgroup_delegation.test"
	var v dns.NsgroupDelegation
	name := acctest.RandomName()
	delegateTo := []map[string]interface{}{
		{
			"name":"delegate_to_ns_group",
			"address":"2.3.4.5",
		},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupDelegationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupDelegationBasicConfig(name, delegateTo),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					testAccCheckNsgroupDelegationDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNsgroupDelegationResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_delegation.test_comment"
	var v dns.NsgroupDelegation
	name := acctest.RandomName()
	delegateTo := []map[string]interface{}{
		{
			"name":"delegate_to_ns_group",
			"address":"2.3.4.5",
		},
	}
	comment := "comment ns group"
	commentUpdate := "updated comment ns group"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupDelegationComment(name, comment , delegateTo),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupDelegationComment(name, commentUpdate, delegateTo),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", commentUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupDelegationResource_DelegateTo(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_delegation.test_delegate_to"
	var v dns.NsgroupDelegation
	name := acctest.RandomName()
	delegateTo := []map[string]interface{}{
		{
			"name":"delegate_to_ns_group",
			"address":"2.3.4.5",
		},
	}
	delegateToUpdate := []map[string]interface{}{
		{
			"name":"delegate_to_ns_group_update",
			"address":"2.3.4.6",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupDelegationDelegateTo(name , delegateTo),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegate_to.0.address", "2.3.4.5"),
					resource.TestCheckResourceAttr(resourceName, "delegate_to.0.name", "delegate_to_ns_group"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupDelegationDelegateTo(name, delegateToUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegate_to.0.address", "2.3.4.6"),
					resource.TestCheckResourceAttr(resourceName, "delegate_to.0.name", "delegate_to_ns_group_update"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupDelegationResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_delegation.test_extattrs"
	var v dns.NsgroupDelegation
	name := acctest.RandomName()
	delegateTo := []map[string]interface{}{
		{
			"name":"delegate_to_ns_group",
			"address":"2.3.4.5",
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
				Config: testAccNsgroupDelegationExtAttrs(name , delegateTo , map[string]any{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupDelegationExtAttrs(name , delegateTo , map[string]any{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupDelegationResource_Name(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_delegation.test_name"
	var v dns.NsgroupDelegation
	name := acctest.RandomName()
	delegateTo := []map[string]interface{}{
		{
			"name":"delegate_to_ns_group",
			"address":"2.3.4.5",
		},
	}
	nameUpdate := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupDelegationName(name , delegateTo),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupDelegationName(nameUpdate, delegateTo),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
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

func testAccNsgroupDelegationBasicConfig(name string, delegateTo []map[string]any) string {
	delegateToStr := utils.ConvertSliceOfMapsToHCL(delegateTo)
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_delegation" "test" {
    name = %q
    delegate_to = %s
}
`, name, delegateToStr)
}

func testAccNsgroupDelegationComment(name , comment string, delegateTo []map[string]any) string {
	delegateToStr := utils.ConvertSliceOfMapsToHCL(delegateTo)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_delegation" "test_comment" {
    name = %q
	delegate_to = %s
	comment = %q
}
`, name, delegateToStr, comment)
}

func testAccNsgroupDelegationDelegateTo(name string, delegateTo []map[string]any) string {
	delegateToStr := utils.ConvertSliceOfMapsToHCL(delegateTo)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_delegation" "test_delegate_to" {
    name = %q
    delegate_to = %s
}
`, name, delegateToStr)
}

func testAccNsgroupDelegationExtAttrs(name string , delegateTo []map[string]any , extAttrs map[string]any) string {
	delegateToStr := utils.ConvertSliceOfMapsToHCL(delegateTo)
	extAttrsStr := utils.ConvertMapToHCL(extAttrs)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_delegation" "test_extattrs" {
    name = %q
    delegate_to = %s
    extattrs = %s
}
`, name, delegateToStr, extAttrsStr)
}

func testAccNsgroupDelegationName(name string , delegateTo []map[string]any) string {
	delegateToStr := utils.ConvertSliceOfMapsToHCL(delegateTo)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_delegation" "test_name" {
    name = %q
    delegate_to = %s
}
`, name, delegateToStr)
}
