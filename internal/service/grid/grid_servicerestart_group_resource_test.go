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

// TODO: Objects required to be set in the grid
// - Members - infoblox.member, infoblox.localdomain

var readableAttributesForGridServicerestartGroup = "comment,extattrs,is_default,last_updated_time,members,mode,name,position,recurring_schedule,requests,service,status"

func TestAccGridServicerestartGroupResource_basic(t *testing.T) {
	var resourceName = "nios_grid_servicerestart_group.test"
	var v grid.GridServicerestartGroup
	name := acctest.RandomNameWithPrefix("grid-service")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGridServicerestartGroupBasicConfig(name, "DNS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "service", "DNS"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "mode", "SIMULTANEOUS"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccGridServicerestartGroupResource_disappears(t *testing.T) {
	resourceName := "nios_grid_servicerestart_group.test"
	var v grid.GridServicerestartGroup
	name := acctest.RandomNameWithPrefix("grid-service")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGridServicerestartGroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccGridServicerestartGroupBasicConfig(name, "DNS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupExists(context.Background(), resourceName, &v),
					testAccCheckGridServicerestartGroupDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccGridServicerestartGroupResource_Comment(t *testing.T) {
	var resourceName = "nios_grid_servicerestart_group.test_comment"
	var v grid.GridServicerestartGroup
	name := acctest.RandomNameWithPrefix("grid-service")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGridServicerestartGroupComment(name, "DNS", "Sample Comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Sample Comment"),
				),
			},
			// Update and Read
			{
				Config: testAccGridServicerestartGroupComment(name, "DNS", "Updated Comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated Comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccGridServicerestartGroupResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_grid_servicerestart_group.test_extattrs"
	var v grid.GridServicerestartGroup
	name := acctest.RandomNameWithPrefix("grid-service")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGridServicerestartGroupExtAttrs(name, "DNS", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccGridServicerestartGroupExtAttrs(name, "DNS", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccGridServicerestartGroupResource_Members(t *testing.T) {
	t.Skip("Requires members to be present in the grid to test against")
	var resourceName = "nios_grid_servicerestart_group.test_members"
	var v grid.GridServicerestartGroup
	name := acctest.RandomNameWithPrefix("grid-service")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGridServicerestartGroupMembers(name, "DHCP", "infoblox.member"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "members.0", "infoblox.member"),
				),
			},
			// Update and Read
			{
				Config: testAccGridServicerestartGroupMembers(name, "DNS", "infoblox.localdomain"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "members.0", "infoblox.localdomain"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccGridServicerestartGroupResource_Mode(t *testing.T) {
	var resourceName = "nios_grid_servicerestart_group.test_mode"
	var v grid.GridServicerestartGroup
	name := acctest.RandomNameWithPrefix("grid-service")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGridServicerestartGroupMode(name, "DNS", "SEQUENTIAL"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mode", "SEQUENTIAL"),
				),
			},
			// Update and Read
			{
				Config: testAccGridServicerestartGroupMode(name, "DNS", "SIMULTANEOUS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mode", "SIMULTANEOUS"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccGridServicerestartGroupResource_Name(t *testing.T) {
	var resourceName = "nios_grid_servicerestart_group.test_name"
	var v grid.GridServicerestartGroup
	name1 := acctest.RandomNameWithPrefix("grid-service")
	name2 := acctest.RandomNameWithPrefix("grid-service")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGridServicerestartGroupName(name1, "DNS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccGridServicerestartGroupName(name2, "DNS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccGridServicerestartGroupResource_RecurringSchedule(t *testing.T) {
	var resourceName = "nios_grid_servicerestart_group.test_recurring_schedule"
	var v grid.GridServicerestartGroup
	name := acctest.RandomNameWithPrefix("grid-service")

	recurringSchedule := map[string]any{
		"services": []string{"DHCPV6", "DNS"},
		"mode":     "SIMULTANEOUS",
		"force":    false,
		"schedule": map[string]any{
			"weekdays":          []string{"TUESDAY", "WEDNESDAY", "MONDAY"},
			"frequency":         "WEEKLY",
			"every":             15,
			"minutes_past_hour": 6,
			"disable":           false,
			"repeat":            "RECUR",
			"hour_of_day":       20,
		},
	}
	recurringScheduleUpdated := map[string]any{
		"services": []string{"ALL"},
		"mode":     "SIMULTANEOUS",
		"force":    true,
		"schedule": map[string]any{
			"minutes_past_hour": 6,
			"repeat":            "ONCE",
			"day_of_month":      30,
			"month":             1,
			"year":              2026,
			"hour_of_day":       20,
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGridServicerestartGroupRecurringSchedule(name, "DNS", recurringSchedule),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.services.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.services.0", "DHCPV6"),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.services.1", "DNS"),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.mode", "SIMULTANEOUS"),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.force", "false"),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.schedule.weekdays.0", "TUESDAY"),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.schedule.weekdays.1", "WEDNESDAY"),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.schedule.weekdays.2", "MONDAY"),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.schedule.frequency", "WEEKLY"),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.schedule.every", "15"),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.schedule.minutes_past_hour", "6"),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.schedule.disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.schedule.repeat", "RECUR"),
				),
			},
			// Update and Read
			{
				Config: testAccGridServicerestartGroupRecurringSchedule(name, "DNS", recurringScheduleUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.services.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.services.0", "ALL"),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.mode", "SIMULTANEOUS"),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.force", "true"),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.schedule.minutes_past_hour", "6"),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.schedule.repeat", "ONCE"),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.schedule.day_of_month", "30"),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.schedule.month", "1"),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.schedule.year", "2026"),
					resource.TestCheckResourceAttr(resourceName, "recurring_schedule.schedule.hour_of_day", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccGridServicerestartGroupResource_Service(t *testing.T) {
	var resourceName = "nios_grid_servicerestart_group.test_service"
	var v grid.GridServicerestartGroup
	name := acctest.RandomNameWithPrefix("grid-service")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccGridServicerestartGroupService(name, "DNS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service", "DNS"),
				),
			},
			// Update and Read
			{
				Config: testAccGridServicerestartGroupService(name, "DHCP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGridServicerestartGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service", "DHCP"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckGridServicerestartGroupExists(ctx context.Context, resourceName string, v *grid.GridServicerestartGroup) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.GridAPI.
			GridServicerestartGroupAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForGridServicerestartGroup).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetGridServicerestartGroupResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetGridServicerestartGroupResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckGridServicerestartGroupDestroy(ctx context.Context, v *grid.GridServicerestartGroup) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.GridAPI.
			GridServicerestartGroupAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForGridServicerestartGroup).
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

func testAccCheckGridServicerestartGroupDisappears(ctx context.Context, v *grid.GridServicerestartGroup) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.GridAPI.
			GridServicerestartGroupAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccGridServicerestartGroupBasicConfig(name, service string) string {
	return fmt.Sprintf(`
resource "nios_grid_servicerestart_group" "test" {
	name = %q
	service = %q
}
`, name, service)
}

func testAccGridServicerestartGroupComment(name, service, comment string) string {
	return fmt.Sprintf(`
resource "nios_grid_servicerestart_group" "test_comment" {
	name = %q
	comment = %q
	service = %q
}
`, name, comment, service)
}

func testAccGridServicerestartGroupExtAttrs(name, service string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
		  %s = %q
		`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_grid_servicerestart_group" "test_extattrs" {
	name = %q
	service = %q
	extattrs = %s
}
`, name, service, extattrsStr)
}

func testAccGridServicerestartGroupMembers(name, service, members string) string {
	return fmt.Sprintf(`
resource "nios_grid_servicerestart_group" "test_members" {
	name = %q
	service = %q
	members = [%q]
}
`, name, service, members)
}

func testAccGridServicerestartGroupMode(name, service, mode string) string {
	return fmt.Sprintf(`
resource "nios_grid_servicerestart_group" "test_mode" {
	name = %q
	service = %q
	mode = %q
}
`, name, service, mode)
}

func testAccGridServicerestartGroupName(name, service string) string {
	return fmt.Sprintf(`
resource "nios_grid_servicerestart_group" "test_name" {
	name = %q
	service = %q
}
`, name, service)
}

func testAccGridServicerestartGroupRecurringSchedule(name, service string, recurringSchedule map[string]any) string {
	recurringScheduleMap := utils.ConvertMapToHCL(recurringSchedule)
	return fmt.Sprintf(`
resource "nios_grid_servicerestart_group" "test_recurring_schedule" {
	name = %q
	service = %q
	recurring_schedule = %s
}
`, name, service, recurringScheduleMap)
}

func testAccGridServicerestartGroupService(name, service string) string {
	return fmt.Sprintf(`
resource "nios_grid_servicerestart_group" "test_service" {
	name = %q
	service = %q
}
`, name, service)
}
