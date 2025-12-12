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

var readableAttributesForFilternac = "comment,expression,extattrs,lease_time,name,options"

func TestAccFilternacResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_filternac.test"
	var v dhcp.Filternac
	name := acctest.RandomNameWithPrefix("tf-filternac-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilternacBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilternacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilternacResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_filternac.test"
	var v dhcp.Filternac
	name := acctest.RandomNameWithPrefix("tf-filternac-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFilternacDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFilternacBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilternacExists(context.Background(), resourceName, &v),
					testAccCheckFilternacDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccFilternacResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_filternac.test_comment"
	var v dhcp.Filternac
	name := acctest.RandomNameWithPrefix("tf-filternac-")
	comment := "Initial comment"
	updatedComment := "Updated comment"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilternacComment(name, comment),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilternacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment),
				),
			},
			// Update and Read
			{
				Config: testAccFilternacComment(name, updatedComment),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilternacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", updatedComment),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilternacResource_Expression(t *testing.T) {
	var resourceName = "nios_dhcp_filternac.test_expression"
	var v dhcp.Filternac
	name := acctest.RandomNameWithPrefix("tf-filternac-")
	expression := `(Radius.ServerError= "false")`
	updatedExpression := `(Radius.ServerError= "true" AND Sophos.ComplianceState = "NonCompliant")`

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilternacExpression(name, expression),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilternacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "expression", expression),
				),
			},
			// Update and Read
			{
				Config: testAccFilternacExpression(name, updatedExpression),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilternacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "expression", updatedExpression),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilternacResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dhcp_filternac.test_extattrs"
	var v dhcp.Filternac
	name := acctest.RandomNameWithPrefix("tf-filternac-")
	site := acctest.RandomNameWithPrefix("Location-")
	updatedSite := acctest.RandomNameWithPrefix("Location-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilternacExtAttrs(name, map[string]string{"Site": site}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilternacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", site),
				),
			},
			// Update and Read
			{
				Config: testAccFilternacExtAttrs(name, map[string]string{"Site": updatedSite}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilternacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", updatedSite),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilternacResource_LeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_filternac.test_lease_time"
	var v dhcp.Filternac
	name := acctest.RandomNameWithPrefix("tf-filternac-")
	lease_time := "3600"
	updated_lease_time := "7200"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilternacLeaseTime(name, lease_time),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilternacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_time", lease_time),
				),
			},
			// Update and Read
			{
				Config: testAccFilternacLeaseTime(name, updated_lease_time),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilternacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_time", updated_lease_time),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilternacResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_filternac.test_name"
	var v dhcp.Filternac
	name := acctest.RandomNameWithPrefix("tf-filternac-")
	updatedName := acctest.RandomNameWithPrefix("tf-filternac-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilternacName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilternacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccFilternacName(updatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilternacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilternacResource_Options(t *testing.T) {
	var resourceName = "nios_dhcp_filternac.test_options"
	var v dhcp.Filternac
	name := acctest.RandomNameWithPrefix("mac_filter")
	options := []map[string]any{
		{
			"name":  "dhcp-lease-time",
			"num":   "51",
			"value": "1200",
		},
		{
			"name":  "time-offset",
			"num":   2,
			"value": "3600",
		},
	}
	updatedOptions := []map[string]any{
		{
			"name":  "dhcp-lease-time",
			"num":   "51",
			"value": "1800",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilternacOptions(name, options),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilternacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr(resourceName, "options.0.num", "51"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "1200"),
					resource.TestCheckResourceAttr(resourceName, "options.1.name", "time-offset"),
					resource.TestCheckResourceAttr(resourceName, "options.1.num", "2"),
					resource.TestCheckResourceAttr(resourceName, "options.1.value", "3600"),
				),
			},
			// Update and Read
			{
				Config: testAccFilternacOptions(name, updatedOptions),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilternacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr(resourceName, "options.0.num", "51"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "1800"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckFilternacExists(ctx context.Context, resourceName string, v *dhcp.Filternac) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			FilternacAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForFilternac).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetFilternacResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetFilternacResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckFilternacDestroy(ctx context.Context, v *dhcp.Filternac) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			FilternacAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForFilternac).
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

func testAccCheckFilternacDisappears(ctx context.Context, v *dhcp.Filternac) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			FilternacAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccFilternacBasicConfig(name string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dhcp_filternac" "test" {
    name = %q
}
`, name)
}

func testAccFilternacComment(name string, comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filternac" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccFilternacExpression(name string, expression string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filternac" "test_expression" {
    name = %q
    expression = %q
}
`, name, expression)
}

func testAccFilternacExtAttrs(name string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf("  %s = %q\n", k, v)
	}
	extattrsStr += "}"
	return fmt.Sprintf(`
resource "nios_dhcp_filternac" "test_extattrs" {
    name = %q
    extattrs = %s
}
`, name, extattrsStr)
}

func testAccFilternacLeaseTime(name string, leaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filternac" "test_lease_time" {
    name = %q
    lease_time = %q
}
`, name, leaseTime)
}

func testAccFilternacName(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filternac" "test_name" {
    name = %q
}
`, name)
}

func testAccFilternacOptions(name string, options []map[string]any) string {
	optionsStr := utils.ConvertSliceOfMapsToHCL(options)
	return fmt.Sprintf(`
resource "nios_dhcp_filternac" "test_options" {
	name = %q
    options = %s
}
`, name, optionsStr)
}
