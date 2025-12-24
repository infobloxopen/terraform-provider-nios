package dhcp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccFilterrelayagentDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_filterrelayagent.test"
	resourceName := "nios_dhcp_filterrelayagent.test"
	var v dhcp.Filterrelayagent
	name := acctest.RandomNameWithPrefix("filterrelayagent")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFilterrelayagentDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFilterrelayagentDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					}, testAccCheckFilterrelayagentResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccFilterrelayagentDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_filterrelayagent.test"
	resourceName := "nios_dhcp_filterrelayagent.test"
	var v dhcp.Filterrelayagent
	name := acctest.RandomNameWithPrefix("filterrelayagent")
	siteExtAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFilterrelayagentDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFilterrelayagentDataSourceConfigExtAttrFilters(name, siteExtAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					}, testAccCheckFilterrelayagentResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckFilterrelayagentResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "circuit_id_name", dataSourceName, "result.0.circuit_id_name"),
		resource.TestCheckResourceAttrPair(resourceName, "circuit_id_substring_length", dataSourceName, "result.0.circuit_id_substring_length"),
		resource.TestCheckResourceAttrPair(resourceName, "circuit_id_substring_offset", dataSourceName, "result.0.circuit_id_substring_offset"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "is_circuit_id", dataSourceName, "result.0.is_circuit_id"),
		resource.TestCheckResourceAttrPair(resourceName, "is_circuit_id_substring", dataSourceName, "result.0.is_circuit_id_substring"),
		resource.TestCheckResourceAttrPair(resourceName, "is_remote_id", dataSourceName, "result.0.is_remote_id"),
		resource.TestCheckResourceAttrPair(resourceName, "is_remote_id_substring", dataSourceName, "result.0.is_remote_id_substring"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "remote_id_name", dataSourceName, "result.0.remote_id_name"),
		resource.TestCheckResourceAttrPair(resourceName, "remote_id_substring_length", dataSourceName, "result.0.remote_id_substring_length"),
		resource.TestCheckResourceAttrPair(resourceName, "remote_id_substring_offset", dataSourceName, "result.0.remote_id_substring_offset"),
	}
}

func testAccFilterrelayagentDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filterrelayagent" "test" {
  name = %q
  is_circuit_id = "NOT_SET"
}

data "nios_dhcp_filterrelayagent" "test" {
  filters = {
	name = nios_dhcp_filterrelayagent.test.name
  }
}
`, name)
}

func testAccFilterrelayagentDataSourceConfigExtAttrFilters(name string, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filterrelayagent" "test" {
  name = %q
  is_circuit_id = "NOT_SET"
  extattrs = {
    Site = %q
  }
}

data "nios_dhcp_filterrelayagent" "test" {
  extattrfilters = {
	Site = nios_dhcp_filterrelayagent.test.extattrs.Site
  }
}
`, name, extAttrsValue)
}
