package misc_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForPxgridEndpoint = "address,client_certificate_subject,client_certificate_valid_from,client_certificate_valid_to,comment,disable,extattrs,log_level,name,network_view,outbound_member_type,outbound_members,publish_settings,subscribe_settings,template_instance,timeout,vendor_identifier,wapi_user_name"

func TestAccPxgridEndpointResource_basic(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test"
	var v misc.PxgridEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPxgridEndpointResource_disappears(t *testing.T) {
	resourceName := "nios_misc_pxgrid_endpoint.test"
	var v misc.PxgridEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPxgridEndpointDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccPxgridEndpointBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					testAccCheckPxgridEndpointDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccPxgridEndpointResource_Ref(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_ref"
	var v misc.PxgridEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPxgridEndpointResource_Address(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_address"
	var v misc.PxgridEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointAddress("ADDRESS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "ADDRESS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointAddress("ADDRESS_UPDATE_REPLACE_ME"),
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
				Config: testAccPxgridEndpointClientCertificateToken("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_certificate_token", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointClientCertificateToken("CLIENT_CERTIFICATE_TOKEN_UPDATE_REPLACE_ME"),
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
				Config: testAccPxgridEndpointComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
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
				Config: testAccPxgridEndpointDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPxgridEndpointResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_extattrs"
	var v misc.PxgridEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
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
				Config: testAccPxgridEndpointLogLevel("LOG_LEVEL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "LOG_LEVEL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointLogLevel("LOG_LEVEL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "LOG_LEVEL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
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
				Config: testAccPxgridEndpointName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointName("NAME_UPDATE_REPLACE_ME"),
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
				Config: testAccPxgridEndpointNetworkView("NETWORK_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointNetworkView("NETWORK_VIEW_UPDATE_REPLACE_ME"),
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
				Config: testAccPxgridEndpointOutboundMemberType("OUTBOUND_MEMBER_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", "OUTBOUND_MEMBER_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointOutboundMemberType("OUTBOUND_MEMBER_TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", "OUTBOUND_MEMBER_TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccPxgridEndpointResource_OutboundMembers(t *testing.T) {
	var resourceName = "nios_misc_pxgrid_endpoint.test_outbound_members"
	var v misc.PxgridEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointOutboundMembers("OUTBOUND_MEMBERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_members", "OUTBOUND_MEMBERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointOutboundMembers("OUTBOUND_MEMBERS_UPDATE_REPLACE_ME"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointPublishSettings("PUBLISH_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "publish_settings", "PUBLISH_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointPublishSettings("PUBLISH_SETTINGS_UPDATE_REPLACE_ME"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointSubscribeSettings("SUBSCRIBE_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscribe_settings", "SUBSCRIBE_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointSubscribeSettings("SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccPxgridEndpointTemplateInstance("TEMPLATE_INSTANCE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_instance", "TEMPLATE_INSTANCE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointTemplateInstance("TEMPLATE_INSTANCE_UPDATE_REPLACE_ME"),
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
				Config: testAccPxgridEndpointTimeout("TIMEOUT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "TIMEOUT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointTimeout("TIMEOUT_UPDATE_REPLACE_ME"),
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
				Config: testAccPxgridEndpointVendorIdentifier("VENDOR_IDENTIFIER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_identifier", "VENDOR_IDENTIFIER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointVendorIdentifier("VENDOR_IDENTIFIER_UPDATE_REPLACE_ME"),
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
				Config: testAccPxgridEndpointWapiUserName("WAPI_USER_NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_name", "WAPI_USER_NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointWapiUserName("WAPI_USER_NAME_UPDATE_REPLACE_ME"),
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
				Config: testAccPxgridEndpointWapiUserPassword("WAPI_USER_PASSWORD_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPxgridEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_password", "WAPI_USER_PASSWORD_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccPxgridEndpointWapiUserPassword("WAPI_USER_PASSWORD_UPDATE_REPLACE_ME"),
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

func testAccPxgridEndpointBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return `
resource "nios_misc_pxgrid_endpoint" "test" {
}
`
}

func testAccPxgridEndpointRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccPxgridEndpointAddress(address string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_address" {
    address = %q
}
`, address)
}

func testAccPxgridEndpointClientCertificateToken(clientCertificateToken string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_client_certificate_token" {
    client_certificate_token = %q
}
`, clientCertificateToken)
}

func testAccPxgridEndpointComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccPxgridEndpointDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccPxgridEndpointExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccPxgridEndpointLogLevel(logLevel string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_log_level" {
    log_level = %q
}
`, logLevel)
}

func testAccPxgridEndpointName(name string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_name" {
    name = %q
}
`, name)
}

func testAccPxgridEndpointNetworkView(networkView string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_network_view" {
    network_view = %q
}
`, networkView)
}

func testAccPxgridEndpointOutboundMemberType(outboundMemberType string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_outbound_member_type" {
    outbound_member_type = %q
}
`, outboundMemberType)
}

func testAccPxgridEndpointOutboundMembers(outboundMembers string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_outbound_members" {
    outbound_members = %q
}
`, outboundMembers)
}

func testAccPxgridEndpointPublishSettings(publishSettings string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_publish_settings" {
    publish_settings = %q
}
`, publishSettings)
}

func testAccPxgridEndpointSubscribeSettings(subscribeSettings string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_subscribe_settings" {
    subscribe_settings = %q
}
`, subscribeSettings)
}

func testAccPxgridEndpointTemplateInstance(templateInstance string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_template_instance" {
    template_instance = %q
}
`, templateInstance)
}

func testAccPxgridEndpointTimeout(timeout string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_timeout" {
    timeout = %q
}
`, timeout)
}

func testAccPxgridEndpointVendorIdentifier(vendorIdentifier string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_vendor_identifier" {
    vendor_identifier = %q
}
`, vendorIdentifier)
}

func testAccPxgridEndpointWapiUserName(wapiUserName string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_wapi_user_name" {
    wapi_user_name = %q
}
`, wapiUserName)
}

func testAccPxgridEndpointWapiUserPassword(wapiUserPassword string) string {
	return fmt.Sprintf(`
resource "nios_misc_pxgrid_endpoint" "test_wapi_user_password" {
    wapi_user_password = %q
}
`, wapiUserPassword)
}
