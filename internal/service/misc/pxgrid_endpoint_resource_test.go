package misc_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

/*
// Manage misc PxgridEndpoint with Basic Fields
resource "nios_misc_pxgrid_endpoint" "misc_pxgrid_endpoint_basic" {
    address = "ADDRESS_REPLACE_ME"
    client_certificate_token = "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME"
    name = "NAME_REPLACE_ME"
    outbound_member_type = "OUTBOUND_MEMBER_TYPE_REPLACE_ME"
    subscribe_settings = "SUBSCRIBE_SETTINGS_REPLACE_ME"
}

// Manage misc PxgridEndpoint with Additional Fields
resource "nios_misc_pxgrid_endpoint" "misc_pxgrid_endpoint_with_additional_fields" {
    address = "ADDRESS_REPLACE_ME"
    client_certificate_token = "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME"
    name = "NAME_REPLACE_ME"
    outbound_member_type = "OUTBOUND_MEMBER_TYPE_REPLACE_ME"
    subscribe_settings = "SUBSCRIBE_SETTINGS_REPLACE_ME"

// TODO : Add additional optional fields below

    //Extensible Attributes
    extattrs = {
        Site = "location-1"
    }
}
*/

var readableAttributesForPxgridEndpoint = "address,client_certificate_subject,client_certificate_valid_from,client_certificate_valid_to,comment,disable,extattrs,log_level,name,network_view,outbound_member_type,outbound_members,publish_settings,subscribe_settings,template_instance,timeout,vendor_identifier,wapi_user_name"

var (
	testDataPath          = getTestDataPath()
	clientCertificateFile = filepath.Join(testDataPath, "client.pem")
	subscribeSettings     = map[string]any{
		"enabled_attributes": []string{"ENDPOINT_PROFILE", "DOMAINNAME", "USERNAME"},
	}
	publishSettings = map[string]any{
		"enabled_attributes": []string{"IPADDRESS"},
	}
)

