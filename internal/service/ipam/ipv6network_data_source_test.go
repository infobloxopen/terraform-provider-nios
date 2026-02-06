package ipam_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccIpv6networkDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_ipam_ipv6network.test"
	resourceName := "nios_ipam_ipv6network.test"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6networkDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6networkDataSourceConfigFilters(network),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					}, testAccCheckIpv6networkResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccIpv6networkDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_ipam_ipv6network.test"
	resourceName := "nios_ipam_ipv6network.test"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6networkDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6networkDataSourceConfigExtAttrFilters(network, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					}, testAccCheckIpv6networkResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckIpv6networkResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_info", dataSourceName, "result.0.cloud_info"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_domainname", dataSourceName, "result.0.ddns_domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_enable_option_fqdn", dataSourceName, "result.0.ddns_enable_option_fqdn"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_generate_hostname", dataSourceName, "result.0.ddns_generate_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_server_always_updates", dataSourceName, "result.0.ddns_server_always_updates"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_ttl", dataSourceName, "result.0.ddns_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "discover_now_status", dataSourceName, "result.0.discover_now_status"),
		resource.TestCheckResourceAttrPair(resourceName, "discovered_bgp_as", dataSourceName, "result.0.discovered_bgp_as"),
		resource.TestCheckResourceAttrPair(resourceName, "discovered_bridge_domain", dataSourceName, "result.0.discovered_bridge_domain"),
		resource.TestCheckResourceAttrPair(resourceName, "discovered_tenant", dataSourceName, "result.0.discovered_tenant"),
		resource.TestCheckResourceAttrPair(resourceName, "discovered_vlan_id", dataSourceName, "result.0.discovered_vlan_id"),
		resource.TestCheckResourceAttrPair(resourceName, "discovered_vlan_name", dataSourceName, "result.0.discovered_vlan_name"),
		resource.TestCheckResourceAttrPair(resourceName, "discovered_vrf_description", dataSourceName, "result.0.discovered_vrf_description"),
		resource.TestCheckResourceAttrPair(resourceName, "discovered_vrf_name", dataSourceName, "result.0.discovered_vrf_name"),
		resource.TestCheckResourceAttrPair(resourceName, "discovered_vrf_rd", dataSourceName, "result.0.discovered_vrf_rd"),
		resource.TestCheckResourceAttrPair(resourceName, "discovery_basic_poll_settings", dataSourceName, "result.0.discovery_basic_poll_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "discovery_blackout_setting", dataSourceName, "result.0.discovery_blackout_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "discovery_engine_type", dataSourceName, "result.0.discovery_engine_type"),
		resource.TestCheckResourceAttrPair(resourceName, "discovery_member", dataSourceName, "result.0.discovery_member"),
		resource.TestCheckResourceAttrPair(resourceName, "domain_name", dataSourceName, "result.0.domain_name"),
		resource.TestCheckResourceAttrPair(resourceName, "domain_name_servers", dataSourceName, "result.0.domain_name_servers"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_ddns", dataSourceName, "result.0.enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_discovery", dataSourceName, "result.0.enable_discovery"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_ifmap_publishing", dataSourceName, "result.0.enable_ifmap_publishing"),
		resource.TestCheckResourceAttrPair(resourceName, "endpoint_sources", dataSourceName, "result.0.endpoint_sources"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "federated_realms", dataSourceName, "result.0.federated_realms"),
		resource.TestCheckResourceAttrPair(resourceName, "last_rir_registration_update_sent", dataSourceName, "result.0.last_rir_registration_update_sent"),
		resource.TestCheckResourceAttrPair(resourceName, "last_rir_registration_update_status", dataSourceName, "result.0.last_rir_registration_update_status"),
		resource.TestCheckResourceAttrPair(resourceName, "logic_filter_rules", dataSourceName, "result.0.logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "members", dataSourceName, "result.0.members"),
		resource.TestCheckResourceAttrPair(resourceName, "mgm_private", dataSourceName, "result.0.mgm_private"),
		resource.TestCheckResourceAttrPair(resourceName, "mgm_private_overridable", dataSourceName, "result.0.mgm_private_overridable"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_ad_user_data", dataSourceName, "result.0.ms_ad_user_data"),
		resource.TestCheckResourceAttrPair(resourceName, "network", dataSourceName, "result.0.network"),
		resource.TestCheckResourceAttrPair(resourceName, "func_call", dataSourceName, "result.0.func_call"),
		resource.TestCheckResourceAttrPair(resourceName, "network_container", dataSourceName, "result.0.network_container"),
		resource.TestCheckResourceAttrPair(resourceName, "network_view", dataSourceName, "result.0.network_view"),
		resource.TestCheckResourceAttrPair(resourceName, "options", dataSourceName, "result.0.options"),
		resource.TestCheckResourceAttrPair(resourceName, "port_control_blackout_setting", dataSourceName, "result.0.port_control_blackout_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "preferred_lifetime", dataSourceName, "result.0.preferred_lifetime"),
		resource.TestCheckResourceAttrPair(resourceName, "recycle_leases", dataSourceName, "result.0.recycle_leases"),
		resource.TestCheckResourceAttrPair(resourceName, "rir", dataSourceName, "result.0.rir"),
		resource.TestCheckResourceAttrPair(resourceName, "rir_organization", dataSourceName, "result.0.rir_organization"),
		resource.TestCheckResourceAttrPair(resourceName, "rir_registration_status", dataSourceName, "result.0.rir_registration_status"),
		resource.TestCheckResourceAttrPair(resourceName, "same_port_control_discovery_blackout", dataSourceName, "result.0.same_port_control_discovery_blackout"),
		resource.TestCheckResourceAttrPair(resourceName, "subscribe_settings", dataSourceName, "result.0.subscribe_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "unmanaged", dataSourceName, "result.0.unmanaged"),
		resource.TestCheckResourceAttrPair(resourceName, "unmanaged_count", dataSourceName, "result.0.unmanaged_count"),
		resource.TestCheckResourceAttrPair(resourceName, "update_dns_on_lease_renewal", dataSourceName, "result.0.update_dns_on_lease_renewal"),
		resource.TestCheckResourceAttrPair(resourceName, "use_blackout_setting", dataSourceName, "result.0.use_blackout_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_domainname", dataSourceName, "result.0.use_ddns_domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_enable_option_fqdn", dataSourceName, "result.0.use_ddns_enable_option_fqdn"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_generate_hostname", dataSourceName, "result.0.use_ddns_generate_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_ttl", dataSourceName, "result.0.use_ddns_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_discovery_basic_polling_settings", dataSourceName, "result.0.use_discovery_basic_polling_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "use_domain_name", dataSourceName, "result.0.use_domain_name"),
		resource.TestCheckResourceAttrPair(resourceName, "use_domain_name_servers", dataSourceName, "result.0.use_domain_name_servers"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_ddns", dataSourceName, "result.0.use_enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_discovery", dataSourceName, "result.0.use_enable_discovery"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_ifmap_publishing", dataSourceName, "result.0.use_enable_ifmap_publishing"),
		resource.TestCheckResourceAttrPair(resourceName, "use_logic_filter_rules", dataSourceName, "result.0.use_logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "use_mgm_private", dataSourceName, "result.0.use_mgm_private"),
		resource.TestCheckResourceAttrPair(resourceName, "use_options", dataSourceName, "result.0.use_options"),
		resource.TestCheckResourceAttrPair(resourceName, "use_preferred_lifetime", dataSourceName, "result.0.use_preferred_lifetime"),
		resource.TestCheckResourceAttrPair(resourceName, "use_recycle_leases", dataSourceName, "result.0.use_recycle_leases"),
		resource.TestCheckResourceAttrPair(resourceName, "use_subscribe_settings", dataSourceName, "result.0.use_subscribe_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "use_update_dns_on_lease_renewal", dataSourceName, "result.0.use_update_dns_on_lease_renewal"),
		resource.TestCheckResourceAttrPair(resourceName, "use_valid_lifetime", dataSourceName, "result.0.use_valid_lifetime"),
		resource.TestCheckResourceAttrPair(resourceName, "use_zone_associations", dataSourceName, "result.0.use_zone_associations"),
		resource.TestCheckResourceAttrPair(resourceName, "valid_lifetime", dataSourceName, "result.0.valid_lifetime"),
		resource.TestCheckResourceAttrPair(resourceName, "vlans", dataSourceName, "result.0.vlans"),
		resource.TestCheckResourceAttrPair(resourceName, "zone_associations", dataSourceName, "result.0.zone_associations"),
	}
}

func testAccIpv6networkDataSourceConfigFilters(network string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test" {
  network = %q
}

data "nios_ipam_ipv6network" "test" {
  filters = {
	 network = nios_ipam_ipv6network.test.network
  }
}
`, network)
}

func testAccIpv6networkDataSourceConfigExtAttrFilters(network, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test" {
  network = %q
  extattrs = {
    Site = %q
  }
}

data "nios_ipam_ipv6network" "test" {
  extattrfilters = {
	Site = nios_ipam_ipv6network.test.extattrs.Site
  }
}
`, network, extAttrsValue)
}
