package threatinsight_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/threatinsight"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccThreatinsightAllowlistDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_threatinsight_allowlist.test"
	resourceName := "nios_threatinsight_allowlist.test"
	var v threatinsight.ThreatinsightAllowlist

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckThreatinsightAllowlistDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccThreatinsightAllowlistDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckThreatinsightAllowlistExists(context.Background(), resourceName, &v),
					}, testAccCheckThreatinsightAllowlistResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckThreatinsightAllowlistResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "fqdn", dataSourceName, "result.0.fqdn"),
		resource.TestCheckResourceAttrPair(resourceName, "type", dataSourceName, "result.0.type"),
	}
}

func testAccThreatinsightAllowlistDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_threatinsight_allowlist" "test" {
}

data "nios_threatinsight_allowlist" "test" {
  filters = {
	 = nios_threatinsight_allowlist.test.
  }
}
`)
}
