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

func TestAccRecordRpzCnameIpaddressdnDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_rpz_record_cname_ipaddressdn.test"
	resourceName := "nios_rpz_record_cname_ipaddressdn.test"
	var v rpz.RecordRpzCnameIpaddressdn
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone
	canonical := "test-cname"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzCnameIpaddressdnDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzCnameIpaddressdnDataSourceConfigFilters(name, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordRpzCnameIpaddressdnExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordRpzCnameIpaddressdnResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordRpzCnameIpaddressdnDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_rpz_record_cname_ipaddressdn.test"
	resourceName := "nios_rpz_record_cname_ipaddressdn.test"
	var v rpz.RecordRpzCnameIpaddressdn
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone
	canonical := "test-cname"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzCnameIpaddressdnDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzCnameIpaddressdnDataSourceConfigExtAttrFilters(name, canonical, rpZone, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordRpzCnameIpaddressdnExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordRpzCnameIpaddressdnResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordRpzCnameIpaddressdnResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
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

func testAccRecordRpzCnameIpaddressdnDataSourceConfigFilters(name, canonical, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_ipaddressdn" "test" {
	name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}

data "nios_rpz_record_cname_ipaddressdn" "test" {
  filters = {
	name = nios_rpz_record_cname_ipaddressdn.test.name
  }
}
`, name, canonical)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameIpaddressdnDataSourceConfigExtAttrFilters(name, canonical, rpZone, extAttrsValue string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_ipaddressdn" "test" {
	name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	extattrs = {
		Site = %q
	} 
}

data "nios_rpz_record_cname_ipaddressdn" "test" {
  extattrfilters = {
	Site = nios_rpz_record_cname_ipaddressdn.test.extattrs.Site
  }
}
`, name, canonical, extAttrsValue)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}
