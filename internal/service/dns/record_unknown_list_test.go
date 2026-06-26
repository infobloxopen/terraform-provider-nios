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

func TestAccRecordUnknownList_basic(t *testing.T) {
	var resourceName = "nios_dns_record_unknown.test"
	var v dns.RecordUnknown
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-unknown")
	subfieldValues := []map[string]any{
		{
			"field_type":     "T",
			"field_value":    "example-text",
			"include_length": "8_BIT",
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
				Config:                   testAccRecordUnknownBasicConfig(zoneFqdn, name, "SPF", subfieldValues),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccRecordUnknownListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_record_unknown.test", 1),
				},
			},
		},
	})
}

func TestAccRecordUnknownList_Filters(t *testing.T) {
	var resourceName = "nios_dns_record_unknown.test"
	var v dns.RecordUnknown
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-unknown")
	subfieldValues := []map[string]any{
		{
			"field_type":     "T",
			"field_value":    "example-text",
			"include_length": "8_BIT",
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
				Config:                   testAccRecordUnknownBasicConfig(zoneFqdn, name, "SPF", subfieldValues),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "record_type", "SPF"),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccRecordUnknownListConfigFilters(name + "." + zoneFqdn),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_record_unknown.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("record:unknown/")),
						}),
						[]querycheck.KnownValueCheck{
							{
								Path:       tfjsonpath.New("name"),
								KnownValue: knownvalue.StringExact(name + "." + zoneFqdn),
							},
							{
								Path:       tfjsonpath.New("record_type"),
								KnownValue: knownvalue.StringExact("SPF"),
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

func TestAccRecordUnknownList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dns_record_unknown.test_extattrs"
	var v dns.RecordUnknown
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-unknown")
	subfieldValues := []map[string]any{
		{
			"field_type":     "T",
			"field_value":    "example-text",
			"include_length": "8_BIT",
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
				Config: testAccRecordUnknownExtAttrs(zoneFqdn, name, "SPF", subfieldValues, map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccRecordUnknownListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_record_unknown.test", 1),
				},
			},
		},
	})
}

func testAccRecordUnknownListBasicConfig() string {
	return `
list "nios_dns_record_unknown" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccRecordUnknownListConfigFilters(filterValue string) string {
	return fmt.Sprintf(`
list "nios_dns_record_unknown" "test" {
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

func testAccRecordUnknownListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_dns_record_unknown" "test" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}
