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

// TODO : OBJECTS TO BE PRESENT IN GRID FOR TESTS
// Grid Members: infoblox.172_28_83_29, infoblox.172_28_82_115
// Distribution Dependent Groups: example_distribution_dependent_group1, example_distribution_dependent_group2
// Upgrade Dependent Groups: example_upgrade_dependent_group1, example_upgrade_dependent_group2

var readableAttributesForUpgradegroup = "comment,distribution_dependent_group,distribution_policy,distribution_time,members,name,time_zone,upgrade_dependent_group,upgrade_policy,upgrade_time"

func TestAccUpgradegroupResource_basic(t *testing.T) {
	var resourceName = "nios_grid_upgradegroup.test"
	var v grid.Upgradegroup

	name := acctest.RandomNameWithPrefix("example-upgradegroup-")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccUpgradegroupBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradegroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "distribution_policy", "SIMULTANEOUSLY"),
					resource.TestCheckResourceAttr(resourceName, "time_zone", ""),
					resource.TestCheckResourceAttr(resourceName, "upgrade_policy", "SEQUENTIALLY"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccUpgradegroupResource_disappears(t *testing.T) {
	resourceName := "nios_grid_upgradegroup.test"
	var v grid.Upgradegroup

	name := acctest.RandomNameWithPrefix("example-upgradegroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckUpgradegroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccUpgradegroupBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradegroupExists(context.Background(), resourceName, &v),
					testAccCheckUpgradegroupDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccUpgradegroupResource_Comment(t *testing.T) {
	var resourceName = "nios_grid_upgradegroup.test_comment"
	var v grid.Upgradegroup

	name := acctest.RandomNameWithPrefix("example-upgradegroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccUpgradegroupComment(name, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradegroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccUpgradegroupComment(name, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradegroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccUpgradegroupResource_DistributionDependentGroup(t *testing.T) {
	var resourceName = "nios_grid_upgradegroup.test_distribution_dependent_group"
	var v grid.Upgradegroup

	name := acctest.RandomNameWithPrefix("example-upgradegroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccUpgradegroupDistributionDependentGroup(name, "example_distribution_dependent_group1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradegroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "distribution_dependent_group", "example_distribution_dependent_group1"),
				),
			},
			// Update and Read
			{
				Config: testAccUpgradegroupDistributionDependentGroup(name, "example_distribution_dependent_group2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradegroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "distribution_dependent_group", "example_distribution_dependent_group2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccUpgradegroupResource_DistributionPolicy(t *testing.T) {
	var resourceName = "nios_grid_upgradegroup.test_distribution_policy"
	var v grid.Upgradegroup

	name := acctest.RandomNameWithPrefix("example-upgradegroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccUpgradegroupDistributionPolicy(name, "SEQUENTIALLY"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradegroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "distribution_policy", "SEQUENTIALLY"),
				),
			},
			// Update and Read
			{
				Config: testAccUpgradegroupDistributionPolicy(name, "SIMULTANEOUSLY"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradegroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "distribution_policy", "SIMULTANEOUSLY"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccUpgradegroupResource_DistributionTime(t *testing.T) {
	var resourceName = "nios_grid_upgradegroup.test_distribution_time"
	var v grid.Upgradegroup

	name := acctest.RandomNameWithPrefix("example-upgradegroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccUpgradegroupDistributionTime(name, 1724965800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradegroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "distribution_time", "1724965800"),
				),
			},
			// Update and Read
			{
				Config: testAccUpgradegroupDistributionTime(name, 1724965800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradegroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "distribution_time", "1724965800"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccUpgradegroupResource_Members(t *testing.T) {
	var resourceName = "nios_grid_upgradegroup.test_members"
	var v grid.Upgradegroup

	name := acctest.RandomNameWithPrefix("example-upgradegroup-")
	member1 := []map[string]any{
		{"member": "infoblox.172_28_83_29"},
	}
	member2 := []map[string]any{
		{"member": "infoblox.172_28_82_115"},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccUpgradegroupMembers(name, member1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradegroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "members.0.member", "infoblox.172_28_83_29"),
				),
			},
			// // Update and Read
			{
				Config: testAccUpgradegroupMembers(name, member2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradegroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "members.0.member", "infoblox.172_28_82_115"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccUpgradegroupResource_Name(t *testing.T) {
	var resourceName = "nios_grid_upgradegroup.test_name"
	var v grid.Upgradegroup

	name1 := acctest.RandomNameWithPrefix("example-upgradegroup")
	name2 := acctest.RandomNameWithPrefix("updated-upgradegroup")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccUpgradegroupName(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradegroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccUpgradegroupName(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradegroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccUpgradegroupResource_UpgradeDependentGroup(t *testing.T) {
	var resourceName = "nios_grid_upgradegroup.test_upgrade_dependent_group"
	var v grid.Upgradegroup

	name := acctest.RandomNameWithPrefix("example-upgradegroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccUpgradegroupUpgradeDependentGroup(name, "example_upgrade_dependent_group1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradegroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "upgrade_dependent_group", "example_upgrade_dependent_group1"),
				),
			},
			// Update and Read
			{
				Config: testAccUpgradegroupUpgradeDependentGroup(name, "example_upgrade_dependent_group2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradegroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "upgrade_dependent_group", "example_upgrade_dependent_group2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccUpgradegroupResource_UpgradePolicy(t *testing.T) {
	var resourceName = "nios_grid_upgradegroup.test_upgrade_policy"
	var v grid.Upgradegroup

	name := acctest.RandomNameWithPrefix("example-upgradegroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccUpgradegroupUpgradePolicy(name, "SIMULTANEOUSLY"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradegroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "upgrade_policy", "SIMULTANEOUSLY"),
				),
			},
			// Update and Read
			{
				Config: testAccUpgradegroupUpgradePolicy(name, "SEQUENTIALLY"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradegroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "upgrade_policy", "SEQUENTIALLY"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccUpgradegroupResource_UpgradeTime(t *testing.T) {
	var resourceName = "nios_grid_upgradegroup.test_upgrade_time"
	var v grid.Upgradegroup

	name := acctest.RandomNameWithPrefix("example-upgradegroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccUpgradegroupUpgradeTime(name, 2724965800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradegroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "upgrade_time", "2724965800"),
				),
			},
			// Update and Read
			{
				Config: testAccUpgradegroupUpgradeTime(name, 2724965800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradegroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "upgrade_time", "2724965800"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckUpgradegroupExists(ctx context.Context, resourceName string, v *grid.Upgradegroup) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.GridAPI.
			UpgradegroupAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForUpgradegroup).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetUpgradegroupResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetUpgradegroupResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckUpgradegroupDestroy(ctx context.Context, v *grid.Upgradegroup) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.GridAPI.
			UpgradegroupAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForUpgradegroup).
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

func testAccCheckUpgradegroupDisappears(ctx context.Context, v *grid.Upgradegroup) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.GridAPI.
			UpgradegroupAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccUpgradegroupBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_grid_upgradegroup" "test" {
    name = %q
}
`, name)
}

func testAccUpgradegroupComment(name, comment string) string {
	return fmt.Sprintf(`
resource "nios_grid_upgradegroup" "test_comment" {
	name = %q	
    comment = %q
}
`, name, comment)
}

func testAccUpgradegroupDistributionDependentGroup(name, distributionDependentGroup string) string {
	return fmt.Sprintf(`
resource "nios_grid_upgradegroup" "test_distribution_dependent_group" {
	name = %q
    distribution_dependent_group = %q
}
`, name, distributionDependentGroup)
}

func testAccUpgradegroupDistributionPolicy(name, distributionPolicy string) string {
	return fmt.Sprintf(`
resource "nios_grid_upgradegroup" "test_distribution_policy" {
	name = %q
    distribution_policy = %q
}
`, name, distributionPolicy)
}

func testAccUpgradegroupDistributionTime(name string, distributionTime int64) string {
	return fmt.Sprintf(`
resource "nios_grid_upgradegroup" "test_distribution_time" {
	name = %q
    distribution_time = %d
}
`, name, distributionTime)
}

func testAccUpgradegroupMembers(name string, members []map[string]any) string {
	membersHCL := utils.ConvertSliceOfMapsToHCL(members)
	return fmt.Sprintf(`
resource "nios_grid_upgradegroup" "test_members" {
	name = %q
    members = %s
}
`, name, membersHCL)
}

func testAccUpgradegroupName(name string) string {
	return fmt.Sprintf(`
resource "nios_grid_upgradegroup" "test_name" {
    name = %q
}
`, name)
}

func testAccUpgradegroupUpgradeDependentGroup(name, upgradeDependentGroup string) string {
	return fmt.Sprintf(`
resource "nios_grid_upgradegroup" "test_upgrade_dependent_group" {
	name = %q
    upgrade_dependent_group = %q
}
`, name, upgradeDependentGroup)
}

func testAccUpgradegroupUpgradePolicy(name, upgradePolicy string) string {
	return fmt.Sprintf(`
resource "nios_grid_upgradegroup" "test_upgrade_policy" {
	name = %q
    upgrade_policy = %q
}
`, name, upgradePolicy)
}

func testAccUpgradegroupUpgradeTime(name string, upgradeTime int64) string {
	return fmt.Sprintf(`
resource "nios_grid_upgradegroup" "test_upgrade_time" {
	name = %q
    upgrade_time = %d
}
`, name, upgradeTime)
}
