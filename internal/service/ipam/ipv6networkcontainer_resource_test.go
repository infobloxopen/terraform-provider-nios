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
			// Delete testing automatically occurs in TestCase
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerDdnsDomainname("DDNS_DOMAINNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "DDNS_DOMAINNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerDdnsDomainname("DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_DdnsEnableOptionFqdn(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_ddns_enable_option_fqdn"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerDdnsEnableOptionFqdn("DDNS_ENABLE_OPTION_FQDN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_enable_option_fqdn", "DDNS_ENABLE_OPTION_FQDN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerDdnsEnableOptionFqdn("DDNS_ENABLE_OPTION_FQDN_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_enable_option_fqdn", "DDNS_ENABLE_OPTION_FQDN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_DdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_ddns_generate_hostname"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerDdnsGenerateHostname("DDNS_GENERATE_HOSTNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "DDNS_GENERATE_HOSTNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerDdnsGenerateHostname("DDNS_GENERATE_HOSTNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "DDNS_GENERATE_HOSTNAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_DdnsServerAlwaysUpdates(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_ddns_server_always_updates"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerDdnsServerAlwaysUpdates("DDNS_SERVER_ALWAYS_UPDATES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "DDNS_SERVER_ALWAYS_UPDATES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerDdnsServerAlwaysUpdates("DDNS_SERVER_ALWAYS_UPDATES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "DDNS_SERVER_ALWAYS_UPDATES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_DdnsTtl(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_ddns_ttl"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerDdnsTtl("DDNS_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "DDNS_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerDdnsTtl("DDNS_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "DDNS_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_DiscoveryBasicPollSettings(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_discovery_basic_poll_settings"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerDiscoveryBasicPollSettings("DISCOVERY_BASIC_POLL_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings", "DISCOVERY_BASIC_POLL_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerDiscoveryBasicPollSettings("DISCOVERY_BASIC_POLL_SETTINGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_basic_poll_settings", "DISCOVERY_BASIC_POLL_SETTINGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_DiscoveryBlackoutSetting(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_discovery_blackout_setting"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerDiscoveryBlackoutSetting("DISCOVERY_BLACKOUT_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_blackout_setting", "DISCOVERY_BLACKOUT_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerDiscoveryBlackoutSetting("DISCOVERY_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_blackout_setting", "DISCOVERY_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_DiscoveryMember(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_discovery_member"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerDiscoveryMember("DISCOVERY_MEMBER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_member", "DISCOVERY_MEMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerDiscoveryMember("DISCOVERY_MEMBER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_member", "DISCOVERY_MEMBER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_DomainNameServers(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_domain_name_servers"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerDomainNameServers("DOMAIN_NAME_SERVERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers", "DOMAIN_NAME_SERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerDomainNameServers("DOMAIN_NAME_SERVERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers", "DOMAIN_NAME_SERVERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_EnableDdns(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_enable_ddns"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerEnableDdns("ENABLE_DDNS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "ENABLE_DDNS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerEnableDdns("ENABLE_DDNS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "ENABLE_DDNS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_EnableDiscovery(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_enable_discovery"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerEnableDiscovery("ENABLE_DISCOVERY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_discovery", "ENABLE_DISCOVERY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerEnableDiscovery("ENABLE_DISCOVERY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_discovery", "ENABLE_DISCOVERY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_extattrs"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_FederatedRealms(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_federated_realms"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerFederatedRealms("FEDERATED_REALMS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "federated_realms", "FEDERATED_REALMS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerFederatedRealms("FEDERATED_REALMS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "federated_realms", "FEDERATED_REALMS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_LogicFilterRules(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_logic_filter_rules"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerLogicFilterRules("LOGIC_FILTER_RULES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerLogicFilterRules("LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_MgmPrivate(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_mgm_private"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerMgmPrivate("MGM_PRIVATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mgm_private", "MGM_PRIVATE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerMgmPrivate("MGM_PRIVATE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mgm_private", "MGM_PRIVATE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_MsAdUserData(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_ms_ad_user_data"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerMsAdUserData("MS_AD_USER_DATA_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_user_data", "MS_AD_USER_DATA_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerMsAdUserData("MS_AD_USER_DATA_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_user_data", "MS_AD_USER_DATA_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_Network(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_network"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerNetwork("NETWORK_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", "NETWORK_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerNetwork("NETWORK_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", "NETWORK_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_NetworkView(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_network_view"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerNetworkView("NETWORK_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerNetworkView("NETWORK_VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_Options(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_options"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerOptions("OPTIONS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerOptions("OPTIONS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_PortControlBlackoutSetting(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_port_control_blackout_setting"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerPortControlBlackoutSetting("PORT_CONTROL_BLACKOUT_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port_control_blackout_setting", "PORT_CONTROL_BLACKOUT_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerPortControlBlackoutSetting("PORT_CONTROL_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port_control_blackout_setting", "PORT_CONTROL_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_PreferredLifetime(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_preferred_lifetime"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerPreferredLifetime("PREFERRED_LIFETIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "PREFERRED_LIFETIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerPreferredLifetime("PREFERRED_LIFETIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "PREFERRED_LIFETIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_RirOrganization(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_rir_organization"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerRirOrganization("RIR_ORGANIZATION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_organization", "RIR_ORGANIZATION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerRirOrganization("RIR_ORGANIZATION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_organization", "RIR_ORGANIZATION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_RirRegistrationStatus(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_rir_registration_status"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerRirRegistrationStatus("RIR_REGISTRATION_STATUS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_status", "RIR_REGISTRATION_STATUS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerRirRegistrationStatus("RIR_REGISTRATION_STATUS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_status", "RIR_REGISTRATION_STATUS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_SamePortControlDiscoveryBlackout(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_same_port_control_discovery_blackout"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerSamePortControlDiscoveryBlackout("SAME_PORT_CONTROL_DISCOVERY_BLACKOUT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "SAME_PORT_CONTROL_DISCOVERY_BLACKOUT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerSamePortControlDiscoveryBlackout("SAME_PORT_CONTROL_DISCOVERY_BLACKOUT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "SAME_PORT_CONTROL_DISCOVERY_BLACKOUT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_SubscribeSettings(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_subscribe_settings"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerSubscribeSettings("SUBSCRIBE_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscribe_settings", "SUBSCRIBE_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerSubscribeSettings("SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscribe_settings", "SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_Unmanaged(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_unmanaged"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUnmanaged("UNMANAGED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "unmanaged", "UNMANAGED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUnmanaged("UNMANAGED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "unmanaged", "UNMANAGED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_update_dns_on_lease_renewal"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUpdateDnsOnLeaseRenewal("UPDATE_DNS_ON_LEASE_RENEWAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "UPDATE_DNS_ON_LEASE_RENEWAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUpdateDnsOnLeaseRenewal("UPDATE_DNS_ON_LEASE_RENEWAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "UPDATE_DNS_ON_LEASE_RENEWAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseBlackoutSetting(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_blackout_setting"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseBlackoutSetting("USE_BLACKOUT_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "USE_BLACKOUT_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseBlackoutSetting("USE_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "USE_BLACKOUT_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseDdnsDomainname(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_ddns_domainname"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseDdnsDomainname("USE_DDNS_DOMAINNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "USE_DDNS_DOMAINNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseDdnsDomainname("USE_DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "USE_DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseDdnsEnableOptionFqdn(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_ddns_enable_option_fqdn"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseDdnsEnableOptionFqdn("USE_DDNS_ENABLE_OPTION_FQDN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_enable_option_fqdn", "USE_DDNS_ENABLE_OPTION_FQDN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseDdnsEnableOptionFqdn("USE_DDNS_ENABLE_OPTION_FQDN_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_enable_option_fqdn", "USE_DDNS_ENABLE_OPTION_FQDN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseDdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_ddns_generate_hostname"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseDdnsGenerateHostname("USE_DDNS_GENERATE_HOSTNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "USE_DDNS_GENERATE_HOSTNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseDdnsGenerateHostname("USE_DDNS_GENERATE_HOSTNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "USE_DDNS_GENERATE_HOSTNAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseDdnsTtl(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_ddns_ttl"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseDdnsTtl("USE_DDNS_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "USE_DDNS_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseDdnsTtl("USE_DDNS_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "USE_DDNS_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseDiscoveryBasicPollingSettings(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_discovery_basic_polling_settings"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseDiscoveryBasicPollingSettings("USE_DISCOVERY_BASIC_POLLING_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "USE_DISCOVERY_BASIC_POLLING_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseDiscoveryBasicPollingSettings("USE_DISCOVERY_BASIC_POLLING_SETTINGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "USE_DISCOVERY_BASIC_POLLING_SETTINGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseDomainNameServers(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_domain_name_servers"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseDomainNameServers("USE_DOMAIN_NAME_SERVERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "USE_DOMAIN_NAME_SERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseDomainNameServers("USE_DOMAIN_NAME_SERVERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "USE_DOMAIN_NAME_SERVERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseEnableDdns(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_enable_ddns"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseEnableDdns("USE_ENABLE_DDNS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "USE_ENABLE_DDNS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseEnableDdns("USE_ENABLE_DDNS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "USE_ENABLE_DDNS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseEnableDiscovery(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_enable_discovery"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseEnableDiscovery("USE_ENABLE_DISCOVERY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "USE_ENABLE_DISCOVERY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseEnableDiscovery("USE_ENABLE_DISCOVERY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "USE_ENABLE_DISCOVERY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseLogicFilterRules(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_logic_filter_rules"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseLogicFilterRules("USE_LOGIC_FILTER_RULES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "USE_LOGIC_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseLogicFilterRules("USE_LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "USE_LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseMgmPrivate(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_mgm_private"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseMgmPrivate("USE_MGM_PRIVATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_mgm_private", "USE_MGM_PRIVATE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseMgmPrivate("USE_MGM_PRIVATE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_mgm_private", "USE_MGM_PRIVATE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseOptions(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_options"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseOptions("USE_OPTIONS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "USE_OPTIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseOptions("USE_OPTIONS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "USE_OPTIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UsePreferredLifetime(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_preferred_lifetime"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUsePreferredLifetime("USE_PREFERRED_LIFETIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "USE_PREFERRED_LIFETIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUsePreferredLifetime("USE_PREFERRED_LIFETIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "USE_PREFERRED_LIFETIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseSubscribeSettings(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_subscribe_settings"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseSubscribeSettings("USE_SUBSCRIBE_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_subscribe_settings", "USE_SUBSCRIBE_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseSubscribeSettings("USE_SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_subscribe_settings", "USE_SUBSCRIBE_SETTINGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseUpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_update_dns_on_lease_renewal"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseUpdateDnsOnLeaseRenewal("USE_UPDATE_DNS_ON_LEASE_RENEWAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "USE_UPDATE_DNS_ON_LEASE_RENEWAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseUpdateDnsOnLeaseRenewal("USE_UPDATE_DNS_ON_LEASE_RENEWAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "USE_UPDATE_DNS_ON_LEASE_RENEWAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseValidLifetime(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_valid_lifetime"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseValidLifetime("USE_VALID_LIFETIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "USE_VALID_LIFETIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseValidLifetime("USE_VALID_LIFETIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "USE_VALID_LIFETIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_UseZoneAssociations(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_use_zone_associations"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerUseZoneAssociations("USE_ZONE_ASSOCIATIONS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_zone_associations", "USE_ZONE_ASSOCIATIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerUseZoneAssociations("USE_ZONE_ASSOCIATIONS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_zone_associations", "USE_ZONE_ASSOCIATIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_ValidLifetime(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_valid_lifetime"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerValidLifetime("VALID_LIFETIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "VALID_LIFETIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerValidLifetime("VALID_LIFETIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "VALID_LIFETIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networkcontainerResource_ZoneAssociations(t *testing.T) {
	var resourceName = "nios_ipam_ipv6network_container.test_zone_associations"
	var v ipam.Ipv6networkcontainer

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networkcontainerZoneAssociations("ZONE_ASSOCIATIONS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone_associations", "ZONE_ASSOCIATIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networkcontainerZoneAssociations("ZONE_ASSOCIATIONS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone_associations", "ZONE_ASSOCIATIONS_UPDATE_REPLACE_ME"),
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

func testAccIpv6networkcontainerDdnsDomainname(ddnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_ddns_domainname" {
    ddns_domainname = %q
}
`, ddnsDomainname)
}

func testAccIpv6networkcontainerDdnsEnableOptionFqdn(ddnsEnableOptionFqdn string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_ddns_enable_option_fqdn" {
    ddns_enable_option_fqdn = %q
}
`, ddnsEnableOptionFqdn)
}

func testAccIpv6networkcontainerDdnsGenerateHostname(ddnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_ddns_generate_hostname" {
    ddns_generate_hostname = %q
}
`, ddnsGenerateHostname)
}

func testAccIpv6networkcontainerDdnsServerAlwaysUpdates(ddnsServerAlwaysUpdates string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_ddns_server_always_updates" {
    ddns_server_always_updates = %q
}
`, ddnsServerAlwaysUpdates)
}

func testAccIpv6networkcontainerDdnsTtl(ddnsTtl string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_ddns_ttl" {
    ddns_ttl = %q
}
`, ddnsTtl)
}

func testAccIpv6networkcontainerDiscoveryBasicPollSettings(discoveryBasicPollSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_discovery_basic_poll_settings" {
    discovery_basic_poll_settings = %q
}
`, discoveryBasicPollSettings)
}

func testAccIpv6networkcontainerDiscoveryBlackoutSetting(discoveryBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_discovery_blackout_setting" {
    discovery_blackout_setting = %q
}
`, discoveryBlackoutSetting)
}

func testAccIpv6networkcontainerDiscoveryMember(discoveryMember string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_discovery_member" {
    discovery_member = %q
}
`, discoveryMember)
}

func testAccIpv6networkcontainerDomainNameServers(domainNameServers string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_domain_name_servers" {
    domain_name_servers = %q
}
`, domainNameServers)
}

func testAccIpv6networkcontainerEnableDdns(enableDdns string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_enable_ddns" {
    enable_ddns = %q
}
`, enableDdns)
}

func testAccIpv6networkcontainerEnableDiscovery(enableDiscovery string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_enable_discovery" {
    enable_discovery = %q
}
`, enableDiscovery)
}

func testAccIpv6networkcontainerExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccIpv6networkcontainerFederatedRealms(federatedRealms string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_federated_realms" {
    federated_realms = %q
}
`, federatedRealms)
}

func testAccIpv6networkcontainerLogicFilterRules(logicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_logic_filter_rules" {
    logic_filter_rules = %q
}
`, logicFilterRules)
}

func testAccIpv6networkcontainerMgmPrivate(mgmPrivate string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_mgm_private" {
    mgm_private = %q
}
`, mgmPrivate)
}

func testAccIpv6networkcontainerMsAdUserData(msAdUserData string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_ms_ad_user_data" {
    ms_ad_user_data = %q
}
`, msAdUserData)
}

func testAccIpv6networkcontainerNetwork(network string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_network" {
    network = %q
}
`, network)
}

func testAccIpv6networkcontainerNetworkView(networkView string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_network_view" {
    network_view = %q
}
`, networkView)
}

func testAccIpv6networkcontainerOptions(options string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_options" {
    options = %q
}
`, options)
}

func testAccIpv6networkcontainerPortControlBlackoutSetting(portControlBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_port_control_blackout_setting" {
    port_control_blackout_setting = %q
}
`, portControlBlackoutSetting)
}

func testAccIpv6networkcontainerPreferredLifetime(preferredLifetime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_preferred_lifetime" {
    preferred_lifetime = %q
}
`, preferredLifetime)
}

func testAccIpv6networkcontainerRirOrganization(rirOrganization string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_rir_organization" {
    rir_organization = %q
}
`, rirOrganization)
}

func testAccIpv6networkcontainerRirRegistrationStatus(rirRegistrationStatus string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_rir_registration_status" {
    rir_registration_status = %q
}
`, rirRegistrationStatus)
}

func testAccIpv6networkcontainerSamePortControlDiscoveryBlackout(samePortControlDiscoveryBlackout string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_same_port_control_discovery_blackout" {
    same_port_control_discovery_blackout = %q
}
`, samePortControlDiscoveryBlackout)
}

func testAccIpv6networkcontainerSubscribeSettings(subscribeSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_subscribe_settings" {
    subscribe_settings = %q
}
`, subscribeSettings)
}

func testAccIpv6networkcontainerUnmanaged(unmanaged string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_unmanaged" {
    unmanaged = %q
}
`, unmanaged)
}

func testAccIpv6networkcontainerUpdateDnsOnLeaseRenewal(updateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_update_dns_on_lease_renewal" {
    update_dns_on_lease_renewal = %q
}
`, updateDnsOnLeaseRenewal)
}

func testAccIpv6networkcontainerUseBlackoutSetting(useBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_blackout_setting" {
    use_blackout_setting = %q
}
`, useBlackoutSetting)
}

func testAccIpv6networkcontainerUseDdnsDomainname(useDdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_ddns_domainname" {
    use_ddns_domainname = %q
}
`, useDdnsDomainname)
}

func testAccIpv6networkcontainerUseDdnsEnableOptionFqdn(useDdnsEnableOptionFqdn string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_ddns_enable_option_fqdn" {
    use_ddns_enable_option_fqdn = %q
}
`, useDdnsEnableOptionFqdn)
}

func testAccIpv6networkcontainerUseDdnsGenerateHostname(useDdnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_ddns_generate_hostname" {
    use_ddns_generate_hostname = %q
}
`, useDdnsGenerateHostname)
}

func testAccIpv6networkcontainerUseDdnsTtl(useDdnsTtl string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_ddns_ttl" {
    use_ddns_ttl = %q
}
`, useDdnsTtl)
}

func testAccIpv6networkcontainerUseDiscoveryBasicPollingSettings(useDiscoveryBasicPollingSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_discovery_basic_polling_settings" {
    use_discovery_basic_polling_settings = %q
}
`, useDiscoveryBasicPollingSettings)
}

func testAccIpv6networkcontainerUseDomainNameServers(useDomainNameServers string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_domain_name_servers" {
    use_domain_name_servers = %q
}
`, useDomainNameServers)
}

func testAccIpv6networkcontainerUseEnableDdns(useEnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_enable_ddns" {
    use_enable_ddns = %q
}
`, useEnableDdns)
}

func testAccIpv6networkcontainerUseEnableDiscovery(useEnableDiscovery string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_enable_discovery" {
    use_enable_discovery = %q
}
`, useEnableDiscovery)
}

func testAccIpv6networkcontainerUseLogicFilterRules(useLogicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_logic_filter_rules" {
    use_logic_filter_rules = %q
}
`, useLogicFilterRules)
}

func testAccIpv6networkcontainerUseMgmPrivate(useMgmPrivate string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_mgm_private" {
    use_mgm_private = %q
}
`, useMgmPrivate)
}

func testAccIpv6networkcontainerUseOptions(useOptions string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_options" {
    use_options = %q
}
`, useOptions)
}

func testAccIpv6networkcontainerUsePreferredLifetime(usePreferredLifetime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_preferred_lifetime" {
    use_preferred_lifetime = %q
}
`, usePreferredLifetime)
}

func testAccIpv6networkcontainerUseSubscribeSettings(useSubscribeSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_subscribe_settings" {
    use_subscribe_settings = %q
}
`, useSubscribeSettings)
}

func testAccIpv6networkcontainerUseUpdateDnsOnLeaseRenewal(useUpdateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_update_dns_on_lease_renewal" {
    use_update_dns_on_lease_renewal = %q
}
`, useUpdateDnsOnLeaseRenewal)
}

func testAccIpv6networkcontainerUseValidLifetime(useValidLifetime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_valid_lifetime" {
    use_valid_lifetime = %q
}
`, useValidLifetime)
}

func testAccIpv6networkcontainerUseZoneAssociations(useZoneAssociations string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_use_zone_associations" {
    use_zone_associations = %q
}
`, useZoneAssociations)
}

func testAccIpv6networkcontainerValidLifetime(validLifetime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_valid_lifetime" {
    valid_lifetime = %q
}
`, validLifetime)
}

func testAccIpv6networkcontainerZoneAssociations(zoneAssociations string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6network_container" "test_zone_associations" {
    zone_associations = %q
}
`, zoneAssociations)
}
