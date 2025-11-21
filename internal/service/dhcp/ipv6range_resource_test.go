package dhcp_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
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
				ExpectError:                          regexp.MustCompile(`ImportStateVerify attributes not equivalent`),
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccIpv6rangeImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all"},
				ImportStateVerifyIdentifierAttribute: "ref",
				ExpectError:                          regexp.MustCompile(`ImportStateVerify attributes not equivalent`),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_AddressType(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_address_type"
	var v dhcp.Ipv6range
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeAddressType(view, "14::1", "14::10", "ADDRESS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address_type", "ADDRESS"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeAddressTypeUpdate(view, "PREFIX", "14:0:0:1000::", "14:0:0:1fff::", "80"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address_type", "PREFIX"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_CloudInfo(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_cloud_info"
	var v dhcp.Ipv6range
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeCloudInfo(view, "14::1", "14::10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.authority_type", "GM"),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.delegated_scope", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.owned_by_adaptor", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccIpv6rangeResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_comment"
	var v dhcp.Ipv6range
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeComment(view, "14::1", "14::10", "Comment for the Ipv6range object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the Ipv6range object"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeComment(view, "14::1", "14::10", "Updated comment for the Ipv6range object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the Ipv6range object"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_Disable(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_disable"
	var v dhcp.Ipv6range
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeDisable(view, "14::1", "14::10", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeDisable(view, "14::1", "14::10", "false"),
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
	view := acctest.RandomNameWithPrefix("network-view")
	discoveryBasicPollSettings := map[string]any{
		"auto_arp_refresh_before_switch_port_polling": true,
		"cli_collection":                      false,
		"complete_ping_sweep":                 false,
		"device_profile":                      false,
		"switch_port_data_collection_polling": "PERIODIC",
	}
	updatedDiscoveryBasicPollSettings := map[string]any{
		"auto_arp_refresh_before_switch_port_polling": true,
		"cli_collection":                      true,
		"complete_ping_sweep":                 false,
		"device_profile":                      false,
		"switch_port_data_collection_polling": "SCHEDULED",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeDiscoveryBasicPollSettings(view, "14::1", "14::10", discoveryBasicPollSettings, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.cli_collection", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.switch_port_data_collection_polling", "PERIODIC"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.auto_arp_refresh_before_switch_port_polling", "true"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.complete_ping_sweep", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.device_profile", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeDiscoveryBasicPollSettings(view, "14::1", "14::10", updatedDiscoveryBasicPollSettings, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.cli_collection", "true"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.switch_port_data_collection_polling", "SCHEDULED"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.auto_arp_refresh_before_switch_port_polling", "true"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.complete_ping_sweep", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.device_profile", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccIpv6rangeResource_DiscoveryBlackoutSetting(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_discovery_blackout_setting"
	var v dhcp.Ipv6range
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeDiscoveryBlackoutSetting(view, "14::1", "14::10", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_blackout_setting.enable_blackout", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccIpv6rangeResource_DiscoveryMember(t *testing.T) {
	t.Skip("Additional configuration is required to run this test")
	var resourceName = "nios_dhcp_ipv6range.test_discovery_member"
	var v dhcp.Ipv6range
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeDiscoveryMember(view, "14::1", "14::10", "infoblox.member1", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_member", "infoblox.member1"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeDiscoveryMember(view, "14::1", "14::10", "infoblox.member2", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_member", "infoblox.member2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_EnableDiscovery(t *testing.T) {
	t.Skip("Additional configuration is required to run this test")
	var resourceName = "nios_dhcp_ipv6range.test_enable_discovery"
	var v dhcp.Ipv6range
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeEnableDiscovery(view, "14::1", "14::10", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_discovery", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeEnableDiscovery(view, "14::1", "14::10", "false"),
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
	t.Skip("Additional configuration is required to run this test")
	var resourceName = "nios_dhcp_ipv6range.test_enable_immediate_discovery"
	var v dhcp.Ipv6range
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeEnableImmediateDiscovery(view, "14::1", "14::10", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_immediate_discovery", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeEnableImmediateDiscovery(view, "14::1", "14::10", "false"),
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeEndAddr(view, "14::1", "14::10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "end_addr", "14::10"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeEndAddr(view, "14::1", "14::20"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "end_addr", "14::20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6rangeResource_Exclude(t *testing.T) {
	var resourceName = "nios_dhcp_ipv6range.test_exclude"
	var v dhcp.Ipv6range
	view := acctest.RandomNameWithPrefix("network-view")
	excludeVal := []map[string]any{
		{
			"start_address": "14::3",
			"end_address":   "14::5",
			"comment":       "Exclude range 14::3 - 14::5",
		},
	}
	excludeValUpdated := []map[string]any{
		{
			"start_address": "14::4",
			"end_address":   "14::6",
			"comment":       "Updated exclude range 14::4 - 14::6",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeExclude(view, "14::1", "14::10", excludeVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "exclude.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "exclude.0.start_address", "14::3"),
					resource.TestCheckResourceAttr(resourceName, "exclude.0.end_address", "14::5"),
					resource.TestCheckResourceAttr(resourceName, "exclude.0.comment", "Exclude range 14::3 - 14::5"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeExclude(view, "14::1", "14::10", excludeValUpdated),
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
	view := acctest.RandomNameWithPrefix("network-view")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeExtAttrs(view, "14::1", "14::10", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeExtAttrs(view, "14::1", "14::10", map[string]string{
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeIpv6EndPrefix(view, "14::1", "14::10", "IPV6_END_PREFIX_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_end_prefix", "IPV6_END_PREFIX_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeIpv6EndPrefix(view, "14::1", "14::10", "IPV6_END_PREFIX_UPDATE_REPLACE_ME"),
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeIpv6PrefixBits(view, "14::1", "14::10", "IPV6_PREFIX_BITS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_prefix_bits", "IPV6_PREFIX_BITS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeIpv6PrefixBits(view, "14::1", "14::10", "IPV6_PREFIX_BITS_UPDATE_REPLACE_ME"),
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeIpv6StartPrefix(view, "14::1", "14::10", "IPV6_START_PREFIX_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_start_prefix", "IPV6_START_PREFIX_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeIpv6StartPrefix(view, "14::1", "14::10", "IPV6_START_PREFIX_UPDATE_REPLACE_ME"),
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
	view := acctest.RandomNameWithPrefix("network-view")
	logic_filter_rulesVal := []map[string]any{}
	logic_filter_rulesValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeLogicFilterRules(view, "14::1", "14::10", logic_filter_rulesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeLogicFilterRules(view, "14::1", "14::10", logic_filter_rulesValUpdated),
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeMember(view, "14::1", "14::10", "MEMBER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member", "MEMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeMember(view, "14::1", "14::10", "MEMBER_UPDATE_REPLACE_ME"),
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeName(view, "14::1", "14::10", "NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeName(view, "14::1", "14::10", "NAME_UPDATE_REPLACE_ME"),
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeNetwork(view, "14::1", "14::10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", "NETWORK_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeNetwork(view, "14::1", "14::10"),
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeNetworkView(view, "14::1", "14::10", "NETWORK_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeNetworkView(view, "14::1", "14::10", "NETWORK_VIEW_UPDATE_REPLACE_ME"),
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
	view := acctest.RandomNameWithPrefix("network-view")
	option_filter_rulesVal := []map[string]any{}
	option_filter_rulesValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeOptionFilterRules(view, "14::1", "14::10", option_filter_rulesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_filter_rules", "OPTION_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeOptionFilterRules(view, "14::1", "14::10", option_filter_rulesValUpdated),
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangePortControlBlackoutSetting(view, "14::1", "14::10", "PORT_CONTROL_BLACKOUT_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port_control_blackout_setting", "PORT_CONTROL_BLACKOUT_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangePortControlBlackoutSetting(view, "14::1", "14::10", "PORT_CONTROL_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeRecycleLeases(view, "14::1", "14::10", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeRecycleLeases(view, "14::1", "14::10", "false"),
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeRestartIfNeeded(view, "14::1", "14::10", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "restart_if_needed", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeRestartIfNeeded(view, "14::1", "14::10", "false"),
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeSamePortControlDiscoveryBlackout(view, "14::1", "14::10", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeSamePortControlDiscoveryBlackout(view, "14::1", "14::10", "false"),
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeServerAssociationType(view, "14::1", "14::10", "SERVER_ASSOCIATION_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "server_association_type", "SERVER_ASSOCIATION_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeServerAssociationType(view, "14::1", "14::10", "SERVER_ASSOCIATION_TYPE_UPDATE_REPLACE_ME"),
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeStartAddr(view, "14::1", "14::10", "START_ADDR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_addr", "START_ADDR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeStartAddr(view, "14::1", "14::10", "START_ADDR_UPDATE_REPLACE_ME"),
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeSubscribeSettings(view, "14::1", "14::10", "SUBSCRIBE_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscribe_settings", "SUBSCRIBE_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeSubscribeSettings(view, "14::1", "14::10", "SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeTemplate(view, "14::1", "14::10", "TEMPLATE_REPLACE_ME"),
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeUseBlackoutSetting(view, "14::1", "14::10", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeUseBlackoutSetting(view, "14::1", "14::10", "false"),
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeUseDiscoveryBasicPollingSettings(view, "14::1", "14::10", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeUseDiscoveryBasicPollingSettings(view, "14::1", "14::10", "false"),
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeUseEnableDiscovery(view, "14::1", "14::10", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeUseEnableDiscovery(view, "14::1", "14::10", "false"),
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeUseLogicFilterRules(view, "14::1", "14::10", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeUseLogicFilterRules(view, "14::1", "14::10", "false"),
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeUseRecycleLeases(view, "14::1", "14::10", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeUseRecycleLeases(view, "14::1", "14::10", "false"),
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
	view := acctest.RandomNameWithPrefix("network-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6rangeUseSubscribeSettings(view, "14::1", "14::10", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6rangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_subscribe_settings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6rangeUseSubscribeSettings(view, "14::1", "14::10", "false"),
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

func testAccIpv6rangeAddressType(view, startAddr, endAddr string, addressType string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_address_type" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    address_type = %q
}
`, startAddr, endAddr, addressType)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeAddressTypeUpdate(view, addressType, ipv6StartPrefix, ipv6EndPrefix string, prefixBits string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_address_type" {
    network = nios_ipam_ipv6network.test.network
    ipv6_start_prefix = %q
    ipv6_end_prefix = %q
    network_view = nios_ipam_network_view.test.name
    address_type = %q
	ipv6_prefix_bits = %q
}
`, ipv6StartPrefix, ipv6EndPrefix, addressType, prefixBits)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeCloudInfo(view, startAddr, endAddr string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_cloud_info" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
}
`, startAddr, endAddr)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeComment(view, startAddr, endAddr string, comment string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_comment" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    comment = %q
}
`, startAddr, endAddr, comment)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeDisable(view, startAddr, endAddr string, disable string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_disable" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    disable = %q
}
`, startAddr, endAddr, disable)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeDiscoveryBasicPollSettings(view, startAddr, endAddr string, discoveryBasicPollSettings map[string]any, useDiscoveryBasicPollingSettings string) string {
	discoveryBasicPollSettingsStr := utils.ConvertMapToHCL(discoveryBasicPollSettings)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_discovery_basic_poll_settings" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    discovery_basic_poll_settings = %s
	use_discovery_basic_polling_settings = %q
}
`, startAddr, endAddr, discoveryBasicPollSettingsStr, useDiscoveryBasicPollingSettings)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeDiscoveryBlackoutSetting(view, startAddr, endAddr string, useBlackoutSetting string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_discovery_blackout_setting" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
	use_blackout_setting = %q
}
`, startAddr, endAddr, useBlackoutSetting)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeDiscoveryMember(view, startAddr, endAddr string, discoveryMember string, useEnableDiscovery string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_discovery_member" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    discovery_member = %q
	use_enable_discovery = %q
}
`, startAddr, endAddr, discoveryMember, useEnableDiscovery)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeEnableDiscovery(view, startAddr, endAddr string, enableDiscovery string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_enable_discovery" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    enable_discovery = %q
}
`, startAddr, endAddr, enableDiscovery)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeEnableImmediateDiscovery(view, startAddr, endAddr string, enableImmediateDiscovery string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_enable_immediate_discovery" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    enable_immediate_discovery = %q
}
`, startAddr, endAddr, enableImmediateDiscovery)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeEndAddr(view, startAddr, endAddr string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_end_addr" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
}
`, startAddr, endAddr)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeExclude(view, startAddr, endAddr string, exclude []map[string]any) string {
	excludeStr := utils.ConvertSliceOfMapsToHCL(exclude)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_exclude" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    exclude = %q
}
`, startAddr, endAddr, excludeStr)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeExtAttrs(view, startAddr, endAddr string, extAttrs map[string]string) string {
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_extattrs" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    extattrs = %s
}
`, startAddr, endAddr, extAttrsStr)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeIpv6EndPrefix(view, startAddr, endAddr string, ipv6EndPrefix string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_ipv6_end_prefix" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    ipv6_end_prefix = %q
}
`, startAddr, endAddr, ipv6EndPrefix)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeIpv6PrefixBits(view, startAddr, endAddr string, ipv6PrefixBits string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_ipv6_prefix_bits" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    ipv6_prefix_bits = %q
}
`, startAddr, endAddr, ipv6PrefixBits)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeIpv6StartPrefix(view, startAddr, endAddr string, ipv6StartPrefix string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_ipv6_start_prefix" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    ipv6_start_prefix = %q
}
`, startAddr, endAddr, ipv6StartPrefix)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeLogicFilterRules(view, startAddr, endAddr string, logicFilterRules []map[string]any) string {
	logicFilterRulesStr := utils.ConvertSliceOfMapsToHCL(logicFilterRules)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_logic_filter_rules" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    logic_filter_rules = %q
}
`, startAddr, endAddr, logicFilterRulesStr)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeMember(view, startAddr, endAddr string, member string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_member" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    member = %q
}
`, startAddr, endAddr, member)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeName(view, startAddr, endAddr string, name string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_name" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    name = %q
}
`, startAddr, endAddr, name)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeNetwork(view, startAddr, endAddr string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_network" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
}
`, startAddr, endAddr)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeNetworkView(view, startAddr, endAddr string, networkView string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_network_view" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = %q
}
`, startAddr, endAddr, networkView)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeOptionFilterRules(view, startAddr, endAddr string, optionFilterRules []map[string]any) string {
	optionFilterRulesStr := utils.ConvertSliceOfMapsToHCL(optionFilterRules)
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_option_filter_rules" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    option_filter_rules = %q
}
`, startAddr, endAddr, optionFilterRulesStr)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangePortControlBlackoutSetting(view, startAddr, endAddr string, portControlBlackoutSetting string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_port_control_blackout_setting" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    port_control_blackout_setting = %q
}
`, startAddr, endAddr, portControlBlackoutSetting)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeRecycleLeases(view, startAddr, endAddr string, recycleLeases string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_recycle_leases" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    recycle_leases = %q
}
`, startAddr, endAddr, recycleLeases)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeRestartIfNeeded(view, startAddr, endAddr string, restartIfNeeded string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_restart_if_needed" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    restart_if_needed = %q
}
`, startAddr, endAddr, restartIfNeeded)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeSamePortControlDiscoveryBlackout(view, startAddr, endAddr string, samePortControlDiscoveryBlackout string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_same_port_control_discovery_blackout" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    same_port_control_discovery_blackout = %q
}
`, startAddr, endAddr, samePortControlDiscoveryBlackout)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeServerAssociationType(view, startAddr, endAddr string, serverAssociationType string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_server_association_type" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    server_association_type = %q
}
`, startAddr, endAddr, serverAssociationType)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeStartAddr(view, startAddr, endAddr string, startAddrValue string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_start_addr" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
}
`, startAddrValue, endAddr)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeSubscribeSettings(view, startAddr, endAddr string, subscribeSettings string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_subscribe_settings" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    subscribe_settings = %q
}
`, startAddr, endAddr, subscribeSettings)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeTemplate(view, startAddr, endAddr string, template string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_template" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    template = %q
}
`, startAddr, endAddr, template)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeUseBlackoutSetting(view, startAddr, endAddr string, useBlackoutSetting string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_use_blackout_setting" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    use_blackout_setting = %q
}
`, startAddr, endAddr, useBlackoutSetting)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeUseDiscoveryBasicPollingSettings(view, startAddr, endAddr string, useDiscoveryBasicPollingSettings string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_use_discovery_basic_polling_settings" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    use_discovery_basic_polling_settings = %q
}
`, startAddr, endAddr, useDiscoveryBasicPollingSettings)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeUseEnableDiscovery(view, startAddr, endAddr string, useEnableDiscovery string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_use_enable_discovery" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    use_enable_discovery = %q
}
`, startAddr, endAddr, useEnableDiscovery)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeUseLogicFilterRules(view, startAddr, endAddr string, useLogicFilterRules string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_use_logic_filter_rules" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    use_logic_filter_rules = %q
}
`, startAddr, endAddr, useLogicFilterRules)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeUseRecycleLeases(view, startAddr, endAddr string, useRecycleLeases string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_use_recycle_leases" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    use_recycle_leases = %q
}
`, startAddr, endAddr, useRecycleLeases)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
}

func testAccIpv6rangeUseSubscribeSettings(view, startAddr, endAddr string, useSubscribeSettings string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_ipv6range" "test_use_subscribe_settings" {
    network = nios_ipam_ipv6network.test.network
    start_addr = %q
    end_addr = %q
    network_view = nios_ipam_network_view.test.name
    use_subscribe_settings = %q
}
`, startAddr, endAddr, useSubscribeSettings)
	return strings.Join([]string{testAccBaseWithIpv6NetworkandView(view), config}, "")
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
