package ipam_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/ipam"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/acctest"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/utils"
)

// TODO : Add readable attributes for the resource
var readableAttributesForNetwork = "authority,bootfile,bootserver,cloud_info,cloud_shared,comment,conflict_count,ddns_domainname,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,ddns_update_fixed_addresses,ddns_use_option81,deny_bootp,dhcp_utilization,dhcp_utilization_status,disable,discover_now_status,discovered_bgp_as,discovered_bridge_domain,discovered_tenant,discovered_vlan_id,discovered_vlan_name,discovered_vrf_description,discovered_vrf_name,discovered_vrf_rd,discovery_basic_poll_settings,discovery_blackout_setting,discovery_engine_type,discovery_member,dynamic_hosts,email_list,enable_ddns,enable_dhcp_thresholds,enable_discovery,enable_email_warnings,enable_ifmap_publishing,enable_snmp_warnings,endpoint_sources,extattrs,federated_realms,high_water_mark,high_water_mark_reset,ignore_dhcp_option_list_request,ignore_id,ignore_mac_addresses,ipam_email_addresses,ipam_threshold_settings,ipam_trap_settings,ipv4addr,last_rir_registration_update_sent,last_rir_registration_update_status,lease_scavenge_time,logic_filter_rules,low_water_mark,low_water_mark_reset,members,mgm_private,mgm_private_overridable,ms_ad_user_data,netmask,network,network_container,network_view,nextserver,options,port_control_blackout_setting,pxe_lease_time,recycle_leases,rir,rir_organization,rir_registration_status,same_port_control_discovery_blackout,static_hosts,subscribe_settings,total_hosts,unmanaged,unmanaged_count,update_dns_on_lease_renewal,use_authority,use_blackout_setting,use_bootfile,use_bootserver,use_ddns_domainname,use_ddns_generate_hostname,use_ddns_ttl,use_ddns_update_fixed_addresses,use_ddns_use_option81,use_deny_bootp,use_discovery_basic_polling_settings,use_email_list,use_enable_ddns,use_enable_dhcp_thresholds,use_enable_discovery,use_enable_ifmap_publishing,use_ignore_dhcp_option_list_request,use_ignore_id,use_ipam_email_addresses,use_ipam_threshold_settings,use_ipam_trap_settings,use_lease_scavenge_time,use_logic_filter_rules,use_mgm_private,use_nextserver,use_options,use_pxe_lease_time,use_recycle_leases,use_subscribe_settings,use_update_dns_on_lease_renewal,use_zone_associations,utilization,utilization_update,vlans,zone_associations"

