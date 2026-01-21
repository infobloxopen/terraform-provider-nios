package dhcp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccIpv6dhcpoptionspaceDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_ipv6optionspace.test"
	resourceName := "nios_dhcp_ipv6optionspace.test"
	var v dhcp.Ipv6dhcpoptionspace
	name := acctest.RandomNameWithPrefix("option-space")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6dhcpoptionspaceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6dhcpoptionspaceDataSourceConfigFilters("5674", name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIpv6dhcpoptionspaceExists(context.Background(), resourceName, &v),
					}, testAccCheckIpv6dhcpoptionspaceResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckIpv6dhcpoptionspaceResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "enterprise_number", dataSourceName, "result.0.enterprise_number"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "option_definitions", dataSourceName, "result.0.option_definitions"),
	}
}

func testAccIpv6dhcpoptionspaceDataSourceConfigFilters(enterpriseNumber, name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6optionspace" "test" {
	enterprise_number = %q
	name = %q
}

data "nios_dhcp_ipv6optionspace" "test" {
	filters = {
		enterprise_number = nios_dhcp_ipv6optionspace.test.enterprise_number
		name = nios_dhcp_ipv6optionspace.test.name
	}
}
`, enterpriseNumber, name)
}
