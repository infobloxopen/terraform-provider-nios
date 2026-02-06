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

func TestAccSharedrecordMxDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_sharedrecord_mx.test"
	resourceName := "nios_dns_sharedrecord_mx.test"
	var v dns.SharedrecordMx
	name := acctest.RandomNameWithPrefix("sharedrecord-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharedrecordMxDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharedrecordMxDataSourceConfigFilters(mail_exchanger, name, 10, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					}, testAccCheckSharedrecordMxResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccSharedrecordMxDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_sharedrecord_mx.test"
	resourceName := "nios_dns_sharedrecord_mx.test"
	name := acctest.RandomNameWithPrefix("sharedrecord-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")
	extAttrValue := acctest.RandomName()

	var v dns.SharedrecordMx
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharedrecordMxDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharedrecordMxDataSourceConfigExtAttrFilters(mail_exchanger, name, 10, sharedRecordGroup, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					}, testAccCheckSharedrecordMxResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckSharedrecordMxResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_mail_exchanger", dataSourceName, "result.0.dns_mail_exchanger"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_name", dataSourceName, "result.0.dns_name"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "mail_exchanger", dataSourceName, "result.0.mail_exchanger"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "preference", dataSourceName, "result.0.preference"),
		resource.TestCheckResourceAttrPair(resourceName, "shared_record_group", dataSourceName, "result.0.shared_record_group"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
	}
}
func testAccSharedrecordMxDataSourceConfigFilters(mailExchanger, name string, preference int, sharedRecordGroup string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_mx" "test" {
	mail_exchanger     = %q
	name               = %q
  	preference          = %d
	shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
}

data "nios_dns_sharedrecord_mx" "test" {
	filters = {
    	name = nios_dns_sharedrecord_mx.test.name
  	}
}
`, mailExchanger, name, preference)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordMxDataSourceConfigExtAttrFilters(mailExchanger, name string, preference int, sharedRecordGroup, extAttrsValue string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_mx" "test" {
	mail_exchanger     = %q
	name               = %q
	preference         = %d
	shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
  	extattrs = {
   		Site = %q
  	}
}

data "nios_dns_sharedrecord_mx" "test" {
	extattrfilters = {
    	Site = nios_dns_sharedrecord_mx.test.extattrs.Site
	}
}
`, mailExchanger, name, preference, extAttrsValue)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}
