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

var readableAttributesForIpv6sharednetwork = "comment,ddns_domainname,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,ddns_use_option81,disable,domain_name,domain_name_servers,enable_ddns,extattrs,logic_filter_rules,name,network_view,networks,options,preferred_lifetime,update_dns_on_lease_renewal,use_ddns_domainname,use_ddns_generate_hostname,use_ddns_ttl,use_ddns_use_option81,use_domain_name,use_domain_name_servers,use_enable_ddns,use_logic_filter_rules,use_options,use_preferred_lifetime,use_update_dns_on_lease_renewal,use_valid_lifetime,valid_lifetime"

func TestAccIpv6sharednetworkResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test"
	var v dhcp.Ipv6sharednetwork
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkBasicConfig(name, networks, network1, network2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "networks.#", "2"),
					resource.TestCheckResourceAttrPair(resourceName, "networks.0", "nios_ipam_ipv6network.test1", "ref"),
					resource.TestCheckResourceAttrPair(resourceName, "networks.1", "nios_ipam_ipv6network.test2", "ref"),
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
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6sharednetworkDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6sharednetworkBasicConfig(name, networks, network1, network2),
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
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkBasicConfig(name, networks, network1, network2),
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
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccIpv6sharednetworkImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"networks"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_comment"
	var v dhcp.Ipv6sharednetwork
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkComment(name, networks, network1, network2, "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkComment(name, networks, network1, network2, "Updated comment for the object"),
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
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkDdnsDomainname(name, networks, network1, network2, "example.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkDdnsDomainname(name, networks, network1, network2, "updated-example.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "updated-example.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_DdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_ddns_generate_hostname"
	var v dhcp.Ipv6sharednetwork
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkDdnsGenerateHostname(name, networks, network1, network2, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkDdnsGenerateHostname(name, networks, network1, network2, "false", "true"),
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
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkDdnsServerAlwaysUpdates(name, networks, network1, network2, "true", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkDdnsServerAlwaysUpdates(name, networks, network1, network2, "false", "true", "true"),
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
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkDdnsTtl(name, networks, network1, network2, "100", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "100"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkDdnsTtl(name, networks, network1, network2, "200", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccIpv6sharednetworkResource_DdnsUseOption81(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_ddns_use_option81"
	var v dhcp.Ipv6sharednetwork
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkDdnsUseOption81(name, networks, network1, network2, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkDdnsUseOption81(name, networks, network1, network2, "false", "true"),
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
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkDisable(name, networks, network1, network2, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkDisable(name, networks, network1, network2, "false"),
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
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkDomainName(name, networks, network1, network2, "example.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name", "example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkDomainName(name, networks, network1, network2, "updated-example.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name", "updated-example.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_DomainNameServers(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_domain_name_servers"
	var v dhcp.Ipv6sharednetwork
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}
	domainNameServersVal := []string{"2001:4860:4860::8888", "2001:4860:4860::9999", "2001:4860:4860::8899"}
	domainNameServersValUpdated := []string{"2001:4860:4860::8881", "2001:4860:4860::9991"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkDomainNameServers(name, networks, network1, network2, domainNameServersVal, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.0", "2001:4860:4860::8888"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.1", "2001:4860:4860::9999"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.2", "2001:4860:4860::8899"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkDomainNameServers(name, networks, network1, network2, domainNameServersValUpdated, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.0", "2001:4860:4860::8881"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.1", "2001:4860:4860::9991"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_EnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_enable_ddns"
	var v dhcp.Ipv6sharednetwork
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkEnableDdns(name, networks, network1, network2, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkEnableDdns(name, networks, network1, network2, "false", "true"),
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
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkExtAttrs(name, networks, network1, network2, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkExtAttrs(name, networks, network1, network2, map[string]string{
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
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}
	logicFilterRulesVal := []map[string]any{
		{
			"filter": "example-ipv6-option-filter-1",
			"type":   "Option",
		},
	}
	logicFilterRulesValUpdated := []map[string]any{
		{
			"filter": "example-ipv6-option-filter-2",
			"type":   "Option",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkLogicFilterRules(name, networks, network1, network2, logicFilterRulesVal, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.filter", "example-ipv6-option-filter-1"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.type", "Option"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkLogicFilterRules(name, networks, network1, network2, logicFilterRulesValUpdated, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.filter", "example-ipv6-option-filter-2"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.type", "Option"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_name"
	var v dhcp.Ipv6sharednetwork
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	nameUpdated := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkName(name, networks, network1, network2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkName(nameUpdated, networks, network1, network2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_NetworkView(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_network_view"
	var v dhcp.Ipv6sharednetwork
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkNetworkView(name, networks, network1, network2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "default"),
				),
			},
		},
	})
}

func TestAccIpv6sharednetworkResource_Networks(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_networks"
	var v dhcp.Ipv6sharednetwork
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networksVal := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}
	networksValUpdated := []string{
		"${nios_ipam_ipv6network.test1.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkNetworks(name, networksVal, network1, network2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "networks.#", "2"),
					resource.TestCheckResourceAttrPair(resourceName, "networks.0", "nios_ipam_ipv6network.test1", "ref"),
					resource.TestCheckResourceAttrPair(resourceName, "networks.1", "nios_ipam_ipv6network.test2", "ref"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkNetworks(name, networksValUpdated, network1, network2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "networks.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "networks.0", "nios_ipam_ipv6network.test1", "ref"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_Options(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_options"
	var v dhcp.Ipv6sharednetwork
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}
	optionsVal := []map[string]any{
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
	optionsValUpdated := []map[string]any{
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
				Config: testAccIpv6sharednetworkOptions(name, networks, network1, network2, optionsVal, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "domain-name"),
					resource.TestCheckResourceAttr(resourceName, "options.0.num", "15"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "options.1.num", "37"),
					resource.TestCheckResourceAttr(resourceName, "options.1.value", "remote-id"),
					resource.TestCheckResourceAttr(resourceName, "options.1.vendor_class", "DHCPv6"),
					resource.TestCheckResourceAttr(resourceName, "options.2.name", "dhcp6.subscriber-id"),
					resource.TestCheckResourceAttr(resourceName, "options.2.value", "subscriber-id"),
					resource.TestCheckResourceAttr(resourceName, "options.2.vendor_class", "DHCPv6"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkOptions(name, networks, network1, network2, optionsValUpdated, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "domain-name"),
					resource.TestCheckResourceAttr(resourceName, "options.0.num", "15"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "example.org"),
					resource.TestCheckResourceAttr(resourceName, "options.1.num", "37"),
					resource.TestCheckResourceAttr(resourceName, "options.1.value", "remote-id-updated"),
					resource.TestCheckResourceAttr(resourceName, "options.1.vendor_class", "DHCPv6"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_PreferredLifetime(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_preferred_lifetime"
	var v dhcp.Ipv6sharednetwork
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkPreferredLifetime(name, networks, network1, network2, "200", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "200"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkPreferredLifetime(name, networks, network1, network2, "400", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "400"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_UpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_update_dns_on_lease_renewal"
	var v dhcp.Ipv6sharednetwork
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUpdateDnsOnLeaseRenewal(name, networks, network1, network2, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUpdateDnsOnLeaseRenewal(name, networks, network1, network2, "false", "true"),
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
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseDdnsDomainname(name, networks, network1, network2, "example.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseDdnsDomainname(name, networks, network1, network2, "example.com", "false"),
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
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseDdnsGenerateHostname(name, networks, network1, network2, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseDdnsGenerateHostname(name, networks, network1, network2, "true", "false"),
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
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseDdnsTtl(name, networks, network1, network2, "100", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseDdnsTtl(name, networks, network1, network2, "100", "false"),
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
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseDdnsUseOption81(name, networks, network1, network2, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseDdnsUseOption81(name, networks, network1, network2, "false", "true"),
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
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseDomainName(name, networks, network1, network2, "true", "example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseDomainName(name, networks, network1, network2, "false", "example.com"),
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
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}
	domainNameServersVal := []string{"2001:4860:4860::8888", "2001:4860:4860::9999", "2001:4860:4860::8899"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseDomainNameServers(name, networks, network1, network2, "true", domainNameServersVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseDomainNameServers(name, networks, network1, network2, "false", domainNameServersVal),
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
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseEnableDdns(name, networks, network1, network2, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseEnableDdns(name, networks, network1, network2, "false", "true"),
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
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}
	logicFilterRulesVal := []map[string]any{
		{
			"filter": "example-ipv6-option-filter-1",
			"type":   "Option",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseLogicFilterRules(name, networks, network1, network2, "true", logicFilterRulesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseLogicFilterRules(name, networks, network1, network2, "false", logicFilterRulesVal),
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
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}
	optionsVal := []map[string]any{
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
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseOptions(name, networks, network1, network2, "true", optionsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseOptions(name, networks, network1, network2, "false", optionsVal),
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
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUsePreferredLifetime(name, networks, network1, network2, "true", "1000"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUsePreferredLifetime(name, networks, network1, network2, "false", "1000"),
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
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseUpdateDnsOnLeaseRenewal(name, networks, network1, network2, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseUpdateDnsOnLeaseRenewal(name, networks, network1, network2, "false", "true"),
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
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkUseValidLifetime(name, networks, network1, network2, "24000", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkUseValidLifetime(name, networks, network1, network2, "25000", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6sharednetworkResource_ValidLifetime(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6sharednetwork.test_valid_lifetime"
	var v dhcp.Ipv6sharednetwork
	name := acctest.RandomNameWithPrefix("ipv6sharednetwork")
	network1 := acctest.RandomIPv6Network()
	network2 := acctest.RandomIPv6Network()
	networks := []string{
		"${nios_ipam_ipv6network.test1.ref}",
		"${nios_ipam_ipv6network.test2.ref}",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6sharednetworkValidLifetime(name, networks, network1, network2, "30000", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "30000"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6sharednetworkValidLifetime(name, networks, network1, network2, "40000", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6sharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "40000"),
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

func testAccIpv6sharednetworkBasicConfig(name string, networks []string, network1 string, network2 string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test" {
    name = %q
    networks = %s
}
`, name, networksStr)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkComment(name string, networks []string, network1 string, network2 string, comment string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_comment" {
    name = %q
    networks = %s
    comment = %q
}
`, name, networksStr, comment)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkDdnsDomainname(name string, networks []string, network1 string, network2 string, ddnsDomainname string, useDdnsDomainname string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_ddns_domainname" {
    name = %q
    networks = %s
    ddns_domainname = %q
	use_ddns_domainname = %q
}
`, name, networksStr, ddnsDomainname, useDdnsDomainname)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkDdnsGenerateHostname(name string, networks []string, network1 string, network2 string, ddnsGenerateHostname string, useDdnsGenerateHostname string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_ddns_generate_hostname" {
    name = %q
    networks = %s
    ddns_generate_hostname = %q
	use_ddns_generate_hostname = %q
}
`, name, networksStr, ddnsGenerateHostname, useDdnsGenerateHostname)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkDdnsServerAlwaysUpdates(name string, networks []string, network1 string, network2 string, ddnsServerAlwaysUpdates string, ddnsUseOption81 string, useDdnsUseOption81 string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_ddns_server_always_updates" {
    name = %q
    networks = %s
    ddns_server_always_updates = %q
	ddns_use_option81 = %q
	use_ddns_use_option81 = %q
}
`, name, networksStr, ddnsServerAlwaysUpdates, ddnsUseOption81, useDdnsUseOption81)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkDdnsTtl(name string, networks []string, network1 string, network2 string, ddnsTtl string, useDdnsTtl string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_ddns_ttl" {
    name = %q
    networks = %s
    ddns_ttl = %q
	use_ddns_ttl = %q
}
`, name, networksStr, ddnsTtl, useDdnsTtl)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkDdnsUseOption81(name string, networks []string, network1 string, network2 string, ddnsUseOption81 string, useDdnsUseOption81 string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_ddns_use_option81" {
    name = %q
    networks = %s
    ddns_use_option81 = %q
	use_ddns_use_option81 = %q
}
`, name, networksStr, ddnsUseOption81, useDdnsUseOption81)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkDisable(name string, networks []string, network1 string, network2 string, disable string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_disable" {
    name = %q
    networks = %s
    disable = %q
}
`, name, networksStr, disable)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkDomainName(name string, networks []string, network1 string, network2 string, domainName string, useDomainName string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_domain_name" {
    name = %q
    networks = %s
    domain_name = %q
	use_domain_name = %q
}
`, name, networksStr, domainName, useDomainName)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkDomainNameServers(name string, networks []string, network1 string, network2 string, domainNameServers []string, useDomainNameServers string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	domainNameServersStr := utils.ConvertStringSliceToHCL(domainNameServers)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_domain_name_servers" {
    name = %q
    networks = %s
    domain_name_servers = %s
	use_domain_name_servers = %q
}
`, name, networksStr, domainNameServersStr, useDomainNameServers)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkEnableDdns(name string, networks []string, network1 string, network2 string, enableDdns string, useEnableDdns string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_enable_ddns" {
    name = %q
    networks = %s
    enable_ddns = %q
	use_enable_ddns = %q
}
`, name, networksStr, enableDdns, useEnableDdns)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkExtAttrs(name string, networks []string, network1 string, network2 string, extAttrs map[string]string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_extattrs" {
    name = %q
    networks = %s
    extattrs = %s
}
`, name, networksStr, extAttrsStr)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkLogicFilterRules(name string, networks []string, network1 string, network2 string, logicFilterRules []map[string]any, useLogicFilterRules string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	logicFilterRulesStr := utils.ConvertSliceOfMapsToHCL(logicFilterRules)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_logic_filter_rules" {
    name = %q
    networks = %s
    logic_filter_rules = %s
	use_logic_filter_rules = %q
}
`, name, networksStr, logicFilterRulesStr, useLogicFilterRules)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkName(name string, networks []string, network1 string, network2 string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_name" {
    name = %q
    networks = %s
}
`, name, networksStr)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkNetworkView(name string, networks []string, network1 string, network2 string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_network_view" {
    name = %q
    networks = %s
}
`, name, networksStr)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkNetworks(name string, networks []string, network1 string, network2 string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_networks" {
    name = %q
    networks = %s
}
`, name, networksStr)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkOptions(name string, networks []string, network1 string, network2 string, options []map[string]any, useOptions string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	optionsStr := utils.ConvertSliceOfMapsToHCL(options)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_options" {
    name = %q
    networks = %s
    options = %s
	use_options = %q
}
`, name, networksStr, optionsStr, useOptions)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkPreferredLifetime(name string, networks []string, network1 string, network2 string, preferredLifetime string, usePreferredLifetime string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_preferred_lifetime" {
    name = %q
    networks = %s
    preferred_lifetime = %q
	use_preferred_lifetime = %q
	valid_lifetime = 43200
	use_valid_lifetime = "true"
}
`, name, networksStr, preferredLifetime, usePreferredLifetime)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkUpdateDnsOnLeaseRenewal(name string, networks []string, network1 string, network2 string, updateDnsOnLeaseRenewal string, useUpdateDnsOnLeaseRenewal string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_update_dns_on_lease_renewal" {
    name = %q
    networks = %s
    update_dns_on_lease_renewal = %q
	use_update_dns_on_lease_renewal = %q
}
`, name, networksStr, updateDnsOnLeaseRenewal, useUpdateDnsOnLeaseRenewal)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkUseDdnsDomainname(name string, networks []string, network1 string, network2 string, ddnsDomainname string, useDdnsDomainname string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_ddns_domainname" {
    name = %q
    networks = %s
	ddns_domainname = %q
    use_ddns_domainname = %q
}
`, name, networksStr, ddnsDomainname, useDdnsDomainname)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkUseDdnsGenerateHostname(name string, networks []string, network1 string, network2 string, ddnsGenerateHostname string, useDdnsGenerateHostname string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_ddns_generate_hostname" {
    name = %q
    networks = %s
	ddns_generate_hostname = %q
    use_ddns_generate_hostname = %q
}
`, name, networksStr, ddnsGenerateHostname, useDdnsGenerateHostname)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkUseDdnsTtl(name string, networks []string, network1 string, network2 string, ddnsTtl string, useDdnsTtl string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_ddns_ttl" {
    name = %q
    networks = %s
	ddns_ttl = %q
    use_ddns_ttl = %q
}
`, name, networksStr, ddnsTtl, useDdnsTtl)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkUseDdnsUseOption81(name string, networks []string, network1 string, network2 string, useDdnsUseOption81 string, ddnsUseOption81 string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_ddns_use_option81" {
    name = %q
    networks = %s
    use_ddns_use_option81 = %q
	ddns_use_option81 = %q
}
`, name, networksStr, useDdnsUseOption81, ddnsUseOption81)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkUseDomainName(name string, networks []string, network1 string, network2 string, useDomainName string, domainName string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_domain_name" {
    name = %q
    networks = %s
    use_domain_name = %q
	domain_name = %q
}
`, name, networksStr, useDomainName, domainName)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkUseDomainNameServers(name string, networks []string, network1 string, network2 string, useDomainNameServers string, domainNameServers []string) string {
	domainNameServersStr := utils.ConvertStringSliceToHCL(domainNameServers)
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_domain_name_servers" {
    name = %q
    networks = %s
    use_domain_name_servers = %q
	domain_name_servers = %s
}
`, name, networksStr, useDomainNameServers, domainNameServersStr)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkUseEnableDdns(name string, networks []string, network1 string, network2 string, useEnableDdns string, enableDdns string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_enable_ddns" {
    name = %q
    networks = %s
    use_enable_ddns = %q
	enable_ddns = %q
}
`, name, networksStr, useEnableDdns, enableDdns)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkUseLogicFilterRules(name string, networks []string, network1 string, network2 string, useLogicFilterRules string, logicFilterRules []map[string]any) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	logicFilterRulesStr := utils.ConvertSliceOfMapsToHCL(logicFilterRules)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_logic_filter_rules" {
    name = %q
    networks = %s
    use_logic_filter_rules = %q
	logic_filter_rules = %s
}
`, name, networksStr, useLogicFilterRules, logicFilterRulesStr)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkUseOptions(name string, networks []string, network1 string, network2 string, useOptions string, options []map[string]any) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	optionsStr := utils.ConvertSliceOfMapsToHCL(options)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_options" {
    name = %q
    networks = %s
    use_options = %q
	options = %s
}
`, name, networksStr, useOptions, optionsStr)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkUsePreferredLifetime(name string, networks []string, network1 string, network2 string, usePreferredLifetime string, preferredLifetime string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_preferred_lifetime" {
    name = %q
    networks = %s
    use_preferred_lifetime = %q
	preferred_lifetime = %q
	valid_lifetime = 43200
	use_valid_lifetime = "true"
}
`, name, networksStr, usePreferredLifetime, preferredLifetime)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkUseUpdateDnsOnLeaseRenewal(name string, networks []string, network1 string, network2 string, useUpdateDnsOnLeaseRenewal string, updateDnsOnLeaseRenewal string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_update_dns_on_lease_renewal" {
    name = %q
    networks = %s
    use_update_dns_on_lease_renewal = %q
	update_dns_on_lease_renewal = %q
}
`, name, networksStr, useUpdateDnsOnLeaseRenewal, updateDnsOnLeaseRenewal)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkUseValidLifetime(name string, networks []string, network1 string, network2 string, validLifetime string, useValidLifetime string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_use_valid_lifetime" {
    name = %q
    networks = %s
	valid_lifetime = %q
    use_valid_lifetime = %q
}
`, name, networksStr, validLifetime, useValidLifetime)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccIpv6sharednetworkValidLifetime(name string, networks []string, network1 string, network2 string, validLifetime string, useValidLifetime string) string {
	networksStr := utils.ConvertStringSliceToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6sharednetwork" "test_valid_lifetime" {
    name = %q
    networks = %s
    valid_lifetime = %q
    use_valid_lifetime = %q
}
`, name, networksStr, validLifetime, useValidLifetime)
	return strings.Join([]string{testAccBaseWithwoIPv6Networks(network1, network2), config}, "")
}

func testAccBaseWithwoIPv6Networks(network1 string, network2 string) string {
	return fmt.Sprintf(`
 resource "nios_ipam_ipv6network" "test1" {
	network = %q
 }
 resource "nios_ipam_ipv6network" "test2" {
	network = %q
 }
 `, network1, network2)
}
