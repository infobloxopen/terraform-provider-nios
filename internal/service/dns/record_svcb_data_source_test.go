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

func TestAccRecordSvcbDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_svcb.test"
	resourceName := "nios_dns_record_svcb.test"
	var v dns.RecordSvcb
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-svcb")
	priority := "10"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordSvcbDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordSvcbDataSourceConfigFilters(zoneFqdn, name, priority),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordSvcbResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordSvcbDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_svcb.test"
	resourceName := "nios_dns_record_svcb.test"
	var v dns.RecordSvcb
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-svcb")
	priority := "10"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordSvcbDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordSvcbDataSourceConfigExtAttrFilters(zoneFqdn, name, priority, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordSvcbResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordSvcbResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "aws_rte53_record_info", dataSourceName, "result.0.aws_rte53_record_info"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_info", dataSourceName, "result.0.cloud_info"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "creation_time", dataSourceName, "result.0.creation_time"),
		resource.TestCheckResourceAttrPair(resourceName, "creator", dataSourceName, "result.0.creator"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_principal", dataSourceName, "result.0.ddns_principal"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_protected", dataSourceName, "result.0.ddns_protected"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "forbid_reclamation", dataSourceName, "result.0.forbid_reclamation"),
		resource.TestCheckResourceAttrPair(resourceName, "last_queried", dataSourceName, "result.0.last_queried"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "priority", dataSourceName, "result.0.priority"),
		resource.TestCheckResourceAttrPair(resourceName, "reclaimable", dataSourceName, "result.0.reclaimable"),
		resource.TestCheckResourceAttrPair(resourceName, "svc_parameters", dataSourceName, "result.0.svc_parameters"),
		resource.TestCheckResourceAttrPair(resourceName, "target_name", dataSourceName, "result.0.target_name"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordSvcbDataSourceConfigFilters(zoneFqdn, name, priority string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_svcb" "test" {
	name = "%s.%s"
	target_name = nios_dns_zone_auth.test.fqdn
	priority = %q
}

data "nios_dns_record_svcb" "test" {
	filters = {
		name = nios_dns_record_svcb.test.name
  	}
}
`, name, zoneFqdn, priority)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordSvcbDataSourceConfigExtAttrFilters(zoneFqdn, name, priority, extAttrsValue string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_svcb" "test" {
	name = "%s.%s"
	target_name = nios_dns_zone_auth.test.fqdn
	priority = %q
	extattrs = {
		Site = %q
	} 
}

data "nios_dns_record_svcb" "test" {
	extattrfilters = {
		Site = nios_dns_record_svcb.test.extattrs.Site
	}
}
`, name, zoneFqdn, priority, extAttrsValue)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}
