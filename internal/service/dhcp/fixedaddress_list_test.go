package dhcp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/querycheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

func TestAccFixedaddressList_basic(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test"
	var v dhcp.Fixedaddress
	ip := "15.0.0.111"
	agentCircuitID := acctest.RandomNumber(1000)

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
				Config: testAccFixedaddressBasicConfig(ip, "CIRCUIT_ID", agentCircuitID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
				),
			},
			// Query the object
			{
				Query:  true,
				Config: testAccFixedaddressListBasicConfig(),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLengthAtLeast("nios_dhcp_fixed_address.test", 1),
				},
			},
		},
	})
}

func TestAccFixedaddressList_Filters(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test"
	var v dhcp.Fixedaddress
	ip := "15.0.0.112"
	agentCircuitID := acctest.RandomNumber(1000)

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
				Config: testAccFixedaddressBasicConfig(ip, "CIRCUIT_ID", agentCircuitID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", ip),
				),
			},
			// Query the object
			{
				Query:  true,
				Config: testAccFixedaddressListConfigFilters(ip),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dhcp_fixed_address.test", 1),
				},
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFixedaddressList_ExtAttrFilters(t *testing.T) {
	var resourceName = "nios_dhcp_fixed_address.test_extattrs"
	var v dhcp.Fixedaddress
	ip := "15.0.0.113"
	agentCircuitID := acctest.RandomNumber(1000)

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
				Config: testAccFixedaddressExtAttrs(ip, "CIRCUIT_ID", agentCircuitID, map[string]string{
					"Site": extAttrValue,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFixedaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Query the object
			{
				Query:  true,
				Config: testAccFixedaddressListConfigExtAttrFilters(extAttrValue),
				QueryResultChecks: []querycheck.QueryResultCheck{
					querycheck.ExpectLength("nios_dhcp_fixed_address.test", 1),
				},
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccFixedaddressListBasicConfig() string {
	return `
list "nios_dhcp_fixed_address" "test" {
	provider = nios
	limit = 5
}
`
}

func testAccFixedaddressListConfigFilters(ip4addr string) string {
	return fmt.Sprintf(`
list "nios_dhcp_fixed_address" "test" {
	provider = nios
	config {
		filters = {
			ipv4addr =  %q
		}
	}
}
`, ip4addr)
}

func testAccFixedaddressListConfigExtAttrFilters(extAttrVal string) string {
	return fmt.Sprintf(`
list "nios_dhcp_fixed_address" "test" {
	provider = nios
	config {
		extattrfilters = {
			Site =  %q
		}
	}
}
`, extAttrVal)
}
