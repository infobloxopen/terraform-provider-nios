package acl_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/acl"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

func TestAccNamedaclDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_acl_namedacl.test"
	resourceName := "nios_acl_namedacl.test"
	var v acl.Namedacl
	name := acctest.RandomNameWithPrefix("namedacl")
	acl := []map[string]any{
		{
			"struct":     "addressac",
			"address":    "192.168.1.10",
			"permission": "ALLOW",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNamedaclDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNamedaclDataSourceConfigFilters(name, acl),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckNamedaclExists(context.Background(), resourceName, &v),
					}, testAccCheckNamedaclResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccNamedaclDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_acl_namedacl.test"
	resourceName := "nios_acl_namedacl.test"
	var v acl.Namedacl
	name := acctest.RandomNameWithPrefix("namedacl")
	extAttrValue := acctest.RandomName()
	acl := []map[string]any{
		{
			"struct":     "addressac",
			"address":    "10.0.0.1",
			"permission": "ALLOW",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNamedaclDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNamedaclDataSourceConfigExtAttrFilters(name, extAttrValue, acl),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckNamedaclExists(context.Background(), resourceName, &v),
					}, testAccCheckNamedaclResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckNamedaclResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "access_list", dataSourceName, "result.0.access_list"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "exploded_access_list", dataSourceName, "result.0.exploded_access_list"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
	}
}

func testAccNamedaclDataSourceConfigFilters(name string, acl []map[string]any) string {
	aclHCL := utils.ConvertSliceOfMapsToHCL(acl)
	return fmt.Sprintf(`
resource "nios_acl_namedacl" "test" {
	name = %q
	access_list = %s
}

data "nios_acl_namedacl" "test" {
  filters = {
	name = nios_acl_namedacl.test.name
  }
}
`, name, aclHCL)
}

func testAccNamedaclDataSourceConfigExtAttrFilters(name string, extAttrsValue string, acl []map[string]any) string {
	aclHCL := utils.ConvertSliceOfMapsToHCL(acl)
	return fmt.Sprintf(`
resource "nios_acl_namedacl" "test" {
	name = %q
	access_list = %s
	extattrs = {
		Site = %q
	}
}

data "nios_acl_namedacl" "test" {
  extattrfilters = {
	Site = nios_acl_namedacl.test.extattrs.Site
  }
}
`, name, aclHCL, extAttrsValue)
}
