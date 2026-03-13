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

func TestAccNsgroupForwardingmemberDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_nsgroup_forwardingmember.test"
	resourceName := "nios_dns_nsgroup_forwardingmember.test"
	var v dns.NsgroupForwardingmember
	name := acctest.RandomNameWithPrefix("ns-group-forwardingMember")
	forwardingServers := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupForwardingmemberDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupForwardingmemberDataSourceConfigFilters(name, forwardingServers),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					}, testAccCheckNsgroupForwardingmemberResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccNsgroupForwardingmemberDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_nsgroup_forwardingmember.test"
	resourceName := "nios_dns_nsgroup_forwardingmember.test"
	var v dns.NsgroupForwardingmember
	name := acctest.RandomNameWithPrefix("ns-group-forwardingMember")
	forwardingServers := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},
	}
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupForwardingmemberDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupForwardingmemberDataSourceConfigExtAttrFilters(name, forwardingServers, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					}, testAccCheckNsgroupForwardingmemberResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckNsgroupForwardingmemberResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "forwarding_servers", dataSourceName, "result.0.forwarding_servers"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
	}
}

func testAccNsgroupForwardingmemberDataSourceConfigFilters(name string, forwardingServers []map[string]any) string {
	forwardingServersStr := utils.ConvertSliceOfMapsToHCL(forwardingServers)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardingmember" "test" {
  name = %q
  forwarding_servers = %s
}

data "nios_dns_nsgroup_forwardingmember" "test" {
  filters = {
    name = nios_dns_nsgroup_forwardingmember.test.name
  }
}
`, name, forwardingServersStr)
}

func testAccNsgroupForwardingmemberDataSourceConfigExtAttrFilters(name string, forwardingServers []map[string]any, extAttrsValue string) string {
	forwardingServersStr := utils.ConvertSliceOfMapsToHCL(forwardingServers)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardingmember" "test" {
  name = %q
  forwarding_servers = %s
  extattrs = {
    Site = %q
  } 
}

data "nios_dns_nsgroup_forwardingmember" "test" {
  extattrfilters = {
	Site = nios_dns_nsgroup_forwardingmember.test.extattrs.Site
  }
}
`, name, forwardingServersStr, extAttrsValue)
}
