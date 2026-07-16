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

func TestAccZoneDelegatedList_basic(t *testing.T) {
	var resourceName = "nios_dns_zone_delegated.test"
	var v dns.ZoneDelegated
	fqdn := acctest.RandomNameWithPrefix("zone-delegated") + ".example.com"
	delegatedToName := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccZoneDelegatedBasicConfig(fqdn, delegatedToName, "10.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccZoneDelegatedListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_zone_delegated.test", 1),
				},
			},
		},
	})
}

func TestAccZoneDelegatedList_Filters(t *testing.T) {
	var resourceName = "nios_dns_zone_delegated.test"
	var v dns.ZoneDelegated
	fqdn := acctest.RandomNameWithPrefix("zone-delegated") + ".example.com"
	delegatedToName := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccZoneDelegatedBasicConfig(fqdn, delegatedToName, "10.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", fqdn),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccZoneDelegatedListConfigFilters(fqdn),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_zone_delegated.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("zone_delegated/")),
						}),
						[]querycheck.KnownValueCheck{
							{
								Path:       tfjsonpath.New("fqdn"),
								KnownValue: knownvalue.StringExact(fqdn),
							},
						},
					),
				},
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneDelegatedList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dns_zone_delegated.test_extattrs"
	var v dns.ZoneDelegated
	fqdn := acctest.RandomNameWithPrefix("zone-delegated") + ".example.com"
	delegatedToName := acctest.RandomName()
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
				Config: testAccZoneDelegatedExtAttrs(fqdn, delegatedToName, "10.0.0.1", map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccZoneDelegatedListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_zone_delegated.test", 1),
				},
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccZoneDelegatedListBasicConfig() string {
	return `
list "nios_dns_zone_delegated" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccZoneDelegatedListConfigFilters(filterValue string) string {
	return fmt.Sprintf(`
list "nios_dns_zone_delegated" "test" {
	provider = nios
	include_resource = true
	config {
		filters = {
			fqdn =  %q
		}
	}
}
`, filterValue)
}

func testAccZoneDelegatedListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_dns_zone_delegated" "test" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}
