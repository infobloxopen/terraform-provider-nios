package ipam_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/ipam"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/acctest"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/utils"
)

var readableAttributesForNetwork = "authority,bootfile,bootserver,cloud_info,cloud_shared,comment,conflict_count,ddns_domainname,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,ddns_update_fixed_addresses,ddns_use_option81,deny_bootp,dhcp_utilization,dhcp_utilization_status,disable,discover_now_status,discovered_bgp_as,discovered_bridge_domain,discovered_tenant,discovered_vlan_id,discovered_vlan_name,discovered_vrf_description,discovered_vrf_name,discovered_vrf_rd,discovery_basic_poll_settings,discovery_blackout_setting,discovery_engine_type,discovery_member,dynamic_hosts,email_list,enable_ddns,enable_dhcp_thresholds,enable_discovery,enable_email_warnings,enable_ifmap_publishing,enable_pxe_lease_time,enable_snmp_warnings,endpoint_sources,extattrs,federated_realms,high_water_mark,high_water_mark_reset,ignore_dhcp_option_list_request,ignore_id,ignore_mac_addresses,ipam_email_addresses,ipam_threshold_settings,ipam_trap_settings,ipv4addr,last_rir_registration_update_sent,last_rir_registration_update_status,lease_scavenge_time,logic_filter_rules,low_water_mark,low_water_mark_reset,members,mgm_private,mgm_private_overridable,ms_ad_user_data,netmask,network,network_container,network_view,nextserver,options,port_control_blackout_setting,pxe_lease_time,recycle_leases,rir,rir_organization,rir_registration_status,same_port_control_discovery_blackout,static_hosts,subscribe_settings,total_hosts,unmanaged,unmanaged_count,update_dns_on_lease_renewal,use_authority,use_blackout_setting,use_bootfile,use_bootserver,use_ddns_domainname,use_ddns_generate_hostname,use_ddns_ttl,use_ddns_update_fixed_addresses,use_ddns_use_option81,use_deny_bootp,use_discovery_basic_polling_settings,use_email_list,use_enable_ddns,use_enable_dhcp_thresholds,use_enable_discovery,use_enable_ifmap_publishing,use_ignore_dhcp_option_list_request,use_ignore_id,use_ipam_email_addresses,use_ipam_threshold_settings,use_ipam_trap_settings,use_lease_scavenge_time,use_logic_filter_rules,use_mgm_private,use_nextserver,use_options,use_pxe_lease_time,use_recycle_leases,use_subscribe_settings,use_update_dns_on_lease_renewal,use_zone_associations,utilization,utilization_update,vlans,zone_associations"

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

