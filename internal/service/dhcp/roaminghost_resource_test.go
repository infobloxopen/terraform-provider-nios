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

/*
// Manage dhcp Roaminghost with Basic Fields
resource "nios_dhcp_roaminghost" "dhcp_roaminghost_basic" {
    name = "NAME_REPLACE_ME"
}

// Manage dhcp Roaminghost with Additional Fields
resource "nios_dhcp_roaminghost" "dhcp_roaminghost_with_additional_fields" {
    name = "NAME_REPLACE_ME"

// TODO : Add additional optional fields below

    //Extensible Attributes
    extattrs = {
        Site = "location-1"
    }
}
*/

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
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "43200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// func TestAccRoaminghostResource_disappears(t *testing.T) {
// 	resourceName := "nios_dhcp_roaminghost.test"
// 	var v dhcp.Roaminghost

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		CheckDestroy:             testAccCheckRoaminghostDestroy(context.Background(), &v),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccRoaminghostBasicConfig("NAME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
// 					testAccCheckRoaminghostDisappears(context.Background(), &v),
// 				),
// 				ExpectNonEmptyPlan: true,
// 			},
// 		},
// 	})
// }

// func TestAccRoaminghostResource_Import(t *testing.T) {
// 	var resourceName = "nios_dhcp_roaminghost.test"
// 	var v dhcp.Roaminghost

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccRoaminghostBasicConfig("NAME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
// 				),
// 			},
// 			// Import with PlanOnly to detect differences
// 			{
// 				ResourceName:                         resourceName,
// 				ImportState:                          true,
// 				ImportStateIdFunc:                    testAccRoaminghostImportStateIdFunc(resourceName),
// 				ImportStateVerify:                    true,
// 				ImportStateVerifyIdentifierAttribute: "ref",
// 				PlanOnly:                             true,
// 			},
// 			// Import and Verify
// 			{
// 				ResourceName:                         resourceName,
// 				ImportState:                          true,
// 				ImportStateIdFunc:                    testAccRoaminghostImportStateIdFunc(resourceName),
// 				ImportStateVerify:                    true,
// 				ImportStateVerifyIgnore:              []string{"extattrs_all"},
// 				ImportStateVerifyIdentifierAttribute: "ref",
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

func TestAccRoaminghostResource_AddressType(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_address_type"
	var v dhcp.Roaminghost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostAddressType("ADDRESS_TYPE_REPLACE_ME", "BOTH"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address_type", "BOTH"),
				),
			},
			{
				Config: testAccRoaminghostAddressType("ADDRESS_TYPE_REPLACE_ME", "IPV4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address_type", "IPV4"),
				),
			},
			{
				Config: testAccRoaminghostAddressType("ADDRESS_TYPE_REPLACE_ME", "IPV6"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostBootfile("NAME_REPLACE_ME", "BOOTFILE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "BOOTFILE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostBootfile("NAME_REPLACE_ME", "BOOTFILE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "BOOTFILE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Bootserver(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_bootserver"
	var v dhcp.Roaminghost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostBootserver("NAME_REPLACE_ME", "BOOTSERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "BOOTSERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostBootserver("NAME_REPLACE_ME", "BOOTSERVER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "BOOTSERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_ClientIdentifierPrependZero(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_client_identifier_prepend_zero"
	var v dhcp.Roaminghost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostClientIdentifierPrependZero("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_identifier_prepend_zero", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostClientIdentifierPrependZero("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostComment("NAME_REPLACE_ME", "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostComment("NAME_REPLACE_ME", "Updated comment for the object"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostDdnsDomainname("NAME_REPLACE_ME", "DDNS_DOMAINNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "DDNS_DOMAINNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostDdnsDomainname("NAME_REPLACE_ME", "DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_DdnsHostname(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_ddns_hostname"
	var v dhcp.Roaminghost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostDdnsHostname("NAME_REPLACE_ME", "DDNS_HOSTNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_hostname", "DDNS_HOSTNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostDdnsHostname("NAME_REPLACE_ME", "DDNS_HOSTNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_hostname", "DDNS_HOSTNAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_DenyBootp(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_deny_bootp"
	var v dhcp.Roaminghost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostDenyBootp("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostDenyBootp("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostDhcpClientIdentifier("NAME_REPLACE_ME", "DHCP_CLIENT_IDENTIFIER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_client_identifier", "DHCP_CLIENT_IDENTIFIER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostDhcpClientIdentifier("NAME_REPLACE_ME", "DHCP_CLIENT_IDENTIFIER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_client_identifier", "DHCP_CLIENT_IDENTIFIER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Disable(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_disable"
	var v dhcp.Roaminghost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostDisable("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostDisable("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostEnableDdns("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostEnableDdns("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostEnablePxeLeaseTime("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostEnablePxeLeaseTime("NAME_REPLACE_ME", "false"),
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
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostExtAttrs("NAME_REPLACE_ME", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostExtAttrs("NAME_REPLACE_ME", map[string]string{
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostForceRoamingHostname("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "force_roaming_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostForceRoamingHostname("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIgnoreDhcpOptionListRequest("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostIgnoreDhcpOptionListRequest("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIpv6DdnsDomainname("NAME_REPLACE_ME", "IPV6_DDNS_DOMAINNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_ddns_domainname", "IPV6_DDNS_DOMAINNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostIpv6DdnsDomainname("NAME_REPLACE_ME", "IPV6_DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_ddns_domainname", "IPV6_DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Ipv6DdnsHostname(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_ipv6_ddns_hostname"
	var v dhcp.Roaminghost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIpv6DdnsHostname("NAME_REPLACE_ME", "IPV6_DDNS_HOSTNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_ddns_hostname", "IPV6_DDNS_HOSTNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostIpv6DdnsHostname("NAME_REPLACE_ME", "IPV6_DDNS_HOSTNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_ddns_hostname", "IPV6_DDNS_HOSTNAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Ipv6DomainName(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_ipv6_domain_name"
	var v dhcp.Roaminghost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIpv6DomainName("NAME_REPLACE_ME", "IPV6_DOMAIN_NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_domain_name", "IPV6_DOMAIN_NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostIpv6DomainName("NAME_REPLACE_ME", "IPV6_DOMAIN_NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_domain_name", "IPV6_DOMAIN_NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Ipv6DomainNameServers(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_ipv6_domain_name_servers"
	var v dhcp.Roaminghost
	ipv6DomainNameServersVal := []string{"IPV6_DOMAIN_NAME_SERVERS_REPLACE_ME1", "IPV6_DOMAIN_NAME_SERVERS_REPLACE_ME2"}
	ipv6DomainNameServersValUpdated := []string{"IPV6_DOMAIN_NAME_SERVERS_REPLACE_ME1", "IPV6_DOMAIN_NAME_SERVERS_REPLACE_ME2"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIpv6DomainNameServers("NAME_REPLACE_ME", ipv6DomainNameServersVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_domain_name_servers", "IPV6_DOMAIN_NAME_SERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostIpv6DomainNameServers("NAME_REPLACE_ME", ipv6DomainNameServersValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_domain_name_servers", "IPV6_DOMAIN_NAME_SERVERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Ipv6Duid(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_ipv6_duid"
	var v dhcp.Roaminghost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIpv6Duid("NAME_REPLACE_ME", "IPV6_DUID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_duid", "IPV6_DUID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostIpv6Duid("NAME_REPLACE_ME", "IPV6_DUID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_duid", "IPV6_DUID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Ipv6EnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_ipv6_enable_ddns"
	var v dhcp.Roaminghost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIpv6EnableDdns("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostIpv6EnableDdns("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIpv6ForceRoamingHostname("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_force_roaming_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostIpv6ForceRoamingHostname("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIpv6MacAddress("NAME_REPLACE_ME", "IPV6_MAC_ADDRESS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_mac_address", "IPV6_MAC_ADDRESS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostIpv6MacAddress("NAME_REPLACE_ME", "IPV6_MAC_ADDRESS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_mac_address", "IPV6_MAC_ADDRESS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Ipv6MatchOption(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_ipv6_match_option"
	var v dhcp.Roaminghost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIpv6MatchOption("IPV6_MATCH_OPTION_REPLACE_ME", "DUID"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_match_option", "DUID"),
				),
			},
			{
				Config: testAccRoaminghostIpv6MatchOption("IPV6_MATCH_OPTION_REPLACE_ME", "V6_MAC_ADDRESS"),
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
	ipv6OptionsVal := []map[string]any{}
	ipv6OptionsValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIpv6Options("NAME_REPLACE_ME", ipv6OptionsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_options", "IPV6_OPTIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostIpv6Options("NAME_REPLACE_ME", ipv6OptionsValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_options", "IPV6_OPTIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Ipv6Template(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_ipv6_template"
	var v dhcp.Roaminghost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostIpv6Template("NAME_REPLACE_ME", "IPV6_TEMPLATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_template", "IPV6_TEMPLATE_REPLACE_ME"),
				),
			},
			// Skip Update testing as this field cannot be updated
		},
	})
}

func TestAccRoaminghostResource_Mac(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_mac"
	var v dhcp.Roaminghost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostMac("NAME_REPLACE_ME", "MAC_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mac", "MAC_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostMac("NAME_REPLACE_ME", "MAC_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mac", "MAC_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_MatchClient(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_match_client"
	var v dhcp.Roaminghost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostMatchClient("MATCH_CLIENT_REPLACE_ME", "CLIENT_ID"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_client", "CLIENT_ID"),
				),
			},
			{
				Config: testAccRoaminghostMatchClient("MATCH_CLIENT_REPLACE_ME", "MAC_ADDRESS"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_NetworkView(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_network_view"
	var v dhcp.Roaminghost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostNetworkView("NAME_REPLACE_ME", "NETWORK_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostNetworkView("NAME_REPLACE_ME", "NETWORK_VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Nextserver(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_nextserver"
	var v dhcp.Roaminghost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostNextserver("NAME_REPLACE_ME", "NEXTSERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", "NEXTSERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostNextserver("NAME_REPLACE_ME", "NEXTSERVER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", "NEXTSERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_Options(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_options"
	var v dhcp.Roaminghost
	optionsVal := []map[string]any{}
	optionsValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostOptions("NAME_REPLACE_ME", optionsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostOptions("NAME_REPLACE_ME", optionsValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRoaminghostResource_PreferredLifetime(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_preferred_lifetime"
	var v dhcp.Roaminghost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostPreferredLifetime("NAME_REPLACE_ME", "PREFERRED_LIFETIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "PREFERRED_LIFETIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostPreferredLifetime("NAME_REPLACE_ME", "PREFERRED_LIFETIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "PREFERRED_LIFETIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccRoaminghostResource_PxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_pxe_lease_time"
	var v dhcp.Roaminghost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostPxeLeaseTime("NAME_REPLACE_ME", "PXE_LEASE_TIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "PXE_LEASE_TIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostPxeLeaseTime("NAME_REPLACE_ME", "PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccRoaminghostResource_Template(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_template"
	var v dhcp.Roaminghost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostTemplate("NAME_REPLACE_ME", "TEMPLATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template", "TEMPLATE_REPLACE_ME"),
				),
			},
			// Skip Update testing as this field cannot be updated
		},
	})
}

func TestAccRoaminghostResource_UseBootfile(t *testing.T) {
	var resourceName = "nios_dhcp_roaminghost.test_use_bootfile"
	var v dhcp.Roaminghost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseBootfile("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseBootfile("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseBootserver("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseBootserver("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseDdnsDomainname("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseDdnsDomainname("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseDenyBootp("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseDenyBootp("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseEnableDdns("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseEnableDdns("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseIgnoreDhcpOptionListRequest("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseIgnoreDhcpOptionListRequest("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseIpv6DdnsDomainname("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipv6_ddns_domainname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseIpv6DdnsDomainname("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseIpv6DomainName("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipv6_domain_name", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseIpv6DomainName("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseIpv6DomainNameServers("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipv6_domain_name_servers", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseIpv6DomainNameServers("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseIpv6EnableDdns("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipv6_enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseIpv6EnableDdns("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseIpv6Options("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipv6_options", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseIpv6Options("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseNextserver("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseNextserver("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseOptions("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseOptions("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUsePreferredLifetime("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUsePreferredLifetime("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUsePxeLeaseTime("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUsePxeLeaseTime("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostUseValidLifetime("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostUseValidLifetime("NAME_REPLACE_ME", "false"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRoaminghostValidLifetime("NAME_REPLACE_ME", "VALID_LIFETIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "VALID_LIFETIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRoaminghostValidLifetime("NAME_REPLACE_ME", "VALID_LIFETIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoaminghostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "VALID_LIFETIME_UPDATE_REPLACE_ME"),
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

func testAccRoaminghostAddressType(name string, addressType string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_address_type" {
    name = %q
    address_type = %q
}
`, name, addressType)
}

func testAccRoaminghostBootfile(name string, bootfile string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_bootfile" {
    name = %q
    bootfile = %q
    use_bootfile = true
}
`, name, bootfile)
}

func testAccRoaminghostBootserver(name string, bootserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_bootserver" {
    name = %q
    bootserver = %q
    use_bootserver = true
}
`, name, bootserver)
}

func testAccRoaminghostClientIdentifierPrependZero(name string, clientIdentifierPrependZero string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_client_identifier_prepend_zero" {
    name = %q
    client_identifier_prepend_zero = %q
}
`, name, clientIdentifierPrependZero)
}

func testAccRoaminghostComment(name string, comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccRoaminghostDdnsDomainname(name string, ddnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ddns_domainname" {
    name = %q
    ddns_domainname = %q
    use_ddns_domainname = true
}
`, name, ddnsDomainname)
}

func testAccRoaminghostDdnsHostname(name string, ddnsHostname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ddns_hostname" {
    name = %q
    ddns_hostname = %q
}
`, name, ddnsHostname)
}

func testAccRoaminghostDenyBootp(name string, denyBootp string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_deny_bootp" {
    name = %q
    deny_bootp = %q
    use_deny_bootp = true
}
`, name, denyBootp)
}

func testAccRoaminghostDhcpClientIdentifier(name string, dhcpClientIdentifier string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_dhcp_client_identifier" {
    name = %q
    dhcp_client_identifier = %q
}
`, name, dhcpClientIdentifier)
}

func testAccRoaminghostDisable(name string, disable string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_disable" {
    name = %q
    disable = %q
}
`, name, disable)
}

func testAccRoaminghostEnableDdns(name string, enableDdns string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_enable_ddns" {
    name = %q
    enable_ddns = %q
    use_enable_ddns = true
}
`, name, enableDdns)
}

func testAccRoaminghostEnablePxeLeaseTime(name string, enablePxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_enable_pxe_lease_time" {
    name = %q
    enable_pxe_lease_time = %q
}
`, name, enablePxeLeaseTime)
}

func testAccRoaminghostExtAttrs(name string, extAttrs map[string]string) string {
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_extattrs" {
    name = %q
    extattrs = %s
}
`, name, extAttrsStr)
}

func testAccRoaminghostForceRoamingHostname(name string, forceRoamingHostname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_force_roaming_hostname" {
    name = %q
    force_roaming_hostname = %q
}
`, name, forceRoamingHostname)
}

func testAccRoaminghostIgnoreDhcpOptionListRequest(name string, ignoreDhcpOptionListRequest string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ignore_dhcp_option_list_request" {
    name = %q
    ignore_dhcp_option_list_request = %q
    use_ignore_dhcp_option_list_request = true
}
`, name, ignoreDhcpOptionListRequest)
}

func testAccRoaminghostIpv6DdnsDomainname(name string, ipv6DdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_ddns_domainname" {
    name = %q
    ipv6_ddns_domainname = %q
    use_ipv6_ddns_domainname = true
}
`, name, ipv6DdnsDomainname)
}

func testAccRoaminghostIpv6DdnsHostname(name string, ipv6DdnsHostname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_ddns_hostname" {
    name = %q
    ipv6_ddns_hostname = %q
}
`, name, ipv6DdnsHostname)
}

func testAccRoaminghostIpv6DomainName(name string, ipv6DomainName string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_domain_name" {
    name = %q
    ipv6_domain_name = %q
    use_ipv6_domain_name = true
}
`, name, ipv6DomainName)
}

func testAccRoaminghostIpv6DomainNameServers(name string, ipv6DomainNameServers []string) string {
	ipv6DomainNameServersStr := utils.ConvertStringSliceToHCL(ipv6DomainNameServers)
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_domain_name_servers" {
    name = %q
    ipv6_domain_name_servers = %s
    use_ipv6_domain_name_servers = true
}
`, name, ipv6DomainNameServersStr)
}

func testAccRoaminghostIpv6Duid(name string, ipv6Duid string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_duid" {
    name = %q
    ipv6_duid = %q
}
`, name, ipv6Duid)
}

func testAccRoaminghostIpv6EnableDdns(name string, ipv6EnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_enable_ddns" {
    name = %q
    ipv6_enable_ddns = %q
    use_ipv6_enable_ddns = true
}
`, name, ipv6EnableDdns)
}

func testAccRoaminghostIpv6ForceRoamingHostname(name string, ipv6ForceRoamingHostname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_force_roaming_hostname" {
    name = %q
    ipv6_force_roaming_hostname = %q
}
`, name, ipv6ForceRoamingHostname)
}

func testAccRoaminghostIpv6MacAddress(name string, ipv6MacAddress string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_mac_address" {
    name = %q
    ipv6_mac_address = %q
}
`, name, ipv6MacAddress)
}

func testAccRoaminghostIpv6MatchOption(name string, ipv6MatchOption string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_match_option" {
    name = %q
    ipv6_match_option = %q
}
`, name, ipv6MatchOption)
}

func testAccRoaminghostIpv6Options(name string, ipv6Options []map[string]any) string {
	ipv6OptionsStr := utils.ConvertSliceOfMapsToHCL(ipv6Options)
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_options" {
    name = %q
    ipv6_options = %s
    use_ipv6_options = true
}
`, name, ipv6OptionsStr)
}

func testAccRoaminghostIpv6Template(name string, ipv6Template string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_ipv6_template" {
    name = %q
    ipv6_template = %q
}
`, name, ipv6Template)
}

func testAccRoaminghostMac(name string, mac string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_mac" {
    name = %q
    mac = %q
}
`, name, mac)
}

func testAccRoaminghostMatchClient(name string, matchClient string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_match_client" {
    name = %q
    match_client = %q
}
`, name, matchClient)
}

func testAccRoaminghostName(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_name" {
    name = %q
}
`, name)
}

func testAccRoaminghostNetworkView(name string, networkView string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_network_view" {
    name = %q
    network_view = %q
}
`, name, networkView)
}

func testAccRoaminghostNextserver(name string, nextserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_nextserver" {
    name = %q
    nextserver = %q
    use_nextserver = true
}
`, name, nextserver)
}

func testAccRoaminghostOptions(name string, options []map[string]any) string {
	optionsStr := utils.ConvertSliceOfMapsToHCL(options)
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_options" {
    name = %q
    options = %s
    use_options = true
}
`, name, optionsStr)
}

func testAccRoaminghostPreferredLifetime(name string, preferredLifetime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_preferred_lifetime" {
    name = %q
    preferred_lifetime = %q
    use_preferred_lifetime = true
}
`, name, preferredLifetime)
}

func testAccRoaminghostPxeLeaseTime(name string, pxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_pxe_lease_time" {
    name = %q
    pxe_lease_time = %q
    use_pxe_lease_time = true
}
`, name, pxeLeaseTime)
}

func testAccRoaminghostTemplate(name string, template string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_template" {
    name = %q
    template = %q
}
`, name, template)
}

func testAccRoaminghostUseBootfile(name string, useBootfile string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_bootfile" {
    name = %q
    use_bootfile = %q
}
`, name, useBootfile)
}

func testAccRoaminghostUseBootserver(name string, useBootserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_bootserver" {
    name = %q
    use_bootserver = %q
}
`, name, useBootserver)
}

func testAccRoaminghostUseDdnsDomainname(name string, useDdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_ddns_domainname" {
    name = %q
    use_ddns_domainname = %q
}
`, name, useDdnsDomainname)
}

func testAccRoaminghostUseDenyBootp(name string, useDenyBootp string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_deny_bootp" {
    name = %q
    use_deny_bootp = %q
}
`, name, useDenyBootp)
}

func testAccRoaminghostUseEnableDdns(name string, useEnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_enable_ddns" {
    name = %q
    use_enable_ddns = %q
}
`, name, useEnableDdns)
}

func testAccRoaminghostUseIgnoreDhcpOptionListRequest(name string, useIgnoreDhcpOptionListRequest string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_ignore_dhcp_option_list_request" {
    name = %q
    use_ignore_dhcp_option_list_request = %q
}
`, name, useIgnoreDhcpOptionListRequest)
}

func testAccRoaminghostUseIpv6DdnsDomainname(name string, useIpv6DdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_ipv6_ddns_domainname" {
    name = %q
    use_ipv6_ddns_domainname = %q
}
`, name, useIpv6DdnsDomainname)
}

func testAccRoaminghostUseIpv6DomainName(name string, useIpv6DomainName string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_ipv6_domain_name" {
    name = %q
    use_ipv6_domain_name = %q
}
`, name, useIpv6DomainName)
}

func testAccRoaminghostUseIpv6DomainNameServers(name string, useIpv6DomainNameServers string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_ipv6_domain_name_servers" {
    name = %q
    use_ipv6_domain_name_servers = %q
}
`, name, useIpv6DomainNameServers)
}

func testAccRoaminghostUseIpv6EnableDdns(name string, useIpv6EnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_ipv6_enable_ddns" {
    name = %q
    use_ipv6_enable_ddns = %q
}
`, name, useIpv6EnableDdns)
}

func testAccRoaminghostUseIpv6Options(name string, useIpv6Options string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_ipv6_options" {
    name = %q
    use_ipv6_options = %q
}
`, name, useIpv6Options)
}

func testAccRoaminghostUseNextserver(name string, useNextserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_nextserver" {
    name = %q
    use_nextserver = %q
}
`, name, useNextserver)
}

func testAccRoaminghostUseOptions(name string, useOptions string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_options" {
    name = %q
    use_options = %q
}
`, name, useOptions)
}

func testAccRoaminghostUsePreferredLifetime(name string, usePreferredLifetime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_preferred_lifetime" {
    name = %q
    use_preferred_lifetime = %q
}
`, name, usePreferredLifetime)
}

func testAccRoaminghostUsePxeLeaseTime(name string, usePxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_pxe_lease_time" {
    name = %q
    use_pxe_lease_time = %q
}
`, name, usePxeLeaseTime)
}

func testAccRoaminghostUseValidLifetime(name string, useValidLifetime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_use_valid_lifetime" {
    name = %q
    use_valid_lifetime = %q
}
`, name, useValidLifetime)
}

func testAccRoaminghostValidLifetime(name string, validLifetime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_roaminghost" "test_valid_lifetime" {
    name = %q
    valid_lifetime = %q
    use_valid_lifetime = true
}
`, name, validLifetime)
}
