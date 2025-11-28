package dhcp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccFixedaddresstemplateDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_fixedaddresstemplate.test"
	resourceName := "nios_dhcp_fixedaddresstemplate.test"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFixedaddresstemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFixedaddresstemplateDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					}, testAccCheckFixedaddresstemplateResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccFixedaddresstemplateDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_fixedaddresstemplate.test"
	resourceName := "nios_dhcp_fixedaddresstemplate.test"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFixedaddresstemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFixedaddresstemplateDataSourceConfigExtAttrFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					}, testAccCheckFixedaddresstemplateResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckFixedaddresstemplateResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "bootfile", dataSourceName, "result.0.bootfile"),
		resource.TestCheckResourceAttrPair(resourceName, "bootserver", dataSourceName, "result.0.bootserver"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_domainname", dataSourceName, "result.0.ddns_domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_hostname", dataSourceName, "result.0.ddns_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "deny_bootp", dataSourceName, "result.0.deny_bootp"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_ddns", dataSourceName, "result.0.enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_pxe_lease_time", dataSourceName, "result.0.enable_pxe_lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "ignore_dhcp_option_list_request", dataSourceName, "result.0.ignore_dhcp_option_list_request"),
		resource.TestCheckResourceAttrPair(resourceName, "logic_filter_rules", dataSourceName, "result.0.logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "nextserver", dataSourceName, "result.0.nextserver"),
		resource.TestCheckResourceAttrPair(resourceName, "number_of_addresses", dataSourceName, "result.0.number_of_addresses"),
		resource.TestCheckResourceAttrPair(resourceName, "offset", dataSourceName, "result.0.offset"),
		resource.TestCheckResourceAttrPair(resourceName, "options", dataSourceName, "result.0.options"),
		resource.TestCheckResourceAttrPair(resourceName, "pxe_lease_time", dataSourceName, "result.0.pxe_lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "use_bootfile", dataSourceName, "result.0.use_bootfile"),
		resource.TestCheckResourceAttrPair(resourceName, "use_bootserver", dataSourceName, "result.0.use_bootserver"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_domainname", dataSourceName, "result.0.use_ddns_domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "use_deny_bootp", dataSourceName, "result.0.use_deny_bootp"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_ddns", dataSourceName, "result.0.use_enable_ddns"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ignore_dhcp_option_list_request", dataSourceName, "result.0.use_ignore_dhcp_option_list_request"),
		resource.TestCheckResourceAttrPair(resourceName, "use_logic_filter_rules", dataSourceName, "result.0.use_logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "use_nextserver", dataSourceName, "result.0.use_nextserver"),
		resource.TestCheckResourceAttrPair(resourceName, "use_options", dataSourceName, "result.0.use_options"),
		resource.TestCheckResourceAttrPair(resourceName, "use_pxe_lease_time", dataSourceName, "result.0.use_pxe_lease_time"),
	}
}

func testAccFixedaddresstemplateDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test" {
  name = %q
}

data "nios_dhcp_fixedaddresstemplate" "test" {
  filters = {
	name = nios_dhcp_fixedaddresstemplate.test.name
  }
}
`, name)
}

func testAccFixedaddresstemplateDataSourceConfigExtAttrFilters(name string, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test" {
  name = %q
  extattrs = {
    Site = %q
  } 
}

data "nios_dhcp_fixedaddresstemplate" "test" {
  extattrfilters = {
	Site = nios_dhcp_fixedaddresstemplate.test.extattrs.Site
  }
}
`, name, extAttrsValue)
}
