package grid_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccExtensibleattributedefDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_grid_extensibleattributedef.test"
	resourceName := "nios_grid_extensibleattributedef.test"
	var v grid.Extensibleattributedef
	name := acctest.RandomNameWithPrefix("tf_test_ea_")
	eaType := "STRING"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckExtensibleattributedefDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccExtensibleattributedefDataSourceConfigFilters(name, eaType),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckExtensibleattributedefExists(context.Background(), resourceName, &v),
					}, testAccCheckExtensibleattributedefResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckExtensibleattributedefResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "allowed_object_types", dataSourceName, "result.0.allowed_object_types"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "default_value", dataSourceName, "result.0.default_value"),
		resource.TestCheckResourceAttrPair(resourceName, "descendants_action", dataSourceName, "result.0.descendants_action"),
		resource.TestCheckResourceAttrPair(resourceName, "flags", dataSourceName, "result.0.flags"),
		resource.TestCheckResourceAttrPair(resourceName, "list_values", dataSourceName, "result.0.list_values"),
		resource.TestCheckResourceAttrPair(resourceName, "max", dataSourceName, "result.0.max"),
		resource.TestCheckResourceAttrPair(resourceName, "min", dataSourceName, "result.0.min"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "namespace", dataSourceName, "result.0.namespace"),
		resource.TestCheckResourceAttrPair(resourceName, "type", dataSourceName, "result.0.type"),
	}
}

func testAccExtensibleattributedefDataSourceConfigFilters(name, eaType string) string {
	return fmt.Sprintf(`
resource "nios_grid_extensibleattributedef" "test" {
    name = %q
    type = %q
}

data "nios_grid_extensibleattributedef" "test" {
    filters = {
        name = nios_grid_extensibleattributedef.test.name
    }
}
`, name, eaType)
}
