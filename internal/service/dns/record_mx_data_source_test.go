package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccRecordMxDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_mx.test"
	resourceName := "nios_dns_record_mx.test"
	var v dns.RecordMx
	name := acctest.RandomNameWithPrefix("record-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordMxDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordMxDataSourceConfigFilters(name, mail_exchanger, 10, "default"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordMxResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordMxDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_mx.test"
	resourceName := "nios_dns_record_mx.test"
	var v dns.RecordMx
	name := acctest.RandomNameWithPrefix("record-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordMxDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordMxDataSourceConfigExtAttrFilters(name, mail_exchanger, 10, "default", extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordMxResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordMxResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
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
		resource.TestCheckResourceAttrPair(resourceName, "dns_mail_exchanger", dataSourceName, "result.0.dns_mail_exchanger"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_name", dataSourceName, "result.0.dns_name"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "forbid_reclamation", dataSourceName, "result.0.forbid_reclamation"),
		resource.TestCheckResourceAttrPair(resourceName, "last_queried", dataSourceName, "result.0.last_queried"),
		resource.TestCheckResourceAttrPair(resourceName, "mail_exchanger", dataSourceName, "result.0.mail_exchanger"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "preference", dataSourceName, "result.0.preference"),
		resource.TestCheckResourceAttrPair(resourceName, "reclaimable", dataSourceName, "result.0.reclaimable"),
		resource.TestCheckResourceAttrPair(resourceName, "shared_record_group", dataSourceName, "result.0.shared_record_group"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordMxDataSourceConfigFilters(name, mail_exchanger string, preference int64, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_mx" "test" {
    name            = %q
    mail_exchanger  = %q
    preference      = %d
    view            = %q
}

data "nios_dns_record_mx" "test" {
  filters = {
	name = nios_dns_record_mx.test.name
  }
}
`, name, mail_exchanger, preference, view)
}

func testAccRecordMxDataSourceConfigExtAttrFilters(name, mail_exchanger string, preference int64, view string, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_mx" "test" {
	name            = %q
	mail_exchanger  = %q
	preference      = %d
	view            = %q
	extattrs = {
		Site = %q
	} 
}

data "nios_dns_record_mx" "test" {
  extattrfilters = {
	Site = nios_dns_record_mx.test.extattrs.Site
  }
}
`, name, mail_exchanger, preference, view, extAttrsValue)
}
