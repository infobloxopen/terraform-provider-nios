
package misc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccSyslogEndpointDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_misc_syslog_endpoint.test"
	resourceName := "nios_misc_syslog_endpoint.test"
	var v misc.SyslogEndpoint

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSyslogEndpointDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSyslogEndpointDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
						}, testAccCheckSyslogEndpointResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccSyslogEndpointDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_misc_syslog_endpoint.test"
	resourceName := "nios_misc_syslog_endpoint.test"
	var v misc.SyslogEndpoint
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSyslogEndpointDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSyslogEndpointDataSourceConfigExtAttrFilters( acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
						}, testAccCheckSyslogEndpointResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckSyslogEndpointResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
    return []resource.TestCheckFunc{
        resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
        resource.TestCheckResourceAttrPair(resourceName, "log_level", dataSourceName, "result.0.log_level"),
        resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
        resource.TestCheckResourceAttrPair(resourceName, "outbound_member_type", dataSourceName, "result.0.outbound_member_type"),
        resource.TestCheckResourceAttrPair(resourceName, "outbound_members", dataSourceName, "result.0.outbound_members"),
        resource.TestCheckResourceAttrPair(resourceName, "syslog_servers", dataSourceName, "result.0.syslog_servers"),
        resource.TestCheckResourceAttrPair(resourceName, "template_instance", dataSourceName, "result.0.template_instance"),
        resource.TestCheckResourceAttrPair(resourceName, "timeout", dataSourceName, "result.0.timeout"),
        resource.TestCheckResourceAttrPair(resourceName, "vendor_identifier", dataSourceName, "result.0.vendor_identifier"),
        resource.TestCheckResourceAttrPair(resourceName, "wapi_user_name", dataSourceName, "result.0.wapi_user_name"),
        resource.TestCheckResourceAttrPair(resourceName, "wapi_user_password", dataSourceName, "result.0.wapi_user_password"),
    }
}

func testAccSyslogEndpointDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test" {
}

data "nios_misc_syslog_endpoint" "test" {
  filters = {
	 = nios_misc_syslog_endpoint.test.
  }
}
`)
}

func testAccSyslogEndpointDataSourceConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test" {
  extattrs = {
    Site = %q
  } 
}

data "nios_misc_syslog_endpoint" "test" {
  extattrfilters = {
	Site = nios_misc_syslog_endpoint.test.extattrs.Site
  }
}
`,extAttrsValue)
}

