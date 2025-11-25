package dtc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccDtcRecordADataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dtc_record_a.test"
	resourceName := "nios_dtc_record_a.test"
	var v dtc.DtcRecordA

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcRecordADestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcRecordADataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					}, testAccCheckDtcRecordAResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDtcRecordAResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "auto_created", dataSourceName, "result.0.auto_created"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "dtc_server", dataSourceName, "result.0.dtc_server"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv4addr", dataSourceName, "result.0.ipv4addr"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
	}
}

func testAccDtcRecordADataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_dtc_record_a" "test" {
}

data "nios_dtc_record_a" "test" {
  filters = {
	 = nios_dtc_record_a.test.
  }
}
`)
}
