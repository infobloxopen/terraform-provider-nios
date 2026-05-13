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

var readableAttributesForUpgradeschedule = "active,start_time,time_zone,upgrade_groups"

func TestAccUpgradescheduleResource_basic(t *testing.T) {
	var resourceName = "nios_grid_upgradeschedule.test"
	var v grid.Upgradeschedule
	startTime := time.Now().Add(24 * time.Hour).Format(utils.NaiveDatetimeLayout)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccUpgradescheduleBasicConfig(false, startTime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradescheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_time", startTime),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "active", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccUpgradescheduleResource_Active(t *testing.T) {
	var resourceName = "nios_grid_upgradeschedule.test_active"
	var v grid.Upgradeschedule

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccUpgradescheduleActive(true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradescheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "active", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccUpgradescheduleActive(false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradescheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "active", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccUpgradescheduleResource_StartTime(t *testing.T) {
	var resourceName = "nios_grid_upgradeschedule.test_start_time"
	var v grid.Upgradeschedule
	now := time.Now()
	startTime := now.Add(24 * time.Hour).Format(utils.NaiveDatetimeLayout)
	updatedStartTime := now.Add(36 * time.Hour).Format(utils.NaiveDatetimeLayout)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccUpgradescheduleStartTime(startTime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradescheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_time", startTime),
				),
			},
			// Update and Read
			{
				Config: testAccUpgradescheduleStartTime(updatedStartTime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradescheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_time", updatedStartTime),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccUpgradescheduleResource_UpgradeGroups(t *testing.T) {
	var resourceName = "nios_grid_upgradeschedule.test_upgrade_groups"
	var v grid.Upgradeschedule

	groupName := acctest.RandomNameWithPrefix("example-upgradegroup-")

	now := time.Now()

	startTime := now.Add(24 * time.Hour).Format(utils.NaiveDatetimeLayout)

	upgradeTime := now.Add(36 * time.Hour).Format(utils.NaiveDatetimeLayout)

	upgradeGroups := []map[string]any{
		{
			"upgrade_time": upgradeTime,
			"name":         "Default",
		},
		{
			"upgrade_time": upgradeTime,
			"name":         groupName,
		},
	}

	updatedUpgradeTime := now.Add(48 * time.Hour).Format(utils.NaiveDatetimeLayout)

	updatedUpgradeGroups := []map[string]any{
		{
			"upgrade_time": updatedUpgradeTime,
			"name":         "Default",
		},
		{
			"upgrade_time": updatedUpgradeTime,
			"name":         groupName,
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccUpgradescheduleUpgradeGroups(groupName, startTime, upgradeGroups),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradescheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.0.name", "Default"),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.0.upgrade_time", upgradeTime),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.1.name", groupName),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.1.upgrade_time", upgradeTime),
				),
			},
			// Update and Read
			{
				Config: testAccUpgradescheduleUpgradeGroups(groupName, startTime, updatedUpgradeGroups),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradescheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.0.name", "Default"),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.0.upgrade_time", updatedUpgradeTime),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.1.name", groupName),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.1.upgrade_time", updatedUpgradeTime),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckUpgradescheduleExists(ctx context.Context, resourceName string, v *grid.Upgradeschedule) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.GridAPI.
			UpgradescheduleAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForUpgradeschedule).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetUpgradescheduleResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetUpgradescheduleResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccUpgradescheduleBasicConfig(active bool, start_time string) string {
	return fmt.Sprintf(`
resource "nios_grid_upgradeschedule" "test" {
    active = %t
    start_time = %q
}
`, active, start_time)
}

func testAccUpgradescheduleActive(active bool) string {
	return fmt.Sprintf(`
resource "nios_grid_upgradeschedule" "test_active" {
    active = %t
}
`, active)
}

func testAccUpgradescheduleStartTime(startTime string) string {
	return fmt.Sprintf(`
resource "nios_grid_upgradeschedule" "test_start_time" {
    start_time = %q
}
`, startTime)
}

func testAccUpgradescheduleUpgradeGroups(groupName, startTime string, upgradeGroups []map[string]any) string {
	upgradeGroupsHCL := utils.ConvertSliceOfMapsToHCL(upgradeGroups)

	return fmt.Sprintf(`
resource "nios_grid_upgradegroup" "test" {
    name = %q
}

resource "nios_grid_upgradeschedule" "test_upgrade_groups" {
  	start_time = %q
    upgrade_groups = %s
}
`, groupName, startTime, upgradeGroupsHCL)
}
