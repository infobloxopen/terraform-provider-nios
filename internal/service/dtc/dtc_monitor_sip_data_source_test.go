package dtc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccDtcMonitorSipDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dtc_monitor_sip.test"
	resourceName := "nios_dtc_monitor_sip.test"
	var v dtc.DtcMonitorSip
	name := acctest.RandomNameWithPrefix("dtc-monitor-sip")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcMonitorSipDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcMonitorSipDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					}, testAccCheckDtcMonitorSipResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccDtcMonitorSipDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dtc_monitor_sip.test"
	resourceName := "nios_dtc_monitor_sip.test"
	var v dtc.DtcMonitorSip
	name := acctest.RandomNameWithPrefix("dtc-monitor-sip")
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcMonitorSipDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcMonitorSipDataSourceConfigExtAttrFilters(name, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					}, testAccCheckDtcMonitorSipResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDtcMonitorSipResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "ciphers", dataSourceName, "result.0.ciphers"),
		resource.TestCheckResourceAttrPair(resourceName, "client_cert", dataSourceName, "result.0.client_cert"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "interval", dataSourceName, "result.0.interval"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "port", dataSourceName, "result.0.port"),
		resource.TestCheckResourceAttrPair(resourceName, "request", dataSourceName, "result.0.request"),
		resource.TestCheckResourceAttrPair(resourceName, "result", dataSourceName, "result.0.result"),
		resource.TestCheckResourceAttrPair(resourceName, "result_code", dataSourceName, "result.0.result_code"),
		resource.TestCheckResourceAttrPair(resourceName, "retry_down", dataSourceName, "result.0.retry_down"),
		resource.TestCheckResourceAttrPair(resourceName, "retry_up", dataSourceName, "result.0.retry_up"),
		resource.TestCheckResourceAttrPair(resourceName, "timeout", dataSourceName, "result.0.timeout"),
		resource.TestCheckResourceAttrPair(resourceName, "transport", dataSourceName, "result.0.transport"),
		resource.TestCheckResourceAttrPair(resourceName, "validate_cert", dataSourceName, "result.0.validate_cert"),
	}
}

func testAccDtcMonitorSipDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_sip" "test" {
  name = %q
}

data "nios_dtc_monitor_sip" "test" {
  filters = {
    name = nios_dtc_monitor_sip.test.name
  }
}
`, name)
}

func testAccDtcMonitorSipDataSourceConfigExtAttrFilters(name, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_sip" "test" {
  name = %q
  extattrs = {
    Site = %q
  } 
}

data "nios_dtc_monitor_sip" "test" {
  extattrfilters = {
	Site = nios_dtc_monitor_sip.test.extattrs.Site
  }
}
`, name, extAttrsValue)
}
