package ipam_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccNetworkcontainerDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_ipam_network_container.test"
	resourceName := "nios_ipam_network_container.test"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetworkcontainerDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkcontainerDataSourceConfigFilters(network),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					}, testAccCheckNetworkcontainerResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccNetworkcontainerDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_ipam_network_container.test"
	resourceName := "nios_ipam_network_container.test"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetworkcontainerDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkcontainerDataSourceConfigExtAttrFilters(network, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					}, testAccCheckNetworkcontainerResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckNetworkcontainerResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "authority", dataSourceName, "result.0.authority"),
		resource.TestCheckResourceAttrPair(resourceName, "bootfile", dataSourceName, "result.0.bootfile"),
		resource.TestCheckResourceAttrPair(resourceName, "bootserver", dataSourceName, "result.0.bootserver"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_info", dataSourceName, "result.0.cloud_info"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_domainname", dataSourceName, "result.0.ddns_domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_generate_hostname", dataSourceName, "result.0.ddns_generate_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_server_always_updates", dataSourceName, "result.0.ddns_server_always_updates"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_ttl", dataSourceName, "result.0.ddns_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_update_fixed_addresses", dataSourceName, "result.0.ddns_update_fixed_addresses"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_use_option81", dataSourceName, "result.0.ddns_use_option81"),
		resource.TestCheckResourceAttrPair(resourceName, "delete_reason", dataSourceName, "result.0.delete_reason"),
		resource.TestCheckResourceAttrPair(resourceName, "deny_bootp", dataSourceName, "result.0.deny_bootp"),
		resource.TestCheckResourceAttrPair(resourceName, "discover_now_status", dataSourceName, "result.0.discover_now_status"),
		resource.TestCheckResourceAttrPair(resourceName, "discovery_basic_poll_settings", dataSourceName, "result.0.discovery_basic_poll_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "discovery_blackout_setting", dataSourceName, "result.0.discovery_blackout_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "discovery_engine_type", dataSourceName, "result.0.discovery_engine_type"),
		resource.TestCheckResourceAttrPair(resourceName, "discovery_member", dataSourceName, "result.0.discovery_member"),
		resource.TestCheckResourceAttrPair(resourceName, "email_list", dataSourceName, "result.0.email_list"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_ddns", dataSourceName, "result.0.enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_dhcp_thresholds", dataSourceName, "result.0.enable_dhcp_thresholds"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_discovery", dataSourceName, "result.0.enable_discovery"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_email_warnings", dataSourceName, "result.0.enable_email_warnings"),
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
		resource.TestCheckResourceAttrPair(resourceName, "last_rir_registration_update_sent", dataSourceName, "result.0.last_rir_registration_update_sent"),
		resource.TestCheckResourceAttrPair(resourceName, "last_rir_registration_update_status", dataSourceName, "result.0.last_rir_registration_update_status"),
		resource.TestCheckResourceAttrPair(resourceName, "lease_scavenge_time", dataSourceName, "result.0.lease_scavenge_time"),
		resource.TestCheckResourceAttrPair(resourceName, "logic_filter_rules", dataSourceName, "result.0.logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "low_water_mark", dataSourceName, "result.0.low_water_mark"),
		resource.TestCheckResourceAttrPair(resourceName, "low_water_mark_reset", dataSourceName, "result.0.low_water_mark_reset"),
		resource.TestCheckResourceAttrPair(resourceName, "mgm_private", dataSourceName, "result.0.mgm_private"),
		resource.TestCheckResourceAttrPair(resourceName, "mgm_private_overridable", dataSourceName, "result.0.mgm_private_overridable"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_ad_user_data", dataSourceName, "result.0.ms_ad_user_data"),
		resource.TestCheckResourceAttrPair(resourceName, "network", dataSourceName, "result.0.network"),
		resource.TestCheckResourceAttrPair(resourceName, "func_call", dataSourceName, "result.0.func_call"),
		resource.TestCheckResourceAttrPair(resourceName, "network_container", dataSourceName, "result.0.network_container"),
		resource.TestCheckResourceAttrPair(resourceName, "network_view", dataSourceName, "result.0.network_view"),
		resource.TestCheckResourceAttrPair(resourceName, "nextserver", dataSourceName, "result.0.nextserver"),
		resource.TestCheckResourceAttrPair(resourceName, "options", dataSourceName, "result.0.options"),
		resource.TestCheckResourceAttrPair(resourceName, "port_control_blackout_setting", dataSourceName, "result.0.port_control_blackout_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "pxe_lease_time", dataSourceName, "result.0.pxe_lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "recycle_leases", dataSourceName, "result.0.recycle_leases"),
		resource.TestCheckResourceAttrPair(resourceName, "remove_subnets", dataSourceName, "result.0.remove_subnets"),
		resource.TestCheckResourceAttrPair(resourceName, "rir", dataSourceName, "result.0.rir"),
		resource.TestCheckResourceAttrPair(resourceName, "rir_organization", dataSourceName, "result.0.rir_organization"),
		resource.TestCheckResourceAttrPair(resourceName, "rir_registration_action", dataSourceName, "result.0.rir_registration_action"),
		resource.TestCheckResourceAttrPair(resourceName, "rir_registration_status", dataSourceName, "result.0.rir_registration_status"),
		resource.TestCheckResourceAttrPair(resourceName, "same_port_control_discovery_blackout", dataSourceName, "result.0.same_port_control_discovery_blackout"),
		resource.TestCheckResourceAttrPair(resourceName, "send_rir_request", dataSourceName, "result.0.send_rir_request"),
		resource.TestCheckResourceAttrPair(resourceName, "subscribe_settings", dataSourceName, "result.0.subscribe_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "unmanaged", dataSourceName, "result.0.unmanaged"),
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
		resource.TestCheckResourceAttrPair(resourceName, "zone_associations", dataSourceName, "result.0.zone_associations"),
	}
}

func testAccNetworkcontainerDataSourceConfigFilters(network string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test" {
  network = %q
}

data "nios_ipam_network_container" "test" {
  filters = {
	network = nios_ipam_network_container.test.network
  }
}
`, network)
}

func testAccNetworkcontainerDataSourceConfigExtAttrFilters(network, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test" {
  network = %q
  extattrs = {
    Site = %q
  }
}

data "nios_ipam_network_container" "test" {
  extattrfilters = {
	"Site" = nios_ipam_network_container.test.extattrs.Site
  }
}
`, network, extAttrsValue)
}
