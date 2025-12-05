package microsoftserver_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/microsoftserver"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccMsserverDnsDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_microsoftserver_msserver_dns.test"
	resourceName := "nios_microsoftserver_msserver_dns.test"
	var v microsoftserver.MsserverDns

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMsserverDnsDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMsserverDnsDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckMsserverDnsExists(context.Background(), resourceName, &v),
					}, testAccCheckMsserverDnsResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckMsserverDnsResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "result.0.address"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_dns_reports_sync", dataSourceName, "result.0.enable_dns_reports_sync"),
		resource.TestCheckResourceAttrPair(resourceName, "login_name", dataSourceName, "result.0.login_name"),
		resource.TestCheckResourceAttrPair(resourceName, "login_password", dataSourceName, "result.0.login_password"),
		resource.TestCheckResourceAttrPair(resourceName, "synchronization_interval", dataSourceName, "result.0.synchronization_interval"),
		resource.TestCheckResourceAttrPair(resourceName, "use_enable_dns_reports_sync", dataSourceName, "result.0.use_enable_dns_reports_sync"),
		resource.TestCheckResourceAttrPair(resourceName, "use_login", dataSourceName, "result.0.use_login"),
		resource.TestCheckResourceAttrPair(resourceName, "use_synchronization_interval", dataSourceName, "result.0.use_synchronization_interval"),
	}
}

func testAccMsserverDnsDataSourceConfigFilters() string {
	return `
resource "nios_microsoftserver_msserver_dns" "test" {
}

data "nios_microsoftserver_msserver_dns" "test" {
  filters = {
	 = nios_microsoftserver_msserver_dns.test.
  }
}
`
}
