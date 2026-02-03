package ipam_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccSuperhostDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_ipam_superhost.test"
	resourceName := "nios_ipam_superhost.test"
	var v ipam.Superhost
	name := acctest.RandomNameWithPrefix("super-host")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSuperhostDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSuperhostDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					}, testAccCheckSuperhostResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccSuperhostDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_ipam_superhost.test"
	resourceName := "nios_ipam_superhost.test"
	var v ipam.Superhost
	name := acctest.RandomNameWithPrefix("super-host")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSuperhostDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSuperhostDataSourceConfigExtAttrFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					}, testAccCheckSuperhostResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckSuperhostResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_associated_objects", dataSourceName, "result.0.dhcp_associated_objects"),
		resource.TestCheckResourceAttrPair(resourceName, "disabled", dataSourceName, "result.0.disabled"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_associated_objects", dataSourceName, "result.0.dns_associated_objects"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
	}
}

func testAccSuperhostDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_ipam_superhost" "test" {
  name = %q
}

data "nios_ipam_superhost" "test" {
  filters = {
	name = nios_ipam_superhost.test.name
  }
}
`, name)
}

func testAccSuperhostDataSourceConfigExtAttrFilters(name, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_ipam_superhost" "test" {
  name = %q
  extattrs = {
    Site = %q
  } 
}

data "nios_ipam_superhost" "test" {
  extattrfilters = {
    Site = nios_ipam_superhost.test.extattrs.Site
  }
}
`, name, extAttrsValue)
}
