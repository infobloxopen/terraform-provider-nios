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

func TestAccRecordDnameList_basic(t *testing.T) {
	var resourceName = "nios_dns_record_dname.test"
	var v dns.RecordDname
	target := acctest.RandomNameWithPrefix("test-dname") + ".com"
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccRecordDnameBasicConfig(target, zoneFqdn),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccRecordDnameListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_record_dname.test", 1),
				},
			},
		},
	})
}

func TestAccRecordDnameList_Filters(t *testing.T) {
	var resourceName = "nios_dns_record_dname.test"
	var v dns.RecordDname
	target := acctest.RandomNameWithPrefix("test-dname") + ".com"
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccRecordDnameBasicConfig(target, zoneFqdn),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", zoneFqdn),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccRecordDnameListConfigFilters(zoneFqdn),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_record_dname.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("record:dname/")),
						}),
						[]querycheck.KnownValueCheck{
							{
								Path:       tfjsonpath.New("name"),
								KnownValue: knownvalue.StringExact(zoneFqdn),
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

func TestAccRecordDnameList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dns_record_dname.test_extattrs"
	var v dns.RecordDname
	target := acctest.RandomNameWithPrefix("test-dname") + ".com"
	view := acctest.RandomNameWithPrefix("test-view")
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
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
				Config: testAccRecordDnameExtAttrs(target, view, zoneFqdn, map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccRecordDnameListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_record_dname.test", 1),
				},
			},
		},
	})
}

func testAccRecordDnameListBasicConfig() string {
	return `
list "nios_dns_record_dname" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccRecordDnameListConfigFilters(filterValue string) string {
	return fmt.Sprintf(`
list "nios_dns_record_dname" "test" {
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

func testAccRecordDnameListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_dns_record_dname" "test" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}
