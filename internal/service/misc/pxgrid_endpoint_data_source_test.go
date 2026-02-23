package misc_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

func TestAccPxgridEndpointDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_misc_pxgrid_endpoint.test"
	resourceName := "nios_misc_pxgrid_endpoint.test"
	var v misc.PxgridEndpoint
	view := acctest.RandomNameWithPrefix("network-view")
	name := acctest.RandomNameWithPrefix("pxgrid-endpoint")
	address := acctest.RandomIP()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPxgridEndpointDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccPxgridEndpointDataSourceConfigFilters(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					}, testAccCheckPxgridEndpointResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccPxgridEndpointDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_misc_pxgrid_endpoint.test"
	resourceName := "nios_misc_pxgrid_endpoint.test"
	var v misc.PxgridEndpoint
	view := acctest.RandomNameWithPrefix("network-view")
	name := acctest.RandomNameWithPrefix("pxgrid-endpoint")
	address := acctest.RandomIP()
	extAttrsValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPxgridEndpointDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccPxgridEndpointDataSourceConfigExtAttrFilters(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings, extAttrsValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					}, testAccCheckPxgridEndpointResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckPxgridEndpointResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "result.0.address"),
		resource.TestCheckResourceAttrPair(resourceName, "client_certificate_subject", dataSourceName, "result.0.client_certificate_subject"),
		resource.TestCheckResourceAttrPair(resourceName, "client_certificate_token", dataSourceName, "result.0.client_certificate_token"),
		resource.TestCheckResourceAttrPair(resourceName, "client_certificate_valid_from", dataSourceName, "result.0.client_certificate_valid_from"),
		resource.TestCheckResourceAttrPair(resourceName, "client_certificate_valid_to", dataSourceName, "result.0.client_certificate_valid_to"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "log_level", dataSourceName, "result.0.log_level"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "network_view", dataSourceName, "result.0.network_view"),
		resource.TestCheckResourceAttrPair(resourceName, "outbound_member_type", dataSourceName, "result.0.outbound_member_type"),
		resource.TestCheckResourceAttrPair(resourceName, "outbound_members", dataSourceName, "result.0.outbound_members"),
		resource.TestCheckResourceAttrPair(resourceName, "publish_settings", dataSourceName, "result.0.publish_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "subscribe_settings", dataSourceName, "result.0.subscribe_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "template_instance", dataSourceName, "result.0.template_instance"),
		resource.TestCheckResourceAttrPair(resourceName, "timeout", dataSourceName, "result.0.timeout"),
		resource.TestCheckResourceAttrPair(resourceName, "vendor_identifier", dataSourceName, "result.0.vendor_identifier"),
		resource.TestCheckResourceAttrPair(resourceName, "wapi_user_name", dataSourceName, "result.0.wapi_user_name"),
		resource.TestCheckResourceAttrPair(resourceName, "wapi_user_password", dataSourceName, "result.0.wapi_user_password"),
	}
}

func testAccPxgridEndpointDataSourceConfigFilters(view, address, clientCertificateFile, name, outboundMemberType string, subscribeSettings map[string]any, publishSettings map[string]any) string {
	subscribeSettingsStr := utils.ConvertMapToHCL(subscribeSettings)
	publishSettingsStr := utils.ConvertMapToHCL(publishSettings)
	config := fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test" {
  address = %q
  client_certificate_file = %q
  name = %q
  outbound_member_type = %q
  subscribe_settings = %s
  publish_settings = %s
  network_view = nios_ipam_network_view.test.name
}

data "nios_misc_pxgrid_endpoint" "test" {
  filters = {
	address = nios_misc_pxgrid_endpoint.test.address
  }
}
`, address, clientCertificateFile, name, outboundMemberType, subscribeSettingsStr, publishSettingsStr)
	return strings.Join([]string{testAccBaseWithview(view), config}, "")
}

func testAccPxgridEndpointDataSourceConfigExtAttrFilters(view, address, clientCertificateFile, name, outboundMemberType string, subscribeSettings map[string]any, publishSettings map[string]any, extAttrsValue string) string {
	subscribeSettingsStr := utils.ConvertMapToHCL(subscribeSettings)
	publishSettingsStr := utils.ConvertMapToHCL(publishSettings)
	config := fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test" {
  address = %q
  client_certificate_file = %q
  name = %q
  outbound_member_type = %q
  subscribe_settings = %s
  publish_settings = %s
  network_view = nios_ipam_network_view.test.name
  extattrs = {
    Site = %q
  } 
}

data "nios_misc_pxgrid_endpoint" "test" {
  extattrfilters = {
    Site = nios_misc_pxgrid_endpoint.test.extattrs.Site
  }
}
`, address, clientCertificateFile, name, outboundMemberType, subscribeSettingsStr, publishSettingsStr, extAttrsValue)
	return strings.Join([]string{testAccBaseWithview(view), config}, "")
}
