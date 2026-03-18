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

/*
// Manage grid Member with Basic Fields
resource "nios_grid_member" "grid_member_basic" {
    host_name = "HOST_NAME_REPLACE_ME"
}

// Manage grid Member with Additional Fields
resource "nios_grid_member" "grid_member_with_additional_fields" {
    host_name = "HOST_NAME_REPLACE_ME"

// TODO : Add additional optional fields below

    //Extensible Attributes
    extattrs = {
        Site = "location-1"
    }
}
*/

var readableAttributesForMember = "active_position,additional_ip_list,automated_traffic_capture_setting,bgp_as,comment,config_addr_type,csp_access_key,csp_member_setting,dns_resolver_setting,dscp,email_setting,enable_ha,enable_lom,enable_member_redirect,enable_ro_api_access,extattrs,external_syslog_backup_servers,external_syslog_server_enable,ha_cloud_platform,ha_on_cloud,host_name,ipv6_setting,ipv6_static_routes,is_dscp_capable,lan2_enabled,lan2_port_setting,lom_network_config,lom_users,master_candidate,member_service_communication,mgmt_port_setting,mmdb_ea_build_time,mmdb_geoip_build_time,nat_setting,node_info,ntp_setting,ospf_list,passive_ha_arp_enabled,platform,pre_provisioning,preserve_if_owns_delegation,remote_console_access_enable,router_id,service_status,service_type_configuration,snmp_setting,static_routes,support_access_enable,support_access_info,syslog_proxy_setting,syslog_servers,syslog_size,threshold_traps,time_zone,traffic_capture_auth_dns_setting,traffic_capture_chr_setting,traffic_capture_qps_setting,traffic_capture_rec_dns_setting,traffic_capture_rec_queries_setting,trap_notifications,upgrade_group,use_automated_traffic_capture,use_dns_resolver_setting,use_dscp,use_email_setting,use_enable_lom,use_enable_member_redirect,use_external_syslog_backup_servers,use_remote_console_access_enable,use_snmp_setting,use_support_access_enable,use_syslog_proxy_setting,use_threshold_traps,use_time_zone,use_traffic_capture_auth_dns,use_traffic_capture_chr,use_traffic_capture_qps,use_traffic_capture_rec_dns,use_traffic_capture_rec_queries,use_trap_notifications,use_v4_vrrp,vip_setting,vpn_mtu"

func TestAccMemberResource_basic(t *testing.T) { //works
	var resourceName = "nios_grid_member.test"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					resource.TestCheckResourceAttr(resourceName, "vip_setting.gateway", "172.28.82.1"),
					resource.TestCheckResourceAttr(resourceName, "vip_setting.primary", "true"),
					resource.TestCheckResourceAttr(resourceName, "vip_setting.subnet_mask", "255.255.254.0"),
					resource.TestCheckResourceAttr(resourceName, "vip_setting.use_dscp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_disappears(t *testing.T) { //works
	resourceName := "nios_grid_member.test"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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

func TestAccMemberResource_AdditionalIpList(t *testing.T) { //works
	var resourceName = "nios_grid_member.test_additional_ip_list"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// issue password is senstive
func TestAccMemberResource_AutomatedTrafficCaptureSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_automated_traffic_capture_setting"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// work
func TestAccMemberResource_BgpAs(t *testing.T) {
	var resourceName = "nios_grid_member.test_bgp_as"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// work
func TestAccMemberResource_Comment(t *testing.T) {
	var resourceName = "nios_grid_member.test_comment"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// issue - "Invalid value for address: \"2001:db8:4958:7d81::894d\": Invalid IPv4 address
func TestAccMemberResource_ConfigAddrType(t *testing.T) {
	var resourceName = "nios_grid_member.test_config_addr_type"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress4 := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))
	vipAddress6 := fmt.Sprintf("2001:db8:%x:%x::%x", acctest.RandomNumber(65535), acctest.RandomNumber(65535), acctest.RandomNumber(65535))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberConfigAddrType(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress4,
					"172.28.82.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "config_addr_type", "IPV4"),
				),
			},
			{
				Config: testAccMemberConfigAddrType(
					hostName,
					"IPV6",
					"VNIOS",
					"ALL_V6",
					vipAddress6,
					"2001::1",
					"64",
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
	cspAccessKeyVal := []string{"CSP_ACCESS_KEY_REPLACE_ME1", "CSP_ACCESS_KEY_REPLACE_ME2"}
	cspAccessKeyValUpdated := []string{"CSP_ACCESS_KEY_REPLACE_ME1", "CSP_ACCESS_KEY_REPLACE_ME2"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberCspAccessKey("HOST_NAME_REPLACE_ME", cspAccessKeyVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "csp_access_key", "CSP_ACCESS_KEY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberCspAccessKey("HOST_NAME_REPLACE_ME", cspAccessKeyValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "csp_access_key", "CSP_ACCESS_KEY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// getting error via wapi - IB.Data.Conflict:CSP test connectivity failed: 'Could not perform Test Connection on offline member.')",
// issue
func TestAccMemberResource_CspMemberSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_csp_member_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

	cspMemberSettingVal := map[string]any{
		"csp_dns_resolver":     "52.119.40.101",
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
					"172.28.82.1",
					"255.255.254.0",
					cspMemberSettingVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "host_name", hostName),
					resource.TestCheckResourceAttr(resourceName, "config_addr_type", "IPV4"),
					resource.TestCheckResourceAttr(resourceName, "platform", "VNIOS"),
					resource.TestCheckResourceAttr(resourceName, "service_type_configuration", "ALL_V4"),
					// resource.TestCheckResourceAttr(resourceName, "csp_member_setting.csp_dns_resolver", "8.8.8.8"),
				),
			},
			{
				Config: testAccMemberCspMemberSetting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.82.1",
					"255.255.254.0",
					cspMemberSettingValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					//resource.TestCheckResourceAttr(resourceName, "csp_member_setting.csp_dns_resolver", "1.1.1.1"),
				),
			},
		},
	})
}

