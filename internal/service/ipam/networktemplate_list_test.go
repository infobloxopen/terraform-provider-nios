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

func TestAccNetworktemplateList_basic(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccNetworktemplateBasicConfig(name, 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccNetworktemplateListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_ipam_networktemplate.test", 1),
				},
			},
		},
	})
}

func TestAccNetworktemplateList_Filters(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccNetworktemplateBasicConfig(name, 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccNetworktemplateListConfigFilters(name),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_ipam_networktemplate.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							// TODO : Update the ref prefix with the correct identifying object for the resource
							"ref": knownvalue.StringRegexp(regexp.MustCompile("networktemplate/")),
						}),
						[]querycheck.KnownValueCheck{
							// TODO : Add checks for required fields
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

func TestAccNetworktemplateList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_extattrs"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

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
				Config: testAccNetworktemplateExtAttrs(name, 24, map[string]string{
					"Tenant ID": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Tenant ID", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccNetworktemplateListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_ipam_networktemplate.test", 1),
				},
			},
		},
	})
}

func testAccNetworktemplateListBasicConfig() string {
	return `
list "nios_ipam_networktemplate" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccNetworktemplateListConfigFilters(name string) string {
	return fmt.Sprintf(`
list "nios_ipam_networktemplate" "test" {
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

func testAccNetworktemplateListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_ipam_networktemplate" "test" {
	provider = nios
	config {
		extattrfilters = {
			"Tenant ID" =  %q
		}
	}
}
`, extAttrsValue)
}
