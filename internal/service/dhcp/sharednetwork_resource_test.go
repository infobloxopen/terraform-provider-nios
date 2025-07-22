package dhcp_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// TODO : Required parents for the execution of tests
// - logic_filter_rules (option_filter, option_logic_filter)

var readableAttributesForSharednetwork = "authority,bootfile,bootserver,comment,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,ddns_update_fixed_addresses,ddns_use_option81,deny_bootp,dhcp_utilization,dhcp_utilization_status,disable,dynamic_hosts,enable_ddns,enable_pxe_lease_time,extattrs,ignore_client_identifier,ignore_dhcp_option_list_request,ignore_id,ignore_mac_addresses,lease_scavenge_time,logic_filter_rules,ms_ad_user_data,name,network_view,networks,nextserver,options,pxe_lease_time,static_hosts,total_hosts,update_dns_on_lease_renewal,use_authority,use_bootfile,use_bootserver,use_ddns_generate_hostname,use_ddns_ttl,use_ddns_update_fixed_addresses,use_ddns_use_option81,use_deny_bootp,use_enable_ddns,use_ignore_client_identifier,use_ignore_dhcp_option_list_request,use_ignore_id,use_lease_scavenge_time,use_logic_filter_rules,use_nextserver,use_options,use_pxe_lease_time,use_update_dns_on_lease_renewal"

func TestAccSharednetworkResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_sharednetwork.test"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}

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
					resource.TestCheckResourceAttr(resourceName, "networks.0", networks[0]),
					resource.TestCheckResourceAttr(resourceName, "networks.1", networks[1]),
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
					resource.TestCheckResourceAttr(resourceName, "options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr(resourceName, "options.0.num", "51"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "43200"),
					resource.TestCheckResourceAttr(resourceName, "options.0.use_option", "false"),
					resource.TestCheckResourceAttr(resourceName, "options.0.vendor_class", "DHCP"),
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
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_sharednetwork.test"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}

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
	var resourceName = "nios_dhcp_sharednetwork.test_authority"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
	authority := true
	authorityUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkAuthority(name, networks, authority, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "authority", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkAuthority(name, networks, authorityUpdated, true),
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
	var resourceName = "nios_dhcp_sharednetwork.test_bootfile"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
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
	var resourceName = "nios_dhcp_sharednetwork.test_bootserver"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
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
	var resourceName = "nios_dhcp_sharednetwork.test_comment"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
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
	var resourceName = "nios_dhcp_sharednetwork.test_ddns_generate_hostname"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
	ddnsGenerateHostName := true
	ddnsGenerateHostNameUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkDdnsGenerateHostname(name, networks, ddnsGenerateHostName, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkDdnsGenerateHostname(name, networks, ddnsGenerateHostNameUpdated, true),
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
	var resourceName = "nios_dhcp_sharednetwork.test_ddns_server_always_updates"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
	ddnServerAlwaysUpdate := false
	ddnServerAlwaysUpdateUpdated := true

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkDdnsServerAlwaysUpdates(name, networks, ddnServerAlwaysUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkDdnsServerAlwaysUpdates(name, networks, ddnServerAlwaysUpdateUpdated),
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
	var resourceName = "nios_dhcp_sharednetwork.test_ddns_ttl"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
	ddnsTtl := int64(3600)
	ddnsTtlUpdated := int64(7200)

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
	var resourceName = "nios_dhcp_sharednetwork.test_ddns_update_fixed_addresses"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
	ddnsUpdateFixedaddress := true
	ddnsUpdateFixedaddressUpdated := false
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkDdnsUpdateFixedAddresses(name, networks, ddnsUpdateFixedaddress, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_update_fixed_addresses", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkDdnsUpdateFixedAddresses(name, networks, ddnsUpdateFixedaddressUpdated, true),
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
	var resourceName = "nios_dhcp_sharednetwork.test_ddns_use_option81"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
	ddnsUseOption81 := true
	ddnsUseOption81Updated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkDdnsUseOption81(name, networks, ddnsUseOption81, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkDdnsUseOption81(name, networks, ddnsUseOption81Updated, true),
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
	var resourceName = "nios_dhcp_sharednetwork.test_deny_bootp"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
	denyBootp := true
	denyBootpUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkDenyBootp(name, networks, denyBootp, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkDenyBootp(name, networks, denyBootpUpdated, true),
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
	var resourceName = "nios_dhcp_sharednetwork.test_disable"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
	disable := true
	disableUpdated := false
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkDisable(name, networks, disable),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkDisable(name, networks, disableUpdated),
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
	var resourceName = "nios_dhcp_sharednetwork.test_enable_ddns"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
	enableDdns := true
	enableDdnsUpdated := false
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkEnableDdns(name, networks, enableDdns, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkEnableDdns(name, networks, enableDdnsUpdated, true),
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
	var resourceName = "nios_dhcp_sharednetwork.test_enable_pxe_lease_time"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
	enablePxeLeaseTime := true
	enablePxeLeaseTimeUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkEnablePxeLeaseTime(name, networks, enablePxeLeaseTime),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkEnablePxeLeaseTime(name, networks, enablePxeLeaseTimeUpdated),
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
	var resourceName = "nios_dhcp_sharednetwork.test_extattrs"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkExtAttrs(name, networks, extAttrValue1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkExtAttrs(name, networks, extAttrValue2),
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
	var resourceName = "nios_dhcp_sharednetwork.test_ignore_client_identifier"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
	ignoreClientIdentifier := true
	ignoreClientIdentifierUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkIgnoreClientIdentifier(name, networks, ignoreClientIdentifier, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_client_identifier", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkIgnoreClientIdentifier(name, networks, ignoreClientIdentifierUpdated, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_client_identifier", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_IgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_dhcp_sharednetwork.test_ignore_dhcp_option_list_request"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
	idnoreDhcpListRequest := true
	ignoreDhcpListRequestUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkIgnoreDhcpOptionListRequest(name, networks, idnoreDhcpListRequest, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkIgnoreDhcpOptionListRequest(name, networks, ignoreDhcpListRequestUpdated, true),
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
	var resourceName = "nios_dhcp_sharednetwork.test_ignore_id"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
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
	var resourceName = "nios_dhcp_sharednetwork.test_ignore_mac_addresses"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
	ignoreMacAddresses := []string{"00:11:22:33:44:55", "66:77:88:99:AA:BB"}
	ignoreMacAddressesUpdated := []string{"AA:BB:CC:DD:EE:FF", "00:11:22:33:44:55"}

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
					resource.TestCheckTypeSetElemAttr(resourceName, "ignore_mac_addresses.0", "00:11:22:33:44:55"),
					resource.TestCheckTypeSetElemAttr(resourceName, "ignore_mac_addresses.1", "66:77:88:99:AA:BB"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkIgnoreMacAddresses(name, networks, ignoreMacAddressesUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_mac_addresses.#", "2"),
					resource.TestCheckTypeSetElemAttr(resourceName, "ignore_mac_addresses.0", "AA:BB:CC:DD:EE:FF"),
					resource.TestCheckTypeSetElemAttr(resourceName, "ignore_mac_addresses.1", "00:11:22:33:44:55"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_LeaseScavengeTime(t *testing.T) {
	var resourceName = "nios_dhcp_sharednetwork.test_lease_scavenge_time"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
	leaseScavengeTime := int64(3600)
	leaseScavengeTimeUpdated := int64(7200)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkLeaseScavengeTime(name, networks, leaseScavengeTime, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "3600"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkLeaseScavengeTime(name, networks, leaseScavengeTimeUpdated, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "7200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_LogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_sharednetwork.test_logic_filter_rules"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
	logicFilterRules := []map[string]any{
		{
			"filter": "option_filter",
			"type":   "Option",
		},
	}
	logicFilterRulesUpdated := []map[string]any{
		{
			"filter": "option_logic_filter",
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
					resource.TestCheckTypeSetElemAttr(resourceName, "logic_filter_rules.0.filter", "option_filter"),
					resource.TestCheckTypeSetElemAttr(resourceName, "logic_filter_rules.0.type", "Option"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkLogicFilterRules(name, networks, logicFilterRulesUpdated, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.#", "1"),
					resource.TestCheckTypeSetElemAttr(resourceName, "logic_filter_rules.0.filter", "option_logic_filter"),
					resource.TestCheckTypeSetElemAttr(resourceName, "logic_filter_rules.0.type", "Option"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_sharednetwork.test_name"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
	nameUpdated := acctest.RandomNameWithPrefix("sharednetwork")

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
	var resourceName = "nios_dhcp_sharednetwork.test_networks"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
	networksUpdated := []string{"network/ZG5zLm5ldHdvcmskMTUuMTQuMS4wLzI0LzA:15.14.1.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTYuMC4wLjAvMjQvMA:16.0.0.0/24/default"}

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
					resource.TestCheckResourceAttr(resourceName, "networks.0", "network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default"),
					resource.TestCheckResourceAttr(resourceName, "networks.1", "network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkNetworks(name, networksUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "networks.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "networks.0", "network/ZG5zLm5ldHdvcmskMTUuMTQuMS4wLzI0LzA:15.14.1.0/24/default"),
					resource.TestCheckResourceAttr(resourceName, "networks.1", "network/ZG5zLm5ldHdvcmskMTYuMC4wLjAvMjQvMA:16.0.0.0/24/default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_Nextserver(t *testing.T) {
	var resourceName = "nios_dhcp_sharednetwork.test_nextserver"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
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
	var resourceName = "nios_dhcp_sharednetwork.test_options"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
	options := []map[string]any{
		{
			"name":  "domain-name",
			"value": "aa.bb.com",
		},
		{
			"name":  "dhcp-lease-time",
			"value": "72000",
		},
	}
	optionsUpdated := []map[string]any{
		{
			"name":  "domain-name",
			"value": "cc.dd.com",
		},
		{
			"name":  "dhcp-lease-time",
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
					resource.TestCheckTypeSetElemAttr(resourceName, "options.0.name", "domain-name"),
					resource.TestCheckTypeSetElemAttr(resourceName, "options.0.value", "aa.bb.com"),
					resource.TestCheckTypeSetElemAttr(resourceName, "options.1.name", "dhcp-lease-time"),
					resource.TestCheckTypeSetElemAttr(resourceName, "options.1.value", "72000"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkOptions(name, networks, optionsUpdated, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "2"),
					resource.TestCheckTypeSetElemAttr(resourceName, "options.0.name", "domain-name"),
					resource.TestCheckTypeSetElemAttr(resourceName, "options.0.value", "cc.dd.com"),
					resource.TestCheckTypeSetElemAttr(resourceName, "options.1.name", "dhcp-lease-time"),
					resource.TestCheckTypeSetElemAttr(resourceName, "options.1.value", "82000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharednetworkResource_PxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_sharednetwork.test_pxe_lease_time"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
	pxeLeaseTime := int64(3600)
	pxeLeaseTimeUpdated := int64(7200)

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
	var resourceName = "nios_dhcp_sharednetwork.test_update_dns_on_lease_renewal"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}
	updateDnsOnLeaseRenewal := true
	updateDnsOnLeaseRenewalUpdated := false
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkUpdateDnsOnLeaseRenewal(name, networks, updateDnsOnLeaseRenewal, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkUpdateDnsOnLeaseRenewal(name, networks, updateDnsOnLeaseRenewalUpdated, true),
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
	var resourceName = "nios_dhcp_sharednetwork.test_use_authority"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}

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
	var resourceName = "nios_dhcp_sharednetwork.test_use_bootfile"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}

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
	var resourceName = "nios_dhcp_sharednetwork.test_use_bootserver"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}

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
	var resourceName = "nios_dhcp_sharednetwork.test_use_ddns_generate_hostname"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}

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
	var resourceName = "nios_dhcp_sharednetwork.test_use_ddns_ttl"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}

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
	var resourceName = "nios_dhcp_sharednetwork.test_use_ddns_update_fixed_addresses"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}

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
	var resourceName = "nios_dhcp_sharednetwork.test_use_ddns_use_option81"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}

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
	var resourceName = "nios_dhcp_sharednetwork.test_use_deny_bootp"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}

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
	var resourceName = "nios_dhcp_sharednetwork.test_use_enable_ddns"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}

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
	var resourceName = "nios_dhcp_sharednetwork.test_use_ignore_client_identifier"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharednetworkUseIgnoreClientIdentifier(name, networks, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharednetworkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_client_identifier", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharednetworkUseIgnoreClientIdentifier(name, networks, false),
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
	var resourceName = "nios_dhcp_sharednetwork.test_use_ignore_dhcp_option_list_request"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}

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
	var resourceName = "nios_dhcp_sharednetwork.test_use_ignore_id"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}

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
	var resourceName = "nios_dhcp_sharednetwork.test_use_lease_scavenge_time"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}

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
	var resourceName = "nios_dhcp_sharednetwork.test_use_logic_filter_rules"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}

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
	var resourceName = "nios_dhcp_sharednetwork.test_use_nextserver"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}

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
	var resourceName = "nios_dhcp_sharednetwork.test_use_options"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}

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
	var resourceName = "nios_dhcp_sharednetwork.test_use_pxe_lease_time"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}

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
	var resourceName = "nios_dhcp_sharednetwork.test_use_update_dns_on_lease_renewal"
	var v dhcp.Sharednetwork
	name := acctest.RandomNameWithPrefix("sharednetwork")
	networks := []string{"network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
		"network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"}

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

func testAccSharednetworkBasicConfig(name string, networks []string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test" {
   name = %q
   networks = %s
}
`, name, networks)
}

func testAccSharednetworkAuthority(name string, networks []string, authority bool, useAuthority bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_authority" {
   name = %q
   networks = %s
   authority = %t
   use_authority = %t
}
`, name, networks, authority, useAuthority)
}

func testAccSharednetworkBootfile(name string, networks []string, bootfile string, useBootFile bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_bootfile" {
   name = %q
   networks = %s
   bootfile = %q
   use_bootfile = %t
}
`, name, networks, bootfile, useBootFile)
}

func testAccSharednetworkBootserver(name string, networks []string, bootserver string, useBootServer bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_bootserver" {
   name = %q
   networks = %s
   bootserver = %q
   use_bootserver = %t
}
`, name, networks, bootserver, useBootServer)
}

func testAccSharednetworkComment(name string, networks []string, comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_comment" {
   name = %q
   networks = %s
   comment = %q
}
`, name, networks, comment)
}

func testAccSharednetworkDdnsGenerateHostname(name string, networks []string, ddnsGenerateHostname, useDdnsGenerateHostName bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_ddns_generate_hostname" {
   name = %q
   networks = %s
   ddns_generate_hostname = %t
   use_ddns_generate_hostname = %t
}
`, name, networks, ddnsGenerateHostname, useDdnsGenerateHostName)
}

func testAccSharednetworkDdnsServerAlwaysUpdates(name string, networks []string, ddnsServerAlwaysUpdates bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_ddns_server_always_updates" {
   name = %q
   networks = %s
   ddns_server_always_updates = %t
}
`, name, networks, ddnsServerAlwaysUpdates)
}

func testAccSharednetworkDdnsTtl(name string, networks []string, ddnsTtl int64, useDdnsTtl bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_ddns_ttl" {
   name = %q
   networks = %s
   ddns_ttl = %d
   use_ddns_ttl = %t
}
`, name, networks, ddnsTtl, useDdnsTtl)
}

func testAccSharednetworkDdnsUpdateFixedAddresses(name string, networks []string, ddnsUpdateFixedAddresses, useDdnsUpdateFixedAddresses bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_ddns_update_fixed_addresses" {
   name = %q
   networks = %s
   ddns_update_fixed_addresses = %t
   use_ddns_update_fixed_addresses = %t
}
`, name, networks, ddnsUpdateFixedAddresses, useDdnsUpdateFixedAddresses)
}

func testAccSharednetworkDdnsUseOption81(name string, networks []string, ddnsUseOption81, useDdnsUseOption81 bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_ddns_use_option81" {
   name = %q
   networks = %s
   ddns_use_option81 = %t
   use_ddns_use_option81 = %t
}
`, name, networks, ddnsUseOption81, useDdnsUseOption81)
}

func testAccSharednetworkDenyBootp(name string, networks []string, denyBootp, useDenyBootp bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_deny_bootp" {
   name = %q
   networks = %s
   deny_bootp = %t
   use_deny_bootp = %t
}
`, name, networks, denyBootp, useDenyBootp)
}

func testAccSharednetworkDisable(name string, networks []string, disable bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_disable" {
   name = %q
   networks = %s
   disable = %t
}
`, name, networks, disable)
}

