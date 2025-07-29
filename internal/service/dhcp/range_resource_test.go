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

var readableAttributesForRange = "always_update_dns,bootfile,bootserver,cloud_info,comment,ddns_domainname,ddns_generate_hostname,deny_all_clients,deny_bootp,dhcp_utilization,dhcp_utilization_status,disable,discover_now_status,discovery_basic_poll_settings,discovery_blackout_setting,discovery_member,dynamic_hosts,email_list,enable_ddns,enable_dhcp_thresholds,enable_discovery,enable_email_warnings,enable_ifmap_publishing,enable_pxe_lease_time,enable_snmp_warnings,end_addr,endpoint_sources,exclude,extattrs,failover_association,fingerprint_filter_rules,high_water_mark,high_water_mark_reset,ignore_dhcp_option_list_request,ignore_id,ignore_mac_addresses,is_split_scope,known_clients,lease_scavenge_time,logic_filter_rules,low_water_mark,low_water_mark_reset,mac_filter_rules,member,ms_ad_user_data,ms_options,ms_server,nac_filter_rules,name,network,network_view,nextserver,option_filter_rules,options,port_control_blackout_setting,pxe_lease_time,recycle_leases,relay_agent_filter_rules,same_port_control_discovery_blackout,server_association_type,start_addr,static_hosts,subscribe_settings,total_hosts,unknown_clients,update_dns_on_lease_renewal,use_blackout_setting,use_bootfile,use_bootserver,use_ddns_domainname,use_ddns_generate_hostname,use_deny_bootp,use_discovery_basic_polling_settings,use_email_list,use_enable_ddns,use_enable_dhcp_thresholds,use_enable_discovery,use_enable_ifmap_publishing,use_ignore_dhcp_option_list_request,use_ignore_id,use_known_clients,use_lease_scavenge_time,use_logic_filter_rules,use_ms_options,use_nextserver,use_options,use_pxe_lease_time,use_recycle_leases,use_subscribe_settings,use_unknown_clients,use_update_dns_on_lease_renewal"

func TestAccRangeResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_range.test"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeBasicConfig(startAddr, endAddr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_addr", startAddr),
					resource.TestCheckResourceAttr(resourceName, "end_addr", endAddr),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "always_update_dns", "false"),
					resource.TestCheckResourceAttr(resourceName, "bootfile", ""),
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", ""),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_discovery", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_ifmap_publishing", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_snmp_warnings", "false"),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "95"),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "85"),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_split_scope", "false"),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "-1"),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "0"),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "10"),
					resource.TestCheckResourceAttr(resourceName, "network_view", "default"),
					resource.TestCheckResourceAttr(resourceName, "nextserver", ""),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ifmap_publishing", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_known_clients", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ms_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_unknown_clients", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_email_list", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_subscribe_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_id", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_range.test"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRangeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRangeBasicConfig(startAddr, endAddr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					testAccCheckRangeDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRangeResource_AlwaysUpdateDns(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_always_update_dns"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	alwaysUpdateDns := false
	alwaysUpdateDnsUpdate := true

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeAlwaysUpdateDns(startAddr, endAddr, alwaysUpdateDns),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "always_update_dns", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeAlwaysUpdateDns(startAddr, endAddr, alwaysUpdateDnsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "always_update_dns", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Bootfile(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_bootfile"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	bootFile := "boot.ini"
	boolFileUpdate := "bootfile_update.ini"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeBootfile(startAddr, endAddr, bootFile),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", bootFile),
				),
			},
			// Update and Read
			{
				Config: testAccRangeBootfile(startAddr, endAddr, boolFileUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", boolFileUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Bootserver(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_bootserver"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	bootServer := "bootServer"
	bootServerUpdate := "bootServerUpdate"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeBootserver(startAddr, endAddr, bootServer),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", bootServer),
				),
			},
			// Update and Read
			{
				Config: testAccRangeBootserver(startAddr, endAddr, bootServerUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", bootServerUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_CloudInfo(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_cloud_info"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeCloudInfo(startAddr, endAddr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.authority_type", "GM"),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.delegated_scope", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.mgmt_platform", ""),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.owned_by_adaptor", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_comment"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	comment := "network range"
	commentUpdate := "network range updated"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeComment(startAddr, endAddr, comment),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment),
				),
			},
			// Update and Read
			{
				Config: testAccRangeComment(startAddr, endAddr, commentUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", commentUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_DdnsDomainname(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_ddns_domainname"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	ddnsDomainName := "yourdomain.com"
	ddnsDomainNameUpdate := "yourdomainupdate.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeDdnsDomainname(startAddr, endAddr, ddnsDomainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", ddnsDomainName),
				),
			},
			// Update and Read
			{
				Config: testAccRangeDdnsDomainname(startAddr, endAddr, ddnsDomainNameUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", ddnsDomainNameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_DdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_ddns_generate_hostname"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	ddnsGenerateHostName := true
	ddnsGenerateHostNameUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeDdnsGenerateHostname(startAddr, endAddr, ddnsGenerateHostName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeDdnsGenerateHostname(startAddr, endAddr, ddnsGenerateHostNameUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_DenyAllClients(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_deny_all_clients"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	denyAllClients := true
	denyAllClientsUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeDenyAllClients(startAddr, endAddr, denyAllClients),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_all_clients", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeDenyAllClients(startAddr, endAddr, denyAllClientsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_all_clients", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_DenyBootp(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_deny_bootp"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	denyBootp := true
	denyBootpUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeDenyBootp(startAddr, endAddr, denyBootp),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeDenyBootp(startAddr, endAddr, denyBootpUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Disable(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_disable"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	disable := true
	disableUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeDisable(startAddr, endAddr, disable),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeDisable(startAddr, endAddr, disableUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_DiscoveryBasicPollSettings(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_discovery_basic_poll_settings"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeDiscoveryBasicPollSettings(startAddr, endAddr, true, false, false, false, false, false, false, false, "PERIODIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.auto_arp_refresh_before_switch_port_polling", "true"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.cli_collection", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.complete_ping_sweep", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.device_profile", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.netbios_scanning", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.port_scanning", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.smart_subnet_ping_sweep", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.snmp_collection", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.switch_port_data_collection_polling", "PERIODIC"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeDiscoveryBasicPollSettings(startAddr, endAddr, true, true, false, true, false, true, false, false, "SCHEDULED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.auto_arp_refresh_before_switch_port_polling", "true"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.cli_collection", "true"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.complete_ping_sweep", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.device_profile", "true"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.netbios_scanning", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.port_scanning", "true"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.smart_subnet_ping_sweep", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.snmp_collection", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.switch_port_data_collection_polling", "SCHEDULED"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_DiscoveryBlackoutSetting(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_discovery_blackout_setting"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeDiscoveryBlackoutSetting(startAddr, endAddr, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_blackout_setting.enable_blackout", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// TODO
func TestAccRangeResource_DiscoveryMember(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_discovery_member"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	discoveryMember := "infoblox.172_28_83_235"
	discoveryMemberUpdate := "infoblox.172_28_83_209"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeDiscoveryMember(startAddr, endAddr, discoveryMember),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_member", discoveryMember),
				),
			},
			// Update and Read
			{
				Config: testAccRangeDiscoveryMember(startAddr, endAddr, discoveryMemberUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_member", discoveryMemberUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_EmailList(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_email_list"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	emailList := []string{"example@infoblox.com"}
	emailListUpdate := []string{"example2@example.com"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeEmailList(startAddr, endAddr, emailList),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_list.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "email_list.0", "example@infoblox.com"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeEmailList(startAddr, endAddr, emailListUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_list.0", "example2@example.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_EnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_enable_ddns"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	enableDDNS := true
	enableDDNSUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeEnableDdns(startAddr, endAddr, enableDDNS),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeEnableDdns(startAddr, endAddr, enableDDNSUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_EnableDhcpThresholds(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_enable_dhcp_thresholds"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	enableDhcpThreasholds := true
	enableDhcpThreasholdsUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeEnableDhcpThresholds(startAddr, endAddr, enableDhcpThreasholds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_dhcp_thresholds", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeEnableDhcpThresholds(startAddr, endAddr, enableDhcpThreasholdsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_dhcp_thresholds", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_EnableEmailWarnings(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_enable_email_warnings"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	enableEmailWarnings := true
	enableEmailWarningsUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeEnableEmailWarnings(startAddr, endAddr, enableEmailWarnings),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_email_warnings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeEnableEmailWarnings(startAddr, endAddr, enableEmailWarningsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_email_warnings", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_EnableIfmapPublishing(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_enable_ifmap_publishing"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	enableIfmapPublishing := true
	enableIfmapPublishingUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeEnableIfmapPublishing(startAddr, endAddr, enableIfmapPublishing),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ifmap_publishing", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeEnableIfmapPublishing(startAddr, endAddr, enableIfmapPublishingUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ifmap_publishing", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_EnablePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_enable_pxe_lease_time"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	enablePxeLeaseTime := true
	enablePxeLeaseTimeUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeEnablePxeLeaseTime(startAddr, endAddr, enablePxeLeaseTime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeEnablePxeLeaseTime(startAddr, endAddr, enablePxeLeaseTimeUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_EnableSnmpWarnings(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_enable_snmp_warnings"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	enableSnmpWarnings := true
	enableSnmpWarningsUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeEnableSnmpWarnings(startAddr, endAddr, enableSnmpWarnings),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_snmp_warnings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeEnableSnmpWarnings(startAddr, endAddr, enableSnmpWarningsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_snmp_warnings", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_EndAddr(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_end_addr"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	updateEndAddr := "10.0.0.30"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeEndAddr(startAddr, endAddr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "end_addr", endAddr),
				),
			},
			// Update and Read
			{
				Config: testAccRangeEndAddr(startAddr, updateEndAddr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "end_addr", updateEndAddr),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Exclude(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_exclude"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	exclude := []map[string]any{
		{
			"start_address": "10.0.0.13",
			"end_address":   "10.0.0.15",
		},
	}
	excludeUpdate := []map[string]any{
		{
			"start_address": "10.0.0.16",
			"end_address":   "10.0.0.18",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeExclude(startAddr, endAddr, exclude),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "exclude.0.start_address", "10.0.0.13"),
					resource.TestCheckResourceAttr(resourceName, "exclude.0.end_address", "10.0.0.15"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeExclude(startAddr, endAddr, excludeUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "exclude.0.start_address", "10.0.0.16"),
					resource.TestCheckResourceAttr(resourceName, "exclude.0.end_address", "10.0.0.18"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_extattrs"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeExtAttrs(startAddr, endAddr, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRangeExtAttrs(startAddr, endAddr, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// TODO
func TestAccRangeResource_EnableDiscovery(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_enable_discovery"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	enableDiscovery := true
	enableDiscoveryUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeEnableDiscovery(startAddr, endAddr, enableDiscovery),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_discovery", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeEnableDiscovery(startAddr, endAddr, enableDiscoveryUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_discovery", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// TODO
func TestAccRangeResource_EnableImmediateDiscovery(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_enable_immediate_discovery"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	enableImmediateDiscovery := true
	enableImmediateDiscoveryUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeEnableImmediateDiscovery(startAddr, endAddr, enableImmediateDiscovery),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_immediate_discovery", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeEnableImmediateDiscovery(startAddr, endAddr, enableImmediateDiscoveryUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_immediate_discovery", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_FailoverAssociation(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_failover_association"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	failoverAssociation := "failover_association"
	failoverAssociationUpdate := "failover_association_1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeFailoverAssociation(startAddr, endAddr, failoverAssociation),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "failover_association", failoverAssociation),
				),
			},
			// Update and Read
			{
				Config: testAccRangeFailoverAssociation(startAddr, endAddr, failoverAssociationUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "failover_association", failoverAssociationUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_FingerprintFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_fingerprint_filter_rules"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	fingerprintFilterRules := []map[string]any{
		{
			"filter":     "range_network_filter",
			"permission": "Allow",
		},
	}
	fingerprintFilterRulesUpdate := []map[string]any{
		{
			"filter":     "range_network_filter1",
			"permission": "Allow",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeFingerprintFilterRules(startAddr, endAddr, fingerprintFilterRules),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fingerprint_filter_rules.0.filter", "range_network_filter"),
					resource.TestCheckResourceAttr(resourceName, "fingerprint_filter_rules.0.permission", "Allow"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeFingerprintFilterRules(startAddr, endAddr, fingerprintFilterRulesUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fingerprint_filter_rules.0.filter", "range_network_filter1"),
					resource.TestCheckResourceAttr(resourceName, "fingerprint_filter_rules.0.permission", "Allow"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_HighWaterMark(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_high_water_mark"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	highWaterMark := 23
	highWaterMarkUpdate := 42

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeHighWaterMark(startAddr, endAddr, highWaterMark),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "23"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeHighWaterMark(startAddr, endAddr, highWaterMarkUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "42"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_HighWaterMarkReset(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_high_water_mark_reset"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	highWaterMarkReset := 23
	highWaterMarkResetUpdate := 42

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeHighWaterMarkReset(startAddr, endAddr, highWaterMarkReset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "23"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeHighWaterMarkReset(startAddr, endAddr, highWaterMarkResetUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "42"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_IgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_ignore_dhcp_option_list_request"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	ignoreDhcpOptionListRequest := true
	ignoreDhcpOptionListRequestUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeIgnoreDhcpOptionListRequest(startAddr, endAddr, ignoreDhcpOptionListRequest),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeIgnoreDhcpOptionListRequest(startAddr, endAddr, ignoreDhcpOptionListRequestUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_IgnoreId(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_ignore_id"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	ignoreId := "CLIENT"
	ignoreIdUpdate := "MACADDR"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeIgnoreId(startAddr, endAddr, ignoreId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_id", "CLIENT"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeIgnoreId(startAddr, endAddr, ignoreIdUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_id", "MACADDR"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_IgnoreMacAddresses(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_ignore_mac_addresses"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	ignoreMacAddresses := []string{"00:1a:2b:3c:4d:5e"}
	ignoreMacAddressesUpdate := []string{"00:1a:2b:33:4d:52"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeIgnoreMacAddresses(startAddr, endAddr, ignoreMacAddresses),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_mac_addresses.0", "00:1a:2b:3c:4d:5e"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeIgnoreMacAddresses(startAddr, endAddr, ignoreMacAddressesUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_mac_addresses.0", "00:1a:2b:33:4d:52"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_KnownClients(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_known_clients"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	knownClients := "Deny"
	knownClientsUpdate := "Allow"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeKnownClients(startAddr, endAddr, knownClients),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "known_clients", "Deny"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeKnownClients(startAddr, endAddr, knownClientsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "known_clients", "Allow"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_LeaseScavengeTime(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_lease_scavenge_time"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	leaseScavengeTime := 86420
	leaseScavengeTimeUpdate := 86430
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeLeaseScavengeTime(startAddr, endAddr, leaseScavengeTime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "86420"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeLeaseScavengeTime(startAddr, endAddr, leaseScavengeTimeUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "86430"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_LogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_logic_filter_rules"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	logicFilterRules := []map[string]any{
		{
			"filter": "mac_filter",
			"type":   "MAC",
		},
	}
	logicFilterRulesUpdate := []map[string]any{
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
				Config: testAccRangeLogicFilterRules(startAddr, endAddr, logicFilterRules),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.filter", "mac_filter"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.type", "MAC"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeLogicFilterRules(startAddr, endAddr, logicFilterRulesUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.filter", "option_logic_filter"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.type", "Option"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_LowWaterMark(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_low_water_mark"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	lowWaterMark := 5
	lowWaterMarkUpdate := 1

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeLowWaterMark(startAddr, endAddr, lowWaterMark),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "5"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeLowWaterMark(startAddr, endAddr, lowWaterMarkUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_LowWaterMarkReset(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_low_water_mark_reset"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	lowWaterMarkReset := 5
	lowWaterMarkResetUpdate := 1

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeLowWaterMarkReset(startAddr, endAddr, lowWaterMarkReset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "5"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeLowWaterMarkReset(startAddr, endAddr, lowWaterMarkResetUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_MacFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_mac_filter_rules"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	macFilterRules := []map[string]any{
		{
			"filter":     "mac_filter",
			"permission": "Allow",
		},
	}
	macFilterRulesUpdate := []map[string]any{
		{
			"filter":     "mac_logic_filter",
			"permission": "Deny",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeMacFilterRules(startAddr, endAddr, macFilterRules),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mac_filter_rules.0.filter", "mac_filter"),
					resource.TestCheckResourceAttr(resourceName, "mac_filter_rules.0.permission", "Allow"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeMacFilterRules(startAddr, endAddr, macFilterRulesUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mac_filter_rules.0.filter", "mac_logic_filter"),
					resource.TestCheckResourceAttr(resourceName, "mac_filter_rules.0.permission", "Deny"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Member(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_member"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	member := map[string]any{
		"ipv4addr": "172.28.83.235",
		"name":     "infoblox.172_28_83_235",
	}
	memberUpdate := map[string]any{
		"ipv4addr": "172.28.83.209",
		"name":     "infoblox.172_28_83_209",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeMember(startAddr, endAddr, member),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member.ipv4addr", "172.28.83.235"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeMember(startAddr, endAddr, memberUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member.ipv4addr", "172.28.83.209"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// TODO
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

// TODO
func TestAccRangeResource_MsServer(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_ms_server"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	msServerIp := "10.120.23.22"
	msServerIpUpdate := "10.120.23.23"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeMsServer(startAddr, endAddr, msServerIp),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_server.ipv4addr", "10.120.23.22"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeMsServer(startAddr, endAddr, msServerIpUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_server.ipv4addr", "10.120.23.23"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_NacFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_nac_filter_rules"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	nacFilterRules := []map[string]any{
		{
			"filter":     "nac_filter",
			"permission": "Allow",
		},
	}
	nacFilterRulesUpdate := []map[string]any{
		{
			"filter":     "nac_filter_rule",
			"permission": "Deny",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeNacFilterRules(startAddr, endAddr, nacFilterRules),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nac_filter_rules.0.filter", "nac_filter"),
					resource.TestCheckResourceAttr(resourceName, "nac_filter_rules.0.permission", "Allow"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeNacFilterRules(startAddr, endAddr, nacFilterRulesUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nac_filter_rules.0.filter", "nac_filter_rule"),
					resource.TestCheckResourceAttr(resourceName, "nac_filter_rules.0.permission", "Deny"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_name"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	name := "range 1"
	nameUpdate := "range 2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeName(startAddr, endAddr, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccRangeName(startAddr, endAddr, nameUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Network(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_network"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	network := "10.0.0.0/24"
	networkUpdate := "20.0.0.0/24"
	startAddrUpdate := "20.0.0.20"
	endAddrUpdate := "20.0.0.30"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeNetwork(startAddr, endAddr, network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccRangeNetwork(startAddrUpdate, endAddrUpdate, networkUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", networkUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_NetworkView(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_network_view"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	networkView := "custom_view"
	//networkViewUpdate := "custom_view"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeNetworkView(startAddr, endAddr, networkView),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", networkView),
				),
			},
			//network view is not updatable for the network range
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Nextserver(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_nextserver"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	nextServer := "next_server.com"
	nextServerUpdate := "next_server_update.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeNextserver(startAddr, endAddr, nextServer),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", nextServer),
				),
			},
			// Update and Read
			{
				Config: testAccRangeNextserver(startAddr, endAddr, nextServerUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", nextServerUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_OptionFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_option_filter_rules"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	optionFilterRules := []map[string]any{
		{
			"filter":     "option_filter",
			"permission": "Allow",
		},
	}
	optionFilterRulesUpdate := []map[string]any{
		{
			"filter":     "option_logic_filter",
			"permission": "Deny",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeOptionFilterRules(startAddr, endAddr, optionFilterRules),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_filter_rules.0.filter", "option_filter"),
					resource.TestCheckResourceAttr(resourceName, "option_filter_rules.0.permission", "Allow"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeOptionFilterRules(startAddr, endAddr, optionFilterRulesUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_filter_rules.0.filter", "option_logic_filter"),
					resource.TestCheckResourceAttr(resourceName, "option_filter_rules.0.permission", "Deny"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_Options(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_options"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeOptions(startAddr, endAddr, "dhcp-lease-time", "51", "6739", "DHCP", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr(resourceName, "options.0.num", "51"),
					resource.TestCheckResourceAttr(resourceName, "options.0.vendor_class", "DHCP"),
					resource.TestCheckResourceAttr(resourceName, "options.0.use_option", "true"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "6739"),
				)},
			// Update and Read
			{
				Config: testAccRangeOptions(startAddr, endAddr, "dhcp-lease-time", "51", "7300", "DHCP", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr(resourceName, "options.0.num", "51"),
					resource.TestCheckResourceAttr(resourceName, "options.0.vendor_class", "DHCP"),
					resource.TestCheckResourceAttr(resourceName, "options.0.use_option", "true"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "7300"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_PortControlBlackoutSetting(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_port_control_blackout_setting"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"

	portControlBlackoutSetting := map[string]any{
		"enable_blackout": false,
	}

	portControlBlackoutSettingUpdate := map[string]any{
		"enable_blackout": false,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangePortControlBlackoutSetting(startAddr, endAddr, portControlBlackoutSetting),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port_control_blackout_setting.enable_blackout", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRangePortControlBlackoutSetting(startAddr, endAddr, portControlBlackoutSettingUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port_control_blackout_setting.enable_blackout", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_PxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_pxe_lease_time"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	pxeLeaseTime := "3400"
	pxeLeaseTimeUpdate := "3600"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangePxeLeaseTime(startAddr, endAddr, pxeLeaseTime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "3400"),
				),
			},
			// Update and Read
			{
				Config: testAccRangePxeLeaseTime(startAddr, endAddr, pxeLeaseTimeUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "3600"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_RecycleLeases(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_recycle_leases"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	recycleLeases := true
	recycleLeasesUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeRecycleLeases(startAddr, endAddr, recycleLeases),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeRecycleLeases(startAddr, endAddr, recycleLeasesUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_RelayAgentFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_relay_agent_filter_rules"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	relayAgentFilterRules := []map[string]any{
		{
			"filter":     "relay_agent_filter",
			"permission": "Allow",
		},
	}
	relayAgentFilterRulesUpdate := []map[string]any{
		{
			"filter":     "relay_agent_logic_filter",
			"permission": "Deny",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeRelayAgentFilterRules(startAddr, endAddr, relayAgentFilterRules),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "relay_agent_filter_rules.0.filter", "relay_agent_filter"),
					resource.TestCheckResourceAttr(resourceName, "relay_agent_filter_rules.0.permission", "Allow"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeRelayAgentFilterRules(startAddr, endAddr, relayAgentFilterRulesUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "relay_agent_filter_rules.0.filter", "relay_agent_logic_filter"),
					resource.TestCheckResourceAttr(resourceName, "relay_agent_filter_rules.0.permission", "Deny"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_SamePortControlDiscoveryBlackout(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_same_port_control_discovery_blackout"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	samePortControlDiscoveryBlackout := false
	samePortControlDiscoveryBlackoutUpdate := true

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeSamePortControlDiscoveryBlackout(startAddr, endAddr, samePortControlDiscoveryBlackout),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeSamePortControlDiscoveryBlackout(startAddr, endAddr, samePortControlDiscoveryBlackoutUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_ServerAssociationType(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_server_association_type"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	serverAssociationType := "FAILOVER"
	failoverAssociation := "failover_association"
	serverAssociationTypeUpdate := "MEMBER"
	member := "infoblox.172_28_83_209"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeServerAssociationType(startAddr, endAddr, serverAssociationType, failoverAssociation, member),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "server_association_type", "FAILOVER"),
					resource.TestCheckResourceAttr(resourceName, "failover_association", failoverAssociation),
				),
			},
			// Update and Read
			{
				Config: testAccRangeServerAssociationType(startAddr, endAddr, serverAssociationTypeUpdate, failoverAssociation, member),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "server_association_type", "MEMBER"),
					resource.TestCheckResourceAttr(resourceName, "member.ipv4addr", "172.28.83.209"),
					resource.TestCheckResourceAttr(resourceName, "member.name", "infoblox.172_28_83_209"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// split member is not updatable
func TestAccRangeResource_SplitMember(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_split_member"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	splitMemberipv4Addr := "10.120.23.23"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeSplitMember(startAddr, endAddr, splitMemberipv4Addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "split_member", "10.120.23.22"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_StartAddr(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_start_addr"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	startAddrUpdate := "10.0.0.14"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeStartAddr(startAddr, endAddr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_addr", "10.0.0.10"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeStartAddr(startAddrUpdate, endAddr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_addr", "10.0.0.14"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_SubscribeSettings(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_subscribe_settings"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	enabledAttribute := "DOMAINNAME"
	enabledAttributeUpdate := "ENDPOINT_PROFILE"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeSubscribeSettings(startAddr, endAddr, enabledAttribute),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscribe_settings.enabled_attributes.0", enabledAttribute),
				),
			},
			// Update and Read
			{
				Config: testAccRangeSubscribeSettings(startAddr, endAddr, enabledAttributeUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscribe_settings.enabled_attributes.0", enabledAttributeUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UnknownClients(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_unknown_clients"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	unknownClients := "Allow"
	unknownClientsUpdate := "Deny"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUnknownClients(startAddr, endAddr, unknownClients),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "unknown_clients", "Allow"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUnknownClients(startAddr, endAddr, unknownClientsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "unknown_clients", "Deny"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_update_dns_on_lease_renewal"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	updateDnsOnLeaseRenewal := true
	updateDnsOnLeaseRenewalUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUpdateDnsOnLeaseRenewal(startAddr, endAddr, updateDnsOnLeaseRenewal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUpdateDnsOnLeaseRenewal(startAddr, endAddr, updateDnsOnLeaseRenewalUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseBlackoutSetting(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_blackout_setting"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useBlackoutSetting := true
	useBlackoutSettingUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseBlackoutSetting(startAddr, endAddr, useBlackoutSetting),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseBlackoutSetting(startAddr, endAddr, useBlackoutSettingUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseBootfile(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_bootfile"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useBootfile := false
	bootfile := "bootfile.com"
	useBootfileUpdate := true

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseBootfile(startAddr, endAddr, bootfile, useBootfile),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseBootfile(startAddr, endAddr, bootfile, useBootfileUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseBootserver(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_bootserver"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useBootServer := false
	bootServer := "bootfile.com"
	useBootServerUpdate := true

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseBootserver(startAddr, endAddr, bootServer, useBootServer),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseBootserver(startAddr, endAddr, bootServer, useBootServerUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseDdnsDomainname(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_ddns_domainname"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	ddnsDomainName := "yourdomain.com"
	useddnsDomainName := true
	useddnsDomainNameUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseDdnsDomainname(startAddr, endAddr, ddnsDomainName, useddnsDomainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseDdnsDomainname(startAddr, endAddr, ddnsDomainName, useddnsDomainNameUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// need to check this because if we pass the field as false
func TestAccRangeResource_UseDdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_ddns_generate_hostname"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useDdnsGenerateHostname := true
	useDdnsGenerateHostnameUpdate := false
	ddnsGeneratedHostname := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseDdnsGenerateHostname(startAddr, endAddr, useDdnsGenerateHostname, ddnsGeneratedHostname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseDdnsGenerateHostname(startAddr, endAddr, useDdnsGenerateHostnameUpdate, ddnsGeneratedHostname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseDenyBootp(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_deny_bootp"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useDenyBootp := true
	useDenyBootpUpdate := false
	denyBootp := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseDenyBootp(startAddr, endAddr, useDenyBootp, denyBootp),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseDenyBootp(startAddr, endAddr, useDenyBootpUpdate, denyBootp),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseDiscoveryBasicPollingSettings(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_discovery_basic_polling_settings"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useDiscoveryBasicPollingSettings := true
	useDiscoveryBasicPollingSettingsUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseDiscoveryBasicPollingSettings(startAddr, endAddr, useDiscoveryBasicPollingSettings),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseDiscoveryBasicPollingSettings(startAddr, endAddr, useDiscoveryBasicPollingSettingsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseEmailList(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_email_list"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useEmailList := true
	useEmailListUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseEmailList(startAddr, endAddr, useEmailList),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_email_list", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseEmailList(startAddr, endAddr, useEmailListUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_email_list", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseEnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_enable_ddns"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useEnableDdns := true
	useEnableDdnsUpdate := false
	enableDdns := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseEnableDdns(startAddr, endAddr, useEnableDdns, enableDdns),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseEnableDdns(startAddr, endAddr, useEnableDdnsUpdate, enableDdns),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseEnableDhcpThresholds(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_enable_dhcp_thresholds"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useEnabledhcpThreasholds := true
	useEnabledhcpThreasholdsUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseEnableDhcpThresholds(startAddr, endAddr, useEnabledhcpThreasholds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseEnableDhcpThresholds(startAddr, endAddr, useEnabledhcpThreasholdsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseEnableDiscovery(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_enable_discovery"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useEnableDiscovery := true
	useEnableDiscoveryUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseEnableDiscovery(startAddr, endAddr, useEnableDiscovery),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseEnableDiscovery(startAddr, endAddr, useEnableDiscoveryUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseEnableIfmapPublishing(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_enable_ifmap_publishing"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useEnableIfmapPublishing := true
	useEnableIfmapPublishingUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseEnableIfmapPublishing(startAddr, endAddr, useEnableIfmapPublishing),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ifmap_publishing", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseEnableIfmapPublishing(startAddr, endAddr, useEnableIfmapPublishingUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ifmap_publishing", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseIgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_ignore_dhcp_option_list_request"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useIgnoreDhcpOptionListRequest := true
	useIgnoreDhcpOptionListRequestUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseIgnoreDhcpOptionListRequest(startAddr, endAddr, useIgnoreDhcpOptionListRequest),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseIgnoreDhcpOptionListRequest(startAddr, endAddr, useIgnoreDhcpOptionListRequestUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseIgnoreId(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_ignore_id"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useIgnoreId := true
	useIgnoreIdUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseIgnoreId(startAddr, endAddr, useIgnoreId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_id", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseIgnoreId(startAddr, endAddr, useIgnoreIdUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_id", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseKnownClients(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_known_clients"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useKnownClients := true
	useKnownClientsUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseKnownClients(startAddr, endAddr, useKnownClients),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_known_clients", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseKnownClients(startAddr, endAddr, useKnownClientsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_known_clients", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseLeaseScavengeTime(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_lease_scavenge_time"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useLeaseScavengeTime := true
	useLeaseScavengeTimeUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseLeaseScavengeTime(startAddr, endAddr, useLeaseScavengeTime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseLeaseScavengeTime(startAddr, endAddr, useLeaseScavengeTimeUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseLogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_logic_filter_rules"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useLogicFilterRules := true
	useLogicFilterRulesUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseLogicFilterRules(startAddr, endAddr, useLogicFilterRules),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseLogicFilterRules(startAddr, endAddr, useLogicFilterRulesUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseMsOptions(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_ms_options"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useMsOptions := true
	useMsOptionsUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseMsOptions(startAddr, endAddr, useMsOptions),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ms_options", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseMsOptions(startAddr, endAddr, useMsOptionsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ms_options", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseNextserver(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_nextserver"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useNextServer := true
	useNextServerUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseNextserver(startAddr, endAddr, useNextServer),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseNextserver(startAddr, endAddr, useNextServerUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseOptions(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_options"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useOptions := true
	useOptionsUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseOptions(startAddr, endAddr, useOptions),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseOptions(startAddr, endAddr, useOptionsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UsePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_pxe_lease_time"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	usePxeLeaseTime := true
	usePxeLeaseTimeUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUsePxeLeaseTime(startAddr, endAddr, usePxeLeaseTime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUsePxeLeaseTime(startAddr, endAddr, usePxeLeaseTimeUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseRecycleLeases(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_recycle_leases"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useRecycleLeases := true
	useRecycleLeasesUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseRecycleLeases(startAddr, endAddr, useRecycleLeases),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseRecycleLeases(startAddr, endAddr, useRecycleLeasesUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseSubscribeSettings(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_subscribe_settings"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useSubscribeSettings := true
	useSubscribeSettingsUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseSubscribeSettings(startAddr, endAddr, useSubscribeSettings),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_subscribe_settings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseSubscribeSettings(startAddr, endAddr, useSubscribeSettingsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_subscribe_settings", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseUnknownClients(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_unknown_clients"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useUnkownClients := true
	useUnkownClientsUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseUnknownClients(startAddr, endAddr, useUnkownClients),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_unknown_clients", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseUnknownClients(startAddr, endAddr, useUnkownClientsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_unknown_clients", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangeResource_UseUpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_dhcp_range.test_use_update_dns_on_lease_renewal"
	var v dhcp.Range
	startAddr := "10.0.0.10"
	endAddr := "10.0.0.20"
	useUpdateDnsOnLeaseRenewal := true
	useUpdateDnsOnLeaseRenewalUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangeUseUpdateDnsOnLeaseRenewal(startAddr, endAddr, useUpdateDnsOnLeaseRenewal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangeUseUpdateDnsOnLeaseRenewal(startAddr, endAddr, useUpdateDnsOnLeaseRenewalUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "false"),
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

func testAccRangeBasicConfig(startAddr, endAddr string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test" {
    start_addr = %q
    end_addr   = %q
}
`, startAddr, endAddr)
}

func testAccRangeAlwaysUpdateDns(startAddr, endAddr string, alwaysUpdateDns bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_always_update_dns" {
    start_addr = %q
    end_addr   = %q
    always_update_dns = %t
}
`, startAddr, endAddr, alwaysUpdateDns)
}

func testAccRangeBootfile(startAddr, endAddr, bootfile string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_bootfile" {
    start_addr = %q
    end_addr   = %q
    bootfile = %q
	use_bootfile = true
}
`, startAddr, endAddr, bootfile)
}

func testAccRangeBootserver(startAddr, endAddr, bootserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_bootserver" {
    start_addr = %q
    end_addr   = %q
    bootserver = %q
	use_bootserver = true
}
`, startAddr, endAddr, bootserver)
}

func testAccRangeCloudInfo(startAddr, endAddr string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_cloud_info" {
	start_addr = %q
	end_addr = %q
}
`, startAddr, endAddr)
}

func testAccRangeComment(startAddr, endAddr, comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_comment" {
	start_addr = %q
	end_addr   = %q
    comment = %q
}
`, startAddr, endAddr, comment)
}

func testAccRangeDdnsDomainname(startAddr, endAddr, ddnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_ddns_domainname" {
	start_addr = %q
	end_addr   = %q
	use_ddns_domainname = true
    ddns_domainname = %q
}
`, startAddr, endAddr, ddnsDomainname)
}

func testAccRangeDdnsGenerateHostname(startAddr, endAddr string, ddnsGenerateHostname bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_ddns_generate_hostname" {
	start_addr = %q
	end_addr   = %q
	use_ddns_generate_hostname = true
    ddns_generate_hostname = %t
}
`, startAddr, endAddr, ddnsGenerateHostname)
}

func testAccRangeDenyAllClients(startAddr, endAddr string, denyAllClients bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_deny_all_clients" {
    start_addr = %q
    end_addr   = %q
    deny_all_clients = %t
}
`, startAddr, endAddr, denyAllClients)
}

func testAccRangeDenyBootp(startAddr, endAddr string, denyBootp bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_deny_bootp" {
    start_addr = %q
    end_addr   = %q
    deny_bootp = %t
	use_deny_bootp =true
}
`, startAddr, endAddr, denyBootp)
}

func testAccRangeDisable(startAddr, endAddr string, disable bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_disable" {
    start_addr = %q
    end_addr   = %q
    disable = %t
}
`, startAddr, endAddr, disable)
}

func testAccRangeDiscoveryBasicPollSettings(startAddr, endAddr string, autoArpRefreshBeforeSwitchPortPolling, cliCollection, completePingSweep, deviceProfile, netbiosScanning, portScanning, smartSubnetPingSweep, snmpCollection bool, switchPortDataCollectionPolling string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_discovery_basic_poll_settings" {
	start_addr = %q
	end_addr = %q
	discovery_basic_poll_settings = {
		auto_arp_refresh_before_switch_port_polling = %t
		cli_collection = %t
		complete_ping_sweep = %t
		device_profile = %t
		netbios_scanning = %t
		port_scanning = %t
		smart_subnet_ping_sweep = %t
		snmp_collection = %t
		switch_port_data_collection_polling = %q
	}
		use_discovery_basic_polling_settings = true
}
`, startAddr, endAddr,
		autoArpRefreshBeforeSwitchPortPolling, cliCollection, completePingSweep, deviceProfile,
		netbiosScanning, portScanning, smartSubnetPingSweep, snmpCollection,
		switchPortDataCollectionPolling)
}

func testAccRangeDiscoveryBlackoutSetting(startAddr, endAddr string, enableBlackout bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_discovery_blackout_setting" {
	start_addr = %q
	end_addr = %q
    discovery_blackout_setting = {
		enable_blackout = %t
		}
		use_blackout_setting = true
}
`, startAddr, endAddr, enableBlackout)
}

func testAccRangeDiscoveryMember(startAddr, endAddr, discoveryMember string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_discovery_member" {
    start_addr = %q
    end_addr   = %q
    discovery_member = %q
	use_enable_discovery = true
}
`, startAddr, endAddr, discoveryMember)
}

func testAccRangeEmailList(startAddr, endAddr string, emailList []string) string {
	emailListHCL := utils.ConvertStringSliceToHCL(emailList)
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_email_list" {
	start_addr = %q
	end_addr = %q
	use_email_list = true
    email_list = %s
}
`, startAddr, endAddr, emailListHCL)
}

func testAccRangeEnableDdns(startAddr, endAddr string, enableDdns bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_enable_ddns" {
	start_addr = %q
	end_addr = %q
    enable_ddns = %t
	use_enable_ddns = true
}
`, startAddr, endAddr, enableDdns)
}

func testAccRangeEnableDhcpThresholds(startAddr, endAddr string, enableDhcpThresholds bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_enable_dhcp_thresholds" {
    start_addr = %q
    end_addr = %q
    enable_dhcp_thresholds = %t
	use_enable_dhcp_thresholds = true
}
`, startAddr, endAddr, enableDhcpThresholds)
}

func testAccRangeEnableDiscovery(startAddr, endAddr string, enableDiscovery bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_enable_discovery" {
    start_addr = %q
    end_addr = %q
    enable_discovery = %t
	use_enable_discovery = true
}
`, startAddr, endAddr, enableDiscovery)
}

func testAccRangeEnableEmailWarnings(startAddr, endAddr string, enableEmailWarnings bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_enable_email_warnings" {
	start_addr = %q
	end_addr = %q
    enable_email_warnings = %t
}
`, startAddr, endAddr, enableEmailWarnings)
}

func testAccRangeEnableIfmapPublishing(startAddr, endAddr string, enableIfmapPublishing bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_enable_ifmap_publishing" {
    start_addr = %q
    end_addr = %q
    enable_ifmap_publishing = %t
	use_enable_ifmap_publishing  = true
}
`, startAddr, endAddr, enableIfmapPublishing)
}

func testAccRangeEnableImmediateDiscovery(startAddr, endAddr string, enableImmediateDiscovery bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_enable_immediate_discovery" {
    start_addr = %q
    end_addr = %q
    enable_immediate_discovery = %t
}
`, startAddr, endAddr, enableImmediateDiscovery)
}

func testAccRangeEnablePxeLeaseTime(startAddr, endAddr string, enablePxeLeaseTime bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_enable_pxe_lease_time" {
    start_addr = %q
    end_addr = %q
    enable_pxe_lease_time = %t
	pxe_lease_time= 3600
	use_pxe_lease_time = true
}
`, startAddr, endAddr, enablePxeLeaseTime)
}

func testAccRangeEnableSnmpWarnings(startAddr, endAddr string, enableSnmpWarnings bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_enable_snmp_warnings" {
    enable_snmp_warnings = %t
    start_addr = %q
    end_addr = %q
}
`, enableSnmpWarnings, startAddr, endAddr)
}

func testAccRangeEndAddr(startAddr, endAddr string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_end_addr" {
    start_addr = %q
    end_addr   = %q
}
`, startAddr, endAddr)
}

func testAccRangeExclude(startAddr, endAddr string, exclude []map[string]any) string {
	excludeHCL := utils.ConvertSliceOfMapsToHCL(exclude)
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_exclude" {
	start_addr = %q
	end_addr = %q
    exclude = %s
}
`, startAddr, endAddr, excludeHCL)
}

func testAccRangeExtAttrs(startAddr, endAddr string, extAttrs map[string]string) string {
	extattrsStr := formatExtAttrsForHCL(extAttrs)
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_extattrs" {
	start_addr = %q
	end_addr = %q
    extattrs = %s
}
`, startAddr, endAddr, extattrsStr)
}

func testAccRangeFailoverAssociation(startAddr, endAddr, failoverAssociation string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_failover_association" {
    failover_association = %q
    start_addr = %q
    end_addr = %q
    server_association_type = "FAILOVER"
}
`, failoverAssociation, startAddr, endAddr)
}

func testAccRangeFingerprintFilterRules(startAddr, endAddr string, fingerprintFilterRules []map[string]any) string {
	fingerprintFilterRulesHCL := utils.ConvertSliceOfMapsToHCL(fingerprintFilterRules)
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_fingerprint_filter_rules" {
    start_addr = %q
    end_addr = %q
    fingerprint_filter_rules = %s
}
`, startAddr, endAddr, fingerprintFilterRulesHCL)
}

func testAccRangeHighWaterMark(startAddr, endAddr string, highWaterMark int) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_high_water_mark" {
	start_addr = %q
	end_addr = %q
    high_water_mark = %d
}
`, startAddr, endAddr, highWaterMark)
}

func testAccRangeHighWaterMarkReset(startAddr, endAddr string, highWaterMarkReset int) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_high_water_mark_reset" {
    high_water_mark_reset = %d
    start_addr = %q
    end_addr = %q
}
`, highWaterMarkReset, startAddr, endAddr)
}

func testAccRangeIgnoreDhcpOptionListRequest(startAddr, endAddr string, ignoreDhcpOptionListRequest bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_ignore_dhcp_option_list_request" {
    ignore_dhcp_option_list_request = %t
    start_addr = %q
    end_addr = %q
	use_ignore_dhcp_option_list_request = true
}
`, ignoreDhcpOptionListRequest, startAddr, endAddr)
}

func testAccRangeIgnoreId(startAddr, endAddr, ignoreId string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_ignore_id" {
	start_addr = %q
	end_addr = %q
    ignore_id = %q
	use_ignore_id = true
}
`, startAddr, endAddr, ignoreId)
}

func testAccRangeIgnoreMacAddresses(startAddr, endAddr string, ignoreMacAddresses []string) string {
	ignoreMacAddressesHCL := utils.ConvertStringSliceToHCL(ignoreMacAddresses)
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_ignore_mac_addresses" {
    start_addr = %q
    end_addr = %q
    ignore_mac_addresses = %s
}
`, startAddr, endAddr, ignoreMacAddressesHCL)
}

func testAccRangeKnownClients(startAddr, endAddr, knownClients string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_known_clients" {
    start_addr = %q
    end_addr = %q
    known_clients = %q
	use_known_clients = true
}
`, startAddr, endAddr, knownClients)
}

func testAccRangeLeaseScavengeTime(startAddr, endAddr string, leaseScavengeTime int) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_lease_scavenge_time" {
	start_addr = %q
	end_addr = %q
    lease_scavenge_time = %d
	use_lease_scavenge_time = true
}
`, startAddr, endAddr, leaseScavengeTime)
}

func testAccRangeLogicFilterRules(startAddr, endAddr string, logicFilterRules []map[string]any) string {
	logicFilterRulesHCL := utils.ConvertSliceOfMapsToHCL(logicFilterRules)
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_logic_filter_rules" {
	start_addr = %q
	end_addr = %q
    logic_filter_rules = %s
	use_logic_filter_rules = true
}
`, startAddr, endAddr, logicFilterRulesHCL)
}

func testAccRangeLowWaterMark(startAddr, endAddr string, lowWaterMark int) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_low_water_mark" {
	start_addr = %q
	end_addr = %q
    low_water_mark = %d
}
`, startAddr, endAddr, lowWaterMark)
}

func testAccRangeLowWaterMarkReset(startAddr, endAddr string, lowWaterMarkReset int) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_low_water_mark_reset" {
    low_water_mark_reset = %d
    start_addr = %q
    end_addr = %q
}
`, lowWaterMarkReset, startAddr, endAddr)
}

func testAccRangeMacFilterRules(startAddr, endAddr string, macFilterRules []map[string]any) string {
	macFilterRulesHCL := utils.ConvertSliceOfMapsToHCL(macFilterRules)
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_mac_filter_rules" {
	start_addr = %q
	end_addr = %q
    mac_filter_rules = %s
}
`, startAddr, endAddr, macFilterRulesHCL)
}

func testAccRangeMember(startAddr, endAddr string, member map[string]any) string {
	memberHCL := utils.ConvertMapToHCL(member)
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_member" {
	start_addr = %q
	end_addr = %q
    member = %s
	server_association_type = "MEMBER"
}
`, startAddr, endAddr, memberHCL)
}

func testAccRangeMsOptions(msOptions string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_ms_options" {
    ms_options = %q
}
`, msOptions)
}

func testAccRangeMsServer(startAddr, endAddr, msServer string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_ms_server" {
	start_addr = %q
	end_addr = %q
    ms_server = {
		ipv4addr = %q
	}
		server_association_type = "MS_SERVER"
}
`, startAddr, endAddr, msServer)
}

func testAccRangeNacFilterRules(startAddr, endAddr string, nacFilterRules []map[string]any) string {
	nacFilterRulesHCL := utils.ConvertSliceOfMapsToHCL(nacFilterRules)
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_nac_filter_rules" {
	start_addr = %q
	end_addr = %q
    nac_filter_rules = %s
}
`, startAddr, endAddr, nacFilterRulesHCL)
}

func testAccRangeName(startAddr, endAddr, name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_name" {
    start_addr = %q
    end_addr = %q
    name = %q
}
`, startAddr, endAddr, name)
}

func testAccRangeNetwork(startAddr, endAddr, network string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_network" {
	start_addr = %q
	end_addr = %q
    network = %q
}
`, startAddr, endAddr, network)
}

func testAccRangeNetworkView(startAddr, endAddr, networkView string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_network_view" {
	start_addr = %q
	end_addr = %q
    network_view = %q
}
`, startAddr, endAddr, networkView)
}

func testAccRangeNextserver(startAddr, endAddr, nextserver string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_nextserver" {
	start_addr = %q
	end_addr = %q
    nextserver = %q
	use_nextserver = true
}
`, startAddr, endAddr, nextserver)
}

func testAccRangeOptionFilterRules(startAddr, endAddr string, optionFilterRules []map[string]any) string {
	optionFilterRulesHCL := utils.ConvertSliceOfMapsToHCL(optionFilterRules)
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_option_filter_rules" {
	start_addr = %q
	end_addr = %q
    option_filter_rules = %s
}
`, startAddr, endAddr, optionFilterRulesHCL)
}

func testAccRangeOptions(startAddr, endAddr, name, num, value, vendorClass string, useOptions bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_options" {
	start_addr = %q
	end_addr = %q
    options = [
		{
			name = %q
			num = %q
			value = %q
			vendor_class = %q
			use_option = true
		}
    ]
		use_options = %t
}
`, startAddr, endAddr, name, num, value, vendorClass, useOptions)
}

func testAccRangePortControlBlackoutSetting(startAddr, endAddr string, portControlBlackoutSetting map[string]any) string {
	portControlBlackoutSettingStr := utils.ConvertMapToHCL(portControlBlackoutSetting)
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_port_control_blackout_setting" {
	start_addr = %q
	end_addr = %q
    port_control_blackout_setting = %s
}
`, startAddr, endAddr, portControlBlackoutSettingStr)
}

func testAccRangePxeLeaseTime(startAddr, endAddr, pxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_pxe_lease_time" {
	start_addr = %q
	end_addr = %q
	use_pxe_lease_time = true
    pxe_lease_time = %q
}
`, startAddr, endAddr, pxeLeaseTime)
}

func testAccRangeRecycleLeases(startAddr, endAddr string, recycleLeases bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_recycle_leases" {
	use_recycle_leases = true
    start_addr = %q
    end_addr = %q
    recycle_leases = %t
}
`, startAddr, endAddr, recycleLeases)
}

func testAccRangeRelayAgentFilterRules(startAddr, endAddr string, relayAgentFilterRules []map[string]any) string {
	relayAgentFilterRulesHCL := utils.ConvertSliceOfMapsToHCL(relayAgentFilterRules)
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_relay_agent_filter_rules" {
	start_addr = %q
	end_addr = %q
    relay_agent_filter_rules = %s
}
`, startAddr, endAddr, relayAgentFilterRulesHCL)
}


func testAccRangeSamePortControlDiscoveryBlackout(startAddr, endAddr string, samePortControlDiscoveryBlackout bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_same_port_control_discovery_blackout" {
	start_addr = %q
	end_addr = %q
    same_port_control_discovery_blackout = %t
	use_blackout_setting = true
}
`, startAddr, endAddr, samePortControlDiscoveryBlackout)
}

func testAccRangeServerAssociationType(startAddr, endAddr, serverAssociationType, failoverAssociation, member string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_server_association_type" {
    start_addr = %q
    end_addr = %q
    server_association_type = %q
    failover_association = %q
	member = {
		name = %q
	}
}
`, startAddr, endAddr, serverAssociationType, failoverAssociation, member)
}

func testAccRangeSplitMember(startAddr, endAddr string, splitMember string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_split_member" {
	start_addr = %q
	end_addr = %q
    split_member = {
		ipv4addr = %q
	}
		split_scope_exclusion_percent = 40
		ms_server = {
			ipv4addr = "10.120.23.22"
}
			server_association_type = "MS_SERVER"
}
`, startAddr, endAddr, splitMember)
}


func testAccRangeStartAddr(startAddr, endAddr string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_start_addr" {
    start_addr = %q
    end_addr = %q
}
`, startAddr, endAddr)
}

func testAccRangeSubscribeSettings(startAddr, endAddr string, subscribeSettings string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_subscribe_settings" {
    start_addr = %q
    end_addr = %q
    subscribe_settings = {
	enabled_attributes = [%q]
}
	use_subscribe_settings = true
}
`, startAddr, endAddr, subscribeSettings)
}

func testAccRangeUnknownClients(startAddr, endAddr string, unknownClients string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_unknown_clients" {
	start_addr = %q
	end_addr = %q
    unknown_clients = %q
}
`, startAddr, endAddr, unknownClients)
}

func testAccRangeUpdateDnsOnLeaseRenewal(startAddr, endAddr string, updateDnsOnLeaseRenewal bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_update_dns_on_lease_renewal" {
    start_addr = %q
    end_addr = %q
    update_dns_on_lease_renewal = %t
}
`, startAddr, endAddr, updateDnsOnLeaseRenewal)
}

func testAccRangeUseBlackoutSetting(startAddr, endAddr string, useBlackoutSetting bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_blackout_setting" {
	start_addr = %q
	end_addr = %q
	use_blackout_setting = %t
}
`, startAddr, endAddr, useBlackoutSetting)
}

func testAccRangeUseBootfile(startAddr, endAddr, bootfile string, useBootfile bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_bootfile" {
	start_addr = %q
	end_addr = %q
	bootfile = %q
    use_bootfile = %t
}
`, startAddr, endAddr, bootfile, useBootfile)
}

func testAccRangeUseBootserver(startAddr, endAddr, bootServer string, useBootserver bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_bootserver" {
    start_addr = %q
    end_addr = %q
	bootserver = %q
    use_bootserver = %t
}
`, startAddr, endAddr, bootServer, useBootserver)
}

func testAccRangeUseDdnsDomainname(startAddr, endAddr, ddnsDomainName string, useDdnsDomainname bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_ddns_domainname" {
	start_addr = %q
	end_addr = %q
	ddns_domainname = %q
    use_ddns_domainname = %t
}
`, startAddr, endAddr, ddnsDomainName, useDdnsDomainname)
}

func testAccRangeUseDdnsGenerateHostname(startAddr, endAddr string, useDdnsGenerateHostname, ddnsGenerateHostname bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_ddns_generate_hostname" {
	start_addr = %q
	end_addr = %q
	ddns_generate_hostname = %t
    use_ddns_generate_hostname = %t
}
`, startAddr, endAddr, ddnsGenerateHostname, useDdnsGenerateHostname)
}

func testAccRangeUseDenyBootp(startAddr, endAddr string, useDenyBootp, denyBootp bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_deny_bootp" {
	start_addr = %q
	end_addr = %q
    use_deny_bootp = %t
	deny_bootp = %t
}
`, startAddr, endAddr, useDenyBootp, denyBootp)
}

func testAccRangeUseDiscoveryBasicPollingSettings(startAddr, endAddr string, useDiscoveryBasicPollingSettings bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_discovery_basic_polling_settings" {
	start_addr = %q
	end_addr = %q
	use_discovery_basic_polling_settings = %t
}
`, startAddr, endAddr, useDiscoveryBasicPollingSettings)
}

func testAccRangeUseEmailList(startAddr, endAddr string, useEmailList bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_email_list" {
	start_addr = %q
	end_addr = %q
    use_email_list = %t
}
`, startAddr, endAddr, useEmailList)
}

func testAccRangeUseEnableDdns(startAddr, endAddr string, useEnableDdns, enableDdns bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_enable_ddns" {
    start_addr = %q
    end_addr = %q
	enable_ddns = %t
    use_enable_ddns = %t
}
`, startAddr, endAddr, enableDdns, useEnableDdns)
}

func testAccRangeUseEnableDhcpThresholds(startAddr, endAddr string, useEnableDhcpThresholds bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_enable_dhcp_thresholds" {
    start_addr = %q
    end_addr = %q
    use_enable_dhcp_thresholds = %t
}
`, startAddr, endAddr, useEnableDhcpThresholds)
}

func testAccRangeUseEnableDiscovery(startAddr, endAddr string, useEnableDiscovery bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_enable_discovery" {
	start_addr = %q
	end_addr = %q
    use_enable_discovery = %t
}
`, startAddr, endAddr, useEnableDiscovery)
}

func testAccRangeUseEnableIfmapPublishing(startAddr, endAddr string, useEnableIfmapPublishing bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_enable_ifmap_publishing" {
	start_addr = %q
	end_addr = %q
    use_enable_ifmap_publishing = %t
}
`, startAddr, endAddr, useEnableIfmapPublishing)
}

func testAccRangeUseIgnoreDhcpOptionListRequest(startAddr, endAddr string, useIgnoreDhcpOptionListRequest bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_ignore_dhcp_option_list_request" {
	start_addr = %q
	end_addr = %q
	use_ignore_dhcp_option_list_request = %t
}
`, startAddr, endAddr, useIgnoreDhcpOptionListRequest)
}

func testAccRangeUseIgnoreId(startAddr, endAddr string, useIgnoreId bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_ignore_id" {
    start_addr = %q
    end_addr = %q
    use_ignore_id = %t
}
`, startAddr, endAddr, useIgnoreId)
}

func testAccRangeUseKnownClients(startAddr, endAddr string, useKnownClients bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_known_clients" {
	start_addr = %q
	end_addr = %q
    use_known_clients = %t
}
`, startAddr, endAddr, useKnownClients)
}

func testAccRangeUseLeaseScavengeTime(startAddr, endAddr string, useLeaseScavengeTime bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_lease_scavenge_time" {
	start_addr = %q
	end_addr = %q
	use_lease_scavenge_time = %t
}
`, startAddr, endAddr, useLeaseScavengeTime)
}

func testAccRangeUseLogicFilterRules(startAddr, endAddr string, useLogicFilterRules bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_logic_filter_rules" {
	start_addr = %q
	end_addr = %q
	use_logic_filter_rules = %t
}
`, startAddr, endAddr, useLogicFilterRules)
}

func testAccRangeUseMsOptions(startAddr, endAddr string, useMsOptions bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_ms_options" {
	start_addr = %q
	end_addr = %q
    use_ms_options = %t
}
`, startAddr, endAddr, useMsOptions)
}

func testAccRangeUseNextserver(startAddr, endAddr string, useNextserver bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_nextserver" {
	start_addr = %q
	end_addr = %q
	use_nextserver = %t
}
`, startAddr, endAddr, useNextserver)
}

func testAccRangeUseOptions(startAddr, endAddr string, useOptions bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_options" {
	start_addr = %q
	end_addr = %q
    use_options = %t
}
`, startAddr, endAddr, useOptions)
}

func testAccRangeUsePxeLeaseTime(startAddr, endAddr string, usePxeLeaseTime bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_pxe_lease_time" {
	start_addr = %q
	end_addr = %q
	use_pxe_lease_time = %t
}
`, startAddr, endAddr, usePxeLeaseTime)
}

func testAccRangeUseRecycleLeases(startAddr, endAddr string, useRecycleLeases bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_recycle_leases" {
	start_addr = %q
	end_addr = %q
	use_recycle_leases = %t
}
`, startAddr, endAddr, useRecycleLeases)
}

func testAccRangeUseSubscribeSettings(startAddr, endAddr string, useSubscribeSettings bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_subscribe_settings" {
	start_addr = %q
	end_addr = %q
    use_subscribe_settings = %t
}
`, startAddr, endAddr, useSubscribeSettings)
}

func testAccRangeUseUnknownClients(startAddr, endAddr string, useUnknownClients bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_unknown_clients" {
	start_addr = %q
	end_addr = %q
    use_unknown_clients = %t
}
`, startAddr, endAddr, useUnknownClients)
}

func testAccRangeUseUpdateDnsOnLeaseRenewal(startAddr, endAddr string, useUpdateDnsOnLeaseRenewal bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_range" "test_use_update_dns_on_lease_renewal" {
	start_addr = %q
	end_addr = %q
    use_update_dns_on_lease_renewal = %t
}
`, startAddr, endAddr, useUpdateDnsOnLeaseRenewal)
}

func formatExtAttrsForHCL(extAttrs map[string]string) string {
	var result strings.Builder
	result.WriteString("{\n")
	for k, v := range extAttrs {
		result.WriteString(fmt.Sprintf("        %s = %q\n", k, v))
	}
	result.WriteString("    }")
	return result.String()
}
