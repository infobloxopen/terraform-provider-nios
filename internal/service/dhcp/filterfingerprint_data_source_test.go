package dhcp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

/*
// Retrieve a specific dhcp Filterfingerprint by filters
data "nios_dhcp_filterfingerprint" "get_dhcp_filterfingerprint_using_filters" {
  filters = {
    fingerprint = "FINGERPRINT_REPLACE_ME"
    name = "NAME_REPLACE_ME"
  }
}
// Retrieve specific dhcp Filterfingerprint using Extensible Attributes
data "nios_dhcp_filterfingerprint" "get_dhcp_filterfingerprint_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all dhcp Filterfingerprint
data "nios_dhcp_filterfingerprint" "get_all_dhcp_filterfingerprint" {}
*/

func TestAccFilterfingerprintDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_filterfingerprint.test"
	resourceName := "nios_dhcp_filterfingerprint.test"
	var v dhcp.Filterfingerprint

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFilterfingerprintDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFilterfingerprintDataSourceConfigFilters("FINGERPRINT_REPLACE_ME", "NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFilterfingerprintExists(context.Background(), resourceName, &v),
					}, testAccCheckFilterfingerprintResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccFilterfingerprintDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_filterfingerprint.test"
	resourceName := "nios_dhcp_filterfingerprint.test"
	var v dhcp.Filterfingerprint

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFilterfingerprintDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFilterfingerprintDataSourceConfigExtAttrFilters("FINGERPRINT_REPLACE_ME", "NAME_REPLACE_ME", acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFilterfingerprintExists(context.Background(), resourceName, &v),
					}, testAccCheckFilterfingerprintResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckFilterfingerprintResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "fingerprint", dataSourceName, "result.0.fingerprint"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
	}
}

func testAccFilterfingerprintDataSourceConfigFilters(fingerprint, name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filterfingerprint" "test" {
  fingerprint = %q
  name = %q
}

data "nios_dhcp_filterfingerprint" "test" {
  filters = {
	fingerprint = nios_dhcp_filterfingerprint.test.fingerprint
  }
}
`, fingerprint, name)
}

func testAccFilterfingerprintDataSourceConfigExtAttrFilters(fingerprint, name, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filterfingerprint" "test" {
  fingerprint = %q
  name = %q
  extattrs = {
    Site = %q
  } 
}

data "nios_dhcp_filterfingerprint" "test" {
  extattrfilters = {
    Site = nios_dhcp_filterfingerprint.test.extattrs.Site
  }
}
`, fingerprint, name, extAttrsValue)
}
