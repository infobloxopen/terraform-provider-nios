package ipam_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccIpv6networktemplateDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_ipam_ipv6networktemplate.test"
	resourceName := "nios_ipam_ipv6networktemplate.test"
	var v ipam.Ipv6networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6networktemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6networktemplateDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					}, testAccCheckIpv6networktemplateResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccIpv6networktemplateDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_ipam_ipv6networktemplate.test"
	resourceName := "nios_ipam_ipv6networktemplate.test"
	var v ipam.Ipv6networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6networktemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6networktemplateDataSourceConfigExtAttrFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					}, testAccCheckIpv6networktemplateResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckIpv6networktemplateResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "allow_any_netmask", dataSourceName, "result.0.allow_any_netmask"),
		resource.TestCheckResourceAttrPair(resourceName, "auto_create_reversezone", dataSourceName, "result.0.auto_create_reversezone"),
		resource.TestCheckResourceAttrPair(resourceName, "cidr", dataSourceName, "result.0.cidr"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_api_compatible", dataSourceName, "result.0.cloud_api_compatible"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_domainname", dataSourceName, "result.0.ddns_domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_enable_option_fqdn", dataSourceName, "result.0.ddns_enable_option_fqdn"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_generate_hostname", dataSourceName, "result.0.ddns_generate_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_server_always_updates", dataSourceName, "result.0.ddns_server_always_updates"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_ttl", dataSourceName, "result.0.ddns_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "delegated_member", dataSourceName, "result.0.delegated_member"),
		resource.TestCheckResourceAttrPair(resourceName, "domain_name", dataSourceName, "result.0.domain_name"),
		resource.TestCheckResourceAttrPair(resourceName, "domain_name_servers", dataSourceName, "result.0.domain_name_servers"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_ddns", dataSourceName, "result.0.enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "fixed_address_templates", dataSourceName, "result.0.fixed_address_templates"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6prefix", dataSourceName, "result.0.ipv6prefix"),
		resource.TestCheckResourceAttrPair(resourceName, "logic_filter_rules", dataSourceName, "result.0.logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "members", dataSourceName, "result.0.members"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "options", dataSourceName, "result.0.options"),
		resource.TestCheckResourceAttrPair(resourceName, "preferred_lifetime", dataSourceName, "result.0.preferred_lifetime"),
		resource.TestCheckResourceAttrPair(resourceName, "range_templates", dataSourceName, "result.0.range_templates"),
		resource.TestCheckResourceAttrPair(resourceName, "recycle_leases", dataSourceName, "result.0.recycle_leases"),
		resource.TestCheckResourceAttrPair(resourceName, "rir", dataSourceName, "result.0.rir"),
		resource.TestCheckResourceAttrPair(resourceName, "rir_organization", dataSourceName, "result.0.rir_organization"),
		resource.TestCheckResourceAttrPair(resourceName, "rir_registration_action", dataSourceName, "result.0.rir_registration_action"),
		resource.TestCheckResourceAttrPair(resourceName, "rir_registration_status", dataSourceName, "result.0.rir_registration_status"),
		resource.TestCheckResourceAttrPair(resourceName, "send_rir_request", dataSourceName, "result.0.send_rir_request"),
		resource.TestCheckResourceAttrPair(resourceName, "update_dns_on_lease_renewal", dataSourceName, "result.0.update_dns_on_lease_renewal"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_domainname", dataSourceName, "result.0.use_ddns_domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_enable_option_fqdn", dataSourceName, "result.0.use_ddns_enable_option_fqdn"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_generate_hostname", dataSourceName, "result.0.use_ddns_generate_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_ttl", dataSourceName, "result.0.use_ddns_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_domain_name", dataSourceName, "result.0.use_domain_name"),
		resource.TestCheckResourceAttrPair(resourceName, "use_domain_name_servers", dataSourceName, "result.0.use_domain_name_servers"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_ddns", dataSourceName, "result.0.use_enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "use_logic_filter_rules", dataSourceName, "result.0.use_logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "use_options", dataSourceName, "result.0.use_options"),
		resource.TestCheckResourceAttrPair(resourceName, "use_preferred_lifetime", dataSourceName, "result.0.use_preferred_lifetime"),
		resource.TestCheckResourceAttrPair(resourceName, "use_recycle_leases", dataSourceName, "result.0.use_recycle_leases"),
		resource.TestCheckResourceAttrPair(resourceName, "use_update_dns_on_lease_renewal", dataSourceName, "result.0.use_update_dns_on_lease_renewal"),
		resource.TestCheckResourceAttrPair(resourceName, "use_valid_lifetime", dataSourceName, "result.0.use_valid_lifetime"),
		resource.TestCheckResourceAttrPair(resourceName, "valid_lifetime", dataSourceName, "result.0.valid_lifetime"),
	}
}

func testAccIpv6networktemplateDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test" {
  name = %q
}

data "nios_ipam_ipv6networktemplate" "test" {
  filters = {
    name = nios_ipam_ipv6networktemplate.test.name
  }
}
`, name)
}

func testAccIpv6networktemplateDataSourceConfigExtAttrFilters(name, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test" {
  name = %q
  extattrs = {
    Site = %q
  } 
}

data "nios_ipam_ipv6networktemplate" "test" {
  extattrfilters = {
    Site = nios_ipam_ipv6networktemplate.test.extattrs.Site
  }
}
`, name, extAttrsValue)
}
