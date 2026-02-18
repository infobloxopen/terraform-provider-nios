package misc_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

/*
// Manage misc DxlEndpoint with Basic Fields
resource "nios_misc_dxl_endpoint" "misc_dxl_endpoint_basic" {
    client_certificate_token = "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME"
    name = "NAME_REPLACE_ME"
    outbound_member_type = "OUTBOUND_MEMBER_TYPE_REPLACE_ME"
}

// Manage misc DxlEndpoint with Additional Fields
resource "nios_misc_dxl_endpoint" "misc_dxl_endpoint_with_additional_fields" {
    client_certificate_token = "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME"
    name = "NAME_REPLACE_ME"
    outbound_member_type = "OUTBOUND_MEMBER_TYPE_REPLACE_ME"

// TODO : Add additional optional fields below

    //Extensible Attributes
    extattrs = {
        Site = "location-1"
    }
}
*/

var readableAttributesForDxlEndpoint = "brokers,client_certificate_subject,client_certificate_valid_from,client_certificate_valid_to,comment,disable,extattrs,log_level,name,outbound_member_type,outbound_members,template_instance,timeout,topics,vendor_identifier,wapi_user_name"

var (
	testDataPath          = getTestDataPath()
	clientCertificateFile = filepath.Join(testDataPath, "client.pem")
	broker                = []map[string]any{
		{
			"host_name": "example.com",
		},
	}
)

