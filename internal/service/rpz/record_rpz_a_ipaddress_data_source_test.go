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

func TestAccRecordRpzAIpaddressDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_rpz_record_a_ipaddress.test"
	resourceName := "nios_rpz_record_a_ipaddress.test"
	var v rpz.RecordRpzAIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzAIpaddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzAIpaddressDataSourceConfigFilters(name, "10.10.0.1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordRpzAIpaddressExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordRpzAIpaddressResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordRpzAIpaddressDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_rpz_record_a_ipaddress.test"
	resourceName := "nios_rpz_record_a_ipaddress.test"
	var v rpz.RecordRpzAIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzAIpaddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzAIpaddressDataSourceConfigExtAttrFilters(name, "10.10.0.1", rpZone, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordRpzAIpaddressExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordRpzAIpaddressResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordRpzAIpaddressResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv4addr", dataSourceName, "result.0.ipv4addr"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "rp_zone", dataSourceName, "result.0.rp_zone"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordRpzAIpaddressDataSourceConfigFilters(name, ipV4Addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_a_ipaddress" "test" {
	name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}

data "nios_rpz_record_a_ipaddress" "test" {
  filters = {
	name = nios_rpz_record_a_ipaddress.test.name
  }
}
`, name, ipV4Addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAIpaddressDataSourceConfigExtAttrFilters(name, ipV4Addr, rpZone, extAttrsValue string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_a_ipaddress" "test" {
	name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	extattrs = {
		Site = %q
	} 
}

data "nios_rpz_record_a_ipaddress" "test" {
  extattrfilters = {
	Site = nios_rpz_record_a_ipaddress.test.extattrs.Site
  }
}
`, name, ipV4Addr, extAttrsValue)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}
