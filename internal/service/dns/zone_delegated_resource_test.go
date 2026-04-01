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

//TODO : OBJECTS TO BE PRESENT IN GRID FOR TESTS
// - IPv4 reverse zone - 10.1.0.0
// - IPv6 reverse zone - 2001:db8:85a3::8a2e:370:7334/128

var readableAttributesForZoneDelegated = "address,comment,delegate_to,delegated_ttl,disable,display_domain,dns_fqdn,enable_rfc2317_exclusion,extattrs,fqdn,locked,locked_by,mask_prefix,ms_ad_integrated,ms_ddns_mode,ms_managed,ms_read_only,ms_sync_master_name,ns_group,parent,prefix,use_delegated_ttl,using_srg_associations,view,zone_format"

func TestAccZoneDelegatedResource_basic(t *testing.T) {
	var resourceName = "nios_dns_zone_delegated.test"
	var v dns.ZoneDelegated
	fqdn := acctest.RandomNameWithPrefix("zone-delegated") + ".example.com"
	delegatedToName := acctest.RandomNameWithPrefix("zone-delegated") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneDelegatedBasicConfig(fqdn, delegatedToName, "10.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_rfc2317_exclusion", "false"),
					resource.TestCheckResourceAttr(resourceName, "locked", "false"),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_integrated", "false"),
					resource.TestCheckResourceAttr(resourceName, "ms_ddns_mode", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
					resource.TestCheckResourceAttr(resourceName, "zone_format", "FORWARD"),
					resource.TestCheckResourceAttr(resourceName, "use_delegated_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneDelegatedResource_disappears(t *testing.T) {
	resourceName := "nios_dns_zone_delegated.test"
	var v dns.ZoneDelegated
	fqdn := acctest.RandomNameWithPrefix("zone-delegated") + ".example.com"
	delegatedToName := acctest.RandomNameWithPrefix("zone-delegated") + "com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckZoneDelegatedDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneDelegatedBasicConfig(fqdn, delegatedToName, "10.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					testAccCheckZoneDelegatedDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccZoneDelegatedResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_zone_delegated.test_comment"
	var v dns.ZoneDelegated
	fqdn := acctest.RandomNameWithPrefix("zone-delegated") + ".example.com"
	delegatedToName := acctest.RandomNameWithPrefix("zone-delegated") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneDelegatedComment(fqdn, delegatedToName, "10.0.0.1", "This is a delegated zone"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a delegated zone"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneDelegatedComment(fqdn, delegatedToName, "10.0.0.1", "This is an updated delegated zone"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated delegated zone"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneDelegatedResource_DelegateTo(t *testing.T) {
	var resourceName = "nios_dns_zone_delegated.test_delegate_to"
	var v dns.ZoneDelegated
	fqdn := acctest.RandomNameWithPrefix("zone-delegated") + ".example.com"
	delegatedToName1 := acctest.RandomNameWithPrefix("zone-delegated") + ".com"
	delegatedToName2 := acctest.RandomNameWithPrefix("zone-delegated") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneDelegatedDelegateTo(fqdn, delegatedToName1, "10.0.0.3"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegate_to.0.name", delegatedToName1),
					resource.TestCheckResourceAttr(resourceName, "delegate_to.0.address", "10.0.0.3"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneDelegatedDelegateTo(fqdn, delegatedToName2, "10.0.0.4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegate_to.0.name", delegatedToName2),
					resource.TestCheckResourceAttr(resourceName, "delegate_to.0.address", "10.0.0.4"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneDelegatedResource_DelegatedTtl(t *testing.T) {
	var resourceName = "nios_dns_zone_delegated.test_delegated_ttl"
	var v dns.ZoneDelegated
	fqdn := acctest.RandomNameWithPrefix("zone-delegated") + ".example.com"
	delegatedToName := acctest.RandomNameWithPrefix("zone-delegated") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneDelegatedDelegatedTtl(fqdn, delegatedToName, "10.0.0.1", 3600, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegated_ttl", "3600"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneDelegatedDelegatedTtl(fqdn, delegatedToName, "10.0.0.1", 7200, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegated_ttl", "7200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneDelegatedResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_zone_delegated.test_disable"
	var v dns.ZoneDelegated
	fqdn := acctest.RandomNameWithPrefix("zone-delegated") + ".example.com"
	delegatedToName := acctest.RandomNameWithPrefix("zone-delegated") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneDelegatedDisable(fqdn, delegatedToName, "10.0.0.1", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneDelegatedDisable(fqdn, delegatedToName, "10.0.0.1", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneDelegatedResource_EnableRfc2317Exclusion(t *testing.T) {
	var resourceName = "nios_dns_zone_delegated.test_enable_rfc2317_exclusion"
	var v dns.ZoneDelegated
	fqdn := acctest.RandomNameWithPrefix("zone-delegated") + ".example.com"
	delegatedToName := acctest.RandomNameWithPrefix("zone-delegated") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneDelegatedEnableRfc2317Exclusion(fqdn, delegatedToName, "10.0.0.1", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_rfc2317_exclusion", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneDelegatedEnableRfc2317Exclusion(fqdn, delegatedToName, "10.0.0.1", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_rfc2317_exclusion", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneDelegatedResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_zone_delegated.test_extattrs"
	var v dns.ZoneDelegated
	fqdn := acctest.RandomNameWithPrefix("zone-delegated") + ".example.com"
	delegatedToName := acctest.RandomNameWithPrefix("zone-delegated") + ".com"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneDelegatedExtAttrs(fqdn, delegatedToName, "10.0.0.1", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccZoneDelegatedExtAttrs(fqdn, delegatedToName, "10.0.0.1", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneDelegatedResource_Locked(t *testing.T) {
	var resourceName = "nios_dns_zone_delegated.test_locked"
	var v dns.ZoneDelegated
	fqdn := acctest.RandomNameWithPrefix("zone-delegated") + ".example.com"
	delegatedToName := acctest.RandomNameWithPrefix("zone-delegated") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneDelegatedLocked(fqdn, delegatedToName, "10.0.0.1", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "locked", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneDelegatedLocked(fqdn, delegatedToName, "10.0.0.1", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "locked", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneDelegatedResource_MsAdIntegrated(t *testing.T) {
	var resourceName = "nios_dns_zone_delegated.test_ms_ad_integrated"
	var v dns.ZoneDelegated
	fqdn := acctest.RandomNameWithPrefix("zone-delegated") + ".example.com"
	delegatedToName := acctest.RandomNameWithPrefix("zone-delegated") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneDelegatedMsAdIntegrated(fqdn, delegatedToName, "10.0.0.1", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_integrated", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneDelegatedMsAdIntegrated(fqdn, delegatedToName, "10.0.0.1", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_integrated", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneDelegatedResource_MsDdnsMode(t *testing.T) {
	var resourceName = "nios_dns_zone_delegated.test_ms_ddns_mode"
	var v dns.ZoneDelegated
	fqdn := acctest.RandomNameWithPrefix("zone-delegated") + ".example.com"
	delegatedToName := acctest.RandomNameWithPrefix("zone-delegated") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneDelegatedMsDdnsMode(fqdn, delegatedToName, "10.0.0.1", "NONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ddns_mode", "NONE"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneDelegatedMsDdnsMode(fqdn, delegatedToName, "10.0.0.1", "ANY"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ddns_mode", "ANY"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneDelegatedMsDdnsModeSecure(fqdn, delegatedToName, "10.0.0.1", "SECURE", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ddns_mode", "SECURE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneDelegatedResource_NsGroup(t *testing.T) {
	var resourceName = "nios_dns_zone_delegated.test_ns_group"
	var v dns.ZoneDelegated
	fqdn := acctest.RandomNameWithPrefix("zone-delegated") + ".example.com"
	delegatedToName := acctest.RandomNameWithPrefix("zone-delegated") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneDelegatedNsGroup(fqdn, delegatedToName, "10.0.0.1", "example_nsg1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ns_group", "example_nsg1"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneDelegatedNsGroup(fqdn, delegatedToName, "10.0.0.1", "example_nsg2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ns_group", "example_nsg2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneDelegatedResource_Prefix(t *testing.T) {
	var resourceName = "nios_dns_zone_delegated.test_prefix"
	var v dns.ZoneDelegated
	fqdn := acctest.RandomNameWithPrefix("zone-delegated") + ".example.com"
	delegatedToName := acctest.RandomNameWithPrefix("zone-delegated") + ".com"
	prefix := acctest.RandomName()
	prefixUpdated := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneDelegatedPrefix(fqdn, delegatedToName, "10.0.0.1", prefix),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", prefix),
				),
			},
			// Update and Read
			{
				Config: testAccZoneDelegatedPrefix(fqdn, delegatedToName, "10.0.0.1", prefixUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", prefixUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneDelegatedResource_UseDelegatedTtl(t *testing.T) {
	var resourceName = "nios_dns_zone_delegated.test_use_delegated_ttl"
	var v dns.ZoneDelegated
	fqdn := acctest.RandomNameWithPrefix("zone-delegated") + ".example.com"
	delegatedToName := acctest.RandomNameWithPrefix("zone-delegated") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneDelegatedUseDelegatedTtl(fqdn, delegatedToName, "10.0.0.1", true, 1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_delegated_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneDelegatedUseDelegatedTtl(fqdn, delegatedToName, "10.0.0.1", false, 1800),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_delegated_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneDelegatedResource_ZoneFormatIPV4(t *testing.T) {
	var resourceName = "nios_dns_zone_delegated.test_zone_format"
	var v dns.ZoneDelegated
	delegatedToName := acctest.RandomNameWithPrefix("zone-delegated") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneDelegatedZoneFormat("10.1.0.132/32", delegatedToName, "10.0.0.1", "IPV4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone_format", "IPV4"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneDelegatedResource_ZoneFormatIPV6(t *testing.T) {
	var resourceName = "nios_dns_zone_delegated.test_zone_format"
	var v dns.ZoneDelegated
	delegatedToName := acctest.RandomNameWithPrefix("zone-delegated") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneDelegatedZoneFormat("8.0.0.0.3.a.5.8.8.b.d.0.1.0.0.2.ip6.arpa.", delegatedToName, "10.0.0.1", "IPV6"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone_format", "IPV6"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckZoneDelegatedExists(ctx context.Context, resourceName string, v *dns.ZoneDelegated) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			ZoneDelegatedAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForZoneDelegated).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetZoneDelegatedResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetZoneDelegatedResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckZoneDelegatedDestroy(ctx context.Context, v *dns.ZoneDelegated) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			ZoneDelegatedAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForZoneDelegated).
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

func testAccCheckZoneDelegatedDisappears(ctx context.Context, v *dns.ZoneDelegated) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			ZoneDelegatedAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccZoneDelegatedImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.Attributes["ref"] == "" {
			return "", fmt.Errorf("ref is not set")
		}
		return rs.Primary.Attributes["ref"], nil
	}
}

func testAccZoneDelegatedBasicConfig(fqdn, delegateToName, delegateToAddress string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_delegated" "test" {
	fqdn = %q
    delegate_to = [
		{
			name = %q
			address = %q
		}
	]
}
`, fqdn, delegateToName, delegateToAddress)
}

func testAccZoneDelegatedComment(fqdn, delegateToName, delegateToAddress, comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_delegated" "test_comment" {
	fqdn = %q
    delegate_to = [
		{
			name = %q
			address = %q
		}
	]
    comment = %q
}
`, fqdn, delegateToName, delegateToAddress, comment)
}

func testAccZoneDelegatedDelegateTo(fqdn, delegateToName, delegateToAddress string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_delegated" "test_delegate_to" {
	fqdn = %q
    delegate_to = [
		{
			name = %q
			address = %q
		}
	]
}
`, fqdn, delegateToName, delegateToAddress)
}

func testAccZoneDelegatedDelegatedTtl(fqdn, delegateToName, delegateToAddress string, delegatedTtl int, useDelegatedTtl bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_delegated" "test_delegated_ttl" {
	fqdn = %q
    delegate_to = [
		{
			name = %q
			address = %q
		}
	]
    delegated_ttl = %d
	use_delegated_ttl = %t
}
`, fqdn, delegateToName, delegateToAddress, delegatedTtl, useDelegatedTtl)
}

func testAccZoneDelegatedDisable(fqdn, delegateToName, delegateToAddress string, disable bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_delegated" "test_disable" {
	fqdn = %q
    delegate_to = [
		{
			name = %q
			address = %q
		}
	]
    disable = %t
}
`, fqdn, delegateToName, delegateToAddress, disable)
}

func testAccZoneDelegatedEnableRfc2317Exclusion(fqdn, delegateToName, delegateToAddress string, enableRfc2317Exclusion bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_delegated" "test_enable_rfc2317_exclusion" {
	fqdn = %q
    delegate_to = [
		{
			name = %q
			address = %q
		}
	]
    enable_rfc2317_exclusion = %t
}
`, fqdn, delegateToName, delegateToAddress, enableRfc2317Exclusion)
}

func testAccZoneDelegatedExtAttrs(fqdn, delegateToName, delegateToAddress string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf("    %q = %q\n", k, v)
	}
	extattrsStr += "  }"
	return fmt.Sprintf(`
resource "nios_dns_zone_delegated" "test_extattrs" {
	fqdn = %q
    delegate_to = [
		{
			name = %q
			address = %q
		}
	]
    extattrs = %s
}
`, fqdn, delegateToName, delegateToAddress, extattrsStr)
}

func testAccZoneDelegatedLocked(fqdn, delegateToName, delegateToAddress string, locked bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_delegated" "test_locked" {
	fqdn = %q
    delegate_to = [
		{
			name = %q
			address = %q
		}
	]
    locked = %t
}
`, fqdn, delegateToName, delegateToAddress, locked)
}

func testAccZoneDelegatedMsAdIntegrated(fqdn, delegateToName, delegateToAddress string, msAdIntegrated bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_delegated" "test_ms_ad_integrated" {
	fqdn = %q
    delegate_to = [
		{
			name = %q
			address = %q
		}
	]
    ms_ad_integrated = %t
}
`, fqdn, delegateToName, delegateToAddress, msAdIntegrated)
}

func testAccZoneDelegatedMsDdnsMode(fqdn, delegateToName, delegateToAddress, msDdnsMode string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_delegated" "test_ms_ddns_mode" {
	fqdn = %q
    delegate_to = [
		{
			name = %q
			address = %q
		}
	]
    ms_ddns_mode = %q
}
`, fqdn, delegateToName, delegateToAddress, msDdnsMode)
}

func testAccZoneDelegatedMsDdnsModeSecure(fqdn, delegateToName, delegateToAddress, msDdnsMode, msAdIntegrated string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_delegated" "test_ms_ddns_mode" {
	fqdn = %q
    delegate_to = [
		{
			name = %q
			address = %q
		}
	]
    ms_ddns_mode = %q
    ms_ad_integrated = %q
}
`, fqdn, delegateToName, delegateToAddress, msDdnsMode, msAdIntegrated)
}

func testAccZoneDelegatedNsGroup(fqdn, delegateToName, delegateToAddress, nsGroup string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_delegated" "test_ns_group" {
	fqdn = %q
    delegate_to = [
		{
			name = %q
			address = %q
		}
	]
    ns_group = %q
}
`, fqdn, delegateToName, delegateToAddress, nsGroup)
}

func testAccZoneDelegatedPrefix(fqdn, delegateToName, delegateToAddress, prefix string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_delegated" "test_prefix" {
	fqdn = %q
    delegate_to = [
		{
			name = %q
			address = %q
		}
	]
    prefix = %q
}
`, fqdn, delegateToName, delegateToAddress, prefix)
}

func testAccZoneDelegatedUseDelegatedTtl(fqdn, delegateToName, delegateToAddress string, useDelegatedTtl bool, delegatedTTL int) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_delegated" "test_use_delegated_ttl" {
	fqdn = %q
    delegate_to = [
		{
			name = %q
			address = %q
		}
	]
    use_delegated_ttl = %t
	delegated_ttl = %d
}
`, fqdn, delegateToName, delegateToAddress, useDelegatedTtl, delegatedTTL)
}

func testAccZoneDelegatedZoneFormat(fqdn, delegateToName, delegateToAddress, zoneFormat string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_delegated" "test_zone_format" {
	fqdn = %q
    delegate_to = [
		{
			name = %q
			address = %q
		}
	]
    zone_format = %q
}
`, fqdn, delegateToName, delegateToAddress, zoneFormat)
}

func TestAccZoneDelegatedResource_Import(t *testing.T) {
	var resourceName = "nios_dns_zone_delegated.test"
	var v dns.ZoneDelegated
	fqdn := acctest.RandomNameWithPrefix("zone-delegated") + ".example.com"
	delegatedToName := acctest.RandomNameWithPrefix("zone-delegated") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccZoneDelegatedBasicConfig(fqdn, delegatedToName, "10.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
				),
			},
			// Import with PlanOnly to detect differences
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccZoneDelegatedImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "ref",
				PlanOnly:                             true,
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccZoneDelegatedImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