// works
func TestAccMemberResource_DnsResolverSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_dns_resolver_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
					"255.255.254.0",
					dnsResolverSettingVal,
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "host_name", hostName),
					resource.TestCheckResourceAttr(resourceName, "config_addr_type", "IPV4"),
					resource.TestCheckResourceAttr(resourceName, "platform", "VNIOS"),
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
					"172.28.82.1",
					"255.255.254.0",
					dnsResolverSettingValUpdated,
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_resolver_setting.resolvers.0", "10.0.0.2"),
					resource.TestCheckResourceAttr(resourceName, "dns_resolver_setting.search_domains.0", "b.com"),
				),
			},
		},
	})
}

// works
func TestAccMemberResource_Dscp(t *testing.T) {
	var resourceName = "nios_grid_member.test_dscp"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_EmailSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_email_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

	emailSettingVal := map[string]any{
		"enabled":            false,
		"port_number":        25,
		"relay_enabled":      false,
		"smtps":              false,
		"use_authentication": false,
	}
	emailSettingValUpdated := map[string]any{
		"enabled":            true,
		"port_number":        25,
		"relay_enabled":      false,
		"smtps":              false,
		"use_authentication": false,
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
					"172.28.82.1",
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
					"172.28.82.1",
					"255.255.254.0",
					emailSettingValUpdated,
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_setting.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "email_setting.port_number", "25"),
					resource.TestCheckResourceAttr(resourceName, "email_setting.relay_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "email_setting.smtps", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_email_setting", "true"),
				),
			},
		},
	})
}

// works
func TestAccMemberResource_EnableHa(t *testing.T) {
	var resourceName = "nios_grid_member.test_enable_ha"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberEnableHa(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.82.1", "255.255.254.0",
					true, 112,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ha", "true"),
					resource.TestCheckResourceAttr(resourceName, "router_id", "112"),
					resource.TestCheckResourceAttr(resourceName, "node_info.#", "2"),
				),
			},
		},
	})
}

// works
func TestAccMemberResource_EnableLom(t *testing.T) {
	var resourceName = "nios_grid_member.test_enable_lom"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_EnableMemberRedirect(t *testing.T) {
	var resourceName = "nios_grid_member.test_enable_member_redirect"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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
	t.Skip("Member should be Grid Master Candidate to enable read-only API access.")
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberEnableRoApiAccess(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.82.1",
					"255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ro_api_access", "true"),
				),
			},
			{
				Config: testAccMemberEnableRoApiAccess(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.82.1",
					"255.255.254.0",
					false,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ro_api_access", "false"),
				),
			},
		},
	})
}

// works
func TestAccMemberResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_grid_member.test_extattrs"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))
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
					"172.28.82.1",
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
					"172.28.82.1",
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

// /works
func TestAccMemberResource_ExternalSyslogBackupServers(t *testing.T) {
	var resourceName = "nios_grid_member.test_external_syslog_backup_servers"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
			"password":        "Password123!",
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
					"172.28.82.1",
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
					"172.28.82.1",
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

// issue - checking
func TestAccMemberResource_HaCloudPlatform(t *testing.T) {
	var resourceName = "nios_grid_member.test_ha_cloud_platform"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberHaCloudPlatform("HA_CLOUD_PLATFORM_REPLACE_ME", "AWS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ha_cloud_platform", "AWS"),
				),
			},
			{
				Config: testAccMemberHaCloudPlatform("HA_CLOUD_PLATFORM_REPLACE_ME", "AZURE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ha_cloud_platform", "AZURE"),
				),
			},
			{
				Config: testAccMemberHaCloudPlatform("HA_CLOUD_PLATFORM_REPLACE_ME", "GCP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ha_cloud_platform", "GCP"),
				),
			},
		},
	})
}

