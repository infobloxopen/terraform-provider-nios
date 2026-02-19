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

func TestAccMemberResource_basic(t *testing.T) {
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

// func TestAccMemberResource_disappears(t *testing.T) {
// 	resourceName := "nios_grid_member.test"
// 	var v grid.Member

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		CheckDestroy:             testAccCheckMemberDestroy(context.Background(), &v),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccMemberBasicConfig("HOST_NAME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckMemberExists(context.Background(), resourceName, &v),
// 					testAccCheckMemberDisappears(context.Background(), &v),
// 				),
// 				ExpectNonEmptyPlan: true,
// 			},
// 		},
// 	})
// }

func TestAccMemberResource_Import(t *testing.T) {
	var resourceName = "nios_grid_member.test"
	var v grid.Member

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				//Config: testAccMemberBasicConfig("HOST_NAME_REPLACE_ME"),
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
				ImportStateVerifyIdentifierAttribute: "ref",
				PlanOnly:                             true,
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccMemberImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_AdditionalIpList(t *testing.T) {
	var resourceName = "nios_grid_member.test_additional_ip_list"
	var v grid.Member
	additionalIpListVal := []map[string]any{}
	additionalIpListValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberAdditionalIpList("HOST_NAME_REPLACE_ME", additionalIpListVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "additional_ip_list", "ADDITIONAL_IP_LIST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberAdditionalIpList("HOST_NAME_REPLACE_ME", additionalIpListValUpdated),
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
	automatedTrafficCaptureSettingVal := map[string]any{}
	automatedTrafficCaptureSettingValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberAutomatedTrafficCaptureSetting("HOST_NAME_REPLACE_ME", automatedTrafficCaptureSettingVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "automated_traffic_capture_setting", "AUTOMATED_TRAFFIC_CAPTURE_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberAutomatedTrafficCaptureSetting("HOST_NAME_REPLACE_ME", automatedTrafficCaptureSettingValUpdated),
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
	bgpAsVal := []map[string]any{}
	bgpAsValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberBgpAs("HOST_NAME_REPLACE_ME", bgpAsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bgp_as", "BGP_AS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberBgpAs("HOST_NAME_REPLACE_ME", bgpAsValUpdated),
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
				Config: testAccMemberComment("HOST_NAME_REPLACE_ME", "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberComment("HOST_NAME_REPLACE_ME", "Updated comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the object"),
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
				Config: testAccMemberConfigAddrType("CONFIG_ADDR_TYPE_REPLACE_ME", "BOTH"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "config_addr_type", "BOTH"),
				),
			},
			{
				Config: testAccMemberConfigAddrType("CONFIG_ADDR_TYPE_REPLACE_ME", "IPV4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "config_addr_type", "IPV4"),
				),
			},
			{
				Config: testAccMemberConfigAddrType("CONFIG_ADDR_TYPE_REPLACE_ME", "IPV6"),
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

func TestAccMemberResource_CspMemberSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_csp_member_setting"
	var v grid.Member
	cspMemberSettingVal := map[string]any{}
	cspMemberSettingValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberCspMemberSetting("HOST_NAME_REPLACE_ME", cspMemberSettingVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "csp_member_setting", "CSP_MEMBER_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberCspMemberSetting("HOST_NAME_REPLACE_ME", cspMemberSettingValUpdated),
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
	dnsResolverSettingVal := map[string]any{}
	dnsResolverSettingValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberDnsResolverSetting("HOST_NAME_REPLACE_ME", dnsResolverSettingVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_resolver_setting", "DNS_RESOLVER_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberDnsResolverSetting("HOST_NAME_REPLACE_ME", dnsResolverSettingValUpdated),
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
				Config: testAccMemberDscp("HOST_NAME_REPLACE_ME", "DSCP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dscp", "DSCP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberDscp("HOST_NAME_REPLACE_ME", "DSCP_UPDATE_REPLACE_ME"),
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
	emailSettingVal := map[string]any{}
	emailSettingValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberEmailSetting("HOST_NAME_REPLACE_ME", emailSettingVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_setting", "EMAIL_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberEmailSetting("HOST_NAME_REPLACE_ME", emailSettingValUpdated),
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
				Config: testAccMemberEnableHa("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ha", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberEnableHa("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ha", "false"),
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
				Config: testAccMemberEnableLom("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_lom", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberEnableLom("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_lom", "false"),
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
				Config: testAccMemberEnableMemberRedirect("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_member_redirect", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberEnableMemberRedirect("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_member_redirect", "false"),
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
				Config: testAccMemberEnableRoApiAccess("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ro_api_access", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberEnableRoApiAccess("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ro_api_access", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_grid_member.test_extattrs"
	var v grid.Member
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberExtAttrs("HOST_NAME_REPLACE_ME", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccMemberExtAttrs("HOST_NAME_REPLACE_ME", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_ExternalSyslogBackupServers(t *testing.T) {
	var resourceName = "nios_grid_member.test_external_syslog_backup_servers"
	var v grid.Member
	externalSyslogBackupServersVal := []map[string]any{}
	externalSyslogBackupServersValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberExternalSyslogBackupServers("HOST_NAME_REPLACE_ME", externalSyslogBackupServersVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_syslog_backup_servers", "EXTERNAL_SYSLOG_BACKUP_SERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberExternalSyslogBackupServers("HOST_NAME_REPLACE_ME", externalSyslogBackupServersValUpdated),
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
				Config: testAccMemberExternalSyslogServerEnable("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_syslog_server_enable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberExternalSyslogServerEnable("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_syslog_server_enable", "false"),
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

func TestAccMemberResource_HaOnCloud(t *testing.T) {
	var resourceName = "nios_grid_member.test_ha_on_cloud"
	var v grid.Member

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
				Config: testAccMemberHostName("HOST_NAME_REPLACE_ME"),
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
	ipv6SettingVal := map[string]any{}
	ipv6SettingValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberIpv6Setting("HOST_NAME_REPLACE_ME", ipv6SettingVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_setting", "IPV6_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberIpv6Setting("HOST_NAME_REPLACE_ME", ipv6SettingValUpdated),
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

func TestAccMemberResource_LomUsers(t *testing.T) {
	var resourceName = "nios_grid_member.test_lom_users"
	var v grid.Member
	lomUsersVal := []map[string]any{}
	lomUsersValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberLomUsers("HOST_NAME_REPLACE_ME", lomUsersVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lom_users", "LOM_USERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberLomUsers("HOST_NAME_REPLACE_ME", lomUsersValUpdated),
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
				Config: testAccMemberMasterCandidate("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "master_candidate", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberMasterCandidate("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "master_candidate", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_MemberServiceCommunication(t *testing.T) {
	var resourceName = "nios_grid_member.test_member_service_communication"
	var v grid.Member
	memberServiceCommunicationVal := []map[string]any{}
	memberServiceCommunicationValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberMemberServiceCommunication("HOST_NAME_REPLACE_ME", memberServiceCommunicationVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member_service_communication", "MEMBER_SERVICE_COMMUNICATION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberMemberServiceCommunication("HOST_NAME_REPLACE_ME", memberServiceCommunicationValUpdated),
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
func TestAccMemberResource_NatSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_nat_setting"
	var v grid.Member
	natSettingVal := map[string]any{}
	natSettingValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberNatSetting("HOST_NAME_REPLACE_ME", natSettingVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nat_setting", "NAT_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberNatSetting("HOST_NAME_REPLACE_ME", natSettingValUpdated),
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
	nodeInfoVal := []map[string]any{}
	nodeInfoValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberNodeInfo("HOST_NAME_REPLACE_ME", nodeInfoVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "node_info", "NODE_INFO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberNodeInfo("HOST_NAME_REPLACE_ME", nodeInfoValUpdated),
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
	ntpSettingVal := map[string]any{}
	ntpSettingValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberNtpSetting("HOST_NAME_REPLACE_ME", ntpSettingVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ntp_setting", "NTP_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberNtpSetting("HOST_NAME_REPLACE_ME", ntpSettingValUpdated),
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
	ospfListVal := []map[string]any{}
	ospfListValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberOspfList("HOST_NAME_REPLACE_ME", ospfListVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ospf_list", "OSPF_LIST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberOspfList("HOST_NAME_REPLACE_ME", ospfListValUpdated),
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
				Config: testAccMemberPassiveHaArpEnabled("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "passive_ha_arp_enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberPassiveHaArpEnabled("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "passive_ha_arp_enabled", "false"),
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
				Config: testAccMemberPlatform("PLATFORM_REPLACE_ME", "CISCO"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "platform", "CISCO"),
				),
			},
			{
				Config: testAccMemberPlatform("PLATFORM_REPLACE_ME", "IBVM"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "platform", "IBVM"),
				),
			},
			{
				Config: testAccMemberPlatform("PLATFORM_REPLACE_ME", "INFOBLOX"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "platform", "INFOBLOX"),
				),
			},
			{
				Config: testAccMemberPlatform("PLATFORM_REPLACE_ME", "RIVERBED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "platform", "RIVERBED"),
				),
			},
			{
				Config: testAccMemberPlatform("PLATFORM_REPLACE_ME", "VNIOS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "platform", "VNIOS"),
				),
			},
		},
	})
}

func TestAccMemberResource_PreProvisioning(t *testing.T) {
	var resourceName = "nios_grid_member.test_pre_provisioning"
	var v grid.Member
	preProvisioningVal := map[string]any{}
	preProvisioningValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberPreProvisioning("HOST_NAME_REPLACE_ME", preProvisioningVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pre_provisioning", "PRE_PROVISIONING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberPreProvisioning("HOST_NAME_REPLACE_ME", preProvisioningValUpdated),
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
				Config: testAccMemberPreserveIfOwnsDelegation("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preserve_if_owns_delegation", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberPreserveIfOwnsDelegation("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preserve_if_owns_delegation", "false"),
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
				Config: testAccMemberRemoteConsoleAccessEnable("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_console_access_enable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberRemoteConsoleAccessEnable("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_console_access_enable", "false"),
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
				Config: testAccMemberRouterId("HOST_NAME_REPLACE_ME", "ROUTER_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "router_id", "ROUTER_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberRouterId("HOST_NAME_REPLACE_ME", "ROUTER_ID_UPDATE_REPLACE_ME"),
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
				Config: testAccMemberServiceTypeConfiguration("SERVICE_TYPE_CONFIGURATION_REPLACE_ME", "ALL_V4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service_type_configuration", "ALL_V4"),
				),
			},
			{
				Config: testAccMemberServiceTypeConfiguration("SERVICE_TYPE_CONFIGURATION_REPLACE_ME", "ALL_V6"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service_type_configuration", "ALL_V6"),
				),
			},
			{
				Config: testAccMemberServiceTypeConfiguration("SERVICE_TYPE_CONFIGURATION_REPLACE_ME", "CUSTOM"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service_type_configuration", "CUSTOM"),
				),
			},
		},
	})
}

func TestAccMemberResource_SnmpSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_snmp_setting"
	var v grid.Member
	snmpSettingVal := map[string]any{}
	snmpSettingValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberSnmpSetting("HOST_NAME_REPLACE_ME", snmpSettingVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "snmp_setting", "SNMP_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberSnmpSetting("HOST_NAME_REPLACE_ME", snmpSettingValUpdated),
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
	staticRoutesVal := []map[string]any{}
	staticRoutesValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberStaticRoutes("HOST_NAME_REPLACE_ME", staticRoutesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "static_routes", "STATIC_ROUTES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberStaticRoutes("HOST_NAME_REPLACE_ME", staticRoutesValUpdated),
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
				Config: testAccMemberSupportAccessEnable("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "support_access_enable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberSupportAccessEnable("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "support_access_enable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_SyslogProxySetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_syslog_proxy_setting"
	var v grid.Member
	syslogProxySettingVal := map[string]any{}
	syslogProxySettingValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberSyslogProxySetting("HOST_NAME_REPLACE_ME", syslogProxySettingVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "syslog_proxy_setting", "SYSLOG_PROXY_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberSyslogProxySetting("HOST_NAME_REPLACE_ME", syslogProxySettingValUpdated),
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
	syslogServersVal := []map[string]any{}
	syslogServersValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberSyslogServers("HOST_NAME_REPLACE_ME", syslogServersVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "syslog_servers", "SYSLOG_SERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberSyslogServers("HOST_NAME_REPLACE_ME", syslogServersValUpdated),
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
				Config: testAccMemberSyslogSize("HOST_NAME_REPLACE_ME", "SYSLOG_SIZE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "syslog_size", "SYSLOG_SIZE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberSyslogSize("HOST_NAME_REPLACE_ME", "SYSLOG_SIZE_UPDATE_REPLACE_ME"),
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
	thresholdTrapsVal := []map[string]any{}
	thresholdTrapsValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberThresholdTraps("HOST_NAME_REPLACE_ME", thresholdTrapsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "threshold_traps", "THRESHOLD_TRAPS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberThresholdTraps("HOST_NAME_REPLACE_ME", thresholdTrapsValUpdated),
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
				Config: testAccMemberTimeZone("HOST_NAME_REPLACE_ME", "TIME_ZONE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "time_zone", "TIME_ZONE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberTimeZone("HOST_NAME_REPLACE_ME", "TIME_ZONE_UPDATE_REPLACE_ME"),
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
	trafficCaptureAuthDnsSettingVal := map[string]any{}
	trafficCaptureAuthDnsSettingValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberTrafficCaptureAuthDnsSetting("HOST_NAME_REPLACE_ME", trafficCaptureAuthDnsSettingVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_auth_dns_setting", "TRAFFIC_CAPTURE_AUTH_DNS_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberTrafficCaptureAuthDnsSetting("HOST_NAME_REPLACE_ME", trafficCaptureAuthDnsSettingValUpdated),
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
	trafficCaptureChrSettingVal := map[string]any{}
	trafficCaptureChrSettingValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberTrafficCaptureChrSetting("HOST_NAME_REPLACE_ME", trafficCaptureChrSettingVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_chr_setting", "TRAFFIC_CAPTURE_CHR_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberTrafficCaptureChrSetting("HOST_NAME_REPLACE_ME", trafficCaptureChrSettingValUpdated),
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
	trafficCaptureQpsSettingVal := map[string]any{}
	trafficCaptureQpsSettingValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberTrafficCaptureQpsSetting("HOST_NAME_REPLACE_ME", trafficCaptureQpsSettingVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_qps_setting", "TRAFFIC_CAPTURE_QPS_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberTrafficCaptureQpsSetting("HOST_NAME_REPLACE_ME", trafficCaptureQpsSettingValUpdated),
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
	trafficCaptureRecDnsSettingVal := map[string]any{}
	trafficCaptureRecDnsSettingValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberTrafficCaptureRecDnsSetting("HOST_NAME_REPLACE_ME", trafficCaptureRecDnsSettingVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_rec_dns_setting", "TRAFFIC_CAPTURE_REC_DNS_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberTrafficCaptureRecDnsSetting("HOST_NAME_REPLACE_ME", trafficCaptureRecDnsSettingValUpdated),
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
	trafficCaptureRecQueriesSettingVal := map[string]any{}
	trafficCaptureRecQueriesSettingValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberTrafficCaptureRecQueriesSetting("HOST_NAME_REPLACE_ME", trafficCaptureRecQueriesSettingVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "traffic_capture_rec_queries_setting", "TRAFFIC_CAPTURE_REC_QUERIES_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberTrafficCaptureRecQueriesSetting("HOST_NAME_REPLACE_ME", trafficCaptureRecQueriesSettingValUpdated),
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
	trapNotificationsVal := []map[string]any{}
	trapNotificationsValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberTrapNotifications("HOST_NAME_REPLACE_ME", trapNotificationsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "trap_notifications", "TRAP_NOTIFICATIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberTrapNotifications("HOST_NAME_REPLACE_ME", trapNotificationsValUpdated),
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
				Config: testAccMemberUpgradeGroup("HOST_NAME_REPLACE_ME", "UPGRADE_GROUP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "upgrade_group", "UPGRADE_GROUP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUpgradeGroup("HOST_NAME_REPLACE_ME", "UPGRADE_GROUP_UPDATE_REPLACE_ME"),
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
				Config: testAccMemberUseAutomatedTrafficCapture("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_automated_traffic_capture", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseAutomatedTrafficCapture("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_automated_traffic_capture", "false"),
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
				Config: testAccMemberUseDnsResolverSetting("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dns_resolver_setting", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseDnsResolverSetting("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dns_resolver_setting", "false"),
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
				Config: testAccMemberUseDscp("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dscp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseDscp("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dscp", "false"),
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
				Config: testAccMemberUseEmailSetting("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_email_setting", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseEmailSetting("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_email_setting", "false"),
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
				Config: testAccMemberUseEnableLom("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_lom", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseEnableLom("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_lom", "false"),
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
				Config: testAccMemberUseEnableMemberRedirect("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_member_redirect", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseEnableMemberRedirect("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_member_redirect", "false"),
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
				Config: testAccMemberUseExternalSyslogBackupServers("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_external_syslog_backup_servers", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseExternalSyslogBackupServers("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_external_syslog_backup_servers", "false"),
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
				Config: testAccMemberUseRemoteConsoleAccessEnable("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_remote_console_access_enable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseRemoteConsoleAccessEnable("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_remote_console_access_enable", "false"),
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
				Config: testAccMemberUseSnmpSetting("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp_setting", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseSnmpSetting("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_snmp_setting", "false"),
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
				Config: testAccMemberUseSupportAccessEnable("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_support_access_enable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseSupportAccessEnable("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_support_access_enable", "false"),
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
				Config: testAccMemberUseSyslogProxySetting("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_syslog_proxy_setting", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseSyslogProxySetting("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_syslog_proxy_setting", "false"),
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
				Config: testAccMemberUseThresholdTraps("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_threshold_traps", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseThresholdTraps("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_threshold_traps", "false"),
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
				Config: testAccMemberUseTimeZone("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_time_zone", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseTimeZone("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_time_zone", "false"),
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
				Config: testAccMemberUseTrafficCaptureAuthDns("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_auth_dns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseTrafficCaptureAuthDns("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_auth_dns", "false"),
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
				Config: testAccMemberUseTrafficCaptureChr("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_chr", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseTrafficCaptureChr("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_chr", "false"),
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
				Config: testAccMemberUseTrafficCaptureQps("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_qps", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseTrafficCaptureQps("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_qps", "false"),
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
				Config: testAccMemberUseTrafficCaptureRecDns("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_rec_dns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseTrafficCaptureRecDns("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_rec_dns", "false"),
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
				Config: testAccMemberUseTrafficCaptureRecQueries("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_rec_queries", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseTrafficCaptureRecQueries("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_traffic_capture_rec_queries", "false"),
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
				Config: testAccMemberUseTrapNotifications("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_trap_notifications", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseTrapNotifications("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_trap_notifications", "false"),
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
				Config: testAccMemberUseV4Vrrp("HOST_NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_v4_vrrp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberUseV4Vrrp("HOST_NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_v4_vrrp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMemberResource_VipSetting(t *testing.T) {
	var resourceName = "nios_grid_member.test_vip_setting"
	var v grid.Member
	vipSettingVal := map[string]any{}
	vipSettingValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMemberVipSetting("HOST_NAME_REPLACE_ME", vipSettingVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vip_setting", "VIP_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberVipSetting("HOST_NAME_REPLACE_ME", vipSettingValUpdated),
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
				Config: testAccMemberVpnMtu("HOST_NAME_REPLACE_ME", "VPN_MTU_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vpn_mtu", "VPN_MTU_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMemberVpnMtu("HOST_NAME_REPLACE_ME", "VPN_MTU_UPDATE_REPLACE_ME"),
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

func testAccMemberImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccMemberAdditionalIpList(hostName string, additionalIpList []map[string]any) string {
	additionalIpListStr := utils.ConvertSliceOfMapsToHCL(additionalIpList)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_additional_ip_list" {
    host_name = %q
    additional_ip_list = %s
}
`, hostName, additionalIpListStr)
}

func testAccMemberAutomatedTrafficCaptureSetting(hostName string, automatedTrafficCaptureSetting map[string]any) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_automated_traffic_capture_setting" {
    host_name = %q
    automated_traffic_capture_setting = %s
    use_automated_traffic_capture = true
}
`, hostName, automatedTrafficCaptureSetting)
}

func testAccMemberBgpAs(hostName string, bgpAs []map[string]any) string {
	bgpAsStr := utils.ConvertSliceOfMapsToHCL(bgpAs)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_bgp_as" {
    host_name = %q
    bgp_as = %s
}
`, hostName, bgpAsStr)
}

func testAccMemberComment(hostName string, comment string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_comment" {
    host_name = %q
    comment = %q
}
`, hostName, comment)
}

func testAccMemberConfigAddrType(hostName string, configAddrType string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_config_addr_type" {
    host_name = %q
    config_addr_type = %q
}
`, hostName, configAddrType)
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

func testAccMemberCspMemberSetting(hostName string, cspMemberSetting map[string]any) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_csp_member_setting" {
    host_name = %q
    csp_member_setting = %s
}
`, hostName, cspMemberSetting)
}

func testAccMemberDnsResolverSetting(hostName string, dnsResolverSetting map[string]any) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_dns_resolver_setting" {
    host_name = %q
    dns_resolver_setting = %s
    use_dns_resolver_setting = true
}
`, hostName, dnsResolverSetting)
}

func testAccMemberDscp(hostName string, dscp string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_dscp" {
    host_name = %q
    dscp = %q
    use_dscp = true
}
`, hostName, dscp)
}

func testAccMemberEmailSetting(hostName string, emailSetting map[string]any) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_email_setting" {
    host_name = %q
    email_setting = %s
    use_email_setting = true
}
`, hostName, emailSetting)
}

func testAccMemberEnableHa(hostName string, enableHa string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_enable_ha" {
    host_name = %q
    enable_ha = %q
}
`, hostName, enableHa)
}

func testAccMemberEnableLom(hostName string, enableLom string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_enable_lom" {
    host_name = %q
    enable_lom = %q
    use_enable_lom = true
}
`, hostName, enableLom)
}

func testAccMemberEnableMemberRedirect(hostName string, enableMemberRedirect string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_enable_member_redirect" {
    host_name = %q
    enable_member_redirect = %q
    use_enable_member_redirect = true
}
`, hostName, enableMemberRedirect)
}

func testAccMemberEnableRoApiAccess(hostName string, enableRoApiAccess string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_enable_ro_api_access" {
    host_name = %q
    enable_ro_api_access = %q
}
`, hostName, enableRoApiAccess)
}

func testAccMemberExtAttrs(hostName string, extAttrs map[string]string) string {
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	return fmt.Sprintf(`
resource "nios_grid_member" "test_extattrs" {
    host_name = %q
    extattrs = %s
}
`, hostName, extAttrsStr)
}

func testAccMemberExternalSyslogBackupServers(hostName string, externalSyslogBackupServers []map[string]any) string {
	externalSyslogBackupServersStr := utils.ConvertSliceOfMapsToHCL(externalSyslogBackupServers)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_external_syslog_backup_servers" {
    host_name = %q
    external_syslog_backup_servers = %s
    use_external_syslog_backup_servers = true
}
`, hostName, externalSyslogBackupServersStr)
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

func testAccMemberHostName(hostName string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_host_name" {
    host_name = %q
}
`, hostName)
}

func testAccMemberIpv6Setting(hostName string, ipv6Setting map[string]any) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_ipv6_setting" {
    host_name = %q
    ipv6_setting = %s
}
`, hostName, ipv6Setting)
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

func testAccMemberLomUsers(hostName string, lomUsers []map[string]any) string {
	lomUsersStr := utils.ConvertSliceOfMapsToHCL(lomUsers)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_lom_users" {
    host_name = %q
    lom_users = %s
}
`, hostName, lomUsersStr)
}

func testAccMemberMasterCandidate(hostName string, masterCandidate string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_master_candidate" {
    host_name = %q
    master_candidate = %q
}
`, hostName, masterCandidate)
}

func testAccMemberMemberServiceCommunication(hostName string, memberServiceCommunication []map[string]any) string {
	memberServiceCommunicationStr := utils.ConvertSliceOfMapsToHCL(memberServiceCommunication)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_member_service_communication" {
    host_name = %q
    member_service_communication = %s
}
`, hostName, memberServiceCommunicationStr)
}

func testAccMemberMgmtPortSetting(hostName string, mgmtPortSetting map[string]any) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_mgmt_port_setting" {
    host_name = %q
    mgmt_port_setting = %s
}
`, hostName, mgmtPortSetting)
}

func testAccMemberNatSetting(hostName string, natSetting map[string]any) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_nat_setting" {
    host_name = %q
    nat_setting = %s
}
`, hostName, natSetting)
}

func testAccMemberNodeInfo(hostName string, nodeInfo []map[string]any) string {
	nodeInfoStr := utils.ConvertSliceOfMapsToHCL(nodeInfo)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_node_info" {
    host_name = %q
    node_info = %s
}
`, hostName, nodeInfoStr)
}

func testAccMemberNtpSetting(hostName string, ntpSetting map[string]any) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_ntp_setting" {
    host_name = %q
    ntp_setting = %s
}
`, hostName, ntpSetting)
}

func testAccMemberOspfList(hostName string, ospfList []map[string]any) string {
	ospfListStr := utils.ConvertSliceOfMapsToHCL(ospfList)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_ospf_list" {
    host_name = %q
    ospf_list = %s
}
`, hostName, ospfListStr)
}

func testAccMemberPassiveHaArpEnabled(hostName string, passiveHaArpEnabled string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_passive_ha_arp_enabled" {
    host_name = %q
    passive_ha_arp_enabled = %q
}
`, hostName, passiveHaArpEnabled)
}

func testAccMemberPlatform(hostName string, platform string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_platform" {
    host_name = %q
    platform = %q
}
`, hostName, platform)
}

func testAccMemberPreProvisioning(hostName string, preProvisioning map[string]any) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_pre_provisioning" {
    host_name = %q
    pre_provisioning = %s
}
`, hostName, preProvisioning)
}

func testAccMemberPreserveIfOwnsDelegation(hostName string, preserveIfOwnsDelegation string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_preserve_if_owns_delegation" {
    host_name = %q
    preserve_if_owns_delegation = %q
}
`, hostName, preserveIfOwnsDelegation)
}

func testAccMemberRemoteConsoleAccessEnable(hostName string, remoteConsoleAccessEnable string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_remote_console_access_enable" {
    host_name = %q
    remote_console_access_enable = %q
    use_remote_console_access_enable = true
}
`, hostName, remoteConsoleAccessEnable)
}

func testAccMemberRouterId(hostName string, routerId string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_router_id" {
    host_name = %q
    router_id = %q
}
`, hostName, routerId)
}

func testAccMemberServiceTypeConfiguration(hostName string, serviceTypeConfiguration string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_service_type_configuration" {
    host_name = %q
    service_type_configuration = %q
}
`, hostName, serviceTypeConfiguration)
}

func testAccMemberSnmpSetting(hostName string, snmpSetting map[string]any) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_snmp_setting" {
    host_name = %q
    snmp_setting = %s
    use_snmp_setting = true
}
`, hostName, snmpSetting)
}

func testAccMemberStaticRoutes(hostName string, staticRoutes []map[string]any) string {
	staticRoutesStr := utils.ConvertSliceOfMapsToHCL(staticRoutes)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_static_routes" {
    host_name = %q
    static_routes = %s
}
`, hostName, staticRoutesStr)
}

func testAccMemberSupportAccessEnable(hostName string, supportAccessEnable string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_support_access_enable" {
    host_name = %q
    support_access_enable = %q
    use_support_access_enable = true
}
`, hostName, supportAccessEnable)
}

func testAccMemberSyslogProxySetting(hostName string, syslogProxySetting map[string]any) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_syslog_proxy_setting" {
    host_name = %q
    syslog_proxy_setting = %s
    use_syslog_proxy_setting = true
}
`, hostName, syslogProxySetting)
}

func testAccMemberSyslogServers(hostName string, syslogServers []map[string]any) string {
	syslogServersStr := utils.ConvertSliceOfMapsToHCL(syslogServers)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_syslog_servers" {
    host_name = %q
    syslog_servers = %s
    use_syslog_proxy_setting = true
}
`, hostName, syslogServersStr)
}

func testAccMemberSyslogSize(hostName string, syslogSize string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_syslog_size" {
    host_name = %q
    syslog_size = %q
    use_syslog_proxy_setting = true
}
`, hostName, syslogSize)
}

func testAccMemberThresholdTraps(hostName string, thresholdTraps []map[string]any) string {
	thresholdTrapsStr := utils.ConvertSliceOfMapsToHCL(thresholdTraps)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_threshold_traps" {
    host_name = %q
    threshold_traps = %s
    use_threshold_traps = true
}
`, hostName, thresholdTrapsStr)
}

func testAccMemberTimeZone(hostName string, timeZone string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_time_zone" {
    host_name = %q
    time_zone = %q
    use_time_zone = true
}
`, hostName, timeZone)
}

func testAccMemberTrafficCaptureAuthDnsSetting(hostName string, trafficCaptureAuthDnsSetting map[string]any) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_traffic_capture_auth_dns_setting" {
    host_name = %q
    traffic_capture_auth_dns_setting = %s
    use_traffic_capture_auth_dns = true
}
`, hostName, trafficCaptureAuthDnsSetting)
}

func testAccMemberTrafficCaptureChrSetting(hostName string, trafficCaptureChrSetting map[string]any) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_traffic_capture_chr_setting" {
    host_name = %q
    traffic_capture_chr_setting = %s
    use_traffic_capture_chr = true
}
`, hostName, trafficCaptureChrSetting)
}

func testAccMemberTrafficCaptureQpsSetting(hostName string, trafficCaptureQpsSetting map[string]any) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_traffic_capture_qps_setting" {
    host_name = %q
    traffic_capture_qps_setting = %s
    use_traffic_capture_qps = true
}
`, hostName, trafficCaptureQpsSetting)
}

func testAccMemberTrafficCaptureRecDnsSetting(hostName string, trafficCaptureRecDnsSetting map[string]any) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_traffic_capture_rec_dns_setting" {
    host_name = %q
    traffic_capture_rec_dns_setting = %s
    use_traffic_capture_rec_dns = true
}
`, hostName, trafficCaptureRecDnsSetting)
}

func testAccMemberTrafficCaptureRecQueriesSetting(hostName string, trafficCaptureRecQueriesSetting map[string]any) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_traffic_capture_rec_queries_setting" {
    host_name = %q
    traffic_capture_rec_queries_setting = %s
    use_traffic_capture_rec_queries = true
}
`, hostName, trafficCaptureRecQueriesSetting)
}

func testAccMemberTrapNotifications(hostName string, trapNotifications []map[string]any) string {
	trapNotificationsStr := utils.ConvertSliceOfMapsToHCL(trapNotifications)
	return fmt.Sprintf(`
resource "nios_grid_member" "test_trap_notifications" {
    host_name = %q
    trap_notifications = %s
    use_trap_notifications = true
}
`, hostName, trapNotificationsStr)
}

func testAccMemberUpgradeGroup(hostName string, upgradeGroup string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_upgrade_group" {
    host_name = %q
    upgrade_group = %q
}
`, hostName, upgradeGroup)
}

func testAccMemberUseAutomatedTrafficCapture(hostName string, useAutomatedTrafficCapture string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_automated_traffic_capture" {
    host_name = %q
    use_automated_traffic_capture = %q
}
`, hostName, useAutomatedTrafficCapture)
}

func testAccMemberUseDnsResolverSetting(hostName string, useDnsResolverSetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_dns_resolver_setting" {
    host_name = %q
    use_dns_resolver_setting = %q
}
`, hostName, useDnsResolverSetting)
}

func testAccMemberUseDscp(hostName string, useDscp string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_dscp" {
    host_name = %q
    use_dscp = %q
}
`, hostName, useDscp)
}

func testAccMemberUseEmailSetting(hostName string, useEmailSetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_email_setting" {
    host_name = %q
    use_email_setting = %q
}
`, hostName, useEmailSetting)
}

func testAccMemberUseEnableLom(hostName string, useEnableLom string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_enable_lom" {
    host_name = %q
    use_enable_lom = %q
}
`, hostName, useEnableLom)
}

func testAccMemberUseEnableMemberRedirect(hostName string, useEnableMemberRedirect string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_enable_member_redirect" {
    host_name = %q
    use_enable_member_redirect = %q
}
`, hostName, useEnableMemberRedirect)
}

func testAccMemberUseExternalSyslogBackupServers(hostName string, useExternalSyslogBackupServers string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_external_syslog_backup_servers" {
    host_name = %q
    use_external_syslog_backup_servers = %q
}
`, hostName, useExternalSyslogBackupServers)
}

func testAccMemberUseRemoteConsoleAccessEnable(hostName string, useRemoteConsoleAccessEnable string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_remote_console_access_enable" {
    host_name = %q
    use_remote_console_access_enable = %q
}
`, hostName, useRemoteConsoleAccessEnable)
}

func testAccMemberUseSnmpSetting(hostName string, useSnmpSetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_snmp_setting" {
    host_name = %q
    use_snmp_setting = %q
}
`, hostName, useSnmpSetting)
}

func testAccMemberUseSupportAccessEnable(hostName string, useSupportAccessEnable string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_support_access_enable" {
    host_name = %q
    use_support_access_enable = %q
}
`, hostName, useSupportAccessEnable)
}

func testAccMemberUseSyslogProxySetting(hostName string, useSyslogProxySetting string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_syslog_proxy_setting" {
    host_name = %q
    use_syslog_proxy_setting = %q
}
`, hostName, useSyslogProxySetting)
}

func testAccMemberUseThresholdTraps(hostName string, useThresholdTraps string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_threshold_traps" {
    host_name = %q
    use_threshold_traps = %q
}
`, hostName, useThresholdTraps)
}

func testAccMemberUseTimeZone(hostName string, useTimeZone string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_time_zone" {
    host_name = %q
    use_time_zone = %q
}
`, hostName, useTimeZone)
}

