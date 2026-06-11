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

func TestAccRecordSrvList_basic(t *testing.T) {
	var resourceName = "nios_dns_record_srv.test"
	var v dns.RecordSrv
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccRecordSrvBasicConfig(name, target, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccRecordSrvListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_record_srv.test", 1),
				},
			},
		},
	})
}

func TestAccRecordSrvList_Filters(t *testing.T) {
	var resourceName = "nios_dns_record_srv.test"
	var v dns.RecordSrv
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccRecordSrvBasicConfig(name, target, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccRecordSrvListConfigFilters(name),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_record_srv.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("record:srv/")),
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

func TestAccRecordSrvList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dns_record_srv.test_extattrs"
	var v dns.RecordSrv
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"

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
				Config: testAccRecordSrvExtAttrs(name, target, 80, 10, 360, map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccRecordSrvListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_record_srv.test_extattrs", 1),
				},
			},
		},
	})
}

func testAccRecordSrvListBasicConfig() string {
	return `
list "nios_dns_record_srv" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccRecordSrvListConfigFilters(filterValue string) string {
	return fmt.Sprintf(`
list "nios_dns_record_srv" "test" {
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

func testAccRecordSrvListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_dns_record_srv" "test" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}