func TestAccDxlEndpointResource_basic(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test"
	var v misc.DxlEndpoint
	name := acctest.RandomNameWithPrefix("dxl-endpoint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointBasicConfig(clientCertificateFile, name, "GM", broker),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", "GM"),
					resource.TestCheckResourceAttr(resourceName, "brokers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "brokers.0.host_name", "example.com"),
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

func TestAccDxlEndpointResource_disappears(t *testing.T) {
	resourceName := "nios_misc_dxl_endpoint.test"
	var v misc.DxlEndpoint
	name := acctest.RandomNameWithPrefix("dxl-endpoint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDxlEndpointDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDxlEndpointBasicConfig(clientCertificateFile, name, "GM", broker),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					testAccCheckDxlEndpointDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDxlEndpointResource_Brokers(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_brokers"
	var v misc.DxlEndpoint
	name := acctest.RandomNameWithPrefix("dxl-endpoint")
	brokersUpdated := []map[string]any{
		{
			"host_name": "example.com",
			"port":      1234,
		},
		{
			"host_name": "example2.com",
			"port":      5678,
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointBrokers(clientCertificateFile, name, "GM", broker),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "brokers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "brokers.0.host_name", "example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointBrokers(clientCertificateFile, name, "GM", brokersUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "brokers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "brokers.0.host_name", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "brokers.0.port", "1234"),
					resource.TestCheckResourceAttr(resourceName, "brokers.1.host_name", "example2.com"),
					resource.TestCheckResourceAttr(resourceName, "brokers.1.port", "5678"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_BrokersImportToken(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_brokers_import_token"
	var v misc.DxlEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointBrokersImportToken("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "BROKERS_IMPORT_TOKEN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "brokers_import_token", "BROKERS_IMPORT_TOKEN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointBrokersImportToken("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "BROKERS_IMPORT_TOKEN_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "brokers_import_token", "BROKERS_IMPORT_TOKEN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_ClientCertificateToken(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_client_certificate_token"
	var v misc.DxlEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointClientCertificateToken("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_certificate_token", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointClientCertificateToken("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_certificate_token", "CLIENT_CERTIFICATE_TOKEN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_Comment(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_comment"
	var v misc.DxlEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointComment("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointComment("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "Updated comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the object"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_Disable(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_disable"
	var v misc.DxlEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointDisable("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointDisable("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_extattrs"
	var v misc.DxlEndpoint
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointExtAttrs("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointExtAttrs("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_LogLevel(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_log_level"
	var v misc.DxlEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointLogLevel("LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "DEBUG"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "DEBUG"),
				),
			},
			{
				Config: testAccDxlEndpointLogLevel("LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "ERROR"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "ERROR"),
				),
			},
			{
				Config: testAccDxlEndpointLogLevel("LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "INFO"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "INFO"),
				),
			},
			{
				Config: testAccDxlEndpointLogLevel("LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "LOG_LEVEL_REPLACE_ME", "WARNING"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "WARNING"),
				),
			},
		},
	})
}

func TestAccDxlEndpointResource_Name(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_name"
	var v misc.DxlEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointName("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointName("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_OutboundMemberType(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_outbound_member_type"
	var v misc.DxlEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointOutboundMemberType("OUTBOUND_MEMBER_TYPE_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", "GM"),
				),
			},
			{
				Config: testAccDxlEndpointOutboundMemberType("OUTBOUND_MEMBER_TYPE_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", "MEMBER"),
				),
			},
		},
	})
}

func TestAccDxlEndpointResource_OutboundMembers(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_outbound_members"
	var v misc.DxlEndpoint
	outboundMembersVal := []string{"OUTBOUND_MEMBERS_REPLACE_ME1", "OUTBOUND_MEMBERS_REPLACE_ME2"}
	outboundMembersValUpdated := []string{"OUTBOUND_MEMBERS_REPLACE_ME1", "OUTBOUND_MEMBERS_REPLACE_ME2"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointOutboundMembers("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", outboundMembersVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_members", "OUTBOUND_MEMBERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointOutboundMembers("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", outboundMembersValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_members", "OUTBOUND_MEMBERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_TemplateInstance(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_template_instance"
	var v misc.DxlEndpoint
	templateInstanceVal := map[string]any{}
	templateInstanceValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointTemplateInstance("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", templateInstanceVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_instance", "TEMPLATE_INSTANCE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointTemplateInstance("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", templateInstanceValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_instance", "TEMPLATE_INSTANCE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccDxlEndpointResource_Timeout(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_timeout"
	var v misc.DxlEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointTimeout("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "TIMEOUT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "TIMEOUT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointTimeout("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "TIMEOUT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "TIMEOUT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccDxlEndpointResource_Topics(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_topics"
	var v misc.DxlEndpoint
	topicsVal := []string{"TOPICS_REPLACE_ME1", "TOPICS_REPLACE_ME2"}
	topicsValUpdated := []string{"TOPICS_REPLACE_ME1", "TOPICS_REPLACE_ME2"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointTopics("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", topicsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "topics", "TOPICS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointTopics("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", topicsValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "topics", "TOPICS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_VendorIdentifier(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_vendor_identifier"
	var v misc.DxlEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointVendorIdentifier("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "VENDOR_IDENTIFIER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_identifier", "VENDOR_IDENTIFIER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointVendorIdentifier("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "VENDOR_IDENTIFIER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_identifier", "VENDOR_IDENTIFIER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_WapiUserName(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_wapi_user_name"
	var v misc.DxlEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointWapiUserName("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "WAPI_USER_NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_name", "WAPI_USER_NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointWapiUserName("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "WAPI_USER_NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_name", "WAPI_USER_NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_WapiUserPassword(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_wapi_user_password"
	var v misc.DxlEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointWapiUserPassword("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "WAPI_USER_PASSWORD_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_password", "WAPI_USER_PASSWORD_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointWapiUserPassword("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME", "NAME_REPLACE_ME", "OUTBOUND_MEMBER_TYPE_REPLACE_ME", "WAPI_USER_PASSWORD_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_password", "WAPI_USER_PASSWORD_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDxlEndpointExists(ctx context.Context, resourceName string, v *misc.DxlEndpoint) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.MiscAPI.
			DxlEndpointAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDxlEndpoint).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDxlEndpointResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDxlEndpointResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDxlEndpointDestroy(ctx context.Context, v *misc.DxlEndpoint) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.MiscAPI.
			DxlEndpointAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDxlEndpoint).
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

func testAccCheckDxlEndpointDisappears(ctx context.Context, v *misc.DxlEndpoint) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.MiscAPI.
			DxlEndpointAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDxlEndpointImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccDxlEndpointBasicConfig(clientCertificateToken, name, outboundMemberType string, broker []map[string]any) string {
	brokerStr := utils.ConvertSliceOfMapsToHCL(broker)
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
	brokers = %s
}
`, clientCertificateToken, name, outboundMemberType, brokerStr)
}

func testAccDxlEndpointBrokers(clientCertificateToken string, name string, outboundMemberType string, brokers []map[string]any) string {
	brokersStr := utils.ConvertSliceOfMapsToHCL(brokers)
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_brokers" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    brokers = %s
}
`, clientCertificateToken, name, outboundMemberType, brokersStr)
}

func testAccDxlEndpointBrokersImportToken(clientCertificateToken string, name string, outboundMemberType string, brokersImportToken string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_brokers_import_token" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    brokers_import_token = %q
}
`, clientCertificateToken, name, outboundMemberType, brokersImportToken)
}

func testAccDxlEndpointClientCertificateToken(clientCertificateToken string, name string, outboundMemberType string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_client_certificate_token" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
}
`, clientCertificateToken, name, outboundMemberType)
}

func testAccDxlEndpointComment(clientCertificateToken string, name string, outboundMemberType string, comment string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_comment" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    comment = %q
}
`, clientCertificateToken, name, outboundMemberType, comment)
}

func testAccDxlEndpointDisable(clientCertificateToken string, name string, outboundMemberType string, disable string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_disable" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    disable = %q
}
`, clientCertificateToken, name, outboundMemberType, disable)
}

func testAccDxlEndpointExtAttrs(clientCertificateToken string, name string, outboundMemberType string, extAttrs map[string]string) string {
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_extattrs" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    extattrs = %s
}
`, clientCertificateToken, name, outboundMemberType, extAttrsStr)
}

func testAccDxlEndpointLogLevel(clientCertificateToken string, name string, outboundMemberType string, logLevel string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_log_level" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    log_level = %q
}
`, clientCertificateToken, name, outboundMemberType, logLevel)
}

func testAccDxlEndpointName(clientCertificateToken string, name string, outboundMemberType string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_name" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
}
`, clientCertificateToken, name, outboundMemberType)
}

func testAccDxlEndpointOutboundMemberType(clientCertificateToken string, name string, outboundMemberType string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_outbound_member_type" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
}
`, clientCertificateToken, name, outboundMemberType)
}

func testAccDxlEndpointOutboundMembers(clientCertificateToken string, name string, outboundMemberType string, outboundMembers []string) string {
	outboundMembersStr := utils.ConvertStringSliceToHCL(outboundMembers)
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_outbound_members" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    outbound_members = %s
}
`, clientCertificateToken, name, outboundMemberType, outboundMembersStr)
}

func testAccDxlEndpointTemplateInstance(clientCertificateToken string, name string, outboundMemberType string, templateInstance map[string]any) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_template_instance" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    template_instance = %s
}
`, clientCertificateToken, name, outboundMemberType, templateInstance)
}

func testAccDxlEndpointTimeout(clientCertificateToken string, name string, outboundMemberType string, timeout string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_timeout" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    timeout = %q
}
`, clientCertificateToken, name, outboundMemberType, timeout)
}

func testAccDxlEndpointTopics(clientCertificateToken string, name string, outboundMemberType string, topics []string) string {
	topicsStr := utils.ConvertStringSliceToHCL(topics)
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_topics" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    topics = %s
}
`, clientCertificateToken, name, outboundMemberType, topicsStr)
}

func testAccDxlEndpointVendorIdentifier(clientCertificateToken string, name string, outboundMemberType string, vendorIdentifier string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_vendor_identifier" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    vendor_identifier = %q
}
`, clientCertificateToken, name, outboundMemberType, vendorIdentifier)
}

func testAccDxlEndpointWapiUserName(clientCertificateToken string, name string, outboundMemberType string, wapiUserName string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_wapi_user_name" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    wapi_user_name = %q
}
`, clientCertificateToken, name, outboundMemberType, wapiUserName)
}

func testAccDxlEndpointWapiUserPassword(clientCertificateToken string, name string, outboundMemberType string, wapiUserPassword string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_wapi_user_password" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    wapi_user_password = %q
}
`, clientCertificateToken, name, outboundMemberType, wapiUserPassword)
}

func getTestDataPath() string {
	wd, err := os.Getwd()
	if err != nil {
		return "../../testdata/nios_misc_dxl_endpoint"
	}
	return filepath.Join(wd, "../../testdata/nios_misc_dxl_endpoint")
}
