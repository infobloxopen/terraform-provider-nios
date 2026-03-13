package dtc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccDtcMonitorIcmpDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dtc_monitor_icmp.test"
	resourceName := "nios_dtc_monitor_icmp.test"
	var v dtc.DtcMonitorIcmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-icmp")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcMonitorIcmpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcMonitorIcmpDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDtcMonitorIcmpExists(context.Background(), resourceName, &v),
					}, testAccCheckDtcMonitorIcmpResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccDtcMonitorIcmpDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dtc_monitor_icmp.test"
	resourceName := "nios_dtc_monitor_icmp.test"
	var v dtc.DtcMonitorIcmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-icmp")
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcMonitorIcmpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcMonitorIcmpDataSourceConfigExtAttrFilters(name, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDtcMonitorIcmpExists(context.Background(), resourceName, &v),
					}, testAccCheckDtcMonitorIcmpResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDtcMonitorIcmpResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "interval", dataSourceName, "result.0.interval"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "retry_down", dataSourceName, "result.0.retry_down"),
		resource.TestCheckResourceAttrPair(resourceName, "retry_up", dataSourceName, "result.0.retry_up"),
		resource.TestCheckResourceAttrPair(resourceName, "timeout", dataSourceName, "result.0.timeout"),
	}
}

func testAccDtcMonitorIcmpDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_icmp" "test" {
  name = %q
}

data "nios_dtc_monitor_icmp" "test" {
  filters = {
    name = nios_dtc_monitor_icmp.test.name
  }
}
`, name)
}

func testAccDtcMonitorIcmpDataSourceConfigExtAttrFilters(name, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_icmp" "test" {
  name = %q
  extattrs = {
    Site = %q
  }
}

data "nios_dtc_monitor_icmp" "test" {
  extattrfilters = {
	Site = nios_dtc_monitor_icmp.test.extattrs.Site
  }
}
`, name, extAttrsValue)
}
