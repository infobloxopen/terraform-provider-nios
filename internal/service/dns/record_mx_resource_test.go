package dns_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// TODO: OBJECTS TO BE PRESENT IN GRID FOR TESTS
// - Parent Zone: example.com (in default view)

var readableAttributesForRecordMx = "aws_rte53_record_info,cloud_info,comment,creation_time,creator,ddns_principal,ddns_protected,disable,dns_mail_exchanger,dns_name,extattrs,forbid_reclamation,last_queried,mail_exchanger,name,preference,reclaimable,shared_record_group,ttl,use_ttl,view,zone"

func TestAccRecordMxResource_basic(t *testing.T) {
	var resourceName = "nios_dns_record_mx.test"
	var v dns.RecordMx
	name := acctest.RandomNameWithPrefix("record-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordMxBasicConfig(name, mail_exchanger, 10, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mail_exchanger", mail_exchanger),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "preference", "10"),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "false"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordMxResource_disappears(t *testing.T) {
	resourceName := "nios_dns_record_mx.test"
	var v dns.RecordMx
	name := acctest.RandomNameWithPrefix("record-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordMxDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordMxBasicConfig(name, mail_exchanger, 10, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					testAccCheckRecordMxDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordMxResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_record_mx.test_comment"
	var v dns.RecordMx
	name := acctest.RandomNameWithPrefix("record-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordMxComment(name, mail_exchanger, 10, "default", "This is a comment."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment."),
				),
			},
			// Update and Read
			{
				Config: testAccRecordMxComment(name, mail_exchanger, 10, "default", "This is an updated comment."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordMxResource_Creator(t *testing.T) {
	var resourceName = "nios_dns_record_mx.test_creator"
	var v dns.RecordMx
	name := acctest.RandomNameWithPrefix("record-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordMxCreator(name, mail_exchanger, 10, "default", "STATIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordMxCreator(name, mail_exchanger, 10, "default", "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "DYNAMIC"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordMxResource_DdnsProtected(t *testing.T) {
	var resourceName = "nios_dns_record_mx.test_ddns_protected"
	var v dns.RecordMx
	name := acctest.RandomNameWithPrefix("record-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordMxDdnsProtected(name, mail_exchanger, 10, "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordMxDdnsProtected(name, mail_exchanger, 10, "default", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordMxResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_record_mx.test_disable"
	var v dns.RecordMx
	name := acctest.RandomNameWithPrefix("record-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordMxDisable(name, mail_exchanger, 10, "default", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordMxDisable(name, mail_exchanger, 10, "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordMxResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_record_mx.test_extattrs"
	var v dns.RecordMx
	name := acctest.RandomNameWithPrefix("record-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordMxExtAttrs(name, mail_exchanger, 10, "default", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordMxExtAttrs(name, mail_exchanger, 10, "default", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordMxResource_ForbidReclamation(t *testing.T) {
	var resourceName = "nios_dns_record_mx.test_forbid_reclamation"
	var v dns.RecordMx
	name := acctest.RandomNameWithPrefix("record-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordMxForbidReclamation(name, mail_exchanger, 10, "default", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordMxForbidReclamation(name, mail_exchanger, 10, "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordMxResource_MailExchanger(t *testing.T) {
	var resourceName = "nios_dns_record_mx.test_mail_exchanger"
	var v dns.RecordMx
	name := acctest.RandomNameWithPrefix("record-mx") + ".example.com"
	mail_exchanger1 := acctest.RandomNameWithPrefix("mail-exchanger1") + ".example.com"
	mail_exchanger2 := acctest.RandomNameWithPrefix("mail-exchanger2") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordMxMailExchanger(name, mail_exchanger1, 10, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mail_exchanger", mail_exchanger1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordMxMailExchanger(name, mail_exchanger2, 10, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mail_exchanger", mail_exchanger2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordMxResource_Name(t *testing.T) {
	var resourceName = "nios_dns_record_mx.test_name"
	var v dns.RecordMx
	name1 := acctest.RandomName() + ".example.com"
	name2 := acctest.RandomName() + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordMxName(name1, mail_exchanger, 10, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordMxName(name2, mail_exchanger, 10, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordMxResource_Preference(t *testing.T) {
	var resourceName = "nios_dns_record_mx.test_preference"
	var v dns.RecordMx
	name := acctest.RandomNameWithPrefix("record-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordMxPreference(name, mail_exchanger, 10, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preference", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordMxPreference(name, mail_exchanger, 20, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preference", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordMxResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_record_mx.test_ttl"
	var v dns.RecordMx
	name := acctest.RandomNameWithPrefix("record-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordMxTtl(name, mail_exchanger, 10, "default", 10, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordMxTtl(name, mail_exchanger, 0, "default", 0, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordMxResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_record_mx.test_use_ttl"
	var v dns.RecordMx
	name := acctest.RandomNameWithPrefix("record-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordMxUseTtl(name, mail_exchanger, 10, "default", "false", 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordMxUseTtl(name, mail_exchanger, 10, "default", "true", 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordMxResource_View(t *testing.T) {
	var resourceName = "nios_dns_record_mx.test_view"
	var v dns.RecordMx
	name := acctest.RandomNameWithPrefix("record-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordMxView(name, mail_exchanger, 10, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordMxView(name, mail_exchanger, 10, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordMxExists(ctx context.Context, resourceName string, v *dns.RecordMx) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			RecordMxAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordMx).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordMxResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordMxResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordMxDestroy(ctx context.Context, v *dns.RecordMx) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			RecordMxAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordMx).
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

func testAccCheckRecordMxDisappears(ctx context.Context, v *dns.RecordMx) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			RecordMxAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}
func testAccRecordMxBasicConfig(name, mail_exchanger string, preference int64, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_mx" "test" {
    name            = %q
    mail_exchanger  = %q
    preference      = %d
    view            = %q
}
`, name, mail_exchanger, preference, view)
}

func testAccRecordMxComment(name, mail_exchanger string, preference int64, view, comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_mx" "test_comment" {
    name            = %q
    mail_exchanger  = %q
    preference      = %d
    view            = %q
    comment         = %q
}
`, name, mail_exchanger, preference, view, comment)
}

func testAccRecordMxCreator(name, mail_exchanger string, preference int64, view, creator string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_mx" "test_creator" {
    name            = %q
    mail_exchanger  = %q
    preference      = %d
    view            = %q
    creator         = %q
}
`, name, mail_exchanger, preference, view, creator)
}

func testAccRecordMxDdnsProtected(name, mail_exchanger string, preference int64, view, ddnsProtected string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_mx" "test_ddns_protected" {
    name             = %q
    mail_exchanger   = %q
    preference       = %d
    view             = %q
    ddns_protected   = %q
}
`, name, mail_exchanger, preference, view, ddnsProtected)
}

func testAccRecordMxDisable(name, mail_exchanger string, preference int64, view, disable string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_mx" "test_disable" {
    name            = %q
    mail_exchanger  = %q
    preference      = %d
    view            = %q
    disable         = %q
}
`, name, mail_exchanger, preference, view, disable)
}

func testAccRecordMxExtAttrs(name, mail_exchanger string, preference int64, view string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf("    %q = %q\n", k, v)
	}
	extattrsStr += "  }"
	return fmt.Sprintf(`
resource "nios_dns_record_mx" "test_extattrs" {
    name            = %q
    mail_exchanger  = %q
    preference      = %d
    view            = %q
    extattrs        = %s
}
`, name, mail_exchanger, preference, view, extattrsStr)
}

func testAccRecordMxForbidReclamation(name, mail_exchanger string, preference int64, view, forbidReclamation string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_mx" "test_forbid_reclamation" {
    name               = %q
    mail_exchanger     = %q
    preference         = %d
    view               = %q
    forbid_reclamation = %q
}
`, name, mail_exchanger, preference, view, forbidReclamation)
}

func testAccRecordMxMailExchanger(name, mail_exchanger string, preference int64, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_mx" "test_mail_exchanger" {
    name            = %q
    mail_exchanger  = %q
    preference      = %d
    view            = %q
}
`, name, mail_exchanger, preference, view)
}

func testAccRecordMxName(name, mail_exchanger string, preference int64, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_mx" "test_name" {
    name            = %q
    mail_exchanger  = %q
    preference      = %d
    view            = %q
}
`, name, mail_exchanger, preference, view)
}

func testAccRecordMxPreference(name, mail_exchanger string, preference int64, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_mx" "test_preference" {
    name            = %q
    mail_exchanger  = %q
    preference      = %d
    view            = %q
}
`, name, mail_exchanger, preference, view)
}

func testAccRecordMxTtl(name, mail_exchanger string, preference int64, view string, ttl int32, use_ttl string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_mx" "test_ttl" {
    name            = %q
    mail_exchanger  = %q
    preference      = %d
    view            = %q
    ttl             = %d
	use_ttl  		= %q
}
`, name, mail_exchanger, preference, view, ttl, use_ttl)
}

func testAccRecordMxUseTtl(name, mail_exchanger string, preference int64, view, useTtl string, ttl int32) string {
	return fmt.Sprintf(`
resource "nios_dns_record_mx" "test_use_ttl" {
    name            = %q
    mail_exchanger  = %q
    preference      = %d
    view            = %q
	use_ttl         = %q
	ttl			    = %d
}
`, name, mail_exchanger, preference, view, useTtl, ttl)
}

func testAccRecordMxView(name, mail_exchanger string, preference int64, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_mx" "test_view" {
    name            = %q
    mail_exchanger  = %q
    preference      = %d
    view            = %q
}
`, name, mail_exchanger, preference, view)
}
