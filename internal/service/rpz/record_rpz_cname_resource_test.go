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

var readableAttributesForRecordRpzCname = "canonical,comment,disable,extattrs,name,rp_zone,ttl,use_ttl,view,zone"

func TestAccRecordRpzCnameResource_basic(t *testing.T) {
	var resourceName = "nios_rpz_record_cname.test"
	var v rpz.RecordRpzCname
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameBasicConfig(name, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameExists(context.Background(), resourceName, &v),
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

func TestAccRecordRpzCnameResource_disappears(t *testing.T) {
	resourceName := "nios_rpz_record_cname.test"
	var v rpz.RecordRpzCname
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzCnameDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzCnameBasicConfig(name, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameExists(context.Background(), resourceName, &v),
					testAccCheckRecordRpzCnameDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordRpzCnameResource_Canonical(t *testing.T) {
	var resourceName = "nios_rpz_record_cname.test_canonical"
	var v rpz.RecordRpzCname
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameCanonical(name, "", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", ""),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameCanonical(name, "*", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", "*"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameResource_Comment(t *testing.T) {
	var resourceName = "nios_rpz_record_cname.test_comment"
	var v rpz.RecordRpzCname
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
				Config: testAccRecordRpzCnameComment(name, canonical, rpZone, comment1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameComment(name, canonical, rpZone, comment2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameResource_Disable(t *testing.T) {
	var resourceName = "nios_rpz_record_cname.test_disable"
	var v rpz.RecordRpzCname
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameDisable(name, canonical, rpZone, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameDisable(name, canonical, rpZone, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_rpz_record_cname.test_extattrs"
	var v rpz.RecordRpzCname
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
				Config: testAccRecordRpzCnameExtAttrs(name, canonical, rpZone, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameExtAttrs(name, canonical, rpZone, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameResource_Name(t *testing.T) {
	var resourceName = "nios_rpz_record_cname.test_name"
	var v rpz.RecordRpzCname
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
				Config: testAccRecordRpzCnameName(name1, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameName(name2, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameResource_RpZone(t *testing.T) {
	var resourceName = "nios_rpz_record_cname.test_rp_zone"
	var v rpz.RecordRpzCname
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameRpZone(name, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
				),
			},
			// Can't update rp_zone as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameResource_Ttl(t *testing.T) {
	var resourceName = "nios_rpz_record_cname.test_ttl"
	var v rpz.RecordRpzCname
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameTtl(name, canonical, rpZone, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameTtl(name, canonical, rpZone, "true", 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameResource_UseTtl(t *testing.T) {
	var resourceName = "nios_rpz_record_cname.test_use_ttl"
	var v rpz.RecordRpzCname
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameUseTtl(name, canonical, rpZone, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameUseTtl(name, canonical, rpZone, "false", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameResource_View(t *testing.T) {
	var resourceName = "nios_rpz_record_cname.test_view"
	var v rpz.RecordRpzCname
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	canonical := ""
	view := "custom_view"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameView(name, canonical, rpZone, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", view),
				),
			},
			// Can't update view as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordRpzCnameExists(ctx context.Context, resourceName string, v *rpz.RecordRpzCname) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.RPZAPI.
			RecordRpzCnameAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordRpzCname).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordRpzCnameResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordRpzCnameResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordRpzCnameDestroy(ctx context.Context, v *rpz.RecordRpzCname) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.RPZAPI.
			RecordRpzCnameAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordRpzCname).
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

func testAccCheckRecordRpzCnameDisappears(ctx context.Context, v *rpz.RecordRpzCname) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.RPZAPI.
			RecordRpzCnameAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordRpzCnameBasicConfig(name, canonical, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname" "test" {
	name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, canonical)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameCanonical(name, canonical, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname" "test_canonical" {
    name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, canonical)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameComment(name, canonical, rpZone, comment string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname" "test_comment" {
	name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    comment = %q
}
`, name, canonical, comment)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameDisable(name, canonical, rpZone, disable string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname" "test_disable" {
	name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    disable = %q
}
`, name, canonical, disable)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameExtAttrs(name, canonical, rpZone string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
	%s = %q
	`, k, v)
	}
	extattrsStr += "\t}"
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname" "test_extattrs" {
	name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    extattrs = %s
}
`, name, canonical, extattrsStr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameName(name, canonical, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname" "test_name" {
    name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, canonical)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameRpZone(name, canonical, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname" "test_rp_zone" {
    name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, name, canonical)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameTtl(name, canonical, rpZone, useTtl string, ttl int32) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname" "test_ttl" {
    name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	ttl = %d
	use_ttl = %q

}
`, name, canonical, ttl, useTtl)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameUseTtl(name, canonical, rpZone, useTtl string, ttl int32) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname" "test_use_ttl" {
    name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	ttl = %d
	use_ttl = %q
}
`, useTtl)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameView(name, canonical, rpZone, view string) string {
	return fmt.Sprintf(`
resource "nios_rpz_record_cname" "test_view" {
	name = %q
	canonical = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    view = %q
}
`, name, canonical, rpZone, view)
}
