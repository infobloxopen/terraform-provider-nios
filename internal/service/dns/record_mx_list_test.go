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

func TestAccRecordMxList_basic(t *testing.T) {
	var resourceName = "nios_dns_record_mx.test"
	var v dns.RecordMx
	name := acctest.RandomNameWithPrefix("record-mx") + ".example.com"
	mailExchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccRecordMxBasicConfig(name, mailExchanger, 10, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccRecordMxListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_record_mx.test", 1),
				},
			},
		},
	})
}

func TestAccRecordMxList_Filters(t *testing.T) {
	var resourceName = "nios_dns_record_mx.test"
	var v dns.RecordMx
	name := acctest.RandomNameWithPrefix("record-mx") + ".example.com"
	mailExchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccRecordMxBasicConfig(name, mailExchanger, 10, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccRecordMxListConfigFilters(name),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_record_mx.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("record:mx/")),
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

func TestAccRecordMxList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dns_record_mx.test_extattrs"
	var v dns.RecordMx
	name := acctest.RandomNameWithPrefix("record-mx") + ".example.com"
	mailExchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
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
				Config: testAccRecordMxExtAttrs(name, mailExchanger, 10, "default", map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccRecordMxListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_record_mx.test", 1),
				},
			},
		},
	})
}

func testAccRecordMxListBasicConfig() string {
	return `
list "nios_dns_record_mx" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccRecordMxListConfigFilters(filterValue string) string {
	return fmt.Sprintf(`
list "nios_dns_record_mx" "test" {
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

func testAccRecordMxListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_dns_record_mx" "test" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}
