package ipam_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccNetworkviewDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_ipam_network_view.test"
	resourceName := "nios_ipam_network_view.test"
	var v ipam.Networkview
	name := acctest.RandomNameWithPrefix("test-network-view")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetworkviewDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkviewDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					}, testAccCheckNetworkviewResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccNetworkviewDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_ipam_network_view.test"
	resourceName := "nios_ipam_network_view.test"
	var v ipam.Networkview
	name := acctest.RandomNameWithPrefix("test-network-view")
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetworkviewDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkviewDataSourceConfigExtAttrFilters(name, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					}, testAccCheckNetworkviewResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckNetworkviewResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "associated_dns_views", dataSourceName, "result.0.associated_dns_views"),
		resource.TestCheckResourceAttrPair(resourceName, "associated_members", dataSourceName, "result.0.associated_members"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_info", dataSourceName, "result.0.cloud_info"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_dns_view", dataSourceName, "result.0.ddns_dns_view"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_zone_primaries", dataSourceName, "result.0.ddns_zone_primaries"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "federated_realms", dataSourceName, "result.0.federated_realms"),
		resource.TestCheckResourceAttrPair(resourceName, "internal_forward_zones", dataSourceName, "result.0.internal_forward_zones"),
		resource.TestCheckResourceAttrPair(resourceName, "is_default", dataSourceName, "result.0.is_default"),
		resource.TestCheckResourceAttrPair(resourceName, "mgm_private", dataSourceName, "result.0.mgm_private"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_ad_user_data", dataSourceName, "result.0.ms_ad_user_data"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "remote_forward_zones", dataSourceName, "result.0.remote_forward_zones"),
		resource.TestCheckResourceAttrPair(resourceName, "remote_reverse_zones", dataSourceName, "result.0.remote_reverse_zones"),
	}
}

func testAccNetworkviewDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_view" "test" {
	name = %q
}

data "nios_ipam_network_view" "test" {
  filters = {
	name = nios_ipam_network_view.test.name
  }
}
`, name)
}

func testAccNetworkviewDataSourceConfigExtAttrFilters(name, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_view" "test" {
  name = %q
    extattrs = {
    	Site = %q
  } 
}

data "nios_ipam_network_view" "test" {
  extattrfilters = {
	Site = nios_ipam_network_view.test.extattrs.Site
  }
}
`, name, extAttrsValue)
}
