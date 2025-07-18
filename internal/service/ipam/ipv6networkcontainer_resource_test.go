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

var readableAttributesForIpv6networkcontainer = "cloud_info,comment,ddns_domainname,ddns_enable_option_fqdn,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,discover_now_status,discovery_basic_poll_settings,discovery_blackout_setting,discovery_engine_type,discovery_member,domain_name_servers,enable_ddns,enable_discovery,endpoint_sources,extattrs,last_rir_registration_update_sent,last_rir_registration_update_status,logic_filter_rules,mgm_private,mgm_private_overridable,ms_ad_user_data,network,network_container,network_view,options,port_control_blackout_setting,preferred_lifetime,rir,rir_organization,rir_registration_status,same_port_control_discovery_blackout,subscribe_settings,unmanaged,update_dns_on_lease_renewal,use_blackout_setting,use_ddns_domainname,use_ddns_enable_option_fqdn,use_ddns_generate_hostname,use_ddns_ttl,use_discovery_basic_polling_settings,use_domain_name_servers,use_enable_ddns,use_enable_discovery,use_logic_filter_rules,use_mgm_private,use_options,use_preferred_lifetime,use_subscribe_settings,use_update_dns_on_lease_renewal,use_valid_lifetime,use_zone_associations,utilization,valid_lifetime,zone_associations"

func TestAccIpv6networkcontainerResource_basic(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerBasicConfig(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					// Check default values are populated correctly
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "true"),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "0"),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_discovery", "false"),
					resource.TestCheckResourceAttr(resourceName, "mgm_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "mgm_private_overridable", "true"),
					resource.TestCheckResourceAttr(resourceName, "network_view", "default"),
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
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_mgm_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_subscribe_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_zone_associations", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_disappears(t *testing.T) {
	resourceName := "nios_ipam_ipv6network_container.test"
	var v ipam.Ipv6networkcontainer
	// Generate a random IPv6 network for the test
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6networkcontainerDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6networkcontainerBasicConfig(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					testAccCheckIpv6networkcontainerDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccIpv6networkcontainerResource_CloudInfo(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_cloud_info"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerCloudInfo(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.authority_type", "GM"),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.delegated_scope", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.mgmt_platform", ""),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.owned_by_adaptor", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
		},
	})
}

