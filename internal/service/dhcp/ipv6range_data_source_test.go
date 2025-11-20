package dhcp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

/*
// Retrieve a specific dhcp Ipv6range by filters
data "nios_dhcp_ipv6range" "get_dhcp_ipv6range_using_filters" {
  filters = {
    network = "NETWORK_REPLACE_ME"
  }
}
// Retrieve specific dhcp Ipv6range using Extensible Attributes
data "nios_" "get__using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all dhcp Ipv6range
data "nios_dhcp_ipv6range" "get_all_dhcp_ipv6range" {}
*/

func TestAccIpv6rangeDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_ipv6range.test"
	resourceName := "nios_dhcp_ipv6range.test"
	var v dhcp.Ipv6range

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6rangeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6rangeDataSourceConfigFilters("NETWORK_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					}, testAccCheckIpv6rangeResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccIpv6rangeDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_ipv6range.test"
	resourceName := "nios_dhcp_ipv6range.test"
	var v dhcp.Ipv6range
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6rangeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6rangeDataSourceConfigExtAttrFilters("NETWORK_REPLACE_ME", acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					}, testAccCheckIpv6rangeResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckIpv6rangeResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "address_type", dataSourceName, "result.0.address_type"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_info", dataSourceName, "result.0.cloud_info"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "discover_now_status", dataSourceName, "result.0.discover_now_status"),
		resource.TestCheckResourceAttrPair(resourceName, "discovery_basic_poll_settings", dataSourceName, "result.0.discovery_basic_poll_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "discovery_blackout_setting", dataSourceName, "result.0.discovery_blackout_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "discovery_member", dataSourceName, "result.0.discovery_member"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_discovery", dataSourceName, "result.0.enable_discovery"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_immediate_discovery", dataSourceName, "result.0.enable_immediate_discovery"),
		resource.TestCheckResourceAttrPair(resourceName, "end_addr", dataSourceName, "result.0.end_addr"),
		resource.TestCheckResourceAttrPair(resourceName, "endpoint_sources", dataSourceName, "result.0.endpoint_sources"),
		resource.TestCheckResourceAttrPair(resourceName, "exclude", dataSourceName, "result.0.exclude"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6_end_prefix", dataSourceName, "result.0.ipv6_end_prefix"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6_prefix_bits", dataSourceName, "result.0.ipv6_prefix_bits"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6_start_prefix", dataSourceName, "result.0.ipv6_start_prefix"),
		resource.TestCheckResourceAttrPair(resourceName, "logic_filter_rules", dataSourceName, "result.0.logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "member", dataSourceName, "result.0.member"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "network", dataSourceName, "result.0.network"),
		resource.TestCheckResourceAttrPair(resourceName, "network_view", dataSourceName, "result.0.network_view"),
		resource.TestCheckResourceAttrPair(resourceName, "option_filter_rules", dataSourceName, "result.0.option_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "port_control_blackout_setting", dataSourceName, "result.0.port_control_blackout_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "recycle_leases", dataSourceName, "result.0.recycle_leases"),
		resource.TestCheckResourceAttrPair(resourceName, "restart_if_needed", dataSourceName, "result.0.restart_if_needed"),
		resource.TestCheckResourceAttrPair(resourceName, "same_port_control_discovery_blackout", dataSourceName, "result.0.same_port_control_discovery_blackout"),
		resource.TestCheckResourceAttrPair(resourceName, "server_association_type", dataSourceName, "result.0.server_association_type"),
		resource.TestCheckResourceAttrPair(resourceName, "start_addr", dataSourceName, "result.0.start_addr"),
		resource.TestCheckResourceAttrPair(resourceName, "subscribe_settings", dataSourceName, "result.0.subscribe_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "template", dataSourceName, "result.0.template"),
		resource.TestCheckResourceAttrPair(resourceName, "use_blackout_setting", dataSourceName, "result.0.use_blackout_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "use_discovery_basic_polling_settings", dataSourceName, "result.0.use_discovery_basic_polling_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_discovery", dataSourceName, "result.0.use_enable_discovery"),
		resource.TestCheckResourceAttrPair(resourceName, "use_logic_filter_rules", dataSourceName, "result.0.use_logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "use_recycle_leases", dataSourceName, "result.0.use_recycle_leases"),
		resource.TestCheckResourceAttrPair(resourceName, "use_subscribe_settings", dataSourceName, "result.0.use_subscribe_settings"),
	}
}

func testAccIpv6rangeDataSourceConfigFilters(network string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test" {
  network = %q
}

data "nios_dhcp_ipv6range" "test" {
  filters = {
	network = nios_dhcp_ipv6range.test.network
  }
}
`, network)
}

func testAccIpv6rangeDataSourceConfigExtAttrFilters(network, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test" {
  network = %q
  extattrs = {
    Site = %q
  } 
}

data "nios_dhcp_ipv6range" "test" {
  extattrfilters = {
	Site = nios_dhcp_ipv6range.test.extattrs.Site
  }
}
`, network, extAttrsValue)
}
