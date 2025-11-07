package grid_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForMemberdfp = "dfp_forward_first,extattrs,host_name,is_dfp_override"

func TestAccMemberdfpResource_basic(t *testing.T) {
	var resourceName = "nios_grid_memberdfp.test"
	var v grid.Memberdfp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberdfpBasicConfig("grid"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberdfpExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberdfpResource_disappears(t *testing.T) {
	resourceName := "nios_grid_memberdfp.test"
	var v grid.Memberdfp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMemberdfpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMemberdfpBasicConfig("grid"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberdfpExists(context.Background(), resourceName, &v),
					testAccCheckMemberdfpDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccMemberdfpResource_Ref(t *testing.T) {
	var resourceName = "nios_grid_memberdfp.test_ref"
	var v grid.Memberdfp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberdfpRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberdfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberdfpRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberdfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberdfpResource_DfpForwardFirst(t *testing.T) {
	var resourceName = "nios_grid_memberdfp.test_dfp_forward_first"
	var v grid.Memberdfp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberdfpDfpForwardFirst("DFP_FORWARD_FIRST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberdfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dfp_forward_first", "DFP_FORWARD_FIRST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberdfpDfpForwardFirst("DFP_FORWARD_FIRST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberdfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dfp_forward_first", "DFP_FORWARD_FIRST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberdfpResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_grid_memberdfp.test_extattrs"
	var v grid.Memberdfp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberdfpExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberdfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberdfpExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberdfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberdfpResource_IsDfpOverride(t *testing.T) {
	var resourceName = "nios_grid_memberdfp.test_is_dfp_override"
	var v grid.Memberdfp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberdfpIsDfpOverride("IS_DFP_OVERRIDE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberdfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "is_dfp_override", "IS_DFP_OVERRIDE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberdfpIsDfpOverride("IS_DFP_OVERRIDE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberdfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "is_dfp_override", "IS_DFP_OVERRIDE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckMemberdfpExists(ctx context.Context, resourceName string, v *grid.Memberdfp) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.GridAPI.
			MemberdfpAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForMemberdfp).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetMemberdfpResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetMemberdfpResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckMemberdfpDestroy(ctx context.Context, v *grid.Memberdfp) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.GridAPI.
			MemberdfpAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForMemberdfp).
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

func testAccCheckMemberdfpDisappears(ctx context.Context, v *grid.Memberdfp) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.GridAPI.
			MemberdfpAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccMemberdfpBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_grid_memberdfp" "test" {
}
`)
}

func testAccMemberdfpRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_grid_memberdfp" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccMemberdfpDfpForwardFirst(dfpForwardFirst string) string {
	return fmt.Sprintf(`
resource "nios_grid_memberdfp" "test_dfp_forward_first" {
    dfp_forward_first = %q
}
`, dfpForwardFirst)
}

func testAccMemberdfpExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_grid_memberdfp" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccMemberdfpIsDfpOverride(isDfpOverride string) string {
	return fmt.Sprintf(`
resource "nios_grid_memberdfp" "test_is_dfp_override" {
    is_dfp_override = %q
}
`, isDfpOverride)
}
