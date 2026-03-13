package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccRecordNsDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_ns.test"
	resourceName := "nios_dns_record_ns.test"
	var v dns.RecordNs
	name := "example.com"
	nameserver := acctest.RandomNameWithPrefix("nameserver") + ".example.com"
	addresses := []map[string]any{
		{
			"address":         "20.0.0.0",
			"auto_create_ptr": false,
		},
	}
	addressesHCL := FormatZoneNameServersToHCL(addresses)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordNsDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordNsDataSourceConfigFilters(name, nameserver, addressesHCL, "default"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordNsExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordNsResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordNsResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "addresses", dataSourceName, "result.0.addresses"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_info", dataSourceName, "result.0.cloud_info"),
		resource.TestCheckResourceAttrPair(resourceName, "creator", dataSourceName, "result.0.creator"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_name", dataSourceName, "result.0.dns_name"),
		resource.TestCheckResourceAttrPair(resourceName, "last_queried", dataSourceName, "result.0.last_queried"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_delegation_name", dataSourceName, "result.0.ms_delegation_name"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "nameserver", dataSourceName, "result.0.nameserver"),
		resource.TestCheckResourceAttrPair(resourceName, "policy", dataSourceName, "result.0.policy"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordNsDataSourceConfigFilters(name, nameserver, addresses, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ns" "test" {
    name 	   = %q
	nameserver = %q
	addresses  = %s
	view       = %q
}

data "nios_dns_record_ns" "test" {
  filters = {
	name = nios_dns_record_ns.test.name
	nameserver = nios_dns_record_ns.test.nameserver
  }
}
`, name, nameserver, addresses, view)
}
