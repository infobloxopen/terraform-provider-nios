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

func TestAccZoneStubList_basic(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test"
	var v dns.ZoneStub
	zoneFqdn := acctest.RandomNameWithPrefix("zone-stub") + ".com"
	stubServerName := acctest.RandomNameWithPrefix("stub_server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccZoneStubBasicConfig(zoneFqdn, "1.1.1.1", stubServerName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccZoneStubListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_zone_stub.test", 1),
				},
			},
		},
	})
}

func TestAccZoneStubList_Filters(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test"
	var v dns.ZoneStub
	zoneFqdn := acctest.RandomNameWithPrefix("zone-stub") + ".com"
	stubServerName := acctest.RandomNameWithPrefix("stub_server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccZoneStubBasicConfig(zoneFqdn, "1.1.1.1", stubServerName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", zoneFqdn),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccZoneStubListConfigFilters(zoneFqdn),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_zone_stub.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("zone_stub/")),
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

func TestAccZoneStubList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_extattrs"
	var v dns.ZoneStub
	zoneFqdn := acctest.RandomNameWithPrefix("zone-stub") + ".com"
	stubServerName := acctest.RandomNameWithPrefix("stub_server")
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
				Config: testAccZoneStubExtAttrs(zoneFqdn, "1.1.1.1", stubServerName, map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccZoneStubListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_zone_stub.test", 1),
				},
			},
		},
	})
}

func testAccZoneStubListBasicConfig() string {
	return `
list "nios_dns_zone_stub" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccZoneStubListConfigFilters(fqdn string) string {
	return fmt.Sprintf(`
list "nios_dns_zone_stub" "test" {
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

func testAccZoneStubListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_dns_zone_stub" "test" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}
