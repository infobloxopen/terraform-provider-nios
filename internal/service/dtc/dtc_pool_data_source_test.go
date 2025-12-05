package dtc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccDtcPoolDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dtc_pool.test"
	resourceName := "nios_dtc_pool.test"
	var v dtc.DtcPool

	name := acctest.RandomNameWithPrefix("dtc-pool")
	lbPreferredMethod := "ROUND_ROBIN"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcPoolDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcPoolDataSourceConfigFilters(name, lbPreferredMethod),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					}, testAccCheckDtcPoolResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccDtcPoolDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dtc_pool.test"
	resourceName := "nios_dtc_pool.test"
	var v dtc.DtcPool
	name := acctest.RandomNameWithPrefix("dtc-pool")
	lbPreferredMethod := "ROUND_ROBIN"
	extAttrValue := acctest.RandomName()
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcPoolDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcPoolDataSourceConfigExtAttrFilters(name, lbPreferredMethod, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					}, testAccCheckDtcPoolResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDtcPoolResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "auto_consolidated_monitors", dataSourceName, "result.0.auto_consolidated_monitors"),
		resource.TestCheckResourceAttrPair(resourceName, "availability", dataSourceName, "result.0.availability"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "consolidated_monitors", dataSourceName, "result.0.consolidated_monitors"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "health", dataSourceName, "result.0.health"),
		resource.TestCheckResourceAttrPair(resourceName, "lb_alternate_method", dataSourceName, "result.0.lb_alternate_method"),
		resource.TestCheckResourceAttrPair(resourceName, "lb_alternate_topology", dataSourceName, "result.0.lb_alternate_topology"),
		resource.TestCheckResourceAttrPair(resourceName, "lb_dynamic_ratio_alternate", dataSourceName, "result.0.lb_dynamic_ratio_alternate"),
		resource.TestCheckResourceAttrPair(resourceName, "lb_dynamic_ratio_preferred", dataSourceName, "result.0.lb_dynamic_ratio_preferred"),
		resource.TestCheckResourceAttrPair(resourceName, "lb_preferred_method", dataSourceName, "result.0.lb_preferred_method"),
		resource.TestCheckResourceAttrPair(resourceName, "lb_preferred_topology", dataSourceName, "result.0.lb_preferred_topology"),
		resource.TestCheckResourceAttrPair(resourceName, "monitors", dataSourceName, "result.0.monitors"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "quorum", dataSourceName, "result.0.quorum"),
		resource.TestCheckResourceAttrPair(resourceName, "servers", dataSourceName, "result.0.servers"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
	}
}

func testAccDtcPoolDataSourceConfigFilters(name, lbPreferredMethod string) string {
	return fmt.Sprintf(`
resource "nios_dtc_pool" "test" {
  name = "%s"
  lb_preferred_method = "%s"
}

data "nios_dtc_pool" "test" {
  filters = {
    name = nios_dtc_pool.test.name 
  }
}
`, name, lbPreferredMethod)
}

func testAccDtcPoolDataSourceConfigExtAttrFilters(name, lbPreferredMethod, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dtc_pool" "test" {
  name = "%s"
  lb_preferred_method = "%s"
  extattrs = {
    Site = %q
  	}
}

data "nios_dtc_pool" "test" {
  extattrfilters = {
	Site = nios_dtc_pool.test.extattrs.Site
  }
}
`, name, lbPreferredMethod, extAttrsValue)
}
