package grid_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccMemberDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_grid_member.test"
	resourceName := "nios_grid_member.test"
	var v grid.Member

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMemberDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMemberDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckMemberExists(context.Background(), resourceName, &v),
					}, testAccCheckMemberResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccMemberDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_grid_member.test"
	resourceName := "nios_grid_member.test"
	var v grid.Member
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMemberDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMemberDataSourceConfigExtAttrFilters(acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckMemberExists(context.Background(), resourceName, &v),
					}, testAccCheckMemberResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckMemberResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "active_position", dataSourceName, "result.0.active_position"),
		resource.TestCheckResourceAttrPair(resourceName, "additional_ip_list", dataSourceName, "result.0.additional_ip_list"),
		resource.TestCheckResourceAttrPair(resourceName, "automated_traffic_capture_setting", dataSourceName, "result.0.automated_traffic_capture_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "bgp_as", dataSourceName, "result.0.bgp_as"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "config_addr_type", dataSourceName, "result.0.config_addr_type"),
		resource.TestCheckResourceAttrPair(resourceName, "csp_access_key", dataSourceName, "result.0.csp_access_key"),
		resource.TestCheckResourceAttrPair(resourceName, "csp_member_setting", dataSourceName, "result.0.csp_member_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_resolver_setting", dataSourceName, "result.0.dns_resolver_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "dscp", dataSourceName, "result.0.dscp"),
		resource.TestCheckResourceAttrPair(resourceName, "email_setting", dataSourceName, "result.0.email_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_ha", dataSourceName, "result.0.enable_ha"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_lom", dataSourceName, "result.0.enable_lom"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_member_redirect", dataSourceName, "result.0.enable_member_redirect"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_ro_api_access", dataSourceName, "result.0.enable_ro_api_access"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "external_syslog_backup_servers", dataSourceName, "result.0.external_syslog_backup_servers"),
		resource.TestCheckResourceAttrPair(resourceName, "external_syslog_server_enable", dataSourceName, "result.0.external_syslog_server_enable"),
		resource.TestCheckResourceAttrPair(resourceName, "ha_cloud_platform", dataSourceName, "result.0.ha_cloud_platform"),
		resource.TestCheckResourceAttrPair(resourceName, "ha_on_cloud", dataSourceName, "result.0.ha_on_cloud"),
		resource.TestCheckResourceAttrPair(resourceName, "host_name", dataSourceName, "result.0.host_name"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6_setting", dataSourceName, "result.0.ipv6_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6_static_routes", dataSourceName, "result.0.ipv6_static_routes"),
		resource.TestCheckResourceAttrPair(resourceName, "is_dscp_capable", dataSourceName, "result.0.is_dscp_capable"),
		resource.TestCheckResourceAttrPair(resourceName, "lan2_enabled", dataSourceName, "result.0.lan2_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "lan2_port_setting", dataSourceName, "result.0.lan2_port_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "lom_network_config", dataSourceName, "result.0.lom_network_config"),
		resource.TestCheckResourceAttrPair(resourceName, "lom_users", dataSourceName, "result.0.lom_users"),
		resource.TestCheckResourceAttrPair(resourceName, "master_candidate", dataSourceName, "result.0.master_candidate"),
		resource.TestCheckResourceAttrPair(resourceName, "member_service_communication", dataSourceName, "result.0.member_service_communication"),
		resource.TestCheckResourceAttrPair(resourceName, "mgmt_port_setting", dataSourceName, "result.0.mgmt_port_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "mmdb_ea_build_time", dataSourceName, "result.0.mmdb_ea_build_time"),
		resource.TestCheckResourceAttrPair(resourceName, "mmdb_geoip_build_time", dataSourceName, "result.0.mmdb_geoip_build_time"),
		resource.TestCheckResourceAttrPair(resourceName, "nat_setting", dataSourceName, "result.0.nat_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "node_info", dataSourceName, "result.0.node_info"),
		resource.TestCheckResourceAttrPair(resourceName, "ntp_setting", dataSourceName, "result.0.ntp_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "ospf_list", dataSourceName, "result.0.ospf_list"),
		resource.TestCheckResourceAttrPair(resourceName, "passive_ha_arp_enabled", dataSourceName, "result.0.passive_ha_arp_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "platform", dataSourceName, "result.0.platform"),
		resource.TestCheckResourceAttrPair(resourceName, "pre_provisioning", dataSourceName, "result.0.pre_provisioning"),
		resource.TestCheckResourceAttrPair(resourceName, "preserve_if_owns_delegation", dataSourceName, "result.0.preserve_if_owns_delegation"),
		resource.TestCheckResourceAttrPair(resourceName, "remote_console_access_enable", dataSourceName, "result.0.remote_console_access_enable"),
		resource.TestCheckResourceAttrPair(resourceName, "router_id", dataSourceName, "result.0.router_id"),
		resource.TestCheckResourceAttrPair(resourceName, "service_status", dataSourceName, "result.0.service_status"),
		resource.TestCheckResourceAttrPair(resourceName, "service_type_configuration", dataSourceName, "result.0.service_type_configuration"),
		resource.TestCheckResourceAttrPair(resourceName, "snmp_setting", dataSourceName, "result.0.snmp_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "static_routes", dataSourceName, "result.0.static_routes"),
		resource.TestCheckResourceAttrPair(resourceName, "support_access_enable", dataSourceName, "result.0.support_access_enable"),
		resource.TestCheckResourceAttrPair(resourceName, "support_access_info", dataSourceName, "result.0.support_access_info"),
		resource.TestCheckResourceAttrPair(resourceName, "syslog_proxy_setting", dataSourceName, "result.0.syslog_proxy_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "syslog_servers", dataSourceName, "result.0.syslog_servers"),
		resource.TestCheckResourceAttrPair(resourceName, "syslog_size", dataSourceName, "result.0.syslog_size"),
		resource.TestCheckResourceAttrPair(resourceName, "threshold_traps", dataSourceName, "result.0.threshold_traps"),
		resource.TestCheckResourceAttrPair(resourceName, "time_zone", dataSourceName, "result.0.time_zone"),
		resource.TestCheckResourceAttrPair(resourceName, "traffic_capture_auth_dns_setting", dataSourceName, "result.0.traffic_capture_auth_dns_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "traffic_capture_chr_setting", dataSourceName, "result.0.traffic_capture_chr_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "traffic_capture_qps_setting", dataSourceName, "result.0.traffic_capture_qps_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "traffic_capture_rec_dns_setting", dataSourceName, "result.0.traffic_capture_rec_dns_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "traffic_capture_rec_queries_setting", dataSourceName, "result.0.traffic_capture_rec_queries_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "trap_notifications", dataSourceName, "result.0.trap_notifications"),
		resource.TestCheckResourceAttrPair(resourceName, "upgrade_group", dataSourceName, "result.0.upgrade_group"),
		resource.TestCheckResourceAttrPair(resourceName, "use_automated_traffic_capture", dataSourceName, "result.0.use_automated_traffic_capture"),
		resource.TestCheckResourceAttrPair(resourceName, "use_dns_resolver_setting", dataSourceName, "result.0.use_dns_resolver_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "use_dscp", dataSourceName, "result.0.use_dscp"),
		resource.TestCheckResourceAttrPair(resourceName, "use_email_setting", dataSourceName, "result.0.use_email_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_lom", dataSourceName, "result.0.use_enable_lom"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_member_redirect", dataSourceName, "result.0.use_enable_member_redirect"),
		resource.TestCheckResourceAttrPair(resourceName, "use_external_syslog_backup_servers", dataSourceName, "result.0.use_external_syslog_backup_servers"),
		resource.TestCheckResourceAttrPair(resourceName, "use_remote_console_access_enable", dataSourceName, "result.0.use_remote_console_access_enable"),
		resource.TestCheckResourceAttrPair(resourceName, "use_snmp_setting", dataSourceName, "result.0.use_snmp_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "use_support_access_enable", dataSourceName, "result.0.use_support_access_enable"),
		resource.TestCheckResourceAttrPair(resourceName, "use_syslog_proxy_setting", dataSourceName, "result.0.use_syslog_proxy_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "use_threshold_traps", dataSourceName, "result.0.use_threshold_traps"),
		resource.TestCheckResourceAttrPair(resourceName, "use_time_zone", dataSourceName, "result.0.use_time_zone"),
		resource.TestCheckResourceAttrPair(resourceName, "use_traffic_capture_auth_dns", dataSourceName, "result.0.use_traffic_capture_auth_dns"),
		resource.TestCheckResourceAttrPair(resourceName, "use_traffic_capture_chr", dataSourceName, "result.0.use_traffic_capture_chr"),
		resource.TestCheckResourceAttrPair(resourceName, "use_traffic_capture_qps", dataSourceName, "result.0.use_traffic_capture_qps"),
		resource.TestCheckResourceAttrPair(resourceName, "use_traffic_capture_rec_dns", dataSourceName, "result.0.use_traffic_capture_rec_dns"),
		resource.TestCheckResourceAttrPair(resourceName, "use_traffic_capture_rec_queries", dataSourceName, "result.0.use_traffic_capture_rec_queries"),
		resource.TestCheckResourceAttrPair(resourceName, "use_trap_notifications", dataSourceName, "result.0.use_trap_notifications"),
		resource.TestCheckResourceAttrPair(resourceName, "use_v4_vrrp", dataSourceName, "result.0.use_v4_vrrp"),
		resource.TestCheckResourceAttrPair(resourceName, "vip_setting", dataSourceName, "result.0.vip_setting"),
		resource.TestCheckResourceAttrPair(resourceName, "vpn_mtu", dataSourceName, "result.0.vpn_mtu"),
	}
}

func testAccMemberDataSourceConfigFilters() string {
	return `
resource "nios_grid_member" "test" {
}

data "nios_grid_member" "test" {
  filters = {
	 = nios_grid_member.test.
  }
}
`
}

func testAccMemberDataSourceConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test" {
  extattrs = {
    Site = %q
  } 
}

data "nios_grid_member" "test" {
  extattrfilters = {
	Site = nios_grid_member.test.extattrs.Site
  }
}
`, extAttrsValue)
}
