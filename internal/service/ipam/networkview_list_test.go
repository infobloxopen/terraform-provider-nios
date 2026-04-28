package ipam_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/querycheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

func TestAccNetworkviewList_basic(t *testing.T) {
	var resourceName = "nios_ipam_network_view.test"
	var v ipam.Networkview
	name := acctest.RandomNameWithPrefix("test-network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version0_14_0),
		},
		Steps: []resource.TestStep{
			//Provider Setup
			{
				Config: utils.ProviderSetup(),
			},
			// Create and Read
			{
				Config: testAccNetworkviewBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				Query:  true,
				Config: testAccNetworkviewListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_ipam_network_view.test", 1),
				},
			},
		},
	})
}

func TestAccNetworkviewList_Filters(t *testing.T) {
	var resourceName = "nios_ipam_network_view.test"
	var v ipam.Networkview
	name := acctest.RandomNameWithPrefix("test-network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version0_14_0),
		},
		Steps: []resource.TestStep{
			// Provider Setup
			{
				Config: utils.ProviderSetup(),
			},
			// Create and Read
			{
				Config: testAccNetworkviewBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Query the object
			{
				Query:  true,
				Config: testAccNetworkviewListConfigFilters(name),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_ipam_network_view.test", 1),
				},
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_ipam_network_view.test_extattrs"
	var v ipam.Networkview
	name := acctest.RandomNameWithPrefix("test-network-view")

	extAttrValue := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version0_14_0),
		},
		Steps: []resource.TestStep{
			// Provider Setup
			{
				Config: utils.ProviderSetup(),
			},
			// Create and Read
			{
				Config: testAccNetworkviewExtAttrs(name, map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				Query: true,

				Config: testAccNetworkviewListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_ipam_network_view.test", 1),
				},
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccNetworkviewListBasicConfig() string {
	return `
list "nios_ipam_network_view" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccNetworkviewListConfigFilters(name string) string {
	return fmt.Sprintf(`
list "nios_ipam_network_view" "test" {
	provider = nios
	config {
		filters = {
			name =  %q
		}
	}
}
`, name)
}

func testAccNetworkviewListConfigExtAttrFilters(name string) string {
	return fmt.Sprintf(`
list "nios_ipam_network_view" "test" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, name)
}
