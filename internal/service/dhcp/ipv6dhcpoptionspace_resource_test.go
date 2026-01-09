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

var readableAttributesForIpv6dhcpoptionspace = "comment,enterprise_number,name,option_definitions"

func TestAccIpv6dhcpoptionspaceResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6optionspace.test"
	var v dhcp.Ipv6dhcpoptionspace
	name := acctest.RandomNameWithPrefix("option-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6dhcpoptionspaceBasicConfig("5896", name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6dhcpoptionspaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enterprise_number", "5896"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6dhcpoptionspaceResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_ipv6optionspace.test"
	var v dhcp.Ipv6dhcpoptionspace
	name := acctest.RandomNameWithPrefix("option-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6dhcpoptionspaceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6dhcpoptionspaceBasicConfig("5896", name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6dhcpoptionspaceExists(context.Background(), resourceName, &v),
					testAccCheckIpv6dhcpoptionspaceDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccIpv6dhcpoptionspaceResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6optionspace.test_comment"
	var v dhcp.Ipv6dhcpoptionspace
	name := acctest.RandomNameWithPrefix("option-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6dhcpoptionspaceComment("5896", name, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6dhcpoptionspaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6dhcpoptionspaceComment("5896", name, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6dhcpoptionspaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6dhcpoptionspaceResource_EnterpriseNumber(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6optionspace.test_enterprise_number"
	var v dhcp.Ipv6dhcpoptionspace
	name := acctest.RandomNameWithPrefix("option-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6dhcpoptionspaceEnterpriseNumber("5896", name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6dhcpoptionspaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enterprise_number", "5896"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6dhcpoptionspaceEnterpriseNumber("5123", name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6dhcpoptionspaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enterprise_number", "5123"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6dhcpoptionspaceResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6optionspace.test_name"
	var v dhcp.Ipv6dhcpoptionspace
	name := acctest.RandomNameWithPrefix("option-space")
	updatedName := acctest.RandomNameWithPrefix("option-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6dhcpoptionspaceName("5896", name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6dhcpoptionspaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6dhcpoptionspaceName("5896", updatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6dhcpoptionspaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckIpv6dhcpoptionspaceExists(ctx context.Context, resourceName string, v *dhcp.Ipv6dhcpoptionspace) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			Ipv6dhcpoptionspaceAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForIpv6dhcpoptionspace).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetIpv6dhcpoptionspaceResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetIpv6dhcpoptionspaceResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckIpv6dhcpoptionspaceDestroy(ctx context.Context, v *dhcp.Ipv6dhcpoptionspace) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			Ipv6dhcpoptionspaceAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForIpv6dhcpoptionspace).
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

func testAccCheckIpv6dhcpoptionspaceDisappears(ctx context.Context, v *dhcp.Ipv6dhcpoptionspace) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			Ipv6dhcpoptionspaceAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccIpv6dhcpoptionspaceBasicConfig(enterpriseNumber, name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6optionspace" "test" {
    enterprise_number = %q
    name = %q
}
`, enterpriseNumber, name)
}

func testAccIpv6dhcpoptionspaceComment(enterpriseNumber string, name string, comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6optionspace" "test_comment" {
    enterprise_number = %q
    name = %q
    comment = %q
}
`, enterpriseNumber, name, comment)
}

func testAccIpv6dhcpoptionspaceEnterpriseNumber(enterpriseNumber string, name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6optionspace" "test_enterprise_number" {
    enterprise_number = %q
    name = %q
}
`, enterpriseNumber, name)
}

func testAccIpv6dhcpoptionspaceName(enterpriseNumber string, name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6optionspace" "test_name" {
    enterprise_number = %q
    name = %q
}
`, enterpriseNumber, name)
}
