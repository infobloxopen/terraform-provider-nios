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

func TestAccNetworkList_basic(t *testing.T) {
	var resourceName = "nios_ipam_network.test"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccNetworkBasicConfig(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccNetworkListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_ipam_network.test", 1),
				},
			},
		},
	})
}

func TestAccNetworkList_Filters(t *testing.T) {
	var resourceName = "nios_ipam_network.test"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccNetworkBasicConfig(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					// TODO : Update with required fields to verify the object was created with expected values
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccNetworkListConfigFilters(network),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_ipam_network.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("network/")),
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

func TestAccNetworkList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_ipam_network.test_extattrs"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

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
				Config: testAccNetworkExtAttrs(network, map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccNetworkListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_ipam_network.test", 1),
				},
			},
		},
	})
}

func testAccNetworkListBasicConfig() string {
	return `
list "nios_ipam_network" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccNetworkListConfigFilters(filterValue string) string {
	return fmt.Sprintf(`
list "nios_ipam_network" "test" {
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

func testAccNetworkListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_ipam_network" "test" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}
