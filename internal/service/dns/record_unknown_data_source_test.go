package dns_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

func TestAccRecordUnknownDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_unknown.test"
	resourceName := "nios_dns_record_unknown.test"
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-unknown")
	subfieldValues := []map[string]any{
		{
			"field_type":     "T",
			"field_value":    "example-text",
			"include_length": "8_BIT",
		},
	}
	var v dns.RecordUnknown

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordUnknownDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordUnknownDataSourceConfigFilters(zoneFqdn, name, "SPF", subfieldValues),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordUnknownResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordUnknownDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_unknown.test"
	resourceName := "nios_dns_record_unknown.test"
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-unknown")
	subfieldValues := []map[string]any{
		{
			"field_type":     "T",
			"field_value":    "example-text",
			"include_length": "8_BIT",
		},
	}
	var v dns.RecordUnknown

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordUnknownDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordUnknownDataSourceConfigExtAttrFilters(zoneFqdn, name, "SPF", subfieldValues, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordUnknownResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordUnknownResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_info", dataSourceName, "result.0.cloud_info"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "creator", dataSourceName, "result.0.creator"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "display_rdata", dataSourceName, "result.0.display_rdata"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_name", dataSourceName, "result.0.dns_name"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_host_name_policy", dataSourceName, "result.0.enable_host_name_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "last_queried", dataSourceName, "result.0.last_queried"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "policy", dataSourceName, "result.0.policy"),
		resource.TestCheckResourceAttrPair(resourceName, "record_type", dataSourceName, "result.0.record_type"),
		resource.TestCheckResourceAttrPair(resourceName, "subfield_values", dataSourceName, "result.0.subfield_values"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordUnknownDataSourceConfigFilters(zoneFqdn, name, recordType string, subfieldValues []map[string]any) string {
	subfieldValuesHCL := utils.ConvertSliceOfMapsToHCL(subfieldValues)
	config := fmt.Sprintf(`
resource "nios_dns_record_unknown" "test" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    record_type = %q
    subfield_values = %s
}

data "nios_dns_record_unknown" "test" {
	filters = {
		name = nios_dns_record_unknown.test.name
	}
}`, name, recordType, subfieldValuesHCL)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordUnknownDataSourceConfigExtAttrFilters(zoneFqdn, name, recordType string, subfieldValues []map[string]any, extAttrsValue string) string {
	subfieldValuesHCL := utils.ConvertSliceOfMapsToHCL(subfieldValues)
	config := fmt.Sprintf(`
resource "nios_dns_record_unknown" "test" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
	record_type = %q
	subfield_values = %s
	extattrs = {
		Site = %q
	} 
}

data "nios_dns_record_unknown" "test" {
	extattrfilters = {
		Site = nios_dns_record_unknown.test.extattrs.Site
	}
}
`, name, recordType, subfieldValuesHCL, extAttrsValue)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}
