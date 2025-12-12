package dhcp_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccMacfilteraddressDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_macfilteraddress.test"
	resourceName := "nios_dhcp_macfilteraddress.test"
	var v dhcp.Macfilteraddress
	mac := "00:1a:2b:3c:3d:5e"
	filter := acctest.RandomNameWithPrefix("mac-filter")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMacfilteraddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMacfilteraddressDataSourceConfigFilters(filter, mac),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					}, testAccCheckMacfilteraddressResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccMacfilteraddressDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_macfilteraddress.test"
	resourceName := "nios_dhcp_macfilteraddress.test"
	var v dhcp.Macfilteraddress
	mac := "00:1a:2b:3c:3d:5e"
	filter := acctest.RandomNameWithPrefix("mac-filter")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMacfilteraddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMacfilteraddressDataSourceConfigExtAttrFilters(filter, mac, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckMacfilteraddressExists(context.Background(), resourceName, &v),
					}, testAccCheckMacfilteraddressResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckMacfilteraddressResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "authentication_time", dataSourceName, "result.0.authentication_time"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "expiration_time", dataSourceName, "result.0.expiration_time"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "filter", dataSourceName, "result.0.filter"),
		resource.TestCheckResourceAttrPair(resourceName, "fingerprint", dataSourceName, "result.0.fingerprint"),
		resource.TestCheckResourceAttrPair(resourceName, "guest_custom_field1", dataSourceName, "result.0.guest_custom_field1"),
		resource.TestCheckResourceAttrPair(resourceName, "guest_custom_field2", dataSourceName, "result.0.guest_custom_field2"),
		resource.TestCheckResourceAttrPair(resourceName, "guest_custom_field3", dataSourceName, "result.0.guest_custom_field3"),
		resource.TestCheckResourceAttrPair(resourceName, "guest_custom_field4", dataSourceName, "result.0.guest_custom_field4"),
		resource.TestCheckResourceAttrPair(resourceName, "guest_email", dataSourceName, "result.0.guest_email"),
		resource.TestCheckResourceAttrPair(resourceName, "guest_first_name", dataSourceName, "result.0.guest_first_name"),
		resource.TestCheckResourceAttrPair(resourceName, "guest_last_name", dataSourceName, "result.0.guest_last_name"),
		resource.TestCheckResourceAttrPair(resourceName, "guest_middle_name", dataSourceName, "result.0.guest_middle_name"),
		resource.TestCheckResourceAttrPair(resourceName, "guest_phone", dataSourceName, "result.0.guest_phone"),
		resource.TestCheckResourceAttrPair(resourceName, "is_registered_user", dataSourceName, "result.0.is_registered_user"),
		resource.TestCheckResourceAttrPair(resourceName, "mac", dataSourceName, "result.0.mac"),
		resource.TestCheckResourceAttrPair(resourceName, "never_expires", dataSourceName, "result.0.never_expires"),
		resource.TestCheckResourceAttrPair(resourceName, "reserved_for_infoblox", dataSourceName, "result.0.reserved_for_infoblox"),
		resource.TestCheckResourceAttrPair(resourceName, "username", dataSourceName, "result.0.username"),
	}
}

func testAccMacfilteraddressDataSourceConfigFilters(filter, mac string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_macfilteraddress" "test" {
  filter = nios_dhcp_filtermac.parent_filter_mac.name
  mac    = %q
}

data "nios_dhcp_macfilteraddress" "test" {
  filters = {
	 filter = nios_dhcp_macfilteraddress.test.filter
  }
}
`, mac)
	return strings.Join([]string{testAccBaseWithMacFilter(filter), config}, "")
}

func testAccMacfilteraddressDataSourceConfigExtAttrFilters(filter, mac, extAttrsValue string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_macfilteraddress" "test" {
  filter = nios_dhcp_filtermac.parent_filter_mac.name
  mac    = %q
  extattrs = {
    Site = %q
  } 
}

data "nios_dhcp_macfilteraddress" "test" {
  extattrfilters = {
	Site = nios_dhcp_macfilteraddress.test.extattrs.Site
  }
}
`, mac, extAttrsValue)
	return strings.Join([]string{testAccBaseWithMacFilter(filter), config}, "")
}
