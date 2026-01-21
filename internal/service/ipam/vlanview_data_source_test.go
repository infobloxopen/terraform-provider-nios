package ipam_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccVlanviewDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_ipam_vlanview.test"
	resourceName := "nios_ipam_vlanview.test"
	var v ipam.Vlanview
	name := acctest.RandomNameWithPrefix("vlan_view")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVlanviewDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccVlanviewDataSourceConfigFilters(15, name, 10),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					}, testAccCheckVlanviewResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccVlanviewDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_ipam_vlanview.test"
	resourceName := "nios_ipam_vlanview.test"
	var v ipam.Vlanview
	name := acctest.RandomNameWithPrefix("vlan_view")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVlanviewDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccVlanviewDataSourceConfigExtAttrFilters(15, name, 10, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					}, testAccCheckVlanviewResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckVlanviewResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "allow_range_overlapping", dataSourceName, "result.0.allow_range_overlapping"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "end_vlan_id", dataSourceName, "result.0.end_vlan_id"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "pre_create_vlan", dataSourceName, "result.0.pre_create_vlan"),
		resource.TestCheckResourceAttrPair(resourceName, "start_vlan_id", dataSourceName, "result.0.start_vlan_id"),
		resource.TestCheckResourceAttrPair(resourceName, "vlan_name_prefix", dataSourceName, "result.0.vlan_name_prefix"),
	}
}

func testAccVlanviewDataSourceConfigFilters(endVlanId int, name string, startVlanId int) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "test" {
  end_vlan_id = %d
  name = %q
  start_vlan_id = %d
}

data "nios_ipam_vlanview" "test" {
  filters = {
    end_vlan_id = nios_ipam_vlanview.test.end_vlan_id
    name = nios_ipam_vlanview.test.name
  }
}
`, endVlanId, name, startVlanId)
}

func testAccVlanviewDataSourceConfigExtAttrFilters(endVlanId int, name string, startVlanId int, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "test" {
  end_vlan_id = %d
  name = %q
  start_vlan_id = %d
  extattrs = {
    Site = %q
  } 
}

data "nios_ipam_vlanview" "test" {
  extattrfilters = {
	Site = nios_ipam_vlanview.test.extattrs.Site
  }
}
`, endVlanId, name, startVlanId, extAttrsValue)
}
