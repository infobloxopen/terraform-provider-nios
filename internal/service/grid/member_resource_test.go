package grid_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// License Require for Test Execution - GRID

var readableAttributesForMember = "active_position,additional_ip_list,automated_traffic_capture_setting,bgp_as,comment,config_addr_type,csp_access_key,csp_member_setting,dns_resolver_setting,dscp,email_setting,enable_ha,enable_lom,enable_member_redirect,enable_ro_api_access,extattrs,external_syslog_backup_servers,external_syslog_server_enable,ha_cloud_platform,ha_on_cloud,host_name,ipv6_setting,ipv6_static_routes,is_dscp_capable,lan2_enabled,lan2_port_setting,lom_network_config,lom_users,master_candidate,member_service_communication,mgmt_port_setting,mmdb_ea_build_time,mmdb_geoip_build_time,nat_setting,node_info,ntp_setting,ospf_list,passive_ha_arp_enabled,platform,pre_provisioning,preserve_if_owns_delegation,remote_console_access_enable,router_id,service_status,service_type_configuration,snmp_setting,static_routes,support_access_enable,support_access_info,syslog_proxy_setting,syslog_servers,syslog_size,threshold_traps,time_zone,traffic_capture_auth_dns_setting,traffic_capture_chr_setting,traffic_capture_qps_setting,traffic_capture_rec_dns_setting,traffic_capture_rec_queries_setting,trap_notifications,upgrade_group,use_automated_traffic_capture,use_dns_resolver_setting,use_dscp,use_email_setting,use_enable_lom,use_enable_member_redirect,use_external_syslog_backup_servers,use_remote_console_access_enable,use_snmp_setting,use_support_access_enable,use_syslog_proxy_setting,use_threshold_traps,use_time_zone,use_traffic_capture_auth_dns,use_traffic_capture_chr,use_traffic_capture_qps,use_traffic_capture_rec_dns,use_traffic_capture_rec_queries,use_trap_notifications,use_v4_vrrp,vip_setting,vpn_mtu"

func TestAccMemberResource_basic(t *testing.T) {
	var resourceName = "nios_grid_member.test"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.100"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberBasicConfig(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "host_name", hostName),
					resource.TestCheckResourceAttr(resourceName, "config_addr_type", "IPV4"),
					resource.TestCheckResourceAttr(resourceName, "platform", "VNIOS"),
					resource.TestCheckResourceAttr(resourceName, "service_type_configuration", "ALL_V4"),

					// ipv6_setting validations
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.auto_router_config_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.dscp", "0"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.primary", "true"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.use_dscp", "false"),

					// vip_setting validations
					resource.TestCheckResourceAttr(resourceName, "vip_setting.address", vipAddress),
					resource.TestCheckResourceAttr(resourceName, "vip_setting.dscp", "0"),
					resource.TestCheckResourceAttr(resourceName, "vip_setting.gateway", "172.28.38.1"),
					resource.TestCheckResourceAttr(resourceName, "vip_setting.primary", "true"),
					resource.TestCheckResourceAttr(resourceName, "vip_setting.subnet_mask", "255.255.254.0"),
					resource.TestCheckResourceAttr(resourceName, "vip_setting.use_dscp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_disappears(t *testing.T) {
	resourceName := "nios_grid_member.test"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.101"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMemberDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMemberBasicConfig(hostName, "IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					testAccCheckMemberDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccMemberResource_AdditionalIpList(t *testing.T) {
	var resourceName = "nios_grid_member.test_additional_ip_list"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.248"

	additionalIpListVal := []map[string]any{
		{
			"anycast":     true,
			"enable_bgp":  true,
			"enable_ospf": false,
			"interface":   "LOOPBACK",
			"ipv4_network_setting": map[string]any{
				"address":     "172.21.3.2",
				"dscp":        0,
				"subnet_mask": "255.255.255.255",
				"use_dscp":    false,
			},
		},
	}
	bgpAsVal := []map[string]any{
		{
			"as":          100,
			"holddown":    16,
			"keepalive":   4,
			"link_detect": false,
			"neighbors": []map[string]any{
				{
					"authentication_mode": "NONE",
					"enable_bfd":          false,
					"enable_bfd_dnscheck": true,
					"interface":           "LAN_HA",
					"multihop":            false,
					"multihop_ttl":        255,
					"neighbor_ip":         "172.21.3.41",
					"remote_as":           1233,
				},
			},
		},
	}
	additionalIpListValUpdated := []map[string]any{
		{
			"anycast":     false,
			"enable_bgp":  false,
			"enable_ospf": true,
			"interface":   "LOOPBACK",
			"ipv4_network_setting": map[string]any{
				"address":     "172.21.3.4",
				"dscp":        0,
				"subnet_mask": "255.255.255.255",
				"use_dscp":    false,
			},
		},
	}
	ospfListVal := []map[string]any{
		{
			"area_id":                "121",
			"area_type":              "STANDARD",
			"authentication_type":    "NONE",
			"auto_calc_cost_enabled": true,
			"cost":                   1,
			"dead_interval":          40,
			"enable_bfd":             false,
			"enable_bfd_dnscheck":    true,
			"hello_interval":         10,
			"interface":              "LAN_HA",
			"is_ipv4":                true,
			"key_id":                 1,
			"retransmit_interval":    5,
			"transmit_delay":         1,
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberAdditionalIpList(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					additionalIpListVal,
					bgpAsVal,
					[]map[string]any{},
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "additional_ip_list.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "additional_ip_list.0.anycast", "true"),
					resource.TestCheckResourceAttr(resourceName, "additional_ip_list.0.enable_bgp", "true"),
					resource.TestCheckResourceAttr(resourceName, "additional_ip_list.0.enable_ospf", "false"),
					resource.TestCheckResourceAttr(resourceName, "additional_ip_list.0.interface", "LOOPBACK"),
					resource.TestCheckResourceAttr(resourceName, "additional_ip_list.0.ipv4_network_setting.address", "172.21.3.2"),
					resource.TestCheckResourceAttr(resourceName, "additional_ip_list.0.ipv4_network_setting.dscp", "0"),
					resource.TestCheckResourceAttr(resourceName, "additional_ip_list.0.ipv4_network_setting.primary", "false"),
					resource.TestCheckResourceAttr(resourceName, "additional_ip_list.0.ipv4_network_setting.subnet_mask", "255.255.255.255"),
					resource.TestCheckResourceAttr(resourceName, "additional_ip_list.0.ipv4_network_setting.use_dscp", "false"),
				),
			},
			{
				Config: testAccMemberAdditionalIpList(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					additionalIpListValUpdated,
					[]map[string]any{},
					ospfListVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "additional_ip_list.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "additional_ip_list.0.anycast", "false"),
					resource.TestCheckResourceAttr(resourceName, "additional_ip_list.0.enable_bgp", "false"),
					resource.TestCheckResourceAttr(resourceName, "additional_ip_list.0.enable_ospf", "true"),
					resource.TestCheckResourceAttr(resourceName, "additional_ip_list.0.interface", "LOOPBACK"),
					resource.TestCheckResourceAttr(resourceName, "additional_ip_list.0.ipv4_network_setting.address", "172.21.3.4"),
				),
			},
		},
	})
}

func TestAccMemberResource_AutomatedTrafficCaptureSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_automated_traffic_capture_setting"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.103"

	automatedTrafficCaptureSettingVal := map[string]any{
		"destination":            "NONE",
		"include_support_bundle": false,
		"keep_local_copy":        false,
		"traffic_capture_enable": false,
	}
	automatedTrafficCaptureSettingValUpdated := map[string]any{
		"destination":               "FTP",
		"include_support_bundle":    true,
		"keep_local_copy":           true,
		"traffic_capture_enable":    true,
		"duration":                  5,
		"destination_host":          "192.28.0.1",
		"traffic_capture_directory": "192.28.0.1",
		"username":                  "user",
		"password":                  "nios",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberAutomatedTrafficCaptureSetting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					automatedTrafficCaptureSettingVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "automated_traffic_capture_setting.destination", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "automated_traffic_capture_setting.include_support_bundle", "false"),
					resource.TestCheckResourceAttr(resourceName, "automated_traffic_capture_setting.keep_local_copy", "false"),
					resource.TestCheckResourceAttr(resourceName, "automated_traffic_capture_setting.traffic_capture_enable", "false"),
				),
			},
			{
				Config: testAccMemberAutomatedTrafficCaptureSetting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					automatedTrafficCaptureSettingValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "automated_traffic_capture_setting.destination", "FTP"),
					resource.TestCheckResourceAttr(resourceName, "automated_traffic_capture_setting.include_support_bundle", "true"),
					resource.TestCheckResourceAttr(resourceName, "automated_traffic_capture_setting.keep_local_copy", "true"),
					resource.TestCheckResourceAttr(resourceName, "automated_traffic_capture_setting.traffic_capture_enable", "true"),
				),
			},
		},
	})
}

func TestAccMemberResource_BgpAs(t *testing.T) {
	var resourceName = "nios_grid_member.test_bgp_as"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.104"

	bgpAsVal := []map[string]any{
		{
			"as":          100,
			"holddown":    16,
			"keepalive":   4,
			"link_detect": false,
			"neighbors": []map[string]any{
				{
					"authentication_mode": "NONE",
					"enable_bfd":          false,
					"enable_bfd_dnscheck": true,
					"interface":           "LAN_HA",
					"multihop":            false,
					"multihop_ttl":        255,
					"neighbor_ip":         "172.21.3.41",
					"remote_as":           1233,
				},
			},
		},
	}
	bgpAsValUpdated := []map[string]any{
		{
			"as":          200,
			"holddown":    20,
			"keepalive":   5,
			"link_detect": true,
			"neighbors": []map[string]any{
				{
					"authentication_mode": "NONE",
					"enable_bfd":          false,
					"enable_bfd_dnscheck": true,
					"interface":           "LAN_HA",
					"multihop":            false,
					"multihop_ttl":        255,
					"neighbor_ip":         "172.21.3.42",
					"remote_as":           2233,
				},
			},
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberBgpAs(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					bgpAsVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bgp_as.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "bgp_as.0.as", "100"),
					resource.TestCheckResourceAttr(resourceName, "bgp_as.0.holddown", "16"),
					resource.TestCheckResourceAttr(resourceName, "bgp_as.0.keepalive", "4"),
					resource.TestCheckResourceAttr(resourceName, "bgp_as.0.link_detect", "false"),
					resource.TestCheckResourceAttr(resourceName, "bgp_as.0.neighbors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "bgp_as.0.neighbors.0.neighbor_ip", "172.21.3.41"),
					resource.TestCheckResourceAttr(resourceName, "bgp_as.0.neighbors.0.remote_as", "1233"),
				),
			},
			{
				Config: testAccMemberBgpAs(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					bgpAsValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bgp_as.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "bgp_as.0.as", "200"),
					resource.TestCheckResourceAttr(resourceName, "bgp_as.0.holddown", "20"),
					resource.TestCheckResourceAttr(resourceName, "bgp_as.0.keepalive", "5"),
					resource.TestCheckResourceAttr(resourceName, "bgp_as.0.link_detect", "true"),
					resource.TestCheckResourceAttr(resourceName, "bgp_as.0.neighbors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "bgp_as.0.neighbors.0.neighbor_ip", "172.21.3.42"),
					resource.TestCheckResourceAttr(resourceName, "bgp_as.0.neighbors.0.remote_as", "2233"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_Comment(t *testing.T) {
	var resourceName = "nios_grid_member.test_comment"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.105"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberComment(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					"Test comment",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Test comment"),
				),
			},
			{
				Config: testAccMemberComment(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					"Test comment updated",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Test comment updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_ConfigAddrType(t *testing.T) {
	var resourceName = "nios_grid_member.test_config_addr_type"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress4 := "172.28.38.106"
	vipAddress6 := fmt.Sprintf("2001:db8:%x:%x::%x", acctest.RandomNumber(65535), acctest.RandomNumber(65535), acctest.RandomNumber(65535))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberConfigAddrTypeIPv4(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress4,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "config_addr_type", "IPV4"),
				),
			},
			{
				Config: testAccMemberConfigAddrTypeIPv6(
					hostName,
					"IPV6",
					"VNIOS",
					"ALL_V6",
					vipAddress6,
					"2001::1",
					8,
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "config_addr_type", "IPV6"),
				),
			},
		},
	})
}

