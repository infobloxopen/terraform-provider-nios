package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/querycheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

func TestAccRecordAList_basic(t *testing.T) {
	var resourceName = "nios_dns_record_a.test"
	var v dns.RecordA
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version0_14_0),
		},
		Steps: []resource.TestStep{
			//Provider Setup
			{
				Config: utils.ProviderSetup(),
			},
			// Create and Read
			{
				Config: testAccRecordABasicConfig(name, "10.0.0.20", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
				),
			},
			{
				Query:  true,
				Config: testAccRecordAListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dns_record_a.test", 1),
				},
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAList_Filters(t *testing.T) {
	var resourceName = "nios_dns_record_a.test"
	var v dns.RecordA
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version0_14_0),
		},
		Steps: []resource.TestStep{
			// Provider Setup
			{
				Config: utils.ProviderSetup(),
			},
			// Create and Read
			{
				Config: testAccRecordABasicConfig(name, "10.0.0.21", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Query:  true,
				Config: testAccRecordAListConfigFilters(name),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_record_a.test", 1),
				},
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dns_record_a.test_extattrs"
	var v dns.RecordA
	name := acctest.RandomName() + ".example.com"
	extAttrValue := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version0_14_0),
		},
		Steps: []resource.TestStep{
			// Provider Setup
			{
				Config: utils.ProviderSetup(),
			},
			// Create and Read
			{
				Config: testAccRecordAExtattrs(name, "10.0.0.22", "default", map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			{
				Query: true,

				Config: testAccRecordAListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dns_record_a.test", 1),
				},
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccRecordAListBasicConfig() string {
	return `
list "nios_dns_record_a" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccRecordAListConfigFilters(name string) string {
	return fmt.Sprintf(`
list "nios_dns_record_a" "test" {
	provider = nios
	config {
		filters = {
			name =  %q
		}
	}
}
`, name)
}

func testAccRecordAListConfigExtAttrFilters(name string) string {
	return fmt.Sprintf(`
list "nios_dns_record_a" "test" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, name)
}
