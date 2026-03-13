package dhcp_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

func TestAccIpv6sharednetworkDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_ipv6sharednetwork.test"
	resourceName := "nios_dhcp_ipv6sharednetwork.test"
	var v dhcp.Ipv6sharednetwork
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6sharednetworkDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6sharednetworkDataSourceConfigFilters(name, networks, network1, network2),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					}, testAccCheckIpv6sharednetworkResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccIpv6sharednetworkDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_ipv6sharednetwork.test"
	resourceName := "nios_dhcp_ipv6sharednetwork.test"
	var v dhcp.Ipv6sharednetwork
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6sharednetworkDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6sharednetworkDataSourceConfigExtAttrFilters(name, networks, network1, network2, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					}, testAccCheckIpv6sharednetworkResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckIpv6sharednetworkResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_domainname", dataSourceName, "result.0.ddns_domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_generate_hostname", dataSourceName, "result.0.ddns_generate_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_server_always_updates", dataSourceName, "result.0.ddns_server_always_updates"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_ttl", dataSourceName, "result.0.ddns_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_use_option81", dataSourceName, "result.0.ddns_use_option81"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "domain_name", dataSourceName, "result.0.domain_name"),
		resource.TestCheckResourceAttrPair(resourceName, "domain_name_servers", dataSourceName, "result.0.domain_name_servers"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_ddns", dataSourceName, "result.0.enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "logic_filter_rules", dataSourceName, "result.0.logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "network_view", dataSourceName, "result.0.network_view"),
		resource.TestCheckResourceAttrPair(resourceName, "networks", dataSourceName, "result.0.networks"),
		resource.TestCheckResourceAttrPair(resourceName, "options", dataSourceName, "result.0.options"),
		resource.TestCheckResourceAttrPair(resourceName, "preferred_lifetime", dataSourceName, "result.0.preferred_lifetime"),
		resource.TestCheckResourceAttrPair(resourceName, "update_dns_on_lease_renewal", dataSourceName, "result.0.update_dns_on_lease_renewal"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_domainname", dataSourceName, "result.0.use_ddns_domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_generate_hostname", dataSourceName, "result.0.use_ddns_generate_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_ttl", dataSourceName, "result.0.use_ddns_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_use_option81", dataSourceName, "result.0.use_ddns_use_option81"),
		resource.TestCheckResourceAttrPair(resourceName, "use_domain_name", dataSourceName, "result.0.use_domain_name"),
		resource.TestCheckResourceAttrPair(resourceName, "use_domain_name_servers", dataSourceName, "result.0.use_domain_name_servers"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_ddns", dataSourceName, "result.0.use_enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "use_logic_filter_rules", dataSourceName, "result.0.use_logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "use_options", dataSourceName, "result.0.use_options"),
		resource.TestCheckResourceAttrPair(resourceName, "use_preferred_lifetime", dataSourceName, "result.0.use_preferred_lifetime"),
		resource.TestCheckResourceAttrPair(resourceName, "use_update_dns_on_lease_renewal", dataSourceName, "result.0.use_update_dns_on_lease_renewal"),
		resource.TestCheckResourceAttrPair(resourceName, "use_valid_lifetime", dataSourceName, "result.0.use_valid_lifetime"),
		resource.TestCheckResourceAttrPair(resourceName, "valid_lifetime", dataSourceName, "result.0.valid_lifetime"),
	}
}

func testAccIpv6sharednetworkDataSourceConfigFilters(name string, networks []string, network1, network2 string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test" {
  name = %q
  networks = %s
}

data "nios_dhcp_ipv6sharednetwork" "test" {
  filters = {
	name = nios_dhcp_ipv6sharednetwork.test.name
  }
}
`, name, networksStr)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkDataSourceConfigExtAttrFilters(name string, networks []string, network1, network2, extAttrsValue string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test" {
  name = %q
  networks = %s
  extattrs = {
    Site = %q
  } 
}

data "nios_dhcp_ipv6sharednetwork" "test" {
  extattrfilters = {
	Site = nios_dhcp_ipv6sharednetwork.test.extattrs.Site
  }
}
`, name, networksStr, extAttrsValue)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}
