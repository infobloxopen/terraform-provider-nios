package ipam_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/ipam"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/acctest"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// TODO : Add readable attributes for the resource
var readableAttributesForNetworkcontainer = "authority,bootfile,bootserver,cloud_info,comment,ddns_domainname,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,ddns_update_fixed_addresses,ddns_use_option81,deny_bootp,discover_now_status,discovery_basic_poll_settings,discovery_blackout_setting,discovery_engine_type,discovery_member,email_list,enable_ddns,enable_dhcp_thresholds,enable_discovery,enable_email_warnings,enable_pxe_lease_time,enable_snmp_warnings,endpoint_sources,extattrs,federated_realms,high_water_mark,high_water_mark_reset,ignore_dhcp_option_list_request,ignore_id,ignore_mac_addresses,ipam_email_addresses,ipam_threshold_settings,ipam_trap_settings,last_rir_registration_update_sent,last_rir_registration_update_status,lease_scavenge_time,logic_filter_rules,low_water_mark,low_water_mark_reset,mgm_private,mgm_private_overridable,ms_ad_user_data,network,network_container,network_view,nextserver,options,port_control_blackout_setting,pxe_lease_time,recycle_leases,rir,rir_organization,rir_registration_status,same_port_control_discovery_blackout,subscribe_settings,unmanaged,update_dns_on_lease_renewal,use_authority,use_blackout_setting,use_bootfile,use_bootserver,use_ddns_domainname,use_ddns_generate_hostname,use_ddns_ttl,use_ddns_update_fixed_addresses,use_ddns_use_option81,use_deny_bootp,use_discovery_basic_polling_settings,use_email_list,use_enable_ddns,use_enable_dhcp_thresholds,use_enable_discovery,use_ignore_dhcp_option_list_request,use_ignore_id,use_ipam_email_addresses,use_ipam_threshold_settings,use_ipam_trap_settings,use_lease_scavenge_time,use_logic_filter_rules,use_mgm_private,use_nextserver,use_options,use_pxe_lease_time,use_recycle_leases,use_subscribe_settings,use_update_dns_on_lease_renewal,use_zone_associations,utilization,zone_associations"

