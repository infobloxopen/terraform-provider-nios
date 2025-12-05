package dhcp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccDhcpfailoverDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_dhcpfailover.test"
	resourceName := "nios_dhcp_dhcpfailover.test"
	var v dhcp.Dhcpfailover
	failoverName := acctest.RandomNameWithPrefix("failover")
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDhcpfailoverDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDhcpfailoverDataSourceConfigFilters(failoverName, primary, secondary, primaryServerType, secondaryServerType),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					}, testAccCheckDhcpfailoverResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccDhcpfailoverDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_dhcpfailover.test"
	resourceName := "nios_dhcp_dhcpfailover.test"
	var v dhcp.Dhcpfailover
	extAttrValue := acctest.RandomNameWithPrefix("Site")
	primary := "infoblox.localdomain"
	secondary := "infoblox.member1"
	primaryServerType := "GRID"
	secondaryServerType := "GRID"
	dhcpfailoverName := acctest.RandomNameWithPrefix("failover")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDhcpfailoverDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDhcpfailoverDataSourceConfigExtAttrFilters(dhcpfailoverName, primary, secondary, primaryServerType, secondaryServerType, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDhcpfailoverExists(context.Background(), resourceName, &v),
					}, testAccCheckDhcpfailoverResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDhcpfailoverResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "association_type", dataSourceName, "result.0.association_type"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "failover_port", dataSourceName, "result.0.failover_port"),
		resource.TestCheckResourceAttrPair(resourceName, "load_balance_split", dataSourceName, "result.0.load_balance_split"),
		resource.TestCheckResourceAttrPair(resourceName, "max_client_lead_time", dataSourceName, "result.0.max_client_lead_time"),
		resource.TestCheckResourceAttrPair(resourceName, "max_load_balance_delay", dataSourceName, "result.0.max_load_balance_delay"),
		resource.TestCheckResourceAttrPair(resourceName, "max_response_delay", dataSourceName, "result.0.max_response_delay"),
		resource.TestCheckResourceAttrPair(resourceName, "max_unacked_updates", dataSourceName, "result.0.max_unacked_updates"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_association_mode", dataSourceName, "result.0.ms_association_mode"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_enable_authentication", dataSourceName, "result.0.ms_enable_authentication"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_enable_switchover_interval", dataSourceName, "result.0.ms_enable_switchover_interval"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_failover_mode", dataSourceName, "result.0.ms_failover_mode"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_failover_partner", dataSourceName, "result.0.ms_failover_partner"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_hotstandby_partner_role", dataSourceName, "result.0.ms_hotstandby_partner_role"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_is_conflict", dataSourceName, "result.0.ms_is_conflict"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_previous_state", dataSourceName, "result.0.ms_previous_state"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_server", dataSourceName, "result.0.ms_server"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_shared_secret", dataSourceName, "result.0.ms_shared_secret"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_state", dataSourceName, "result.0.ms_state"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_switchover_interval", dataSourceName, "result.0.ms_switchover_interval"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "primary", dataSourceName, "result.0.primary"),
		resource.TestCheckResourceAttrPair(resourceName, "primary_server_type", dataSourceName, "result.0.primary_server_type"),
		resource.TestCheckResourceAttrPair(resourceName, "primary_state", dataSourceName, "result.0.primary_state"),
		resource.TestCheckResourceAttrPair(resourceName, "recycle_leases", dataSourceName, "result.0.recycle_leases"),
		resource.TestCheckResourceAttrPair(resourceName, "secondary", dataSourceName, "result.0.secondary"),
		resource.TestCheckResourceAttrPair(resourceName, "secondary_server_type", dataSourceName, "result.0.secondary_server_type"),
		resource.TestCheckResourceAttrPair(resourceName, "secondary_state", dataSourceName, "result.0.secondary_state"),
		resource.TestCheckResourceAttrPair(resourceName, "use_failover_port", dataSourceName, "result.0.use_failover_port"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ms_switchover_interval", dataSourceName, "result.0.use_ms_switchover_interval"),
		resource.TestCheckResourceAttrPair(resourceName, "use_recycle_leases", dataSourceName, "result.0.use_recycle_leases"),
	}
}

func testAccDhcpfailoverDataSourceConfigFilters(name string, primary string, secondary string, primaryServerType string, secondaryServerType string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test" {
  name = %q
  primary = %q
  secondary = %q
  primary_server_type = %q
  secondary_server_type = %q
}
data "nios_dhcp_dhcpfailover" "test" {
  filters = {
    name = nios_dhcp_dhcpfailover.test.name
  }
}
`, name, primary, secondary, primaryServerType, secondaryServerType)
}

func testAccDhcpfailoverDataSourceConfigExtAttrFilters(name string, primary string, secondary string, primaryServerType string, secondaryServerType string, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_dhcpfailover" "test" {
  name = %q
  primary = %q
  secondary = %q
  primary_server_type = %q
  secondary_server_type = %q
  extattrs = {
    Site = %q
  } 
}

data "nios_dhcp_dhcpfailover" "test" {
  extattrfilters = {
    Site = nios_dhcp_dhcpfailover.test.extattrs.Site
  }
}
`, name, primary, secondary, primaryServerType, secondaryServerType, extAttrsValue)
}
