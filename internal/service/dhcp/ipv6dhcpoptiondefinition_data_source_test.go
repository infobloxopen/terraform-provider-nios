package dhcp_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccIpv6dhcpoptiondefinitionDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_ipv6dhcpoptiondefinition.test"
	resourceName := "nios_dhcp_ipv6dhcpoptiondefinition.test"
	var v dhcp.Ipv6dhcpoptiondefinition
	name := acctest.RandomNameWithPrefix("option-definition")
	optionSpace := acctest.RandomNameWithPrefix("option-space")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6dhcpoptiondefinitionDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6dhcpoptiondefinitionDataSourceConfigFilters(optionSpace, "10", name, "string"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIpv6dhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					}, testAccCheckIpv6dhcpoptiondefinitionResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckIpv6dhcpoptiondefinitionResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "code", dataSourceName, "result.0.code"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "space", dataSourceName, "result.0.space"),
		resource.TestCheckResourceAttrPair(resourceName, "type", dataSourceName, "result.0.type"),
	}
}

func testAccIpv6dhcpoptiondefinitionDataSourceConfigFilters(optionSpace, code, name, optionType string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6dhcpoptiondefinition" "test" {
  code = %q
  name = %q
  type = %q
  space = nios_dhcp_ipv6dhcpoptionspace.test.name
}

data "nios_dhcp_ipv6dhcpoptiondefinition" "test" {
  filters = {
	code = nios_dhcp_ipv6dhcpoptiondefinition.test.code
	name = nios_dhcp_ipv6dhcpoptiondefinition.test.name
	type = nios_dhcp_ipv6dhcpoptiondefinition.test.type
  }
}
`, code, name, optionType)
	return strings.Join([]string{testAccBaseWithIpv6DHCPOptionSpace(optionSpace, "10"), config}, "")
}
