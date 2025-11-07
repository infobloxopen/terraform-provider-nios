package microsoftserver_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/microsoftserver"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccMsserverDhcpDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_microsoftserver_msserver_dhcp.test"
	resourceName := "nios_microsoftserver_msserver_dhcp.test"
	var v microsoftserver.MsserverDhcp

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMsserverDhcpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMsserverDhcpDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckMsserverDhcpExists(context.Background(), resourceName, &v),
					}, testAccCheckMsserverDhcpResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckMsserverDhcpResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "result.0.address"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_utilization", dataSourceName, "result.0.dhcp_utilization"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_utilization_status", dataSourceName, "result.0.dhcp_utilization_status"),
		resource.TestCheckResourceAttrPair(resourceName, "dynamic_hosts", dataSourceName, "result.0.dynamic_hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "last_sync_ts", dataSourceName, "result.0.last_sync_ts"),
		resource.TestCheckResourceAttrPair(resourceName, "login_name", dataSourceName, "result.0.login_name"),
		resource.TestCheckResourceAttrPair(resourceName, "login_password", dataSourceName, "result.0.login_password"),
		resource.TestCheckResourceAttrPair(resourceName, "network_view", dataSourceName, "result.0.network_view"),
		resource.TestCheckResourceAttrPair(resourceName, "next_sync_control", dataSourceName, "result.0.next_sync_control"),
		resource.TestCheckResourceAttrPair(resourceName, "read_only", dataSourceName, "result.0.read_only"),
		resource.TestCheckResourceAttrPair(resourceName, "server_name", dataSourceName, "result.0.server_name"),
		resource.TestCheckResourceAttrPair(resourceName, "static_hosts", dataSourceName, "result.0.static_hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "status", dataSourceName, "result.0.status"),
		resource.TestCheckResourceAttrPair(resourceName, "status_detail", dataSourceName, "result.0.status_detail"),
		resource.TestCheckResourceAttrPair(resourceName, "status_last_updated", dataSourceName, "result.0.status_last_updated"),
		resource.TestCheckResourceAttrPair(resourceName, "supports_failover", dataSourceName, "result.0.supports_failover"),
		resource.TestCheckResourceAttrPair(resourceName, "synchronization_interval", dataSourceName, "result.0.synchronization_interval"),
		resource.TestCheckResourceAttrPair(resourceName, "total_hosts", dataSourceName, "result.0.total_hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "use_login", dataSourceName, "result.0.use_login"),
		resource.TestCheckResourceAttrPair(resourceName, "use_synchronization_interval", dataSourceName, "result.0.use_synchronization_interval"),
	}
}

func testAccMsserverDhcpDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_dhcp" "test" {
}

data "nios_microsoftserver_msserver_dhcp" "test" {
  filters = {
	 = nios_microsoftserver_msserver_dhcp.test.
  }
}
`)
}
