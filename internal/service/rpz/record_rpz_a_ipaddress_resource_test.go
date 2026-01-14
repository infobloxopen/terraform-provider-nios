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

var readableAttributesForRecordRpzAIpaddress = "comment,disable,extattrs,ipv4addr,name,rp_zone,ttl,use_ttl,view,zone"

func TestAccRecordRpzAIpaddressResource_basic(t *testing.T) {
	var resourceName = "nios_rpz_record_a_ipaddress.test"
	var v rpz.RecordRpzAIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAIpaddressBasicConfig(name, "10.10.0.1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "10.10.0.1"),
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

func TestAccRecordRpzAIpaddressResource_disappears(t *testing.T) {
	resourceName := "nios_rpz_record_a_ipaddress.test"
	var v rpz.RecordRpzAIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzAIpaddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzAIpaddressBasicConfig(name, "10.10.0.1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAIpaddressExists(context.Background(), resourceName, &v),
					testAccCheckRecordRpzAIpaddressDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordRpzAIpaddressResource_Comment(t *testing.T) {
	var resourceName = "nios_rpz_record_a_ipaddress.test_comment"
	var v rpz.RecordRpzAIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone
	comment1 := "This is a new rpz a ipaddress record"
	comment2 := "This is an updated rpz a ipaddress record"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAIpaddressComment(name, "10.10.0.1", rpZone, comment1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAIpaddressComment(name, "10.10.0.1", rpZone, comment2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAIpaddressResource_Disable(t *testing.T) {
	var resourceName = "nios_rpz_record_a_ipaddress.test_disable"
	var v rpz.RecordRpzAIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAIpaddressDisable(name, "10.10.0.1", rpZone, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAIpaddressDisable(name, "10.10.0.1", rpZone, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAIpaddressResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_rpz_record_a_ipaddress.test_extattrs"
	var v rpz.RecordRpzAIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAIpaddressExtAttrs(name, "10.10.0.1", rpZone, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAIpaddressExtAttrs(name, "10.10.0.1", rpZone, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAIpaddressResource_Ipv4addr(t *testing.T) {
	var resourceName = "nios_rpz_record_a_ipaddress.test_ipv4addr"
	var v rpz.RecordRpzAIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAIpaddressIpv4addr(name, "10.10.0.1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "10.10.0.1"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAIpaddressIpv4addr(name, "10.10.0.2", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "10.10.0.2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAIpaddressResource_Name(t *testing.T) {
	var resourceName = "nios_rpz_record_a_ipaddress.test_name"
	var v rpz.RecordRpzAIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name1 := "10.10.0.0/16" + "." + rpZone
	name2 := "10.15.0.0/16" + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAIpaddressName(name1, "10.10.0.1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAIpaddressName(name2, "10.10.0.1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAIpaddressResource_RpZone(t *testing.T) {
	var resourceName = "nios_rpz_record_a_ipaddress.test_rp_zone"
	var v rpz.RecordRpzAIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAIpaddressRpZone(name, "10.10.0.1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
				),
			},
			// Can't update rp_zone as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAIpaddressResource_Ttl(t *testing.T) {
	var resourceName = "nios_rpz_record_a_ipaddress.test_ttl"
	var v rpz.RecordRpzAIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAIpaddressTtl(name, "10.10.0.1", rpZone, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAIpaddressTtl(name, "10.10.0.1", rpZone, "true", 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAIpaddressResource_UseTtl(t *testing.T) {
	var resourceName = "nios_rpz_record_a_ipaddress.test_use_ttl"
	var v rpz.RecordRpzAIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAIpaddressUseTtl(name, "10.10.0.1", rpZone, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAIpaddressUseTtl(name, "10.10.0.1", rpZone, "false", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAIpaddressResource_View(t *testing.T) {
	var resourceName = "nios_rpz_record_a_ipaddress.test_view"
	var v rpz.RecordRpzAIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone
	view := acctest.RandomNameWithPrefix("test-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAIpaddressView(name, "10.10.0.1", rpZone, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", view),
				),
			},
			// Can't update view as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordRpzAIpaddressExists(ctx context.Context, resourceName string, v *rpz.RecordRpzAIpaddress) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.RPZAPI.
			RecordRpzAIpaddressAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordRpzAIpaddress).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordRpzAIpaddressResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordRpzAIpaddressResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordRpzAIpaddressDestroy(ctx context.Context, v *rpz.RecordRpzAIpaddress) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.RPZAPI.
			RecordRpzAIpaddressAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordRpzAIpaddress).
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

func testAccCheckRecordRpzAIpaddressDisappears(ctx context.Context, v *rpz.RecordRpzAIpaddress) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.RPZAPI.
			RecordRpzAIpaddressAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordRpzAIpaddressBasicConfig(name, ipV4Addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_a_ipaddress" "test" {
	name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, ipV4Addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAIpaddressComment(name, ipV4Addr, rpZone, comment string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_a_ipaddress" "test_comment" {
	name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    comment = %q
}
`, name, ipV4Addr, comment)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAIpaddressDisable(name, ipV4Addr, rpZone, disable string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_a_ipaddress" "test_disable" {
	name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    disable = %q
}
`, name, ipV4Addr, disable)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")

}

func testAccRecordRpzAIpaddressExtAttrs(name, ipV4Addr, rpZone string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
	%s = %q
	`, k, v)
	}
	extattrsStr += "\t}"
	config := fmt.Sprintf(`
resource "nios_rpz_record_a_ipaddress" "test_extattrs" {
	name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    extattrs = %s
}
`, name, ipV4Addr, extattrsStr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAIpaddressIpv4addr(name, ipV4Addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_a_ipaddress" "test_ipv4addr" {
    name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, ipV4Addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAIpaddressName(name, ipV4Addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_a_ipaddress" "test_name" {
    name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, ipV4Addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAIpaddressRpZone(name, ipV4Addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_a_ipaddress" "test_rp_zone" {
    name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, ipV4Addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAIpaddressTtl(name, ipV4Addr, rpZone string, use_ttl string, ttl int32) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_a_ipaddress" "test_ttl" {
	name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    ttl = %d
	use_ttl = %q
}
`, name, ipV4Addr, ttl, use_ttl)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAIpaddressUseTtl(name, ipV4Addr, rpZone string, use_ttl string, ttl int32) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_a_ipaddress" "test_use_ttl" {
    name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    ttl = %d
	use_ttl = %q
}
`, name, ipV4Addr, ttl, use_ttl)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAIpaddressView(name, ipV4Addr, rpZone, view string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_a_ipaddress" "test_view" {
	name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    view = nios_dns_view.custom_view.name
}
`, name, ipV4Addr)

	return strings.Join([]string{testAccBaseWithView(view), testAccBaseWithZone(rpZone, "nios_dns_view.custom_view.name"), config}, "")
}
