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

func TestAccRecordNsList_basic(t *testing.T) {
	var resourceName = "nios_dns_record_ns.test"
	var v dns.RecordNs
	name := "example.com"
	nameserver := acctest.RandomNameWithPrefix("nameserver") + ".example.com"
	addresses := []map[string]any{
		{
			"address":         "20.0.0.0",
			"auto_create_ptr": false,
		},
	}
	addressesHCL := FormatZoneNameServersToHCL(addresses)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccRecordNsBasicConfig(name, nameserver, addressesHCL, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNsExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccRecordNsListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_record_ns.test", 1),
				},
			},
		},
	})
}

func TestAccRecordNsList_Filters(t *testing.T) {
	var resourceName = "nios_dns_record_ns.test"
	var v dns.RecordNs
	name := "example.com"
	nameserver := acctest.RandomNameWithPrefix("nameserver") + ".example.com"
	addresses := []map[string]any{
		{
			"address":         "20.0.0.0",
			"auto_create_ptr": false,
		},
	}
	addressesHCL := FormatZoneNameServersToHCL(addresses)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccRecordNsBasicConfig(name, nameserver, addressesHCL, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccRecordNsListConfigFilters(name, nameserver),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_record_ns.test", 1),
					querycheck.ExpectResourceKnownValues(
						resourceName,
						queryfilter.ByResourceIdentity(map[string]knownvalue.Check{
							"ref": knownvalue.StringRegexp(regexp.MustCompile("record:ns/")),
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

func testAccRecordNsListBasicConfig() string {
	return `
list "nios_dns_record_ns" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccRecordNsListConfigFilters(name, nameserver string) string {
	return fmt.Sprintf(`
list "nios_dns_record_ns" "test" {
	provider = nios
	include_resource = true
	config {
		filters = {
			name = %q
			nameserver = %q
		}
	}
}
`, name, nameserver)
}
