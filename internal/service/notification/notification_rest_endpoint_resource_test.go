package notification_test

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

	"github.com/infobloxopen/infoblox-nios-go-client/notification"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// OBJECTS TO BE PRESENT IN GRID FOR TESTS
// Notification Rest Template : "REST API Template", "REST_API_Session_Template_3"
// Grid Master Candidate : "infoblox.grid_master_candidate1", "infoblox.grid_master_candidate2"

var readableAttributesForNotificationRestEndpoint = "client_certificate_subject,client_certificate_valid_from,client_certificate_valid_to,comment,extattrs,log_level,name,outbound_member_type,outbound_members,server_cert_validation,sync_disabled,template_instance,timeout,uri,username,vendor_identifier,wapi_user_name"

var (
	uri                = "https://example.com"
	outboundMemberType = "GM"
)

func TestAccNotificationRestEndpointResource_basic(t *testing.T) {
	var resourceName = "nios_notification_rest_endpoint.test"
	var v notification.NotificationRestEndpoint
	name := acctest.RandomNameWithPrefix("notification-rest-endpoint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRestEndpointBasicConfig(name, outboundMemberType, uri),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", outboundMemberType),
					resource.TestCheckResourceAttr(resourceName, "uri", uri),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "log_level", "WARNING"),
					resource.TestCheckResourceAttr(resourceName, "server_cert_validation", "CA_CERT"),
					resource.TestCheckResourceAttr(resourceName, "sync_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "timeout", "30"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRestEndpointResource_disappears(t *testing.T) {
	resourceName := "nios_notification_rest_endpoint.test"
	var v notification.NotificationRestEndpoint
	name := acctest.RandomNameWithPrefix("notification-rest-endpoint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNotificationRestEndpointDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNotificationRestEndpointBasicConfig(name, outboundMemberType, uri),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					testAccCheckNotificationRestEndpointDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNotificationRestEndpointResource_Comment(t *testing.T) {
	var resourceName = "nios_notification_rest_endpoint.test_comment"
	var v notification.NotificationRestEndpoint
	name := acctest.RandomNameWithPrefix("notification-rest-endpoint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRestEndpointComment(name, outboundMemberType, uri, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRestEndpointComment(name, outboundMemberType, uri, "This is a updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRestEndpointResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_notification_rest_endpoint.test_extattrs"
	var v notification.NotificationRestEndpoint
	name := acctest.RandomNameWithPrefix("notification-rest-endpoint")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRestEndpointExtAttrs(name, outboundMemberType, uri, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRestEndpointExtAttrs(name, outboundMemberType, uri, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRestEndpointResource_LogLevel(t *testing.T) {
	var resourceName = "nios_notification_rest_endpoint.test_log_level"
	var v notification.NotificationRestEndpoint
	name := acctest.RandomNameWithPrefix("notification-rest-endpoint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRestEndpointLogLevel(name, outboundMemberType, uri, "DEBUG"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "DEBUG"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRestEndpointLogLevel(name, outboundMemberType, uri, "ERROR"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "ERROR"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRestEndpointResource_Name(t *testing.T) {
	var resourceName = "nios_notification_rest_endpoint.test_name"
	var v notification.NotificationRestEndpoint
	name := acctest.RandomNameWithPrefix("notification-rest-endpoint")
	updatedName := acctest.RandomNameWithPrefix("notification-rest-endpoint-updated")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRestEndpointName(name, outboundMemberType, uri),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRestEndpointName(updatedName, outboundMemberType, uri),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRestEndpointResource_ClientCertificateFile(t *testing.T) {
	var resourceName = "nios_notification_rest_endpoint.test_client_certificate_file"
	var v notification.NotificationRestEndpoint
	name := acctest.RandomNameWithPrefix("notification-rest-endpoint")
	testDataPath := getTestDataPath()
	clientCertificateFile := filepath.Join(testDataPath, "dummy-bundle.pem")
	updatedClientCertificateFile := filepath.Join(testDataPath, "dummy-bundle2.pem")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRestEndpointClientCertificateFile(name, outboundMemberType, uri, clientCertificateFile),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRestEndpointClientCertificateFile(name, outboundMemberType, uri, updatedClientCertificateFile),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRestEndpointResource_OutboundMemberType(t *testing.T) {
	var resourceName = "nios_notification_rest_endpoint.test_outbound_member_type"
	var v notification.NotificationRestEndpoint
	name := acctest.RandomNameWithPrefix("notification-rest-endpoint")
	outboundMembers := []string{"infoblox.grid_master_candidate1"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRestEndpointOutboundMemberType(name, outboundMemberType, uri),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", "GM"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRestEndpointOutboundMemberTypeUpdate(name, "MEMBER", uri, outboundMembers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", "MEMBER"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRestEndpointResource_OutboundMembers(t *testing.T) {
	var resourceName = "nios_notification_rest_endpoint.test_outbound_members"
	var v notification.NotificationRestEndpoint
	name := acctest.RandomNameWithPrefix("notification-rest-endpoint")
	outboundMembers := []string{"infoblox.grid_master_candidate1"}
	updatedOutboundMembers := []string{"infoblox.grid_master_candidate2"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRestEndpointOutboundMembers(name, "MEMBER", uri, outboundMembers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_members.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "outbound_members.0", "infoblox.grid_master_candidate1"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRestEndpointOutboundMembers(name, "MEMBER", uri, updatedOutboundMembers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_members.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "outbound_members.0", "infoblox.grid_master_candidate2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRestEndpointResource_ServerCertValidation(t *testing.T) {
	var resourceName = "nios_notification_rest_endpoint.test_server_cert_validation"
	var v notification.NotificationRestEndpoint
	name := acctest.RandomNameWithPrefix("notification-rest-endpoint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRestEndpointServerCertValidation(name, outboundMemberType, uri, "CA_CERT_NO_HOSTNAME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "server_cert_validation", "CA_CERT_NO_HOSTNAME"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRestEndpointServerCertValidation(name, outboundMemberType, uri, "NO_VALIDATION"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "server_cert_validation", "NO_VALIDATION"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRestEndpointResource_SyncDisabled(t *testing.T) {
	var resourceName = "nios_notification_rest_endpoint.test_sync_disabled"
	var v notification.NotificationRestEndpoint
	name := acctest.RandomNameWithPrefix("notification-rest-endpoint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRestEndpointSyncDisabled(name, outboundMemberType, uri, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sync_disabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRestEndpointSyncDisabled(name, outboundMemberType, uri, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sync_disabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRestEndpointResource_TemplateInstance(t *testing.T) {
	var resourceName = "nios_notification_rest_endpoint.test_template_instance"
	var v notification.NotificationRestEndpoint
	name := acctest.RandomNameWithPrefix("notification-rest-endpoint")
	templateInstance := map[string]any{
		"parameters": []map[string]any{
			{
				"name":   "SPECIAL1",
				"syntax": "INT",
			},
			{
				"name":   "SPECIAL2",
				"syntax": "STR",
			},
			{
				"name":   "SPECIAL3",
				"syntax": "BOOL",
			},
		},
		"template": "REST API Template",
	}

	updatedTemplateInstance := map[string]any{
		"template": "REST_API_Session_Template_3",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRestEndpointTemplateInstance(name, outboundMemberType, uri, templateInstance),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_instance.parameters.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "template_instance.parameters.0.name", "SPECIAL1"),
					resource.TestCheckResourceAttr(resourceName, "template_instance.parameters.0.syntax", "INT"),
					resource.TestCheckResourceAttr(resourceName, "template_instance.parameters.1.name", "SPECIAL2"),
					resource.TestCheckResourceAttr(resourceName, "template_instance.parameters.1.syntax", "STR"),
					resource.TestCheckResourceAttr(resourceName, "template_instance.parameters.2.name", "SPECIAL3"),
					resource.TestCheckResourceAttr(resourceName, "template_instance.parameters.2.syntax", "BOOL"),
					resource.TestCheckResourceAttr(resourceName, "template_instance.template", "REST API Template"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRestEndpointTemplateInstance(name, outboundMemberType, uri, updatedTemplateInstance),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_instance.template", "REST_API_Session_Template_3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRestEndpointResource_Timeout(t *testing.T) {
	var resourceName = "nios_notification_rest_endpoint.test_timeout"
	var v notification.NotificationRestEndpoint
	name := acctest.RandomNameWithPrefix("notification-rest-endpoint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRestEndpointTimeout(name, outboundMemberType, uri, "100"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "100"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRestEndpointTimeout(name, outboundMemberType, uri, "200"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRestEndpointResource_Uri(t *testing.T) {
	var resourceName = "nios_notification_rest_endpoint.test_uri"
	var v notification.NotificationRestEndpoint
	name := acctest.RandomNameWithPrefix("notification-rest-endpoint")
	updatedUri := "https://example-updated.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRestEndpointUri(name, outboundMemberType, uri),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "uri", uri),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRestEndpointUri(name, outboundMemberType, updatedUri),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "uri", updatedUri),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRestEndpointResource_Username(t *testing.T) {
	var resourceName = "nios_notification_rest_endpoint.test_username"
	var v notification.NotificationRestEndpoint
	name := acctest.RandomNameWithPrefix("notification-rest-endpoint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRestEndpointUsername(name, outboundMemberType, uri, "example_username", "example_password"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "username", "example_username"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRestEndpointUsername(name, outboundMemberType, uri, "example_username_updated", "example_password_updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "username", "example_username_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRestEndpointResource_VendorIdentifier(t *testing.T) {
	var resourceName = "nios_notification_rest_endpoint.test_vendor_identifier"
	var v notification.NotificationRestEndpoint
	name := acctest.RandomNameWithPrefix("notification-rest-endpoint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRestEndpointVendorIdentifier(name, outboundMemberType, uri, "WAPI"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_identifier", "WAPI"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRestEndpointVendorIdentifier(name, outboundMemberType, uri, "CarbonBlack"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_identifier", "CarbonBlack"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNotificationRestEndpointResource_WapiUserName(t *testing.T) {
	var resourceName = "nios_notification_rest_endpoint.test_wapi_user_name"
	var v notification.NotificationRestEndpoint
	name := acctest.RandomNameWithPrefix("notification-rest-endpoint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNotificationRestEndpointWapiUserName(name, outboundMemberType, uri, "example_wapi_username", "example_wapi_password"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_name", "example_wapi_username"),
				),
			},
			// Update and Read
			{
				Config: testAccNotificationRestEndpointWapiUserName(name, outboundMemberType, uri, "example_wapi_username_updated", "example_wapi_password_updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNotificationRestEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_name", "example_wapi_username_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckNotificationRestEndpointExists(ctx context.Context, resourceName string, v *notification.NotificationRestEndpoint) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.NotificationAPI.
			NotificationRestEndpointAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForNotificationRestEndpoint).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetNotificationRestEndpointResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetNotificationRestEndpointResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckNotificationRestEndpointDestroy(ctx context.Context, v *notification.NotificationRestEndpoint) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.NotificationAPI.
			NotificationRestEndpointAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForNotificationRestEndpoint).
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

func testAccCheckNotificationRestEndpointDisappears(ctx context.Context, v *notification.NotificationRestEndpoint) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.NotificationAPI.
			NotificationRestEndpointAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccNotificationRestEndpointBasicConfig(name, outboundMemberType, uri string) string {
	return fmt.Sprintf(`
resource "nios_notification_rest_endpoint" "test" {
    name = %q
    outbound_member_type = %q
    uri = %q
}
`, name, outboundMemberType, uri)
}

func testAccNotificationRestEndpointComment(name string, outboundMemberType string, uri string, comment string) string {
	return fmt.Sprintf(`
resource "nios_notification_rest_endpoint" "test_comment" {
    name = %q
    outbound_member_type = %q
    uri = %q
    comment = %q
}
`, name, outboundMemberType, uri, comment)
}

func testAccNotificationRestEndpointExtAttrs(name string, outboundMemberType string, uri string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_notification_rest_endpoint" "test_extattrs" {
    name = %q
    outbound_member_type = %q
    uri = %q
    extattrs = %s
}
`, name, outboundMemberType, uri, extattrsStr)
}

func testAccNotificationRestEndpointLogLevel(name string, outboundMemberType string, uri string, logLevel string) string {
	return fmt.Sprintf(`
resource "nios_notification_rest_endpoint" "test_log_level" {
    name = %q
    outbound_member_type = %q
    uri = %q
    log_level = %q
}
`, name, outboundMemberType, uri, logLevel)
}

func testAccNotificationRestEndpointName(name string, outboundMemberType string, uri string) string {
	return fmt.Sprintf(`
resource "nios_notification_rest_endpoint" "test_name" {
    name = %q
    outbound_member_type = %q
    uri = %q
}
`, name, outboundMemberType, uri)
}

func testAccNotificationRestEndpointClientCertificateFile(name string, outboundMemberType string, uri string, clientCertificateFile string) string {
	return fmt.Sprintf(`
resource "nios_notification_rest_endpoint" "test_client_certificate_file" {
    name = %q
    outbound_member_type = %q
    uri = %q
    client_certificate_file = %q
}
`, name, outboundMemberType, uri, clientCertificateFile)
}
func testAccNotificationRestEndpointOutboundMemberType(name string, outboundMemberType string, uri string) string {
	return fmt.Sprintf(`
resource "nios_notification_rest_endpoint" "test_outbound_member_type" {
    name = %q
    outbound_member_type = %q
    uri = %q
}
`, name, outboundMemberType, uri)
}

func testAccNotificationRestEndpointOutboundMemberTypeUpdate(name string, outboundMemberType string, uri string, outboundMembers []string) string {
	outboundMembersHCL := utils.ConvertStringSliceToHCL(outboundMembers)
	return fmt.Sprintf(`
resource "nios_notification_rest_endpoint" "test_outbound_member_type" {
    name = %q
    outbound_member_type = %q
    uri = %q
	outbound_members = %s
}
`, name, outboundMemberType, uri, outboundMembersHCL)
}

func testAccNotificationRestEndpointOutboundMembers(name string, outboundMemberType string, uri string, outboundMembers []string) string {
	outboundMembersHCL := utils.ConvertStringSliceToHCL(outboundMembers)
	return fmt.Sprintf(`
resource "nios_notification_rest_endpoint" "test_outbound_members" {
    name = %q
    outbound_member_type = %q
    uri = %q
    outbound_members = %s
}
`, name, outboundMemberType, uri, outboundMembersHCL)
}

func testAccNotificationRestEndpointServerCertValidation(name string, outboundMemberType string, uri string, serverCertValidation string) string {
	return fmt.Sprintf(`
resource "nios_notification_rest_endpoint" "test_server_cert_validation" {
    name = %q
    outbound_member_type = %q
    uri = %q
    server_cert_validation = %q
}
`, name, outboundMemberType, uri, serverCertValidation)
}

func testAccNotificationRestEndpointSyncDisabled(name string, outboundMemberType string, uri string, syncDisabled string) string {
	return fmt.Sprintf(`
resource "nios_notification_rest_endpoint" "test_sync_disabled" {
    name = %q
    outbound_member_type = %q
    uri = %q
    sync_disabled = %q
}
`, name, outboundMemberType, uri, syncDisabled)
}

func testAccNotificationRestEndpointTemplateInstance(name string, outboundMemberType string, uri string, templateInstance map[string]any) string {
	templateInstanceHCL := utils.ConvertMapToHCL(templateInstance)
	return fmt.Sprintf(`
resource "nios_notification_rest_endpoint" "test_template_instance" {
    name = %q
    outbound_member_type = %q
    uri = %q
    template_instance = %s
}
`, name, outboundMemberType, uri, templateInstanceHCL)
}

func testAccNotificationRestEndpointTimeout(name string, outboundMemberType string, uri string, timeout string) string {
	return fmt.Sprintf(`
resource "nios_notification_rest_endpoint" "test_timeout" {
    name = %q
    outbound_member_type = %q
    uri = %q
    timeout = %q
}
`, name, outboundMemberType, uri, timeout)
}

func testAccNotificationRestEndpointUri(name string, outboundMemberType string, uri string) string {
	return fmt.Sprintf(`
resource "nios_notification_rest_endpoint" "test_uri" {
    name = %q
    outbound_member_type = %q
    uri = %q
}
`, name, outboundMemberType, uri)
}

func testAccNotificationRestEndpointUsername(name string, outboundMemberType string, uri string, username string, password string) string {
	return fmt.Sprintf(`
resource "nios_notification_rest_endpoint" "test_username" {
    name = %q
    outbound_member_type = %q
    uri = %q
    username = %q
	password = %q
}
`, name, outboundMemberType, uri, username, password)
}

func testAccNotificationRestEndpointVendorIdentifier(name string, outboundMemberType string, uri string, vendorIdentifier string) string {
	return fmt.Sprintf(`
resource "nios_notification_rest_endpoint" "test_vendor_identifier" {
    name = %q
    outbound_member_type = %q
    uri = %q
    vendor_identifier = %q
}
`, name, outboundMemberType, uri, vendorIdentifier)
}

func testAccNotificationRestEndpointWapiUserName(name string, outboundMemberType string, uri string, wapiUserName string, wapiUserPassword string) string {
	return fmt.Sprintf(`
resource "nios_notification_rest_endpoint" "test_wapi_user_name" {
    name = %q
    outbound_member_type = %q
    uri = %q
    wapi_user_name = %q
	wapi_user_password = %q
}
`, name, outboundMemberType, uri, wapiUserName, wapiUserPassword)
}

func getTestDataPath() string {
	wd, err := os.Getwd()
	if err != nil {
		return "../../testdata/nios_notification_rest_endpoint"
	}
	return filepath.Join(wd, "../../testdata/nios_notification_rest_endpoint")
}
