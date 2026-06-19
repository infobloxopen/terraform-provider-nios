
package ipam_test

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

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccIpv6networkList_basic(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config: testAccIpv6networkBasicConfig(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
				),
			},
            // Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:  true,
				Config: testAccIpv6networkListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_ipam_ipv6network.test", 1),
				},
			},
		},
	})
}

func TestAccIpv6networkList_Filters(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config: testAccIpv6networkBasicConfig(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:  true,
				Config: testAccIpv6networkListConfigFilters(network),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_ipam_ipv6network.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
						    // TODO : Update the ref prefix with the correct identifying object for the resource
							"ref": knownvalue.StringRegexp(regexp.MustCompile("ipv6network/")),
						}),
						[]querycheck.KnownValueCheck{
						    // TODO : Add checks for required fields
							{
								Path:       tfjsonpath.New("network"),
								KnownValue: knownvalue.StringExact(network),
							},
						},
					),
				},
			},
		},
	})
}

func TestAccIpv6networkList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_extattrs"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	extAttrValue := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config: testAccIpv6networkExtAttrs(network, map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query: true,
				Config: testAccIpv6networkListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_ipam_ipv6network.test", 1),
				},
			},
		},
	})
}

func testAccIpv6networkListBasicConfig() string {
	return `
list "nios_ipam_ipv6network" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccIpv6networkListConfigFilters(filterValue string) string {
	return fmt.Sprintf(`
list "nios_ipam_ipv6network" "test" {
	provider = nios
	include_resource = true
	config {
		filters = {
			network =  %q
		}
	}
}
`, filterValue)
}

func testAccIpv6networkListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_ipam_ipv6network" "test" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}

