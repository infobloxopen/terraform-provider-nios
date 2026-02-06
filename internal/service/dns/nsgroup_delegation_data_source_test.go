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

func TestAccNsgroupDelegationDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_nsgroup_delegation.test"
	resourceName := "nios_dns_nsgroup_delegation.test"
	var v dns.NsgroupDelegation
	name := acctest.RandomName()
	delegateTo := []map[string]interface{}{
		{
			"name":    "delegate_to_ns_group",
			"address": "2.3.4.5",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupDelegationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupDelegationDataSourceConfigFilters(name, delegateTo),
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
	name := acctest.RandomName()
	delegateTo := []map[string]interface{}{
		{
			"name":    "delegate_to_ns_group",
			"address": "2.3.4.5",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupDelegationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupDelegationDataSourceConfigExtAttrFilters(name, delegateTo, acctest.RandomName()),
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

func testAccCheckNsgroupDelegationResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "delegate_to", dataSourceName, "result.0.delegate_to"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
	}
}

func testAccNsgroupDelegationDataSourceConfigFilters(name string, delegateTo []map[string]any) string {
	delegateToStr := utils.ConvertSliceOfMapsToHCL(delegateTo)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_delegation" "test" {
    name = %q
    delegate_to = %s
}

data "nios_dns_nsgroup_delegation" "test" {
  filters = {
	name = nios_dns_nsgroup_delegation.test.name
  }
}
`, name, delegateToStr)
}

func testAccNsgroupDelegationDataSourceConfigExtAttrFilters(name string, delegateTo []map[string]interface{}, extAttrsValue string) string {
	delegateToStr := utils.ConvertSliceOfMapsToHCL(delegateTo)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_delegation" "test" {
  name = %q
  delegate_to = %s
  extattrs = {
    Site = %q
  } 
}

data "nios_dns_nsgroup_delegation" "test" {
  extattrfilters = {
	Site = nios_dns_nsgroup_delegation.test.extattrs.Site
  }
}
`, name, delegateToStr, extAttrsValue)
}
