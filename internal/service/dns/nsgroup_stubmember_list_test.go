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

func TestAccNsgroupStubmemberList_basic(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_stubmember.test"
	var v dns.NsgroupStubmember
	name := acctest.RandomNameWithPrefix("test-nsgroup-stubmember")
	memberName := utils.GetNIOSGridMemberHostName()
	stubMember := []map[string]any{
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
				Config:                   testAccNsgroupStubmemberBasicConfig(name, stubMember),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupStubmemberExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccNsgroupStubmemberListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_nsgroup_stubmember.test", 1),
				},
			},
		},
	})
}

func TestAccNsgroupStubmemberList_Filters(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_stubmember.test"
	var v dns.NsgroupStubmember
	name := acctest.RandomNameWithPrefix("test-nsgroup-stubmember")
	memberName := utils.GetNIOSGridMemberHostName()
	stubMember := []map[string]any{
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
				Config:                   testAccNsgroupStubmemberBasicConfig(name, stubMember),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupStubmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccNsgroupStubmemberListConfigFilters(name),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_nsgroup_stubmember.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("nsgroup:stubmember/")),
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

func TestAccNsgroupStubmemberList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_stubmember.test_extattrs"
	var v dns.NsgroupStubmember
	name := acctest.RandomNameWithPrefix("test-nsgroup-stubmember")
	memberName := utils.GetNIOSGridMemberHostName()
	stubMember := []map[string]any{
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
				Config: testAccNsgroupStubmemberExtAttrs(name, stubMember, map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupStubmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccNsgroupStubmemberListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_nsgroup_stubmember.test_extattrs", 1),
				},
			},
		},
	})
}

func testAccNsgroupStubmemberListBasicConfig() string {
	return `
list "nios_dns_nsgroup_stubmember" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccNsgroupStubmemberListConfigFilters(filterValue string) string {
	return fmt.Sprintf(`
list "nios_dns_nsgroup_stubmember" "test" {
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

func testAccNsgroupStubmemberListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_dns_nsgroup_stubmember" "test_extattrs" {
	provider = nios
	config {
		extattrfilters = {
			Site = %q
		}
	}
}
`, extAttrsValue)
}
