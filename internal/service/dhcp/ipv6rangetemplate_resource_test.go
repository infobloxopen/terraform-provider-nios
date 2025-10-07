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
	name := acctest.RandomNameWithPrefix("ipv6-range-template")
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
	name := acctest.RandomNameWithPrefix("ipv6-range-template")
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
	name := acctest.RandomNameWithPrefix("ipv6-range-template")
	numberOfAdresses := 100
	offset := 50
	delegatedMember1 := map[string]any{

		"name":     "infoblox.172_28_82_213",
		"ipv4addr": "172.28.82.213",
	}
	delegatedMember2 := map[string]any{
		"name":     "infoblox.172_28_82_185",
		"ipv4addr": "172.28.82.185",
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateDelegatedMember(name, numberOfAdresses, offset, delegatedMember1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegated_member.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegated_member.name", "infoblox.172_28_82_213"),
					resource.TestCheckResourceAttr(resourceName, "delegated_member.ipv4addr", "172.28.82.213"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateDelegatedMember(name, numberOfAdresses, offset, delegatedMember2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegated_member.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegated_member.name", "infoblox.172_28_82_185"),
					resource.TestCheckResourceAttr(resourceName, "delegated_member.ipv4addr", "172.28.82.185"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_Exclude(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_exclude"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("ipv6-range-template")
	numberOfAdresses := 100
	offset := 50
	exclude := []map[string]any{
		{
			"number_of_addresses": 10,
			"offset":              20,
		},
	}
	excludeUpdated := []map[string]any{
		{
			"number_of_addresses": 15,
			"offset":              25,
			"comment":             "exclude for range template",
		},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateExclude(name, numberOfAdresses, offset, exclude),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "exclude.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "exclude.0.number_of_addresses", "10"),
					resource.TestCheckResourceAttr(resourceName, "exclude.0.offset", "20"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateExclude(name, numberOfAdresses, offset, excludeUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "exclude.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "exclude.0.number_of_addresses", "15"),
					resource.TestCheckResourceAttr(resourceName, "exclude.0.offset", "25"),
					resource.TestCheckResourceAttr(resourceName, "exclude.0.comment", "exclude for range template"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_LogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_logic_filter_rules"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("ipv6-range-template")
	numberOfAdresses := 100
	offset := 50
	logicFilterRules := []map[string]any{
		{
			"filter": "option_filter",
			"type":   "Option",
		},
	}
	logicFilterRulesUpdated := []map[string]any{
		{
			"filter": "option_logic_filter",
			"type":   "Option",
		},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateLogicFilterRules(name, numberOfAdresses, offset, logicFilterRules),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.filter", "option_filter"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.type", "Option"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateLogicFilterRules(name, numberOfAdresses, offset, logicFilterRulesUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.filter", "option_logic_filter"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.type", "Option"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_Member(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_member"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("ipv6-range-template")
	numberOfAdresses := 100
	offset := 50
	member1 := map[string]any{
		"name":     "infoblox.172_28_82_213",
		"ipv4addr": "172.28.82.213",
	}
	member2 := map[string]any{
		"name":     "infoblox.172_28_82_185",
		"ipv4addr": "172.28.82.185",
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateMember(name, numberOfAdresses, offset, member1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "member.name", "infoblox.172_28_82_213"),
					resource.TestCheckResourceAttr(resourceName, "member.ipv4addr", "172.28.82.213"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateMember(name, numberOfAdresses, offset, member2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "member.name", "infoblox.172_28_82_185"),
					resource.TestCheckResourceAttr(resourceName, "member.ipv4addr", "172.28.82.185"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_name"
	var v dhcp.Ipv6rangetemplate
	name1 := acctest.RandomNameWithPrefix("ipv6-range-template")
	name2 := acctest.RandomNameWithPrefix("ipv6-range-template")
	numberOfAdresses := 100
	offset := 50
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateName(name1, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateName(name2, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_NumberOfAddresses(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_number_of_addresses"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("ipv6-range-template")
	offset := 50
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateNumberOfAddresses(name, 100, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "number_of_addresses", "100"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateNumberOfAddresses(name, 150, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "number_of_addresses", "150"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_Offset(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_offset"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("ipv6-range-template")
	numberOfAdresses := 100
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateOffset(name, numberOfAdresses, 200),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "offset", "200"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateOffset(name, numberOfAdresses, 250),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "offset", "250"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_OptionFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_option_filter_rules"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("ipv6-range-template")
	numberOfAdresses := 100
	offset := 50
	optionFilterRules := []map[string]any{
		{
			"filter":     "option_filter",
			"permission": "Allow",
		},
	}
	optionFilterRulesUpdated := []map[string]any{
		{
			"filter":     "option_filter",
			"permission": "Deny",
		},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateOptionFilterRules(name, numberOfAdresses, offset, optionFilterRules),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_filter_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "option_filter_rules.0.filter", "option_filter"),
					resource.TestCheckResourceAttr(resourceName, "option_filter_rules.0.permission", "Allow"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateOptionFilterRules(name, numberOfAdresses, offset, optionFilterRulesUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_filter_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "option_filter_rules.0.filter", "option_filter"),
					resource.TestCheckResourceAttr(resourceName, "option_filter_rules.0.permission", "Deny"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_RecycleLeases(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_recycle_leases"
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
				Config: testAccIpv6rangetemplateRecycleLeases(name, numberOfAdresses, offset, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateRecycleLeases(name, numberOfAdresses, offset, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_ServerAssociationType(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_server_association_type"
	var v dhcp.Ipv6rangetemplate
	name := acctest.RandomNameWithPrefix("ipv6-range-template")
	numberOfAdresses := 100
	offset := 50
	member := map[string]any{
		"name":     "infoblox.172_28_82_213",
		"ipv4addr": "172.28.82.213",
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangetemplateServerAssociationType(name, numberOfAdresses, offset, "MEMBER", member),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "server_association_type", "MEMBER"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateServerAssociationType(name, numberOfAdresses, offset, "NONE", nil),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "server_association_type", "NONE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_UseLogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_use_logic_filter_rules"
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
				Config: testAccIpv6rangetemplateUseLogicFilterRules(name, numberOfAdresses, offset, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateUseLogicFilterRules(name, numberOfAdresses, offset, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangetemplateResource_UseRecycleLeases(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6_range_template.test_use_recycle_leases"
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
				Config: testAccIpv6rangetemplateUseRecycleLeases(name, numberOfAdresses, offset, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangetemplateUseRecycleLeases(name, numberOfAdresses, offset, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "false"),
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
	name               = %q
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
}
`, name, numberOfAddresses, offset)
}

func testAccIpv6rangetemplateCloudApiCompatible(name string, numberOfAddresses, offset int, cloudApiCompatible bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_cloud_api_compatible" {
    name               = %q
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = %t
}
`, name, numberOfAddresses, offset, cloudApiCompatible)
}

func testAccIpv6rangetemplateComment(name string, numberOfAddresses, offset int, comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_comment" {
	name               = %q
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
	comment = %q
}
`, name, numberOfAddresses, offset, comment)
}

func testAccIpv6rangetemplateDelegatedMember(name string, numberOfAddresses, offset int, delegatedMember map[string]any) string {
	delegatedMemberStr := utils.ConvertMapToHCL(delegatedMember)
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_delegated_member" {
    name               = %q
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
    delegated_member = %s
}
`, name, numberOfAddresses, offset, delegatedMemberStr)
}

func testAccIpv6rangetemplateExclude(name string, numberOfAddresses, offset int, exclude []map[string]any) string {
	excludeStr := utils.ConvertSliceOfMapsToHCL(exclude)
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_exclude" {
    name               = %q
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
    exclude = %s
}
`, name, numberOfAddresses, offset, excludeStr)
}

func testAccIpv6rangetemplateLogicFilterRules(name string, numberOfAddresses, offset int, logicFilterRules []map[string]any) string {
	logicFilterRulesStr := utils.ConvertSliceOfMapsToHCL(logicFilterRules)
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_logic_filter_rules" {
    name               = %q
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
    logic_filter_rules = %s
    use_logic_filter_rules = true 
}
`, name, numberOfAddresses, offset, logicFilterRulesStr)
}

func testAccIpv6rangetemplateMember(name string, numberOfAddresses, offset int, member map[string]any) string {
	membersStr := utils.ConvertMapToHCL(member)
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_member" {
    name               = %q
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
    member = %s
}
`, name, numberOfAddresses, offset, membersStr)
}

func testAccIpv6rangetemplateName(name string, numberOfAddresses, offset int) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_name" {
    name               = %q
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
}
`, name, numberOfAddresses, offset)
}

func testAccIpv6rangetemplateNumberOfAddresses(name string, numberOfAddresses, offset int) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_number_of_addresses" {
    name               = %q
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
}
`, name, numberOfAddresses, offset)
}

func testAccIpv6rangetemplateOffset(name string, numberOfAddresses, offset int) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_offset" {
    name               = %q
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
}
`, name, numberOfAddresses, offset)
}

func testAccIpv6rangetemplateOptionFilterRules(name string, numberOfAddresses, offset int, optionFilterRules []map[string]any) string {
	optionFilterRulesStr := utils.ConvertSliceOfMapsToHCL(optionFilterRules)
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_option_filter_rules" {
    name               = %q
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
    option_filter_rules = %s
}
`, name, numberOfAddresses, offset, optionFilterRulesStr)
}

func testAccIpv6rangetemplateRecycleLeases(name string, numberOfAddresses, offset int, recycleLeases bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_recycle_leases" {
    name               = %q
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
    recycle_leases = %t
    use_recycle_leases = true
}
`, name, numberOfAddresses, offset, recycleLeases)
}

func testAccIpv6rangetemplateServerAssociationType(name string, numberOfAddresses, offset int, serverAssociationType string, member map[string]any) string {
	var extraConfig string
	if member != nil {
		extraConfig = fmt.Sprintf(`member = %q`, utils.ConvertMapToHCL(member))
	}
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_server_association_type" {
    name               = %q
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
    server_association_type = %q
    %s
}
`, name, numberOfAddresses, offset, serverAssociationType, extraConfig)
}

func testAccIpv6rangetemplateUseLogicFilterRules(name string, numberOfAddresses, offset int, useLogicFilterRules bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "test_use_logic_filter_rules" {
    name               = %q
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
    name               = %q
	number_of_addresses = %d
	offset = %d
    cloud_api_compatible = true
    use_recycle_leases = %t
}
`, name, numberOfAddresses, offset, useRecycleLeases)
}