func testAccSharednetworkEnableDdns(name string, networks []string, enableDdns, useEnableDdns bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_enable_ddns" {
   name = %q
   networks = %s
   enable_ddns = %t
   use_enable_ddns = %t
}
`, name, networks, enableDdns, useEnableDdns)
}

func testAccSharednetworkEnablePxeLeaseTime(name string, networks []string, enablePxeLeaseTime bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_enable_pxe_lease_time" {
   name = %q
   networks = %s
   enable_pxe_lease_time = %t
}
`, name, networks, enablePxeLeaseTime)
}

func testAccSharednetworkExtAttrs(name string, networks []string, extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_extattrs" {
   name = %q
   networks = %s
   extattrs = %q
}
`, name, networks, extAttrs)
}

func testAccSharednetworkIgnoreClientIdentifier(name string, networks []string, ignoreClientIdentifier, useIgnoreClientIdentifier bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_ignore_client_identifier" {
   name = %q
   networks = %s
   ignore_client_identifier = %t
   use_ignore_client_identifier = %t
}
`, name, networks, ignoreClientIdentifier, useIgnoreClientIdentifier)
}

func testAccSharednetworkIgnoreDhcpOptionListRequest(name string, networks []string, ignoreDhcpOptionListRequest, useIgnoreDhcpOptionListRequest bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_ignore_dhcp_option_list_request" {
   name = %q
   networks = %s
   ignore_dhcp_option_list_request = %t
   use_ignore_dhcp_option_list_request = %t
}
`, name, networks, ignoreDhcpOptionListRequest, useIgnoreDhcpOptionListRequest)
}

func testAccSharednetworkIgnoreId(name string, networks []string, ignoreId string, useIgnoreId bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_ignore_id" {
   name = %q
   networks = %s
   ignore_id = %q
   use_ignore_id = %t
}
`, name, networks, ignoreId, useIgnoreId)
}

