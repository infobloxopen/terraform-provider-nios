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

var readableAttributesForFilteroption = "apply_as_class,bootfile,bootserver,comment,expression,extattrs,lease_time,name,next_server,option_list,option_space,pxe_lease_time"

func TestAccFilteroptionResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_filteroption.test"
	var v dhcp.Filteroption
	name := acctest.RandomNameWithPrefix("filteroption")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilteroptionBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "apply_as_class", "true"),
					resource.TestCheckResourceAttr(resourceName, "option_space", "DHCP"),
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "expression", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilteroptionResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_filteroption.test"
	var v dhcp.Filteroption
	name := acctest.RandomNameWithPrefix("filteroption")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFilteroptionDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFilteroptionBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					testAccCheckFilteroptionDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccFilteroptionResource_ApplyAsClass(t *testing.T) {
	var resourceName = "nios_dhcp_filteroption.test_apply_as_class"
	var v dhcp.Filteroption
	name := acctest.RandomNameWithPrefix("filteroption")
	applyAsClass := "true"
	updatedApplyAsClass := "false"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilteroptionApplyAsClass(name, applyAsClass),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "apply_as_class", applyAsClass),
				),
			},
			// Update and Read
			{
				Config: testAccFilteroptionApplyAsClass(name, updatedApplyAsClass),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "apply_as_class", updatedApplyAsClass),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilteroptionResource_Bootfile(t *testing.T) {
	var resourceName = "nios_dhcp_filteroption.test_bootfile"
	var v dhcp.Filteroption
	name := acctest.RandomNameWithPrefix("filteroption")
	bootfile := "BOOTFILE_REPLACE_ME"
	updatedBootfile := "BOOTFILE_UPDATE_REPLACE_ME"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilteroptionBootfile(name, bootfile),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "BOOTFILE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFilteroptionBootfile(name, updatedBootfile),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "BOOTFILE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilteroptionResource_Bootserver(t *testing.T) {
	var resourceName = "nios_dhcp_filteroption.test_bootserver"
	var v dhcp.Filteroption
	name := acctest.RandomNameWithPrefix("filteroption")
	bootserver := "BOOTSERVER_REPLACE_ME"
	updatedBootserver := "BOOTSERVER_UPDATE_REPLACE_ME"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilteroptionBootserver(name, bootserver),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "BOOTSERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFilteroptionBootserver(name, updatedBootserver),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "BOOTSERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilteroptionResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_filteroption.test_comment"
	var v dhcp.Filteroption
	name := acctest.RandomNameWithPrefix("filteroption")
	comment := "COMMENT_REPLACE_ME"
	updatedComment := "COMMENT_UPDATE_REPLACE_ME"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilteroptionComment(name, comment),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFilteroptionComment(name, updatedComment),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilteroptionResource_Expression(t *testing.T) {
	var resourceName = "nios_dhcp_filteroption.test_expression"
	var v dhcp.Filteroption
	name := acctest.RandomNameWithPrefix("filteroption")
	expression := "(option domain-name=\"example.com\")"
	updatedExpression := "(option ntp-servers=\"2.2.2.2\")"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilteroptionExpression(name, expression),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "expression", expression),
				),
			},
			// Update and Read
			{
				Config: testAccFilteroptionExpression(name, updatedExpression),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "expression", updatedExpression),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilteroptionResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dhcp_filteroption.test_extattrs"
	var v dhcp.Filteroption
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()
	name := acctest.RandomNameWithPrefix("filteroption")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilteroptionExtAttrs(name, map[string]string{"Site": extAttrValue1}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccFilteroptionExtAttrs(name, map[string]string{"Site": extAttrValue2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilteroptionResource_LeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_filteroption.test_lease_time"
	var v dhcp.Filteroption
	name := acctest.RandomNameWithPrefix("filteroption")
	leaseTime := "600"
	updatedLeaseTime := "1200"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilteroptionLeaseTime(name, leaseTime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_time", leaseTime),
				),
			},
			// Update and Read
			{
				Config: testAccFilteroptionLeaseTime(name, updatedLeaseTime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_time", updatedLeaseTime),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilteroptionResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_filteroption.test_name"
	var v dhcp.Filteroption
	name := acctest.RandomNameWithPrefix("filteroption")
	updateName := acctest.RandomNameWithPrefix("filteroption")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilteroptionName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccFilteroptionName(updateName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", updateName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilteroptionResource_NextServer(t *testing.T) {
	var resourceName = "nios_dhcp_filteroption.test_next_server"
	var v dhcp.Filteroption
	name := acctest.RandomNameWithPrefix("filteroption")
	nextServer := "NEXT_SERVER_REPLACE_ME"
	updatedNextServer := "NEXT_SERVER_UPDATE_REPLACE_ME"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilteroptionNextServer(name, nextServer),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "next_server", nextServer),
				),
			},
			// Update and Read
			{
				Config: testAccFilteroptionNextServer(name, updatedNextServer),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "next_server", updatedNextServer),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilteroptionResource_OptionList(t *testing.T) {
	var resourceName = "nios_dhcp_filteroption.test_option_list"
	var v dhcp.Filteroption
	name := acctest.RandomNameWithPrefix("filteroption")
	optionList := []map[string]any{
		{
			"name":  "domain-name",
			"value": "example.com",
		},
	}
	updatedOptionList := []map[string]any{
		{
			"name":  "time-offset",
			"num":   2,
			"value": "1200",
		},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilteroptionOptionList(name, optionList),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_list.0.name", "domain-name"),
					resource.TestCheckResourceAttr(resourceName, "option_list.0.value", "example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccFilteroptionOptionList(name, updatedOptionList),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_list.0.name", "time-offset"),
					resource.TestCheckResourceAttr(resourceName, "option_list.0.value", "1200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilteroptionResource_OptionSpace(t *testing.T) {
	var resourceName = "nios_dhcp_filteroption.test_option_space"
	var v dhcp.Filteroption
	name := acctest.RandomNameWithPrefix("filteroption")
	optionSpace := "DHCP"
	updatedOptionSpace := acctest.RandomNameWithPrefix("optionspace")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilteroptionOptionSpace(name, optionSpace),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_space", "DHCP"),
				),
			},
			// Update and Read
			{
				Config: testAccFilteroptionOptionSpaceUpdated(name, updatedOptionSpace),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_space", updatedOptionSpace),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilteroptionResource_PxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_filteroption.test_pxe_lease_time"
	var v dhcp.Filteroption
	name := acctest.RandomNameWithPrefix("filteroption")
	pxeLeaseTime := "1200"
	updatedPxeLeaseTime := "1800"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilteroptionPxeLeaseTime(name, pxeLeaseTime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "1200"),
				),
			},
			// Update and Read
			{
				Config: testAccFilteroptionPxeLeaseTime(name, updatedPxeLeaseTime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "1800"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckFilteroptionExists(ctx context.Context, resourceName string, v *dhcp.Filteroption) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			FilteroptionAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForFilteroption).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetFilteroptionResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetFilteroptionResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckFilteroptionDestroy(ctx context.Context, v *dhcp.Filteroption) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			FilteroptionAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForFilteroption).
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

func testAccCheckFilteroptionDisappears(ctx context.Context, v *dhcp.Filteroption) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			FilteroptionAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccFilteroptionBasicConfig(name string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dhcp_filteroption" "test" {
    name = %q
}
`, name)
}

func testAccFilteroptionApplyAsClass(name string, applyAsClass string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filteroption" "test_apply_as_class" {
    name = %q
    apply_as_class = %q
}
`, name, applyAsClass)
}