func TestAccNetworkResource_Bootfile(t *testing.T) {
	var resourceName = "nios_ipam_network.test_bootfile"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkBootfile(network, "bootfile", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "bootfile"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkBootfile(network, "bootfile_updated", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "bootfile_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_Bootserver(t *testing.T) {
	var resourceName = "nios_ipam_network.test_bootserver"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkBootserver(network, "test_bootserver", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "test_bootserver"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkBootserver(network, "test_bootserver_updated", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "test_bootserver_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_CloudInfo(t *testing.T) {
	var resourceName = "nios_ipam_network.test_cloud_info"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkCloudInfo(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
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

func TestAccNetworkResource_CloudShared(t *testing.T) {
	var resourceName = "nios_ipam_network.test_cloud_shared"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkCloudShared(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_shared", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkCloudShared(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_shared", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_Comment(t *testing.T) {
	var resourceName = "nios_ipam_network.test_comment"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkComment(network, "test comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkComment(network, "updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_DdnsDomainname(t *testing.T) {
	var resourceName = "nios_ipam_network.test_ddns_domainname"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkDdnsDomainname(network, "test.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "test.com"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkDdnsDomainname(network, "testupdated.com", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "testupdated.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_DdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_ipam_network.test_ddns_generate_hostname"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkDdnsGenerateHostname(network, "true", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkDdnsGenerateHostname(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_DdnsServerAlwaysUpdates(t *testing.T) {
	var resourceName = "nios_ipam_network.test_ddns_server_always_updates"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkDdnsServerAlwaysUpdates(network, "true", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkDdnsServerAlwaysUpdates(network, "false", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_DdnsTtl(t *testing.T) {
	var resourceName = "nios_ipam_network.test_ddns_ttl"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkDdnsTtl(network, "1", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "1"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkDdnsTtl(network, "600", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "600"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_DdnsUpdateFixedAddresses(t *testing.T) {
	var resourceName = "nios_ipam_network.test_ddns_update_fixed_addresses"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkDdnsUpdateFixedAddresses(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_update_fixed_addresses", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkDdnsUpdateFixedAddresses(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_update_fixed_addresses", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_DdnsUseOption81(t *testing.T) {
	var resourceName = "nios_ipam_network.test_ddns_use_option81"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkDdnsUseOption81(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkDdnsUseOption81(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_DeleteReason(t *testing.T) {
	var resourceName = "nios_ipam_network.test_delete_reason"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkDeleteReason(network, "Test deletion reason"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delete_reason", "Test deletion reason"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkDeleteReason(network, "Updated deletion reason"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delete_reason", "Updated deletion reason"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_DenyBootp(t *testing.T) {
	var resourceName = "nios_ipam_network.test_deny_bootp"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkDenyBootp(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkDenyBootp(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_Disable(t *testing.T) {
	var resourceName = "nios_ipam_network.test_disable"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkDisable(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkDisable(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_DiscoveredBridgeDomain(t *testing.T) {
	var resourceName = "nios_ipam_network.test_discovered_bridge_domain"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkDiscoveredBridgeDomain(network, "bridge-domain-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovered_bridge_domain", "bridge-domain-1"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkDiscoveredBridgeDomain(network, "bridge-domain-2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovered_bridge_domain", "bridge-domain-2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_DiscoveredTenant(t *testing.T) {
	var resourceName = "nios_ipam_network.test_discovered_tenant"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkDiscoveredTenant(network, "tenant-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovered_tenant", "tenant-1"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkDiscoveredTenant(network, "tenant-2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovered_tenant", "tenant-2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_DiscoveryBasicPollSettings(t *testing.T) {
	var resourceName = "nios_ipam_network.test_discovery_basic_poll_settings"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkDiscoveryBasicPollSettings(network, "true", "true", "false", "default", "false", "false", "1", "false", "false", "true", "PERIODIC", "3600", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.auto_arp_refresh_before_switch_port_polling", "true"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.cli_collection", "true"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.complete_ping_sweep", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.credential_group", "default"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.device_profile", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.netbios_scanning", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.polling_frequency_modifier", "1"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.port_scanning", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.smart_subnet_ping_sweep", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.snmp_collection", "true"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.switch_port_data_collection_polling", "PERIODIC"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.switch_port_data_collection_polling_interval", "3600"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.use_global_polling_frequency_modifier", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkDiscoveryBasicPollSettings(network, "false", "true", "false", "default", "false", "false", "1", "false", "false", "true", "PERIODIC", "3600", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.auto_arp_refresh_before_switch_port_polling", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.cli_collection", "true"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.complete_ping_sweep", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.credential_group", "default"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.device_profile", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.netbios_scanning", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.polling_frequency_modifier", "1"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.port_scanning", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.smart_subnet_ping_sweep", "false"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.snmp_collection", "true"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.switch_port_data_collection_polling", "PERIODIC"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.switch_port_data_collection_polling_interval", "3600"),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings.use_global_polling_frequency_modifier", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_DiscoveryBlackoutSetting(t *testing.T) {
	var resourceName = "nios_ipam_network.test_discovery_blackout_setting"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkDiscoveryBlackoutSetting(network, "false", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_blackout_setting.enable_blackout", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_EmailList(t *testing.T) {
	var resourceName = "nios_ipam_network.test_email_list"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkEmailList(network, "admin@example.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_list.0", "admin@example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkEmailList(network, "admin@updated.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_list.0", "admin@updated.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_EnableDdns(t *testing.T) {
	var resourceName = "nios_ipam_network.test_enable_ddns"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkEnableDdns(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkEnableDdns(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_EnableDhcpThresholds(t *testing.T) {
	var resourceName = "nios_ipam_network.test_enable_dhcp_thresholds"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkEnableDhcpThresholds(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_dhcp_thresholds", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkEnableDhcpThresholds(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_dhcp_thresholds", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworkResource_EnableEmailWarnings(t *testing.T) {
	var resourceName = "nios_ipam_network.test_enable_email_warnings"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkEnableEmailWarnings(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_email_warnings", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkEnableEmailWarnings(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_email_warnings", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_EnableIfmapPublishing(t *testing.T) {
	var resourceName = "nios_ipam_network.test_enable_ifmap_publishing"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkEnableIfmapPublishing(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ifmap_publishing", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ifmap_publishing", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkEnableIfmapPublishing(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ifmap_publishing", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ifmap_publishing", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_EnablePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_ipam_network.test_enable_pxe_lease_time"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkEnablePxeLeaseTime(network, "100", "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkEnablePxeLeaseTime(network, "100", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_EnableSnmpWarnings(t *testing.T) {
	var resourceName = "nios_ipam_network.test_enable_snmp_warnings"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkEnableSnmpWarnings(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_snmp_warnings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkEnableSnmpWarnings(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_snmp_warnings", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_ipam_network.test_extattrs"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkExtAttrs(network, map[string]string{"Site": extAttrValue1}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkExtAttrs(network, map[string]string{"Site": extAttrValue2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_HighWaterMark(t *testing.T) {
	var resourceName = "nios_ipam_network.test_high_water_mark"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkHighWaterMark(network, "80"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "80"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkHighWaterMark(network, "90"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "90"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_HighWaterMarkReset(t *testing.T) {
	var resourceName = "nios_ipam_network.test_high_water_mark_reset"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkHighWaterMarkReset(network, "70"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "70"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkHighWaterMarkReset(network, "80"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "80"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_IgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_ipam_network.test_ignore_dhcp_option_list_request"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkIgnoreDhcpOptionListRequest(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkIgnoreDhcpOptionListRequest(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_IgnoreId(t *testing.T) {
	var resourceName = "nios_ipam_network.test_ignore_id"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkIgnoreId(network, "NONE", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_id", "NONE"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkIgnoreId(network, "MACADDR", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_id", "MACADDR"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_IgnoreMacAddresses(t *testing.T) {
	var resourceName = "nios_ipam_network.test_ignore_mac_addresses"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkIgnoreMacAddresses(network, "aa:bb:cc:dd:ee:ff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_mac_addresses.0", "aa:bb:cc:dd:ee:ff"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkIgnoreMacAddresses(network, "ff:ee:dd:cc:bb:aa"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_mac_addresses.0", "ff:ee:dd:cc:bb:aa"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_IpamEmailAddresses(t *testing.T) {
	var resourceName = "nios_ipam_network.test_ipam_email_addresses"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkIpamEmailAddresses(network, "testuser@infoblox.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_email_addresses.0", "testuser@infoblox.com"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkIpamEmailAddresses(network, "testuser2@infoblox.com", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_email_addresses.0", "testuser2@infoblox.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_IpamThresholdSettings(t *testing.T) {
	var resourceName = "nios_ipam_network.test_ipam_threshold_settings"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkIpamThresholdSettings(network, "85", "95", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_threshold_settings.reset_value", "85"),
					resource.TestCheckResourceAttr(resourceName, "ipam_threshold_settings.trigger_value", "95"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkIpamThresholdSettings(network, "75", "80", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_threshold_settings.reset_value", "75"),
					resource.TestCheckResourceAttr(resourceName, "ipam_threshold_settings.trigger_value", "80"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_IpamTrapSettings(t *testing.T) {
	var resourceName = "nios_ipam_network.test_ipam_trap_settings"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkIpamTrapSettings(network, "false", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_trap_settings.enable_email_warnings", "false"),
					resource.TestCheckResourceAttr(resourceName, "ipam_trap_settings.enable_snmp_warnings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkIpamTrapSettings(network, "true", "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_trap_settings.enable_email_warnings", "true"),
					resource.TestCheckResourceAttr(resourceName, "ipam_trap_settings.enable_snmp_warnings", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// func TestAccNetworkResource_Ipv4addr(t *testing.T) {
// 	var resourceName = "nios_ipam_network.test_ipv4addr"
// 	var v ipam.Network

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkIpv4addr("IPV4ADDR_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "IPV4ADDR_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkIpv4addr("IPV4ADDR_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "IPV4ADDR_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

func TestAccNetworkResource_LeaseScavengeTime(t *testing.T) {
	var resourceName = "nios_ipam_network.test_lease_scavenge_time"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkLeaseScavengeTime(network, "-1", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "-1"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkLeaseScavengeTime(network, "86400", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "86400"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_LowWaterMark(t *testing.T) {
	var resourceName = "nios_ipam_network.test_low_water_mark"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkLowWaterMark(network, "0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "0"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkLowWaterMark(network, "50"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "50"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_LowWaterMarkReset(t *testing.T) {
	var resourceName = "nios_ipam_network.test_low_water_mark_reset"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkLowWaterMarkReset(network, "10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkLowWaterMarkReset(network, "20"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_MgmPrivate(t *testing.T) {
	var resourceName = "nios_ipam_network.test_mgm_private"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkMgmPrivate(network, "false", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mgm_private", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkMgmPrivate(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mgm_private", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_Netmask(t *testing.T) {
	var resourceName = "nios_ipam_network.test_netmask"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkNetmask(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "netmask", network[strings.LastIndex(network, "/")+1:])),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_Network(t *testing.T) {
	var resourceName = "nios_ipam_network.test_network"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkNetwork(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_NetworkView(t *testing.T) {
	var resourceName = "nios_ipam_network.test_network_view"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkNetworkView(network, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_Nextserver(t *testing.T) {
	var resourceName = "nios_ipam_network.test_nextserver"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkNextserver(network, "1.1.1.1", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", "1.1.1.1"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkNextserver(network, "1.1.1.2", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", "1.1.1.2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_Options(t *testing.T) {
	var resourceName = "nios_ipam_network.test_options"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkOptions(network, "dhcp-lease-time", "51", "7200", "DHCP", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "7200"),
					resource.TestCheckResourceAttr(resourceName, "options.0.vendor_class", "DHCP"),
					resource.TestCheckResourceAttr(resourceName, "options.0.use_option", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkOptions(network, "dhcp-lease-time", "51", "7300", "DHCP", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "7300"),
					resource.TestCheckResourceAttr(resourceName, "options.0.vendor_class", "DHCP"),
					resource.TestCheckResourceAttr(resourceName, "options.0.use_option", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_PortControlBlackoutSetting(t *testing.T) {
	var resourceName = "nios_ipam_network.test_port_control_blackout_setting"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkPortControlBlackoutSetting(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port_control_blackout_setting.enable_blackout", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_PxeLeaseTime(t *testing.T) {
	var resourceName = "nios_ipam_network.test_pxe_lease_time"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkPxeLeaseTime(network, "0", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "0"),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkPxeLeaseTime(network, "40000", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "40000"),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_RecycleLeases(t *testing.T) {
	var resourceName = "nios_ipam_network.test_recycle_leases"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkRecycleLeases(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkRecycleLeases(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_RirRegistrationStatus(t *testing.T) {
	var resourceName = "nios_ipam_network.test_rir_registration_status"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkRirRegistrationStatus(network, "NOT_REGISTERED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_status", "NOT_REGISTERED"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_SamePortControlDiscoveryBlackout(t *testing.T) {
	var resourceName = "nios_ipam_network.test_same_port_control_discovery_blackout"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkSamePortControlDiscoveryBlackout(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkSamePortControlDiscoveryBlackout(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// func TestAccNetworkResource_SubscribeSettings(t *testing.T) {
// 	var resourceName = "nios_ipam_network.test_subscribe_settings"
// 	var v ipam.Network

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccNetworkSubscribeSettings("SUBSCRIBE_SETTINGS_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "subscribe_settings", "SUBSCRIBE_SETTINGS_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccNetworkSubscribeSettings("SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckNetworkExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "subscribe_settings", "SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

func TestAccNetworkResource_Unmanaged(t *testing.T) {
	var resourceName = "nios_ipam_network.test_unmanaged"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUnmanaged(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "unmanaged", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_ipam_network.test_update_dns_on_lease_renewal"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUpdateDnsOnLeaseRenewal(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUpdateDnsOnLeaseRenewal(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseAuthority(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_authority"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseAuthority(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_authority", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseAuthority(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_authority", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseBlackoutSetting(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_blackout_setting"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseBlackoutSetting(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseBlackoutSetting(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseBootfile(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_bootfile"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseBootfile(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseBootfile(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseBootserver(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_bootserver"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseBootserver(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseBootserver(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseDdnsDomainname(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_ddns_domainname"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseDdnsDomainname(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseDdnsDomainname(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseDdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_ddns_generate_hostname"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseDdnsGenerateHostname(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseDdnsGenerateHostname(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseDdnsTtl(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_ddns_ttl"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseDdnsTtl(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseDdnsTtl(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseDdnsUpdateFixedAddresses(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_ddns_update_fixed_addresses"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseDdnsUpdateFixedAddresses(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_update_fixed_addresses", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseDdnsUpdateFixedAddresses(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_update_fixed_addresses", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseDdnsUseOption81(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_ddns_use_option81"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseDdnsUseOption81(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseDdnsUseOption81(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseDenyBootp(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_deny_bootp"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseDenyBootp(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseDenyBootp(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseDiscoveryBasicPollingSettings(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_discovery_basic_polling_settings"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseDiscoveryBasicPollingSettings(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseDiscoveryBasicPollingSettings(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseEmailList(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_email_list"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseEmailList(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_email_list", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseEmailList(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_email_list", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseEnableDdns(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_enable_ddns"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseEnableDdns(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseEnableDdns(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseEnableDhcpThresholds(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_enable_dhcp_thresholds"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseEnableDhcpThresholds(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseEnableDhcpThresholds(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseEnableDiscovery(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_enable_discovery"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseEnableDiscovery(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseEnableDiscovery(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseEnableIfmapPublishing(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_enable_ifmap_publishing"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseEnableIfmapPublishing(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ifmap_publishing", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseEnableIfmapPublishing(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ifmap_publishing", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseIgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_ignore_dhcp_option_list_request"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseIgnoreDhcpOptionListRequest(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseIgnoreDhcpOptionListRequest(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseIgnoreId(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_ignore_id"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseIgnoreId(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_id", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseIgnoreId(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_id", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseIpamEmailAddresses(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_ipam_email_addresses"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseIpamEmailAddresses(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_email_addresses", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseIpamEmailAddresses(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_email_addresses", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseIpamThresholdSettings(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_ipam_threshold_settings"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseIpamThresholdSettings(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_threshold_settings", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseIpamThresholdSettings(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_threshold_settings", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseIpamTrapSettings(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_ipam_trap_settings"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseIpamTrapSettings(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_trap_settings", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseIpamTrapSettings(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_trap_settings", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseLeaseScavengeTime(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_lease_scavenge_time"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseLeaseScavengeTime(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseLeaseScavengeTime(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseLogicFilterRules(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_logic_filter_rules"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseLogicFilterRules(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseLogicFilterRules(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseMgmPrivate(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_mgm_private"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseMgmPrivate(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_mgm_private", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseNextserver(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_nextserver"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseNextserver(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseNextserver(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseOptions(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_options"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseOptions(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseOptions(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UsePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_pxe_lease_time"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUsePxeLeaseTime(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUsePxeLeaseTime(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseRecycleLeases(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_recycle_leases"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseRecycleLeases(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseRecycleLeases(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseSubscribeSettings(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_subscribe_settings"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseSubscribeSettings(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_subscribe_settings", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseUpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_update_dns_on_lease_renewal"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseUpdateDnsOnLeaseRenewal(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkUseUpdateDnsOnLeaseRenewal(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkResource_UseZoneAssociations(t *testing.T) {
	var resourceName = "nios_ipam_network.test_use_zone_associations"
	var v ipam.Network
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkUseZoneAssociations(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_zone_associations", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

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

func testAccNetworkBootfile(network, bootfile, useBootfile string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_bootfile" {
    network = %q
    bootfile = %q
    use_bootfile = %q
}
`, network, bootfile, useBootfile)
}

func testAccNetworkBootserver(network, bootserver, useBootserver string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_bootserver" {
    network = %q
    bootserver = %q
    use_bootserver = %q
}
`, network, bootserver, useBootserver)
}

func testAccNetworkCloudInfo(network string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_cloud_info" {
    network = %q
}
`, network)
}

func testAccNetworkCloudShared(network, cloudShared string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_cloud_shared" {
    network = %q
    cloud_shared = %q
}
`, network, cloudShared)
}

func testAccNetworkComment(network, comment string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_comment" {
    network = %q
    comment = %q
}
`, network, comment)
}

func testAccNetworkDdnsDomainname(network, ddnsDomainname, useDdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_ddns_domainname" {
    network = %q
    ddns_domainname = %q
    use_ddns_domainname = %q
}
`, network, ddnsDomainname, useDdnsDomainname)
}

func testAccNetworkDdnsGenerateHostname(network, ddnsGenerateHostname, useDdnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_ddns_generate_hostname" {
    network = %q
    ddns_generate_hostname = %q
    use_ddns_generate_hostname = %q
}
`, network, ddnsGenerateHostname, useDdnsGenerateHostname)
}

func testAccNetworkDdnsServerAlwaysUpdates(network, ddnsServerAlwaysUpdates, ddnsUseOption81, useDdnsUseOption81 string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_ddns_server_always_updates" {
    network = %q
    ddns_server_always_updates = %q
    ddns_use_option81 = %q
    use_ddns_use_option81 = %q
}
`, network, ddnsServerAlwaysUpdates, ddnsUseOption81, useDdnsUseOption81)
}

func testAccNetworkDdnsTtl(network, ddnsTtl, useDdnsTtl string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_ddns_ttl" {
    network = %q
    ddns_ttl = %q
    use_ddns_ttl = %q
}
`, network, ddnsTtl, useDdnsTtl)
}

func testAccNetworkDdnsUpdateFixedAddresses(network, ddnsUpdateFixedAddresses, useDdnsUpdateFixedAddresses string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_ddns_update_fixed_addresses" {
    network = %q
    ddns_update_fixed_addresses = %q
    use_ddns_update_fixed_addresses = %q
}
`, network, ddnsUpdateFixedAddresses, useDdnsUpdateFixedAddresses)
}

func testAccNetworkDdnsUseOption81(network, ddnsUseOption81, useDdnsUseOption81 string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_ddns_use_option81" {
    network = %q
    ddns_use_option81 = %q
    use_ddns_use_option81 = %q
}
`, network, ddnsUseOption81, useDdnsUseOption81)
}

func testAccNetworkDeleteReason(network, deleteReason string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_delete_reason" {
    network = %q
    delete_reason = %q
}
`, network, deleteReason)
}

func testAccNetworkDenyBootp(network, denyBootp, useDenyBootp string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_deny_bootp" {
    network = %q
    deny_bootp = %q
    use_deny_bootp = %q
}
`, network, denyBootp, useDenyBootp)
}

func testAccNetworkDisable(network, disable string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_disable" {
    network = %q
    disable = %q
}
`, network, disable)
}

func testAccNetworkDiscoveredBridgeDomain(network, discoveredBridgeDomain string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_discovered_bridge_domain" {
    network = %q
    discovered_bridge_domain = %q
}
`, network, discoveredBridgeDomain)
}

func testAccNetworkDiscoveredTenant(network, discoveredTenant string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_discovered_tenant" {
    network = %q
    discovered_tenant = %q
}
`, network, discoveredTenant)
}

func testAccNetworkDiscoveryBasicPollSettings(network, autoArpRefreshBeforeSwitchPortPolling, cliCollection, completePingSweep, credentialGroup, deviceProfile, netbiosScanning, pollingFrequencyModifier, portScanning, smartSubnetPingSweep, snmpCollection, switchPortDataCollectionPolling, switchPortDataCollectionPollingInterval, useGlobalPollingFrequencyModifier, useDiscoveryBasicPollSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_discovery_basic_poll_settings" {
  network = %q
    discovery_basic_poll_settings = {
        auto_arp_refresh_before_switch_port_polling = %s
        cli_collection = %s
        complete_ping_sweep = %s
        credential_group = %q
        device_profile = %s
        netbios_scanning = %s
        polling_frequency_modifier = %q
        port_scanning = %s
        smart_subnet_ping_sweep = %s
        snmp_collection = %s
        switch_port_data_collection_polling = %q
        switch_port_data_collection_polling_interval = %s
        use_global_polling_frequency_modifier = %s
    }
    use_discovery_basic_polling_settings = %s
}
`, network, autoArpRefreshBeforeSwitchPortPolling, cliCollection, completePingSweep, credentialGroup, deviceProfile, netbiosScanning, pollingFrequencyModifier, portScanning, smartSubnetPingSweep, snmpCollection, switchPortDataCollectionPolling, switchPortDataCollectionPollingInterval, useGlobalPollingFrequencyModifier, useDiscoveryBasicPollSettings)
}

func testAccNetworkDiscoveryBlackoutSetting(network, enabledBlackout, useBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_discovery_blackout_setting" {
  network = %q
    discovery_blackout_setting = {
		enabled_blackout = %q
	}
    use_blackout_setting = %q
}
`, network, enabledBlackout, useBlackoutSetting)
}

func testAccNetworkEmailList(network, emailList, useEmailList string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_email_list" {
  network    = %q
  email_list = [%q]
  use_email_list = %q
}
`, network, emailList, useEmailList)
}

func testAccNetworkEnableDdns(network, enableDdns, useEnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_enable_ddns" {
  network = %q
  enable_ddns = %q
  use_enable_ddns = %q
}
`, network, enableDdns, useEnableDdns)
}

func testAccNetworkEnableDhcpThresholds(network, enableDhcpThresholds, useEnableDhcpThresholds string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_enable_dhcp_thresholds" {
  network                  = %q
  enable_dhcp_thresholds   = %q
  use_enable_dhcp_thresholds = %q
}
`, network, enableDhcpThresholds, useEnableDhcpThresholds)
}

func testAccNetworkEnableEmailWarnings(network, enableEmailWarnings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_enable_email_warnings" {
  network               = "%s"
  enable_email_warnings = "%s"
}
`, network, enableEmailWarnings)
}

func testAccNetworkEnableIfmapPublishing(network, enableIfmapPublishing, useEnableIfmapPublishing string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_enable_ifmap_publishing" {
  network                 = %q
  enable_ifmap_publishing = %q
  use_enable_ifmap_publishing = %q
}
`, network, enableIfmapPublishing, useEnableIfmapPublishing)
}

func testAccNetworkEnablePxeLeaseTime(network, pxeLeaseTime, enablePxeLeaseTime, usePxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_enable_pxe_lease_time" {
  network = %q
    pxe_lease_time = %q
    enable_pxe_lease_time = %q
	use_pxe_lease_time = %q
}
`, network, pxeLeaseTime, enablePxeLeaseTime, usePxeLeaseTime)
}

func testAccNetworkEnableSnmpWarnings(network, enableSnmpWarnings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_enable_snmp_warnings" {
  network              = "%s"
  enable_snmp_warnings = "%s"
}
`, network, enableSnmpWarnings)
}

func testAccNetworkExtAttrs(network string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_extattrs" {
  network  = %q
  extattrs = %s
}
`, network, extattrsStr)
}

func testAccNetworkHighWaterMark(network, highWaterMark string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_high_water_mark" {
    network = %q
    high_water_mark = %q
}
`, network, highWaterMark)
}

func testAccNetworkHighWaterMarkReset(network, highWaterMarkReset string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_high_water_mark_reset" {
    network = %q
    high_water_mark_reset = %q
}
`, network, highWaterMarkReset)
}

func testAccNetworkIgnoreDhcpOptionListRequest(network, ignoreDhcpOptionListRequest, useIgnoreDhcpOptionListRequest string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_ignore_dhcp_option_list_request" {
    network = %q
    ignore_dhcp_option_list_request = %q
    use_ignore_dhcp_option_list_request = %q
}
`, network, ignoreDhcpOptionListRequest, useIgnoreDhcpOptionListRequest)
}

func testAccNetworkIgnoreId(network, ignoreId, useIgnoreId, useBootfile string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_ignore_id" {
    network = %q
    ignore_id = %q
    use_ignore_id = %q
    use_bootfile = %q
}
`, network, ignoreId, useIgnoreId, useBootfile)
}

func testAccNetworkIgnoreMacAddresses(network, ignoreMacAddresses string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_ignore_mac_addresses" {
    network = %q
    ignore_mac_addresses = [%q]
}
`, network, ignoreMacAddresses)
}

func testAccNetworkIpamEmailAddresses(network, ipamEmailAddresses, useIpamEmailAddresses string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_ipam_email_addresses" {
    network = %q
    ipam_email_addresses = [%q]
    use_ipam_email_addresses = %q
}
`, network, ipamEmailAddresses, useIpamEmailAddresses)
}

func testAccNetworkIpamThresholdSettings(network, resetValue, triggerValue, useIpamThresholdSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_ipam_threshold_settings" {
    network = %q
    ipam_threshold_settings = {
        reset_value = %q
        trigger_value = %q
    }
	use_ipam_threshold_settings = %q
}
`, network, resetValue, triggerValue, useIpamThresholdSettings)
}

func testAccNetworkIpamTrapSettings(network, enableEmailWarnings, enableSnmpWarnings, useIpamTrapSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_ipam_trap_settings" {
    network = %q
    ipam_trap_settings = {
        enable_email_warnings = %q
        enable_snmp_warnings = %q
    }
	use_ipam_trap_settings = %q
}
`, network, enableEmailWarnings, enableSnmpWarnings, useIpamTrapSettings)
}

// func testAccNetworkIpv4addr(ipv4addr string) string {
// 	return fmt.Sprintf(`
// resource "nios_ipam_network" "test_ipv4addr" {
//     ipv4addr = %q
// }
// `, ipv4addr)
// }

func testAccNetworkLeaseScavengeTime(network, leaseScavengeTime, useLeaseScavengeTime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_lease_scavenge_time" {
    network = %q
    lease_scavenge_time = %q
    use_lease_scavenge_time = %q
}
`, network, leaseScavengeTime, useLeaseScavengeTime)
}

func testAccNetworkLowWaterMark(network, lowWaterMark string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_low_water_mark" {
    network = %q
    low_water_mark = %q
}
`, network, lowWaterMark)
}

func testAccNetworkLowWaterMarkReset(network, lowWaterMarkReset string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_low_water_mark_reset" {
    network = %q
    low_water_mark_reset = %q
}
`, network, lowWaterMarkReset)
}

func testAccNetworkMgmPrivate(network, mgmPrivate, useMgmPrivate string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_mgm_private" {
    network = %q
    mgm_private = %q
    use_mgm_private = %q
}
`, network, mgmPrivate, useMgmPrivate)
}

func testAccNetworkNetmask(network string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_netmask" {
    network = %q
}
`, network)
}

func testAccNetworkNetwork(network string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_network" {
    network = %q
}
`, network)
}

func testAccNetworkNetworkView(network, networkView string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_network_view" {
    network = %q
    network_view = %q
}
`, network, networkView)
}

func testAccNetworkNextserver(network, nextserver, useNextserver string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_nextserver" {
    network = %q
    nextserver = %q
    use_nextserver = %q
}
`, network, nextserver, useNextserver)
}

func testAccNetworkOptions(network, name, num, value, vendorClass, useOption, useOptions string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_options" {
    network = %q
    options = [
		{
			name = %q
			num = %q
			value = %q
			vendor_class = %q
			use_option = %s
		}
    ]
    use_options = %q
}
`, network, name, num, value, vendorClass, useOption, useOptions)
}

func testAccNetworkPortControlBlackoutSetting(network, enableBlackout, useBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_port_control_blackout_setting" {
    network = %q
    port_control_blackout_setting = {
		enable_blackout = %q
	}
	use_blackout_setting = %q
}
`, network, enableBlackout, useBlackoutSetting)
}

func testAccNetworkPxeLeaseTime(network, pxeLeaseTime, usePxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_pxe_lease_time" {
    network = %q
    pxe_lease_time = %q
    use_pxe_lease_time = %q
}
`, network, pxeLeaseTime, usePxeLeaseTime)
}

func testAccNetworkRecycleLeases(network, recycleLeases, useRecycleLeases string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_recycle_leases" {
    network = %q
    recycle_leases = %q
    use_recycle_leases = %q
}
`, network, recycleLeases, useRecycleLeases)
}

func testAccNetworkRirRegistrationStatus(network, rirRegistrationStatus string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_rir_registration_status" {
    network = %q
    rir_registration_status = %q
}
`, network, rirRegistrationStatus)
}

func testAccNetworkSamePortControlDiscoveryBlackout(network, samePortControlDiscoveryBlackout, useBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_same_port_control_discovery_blackout" {
    network = %q
    same_port_control_discovery_blackout = %q
    use_blackout_setting = %q
}
`, network, samePortControlDiscoveryBlackout, useBlackoutSetting)
}

func testAccNetworkUnmanaged(network, unmanaged string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_unmanaged" {
    network = %q
    unmanaged = %q
}
`, network, unmanaged)
}

func testAccNetworkUpdateDnsOnLeaseRenewal(network, updateDnsOnLeaseRenewal, useUpdateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_update_dns_on_lease_renewal" {
    network = %q
    update_dns_on_lease_renewal = %q
    use_update_dns_on_lease_renewal = %q
}
`, network, updateDnsOnLeaseRenewal, useUpdateDnsOnLeaseRenewal)
}

func testAccNetworkUseAuthority(network, useAuthority string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_authority" {
	network = %q
    use_authority = %q
}
`, network, useAuthority)
}

func testAccNetworkUseBlackoutSetting(network, useBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_blackout_setting" {
    network = %q
    use_blackout_setting = %q
}
`, network, useBlackoutSetting)
}

func testAccNetworkUseBootfile(network, useBootfile string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_bootfile" {
    network = %q
    use_bootfile = %q
}
`, network, useBootfile)
}

func testAccNetworkUseBootserver(network, useBootserver string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_bootserver" {
    network = %q
    use_bootserver = %q
}
`, network, useBootserver)
}

func testAccNetworkUseDdnsDomainname(network, useDdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_ddns_domainname" {
	network = %q
    use_ddns_domainname = %q
}
`, network, useDdnsDomainname)
}

func testAccNetworkUseDdnsGenerateHostname(network, useDdnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_ddns_generate_hostname" {
    network = %q
    use_ddns_generate_hostname = %q
}
`, network, useDdnsGenerateHostname)
}

func testAccNetworkUseDdnsTtl(network, useDdnsTtl string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_ddns_ttl" {
	network = %q
    use_ddns_ttl = %q
}
`, network, useDdnsTtl)
}

func testAccNetworkUseDdnsUpdateFixedAddresses(network, useDdnsUpdateFixedAddresses string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_ddns_update_fixed_addresses" {
    network = %q
    use_ddns_update_fixed_addresses = %q
}
`, network, useDdnsUpdateFixedAddresses)
}

func testAccNetworkUseDdnsUseOption81(network, useDdnsUseOption81 string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_ddns_use_option81" {
    network = %q
    use_ddns_use_option81 = %q
}
`, network, useDdnsUseOption81)
}

func testAccNetworkUseDenyBootp(network, useDenyBootp string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_deny_bootp" {
    network = %q
    use_deny_bootp = %q
}
`, network, useDenyBootp)
}

func testAccNetworkUseDiscoveryBasicPollingSettings(network, useDiscoveryBasicPollingSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_discovery_basic_polling_settings" {
    network = %q
    use_discovery_basic_polling_settings = %q
}
`, network, useDiscoveryBasicPollingSettings)
}

func testAccNetworkUseEmailList(network, useEmailList string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_email_list" {
    network = %q
    use_email_list = %q
}
`, network, useEmailList)
}

func testAccNetworkUseEnableDdns(network, useEnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_enable_ddns" {
    network = %q
    use_enable_ddns = %q
}
`, network, useEnableDdns)
}

func testAccNetworkUseEnableDhcpThresholds(network, useEnableDhcpThresholds string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_enable_dhcp_thresholds" {
    network = %q
    use_enable_dhcp_thresholds = %q
}
`, network, useEnableDhcpThresholds)
}

func testAccNetworkUseEnableDiscovery(network, useEnableDiscovery string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_enable_discovery" {
    network = %q
    use_enable_discovery = %q
}
`, network, useEnableDiscovery)
}

func testAccNetworkUseEnableIfmapPublishing(network, useEnableIfmapPublishing string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_enable_ifmap_publishing" {
    network = %q
    use_enable_ifmap_publishing = %q
}
`, network, useEnableIfmapPublishing)
}

func testAccNetworkUseIgnoreDhcpOptionListRequest(network, useIgnoreDhcpOptionListRequest string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_ignore_dhcp_option_list_request" {
    network = %q
    use_ignore_dhcp_option_list_request = %q
}
`, network, useIgnoreDhcpOptionListRequest)
}

func testAccNetworkUseIgnoreId(network, useIgnoreId string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_ignore_id" {
    network = %q
    use_ignore_id = %q
}
`, network, useIgnoreId)
}

func testAccNetworkUseIpamEmailAddresses(network, useIpamEmailAddresses string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_ipam_email_addresses" {
    network = %q
    use_ipam_email_addresses = %q
}
`, network, useIpamEmailAddresses)
}

func testAccNetworkUseIpamThresholdSettings(network, useIpamThresholdSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_ipam_threshold_settings" {
    network = %q
    use_ipam_threshold_settings = %q
}
`, network, useIpamThresholdSettings)
}

func testAccNetworkUseIpamTrapSettings(network, useIpamTrapSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_ipam_trap_settings" {
    network = %q
    use_ipam_trap_settings = %q
}
`, network, useIpamTrapSettings)
}

func testAccNetworkUseLeaseScavengeTime(network, useLeaseScavengeTime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_lease_scavenge_time" {
    network = %q
    use_lease_scavenge_time = %q
}
`, network, useLeaseScavengeTime)
}

func testAccNetworkUseLogicFilterRules(network, useLogicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_logic_filter_rules" {
    network = %q
    use_logic_filter_rules = %q
}
`, network, useLogicFilterRules)
}

func testAccNetworkUseMgmPrivate(network, useMgmPrivate string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_mgm_private" {
    network = %q
    use_mgm_private = %q
}
`, network, useMgmPrivate)
}

func testAccNetworkUseNextserver(network, useNextserver string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_nextserver" {
    network = %q
    use_nextserver = %q
}
`, network, useNextserver)
}

func testAccNetworkUseOptions(network, useOptions string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_options" {
    network = %q
    use_options = %q
}
`, network, useOptions)
}

func testAccNetworkUsePxeLeaseTime(network, usePxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_pxe_lease_time" {
    network = %q
    use_pxe_lease_time = %q
}
`, network, usePxeLeaseTime)
}

func testAccNetworkUseRecycleLeases(network, useRecycleLeases string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_recycle_leases" {
    network = %q
    use_recycle_leases = %q
}
`, network, useRecycleLeases)
}

func testAccNetworkUseSubscribeSettings(network, useSubscribeSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_subscribe_settings" {
    network = %q
    use_subscribe_settings = %q
}
`, network, useSubscribeSettings)
}

func testAccNetworkUseUpdateDnsOnLeaseRenewal(network, useUpdateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_update_dns_on_lease_renewal" {
    network = %q
    use_update_dns_on_lease_renewal = %q
}
`, network, useUpdateDnsOnLeaseRenewal)
}

func testAccNetworkUseZoneAssociations(network, useZoneAssociations string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_use_zone_associations" {
    network = %q
    use_zone_associations = %q
}
`, network, useZoneAssociations)
}
