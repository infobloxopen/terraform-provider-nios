package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccRecordSrvDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_srv.test"
	resourceName := "nios_dns_record_srv.test"
	var v dns.RecordSrv
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordSrvDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordSrvDataSourceConfigFilters(name, target, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordSrvResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordSrvDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_srv.test"
	resourceName := "nios_dns_record_srv.test"
	var v dns.RecordSrv
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordSrvDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordSrvDataSourceConfigExtAttrFilters(name, target, 80, 10, 360, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordSrvResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordSrvResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "aws_rte53_record_info", dataSourceName, "result.0.aws_rte53_record_info"),
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
		resource.TestCheckResourceAttrPair(resourceName, "port", dataSourceName, "result.0.port"),
		resource.TestCheckResourceAttrPair(resourceName, "priority", dataSourceName, "result.0.priority"),
		resource.TestCheckResourceAttrPair(resourceName, "reclaimable", dataSourceName, "result.0.reclaimable"),
		resource.TestCheckResourceAttrPair(resourceName, "shared_record_group", dataSourceName, "result.0.shared_record_group"),
		resource.TestCheckResourceAttrPair(resourceName, "target", dataSourceName, "result.0.target"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "weight", dataSourceName, "result.0.weight"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordSrvDataSourceConfigFilters(name, target string, port, priority, weight int) string {
	return fmt.Sprintf(`
resource "nios_dns_record_srv" "test" {
	name = %q
	target = %q
	port = %d
	priority = %d
	weight = %d
}

data "nios_dns_record_srv" "test" {
  filters = {
	name = nios_dns_record_srv.test.name
  }
}
`, name, target, port, priority, weight)
}

func testAccRecordSrvDataSourceConfigExtAttrFilters(name, target string, port, priority, weight int, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_srv" "test" {
	name = %q
	target = %q
	port = %d
	priority = %d
	weight = %d
	extattrs = {
    	Site = %q
  } 
}

data "nios_dns_record_srv" "test" {
  extattrfilters = {
	Site = nios_dns_record_srv.test.extattrs.Site
  }
}
`, name, target, port, priority, weight, extAttrsValue)
}
