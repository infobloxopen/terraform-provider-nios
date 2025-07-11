package dns_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/acctest"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/utils"
)

//TODO : OBJECTS TO BE PRESENT IN GRID FOR TESTS
// - externalNsGroup = ensg1, ensg2
// - nsGroup = ns_group1, ns_group2

var readableAttributesForZoneForward = "address,comment,disable,disable_ns_generation,display_domain,dns_fqdn,extattrs,external_ns_group,forward_to,forwarders_only,forwarding_servers,fqdn,locked,locked_by,mask_prefix,ms_ad_integrated,ms_ddns_mode,ms_managed,ms_read_only,ms_sync_master_name,ns_group,parent,prefix,using_srg_associations,view,zone_format"

func TestAccZoneForwardResource_basic(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test"
	var v dns.ZoneForward
	fqdn := acctest.RandomNameWithPrefix("zone-forward") + ".example.com"
	externalNsGroup := "ensg1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardBasicConfig(fqdn, externalNsGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "fqdn", fqdn),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "disable_ns_generation", "false"),
					resource.TestCheckResourceAttr(resourceName, "forwarders_only", "false"),
					resource.TestCheckResourceAttr(resourceName, "locked", "false"),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_integrated", "false"),
					resource.TestCheckResourceAttr(resourceName, "ms_ddns_mode", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
					resource.TestCheckResourceAttr(resourceName, "zone_format", "FORWARD"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_disappears(t *testing.T) {
	resourceName := "nios_dns_zone_forward.test"
	var v dns.ZoneForward
	fqdn := acctest.RandomNameWithPrefix("zone-forward") + ".example.com"
	externalNsGroup := "ensg1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckZoneForwardDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneForwardBasicConfig(fqdn, externalNsGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					testAccCheckZoneForwardDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccZoneForwardResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_comment"
	var v dns.ZoneForward
	fqdn := acctest.RandomNameWithPrefix("zone-forward") + ".example.com"
	externalNsGroup := "ensg1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardComment(fqdn, externalNsGroup, "Zone forward comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Zone forward comment"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardComment(fqdn, externalNsGroup, "Zone forward comment updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Zone forward comment updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_disable"
	var v dns.ZoneForward
	fqdn := acctest.RandomNameWithPrefix("zone-forward") + ".example.com"
	externalNsGroup := "ensg1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardDisable(fqdn, externalNsGroup, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardDisable(fqdn, externalNsGroup, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_DisableNsGeneration(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_disable_ns_generation"
	var v dns.ZoneForward
	fqdn := acctest.RandomNameWithPrefix("zone-forward") + ".example.com"
	externalNsGroup := "ensg1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardDisableNsGeneration(fqdn, externalNsGroup, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_ns_generation", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardDisableNsGeneration(fqdn, externalNsGroup, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_ns_generation", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_extattrs"
	var v dns.ZoneForward
	fqdn := acctest.RandomNameWithPrefix("zone-forward") + ".example.com"
	externalNsGroup := "ensg1"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardExtAttrs(fqdn, externalNsGroup, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardExtAttrs(fqdn, externalNsGroup, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_ExternalNsGroup(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_external_ns_group"
	var v dns.ZoneForward
	fqdn := acctest.RandomNameWithPrefix("zone-forward") + ".example.com"
	externalNsGroup1 := "ensg1"
	externalNsGroup2 := "ensg2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardExternalNsGroup(fqdn, externalNsGroup1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_ns_group", externalNsGroup1),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardExternalNsGroup(fqdn, externalNsGroup2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_ns_group", externalNsGroup2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_ForwardTo(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_forward_to"
	var v dns.ZoneForward
	fqdn := acctest.RandomNameWithPrefix("zone-forward") + ".example.com"
	forwardTo1 := []map[string]string{
		{
			"name":    "example1.org",
			"address": "10.1.0.1",
		},
	}
	forwardTo2 := []map[string]string{
		{
			"name":    "example2.org",
			"address": "10.1.0.2",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardForwardTo(fqdn, forwardTo1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forward_to.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "forward_to.0.name", "example1.org"),
					resource.TestCheckResourceAttr(resourceName, "forward_to.0.address", "10.1.0.1"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardForwardTo(fqdn, forwardTo2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forward_to.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "forward_to.0.name", "example2.org"),
					resource.TestCheckResourceAttr(resourceName, "forward_to.0.address", "10.1.0.2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_ForwardersOnly(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_forwarders_only"
	var v dns.ZoneForward
	fqdn := acctest.RandomNameWithPrefix("zone-forward") + ".example.com"
	externalNsGroup := "ensg1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardForwardersOnly(fqdn, externalNsGroup, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarders_only", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardForwardersOnly(fqdn, externalNsGroup, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarders_only", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_ForwardingServers(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_forwarding_servers"
	var v dns.ZoneForward
	fqdn := acctest.RandomNameWithPrefix("zone-forward") + ".example.com"
	externalNsGroup := "ensg1"
	forwardingServer1 := []dns.ZoneForwardForwardingServers{
		{
			Name:                  utils.Ptr("infoblox.172_28_82_248"),
			ForwardersOnly:        utils.Ptr(true),
			UseOverrideForwarders: utils.Ptr(true),
			ForwardTo: []dns.ZoneforwardforwardingserversForwardTo{
				{
					Name:    utils.Ptr("example1.org"),
					Address: utils.Ptr("11.10.1.2"),
				},
			},
		},
	}
	forwardingServer2 := []dns.ZoneForwardForwardingServers{
		{
			Name:                  utils.Ptr("infoblox.172_28_82_248"),
			ForwardersOnly:        utils.Ptr(false),
			UseOverrideForwarders: utils.Ptr(false),
			ForwardTo: []dns.ZoneforwardforwardingserversForwardTo{
				{
					Name:    utils.Ptr("example22.org"),
					Address: utils.Ptr("11.10.11.22"),
				},
			},
		},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardForwardingServers(fqdn, externalNsGroup, forwardingServer1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarding_servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "forwarding_servers.0.name", "infoblox.172_28_82_248"),
					resource.TestCheckResourceAttr(resourceName, "forwarding_servers.0.forwarders_only", "true"),
					resource.TestCheckResourceAttr(resourceName, "forwarding_servers.0.use_override_forwarders", "true"),
					resource.TestCheckResourceAttr(resourceName, "forwarding_servers.0.forward_to.0.name", "example1.org"),
					resource.TestCheckResourceAttr(resourceName, "forwarding_servers.0.forward_to.0.address", "11.10.1.2"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardForwardingServers(fqdn, externalNsGroup, forwardingServer2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarding_servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "forwarding_servers.0.name", "infoblox.172_28_82_248"),
					resource.TestCheckResourceAttr(resourceName, "forwarding_servers.0.forwarders_only", "false"),
					resource.TestCheckResourceAttr(resourceName, "forwarding_servers.0.use_override_forwarders", "false"),
					resource.TestCheckResourceAttr(resourceName, "forwarding_servers.0.forward_to.0.name", "example22.org"),
					resource.TestCheckResourceAttr(resourceName, "forwarding_servers.0.forward_to.0.address", "11.10.11.22"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_Locked(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_locked"
	var v dns.ZoneForward
	fqdn := acctest.RandomNameWithPrefix("zone-forward") + ".example.com"
	externalNsGroup := "ensg1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardLocked(fqdn, externalNsGroup, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "locked", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardLocked(fqdn, externalNsGroup, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "locked", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_MsAdIntegrated(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_ms_ad_integrated"
	var v dns.ZoneForward
	fqdn := acctest.RandomNameWithPrefix("zone-forward") + ".example.com"
	externalNsGroup := "ensg1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardMsAdIntegrated(fqdn, externalNsGroup, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_integrated", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardMsAdIntegrated(fqdn, externalNsGroup, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_integrated", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_MsDdnsMode(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_ms_ddns_mode"
	var v dns.ZoneForward
	fqdn := acctest.RandomNameWithPrefix("zone-forward") + ".example.com"
	externalNsGroup := "ensg1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardMsDdnsMode(fqdn, externalNsGroup, "ANY"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ddns_mode", "ANY"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardMsDdnsMode(fqdn, externalNsGroup, "NONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ddns_mode", "NONE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_NsGroup(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_ns_group"
	var v dns.ZoneForward
	fqdn := acctest.RandomNameWithPrefix("zone-forward") + ".example.com"
	externalNsGroup := "ensg1"
	nsGroup1 := "ns_group1"
	nsGroup2 := "ns_group2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardNsGroup(fqdn, externalNsGroup, nsGroup1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ns_group", nsGroup1),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardNsGroup(fqdn, externalNsGroup, nsGroup2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ns_group", nsGroup2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_Prefix(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_prefix"
	var v dns.ZoneForward
	fqdn := acctest.RandomNameWithPrefix("zone-forward") + ".example.com"
	externalNsGroup := "ensg1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardPrefix(fqdn, externalNsGroup, "0-127"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", "0-127"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardPrefix(fqdn, externalNsGroup, "128/26"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", "128/26"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckZoneForwardExists(ctx context.Context, resourceName string, v *dns.ZoneForward) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			ZoneForwardAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForZoneForward).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetZoneForwardResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetZoneForwardResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckZoneForwardDestroy(ctx context.Context, v *dns.ZoneForward) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			ZoneForwardAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForZoneForward).
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

func testAccCheckZoneForwardDisappears(ctx context.Context, v *dns.ZoneForward) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			ZoneForwardAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccZoneForwardBasicConfig(fqdn, nsGroup string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test" {
	fqdn = %q
    external_ns_group = %q
}
`, fqdn, nsGroup)
}

func testAccZoneForwardComment(fqdn, externalNsGroup, comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_comment" {
   fqdn = %q
   external_ns_group = %q
   comment = %q
}
`, fqdn, externalNsGroup, comment)
}

func testAccZoneForwardDisable(fqdn, externalNsGroup string, disable bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_disable" {
   fqdn = %q
   external_ns_group = %q
   disable = %t
}
`, fqdn, externalNsGroup, disable)
}

func testAccZoneForwardDisableNsGeneration(fqdn, externalNsGroup string, disableNsGeneration bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_disable_ns_generation" {
   fqdn = %q
   external_ns_group = %q
   disable_ns_generation = %t 
}
`, fqdn, externalNsGroup, disableNsGeneration)
}

func testAccZoneForwardExtAttrs(fqdn, externalNsGroup string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_extattrs" {
   fqdn = %q
   external_ns_group = %q
   extattrs = %s
}
`, fqdn, externalNsGroup, extattrsStr)
}

func testAccZoneForwardExternalNsGroup(fqdn, externalNsGroup string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_external_ns_group" {
   fqdn = %q
   external_ns_group = %q
}
`, fqdn, externalNsGroup)
}

func testAccZoneForwardForwardTo(fqdn string, forwardTo []map[string]string) string {
	// Format forward_to
	forwardToStr := "[\n"
	for _, entry := range forwardTo {
		forwardToStr += "{\n"
		if name, ok := entry["name"]; ok {
			forwardToStr += fmt.Sprintf(" name = %q\n", name)
		}
		if address, ok := entry["address"]; ok {
			forwardToStr += fmt.Sprintf(" address = %q\n", address)
		}
		forwardToStr += " },\n"
	}
	forwardToStr += " ]"
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_forward_to" {
   fqdn = %q
   forward_to = %s
}
`, fqdn, forwardToStr)
}

func testAccZoneForwardForwardersOnly(fqdn, externalNsGroup string, forwardersOnly bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_forwarders_only" {
   fqdn = %q
   external_ns_group = %q
   forwarders_only = %t
}
`, fqdn, externalNsGroup, forwardersOnly)
}

func testAccZoneForwardForwardingServers(fqdn, externalNsGroup string, forwardingServers []dns.ZoneForwardForwardingServers) string {

	forwardingServersStr, _ := json.MarshalIndent(forwardingServers, "", " ")
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_forwarding_servers" {
   fqdn = %q
   external_ns_group = %q
   forwarding_servers = %s
}
`, fqdn, externalNsGroup, string(forwardingServersStr))
}

func testAccZoneForwardLocked(fqdn, externalNsGroup string, locked bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_locked" {
   fqdn = %q
   external_ns_group = %q
   locked = %t
}
`, fqdn, externalNsGroup, locked)
}

func testAccZoneForwardMsAdIntegrated(fqdn, externalNsGroup string, msAdIntegrated bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_ms_ad_integrated" {
   fqdn = %q
   external_ns_group = %q
   ms_ad_integrated = %t
}
`, fqdn, externalNsGroup, msAdIntegrated)
}

func testAccZoneForwardMsDdnsMode(fqdn, externalNsGrouop, msDdnsMode string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_ms_ddns_mode" {
   fqdn = %q
   external_ns_group = %q
   ms_ddns_mode = %q
}
`, fqdn, externalNsGrouop, msDdnsMode)
}

func testAccZoneForwardNsGroup(fqdn, extaernalNsGroup, nsGroup string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_ns_group" {
   fqdn = %q
   external_ns_group = %q
   ns_group = %q
}
`, fqdn, extaernalNsGroup, nsGroup)
}

func testAccZoneForwardPrefix(fqdn, externalNsGroup, prefix string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_prefix" {
   fqdn = %q
   external_ns_group = %q
   prefix = %q
}
`, fqdn, externalNsGroup, prefix)
}
