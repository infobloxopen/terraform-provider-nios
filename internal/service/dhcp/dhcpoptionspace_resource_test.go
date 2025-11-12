package dhcp_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForDhcpoptionspace = "comment,name,option_definitions,space_type"

func TestAccDhcpoptionspaceResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpoptionspace.test"
	var v dhcp.Dhcpoptionspace
	name := acctest.RandomNameWithPrefix("dhcp-option-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpoptionspaceBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptionspaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpoptionspaceResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_dhcpoptionspace.test"
	var v dhcp.Dhcpoptionspace
	name := acctest.RandomNameWithPrefix("dhcp-option-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDhcpoptionspaceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDhcpoptionspaceBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptionspaceExists(context.Background(), resourceName, &v),
					testAccCheckDhcpoptionspaceDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDhcpoptionspaceResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpoptionspace.test_comment"
	var v dhcp.Dhcpoptionspace
	name := acctest.RandomNameWithPrefix("dhcp-option-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpoptionspaceComment(name, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptionspaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpoptionspaceComment(name, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptionspaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpoptionspaceResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpoptionspace.test_name"
	var v dhcp.Dhcpoptionspace
	name := acctest.RandomNameWithPrefix("dhcp-option-space")
	updatedName := acctest.RandomNameWithPrefix("dhcp-option-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpoptionspaceName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptionspaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpoptionspaceName(updatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptionspaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpoptionspaceResource_OptionDefinitions(t *testing.T) {
	var resourceName = "nios_dhcp_dhcpoptionspace.test_option_definitions"
	var v dhcp.Dhcpoptionspace
	name := acctest.RandomNameWithPrefix("dhcp-option-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpoptionspaceOptionDefinitions(name, "OPTION_DEFINITIONS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptionspaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_definitions", "OPTION_DEFINITIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpoptionspaceOptionDefinitions(name, "OPTION_DEFINITIONS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptionspaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_definitions", "OPTION_DEFINITIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDhcpoptionspaceExists(ctx context.Context, resourceName string, v *dhcp.Dhcpoptionspace) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			DhcpoptionspaceAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDhcpoptionspace).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDhcpoptionspaceResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDhcpoptionspaceResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDhcpoptionspaceDestroy(ctx context.Context, v *dhcp.Dhcpoptionspace) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			DhcpoptionspaceAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDhcpoptionspace).
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

func testAccCheckDhcpoptionspaceDisappears(ctx context.Context, v *dhcp.Dhcpoptionspace) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			DhcpoptionspaceAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDhcpoptionspaceBasicConfig(name string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpoptionspace" "test" {
    name = %q
}
`, name)
}

func testAccDhcpoptionspaceComment(name string, comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpoptionspace" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccDhcpoptionspaceName(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpoptionspace" "test_name" {
    name = %q
}
`, name)
}

func testAccDhcpoptionspaceOptionDefinitions(name string, optionDefinitions string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpoptionspace" "test_option_definitions" {
    name = %q
    option_definitions = %q
}
`, name, optionDefinitions)
}
