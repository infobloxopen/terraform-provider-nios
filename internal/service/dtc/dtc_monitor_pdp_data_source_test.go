
package dtc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccDtcMonitorPdpDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dtc_monitor_pdp.test"
	resourceName := "nios_dtc_monitor_pdp.test"
	var v dtc.DtcMonitorPdp
	name := acctest.RandomNameWithPrefix("dtc-monitor-pdp")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcMonitorPdpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcMonitorPdpDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckDtcMonitorPdpExists(context.Background(), resourceName, &v),
						}, testAccCheckDtcMonitorPdpResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccDtcMonitorPdpDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dtc_monitor_pdp.test"
	resourceName := "nios_dtc_monitor_pdp.test"
	var v dtc.DtcMonitorPdp
	name := acctest.RandomNameWithPrefix("dtc-monitor-pdp")
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcMonitorPdpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcMonitorPdpDataSourceConfigExtAttrFilters(name , extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckDtcMonitorPdpExists(context.Background(), resourceName, &v),
						}, testAccCheckDtcMonitorPdpResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDtcMonitorPdpResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
    return []resource.TestCheckFunc{
        resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
        resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
        resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
        resource.TestCheckResourceAttrPair(resourceName, "interval", dataSourceName, "result.0.interval"),
        resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
        resource.TestCheckResourceAttrPair(resourceName, "port", dataSourceName, "result.0.port"),
        resource.TestCheckResourceAttrPair(resourceName, "retry_down", dataSourceName, "result.0.retry_down"),
        resource.TestCheckResourceAttrPair(resourceName, "retry_up", dataSourceName, "result.0.retry_up"),
        resource.TestCheckResourceAttrPair(resourceName, "timeout", dataSourceName, "result.0.timeout"),
    }
}

func testAccDtcMonitorPdpDataSourceConfigFilters(name string ) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_pdp" "test" {
  name = %q	
}

data "nios_dtc_monitor_pdp" "test" {
  filters = {
    name = nios_dtc_monitor_pdp.test.name
  }
}
`, name)
}

func testAccDtcMonitorPdpDataSourceConfigExtAttrFilters(name, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_pdp" "test" {
  name = %q	
  extattrs = {
    Site = %q
  } 
}

data "nios_dtc_monitor_pdp" "test" {
  extattrfilters = {
	Site = nios_dtc_monitor_pdp.test.extattrs.Site
  }
}
`, name, extAttrsValue)
}

