package ipam_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/querycheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccBulkhostnametemplateList_basic(t *testing.T) {
	var resourceName = "nios_ipam_bulk_hostname_template.test"
	var v ipam.Bulkhostnametemplate
	templateName := acctest.RandomNameWithPrefix("test-template")
	templateFormat := "host-$4"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config: testAccBulkhostnametemplateBasicConfig(templateName, templateFormat),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBulkhostnametemplateExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:  true,
				Config: testAccBulkhostnametemplateListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_ipam_bulk_hostname_template.test", 1),
				},
			},
		},
	})
}

func TestAccBulkhostnametemplateList_Filters(t *testing.T) {
	var resourceName = "nios_ipam_bulk_hostname_template.test"
	var v ipam.Bulkhostnametemplate
	templateName := acctest.RandomNameWithPrefix("test-template")
	templateFormat := "host-$4"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config: testAccBulkhostnametemplateBasicConfig(templateName, templateFormat),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBulkhostnametemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_name", templateName),
				),
			},
			// Query the object
			{

				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:  true,
				Config: testAccBulkhostnametemplateListConfigFilters(templateName),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_ipam_bulk_hostname_template.test", 1),
				},
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccBulkhostnametemplateListBasicConfig() string {
	return `
list "nios_ipam_bulk_hostname_template" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccBulkhostnametemplateListConfigFilters(name string) string {
	return fmt.Sprintf(`
list "nios_ipam_bulk_hostname_template" "test" {
	provider = nios
	config {
		filters = {
			template_name =  %q
		}
	}
}
`, name)
}