// issue - checking
func TestAccMemberResource_HaOnCloud(t *testing.T) {
	var resourceName = "nios_grid_member.test_ha_on_cloud"
	var v grid.Member
	t.Skip("Unknown argument/field: 'ha_on_cloud'")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberHaOnCloud("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ha_on_cloud", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberHaOnCloud("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ha_on_cloud", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// work
func TestAccMemberResource_HostName(t *testing.T) {
	var resourceName = "nios_grid_member.test_host_name"
	var v grid.Member

	hostName1 := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	hostName2 := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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
	t.Skip("need actual member with hardware licence - DSCP cannot be configured on an IB-UNKNOWN appliance. wapi will not set vaule even if you set")
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

	ipv6SettingVal := map[string]any{
		"auto_router_config_enabled": false,
		"dscp":                       0,
		"enabled":                    false,
		"primary":                    true,
		"use_dscp":                   true,
	}
	ipv6SettingValUpdated := map[string]any{
		"auto_router_config_enabled": true,
		"dscp":                       0,
		"enabled":                    true,
		"primary":                    true,
		"use_dscp":                   false,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberIpv6Setting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.82.1",
					"255.255.254.0",
					ipv6SettingVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.auto_router_config_enabled", "false"),
				),
			},
			{
				Config: testAccMemberIpv6Setting(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.82.1",
					"255.255.254.0",
					ipv6SettingValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting.auto_router_config_enabled", "true"),
				),
			},
		},
	})
}

