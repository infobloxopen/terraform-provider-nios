package security_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccNetworkuserDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_security_networkuser.test"
	resourceName := "nios_security_networkuser.test"
	var v security.Networkuser

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetworkuserDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkuserDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					}, testAccCheckNetworkuserResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckNetworkuserResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "result.0.address"),
		resource.TestCheckResourceAttrPair(resourceName, "address_object", dataSourceName, "result.0.address_object"),
		resource.TestCheckResourceAttrPair(resourceName, "data_source", dataSourceName, "result.0.data_source"),
		resource.TestCheckResourceAttrPair(resourceName, "data_source_ip", dataSourceName, "result.0.data_source_ip"),
		resource.TestCheckResourceAttrPair(resourceName, "domainname", dataSourceName, "result.0.domainname"),
		resource.TestCheckResourceAttrPair(resourceName, "first_seen_time", dataSourceName, "result.0.first_seen_time"),
		resource.TestCheckResourceAttrPair(resourceName, "guid", dataSourceName, "result.0.guid"),
		resource.TestCheckResourceAttrPair(resourceName, "last_seen_time", dataSourceName, "result.0.last_seen_time"),
		resource.TestCheckResourceAttrPair(resourceName, "last_updated_time", dataSourceName, "result.0.last_updated_time"),
		resource.TestCheckResourceAttrPair(resourceName, "logon_id", dataSourceName, "result.0.logon_id"),
		resource.TestCheckResourceAttrPair(resourceName, "logout_time", dataSourceName, "result.0.logout_time"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "network", dataSourceName, "result.0.network"),
		resource.TestCheckResourceAttrPair(resourceName, "network_view", dataSourceName, "result.0.network_view"),
		resource.TestCheckResourceAttrPair(resourceName, "user_status", dataSourceName, "result.0.user_status"),
	}
}

func testAccNetworkuserDataSourceConfigFilters() string {
	return `
resource "nios_security_networkuser" "test" {
}

data "nios_security_networkuser" "test" {
  filters = {
	 = nios_security_networkuser.test.
  }
}
`
}
