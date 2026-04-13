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

// NOTE:
// - The following Grid Masters (GMs) must be present:
// - infoblox.grid_master_candidate1
// - infoblox.grid_master_candidate2
// - The following templates must be added to the Grid:
// - Version5_Syslog_Session_Template
// - Version5_Syslog_Session_Template1
// - Special characters and whitespaces are not accepted in the "name" field.
// - The "address" field must be a valid IPv4 address.

var readableAttributesForSyslogEndpoint = "extattrs,log_level,name,outbound_member_type,outbound_members,syslog_servers,template_instance,timeout,vendor_identifier,wapi_user_name"

func TestAccSyslogEndpointResource_basic(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test"
	var v misc.SyslogEndpoint
	name := "syslogserverbasic"
	outboundMemberType := "GM"
	syslogServer := "10.120.21.82"
	connectionType := "stcp"
	format := "formatted"
	certificateFilePath := "client.crt"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointBasicConfig(name, outboundMemberType, syslogServer, connectionType, format, certificateFilePath),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", outboundMemberType),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.address", syslogServer),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.format", format),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.connection_type", connectionType),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "log_level", "WARNING"),
					resource.TestCheckResourceAttr(resourceName, "timeout", "30"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_disappears(t *testing.T) {
	resourceName := "nios_misc_syslog_endpoint.test"
	var v misc.SyslogEndpoint
	name := "syslogserverdisappears"
	outboundMemberType := "GM"
	syslogServer := "10.1.1.1"
	connectionType := "udp"
	format := "formatted"
	certificateFilePath := "client.crt"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSyslogEndpointDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSyslogEndpointBasicConfig(name, outboundMemberType, syslogServer, connectionType, format, certificateFilePath),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					testAccCheckSyslogEndpointDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSyslogEndpointResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_extattrs"
	var v misc.SyslogEndpoint
	site := acctest.RandomNameWithPrefix("site")
	name := "syslogserverextattrs"
	outboundMemberType := "GM"
	syslogServer := "10.1.1.1"
	connectionType := "udp"
	format := "formatted"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointExtAttrs(name, outboundMemberType, syslogServer, connectionType, format, map[string]string{"Site": site}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", site),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointExtAttrs(name, outboundMemberType, syslogServer, connectionType, format, map[string]string{"Site": site}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", site),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_LogLevel(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_log_level"
	var v misc.SyslogEndpoint
	name := "syslogserverloglevel"
	outboundMemberType := "GM"
	syslogServer := "10.120.21.82"
	connectionType := "udp"
	format := "formatted"
	logLevel := "DEBUG"
	updatedLogLevel := "INFO"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointLogLevel(name, outboundMemberType, syslogServer, connectionType, format, logLevel),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", logLevel),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointLogLevel(name, outboundMemberType, syslogServer, connectionType, format, updatedLogLevel),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", updatedLogLevel),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_Name(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_name"
	var v misc.SyslogEndpoint
	name := "syslogservername"
	outboundMemberType := "GM"
	syslogServer := "10.1.1.1"
	connectionType := "udp"
	format := "formatted"
	updatedName := "syslogservernameupdated"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointName(name, outboundMemberType, syslogServer, connectionType, format),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointName(updatedName, outboundMemberType, syslogServer, connectionType, format),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_OutboundMemberType(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_outbound_member_type"
	var v misc.SyslogEndpoint
	name := "syslogserveroutboundmembertype"
	outboundMemberType := "GM"
	syslogServer := "10.1.1.1"
	connectionType := "udp"
	format := "formatted"
	updatedOutboundMember := "MEMBER"
	outboundMember := "infoblox.member"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointOutboundMemberType(name, outboundMemberType, syslogServer, connectionType, format),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", outboundMemberType),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointOutboundMemberTypeUpdated(name, updatedOutboundMember, syslogServer, connectionType, format, outboundMember),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", updatedOutboundMember),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_OutboundMembers(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_outbound_members"
	var v misc.SyslogEndpoint
	name := "syslogserveroutboundmembers"
	outboundMemberType := "MEMBER"
	syslogServer := "10.1.1.1"
	connectionType := "udp"
	format := "formatted"
	outboundMemberTypeUpdated := "MEMBER"
	outboundMember := "infoblox.member"
	outboundMemberUpdated := "infoblox.grid_master_candidate1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointOutboundMembers(name, outboundMemberType, syslogServer, connectionType, format, outboundMember),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_members.0", outboundMember),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointOutboundMembers(name, outboundMemberTypeUpdated, syslogServer, connectionType, format, outboundMemberUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_members.0", outboundMemberUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_SyslogServers(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_syslog_servers"
	var v misc.SyslogEndpoint
	name := "syslogservers"
	outboundMemberType := "GM"
	syslogServer := "1.1.1.1"
	connectionType := "udp"
	format := "formatted"
	updatedSyslogServer := "10.1.1.12"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointSyslogServers(name, outboundMemberType, syslogServer, connectionType, format),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.address", syslogServer),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointSyslogServers(name, outboundMemberType, updatedSyslogServer, connectionType, format),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.address", updatedSyslogServer),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_TemplateInstance(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_template_instance"
	var v misc.SyslogEndpoint
	templateInstance := map[string]any{
		"template": "Version5_Syslog_Session_Template",
	}
	name := "syslogservertemplateinstance"
	outboundMemberType := "GM"
	syslogServer := "10.1.1.1"
	connectionType := "udp"
	format := "formatted"
	updatedTemplateInstance := map[string]any{
		"template": "Version5_Syslog_Session_Template1",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointTemplateInstance(name, outboundMemberType, syslogServer, connectionType, format, templateInstance),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_instance.template", templateInstance["template"].(string)),
				),
			},
			// Update and Read - Immutable Attribute template_instance
			{
				Config: testAccSyslogEndpointTemplateInstance(name, outboundMemberType, syslogServer, connectionType, format, updatedTemplateInstance),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_instance.template", updatedTemplateInstance["template"].(string)),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_Timeout(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_timeout"
	var v misc.SyslogEndpoint
	name := "syslogservertimeout"
	outboundMemberType := "GM"
	syslogServer := "10.1.1.1"
	connectionType := "udp"
	format := "formatted"
	timeout := "45"
	timeoutUpdated := "60"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointTimeout(name, outboundMemberType, syslogServer, connectionType, format, timeout),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", timeout),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointTimeout(name, outboundMemberType, syslogServer, connectionType, format, timeoutUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", timeoutUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_VendorIdentifier(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_vendor_identifier"
	var v misc.SyslogEndpoint
	name := "syslogservervendoridentifier"
	outboundMemberType := "GM"
	syslogServer := "10.1.1.1"
	connectionType := "udp"
	format := "formatted"
	vendorIdentifier := "WAPI"
	updatedVendorIdentifier := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointVendorIdentifier(name, outboundMemberType, syslogServer, connectionType, format, vendorIdentifier),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_identifier", vendorIdentifier),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointVendorIdentifier(name, outboundMemberType, syslogServer, connectionType, format, updatedVendorIdentifier),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_identifier", updatedVendorIdentifier),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_WapiUserName(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_wapi_user_name"
	var v misc.SyslogEndpoint
	name := "syslogserverwapiusername"
	outboundMemberType := "GM"
	syslogServer := "10.1.1.2"
	connectionType := "udp"
	format := "formatted"
	wapiUserName := "admin"
	updatedWapiUserName := "admin1"
	wapiPassword := "Example-Admin123!"
	updatedWapiPassword := "Example-Admin1234!"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointWapiUserName(name, outboundMemberType, syslogServer, connectionType, format, wapiUserName, wapiPassword),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_name", wapiUserName),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointWapiUserName(name, outboundMemberType, syslogServer, connectionType, format, updatedWapiUserName, updatedWapiPassword),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_name", updatedWapiUserName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSyslogEndpointResource_WapiUserPassword(t *testing.T) {
	var resourceName = "nios_misc_syslog_endpoint.test_wapi_user_password"
	var v misc.SyslogEndpoint
	name := "syslogserverwapiuserpassword"
	outboundMemberType := "GM"
	syslogServer := "10.1.1.2"
	connectionType := "udp"
	format := "formatted"
	wapiUserName := "admin"
	wapiUserPassword := "WAPI_USER_PASSWORD_REPLACE_ME"
	updatedWapiUserPassword := "WAPI_USER_PASSWORD_UPDATE_REPLACE_ME"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSyslogEndpointWapiUserPassword(name, outboundMemberType, syslogServer, connectionType, format, wapiUserName, wapiUserPassword),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSyslogEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_password", "WAPI_USER_PASSWORD_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSyslogEndpointWapiUserPassword(name, outboundMemberType, syslogServer, connectionType, format, wapiUserName, updatedWapiUserPassword),
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

func testAccSyslogEndpointBasicConfig(name string, outboundMemberType string, syslogServer string, connectionType string, format string, certificateFilePath string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test" {
    name = %q
    outbound_member_type = %q
    syslog_servers = [
        {
            address = %q
			connection_type = %q
			format = %q
			certificate_file_path = %q
        }
    ]
}
`, name, outboundMemberType, syslogServer, connectionType, format, certificateFilePath)
}

func testAccSyslogEndpointExtAttrs(name string, outboundMemberType string, syslogServer string, connectionType string, format string, extAttrsMap map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrsMap {
		extattrsStr += fmt.Sprintf("  %s = %q\n", k, v)
	}
	extattrsStr += "}"
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_extattrs" {
    name = %q
    outbound_member_type = %q
    syslog_servers = [
        {
            address = %q
			connection_type = %q
			format = %q
        }
    ]
    extattrs = %s
}
`, name, outboundMemberType, syslogServer, connectionType, format, extattrsStr)
}

func testAccSyslogEndpointLogLevel(name string, outboundMemberType string, syslogServer string, connectionType string, format string, logLevel string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_log_level" {
    name = %q
    outbound_member_type = %q
    syslog_servers = [
        {
            address = %q
			connection_type = %q
			format = %q
        }
    ]
    log_level = %q
}
`, name, outboundMemberType, syslogServer, connectionType, format, logLevel)
}

func testAccSyslogEndpointName(name string, outboundMemberType string, syslogServer string, connectionType string, format string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_name" {
    name = %q
    outbound_member_type = %q
    syslog_servers = [
        {
            address = %q
			connection_type = %q
			format = %q
        }
    ]	
}
`, name, outboundMemberType, syslogServer, connectionType, format)
}

func testAccSyslogEndpointOutboundMemberType(name string, outboundMemberType string, syslogServer string, connectionType string, format string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_outbound_member_type" {
    name = %q
    outbound_member_type = %q
    syslog_servers = [
        {
            address = %q
			connection_type = %q
			format = %q
        }
    ]
}
`, name, outboundMemberType, syslogServer, connectionType, format)
}

func testAccSyslogEndpointOutboundMemberTypeUpdated(name string, outboundMemberType string, syslogServer string, connectionType string, format string, outboundMember string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_outbound_member_type" {
    name = %q
    outbound_member_type = %q
    syslog_servers = [
        {
            address = %q
			connection_type = %q
			format = %q
        }
    ]
	outbound_members = [%q]
}
`, name, outboundMemberType, syslogServer, connectionType, format, outboundMember)
}

