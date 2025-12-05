package dhcp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccRoaminghostDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_roaminghost.test"
	resourceName := "nios_dhcp_roaminghost.test"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRoaminghostDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRoaminghostDataSourceConfigFilters(name, mac, "IPV4", "MAC_ADDRESS"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					}, testAccCheckRoaminghostResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRoaminghostDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_roaminghost.test"
	resourceName := "nios_dhcp_roaminghost.test"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRoaminghostDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRoaminghostDataSourceConfigExtAttrFilters(name, mac, "IPV4", "MAC_ADDRESS", acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					}, testAccCheckRoaminghostResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRoaminghostResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "address_type", dataSourceName, "result.0.address_type"),
		resource.TestCheckResourceAttrPair(resourceName, "bootfile", dataSourceName, "result.0.bootfile"),
		resource.TestCheckResourceAttrPair(resourceName, "bootserver", dataSourceName, "result.0.bootserver"),
		resource.TestCheckResourceAttrPair(resourceName, "client_identifier_prepend_zero", dataSourceName, "result.0.client_identifier_prepend_zero"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_domainname", dataSourceName, "result.0.ddns_domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_hostname", dataSourceName, "result.0.ddns_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "deny_bootp", dataSourceName, "result.0.deny_bootp"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_client_identifier", dataSourceName, "result.0.dhcp_client_identifier"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_ddns", dataSourceName, "result.0.enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_pxe_lease_time", dataSourceName, "result.0.enable_pxe_lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "force_roaming_hostname", dataSourceName, "result.0.force_roaming_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "ignore_dhcp_option_list_request", dataSourceName, "result.0.ignore_dhcp_option_list_request"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6_client_hostname", dataSourceName, "result.0.ipv6_client_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6_ddns_domainname", dataSourceName, "result.0.ipv6_ddns_domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6_ddns_hostname", dataSourceName, "result.0.ipv6_ddns_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6_domain_name", dataSourceName, "result.0.ipv6_domain_name"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6_domain_name_servers", dataSourceName, "result.0.ipv6_domain_name_servers"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6_duid", dataSourceName, "result.0.ipv6_duid"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6_enable_ddns", dataSourceName, "result.0.ipv6_enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6_force_roaming_hostname", dataSourceName, "result.0.ipv6_force_roaming_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6_mac_address", dataSourceName, "result.0.ipv6_mac_address"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6_match_option", dataSourceName, "result.0.ipv6_match_option"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6_options", dataSourceName, "result.0.ipv6_options"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6_template", dataSourceName, "result.0.ipv6_template"),
		resource.TestCheckResourceAttrPair(resourceName, "mac", dataSourceName, "result.0.mac"),
		resource.TestCheckResourceAttrPair(resourceName, "match_client", dataSourceName, "result.0.match_client"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "network_view", dataSourceName, "result.0.network_view"),
		resource.TestCheckResourceAttrPair(resourceName, "nextserver", dataSourceName, "result.0.nextserver"),
		resource.TestCheckResourceAttrPair(resourceName, "options", dataSourceName, "result.0.options"),
		resource.TestCheckResourceAttrPair(resourceName, "preferred_lifetime", dataSourceName, "result.0.preferred_lifetime"),
		resource.TestCheckResourceAttrPair(resourceName, "pxe_lease_time", dataSourceName, "result.0.pxe_lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "template", dataSourceName, "result.0.template"),
		resource.TestCheckResourceAttrPair(resourceName, "use_bootfile", dataSourceName, "result.0.use_bootfile"),
		resource.TestCheckResourceAttrPair(resourceName, "use_bootserver", dataSourceName, "result.0.use_bootserver"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_domainname", dataSourceName, "result.0.use_ddns_domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "use_deny_bootp", dataSourceName, "result.0.use_deny_bootp"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_ddns", dataSourceName, "result.0.use_enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ignore_dhcp_option_list_request", dataSourceName, "result.0.use_ignore_dhcp_option_list_request"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ipv6_ddns_domainname", dataSourceName, "result.0.use_ipv6_ddns_domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ipv6_domain_name", dataSourceName, "result.0.use_ipv6_domain_name"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ipv6_domain_name_servers", dataSourceName, "result.0.use_ipv6_domain_name_servers"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ipv6_enable_ddns", dataSourceName, "result.0.use_ipv6_enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ipv6_options", dataSourceName, "result.0.use_ipv6_options"),
		resource.TestCheckResourceAttrPair(resourceName, "use_nextserver", dataSourceName, "result.0.use_nextserver"),
		resource.TestCheckResourceAttrPair(resourceName, "use_options", dataSourceName, "result.0.use_options"),
		resource.TestCheckResourceAttrPair(resourceName, "use_preferred_lifetime", dataSourceName, "result.0.use_preferred_lifetime"),
		resource.TestCheckResourceAttrPair(resourceName, "use_pxe_lease_time", dataSourceName, "result.0.use_pxe_lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "use_valid_lifetime", dataSourceName, "result.0.use_valid_lifetime"),
		resource.TestCheckResourceAttrPair(resourceName, "valid_lifetime", dataSourceName, "result.0.valid_lifetime"),
	}
}

func testAccRoaminghostDataSourceConfigFilters(name, mac, addressType, matchClient string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test" {
  name = %q
  mac = %q
  address_type = %q
  match_client = %q
}

data "nios_dhcp_roaminghost" "test" {
  filters = {
	name = nios_dhcp_roaminghost.test.name
  }
}
`, name, mac, addressType, matchClient)
}

func testAccRoaminghostDataSourceConfigExtAttrFilters(name, mac, addressType, matchClient, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test" {
  name = %q
  mac = %q
  address_type = %q
  match_client = %q
  extattrs = {
    Site = %q
  } 
}

data "nios_dhcp_roaminghost" "test" {
  extattrfilters = {
    Site = nios_dhcp_roaminghost.test.extattrs.Site
  }
}
`, name, mac, addressType, matchClient, extAttrsValue)
}
