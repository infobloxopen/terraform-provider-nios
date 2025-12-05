package grid_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccMemberdfpDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_grid_memberdfp.test"
	resourceName := "nios_grid_memberdfp.test"
	var v grid.Memberdfp

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMemberdfpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMemberdfpDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckMemberdfpExists(context.Background(), resourceName, &v),
					}, testAccCheckMemberdfpResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccMemberdfpDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_grid_memberdfp.test"
	resourceName := "nios_grid_memberdfp.test"
	var v grid.Memberdfp
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMemberdfpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMemberdfpDataSourceConfigExtAttrFilters(acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckMemberdfpExists(context.Background(), resourceName, &v),
					}, testAccCheckMemberdfpResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckMemberdfpResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "dfp_forward_first", dataSourceName, "result.0.dfp_forward_first"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "host_name", dataSourceName, "result.0.host_name"),
		resource.TestCheckResourceAttrPair(resourceName, "is_dfp_override", dataSourceName, "result.0.is_dfp_override"),
	}
}

func testAccMemberdfpDataSourceConfigFilters() string {
	return `
resource "nios_grid_memberdfp" "test" {
}

data "nios_grid_memberdfp" "test" {
  filters = {
	 = nios_grid_memberdfp.test.
  }
}
`
}

func testAccMemberdfpDataSourceConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_grid_memberdfp" "test" {
  extattrs = {
    Site = %q
  } 
}

data "nios_grid_memberdfp" "test" {
  extattrfilters = {
	Site = nios_grid_memberdfp.test.extattrs.Site
  }
}
`, extAttrsValue)
}
