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

var readableAttributesForSyslogEndpoint = "extattrs,log_level,name,outbound_member_type,outbound_members,syslog_servers,template_instance,timeout,vendor_identifier,wapi_user_name"

func TestAccSyslogEndpointResource_basic(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test"
	var v misc.SyslogEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_disappears(t *testing.T) {
	resourceName := "nios_misc_syslog_endpoint.test"
	var v misc.SyslogEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSyslogEndpointDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSyslogEndpointBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					testAccCheckSyslogEndpointDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSyslogEndpointResource_Ref(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_ref"
	var v misc.SyslogEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_extattrs"
	var v misc.SyslogEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_LogLevel(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_log_level"
	var v misc.SyslogEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointLogLevel("LOG_LEVEL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "LOG_LEVEL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointLogLevel("LOG_LEVEL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "LOG_LEVEL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_Name(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_name"
	var v misc.SyslogEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_OutboundMemberType(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_outbound_member_type"
	var v misc.SyslogEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointOutboundMemberType("OUTBOUND_MEMBER_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", "OUTBOUND_MEMBER_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointOutboundMemberType("OUTBOUND_MEMBER_TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", "OUTBOUND_MEMBER_TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_OutboundMembers(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_outbound_members"
	var v misc.SyslogEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointOutboundMembers("OUTBOUND_MEMBERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_members", "OUTBOUND_MEMBERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointOutboundMembers("OUTBOUND_MEMBERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_members", "OUTBOUND_MEMBERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_SyslogServers(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_syslog_servers"
	var v misc.SyslogEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointSyslogServers("SYSLOG_SERVERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers", "SYSLOG_SERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointSyslogServers("SYSLOG_SERVERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers", "SYSLOG_SERVERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_TemplateInstance(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_template_instance"
	var v misc.SyslogEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointTemplateInstance("TEMPLATE_INSTANCE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_instance", "TEMPLATE_INSTANCE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointTemplateInstance("TEMPLATE_INSTANCE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_instance", "TEMPLATE_INSTANCE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_Timeout(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_timeout"
	var v misc.SyslogEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointTimeout("TIMEOUT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "TIMEOUT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointTimeout("TIMEOUT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "TIMEOUT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_VendorIdentifier(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_vendor_identifier"
	var v misc.SyslogEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointVendorIdentifier("VENDOR_IDENTIFIER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_identifier", "VENDOR_IDENTIFIER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointVendorIdentifier("VENDOR_IDENTIFIER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_identifier", "VENDOR_IDENTIFIER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_WapiUserName(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_wapi_user_name"
	var v misc.SyslogEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointWapiUserName("WAPI_USER_NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_name", "WAPI_USER_NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointWapiUserName("WAPI_USER_NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_name", "WAPI_USER_NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_WapiUserPassword(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_wapi_user_password"
	var v misc.SyslogEndpoint

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointWapiUserPassword("WAPI_USER_PASSWORD_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_password", "WAPI_USER_PASSWORD_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointWapiUserPassword("WAPI_USER_PASSWORD_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_password", "WAPI_USER_PASSWORD_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckSyslogEndpointExists(ctx context.Context, resourceName string, v *misc.SyslogEndpoint) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.MiscAPI.
			SyslogEndpointAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForSyslogEndpoint).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetSyslogEndpointResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetSyslogEndpointResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckSyslogEndpointDestroy(ctx context.Context, v *misc.SyslogEndpoint) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.MiscAPI.
			SyslogEndpointAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForSyslogEndpoint).
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

func testAccCheckSyslogEndpointDisappears(ctx context.Context, v *misc.SyslogEndpoint) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.MiscAPI.
			SyslogEndpointAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccSyslogEndpointBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test" {
}
`)
}

func testAccSyslogEndpointRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccSyslogEndpointExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccSyslogEndpointLogLevel(logLevel string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_log_level" {
    log_level = %q
}
`, logLevel)
}

func testAccSyslogEndpointName(name string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_name" {
    name = %q
}
`, name)
}

func testAccSyslogEndpointOutboundMemberType(outboundMemberType string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_outbound_member_type" {
    outbound_member_type = %q
}
`, outboundMemberType)
}

func testAccSyslogEndpointOutboundMembers(outboundMembers string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_outbound_members" {
    outbound_members = %q
}
`, outboundMembers)
}

func testAccSyslogEndpointSyslogServers(syslogServers string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_syslog_servers" {
    syslog_servers = %q
}
`, syslogServers)
}

func testAccSyslogEndpointTemplateInstance(templateInstance string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_template_instance" {
    template_instance = %q
}
`, templateInstance)
}

func testAccSyslogEndpointTimeout(timeout string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_timeout" {
    timeout = %q
}
`, timeout)
}

func testAccSyslogEndpointVendorIdentifier(vendorIdentifier string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_vendor_identifier" {
    vendor_identifier = %q
}
`, vendorIdentifier)
}

func testAccSyslogEndpointWapiUserName(wapiUserName string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_wapi_user_name" {
    wapi_user_name = %q
}
`, wapiUserName)
}

func testAccSyslogEndpointWapiUserPassword(wapiUserPassword string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_wapi_user_password" {
    wapi_user_password = %q
}
`, wapiUserPassword)
}
