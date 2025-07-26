package dhcp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccSharednetworkDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_sharednetwork.test"
	resourceName := "nios_dhcp_sharednetwork.test"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMjEuMjEuNi4wLzI0LzA:21.21.6.0/24/default",
		"network/ZG5zLm5ldHdvcmskMjEuMjEuNy4wLzI0LzA:21.21.7.0/24/default"}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharednetworkDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharednetworkDataSourceConfigFilters(name, networks),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					}, testAccCheckSharednetworkResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccSharednetworkDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_sharednetwork.test"
	resourceName := "nios_dhcp_sharednetwork.test"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMjEuMjEuNi4wLzI0LzA:21.21.6.0/24/default",
		"network/ZG5zLm5ldHdvcmskMjEuMjEuNy4wLzI0LzA:21.21.7.0/24/default"}
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharednetworkDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharednetworkDataSourceConfigExtAttrFilters(name, networks, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					}, testAccCheckSharednetworkResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckSharednetworkResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "authority", dataSourceName, "result.0.authority"),
		resource.TestCheckResourceAttrPair(resourceName, "bootfile", dataSourceName, "result.0.bootfile"),
		resource.TestCheckResourceAttrPair(resourceName, "bootserver", dataSourceName, "result.0.bootserver"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_generate_hostname", dataSourceName, "result.0.ddns_generate_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_server_always_updates", dataSourceName, "result.0.ddns_server_always_updates"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_ttl", dataSourceName, "result.0.ddns_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_update_fixed_addresses", dataSourceName, "result.0.ddns_update_fixed_addresses"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_use_option81", dataSourceName, "result.0.ddns_use_option81"),
		resource.TestCheckResourceAttrPair(resourceName, "deny_bootp", dataSourceName, "result.0.deny_bootp"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_utilization", dataSourceName, "result.0.dhcp_utilization"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_utilization_status", dataSourceName, "result.0.dhcp_utilization_status"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "dynamic_hosts", dataSourceName, "result.0.dynamic_hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_ddns", dataSourceName, "result.0.enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_pxe_lease_time", dataSourceName, "result.0.enable_pxe_lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "ignore_client_identifier", dataSourceName, "result.0.ignore_client_identifier"),
		resource.TestCheckResourceAttrPair(resourceName, "ignore_dhcp_option_list_request", dataSourceName, "result.0.ignore_dhcp_option_list_request"),
		resource.TestCheckResourceAttrPair(resourceName, "ignore_id", dataSourceName, "result.0.ignore_id"),
		resource.TestCheckResourceAttrPair(resourceName, "ignore_mac_addresses", dataSourceName, "result.0.ignore_mac_addresses"),
		resource.TestCheckResourceAttrPair(resourceName, "lease_scavenge_time", dataSourceName, "result.0.lease_scavenge_time"),
		resource.TestCheckResourceAttrPair(resourceName, "logic_filter_rules", dataSourceName, "result.0.logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_ad_user_data", dataSourceName, "result.0.ms_ad_user_data"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "network_view", dataSourceName, "result.0.network_view"),
		resource.TestCheckResourceAttrPair(resourceName, "networks", dataSourceName, "result.0.networks"),
		resource.TestCheckResourceAttrPair(resourceName, "nextserver", dataSourceName, "result.0.nextserver"),
		resource.TestCheckResourceAttrPair(resourceName, "options", dataSourceName, "result.0.options"),
		resource.TestCheckResourceAttrPair(resourceName, "pxe_lease_time", dataSourceName, "result.0.pxe_lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "static_hosts", dataSourceName, "result.0.static_hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "total_hosts", dataSourceName, "result.0.total_hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "update_dns_on_lease_renewal", dataSourceName, "result.0.update_dns_on_lease_renewal"),
		resource.TestCheckResourceAttrPair(resourceName, "use_authority", dataSourceName, "result.0.use_authority"),
		resource.TestCheckResourceAttrPair(resourceName, "use_bootfile", dataSourceName, "result.0.use_bootfile"),
		resource.TestCheckResourceAttrPair(resourceName, "use_bootserver", dataSourceName, "result.0.use_bootserver"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_generate_hostname", dataSourceName, "result.0.use_ddns_generate_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_ttl", dataSourceName, "result.0.use_ddns_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_update_fixed_addresses", dataSourceName, "result.0.use_ddns_update_fixed_addresses"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_use_option81", dataSourceName, "result.0.use_ddns_use_option81"),
		resource.TestCheckResourceAttrPair(resourceName, "use_deny_bootp", dataSourceName, "result.0.use_deny_bootp"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_ddns", dataSourceName, "result.0.use_enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ignore_client_identifier", dataSourceName, "result.0.use_ignore_client_identifier"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ignore_dhcp_option_list_request", dataSourceName, "result.0.use_ignore_dhcp_option_list_request"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ignore_id", dataSourceName, "result.0.use_ignore_id"),
		resource.TestCheckResourceAttrPair(resourceName, "use_lease_scavenge_time", dataSourceName, "result.0.use_lease_scavenge_time"),
		resource.TestCheckResourceAttrPair(resourceName, "use_logic_filter_rules", dataSourceName, "result.0.use_logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "use_nextserver", dataSourceName, "result.0.use_nextserver"),
		resource.TestCheckResourceAttrPair(resourceName, "use_options", dataSourceName, "result.0.use_options"),
		resource.TestCheckResourceAttrPair(resourceName, "use_pxe_lease_time", dataSourceName, "result.0.use_pxe_lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "use_update_dns_on_lease_renewal", dataSourceName, "result.0.use_update_dns_on_lease_renewal"),
	}
}

func testAccSharednetworkDataSourceConfigFilters(name string, networks []string) string {
	networksStr := formatNetworksToHCL(networks)
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test" {
 name = "%s"
 networks = %s
}

data "nios_dhcp_sharednetwork" "test" {
 filters = {
	 name = nios_dhcp_sharednetwork.test.name
 }
}
`, name, networksStr)
}

func testAccSharednetworkDataSourceConfigExtAttrFilters(name string, netwrorks []string, extAttrsValue string) string {
	netwrorksStr := formatNetworksToHCL(netwrorks)
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test" {
 name = %q
 networks = %s
 extattrs = {
   Site = %q
 }
}

data "nios_dhcp_sharednetwork" "test" {
 extattrfilters = {
	Site = nios_dhcp_sharednetwork.test.extattrs.Site
 }
}
`, name, netwrorksStr, extAttrsValue)
}