func testAccMemberUseTrafficCaptureAuthDns(hostName string, useTrafficCaptureAuthDns string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_traffic_capture_auth_dns" {
    host_name = %q
    use_traffic_capture_auth_dns = %q
}
`, hostName, useTrafficCaptureAuthDns)
}

func testAccMemberUseTrafficCaptureChr(hostName string, useTrafficCaptureChr string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_traffic_capture_chr" {
    host_name = %q
    use_traffic_capture_chr = %q
}
`, hostName, useTrafficCaptureChr)
}

func testAccMemberUseTrafficCaptureQps(hostName string, useTrafficCaptureQps string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_traffic_capture_qps" {
    host_name = %q
    use_traffic_capture_qps = %q
}
`, hostName, useTrafficCaptureQps)
}

func testAccMemberUseTrafficCaptureRecDns(hostName string, useTrafficCaptureRecDns string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_traffic_capture_rec_dns" {
    host_name = %q
    use_traffic_capture_rec_dns = %q
}
`, hostName, useTrafficCaptureRecDns)
}

func testAccMemberUseTrafficCaptureRecQueries(hostName string, useTrafficCaptureRecQueries string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_traffic_capture_rec_queries" {
    host_name = %q
    use_traffic_capture_rec_queries = %q
}
`, hostName, useTrafficCaptureRecQueries)
}

func testAccMemberUseTrapNotifications(hostName string, useTrapNotifications string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_trap_notifications" {
    host_name = %q
    use_trap_notifications = %q
}
`, hostName, useTrapNotifications)
}

func testAccMemberUseV4Vrrp(hostName string, useV4Vrrp string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_use_v4_vrrp" {
    host_name = %q
    use_v4_vrrp = %q
}
`, hostName, useV4Vrrp)
}

func testAccMemberVipSetting(hostName string, vipSetting map[string]any) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_vip_setting" {
    host_name = %q
    vip_setting = %s
}
`, hostName, vipSetting)
}

func testAccMemberVpnMtu(hostName string, vpnMtu string) string {
	return fmt.Sprintf(`
resource "nios_grid_member" "test_vpn_mtu" {
    host_name = %q
    vpn_mtu = %q
}
`, hostName, vpnMtu)
}
