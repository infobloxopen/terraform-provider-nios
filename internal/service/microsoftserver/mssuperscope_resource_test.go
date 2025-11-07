package microsoftserver_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/microsoftserver"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForMssuperscope = "comment,dhcp_utilization,dhcp_utilization_status,disable,dynamic_hosts,extattrs,high_water_mark,high_water_mark_reset,low_water_mark,low_water_mark_reset,name,network_view,ranges,static_hosts,total_hosts"

func TestAccMssuperscopeResource_basic(t *testing.T) {
	var resourceName = "nios_microsoftserver_mssuperscope.test"
	var v microsoftserver.Mssuperscope

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMssuperscopeResource_disappears(t *testing.T) {
	resourceName := "nios_microsoftserver_mssuperscope.test"
	var v microsoftserver.Mssuperscope

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMssuperscopeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMssuperscopeBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					testAccCheckMssuperscopeDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccMssuperscopeResource_Ref(t *testing.T) {
	var resourceName = "nios_microsoftserver_mssuperscope.test_ref"
	var v microsoftserver.Mssuperscope

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMssuperscopeRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMssuperscopeResource_Comment(t *testing.T) {
	var resourceName = "nios_microsoftserver_mssuperscope.test_comment"
	var v microsoftserver.Mssuperscope

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMssuperscopeComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMssuperscopeResource_Disable(t *testing.T) {
	var resourceName = "nios_microsoftserver_mssuperscope.test_disable"
	var v microsoftserver.Mssuperscope

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMssuperscopeDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMssuperscopeResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_microsoftserver_mssuperscope.test_extattrs"
	var v microsoftserver.Mssuperscope

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMssuperscopeExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMssuperscopeResource_Name(t *testing.T) {
	var resourceName = "nios_microsoftserver_mssuperscope.test_name"
	var v microsoftserver.Mssuperscope

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMssuperscopeName("NAME_UPDATE_REPLACE_ME"),
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
	var resourceName = "nios_microsoftserver_mssuperscope.test_network_view"
	var v microsoftserver.Mssuperscope

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeNetworkView("NETWORK_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMssuperscopeNetworkView("NETWORK_VIEW_UPDATE_REPLACE_ME"),
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
	var resourceName = "nios_microsoftserver_mssuperscope.test_ranges"
	var v microsoftserver.Mssuperscope

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeRanges("RANGES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ranges", "RANGES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMssuperscopeRanges("RANGES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ranges", "RANGES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckMssuperscopeExists(ctx context.Context, resourceName string, v *microsoftserver.Mssuperscope) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.MicrosoftServerAPI.
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

func testAccCheckMssuperscopeDestroy(ctx context.Context, v *microsoftserver.Mssuperscope) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.MicrosoftServerAPI.
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

func testAccCheckMssuperscopeDisappears(ctx context.Context, v *microsoftserver.Mssuperscope) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.MicrosoftServerAPI.
			MssuperscopeAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccMssuperscopeBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_microsoftserver_mssuperscope" "test" {
}
`)
}

func testAccMssuperscopeRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_mssuperscope" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccMssuperscopeComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_mssuperscope" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccMssuperscopeDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_mssuperscope" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccMssuperscopeExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_mssuperscope" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccMssuperscopeName(name string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_mssuperscope" "test_name" {
    name = %q
}
`, name)
}

func testAccMssuperscopeNetworkView(networkView string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_mssuperscope" "test_network_view" {
    network_view = %q
}
`, networkView)
}

func testAccMssuperscopeRanges(ranges string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_mssuperscope" "test_ranges" {
    ranges = %q
}
`, ranges)
}
