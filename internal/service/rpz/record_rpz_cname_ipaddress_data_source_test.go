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

func TestAccRecordRpzCnameIpaddressDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_rpz_record_rpz_cname_ipaddress.test"
	resourceName := "nios_rpz_record_rpz_cname_ipaddress.test"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIp := "11.0.0.30"
	canonical := "11.0.0.30"
	view := acctest.RandomNameWithPrefix("view")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzCnameIpaddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzCnameIpaddressDataSourceConfigFilters(nameIp, canonical, rpZone, view),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordRpzCnameIpaddressResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordRpzCnameIpaddressDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_rpz_record_rpz_cname_ipaddress.test"
	resourceName := "nios_rpz_record_rpz_cname_ipaddress.test"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIp := "11.0.0.31"
	canonical := ""
	view := acctest.RandomNameWithPrefix("view")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzCnameIpaddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzCnameIpaddressDataSourceConfigExtAttrFilters(nameIp, canonical, rpZone, view),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordRpzCnameIpaddressResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordRpzCnameIpaddressResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
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

func testAccRecordRpzCnameIpaddressDataSourceConfigFilters(nameIp, canonical, rpZone, view string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_rpz_cname_ipaddress" "test" {
	name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
}

data "nios_rpz_record_rpz_cname_ipaddress" "test" {
  filters = {
	name = nios_rpz_record_rpz_cname_ipaddress.test.name
  }
}
`, nameIp, canonical)
	return strings.Join([]string{testAccBaseWithView(view), testAccBaseWithZoneNetworkZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameIpaddressDataSourceConfigExtAttrFilters(nameIp, canonical, rpZone, extAttrsValue string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_rpz_cname_ipaddress" "test" {
	name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
	extattrs = {
		Site = %q
	}
}

data "nios_rpz_record_rpz_cname_ipaddress" "test" {
  extattrfilters = {
	Site = nios_rpz_record_rpz_cname_ipaddress.test.extattrs.Site
  }
}
`, nameIp, canonical, extAttrsValue)
	return strings.Join([]string{testAccBaseWithZoneNetworkZone(rpZone, ""), config}, "")
}
