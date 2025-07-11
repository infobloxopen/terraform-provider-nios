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

// TODO: DiscoveryMember requires discovering to be enabled.
// TODO: EnableDiscovery requires a valid discovery member.
// TODO: EnableImmediateDiscovery requires a valid discovery member.
// TODO: Federated realms serve need to enabled
// TODO: LogicFilterRules Logic filter rule required
// TODO: RemoveSubnets Need child objects and only delete param
// TODO: RirOrganization rir organization configuration required
// TODO: RirOrganizationAction rir organization configuration required
// TODO: ZoneAssociations Need dns zone to test associations
// TODO: MappedEAAttributes Need ISE server to test mapped_ea_attributes

var readableAttributesForNetworkcontainer = "authority,bootfile,bootserver,cloud_info,comment,ddns_domainname,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,ddns_update_fixed_addresses,ddns_use_option81,deny_bootp,discover_now_status,discovery_basic_poll_settings,discovery_blackout_setting,discovery_engine_type,discovery_member,email_list,enable_ddns,enable_dhcp_thresholds,enable_discovery,enable_email_warnings,enable_pxe_lease_time,enable_snmp_warnings,endpoint_sources,extattrs,federated_realms,high_water_mark,high_water_mark_reset,ignore_dhcp_option_list_request,ignore_id,ignore_mac_addresses,ipam_email_addresses,ipam_threshold_settings,ipam_trap_settings,last_rir_registration_update_sent,last_rir_registration_update_status,lease_scavenge_time,logic_filter_rules,low_water_mark,low_water_mark_reset,mgm_private,mgm_private_overridable,ms_ad_user_data,network,network_container,network_view,nextserver,options,port_control_blackout_setting,pxe_lease_time,recycle_leases,rir,rir_organization,rir_registration_status,same_port_control_discovery_blackout,subscribe_settings,unmanaged,update_dns_on_lease_renewal,use_authority,use_blackout_setting,use_bootfile,use_bootserver,use_ddns_domainname,use_ddns_generate_hostname,use_ddns_ttl,use_ddns_update_fixed_addresses,use_ddns_use_option81,use_deny_bootp,use_discovery_basic_polling_settings,use_email_list,use_enable_ddns,use_enable_dhcp_thresholds,use_enable_discovery,use_ignore_dhcp_option_list_request,use_ignore_id,use_ipam_email_addresses,use_ipam_threshold_settings,use_ipam_trap_settings,use_lease_scavenge_time,use_logic_filter_rules,use_mgm_private,use_nextserver,use_options,use_pxe_lease_time,use_recycle_leases,use_subscribe_settings,use_update_dns_on_lease_renewal,use_zone_associations,utilization,zone_associations"

