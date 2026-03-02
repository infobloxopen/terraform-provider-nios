package microsoft_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/microsoft"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForMssuperscope = "comment,dhcp_utilization,dhcp_utilization_status,disable,dynamic_hosts,extattrs,high_water_mark,high_water_mark_reset,low_water_mark,low_water_mark_reset,name,network_view,ranges,static_hosts,total_hosts"

func TestAccMssuperscopeResource_basic(t *testing.T) {
	var resourceName = "nios_microsoft_mssuperscope.test"
	var v microsoft.Mssuperscope

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeBasicConfig("NAME_REPLACE_ME", "RANGES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
					resource.TestCheckResourceAttr(resourceName, "ranges", "RANGES_REPLACE_ME"),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMssuperscopeResource_disappears(t *testing.T) {
	resourceName := "nios_microsoft_mssuperscope.test"
	var v microsoft.Mssuperscope

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMssuperscopeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMssuperscopeBasicConfig("NAME_REPLACE_ME", "RANGES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					testAccCheckMssuperscopeDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccMssuperscopeResource_Import(t *testing.T) {
	var resourceName = "nios_microsoft_mssuperscope.test"
	var v microsoft.Mssuperscope

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeBasicConfig("NAME_REPLACE_ME", "RANGES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
				),
			},
			// Import with PlanOnly to detect differences
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccMssuperscopeImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "ref",
				PlanOnly:                             true,
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccMssuperscopeImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMssuperscopeResource_Comment(t *testing.T) {
	var resourceName = "nios_microsoft_mssuperscope.test_comment"
	var v microsoft.Mssuperscope

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeComment("NAME_REPLACE_ME", "RANGES_REPLACE_ME", "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccMssuperscopeComment("NAME_REPLACE_ME", "RANGES_REPLACE_ME", "Updated comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the object"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMssuperscopeResource_Disable(t *testing.T) {
	var resourceName = "nios_microsoft_mssuperscope.test_disable"
	var v microsoft.Mssuperscope

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeDisable("NAME_REPLACE_ME", "RANGES_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMssuperscopeDisable("NAME_REPLACE_ME", "RANGES_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMssuperscopeResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_microsoft_mssuperscope.test_extattrs"
	var v microsoft.Mssuperscope
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeExtAttrs("NAME_REPLACE_ME", "RANGES_REPLACE_ME", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccMssuperscopeExtAttrs("NAME_REPLACE_ME", "RANGES_REPLACE_ME", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMssuperscopeResource_Name(t *testing.T) {
	var resourceName = "nios_microsoft_mssuperscope.test_name"
	var v microsoft.Mssuperscope

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeName("NAME_REPLACE_ME", "RANGES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMssuperscopeName("NAME_REPLACE_ME", "RANGES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMssuperscopeResource_NetworkView(t *testing.T) {
	var resourceName = "nios_microsoft_mssuperscope.test_network_view"
	var v microsoft.Mssuperscope

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeNetworkView("NAME_REPLACE_ME", "RANGES_REPLACE_ME", "NETWORK_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMssuperscopeNetworkView("NAME_REPLACE_ME", "RANGES_REPLACE_ME", "NETWORK_VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMssuperscopeResource_Ranges(t *testing.T) {
	var resourceName = "nios_microsoft_mssuperscope.test_ranges"
	var v microsoft.Mssuperscope
	rangesVal := []string{"RANGES_REPLACE_ME1", "RANGES_REPLACE_ME2"}
	rangesValUpdated := []string{"RANGES_REPLACE_ME1", "RANGES_REPLACE_ME2"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeRanges("NAME_REPLACE_ME", rangesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ranges", "RANGES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMssuperscopeRanges("NAME_REPLACE_ME", rangesValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ranges", "RANGES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckMssuperscopeExists(ctx context.Context, resourceName string, v *microsoft.Mssuperscope) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.MicrosoftAPI.
			MssuperscopeAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForMssuperscope).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetMssuperscopeResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetMssuperscopeResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckMssuperscopeDestroy(ctx context.Context, v *microsoft.Mssuperscope) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.MicrosoftAPI.
			MssuperscopeAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForMssuperscope).
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

func testAccCheckMssuperscopeDisappears(ctx context.Context, v *microsoft.Mssuperscope) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.MicrosoftAPI.
			MssuperscopeAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccMssuperscopeImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccMssuperscopeBasicConfig(name, ranges string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_mssuperscope" "test" {
    name = %q
    ranges = %q
}
`, name, ranges)
}

func testAccMssuperscopeComment(name string, ranges string, comment string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_mssuperscope" "test_comment" {
    name = %q
    ranges = %q
    comment = %q
}
`, name, ranges, comment)
}

func testAccMssuperscopeDisable(name string, ranges string, disable string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_mssuperscope" "test_disable" {
    name = %q
    ranges = %q
    disable = %q
}
`, name, ranges, disable)
}

func testAccMssuperscopeExtAttrs(name string, ranges string, extAttrs map[string]string) string {
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	return fmt.Sprintf(`
resource "nios_microsoft_mssuperscope" "test_extattrs" {
    name = %q
    ranges = %q
    extattrs = %s
}
`, name, ranges, extAttrsStr)
}

func testAccMssuperscopeName(name string, ranges string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_mssuperscope" "test_name" {
    name = %q
    ranges = %q
}
`, name, ranges)
}

func testAccMssuperscopeNetworkView(name string, ranges string, networkView string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_mssuperscope" "test_network_view" {
    name = %q
    ranges = %q
    network_view = %q
}
`, name, ranges, networkView)
}

func testAccMssuperscopeRanges(name string, ranges []string) string {
	rangesStr := utils.ConvertStringSliceToHCL(ranges)
	return fmt.Sprintf(`
resource "nios_microsoft_mssuperscope" "test_ranges" {
    name = %q
    ranges = %q
}
`, name, rangesStr)
}
