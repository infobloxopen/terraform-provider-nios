package dns_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForSharedrecordAaaa = "comment,disable,dns_name,extattrs,ipv6addr,name,shared_record_group,ttl,use_ttl"

func TestAccSharedrecordAaaaResource_basic(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_aaaa.test"
	var v dns.SharedrecordAaaa
	name := acctest.RandomNameWithPrefix("sharedrecord-aaaa")
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordAaaaBasicConfig(name, "2001:db8::1", sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", "2001:db8::1"),
					resource.TestCheckResourceAttr(resourceName, "shared_record_group", sharedRecordGroup),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordAaaaResource_disappears(t *testing.T) {
	resourceName := "nios_dns_sharedrecord_aaaa.test"
	var v dns.SharedrecordAaaa
	name := acctest.RandomNameWithPrefix("sharedrecord-aaaa")
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharedrecordAaaaDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharedrecordAaaaBasicConfig(name, "2001:db8::1", sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAaaaExists(context.Background(), resourceName, &v),
					testAccCheckSharedrecordAaaaDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSharedrecordAaaaResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_aaaa.test_comment"
	var v dns.SharedrecordAaaa
	name := acctest.RandomNameWithPrefix("sharedrecord-aaaa")
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordAaaaComment(name, "2001:db8::1", sharedRecordGroup, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordAaaaComment(name, "2001:db8::1", sharedRecordGroup, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordAaaaResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_aaaa.test_disable"
	var v dns.SharedrecordAaaa
	name := acctest.RandomNameWithPrefix("sharedrecord-aaaa")
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordAaaaDisable(name, "2001:db8::1", sharedRecordGroup, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordAaaaDisable(name, "2001:db8::1", sharedRecordGroup, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordAaaaResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_aaaa.test_extattrs"
	var v dns.SharedrecordAaaa
	name := acctest.RandomNameWithPrefix("sharedrecord-aaaa") + ".example.com"
	ipv6addr := "2001:db8::1"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordAaaaExtAttrs(name, ipv6addr, sharedRecordGroup, map[string]any{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordAaaaExtAttrs(name, ipv6addr, sharedRecordGroup, map[string]any{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordAaaaResource_Ipv6addr(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_aaaa.test_ipv6addr"
	var v dns.SharedrecordAaaa
	name := acctest.RandomNameWithPrefix("sharedrecord-aaaa")
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordAaaaIpv6addr(name, "2001:db8::1", sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", "2001:db8::1"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordAaaaIpv6addr(name, "2001:db8::2", sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", "2001:db8::2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordAaaaResource_Name(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_aaaa.test_name"
	var v dns.SharedrecordAaaa
	name1 := acctest.RandomNameWithPrefix("sharedrecord-aaaa")
	name2 := acctest.RandomNameWithPrefix("updatedsharedrecord-aaaa")
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordAaaaName(name1, "2001:db8::1", sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordAaaaName(name2, "2001:db8::1", sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordAaaaResource_SharedRecordGroup(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_aaaa.test_shared_record_group"
	var v dns.SharedrecordAaaa
	name := acctest.RandomNameWithPrefix("sharedrecord-aaaa")
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordAaaaSharedRecordGroup(name, "2001:db8::1", sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "shared_record_group", sharedRecordGroup),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordAaaaResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_aaaa.test_ttl"
	var v dns.SharedrecordAaaa
	name := acctest.RandomNameWithPrefix("sharedrecord-aaaa")
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordAaaaTtl(name, "2001:db8::1", sharedRecordGroup, 3600, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "3600"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordAaaaTtl(name, "2001:db8::1", sharedRecordGroup, 7200, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "7200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordAaaaResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_aaaa.test_use_ttl"
	var v dns.SharedrecordAaaa
	name := acctest.RandomNameWithPrefix("sharedrecord-aaaa")
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordAaaaUseTtl(name, "2001:db8::1", sharedRecordGroup, 3600, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordAaaaUseTtl(name, "2001:db8::1", sharedRecordGroup, 7200, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckSharedrecordAaaaExists(ctx context.Context, resourceName string, v *dns.SharedrecordAaaa) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			SharedrecordAaaaAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForSharedrecordAaaa).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetSharedrecordAaaaResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetSharedrecordAaaaResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckSharedrecordAaaaDestroy(ctx context.Context, v *dns.SharedrecordAaaa) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		if v == nil || v.Ref == nil { // Add nil check
			return nil
		}
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			SharedrecordAaaaAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForSharedrecordAaaa).
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

func testAccCheckSharedrecordAaaaDisappears(ctx context.Context, v *dns.SharedrecordAaaa) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			SharedrecordAaaaAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccSharedrecordAaaaBasicConfig(name, ipv6addr, sharedRecordGroup string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_aaaa" "test" {
    name                = %q
    ipv6addr            = %q
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
}
`, name, ipv6addr)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordAaaaComment(name, ipv6addr, sharedRecordGroup, comment string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_aaaa" "test_comment" {
    name                = %q
    ipv6addr            = %q
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
    comment             = %q
}
`, name, ipv6addr, comment)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordAaaaDisable(name, ipv6addr, sharedRecordGroup string, disable bool) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_aaaa" "test_disable" {
    name                = %q
    ipv6addr            = %q
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
    disable             = %t
}
`, name, ipv6addr, disable)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordAaaaExtAttrs(name, ipv6addr, sharedRecordGroup string, extAttrs map[string]any) string {
	extAttrsStr := utils.ConvertMapToHCL(extAttrs)
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_aaaa" "test_extattrs" {
    name                = %q
    ipv6addr            = %q
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
    extattrs            = %s
}
`, name, ipv6addr, extAttrsStr)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordAaaaIpv6addr(name, ipv6addr, sharedRecordGroup string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_aaaa" "test_ipv6addr" {
    name                = %q
    ipv6addr            = %q
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
}
`, name, ipv6addr)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordAaaaName(name, ipv6addr, sharedRecordGroup string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_aaaa" "test_name" {
    name                = %q
    ipv6addr            = %q
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
}
`, name, ipv6addr)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordAaaaSharedRecordGroup(name, ipv6addr, sharedRecordGroup string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_aaaa" "test_shared_record_group" {
    name                = %q
    ipv6addr            = %q
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
}
`, name, ipv6addr)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordAaaaTtl(name, ipv6addr, sharedRecordGroup string, ttl int, useTtl bool) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_aaaa" "test_ttl" {
    name                = %q
    ipv6addr            = %q
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
    ttl                 = %d
    use_ttl             = %t
}
`, name, ipv6addr, ttl, useTtl)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordAaaaUseTtl(name, ipv6addr, sharedRecordGroup string, ttl int, useTtl bool) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_aaaa" "test_use_ttl" {
    name                = %q
    ipv6addr            = %q
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
    ttl                 = %d
    use_ttl             = %t
}
`, name, ipv6addr, ttl, useTtl)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}
