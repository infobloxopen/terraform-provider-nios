package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccRecordPtrDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_ptr.test"
	resourceName := "nios_dns_record_ptr.test"
	var v dns.RecordPtr
	ptrDName := acctest.RandomNameWithPrefix("ptr") + ".example.com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordPtrDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordPtrDataSourceConfigFilters("192.168.10.22", ptrDName, "default"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordPtrResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordPtrDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_ptr.test"
	resourceName := "nios_dns_record_ptr.test"
	var v dns.RecordPtr
	ptrDName := acctest.RandomNameWithPrefix("ptr") + ".example.com"
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordPtrDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordPtrDataSourceConfigExtAttrFilters("192.168.10.22", ptrDName, "default", extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordPtrResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordPtrResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
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
		resource.TestCheckResourceAttrPair(resourceName, "discovered_data", dataSourceName, "result.0.discovered_data"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_name", dataSourceName, "result.0.dns_name"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_ptrdname", dataSourceName, "result.0.dns_ptrdname"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "forbid_reclamation", dataSourceName, "result.0.forbid_reclamation"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv4addr", dataSourceName, "result.0.ipv4addr"),
		resource.TestCheckResourceAttrPair(resourceName, "func_call", dataSourceName, "result.0.func_call"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6addr", dataSourceName, "result.0.ipv6addr"),
		resource.TestCheckResourceAttrPair(resourceName, "last_queried", dataSourceName, "result.0.last_queried"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_ad_user_data", dataSourceName, "result.0.ms_ad_user_data"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "ptrdname", dataSourceName, "result.0.ptrdname"),
		resource.TestCheckResourceAttrPair(resourceName, "reclaimable", dataSourceName, "result.0.reclaimable"),
		resource.TestCheckResourceAttrPair(resourceName, "shared_record_group", dataSourceName, "result.0.shared_record_group"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordPtrDataSourceConfigFilters(ipv4addr, ptrdname, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test" {
	ipv4addr = %q
	ptrdname = %q
	view = %q
}

data "nios_dns_record_ptr" "test" {
  filters = {
	ipv4addr = nios_dns_record_ptr.test.ipv4addr
	ptrdname = nios_dns_record_ptr.test.ptrdname
	view = nios_dns_record_ptr.test.view
  }
}
`, ipv4addr, ptrdname, view)
}

func testAccRecordPtrDataSourceConfigExtAttrFilters(ipv4addr, ptrdname, view, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test" {
	ipv4addr = %q
	ptrdname = %q
	view = %q
	extattrs = {
		Site =  %q
	}
}

data "nios_dns_record_ptr" "test" {
  extattrfilters = {
	Site = nios_dns_record_ptr.test.extattrs.Site
  }
}
`, ipv4addr, ptrdname, view, extAttrsValue)
}
