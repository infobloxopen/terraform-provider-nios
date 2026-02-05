package dtc_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccDtcRecordCnameDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dtc_record_cname.test"
	resourceName := "nios_dtc_record_cname.test"
	var v dtc.DtcRecordCname
	name := acctest.RandomNameWithPrefix("dtc-cname")
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcRecordCnameDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcRecordCnameDataSourceConfigFilters(name, serverName),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDtcRecordCnameExists(context.Background(), resourceName, &v),
					}, testAccCheckDtcRecordCnameResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDtcRecordCnameResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "auto_created", dataSourceName, "result.0.auto_created"),
		resource.TestCheckResourceAttrPair(resourceName, "canonical", dataSourceName, "result.0.canonical"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_canonical", dataSourceName, "result.0.dns_canonical"),
		resource.TestCheckResourceAttrPair(resourceName, "dtc_server", dataSourceName, "result.0.dtc_server"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
	}
}

func testAccDtcRecordCnameDataSourceConfigFilters(name, serverName string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_cname" "test" {
	  canonical = %q
	  dtc_server = nios_dtc_server.test.name
}

data "nios_dtc_record_cname" "test" {
  filters = {
	 dtc_server= nios_dtc_record_cname.test.dtc_server
	 canonical= nios_dtc_record_cname.test.canonical
  }
}`, name)
	return strings.Join([]string{testAccBaseWithDtcServerDisable(serverName, "2.2.2.2"), config}, "")
}
