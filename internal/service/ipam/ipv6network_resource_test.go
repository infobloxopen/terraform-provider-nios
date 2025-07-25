package ipam_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForIpv6network = "cloud_info,comment,ddns_domainname,ddns_enable_option_fqdn,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,disable,discover_now_status,discovered_bgp_as,discovered_bridge_domain,discovered_tenant,discovered_vlan_id,discovered_vlan_name,discovered_vrf_description,discovered_vrf_name,discovered_vrf_rd,discovery_basic_poll_settings,discovery_blackout_setting,discovery_engine_type,discovery_member,domain_name,domain_name_servers,enable_ddns,enable_discovery,enable_ifmap_publishing,endpoint_sources,extattrs,last_rir_registration_update_sent,last_rir_registration_update_status,logic_filter_rules,members,mgm_private,mgm_private_overridable,ms_ad_user_data,network,network_container,network_view,options,port_control_blackout_setting,preferred_lifetime,recycle_leases,rir,rir_organization,rir_registration_status,same_port_control_discovery_blackout,subscribe_settings,unmanaged,unmanaged_count,update_dns_on_lease_renewal,use_blackout_setting,use_ddns_domainname,use_ddns_enable_option_fqdn,use_ddns_generate_hostname,use_ddns_ttl,use_discovery_basic_polling_settings,use_domain_name,use_domain_name_servers,use_enable_ddns,use_enable_discovery,use_enable_ifmap_publishing,use_logic_filter_rules,use_mgm_private,use_options,use_preferred_lifetime,use_recycle_leases,use_subscribe_settings,use_update_dns_on_lease_renewal,use_valid_lifetime,use_zone_associations,valid_lifetime,vlans,zone_associations"