func TestAccMemberResource_CspAccessKey(t *testing.T) {
	var resourceName = "nios_grid_member.test_csp_access_key"
	var v grid.Member
	t.Skip("Insertion and update not allowed for csp access key")
	vipAddress := "172.28.38.107"
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	cspAccessKeyVal := []string{"CSP_ACCESS_KEY_REPLACE_ME1"}
	cspAccessKeyValUpdated := []string{"CSP_ACCESS_KEY_REPLACE_ME1"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberCspAccessKey(hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					cspAccessKeyVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "csp_access_key", "CSP_ACCESS_KEY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberCspAccessKey(hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					cspAccessKeyValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "csp_access_key", "CSP_ACCESS_KEY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_CspMemberSetting(t *testing.T) {
	t.Skip("CSP member setting cannot be updated due to member being offline")
	var resourceName = "nios_grid_member.test_csp_member_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.108"

	cspMemberSettingVal := map[string]any{
		"csp_dns_resolver":     "2.2.2.2",
		"use_csp_dns_resolver": false,
		"use_csp_https_proxy":  false,
		"use_csp_join_token":   false,
	}
	cspMemberSettingValUpdated := map[string]any{
		"csp_dns_resolver":     "1.1.1.1",
		"use_csp_dns_resolver": true,
		"use_csp_https_proxy":  true,
		"use_csp_join_token":   true,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberCspMemberSetting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					cspMemberSettingVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "csp_member_setting.csp_dns_resolver", "2.2.2.2"),
					resource.TestCheckResourceAttr(resourceName, "csp_member_setting.use_csp_dns_resolver", "false"),
					resource.TestCheckResourceAttr(resourceName, "csp_member_setting.use_csp_https_proxy", "false"),
					resource.TestCheckResourceAttr(resourceName, "csp_member_setting.use_csp_join_token", "false"),
				),
			},
			{
				Config: testAccMemberCspMemberSetting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					cspMemberSettingValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "csp_member_setting.csp_dns_resolver", "1.1.1.1"),
					resource.TestCheckResourceAttr(resourceName, "csp_member_setting.use_csp_dns_resolver", "true"),
					resource.TestCheckResourceAttr(resourceName, "csp_member_setting.use_csp_https_proxy", "true"),
					resource.TestCheckResourceAttr(resourceName, "csp_member_setting.use_csp_join_token", "true"),
				),
			},
		},
	})
}

func TestAccMemberResource_DnsResolverSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_dns_resolver_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.109"

	dnsResolverSettingVal := map[string]any{
		"resolvers":      []string{"10.0.0.1"},
		"search_domains": []string{"a.com"},
	}
	dnsResolverSettingValUpdated := map[string]any{
		"resolvers":      []string{"10.0.0.2"},
		"search_domains": []string{"b.com"},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberDnsResolverSetting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					dnsResolverSettingVal,
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service_type_configuration", "ALL_V4"),
					resource.TestCheckResourceAttr(resourceName, "dns_resolver_setting.resolvers.0", "10.0.0.1"),
					resource.TestCheckResourceAttr(resourceName, "dns_resolver_setting.search_domains.0", "a.com"),
					resource.TestCheckResourceAttr(resourceName, "use_dns_resolver_setting", "true"),
				),
			},
			{
				Config: testAccMemberDnsResolverSetting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					dnsResolverSettingValUpdated,
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_resolver_setting.resolvers.0", "10.0.0.2"),
					resource.TestCheckResourceAttr(resourceName, "dns_resolver_setting.search_domains.0", "b.com"),
					resource.TestCheckResourceAttr(resourceName, "use_dns_resolver_setting", "true"),
				),
			},
		},
	})
}

func TestAccMemberResource_Dscp(t *testing.T) {
	var resourceName = "nios_grid_member.test_dscp"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.110"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberDscp(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					0, //DSCP cannot be configured on an IB-UNKNOWN appliance, can only check the default value which is 0
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "host_name", hostName),
					resource.TestCheckResourceAttr(resourceName, "config_addr_type", "IPV4"),
					resource.TestCheckResourceAttr(resourceName, "platform", "VNIOS"),
					resource.TestCheckResourceAttr(resourceName, "service_type_configuration", "ALL_V4"),
					resource.TestCheckResourceAttr(resourceName, "use_dscp", "true"),
					resource.TestCheckResourceAttr(resourceName, "dscp", "0"),
				),
			},
			{
				Config: testAccMemberDscp(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					0,
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dscp", "0"),
					resource.TestCheckResourceAttr(resourceName, "use_dscp", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_EmailSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_email_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.111"

	emailSettingVal := map[string]any{
		"enabled":            false,
		"port_number":        25,
		"relay_enabled":      false,
		"smtps":              false,
		"use_authentication": false,
		"from_address":       "nios.provider@infoblox.com",
	}
	emailSettingValUpdated := map[string]any{
		"enabled":            true,
		"port_number":        587,
		"relay_enabled":      true,
		"relay":              "smtp.relay.com",
		"smtps":              true,
		"use_authentication": true,
		"password":           "nios",
		"from_address":       "nios_sender.provider@infoblox.com",
		"address":            "nios.provider@infoblox.com",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberEmailSetting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					emailSettingVal,
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_setting.enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "email_setting.port_number", "25"),
					resource.TestCheckResourceAttr(resourceName, "email_setting.relay_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "email_setting.smtps", "false"),
					resource.TestCheckResourceAttr(resourceName, "email_setting.use_authentication", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_email_setting", "false"),
				),
			},
			{
				Config: testAccMemberEmailSetting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					emailSettingValUpdated,
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_setting.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "email_setting.port_number", "587"),
					resource.TestCheckResourceAttr(resourceName, "email_setting.relay_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "email_setting.smtps", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_email_setting", "true"),
					resource.TestCheckResourceAttr(resourceName, "email_setting.use_authentication", "true"),
					resource.TestCheckResourceAttr(resourceName, "email_setting.from_address", "nios_sender.provider@infoblox.com"),
					resource.TestCheckResourceAttr(resourceName, "email_setting.address", "nios.provider@infoblox.com"),
				),
			},
		},
	})
}

func TestAccMemberResource_EnableHa(t *testing.T) {
	var resourceName = "nios_grid_member.test_enable_ha"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.112"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberEnableHa(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					true, 197,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ha", "true"),
					resource.TestCheckResourceAttr(resourceName, "router_id", "197"),
					resource.TestCheckResourceAttr(resourceName, "node_info.#", "2"),
				),
			},
		},
	})
}

func TestAccMemberResource_EnableLom(t *testing.T) {
	var resourceName = "nios_grid_member.test_enable_lom"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.113"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberEnableLom(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_lom", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_lom", "true"),
				),
			},
			{
				Config: testAccMemberEnableLom(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_lom", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_lom", "true"),
				),
			},
		},
	})
}

func TestAccMemberResource_EnableMemberRedirect(t *testing.T) {
	var resourceName = "nios_grid_member.test_enable_member_redirect"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.114"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberEnableMemberRedirect(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_member_redirect", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_member_redirect", "true"),
				),
			},
			{
				Config: testAccMemberEnableMemberRedirect(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_member_redirect", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_member_redirect", "true"),
				),
			},
		},
	})
}

func TestAccMemberResource_EnableRoApiAccess(t *testing.T) {
	var resourceName = "nios_grid_member.test_enable_ro_api_access"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.115"

	networkSettingAddress6 := fmt.Sprintf("2001:db8:%x:%x::%x", acctest.RandomNumber(65535), acctest.RandomNumber(65535), acctest.RandomNumber(65535))
	ipv6SettingVal := map[string]any{
		"auto_router_config_enabled": false,
		"dscp":                       0,
		"enabled":                    true,
		"use_dscp":                   false,
		"virtual_ip":                 networkSettingAddress6,
		"gateway":                    "2001::1",
		"primary":                    true,
		"cidr_prefix":                8,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberEnableRoApiAccess(
					hostName,
					"BOTH",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
					ipv6SettingVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ro_api_access", "true"),
				),
			},
			{
				Config: testAccMemberEnableRoApiAccess(
					hostName,
					"BOTH",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
					ipv6SettingVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ro_api_access", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_grid_member.test_extattrs"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.116"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberExtAttrs(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					map[string]any{
						"Site": extAttrValue1,
					},
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			{
				Config: testAccMemberExtAttrs(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					map[string]any{
						"Site": extAttrValue2,
					},
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
		},
	})
}

func TestAccMemberResource_ExternalSyslogBackupServers(t *testing.T) {
	var resourceName = "nios_grid_member.test_external_syslog_backup_servers"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.117"

	externalSyslogBackupServersVal := []map[string]any{
		{
			"address_or_fqdn": "192.0.2.10",
			"directory_path":  "/var/log/backup",
			"enable":          true,
			"port":            21,
			"protocol":        "FTP",
			"username":        "admin1",
			"password":        "Password123!",
		},
	}
	externalSyslogBackupServersValUpdated := []map[string]any{
		{
			"address_or_fqdn": "192.0.2.20",
			"directory_path":  "/var/log/backup2",
			"enable":          false,
			"port":            22,
			"protocol":        "SCP",
			"username":        "admin2",
			"password":        "Password456!",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberExternalSyslogBackupServers(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					externalSyslogBackupServersVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_syslog_backup_servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "external_syslog_backup_servers.0.address_or_fqdn", "192.0.2.10"),
					resource.TestCheckResourceAttr(resourceName, "external_syslog_backup_servers.0.protocol", "FTP"),
					resource.TestCheckResourceAttr(resourceName, "external_syslog_backup_servers.0.port", "21"),
					resource.TestCheckResourceAttr(resourceName, "external_syslog_backup_servers.0.enable", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_external_syslog_backup_servers", "true"),
				),
			},
			{
				Config: testAccMemberExternalSyslogBackupServers(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					externalSyslogBackupServersValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_syslog_backup_servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "external_syslog_backup_servers.0.address_or_fqdn", "192.0.2.20"),
					resource.TestCheckResourceAttr(resourceName, "external_syslog_backup_servers.0.protocol", "SCP"),
					resource.TestCheckResourceAttr(resourceName, "external_syslog_backup_servers.0.port", "22"),
					resource.TestCheckResourceAttr(resourceName, "external_syslog_backup_servers.0.enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_external_syslog_backup_servers", "true"),
				),
			},
		},
	})
}

func TestAccMemberResource_HaCloudPlatform(t *testing.T) {
	var resourceName = "nios_grid_member.test_ha_cloud_platform"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.118"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberHaCloudPlatform(hostName,
					"AWS",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ha_cloud_platform", "AWS"),
				),
			},
			{
				Config: testAccMemberHaCloudPlatform(hostName,
					"AZURE",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ha_cloud_platform", "AZURE"),
				),
			},
			{
				Config: testAccMemberHaCloudPlatform(hostName,
					"GCP",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ha_cloud_platform", "GCP"),
				),
			},
		},
	})
}

func TestAccMemberResource_HaOnCloud(t *testing.T) {
	var resourceName = "nios_grid_member.test_ha_on_cloud"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.119"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberHaOnCloud(hostName,
					"false", "None",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0", "false", 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ha_on_cloud", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberHaOnCloud(hostName,
					"true", "AWS",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0", "true", 115),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ha_on_cloud", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_HostName(t *testing.T) {
	var resourceName = "nios_grid_member.test_host_name"
	var v grid.Member

	hostName1 := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	hostName2 := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.120"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberHostName(
					hostName1,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "host_name", hostName1),
				),
			},
			{
				Config: testAccMemberHostName(
					hostName2,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "host_name", hostName2),
				),
			},
		},
	})
}

func TestAccMemberResource_Ipv6Setting(t *testing.T) {
	var resourceName = "nios_grid_member.test_ipv6_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	networkSettingAddress := fmt.Sprintf("2001:db8:%x:%x::%x", acctest.RandomNumber(65535), acctest.RandomNumber(65535), acctest.RandomNumber(65535))
	networkSettingAddress6 := fmt.Sprintf("2001:db8:%x:%x::%x", acctest.RandomNumber(65535), acctest.RandomNumber(65535), acctest.RandomNumber(65535))

	vipAddress := "172.28.38.121"

	ipv6SettingVal := map[string]any{
		"auto_router_config_enabled": false,
		"dscp":                       0,
		"enabled":                    true,
		"use_dscp":                   false,
		"virtual_ip":                 networkSettingAddress,
		"gateway":                    "2001::1",
		"primary":                    true,
		"cidr_prefix":                8,
	}
	ipv6SettingValUpdated := map[string]any{
		"auto_router_config_enabled": false,
		"dscp":                       0,
		"enabled":                    true,
		"use_dscp":                   false,
		"virtual_ip":                 networkSettingAddress6,
		"gateway":                    "2001::1",
		"primary":                    true,
		"cidr_prefix":                8,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberIpv6Setting(
					hostName,
					"BOTH",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					ipv6SettingVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.auto_router_config_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.dscp", "0"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.use_dscp", "false"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.virtual_ip", networkSettingAddress),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.gateway", "2001::1"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.primary", "true"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.cidr_prefix", "8"),
				),
			},
			{
				Config: testAccMemberIpv6Setting(
					hostName,
					"BOTH",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					ipv6SettingValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.auto_router_config_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.dscp", "0"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.use_dscp", "false"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.virtual_ip", networkSettingAddress6),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.gateway", "2001::1"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.primary", "true"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.cidr_prefix", "8"),
				),
			},
		},
	})
}

