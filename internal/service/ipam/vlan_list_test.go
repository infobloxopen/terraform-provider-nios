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

func TestAccVlanList_basic(t *testing.T) {
	var resourceName = "nios_ipam_vlan.test"
	var v ipam.Vlan
	name := acctest.RandomNameWithPrefix("vlan")
	view := acctest.RandomNameWithPrefix("example-vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccVlanBasicConfig(51, name, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccVlanListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_ipam_vlan.test", 1),
				},
			},
		},
	})
}

func TestAccVlanList_Filters(t *testing.T) {
	var resourceName = "nios_ipam_vlan.test"
	var v ipam.Vlan
	name := acctest.RandomNameWithPrefix("vlan")
	view := acctest.RandomNameWithPrefix("example-vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccVlanBasicConfig(51, name, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccVlanListConfigFilters(name),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_ipam_vlan.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("vlan/")),
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

func TestAccVlanList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_ipam_vlan.test_extattrs"
	var v ipam.Vlan
	name := acctest.RandomNameWithPrefix("vlan")
	view := acctest.RandomNameWithPrefix("example-vlan-view")

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
				Config: testAccVlanExtAttrs(51, name, view, map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccVlanListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_ipam_vlan.test", 1),
				},
			},
		},
	})
}

func testAccVlanListBasicConfig() string {
	return `
list "nios_ipam_vlan" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccVlanListConfigFilters(name string) string {
	return fmt.Sprintf(`
list "nios_ipam_vlan" "test" {
	provider = nios
	include_resource = true
	config {
		filters = {
			name =  %q
		}
	}
}
`, name)
}

func testAccVlanListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_ipam_vlan" "test" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}
