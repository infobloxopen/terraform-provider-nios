package grid_test

import (
	"context"
	"fmt"
	"strconv"
	"strings"
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
	start_time := time.Now().Add(12 * time.Hour).Unix()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDistributionscheduleBasicConfig(false, start_time),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDistributionscheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_time", strconv.FormatInt(start_time, 10)),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "active", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDistributionscheduleResource_disappears(t *testing.T) {
	t.Skip("Distributionschedule cannot be deleted from NIOS, skipping disappears test")

	// resourceName := "nios_grid_distributionschedule.test"
	// var v grid.Distributionschedule
	// start_time := time.Now().Add(24 * time.Hour).Unix()

	// resource.ParallelTest(t, resource.TestCase{
	// 	PreCheck:                 func() { acctest.PreCheck(t) },
	// 	ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
	// 	CheckDestroy:             testAccCheckDistributionscheduleDestroy(context.Background(), &v),
	// 	Steps: []resource.TestStep{
	// 		{
	// 			Config: testAccDistributionscheduleBasicConfig(false, start_time),
	// 			Check: resource.ComposeTestCheckFunc(
	// 				testAccCheckDistributionscheduleExists(context.Background(), resourceName, &v),
	// 				testAccCheckDistributionscheduleDisappears(context.Background(), &v),
	// 			),
	// 			ExpectNonEmptyPlan: true,
	// 		},
	// 	},
	// })
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
				Config: testAccDistributionscheduleActive(false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDistributionscheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "active", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccDistributionscheduleActive(true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDistributionscheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "active", "true"),
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
	start_time := strconv.FormatInt(now.Add(1*time.Hour).Unix(), 10)
	updated_start_time := strconv.FormatInt(now.Add(5*time.Hour).Unix(), 10)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDistributionscheduleStartTime(start_time),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDistributionscheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_time", start_time),
				),
			},
			// Update and Read
			{
				Config: testAccDistributionscheduleStartTime(updated_start_time),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDistributionscheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_time", updated_start_time),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDistributionscheduleResource_UpgradeGroups(t *testing.T) {
	var resourceName = "nios_grid_distributionschedule.test_upgrade_groups"
	var v grid.Distributionschedule

	now := time.Now()

	startTime := now.Add(12 * time.Hour).Unix()

	distribution_time := now.Add(24 * time.Hour).Unix()

	upgrade_groups := []map[string]any{
		{
			"distribution_time": distribution_time,
			"name":              "Default",
		},
	}

	updated_distribution_time := now.Add(48 * time.Hour).Unix()

	updated_upgrade_groups := []map[string]any{
		{
			"distribution_time": updated_distribution_time,
			"name":              "Default",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDistributionscheduleUpgradeGroups(startTime, upgrade_groups),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDistributionscheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.0.name", "Default"),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.0.distribution_time", strconv.FormatInt(distribution_time, 10)),
				),
			},
			// Update and Read
			{
				Config: testAccDistributionscheduleUpgradeGroups(startTime, updated_upgrade_groups),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDistributionscheduleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.0.name", "Default"),
					resource.TestCheckResourceAttr(resourceName, "upgrade_groups.0.distribution_time", strconv.FormatInt(updated_distribution_time, 10))),
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

// func testAccCheckDistributionscheduleDestroy(ctx context.Context, v *grid.Distributionschedule) resource.TestCheckFunc {
// // Verify the resource was destroyed
// return func(state *terraform.State) error {
// 	_, httpRes, err := acctest.NIOSClient.GridAPI.
// 		DistributionscheduleAPI.
// 		Read(ctx, utils.ExtractResourceRef(*v.Ref)).
// 		ReturnAsObject(1).
// 		ReturnFieldsPlus(readableAttributesForDistributionschedule).
// 		Execute()
// 	if err != nil {
// 		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
// 			// resource was deleted
// 			return nil
// 		}
// 		return err
// 	}
// 	return errors.New("expected to be deleted")
// }
// return nil
// }

// func testAccCheckDistributionscheduleDisappears(ctx context.Context, v *grid.Distributionschedule) resource.TestCheckFunc {
// 	// Delete the resource externally to verify disappears test
// 	return func(state *terraform.State) error {
// 		_, err := acctest.NIOSClient.GridAPI.
// 			DistributionscheduleAPI.
// 			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
// 			Execute()
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	}
// }

func testAccDistributionscheduleBasicConfig(active bool, start_time int64) string {
	return fmt.Sprintf(`
resource "nios_grid_distributionschedule" "test" {
	active = %t
	start_time = %d
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

func testAccDistributionscheduleUpgradeGroups(startTime int64, upgradeGroups []map[string]any) string {
	hclGroups := []string{}
	for _, g := range upgradeGroups {
		groupHCL := fmt.Sprintf(`
    {
      name = %q
      distribution_time = %d
    }`, g["name"], g["distribution_time"])
		hclGroups = append(hclGroups, groupHCL)
	}

	return fmt.Sprintf(`
resource "nios_grid_distributionschedule" "test_upgrade_groups" {
  start_time = %d
  upgrade_groups = [
    %s
  ]
}
`, startTime, strings.Join(hclGroups, ","))
}