func TestAccMemberResource_Ipv6StaticRoutes(t *testing.T) {
	t.Skip("IPv6 Static Routes cannot be created if other routes exist.")
	var resourceName = "nios_grid_member.test_ipv6_static_routes"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.122"
	networkSettingAddress := fmt.Sprintf("2001:db8:%x:%x::%x", acctest.RandomNumber(65535), acctest.RandomNumber(65535), acctest.RandomNumber(65535))
	ipv6StaticRoutesVal := []map[string]any{
		{
			"address": networkSettingAddress,
			"cidr":    64,
			"gateway": "2001::1",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberIpv6StaticRoutes(hostName,
					ipv6StaticRoutesVal,
					vipAddress,
					"172.28.38.1",
					"255.255.254.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_static_routes", "IPV6_STATIC_ROUTES_REPLACE_ME"),
				),
			},
			// Cannot change or remove IP address configuration due to existing static routes.
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_Lan2Enabled(t *testing.T) {
	var resourceName = "nios_grid_member.test_lan2_enabled"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.123"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberLan2Enabled(hostName,
					"true",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lan2_enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberLan2EnabledFalse(hostName,
					"false",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lan2_enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_Lan2PortSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_lan2_port_setting"
	var v grid.Member
	networkSettingAddress := fmt.Sprintf("172.29.38.%d", acctest.RandomNumber(254))
	networkSettingAddress6 := fmt.Sprintf("2001:db8:%x:%x::%x", acctest.RandomNumber(65535), acctest.RandomNumber(65535), acctest.RandomNumber(65535))
	lan2PortSettingVal := map[string]any{
		"virtual_router_id": 10,
		"enabled":           true,
		"network_setting": map[string]any{
			"address":     networkSettingAddress,
			"gateway":     "172.29.38.1",
			"primary":     true,
			"subnet_mask": "255.255.0.0",
		},
		"nic_failover_enabled":           false,
		"nic_failover_enable_primary":    false,
		"default_route_failover_enabled": false,
	}
	lan2PortSettingValUpdated := map[string]any{
		"virtual_router_id": 10,
		"enabled":           true,
		"v6_network_setting": map[string]any{
			"virtual_ip":  networkSettingAddress6,
			"gateway":     "2001::1",
			"primary":     true,
			"cidr_prefix": 8,
		},
	}
	lan2PortSettingValUpdated2 := map[string]any{
		"virtual_router_id":              10,
		"enabled":                        true,
		"nic_failover_enabled":           true,
		"nic_failover_enable_primary":    true,
		"default_route_failover_enabled": false,
	}

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.124"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberLan2PortSetting(hostName,
					lan2PortSettingVal,
					vipAddress,
					"172.28.38.1",
					"255.255.254.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting.virtual_router_id", "10"),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting.network_setting.address", networkSettingAddress),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting.network_setting.gateway", "172.29.38.1"),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting.network_setting.primary", "true"),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting.network_setting.subnet_mask", "255.255.0.0"),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting.nic_failover_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting.nic_failover_enable_primary", "false"),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting.default_route_failover_enabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberLan2PortSetting(hostName,
					lan2PortSettingValUpdated,
					vipAddress,
					"172.28.38.1",
					"255.255.254.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting.v6_network_setting.virtual_ip", networkSettingAddress6),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting.v6_network_setting.gateway", "2001::1"),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting.v6_network_setting.primary", "true"),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting.v6_network_setting.cidr_prefix", "8"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberLan2PortSetting(hostName,
					lan2PortSettingValUpdated2,
					vipAddress,
					"172.28.38.1",
					"255.255.254.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting.virtual_router_id", "10"),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting.nic_failover_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting.nic_failover_enable_primary", "true"),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting.default_route_failover_enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_LomNetworkConfig(t *testing.T) {
	var resourceName = "nios_grid_member.test_lom_network_config"
	var v grid.Member
	networkSettingAddress := fmt.Sprintf("172.29.38.%d", acctest.RandomNumber(254))
	vipAddress := "172.28.38.125"
	lomNetworkConfigVal := []map[string]any{
		{
			"address":     networkSettingAddress,
			"gateway":     "172.29.38.1",
			"subnet_mask": "255.255.0.0",
		},
	}
	lomNetworkConfigValUpdated := []map[string]any{
		{
			"address":     networkSettingAddress,
			"gateway":     "172.29.38.1",
			"subnet_mask": "255.255.0.0",
		},
	}

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberLomNetworkConfig(hostName, lomNetworkConfigVal,
					vipAddress,
					"172.28.38.1",
					"255.255.254.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lom_network_config.0.address", networkSettingAddress),
					resource.TestCheckResourceAttr(resourceName, "lom_network_config.0.gateway", "172.29.38.1"),
					resource.TestCheckResourceAttr(resourceName, "lom_network_config.0.subnet_mask", "255.255.0.0"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberLomNetworkConfig(hostName,
					lomNetworkConfigValUpdated,
					vipAddress,
					"172.28.38.1",
					"255.255.254.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lom_network_config.0.address", networkSettingAddress),
					resource.TestCheckResourceAttr(resourceName, "lom_network_config.0.gateway", "172.29.38.1"),
					resource.TestCheckResourceAttr(resourceName, "lom_network_config.0.subnet_mask", "255.255.0.0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_LomUsers(t *testing.T) {
	var resourceName = "nios_grid_member.test_lom_users"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.126"

	lomUsersVal := []map[string]any{
		{
			"disable":  false,
			"name":     "LOMuser1",
			"password": "@#nios123",
			"role":     "USER",
		},
	}

	lomUsersValUpdated := []map[string]any{
		{
			"disable":  true,
			"name":     "LOMuser2",
			"password": "@#nios1234",
			"role":     "OPERATOR",
			"comment":  "Updated user",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberLomUsers(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					lomUsersVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lom_users.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "lom_users.0.name", "LOMuser1"),
					resource.TestCheckResourceAttr(resourceName, "lom_users.0.role", "USER"),
					resource.TestCheckResourceAttr(resourceName, "lom_users.0.disable", "false"),
				),
			},
			// Update and Read - disable the user
			{
				Config: testAccMemberLomUsers(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					lomUsersValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lom_users.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "lom_users.0.name", "LOMuser2"),
					resource.TestCheckResourceAttr(resourceName, "lom_users.0.role", "OPERATOR"),
					resource.TestCheckResourceAttr(resourceName, "lom_users.0.disable", "true"),
					resource.TestCheckResourceAttr(resourceName, "lom_users.0.comment", "Updated user"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_MgmtPortSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_mgmt_port_setting"
	var v grid.Member

	mgmtPortSettingVal := map[string]any{
		"enabled":                 true,
		"vpn_enabled":             true,
		"security_access_enabled": true,
	}
	mgmtPortSettingValUpdated := map[string]any{
		"enabled":                 true,
		"vpn_enabled":             false,
		"security_access_enabled": false,
	}
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.128"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberMgmtPortSetting(hostName,
					mgmtPortSettingVal,
					vipAddress,
					"172.28.38.1",
					"255.255.254.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mgmt_port_setting.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "mgmt_port_setting.vpn_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "mgmt_port_setting.security_access_enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberMgmtPortSetting(hostName,
					mgmtPortSettingValUpdated,
					vipAddress,
					"172.28.38.1",
					"255.255.254.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mgmt_port_setting.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "mgmt_port_setting.vpn_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "mgmt_port_setting.security_access_enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccMemberResource_NatSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_nat_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.129"
	externalVirtualIp := fmt.Sprintf("172.28.1.%d", acctest.RandomNumber(254))
	externalVirtualIpUpdated := fmt.Sprintf("172.28.1.%d", acctest.RandomNumber(254))

	natSettingVal := map[string]any{
		"enabled":             true,
		"external_virtual_ip": externalVirtualIp,
	}
	natSettingValUpdated := map[string]any{
		"enabled":             true,
		"external_virtual_ip": externalVirtualIpUpdated,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberNatSetting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					natSettingVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nat_setting.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "nat_setting.external_virtual_ip", externalVirtualIp),
				),
			},
			// Update and Read
			{
				Config: testAccMemberNatSetting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					natSettingValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nat_setting.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "nat_setting.external_virtual_ip", externalVirtualIpUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_NodeInfo(t *testing.T) {
	var resourceName = "nios_grid_member.test_node_info"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.130"
	networkSettingAddress6 := fmt.Sprintf("2001:db8:%x:%x::%x", acctest.RandomNumber(65535), acctest.RandomNumber(65535), acctest.RandomNumber(65535))

	mgmgtPhysicalSetting := map[string]any{
		"auto_port_setting_enabled": false,
		"speed":                     "100",
		"duplex":                    "FULL",
	}

	nodeInfoVal := []map[string]any{
		{
			"lan_ha_port_setting": map[string]any{
				"ha_cloud_attribute": "UNK",
				"ha_ip_address":      "172.28.38.121",
				"ha_port_setting": map[string]any{
					"auto_port_setting_enabled": true,
					"speed":                     "10",
				},
				"lan_port_setting": map[string]any{
					"auto_port_setting_enabled": true,
				},
				"mgmt_lan": "172.28.38.132",
			},
		},
		{
			"lan_ha_port_setting": map[string]any{
				"ha_cloud_attribute": "UNK",
				"ha_ip_address":      "172.28.38.141",
				"ha_port_setting": map[string]any{
					"auto_port_setting_enabled": true,
					"speed":                     "10",
				},
				"lan_port_setting": map[string]any{
					"auto_port_setting_enabled": true,
				},
				"mgmt_lan": "172.28.38.143",
			},
		},
	}

	nodeInfoValUpdated := []map[string]any{
		{
			"lan_ha_port_setting": map[string]any{
				"ha_cloud_attribute": "UNK",
				"ha_ip_address":      "172.28.38.12",
				"ha_port_setting": map[string]any{
					"auto_port_setting_enabled": true,
					"speed":                     "10",
				},
				"lan_port_setting": map[string]any{
					"auto_port_setting_enabled": true,
				},
				"mgmt_lan": "172.28.38.33",
			},
		},
		{
			"lan_ha_port_setting": map[string]any{
				"ha_cloud_attribute": "UNK",
				"ha_ip_address":      "172.28.38.42",
				"ha_port_setting": map[string]any{
					"auto_port_setting_enabled": true,
					"speed":                     "10",
				},
				"lan_port_setting": map[string]any{
					"auto_port_setting_enabled": true,
				},
				"mgmt_lan": "172.28.38.44",
			},
		},
	}

	nodeInfoMGMTIPv4 := []map[string]any{
		{
			"mgmt_network_setting": map[string]any{
				"address":     "172.28.38.254",
				"gateway":     "172.28.38.1",
				"subnet_mask": "255.255.255.0",
			},
			"mgmt_physical_setting": mgmgtPhysicalSetting,
		},
	}

	nodeInfoMGMTIPv6 := []map[string]any{
		{
			"v6_mgmt_network_setting": map[string]any{
				"enabled":                    true,
				"virtual_ip":                 networkSettingAddress6,
				"gateway":                    "2001::1",
				"auto_router_config_enabled": false,
				"cidr_prefix":                8,
			},
			"mgmt_physical_setting": mgmgtPhysicalSetting,
		},
	}

	mgmtPortSettingVal := map[string]any{
		"enabled": false,
	}
	mgmtPortSettingValUpdated := map[string]any{
		"enabled": true,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberNodeInfo(hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0", "true", 112, nodeInfoVal, mgmtPortSettingVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "node_info.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.lan_ha_port_setting.ha_ip_address", "172.28.38.121"),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.lan_ha_port_setting.mgmt_lan", "172.28.38.132"),
					resource.TestCheckResourceAttr(resourceName, "node_info.1.lan_ha_port_setting.ha_ip_address", "172.28.38.141"),
					resource.TestCheckResourceAttr(resourceName, "node_info.1.lan_ha_port_setting.mgmt_lan", "172.28.38.143"),
				),
			},
			// Update and Read
			// Invalid Value for Router ID provided here since Router ID is not updated here. ID must be between 1-255
			{
				Config: testAccMemberNodeInfo(hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0", "false", 0, nodeInfoMGMTIPv4, mgmtPortSettingValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "node_info.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.mgmt_network_setting.address", "172.28.38.254"),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.mgmt_network_setting.gateway", "172.28.38.1"),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.mgmt_network_setting.subnet_mask", "255.255.255.0"),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.mgmt_physical_setting.auto_port_setting_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.mgmt_physical_setting.speed", "100"),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.mgmt_physical_setting.duplex", "FULL"),
				),
			},
			// Update and Read
			// Invalid Value for Router ID provided here since Router ID is not updated here. ID must be between 1-255
			{
				Config: testAccMemberNodeInfo(hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0", "false", 0, nodeInfoMGMTIPv6, mgmtPortSettingValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "node_info.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.v6_mgmt_network_setting.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.v6_mgmt_network_setting.virtual_ip", networkSettingAddress6),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.v6_mgmt_network_setting.gateway", "2001::1"),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.v6_mgmt_network_setting.auto_router_config_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.v6_mgmt_network_setting.cidr_prefix", "8"),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.mgmt_physical_setting.auto_port_setting_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.mgmt_physical_setting.speed", "100"),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.mgmt_physical_setting.duplex", "FULL"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberNodeInfo(hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0", "true", 113, nodeInfoValUpdated, mgmtPortSettingVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "node_info.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.lan_ha_port_setting.ha_ip_address", "172.28.38.12"),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.lan_ha_port_setting.mgmt_lan", "172.28.38.33"),
					resource.TestCheckResourceAttr(resourceName, "node_info.1.lan_ha_port_setting.ha_ip_address", "172.28.38.42"),
					resource.TestCheckResourceAttr(resourceName, "node_info.1.lan_ha_port_setting.mgmt_lan", "172.28.38.44"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_NtpSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_ntp_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.131"

	ntpSettingVal := map[string]any{
		"enable_external_ntp_servers":    false,
		"enable_ntp":                     false,
		"exclude_grid_master_ntp_server": false,
		"local_ntp_stratum":              15,
		"use_default_stratum":            true,
		"use_local_ntp_stratum":          false,
	}

	ntpSettingValUpdated := map[string]any{
		"enable_external_ntp_servers":    false,
		"enable_ntp":                     true,
		"exclude_grid_master_ntp_server": false,
		"local_ntp_stratum":              15,
		"use_default_stratum":            true,
		"use_local_ntp_stratum":          false,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberNtpSetting(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					ntpSettingVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ntp_setting.enable_ntp", "false"),
					resource.TestCheckResourceAttr(resourceName, "ntp_setting.local_ntp_stratum", "15"),
				),
			},
			{
				Config: testAccMemberNtpSetting(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					ntpSettingValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ntp_setting.enable_ntp", "true"),
				),
			},
		},
	})
}

