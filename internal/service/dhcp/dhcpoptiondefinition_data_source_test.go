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

func TestAccDhcpoptiondefinitionDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_optiondefinition.test"
	resourceName := "nios_dhcp_optiondefinition.test"
	var v dhcp.Dhcpoptiondefinition
	name := acctest.RandomNameWithPrefix("dhcp-option-definition")
	space := acctest.RandomNameWithPrefix("dhcp-option-space")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDhcpoptiondefinitionDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDhcpoptiondefinitionDataSourceConfigFilters("10", name, "string", space),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDhcpoptiondefinitionExists(context.Background(), resourceName, &v),
					}, testAccCheckDhcpoptiondefinitionResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDhcpoptiondefinitionResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "code", dataSourceName, "result.0.code"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "space", dataSourceName, "result.0.space"),
		resource.TestCheckResourceAttrPair(resourceName, "type", dataSourceName, "result.0.type"),
	}
}

func testAccDhcpoptiondefinitionDataSourceConfigFilters(code, name, optionType, optionSpace string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_optiondefinition" "test" {
  code = %q
  name = %q
  type = %q
  space = nios_dhcp_optionspace.test.name
}

data "nios_dhcp_optiondefinition" "test" {
  filters = {
	name = nios_dhcp_optiondefinition.test.name
	code = nios_dhcp_optiondefinition.test.code
	space = nios_dhcp_optionspace.test.name
	type = nios_dhcp_optiondefinition.test.type
  }
}
`, code, name, optionType)
	return strings.Join([]string{testAccBaseWithDHCPOptionSpace(optionSpace), config}, "")
}
