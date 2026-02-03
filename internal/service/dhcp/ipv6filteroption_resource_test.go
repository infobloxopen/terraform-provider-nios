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

var readableAttributesForIpv6filteroption = "apply_as_class,comment,expression,extattrs,lease_time,name,option_list,option_space"

func TestAccIpv6filteroptionResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6filteroption.test"
	var v dhcp.Ipv6filteroption
	name := acctest.RandomNameWithPrefix("ipv6filteroption")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6filteroptionBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6filteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "apply_as_class", "true"),
					resource.TestCheckResourceAttr(resourceName, "option_space", "DHCPv6"),
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "expression", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6filteroptionResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_ipv6filteroption.test"
	var v dhcp.Ipv6filteroption
	name := acctest.RandomNameWithPrefix("ipv6filteroption")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6filteroptionDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6filteroptionBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6filteroptionExists(context.Background(), resourceName, &v),
					testAccCheckIpv6filteroptionDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccIpv6filteroptionResource_Import(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6filteroption.test"
	var v dhcp.Ipv6filteroption
	name := acctest.RandomNameWithPrefix("ipv6filteroption")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6filteroptionBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6filteroptionExists(context.Background(), resourceName, &v),
				),
			},
			// Import with PlanOnly to detect differences
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccIpv6filteroptionImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "ref",
				PlanOnly:                             true,
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccIpv6filteroptionImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6filteroptionResource_ApplyAsClass(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6filteroption.test_apply_as_class"
	var v dhcp.Ipv6filteroption
	name := acctest.RandomNameWithPrefix("ipv6filteroption")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6filteroptionApplyAsClass(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6filteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "apply_as_class", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6filteroptionApplyAsClass(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6filteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "apply_as_class", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6filteroptionResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6filteroption.test_comment"
	var v dhcp.Ipv6filteroption
	name := acctest.RandomNameWithPrefix("ipv6filteroption")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6filteroptionComment(name, "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6filteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6filteroptionComment(name, "Updated comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6filteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the object"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6filteroptionResource_Expression(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6filteroption.test_expression"
	var v dhcp.Ipv6filteroption
	name := acctest.RandomNameWithPrefix("ipv6filteroption")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6filteroptionExpression(name, "(option dhcp6.server-id=\"server-id\")"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6filteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "expression", "(option dhcp6.server-id=\"server-id\")"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6filteroptionExpression(name, "(option dhcp6.server-id=\"server-id\" AND option dhcp6.vendor-class=\"DHCPv6\")"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6filteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "expression", "(option dhcp6.server-id=\"server-id\" AND option dhcp6.vendor-class=\"DHCPv6\")"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6filteroptionResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6filteroption.test_extattrs"
	var v dhcp.Ipv6filteroption
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()
	name := acctest.RandomNameWithPrefix("ipv6filteroption")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6filteroptionExtAttrs(name, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6filteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6filteroptionExtAttrs(name, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6filteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6filteroptionResource_LeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6filteroption.test_lease_time"
	var v dhcp.Ipv6filteroption
	name := acctest.RandomNameWithPrefix("ipv6filteroption")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6filteroptionLeaseTime(name, "1000"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6filteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_time", "1000"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6filteroptionLeaseTime(name, "3000"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6filteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_time", "3000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccIpv6filteroptionResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6filteroption.test_name"
	var v dhcp.Ipv6filteroption
	name := acctest.RandomNameWithPrefix("ipv6filteroption")
	nameUpdated := acctest.RandomNameWithPrefix("ipv6filteroption")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6filteroptionName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6filteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6filteroptionName(nameUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6filteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6filteroptionResource_OptionList(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6filteroption.test_option_list"
	var v dhcp.Ipv6filteroption
	name := acctest.RandomNameWithPrefix("ipv6filteroption")
	optionList := []map[string]any{
		{
			"name":  "dhcp6.subscriber-id",
			"value": "subscriber-id",
		},
		{
			"num":   23,
			"value": "fc00::,2001:db8::",
		},
		{
			"name":         "dhcp6.remote-id",
			"num":          37,
			"value":        "remote-id",
			"vendor_class": "DHCPv6",
		},
		{
			"name":         "dhcp6.fqdn",
			"num":          39,
			"value":        "example.com",
			"vendor_class": "DHCPv6",
		},
	}
	optionListUpdated := []map[string]any{
		{
			"name":         "dhcp6.remote-id",
			"num":          37,
			"value":        "remote-id",
			"vendor_class": "DHCPv6",
		},
		{
			"name":  "dhcp6.subscriber-id",
			"value": "subscriber-id",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6filteroptionOptionList(name, optionList),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6filteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_list.#", "4"),
					resource.TestCheckResourceAttr(resourceName, "option_list.0.name", "dhcp6.subscriber-id"),
					resource.TestCheckResourceAttr(resourceName, "option_list.0.value", "subscriber-id"),
					resource.TestCheckResourceAttr(resourceName, "option_list.1.num", "23"),
					resource.TestCheckResourceAttr(resourceName, "option_list.1.value", "fc00::,2001:db8::"),
					resource.TestCheckResourceAttr(resourceName, "option_list.2.vendor_class", "DHCPv6"),
					resource.TestCheckResourceAttr(resourceName, "option_list.2.name", "dhcp6.remote-id"),
					resource.TestCheckResourceAttr(resourceName, "option_list.2.num", "37"),
					resource.TestCheckResourceAttr(resourceName, "option_list.2.value", "remote-id"),
					resource.TestCheckResourceAttr(resourceName, "option_list.3.vendor_class", "DHCPv6"),
					resource.TestCheckResourceAttr(resourceName, "option_list.3.name", "dhcp6.fqdn"),
					resource.TestCheckResourceAttr(resourceName, "option_list.3.num", "39"),
					resource.TestCheckResourceAttr(resourceName, "option_list.3.value", "example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6filteroptionOptionList(name, optionListUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6filteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_list.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "option_list.0.vendor_class", "DHCPv6"),
					resource.TestCheckResourceAttr(resourceName, "option_list.0.name", "dhcp6.remote-id"),
					resource.TestCheckResourceAttr(resourceName, "option_list.0.num", "37"),
					resource.TestCheckResourceAttr(resourceName, "option_list.0.value", "remote-id"),
					resource.TestCheckResourceAttr(resourceName, "option_list.1.name", "dhcp6.subscriber-id"),
					resource.TestCheckResourceAttr(resourceName, "option_list.1.value", "subscriber-id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6filteroptionResource_OptionSpace(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6filteroption.test_option_space"
	var v dhcp.Ipv6filteroption
	name := acctest.RandomNameWithPrefix("ipv6filteroption")
	optionSpace := acctest.RandomNameWithPrefix("ipv6_option_space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6filteroptionOptionSpace(name, "DHCPv6"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6filteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_space", "DHCPv6"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6filteroptionOptionSpaceUpdated(name, optionSpace),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6filteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_space", optionSpace),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckIpv6filteroptionExists(ctx context.Context, resourceName string, v *dhcp.Ipv6filteroption) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			Ipv6filteroptionAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForIpv6filteroption).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetIpv6filteroptionResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetIpv6filteroptionResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckIpv6filteroptionDestroy(ctx context.Context, v *dhcp.Ipv6filteroption) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			Ipv6filteroptionAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForIpv6filteroption).
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

func testAccCheckIpv6filteroptionDisappears(ctx context.Context, v *dhcp.Ipv6filteroption) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			Ipv6filteroptionAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccIpv6filteroptionImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.Attributes["ref"] == "" {
			return "", fmt.Errorf("ref is not set")
		}
		return rs.Primary.Attributes["ref"], nil
	}
}

func testAccIpv6filteroptionBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6filteroption" "test" {
    name = %q
}
`, name)
}

func testAccIpv6filteroptionApplyAsClass(name string, applyAsClass string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6filteroption" "test_apply_as_class" {
    name = %q
    apply_as_class = %q
}
`, name, applyAsClass)
}

func testAccIpv6filteroptionComment(name string, comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6filteroption" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccIpv6filteroptionExpression(name string, expression string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6filteroption" "test_expression" {
    name = %q
    expression = %q
}
`, name, expression)
}

func testAccIpv6filteroptionExtAttrs(name string, extAttrs map[string]string) string {
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6filteroption" "test_extattrs" {
    name = %q
    extattrs = %s
}
`, name, extAttrsStr)
}

func testAccIpv6filteroptionLeaseTime(name string, leaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6filteroption" "test_lease_time" {
    name = %q
    lease_time = %q
}
`, name, leaseTime)
}

func testAccIpv6filteroptionName(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6filteroption" "test_name" {
    name = %q
}
`, name)
}

func testAccIpv6filteroptionOptionList(name string, optionList []map[string]any) string {
	optionListStr := utils.ConvertSliceOfMapsToHCL(optionList)
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6filteroption" "test_option_list" {
    name = %q
    option_list = %s
}
`, name, optionListStr)
}

func testAccIpv6filteroptionOptionSpace(name string, optionSpace string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6filteroption" "test_option_space" {
    name = %q
    option_space = %q
}
`, name, optionSpace)
}

func testAccIpv6filteroptionOptionSpaceUpdated(name string, optionSpace string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6filteroption" "test_option_space" {
    name = %q
    option_space = nios_dhcp_ipv6optionspace.test.name
}
`, name)
	return strings.Join([]string{testAccBaseWithIpv6DHCPOptionSpace(optionSpace, "10"), config}, "")
}
