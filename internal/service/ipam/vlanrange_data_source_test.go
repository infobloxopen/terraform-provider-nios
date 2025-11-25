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

func TestAccVlanrangeDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_ipam_vlanrange.test"
	resourceName := "nios_ipam_vlanrange.test"
	var v ipam.Vlanrange
	vlanRange := acctest.RandomNameWithPrefix("vlan-range")
	vlanView := acctest.RandomNameWithPrefix("vlan-view")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVlanrangeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccVlanrangeDataSourceConfigFilters(15, vlanRange, 10, vlanView),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					}, testAccCheckVlanrangeResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccVlanrangeDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_ipam_vlanrange.test"
	resourceName := "nios_ipam_vlanrange.test"
	var v ipam.Vlanrange
	vlanRange := acctest.RandomNameWithPrefix("vlan-range")
	vlanView := acctest.RandomNameWithPrefix("vlan-view")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVlanrangeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccVlanrangeDataSourceConfigExtAttrFilters(15, vlanRange, 10, vlanView, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					}, testAccCheckVlanrangeResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckVlanrangeResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "delete_vlans", dataSourceName, "result.0.delete_vlans"),
		resource.TestCheckResourceAttrPair(resourceName, "end_vlan_id", dataSourceName, "result.0.end_vlan_id"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "pre_create_vlan", dataSourceName, "result.0.pre_create_vlan"),
		resource.TestCheckResourceAttrPair(resourceName, "start_vlan_id", dataSourceName, "result.0.start_vlan_id"),
		resource.TestCheckResourceAttrPair(resourceName, "vlan_name_prefix", dataSourceName, "result.0.vlan_name_prefix"),
		resource.TestCheckResourceAttrPair(resourceName, "vlan_view", dataSourceName, "result.0.vlan_view"),
	}
}

func testAccVlanrangeDataSourceConfigFilters(endVlanId int, name string, startVlanId int, vlanView string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_vlanrange" "test" {
  end_vlan_id = %d
  name = %q
  start_vlan_id = %d
  vlan_view = nios_ipam_vlanview.test.ref
}

data "nios_ipam_vlanrange" "test" {
  filters = {
	end_vlan_id = nios_ipam_vlanrange.test.end_vlan_id
  }
}
`, endVlanId, name, startVlanId)
	return strings.Join([]string{testAccBaseWithVlanView(acctest.RandomNameWithPrefix("vlan-view")), config}, "")
}

func testAccVlanrangeDataSourceConfigExtAttrFilters(endVlanId int, name string, startVlanId int, vlanView, extAttrsValue string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_vlanrange" "test" {
  end_vlan_id = %d
  name = %q
  start_vlan_id = %d
  vlan_view = nios_ipam_vlanview.test.ref
  extattrs = {
    Site = %q
  } 
}

data "nios_ipam_vlanrange" "test" {
  extattrfilters = {
	Site = nios_ipam_vlanrange.test.extattrs.Site
  }
}
`, endVlanId, name, startVlanId, extAttrsValue)
	return strings.Join([]string{testAccBaseWithVlanView(acctest.RandomNameWithPrefix("vlan-view")), config}, "")
}
