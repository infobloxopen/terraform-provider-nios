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

var readableAttributesForZoneStub = "address,comment,disable,disable_forwarding,display_domain,dns_fqdn,extattrs,external_ns_group,fqdn,locked,locked_by,mask_prefix,ms_ad_integrated,ms_ddns_mode,ms_managed,ms_read_only,ms_sync_master_name,ns_group,parent,prefix,soa_email,soa_expire,soa_mname,soa_negative_ttl,soarefresh,soa_retry,soa_serial_number,stub_from,stub_members,stub_msservers,using_srg_associations,view,zone_format"

func TestAccZoneStubResource_basic(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test"
	var v dns.ZoneStub

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubBasicConfig(""),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckZoneStubDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneStubBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					testAccCheckZoneStubDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccZoneStubResource_Ref(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_ref"
	var v dns.ZoneStub

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_comment"
	var v dns.ZoneStub

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_disable"
	var v dns.ZoneStub

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_DisableForwarding(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_disable_forwarding"
	var v dns.ZoneStub

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubDisableForwarding("DISABLE_FORWARDING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_forwarding", "DISABLE_FORWARDING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubDisableForwarding("DISABLE_FORWARDING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_forwarding", "DISABLE_FORWARDING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_extattrs"
	var v dns.ZoneStub

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_ExternalNsGroup(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_external_ns_group"
	var v dns.ZoneStub

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubExternalNsGroup("EXTERNAL_NS_GROUP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_ns_group", "EXTERNAL_NS_GROUP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubExternalNsGroup("EXTERNAL_NS_GROUP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_ns_group", "EXTERNAL_NS_GROUP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_Fqdn(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_fqdn"
	var v dns.ZoneStub

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubFqdn("FQDN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "FQDN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubFqdn("FQDN_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "FQDN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_Locked(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_locked"
	var v dns.ZoneStub

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubLocked("LOCKED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "locked", "LOCKED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubLocked("LOCKED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "locked", "LOCKED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_MsAdIntegrated(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_ms_ad_integrated"
	var v dns.ZoneStub

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubMsAdIntegrated("MS_AD_INTEGRATED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_integrated", "MS_AD_INTEGRATED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubMsAdIntegrated("MS_AD_INTEGRATED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_integrated", "MS_AD_INTEGRATED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_MsDdnsMode(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_ms_ddns_mode"
	var v dns.ZoneStub

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubMsDdnsMode("MS_DDNS_MODE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ddns_mode", "MS_DDNS_MODE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubMsDdnsMode("MS_DDNS_MODE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ddns_mode", "MS_DDNS_MODE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_NsGroup(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_ns_group"
	var v dns.ZoneStub

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubNsGroup("NS_GROUP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ns_group", "NS_GROUP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubNsGroup("NS_GROUP_UPDATE_REPLACE_ME"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubPrefix("PREFIX_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", "PREFIX_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubPrefix("PREFIX_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", "PREFIX_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_StubFrom(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_stub_from"
	var v dns.ZoneStub

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubStubFrom("STUB_FROM_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "stub_from", "STUB_FROM_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubStubFrom("STUB_FROM_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "stub_from", "STUB_FROM_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_StubMembers(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_stub_members"
	var v dns.ZoneStub

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubStubMembers("STUB_MEMBERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "stub_members", "STUB_MEMBERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubStubMembers("STUB_MEMBERS_UPDATE_REPLACE_ME"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubStubMsservers("STUB_MSSERVERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "stub_msservers", "STUB_MSSERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubStubMsservers("STUB_MSSERVERS_UPDATE_REPLACE_ME"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubView("VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubView("VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneStubResource_ZoneFormat(t *testing.T) {
	var resourceName = "nios_dns_zone_stub.test_zone_format"
	var v dns.ZoneStub

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneStubZoneFormat("ZONE_FORMAT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone_format", "ZONE_FORMAT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneStubZoneFormat("ZONE_FORMAT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone_format", "ZONE_FORMAT_UPDATE_REPLACE_ME"),
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

func testAccZoneStubBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test" {
}
`)
}

func testAccZoneStubRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccZoneStubComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccZoneStubDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccZoneStubDisableForwarding(disableForwarding string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_disable_forwarding" {
    disable_forwarding = %q
}
`, disableForwarding)
}

func testAccZoneStubExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccZoneStubExternalNsGroup(externalNsGroup string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_external_ns_group" {
    external_ns_group = %q
}
`, externalNsGroup)
}

func testAccZoneStubFqdn(fqdn string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_fqdn" {
    fqdn = %q
}
`, fqdn)
}

func testAccZoneStubLocked(locked string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_locked" {
    locked = %q
}
`, locked)
}

func testAccZoneStubMsAdIntegrated(msAdIntegrated string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_ms_ad_integrated" {
    ms_ad_integrated = %q
}
`, msAdIntegrated)
}

func testAccZoneStubMsDdnsMode(msDdnsMode string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_ms_ddns_mode" {
    ms_ddns_mode = %q
}
`, msDdnsMode)
}

func testAccZoneStubNsGroup(nsGroup string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_ns_group" {
    ns_group = %q
}
`, nsGroup)
}

func testAccZoneStubPrefix(prefix string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_prefix" {
    prefix = %q
}
`, prefix)
}

func testAccZoneStubStubFrom(stubFrom string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_stub_from" {
    stub_from = %q
}
`, stubFrom)
}

func testAccZoneStubStubMembers(stubMembers string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_stub_members" {
    stub_members = %q
}
`, stubMembers)
}

func testAccZoneStubStubMsservers(stubMsservers string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_stub_msservers" {
    stub_msservers = %q
}
`, stubMsservers)
}

func testAccZoneStubView(view string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_view" {
    view = %q
}
`, view)
}

func testAccZoneStubZoneFormat(zoneFormat string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test_zone_format" {
    zone_format = %q
}
`, zoneFormat)
}
