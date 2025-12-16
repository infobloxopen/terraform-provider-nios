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

var readableAttributesForRecordRpzCnameIpaddress = "canonical,comment,disable,extattrs,is_ipv4,name,rp_zone,ttl,use_ttl,view,zone"

func TestAccRecordRpzCnameIpaddressResource_basic(t *testing.T) {
	var resourceName = "nios_rpz_record_rpz_cname_ipaddress.test"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameIpaddressBasicConfig(name, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "canonical", canonical),
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

func TestAccRecordRpzCnameIpaddressResource_disappears(t *testing.T) {
	resourceName := "nios_rpz_record_rpz_cname_ipaddress.test"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzCnameIpaddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzCnameIpaddressBasicConfig(name, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					testAccCheckRecordRpzCnameIpaddressDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordRpzCnameIpaddressResource_Canonical(t *testing.T) {
	var resourceName = "nios_rpz_record_rpz_cname_ipaddress.test_canonical"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	baseName := acctest.RandomName()
	name := baseName + "." + rpZone
	canonical := acctest.RandomNameWithPrefix("test-canonical")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameIpaddressCanonical(name, "", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", ""),
				),
			},
			// Update to No Data rule
			{
				Config: testAccRecordRpzCnameIpaddressCanonical(name, "*", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", "*"),
				),
			},
			// Update to Passthru Domain Name Rule
			{
				Config: testAccRecordRpzCnameIpaddressCanonical(name, baseName, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", baseName),
				),
			},
			// update to Substitution rule
			{
				Config: testAccRecordRpzCnameIpaddressCanonical(name, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", canonical),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameIpaddressResource_Comment(t *testing.T) {
	var resourceName = "nios_rpz_record_rpz_cname_ipaddress.test_comment"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	canonical := ""
	comment1 := "This is a new rpz cname record"
	comment2 := "This is an updated rpz cname record"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameIpaddressComment(name, canonical, rpZone, comment1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameIpaddressComment(name, canonical, rpZone, comment2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameIpaddressResource_Disable(t *testing.T) {
	var resourceName = "nios_rpz_record_rpz_cname_ipaddress.test_disable"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameIpaddressDisable(name, canonical, rpZone, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameIpaddressDisable(name, canonical, rpZone, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameIpaddressResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_rpz_record_rpz_cname_ipaddress.test_extattrs"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	canonical := ""
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameIpaddressExtAttrs(name, canonical, rpZone, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameIpaddressExtAttrs(name, canonical, rpZone, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameIpaddressResource_Name(t *testing.T) {
	var resourceName = "nios_rpz_record_rpz_cname_ipaddress.test_name"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name1 := acctest.RandomName() + "." + rpZone
	name2 := acctest.RandomName() + "." + rpZone
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameIpaddressName(name1, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameIpaddressName(name2, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameIpaddressResource_RpZone(t *testing.T) {
	var resourceName = "nios_rpz_record_rpz_cname_ipaddress.test_rp_zone"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameIpaddressRpZone(name, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
				),
			},
			// Can't update rp_zone as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameIpaddressResource_Ttl(t *testing.T) {
	var resourceName = "nios_rpz_record_rpz_cname_ipaddress.test_ttl"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameIpaddressTtl(name, canonical, rpZone, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameIpaddressTtl(name, canonical, rpZone, "true", 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameIpaddressResource_UseTtl(t *testing.T) {
	var resourceName = "nios_rpz_record_rpz_cname_ipaddress.test_use_ttl"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameIpaddressUseTtl(name, canonical, rpZone, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameIpaddressUseTtl(name, canonical, rpZone, "false", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameIpaddressResource_View(t *testing.T) {
	var resourceName = "nios_rpz_record_rpz_cname_ipaddress.test_view"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	canonical := ""
	view := acctest.RandomNameWithPrefix("test-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameIpaddressView(name, canonical, rpZone, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", view),
				),
			},
			// Can't update view as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordRpzCnameIpaddressExists(ctx context.Context, resourceName string, v *rpz.RecordRpzCnameIpaddress) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.RPZAPI.
			RecordRpzCnameIpaddressAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordRpzCnameIpaddress).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordRpzCnameIpaddressResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordRpzCnameIpaddressResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordRpzCnameIpaddressDestroy(ctx context.Context, v *rpz.RecordRpzCnameIpaddress) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.RPZAPI.
			RecordRpzCnameIpaddressAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordRpzCnameIpaddress).
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

func testAccCheckRecordRpzCnameIpaddressDisappears(ctx context.Context, v *rpz.RecordRpzCnameIpaddress) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.RPZAPI.
			RecordRpzCnameIpaddressAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordRpzCnameIpaddressBasicConfig(name, canonical, rpZone string) string {
	// TODO: create basic resource with required fields
	config := fmt.Sprintf(`
resource "nios_rpz_record_rpz_cname_ipaddress" "test" {
	name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, canonical)
	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameIpaddressCanonical(cname, canonical, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_rpz_cname_ipaddress" "test_canonical" {
    name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn}
`, cname, canonical)
	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameIpaddressComment(name, canonical, rpZone, comment string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_rpz_cname_ipaddress" "test_comment" {
    name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    comment = %q
}
`, name, canonical, comment)
	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameIpaddressDisable(name, canonical, rpZone, disable string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_rpz_cname_ipaddress" "test_disable" {
    name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    disable = %q
}
`, name, canonical, disable)
	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameIpaddressExtAttrs(name, canonical, rpZone string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
	%s = %q
	`, k, v)
	}
	extattrsStr += "\t}"

	config := fmt.Sprintf(`
resource "nios_rpz_record_rpz_cname_ipaddress" "test_extattrs" {
    name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    extattrs = %s
}
`, name, canonical, extattrsStr)
	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameIpaddressName(name, canonical, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_rpz_cname_ipaddress" "test_name" {
    name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, canonical)
	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameIpaddressRpZone(name, canonical, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_rpz_cname_ipaddress" "test_rp_zone" {
	name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, canonical)
	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameIpaddressTtl(name, canonical, rpZone, useTtl string, ttl int32) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_rpz_cname_ipaddress" "test_ttl" {
	name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	ttl = %d
	use_ttl = %q
}
`, name, canonical, ttl, useTtl)
	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameIpaddressUseTtl(uname, canonical, rpZone, useTtl string, ttl int32) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_rpz_cname_ipaddress" "test_use_ttl" {
    name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    use_ttl = %q
	ttl = %d
}
`, uname, canonical, useTtl, ttl)
	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameIpaddressView(name, canonical, rpZone, view string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_rpz_cname_ipaddress" "test_view" {
	name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    view = nios_dns_view.custom_view.name
}
`, name, canonical)
	return strings.Join([]string{testAccBaseWithView(view), testAccBaseWithZone(rpZone, "nios_dns_view.custom_view.name"), config}, "")
}