// pending
func TestAccMemberResource_Ipv6StaticRoutes(t *testing.T) {
	var resourceName = "nios_grid_member.test_ipv6_static_routes"
	var v grid.Member
	ipv6StaticRoutesVal := []map[string]any{}
	ipv6StaticRoutesValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberIpv6StaticRoutes("HOST_NAME_REPLACE_ME", ipv6StaticRoutesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_static_routes", "IPV6_STATIC_ROUTES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberIpv6StaticRoutes("HOST_NAME_REPLACE_ME", ipv6StaticRoutesValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_static_routes", "IPV6_STATIC_ROUTES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// pending
func TestAccMemberResource_Lan2Enabled(t *testing.T) {
	var resourceName = "nios_grid_member.test_lan2_enabled"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberLan2Enabled("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lan2_enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberLan2Enabled("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lan2_enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// pending
func TestAccMemberResource_Lan2PortSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_lan2_port_setting"
	var v grid.Member
	lan2PortSettingVal := map[string]any{}
	lan2PortSettingValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberLan2PortSetting("HOST_NAME_REPLACE_ME", lan2PortSettingVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting", "LAN2_PORT_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberLan2PortSetting("HOST_NAME_REPLACE_ME", lan2PortSettingValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting", "LAN2_PORT_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// pending
func TestAccMemberResource_LomNetworkConfig(t *testing.T) {
	var resourceName = "nios_grid_member.test_lom_network_config"
	var v grid.Member
	lomNetworkConfigVal := []map[string]any{}
	lomNetworkConfigValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberLomNetworkConfig("HOST_NAME_REPLACE_ME", lomNetworkConfigVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lom_network_config", "LOM_NETWORK_CONFIG_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberLomNetworkConfig("HOST_NAME_REPLACE_ME", lomNetworkConfigValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lom_network_config", "LOM_NETWORK_CONFIG_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// works
func TestAccMemberResource_LomUsers(t *testing.T) {
	var resourceName = "nios_grid_member.test_lom_users"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
			"name":     "LOMuser1",
			"password": "@#nios123",
			"role":     "USER",
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
					"172.28.82.1",
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
					"172.28.82.1",
					"255.255.254.0",
					lomUsersValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lom_users.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "lom_users.0.name", "LOMuser1"),
					resource.TestCheckResourceAttr(resourceName, "lom_users.0.role", "USER"),
					resource.TestCheckResourceAttr(resourceName, "lom_users.0.disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// issue
// "text": "Cannot delete or add member services. only edit "
func TestAccMemberResource_MemberServiceCommunication(t *testing.T) {
	var resourceName = "nios_grid_member.test_member_service_communication"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

	memberServiceCommunicationVal := []map[string]any{
		{
			"service": "GRID_BACKUP",
			"type":    "IPV4",
		},
	}

	memberServiceCommunicationValUpdated := []map[string]any{
		{
			"service": "GRID_BACKUP",
			"type":    "IPV6",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberMemberServiceCommunication(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.82.1",
					"255.255.254.0",
					memberServiceCommunicationVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member_service_communication.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "member_service_communication.0.service", "GRID_BACKUP"),
					resource.TestCheckResourceAttr(resourceName, "member_service_communication.0.type", "IPV4"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberMemberServiceCommunication(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.82.1",
					"255.255.254.0",
					memberServiceCommunicationValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member_service_communication.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "member_service_communication.0.service", "GRID_BACKUP"),
					resource.TestCheckResourceAttr(resourceName, "member_service_communication.0.type", "IPV6"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// issue
func TestAccMemberResource_MgmtPortSetting(t *testing.T) { // ui cant change settings
	var resourceName = "nios_grid_member.test_mgmt_port_setting"
	var v grid.Member
	mgmtPortSettingVal := map[string]any{}
	mgmtPortSettingValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberMgmtPortSetting("HOST_NAME_REPLACE_ME", mgmtPortSettingVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mgmt_port_setting", "MGMT_PORT_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberMgmtPortSetting("HOST_NAME_REPLACE_ME", mgmtPortSettingValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mgmt_port_setting", "MGMT_PORT_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccMemberResource_NatSetting(t *testing.T) { // work
	var resourceName = "nios_grid_member.test_nat_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))
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
					"172.28.82.1",
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
					"172.28.82.1",
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
func TestAccMemberResource_NodeInfo(t *testing.T) { //works
	var resourceName = "nios_grid_member.test_node_info"
	var v grid.Member
	t.Skip("HA node has to be configured as per member and grid, this is tested locally and it passes")
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

	nodeInfoVal := []map[string]any{
		{
			"lan_ha_port_setting": map[string]any{
				"ha_cloud_attribute": "UNK",
				"ha_ip_address":      "172.28.82.11",
				"ha_port_setting": map[string]any{
					"auto_port_setting_enabled": true,
					"speed":                     "10",
				},
				"lan_port_setting": map[string]any{
					"auto_port_setting_enabled": true,
				},
				"mgmt_lan": "172.28.82.32",
			},
		},
		{
			"lan_ha_port_setting": map[string]any{
				"ha_cloud_attribute": "UNK",
				"ha_ip_address":      "172.28.82.41",
				"ha_port_setting": map[string]any{
					"auto_port_setting_enabled": true,
					"speed":                     "10",
				},
				"lan_port_setting": map[string]any{
					"auto_port_setting_enabled": true,
				},
				"mgmt_lan": "172.28.82.43",
			},
		},
	}

	nodeInfoValUpdated := []map[string]any{
		{
			"lan_ha_port_setting": map[string]any{
				"ha_cloud_attribute": "UNK",
				"ha_ip_address":      "172.28.82.12",
				"ha_port_setting": map[string]any{
					"auto_port_setting_enabled": true,
					"speed":                     "10",
				},
				"lan_port_setting": map[string]any{
					"auto_port_setting_enabled": true,
				},
				"mgmt_lan": "172.28.82.33",
			},
		},
		{
			"lan_ha_port_setting": map[string]any{
				"ha_cloud_attribute": "UNK",
				"ha_ip_address":      "172.28.82.42",
				"ha_port_setting": map[string]any{
					"auto_port_setting_enabled": true,
					"speed":                     "10",
				},
				"lan_port_setting": map[string]any{
					"auto_port_setting_enabled": true,
				},
				"mgmt_lan": "172.28.82.44",
			},
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberNodeInfo(hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.82.1", "255.255.254.0", nodeInfoVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "node_info.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.lan_ha_port_setting.ha_ip_address", "172.28.82.11"),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.lan_ha_port_setting.mgmt_lan", "172.28.82.32"),
					resource.TestCheckResourceAttr(resourceName, "node_info.1.lan_ha_port_setting.ha_ip_address", "172.28.82.41"),
					resource.TestCheckResourceAttr(resourceName, "node_info.1.lan_ha_port_setting.mgmt_lan", "172.28.82.43"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberNodeInfo(hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.82.1", "255.255.254.0", nodeInfoValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "node_info.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.lan_ha_port_setting.ha_ip_address", "172.28.82.12"),
					resource.TestCheckResourceAttr(resourceName, "node_info.0.lan_ha_port_setting.mgmt_lan", "172.28.82.33"),
					resource.TestCheckResourceAttr(resourceName, "node_info.1.lan_ha_port_setting.ha_ip_address", "172.28.82.42"),
					resource.TestCheckResourceAttr(resourceName, "node_info.1.lan_ha_port_setting.mgmt_lan", "172.28.82.44"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_NtpSetting(t *testing.T) { //works
	var resourceName = "nios_grid_member.test_ntp_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					vipAddress, "172.28.82.1", "255.255.254.0",
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
					vipAddress, "172.28.82.1", "255.255.254.0",
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

// works
func TestAccMemberResource_OspfList(t *testing.T) {
	var resourceName = "nios_grid_member.test_ospf_list"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
			"authentication_type":    "NONE",
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
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberOspfList(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.82.1", "255.255.254.0",
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
					vipAddress, "172.28.82.1", "255.255.254.0",
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

// needs a special HA grid - SA-HA type - deploying and checking ..
func TestAccMemberResource_PassiveHaArpEnabled(t *testing.T) {
	var resourceName = "nios_grid_member.test_passive_ha_arp_enabled"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberPassiveHaArpEnabled(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.82.1", "255.255.254.0",
					true,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ha", "true"),
					resource.TestCheckResourceAttr(resourceName, "router_id", "112"),
					resource.TestCheckResourceAttr(resourceName, "node_info.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "passive_ha_arp_enabled", "true"),
				),
			},
		},
	})
}

func TestAccMemberResource_Platform(t *testing.T) { //works
	var resourceName = "nios_grid_member.test_platform"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberPlatform(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.82.1", "255.255.254.0",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "platform", "VNIOS"),
				),
			},
			{
				Config: testAccMemberPlatform(
					hostName, "IPV4", "INFOBLOX", "ALL_V4",
					vipAddress, "172.28.82.1", "255.255.254.0",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "platform", "INFOBLOX"),
				),
			},
		},
	})
}

// unable to get pre_provisioning - only edit and delete can be done
func TestAccMemberResource_PreProvisioning(t *testing.T) {
	var resourceName = "nios_grid_member.test"
	var v grid.Member
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
			// Create without pre_provisioning
			{
				Config: testAccMemberPreProvisioningUpdate(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.82.1",
					"255.255.254.0",
					nil,
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
					resource.TestCheckResourceAttr(resourceName, "vip_setting.gateway", "172.28.82.1"),
					resource.TestCheckResourceAttr(resourceName, "vip_setting.primary", "true"),
					resource.TestCheckResourceAttr(resourceName, "vip_setting.subnet_mask", "255.255.254.0"),
					resource.TestCheckResourceAttr(resourceName, "vip_setting.use_dscp", "false"),
				),
			},
			// Update with pre_provisioning
			{
				Config: testAccMemberPreProvisioningUpdate(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.82.1",
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
				Config: testAccMemberPreProvisioningUpdate(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.82.1",
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

func TestAccMemberResource_PreserveIfOwnsDelegation(t *testing.T) { //works
	var resourceName = "nios_grid_member.test_preserve_if_owns_delegation"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberPreserveIfOwnsDelegation(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.82.1", "255.255.254.0",
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
					vipAddress, "172.28.82.1", "255.255.254.0",
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

func TestAccMemberResource_RemoteConsoleAccessEnable(t *testing.T) { //works
	var resourceName = "nios_grid_member.test_remote_console_access_enable"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberRemoteConsoleAccessEnable(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.82.1", "255.255.254.0",
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
					vipAddress, "172.28.82.1", "255.255.254.0",
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

func TestAccMemberResource_RouterId(t *testing.T) { //works
	var resourceName = "nios_grid_member.test_router_id"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberRouterId(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.82.1", "255.255.254.0",
					111,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "router_id", "111"),
				),
			},
			{
				Config: testAccMemberRouterId(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.82.1", "255.255.254.0",
					112,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "router_id", "112"),
				),
			},
		},
	})
}

// issue - rajat twm
func TestAccMemberResource_ServiceTypeConfiguration(t *testing.T) {
	var resourceName = "nios_grid_member.test_service_type_configuration"
	var v grid.Member
	t.Skip("you would need a grid member which supports both ipv4 and v6- else error - ")
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress4 := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))
	vipAddress6 := fmt.Sprintf("2001:db8:%x:%x::%x", acctest.RandomNumber(65535), acctest.RandomNumber(65535), acctest.RandomNumber(65535))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberServiceTypeConfiguration(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress4, "172.28.82.1", "255.255.254.0",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service_type_configuration", "ALL_V4"),
				),
			},
			{
				Config: testAccMemberServiceTypeConfiguration(
					hostName, "IPV6", "VNIOS", "ALL_V6",
					vipAddress6, "2001::1", "64",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service_type_configuration", "ALL_V6"),
				),
			},
		},
	})
}

func TestAccMemberResource_SnmpSetting(t *testing.T) { //works
	var resourceName = "nios_grid_member.test_snmp_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					vipAddress, "172.28.82.1", "255.255.254.0",
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
					vipAddress, "172.28.82.1", "255.255.254.0",
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

// dont know how to do set on UI///issue - rajat
func TestAccMemberResource_StaticRoutes(t *testing.T) {
	var resourceName = "nios_grid_member.test_static_routes"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

	staticRoutesVal := []map[string]any{
		{
			"address":     "172.28.90.0",
			"gateway":     "172.28.82.1",
			"subnet_mask": "255.255.254.0",
			"primary":     true,
			"dscp":        0,
			"use_dscp":    true,
		},
	}
	staticRoutesValUpdated := []map[string]any{
		{
			"address":     "172.28.90.0",
			"gateway":     "172.28.82.1",
			"subnet_mask": "255.255.254.0",
			"primary":     true,
			"dscp":        0,
			"use_dscp":    true,
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberStaticRoutes(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.82.1", "255.255.254.0",
					staticRoutesVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "static_routes.#", "1"),
				),
			},
			{
				Config: testAccMemberStaticRoutes(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.82.1", "255.255.254.0",
					staticRoutesValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "static_routes.#", "1"),
				),
			},
		},
	})
}

// works
func TestAccMemberResource_SupportAccessEnable(t *testing.T) {
	var resourceName = "nios_grid_member.test_support_access_enable"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberSupportAccessEnable(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.82.1", "255.255.254.0",
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
					vipAddress, "172.28.82.1", "255.255.254.0",
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

// issue - chai
func TestAccMemberResource_SyslogProxySetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_syslog_proxy_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))
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
	}

	syslogProxySettingVal := map[string]any{

		"client_acls": []map[string]any{
			{
				"struct":     "addressac",
				"address":    "192.0.0.1",
				"permission": "ALLOW",
			},
		},
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
				"address":    "192.0.0.1",
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
					vipAddress, "172.28.82.1", "255.255.254.0",
					syslogProxySettingVal,
					syslogServersVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					// resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.client_acls.#", "1"),
					// resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.enable", "false"),
					// resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.tcp_enable", "false"),
					// resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.tcp_port", "514"),
					// resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.udp_enable", "true"),
					// resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.udp_port", "514"),
					// resource.TestCheckResourceAttr(resourceName, "use_syslog_proxy_setting", "true"),
				),
			},
			{
				Config: testAccMemberSyslogProxySetting(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.82.1", "255.255.254.0",
					syslogProxySettingValUpdated,
					syslogServersVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					// resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.client_acls.#", "1"),
					// resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.enable", "true"),
					// resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.tcp_enable", "true"),
					// resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.tcp_port", "1514"),
					// resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.udp_enable", "false"),
					// resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting.udp_port", "514"),
					// resource.TestCheckResourceAttr(resourceName, "use_syslog_proxy_setting", "true"),
				),
			},
		},
	})
}

// issue
func TestAccMemberResource_SyslogServers(t *testing.T) {
	var resourceName = "nios_grid_member.test_syslog_servers"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))
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
					"172.28.82.1",
					"255.255.254.0",
					syslogServersVal,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.connection_type", "TCP"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.port", "514"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.severity", "DEBUG"),
					resource.TestCheckResourceAttr(resourceName, "use_syslog_proxy_setting", "true"),
				),
			},
			{
				Config: testAccMemberSyslogServers(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.82.1",
					"255.255.254.0",
					syslogServersValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.connection_type", "TCP"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.port", "515"),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers.0.severity", "INFO"),
					resource.TestCheckResourceAttr(resourceName, "use_syslog_proxy_setting", "true"),
				),
			},
		},
	})
}

