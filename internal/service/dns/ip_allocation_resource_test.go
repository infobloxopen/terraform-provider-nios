package dns_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForIPAllocation = "aliases,allow_telnet,cli_credentials,cloud_info,comment,configure_for_dns,creation_time,ddns_protected,device_description,device_location,device_type,device_vendor,disable,disable_discovery,dns_aliases,dns_name,extattrs,ipv4addrs,ipv6addrs,last_queried,ms_ad_user_data,name,network_view,rrset_order,snmp3_credential,snmp_credential,ttl,use_cli_credentials,use_dns_ea_inheritance,use_snmp3_credential,use_snmp_credential,use_ttl,view,zone"

func TestAccIPAllocationResource_basic(t *testing.T) {
	var resourceName = "nios_ip_allocation.test"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.10",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationBasicConfig(name, "default", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
					resource.TestCheckResourceAttr(resourceName, "ipv4addrs.0.ipv4addr", "192.168.1.10"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "allow_telnet", "false"),
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "configure_for_dns", "true"),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "false"),
					resource.TestCheckResourceAttr(resourceName, "device_description", ""),
					resource.TestCheckResourceAttr(resourceName, "device_location", ""),
					resource.TestCheckResourceAttr(resourceName, "device_type", ""),
					resource.TestCheckResourceAttr(resourceName, "device_vendor", ""),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "disable_discovery", "false"),
					resource.TestCheckResourceAttr(resourceName, "network_view", "default"),
					resource.TestCheckResourceAttr(resourceName, "rrset_order", "cyclic"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_disappears(t *testing.T) {
	resourceName := "nios_ip_allocation.test"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.11",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIPAllocationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIPAllocationBasicConfig(name, "default", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					testAccCheckIPAllocationDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccIPAllocationResource_Aliases(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_aliases"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	alias := acctest.RandomName() + ".example.com"
	aliasUpdate := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.12",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationAliases(name, "default", []string{alias}, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "aliases.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "aliases.0", alias),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationAliases(name, "default", []string{aliasUpdate}, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "aliases.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "aliases.0", aliasUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_AllowTelnet(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_allow_telnet"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.13",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationAllowTelnet(name, "default", "false", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_telnet", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationAllowTelnet(name, "default", "true", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_telnet", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_CliCredentials(t *testing.T) {
	t.Skip("Skipping test as CLI Credential are not set up in the GRID")
	var resourceName = "nios_ip_allocation.test_cli_credentials"
	var v dns.RecordHost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationCliCredentials("CLI_CREDENTIALS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials", "CLI_CREDENTIALS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationCliCredentials("CLI_CREDENTIALS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cli_credentials", "CLI_CREDENTIALS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_CloudInfo(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_cloud_info"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.14",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationCloudInfo(name, "default", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.authority_type", "GM"),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.delegated_scope", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.mgmt_platform", ""),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.owned_by_adaptor", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_Comment(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_comment"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.15",
		},
	}
	comment := "new host record"
	updatedComment := "updated host record"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationComment(name, "default", comment, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationComment(name, "default", updatedComment, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", updatedComment),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_ConfigureForDns(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_configure_for_dns"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.16",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationConfigureForDns(name, "default", "true", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "configure_for_dns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationConfigureForDns(name, "default", "false", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "configure_for_dns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_DdnsProtected(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_ddns_protected"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.17",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationDdnsProtected(name, "default", "false", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationDdnsProtected(name, "default", "true", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_DeviceDescription(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_device_description"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.18",
		},
	}
	deviceDesc := "device description"
	updatedDeviceDesc := "updated device description"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationDeviceDescription(name, "default", deviceDesc, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_description", deviceDesc),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationDeviceDescription(name, "default", updatedDeviceDesc, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_description", updatedDeviceDesc),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_DeviceLocation(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_device_location"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.19",
		},
	}
	deviceLocn := "device location"
	updatedDeviceLocn := "updated device location"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationDeviceLocation(name, "default", deviceLocn, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_location", deviceLocn),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationDeviceLocation(name, "default", updatedDeviceLocn, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_location", updatedDeviceLocn),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_DeviceType(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_device_type"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.20",
		},
	}
	deviceType := "device type"
	updatedDeviceType := "updated device type"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationDeviceType(name, "default", deviceType, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_type", deviceType),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationDeviceType(name, "default", updatedDeviceType, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_type", updatedDeviceType),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_DeviceVendor(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_device_vendor"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.21",
		},
	}
	deviceVendor := "device vendor"
	updatedDeviceVendor := "updated device vendor"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationDeviceVendor(name, "default", deviceVendor, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_vendor", deviceVendor),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationDeviceVendor(name, "default", updatedDeviceVendor, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_vendor", updatedDeviceVendor),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_Disable(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_disable"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.22",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationDisable(name, "default", "false", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationDisable(name, "default", "true", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_DisableDiscovery(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_disable_discovery"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.23",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationDisableDiscovery(name, "default", "true", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_discovery", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationDisableDiscovery(name, "default", "false", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_discovery", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_DnsAliases(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_dns_aliases"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	alias := acctest.RandomName() + ".example.com"
	aliasUpdate := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.24",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationDnsAliases(name, "default", []string{alias}, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "aliases.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "aliases.0", alias),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationDnsAliases(name, "default", []string{aliasUpdate}, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "aliases.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "aliases.0", aliasUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_extattrs"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.26",
		},
	}
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationExtAttrs(name, "default", ipv4addr, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationExtAttrs(name, "default", ipv4addr, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_Ipv4addrs(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_ipv4addrs"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.27",
		},
	}
	updatedIpv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.28",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationIpv4addrs(name, "default", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addrs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ipv4addrs.0.ipv4addr", "192.168.1.27"),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationIpv4addrs(name, "default", updatedIpv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addrs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ipv4addrs.0.ipv4addr", "192.168.1.28"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_Ipv6addrs(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_ipv6addrs"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv6addr := []map[string]any{
		{
			"ipv6addr": "fd00:1234:5678::1",
		},
	}
	updatedIpv6addr := []map[string]any{
		{
			"ipv6addr": "fd00:1234:5678::12",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationIpv6addrs(name, "default", ipv6addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addrs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ipv6addrs.0.ipv6addr", "fd00:1234:5678::1"),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationIpv6addrs(name, "default", updatedIpv6addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addrs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ipv6addrs.0.ipv6addr", "fd00:1234:5678::12"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_Name(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_name"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	updatedName := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.2.10",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationName(name, "default", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationName(updatedName, "default", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_NetworkView(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_network_view"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.2.11",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationNetworkView(name, "default", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_RrsetOrder(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_rrset_order"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.2.12",
		},
	}
	rrsetOrder := "cyclic"
	updatedRrsetOrder := "random"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationRrsetOrder(name, "default", rrsetOrder, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rrset_order", rrsetOrder),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationRrsetOrder(name, "default", updatedRrsetOrder, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rrset_order", updatedRrsetOrder),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_Snmp3Credential(t *testing.T) {
	t.Skip("Skipping test as SNMP3 Credential is not supported yet")
	var resourceName = "nios_ip_allocation.test_snmp3_credential"
	var v dns.RecordHost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationSnmp3Credential("SNMP3_CREDENTIAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential", "SNMP3_CREDENTIAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationSnmp3Credential("SNMP3_CREDENTIAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "snmp3_credential", "SNMP3_CREDENTIAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_SnmpCredential(t *testing.T) {
	t.Skip("Skipping test as SNMP Credential are not set up in the GRID")
	var resourceName = "nios_ip_allocation.test_snmp_credential"
	var v dns.RecordHost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationSnmpCredential("SNMP_CREDENTIAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "snmp_credential", "SNMP_CREDENTIAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationSnmpCredential("SNMP_CREDENTIAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "snmp_credential", "SNMP_CREDENTIAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_Ttl(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_ttl"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.2.13",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationTtl(name, "default", 10, "true", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationTtl(name, "default", 0, "true", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_UseCliCredentials(t *testing.T) {
	t.Skip("Skipping test as CLI Credential are not set up in the GRID")
	var resourceName = "nios_ip_allocation.test_use_cli_credentials"
	var v dns.RecordHost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationUseCliCredentials("USE_CLI_CREDENTIALS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_cli_credentials", "USE_CLI_CREDENTIALS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationUseCliCredentials("USE_CLI_CREDENTIALS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_cli_credentials", "USE_CLI_CREDENTIALS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_UseDnsEaInheritance(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_use_dns_ea_inheritance"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.2.14",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationUseDnsEaInheritance(name, "default", "true", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dns_ea_inheritance", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationUseDnsEaInheritance(name, "default", "false", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dns_ea_inheritance", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_UseSnmp3Credential(t *testing.T) {
	t.Skip("Skipping test as SNMP3 Credential is not supported yet")
	var resourceName = "nios_ip_allocation.test_use_snmp3_credential"
	var v dns.RecordHost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationUseSnmp3Credential("USE_SNMP3_CREDENTIAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp3_credential", "USE_SNMP3_CREDENTIAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationUseSnmp3Credential("USE_SNMP3_CREDENTIAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp3_credential", "USE_SNMP3_CREDENTIAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_UseSnmpCredential(t *testing.T) {
	t.Skip("Skipping test as SNMP Credential are not set up in the GRID")
	var resourceName = "nios_ip_allocation.test_use_snmp_credential"
	var v dns.RecordHost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationUseSnmpCredential("USE_SNMP_CREDENTIAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp_credential", "USE_SNMP_CREDENTIAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationUseSnmpCredential("USE_SNMP_CREDENTIAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp_credential", "USE_SNMP_CREDENTIAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_UseTtl(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_use_ttl"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.2.15",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationUseTtl(name, "default", "true", 10, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationUseTtl(name, "default", "false", 10, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAllocationResource_View(t *testing.T) {
	var resourceName = "nios_ip_allocation.test_view"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.2.16",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAllocationView(name, "default", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccIPAllocationView(name, "default", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckIPAllocationExists(ctx context.Context, resourceName string, v *dns.RecordHost) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			RecordHostAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForIPAllocation).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordHostResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordHostResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckIPAllocationDestroy(ctx context.Context, v *dns.RecordHost) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			RecordHostAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForIPAllocation).
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

func testAccCheckIPAllocationDisappears(ctx context.Context, v *dns.RecordHost) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			RecordHostAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccIPAllocationBasicConfig(name, view string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test" {
	name = %q
	ipv4addrs = %s
	view = %q
}
`, name, ipv4addrHCL, view)
}

func testAccIPAllocationAliases(name, view string, aliases []string, ipv4addr []map[string]any) string {
	aliasesHCL := utils.ConvertStringSliceToHCL(aliases)
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_aliases" {
	name = %q
	view = %q
    aliases = %s
	ipv4addrs = %s
}
`, name, view, aliasesHCL, ipv4addrHCL)
}

func testAccIPAllocationAllowTelnet(name, view, allowTelnet string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_allow_telnet" {
	name = %q
	view = %q
	ipv4addrs = %s
    allow_telnet = %q
}
`, name, view, ipv4addrHCL, allowTelnet)
}

func testAccIPAllocationCliCredentials(cliCredentials string) string {
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_cli_credentials" {
    cli_credentials = %q
}
`, cliCredentials)
}

func testAccIPAllocationCloudInfo(name, view string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_cloud_info" {
	name = %q
	view = %q
	ipv4addrs = %s
}
`, name, view, ipv4addrHCL)
}

func testAccIPAllocationComment(name, view, comment string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_comment" {
	name = %q
	view = %q
	ipv4addrs = %s
    comment = %q
}
`, name, view, ipv4addrHCL, comment)
}

func testAccIPAllocationConfigureForDns(name, view, configureForDns string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_configure_for_dns" {
	name = %q
	view = %q
	ipv4addrs = %s
    configure_for_dns = %q
}
`, name, view, ipv4addrHCL, configureForDns)
}

func testAccIPAllocationDdnsProtected(name, view, ddnsProtected string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_ddns_protected" {
	name = %q
	view = %q
	ipv4addrs = %s
    ddns_protected = %q
}
`, name, view, ipv4addrHCL, ddnsProtected)
}

func testAccIPAllocationDeviceDescription(name, view, deviceDescription string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_device_description" {
	name = %q
	view = %q
	ipv4addrs = %s
    device_description = %q
}
`, name, view, ipv4addrHCL, deviceDescription)
}

func testAccIPAllocationDeviceLocation(name, view, deviceLocation string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_device_location" {
	name = %q
	view = %q
	ipv4addrs = %s
    device_location = %q
}
`, name, view, ipv4addrHCL, deviceLocation)
}

func testAccIPAllocationDeviceType(name, view, deviceType string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_device_type" {
	name = %q
	view = %q
	ipv4addrs = %s
    device_type = %q
}
`, name, view, ipv4addrHCL, deviceType)
}

func testAccIPAllocationDeviceVendor(name, view, deviceVendor string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_device_vendor" {
	name = %q
	view = %q
	ipv4addrs = %s
    device_vendor = %q
}
`, name, view, ipv4addrHCL, deviceVendor)
}

func testAccIPAllocationDisable(name, view, disable string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_disable" {
	name = %q
	view = %q
	ipv4addrs = %s
    disable = %q
}
`, name, view, ipv4addrHCL, disable)
}

func testAccIPAllocationDisableDiscovery(name, view, disableDiscovery string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_disable_discovery" {
	name = %q
	view = %q
	ipv4addrs = %s
    disable_discovery = %q
}
`, name, view, ipv4addrHCL, disableDiscovery)
}

func testAccIPAllocationDnsAliases(name, view string, dnsAliases []string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	dnsAliasesHCL := utils.ConvertStringSliceToHCL(dnsAliases)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_dns_aliases" {
	name = %q
	view = %q
	ipv4addrs = %s
    dns_aliases = %s
}
`, name, view, ipv4addrHCL, dnsAliasesHCL)
}

func testAccIPAllocationExtAttrs(name, view string, ipv4addr []map[string]any, extAttrs map[string]string) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_extattrs" {
	name = %q
	view = %q
	ipv4addrs = %s
    extattrs = %s
}
`, name, view, ipv4addrHCL, extattrsStr)
}

func testAccIPAllocationIpv4addrs(name, view string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_ipv4addrs" {
	name = %q
	view = %q
	ipv4addrs = %s
}
`, name, view, ipv4addrHCL)
}

func testAccIPAllocationIpv6addrs(name, view string, ipv6addrs []map[string]any) string {
	ipv6addrHCL := utils.ConvertSliceOfMapsToHCL(ipv6addrs)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_ipv6addrs" {
	name = %q
	view = %q
    ipv6addrs = %s
}
`, name, view, ipv6addrHCL)
}

func testAccIPAllocationName(name, view string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_name" {
    name = %q
	view = %q
	ipv4addrs = %s
}
`, name, view, ipv4addrHCL)
}

func testAccIPAllocationNetworkView(name, view string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_network_view" {
    name = %q
	view = %q
	ipv4addrs = %s
}
`, name, view, ipv4addrHCL)
}

func testAccIPAllocationRrsetOrder(name, view, rrsetOrder string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_rrset_order" {
	name = %q
	view = %q
	ipv4addrs = %s
    rrset_order = %q
}
`, name, view, ipv4addrHCL, rrsetOrder)
}

func testAccIPAllocationSnmp3Credential(snmp3Credential string) string {
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_snmp3_credential" {
    snmp3_credential = %q
}
`, snmp3Credential)
}

func testAccIPAllocationSnmpCredential(snmpCredential string) string {
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_snmp_credential" {
    snmp_credential = %q
}
`, snmpCredential)
}

func testAccIPAllocationTtl(name, view string, ttl int32, useTtl string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_ttl" {
	name = %q
	view = %q
	ipv4addrs = %s
    ttl = %d
	use_ttl = %q
}
`, name, view, ipv4addrHCL, ttl, useTtl)
}

func testAccIPAllocationUseCliCredentials(useCliCredentials string) string {
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_use_cli_credentials" {
    use_cli_credentials = %q
}
`, useCliCredentials)
}

func testAccIPAllocationUseDnsEaInheritance(name, view, useDnsEaInheritance string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_use_dns_ea_inheritance" {
	name = %q
	view = %q
	ipv4addrs = %s
    use_dns_ea_inheritance = %q
}
`, name, view, ipv4addrHCL, useDnsEaInheritance)
}

func testAccIPAllocationUseSnmp3Credential(useSnmp3Credential string) string {
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_use_snmp3_credential" {
    use_snmp3_credential = %q
}
`, useSnmp3Credential)
}

func testAccIPAllocationUseSnmpCredential(useSnmpCredential string) string {
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_use_snmp_credential" {
    use_snmp_credential = %q
}
`, useSnmpCredential)
}

func testAccIPAllocationUseTtl(name, view, useTtl string, ttl int32, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_use_ttl" {
	name = %q
	view = %q
	ipv4addrs = %s
	ttl = %d
    use_ttl = %q
}
`, name, view, ipv4addrHCL, ttl, useTtl)
}

func testAccIPAllocationView(name, view string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "test_view" {
  	name = %q
	view = %q
	ipv4addrs = %s
}
`, name, view, ipv4addrHCL)
}
