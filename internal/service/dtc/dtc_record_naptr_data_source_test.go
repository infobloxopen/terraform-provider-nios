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

func TestAccDtcRecordNaptrDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dtc_record_naptr.test"
	resourceName := "nios_dtc_record_naptr.test"
	var v dtc.DtcRecordNaptr
	serverName := acctest.RandomNameWithPrefix("dtc-server")
	serverIp := acctest.RandomIP()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcRecordNaptrDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcRecordNaptrDataSourceConfigFilters(serverName, serverIp, 2, 5, "example.com"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					}, testAccCheckDtcRecordNaptrResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDtcRecordNaptrResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "dtc_server", dataSourceName, "result.0.dtc_server"),
		resource.TestCheckResourceAttrPair(resourceName, "flags", dataSourceName, "result.0.flags"),
		resource.TestCheckResourceAttrPair(resourceName, "order", dataSourceName, "result.0.order"),
		resource.TestCheckResourceAttrPair(resourceName, "preference", dataSourceName, "result.0.preference"),
		resource.TestCheckResourceAttrPair(resourceName, "regexp", dataSourceName, "result.0.regexp"),
		resource.TestCheckResourceAttrPair(resourceName, "replacement", dataSourceName, "result.0.replacement"),
		resource.TestCheckResourceAttrPair(resourceName, "services", dataSourceName, "result.0.services"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
	}
}

func testAccDtcRecordNaptrDataSourceConfigFilters(serverName, serverIP string, order, preference int, replacement string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_naptr" "test" {
  		dtc_server = nios_dtc_server.test.name
		order	  = %d
		preference = %d
		replacement = "%s"
	}
	data "nios_dtc_record_naptr" "test" {
  	filters = {
		dtc_server = nios_dtc_server.test.name
		replacement = nios_dtc_record_naptr.test.replacement
	}
}
  	`, order, preference, replacement)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, serverIP), config}, "\n")
}
