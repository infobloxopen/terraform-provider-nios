package dtc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccDtcServerDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dtc_server.test"
	resourceName := "nios_dtc_server.test"
	var v dtc.DtcServer
	name := acctest.RandomNameWithPrefix("dtc-server")
	host := acctest.RandomIP()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcServerDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcServerDataSourceConfigFilters(name, host),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					}, testAccCheckDtcServerResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccDtcServerDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dtc_server.test"
	resourceName := "nios_dtc_server.test"
	var v dtc.DtcServer
	name := acctest.RandomNameWithPrefix("dtc-server")
	host := acctest.RandomIP()
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcServerDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcServerDataSourceConfigExtAttrFilters(name, host, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					}, testAccCheckDtcServerResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDtcServerResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "auto_create_host_record", dataSourceName, "result.0.auto_create_host_record"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "health", dataSourceName, "result.0.health"),
		resource.TestCheckResourceAttrPair(resourceName, "host", dataSourceName, "result.0.host"),
		resource.TestCheckResourceAttrPair(resourceName, "monitors", dataSourceName, "result.0.monitors"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "sni_hostname", dataSourceName, "result.0.sni_hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "use_sni_hostname", dataSourceName, "result.0.use_sni_hostname"),
	}
}

func testAccDtcServerDataSourceConfigFilters(name, host string) string {
	return fmt.Sprintf(`
resource "nios_dtc_server" "test" {
name = "%s"
host = "%s"
}

data "nios_dtc_server" "test" {
  filters = {
	 name = nios_dtc_server.test.name 
  }
}
`, name, host)
}

func testAccDtcServerDataSourceConfigExtAttrFilters(name, host, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dtc_server" "test" {
	name = %q
	host = %q
  extattrs = {
    Site =  %q
  	}
}

data "nios_dtc_server" "test" {
  extattrfilters = {
	Site = nios_dtc_server.test.extattrs.Site
  }
}
`, name, host, extAttrsValue)
}
