
package parentalcontrol_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/parentalcontrol"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccParentalcontrolSubscribersiteDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_parentalcontrol_subscribersite.test"
	resourceName := "nios_parentalcontrol_subscribersite.test"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckParentalcontrolSubscribersiteDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccParentalcontrolSubscribersiteDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
						}, testAccCheckParentalcontrolSubscribersiteResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccParentalcontrolSubscribersiteDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_parentalcontrol_subscribersite.test"
	resourceName := "nios_parentalcontrol_subscribersite.test"
	var v parentalcontrol.ParentalcontrolSubscribersite
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckParentalcontrolSubscribersiteDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccParentalcontrolSubscribersiteDataSourceConfigExtAttrFilters( acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
						}, testAccCheckParentalcontrolSubscribersiteResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckParentalcontrolSubscribersiteResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
    return []resource.TestCheckFunc{
        resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "abss", dataSourceName, "result.0.abss"),
        resource.TestCheckResourceAttrPair(resourceName, "api_members", dataSourceName, "result.0.api_members"),
        resource.TestCheckResourceAttrPair(resourceName, "api_port", dataSourceName, "result.0.api_port"),
        resource.TestCheckResourceAttrPair(resourceName, "block_size", dataSourceName, "result.0.block_size"),
        resource.TestCheckResourceAttrPair(resourceName, "blocking_ipv4_vip1", dataSourceName, "result.0.blocking_ipv4_vip1"),
        resource.TestCheckResourceAttrPair(resourceName, "blocking_ipv4_vip2", dataSourceName, "result.0.blocking_ipv4_vip2"),
        resource.TestCheckResourceAttrPair(resourceName, "blocking_ipv6_vip1", dataSourceName, "result.0.blocking_ipv6_vip1"),
        resource.TestCheckResourceAttrPair(resourceName, "blocking_ipv6_vip2", dataSourceName, "result.0.blocking_ipv6_vip2"),
        resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
        resource.TestCheckResourceAttrPair(resourceName, "dca_sub_bw_list", dataSourceName, "result.0.dca_sub_bw_list"),
        resource.TestCheckResourceAttrPair(resourceName, "dca_sub_query_count", dataSourceName, "result.0.dca_sub_query_count"),
        resource.TestCheckResourceAttrPair(resourceName, "enable_global_allow_list_rpz", dataSourceName, "result.0.enable_global_allow_list_rpz"),
        resource.TestCheckResourceAttrPair(resourceName, "enable_rpz_filtering_bypass", dataSourceName, "result.0.enable_rpz_filtering_bypass"),
        resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
        resource.TestCheckResourceAttrPair(resourceName, "first_port", dataSourceName, "result.0.first_port"),
        resource.TestCheckResourceAttrPair(resourceName, "global_allow_list_rpz", dataSourceName, "result.0.global_allow_list_rpz"),
        resource.TestCheckResourceAttrPair(resourceName, "maximum_subscribers", dataSourceName, "result.0.maximum_subscribers"),
        resource.TestCheckResourceAttrPair(resourceName, "members", dataSourceName, "result.0.members"),
        resource.TestCheckResourceAttrPair(resourceName, "msps", dataSourceName, "result.0.msps"),
        resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
        resource.TestCheckResourceAttrPair(resourceName, "nas_gateways", dataSourceName, "result.0.nas_gateways"),
        resource.TestCheckResourceAttrPair(resourceName, "nas_port", dataSourceName, "result.0.nas_port"),
        resource.TestCheckResourceAttrPair(resourceName, "proxy_rpz_passthru", dataSourceName, "result.0.proxy_rpz_passthru"),
        resource.TestCheckResourceAttrPair(resourceName, "spms", dataSourceName, "result.0.spms"),
        resource.TestCheckResourceAttrPair(resourceName, "stop_anycast", dataSourceName, "result.0.stop_anycast"),
        resource.TestCheckResourceAttrPair(resourceName, "strict_nat", dataSourceName, "result.0.strict_nat"),
        resource.TestCheckResourceAttrPair(resourceName, "subscriber_collection_type", dataSourceName, "result.0.subscriber_collection_type"),
    }
}

func testAccParentalcontrolSubscribersiteDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test" {
}

data "nios_parentalcontrol_subscribersite" "test" {
  filters = {
	 = nios_parentalcontrol_subscribersite.test.
  }
}
`)
}

func testAccParentalcontrolSubscribersiteDataSourceConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test" {
  extattrs = {
    Site = %q
  } 
}

data "nios_parentalcontrol_subscribersite" "test" {
  extattrfilters = {
	Site = nios_parentalcontrol_subscribersite.test.extattrs.Site
  }
}
`,extAttrsValue)
}

