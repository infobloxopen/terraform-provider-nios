package ipam_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccBulkhostnametemplateDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_ipam_bulkhostnametemplate.test"
	resourceName := "nios_ipam_bulkhostnametemplate.test"
	var v ipam.Bulkhostnametemplate

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBulkhostnametemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccBulkhostnametemplateDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckBulkhostnametemplateExists(context.Background(), resourceName, &v),
					}, testAccCheckBulkhostnametemplateResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccBulkhostnametemplateDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_ipam_bulkhostnametemplate.test"
	resourceName := "nios_ipam_bulkhostnametemplate.test"
	var v ipam.Bulkhostnametemplate
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBulkhostnametemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccBulkhostnametemplateDataSourceConfigExtAttrFilters(acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckBulkhostnametemplateExists(context.Background(), resourceName, &v),
					}, testAccCheckBulkhostnametemplateResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckBulkhostnametemplateResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "is_grid_default", dataSourceName, "result.0.is_grid_default"),
		resource.TestCheckResourceAttrPair(resourceName, "pre_defined", dataSourceName, "result.0.pre_defined"),
		resource.TestCheckResourceAttrPair(resourceName, "template_format", dataSourceName, "result.0.template_format"),
		resource.TestCheckResourceAttrPair(resourceName, "template_name", dataSourceName, "result.0.template_name"),
	}
}

func testAccBulkhostnametemplateDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_ipam_bulkhostnametemplate" "test" {
}

data "nios_ipam_bulkhostnametemplate" "test" {
  filters = {
	 = nios_ipam_bulkhostnametemplate.test.
  }
}
`)
}

func testAccBulkhostnametemplateDataSourceConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_ipam_bulkhostnametemplate" "test" {
  extattrs = {
    Site = %q
  } 
}

data "nios_ipam_bulkhostnametemplate" "test" {
  extattrfilters = {
	Site = nios_ipam_bulkhostnametemplate.test.extattrs.Site
  }
}
`, extAttrsValue)
}
