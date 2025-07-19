package dhcp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccFixedaddressDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_fixed_address.test"
	resourceName := "nios_dhcp_fixed_address.test"
	var v dhcp.Fixedaddress
	ip := acctest.RandomIPWithSpecificOctetsSet("16.0.0")
	agentCircuitID := acctest.RandomNumber(255)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFixedaddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFixedaddressDataSourceConfigFilters(ip, "CIRCUIT_ID", agentCircuitID),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					}, testAccCheckFixedaddressResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccFixedaddressDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_fixed_address.test"
	resourceName := "nios_dhcp_fixed_address.test"
	var v dhcp.Fixedaddress
	ip := acctest.RandomIPWithSpecificOctetsSet("16.0.0")
	agentCircuitID := acctest.RandomNumber(255)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFixedaddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFixedaddressDataSourceConfigExtAttrFilters(ip, "CIRCUIT_ID", agentCircuitID, "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					}, testAccCheckFixedaddressResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccFixedaddressDataSource_MsServerStruct(t *testing.T) {
	dataSourceName := "data.nios_dhcp_fixed_address.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFixedaddressDataSourcConfigeMsServerStruct("msdhcpserver", "10.34.98.68"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "result.#", "1"),
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckFixedaddressResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "agent_circuit_id", dataSourceName, "result.0.agent_circuit_id"),
		resource.TestCheckResourceAttrPair(resourceName, "agent_remote_id", dataSourceName, "result.0.agent_remote_id"),
		resource.TestCheckResourceAttrPair(resourceName, "allow_telnet", dataSourceName, "result.0.allow_telnet"),
		resource.TestCheckResourceAttrPair(resourceName, "always_update_dns", dataSourceName, "result.0.always_update_dns"),
		resource.TestCheckResourceAttrPair(resourceName, "bootfile", dataSourceName, "result.0.bootfile"),
		resource.TestCheckResourceAttrPair(resourceName, "bootserver", dataSourceName, "result.0.bootserver"),
		resource.TestCheckResourceAttrPair(resourceName, "cli_credentials", dataSourceName, "result.0.cli_credentials"),
		resource.TestCheckResourceAttrPair(resourceName, "client_identifier_prepend_zero", dataSourceName, "result.0.client_identifier_prepend_zero"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_info", dataSourceName, "result.0.cloud_info"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_domainname", dataSourceName, "result.0.ddns_domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_hostname", dataSourceName, "result.0.ddns_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "deny_bootp", dataSourceName, "result.0.deny_bootp"),
		resource.TestCheckResourceAttrPair(resourceName, "device_description", dataSourceName, "result.0.device_description"),
		resource.TestCheckResourceAttrPair(resourceName, "device_location", dataSourceName, "result.0.device_location"),
		resource.TestCheckResourceAttrPair(resourceName, "device_type", dataSourceName, "result.0.device_type"),
		resource.TestCheckResourceAttrPair(resourceName, "device_vendor", dataSourceName, "result.0.device_vendor"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_client_identifier", dataSourceName, "result.0.dhcp_client_identifier"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "disable_discovery", dataSourceName, "result.0.disable_discovery"),
		resource.TestCheckResourceAttrPair(resourceName, "discover_now_status", dataSourceName, "result.0.discover_now_status"),
		resource.TestCheckResourceAttrPair(resourceName, "discovered_data", dataSourceName, "result.0.discovered_data"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_ddns", dataSourceName, "result.0.enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_immediate_discovery", dataSourceName, "result.0.enable_immediate_discovery"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_pxe_lease_time", dataSourceName, "result.0.enable_pxe_lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "ignore_dhcp_option_list_request", dataSourceName, "result.0.ignore_dhcp_option_list_request"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv4addr", dataSourceName, "result.0.ipv4addr"),
		resource.TestCheckResourceAttrPair(resourceName, "func_call", dataSourceName, "result.0.func_call"),
		resource.TestCheckResourceAttrPair(resourceName, "is_invalid_mac", dataSourceName, "result.0.is_invalid_mac"),
		resource.TestCheckResourceAttrPair(resourceName, "logic_filter_rules", dataSourceName, "result.0.logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "mac", dataSourceName, "result.0.mac"),
		resource.TestCheckResourceAttrPair(resourceName, "match_client", dataSourceName, "result.0.match_client"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_ad_user_data", dataSourceName, "result.0.ms_ad_user_data"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_options", dataSourceName, "result.0.ms_options"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_server", dataSourceName, "result.0.ms_server"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "network", dataSourceName, "result.0.network"),
		resource.TestCheckResourceAttrPair(resourceName, "network_view", dataSourceName, "result.0.network_view"),
		resource.TestCheckResourceAttrPair(resourceName, "nextserver", dataSourceName, "result.0.nextserver"),
		resource.TestCheckResourceAttrPair(resourceName, "options", dataSourceName, "result.0.options"),
		resource.TestCheckResourceAttrPair(resourceName, "pxe_lease_time", dataSourceName, "result.0.pxe_lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "reserved_interface", dataSourceName, "result.0.reserved_interface"),
		resource.TestCheckResourceAttrPair(resourceName, "restart_if_needed", dataSourceName, "result.0.restart_if_needed"),
		resource.TestCheckResourceAttrPair(resourceName, "snmp3_credential", dataSourceName, "result.0.snmp3_credential"),
		resource.TestCheckResourceAttrPair(resourceName, "snmp_credential", dataSourceName, "result.0.snmp_credential"),
		resource.TestCheckResourceAttrPair(resourceName, "template", dataSourceName, "result.0.template"),
		resource.TestCheckResourceAttrPair(resourceName, "use_bootfile", dataSourceName, "result.0.use_bootfile"),
		resource.TestCheckResourceAttrPair(resourceName, "use_bootserver", dataSourceName, "result.0.use_bootserver"),
		resource.TestCheckResourceAttrPair(resourceName, "use_cli_credentials", dataSourceName, "result.0.use_cli_credentials"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_domainname", dataSourceName, "result.0.use_ddns_domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "use_deny_bootp", dataSourceName, "result.0.use_deny_bootp"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_ddns", dataSourceName, "result.0.use_enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ignore_dhcp_option_list_request", dataSourceName, "result.0.use_ignore_dhcp_option_list_request"),
		resource.TestCheckResourceAttrPair(resourceName, "use_logic_filter_rules", dataSourceName, "result.0.use_logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ms_options", dataSourceName, "result.0.use_ms_options"),
		resource.TestCheckResourceAttrPair(resourceName, "use_nextserver", dataSourceName, "result.0.use_nextserver"),
		resource.TestCheckResourceAttrPair(resourceName, "use_options", dataSourceName, "result.0.use_options"),
		resource.TestCheckResourceAttrPair(resourceName, "use_pxe_lease_time", dataSourceName, "result.0.use_pxe_lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "use_snmp3_credential", dataSourceName, "result.0.use_snmp3_credential"),
		resource.TestCheckResourceAttrPair(resourceName, "use_snmp_credential", dataSourceName, "result.0.use_snmp_credential"),
	}
}

func testAccFixedaddressDataSourceConfigFilters(ip, matchClient string, agentCircuitID int) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
}

data "nios_dhcp_fixed_address" "test" {
  filters = {
	ipv4addr = nios_dhcp_fixed_address.test.ipv4addr
  }
}
`, ip, matchClient, agentCircuitID)
}

func testAccFixedaddressDataSourceConfigExtAttrFilters(ip, matchClient string, agentCircuitID int, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	extattrs = {
		Site = %q
	}
}

data "nios_dhcp_fixed_address" "test" {
  extattrfilters = {
	Site = nios_dhcp_fixed_address.test.extattrs.Site
  }
}
`, ip, matchClient, agentCircuitID, extAttrsValue)
}

func testAccFixedaddressDataSourcConfigeMsServerStruct(structValue, ipv4Adrr string) string {
	return fmt.Sprintf(`
data "nios_dhcp_fixed_address" "test" {
	body = {
		ms_server = {
			struct = %q
			ipv4addr = %q
		}
	}
	max_results = 1
	paging = 0
}
`, structValue, ipv4Adrr)
}
