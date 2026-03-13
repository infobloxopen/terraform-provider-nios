package microsoft_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/microsoft"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccMsserverDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_microsoft_msserver.test"
	resourceName := "nios_microsoft_msserver.test"
	var v microsoft.Msserver
	address := "10.10.0.1"
	loginName := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMsserverDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMsserverDataSourceConfigFilters(address, loginName),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckMsserverExists(context.Background(), resourceName, &v),
					}, testAccCheckMsserverResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccMsserverDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_microsoft_msserver.test"
	resourceName := "nios_microsoft_msserver.test"
	var v microsoft.Msserver
	address := "10.10.0.1"
	loginName := acctest.RandomName()
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMsserverDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMsserverDataSourceConfigExtAttrFilters(address, loginName, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckMsserverExists(context.Background(), resourceName, &v),
					}, testAccCheckMsserverResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckMsserverResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "ad_domain", dataSourceName, "result.0.ad_domain"),
		resource.TestCheckResourceAttrPair(resourceName, "ad_sites", dataSourceName, "result.0.ad_sites"),
		resource.TestCheckResourceAttrPair(resourceName, "ad_user", dataSourceName, "result.0.ad_user"),
		resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "result.0.address"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "connection_status", dataSourceName, "result.0.connection_status"),
		resource.TestCheckResourceAttrPair(resourceName, "connection_status_detail", dataSourceName, "result.0.connection_status_detail"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_server", dataSourceName, "result.0.dhcp_server"),
		resource.TestCheckResourceAttrPair(resourceName, "disabled", dataSourceName, "result.0.disabled"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_server", dataSourceName, "result.0.dns_server"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_view", dataSourceName, "result.0.dns_view"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "grid_member", dataSourceName, "result.0.grid_member"),
		resource.TestCheckResourceAttrPair(resourceName, "last_seen", dataSourceName, "result.0.last_seen"),
		resource.TestCheckResourceAttrPair(resourceName, "log_destination", dataSourceName, "result.0.log_destination"),
		resource.TestCheckResourceAttrPair(resourceName, "log_level", dataSourceName, "result.0.log_level"),
		resource.TestCheckResourceAttrPair(resourceName, "login_name", dataSourceName, "result.0.login_name"),
		resource.TestCheckResourceAttrPair(resourceName, "login_password", dataSourceName, "result.0.login_password"),
		resource.TestCheckResourceAttrPair(resourceName, "managing_member", dataSourceName, "result.0.managing_member"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_max_connection", dataSourceName, "result.0.ms_max_connection"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_rpc_timeout_in_seconds", dataSourceName, "result.0.ms_rpc_timeout_in_seconds"),
		resource.TestCheckResourceAttrPair(resourceName, "network_view", dataSourceName, "result.0.network_view"),
		resource.TestCheckResourceAttrPair(resourceName, "read_only", dataSourceName, "result.0.read_only"),
		resource.TestCheckResourceAttrPair(resourceName, "root_ad_domain", dataSourceName, "result.0.root_ad_domain"),
		resource.TestCheckResourceAttrPair(resourceName, "server_name", dataSourceName, "result.0.server_name"),
		resource.TestCheckResourceAttrPair(resourceName, "synchronization_min_delay", dataSourceName, "result.0.synchronization_min_delay"),
		resource.TestCheckResourceAttrPair(resourceName, "synchronization_status", dataSourceName, "result.0.synchronization_status"),
		resource.TestCheckResourceAttrPair(resourceName, "synchronization_status_detail", dataSourceName, "result.0.synchronization_status_detail"),
		resource.TestCheckResourceAttrPair(resourceName, "use_log_destination", dataSourceName, "result.0.use_log_destination"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ms_max_connection", dataSourceName, "result.0.use_ms_max_connection"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ms_rpc_timeout_in_seconds", dataSourceName, "result.0.use_ms_rpc_timeout_in_seconds"),
		resource.TestCheckResourceAttrPair(resourceName, "version", dataSourceName, "result.0.version"),
	}
}

func testAccMsserverDataSourceConfigFilters(address, loginName string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test" {
  address = %q
  login_name = %q
}

data "nios_microsoft_msserver" "test" {
  filters = {
	address = nios_microsoft_msserver.test.address
  }
}
`, address, loginName)
}

func testAccMsserverDataSourceConfigExtAttrFilters(address, loginName, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test" {
  address = %q
  login_name = %q
  extattrs = {
    Site = %q
  } 
}

data "nios_microsoft_msserver" "test" {
  extattrfilters = {
	Site = nios_microsoft_msserver.test.extattrs.Site
  }
}
`, address, loginName, extAttrsValue)
}
