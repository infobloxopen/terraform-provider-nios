package dns_test

import (
	"context"
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

// TODO : Add readable attributes for the resource
var readableAttributesForZoneForward = ""

func TestAccZoneForwardResource_basic(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test"
	var v dns.ZoneForward

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_disappears(t *testing.T) {
	resourceName := "nios_dns_zone_forward.test"
	var v dns.ZoneForward

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckZoneForwardDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneForwardBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					testAccCheckZoneForwardDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccZoneForwardResource_Ref(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test__ref"
	var v dns.ZoneForward

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "_ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "_ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_comment"
	var v dns.ZoneForward

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_disable"
	var v dns.ZoneForward

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_DisableNsGeneration(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_disable_ns_generation"
	var v dns.ZoneForward

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardDisableNsGeneration("DISABLE_NS_GENERATION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_ns_generation", "DISABLE_NS_GENERATION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardDisableNsGeneration("DISABLE_NS_GENERATION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_ns_generation", "DISABLE_NS_GENERATION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_extattrs"
	var v dns.ZoneForward

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_ExternalNsGroup(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_external_ns_group"
	var v dns.ZoneForward

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardExternalNsGroup("EXTERNAL_NS_GROUP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_ns_group", "EXTERNAL_NS_GROUP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardExternalNsGroup("EXTERNAL_NS_GROUP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_ns_group", "EXTERNAL_NS_GROUP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_ForwardTo(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_forward_to"
	var v dns.ZoneForward

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardForwardTo("FORWARD_TO_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forward_to", "FORWARD_TO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardForwardTo("FORWARD_TO_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forward_to", "FORWARD_TO_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_ForwardersOnly(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_forwarders_only"
	var v dns.ZoneForward

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardForwardersOnly("FORWARDERS_ONLY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarders_only", "FORWARDERS_ONLY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardForwardersOnly("FORWARDERS_ONLY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarders_only", "FORWARDERS_ONLY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_ForwardingServers(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_forwarding_servers"
	var v dns.ZoneForward

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardForwardingServers("FORWARDING_SERVERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarding_servers", "FORWARDING_SERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardForwardingServers("FORWARDING_SERVERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarding_servers", "FORWARDING_SERVERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_Fqdn(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_fqdn"
	var v dns.ZoneForward

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardFqdn("FQDN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "FQDN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardFqdn("FQDN_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "FQDN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_Locked(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_locked"
	var v dns.ZoneForward

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardLocked("LOCKED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "locked", "LOCKED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardLocked("LOCKED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "locked", "LOCKED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_MsAdIntegrated(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_ms_ad_integrated"
	var v dns.ZoneForward

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardMsAdIntegrated("MS_AD_INTEGRATED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_integrated", "MS_AD_INTEGRATED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardMsAdIntegrated("MS_AD_INTEGRATED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_integrated", "MS_AD_INTEGRATED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_MsDdnsMode(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_ms_ddns_mode"
	var v dns.ZoneForward

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardMsDdnsMode("MS_DDNS_MODE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ddns_mode", "MS_DDNS_MODE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardMsDdnsMode("MS_DDNS_MODE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ddns_mode", "MS_DDNS_MODE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_NsGroup(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_ns_group"
	var v dns.ZoneForward

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardNsGroup("NS_GROUP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ns_group", "NS_GROUP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardNsGroup("NS_GROUP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ns_group", "NS_GROUP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_Prefix(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_prefix"
	var v dns.ZoneForward

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardPrefix("PREFIX_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", "PREFIX_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardPrefix("PREFIX_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", "PREFIX_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_View(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_view"
	var v dns.ZoneForward

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardView("VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardView("VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneForwardResource_ZoneFormat(t *testing.T) {
	var resourceName = "nios_dns_zone_forward.test_zone_format"
	var v dns.ZoneForward

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneForwardZoneFormat("ZONE_FORMAT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone_format", "ZONE_FORMAT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneForwardZoneFormat("ZONE_FORMAT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone_format", "ZONE_FORMAT_UPDATE_REPLACE_ME"),
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

func testAccZoneForwardBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test" {
}
`)
}

func testAccZoneForwardRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test__ref" {
    _ref = %q
}
`, ref)
}

func testAccZoneForwardComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccZoneForwardDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccZoneForwardDisableNsGeneration(disableNsGeneration string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_disable_ns_generation" {
    disable_ns_generation = %q
}
`, disableNsGeneration)
}

func testAccZoneForwardExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccZoneForwardExternalNsGroup(externalNsGroup string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_external_ns_group" {
    external_ns_group = %q
}
`, externalNsGroup)
}

func testAccZoneForwardForwardTo(forwardTo string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_forward_to" {
    forward_to = %q
}
`, forwardTo)
}

func testAccZoneForwardForwardersOnly(forwardersOnly string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_forwarders_only" {
    forwarders_only = %q
}
`, forwardersOnly)
}

func testAccZoneForwardForwardingServers(forwardingServers string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_forwarding_servers" {
    forwarding_servers = %q
}
`, forwardingServers)
}

func testAccZoneForwardFqdn(fqdn string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_fqdn" {
    fqdn = %q
}
`, fqdn)
}

func testAccZoneForwardLocked(locked string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_locked" {
    locked = %q
}
`, locked)
}

func testAccZoneForwardMsAdIntegrated(msAdIntegrated string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_ms_ad_integrated" {
    ms_ad_integrated = %q
}
`, msAdIntegrated)
}

func testAccZoneForwardMsDdnsMode(msDdnsMode string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_ms_ddns_mode" {
    ms_ddns_mode = %q
}
`, msDdnsMode)
}

func testAccZoneForwardNsGroup(nsGroup string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_ns_group" {
    ns_group = %q
}
`, nsGroup)
}

func testAccZoneForwardPrefix(prefix string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_prefix" {
    prefix = %q
}
`, prefix)
}

func testAccZoneForwardView(view string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_view" {
    view = %q
}
`, view)
}

func testAccZoneForwardZoneFormat(zoneFormat string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test_zone_format" {
    zone_format = %q
}
`, zoneFormat)
}
