package dhcp_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForDhcpoptiondefinition = "code,name,space,type"

func TestAccDhcpoptiondefinitionResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_optiondefinition.test"
	var v dhcp.Dhcpoptiondefinition
	name := acctest.RandomNameWithPrefix("dhcp-option-definition")
	space := acctest.RandomNameWithPrefix("dhcp-option-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpoptiondefinitionBasicConfig("10", name, "string", space),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "code", "10"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "type", "string"),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpoptiondefinitionResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_optiondefinition.test"
	var v dhcp.Dhcpoptiondefinition
	name := acctest.RandomNameWithPrefix("dhcp-option-definition")
	space := acctest.RandomNameWithPrefix("dhcp-option-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDhcpoptiondefinitionDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDhcpoptiondefinitionBasicConfig("10", name, "string", space),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					testAccCheckDhcpoptiondefinitionDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDhcpoptiondefinitionResource_Code(t *testing.T) {
	var resourceName = "nios_dhcp_optiondefinition.test_code"
	var v dhcp.Dhcpoptiondefinition
	name := acctest.RandomNameWithPrefix("dhcp-option-definition")
	space := acctest.RandomNameWithPrefix("dhcp-option-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpoptiondefinitionCode("10", name, "string", space),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "code", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpoptiondefinitionCode("20", name, "string", space),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "code", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpoptiondefinitionResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_optiondefinition.test_name"
	var v dhcp.Dhcpoptiondefinition
	name := acctest.RandomNameWithPrefix("dhcp-option-definition")
	updatedName := acctest.RandomNameWithPrefix("dhcp-option-definition")
	space := acctest.RandomNameWithPrefix("dhcp-option-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpoptiondefinitionName("10", name, "string", space),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpoptiondefinitionName("10", updatedName, "string", space),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpoptiondefinitionResource_Space(t *testing.T) {
	var resourceName = "nios_dhcp_optiondefinition.test_space"
	var v dhcp.Dhcpoptiondefinition
	name := acctest.RandomNameWithPrefix("dhcp-option-definition")
	optionSpace1 := acctest.RandomNameWithPrefix("option-space")
	optionSpace2 := acctest.RandomNameWithPrefix("option-space")
	optionSpace1ResourceName := "nios_dhcp_optionspace.test1"
	optionSpace2ResourceName := "nios_dhcp_optionspace.test2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpoptiondefinitionSpace("10", name, "string", optionSpace1, optionSpace2, optionSpace1ResourceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "space", optionSpace1),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpoptiondefinitionSpace("10", name, "string", optionSpace1, optionSpace2, optionSpace2ResourceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "space", optionSpace2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpoptiondefinitionResource_Type(t *testing.T) {
	var resourceName = "nios_dhcp_optiondefinition.test_type"
	var v dhcp.Dhcpoptiondefinition
	name := acctest.RandomNameWithPrefix("dhcp-option-definition")
	space := acctest.RandomNameWithPrefix("dhcp-option-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpoptiondefinitionType("10", name, "string", space),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "string"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpoptiondefinitionType("10", name, "ip-address", space),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "ip-address"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpoptiondefinitionType("10", name, "8-bit unsigned integer", space),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "8-bit unsigned integer"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpoptiondefinitionType("10", name, "domain-name", space),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "domain-name"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpoptiondefinitionType("10", name, "array of ip-address", space),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "array of ip-address"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpoptiondefinitionType("10", name, "boolean", space),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "boolean"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDhcpoptiondefinitionExists(ctx context.Context, resourceName string, v *dhcp.Dhcpoptiondefinition) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			DhcpoptiondefinitionAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDhcpoptiondefinition).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDhcpoptiondefinitionResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDhcpoptiondefinitionResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDhcpoptiondefinitionDestroy(ctx context.Context, v *dhcp.Dhcpoptiondefinition) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			DhcpoptiondefinitionAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDhcpoptiondefinition).
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

func testAccCheckDhcpoptiondefinitionDisappears(ctx context.Context, v *dhcp.Dhcpoptiondefinition) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			DhcpoptiondefinitionAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDhcpoptiondefinitionBasicConfig(code, name, optionType, optionSpace string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_optiondefinition" "test" {
    code = %q
    name = %q
    type = %q
	space = nios_dhcp_optionspace.test.name
}
`, code, name, optionType)
	return strings.Join([]string{testAccBaseWithDHCPOptionSpace(optionSpace), config}, "")
}

func testAccDhcpoptiondefinitionCode(code string, name string, optionType, optionSpace string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_optiondefinition" "test_code" {
    code = %q
    name = %q
    type = %q
	space = nios_dhcp_optionspace.test.name
}
`, code, name, optionType)
	return strings.Join([]string{testAccBaseWithDHCPOptionSpace(optionSpace), config}, "")
}

func testAccDhcpoptiondefinitionName(code string, name string, optionType, optionSpace string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_optiondefinition" "test_name" {
    code = %q
    name = %q
    type = %q
	space = nios_dhcp_optionspace.test.name
}
`, code, name, optionType)
	return strings.Join([]string{testAccBaseWithDHCPOptionSpace(optionSpace), config}, "")
}

func testAccDhcpoptiondefinitionSpace(code string, name string, optionType string, optionSpace1, optionSpace2, optionSpaceResourceName string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_optiondefinition" "test_space" {
    code = %q
    name = %q
    type = %q
    space = %s.name
}
`, code, name, optionType, optionSpaceResourceName)
	return strings.Join([]string{testAccBaseWithTwoDHCPOptionSpaces(optionSpace1, optionSpace2), config}, "")
}

func testAccDhcpoptiondefinitionType(code string, name string, optionType, optionSpace string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_optiondefinition" "test_type" {
    code = %q
    name = %q
    type = %q
	space = nios_dhcp_optionspace.test.name
}
`, code, name, optionType)
	return strings.Join([]string{testAccBaseWithDHCPOptionSpace(optionSpace), config}, "")
}

func testAccBaseWithDHCPOptionSpace(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_optionspace" "test" {
  name = %q
}`, name)
}

func testAccBaseWithTwoDHCPOptionSpaces(name1, name2 string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_optionspace" "test1" {
  name = %q
}

resource "nios_dhcp_optionspace" "test2" {
  name = %q
}`, name1, name2)
}
