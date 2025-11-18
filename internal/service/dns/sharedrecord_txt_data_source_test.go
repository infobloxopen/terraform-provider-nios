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

func TestAccSharedrecordTxtDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_sharedrecord_txt.test"
	resourceName := "nios_dns_sharedrecord_txt.test"
	var v dns.SharedrecordTxt
	name := acctest.RandomNameWithPrefix("sharedrecord-txt-")
	text := "This is a shared record TXT record"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharedrecordTxtDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharedrecordTxtDataSourceConfigFilters(name, text, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSharedrecordTxtExists(context.Background(), resourceName, &v),
					}, testAccCheckSharedrecordTxtResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccSharedrecordTxtDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_sharedrecord_txt.test"
	resourceName := "nios_dns_sharedrecord_txt.test"
	name := acctest.RandomNameWithPrefix("sharedrecord-txt-")
	text := "This is a shared record TXT record"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")
	var v dns.SharedrecordTxt
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharedrecordTxtDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharedrecordTxtDataSourceConfigExtAttrFilters(name, sharedRecordGroup, text, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSharedrecordTxtExists(context.Background(), resourceName, &v),
					}, testAccCheckSharedrecordTxtResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckSharedrecordTxtResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_name", dataSourceName, "result.0.dns_name"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "shared_record_group", dataSourceName, "result.0.shared_record_group"),
		resource.TestCheckResourceAttrPair(resourceName, "text", dataSourceName, "result.0.text"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
	}
}

func testAccSharedrecordTxtDataSourceConfigFilters(name, sharedRecordGroup, text string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_txt" "test" {
  name = %q
  text = %q
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
}

data "nios_dns_sharedrecord_txt" "test" {
  filters = {
	 name = nios_dns_sharedrecord_txt.test.name
  }
}
`, name, text)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordTxtDataSourceConfigExtAttrFilters(name, sharedRecordGroup, text, extAttrsValue string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_txt" "test" {
  name = %q
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
  text = %q
  extattrs = {
    Site = %q
  } 
}

data "nios_dns_sharedrecord_txt" "test" {
  extattrfilters = {
	Site = nios_dns_sharedrecord_txt.test.extattrs.Site
  }
}
`, name, text, extAttrsValue)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}
