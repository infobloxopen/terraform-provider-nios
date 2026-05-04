package grid_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForDistributionschedule = "active,start_time,time_zone,upgrade_groups"

func TestAccDistributionscheduleResource_basic(t *testing.T) {
	var resourceName = "nios_grid_distributionschedule.test"
	var v grid.Distributionschedule
	startTime := time.Now().Add(12 * time.Hour).Format(utils.NaiveDatetimeLayout)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDistributionscheduleBasicConfig(false, startTime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDistributionscheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_time", startTime),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "active", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDistributionscheduleResource_Active(t *testing.T) {
	var resourceName = "nios_grid_distributionschedule.test_active"
	var v grid.Distributionschedule

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDistributionscheduleActive(true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDistributionscheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "active", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDistributionscheduleActive(false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDistributionscheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "active", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDistributionscheduleResource_StartTime(t *testing.T) {
	var resourceName = "nios_grid_distributionschedule.test_start_time"
	var v grid.Distributionschedule
	now := time.Now()
	startTime := now.Add(6 * time.Hour).Format(utils.NaiveDatetimeLayout)
	updatedStartTime := now.Add(10 * time.Hour).Format(utils.NaiveDatetimeLayout)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDistributionscheduleStartTime(startTime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDistributionscheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_time", startTime),
				),
			},
			// Update and Read
			{
				Config: testAccDistributionscheduleStartTime(updatedStartTime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDistributionscheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_time", updatedStartTime),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDistributionscheduleResource_UpgradeGroups(t *testing.T) {
	var resourceName = "nios_grid_distributionschedule.test_upgrade_groups"
	var v grid.Distributionschedule

	groupName := acctest.RandomNameWithPrefix("example-upgradegroup-")

	now := time.Now()

	startTime := now.Add(12 * time.Hour).Format(utils.NaiveDatetimeLayout)

	distributionTime := now.Add(20 * time.Hour).Format(utils.NaiveDatetimeLayout)

	upgradeGroups := []map[string]any{
		{
			"distribution_time": distributionTime,
			"name":              "Default",
		},
		{
			"distribution_time": distributionTime,
			"name":              groupName,
		},
	}

	updatedDistributionTime := now.Add(23 * time.Hour).Format(utils.NaiveDatetimeLayout)

	updatedUpgradeGroups := []map[string]any{
		{
			"distribution_time": updatedDistributionTime,
			"name":              "Default",
		},
		{
			"distribution_time": updatedDistributionTime,
			"name":              groupName,
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Deactivate schedule for Integration Testing
			{
				Config: testAccDistributionscheduleDeactivate(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nios_grid_distributionschedule.deactivate_schedule", "active", "false"),
				),
			},
			// Create and Read
			{
				Config: testAccDistributionscheduleUpgradeGroups(groupName, startTime, upgradeGroups),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDistributionscheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.0.name", "Default"),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.0.distribution_time", distributionTime),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.1.name", groupName),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.1.distribution_time", distributionTime),
				),
			},
			// Update and Read
			{
				Config: testAccDistributionscheduleUpgradeGroups(groupName, startTime, updatedUpgradeGroups),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDistributionscheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.0.name", "Default"),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.0.distribution_time", updatedDistributionTime),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.1.name", groupName),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.1.distribution_time", updatedDistributionTime),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDistributionscheduleExists(ctx context.Context, resourceName string, v *grid.Distributionschedule) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.GridAPI.
			DistributionscheduleAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDistributionschedule).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDistributionscheduleResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDistributionscheduleResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccDistributionscheduleBasicConfig(active bool, start_time string) string {
	return fmt.Sprintf(`
resource "nios_grid_distributionschedule" "test" {
	active = %t
	start_time = %q
}
`, active, start_time)
}

func testAccDistributionscheduleActive(active bool) string {
	return fmt.Sprintf(`
resource "nios_grid_distributionschedule" "test_active" {
    active = %t
}
`, active)
}

func testAccDistributionscheduleStartTime(startTime string) string {
	return fmt.Sprintf(`
resource "nios_grid_distributionschedule" "test_start_time" {
    start_time = %q
}
`, startTime)
}

func testAccDistributionscheduleUpgradeGroups(groupName, startTime string, upgradeGroups []map[string]any) string {
	upgradeGroupsHCL := utils.ConvertSliceOfMapsToHCL(upgradeGroups)

	return fmt.Sprintf(`
resource "nios_grid_upgradegroup" "test" {
    name = %q
}

resource "nios_grid_distributionschedule" "test_upgrade_groups" {
  start_time = %q
  upgrade_groups = %s
}
`, groupName, startTime, upgradeGroupsHCL)
}

func testAccDistributionscheduleDeactivate() string {
	return `
resource "nios_grid_distributionschedule" "deactivate_schedule" {
    active = false
}
`
}
