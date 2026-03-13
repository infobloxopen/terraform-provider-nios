package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccRecordTxtDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_txt.test"
	resourceName := "nios_dns_record_txt.test"
	var v dns.RecordTxt
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordTxtDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordTxtDataSourceConfigFilters(name, "record text"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordTxtResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordTxtDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_txt.test"
	resourceName := "nios_dns_record_txt.test"
	var v dns.RecordTxt
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordTxtDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordTxtDataSourceConfigExtAttrFilters(name, "record text", extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordTxtResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordTxtResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
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
		resource.TestCheckResourceAttrPair(resourceName, "dns_name", dataSourceName, "result.0.dns_name"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "forbid_reclamation", dataSourceName, "result.0.forbid_reclamation"),
		resource.TestCheckResourceAttrPair(resourceName, "last_queried", dataSourceName, "result.0.last_queried"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "reclaimable", dataSourceName, "result.0.reclaimable"),
		resource.TestCheckResourceAttrPair(resourceName, "shared_record_group", dataSourceName, "result.0.shared_record_group"),
		resource.TestCheckResourceAttrPair(resourceName, "text", dataSourceName, "result.0.text"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordTxtDataSourceConfigFilters(name, text string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_txt" "test" {
	name = %q
	text = %q
}

data "nios_dns_record_txt" "test" {
  filters = {
	name  = nios_dns_record_txt.test.name
  }
}
`, name, text)
}

func testAccRecordTxtDataSourceConfigExtAttrFilters(name, text, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_txt" "test" {
	name = %q
	text = %q  
	extattrs = {
    	Site = %q
  } 
}

data "nios_dns_record_txt" "test" {
  extattrfilters = {
	Site = nios_dns_record_txt.test.extattrs.Site
  }
}
`, name, text, extAttrsValue)
}
