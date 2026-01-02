package rpz_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/rpz"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForRecordRpzAaaaIpaddress = "comment,disable,extattrs,ipv6addr,name,rp_zone,ttl,use_ttl,view,zone"

func TestAccRecordRpzAaaaIpaddressResource_basic(t *testing.T) {
	var resourceName = "nios_rpz_record_aaaa_ipaddress.test"
	var v rpz.RecordRpzAaaaIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "2001:db8::/64" + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAaaaIpaddressBasicConfig(name, "2001:db8::10", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", "2001:db8::10"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAaaaIpaddressResource_disappears(t *testing.T) {
	resourceName := "nios_rpz_record_aaaa_ipaddress.test"
	var v rpz.RecordRpzAaaaIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "2001:db8::/64" + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzAaaaIpaddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzAaaaIpaddressBasicConfig(name, "2001:db8::10", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaIpaddressExists(context.Background(), resourceName, &v),
					testAccCheckRecordRpzAaaaIpaddressDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordRpzAaaaIpaddressResource_Comment(t *testing.T) {
	var resourceName = "nios_rpz_record_aaaa_ipaddress.test_comment"
	var v rpz.RecordRpzAaaaIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "2001:db8::/64" + "." + rpZone
	comment1 := "This is a new rpz aaaa ipaddress record"
	comment2 := "This is an updated rpz aaaa ipaddress record"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAaaaIpaddressComment(name, "2001:db8::10", rpZone, comment1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAaaaIpaddressComment(name, "2001:db8::10", rpZone, comment2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAaaaIpaddressResource_Disable(t *testing.T) {
	var resourceName = "nios_rpz_record_aaaa_ipaddress.test_disable"
	var v rpz.RecordRpzAaaaIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "2001:db8::/64" + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAaaaIpaddressDisable(name, "2001:db8::10", rpZone, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAaaaIpaddressDisable(name, "2001:db8::10", rpZone, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAaaaIpaddressResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_rpz_record_aaaa_ipaddress.test_extattrs"
	var v rpz.RecordRpzAaaaIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "2001:db8::/64" + "." + rpZone
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAaaaIpaddressExtAttrs(name, "2001:db8::10", rpZone, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAaaaIpaddressExtAttrs(name, "2001:db8::10", rpZone, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAaaaIpaddressResource_Ipv6addr(t *testing.T) {
	var resourceName = "nios_rpz_record_aaaa_ipaddress.test_ipv6addr"
	var v rpz.RecordRpzAaaaIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "2001:db8::/64" + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAaaaIpaddressIpv6addr(name, "2001:db8::10", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", "2001:db8::10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAaaaIpaddressIpv6addr(name, "2001:db8::20", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", "2001:db8::20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAaaaIpaddressResource_Name(t *testing.T) {
	var resourceName = "nios_rpz_record_aaaa_ipaddress.test_name"
	var v rpz.RecordRpzAaaaIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name1 := "2001:db8::/64" + "." + rpZone
	name2 := "2001:db8:1::/64" + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAaaaIpaddressName(name1, "2001:db8::10", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAaaaIpaddressName(name2, "2001:db8::10", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAaaaIpaddressResource_RpZone(t *testing.T) {
	var resourceName = "nios_rpz_record_aaaa_ipaddress.test_rp_zone"
	var v rpz.RecordRpzAaaaIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "2001:db8::/64" + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAaaaIpaddressRpZone(name, "2001:db8::10", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
				),
			},
			// Can't update rp_zone as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAaaaIpaddressResource_Ttl(t *testing.T) {
	var resourceName = "nios_rpz_record_aaaa_ipaddress.test_ttl"
	var v rpz.RecordRpzAaaaIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "2001:db8::/64" + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAaaaIpaddressTtl(name, "2001:db8::10", rpZone, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAaaaIpaddressTtl(name, "2001:db8::10", rpZone, "true", 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAaaaIpaddressResource_UseTtl(t *testing.T) {
	var resourceName = "nios_rpz_record_aaaa_ipaddress.test_use_ttl"
	var v rpz.RecordRpzAaaaIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "2001:db8::/64" + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAaaaIpaddressUseTtl(name, "2001:db8::10", rpZone, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAaaaIpaddressUseTtl(name, "2001:db8::10", rpZone, "false", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAaaaIpaddressResource_View(t *testing.T) {
	var resourceName = "nios_rpz_record_aaaa_ipaddress.test_view"
	var v rpz.RecordRpzAaaaIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "2001:db8::/64" + "." + rpZone
	view := acctest.RandomNameWithPrefix("test-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAaaaIpaddressView(name, "2001:db8::10", rpZone, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", view),
				),
			},
			// Can't update view as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordRpzAaaaIpaddressExists(ctx context.Context, resourceName string, v *rpz.RecordRpzAaaaIpaddress) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.RPZAPI.
			RecordRpzAaaaIpaddressAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordRpzAaaaIpaddress).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordRpzAaaaIpaddressResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordRpzAaaaIpaddressResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordRpzAaaaIpaddressDestroy(ctx context.Context, v *rpz.RecordRpzAaaaIpaddress) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.RPZAPI.
			RecordRpzAaaaIpaddressAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordRpzAaaaIpaddress).
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

func testAccCheckRecordRpzAaaaIpaddressDisappears(ctx context.Context, v *rpz.RecordRpzAaaaIpaddress) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.RPZAPI.
			RecordRpzAaaaIpaddressAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordRpzAaaaIpaddressBasicConfig(name, ipV6Addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa_ipaddress" "test" {
	name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, ipV6Addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAaaaIpaddressComment(name, ipV6Addr, rpZone, comment string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa_ipaddress" "test_comment" {
	name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    comment = %q
}
`, name, ipV6Addr, comment)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAaaaIpaddressDisable(name, ipV6Addr, rpZone, disable string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa_ipaddress" "test_disable" {
    name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	disable = %q
}
`, name, ipV6Addr, disable)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAaaaIpaddressExtAttrs(name, ipV6Addr, rpZone string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
	%s = %q
	`, k, v)
	}
	extattrsStr += "\t}"
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa_ipaddress" "test_extattrs" {
    name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	extattrs = %s
}
`, name, ipV6Addr, extattrsStr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAaaaIpaddressIpv6addr(name, ipv6addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa_ipaddress" "test_ipv6addr" {
    name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, ipv6addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAaaaIpaddressName(name, ipv6addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa_ipaddress" "test_name" {
    name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, ipv6addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAaaaIpaddressRpZone(name, ipv6addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa_ipaddress" "test_rp_zone" {
    name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, ipv6addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAaaaIpaddressTtl(name, ipv6addr, rpZone, use_ttl string, ttl int32) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa_ipaddress" "test_ttl" {
    name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	ttl = %d
	use_ttl = %q
}
`, name, ipv6addr, ttl, use_ttl)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAaaaIpaddressUseTtl(name, ipv6addr, rpZone, use_ttl string, ttl int32) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa_ipaddress" "test_use_ttl" {
    name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	ttl = %d
	use_ttl = %q
}
`, name, ipv6addr, ttl, use_ttl)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAaaaIpaddressView(name, ipV6Addr, rpZone, view string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa_ipaddress" "test_view" {
    name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	view = nios_dns_view.custom_view.name
}
`, name, ipV6Addr)

	return strings.Join([]string{testAccBaseWithView(view), testAccBaseWithZone(rpZone, "nios_dns_view.custom_view.name"), config}, "")
}
