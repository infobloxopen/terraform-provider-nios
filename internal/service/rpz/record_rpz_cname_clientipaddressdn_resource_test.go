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

var readableAttributesForRecordRpzCnameClientipaddressdn = "canonical,comment,disable,extattrs,is_ipv4,name,rp_zone,ttl,use_ttl,view,zone"

func TestAccRecordRpzCnameClientipaddressdnResource_basic(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_clientipaddressdn.test"
	var v rpz.RecordRpzCnameClientipaddressdn
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone
	canonical := "test-cname"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameClientipaddressdnBasicConfig(name, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "canonical", canonical),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_ipv4", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameClientipaddressdnResource_disappears(t *testing.T) {
	resourceName := "nios_rpz_record_cname_clientipaddressdn.test"
	var v rpz.RecordRpzCnameClientipaddressdn
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone
	canonical := "test-cname"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzCnameClientipaddressdnDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzCnameClientipaddressdnBasicConfig(name, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressdnExists(context.Background(), resourceName, &v),
					testAccCheckRecordRpzCnameClientipaddressdnDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordRpzCnameClientipaddressdnResource_Canonical(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_clientipaddressdn.test_canonical"
	var v rpz.RecordRpzCnameClientipaddressdn
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone
	canonical1 := "test-cname1"
	canonical2 := "test-cname2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameClientipaddressdnCanonical(name, canonical1, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", canonical1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameClientipaddressdnCanonical(name, canonical2, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", canonical2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameClientipaddressdnResource_Comment(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_clientipaddressdn.test_comment"
	var v rpz.RecordRpzCnameClientipaddressdn
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone
	canonical := "test-cname"
	comment1 := "This is a new rpz cname client ipaddress dn record"
	comment2 := "This is an updated rpz cname client ipaddress dn record"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameClientipaddressdnComment(name, canonical, rpZone, comment1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameClientipaddressdnComment(name, canonical, rpZone, comment2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameClientipaddressdnResource_Disable(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_clientipaddressdn.test_disable"
	var v rpz.RecordRpzCnameClientipaddressdn
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone
	canonical := "test-cname"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameClientipaddressdnDisable(name, canonical, rpZone, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameClientipaddressdnDisable(name, canonical, rpZone, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameClientipaddressdnResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_clientipaddressdn.test_extattrs"
	var v rpz.RecordRpzCnameClientipaddressdn
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone
	canonical := "test-cname"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameClientipaddressdnExtAttrs(name, canonical, rpZone, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameClientipaddressdnExtAttrs(name, canonical, rpZone, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameClientipaddressdnResource_Name(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_clientipaddressdn.test_name"
	var v rpz.RecordRpzCnameClientipaddressdn
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name1 := "10.10.0.0/16" + "." + rpZone
	name2 := "2001:db8::/64" + "." + rpZone
	canonical := "test-cname"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameClientipaddressdnName(name1, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
					resource.TestCheckResourceAttr(resourceName, "is_ipv4", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameClientipaddressdnName(name2, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
					resource.TestCheckResourceAttr(resourceName, "is_ipv4", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameClientipaddressdnResource_RpZone(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_clientipaddressdn.test_rp_zone"
	var v rpz.RecordRpzCnameClientipaddressdn
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone
	canonical := "test-cname"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameClientipaddressdnRpZone(name, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
				),
			},
			// Can't update rp_zone as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameClientipaddressdnResource_Ttl(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_clientipaddressdn.test_ttl"
	var v rpz.RecordRpzCnameClientipaddressdn
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone
	canonical := "test-cname"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameClientipaddressdnTtl(name, canonical, rpZone, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameClientipaddressdnTtl(name, canonical, rpZone, "true", 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameClientipaddressdnResource_UseTtl(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_clientipaddressdn.test_use_ttl"
	var v rpz.RecordRpzCnameClientipaddressdn
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone
	canonical := "test-cname"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameClientipaddressdnUseTtl(name, canonical, rpZone, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameClientipaddressdnUseTtl(name, canonical, rpZone, "false", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameClientipaddressdnResource_View(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_clientipaddressdn.test_view"
	var v rpz.RecordRpzCnameClientipaddressdn
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := "10.10.0.0/16" + "." + rpZone
	canonical := "test-cname"
	view := acctest.RandomNameWithPrefix("test-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameClientipaddressdnView(name, canonical, rpZone, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", view),
				),
			},
			// Can't update view as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordRpzCnameClientipaddressdnExists(ctx context.Context, resourceName string, v *rpz.RecordRpzCnameClientipaddressdn) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.RPZAPI.
			RecordRpzCnameClientipaddressdnAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordRpzCnameClientipaddressdn).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordRpzCnameClientipaddressdnResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordRpzCnameClientipaddressdnResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordRpzCnameClientipaddressdnDestroy(ctx context.Context, v *rpz.RecordRpzCnameClientipaddressdn) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.RPZAPI.
			RecordRpzCnameClientipaddressdnAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordRpzCnameClientipaddressdn).
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

func testAccCheckRecordRpzCnameClientipaddressdnDisappears(ctx context.Context, v *rpz.RecordRpzCnameClientipaddressdn) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.RPZAPI.
			RecordRpzCnameClientipaddressdnAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordRpzCnameClientipaddressdnBasicConfig(name, canonical, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddressdn" "test" {
	name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, canonical)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameClientipaddressdnCanonical(name, canonical, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddressdn" "test_canonical" {
    name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, canonical)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameClientipaddressdnComment(name, canonical, rpZone, comment string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddressdn" "test_comment" {
	name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    comment = %q
}
`, name, canonical, comment)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameClientipaddressdnDisable(name, canonical, rpZone, disable string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddressdn" "test_disable" {
	name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    disable = %q
}
`, name, canonical, disable)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameClientipaddressdnExtAttrs(name, canonical, rpZone string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
	%s = %q
	`, k, v)
	}
	extattrsStr += "\t}"
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddressdn" "test_extattrs" {
    name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	extattrs = %s
}
`, name, canonical, extattrsStr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameClientipaddressdnName(name, canonical, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddressdn" "test_name" {
    name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, canonical)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameClientipaddressdnRpZone(name, canonical, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddressdn" "test_rp_zone" {
    name = %q
	canonical = %q
    rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, canonical)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameClientipaddressdnTtl(name, canonical, rpZone, useTtl string, ttl int32) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddressdn" "test_ttl" {
    name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	ttl = %d
	use_ttl = %q
}
`, name, canonical, ttl, useTtl)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameClientipaddressdnUseTtl(name, canonical, rpZone, useTtl string, ttl int32) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddressdn" "test_use_ttl" {
    name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	ttl = %d
	use_ttl = %q
}
`, name, canonical, ttl, useTtl)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameClientipaddressdnView(name, canonical, rpZone, view string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddressdn" "test_view" {
    name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    view = nios_dns_view.custom_view.name
}
`, name, canonical)

	return strings.Join([]string{testAccBaseWithView(view), testAccBaseWithZone(rpZone, "nios_dns_view.custom_view.name"), config}, "")
}
