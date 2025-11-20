package dhcp_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForIpv6range = "address_type,cloud_info,comment,disable,discover_now_status,discovery_basic_poll_settings,discovery_blackout_setting,discovery_member,enable_discovery,end_addr,endpoint_sources,exclude,extattrs,ipv6_end_prefix,ipv6_prefix_bits,ipv6_start_prefix,logic_filter_rules,member,name,network,network_view,option_filter_rules,port_control_blackout_setting,recycle_leases,same_port_control_discovery_blackout,server_association_type,start_addr,subscribe_settings,use_blackout_setting,use_discovery_basic_polling_settings,use_enable_discovery,use_logic_filter_rules,use_recycle_leases,use_subscribe_settings"

func TestAccIpv6rangeResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test"
	var v dhcp.Ipv6range
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeBasicConfig(view, "14::1", "14::10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", "14::/64"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "address_type", "ADDRESS"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_discovery", "false"),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "true"),
					resource.TestCheckResourceAttr(resourceName, "restart_if_needed", "false"),
					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "false"),
					resource.TestCheckResourceAttr(resourceName, "server_association_type", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_subscribe_settings", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_ipv6range.test"
	var v dhcp.Ipv6range
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6rangeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6rangeBasicConfig(view, "14::20", "14::30"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					testAccCheckIpv6rangeDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccIpv6rangeResource_Import(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test"
	var v dhcp.Ipv6range
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeBasicConfig(view, "14::40", "14::50"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
				),
			},
			// Import with PlanOnly to detect differences
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccIpv6rangeImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "ref",
				PlanOnly:                             true,
				//ExpectError:                          regexp.MustCompile(`ImportStateVerify attributes not equivalent`),
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccIpv6rangeImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_AddressType(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_address_type"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeAddressType("NETWORK_REPLACE_ME", "ADDRESS_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address_type", "ADDRESS_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeAddressType("NETWORK_REPLACE_ME", "ADDRESS_TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address_type", "ADDRESS_TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_CloudInfo(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_cloud_info"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeCloudInfo("NETWORK_REPLACE_ME", "CLOUD_INFO_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_info", "CLOUD_INFO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeCloudInfo("NETWORK_REPLACE_ME", "CLOUD_INFO_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_info", "CLOUD_INFO_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccIpv6rangeResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_comment"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeComment("NETWORK_REPLACE_ME", "Comment for the Ipv6range object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeComment("NETWORK_REPLACE_ME", "Updated comment for the Ipv6range object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the object"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_Disable(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_disable"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeDisable("NETWORK_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeDisable("NETWORK_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_DiscoveryBasicPollSettings(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_discovery_basic_poll_settings"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeDiscoveryBasicPollSettings("NETWORK_REPLACE_ME", "DISCOVERY_BASIC_POLL_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings", "DISCOVERY_BASIC_POLL_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeDiscoveryBasicPollSettings("NETWORK_REPLACE_ME", "DISCOVERY_BASIC_POLL_SETTINGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings", "DISCOVERY_BASIC_POLL_SETTINGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccIpv6rangeResource_DiscoveryBlackoutSetting(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_discovery_blackout_setting"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeDiscoveryBlackoutSetting("NETWORK_REPLACE_ME", "DISCOVERY_BLACKOUT_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_blackout_setting", "DISCOVERY_BLACKOUT_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeDiscoveryBlackoutSetting("NETWORK_REPLACE_ME", "DISCOVERY_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_blackout_setting", "DISCOVERY_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccIpv6rangeResource_DiscoveryMember(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_discovery_member"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeDiscoveryMember("NETWORK_REPLACE_ME", "DISCOVERY_MEMBER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_member", "DISCOVERY_MEMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeDiscoveryMember("NETWORK_REPLACE_ME", "DISCOVERY_MEMBER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_member", "DISCOVERY_MEMBER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_EnableDiscovery(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_enable_discovery"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeEnableDiscovery("NETWORK_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_discovery", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeEnableDiscovery("NETWORK_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_discovery", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_EnableImmediateDiscovery(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_enable_immediate_discovery"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeEnableImmediateDiscovery("NETWORK_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_immediate_discovery", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeEnableImmediateDiscovery("NETWORK_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_immediate_discovery", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_EndAddr(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_end_addr"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeEndAddr("NETWORK_REPLACE_ME", "END_ADDR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "end_addr", "END_ADDR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeEndAddr("NETWORK_REPLACE_ME", "END_ADDR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "end_addr", "END_ADDR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_Exclude(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_exclude"
	var v dhcp.Ipv6range
	excludeVal := []map[string]any{}
	excludeValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeExclude("NETWORK_REPLACE_ME", excludeVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "exclude", "EXCLUDE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeExclude("NETWORK_REPLACE_ME", excludeValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "exclude", "EXCLUDE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_extattrs"
	var v dhcp.Ipv6range
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeExtAttrs("NETWORK_REPLACE_ME", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeExtAttrs("NETWORK_REPLACE_ME", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_Ipv6EndPrefix(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_ipv6_end_prefix"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeIpv6EndPrefix("NETWORK_REPLACE_ME", "IPV6_END_PREFIX_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_end_prefix", "IPV6_END_PREFIX_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeIpv6EndPrefix("NETWORK_REPLACE_ME", "IPV6_END_PREFIX_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_end_prefix", "IPV6_END_PREFIX_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_Ipv6PrefixBits(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_ipv6_prefix_bits"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeIpv6PrefixBits("NETWORK_REPLACE_ME", "IPV6_PREFIX_BITS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_prefix_bits", "IPV6_PREFIX_BITS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeIpv6PrefixBits("NETWORK_REPLACE_ME", "IPV6_PREFIX_BITS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_prefix_bits", "IPV6_PREFIX_BITS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccIpv6rangeResource_Ipv6StartPrefix(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_ipv6_start_prefix"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeIpv6StartPrefix("NETWORK_REPLACE_ME", "IPV6_START_PREFIX_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_start_prefix", "IPV6_START_PREFIX_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeIpv6StartPrefix("NETWORK_REPLACE_ME", "IPV6_START_PREFIX_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_start_prefix", "IPV6_START_PREFIX_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_LogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_logic_filter_rules"
	var v dhcp.Ipv6range
	logic_filter_rulesVal := []map[string]any{}
	logic_filter_rulesValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeLogicFilterRules("NETWORK_REPLACE_ME", logic_filter_rulesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeLogicFilterRules("NETWORK_REPLACE_ME", logic_filter_rulesValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_Member(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_member"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeMember("NETWORK_REPLACE_ME", "MEMBER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member", "MEMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeMember("NETWORK_REPLACE_ME", "MEMBER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member", "MEMBER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccIpv6rangeResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_name"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeName("NETWORK_REPLACE_ME", "NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeName("NETWORK_REPLACE_ME", "NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_Network(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_network"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeNetwork("NETWORK_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", "NETWORK_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeNetwork("NETWORK_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", "NETWORK_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_NetworkView(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_network_view"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeNetworkView("NETWORK_REPLACE_ME", "NETWORK_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeNetworkView("NETWORK_REPLACE_ME", "NETWORK_VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_OptionFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_option_filter_rules"
	var v dhcp.Ipv6range
	option_filter_rulesVal := []map[string]any{}
	option_filter_rulesValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeOptionFilterRules("NETWORK_REPLACE_ME", option_filter_rulesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_filter_rules", "OPTION_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeOptionFilterRules("NETWORK_REPLACE_ME", option_filter_rulesValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_filter_rules", "OPTION_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_PortControlBlackoutSetting(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_port_control_blackout_setting"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangePortControlBlackoutSetting("NETWORK_REPLACE_ME", "PORT_CONTROL_BLACKOUT_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port_control_blackout_setting", "PORT_CONTROL_BLACKOUT_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangePortControlBlackoutSetting("NETWORK_REPLACE_ME", "PORT_CONTROL_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port_control_blackout_setting", "PORT_CONTROL_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccIpv6rangeResource_RecycleLeases(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_recycle_leases"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeRecycleLeases("NETWORK_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeRecycleLeases("NETWORK_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_RestartIfNeeded(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_restart_if_needed"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeRestartIfNeeded("NETWORK_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "restart_if_needed", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeRestartIfNeeded("NETWORK_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "restart_if_needed", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_SamePortControlDiscoveryBlackout(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_same_port_control_discovery_blackout"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeSamePortControlDiscoveryBlackout("NETWORK_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeSamePortControlDiscoveryBlackout("NETWORK_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_ServerAssociationType(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_server_association_type"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeServerAssociationType("NETWORK_REPLACE_ME", "SERVER_ASSOCIATION_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "server_association_type", "SERVER_ASSOCIATION_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeServerAssociationType("NETWORK_REPLACE_ME", "SERVER_ASSOCIATION_TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "server_association_type", "SERVER_ASSOCIATION_TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_StartAddr(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_start_addr"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeStartAddr("NETWORK_REPLACE_ME", "START_ADDR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_addr", "START_ADDR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeStartAddr("NETWORK_REPLACE_ME", "START_ADDR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_addr", "START_ADDR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_SubscribeSettings(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_subscribe_settings"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeSubscribeSettings("NETWORK_REPLACE_ME", "SUBSCRIBE_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscribe_settings", "SUBSCRIBE_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeSubscribeSettings("NETWORK_REPLACE_ME", "SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscribe_settings", "SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccIpv6rangeResource_Template(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_template"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeTemplate("NETWORK_REPLACE_ME", "TEMPLATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template", "TEMPLATE_REPLACE_ME"),
				),
			},
			// Skip Update testing as this field cannot be updated
		},
	})
}

func TestAccIpv6rangeResource_UseBlackoutSetting(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_use_blackout_setting"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeUseBlackoutSetting("NETWORK_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeUseBlackoutSetting("NETWORK_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_UseDiscoveryBasicPollingSettings(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_use_discovery_basic_polling_settings"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeUseDiscoveryBasicPollingSettings("NETWORK_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeUseDiscoveryBasicPollingSettings("NETWORK_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_UseEnableDiscovery(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_use_enable_discovery"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeUseEnableDiscovery("NETWORK_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeUseEnableDiscovery("NETWORK_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_UseLogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_use_logic_filter_rules"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeUseLogicFilterRules("NETWORK_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeUseLogicFilterRules("NETWORK_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_UseRecycleLeases(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_use_recycle_leases"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeUseRecycleLeases("NETWORK_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeUseRecycleLeases("NETWORK_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_UseSubscribeSettings(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_use_subscribe_settings"
	var v dhcp.Ipv6range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeUseSubscribeSettings("NETWORK_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_subscribe_settings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeUseSubscribeSettings("NETWORK_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_subscribe_settings", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckIpv6rangeExists(ctx context.Context, resourceName string, v *dhcp.Ipv6range) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			Ipv6rangeAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForIpv6range).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetIpv6rangeResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetIpv6rangeResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckIpv6rangeDestroy(ctx context.Context, v *dhcp.Ipv6range) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			Ipv6rangeAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForIpv6range).
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

func testAccCheckIpv6rangeDisappears(ctx context.Context, v *dhcp.Ipv6range) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			Ipv6rangeAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccIpv6rangeImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.Attributes["ref"] == "" {
			return "", fmt.Errorf("ref is not set")
		}
		return rs.Primary.Attributes["ref"], nil
	}
}

func testAccIpv6rangeBasicConfig(view, startAddr, endAddr string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
	network_view = nios_ipam_network_view.test.name
}
`, startAddr, endAddr)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeAddressType(network string, addressType string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_address_type" {
    network = %q
    address_type = %q
}
`, network, addressType)
}

func testAccIpv6rangeCloudInfo(network string, cloudInfo string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_cloud_info" {
    network = %q
    cloud_info = %q
}
`, network, cloudInfo)
}

func testAccIpv6rangeComment(network string, comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_comment" {
    network = %q
    comment = %q
}
`, network, comment)
}

func testAccIpv6rangeDisable(network string, disable string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_disable" {
    network = %q
    disable = %q
}
`, network, disable)
}

func testAccIpv6rangeDiscoveryBasicPollSettings(network string, discoveryBasicPollSettings string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_discovery_basic_poll_settings" {
    network = %q
    discovery_basic_poll_settings = %q
}
`, network, discoveryBasicPollSettings)
}

func testAccIpv6rangeDiscoveryBlackoutSetting(network string, discoveryBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_discovery_blackout_setting" {
    network = %q
    discovery_blackout_setting = %q
}
`, network, discoveryBlackoutSetting)
}

func testAccIpv6rangeDiscoveryMember(network string, discoveryMember string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_discovery_member" {
    network = %q
    discovery_member = %q
}
`, network, discoveryMember)
}

func testAccIpv6rangeEnableDiscovery(network string, enableDiscovery string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_enable_discovery" {
    network = %q
    enable_discovery = %q
}
`, network, enableDiscovery)
}

func testAccIpv6rangeEnableImmediateDiscovery(network string, enableImmediateDiscovery string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_enable_immediate_discovery" {
    network = %q
    enable_immediate_discovery = %q
}
`, network, enableImmediateDiscovery)
}

func testAccIpv6rangeEndAddr(network string, endAddr string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_end_addr" {
    network = %q
    end_addr = %q
}
`, network, endAddr)
}

func testAccIpv6rangeExclude(network string, exclude []map[string]any) string {
	excludeStr := utils.ConvertSliceOfMapsToHCL(exclude)
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_exclude" {
    network = %q
    exclude = %q
}
`, network, excludeStr)
}

func testAccIpv6rangeExtAttrs(network string, extAttrs map[string]string) string {
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_extattrs" {
    network = %q
    extattrs = %q
}
`, network, extAttrsStr)
}

func testAccIpv6rangeIpv6EndPrefix(network string, ipv6EndPrefix string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_ipv6_end_prefix" {
    network = %q
    ipv6_end_prefix = %q
}
`, network, ipv6EndPrefix)
}

func testAccIpv6rangeIpv6PrefixBits(network string, ipv6PrefixBits string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_ipv6_prefix_bits" {
    network = %q
    ipv6_prefix_bits = %q
}
`, network, ipv6PrefixBits)
}

func testAccIpv6rangeIpv6StartPrefix(network string, ipv6StartPrefix string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_ipv6_start_prefix" {
    network = %q
    ipv6_start_prefix = %q
}
`, network, ipv6StartPrefix)
}

func testAccIpv6rangeLogicFilterRules(network string, logicFilterRules []map[string]any) string {
	logicFilterRulesStr := utils.ConvertSliceOfMapsToHCL(logicFilterRules)
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_logic_filter_rules" {
    network = %q
    logic_filter_rules = %q
}
`, network, logicFilterRulesStr)
}

func testAccIpv6rangeMember(network string, member string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_member" {
    network = %q
    member = %q
}
`, network, member)
}

func testAccIpv6rangeName(network string, name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_name" {
    network = %q
    name = %q
}
`, network, name)
}

func testAccIpv6rangeNetwork(network string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_network" {
    network = %q
}
`, network)
}

func testAccIpv6rangeNetworkView(network string, networkView string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_network_view" {
    network = %q
    network_view = %q
}
`, network, networkView)
}

func testAccIpv6rangeOptionFilterRules(network string, optionFilterRules []map[string]any) string {
	optionFilterRulesStr := utils.ConvertSliceOfMapsToHCL(optionFilterRules)
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_option_filter_rules" {
    network = %q
    option_filter_rules = %q
}
`, network, optionFilterRulesStr)
}

func testAccIpv6rangePortControlBlackoutSetting(network string, portControlBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_port_control_blackout_setting" {
    network = %q
    port_control_blackout_setting = %q
}
`, network, portControlBlackoutSetting)
}

func testAccIpv6rangeRecycleLeases(network string, recycleLeases string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_recycle_leases" {
    network = %q
    recycle_leases = %q
}
`, network, recycleLeases)
}

func testAccIpv6rangeRestartIfNeeded(network string, restartIfNeeded string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_restart_if_needed" {
    network = %q
    restart_if_needed = %q
}
`, network, restartIfNeeded)
}

func testAccIpv6rangeSamePortControlDiscoveryBlackout(network string, samePortControlDiscoveryBlackout string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_same_port_control_discovery_blackout" {
    network = %q
    same_port_control_discovery_blackout = %q
}
`, network, samePortControlDiscoveryBlackout)
}

func testAccIpv6rangeServerAssociationType(network string, serverAssociationType string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_server_association_type" {
    network = %q
    server_association_type = %q
}
`, network, serverAssociationType)
}

func testAccIpv6rangeStartAddr(network string, startAddr string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_start_addr" {
    network = %q
    start_addr = %q
}
`, network, startAddr)
}

func testAccIpv6rangeSubscribeSettings(network string, subscribeSettings string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_subscribe_settings" {
    network = %q
    subscribe_settings = %q
}
`, network, subscribeSettings)
}

func testAccIpv6rangeTemplate(network string, template string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_template" {
    network = %q
    template = %q
}
`, network, template)
}

func testAccIpv6rangeUseBlackoutSetting(network string, useBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_use_blackout_setting" {
    network = %q
    use_blackout_setting = %q
}
`, network, useBlackoutSetting)
}

func testAccIpv6rangeUseDiscoveryBasicPollingSettings(network string, useDiscoveryBasicPollingSettings string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_use_discovery_basic_polling_settings" {
    network = %q
    use_discovery_basic_polling_settings = %q
}
`, network, useDiscoveryBasicPollingSettings)
}

func testAccIpv6rangeUseEnableDiscovery(network string, useEnableDiscovery string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_use_enable_discovery" {
    network = %q
    use_enable_discovery = %q
}
`, network, useEnableDiscovery)
}

func testAccIpv6rangeUseLogicFilterRules(network string, useLogicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_use_logic_filter_rules" {
    network = %q
    use_logic_filter_rules = %q
}
`, network, useLogicFilterRules)
}

func testAccIpv6rangeUseRecycleLeases(network string, useRecycleLeases string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_use_recycle_leases" {
    network = %q
    use_recycle_leases = %q
}
`, network, useRecycleLeases)
}

func testAccIpv6rangeUseSubscribeSettings(network string, useSubscribeSettings string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_use_subscribe_settings" {
    network = %q
    use_subscribe_settings = %q
}
`, network, useSubscribeSettings)
}

func testAccBaseWithIpv6NetworkandView(view string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_view" "test" {
	name = %q
}

resource "nios_ipam_ipv6network" "test" {
    network = "14::/64"
	network_view = nios_ipam_network_view.test.name
}
`, view)
}
