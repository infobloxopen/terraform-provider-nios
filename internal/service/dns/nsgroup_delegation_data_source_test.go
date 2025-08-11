
package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccNsgroupDelegationDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_nsgroup_delegation.test"
	resourceName := "nios_dns_nsgroup_delegation.test"
	var v dns.NsgroupDelegation

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupDelegationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupDelegationDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
						}, testAccCheckNsgroupDelegationResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccNsgroupDelegationDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_nsgroup_delegation.test"
	resourceName := "nios_dns_nsgroup_delegation.test"
	var v dns.NsgroupDelegation
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupDelegationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupDelegationDataSourceConfigExtAttrFilters( acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckNsgroupDelegationExists(context.Background(), resourceName, &v),
						}, testAccCheckNsgroupDelegationResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckNsgroupDelegationResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
    return []resource.TestCheckFunc{
        resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
        resource.TestCheckResourceAttrPair(resourceName, "delegate_to", dataSourceName, "result.0.delegate_to"),
        resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
        resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
    }
}

func testAccNsgroupDelegationDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_delegation" "test" {
}

data "nios_dns_nsgroup_delegation" "test" {
  filters = {
	 = nios_dns_nsgroup_delegation.test.
  }
}
`)
}

func testAccNsgroupDelegationDataSourceConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_delegation" "test" {
  extattrs = {
    Site = %q
  } 
}

data "nios_dns_nsgroup_delegation" "test" {
  extattrfilters = {
	Site = nios_dns_nsgroup_delegation.test.extattrs.Site
  }
}
`,extAttrsValue)
}

