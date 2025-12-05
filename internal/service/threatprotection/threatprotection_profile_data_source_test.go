
package threatprotection_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/threatprotection"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccThreatprotectionProfileDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_threatprotection_profile.test"
	resourceName := "nios_threatprotection_profile.test"
	var v threatprotection.ThreatprotectionProfile

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckThreatprotectionProfileDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccThreatprotectionProfileDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
						}, testAccCheckThreatprotectionProfileResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccThreatprotectionProfileDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_threatprotection_profile.test"
	resourceName := "nios_threatprotection_profile.test"
	var v threatprotection.ThreatprotectionProfile
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckThreatprotectionProfileDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccThreatprotectionProfileDataSourceConfigExtAttrFilters( acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
						}, testAccCheckThreatprotectionProfileResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckThreatprotectionProfileResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
    return []resource.TestCheckFunc{
        resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
        resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
        resource.TestCheckResourceAttrPair(resourceName, "current_ruleset", dataSourceName, "result.0.current_ruleset"),
        resource.TestCheckResourceAttrPair(resourceName, "disable_multiple_dns_tcp_request", dataSourceName, "result.0.disable_multiple_dns_tcp_request"),
        resource.TestCheckResourceAttrPair(resourceName, "events_per_second_per_rule", dataSourceName, "result.0.events_per_second_per_rule"),
        resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
        resource.TestCheckResourceAttrPair(resourceName, "members", dataSourceName, "result.0.members"),
        resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
        resource.TestCheckResourceAttrPair(resourceName, "source_member", dataSourceName, "result.0.source_member"),
        resource.TestCheckResourceAttrPair(resourceName, "source_profile", dataSourceName, "result.0.source_profile"),
        resource.TestCheckResourceAttrPair(resourceName, "use_current_ruleset", dataSourceName, "result.0.use_current_ruleset"),
        resource.TestCheckResourceAttrPair(resourceName, "use_disable_multiple_dns_tcp_request", dataSourceName, "result.0.use_disable_multiple_dns_tcp_request"),
        resource.TestCheckResourceAttrPair(resourceName, "use_events_per_second_per_rule", dataSourceName, "result.0.use_events_per_second_per_rule"),
    }
}

func testAccThreatprotectionProfileDataSourceConfigFilters() string {
	return `
resource "nios_threatprotection_profile" "test" {
}

data "nios_threatprotection_profile" "test" {
  filters = {
	 = nios_threatprotection_profile.test.
  }
}
`
}

func testAccThreatprotectionProfileDataSourceConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_profile" "test" {
  extattrs = {
    Site = %q
  } 
}

data "nios_threatprotection_profile" "test" {
  extattrfilters = {
	Site = nios_threatprotection_profile.test.extattrs.Site
  }
}
`,extAttrsValue)
}