func testAccSyslogEndpointOutboundMembers(name string, outboundMemberType string, syslogServer string, connectionType string, format string, outboundMember string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_outbound_members" {
    name = %q
    outbound_member_type = %q
    syslog_servers = [
        {
            address = %q
			connection_type = %q
			format = %q
        }
    ]
	outbound_members = [%q]
}
`, name, outboundMemberType, syslogServer, connectionType, format, outboundMember)
}

func testAccSyslogEndpointSyslogServers(name string, outboundMemberType string, syslogServer string, connectionType string, format string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_syslog_servers" {
    name = %q
    outbound_member_type = %q
    syslog_servers = [
        {
            address = %q
			connection_type = %q
			format = %q
        }
    ]
}
`, name, outboundMemberType, syslogServer, connectionType, format)

}

func testAccSyslogEndpointTemplateInstance(name string, outboundMemberType string, syslogServer string, connectionType string, format string, templateInstance map[string]any) string {
	templateInstanceHCL := utils.ConvertMapToHCL(templateInstance)
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_template_instance" {
    name = %q
    outbound_member_type = %q
    syslog_servers = [
        {
            address = %q
			connection_type = %q
			format = %q
        }
    ]
    template_instance = %s
}
`, name, outboundMemberType, syslogServer, connectionType, format, templateInstanceHCL)
}

func testAccSyslogEndpointTimeout(name string, outboundMemberType string, syslogServer string, connectionType string, format string, timeout string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_timeout" {
    name = %q
    outbound_member_type = %q
    syslog_servers = [
        {
            address = %q
			connection_type = %q
			format = %q
        }
    ]
    timeout = %q
}
`, name, outboundMemberType, syslogServer, connectionType, format, timeout)
}

