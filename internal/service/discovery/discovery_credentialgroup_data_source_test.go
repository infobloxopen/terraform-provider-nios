package discovery_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/discovery"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccDiscoveryCredentialgroupDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_discovery_credentialgroup.test"
	resourceName := "nios_discovery_credentialgroup.test"
	var v discovery.DiscoveryCredentialgroup
	name := acctest.RandomNameWithPrefix("example-discovery-credentialgroup")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDiscoveryCredentialgroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDiscoveryCredentialgroupDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDiscoveryCredentialgroupExists(context.Background(), resourceName, &v),
					}, testAccCheckDiscoveryCredentialgroupResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDiscoveryCredentialgroupResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.*.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.*.name"),
	}
}

func testAccDiscoveryCredentialgroupDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_discovery_credentialgroup" "test" {
  name = %q
}

data "nios_discovery_credentialgroup" "test" {
  depends_on = [nios_discovery_credentialgroup.test]
}
`, name)
}
