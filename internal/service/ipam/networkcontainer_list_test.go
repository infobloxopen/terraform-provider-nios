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

func TestAccNetworkcontainerList_basic(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test"
	var v ipam.Networkcontainer
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
				Config:                   testAccNetworkcontainerBasicConfig(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccNetworkcontainerListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_ipam_network_container.test", 1),
				},
			},
		},
	})
}

func TestAccNetworkcontainerList_Filters(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test"
	var v ipam.Networkcontainer
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
				Config:                   testAccNetworkcontainerBasicConfig(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					// TODO : Update with required fields to verify the object was created with expected values
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccNetworkcontainerListConfigFilters(network),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_ipam_network_container.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							// TODO : Update the ref prefix with the correct identifying object for the resource
							"ref": knownvalue.StringRegexp(regexp.MustCompile("networkcontainer/")),
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

func TestAccNetworkcontainerList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_extattrs"
	var v ipam.Networkcontainer
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
				Config: testAccNetworkcontainerExtAttrs(network, map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccNetworkcontainerListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_ipam_network_container.test", 1),
				},
			},
		},
	})
}

func testAccNetworkcontainerListBasicConfig() string {
	return `
list "nios_ipam_network_container" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccNetworkcontainerListConfigFilters(filterValue string) string {
	return fmt.Sprintf(`
list "nios_ipam_network_container" "test" {
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

func testAccNetworkcontainerListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_ipam_network_container" "test" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}