// works
func TestAccMemberResource_SyslogSize(t *testing.T) {
	var resourceName = "nios_grid_member.test_syslog_size"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_ThresholdTraps(t *testing.T) {
	var resourceName = "nios_grid_member.test_threshold_traps"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					vipAddress, "172.28.82.1", "255.255.254.0",
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
					vipAddress, "172.28.82.1", "255.255.254.0",
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

// works
func TestAccMemberResource_TimeZone(t *testing.T) {
	var resourceName = "nios_grid_member.test_time_zone"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_TrafficCaptureAuthDnsSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_traffic_capture_auth_dns_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					vipAddress, "172.28.82.1", "255.255.254.0",
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
					vipAddress, "172.28.82.1", "255.255.254.0",
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

// works
func TestAccMemberResource_TrafficCaptureChrSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_traffic_capture_chr_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					vipAddress, "172.28.82.1", "255.255.254.0",
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
					vipAddress, "172.28.82.1", "255.255.254.0",
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

// works
func TestAccMemberResource_TrafficCaptureQpsSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_traffic_capture_qps_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					vipAddress, "172.28.82.1", "255.255.254.0",
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
					vipAddress, "172.28.82.1", "255.255.254.0",
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

// issue
func TestAccMemberResource_TrafficCaptureRecDnsSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_traffic_capture_rec_dns_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

	trafficCaptureRecDnsSettingVal := map[string]any{
		"kpi_monitored_domains":            []string{},
		"rec_dns_latency_listen_on_source": "VIP_V4",
		"rec_dns_latency_trigger_enable":   false,
	}
	trafficCaptureRecDnsSettingValUpdated := map[string]any{
		"kpi_monitored_domains":            []map[string]any{{"domain_name": "a.com", "record_type": "A"}},
		"rec_dns_latency_listen_on_source": "VIP_V4",
		"rec_dns_latency_trigger_enable":   false,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberTrafficCaptureRecDnsSetting(
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.82.1", "255.255.254.0",
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
					hostName, "IPV4", "VNIOS", "ALL_V4",
					vipAddress, "172.28.82.1", "255.255.254.0",
					trafficCaptureRecDnsSettingValUpdated,
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_rec_dns_setting.rec_dns_latency_listen_on_source", "VIP_V4"),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_rec_dns_setting.rec_dns_latency_trigger_enable", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_rec_dns", "true"),
				),
			},
		},
	})
}

