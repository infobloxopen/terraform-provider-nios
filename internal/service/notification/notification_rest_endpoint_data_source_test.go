package notification_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/notification"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccNotificationRestEndpointDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_notification_rest_endpoint.test"
	resourceName := "nios_notification_rest_endpoint.test"
	var v notification.NotificationRestEndpoint
	name := acctest.RandomNameWithPrefix("notification-rest-endpoint")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNotificationRestEndpointDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNotificationRestEndpointDataSourceConfigFilters(name, outboundMemberType, uri),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					}, testAccCheckNotificationRestEndpointResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccNotificationRestEndpointDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_notification_rest_endpoint.test"
	resourceName := "nios_notification_rest_endpoint.test"
	var v notification.NotificationRestEndpoint
	name := acctest.RandomNameWithPrefix("notification-rest-endpoint")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNotificationRestEndpointDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNotificationRestEndpointDataSourceConfigExtAttrFilters(name, outboundMemberType, uri, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					}, testAccCheckNotificationRestEndpointResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckNotificationRestEndpointResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "client_certificate_subject", dataSourceName, "result.0.client_certificate_subject"),
		resource.TestCheckResourceAttrPair(resourceName, "client_certificate_token", dataSourceName, "result.0.client_certificate_token"),
		resource.TestCheckResourceAttrPair(resourceName, "client_certificate_valid_from", dataSourceName, "result.0.client_certificate_valid_from"),
		resource.TestCheckResourceAttrPair(resourceName, "client_certificate_valid_to", dataSourceName, "result.0.client_certificate_valid_to"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "log_level", dataSourceName, "result.0.log_level"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "outbound_member_type", dataSourceName, "result.0.outbound_member_type"),
		resource.TestCheckResourceAttrPair(resourceName, "outbound_members", dataSourceName, "result.0.outbound_members"),
		// resource.TestCheckResourceAttrPair(resourceName, "password", dataSourceName, "result.0.password"),
		resource.TestCheckResourceAttrPair(resourceName, "server_cert_validation", dataSourceName, "result.0.server_cert_validation"),
		resource.TestCheckResourceAttrPair(resourceName, "sync_disabled", dataSourceName, "result.0.sync_disabled"),
		resource.TestCheckResourceAttrPair(resourceName, "template_instance", dataSourceName, "result.0.template_instance"),
		resource.TestCheckResourceAttrPair(resourceName, "timeout", dataSourceName, "result.0.timeout"),
		resource.TestCheckResourceAttrPair(resourceName, "uri", dataSourceName, "result.0.uri"),
		resource.TestCheckResourceAttrPair(resourceName, "username", dataSourceName, "result.0.username"),
		resource.TestCheckResourceAttrPair(resourceName, "vendor_identifier", dataSourceName, "result.0.vendor_identifier"),
		resource.TestCheckResourceAttrPair(resourceName, "wapi_user_name", dataSourceName, "result.0.wapi_user_name"),
		// resource.TestCheckResourceAttrPair(resourceName, "wapi_user_password", dataSourceName, "result.0.wapi_user_password"),
	}
}

func testAccNotificationRestEndpointDataSourceConfigFilters(name, outboundMemberType, uri string) string {
	return fmt.Sprintf(`
resource "nios_notification_rest_endpoint" "test" {
	name = %q
	outbound_member_type = %q
	uri = %q
}

data "nios_notification_rest_endpoint" "test" {
	filters = {
		name = nios_notification_rest_endpoint.test.name
	}
}
`, name, outboundMemberType, uri)
}

func testAccNotificationRestEndpointDataSourceConfigExtAttrFilters(name, outboundMemberType, uri string, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_notification_rest_endpoint" "test" {
	name = %q
	outbound_member_type = %q
	uri = %q
	extattrs = {
		Site = %q
	} 
}

data "nios_notification_rest_endpoint" "test" {
	extattrfilters = {
		Site = nios_notification_rest_endpoint.test.extattrs.Site
	}
}
`, name, outboundMemberType, uri, extAttrsValue)
}
