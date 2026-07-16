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

func TestAccSharedrecordCnameList_basic(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_cname.test"
	var v dns.SharedrecordCname
	name := acctest.RandomNameWithPrefix("sharedrecord-cname-")
	canonical := acctest.RandomNameWithPrefix("canonical-name") + ".com"
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
				Config:                   testAccSharedrecordCnameBasicConfig(name, canonical, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccSharedrecordCnameListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_sharedrecord_cname.test", 1),
				},
			},
		},
	})
}

func TestAccSharedrecordCnameList_Filters(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_cname.test"
	var v dns.SharedrecordCname
	name := acctest.RandomNameWithPrefix("sharedrecord-cname-")
	canonical := acctest.RandomNameWithPrefix("canonical-name") + ".com"
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
				Config:                   testAccSharedrecordCnameBasicConfig(name, canonical, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccSharedrecordCnameListConfigFilters(name),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_sharedrecord_cname.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("sharedrecord:cname/")),
						}),
						[]querycheck.KnownValueCheck{
							{
								Path:       tfjsonpath.New("name"),
								KnownValue: knownvalue.StringExact(name),
							},
							{
								Path:       tfjsonpath.New("canonical"),
								KnownValue: knownvalue.StringExact(canonical),
							},
						},
					),
				},
			},
		},
	})
}

func TestAccSharedrecordCnameList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_cname.test_extattrs"
	var v dns.SharedrecordCname

	extAttrValue := acctest.RandomName()
	name := acctest.RandomNameWithPrefix("sharedrecord-cname-")
	canonical := acctest.RandomNameWithPrefix("canonical-name") + ".com"
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
				Config: testAccSharedrecordCnameExtAttrs(name, canonical, sharedRecordGroup, map[string]any{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccSharedrecordCnameListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_sharedrecord_cname.test_extattrs", 1),
				},
			},
		},
	})
}

func testAccSharedrecordCnameListBasicConfig() string {
	return `
list "nios_dns_sharedrecord_cname" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccSharedrecordCnameListConfigFilters(filterValue string) string {
	return fmt.Sprintf(`
list "nios_dns_sharedrecord_cname" "test" {
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

func testAccSharedrecordCnameListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_dns_sharedrecord_cname" "test_extattrs" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}
