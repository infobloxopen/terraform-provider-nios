package grid_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

func TestRestartServicesAction_DNS(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Forces restart of DNS Service in GROUPED mode after DNS View creation
			{
				Config: testaccrestartservicesactionDns(),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestRestartServicesAction_DHCPv4(t *testing.T) {
	startAddr := "10.0.0.111"
	endAddr := "10.0.0.121"
	memberName := utils.GetNIOSGridMemberHostName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Forces restart of DHCPv4 Service in SEQUENTIAL mode after DHCP Range creation
			{
				Config: testAccRestartServicesAction_DHCPv4(startAddr, endAddr, "initial comment", "infoblox.member2", "admin"),
			},
			// Forces restart of DHCPv4 Service in SEQUENTIAL mode after DHCP Range comment update
			{
				Config: testAccRestartServicesAction_DHCPv4(startAddr, endAddr, "updated comment", memberName, "admin"),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestRestartServicesAction_DHCPv6(t *testing.T) {
	ipv6Network := "2001:db8:abcd:1241::/64"
	ipv6addr := "2001:db8:abcd:1241::1"
	networkView := acctest.RandomNameWithPrefix("network-view")
	duid := "00:01:00:01:1d:2b:3c:4d:00:0c:19:ab:cd:ef"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Restarts DHCPv6 Service in SIMULTANEOUS mode after IPv6 Fixed Address creation (RESTART_IF_NEEDED)
			{
				Config: testAccRestartServicesAction_DHCPv6(ipv6addr, duid, networkView, ipv6Network),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testaccrestartservicesactionDns() string {
	return `
action "nios_restart_services" "dns" {
  config {
    services       = ["DNS"]
    restart_option = "FORCE_RESTART"
    mode           = "GROUPED"
  }
}

resource "nios_dns_view" "test" {
  name = "new_view"

  lifecycle {
    action_trigger {
      events  = [after_create]
      actions = [action.nios_restart_services.dns]
    }
  }
}
`
}

func testAccRestartServicesAction_DHCPv4(startAddr, endAddr, comment, member, userName string) string {
	return fmt.Sprintf(`
action "nios_restart_services" "dhcpv4" {
  config {
    services       = ["DHCPV4"]
    restart_option = "FORCE_RESTART"
    members        = [%q]
    user_name      = %q
    mode           = "SEQUENTIAL"
  }
}

resource "nios_dhcp_range" "test" {
  start_addr = %q
  end_addr   = %q
  comment    = %q

  lifecycle {
    action_trigger {
      events  = [after_update]
      actions = [action.nios_restart_services.dhcpv4]
    }
  }
}
`, member, userName, startAddr, endAddr, comment)
}

func testAccRestartServicesAction_DHCPv6(ipv6addr, duid, networkView, ipv6Network string) string {
	return fmt.Sprintf(`
action "nios_restart_services" "dhcpv6" {
  config {
    services       = ["DHCPV6"]
    restart_option = "RESTART_IF_NEEDED"
    mode           = "SIMULTANEOUS"
  }
}

resource "nios_ipam_network_view" "parent_network_view" {
  name = %q
}

resource "nios_ipam_ipv6network" "test_ipv6_network" {
  network      = %q
  network_view = nios_ipam_network_view.parent_network_view.name
}

resource "nios_dhcp_ipv6fixedaddress" "test" {
  ipv6addr     = %q
  duid         = %q
  network      = nios_ipam_ipv6network.test_ipv6_network.network
  network_view = nios_ipam_network_view.parent_network_view.name

  lifecycle {
    action_trigger {
      events  = [after_create]
      actions = [action.nios_restart_services.dhcpv6]
    }
  }
}
`, networkView, ipv6Network, ipv6addr, duid)
}
