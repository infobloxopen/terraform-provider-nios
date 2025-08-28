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

func TestAccRecordCaaDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_caa.test"
	resourceName := "nios_dns_record_caa.test"
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-caa")
	var v dns.RecordCaa

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordCaaDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordCaaDataSourceConfigFilters(zoneFqdn, name, 0, "issue", "digicert.com"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordCaaResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordCaaDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_caa.test"
	resourceName := "nios_dns_record_caa.test"
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-caa")
	var v dns.RecordCaa

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordCaaDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordCaaDataSourceConfigExtAttrFilters(zoneFqdn, name, 0, "issue", "digicert.com", acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordCaaResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordCaaResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "ca_flag", dataSourceName, "result.0.ca_flag"),
		resource.TestCheckResourceAttrPair(resourceName, "ca_tag", dataSourceName, "result.0.ca_tag"),
		resource.TestCheckResourceAttrPair(resourceName, "ca_value", dataSourceName, "result.0.ca_value"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_info", dataSourceName, "result.0.cloud_info"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "creation_time", dataSourceName, "result.0.creation_time"),
		resource.TestCheckResourceAttrPair(resourceName, "creator", dataSourceName, "result.0.creator"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_principal", dataSourceName, "result.0.ddns_principal"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_protected", dataSourceName, "result.0.ddns_protected"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_name", dataSourceName, "result.0.dns_name"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "forbid_reclamation", dataSourceName, "result.0.forbid_reclamation"),
		resource.TestCheckResourceAttrPair(resourceName, "last_queried", dataSourceName, "result.0.last_queried"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "reclaimable", dataSourceName, "result.0.reclaimable"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordCaaDataSourceConfigFilters(zoneFqdn, name string, caFlag int, caTag, caValue string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_caa" "test" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    ca_flag = %d
    ca_tag = %q
    ca_value = %q
}

data "nios_dns_record_caa" "test" {
	filters = {
		name = nios_dns_record_caa.test.name
	}
}
`, name, caFlag, caTag, caValue)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordCaaDataSourceConfigExtAttrFilters(zoneFqdn, name string, caFlag int, caTag, caValue string, extAttrsValue string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_caa" "test" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
	ca_flag = %d
	ca_tag = %q
	ca_value = %q
	extattrs = {
		Site = %q
	} 
}

data "nios_dns_record_caa" "test" {
	extattrfilters = {
		Site = nios_dns_record_caa.test.extattrs.Site
	}
}
`, name, caFlag, caTag, caValue, extAttrsValue)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}
