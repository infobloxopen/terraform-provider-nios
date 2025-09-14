package dhcp_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

//TODO : Following Objects required configurations to be defined in the grid
// - CLI Credentials
// - Restart if Needed
// - MS Server , Options

//TODO : add tests
// The following require additional resource/data source objects to be supported.
// - Logic Filter Rules
// - Reserved Interface
// - Discovery
// - Fixed Address Template
// - Network - To randomize the test IPs

//TODO : OBJECTS TO BE PRESENT IN GRID FOR TESTS
// - Network View - default , test_fixed_address
// - Ipv4 Network - 15.0.0.0/24 , 16.0.0.0/24

var readableAttributesForFixedaddress = "agent_circuit_id,agent_remote_id,allow_telnet,always_update_dns,bootfile,bootserver,cli_credentials,client_identifier_prepend_zero,cloud_info,comment,ddns_domainname,ddns_hostname,deny_bootp,device_description,device_location,device_type,device_vendor,dhcp_client_identifier,disable,disable_discovery,discover_now_status,discovered_data,enable_ddns,enable_pxe_lease_time,extattrs,ignore_dhcp_option_list_request,ipv4addr,is_invalid_mac,logic_filter_rules,mac,match_client,ms_ad_user_data,ms_options,ms_server,name,network,network_view,nextserver,options,pxe_lease_time,reserved_interface,snmp3_credential,snmp_credential,use_bootfile,use_bootserver,use_cli_credentials,use_ddns_domainname,use_deny_bootp,use_enable_ddns,use_ignore_dhcp_option_list_request,use_logic_filter_rules,use_ms_options,use_nextserver,use_options,use_pxe_lease_time,use_snmp3_credential,use_snmp_credential"

func TestAccFixedaddressResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test"
	var v dhcp.Fixedaddress
	ip := "15.0.0.1"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressBasicConfig(ip, "CIRCUIT_ID", agentCircuitID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", ip),
					resource.TestCheckResourceAttr(resourceName, "match_client", "CIRCUIT_ID"),
					resource.TestCheckResourceAttr(resourceName, "agent_circuit_id", fmt.Sprintf("%v", agentCircuitID)),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "allow_telnet", "false"),
					resource.TestCheckResourceAttr(resourceName, "always_update_dns", "false"),
					resource.TestCheckResourceAttr(resourceName, "client_identifier_prepend_zero", "false"),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "false"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "disable_discovery", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_cli_credentials", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ms_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_snmp3_credential", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_snmp_credential", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_fixed_address.test"
	var v dhcp.Fixedaddress
	ip := "15.0.0.2"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFixedaddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFixedaddressBasicConfig(ip, "CIRCUIT_ID", agentCircuitID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					testAccCheckFixedaddressDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccFixedaddressResource_AgentCircuitId(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_agent_circuit_id"
	var v dhcp.Fixedaddress
	ip := "15.0.0.3"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressAgentCircuitId(ip, "CIRCUIT_ID", 30),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "agent_circuit_id", "30"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressAgentCircuitId(ip, "CIRCUIT_ID", 32),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "agent_circuit_id", "32"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_AgentRemoteId(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_agent_remote_id"
	var v dhcp.Fixedaddress
	ip := "15.0.0.4"
	agentRemoteID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressAgentRemoteId(ip, "REMOTE_ID", agentRemoteID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "agent_remote_id", fmt.Sprintf("%v", agentRemoteID)),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressAgentRemoteId(ip, "REMOTE_ID", agentRemoteID+10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "agent_remote_id", fmt.Sprintf("%v", agentRemoteID+10)),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_AllowTelnet(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_allow_telnet"
	var v dhcp.Fixedaddress
	ip := "15.0.0.5"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressAllowTelnet(ip, "CIRCUIT_ID", agentCircuitID, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_telnet", "true"),
				),
			},
		},
	})
}