func TestAccMemberResource_OspfList(t *testing.T) {
	// Authentication Key Issue
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_grid_member.test_ospf_list"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.132"

	ospfListVal := []map[string]any{
		{
			"area_id":                "121",
			"area_type":              "STANDARD",
			"authentication_type":    "NONE",
			"auto_calc_cost_enabled": true,
			"cost":                   1,
			"dead_interval":          40,
			"enable_bfd":             false,
			"enable_bfd_dnscheck":    true,
			"hello_interval":         10,
			"interface":              "LAN_HA",
			"is_ipv4":                true,
			"key_id":                 1,
			"retransmit_interval":    5,
			"transmit_delay":         1,
		},
	}
	ospfListValUpdated := []map[string]any{
		{
			"area_id":                "121",
			"area_type":              "STANDARD",
			"authentication_type":    "SIMPLE",
			"auto_calc_cost_enabled": true,
			"cost":                   2,
			"dead_interval":          40,
			"enable_bfd":             false,
			"enable_bfd_dnscheck":    true,
			"hello_interval":         10,
			"interface":              "LAN_HA",
			"is_ipv4":                true,
			"key_id":                 1,
			"retransmit_interval":    5,
			"transmit_delay":         1,
			"authentication_key":     "key",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberOspfList(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					ospfListVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ospf_list.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ospf_list.0.area_id", "121"),
					resource.TestCheckResourceAttr(resourceName, "ospf_list.0.interface", "LAN_HA"),
					resource.TestCheckResourceAttr(resourceName, "ospf_list.0.cost", "1"),
				),
			},
			{
				Config: testAccMemberOspfList(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					ospfListValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ospf_list.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ospf_list.0.cost", "2"),
				),
			},
		},
	})
}

func TestAccMemberResource_PassiveHaArpEnabled(t *testing.T) {
	var resourceName = "nios_grid_member.test_passive_ha_arp_enabled"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.133"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberPassiveHaArpEnabled(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ha", "true"),
					resource.TestCheckResourceAttr(resourceName, "router_id", "198"),
					resource.TestCheckResourceAttr(resourceName, "node_info.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "passive_ha_arp_enabled", "true"),
				),
			},
		},
	})
}

func TestAccMemberResource_Platform(t *testing.T) {
	var resourceName = "nios_grid_member.test_platform"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.134"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberPlatform(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "platform", "VNIOS"),
				),
			},
			{
				Config: testAccMemberPlatform(
					hostName, "IPV4", "INFOBLOX", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "platform", "INFOBLOX"),
				),
			},
		},
	})
}

func TestAccMemberResource_PreProvisioning(t *testing.T) {
	t.Skip("Skipping test due to NIOS-109825")
	var resourceName = "nios_grid_member.test"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.135"

	preProvisioningVal := map[string]any{
		"hardware_info": []map[string]any{
			{
				"hwtype": "IB-V926",
			},
		},
		"licenses": []string{"dns", "dhcp"},
	}

	preProvisioningValUpdated := map[string]any{
		"hardware_info": []map[string]any{
			{
				"hwtype": "IB-V926",
			},
		},
		"licenses": []string{"dns", "dhcp"},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberPreProvisioning(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					preProvisioningVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pre_provisioning.hardware_info.0.hwmodel", "IB-VM-820"),
					resource.TestCheckResourceAttr(resourceName, "pre_provisioning.hardware_info.0.hwtype", "IB-VNIOS"),
					resource.TestCheckResourceAttr(resourceName, "pre_provisioning.licenses.#", "2"),
				),
			},
			// Update pre_provisioning with new values
			{
				Config: testAccMemberPreProvisioning(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					preProvisioningValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pre_provisioning.hardware_info.0.hwmodel", "IB-VM-1420"),
					resource.TestCheckResourceAttr(resourceName, "pre_provisioning.licenses.#", "2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_PreserveIfOwnsDelegation(t *testing.T) {
	var resourceName = "nios_grid_member.test_preserve_if_owns_delegation"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.136"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberPreserveIfOwnsDelegation(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preserve_if_owns_delegation", "true"),
				),
			},
			{
				Config: testAccMemberPreserveIfOwnsDelegation(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preserve_if_owns_delegation", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_RemoteConsoleAccessEnable(t *testing.T) {
	var resourceName = "nios_grid_member.test_remote_console_access_enable"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.137"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberRemoteConsoleAccessEnable(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_console_access_enable", "true"),
				),
			},
			{
				Config: testAccMemberRemoteConsoleAccessEnable(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_console_access_enable", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_RouterId(t *testing.T) {
	var resourceName = "nios_grid_member.test_router_id"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.138"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberRouterId(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					199,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "router_id", "199"),
				),
			},
			{
				Config: testAccMemberRouterId(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					201,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "router_id", "201"),
				),
			},
		},
	})
}

func TestAccMemberResource_ServiceTypeConfiguration(t *testing.T) {
	var resourceName = "nios_grid_member.test_service_type_configuration"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress4 := "172.28.38.139"
	networkSettingAddress6 := fmt.Sprintf("2001:db8:%x:%x::%x", acctest.RandomNumber(65535), acctest.RandomNumber(65535), acctest.RandomNumber(65535))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberServiceTypeConfiguration(
					hostName, "IPV4", "VNIOS", "ALL_V4", networkSettingAddress6,
					vipAddress4, "172.28.38.1", "255.255.254.0",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service_type_configuration", "ALL_V4"),
				),
			},
			{
				Config: testAccMemberServiceTypeConfiguration(
					hostName, "IPV6", "VNIOS", "ALL_V6", networkSettingAddress6,
					vipAddress4, "2001::1", "64",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service_type_configuration", "ALL_V6"),
				),
			},
		},
	})
}

func TestAccMemberResource_SnmpSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_snmp_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.140"

	snmpSettingVal := map[string]any{
		"queries_enable":           true,
		"queries_community_string": "example_community_string",
		"snmpv3_queries_enable":    true,
		"snmpv3_traps_enable":      true,
		"traps_enable":             false,
		"snmpv3_queries_users": []map[string]any{
			{
				"user": "${nios_security_snmp_user.test.ref}",
			},
		},
	}

	snmpSettingValUpdated := map[string]any{
		"queries_enable":        false,
		"snmpv3_queries_enable": false,
		"snmpv3_traps_enable":   false,
		"traps_enable":          false,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberSnmpSetting(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					snmpSettingVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "snmp_setting.queries_enable", "true"),
					resource.TestCheckResourceAttr(resourceName, "snmp_setting.queries_community_string", "example_community_string"),
					resource.TestCheckResourceAttr(resourceName, "snmp_setting.snmpv3_queries_enable", "true"),
					resource.TestCheckResourceAttr(resourceName, "snmp_setting.snmpv3_traps_enable", "true"),
					resource.TestCheckResourceAttr(resourceName, "snmp_setting.traps_enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "snmp_setting.snmpv3_queries_users.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "use_snmp_setting", "true"),
				),
			},
			{
				Config: testAccMemberSnmpSetting(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					snmpSettingValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "snmp_setting.queries_enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "snmp_setting.snmpv3_queries_enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "snmp_setting.snmpv3_traps_enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "snmp_setting.traps_enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_snmp_setting", "true"),
				),
			},
		},
	})
}

func TestAccMemberResource_StaticRoutes(t *testing.T) {
	t.Skip("Cannot create Static Routes if Static Routes already exist")
	var resourceName = "nios_grid_member.test_static_routes"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.141"

	staticRoutesVal := []map[string]any{
		{
			"address":     "172.28.90.10",
			"gateway":     "172.28.90.1",
			"subnet_mask": "255.255.254.0",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberStaticRoutes(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					staticRoutesVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "static_routes.#", "1"),
				),
			},
		},
	})
}

func TestAccMemberResource_SupportAccessEnable(t *testing.T) {
	var resourceName = "nios_grid_member.test_support_access_enable"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.142"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberSupportAccessEnable(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "support_access_enable", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_support_access_enable", "true"),
				),
			},
			{
				Config: testAccMemberSupportAccessEnable(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "support_access_enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_support_access_enable", "true"),
				),
			},
		},
	})
}

func TestAccMemberResource_SyslogProxySetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_syslog_proxy_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.249"
	testDataPath := getTestDataPath()
	syslogServersVal := []map[string]any{
		{
			"address_or_fqdn":       "19245.com",
			"category_list":         []string{"AUTH_ACTIVE_DIRECTORY"},
			"certificate_file_path": filepath.Join(testDataPath, "client.crt"),
			"connection_type":       "STCP",
			"local_interface":       "ANY",
			"message_node_id":       "LAN",
			"message_source":        "ANY",
			"only_category_list":    false,
			"port":                  514,
			"severity":              "DEBUG",
			"username":              "admin1",
			"password":              "Password123!",
		},
	}

	syslogProxySettingVal := map[string]any{
		"enable":     false,
		"tcp_enable": false,
		"tcp_port":   514,
		"udp_enable": true,
		"udp_port":   514,
	}

	syslogProxySettingValUpdated := map[string]any{
		"client_acls": []map[string]any{
			{
				"struct":     "addressac",
				"address":    "19.0.0.1",
				"permission": "ALLOW",
			},
		},
		"enable":     true,
		"tcp_enable": true,
		"tcp_port":   1514,
		"udp_enable": false,
		"udp_port":   514,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberSyslogProxySetting(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.13", "255.255.254.0",
					syslogProxySettingVal,
					syslogServersVal, syslogServersVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.tcp_enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.tcp_port", "514"),
					resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.udp_enable", "true"),
					resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.udp_port", "514"),
					resource.TestCheckResourceAttr(resourceName, "use_syslog_proxy_setting", "true"),
				),
			},
			{
				Config: testAccMemberSyslogProxySetting(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					syslogProxySettingValUpdated,
					syslogServersVal, syslogServersVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.client_acls.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.client_acls.0.struct", "addressac"),
					resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.client_acls.0.address", "19.0.0.1"),
					resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.client_acls.0.permission", "ALLOW"),
					resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.enable", "true"),
					resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.tcp_enable", "true"),
					resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.tcp_port", "1514"),
					resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.udp_enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.udp_port", "514"),
					resource.TestCheckResourceAttr(resourceName, "use_syslog_proxy_setting", "true"),
				),
			},
		},
	})
}

func TestAccMemberResource_SyslogServers(t *testing.T) {
	var resourceName = "nios_grid_member.test_syslog_servers"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.144"
	testDataPath := getTestDataPath()

	syslogServersVal := []map[string]any{
		{
			"address_or_fqdn":       "192.com",
			"category_list":         []string{"AUTH_ACTIVE_DIRECTORY"},
			"certificate_file_path": filepath.Join(testDataPath, "client.crt"),
			"connection_type":       "STCP",
			"local_interface":       "ANY",
			"message_node_id":       "LAN",
			"message_source":        "ANY",
			"only_category_list":    false,
			"port":                  514,
			"severity":              "DEBUG",
		},
		{
			"address_or_fqdn":    "abc.example.com",
			"category_list":      []string{"AUTH_ACTIVE_DIRECTORY"},
			"connection_type":    "TCP",
			"local_interface":    "ANY",
			"message_node_id":    "LAN",
			"message_source":     "EXTERNAL",
			"only_category_list": false,
			"port":               515,
			"severity":           "INFO",
		},
	}

	syslogServersValUpdated := []map[string]any{
		{
			"address_or_fqdn":    "192.com",
			"category_list":      []string{"AUTH_ACTIVE_DIRECTORY"},
			"connection_type":    "TCP",
			"local_interface":    "ANY",
			"message_node_id":    "LAN",
			"message_source":     "ANY",
			"only_category_list": false,
			"port":               515,
			"severity":           "INFO",
		},
	}
	syslogProxySettingVal := map[string]any{

		"client_acls": []map[string]any{
			{
				"struct":     "addressac",
				"address":    "192.0.0.1",
				"permission": "ALLOW",
			},
		},
		"enable":     true,
		"tcp_enable": false,
		"tcp_port":   514,
		"udp_enable": true,
		"udp_port":   514,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberSyslogServers(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					syslogServersVal,
					syslogProxySettingVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.address_or_fqdn", "192.com"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.category_list.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.category_list.0", "AUTH_ACTIVE_DIRECTORY"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.connection_type", "STCP"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.local_interface", "ANY"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.message_node_id", "LAN"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.message_source", "ANY"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.only_category_list", "false"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.port", "514"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.severity", "DEBUG"),
				),
			},
			{
				Config: testAccMemberSyslogServers(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					syslogServersValUpdated,
					syslogProxySettingVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.address_or_fqdn", "192.com"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.category_list.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.category_list.0", "AUTH_ACTIVE_DIRECTORY"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.connection_type", "TCP"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.local_interface", "ANY"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.message_node_id", "LAN"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.message_source", "ANY"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.only_category_list", "false"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.port", "515"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.severity", "INFO"),
				),
			},
		},
	})
}

