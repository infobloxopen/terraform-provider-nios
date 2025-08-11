package security_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccAdminuserDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_security_admin_user.test"
	resourceName := "nios_security_admin_user.test"
	var v security.Adminuser
	name := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAdminuserDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAdminuserDataSourceConfigFilters(name, "ExamplePassword12!", `{"admin-group"}`),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					}, testAccCheckAdminuserResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccAdminuserDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_security_admin_user.test"
	resourceName := "nios_security_admin_user.test"
	extAttrValue := acctest.RandomName()
	name := acctest.RandomName()
	var v security.Adminuser
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAdminuserDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAdminuserDataSourceConfigExtAttrFilters(name, "ExamplePassword12!", `{"admin-group"}`, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					}, testAccCheckAdminuserResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckAdminuserResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "admin_groups", dataSourceName, "result.0.admin_groups"),
		resource.TestCheckResourceAttrPair(resourceName, "auth_method", dataSourceName, "result.0.auth_method"),
		resource.TestCheckResourceAttrPair(resourceName, "auth_type", dataSourceName, "result.0.auth_type"),
		resource.TestCheckResourceAttrPair(resourceName, "ca_certificate_issuer", dataSourceName, "result.0.ca_certificate_issuer"),
		resource.TestCheckResourceAttrPair(resourceName, "client_certificate_serial_number", dataSourceName, "result.0.client_certificate_serial_number"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "email", dataSourceName, "result.0.email"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_certificate_authentication", dataSourceName, "result.0.enable_certificate_authentication"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "password", dataSourceName, "result.0.password"),
		resource.TestCheckResourceAttrPair(resourceName, "ssh_keys", dataSourceName, "result.0.ssh_keys"),
		resource.TestCheckResourceAttrPair(resourceName, "status", dataSourceName, "result.0.status"),
		resource.TestCheckResourceAttrPair(resourceName, "time_zone", dataSourceName, "result.0.time_zone"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ssh_keys", dataSourceName, "result.0.use_ssh_keys"),
		resource.TestCheckResourceAttrPair(resourceName, "use_time_zone", dataSourceName, "result.0.use_time_zone"),
	}
}

func testAccAdminuserDataSourceConfigFilters(name, password string, adminGroups string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test" {
	  name            = %q
	  password        = %q
	  admin_groups    = %q
}

data "nios_security_admin_user" "test" {
  filters = {
	 name = nios_security_admin_user.test.name
  }
}
`, name, password, adminGroups)
}

func testAccAdminuserDataSourceConfigExtAttrFilters(name, password, adminGroups, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test" {
  name = %q
  password = %q
  admin_groups = %q
  extattrs = {
    Site = %q
  } 
}

data "nios_security_admin_user" "test" {
  extattrfilters = {
	Site = nios_security_admin_user.test.extattrs.Site
  }
}
`, name, password, adminGroups, extAttrsValue)
}