func TestAccFixedaddressResource_AlwaysUpdateDns(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_always_update_dns"
	var v dhcp.Fixedaddress
	ip := "15.0.0.6"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressAlwaysUpdateDns(ip, "CIRCUIT_ID", agentCircuitID, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "always_update_dns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressAlwaysUpdateDns(ip, "CIRCUIT_ID", agentCircuitID, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "always_update_dns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Bootfile(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_bootfile"
	var v dhcp.Fixedaddress
	ip := "15.0.0.7"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressBootfile(ip, "CIRCUIT_ID", agentCircuitID, "file", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "file"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressBootfile(ip, "CIRCUIT_ID", agentCircuitID, "file1", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "file1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Bootserver(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_bootserver"
	var v dhcp.Fixedaddress
	ip := "15.0.0.8"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressBootserver(ip, "CIRCUIT_ID", agentCircuitID, "boot_server_example.com", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "boot_server_example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressBootserver(ip, "CIRCUIT_ID", agentCircuitID, "boot_server_updated_example.com", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "boot_server_updated_example.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_CliCredentials(t *testing.T) {
	//t.Skip("Skipping test as CLI Credential are not set up in the GRID")
	var resourceName = "nios_dhcp_fixed_address.test_cli_credentials"
	var v dhcp.Fixedaddress
	ip := "15.0.0.9"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressCliCredentialsSSH(ip, "CIRCUIT_ID", agentCircuitID, "Comment for CLI Credentials", "NIOS_USER", "NIOS_PASSWORD", "SSH", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.0.comment", "Comment for CLI Credentials"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.0.user", "NIOS_USER"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.0.password", "NIOS_PASSWORD"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.0.credential_type", "SSH"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.0.credential_group", "default"),
				),
			},
			//Update and Read
			{
				Config: testAccFixedaddressCliCredentials(ip, "CIRCUIT_ID", agentCircuitID, "Updated Comment for CLI Credentials", "NIOS_USER", "NIOS_PASSWORD", "TELNET", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.comment", "Updated Comment for CLI Credentials"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.user", "NIOS_USER"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.password", "NIOS_PASSWORD"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.credential_type", "TELNET"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.credential_group", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressCliCredentials(ip, "CIRCUIT_ID", agentCircuitID, "Updated Comment for CLI Credentials", "NIOS_USER", "NIOS_PASSWORD", "ENABLE_TELNET", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.comment", "Updated Comment for CLI Credentials"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.user", "NIOS_USER"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.password", "NIOS_PASSWORD"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.credential_type", "ENABLE_TELNET"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.credential_group", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressCliCredentials(ip, "CIRCUIT_ID", agentCircuitID, "Updated Comment for CLI Credentials", "NIOS_USER", "NIOS_PASSWORD", "ENABLE_SSH", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.comment", "Updated Comment for CLI Credentials"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.user", "NIOS_USER"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.password", "NIOS_PASSWORD"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.credential_type", "ENABLE_SSH"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.credential_group", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_ClientIdentifierPrependZero(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_client_identifier_prepend_zero"
	var v dhcp.Fixedaddress
	ip := "15.0.0.10"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressClientIdentifierPrependZero(ip, "CIRCUIT_ID", agentCircuitID, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_identifier_prepend_zero", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressClientIdentifierPrependZero(ip, "CIRCUIT_ID", agentCircuitID, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_identifier_prepend_zero", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_comment"
	var v dhcp.Fixedaddress
	ip := "15.0.0.11"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressComment(ip, "CIRCUIT_ID", agentCircuitID, "Comment for Fixed Address"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for Fixed Address"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressComment(ip, "CIRCUIT_ID", agentCircuitID, "Updated Comment for Fixed Address"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated Comment for Fixed Address"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_DdnsDomainname(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_ddns_domainname"
	var v dhcp.Fixedaddress
	ip := "15.0.0.12"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressDdnsDomainname(ip, "CIRCUIT_ID", agentCircuitID, "ddns_domain.name", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "ddns_domain.name"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressDdnsDomainname(ip, "CIRCUIT_ID", agentCircuitID, "updated_ddns_domain.name", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "updated_ddns_domain.name"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_DdnsHostname(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_ddns_hostname"
	var v dhcp.Fixedaddress
	ip := "15.0.0.13"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressDdnsHostname(ip, "CIRCUIT_ID", agentCircuitID, "ddns_host.name"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_hostname", "ddns_host.name"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressDdnsHostname(ip, "CIRCUIT_ID", agentCircuitID, "updated_ddns_host.name"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_hostname", "updated_ddns_host.name"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_DenyBootp(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_deny_bootp"
	var v dhcp.Fixedaddress
	ip := "15.0.0.14"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressDenyBootp(ip, "CIRCUIT_ID", agentCircuitID, "true", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressDenyBootp(ip, "CIRCUIT_ID", agentCircuitID, "false", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_DeviceDescription(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_device_description"
	var v dhcp.Fixedaddress
	ip := "15.0.0.15"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressDeviceDescription(ip, "CIRCUIT_ID", agentCircuitID, "DEVICE_DESCRIPTION"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_description", "DEVICE_DESCRIPTION"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressDeviceDescription(ip, "CIRCUIT_ID", agentCircuitID, "DEVICE_DESCRIPTION_UPDATED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_description", "DEVICE_DESCRIPTION_UPDATED"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_DeviceLocation(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_device_location"
	var v dhcp.Fixedaddress
	ip := "15.0.0.16"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressDeviceLocation(ip, "CIRCUIT_ID", agentCircuitID, "DEVICE_LOCATION"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_location", "DEVICE_LOCATION"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressDeviceLocation(ip, "CIRCUIT_ID", agentCircuitID, "DEVICE_LOCATION_UPDATED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_location", "DEVICE_LOCATION_UPDATED"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_DeviceType(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_device_type"
	var v dhcp.Fixedaddress
	ip := "15.0.0.17"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressDeviceType(ip, "CIRCUIT_ID", agentCircuitID, "DEVICE_TYPE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_type", "DEVICE_TYPE"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressDeviceType(ip, "CIRCUIT_ID", agentCircuitID, "DEVICE_TYPE_UPDATED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_type", "DEVICE_TYPE_UPDATED"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_DeviceVendor(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_device_vendor"
	var v dhcp.Fixedaddress
	ip := "15.0.0.18"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressDeviceVendor(ip, "CIRCUIT_ID", agentCircuitID, "DEVICE_VENDOR"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_vendor", "DEVICE_VENDOR"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressDeviceVendor(ip, "CIRCUIT_ID", agentCircuitID, "DEVICE_VENDOR_UPDATED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_vendor", "DEVICE_VENDOR_UPDATED"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_DhcpClientIdentifier(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_dhcp_client_identifier"
	var v dhcp.Fixedaddress
	ip := "15.0.0.19"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressDhcpClientIdentifier(ip, "CLIENT_ID", "DHCP_CLIENT_IDENTIFIER"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_client_identifier", "DHCP_CLIENT_IDENTIFIER"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressDhcpClientIdentifier(ip, "CLIENT_ID", "DHCP_CLIENT_IDENTIFIER_UPDATED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_client_identifier", "DHCP_CLIENT_IDENTIFIER_UPDATED"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Disable(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_disable"
	var v dhcp.Fixedaddress
	ip := "15.0.0.20"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressDisable(ip, "CIRCUIT_ID", agentCircuitID, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressDisable(ip, "CIRCUIT_ID", agentCircuitID, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_DisableDiscovery(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_disable_discovery"
	var v dhcp.Fixedaddress
	ip := "15.0.0.21"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressDisableDiscovery(ip, "CIRCUIT_ID", agentCircuitID, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_discovery", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressDisableDiscovery(ip, "CIRCUIT_ID", agentCircuitID, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_discovery", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_EnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_enable_ddns"
	var v dhcp.Fixedaddress
	ip := "15.0.0.22"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressEnableDdns(ip, "CIRCUIT_ID", agentCircuitID, "true", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressEnableDdns(ip, "CIRCUIT_ID", agentCircuitID, "false", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_EnableImmediateDiscovery(t *testing.T) {
	t.Skip("Skipping test as Discovery is not supported")
	var resourceName = "nios_dhcp_fixed_address.test_enable_immediate_discovery"
	var v dhcp.Fixedaddress
	ip := "15.0.0.23"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressEnableImmediateDiscovery(ip, "CIRCUIT_ID", agentCircuitID, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_immediate_discovery", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressEnableImmediateDiscovery(ip, "CIRCUIT_ID", agentCircuitID, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_immediate_discovery", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_EnablePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_enable_pxe_lease_time"
	var v dhcp.Fixedaddress
	ip := "15.0.0.24"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressEnablePxeLeaseTime(ip, "CIRCUIT_ID", agentCircuitID, "true", 3600, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressEnablePxeLeaseTime(ip, "CIRCUIT_ID", agentCircuitID, "false", 3600, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_extattrs"
	var v dhcp.Fixedaddress
	ip := "15.0.0.25"
	agentCircuitID := acctest.RandomNumber(1000)
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressExtAttrs(ip, "CIRCUIT_ID", agentCircuitID, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressExtAttrs(ip, "CIRCUIT_ID", agentCircuitID, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_IgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_ignore_dhcp_option_list_request"
	var v dhcp.Fixedaddress
	ip := "15.0.0.26"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressIgnoreDhcpOptionListRequest(ip, "CIRCUIT_ID", agentCircuitID, "true", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressIgnoreDhcpOptionListRequest(ip, "CIRCUIT_ID", agentCircuitID, "false", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Ipv4addr(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_ipv4addr"
	var v dhcp.Fixedaddress
	ip := "15.0.0.27"
	IPUpdated := "15.0.0.60"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressIpv4addr(ip, "CIRCUIT_ID", agentCircuitID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", ip),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressIpv4addr(IPUpdated, "CIRCUIT_ID", agentCircuitID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", IPUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// TestAccFixedaddressResource_FuncCall tests the "func_call" attribute functionality
// which allocates IPv4 addresses using next_available_ip. Since func_call attribute can't be
// updated, the comment is updated to demonstrate an update to the resource
func TestAccFixedaddressResource_FuncCall(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_func_call"
	var v dhcp.Fixedaddress
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressFuncCall("CIRCUIT_ID", agentCircuitID, "ipv4addr", "next_available_ip", "ips", "network", "15.0.0.0/24", "Original Function Call"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressFuncCall("CIRCUIT_ID", agentCircuitID, "comment", "next_available_ip", "ips", "network", "15.0.0.0/24", "Function Call with Update"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
				),
			},
		},
	})
}

func TestAccFixedaddressResource_LogicFilterRules(t *testing.T) {
	t.Skip("Skipping test as support for MAC/NAC/Option logic filter rules is not implemented yet")
	var resourceName = "nios_dhcp_fixed_address.test_logic_filter_rules"
	var v dhcp.Fixedaddress
	ip := "15.0.0.28"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressLogicFilterRules(ip, "CIRCUIT_ID", agentCircuitID, "FILTER_NAME", "MAC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.filter", "FILTER_NAME"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.type", "MAC"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressLogicFilterRules(ip, "CIRCUIT_ID", agentCircuitID, "FILTER_NAME_UPDATED", "Options"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.filter", "FILTER_NAME_UPDATED"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.type", "Option"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Mac(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_mac"
	var v dhcp.Fixedaddress
	ip := "15.0.0.29"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressMac(ip, "MAC_ADDRESS", "00:1a:2b:3c:4d:5e"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mac", "00:1a:2b:3c:4d:5e"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressMac(ip, "MAC_ADDRESS", "10:9a:dd:ee:ff:01"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mac", "10:9a:dd:ee:ff:01"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressMac(ip, "RESERVED", "00:00:00:00:00:00"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_client", "RESERVED"),
					resource.TestCheckResourceAttr(resourceName, "mac", "00:00:00:00:00:00"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_MatchClient(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_match_client"
	var v dhcp.Fixedaddress
	ip := "15.0.0.30"
	agentCircuitID := acctest.RandomNumber(1000)
	agentRemoteID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressMatchClientCircuitID(ip, "CIRCUIT_ID", agentCircuitID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_client", "CIRCUIT_ID"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressMatchClientRemoteID(ip, "REMOTE_ID", agentRemoteID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_client", "REMOTE_ID"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_MsOptions(t *testing.T) {
	t.Skip("Skipping test as appropriate MS options are not set up in the GRID")
	var resourceName = "nios_dhcp_fixed_address.test_ms_options"
	var v dhcp.Fixedaddress
	ip := "15.0.0.31"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressMsOptions(ip, "CIRCUIT_ID", agentCircuitID, "time-offset", "2", "50"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_options", "50"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressMsOptions(ip, "CIRCUIT_ID", agentCircuitID, "dhcp-lease-time", "56", "100"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_options", "100"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_MsServer(t *testing.T) {
	t.Skip("Skipping test as MS Server is supposed to be set up on the GRID")
	var resourceName = "nios_dhcp_fixed_address.test_ms_server"
	var v dhcp.Fixedaddress
	ip := "50.0.0.132"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressMsServer(ip, "MAC_ADDRESS", "00:1a:6b:3c:4d:5e", "10.34.98.68", "ms_server", "example_fixed_address"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_server", "10.34.98.68"),
				),
			},
			// Update and Read
			//{
			//	Config: testAccFixedaddressMsServer(ip, "CIRCUIT_ID", agentCircuitID, "2.2.2.2"),
			//	Check: resource.ComposeTestCheckFunc(
			//		testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
			//		resource.TestCheckResourceAttr(resourceName, "ms_server", "2.2.2.2"),
			//	),
			//},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_name"
	var v dhcp.Fixedaddress
	ip := "15.0.0.33"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressName(ip, "CIRCUIT_ID", agentCircuitID, "example_fixed_address"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "example_fixed_address"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressName(ip, "CIRCUIT_ID", agentCircuitID, "example_updated_fixed_address"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "example_updated_fixed_address"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Network(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_network"
	var v dhcp.Fixedaddress
	ip := "15.0.0.34"
	ipUpdated := acctest.RandomIPWithSpecificOctetsSet("16.0.0")
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressNetwork(ip, "CIRCUIT_ID", agentCircuitID, "15.0.0.0/24"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", "15.0.0.0/24"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressNetwork(ipUpdated, "CIRCUIT_ID", agentCircuitID, "16.0.0.0/24"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", "16.0.0.0/24"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_NetworkView(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_network_view"
	var v dhcp.Fixedaddress
	ip := "15.0.0.35"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressNetworkView(ip, "CIRCUIT_ID", agentCircuitID, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Nextserver(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_nextserver"
	var v dhcp.Fixedaddress
	ip := "15.0.0.36"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressNextserver(ip, "CIRCUIT_ID", agentCircuitID, "example_next_server.com", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", "example_next_server.com"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressNextserver(ip, "CIRCUIT_ID", agentCircuitID, "example_updated_next_server.com", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", "example_updated_next_server.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Options(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_options"
	var v dhcp.Fixedaddress
	ip := "15.0.0.37"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressOptions(ip, "CIRCUIT_ID", agentCircuitID, "time-offset", "50", true, "2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "time-offset"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "50"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressOptions(ip, "CIRCUIT_ID", agentCircuitID, "dhcp-lease-time", "7200", true, "51"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "7200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_PxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_pxe_lease_time"
	var v dhcp.Fixedaddress
	ip := "15.0.0.38"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressPxeLeaseTime(ip, "CIRCUIT_ID", agentCircuitID, "3600", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "3600"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressPxeLeaseTime(ip, "CIRCUIT_ID", agentCircuitID, "4800", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "4800"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_ReservedInterface(t *testing.T) {
	t.Skip("Skipping test as reserved_interface is not implemented yet")
	var resourceName = "nios_dhcp_fixed_address.test_reserved_interface"
	var v dhcp.Fixedaddress
	ip := "15.0.0.39"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressReservedInterface(ip, "CIRCUIT_ID", agentCircuitID, "ref"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "reserved_interface", "ref"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressReservedInterface(ip, "CIRCUIT_ID", agentCircuitID, "ref2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "reserved_interface", "ref2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_RestartIfNeeded(t *testing.T) {
	t.Skip("Skipping test as it requires a CP member to be configured")
	var resourceName = "nios_dhcp_fixed_address.test_restart_if_needed"
	var v dhcp.Fixedaddress
	ip := "15.0.0.40"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressRestartIfNeeded(ip, "CIRCUIT_ID", agentCircuitID, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "restart_if_needed", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressRestartIfNeeded(ip, "CIRCUIT_ID", agentCircuitID, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "restart_if_needed", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Snmp3Credential(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_snmp3_credential"
	var v dhcp.Fixedaddress
	ip := "16.0.0.121"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressSnmp3Credential(ip, "CIRCUIT_ID", agentCircuitID, "snmp", "MD5", "snmp1234", "3DES", "snmp1234", "SNMP3 Credential Comment", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.user", "snmp"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.authentication_protocol", "MD5"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.authentication_password", "snmp1234"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.privacy_protocol", "3DES"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.privacy_password", "snmp1234"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.comment", "SNMP3 Credential Comment"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.credential_group", "default"),
				),
			},
			//Update and Read
			{
				Config: testAccFixedaddressSnmp3Credential(ip, "CIRCUIT_ID", agentCircuitID, "SNMP3_USER_UPDATE", "SHA-224", "AUTH_PASSWORD", "AES-256", "PRIVACY_PASSWORD", "SNMP3 Credential Comment Updated", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.user", "SNMP3_USER_UPDATE"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.authentication_protocol", "SHA-224"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.authentication_password", "AUTH_PASSWORD"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.privacy_protocol", "AES-256"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.privacy_password", "PRIVACY_PASSWORD"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.comment", "SNMP3 Credential Comment Updated"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.credential_group", "default"),
				),
			},
			//Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_SnmpCredential(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_snmp_credential"
	var v dhcp.Fixedaddress
	ip := "15.0.0.42"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressSnmpCredential(ip, "CIRCUIT_ID", agentCircuitID, "COMMUNITY_STRING", "SNMP Credential Comment", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "snmp_credential.community_string", "COMMUNITY_STRING"),
					resource.TestCheckResourceAttr(resourceName, "snmp_credential.comment", "SNMP Credential Comment"),
					resource.TestCheckResourceAttr(resourceName, "snmp_credential.credential_group", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressSnmpCredential(ip, "CIRCUIT_ID", agentCircuitID, "COMMUNITY_STRING_UPDATED", "SNMP Credential Comment Updated", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "snmp_credential.community_string", "COMMUNITY_STRING_UPDATED"),
					resource.TestCheckResourceAttr(resourceName, "snmp_credential.comment", "SNMP Credential Comment Updated"),
					resource.TestCheckResourceAttr(resourceName, "snmp_credential.credential_group", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_Template(t *testing.T) {
	t.Skip("Skipping test as Fixed Address Template is not implemented yet")
	var resourceName = "nios_dhcp_fixed_address.test_template"
	var v dhcp.Fixedaddress
	ip := "15.0.0.43"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressTemplate(ip, "CIRCUIT_ID", agentCircuitID, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template", "TEMPLATE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressTemplate(ip, "CIRCUIT_ID", agentCircuitID, ""),
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
	var resourceName = "nios_dhcp_fixed_address.test_use_bootfile"
	var v dhcp.Fixedaddress
	ip := "15.0.0.44"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseBootfile(ip, "CIRCUIT_ID", agentCircuitID, "true", "file"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseBootfile(ip, "CIRCUIT_ID", agentCircuitID, "false", "file"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseBootserver(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_use_bootserver"
	var v dhcp.Fixedaddress
	ip := "15.0.0.45"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseBootserver(ip, "CIRCUIT_ID", agentCircuitID, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseBootserver(ip, "CIRCUIT_ID", agentCircuitID, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseCliCredentials(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_use_cli_credentials"
	var v dhcp.Fixedaddress
	ip := "15.0.0.46"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseCliCredentials(ip, "CIRCUIT_ID", agentCircuitID, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_cli_credentials", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseCliCredentials(ip, "CIRCUIT_ID", agentCircuitID, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_cli_credentials", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseDdnsDomainname(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_use_ddns_domainname"
	var v dhcp.Fixedaddress
	ip := "15.0.0.47"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseDdnsDomainname(ip, "CIRCUIT_ID", agentCircuitID, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseDdnsDomainname(ip, "CIRCUIT_ID", agentCircuitID, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseDenyBootp(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_use_deny_bootp"
	var v dhcp.Fixedaddress
	ip := "15.0.0.48"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseDenyBootp(ip, "CIRCUIT_ID", agentCircuitID, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseDenyBootp(ip, "CIRCUIT_ID", agentCircuitID, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseEnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_use_enable_ddns"
	var v dhcp.Fixedaddress
	ip := "15.0.0.49"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseEnableDdns(ip, "CIRCUIT_ID", agentCircuitID, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseEnableDdns(ip, "CIRCUIT_ID", agentCircuitID, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseIgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_use_ignore_dhcp_option_list_request"
	var v dhcp.Fixedaddress
	ip := "15.0.0.50"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseIgnoreDhcpOptionListRequest(ip, "CIRCUIT_ID", agentCircuitID, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseIgnoreDhcpOptionListRequest(ip, "CIRCUIT_ID", agentCircuitID, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseLogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_use_logic_filter_rules"
	var v dhcp.Fixedaddress
	ip := "15.0.0.51"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseLogicFilterRules(ip, "CIRCUIT_ID", agentCircuitID, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseLogicFilterRules(ip, "CIRCUIT_ID", agentCircuitID, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseMsOptions(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_use_ms_options"
	var v dhcp.Fixedaddress
	ip := "15.0.0.52"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseMsOptions(ip, "CIRCUIT_ID", agentCircuitID, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ms_options", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseMsOptions(ip, "CIRCUIT_ID", agentCircuitID, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ms_options", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseNextserver(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_use_nextserver"
	var v dhcp.Fixedaddress
	ip := "15.0.0.53"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseNextserver(ip, "CIRCUIT_ID", agentCircuitID, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseNextserver(ip, "CIRCUIT_ID", agentCircuitID, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseOptions(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_use_options"
	var v dhcp.Fixedaddress
	ip := "15.0.0.54"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseOptions(ip, "CIRCUIT_ID", agentCircuitID, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseOptions(ip, "CIRCUIT_ID", agentCircuitID, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UsePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_use_pxe_lease_time"
	var v dhcp.Fixedaddress
	ip := "15.0.0.55"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUsePxeLeaseTime(ip, "CIRCUIT_ID", agentCircuitID, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUsePxeLeaseTime(ip, "CIRCUIT_ID", agentCircuitID, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseSnmp3Credential(t *testing.T) {
	t.Skip("Skipping test as SNMP3 Credential is not supported yet")
	var resourceName = "nios_dhcp_fixed_address.test_use_snmp3_credential"
	var v dhcp.Fixedaddress
	ip := "15.0.0.56"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseSnmp3Credential(ip, "CIRCUIT_ID", agentCircuitID, "true", "SNMP3_USER", "MD5", "AUTH_PASSWORD", "3DES", "PRIVACY_PASSWORD", "SNMP3 Credential Comment", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp3_credential", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseSnmp3Credential(ip, "CIRCUIT_ID", agentCircuitID, "false", "SNMP3_USER", "MD5", "AUTH_PASSWORD", "3DES", "PRIVACY_PASSWORD", "SNMP3 Credential Comment", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp3_credential", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressResource_UseSnmpCredential(t *testing.T) {
	t.Skip("Skipping test as SNMP Credential are not set up in the GRID")
	var resourceName = "nios_dhcp_fixed_address.test_use_snmp_credential"
	var v dhcp.Fixedaddress
	ip := "15.0.0.57"
	agentCircuitID := acctest.RandomNumber(1000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddressUseSnmpCredential(ip, "CIRCUIT_ID", agentCircuitID, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp_credential", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddressUseSnmpCredential(ip, "CIRCUIT_ID", agentCircuitID, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp_credential", "false"),
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

func testAccFixedaddressBasicConfig(ip4addr, matchClient string, agentCircuitId int) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
}
`, ip4addr, matchClient, agentCircuitId)
}

// All test config generators now include the basic fields and the field under test

func testAccFixedaddressAgentCircuitId(ip, matchClient string, agentCircuitId int) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_agent_circuit_id" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
}
`, ip, matchClient, agentCircuitId)
}

func testAccFixedaddressAgentRemoteId(ip, matchClient string, agentRemoteId int) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_agent_remote_id" {
	ipv4addr = %q
	match_client = %q
	agent_remote_id = %d
}
`, ip, matchClient, agentRemoteId)
}

func testAccFixedaddressAllowTelnet(ip, matchClient string, agentCircuitID int, allowTelnet string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_allow_telnet" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	allow_telnet = %q
}
`, ip, matchClient, agentCircuitID, allowTelnet)
}

func testAccFixedaddressAlwaysUpdateDns(ip, matchClient string, agentCircuitID int, alwaysUpdateDns string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_always_update_dns" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	always_update_dns = %q
}
`, ip, matchClient, agentCircuitID, alwaysUpdateDns)
}

func testAccFixedaddressBootfile(ip, matchClient string, agentCircuitID int, bootfile string, useBootfile bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_bootfile" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	bootfile = %q
	use_bootfile = %t
}
`, ip, matchClient, agentCircuitID, bootfile, useBootfile)
}

func testAccFixedaddressBootserver(ip, matchClient string, agentCircuitID int, bootserver string, useBootServer bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_bootserver" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	bootserver = %q
	use_bootserver = %t
}
`, ip, matchClient, agentCircuitID, bootserver, useBootServer)
}

func testAccFixedaddressCliCredentials(ip, matchClient string, agentCircuitID int, cliCredComment, cliCredUser, cliCredPassword, cliCredType, cliCredGroup, useCLICredentials string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_cli_credentials" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	cli_credentials = [{
		comment          = "Comment for SSH Credentials"
		user             = "NIOS_USER"
		password         = "NIOS_PASSWORD"
		credential_type  = "SSH"
		credential_group = "default"
	},
	{
		comment          = %q
		user             = %q
		password         = %q
		credential_type  = %q
		credential_group = %q
	},
	]
	use_cli_credentials = %q
}
`, ip, matchClient, agentCircuitID, cliCredComment, cliCredUser, cliCredPassword, cliCredType, cliCredGroup, useCLICredentials)
}

func testAccFixedaddressCliCredentialsSSH(ip, matchClient string, agentCircuitID int, cliCredComment, cliCredUser, cliCredPassword, cliCredType, cliCredGroup, useCLICredentials string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_cli_credentials" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	cli_credentials = [{
		comment          = %q
		user             = %q
		password         = %q
		credential_type  = %q
		credential_group = %q
	}]
	use_cli_credentials = %q
}
`, ip, matchClient, agentCircuitID, cliCredComment, cliCredUser, cliCredPassword, cliCredType, cliCredGroup, useCLICredentials)
}

func testAccFixedaddressClientIdentifierPrependZero(ip, matchClient string, agentCircuitID int, clientIdentifierPrependZero string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_client_identifier_prepend_zero" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	client_identifier_prepend_zero = %q
}
`, ip, matchClient, agentCircuitID, clientIdentifierPrependZero)
}

func testAccFixedaddressComment(ip, matchClient string, agentCircuitID int, comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_comment" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	comment = %q
}
`, ip, matchClient, agentCircuitID, comment)
}

func testAccFixedaddressDdnsDomainname(ip, matchClient string, agentCircuitID int, ddnsDomainname string, useDDNSDomainName bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_ddns_domainname" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	ddns_domainname = %q
	use_ddns_domainname = %t
}
`, ip, matchClient, agentCircuitID, ddnsDomainname, useDDNSDomainName)
}

func testAccFixedaddressDdnsHostname(ip, matchClient string, agentCircuitID int, ddnsHostname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_ddns_hostname" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	ddns_hostname = %q
}
`, ip, matchClient, agentCircuitID, ddnsHostname)
}

func testAccFixedaddressDenyBootp(ip, matchClient string, agentCircuitID int, denyBootp string, useDenyBootp bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_deny_bootp" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	deny_bootp = %q
	use_deny_bootp = %t
}
`, ip, matchClient, agentCircuitID, denyBootp, useDenyBootp)
}

func testAccFixedaddressDeviceDescription(ip, matchClient string, agentCircuitID int, deviceDescription string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_device_description" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	device_description = %q
}
`, ip, matchClient, agentCircuitID, deviceDescription)
}

func testAccFixedaddressDeviceLocation(ip, matchClient string, agentCircuitID int, deviceLocation string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_device_location" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	device_location = %q
}
`, ip, matchClient, agentCircuitID, deviceLocation)
}

func testAccFixedaddressDeviceType(ip, matchClient string, agentCircuitID int, deviceType string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_device_type" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	device_type = %q
}
`, ip, matchClient, agentCircuitID, deviceType)
}

func testAccFixedaddressDeviceVendor(ip, matchClient string, agentCircuitID int, deviceVendor string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_device_vendor" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	device_vendor = %q
}
`, ip, matchClient, agentCircuitID, deviceVendor)
}

func testAccFixedaddressDhcpClientIdentifier(ip, matchClient string, dhcpClientIdentifier string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_dhcp_client_identifier" {
	ipv4addr = %q
	match_client = %q
	dhcp_client_identifier = %q
}
`, ip, matchClient, dhcpClientIdentifier)
}

func testAccFixedaddressDisable(ip, matchClient string, agentCircuitID int, disable string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_disable" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	disable = %q
}
`, ip, matchClient, agentCircuitID, disable)
}

func testAccFixedaddressDisableDiscovery(ip, matchClient string, agentCircuitID int, disableDiscovery string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_disable_discovery" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	disable_discovery = %q
}
`, ip, matchClient, agentCircuitID, disableDiscovery)
}

func testAccFixedaddressEnableDdns(ip, matchClient string, agentCircuitID int, enableDdns string, useEnableDDNS bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_enable_ddns" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	enable_ddns = %q
	use_enable_ddns = %t
}
`, ip, matchClient, agentCircuitID, enableDdns, useEnableDDNS)
}

func testAccFixedaddressEnableImmediateDiscovery(ip, matchClient string, agentCircuitID int, enableImmediateDiscovery string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_enable_immediate_discovery" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	enable_immediate_discovery = %q
}
`, ip, matchClient, agentCircuitID, enableImmediateDiscovery)
}

func testAccFixedaddressEnablePxeLeaseTime(ip, matchClient string, agentCircuitID int, enablePxeLeaseTime string, pxeLeaseTime int, usePXELeaseTime bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_enable_pxe_lease_time" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	enable_pxe_lease_time = %q
	pxe_lease_time = %d
	use_pxe_lease_time = %t
}
`, ip, matchClient, agentCircuitID, enablePxeLeaseTime, pxeLeaseTime, usePXELeaseTime)
}

func testAccFixedaddressExtAttrs(ip, matchClient string, agentCircuitID int, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
		  %s = %q
		`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_extattrs" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	extattrs = %s
}
`, ip, matchClient, agentCircuitID, extattrsStr)
}

func testAccFixedaddressFuncCall(matchClient string, agentCircuitID int, attributeName, objFunc, resultField, object, objectParameters, comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_func_call" {
	match_client = %q
	agent_circuit_id = %d
	func_call = {
		"attribute_name" = %q
		"object_function" = %q
		"result_field" = %q
		"object" = %q
		"object_parameters" = {
			"network" = %q
			"network_view" = "default"
		}
	}
	comment = %q
}
`, matchClient, agentCircuitID, attributeName, objFunc, resultField, object, objectParameters, comment)
}

func testAccFixedaddressIgnoreDhcpOptionListRequest(ip, matchClient string, agentCircuitID int, ignoreDhcpOptionListRequest string, useIgnoreDhcpOptionListRequest bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_ignore_dhcp_option_list_request" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	ignore_dhcp_option_list_request = %q
	use_ignore_dhcp_option_list_request = %t
}
`, ip, matchClient, agentCircuitID, ignoreDhcpOptionListRequest, useIgnoreDhcpOptionListRequest)
}

func testAccFixedaddressIpv4addr(ip, matchClient string, agentCircuitID int) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_ipv4addr" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
}
`, ip, matchClient, agentCircuitID)
}

func testAccFixedaddressLogicFilterRules(ip, matchClient string, agentCircuitID int, logicFilterRuleFilter, logicFilterRuleType string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_logic_filter_rules" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	logic_filter_rules = [{
		filter = %q
		type = %q
	}]
}
`, ip, matchClient, agentCircuitID, logicFilterRuleFilter, logicFilterRuleType)
}

func testAccFixedaddressMac(ip, matchClient string, mac string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_mac" {
	ipv4addr = %q
	match_client = %q
	mac = %q
}
`, ip, matchClient, mac)
}

func testAccFixedaddressMatchClientCircuitID(ip, matchClient string, agentCircuitID int) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_match_client" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
}
`, ip, matchClient, agentCircuitID)
}

func testAccFixedaddressMatchClientRemoteID(ip, matchClient string, agentRemoteID int) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_match_client" {
	ipv4addr = %q
	match_client = %q
	agent_remote_id = %d
}
`, ip, matchClient, agentRemoteID)
}

func testAccFixedaddressMsOptions(ip, matchClient string, agentCircuitID int, msOptionsName, msOptionsNum, msOptionsValue string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_ms_options" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	ms_options = [
		{
			name  = %q
			num = %q
			value = %q
		}
	]
}
`, ip, matchClient, agentCircuitID, msOptionsName, msOptionsNum, msOptionsValue)
}

func testAccFixedaddressMsServer(ip, matchClient string, macAddress string, msServerIpv4Addr, networkView, name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_ms_server" {
	ipv4addr = %q
	match_client = %q
	mac = %q
	ms_server = {
		ipv4addr = %q	
	}
	network_view = %q
	name = %q
}
`, ip, matchClient, macAddress, msServerIpv4Addr, networkView, name)
}

func testAccFixedaddressName(ip, matchClient string, agentCircuitID int, name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_name" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	name = %q
}
`, ip, matchClient, agentCircuitID, name)
}

func testAccFixedaddressNetwork(ip, matchClient string, agentCircuitID int, network string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_network" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	network = %q
}
`, ip, matchClient, agentCircuitID, network)
}

func testAccFixedaddressNetworkView(ip, matchClient string, agentCircuitID int, networkView string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_network_view" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	network_view = %q
}
`, ip, matchClient, agentCircuitID, networkView)
}

func testAccFixedaddressNextserver(ip, matchClient string, agentCircuitID int, nextserver string, useNextServer bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_nextserver" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	nextserver = %q
	use_nextserver = %t
}
`, ip, matchClient, agentCircuitID, nextserver, useNextServer)
}

func testAccFixedaddressOptions(ip, matchClient string, agentCircuitID int, optionsName, optionValue string, useOptions bool, num string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_options" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	options = [ {
		name         = %q
		num 		= %q
		value        = %q
	} ]
	use_options = %t
	
}
`, ip, matchClient, agentCircuitID, optionsName, num, optionValue, useOptions)
}

func testAccFixedaddressPxeLeaseTime(ip, matchClient string, agentCircuitID int, pxeLeaseTime string, usePXELeaseTime bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_pxe_lease_time" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	pxe_lease_time = %q
	use_pxe_lease_time = %t
}
`, ip, matchClient, agentCircuitID, pxeLeaseTime, usePXELeaseTime)
}

func testAccFixedaddressReservedInterface(ip, matchClient string, agentCircuitID int, reservedInterface string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_reserved_interface" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	reserved_interface = %q
}
`, ip, matchClient, agentCircuitID, reservedInterface)
}

func testAccFixedaddressRestartIfNeeded(ip, matchClient string, agentCircuitID int, restartIfNeeded string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_restart_if_needed" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	restart_if_needed = %q
}
`, ip, matchClient, agentCircuitID, restartIfNeeded)
}

func testAccFixedaddressSnmp3Credential(ip, matchClient string, agentCircuitID int, snmp3CredentialUser, snmp3CredentialAuthProtocol, snmp3CredentialAuthPass, snmp3CredentialPrvProtocol, snmp3CredentialPrvPass, snmp3CredentialComment, snmp3CredentialGroup, useSnmp3Credentials string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_snmp3_credential" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	snmp3_credential = {
		user = %q
		authentication_protocol = %q
		authentication_password = %q
		privacy_protocol = %q
		privacy_password = %q
		comment = %q
		credential_group = %q
	}
	use_snmp3_credential = %q
	use_cli_credentials = true
}
`, ip, matchClient, agentCircuitID, snmp3CredentialUser, snmp3CredentialAuthProtocol, snmp3CredentialAuthPass, snmp3CredentialPrvProtocol, snmp3CredentialPrvPass, snmp3CredentialComment, snmp3CredentialGroup, useSnmp3Credentials)
}

func testAccFixedaddressSnmpCredential(ip, matchClient string, agentCircuitID int, snmpCredentialCommStr, snmpCredentialComment, snmpCredentialGroup, useSnmpCredentials string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_snmp_credential" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	snmp_credential = {
		community_string = %q
		comment = %q
		credential_group = %q
	}
	use_snmp_credential = %q
}
`, ip, matchClient, agentCircuitID, snmpCredentialCommStr, snmpCredentialComment, snmpCredentialGroup, useSnmpCredentials)
}

func testAccFixedaddressTemplate(ip, matchClient string, agentCircuitID int, template string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_template" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	template = %q
}
`, ip, matchClient, agentCircuitID, template)
}

func testAccFixedaddressUseBootfile(ip, matchClient string, agentCircuitID int, useBootFile, bootFile string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_use_bootfile" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	use_bootfile = %q
	bootfile = %q
}
`, ip, matchClient, agentCircuitID, useBootFile, bootFile)
}

func testAccFixedaddressUseBootserver(ip, matchClient string, agentCircuitID int, useBootserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_use_bootserver" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	use_bootserver = %q
}
`, ip, matchClient, agentCircuitID, useBootserver)
}

func testAccFixedaddressUseCliCredentials(ip, matchClient string, agentCircuitID int, useCliCredentials string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_use_cli_credentials" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	use_cli_credentials = %q
}
`, ip, matchClient, agentCircuitID, useCliCredentials)
}

func testAccFixedaddressUseDdnsDomainname(ip, matchClient string, agentCircuitID int, useDdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_use_ddns_domainname" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	use_ddns_domainname = %q
}
`, ip, matchClient, agentCircuitID, useDdnsDomainname)
}

func testAccFixedaddressUseDenyBootp(ip, matchClient string, agentCircuitID int, useDenyBootp string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_use_deny_bootp" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	use_deny_bootp = %q
}
`, ip, matchClient, agentCircuitID, useDenyBootp)
}

func testAccFixedaddressUseEnableDdns(ip, matchClient string, agentCircuitID int, useEnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_use_enable_ddns" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	use_enable_ddns = %q
}
`, ip, matchClient, agentCircuitID, useEnableDdns)
}

func testAccFixedaddressUseIgnoreDhcpOptionListRequest(ip, matchClient string, agentCircuitID int, useIgnoreDhcpOptionListRequest string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_use_ignore_dhcp_option_list_request" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	use_ignore_dhcp_option_list_request = %q
}
`, ip, matchClient, agentCircuitID, useIgnoreDhcpOptionListRequest)
}

func testAccFixedaddressUseLogicFilterRules(ip, matchClient string, agentCircuitID int, useLogicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_use_logic_filter_rules" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	use_logic_filter_rules = %q
}
`, ip, matchClient, agentCircuitID, useLogicFilterRules)
}

func testAccFixedaddressUseMsOptions(ip, matchClient string, agentCircuitID int, useMsOptions string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_use_ms_options" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	use_ms_options = %q
}
`, ip, matchClient, agentCircuitID, useMsOptions)
}

func testAccFixedaddressUseNextserver(ip, matchClient string, agentCircuitID int, useNextserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_use_nextserver" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	use_nextserver = %q
}
`, ip, matchClient, agentCircuitID, useNextserver)
}

func testAccFixedaddressUseOptions(ip, matchClient string, agentCircuitID int, useOptions string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_use_options" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	use_options = %q
}
`, ip, matchClient, agentCircuitID, useOptions)
}

func testAccFixedaddressUsePxeLeaseTime(ip, matchClient string, agentCircuitID int, usePxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_use_pxe_lease_time" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	use_pxe_lease_time = %q
}
`, ip, matchClient, agentCircuitID, usePxeLeaseTime)
}

func testAccFixedaddressUseSnmp3Credential(ip, matchClient string, agentCircuitID int, useSnmp3Credential, snmp3CredentialUser, snmp3CredentialAuthProtocol, snmp3CredentialAuthPass, snmp3CredentialPrvProtocol, snmp3CredentialPrvPass, snmp3CredentialComment, snmp3CredentialGroup string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_use_snmp3_credential" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	use_snmp3_credential = %q
	snmp3_credential = {
		user = %q
		authentication_protocol = %q
		authentication_password = %q
		privacy_protocol = %q
		privacy_password = %q
		comment = %q
		credential_group = %q
	}
}
`, ip, matchClient, agentCircuitID, useSnmp3Credential, snmp3CredentialUser, snmp3CredentialAuthProtocol, snmp3CredentialAuthPass, snmp3CredentialPrvProtocol, snmp3CredentialPrvPass, snmp3CredentialComment, snmp3CredentialGroup)
}

func testAccFixedaddressUseSnmpCredential(ip, matchClient string, agentCircuitID int, useSnmpCredential string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixed_address" "test_use_snmp_credential" {
	ipv4addr = %q
	match_client = %q
	agent_circuit_id = %d
	use_snmp_credential = %q
}
`, ip, matchClient, agentCircuitID, useSnmpCredential)
}
