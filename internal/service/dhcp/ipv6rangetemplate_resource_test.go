package dhcp_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForIpv6rangetemplate = "cloud_api_compatible,comment,delegated_member,exclude,logic_filter_rules,member,name,number_of_addresses,offset,option_filter_rules,recycle_leases,server_association_type,use_logic_filter_rules,use_recycle_leases"

func TestAccIpv6rangetemplateResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("ipv6-range-template")
	numberOfAdresses := 100
	offset := 50

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateBasicConfig(name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_ipv6_range_template.test"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("ipv6-range-template")
	numberOfAdresses := 100
	offset := 50
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6rangetemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6rangetemplateBasicConfig(name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					testAccCheckIpv6rangetemplateDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// The testcase will fail, as this is a known issue
// If the user is a cloud-user, then they need Terraform internal ID with cloud permission and enable cloud delegation for the user to create a range template.
// if the user is a non cloud-user, they need to have  Terraform internal ID without cloud permission.
func TestAccIpv6rangetemplateResource_CloudApiCompatible(t *testing.T) {
	t.Skip("Skipping this test as it is a known issue.")
	var resourceName = "nios_dhcp_ipv6_range_template.test_cloud_api_compatible"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := 100
	offset := 50
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateCloudApiCompatible(name, numberOfAdresses, offset, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_api_compatible", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateCloudApiCompatible(name, numberOfAdresses, offset, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_api_compatible", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_comment"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := 100
	offset := 50
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateComment(name, numberOfAdresses, offset, "example comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "example comment"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateComment(name, numberOfAdresses, offset, "example comment updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "example comment updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_DelegatedMember(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_delegated_member"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := 100
	offset := 50
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateDelegatedMember(name, numberOfAdresses, offset, "DELEGATED_MEMBER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegated_member", "DELEGATED_MEMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateDelegatedMember(name, numberOfAdresses, offset, "DELEGATED_MEMBER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegated_member", "DELEGATED_MEMBER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_Exclude(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_exclude"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := 100
	offset := 50
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateExclude(name, numberOfAdresses, offset, "EXCLUDE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "exclude", "EXCLUDE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateExclude(name, numberOfAdresses, offset, "EXCLUDE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "exclude", "EXCLUDE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_LogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_logic_filter_rules"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := 100
	offset := 50
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateLogicFilterRules("LOGIC_FILTER_RULES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateLogicFilterRules("LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_Member(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_member"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := 100
	offset := 50
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateMember("MEMBER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member", "MEMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateMember("MEMBER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member", "MEMBER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_name"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := 100
	offset := 50
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_NumberOfAddresses(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_number_of_addresses"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := 100
	offset := 50
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateNumberOfAddresses("NUMBER_OF_ADDRESSES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "number_of_addresses", "NUMBER_OF_ADDRESSES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateNumberOfAddresses("NUMBER_OF_ADDRESSES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "number_of_addresses", "NUMBER_OF_ADDRESSES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_Offset(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_offset"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := 100
	offset := 50
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateOffset("OFFSET_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "offset", "OFFSET_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateOffset("OFFSET_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "offset", "OFFSET_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_OptionFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_option_filter_rules"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := 100
	offset := 50
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateOptionFilterRules("OPTION_FILTER_RULES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_filter_rules", "OPTION_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateOptionFilterRules("OPTION_FILTER_RULES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_filter_rules", "OPTION_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_RecycleLeases(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_recycle_leases"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := 100
	offset := 50
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateRecycleLeases("RECYCLE_LEASES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "RECYCLE_LEASES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateRecycleLeases("RECYCLE_LEASES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "RECYCLE_LEASES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_ServerAssociationType(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_server_association_type"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := 100
	offset := 50
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateServerAssociationType("SERVER_ASSOCIATION_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "server_association_type", "SERVER_ASSOCIATION_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateServerAssociationType("SERVER_ASSOCIATION_TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "server_association_type", "SERVER_ASSOCIATION_TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_UseLogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_use_logic_filter_rules"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := 100
	offset := 50
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateUseLogicFilterRules("USE_LOGIC_FILTER_RULES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "USE_LOGIC_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateUseLogicFilterRules("USE_LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "USE_LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_UseRecycleLeases(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_use_recycle_leases"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := 100
	offset := 50
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateUseRecycleLeases("USE_RECYCLE_LEASES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "USE_RECYCLE_LEASES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateUseRecycleLeases("USE_RECYCLE_LEASES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "USE_RECYCLE_LEASES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckIpv6rangetemplateExists(ctx context.Context, resourceName string, v *dhcp.Ipv6rangetemplate) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			Ipv6rangetemplateAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForIpv6rangetemplate).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetIpv6rangetemplateResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetIpv6rangetemplateResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckIpv6rangetemplateDestroy(ctx context.Context, v *dhcp.Ipv6rangetemplate) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			Ipv6rangetemplateAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForIpv6rangetemplate).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				// resource was deleted
				return nil
			}
			return err
		}
		return errors.New("expected to be deleted")
	}
}

func testAccCheckIpv6rangetemplateDisappears(ctx context.Context, v *dhcp.Ipv6rangetemplate) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			Ipv6rangetemplateAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccIpv6rangetemplateBasicConfig(name string, numberOfAddresses, offset int) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test" {
	name               = %s
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
}
`, name, numberOfAddresses, offset)
}

func testAccIpv6rangetemplateCloudApiCompatible(name string, numberOfAddresses, offset int, cloudApiCompatible string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_cloud_api_compatible" {
    name               = %s
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = %q
}
`, name, numberOfAddresses, offset, cloudApiCompatible)
}

func testAccIpv6rangetemplateComment(name string, numberOfAddresses, offset int, comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_comment" {
	name               = %s
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
	comment = %q
}
`, name, numberOfAddresses, offset, comment)
}

func testAccIpv6rangetemplateDelegatedMember(name string, numberOfAddresses, offset int, delegatedMember string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_delegated_member" {
    name               = %s
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
    delegated_member = %q
}
`, name, numberOfAddresses, offset, delegatedMember)
}

func testAccIpv6rangetemplateExclude(name string, numberOfAddresses, offset int, exclude string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_exclude" {
    name               = %s
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
    exclude = %q
}
`, name, numberOfAddresses, offset, exclude)
}

func testAccIpv6rangetemplateLogicFilterRules(name string, numberOfAddresses, offset int, logicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_logic_filter_rules" {
    name               = %s
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
    logic_filter_rules = %q
}
`, name, numberOfAddresses, offset, logicFilterRules)
}

func testAccIpv6rangetemplateMember(name string, numberOfAddresses, offset int, member string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_member" {
    name               = %s
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
    member = %q
}
`, name, numberOfAddresses, offset, member)
}

func testAccIpv6rangetemplateName(name string, numberOfAddresses, offset int) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_name" {
    name               = %s
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
}
`, name, numberOfAddresses, offset)
}

func testAccIpv6rangetemplateNumberOfAddresses(name string, numberOfAddresses, offset int) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_number_of_addresses" {
    name               = %s
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
}
`, name, numberOfAddresses, offset)
}

func testAccIpv6rangetemplateOffset(name string, numberOfAddresses, offset int) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_offset" {
    name               = %s
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
}
`, name, numberOfAddresses, offset)
}

func testAccIpv6rangetemplateOptionFilterRules(name string, numberOfAddresses, offset int, optionFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_option_filter_rules" {
    name               = %s
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
    option_filter_rules = %q
}
`, name, numberOfAddresses, offset, optionFilterRules)
}

func testAccIpv6rangetemplateRecycleLeases(name string, numberOfAddresses, offset int, recycleLeases bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_recycle_leases" {
    name               = %s
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
    recycle_leases = %t
}
`, name, numberOfAddresses, offset, recycleLeases)
}

func testAccIpv6rangetemplateServerAssociationType(name string, numberOfAddresses, offset int, serverAssociationType string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_server_association_type" {
    name               = %s
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
    server_association_type = %q
}
`, name, numberOfAddresses, offset, serverAssociationType)
}

func testAccIpv6rangetemplateUseLogicFilterRules(name string, numberOfAddresses, offset int, useLogicFilterRules bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_use_logic_filter_rules" {
    name               = %s
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
    use_logic_filter_rules = %t
}
`, name, numberOfAddresses, offset, useLogicFilterRules)
}

func testAccIpv6rangetemplateUseRecycleLeases(name string, numberOfAddresses, offset int, useRecycleLeases bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_use_recycle_leases" {
    name               = %s
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
    use_recycle_leases = %t
}
`, name, numberOfAddresses, offset, useRecycleLeases)
}
