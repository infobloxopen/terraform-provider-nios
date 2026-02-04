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

func TestAccRecordTlsaDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_tlsa.test"
	resourceName := "nios_dns_record_tlsa.test"
	var v dns.RecordTlsa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-tlsa")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordTlsaDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordTlsaDataSourceConfigFilters(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordTlsaResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordTlsaDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_tlsa.test"
	resourceName := "nios_dns_record_tlsa.test"
	var v dns.RecordTlsa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-tlsa")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordTlsaDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordTlsaDataSourceConfigExtAttrFilters(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordTlsaResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordTlsaResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "certificate_data", dataSourceName, "result.0.certificate_data"),
		resource.TestCheckResourceAttrPair(resourceName, "certificate_usage", dataSourceName, "result.0.certificate_usage"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_info", dataSourceName, "result.0.cloud_info"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "creator", dataSourceName, "result.0.creator"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_name", dataSourceName, "result.0.dns_name"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "last_queried", dataSourceName, "result.0.last_queried"),
		resource.TestCheckResourceAttrPair(resourceName, "matched_type", dataSourceName, "result.0.matched_type"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "selector", dataSourceName, "result.0.selector"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordTlsaDataSourceConfigFilters(zoneFqdn, name, certificateData string, certificateUsage, matchedType, selector int) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_tlsa" "test" {
	certificate_data = %q
    certificate_usage = %d
    matched_type = %d
    selector = %d
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
}

data "nios_dns_record_tlsa" "test" {
	filters = {
		name = nios_dns_record_tlsa.test.name
	}
}
`, certificateData, certificateUsage, matchedType, selector, name)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordTlsaDataSourceConfigExtAttrFilters(zoneFqdn, name, certificateData string, certificateUsage, matchedType, selector int, extAttrsValue string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_tlsa" "test" {
	certificate_data = %q
	certificate_usage = %d
	matched_type = %d
	selector = %d
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
	extattrs = {
		Site = %q
	} 
}

data "nios_dns_record_tlsa" "test" {
	extattrfilters = {
		Site = nios_dns_record_tlsa.test.extattrs.Site
	}
}
`, certificateData, certificateUsage, matchedType, selector, name, extAttrsValue)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}
