package dhcp_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccIpv6fixedaddressDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_ipv6fixedaddress.test"
	resourceName := "nios_dhcp_ipv6fixedaddress.test"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	macAddr := "00:0c:29:ab:cd:ef"
	matchClient := "MAC_ADDRESS"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6fixedaddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6fixedaddressDataSourceConfigFilters(ipv6addr, matchClient, macAddr, networkView, ipv6Network),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					}, testAccCheckIpv6fixedaddressResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccIpv6fixedaddressDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_ipv6fixedaddress.test"
	resourceName := "nios_dhcp_ipv6fixedaddress.test"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	matchClient := "MAC_ADDRESS"
	networkView := acctest.RandomNameWithPrefix("network-view")
	macAddr := "00:0c:29:ab:cd:ef"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6fixedaddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6fixedaddressDataSourceConfigExtAttrFilters(ipv6addr, matchClient, macAddr, acctest.RandomName(), networkView, ipv6Network),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					}, testAccCheckIpv6fixedaddressResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckIpv6fixedaddressResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "address_type", dataSourceName, "result.0.address_type"),
		resource.TestCheckResourceAttrPair(resourceName, "allow_telnet", dataSourceName, "result.0.allow_telnet"),
		resource.TestCheckResourceAttrPair(resourceName, "cli_credentials", dataSourceName, "result.0.cli_credentials"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_info", dataSourceName, "result.0.cloud_info"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "device_description", dataSourceName, "result.0.device_description"),
		resource.TestCheckResourceAttrPair(resourceName, "device_location", dataSourceName, "result.0.device_location"),
		resource.TestCheckResourceAttrPair(resourceName, "device_type", dataSourceName, "result.0.device_type"),
		resource.TestCheckResourceAttrPair(resourceName, "device_vendor", dataSourceName, "result.0.device_vendor"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "disable_discovery", dataSourceName, "result.0.disable_discovery"),
		resource.TestCheckResourceAttrPair(resourceName, "discover_now_status", dataSourceName, "result.0.discover_now_status"),
		resource.TestCheckResourceAttrPair(resourceName, "discovered_data", dataSourceName, "result.0.discovered_data"),
		resource.TestCheckResourceAttrPair(resourceName, "domain_name", dataSourceName, "result.0.domain_name"),
		resource.TestCheckResourceAttrPair(resourceName, "domain_name_servers", dataSourceName, "result.0.domain_name_servers"),
		resource.TestCheckResourceAttrPair(resourceName, "duid", dataSourceName, "result.0.duid"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_immediate_discovery", dataSourceName, "result.0.enable_immediate_discovery"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6addr", dataSourceName, "result.0.ipv6addr"),
		resource.TestCheckResourceAttrPair(resourceName, "func_call", dataSourceName, "result.0.func_call"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6prefix", dataSourceName, "result.0.ipv6prefix"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6prefix_bits", dataSourceName, "result.0.ipv6prefix_bits"),
		resource.TestCheckResourceAttrPair(resourceName, "logic_filter_rules", dataSourceName, "result.0.logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "mac_address", dataSourceName, "result.0.mac_address"),
		resource.TestCheckResourceAttrPair(resourceName, "match_client", dataSourceName, "result.0.match_client"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_ad_user_data", dataSourceName, "result.0.ms_ad_user_data"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "network", dataSourceName, "result.0.network"),
		resource.TestCheckResourceAttrPair(resourceName, "network_view", dataSourceName, "result.0.network_view"),
		resource.TestCheckResourceAttrPair(resourceName, "options", dataSourceName, "result.0.options"),
		resource.TestCheckResourceAttrPair(resourceName, "preferred_lifetime", dataSourceName, "result.0.preferred_lifetime"),
		resource.TestCheckResourceAttrPair(resourceName, "reserved_interface", dataSourceName, "result.0.reserved_interface"),
		resource.TestCheckResourceAttrPair(resourceName, "restart_if_needed", dataSourceName, "result.0.restart_if_needed"),
		resource.TestCheckResourceAttrPair(resourceName, "snmp3_credential", dataSourceName, "result.0.snmp3_credential"),
		resource.TestCheckResourceAttrPair(resourceName, "snmp_credential", dataSourceName, "result.0.snmp_credential"),
		resource.TestCheckResourceAttrPair(resourceName, "template", dataSourceName, "result.0.template"),
		resource.TestCheckResourceAttrPair(resourceName, "use_cli_credentials", dataSourceName, "result.0.use_cli_credentials"),
		resource.TestCheckResourceAttrPair(resourceName, "use_domain_name", dataSourceName, "result.0.use_domain_name"),
		resource.TestCheckResourceAttrPair(resourceName, "use_domain_name_servers", dataSourceName, "result.0.use_domain_name_servers"),
		resource.TestCheckResourceAttrPair(resourceName, "use_logic_filter_rules", dataSourceName, "result.0.use_logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "use_options", dataSourceName, "result.0.use_options"),
		resource.TestCheckResourceAttrPair(resourceName, "use_preferred_lifetime", dataSourceName, "result.0.use_preferred_lifetime"),
		resource.TestCheckResourceAttrPair(resourceName, "use_snmp3_credential", dataSourceName, "result.0.use_snmp3_credential"),
		resource.TestCheckResourceAttrPair(resourceName, "use_snmp_credential", dataSourceName, "result.0.use_snmp_credential"),
		resource.TestCheckResourceAttrPair(resourceName, "use_valid_lifetime", dataSourceName, "result.0.use_valid_lifetime"),
		resource.TestCheckResourceAttrPair(resourceName, "valid_lifetime", dataSourceName, "result.0.valid_lifetime"),
	}
}

func testAccIpv6fixedaddressDataSourceConfigFilters(ipv6addr, matchClient, macAddress, networkView, ipv6Network string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test" {
 ipv6addr = %q
 match_client = %q
 mac_address = %q
 network = nios_ipam_ipv6network.test_ipv6_network.network
 network_view = nios_ipam_network_view.parent_network_view.name
}

data "nios_dhcp_ipv6fixedaddress" "test" {
 filters = {
	 ipv6addr = nios_dhcp_ipv6fixedaddress.test.ipv6addr
 }
}
`, ipv6addr, matchClient, macAddress)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressDataSourceConfigExtAttrFilters(ipv6addr, matchClient, macAddress, extAttrsValue, networkView, ipv6Network string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test" {
 ipv6addr = %q
 match_client = %q
 mac_address = %q
 network = nios_ipam_ipv6network.test_ipv6_network.network
 network_view = nios_ipam_network_view.parent_network_view.name
 extattrs = {
   Site = %q
 }
}

data "nios_dhcp_ipv6fixedaddress" "test" {
 extattrfilters = {
	Site = nios_dhcp_ipv6fixedaddress.test.extattrs.Site
 }
}
`, ipv6addr, matchClient, macAddress, extAttrsValue)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccBaseNetworkView(networkView, ipv6Network string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_ipv6_network" {
  network      = %q
  network_view = nios_ipam_network_view.parent_network_view.name
}

resource "nios_ipam_network_view" "parent_network_view" {
  name = %q
}
`, ipv6Network, networkView)
}
