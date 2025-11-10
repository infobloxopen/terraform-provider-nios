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

var readableAttributesForIpv6dhcpoptiondefinition = "code,name,space,type"

func TestAccIpv6dhcpoptiondefinitionResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6dhcpoptiondefinition.test"
	var v dhcp.Ipv6dhcpoptiondefinition
	name := acctest.RandomNameWithPrefix("option-definition")
	optionSpace := acctest.RandomNameWithPrefix("option-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6dhcpoptiondefinitionBasicConfig(optionSpace, "10", name, "string"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6dhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "code", "10"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "type", "string"),
					resource.TestCheckResourceAttr(resourceName, "space", optionSpace),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6dhcpoptiondefinitionResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_ipv6dhcpoptiondefinition.test"
	var v dhcp.Ipv6dhcpoptiondefinition
	name := acctest.RandomNameWithPrefix("option-definition")
	optionSpace := acctest.RandomNameWithPrefix("option-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6dhcpoptiondefinitionDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6dhcpoptiondefinitionBasicConfig(optionSpace, "10", name, "string"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6dhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					testAccCheckIpv6dhcpoptiondefinitionDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccIpv6dhcpoptiondefinitionResource_Code(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6dhcpoptiondefinition.test_code"
	var v dhcp.Ipv6dhcpoptiondefinition
	name := acctest.RandomNameWithPrefix("option-definition")
	optionSpace := acctest.RandomNameWithPrefix("option-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6dhcpoptiondefinitionCode(optionSpace, "10", name, "string"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6dhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "code", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6dhcpoptiondefinitionCode(optionSpace, "20", name, "string"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6dhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "code", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6dhcpoptiondefinitionResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6dhcpoptiondefinition.test_name"
	var v dhcp.Ipv6dhcpoptiondefinition
	name := acctest.RandomNameWithPrefix("option-definition")
	updatedOpttionDefinitionName := acctest.RandomNameWithPrefix("option-definition")
	optionSpace := acctest.RandomNameWithPrefix("option-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6dhcpoptiondefinitionName(optionSpace, "10", name, "string"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6dhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6dhcpoptiondefinitionName(optionSpace, "10", updatedOpttionDefinitionName, "string"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6dhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", updatedOpttionDefinitionName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6dhcpoptiondefinitionResource_Space(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6dhcpoptiondefinition.test_space"
	var v dhcp.Ipv6dhcpoptiondefinition
	name := acctest.RandomNameWithPrefix("option-definition")
	optionSpace1 := acctest.RandomNameWithPrefix("option-space")
	optionSpace2 := acctest.RandomNameWithPrefix("option-space")
	optionSpace1ResourceName := "nios_dhcp_ipv6dhcpoptionspace.test1"
	optionSpace2ResourceName := "nios_dhcp_ipv6dhcpoptionspace.test2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6dhcpoptiondefinitionSpace(optionSpace1, optionSpace2, "10", name, "string", optionSpace1ResourceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6dhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "space", optionSpace1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6dhcpoptiondefinitionSpace(optionSpace1, optionSpace2, "10", name, "string", optionSpace2ResourceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6dhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "space", optionSpace2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6dhcpoptiondefinitionResource_Type(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6dhcpoptiondefinition.test_type"
	var v dhcp.Ipv6dhcpoptiondefinition
	name := acctest.RandomNameWithPrefix("option-definition")
	optionSpace := acctest.RandomNameWithPrefix("option-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6dhcpoptiondefinitionType(optionSpace, "10", name, "string"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6dhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "string"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6dhcpoptiondefinitionType(optionSpace, "10", name, "boolean"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6dhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "boolean"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6dhcpoptiondefinitionType(optionSpace, "10", name, "8-bit unsigned integer"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6dhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "8-bit unsigned integer"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckIpv6dhcpoptiondefinitionExists(ctx context.Context, resourceName string, v *dhcp.Ipv6dhcpoptiondefinition) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			Ipv6dhcpoptiondefinitionAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForIpv6dhcpoptiondefinition).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetIpv6dhcpoptiondefinitionResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetIpv6dhcpoptiondefinitionResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckIpv6dhcpoptiondefinitionDestroy(ctx context.Context, v *dhcp.Ipv6dhcpoptiondefinition) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			Ipv6dhcpoptiondefinitionAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForIpv6dhcpoptiondefinition).
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

func testAccCheckIpv6dhcpoptiondefinitionDisappears(ctx context.Context, v *dhcp.Ipv6dhcpoptiondefinition) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			Ipv6dhcpoptiondefinitionAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccIpv6dhcpoptiondefinitionBasicConfig(optionSpace, code, name, optionType string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6dhcpoptiondefinition" "test" {
    code = %q
    name = %q
    type = %q
	space = nios_dhcp_ipv6dhcpoptionspace.test.name
}
`, code, name, optionType)
	return strings.Join([]string{testAccBaseWithIpv6DHCPOptionSpace(optionSpace, "10"), config}, "")
}

func testAccIpv6dhcpoptiondefinitionCode(optionSpace, code string, name string, optionType string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6dhcpoptiondefinition" "test_code" {
    code = %q
    name = %q
    type = %q
	space = nios_dhcp_ipv6dhcpoptionspace.test.name
}
`, code, name, optionType)
	return strings.Join([]string{testAccBaseWithIpv6DHCPOptionSpace(optionSpace, "10"), config}, "")
}

func testAccIpv6dhcpoptiondefinitionName(optionSpace, code string, name string, optionType string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6dhcpoptiondefinition" "test_name" {
    code = %q
    name = %q
    type = %q
	space = nios_dhcp_ipv6dhcpoptionspace.test.name
}
`, code, name, optionType)
	return strings.Join([]string{testAccBaseWithIpv6DHCPOptionSpace(optionSpace, "10"), config}, "")
}

func testAccIpv6dhcpoptiondefinitionSpace(optionSpace1, optionSpace2, code string, name string, optionType, optionSpaceResourceName string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6dhcpoptiondefinition" "test_space" {
    code = %q
    name = %q
    type = %q
	space = %s.name
}
`, code, name, optionType, optionSpaceResourceName)
	return strings.Join([]string{testAccBaseWithTwoIpv6DHCPOptionSpaces(optionSpace1, "10", optionSpace2, "20"), config}, "")
}

func testAccIpv6dhcpoptiondefinitionType(optionSpace, code string, name string, optionType string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6dhcpoptiondefinition" "test_type" {
    code = %q
    name = %q
    type = %q
	space = nios_dhcp_ipv6dhcpoptionspace.test.name
}
`, code, name, optionType)
	return strings.Join([]string{testAccBaseWithIpv6DHCPOptionSpace(optionSpace, "10"), config}, "")
}

func testAccBaseWithIpv6DHCPOptionSpace(name, enterpriseNumber string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6dhcpoptionspace" "test" {
  name = %q
  enterprise_number = %q
}`, name, enterpriseNumber)
}

func testAccBaseWithTwoIpv6DHCPOptionSpaces(name1, enterpriseNumber1, name2, enterpriseNumber2 string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6dhcpoptionspace" "test1" {
  name = %q
  enterprise_number = %q
}
resource "nios_dhcp_ipv6dhcpoptionspace" "test2" {
  name = %q
  enterprise_number = %q
}
`, name1, enterpriseNumber1, name2, enterpriseNumber2)
}