func TestAccMemberResource_SyslogSize(t *testing.T) {
	var resourceName = "nios_grid_member.test_syslog_size"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.145"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberSyslogSize(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					10,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "syslog_size", "10"),
				),
			},
			{
				Config: testAccMemberSyslogSize(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					300,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "syslog_size", "300"),
				),
			},
		},
	})
}

func TestAccMemberResource_ThresholdTraps(t *testing.T) {
	var resourceName = "nios_grid_member.test_threshold_traps"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.146"

	thresholdTrapsVal := []map[string]any{
		{
			"trap_reset":   10,
			"trap_trigger": 100,
			"trap_type":    "CpuUsage",
		},
	}
	thresholdTrapsValUpdated := []map[string]any{
		{
			"trap_reset":   15,
			"trap_trigger": 150,
			"trap_type":    "CpuUsage",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberThresholdTraps(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					thresholdTrapsVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "threshold_traps.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "threshold_traps.0.trap_type", "CpuUsage"),
					resource.TestCheckResourceAttr(resourceName, "threshold_traps.0.trap_trigger", "100"),
					resource.TestCheckResourceAttr(resourceName, "threshold_traps.0.trap_reset", "10"),
					resource.TestCheckResourceAttr(resourceName, "use_threshold_traps", "true"),
				),
			},
			{
				Config: testAccMemberThresholdTraps(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					thresholdTrapsValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "threshold_traps.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "threshold_traps.0.trap_type", "CpuUsage"),
					resource.TestCheckResourceAttr(resourceName, "threshold_traps.0.trap_trigger", "150"),
					resource.TestCheckResourceAttr(resourceName, "threshold_traps.0.trap_reset", "15"),
					resource.TestCheckResourceAttr(resourceName, "use_threshold_traps", "true"),
				),
			},
		},
	})
}

func TestAccMemberResource_TimeZone(t *testing.T) {
	var resourceName = "nios_grid_member.test_time_zone"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.147"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberTimeZone(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					"UTC",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "time_zone", "UTC"),
				),
			},
			{
				Config: testAccMemberTimeZone(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					"UTC",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "time_zone", "UTC"),
				),
			},
		},
	})
}

func TestAccMemberResource_TrafficCaptureAuthDnsSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_traffic_capture_auth_dns_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.148"

	trafficCaptureAuthDnsSettingVal := map[string]any{
		"auth_dns_latency_listen_on_source": "VIP_V4",
		"auth_dns_latency_trigger_enable":   false,
	}
	trafficCaptureAuthDnsSettingValUpdated := map[string]any{
		"auth_dns_latency_listen_on_source": "VIP_V4",
		"auth_dns_latency_trigger_enable":   true,
		"auth_dns_latency_threshold":        15,
		"auth_dns_latency_reset":            15,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberTrafficCaptureAuthDnsSetting(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					trafficCaptureAuthDnsSettingVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_auth_dns_setting.auth_dns_latency_listen_on_source", "VIP_V4"),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_auth_dns_setting.auth_dns_latency_trigger_enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_auth_dns", "true"),
				),
			},
			{
				Config: testAccMemberTrafficCaptureAuthDnsSetting(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					trafficCaptureAuthDnsSettingValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_auth_dns_setting.auth_dns_latency_listen_on_source", "VIP_V4"),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_auth_dns_setting.auth_dns_latency_trigger_enable", "true"),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_auth_dns_setting.auth_dns_latency_threshold", "15"),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_auth_dns_setting.auth_dns_latency_reset", "15"),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_auth_dns", "true"),
				),
			},
		},
	})
}

func TestAccMemberResource_TrafficCaptureChrSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_traffic_capture_chr_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.149"

	trafficCaptureChrSettingVal := map[string]any{
		"chr_trigger_enable": false,
	}
	trafficCaptureChrSettingValUpdated := map[string]any{
		"chr_trigger_enable":        true,
		"chr_threshold":             15,
		"chr_reset":                 15,
		"chr_min_cache_utilization": 15,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberTrafficCaptureChrSetting(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					trafficCaptureChrSettingVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_chr_setting.chr_trigger_enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_chr", "true"),
				),
			},
			{
				Config: testAccMemberTrafficCaptureChrSetting(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					trafficCaptureChrSettingValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_chr_setting.chr_trigger_enable", "true"),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_chr_setting.chr_threshold", "15"),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_chr_setting.chr_reset", "15"),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_chr_setting.chr_min_cache_utilization", "15"),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_chr", "true"),
				),
			},
		},
	})
}

func TestAccMemberResource_TrafficCaptureQpsSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_traffic_capture_qps_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.150"

	trafficCaptureQpsSettingVal := map[string]any{
		"qps_trigger_enable": false,
	}
	trafficCaptureQpsSettingValUpdated := map[string]any{
		"qps_trigger_enable": true,
		"qps_threshold":      15,
		"qps_reset":          15,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberTrafficCaptureQpsSetting(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					trafficCaptureQpsSettingVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_qps_setting.qps_trigger_enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_qps", "true"),
				),
			},
			{
				Config: testAccMemberTrafficCaptureQpsSetting(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					trafficCaptureQpsSettingValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_qps_setting.qps_trigger_enable", "true"),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_qps_setting.qps_threshold", "15"),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_qps_setting.qps_reset", "15"),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_qps", "true"),
				),
			},
		},
	})
}

func TestAccMemberResource_TrafficCaptureRecDnsSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_traffic_capture_rec_dns_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.151"
	networkSettingAddress6 := fmt.Sprintf("2001:db8:%x:%x::%x", acctest.RandomNumber(65535), acctest.RandomNumber(65535), acctest.RandomNumber(65535))

	trafficCaptureRecDnsSettingVal := map[string]any{
		"rec_dns_latency_listen_on_source": "VIP_V4",
		"rec_dns_latency_trigger_enable":   false,
	}
	trafficCaptureRecDnsSettingValUpdated := map[string]any{
		"rec_dns_latency_listen_on_source": "VIP_V6",
		"rec_dns_latency_trigger_enable":   false,
		"rec_dns_latency_reset":            20,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberTrafficCaptureRecDnsSetting(
					hostName, "IPV4", "VNIOS", "ALL_V4", networkSettingAddress6,
					vipAddress, "172.28.38.1", "255.255.254.0",
					trafficCaptureRecDnsSettingVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_rec_dns_setting.rec_dns_latency_listen_on_source", "VIP_V4"),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_rec_dns_setting.rec_dns_latency_trigger_enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_rec_dns", "true"),
				),
			},
			{
				Config: testAccMemberTrafficCaptureRecDnsSetting(
					hostName, "IPV6", "VNIOS", "ALL_V6", networkSettingAddress6,
					vipAddress, "172.28.38.1", "255.255.254.0",
					trafficCaptureRecDnsSettingValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_rec_dns_setting.rec_dns_latency_listen_on_source", "VIP_V6"),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_rec_dns_setting.rec_dns_latency_trigger_enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_rec_dns_setting.rec_dns_latency_reset", "20"),
				),
			},
		},
	})
}

func TestAccMemberResource_TrafficCaptureRecQueriesSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_traffic_capture_rec_queries_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.152"

	trafficCaptureRecQueriesSettingVal := map[string]any{
		"recursive_clients_count_trigger_enable": false,
	}
	trafficCaptureRecQueriesSettingValUpdated := map[string]any{
		"recursive_clients_count_trigger_enable": true,
		"recursive_clients_count_threshold":      15,
		"recursive_clients_count_reset":          15,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberTrafficCaptureRecQueriesSetting(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					trafficCaptureRecQueriesSettingVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_rec_queries_setting.recursive_clients_count_trigger_enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_rec_queries", "true"),
				),
			},
			{
				Config: testAccMemberTrafficCaptureRecQueriesSetting(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					trafficCaptureRecQueriesSettingValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_rec_queries_setting.recursive_clients_count_trigger_enable", "true"),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_rec_queries_setting.recursive_clients_count_threshold", "15"),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_rec_queries_setting.recursive_clients_count_reset", "15"),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_rec_queries", "true"),
				),
			},
		},
	})
}

func TestAccMemberResource_TrapNotifications(t *testing.T) {
	var resourceName = "nios_grid_member.test_trap_notifications"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.153"

	trapNotificationsVal := []map[string]any{
		{"enable_email": false, "enable_trap": true, "trap_type": "AnalyticsRPZ"},
		{"enable_email": false, "enable_trap": true, "trap_type": "DNS"},
	}
	trapNotificationsValUpdated := []map[string]any{
		{"enable_email": true, "enable_trap": true, "trap_type": "AnalyticsRPZ"},
		{"enable_email": false, "enable_trap": true, "trap_type": "DNS"},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberTrapNotifications(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					trapNotificationsVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrSet(resourceName, "trap_notifications.#"),
					resource.TestCheckResourceAttr(resourceName, "use_trap_notifications", "true"),
				),
			},
			{
				Config: testAccMemberTrapNotifications(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.38.1", "255.255.254.0",
					trapNotificationsValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrSet(resourceName, "trap_notifications.#"),
					resource.TestCheckResourceAttr(resourceName, "use_trap_notifications", "true"),
				),
			},
		},
	})
}

func TestAccMemberResource_UpgradeGroup(t *testing.T) {
	var resourceName = "nios_grid_member.test_upgrade_group"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.247"
	upgradeGroupVal := fmt.Sprintf("new-group-%s", hostName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberUpgradeGroup(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					"Default",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "upgrade_group", "Default"),
				),
			},
			{
				Config: testAccMemberUpgradeGroup(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					upgradeGroupVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "upgrade_group", upgradeGroupVal),
				),
			},
		},
	})
}

