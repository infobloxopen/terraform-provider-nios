package dhcp_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dhcp"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/acctest"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/utils"
)

var readableAttributesForRange = "always_update_dns,bootfile,bootserver,cloud_info,comment,ddns_domainname,ddns_generate_hostname,deny_all_clients,deny_bootp,dhcp_utilization,dhcp_utilization_status,disable,discover_now_status,discovery_basic_poll_settings,discovery_blackout_setting,discovery_member,dynamic_hosts,email_list,enable_ddns,enable_dhcp_thresholds,enable_discovery,enable_email_warnings,enable_ifmap_publishing,enable_pxe_lease_time,enable_snmp_warnings,end_addr,endpoint_sources,exclude,extattrs,failover_association,fingerprint_filter_rules,high_water_mark,high_water_mark_reset,ignore_dhcp_option_list_request,ignore_id,ignore_mac_addresses,is_split_scope,known_clients,lease_scavenge_time,logic_filter_rules,low_water_mark,low_water_mark_reset,mac_filter_rules,member,ms_ad_user_data,ms_options,ms_server,nac_filter_rules,name,network,network_view,nextserver,option_filter_rules,options,port_control_blackout_setting,pxe_lease_time,recycle_leases,relay_agent_filter_rules,same_port_control_discovery_blackout,server_association_type,start_addr,static_hosts,subscribe_settings,total_hosts,unknown_clients,update_dns_on_lease_renewal,use_blackout_setting,use_bootfile,use_bootserver,use_ddns_domainname,use_ddns_generate_hostname,use_deny_bootp,use_discovery_basic_polling_settings,use_email_list,use_enable_ddns,use_enable_dhcp_thresholds,use_enable_discovery,use_enable_ifmap_publishing,use_ignore_dhcp_option_list_request,use_ignore_id,use_known_clients,use_lease_scavenge_time,use_logic_filter_rules,use_ms_options,use_nextserver,use_options,use_pxe_lease_time,use_recycle_leases,use_subscribe_settings,use_unknown_clients,use_update_dns_on_lease_renewal"

func TestAccRangeResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_range.test"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_range.test"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRangeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRangeBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					testAccCheckRangeDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRangeResource_Ref(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_ref"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_AlwaysUpdateDns(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_always_update_dns"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeAlwaysUpdateDns("ALWAYS_UPDATE_DNS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "always_update_dns", "ALWAYS_UPDATE_DNS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeAlwaysUpdateDns("ALWAYS_UPDATE_DNS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "always_update_dns", "ALWAYS_UPDATE_DNS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Bootfile(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_bootfile"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeBootfile("BOOTFILE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "BOOTFILE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeBootfile("BOOTFILE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "BOOTFILE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Bootserver(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_bootserver"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeBootserver("BOOTSERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "BOOTSERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeBootserver("BOOTSERVER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "BOOTSERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_CloudInfo(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_cloud_info"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeCloudInfo("CLOUD_INFO_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_info", "CLOUD_INFO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeCloudInfo("CLOUD_INFO_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_info", "CLOUD_INFO_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_comment"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_DdnsDomainname(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_ddns_domainname"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeDdnsDomainname("DDNS_DOMAINNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "DDNS_DOMAINNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeDdnsDomainname("DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_DdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_ddns_generate_hostname"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeDdnsGenerateHostname("DDNS_GENERATE_HOSTNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "DDNS_GENERATE_HOSTNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeDdnsGenerateHostname("DDNS_GENERATE_HOSTNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "DDNS_GENERATE_HOSTNAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_DenyAllClients(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_deny_all_clients"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeDenyAllClients("DENY_ALL_CLIENTS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_all_clients", "DENY_ALL_CLIENTS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeDenyAllClients("DENY_ALL_CLIENTS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_all_clients", "DENY_ALL_CLIENTS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_DenyBootp(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_deny_bootp"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeDenyBootp("DENY_BOOTP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "DENY_BOOTP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeDenyBootp("DENY_BOOTP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "DENY_BOOTP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Disable(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_disable"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_DiscoveryBasicPollSettings(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_discovery_basic_poll_settings"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeDiscoveryBasicPollSettings("DISCOVERY_BASIC_POLL_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings", "DISCOVERY_BASIC_POLL_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeDiscoveryBasicPollSettings("DISCOVERY_BASIC_POLL_SETTINGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings", "DISCOVERY_BASIC_POLL_SETTINGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_DiscoveryBlackoutSetting(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_discovery_blackout_setting"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeDiscoveryBlackoutSetting("DISCOVERY_BLACKOUT_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_blackout_setting", "DISCOVERY_BLACKOUT_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeDiscoveryBlackoutSetting("DISCOVERY_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_blackout_setting", "DISCOVERY_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_DiscoveryMember(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_discovery_member"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeDiscoveryMember("DISCOVERY_MEMBER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_member", "DISCOVERY_MEMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeDiscoveryMember("DISCOVERY_MEMBER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_member", "DISCOVERY_MEMBER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_EmailList(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_email_list"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeEmailList("EMAIL_LIST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_list", "EMAIL_LIST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeEmailList("EMAIL_LIST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_list", "EMAIL_LIST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_EnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_enable_ddns"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeEnableDdns("ENABLE_DDNS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "ENABLE_DDNS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeEnableDdns("ENABLE_DDNS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "ENABLE_DDNS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_EnableDhcpThresholds(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_enable_dhcp_thresholds"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeEnableDhcpThresholds("ENABLE_DHCP_THRESHOLDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_dhcp_thresholds", "ENABLE_DHCP_THRESHOLDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeEnableDhcpThresholds("ENABLE_DHCP_THRESHOLDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_dhcp_thresholds", "ENABLE_DHCP_THRESHOLDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_EnableDiscovery(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_enable_discovery"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeEnableDiscovery("ENABLE_DISCOVERY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_discovery", "ENABLE_DISCOVERY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeEnableDiscovery("ENABLE_DISCOVERY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_discovery", "ENABLE_DISCOVERY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_EnableEmailWarnings(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_enable_email_warnings"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeEnableEmailWarnings("ENABLE_EMAIL_WARNINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_email_warnings", "ENABLE_EMAIL_WARNINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeEnableEmailWarnings("ENABLE_EMAIL_WARNINGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_email_warnings", "ENABLE_EMAIL_WARNINGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_EnableIfmapPublishing(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_enable_ifmap_publishing"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeEnableIfmapPublishing("ENABLE_IFMAP_PUBLISHING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ifmap_publishing", "ENABLE_IFMAP_PUBLISHING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeEnableIfmapPublishing("ENABLE_IFMAP_PUBLISHING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ifmap_publishing", "ENABLE_IFMAP_PUBLISHING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_EnableImmediateDiscovery(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_enable_immediate_discovery"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeEnableImmediateDiscovery("ENABLE_IMMEDIATE_DISCOVERY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_immediate_discovery", "ENABLE_IMMEDIATE_DISCOVERY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeEnableImmediateDiscovery("ENABLE_IMMEDIATE_DISCOVERY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_immediate_discovery", "ENABLE_IMMEDIATE_DISCOVERY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_EnablePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_enable_pxe_lease_time"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeEnablePxeLeaseTime("ENABLE_PXE_LEASE_TIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "ENABLE_PXE_LEASE_TIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeEnablePxeLeaseTime("ENABLE_PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "ENABLE_PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_EnableSnmpWarnings(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_enable_snmp_warnings"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeEnableSnmpWarnings("ENABLE_SNMP_WARNINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_snmp_warnings", "ENABLE_SNMP_WARNINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeEnableSnmpWarnings("ENABLE_SNMP_WARNINGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_snmp_warnings", "ENABLE_SNMP_WARNINGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_EndAddr(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_end_addr"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeEndAddr("END_ADDR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "end_addr", "END_ADDR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeEndAddr("END_ADDR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "end_addr", "END_ADDR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Exclude(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_exclude"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeExclude("EXCLUDE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "exclude", "EXCLUDE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeExclude("EXCLUDE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "exclude", "EXCLUDE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_extattrs"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_FailoverAssociation(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_failover_association"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeFailoverAssociation("FAILOVER_ASSOCIATION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "failover_association", "FAILOVER_ASSOCIATION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeFailoverAssociation("FAILOVER_ASSOCIATION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "failover_association", "FAILOVER_ASSOCIATION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_FingerprintFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_fingerprint_filter_rules"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeFingerprintFilterRules("FINGERPRINT_FILTER_RULES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fingerprint_filter_rules", "FINGERPRINT_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeFingerprintFilterRules("FINGERPRINT_FILTER_RULES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fingerprint_filter_rules", "FINGERPRINT_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_HighWaterMark(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_high_water_mark"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeHighWaterMark("HIGH_WATER_MARK_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "HIGH_WATER_MARK_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeHighWaterMark("HIGH_WATER_MARK_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "HIGH_WATER_MARK_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_HighWaterMarkReset(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_high_water_mark_reset"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeHighWaterMarkReset("HIGH_WATER_MARK_RESET_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "HIGH_WATER_MARK_RESET_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeHighWaterMarkReset("HIGH_WATER_MARK_RESET_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "HIGH_WATER_MARK_RESET_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_IgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_ignore_dhcp_option_list_request"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeIgnoreDhcpOptionListRequest("IGNORE_DHCP_OPTION_LIST_REQUEST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "IGNORE_DHCP_OPTION_LIST_REQUEST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeIgnoreDhcpOptionListRequest("IGNORE_DHCP_OPTION_LIST_REQUEST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "IGNORE_DHCP_OPTION_LIST_REQUEST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_IgnoreId(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_ignore_id"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeIgnoreId("IGNORE_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_id", "IGNORE_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeIgnoreId("IGNORE_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_id", "IGNORE_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_IgnoreMacAddresses(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_ignore_mac_addresses"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeIgnoreMacAddresses("IGNORE_MAC_ADDRESSES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_mac_addresses", "IGNORE_MAC_ADDRESSES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeIgnoreMacAddresses("IGNORE_MAC_ADDRESSES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_mac_addresses", "IGNORE_MAC_ADDRESSES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_KnownClients(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_known_clients"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeKnownClients("KNOWN_CLIENTS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "known_clients", "KNOWN_CLIENTS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeKnownClients("KNOWN_CLIENTS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "known_clients", "KNOWN_CLIENTS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_LeaseScavengeTime(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_lease_scavenge_time"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeLeaseScavengeTime("LEASE_SCAVENGE_TIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "LEASE_SCAVENGE_TIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeLeaseScavengeTime("LEASE_SCAVENGE_TIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "LEASE_SCAVENGE_TIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_LogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_logic_filter_rules"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeLogicFilterRules("LOGIC_FILTER_RULES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeLogicFilterRules("LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_LowWaterMark(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_low_water_mark"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeLowWaterMark("LOW_WATER_MARK_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "LOW_WATER_MARK_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeLowWaterMark("LOW_WATER_MARK_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "LOW_WATER_MARK_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_LowWaterMarkReset(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_low_water_mark_reset"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeLowWaterMarkReset("LOW_WATER_MARK_RESET_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "LOW_WATER_MARK_RESET_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeLowWaterMarkReset("LOW_WATER_MARK_RESET_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "LOW_WATER_MARK_RESET_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_MacFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_mac_filter_rules"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeMacFilterRules("MAC_FILTER_RULES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mac_filter_rules", "MAC_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeMacFilterRules("MAC_FILTER_RULES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mac_filter_rules", "MAC_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Member(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_member"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeMember("MEMBER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member", "MEMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeMember("MEMBER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member", "MEMBER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_MsAdUserData(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_ms_ad_user_data"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeMsAdUserData("MS_AD_USER_DATA_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_user_data", "MS_AD_USER_DATA_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeMsAdUserData("MS_AD_USER_DATA_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_user_data", "MS_AD_USER_DATA_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_MsOptions(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_ms_options"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeMsOptions("MS_OPTIONS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_options", "MS_OPTIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeMsOptions("MS_OPTIONS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_options", "MS_OPTIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_MsServer(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_ms_server"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeMsServer("MS_SERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_server", "MS_SERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeMsServer("MS_SERVER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_server", "MS_SERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_NacFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_nac_filter_rules"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeNacFilterRules("NAC_FILTER_RULES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nac_filter_rules", "NAC_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeNacFilterRules("NAC_FILTER_RULES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nac_filter_rules", "NAC_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_name"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Network(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_network"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeNetwork("NETWORK_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", "NETWORK_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeNetwork("NETWORK_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", "NETWORK_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_NetworkView(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_network_view"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeNetworkView("NETWORK_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeNetworkView("NETWORK_VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Nextserver(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_nextserver"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeNextserver("NEXTSERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", "NEXTSERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeNextserver("NEXTSERVER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", "NEXTSERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_OptionFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_option_filter_rules"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeOptionFilterRules("OPTION_FILTER_RULES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_filter_rules", "OPTION_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeOptionFilterRules("OPTION_FILTER_RULES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_filter_rules", "OPTION_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Options(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_options"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeOptions("OPTIONS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeOptions("OPTIONS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_PortControlBlackoutSetting(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_port_control_blackout_setting"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangePortControlBlackoutSetting("PORT_CONTROL_BLACKOUT_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port_control_blackout_setting", "PORT_CONTROL_BLACKOUT_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangePortControlBlackoutSetting("PORT_CONTROL_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port_control_blackout_setting", "PORT_CONTROL_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_PxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_pxe_lease_time"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangePxeLeaseTime("PXE_LEASE_TIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "PXE_LEASE_TIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangePxeLeaseTime("PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_RecycleLeases(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_recycle_leases"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeRecycleLeases("RECYCLE_LEASES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "RECYCLE_LEASES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeRecycleLeases("RECYCLE_LEASES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "RECYCLE_LEASES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_RelayAgentFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_relay_agent_filter_rules"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeRelayAgentFilterRules("RELAY_AGENT_FILTER_RULES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "relay_agent_filter_rules", "RELAY_AGENT_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeRelayAgentFilterRules("RELAY_AGENT_FILTER_RULES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "relay_agent_filter_rules", "RELAY_AGENT_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_RestartIfNeeded(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_restart_if_needed"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeRestartIfNeeded("RESTART_IF_NEEDED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "restart_if_needed", "RESTART_IF_NEEDED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeRestartIfNeeded("RESTART_IF_NEEDED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "restart_if_needed", "RESTART_IF_NEEDED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_SamePortControlDiscoveryBlackout(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_same_port_control_discovery_blackout"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeSamePortControlDiscoveryBlackout("SAME_PORT_CONTROL_DISCOVERY_BLACKOUT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "SAME_PORT_CONTROL_DISCOVERY_BLACKOUT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeSamePortControlDiscoveryBlackout("SAME_PORT_CONTROL_DISCOVERY_BLACKOUT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "SAME_PORT_CONTROL_DISCOVERY_BLACKOUT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_ServerAssociationType(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_server_association_type"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeServerAssociationType("SERVER_ASSOCIATION_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "server_association_type", "SERVER_ASSOCIATION_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeServerAssociationType("SERVER_ASSOCIATION_TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "server_association_type", "SERVER_ASSOCIATION_TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_SplitMember(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_split_member"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeSplitMember("SPLIT_MEMBER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "split_member", "SPLIT_MEMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeSplitMember("SPLIT_MEMBER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "split_member", "SPLIT_MEMBER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_SplitScopeExclusionPercent(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_split_scope_exclusion_percent"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeSplitScopeExclusionPercent("SPLIT_SCOPE_EXCLUSION_PERCENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "split_scope_exclusion_percent", "SPLIT_SCOPE_EXCLUSION_PERCENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeSplitScopeExclusionPercent("SPLIT_SCOPE_EXCLUSION_PERCENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "split_scope_exclusion_percent", "SPLIT_SCOPE_EXCLUSION_PERCENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_StartAddr(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_start_addr"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeStartAddr("START_ADDR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_addr", "START_ADDR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeStartAddr("START_ADDR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_addr", "START_ADDR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_SubscribeSettings(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_subscribe_settings"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeSubscribeSettings("SUBSCRIBE_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscribe_settings", "SUBSCRIBE_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeSubscribeSettings("SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscribe_settings", "SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Template(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_template"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeTemplate("TEMPLATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template", "TEMPLATE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeTemplate("TEMPLATE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template", "TEMPLATE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UnknownClients(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_unknown_clients"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUnknownClients("UNKNOWN_CLIENTS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "unknown_clients", "UNKNOWN_CLIENTS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUnknownClients("UNKNOWN_CLIENTS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "unknown_clients", "UNKNOWN_CLIENTS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_update_dns_on_lease_renewal"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUpdateDnsOnLeaseRenewal("UPDATE_DNS_ON_LEASE_RENEWAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "UPDATE_DNS_ON_LEASE_RENEWAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUpdateDnsOnLeaseRenewal("UPDATE_DNS_ON_LEASE_RENEWAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "UPDATE_DNS_ON_LEASE_RENEWAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseBlackoutSetting(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_blackout_setting"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseBlackoutSetting("USE_BLACKOUT_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "USE_BLACKOUT_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseBlackoutSetting("USE_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "USE_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseBootfile(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_bootfile"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseBootfile("USE_BOOTFILE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "USE_BOOTFILE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseBootfile("USE_BOOTFILE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "USE_BOOTFILE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseBootserver(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_bootserver"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseBootserver("USE_BOOTSERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "USE_BOOTSERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseBootserver("USE_BOOTSERVER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "USE_BOOTSERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseDdnsDomainname(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_ddns_domainname"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseDdnsDomainname("USE_DDNS_DOMAINNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "USE_DDNS_DOMAINNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseDdnsDomainname("USE_DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "USE_DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseDdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_ddns_generate_hostname"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseDdnsGenerateHostname("USE_DDNS_GENERATE_HOSTNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "USE_DDNS_GENERATE_HOSTNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseDdnsGenerateHostname("USE_DDNS_GENERATE_HOSTNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "USE_DDNS_GENERATE_HOSTNAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseDenyBootp(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_deny_bootp"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseDenyBootp("USE_DENY_BOOTP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "USE_DENY_BOOTP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseDenyBootp("USE_DENY_BOOTP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "USE_DENY_BOOTP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseDiscoveryBasicPollingSettings(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_discovery_basic_polling_settings"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseDiscoveryBasicPollingSettings("USE_DISCOVERY_BASIC_POLLING_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "USE_DISCOVERY_BASIC_POLLING_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseDiscoveryBasicPollingSettings("USE_DISCOVERY_BASIC_POLLING_SETTINGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "USE_DISCOVERY_BASIC_POLLING_SETTINGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseEmailList(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_email_list"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseEmailList("USE_EMAIL_LIST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_email_list", "USE_EMAIL_LIST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseEmailList("USE_EMAIL_LIST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_email_list", "USE_EMAIL_LIST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseEnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_enable_ddns"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseEnableDdns("USE_ENABLE_DDNS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "USE_ENABLE_DDNS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseEnableDdns("USE_ENABLE_DDNS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "USE_ENABLE_DDNS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseEnableDhcpThresholds(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_enable_dhcp_thresholds"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseEnableDhcpThresholds("USE_ENABLE_DHCP_THRESHOLDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "USE_ENABLE_DHCP_THRESHOLDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseEnableDhcpThresholds("USE_ENABLE_DHCP_THRESHOLDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "USE_ENABLE_DHCP_THRESHOLDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseEnableDiscovery(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_enable_discovery"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseEnableDiscovery("USE_ENABLE_DISCOVERY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "USE_ENABLE_DISCOVERY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseEnableDiscovery("USE_ENABLE_DISCOVERY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "USE_ENABLE_DISCOVERY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseEnableIfmapPublishing(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_enable_ifmap_publishing"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseEnableIfmapPublishing("USE_ENABLE_IFMAP_PUBLISHING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ifmap_publishing", "USE_ENABLE_IFMAP_PUBLISHING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseEnableIfmapPublishing("USE_ENABLE_IFMAP_PUBLISHING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ifmap_publishing", "USE_ENABLE_IFMAP_PUBLISHING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseIgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_ignore_dhcp_option_list_request"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseIgnoreDhcpOptionListRequest("USE_IGNORE_DHCP_OPTION_LIST_REQUEST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "USE_IGNORE_DHCP_OPTION_LIST_REQUEST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseIgnoreDhcpOptionListRequest("USE_IGNORE_DHCP_OPTION_LIST_REQUEST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "USE_IGNORE_DHCP_OPTION_LIST_REQUEST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseIgnoreId(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_ignore_id"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseIgnoreId("USE_IGNORE_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_id", "USE_IGNORE_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseIgnoreId("USE_IGNORE_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_id", "USE_IGNORE_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseKnownClients(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_known_clients"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseKnownClients("USE_KNOWN_CLIENTS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_known_clients", "USE_KNOWN_CLIENTS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseKnownClients("USE_KNOWN_CLIENTS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_known_clients", "USE_KNOWN_CLIENTS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseLeaseScavengeTime(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_lease_scavenge_time"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseLeaseScavengeTime("USE_LEASE_SCAVENGE_TIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "USE_LEASE_SCAVENGE_TIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseLeaseScavengeTime("USE_LEASE_SCAVENGE_TIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "USE_LEASE_SCAVENGE_TIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseLogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_logic_filter_rules"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseLogicFilterRules("USE_LOGIC_FILTER_RULES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "USE_LOGIC_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseLogicFilterRules("USE_LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "USE_LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseMsOptions(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_ms_options"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseMsOptions("USE_MS_OPTIONS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ms_options", "USE_MS_OPTIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseMsOptions("USE_MS_OPTIONS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ms_options", "USE_MS_OPTIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseNextserver(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_nextserver"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseNextserver("USE_NEXTSERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "USE_NEXTSERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseNextserver("USE_NEXTSERVER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "USE_NEXTSERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseOptions(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_options"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseOptions("USE_OPTIONS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "USE_OPTIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseOptions("USE_OPTIONS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "USE_OPTIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UsePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_pxe_lease_time"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUsePxeLeaseTime("USE_PXE_LEASE_TIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "USE_PXE_LEASE_TIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUsePxeLeaseTime("USE_PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "USE_PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseRecycleLeases(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_recycle_leases"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseRecycleLeases("USE_RECYCLE_LEASES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "USE_RECYCLE_LEASES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseRecycleLeases("USE_RECYCLE_LEASES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "USE_RECYCLE_LEASES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseSubscribeSettings(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_subscribe_settings"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseSubscribeSettings("USE_SUBSCRIBE_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_subscribe_settings", "USE_SUBSCRIBE_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseSubscribeSettings("USE_SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_subscribe_settings", "USE_SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseUnknownClients(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_unknown_clients"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseUnknownClients("USE_UNKNOWN_CLIENTS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_unknown_clients", "USE_UNKNOWN_CLIENTS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseUnknownClients("USE_UNKNOWN_CLIENTS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_unknown_clients", "USE_UNKNOWN_CLIENTS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseUpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_update_dns_on_lease_renewal"
	var v dhcp.Range

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseUpdateDnsOnLeaseRenewal("USE_UPDATE_DNS_ON_LEASE_RENEWAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "USE_UPDATE_DNS_ON_LEASE_RENEWAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseUpdateDnsOnLeaseRenewal("USE_UPDATE_DNS_ON_LEASE_RENEWAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "USE_UPDATE_DNS_ON_LEASE_RENEWAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRangeExists(ctx context.Context, resourceName string, v *dhcp.Range) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			RangeAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRange).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRangeResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRangeResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRangeDestroy(ctx context.Context, v *dhcp.Range) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			RangeAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRange).
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

func testAccCheckRangeDisappears(ctx context.Context, v *dhcp.Range) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			RangeAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRangeBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test" {
}
`)
}

func testAccRangeRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccRangeAlwaysUpdateDns(alwaysUpdateDns string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_always_update_dns" {
    always_update_dns = %q
}
`, alwaysUpdateDns)
}

func testAccRangeBootfile(bootfile string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_bootfile" {
    bootfile = %q
}
`, bootfile)
}

func testAccRangeBootserver(bootserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_bootserver" {
    bootserver = %q
}
`, bootserver)
}

func testAccRangeCloudInfo(cloudInfo string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_cloud_info" {
    cloud_info = %q
}
`, cloudInfo)
}

func testAccRangeComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccRangeDdnsDomainname(ddnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_ddns_domainname" {
    ddns_domainname = %q
}
`, ddnsDomainname)
}

func testAccRangeDdnsGenerateHostname(ddnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_ddns_generate_hostname" {
    ddns_generate_hostname = %q
}
`, ddnsGenerateHostname)
}

func testAccRangeDenyAllClients(denyAllClients string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_deny_all_clients" {
    deny_all_clients = %q
}
`, denyAllClients)
}

func testAccRangeDenyBootp(denyBootp string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_deny_bootp" {
    deny_bootp = %q
}
`, denyBootp)
}

func testAccRangeDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccRangeDiscoveryBasicPollSettings(discoveryBasicPollSettings string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_discovery_basic_poll_settings" {
    discovery_basic_poll_settings = %q
}
`, discoveryBasicPollSettings)
}

func testAccRangeDiscoveryBlackoutSetting(discoveryBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_discovery_blackout_setting" {
    discovery_blackout_setting = %q
}
`, discoveryBlackoutSetting)
}

func testAccRangeDiscoveryMember(discoveryMember string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_discovery_member" {
    discovery_member = %q
}
`, discoveryMember)
}

func testAccRangeEmailList(emailList string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_email_list" {
    email_list = %q
}
`, emailList)
}

func testAccRangeEnableDdns(enableDdns string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_enable_ddns" {
    enable_ddns = %q
}
`, enableDdns)
}

func testAccRangeEnableDhcpThresholds(enableDhcpThresholds string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_enable_dhcp_thresholds" {
    enable_dhcp_thresholds = %q
}
`, enableDhcpThresholds)
}

func testAccRangeEnableDiscovery(enableDiscovery string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_enable_discovery" {
    enable_discovery = %q
}
`, enableDiscovery)
}

func testAccRangeEnableEmailWarnings(enableEmailWarnings string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_enable_email_warnings" {
    enable_email_warnings = %q
}
`, enableEmailWarnings)
}

func testAccRangeEnableIfmapPublishing(enableIfmapPublishing string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_enable_ifmap_publishing" {
    enable_ifmap_publishing = %q
}
`, enableIfmapPublishing)
}

func testAccRangeEnableImmediateDiscovery(enableImmediateDiscovery string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_enable_immediate_discovery" {
    enable_immediate_discovery = %q
}
`, enableImmediateDiscovery)
}

func testAccRangeEnablePxeLeaseTime(enablePxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_enable_pxe_lease_time" {
    enable_pxe_lease_time = %q
}
`, enablePxeLeaseTime)
}

func testAccRangeEnableSnmpWarnings(enableSnmpWarnings string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_enable_snmp_warnings" {
    enable_snmp_warnings = %q
}
`, enableSnmpWarnings)
}

func testAccRangeEndAddr(endAddr string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_end_addr" {
    end_addr = %q
}
`, endAddr)
}

func testAccRangeExclude(exclude string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_exclude" {
    exclude = %q
}
`, exclude)
}

func testAccRangeExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccRangeFailoverAssociation(failoverAssociation string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_failover_association" {
    failover_association = %q
}
`, failoverAssociation)
}

func testAccRangeFingerprintFilterRules(fingerprintFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_fingerprint_filter_rules" {
    fingerprint_filter_rules = %q
}
`, fingerprintFilterRules)
}

func testAccRangeHighWaterMark(highWaterMark string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_high_water_mark" {
    high_water_mark = %q
}
`, highWaterMark)
}

func testAccRangeHighWaterMarkReset(highWaterMarkReset string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_high_water_mark_reset" {
    high_water_mark_reset = %q
}
`, highWaterMarkReset)
}

func testAccRangeIgnoreDhcpOptionListRequest(ignoreDhcpOptionListRequest string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_ignore_dhcp_option_list_request" {
    ignore_dhcp_option_list_request = %q
}
`, ignoreDhcpOptionListRequest)
}

func testAccRangeIgnoreId(ignoreId string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_ignore_id" {
    ignore_id = %q
}
`, ignoreId)
}

func testAccRangeIgnoreMacAddresses(ignoreMacAddresses string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_ignore_mac_addresses" {
    ignore_mac_addresses = %q
}
`, ignoreMacAddresses)
}

func testAccRangeKnownClients(knownClients string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_known_clients" {
    known_clients = %q
}
`, knownClients)
}

func testAccRangeLeaseScavengeTime(leaseScavengeTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_lease_scavenge_time" {
    lease_scavenge_time = %q
}
`, leaseScavengeTime)
}

func testAccRangeLogicFilterRules(logicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_logic_filter_rules" {
    logic_filter_rules = %q
}
`, logicFilterRules)
}

func testAccRangeLowWaterMark(lowWaterMark string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_low_water_mark" {
    low_water_mark = %q
}
`, lowWaterMark)
}

func testAccRangeLowWaterMarkReset(lowWaterMarkReset string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_low_water_mark_reset" {
    low_water_mark_reset = %q
}
`, lowWaterMarkReset)
}

func testAccRangeMacFilterRules(macFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_mac_filter_rules" {
    mac_filter_rules = %q
}
`, macFilterRules)
}

func testAccRangeMember(member string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_member" {
    member = %q
}
`, member)
}

func testAccRangeMsAdUserData(msAdUserData string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_ms_ad_user_data" {
    ms_ad_user_data = %q
}
`, msAdUserData)
}

func testAccRangeMsOptions(msOptions string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_ms_options" {
    ms_options = %q
}
`, msOptions)
}

func testAccRangeMsServer(msServer string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_ms_server" {
    ms_server = %q
}
`, msServer)
}

func testAccRangeNacFilterRules(nacFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_nac_filter_rules" {
    nac_filter_rules = %q
}
`, nacFilterRules)
}

func testAccRangeName(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_name" {
    name = %q
}
`, name)
}

func testAccRangeNetwork(network string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_network" {
    network = %q
}
`, network)
}

func testAccRangeNetworkView(networkView string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_network_view" {
    network_view = %q
}
`, networkView)
}

func testAccRangeNextserver(nextserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_nextserver" {
    nextserver = %q
}
`, nextserver)
}

func testAccRangeOptionFilterRules(optionFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_option_filter_rules" {
    option_filter_rules = %q
}
`, optionFilterRules)
}

func testAccRangeOptions(options string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_options" {
    options = %q
}
`, options)
}

func testAccRangePortControlBlackoutSetting(portControlBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_port_control_blackout_setting" {
    port_control_blackout_setting = %q
}
`, portControlBlackoutSetting)
}

func testAccRangePxeLeaseTime(pxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_pxe_lease_time" {
    pxe_lease_time = %q
}
`, pxeLeaseTime)
}

func testAccRangeRecycleLeases(recycleLeases string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_recycle_leases" {
    recycle_leases = %q
}
`, recycleLeases)
}

func testAccRangeRelayAgentFilterRules(relayAgentFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_relay_agent_filter_rules" {
    relay_agent_filter_rules = %q
}
`, relayAgentFilterRules)
}

func testAccRangeRestartIfNeeded(restartIfNeeded string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_restart_if_needed" {
    restart_if_needed = %q
}
`, restartIfNeeded)
}

func testAccRangeSamePortControlDiscoveryBlackout(samePortControlDiscoveryBlackout string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_same_port_control_discovery_blackout" {
    same_port_control_discovery_blackout = %q
}
`, samePortControlDiscoveryBlackout)
}

func testAccRangeServerAssociationType(serverAssociationType string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_server_association_type" {
    server_association_type = %q
}
`, serverAssociationType)
}

func testAccRangeSplitMember(splitMember string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_split_member" {
    split_member = %q
}
`, splitMember)
}

func testAccRangeSplitScopeExclusionPercent(splitScopeExclusionPercent string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_split_scope_exclusion_percent" {
    split_scope_exclusion_percent = %q
}
`, splitScopeExclusionPercent)
}

func testAccRangeStartAddr(startAddr string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_start_addr" {
    start_addr = %q
}
`, startAddr)
}

func testAccRangeSubscribeSettings(subscribeSettings string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_subscribe_settings" {
    subscribe_settings = %q
}
`, subscribeSettings)
}

func testAccRangeTemplate(template string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_template" {
    template = %q
}
`, template)
}

func testAccRangeUnknownClients(unknownClients string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_unknown_clients" {
    unknown_clients = %q
}
`, unknownClients)
}

func testAccRangeUpdateDnsOnLeaseRenewal(updateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_update_dns_on_lease_renewal" {
    update_dns_on_lease_renewal = %q
}
`, updateDnsOnLeaseRenewal)
}

func testAccRangeUseBlackoutSetting(useBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_blackout_setting" {
    use_blackout_setting = %q
}
`, useBlackoutSetting)
}

func testAccRangeUseBootfile(useBootfile string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_bootfile" {
    use_bootfile = %q
}
`, useBootfile)
}

func testAccRangeUseBootserver(useBootserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_bootserver" {
    use_bootserver = %q
}
`, useBootserver)
}

func testAccRangeUseDdnsDomainname(useDdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_ddns_domainname" {
    use_ddns_domainname = %q
}
`, useDdnsDomainname)
}

func testAccRangeUseDdnsGenerateHostname(useDdnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_ddns_generate_hostname" {
    use_ddns_generate_hostname = %q
}
`, useDdnsGenerateHostname)
}

func testAccRangeUseDenyBootp(useDenyBootp string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_deny_bootp" {
    use_deny_bootp = %q
}
`, useDenyBootp)
}

func testAccRangeUseDiscoveryBasicPollingSettings(useDiscoveryBasicPollingSettings string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_discovery_basic_polling_settings" {
    use_discovery_basic_polling_settings = %q
}
`, useDiscoveryBasicPollingSettings)
}

func testAccRangeUseEmailList(useEmailList string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_email_list" {
    use_email_list = %q
}
`, useEmailList)
}

func testAccRangeUseEnableDdns(useEnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_enable_ddns" {
    use_enable_ddns = %q
}
`, useEnableDdns)
}

func testAccRangeUseEnableDhcpThresholds(useEnableDhcpThresholds string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_enable_dhcp_thresholds" {
    use_enable_dhcp_thresholds = %q
}
`, useEnableDhcpThresholds)
}

func testAccRangeUseEnableDiscovery(useEnableDiscovery string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_enable_discovery" {
    use_enable_discovery = %q
}
`, useEnableDiscovery)
}

func testAccRangeUseEnableIfmapPublishing(useEnableIfmapPublishing string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_enable_ifmap_publishing" {
    use_enable_ifmap_publishing = %q
}
`, useEnableIfmapPublishing)
}

func testAccRangeUseIgnoreDhcpOptionListRequest(useIgnoreDhcpOptionListRequest string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_ignore_dhcp_option_list_request" {
    use_ignore_dhcp_option_list_request = %q
}
`, useIgnoreDhcpOptionListRequest)
}

func testAccRangeUseIgnoreId(useIgnoreId string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_ignore_id" {
    use_ignore_id = %q
}
`, useIgnoreId)
}

func testAccRangeUseKnownClients(useKnownClients string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_known_clients" {
    use_known_clients = %q
}
`, useKnownClients)
}

func testAccRangeUseLeaseScavengeTime(useLeaseScavengeTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_lease_scavenge_time" {
    use_lease_scavenge_time = %q
}
`, useLeaseScavengeTime)
}

func testAccRangeUseLogicFilterRules(useLogicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_logic_filter_rules" {
    use_logic_filter_rules = %q
}
`, useLogicFilterRules)
}

func testAccRangeUseMsOptions(useMsOptions string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_ms_options" {
    use_ms_options = %q
}
`, useMsOptions)
}

func testAccRangeUseNextserver(useNextserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_nextserver" {
    use_nextserver = %q
}
`, useNextserver)
}

func testAccRangeUseOptions(useOptions string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_options" {
    use_options = %q
}
`, useOptions)
}

func testAccRangeUsePxeLeaseTime(usePxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_pxe_lease_time" {
    use_pxe_lease_time = %q
}
`, usePxeLeaseTime)
}

func testAccRangeUseRecycleLeases(useRecycleLeases string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_recycle_leases" {
    use_recycle_leases = %q
}
`, useRecycleLeases)
}

func testAccRangeUseSubscribeSettings(useSubscribeSettings string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_subscribe_settings" {
    use_subscribe_settings = %q
}
`, useSubscribeSettings)
}

func testAccRangeUseUnknownClients(useUnknownClients string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_unknown_clients" {
    use_unknown_clients = %q
}
`, useUnknownClients)
}

func testAccRangeUseUpdateDnsOnLeaseRenewal(useUpdateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_update_dns_on_lease_renewal" {
    use_update_dns_on_lease_renewal = %q
}
`, useUpdateDnsOnLeaseRenewal)
}