func testAccFilteroptionBootfile(name string, bootfile string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filteroption" "test_bootfile" {
    name = %q
    bootfile = %q
}
`, name, bootfile)
}

func testAccFilteroptionBootserver(name string, bootserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filteroption" "test_bootserver" {
    name = %q
    bootserver = %q
}
`, name, bootserver)
}

func testAccFilteroptionComment(name string, comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filteroption" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccFilteroptionExpression(name string, expression string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filteroption" "test_expression" {
    name = %q
    expression = %q
}
`, name, expression)
}

func testAccFilteroptionExtAttrs(name string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf("  %s = %q\n", k, v)
	}
	extattrsStr += "}"
	return fmt.Sprintf(`
resource "nios_dhcp_filteroption" "test_extattrs" {
	name = %q
	extattrs = %s
}
`, name, extattrsStr)
}

func testAccFilteroptionLeaseTime(name string, leaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filteroption" "test_lease_time" {
    name = %q
    lease_time = %q
}
`, name, leaseTime)
}

func testAccFilteroptionName(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filteroption" "test_name" {
    name = %q
}
`, name)
}

func testAccFilteroptionNextServer(name string, nextServer string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filteroption" "test_next_server" {
    name = %q
    next_server = %q
}
`, name, nextServer)
}

func testAccFilteroptionOptionList(name string, optionList []map[string]any) string {
	optionListStr := utils.ConvertSliceOfMapsToHCL(optionList)
	return fmt.Sprintf(`
resource "nios_dhcp_filteroption" "test_option_list" {
    name = %q
    option_list = %s
}
`, name, optionListStr)
}

func testAccFilteroptionOptionSpace(name string, optionSpace string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filteroption" "test_option_space" {
    name = %q
    option_space = %q
}
`, name, optionSpace)
}

func testAccFilteroptionPxeLeaseTime(name string, pxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filteroption" "test_pxe_lease_time" {
    name = %q
    pxe_lease_time = %q
}
`, name, pxeLeaseTime)
}

func testAccFilteroptionOptionSpaceUpdated(name string, optionSpace string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_filteroption" "test_option_space" {
    name = %q
    option_space = nios_dhcp_optionspace.test.name
}
`, name)
	return strings.Join([]string{testAccBaseWithDHCPOptionSpace(optionSpace), config}, "")
}
