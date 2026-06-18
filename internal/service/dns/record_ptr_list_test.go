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

func TestAccRecordPtrList_basic(t *testing.T) {
	var resourceName = "nios_dns_record_ptr.test"
	var v dns.RecordPtr
	ptrDName := acctest.RandomNameWithPrefix("ptr") + ".example.com"
	ipv4addr := "192.168.10.22"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccRecordPtrBasicConfig(ipv4addr, ptrDName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccRecordPtrListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_record_ptr.test", 1),
				},
			},
		},
	})
}

func TestAccRecordPtrList_Filters(t *testing.T) {
	var resourceName = "nios_dns_record_ptr.test"
	var v dns.RecordPtr
	ptrDName := acctest.RandomNameWithPrefix("ptr") + ".example.com"
	ipv4addr := "192.168.10.22"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccRecordPtrBasicConfig(ipv4addr, ptrDName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", ipv4addr),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccRecordPtrListConfigFilters(ptrDName),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_record_ptr.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("record:ptr/")),
						}),
						[]querycheck.KnownValueCheck{
							{
								Path:       tfjsonpath.New("name"),
								KnownValue: knownvalue.StringExact("22.10.168.192.in-addr.arpa"),
							},
							{
								Path:       tfjsonpath.New("ptrdname"),
								KnownValue: knownvalue.StringExact(ptrDName),
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

func TestAccRecordPtrList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dns_record_ptr.test_extattrs"
	var v dns.RecordPtr
	ptrDName := acctest.RandomNameWithPrefix("ptr") + ".example.com"
	name := "22.10.168.192.in-addr.arpa"

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
				Config: testAccRecordPtrExtAttrs(name, ptrDName, "default", map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccRecordPtrListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_record_ptr.test", 1),
				},
			},
		},
	})
}

func testAccRecordPtrListBasicConfig() string {
	return `
list "nios_dns_record_ptr" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccRecordPtrListConfigFilters(filterValue string) string {
	return fmt.Sprintf(`
list "nios_dns_record_ptr" "test" {
	provider = nios
	include_resource = true
	config {
		filters = {
			ptrdname =  %q
		}
	}
}
`, filterValue)
}

func testAccRecordPtrListConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
list "nios_dns_record_ptr" "test" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrsValue)
}
