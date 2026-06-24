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

func TestAccIpv6networktemplateList_basic(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test"
	var v ipam.Ipv6networktemplate
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
				Config:                   testAccIpv6networktemplateBasicConfig(name, 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccIpv6networktemplateListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_ipam_ipv6networktemplate.test", 1),
				},
			},
		},
	})
}

func TestAccIpv6networktemplateList_Filters(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test"
	var v ipam.Ipv6networktemplate
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
				Config:                   testAccIpv6networktemplateBasicConfig(name, 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccIpv6networktemplateListConfigFilters(name),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_ipam_ipv6networktemplate.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("ipv6networktemplate/")),
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

func TestAccIpv6networktemplateList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_extattrs"
	var v ipam.Ipv6networktemplate
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
				Config: testAccIpv6networktemplateExtAttrs(name, 24, map[string]string{
					"Tenant ID": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Tenant ID", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccIpv6networktemplateListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_ipam_ipv6networktemplate.test", 1),
				},
			},
		},
	})
}

func testAccIpv6networktemplateListBasicConfig() string {
	return `
list "nios_ipam_ipv6networktemplate" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccIpv6networktemplateListConfigFilters(filterValue string) string {
	return fmt.Sprintf(`
list "nios_ipam_ipv6networktemplate" "test" {
	provider = nios
	include_resource = true
	config {
		filters = {
			name =  %q
		}
	}
}
`, filterValue)
}

func testAccIpv6networktemplateListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_ipam_ipv6networktemplate" "test" {
	provider = nios
	config {
		extattrfilters = {
			"Tenant ID" =  %q
		}
	}
}
`, extAttrsValue)
}
