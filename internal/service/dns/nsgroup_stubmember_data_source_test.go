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

func TestAccNsgroupStubmemberDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_nsgroup_stubmember.test"
	resourceName := "nios_dns_nsgroup_stubmember.test"
	var v dns.NsgroupStubmember
	name := acctest.RandomNameWithPrefix("test-nsgroup-stubmember")
	stubMember := []map[string]any{
		{
			"name": "member.com",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupStubmemberDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupStubmemberDataSourceConfigFilters(name, stubMember),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckNsgroupStubmemberExists(context.Background(), resourceName, &v),
					}, testAccCheckNsgroupStubmemberResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccNsgroupStubmemberDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_nsgroup_stubmember.test"
	resourceName := "nios_dns_nsgroup_stubmember.test"
	var v dns.NsgroupStubmember
	name := acctest.RandomNameWithPrefix("test-nsgroup-stubmember")
	stubMember := []map[string]any{
		{
			"name": "member.com",
		},
	}
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupStubmemberDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupStubmemberDataSourceConfigExtAttrFilters(name, stubMember, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckNsgroupStubmemberExists(context.Background(), resourceName, &v),
					}, testAccCheckNsgroupStubmemberResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckNsgroupStubmemberResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "stub_members", dataSourceName, "result.0.stub_members"),
	}
}

func testAccNsgroupStubmemberDataSourceConfigFilters(name string, stubMember []map[string]any) string {
	stubMemberStr := utils.ConvertSliceOfMapsToHCL(stubMember)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_stubmember" "test" {
    name = %q
    stub_members = %s
}

data "nios_dns_nsgroup_stubmember" "test" {
  filters = {
    name = nios_dns_nsgroup_stubmember.test.name
  }
}
`, name, stubMemberStr)
}

func testAccNsgroupStubmemberDataSourceConfigExtAttrFilters(name string, stubMember []map[string]any, extAttrsValue string) string {
	stubMemberStr := utils.ConvertSliceOfMapsToHCL(stubMember)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_stubmember" "test" {
	name = %q
	stub_members = %s
  	extattrs = {
    	Site = %q
  	} 
}

data "nios_dns_nsgroup_stubmember" "test" {
  extattrfilters = {
	Site = nios_dns_nsgroup_stubmember.test.extattrs.Site
  }
}
`, name, stubMemberStr, extAttrsValue)
}
