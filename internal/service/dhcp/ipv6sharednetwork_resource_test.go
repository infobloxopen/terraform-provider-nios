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
// Manage dhcp Ipv6sharednetwork with Basic Fields
resource "nios_dhcp_ipv6sharednetwork" "dhcp_ipv6sharednetwork_basic" {
    name = "NAME_REPLACE_ME"
    networks = "NETWORKS_REPLACE_ME"
}

// Manage dhcp Ipv6sharednetwork with Additional Fields
resource "nios_dhcp_ipv6sharednetwork" "dhcp_ipv6sharednetwork_with_additional_fields" {
    name = "NAME_REPLACE_ME"
    networks = "NETWORKS_REPLACE_ME"

// TODO : Add additional optional fields below

    //Extensible Attributes
    extattrs = {
        Site = "location-1"
    }
}
*/

var readableAttributesForIpv6sharednetwork = "comment,ddns_domainname,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,ddns_use_option81,disable,domain_name,domain_name_servers,enable_ddns,extattrs,logic_filter_rules,name,network_view,networks,options,preferred_lifetime,update_dns_on_lease_renewal,use_ddns_domainname,use_ddns_generate_hostname,use_ddns_ttl,use_ddns_use_option81,use_domain_name,use_domain_name_servers,use_enable_ddns,use_logic_filter_rules,use_options,use_preferred_lifetime,use_update_dns_on_lease_renewal,use_valid_lifetime,valid_lifetime"

