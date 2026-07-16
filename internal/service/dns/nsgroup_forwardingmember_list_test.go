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
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

func TestAccNsgroupForwardingmemberList_basic(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardingmember.test"
	var v dns.NsgroupForwardingmember
	name := acctest.RandomNameWithPrefix("ns-group-forwardingMember")
	memberName := utils.GetNIOSGridMasterHostName()
	forwardingServers := []map[string]any{
		{
			"name": memberName,
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
				Config:                   testAccNsgroupForwardingmemberBasicConfig(name, forwardingServers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccNsgroupForwardingmemberListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_nsgroup_forwardingmember.test", 1),
				},
			},
		},
	})
}

func TestAccNsgroupForwardingmemberList_Filters(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardingmember.test"
	var v dns.NsgroupForwardingmember
	name := acctest.RandomNameWithPrefix("ns-group-forwardingMember")
	memberName := utils.GetNIOSGridMasterHostName()
	forwardingServers := []map[string]any{
		{
			"name": memberName,
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
				Config:                   testAccNsgroupForwardingmemberBasicConfig(name, forwardingServers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccNsgroupForwardingmemberListConfigFilters(name),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_nsgroup_forwardingmember.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("nsgroup:forwardingmember/")),
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

func TestAccNsgroupForwardingmemberList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardingmember.test_extattrs"
	var v dns.NsgroupForwardingmember
	name := acctest.RandomNameWithPrefix("ns-group-forwardingMember")
	memberName := utils.GetNIOSGridMasterHostName()
	forwardingServers := []map[string]any{
		{
			"name": memberName,
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
				Config: testAccNsgroupForwardingmemberExtAttrs(name, forwardingServers, map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccNsgroupForwardingmemberListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_nsgroup_forwardingmember.test_extattrs", 1),
				},
			},
		},
	})
}

func testAccNsgroupForwardingmemberListBasicConfig() string {
	return `
list "nios_dns_nsgroup_forwardingmember" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccNsgroupForwardingmemberListConfigFilters(filterValue string) string {
	return fmt.Sprintf(`
list "nios_dns_nsgroup_forwardingmember" "test" {
	provider = nios
	include_resource = true
	config {
		filters = {
			name = %q
		}
	}
}
`, filterValue)
}

func testAccNsgroupForwardingmemberListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_dns_nsgroup_forwardingmember" "test_extattrs" {
	provider = nios
	config {
		extattrfilters = {
			Site = %q
		}
	}
}
`, extAttrsValue)
}
