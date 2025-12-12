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

var readableAttributesForRecordRpzAaaa = "comment,disable,extattrs,ipv6addr,name,rp_zone,ttl,use_ttl,view,zone"

func TestAccRecordRpzAaaaResource_basic(t *testing.T) {
	var resourceName = "nios_rpz_record_aaaa.test"
	var v rpz.RecordRpzAaaa
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAaaaBasicConfig(name, "2002:1f93::12:1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", "2002:1f93::12:1"),
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

func TestAccRecordRpzAaaaResource_disappears(t *testing.T) {
	resourceName := "nios_rpz_record_aaaa.test"
	var v rpz.RecordRpzAaaa
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzAaaaDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzAaaaBasicConfig(name, "2002:1f93::12:1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaExists(context.Background(), resourceName, &v),
					testAccCheckRecordRpzAaaaDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordRpzAaaaResource_Comment(t *testing.T) {
	var resourceName = "nios_rpz_record_aaaa.test_comment"
	var v rpz.RecordRpzAaaa
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	comment1 := "This is a new rpz aaaa record"
	comment2 := "This is an updated rpz aaaa record"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAaaaComment(name, "2002:1f93::12:1", rpZone, comment1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAaaaComment(name, "2002:1f93::12:1", rpZone, comment2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAaaaResource_Disable(t *testing.T) {
	var resourceName = "nios_rpz_record_aaaa.test_disable"
	var v rpz.RecordRpzAaaa
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAaaaDisable(name, "2002:1f93::12:1", rpZone, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAaaaDisable(name, "2002:1f93::12:1", rpZone, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAaaaResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_rpz_record_aaaa.test_extattrs"
	var v rpz.RecordRpzAaaa
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAaaaExtAttrs(name, "2002:1f93::12:1", rpZone, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAaaaExtAttrs(name, "2002:1f93::12:1", rpZone, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAaaaResource_Ipv6addr(t *testing.T) {
	var resourceName = "nios_rpz_record_aaaa.test_ipv6addr"
	var v rpz.RecordRpzAaaa
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAaaaIpv6addr(name, "2002:1f93::12:1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", "2002:1f93::12:1"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAaaaIpv6addr(name, "2002:1f93::12:10", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", "2002:1f93::12:10"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAaaaResource_Name(t *testing.T) {
	var resourceName = "nios_rpz_record_aaaa.test_name"
	var v rpz.RecordRpzAaaa
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name1 := acctest.RandomName() + "." + rpZone
	name2 := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAaaaName(name1, "2002:1f93::12:1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAaaaName(name2, "2002:1f93::12:1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAaaaResource_RpZone(t *testing.T) {
	var resourceName = "nios_rpz_record_aaaa.test_rp_zone"
	var v rpz.RecordRpzAaaa
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAaaaRpZone(name, "2002:1f93::12:1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
				),
			},
			// Can't update rp_zone as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAaaaResource_Ttl(t *testing.T) {
	var resourceName = "nios_rpz_record_aaaa.test_ttl"
	var v rpz.RecordRpzAaaa
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAaaaTtl(name, "2002:1f93::12:1", rpZone, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAaaaTtl(name, "2002:1f93::12:1", rpZone, "true", 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAaaaResource_UseTtl(t *testing.T) {
	var resourceName = "nios_rpz_record_aaaa.test_use_ttl"
	var v rpz.RecordRpzAaaa
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAaaaUseTtl(name, "2002:1f93::12:1", rpZone, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAaaaUseTtl(name, "2002:1f93::12:1", rpZone, "false", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAaaaResource_View(t *testing.T) {
	var resourceName = "nios_rpz_record_aaaa.test_view"
	var v rpz.RecordRpzAaaa
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	view := acctest.RandomNameWithPrefix("test-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAaaaView(name, "2002:1f93::12:1", rpZone, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", view),
				),
			},
			// Can't update view as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordRpzAaaaExists(ctx context.Context, resourceName string, v *rpz.RecordRpzAaaa) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.RPZAPI.
			RecordRpzAaaaAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordRpzAaaa).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordRpzAaaaResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordRpzAaaaResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordRpzAaaaDestroy(ctx context.Context, v *rpz.RecordRpzAaaa) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.RPZAPI.
			RecordRpzAaaaAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordRpzAaaa).
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

func testAccCheckRecordRpzAaaaDisappears(ctx context.Context, v *rpz.RecordRpzAaaa) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.RPZAPI.
			RecordRpzAaaaAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordRpzAaaaBasicConfig(name, ipV6Addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa" "test" {
	name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, ipV6Addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAaaaComment(name, ipV6Addr, rpZone, comment string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa" "test_comment" {
	name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    comment = %q
}
`, name, ipV6Addr, comment)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAaaaDisable(name, ipV6Addr, rpZone, disable string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa" "test_disable" {
	name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    disable = %q
}
`, name, ipV6Addr, disable)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAaaaExtAttrs(name, ipV6Addr, rpZone string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
	%s = %q
	`, k, v)
	}
	extattrsStr += "\t}"
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa" "test_extattrs" {
	name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    extattrs = %s
}
`, name, ipV6Addr, extattrsStr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAaaaIpv6addr(name, ipV6Addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa" "test_ipv6addr" {
    name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, ipV6Addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAaaaName(name, ipV6Addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa" "test_name" {
    name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, ipV6Addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAaaaRpZone(name, ipV6Addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa" "test_rp_zone" {
    name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, ipV6Addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAaaaTtl(name, ipV6Addr, rpZone string, use_ttl string, ttl int32) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa" "test_ttl" {
	name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	ttl = %d
	use_ttl = %q
}
`, name, ipV6Addr, ttl, use_ttl)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAaaaUseTtl(name, ipV6Addr, rpZone string, use_ttl string, ttl int32) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa" "test_use_ttl" {
    name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	ttl = %d
	use_ttl = %q
}
`, name, ipV6Addr, ttl, use_ttl)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAaaaView(name, ipV6Addr, rpZone, view string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa" "test_view" {
    name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	view = nios_dns_view.custom_view.name

}
`, name, ipV6Addr)

	return strings.Join([]string{testAccBaseWithView(view), testAccBaseWithZone(rpZone, "nios_dns_view.custom_view.name"), config}, "")
}