// works
func TestAccMemberResource_TrafficCaptureRecQueriesSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_traffic_capture_rec_queries_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					vipAddress, "172.28.82.1", "255.255.254.0",
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
					vipAddress, "172.28.82.1", "255.255.254.0",
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

// works
func TestAccMemberResource_TrapNotifications(t *testing.T) {
	var resourceName = "nios_grid_member.test_trap_notifications"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					vipAddress, "172.28.82.1", "255.255.254.0",
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
					vipAddress, "172.28.82.1", "255.255.254.0",
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
	t.Skip("UpgradeGroup is either master or default and cannot be updated.")
	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
					"255.255.254.0",
					"Grid Master",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "upgrade_group", "Grid Master"),
				),
			},
		},
	})
}

// works
func TestAccMemberResource_UseAutomatedTrafficCapture(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_automated_traffic_capture"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_UseDnsResolverSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_dns_resolver_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_UseDscp(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_dscp"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_UseEmailSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_email_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_UseEnableLom(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_enable_lom"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_UseEnableMemberRedirect(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_enable_member_redirect"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_UseExternalSyslogBackupServers(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_external_syslog_backup_servers"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_UseRemoteConsoleAccessEnable(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_remote_console_access_enable"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_UseSnmpSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_snmp_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_UseSupportAccessEnable(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_support_access_enable"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_UseSyslogProxySetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_syslog_proxy_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_UseThresholdTraps(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_threshold_traps"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_UseTimeZone(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_time_zone"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_UseTrafficCaptureAuthDns(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_traffic_capture_auth_dns"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_UseTrafficCaptureChr(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_traffic_capture_chr"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_UseTrafficCaptureQps(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_traffic_capture_qps"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_UseTrafficCaptureRecDns(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_traffic_capture_rec_dns"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_UseTrafficCaptureRecQueries(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_traffic_capture_rec_queries"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_UseTrapNotifications(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_trap_notifications"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_UseV4Vrrp(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_v4_vrrp"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

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
					"172.28.82.1",
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
					"172.28.82.1",
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

// works
func TestAccMemberResource_MasterCandidate(t *testing.T) {
	var resourceName = "nios_grid_member.test_master_candidate"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMemberMasterCandidate(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.82.1",
					"255.255.254.0",
					"true",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "host_name", hostName),
					resource.TestCheckResourceAttr(resourceName, "config_addr_type", "IPV4"),
					resource.TestCheckResourceAttr(resourceName, "platform", "VNIOS"),
					resource.TestCheckResourceAttr(resourceName, "service_type_configuration", "ALL_V4"),
					resource.TestCheckResourceAttr(resourceName, "master_candidate", "true"),
				),
			},
			{
				Config: testAccMemberMasterCandidate(
					hostName,
					"IPV4",
					"VNIOS",
					"ALL_V4",
					vipAddress,
					"172.28.82.1",
					"255.255.254.0",
					"false",
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "master_candidate", "false"),
				),
			},
		},
	})
}

// works
func TestAccMemberResource_VipSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_vip_setting"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))

	vipSettingVal := map[string]any{
		"address":     vipAddress,
		"dscp":        0,
		"gateway":     "172.28.82.1",
		"primary":     true,
		"subnet_mask": "255.255.254.0",
		"use_dscp":    false,
	}
	vipSettingValUpdated := map[string]any{
		"address":     vipAddress,
		"dscp":        0,
		"gateway":     "172.28.82.2",
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
					resource.TestCheckResourceAttr(resourceName, "vip_setting.gateway", "172.28.82.2"),
				),
			},
		},
	})
}

