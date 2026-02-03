package rpz_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/rpz"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccRecordRpzCnameClientipaddressDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_rpz_record_cname_clientipaddress.test"
	resourceName := "nios_rpz_record_cname_clientipaddress.test"
	var v rpz.RecordRpzCnameClientipaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP := "12.0.0.40"
	canonical := "rpz-passthru"
	view := acctest.RandomNameWithPrefix("view")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzCnameClientipaddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzCnameClientipaddressDataSourceConfigFilters(nameIP, rpZone, canonical, view),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordRpzCnameClientipaddressExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordRpzCnameClientipaddressResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordRpzCnameClientipaddressDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_rpz_record_cname_clientipaddress.test"
	resourceName := "nios_rpz_record_cname_clientipaddress.test"
	var v rpz.RecordRpzCnameClientipaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP := "12.0.0.41"
	canonical := "rpz-passthru"
	view := acctest.RandomNameWithPrefix("view")
	site := view

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzCnameClientipaddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzCnameClientipaddressDataSourceConfigExtAttrFilters(nameIP, rpZone, canonical, view, site),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordRpzCnameClientipaddressExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordRpzCnameClientipaddressResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordRpzCnameClientipaddressResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "canonical", dataSourceName, "result.0.canonical"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "is_ipv4", dataSourceName, "result.0.is_ipv4"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "rp_zone", dataSourceName, "result.0.rp_zone"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordRpzCnameClientipaddressDataSourceConfigFilters(nameIP, rpZone, canonical, view string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddress" "test" {
	name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	view = nios_dns_view.custom_view.name
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
}

data "nios_rpz_record_cname_clientipaddress" "test" {
  filters = {
	name = nios_rpz_record_cname_clientipaddress.test.name
  }
}`, nameIP, canonical)
	return strings.Join([]string{testAccBaseWithView(view), testAccBaseWithZoneRPNetwork(rpZone, "nios_dns_view.custom_view.name"), config}, "")
}

func testAccRecordRpzCnameClientipaddressDataSourceConfigExtAttrFilters(nameIP, rpZone, canonical, view, site string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddress" "test" {
	name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
	view = nios_dns_view.custom_view.name
	extattrs = {
		Site = %q
	}
}

data "nios_rpz_record_cname_clientipaddress" "test" {
  extattrfilters = {
	Site = nios_rpz_record_cname_clientipaddress.test.extattrs.Site
  }
}
`, nameIP, canonical, site)
	return strings.Join([]string{testAccBaseWithView(view), testAccBaseWithZoneRPNetwork(rpZone, "nios_dns_view.custom_view.name"), config}, "")
}
