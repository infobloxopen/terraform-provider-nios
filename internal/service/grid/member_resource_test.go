package grid_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForMember = "active_position,additional_ip_list,automated_traffic_capture_setting,bgp_as,comment,config_addr_type,csp_access_key,csp_member_setting,dns_resolver_setting,dscp,email_setting,enable_ha,enable_lom,enable_member_redirect,enable_ro_api_access,extattrs,external_syslog_backup_servers,external_syslog_server_enable,ha_cloud_platform,ha_on_cloud,host_name,ipv6_setting,ipv6_static_routes,is_dscp_capable,lan2_enabled,lan2_port_setting,lom_network_config,lom_users,master_candidate,member_service_communication,mgmt_port_setting,mmdb_ea_build_time,mmdb_geoip_build_time,nat_setting,node_info,ntp_setting,ospf_list,passive_ha_arp_enabled,platform,pre_provisioning,preserve_if_owns_delegation,remote_console_access_enable,router_id,service_status,service_type_configuration,snmp_setting,static_routes,support_access_enable,support_access_info,syslog_proxy_setting,syslog_servers,syslog_size,threshold_traps,time_zone,traffic_capture_auth_dns_setting,traffic_capture_chr_setting,traffic_capture_qps_setting,traffic_capture_rec_dns_setting,traffic_capture_rec_queries_setting,trap_notifications,upgrade_group,use_automated_traffic_capture,use_dns_resolver_setting,use_dscp,use_email_setting,use_enable_lom,use_enable_member_redirect,use_external_syslog_backup_servers,use_remote_console_access_enable,use_snmp_setting,use_support_access_enable,use_syslog_proxy_setting,use_threshold_traps,use_time_zone,use_traffic_capture_auth_dns,use_traffic_capture_chr,use_traffic_capture_qps,use_traffic_capture_rec_dns,use_traffic_capture_rec_queries,use_trap_notifications,use_v4_vrrp,vip_setting,vpn_mtu"

