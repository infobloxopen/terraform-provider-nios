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

func TestAccRecordDnameDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_dname.test"
	resourceName := "nios_dns_record_dname.test"
	var v dns.RecordDname
	target := acctest.RandomNameWithPrefix("test-dname") + ".com"
	view := acctest.RandomNameWithPrefix("test-view")
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordDnameDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordDnameDataSourceConfigFilters(target, view, zoneFqdn),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordDnameResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordDnameDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_dname.test"
	resourceName := "nios_dns_record_dname.test"
	var v dns.RecordDname
	target := acctest.RandomNameWithPrefix("test-dname") + ".com"
	view := acctest.RandomNameWithPrefix("test-view")
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordDnameDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordDnameDataSourceConfigExtAttrFilters(target, view, zoneFqdn, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordDnameResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordDnameResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
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
		resource.TestCheckResourceAttrPair(resourceName, "dns_target", dataSourceName, "result.0.dns_target"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "forbid_reclamation", dataSourceName, "result.0.forbid_reclamation"),
		resource.TestCheckResourceAttrPair(resourceName, "last_queried", dataSourceName, "result.0.last_queried"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "reclaimable", dataSourceName, "result.0.reclaimable"),
		resource.TestCheckResourceAttrPair(resourceName, "shared_record_group", dataSourceName, "result.0.shared_record_group"),
		resource.TestCheckResourceAttrPair(resourceName, "target", dataSourceName, "result.0.target"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordDnameDataSourceConfigFilters(target, view, zoneFqdn string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_dname" "test" {
	name = nios_dns_zone_auth.test.fqdn
	target = %q
	view = nios_dns_zone_auth.test.view
}

data "nios_dns_record_dname" "test" {
  filters = {
    name = nios_dns_record_dname.test.name
    target = nios_dns_record_dname.test.target
    view = nios_dns_record_dname.test.view
  }
}
`, target)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordDnameDataSourceConfigExtAttrFilters(target, view, zoneFqdn, extAttrsValue string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_dname" "test" {
	name = nios_dns_zone_auth.test.fqdn
	target = %q
	view = nios_dns_zone_auth.test.view
    extattrs = {
    	Site = %q
  } 
}

data "nios_dns_record_dname" "test" {
	extattrfilters = {
		Site = nios_dns_record_dname.test.extattrs.Site
	}
}
`, target, extAttrsValue)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}