func TestAccNetworkcontainerResource_basic(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerBasicConfig(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					// Check default values are populated correctly
					resource.TestCheckResourceAttr(resourceName, "authority", "false"),
					resource.TestCheckResourceAttr(resourceName, "auto_create_reversezone", "false"),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.authority_type", "GM"),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.delegated_scope", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.mgmt_platform", ""),
					resource.TestCheckResourceAttr(resourceName, "cloud_info.owned_by_adaptor", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "true"),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "0"),
					resource.TestCheckResourceAttr(resourceName, "ddns_update_fixed_addresses", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "false"),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_dhcp_thresholds", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_email_warnings", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_snmp_warnings", "false"),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "95"),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "85"),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "false"),
					resource.TestCheckResourceAttr(resourceName, "ignore_id", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "-1"),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "0"),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "10"),
					resource.TestCheckResourceAttr(resourceName, "mgm_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "mgm_private_overridable", "true"),
					resource.TestCheckResourceAttr(resourceName, "network_view", "default"),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "true"),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_status", "NOT_REGISTERED"),
					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "false"),
					resource.TestCheckResourceAttr(resourceName, "unmanaged", "false"),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_authority", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_update_fixed_addresses", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_email_list", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_id", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_email_addresses", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_threshold_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_trap_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_mgm_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_subscribe_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_zone_associations", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_disappears(t *testing.T) {
	resourceName := "nios_ipam_network_container.test"
	var v ipam.Networkcontainer
	// Generate a random CIDR network for the test
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetworkcontainerDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkcontainerBasicConfig(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					testAccCheckNetworkcontainerDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNetworkcontainerResource_Authority(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_authority"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerAuthority(network, "false", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "authority", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_authority", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerAuthority(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "authority", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_authority", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_Bootfile(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_bootfile"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerBootfile(network, "bootfile", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "bootfile"),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerBootfile(network, "bootfile_updated", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "bootfile_updated"),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_Bootserver(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_bootserver"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerBootserver(network, "test_bootserver", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "test_bootserver"),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerBootserver(network, "test_bootserver_updated", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "test_bootserver_updated"),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_CloudInfo(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_cloud_info"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerCloudInfo(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
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

func TestAccNetworkcontainerResource_Comment(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_comment"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerComment(network, "test comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerComment(network, "updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "updated comment"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_DdnsDomainname(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_ddns_domainname"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerDdnsDomainname(network, "test.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "test.com"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerDdnsDomainname(network, "testupdated.com", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "testupdated.com"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_DdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_ddns_generate_hostname"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerDdnsGenerateHostname(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerDdnsGenerateHostname(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_DdnsServerAlwaysUpdates(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_ddns_server_always_updates"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerDdnsServerAlwaysUpdates(network, "true", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerDdnsServerAlwaysUpdates(network, "false", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_DdnsTtl(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_ddns_ttl"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerDdnsTtl(network, "1", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "1"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerDdnsTtl(network, "2", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "2"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_DdnsUpdateFixedAddresses(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_ddns_update_fixed_addresses"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerDdnsUpdateFixedAddresses(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_update_fixed_addresses", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_update_fixed_addresses", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerDdnsUpdateFixedAddresses(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_update_fixed_addresses", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_update_fixed_addresses", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_DdnsUseOption81(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_ddns_use_option81"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerDdnsUseOption81(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerDdnsUseOption81(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_DenyBootp(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_deny_bootp"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerDenyBootp(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerDenyBootp(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_DiscoveryBasicPollSettings(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_discovery_basic_poll_settings"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerDiscoveryBasicPollSettings(network, "true", "true", "false", "default", "false", "false", "1", "false", "false", "true", "PERIODIC", "3600", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
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
				Config: testAccNetworkcontainerDiscoveryBasicPollSettings(network, "false", "true", "false", "default", "false", "false", "1", "false", "false", "true", "PERIODIC", "3600", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
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

func TestAccNetworkcontainerResource_DiscoveryBlackoutSetting(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_discovery_blackout_setting"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerDiscoveryBlackoutSetting(network, "false", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovery_blackout_setting.enable_blackout", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_EmailList(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_email_list"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerEmailList(network, "test@infoblox.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_list.0", "test@infoblox.com"),
					resource.TestCheckResourceAttr(resourceName, "use_email_list", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerEmailList(network, "update@test.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_list.0", "update@test.com"),
					resource.TestCheckResourceAttr(resourceName, "use_email_list", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_EnableDdns(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_enable_ddns"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerEnableDdns(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerEnableDdns(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_EnableDhcpThresholds(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_enable_dhcp_thresholds"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerEnableDhcpThresholds(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_dhcp_thresholds", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerEnableDhcpThresholds(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_dhcp_thresholds", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_EnableEmailWarnings(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_enable_email_warnings"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerEnableEmailWarnings(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_email_warnings", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerEnableEmailWarnings(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_email_warnings", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_EnablePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_enable_pxe_lease_time"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerEnablePxeLeaseTime(network, "100", "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "100"),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerEnablePxeLeaseTime(network, "100", "true", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "true"),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "100"),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_EnableSnmpWarnings(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_enable_snmp_warnings"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerEnableSnmpWarnings(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_snmp_warnings", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerEnableSnmpWarnings(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_snmp_warnings", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_extattrs"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerExtAttrs(network, map[string]string{"Site": extAttrValue1}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerExtAttrs(network, map[string]string{"Site": extAttrValue2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_HighWaterMark(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_high_water_mark"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerHighWaterMark(network, "95"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "95"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerHighWaterMark(network, "90"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "90"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_HighWaterMarkReset(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_high_water_mark_reset"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerHighWaterMarkReset(network, "85"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "85"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerHighWaterMarkReset(network, "80"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "80"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_IgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_ignore_dhcp_option_list_request"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerIgnoreDhcpOptionListRequest(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerIgnoreDhcpOptionListRequest(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_IgnoreId(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_ignore_id"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerIgnoreId(network, "NONE", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_id", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_id", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerIgnoreId(network, "MACADDR", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_id", "MACADDR"),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_id", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_IgnoreMacAddresses(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_ignore_mac_addresses"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerIgnoreMacAddresses(network, "aa:bb:cc:dd:ee:ff"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_mac_addresses.0", "aa:bb:cc:dd:ee:ff"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerIgnoreMacAddresses(network, "ff:ee:dd:cc:bb:aa"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_mac_addresses.0", "ff:ee:dd:cc:bb:aa"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_IpamEmailAddresses(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_ipam_email_addresses"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerIpamEmailAddresses(network, "testuser@infoblox.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_email_addresses.0", "testuser@infoblox.com"),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_email_addresses", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerIpamEmailAddresses(network, "testuserupdated@infoblox.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_email_addresses.0", "testuserupdated@infoblox.com"),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_email_addresses", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_IpamThresholdSettings(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_ipam_threshold_settings"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerIpamThresholdSettings(network, "85", "95", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_threshold_settings.reset_value", "85"),
					resource.TestCheckResourceAttr(resourceName, "ipam_threshold_settings.trigger_value", "95"),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_threshold_settings", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerIpamThresholdSettings(network, "75", "80", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_threshold_settings.reset_value", "75"),
					resource.TestCheckResourceAttr(resourceName, "ipam_threshold_settings.trigger_value", "80"),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_threshold_settings", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_IpamTrapSettings(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_ipam_trap_settings"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerIpamTrapSettings(network, "false", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_trap_settings.enable_email_warnings", "false"),
					resource.TestCheckResourceAttr(resourceName, "ipam_trap_settings.enable_snmp_warnings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerIpamTrapSettings(network, "true", "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_trap_settings.enable_email_warnings", "true"),
					resource.TestCheckResourceAttr(resourceName, "ipam_trap_settings.enable_snmp_warnings", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_LeaseScavengeTime(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_lease_scavenge_time"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerLeaseScavengeTime(network, "-1", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "-1"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerLeaseScavengeTime(network, "86400", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "86400"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_LowWaterMark(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_low_water_mark"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerLowWaterMark(network, "0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "0"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerLowWaterMark(network, "50"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "50"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_LowWaterMarkReset(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_low_water_mark_reset"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerLowWaterMarkReset(network, "10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerLowWaterMarkReset(network, "20"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_MgmPrivate(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_mgm_private"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerMgmPrivate(network, "false", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mgm_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_mgm_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerMgmPrivate(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mgm_private", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_mgm_private", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_Network(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_network"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerNetwork(network),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_NetworkView(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_network_view"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerNetworkView(network, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "default"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_Nextserver(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_nextserver"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerNextserver(network, "1.1.1.1", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", "1.1.1.1"),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerNextserver(network, "1.1.1.2", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", "1.1.1.2"),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_Options(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_options"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read - Test with special option that supports use_option
			{
				Config: testAccNetworkcontainerOptions(network, "dhcp-lease-time", "51", "7200", "DHCP", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "7200"),
					resource.TestCheckResourceAttr(resourceName, "options.0.vendor_class", "DHCP"),
					resource.TestCheckResourceAttr(resourceName, "options.0.use_option", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read - Test with another special option (use_option should be preserved)
			{
				Config: testAccNetworkcontainerOptions(network, "dhcp-lease-time", "51", "7300", "DHCP", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "7300"),
					resource.TestCheckResourceAttr(resourceName, "options.0.vendor_class", "DHCP"),
					resource.TestCheckResourceAttr(resourceName, "options.0.use_option", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_PortControlBlackoutSetting(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_port_control_blackout_setting"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerPortControlBlackoutSetting(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port_control_blackout_setting.enable_blackout", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_PxeLeaseTime(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_pxe_lease_time"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerPxeLeaseTime(network, "0", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "0"),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerPxeLeaseTime(network, "40000", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "40000"),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_RecycleLeases(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_recycle_leases"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerRecycleLeases(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerRecycleLeases(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_RirRegistrationStatus(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_rir_registration_status"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerRirRegistrationStatus(network, "NOT_REGISTERED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_status", "NOT_REGISTERED"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_SamePortControlDiscoveryBlackout(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_same_port_control_discovery_blackout"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerSamePortControlDiscoveryBlackout(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerSamePortControlDiscoveryBlackout(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "same_port_control_discovery_blackout", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_Unmanaged(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_unmanaged"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUnmanaged(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "unmanaged", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_update_dns_on_lease_renewal"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUpdateDnsOnLeaseRenewal(network, "false", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUpdateDnsOnLeaseRenewal(network, "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseAuthority(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_authority"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseAuthority(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_authority", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseAuthority(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_authority", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseBlackoutSetting(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_blackout_setting"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseBlackoutSetting(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseBlackoutSetting(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_blackout_setting", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseBootfile(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_bootfile"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseBootfile(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseBootfile(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseBootserver(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_bootserver"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseBootserver(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseBootserver(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseDdnsDomainname(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_ddns_domainname"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseDdnsDomainname(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseDdnsDomainname(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseDdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_ddns_generate_hostname"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseDdnsGenerateHostname(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseDdnsGenerateHostname(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseDdnsTtl(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_ddns_ttl"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseDdnsTtl(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseDdnsTtl(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseDdnsUpdateFixedAddresses(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_ddns_update_fixed_addresses"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseDdnsUpdateFixedAddresses(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_update_fixed_addresses", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseDdnsUpdateFixedAddresses(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_update_fixed_addresses", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseDdnsUseOption81(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_ddns_use_option81"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseDdnsUseOption81(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseDdnsUseOption81(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseDenyBootp(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_deny_bootp"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseDenyBootp(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseDenyBootp(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseDiscoveryBasicPollingSettings(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_discovery_basic_polling_settings"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseDiscoveryBasicPollingSettings(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseDiscoveryBasicPollingSettings(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_discovery_basic_polling_settings", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseEmailList(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_email_list"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseEmailList(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_email_list", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseEmailList(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_email_list", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseEnableDdns(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_enable_ddns"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseEnableDdns(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseEnableDdns(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseEnableDhcpThresholds(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_enable_dhcp_thresholds"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseEnableDhcpThresholds(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseEnableDhcpThresholds(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseEnableDiscovery(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_enable_discovery"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseEnableDiscovery(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseEnableDiscovery(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_discovery", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseIgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_ignore_dhcp_option_list_request"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseIgnoreDhcpOptionListRequest(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseIgnoreDhcpOptionListRequest(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseIgnoreId(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_ignore_id"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseIgnoreId(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_id", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseIgnoreId(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_id", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseIpamEmailAddresses(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_ipam_email_addresses"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseIpamEmailAddresses(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_email_addresses", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseIpamEmailAddresses(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_email_addresses", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseIpamThresholdSettings(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_ipam_threshold_settings"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseIpamThresholdSettings(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_threshold_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseIpamThresholdSettings(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_threshold_settings", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseIpamTrapSettings(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_ipam_trap_settings"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseIpamTrapSettings(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_trap_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseIpamTrapSettings(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_trap_settings", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseLeaseScavengeTime(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_lease_scavenge_time"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseLeaseScavengeTime(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseLeaseScavengeTime(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseLogicFilterRules(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_logic_filter_rules"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseLogicFilterRules(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseLogicFilterRules(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseMgmPrivate(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_mgm_private"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseMgmPrivate(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_mgm_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseNextserver(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_nextserver"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseNextserver(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseNextserver(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseOptions(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_options"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseOptions(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseOptions(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UsePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_pxe_lease_time"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUsePxeLeaseTime(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUsePxeLeaseTime(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseRecycleLeases(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_recycle_leases"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseRecycleLeases(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseRecycleLeases(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseSubscribeSettings(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_subscribe_settings"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseSubscribeSettings(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_subscribe_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseUpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_update_dns_on_lease_renewal"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseUpdateDnsOnLeaseRenewal(network, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "false"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkcontainerUseUpdateDnsOnLeaseRenewal(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkcontainerResource_UseZoneAssociations(t *testing.T) {
	var resourceName = "nios_ipam_network_container.test_use_zone_associations"
	var v ipam.Networkcontainer
	network := acctest.RandomCIDRNetwork()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkcontainerUseZoneAssociations(network, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkcontainerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_zone_associations", "true"),
					resource.TestCheckResourceAttr(resourceName, "network", network),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

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

func testAccCheckNetworkcontainerDisappears(ctx context.Context, v *ipam.Networkcontainer) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.IPAMAPI.
			NetworkcontainerAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccNetworkcontainerBasicConfig(network string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test" {
	network = %q
}
`, network)
}

func testAccNetworkcontainerAuthority(network, authority, useAuthority string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_authority" {
	network = %q
    authority = %q
	use_authority = %q
}
`, network, authority, useAuthority)
}

func testAccNetworkcontainerBootfile(network, bootfile, useBootfile string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_bootfile" {
    network = %q
    bootfile = %q
    use_bootfile = %q
}
`, network, bootfile, useBootfile)
}

func testAccNetworkcontainerBootserver(network, bootserver, useBootserver string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_bootserver" {
    network = %q
    bootserver = %q
    use_bootserver = %q
}
`, network, bootserver, useBootserver)
}

func testAccNetworkcontainerCloudInfo(network string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_cloud_info" {
    network = %q
}
`, network)
}

func testAccNetworkcontainerComment(network, comment string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_comment" {
    network = %q
    comment = %q
}
`, network, comment)
}

func testAccNetworkcontainerDdnsDomainname(network, ddnsDomainname, useDdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_ddns_domainname" {
    network = %q
    ddns_domainname = %q
	use_ddns_domainname = %q
}
`, network, ddnsDomainname, useDdnsDomainname)
}

func testAccNetworkcontainerDdnsGenerateHostname(network, ddnsGenerateHostname, useDdnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_ddns_generate_hostname" {
    network = %q
    ddns_generate_hostname = %q
    use_ddns_generate_hostname = %q
}
`, network, ddnsGenerateHostname, useDdnsGenerateHostname)
}

func testAccNetworkcontainerDdnsServerAlwaysUpdates(network, ddnsServerAlwaysUpdates, ddnsUseOption81, useDdnsUseOption81 string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_ddns_server_always_updates" {
    network = %q
    ddns_server_always_updates = %q
    ddns_use_option81 = %q
	use_ddns_use_option81 = %q
}
`, network, ddnsServerAlwaysUpdates, ddnsUseOption81, useDdnsUseOption81)
}

func testAccNetworkcontainerDdnsTtl(network, ddnsTtl, useDdnsTtl string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_ddns_ttl" {
    network = %q
    ddns_ttl = %q
    use_ddns_ttl = %q
}
`, network, ddnsTtl, useDdnsTtl)
}

func testAccNetworkcontainerDdnsUpdateFixedAddresses(network, ddnsUpdateFixedAddresses, useDdnsUpdateFixedAddresses string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_ddns_update_fixed_addresses" {
    network = %q
    ddns_update_fixed_addresses = %q
    use_ddns_update_fixed_addresses = %q
}
`, network, ddnsUpdateFixedAddresses, useDdnsUpdateFixedAddresses)
}

func testAccNetworkcontainerDdnsUseOption81(network, ddnsUseOption81, useDdnsUseOption81 string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_ddns_use_option81" {
    network = %q
    ddns_use_option81 = %q
    use_ddns_use_option81 = %q
}
`, network, ddnsUseOption81, useDdnsUseOption81)
}

func testAccNetworkcontainerDenyBootp(network, denyBootp, useDenyBootp string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_deny_bootp" {
    network = %q
    deny_bootp = %q
    use_deny_bootp = %q
}
`, network, denyBootp, useDenyBootp)
}

func testAccNetworkcontainerDiscoveryBasicPollSettings(network, autoArpRefreshBeforeSwitchPortPolling, cliCollection, completePingSweep, credentialGroup, deviceProfile, netbiosScanning, pollingFrequencyModifier, portScanning, smartSubnetPingSweep, snmpCollection, switchPortDataCollectionPolling, switchPortDataCollectionPollingInterval, useGlobalPollingFrequencyModifier, useDiscoveryBasicPollSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_discovery_basic_poll_settings" {
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

func testAccNetworkcontainerDiscoveryBlackoutSetting(network, enabledBlackout, useBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_discovery_blackout_setting" {
    network = %q
    discovery_blackout_setting = {
		enabled_blackout = %q
	}
    use_blackout_setting = %q
}
`, network, enabledBlackout, useBlackoutSetting)
}

func testAccNetworkcontainerEmailList(network, emailList, useEmailList string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_email_list" {
    network = %q
    email_list = [%q]
    use_email_list = %q
}
`, network, emailList, useEmailList)
}

func testAccNetworkcontainerEnableDdns(network, enableDdns, useEnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_enable_ddns" {
    network = %q
    enable_ddns = %q
    use_enable_ddns = %q
}
`, network, enableDdns, useEnableDdns)
}

func testAccNetworkcontainerEnableDhcpThresholds(network, enableDhcpThresholds, useEnableDhcpThresholds string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_enable_dhcp_thresholds" {
    network = %q
    enable_dhcp_thresholds = %q
    use_enable_dhcp_thresholds = %q
}
`, network, enableDhcpThresholds, useEnableDhcpThresholds)
}

func testAccNetworkcontainerEnableEmailWarnings(network, enableEmailWarnings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_enable_email_warnings" {
    network = %q
    enable_email_warnings = %q
}
`, network, enableEmailWarnings)
}

func testAccNetworkcontainerEnablePxeLeaseTime(network, pxeLeaseTime, enablePxeLeaseTime, usePxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_enable_pxe_lease_time" {
    network = %q
    pxe_lease_time = %q
    enable_pxe_lease_time = %q
	use_pxe_lease_time = %q

}
`, network, pxeLeaseTime, enablePxeLeaseTime, usePxeLeaseTime)
}

func testAccNetworkcontainerEnableSnmpWarnings(network, enableSnmpWarnings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_enable_snmp_warnings" {
    network = %q
    enable_snmp_warnings = %q
}
`, network, enableSnmpWarnings)
}

func testAccNetworkcontainerExtAttrs(network string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_extattrs" {
    network = %q
    extattrs = %s
}
`, network, extattrsStr)
}

func testAccNetworkcontainerHighWaterMark(network, highWaterMark string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_high_water_mark" {
    network = %q
    high_water_mark = %q
}
`, network, highWaterMark)
}

func testAccNetworkcontainerHighWaterMarkReset(network, highWaterMarkReset string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_high_water_mark_reset" {
    network = %q
    high_water_mark_reset = %q
}
`, network, highWaterMarkReset)
}

func testAccNetworkcontainerIgnoreDhcpOptionListRequest(network, ignoreDhcpOptionListRequest, useIgnoreDhcpOptionListRequest string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_ignore_dhcp_option_list_request" {
    network = %q
    ignore_dhcp_option_list_request = %q
    use_ignore_dhcp_option_list_request = %q
}
`, network, ignoreDhcpOptionListRequest, useIgnoreDhcpOptionListRequest)
}

func testAccNetworkcontainerIgnoreId(network, ignoreId, useIgnoreId, useBootfile string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_ignore_id" {
    network = %q
    ignore_id = %q
    use_ignore_id = %q
    use_bootfile = %q
}
`, network, ignoreId, useIgnoreId, useBootfile)
}

func testAccNetworkcontainerIgnoreMacAddresses(network, ignoreMacAddresses string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_ignore_mac_addresses" {
    network = %q
    ignore_mac_addresses = [%q]
}
`, network, ignoreMacAddresses)
}

func testAccNetworkcontainerIpamEmailAddresses(network, ipamEmailAddresses, useIpamEmailAddresses string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_ipam_email_addresses" {
    network = %q
    ipam_email_addresses = [%q]
    use_ipam_email_addresses = %q
}
`, network, ipamEmailAddresses, useIpamEmailAddresses)
}

func testAccNetworkcontainerIpamThresholdSettings(network, resetValue, triggerValue, useIpamThresholdSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_ipam_threshold_settings" {
    network = %q
    ipam_threshold_settings = {
        reset_value = %q
        trigger_value = %q
    }
	use_ipam_threshold_settings = %q
}
`, network, resetValue, triggerValue, useIpamThresholdSettings)
}

func testAccNetworkcontainerIpamTrapSettings(network, enableEmailWarnings, enableSnmpWarnings, useIpamTrapSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_ipam_trap_settings" {
    network = %q
    ipam_trap_settings = {
        enable_email_warnings = %q
        enable_snmp_warnings = %q
    }
	use_ipam_trap_settings = %q
}
`, network, enableEmailWarnings, enableSnmpWarnings, useIpamTrapSettings)
}

func testAccNetworkcontainerLeaseScavengeTime(network, leaseScavengeTime, useLeaseScavengeTime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_lease_scavenge_time" {
    network = %q
    lease_scavenge_time = %q
    use_lease_scavenge_time = %q
}
`, network, leaseScavengeTime, useLeaseScavengeTime)
}

func testAccNetworkcontainerLowWaterMark(network, lowWaterMark string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_low_water_mark" {
    network = %q
    low_water_mark = %q
}
`, network, lowWaterMark)
}

func testAccNetworkcontainerLowWaterMarkReset(network, lowWaterMarkReset string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_low_water_mark_reset" {
    network = %q
    low_water_mark_reset = %q
}
`, network, lowWaterMarkReset)
}

func testAccNetworkcontainerMgmPrivate(network, mgmPrivate, useMgmPrivate string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_mgm_private" {
    network = %q
    mgm_private = %q
    use_mgm_private = %q
}
`, network, mgmPrivate, useMgmPrivate)
}

func testAccNetworkcontainerNetwork(network string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_network" {
    network = %q
}
`, network)
}

func testAccNetworkcontainerNetworkView(network, networkView string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_network_view" {
    network = %q
    network_view = %q
}
`, network, networkView)
}

func testAccNetworkcontainerNextserver(network, nextserver, useNextserver string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_nextserver" {
    network = %q
    nextserver = %q
    use_nextserver = %q
}
`, network, nextserver, useNextserver)
}

func testAccNetworkcontainerOptions(network, name, num, value, vendorClass, useOption, useOptions string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_options" {
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

func testAccNetworkcontainerPortControlBlackoutSetting(network, enableBlackout, useBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_port_control_blackout_setting" {
	network = %q
    port_control_blackout_setting = {
		enable_blackout = %q
	}
	use_blackout_setting = %q
}
`, network, enableBlackout, useBlackoutSetting)
}

func testAccNetworkcontainerPxeLeaseTime(network, pxeLeaseTime, usePxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_pxe_lease_time" {
    network = %q
    pxe_lease_time = %q
    use_pxe_lease_time = %q
}
`, network, pxeLeaseTime, usePxeLeaseTime)
}

func testAccNetworkcontainerRecycleLeases(network, recycleLeases, useRecycleLeases string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_recycle_leases" {
    network = %q
    recycle_leases = %q
    use_recycle_leases = %q
}
`, network, recycleLeases, useRecycleLeases)
}

func testAccNetworkcontainerRirRegistrationStatus(network, rirRegistrationStatus string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_rir_registration_status" {
    network = %q
    rir_registration_status = %q
}
`, network, rirRegistrationStatus)
}

func testAccNetworkcontainerSamePortControlDiscoveryBlackout(network, samePortControlDiscoveryBlackout, useBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_same_port_control_discovery_blackout" {
    network = %q
    same_port_control_discovery_blackout = %q
    use_blackout_setting = %q
}
`, network, samePortControlDiscoveryBlackout, useBlackoutSetting)
}

func testAccNetworkcontainerUnmanaged(network, unmanaged string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_unmanaged" {
    network = %q
    unmanaged = %q
}
`, network, unmanaged)
}

func testAccNetworkcontainerUpdateDnsOnLeaseRenewal(network, updateDnsOnLeaseRenewal, useUpdateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_update_dns_on_lease_renewal" {
    network = %q
    update_dns_on_lease_renewal = %q
    use_update_dns_on_lease_renewal = %q
}
`, network, updateDnsOnLeaseRenewal, useUpdateDnsOnLeaseRenewal)
}

func testAccNetworkcontainerUseAuthority(network, useAuthority string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_authority" {
    network = %q
    use_authority = %q
}
`, network, useAuthority)
}

func testAccNetworkcontainerUseBlackoutSetting(network, useBlackoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_blackout_setting" {
	network = %q
    use_blackout_setting = %q
}
`, network, useBlackoutSetting)
}

func testAccNetworkcontainerUseBootfile(network, useBootfile string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_bootfile" {
    network = %q
    use_bootfile = %q
}
`, network, useBootfile)
}

func testAccNetworkcontainerUseBootserver(network, useBootserver string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_bootserver" {
    network = %q
    use_bootserver = %q
}
`, network, useBootserver)
}

func testAccNetworkcontainerUseDdnsDomainname(network, useDdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_ddns_domainname" {
    network = %q
    use_ddns_domainname = %q
}
`, network, useDdnsDomainname)
}

func testAccNetworkcontainerUseDdnsGenerateHostname(network, useDdnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_ddns_generate_hostname" {
    network = %q
    use_ddns_generate_hostname = %q
}
`, network, useDdnsGenerateHostname)
}

func testAccNetworkcontainerUseDdnsTtl(network, useDdnsTtl string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_ddns_ttl" {
    network = %q
    use_ddns_ttl = %q
}
`, network, useDdnsTtl)
}

func testAccNetworkcontainerUseDdnsUpdateFixedAddresses(network, useDdnsUpdateFixedAddresses string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_ddns_update_fixed_addresses" {
    network = %q
    use_ddns_update_fixed_addresses = %q
}
`, network, useDdnsUpdateFixedAddresses)
}

func testAccNetworkcontainerUseDdnsUseOption81(network, useDdnsUseOption81 string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_ddns_use_option81" {
    network = %q
    use_ddns_use_option81 = %q
}
`, network, useDdnsUseOption81)
}

func testAccNetworkcontainerUseDenyBootp(network, useDenyBootp string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_deny_bootp" {
    network = %q
    use_deny_bootp = %q
}
`, network, useDenyBootp)
}

func testAccNetworkcontainerUseDiscoveryBasicPollingSettings(network, useDiscoveryBasicPollingSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_discovery_basic_polling_settings" {
    network = %q
    use_discovery_basic_polling_settings = %q
}
`, network, useDiscoveryBasicPollingSettings)
}

func testAccNetworkcontainerUseEmailList(network, useEmailList string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_email_list" {
    network = %q
    use_email_list = %q
}
`, network, useEmailList)
}

func testAccNetworkcontainerUseEnableDdns(network, useEnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_enable_ddns" {
    network = %q
    use_enable_ddns = %q
}
`, network, useEnableDdns)
}

func testAccNetworkcontainerUseEnableDhcpThresholds(network, useEnableDhcpThresholds string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_enable_dhcp_thresholds" {
    network = %q
    use_enable_dhcp_thresholds = %q
}
`, network, useEnableDhcpThresholds)
}

func testAccNetworkcontainerUseEnableDiscovery(network, useEnableDiscovery string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_enable_discovery" {
    network = %q
    use_enable_discovery = %q
}
`, network, useEnableDiscovery)
}

func testAccNetworkcontainerUseIgnoreDhcpOptionListRequest(network, useIgnoreDhcpOptionListRequest string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_ignore_dhcp_option_list_request" {
    network = %q
    use_ignore_dhcp_option_list_request = %q
}
`, network, useIgnoreDhcpOptionListRequest)
}

func testAccNetworkcontainerUseIgnoreId(network, useIgnoreId string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_ignore_id" {
    network = %q
    use_ignore_id = %q
}
`, network, useIgnoreId)
}

func testAccNetworkcontainerUseIpamEmailAddresses(network, useIpamEmailAddresses string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_ipam_email_addresses" {
    network = %q
    use_ipam_email_addresses = %q
}
`, network, useIpamEmailAddresses)
}

func testAccNetworkcontainerUseIpamThresholdSettings(network, useIpamThresholdSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_ipam_threshold_settings" {
    network = %q
    use_ipam_threshold_settings = %q
}
`, network, useIpamThresholdSettings)
}

func testAccNetworkcontainerUseIpamTrapSettings(network, useIpamTrapSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_ipam_trap_settings" {
    network = %q
    use_ipam_trap_settings = %q
}
`, network, useIpamTrapSettings)
}

func testAccNetworkcontainerUseLeaseScavengeTime(network, useLeaseScavengeTime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_lease_scavenge_time" {
    network = %q
    use_lease_scavenge_time = %q
}
`, network, useLeaseScavengeTime)
}

func testAccNetworkcontainerUseLogicFilterRules(network, useLogicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_logic_filter_rules" {
    network = %q
    use_logic_filter_rules = %q
}
`, network, useLogicFilterRules)
}

func testAccNetworkcontainerUseMgmPrivate(network, useMgmPrivate string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_mgm_private" {
    network = %q
    use_mgm_private = %q
}
`, network, useMgmPrivate)
}

func testAccNetworkcontainerUseNextserver(network, useNextserver string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_nextserver" {
    network = %q
    use_nextserver = %q
}
`, network, useNextserver)
}

func testAccNetworkcontainerUseOptions(network, useOptions string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_options" {
    network = %q
    use_options = %q
}
`, network, useOptions)
}

func testAccNetworkcontainerUsePxeLeaseTime(network, usePxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_pxe_lease_time" {
    network = %q
    use_pxe_lease_time = %q
}
`, network, usePxeLeaseTime)
}

func testAccNetworkcontainerUseRecycleLeases(network, useRecycleLeases string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_recycle_leases" {
    network = %q
    use_recycle_leases = %q
}
`, network, useRecycleLeases)
}

func testAccNetworkcontainerUseSubscribeSettings(network, useSubscribeSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_subscribe_settings" {
    network = %q
    use_subscribe_settings = %q
}
`, network, useSubscribeSettings)
}

func testAccNetworkcontainerUseUpdateDnsOnLeaseRenewal(network, useUpdateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_update_dns_on_lease_renewal" {
    network = %q
    use_update_dns_on_lease_renewal = %q
}
`, network, useUpdateDnsOnLeaseRenewal)
}

func testAccNetworkcontainerUseZoneAssociations(network, useZoneAssociations string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network_container" "test_use_zone_associations" {
    network = %q
    use_zone_associations = %q
}
`, network, useZoneAssociations)
}