// works
func TestAccMemberResource_VpnMtu(t *testing.T) {
	var resourceName = "nios_grid_member.test_vpn_mtu"
	var v grid.Member

	hostName := fmt.Sprintf("infoblox-%s.localdomain", acctest.RandomName())
	vipAddress := fmt.Sprintf("172.28.83.%d", acctest.RandomNumber(254))
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
					"172.28.82.1",
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
					"172.28.82.1",
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

func testAccMemberConfigAddrType(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string, useV4Vrrp bool,
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
	use_v4_vrrp = %t // set use_v4_vrrp to false to avoid dependency on vrrp resource for config_addr_type testing
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, useV4Vrrp)
}

func testAccMemberCspAccessKey(hostName string, cspAccessKey []string) string {
	cspAccessKeyStr := utils.ConvertStringSliceToHCL(cspAccessKey)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_csp_access_key" {
    host_name = %q
    csp_access_key = %s
}
`, hostName, cspAccessKeyStr)
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
	dscp int, use_dscp bool,
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
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, dscp, use_dscp)
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
        ha_ip_address = "172.28.82.11"
        ha_port_setting = {
          auto_port_setting_enabled = true
          speed = "10"
        }
        lan_port_setting = {
          auto_port_setting_enabled = true
        }
        mgmt_lan = "172.28.82.32"
      }
    },
    {
      lan_ha_port_setting = {
        ha_cloud_attribute = "UNK"
        ha_ip_address = "172.28.82.41"
        ha_port_setting = {
          auto_port_setting_enabled = true
          speed = "10"
        }
        lan_port_setting = {
          auto_port_setting_enabled = true
        }
        mgmt_lan = "172.28.82.43"
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
	enableRoApiAccess bool,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_enable_ro_api_access" {
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

    enable_ro_api_access = %t
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, enableRoApiAccess)
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

func testAccMemberExternalSyslogServerEnable(hostName string, externalSyslogServerEnable string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_external_syslog_server_enable" {
    host_name = %q
    external_syslog_server_enable = %q
    use_syslog_proxy_setting = true
}
`, hostName, externalSyslogServerEnable)
}

