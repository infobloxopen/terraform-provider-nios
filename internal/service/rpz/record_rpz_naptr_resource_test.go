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

var readableAttributesForRecordRpzNaptr = "comment,disable,extattrs,flags,last_queried,name,order,preference,regexp,replacement,rp_zone,services,ttl,use_ttl,view,zone"

func TestAccRecordRpzNaptrResource_basic(t *testing.T) {
	var resourceName = "nios_rpz_record_naptr.test"
	var v rpz.RecordRpzNaptr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzNaptrBasicConfig(name, rpZone, ".", 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
					resource.TestCheckResourceAttr(resourceName, "order", "10"),
					resource.TestCheckResourceAttr(resourceName, "preference", "10"),
					resource.TestCheckResourceAttr(resourceName, "replacement", "."),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
					resource.TestCheckResourceAttr(resourceName, "services", ""),
					resource.TestCheckResourceAttr(resourceName, "flags", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzNaptrResource_disappears(t *testing.T) {
	resourceName := "nios_rpz_record_naptr.test"
	var v rpz.RecordRpzNaptr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzNaptrDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzNaptrBasicConfig(name, rpZone, ".", 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					testAccCheckRecordRpzNaptrDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordRpzNaptrResource_Comment(t *testing.T) {
	var resourceName = "nios_rpz_record_naptr.test_comment"
	var v rpz.RecordRpzNaptr
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
				Config: testAccRecordRpzNaptrComment(name, rpZone, ".", comment1, 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzNaptrComment(name, rpZone, ".", comment2, 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzNaptrResource_Disable(t *testing.T) {
	var resourceName = "nios_rpz_record_naptr.test_disable"
	var v rpz.RecordRpzNaptr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzNaptrDisable(name, rpZone, ".", "false", 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzNaptrDisable(name, rpZone, ".", "true", 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzNaptrResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_rpz_record_naptr.test_extattrs"
	var v rpz.RecordRpzNaptr
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
				Config: testAccRecordRpzNaptrExtAttrs(name, rpZone, ".", 10, 10, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzNaptrExtAttrs(name, rpZone, ".", 10, 10, map[string]string{"Site": extAttrValue2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzNaptrResource_Flags(t *testing.T) {
	var resourceName = "nios_rpz_record_naptr.test_flags"
	var v rpz.RecordRpzNaptr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzNaptrFlags(name, rpZone, ".", "U", 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "flags", "U"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzNaptrFlags(name, rpZone, ".", "S", 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "flags", "S"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzNaptrResource_Name(t *testing.T) {
	var resourceName = "nios_rpz_record_naptr.test_name"
	var v rpz.RecordRpzNaptr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name1 := acctest.RandomName() + "." + rpZone
	name2 := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzNaptrName(name1, rpZone, ".", 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzNaptrName(name2, rpZone, ".", 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzNaptrResource_Order(t *testing.T) {
	var resourceName = "nios_rpz_record_naptr.test_order"
	var v rpz.RecordRpzNaptr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzNaptrOrder(name, rpZone, ".", 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "order", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzNaptrOrder(name, rpZone, ".", 20, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "order", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzNaptrResource_Preference(t *testing.T) {
	var resourceName = "nios_rpz_record_naptr.test_preference"
	var v rpz.RecordRpzNaptr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzNaptrPreference(name, rpZone, ".", 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preference", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzNaptrPreference(name, rpZone, ".", 10, 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preference", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzNaptrResource_Regexp(t *testing.T) {
	var resourceName = "nios_rpz_record_naptr.test_regexp"
	var v rpz.RecordRpzNaptr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzNaptrRegexp(name, rpZone, ".", "", 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "regexp", ""),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzNaptrRegexp(name, rpZone, ".", "!^.*$!sip:jdoe@corpxyz.com!", 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "regexp", "!^.*$!sip:jdoe@corpxyz.com!"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzNaptrResource_Replacement(t *testing.T) {
	var resourceName = "nios_rpz_record_naptr.test_replacement"
	var v rpz.RecordRpzNaptr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzNaptrReplacement(name, rpZone, ".", 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "replacement", "."),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzNaptrReplacement(name, rpZone, "test.com", 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "replacement", "test.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzNaptrResource_RpZone(t *testing.T) {
	var resourceName = "nios_rpz_record_naptr.test_rp_zone"
	var v rpz.RecordRpzNaptr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzNaptrRpZone(name, rpZone, ".", 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
				),
			},
			// Can't update rp_zone as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzNaptrResource_Services(t *testing.T) {
	var resourceName = "nios_rpz_record_naptr.test_services"
	var v rpz.RecordRpzNaptr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzNaptrServices(name, rpZone, ".", "http+E2U", 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "services", "http+E2U"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzNaptrServices(name, rpZone, ".", "SIPS+D2T", 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "services", "SIPS+D2T"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzNaptrResource_Ttl(t *testing.T) {
	var resourceName = "nios_rpz_record_naptr.test_ttl"
	var v rpz.RecordRpzNaptr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzNaptrTtl(name, rpZone, ".", "true", 10, 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzNaptrTtl(name, rpZone, ".", "true", 10, 10, 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzNaptrResource_UseTtl(t *testing.T) {
	var resourceName = "nios_rpz_record_naptr.test_use_ttl"
	var v rpz.RecordRpzNaptr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzNaptrUseTtl(name, rpZone, ".", "true", 10, 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzNaptrUseTtl(name, rpZone, ".", "false", 10, 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzNaptrResource_View(t *testing.T) {
	var resourceName = "nios_rpz_record_naptr.test_view"
	var v rpz.RecordRpzNaptr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	view := "custom_view"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzNaptrView(name, rpZone, ".", view, 10, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", view),
				),
			},

			// Can't update view as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordRpzNaptrExists(ctx context.Context, resourceName string, v *rpz.RecordRpzNaptr) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.RPZAPI.
			RecordRpzNaptrAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordRpzNaptr).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordRpzNaptrResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordRpzNaptrResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordRpzNaptrDestroy(ctx context.Context, v *rpz.RecordRpzNaptr) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.RPZAPI.
			RecordRpzNaptrAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordRpzNaptr).
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

func testAccCheckRecordRpzNaptrDisappears(ctx context.Context, v *rpz.RecordRpzNaptr) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.RPZAPI.
			RecordRpzNaptrAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordRpzNaptrBasicConfig(name, rpZone, replacement string, order, preference int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_naptr" "test" {
	name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	order = %d
    preference = %d
    replacement = %q
}
`, name, order, preference, replacement)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzNaptrComment(name, rpZone, replacement, comment string, order, preference int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_naptr" "test_comment" {
	name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	order = %d
    preference = %d
    replacement = %q
    comment = %q
}
`, name, order, preference, replacement, comment)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzNaptrDisable(name, rpZone, replacement, disable string, order, preference int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_naptr" "test_disable" {
	name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	order = %d
    preference = %d
    replacement = %q
    disable = %q
}
`, name, order, preference, replacement, disable)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzNaptrExtAttrs(name, rpZone, replacement string, order, preference int, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
	%s = %q
	`, k, v)
	}
	extattrsStr += "\t}"

	config := fmt.Sprintf(`
resource "nios_rpz_record_naptr" "test_extattrs" {
	name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	order = %d
    preference = %d
    replacement = %q
    extattrs = %s
}
`, name, order, preference, replacement, extattrsStr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzNaptrFlags(name, rpZone, replacement, flags string, order, preference int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_naptr" "test_flags" {
    name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	order = %d
    preference = %d
    replacement = %q
	flags = %q
}
`, name, order, preference, replacement, flags)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzNaptrName(name, rpZone, replacement string, order, preference int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_naptr" "test_name" {
    name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	order = %d
    preference = %d
    replacement = %q
}
`, name, order, preference, replacement)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzNaptrOrder(name, rpZone, replacement string, order, preference int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_naptr" "test_order" {
    name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	order = %d
    preference = %d
    replacement = %q
}
`, name, order, preference, replacement)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzNaptrPreference(name, rpZone, replacement string, order, preference int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_naptr" "test_preference" {
    name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	order = %d
    preference = %d
    replacement = %q
}
`, name, order, preference, replacement)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzNaptrRegexp(name, rpZone, replacement, regexp string, order, preference int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_naptr" "test_regexp" {
    name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	order = %d
    preference = %d
    replacement = %q
	regexp = %q
}
`, name, order, preference, replacement, regexp)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzNaptrReplacement(name, rpZone, replacement string, order, preference int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_naptr" "test_replacement" {
    name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	order = %d
    preference = %d
    replacement = %q
}
`, name, order, preference, replacement)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzNaptrRpZone(name, rpZone, replacement string, order, preference int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_naptr" "test_rp_zone" {
    name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	order = %d
    preference = %d
    replacement = %q
}
`, name, order, preference, replacement)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzNaptrServices(name, rpZone, replacement, services string, order, preference int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_naptr" "test_services" {
    name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	order = %d
    preference = %d
    replacement = %q
	services = %q
}
`, name, order, preference, replacement, services)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzNaptrTtl(name, rpZone, replacement, use_ttl string, order, preference, ttl int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_naptr" "test_ttl" {
    name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	order = %d
    preference = %d
    replacement = %q
	ttl = %d
	use_ttl = %q
}
`, name, order, preference, replacement, ttl, use_ttl)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzNaptrUseTtl(name, rpZone, replacement, use_ttl string, order, preference, ttl int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_naptr" "test_use_ttl" {
    name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	order = %d
    preference = %d
    replacement = %q
	ttl = %d
	use_ttl = %q
}
`, name, order, preference, replacement, ttl, use_ttl)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzNaptrView(name, rpZone, replacement, view string, order, preference int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_naptr" "test_view" {
    name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	order = %d
    preference = %d
    replacement = %q
	view = %q
}
`, name, order, preference, replacement, view)

	return strings.Join([]string{testAccBaseWithZone(rpZone, view), config}, "")
}

func testAccBaseWithZone(zoneFqdn, view string) string {
	if view == "" {
		view = "default"
	}
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test" {
    fqdn = %q
	view = %q
}
`, zoneFqdn, view)
}
