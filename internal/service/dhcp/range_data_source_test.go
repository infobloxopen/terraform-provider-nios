package dhcp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccRangeDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_range.test"
	resourceName := "nios_dhcp_range.test"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRangeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRangeDataSourceConfigFilters(startAddr, endAddr),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRangeExists(context.Background(), resourceName, &v),
					}, testAccCheckRangeResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRangeDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_range.test"
	resourceName := "nios_dhcp_range.test"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRangeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRangeDataSourceConfigExtAttrFilters(startAddr, endAddr, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRangeExists(context.Background(), resourceName, &v),
					}, testAccCheckRangeResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRangeResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "always_update_dns", dataSourceName, "result.0.always_update_dns"),
		resource.TestCheckResourceAttrPair(resourceName, "bootfile", dataSourceName, "result.0.bootfile"),
		resource.TestCheckResourceAttrPair(resourceName, "bootserver", dataSourceName, "result.0.bootserver"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_info", dataSourceName, "result.0.cloud_info"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_domainname", dataSourceName, "result.0.ddns_domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_generate_hostname", dataSourceName, "result.0.ddns_generate_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "deny_all_clients", dataSourceName, "result.0.deny_all_clients"),
		resource.TestCheckResourceAttrPair(resourceName, "deny_bootp", dataSourceName, "result.0.deny_bootp"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_utilization", dataSourceName, "result.0.dhcp_utilization"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_utilization_status", dataSourceName, "result.0.dhcp_utilization_status"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "discover_now_status", dataSourceName, "result.0.discover_now_status"),
		resource.TestCheckResourceAttrPair(resourceName, "discovery_basic_poll_settings", dataSourceName, "result.0.discovery_basic_poll_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "discovery_blackout_setting", dataSourceName, "result.0.discovery_blackout_setting"),
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
		resource.TestCheckResourceAttrPair(resourceName, "end_addr", dataSourceName, "result.0.end_addr"),
		resource.TestCheckResourceAttrPair(resourceName, "endpoint_sources", dataSourceName, "result.0.endpoint_sources"),
		resource.TestCheckResourceAttrPair(resourceName, "exclude", dataSourceName, "result.0.exclude"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "failover_association", dataSourceName, "result.0.failover_association"),
		resource.TestCheckResourceAttrPair(resourceName, "fingerprint_filter_rules", dataSourceName, "result.0.fingerprint_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "high_water_mark", dataSourceName, "result.0.high_water_mark"),
		resource.TestCheckResourceAttrPair(resourceName, "high_water_mark_reset", dataSourceName, "result.0.high_water_mark_reset"),
		resource.TestCheckResourceAttrPair(resourceName, "ignore_dhcp_option_list_request", dataSourceName, "result.0.ignore_dhcp_option_list_request"),
		resource.TestCheckResourceAttrPair(resourceName, "ignore_id", dataSourceName, "result.0.ignore_id"),
		resource.TestCheckResourceAttrPair(resourceName, "ignore_mac_addresses", dataSourceName, "result.0.ignore_mac_addresses"),
		resource.TestCheckResourceAttrPair(resourceName, "is_split_scope", dataSourceName, "result.0.is_split_scope"),
		resource.TestCheckResourceAttrPair(resourceName, "known_clients", dataSourceName, "result.0.known_clients"),
		resource.TestCheckResourceAttrPair(resourceName, "lease_scavenge_time", dataSourceName, "result.0.lease_scavenge_time"),
		resource.TestCheckResourceAttrPair(resourceName, "logic_filter_rules", dataSourceName, "result.0.logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "low_water_mark", dataSourceName, "result.0.low_water_mark"),
		resource.TestCheckResourceAttrPair(resourceName, "low_water_mark_reset", dataSourceName, "result.0.low_water_mark_reset"),
		resource.TestCheckResourceAttrPair(resourceName, "mac_filter_rules", dataSourceName, "result.0.mac_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "member", dataSourceName, "result.0.member"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_ad_user_data", dataSourceName, "result.0.ms_ad_user_data"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_options", dataSourceName, "result.0.ms_options"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_server", dataSourceName, "result.0.ms_server"),
		resource.TestCheckResourceAttrPair(resourceName, "nac_filter_rules", dataSourceName, "result.0.nac_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "network", dataSourceName, "result.0.network"),
		resource.TestCheckResourceAttrPair(resourceName, "network_view", dataSourceName, "result.0.network_view"),
		resource.TestCheckResourceAttrPair(resourceName, "nextserver", dataSourceName, "result.0.nextserver"),
		resource.TestCheckResourceAttrPair(resourceName, "option_filter_rules", dataSourceName, "result.0.option_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "options", dataSourceName, "result.0.options"),
		resource.TestCheckResourceAttrPair(resourceName, "port_control_blackout_setting", dataSourceName, "result.0.port_control_blackout_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "pxe_lease_time", dataSourceName, "result.0.pxe_lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "recycle_leases", dataSourceName, "result.0.recycle_leases"),
		resource.TestCheckResourceAttrPair(resourceName, "relay_agent_filter_rules", dataSourceName, "result.0.relay_agent_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "restart_if_needed", dataSourceName, "result.0.restart_if_needed"),
		resource.TestCheckResourceAttrPair(resourceName, "same_port_control_discovery_blackout", dataSourceName, "result.0.same_port_control_discovery_blackout"),
		resource.TestCheckResourceAttrPair(resourceName, "server_association_type", dataSourceName, "result.0.server_association_type"),
		resource.TestCheckResourceAttrPair(resourceName, "split_member", dataSourceName, "result.0.split_member"),
		resource.TestCheckResourceAttrPair(resourceName, "split_scope_exclusion_percent", dataSourceName, "result.0.split_scope_exclusion_percent"),
		resource.TestCheckResourceAttrPair(resourceName, "start_addr", dataSourceName, "result.0.start_addr"),
		resource.TestCheckResourceAttrPair(resourceName, "static_hosts", dataSourceName, "result.0.static_hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "subscribe_settings", dataSourceName, "result.0.subscribe_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "template", dataSourceName, "result.0.template"),
		resource.TestCheckResourceAttrPair(resourceName, "total_hosts", dataSourceName, "result.0.total_hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "unknown_clients", dataSourceName, "result.0.unknown_clients"),
		resource.TestCheckResourceAttrPair(resourceName, "update_dns_on_lease_renewal", dataSourceName, "result.0.update_dns_on_lease_renewal"),
		resource.TestCheckResourceAttrPair(resourceName, "use_blackout_setting", dataSourceName, "result.0.use_blackout_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "use_bootfile", dataSourceName, "result.0.use_bootfile"),
		resource.TestCheckResourceAttrPair(resourceName, "use_bootserver", dataSourceName, "result.0.use_bootserver"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_domainname", dataSourceName, "result.0.use_ddns_domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_generate_hostname", dataSourceName, "result.0.use_ddns_generate_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "use_deny_bootp", dataSourceName, "result.0.use_deny_bootp"),
		resource.TestCheckResourceAttrPair(resourceName, "use_discovery_basic_polling_settings", dataSourceName, "result.0.use_discovery_basic_polling_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "use_email_list", dataSourceName, "result.0.use_email_list"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_ddns", dataSourceName, "result.0.use_enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_dhcp_thresholds", dataSourceName, "result.0.use_enable_dhcp_thresholds"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_discovery", dataSourceName, "result.0.use_enable_discovery"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_ifmap_publishing", dataSourceName, "result.0.use_enable_ifmap_publishing"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ignore_dhcp_option_list_request", dataSourceName, "result.0.use_ignore_dhcp_option_list_request"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ignore_id", dataSourceName, "result.0.use_ignore_id"),
		resource.TestCheckResourceAttrPair(resourceName, "use_known_clients", dataSourceName, "result.0.use_known_clients"),
		resource.TestCheckResourceAttrPair(resourceName, "use_lease_scavenge_time", dataSourceName, "result.0.use_lease_scavenge_time"),
		resource.TestCheckResourceAttrPair(resourceName, "use_logic_filter_rules", dataSourceName, "result.0.use_logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ms_options", dataSourceName, "result.0.use_ms_options"),
		resource.TestCheckResourceAttrPair(resourceName, "use_nextserver", dataSourceName, "result.0.use_nextserver"),
		resource.TestCheckResourceAttrPair(resourceName, "use_options", dataSourceName, "result.0.use_options"),
		resource.TestCheckResourceAttrPair(resourceName, "use_pxe_lease_time", dataSourceName, "result.0.use_pxe_lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "use_recycle_leases", dataSourceName, "result.0.use_recycle_leases"),
		resource.TestCheckResourceAttrPair(resourceName, "use_subscribe_settings", dataSourceName, "result.0.use_subscribe_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "use_unknown_clients", dataSourceName, "result.0.use_unknown_clients"),
		resource.TestCheckResourceAttrPair(resourceName, "use_update_dns_on_lease_renewal", dataSourceName, "result.0.use_update_dns_on_lease_renewal"),
	}
}

func testAccRangeDataSourceConfigFilters(startAddr, endAddr string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test" {
  start_addr = %q
  end_addr = %q
}

data "nios_dhcp_range" "test" {
  filters = {
	start_addr = nios_dhcp_range.test.start_addr
  }
}
`, startAddr, endAddr)
}

func testAccRangeDataSourceConfigExtAttrFilters(startAddr, endAddr, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test" {
  start_addr = %q
  end_addr = %q
  extattrs = {
    Site = %q
  } 
}

data "nios_dhcp_range" "test" {
  extattrfilters = {
	Site = nios_dhcp_range.test.extattrs.Site
  }
}
`, startAddr, endAddr, extAttrsValue)
}
