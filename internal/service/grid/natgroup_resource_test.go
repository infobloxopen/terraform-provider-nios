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

var readableAttributesForNatgroup = "comment,name"

func TestAccNatgroupResource_basic(t *testing.T) {
	var resourceName = "nios_grid_natgroup.test"
	var v grid.Natgroup
	name := acctest.RandomNameWithPrefix("natgroup")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNatgroupBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNatgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNatgroupResource_disappears(t *testing.T) {
	resourceName := "nios_grid_natgroup.test"
	var v grid.Natgroup
	name := acctest.RandomNameWithPrefix("natgroup")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNatgroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNatgroupBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNatgroupExists(context.Background(), resourceName, &v),
					testAccCheckNatgroupDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNatgroupResource_Comment(t *testing.T) {
	var resourceName = "nios_grid_natgroup.test_comment"
	var v grid.Natgroup
	name := acctest.RandomNameWithPrefix("natgroup")
	comment1 := "This is a new nat group"
	comment2 := "This is a updated nat group"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNatgroupComment(name, comment1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNatgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment1),
				),
			},
			// Update and Read
			{
				Config: testAccNatgroupComment(name, comment2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNatgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNatgroupResource_Name(t *testing.T) {
	var resourceName = "nios_grid_natgroup.test_name"
	var v grid.Natgroup
	name1 := acctest.RandomNameWithPrefix("natgroup")
	name2 := acctest.RandomNameWithPrefix("natgroup")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNatgroupName(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNatgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccNatgroupName(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNatgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckNatgroupExists(ctx context.Context, resourceName string, v *grid.Natgroup) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.GridAPI.
			NatgroupAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForNatgroup).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetNatgroupResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetNatgroupResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckNatgroupDestroy(ctx context.Context, v *grid.Natgroup) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.GridAPI.
			NatgroupAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForNatgroup).
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

func testAccCheckNatgroupDisappears(ctx context.Context, v *grid.Natgroup) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.GridAPI.
			NatgroupAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccNatgroupBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_grid_natgroup" "test" {
	name = %q
}
`, name)
}

func testAccNatgroupComment(name, comment string) string {
	return fmt.Sprintf(`
resource "nios_grid_natgroup" "test_comment" {
	name = %q
    comment = %q
}
`, name, comment)
}

func testAccNatgroupName(name string) string {
	return fmt.Sprintf(`
resource "nios_grid_natgroup" "test_name" {
    name = %q
}
`, name)
}
