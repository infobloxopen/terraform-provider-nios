package grid_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccGridServicerestartGroupDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_grid_servicerestart_group.test"
	resourceName := "nios_grid_servicerestart_group.test"
	var v grid.GridServicerestartGroup
	name := acctest.RandomNameWithPrefix("grid-service")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGridServicerestartGroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccGridServicerestartGroupDataSourceConfigFilters(name, "DNS"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckGridServicerestartGroupExists(context.Background(), resourceName, &v),
					}, testAccCheckGridServicerestartGroupResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccGridServicerestartGroupDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_grid_servicerestart_group.test"
	resourceName := "nios_grid_servicerestart_group.test"
	var v grid.GridServicerestartGroup
	name := acctest.RandomNameWithPrefix("grid-service")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGridServicerestartGroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccGridServicerestartGroupDataSourceConfigExtAttrFilters(name, "DNS", acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckGridServicerestartGroupExists(context.Background(), resourceName, &v),
					}, testAccCheckGridServicerestartGroupResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckGridServicerestartGroupResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "is_default", dataSourceName, "result.0.is_default"),
		resource.TestCheckResourceAttrPair(resourceName, "last_updated_time", dataSourceName, "result.0.last_updated_time"),
		resource.TestCheckResourceAttrPair(resourceName, "members", dataSourceName, "result.0.members"),
		resource.TestCheckResourceAttrPair(resourceName, "mode", dataSourceName, "result.0.mode"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "position", dataSourceName, "result.0.position"),
		resource.TestCheckResourceAttrPair(resourceName, "recurring_schedule", dataSourceName, "result.0.recurring_schedule"),
		resource.TestCheckResourceAttrPair(resourceName, "requests", dataSourceName, "result.0.requests"),
		resource.TestCheckResourceAttrPair(resourceName, "service", dataSourceName, "result.0.service"),
		resource.TestCheckResourceAttrPair(resourceName, "status", dataSourceName, "result.0.status"),
	}
}

func testAccGridServicerestartGroupDataSourceConfigFilters(name, service string) string {
	return fmt.Sprintf(`
resource "nios_grid_servicerestart_group" "test" {
	name = %q
	service = %q
}

data "nios_grid_servicerestart_group" "test" {
	filters = {
		name = nios_grid_servicerestart_group.test.name
	}
}
`, name, service)
}

func testAccGridServicerestartGroupDataSourceConfigExtAttrFilters(name, service, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_grid_servicerestart_group" "test" {
	name = %q
	service = %q
	extattrs = {
		Site = %q
	} 
}

data "nios_grid_servicerestart_group" "test" {
  extattrfilters = {
	Site = nios_grid_servicerestart_group.test.extattrs.Site
  }
}
`, name, service, extAttrsValue)
}
