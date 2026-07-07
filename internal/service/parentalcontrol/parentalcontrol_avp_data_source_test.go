package parentalcontrol_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/parentalcontrol"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccParentalcontrolAvpDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_parentalcontrol_avp.test"
	resourceName := "nios_parentalcontrol_avp.test"
	var v parentalcontrol.ParentalcontrolAvp
	name := acctest.RandomNameWithPrefix("parentalcontrol-avp")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckParentalcontrolAvpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccParentalcontrolAvpDataSourceConfigFilters(name, "STRING", 220),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					}, testAccCheckParentalcontrolAvpResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckParentalcontrolAvpResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "domain_types", dataSourceName, "result.0.domain_types"),
		resource.TestCheckResourceAttrPair(resourceName, "is_restricted", dataSourceName, "result.0.is_restricted"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "type", dataSourceName, "result.0.type"),
		resource.TestCheckResourceAttrPair(resourceName, "user_defined", dataSourceName, "result.0.user_defined"),
		resource.TestCheckResourceAttrPair(resourceName, "value_type", dataSourceName, "result.0.value_type"),
		resource.TestCheckResourceAttrPair(resourceName, "vendor_id", dataSourceName, "result.0.vendor_id"),
		resource.TestCheckResourceAttrPair(resourceName, "vendor_type", dataSourceName, "result.0.vendor_type"),
	}
}

func testAccParentalcontrolAvpDataSourceConfigFilters(name, valueType string, type_ int32) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_avp" "test" {
    name = %q
	value_type = %q
	type = %d
}

data "nios_parentalcontrol_avp" "test" {
  filters = {
	 name = nios_parentalcontrol_avp.test.name
  }
}
`, name, valueType, type_)
}
