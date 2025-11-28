package dhcp_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForIpv6fixedaddress = "address_type,allow_telnet,cli_credentials,cloud_info,comment,device_description,device_location,device_type,device_vendor,disable,disable_discovery,discover_now_status,discovered_data,domain_name,domain_name_servers,duid,extattrs,ipv6addr,ipv6prefix,ipv6prefix_bits,logic_filter_rules,mac_address,match_client,ms_ad_user_data,name,network,network_view,options,preferred_lifetime,reserved_interface,snmp3_credential,snmp_credential,use_cli_credentials,use_domain_name,use_domain_name_servers,use_logic_filter_rules,use_options,use_preferred_lifetime,use_snmp3_credential,use_snmp_credential,use_valid_lifetime,valid_lifetime"

// TODO: Add tests:
// The following require additional resource/data source objects to be supported.
// - Logic Filter Rules
// - Reserved Interface
// - IPV6 Fixed Address Template
// - SNMP credentials

func TestAccIpv6fixedaddressResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	//macAddr := "00:0c:29:AB:CD:EF"
	//matchClient := "MAC_ADDRESS"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressBasicConfig(ipv6addr, duid, networkView, ipv6Network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", ipv6addr),
					resource.TestCheckResourceAttr(resourceName, "duid", duid),
					resource.TestCheckResourceAttr(resourceName, "network_view", networkView),
					resource.TestCheckResourceAttr(resourceName, "network", ipv6Network),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "address_type", "ADDRESS"),
					resource.TestCheckResourceAttr(resourceName, "match_client", "DUID"),
					resource.TestCheckResourceAttr(resourceName, "allow_telnet", "false"),
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "device_description", ""),
					resource.TestCheckResourceAttr(resourceName, "device_location", ""),
					resource.TestCheckResourceAttr(resourceName, "device_type", ""),
					resource.TestCheckResourceAttr(resourceName, "device_vendor", ""),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "disable_discovery", "false"),
					resource.TestCheckResourceAttr(resourceName, "name", ""),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "27000"),
					resource.TestCheckResourceAttr(resourceName, "use_cli_credentials", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_snmp3_credential", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_snmp_credential", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "false"),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "43200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_ipv6fixedaddress.test"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6fixedaddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6fixedaddressBasicConfig(ipv6addr, duid, networkView, ipv6Network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					testAccCheckIpv6fixedaddressDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccIpv6fixedaddressResource_AddressType(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_address_type"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	ipv6adress1 := "2001:db8:abcd:1231::2"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	ipv6Prefix := "2001:db8:abcd:1231::"
	ipv6Prefix1 := "2001:db8:abcd:1232::"
	//ipv6Prefix2 := "2001:db8:abcd:1231::"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressAddressType(ipv6addr, duid, networkView, ipv6Network, "ADDRESS", "", 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address_type", "ADDRESS"),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", ipv6addr),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressAddressType("", duid, networkView, ipv6Network, "PREFIX", ipv6Prefix, 64),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address_type", "PREFIX"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressAddressType(ipv6adress1, duid, networkView, ipv6Network, "BOTH", ipv6Prefix1, 64),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address_type", "BOTH"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_AllowTelnet(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_allow_telnet"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressAllowTelnet(ipv6addr, duid, networkView, ipv6Network, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_telnet", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressAllowTelnet(ipv6addr, duid, networkView, ipv6Network, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_telnet", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_CliCredentials(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_cli_credentials"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressCliCredentials(ipv6addr, duid, networkView, ipv6Network, "Comment for CLI Credentials", "NIOS_USER", "NIOS_PASSWORD", "SSH", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.0.comment", "Comment for CLI Credentials"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.0.user", "NIOS_USER"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.0.password", "NIOS_PASSWORD"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.0.credential_type", "SSH"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.0.credential_group", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressCliCredentials(ipv6addr, duid, networkView, ipv6Network, "Updated Comment for CLI Credentials", "NIOS_USER", "NIOS_PASSWORD", "TELNET", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.comment", "Updated Comment for CLI Credentials"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.user", "NIOS_USER"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.password", "NIOS_PASSWORD"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.credential_type", "TELNET"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.credential_group", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressCliCredentials(ipv6addr, duid, networkView, ipv6Network, "Updated Comment for CLI Credentials", "NIOS_USER", "NIOS_PASSWORD", "ENABLE_SSH", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.comment", "Updated Comment for CLI Credentials"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.user", "NIOS_USER"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.password", "NIOS_PASSWORD"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.credential_type", "ENABLE_TELNET"),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials.1.credential_group", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressCliCredentials(ipv6addr, duid, networkView, ipv6Network, "Updated Comment for CLI Credentials", "NIOS_USER", "NIOS_PASSWORD", "TELNET", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
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

func TestAccIpv6fixedaddressResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_comment"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressComment(ipv6addr, duid, networkView, ipv6Network, "IPV6 Fixed Address Comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "IPV6 Fixed Address Comment"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressComment(ipv6addr, duid, networkView, ipv6Network, "IPV6 Fixed Address Comment Updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "IPV6 Fixed Address Comment Updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_DeviceDescription(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_device_description"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	deviceDesc1 := acctest.RandomNameWithPrefix("device-description-")
	deviceDesc2 := acctest.RandomNameWithPrefix("device-description-")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressDeviceDescription(ipv6addr, duid, networkView, ipv6Network, deviceDesc1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_description", deviceDesc1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressDeviceDescription(ipv6addr, duid, networkView, ipv6Network, deviceDesc2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_description", deviceDesc2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_DeviceLocation(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_device_location"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	devLocation1 := acctest.RandomNameWithPrefix("device-location-")
	devLocation2 := acctest.RandomNameWithPrefix("device-location-")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressDeviceLocation(ipv6addr, duid, networkView, ipv6Network, devLocation1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_location", devLocation1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressDeviceLocation(ipv6addr, duid, networkView, ipv6Network, devLocation2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_location", devLocation2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_DeviceType(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_device_type"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	devType1 := acctest.RandomNameWithPrefix("device-type-")
	devType2 := acctest.RandomNameWithPrefix("device-type-")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressDeviceType(ipv6addr, duid, networkView, ipv6Network, devType1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_type", devType1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressDeviceType(ipv6addr, duid, networkView, ipv6Network, devType2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_type", devType2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_DeviceVendor(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_device_vendor"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	devVendor1 := acctest.RandomNameWithPrefix("device-vendor-")
	devVendor2 := acctest.RandomNameWithPrefix("device-vendor-")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressDeviceVendor(ipv6addr, duid, networkView, ipv6Network, devVendor1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_vendor", devVendor1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressDeviceVendor(ipv6addr, duid, networkView, ipv6Network, devVendor2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_vendor", devVendor2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_Disable(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_disable"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressDisable(ipv6addr, duid, networkView, ipv6Network, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressDisable(ipv6addr, duid, networkView, ipv6Network, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_DisableDiscovery(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_disable_discovery"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressDisableDiscovery(ipv6addr, duid, networkView, ipv6Network, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_discovery", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressDisableDiscovery(ipv6addr, duid, networkView, ipv6Network, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_discovery", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_DomainName(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_domain_name"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	domainName1 := acctest.RandomName()
	domainName2 := acctest.RandomName()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressDomainName(ipv6addr, duid, networkView, ipv6Network, domainName1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name", domainName1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressDomainName(ipv6addr, duid, networkView, ipv6Network, domainName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name", domainName2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_DomainNameServers(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_domain_name_servers"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	domainNameServers1 := `["2001:4860:4860::8888","2001:4860:4860::8844"]`
	domainNameServers2 := `["2620:fe::9","2001:4860:4860::6844"]`
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressDomainNameServers(ipv6addr, duid, networkView, ipv6Network, domainNameServers1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.0", "2001:4860:4860::8888"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.1", "2001:4860:4860::8844"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressDomainNameServers(ipv6addr, duid, networkView, ipv6Network, domainNameServers2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.0", "2620:fe::9"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.1", "2001:4860:4860::6844"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_Duid(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_duid"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid1 := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	duid2 := "00:01:00:11:11:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressDuid(ipv6addr, duid1, networkView, ipv6Network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "duid", duid1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressDuid(ipv6addr, duid2, networkView, ipv6Network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "duid", duid2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_extattrs"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid1 := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	extAttrs1 := acctest.RandomName()
	extAttrs2 := acctest.RandomName()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressExtAttrs(ipv6addr, duid1, networkView, ipv6Network, map[string]any{"Site": extAttrs1}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrs1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressExtAttrs(ipv6addr, duid1, networkView, ipv6Network, map[string]any{"Site": extAttrs2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrs2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_Ipv6addr(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_ipv6addr"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr1 := "2001:db8:abcd:1231::1"
	ipv6addr2 := "2001:db8:abcd:1231::2"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressIpv6addr(ipv6addr1, duid, networkView, ipv6Network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", ipv6addr1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressIpv6addr(ipv6addr2, duid, networkView, ipv6Network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", ipv6addr2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// TestAccIpv6fixedaddressResource_FuncCall tests the "func_call" attribute functionality
// which allocates IPv6 addresses using next_available_ip. Since func_call attribute can't be
// updated, the comment is updated to demonstrate an update to the resource
func TestAccIpv6fixedaddressResource_FuncCall(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_func_call"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	//ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressFuncCall(duid, networkView, ipv6Network, "ipv6addr", "next_available_ip", "ips", "ipv6network", "create IPV6 Fixed Address using Func call"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "create IPV6 Fixed Address using Func call"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressFuncCall(duid, networkView, ipv6Network, "ipv6addr", "next_available_ip", "ips", "ipv6network", "update IPV6 Fixed Address using Func call"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "update IPV6 Fixed Address using Func call"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_Ipv6prefix(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_ipv6prefix"
	var v dhcp.Ipv6fixedaddress
	ipv6Network1 := "2001:db8:abcd:1231::/64"
	ipv6Prefix1 := "2001:db8:abcd:1231::"
	ipv6Prefix2 := "2001:db8:abcd:1241::"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressIpv6prefix(duid, networkView, ipv6Network1, "PREFIX", ipv6Prefix1, 64),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6prefix", ipv6Prefix1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressIpv6prefix(duid, networkView, ipv6Network1, "PREFIX", ipv6Prefix2, 64),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6prefix", ipv6Prefix2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_Ipv6prefixBits(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_ipv6prefix_bits"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	ipv6Prefix := "2001:db8:abcd:1231::"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressIpv6prefixBits(duid, networkView, ipv6Network, "BOTH", ipv6addr, ipv6Prefix, 64),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6prefix_bits", "64"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressIpv6prefixBits(duid, networkView, ipv6Network, "BOTH", ipv6addr, ipv6Prefix, 65),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6prefix_bits", "65"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_LogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_logic_filter_rules"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	logicFilterRules := []map[string]any{
		{
			"filter": "ipv6_option_filter",
			"type":   "Option",
		},
	}
	logicFilterRulesUpdated := []map[string]any{
		{
			"filter": "ipv6_option_filter1",
			"type":   "Option",
		},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressLogicFilterRules(ipv6addr, duid, networkView, ipv6Network, logicFilterRules),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.filter", "ipv6_option_filter"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.type", "Option"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressLogicFilterRules(ipv6addr, duid, networkView, ipv6Network, logicFilterRulesUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.filter", "ipv6_option_filter1"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.type", "Option"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_MacAddress(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_mac_address"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	macAddr1 := "00:0c:29:ab:cd:ef"
	macAddr2 := "01:2c:39:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressMacAddress(ipv6addr, "MAC_ADDRESS", macAddr1, networkView, ipv6Network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mac_address", macAddr1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressMacAddress(ipv6addr, "MAC_ADDRESS", macAddr2, networkView, ipv6Network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mac_address", macAddr2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_MatchClient(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_match_client"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	macAddr := "00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressMatchClient(ipv6addr, duid, "", networkView, ipv6Network, "DUID"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_client", "DUID"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressMatchClient(ipv6addr, "", macAddr, networkView, ipv6Network, "MAC_ADDRESS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_client", "MAC_ADDRESS"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_name"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	name1 := acctest.RandomNameWithPrefix("ipv6fixedaddress-name-")
	name2 := acctest.RandomNameWithPrefix("ipv6fixedaddress-name-")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressName(ipv6addr, duid, networkView, ipv6Network, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressName(ipv6addr, duid, networkView, ipv6Network, name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_Network(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_network"
	var v dhcp.Ipv6fixedaddress
	ipv6Network1 := "2001:db8:abcd:1231::/64"
	ipv6addr1 := "2001:db8:abcd:1231::1"
	ipv6Network2 := "2001:db8:abcd:1241::/64"
	ipv6addr2 := "2001:db8:abcd:1241::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressNetwork(ipv6addr1, duid, networkView, ipv6Network1, ipv6Network2, "test_ipv6network1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", ipv6Network1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressNetwork(ipv6addr2, duid, networkView, ipv6Network1, ipv6Network2, "test_ipv6network2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", ipv6Network2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_Options(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_options"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	options := []map[string]any{
		{
			"name":  "domain-name",
			"num":   "15",
			"value": "aa.bb.com",
		},
	}
	optionsUpdated := []map[string]any{
		{
			"name":  "domain-name",
			"value": "bb.cc.com",
		},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressOptions(ipv6addr, duid, networkView, ipv6Network, options),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "domain-name"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "aa.bb.com"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressOptions(ipv6addr, duid, networkView, ipv6Network, optionsUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "bb.cc.com"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "domain-name"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_PreferredLifetime(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_preferred_lifetime"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressPreferredLifetime(ipv6addr, duid, networkView, ipv6Network, 6200),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "6200"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressPreferredLifetime(ipv6addr, duid, networkView, ipv6Network, 4800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "4800"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_ReservedInterface(t *testing.T) {
	t.Skip("Skipping test as reserved_interface is not implemented yet")
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_reserved_interface"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	reservedInterface1 := acctest.RandomNameWithPrefix("reserved-interface-")
	reservedInterface2 := acctest.RandomNameWithPrefix("reserved-interface-")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressReservedInterface(ipv6addr, duid, networkView, ipv6Network, reservedInterface1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "reserved_interface", reservedInterface1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressReservedInterface(ipv6addr, duid, networkView, ipv6Network, reservedInterface2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "reserved_interface", reservedInterface2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_Snmp3Credential(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_snmp3_credential"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressSnmp3Credential(ipv6addr, duid, networkView, ipv6Network, "snmp", "MD5", "snmp1234", "3DES", "snmp1234", "SNMP3 Credential Comment", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.user", "snmp"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.authentication_protocol", "MD5"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.authentication_password", "snmp1234"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.privacy_protocol", "3DES"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.privacy_password", "snmp1234"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.comment", "SNMP3 Credential Comment"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.credential_group", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressSnmp3Credential(ipv6addr, duid, networkView, ipv6Network, "SNMP3_USER_UPDATE", "SHA-224", "AUTH_PASSWORD", "AES-256", "PRIVACY_PASSWORD", "SNMP3 Credential Comment Updated", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.user", "SNMP3_USER_UPDATE"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.authentication_protocol", "SHA-224"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.authentication_password", "AUTH_PASSWORD"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.privacy_protocol", "AES-256"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.privacy_password", "PRIVACY_PASSWORD"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.comment", "SNMP3 Credential Comment Updated"),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential.credential_group", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_SnmpCredential(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_snmp_credential"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressSnmpCredential(ipv6addr, duid, networkView, ipv6Network, "COMMUNITY_STRING", "SNMP Credential Comment", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "snmp_credential.community_string", "COMMUNITY_STRING"),
					resource.TestCheckResourceAttr(resourceName, "snmp_credential.comment", "SNMP Credential Comment"),
					resource.TestCheckResourceAttr(resourceName, "snmp_credential.credential_group", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressSnmpCredential(ipv6addr, duid, networkView, ipv6Network, "COMMUNITY_STRING_UPDATED", "SNMP Credential Comment Updated", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "snmp_credential.community_string", "COMMUNITY_STRING_UPDATED"),
					resource.TestCheckResourceAttr(resourceName, "snmp_credential.comment", "SNMP Credential Comment Updated"),
					resource.TestCheckResourceAttr(resourceName, "snmp_credential.credential_group", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_UseCliCredentials(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_use_cli_credentials"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressUseCliCredentials(ipv6addr, duid, networkView, ipv6Network, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_cli_credentials", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressUseCliCredentials(ipv6addr, duid, networkView, ipv6Network, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_cli_credentials", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_UseDomainName(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_use_domain_name"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	domainName := acctest.RandomName()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressUseDomainName(ipv6addr, duid, networkView, ipv6Network, true, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressUseDomainName(ipv6addr, duid, networkView, ipv6Network, false, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_UseDomainNameServers(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_use_domain_name_servers"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressUseDomainNameServers(ipv6addr, duid, networkView, ipv6Network, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressUseDomainNameServers(ipv6addr, duid, networkView, ipv6Network, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_UseLogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_use_logic_filter_rules"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressUseLogicFilterRules(ipv6addr, duid, networkView, ipv6Network, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressUseLogicFilterRules(ipv6addr, duid, networkView, ipv6Network, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_UseOptions(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_use_options"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressUseOptions(ipv6addr, duid, networkView, ipv6Network, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressUseOptions(ipv6addr, duid, networkView, ipv6Network, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_UsePreferredLifetime(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_use_preferred_lifetime"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressUsePreferredLifetime(ipv6addr, duid, networkView, ipv6Network, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressUsePreferredLifetime(ipv6addr, duid, networkView, ipv6Network, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_UseSnmp3Credential(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_use_snmp3_credential"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressUseSnmp3Credential(ipv6addr, duid, networkView, ipv6Network, true, true, "SNMP3_USER", "MD5", "AUTH_PASSWORD", "3DES", "PRIVACY_PASSWORD", "SNMP3 Credential Comment", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp3_credential", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressUseSnmp3Credential(ipv6addr, duid, networkView, ipv6Network, false, false, "", "", "", "", "", "", ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp3_credential", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_UseSnmpCredential(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_use_snmp_credential"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressUseSnmpCredential(ipv6addr, duid, networkView, ipv6Network, true, "COMMUNITY_STRING", "SNMP Credential Comment", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp_credential", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressUseSnmpCredential(ipv6addr, duid, networkView, ipv6Network, false, "", "", ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp_credential", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_UseValidLifetime(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_use_valid_lifetime"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressUseValidLifetime(ipv6addr, duid, networkView, ipv6Network, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressUseValidLifetime(ipv6addr, duid, networkView, ipv6Network, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddressResource_ValidLifetime(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddress.test_valid_lifetime"
	var v dhcp.Ipv6fixedaddress
	ipv6Network := "2001:db8:abcd:1231::/64"
	ipv6addr := "2001:db8:abcd:1231::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddressValidLifetime(ipv6addr, duid, networkView, ipv6Network, 42800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "42800"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddressValidLifetime(ipv6addr, duid, networkView, ipv6Network, 56000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "56000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckIpv6fixedaddressExists(ctx context.Context, resourceName string, v *dhcp.Ipv6fixedaddress) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			Ipv6fixedaddressAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForIpv6fixedaddress).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetIpv6fixedaddressResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetIpv6fixedaddressResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckIpv6fixedaddressDestroy(ctx context.Context, v *dhcp.Ipv6fixedaddress) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			Ipv6fixedaddressAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForIpv6fixedaddress).
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

func testAccCheckIpv6fixedaddressDisappears(ctx context.Context, v *dhcp.Ipv6fixedaddress) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			Ipv6fixedaddressAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccIpv6fixedaddressBasicConfig(ipv6addr, duid, networkView, ipv6Network string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test" {
 ipv6addr = %q
 duid = %q
 network = nios_ipam_ipv6network.test_ipv6_network.network
 network_view = nios_ipam_network_view.parent_network_view.name
}`, ipv6addr, duid)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressAddressType(ipv6addr, duid, networkView, ipv6Network, addressType string, ipv6Prefix string, ipv6PrefixBits int32) string {
	var ipv6Config string
	switch addressType {
	case "BOTH":
		ipv6Config = fmt.Sprintf(`ipv6prefix = %q, ipv6prefix_bits = %d`, ipv6Prefix, ipv6PrefixBits)
	case "PREFIX":
		ipv6Config = fmt.Sprintf(`ipv6prefix = %q`, ipv6Prefix)
	case "ADDRESS":
		ipv6Config = fmt.Sprintf(`ipv6addr = %q`, ipv6addr)
	}
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_address_type" {
    %s
    duid = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
    address_type = %q
}
`, ipv6Config, duid, addressType)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressAllowTelnet(ipv6addr, duid, networkView, ipv6Network string, allowTelnet bool) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_allow_telnet" {
    ipv6addr = %q
    duid = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
    allow_telnet = %t
}
`, ipv6addr, duid, allowTelnet)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressCliCredentials(ipv6addr, duid, networkView, ipv6Network, cliCredComment, cliCredUser, cliCredPassword, cliCredType, cliCredGroup, useCLICredentials string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_cli_credentials" {
    ipv6addr = %q
    duid = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
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
`, ipv6addr, duid, cliCredComment, cliCredUser, cliCredPassword, cliCredType, cliCredGroup, useCLICredentials)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressComment(ipv6addr, duid, networkView, ipv6Network, comment string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_comment" {
    ipv6addr = %q
    duid = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
    comment = %q
}
`, ipv6addr, duid, comment)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressDeviceDescription(ipv6addr, duid, networkView, ipv6Network, deviceDescription string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_device_description" {
    ipv6addr = %q
    duid = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
    device_description = %q
}
`, ipv6addr, duid, deviceDescription)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressDeviceLocation(ipv6addr, duid, networkView, ipv6Network, deviceLocation string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_device_location" {
    ipv6addr = %q
    duid = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
    device_location = %q
}
`, ipv6addr, duid, deviceLocation)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressDeviceType(ipv6addr, duid, networkView, ipv6Network, deviceType string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_device_type" {
    ipv6addr = %q
    duid = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
    device_type = %q
}
`, ipv6addr, duid, deviceType)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressDeviceVendor(ipv6addr, duid, networkView, ipv6Network, deviceVendor string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_device_vendor" {
    ipv6addr = %q
    duid = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
    device_vendor = %q
}
`, ipv6addr, duid, deviceVendor)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressDisable(ipv6addr, duid, networkView, ipv6Network string, disable bool) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_disable" {
    ipv6addr = %q
    duid = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
    disable = %t
}
`, ipv6addr, duid, disable)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressDisableDiscovery(ipv6addr, duid, networkView, ipv6Network string, disableDiscovery bool) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_disable_discovery" {
    ipv6addr = %q
    duid = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
    disable_discovery = %t
}
`, ipv6addr, duid, disableDiscovery)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressDomainName(ipv6addr, duid, networkView, ipv6Network, domainName string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_domain_name" {
    ipv6addr = %q
    duid = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
    domain_name = %q
    use_domain_name = true
}
`, ipv6addr, duid, domainName)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressDomainNameServers(ipv6addr, duid, networkView, ipv6Network, domainNameServers string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_domain_name_servers" {
    ipv6addr = %q
    duid = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
    domain_name_servers = %s
    use_domain_name_servers = true
}
`, ipv6addr, duid, domainNameServers)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressDuid(ipv6addr, duid, networkView, ipv6Network string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_duid" {
    ipv6addr = %q
    duid = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressExtAttrs(ipv6addr, duid, networkView, ipv6Network string, extAttrs map[string]any) string {
	extAttrsStr := utils.ConvertMapToHCL(extAttrs)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_extattrs" {
    ipv6addr = %q
    duid = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
    extattrs = %s
}
`, ipv6addr, duid, extAttrsStr)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressIpv6addr(ipv6addr, duid, networkView, ipv6Network string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_ipv6addr" {
    ipv6addr = %q
    duid = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressFuncCall(duid, networkView, ipv6Network, attributeName, objFunc, resultField, object, comment string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_func_call" {
    duid = %q
    comment = %q
    func_call = {
		"attribute_name" = %q
		"object_function" = %q
		"result_field" = %q
		"object" = %q
		"object_parameters" = {
			"network" = nios_ipam_ipv6network.test_ipv6_network.network
			"network_view" = nios_ipam_network_view.parent_network_view.name
		}
	}
}
`, duid, comment, attributeName, objFunc, resultField, object)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressIpv6prefix(duid, networkView, ipv6Network, addressType, ipv6prefix string, ipv6PrefixBits int) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_ipv6prefix" {
    duid = %q
    address_type = %q
    ipv6prefix = %q
    ipv6prefix_bits = %d
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, duid, addressType, ipv6prefix, ipv6PrefixBits)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressIpv6prefixBits(duid, networkView, ipv6Network, addressType, ipv6addr, ipv6prefix string, ipv6prefixBits int32) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_ipv6prefix_bits" {
    duid = %q
    address_type = %q
    ipv6addr = %q
    ipv6prefix = %q
    ipv6prefix_bits = %d
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, duid, addressType, ipv6addr, ipv6prefix, ipv6prefixBits)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressLogicFilterRules(ipv6addr, duid, networkView, ipv6Network string, logicFilterRules []map[string]any) string {
	logicFilterRulesStr := convertSliceOfMapsToString(logicFilterRules)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_logic_filter_rules" {
    ipv6addr = %q
    duid = %q
    logic_filter_rules = %s
    use_logic_filter_rules = true
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid, logicFilterRulesStr)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressMacAddress(ipv6addr, matchClient, macAddress, networkView, ipv6Network string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_mac_address" {
    ipv6addr = %q
    match_client = %q
    mac_address = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, matchClient, macAddress)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressMatchClient(ipv6addr, duid, macAddress, networkView, ipv6Network, matchClient string) string {
	var extraConfig string
	switch matchClient {
	case "DUID":
		extraConfig = fmt.Sprintf(`duid = %q`, duid)
	case "MAC_ADDRESS":
		extraConfig = fmt.Sprintf(`mac_address = %q`, macAddress)
	}
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_match_client" {
    ipv6addr = %q
    %s
    match_client = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, extraConfig, matchClient)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressMsAdUserData(ipv6addr, duid, networkView, ipv6Network string, msAdUserData string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_ms_ad_user_data" {
    ipv6addr = %q
    duid = %q
    ms_ad_user_data = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid, msAdUserData)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressName(ipv6addr, duid, networkView, ipv6Network, name string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_name" {
    ipv6addr = %q
    duid = %q
    name = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid, name)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressNetwork(ipv6addr, duid, networkView, ipv6Network1, ipv6Network2, networkName string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_network" {
    ipv6addr = %q
    duid = %q
    network = nios_ipam_ipv6network.%s.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid, networkName)
	return strings.Join([]string{testAccBaseNetworkViewCreateTwoNetworks(networkView, ipv6Network1, ipv6Network2), config}, "")
}

func testAccIpv6fixedaddressOptions(ipv6addr, duid, networkView, ipv6Network string, options []map[string]any) string {
	optionsStr := convertSliceOfMapsToString(options)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_options" {
    ipv6addr = %q
    duid = %q
    options = %s
    use_options = true
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid, optionsStr)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressPreferredLifetime(ipv6addr, duid, networkView, ipv6Network string, preferredLifetime int) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_preferred_lifetime" {
    ipv6addr = %q
    duid = %q
    preferred_lifetime = %d
    use_preferred_lifetime = true
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid, preferredLifetime)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressReservedInterface(ipv6addr, duid, networkView, ipv6Network string, reservedInterface string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_reserved_interface" {
    ipv6addr = %q
    duid = %q
    reserved_interface = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid, reservedInterface)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressRestartIfNeeded(ipv6addr, duid, networkView, ipv6Network string, restartIfNeeded bool) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_restart_if_needed" {
    ipv6addr = %q
    duid = %q
    restart_if_needed = %t
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid, restartIfNeeded)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressSnmp3Credential(ipv6addr, duid, networkView, ipv6Network, snmp3CredentialUser, snmp3CredentialAuthProtocol, snmp3CredentialAuthPass, snmp3CredentialPrvProtocol, snmp3CredentialPrvPass, snmp3CredentialComment, snmp3CredentialGroup, useSnmp3Credentials string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_snmp3_credential" {
    ipv6addr = %q
    duid = %q
    snmp3_credential = {
		user = %q
		authentication_protocol = %q
		authentication_password = %q
		privacy_protocol = %q
		privacy_password = %q
		comment = %q
		credential_group = %q
	}
	use_snmp3_credential = true
	use_cli_credentials = true
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid, snmp3CredentialUser, snmp3CredentialAuthProtocol, snmp3CredentialAuthPass, snmp3CredentialPrvProtocol, snmp3CredentialPrvPass, snmp3CredentialComment, snmp3CredentialGroup)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressSnmpCredential(ipv6addr, duid, networkView, ipv6Network, snmpCredentialCommStr, snmpCredentialComment, snmpCredentialGroup, useSnmpCredentials string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_snmp_credential" {
    ipv6addr = %q
    duid = %q
    snmp_credential = {
		community_string = %q
		comment = %q
		credential_group = %q
	}
	use_snmp_credential = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid, snmpCredentialCommStr, snmpCredentialComment, snmpCredentialGroup, useSnmpCredentials)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressTemplate(ipv6addr, duid, networkView, ipv6Network, template string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_template" {
    ipv6addr = %q
    duid = %q
    template = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid, template)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressUseCliCredentials(ipv6addr, duid, networkView, ipv6Network string, useCliCredentials bool) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_use_cli_credentials" {
    ipv6addr = %q
    duid = %q
    use_cli_credentials = %t
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid, useCliCredentials)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressUseDomainName(ipv6addr, duid, networkView, ipv6Network string, useDomainName bool, domainName string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_use_domain_name" {
    ipv6addr = %q
    duid = %q
    use_domain_name = %t
    domain_name = %q
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid, useDomainName, domainName)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressUseDomainNameServers(ipv6addr, duid, networkView, ipv6Network string, useDomainNameServers bool) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_use_domain_name_servers" {
    ipv6addr = %q
    duid = %q
    use_domain_name_servers = %t
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid, useDomainNameServers)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressUseLogicFilterRules(ipv6addr, duid, networkView, ipv6Network string, useLogicFilterRules bool) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_use_logic_filter_rules" {
    ipv6addr = %q
    duid = %q
    use_logic_filter_rules = %t
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid, useLogicFilterRules)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressUseOptions(ipv6addr, duid, networkView, ipv6Network string, useOptions bool) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_use_options" {
    ipv6addr = %q
    duid = %q
    use_options = %t
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid, useOptions)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressUsePreferredLifetime(ipv6addr, duid, networkView, ipv6Network string, usePreferredLifetime bool) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_use_preferred_lifetime" {
    ipv6addr = %q
    duid = %q
    use_preferred_lifetime = %t
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid, usePreferredLifetime)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressUseSnmp3Credential(ipv6addr, duid, networkView, ipv6Network string, useSnmp3Credential, useCliCredentials bool, snmp3CredentialUser, snmp3CredentialAuthProtocol, snmp3CredentialAuthPass, snmp3CredentialPrvProtocol, snmp3CredentialPrvPass, snmp3CredentialComment, snmp3CredentialGroup string) string {
	var snmp3Config string
	if snmp3CredentialUser != "" && snmp3CredentialAuthProtocol != "" && snmp3CredentialAuthPass != "" && snmp3CredentialPrvProtocol != "" && snmp3CredentialPrvPass != "" && snmp3CredentialComment != "" && snmp3CredentialGroup != "" {
		snmp3Config = fmt.Sprintf(`snmp3_credential = {
		user = %q
		authentication_protocol = %q
		authentication_password = %q
		privacy_protocol = %q
		privacy_password = %q
		comment = %q
		credential_group = %q
	}`,
			snmp3CredentialUser, snmp3CredentialAuthProtocol, snmp3CredentialAuthPass, snmp3CredentialPrvProtocol, snmp3CredentialPrvPass, snmp3CredentialComment, snmp3CredentialGroup)
	}
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_use_snmp3_credential" {
    ipv6addr = %q
    duid = %q
    use_snmp3_credential = %t
    use_cli_credentials = %t
    %s
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid, useSnmp3Credential, useCliCredentials, snmp3Config)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressUseSnmpCredential(ipv6addr, duid, networkView, ipv6Network string, useSnmpCredential bool, snmpCredentialCommStr, snmpCredentialComment, snmpCredentialGroup string) string {
	var snmpConfig string
	if snmpCredentialCommStr != "" && snmpCredentialComment != "" && snmpCredentialGroup != "" {
		snmpConfig = fmt.Sprintf(`snmp_credential = {
		community_string = %q
		comment = %q
		credential_group = %q
	}`,
			snmpCredentialCommStr, snmpCredentialComment, snmpCredentialGroup)
	}
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_use_snmp_credential" {
    ipv6addr = %q
    duid = %q
    use_snmp_credential = %t
    %s
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid, useSnmpCredential, snmpConfig)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressUseValidLifetime(ipv6addr, duid, networkView, ipv6Network string, useValidLifetime bool) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_use_valid_lifetime" {
    ipv6addr = %q
    duid = %q
    use_valid_lifetime = %t
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid, useValidLifetime)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccIpv6fixedaddressValidLifetime(ipv6addr, duid, networkView, ipv6Network string, validLifetime int) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddress" "test_valid_lifetime" {
    ipv6addr = %q
    duid = %q
    valid_lifetime = %d
    use_valid_lifetime = true
    network = nios_ipam_ipv6network.test_ipv6_network.network
    network_view = nios_ipam_network_view.parent_network_view.name
}
`, ipv6addr, duid, validLifetime)
	return strings.Join([]string{testAccBaseNetworkView(networkView, ipv6Network), config}, "")
}

func testAccBaseNetworkViewCreateTwoNetworks(networkView, ipv6Network1, ipv6Network2 string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_ipv6network1" {
  network      = %q
  network_view = nios_ipam_network_view.parent_network_view.name
}

resource "nios_ipam_ipv6network" "test_ipv6network2" {
  network      = %q
  network_view = nios_ipam_network_view.parent_network_view.name
}

resource "nios_ipam_network_view" "parent_network_view" {
  name = %q
}
`, ipv6Network1, ipv6Network2, networkView)
}
