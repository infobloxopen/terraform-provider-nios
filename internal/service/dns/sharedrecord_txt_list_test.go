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

func TestAccSharedrecordTxtList_basic(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_txt.test"
	var v dns.SharedrecordTxt
	name := acctest.RandomNameWithPrefix("sharedrecord-txt-")
	text := "This is a shared record TXT record"
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
				Config:                   testAccSharedrecordTxtBasicConfig(name, sharedRecordGroup, text),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordTxtExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccSharedrecordTxtListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_sharedrecord_txt.test", 1),
				},
			},
		},
	})
}

func TestAccSharedrecordTxtList_Filters(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_txt.test"
	var v dns.SharedrecordTxt
	name := acctest.RandomNameWithPrefix("sharedrecord-txt-")
	text := "This is a shared record TXT record"
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
				Config:                   testAccSharedrecordTxtBasicConfig(name, sharedRecordGroup, text),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccSharedrecordTxtListConfigFilters(name),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_sharedrecord_txt.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("sharedrecord:txt/")),
						}),
						[]querycheck.KnownValueCheck{
							{
								Path:       tfjsonpath.New("name"),
								KnownValue: knownvalue.StringExact(name),
							},
							{
								Path:       tfjsonpath.New("text"),
								KnownValue: knownvalue.StringExact(text),
							},
						},
					),
				},
			},
		},
	})
}

func TestAccSharedrecordTxtList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_txt.test_extattrs"
	var v dns.SharedrecordTxt

	extAttrValue := acctest.RandomName()
	name := acctest.RandomNameWithPrefix("sharedrecord-txt-")
	text := "This is a shared record TXT record"
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
				Config: testAccSharedrecordTxtExtAttrs(name, sharedRecordGroup, text, map[string]any{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccSharedrecordTxtListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_sharedrecord_txt.test_extattrs", 1),
				},
			},
		},
	})
}

func testAccSharedrecordTxtListBasicConfig() string {
	return `
list "nios_dns_sharedrecord_txt" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccSharedrecordTxtListConfigFilters(filterValue string) string {
	return fmt.Sprintf(`
list "nios_dns_sharedrecord_txt" "test" {
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

func testAccSharedrecordTxtListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_dns_sharedrecord_txt" "test_extattrs" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}
