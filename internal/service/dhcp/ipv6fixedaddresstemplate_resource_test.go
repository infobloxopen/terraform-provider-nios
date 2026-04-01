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

// Objects to be present in the grid for tests
// ipv6filteroption - example-ipv6-option-filter-1 , example-ipv6-option-filter-2

var readableAttributesForIpv6fixedaddresstemplate = "comment,domain_name,domain_name_servers,extattrs,logic_filter_rules,name,number_of_addresses,offset,options,preferred_lifetime,use_domain_name,use_domain_name_servers,use_logic_filter_rules,use_options,use_preferred_lifetime,use_valid_lifetime,valid_lifetime"

func TestAccIpv6fixedaddresstemplateResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddresstemplate.test"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "false"),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "43200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddresstemplateResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_ipv6fixedaddresstemplate.test"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6fixedaddresstemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6fixedaddresstemplateBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					testAccCheckIpv6fixedaddresstemplateDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccIpv6fixedaddresstemplateResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddresstemplate.test_comment"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateComment(name, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateComment(name, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddresstemplateResource_DomainName(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddresstemplate.test_domain_name"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateDomainName(name, "example.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name", "example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateDomainName(name, "example.org", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name", "example.org"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddresstemplateResource_DomainNameServers(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddresstemplate.test_domain_name_servers"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")
	domainNameServers := []string{"2001:4860:4860::8888", "2001:4860:4860::9999", "2001:4860:4860::8899"}
	updatedDomainNameServers := []string{"2001:4860:4860::8888", "2001:4860:4860::8844"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateDomainNameServers(name, domainNameServers, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.0", "2001:4860:4860::8888"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.1", "2001:4860:4860::9999"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.2", "2001:4860:4860::8899"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateDomainNameServers(name, updatedDomainNameServers, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.0", "2001:4860:4860::8888"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.1", "2001:4860:4860::8844"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddresstemplateResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddresstemplate.test_extattrs"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()
	extAttrs := map[string]any{
		"Site": extAttrValue1,
	}
	updatedExtAttrs := map[string]any{
		"Site": extAttrValue2,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateExtAttrs(name, extAttrs),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateExtAttrs(name, updatedExtAttrs),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddresstemplateResource_LogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddresstemplate.test_logic_filter_rules"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")
	logicFilterRules := []map[string]any{
		{
			"filter": "example-ipv6-option-filter-1",
			"type":   "Option",
		},
	}
	updatedLogicFilterRules := []map[string]any{
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
				Config: testAccIpv6fixedaddresstemplateLogicFilterRules(name, logicFilterRules, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.filter", "example-ipv6-option-filter-1"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.type", "Option"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateLogicFilterRules(name, updatedLogicFilterRules, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.filter", "example-ipv6-option-filter-2"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.type", "Option"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddresstemplateResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddresstemplate.test_name"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")
	updatedName := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateName(updatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddresstemplateResource_NumberOfAddresses(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddresstemplate.test_number_of_addresses"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateNumberOfAddresses(name, "10", "10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "number_of_addresses", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateNumberOfAddresses(name, "20", "10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "number_of_addresses", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddresstemplateResource_Offset(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddresstemplate.test_offset"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateOffset(name, "10", "20"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "offset", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateOffset(name, "15", "20"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "offset", "15"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddresstemplateResource_Options(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddresstemplate.test_options"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")
	options := []map[string]any{
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
	updatedOptions := []map[string]any{
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
				Config: testAccIpv6fixedaddresstemplateOptions(name, options, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
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
				Config: testAccIpv6fixedaddresstemplateOptions(name, updatedOptions, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
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

func TestAccIpv6fixedaddresstemplateResource_PreferredLifetime(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddresstemplate.test_preferred_lifetime"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplatePreferredLifetime(name, "200", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "200"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplatePreferredLifetime(name, "600", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "600"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddresstemplateResource_UseDomainName(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddresstemplate.test_use_domain_name"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateUseDomainName(name, "example.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateUseDomainName(name, "example.com", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddresstemplateResource_UseDomainNameServers(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddresstemplate.test_use_domain_name_servers"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")
	domainNameServers := []string{"2001:4860:4860::8888", "2001:4860:4860::8844"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateUseDomainNameServers(name, "true", domainNameServers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateUseDomainNameServers(name, "false", domainNameServers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddresstemplateResource_UseLogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddresstemplate.test_use_logic_filter_rules"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")
	logicFilterRules := []map[string]any{
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
				Config: testAccIpv6fixedaddresstemplateUseLogicFilterRules(name, "true", logicFilterRules),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateUseLogicFilterRules(name, "false", logicFilterRules),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddresstemplateResource_UseOptions(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddresstemplate.test_use_options"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")
	options := []map[string]any{
		{
			"name":  "domain-name",
			"num":   "15",
			"value": "example.com",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateUseOptions(name, "true", options),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateUseOptions(name, "false", options),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddresstemplateResource_UsePreferredLifetime(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddresstemplate.test_use_preferred_lifetime"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateUsePreferredLifetime(name, "true", "100"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateUsePreferredLifetime(name, "false", "100"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddresstemplateResource_UseValidLifetime(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddresstemplate.test_use_valid_lifetime"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateUseValidLifetime(name, "true", "200"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateUseValidLifetime(name, "false", "200"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6fixedaddresstemplateResource_ValidLifetime(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddresstemplate.test_valid_lifetime"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateValidLifetime(name, "200", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "200"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateValidLifetime(name, "400", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "400"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckIpv6fixedaddresstemplateExists(ctx context.Context, resourceName string, v *dhcp.Ipv6fixedaddresstemplate) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			Ipv6fixedaddresstemplateAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForIpv6fixedaddresstemplate).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetIpv6fixedaddresstemplateResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetIpv6fixedaddresstemplateResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckIpv6fixedaddresstemplateDestroy(ctx context.Context, v *dhcp.Ipv6fixedaddresstemplate) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			Ipv6fixedaddresstemplateAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForIpv6fixedaddresstemplate).
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

func testAccCheckIpv6fixedaddresstemplateDisappears(ctx context.Context, v *dhcp.Ipv6fixedaddresstemplate) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			Ipv6fixedaddresstemplateAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccIpv6fixedaddresstemplateBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test" {
    name = %q
}
`, name)
}

func testAccIpv6fixedaddresstemplateComment(name string, comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccIpv6fixedaddresstemplateDomainName(name string, domainName, useDomainName string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_domain_name" {
    name = %q
    domain_name = %q
	use_domain_name = %q
}
`, name, domainName, useDomainName)
}

func testAccIpv6fixedaddresstemplateDomainNameServers(name string, domainNameServers []string, useDomainNameServers string) string {
	domainNameServersStr := utils.ConvertStringSliceToHCL(domainNameServers)
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_domain_name_servers" {
    name = %q
    domain_name_servers = %s
    use_domain_name_servers = %q
}
`, name, domainNameServersStr, useDomainNameServers)
}

func testAccIpv6fixedaddresstemplateExtAttrs(name string, extAttrs map[string]any) string {
	extAttrsStr := utils.ConvertMapToHCL(extAttrs)
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_extattrs" {
    name = %q
    extattrs = %s
}
`, name, extAttrsStr)
}

func testAccIpv6fixedaddresstemplateLogicFilterRules(name string, logicFilterRules []map[string]any, useLogicFilterRules string) string {
	logicFilterRulesStr := utils.ConvertSliceOfMapsToHCL(logicFilterRules)
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_logic_filter_rules" {
    name = %q
    logic_filter_rules = %s
	use_logic_filter_rules = %q
}
`, name, logicFilterRulesStr, useLogicFilterRules)
}

func testAccIpv6fixedaddresstemplateName(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_name" {
    name = %q
}
`, name)
}

func testAccIpv6fixedaddresstemplateNumberOfAddresses(name string, numberOfAddresses, offset string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_number_of_addresses" {
    name = %q
    number_of_addresses = %q
	offset = %q
}
`, name, numberOfAddresses, offset)
}

func testAccIpv6fixedaddresstemplateOffset(name string, offset, numberOfAddresses string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_offset" {
    name = %q
    offset = %q
    number_of_addresses = %q
}
`, name, offset, numberOfAddresses)
}

func testAccIpv6fixedaddresstemplateOptions(name string, options []map[string]any, useOptions string) string {
	optionsStr := utils.ConvertSliceOfMapsToHCL(options)
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_options" {
    name = %q
    options = %s
	use_options = %q
}
`, name, optionsStr, useOptions)
}

func testAccIpv6fixedaddresstemplatePreferredLifetime(name string, preferredLifetime, usePreferredLifetime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_preferred_lifetime" {
    name = %q
    preferred_lifetime = %q
    use_preferred_lifetime = %q
	valid_lifetime = 43200
	use_valid_lifetime = true
}
`, name, preferredLifetime, usePreferredLifetime)
}

func testAccIpv6fixedaddresstemplateUseDomainName(name string, domainName, useDomainName string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_use_domain_name" {
    name = %q
	domain_name = %q
    use_domain_name = %q
}
`, name, domainName, useDomainName)
}

func testAccIpv6fixedaddresstemplateUseDomainNameServers(name string, useDomainNameServers string, domainNameServers []string) string {
	domainNameServerStr := utils.ConvertStringSliceToHCL(domainNameServers)
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_use_domain_name_servers" {
    name = %q
    use_domain_name_servers = %q
    domain_name_servers = %s
}
`, name, useDomainNameServers, domainNameServerStr)
}

func testAccIpv6fixedaddresstemplateUseLogicFilterRules(name string, useLogicFilterRules string, logicFilterRules []map[string]any) string {
	logicFilterRulesStr := utils.ConvertSliceOfMapsToHCL(logicFilterRules)
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_use_logic_filter_rules" {
    name = %q
    use_logic_filter_rules = %q
    logic_filter_rules = %s
}
`, name, useLogicFilterRules, logicFilterRulesStr)
}

func testAccIpv6fixedaddresstemplateUseOptions(name string, useOptions string, options []map[string]any) string {
	optionsStr := utils.ConvertSliceOfMapsToHCL(options)
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_use_options" {
    name = %q
    use_options = %q
    options = %s
}
`, name, useOptions, optionsStr)
}

func testAccIpv6fixedaddresstemplateUsePreferredLifetime(name string, usePreferredLifetime, preferredLifetime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_use_preferred_lifetime" {
    name = %q
    use_preferred_lifetime = %q
    preferred_lifetime = %q
	valid_lifetime = 43200
	use_valid_lifetime = true
}
`, name, usePreferredLifetime, preferredLifetime)
}

func testAccIpv6fixedaddresstemplateUseValidLifetime(name string, useValidLifetime, validLifetime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_use_valid_lifetime" {
    name = %q
    use_valid_lifetime = %q
    valid_lifetime = %q
}
`, name, useValidLifetime, validLifetime)
}

func testAccIpv6fixedaddresstemplateValidLifetime(name string, validLifetime, useValidLifetime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_valid_lifetime" {
    name = %q
    valid_lifetime = %q
    use_valid_lifetime = %q
}
`, name, validLifetime, useValidLifetime)
}