func TestAccIpv6networkResource_basic(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkBasicConfig(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.authority_type", "GM"),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.delegated_scope", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.mgmt_platform", ""),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.owned_by_adaptor", "false"),
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", ""),
					resource.TestCheckResourceAttr(resourceName, "ddns_enable_option_fqdn", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "true"),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "0"),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_discovery", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_ifmap_publishing", "false"),
					resource.TestCheckResourceAttr(resourceName, "mgm_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "mgm_private_overridable", "true"),
					resource.TestCheckResourceAttr(resourceName, "network_view", "default"),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "27000"),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "true"),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_status", "NOT_REGISTERED"),
					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "false"),
					resource.TestCheckResourceAttr(resourceName, "unmanaged", "false"),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_enable_option_fqdn", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ifmap_publishing", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_mgm_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_subscribe_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_zone_associations", "true"),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "43200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_disappears(t *testing.T) {
	resourceName := "nios_ipam_ipv6network.test"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6networkDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6networkBasicConfig(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					testAccCheckIpv6networkDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccIpv6networkResource_CloudInfo(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_cloud_info"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkCloudInfo(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.authority_type", "GM"),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.delegated_scope", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.mgmt_platform", ""),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.owned_by_adaptor", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_Comment(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_comment"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkComment(network, "test comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkComment(network, "updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "updated comment"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_DdnsDomainname(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_ddns_domainname"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkDdnsDomainname(network, "test.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "test.com"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkDdnsDomainname(network, "testupdated.com", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "testupdated.com"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_DdnsEnableOptionFqdn(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_ddns_enable_option_fqdn"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkDdnsEnableOptionFqdn(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_enable_option_fqdn", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_enable_option_fqdn", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkDdnsEnableOptionFqdn(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_enable_option_fqdn", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_enable_option_fqdn", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_DdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_ddns_generate_hostname"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkDdnsGenerateHostname(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkDdnsGenerateHostname(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_DdnsServerAlwaysUpdates(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_ddns_server_always_updates"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkDdnsServerAlwaysUpdates(network, "true", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "ddns_enable_option_fqdn", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_enable_option_fqdn", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkDdnsServerAlwaysUpdates(network, "false", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "ddns_enable_option_fqdn", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_enable_option_fqdn", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_DdnsTtl(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_ddns_ttl"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkDdnsTtl(network, "1", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "1"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkDdnsTtl(network, "2", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "2"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_Disable(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_disable"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkDisable(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkDisable(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_DiscoveredBridgeDomain(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_discovered_bridge_domain"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkDiscoveredBridgeDomain(network, "test_bridge_domain.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovered_bridge_domain", "test_bridge_domain.com"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkDiscoveredBridgeDomain(network, "updated_bridge_domain.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovered_bridge_domain", "updated_bridge_domain.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_DiscoveredTenant(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_discovered_tenant"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkDiscoveredTenant(network, "test_tenant.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovered_tenant", "test_tenant.com"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkDiscoveredTenant(network, "updated_tenant.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovered_tenant", "updated_tenant.com"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_DiscoveryBasicPollSettings(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_discovery_basic_poll_settings"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkDiscoveryBasicPollSettings(network, "true", "true", "false", "default", "false", "false", "1", "false", "false", "true", "PERIODIC", "3600", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
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
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkDiscoveryBasicPollSettings(network, "false", "true", "false", "default", "false", "false", "1", "false", "false", "true", "PERIODIC", "3600", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
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
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_DiscoveryBlackoutSetting(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_discovery_blackout_setting"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkDiscoveryBlackoutSetting(network, "false", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_blackout_setting.enable_blackout", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_DomainName(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_domain_name"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkDomainName(network, "test_domain.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name", "test_domain.com"),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkDomainName(network, "updated_domain.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name", "updated_domain.com"),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_DomainNameServers(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_domain_name_servers"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkDomainNameServers(network, "11::22:33:44:55:66", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.0", "11::22:33:44:55:66"),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkDomainNameServers(network, "11::22:33:44:55:67", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.0", "11::22:33:44:55:67"),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_EnableDdns(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_enable_ddns"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkEnableDdns(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkEnableDdns(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_EnableIfmapPublishing(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_enable_ifmap_publishing"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkEnableIfmapPublishing(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ifmap_publishing", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ifmap_publishing", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkEnableIfmapPublishing(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ifmap_publishing", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ifmap_publishing", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_extattrs"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkExtAttrs(network, map[string]string{"Site": extAttrValue1}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkExtAttrs(network, map[string]string{"Site": extAttrValue2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_MgmPrivate(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_mgm_private"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkMgmPrivate(network, "false", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mgm_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_mgm_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkMgmPrivate(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mgm_private", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_mgm_private", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_NetworkView(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_network_view"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkNetworkView(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "default"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_Options(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_options"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkOptions(network, "dhcp6.fqdn", "39", "test_options.com", "DHCPv6", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "dhcp6.fqdn"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "test_options.com"),
					resource.TestCheckResourceAttr(resourceName, "options.0.vendor_class", "DHCPv6"),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkOptions(network, "dhcp6.fqdn", "39", "updated_options.com", "DHCPv6", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "dhcp6.fqdn"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "updated_options.com"),
					resource.TestCheckResourceAttr(resourceName, "options.0.vendor_class", "DHCPv6"),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_PortControlBlackoutSetting(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_port_control_blackout_setting"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkPortControlBlackoutSetting(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port_control_blackout_setting.enable_blackout", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_PreferredLifetime(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_preferred_lifetime"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkPreferredLifetime(network, "28000", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "28000"),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkPreferredLifetime(network, "30000", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "30000"),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_RecycleLeases(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_recycle_leases"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkRecycleLeases(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkRecycleLeases(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_RirRegistrationStatus(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_rir_registration_status"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkRirRegistrationStatus(network, "NOT_REGISTERED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_status", "NOT_REGISTERED"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_SamePortControlDiscoveryBlackout(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_same_port_control_discovery_blackout"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkSamePortControlDiscoveryBlackout(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkSamePortControlDiscoveryBlackout(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_Unmanaged(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_unmanaged"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkUnmanaged(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "unmanaged", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_UpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_update_dns_on_lease_renewal"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkUpdateDnsOnLeaseRenewal(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkUpdateDnsOnLeaseRenewal(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_UseBlackoutSetting(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_use_blackout_setting"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkUseBlackoutSetting(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkUseBlackoutSetting(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_UseDdnsDomainname(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_use_ddns_domainname"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkUseDdnsDomainname(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkUseDdnsDomainname(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_UseDdnsEnableOptionFqdn(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_use_ddns_enable_option_fqdn"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkUseDdnsEnableOptionFqdn(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_enable_option_fqdn", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkUseDdnsEnableOptionFqdn(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_enable_option_fqdn", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_UseDdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_use_ddns_generate_hostname"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkUseDdnsGenerateHostname(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkUseDdnsGenerateHostname(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_UseDdnsTtl(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_use_ddns_ttl"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkUseDdnsTtl(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkUseDdnsTtl(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_UseDiscoveryBasicPollingSettings(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_use_discovery_basic_polling_settings"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkUseDiscoveryBasicPollingSettings(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkUseDiscoveryBasicPollingSettings(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_UseDomainName(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_use_domain_name"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkUseDomainName(network, "test_domain.com", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "domain_name", "test_domain.com"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkUseDomainName(network, "updated_domain.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "domain_name", "updated_domain.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_UseDomainNameServers(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_use_domain_name_servers"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkUseDomainNameServers(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkUseDomainNameServers(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_UseEnableDdns(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_use_enable_ddns"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkUseEnableDdns(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkUseEnableDdns(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_UseEnableDiscovery(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_use_enable_discovery"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkUseEnableDiscovery(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkUseEnableDiscovery(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_UseEnableIfmapPublishing(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_use_enable_ifmap_publishing"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkUseEnableIfmapPublishing(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ifmap_publishing", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkUseEnableIfmapPublishing(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ifmap_publishing", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_UseLogicFilterRules(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_use_logic_filter_rules"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkUseLogicFilterRules(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkUseLogicFilterRules(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_UseMgmPrivate(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_use_mgm_private"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkUseMgmPrivate(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_mgm_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_UseOptions(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_use_options"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkUseOptions(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkUseOptions(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_UsePreferredLifetime(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_use_preferred_lifetime"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkUsePreferredLifetime(network, "28000", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "28000"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkUsePreferredLifetime(network, "30000", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "30000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_UseRecycleLeases(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_use_recycle_leases"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkUseRecycleLeases(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkUseRecycleLeases(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_UseUpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_use_update_dns_on_lease_renewal"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkUseUpdateDnsOnLeaseRenewal(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkUseUpdateDnsOnLeaseRenewal(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_UseValidLifetime(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_use_valid_lifetime"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkUseValidLifetime(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkUseValidLifetime(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_UseZoneAssociations(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_use_zone_associations"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkUseZoneAssociations(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_zone_associations", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkResource_ValidLifetime(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network.test_valid_lifetime"
	var v ipam.Ipv6network
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkValidLifetime(network, "43200", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "43200"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkValidLifetime(network, "40000", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "40000"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckIpv6networkExists(ctx context.Context, resourceName string, v *ipam.Ipv6network) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.IPAMAPI.
			Ipv6networkAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForIpv6network).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetIpv6networkResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetIpv6networkResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckIpv6networkDestroy(ctx context.Context, v *ipam.Ipv6network) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.IPAMAPI.
			Ipv6networkAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForIpv6network).
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

func testAccCheckIpv6networkDisappears(ctx context.Context, v *ipam.Ipv6network) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.IPAMAPI.
			Ipv6networkAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccIpv6networkBasicConfig(network string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test" {
    network = %q
}
`, network)
}

func testAccIpv6networkCloudInfo(network string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_cloud_info" {
    network = %q
}
`, network)
}

func testAccIpv6networkComment(network, comment string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_comment" {
	network = %q
    comment = %q
}
`, network, comment)
}

func testAccIpv6networkDdnsDomainname(network, ddnsDomainname, useDdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_ddns_domainname" {
    ddns_domainname = %q
    use_ddns_domainname = %q
    network = %q
}
`, ddnsDomainname, useDdnsDomainname, network)
}

func testAccIpv6networkDdnsEnableOptionFqdn(network, ddnsEnableOptionFqdn, useDdnsEnableOptionFqdn string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_ddns_enable_option_fqdn" {
    ddns_enable_option_fqdn = %q
    use_ddns_enable_option_fqdn = %q
    network = %q
}
`, ddnsEnableOptionFqdn, useDdnsEnableOptionFqdn, network)
}

func testAccIpv6networkDdnsGenerateHostname(network, ddnsGenerateHostname, useDdnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_ddns_generate_hostname" {
    ddns_generate_hostname = %q
    use_ddns_generate_hostname = %q
    network = %q
}
`, ddnsGenerateHostname, useDdnsGenerateHostname, network)
}

func testAccIpv6networkDdnsServerAlwaysUpdates(network, ddnsServerAlwaysUpdates, ddnsEnableOptionFqdn, useDdnsEnableOptionFqdn string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_ddns_server_always_updates" {
    ddns_server_always_updates = %q
    network = %q
    ddns_enable_option_fqdn = %q
    use_ddns_enable_option_fqdn = %q
}
`, ddnsServerAlwaysUpdates, network, ddnsEnableOptionFqdn, useDdnsEnableOptionFqdn)
}

func testAccIpv6networkDdnsTtl(network, ddnsTtl, useDdnsTtl string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_ddns_ttl" {
    ddns_ttl = %q
    use_ddns_ttl = %q
    network = %q
}
`, ddnsTtl, useDdnsTtl, network)
}

func testAccIpv6networkDisable(network, disable string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_disable" {
    disable = %q
    network = %q
}
`, disable, network)
}

func testAccIpv6networkDiscoveredBridgeDomain(network, discoveredBridgeDomain string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_discovered_bridge_domain" {
    network = %q
    discovered_bridge_domain = %q
}
`, network, discoveredBridgeDomain)
}

func testAccIpv6networkDiscoveredTenant(network, discoveredTenant string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_discovered_tenant" {
    network = %q
    discovered_tenant = %q
}
`, network, discoveredTenant)
}

func testAccIpv6networkDiscoveryBasicPollSettings(network, autoArpRefreshBeforeSwitchPortPolling, cliCollection, completePingSweep, credentialGroup, deviceProfile, netbiosScanning, pollingFrequencyModifier, portScanning, smartSubnetPingSweep, snmpCollection, switchPortDataCollectionPolling, switchPortDataCollectionPollingInterval, useGlobalPollingFrequencyModifier, useDiscoveryBasicPollSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_discovery_basic_poll_settings" {
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

func testAccIpv6networkDiscoveryBlackoutSetting(network, enableBlackout, useBlackout string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_discovery_blackout_setting" {
    network = %q
    discovery_blackout_setting = {
        enable_blackout = %q
    }
    use_blackout_setting = %q
}
`, network, enableBlackout, useBlackout)
}

func testAccIpv6networkDomainName(network, domainName, useDomainName string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_domain_name" {
    domain_name = %q
    use_domain_name = %q
    network = %q
}
`, domainName, useDomainName, network)
}

func testAccIpv6networkDomainNameServers(network, domainNameServers, useDomainNameServers string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_domain_name_servers" {
    domain_name_servers = [%q]
    use_domain_name_servers = %q
    network = %q
}
`, domainNameServers, useDomainNameServers, network)
}

func testAccIpv6networkEnableDdns(network, enableDdns, useEnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_enable_ddns" {
    enable_ddns = %q
    use_enable_ddns = %q
    network = %q
}
`, enableDdns, useEnableDdns, network)
}

func testAccIpv6networkEnableIfmapPublishing(network, enableIfmapPublishing, useEnableIfmapPublishing string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_enable_ifmap_publishing" {
    enable_ifmap_publishing = %q
    use_enable_ifmap_publishing = %q
    network = %q
}
`, enableIfmapPublishing, useEnableIfmapPublishing, network)
}

func testAccIpv6networkExtAttrs(network string, extAttrs map[string]string) string {
	extAttrsStr := ""
	for key, value := range extAttrs {
		if extAttrsStr != "" {
			extAttrsStr += "\n    "
		}
		extAttrsStr += fmt.Sprintf("%s = %q", key, value)
	}

	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_extattrs" {
    network = %q
    extattrs = {
        %s
    }
}
`, network, extAttrsStr)
}

func testAccIpv6networkMgmPrivate(network, mgmPrivate, useMgmPrivate string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_mgm_private" {
    network = %q
    mgm_private = %q
    use_mgm_private = %q
}
`, network, mgmPrivate, useMgmPrivate)
}

func testAccIpv6networkNetworkView(network string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_network_view" {
    network = %q
}
`, network)
}

func testAccIpv6networkOptions(network, name, num, value, vendorClass, useOptions string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_options" {
    network = %q
    options = [
		{
			name = %q
			num = %q
			value = %q
			vendor_class = %q
		}
    ]
    use_options = %q
}
`, network, name, num, value, vendorClass, useOptions)
}

func testAccIpv6networkPortControlBlackoutSetting(network, enableBlackout, useBlackout string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_port_control_blackout_setting" {
    port_control_blackout_setting = {
        enable_blackout = %q
    }
    use_blackout_setting = %q
    network = %q
}
`, enableBlackout, useBlackout, network)
}

func testAccIpv6networkPreferredLifetime(network, preferredLifetime, usePreferredLifetime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_preferred_lifetime" {
    network = %q
    preferred_lifetime = %q
    use_preferred_lifetime = %q
}
`, network, preferredLifetime, usePreferredLifetime)
}

func testAccIpv6networkRecycleLeases(network, recycleLeases, useRecycleLeases string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_recycle_leases" {
    network = %q
    recycle_leases = %q
    use_recycle_leases = %q
}
`, network, recycleLeases, useRecycleLeases)
}

func testAccIpv6networkRirRegistrationStatus(network, rirRegistrationStatus string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_rir_registration_status" {
    network = %q
    rir_registration_status = %q
}
`, network, rirRegistrationStatus)
}

func testAccIpv6networkSamePortControlDiscoveryBlackout(network, samePortControlDiscoveryBlackout, useBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_same_port_control_discovery_blackout" {
    network = %q
    same_port_control_discovery_blackout = %q
    use_blackout_setting = %q
}
`, network, samePortControlDiscoveryBlackout, useBlackoutSetting)
}

func testAccIpv6networkUnmanaged(network, unmanaged string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_unmanaged" {
    network = %q
    unmanaged = %q
}
`, network, unmanaged)
}

func testAccIpv6networkUpdateDnsOnLeaseRenewal(network, updateDnsOnLeaseRenewal, useUpdateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_update_dns_on_lease_renewal" {
    network = %q
    update_dns_on_lease_renewal = %q
    use_update_dns_on_lease_renewal = %q
}
`, network, updateDnsOnLeaseRenewal, useUpdateDnsOnLeaseRenewal)
}

func testAccIpv6networkUseBlackoutSetting(network, useBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_use_blackout_setting" {
    network = %q
    use_blackout_setting = %q
}
`, network, useBlackoutSetting)
}

func testAccIpv6networkUseDdnsDomainname(network, useDdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_use_ddns_domainname" {
    network = %q
    use_ddns_domainname = %q
}
`, network, useDdnsDomainname)
}

func testAccIpv6networkUseDdnsEnableOptionFqdn(network, useDdnsEnableOptionFqdn string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_use_ddns_enable_option_fqdn" {
    network = %q
    use_ddns_enable_option_fqdn = %q
}
`, network, useDdnsEnableOptionFqdn)
}

func testAccIpv6networkUseDdnsGenerateHostname(network, useDdnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_use_ddns_generate_hostname" {
    network = %q
    use_ddns_generate_hostname = %q
}
`, network, useDdnsGenerateHostname)
}

func testAccIpv6networkUseDdnsTtl(network, useDdnsTtl string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_use_ddns_ttl" {
    network = %q
    use_ddns_ttl = %q
}
`, network, useDdnsTtl)
}

func testAccIpv6networkUseDiscoveryBasicPollingSettings(network, useDiscoveryBasicPollingSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_use_discovery_basic_polling_settings" {
    network = %q
    use_discovery_basic_polling_settings = %q
}
`, network, useDiscoveryBasicPollingSettings)
}

func testAccIpv6networkUseDomainName(network, domainName, useDomainName string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_use_domain_name" {
    network = %q
    domain_name = %q
    use_domain_name = %q
}
`, network, domainName, useDomainName)
}

func testAccIpv6networkUseDomainNameServers(network, useDomainNameServers string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_use_domain_name_servers" {
    network = %q
    use_domain_name_servers = %q
}
`, network, useDomainNameServers)
}

func testAccIpv6networkUseEnableDdns(network, useEnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_use_enable_ddns" {
    network = %q
    use_enable_ddns = %q
}
`, network, useEnableDdns)
}

func testAccIpv6networkUseEnableDiscovery(network, useEnableDiscovery string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_use_enable_discovery" {
    network = %q
    use_enable_discovery = %q
}
`, network, useEnableDiscovery)
}

func testAccIpv6networkUseEnableIfmapPublishing(network, useEnableIfmapPublishing string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_use_enable_ifmap_publishing" {
    network = %q
    use_enable_ifmap_publishing = %q
}
`, network, useEnableIfmapPublishing)
}

func testAccIpv6networkUseLogicFilterRules(network, useLogicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_use_logic_filter_rules" {
    network = %q
    use_logic_filter_rules = %q
}
`, network, useLogicFilterRules)
}

func testAccIpv6networkUseMgmPrivate(network, useMgmPrivate string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_use_mgm_private" {
    network = %q
    use_mgm_private = %q
}
`, network, useMgmPrivate)
}

func testAccIpv6networkUseOptions(network, useOptions string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_use_options" {
    network = %q
    use_options = %q
}
`, network, useOptions)
}

func testAccIpv6networkUsePreferredLifetime(network, preferredLifetime, usePreferredLifetime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_use_preferred_lifetime" {
    network = %q
    preferred_lifetime = %q
    use_preferred_lifetime = %q
}
`, network, preferredLifetime, usePreferredLifetime)
}

func testAccIpv6networkUseRecycleLeases(network, useRecycleLeases string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_use_recycle_leases" {
    network = %q
    use_recycle_leases = %q
}
`, network, useRecycleLeases)
}

func testAccIpv6networkUseUpdateDnsOnLeaseRenewal(network, useUpdateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_use_update_dns_on_lease_renewal" {
    network = %q
    use_update_dns_on_lease_renewal = %q
}
`, network, useUpdateDnsOnLeaseRenewal)
}

func testAccIpv6networkUseValidLifetime(network, useValidLifetime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_use_valid_lifetime" {
    network = %q
    use_valid_lifetime = %q
}
`, network, useValidLifetime)
}

func testAccIpv6networkUseZoneAssociations(network, useZoneAssociations string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_use_zone_associations" {
    network = %q
    use_zone_associations = %q
}
`, network, useZoneAssociations)
}

func testAccIpv6networkValidLifetime(network, validLifetime, useValidLifetime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network" "test_valid_lifetime" {
    network = %q
    valid_lifetime = %q
    use_valid_lifetime = %q
}
`, network, validLifetime, useValidLifetime)
}
