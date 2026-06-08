
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

func TestAccRecordCaaList_basic(t *testing.T) {
	var resourceName = "nios_dns_record_caa.test"
	var v dns.RecordCaa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-caa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccRecordCaaBasicConfig(zoneFqdn, name, 0, "issue", "digicert.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:  true,
				Config: testAccRecordCaaListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_record_caa.test", 1),
				},
			},
		},
	})
}

func TestAccRecordCaaList_Filters(t *testing.T) {
	var resourceName = "nios_dns_record_caa.test"
	var v dns.RecordCaa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-caa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccRecordCaaBasicConfig(zoneFqdn, name, 0, "issue", "digicert.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_tag", "issue"),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:  true,
				Config: testAccRecordCaaListConfigFilters(fmt.Sprintf("%s.%s", name, zoneFqdn)),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_record_caa.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("record:caa/")),
						}),
						[]querycheck.KnownValueCheck{
							{
								Path:       tfjsonpath.New("ca_tag"),
								KnownValue: knownvalue.StringExact("issue"),
							},
							{
								Path:       tfjsonpath.New("ca_value"),
								KnownValue: knownvalue.StringExact("digicert.com"),
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

func TestAccRecordCaaList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dns_record_caa.test_extattrs"
	var v dns.RecordCaa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-caa")
	extAttrValue := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config: testAccRecordCaaExtAttrs(zoneFqdn, name, 0, "issue", "digicert.com", map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query: true,
				Config: testAccRecordCaaListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_record_caa.test", 1),
				},
			},
		},
	})
}

func testAccRecordCaaListBasicConfig() string {
	return `
list "nios_dns_record_caa" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccRecordCaaListConfigFilters(filterValue string) string {
	return fmt.Sprintf(`
list "nios_dns_record_caa" "test" {
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

func testAccRecordCaaListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_dns_record_caa" "test" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}

