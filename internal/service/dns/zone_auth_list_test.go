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

func TestAccZoneAuthList_basic(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("auth-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccZoneAuthBasicConfig(zoneFqdn, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccZoneAuthListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_zone_auth.test", 1),
				},
			},
		},
	})
}

func TestAccZoneAuthList_Filters(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("auth-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccZoneAuthBasicConfig(zoneFqdn, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", zoneFqdn),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccZoneAuthListConfigFilters(zoneFqdn),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_zone_auth.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("zone_auth/")),
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

func TestAccZoneAuthList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_extattrs"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("auth-zone") + ".com"
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
				Config: testAccZoneAuthExtAttrs(zoneFqdn, "default", map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccZoneAuthListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_zone_auth.test", 1),
				},
			},
		},
	})
}

func testAccZoneAuthListBasicConfig() string {
	return `
list "nios_dns_zone_auth" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccZoneAuthListConfigFilters(fqdn string) string {
	return fmt.Sprintf(`
list "nios_dns_zone_auth" "test" {
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

func testAccZoneAuthListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_dns_zone_auth" "test" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}
