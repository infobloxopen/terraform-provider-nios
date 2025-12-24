package dhcp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

func TestAccFingerprintDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_fingerprint.test"
	resourceName := "nios_dhcp_fingerprint.test"
	var v dhcp.Fingerprint
	name := acctest.RandomNameWithPrefix("fingerprint")
	optionSequence := []string{"1,2,3,4,5,99,199,200"}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFingerprintDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFingerprintDataSourceConfigFilters("Windows OS", name, optionSequence),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFingerprintExists(context.Background(), resourceName, &v),
					}, testAccCheckFingerprintResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccFingerprintDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_fingerprint.test"
	resourceName := "nios_dhcp_fingerprint.test"
	var v dhcp.Fingerprint
	name := acctest.RandomNameWithPrefix("fingerprint")
	optionSequence := []string{"1,2,3,4,5,99,199,200,250"}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFingerprintDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFingerprintDataSourceConfigExtAttrFilters("Windows OS", name, optionSequence, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFingerprintExists(context.Background(), resourceName, &v),
					}, testAccCheckFingerprintResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckFingerprintResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "device_class", dataSourceName, "result.0.device_class"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6_option_sequence", dataSourceName, "result.0.ipv6_option_sequence"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "option_sequence", dataSourceName, "result.0.option_sequence"),
		resource.TestCheckResourceAttrPair(resourceName, "type", dataSourceName, "result.0.type"),
		resource.TestCheckResourceAttrPair(resourceName, "vendor_id", dataSourceName, "result.0.vendor_id"),
	}
}

func testAccFingerprintDataSourceConfigFilters(deviceClass, name string, optionSequence []string) string {
	optionSequenceStr := utils.ConvertStringSliceToHCL(optionSequence)
	return fmt.Sprintf(`
resource "nios_dhcp_fingerprint" "test" {
  device_class = %q
  name = %q
  option_sequence = %s
}

data "nios_dhcp_fingerprint" "test" {
  filters = {
	device_class = nios_dhcp_fingerprint.test.device_class
	name = nios_dhcp_fingerprint.test.name
  }
}
`, deviceClass, name, optionSequenceStr)
}

func testAccFingerprintDataSourceConfigExtAttrFilters(deviceClass, name string, optionSequence []string, extAttrsValue string) string {
	optionSequenceStr := utils.ConvertStringSliceToHCL(optionSequence)
	return fmt.Sprintf(`
resource "nios_dhcp_fingerprint" "test" {
  device_class = %q
  name = %q
  option_sequence = %s
  extattrs = {
    Site = %q
  } 
}

data "nios_dhcp_fingerprint" "test" {
  extattrfilters = {
    Site = nios_dhcp_fingerprint.test.extattrs.Site
  }
}
`, deviceClass, name, optionSequenceStr, extAttrsValue)
}
