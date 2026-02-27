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

// OBJECTS TO BE PRESENT IN GRID FOR TESTS
// Pxgrid Templates - PxgrdiSessionTemplate, PxgrdiSessionTemplate_Alt
// Grid Master Candidate - infoblox.grid_master_candidate1

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
	view := acctest.RandomNameWithPrefix("network-view")
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

func TestAccPxgridEndpointResource_disappears(t *testing.T) {
	resourceName := "nios_misc_pxgrid_endpoint.test"
	var v misc.PxgridEndpoint
	view := acctest.RandomNameWithPrefix("network-view")
	name := acctest.RandomNameWithPrefix("pxgrid-endpoint")
	address := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPxgridEndpointDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccPxgridEndpointBasicConfig(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					testAccCheckPxgridEndpointDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccPxgridEndpointResource_Address(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_address"
	var v misc.PxgridEndpoint
	view := acctest.RandomNameWithPrefix("network-view")
	name := acctest.RandomNameWithPrefix("pxgrid-endpoint")
	address := acctest.RandomIP()
	addressUpdate := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointAddress(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", address),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointAddress(view, addressUpdate, clientCertificateFile, name, "GM", subscribeSettings, publishSettings),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", addressUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPxgridEndpointResource_ClientCertificateToken(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_client_certificate_token"
	var v misc.PxgridEndpoint
	view := acctest.RandomNameWithPrefix("network-view")
	name := acctest.RandomNameWithPrefix("pxgrid-endpoint")
	address := acctest.RandomIP()
	updatedClientCertificateFile := filepath.Join(testDataPath, "client_updated.pem")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointClientCertificateToken(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointClientCertificateToken(view, address, updatedClientCertificateFile, name, "GM", subscribeSettings, publishSettings),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPxgridEndpointResource_Comment(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_comment"
	var v misc.PxgridEndpoint
	view := acctest.RandomNameWithPrefix("network-view")
	name := acctest.RandomNameWithPrefix("pxgrid-endpoint")
	address := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointComment(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings, "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointComment(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings, "Updated comment for the object"),
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
	view := acctest.RandomNameWithPrefix("network-view")
	name := acctest.RandomNameWithPrefix("pxgrid-endpoint")
	address := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointDisable(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointDisable(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings, "false"),
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
	view := acctest.RandomNameWithPrefix("network-view")
	name := acctest.RandomNameWithPrefix("pxgrid-endpoint")
	address := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointExtAttrs(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointExtAttrs(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings, map[string]string{
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
	view := acctest.RandomNameWithPrefix("network-view")
	name := acctest.RandomNameWithPrefix("pxgrid-endpoint")
	address := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointLogLevel(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings, "DEBUG"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "DEBUG"),
				),
			},
			{
				Config: testAccPxgridEndpointLogLevel(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings, "ERROR"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "ERROR"),
				),
			},
			{
				Config: testAccPxgridEndpointLogLevel(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings, "INFO"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "INFO"),
				),
			},
			{
				Config: testAccPxgridEndpointLogLevel(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings, "WARNING"),
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
	view := acctest.RandomNameWithPrefix("network-view")
	name := acctest.RandomNameWithPrefix("pxgrid-endpoint")
	nameUpdate := acctest.RandomNameWithPrefix("pxgrid-endpoint-updated")
	address := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointName(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointName(view, address, clientCertificateFile, nameUpdate, "GM", subscribeSettings, publishSettings),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPxgridEndpointResource_NetworkView(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_network_view"
	var v misc.PxgridEndpoint
	view := acctest.RandomNameWithPrefix("network-view")
	viewUpdate := acctest.RandomNameWithPrefix("view-updated")
	name := acctest.RandomNameWithPrefix("pxgrid-endpoint")
	address := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointNetworkView(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", view),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointNetworkView(viewUpdate, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings, viewUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", viewUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPxgridEndpointResource_OutboundMemberType(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_outbound_member_type"
	var v misc.PxgridEndpoint
	view := acctest.RandomNameWithPrefix("network-view")
	name := acctest.RandomNameWithPrefix("pxgrid-endpoint")
	address := acctest.RandomIP()
	outboundMembers := []string{"infoblox.grid_master_candidate1"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointOutboundMemberType(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", "GM"),
				),
			},
			{
				Config: testAccPxgridEndpointOutboundMemberTypeUpdate(view, address, clientCertificateFile, name, "MEMBER", subscribeSettings, publishSettings, outboundMembers),
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
	outboundMembersVal := []string{"infoblox.grid_master_candidate1"}
	outboundMembersValUpdated := []string{"infoblox.grid_master_candidate2"}
	view := acctest.RandomNameWithPrefix("network-view")
	name := acctest.RandomNameWithPrefix("pxgrid-endpoint")
	address := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointOutboundMembers(view, address, clientCertificateFile, name, "MEMBER", subscribeSettings, publishSettings, outboundMembersVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_members.0", outboundMembersVal[0]),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointOutboundMembers(view, address, clientCertificateFile, name, "MEMBER", subscribeSettings, publishSettings, outboundMembersValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_members.0", outboundMembersValUpdated[0]),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPxgridEndpointResource_PublishSettings(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_publish_settings"
	var v misc.PxgridEndpoint
	publishSettingsVal := map[string]any{
		"enabled_attributes": []string{"IPADDRESS"},
	}
	publishSettingsValUpdated := map[string]any{
		"enabled_attributes": []string{"FINGERPRINT", "IPADDRESS", "HOSTNAME"},
	}
	view := acctest.RandomNameWithPrefix("network-view")
	name := acctest.RandomNameWithPrefix("pxgrid-endpoint")
	address := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointPublishSettings(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettingsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "publish_settings.enabled_attributes.0", "IPADDRESS"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointPublishSettings(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettingsValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "publish_settings.enabled_attributes.0", "FINGERPRINT"),
					resource.TestCheckResourceAttr(resourceName, "publish_settings.enabled_attributes.1", "IPADDRESS"),
					resource.TestCheckResourceAttr(resourceName, "publish_settings.enabled_attributes.2", "HOSTNAME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccPxgridEndpointResource_SubscribeSettings(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_subscribe_settings"
	var v misc.PxgridEndpoint
	subscribeSettingsVal := map[string]any{
		"enabled_attributes": []string{"ENDPOINT_PROFILE", "DOMAINNAME"},
	}
	subscribeSettingsValUpdated := map[string]any{
		"enabled_attributes": []string{"ENDPOINT_PROFILE", "DOMAINNAME", "USERNAME"},
	}
	view := acctest.RandomNameWithPrefix("network-view")
	name := acctest.RandomNameWithPrefix("pxgrid-endpoint")
	address := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointSubscribeSettings(view, address, clientCertificateFile, name, "GM", subscribeSettingsVal, publishSettings),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscribe_settings.enabled_attributes.0", "ENDPOINT_PROFILE"),
					resource.TestCheckResourceAttr(resourceName, "subscribe_settings.enabled_attributes.1", "DOMAINNAME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointSubscribeSettings(view, address, clientCertificateFile, name, "GM", subscribeSettingsValUpdated, publishSettings),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscribe_settings.enabled_attributes.0", "ENDPOINT_PROFILE"),
					resource.TestCheckResourceAttr(resourceName, "subscribe_settings.enabled_attributes.1", "DOMAINNAME"),
					resource.TestCheckResourceAttr(resourceName, "subscribe_settings.enabled_attributes.2", "USERNAME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccPxgridEndpointResource_TemplateInstance(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_template_instance"
	var v misc.PxgridEndpoint
	templateInstanceVal := map[string]any{
		"template": "PxgrdiSessionTemplate",
	}
	templateInstanceValUpdated := map[string]any{
		"template": "PxgridSessionTemplate_Alt",
	}
	view := acctest.RandomNameWithPrefix("network-view")
	name := acctest.RandomNameWithPrefix("pxgrid-endpoint")
	address := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointTemplateInstance(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings, templateInstanceVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_instance.template", "PxgrdiSessionTemplate"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointTemplateInstance(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings, templateInstanceValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_instance.template", "PxgridSessionTemplate_Alt"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccPxgridEndpointResource_Timeout(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_timeout"
	var v misc.PxgridEndpoint
	view := acctest.RandomNameWithPrefix("network-view")
	name := acctest.RandomNameWithPrefix("pxgrid-endpoint")
	address := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointTimeout(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings, "60"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "60"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointTimeout(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings, "120"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "120"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccPxgridEndpointResource_VendorIdentifier(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_vendor_identifier"
	var v misc.PxgridEndpoint
	view := acctest.RandomNameWithPrefix("network-view")
	name := acctest.RandomNameWithPrefix("pxgrid-endpoint")
	address := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointVendorIdentifier(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings, "pxgrid"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_identifier", "pxgrid"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointVendorIdentifier(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_identifier", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPxgridEndpointResource_WapiUserName(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_wapi_user_name"
	var v misc.PxgridEndpoint
	view := acctest.RandomNameWithPrefix("network-view")
	name := acctest.RandomNameWithPrefix("pxgrid-endpoint")
	address := acctest.RandomIP()
	wapiUserName := "wapi-user"
	wapiUserNameUpdate := "wapi-user-updated"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointWapiUserName(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings, wapiUserName, "password"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_name", wapiUserName),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointWapiUserName(view, address, clientCertificateFile, name, "GM", subscribeSettings, publishSettings, wapiUserNameUpdate, "password"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_name", wapiUserNameUpdate),
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

func testAccPxgridEndpointAddress(view, address, clientCertificateFile, name, outboundMemberType string, subscribeSettings map[string]any, publishSettings map[string]any) string {
	subscribeSettingsStr := utils.ConvertMapToHCL(subscribeSettings)
	publishSettingsStr := utils.ConvertMapToHCL(publishSettings)
	config := fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_address" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %s
	publish_settings = %s
	network_view = nios_ipam_network_view.test.name
}
`, address, clientCertificateFile, name, outboundMemberType, subscribeSettingsStr, publishSettingsStr)
	return strings.Join([]string{testAccBaseWithview(view), config}, "")
}

func testAccPxgridEndpointClientCertificateToken(view, address, clientCertificateFile, name, outboundMemberType string, subscribeSettings map[string]any, publishSettings map[string]any) string {
	subscribeSettingsStr := utils.ConvertMapToHCL(subscribeSettings)
	publishSettingsStr := utils.ConvertMapToHCL(publishSettings)
	config := fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_client_certificate_token" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %s
	publish_settings = %s
	network_view = nios_ipam_network_view.test.name
}
`, address, clientCertificateFile, name, outboundMemberType, subscribeSettingsStr, publishSettingsStr)
	return strings.Join([]string{testAccBaseWithview(view), config}, "")
}

func testAccPxgridEndpointComment(view, address, clientCertificateFile, name, outboundMemberType string, subscribeSettings map[string]any, publishSettings map[string]any, comment string) string {
	subscribeSettingsStr := utils.ConvertMapToHCL(subscribeSettings)
	publishSettingsStr := utils.ConvertMapToHCL(publishSettings)
	config := fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_comment" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %s
	publish_settings = %s
	network_view = nios_ipam_network_view.test.name
    comment = %q
}
`, address, clientCertificateFile, name, outboundMemberType, subscribeSettingsStr, publishSettingsStr, comment)
	return strings.Join([]string{testAccBaseWithview(view), config}, "")
}

func testAccPxgridEndpointDisable(view, address, clientCertificateFile, name, outboundMemberType string, subscribeSettings map[string]any, publishSettings map[string]any, disable string) string {
	subscribeSettingsStr := utils.ConvertMapToHCL(subscribeSettings)
	publishSettingsStr := utils.ConvertMapToHCL(publishSettings)
	config := fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_disable" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %s
	publish_settings = %s
	network_view = nios_ipam_network_view.test.name
    disable = %q
}
`, address, clientCertificateFile, name, outboundMemberType, subscribeSettingsStr, publishSettingsStr, disable)
	return strings.Join([]string{testAccBaseWithview(view), config}, "")
}

func testAccPxgridEndpointExtAttrs(view, address, clientCertificateFile, name, outboundMemberType string, subscribeSettings map[string]any, publishSettings map[string]any, extAttrs map[string]string) string {
	subscribeSettingsStr := utils.ConvertMapToHCL(subscribeSettings)
	publishSettingsStr := utils.ConvertMapToHCL(publishSettings)
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	config := fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_extattrs" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %s
	publish_settings = %s
	network_view = nios_ipam_network_view.test.name
    extattrs = %s
}
`, address, clientCertificateFile, name, outboundMemberType, subscribeSettingsStr, publishSettingsStr, extAttrsStr)
	return strings.Join([]string{testAccBaseWithview(view), config}, "")
}

func testAccPxgridEndpointLogLevel(view, address, clientCertificateFile, name, outboundMemberType string, subscribeSettings map[string]any, publishSettings map[string]any, logLevel string) string {
	subscribeSettingsStr := utils.ConvertMapToHCL(subscribeSettings)
	publishSettingsStr := utils.ConvertMapToHCL(publishSettings)
	config := fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_log_level" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %s
	publish_settings = %s
	network_view = nios_ipam_network_view.test.name
    log_level = %q
}
`, address, clientCertificateFile, name, outboundMemberType, subscribeSettingsStr, publishSettingsStr, logLevel)
	return strings.Join([]string{testAccBaseWithview(view), config}, "")
}

func testAccPxgridEndpointName(view, address, clientCertificateFile, name, outboundMemberType string, subscribeSettings map[string]any, publishSettings map[string]any) string {
	subscribeSettingsStr := utils.ConvertMapToHCL(subscribeSettings)
	publishSettingsStr := utils.ConvertMapToHCL(publishSettings)
	config := fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_name" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %s
	publish_settings = %s
	network_view = nios_ipam_network_view.test.name
}
`, address, clientCertificateFile, name, outboundMemberType, subscribeSettingsStr, publishSettingsStr)
	return strings.Join([]string{testAccBaseWithview(view), config}, "")
}

func testAccPxgridEndpointNetworkView(view, address, clientCertificateFile, name, outboundMemberType string, subscribeSettings map[string]any, publishSettings map[string]any, networkView string) string {
	subscribeSettingsStr := utils.ConvertMapToHCL(subscribeSettings)
	publishSettingsStr := utils.ConvertMapToHCL(publishSettings)
	config := fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_network_view" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %s
	publish_settings = %s
	network_view = %q
}
`, address, clientCertificateFile, name, outboundMemberType, subscribeSettingsStr, publishSettingsStr, networkView)
	return strings.Join([]string{testAccBaseWithview(view), config}, "")
}

func testAccPxgridEndpointOutboundMemberType(view, address, clientCertificateFile, name, outboundMemberType string, subscribeSettings map[string]any, publishSettings map[string]any) string {
	subscribeSettingsStr := utils.ConvertMapToHCL(subscribeSettings)
	publishSettingsStr := utils.ConvertMapToHCL(publishSettings)
	config := fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_outbound_member_type" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %s
	publish_settings = %s
	network_view = nios_ipam_network_view.test.name
}
`, address, clientCertificateFile, name, outboundMemberType, subscribeSettingsStr, publishSettingsStr)
	return strings.Join([]string{testAccBaseWithview(view), config}, "")
}

func testAccPxgridEndpointOutboundMemberTypeUpdate(view, address, clientCertificateFile, name, outboundMemberType string, subscribeSettings map[string]any, publishSettings map[string]any, outboundMembers []string) string {
	subscribeSettingsStr := utils.ConvertMapToHCL(subscribeSettings)
	publishSettingsStr := utils.ConvertMapToHCL(publishSettings)
	outboundMembersStr := utils.ConvertStringSliceToHCL(outboundMembers)
	config := fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_outbound_member_type" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
	outbound_members = %s
    subscribe_settings = %s
	publish_settings = %s
	network_view = nios_ipam_network_view.test.name
}
`, address, clientCertificateFile, name, outboundMemberType, outboundMembersStr, subscribeSettingsStr, publishSettingsStr)
	return strings.Join([]string{testAccBaseWithview(view), config}, "")
}

func testAccPxgridEndpointOutboundMembers(view, address, clientCertificateFile, name, outboundMemberType string, subscribeSettings map[string]any, publishSettings map[string]any, outboundMembers []string) string {
	subscribeSettingsStr := utils.ConvertMapToHCL(subscribeSettings)
	publishSettingsStr := utils.ConvertMapToHCL(publishSettings)
	outboundMembersStr := utils.ConvertStringSliceToHCL(outboundMembers)
	config := fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_outbound_members" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %s
	publish_settings = %s
	network_view = nios_ipam_network_view.test.name
    outbound_members = %s
}
`, address, clientCertificateFile, name, outboundMemberType, subscribeSettingsStr, publishSettingsStr, outboundMembersStr)
	return strings.Join([]string{testAccBaseWithview(view), config}, "")
}

func testAccPxgridEndpointPublishSettings(view, address, clientCertificateFile, name, outboundMemberType string, subscribeSettings map[string]any, publishSettings map[string]any) string {
	subscribeSettingsStr := utils.ConvertMapToHCL(subscribeSettings)
	publishSettingsStr := utils.ConvertMapToHCL(publishSettings)
	config := fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_publish_settings" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %s
	publish_settings = %s
	network_view = nios_ipam_network_view.test.name
}
`, address, clientCertificateFile, name, outboundMemberType, subscribeSettingsStr, publishSettingsStr)
	return strings.Join([]string{testAccBaseWithview(view), config}, "")
}

func testAccPxgridEndpointSubscribeSettings(view, address, clientCertificateFile, name, outboundMemberType string, subscribeSettings map[string]any, publishSettings map[string]any) string {
	subscribeSettingsStr := utils.ConvertMapToHCL(subscribeSettings)
	publishSettingsStr := utils.ConvertMapToHCL(publishSettings)
	config := fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_subscribe_settings" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %s
	publish_settings = %s
	network_view = nios_ipam_network_view.test.name
}
`, address, clientCertificateFile, name, outboundMemberType, subscribeSettingsStr, publishSettingsStr)
	return strings.Join([]string{testAccBaseWithview(view), config}, "")
}

func testAccPxgridEndpointTemplateInstance(view, address, clientCertificateFile, name, outboundMemberType string, subscribeSettings map[string]any, publishSettings map[string]any, templateInstance map[string]any) string {
	subscribeSettingsStr := utils.ConvertMapToHCL(subscribeSettings)
	publishSettingsStr := utils.ConvertMapToHCL(publishSettings)
	templateInstanceStr := utils.ConvertMapToHCL(templateInstance)
	config := fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_template_instance" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %s
	publish_settings = %s
	network_view = nios_ipam_network_view.test.name
    template_instance = %s
}
`, address, clientCertificateFile, name, outboundMemberType, subscribeSettingsStr, publishSettingsStr, templateInstanceStr)
	return strings.Join([]string{testAccBaseWithview(view), config}, "")
}

func testAccPxgridEndpointTimeout(view, address, clientCertificateFile, name, outboundMemberType string, subscribeSettings map[string]any, publishSettings map[string]any, timeout string) string {
	subscribeSettingsStr := utils.ConvertMapToHCL(subscribeSettings)
	publishSettingsStr := utils.ConvertMapToHCL(publishSettings)
	config := fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_timeout" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %s
	publish_settings = %s
	network_view = nios_ipam_network_view.test.name
    timeout = %q
}
`, address, clientCertificateFile, name, outboundMemberType, subscribeSettingsStr, publishSettingsStr, timeout)
	return strings.Join([]string{testAccBaseWithview(view), config}, "")
}

func testAccPxgridEndpointVendorIdentifier(view, address, clientCertificateFile, name, outboundMemberType string, subscribeSettings map[string]any, publishSettings map[string]any, vendorIdentifier string) string {
	subscribeSettingsStr := utils.ConvertMapToHCL(subscribeSettings)
	publishSettingsStr := utils.ConvertMapToHCL(publishSettings)
	config := fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_vendor_identifier" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %s
	publish_settings = %s
	network_view = nios_ipam_network_view.test.name
    vendor_identifier = %q
}
`, address, clientCertificateFile, name, outboundMemberType, subscribeSettingsStr, publishSettingsStr, vendorIdentifier)
	return strings.Join([]string{testAccBaseWithview(view), config}, "")
}

func testAccPxgridEndpointWapiUserName(view, address, clientCertificateFile, name, outboundMemberType string, subscribeSettings map[string]any, publishSettings map[string]any, wapiUserName, wapiUserPassword string) string {
	subscribeSettingsStr := utils.ConvertMapToHCL(subscribeSettings)
	publishSettingsStr := utils.ConvertMapToHCL(publishSettings)
	config := fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_wapi_user_name" {
    address = %q
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    subscribe_settings = %s
	publish_settings = %s
	network_view = nios_ipam_network_view.test.name
    wapi_user_name = %q
	wapi_user_password = %q
}
`, address, clientCertificateFile, name, outboundMemberType, subscribeSettingsStr, publishSettingsStr, wapiUserName, wapiUserPassword)
	return strings.Join([]string{testAccBaseWithview(view), config}, "")
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
