package dhcp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccIpv6fixedaddresstemplateDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_ipv6fixedaddresstemplate.test"
	resourceName := "nios_dhcp_ipv6fixedaddresstemplate.test"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6fixedaddresstemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6fixedaddresstemplateDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					}, testAccCheckIpv6fixedaddresstemplateResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccIpv6fixedaddresstemplateDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_ipv6fixedaddresstemplate.test"
	resourceName := "nios_dhcp_ipv6fixedaddresstemplate.test"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6fixedaddresstemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6fixedaddresstemplateDataSourceConfigExtAttrFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					}, testAccCheckIpv6fixedaddresstemplateResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckIpv6fixedaddresstemplateResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "domain_name", dataSourceName, "result.0.domain_name"),
		resource.TestCheckResourceAttrPair(resourceName, "domain_name_servers", dataSourceName, "result.0.domain_name_servers"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "logic_filter_rules", dataSourceName, "result.0.logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "number_of_addresses", dataSourceName, "result.0.number_of_addresses"),
		resource.TestCheckResourceAttrPair(resourceName, "offset", dataSourceName, "result.0.offset"),
		resource.TestCheckResourceAttrPair(resourceName, "options", dataSourceName, "result.0.options"),
		resource.TestCheckResourceAttrPair(resourceName, "preferred_lifetime", dataSourceName, "result.0.preferred_lifetime"),
		resource.TestCheckResourceAttrPair(resourceName, "use_domain_name", dataSourceName, "result.0.use_domain_name"),
		resource.TestCheckResourceAttrPair(resourceName, "use_domain_name_servers", dataSourceName, "result.0.use_domain_name_servers"),
		resource.TestCheckResourceAttrPair(resourceName, "use_logic_filter_rules", dataSourceName, "result.0.use_logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "use_options", dataSourceName, "result.0.use_options"),
		resource.TestCheckResourceAttrPair(resourceName, "use_preferred_lifetime", dataSourceName, "result.0.use_preferred_lifetime"),
		resource.TestCheckResourceAttrPair(resourceName, "use_valid_lifetime", dataSourceName, "result.0.use_valid_lifetime"),
		resource.TestCheckResourceAttrPair(resourceName, "valid_lifetime", dataSourceName, "result.0.valid_lifetime"),
	}
}

func testAccIpv6fixedaddresstemplateDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test" {
  name = %q
}

data "nios_dhcp_ipv6fixedaddresstemplate" "test" {
  filters = {
	name = nios_dhcp_ipv6fixedaddresstemplate.test.name
  }
}
`, name)
}

func testAccIpv6fixedaddresstemplateDataSourceConfigExtAttrFilters(name string, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test" {
  name = %q
  extattrs = {
    Site = %q
  } 
}

data "nios_dhcp_ipv6fixedaddresstemplate" "test" {
  extattrfilters = {
	Site = nios_dhcp_ipv6fixedaddresstemplate.test.extattrs.Site
  }
}
`, name, extAttrsValue)
}