func TestAccNetworkcontainerResource_basic(t *testing.T) {
	var resourceName = "nios_nios_ipam_networkcontainer.test"
	var v ipam.Networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerBasicConfig("11.0.0.0/16"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "network", "11.0.0.0/16"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_disappears(t *testing.T) {
	resourceName := "nios_nios_ipam_networkcontainer.test"
	var v ipam.Networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetworkcontainerDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkcontainerBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					testAccCheckNetworkcontainerDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// func TestAccNetworkcontainerResource_Ref(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test__ref"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerRef("REF_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "_ref", "REF_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerRef("REF_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "_ref", "REF_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_Authority(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_authority"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerAuthority("AUTHORITY_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "authority", "AUTHORITY_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerAuthority("AUTHORITY_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "authority", "AUTHORITY_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_AutoCreateReversezone(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_auto_create_reversezone"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerAutoCreateReversezone("AUTO_CREATE_REVERSEZONE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "auto_create_reversezone", "AUTO_CREATE_REVERSEZONE_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerAutoCreateReversezone("AUTO_CREATE_REVERSEZONE_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "auto_create_reversezone", "AUTO_CREATE_REVERSEZONE_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_Bootfile(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_bootfile"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerBootfile("BOOTFILE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "bootfile", "BOOTFILE_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerBootfile("BOOTFILE_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "bootfile", "BOOTFILE_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_Bootserver(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_bootserver"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerBootserver("BOOTSERVER_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "bootserver", "BOOTSERVER_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerBootserver("BOOTSERVER_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "bootserver", "BOOTSERVER_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_CloudInfo(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_cloud_info"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerCloudInfo("CLOUD_INFO_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "cloud_info", "CLOUD_INFO_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerCloudInfo("CLOUD_INFO_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "cloud_info", "CLOUD_INFO_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_Comment(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_comment"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerComment("COMMENT_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerComment("COMMENT_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_DdnsDomainname(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_ddns_domainname"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerDdnsDomainname("DDNS_DOMAINNAME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "DDNS_DOMAINNAME_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerDdnsDomainname("DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_DdnsGenerateHostname(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_ddns_generate_hostname"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerDdnsGenerateHostname("DDNS_GENERATE_HOSTNAME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "DDNS_GENERATE_HOSTNAME_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerDdnsGenerateHostname("DDNS_GENERATE_HOSTNAME_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "DDNS_GENERATE_HOSTNAME_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_DdnsServerAlwaysUpdates(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_ddns_server_always_updates"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerDdnsServerAlwaysUpdates("DDNS_SERVER_ALWAYS_UPDATES_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "DDNS_SERVER_ALWAYS_UPDATES_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerDdnsServerAlwaysUpdates("DDNS_SERVER_ALWAYS_UPDATES_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "DDNS_SERVER_ALWAYS_UPDATES_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_DdnsTtl(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_ddns_ttl"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerDdnsTtl("DDNS_TTL_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "DDNS_TTL_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerDdnsTtl("DDNS_TTL_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "DDNS_TTL_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_DdnsUpdateFixedAddresses(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_ddns_update_fixed_addresses"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerDdnsUpdateFixedAddresses("DDNS_UPDATE_FIXED_ADDRESSES_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_update_fixed_addresses", "DDNS_UPDATE_FIXED_ADDRESSES_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerDdnsUpdateFixedAddresses("DDNS_UPDATE_FIXED_ADDRESSES_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_update_fixed_addresses", "DDNS_UPDATE_FIXED_ADDRESSES_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_DdnsUseOption81(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_ddns_use_option81"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerDdnsUseOption81("DDNS_USE_OPTION81_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "DDNS_USE_OPTION81_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerDdnsUseOption81("DDNS_USE_OPTION81_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "DDNS_USE_OPTION81_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_DeleteReason(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_delete_reason"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerDeleteReason("DELETE_REASON_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "delete_reason", "DELETE_REASON_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerDeleteReason("DELETE_REASON_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "delete_reason", "DELETE_REASON_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_DenyBootp(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_deny_bootp"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerDenyBootp("DENY_BOOTP_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "DENY_BOOTP_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerDenyBootp("DENY_BOOTP_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "DENY_BOOTP_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_DiscoveryBasicPollSettings(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_discovery_basic_poll_settings"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerDiscoveryBasicPollSettings("DISCOVERY_BASIC_POLL_SETTINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings", "DISCOVERY_BASIC_POLL_SETTINGS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerDiscoveryBasicPollSettings("DISCOVERY_BASIC_POLL_SETTINGS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings", "DISCOVERY_BASIC_POLL_SETTINGS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_DiscoveryBlackoutSetting(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_discovery_blackout_setting"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerDiscoveryBlackoutSetting("DISCOVERY_BLACKOUT_SETTING_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "discovery_blackout_setting", "DISCOVERY_BLACKOUT_SETTING_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerDiscoveryBlackoutSetting("DISCOVERY_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "discovery_blackout_setting", "DISCOVERY_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_DiscoveryMember(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_discovery_member"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerDiscoveryMember("DISCOVERY_MEMBER_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "discovery_member", "DISCOVERY_MEMBER_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerDiscoveryMember("DISCOVERY_MEMBER_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "discovery_member", "DISCOVERY_MEMBER_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_EmailList(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_email_list"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerEmailList("EMAIL_LIST_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "email_list", "EMAIL_LIST_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerEmailList("EMAIL_LIST_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "email_list", "EMAIL_LIST_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_EnableDdns(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_enable_ddns"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerEnableDdns("ENABLE_DDNS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "ENABLE_DDNS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerEnableDdns("ENABLE_DDNS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "ENABLE_DDNS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_EnableDhcpThresholds(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_enable_dhcp_thresholds"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerEnableDhcpThresholds("ENABLE_DHCP_THRESHOLDS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_dhcp_thresholds", "ENABLE_DHCP_THRESHOLDS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerEnableDhcpThresholds("ENABLE_DHCP_THRESHOLDS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_dhcp_thresholds", "ENABLE_DHCP_THRESHOLDS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_EnableDiscovery(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_enable_discovery"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerEnableDiscovery("ENABLE_DISCOVERY_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_discovery", "ENABLE_DISCOVERY_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerEnableDiscovery("ENABLE_DISCOVERY_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_discovery", "ENABLE_DISCOVERY_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_EnableEmailWarnings(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_enable_email_warnings"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerEnableEmailWarnings("ENABLE_EMAIL_WARNINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_email_warnings", "ENABLE_EMAIL_WARNINGS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerEnableEmailWarnings("ENABLE_EMAIL_WARNINGS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_email_warnings", "ENABLE_EMAIL_WARNINGS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_EnableImmediateDiscovery(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_enable_immediate_discovery"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerEnableImmediateDiscovery("ENABLE_IMMEDIATE_DISCOVERY_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_immediate_discovery", "ENABLE_IMMEDIATE_DISCOVERY_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerEnableImmediateDiscovery("ENABLE_IMMEDIATE_DISCOVERY_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_immediate_discovery", "ENABLE_IMMEDIATE_DISCOVERY_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_EnablePxeLeaseTime(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_enable_pxe_lease_time"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerEnablePxeLeaseTime("ENABLE_PXE_LEASE_TIME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "ENABLE_PXE_LEASE_TIME_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerEnablePxeLeaseTime("ENABLE_PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "ENABLE_PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_EnableSnmpWarnings(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_enable_snmp_warnings"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerEnableSnmpWarnings("ENABLE_SNMP_WARNINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_snmp_warnings", "ENABLE_SNMP_WARNINGS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerEnableSnmpWarnings("ENABLE_SNMP_WARNINGS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_snmp_warnings", "ENABLE_SNMP_WARNINGS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_ExtAttrs(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_extattrs"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerExtAttrs("EXT_ATTRS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_FederatedRealms(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_federated_realms"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerFederatedRealms("FEDERATED_REALMS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "federated_realms", "FEDERATED_REALMS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerFederatedRealms("FEDERATED_REALMS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "federated_realms", "FEDERATED_REALMS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_HighWaterMark(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_high_water_mark"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerHighWaterMark("HIGH_WATER_MARK_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "HIGH_WATER_MARK_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerHighWaterMark("HIGH_WATER_MARK_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "HIGH_WATER_MARK_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_HighWaterMarkReset(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_high_water_mark_reset"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerHighWaterMarkReset("HIGH_WATER_MARK_RESET_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "HIGH_WATER_MARK_RESET_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerHighWaterMarkReset("HIGH_WATER_MARK_RESET_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "HIGH_WATER_MARK_RESET_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_IgnoreDhcpOptionListRequest(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_ignore_dhcp_option_list_request"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerIgnoreDhcpOptionListRequest("IGNORE_DHCP_OPTION_LIST_REQUEST_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "IGNORE_DHCP_OPTION_LIST_REQUEST_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerIgnoreDhcpOptionListRequest("IGNORE_DHCP_OPTION_LIST_REQUEST_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "IGNORE_DHCP_OPTION_LIST_REQUEST_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_IgnoreId(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_ignore_id"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerIgnoreId("IGNORE_ID_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ignore_id", "IGNORE_ID_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerIgnoreId("IGNORE_ID_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ignore_id", "IGNORE_ID_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_IgnoreMacAddresses(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_ignore_mac_addresses"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerIgnoreMacAddresses("IGNORE_MAC_ADDRESSES_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ignore_mac_addresses", "IGNORE_MAC_ADDRESSES_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerIgnoreMacAddresses("IGNORE_MAC_ADDRESSES_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ignore_mac_addresses", "IGNORE_MAC_ADDRESSES_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_IpamEmailAddresses(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_ipam_email_addresses"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerIpamEmailAddresses("IPAM_EMAIL_ADDRESSES_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ipam_email_addresses", "IPAM_EMAIL_ADDRESSES_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerIpamEmailAddresses("IPAM_EMAIL_ADDRESSES_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ipam_email_addresses", "IPAM_EMAIL_ADDRESSES_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_IpamThresholdSettings(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_ipam_threshold_settings"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerIpamThresholdSettings("IPAM_THRESHOLD_SETTINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ipam_threshold_settings", "IPAM_THRESHOLD_SETTINGS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerIpamThresholdSettings("IPAM_THRESHOLD_SETTINGS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ipam_threshold_settings", "IPAM_THRESHOLD_SETTINGS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_IpamTrapSettings(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_ipam_trap_settings"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerIpamTrapSettings("IPAM_TRAP_SETTINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ipam_trap_settings", "IPAM_TRAP_SETTINGS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerIpamTrapSettings("IPAM_TRAP_SETTINGS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ipam_trap_settings", "IPAM_TRAP_SETTINGS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_LeaseScavengeTime(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_lease_scavenge_time"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerLeaseScavengeTime("LEASE_SCAVENGE_TIME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "LEASE_SCAVENGE_TIME_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerLeaseScavengeTime("LEASE_SCAVENGE_TIME_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "LEASE_SCAVENGE_TIME_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_LogicFilterRules(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_logic_filter_rules"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerLogicFilterRules("LOGIC_FILTER_RULES_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerLogicFilterRules("LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_LowWaterMark(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_low_water_mark"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerLowWaterMark("LOW_WATER_MARK_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "LOW_WATER_MARK_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerLowWaterMark("LOW_WATER_MARK_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "LOW_WATER_MARK_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_LowWaterMarkReset(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_low_water_mark_reset"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerLowWaterMarkReset("LOW_WATER_MARK_RESET_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "LOW_WATER_MARK_RESET_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerLowWaterMarkReset("LOW_WATER_MARK_RESET_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "LOW_WATER_MARK_RESET_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_MgmPrivate(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_mgm_private"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerMgmPrivate("MGM_PRIVATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "mgm_private", "MGM_PRIVATE_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerMgmPrivate("MGM_PRIVATE_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "mgm_private", "MGM_PRIVATE_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_MsAdUserData(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_ms_ad_user_data"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerMsAdUserData("MS_AD_USER_DATA_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ms_ad_user_data", "MS_AD_USER_DATA_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerMsAdUserData("MS_AD_USER_DATA_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ms_ad_user_data", "MS_AD_USER_DATA_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_Network(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_network"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerNetwork("NETWORK_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "network", "NETWORK_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerNetwork("NETWORK_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "network", "NETWORK_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_FuncCall(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_func_call"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerFuncCall("FUNC_CALL_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "func_call", "FUNC_CALL_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerFuncCall("FUNC_CALL_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "func_call", "FUNC_CALL_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_NetworkView(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_network_view"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerNetworkView("NETWORK_VIEW_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerNetworkView("NETWORK_VIEW_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_Nextserver(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_nextserver"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerNextserver("NEXTSERVER_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "nextserver", "NEXTSERVER_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerNextserver("NEXTSERVER_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "nextserver", "NEXTSERVER_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_Options(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_options"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerOptions("OPTIONS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerOptions("OPTIONS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_PortControlBlackoutSetting(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_port_control_blackout_setting"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerPortControlBlackoutSetting("PORT_CONTROL_BLACKOUT_SETTING_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "port_control_blackout_setting", "PORT_CONTROL_BLACKOUT_SETTING_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerPortControlBlackoutSetting("PORT_CONTROL_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "port_control_blackout_setting", "PORT_CONTROL_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_PxeLeaseTime(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_pxe_lease_time"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerPxeLeaseTime("PXE_LEASE_TIME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "PXE_LEASE_TIME_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerPxeLeaseTime("PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_RecycleLeases(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_recycle_leases"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerRecycleLeases("RECYCLE_LEASES_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "RECYCLE_LEASES_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerRecycleLeases("RECYCLE_LEASES_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "RECYCLE_LEASES_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_RemoveSubnets(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_remove_subnets"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerRemoveSubnets("REMOVE_SUBNETS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "remove_subnets", "REMOVE_SUBNETS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerRemoveSubnets("REMOVE_SUBNETS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "remove_subnets", "REMOVE_SUBNETS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_RestartIfNeeded(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_restart_if_needed"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerRestartIfNeeded("RESTART_IF_NEEDED_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "restart_if_needed", "RESTART_IF_NEEDED_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerRestartIfNeeded("RESTART_IF_NEEDED_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "restart_if_needed", "RESTART_IF_NEEDED_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_RirOrganization(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_rir_organization"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerRirOrganization("RIR_ORGANIZATION_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "rir_organization", "RIR_ORGANIZATION_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerRirOrganization("RIR_ORGANIZATION_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "rir_organization", "RIR_ORGANIZATION_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_RirRegistrationAction(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_rir_registration_action"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerRirRegistrationAction("RIR_REGISTRATION_ACTION_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "rir_registration_action", "RIR_REGISTRATION_ACTION_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerRirRegistrationAction("RIR_REGISTRATION_ACTION_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "rir_registration_action", "RIR_REGISTRATION_ACTION_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_RirRegistrationStatus(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_rir_registration_status"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerRirRegistrationStatus("RIR_REGISTRATION_STATUS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "rir_registration_status", "RIR_REGISTRATION_STATUS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerRirRegistrationStatus("RIR_REGISTRATION_STATUS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "rir_registration_status", "RIR_REGISTRATION_STATUS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_SamePortControlDiscoveryBlackout(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_same_port_control_discovery_blackout"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerSamePortControlDiscoveryBlackout("SAME_PORT_CONTROL_DISCOVERY_BLACKOUT_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "SAME_PORT_CONTROL_DISCOVERY_BLACKOUT_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerSamePortControlDiscoveryBlackout("SAME_PORT_CONTROL_DISCOVERY_BLACKOUT_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "SAME_PORT_CONTROL_DISCOVERY_BLACKOUT_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_SendRirRequest(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_send_rir_request"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerSendRirRequest("SEND_RIR_REQUEST_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "send_rir_request", "SEND_RIR_REQUEST_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerSendRirRequest("SEND_RIR_REQUEST_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "send_rir_request", "SEND_RIR_REQUEST_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_SubscribeSettings(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_subscribe_settings"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerSubscribeSettings("SUBSCRIBE_SETTINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "subscribe_settings", "SUBSCRIBE_SETTINGS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerSubscribeSettings("SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "subscribe_settings", "SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_Unmanaged(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_unmanaged"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUnmanaged("UNMANAGED_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "unmanaged", "UNMANAGED_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUnmanaged("UNMANAGED_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "unmanaged", "UNMANAGED_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UpdateDnsOnLeaseRenewal(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_update_dns_on_lease_renewal"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUpdateDnsOnLeaseRenewal("UPDATE_DNS_ON_LEASE_RENEWAL_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "UPDATE_DNS_ON_LEASE_RENEWAL_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUpdateDnsOnLeaseRenewal("UPDATE_DNS_ON_LEASE_RENEWAL_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "UPDATE_DNS_ON_LEASE_RENEWAL_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseAuthority(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_authority"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseAuthority("USE_AUTHORITY_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_authority", "USE_AUTHORITY_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseAuthority("USE_AUTHORITY_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_authority", "USE_AUTHORITY_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseBlackoutSetting(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_blackout_setting"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseBlackoutSetting("USE_BLACKOUT_SETTING_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "USE_BLACKOUT_SETTING_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseBlackoutSetting("USE_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "USE_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseBootfile(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_bootfile"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseBootfile("USE_BOOTFILE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "USE_BOOTFILE_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseBootfile("USE_BOOTFILE_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "USE_BOOTFILE_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseBootserver(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_bootserver"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseBootserver("USE_BOOTSERVER_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "USE_BOOTSERVER_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseBootserver("USE_BOOTSERVER_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "USE_BOOTSERVER_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseDdnsDomainname(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_ddns_domainname"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseDdnsDomainname("USE_DDNS_DOMAINNAME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "USE_DDNS_DOMAINNAME_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseDdnsDomainname("USE_DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "USE_DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseDdnsGenerateHostname(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_ddns_generate_hostname"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseDdnsGenerateHostname("USE_DDNS_GENERATE_HOSTNAME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "USE_DDNS_GENERATE_HOSTNAME_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseDdnsGenerateHostname("USE_DDNS_GENERATE_HOSTNAME_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "USE_DDNS_GENERATE_HOSTNAME_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseDdnsTtl(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_ddns_ttl"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseDdnsTtl("USE_DDNS_TTL_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "USE_DDNS_TTL_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseDdnsTtl("USE_DDNS_TTL_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "USE_DDNS_TTL_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseDdnsUpdateFixedAddresses(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_ddns_update_fixed_addresses"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseDdnsUpdateFixedAddresses("USE_DDNS_UPDATE_FIXED_ADDRESSES_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ddns_update_fixed_addresses", "USE_DDNS_UPDATE_FIXED_ADDRESSES_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseDdnsUpdateFixedAddresses("USE_DDNS_UPDATE_FIXED_ADDRESSES_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ddns_update_fixed_addresses", "USE_DDNS_UPDATE_FIXED_ADDRESSES_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseDdnsUseOption81(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_ddns_use_option81"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseDdnsUseOption81("USE_DDNS_USE_OPTION81_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "USE_DDNS_USE_OPTION81_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseDdnsUseOption81("USE_DDNS_USE_OPTION81_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "USE_DDNS_USE_OPTION81_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseDenyBootp(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_deny_bootp"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseDenyBootp("USE_DENY_BOOTP_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "USE_DENY_BOOTP_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseDenyBootp("USE_DENY_BOOTP_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "USE_DENY_BOOTP_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseDiscoveryBasicPollingSettings(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_discovery_basic_polling_settings"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseDiscoveryBasicPollingSettings("USE_DISCOVERY_BASIC_POLLING_SETTINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "USE_DISCOVERY_BASIC_POLLING_SETTINGS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseDiscoveryBasicPollingSettings("USE_DISCOVERY_BASIC_POLLING_SETTINGS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "USE_DISCOVERY_BASIC_POLLING_SETTINGS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseEmailList(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_email_list"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseEmailList("USE_EMAIL_LIST_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_email_list", "USE_EMAIL_LIST_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseEmailList("USE_EMAIL_LIST_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_email_list", "USE_EMAIL_LIST_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseEnableDdns(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_enable_ddns"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseEnableDdns("USE_ENABLE_DDNS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "USE_ENABLE_DDNS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseEnableDdns("USE_ENABLE_DDNS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "USE_ENABLE_DDNS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseEnableDhcpThresholds(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_enable_dhcp_thresholds"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseEnableDhcpThresholds("USE_ENABLE_DHCP_THRESHOLDS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "USE_ENABLE_DHCP_THRESHOLDS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseEnableDhcpThresholds("USE_ENABLE_DHCP_THRESHOLDS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "USE_ENABLE_DHCP_THRESHOLDS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseEnableDiscovery(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_enable_discovery"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseEnableDiscovery("USE_ENABLE_DISCOVERY_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "USE_ENABLE_DISCOVERY_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseEnableDiscovery("USE_ENABLE_DISCOVERY_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "USE_ENABLE_DISCOVERY_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseIgnoreDhcpOptionListRequest(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_ignore_dhcp_option_list_request"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseIgnoreDhcpOptionListRequest("USE_IGNORE_DHCP_OPTION_LIST_REQUEST_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "USE_IGNORE_DHCP_OPTION_LIST_REQUEST_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseIgnoreDhcpOptionListRequest("USE_IGNORE_DHCP_OPTION_LIST_REQUEST_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "USE_IGNORE_DHCP_OPTION_LIST_REQUEST_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseIgnoreId(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_ignore_id"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseIgnoreId("USE_IGNORE_ID_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ignore_id", "USE_IGNORE_ID_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseIgnoreId("USE_IGNORE_ID_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ignore_id", "USE_IGNORE_ID_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseIpamEmailAddresses(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_ipam_email_addresses"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseIpamEmailAddresses("USE_IPAM_EMAIL_ADDRESSES_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ipam_email_addresses", "USE_IPAM_EMAIL_ADDRESSES_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseIpamEmailAddresses("USE_IPAM_EMAIL_ADDRESSES_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ipam_email_addresses", "USE_IPAM_EMAIL_ADDRESSES_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseIpamThresholdSettings(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_ipam_threshold_settings"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseIpamThresholdSettings("USE_IPAM_THRESHOLD_SETTINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ipam_threshold_settings", "USE_IPAM_THRESHOLD_SETTINGS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseIpamThresholdSettings("USE_IPAM_THRESHOLD_SETTINGS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ipam_threshold_settings", "USE_IPAM_THRESHOLD_SETTINGS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseIpamTrapSettings(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_ipam_trap_settings"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseIpamTrapSettings("USE_IPAM_TRAP_SETTINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ipam_trap_settings", "USE_IPAM_TRAP_SETTINGS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseIpamTrapSettings("USE_IPAM_TRAP_SETTINGS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ipam_trap_settings", "USE_IPAM_TRAP_SETTINGS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseLeaseScavengeTime(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_lease_scavenge_time"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseLeaseScavengeTime("USE_LEASE_SCAVENGE_TIME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "USE_LEASE_SCAVENGE_TIME_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseLeaseScavengeTime("USE_LEASE_SCAVENGE_TIME_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "USE_LEASE_SCAVENGE_TIME_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseLogicFilterRules(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_logic_filter_rules"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseLogicFilterRules("USE_LOGIC_FILTER_RULES_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "USE_LOGIC_FILTER_RULES_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseLogicFilterRules("USE_LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "USE_LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseMgmPrivate(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_mgm_private"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseMgmPrivate("USE_MGM_PRIVATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_mgm_private", "USE_MGM_PRIVATE_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseMgmPrivate("USE_MGM_PRIVATE_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_mgm_private", "USE_MGM_PRIVATE_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseNextserver(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_nextserver"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseNextserver("USE_NEXTSERVER_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "USE_NEXTSERVER_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseNextserver("USE_NEXTSERVER_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "USE_NEXTSERVER_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseOptions(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_options"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseOptions("USE_OPTIONS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_options", "USE_OPTIONS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseOptions("USE_OPTIONS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_options", "USE_OPTIONS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UsePxeLeaseTime(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_pxe_lease_time"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUsePxeLeaseTime("USE_PXE_LEASE_TIME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "USE_PXE_LEASE_TIME_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUsePxeLeaseTime("USE_PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "USE_PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseRecycleLeases(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_recycle_leases"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseRecycleLeases("USE_RECYCLE_LEASES_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "USE_RECYCLE_LEASES_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseRecycleLeases("USE_RECYCLE_LEASES_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "USE_RECYCLE_LEASES_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseSubscribeSettings(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_subscribe_settings"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseSubscribeSettings("USE_SUBSCRIBE_SETTINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_subscribe_settings", "USE_SUBSCRIBE_SETTINGS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseSubscribeSettings("USE_SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_subscribe_settings", "USE_SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseUpdateDnsOnLeaseRenewal(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_update_dns_on_lease_renewal"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseUpdateDnsOnLeaseRenewal("USE_UPDATE_DNS_ON_LEASE_RENEWAL_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "USE_UPDATE_DNS_ON_LEASE_RENEWAL_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseUpdateDnsOnLeaseRenewal("USE_UPDATE_DNS_ON_LEASE_RENEWAL_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "USE_UPDATE_DNS_ON_LEASE_RENEWAL_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_UseZoneAssociations(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_use_zone_associations"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerUseZoneAssociations("USE_ZONE_ASSOCIATIONS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_zone_associations", "USE_ZONE_ASSOCIATIONS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerUseZoneAssociations("USE_ZONE_ASSOCIATIONS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_zone_associations", "USE_ZONE_ASSOCIATIONS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

// func TestAccNetworkcontainerResource_ZoneAssociations(t *testing.T) {
// 	var resourceName = "nios_nios_ipam_networkcontainer.test_zone_associations"
// 	var v ipam.Networkcontainer

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkcontainerZoneAssociations("ZONE_ASSOCIATIONS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "zone_associations", "ZONE_ASSOCIATIONS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkcontainerZoneAssociations("ZONE_ASSOCIATIONS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "zone_associations", "ZONE_ASSOCIATIONS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

func testAccCheckNetworkcontainerExists(ctx context.Context, resourceName string, v *ipam.Networkcontainer) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.IPAMAPI.
			NetworkcontainerAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForNetworkcontainer).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetNetworkcontainerResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetNetworkcontainerResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckNetworkcontainerDestroy(ctx context.Context, v *ipam.Networkcontainer) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.IPAMAPI.
			NetworkcontainerAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForNetworkcontainer).
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

// func testAccCheckNetworkcontainerDisappears(ctx context.Context, v *ipam.Networkcontainer) resource.TestCheckFunc {
// 	// Delete the resource externally to verify disappears test
// 	return func(state *terraform.State) error {
// 		_, err := acctest.NIOSClient.IPAMAPI.
// 			NetworkcontainerAPI.
// 			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
// 			Execute()
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	}
// }

func testAccNetworkcontainerBasicConfig(network string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_nios_ipam_networkcontainer" "test" {
	network = %q
}
`, network)
}

// func testAccNetworkcontainerRef(ref string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test__ref" {
//     _ref = %q
// }
// `, ref)
// }

// func testAccNetworkcontainerAuthority(authority string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_authority" {
//     authority = %q
// }
// `, authority)
// }

// func testAccNetworkcontainerAutoCreateReversezone(autoCreateReversezone string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_auto_create_reversezone" {
//     auto_create_reversezone = %q
// }
// `, autoCreateReversezone)
// }

// func testAccNetworkcontainerBootfile(bootfile string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_bootfile" {
//     bootfile = %q
// }
// `, bootfile)
// }

// func testAccNetworkcontainerBootserver(bootserver string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_bootserver" {
//     bootserver = %q
// }
// `, bootserver)
// }

// func testAccNetworkcontainerCloudInfo(cloudInfo string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_cloud_info" {
//     cloud_info = %q
// }
// `, cloudInfo)
// }

// func testAccNetworkcontainerComment(comment string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_comment" {
//     comment = %q
// }
// `, comment)
// }

// func testAccNetworkcontainerDdnsDomainname(ddnsDomainname string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_ddns_domainname" {
//     ddns_domainname = %q
// }
// `, ddnsDomainname)
// }

// func testAccNetworkcontainerDdnsGenerateHostname(ddnsGenerateHostname string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_ddns_generate_hostname" {
//     ddns_generate_hostname = %q
// }
// `, ddnsGenerateHostname)
// }

// func testAccNetworkcontainerDdnsServerAlwaysUpdates(ddnsServerAlwaysUpdates string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_ddns_server_always_updates" {
//     ddns_server_always_updates = %q
// }
// `, ddnsServerAlwaysUpdates)
// }

// func testAccNetworkcontainerDdnsTtl(ddnsTtl string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_ddns_ttl" {
//     ddns_ttl = %q
// }
// `, ddnsTtl)
// }

// func testAccNetworkcontainerDdnsUpdateFixedAddresses(ddnsUpdateFixedAddresses string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_ddns_update_fixed_addresses" {
//     ddns_update_fixed_addresses = %q
// }
// `, ddnsUpdateFixedAddresses)
// }

// func testAccNetworkcontainerDdnsUseOption81(ddnsUseOption81 string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_ddns_use_option81" {
//     ddns_use_option81 = %q
// }
// `, ddnsUseOption81)
// }

// func testAccNetworkcontainerDeleteReason(deleteReason string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_delete_reason" {
//     delete_reason = %q
// }
// `, deleteReason)
// }

// func testAccNetworkcontainerDenyBootp(denyBootp string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_deny_bootp" {
//     deny_bootp = %q
// }
// `, denyBootp)
// }

// func testAccNetworkcontainerDiscoveryBasicPollSettings(discoveryBasicPollSettings string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_discovery_basic_poll_settings" {
//     discovery_basic_poll_settings = %q
// }
// `, discoveryBasicPollSettings)
// }

// func testAccNetworkcontainerDiscoveryBlackoutSetting(discoveryBlackoutSetting string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_discovery_blackout_setting" {
//     discovery_blackout_setting = %q
// }
// `, discoveryBlackoutSetting)
// }

// func testAccNetworkcontainerDiscoveryMember(discoveryMember string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_discovery_member" {
//     discovery_member = %q
// }
// `, discoveryMember)
// }

// func testAccNetworkcontainerEmailList(emailList string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_email_list" {
//     email_list = %q
// }
// `, emailList)
// }

// func testAccNetworkcontainerEnableDdns(enableDdns string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_enable_ddns" {
//     enable_ddns = %q
// }
// `, enableDdns)
// }

// func testAccNetworkcontainerEnableDhcpThresholds(enableDhcpThresholds string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_enable_dhcp_thresholds" {
//     enable_dhcp_thresholds = %q
// }
// `, enableDhcpThresholds)
// }

// func testAccNetworkcontainerEnableDiscovery(enableDiscovery string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_enable_discovery" {
//     enable_discovery = %q
// }
// `, enableDiscovery)
// }

// func testAccNetworkcontainerEnableEmailWarnings(enableEmailWarnings string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_enable_email_warnings" {
//     enable_email_warnings = %q
// }
// `, enableEmailWarnings)
// }

// func testAccNetworkcontainerEnableImmediateDiscovery(enableImmediateDiscovery string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_enable_immediate_discovery" {
//     enable_immediate_discovery = %q
// }
// `, enableImmediateDiscovery)
// }

// func testAccNetworkcontainerEnablePxeLeaseTime(enablePxeLeaseTime string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_enable_pxe_lease_time" {
//     enable_pxe_lease_time = %q
// }
// `, enablePxeLeaseTime)
// }

// func testAccNetworkcontainerEnableSnmpWarnings(enableSnmpWarnings string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_enable_snmp_warnings" {
//     enable_snmp_warnings = %q
// }
// `, enableSnmpWarnings)
// }

// func testAccNetworkcontainerExtAttrs(extAttrs string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_extattrs" {
//     extattrs = %q
// }
// `, extAttrs)
// }

// func testAccNetworkcontainerFederatedRealms(federatedRealms string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_federated_realms" {
//     federated_realms = %q
// }
// `, federatedRealms)
// }

// func testAccNetworkcontainerHighWaterMark(highWaterMark string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_high_water_mark" {
//     high_water_mark = %q
// }
// `, highWaterMark)
// }

// func testAccNetworkcontainerHighWaterMarkReset(highWaterMarkReset string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_high_water_mark_reset" {
//     high_water_mark_reset = %q
// }
// `, highWaterMarkReset)
// }

// func testAccNetworkcontainerIgnoreDhcpOptionListRequest(ignoreDhcpOptionListRequest string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_ignore_dhcp_option_list_request" {
//     ignore_dhcp_option_list_request = %q
// }
// `, ignoreDhcpOptionListRequest)
// }

// func testAccNetworkcontainerIgnoreId(ignoreId string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_ignore_id" {
//     ignore_id = %q
// }
// `, ignoreId)
// }

// func testAccNetworkcontainerIgnoreMacAddresses(ignoreMacAddresses string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_ignore_mac_addresses" {
//     ignore_mac_addresses = %q
// }
// `, ignoreMacAddresses)
// }

// func testAccNetworkcontainerIpamEmailAddresses(ipamEmailAddresses string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_ipam_email_addresses" {
//     ipam_email_addresses = %q
// }
// `, ipamEmailAddresses)
// }

// func testAccNetworkcontainerIpamThresholdSettings(ipamThresholdSettings string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_ipam_threshold_settings" {
//     ipam_threshold_settings = %q
// }
// `, ipamThresholdSettings)
// }

// func testAccNetworkcontainerIpamTrapSettings(ipamTrapSettings string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_ipam_trap_settings" {
//     ipam_trap_settings = %q
// }
// `, ipamTrapSettings)
// }

// func testAccNetworkcontainerLeaseScavengeTime(leaseScavengeTime string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_lease_scavenge_time" {
//     lease_scavenge_time = %q
// }
// `, leaseScavengeTime)
// }

// func testAccNetworkcontainerLogicFilterRules(logicFilterRules string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_logic_filter_rules" {
//     logic_filter_rules = %q
// }
// `, logicFilterRules)
// }

// func testAccNetworkcontainerLowWaterMark(lowWaterMark string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_low_water_mark" {
//     low_water_mark = %q
// }
// `, lowWaterMark)
// }

// func testAccNetworkcontainerLowWaterMarkReset(lowWaterMarkReset string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_low_water_mark_reset" {
//     low_water_mark_reset = %q
// }
// `, lowWaterMarkReset)
// }

// func testAccNetworkcontainerMgmPrivate(mgmPrivate string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_mgm_private" {
//     mgm_private = %q
// }
// `, mgmPrivate)
// }

// func testAccNetworkcontainerMsAdUserData(msAdUserData string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_ms_ad_user_data" {
//     ms_ad_user_data = %q
// }
// `, msAdUserData)
// }

// func testAccNetworkcontainerNetwork(network string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_network" {
//     network = %q
// }
// `, network)
// }

// func testAccNetworkcontainerFuncCall(funcCall string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_func_call" {
//     func_call = %q
// }
// `, funcCall)
// }

// func testAccNetworkcontainerNetworkView(networkView string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_network_view" {
//     network_view = %q
// }
// `, networkView)
// }

// func testAccNetworkcontainerNextserver(nextserver string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_nextserver" {
//     nextserver = %q
// }
// `, nextserver)
// }

// func testAccNetworkcontainerOptions(options string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_options" {
//     options = %q
// }
// `, options)
// }

// func testAccNetworkcontainerPortControlBlackoutSetting(portControlBlackoutSetting string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_port_control_blackout_setting" {
//     port_control_blackout_setting = %q
// }
// `, portControlBlackoutSetting)
// }

// func testAccNetworkcontainerPxeLeaseTime(pxeLeaseTime string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_pxe_lease_time" {
//     pxe_lease_time = %q
// }
// `, pxeLeaseTime)
// }

// func testAccNetworkcontainerRecycleLeases(recycleLeases string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_recycle_leases" {
//     recycle_leases = %q
// }
// `, recycleLeases)
// }

// func testAccNetworkcontainerRemoveSubnets(removeSubnets string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_remove_subnets" {
//     remove_subnets = %q
// }
// `, removeSubnets)
// }

// func testAccNetworkcontainerRestartIfNeeded(restartIfNeeded string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_restart_if_needed" {
//     restart_if_needed = %q
// }
// `, restartIfNeeded)
// }

// func testAccNetworkcontainerRirOrganization(rirOrganization string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_rir_organization" {
//     rir_organization = %q
// }
// `, rirOrganization)
// }

// func testAccNetworkcontainerRirRegistrationAction(rirRegistrationAction string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_rir_registration_action" {
//     rir_registration_action = %q
// }
// `, rirRegistrationAction)
// }

// func testAccNetworkcontainerRirRegistrationStatus(rirRegistrationStatus string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_rir_registration_status" {
//     rir_registration_status = %q
// }
// `, rirRegistrationStatus)
// }

// func testAccNetworkcontainerSamePortControlDiscoveryBlackout(samePortControlDiscoveryBlackout string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_same_port_control_discovery_blackout" {
//     same_port_control_discovery_blackout = %q
// }
// `, samePortControlDiscoveryBlackout)
// }

// func testAccNetworkcontainerSendRirRequest(sendRirRequest string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_send_rir_request" {
//     send_rir_request = %q
// }
// `, sendRirRequest)
// }

// func testAccNetworkcontainerSubscribeSettings(subscribeSettings string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_subscribe_settings" {
//     subscribe_settings = %q
// }
// `, subscribeSettings)
// }

// func testAccNetworkcontainerUnmanaged(unmanaged string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_unmanaged" {
//     unmanaged = %q
// }
// `, unmanaged)
// }

// func testAccNetworkcontainerUpdateDnsOnLeaseRenewal(updateDnsOnLeaseRenewal string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_update_dns_on_lease_renewal" {
//     update_dns_on_lease_renewal = %q
// }
// `, updateDnsOnLeaseRenewal)
// }

// func testAccNetworkcontainerUseAuthority(useAuthority string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_authority" {
//     use_authority = %q
// }
// `, useAuthority)
// }

// func testAccNetworkcontainerUseBlackoutSetting(useBlackoutSetting string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_blackout_setting" {
//     use_blackout_setting = %q
// }
// `, useBlackoutSetting)
// }

// func testAccNetworkcontainerUseBootfile(useBootfile string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_bootfile" {
//     use_bootfile = %q
// }
// `, useBootfile)
// }

// func testAccNetworkcontainerUseBootserver(useBootserver string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_bootserver" {
//     use_bootserver = %q
// }
// `, useBootserver)
// }

// func testAccNetworkcontainerUseDdnsDomainname(useDdnsDomainname string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_ddns_domainname" {
//     use_ddns_domainname = %q
// }
// `, useDdnsDomainname)
// }

// func testAccNetworkcontainerUseDdnsGenerateHostname(useDdnsGenerateHostname string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_ddns_generate_hostname" {
//     use_ddns_generate_hostname = %q
// }
// `, useDdnsGenerateHostname)
// }

// func testAccNetworkcontainerUseDdnsTtl(useDdnsTtl string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_ddns_ttl" {
//     use_ddns_ttl = %q
// }
// `, useDdnsTtl)
// }

// func testAccNetworkcontainerUseDdnsUpdateFixedAddresses(useDdnsUpdateFixedAddresses string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_ddns_update_fixed_addresses" {
//     use_ddns_update_fixed_addresses = %q
// }
// `, useDdnsUpdateFixedAddresses)
// }

// func testAccNetworkcontainerUseDdnsUseOption81(useDdnsUseOption81 string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_ddns_use_option81" {
//     use_ddns_use_option81 = %q
// }
// `, useDdnsUseOption81)
// }

// func testAccNetworkcontainerUseDenyBootp(useDenyBootp string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_deny_bootp" {
//     use_deny_bootp = %q
// }
// `, useDenyBootp)
// }

// func testAccNetworkcontainerUseDiscoveryBasicPollingSettings(useDiscoveryBasicPollingSettings string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_discovery_basic_polling_settings" {
//     use_discovery_basic_polling_settings = %q
// }
// `, useDiscoveryBasicPollingSettings)
// }

// func testAccNetworkcontainerUseEmailList(useEmailList string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_email_list" {
//     use_email_list = %q
// }
// `, useEmailList)
// }

// func testAccNetworkcontainerUseEnableDdns(useEnableDdns string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_enable_ddns" {
//     use_enable_ddns = %q
// }
// `, useEnableDdns)
// }

// func testAccNetworkcontainerUseEnableDhcpThresholds(useEnableDhcpThresholds string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_enable_dhcp_thresholds" {
//     use_enable_dhcp_thresholds = %q
// }
// `, useEnableDhcpThresholds)
// }

// func testAccNetworkcontainerUseEnableDiscovery(useEnableDiscovery string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_enable_discovery" {
//     use_enable_discovery = %q
// }
// `, useEnableDiscovery)
// }

// func testAccNetworkcontainerUseIgnoreDhcpOptionListRequest(useIgnoreDhcpOptionListRequest string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_ignore_dhcp_option_list_request" {
//     use_ignore_dhcp_option_list_request = %q
// }
// `, useIgnoreDhcpOptionListRequest)
// }

// func testAccNetworkcontainerUseIgnoreId(useIgnoreId string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_ignore_id" {
//     use_ignore_id = %q
// }
// `, useIgnoreId)
// }

// func testAccNetworkcontainerUseIpamEmailAddresses(useIpamEmailAddresses string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_ipam_email_addresses" {
//     use_ipam_email_addresses = %q
// }
// `, useIpamEmailAddresses)
// }

// func testAccNetworkcontainerUseIpamThresholdSettings(useIpamThresholdSettings string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_ipam_threshold_settings" {
//     use_ipam_threshold_settings = %q
// }
// `, useIpamThresholdSettings)
// }

// func testAccNetworkcontainerUseIpamTrapSettings(useIpamTrapSettings string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_ipam_trap_settings" {
//     use_ipam_trap_settings = %q
// }
// `, useIpamTrapSettings)
// }

// func testAccNetworkcontainerUseLeaseScavengeTime(useLeaseScavengeTime string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_lease_scavenge_time" {
//     use_lease_scavenge_time = %q
// }
// `, useLeaseScavengeTime)
// }

// func testAccNetworkcontainerUseLogicFilterRules(useLogicFilterRules string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_logic_filter_rules" {
//     use_logic_filter_rules = %q
// }
// `, useLogicFilterRules)
// }

// func testAccNetworkcontainerUseMgmPrivate(useMgmPrivate string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_mgm_private" {
//     use_mgm_private = %q
// }
// `, useMgmPrivate)
// }

// func testAccNetworkcontainerUseNextserver(useNextserver string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_nextserver" {
//     use_nextserver = %q
// }
// `, useNextserver)
// }

// func testAccNetworkcontainerUseOptions(useOptions string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_options" {
//     use_options = %q
// }
// `, useOptions)
// }

// func testAccNetworkcontainerUsePxeLeaseTime(usePxeLeaseTime string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_pxe_lease_time" {
//     use_pxe_lease_time = %q
// }
// `, usePxeLeaseTime)
// }

// func testAccNetworkcontainerUseRecycleLeases(useRecycleLeases string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_recycle_leases" {
//     use_recycle_leases = %q
// }
// `, useRecycleLeases)
// }

// func testAccNetworkcontainerUseSubscribeSettings(useSubscribeSettings string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_subscribe_settings" {
//     use_subscribe_settings = %q
// }
// `, useSubscribeSettings)
// }

// func testAccNetworkcontainerUseUpdateDnsOnLeaseRenewal(useUpdateDnsOnLeaseRenewal string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_update_dns_on_lease_renewal" {
//     use_update_dns_on_lease_renewal = %q
// }
// `, useUpdateDnsOnLeaseRenewal)
// }

// func testAccNetworkcontainerUseZoneAssociations(useZoneAssociations string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_use_zone_associations" {
//     use_zone_associations = %q
// }
// `, useZoneAssociations)
// }

// func testAccNetworkcontainerZoneAssociations(zoneAssociations string) string {
// 	return fmt.Sprintf(`
// resource "nios_nios_ipam_networkcontainer" "test_zone_associations" {
//     zone_associations = %q
// }
// `, zoneAssociations)
// }
