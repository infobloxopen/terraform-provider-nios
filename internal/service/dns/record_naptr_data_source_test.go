package dns_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccRecordNaptrDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_naptr.test"
	resourceName := "nios_dns_record_naptr.test"
	var v dns.RecordNaptr
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("test-naptr")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordNaptrDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordNaptrDataSourceConfigFilters(zoneFqdn, name, 10, 10, "."),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordNaptrResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordNaptrDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_naptr.test"
	resourceName := "nios_dns_record_naptr.test"
	var v dns.RecordNaptr
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("test-naptr")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordNaptrDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordNaptrDataSourceConfigExtAttrFilters(zoneFqdn, name, 10, 10, ".", acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordNaptrResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordNaptrResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_info", dataSourceName, "result.0.cloud_info"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "creation_time", dataSourceName, "result.0.creation_time"),
		resource.TestCheckResourceAttrPair(resourceName, "creator", dataSourceName, "result.0.creator"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_principal", dataSourceName, "result.0.ddns_principal"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_protected", dataSourceName, "result.0.ddns_protected"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_name", dataSourceName, "result.0.dns_name"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_replacement", dataSourceName, "result.0.dns_replacement"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "flags", dataSourceName, "result.0.flags"),
		resource.TestCheckResourceAttrPair(resourceName, "forbid_reclamation", dataSourceName, "result.0.forbid_reclamation"),
		resource.TestCheckResourceAttrPair(resourceName, "last_queried", dataSourceName, "result.0.last_queried"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "order", dataSourceName, "result.0.order"),
		resource.TestCheckResourceAttrPair(resourceName, "preference", dataSourceName, "result.0.preference"),
		resource.TestCheckResourceAttrPair(resourceName, "reclaimable", dataSourceName, "result.0.reclaimable"),
		resource.TestCheckResourceAttrPair(resourceName, "regexp", dataSourceName, "result.0.regexp"),
		resource.TestCheckResourceAttrPair(resourceName, "replacement", dataSourceName, "result.0.replacement"),
		resource.TestCheckResourceAttrPair(resourceName, "services", dataSourceName, "result.0.services"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordNaptrDataSourceConfigFilters(zoneFqdn, name string, order, preference int, replacement string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test" {
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
 	order = %d
    preference = %d
    replacement = %q
}

data "nios_dns_record_naptr" "test" {
	filters = {
		name = nios_dns_record_naptr.test.name
	}
}
`, name, order, preference, replacement)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrDataSourceConfigExtAttrFilters(zoneFqdn, name string, order, preference int, replacement string, extAttrsValue string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
	order = %d
	preference = %d
	replacement = %q
	extattrs = {
		Site = %q
	}
}

data "nios_dns_record_naptr" "test" {
	extattrfilters = {
		Site = nios_dns_record_naptr.test.extattrs.Site
	}
}
`, name, order, preference, replacement, extAttrsValue)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")

}
