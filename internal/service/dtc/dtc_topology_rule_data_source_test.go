package dtc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

var readableAttributesForDtcTopologyRule = "dest_type,destination_link,return_type,sources,topology,valid"

func TestAccDtcTopologyRuleDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dtc_topology_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDtcTopologyRuleDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "result.#"),
					resource.TestCheckResourceAttrSet(dataSourceName, "result.0.dest_type"),
					resource.TestCheckResourceAttrSet(dataSourceName, "result.0.destination_link"),
					resource.TestCheckResourceAttrSet(dataSourceName, "result.0.return_type"),
					resource.TestCheckResourceAttrSet(dataSourceName, "result.0.sources.#"),
					resource.TestCheckResourceAttrSet(dataSourceName, "result.0.topology"),
					resource.TestCheckResourceAttrSet(dataSourceName, "result.0.valid"),
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccDtcTopologyRuleDataSourceConfigFilters() string {
	return fmt.Sprintf(`
data "nios_dtc_topology_rule" "test" {
  filters = {
	 = nios_dtc_topology_rule.test.
  }
}
`)
}
