package ipam_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccNetworkDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_ipam_network.test"
	resourceName := "nios_ipam_network.test"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetworkDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkDataSourceConfigFilters(network),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckNetworkExists(context.Background(), resourceName, &v),
					}, testAccCheckNetworkResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccNetworkDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_ipam_network.test"
	resourceName := "nios_ipam_network.test"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetworkDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkDataSourceConfigExtAttrFilters(network, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckNetworkExists(context.Background(), resourceName, &v),
					}, testAccCheckNetworkResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckNetworkResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "authority", dataSourceName, "result.0.authority"),
		resource.TestCheckResourceAttrPair(resourceName, "auto_create_reversezone", dataSourceName, "result.0.auto_create_reversezone"),
		resource.TestCheckResourceAttrPair(resourceName, "bootfile", dataSourceName, "result.0.bootfile"),
		resource.TestCheckResourceAttrPair(resourceName, "bootserver", dataSourceName, "result.0.bootserver"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_info", dataSourceName, "result.0.cloud_info"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_shared", dataSourceName, "result.0.cloud_shared"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "conflict_count", dataSourceName, "result.0.conflict_count"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_domainname", dataSourceName, "result.0.ddns_domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_generate_hostname", dataSourceName, "result.0.ddns_generate_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_server_always_updates", dataSourceName, "result.0.ddns_server_always_updates"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_ttl", dataSourceName, "result.0.ddns_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_update_fixed_addresses", dataSourceName, "result.0.ddns_update_fixed_addresses"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_use_option81", dataSourceName, "result.0.ddns_use_option81"),
		resource.TestCheckResourceAttrPair(resourceName, "delete_reason", dataSourceName, "result.0.delete_reason"),
		resource.TestCheckResourceAttrPair(resourceName, "deny_bootp", dataSourceName, "result.0.deny_bootp"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_utilization", dataSourceName, "result.0.dhcp_utilization"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_utilization_status", dataSourceName, "result.0.dhcp_utilization_status"),
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
		resource.TestCheckResourceAttrPair(resourceName, "dynamic_hosts", dataSourceName, "result.0.dynamic_hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "email_list", dataSourceName, "result.0.email_list"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_ddns", dataSourceName, "result.0.enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_dhcp_thresholds", dataSourceName, "result.0.enable_dhcp_thresholds"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_discovery", dataSourceName, "result.0.enable_discovery"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_email_warnings", dataSourceName, "result.0.enable_email_warnings"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_ifmap_publishing", dataSourceName, "result.0.enable_ifmap_publishing"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_immediate_discovery", dataSourceName, "result.0.enable_immediate_discovery"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_pxe_lease_time", dataSourceName, "result.0.enable_pxe_lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_snmp_warnings", dataSourceName, "result.0.enable_snmp_warnings"),
		resource.TestCheckResourceAttrPair(resourceName, "endpoint_sources", dataSourceName, "result.0.endpoint_sources"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "federated_realms", dataSourceName, "result.0.federated_realms"),
		resource.TestCheckResourceAttrPair(resourceName, "high_water_mark", dataSourceName, "result.0.high_water_mark"),
		resource.TestCheckResourceAttrPair(resourceName, "high_water_mark_reset", dataSourceName, "result.0.high_water_mark_reset"),
		resource.TestCheckResourceAttrPair(resourceName, "ignore_dhcp_option_list_request", dataSourceName, "result.0.ignore_dhcp_option_list_request"),
		resource.TestCheckResourceAttrPair(resourceName, "ignore_id", dataSourceName, "result.0.ignore_id"),
		resource.TestCheckResourceAttrPair(resourceName, "ignore_mac_addresses", dataSourceName, "result.0.ignore_mac_addresses"),
		resource.TestCheckResourceAttrPair(resourceName, "ipam_email_addresses", dataSourceName, "result.0.ipam_email_addresses"),
		resource.TestCheckResourceAttrPair(resourceName, "ipam_threshold_settings", dataSourceName, "result.0.ipam_threshold_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "ipam_trap_settings", dataSourceName, "result.0.ipam_trap_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv4addr", dataSourceName, "result.0.ipv4addr"),
		resource.TestCheckResourceAttrPair(resourceName, "last_rir_registration_update_sent", dataSourceName, "result.0.last_rir_registration_update_sent"),
		resource.TestCheckResourceAttrPair(resourceName, "last_rir_registration_update_status", dataSourceName, "result.0.last_rir_registration_update_status"),
		resource.TestCheckResourceAttrPair(resourceName, "lease_scavenge_time", dataSourceName, "result.0.lease_scavenge_time"),
		resource.TestCheckResourceAttrPair(resourceName, "logic_filter_rules", dataSourceName, "result.0.logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "low_water_mark", dataSourceName, "result.0.low_water_mark"),
		resource.TestCheckResourceAttrPair(resourceName, "low_water_mark_reset", dataSourceName, "result.0.low_water_mark_reset"),
		resource.TestCheckResourceAttrPair(resourceName, "members", dataSourceName, "result.0.members"),
		resource.TestCheckResourceAttrPair(resourceName, "mgm_private", dataSourceName, "result.0.mgm_private"),
		resource.TestCheckResourceAttrPair(resourceName, "mgm_private_overridable", dataSourceName, "result.0.mgm_private_overridable"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_ad_user_data", dataSourceName, "result.0.ms_ad_user_data"),
		resource.TestCheckResourceAttrPair(resourceName, "netmask", dataSourceName, "result.0.netmask"),
		resource.TestCheckResourceAttrPair(resourceName, "network", dataSourceName, "result.0.network"),
		resource.TestCheckResourceAttrPair(resourceName, "func_call", dataSourceName, "result.0.func_call"),
		resource.TestCheckResourceAttrPair(resourceName, "network_container", dataSourceName, "result.0.network_container"),
		resource.TestCheckResourceAttrPair(resourceName, "network_view", dataSourceName, "result.0.network_view"),
		resource.TestCheckResourceAttrPair(resourceName, "nextserver", dataSourceName, "result.0.nextserver"),
		resource.TestCheckResourceAttrPair(resourceName, "options", dataSourceName, "result.0.options"),
		resource.TestCheckResourceAttrPair(resourceName, "port_control_blackout_setting", dataSourceName, "result.0.port_control_blackout_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "pxe_lease_time", dataSourceName, "result.0.pxe_lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "recycle_leases", dataSourceName, "result.0.recycle_leases"),
		resource.TestCheckResourceAttrPair(resourceName, "rir", dataSourceName, "result.0.rir"),
		resource.TestCheckResourceAttrPair(resourceName, "rir_organization", dataSourceName, "result.0.rir_organization"),
		resource.TestCheckResourceAttrPair(resourceName, "rir_registration_action", dataSourceName, "result.0.rir_registration_action"),
		resource.TestCheckResourceAttrPair(resourceName, "rir_registration_status", dataSourceName, "result.0.rir_registration_status"),
		resource.TestCheckResourceAttrPair(resourceName, "same_port_control_discovery_blackout", dataSourceName, "result.0.same_port_control_discovery_blackout"),
		resource.TestCheckResourceAttrPair(resourceName, "send_rir_request", dataSourceName, "result.0.send_rir_request"),
		resource.TestCheckResourceAttrPair(resourceName, "static_hosts", dataSourceName, "result.0.static_hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "subscribe_settings", dataSourceName, "result.0.subscribe_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "template", dataSourceName, "result.0.template"),
		resource.TestCheckResourceAttrPair(resourceName, "total_hosts", dataSourceName, "result.0.total_hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "unmanaged", dataSourceName, "result.0.unmanaged"),
		resource.TestCheckResourceAttrPair(resourceName, "unmanaged_count", dataSourceName, "result.0.unmanaged_count"),
		resource.TestCheckResourceAttrPair(resourceName, "update_dns_on_lease_renewal", dataSourceName, "result.0.update_dns_on_lease_renewal"),
		resource.TestCheckResourceAttrPair(resourceName, "use_authority", dataSourceName, "result.0.use_authority"),
		resource.TestCheckResourceAttrPair(resourceName, "use_blackout_setting", dataSourceName, "result.0.use_blackout_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "use_bootfile", dataSourceName, "result.0.use_bootfile"),
		resource.TestCheckResourceAttrPair(resourceName, "use_bootserver", dataSourceName, "result.0.use_bootserver"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_domainname", dataSourceName, "result.0.use_ddns_domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_generate_hostname", dataSourceName, "result.0.use_ddns_generate_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_ttl", dataSourceName, "result.0.use_ddns_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_update_fixed_addresses", dataSourceName, "result.0.use_ddns_update_fixed_addresses"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_use_option81", dataSourceName, "result.0.use_ddns_use_option81"),
		resource.TestCheckResourceAttrPair(resourceName, "use_deny_bootp", dataSourceName, "result.0.use_deny_bootp"),
		resource.TestCheckResourceAttrPair(resourceName, "use_discovery_basic_polling_settings", dataSourceName, "result.0.use_discovery_basic_polling_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "use_email_list", dataSourceName, "result.0.use_email_list"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_ddns", dataSourceName, "result.0.use_enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_dhcp_thresholds", dataSourceName, "result.0.use_enable_dhcp_thresholds"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_discovery", dataSourceName, "result.0.use_enable_discovery"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_ifmap_publishing", dataSourceName, "result.0.use_enable_ifmap_publishing"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ignore_dhcp_option_list_request", dataSourceName, "result.0.use_ignore_dhcp_option_list_request"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ignore_id", dataSourceName, "result.0.use_ignore_id"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ipam_email_addresses", dataSourceName, "result.0.use_ipam_email_addresses"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ipam_threshold_settings", dataSourceName, "result.0.use_ipam_threshold_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ipam_trap_settings", dataSourceName, "result.0.use_ipam_trap_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "use_lease_scavenge_time", dataSourceName, "result.0.use_lease_scavenge_time"),
		resource.TestCheckResourceAttrPair(resourceName, "use_logic_filter_rules", dataSourceName, "result.0.use_logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "use_mgm_private", dataSourceName, "result.0.use_mgm_private"),
		resource.TestCheckResourceAttrPair(resourceName, "use_nextserver", dataSourceName, "result.0.use_nextserver"),
		resource.TestCheckResourceAttrPair(resourceName, "use_options", dataSourceName, "result.0.use_options"),
		resource.TestCheckResourceAttrPair(resourceName, "use_pxe_lease_time", dataSourceName, "result.0.use_pxe_lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "use_recycle_leases", dataSourceName, "result.0.use_recycle_leases"),
		resource.TestCheckResourceAttrPair(resourceName, "use_subscribe_settings", dataSourceName, "result.0.use_subscribe_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "use_update_dns_on_lease_renewal", dataSourceName, "result.0.use_update_dns_on_lease_renewal"),
		resource.TestCheckResourceAttrPair(resourceName, "use_zone_associations", dataSourceName, "result.0.use_zone_associations"),
		resource.TestCheckResourceAttrPair(resourceName, "utilization", dataSourceName, "result.0.utilization"),
		resource.TestCheckResourceAttrPair(resourceName, "utilization_update", dataSourceName, "result.0.utilization_update"),
		resource.TestCheckResourceAttrPair(resourceName, "vlans", dataSourceName, "result.0.vlans"),
		resource.TestCheckResourceAttrPair(resourceName, "zone_associations", dataSourceName, "result.0.zone_associations"),
	}
}

func testAccNetworkDataSourceConfigFilters(network string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test" {
  network = %q
}

data "nios_ipam_network" "test" {
  filters = {
	  network = nios_ipam_network.test.network
  }
}
`, network)
}

func testAccNetworkDataSourceConfigExtAttrFilters(network, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test" {
  network = %q
  extattrs = {
    Site = %q
  }
}

data "nios_ipam_network" "test" {
  extattrfilters = {
	"Site" = nios_ipam_network.test.extattrs.Site
  }
}
`, network, extAttrsValue)
}
