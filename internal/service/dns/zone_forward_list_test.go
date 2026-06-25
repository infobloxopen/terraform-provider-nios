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

func TestAccZoneForwardList_basic(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test"
	var v dns.ZoneForward
	zoneFqdn := acctest.RandomNameWithPrefix("zone-forward") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccZoneForwardBasicConfig(zoneFqdn, "ensg1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccZoneForwardListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_zone_forward.test", 1),
				},
			},
		},
	})
}

func TestAccZoneForwardList_Filters(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test"
	var v dns.ZoneForward
	zoneFqdn := acctest.RandomNameWithPrefix("zone-forward") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccZoneForwardBasicConfig(zoneFqdn, "ensg1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", zoneFqdn),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccZoneForwardListConfigFilters(zoneFqdn),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_zone_forward.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("zone_forward/")),
						}),
						[]querycheck.KnownValueCheck{
							{
								Path:       tfjsonpath.New("fqdn"),
								KnownValue: knownvalue.StringExact(zoneFqdn),
							},
							{
								Path:       tfjsonpath.New("view"),
								KnownValue: knownvalue.StringExact("default"),
							},
						},
					),
				},
			},
		},
	})
}

func TestAccZoneForwardList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_extattrs"
	var v dns.ZoneForward
	zoneFqdn := acctest.RandomNameWithPrefix("zone-forward") + ".example.com"
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
				Config: testAccZoneForwardExtAttrs(zoneFqdn, "ensg1", map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccZoneForwardListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_zone_forward.test_extattrs", 1),
				},
			},
		},
	})
}

func testAccZoneForwardListBasicConfig() string {
	return `
list "nios_dns_zone_forward" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccZoneForwardListConfigFilters(fqdn string) string {
	return fmt.Sprintf(`
list "nios_dns_zone_forward" "test" {
	provider = nios
	include_resource = true
	config {
		filters = {
			fqdn = %q
		}
	}
}
`, fqdn)
}

func testAccZoneForwardListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_dns_zone_forward" "test" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}
