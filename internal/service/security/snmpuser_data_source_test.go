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
	dataSourceName := "data.nios_security_snmp_user.test"
	resourceName := "nios_security_snmp_user.test"
	var v security.Snmpuser

	name := acctest.RandomNameWithPrefix("example-snmpuser-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpuserDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpuserDataSourceConfigFilters(name, "NONE", "NONE"),
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
	dataSourceName := "data.nios_security_snmp_user.test"
	resourceName := "nios_security_snmp_user.test"
	var v security.Snmpuser

	name := acctest.RandomNameWithPrefix("example-snmpuser-")
	extAttrValue := acctest.RandomNameWithPrefix("snmp-user")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpuserDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpuserDataSourceConfigExtAttrFilters(name, "NONE", "NONE", extAttrValue),
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

func testAccSnmpuserDataSourceConfigFilters(name, authentication_protocol, privacy_protocol string) string {
	return fmt.Sprintf(`
resource "nios_security_snmp_user" "test" {
	  name 					  = %q
	  authentication_protocol = %q
	  privacy_protocol 		  = %q
}

data "nios_security_snmp_user" "test" {
  filters = {
	name  = nios_security_snmp_user.test.name
  }
}
`, name, authentication_protocol, privacy_protocol)
}

func testAccSnmpuserDataSourceConfigExtAttrFilters(name, authentication_protocol, privacy_protocol, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_security_snmp_user" "test" {
  name 					  = %q
  authentication_protocol = %q
  privacy_protocol 		  = %q
  extattrs = {
    Site   = %q
  } 
}

data "nios_security_snmp_user" "test" {
  extattrfilters = {
	Site = nios_security_snmp_user.test.extattrs.Site
  }
}
`, name, authentication_protocol, privacy_protocol, extAttrsValue)
}