func TestAccMemberResource_basic(t *testing.T) {
	var resourceName = "nios_grid_member.test"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberBasicConfig("grid"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_disappears(t *testing.T) {
	resourceName := "nios_grid_member.test"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMemberDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMemberBasicConfig("grid"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					testAccCheckMemberDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccMemberResource_Ref(t *testing.T) {
	var resourceName = "nios_grid_member.test_ref"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_AdditionalIpList(t *testing.T) {
	var resourceName = "nios_grid_member.test_additional_ip_list"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberAdditionalIpList("ADDITIONAL_IP_LIST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "additional_ip_list", "ADDITIONAL_IP_LIST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberAdditionalIpList("ADDITIONAL_IP_LIST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "additional_ip_list", "ADDITIONAL_IP_LIST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_AutomatedTrafficCaptureSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_automated_traffic_capture_setting"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberAutomatedTrafficCaptureSetting("AUTOMATED_TRAFFIC_CAPTURE_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "automated_traffic_capture_setting", "AUTOMATED_TRAFFIC_CAPTURE_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberAutomatedTrafficCaptureSetting("AUTOMATED_TRAFFIC_CAPTURE_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "automated_traffic_capture_setting", "AUTOMATED_TRAFFIC_CAPTURE_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_BgpAs(t *testing.T) {
	var resourceName = "nios_grid_member.test_bgp_as"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberBgpAs("BGP_AS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bgp_as", "BGP_AS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberBgpAs("BGP_AS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bgp_as", "BGP_AS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_Comment(t *testing.T) {
	var resourceName = "nios_grid_member.test_comment"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_ConfigAddrType(t *testing.T) {
	var resourceName = "nios_grid_member.test_config_addr_type"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberConfigAddrType("CONFIG_ADDR_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "config_addr_type", "CONFIG_ADDR_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberConfigAddrType("CONFIG_ADDR_TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "config_addr_type", "CONFIG_ADDR_TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_CspAccessKey(t *testing.T) {
	var resourceName = "nios_grid_member.test_csp_access_key"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberCspAccessKey("CSP_ACCESS_KEY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "csp_access_key", "CSP_ACCESS_KEY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberCspAccessKey("CSP_ACCESS_KEY_UPDATE_REPLACE_ME"),
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
	var resourceName = "nios_grid_member.test_csp_member_setting"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberCspMemberSetting("CSP_MEMBER_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "csp_member_setting", "CSP_MEMBER_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberCspMemberSetting("CSP_MEMBER_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "csp_member_setting", "CSP_MEMBER_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_DnsResolverSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_dns_resolver_setting"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberDnsResolverSetting("DNS_RESOLVER_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_resolver_setting", "DNS_RESOLVER_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberDnsResolverSetting("DNS_RESOLVER_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_resolver_setting", "DNS_RESOLVER_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_Dscp(t *testing.T) {
	var resourceName = "nios_grid_member.test_dscp"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberDscp("DSCP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dscp", "DSCP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberDscp("DSCP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dscp", "DSCP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_EmailSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_email_setting"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberEmailSetting("EMAIL_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_setting", "EMAIL_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberEmailSetting("EMAIL_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_setting", "EMAIL_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_EnableHa(t *testing.T) {
	var resourceName = "nios_grid_member.test_enable_ha"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberEnableHa("ENABLE_HA_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ha", "ENABLE_HA_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberEnableHa("ENABLE_HA_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ha", "ENABLE_HA_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_EnableLom(t *testing.T) {
	var resourceName = "nios_grid_member.test_enable_lom"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberEnableLom("ENABLE_LOM_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_lom", "ENABLE_LOM_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberEnableLom("ENABLE_LOM_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_lom", "ENABLE_LOM_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_EnableMemberRedirect(t *testing.T) {
	var resourceName = "nios_grid_member.test_enable_member_redirect"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberEnableMemberRedirect("ENABLE_MEMBER_REDIRECT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_member_redirect", "ENABLE_MEMBER_REDIRECT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberEnableMemberRedirect("ENABLE_MEMBER_REDIRECT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_member_redirect", "ENABLE_MEMBER_REDIRECT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_EnableRoApiAccess(t *testing.T) {
	var resourceName = "nios_grid_member.test_enable_ro_api_access"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberEnableRoApiAccess("ENABLE_RO_API_ACCESS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ro_api_access", "ENABLE_RO_API_ACCESS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberEnableRoApiAccess("ENABLE_RO_API_ACCESS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ro_api_access", "ENABLE_RO_API_ACCESS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_grid_member.test_extattrs"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_ExternalSyslogBackupServers(t *testing.T) {
	var resourceName = "nios_grid_member.test_external_syslog_backup_servers"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberExternalSyslogBackupServers("EXTERNAL_SYSLOG_BACKUP_SERVERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_syslog_backup_servers", "EXTERNAL_SYSLOG_BACKUP_SERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberExternalSyslogBackupServers("EXTERNAL_SYSLOG_BACKUP_SERVERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_syslog_backup_servers", "EXTERNAL_SYSLOG_BACKUP_SERVERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_ExternalSyslogServerEnable(t *testing.T) {
	var resourceName = "nios_grid_member.test_external_syslog_server_enable"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberExternalSyslogServerEnable("EXTERNAL_SYSLOG_SERVER_ENABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_syslog_server_enable", "EXTERNAL_SYSLOG_SERVER_ENABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberExternalSyslogServerEnable("EXTERNAL_SYSLOG_SERVER_ENABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_syslog_server_enable", "EXTERNAL_SYSLOG_SERVER_ENABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_HaCloudPlatform(t *testing.T) {
	var resourceName = "nios_grid_member.test_ha_cloud_platform"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberHaCloudPlatform("HA_CLOUD_PLATFORM_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ha_cloud_platform", "HA_CLOUD_PLATFORM_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberHaCloudPlatform("HA_CLOUD_PLATFORM_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ha_cloud_platform", "HA_CLOUD_PLATFORM_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_HaOnCloud(t *testing.T) {
	var resourceName = "nios_grid_member.test_ha_on_cloud"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberHaOnCloud("HA_ON_CLOUD_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ha_on_cloud", "HA_ON_CLOUD_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberHaOnCloud("HA_ON_CLOUD_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ha_on_cloud", "HA_ON_CLOUD_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_HostName(t *testing.T) {
	var resourceName = "nios_grid_member.test_host_name"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberHostName("HOST_NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "host_name", "HOST_NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberHostName("HOST_NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "host_name", "HOST_NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_Ipv6Setting(t *testing.T) {
	var resourceName = "nios_grid_member.test_ipv6_setting"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberIpv6Setting("IPV6_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting", "IPV6_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberIpv6Setting("IPV6_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting", "IPV6_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_Ipv6StaticRoutes(t *testing.T) {
	var resourceName = "nios_grid_member.test_ipv6_static_routes"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberIpv6StaticRoutes("IPV6_STATIC_ROUTES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_static_routes", "IPV6_STATIC_ROUTES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberIpv6StaticRoutes("IPV6_STATIC_ROUTES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_static_routes", "IPV6_STATIC_ROUTES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_Lan2Enabled(t *testing.T) {
	var resourceName = "nios_grid_member.test_lan2_enabled"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberLan2Enabled("LAN2_ENABLED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lan2_enabled", "LAN2_ENABLED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberLan2Enabled("LAN2_ENABLED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lan2_enabled", "LAN2_ENABLED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_Lan2PortSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_lan2_port_setting"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberLan2PortSetting("LAN2_PORT_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting", "LAN2_PORT_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberLan2PortSetting("LAN2_PORT_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lan2_port_setting", "LAN2_PORT_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_LomNetworkConfig(t *testing.T) {
	var resourceName = "nios_grid_member.test_lom_network_config"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberLomNetworkConfig("LOM_NETWORK_CONFIG_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lom_network_config", "LOM_NETWORK_CONFIG_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberLomNetworkConfig("LOM_NETWORK_CONFIG_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lom_network_config", "LOM_NETWORK_CONFIG_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_LomUsers(t *testing.T) {
	var resourceName = "nios_grid_member.test_lom_users"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberLomUsers("LOM_USERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lom_users", "LOM_USERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberLomUsers("LOM_USERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lom_users", "LOM_USERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_MasterCandidate(t *testing.T) {
	var resourceName = "nios_grid_member.test_master_candidate"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberMasterCandidate("MASTER_CANDIDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "master_candidate", "MASTER_CANDIDATE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberMasterCandidate("MASTER_CANDIDATE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "master_candidate", "MASTER_CANDIDATE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_MemberServiceCommunication(t *testing.T) {
	var resourceName = "nios_grid_member.test_member_service_communication"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberMemberServiceCommunication("MEMBER_SERVICE_COMMUNICATION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member_service_communication", "MEMBER_SERVICE_COMMUNICATION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberMemberServiceCommunication("MEMBER_SERVICE_COMMUNICATION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member_service_communication", "MEMBER_SERVICE_COMMUNICATION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_MgmtPortSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_mgmt_port_setting"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberMgmtPortSetting("MGMT_PORT_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mgmt_port_setting", "MGMT_PORT_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberMgmtPortSetting("MGMT_PORT_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mgmt_port_setting", "MGMT_PORT_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_NatSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_nat_setting"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberNatSetting("NAT_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nat_setting", "NAT_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberNatSetting("NAT_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nat_setting", "NAT_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_NodeInfo(t *testing.T) {
	var resourceName = "nios_grid_member.test_node_info"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberNodeInfo("NODE_INFO_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "node_info", "NODE_INFO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberNodeInfo("NODE_INFO_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "node_info", "NODE_INFO_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_NtpSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_ntp_setting"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberNtpSetting("NTP_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ntp_setting", "NTP_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberNtpSetting("NTP_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ntp_setting", "NTP_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_OspfList(t *testing.T) {
	var resourceName = "nios_grid_member.test_ospf_list"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberOspfList("OSPF_LIST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ospf_list", "OSPF_LIST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberOspfList("OSPF_LIST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ospf_list", "OSPF_LIST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_PassiveHaArpEnabled(t *testing.T) {
	var resourceName = "nios_grid_member.test_passive_ha_arp_enabled"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberPassiveHaArpEnabled("PASSIVE_HA_ARP_ENABLED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "passive_ha_arp_enabled", "PASSIVE_HA_ARP_ENABLED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberPassiveHaArpEnabled("PASSIVE_HA_ARP_ENABLED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "passive_ha_arp_enabled", "PASSIVE_HA_ARP_ENABLED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_Platform(t *testing.T) {
	var resourceName = "nios_grid_member.test_platform"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberPlatform("PLATFORM_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "platform", "PLATFORM_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberPlatform("PLATFORM_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "platform", "PLATFORM_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_PreProvisioning(t *testing.T) {
	var resourceName = "nios_grid_member.test_pre_provisioning"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberPreProvisioning("PRE_PROVISIONING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pre_provisioning", "PRE_PROVISIONING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberPreProvisioning("PRE_PROVISIONING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pre_provisioning", "PRE_PROVISIONING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_PreserveIfOwnsDelegation(t *testing.T) {
	var resourceName = "nios_grid_member.test_preserve_if_owns_delegation"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberPreserveIfOwnsDelegation("PRESERVE_IF_OWNS_DELEGATION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preserve_if_owns_delegation", "PRESERVE_IF_OWNS_DELEGATION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberPreserveIfOwnsDelegation("PRESERVE_IF_OWNS_DELEGATION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preserve_if_owns_delegation", "PRESERVE_IF_OWNS_DELEGATION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_RemoteConsoleAccessEnable(t *testing.T) {
	var resourceName = "nios_grid_member.test_remote_console_access_enable"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberRemoteConsoleAccessEnable("REMOTE_CONSOLE_ACCESS_ENABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_console_access_enable", "REMOTE_CONSOLE_ACCESS_ENABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberRemoteConsoleAccessEnable("REMOTE_CONSOLE_ACCESS_ENABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_console_access_enable", "REMOTE_CONSOLE_ACCESS_ENABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_RouterId(t *testing.T) {
	var resourceName = "nios_grid_member.test_router_id"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberRouterId("ROUTER_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "router_id", "ROUTER_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberRouterId("ROUTER_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "router_id", "ROUTER_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_ServiceTypeConfiguration(t *testing.T) {
	var resourceName = "nios_grid_member.test_service_type_configuration"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberServiceTypeConfiguration("SERVICE_TYPE_CONFIGURATION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service_type_configuration", "SERVICE_TYPE_CONFIGURATION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberServiceTypeConfiguration("SERVICE_TYPE_CONFIGURATION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service_type_configuration", "SERVICE_TYPE_CONFIGURATION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_SnmpSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_snmp_setting"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberSnmpSetting("SNMP_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "snmp_setting", "SNMP_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberSnmpSetting("SNMP_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "snmp_setting", "SNMP_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_StaticRoutes(t *testing.T) {
	var resourceName = "nios_grid_member.test_static_routes"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberStaticRoutes("STATIC_ROUTES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "static_routes", "STATIC_ROUTES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberStaticRoutes("STATIC_ROUTES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "static_routes", "STATIC_ROUTES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_SupportAccessEnable(t *testing.T) {
	var resourceName = "nios_grid_member.test_support_access_enable"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberSupportAccessEnable("SUPPORT_ACCESS_ENABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "support_access_enable", "SUPPORT_ACCESS_ENABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberSupportAccessEnable("SUPPORT_ACCESS_ENABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "support_access_enable", "SUPPORT_ACCESS_ENABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_SyslogProxySetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_syslog_proxy_setting"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberSyslogProxySetting("SYSLOG_PROXY_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting", "SYSLOG_PROXY_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberSyslogProxySetting("SYSLOG_PROXY_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting", "SYSLOG_PROXY_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_SyslogServers(t *testing.T) {
	var resourceName = "nios_grid_member.test_syslog_servers"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberSyslogServers("SYSLOG_SERVERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers", "SYSLOG_SERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberSyslogServers("SYSLOG_SERVERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers", "SYSLOG_SERVERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_SyslogSize(t *testing.T) {
	var resourceName = "nios_grid_member.test_syslog_size"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberSyslogSize("SYSLOG_SIZE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "syslog_size", "SYSLOG_SIZE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberSyslogSize("SYSLOG_SIZE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "syslog_size", "SYSLOG_SIZE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_ThresholdTraps(t *testing.T) {
	var resourceName = "nios_grid_member.test_threshold_traps"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberThresholdTraps("THRESHOLD_TRAPS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "threshold_traps", "THRESHOLD_TRAPS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberThresholdTraps("THRESHOLD_TRAPS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "threshold_traps", "THRESHOLD_TRAPS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_TimeZone(t *testing.T) {
	var resourceName = "nios_grid_member.test_time_zone"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberTimeZone("TIME_ZONE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "time_zone", "TIME_ZONE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberTimeZone("TIME_ZONE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "time_zone", "TIME_ZONE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_TrafficCaptureAuthDnsSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_traffic_capture_auth_dns_setting"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberTrafficCaptureAuthDnsSetting("TRAFFIC_CAPTURE_AUTH_DNS_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_auth_dns_setting", "TRAFFIC_CAPTURE_AUTH_DNS_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberTrafficCaptureAuthDnsSetting("TRAFFIC_CAPTURE_AUTH_DNS_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_auth_dns_setting", "TRAFFIC_CAPTURE_AUTH_DNS_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_TrafficCaptureChrSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_traffic_capture_chr_setting"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberTrafficCaptureChrSetting("TRAFFIC_CAPTURE_CHR_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_chr_setting", "TRAFFIC_CAPTURE_CHR_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberTrafficCaptureChrSetting("TRAFFIC_CAPTURE_CHR_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_chr_setting", "TRAFFIC_CAPTURE_CHR_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_TrafficCaptureQpsSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_traffic_capture_qps_setting"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberTrafficCaptureQpsSetting("TRAFFIC_CAPTURE_QPS_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_qps_setting", "TRAFFIC_CAPTURE_QPS_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberTrafficCaptureQpsSetting("TRAFFIC_CAPTURE_QPS_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_qps_setting", "TRAFFIC_CAPTURE_QPS_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_TrafficCaptureRecDnsSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_traffic_capture_rec_dns_setting"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberTrafficCaptureRecDnsSetting("TRAFFIC_CAPTURE_REC_DNS_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_rec_dns_setting", "TRAFFIC_CAPTURE_REC_DNS_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberTrafficCaptureRecDnsSetting("TRAFFIC_CAPTURE_REC_DNS_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_rec_dns_setting", "TRAFFIC_CAPTURE_REC_DNS_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_TrafficCaptureRecQueriesSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_traffic_capture_rec_queries_setting"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberTrafficCaptureRecQueriesSetting("TRAFFIC_CAPTURE_REC_QUERIES_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_rec_queries_setting", "TRAFFIC_CAPTURE_REC_QUERIES_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberTrafficCaptureRecQueriesSetting("TRAFFIC_CAPTURE_REC_QUERIES_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_rec_queries_setting", "TRAFFIC_CAPTURE_REC_QUERIES_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_TrapNotifications(t *testing.T) {
	var resourceName = "nios_grid_member.test_trap_notifications"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberTrapNotifications("TRAP_NOTIFICATIONS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "trap_notifications", "TRAP_NOTIFICATIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberTrapNotifications("TRAP_NOTIFICATIONS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "trap_notifications", "TRAP_NOTIFICATIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_UpgradeGroup(t *testing.T) {
	var resourceName = "nios_grid_member.test_upgrade_group"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberUpgradeGroup("UPGRADE_GROUP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "upgrade_group", "UPGRADE_GROUP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUpgradeGroup("UPGRADE_GROUP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "upgrade_group", "UPGRADE_GROUP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_UseAutomatedTrafficCapture(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_automated_traffic_capture"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberUseAutomatedTrafficCapture("USE_AUTOMATED_TRAFFIC_CAPTURE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_automated_traffic_capture", "USE_AUTOMATED_TRAFFIC_CAPTURE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseAutomatedTrafficCapture("USE_AUTOMATED_TRAFFIC_CAPTURE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_automated_traffic_capture", "USE_AUTOMATED_TRAFFIC_CAPTURE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_UseDnsResolverSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_dns_resolver_setting"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberUseDnsResolverSetting("USE_DNS_RESOLVER_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dns_resolver_setting", "USE_DNS_RESOLVER_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseDnsResolverSetting("USE_DNS_RESOLVER_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dns_resolver_setting", "USE_DNS_RESOLVER_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_UseDscp(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_dscp"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberUseDscp("USE_DSCP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dscp", "USE_DSCP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseDscp("USE_DSCP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dscp", "USE_DSCP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_UseEmailSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_email_setting"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberUseEmailSetting("USE_EMAIL_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_email_setting", "USE_EMAIL_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseEmailSetting("USE_EMAIL_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_email_setting", "USE_EMAIL_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_UseEnableLom(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_enable_lom"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberUseEnableLom("USE_ENABLE_LOM_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_lom", "USE_ENABLE_LOM_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseEnableLom("USE_ENABLE_LOM_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_lom", "USE_ENABLE_LOM_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_UseEnableMemberRedirect(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_enable_member_redirect"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberUseEnableMemberRedirect("USE_ENABLE_MEMBER_REDIRECT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_member_redirect", "USE_ENABLE_MEMBER_REDIRECT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseEnableMemberRedirect("USE_ENABLE_MEMBER_REDIRECT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_member_redirect", "USE_ENABLE_MEMBER_REDIRECT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_UseExternalSyslogBackupServers(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_external_syslog_backup_servers"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberUseExternalSyslogBackupServers("USE_EXTERNAL_SYSLOG_BACKUP_SERVERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_external_syslog_backup_servers", "USE_EXTERNAL_SYSLOG_BACKUP_SERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseExternalSyslogBackupServers("USE_EXTERNAL_SYSLOG_BACKUP_SERVERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_external_syslog_backup_servers", "USE_EXTERNAL_SYSLOG_BACKUP_SERVERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_UseRemoteConsoleAccessEnable(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_remote_console_access_enable"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberUseRemoteConsoleAccessEnable("USE_REMOTE_CONSOLE_ACCESS_ENABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_remote_console_access_enable", "USE_REMOTE_CONSOLE_ACCESS_ENABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseRemoteConsoleAccessEnable("USE_REMOTE_CONSOLE_ACCESS_ENABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_remote_console_access_enable", "USE_REMOTE_CONSOLE_ACCESS_ENABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_UseSnmpSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_snmp_setting"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberUseSnmpSetting("USE_SNMP_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp_setting", "USE_SNMP_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseSnmpSetting("USE_SNMP_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp_setting", "USE_SNMP_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_UseSupportAccessEnable(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_support_access_enable"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberUseSupportAccessEnable("USE_SUPPORT_ACCESS_ENABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_support_access_enable", "USE_SUPPORT_ACCESS_ENABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseSupportAccessEnable("USE_SUPPORT_ACCESS_ENABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_support_access_enable", "USE_SUPPORT_ACCESS_ENABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_UseSyslogProxySetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_syslog_proxy_setting"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberUseSyslogProxySetting("USE_SYSLOG_PROXY_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_syslog_proxy_setting", "USE_SYSLOG_PROXY_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseSyslogProxySetting("USE_SYSLOG_PROXY_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_syslog_proxy_setting", "USE_SYSLOG_PROXY_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_UseThresholdTraps(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_threshold_traps"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberUseThresholdTraps("USE_THRESHOLD_TRAPS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_threshold_traps", "USE_THRESHOLD_TRAPS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseThresholdTraps("USE_THRESHOLD_TRAPS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_threshold_traps", "USE_THRESHOLD_TRAPS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_UseTimeZone(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_time_zone"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberUseTimeZone("USE_TIME_ZONE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_time_zone", "USE_TIME_ZONE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseTimeZone("USE_TIME_ZONE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_time_zone", "USE_TIME_ZONE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_UseTrafficCaptureAuthDns(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_traffic_capture_auth_dns"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberUseTrafficCaptureAuthDns("USE_TRAFFIC_CAPTURE_AUTH_DNS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_auth_dns", "USE_TRAFFIC_CAPTURE_AUTH_DNS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseTrafficCaptureAuthDns("USE_TRAFFIC_CAPTURE_AUTH_DNS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_auth_dns", "USE_TRAFFIC_CAPTURE_AUTH_DNS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_UseTrafficCaptureChr(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_traffic_capture_chr"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberUseTrafficCaptureChr("USE_TRAFFIC_CAPTURE_CHR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_chr", "USE_TRAFFIC_CAPTURE_CHR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseTrafficCaptureChr("USE_TRAFFIC_CAPTURE_CHR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_chr", "USE_TRAFFIC_CAPTURE_CHR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_UseTrafficCaptureQps(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_traffic_capture_qps"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberUseTrafficCaptureQps("USE_TRAFFIC_CAPTURE_QPS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_qps", "USE_TRAFFIC_CAPTURE_QPS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseTrafficCaptureQps("USE_TRAFFIC_CAPTURE_QPS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_qps", "USE_TRAFFIC_CAPTURE_QPS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_UseTrafficCaptureRecDns(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_traffic_capture_rec_dns"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberUseTrafficCaptureRecDns("USE_TRAFFIC_CAPTURE_REC_DNS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_rec_dns", "USE_TRAFFIC_CAPTURE_REC_DNS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseTrafficCaptureRecDns("USE_TRAFFIC_CAPTURE_REC_DNS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_rec_dns", "USE_TRAFFIC_CAPTURE_REC_DNS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_UseTrafficCaptureRecQueries(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_traffic_capture_rec_queries"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberUseTrafficCaptureRecQueries("USE_TRAFFIC_CAPTURE_REC_QUERIES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_rec_queries", "USE_TRAFFIC_CAPTURE_REC_QUERIES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseTrafficCaptureRecQueries("USE_TRAFFIC_CAPTURE_REC_QUERIES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_rec_queries", "USE_TRAFFIC_CAPTURE_REC_QUERIES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_UseTrapNotifications(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_trap_notifications"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberUseTrapNotifications("USE_TRAP_NOTIFICATIONS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_trap_notifications", "USE_TRAP_NOTIFICATIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseTrapNotifications("USE_TRAP_NOTIFICATIONS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_trap_notifications", "USE_TRAP_NOTIFICATIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_UseV4Vrrp(t *testing.T) {
	var resourceName = "nios_grid_member.test_use_v4_vrrp"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberUseV4Vrrp("USE_V4_VRRP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_v4_vrrp", "USE_V4_VRRP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseV4Vrrp("USE_V4_VRRP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_v4_vrrp", "USE_V4_VRRP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_VipSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_vip_setting"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberVipSetting("VIP_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vip_setting", "VIP_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberVipSetting("VIP_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vip_setting", "VIP_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_VpnMtu(t *testing.T) {
	var resourceName = "nios_grid_member.test_vpn_mtu"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberVpnMtu("VPN_MTU_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vpn_mtu", "VPN_MTU_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberVpnMtu("VPN_MTU_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vpn_mtu", "VPN_MTU_UPDATE_REPLACE_ME"),
				),
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

func testAccMemberBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return `
resource "nios_grid_member" "test" {
}
`
}

func testAccMemberRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccMemberAdditionalIpList(additionalIpList string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_additional_ip_list" {
    additional_ip_list = %q
}
`, additionalIpList)
}

func testAccMemberAutomatedTrafficCaptureSetting(automatedTrafficCaptureSetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_automated_traffic_capture_setting" {
    automated_traffic_capture_setting = %q
}
`, automatedTrafficCaptureSetting)
}

func testAccMemberBgpAs(bgpAs string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_bgp_as" {
    bgp_as = %q
}
`, bgpAs)
}

func testAccMemberComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccMemberConfigAddrType(configAddrType string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_config_addr_type" {
    config_addr_type = %q
}
`, configAddrType)
}

func testAccMemberCspAccessKey(cspAccessKey string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_csp_access_key" {
    csp_access_key = %q
}
`, cspAccessKey)
}

func testAccMemberCspMemberSetting(cspMemberSetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_csp_member_setting" {
    csp_member_setting = %q
}
`, cspMemberSetting)
}

func testAccMemberDnsResolverSetting(dnsResolverSetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_dns_resolver_setting" {
    dns_resolver_setting = %q
}
`, dnsResolverSetting)
}

func testAccMemberDscp(dscp string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_dscp" {
    dscp = %q
}
`, dscp)
}

func testAccMemberEmailSetting(emailSetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_email_setting" {
    email_setting = %q
}
`, emailSetting)
}

func testAccMemberEnableHa(enableHa string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_enable_ha" {
    enable_ha = %q
}
`, enableHa)
}

func testAccMemberEnableLom(enableLom string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_enable_lom" {
    enable_lom = %q
}
`, enableLom)
}

func testAccMemberEnableMemberRedirect(enableMemberRedirect string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_enable_member_redirect" {
    enable_member_redirect = %q
}
`, enableMemberRedirect)
}

func testAccMemberEnableRoApiAccess(enableRoApiAccess string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_enable_ro_api_access" {
    enable_ro_api_access = %q
}
`, enableRoApiAccess)
}

func testAccMemberExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccMemberExternalSyslogBackupServers(externalSyslogBackupServers string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_external_syslog_backup_servers" {
    external_syslog_backup_servers = %q
}
`, externalSyslogBackupServers)
}

func testAccMemberExternalSyslogServerEnable(externalSyslogServerEnable string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_external_syslog_server_enable" {
    external_syslog_server_enable = %q
}
`, externalSyslogServerEnable)
}

func testAccMemberHaCloudPlatform(haCloudPlatform string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_ha_cloud_platform" {
    ha_cloud_platform = %q
}
`, haCloudPlatform)
}

func testAccMemberHaOnCloud(haOnCloud string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_ha_on_cloud" {
    ha_on_cloud = %q
}
`, haOnCloud)
}

func testAccMemberHostName(hostName string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_host_name" {
    host_name = %q
}
`, hostName)
}

func testAccMemberIpv6Setting(ipv6Setting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_ipv6_setting" {
    ipv6_setting = %q
}
`, ipv6Setting)
}

func testAccMemberIpv6StaticRoutes(ipv6StaticRoutes string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_ipv6_static_routes" {
    ipv6_static_routes = %q
}
`, ipv6StaticRoutes)
}

func testAccMemberLan2Enabled(lan2Enabled string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_lan2_enabled" {
    lan2_enabled = %q
}
`, lan2Enabled)
}

func testAccMemberLan2PortSetting(lan2PortSetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_lan2_port_setting" {
    lan2_port_setting = %q
}
`, lan2PortSetting)
}

func testAccMemberLomNetworkConfig(lomNetworkConfig string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_lom_network_config" {
    lom_network_config = %q
}
`, lomNetworkConfig)
}

func testAccMemberLomUsers(lomUsers string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_lom_users" {
    lom_users = %q
}
`, lomUsers)
}

func testAccMemberMasterCandidate(masterCandidate string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_master_candidate" {
    master_candidate = %q
}
`, masterCandidate)
}

func testAccMemberMemberServiceCommunication(memberServiceCommunication string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_member_service_communication" {
    member_service_communication = %q
}
`, memberServiceCommunication)
}

func testAccMemberMgmtPortSetting(mgmtPortSetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_mgmt_port_setting" {
    mgmt_port_setting = %q
}
`, mgmtPortSetting)
}

func testAccMemberNatSetting(natSetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_nat_setting" {
    nat_setting = %q
}
`, natSetting)
}

func testAccMemberNodeInfo(nodeInfo string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_node_info" {
    node_info = %q
}
`, nodeInfo)
}

func testAccMemberNtpSetting(ntpSetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_ntp_setting" {
    ntp_setting = %q
}
`, ntpSetting)
}

func testAccMemberOspfList(ospfList string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_ospf_list" {
    ospf_list = %q
}
`, ospfList)
}

func testAccMemberPassiveHaArpEnabled(passiveHaArpEnabled string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_passive_ha_arp_enabled" {
    passive_ha_arp_enabled = %q
}
`, passiveHaArpEnabled)
}

func testAccMemberPlatform(platform string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_platform" {
    platform = %q
}
`, platform)
}

func testAccMemberPreProvisioning(preProvisioning string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_pre_provisioning" {
    pre_provisioning = %q
}
`, preProvisioning)
}

func testAccMemberPreserveIfOwnsDelegation(preserveIfOwnsDelegation string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_preserve_if_owns_delegation" {
    preserve_if_owns_delegation = %q
}
`, preserveIfOwnsDelegation)
}

func testAccMemberRemoteConsoleAccessEnable(remoteConsoleAccessEnable string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_remote_console_access_enable" {
    remote_console_access_enable = %q
}
`, remoteConsoleAccessEnable)
}

func testAccMemberRouterId(routerId string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_router_id" {
    router_id = %q
}
`, routerId)
}

func testAccMemberServiceTypeConfiguration(serviceTypeConfiguration string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_service_type_configuration" {
    service_type_configuration = %q
}
`, serviceTypeConfiguration)
}

func testAccMemberSnmpSetting(snmpSetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_snmp_setting" {
    snmp_setting = %q
}
`, snmpSetting)
}

func testAccMemberStaticRoutes(staticRoutes string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_static_routes" {
    static_routes = %q
}
`, staticRoutes)
}

func testAccMemberSupportAccessEnable(supportAccessEnable string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_support_access_enable" {
    support_access_enable = %q
}
`, supportAccessEnable)
}

func testAccMemberSyslogProxySetting(syslogProxySetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_syslog_proxy_setting" {
    syslog_proxy_setting = %q
}
`, syslogProxySetting)
}

func testAccMemberSyslogServers(syslogServers string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_syslog_servers" {
    syslog_servers = %q
}
`, syslogServers)
}

func testAccMemberSyslogSize(syslogSize string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_syslog_size" {
    syslog_size = %q
}
`, syslogSize)
}

func testAccMemberThresholdTraps(thresholdTraps string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_threshold_traps" {
    threshold_traps = %q
}
`, thresholdTraps)
}

func testAccMemberTimeZone(timeZone string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_time_zone" {
    time_zone = %q
}
`, timeZone)
}

func testAccMemberTrafficCaptureAuthDnsSetting(trafficCaptureAuthDnsSetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_traffic_capture_auth_dns_setting" {
    traffic_capture_auth_dns_setting = %q
}
`, trafficCaptureAuthDnsSetting)
}

func testAccMemberTrafficCaptureChrSetting(trafficCaptureChrSetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_traffic_capture_chr_setting" {
    traffic_capture_chr_setting = %q
}
`, trafficCaptureChrSetting)
}

func testAccMemberTrafficCaptureQpsSetting(trafficCaptureQpsSetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_traffic_capture_qps_setting" {
    traffic_capture_qps_setting = %q
}
`, trafficCaptureQpsSetting)
}

func testAccMemberTrafficCaptureRecDnsSetting(trafficCaptureRecDnsSetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_traffic_capture_rec_dns_setting" {
    traffic_capture_rec_dns_setting = %q
}
`, trafficCaptureRecDnsSetting)
}

func testAccMemberTrafficCaptureRecQueriesSetting(trafficCaptureRecQueriesSetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_traffic_capture_rec_queries_setting" {
    traffic_capture_rec_queries_setting = %q
}
`, trafficCaptureRecQueriesSetting)
}

func testAccMemberTrapNotifications(trapNotifications string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_trap_notifications" {
    trap_notifications = %q
}
`, trapNotifications)
}

func testAccMemberUpgradeGroup(upgradeGroup string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_upgrade_group" {
    upgrade_group = %q
}
`, upgradeGroup)
}

func testAccMemberUseAutomatedTrafficCapture(useAutomatedTrafficCapture string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_automated_traffic_capture" {
    use_automated_traffic_capture = %q
}
`, useAutomatedTrafficCapture)
}

func testAccMemberUseDnsResolverSetting(useDnsResolverSetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_dns_resolver_setting" {
    use_dns_resolver_setting = %q
}
`, useDnsResolverSetting)
}

func testAccMemberUseDscp(useDscp string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_dscp" {
    use_dscp = %q
}
`, useDscp)
}

func testAccMemberUseEmailSetting(useEmailSetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_email_setting" {
    use_email_setting = %q
}
`, useEmailSetting)
}

func testAccMemberUseEnableLom(useEnableLom string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_enable_lom" {
    use_enable_lom = %q
}
`, useEnableLom)
}

func testAccMemberUseEnableMemberRedirect(useEnableMemberRedirect string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_enable_member_redirect" {
    use_enable_member_redirect = %q
}
`, useEnableMemberRedirect)
}

func testAccMemberUseExternalSyslogBackupServers(useExternalSyslogBackupServers string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_external_syslog_backup_servers" {
    use_external_syslog_backup_servers = %q
}
`, useExternalSyslogBackupServers)
}

func testAccMemberUseRemoteConsoleAccessEnable(useRemoteConsoleAccessEnable string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_remote_console_access_enable" {
    use_remote_console_access_enable = %q
}
`, useRemoteConsoleAccessEnable)
}

func testAccMemberUseSnmpSetting(useSnmpSetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_snmp_setting" {
    use_snmp_setting = %q
}
`, useSnmpSetting)
}

func testAccMemberUseSupportAccessEnable(useSupportAccessEnable string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_support_access_enable" {
    use_support_access_enable = %q
}
`, useSupportAccessEnable)
}

func testAccMemberUseSyslogProxySetting(useSyslogProxySetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_syslog_proxy_setting" {
    use_syslog_proxy_setting = %q
}
`, useSyslogProxySetting)
}

func testAccMemberUseThresholdTraps(useThresholdTraps string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_threshold_traps" {
    use_threshold_traps = %q
}
`, useThresholdTraps)
}

func testAccMemberUseTimeZone(useTimeZone string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_time_zone" {
    use_time_zone = %q
}
`, useTimeZone)
}

func testAccMemberUseTrafficCaptureAuthDns(useTrafficCaptureAuthDns string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_traffic_capture_auth_dns" {
    use_traffic_capture_auth_dns = %q
}
`, useTrafficCaptureAuthDns)
}

func testAccMemberUseTrafficCaptureChr(useTrafficCaptureChr string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_traffic_capture_chr" {
    use_traffic_capture_chr = %q
}
`, useTrafficCaptureChr)
}

func testAccMemberUseTrafficCaptureQps(useTrafficCaptureQps string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_traffic_capture_qps" {
    use_traffic_capture_qps = %q
}
`, useTrafficCaptureQps)
}

func testAccMemberUseTrafficCaptureRecDns(useTrafficCaptureRecDns string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_traffic_capture_rec_dns" {
    use_traffic_capture_rec_dns = %q
}
`, useTrafficCaptureRecDns)
}

func testAccMemberUseTrafficCaptureRecQueries(useTrafficCaptureRecQueries string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_traffic_capture_rec_queries" {
    use_traffic_capture_rec_queries = %q
}
`, useTrafficCaptureRecQueries)
}

func testAccMemberUseTrapNotifications(useTrapNotifications string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_trap_notifications" {
    use_trap_notifications = %q
}
`, useTrapNotifications)
}

func testAccMemberUseV4Vrrp(useV4Vrrp string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_v4_vrrp" {
    use_v4_vrrp = %q
}
`, useV4Vrrp)
}

func testAccMemberVipSetting(vipSetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_vip_setting" {
    vip_setting = %q
}
`, vipSetting)
}

func testAccMemberVpnMtu(vpnMtu string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_vpn_mtu" {
    vpn_mtu = %q
}
`, vpnMtu)
}
