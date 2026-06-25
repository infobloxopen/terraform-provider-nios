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

func TestAccSharedrecordSrvList_basic(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_srv.test"
	var v dns.SharedrecordSrv
	name := acctest.RandomNameWithPrefix("sharedrecord-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccSharedrecordSrvBasicConfig(name, 10, 80, sharedRecordGroup, target, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccSharedrecordSrvListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_sharedrecord_srv.test", 1),
				},
			},
		},
	})
}

func TestAccSharedrecordSrvList_Filters(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_srv.test"
	var v dns.SharedrecordSrv
	name := acctest.RandomNameWithPrefix("sharedrecord-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccSharedrecordSrvBasicConfig(name, 10, 80, sharedRecordGroup, target, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccSharedrecordSrvListConfigFilters(name),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_sharedrecord_srv.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("sharedrecord:srv/")),
						}),
						[]querycheck.KnownValueCheck{
							{
								Path:       tfjsonpath.New("name"),
								KnownValue: knownvalue.StringExact(name),
							},
							{
								Path:       tfjsonpath.New("target"),
								KnownValue: knownvalue.StringExact(target),
							},
						},
					),
				},
			},
		},
	})
}

func TestAccSharedrecordSrvList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_srv.test_extattrs"
	var v dns.SharedrecordSrv

	extAttrValue := acctest.RandomName()
	name := acctest.RandomNameWithPrefix("sharedrecord-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config: testAccSharedrecordSrvExtAttrs(name, 10, 80, sharedRecordGroup, target, 10, map[string]any{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccSharedrecordSrvListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_sharedrecord_srv.test_extattrs", 1),
				},
			},
		},
	})
}

func testAccSharedrecordSrvListBasicConfig() string {
	return `
list "nios_dns_sharedrecord_srv" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccSharedrecordSrvListConfigFilters(filterValue string) string {
	return fmt.Sprintf(`
list "nios_dns_sharedrecord_srv" "test" {
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

func testAccSharedrecordSrvListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_dns_sharedrecord_srv" "test_extattrs" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}
