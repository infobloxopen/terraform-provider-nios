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

var readableAttributesForRoaminghost = "address_type,bootfile,bootserver,client_identifier_prepend_zero,comment,ddns_domainname,ddns_hostname,deny_bootp,dhcp_client_identifier,disable,enable_ddns,enable_pxe_lease_time,extattrs,force_roaming_hostname,ignore_dhcp_option_list_request,ipv6_client_hostname,ipv6_ddns_domainname,ipv6_ddns_hostname,ipv6_domain_name,ipv6_domain_name_servers,ipv6_duid,ipv6_enable_ddns,ipv6_force_roaming_hostname,ipv6_mac_address,ipv6_match_option,ipv6_options,mac,match_client,name,network_view,nextserver,options,preferred_lifetime,pxe_lease_time,use_bootfile,use_bootserver,use_ddns_domainname,use_deny_bootp,use_enable_ddns,use_ignore_dhcp_option_list_request,use_ipv6_ddns_domainname,use_ipv6_domain_name,use_ipv6_domain_name_servers,use_ipv6_enable_ddns,use_ipv6_options,use_nextserver,use_options,use_preferred_lifetime,use_pxe_lease_time,use_valid_lifetime,valid_lifetime"

func TestAccRoaminghostResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostBasicConfig(name, mac, "IPV4", "MAC_ADDRESS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "mac", mac),
					resource.TestCheckResourceAttr(resourceName, "address_type", "IPV4"),
					resource.TestCheckResourceAttr(resourceName, "match_client", "MAC_ADDRESS"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "client_identifier_prepend_zero", "false"),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "false"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "force_roaming_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "false"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_force_roaming_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "27000"),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ipv6_ddns_domainname", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ipv6_domain_name", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ipv6_domain_name_servers", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ipv6_enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ipv6_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_roaminghost.test"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRoaminghostDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRoaminghostBasicConfig(name, mac, "IPV4", "MAC_ADDRESS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					testAccCheckRoaminghostDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRoaminghostResource_Import(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostBasicConfig(name, mac, "IPV4", "MAC_ADDRESS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
				),
			},
			// Import with PlanOnly to detect differences
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccRoaminghostImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "ref",
				PlanOnly:                             true,
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccRoaminghostImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_AddressType(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_address_type"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()
	ipv6mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostAddressTypeIPv4(name, mac, "IPV4", "MAC_ADDRESS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address_type", "IPV4"),
				),
			},
			{
				Config: testAccRoaminghostAddressTypeBoth(name, mac, ipv6mac, "BOTH", "MAC_ADDRESS", "V6_MAC_ADDRESS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address_type", "BOTH"),
				),
			},
			{
				Config: testAccRoaminghostAddressTypeIPv6(name, mac, "IPV6", "V6_MAC_ADDRESS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address_type", "IPV6"),
				),
			},
		},
	})
}