func TestAccMemberResource_UseAutomatedTrafficCapture(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_automated_traffic_capture"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.155"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberUseAutomatedTrafficCapture(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_automated_traffic_capture", "true"),
				),
			},
			{
				Config: testAccMemberUseAutomatedTrafficCapture(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_automated_traffic_capture", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_UseDnsResolverSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_dns_resolver_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.156"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberUseDnsResolverSetting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dns_resolver_setting", "true"),
				),
			},
			{
				Config: testAccMemberUseDnsResolverSetting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dns_resolver_setting", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_UseDscp(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_dscp"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.157"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberUseDscp(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dscp", "true"),
				),
			},
			{
				Config: testAccMemberUseDscp(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dscp", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_UseEmailSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_email_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.158"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberUseEmailSetting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_email_setting", "true"),
				),
			},
			{
				Config: testAccMemberUseEmailSetting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_email_setting", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_UseEnableLom(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_enable_lom"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.159"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberUseEnableLom(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_lom", "true"),
				),
			},
			{
				Config: testAccMemberUseEnableLom(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_lom", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_UseEnableMemberRedirect(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_enable_member_redirect"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.160"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberUseEnableMemberRedirect(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_member_redirect", "true"),
				),
			},
			{
				Config: testAccMemberUseEnableMemberRedirect(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_member_redirect", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_UseExternalSyslogBackupServers(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_external_syslog_backup_servers"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.161"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberUseExternalSyslogBackupServers(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_external_syslog_backup_servers", "true"),
				),
			},
			{
				Config: testAccMemberUseExternalSyslogBackupServers(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_external_syslog_backup_servers", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_UseRemoteConsoleAccessEnable(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_remote_console_access_enable"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.162"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberUseRemoteConsoleAccessEnable(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_remote_console_access_enable", "true"),
				),
			},
			{
				Config: testAccMemberUseRemoteConsoleAccessEnable(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_remote_console_access_enable", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_UseSnmpSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_snmp_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.163"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberUseSnmpSetting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp_setting", "true"),
				),
			},
			{
				Config: testAccMemberUseSnmpSetting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp_setting", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_UseSupportAccessEnable(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_support_access_enable"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.164"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberUseSupportAccessEnable(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_support_access_enable", "true"),
				),
			},
			{
				Config: testAccMemberUseSupportAccessEnable(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_support_access_enable", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_UseSyslogProxySetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_syslog_proxy_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.165"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberUseSyslogProxySetting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_syslog_proxy_setting", "true"),
				),
			},
			{
				Config: testAccMemberUseSyslogProxySetting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_syslog_proxy_setting", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_UseThresholdTraps(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_threshold_traps"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.166"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberUseThresholdTraps(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_threshold_traps", "true"),
				),
			},
			{
				Config: testAccMemberUseThresholdTraps(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_threshold_traps", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_UseTimeZone(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_time_zone"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.167"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberUseTimeZone(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_time_zone", "true"),
				),
			},
			{
				Config: testAccMemberUseTimeZone(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_time_zone", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_UseTrafficCaptureAuthDns(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_traffic_capture_auth_dns"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.168"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberUseTrafficCaptureAuthDns(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_auth_dns", "true"),
				),
			},
			{
				Config: testAccMemberUseTrafficCaptureAuthDns(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_auth_dns", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_UseTrafficCaptureChr(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_traffic_capture_chr"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.169"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberUseTrafficCaptureChr(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_chr", "true"),
				),
			},
			{
				Config: testAccMemberUseTrafficCaptureChr(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_chr", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_UseTrafficCaptureQps(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_traffic_capture_qps"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.170"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberUseTrafficCaptureQps(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_qps", "true"),
				),
			},
			{
				Config: testAccMemberUseTrafficCaptureQps(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_qps", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_UseTrafficCaptureRecDns(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_traffic_capture_rec_dns"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.171"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberUseTrafficCaptureRecDns(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_rec_dns", "true"),
				),
			},
			{
				Config: testAccMemberUseTrafficCaptureRecDns(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_rec_dns", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_UseTrafficCaptureRecQueries(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_traffic_capture_rec_queries"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.172"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberUseTrafficCaptureRecQueries(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_rec_queries", "true"),
				),
			},
			{
				Config: testAccMemberUseTrafficCaptureRecQueries(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_rec_queries", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_UseTrapNotifications(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_trap_notifications"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.173"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberUseTrapNotifications(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_trap_notifications", "true"),
				),
			},
			{
				Config: testAccMemberUseTrapNotifications(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_trap_notifications", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_UseV4Vrrp(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_v4_vrrp"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.174"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberUseV4Vrrp(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_v4_vrrp", "true"),
				),
			},
			{
				Config: testAccMemberUseV4Vrrp(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_v4_vrrp", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_MasterCandidate(t *testing.T) {
	var resourceName = "nios_grid_member.test_master_candidate"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.175"
	networkSettingAddress6 := fmt.Sprintf("2001:db8:%x:%x::%x", acctest.RandomNumber(65535), acctest.RandomNumber(65535), acctest.RandomNumber(65535))
	ipv6SettingVal := map[string]any{
		"auto_router_config_enabled": false,
		"dscp":                       0,
		"enabled":                    true,
		"use_dscp":                   false,
		"virtual_ip":                 networkSettingAddress6,
		"gateway":                    "2001::1",
		"primary":                    true,
		"cidr_prefix":                8,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberMasterCandidate(
					hostName,
					"BOTH",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					"true",
					ipv6SettingVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "master_candidate", "true"),
				),
			},
			{
				Config: testAccMemberMasterCandidate(
					hostName,
					"BOTH",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					"false",
					ipv6SettingVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "master_candidate", "false"),
				),
			},
		},
	})
}

func TestAccMemberResource_VipSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_vip_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.176"

	vipSettingVal := map[string]any{
		"address":     vipAddress,
		"dscp":        0,
		"gateway":     "172.28.38.1",
		"primary":     true,
		"subnet_mask": "255.255.254.0",
		"use_dscp":    false,
	}
	vipSettingValUpdated := map[string]any{
		"address":     vipAddress,
		"dscp":        0,
		"gateway":     "172.28.38.2",
		"primary":     true,
		"subnet_mask": "255.255.254.0",
		"use_dscp":    false,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberVipSetting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipSettingVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vip_setting.address", vipAddress),
					resource.TestCheckResourceAttr(resourceName, "vip_setting.dscp", "0"),
					resource.TestCheckResourceAttr(resourceName, "vip_setting.use_dscp", "false"),
				),
			},
			{
				Config: testAccMemberVipSetting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipSettingValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vip_setting.address", vipAddress),
					resource.TestCheckResourceAttr(resourceName, "vip_setting.gateway", "172.28.38.2"),
				),
			},
		},
	})
}

func TestAccMemberResource_VpnMtu(t *testing.T) {
	var resourceName = "nios_grid_member.test_vpn_mtu"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.177"
	vpnMtu1 := 1450
	vpnMtu2 := 1400

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberVpnMtu(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					vpnMtu1,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vpn_mtu", "1450"),
				),
			},
			{
				Config: testAccMemberVpnMtu(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
					vpnMtu2,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vpn_mtu", "1400"),
				),
			},
		},
	})
}

func TestAccMemberResource_Import(t *testing.T) {
	var resourceName = "nios_grid_member.test"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := "172.28.38.178"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberBasicConfig(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.38.1",
					"255.255.254.0",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
				),
			},
			// Import with PlanOnly to detect differences
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccMemberImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"configure_csp_member_setting", "support_access_info"},
				ImportStateVerifyIdentifierAttribute: "uuid",
				PlanOnly:                             true,
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccMemberImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all", "configure_csp_member_setting", "support_access_info"},
				ImportStateVerifyIdentifierAttribute: "uuid",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckMemberExists(ctx context.Context, resourceName string, v *grid.Member) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.GridAPI.
			MemberAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForMember).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetMemberResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetMemberResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckMemberDestroy(ctx context.Context, v *grid.Member) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.GridAPI.
			MemberAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForMember).
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

func testAccCheckMemberDisappears(ctx context.Context, v *grid.Member) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.GridAPI.
			MemberAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccMemberImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.Attributes["uuid"] == "" {
			return "", fmt.Errorf("uuid is not set")
		}
		return rs.Primary.Attributes["uuid"], nil
	}
}

func testAccMemberBasicConfig(hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test" {
	host_name = %q
	config_addr_type = %q
	platform = %q
	service_type_configuration = %q
    
	ipv6_setting = {
		auto_router_config_enabled = false
		dscp = 0
		enabled = false
		primary = true
		use_dscp = false
	}
    
	vip_setting = {
		address = %q
		dscp = 0
		gateway = %q
		primary = true
		subnet_mask = %q
		use_dscp = false
	}
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask)
}

func testAccMemberAdditionalIpList(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	additionalIpList []map[string]any,
	bgpAs []map[string]any,
	ospfList []map[string]any,
) string {

	additionalIpListStr := utils.ConvertSliceOfMapsToHCL(additionalIpList)

	bgpAsStr := "null"
	if len(bgpAs) > 0 {
		bgpAsStr = utils.ConvertSliceOfMapsToHCL(bgpAs)
	}

	ospfListStr := "null"
	if len(ospfList) > 0 {
		ospfListStr = utils.ConvertSliceOfMapsToHCL(ospfList)
	}

	return fmt.Sprintf(`
resource "nios_grid_member" "test_additional_ip_list" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    use_dscp = false
    additional_ip_list = %s
    bgp_as = %s
    ospf_list = %s
}
`, hostName, configAddrType, platform, serviceTypeConfig,
		vipAddress, vipGateway, vipSubnetMask,
		additionalIpListStr, bgpAsStr, ospfListStr)
}

func testAccMemberAutomatedTrafficCaptureSetting(hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string, automatedTrafficCaptureSetting map[string]any) string {
	automatedTrafficCaptureSettingStr := utils.ConvertMapToHCL(automatedTrafficCaptureSetting)

	return fmt.Sprintf(`
resource "nios_grid_member" "test_automated_traffic_capture_setting" {
    host_name = %q
	config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }
    automated_traffic_capture_setting = %s
    use_automated_traffic_capture = true
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, automatedTrafficCaptureSettingStr)
}

func testAccMemberBgpAs(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	bgpAs []map[string]any,
) string {
	bgpAsStr := "null"
	if len(bgpAs) > 0 {
		bgpAsStr = utils.ConvertSliceOfMapsToHCL(bgpAs)
	}

	return fmt.Sprintf(`
resource "nios_grid_member" "test_bgp_as" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    bgp_as = %s
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, bgpAsStr)
}

func testAccMemberComment(hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask string, comment string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_comment" {
    host_name = %q
	config_addr_type = %q
	platform = %q
	service_type_configuration = %q
	
	ipv6_setting = {
		auto_router_config_enabled = false
		dscp = 0
		enabled = false
		primary = true
		use_dscp = false
	}

	vip_setting = {
		address = %q
		dscp = 0
		gateway = %q
		primary = true
		subnet_mask = %q
		use_dscp = false
	}
    comment = %q
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, comment)
}

func testAccMemberConfigAddrTypeIPv4(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string, useV4Vrrp bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_config_addr_type" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }
	use_v4_vrrp = %t // set use_v4_vrrp to false to avoid dependency on vrrp resource for config_addr_type testing
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useV4Vrrp)
}

func testAccMemberConfigAddrTypeIPv6(
	hostName, configAddrType, platform, serviceTypeConfig,
	ipv6SettingAddress, ipv6SettingGateway string, ipv6SettingCIDRPrefix int, useV4Vrrp bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_config_addr_type" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q
    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = true
        primary = true
        use_dscp = false
        virtual_ip = %q
        gateway = %q
        cidr_prefix = %d
    }

    use_v4_vrrp = %t 
}
`, hostName, configAddrType, platform, serviceTypeConfig, ipv6SettingAddress, ipv6SettingGateway, ipv6SettingCIDRPrefix, useV4Vrrp)
}

func testAccMemberCspAccessKey(hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask string, cspAccessKey []string) string {
	cspAccessKeyStr := utils.ConvertStringSliceToHCL(cspAccessKey)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_csp_access_key" {
    host_name = %q
	config_addr_type = %q
    platform = %q
    service_type_configuration = %q
    
    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }
    
    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }
    csp_access_key = %s
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, cspAccessKeyStr)
}

func testAccMemberCspMemberSetting(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	cspMemberSetting map[string]any,
) string {
	cspMemberSettingStr := utils.ConvertMapToHCL(cspMemberSetting)

	return fmt.Sprintf(`
resource "nios_grid_member" "test_csp_member_setting" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    csp_member_setting = %s
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, cspMemberSettingStr)
}

func testAccMemberDnsResolverSetting(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	dnsResolverSetting map[string]any,
	useDNSResolverSetting bool,
) string {
	dnsResolverSettingStr := utils.ConvertMapToHCL(dnsResolverSetting)

	return fmt.Sprintf(`
resource "nios_grid_member" "test_dns_resolver_setting" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    dns_resolver_setting = %s
    use_dns_resolver_setting = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, dnsResolverSettingStr, useDNSResolverSetting)
}

func testAccMemberDscp(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	dscp int, useDSCP bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_dscp" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    dscp = %d
    use_dscp = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, dscp, useDSCP)
}

func testAccMemberEmailSetting(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	emailSetting map[string]any,
	useEmailSetting bool,
) string {
	emailSettingStr := utils.ConvertMapToHCL(emailSetting)

	return fmt.Sprintf(`
resource "nios_grid_member" "test_email_setting" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    email_setting = %s
    use_email_setting = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, emailSettingStr, useEmailSetting)
}

func testAccMemberEnableHa(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	enableHa bool,
	routerID int,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_enable_ha" {
	host_name = %q
	config_addr_type = %q
	platform = %q
	service_type_configuration = %q

	ipv6_setting = {
		auto_router_config_enabled = false
		enabled = false
		primary = true
		use_dscp = false
		dscp = 0
	}

	vip_setting = {
		address = %q
		gateway = %q
		subnet_mask = %q
		primary = true
		use_dscp = false
		dscp = 0
	}
	enable_ha = %t
	router_id = %d

	node_info = [
	{
	  lan_ha_port_setting = {
		ha_cloud_attribute = "UNK"
		ha_ip_address = "172.28.38.11"
		ha_port_setting = {
			auto_port_setting_enabled = true
			speed = "10"
		}
		lan_port_setting = {
			auto_port_setting_enabled = true
		}
		mgmt_lan = "172.28.38.32"
	  }
	},
    {
      lan_ha_port_setting = {
      	ha_cloud_attribute = "UNK"
      	ha_ip_address = "172.28.38.41"
      	ha_port_setting = {
        	auto_port_setting_enabled = true
        	speed = "10"
        }
      	lan_port_setting = {
        	auto_port_setting_enabled = true
      	}
      	mgmt_lan = "172.28.38.43"
      }
    }
  ]
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, enableHa, routerID)
}

func testAccMemberEnableLom(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	enableLom bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_enable_lom" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    enable_lom = %t
    use_enable_lom = true
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, enableLom)
}

func testAccMemberEnableMemberRedirect(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	enableMemberRedirect bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_enable_member_redirect" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    enable_member_redirect = %t
    use_enable_member_redirect = true
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, enableMemberRedirect)
}

func testAccMemberEnableRoApiAccess(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	enableRoApiAccess bool, ipv6Setting map[string]any,
) string {
	ipv6SettingStr := utils.ConvertMapToHCL(ipv6Setting)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_enable_ro_api_access" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    enable_ro_api_access = %t
	master_candidate = true
	ipv6_setting = %s
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, enableRoApiAccess, ipv6SettingStr)
}