func TestAccNetworkResource_basic(t *testing.T) {
	var resourceName = "nios_ipam_network.test"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkBasicConfig(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_disappears(t *testing.T) {
	resourceName := "nios_ipam_network.test"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetworkDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkBasicConfig(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					testAccCheckNetworkDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNetworkResource_Authority(t *testing.T) {
	var resourceName = "nios_ipam_network.test_authority"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkAuthority(network, "false", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "authority", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkAuthority(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "authority", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

//     func TestAccNetworkResource_AutoCreateReversezone(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_auto_create_reversezone"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkAutoCreateReversezone("AUTO_CREATE_REVERSEZONE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "auto_create_reversezone", "AUTO_CREATE_REVERSEZONE_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkAutoCreateReversezone("AUTO_CREATE_REVERSEZONE_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "auto_create_reversezone", "AUTO_CREATE_REVERSEZONE_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_Bootfile(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_bootfile"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkBootfile("BOOTFILE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "bootfile", "BOOTFILE_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkBootfile("BOOTFILE_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "bootfile", "BOOTFILE_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_Bootserver(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_bootserver"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkBootserver("BOOTSERVER_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "bootserver", "BOOTSERVER_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkBootserver("BOOTSERVER_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "bootserver", "BOOTSERVER_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_CloudInfo(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_cloud_info"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkCloudInfo("CLOUD_INFO_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "cloud_info", "CLOUD_INFO_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkCloudInfo("CLOUD_INFO_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "cloud_info", "CLOUD_INFO_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_CloudShared(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_cloud_shared"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkCloudShared("CLOUD_SHARED_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "cloud_shared", "CLOUD_SHARED_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkCloudShared("CLOUD_SHARED_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "cloud_shared", "CLOUD_SHARED_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_Comment(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_comment"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkComment("COMMENT_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkComment("COMMENT_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_DdnsDomainname(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_ddns_domainname"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkDdnsDomainname("DDNS_DOMAINNAME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "DDNS_DOMAINNAME_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkDdnsDomainname("DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_DdnsGenerateHostname(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_ddns_generate_hostname"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkDdnsGenerateHostname("DDNS_GENERATE_HOSTNAME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "DDNS_GENERATE_HOSTNAME_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkDdnsGenerateHostname("DDNS_GENERATE_HOSTNAME_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "DDNS_GENERATE_HOSTNAME_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_DdnsServerAlwaysUpdates(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_ddns_server_always_updates"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkDdnsServerAlwaysUpdates("DDNS_SERVER_ALWAYS_UPDATES_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "DDNS_SERVER_ALWAYS_UPDATES_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkDdnsServerAlwaysUpdates("DDNS_SERVER_ALWAYS_UPDATES_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "DDNS_SERVER_ALWAYS_UPDATES_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_DdnsTtl(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_ddns_ttl"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkDdnsTtl("DDNS_TTL_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "DDNS_TTL_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkDdnsTtl("DDNS_TTL_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "DDNS_TTL_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_DdnsUpdateFixedAddresses(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_ddns_update_fixed_addresses"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkDdnsUpdateFixedAddresses("DDNS_UPDATE_FIXED_ADDRESSES_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_update_fixed_addresses", "DDNS_UPDATE_FIXED_ADDRESSES_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkDdnsUpdateFixedAddresses("DDNS_UPDATE_FIXED_ADDRESSES_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_update_fixed_addresses", "DDNS_UPDATE_FIXED_ADDRESSES_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_DdnsUseOption81(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_ddns_use_option81"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkDdnsUseOption81("DDNS_USE_OPTION81_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "DDNS_USE_OPTION81_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkDdnsUseOption81("DDNS_USE_OPTION81_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "DDNS_USE_OPTION81_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_DeleteReason(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_delete_reason"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkDeleteReason("DELETE_REASON_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "delete_reason", "DELETE_REASON_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkDeleteReason("DELETE_REASON_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "delete_reason", "DELETE_REASON_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_DenyBootp(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_deny_bootp"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkDenyBootp("DENY_BOOTP_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "DENY_BOOTP_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkDenyBootp("DENY_BOOTP_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "DENY_BOOTP_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_Disable(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_disable"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkDisable("DISABLE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkDisable("DISABLE_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_DiscoveredBridgeDomain(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_discovered_bridge_domain"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkDiscoveredBridgeDomain("DISCOVERED_BRIDGE_DOMAIN_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "discovered_bridge_domain", "DISCOVERED_BRIDGE_DOMAIN_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkDiscoveredBridgeDomain("DISCOVERED_BRIDGE_DOMAIN_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "discovered_bridge_domain", "DISCOVERED_BRIDGE_DOMAIN_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_DiscoveredTenant(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_discovered_tenant"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkDiscoveredTenant("DISCOVERED_TENANT_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "discovered_tenant", "DISCOVERED_TENANT_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkDiscoveredTenant("DISCOVERED_TENANT_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "discovered_tenant", "DISCOVERED_TENANT_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_DiscoveryBasicPollSettings(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_discovery_basic_poll_settings"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkDiscoveryBasicPollSettings("DISCOVERY_BASIC_POLL_SETTINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings", "DISCOVERY_BASIC_POLL_SETTINGS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkDiscoveryBasicPollSettings("DISCOVERY_BASIC_POLL_SETTINGS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings", "DISCOVERY_BASIC_POLL_SETTINGS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_DiscoveryBlackoutSetting(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_discovery_blackout_setting"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkDiscoveryBlackoutSetting("DISCOVERY_BLACKOUT_SETTING_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "discovery_blackout_setting", "DISCOVERY_BLACKOUT_SETTING_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkDiscoveryBlackoutSetting("DISCOVERY_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "discovery_blackout_setting", "DISCOVERY_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_DiscoveryMember(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_discovery_member"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkDiscoveryMember("DISCOVERY_MEMBER_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "discovery_member", "DISCOVERY_MEMBER_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkDiscoveryMember("DISCOVERY_MEMBER_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "discovery_member", "DISCOVERY_MEMBER_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_EmailList(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_email_list"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkEmailList("EMAIL_LIST_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "email_list", "EMAIL_LIST_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkEmailList("EMAIL_LIST_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "email_list", "EMAIL_LIST_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_EnableDdns(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_enable_ddns"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkEnableDdns("ENABLE_DDNS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "ENABLE_DDNS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkEnableDdns("ENABLE_DDNS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "ENABLE_DDNS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_EnableDhcpThresholds(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_enable_dhcp_thresholds"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkEnableDhcpThresholds("ENABLE_DHCP_THRESHOLDS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_dhcp_thresholds", "ENABLE_DHCP_THRESHOLDS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkEnableDhcpThresholds("ENABLE_DHCP_THRESHOLDS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_dhcp_thresholds", "ENABLE_DHCP_THRESHOLDS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_EnableDiscovery(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_enable_discovery"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkEnableDiscovery("ENABLE_DISCOVERY_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_discovery", "ENABLE_DISCOVERY_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkEnableDiscovery("ENABLE_DISCOVERY_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_discovery", "ENABLE_DISCOVERY_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_EnableEmailWarnings(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_enable_email_warnings"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkEnableEmailWarnings("ENABLE_EMAIL_WARNINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_email_warnings", "ENABLE_EMAIL_WARNINGS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkEnableEmailWarnings("ENABLE_EMAIL_WARNINGS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_email_warnings", "ENABLE_EMAIL_WARNINGS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_EnableIfmapPublishing(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_enable_ifmap_publishing"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkEnableIfmapPublishing("ENABLE_IFMAP_PUBLISHING_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_ifmap_publishing", "ENABLE_IFMAP_PUBLISHING_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkEnableIfmapPublishing("ENABLE_IFMAP_PUBLISHING_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_ifmap_publishing", "ENABLE_IFMAP_PUBLISHING_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_EnableImmediateDiscovery(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_enable_immediate_discovery"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkEnableImmediateDiscovery("ENABLE_IMMEDIATE_DISCOVERY_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_immediate_discovery", "ENABLE_IMMEDIATE_DISCOVERY_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkEnableImmediateDiscovery("ENABLE_IMMEDIATE_DISCOVERY_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_immediate_discovery", "ENABLE_IMMEDIATE_DISCOVERY_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_EnablePxeLeaseTime(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_enable_pxe_lease_time"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkEnablePxeLeaseTime("ENABLE_PXE_LEASE_TIME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "ENABLE_PXE_LEASE_TIME_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkEnablePxeLeaseTime("ENABLE_PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "ENABLE_PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_EnableSnmpWarnings(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_enable_snmp_warnings"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkEnableSnmpWarnings("ENABLE_SNMP_WARNINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_snmp_warnings", "ENABLE_SNMP_WARNINGS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkEnableSnmpWarnings("ENABLE_SNMP_WARNINGS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "enable_snmp_warnings", "ENABLE_SNMP_WARNINGS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_ExtAttrs(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_extattrs"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkExtAttrs("EXT_ATTRS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_FederatedRealms(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_federated_realms"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkFederatedRealms("FEDERATED_REALMS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "federated_realms", "FEDERATED_REALMS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkFederatedRealms("FEDERATED_REALMS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "federated_realms", "FEDERATED_REALMS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_HighWaterMark(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_high_water_mark"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkHighWaterMark("HIGH_WATER_MARK_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "HIGH_WATER_MARK_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkHighWaterMark("HIGH_WATER_MARK_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "HIGH_WATER_MARK_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_HighWaterMarkReset(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_high_water_mark_reset"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkHighWaterMarkReset("HIGH_WATER_MARK_RESET_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "HIGH_WATER_MARK_RESET_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkHighWaterMarkReset("HIGH_WATER_MARK_RESET_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "HIGH_WATER_MARK_RESET_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_IgnoreDhcpOptionListRequest(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_ignore_dhcp_option_list_request"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkIgnoreDhcpOptionListRequest("IGNORE_DHCP_OPTION_LIST_REQUEST_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "IGNORE_DHCP_OPTION_LIST_REQUEST_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkIgnoreDhcpOptionListRequest("IGNORE_DHCP_OPTION_LIST_REQUEST_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "IGNORE_DHCP_OPTION_LIST_REQUEST_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_IgnoreId(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_ignore_id"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkIgnoreId("IGNORE_ID_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ignore_id", "IGNORE_ID_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkIgnoreId("IGNORE_ID_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ignore_id", "IGNORE_ID_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_IgnoreMacAddresses(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_ignore_mac_addresses"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkIgnoreMacAddresses("IGNORE_MAC_ADDRESSES_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ignore_mac_addresses", "IGNORE_MAC_ADDRESSES_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkIgnoreMacAddresses("IGNORE_MAC_ADDRESSES_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ignore_mac_addresses", "IGNORE_MAC_ADDRESSES_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_IpamEmailAddresses(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_ipam_email_addresses"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkIpamEmailAddresses("IPAM_EMAIL_ADDRESSES_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ipam_email_addresses", "IPAM_EMAIL_ADDRESSES_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkIpamEmailAddresses("IPAM_EMAIL_ADDRESSES_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ipam_email_addresses", "IPAM_EMAIL_ADDRESSES_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_IpamThresholdSettings(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_ipam_threshold_settings"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkIpamThresholdSettings("IPAM_THRESHOLD_SETTINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ipam_threshold_settings", "IPAM_THRESHOLD_SETTINGS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkIpamThresholdSettings("IPAM_THRESHOLD_SETTINGS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ipam_threshold_settings", "IPAM_THRESHOLD_SETTINGS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_IpamTrapSettings(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_ipam_trap_settings"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkIpamTrapSettings("IPAM_TRAP_SETTINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ipam_trap_settings", "IPAM_TRAP_SETTINGS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkIpamTrapSettings("IPAM_TRAP_SETTINGS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ipam_trap_settings", "IPAM_TRAP_SETTINGS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_Ipv4addr(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_ipv4addr"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkIpv4addr("IPV4ADDR_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "IPV4ADDR_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkIpv4addr("IPV4ADDR_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "IPV4ADDR_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_LeaseScavengeTime(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_lease_scavenge_time"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkLeaseScavengeTime("LEASE_SCAVENGE_TIME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "LEASE_SCAVENGE_TIME_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkLeaseScavengeTime("LEASE_SCAVENGE_TIME_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "LEASE_SCAVENGE_TIME_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_LogicFilterRules(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_logic_filter_rules"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkLogicFilterRules("LOGIC_FILTER_RULES_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkLogicFilterRules("LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_LowWaterMark(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_low_water_mark"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkLowWaterMark("LOW_WATER_MARK_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "LOW_WATER_MARK_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkLowWaterMark("LOW_WATER_MARK_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "LOW_WATER_MARK_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_LowWaterMarkReset(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_low_water_mark_reset"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkLowWaterMarkReset("LOW_WATER_MARK_RESET_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "LOW_WATER_MARK_RESET_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkLowWaterMarkReset("LOW_WATER_MARK_RESET_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "LOW_WATER_MARK_RESET_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_Members(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_members"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkMembers("MEMBERS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "members", "MEMBERS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkMembers("MEMBERS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "members", "MEMBERS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_MgmPrivate(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_mgm_private"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkMgmPrivate("MGM_PRIVATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "mgm_private", "MGM_PRIVATE_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkMgmPrivate("MGM_PRIVATE_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "mgm_private", "MGM_PRIVATE_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_MsAdUserData(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_ms_ad_user_data"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkMsAdUserData("MS_AD_USER_DATA_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ms_ad_user_data", "MS_AD_USER_DATA_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkMsAdUserData("MS_AD_USER_DATA_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ms_ad_user_data", "MS_AD_USER_DATA_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_Netmask(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_netmask"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkNetmask("NETMASK_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "netmask", "NETMASK_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkNetmask("NETMASK_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "netmask", "NETMASK_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_Network(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_network"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkNetwork("NETWORK_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "network", "NETWORK_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkNetwork("NETWORK_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "network", "NETWORK_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_FuncCall(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_func_call"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkFuncCall("FUNC_CALL_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "func_call", "FUNC_CALL_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkFuncCall("FUNC_CALL_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "func_call", "FUNC_CALL_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_NetworkView(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_network_view"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkNetworkView("NETWORK_VIEW_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkNetworkView("NETWORK_VIEW_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_Nextserver(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_nextserver"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkNextserver("NEXTSERVER_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "nextserver", "NEXTSERVER_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkNextserver("NEXTSERVER_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "nextserver", "NEXTSERVER_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_Options(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_options"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkOptions("OPTIONS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkOptions("OPTIONS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_PortControlBlackoutSetting(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_port_control_blackout_setting"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkPortControlBlackoutSetting("PORT_CONTROL_BLACKOUT_SETTING_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "port_control_blackout_setting", "PORT_CONTROL_BLACKOUT_SETTING_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkPortControlBlackoutSetting("PORT_CONTROL_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "port_control_blackout_setting", "PORT_CONTROL_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_PxeLeaseTime(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_pxe_lease_time"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkPxeLeaseTime("PXE_LEASE_TIME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "PXE_LEASE_TIME_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkPxeLeaseTime("PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_RecycleLeases(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_recycle_leases"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkRecycleLeases("RECYCLE_LEASES_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "RECYCLE_LEASES_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkRecycleLeases("RECYCLE_LEASES_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "RECYCLE_LEASES_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_RestartIfNeeded(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_restart_if_needed"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkRestartIfNeeded("RESTART_IF_NEEDED_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "restart_if_needed", "RESTART_IF_NEEDED_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkRestartIfNeeded("RESTART_IF_NEEDED_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "restart_if_needed", "RESTART_IF_NEEDED_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_RirOrganization(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_rir_organization"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkRirOrganization("RIR_ORGANIZATION_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "rir_organization", "RIR_ORGANIZATION_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkRirOrganization("RIR_ORGANIZATION_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "rir_organization", "RIR_ORGANIZATION_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_RirRegistrationAction(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_rir_registration_action"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkRirRegistrationAction("RIR_REGISTRATION_ACTION_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "rir_registration_action", "RIR_REGISTRATION_ACTION_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkRirRegistrationAction("RIR_REGISTRATION_ACTION_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "rir_registration_action", "RIR_REGISTRATION_ACTION_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_RirRegistrationStatus(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_rir_registration_status"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkRirRegistrationStatus("RIR_REGISTRATION_STATUS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "rir_registration_status", "RIR_REGISTRATION_STATUS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkRirRegistrationStatus("RIR_REGISTRATION_STATUS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "rir_registration_status", "RIR_REGISTRATION_STATUS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_SamePortControlDiscoveryBlackout(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_same_port_control_discovery_blackout"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkSamePortControlDiscoveryBlackout("SAME_PORT_CONTROL_DISCOVERY_BLACKOUT_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "SAME_PORT_CONTROL_DISCOVERY_BLACKOUT_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkSamePortControlDiscoveryBlackout("SAME_PORT_CONTROL_DISCOVERY_BLACKOUT_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "SAME_PORT_CONTROL_DISCOVERY_BLACKOUT_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_SendRirRequest(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_send_rir_request"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkSendRirRequest("SEND_RIR_REQUEST_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "send_rir_request", "SEND_RIR_REQUEST_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkSendRirRequest("SEND_RIR_REQUEST_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "send_rir_request", "SEND_RIR_REQUEST_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_SubscribeSettings(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_subscribe_settings"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkSubscribeSettings("SUBSCRIBE_SETTINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "subscribe_settings", "SUBSCRIBE_SETTINGS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkSubscribeSettings("SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "subscribe_settings", "SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_Template(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_template"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkTemplate("TEMPLATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "template", "TEMPLATE_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkTemplate("TEMPLATE_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "template", "TEMPLATE_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_Unmanaged(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_unmanaged"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUnmanaged("UNMANAGED_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "unmanaged", "UNMANAGED_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUnmanaged("UNMANAGED_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "unmanaged", "UNMANAGED_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UpdateDnsOnLeaseRenewal(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_update_dns_on_lease_renewal"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUpdateDnsOnLeaseRenewal("UPDATE_DNS_ON_LEASE_RENEWAL_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "UPDATE_DNS_ON_LEASE_RENEWAL_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUpdateDnsOnLeaseRenewal("UPDATE_DNS_ON_LEASE_RENEWAL_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "UPDATE_DNS_ON_LEASE_RENEWAL_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseAuthority(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_authority"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseAuthority("USE_AUTHORITY_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_authority", "USE_AUTHORITY_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseAuthority("USE_AUTHORITY_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_authority", "USE_AUTHORITY_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseBlackoutSetting(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_blackout_setting"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseBlackoutSetting("USE_BLACKOUT_SETTING_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "USE_BLACKOUT_SETTING_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseBlackoutSetting("USE_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "USE_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseBootfile(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_bootfile"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseBootfile("USE_BOOTFILE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "USE_BOOTFILE_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseBootfile("USE_BOOTFILE_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "USE_BOOTFILE_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseBootserver(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_bootserver"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseBootserver("USE_BOOTSERVER_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "USE_BOOTSERVER_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseBootserver("USE_BOOTSERVER_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "USE_BOOTSERVER_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseDdnsDomainname(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_ddns_domainname"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseDdnsDomainname("USE_DDNS_DOMAINNAME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "USE_DDNS_DOMAINNAME_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseDdnsDomainname("USE_DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "USE_DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseDdnsGenerateHostname(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_ddns_generate_hostname"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseDdnsGenerateHostname("USE_DDNS_GENERATE_HOSTNAME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "USE_DDNS_GENERATE_HOSTNAME_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseDdnsGenerateHostname("USE_DDNS_GENERATE_HOSTNAME_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "USE_DDNS_GENERATE_HOSTNAME_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseDdnsTtl(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_ddns_ttl"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseDdnsTtl("USE_DDNS_TTL_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "USE_DDNS_TTL_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseDdnsTtl("USE_DDNS_TTL_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "USE_DDNS_TTL_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseDdnsUpdateFixedAddresses(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_ddns_update_fixed_addresses"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseDdnsUpdateFixedAddresses("USE_DDNS_UPDATE_FIXED_ADDRESSES_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ddns_update_fixed_addresses", "USE_DDNS_UPDATE_FIXED_ADDRESSES_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseDdnsUpdateFixedAddresses("USE_DDNS_UPDATE_FIXED_ADDRESSES_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ddns_update_fixed_addresses", "USE_DDNS_UPDATE_FIXED_ADDRESSES_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseDdnsUseOption81(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_ddns_use_option81"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseDdnsUseOption81("USE_DDNS_USE_OPTION81_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "USE_DDNS_USE_OPTION81_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseDdnsUseOption81("USE_DDNS_USE_OPTION81_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "USE_DDNS_USE_OPTION81_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseDenyBootp(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_deny_bootp"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseDenyBootp("USE_DENY_BOOTP_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "USE_DENY_BOOTP_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseDenyBootp("USE_DENY_BOOTP_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "USE_DENY_BOOTP_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseDiscoveryBasicPollingSettings(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_discovery_basic_polling_settings"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseDiscoveryBasicPollingSettings("USE_DISCOVERY_BASIC_POLLING_SETTINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "USE_DISCOVERY_BASIC_POLLING_SETTINGS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseDiscoveryBasicPollingSettings("USE_DISCOVERY_BASIC_POLLING_SETTINGS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "USE_DISCOVERY_BASIC_POLLING_SETTINGS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseEmailList(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_email_list"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseEmailList("USE_EMAIL_LIST_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_email_list", "USE_EMAIL_LIST_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseEmailList("USE_EMAIL_LIST_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_email_list", "USE_EMAIL_LIST_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseEnableDdns(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_enable_ddns"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseEnableDdns("USE_ENABLE_DDNS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "USE_ENABLE_DDNS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseEnableDdns("USE_ENABLE_DDNS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "USE_ENABLE_DDNS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseEnableDhcpThresholds(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_enable_dhcp_thresholds"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseEnableDhcpThresholds("USE_ENABLE_DHCP_THRESHOLDS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "USE_ENABLE_DHCP_THRESHOLDS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseEnableDhcpThresholds("USE_ENABLE_DHCP_THRESHOLDS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "USE_ENABLE_DHCP_THRESHOLDS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseEnableDiscovery(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_enable_discovery"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseEnableDiscovery("USE_ENABLE_DISCOVERY_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "USE_ENABLE_DISCOVERY_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseEnableDiscovery("USE_ENABLE_DISCOVERY_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "USE_ENABLE_DISCOVERY_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseEnableIfmapPublishing(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_enable_ifmap_publishing"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseEnableIfmapPublishing("USE_ENABLE_IFMAP_PUBLISHING_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_enable_ifmap_publishing", "USE_ENABLE_IFMAP_PUBLISHING_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseEnableIfmapPublishing("USE_ENABLE_IFMAP_PUBLISHING_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_enable_ifmap_publishing", "USE_ENABLE_IFMAP_PUBLISHING_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseIgnoreDhcpOptionListRequest(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_ignore_dhcp_option_list_request"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseIgnoreDhcpOptionListRequest("USE_IGNORE_DHCP_OPTION_LIST_REQUEST_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "USE_IGNORE_DHCP_OPTION_LIST_REQUEST_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseIgnoreDhcpOptionListRequest("USE_IGNORE_DHCP_OPTION_LIST_REQUEST_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "USE_IGNORE_DHCP_OPTION_LIST_REQUEST_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseIgnoreId(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_ignore_id"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseIgnoreId("USE_IGNORE_ID_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ignore_id", "USE_IGNORE_ID_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseIgnoreId("USE_IGNORE_ID_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ignore_id", "USE_IGNORE_ID_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseIpamEmailAddresses(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_ipam_email_addresses"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseIpamEmailAddresses("USE_IPAM_EMAIL_ADDRESSES_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ipam_email_addresses", "USE_IPAM_EMAIL_ADDRESSES_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseIpamEmailAddresses("USE_IPAM_EMAIL_ADDRESSES_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ipam_email_addresses", "USE_IPAM_EMAIL_ADDRESSES_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseIpamThresholdSettings(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_ipam_threshold_settings"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseIpamThresholdSettings("USE_IPAM_THRESHOLD_SETTINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ipam_threshold_settings", "USE_IPAM_THRESHOLD_SETTINGS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseIpamThresholdSettings("USE_IPAM_THRESHOLD_SETTINGS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ipam_threshold_settings", "USE_IPAM_THRESHOLD_SETTINGS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseIpamTrapSettings(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_ipam_trap_settings"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseIpamTrapSettings("USE_IPAM_TRAP_SETTINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ipam_trap_settings", "USE_IPAM_TRAP_SETTINGS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseIpamTrapSettings("USE_IPAM_TRAP_SETTINGS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_ipam_trap_settings", "USE_IPAM_TRAP_SETTINGS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseLeaseScavengeTime(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_lease_scavenge_time"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseLeaseScavengeTime("USE_LEASE_SCAVENGE_TIME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "USE_LEASE_SCAVENGE_TIME_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseLeaseScavengeTime("USE_LEASE_SCAVENGE_TIME_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "USE_LEASE_SCAVENGE_TIME_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseLogicFilterRules(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_logic_filter_rules"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseLogicFilterRules("USE_LOGIC_FILTER_RULES_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "USE_LOGIC_FILTER_RULES_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseLogicFilterRules("USE_LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "USE_LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseMgmPrivate(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_mgm_private"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseMgmPrivate("USE_MGM_PRIVATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_mgm_private", "USE_MGM_PRIVATE_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseMgmPrivate("USE_MGM_PRIVATE_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_mgm_private", "USE_MGM_PRIVATE_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseNextserver(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_nextserver"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseNextserver("USE_NEXTSERVER_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "USE_NEXTSERVER_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseNextserver("USE_NEXTSERVER_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "USE_NEXTSERVER_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseOptions(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_options"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseOptions("USE_OPTIONS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_options", "USE_OPTIONS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseOptions("USE_OPTIONS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_options", "USE_OPTIONS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UsePxeLeaseTime(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_pxe_lease_time"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUsePxeLeaseTime("USE_PXE_LEASE_TIME_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "USE_PXE_LEASE_TIME_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUsePxeLeaseTime("USE_PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "USE_PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseRecycleLeases(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_recycle_leases"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseRecycleLeases("USE_RECYCLE_LEASES_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "USE_RECYCLE_LEASES_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseRecycleLeases("USE_RECYCLE_LEASES_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "USE_RECYCLE_LEASES_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseSubscribeSettings(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_subscribe_settings"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseSubscribeSettings("USE_SUBSCRIBE_SETTINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_subscribe_settings", "USE_SUBSCRIBE_SETTINGS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseSubscribeSettings("USE_SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_subscribe_settings", "USE_SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseUpdateDnsOnLeaseRenewal(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_update_dns_on_lease_renewal"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseUpdateDnsOnLeaseRenewal("USE_UPDATE_DNS_ON_LEASE_RENEWAL_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "USE_UPDATE_DNS_ON_LEASE_RENEWAL_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseUpdateDnsOnLeaseRenewal("USE_UPDATE_DNS_ON_LEASE_RENEWAL_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "USE_UPDATE_DNS_ON_LEASE_RENEWAL_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_UseZoneAssociations(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_use_zone_associations"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkUseZoneAssociations("USE_ZONE_ASSOCIATIONS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_zone_associations", "USE_ZONE_ASSOCIATIONS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkUseZoneAssociations("USE_ZONE_ASSOCIATIONS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "use_zone_associations", "USE_ZONE_ASSOCIATIONS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_Vlans(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_vlans"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkVlans("VLANS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "vlans", "VLANS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkVlans("VLANS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "vlans", "VLANS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

//     func TestAccNetworkResource_ZoneAssociations(t *testing.T) {
//     var resourceName = "nios_ipam_network.test_zone_associations"
//     var v ipam.Network

//     resource.ParallelTest(t, resource.TestCase{
//         PreCheck:                 func() { acctest.PreCheck(t) },
//         ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkZoneAssociations("ZONE_ASSOCIATIONS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "zone_associations", "ZONE_ASSOCIATIONS_REPLACE_ME"),
//                 ),
//             },
//             // Update and Read
// 			{
// 				Config: testAccNetworkZoneAssociations("ZONE_ASSOCIATIONS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "zone_associations", "ZONE_ASSOCIATIONS_UPDATE_REPLACE_ME"),
//                 ),
//             },
//    			// Delete testing automatically occurs in TestCase
//         },
//     })
// }

func testAccCheckNetworkExists(ctx context.Context, resourceName string, v *ipam.Network) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.IPAMAPI.
			NetworkAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForNetwork).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetNetworkResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetNetworkResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckNetworkDestroy(ctx context.Context, v *ipam.Network) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.IPAMAPI.
			NetworkAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForNetwork).
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

func testAccCheckNetworkDisappears(ctx context.Context, v *ipam.Network) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.IPAMAPI.
			NetworkAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccNetworkBasicConfig(network string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_ipam_network" "test" {
    network = %q
}
`, network)
}

func testAccNetworkAuthority(network, authority, useAuthority string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_authority" {
	network = %q
    authority = %q
    use_authority = %q
}
`, network, authority, useAuthority)
}

// func testAccNetworkAutoCreateReversezone(autoCreateReversezone string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_auto_create_reversezone" {
//     auto_create_reversezone = %q
// }
// `,autoCreateReversezone)
// }

// func testAccNetworkBootfile(bootfile string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_bootfile" {
//     bootfile = %q
// }
// `,bootfile)
// }

// func testAccNetworkBootserver(bootserver string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_bootserver" {
//     bootserver = %q
// }
// `,bootserver)
// }

// func testAccNetworkCloudInfo(cloudInfo string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_cloud_info" {
//     cloud_info = %q
// }
// `,cloudInfo)
// }

// func testAccNetworkCloudShared(cloudShared string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_cloud_shared" {
//     cloud_shared = %q
// }
// `,cloudShared)
// }

// func testAccNetworkComment(comment string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_comment" {
//     comment = %q
// }
// `,comment)
// }

// func testAccNetworkDdnsDomainname(ddnsDomainname string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_ddns_domainname" {
//     ddns_domainname = %q
// }
// `,ddnsDomainname)
// }

// func testAccNetworkDdnsGenerateHostname(ddnsGenerateHostname string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_ddns_generate_hostname" {
//     ddns_generate_hostname = %q
// }
// `,ddnsGenerateHostname)
// }

// func testAccNetworkDdnsServerAlwaysUpdates(ddnsServerAlwaysUpdates string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_ddns_server_always_updates" {
//     ddns_server_always_updates = %q
// }
// `,ddnsServerAlwaysUpdates)
// }

// func testAccNetworkDdnsTtl(ddnsTtl string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_ddns_ttl" {
//     ddns_ttl = %q
// }
// `,ddnsTtl)
// }

// func testAccNetworkDdnsUpdateFixedAddresses(ddnsUpdateFixedAddresses string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_ddns_update_fixed_addresses" {
//     ddns_update_fixed_addresses = %q
// }
// `,ddnsUpdateFixedAddresses)
// }

// func testAccNetworkDdnsUseOption81(ddnsUseOption81 string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_ddns_use_option81" {
//     ddns_use_option81 = %q
// }
// `,ddnsUseOption81)
// }

// func testAccNetworkDeleteReason(deleteReason string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_delete_reason" {
//     delete_reason = %q
// }
// `,deleteReason)
// }

// func testAccNetworkDenyBootp(denyBootp string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_deny_bootp" {
//     deny_bootp = %q
// }
// `,denyBootp)
// }

// func testAccNetworkDisable(disable string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_disable" {
//     disable = %q
// }
// `,disable)
// }

// func testAccNetworkDiscoveredBridgeDomain(discoveredBridgeDomain string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_discovered_bridge_domain" {
//     discovered_bridge_domain = %q
// }
// `,discoveredBridgeDomain)
// }

// func testAccNetworkDiscoveredTenant(discoveredTenant string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_discovered_tenant" {
//     discovered_tenant = %q
// }
// `,discoveredTenant)
// }

// func testAccNetworkDiscoveryBasicPollSettings(discoveryBasicPollSettings string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_discovery_basic_poll_settings" {
//     discovery_basic_poll_settings = %q
// }
// `,discoveryBasicPollSettings)
// }

// func testAccNetworkDiscoveryBlackoutSetting(discoveryBlackoutSetting string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_discovery_blackout_setting" {
//     discovery_blackout_setting = %q
// }
// `,discoveryBlackoutSetting)
// }

// func testAccNetworkDiscoveryMember(discoveryMember string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_discovery_member" {
//     discovery_member = %q
// }
// `,discoveryMember)
// }

// func testAccNetworkEmailList(emailList string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_email_list" {
//     email_list = %q
// }
// `,emailList)
// }

// func testAccNetworkEnableDdns(enableDdns string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_enable_ddns" {
//     enable_ddns = %q
// }
// `,enableDdns)
// }

// func testAccNetworkEnableDhcpThresholds(enableDhcpThresholds string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_enable_dhcp_thresholds" {
//     enable_dhcp_thresholds = %q
// }
// `,enableDhcpThresholds)
// }

// func testAccNetworkEnableDiscovery(enableDiscovery string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_enable_discovery" {
//     enable_discovery = %q
// }
// `,enableDiscovery)
// }

// func testAccNetworkEnableEmailWarnings(enableEmailWarnings string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_enable_email_warnings" {
//     enable_email_warnings = %q
// }
// `,enableEmailWarnings)
// }

// func testAccNetworkEnableIfmapPublishing(enableIfmapPublishing string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_enable_ifmap_publishing" {
//     enable_ifmap_publishing = %q
// }
// `,enableIfmapPublishing)
// }

// func testAccNetworkEnableImmediateDiscovery(enableImmediateDiscovery string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_enable_immediate_discovery" {
//     enable_immediate_discovery = %q
// }
// `,enableImmediateDiscovery)
// }

// func testAccNetworkEnablePxeLeaseTime(enablePxeLeaseTime string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_enable_pxe_lease_time" {
//     enable_pxe_lease_time = %q
// }
// `,enablePxeLeaseTime)
// }

// func testAccNetworkEnableSnmpWarnings(enableSnmpWarnings string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_enable_snmp_warnings" {
//     enable_snmp_warnings = %q
// }
// `,enableSnmpWarnings)
// }

// func testAccNetworkExtAttrs(extAttrs string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_extattrs" {
//     extattrs = %q
// }
// `,extAttrs)
// }

// func testAccNetworkFederatedRealms(federatedRealms string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_federated_realms" {
//     federated_realms = %q
// }
// `,federatedRealms)
// }

// func testAccNetworkHighWaterMark(highWaterMark string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_high_water_mark" {
//     high_water_mark = %q
// }
// `,highWaterMark)
// }

// func testAccNetworkHighWaterMarkReset(highWaterMarkReset string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_high_water_mark_reset" {
//     high_water_mark_reset = %q
// }
// `,highWaterMarkReset)
// }

// func testAccNetworkIgnoreDhcpOptionListRequest(ignoreDhcpOptionListRequest string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_ignore_dhcp_option_list_request" {
//     ignore_dhcp_option_list_request = %q
// }
// `,ignoreDhcpOptionListRequest)
// }

// func testAccNetworkIgnoreId(ignoreId string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_ignore_id" {
//     ignore_id = %q
// }
// `,ignoreId)
// }

// func testAccNetworkIgnoreMacAddresses(ignoreMacAddresses string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_ignore_mac_addresses" {
//     ignore_mac_addresses = %q
// }
// `,ignoreMacAddresses)
// }

// func testAccNetworkIpamEmailAddresses(ipamEmailAddresses string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_ipam_email_addresses" {
//     ipam_email_addresses = %q
// }
// `,ipamEmailAddresses)
// }

// func testAccNetworkIpamThresholdSettings(ipamThresholdSettings string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_ipam_threshold_settings" {
//     ipam_threshold_settings = %q
// }
// `,ipamThresholdSettings)
// }

// func testAccNetworkIpamTrapSettings(ipamTrapSettings string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_ipam_trap_settings" {
//     ipam_trap_settings = %q
// }
// `,ipamTrapSettings)
// }

// func testAccNetworkIpv4addr(ipv4addr string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_ipv4addr" {
//     ipv4addr = %q
// }
// `,ipv4addr)
// }

// func testAccNetworkLeaseScavengeTime(leaseScavengeTime string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_lease_scavenge_time" {
//     lease_scavenge_time = %q
// }
// `,leaseScavengeTime)
// }

// func testAccNetworkLogicFilterRules(logicFilterRules string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_logic_filter_rules" {
//     logic_filter_rules = %q
// }
// `,logicFilterRules)
// }

// func testAccNetworkLowWaterMark(lowWaterMark string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_low_water_mark" {
//     low_water_mark = %q
// }
// `,lowWaterMark)
// }

// func testAccNetworkLowWaterMarkReset(lowWaterMarkReset string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_low_water_mark_reset" {
//     low_water_mark_reset = %q
// }
// `,lowWaterMarkReset)
// }

// func testAccNetworkMembers(members string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_members" {
//     members = %q
// }
// `,members)
// }

// func testAccNetworkMgmPrivate(mgmPrivate string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_mgm_private" {
//     mgm_private = %q
// }
// `,mgmPrivate)
// }

// func testAccNetworkMsAdUserData(msAdUserData string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_ms_ad_user_data" {
//     ms_ad_user_data = %q
// }
// `,msAdUserData)
// }

// func testAccNetworkNetmask(netmask string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_netmask" {
//     netmask = %q
// }
// `,netmask)
// }

// func testAccNetworkNetwork(network string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_network" {
//     network = %q
// }
// `,network)
// }

// func testAccNetworkFuncCall(funcCall string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_func_call" {
//     func_call = %q
// }
// `,funcCall)
// }

// func testAccNetworkNetworkView(networkView string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_network_view" {
//     network_view = %q
// }
// `,networkView)
// }

// func testAccNetworkNextserver(nextserver string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_nextserver" {
//     nextserver = %q
// }
// `,nextserver)
// }

// func testAccNetworkOptions(options string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_options" {
//     options = %q
// }
// `,options)
// }

// func testAccNetworkPortControlBlackoutSetting(portControlBlackoutSetting string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_port_control_blackout_setting" {
//     port_control_blackout_setting = %q
// }
// `,portControlBlackoutSetting)
// }

// func testAccNetworkPxeLeaseTime(pxeLeaseTime string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_pxe_lease_time" {
//     pxe_lease_time = %q
// }
// `,pxeLeaseTime)
// }

// func testAccNetworkRecycleLeases(recycleLeases string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_recycle_leases" {
//     recycle_leases = %q
// }
// `,recycleLeases)
// }

// func testAccNetworkRestartIfNeeded(restartIfNeeded string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_restart_if_needed" {
//     restart_if_needed = %q
// }
// `,restartIfNeeded)
// }

// func testAccNetworkRirOrganization(rirOrganization string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_rir_organization" {
//     rir_organization = %q
// }
// `,rirOrganization)
// }

// func testAccNetworkRirRegistrationAction(rirRegistrationAction string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_rir_registration_action" {
//     rir_registration_action = %q
// }
// `,rirRegistrationAction)
// }

// func testAccNetworkRirRegistrationStatus(rirRegistrationStatus string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_rir_registration_status" {
//     rir_registration_status = %q
// }
// `,rirRegistrationStatus)
// }

// func testAccNetworkSamePortControlDiscoveryBlackout(samePortControlDiscoveryBlackout string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_same_port_control_discovery_blackout" {
//     same_port_control_discovery_blackout = %q
// }
// `,samePortControlDiscoveryBlackout)
// }

// func testAccNetworkSendRirRequest(sendRirRequest string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_send_rir_request" {
//     send_rir_request = %q
// }
// `,sendRirRequest)
// }

// func testAccNetworkSubscribeSettings(subscribeSettings string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_subscribe_settings" {
//     subscribe_settings = %q
// }
// `,subscribeSettings)
// }

// func testAccNetworkTemplate(template string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_template" {
//     template = %q
// }
// `,template)
// }

// func testAccNetworkUnmanaged(unmanaged string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_unmanaged" {
//     unmanaged = %q
// }
// `,unmanaged)
// }

// func testAccNetworkUpdateDnsOnLeaseRenewal(updateDnsOnLeaseRenewal string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_update_dns_on_lease_renewal" {
//     update_dns_on_lease_renewal = %q
// }
// `,updateDnsOnLeaseRenewal)
// }

// func testAccNetworkUseAuthority(useAuthority string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_authority" {
//     use_authority = %q
// }
// `,useAuthority)
// }

// func testAccNetworkUseBlackoutSetting(useBlackoutSetting string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_blackout_setting" {
//     use_blackout_setting = %q
// }
// `,useBlackoutSetting)
// }

// func testAccNetworkUseBootfile(useBootfile string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_bootfile" {
//     use_bootfile = %q
// }
// `,useBootfile)
// }

// func testAccNetworkUseBootserver(useBootserver string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_bootserver" {
//     use_bootserver = %q
// }
// `,useBootserver)
// }

// func testAccNetworkUseDdnsDomainname(useDdnsDomainname string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_ddns_domainname" {
//     use_ddns_domainname = %q
// }
// `,useDdnsDomainname)
// }

// func testAccNetworkUseDdnsGenerateHostname(useDdnsGenerateHostname string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_ddns_generate_hostname" {
//     use_ddns_generate_hostname = %q
// }
// `,useDdnsGenerateHostname)
// }

// func testAccNetworkUseDdnsTtl(useDdnsTtl string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_ddns_ttl" {
//     use_ddns_ttl = %q
// }
// `,useDdnsTtl)
// }

// func testAccNetworkUseDdnsUpdateFixedAddresses(useDdnsUpdateFixedAddresses string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_ddns_update_fixed_addresses" {
//     use_ddns_update_fixed_addresses = %q
// }
// `,useDdnsUpdateFixedAddresses)
// }

// func testAccNetworkUseDdnsUseOption81(useDdnsUseOption81 string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_ddns_use_option81" {
//     use_ddns_use_option81 = %q
// }
// `,useDdnsUseOption81)
// }

// func testAccNetworkUseDenyBootp(useDenyBootp string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_deny_bootp" {
//     use_deny_bootp = %q
// }
// `,useDenyBootp)
// }

// func testAccNetworkUseDiscoveryBasicPollingSettings(useDiscoveryBasicPollingSettings string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_discovery_basic_polling_settings" {
//     use_discovery_basic_polling_settings = %q
// }
// `,useDiscoveryBasicPollingSettings)
// }

// func testAccNetworkUseEmailList(useEmailList string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_email_list" {
//     use_email_list = %q
// }
// `,useEmailList)
// }

// func testAccNetworkUseEnableDdns(useEnableDdns string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_enable_ddns" {
//     use_enable_ddns = %q
// }
// `,useEnableDdns)
// }

// func testAccNetworkUseEnableDhcpThresholds(useEnableDhcpThresholds string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_enable_dhcp_thresholds" {
//     use_enable_dhcp_thresholds = %q
// }
// `,useEnableDhcpThresholds)
// }

// func testAccNetworkUseEnableDiscovery(useEnableDiscovery string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_enable_discovery" {
//     use_enable_discovery = %q
// }
// `,useEnableDiscovery)
// }

// func testAccNetworkUseEnableIfmapPublishing(useEnableIfmapPublishing string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_enable_ifmap_publishing" {
//     use_enable_ifmap_publishing = %q
// }
// `,useEnableIfmapPublishing)
// }

// func testAccNetworkUseIgnoreDhcpOptionListRequest(useIgnoreDhcpOptionListRequest string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_ignore_dhcp_option_list_request" {
//     use_ignore_dhcp_option_list_request = %q
// }
// `,useIgnoreDhcpOptionListRequest)
// }

// func testAccNetworkUseIgnoreId(useIgnoreId string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_ignore_id" {
//     use_ignore_id = %q
// }
// `,useIgnoreId)
// }

// func testAccNetworkUseIpamEmailAddresses(useIpamEmailAddresses string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_ipam_email_addresses" {
//     use_ipam_email_addresses = %q
// }
// `,useIpamEmailAddresses)
// }

// func testAccNetworkUseIpamThresholdSettings(useIpamThresholdSettings string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_ipam_threshold_settings" {
//     use_ipam_threshold_settings = %q
// }
// `,useIpamThresholdSettings)
// }

// func testAccNetworkUseIpamTrapSettings(useIpamTrapSettings string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_ipam_trap_settings" {
//     use_ipam_trap_settings = %q
// }
// `,useIpamTrapSettings)
// }

// func testAccNetworkUseLeaseScavengeTime(useLeaseScavengeTime string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_lease_scavenge_time" {
//     use_lease_scavenge_time = %q
// }
// `,useLeaseScavengeTime)
// }

// func testAccNetworkUseLogicFilterRules(useLogicFilterRules string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_logic_filter_rules" {
//     use_logic_filter_rules = %q
// }
// `,useLogicFilterRules)
// }

// func testAccNetworkUseMgmPrivate(useMgmPrivate string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_mgm_private" {
//     use_mgm_private = %q
// }
// `,useMgmPrivate)
// }

// func testAccNetworkUseNextserver(useNextserver string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_nextserver" {
//     use_nextserver = %q
// }
// `,useNextserver)
// }

// func testAccNetworkUseOptions(useOptions string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_options" {
//     use_options = %q
// }
// `,useOptions)
// }

// func testAccNetworkUsePxeLeaseTime(usePxeLeaseTime string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_pxe_lease_time" {
//     use_pxe_lease_time = %q
// }
// `,usePxeLeaseTime)
// }

// func testAccNetworkUseRecycleLeases(useRecycleLeases string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_recycle_leases" {
//     use_recycle_leases = %q
// }
// `,useRecycleLeases)
// }

// func testAccNetworkUseSubscribeSettings(useSubscribeSettings string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_subscribe_settings" {
//     use_subscribe_settings = %q
// }
// `,useSubscribeSettings)
// }

// func testAccNetworkUseUpdateDnsOnLeaseRenewal(useUpdateDnsOnLeaseRenewal string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_update_dns_on_lease_renewal" {
//     use_update_dns_on_lease_renewal = %q
// }
// `,useUpdateDnsOnLeaseRenewal)
// }

// func testAccNetworkUseZoneAssociations(useZoneAssociations string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_use_zone_associations" {
//     use_zone_associations = %q
// }
// `,useZoneAssociations)
// }

// func testAccNetworkVlans(vlans string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_vlans" {
//     vlans = %q
// }
// `,vlans)
// }

// func testAccNetworkZoneAssociations(zoneAssociations string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_zone_associations" {
//     zone_associations = %q
// }
// `,zoneAssociations)
// }