func TestAccIpv6sharednetworkResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkBasicConfig("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
					resource.TestCheckResourceAttr(resourceName, "networks", "NETWORKS_REPLACE_ME"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "true"),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "0"),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "false"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "27000"),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "false"),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "43200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_ipv6sharednetwork.test"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6sharednetworkDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6sharednetworkBasicConfig("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					testAccCheckIpv6sharednetworkDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccIpv6sharednetworkResource_Import(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkBasicConfig("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
				),
			},
			// Import with PlanOnly to detect differences
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccIpv6sharednetworkImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "ref",
				PlanOnly:                             true,
				//ExpectError:                          regexp.MustCompile(`ImportStateVerify attributes not equivalent`),
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccIpv6sharednetworkImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_comment"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkComment("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkComment("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "Updated comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the object"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_DdnsDomainname(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_ddns_domainname"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkDdnsDomainname("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "DDNS_DOMAINNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "DDNS_DOMAINNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkDdnsDomainname("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_DdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_ddns_generate_hostname"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkDdnsGenerateHostname("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkDdnsGenerateHostname("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_DdnsServerAlwaysUpdates(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_ddns_server_always_updates"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkDdnsServerAlwaysUpdates("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkDdnsServerAlwaysUpdates("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_DdnsTtl(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_ddns_ttl"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkDdnsTtl("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "DDNS_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "DDNS_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkDdnsTtl("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "DDNS_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "DDNS_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccIpv6sharednetworkResource_DdnsUseOption81(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_ddns_use_option81"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkDdnsUseOption81("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkDdnsUseOption81("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_Disable(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_disable"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkDisable("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkDisable("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_DomainName(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_domain_name"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkDomainName("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "DOMAIN_NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name", "DOMAIN_NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkDomainName("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "DOMAIN_NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name", "DOMAIN_NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_DomainNameServers(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_domain_name_servers"
	var v dhcp.Ipv6sharednetwork
	domainNameServersVal := []string{"DOMAIN_NAME_SERVERS_REPLACE_ME1", "DOMAIN_NAME_SERVERS_REPLACE_ME2"}
	domainNameServersValUpdated := []string{"DOMAIN_NAME_SERVERS_REPLACE_ME1", "DOMAIN_NAME_SERVERS_REPLACE_ME2"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkDomainNameServers("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", domainNameServersVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers", "DOMAIN_NAME_SERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkDomainNameServers("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", domainNameServersValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers", "DOMAIN_NAME_SERVERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_EnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_enable_ddns"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkEnableDdns("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkEnableDdns("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_extattrs"
	var v dhcp.Ipv6sharednetwork
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkExtAttrs("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkExtAttrs("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_LogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_logic_filter_rules"
	var v dhcp.Ipv6sharednetwork
	logicFilterRulesVal := []map[string]any{}
	logicFilterRulesValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkLogicFilterRules("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", logicFilterRulesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkLogicFilterRules("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", logicFilterRulesValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_name"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkName("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkName("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_NetworkView(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_network_view"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkNetworkView("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "NETWORK_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_REPLACE_ME"),
				),
			},
			// Skip Update testing as this field cannot be updated
		},
	})
}

func TestAccIpv6sharednetworkResource_Networks(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_networks"
	var v dhcp.Ipv6sharednetwork
	networksVal := []map[string]any{}
	networksValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkNetworks("NAME_REPLACE_ME", networksVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "networks", "NETWORKS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkNetworks("NAME_REPLACE_ME", networksValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "networks", "NETWORKS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_Options(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_options"
	var v dhcp.Ipv6sharednetwork
	optionsVal := []map[string]any{}
	optionsValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkOptions("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", optionsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkOptions("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", optionsValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_PreferredLifetime(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_preferred_lifetime"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkPreferredLifetime("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "PREFERRED_LIFETIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "PREFERRED_LIFETIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkPreferredLifetime("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "PREFERRED_LIFETIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "PREFERRED_LIFETIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccIpv6sharednetworkResource_UpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_update_dns_on_lease_renewal"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUpdateDnsOnLeaseRenewal("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUpdateDnsOnLeaseRenewal("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_UseDdnsDomainname(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_use_ddns_domainname"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseDdnsDomainname("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseDdnsDomainname("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_UseDdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_use_ddns_generate_hostname"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseDdnsGenerateHostname("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseDdnsGenerateHostname("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_UseDdnsTtl(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_use_ddns_ttl"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseDdnsTtl("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseDdnsTtl("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_UseDdnsUseOption81(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_use_ddns_use_option81"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseDdnsUseOption81("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseDdnsUseOption81("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_UseDomainName(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_use_domain_name"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseDomainName("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseDomainName("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_UseDomainNameServers(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_use_domain_name_servers"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseDomainNameServers("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseDomainNameServers("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_UseEnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_use_enable_ddns"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseEnableDdns("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseEnableDdns("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_UseLogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_use_logic_filter_rules"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseLogicFilterRules("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseLogicFilterRules("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_UseOptions(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_use_options"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseOptions("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseOptions("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_UsePreferredLifetime(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_use_preferred_lifetime"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUsePreferredLifetime("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUsePreferredLifetime("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_UseUpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_use_update_dns_on_lease_renewal"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseUpdateDnsOnLeaseRenewal("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseUpdateDnsOnLeaseRenewal("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_UseValidLifetime(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_use_valid_lifetime"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseValidLifetime("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseValidLifetime("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_ValidLifetime(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_valid_lifetime"
	var v dhcp.Ipv6sharednetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkValidLifetime("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "VALID_LIFETIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "VALID_LIFETIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkValidLifetime("NAME_REPLACE_ME", "NETWORKS_REPLACE_ME", "VALID_LIFETIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "VALID_LIFETIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckIpv6sharednetworkExists(ctx context.Context, resourceName string, v *dhcp.Ipv6sharednetwork) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			Ipv6sharednetworkAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForIpv6sharednetwork).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetIpv6sharednetworkResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetIpv6sharednetworkResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckIpv6sharednetworkDestroy(ctx context.Context, v *dhcp.Ipv6sharednetwork) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			Ipv6sharednetworkAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForIpv6sharednetwork).
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

func testAccCheckIpv6sharednetworkDisappears(ctx context.Context, v *dhcp.Ipv6sharednetwork) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			Ipv6sharednetworkAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccIpv6sharednetworkImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccIpv6sharednetworkBasicConfig(name, networks string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test" {
    name = %q
    networks = %q
}
`, name, networks)
}

func testAccIpv6sharednetworkComment(name string, networks string, comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_comment" {
    name = %q
    networks = %q
    comment = %q
}
`, name, networks, comment)
}

func testAccIpv6sharednetworkDdnsDomainname(name string, networks string, ddnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_ddns_domainname" {
    name = %q
    networks = %q
    ddns_domainname = %q
}
`, name, networks, ddnsDomainname)
}

func testAccIpv6sharednetworkDdnsGenerateHostname(name string, networks string, ddnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_ddns_generate_hostname" {
    name = %q
    networks = %q
    ddns_generate_hostname = %q
}
`, name, networks, ddnsGenerateHostname)
}

func testAccIpv6sharednetworkDdnsServerAlwaysUpdates(name string, networks string, ddnsServerAlwaysUpdates string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_ddns_server_always_updates" {
    name = %q
    networks = %q
    ddns_server_always_updates = %q
}
`, name, networks, ddnsServerAlwaysUpdates)
}

func testAccIpv6sharednetworkDdnsTtl(name string, networks string, ddnsTtl string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_ddns_ttl" {
    name = %q
    networks = %q
    ddns_ttl = %q
}
`, name, networks, ddnsTtl)
}

func testAccIpv6sharednetworkDdnsUseOption81(name string, networks string, ddnsUseOption81 string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_ddns_use_option81" {
    name = %q
    networks = %q
    ddns_use_option81 = %q
}
`, name, networks, ddnsUseOption81)
}

func testAccIpv6sharednetworkDisable(name string, networks string, disable string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_disable" {
    name = %q
    networks = %q
    disable = %q
}
`, name, networks, disable)
}

func testAccIpv6sharednetworkDomainName(name string, networks string, domainName string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_domain_name" {
    name = %q
    networks = %q
    domain_name = %q
}
`, name, networks, domainName)
}

func testAccIpv6sharednetworkDomainNameServers(name string, networks string, domainNameServers []string) string {
	domainNameServersStr := utils.ConvertStringSliceToHCL(domainNameServers)
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_domain_name_servers" {
    name = %q
    networks = %q
    domain_name_servers = %q
}
`, name, networks, domainNameServersStr)
}

func testAccIpv6sharednetworkEnableDdns(name string, networks string, enableDdns string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_enable_ddns" {
    name = %q
    networks = %q
    enable_ddns = %q
}
`, name, networks, enableDdns)
}

func testAccIpv6sharednetworkExtAttrs(name string, networks string, extAttrs map[string]string) string {
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_extattrs" {
    name = %q
    networks = %q
    extattrs = %s
}
`, name, networks, extAttrsStr)
}

func testAccIpv6sharednetworkLogicFilterRules(name string, networks string, logicFilterRules []map[string]any) string {
	logicFilterRulesStr := utils.ConvertSliceOfMapsToHCL(logicFilterRules)
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_logic_filter_rules" {
    name = %q
    networks = %q
    logic_filter_rules = %s
}
`, name, networks, logicFilterRulesStr)
}

func testAccIpv6sharednetworkName(name string, networks string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_name" {
    name = %q
    networks = %q
}
`, name, networks)
}

func testAccIpv6sharednetworkNetworkView(name string, networks string, networkView string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_network_view" {
    name = %q
    networks = %q
    network_view = %q
}
`, name, networks, networkView)
}

func testAccIpv6sharednetworkNetworks(name string, networks []map[string]any) string {
	networksStr := utils.ConvertSliceOfMapsToHCL(networks)
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_networks" {
    name = %q
    networks = %s
}
`, name, networksStr)
}

func testAccIpv6sharednetworkOptions(name string, networks string, options []map[string]any) string {
	optionsStr := utils.ConvertSliceOfMapsToHCL(options)
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_options" {
    name = %q
    networks = %q
    options = %s
}
`, name, networks, optionsStr)
}

func testAccIpv6sharednetworkPreferredLifetime(name string, networks string, preferredLifetime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_preferred_lifetime" {
    name = %q
    networks = %q
    preferred_lifetime = %q
}
`, name, networks, preferredLifetime)
}

func testAccIpv6sharednetworkUpdateDnsOnLeaseRenewal(name string, networks string, updateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_update_dns_on_lease_renewal" {
    name = %q
    networks = %q
    update_dns_on_lease_renewal = %q
}
`, name, networks, updateDnsOnLeaseRenewal)
}

func testAccIpv6sharednetworkUseDdnsDomainname(name string, networks string, useDdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_ddns_domainname" {
    name = %q
    networks = %q
    use_ddns_domainname = %q
}
`, name, networks, useDdnsDomainname)
}

func testAccIpv6sharednetworkUseDdnsGenerateHostname(name string, networks string, useDdnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_ddns_generate_hostname" {
    name = %q
    networks = %q
    use_ddns_generate_hostname = %q
}
`, name, networks, useDdnsGenerateHostname)
}

func testAccIpv6sharednetworkUseDdnsTtl(name string, networks string, useDdnsTtl string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_ddns_ttl" {
    name = %q
    networks = %q
    use_ddns_ttl = %q
}
`, name, networks, useDdnsTtl)
}

func testAccIpv6sharednetworkUseDdnsUseOption81(name string, networks string, useDdnsUseOption81 string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_ddns_use_option81" {
    name = %q
    networks = %q
    use_ddns_use_option81 = %q
}
`, name, networks, useDdnsUseOption81)
}

func testAccIpv6sharednetworkUseDomainName(name string, networks string, useDomainName string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_domain_name" {
    name = %q
    networks = %q
    use_domain_name = %q
}
`, name, networks, useDomainName)
}

func testAccIpv6sharednetworkUseDomainNameServers(name string, networks string, useDomainNameServers string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_domain_name_servers" {
    name = %q
    networks = %q
    use_domain_name_servers = %q
}
`, name, networks, useDomainNameServers)
}

func testAccIpv6sharednetworkUseEnableDdns(name string, networks string, useEnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_enable_ddns" {
    name = %q
    networks = %q
    use_enable_ddns = %q
}
`, name, networks, useEnableDdns)
}

func testAccIpv6sharednetworkUseLogicFilterRules(name string, networks string, useLogicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_logic_filter_rules" {
    name = %q
    networks = %q
    use_logic_filter_rules = %q
}
`, name, networks, useLogicFilterRules)
}

func testAccIpv6sharednetworkUseOptions(name string, networks string, useOptions string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_options" {
    name = %q
    networks = %q
    use_options = %q
}
`, name, networks, useOptions)
}

func testAccIpv6sharednetworkUsePreferredLifetime(name string, networks string, usePreferredLifetime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_preferred_lifetime" {
    name = %q
    networks = %q
    use_preferred_lifetime = %q
}
`, name, networks, usePreferredLifetime)
}

func testAccIpv6sharednetworkUseUpdateDnsOnLeaseRenewal(name string, networks string, useUpdateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_update_dns_on_lease_renewal" {
    name = %q
    networks = %q
    use_update_dns_on_lease_renewal = %q
}
`, name, networks, useUpdateDnsOnLeaseRenewal)
}

func testAccIpv6sharednetworkUseValidLifetime(name string, networks string, useValidLifetime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_valid_lifetime" {
    name = %q
    networks = %q
    use_valid_lifetime = %q
}
`, name, networks, useValidLifetime)
}

func testAccIpv6sharednetworkValidLifetime(name string, networks string, validLifetime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_valid_lifetime" {
    name = %q
    networks = %q
    valid_lifetime = %q
}
`, name, networks, validLifetime)
}
