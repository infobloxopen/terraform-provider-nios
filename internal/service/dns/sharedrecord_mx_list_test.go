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

func TestAccSharedrecordMxList_basic(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_mx.test"
	var v dns.SharedrecordMx
	name := acctest.RandomNameWithPrefix("sharedrecord-mx") + ".example.com"
	mailExchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
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
				Config:                   testAccSharedrecordMxBasicConfig(mailExchanger, name, 10, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccSharedrecordMxListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_sharedrecord_mx.test", 1),
				},
			},
		},
	})
}

func TestAccSharedrecordMxList_Filters(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_mx.test"
	var v dns.SharedrecordMx
	name := acctest.RandomNameWithPrefix("sharedrecord-mx") + ".example.com"
	mailExchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
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
				Config:                   testAccSharedrecordMxBasicConfig(mailExchanger, name, 10, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccSharedrecordMxListConfigFilters(name),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_sharedrecord_mx.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("sharedrecord:mx/")),
						}),
						[]querycheck.KnownValueCheck{
							{
								Path:       tfjsonpath.New("name"),
								KnownValue: knownvalue.StringExact(name),
							},
							{
								Path:       tfjsonpath.New("mail_exchanger"),
								KnownValue: knownvalue.StringExact(mailExchanger),
							},
						},
					),
				},
			},
		},
	})
}

func TestAccSharedrecordMxList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_mx.test_extattrs"
	var v dns.SharedrecordMx

	extAttrValue := acctest.RandomName()
	name := acctest.RandomNameWithPrefix("sharedrecord-mx") + ".example.com"
	mailExchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
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
				Config: testAccSharedrecordMxExtAttrs(mailExchanger, name, 10, sharedRecordGroup, map[string]any{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccSharedrecordMxListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_sharedrecord_mx.test_extattrs", 1),
				},
			},
		},
	})
}

func testAccSharedrecordMxListBasicConfig() string {
	return `
list "nios_dns_sharedrecord_mx" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccSharedrecordMxListConfigFilters(filterValue string) string {
	return fmt.Sprintf(`
list "nios_dns_sharedrecord_mx" "test" {
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

func testAccSharedrecordMxListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_dns_sharedrecord_mx" "test_extattrs" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}
