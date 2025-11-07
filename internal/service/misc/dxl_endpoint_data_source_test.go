
package misc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccDxlEndpointDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_misc_dxl_endpoint.test"
	resourceName := "nios_misc_dxl_endpoint.test"
	var v misc.DxlEndpoint

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDxlEndpointDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDxlEndpointDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
						}, testAccCheckDxlEndpointResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccDxlEndpointDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_misc_dxl_endpoint.test"
	resourceName := "nios_misc_dxl_endpoint.test"
	var v misc.DxlEndpoint
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDxlEndpointDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDxlEndpointDataSourceConfigExtAttrFilters( acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
						}, testAccCheckDxlEndpointResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDxlEndpointResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
    return []resource.TestCheckFunc{
        resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "brokers", dataSourceName, "result.0.brokers"),
        resource.TestCheckResourceAttrPair(resourceName, "brokers_import_token", dataSourceName, "result.0.brokers_import_token"),
        resource.TestCheckResourceAttrPair(resourceName, "client_certificate_subject", dataSourceName, "result.0.client_certificate_subject"),
        resource.TestCheckResourceAttrPair(resourceName, "client_certificate_token", dataSourceName, "result.0.client_certificate_token"),
        resource.TestCheckResourceAttrPair(resourceName, "client_certificate_valid_from", dataSourceName, "result.0.client_certificate_valid_from"),
        resource.TestCheckResourceAttrPair(resourceName, "client_certificate_valid_to", dataSourceName, "result.0.client_certificate_valid_to"),
        resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
        resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
        resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
        resource.TestCheckResourceAttrPair(resourceName, "log_level", dataSourceName, "result.0.log_level"),
        resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
        resource.TestCheckResourceAttrPair(resourceName, "outbound_member_type", dataSourceName, "result.0.outbound_member_type"),
        resource.TestCheckResourceAttrPair(resourceName, "outbound_members", dataSourceName, "result.0.outbound_members"),
        resource.TestCheckResourceAttrPair(resourceName, "template_instance", dataSourceName, "result.0.template_instance"),
        resource.TestCheckResourceAttrPair(resourceName, "timeout", dataSourceName, "result.0.timeout"),
        resource.TestCheckResourceAttrPair(resourceName, "topics", dataSourceName, "result.0.topics"),
        resource.TestCheckResourceAttrPair(resourceName, "vendor_identifier", dataSourceName, "result.0.vendor_identifier"),
        resource.TestCheckResourceAttrPair(resourceName, "wapi_user_name", dataSourceName, "result.0.wapi_user_name"),
        resource.TestCheckResourceAttrPair(resourceName, "wapi_user_password", dataSourceName, "result.0.wapi_user_password"),
    }
}

func testAccDxlEndpointDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test" {
}

data "nios_misc_dxl_endpoint" "test" {
  filters = {
	 = nios_misc_dxl_endpoint.test.
  }
}
`)
}

func testAccDxlEndpointDataSourceConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test" {
  extattrs = {
    Site = %q
  } 
}

data "nios_misc_dxl_endpoint" "test" {
  extattrfilters = {
	Site = nios_misc_dxl_endpoint.test.extattrs.Site
  }
}
`,extAttrsValue)
}

