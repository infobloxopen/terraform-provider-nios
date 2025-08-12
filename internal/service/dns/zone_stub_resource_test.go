package dns_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForZoneStub = "address,comment,disable,disable_forwarding,display_domain,dns_fqdn,extattrs,external_ns_group,fqdn,locked,locked_by,mask_prefix,ms_ad_integrated,ms_ddns_mode,ms_managed,ms_read_only,ms_sync_master_name,ns_group,parent,prefix,soa_email,soa_expire,soa_mname,soa_negative_ttl,soa_refresh,soa_retry,soa_serial_number,stub_from,stub_members,stub_msservers,using_srg_associations,view,zone_format"

func TestAccZoneStubResource_basic(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test"
	var v dns.ZoneStub
	fqdn := acctest.RandomNameWithPrefix("zone_stub")
	stubServerName := acctest.RandomNameWithPrefix("stub_server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubBasicConfig(fqdn, "1.1.1.1", stubServerName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_disappears(t *testing.T) {
	resourceName := "nios_dns_zone_stub.test"
	var v dns.ZoneStub
	fqdn := acctest.RandomNameWithPrefix("zone_stub")
	stubServerName := acctest.RandomNameWithPrefix("stub_server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckZoneStubDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneStubBasicConfig(fqdn, "1.1.1.1", stubServerName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					testAccCheckZoneStubDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccZoneStubResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_comment"
	var v dns.ZoneStub
	fqdn := acctest.RandomNameWithPrefix("zone_stub")
	stubServerName := acctest.RandomNameWithPrefix("stub_server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubComment(fqdn, "1.1.1.1", stubServerName, "Example Comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Example Comment"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubComment(fqdn, "1.1.1.1", stubServerName, "Updated Comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated Comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_disable"
	var v dns.ZoneStub
	fqdn := acctest.RandomNameWithPrefix("zone_stub")
	stubServerName := acctest.RandomNameWithPrefix("stub_server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubDisable(fqdn, "1.1.1.1", stubServerName, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubDisable(fqdn, "1.1.1.1", stubServerName, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_DisableForwarding(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_disable_forwarding"
	var v dns.ZoneStub
	fqdn := acctest.RandomNameWithPrefix("zone_stub")
	stubServerName := acctest.RandomNameWithPrefix("stub_server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubDisableForwarding(fqdn, "1.1.1.1", stubServerName, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_forwarding", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubDisableForwarding(fqdn, "1.1.1.1", stubServerName, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_forwarding", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_extattrs"
	var v dns.ZoneStub
	fqdn := acctest.RandomNameWithPrefix("zone_stub")
	stubServerName := acctest.RandomNameWithPrefix("stub_server")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubExtAttrs(fqdn, "1.1.1.1", stubServerName, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubExtAttrs(fqdn, "1.1.1.1", stubServerName, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_ExternalNsGroup(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_external_ns_group"
	var v dns.ZoneStub
	fqdn := acctest.RandomNameWithPrefix("zone_stub")
	stubServerName := acctest.RandomNameWithPrefix("stub_server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubExternalNsGroup(fqdn, "1.1.1.1", stubServerName, "EXTERNAL_NS_GROUP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_ns_group", "EXTERNAL_NS_GROUP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubExternalNsGroup(fqdn, "1.1.1.1", stubServerName, "EXTERNAL_NS_GROUP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_ns_group", "EXTERNAL_NS_GROUP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_Locked(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_locked"
	var v dns.ZoneStub
	fqdn := acctest.RandomNameWithPrefix("zone_stub")
	stubServerName := acctest.RandomNameWithPrefix("stub_server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubLocked(fqdn, "1.1.1.1", stubServerName, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "locked", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubLocked(fqdn, "1.1.1.1", stubServerName, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "locked", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_MsAdIntegrated(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_ms_ad_integrated"
	var v dns.ZoneStub
	fqdn := acctest.RandomNameWithPrefix("zone_stub")
	stubServerName := acctest.RandomNameWithPrefix("stub_server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubMsAdIntegrated(fqdn, "1.1.1.1", stubServerName, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_integrated", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubMsAdIntegrated(fqdn, "1.1.1.1", stubServerName, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_integrated", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_MsDdnsMode(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_ms_ddns_mode"
	var v dns.ZoneStub
	fqdn := acctest.RandomNameWithPrefix("zone_stub")
	stubServerName := acctest.RandomNameWithPrefix("stub_server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubMsDdnsMode(fqdn, "1.1.1.1", stubServerName, "ANY"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ddns_mode", "ANY"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubMsDdnsMode(fqdn, "1.1.1.1", stubServerName, "SECURE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ddns_mode", "SECURE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_NsGroup(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_ns_group"
	var v dns.ZoneStub
	fqdn := acctest.RandomNameWithPrefix("zone_stub")
	stubServerName := acctest.RandomNameWithPrefix("stub_server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubNsGroup(fqdn, "1.1.1.1", stubServerName, "NS_GROUP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ns_group", "NS_GROUP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubNsGroup(fqdn, "1.1.1.1", stubServerName, "NS_GROUP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ns_group", "NS_GROUP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_Prefix(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_prefix"
	var v dns.ZoneStub
	fqdn := acctest.RandomNameWithPrefix("zone_stub")
	stubServerName := acctest.RandomNameWithPrefix("stub_server")
	prefix := acctest.RandomName()
	prefixUpdated := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubPrefix(fqdn, "1.1.1.1", stubServerName, prefix),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", prefix),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubPrefix(fqdn, "1.1.1.1", stubServerName, prefixUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", prefixUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_StubFrom(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_stub_from"
	var v dns.ZoneStub
	fqdn := acctest.RandomNameWithPrefix("zone_stub")
	stubServerName := acctest.RandomNameWithPrefix("stub_server")
	stubServerName2 := acctest.RandomNameWithPrefix("stub_server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubStubFrom(fqdn, "1.1.1.1", stubServerName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "stub_from.0.name", stubServerName),
					resource.TestCheckResourceAttr(resourceName, "stub_from.0.address", "1.1.1.1"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubStubFrom(fqdn, "2.2.2.2", stubServerName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "stub_from.0.name", stubServerName2),
					resource.TestCheckResourceAttr(resourceName, "stub_from.0.address", "2.2.2.2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_StubMembers(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_stub_members"
	var v dns.ZoneStub
	fqdn := acctest.RandomNameWithPrefix("zone_stub")
	stubServerName := acctest.RandomNameWithPrefix("stub_server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubStubMembers(fqdn, "1.1.1.1", stubServerName, "STUB_MEMBERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "stub_members", "STUB_MEMBERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubStubMembers(fqdn, "1.1.1.1", stubServerName, "STUB_MEMBERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "stub_members", "STUB_MEMBERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_StubMsservers(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_stub_msservers"
	var v dns.ZoneStub
	fqdn := acctest.RandomNameWithPrefix("zone_stub")
	stubServerName := acctest.RandomNameWithPrefix("stub_server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubStubMsservers(fqdn, "1.1.1.1", stubServerName, "STUB_MSSERVERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "stub_msservers", "STUB_MSSERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubStubMsservers(fqdn, "1.1.1.1", stubServerName, "STUB_MSSERVERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "stub_msservers", "STUB_MSSERVERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_View(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_view"
	var v dns.ZoneStub
	fqdn := acctest.RandomNameWithPrefix("zone_stub")
	stubServerName := acctest.RandomNameWithPrefix("stub_server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubView(fqdn, "1.1.1.1", stubServerName, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubView(fqdn, "1.1.1.1", stubServerName, "example_view"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "example_view"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// TestAccZoneStubResource_ZoneFormatIPv4 tests the IPv4 zone format for a stub zone.
// It is mandatory to provide a prefix for classless reverse zones.
func TestAccZoneStubResource_ZoneFormatIPv4(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_zone_format"
	var v dns.ZoneStub
	stubServerName := acctest.RandomNameWithPrefix("stub_server")
	prefix := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubZoneFormatIPv4("10.1.0.0/25", "1.1.1.1", stubServerName, "IPV4", prefix),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone_format", "IPV4"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_ZoneFormatIPv6(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_zone_format"
	var v dns.ZoneStub
	//fqdn := acctest.RandomNameWithPrefix("zone_stub")
	stubServerName := acctest.RandomNameWithPrefix("stub_server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubZoneFormatIpv6("2001:db8:85a3:8::/64", "1.1.1.1", stubServerName, "IPV6"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone_format", "IPV6"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckZoneStubExists(ctx context.Context, resourceName string, v *dns.ZoneStub) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			ZoneStubAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForZoneStub).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetZoneStubResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetZoneStubResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckZoneStubDestroy(ctx context.Context, v *dns.ZoneStub) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			ZoneStubAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForZoneStub).
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

func testAccCheckZoneStubDisappears(ctx context.Context, v *dns.ZoneStub) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			ZoneStubAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccZoneStubBasicConfig(fqdn, stubAddress, stubName string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test" {
	fqdn = %q
	stub_from = [{
		address = %q
		name  = %q
	}]
}
`, fqdn, stubAddress, stubName)
}

func testAccZoneStubComment(fqdn, stubAddress, stubName, comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_comment" {
	fqdn = %q
	stub_from = [{
		address = %q
		name  = %q
	}]
    comment = %q
}
`, fqdn, stubAddress, stubName, comment)
}

func testAccZoneStubDisable(fqdn, stubAddress, stubName, disable string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_disable" {
    fqdn = %q
	stub_from = [{
		address = %q
		name  = %q
	}]
	disable = %q
}
`, fqdn, stubAddress, stubName, disable)
}

func testAccZoneStubDisableForwarding(fqdn, stubAddress, stubName, disableForwarding string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_disable_forwarding" {
    fqdn = %q
	stub_from = [{
		address = %q
		name  = %q
	}]
	disable_forwarding = %q
}
`, fqdn, stubAddress, stubName, disableForwarding)
}

func testAccZoneStubExtAttrs(fqdn, stubAddress, stubName string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_extattrs" {
    fqdn = %q
	stub_from = [{
		address = %q
		name  = %q
	}]
	extattrs = %s
}
`, fqdn, stubAddress, stubName, extattrsStr)
}

func testAccZoneStubExternalNsGroup(fqdn, stubAddress, stubName, externalNsGroup string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_external_ns_group" {
    fqdn = %q
	stub_from = [{
		address = %q
		name  = %q
	}]
	external_ns_group = %q
}
`, fqdn, stubAddress, stubName, externalNsGroup)
}

func testAccZoneStubLocked(fqdn, stubAddress, stubName, locked string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_locked" {
    fqdn = %q
	stub_from = [{
		address = %q
		name  = %q
	}]
	locked = %q
}
`, fqdn, stubAddress, stubName, locked)
}

func testAccZoneStubMsAdIntegrated(fqdn, stubAddress, stubName, msAdIntegrated string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_ms_ad_integrated" {
    fqdn = %q
	stub_from = [{
		address = %q
		name  = %q
	}]
	ms_ad_integrated = %q
}
`, fqdn, stubAddress, stubName, msAdIntegrated)
}

func testAccZoneStubMsDdnsMode(fqdn, stubAddress, stubName, msDdnsMode string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_ms_ddns_mode" {
    fqdn = %q
	stub_from = [{
		address = %q
		name  = %q
	}]
	ms_ddns_mode = %q
}
`, fqdn, stubAddress, stubName, msDdnsMode)
}

func testAccZoneStubNsGroup(fqdn, stubAddress, stubName, nsGroup string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_ns_group" {
    fqdn = %q
	stub_from = [{
		address = %q
		name  = %q
	}]
	ns_group = %q
}
`, fqdn, stubAddress, stubName, nsGroup)
}

func testAccZoneStubPrefix(fqdn, stubAddress, stubName, prefix string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_prefix" {
    fqdn = %q
	stub_from = [{
		address = %q
		name  = %q
	}]
	prefix = %q
}
`, fqdn, stubAddress, stubName, prefix)
}

func testAccZoneStubStubFrom(fqdn, stubAddress, stubName string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_stub_from" {
    fqdn = %q
	stub_from = [{
		address = %q
		name  = %q
	}]
}
`, fqdn, stubAddress, stubName)
}

func testAccZoneStubStubMembers(fqdn, stubAddress, stubName, stubMembers string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_stub_members" {
    fqdn = %q
	stub_from = [{
		address = %q
		name  = %q
	}]
	stub_members = %q
}
`, fqdn, stubAddress, stubName, stubMembers)
}

func testAccZoneStubStubMsservers(fqdn, stubAddress, stubName, stubMsservers string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_stub_msservers" {
    fqdn = %q
	stub_from = [{
		address = %q
		name  = %q
	}]
	stub_msservers = %q
}
`, fqdn, stubAddress, stubName, stubMsservers)
}

func testAccZoneStubView(fqdn, stubAddress, stubName, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_view" {
    fqdn = %q
	stub_from = [{
		address = %q
		name  = %q
	}]
	view = %q
}
`, fqdn, stubAddress, stubName, view)
}

func testAccZoneStubZoneFormatIPv4(fqdn, stubAddress, stubName, zoneFormat, prefix string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_zone_format" {
    fqdn = %q
	stub_from = [{
		address = %q
		name  = %q
	}]
	zone_format = %q
	prefix = %q
}
`, fqdn, stubAddress, stubName, zoneFormat, prefix)
}

func testAccZoneStubZoneFormatIpv6(fqdn, stubAddress, stubName, zoneFormat string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_zone_format" {
    fqdn = %q
	stub_from = [{
		address = %q
		name  = %q
	}]
	zone_format = %q
}
`, fqdn, stubAddress, stubName, zoneFormat)
}
