package grid_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// NOTE : Ensure two member are created to run the tests	
//		  - infoblox.member1 
// 		  - infoblox.member2

var readableAttributesForGmcgroup = "comment,gmc_promotion_policy,members,name,scheduled_time,time_zone"

func TestAccGmcgroupResource_basic(t *testing.T) {
	var resourceName = "nios_grid_gmcgroup.test"
	var v grid.Gmcgroup
	name := acctest.RandomNameWithPrefix("gmcgroup")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGmcgroupBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "gmc_promotion_policy", "SIMULTANEOUSLY"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccGmcgroupResource_disappears(t *testing.T) {
	resourceName := "nios_grid_gmcgroup.test"
	var v grid.Gmcgroup
	name := acctest.RandomNameWithPrefix("gmcgroup")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGmcgroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccGmcgroupBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					testAccCheckGmcgroupDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// TO DO
func TestAccGmcgroupResource_Import(t *testing.T) {
	var resourceName = "nios_grid_gmcgroup.test"
	var v grid.Gmcgroup
	name := acctest.RandomNameWithPrefix("gmcgroup")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGmcgroupBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
				),
			},
			// Import with PlanOnly to detect differences
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccGmcgroupImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "ref",
				PlanOnly:                             true,
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccGmcgroupImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccGmcgroupResource_Comment(t *testing.T) {
	var resourceName = "nios_grid_gmcgroup.test_comment"
	var v grid.Gmcgroup
	name := acctest.RandomNameWithPrefix("gmcgroup")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGmcgroupComment(name, "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccGmcgroupComment(name, "Updated comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the object"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccGmcgroupResource_GmcPromotionPolicy(t *testing.T) {
	var resourceName = "nios_grid_gmcgroup.test_gmc_promotion_policy"
	var v grid.Gmcgroup
	name := acctest.RandomNameWithPrefix("gmcgroup")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGmcgroupGmcPromotionPolicy(name, "SEQUENTIALLY"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "gmc_promotion_policy", "SEQUENTIALLY"),
				),
			},
			{
				Config: testAccGmcgroupGmcPromotionPolicy(name, "SIMULTANEOUSLY"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "gmc_promotion_policy", "SIMULTANEOUSLY"),
				),
			},
		},
	})
}

func TestAccGmcgroupResource_Members(t *testing.T) {
	var resourceName = "nios_grid_gmcgroup.test_members"
	var v grid.Gmcgroup
	membersVal := []map[string]any{
		{
			"member": "infoblox.member1",
		},
	}
	membersValUpdated := []map[string]any{
		{
			"member": "infoblox.member1",
		},
		{
			"member": "infoblox.member2",
		},
	}
	name := acctest.RandomNameWithPrefix("gmcgroup")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGmcgroupMembers(name, membersVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "members.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "members.0.member", "infoblox.member1"),
				),
			},
			// Update and Read
			{
				Config: testAccGmcgroupMembers(name, membersValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "members.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "members.0.member", "infoblox.member1"),
					resource.TestCheckResourceAttr(resourceName, "members.1.member", "infoblox.member2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccGmcgroupResource_Name(t *testing.T) {
	var resourceName = "nios_grid_gmcgroup.test_name"
	var v grid.Gmcgroup
	name := acctest.RandomNameWithPrefix("gmcgroup")
	name_update := acctest.RandomNameWithPrefix("gmcgroup_update")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGmcgroupName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccGmcgroupName(name_update),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name_update),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccGmcgroupResource_ScheduledTime(t *testing.T) {
	var resourceName = "nios_grid_gmcgroup.test_scheduled_time"
	var v grid.Gmcgroup
	name := acctest.RandomNameWithPrefix("gmcgroup")
	scheduled_time1 := time.Now().Add(3 * time.Hour).Unix()
	scheduled_time2 := time.Now().Add(4 * time.Hour).Unix()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGmcgroupScheduledTime(name, fmt.Sprintf("%d", scheduled_time1)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "scheduled_time", fmt.Sprintf("%d", scheduled_time1)),
				),
			},
			// Update and Read
			{
				Config: testAccGmcgroupScheduledTime(name, fmt.Sprintf("%d", scheduled_time2)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "scheduled_time", fmt.Sprintf("%d", scheduled_time2)),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckGmcgroupExists(ctx context.Context, resourceName string, v *grid.Gmcgroup) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.GridAPI.
			GmcgroupAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForGmcgroup).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetGmcgroupResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetGmcgroupResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckGmcgroupDestroy(ctx context.Context, v *grid.Gmcgroup) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.GridAPI.
			GmcgroupAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForGmcgroup).
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

func testAccCheckGmcgroupDisappears(ctx context.Context, v *grid.Gmcgroup) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.GridAPI.
			GmcgroupAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccGmcgroupImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccGmcgroupBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_grid_gmcgroup" "test" {
    name = %q
}
`, name)
}

func testAccGmcgroupComment(name string, comment string) string {
	return fmt.Sprintf(`
resource "nios_grid_gmcgroup" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccGmcgroupGmcPromotionPolicy(name string, gmcPromotionPolicy string) string {
	return fmt.Sprintf(`
resource "nios_grid_gmcgroup" "test_gmc_promotion_policy" {
    name = %q
    gmc_promotion_policy = %q
}
`, name, gmcPromotionPolicy)
}

func testAccGmcgroupMembers(name string, members []map[string]any) string {
	membersStr := utils.ConvertSliceOfMapsToHCL(members)
	return fmt.Sprintf(`
resource "nios_grid_gmcgroup" "test_members" {
    name = %q
    members = %s
}
`, name, membersStr)
}

func testAccGmcgroupName(name string) string {
	return fmt.Sprintf(`
resource "nios_grid_gmcgroup" "test_name" {
    name = %q
}
`, name)
}

func testAccGmcgroupScheduledTime(name string, scheduledTime string) string {
	return fmt.Sprintf(`
resource "nios_grid_gmcgroup" "test_scheduled_time" {
    name = %q
    scheduled_time = %q
}
`, name, scheduledTime)
}
