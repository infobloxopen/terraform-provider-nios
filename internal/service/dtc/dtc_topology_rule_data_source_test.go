package dtc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccDtcTopologyRuleDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dtc_topology_rule.test"
	resourceName := "nios_dtc_topology.test"
	serverName := acctest.RandomNameWithPrefix("dtc-server")
	topologyName := acctest.RandomNameWithPrefix("dtc-topology")
	randomIp := acctest.RandomIP()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDtcTopologyRuleDataSourceConfigFilters(serverName, topologyName, randomIp, "Africa"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{}, testAccCheckDtcTopologyRuleResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDtcTopologyRuleResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.topology"),
		resource.TestCheckResourceAttrPair(resourceName, "rules.0.destination_link", dataSourceName, "result.0.destination_link"),
		resource.TestCheckResourceAttrPair(resourceName, "rules.0.sources.0.source_op", dataSourceName, "result.0.sources.0.source_op"),
		resource.TestCheckResourceAttrPair(resourceName, "rules.0.sources.0.source_type", dataSourceName, "result.0.sources.0.source_type"),
		resource.TestCheckResourceAttrPair(resourceName, "rules.0.sources.0.source_value", dataSourceName, "result.0.sources.0.source_value"),
	}
}

func testAccDtcTopologyRuleDataSourceConfigFilters(serverName, topologyName, serverIP, sourceValue string) string {
	return fmt.Sprintf(`
resource "nios_dtc_server" "create_dtc_server" {
  name = %q
  host = %q
}
resource "nios_dtc_topology" "test" {
  name    = %q
  rules = [
    {
      dest_type        = "SERVER"
      destination_link = nios_dtc_server.create_dtc_server.ref
	  sources  = [
        {
            source_op   = "IS",
            source_type = "CONTINENT",
            source_value = %q
        }
      ]
    }
  ]
}

data "nios_dtc_topology_rule" "test" {
  filters = {
	topology = nios_dtc_topology.test.ref
  }
}
`, serverName, serverIP, topologyName, sourceValue)
}