func TestAccIpv6networkcontainerResource_Comment(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_comment"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerComment(network, "test comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerComment(network, "updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "updated comment"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_DdnsDomainname(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_ddns_domainname"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerDdnsDomainname(network, "test.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "test.com"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerDdnsDomainname(network, "testupdated.com", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "testupdated.com"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_DdnsEnableOptionFqdn(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_ddns_enable_option_fqdn"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerDdnsEnableOptionFqdn(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_enable_option_fqdn", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_enable_option_fqdn", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerDdnsEnableOptionFqdn(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_enable_option_fqdn", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_enable_option_fqdn", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_DdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_ddns_generate_hostname"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerDdnsGenerateHostname(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerDdnsGenerateHostname(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_DdnsServerAlwaysUpdates(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_ddns_server_always_updates"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerDdnsServerAlwaysUpdates(network, "true", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "true"),
					resource.TestCheckResourceAttr(resourceName, "ddns_enable_option_fqdn", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_enable_option_fqdn", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerDdnsServerAlwaysUpdates(network, "false", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_enable_option_fqdn", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_enable_option_fqdn", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_DdnsTtl(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_ddns_ttl"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerDdnsTtl(network, "1", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "1"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerDdnsTtl(network, "2", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "2"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_DiscoveryBasicPollSettings(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_discovery_basic_poll_settings"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerDiscoveryBasicPollSettings(network, "true", "true", "false", "default", "false", "false", "1", "false", "false", "true", "PERIODIC", "3600", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
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
				Config: testAccIpv6networkcontainerDiscoveryBasicPollSettings(network, "false", "true", "false", "default", "false", "false", "1", "false", "false", "true", "PERIODIC", "3600", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
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

func TestAccIpv6networkcontainerResource_DiscoveryBlackoutSetting(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_discovery_blackout_setting"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerDiscoveryBlackoutSetting(network, "false", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_blackout_setting.enable_blackout", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_DomainNameServers(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_domain_name_servers"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerDomainNameServers(network, "100::1", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.0", "100::1"),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerDomainNameServers(network, "100::2", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.0", "100::2"),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_EnableDdns(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_enable_ddns"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerEnableDdns(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerEnableDdns(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_extattrs"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerExtAttrs(network, map[string]string{"Site": extAttrValue1}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerExtAttrs(network, map[string]string{"Site": extAttrValue2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_MgmPrivate(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_mgm_private"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerMgmPrivate(network, "false", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mgm_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_mgm_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerMgmPrivate(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mgm_private", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_mgm_private", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_Network(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_network"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerNetwork(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_NetworkView(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_network_view"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerNetworkView(network, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "default"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_Options(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_options"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerOptions(network, "dhcp6.fqdn", "39", "test.com", "DHCPv6", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "dhcp6.fqdn"),
					resource.TestCheckResourceAttr(resourceName, "options.0.num", "39"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "test.com"),
					resource.TestCheckResourceAttr(resourceName, "options.0.vendor_class", "DHCPv6"),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerOptions(network, "dhcp-rebinding-time", "59", "100", "DHCP", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "dhcp-rebinding-time"),
					resource.TestCheckResourceAttr(resourceName, "options.0.num", "59"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "100"),
					resource.TestCheckResourceAttr(resourceName, "options.0.vendor_class", "DHCP"),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_PortControlBlackoutSetting(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_port_control_blackout_setting"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerPortControlBlackoutSetting(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port_control_blackout_setting.enable_blackout", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_PreferredLifetime(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_preferred_lifetime"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerPreferredLifetime(network, "27000", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "27000"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerPreferredLifetime(network, "30000", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "30000"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_RirRegistrationStatus(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_rir_registration_status"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerRirRegistrationStatus(network, "NOT_REGISTERED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_status", "NOT_REGISTERED"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_SamePortControlDiscoveryBlackout(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_same_port_control_discovery_blackout"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerSamePortControlDiscoveryBlackout(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerSamePortControlDiscoveryBlackout(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_Unmanaged(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_unmanaged"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUnmanaged(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "unmanaged", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_update_dns_on_lease_renewal"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUpdateDnsOnLeaseRenewal(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUpdateDnsOnLeaseRenewal(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseBlackoutSetting(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_blackout_setting"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseBlackoutSetting(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseBlackoutSetting(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseDdnsDomainname(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_ddns_domainname"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseDdnsDomainname(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseDdnsDomainname(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseDdnsEnableOptionFqdn(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_ddns_enable_option_fqdn"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseDdnsEnableOptionFqdn(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_enable_option_fqdn", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseDdnsEnableOptionFqdn(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_enable_option_fqdn", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseDdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_ddns_generate_hostname"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseDdnsGenerateHostname(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseDdnsGenerateHostname(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseDdnsTtl(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_ddns_ttl"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseDdnsTtl(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseDdnsTtl(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseDiscoveryBasicPollingSettings(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_discovery_basic_polling_settings"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseDiscoveryBasicPollingSettings(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseDiscoveryBasicPollingSettings(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseDomainNameServers(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_domain_name_servers"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseDomainNameServers(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseDomainNameServers(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseEnableDdns(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_enable_ddns"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseEnableDdns(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseEnableDdns(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseEnableDiscovery(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_enable_discovery"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseEnableDiscovery(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseEnableDiscovery(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseLogicFilterRules(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_logic_filter_rules"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseLogicFilterRules(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseLogicFilterRules(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseMgmPrivate(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_mgm_private"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseMgmPrivate(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_mgm_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseOptions(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_options"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseOptions(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseOptions(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UsePreferredLifetime(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_preferred_lifetime"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUsePreferredLifetime(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUsePreferredLifetime(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseSubscribeSettings(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_subscribe_settings"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseSubscribeSettings(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_subscribe_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseUpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_update_dns_on_lease_renewal"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseUpdateDnsOnLeaseRenewal(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseUpdateDnsOnLeaseRenewal(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseValidLifetime(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_valid_lifetime"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseValidLifetime(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseValidLifetime(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseZoneAssociations(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_zone_associations"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseZoneAssociations(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_zone_associations", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_ValidLifetime(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_valid_lifetime"
	var v ipam.Ipv6networkcontainer
	network := acctest.RandomIPv6Network()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerValidLifetime(network, "43200", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "43200"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerValidLifetime(network, "50000", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "50000"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckIpv6networkcontainerExists(ctx context.Context, resourceName string, v *ipam.Ipv6networkcontainer) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.IPAMAPI.
			Ipv6networkcontainerAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForIpv6networkcontainer).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetIpv6networkcontainerResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetIpv6networkcontainerResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckIpv6networkcontainerDestroy(ctx context.Context, v *ipam.Ipv6networkcontainer) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.IPAMAPI.
			Ipv6networkcontainerAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForIpv6networkcontainer).
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

func testAccCheckIpv6networkcontainerDisappears(ctx context.Context, v *ipam.Ipv6networkcontainer) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.IPAMAPI.
			Ipv6networkcontainerAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccIpv6networkcontainerBasicConfig(network string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test" {
    network = %q
}
`, network)
}

func testAccIpv6networkcontainerCloudInfo(network string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_cloud_info" {
    network = %q
}
`, network)
}

func testAccIpv6networkcontainerComment(network, comment string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_comment" {
    network = %q
    comment = %q
}
`, network, comment)
}

func testAccIpv6networkcontainerDdnsDomainname(network, ddnsDomainname, useDdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_ddns_domainname" {
    network = %q
    ddns_domainname = %q
    use_ddns_domainname = %q
}
`, network, ddnsDomainname, useDdnsDomainname)
}

func testAccIpv6networkcontainerDdnsEnableOptionFqdn(network, ddnsEnableOptionFqdn, useDdnsEnableOptionFqdn string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_ddns_enable_option_fqdn" {
    network = %q
    ddns_enable_option_fqdn = %q
    use_ddns_enable_option_fqdn = %q
}
`, network, ddnsEnableOptionFqdn, useDdnsEnableOptionFqdn)
}

func testAccIpv6networkcontainerDdnsGenerateHostname(network, ddnsGenerateHostname, useDdnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_ddns_generate_hostname" {
    network = %q
    ddns_generate_hostname = %q
    use_ddns_generate_hostname = %q
}
`, network, ddnsGenerateHostname, useDdnsGenerateHostname)
}

func testAccIpv6networkcontainerDdnsServerAlwaysUpdates(network, ddnsServerAlwaysUpdates, ddnsEnableOptionFqdn, useDdnsEnableOptionFqdn string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_ddns_server_always_updates" {
    network = %q
    ddns_server_always_updates = %q
	ddns_enable_option_fqdn = %q
    use_ddns_enable_option_fqdn = %q
}
`, network, ddnsServerAlwaysUpdates, ddnsEnableOptionFqdn, useDdnsEnableOptionFqdn)
}

func testAccIpv6networkcontainerDdnsTtl(network, ddnsTtl, useDdnsTtl string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_ddns_ttl" {
    network = %q
    ddns_ttl = %q
    use_ddns_ttl = %q
}
`, network, ddnsTtl, useDdnsTtl)
}

func testAccIpv6networkcontainerDiscoveryBasicPollSettings(network, autoArpRefreshBeforeSwitchPortPolling, cliCollection, completePingSweep, credentialGroup, deviceProfile, netbiosScanning, pollingFrequencyModifier, portScanning, smartSubnetPingSweep, snmpCollection, switchPortDataCollectionPolling, switchPortDataCollectionPollingInterval, useGlobalPollingFrequencyModifier, useDiscoveryBasicPollSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_discovery_basic_poll_settings" {
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

func testAccIpv6networkcontainerDiscoveryBlackoutSetting(network, enabledBlackout, useBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_discovery_blackout_setting" {
    network = %q
    discovery_blackout_setting = {
        enabled = %q
    }
    use_blackout_setting = %q
}
`, network, enabledBlackout, useBlackoutSetting)
}

func testAccIpv6networkcontainerDomainNameServers(network, domainNameServers, useDomainNameServers string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_domain_name_servers" {
    network = %q
    domain_name_servers = [%q]
	use_domain_name_servers = %q
}
`, network, domainNameServers, useDomainNameServers)
}

func testAccIpv6networkcontainerEnableDdns(network, enableDdns, useEnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_enable_ddns" {
    network = %q
    enable_ddns = %q
    use_enable_ddns = %q
}
`, network, enableDdns, useEnableDdns)
}

func testAccIpv6networkcontainerExtAttrs(network string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_extattrs" {
    network = %q
    extattrs = %s
}
`, network, extattrsStr)
}

func testAccIpv6networkcontainerMgmPrivate(network, mgmPrivate, useMgmPrivate string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_mgm_private" {
    network = %q
    mgm_private = %q
    use_mgm_private = %q
}
`, network, mgmPrivate, useMgmPrivate)
}

func testAccIpv6networkcontainerNetwork(network string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_network" {
    network = %q
}
`, network)
}

func testAccIpv6networkcontainerNetworkView(network, networkView string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_network_view" {
    network = %q
    network_view = %q
}
`, network, networkView)
}

func testAccIpv6networkcontainerOptions(network, name, num, value, vendorClass, useOptions string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_options" {
    network = %q
    options = [{
        name = %q
        num = %q
        value = %q
        vendor_class = %q
    }]
    use_options = %q
}
`, network, name, num, value, vendorClass, useOptions)
}

func testAccIpv6networkcontainerPortControlBlackoutSetting(network, portControlBlackoutSetting, useBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_port_control_blackout_setting" {
    network = %q
    port_control_blackout_setting = {
        enable_blackout = %q
    }
    use_blackout_setting = %q
}
`, network, portControlBlackoutSetting, useBlackoutSetting)
}

func testAccIpv6networkcontainerPreferredLifetime(network, preferredLifetime, usePreferredLifetime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_preferred_lifetime" {
    network = %q
    preferred_lifetime = %q
    use_preferred_lifetime = %q
}
`, network, preferredLifetime, usePreferredLifetime)
}

func testAccIpv6networkcontainerRirOrganization(rirOrganization string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_rir_organization" {
    rir_organization = %q
}
`, rirOrganization)
}

func testAccIpv6networkcontainerRirRegistrationStatus(network, rirRegistrationStatus string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_rir_registration_status" {
	network = %q
    rir_registration_status = %q
}
`, network, rirRegistrationStatus)
}

func testAccIpv6networkcontainerSamePortControlDiscoveryBlackout(network, samePortControlDiscoveryBlackout, useBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_same_port_control_discovery_blackout" {
    network = %q
    same_port_control_discovery_blackout = %q
    use_blackout_setting = %q
}
`, network, samePortControlDiscoveryBlackout, useBlackoutSetting)
}

func testAccIpv6networkcontainerSubscribeSettings(subscribeSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_subscribe_settings" {
    subscribe_settings = %q
}
`, subscribeSettings)
}

func testAccIpv6networkcontainerUnmanaged(network, unmanaged string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_unmanaged" {
    network = %q
    unmanaged = %q
}
`, network, unmanaged)
}

func testAccIpv6networkcontainerUpdateDnsOnLeaseRenewal(network, updateDnsOnLeaseRenewal, useUpdateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_update_dns_on_lease_renewal" {
    network = %q
    update_dns_on_lease_renewal = %q
    use_update_dns_on_lease_renewal = %q
}
`, network, updateDnsOnLeaseRenewal, useUpdateDnsOnLeaseRenewal)
}

func testAccIpv6networkcontainerUseBlackoutSetting(network, useBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_blackout_setting" {
    network = %q
    use_blackout_setting = %q
}
`, network, useBlackoutSetting)
}

func testAccIpv6networkcontainerUseDdnsDomainname(network, useDdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_ddns_domainname" {
    network = %q
    use_ddns_domainname = %q
}
`, network, useDdnsDomainname)
}

func testAccIpv6networkcontainerUseDdnsEnableOptionFqdn(network, useDdnsEnableOptionFqdn string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_ddns_enable_option_fqdn" {
    network = %q
    use_ddns_enable_option_fqdn = %q
}
`, network, useDdnsEnableOptionFqdn)
}

func testAccIpv6networkcontainerUseDdnsGenerateHostname(network, useDdnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_ddns_generate_hostname" {
    network = %q
    use_ddns_generate_hostname = %q
}
`, network, useDdnsGenerateHostname)
}

func testAccIpv6networkcontainerUseDdnsTtl(network, useDdnsTtl string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_ddns_ttl" {
    network = %q
    use_ddns_ttl = %q
}
`, network, useDdnsTtl)
}

func testAccIpv6networkcontainerUseDiscoveryBasicPollingSettings(network, useDiscoveryBasicPollingSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_discovery_basic_polling_settings" {
    network = %q
    use_discovery_basic_polling_settings = %q
}
`, network, useDiscoveryBasicPollingSettings)
}

func testAccIpv6networkcontainerUseDomainNameServers(network, useDomainNameServers string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_domain_name_servers" {
    network = %q
    use_domain_name_servers = %q
}
`, network, useDomainNameServers)
}

func testAccIpv6networkcontainerUseEnableDdns(network, useEnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_enable_ddns" {
    network = %q
    use_enable_ddns = %q
}
`, network, useEnableDdns)
}

func testAccIpv6networkcontainerUseEnableDiscovery(network, useEnableDiscovery string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_enable_discovery" {
    network = %q
    use_enable_discovery = %q
}
`, network, useEnableDiscovery)
}

func testAccIpv6networkcontainerUseLogicFilterRules(network, useLogicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_logic_filter_rules" {
    network = %q
    use_logic_filter_rules = %q
}
`, network, useLogicFilterRules)
}

func testAccIpv6networkcontainerUseMgmPrivate(network, useMgmPrivate string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_mgm_private" {
    network = %q
    use_mgm_private = %q
}
`, network, useMgmPrivate)
}

func testAccIpv6networkcontainerUseOptions(network, useOptions string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_options" {
    network = %q
    use_options = %q
}
`, network, useOptions)
}

func testAccIpv6networkcontainerUsePreferredLifetime(network, usePreferredLifetime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_preferred_lifetime" {
    network = %q
    use_preferred_lifetime = %q
}
`, network, usePreferredLifetime)
}

func testAccIpv6networkcontainerUseSubscribeSettings(network, useSubscribeSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_subscribe_settings" {
    network = %q
    use_subscribe_settings = %q
}
`, network, useSubscribeSettings)
}

func testAccIpv6networkcontainerUseUpdateDnsOnLeaseRenewal(network, useUpdateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_update_dns_on_lease_renewal" {
    network = %q
    use_update_dns_on_lease_renewal = %q
}
`, network, useUpdateDnsOnLeaseRenewal)
}

func testAccIpv6networkcontainerUseValidLifetime(network, useValidLifetime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_valid_lifetime" {
    network = %q
    use_valid_lifetime = %q
}
`, network, useValidLifetime)
}

func testAccIpv6networkcontainerUseZoneAssociations(network, useZoneAssociations string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_zone_associations" {
    network = %q
    use_zone_associations = %q
}
`, network, useZoneAssociations)
}

func testAccIpv6networkcontainerValidLifetime(network, validLifetime, useValidLifetime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_valid_lifetime" {
    network = %q
    valid_lifetime = %q
    use_valid_lifetime = %q
}
`, network, validLifetime, useValidLifetime)
}
