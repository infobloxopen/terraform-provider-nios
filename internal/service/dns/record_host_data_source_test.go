package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

func TestAccRecordHostDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_record_host.test"
	resourceName := "nios_ip_allocation.test"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.10",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIPAllocationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordHostDataSourceConfigFilters(name, "default", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordHostResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordHostDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_record_host.test"
	resourceName := "nios_ip_allocation.test"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.10",
		},
	}
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIPAllocationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordHostDataSourceConfigExtAttrFilters(name, "default", extAttrValue, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordHostResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordHostResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "aliases", dataSourceName, "result.0.aliases"),
		resource.TestCheckResourceAttrPair(resourceName, "allow_telnet", dataSourceName, "result.0.allow_telnet"),
		resource.TestCheckResourceAttrPair(resourceName, "cli_credentials", dataSourceName, "result.0.cli_credentials"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_info", dataSourceName, "result.0.cloud_info"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "configure_for_dns", dataSourceName, "result.0.configure_for_dns"),
		resource.TestCheckResourceAttrPair(resourceName, "creation_time", dataSourceName, "result.0.creation_time"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_protected", dataSourceName, "result.0.ddns_protected"),
		resource.TestCheckResourceAttrPair(resourceName, "device_description", dataSourceName, "result.0.device_description"),
		resource.TestCheckResourceAttrPair(resourceName, "device_location", dataSourceName, "result.0.device_location"),
		resource.TestCheckResourceAttrPair(resourceName, "device_type", dataSourceName, "result.0.device_type"),
		resource.TestCheckResourceAttrPair(resourceName, "device_vendor", dataSourceName, "result.0.device_vendor"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "disable_discovery", dataSourceName, "result.0.disable_discovery"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_aliases", dataSourceName, "result.0.dns_aliases"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_name", dataSourceName, "result.0.dns_name"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_immediate_discovery", dataSourceName, "result.0.enable_immediate_discovery"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv4addrs", dataSourceName, "result.0.ipv4addrs"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6addrs", dataSourceName, "result.0.ipv6addrs"),
		resource.TestCheckResourceAttrPair(resourceName, "last_queried", dataSourceName, "result.0.last_queried"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_ad_user_data", dataSourceName, "result.0.ms_ad_user_data"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "network_view", dataSourceName, "result.0.network_view"),
		resource.TestCheckResourceAttrPair(resourceName, "restart_if_needed", dataSourceName, "result.0.restart_if_needed"),
		resource.TestCheckResourceAttrPair(resourceName, "rrset_order", dataSourceName, "result.0.rrset_order"),
		resource.TestCheckResourceAttrPair(resourceName, "snmp3_credential", dataSourceName, "result.0.snmp3_credential"),
		resource.TestCheckResourceAttrPair(resourceName, "snmp_credential", dataSourceName, "result.0.snmp_credential"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_cli_credentials", dataSourceName, "result.0.use_cli_credentials"),
		resource.TestCheckResourceAttrPair(resourceName, "use_dns_ea_inheritance", dataSourceName, "result.0.use_dns_ea_inheritance"),
		resource.TestCheckResourceAttrPair(resourceName, "use_snmp3_credential", dataSourceName, "result.0.use_snmp3_credential"),
		resource.TestCheckResourceAttrPair(resourceName, "use_snmp_credential", dataSourceName, "result.0.use_snmp_credential"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordHostDataSourceConfigFilters(name, view string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test" {
	name = %q
	ipv4addrs = %s
	view = %q
}

data "nios_record_host" "test" {
  filters = {
	name = nios_ip_allocation.test.name
  }
}
`, name, ipv4addrHCL, view)
}

func testAccRecordHostDataSourceConfigExtAttrFilters(name, view, extAttrsValue string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test" {
  name = %q
  ipv4addrs = %s
  view = %q
  extattrs = {
	Site = %q
  }
}

data "nios_record_host" "test" {
  extattrfilters = {
	Site = nios_ip_allocation.test.extattrs.Site
  }
}
`, name, ipv4addrHCL, view, extAttrsValue)
}
