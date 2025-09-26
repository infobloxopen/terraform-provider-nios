package dhcp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccIpv6rangetemplateDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_ipv6rangetemplate.test"
	resourceName := "nios_dhcp_ipv6rangetemplate.test"
	var v dhcp.Ipv6rangetemplate

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6rangetemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6rangetemplateDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					}, testAccCheckIpv6rangetemplateResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckIpv6rangetemplateResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_api_compatible", dataSourceName, "result.0.cloud_api_compatible"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "delegated_member", dataSourceName, "result.0.delegated_member"),
		resource.TestCheckResourceAttrPair(resourceName, "exclude", dataSourceName, "result.0.exclude"),
		resource.TestCheckResourceAttrPair(resourceName, "logic_filter_rules", dataSourceName, "result.0.logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "member", dataSourceName, "result.0.member"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "number_of_addresses", dataSourceName, "result.0.number_of_addresses"),
		resource.TestCheckResourceAttrPair(resourceName, "offset", dataSourceName, "result.0.offset"),
		resource.TestCheckResourceAttrPair(resourceName, "option_filter_rules", dataSourceName, "result.0.option_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "recycle_leases", dataSourceName, "result.0.recycle_leases"),
		resource.TestCheckResourceAttrPair(resourceName, "server_association_type", dataSourceName, "result.0.server_association_type"),
		resource.TestCheckResourceAttrPair(resourceName, "use_logic_filter_rules", dataSourceName, "result.0.use_logic_filter_rules"),
		resource.TestCheckResourceAttrPair(resourceName, "use_recycle_leases", dataSourceName, "result.0.use_recycle_leases"),
	}
}

func testAccIpv6rangetemplateDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6rangetemplate" "test" {
}

data "nios_dhcp_ipv6rangetemplate" "test" {
  filters = {
	 = nios_dhcp_ipv6rangetemplate.test.
  }
}
`)
}
