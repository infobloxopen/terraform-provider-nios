package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

func TestAccNsgroupForwardstubserverDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_nsgroup_forwardstubserver.test"
	resourceName := "nios_dns_nsgroup_forwardstubserver.test"
	var v dns.NsgroupForwardstubserver
	name := acctest.RandomNameWithPrefix("ns-group-forwardstubserver")
	externalServers := []map[string]any{
		{
			"name":    "example.com",
			"address": "2.3.3.4",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupForwardstubserverDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupForwardstubserverDataSourceConfigFilters(name, externalServers),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					}, testAccCheckNsgroupForwardstubserverResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccNsgroupForwardstubserverDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_nsgroup_forwardstubserver.test"
	resourceName := "nios_dns_nsgroup_forwardstubserver.test"
	var v dns.NsgroupForwardstubserver
	name := acctest.RandomNameWithPrefix("ns-group-forwardstubserver")
	externalServers := []map[string]any{
		{
			"name":    "example.com",
			"address": "2.3.3.4",
		},
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupForwardstubserverDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupForwardstubserverDataSourceConfigExtAttrFilters(name, externalServers, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					}, testAccCheckNsgroupForwardstubserverResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckNsgroupForwardstubserverResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "external_servers", dataSourceName, "result.0.external_servers"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
	}
}

func testAccNsgroupForwardstubserverDataSourceConfigFilters(name string, externalServers []map[string]any) string {
	externalServersStr := utils.ConvertSliceOfMapsToHCL(externalServers)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardstubserver" "test" {
    name = %q
    external_servers = %s
}

data "nios_dns_nsgroup_forwardstubserver" "test" {
  filters = {
    name = nios_dns_nsgroup_forwardstubserver.test.name
  }
}
`, name, externalServersStr)
}

func testAccNsgroupForwardstubserverDataSourceConfigExtAttrFilters(name string, externalServers []map[string]any, extAttrsValue string) string {
	externalServersStr := utils.ConvertSliceOfMapsToHCL(externalServers)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardstubserver" "test" {
  name = %q
  external_servers = %s
  extattrs = {
    Site = %q
  } 
}

data "nios_dns_nsgroup_forwardstubserver" "test" {
  extattrfilters = {
	Site = nios_dns_nsgroup_forwardstubserver.test.extattrs.Site
  }
}
`, name, externalServersStr, extAttrsValue)
}