func testAccSharednetworkIgnoreMacAddresses(name string, networks []string, ignoreMacAddresses []string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_ignore_mac_addresses" {
   name = %q
   networks = %s
   ignore_mac_addresses = %s
}
`, name, networks, ignoreMacAddresses)
}

func testAccSharednetworkLeaseScavengeTime(name string, networks []string, leaseScavengeTime int64, useLeaseScavengeTime bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_lease_scavenge_time" {
   name = %q
   networks = %s
   lease_scavenge_time = %d
   use_lease_scavenge_time = %t
}
`, name, networks, leaseScavengeTime, useLeaseScavengeTime)
}

func testAccSharednetworkLogicFilterRules(name string, networks []string, logicFilterRules []map[string]any, useLogicFilterRules bool) string {
	logicFilterRulesStr := convertSliceOfMapsToString(logicFilterRules)
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_logic_filter_rules" {
   name = %q
   networks = %s
   logic_filter_rules = %q
   use_logic_filter_rules = %t
}
`, name, networks, logicFilterRulesStr, useLogicFilterRules)
}

func testAccSharednetworkName(name string, networks []string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_name" {
   name = %q
   networks = %s
}
`, name, networks)
}

func testAccSharednetworkNetworks(name string, networks []string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_networks" {
   name = %q
   networks = %q
}
`, name, networks)
}

func testAccSharednetworkNextserver(name string, networks []string, nextserver string, useNextserver bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_nextserver" {
   name = %q
   networks = %s
   nextserver = %q
   use_nextserver = %t
}
`, name, networks, nextserver, useNextserver)
}

