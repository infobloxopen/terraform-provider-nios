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
	var resourceName = "nios_dhcp_ipv6rangetemplate.test"
	var v dhcp.Ipv6rangetemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateBasicConfig(),
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
	resourceName := "nios_dhcp_ipv6rangetemplate.test"
	var v dhcp.Ipv6rangetemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6rangetemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6rangetemplateBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					testAccCheckIpv6rangetemplateDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccIpv6rangetemplateResource_Ref(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6rangetemplate.test_ref"
	var v dhcp.Ipv6rangetemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_CloudApiCompatible(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6rangetemplate.test_cloud_api_compatible"
	var v dhcp.Ipv6rangetemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateCloudApiCompatible("CLOUD_API_COMPATIBLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_api_compatible", "CLOUD_API_COMPATIBLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateCloudApiCompatible("CLOUD_API_COMPATIBLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_api_compatible", "CLOUD_API_COMPATIBLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6rangetemplate.test_comment"
	var v dhcp.Ipv6rangetemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_DelegatedMember(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6rangetemplate.test_delegated_member"
	var v dhcp.Ipv6rangetemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateDelegatedMember("DELEGATED_MEMBER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegated_member", "DELEGATED_MEMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateDelegatedMember("DELEGATED_MEMBER_UPDATE_REPLACE_ME"),
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
	var resourceName = "nios_dhcp_ipv6rangetemplate.test_exclude"
	var v dhcp.Ipv6rangetemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateExclude("EXCLUDE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "exclude", "EXCLUDE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateExclude("EXCLUDE_UPDATE_REPLACE_ME"),
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
	var resourceName = "nios_dhcp_ipv6rangetemplate.test_logic_filter_rules"
	var v dhcp.Ipv6rangetemplate

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
	var resourceName = "nios_dhcp_ipv6rangetemplate.test_member"
	var v dhcp.Ipv6rangetemplate

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
	var resourceName = "nios_dhcp_ipv6rangetemplate.test_name"
	var v dhcp.Ipv6rangetemplate

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
	var resourceName = "nios_dhcp_ipv6rangetemplate.test_number_of_addresses"
	var v dhcp.Ipv6rangetemplate

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
	var resourceName = "nios_dhcp_ipv6rangetemplate.test_offset"
	var v dhcp.Ipv6rangetemplate

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
	var resourceName = "nios_dhcp_ipv6rangetemplate.test_option_filter_rules"
	var v dhcp.Ipv6rangetemplate

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
	var resourceName = "nios_dhcp_ipv6rangetemplate.test_recycle_leases"
	var v dhcp.Ipv6rangetemplate

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
	var resourceName = "nios_dhcp_ipv6rangetemplate.test_server_association_type"
	var v dhcp.Ipv6rangetemplate

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
	var resourceName = "nios_dhcp_ipv6rangetemplate.test_use_logic_filter_rules"
	var v dhcp.Ipv6rangetemplate

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
	var resourceName = "nios_dhcp_ipv6rangetemplate.test_use_recycle_leases"
	var v dhcp.Ipv6rangetemplate

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

func testAccIpv6rangetemplateBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6rangetemplate" "test" {
}
`)
}

func testAccIpv6rangetemplateRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6rangetemplate" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccIpv6rangetemplateCloudApiCompatible(cloudApiCompatible string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6rangetemplate" "test_cloud_api_compatible" {
    cloud_api_compatible = %q
}
`, cloudApiCompatible)
}

func testAccIpv6rangetemplateComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6rangetemplate" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccIpv6rangetemplateDelegatedMember(delegatedMember string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6rangetemplate" "test_delegated_member" {
    delegated_member = %q
}
`, delegatedMember)
}

func testAccIpv6rangetemplateExclude(exclude string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6rangetemplate" "test_exclude" {
    exclude = %q
}
`, exclude)
}

func testAccIpv6rangetemplateLogicFilterRules(logicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6rangetemplate" "test_logic_filter_rules" {
    logic_filter_rules = %q
}
`, logicFilterRules)
}

func testAccIpv6rangetemplateMember(member string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6rangetemplate" "test_member" {
    member = %q
}
`, member)
}

func testAccIpv6rangetemplateName(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6rangetemplate" "test_name" {
    name = %q
}
`, name)
}

func testAccIpv6rangetemplateNumberOfAddresses(numberOfAddresses string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6rangetemplate" "test_number_of_addresses" {
    number_of_addresses = %q
}
`, numberOfAddresses)
}

func testAccIpv6rangetemplateOffset(offset string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6rangetemplate" "test_offset" {
    offset = %q
}
`, offset)
}

func testAccIpv6rangetemplateOptionFilterRules(optionFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6rangetemplate" "test_option_filter_rules" {
    option_filter_rules = %q
}
`, optionFilterRules)
}

func testAccIpv6rangetemplateRecycleLeases(recycleLeases string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6rangetemplate" "test_recycle_leases" {
    recycle_leases = %q
}
`, recycleLeases)
}

func testAccIpv6rangetemplateServerAssociationType(serverAssociationType string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6rangetemplate" "test_server_association_type" {
    server_association_type = %q
}
`, serverAssociationType)
}

func testAccIpv6rangetemplateUseLogicFilterRules(useLogicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6rangetemplate" "test_use_logic_filter_rules" {
    use_logic_filter_rules = %q
}
`, useLogicFilterRules)
}

func testAccIpv6rangetemplateUseRecycleLeases(useRecycleLeases string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6rangetemplate" "test_use_recycle_leases" {
    use_recycle_leases = %q
}
`, useRecycleLeases)
}