func TestAccRoaminghostResource_Bootfile(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_bootfile"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostBootfile(name, mac, "IPV4", "MAC_ADDRESS", "boot.ini"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "boot.ini"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostBootfile(name, mac, "IPV4", "MAC_ADDRESS", "setup.boot"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "setup.boot"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Bootserver(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_bootserver"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostBootserver(name, mac, "IPV4", "MAC_ADDRESS", "example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostBootserver(name, mac, "IPV4", "MAC_ADDRESS", "1.1.1.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "1.1.1.1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_ClientIdentifierPrependZero(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_client_identifier_prepend_zero"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostClientIdentifierPrependZero(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_identifier_prepend_zero", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostClientIdentifierPrependZero(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_identifier_prepend_zero", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_comment"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostComment(name, mac, "IPV4", "MAC_ADDRESS", "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostComment(name, mac, "IPV4", "MAC_ADDRESS", "Updated comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the object"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_DdnsDomainname(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_ddns_domainname"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostDdnsDomainname(name, mac, "IPV4", "MAC_ADDRESS", "example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostDdnsDomainname(name, mac, "IPV4", "MAC_ADDRESS", "example.org"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "example.org"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_DdnsHostname(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_ddns_hostname"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostDdnsHostname(name, mac, "IPV4", "MAC_ADDRESS", "example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_hostname", "example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostDdnsHostname(name, mac, "IPV4", "MAC_ADDRESS", "example.org"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_hostname", "example.org"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_DenyBootp(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_deny_bootp"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostDenyBootp(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostDenyBootp(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_DhcpClientIdentifier(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_dhcp_client_identifier"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostDhcpClientIdentifier(name, "IPV4", "CLIENT_ID", "01aa.aabb.bbcc.cc"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_client_identifier", "01aa.aabb.bbcc.cc"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostDhcpClientIdentifier(name, "IPV4", "CLIENT_ID", "c471.fec5.39d9-Gi0/1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_client_identifier", "c471.fec5.39d9-Gi0/1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Disable(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_disable"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostDisable(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostDisable(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_EnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_enable_ddns"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostEnableDdns(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostEnableDdns(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_EnablePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_enable_pxe_lease_time"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostEnablePxeLeaseTime(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostEnablePxeLeaseTime(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_extattrs"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostExtAttrs(name, mac, "IPV4", "MAC_ADDRESS", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostExtAttrs(name, mac, "IPV4", "MAC_ADDRESS", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_ForceRoamingHostname(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_force_roaming_hostname"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostForceRoamingHostname(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "force_roaming_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostForceRoamingHostname(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "force_roaming_hostname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_IgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_ignore_dhcp_option_list_request"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIgnoreDhcpOptionListRequest(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostIgnoreDhcpOptionListRequest(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Ipv6DdnsDomainname(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_ipv6_ddns_domainname"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIpv6DdnsDomainname(name, mac, "IPV4", "MAC_ADDRESS", "example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_ddns_domainname", "example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostIpv6DdnsDomainname(name, mac, "IPV4", "MAC_ADDRESS", "example.org"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_ddns_domainname", "example.org"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Ipv6DdnsHostname(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_ipv6_ddns_hostname"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIpv6DdnsHostname(name, mac, "IPV4", "MAC_ADDRESS", "example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_ddns_hostname", "example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostIpv6DdnsHostname(name, mac, "IPV4", "MAC_ADDRESS", "example.org"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_ddns_hostname", "example.org"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Ipv6DomainName(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_ipv6_domain_name"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIpv6DomainName(name, mac, "IPV4", "MAC_ADDRESS", "example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_domain_name", "example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostIpv6DomainName(name, mac, "IPV4", "MAC_ADDRESS", "example.org"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_domain_name", "example.org"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Ipv6DomainNameServers(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_ipv6_domain_name_servers"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()
	ipv6DomainNameServersVal := []string{"2001:4860:4860::8888", "2001:4860:4860::9999", "2001:4860:4860::8899"}
	ipv6DomainNameServersValUpdated := []string{"2001:4860:4860::8888", "2001:4860:4860::8844"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIpv6DomainNameServers(name, mac, "IPV4", "MAC_ADDRESS", ipv6DomainNameServersVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_domain_name_servers.0", "2001:4860:4860::8888"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_domain_name_servers.1", "2001:4860:4860::9999"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_domain_name_servers.2", "2001:4860:4860::8899"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostIpv6DomainNameServers(name, mac, "IPV4", "MAC_ADDRESS", ipv6DomainNameServersValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_domain_name_servers.0", "2001:4860:4860::8888"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_domain_name_servers.1", "2001:4860:4860::8844"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Ipv6Duid(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_ipv6_duid"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIpv6Duid(name, "IPV6", "DUID", "03:03:01:01:00:90:7f:97:ad:95"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_duid", "03:03:01:01:00:90:7f:97:ad:95"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostIpv6Duid(name, "IPV6", "DUID", "00:03:03:01:00:90:7f:97:ad:95"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_duid", "00:03:03:01:00:90:7f:97:ad:95"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Ipv6EnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_ipv6_enable_ddns"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIpv6EnableDdns(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostIpv6EnableDdns(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_enable_ddns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Ipv6ForceRoamingHostname(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_ipv6_force_roaming_hostname"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIpv6ForceRoamingHostname(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_force_roaming_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostIpv6ForceRoamingHostname(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_force_roaming_hostname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Ipv6MacAddress(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_ipv6_mac_address"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	ipv6Mac := acctest.RandomMACAddress()
	ipv6MacUpdated := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIpv6MacAddress(name, ipv6Mac, "IPV6", "V6_MAC_ADDRESS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_mac_address", ipv6Mac),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostIpv6MacAddress(name, ipv6MacUpdated, "IPV6", "V6_MAC_ADDRESS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_mac_address", ipv6MacUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Ipv6MatchOption(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_ipv6_match_option"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	ipv6Mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIpv6MatchOptionDuid(name, "IPV6", "DUID", "11:03:00:01:00:90:7f:97:ad:95"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_match_option", "DUID"),
				),
			},
			{
				Config: testAccRoaminghostIpv6MatchOptionMacAddress(name, ipv6Mac, "IPV6", "V6_MAC_ADDRESS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_match_option", "V6_MAC_ADDRESS"),
				),
			},
		},
	})
}

func TestAccRoaminghostResource_Ipv6Options(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_ipv6_options"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()
	ipv6OptionsVal := []map[string]any{
		{
			"name":  "domain-name",
			"num":   "15",
			"value": "example.com",
		},
		{
			"num":          "37",
			"value":        "remote-id",
			"vendor_class": "DHCPv6",
		},
		{
			"name":         "dhcp6.subscriber-id",
			"value":        "subscriber-id",
			"vendor_class": "DHCPv6",
		},
	}
	ipv6OptionsValUpdated := []map[string]any{
		{
			"name":  "domain-name",
			"num":   "15",
			"value": "example.org",
		},
		{
			"num":          "37",
			"value":        "remote-id-updated",
			"vendor_class": "DHCPv6",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIpv6Options(name, mac, "IPV4", "MAC_ADDRESS", ipv6OptionsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_options.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_options.0.name", "domain-name"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_options.0.num", "15"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_options.0.value", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_options.1.num", "37"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_options.1.value", "remote-id"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_options.1.vendor_class", "DHCPv6"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_options.2.name", "dhcp6.subscriber-id"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_options.2.value", "subscriber-id"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_options.2.vendor_class", "DHCPv6"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostIpv6Options(name, mac, "IPV4", "MAC_ADDRESS", ipv6OptionsValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_options.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_options.0.name", "domain-name"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_options.0.num", "15"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_options.0.value", "example.org"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_options.1.num", "37"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_options.1.value", "remote-id-updated"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_options.1.vendor_class", "DHCPv6"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Mac(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_mac"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()
	macUpdated := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostMac(name, mac, "IPV4", "MAC_ADDRESS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mac", mac),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostMac(name, macUpdated, "IPV4", "MAC_ADDRESS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mac", macUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_MatchClient(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_match_client"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostMatchClient(name, "IPV4", "CLIENT_ID", "00:03:00:01:01:90:7f:97:dd:95"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_client", "CLIENT_ID"),
				),
			},
			{
				Config: testAccRoaminghostMatchClientMacAddress(name, mac, "IPV4", "MAC_ADDRESS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_client", "MAC_ADDRESS"),
				),
			},
		},
	})
}

func TestAccRoaminghostResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_name"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	nameUpdated := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostName(name, mac, "IPV4", "MAC_ADDRESS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostName(nameUpdated, mac, "IPV4", "MAC_ADDRESS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_NetworkView(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_network_view"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostNetworkView(name, mac, "IPV4", "MAC_ADDRESS", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Nextserver(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_nextserver"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostNextserver(name, mac, "IPV4", "MAC_ADDRESS", "1.1.1.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", "1.1.1.1"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostNextserver(name, mac, "IPV4", "MAC_ADDRESS", "example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", "example.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Options(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_options"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()
	optionsVal := []map[string]any{
		{
			"name":  "time-offset",
			"num":   2,
			"value": "50",
		},
		{
			"name":  "subnet-mask",
			"value": "1.1.1.1",
		},
	}
	optionsValUpdated := []map[string]any{

		{
			"name":  "subnet-mask",
			"value": "1.1.1.1",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostOptions(name, mac, "IPV4", "MAC_ADDRESS", optionsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "time-offset"),
					resource.TestCheckResourceAttr(resourceName, "options.0.num", "2"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "50"),
					resource.TestCheckResourceAttr(resourceName, "options.1.name", "subnet-mask"),
					resource.TestCheckResourceAttr(resourceName, "options.1.value", "1.1.1.1"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostOptions(name, mac, "IPV4", "MAC_ADDRESS", optionsValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "subnet-mask"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "1.1.1.1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_PreferredLifetime(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_preferred_lifetime"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostPreferredLifetime(name, mac, "IPV4", "MAC_ADDRESS", "43000"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "43000"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostPreferredLifetime(name, mac, "IPV4", "MAC_ADDRESS", "45000"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "45000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccRoaminghostResource_PxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_pxe_lease_time"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostPxeLeaseTime(name, mac, "IPV4", "MAC_ADDRESS", "1000"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "1000"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostPxeLeaseTime(name, mac, "IPV4", "MAC_ADDRESS", "2000"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "2000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_UseBootfile(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_use_bootfile"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseBootfile(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseBootfile(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_UseBootserver(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_use_bootserver"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseBootserver(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseBootserver(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_UseDdnsDomainname(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_use_ddns_domainname"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseDdnsDomainname(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseDdnsDomainname(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_UseDenyBootp(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_use_deny_bootp"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseDenyBootp(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseDenyBootp(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_UseEnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_use_enable_ddns"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseEnableDdns(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseEnableDdns(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_UseIgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_use_ignore_dhcp_option_list_request"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseIgnoreDhcpOptionListRequest(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseIgnoreDhcpOptionListRequest(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_UseIpv6DdnsDomainname(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_use_ipv6_ddns_domainname"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseIpv6DdnsDomainname(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipv6_ddns_domainname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseIpv6DdnsDomainname(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipv6_ddns_domainname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_UseIpv6DomainName(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_use_ipv6_domain_name"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseIpv6DomainName(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipv6_domain_name", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseIpv6DomainName(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipv6_domain_name", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_UseIpv6DomainNameServers(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_use_ipv6_domain_name_servers"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseIpv6DomainNameServers(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipv6_domain_name_servers", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseIpv6DomainNameServers(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipv6_domain_name_servers", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_UseIpv6EnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_use_ipv6_enable_ddns"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseIpv6EnableDdns(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipv6_enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseIpv6EnableDdns(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipv6_enable_ddns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_UseIpv6Options(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_use_ipv6_options"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseIpv6Options(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipv6_options", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseIpv6Options(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipv6_options", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_UseNextserver(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_use_nextserver"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseNextserver(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseNextserver(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_UseOptions(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_use_options"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseOptions(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseOptions(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_UsePreferredLifetime(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_use_preferred_lifetime"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUsePreferredLifetime(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUsePreferredLifetime(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_UsePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_use_pxe_lease_time"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUsePxeLeaseTime(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUsePxeLeaseTime(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_UseValidLifetime(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_use_valid_lifetime"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseValidLifetime(name, mac, "IPV4", "MAC_ADDRESS", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseValidLifetime(name, mac, "IPV4", "MAC_ADDRESS", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_ValidLifetime(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_valid_lifetime"
	var v dhcp.Roaminghost
	name := acctest.RandomNameWithPrefix("roaminghost")
	mac := acctest.RandomMACAddress()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostValidLifetime(name, mac, "IPV4", "MAC_ADDRESS", "4000"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "4000"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostValidLifetime(name, mac, "IPV4", "MAC_ADDRESS", "4500"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "4500"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRoaminghostExists(ctx context.Context, resourceName string, v *dhcp.Roaminghost) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			RoaminghostAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRoaminghost).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRoaminghostResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRoaminghostResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRoaminghostDestroy(ctx context.Context, v *dhcp.Roaminghost) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			RoaminghostAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRoaminghost).
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

func testAccCheckRoaminghostDisappears(ctx context.Context, v *dhcp.Roaminghost) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			RoaminghostAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRoaminghostImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccRoaminghostBasicConfig(name string, mac string, addressType string, matchClient string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test" {
    name = %q
	mac = %q
	address_type = %q
	match_client = %q
}
`, name, mac, addressType, matchClient)
}

func testAccRoaminghostAddressTypeIPv4(name string, mac string, addressType string, matchClient string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_address_type" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
}
`, name, mac, addressType, matchClient)
}

func testAccRoaminghostAddressTypeIPv6(name string, mac string, addressType string, ipv6MatchOption string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_address_type" {
    name = %q
    ipv6_mac_address = %q
    address_type = %q
    ipv6_match_option = %q
}
`, name, mac, addressType, ipv6MatchOption)
}

func testAccRoaminghostAddressTypeBoth(name string, mac string, ipv6mac string, addressType string, matchClient string, ipv6MatchOption string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_address_type" {
    name = %q
    mac = %q
    ipv6_mac_address = %q
    address_type = %q
    match_client = %q
	ipv6_match_option = %q
}
`, name, mac, ipv6mac, addressType, matchClient, ipv6MatchOption)
}

func testAccRoaminghostBootfile(name string, mac string, addressType string, matchClient string, bootfile string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_bootfile" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    bootfile = %q
    use_bootfile = true
}
`, name, mac, addressType, matchClient, bootfile)
}

func testAccRoaminghostBootserver(name string, mac string, addressType string, matchClient string, bootserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_bootserver" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    bootserver = %q
    use_bootserver = true
}
`, name, mac, addressType, matchClient, bootserver)
}

func testAccRoaminghostClientIdentifierPrependZero(name string, mac string, addressType string, matchClient string, clientIdentifierPrependZero string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_client_identifier_prepend_zero" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    client_identifier_prepend_zero = %q
}
`, name, mac, addressType, matchClient, clientIdentifierPrependZero)
}

func testAccRoaminghostComment(name string, mac string, addressType string, matchClient string, comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_comment" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    comment = %q
}
`, name, mac, addressType, matchClient, comment)
}

func testAccRoaminghostDdnsDomainname(name string, mac string, addressType string, matchClient string, ddnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ddns_domainname" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    ddns_domainname = %q
    use_ddns_domainname = true
}
`, name, mac, addressType, matchClient, ddnsDomainname)
}

func testAccRoaminghostDdnsHostname(name string, mac string, addressType string, matchClient string, ddnsHostname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ddns_hostname" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    ddns_hostname = %q
}
`, name, mac, addressType, matchClient, ddnsHostname)
}

func testAccRoaminghostDenyBootp(name string, mac string, addressType string, matchClient string, denyBootp string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_deny_bootp" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    deny_bootp = %q
    use_deny_bootp = true
}
`, name, mac, addressType, matchClient, denyBootp)
}

func testAccRoaminghostDhcpClientIdentifier(name string, addressType string, matchClient string, dhcpClientIdentifier string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_dhcp_client_identifier" {
    name = %q
    address_type = %q
    match_client = %q
    dhcp_client_identifier = %q
}
`, name, addressType, matchClient, dhcpClientIdentifier)
}

func testAccRoaminghostDisable(name string, mac string, addressType string, matchClient string, disable string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_disable" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    disable = %q
}
`, name, mac, addressType, matchClient, disable)
}

func testAccRoaminghostEnableDdns(name string, mac string, addressType string, matchClient string, enableDdns string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_enable_ddns" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    enable_ddns = %q
    use_enable_ddns = true
}
`, name, mac, addressType, matchClient, enableDdns)
}

func testAccRoaminghostEnablePxeLeaseTime(name string, mac string, addressType string, matchClient string, enablePxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_enable_pxe_lease_time" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    enable_pxe_lease_time = %q
	pxe_lease_time = 1200
	use_pxe_lease_time = true
}
`, name, mac, addressType, matchClient, enablePxeLeaseTime)
}

func testAccRoaminghostExtAttrs(name string, mac string, addressType string, matchClient string, extAttrs map[string]string) string {
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_extattrs" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    extattrs = %s
}
`, name, mac, addressType, matchClient, extAttrsStr)
}

func testAccRoaminghostForceRoamingHostname(name string, mac string, addressType string, matchClient string, forceRoamingHostname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_force_roaming_hostname" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    force_roaming_hostname = %q
}
`, name, mac, addressType, matchClient, forceRoamingHostname)
}

func testAccRoaminghostIgnoreDhcpOptionListRequest(name string, mac string, addressType string, matchClient string, ignoreDhcpOptionListRequest string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ignore_dhcp_option_list_request" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    ignore_dhcp_option_list_request = %q
    use_ignore_dhcp_option_list_request = true
}
`, name, mac, addressType, matchClient, ignoreDhcpOptionListRequest)
}

func testAccRoaminghostIpv6DdnsDomainname(name string, mac string, addressType string, matchClient string, ipv6DdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_ddns_domainname" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    ipv6_ddns_domainname = %q
    use_ipv6_ddns_domainname = true
}
`, name, mac, addressType, matchClient, ipv6DdnsDomainname)
}

func testAccRoaminghostIpv6DdnsHostname(name string, mac string, addressType string, matchClient string, ipv6DdnsHostname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_ddns_hostname" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    ipv6_ddns_hostname = %q
}
`, name, mac, addressType, matchClient, ipv6DdnsHostname)
}

func testAccRoaminghostIpv6DomainName(name string, mac string, addressType string, matchClient string, ipv6DomainName string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_domain_name" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    ipv6_domain_name = %q
    use_ipv6_domain_name = true
}
`, name, mac, addressType, matchClient, ipv6DomainName)
}

func testAccRoaminghostIpv6DomainNameServers(name string, mac string, addressType string, matchClient string, ipv6DomainNameServers []string) string {
	ipv6DomainNameServersStr := utils.ConvertStringSliceToHCL(ipv6DomainNameServers)
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_domain_name_servers" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    ipv6_domain_name_servers = %s
    use_ipv6_domain_name_servers = true
}
`, name, mac, addressType, matchClient, ipv6DomainNameServersStr)
}

func testAccRoaminghostIpv6Duid(name string, addressType string, ipv6MatchOption string, ipv6Duid string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_duid" {
    name = %q
    address_type = %q
    ipv6_match_option = %q
    ipv6_duid = %q
}
`, name, addressType, ipv6MatchOption, ipv6Duid)
}

func testAccRoaminghostIpv6EnableDdns(name string, mac string, addressType string, matchClient string, ipv6EnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_enable_ddns" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    ipv6_enable_ddns = %q
    use_ipv6_enable_ddns = true
}
`, name, mac, addressType, matchClient, ipv6EnableDdns)
}

func testAccRoaminghostIpv6ForceRoamingHostname(name string, mac string, addressType string, matchClient string, ipv6ForceRoamingHostname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_force_roaming_hostname" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    ipv6_force_roaming_hostname = %q
}
`, name, mac, addressType, matchClient, ipv6ForceRoamingHostname)
}

func testAccRoaminghostIpv6MacAddress(name string, ipv6mac string, addressType string, ipv6MatchOption string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_mac_address" {
    name = %q
    ipv6_mac_address = %q
    address_type = %q
	ipv6_match_option = %q
}
`, name, ipv6mac, addressType, ipv6MatchOption)
}

func testAccRoaminghostIpv6MatchOptionMacAddress(name string, mac string, addressType string, ipv6MatchOption string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_match_option" {
    name = %q
    ipv6_mac_address = %q
    address_type = %q
    ipv6_match_option = %q
}
`, name, mac, addressType, ipv6MatchOption)
}

func testAccRoaminghostIpv6MatchOptionDuid(name string, addressType string, ipv6MatchOption string, duid string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_match_option" {
    name = %q
    address_type = %q
    ipv6_match_option = %q
	ipv6_duid = %q
}
`, name, addressType, ipv6MatchOption, duid)
}

func testAccRoaminghostIpv6Options(name string, mac string, addressType string, matchClient string, ipv6Options []map[string]any) string {
	ipv6OptionsStr := utils.ConvertSliceOfMapsToHCL(ipv6Options)
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_options" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    ipv6_options = %s
    use_ipv6_options = true
}
`, name, mac, addressType, matchClient, ipv6OptionsStr)
}

func testAccRoaminghostMac(name string, mac string, addressType string, matchClient string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_mac" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
}
`, name, mac, addressType, matchClient)
}

func testAccRoaminghostMatchClient(name string, addressType string, matchClient string, dhcpClientIdentifier string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_match_client" {
    name = %q
    address_type = %q
    match_client = %q
	dhcp_client_identifier = %q
}
`, name, addressType, matchClient, dhcpClientIdentifier)
}

func testAccRoaminghostMatchClientMacAddress(name string, mac string, addressType string, matchClient string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_match_client" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
}
`, name, mac, addressType, matchClient)
}

func testAccRoaminghostName(name string, mac string, addressType string, matchClient string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_name" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
}
`, name, mac, addressType, matchClient)
}

func testAccRoaminghostNetworkView(name string, mac string, addressType string, matchClient string, networkView string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_network_view" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    network_view = %q
}
`, name, mac, addressType, matchClient, networkView)
}

func testAccRoaminghostNextserver(name string, mac string, addressType string, matchClient string, nextserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_nextserver" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    nextserver = %q
    use_nextserver = true
}
`, name, mac, addressType, matchClient, nextserver)
}

func testAccRoaminghostOptions(name string, mac string, addressType string, matchClient string, options []map[string]any) string {
	optionsStr := utils.ConvertSliceOfMapsToHCL(options)
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_options" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    options = %s
    use_options = true
}
`, name, mac, addressType, matchClient, optionsStr)
}

func testAccRoaminghostPreferredLifetime(name string, mac string, addressType string, matchClient string, preferredLifetime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_preferred_lifetime" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    preferred_lifetime = %q
    use_preferred_lifetime = true
}
`, name, mac, addressType, matchClient, preferredLifetime)
}

func testAccRoaminghostPxeLeaseTime(name string, mac string, addressType string, matchClient string, pxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_pxe_lease_time" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    pxe_lease_time = %q
    use_pxe_lease_time = true
}
`, name, mac, addressType, matchClient, pxeLeaseTime)
}

func testAccRoaminghostUseBootfile(name string, mac string, addressType string, matchClient string, useBootfile string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_bootfile" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    use_bootfile = %q
}
`, name, mac, addressType, matchClient, useBootfile)
}

func testAccRoaminghostUseBootserver(name string, mac string, addressType string, matchClient string, useBootserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_bootserver" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    use_bootserver = %q
}
`, name, mac, addressType, matchClient, useBootserver)
}

func testAccRoaminghostUseDdnsDomainname(name string, mac string, addressType string, matchClient string, useDdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_ddns_domainname" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    use_ddns_domainname = %q
}
`, name, mac, addressType, matchClient, useDdnsDomainname)
}

func testAccRoaminghostUseDenyBootp(name string, mac string, addressType string, matchClient string, useDenyBootp string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_deny_bootp" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    use_deny_bootp = %q
}
`, name, mac, addressType, matchClient, useDenyBootp)
}

func testAccRoaminghostUseEnableDdns(name string, mac string, addressType string, matchClient string, useEnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_enable_ddns" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    use_enable_ddns = %q
}
`, name, mac, addressType, matchClient, useEnableDdns)
}

func testAccRoaminghostUseIgnoreDhcpOptionListRequest(name string, mac string, addressType string, matchClient string, useIgnoreDhcpOptionListRequest string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_ignore_dhcp_option_list_request" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    use_ignore_dhcp_option_list_request = %q
}
`, name, mac, addressType, matchClient, useIgnoreDhcpOptionListRequest)
}

func testAccRoaminghostUseIpv6DdnsDomainname(name string, mac string, addressType string, matchClient string, useIpv6DdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_ipv6_ddns_domainname" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    use_ipv6_ddns_domainname = %q
}
`, name, mac, addressType, matchClient, useIpv6DdnsDomainname)
}

func testAccRoaminghostUseIpv6DomainName(name string, mac string, addressType string, matchClient string, useIpv6DomainName string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_ipv6_domain_name" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
	ipv6_domain_name = "example.com"
    use_ipv6_domain_name = %q
}
`, name, mac, addressType, matchClient, useIpv6DomainName)
}

func testAccRoaminghostUseIpv6DomainNameServers(name string, mac string, addressType string, matchClient string, useIpv6DomainNameServers string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_ipv6_domain_name_servers" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    use_ipv6_domain_name_servers = %q
}
`, name, mac, addressType, matchClient, useIpv6DomainNameServers)
}

func testAccRoaminghostUseIpv6EnableDdns(name string, mac string, addressType string, matchClient string, useIpv6EnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_ipv6_enable_ddns" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    use_ipv6_enable_ddns = %q
}
`, name, mac, addressType, matchClient, useIpv6EnableDdns)
}

func testAccRoaminghostUseIpv6Options(name string, mac string, addressType string, matchClient string, useIpv6Options string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_ipv6_options" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    use_ipv6_options = %q
}
`, name, mac, addressType, matchClient, useIpv6Options)
}

func testAccRoaminghostUseNextserver(name string, mac string, addressType string, matchClient string, useNextserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_nextserver" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    use_nextserver = %q
}
`, name, mac, addressType, matchClient, useNextserver)
}

func testAccRoaminghostUseOptions(name string, mac string, addressType string, matchClient string, useOptions string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_options" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    use_options = %q
}
`, name, mac, addressType, matchClient, useOptions)
}

func testAccRoaminghostUsePreferredLifetime(name string, mac string, addressType string, matchClient string, usePreferredLifetime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_preferred_lifetime" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    use_preferred_lifetime = %q
}
`, name, mac, addressType, matchClient, usePreferredLifetime)
}

func testAccRoaminghostUsePxeLeaseTime(name string, mac string, addressType string, matchClient string, usePxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_pxe_lease_time" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    use_pxe_lease_time = %q
}
`, name, mac, addressType, matchClient, usePxeLeaseTime)
}

func testAccRoaminghostUseValidLifetime(name string, mac string, addressType string, matchClient string, useValidLifetime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_valid_lifetime" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    use_valid_lifetime = %q
}
`, name, mac, addressType, matchClient, useValidLifetime)
}

func testAccRoaminghostValidLifetime(name string, mac string, addressType string, matchClient string, validLifetime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_valid_lifetime" {
    name = %q
    mac = %q
    address_type = %q
    match_client = %q
    valid_lifetime = %q
    use_valid_lifetime = true
}
`, name, mac, addressType, matchClient, validLifetime)
}