func testAccSyslogEndpointVendorIdentifier(name string, outboundMemberType string, syslogServer string, connectionType string, format string, vendorIdentifier string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_vendor_identifier" {
    name = %q
    outbound_member_type = %q
    syslog_servers = [
        {
            address = %q
			connection_type = %q
			format = %q
        }
    ]
    vendor_identifier = %q
}
`, name, outboundMemberType, syslogServer, connectionType, format, vendorIdentifier)
}

func testAccSyslogEndpointWapiUserName(name string, outboundMemberType string, syslogServer string, connectionType string, format string, wapiUserName string, wapiUserPassword string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_wapi_user_name" {
    name = %q
    outbound_member_type = %q
    syslog_servers = [
        {
            address = %q
			connection_type = %q
			format = %q
        }
    ]
    wapi_user_name = %q
    wapi_user_password = %q
}
`, name, outboundMemberType, syslogServer, connectionType, format, wapiUserName, wapiUserPassword)
}

func testAccSyslogEndpointWapiUserPassword(name string, outboundMemberType string, syslogServer string, connectionType string, format string, wapiUserName string, wapiUserPassword string) string {
	return fmt.Sprintf(`
resource "nios_misc_syslog_endpoint" "test_wapi_user_password" {
    name = %q
    outbound_member_type = %q
    syslog_servers = [
        {
            address = %q
			connection_type = %q
			format = %q
        }
    ]
    wapi_user_name = %q
    wapi_user_password = %q
}
`, name, outboundMemberType, syslogServer, connectionType, format, wapiUserName, wapiUserPassword)
}