func testAccMemberExtAttrs(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	extAttrs map[string]any,
) string {
	extAttrsStr := utils.ConvertMapToHCL(extAttrs)

	return fmt.Sprintf(`
resource "nios_grid_member" "test_extattrs" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    extattrs = %s
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, extAttrsStr)
}

func testAccMemberExternalSyslogBackupServers(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	externalSyslogBackupServers []map[string]any,
) string {
	externalSyslogBackupServersStr := utils.ConvertSliceOfMapsToHCL(externalSyslogBackupServers)

	return fmt.Sprintf(`
resource "nios_grid_member" "test_external_syslog_backup_servers" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    external_syslog_backup_servers = %s
    use_external_syslog_backup_servers = true
	external_syslog_server_enable = false
    use_syslog_proxy_setting = true
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, externalSyslogBackupServersStr)
}

func testAccMemberHaCloudPlatform(hostName string, haCloudPlatform string, vipAddress, vipGateway, vipSubnetMask string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_ha_cloud_platform" {
    host_name = %q
    ha_cloud_platform = %q
	vip_setting = {
		address = %q
		dscp = 0
		gateway = %q
		primary = true
		subnet_mask = %q
		use_dscp = false
	}
}
`, hostName, haCloudPlatform, vipAddress, vipGateway, vipSubnetMask)
}

func testAccMemberHaOnCloud(hostName string, haOnCloud string, haCloudPlatform, vipAddress, vipGateway, vipSubnetMask, enableHA string, routerID int) string {
	routerIDStr := ""
	haNodeInfoStr := ""
	haCloudPlatformStr := ""
	haOnCloudStr := fmt.Sprintf("ha_on_cloud = %q", haOnCloud)
	enableHAStr := fmt.Sprintf("enable_ha = %q", enableHA)
	if enableHA == "true" {
		routerIDStr = fmt.Sprintf("router_id = %d", routerID)
		haNodeInfoStr = `node_info = [
			{
				lan_ha_port_setting = {
					ha_cloud_attribute = "1"
					ha_ip_address      = "172.28.38.230"
					mgmt_lan           = "172.28.38.231"
				}
			},
			{
				lan_ha_port_setting = {
					ha_cloud_attribute = "2"
					ha_ip_address      = "172.28.38.232"
					mgmt_lan           = "172.28.38.233"
				}
			}
		]`
		haCloudPlatformStr = fmt.Sprintf("ha_cloud_platform = %q", haCloudPlatform)
	}
	return fmt.Sprintf(`
resource "nios_grid_member" "test_ha_on_cloud" {
    host_name = %q
    %s
	platform = "VNIOS"
	%s
	vip_setting = {
		address = %q
		dscp = 0
		gateway = %q
		primary = true
		subnet_mask = %q
		use_dscp = false
	}
	%s
	%s
	%s
	dns_resolver_setting = {
		search_domains = ["5.5.5.5"]
	}
	use_dns_resolver_setting = true
}
`, hostName, haOnCloudStr, haCloudPlatformStr, vipAddress, vipGateway, vipSubnetMask, haNodeInfoStr, enableHAStr, routerIDStr)
}

func testAccMemberHostName(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_host_name" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask)
}

func testAccMemberIpv6Setting(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	ipv6Setting map[string]any,
) string {
	ipv6SettingStr := utils.ConvertMapToHCL(ipv6Setting)

	return fmt.Sprintf(`
resource "nios_grid_member" "test_ipv6_setting" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = %s

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }
}
`, hostName, configAddrType, platform, serviceTypeConfig, ipv6SettingStr, vipAddress, vipGateway, vipSubnetMask)
}

func testAccMemberIpv6StaticRoutes(hostName string, ipv6StaticRoutes []map[string]any, vipAddress, vipGateway, vipSubnetMask string) string {
	ipv6StaticRoutesStr := utils.ConvertSliceOfMapsToHCL(ipv6StaticRoutes)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_ipv6_static_routes" {
    host_name = %q
    ipv6_static_routes = %s
	vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }
}
`, hostName, ipv6StaticRoutesStr, vipAddress, vipGateway, vipSubnetMask)
}

func testAccMemberLan2Enabled(hostName string, lan2Enabled string, vipAddress, vipGateway, vipSubnetMask string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_lan2_enabled" {
    host_name = %[1]q
    lan2_enabled = %[2]q
	vip_setting = {
        address = %[3]q
        dscp = 0
        gateway = %[4]q
        primary = true
        subnet_mask = %[5]q
        use_dscp = false
    }
	lan2_port_setting = {
		virtual_router_id = 10
		enabled = true
		network_setting = {
        address = "172.29.38.15"
        gateway = "172.29.38.1"
        primary = true
        subnet_mask = "255.255.0.0"
    }
	lan2_port_setting = null
}
}
`, hostName, lan2Enabled, vipAddress, vipGateway, vipSubnetMask)
}

func testAccMemberLan2EnabledFalse(hostName string, lan2Enabled string, vipAddress, vipGateway, vipSubnetMask string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_lan2_enabled" {
    host_name = %[1]q
    lan2_enabled = %[2]q
	vip_setting = {
        address = %[3]q
        dscp = 0
        gateway = %[4]q
        primary = true
        subnet_mask = %[5]q
        use_dscp = false
    }
}
`, hostName, lan2Enabled, vipAddress, vipGateway, vipSubnetMask)
}

func testAccMemberLan2PortSetting(hostName string, lan2PortSetting map[string]any, vipAddress, vipGateway, vipSubnetMask string) string {
	lan2PortSettingStr := utils.ConvertMapToHCL(lan2PortSetting)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_lan2_port_setting" {
    host_name = %q
    lan2_port_setting = %s
	vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }
	lan2_enabled = true
}
`, hostName, lan2PortSettingStr, vipAddress, vipGateway, vipSubnetMask)
}

func testAccMemberLomNetworkConfig(hostName string, lomNetworkConfig []map[string]any, vipAddress, vipGateway, vipSubnetMask string) string {
	lomNetworkConfigStr := utils.ConvertSliceOfMapsToHCL(lomNetworkConfig)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_lom_network_config" {
    host_name = %q
    lom_network_config = %s
	vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }
}
`, hostName, lomNetworkConfigStr, vipAddress, vipGateway, vipSubnetMask)
}

func testAccMemberLomUsers(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	lomUsers []map[string]any,
) string {
	lomUsersStr := utils.ConvertSliceOfMapsToHCL(lomUsers)

	return fmt.Sprintf(`
resource "nios_grid_member" "test_lom_users" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    enable_lom = true
    use_enable_lom = true
    lom_users = %s
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, lomUsersStr)
}

func testAccMemberMasterCandidate(hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask, masterCandidate string, ipv6Setting map[string]any) string {
	ipv6SettingStr := utils.ConvertMapToHCL(ipv6Setting)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_master_candidate" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q
    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    master_candidate = %q
	ipv6_setting = %s
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, masterCandidate, ipv6SettingStr)
}

func testAccMemberMgmtPortSetting(hostName string, mgmtPortSetting map[string]any, vipAddress, vipGateway, vipSubnetMask string) string {
	mgmtPortSettingStr := utils.ConvertMapToHCL(mgmtPortSetting)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_mgmt_port_setting" {
    host_name = %q
    mgmt_port_setting = %s
	vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }
	node_info = [
	{
		mgmt_network_setting = {
			address     = "1.1.1.2"
			gateway     = "1.1.1.1"
			subnet_mask = "255.255.255.0"
		}
	}
	]
}
`, hostName, mgmtPortSettingStr, vipAddress, vipGateway, vipSubnetMask)
}

func testAccMemberNatSetting(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	natSetting map[string]any,
) string {
	natSettingStr := utils.ConvertMapToHCL(natSetting)

	return fmt.Sprintf(`
resource "nios_grid_member" "test_nat_setting" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    nat_setting = %s
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, natSettingStr)
}

func testAccMemberNodeInfo(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	enableHA string, routerID int, nodeInfo []map[string]any, mgmtPortSetting map[string]any,
) string {
	nodeInfoStr := utils.ConvertSliceOfMapsToHCL(nodeInfo)
	mgmtPortSettingStr := utils.ConvertMapToHCL(mgmtPortSetting)
	routerIDStr := ""
	if enableHA == "true" {
		routerIDStr = fmt.Sprintf("router_id = %d", routerID)
	}

	return fmt.Sprintf(`
resource "nios_grid_member" "test_node_info" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    enable_ha = %q
    %s
    node_info = %s
	mgmt_port_setting = %s
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, enableHA, routerIDStr, nodeInfoStr, mgmtPortSettingStr)
}

func testAccMemberNtpSetting(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	ntpSetting map[string]any,
) string {
	ntpSettingStr := utils.ConvertMapToHCL(ntpSetting)

	return fmt.Sprintf(`
resource "nios_grid_member" "test_ntp_setting" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    ntp_setting = %s
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, ntpSettingStr)
}

func testAccMemberOspfList(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	ospfList []map[string]any,
) string {
	ospfListStr := utils.ConvertSliceOfMapsToHCL(ospfList)

	return fmt.Sprintf(`
resource "nios_grid_member" "test_ospf_list" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    ospf_list = %s
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, ospfListStr)
}

func testAccMemberPassiveHaArpEnabled(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	passiveHaArpEnabled bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_passive_ha_arp_enabled" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    enable_ha = true
    router_id = 198
    passive_ha_arp_enabled = %t

    node_info = [
        {
            lan_ha_port_setting = {
                ha_cloud_attribute = "UNK"
                ha_ip_address = "172.28.38.11"
                ha_port_setting = {
                    auto_port_setting_enabled = true
                    speed = "10"
                }
                lan_port_setting = {
                    auto_port_setting_enabled = true
                }
                mgmt_lan = "172.28.38.32"
            }
        },
        {
            lan_ha_port_setting = {
                ha_cloud_attribute = "UNK"
                ha_ip_address = "172.28.38.41"
                ha_port_setting = {
                    auto_port_setting_enabled = true
                    speed = "10"
                }
                lan_port_setting = {
                    auto_port_setting_enabled = true
                }
                mgmt_lan = "172.28.38.43"
            }
        }
    ]
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, passiveHaArpEnabled)
}

func testAccMemberPlatform(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_platform" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask)
}

func testAccMemberPreProvisioning(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	preProvisioning map[string]any,
) string {
	preProvisioningStr := utils.ConvertMapToHCL(preProvisioning)
	preProvisioningConfigStr := fmt.Sprintf("\n\n    pre_provisioning = %s", preProvisioningStr)

	return fmt.Sprintf(`
resource "nios_grid_member" "test" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }
	%s
	preserve_if_owns_delegation  = false
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, preProvisioningConfigStr)
}

func testAccMemberPreserveIfOwnsDelegation(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	preserveIfOwnsDelegation bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_preserve_if_owns_delegation" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    preserve_if_owns_delegation = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, preserveIfOwnsDelegation)
}

func testAccMemberRemoteConsoleAccessEnable(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	remoteConsoleAccessEnable bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_remote_console_access_enable" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    remote_console_access_enable = %t
    use_remote_console_access_enable = true
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, remoteConsoleAccessEnable)
}

func testAccMemberRouterId(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	routerId int,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_router_id" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    router_id = %d
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, routerId)
}

func testAccMemberServiceTypeConfiguration(
	hostName, configAddrType, platform, serviceTypeConfiguration, ipv6SettingIP,
	vipAddress, vipGateway, vipSubnetMask string,
) string {
	enabledV6 := configAddrType != "IPV4"
	ipv6SettingStr := ""
	if configAddrType != "IPV4" {
		ipv6SettingStr = fmt.Sprintf("enabled = %t\n\t\tvirtual_ip = %q\n\t\tgateway = \"2001::1\"\n\t\tcidr_prefix = 8", enabledV6, ipv6SettingIP)
	}
	vipSettingStr := ""
	if configAddrType == "IPV4" {
		vipSettingStr = fmt.Sprintf("\t\taddress = %q\n\t\tgateway = %q\n\t\tsubnet_mask = %q", vipAddress, vipGateway, vipSubnetMask)
	}
	return fmt.Sprintf(`
resource "nios_grid_member" "test_service_type_configuration" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        %s
        primary = true
        use_dscp = false
    }

    vip_setting = {
        %s
        use_dscp = false
		primary = true
		dscp = 0
    }
}
`, hostName, configAddrType, platform, serviceTypeConfiguration, ipv6SettingStr, vipSettingStr)
}

func testAccMemberSnmpSetting(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	snmpSetting map[string]any,
) string {
	snmpSettingStr := utils.ConvertMapToHCL(snmpSetting)

	return fmt.Sprintf(`

resource "nios_security_snmp_user" "test" {
    name                 	= "example-snmpuser1"
    authentication_protocol = "NONE"
    privacy_protocol     	= "NONE"
}

resource "nios_grid_member" "test_snmp_setting" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    snmp_setting = %s
    use_snmp_setting = true
	
	
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, snmpSettingStr)
}

