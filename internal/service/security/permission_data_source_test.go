package security_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccPermissionDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_security_permission.test"
	resourceName := "nios_security_permission.test"
	var v security.Permission

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPermissionDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccPermissionDataSourceConfigFilters("cloud-api-only", "HOST", "READ"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckPermissionExists(context.Background(), resourceName, &v),
					}, testAccCheckPermissionResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckPermissionResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "group", dataSourceName, "result.0.group"),
		resource.TestCheckResourceAttrPair(resourceName, "object", dataSourceName, "result.0.object"),
		resource.TestCheckResourceAttrPair(resourceName, "permission", dataSourceName, "result.0.permission"),
		resource.TestCheckResourceAttrPair(resourceName, "resource_type", dataSourceName, "result.0.resource_type"),
		resource.TestCheckResourceAttrPair(resourceName, "role", dataSourceName, "result.0.role"),
	}
}

func testAccPermissionDataSourceConfigFilters(group, resourceType, permission string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_view" {
    name = "test-view-ds"
}
resource "nios_security_permission" "test" {
    group = %q
    resource_type = %q
	permission = %q
}
data "nios_security_permission" "test" {
    filters = {
        group = nios_security_permission.test.group
        resource_type = nios_security_permission.test.resource_type
        permission = nios_security_permission.test.permission
    }
}
`, group, resourceType, permission)
}
