package dhcp_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dhcp"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/acctest"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/utils"
)

// TODO : Add readable attributes for the resource
var readableAttributesForFixedaddress = ""

func TestAccFixedaddressResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_fixedaddress.test"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFixedaddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFixedaddressBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					testAccCheckFixedaddressDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccFixedaddressResource_Ref(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test__ref"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "_ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "_ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_AgentCircuitId(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_agent_circuit_id"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressAgentCircuitId("AGENT_CIRCUIT_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "agent_circuit_id", "AGENT_CIRCUIT_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressAgentCircuitId("AGENT_CIRCUIT_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "agent_circuit_id", "AGENT_CIRCUIT_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_AgentRemoteId(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_agent_remote_id"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressAgentRemoteId("AGENT_REMOTE_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "agent_remote_id", "AGENT_REMOTE_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressAgentRemoteId("AGENT_REMOTE_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "agent_remote_id", "AGENT_REMOTE_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_AllowTelnet(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_allow_telnet"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressAllowTelnet("ALLOW_TELNET_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_telnet", "ALLOW_TELNET_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressAllowTelnet("ALLOW_TELNET_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_telnet", "ALLOW_TELNET_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_AlwaysUpdateDns(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_always_update_dns"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressAlwaysUpdateDns("ALWAYS_UPDATE_DNS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "always_update_dns", "ALWAYS_UPDATE_DNS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressAlwaysUpdateDns("ALWAYS_UPDATE_DNS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "always_update_dns", "ALWAYS_UPDATE_DNS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Bootfile(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_bootfile"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressBootfile("BOOTFILE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "BOOTFILE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressBootfile("BOOTFILE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "BOOTFILE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Bootserver(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_bootserver"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressBootserver("BOOTSERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "BOOTSERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressBootserver("BOOTSERVER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "BOOTSERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_CliCredentials(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_cli_credentials"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressCliCredentials("CLI_CREDENTIALS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials", "CLI_CREDENTIALS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressCliCredentials("CLI_CREDENTIALS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials", "CLI_CREDENTIALS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_ClientIdentifierPrependZero(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_client_identifier_prepend_zero"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressClientIdentifierPrependZero("CLIENT_IDENTIFIER_PREPEND_ZERO_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_identifier_prepend_zero", "CLIENT_IDENTIFIER_PREPEND_ZERO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressClientIdentifierPrependZero("CLIENT_IDENTIFIER_PREPEND_ZERO_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_identifier_prepend_zero", "CLIENT_IDENTIFIER_PREPEND_ZERO_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_CloudInfo(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_cloud_info"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressCloudInfo("CLOUD_INFO_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_info", "CLOUD_INFO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressCloudInfo("CLOUD_INFO_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_info", "CLOUD_INFO_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_comment"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_DdnsDomainname(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_ddns_domainname"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressDdnsDomainname("DDNS_DOMAINNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "DDNS_DOMAINNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressDdnsDomainname("DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_DdnsHostname(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_ddns_hostname"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressDdnsHostname("DDNS_HOSTNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_hostname", "DDNS_HOSTNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressDdnsHostname("DDNS_HOSTNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_hostname", "DDNS_HOSTNAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_DenyBootp(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_deny_bootp"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressDenyBootp("DENY_BOOTP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "DENY_BOOTP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressDenyBootp("DENY_BOOTP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "DENY_BOOTP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_DeviceDescription(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_device_description"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressDeviceDescription("DEVICE_DESCRIPTION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_description", "DEVICE_DESCRIPTION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressDeviceDescription("DEVICE_DESCRIPTION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_description", "DEVICE_DESCRIPTION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_DeviceLocation(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_device_location"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressDeviceLocation("DEVICE_LOCATION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_location", "DEVICE_LOCATION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressDeviceLocation("DEVICE_LOCATION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_location", "DEVICE_LOCATION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_DeviceType(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_device_type"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressDeviceType("DEVICE_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_type", "DEVICE_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressDeviceType("DEVICE_TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_type", "DEVICE_TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_DeviceVendor(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_device_vendor"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressDeviceVendor("DEVICE_VENDOR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_vendor", "DEVICE_VENDOR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressDeviceVendor("DEVICE_VENDOR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_vendor", "DEVICE_VENDOR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_DhcpClientIdentifier(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_dhcp_client_identifier"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressDhcpClientIdentifier("DHCP_CLIENT_IDENTIFIER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_client_identifier", "DHCP_CLIENT_IDENTIFIER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressDhcpClientIdentifier("DHCP_CLIENT_IDENTIFIER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_client_identifier", "DHCP_CLIENT_IDENTIFIER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Disable(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_disable"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_DisableDiscovery(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_disable_discovery"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressDisableDiscovery("DISABLE_DISCOVERY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_discovery", "DISABLE_DISCOVERY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressDisableDiscovery("DISABLE_DISCOVERY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_discovery", "DISABLE_DISCOVERY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_DiscoveredData(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_discovered_data"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressDiscoveredData("DISCOVERED_DATA_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovered_data", "DISCOVERED_DATA_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressDiscoveredData("DISCOVERED_DATA_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovered_data", "DISCOVERED_DATA_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_EnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_enable_ddns"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressEnableDdns("ENABLE_DDNS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "ENABLE_DDNS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressEnableDdns("ENABLE_DDNS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "ENABLE_DDNS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_EnableImmediateDiscovery(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_enable_immediate_discovery"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressEnableImmediateDiscovery("ENABLE_IMMEDIATE_DISCOVERY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_immediate_discovery", "ENABLE_IMMEDIATE_DISCOVERY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressEnableImmediateDiscovery("ENABLE_IMMEDIATE_DISCOVERY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_immediate_discovery", "ENABLE_IMMEDIATE_DISCOVERY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_EnablePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_enable_pxe_lease_time"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressEnablePxeLeaseTime("ENABLE_PXE_LEASE_TIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "ENABLE_PXE_LEASE_TIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressEnablePxeLeaseTime("ENABLE_PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "ENABLE_PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_extattrs"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_IgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_ignore_dhcp_option_list_request"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressIgnoreDhcpOptionListRequest("IGNORE_DHCP_OPTION_LIST_REQUEST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "IGNORE_DHCP_OPTION_LIST_REQUEST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressIgnoreDhcpOptionListRequest("IGNORE_DHCP_OPTION_LIST_REQUEST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "IGNORE_DHCP_OPTION_LIST_REQUEST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Ipv4addr(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_ipv4addr"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressIpv4addr("IPV4ADDR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "IPV4ADDR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressIpv4addr("IPV4ADDR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "IPV4ADDR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_FuncCall(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_func_call"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressFuncCall("FUNC_CALL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "func_call", "FUNC_CALL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressFuncCall("FUNC_CALL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "func_call", "FUNC_CALL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_LogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_logic_filter_rules"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressLogicFilterRules("LOGIC_FILTER_RULES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressLogicFilterRules("LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Mac(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_mac"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressMac("MAC_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mac", "MAC_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressMac("MAC_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mac", "MAC_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_MatchClient(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_match_client"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressMatchClient("MATCH_CLIENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_client", "MATCH_CLIENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressMatchClient("MATCH_CLIENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_client", "MATCH_CLIENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_MsAdUserData(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_ms_ad_user_data"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressMsAdUserData("MS_AD_USER_DATA_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_user_data", "MS_AD_USER_DATA_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressMsAdUserData("MS_AD_USER_DATA_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_user_data", "MS_AD_USER_DATA_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_MsOptions(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_ms_options"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressMsOptions("MS_OPTIONS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_options", "MS_OPTIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressMsOptions("MS_OPTIONS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_options", "MS_OPTIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_MsServer(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_ms_server"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressMsServer("MS_SERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_server", "MS_SERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressMsServer("MS_SERVER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_server", "MS_SERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_name"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Network(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_network"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressNetwork("NETWORK_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", "NETWORK_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressNetwork("NETWORK_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", "NETWORK_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_NetworkView(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_network_view"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressNetworkView("NETWORK_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressNetworkView("NETWORK_VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Nextserver(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_nextserver"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressNextserver("NEXTSERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", "NEXTSERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressNextserver("NEXTSERVER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", "NEXTSERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Options(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_options"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressOptions("OPTIONS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressOptions("OPTIONS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_PxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_pxe_lease_time"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressPxeLeaseTime("PXE_LEASE_TIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "PXE_LEASE_TIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressPxeLeaseTime("PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_ReservedInterface(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_reserved_interface"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressReservedInterface("RESERVED_INTERFACE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "reserved_interface", "RESERVED_INTERFACE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressReservedInterface("RESERVED_INTERFACE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "reserved_interface", "RESERVED_INTERFACE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_RestartIfNeeded(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_restart_if_needed"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressRestartIfNeeded("RESTART_IF_NEEDED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "restart_if_needed", "RESTART_IF_NEEDED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressRestartIfNeeded("RESTART_IF_NEEDED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "restart_if_needed", "RESTART_IF_NEEDED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Snmp3Credential(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_snmp3_credential"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressSnmp3Credential("SNMP3_CREDENTIAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential", "SNMP3_CREDENTIAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressSnmp3Credential("SNMP3_CREDENTIAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential", "SNMP3_CREDENTIAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_SnmpCredential(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_snmp_credential"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressSnmpCredential("SNMP_CREDENTIAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "snmp_credential", "SNMP_CREDENTIAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressSnmpCredential("SNMP_CREDENTIAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "snmp_credential", "SNMP_CREDENTIAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Template(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_template"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressTemplate("TEMPLATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template", "TEMPLATE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressTemplate("TEMPLATE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template", "TEMPLATE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseBootfile(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_use_bootfile"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseBootfile("USE_BOOTFILE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "USE_BOOTFILE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseBootfile("USE_BOOTFILE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "USE_BOOTFILE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseBootserver(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_use_bootserver"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseBootserver("USE_BOOTSERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "USE_BOOTSERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseBootserver("USE_BOOTSERVER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "USE_BOOTSERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseCliCredentials(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_use_cli_credentials"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseCliCredentials("USE_CLI_CREDENTIALS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_cli_credentials", "USE_CLI_CREDENTIALS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseCliCredentials("USE_CLI_CREDENTIALS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_cli_credentials", "USE_CLI_CREDENTIALS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseDdnsDomainname(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_use_ddns_domainname"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseDdnsDomainname("USE_DDNS_DOMAINNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "USE_DDNS_DOMAINNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseDdnsDomainname("USE_DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "USE_DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseDenyBootp(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_use_deny_bootp"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseDenyBootp("USE_DENY_BOOTP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "USE_DENY_BOOTP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseDenyBootp("USE_DENY_BOOTP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "USE_DENY_BOOTP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseEnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_use_enable_ddns"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseEnableDdns("USE_ENABLE_DDNS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "USE_ENABLE_DDNS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseEnableDdns("USE_ENABLE_DDNS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "USE_ENABLE_DDNS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseIgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_use_ignore_dhcp_option_list_request"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseIgnoreDhcpOptionListRequest("USE_IGNORE_DHCP_OPTION_LIST_REQUEST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "USE_IGNORE_DHCP_OPTION_LIST_REQUEST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseIgnoreDhcpOptionListRequest("USE_IGNORE_DHCP_OPTION_LIST_REQUEST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "USE_IGNORE_DHCP_OPTION_LIST_REQUEST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseLogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_use_logic_filter_rules"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseLogicFilterRules("USE_LOGIC_FILTER_RULES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "USE_LOGIC_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseLogicFilterRules("USE_LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "USE_LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseMsOptions(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_use_ms_options"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseMsOptions("USE_MS_OPTIONS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ms_options", "USE_MS_OPTIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseMsOptions("USE_MS_OPTIONS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ms_options", "USE_MS_OPTIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseNextserver(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_use_nextserver"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseNextserver("USE_NEXTSERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "USE_NEXTSERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseNextserver("USE_NEXTSERVER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "USE_NEXTSERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseOptions(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_use_options"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseOptions("USE_OPTIONS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "USE_OPTIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseOptions("USE_OPTIONS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "USE_OPTIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UsePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_use_pxe_lease_time"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUsePxeLeaseTime("USE_PXE_LEASE_TIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "USE_PXE_LEASE_TIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUsePxeLeaseTime("USE_PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "USE_PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseSnmp3Credential(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_use_snmp3_credential"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseSnmp3Credential("USE_SNMP3_CREDENTIAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp3_credential", "USE_SNMP3_CREDENTIAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseSnmp3Credential("USE_SNMP3_CREDENTIAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp3_credential", "USE_SNMP3_CREDENTIAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseSnmpCredential(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddress.test_use_snmp_credential"
	var v dhcp.Fixedaddress

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseSnmpCredential("USE_SNMP_CREDENTIAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp_credential", "USE_SNMP_CREDENTIAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseSnmpCredential("USE_SNMP_CREDENTIAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp_credential", "USE_SNMP_CREDENTIAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckFixedaddressExists(ctx context.Context, resourceName string, v *dhcp.Fixedaddress) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			FixedaddressAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForFixedaddress).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetFixedaddressResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetFixedaddressResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckFixedaddressDestroy(ctx context.Context, v *dhcp.Fixedaddress) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			FixedaddressAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForFixedaddress).
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

func testAccCheckFixedaddressDisappears(ctx context.Context, v *dhcp.Fixedaddress) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			FixedaddressAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccFixedaddressBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test" {
}
`)
}

func testAccFixedaddressRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test__ref" {
    _ref = %q
}
`, ref)
}

func testAccFixedaddressAgentCircuitId(agentCircuitId string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_agent_circuit_id" {
    agent_circuit_id = %q
}
`, agentCircuitId)
}

func testAccFixedaddressAgentRemoteId(agentRemoteId string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_agent_remote_id" {
    agent_remote_id = %q
}
`, agentRemoteId)
}

func testAccFixedaddressAllowTelnet(allowTelnet string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_allow_telnet" {
    allow_telnet = %q
}
`, allowTelnet)
}

func testAccFixedaddressAlwaysUpdateDns(alwaysUpdateDns string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_always_update_dns" {
    always_update_dns = %q
}
`, alwaysUpdateDns)
}

func testAccFixedaddressBootfile(bootfile string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_bootfile" {
    bootfile = %q
}
`, bootfile)
}

func testAccFixedaddressBootserver(bootserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_bootserver" {
    bootserver = %q
}
`, bootserver)
}

func testAccFixedaddressCliCredentials(cliCredentials string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_cli_credentials" {
    cli_credentials = %q
}
`, cliCredentials)
}

func testAccFixedaddressClientIdentifierPrependZero(clientIdentifierPrependZero string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_client_identifier_prepend_zero" {
    client_identifier_prepend_zero = %q
}
`, clientIdentifierPrependZero)
}

func testAccFixedaddressCloudInfo(cloudInfo string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_cloud_info" {
    cloud_info = %q
}
`, cloudInfo)
}

func testAccFixedaddressComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccFixedaddressDdnsDomainname(ddnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_ddns_domainname" {
    ddns_domainname = %q
}
`, ddnsDomainname)
}

func testAccFixedaddressDdnsHostname(ddnsHostname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_ddns_hostname" {
    ddns_hostname = %q
}
`, ddnsHostname)
}

func testAccFixedaddressDenyBootp(denyBootp string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_deny_bootp" {
    deny_bootp = %q
}
`, denyBootp)
}

func testAccFixedaddressDeviceDescription(deviceDescription string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_device_description" {
    device_description = %q
}
`, deviceDescription)
}

func testAccFixedaddressDeviceLocation(deviceLocation string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_device_location" {
    device_location = %q
}
`, deviceLocation)
}

func testAccFixedaddressDeviceType(deviceType string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_device_type" {
    device_type = %q
}
`, deviceType)
}

func testAccFixedaddressDeviceVendor(deviceVendor string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_device_vendor" {
    device_vendor = %q
}
`, deviceVendor)
}

func testAccFixedaddressDhcpClientIdentifier(dhcpClientIdentifier string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_dhcp_client_identifier" {
    dhcp_client_identifier = %q
}
`, dhcpClientIdentifier)
}

func testAccFixedaddressDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccFixedaddressDisableDiscovery(disableDiscovery string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_disable_discovery" {
    disable_discovery = %q
}
`, disableDiscovery)
}

func testAccFixedaddressDiscoveredData(discoveredData string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_discovered_data" {
    discovered_data = %q
}
`, discoveredData)
}

func testAccFixedaddressEnableDdns(enableDdns string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_enable_ddns" {
    enable_ddns = %q
}
`, enableDdns)
}

func testAccFixedaddressEnableImmediateDiscovery(enableImmediateDiscovery string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_enable_immediate_discovery" {
    enable_immediate_discovery = %q
}
`, enableImmediateDiscovery)
}

func testAccFixedaddressEnablePxeLeaseTime(enablePxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_enable_pxe_lease_time" {
    enable_pxe_lease_time = %q
}
`, enablePxeLeaseTime)
}

func testAccFixedaddressExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccFixedaddressIgnoreDhcpOptionListRequest(ignoreDhcpOptionListRequest string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_ignore_dhcp_option_list_request" {
    ignore_dhcp_option_list_request = %q
}
`, ignoreDhcpOptionListRequest)
}

func testAccFixedaddressIpv4addr(ipv4addr string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_ipv4addr" {
    ipv4addr = %q
}
`, ipv4addr)
}

func testAccFixedaddressFuncCall(funcCall string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_func_call" {
    func_call = %q
}
`, funcCall)
}

func testAccFixedaddressLogicFilterRules(logicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_logic_filter_rules" {
    logic_filter_rules = %q
}
`, logicFilterRules)
}

func testAccFixedaddressMac(mac string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_mac" {
    mac = %q
}
`, mac)
}

func testAccFixedaddressMatchClient(matchClient string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_match_client" {
    match_client = %q
}
`, matchClient)
}

func testAccFixedaddressMsAdUserData(msAdUserData string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_ms_ad_user_data" {
    ms_ad_user_data = %q
}
`, msAdUserData)
}

func testAccFixedaddressMsOptions(msOptions string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_ms_options" {
    ms_options = %q
}
`, msOptions)
}

func testAccFixedaddressMsServer(msServer string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_ms_server" {
    ms_server = %q
}
`, msServer)
}

func testAccFixedaddressName(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_name" {
    name = %q
}
`, name)
}

func testAccFixedaddressNetwork(network string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_network" {
    network = %q
}
`, network)
}

func testAccFixedaddressNetworkView(networkView string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_network_view" {
    network_view = %q
}
`, networkView)
}

func testAccFixedaddressNextserver(nextserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_nextserver" {
    nextserver = %q
}
`, nextserver)
}

func testAccFixedaddressOptions(options string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_options" {
    options = %q
}
`, options)
}

func testAccFixedaddressPxeLeaseTime(pxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_pxe_lease_time" {
    pxe_lease_time = %q
}
`, pxeLeaseTime)
}

func testAccFixedaddressReservedInterface(reservedInterface string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_reserved_interface" {
    reserved_interface = %q
}
`, reservedInterface)
}

func testAccFixedaddressRestartIfNeeded(restartIfNeeded string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_restart_if_needed" {
    restart_if_needed = %q
}
`, restartIfNeeded)
}

func testAccFixedaddressSnmp3Credential(snmp3Credential string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_snmp3_credential" {
    snmp3_credential = %q
}
`, snmp3Credential)
}

func testAccFixedaddressSnmpCredential(snmpCredential string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_snmp_credential" {
    snmp_credential = %q
}
`, snmpCredential)
}

func testAccFixedaddressTemplate(template string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_template" {
    template = %q
}
`, template)
}

func testAccFixedaddressUseBootfile(useBootfile string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_use_bootfile" {
    use_bootfile = %q
}
`, useBootfile)
}

func testAccFixedaddressUseBootserver(useBootserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_use_bootserver" {
    use_bootserver = %q
}
`, useBootserver)
}

func testAccFixedaddressUseCliCredentials(useCliCredentials string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_use_cli_credentials" {
    use_cli_credentials = %q
}
`, useCliCredentials)
}

func testAccFixedaddressUseDdnsDomainname(useDdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_use_ddns_domainname" {
    use_ddns_domainname = %q
}
`, useDdnsDomainname)
}

func testAccFixedaddressUseDenyBootp(useDenyBootp string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_use_deny_bootp" {
    use_deny_bootp = %q
}
`, useDenyBootp)
}

func testAccFixedaddressUseEnableDdns(useEnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_use_enable_ddns" {
    use_enable_ddns = %q
}
`, useEnableDdns)
}

func testAccFixedaddressUseIgnoreDhcpOptionListRequest(useIgnoreDhcpOptionListRequest string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_use_ignore_dhcp_option_list_request" {
    use_ignore_dhcp_option_list_request = %q
}
`, useIgnoreDhcpOptionListRequest)
}

func testAccFixedaddressUseLogicFilterRules(useLogicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_use_logic_filter_rules" {
    use_logic_filter_rules = %q
}
`, useLogicFilterRules)
}

func testAccFixedaddressUseMsOptions(useMsOptions string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_use_ms_options" {
    use_ms_options = %q
}
`, useMsOptions)
}

func testAccFixedaddressUseNextserver(useNextserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_use_nextserver" {
    use_nextserver = %q
}
`, useNextserver)
}

func testAccFixedaddressUseOptions(useOptions string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_use_options" {
    use_options = %q
}
`, useOptions)
}

func testAccFixedaddressUsePxeLeaseTime(usePxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_use_pxe_lease_time" {
    use_pxe_lease_time = %q
}
`, usePxeLeaseTime)
}

func testAccFixedaddressUseSnmp3Credential(useSnmp3Credential string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_use_snmp3_credential" {
    use_snmp3_credential = %q
}
`, useSnmp3Credential)
}

func testAccFixedaddressUseSnmpCredential(useSnmpCredential string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddress" "test_use_snmp_credential" {
    use_snmp_credential = %q
}
`, useSnmpCredential)
}