func testAccSharednetworkOptions(name string, networks []string, options []map[string]any, useOptions bool) string {
	optionsStr := convertSliceOfMapsToString(options)
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_options" {
   name = %q
   networks = %s
   options = %q
   use_options = %t
}
`, name, networks, optionsStr, useOptions)
}

func testAccSharednetworkPxeLeaseTime(name string, networks []string, pxeLeaseTime int64, usePxeLeaseTime bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_pxe_lease_time" {
   name = %q
   networks = %s
   pxe_lease_time = %d
   use_pxe_lease_time = %t
}
`, name, networks, pxeLeaseTime, usePxeLeaseTime)
}

func testAccSharednetworkUpdateDnsOnLeaseRenewal(name string, networks []string, updateDnsOnLeaseRenewal, useUpdateDnsOnLeaseRenewal bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_update_dns_on_lease_renewal" {
   name = %q
   networks = %s
   update_dns_on_lease_renewal = %t
   use_update_dns_on_lease_renewal = %t
}
`, name, networks, updateDnsOnLeaseRenewal, useUpdateDnsOnLeaseRenewal)
}

func testAccSharednetworkUseAuthority(name string, networks []string, useAuthority bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_use_authority" {
   name = %q
   networks = %s
   use_authority = %t
}
`, name, networks, useAuthority)
}

