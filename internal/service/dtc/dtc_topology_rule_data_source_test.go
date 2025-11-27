package dtc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccDtcTopologyRuleDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dtc_topology_rule.test"
	resourceName := "nios_dtc_topology_rule.test"
	var v dtc.DtcTopologyRule

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcTopologyRuleDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcTopologyRuleDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDtcTopologyRuleExists(context.Background(), resourceName, &v),
					}, testAccCheckDtcTopologyRuleResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDtcTopologyRuleResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "dest_type", dataSourceName, "result.0.dest_type"),
		resource.TestCheckResourceAttrPair(resourceName, "destination_link", dataSourceName, "result.0.destination_link"),
		resource.TestCheckResourceAttrPair(resourceName, "return_type", dataSourceName, "result.0.return_type"),
		resource.TestCheckResourceAttrPair(resourceName, "sources", dataSourceName, "result.0.sources"),
		resource.TestCheckResourceAttrPair(resourceName, "topology", dataSourceName, "result.0.topology"),
		resource.TestCheckResourceAttrPair(resourceName, "valid", dataSourceName, "result.0.valid"),
	}
}

func testAccDtcTopologyRuleDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_dtc_topology_rule" "test" {
}

data "nios_dtc_topology_rule" "test" {
  filters = {
	 = nios_dtc_topology_rule.test.
  }
}
`)
}
