package dhcp_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// TODO : Required parents for the execution of tests - logic_filter_rules (option_filter, option_logic_filter)
// TODO: - create NW using GenerateRandomCIDR, and get references of networks using Network resource
// TODO: - testcases related ignore_id, use_ignore_id, ignore_client_identifier and ignore_client_identifier to be revisited.

var readableAttributesForSharednetwork = "authority,bootfile,bootserver,comment,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,ddns_update_fixed_addresses,ddns_use_option81,deny_bootp,dhcp_utilization,dhcp_utilization_status,disable,dynamic_hosts,enable_ddns,enable_pxe_lease_time,extattrs,ignore_client_identifier,ignore_dhcp_option_list_request,ignore_id,ignore_mac_addresses,lease_scavenge_time,logic_filter_rules,ms_ad_user_data,name,network_view,networks,nextserver,options,pxe_lease_time,static_hosts,total_hosts,update_dns_on_lease_renewal,use_authority,use_bootfile,use_bootserver,use_ddns_generate_hostname,use_ddns_ttl,use_ddns_update_fixed_addresses,use_ddns_use_option81,use_deny_bootp,use_enable_ddns,use_ignore_client_identifier,use_ignore_dhcp_option_list_request,use_ignore_id,use_lease_scavenge_time,use_logic_filter_rules,use_nextserver,use_options,use_pxe_lease_time,use_update_dns_on_lease_renewal"

func TestAccSharednetworkResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkBasicConfig(name, networks),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "networks.#", fmt.Sprintf("%d", len(networks))),
					resource.TestCheckResourceAttrPair(resourceName, "networks.0.ref", "nios_ipam_network.test_network1", "ref"),
					resource.TestCheckResourceAttrPair(resourceName, "networks.1.ref", "nios_ipam_network.test_network2", "ref"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "authority", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "true"),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "0"),
					resource.TestCheckResourceAttr(resourceName, "ddns_update_fixed_addresses", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "false"),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "false"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "ignore_client_identifier", "false"),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "false"),
					resource.TestCheckResourceAttr(resourceName, "ignore_id", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "-1"),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_authority", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_update_fixed_addresses", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_client_identifier", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_id", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_shared_network.test"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharednetworkDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharednetworkBasicConfig(name, networks),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					testAccCheckSharednetworkDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSharednetworkResource_Authority(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_authority"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkAuthority(name, networks, true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "authority", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkAuthority(name, networks, false, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "authority", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_Bootfile(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_bootfile"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	bootFile := "boot.txt"
	bootFileUpdated := "boot_updated.txt"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkBootfile(name, networks, bootFile, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "boot.txt"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkBootfile(name, networks, bootFileUpdated, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "boot_updated.txt"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_Bootserver(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_bootserver"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	bootServer := "boot-server1"
	bootServerUpdated := "boot-server2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkBootserver(name, networks, bootServer, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "boot-server1"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkBootserver(name, networks, bootServerUpdated, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "boot-server2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_comment"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	comment := "shared network comment"
	commentUpdated := "updated shared network comment"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkComment(name, networks, comment),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "shared network comment"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkComment(name, networks, commentUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "updated shared network comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_DdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_ddns_generate_hostname"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkDdnsGenerateHostname(name, networks, true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkDdnsGenerateHostname(name, networks, false, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_DdnsServerAlwaysUpdates(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_ddns_server_always_updates"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkDdnsServerAlwaysUpdates(name, networks, true, true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkDdnsServerAlwaysUpdates(name, networks, false, true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_DdnsTtl(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_ddns_ttl"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	ddnsTtl := 3600
	ddnsTtlUpdated := 7200

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkDdnsTtl(name, networks, ddnsTtl, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "3600"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkDdnsTtl(name, networks, ddnsTtlUpdated, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "7200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_DdnsUpdateFixedAddresses(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_ddns_update_fixed_addresses"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkDdnsUpdateFixedAddresses(name, networks, true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_update_fixed_addresses", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkDdnsUpdateFixedAddresses(name, networks, false, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_update_fixed_addresses", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_DdnsUseOption81(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_ddns_use_option81"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkDdnsUseOption81(name, networks, true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkDdnsUseOption81(name, networks, false, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_DenyBootp(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_deny_bootp"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkDenyBootp(name, networks, true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkDenyBootp(name, networks, false, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_Disable(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_disable"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkDisable(name, networks, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkDisable(name, networks, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_EnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_enable_ddns"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkEnableDdns(name, networks, true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkEnableDdns(name, networks, false, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_EnablePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_enable_pxe_lease_time"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	pxeLeaseTime := 43200
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkEnablePxeLeaseTime(name, networks, true, true, pxeLeaseTime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkEnablePxeLeaseTime(name, networks, false, true, pxeLeaseTime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_extattrs"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkExtAttrs(name, networks, map[string]string{"Site": extAttrValue1}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkExtAttrs(name, networks, map[string]string{"Site": extAttrValue2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_IgnoreClientIdentifier(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_ignore_client_identifier"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkIgnoreClientIdentifier(name, networks, false, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_client_identifier", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkIgnoreClientIdentifier(name, networks, true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_client_identifier", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_IgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_ignore_dhcp_option_list_request"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkIgnoreDhcpOptionListRequest(name, networks, true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkIgnoreDhcpOptionListRequest(name, networks, false, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_IgnoreId(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_ignore_id"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	ignoreId := "CLIENT"
	ignoreIdUpdated := "NONE"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkIgnoreId(name, networks, ignoreId, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_id", "CLIENT"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkIgnoreId(name, networks, ignoreIdUpdated, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_id", "NONE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_IgnoreMacAddresses(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_ignore_mac_addresses"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	ignoreMacAddresses := []string{"00:11:22:33:44:55", "66:77:88:99:aa:bb"}
	ignoreMacAddressesUpdated := []string{"00:11:22:33:44:88", "00:11:22:33:44:55"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkIgnoreMacAddresses(name, networks, ignoreMacAddresses),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_mac_addresses.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "ignore_mac_addresses.0", "00:11:22:33:44:55"),
					resource.TestCheckResourceAttr(resourceName, "ignore_mac_addresses.1", "66:77:88:99:aa:bb"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkIgnoreMacAddresses(name, networks, ignoreMacAddressesUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_mac_addresses.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "ignore_mac_addresses.0", "00:11:22:33:44:88"),
					resource.TestCheckResourceAttr(resourceName, "ignore_mac_addresses.1", "00:11:22:33:44:55"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_LeaseScavengeTime(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_lease_scavenge_time"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	leaseScavengeTime := 86420
	leaseScavengeTimeUpdated := 214440
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkLeaseScavengeTime(name, networks, leaseScavengeTime, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "86420"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkLeaseScavengeTime(name, networks, leaseScavengeTimeUpdated, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "214440"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_LogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_logic_filter_rules"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	logicFilterRules := []map[string]any{
		{
			"filter": "example-option-filter-1",
			"type":   "Option",
		},
	}
	logicFilterRulesUpdated := []map[string]any{
		{
			"filter": "example-option-filter-2",
			"type":   "Option",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkLogicFilterRules(name, networks, logicFilterRules, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.filter", "example-option-filter-1"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.type", "Option"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkLogicFilterRules(name, networks, logicFilterRulesUpdated, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.filter", "example-option-filter-2"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.type", "Option"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_name"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	nameUpdated := acctest.RandomNameWithPrefix("shared_network")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkName(name, networks),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkName(nameUpdated, networks),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_Networks(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_networks"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	networksUpdated := []string{"${nios_ipam_network.test_network3.ref}",
		"${nios_ipam_network.test_network4.ref}"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkNetworks(name, networks),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "networks.#", "2"),
					resource.TestCheckResourceAttrPair(resourceName, "networks.0.ref", "nios_ipam_network.test_network1", "ref"),
					resource.TestCheckResourceAttrPair(resourceName, "networks.1.ref", "nios_ipam_network.test_network2", "ref"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkNetworks(name, networksUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "networks.#", "2"),
					resource.TestCheckResourceAttrPair(resourceName, "networks.0.ref", "nios_ipam_network.test_network3", "ref"),
					resource.TestCheckResourceAttrPair(resourceName, "networks.1.ref", "nios_ipam_network.test_network4", "ref"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_Nextserver(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_nextserver"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	nextServer := "nest-server1"
	nextServerUpdated := "next-server2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkNextserver(name, networks, nextServer, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", nextServer),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkNextserver(name, networks, nextServerUpdated, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", nextServerUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_Options(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_options"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	options := []map[string]any{
		{
			"name":  "domain-name",
			"num":   "15",
			"value": "aa.bb.com",
		},
		{
			"name":  "dhcp-lease-time",
			"num":   "51",
			"value": "72000",
		},
	}
	optionsUpdated := []map[string]any{
		{
			"name":  "time-offset",
			"value": "50",
		},
		{
			"name":  "dhcp-lease-time",
			"num":   "51",
			"value": "82000",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkOptions(name, networks, options, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "domain-name"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "aa.bb.com"),
					resource.TestCheckResourceAttr(resourceName, "options.1.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr(resourceName, "options.1.value", "72000"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkOptions(name, networks, optionsUpdated, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "time-offset"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "50"),
					resource.TestCheckResourceAttr(resourceName, "options.1.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr(resourceName, "options.1.value", "82000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_PxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_pxe_lease_time"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	pxeLeaseTime := 3600
	pxeLeaseTimeUpdated := 7200

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkPxeLeaseTime(name, networks, pxeLeaseTime, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "3600"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkPxeLeaseTime(name, networks, pxeLeaseTimeUpdated, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "7200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_UpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_update_dns_on_lease_renewal"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkUpdateDnsOnLeaseRenewal(name, networks, true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkUpdateDnsOnLeaseRenewal(name, networks, false, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_UseAuthority(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_use_authority"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkUseAuthority(name, networks, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_authority", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkUseAuthority(name, networks, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_authority", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_UseBootfile(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_use_bootfile"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkUseBootfile(name, networks, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkUseBootfile(name, networks, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_UseBootserver(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_use_bootserver"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkUseBootserver(name, networks, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkUseBootserver(name, networks, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_UseDdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_use_ddns_generate_hostname"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkUseDdnsGenerateHostname(name, networks, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkUseDdnsGenerateHostname(name, networks, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_UseDdnsTtl(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_use_ddns_ttl"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkUseDdnsTtl(name, networks, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkUseDdnsTtl(name, networks, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_UseDdnsUpdateFixedAddresses(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_use_ddns_update_fixed_addresses"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkUseDdnsUpdateFixedAddresses(name, networks, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_update_fixed_addresses", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkUseDdnsUpdateFixedAddresses(name, networks, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_update_fixed_addresses", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_UseDdnsUseOption81(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_use_ddns_use_option81"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkUseDdnsUseOption81(name, networks, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkUseDdnsUseOption81(name, networks, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_UseDenyBootp(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_use_deny_bootp"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkUseDenyBootp(name, networks, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkUseDenyBootp(name, networks, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_UseEnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_use_enable_ddns"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkUseEnableDdns(name, networks, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkUseEnableDdns(name, networks, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_UseIgnoreClientIdentifier(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_use_ignore_client_identifier"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkUseIgnoreClientIdentifier(name, networks, true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_client_identifier", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkUseIgnoreClientIdentifier(name, networks, false, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_client_identifier", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_UseIgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_use_ignore_dhcp_option_list_request"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkUseIgnoreDhcpOptionListRequest(name, networks, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkUseIgnoreDhcpOptionListRequest(name, networks, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_UseIgnoreId(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_use_ignore_id"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkUseIgnoreId(name, networks, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_id", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkUseIgnoreId(name, networks, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_id", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_UseLeaseScavengeTime(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_use_lease_scavenge_time"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkUseLeaseScavengeTime(name, networks, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkUseLeaseScavengeTime(name, networks, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_UseLogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_use_logic_filter_rules"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkUseLogicFilterRules(name, networks, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkUseLogicFilterRules(name, networks, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_UseNextserver(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_use_nextserver"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkUseNextserver(name, networks, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkUseNextserver(name, networks, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_UseOptions(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_use_options"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkUseOptions(name, networks, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkUseOptions(name, networks, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_UsePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_use_pxe_lease_time"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkUsePxeLeaseTime(name, networks, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkUsePxeLeaseTime(name, networks, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_UseUpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_dhcp_shared_network.test_use_update_dns_on_lease_renewal"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("shared_network")
	networks := []string{"${nios_ipam_network.test_network1.ref}",
		"${nios_ipam_network.test_network2.ref}"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkUseUpdateDnsOnLeaseRenewal(name, networks, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkUseUpdateDnsOnLeaseRenewal(name, networks, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckSharednetworkExists(ctx context.Context, resourceName string, v *dhcp.Sharednetwork) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			SharednetworkAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForSharednetwork).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetSharednetworkResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetSharednetworkResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckSharednetworkDestroy(ctx context.Context, v *dhcp.Sharednetwork) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			SharednetworkAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForSharednetwork).
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

func testAccCheckSharednetworkDisappears(ctx context.Context, v *dhcp.Sharednetwork) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			SharednetworkAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccBaseWithNetworks(network1, network2 string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "test_network1" {
	network = %q
}	

resource "nios_ipam_network" "test_network2" {
	network = %q	
}
`, network1, network2)
}

func testAccSharednetworkBasicConfig(name string, networks []string) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test" {
    name     = %q
    networks = %s
}`, name, networksStr)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.1.0.0/24", "201.2.0.0/24"), config}, "\n")
}

func formatNetworksToHCL(networks []string) string {
	networksStr := "["
	for i, network := range networks {
		if i > 0 {
			networksStr += ","
		}
		networksStr += fmt.Sprintf(`
        {
            ref = %q
        }`, network)
	}
	networksStr += "]"
	return networksStr
}

func testAccSharednetworkAuthority(name string, networks []string, authority bool, useAuthority bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_authority" {
   name = %q
   networks = %s
   authority = %t
   use_authority = %t
}
`, name, networksStr, authority, useAuthority)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.3.0.0/24", "201.4.0.0/24"), config}, "\n")
}

func testAccSharednetworkBootfile(name string, networks []string, bootfile string, useBootFile bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_bootfile" {
   name = %q
   networks = %s
   bootfile = %q
   use_bootfile = %t
}
`, name, networksStr, bootfile, useBootFile)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.5.0.0/24", "201.6.0.0/24"), config}, "\n")
}

func testAccSharednetworkBootserver(name string, networks []string, bootserver string, useBootServer bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_bootserver" {
   name = %q
   networks = %s
   bootserver = %q
   use_bootserver = %t
}
`, name, networksStr, bootserver, useBootServer)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.7.0.0/24", "201.8.0.0/24"), config}, "\n")
}

func testAccSharednetworkComment(name string, networks []string, comment string) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_comment" {
   name = %q
   networks = %s
   comment = %q
}
`, name, networksStr, comment)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.9.0.0/24", "201.10.0.0/24"), config}, "\n")
}

func testAccSharednetworkDdnsGenerateHostname(name string, networks []string, ddnsGenerateHostname, useDdnsGenerateHostName bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_ddns_generate_hostname" {
   name = %q
   networks = %s
   ddns_generate_hostname = %t
   use_ddns_generate_hostname = %t
}
`, name, networksStr, ddnsGenerateHostname, useDdnsGenerateHostName)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.11.0.0/24", "201.12.0.0/24"), config}, "\n")
}

func testAccSharednetworkDdnsServerAlwaysUpdates(name string, networks []string, ddnsServerAlwaysUpdates bool, ddnsUseOption18, useDdnsUseOption18 bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_ddns_server_always_updates" {
   name = %q
   networks = %s
   ddns_server_always_updates = %t
   ddns_use_option81 = %t
   use_ddns_use_option81 = %t
}
`, name, networksStr, ddnsServerAlwaysUpdates, ddnsUseOption18, useDdnsUseOption18)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.13.0.0/24", "201.14.0.0/24"), config}, "\n")
}

func testAccSharednetworkDdnsTtl(name string, networks []string, ddnsTtl int, useDdnsTtl bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_ddns_ttl" {
   name = %q
   networks = %s
   ddns_ttl = %d
   use_ddns_ttl = %t
}
`, name, networksStr, ddnsTtl, useDdnsTtl)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.15.0.0/24", "201.16.0.0/24"), config}, "\n")
}

func testAccSharednetworkDdnsUpdateFixedAddresses(name string, networks []string, ddnsUpdateFixedAddresses, useDdnsUpdateFixedAddresses bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_ddns_update_fixed_addresses" {
   name = %q
   networks = %s
   ddns_update_fixed_addresses = %t
   use_ddns_update_fixed_addresses = %t
}
`, name, networksStr, ddnsUpdateFixedAddresses, useDdnsUpdateFixedAddresses)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.17.0.0/24", "201.18.0.0/24"), config}, "\n")
}

func testAccSharednetworkDdnsUseOption81(name string, networks []string, ddnsUseOption81, useDdnsUseOption81 bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_ddns_use_option81" {
   name = %q
   networks = %s
   ddns_use_option81 = %t
   use_ddns_use_option81 = %t
}
`, name, networksStr, ddnsUseOption81, useDdnsUseOption81)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.19.0.0/24", "201.20.0.0/24"), config}, "\n")
}

func testAccSharednetworkDenyBootp(name string, networks []string, denyBootp, useDenyBootp bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_deny_bootp" {
   name = %q
   networks = %s
   deny_bootp = %t
   use_deny_bootp = %t
}
`, name, networksStr, denyBootp, useDenyBootp)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.21.0.0/24", "201.22.0.0/24"), config}, "\n")
}

func testAccSharednetworkDisable(name string, networks []string, disable bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_disable" {
   name = %q
   networks = %s
   disable = %t
}
`, name, networksStr, disable)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.23.0.0/24", "201.24.0.0/24"), config}, "\n")
}

func testAccSharednetworkEnableDdns(name string, networks []string, enableDdns, useEnableDdns bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_enable_ddns" {
   name = %q
   networks = %s
   enable_ddns = %t
   use_enable_ddns = %t
}
`, name, networksStr, enableDdns, useEnableDdns)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.25.0.0/24", "201.26.0.0/24"), config}, "\n")
}

func testAccSharednetworkEnablePxeLeaseTime(name string, networks []string, enablePxeLeaseTime bool, usePxeLeaseTime bool, pxeLeaseTime int) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_enable_pxe_lease_time" {
   name = %q
   networks = %s
   enable_pxe_lease_time = %t
   use_pxe_lease_time = %t
   pxe_lease_time = %d
}
`, name, networksStr, enablePxeLeaseTime, usePxeLeaseTime, pxeLeaseTime)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.27.0.0/24", "201.28.0.0/24"), config}, "\n")
}

func testAccSharednetworkExtAttrs(name string, networks []string, extAttrs map[string]string) string {
	networksStr := formatNetworksToHCL(networks)
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_extattrs" {
   name = %q
   networks = %s
   extattrs = %s
}
`, name, networksStr, extattrsStr)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.29.0.0/24", "201.30.0.0/24"), config}, "\n")
}

func testAccSharednetworkIgnoreClientIdentifier(name string, networks []string, ignoreClientIdentifier, useIgnoreClientIdentifier bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_ignore_client_identifier" {
   name = %q
   networks = %s
   ignore_client_identifier = %t
   use_ignore_client_identifier = %t
   use_ignore_id = true
}
`, name, networksStr, ignoreClientIdentifier, useIgnoreClientIdentifier)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.31.0.0/24", "201.32.0.0/24"), config}, "\n")
}

func testAccSharednetworkIgnoreDhcpOptionListRequest(name string, networks []string, ignoreDhcpOptionListRequest, useIgnoreDhcpOptionListRequest bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_ignore_dhcp_option_list_request" {
   name = %q
   networks = %s
   ignore_dhcp_option_list_request = %t
   use_ignore_dhcp_option_list_request = %t
}
`, name, networksStr, ignoreDhcpOptionListRequest, useIgnoreDhcpOptionListRequest)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.33.0.0/24", "201.34.0.0/24"), config}, "\n")
}

func testAccSharednetworkIgnoreId(name string, networks []string, ignoreId string, useIgnoreId bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_ignore_id" {
   name = %q
   networks = %s
   ignore_id = %q
   use_ignore_id = %t
}
`, name, networksStr, ignoreId, useIgnoreId)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.35.0.0/24", "201.36.0.0/24"), config}, "\n")
}

func testAccSharednetworkIgnoreMacAddresses(name string, networks []string, ignoreMacAddresses []string) string {
	networksStr := formatNetworksToHCL(networks)
	ignoreMacStr := formatMacAddrToHCL(ignoreMacAddresses)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_ignore_mac_addresses" {
   name = %q
   networks = %s
   ignore_mac_addresses = %s
}
`, name, networksStr, ignoreMacStr)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.37.0.0/24", "201.38.0.0/24"), config}, "\n")
}

func formatMacAddrToHCL(ignoreMacAddresses []string) string {
	macList := make([]string, len(ignoreMacAddresses))
	for i, mac := range ignoreMacAddresses {
		macList[i] = fmt.Sprintf("%q", mac)
	}
	return fmt.Sprintf("[%s]", strings.Join(macList, ", "))
}

func testAccSharednetworkLeaseScavengeTime(name string, networks []string, leaseScavengeTime int, useLeaseScavengeTime bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_lease_scavenge_time" {
   name = %q
   networks = %s
   lease_scavenge_time = %d
   use_lease_scavenge_time = %t
}
`, name, networksStr, leaseScavengeTime, useLeaseScavengeTime)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.39.0.0/24", "201.40.0.0/24"), config}, "\n")
}

func testAccSharednetworkLogicFilterRules(name string, networks []string, logicFilterRules []map[string]any, useLogicFilterRules bool) string {
	logicFilterRulesStr := convertSliceOfMapsToString(logicFilterRules)
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_logic_filter_rules" {
   name = %q
   networks = %s
   logic_filter_rules = %s
   use_logic_filter_rules = %t
}
`, name, networksStr, logicFilterRulesStr, useLogicFilterRules)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.41.0.0/24", "201.42.0.0/24"), config}, "\n")
}

func testAccSharednetworkName(name string, networks []string) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_name" {
   name = %q
   networks = %s
}
`, name, networksStr)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.43.0.0/24", "201.44.0.0/24"), config}, "\n")
}

func testAccSharednetworkNetworks(name string, networks []string) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_networks" {
   name = %q
   networks = %s
}
`, name, networksStr)
	network3and4 := testAccBaseWithNetworks("202.45.0.0/24", "202.46.0.0/24")
	network3and4Replace1 := strings.Replace(network3and4, "test_network1", "test_network3", 1)
	network3and4Replace2 := strings.Replace(network3and4Replace1, "test_network2", "test_network4", 1)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.45.0.0/24", "201.46.0.0/24"), network3and4Replace2, config}, "\n")
}

func testAccSharednetworkNextserver(name string, networks []string, nextserver string, useNextserver bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_nextserver" {
   name = %q
   networks = %s
   nextserver = %q
   use_nextserver = %t
}
`, name, networksStr, nextserver, useNextserver)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.47.0.0/24", "201.48.0.0/24"), config}, "\n")
}

func testAccSharednetworkOptions(name string, networks []string, options []map[string]any, useOptions bool) string {
	networksStr := formatNetworksToHCL(networks)
	optionsStr := convertSliceOfMapsToString(options)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_options" {
   name = %q
   networks = %s
   options = %s
   use_options = %t
}
`, name, networksStr, optionsStr, useOptions)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.49.0.0/24", "201.50.0.0/24"), config}, "\n")
}

func testAccSharednetworkPxeLeaseTime(name string, networks []string, pxeLeaseTime int, usePxeLeaseTime bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_pxe_lease_time" {
   name = %q
   networks = %s
   pxe_lease_time = %d
   use_pxe_lease_time = %t
}
`, name, networksStr, pxeLeaseTime, usePxeLeaseTime)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.51.0.0/24", "201.52.0.0/24"), config}, "\n")
}

func testAccSharednetworkUpdateDnsOnLeaseRenewal(name string, networks []string, updateDnsOnLeaseRenewal, useUpdateDnsOnLeaseRenewal bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_update_dns_on_lease_renewal" {
   name = %q
   networks = %s
   update_dns_on_lease_renewal = %t
   use_update_dns_on_lease_renewal = %t
}
`, name, networksStr, updateDnsOnLeaseRenewal, useUpdateDnsOnLeaseRenewal)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.53.0.0/24", "201.54.0.0/24"), config}, "\n")
}

func testAccSharednetworkUseAuthority(name string, networks []string, useAuthority bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_use_authority" {
   name = %q
   networks = %s
   use_authority = %t
}
`, name, networksStr, useAuthority)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.55.0.0/24", "201.56.0.0/24"), config}, "\n")
}

func testAccSharednetworkUseBootfile(name string, networks []string, useBootfile bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_use_bootfile" {
   name = %q
   networks = %s
   use_bootfile = %t
}
`, name, networksStr, useBootfile)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.57.0.0/24", "201.58.0.0/24"), config}, "\n")
}

func testAccSharednetworkUseBootserver(name string, networks []string, useBootserver bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_use_bootserver" {
   name = %q
   networks = %s
   use_bootserver = %t
}
`, name, networksStr, useBootserver)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.59.0.0/24", "201.60.0.0/24"), config}, "\n")
}

func testAccSharednetworkUseDdnsGenerateHostname(name string, networks []string, useDdnsGenerateHostname bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_use_ddns_generate_hostname" {
   name = %q
   networks = %s
   use_ddns_generate_hostname = %t
}
`, name, networksStr, useDdnsGenerateHostname)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.61.0.0/24", "201.62.0.0/24"), config}, "\n")
}

func testAccSharednetworkUseDdnsTtl(name string, networks []string, useDdnsTtl bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_use_ddns_ttl" {
   name = %q
   networks = %s
   use_ddns_ttl = %t
}
`, name, networksStr, useDdnsTtl)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.63.0.0/24", "201.64.0.0/24"), config}, "\n")
}

func testAccSharednetworkUseDdnsUpdateFixedAddresses(name string, networks []string, useDdnsUpdateFixedAddresses bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_use_ddns_update_fixed_addresses" {
   name = %q
   networks = %s
   use_ddns_update_fixed_addresses = %t
}
`, name, networksStr, useDdnsUpdateFixedAddresses)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.65.0.0/24", "201.66.0.0/24"), config}, "\n")
}

func testAccSharednetworkUseDdnsUseOption81(name string, networks []string, useDdnsUseOption81 bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_use_ddns_use_option81" {
   name = %q
   networks = %s
   use_ddns_use_option81 = %t
}
`, name, networksStr, useDdnsUseOption81)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.67.0.0/24", "201.68.0.0/24"), config}, "\n")
}

func testAccSharednetworkUseDenyBootp(name string, networks []string, useDenyBootp bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_use_deny_bootp" {
   name = %q
   networks = %s
   use_deny_bootp = %t
}
`, name, networksStr, useDenyBootp)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.69.0.0/24", "201.70.0.0/24"), config}, "\n")
}

func testAccSharednetworkUseEnableDdns(name string, networks []string, useEnableDdns bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_use_enable_ddns" {
   name = %q
   networks = %s
   use_enable_ddns = %t
}
`, name, networksStr, useEnableDdns)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.71.0.0/24", "201.72.0.0/24"), config}, "\n")
}

func testAccSharednetworkUseIgnoreClientIdentifier(name string, networks []string, useIgnoreClientIdentifier bool, ignoreClientIdentifier bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_use_ignore_client_identifier" {
   name = %q
   networks = %s
   use_ignore_client_identifier = %t
   ignore_client_identifier = %t
}
`, name, networksStr, useIgnoreClientIdentifier, ignoreClientIdentifier)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.73.0.0/24", "201.74.0.0/24"), config}, "\n")
}

func testAccSharednetworkUseIgnoreDhcpOptionListRequest(name string, networks []string, useIgnoreDhcpOptionListRequest bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_use_ignore_dhcp_option_list_request" {
   name = %q
   networks = %s
   use_ignore_dhcp_option_list_request = %t
}
`, name, networksStr, useIgnoreDhcpOptionListRequest)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.75.0.0/24", "201.76.0.0/24"), config}, "\n")
}

func testAccSharednetworkUseIgnoreId(name string, networks []string, useIgnoreId bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_use_ignore_id" {
   name = %q
   networks = %s
   use_ignore_id = %t
}
`, name, networksStr, useIgnoreId)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.77.0.0/24", "201.78.0.0/24"), config}, "\n")
}

func testAccSharednetworkUseLeaseScavengeTime(name string, networks []string, useLeaseScavengeTime bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_use_lease_scavenge_time" {
   name = %q
   networks = %s
   use_lease_scavenge_time = %t
}
`, name, networksStr, useLeaseScavengeTime)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.79.0.0/24", "201.80.0.0/24"), config}, "\n")
}

func testAccSharednetworkUseLogicFilterRules(name string, networks []string, useLogicFilterRules bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_use_logic_filter_rules" {
   name = %q
   networks = %s
   use_logic_filter_rules = %t
}
`, name, networksStr, useLogicFilterRules)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.81.0.0/24", "201.82.0.0/24"), config}, "\n")
}

func testAccSharednetworkUseNextserver(name string, networks []string, useNextserver bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_use_nextserver" {
   name = %q
   networks = %s
   use_nextserver = %t
}
`, name, networksStr, useNextserver)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.83.0.0/24", "201.84.0.0/24"), config}, "\n")
}

func testAccSharednetworkUseOptions(name string, networks []string, useOptions bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_use_options" {
   name = %q
   networks = %s
   use_options = %t
}
`, name, networksStr, useOptions)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.85.0.0/24", "201.86.0.0/24"), config}, "\n")
}

func testAccSharednetworkUsePxeLeaseTime(name string, networks []string, usePxeLeaseTime bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_use_pxe_lease_time" {
   name = %q
   networks = %s
   use_pxe_lease_time = %t
}
`, name, networksStr, usePxeLeaseTime)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.87.0.0/24", "201.88.0.0/24"), config}, "\n")
}

func testAccSharednetworkUseUpdateDnsOnLeaseRenewal(name string, networks []string, useUpdateDnsOnLeaseRenewal bool) string {
	networksStr := formatNetworksToHCL(networks)
	config := fmt.Sprintf(`
resource "nios_dhcp_shared_network" "test_use_update_dns_on_lease_renewal" {
   name = %q
   networks = %s
   use_update_dns_on_lease_renewal = %t
}
`, name, networksStr, useUpdateDnsOnLeaseRenewal)
	return strings.Join([]string{testAccBaseWithNetworks(
		"201.89.0.0/24", "201.90.0.0/24"), config}, "\n")
}

func convertSliceOfMapsToString(maps []map[string]any) string {
	mapsStr := "[\n"
	for _, obj := range maps {
		mapsStr += "  {\n"
		for k, v := range obj {
			if strVal, ok := v.(string); ok {
				mapsStr += fmt.Sprintf("    %s = %q\n", k, strVal) // Enclose string values in quotes
			} else {
				mapsStr += fmt.Sprintf("    %s = %v\n", k, v)
			}
		}
		mapsStr += "  },\n"
	}
	mapsStr += "]"
	return mapsStr
}
