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

var readableAttributesForRecordRpzA = "comment,disable,extattrs,ipv4addr,name,rp_zone,ttl,use_ttl,view,zone"

func TestAccRecordRpzAResource_basic(t *testing.T) {
	var resourceName = "nios_rpz_record_a.test"
	var v rpz.RecordRpzA
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzABasicConfig(name, "10.10.0.1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
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

func TestAccRecordRpzAResource_disappears(t *testing.T) {
	resourceName := "nios_rpz_record_a.test"
	var v rpz.RecordRpzA
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzADestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzABasicConfig(name, "10.10.0.1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					testAccCheckRecordRpzADisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordRpzAResource_Comment(t *testing.T) {
	var resourceName = "nios_rpz_record_a.test_comment"
	var v rpz.RecordRpzA
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	comment1 := "This is a new rpz a record"
	comment2 := "This is a updated rpz a record"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAComment(name, "10.10.0.1", rpZone, comment1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAComment(name, "10.10.0.1", rpZone, comment2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAResource_Disable(t *testing.T) {
	var resourceName = "nios_rpz_record_a.test_disable"
	var v rpz.RecordRpzA
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzADisable(name, "10.10.0.1", rpZone, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzADisable(name, "10.10.0.1", rpZone, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_rpz_record_a.test_extattrs"
	var v rpz.RecordRpzA
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
				Config: testAccRecordRpzAExtAttrs(name, "10.10.0.1", rpZone, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAExtAttrs(name, "10.10.0.1", rpZone, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAResource_Ipv4addr(t *testing.T) {
	var resourceName = "nios_rpz_record_a.test_ipv4addr"
	var v rpz.RecordRpzA
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAIpv4addr(name, "10.10.0.1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "10.10.0.1"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAIpv4addr(name, "10.10.0.2", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "10.10.0.2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAResource_Name(t *testing.T) {
	var resourceName = "nios_rpz_record_a.test_name"
	var v rpz.RecordRpzA
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name1 := acctest.RandomName() + "." + rpZone
	name2 := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAName(name1, "10.10.0.1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAName(name2, "10.10.0.1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAResource_RpZone(t *testing.T) {
	var resourceName = "nios_rpz_record_a.test_rp_zone"
	var v rpz.RecordRpzA
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzARpZone(name, "10.10.0.1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
				),
			},
			// Can't update rp_zone as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAResource_Ttl(t *testing.T) {
	var resourceName = "nios_rpz_record_a.test_ttl"
	var v rpz.RecordRpzA
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzATtl(name, "10.10.0.1", rpZone, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzATtl(name, "10.10.0.1", rpZone, "true", 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAResource_UseTtl(t *testing.T) {
	var resourceName = "nios_rpz_record_a.test_use_ttl"
	var v rpz.RecordRpzA
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAUseTtl(name, "10.10.0.1", rpZone, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAUseTtl(name, "10.10.0.1", rpZone, "false", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAResource_View(t *testing.T) {
	var resourceName = "nios_rpz_record_a.test_view"
	var v rpz.RecordRpzA
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	view := acctest.RandomNameWithPrefix("test-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAView(name, "10.10.0.1", rpZone, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", view),
				),
			},
			// Can't update view as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordRpzAExists(ctx context.Context, resourceName string, v *rpz.RecordRpzA) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.RPZAPI.
			RecordRpzAAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordRpzA).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordRpzAResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordRpzAResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordRpzADestroy(ctx context.Context, v *rpz.RecordRpzA) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.RPZAPI.
			RecordRpzAAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordRpzA).
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

func testAccCheckRecordRpzADisappears(ctx context.Context, v *rpz.RecordRpzA) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.RPZAPI.
			RecordRpzAAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordRpzABasicConfig(name, ipV4Addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_a" "test" {
	name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, ipV4Addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAComment(name, ipV4Addr, rpZone, comment string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_a" "test_comment" {
    name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	comment = %q
}
`, name, ipV4Addr, comment)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzADisable(name, ipV4Addr, rpZone, disable string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_a" "test_disable" {
	name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    disable = %q
}
`, name, ipV4Addr, disable)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAExtAttrs(name, ipV4Addr, rpZone string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
	%s = %q
	`, k, v)
	}
	extattrsStr += "\t}"
	config := fmt.Sprintf(`
resource "nios_rpz_record_a" "test_extattrs" {
	name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    extattrs = %s
}
`, name, ipV4Addr, extattrsStr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAIpv4addr(name, ipV4Addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_a" "test_ipv4addr" {
    name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, ipV4Addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAName(name, ipV4Addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_a" "test_name" {
    name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, ipV4Addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzARpZone(name, ipV4Addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_a" "test_rp_zone" {
    name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, ipV4Addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzATtl(name, ipV4Addr, rpZone string, use_ttl string, ttl int32) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_a" "test_ttl" {
    name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	ttl = %d
	use_ttl = %q
}
`, name, ipV4Addr, ttl, use_ttl)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAUseTtl(name, ipV4Addr, rpZone string, use_ttl string, ttl int32) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_a" "test_use_ttl" {
    name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	ttl = %d
	use_ttl = %q
}
`, name, ipV4Addr, ttl, use_ttl)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAView(name, ipV4Addr, rpZone, view string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_a" "test_view" {
    name = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	view = nios_dns_view.custom_view.name
}
`, name, ipV4Addr)

	return strings.Join([]string{testAccBaseWithView(view), testAccBaseWithZone(rpZone, "nios_dns_view.custom_view.name"), config}, "")
}

func testAccBaseWithZone(zoneFqdn, view string) string {
	if view == "" {
		view = `"default"`
	}
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test" {
    fqdn = %q
	view = %s
}
`, zoneFqdn, view)
}

func testAccBaseWithView(view string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "custom_view" {
	name = %q
}
`, view)
}
