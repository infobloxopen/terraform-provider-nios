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
// filteroption - example-option-filter-1 , example-option-filter-2

var readableAttributesForFixedaddresstemplate = "bootfile,bootserver,comment,ddns_domainname,ddns_hostname,deny_bootp,enable_ddns,enable_pxe_lease_time,extattrs,ignore_dhcp_option_list_request,logic_filter_rules,name,nextserver,number_of_addresses,offset,options,pxe_lease_time,use_bootfile,use_bootserver,use_ddns_domainname,use_deny_bootp,use_enable_ddns,use_ignore_dhcp_option_list_request,use_logic_filter_rules,use_nextserver,use_options,use_pxe_lease_time"

func TestAccFixedaddresstemplateResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "extattrs.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_fixedaddresstemplate.test"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFixedaddresstemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFixedaddresstemplateBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					testAccCheckFixedaddresstemplateDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccFixedaddresstemplateResource_Bootfile(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_bootfile"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateBootfile(name, "bootfile-name", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "bootfile-name"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateBootfile(name, "bootfile-name-updated", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "bootfile-name-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_Bootserver(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_bootserver"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateBootserver(name, "10.0.0.0", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "10.0.0.0"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateBootserver(name, "10.1.1.1", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "10.1.1.1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_comment"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateComment(name, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateComment(name, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_DdnsDomainname(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_ddns_domainname"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateDdnsDomainname(name, "example.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateDdnsDomainname(name, "example.org", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "example.org"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_DdnsHostname(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_ddns_hostname"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateDdnsHostname(name, "ddns_host.name"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_hostname", "ddns_host.name"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateDdnsHostname(name, "updated_ddns_host.name"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_hostname", "updated_ddns_host.name"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_DenyBootp(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_deny_bootp"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateDenyBootp(name, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateDenyBootp(name, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_EnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_enable_ddns"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateEnableDdns(name, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateEnableDdns(name, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_EnablePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_enable_pxe_lease_time"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateEnablePxeLeaseTime(name, "true", "100", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateEnablePxeLeaseTime(name, "false", "200", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_extattrs"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")
	extAttr1 := acctest.RandomName()
	extAttr2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateExtAttrs(name, map[string]any{
					"Site": extAttr1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttr1),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateExtAttrs(name, map[string]any{
					"Site": extAttr2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttr2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_IgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_ignore_dhcp_option_list_request"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateIgnoreDhcpOptionListRequest(name, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateIgnoreDhcpOptionListRequest(name, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_LogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_logic_filter_rules"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")
	logicFilterRules := []map[string]any{
		{
			"filter": "example-option-filter-1",
			"type":   "Option",
		},
	}
	updatedLogicFilterRules := []map[string]any{
		{
			"filter": "example-option-filter-2",
			"type":   "Option",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateLogicFilterRules(name, logicFilterRules, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.filter", "example-option-filter-1"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.type", "Option"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateLogicFilterRules(name, updatedLogicFilterRules, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.filter", "example-option-filter-2"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.type", "Option"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_name"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")
	updatedName := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateName(updatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_Nextserver(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_nextserver"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateNextserver(name, "10.0.0.0", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", "10.0.0.0"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateNextserver(name, "10.1.1.1", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", "10.1.1.1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_NumberOfAddresses(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_number_of_addresses"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateNumberOfAddresses(name, "3", "2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "number_of_addresses", "3"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateNumberOfAddresses(name, "5", "2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "number_of_addresses", "5"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_Offset(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_offset"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateOffset(name, "3", "10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "offset", "3"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateOffset(name, "2", "10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "offset", "2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_Options(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_options"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")
	options := []map[string]any{
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

	updatedOptions := []map[string]any{
		{
			"num":   51,
			"value": "7200",
		},
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
				Config: testAccFixedaddresstemplateOptions(name, options, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
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
				Config: testAccFixedaddresstemplateOptions(name, updatedOptions, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "options.0.num", "51"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "7200"),
					resource.TestCheckResourceAttr(resourceName, "options.1.name", "subnet-mask"),
					resource.TestCheckResourceAttr(resourceName, "options.1.value", "1.1.1.1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_PxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_pxe_lease_time"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplatePxeLeaseTime(name, "10", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplatePxeLeaseTime(name, "20", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_UseBootfile(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_use_bootfile"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateUseBootfile(name, "true", "boot-file"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateUseBootfile(name, "false", "boot-file"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_UseBootserver(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_use_bootserver"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateUseBootserver(name, "true", "10.0.0.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateUseBootserver(name, "false", "10.0.0.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_UseDdnsDomainname(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_use_ddns_domainname"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateUseDdnsDomainname(name, "true", "example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateUseDdnsDomainname(name, "false", "example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_UseDenyBootp(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_use_deny_bootp"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateUseDenyBootp(name, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateUseDenyBootp(name, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_UseEnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_use_enable_ddns"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateUseEnableDdns(name, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateUseEnableDdns(name, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_UseIgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_use_ignore_dhcp_option_list_request"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateUseIgnoreDhcpOptionListRequest(name, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateUseIgnoreDhcpOptionListRequest(name, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_UseLogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_use_logic_filter_rules"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")
	logicFilterRules := []map[string]any{
		{
			"filter": "example-option-filter-1",
			"type":   "Option",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateUseLogicFilterRules(name, "true", logicFilterRules),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateUseLogicFilterRules(name, "false", logicFilterRules),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_UseNextserver(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_use_nextserver"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateUseNextserver(name, "true", "10.0.0.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateUseNextserver(name, "false", "10.0.0.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_UseOptions(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_use_options"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")
	options := []map[string]any{
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateUseOptions(name, "true", options),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateUseOptions(name, "false", options),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddresstemplateResource_UsePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_fixedaddresstemplate.test_use_pxe_lease_time"
	var v dhcp.Fixedaddresstemplate
	name := acctest.RandomNameWithPrefix("fixedaddress-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFixedaddresstemplateUsePxeLeaseTime(name, "true", "100"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFixedaddresstemplateUsePxeLeaseTime(name, "false", "100"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddresstemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckFixedaddresstemplateExists(ctx context.Context, resourceName string, v *dhcp.Fixedaddresstemplate) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			FixedaddresstemplateAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForFixedaddresstemplate).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetFixedaddresstemplateResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetFixedaddresstemplateResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckFixedaddresstemplateDestroy(ctx context.Context, v *dhcp.Fixedaddresstemplate) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			FixedaddresstemplateAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForFixedaddresstemplate).
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

func testAccCheckFixedaddresstemplateDisappears(ctx context.Context, v *dhcp.Fixedaddresstemplate) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			FixedaddresstemplateAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccFixedaddresstemplateBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test" {
    name = %q
}
`, name)
}

func testAccFixedaddresstemplateBootfile(name string, bootfile, useBootfile string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_bootfile" {
    name = %q
    bootfile = %q
    use_bootfile = %q
}
`, name, bootfile, useBootfile)
}

func testAccFixedaddresstemplateBootserver(name string, bootserver, useBootserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_bootserver" {
    name = %q
    bootserver = %q
    use_bootserver = %q
}
`, name, bootserver, useBootserver)
}

func testAccFixedaddresstemplateComment(name string, comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccFixedaddresstemplateDdnsDomainname(name string, ddnsDomainname string, useDdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_ddns_domainname" {
    name = %q
    ddns_domainname = %q
	use_ddns_domainname = %q
}
`, name, ddnsDomainname, useDdnsDomainname)
}

func testAccFixedaddresstemplateDdnsHostname(name string, ddnsHostname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_ddns_hostname" {
    name = %q
    ddns_hostname = %q
}
`, name, ddnsHostname)
}

func testAccFixedaddresstemplateDenyBootp(name string, denyBootp string, useDenyBootp string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_deny_bootp" {
    name = %q
    deny_bootp = %q
	use_deny_bootp = %q
}
`, name, denyBootp, useDenyBootp)
}

func testAccFixedaddresstemplateEnableDdns(name string, enableDdns string, useEnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_enable_ddns" {
    name = %q
    enable_ddns = %q
	use_enable_ddns = %q
}
`, name, enableDdns, useEnableDdns)
}

func testAccFixedaddresstemplateEnablePxeLeaseTime(name string, enablePxeLeaseTime string, pxeLeaseTime string, usePxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_enable_pxe_lease_time" {
    name = %q
    enable_pxe_lease_time = %q
    pxe_lease_time = %q
    use_pxe_lease_time = %q
}
`, name, enablePxeLeaseTime, pxeLeaseTime, usePxeLeaseTime)
}

func testAccFixedaddresstemplateExtAttrs(name string, extAttrs map[string]any) string {
	extAttrStr := utils.ConvertMapToHCL(extAttrs)
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_extattrs" {
    name = %q
    extattrs = %s
}
`, name, extAttrStr)
}

func testAccFixedaddresstemplateIgnoreDhcpOptionListRequest(name string, ignoreDhcpOptionListRequest string, useIgnoreDhcpOptionListRequest string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_ignore_dhcp_option_list_request" {
    name = %q
    ignore_dhcp_option_list_request = %q
    use_ignore_dhcp_option_list_request = %q
}
`, name, ignoreDhcpOptionListRequest, useIgnoreDhcpOptionListRequest)
}

func testAccFixedaddresstemplateLogicFilterRules(name string, logicFilterRules []map[string]any, useLogicFilterRules string) string {
	logicFilterRulesStr := utils.ConvertSliceOfMapsToHCL(logicFilterRules)
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_logic_filter_rules" {
    name = %q
    logic_filter_rules = %s
	use_logic_filter_rules = %q
}
`, name, logicFilterRulesStr, useLogicFilterRules)
}

func testAccFixedaddresstemplateName(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_name" {
    name = %q
}
`, name)
}

func testAccFixedaddresstemplateNextserver(name string, nextserver string, useNextserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_nextserver" {
    name = %q
    nextserver = %q
    use_nextserver = %q
}
`, name, nextserver, useNextserver)
}

func testAccFixedaddresstemplateNumberOfAddresses(name string, numberOfAddresses string, offset string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_number_of_addresses" {
    name = %q
    number_of_addresses = %q
    offset = %q
}
`, name, numberOfAddresses, offset)
}

func testAccFixedaddresstemplateOffset(name string, offset string, numberOfAddresses string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_offset" {
    name = %q
    offset = %q
    number_of_addresses = %q
}
`, name, offset, numberOfAddresses)
}

func testAccFixedaddresstemplateOptions(name string, options []map[string]any, useOptions string) string {
	optionsStr := utils.ConvertSliceOfMapsToHCL(options)
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_options" {
    name = %q
    options = %s
    use_options = %q
}
`, name, optionsStr, useOptions)
}

func testAccFixedaddresstemplatePxeLeaseTime(name string, pxeLeaseTime string, usePxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_pxe_lease_time" {
    name = %q
    pxe_lease_time = %q
    use_pxe_lease_time = %q
}
`, name, pxeLeaseTime, usePxeLeaseTime)
}

func testAccFixedaddresstemplateUseBootfile(name string, useBootfile string, bootFile string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_use_bootfile" {
    name = %q
    use_bootfile = %q
    bootfile = %q
}
`, name, useBootfile, bootFile)
}

func testAccFixedaddresstemplateUseBootserver(name string, useBootserver string, bootServer string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_use_bootserver" {
    name = %q
    use_bootserver = %q
    bootserver = %q
}
`, name, useBootserver, bootServer)
}

func testAccFixedaddresstemplateUseDdnsDomainname(name string, useDdnsDomainname string, ddnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_use_ddns_domainname" {
    name = %q
    use_ddns_domainname = %q
    ddns_domainname = %q
}
`, name, useDdnsDomainname, ddnsDomainname)
}

func testAccFixedaddresstemplateUseDenyBootp(name string, useDenyBootp string, denyBootp string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_use_deny_bootp" {
    name = %q
    use_deny_bootp = %q
    deny_bootp = %q
}
`, name, useDenyBootp, denyBootp)
}

func testAccFixedaddresstemplateUseEnableDdns(name string, useEnableDdns string, enableDdns string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_use_enable_ddns" {
    name = %q
    use_enable_ddns = %q
    enable_ddns = %q
}
`, name, useEnableDdns, enableDdns)
}

func testAccFixedaddresstemplateUseIgnoreDhcpOptionListRequest(name string, useIgnoreDhcpOptionListRequest string, ignoreDhcpOptionListRequest string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_use_ignore_dhcp_option_list_request" {
    name = %q
    use_ignore_dhcp_option_list_request = %q
    ignore_dhcp_option_list_request = %q
}
`, name, useIgnoreDhcpOptionListRequest, ignoreDhcpOptionListRequest)
}

func testAccFixedaddresstemplateUseLogicFilterRules(name string, useLogicFilterRules string, logicFilterRules []map[string]any) string {
	logicFilterRulesStr := utils.ConvertSliceOfMapsToHCL(logicFilterRules)
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_use_logic_filter_rules" {
    name = %q
    use_logic_filter_rules = %q
    logic_filter_rules = %s
}
`, name, useLogicFilterRules, logicFilterRulesStr)
}

func testAccFixedaddresstemplateUseNextserver(name string, useNextserver string, nextServer string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_use_nextserver" {
    name = %q
    use_nextserver = %q
    nextserver = %q
}
`, name, useNextserver, nextServer)
}

func testAccFixedaddresstemplateUseOptions(name string, useOptions string, options []map[string]any) string {
	optionsStr := utils.ConvertSliceOfMapsToHCL(options)
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_use_options" {
    name = %q
    use_options = %q
    options = %s
}
`, name, useOptions, optionsStr)
}

func testAccFixedaddresstemplateUsePxeLeaseTime(name string, usePxeLeaseTime string, pxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fixedaddresstemplate" "test_use_pxe_lease_time" {
    name = %q
    use_pxe_lease_time = %q
    pxe_lease_time = %q
}
`, name, usePxeLeaseTime, pxeLeaseTime)
}
