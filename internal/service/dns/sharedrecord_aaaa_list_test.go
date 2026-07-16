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

func TestAccSharedrecordAaaaList_basic(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_aaaa.test"
	var v dns.SharedrecordAaaa
	name := acctest.RandomNameWithPrefix("sharedrecord-aaaa")
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccSharedrecordAaaaBasicConfig(name, "2001:db8::1", sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAaaaExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccSharedrecordAaaaListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_sharedrecord_aaaa.test", 1),
				},
			},
		},
	})
}

func TestAccSharedrecordAaaaList_Filters(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_aaaa.test"
	var v dns.SharedrecordAaaa
	name := acctest.RandomNameWithPrefix("sharedrecord-aaaa")
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccSharedrecordAaaaBasicConfig(name, "2001:db8::1", sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccSharedrecordAaaaListConfigFilters(name),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_sharedrecord_aaaa.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("sharedrecord:aaaa/")),
						}),
						[]querycheck.KnownValueCheck{
							{
								Path:       tfjsonpath.New("name"),
								KnownValue: knownvalue.StringExact(name),
							},
							{
								Path:       tfjsonpath.New("ipv6addr"),
								KnownValue: knownvalue.StringExact("2001:db8::1"),
							},
						},
					),
				},
			},
		},
	})
}

func TestAccSharedrecordAaaaList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_aaaa.test_extattrs"
	var v dns.SharedrecordAaaa

	extAttrValue := acctest.RandomName()
	name := acctest.RandomNameWithPrefix("sharedrecord-aaaa") + ".example.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config: testAccSharedrecordAaaaExtAttrs(name, "2001:db8::1", sharedRecordGroup, map[string]any{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccSharedrecordAaaaListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_sharedrecord_aaaa.test_extattrs", 1),
				},
			},
		},
	})
}

func testAccSharedrecordAaaaListBasicConfig() string {
	return `
list "nios_dns_sharedrecord_aaaa" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccSharedrecordAaaaListConfigFilters(filterValue string) string {
	return fmt.Sprintf(`
list "nios_dns_sharedrecord_aaaa" "test" {
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

func testAccSharedrecordAaaaListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_dns_sharedrecord_aaaa" "test_extattrs" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}
