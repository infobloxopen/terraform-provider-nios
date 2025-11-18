package dns_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// TODO: OBJECTS TO BE PRESENT IN THE GRID FOR TESTS
// sharedrecordgroup: "example-sharedrecordgroup"

var readableAttributesForSharedrecordMx = "comment,disable,dns_mail_exchanger,dns_name,extattrs,mail_exchanger,name,preference,shared_record_group,ttl,use_ttl"

func TestAccSharedrecordMxResource_basic(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_mx.test"
	var v dns.SharedrecordMx
	name := acctest.RandomNameWithPrefix("sharedrecord-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordMxBasicConfig(mail_exchanger, name, 10, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mail_exchanger", mail_exchanger),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "preference", "10"),
					resource.TestCheckResourceAttr(resourceName, "shared_record_group", sharedRecordGroup),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordMxResource_disappears(t *testing.T) {
	resourceName := "nios_dns_sharedrecord_mx.test"
	var v dns.SharedrecordMx
	name := acctest.RandomNameWithPrefix("sharedrecord-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharedrecordMxDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharedrecordMxBasicConfig(mail_exchanger, name, 10, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					testAccCheckSharedrecordMxDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSharedrecordMxResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_mx.test_comment"
	var v dns.SharedrecordMx
	name := acctest.RandomNameWithPrefix("sharedrecord-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordMxComment(mail_exchanger, name, 10, sharedRecordGroup, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordMxComment(mail_exchanger, name, 10, sharedRecordGroup, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordMxResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_mx.test_disable"
	var v dns.SharedrecordMx
	name := acctest.RandomNameWithPrefix("sharedrecord-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordMxDisable(mail_exchanger, name, 10, sharedRecordGroup, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordMxDisable(mail_exchanger, name, 10, sharedRecordGroup, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordMxResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_mx.test_extattrs"
	var v dns.SharedrecordMx
	name := acctest.RandomNameWithPrefix("sharedrecord-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordMxExtAttrs(mail_exchanger, name, 10, sharedRecordGroup, map[string]any{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordMxExtAttrs(mail_exchanger, name, 10, sharedRecordGroup, map[string]any{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordMxResource_MailExchanger(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_mx.test_mail_exchanger"
	var v dns.SharedrecordMx
	name := acctest.RandomNameWithPrefix("sharedrecord-mx") + ".example.com"
	mail_exchanger1 := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
	mail_exchanger2 := acctest.RandomNameWithPrefix("updatedmail-exchanger") + ".example.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordMxMailExchanger(mail_exchanger1, name, 10, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mail_exchanger", mail_exchanger1),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordMxMailExchanger(mail_exchanger2, "example.com", 10, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mail_exchanger", mail_exchanger2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordMxResource_Name(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_mx.test_name"
	var v dns.SharedrecordMx
	name1 := acctest.RandomNameWithPrefix("sharedrecord-mx") + ".example.com"
	name2 := acctest.RandomNameWithPrefix("updatedsharedrecord-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordMxName(mail_exchanger, name1, 10, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordMxName(mail_exchanger, name2, 10, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordMxResource_Preference(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_mx.test_preference"
	var v dns.SharedrecordMx
	name := acctest.RandomNameWithPrefix("sharedrecord-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordMxPreference(mail_exchanger, name, 10, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preference", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordMxPreference(mail_exchanger, name, 20, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preference", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordMxResource_SharedRecordGroup(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_mx.test_shared_record_group"
	var v dns.SharedrecordMx
	name := acctest.RandomNameWithPrefix("sharedrecord-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordMxSharedRecordGroup(mail_exchanger, name, 10, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "shared_record_group", sharedRecordGroup),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordMxResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_mx.test_ttl"
	var v dns.SharedrecordMx
	name := acctest.RandomNameWithPrefix("sharedrecord-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordMxTtl(mail_exchanger, name, 10, sharedRecordGroup, 3600, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "3600"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordMxTtl(mail_exchanger, name, 10, sharedRecordGroup, 4200, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "4200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordMxResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_mx.test_use_ttl"
	var v dns.SharedrecordMx
	name := acctest.RandomNameWithPrefix("sharedrecord-mx") + ".example.com"
	mail_exchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordMxUseTtl(mail_exchanger, name, 10, sharedRecordGroup, "true", 3600),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordMxUseTtl(mail_exchanger, name, 10, sharedRecordGroup, "false", 4200),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordMxExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckSharedrecordMxExists(ctx context.Context, resourceName string, v *dns.SharedrecordMx) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			SharedrecordMxAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForSharedrecordMx).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetSharedrecordMxResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetSharedrecordMxResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckSharedrecordMxDestroy(ctx context.Context, v *dns.SharedrecordMx) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			SharedrecordMxAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForSharedrecordMx).
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

func testAccCheckSharedrecordMxDisappears(ctx context.Context, v *dns.SharedrecordMx) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			SharedrecordMxAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccSharedrecordMxBasicConfig(mailExchanger, name string, preference int, sharedRecordGroup string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_mx" "test" {
    mail_exchanger      = %q
    name                = %q
    preference          = %d
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
}
`, mailExchanger, name, preference)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordMxComment(mailExchanger, name string, preference int, sharedRecordGroup, comment string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_mx" "test_comment" {
    mail_exchanger      = %q
    name                = %q
    preference          = %d
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
    comment             = %q
}
`, mailExchanger, name, preference, comment)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordMxDisable(mailExchanger, name string, preference int, sharedRecordGroup, disable string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_mx" "test_disable" {
    mail_exchanger      = %q
    name                = %q
    preference          = %d
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
    disable             = %s
}
`, mailExchanger, name, preference, disable)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordMxExtAttrs(mailExchanger, name string, preference int, sharedRecordGroup string, extAttrs map[string]any) string {
	extAttrsStr := utils.ConvertMapToHCL(extAttrs)
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_mx" "test_extattrs" {
    mail_exchanger      = %q
    name                = %q
    preference          = %d
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
    extattrs            = %s
}
`, mailExchanger, name, preference, extAttrsStr)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordMxMailExchanger(mailExchanger, name string, preference int, sharedRecordGroup string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_mx" "test_mail_exchanger" {
    mail_exchanger      = %q
    name                = %q
    preference          = %d
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
}
`, mailExchanger, name, preference)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordMxName(mailExchanger, name string, preference int, sharedRecordGroup string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_mx" "test_name" {
    mail_exchanger      = %q
    name                = %q
    preference          = %d
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
}
`, mailExchanger, name, preference)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordMxPreference(mailExchanger, name string, preference int, sharedRecordGroup string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_mx" "test_preference" {
    mail_exchanger      = %q
    name                = %q
    preference          = %d
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
}
`, mailExchanger, name, preference)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordMxSharedRecordGroup(mailExchanger, name string, preference int, sharedRecordGroup string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_mx" "test_shared_record_group" {
    mail_exchanger      = %q
    name                = %q
    preference          = %d
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
}
`, mailExchanger, name, preference)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordMxTtl(mailExchanger, name string, preference int, sharedRecordGroup string, ttl int, useTtl string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_mx" "test_ttl" {
    mail_exchanger      = %q
    name                = %q
    preference          = %d
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
    ttl                 = %d
    use_ttl             = %s
}
`, mailExchanger, name, preference, ttl, useTtl)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordMxUseTtl(mailExchanger, name string, preference int, sharedRecordGroup, useTtl string, ttl int) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_mx" "test_use_ttl" {
    mail_exchanger      = %q
    name                = %q
    preference          = %d
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
    use_ttl             = %s
    ttl                 = %d
}
`, mailExchanger, name, preference, useTtl, ttl)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccBaseSharedRecordGroup(name string) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecordgroup" "parent_sharedrecord_group" {
  name = %q
}
`, name)
}
