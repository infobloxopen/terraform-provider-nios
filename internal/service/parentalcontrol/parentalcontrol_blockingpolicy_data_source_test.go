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
	name := acctest.RandomNameWithPrefix("blockingpolicy-")
	value := "0000000abcd0000000abcdef00000001"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckParentalcontrolBlockingpolicyDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccParentalcontrolBlockingpolicyDataSourceConfigFilters(name, value),
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

func testAccParentalcontrolBlockingpolicyDataSourceConfigFilters(name, value string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_blockingpolicy" "test" {
    name = %q
    value = %q
}

data "nios_parentalcontrol_blockingpolicy" "test" {
  filters = {
	 name = nios_parentalcontrol_blockingpolicy.test.name
  }
}
`, name, value)
}
