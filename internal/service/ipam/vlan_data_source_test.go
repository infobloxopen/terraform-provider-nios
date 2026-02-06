package ipam_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccVlanDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_ipam_vlan.test"
	resourceName := "nios_ipam_vlan.test"
	var v ipam.Vlan
	name := acctest.RandomNameWithPrefix("vlan")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVlanDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccVlanDataSourceConfigFilters(99, name, "example_vlan_view13"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckVlanExists(context.Background(), resourceName, &v),
					}, testAccCheckVlanResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccVlanDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_ipam_vlan.test"
	resourceName := "nios_ipam_vlan.test"
	var v ipam.Vlan
	name := acctest.RandomNameWithPrefix("vlan")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVlanDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccVlanDataSourceConfigExtAttrFilters(100, name, "example_vlan_view14", acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckVlanExists(context.Background(), resourceName, &v),
					}, testAccCheckVlanResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckVlanResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "assigned_to", dataSourceName, "result.0.assigned_to"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "contact", dataSourceName, "result.0.contact"),
		resource.TestCheckResourceAttrPair(resourceName, "department", dataSourceName, "result.0.department"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "result.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "result.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "parent", dataSourceName, "result.0.parent"),
		resource.TestCheckResourceAttrPair(resourceName, "reserved", dataSourceName, "result.0.reserved"),
		resource.TestCheckResourceAttrPair(resourceName, "status", dataSourceName, "result.0.status"),
	}
}

func testAccVlanDataSourceConfigFilters(id int, name, parent string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_vlan" "test" {
  id = %d
  name = %q
  parent = nios_ipam_vlanview.%s.ref
}

data "nios_ipam_vlan" "test" {
  filters = {
	id = nios_ipam_vlan.test.id
  }
}
`, id, name, parent)
	return strings.Join([]string{testAccBaseWithVlanView(parent), config}, "")
}

func testAccVlanDataSourceConfigExtAttrFilters(id int, name, parent, extAttrsValue string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_vlan" "test" {
  id = %d
  name = %q
  parent = nios_ipam_vlanview.%s.ref
  extattrs = {
    Site = %q
  } 
}

data "nios_ipam_vlan" "test" {
  extattrfilters = {
	Site = nios_ipam_vlan.test.extattrs.Site
  }
}
`, id, name, parent, extAttrsValue)
	return strings.Join([]string{testAccBaseWithVlanView(parent), config}, "")
}
