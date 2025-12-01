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

var readableAttributesForRecordRpzMx = "comment,disable,extattrs,mail_exchanger,name,preference,rp_zone,ttl,use_ttl,view,zone"

func TestAccRecordRpzMxResource_basic(t *testing.T) {
	var resourceName = "nios_rpz_record_mx.test"
	var v rpz.RecordRpzMx
	rpZone := acctest.RandomNameWithPrefix("rpz") + ".example.com"
	name := acctest.RandomName() + "." + rpZone
	mailExchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzMxBasicConfig(name, mailExchanger, rpZone, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "mail_exchanger", mailExchanger),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
					resource.TestCheckResourceAttr(resourceName, "preference", "10"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzMxResource_disappears(t *testing.T) {
	resourceName := "nios_rpz_record_mx.test"
	var v rpz.RecordRpzMx
	rpZone := acctest.RandomNameWithPrefix("rpz") + ".example.com"
	name := acctest.RandomName() + "." + rpZone
	mailExchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzMxDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzMxBasicConfig(name, mailExchanger, rpZone, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					testAccCheckRecordRpzMxDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordRpzMxResource_Comment(t *testing.T) {
	var resourceName = "nios_rpz_record_mx.test_comment"
	var v rpz.RecordRpzMx
	rpZone := acctest.RandomNameWithPrefix("rpz") + ".example.com"
	name := acctest.RandomName() + "." + rpZone
	mailExchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzMxComment(name, mailExchanger, rpZone, 10, "This is a comment."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment."),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzMxComment(name, mailExchanger, rpZone, 10, "This is an updated comment."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzMxResource_Disable(t *testing.T) {
	var resourceName = "nios_rpz_record_mx.test_disable"
	var v rpz.RecordRpzMx
	rpZone := acctest.RandomNameWithPrefix("rpz") + ".example.com"
	name := acctest.RandomName() + "." + rpZone
	mailExchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzMxDisable(name, mailExchanger, rpZone, 10, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzMxDisable(name, mailExchanger, rpZone, 10, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzMxResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_rpz_record_mx.test_extattrs"
	var v rpz.RecordRpzMx
	rpZone := acctest.RandomNameWithPrefix("rpz") + ".example.com"
	name := acctest.RandomName() + "." + rpZone
	mailExchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzMxExtAttrs(name, mailExchanger, rpZone, 10, map[string]any{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzMxExtAttrs(name, mailExchanger, rpZone, 10, map[string]any{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzMxResource_MailExchanger(t *testing.T) {
	var resourceName = "nios_rpz_record_mx.test_mail_exchanger"
	var v rpz.RecordRpzMx
	rpZone := acctest.RandomNameWithPrefix("rpz") + ".example.com"
	name := acctest.RandomName() + "." + rpZone
	mailExchanger1 := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
	mailExchanger2 := acctest.RandomNameWithPrefix("updatedmail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzMxMailExchanger(name, mailExchanger1, rpZone, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mail_exchanger", mailExchanger1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzMxMailExchanger(name, mailExchanger2, rpZone, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mail_exchanger", mailExchanger2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzMxResource_Name(t *testing.T) {
	var resourceName = "nios_rpz_record_mx.test_name"
	var v rpz.RecordRpzMx
	rpZone := acctest.RandomNameWithPrefix("rpz") + ".example.com"
	name1 := acctest.RandomName() + "." + rpZone
	name2 := acctest.RandomName() + "." + rpZone
	mailExchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzMxName(name1, mailExchanger, rpZone, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzMxName(name2, mailExchanger, rpZone, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzMxResource_Preference(t *testing.T) {
	var resourceName = "nios_rpz_record_mx.test_preference"
	var v rpz.RecordRpzMx
	rpZone := acctest.RandomNameWithPrefix("rpz") + ".example.com"
	name := acctest.RandomName() + "." + rpZone
	mailExchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzMxPreference(name, mailExchanger, rpZone, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preference", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzMxPreference(name, mailExchanger, rpZone, 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preference", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzMxResource_RpZone(t *testing.T) {
	var resourceName = "nios_rpz_record_mx.test_rp_zone"
	var v rpz.RecordRpzMx
	rpZone := acctest.RandomNameWithPrefix("rpz1") + ".example.com"
	name := acctest.RandomName() + "." + rpZone
	mailExchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzMxRpZone(name, mailExchanger, rpZone, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzMxResource_Ttl(t *testing.T) {
	var resourceName = "nios_rpz_record_mx.test_ttl"
	var v rpz.RecordRpzMx
	rpZone := acctest.RandomNameWithPrefix("rpz") + ".example.com"
	name := acctest.RandomName() + "." + rpZone
	mailExchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzMxTtl(name, mailExchanger, rpZone, 10, 3600, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "3600"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzMxTtl(name, mailExchanger, rpZone, 20, 7200, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "7200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzMxResource_UseTtl(t *testing.T) {
	var resourceName = "nios_rpz_record_mx.test_use_ttl"
	var v rpz.RecordRpzMx
	rpZone := acctest.RandomNameWithPrefix("rpz") + ".example.com"
	name := acctest.RandomName() + "." + rpZone
	mailExchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzMxUseTtl(name, mailExchanger, rpZone, 10, 3600, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzMxUseTtl(name, mailExchanger, rpZone, 20, 3600, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzMxResource_View(t *testing.T) {
	var resourceName = "nios_rpz_record_mx.test_view"
	var v rpz.RecordRpzMx
	rpZone := acctest.RandomNameWithPrefix("rpz") + ".example.com"
	name := acctest.RandomName() + "." + rpZone
	mailExchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzMxView(name, mailExchanger, rpZone, 10, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordRpzMxExists(ctx context.Context, resourceName string, v *rpz.RecordRpzMx) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.RPZAPI.
			RecordRpzMxAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordRpzMx).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordRpzMxResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordRpzMxResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordRpzMxDestroy(ctx context.Context, v *rpz.RecordRpzMx) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.RPZAPI.
			RecordRpzMxAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordRpzMx).
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

func testAccCheckRecordRpzMxDisappears(ctx context.Context, v *rpz.RecordRpzMx) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.RPZAPI.
			RecordRpzMxAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}
func testAccRecordRpzMxBasicConfig(name, mailExchanger, rpZone string, preference int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_mx" "test" {
    name           = %q
    mail_exchanger = %q
    preference     = %d
    rp_zone        = nios_dns_zone_rp.test.fqdn
}
`, name, mailExchanger, preference)

	return strings.Join([]string{testAccBaseWithZone(rpZone), config}, "")
}

func testAccRecordRpzMxComment(name, mailExchanger, rpZone string, preference int, comment string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_mx" "test_comment" {
    name           = %q
    mail_exchanger = %q
    rp_zone        = nios_dns_zone_rp.test.fqdn
    preference     = %d
    comment        = %q
}
`, name, mailExchanger, preference, comment)

	return strings.Join([]string{testAccBaseWithZone(rpZone), config}, "")
}

func testAccRecordRpzMxDisable(name, mailExchanger, rpZone string, preference int, disable bool) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_mx" "test_disable" {
    name           = %q
    mail_exchanger = %q
    preference     = %d
    rp_zone        = nios_dns_zone_rp.test.fqdn
    disable        = %t
}
`, name, mailExchanger, preference, disable)

	return strings.Join([]string{testAccBaseWithZone(rpZone), config}, "")
}

func testAccRecordRpzMxExtAttrs(name, mailExchanger, rpZone string, preference int, extAttrs map[string]any) string {
	extAttrsStr := utils.ConvertMapToHCL(extAttrs)
	config := fmt.Sprintf(`
resource "nios_rpz_record_mx" "test_extattrs" {
    name           = %q
    mail_exchanger = %q
    preference     = %d
    rp_zone        = nios_dns_zone_rp.test.fqdn
    extattrs       = %s
}
`, name, mailExchanger, preference, extAttrsStr)

	return strings.Join([]string{testAccBaseWithZone(rpZone), config}, "")
}

func testAccRecordRpzMxMailExchanger(name, mailExchanger, rpZone string, preference int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_mx" "test_mail_exchanger" {
    name           = %q
    mail_exchanger = %q
    preference     = %d
    rp_zone        = nios_dns_zone_rp.test.fqdn
}
`, name, mailExchanger, preference)

	return strings.Join([]string{testAccBaseWithZone(rpZone), config}, "")
}

func testAccRecordRpzMxName(name, mailExchanger, rpZone string, preference int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_mx" "test_name" {
    name           = %q
    mail_exchanger = %q
    preference     = %d
    rp_zone        = nios_dns_zone_rp.test.fqdn
}
`, name, mailExchanger, preference)

	return strings.Join([]string{testAccBaseWithZone(rpZone), config}, "")
}

func testAccRecordRpzMxPreference(name, mailExchanger, rpZone string, preference int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_mx" "test_preference" {
    name           = %q
    mail_exchanger = %q
    preference     = %d
    rp_zone        = nios_dns_zone_rp.test.fqdn
}
`, name, mailExchanger, preference)

	return strings.Join([]string{testAccBaseWithZone(rpZone), config}, "")
}

func testAccRecordRpzMxRpZone(name, mailExchanger, rpZone string, preference int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_mx" "test_rp_zone" {
    name           = %q
    mail_exchanger = %q
    rp_zone        = nios_dns_zone_rp.test.fqdn
    preference     = %d
}
`, name, mailExchanger, preference)

	return strings.Join([]string{testAccBaseWithZone(rpZone), config}, "")
}

func testAccRecordRpzMxTtl(name, mailExchanger, rpZone string, preference int, ttl int32, useTtl bool) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_mx" "test_ttl" {
    name           = %q
    mail_exchanger = %q
    preference     = %d
    rp_zone        = nios_dns_zone_rp.test.fqdn
    ttl            = %d
    use_ttl        = %t
}
`, name, mailExchanger, preference, ttl, useTtl)

	return strings.Join([]string{testAccBaseWithZone(rpZone), config}, "")
}

func testAccRecordRpzMxUseTtl(name, mailExchanger, rpZone string, preference int, ttl int32, useTtl bool) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_mx" "test_use_ttl" {
    name           = %q
    mail_exchanger = %q
    preference     = %d
    rp_zone        = nios_dns_zone_rp.test.fqdn
    ttl            = %d
    use_ttl        = %t
}
`, name, mailExchanger, preference, ttl, useTtl)

	return strings.Join([]string{testAccBaseWithZone(rpZone), config}, "")
}

func testAccRecordRpzMxView(name, mailExchanger, rpZone string, preference int, view string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_mx" "test_view" {
    name           = %q
    mail_exchanger = %q
    preference     = %d
    rp_zone        = nios_dns_zone_rp.test.fqdn
    view           = %q
}`, name, mailExchanger, preference, view)

	return strings.Join([]string{testAccBaseWithZone(rpZone), config}, "")
}

func testAccBaseWithZone(zoneFqdn string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test" {
    fqdn = %q
}
`, zoneFqdn)
}
