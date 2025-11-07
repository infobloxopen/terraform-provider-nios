package parentalcontrol_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/parentalcontrol"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccParentalcontrolSubscriberrecordDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_parentalcontrol_subscriberrecord.test"
	resourceName := "nios_parentalcontrol_subscriberrecord.test"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckParentalcontrolSubscriberrecordDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccParentalcontrolSubscriberrecordDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					}, testAccCheckParentalcontrolSubscriberrecordResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckParentalcontrolSubscriberrecordResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "accounting_session_id", dataSourceName, "result.0.accounting_session_id"),
		resource.TestCheckResourceAttrPair(resourceName, "alt_ip_addr", dataSourceName, "result.0.alt_ip_addr"),
		resource.TestCheckResourceAttrPair(resourceName, "ans0", dataSourceName, "result.0.ans0"),
		resource.TestCheckResourceAttrPair(resourceName, "ans1", dataSourceName, "result.0.ans1"),
		resource.TestCheckResourceAttrPair(resourceName, "ans2", dataSourceName, "result.0.ans2"),
		resource.TestCheckResourceAttrPair(resourceName, "ans3", dataSourceName, "result.0.ans3"),
		resource.TestCheckResourceAttrPair(resourceName, "ans4", dataSourceName, "result.0.ans4"),
		resource.TestCheckResourceAttrPair(resourceName, "black_list", dataSourceName, "result.0.black_list"),
		resource.TestCheckResourceAttrPair(resourceName, "bwflag", dataSourceName, "result.0.bwflag"),
		resource.TestCheckResourceAttrPair(resourceName, "dynamic_category_policy", dataSourceName, "result.0.dynamic_category_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "flags", dataSourceName, "result.0.flags"),
		resource.TestCheckResourceAttrPair(resourceName, "ip_addr", dataSourceName, "result.0.ip_addr"),
		resource.TestCheckResourceAttrPair(resourceName, "ipsd", dataSourceName, "result.0.ipsd"),
		resource.TestCheckResourceAttrPair(resourceName, "localid", dataSourceName, "result.0.localid"),
		resource.TestCheckResourceAttrPair(resourceName, "nas_contextual", dataSourceName, "result.0.nas_contextual"),
		resource.TestCheckResourceAttrPair(resourceName, "op_code", dataSourceName, "result.0.op_code"),
		resource.TestCheckResourceAttrPair(resourceName, "parental_control_policy", dataSourceName, "result.0.parental_control_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "prefix", dataSourceName, "result.0.prefix"),
		resource.TestCheckResourceAttrPair(resourceName, "proxy_all", dataSourceName, "result.0.proxy_all"),
		resource.TestCheckResourceAttrPair(resourceName, "site", dataSourceName, "result.0.site"),
		resource.TestCheckResourceAttrPair(resourceName, "subscriber_id", dataSourceName, "result.0.subscriber_id"),
		resource.TestCheckResourceAttrPair(resourceName, "subscriber_secure_policy", dataSourceName, "result.0.subscriber_secure_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "unknown_category_policy", dataSourceName, "result.0.unknown_category_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "white_list", dataSourceName, "result.0.white_list"),
		resource.TestCheckResourceAttrPair(resourceName, "wpc_category_policy", dataSourceName, "result.0.wpc_category_policy"),
	}
}

func testAccParentalcontrolSubscriberrecordDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test" {
}

data "nios_parentalcontrol_subscriberrecord" "test" {
  filters = {
	 = nios_parentalcontrol_subscriberrecord.test.
  }
}
`)
}
