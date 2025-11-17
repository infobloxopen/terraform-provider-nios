package dtc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccDtcMonitorTcpDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dtc_monitor_tcp.test"
	resourceName := "nios_dtc_monitor_tcp.test"
	var v dtc.DtcMonitorTcp
	name := acctest.RandomNameWithPrefix("dtc-monitor-tcp")
	port := 49152

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcMonitorTcpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcMonitorTcpDataSourceConfigFilters(name, port),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					}, testAccCheckDtcMonitorTcpResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccDtcMonitorTcpDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dtc_monitor_tcp.test"
	resourceName := "nios_dtc_monitor_tcp.test"
	var v dtc.DtcMonitorTcp
	name := acctest.RandomNameWithPrefix("dtc-monitor-tcp")
	port := 49152
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcMonitorTcpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcMonitorTcpDataSourceConfigExtAttrFilters(name, port, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					}, testAccCheckDtcMonitorTcpResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDtcMonitorTcpResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
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

func testAccDtcMonitorTcpDataSourceConfigFilters(name string, port int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test" {
  name = %q
  port = %d
}

data "nios_dtc_monitor_tcp" "test" {
  filters = {
    name = nios_dtc_monitor_tcp.test.name
  }
}
`, name, port)
}

func testAccDtcMonitorTcpDataSourceConfigExtAttrFilters(name string, port int, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test" {
  name = %q
  port = %d
  extattrs = {
    Site = %q
  } 
}

data "nios_dtc_monitor_tcp" "test" {
  extattrfilters = {
	Site = nios_dtc_monitor_tcp.test.extattrs.Site
  }
}
`, name, port, extAttrsValue)
}
