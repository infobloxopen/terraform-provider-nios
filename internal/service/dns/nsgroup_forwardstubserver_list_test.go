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

func TestAccNsgroupForwardstubserverList_basic(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardstubserver.test"
	var v dns.NsgroupForwardstubserver
	name := acctest.RandomNameWithPrefix("ns-group-forwardstubserver")
	externalServers := []map[string]any{
		{
			"name":    "example.com",
			"address": "2.3.3.4",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccNsgroupForwardstubserverBasicConfig(name, externalServers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccNsgroupForwardstubserverListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_nsgroup_forwardstubserver.test", 1),
				},
			},
		},
	})
}

func TestAccNsgroupForwardstubserverList_Filters(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardstubserver.test"
	var v dns.NsgroupForwardstubserver
	name := acctest.RandomNameWithPrefix("ns-group-forwardstubserver")
	externalServers := []map[string]any{
		{
			"name":    "example.com",
			"address": "2.3.3.4",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccNsgroupForwardstubserverBasicConfig(name, externalServers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccNsgroupForwardstubserverListConfigFilters(name),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_nsgroup_forwardstubserver.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("nsgroup:forwardstubserver/")),
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

func TestAccNsgroupForwardstubserverList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardstubserver.test_extattrs"
	var v dns.NsgroupForwardstubserver
	name := acctest.RandomNameWithPrefix("ns-group-forwardstubserver")
	externalServers := []map[string]any{
		{
			"name":    "example.com",
			"address": "2.3.3.4",
		},
	}

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
				Config: testAccNsgroupForwardstubserverExtAttrs(name, externalServers, map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccNsgroupForwardstubserverListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_nsgroup_forwardstubserver.test_extattrs", 1),
				},
			},
		},
	})
}

func testAccNsgroupForwardstubserverListBasicConfig() string {
	return `
list "nios_dns_nsgroup_forwardstubserver" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccNsgroupForwardstubserverListConfigFilters(filterValue string) string {
	return fmt.Sprintf(`
list "nios_dns_nsgroup_forwardstubserver" "test" {
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

func testAccNsgroupForwardstubserverListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_dns_nsgroup_forwardstubserver" "test_extattrs" {
	provider = nios
	config {
		extattrfilters = {
			Site = %q
		}
	}
}
`, extAttrsValue)
}
