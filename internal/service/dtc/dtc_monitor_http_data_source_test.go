
package dtc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccDtcMonitorHttpDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dtc_monitor_http.test"
	resourceName := "nios_dtc_monitor_http.test"
	var v dtc.DtcMonitorHttp

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcMonitorHttpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcMonitorHttpDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
						}, testAccCheckDtcMonitorHttpResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccDtcMonitorHttpDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dtc_monitor_http.test"
	resourceName := "nios_dtc_monitor_http.test"
	var v dtc.DtcMonitorHttp
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcMonitorHttpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcMonitorHttpDataSourceConfigExtAttrFilters( acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
						}, testAccCheckDtcMonitorHttpResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDtcMonitorHttpResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
    return []resource.TestCheckFunc{
        resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "ciphers", dataSourceName, "result.0.ciphers"),
        resource.TestCheckResourceAttrPair(resourceName, "client_cert", dataSourceName, "result.0.client_cert"),
        resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
        resource.TestCheckResourceAttrPair(resourceName, "content_check", dataSourceName, "result.0.content_check"),
        resource.TestCheckResourceAttrPair(resourceName, "content_check_input", dataSourceName, "result.0.content_check_input"),
        resource.TestCheckResourceAttrPair(resourceName, "content_check_op", dataSourceName, "result.0.content_check_op"),
        resource.TestCheckResourceAttrPair(resourceName, "content_check_regex", dataSourceName, "result.0.content_check_regex"),
        resource.TestCheckResourceAttrPair(resourceName, "content_extract_group", dataSourceName, "result.0.content_extract_group"),
        resource.TestCheckResourceAttrPair(resourceName, "content_extract_type", dataSourceName, "result.0.content_extract_type"),
        resource.TestCheckResourceAttrPair(resourceName, "content_extract_value", dataSourceName, "result.0.content_extract_value"),
        resource.TestCheckResourceAttrPair(resourceName, "enable_sni", dataSourceName, "result.0.enable_sni"),
        resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
        resource.TestCheckResourceAttrPair(resourceName, "interval", dataSourceName, "result.0.interval"),
        resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
        resource.TestCheckResourceAttrPair(resourceName, "port", dataSourceName, "result.0.port"),
        resource.TestCheckResourceAttrPair(resourceName, "request", dataSourceName, "result.0.request"),
        resource.TestCheckResourceAttrPair(resourceName, "result", dataSourceName, "result.0.result"),
        resource.TestCheckResourceAttrPair(resourceName, "result_code", dataSourceName, "result.0.result_code"),
        resource.TestCheckResourceAttrPair(resourceName, "retry_down", dataSourceName, "result.0.retry_down"),
        resource.TestCheckResourceAttrPair(resourceName, "retry_up", dataSourceName, "result.0.retry_up"),
        resource.TestCheckResourceAttrPair(resourceName, "secure", dataSourceName, "result.0.secure"),
        resource.TestCheckResourceAttrPair(resourceName, "timeout", dataSourceName, "result.0.timeout"),
        resource.TestCheckResourceAttrPair(resourceName, "validate_cert", dataSourceName, "result.0.validate_cert"),
    }
}

func testAccDtcMonitorHttpDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test" {
}

data "nios_dtc_monitor_http" "test" {
  filters = {
	 = nios_dtc_monitor_http.test.
  }
}
`)
}

func testAccDtcMonitorHttpDataSourceConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test" {
  extattrs = {
    Site = %q
  } 
}

data "nios_dtc_monitor_http" "test" {
  extattrfilters = {
	Site = nios_dtc_monitor_http.test.extattrs.Site
  }
}
`,extAttrsValue)
}

