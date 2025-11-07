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

var readableAttributesForGmcgroup = "comment,gmc_promotion_policy,members,name,scheduled_time,time_zone"

func TestAccGmcgroupResource_basic(t *testing.T) {
	var resourceName = "nios_grid_gmcgroup.test"
	var v grid.Gmcgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGmcgroupBasicConfig("grid"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccGmcgroupResource_disappears(t *testing.T) {
	resourceName := "nios_grid_gmcgroup.test"
	var v grid.Gmcgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGmcgroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccGmcgroupBasicConfig("grid"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					testAccCheckGmcgroupDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccGmcgroupResource_Ref(t *testing.T) {
	var resourceName = "nios_grid_gmcgroup.test_ref"
	var v grid.Gmcgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGmcgroupRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccGmcgroupRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccGmcgroupResource_Comment(t *testing.T) {
	var resourceName = "nios_grid_gmcgroup.test_comment"
	var v grid.Gmcgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGmcgroupComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccGmcgroupComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccGmcgroupResource_GmcPromotionPolicy(t *testing.T) {
	var resourceName = "nios_grid_gmcgroup.test_gmc_promotion_policy"
	var v grid.Gmcgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGmcgroupGmcPromotionPolicy("GMC_PROMOTION_POLICY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "gmc_promotion_policy", "GMC_PROMOTION_POLICY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccGmcgroupGmcPromotionPolicy("GMC_PROMOTION_POLICY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "gmc_promotion_policy", "GMC_PROMOTION_POLICY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccGmcgroupResource_Members(t *testing.T) {
	var resourceName = "nios_grid_gmcgroup.test_members"
	var v grid.Gmcgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGmcgroupMembers("MEMBERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "members", "MEMBERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccGmcgroupMembers("MEMBERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "members", "MEMBERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccGmcgroupResource_Name(t *testing.T) {
	var resourceName = "nios_grid_gmcgroup.test_name"
	var v grid.Gmcgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGmcgroupName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccGmcgroupName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccGmcgroupResource_ScheduledTime(t *testing.T) {
	var resourceName = "nios_grid_gmcgroup.test_scheduled_time"
	var v grid.Gmcgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGmcgroupScheduledTime("SCHEDULED_TIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "scheduled_time", "SCHEDULED_TIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccGmcgroupScheduledTime("SCHEDULED_TIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGmcgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "scheduled_time", "SCHEDULED_TIME_UPDATE_REPLACE_ME"),
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

func testAccGmcgroupBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_grid_gmcgroup" "test" {
}
`)
}

func testAccGmcgroupRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_grid_gmcgroup" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccGmcgroupComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_grid_gmcgroup" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccGmcgroupGmcPromotionPolicy(gmcPromotionPolicy string) string {
	return fmt.Sprintf(`
resource "nios_grid_gmcgroup" "test_gmc_promotion_policy" {
    gmc_promotion_policy = %q
}
`, gmcPromotionPolicy)
}

func testAccGmcgroupMembers(members string) string {
	return fmt.Sprintf(`
resource "nios_grid_gmcgroup" "test_members" {
    members = %q
}
`, members)
}

func testAccGmcgroupName(name string) string {
	return fmt.Sprintf(`
resource "nios_grid_gmcgroup" "test_name" {
    name = %q
}
`, name)
}

func testAccGmcgroupScheduledTime(scheduledTime string) string {
	return fmt.Sprintf(`
resource "nios_grid_gmcgroup" "test_scheduled_time" {
    scheduled_time = %q
}
`, scheduledTime)
}