func testAccMemberStaticRoutes(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	staticRoutes []map[string]any,
) string {
	staticRoutesStr := utils.ConvertSliceOfMapsToHCL(staticRoutes)

	return fmt.Sprintf(`
resource "nios_grid_member" "test_static_routes" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    static_routes = %s
	use_dscp = true
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, staticRoutesStr)
}

func testAccMemberSupportAccessEnable(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	supportAccessEnable bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_support_access_enable" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    support_access_enable = %t
    use_support_access_enable = true
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, supportAccessEnable)
}

func testAccMemberSyslogProxySetting(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	syslogProxySetting map[string]any,
	syslogServersVal []map[string]any,
	externalSyslogBackupServers []map[string]any,
) string {
	syslogProxySettingStr := utils.ConvertMapToHCL(syslogProxySetting)
	syslogServersValStr := utils.ConvertSliceOfMapsToHCL(syslogServersVal)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_syslog_proxy_setting" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    syslog_proxy_setting = %s
    use_syslog_proxy_setting = true
	syslog_servers = %s
	external_syslog_server_enable = true // System requires a mandatory external Syslog server when syslog proxy setting is enabled.
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, syslogProxySettingStr, syslogServersValStr)
}

func testAccMemberSyslogServers(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	syslogServers []map[string]any, sysLogProxySetting map[string]any,
) string {
	syslogServersStr := utils.ConvertSliceOfMapsToHCL(syslogServers)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_syslog_servers" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    syslog_servers = %s
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, syslogServersStr)
}

func testAccMemberSyslogSize(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	syslogSize int,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_syslog_size" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    syslog_size = %d
    use_syslog_proxy_setting = true
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, syslogSize)
}

func testAccMemberThresholdTraps(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	thresholdTraps []map[string]any,
) string {
	thresholdTrapsStr := utils.ConvertSliceOfMapsToHCL(thresholdTraps)

	return fmt.Sprintf(`
resource "nios_grid_member" "test_threshold_traps" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    threshold_traps = %s
    use_threshold_traps = true
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, thresholdTrapsStr)
}

func testAccMemberTimeZone(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask, timeZone string,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_time_zone" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    time_zone = %q
    use_time_zone = true
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, timeZone)
}

func testAccMemberTrafficCaptureAuthDnsSetting(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	trafficCaptureAuthDnsSetting map[string]any,
) string {
	trafficCaptureAuthDnsSettingStr := utils.ConvertMapToHCL(trafficCaptureAuthDnsSetting)

	return fmt.Sprintf(`
resource "nios_grid_member" "test_traffic_capture_auth_dns_setting" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    traffic_capture_auth_dns_setting = %s
    use_traffic_capture_auth_dns = true
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, trafficCaptureAuthDnsSettingStr)
}

func testAccMemberTrafficCaptureChrSetting(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	trafficCaptureChrSetting map[string]any,
) string {
	trafficCaptureChrSettingStr := utils.ConvertMapToHCL(trafficCaptureChrSetting)

	return fmt.Sprintf(`
resource "nios_grid_member" "test_traffic_capture_chr_setting" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    traffic_capture_chr_setting = %s
    use_traffic_capture_chr = true
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, trafficCaptureChrSettingStr)
}

func testAccMemberTrafficCaptureQpsSetting(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	trafficCaptureQpsSetting map[string]any,
) string {
	trafficCaptureQpsSettingStr := utils.ConvertMapToHCL(trafficCaptureQpsSetting)

	return fmt.Sprintf(`
resource "nios_grid_member" "test_traffic_capture_qps_setting" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    traffic_capture_qps_setting = %s
    use_traffic_capture_qps = true
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, trafficCaptureQpsSettingStr)
}

func testAccMemberTrafficCaptureRecDnsSetting(
	hostName, configAddrType, platform, serviceTypeConfig, ipv6SettingIP,
	vipAddress, vipGateway, vipSubnetMask string,
	trafficCaptureRecDnsSetting map[string]any,
) string {
	trafficCaptureRecDnsSettingStr := utils.ConvertMapToHCL(trafficCaptureRecDnsSetting)
	enabledv6 := configAddrType != "IPV4"
	ipv6SettingStr := ""
	if configAddrType != "IPV4" {
		ipv6SettingStr = fmt.Sprintf("enabled = %t\n\t\tvirtual_ip = %q\n\t\tgateway = \"2001::1\"\n\t\tcidr_prefix = 8", enabledv6, ipv6SettingIP)
	}
	vipSettingStr := ""
	if configAddrType == "IPV4" {
		vipSettingStr = fmt.Sprintf("\t\taddress = %q\n\t\tgateway = %q\n\t\tsubnet_mask = %q", vipAddress, vipGateway, vipSubnetMask)
	}

	return fmt.Sprintf(`
resource "nios_grid_member" "test_traffic_capture_rec_dns_setting" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
		auto_router_config_enabled = false
		dscp = 0
		%s
		primary = true
		use_dscp = false
  }

    vip_setting = {
		%s
        primary = true
		dscp = 0
        use_dscp = false
    }

    traffic_capture_rec_dns_setting = %s
    use_traffic_capture_rec_dns = true
}
`, hostName, configAddrType, platform, serviceTypeConfig, ipv6SettingStr, vipSettingStr, trafficCaptureRecDnsSettingStr)
}

func testAccMemberTrafficCaptureRecQueriesSetting(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	trafficCaptureRecQueriesSetting map[string]any,
) string {
	trafficCaptureRecQueriesSettingStr := utils.ConvertMapToHCL(trafficCaptureRecQueriesSetting)

	return fmt.Sprintf(`
resource "nios_grid_member" "test_traffic_capture_rec_queries_setting" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    traffic_capture_rec_queries_setting = %s
    use_traffic_capture_rec_queries = true
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, trafficCaptureRecQueriesSettingStr)
}

func testAccMemberTrapNotifications(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	trapNotifications []map[string]any,
) string {
	trapNotificationsStr := utils.ConvertSliceOfMapsToHCL(trapNotifications)

	return fmt.Sprintf(`
resource "nios_grid_member" "test_trap_notifications" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    trap_notifications = %s
    use_trap_notifications = true
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, trapNotificationsStr)
}

func testAccMemberUpgradeGroup(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask, upgradeGroup string,
) string {
	return fmt.Sprintf(`
resource "nios_grid_upgradegroup" "test" {
    name = "new-group-%[1]s"
}
resource "nios_grid_member" "test_upgrade_group" {
    host_name = %[1]q
    config_addr_type = %[2]q
    platform = %[3]q
    service_type_configuration = %[4]q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %[5]q
        dscp = 0
        gateway = %[6]q
        primary = true
        subnet_mask = %[7]q
        use_dscp = false
    }

    upgrade_group = %[8]q
	use_traffic_capture_rec_dns = false
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, upgradeGroup)
}

func testAccMemberUseAutomatedTrafficCapture(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	useAutomatedTrafficCapture bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_automated_traffic_capture" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    automated_traffic_capture_setting = {
        destination            = "NONE"
        include_support_bundle = false
        keep_local_copy        = false
        traffic_capture_enable = false
    }

    use_automated_traffic_capture = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useAutomatedTrafficCapture)
}

func testAccMemberUseDnsResolverSetting(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	useDnsResolverSetting bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_dns_resolver_setting" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    dns_resolver_setting = {
        resolvers      = ["10.0.0.1"]
        search_domains = ["a.com"]
    }

    use_dns_resolver_setting = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useDnsResolverSetting)
}

func testAccMemberUseDscp(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	useDscp bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_dscp" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    dscp = 0
    use_dscp = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useDscp)
}

func testAccMemberUseEmailSetting(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	useEmailSetting bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_email_setting" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    use_email_setting = %t
	email_setting = {
		enabled            = false
		port_number        = 25
		relay_enabled      = false
		smtps              = false
		use_authentication = false
    }
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useEmailSetting)
}

func testAccMemberUseEnableLom(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	useEnableLom bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_enable_lom" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    use_enable_lom = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useEnableLom)
}

func testAccMemberUseEnableMemberRedirect(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	useEnableMemberRedirect bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_enable_member_redirect" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    use_enable_member_redirect = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useEnableMemberRedirect)
}

func testAccMemberUseExternalSyslogBackupServers(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	useExternalSyslogBackupServers bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_external_syslog_backup_servers" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    use_external_syslog_backup_servers = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useExternalSyslogBackupServers)
}

func testAccMemberUseRemoteConsoleAccessEnable(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	useRemoteConsoleAccessEnable bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_remote_console_access_enable" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    use_remote_console_access_enable = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useRemoteConsoleAccessEnable)
}

func testAccMemberUseSnmpSetting(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	useSnmpSetting bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_snmp_setting" {
	host_name = %q
	config_addr_type = %q
	platform = %q
	service_type_configuration = %q

	ipv6_setting = {
		auto_router_config_enabled = false
		dscp = 0
		enabled = false
		primary = true
		use_dscp = false
	}

	vip_setting = {
		address = %q
		dscp = 0
		gateway = %q
		primary = true
		subnet_mask = %q
		use_dscp = false
	}

	use_snmp_setting = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useSnmpSetting)
}

func testAccMemberUseSupportAccessEnable(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	useSupportAccessEnable bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_support_access_enable" {
	host_name = %q
	config_addr_type = %q
	platform = %q
	service_type_configuration = %q

	ipv6_setting = {
		auto_router_config_enabled = false
		dscp = 0
		enabled = false
		primary = true
		use_dscp = false
	}

	vip_setting = {
		address = %q
		dscp = 0
		gateway = %q
		primary = true
		subnet_mask = %q
		use_dscp = false
	}

	use_support_access_enable = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useSupportAccessEnable)
}

func testAccMemberUseSyslogProxySetting(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	useSyslogProxySetting bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_syslog_proxy_setting" {
	host_name = %q
	config_addr_type = %q
	platform = %q
	service_type_configuration = %q

	ipv6_setting = {
		auto_router_config_enabled = false
		dscp = 0
		enabled = false
		primary = true
		use_dscp = false
	}

	vip_setting = {
		address = %q
		dscp = 0
		gateway = %q
		primary = true
		subnet_mask = %q
		use_dscp = false
	}

	use_syslog_proxy_setting = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useSyslogProxySetting)
}

func testAccMemberUseThresholdTraps(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	useThresholdTraps bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_threshold_traps" {
	host_name = %q
	config_addr_type = %q
	platform = %q
	service_type_configuration = %q

	ipv6_setting = {
		auto_router_config_enabled = false
		dscp = 0
		enabled = false
		primary = true
		use_dscp = false
	}

	vip_setting = {
		address = %q
		dscp = 0
		gateway = %q
		primary = true
		subnet_mask = %q
		use_dscp = false
	}

	use_threshold_traps = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useThresholdTraps)
}

func testAccMemberUseTimeZone(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	useTimeZone bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_time_zone" {
	host_name = %q
	config_addr_type = %q
	platform = %q
	service_type_configuration = %q

	ipv6_setting = {
		auto_router_config_enabled = false
		dscp = 0
		enabled = false
		primary = true
		use_dscp = false
	}

	vip_setting = {
		address = %q
		dscp = 0
		gateway = %q
		primary = true
		subnet_mask = %q
		use_dscp = false
	}

	use_time_zone = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useTimeZone)
}

func testAccMemberUseTrafficCaptureAuthDns(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	useTrafficCaptureAuthDns bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_traffic_capture_auth_dns" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    use_traffic_capture_auth_dns = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useTrafficCaptureAuthDns)
}

func testAccMemberUseTrafficCaptureChr(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	useTrafficCaptureChr bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_traffic_capture_chr" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    use_traffic_capture_chr = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useTrafficCaptureChr)
}

func testAccMemberUseTrafficCaptureQps(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	useTrafficCaptureQps bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_traffic_capture_qps" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    use_traffic_capture_qps = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useTrafficCaptureQps)
}

func testAccMemberUseTrafficCaptureRecDns(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	useTrafficCaptureRecDns bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_traffic_capture_rec_dns" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    use_traffic_capture_rec_dns = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useTrafficCaptureRecDns)
}

func testAccMemberUseTrafficCaptureRecQueries(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	useTrafficCaptureRecQueries bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_traffic_capture_rec_queries" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    use_traffic_capture_rec_queries = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useTrafficCaptureRecQueries)
}

func testAccMemberUseTrapNotifications(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	useTrapNotifications bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_trap_notifications" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    use_trap_notifications = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useTrapNotifications)
}

func testAccMemberUseV4Vrrp(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	useV4Vrrp bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_v4_vrrp" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    use_v4_vrrp = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useV4Vrrp)
}

func testAccMemberVipSetting(
	hostName, configAddrType, platform, serviceTypeConfig string,
	vipSetting map[string]any,
) string {
	vipSettingStr := utils.ConvertMapToHCL(vipSetting)

	return fmt.Sprintf(`
resource "nios_grid_member" "test_vip_setting" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = %s
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipSettingStr)
}

func testAccMemberVpnMtu(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	vpnMtu int,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_vpn_mtu" {
    host_name = %q
    config_addr_type = %q
    platform = %q
    service_type_configuration = %q

    ipv6_setting = {
        auto_router_config_enabled = false
        dscp = 0
        enabled = false
        primary = true
        use_dscp = false
    }

    vip_setting = {
        address = %q
        dscp = 0
        gateway = %q
        primary = true
        subnet_mask = %q
        use_dscp = false
    }

    vpn_mtu = %d
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, vpnMtu)
}

func getTestDataPath() string {
	wd, err := os.Getwd()
	if err != nil {
		return "../../testdata/nios_grid_member"
	}
	return filepath.Join(wd, "../../testdata/nios_grid_member")
}
