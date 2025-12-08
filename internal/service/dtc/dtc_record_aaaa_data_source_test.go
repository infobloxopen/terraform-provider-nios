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

func TestAccDtcRecordAaaaDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dtc_record_aaaa.test"
	resourceName := "nios_dtc_record_aaaa.test"
	var v dtc.DtcRecordAaaa
	ipv6Addr := "2001:db8:85a3::8a2e:370:7335"
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcRecordAaaaDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcRecordAaaaDataSourceConfigFilters(ipv6Addr, serverName),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					}, testAccCheckDtcRecordAaaaResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDtcRecordAaaaResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "auto_created", dataSourceName, "result.0.auto_created"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "dtc_server", dataSourceName, "result.0.dtc_server"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6addr", dataSourceName, "result.0.ipv6addr"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
	}
}

func testAccDtcRecordAaaaDataSourceConfigFilters(ipv4Addr string, serverName string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_aaaa" "test" {
	ipv6addr = %q
	dtc_server = nios_dtc_server.test.name
}

data "nios_dtc_record_aaaa" "test" {
  filters = {
	dtc_server = nios_dtc_server.test.name
	ipv6addr = nios_dtc_record_aaaa.test.ipv6addr
  }
}
	`, ipv4Addr)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.2.2.2"), config}, "")
}