func TestAccPxgridEndpointResource_basic(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test"
	var v misc.PxgridEndpoint
	view := acctest.RandomNameWithPrefix("view")
	name := acctest.RandomNameWithPrefix("pxgrid-endpoint")
	address := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointBasicConfig(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", address),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", "GM"),
					resource.TestCheckResourceAttr(resourceName, "subscribe_settings.enabled_attributes.0", "ENDPOINT_PROFILE"),
					resource.TestCheckResourceAttr(resourceName, "subscribe_settings.enabled_attributes.1", "DOMAINNAME"),
					resource.TestCheckResourceAttr(resourceName, "subscribe_settings.enabled_attributes.2", "USERNAME"),
					resource.TestCheckResourceAttr(resourceName, "publish_settings.enabled_attributes.0", "IPADDRESS"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "log_level", "WARNING"),
					resource.TestCheckResourceAttr(resourceName, "timeout", "30"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// func TestAccPxgridEndpointResource_disappears(t *testing.T) {
// 	resourceName := "nios_misc_pxgrid_endpoint.test"
// 	var v misc.PxgridEndpoint

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		CheckDestroy:             testAccCheckPxgridEndpointDestroy(context.Background(), &v),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccPxgridEndpointBasicConfig("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
// 					testAccCheckPxgridEndpointDisappears(context.Background(), &v),
// 				),
// 				ExpectNonEmptyPlan: true,
// 			},
// 		},
// 	})
// }

func TestAccPxgridEndpointResource_Address(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_address"
	var v misc.PxgridEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointAddress("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "ADDRESS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointAddress("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "ADDRESS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPxgridEndpointResource_ClientCertificateToken(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_client_certificate_token"
	var v misc.PxgridEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointClientCertificateToken("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_certificate_token", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointClientCertificateToken("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_certificate_token", "CLIENT_CERTIFICATE_TOKEN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPxgridEndpointResource_Comment(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_comment"
	var v misc.PxgridEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointComment("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointComment("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", "Updated comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the object"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPxgridEndpointResource_Disable(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_disable"
	var v misc.PxgridEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointDisable("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointDisable("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPxgridEndpointResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_extattrs"
	var v misc.PxgridEndpoint
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointExtAttrs("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointExtAttrs("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPxgridEndpointResource_LogLevel(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_log_level"
	var v misc.PxgridEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointLogLevel("LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "DEBUG"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "DEBUG"),
				),
			},
			{
				Config: testAccPxgridEndpointLogLevel("LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "ERROR"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "ERROR"),
				),
			},
			{
				Config: testAccPxgridEndpointLogLevel("LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "INFO"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "INFO"),
				),
			},
			{
				Config: testAccPxgridEndpointLogLevel("LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "WARNING"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "WARNING"),
				),
			},
		},
	})
}

func TestAccPxgridEndpointResource_Name(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_name"
	var v misc.PxgridEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointName("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointName("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPxgridEndpointResource_NetworkView(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_network_view"
	var v misc.PxgridEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointNetworkView("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", "NETWORK_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointNetworkView("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", "NETWORK_VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPxgridEndpointResource_OutboundMemberType(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_outbound_member_type"
	var v misc.PxgridEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointOutboundMemberType("OUTBOUND_MEMBER_TYPE_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", "GM"),
				),
			},
			{
				Config: testAccPxgridEndpointOutboundMemberType("OUTBOUND_MEMBER_TYPE_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", "MEMBER"),
				),
			},
		},
	})
}

func TestAccPxgridEndpointResource_OutboundMembers(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_outbound_members"
	var v misc.PxgridEndpoint
	outboundMembersVal := []string{"OUTBOUND_MEMBERS_REPLACE_ME1", "OUTBOUND_MEMBERS_REPLACE_ME2"}
	outboundMembersValUpdated := []string{"OUTBOUND_MEMBERS_REPLACE_ME1", "OUTBOUND_MEMBERS_REPLACE_ME2"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointOutboundMembers("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", outboundMembersVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_members", "OUTBOUND_MEMBERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointOutboundMembers("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", outboundMembersValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_members", "OUTBOUND_MEMBERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPxgridEndpointResource_PublishSettings(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_publish_settings"
	var v misc.PxgridEndpoint
	publishSettingsVal := map[string]any{}
	publishSettingsValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointPublishSettings("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", publishSettingsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "publish_settings", "PUBLISH_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointPublishSettings("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", publishSettingsValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "publish_settings", "PUBLISH_SETTINGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccPxgridEndpointResource_SubscribeSettings(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_subscribe_settings"
	var v misc.PxgridEndpoint
	// subscribeSettingsVal := map[string]any{}
	// subscribeSettingsValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointSubscribeSettings("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscribe_settings", "SUBSCRIBE_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointSubscribeSettings("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscribe_settings", "SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccPxgridEndpointResource_TemplateInstance(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_template_instance"
	var v misc.PxgridEndpoint
	templateInstanceVal := map[string]any{}
	templateInstanceValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointTemplateInstance("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", templateInstanceVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_instance", "TEMPLATE_INSTANCE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointTemplateInstance("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", templateInstanceValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_instance", "TEMPLATE_INSTANCE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccPxgridEndpointResource_Timeout(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_timeout"
	var v misc.PxgridEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointTimeout("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", "TIMEOUT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "TIMEOUT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointTimeout("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", "TIMEOUT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "TIMEOUT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccPxgridEndpointResource_VendorIdentifier(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_vendor_identifier"
	var v misc.PxgridEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointVendorIdentifier("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", "VENDOR_IDENTIFIER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_identifier", "VENDOR_IDENTIFIER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointVendorIdentifier("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", "VENDOR_IDENTIFIER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_identifier", "VENDOR_IDENTIFIER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPxgridEndpointResource_WapiUserName(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_wapi_user_name"
	var v misc.PxgridEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointWapiUserName("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", "WAPI_USER_NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_name", "WAPI_USER_NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointWapiUserName("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", "WAPI_USER_NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_name", "WAPI_USER_NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPxgridEndpointResource_WapiUserPassword(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_wapi_user_password"
	var v misc.PxgridEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointWapiUserPassword("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", "WAPI_USER_PASSWORD_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_password", "WAPI_USER_PASSWORD_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointWapiUserPassword("ADDRESS_REPLACE_ME", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME", "WAPI_USER_PASSWORD_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_password", "WAPI_USER_PASSWORD_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckPxgridEndpointExists(ctx context.Context, resourceName string, v *misc.PxgridEndpoint) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.MiscAPI.
			PxgridEndpointAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForPxgridEndpoint).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetPxgridEndpointResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetPxgridEndpointResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckPxgridEndpointDestroy(ctx context.Context, v *misc.PxgridEndpoint) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.MiscAPI.
			PxgridEndpointAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForPxgridEndpoint).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				// resource was deleted
				return nil
			}
			return err
		}
		return errors.New("expected to be deleted")
	}
}

func testAccCheckPxgridEndpointDisappears(ctx context.Context, v *misc.PxgridEndpoint) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.MiscAPI.
			PxgridEndpointAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccPxgridEndpointImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.Attributes["ref"] == "" {
			return "", fmt.Errorf("ref is not set")
		}
		return rs.Primary.Attributes["ref"], nil
	}
}

func testAccPxgridEndpointBasicConfig(view, address, clientCertificateToken, name, outboundMemberType string, subscribeSettings map[string]any, publishSettings map[string]any) string {
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
`, address, clientCertificateToken, name, outboundMemberType, subscribeSettingsStr, publishSettingsStr)
	return strings.Join([]string{testAccBaseWithview(view), config}, "")
}

func testAccPxgridEndpointAddress(address string, clientCertificateToken string, name string, outboundMemberType string, subscribeSettings string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_address" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %q
}
`, address, clientCertificateToken, name, outboundMemberType, subscribeSettings)
}

func testAccPxgridEndpointClientCertificateToken(address string, clientCertificateToken string, name string, outboundMemberType string, subscribeSettings string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_client_certificate_token" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %q
}
`, address, clientCertificateToken, name, outboundMemberType, subscribeSettings)
}

func testAccPxgridEndpointComment(address string, clientCertificateToken string, name string, outboundMemberType string, subscribeSettings string, comment string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_comment" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %q
    comment = %q
}
`, address, clientCertificateToken, name, outboundMemberType, subscribeSettings, comment)
}

func testAccPxgridEndpointDisable(address string, clientCertificateToken string, name string, outboundMemberType string, subscribeSettings string, disable string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_disable" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %q
    disable = %q
}
`, address, clientCertificateToken, name, outboundMemberType, subscribeSettings, disable)
}

func testAccPxgridEndpointExtAttrs(address string, clientCertificateToken string, name string, outboundMemberType string, subscribeSettings string, extAttrs map[string]string) string {
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_extattrs" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %q
    extattrs = %s
}
`, address, clientCertificateToken, name, outboundMemberType, subscribeSettings, extAttrsStr)
}

func testAccPxgridEndpointLogLevel(address string, clientCertificateToken string, name string, outboundMemberType string, subscribeSettings string, logLevel string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_log_level" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %q
    log_level = %q
}
`, address, clientCertificateToken, name, outboundMemberType, subscribeSettings, logLevel)
}

func testAccPxgridEndpointName(address string, clientCertificateToken string, name string, outboundMemberType string, subscribeSettings string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_name" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %q
}
`, address, clientCertificateToken, name, outboundMemberType, subscribeSettings)
}

func testAccPxgridEndpointNetworkView(address string, clientCertificateToken string, name string, outboundMemberType string, subscribeSettings string, networkView string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_network_view" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %q
    network_view = %q
}
`, address, clientCertificateToken, name, outboundMemberType, subscribeSettings, networkView)
}

func testAccPxgridEndpointOutboundMemberType(address string, clientCertificateToken string, name string, outboundMemberType string, subscribeSettings string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_outbound_member_type" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %q
}
`, address, clientCertificateToken, name, outboundMemberType, subscribeSettings)
}

func testAccPxgridEndpointOutboundMembers(address string, clientCertificateToken string, name string, outboundMemberType string, subscribeSettings string, outboundMembers []string) string {
	outboundMembersStr := utils.ConvertStringSliceToHCL(outboundMembers)
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_outbound_members" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %q
    outbound_members = %s
}
`, address, clientCertificateToken, name, outboundMemberType, subscribeSettings, outboundMembersStr)
}

func testAccPxgridEndpointPublishSettings(address string, clientCertificateToken string, name string, outboundMemberType string, subscribeSettings string, publishSettings map[string]any) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_publish_settings" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %q
    publish_settings = %s
}
`, address, clientCertificateToken, name, outboundMemberType, subscribeSettings, publishSettings)
}

func testAccPxgridEndpointSubscribeSettings(address string, clientCertificateToken string, name string, outboundMemberType string, subscribeSettings string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_subscribe_settings" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %q
}
`, address, clientCertificateToken, name, outboundMemberType, subscribeSettings)
}

func testAccPxgridEndpointTemplateInstance(address string, clientCertificateToken string, name string, outboundMemberType string, subscribeSettings string, templateInstance map[string]any) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_template_instance" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %q
    template_instance = %s
}
`, address, clientCertificateToken, name, outboundMemberType, subscribeSettings, templateInstance)
}

func testAccPxgridEndpointTimeout(address string, clientCertificateToken string, name string, outboundMemberType string, subscribeSettings string, timeout string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_timeout" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %q
    timeout = %q
}
`, address, clientCertificateToken, name, outboundMemberType, subscribeSettings, timeout)
}

func testAccPxgridEndpointVendorIdentifier(address string, clientCertificateToken string, name string, outboundMemberType string, subscribeSettings string, vendorIdentifier string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_vendor_identifier" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %q
    vendor_identifier = %q
}
`, address, clientCertificateToken, name, outboundMemberType, subscribeSettings, vendorIdentifier)
}

func testAccPxgridEndpointWapiUserName(address string, clientCertificateToken string, name string, outboundMemberType string, subscribeSettings string, wapiUserName string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_wapi_user_name" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %q
    wapi_user_name = %q
}
`, address, clientCertificateToken, name, outboundMemberType, subscribeSettings, wapiUserName)
}

func testAccPxgridEndpointWapiUserPassword(address string, clientCertificateToken string, name string, outboundMemberType string, subscribeSettings string, wapiUserPassword string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_wapi_user_password" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %q
    wapi_user_password = %q
}
`, address, clientCertificateToken, name, outboundMemberType, subscribeSettings, wapiUserPassword)
}

func getTestDataPath() string {
	wd, err := os.Getwd()
	if err != nil {
		return "../../testdata/nios_misc_pxgrid_endpoint"
	}
	return filepath.Join(wd, "../../testdata/nios_misc_pxgrid_endpoint")
}

func testAccBaseWithview(name string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_view" "test" {
	name = %q
}
`, name)
}
