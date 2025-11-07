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

var readableAttributesForGridServicerestartGroupOrder = "groups"

func TestAccGridServicerestartGroupOrderResource_basic(t *testing.T) {
	var resourceName = "nios_grid_servicerestart_group_order.test"
	var v grid.GridServicerestartGroupOrder

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGridServicerestartGroupOrderBasicConfig("grid"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupOrderExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccGridServicerestartGroupOrderResource_disappears(t *testing.T) {
	resourceName := "nios_grid_servicerestart_group_order.test"
	var v grid.GridServicerestartGroupOrder

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGridServicerestartGroupOrderDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccGridServicerestartGroupOrderBasicConfig("grid"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupOrderExists(context.Background(), resourceName, &v),
					//testAccCheckGridServicerestartGroupOrderDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccGridServicerestartGroupOrderResource_Ref(t *testing.T) {
	var resourceName = "nios_grid_servicerestart_group_order.test_ref"
	var v grid.GridServicerestartGroupOrder

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGridServicerestartGroupOrderRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupOrderExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccGridServicerestartGroupOrderRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupOrderExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccGridServicerestartGroupOrderResource_Groups(t *testing.T) {
	var resourceName = "nios_grid_servicerestart_group_order.test_groups"
	var v grid.GridServicerestartGroupOrder

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGridServicerestartGroupOrderGroups("GROUPS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupOrderExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "groups", "GROUPS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccGridServicerestartGroupOrderGroups("GROUPS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupOrderExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "groups", "GROUPS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckGridServicerestartGroupOrderExists(ctx context.Context, resourceName string, v *grid.GridServicerestartGroupOrder) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.GridAPI.
			GridServicerestartGroupOrderAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForGridServicerestartGroupOrder).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetGridServicerestartGroupOrderResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetGridServicerestartGroupOrderResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckGridServicerestartGroupOrderDestroy(ctx context.Context, v *grid.GridServicerestartGroupOrder) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.GridAPI.
			GridServicerestartGroupOrderAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForGridServicerestartGroupOrder).
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

func testAccGridServicerestartGroupOrderBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_grid_servicerestart_group_order" "test" {
}
`)
}

func testAccGridServicerestartGroupOrderRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_grid_servicerestart_group_order" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccGridServicerestartGroupOrderGroups(groups string) string {
	return fmt.Sprintf(`
resource "nios_grid_servicerestart_group_order" "test_groups" {
    groups = %q
}
`, groups)
}