func testAccMemberHaCloudPlatform(hostName string, haCloudPlatform string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_ha_cloud_platform" {
    host_name = %q
    ha_cloud_platform = %q
}
`, hostName, haCloudPlatform)
}

func testAccMemberHaOnCloud(hostName string, haOnCloud string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_ha_on_cloud" {
    host_name = %q
    ha_on_cloud = %q
}
`, hostName, haOnCloud)
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
		use_dscp = true
}
`, hostName, configAddrType, platform, serviceTypeConfig, ipv6SettingStr, vipAddress, vipGateway, vipSubnetMask)
}

func testAccMemberIpv6StaticRoutes(hostName string, ipv6StaticRoutes []map[string]any) string {
	ipv6StaticRoutesStr := utils.ConvertSliceOfMapsToHCL(ipv6StaticRoutes)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_ipv6_static_routes" {
    host_name = %q
    ipv6_static_routes = %s
}
`, hostName, ipv6StaticRoutesStr)
}

func testAccMemberLan2Enabled(hostName string, lan2Enabled string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_lan2_enabled" {
    host_name = %q
    lan2_enabled = %q
}
`, hostName, lan2Enabled)
}

func testAccMemberLan2PortSetting(hostName string, lan2PortSetting map[string]any) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_lan2_port_setting" {
    host_name = %q
    lan2_port_setting = %s
}
`, hostName, lan2PortSetting)
}

func testAccMemberLomNetworkConfig(hostName string, lomNetworkConfig []map[string]any) string {
	lomNetworkConfigStr := utils.ConvertSliceOfMapsToHCL(lomNetworkConfig)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_lom_network_config" {
    host_name = %q
    lom_network_config = %s
}
`, hostName, lomNetworkConfigStr)
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
	vipAddress, vipGateway, vipSubnetMask, masterCandidate string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_master_candidate" {
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

    master_candidate = %q
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, masterCandidate)
}

func testAccMemberMemberServiceCommunication(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	memberServiceCommunication []map[string]any,
) string {
	memberServiceCommunicationStr := utils.ConvertSliceOfMapsToHCL(memberServiceCommunication)

	return fmt.Sprintf(`
resource "nios_grid_member" "test_member_service_communication" {
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

    member_service_communication = %s
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, memberServiceCommunicationStr)
}

func testAccMemberMgmtPortSetting(hostName string, mgmtPortSetting map[string]any) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_mgmt_port_setting" {
    host_name = %q
    mgmt_port_setting = %s
}
`, hostName, mgmtPortSetting)
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
	nodeInfo []map[string]any,
) string {
	nodeInfoStr := utils.ConvertSliceOfMapsToHCL(nodeInfo)

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

    enable_ha = true
    router_id = 112
    node_info = %s
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, nodeInfoStr)
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
    router_id = 112
    passive_ha_arp_enabled = %t

    node_info = [
        {
            lan_ha_port_setting = {
                ha_cloud_attribute = "UNK"
                ha_ip_address = "172.28.82.11"
                ha_port_setting = {
                    auto_port_setting_enabled = true
                    speed = "10"
                }
                lan_port_setting = {
                    auto_port_setting_enabled = true
                }
                mgmt_lan = "172.28.82.32"
            }
        },
        {
            lan_ha_port_setting = {
                ha_cloud_attribute = "UNK"
                ha_ip_address = "172.28.82.41"
                ha_port_setting = {
                    auto_port_setting_enabled = true
                    speed = "10"
                }
                lan_port_setting = {
                    auto_port_setting_enabled = true
                }
                mgmt_lan = "172.28.82.43"
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

func testAccMemberPreProvisioningUpdate(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	preProvisioning map[string]any,
) string {
	// Convert map to HCL if pre_provisioning is provided
	var preProvConfig string
	if preProvisioning != nil {
		preProvisioningStr := utils.ConvertMapToHCL(preProvisioning)
		preProvConfig = fmt.Sprintf("\n\n    pre_provisioning = %s", preProvisioningStr)
	}

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
    }%s
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, preProvConfig)
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
	hostName, configAddrType, platform, serviceTypeConfiguration,
	vipAddress, vipGateway, vipSubnetMask string,
) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_service_type_configuration" {
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
`, hostName, configAddrType, platform, serviceTypeConfiguration, vipAddress, vipGateway, vipSubnetMask)
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
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, syslogProxySettingStr, syslogServersValStr)
}

func testAccMemberSyslogServers(
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	syslogServers []map[string]any,
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
    use_syslog_proxy_setting = true
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
	hostName, configAddrType, platform, serviceTypeConfig,
	vipAddress, vipGateway, vipSubnetMask string,
	trafficCaptureRecDnsSetting map[string]any,
) string {
	trafficCaptureRecDnsSettingStr := utils.ConvertMapToHCL(trafficCaptureRecDnsSetting)

	return fmt.Sprintf(`
resource "nios_grid_member" "test_traffic_capture_rec_dns_setting" {
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

    traffic_capture_rec_dns_setting = %s
    use_traffic_capture_rec_dns = true
}
`, hostName, configAddrType, platform, serviceTypeConfig, vipAddress, vipGateway, vipSubnetMask, trafficCaptureRecDnsSettingStr)
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
resource "nios_grid_member" "test_upgrade_group" {
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

    upgrade_group = %q
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
		return "../../testdata/nios_member"
	}
	return filepath.Join(wd, "../../testdata/nios_member")
}
