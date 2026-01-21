package dtc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccDtcMonitorSnmpDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dtc_monitor_snmp.test"
	resourceName := "nios_dtc_monitor_snmp.test"
	var v dtc.DtcMonitorSnmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-snmp")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcMonitorSnmpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcMonitorSnmpDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					}, testAccCheckDtcMonitorSnmpResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccDtcMonitorSnmpDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dtc_monitor_snmp.test"
	resourceName := "nios_dtc_monitor_snmp.test"
	var v dtc.DtcMonitorSnmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-snmp")
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcMonitorSnmpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcMonitorSnmpDataSourceConfigExtAttrFilters(name, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					}, testAccCheckDtcMonitorSnmpResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDtcMonitorSnmpResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "community", dataSourceName, "result.0.community"),
		resource.TestCheckResourceAttrPair(resourceName, "context", dataSourceName, "result.0.context"),
		resource.TestCheckResourceAttrPair(resourceName, "engine_id", dataSourceName, "result.0.engine_id"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "interval", dataSourceName, "result.0.interval"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "oids", dataSourceName, "result.0.oids"),
		resource.TestCheckResourceAttrPair(resourceName, "port", dataSourceName, "result.0.port"),
		resource.TestCheckResourceAttrPair(resourceName, "retry_down", dataSourceName, "result.0.retry_down"),
		resource.TestCheckResourceAttrPair(resourceName, "retry_up", dataSourceName, "result.0.retry_up"),
		resource.TestCheckResourceAttrPair(resourceName, "timeout", dataSourceName, "result.0.timeout"),
		resource.TestCheckResourceAttrPair(resourceName, "user", dataSourceName, "result.0.user"),
		resource.TestCheckResourceAttrPair(resourceName, "version", dataSourceName, "result.0.version"),
	}
}

func testAccDtcMonitorSnmpDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test" {
    name = %q
}

data "nios_dtc_monitor_snmp" "test" {
  filters = {
    name = nios_dtc_monitor_snmp.test.name
  }
}
`, name)
}

func testAccDtcMonitorSnmpDataSourceConfigExtAttrFilters(name, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test" {
  name = %q
  extattrs = {
    Site = %q
  } 
}

data "nios_dtc_monitor_snmp" "test" {
  extattrfilters = {
	Site = nios_dtc_monitor_snmp.test.extattrs.Site
  }
}
`, name, extAttrsValue)
}
