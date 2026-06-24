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

func TestAccRecordTxtList_basic(t *testing.T) {
	var resourceName = "nios_dns_record_txt.test"
	var v dns.RecordTxt
	name := acctest.RandomNameWithPrefix("record-txt") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccRecordTxtBasicConfig(name, "Record Text", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccRecordTxtListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_record_txt.test", 1),
				},
			},
		},
	})
}

func TestAccRecordTxtList_Filters(t *testing.T) {
	var resourceName = "nios_dns_record_txt.test"
	var v dns.RecordTxt
	name := acctest.RandomNameWithPrefix("record-txt") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccRecordTxtBasicConfig(name, "Record Text", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccRecordTxtListConfigFilters(name),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_record_txt.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("record:txt/")),
						}),
						[]querycheck.KnownValueCheck{
							{
								Path:       tfjsonpath.New("name"),
								KnownValue: knownvalue.StringExact(name),
							},
							{
								Path:       tfjsonpath.New("view"),
								KnownValue: knownvalue.StringExact("default"),
							},
						},
					),
				},
			},
		},
	})
}

func TestAccRecordTxtList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dns_record_txt.test_extattrs"
	var v dns.RecordTxt
	name := acctest.RandomNameWithPrefix("record-txt") + ".example.com"
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
				Config: testAccRecordTxtExtAttrs(name, "Record Text", map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccRecordTxtListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_record_txt.test_extattrs", 1),
				},
			},
		},
	})
}

func testAccRecordTxtListBasicConfig() string {
	return `
list "nios_dns_record_txt" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccRecordTxtListConfigFilters(filterValue string) string {
	return fmt.Sprintf(`
list "nios_dns_record_txt" "test" {
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

func testAccRecordTxtListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_dns_record_txt" "test_extattrs" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}
