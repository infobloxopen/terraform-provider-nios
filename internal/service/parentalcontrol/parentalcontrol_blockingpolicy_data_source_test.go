package parentalcontrol_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/parentalcontrol"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccParentalcontrolBlockingpolicyDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_parentalcontrol_blockingpolicy.test"
	resourceName := "nios_parentalcontrol_blockingpolicy.test"
	var v parentalcontrol.ParentalcontrolBlockingpolicy

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckParentalcontrolBlockingpolicyDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccParentalcontrolBlockingpolicyDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckParentalcontrolBlockingpolicyExists(context.Background(), resourceName, &v),
					}, testAccCheckParentalcontrolBlockingpolicyResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckParentalcontrolBlockingpolicyResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "value", dataSourceName, "result.0.value"),
	}
}

func testAccParentalcontrolBlockingpolicyDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_blockingpolicy" "test" {
}

data "nios_parentalcontrol_blockingpolicy" "test" {
  filters = {
	 = nios_parentalcontrol_blockingpolicy.test.
  }
}
`)
}
