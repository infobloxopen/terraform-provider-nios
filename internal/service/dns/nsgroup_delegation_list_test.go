package dns_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/querycheck"
	"github.com/hashicorp/terraform-plugin-testing/querycheck/queryfilter"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccNsgroupDelegationList_basic(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_delegation.test"
	var v dns.NsgroupDelegation
	name := acctest.RandomName()
	delegateTo := []map[string]interface{}{
		{
			"name":    "delegate_to_ns_group",
			"address": "2.3.4.5",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccNsgroupDelegationBasicConfig(name, delegateTo),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccNsgroupDelegationListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_nsgroup_delegation.test", 1),
				},
			},
		},
	})
}

func TestAccNsgroupDelegationList_Filters(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_delegation.test"
	var v dns.NsgroupDelegation
	name := acctest.RandomName()
	delegateTo := []map[string]interface{}{
		{
			"name":    "delegate_to_ns_group",
			"address": "2.3.4.5",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccNsgroupDelegationBasicConfig(name, delegateTo),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccNsgroupDelegationListConfigFilters(name),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_nsgroup_delegation.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("nsgroup:delegation/")),
						}),
						[]querycheck.KnownValueCheck{
							{
								Path:       tfjsonpath.New("name"),
								KnownValue: knownvalue.StringExact(name),
							},
						},
					),
				},
			},
		},
	})
}

func TestAccNsgroupDelegationList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_delegation.test_extattrs"
	var v dns.NsgroupDelegation
	name := acctest.RandomName()
	delegateTo := []map[string]interface{}{
		{
			"name":    "delegate_to_ns_group",
			"address": "2.3.4.5",
		},
	}

	extAttrValue := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config: testAccNsgroupDelegationExtAttrs(name, delegateTo, map[string]any{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccNsgroupDelegationListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_nsgroup_delegation.test_extattrs", 1),
				},
			},
		},
	})
}

func testAccNsgroupDelegationListBasicConfig() string {
	return `
list "nios_dns_nsgroup_delegation" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccNsgroupDelegationListConfigFilters(filterValue string) string {
	return fmt.Sprintf(`
list "nios_dns_nsgroup_delegation" "test" {
	provider = nios
	include_resource = true
	config {
		filters = {
			name =  %q
		}
	}
}
`, filterValue)
}

func testAccNsgroupDelegationListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_dns_nsgroup_delegation" "test" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}
