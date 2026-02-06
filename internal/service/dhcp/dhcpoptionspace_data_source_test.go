package dhcp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccDhcpoptionspaceDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_optionspace.test"
	resourceName := "nios_dhcp_optionspace.test"
	var v dhcp.Dhcpoptionspace
	name := acctest.RandomNameWithPrefix("dhcp-option-space")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDhcpoptionspaceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDhcpoptionspaceDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDhcpoptionspaceExists(context.Background(), resourceName, &v),
					}, testAccCheckDhcpoptionspaceResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDhcpoptionspaceResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "option_definitions", dataSourceName, "result.0.option_definitions"),
		resource.TestCheckResourceAttrPair(resourceName, "space_type", dataSourceName, "result.0.space_type"),
	}
}

func testAccDhcpoptionspaceDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_optionspace" "test" {
  name = %q
}

data "nios_dhcp_optionspace" "test" {
  filters = {
	name = nios_dhcp_optionspace.test.name
  }
}
`, name)
}
