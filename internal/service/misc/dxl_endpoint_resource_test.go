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

var readableAttributesForDxlEndpoint = "brokers,client_certificate_subject,client_certificate_valid_from,client_certificate_valid_to,comment,disable,extattrs,log_level,name,outbound_member_type,outbound_members,template_instance,timeout,topics,vendor_identifier,wapi_user_name"

func TestAccDxlEndpointResource_basic(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test"
	var v misc.DxlEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_disappears(t *testing.T) {
	resourceName := "nios_misc_dxl_endpoint.test"
	var v misc.DxlEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDxlEndpointDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDxlEndpointBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					testAccCheckDxlEndpointDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDxlEndpointResource_Ref(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_ref"
	var v misc.DxlEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_Brokers(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_brokers"
	var v misc.DxlEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointBrokers("BROKERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "brokers", "BROKERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointBrokers("BROKERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "brokers", "BROKERS_UPDATE_REPLACE_ME"),
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
				Config: testAccDxlEndpointBrokersImportToken("BROKERS_IMPORT_TOKEN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "brokers_import_token", "BROKERS_IMPORT_TOKEN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointBrokersImportToken("BROKERS_IMPORT_TOKEN_UPDATE_REPLACE_ME"),
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
				Config: testAccDxlEndpointClientCertificateToken("CLIENT_CERTIFICATE_TOKEN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_certificate_token", "CLIENT_CERTIFICATE_TOKEN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointClientCertificateToken("CLIENT_CERTIFICATE_TOKEN_UPDATE_REPLACE_ME"),
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
				Config: testAccDxlEndpointComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
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
				Config: testAccDxlEndpointDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_extattrs"
	var v misc.DxlEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
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
				Config: testAccDxlEndpointLogLevel("LOG_LEVEL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "LOG_LEVEL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointLogLevel("LOG_LEVEL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "LOG_LEVEL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
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
				Config: testAccDxlEndpointName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointName("NAME_UPDATE_REPLACE_ME"),
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
				Config: testAccDxlEndpointOutboundMemberType("OUTBOUND_MEMBER_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", "OUTBOUND_MEMBER_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointOutboundMemberType("OUTBOUND_MEMBER_TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", "OUTBOUND_MEMBER_TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_OutboundMembers(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_outbound_members"
	var v misc.DxlEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointOutboundMembers("OUTBOUND_MEMBERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_members", "OUTBOUND_MEMBERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointOutboundMembers("OUTBOUND_MEMBERS_UPDATE_REPLACE_ME"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointTemplateInstance("TEMPLATE_INSTANCE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_instance", "TEMPLATE_INSTANCE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointTemplateInstance("TEMPLATE_INSTANCE_UPDATE_REPLACE_ME"),
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
				Config: testAccDxlEndpointTimeout("TIMEOUT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "TIMEOUT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointTimeout("TIMEOUT_UPDATE_REPLACE_ME"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointTopics("TOPICS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "topics", "TOPICS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointTopics("TOPICS_UPDATE_REPLACE_ME"),
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
				Config: testAccDxlEndpointVendorIdentifier("VENDOR_IDENTIFIER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_identifier", "VENDOR_IDENTIFIER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointVendorIdentifier("VENDOR_IDENTIFIER_UPDATE_REPLACE_ME"),
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
				Config: testAccDxlEndpointWapiUserName("WAPI_USER_NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_name", "WAPI_USER_NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointWapiUserName("WAPI_USER_NAME_UPDATE_REPLACE_ME"),
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
				Config: testAccDxlEndpointWapiUserPassword("WAPI_USER_PASSWORD_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_password", "WAPI_USER_PASSWORD_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointWapiUserPassword("WAPI_USER_PASSWORD_UPDATE_REPLACE_ME"),
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

func testAccDxlEndpointBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test" {
}
`)
}

func testAccDxlEndpointRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccDxlEndpointBrokers(brokers string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_brokers" {
    brokers = %q
}
`, brokers)
}

func testAccDxlEndpointBrokersImportToken(brokersImportToken string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_brokers_import_token" {
    brokers_import_token = %q
}
`, brokersImportToken)
}

func testAccDxlEndpointClientCertificateToken(clientCertificateToken string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_client_certificate_token" {
    client_certificate_token = %q
}
`, clientCertificateToken)
}

func testAccDxlEndpointComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccDxlEndpointDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccDxlEndpointExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccDxlEndpointLogLevel(logLevel string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_log_level" {
    log_level = %q
}
`, logLevel)
}

func testAccDxlEndpointName(name string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_name" {
    name = %q
}
`, name)
}

func testAccDxlEndpointOutboundMemberType(outboundMemberType string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_outbound_member_type" {
    outbound_member_type = %q
}
`, outboundMemberType)
}

func testAccDxlEndpointOutboundMembers(outboundMembers string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_outbound_members" {
    outbound_members = %q
}
`, outboundMembers)
}

func testAccDxlEndpointTemplateInstance(templateInstance string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_template_instance" {
    template_instance = %q
}
`, templateInstance)
}

func testAccDxlEndpointTimeout(timeout string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_timeout" {
    timeout = %q
}
`, timeout)
}

func testAccDxlEndpointTopics(topics string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_topics" {
    topics = %q
}
`, topics)
}

func testAccDxlEndpointVendorIdentifier(vendorIdentifier string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_vendor_identifier" {
    vendor_identifier = %q
}
`, vendorIdentifier)
}

func testAccDxlEndpointWapiUserName(wapiUserName string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_wapi_user_name" {
    wapi_user_name = %q
}
`, wapiUserName)
}

func testAccDxlEndpointWapiUserPassword(wapiUserPassword string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_wapi_user_password" {
    wapi_user_password = %q
}
`, wapiUserPassword)
}
