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

func TestAccVlanrangeList_basic(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test"
	var v ipam.Vlanrange
	vlanRange := acctest.RandomNameWithPrefix("vlan-range")
	vlanView := acctest.RandomNameWithPrefix("vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccVlanrangeBasicConfig(71, vlanRange, 61, vlanView),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccVlanrangeListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_ipam_vlanrange.test", 1),
				},
			},
		},
	})
}

func TestAccVlanrangeList_Filters(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test"
	var v ipam.Vlanrange
	vlanRange := acctest.RandomNameWithPrefix("vlan-range")
	vlanView := acctest.RandomNameWithPrefix("vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccVlanrangeBasicConfig(71, vlanRange, 61, vlanView),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", vlanRange),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccVlanrangeListConfigFilters(vlanRange),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_ipam_vlanrange.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("vlanrange/")),
						}),
						[]querycheck.KnownValueCheck{
							{
								Path:       tfjsonpath.New("name"),
								KnownValue: knownvalue.StringExact(vlanRange),
							},
							{
								Path:       tfjsonpath.New("end_vlan_id"),
								KnownValue: knownvalue.Int64Exact(71),
							},
							{
								Path:       tfjsonpath.New("start_vlan_id"),
								KnownValue: knownvalue.Int64Exact(61),
							},
							{
								Path:       tfjsonpath.New("vlan_view"),
								KnownValue: knownvalue.StringRegexp(regexp.MustCompile("vlanview/")),
							},
						},
					),
				},
			},
		},
	})
}

func TestAccVlanrangeList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test_extattrs"
	var v ipam.Vlanrange
	vlanRange := acctest.RandomNameWithPrefix("vlan-range")
	vlanView := acctest.RandomNameWithPrefix("vlan-view")

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
				Config: testAccVlanrangeExtAttrs(71, vlanRange, 61, vlanView, map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccVlanrangeListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_ipam_vlanrange.test", 1),
				},
			},
		},
	})
}

func testAccVlanrangeListBasicConfig() string {
	return `
list "nios_ipam_vlanrange" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccVlanrangeListConfigFilters(name string) string {
	return fmt.Sprintf(`
list "nios_ipam_vlanrange" "test" {
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

func testAccVlanrangeListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_ipam_vlanrange" "test" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}
