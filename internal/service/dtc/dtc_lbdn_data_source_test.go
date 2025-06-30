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
				Config: testAccDtcLbdnDataSourceConfigFilters("lbdn-data-source-test", "ROUND_ROBIN"),
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
				Config: testAccDtcLbdnDataSourceConfigExtAttrFilters("lbdn-data-source-test2", "ROUND_ROBIN", "Denmark"),
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

func testAccCheckDtcLbdnResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
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

func testAccDtcLbdnDataSourceConfigFilters(name, lbMethod string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test" {
	name = %q
	lb_method = %q
}

data "nios_dtc_lbdn" "test" {
  filters = {
	 name = nios_dtc_lbdn.test.name
  }
}
`, name, lbMethod)
}

func testAccDtcLbdnDataSourceConfigExtAttrFilters(name, lbMethod, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test" {
  name = %q
  lb_method = %q
  extattrs = {
    Site = %q
  	}
}

data "nios_dtc_lbdn" "test" {
  extattrfilters = {
	"Site" = nios_dtc_lbdn.test.extattrs.Site
  }
}
`, name, lbMethod, extAttrsValue)
}
