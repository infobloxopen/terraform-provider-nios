
package security_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccAdmingroupDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_security_admingroup.test"
	resourceName := "nios_security_admingroup.test"
	var v security.Admingroup

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAdmingroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAdmingroupDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
						}, testAccCheckAdmingroupResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccAdmingroupDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_security_admingroup.test"
	resourceName := "nios_security_admingroup.test"
	var v security.Admingroup
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAdmingroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAdmingroupDataSourceConfigExtAttrFilters(, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
						}, testAccCheckAdmingroupResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckAdmingroupResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
    return []resource.TestCheckFunc{
        resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "access_method", dataSourceName, "result.0.access_method"),
        resource.TestCheckResourceAttrPair(resourceName, "admin_set_commands", dataSourceName, "result.0.admin_set_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "admin_show_commands", dataSourceName, "result.0.admin_show_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "admin_toplevel_commands", dataSourceName, "result.0.admin_toplevel_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "cloud_set_commands", dataSourceName, "result.0.cloud_set_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "cloud_show_commands", dataSourceName, "result.0.cloud_show_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
        resource.TestCheckResourceAttrPair(resourceName, "database_set_commands", dataSourceName, "result.0.database_set_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "database_show_commands", dataSourceName, "result.0.database_show_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "dhcp_set_commands", dataSourceName, "result.0.dhcp_set_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "dhcp_show_commands", dataSourceName, "result.0.dhcp_show_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
        resource.TestCheckResourceAttrPair(resourceName, "disable_concurrent_login", dataSourceName, "result.0.disable_concurrent_login"),
        resource.TestCheckResourceAttrPair(resourceName, "dns_set_commands", dataSourceName, "result.0.dns_set_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "dns_show_commands", dataSourceName, "result.0.dns_show_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "dns_toplevel_commands", dataSourceName, "result.0.dns_toplevel_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "docker_set_commands", dataSourceName, "result.0.docker_set_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "docker_show_commands", dataSourceName, "result.0.docker_show_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "email_addresses", dataSourceName, "result.0.email_addresses"),
        resource.TestCheckResourceAttrPair(resourceName, "enable_restricted_user_access", dataSourceName, "result.0.enable_restricted_user_access"),
        resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
        resource.TestCheckResourceAttrPair(resourceName, "grid_set_commands", dataSourceName, "result.0.grid_set_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "grid_show_commands", dataSourceName, "result.0.grid_show_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "inactivity_lockout_setting", dataSourceName, "result.0.inactivity_lockout_setting"),
        resource.TestCheckResourceAttrPair(resourceName, "licensing_set_commands", dataSourceName, "result.0.licensing_set_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "licensing_show_commands", dataSourceName, "result.0.licensing_show_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "lockout_setting", dataSourceName, "result.0.lockout_setting"),
        resource.TestCheckResourceAttrPair(resourceName, "machine_control_toplevel_commands", dataSourceName, "result.0.machine_control_toplevel_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
        resource.TestCheckResourceAttrPair(resourceName, "networking_set_commands", dataSourceName, "result.0.networking_set_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "networking_show_commands", dataSourceName, "result.0.networking_show_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "password_setting", dataSourceName, "result.0.password_setting"),
        resource.TestCheckResourceAttrPair(resourceName, "roles", dataSourceName, "result.0.roles"),
        resource.TestCheckResourceAttrPair(resourceName, "saml_setting", dataSourceName, "result.0.saml_setting"),
        resource.TestCheckResourceAttrPair(resourceName, "security_set_commands", dataSourceName, "result.0.security_set_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "security_show_commands", dataSourceName, "result.0.security_show_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "superuser", dataSourceName, "result.0.superuser"),
        resource.TestCheckResourceAttrPair(resourceName, "trouble_shooting_toplevel_commands", dataSourceName, "result.0.trouble_shooting_toplevel_commands"),
        resource.TestCheckResourceAttrPair(resourceName, "use_account_inactivity_lockout_enable", dataSourceName, "result.0.use_account_inactivity_lockout_enable"),
        resource.TestCheckResourceAttrPair(resourceName, "use_disable_concurrent_login", dataSourceName, "result.0.use_disable_concurrent_login"),
        resource.TestCheckResourceAttrPair(resourceName, "use_lockout_setting", dataSourceName, "result.0.use_lockout_setting"),
        resource.TestCheckResourceAttrPair(resourceName, "use_password_setting", dataSourceName, "result.0.use_password_setting"),
        resource.TestCheckResourceAttrPair(resourceName, "user_access", dataSourceName, "result.0.user_access"),
    }
}

func testAccAdmingroupDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_security_admingroup" "test" {
}

data "nios_security_admingroup" "test" {
  filters = {
	 = nios_security_admingroup.test.
  }
}
`)
}

func testAccAdmingroupDataSourceConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_security_admingroup" "test" {
  extattrs = {
    Site = %q
  } 
}

data "nios_security_admingroup" "test" {
  extattrfilters = {
	Site = nios_security_admingroup.test.extattrs.Site
  }
}
`,extAttrsValue)
}

