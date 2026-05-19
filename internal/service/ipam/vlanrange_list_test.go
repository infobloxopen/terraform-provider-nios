package ipam_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/querycheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

func TestAccVlanrangeList_basic(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test"
	var v ipam.Vlanrange
	vlanRange := acctest.RandomNameWithPrefix("vlan-range")
	vlanView := acctest.RandomNameWithPrefix("vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version0_14_0),
		},
		Steps: []resource.TestStep{
			// Provider Setup
			{
				Config: utils.ProviderSetup(),
			},
			// Create and Read
			{
				Config: testAccVlanrangeBasicConfig(71, vlanRange, 61, vlanView),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				Query:  true,
				Config: testAccVlanrangeListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_ipam_vlanrange.test", 1),
				},
			},
		},
	})
}

func TestAccVlanrangeList_Filters(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test"
	var v ipam.Vlanrange
	vlanRange := acctest.RandomNameWithPrefix("vlan-range")
	vlanView := acctest.RandomNameWithPrefix("vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version0_14_0),
		},
		Steps: []resource.TestStep{
			// Provider Setup
			{
				Config: utils.ProviderSetup(),
			},
			// Create and Read
			{
				Config: testAccVlanrangeBasicConfig(71, vlanRange, 61, vlanView),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", vlanRange),
				),
			},
			// Query the object
			{
				Query:  true,
				Config: testAccVlanrangeListConfigFilters(vlanRange),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_ipam_vlanrange.test", 1),
				},
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanrangeList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test_extattrs"
	var v ipam.Vlanrange
	vlanRange := acctest.RandomNameWithPrefix("vlan-range")
	vlanView := acctest.RandomNameWithPrefix("vlan-view")

	extAttrValue := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version0_14_0),
		},
		Steps: []resource.TestStep{
			// Provider Setup
			{
				Config: utils.ProviderSetup(),
			},
			// Create and Read
			{
				Config: testAccVlanrangeExtAttrs(71, vlanRange, 61, vlanView, map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				Query:  true,
				Config: testAccVlanrangeListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_ipam_vlanrange.test", 1),
				},
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccVlanrangeListBasicConfig() string {
	return `
list "nios_ipam_vlanrange" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccVlanrangeListConfigFilters(name string) string {
	return fmt.Sprintf(`
list "nios_ipam_vlanrange" "test" {
	provider = nios
	config {
		filters = {
			name =  %q
		}
	}
}
`, name)
}

func testAccVlanrangeListConfigExtAttrFilters(name string) string {
	return fmt.Sprintf(`
list "nios_ipam_vlanrange" "test" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, name)
}
