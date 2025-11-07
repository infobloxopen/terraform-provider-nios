package microsoftserver_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/microsoftserver"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccMsserverAdsitesSiteDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_microsoftserver_msserver_adsites_site.test"
	resourceName := "nios_microsoftserver_msserver_adsites_site.test"
	var v microsoftserver.MsserverAdsitesSite

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMsserverAdsitesSiteDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMsserverAdsitesSiteDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckMsserverAdsitesSiteExists(context.Background(), resourceName, &v),
					}, testAccCheckMsserverAdsitesSiteResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckMsserverAdsitesSiteResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "domain", dataSourceName, "result.0.domain"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "networks", dataSourceName, "result.0.networks"),
	}
}

func testAccMsserverAdsitesSiteDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_adsites_site" "test" {
}

data "nios_microsoftserver_msserver_adsites_site" "test" {
  filters = {
	 = nios_microsoftserver_msserver_adsites_site.test.
  }
}
`)
}
