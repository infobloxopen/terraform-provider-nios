
package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccNsgroupForwardingmemberDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_nsgroup_forwardingmember.test"
	resourceName := "nios_dns_nsgroup_forwardingmember.test"
	var v dns.NsgroupForwardingmember

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupForwardingmemberDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupForwardingmemberDataSourceConfigFilters(),
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
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupForwardingmemberDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupForwardingmemberDataSourceConfigExtAttrFilters( acctest.RandomName()),
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

func testAccCheckNsgroupForwardingmemberResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
    return []resource.TestCheckFunc{
        resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
        resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
        resource.TestCheckResourceAttrPair(resourceName, "forwarding_servers", dataSourceName, "result.0.forwarding_servers"),
        resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
    }
}

func testAccNsgroupForwardingmemberDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardingmember" "test" {
}

data "nios_dns_nsgroup_forwardingmember" "test" {
  filters = {
	 = nios_dns_nsgroup_forwardingmember.test.
  }
}
`)
}

func testAccNsgroupForwardingmemberDataSourceConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardingmember" "test" {
  extattrs = {
    Site = %q
  } 
}

data "nios_dns_nsgroup_forwardingmember" "test" {
  extattrfilters = {
	Site = nios_dns_nsgroup_forwardingmember.test.extattrs.Site
  }
}
`,extAttrsValue)
}

