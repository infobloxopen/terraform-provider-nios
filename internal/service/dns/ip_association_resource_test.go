package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

func TestAccIPAssociationResource_basic(t *testing.T) {
	var resourceName = "nios_ip_association.association"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.10",
		},
	}
	mac := "12:00:43:fe:9a:8c"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAssociationBasicConfig(name, "default", mac, "true", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mac", mac),
					resource.TestCheckResourceAttr(resourceName, "configure_for_dhcp", "true"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "match_client", "DUID"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAssociationResource_macAssociation(t *testing.T) {
	var resourceName = "nios_ip_association.mac_association"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv6addr := []map[string]any{
		{
			"ipv6addr": "fd00:1234:5678::1",
		},
	}
	mac := "12:00:43:fe:9a:8d"
	macUpdated := "12:00:43:fe:9a:8f"
	matchClient := "MAC_ADDRESS"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAssociationMacConfig(name, "default", mac, "true", matchClient, ipv6addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mac", mac),
					resource.TestCheckResourceAttr(resourceName, "configure_for_dhcp", "true"),
					resource.TestCheckResourceAttr(resourceName, "match_client", matchClient),
				),
			},
			// Update and Read
			{
				Config: testAccIPAssociationMacConfig(name, "default", macUpdated, "true", matchClient, ipv6addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mac", macUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAssociationResource_duidAssociation(t *testing.T) {
	var resourceName = "nios_ip_association.duid_association"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv6addr := []map[string]any{
		{
			"ipv6addr": "fd00:1234:5678::12",
		},
	}
	mac := "12:00:43:fe:9a:3e"
	duid := "00:01:5f:3a:1b:2c:12:34:56:78:9a:bc"
	duidUpdated := "00:01:5f:3a:1b:2c:12:34:56:78:9a:bb"
	matchClient := "DUID"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAssociationDuidConfig(name, "default", mac, duid, "true", matchClient, ipv6addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "duid", duid),
					resource.TestCheckResourceAttr(resourceName, "configure_for_dhcp", "true"),
					resource.TestCheckResourceAttr(resourceName, "match_client", matchClient),
				),
			},
			// Update and Read
			{
				Config: testAccIPAssociationDuidConfig(name, "default", mac, duidUpdated, "true", matchClient, ipv6addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "duid", duidUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAssociationResource_matchClient(t *testing.T) {
	var resourceName = "nios_ip_association.match_client"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv6addr := []map[string]any{
		{
			"ipv6addr": "fd00:1234:5678::12",
		},
	}
	mac := "12:00:43:fe:9a:8e"
	duid := "00:01:5f:3a:1b:2c:12:34:56:78:9b:bc"
	matchClient := "DUID"
	matchClientUpdated := "MAC_ADDRESS"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAssociationMatchClientConfig(name, "default", mac, duid, "true", matchClient, ipv6addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "duid", duid),
					resource.TestCheckResourceAttr(resourceName, "configure_for_dhcp", "true"),
					resource.TestCheckResourceAttr(resourceName, "match_client", matchClient),
				),
			},
			// Update and Read
			{
				Config: testAccIPAssociationMatchClientConfig(name, "default", mac, duid, "true", matchClientUpdated, ipv6addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_client", matchClientUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIPAssociationResource_disableAssociation(t *testing.T) {
	var resourceName = "nios_ip_association.disable_association"
	var v dns.RecordHost

	name := acctest.RandomName() + ".example.com"
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.23",
		},
	}
	mac := "12:00:43:fe:9a:7e"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIPAssociationDisableConfig(name, "default", mac, "true", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "configure_for_dhcp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIPAssociationDisableConfig(name, "default", mac, "false", ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIPAllocationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "configure_for_dhcp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccIPAssociationBasicConfig(name, view, mac, configure_for_dhcp string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "allocation" {
	name = %q
	ipv4addrs = %s
	view = %q
}

resource "nios_ip_association" "association" {
	ref = nios_ip_allocation.allocation.ref
	mac = %q
	configure_for_dhcp = %q
	internal_id = nios_ip_allocation.allocation.internal_id
}
`, name, ipv4addrHCL, view, mac, configure_for_dhcp)
}

func testAccIPAssociationMacConfig(name, view, mac, configure_for_dhcp, match_client string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "allocation" {
	name = %q
	ipv6addrs = %s
	view = %q
}

resource "nios_ip_association" "mac_association" {
	ref = nios_ip_allocation.allocation.ref
	mac = %q
	configure_for_dhcp = %q
	match_client = %q
	internal_id = nios_ip_allocation.allocation.internal_id
}
`, name, ipv4addrHCL, view, mac, configure_for_dhcp, match_client)
}

func testAccIPAssociationDuidConfig(name, view, mac, duid, configure_for_dhcp, match_client string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "allocation" {
	name = %q
	ipv6addrs = %s
	view = %q
}

resource "nios_ip_association" "duid_association" {
	ref = nios_ip_allocation.allocation.ref
	mac = %q
	duid = %q
	configure_for_dhcp = %q
	match_client = %q
	internal_id = nios_ip_allocation.allocation.internal_id
}
`, name, ipv4addrHCL, view, mac, duid, configure_for_dhcp, match_client)
}

func testAccIPAssociationMatchClientConfig(name, view, mac, duid, configure_for_dhcp, match_client string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "allocation" {
	name = %q
	ipv6addrs = %s
	view = %q
}

resource "nios_ip_association" "match_client" {
	ref = nios_ip_allocation.allocation.ref
	mac = %q
	duid = %q
	configure_for_dhcp = %q
	match_client = %q
	internal_id = nios_ip_allocation.allocation.internal_id
}
`, name, ipv4addrHCL, view, mac, duid, configure_for_dhcp, match_client)
}

func testAccIPAssociationDisableConfig(name, view, mac, configure_for_dhcp string, ipv4addr []map[string]any) string {
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_ip_allocation" "allocation" {
	name = %q
	ipv4addrs = %s
	view = %q
}

resource "nios_ip_association" "disable_association" {
	ref = nios_ip_allocation.allocation.ref
	mac = %q
	configure_for_dhcp = %q
	internal_id = nios_ip_allocation.allocation.internal_id
}
`, name, ipv4addrHCL, view, mac, configure_for_dhcp)
}
