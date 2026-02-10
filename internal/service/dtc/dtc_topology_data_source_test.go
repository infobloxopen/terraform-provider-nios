package dtc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccDtcTopologyDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dtc_topology.test"
	resourceName := "nios_dtc_topology.test"
	var v dtc.DtcTopology
	name := acctest.RandomNameWithPrefix("dtc-topology")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcTopologyDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcTopologyDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					}, testAccCheckDtcTopologyResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccDtcTopologyDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dtc_topology.test"
	resourceName := "nios_dtc_topology.test"
	var v dtc.DtcTopology
	name := acctest.RandomNameWithPrefix("dtc-topology")
	extAttrValue := acctest.RandomName()
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcTopologyDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcTopologyDataSourceConfigExtAttrFilters(name, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					}, testAccCheckDtcTopologyResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDtcTopologyResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "rules", dataSourceName, "result.0.rules"),
	}
}

func testAccDtcTopologyDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_dtc_topology" "test" {
  name = %q
}

data "nios_dtc_topology" "test" {
  filters = {
	name = nios_dtc_topology.test.name
  }
}
`, name)
}

func testAccDtcTopologyDataSourceConfigExtAttrFilters(name, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dtc_topology" "test" {
  name = %q
  extattrs = {
    Site = %q
  } 
}

data "nios_dtc_topology" "test" {
  extattrfilters = {
	Site = nios_dtc_topology.test.extattrs.Site
  }
}
`, name, extAttrsValue)
}
