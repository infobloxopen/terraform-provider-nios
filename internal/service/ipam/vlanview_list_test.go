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
)

func TestAccVlanviewList_basic(t *testing.T) {
	var resourceName = "nios_ipam_vlanview.test"
	var v ipam.Vlanview
	name := acctest.RandomNameWithPrefix("vlan_view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccVlanviewBasicConfig(15, name, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccVlanviewListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_ipam_vlanview.test", 1),
				},
			},
		},
	})
}

func TestAccVlanviewList_Filters(t *testing.T) {
	var resourceName = "nios_ipam_vlanview.test"
	var v ipam.Vlanview
	name := acctest.RandomNameWithPrefix("vlan_view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config:                   testAccVlanviewBasicConfig(15, name, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccVlanviewListConfigFilters(name),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_ipam_vlanview.test", 1),
				},
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanviewList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_ipam_vlanview.test_extattrs"
	var v ipam.Vlanview
	name := acctest.RandomNameWithPrefix("vlan_view")

	extAttrValue := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		Steps: []resource.TestStep{
			// Create and Read
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Config: testAccVlanviewExtAttrs(15, name, 10, map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Query:                    true,
				Config:                   testAccVlanviewListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_ipam_vlanview.test", 1),
				},
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccVlanviewListBasicConfig() string {
	return `
list "nios_ipam_vlanview" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccVlanviewListConfigFilters(name string) string {
	return fmt.Sprintf(`
list "nios_ipam_vlanview" "test" {
	provider = nios
	config {
		filters = {
			name =  %q
		}
	}
}
`, name)
}

func testAccVlanviewListConfigExtAttrFilters(name string) string {
	return fmt.Sprintf(`
list "nios_ipam_vlanview" "test" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, name)
}
