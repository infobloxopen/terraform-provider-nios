package dhcp_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

func TestAccFilterfingerprintDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_filterfingerprint.test"
	resourceName := "nios_dhcp_filterfingerprint.test"
	var v dhcp.Filterfingerprint
	name := acctest.RandomNameWithPrefix("filter-fingerprint")
	fingerprint1 := acctest.RandomNameWithPrefix("fingerprint")
	fingerprint2 := acctest.RandomNameWithPrefix("fingerprint")
	fingerprints := []string{
		"${nios_dhcp_fingerprint.test1.name}",
		"${nios_dhcp_fingerprint.test2.name}",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFilterfingerprintDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFilterfingerprintDataSourceConfigFilters(fingerprint1, fingerprint2, fingerprints, name),
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
	name := acctest.RandomNameWithPrefix("filter-fingerprint")
	fingerprint1 := acctest.RandomNameWithPrefix("fingerprint")
	fingerprint2 := acctest.RandomNameWithPrefix("fingerprint")
	extAttrsValue := acctest.RandomName()
	fingerprints := []string{
		"${nios_dhcp_fingerprint.test1.name}",
		"${nios_dhcp_fingerprint.test2.name}",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFilterfingerprintDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFilterfingerprintDataSourceConfigExtAttrFilters(fingerprint1, fingerprint2, fingerprints, name, extAttrsValue),
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
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "fingerprint", dataSourceName, "result.0.fingerprint"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
	}
}

func testAccFilterfingerprintDataSourceConfigFilters(fingerprint1, fingerprint2 string, fingerprints []string, name string) string {
	fingerprintsStr := utils.ConvertStringSliceToHCL(fingerprints)
	config := fmt.Sprintf(`
resource "nios_dhcp_filterfingerprint" "test" {
  fingerprint = %s
  name = %q
}

data "nios_dhcp_filterfingerprint" "test" {
  filters = {
	name = nios_dhcp_filterfingerprint.test.name
  }
}
`, fingerprintsStr, name)
	return strings.Join([]string{testAccBaseWithFingerprint(fingerprint1, fingerprint2), config}, "")
}

func testAccFilterfingerprintDataSourceConfigExtAttrFilters(fingerprint1, fingerprint2 string, fingerprints []string, name, extAttrsValue string) string {
	fingerprintsStr := utils.ConvertStringSliceToHCL(fingerprints)
	config := fmt.Sprintf(`
resource "nios_dhcp_filterfingerprint" "test" {
  fingerprint = %s
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
`, fingerprintsStr, name, extAttrsValue)
	return strings.Join([]string{testAccBaseWithFingerprint(fingerprint1, fingerprint2), config}, "")
}
