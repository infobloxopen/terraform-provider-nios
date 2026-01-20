package security_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccFtpuserDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_security_ftpuser.test"
	resourceName := "nios_security_ftpuser.test"
	var v security.Ftpuser
	username := acctest.RandomNameWithPrefix("ftf-test-user-")
	password := acctest.RandomAlphaNumeric(12)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFtpuserDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFtpuserDataSourceConfigFilters(username, password),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					}, testAccCheckFtpuserResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccFtpuserDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_security_ftpuser.test"
	resourceName := "nios_security_ftpuser.test"
	var v security.Ftpuser
	extAttrValue := acctest.RandomNameWithPrefix("ftp-user")
	username := acctest.RandomNameWithPrefix("ftf-test-user-")
	password := acctest.RandomAlphaNumeric(12)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFtpuserDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFtpuserDataSourceConfigExtAttrFilters(username, password, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFtpuserExists(context.Background(), resourceName, &v),
					}, testAccCheckFtpuserResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckFtpuserResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "home_dir", dataSourceName, "result.0.home_dir"),
		resource.TestCheckResourceAttrPair(resourceName, "permission", dataSourceName, "result.0.permission"),
		resource.TestCheckResourceAttrPair(resourceName, "username", dataSourceName, "result.0.username"),
	}
}

func testAccFtpuserDataSourceConfigFilters(username, password string) string {
	return fmt.Sprintf(`
resource "nios_security_ftpuser" "test" {
  username = %q
  password = %q
}

data "nios_security_ftpuser" "test" {
  filters = {
 	username = nios_security_ftpuser.test.username
  }
}
`, username, password)
}

func testAccFtpuserDataSourceConfigExtAttrFilters(username, password, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_security_ftpuser" "test" {
  username = %q
  password = %q
  extattrs = {
    Site = %q
  }
}
data "nios_security_ftpuser" "test" {
  extattrfilters = {
	Site = nios_security_ftpuser.test.extattrs.Site
  }
}
`, username, password, extAttrsValue)
}
