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
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
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

func TestAccIpv6fixedaddresstemplateResource_Ref(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6fixedaddresstemplate.test_ref"
	var v dhcp.Ipv6fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("ipv6-fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateRef(name, "REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateRef(name, "REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
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
				Config: testAccIpv6fixedaddresstemplateComment(name, "COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateComment(name, "COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
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
				Config: testAccIpv6fixedaddresstemplateDomainName(name, "DOMAIN_NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name", "DOMAIN_NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateDomainName(name, "DOMAIN_NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name", "DOMAIN_NAME_UPDATE_REPLACE_ME"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateDomainNameServers(name, "DOMAIN_NAME_SERVERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers", "DOMAIN_NAME_SERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateDomainNameServers(name, "DOMAIN_NAME_SERVERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers", "DOMAIN_NAME_SERVERS_UPDATE_REPLACE_ME"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateExtAttrs(name, "EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateExtAttrs(name, "EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateLogicFilterRules(name, "LOGIC_FILTER_RULES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateLogicFilterRules(name, "LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
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
				Config: testAccIpv6fixedaddresstemplateName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
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
				Config: testAccIpv6fixedaddresstemplateNumberOfAddresses(name, "NUMBER_OF_ADDRESSES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "number_of_addresses", "NUMBER_OF_ADDRESSES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateNumberOfAddresses(name, "NUMBER_OF_ADDRESSES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "number_of_addresses", "NUMBER_OF_ADDRESSES_UPDATE_REPLACE_ME"),
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
				Config: testAccIpv6fixedaddresstemplateOffset(name, "OFFSET_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "offset", "OFFSET_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateOffset(name, "OFFSET_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "offset", "OFFSET_UPDATE_REPLACE_ME"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateOptions(name, "OPTIONS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateOptions(name, "OPTIONS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_UPDATE_REPLACE_ME"),
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
				Config: testAccIpv6fixedaddresstemplatePreferredLifetime(name, "PREFERRED_LIFETIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "PREFERRED_LIFETIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplatePreferredLifetime(name, "PREFERRED_LIFETIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "PREFERRED_LIFETIME_UPDATE_REPLACE_ME"),
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
				Config: testAccIpv6fixedaddresstemplateUseDomainName(name, "USE_DOMAIN_NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name", "USE_DOMAIN_NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateUseDomainName(name, "USE_DOMAIN_NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name", "USE_DOMAIN_NAME_UPDATE_REPLACE_ME"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateUseDomainNameServers(name, "USE_DOMAIN_NAME_SERVERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "USE_DOMAIN_NAME_SERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateUseDomainNameServers(name, "USE_DOMAIN_NAME_SERVERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "USE_DOMAIN_NAME_SERVERS_UPDATE_REPLACE_ME"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateUseLogicFilterRules(name, "USE_LOGIC_FILTER_RULES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "USE_LOGIC_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateUseLogicFilterRules(name, "USE_LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "USE_LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6fixedaddresstemplateUseOptions(name, "USE_OPTIONS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "USE_OPTIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateUseOptions(name, "USE_OPTIONS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "USE_OPTIONS_UPDATE_REPLACE_ME"),
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
				Config: testAccIpv6fixedaddresstemplateUsePreferredLifetime(name, "USE_PREFERRED_LIFETIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "USE_PREFERRED_LIFETIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateUsePreferredLifetime(name, "USE_PREFERRED_LIFETIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "USE_PREFERRED_LIFETIME_UPDATE_REPLACE_ME"),
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
				Config: testAccIpv6fixedaddresstemplateUseValidLifetime(name, "USE_VALID_LIFETIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "USE_VALID_LIFETIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateUseValidLifetime(name, "USE_VALID_LIFETIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "USE_VALID_LIFETIME_UPDATE_REPLACE_ME"),
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
				Config: testAccIpv6fixedaddresstemplateValidLifetime(name, "VALID_LIFETIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "VALID_LIFETIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6fixedaddresstemplateValidLifetime(name, "VALID_LIFETIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6fixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "VALID_LIFETIME_UPDATE_REPLACE_ME"),
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
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test" {
    name = %q
}
`, name)
}

func testAccIpv6fixedaddresstemplateRef(name string, ref string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_ref" {
    name = %q
    ref = %q
}
`, name, ref)
}

func testAccIpv6fixedaddresstemplateComment(name string, comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccIpv6fixedaddresstemplateDomainName(name string, domainName string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_domain_name" {
    name = %q
    domain_name = %q
}
`, name, domainName)
}

func testAccIpv6fixedaddresstemplateDomainNameServers(name string, domainNameServers string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_domain_name_servers" {
    name = %q
    domain_name_servers = %q
}
`, name, domainNameServers)
}

func testAccIpv6fixedaddresstemplateExtAttrs(name string, extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_extattrs" {
    name = %q
    extattrs = %q
}
`, name, extAttrs)
}

func testAccIpv6fixedaddresstemplateLogicFilterRules(name string, logicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_logic_filter_rules" {
    name = %q
    logic_filter_rules = %q
}
`, name, logicFilterRules)
}

func testAccIpv6fixedaddresstemplateName(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_name" {
    name = %q
}
`, name)
}

func testAccIpv6fixedaddresstemplateNumberOfAddresses(name string, numberOfAddresses string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_number_of_addresses" {
    name = %q
    number_of_addresses = %q
}
`, name, numberOfAddresses)
}

func testAccIpv6fixedaddresstemplateOffset(name string, offset string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_offset" {
    name = %q
    offset = %q
}
`, name, offset)
}

func testAccIpv6fixedaddresstemplateOptions(name string, options string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_options" {
    name = %q
    options = %q
}
`, name, options)
}

func testAccIpv6fixedaddresstemplatePreferredLifetime(name string, preferredLifetime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_preferred_lifetime" {
    name = %q
    preferred_lifetime = %q
}
`, name, preferredLifetime)
}

func testAccIpv6fixedaddresstemplateUseDomainName(name string, useDomainName string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_use_domain_name" {
    name = %q
    use_domain_name = %q
}
`, name, useDomainName)
}

func testAccIpv6fixedaddresstemplateUseDomainNameServers(name string, useDomainNameServers string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_use_domain_name_servers" {
    name = %q
    use_domain_name_servers = %q
}
`, name, useDomainNameServers)
}

func testAccIpv6fixedaddresstemplateUseLogicFilterRules(name string, useLogicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_use_logic_filter_rules" {
    name = %q
    use_logic_filter_rules = %q
}
`, name, useLogicFilterRules)
}

func testAccIpv6fixedaddresstemplateUseOptions(name string, useOptions string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_use_options" {
    name = %q
    use_options = %q
}
`, name, useOptions)
}

func testAccIpv6fixedaddresstemplateUsePreferredLifetime(name string, usePreferredLifetime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_use_preferred_lifetime" {
    name = %q
    use_preferred_lifetime = %q
}
`, name, usePreferredLifetime)
}

func testAccIpv6fixedaddresstemplateUseValidLifetime(name string, useValidLifetime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_use_valid_lifetime" {
    name = %q
    use_valid_lifetime = %q
}
`, name, useValidLifetime)
}

func testAccIpv6fixedaddresstemplateValidLifetime(name string, validLifetime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6fixedaddresstemplate" "test_valid_lifetime" {
    name = %q
    valid_lifetime = %q
}
`, name, validLifetime)
}