func testAccSharednetworkUseBootfile(name string, networks []string, useBootfile bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_use_bootfile" {
   name = %q
   networks = %s
   use_bootfile = %t
}
`, name, networks, useBootfile)
}

func testAccSharednetworkUseBootserver(name string, networks []string, useBootserver bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_use_bootserver" {
   name = %q
   networks = %s
   use_bootserver = %t
}
`, name, networks, useBootserver)
}

func testAccSharednetworkUseDdnsGenerateHostname(name string, networks []string, useDdnsGenerateHostname bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_use_ddns_generate_hostname" {
   name = %q
   networks = %s
   use_ddns_generate_hostname = %t
}
`, name, networks, useDdnsGenerateHostname)
}

func testAccSharednetworkUseDdnsTtl(name string, networks []string, useDdnsTtl bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_use_ddns_ttl" {
   name = %q
   networks = %s
   use_ddns_ttl = %t
}
`, name, networks, useDdnsTtl)
}

func testAccSharednetworkUseDdnsUpdateFixedAddresses(name string, networks []string, useDdnsUpdateFixedAddresses bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_use_ddns_update_fixed_addresses" {
   name = %q
   networks = %s
   use_ddns_update_fixed_addresses = %t
}
`, name, networks, useDdnsUpdateFixedAddresses)
}

func testAccSharednetworkUseDdnsUseOption81(name string, networks []string, useDdnsUseOption81 bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_use_ddns_use_option81" {
   name = %q
   networks = %s
   use_ddns_use_option81 = %t
}
`, name, networks, useDdnsUseOption81)
}

func testAccSharednetworkUseDenyBootp(name string, networks []string, useDenyBootp bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_use_deny_bootp" {
   name = %q
   networks = %s
   use_deny_bootp = %t
}
`, name, networks, useDenyBootp)
}

func testAccSharednetworkUseEnableDdns(name string, networks []string, useEnableDdns bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_use_enable_ddns" {
   name = %q
   networks = %s
   use_enable_ddns = %t
}
`, name, networks, useEnableDdns)
}

func testAccSharednetworkUseIgnoreClientIdentifier(name string, networks []string, useIgnoreClientIdentifier bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_use_ignore_client_identifier" {
   name = %q
   networks = %s
   use_ignore_client_identifier = %t
}
`, name, networks, useIgnoreClientIdentifier)
}

func testAccSharednetworkUseIgnoreDhcpOptionListRequest(name string, networks []string, useIgnoreDhcpOptionListRequest bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_use_ignore_dhcp_option_list_request" {
   name = %q
   networks = %s
   use_ignore_dhcp_option_list_request = %t
}
`, name, networks, useIgnoreDhcpOptionListRequest)
}

func testAccSharednetworkUseIgnoreId(name string, networks []string, useIgnoreId bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_use_ignore_id" {
   name = %q
   networks = %s
   use_ignore_id = %t
}
`, name, networks, useIgnoreId)
}

func testAccSharednetworkUseLeaseScavengeTime(name string, networks []string, useLeaseScavengeTime bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_use_lease_scavenge_time" {
   name = %q
   networks = %s
   use_lease_scavenge_time = %t
}
`, name, networks, useLeaseScavengeTime)
}

func testAccSharednetworkUseLogicFilterRules(name string, networks []string, useLogicFilterRules bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_use_logic_filter_rules" {
   name = %q
   networks = %s
   use_logic_filter_rules = %t
}
`, name, networks, useLogicFilterRules)
}

func testAccSharednetworkUseNextserver(name string, networks []string, useNextserver bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_use_nextserver" {
   name = %q
   networks = %s
   use_nextserver = %t
}
`, name, networks, useNextserver)
}

func testAccSharednetworkUseOptions(name string, networks []string, useOptions bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_use_options" {
   name = %q
   networks = %s
   use_options = %t
}
`, name, networks, useOptions)
}

func testAccSharednetworkUsePxeLeaseTime(name string, networks []string, usePxeLeaseTime bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_use_pxe_lease_time" {
   name = %q
   networks = %s
   use_pxe_lease_time = %t
}
`, name, networks, usePxeLeaseTime)
}

func testAccSharednetworkUseUpdateDnsOnLeaseRenewal(name string, networks []string, useUpdateDnsOnLeaseRenewal bool) string {
	return fmt.Sprintf(`
resource "nios_dhcp_sharednetwork" "test_use_update_dns_on_lease_renewal" {
   name = %q
   networks = %s
   use_update_dns_on_lease_renewal = %t
}
`, name, networks, useUpdateDnsOnLeaseRenewal)
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
