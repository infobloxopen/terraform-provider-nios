package grid_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForMemberdfp = "dfp_forward_first,host_name,is_dfp_override"

func TestAccMemberdfpResource_basic(t *testing.T) {
	var resourceName = "nios_grid_memberdfp.test"
	var v grid.Memberdfp
	host_name := "infoblox.member1"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberdfpBasicConfig(host_name, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberdfpExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "dfp_forward_first", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_dfp_override", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberdfpResource_disappears(t *testing.T) {
	t.Skip("memberdfp cannot be deleted from NIOS, skipping disappears test")
}

func TestAccMemberdfpResource_DfpForwardFirst(t *testing.T) {
	var resourceName = "nios_grid_memberdfp.test_dfp_forward_first"
	var v grid.Memberdfp
	host_name := "infoblox.member1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberdfpDfpForwardFirst(host_name, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberdfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dfp_forward_first", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberdfpDfpForwardFirst(host_name, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberdfpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dfp_forward_first", "false"),
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

func testAccMemberdfpBasicConfig(host_name string, isDfpOverride bool) string {
	return fmt.Sprintf(`
resource "nios_grid_memberdfp" "test" {
	host_name = %q
	is_dfp_override = %t
}
`, host_name, isDfpOverride)
}

func testAccMemberdfpDfpForwardFirst(host_name string, dfpForwardFirst bool) string {
	return fmt.Sprintf(`
resource "nios_grid_memberdfp" "test_dfp_forward_first" {
    host_name = %q
	dfp_forward_first = %t
}
`, host_name, dfpForwardFirst)
}
