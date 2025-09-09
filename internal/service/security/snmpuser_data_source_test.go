package security_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccSnmpuserDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_security_snmpuser.test"
	resourceName := "nios_security_snmpuser.test"
	var v security.Snmpuser

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpuserDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpuserDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					}, testAccCheckSnmpuserResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccSnmpuserDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_security_snmpuser.test"
	resourceName := "nios_security_snmpuser.test"
	var v security.Snmpuser
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpuserDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpuserDataSourceConfigExtAttrFilters("value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSnmpuserExists(context.Background(), resourceName, &v),
					}, testAccCheckSnmpuserResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckSnmpuserResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "authentication_password", dataSourceName, "result.0.authentication_password"),
		resource.TestCheckResourceAttrPair(resourceName, "authentication_protocol", dataSourceName, "result.0.authentication_protocol"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "privacy_password", dataSourceName, "result.0.privacy_password"),
		resource.TestCheckResourceAttrPair(resourceName, "privacy_protocol", dataSourceName, "result.0.privacy_protocol"),
	}
}

func testAccSnmpuserDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_security_snmpuser" "test" {
}

data "nios_security_snmpuser" "test" {
  filters = {
	 = nios_security_snmpuser.test.
  }
}
`)
}

func testAccSnmpuserDataSourceConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_security_snmpuser" "test" {
  extattrs = {
    Site = %q
  } 
}

data "nios_security_snmpuser" "test" {
  extattrfilters = {
	"Site" = nios_security_snmpuser.test.extattrs.Site
  }
}
`, extAttrsValue)
}
