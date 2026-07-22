package grid_test

import (
	"context"
	"fmt"
	"strconv"
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
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_grid_upgradeschedule.test"
	var v grid.Upgradeschedule
	start_time := time.Now().Add(12 * time.Hour).Format(utils.NaiveDatetimeLayout)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccUpgradescheduleBasicConfig(false, start_time),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradescheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_time", start_time),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "active", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccUpgradescheduleResource_Active(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
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
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_grid_upgradeschedule.test_start_time"
	var v grid.Upgradeschedule
	now := time.Now()
	start_time := now.Add(6 * time.Hour).Format(utils.NaiveDatetimeLayout)
	updated_start_time := now.Add(10 * time.Hour).Format(utils.NaiveDatetimeLayout)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccUpgradescheduleStartTime(start_time),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradescheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_time", start_time),
				),
			},
			// Update and Read
			{
				Config: testAccUpgradescheduleStartTime(updated_start_time),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradescheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_time", updated_start_time),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccUpgradescheduleResource_UpgradeGroups(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_grid_upgradeschedule.test_upgrade_groups"
	var v grid.Upgradeschedule

	groupName := acctest.RandomNameWithPrefix("example-upgradegroup-")

	now := time.Now()

	startTime := now.Add(12 * time.Hour).Format(utils.NaiveDatetimeLayout)

	upgrade_time := now.Add(24 * time.Hour).Format(utils.NaiveDatetimeLayout)

	upgrade_groups := []map[string]any{
		{
			"upgrade_time": upgrade_time,
			"name":         "Default",
		},
		{
			"upgrade_time": upgrade_time,
			"name":         groupName,
		},
	}

	updated_upgrade_time := now.Add(48 * time.Hour).Format(utils.NaiveDatetimeLayout)

	updated_upgrade_groups := []map[string]any{
		{
			"upgrade_time": updated_upgrade_time,
			"name":         "Default",
		},
		{
			"upgrade_time": updated_upgrade_time,
			"name":         groupName,
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccUpgradescheduleUpgradeGroups(groupName, startTime, upgrade_groups),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradescheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrSet(resourceName, "upgrade_groups.#"),
					testAccCheckUpgradeGroups(resourceName, "upgrade_time", upgrade_groups),
				),
			},
			// Update and Read
			{
				Config: testAccUpgradescheduleUpgradeGroups(groupName, startTime, updated_upgrade_groups),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUpgradescheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrSet(resourceName, "upgrade_groups.#"),
					testAccCheckUpgradeGroups(resourceName, "upgrade_time", updated_upgrade_groups),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// testAccCheckUpgradeGroups verifies that every entry in expectedGroups
// is present in the resource's upgrade_groups list (by name) and has the correct
// value for timeField (e.g. "upgrade_time" or "distribution_time").
func testAccCheckUpgradeGroups(resourceName, timeField string, expectedGroups []map[string]any) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		countStr := rs.Primary.Attributes["upgrade_groups.#"]
		count, err := strconv.Atoi(countStr)
		if err != nil || count == 0 {
			return fmt.Errorf("upgrade_groups.# is %q, expected a positive integer", countStr)
		}

		// Build a map of name -> timeField value from state (order-independent).
		actual := make(map[string]string, count)
		for i := range count {
			prefix := fmt.Sprintf("upgrade_groups.%d.", i)
			name := rs.Primary.Attributes[prefix+"name"]
			if name != "" {
				actual[name] = rs.Primary.Attributes[prefix+timeField]
			}
		}

		for _, eg := range expectedGroups {
			name, _ := eg["name"].(string)
			wantTime, _ := eg[timeField].(string)
			gotTime, found := actual[name]
			if !found {
				return fmt.Errorf("upgrade group %q not found in state (state has %d groups)", name, count)
			}
			if gotTime != wantTime {
				return fmt.Errorf("upgrade group %q: expected %s=%q, got %q", name, timeField, wantTime, gotTime)
			}
		}
		return nil
	}
}

func testAccCheckUpgradescheduleExists(ctx context.Context, resourceName string, v *grid.Upgradeschedule) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		uuid := rs.Primary.Attributes["uuid"]
		apiRes, _, err := acctest.NIOSClient.GridAPI.
			UpgradescheduleAPI.
			Read(ctx, utils.ResolveObjectIdentifier(&uuid, rs.Primary.Attributes["ref"])).
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
