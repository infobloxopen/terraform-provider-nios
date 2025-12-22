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

var readableAttributesForRecordRpzTxt = "comment,disable,extattrs,name,rp_zone,text,ttl,use_ttl,view,zone"

func TestAccRecordRpzTxtResource_basic(t *testing.T) {
	var resourceName = "nios_rpz_record_txt.test"
	var v rpz.RecordRpzTxt
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	txt := "Record Text"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzTxtBasicConfig(name, rpZone, txt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "text", txt),
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

func TestAccRecordRpzTxtResource_disappears(t *testing.T) {
	resourceName := "nios_rpz_record_txt.test"
	var v rpz.RecordRpzTxt
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	txt := "Record Text"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzTxtDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzTxtBasicConfig(name, rpZone, txt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzTxtExists(context.Background(), resourceName, &v),
					testAccCheckRecordRpzTxtDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordRpzTxtResource_Comment(t *testing.T) {
	var resourceName = "nios_rpz_record_txt.test_comment"
	var v rpz.RecordRpzTxt
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	txt := "Record Text"
	comment1 := "This is a new rpz txt record"
	comment2 := "This is an updated rpz txt record"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzTxtComment(name, rpZone, txt, comment1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzTxtComment(name, rpZone, txt, comment2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzTxtResource_Disable(t *testing.T) {
	var resourceName = "nios_rpz_record_txt.test_disable"
	var v rpz.RecordRpzTxt
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	txt := "Record Text"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzTxtDisable(name, rpZone, txt, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzTxtDisable(name, rpZone, txt, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzTxtResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_rpz_record_txt.test_extattrs"
	var v rpz.RecordRpzTxt
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	txt := "Record Text"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzTxtExtAttrs(name, rpZone, txt, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzTxtExtAttrs(name, rpZone, txt, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzTxtResource_Name(t *testing.T) {
	var resourceName = "nios_rpz_record_txt.test_name"
	var v rpz.RecordRpzTxt
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name1 := acctest.RandomName() + "." + rpZone
	name2 := acctest.RandomName() + "." + rpZone
	txt := "Record Text"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzTxtName(name1, rpZone, txt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzTxtName(name2, rpZone, txt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzTxtResource_RpZone(t *testing.T) {
	var resourceName = "nios_rpz_record_txt.test_rp_zone"
	var v rpz.RecordRpzTxt
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	txt := "Record Text"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzTxtRpZone(name, rpZone, txt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
				),
			},
			// Can't update rp_zone as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzTxtResource_Text(t *testing.T) {
	var resourceName = "nios_rpz_record_txt.test_text"
	var v rpz.RecordRpzTxt
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	txt1 := "Record Text"
	txt2 := "Updated Record Text"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzTxtText(name, rpZone, txt1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "text", txt1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzTxtText(name, rpZone, txt2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "text", txt2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzTxtResource_Ttl(t *testing.T) {
	var resourceName = "nios_rpz_record_txt.test_ttl"
	var v rpz.RecordRpzTxt
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	txt := "Record Text"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzTxtTtl(name, rpZone, txt, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzTxtTtl(name, rpZone, txt, "true", 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzTxtResource_UseTtl(t *testing.T) {
	var resourceName = "nios_rpz_record_txt.test_use_ttl"
	var v rpz.RecordRpzTxt
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	txt := "Record Text"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzTxtUseTtl(name, rpZone, txt, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzTxtUseTtl(name, rpZone, txt, "false", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzTxtResource_View(t *testing.T) {
	var resourceName = "nios_rpz_record_txt.test_view"
	var v rpz.RecordRpzTxt
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	txt := "Record Text"
	view := acctest.RandomNameWithPrefix("test-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzTxtView(name, rpZone, txt, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", view),
				),
			},
			// Can't update view as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordRpzTxtExists(ctx context.Context, resourceName string, v *rpz.RecordRpzTxt) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.RPZAPI.
			RecordRpzTxtAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordRpzTxt).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordRpzTxtResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordRpzTxtResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordRpzTxtDestroy(ctx context.Context, v *rpz.RecordRpzTxt) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.RPZAPI.
			RecordRpzTxtAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordRpzTxt).
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

func testAccCheckRecordRpzTxtDisappears(ctx context.Context, v *rpz.RecordRpzTxt) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.RPZAPI.
			RecordRpzTxtAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordRpzTxtBasicConfig(name, rpZone, txt string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_txt" "test" {
	name = %q
	text = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, txt)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzTxtComment(name, rpZone, txt, comment string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_txt" "test_comment" {
	name = %q
	text = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    comment = %q
}
`, name, txt, comment)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzTxtDisable(name, rpZone, txt, disable string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_txt" "test_disable" {
	name = %q
	text = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	disable = %q
}
`, name, txt, disable)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzTxtExtAttrs(name, rpZone, txt string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
	%s = %q
	`, k, v)
	}
	extattrsStr += "\t}"
	config := fmt.Sprintf(`
resource "nios_rpz_record_txt" "test_extattrs" {
	name = %q
	text = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    extattrs = %s
}
`, name, txt, extattrsStr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzTxtName(name, rpZone, txt string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_txt" "test_name" {
    name = %q
	text = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, txt)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzTxtRpZone(name, rpZone, txt string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_txt" "test_rp_zone" {
    name = %q
	text = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, txt)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzTxtText(name, rpZone, txt string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_txt" "test_text" {
    name = %q
	text = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, txt)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzTxtTtl(name, rpZone, txt, use_ttl string, ttl int32) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_txt" "test_ttl" {
	name = %q
	text = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    ttl = %d
	use_ttl = %q
}
`, name, txt, ttl, use_ttl)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzTxtUseTtl(name, rpZone, txt, use_ttl string, ttl int32) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_txt" "test_use_ttl" {
    name = %q
	text = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    ttl = %d
	use_ttl = %q
}
`, name, txt, ttl, use_ttl)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzTxtView(name, rpZone, txt, view string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_txt" "test_view" {
	name = %q
	text = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    view = nios_dns_view.custom_view.name
}
`, name, txt)

	return strings.Join([]string{testAccBaseWithView(view), testAccBaseWithZone(rpZone, "nios_dns_view.custom_view.name"), config}, "")
}
