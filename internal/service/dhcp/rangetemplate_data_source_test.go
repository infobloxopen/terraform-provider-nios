package dhcp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccRangetemplateDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_rangetemplate.test"
	resourceName := "nios_dhcp_rangetemplate.test"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRangetemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRangetemplateDataSourceConfigFilters(name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					}, testAccCheckRangetemplateResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRangetemplateDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_rangetemplate.test"
	resourceName := "nios_dhcp_rangetemplate.test"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRangetemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRangetemplateDataSourceConfigExtAttrFilters(name, numberOfAdresses, offset, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					}, testAccCheckRangetemplateResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRangetemplateResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "bootfile", dataSourceName, "result.0.bootfile"),
		resource.TestCheckResourceAttrPair(resourceName, "bootserver", dataSourceName, "result.0.bootserver"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_api_compatible", dataSourceName, "result.0.cloud_api_compatible"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_domainname", dataSourceName, "result.0.ddns_domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_generate_hostname", dataSourceName, "result.0.ddns_generate_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "delegated_member", dataSourceName, "result.0.delegated_member"),
		resource.TestCheckResourceAttrPair(resourceName, "deny_all_clients", dataSourceName, "result.0.deny_all_clients"),
		resource.TestCheckResourceAttrPair(resourceName, "deny_bootp", dataSourceName, "result.0.deny_bootp"),
		resource.TestCheckResourceAttrPair(resourceName, "email_list", dataSourceName, "result.0.email_list"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_ddns", dataSourceName, "result.0.enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_dhcp_thresholds", dataSourceName, "result.0.enable_dhcp_thresholds"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_email_warnings", dataSourceName, "result.0.enable_email_warnings"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_pxe_lease_time", dataSourceName, "result.0.enable_pxe_lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_snmp_warnings", dataSourceName, "result.0.enable_snmp_warnings"),
		resource.TestCheckResourceAttrPair(resourceName, "exclude", dataSourceName, "result.0.exclude"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "failover_association", dataSourceName, "result.0.failover_association"),
		resource.TestCheckResourceAttrPair(resourceName, "fingerprint_filter_rules", dataSourceName, "result.0.fingerprint_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "high_water_mark", dataSourceName, "result.0.high_water_mark"),
		resource.TestCheckResourceAttrPair(resourceName, "high_water_mark_reset", dataSourceName, "result.0.high_water_mark_reset"),
		resource.TestCheckResourceAttrPair(resourceName, "ignore_dhcp_option_list_request", dataSourceName, "result.0.ignore_dhcp_option_list_request"),
		resource.TestCheckResourceAttrPair(resourceName, "known_clients", dataSourceName, "result.0.known_clients"),
		resource.TestCheckResourceAttrPair(resourceName, "lease_scavenge_time", dataSourceName, "result.0.lease_scavenge_time"),
		resource.TestCheckResourceAttrPair(resourceName, "logic_filter_rules", dataSourceName, "result.0.logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "low_water_mark", dataSourceName, "result.0.low_water_mark"),
		resource.TestCheckResourceAttrPair(resourceName, "low_water_mark_reset", dataSourceName, "result.0.low_water_mark_reset"),
		resource.TestCheckResourceAttrPair(resourceName, "mac_filter_rules", dataSourceName, "result.0.mac_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "member", dataSourceName, "result.0.member"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_options", dataSourceName, "result.0.ms_options"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_server", dataSourceName, "result.0.ms_server"),
		resource.TestCheckResourceAttrPair(resourceName, "nac_filter_rules", dataSourceName, "result.0.nac_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "nextserver", dataSourceName, "result.0.nextserver"),
		resource.TestCheckResourceAttrPair(resourceName, "number_of_addresses", dataSourceName, "result.0.number_of_addresses"),
		resource.TestCheckResourceAttrPair(resourceName, "offset", dataSourceName, "result.0.offset"),
		resource.TestCheckResourceAttrPair(resourceName, "option_filter_rules", dataSourceName, "result.0.option_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "options", dataSourceName, "result.0.options"),
		resource.TestCheckResourceAttrPair(resourceName, "pxe_lease_time", dataSourceName, "result.0.pxe_lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "recycle_leases", dataSourceName, "result.0.recycle_leases"),
		resource.TestCheckResourceAttrPair(resourceName, "relay_agent_filter_rules", dataSourceName, "result.0.relay_agent_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "server_association_type", dataSourceName, "result.0.server_association_type"),
		resource.TestCheckResourceAttrPair(resourceName, "unknown_clients", dataSourceName, "result.0.unknown_clients"),
		resource.TestCheckResourceAttrPair(resourceName, "update_dns_on_lease_renewal", dataSourceName, "result.0.update_dns_on_lease_renewal"),
		resource.TestCheckResourceAttrPair(resourceName, "use_bootfile", dataSourceName, "result.0.use_bootfile"),
		resource.TestCheckResourceAttrPair(resourceName, "use_bootserver", dataSourceName, "result.0.use_bootserver"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_domainname", dataSourceName, "result.0.use_ddns_domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_generate_hostname", dataSourceName, "result.0.use_ddns_generate_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "use_deny_bootp", dataSourceName, "result.0.use_deny_bootp"),
		resource.TestCheckResourceAttrPair(resourceName, "use_email_list", dataSourceName, "result.0.use_email_list"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_ddns", dataSourceName, "result.0.use_enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_dhcp_thresholds", dataSourceName, "result.0.use_enable_dhcp_thresholds"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ignore_dhcp_option_list_request", dataSourceName, "result.0.use_ignore_dhcp_option_list_request"),
		resource.TestCheckResourceAttrPair(resourceName, "use_known_clients", dataSourceName, "result.0.use_known_clients"),
		resource.TestCheckResourceAttrPair(resourceName, "use_lease_scavenge_time", dataSourceName, "result.0.use_lease_scavenge_time"),
		resource.TestCheckResourceAttrPair(resourceName, "use_logic_filter_rules", dataSourceName, "result.0.use_logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ms_options", dataSourceName, "result.0.use_ms_options"),
		resource.TestCheckResourceAttrPair(resourceName, "use_nextserver", dataSourceName, "result.0.use_nextserver"),
		resource.TestCheckResourceAttrPair(resourceName, "use_options", dataSourceName, "result.0.use_options"),
		resource.TestCheckResourceAttrPair(resourceName, "use_pxe_lease_time", dataSourceName, "result.0.use_pxe_lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "use_recycle_leases", dataSourceName, "result.0.use_recycle_leases"),
		resource.TestCheckResourceAttrPair(resourceName, "use_unknown_clients", dataSourceName, "result.0.use_unknown_clients"),
		resource.TestCheckResourceAttrPair(resourceName, "use_update_dns_on_lease_renewal", dataSourceName, "result.0.use_update_dns_on_lease_renewal"),
	}
}

func testAccRangetemplateDataSourceConfigFilters(name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test" {
	  name = %q
	  number_of_addresses = %d
	  offset = %d
}

data "nios_dhcp_rangetemplate" "test" {
  filters = {
	 name = nios_dhcp_rangetemplate.test.name
  }
}
`, name, numberOfAddresses, offset)
}

func testAccRangetemplateDataSourceConfigExtAttrFilters(name string, numberOfAddresses, offset int64, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test" {
  name = %q
  number_of_addresses = %d
  offset = %d
  extattrs = {
    Site = %q
  } 
}

data "nios_dhcp_rangetemplate" "test" {
  extattrfilters = {
	Site = nios_dhcp_rangetemplate.test.extattrs.Site
  }
}
`, name, numberOfAddresses, offset, extAttrsValue)
}
