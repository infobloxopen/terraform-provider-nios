
package dtc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/acctest"
)

func TestAccDtcLbdnDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dtc_lbdn.test"
	resourceName := "nios_dtc_lbdn.test"
	var v dtc.DtcLbdn

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcLbdnDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcLbdnDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
						}, testAccCheckDtcLbdnResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccDtcLbdnDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.nios_dtc_lbdn.test"
	resourceName := "nios_dtc_lbdn.test"
	var v dtc.DtcLbdn
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcLbdnDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcLbdnDataSourceConfigExtAttrFilters(, "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
						}, testAccCheckDtcLbdnResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDtcLbdnResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
    return []resource.TestCheckFunc{
        resource.TestCheckResourceAttrPair(resourceName, "_ref", dataSourceName, "result.0._ref"),
        resource.TestCheckResourceAttrPair(resourceName, "auth_zones", dataSourceName, "result.0.auth_zones"),
        resource.TestCheckResourceAttrPair(resourceName, "auto_consolidated_monitors", dataSourceName, "result.0.auto_consolidated_monitors"),
        resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
        resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
        resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
        resource.TestCheckResourceAttrPair(resourceName, "health", dataSourceName, "result.0.health"),
        resource.TestCheckResourceAttrPair(resourceName, "lb_method", dataSourceName, "result.0.lb_method"),
        resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
        resource.TestCheckResourceAttrPair(resourceName, "patterns", dataSourceName, "result.0.patterns"),
        resource.TestCheckResourceAttrPair(resourceName, "persistence", dataSourceName, "result.0.persistence"),
        resource.TestCheckResourceAttrPair(resourceName, "pools", dataSourceName, "result.0.pools"),
        resource.TestCheckResourceAttrPair(resourceName, "priority", dataSourceName, "result.0.priority"),
        resource.TestCheckResourceAttrPair(resourceName, "topology", dataSourceName, "result.0.topology"),
        resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
        resource.TestCheckResourceAttrPair(resourceName, "types", dataSourceName, "result.0.types"),
        resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
    }
}

func testAccDtcLbdnDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test" {
}

data "nios_dtc_lbdn" "test" {
  filters = {
	 = nios_dtc_lbdn.test.
  }
}
`)
}

func testAccDtcLbdnDataSourceConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test" {
  extattrs = {
    Site = {
        value = %q
    }
  	}
}

data "nios_dtc_lbdn" "test" {
  extattrfilters = {
	"Site" = nios_dtc_lbdn.test.tags.tag1
  }
}
`,extAttrsValue)
}

