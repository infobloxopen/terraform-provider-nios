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

func TestAccDtcRecordSrvDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dtc_record_srv.test"
	resourceName := "nios_dtc_record_srv.test"
	var v dtc.DtcRecordSrv
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcRecordSrvDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcRecordSrvDataSourceConfigFilters(24, 10, "infoblox.com", 30, serverName),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					}, testAccCheckDtcRecordSrvResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDtcRecordSrvResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "dtc_server", dataSourceName, "result.0.dtc_server"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "port", dataSourceName, "result.0.port"),
		resource.TestCheckResourceAttrPair(resourceName, "priority", dataSourceName, "result.0.priority"),
		resource.TestCheckResourceAttrPair(resourceName, "target", dataSourceName, "result.0.target"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "weight", dataSourceName, "result.0.weight"),
	}
}

func testAccDtcRecordSrvDataSourceConfigFilters(port, priority int, target string, weight int, serverName string) string {
    config := fmt.Sprintf(`
resource "nios_dtc_record_srv" "test" {
    port       = %d
    priority   = %d
    target     = %q
    weight     = %d
    dtc_server = nios_dtc_server.test.name
}

data "nios_dtc_record_srv" "test" {
    filters = {
        dtc_server = nios_dtc_server.test.name
        port       = %d
        priority   = %d
        target     = %q
        weight     = %d
    }
    depends_on = [nios_dtc_record_srv.test]
}
`, port, priority, target, weight, port, priority, target, weight)
    return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.2.2.2"), config}, "")
}
