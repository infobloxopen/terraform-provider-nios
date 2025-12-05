package security_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccAdminroleDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_security_admin_role.test"
	resourceName := "nios_security_admin_role.test"
	var v security.Adminrole

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAdminroleDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAdminroleDataSourceConfigFilters(acctest.RandomNameWithPrefix("admin-role")),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckAdminroleExists(context.Background(), resourceName, &v),
					}, testAccCheckAdminroleResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccAdminroleDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_security_admin_role.test"
	resourceName := "nios_security_admin_role.test"
	var v security.Adminrole
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAdminroleDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAdminroleDataSourceConfigExtAttrFilters(acctest.RandomNameWithPrefix("admin-role"), acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckAdminroleExists(context.Background(), resourceName, &v),
					}, testAccCheckAdminroleResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckAdminroleResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
	}
}

func testAccAdminroleDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_role" "test" {
 name = %q
}

data "nios_security_admin_role" "test" {
 filters = {
	name = nios_security_admin_role.test.name
 }
}
`, name)
}

func testAccAdminroleDataSourceConfigExtAttrFilters(name, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_role" "test" {
 name = %q
 extattrs = {
   Site = %q
 }
}

data "nios_security_admin_role" "test" {
 extattrfilters = {
	Site = nios_security_admin_role.test.extattrs.Site
 }
}
`, name, extAttrsValue)
}
