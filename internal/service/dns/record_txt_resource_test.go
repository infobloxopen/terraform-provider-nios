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

var readableAttributesForRecordTxt = "aws_rte53_record_info,cloud_info,comment,creation_time,creator,ddns_principal,ddns_protected,disable,dns_name,extattrs,forbid_reclamation,last_queried,name,reclaimable,shared_record_group,text,ttl,use_ttl,view,zone"

func TestAccRecordTxtResource_basic(t *testing.T) {
	var resourceName = "nios_dns_record_txt.test"
	var v dns.RecordTxt
	name := acctest.RandomNameWithPrefix("record-txt") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTxtBasicConfig(name, "Record Text", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "text", "Record Text"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
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

func TestAccRecordTxtResource_disappears(t *testing.T) {
	resourceName := "nios_dns_record_txt.test"
	var v dns.RecordTxt
	name := acctest.RandomNameWithPrefix("record-txt") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordTxtDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordTxtBasicConfig(name, "Record Text", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					testAccCheckRecordTxtDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordTxtResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_record_txt.test_comment"
	var v dns.RecordTxt
	name := acctest.RandomNameWithPrefix("record-txt") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTxtComment(name, "Record Text", "This is a new record"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a new record"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTxtComment(name, "Record Text", "This is an updated record"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated record"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTxtResource_Creator(t *testing.T) {
	var resourceName = "nios_dns_record_txt.test_creator"
	var v dns.RecordTxt
	name := acctest.RandomNameWithPrefix("record-txt") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTxtCreator(name, "Record Text", "STATIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTxtCreator(name, "Record Text", "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "DYNAMIC"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTxtResource_DdnsPrincipal(t *testing.T) {
	var resourceName = "nios_dns_record_txt.test_ddns_principal"
	var v dns.RecordTxt
	name := acctest.RandomNameWithPrefix("record-txt") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTxtDdnsPrincipal(name, "Record Text", "dhcp/server1@CORP.LOCAL", "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "dhcp/server1@CORP.LOCAL"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTxtDdnsPrincipal(name, "Record Text", "dhcp/server2@CORP.LOCAL", "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "dhcp/server2@CORP.LOCAL"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTxtResource_DdnsProtected(t *testing.T) {
	var resourceName = "nios_dns_record_txt.test_ddns_protected"
	var v dns.RecordTxt
	name := acctest.RandomNameWithPrefix("record-txt") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTxtDdnsProtected(name, "Record Text", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTxtDdnsProtected(name, "Record Text", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTxtResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_record_txt.test_disable"
	var v dns.RecordTxt
	name := acctest.RandomNameWithPrefix("record-txt") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTxtDisable(name, "Record Text", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTxtDisable(name, "Record Text", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTxtResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_record_txt.test_extattrs"
	var v dns.RecordTxt
	name := acctest.RandomNameWithPrefix("record-txt") + ".example.com"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTxtExtAttrs(name, "Record Text", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTxtExtAttrs(name, "Record Text", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTxtResource_ForbidReclamation(t *testing.T) {
	var resourceName = "nios_dns_record_txt.test_forbid_reclamation"
	var v dns.RecordTxt
	name := acctest.RandomNameWithPrefix("record-txt") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTxtForbidReclamation(name, "Record Text", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTxtForbidReclamation(name, "Record Text", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTxtResource_Name(t *testing.T) {
	var resourceName = "nios_dns_record_txt.test_name"
	var v dns.RecordTxt
	name1 := acctest.RandomNameWithPrefix("record-txt") + ".example.com"
	name2 := acctest.RandomNameWithPrefix("record-txt") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTxtName(name1, "Record Text"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTxtName(name2, "Record Text"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTxtResource_Text(t *testing.T) {
	var resourceName = "nios_dns_record_txt.test_text"
	var v dns.RecordTxt
	name := acctest.RandomNameWithPrefix("record-txt") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTxtText(name, "Record Text"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "text", "Record Text"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTxtText(name, "Record Updated Text"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "text", "Record Updated Text"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTxtResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_record_txt.test_ttl"
	var v dns.RecordTxt
	name := acctest.RandomNameWithPrefix("record-txt") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTxtTtl(name, "Record Text", 10, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTxtTtl(name, "Record Text", 1000, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "1000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTxtResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_record_txt.test_use_ttl"
	var v dns.RecordTxt
	name := acctest.RandomNameWithPrefix("record-txt") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTxtUseTtl(name, "Record Text", "true", 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTxtUseTtl(name, "Record Text", "false", 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordTxtExists(ctx context.Context, resourceName string, v *dns.RecordTxt) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			RecordTxtAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordTxt).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordTxtResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordTxtResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordTxtDestroy(ctx context.Context, v *dns.RecordTxt) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			RecordTxtAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordTxt).
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

func testAccCheckRecordTxtDisappears(ctx context.Context, v *dns.RecordTxt) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			RecordTxtAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordTxtBasicConfig(name, text, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_txt" "test" {
	name = %q
	text = %q
	view = %q
}
`, name, text, view)
}

func testAccRecordTxtComment(name, text, comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_txt" "test_comment" {
	name = %q
	text = %q
    comment = %q
}
`, name, text, comment)
}

func testAccRecordTxtCreator(name, text, creator string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_txt" "test_creator" {
	name = %q
	text = %q
    creator = %q
}
`, name, text, creator)
}

func testAccRecordTxtDdnsPrincipal(name, text, ddnsPrincipal, creator string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_txt" "test_ddns_principal" {
	name = %q
	text = %q
    ddns_principal = %q
	creator = %q
}
`, name, text, ddnsPrincipal, creator)
}

func testAccRecordTxtDdnsProtected(name, text, ddnsProtected string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_txt" "test_ddns_protected" {
	name = %q
	text = %q
    ddns_protected = %q
}
`, name, text, ddnsProtected)
}

func testAccRecordTxtDisable(name, text, disable string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_txt" "test_disable" {
	name = %q
	text = %q
    disable = %q
}
`, name, text, disable)
}

func testAccRecordTxtExtAttrs(name, text string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_dns_record_txt" "test_extattrs" {
	name = %q
	text = %q
    extattrs = %s
}
`, name, text, extattrsStr)
}

func testAccRecordTxtForbidReclamation(name, text, forbidReclamation string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_txt" "test_forbid_reclamation" {
	name = %q
	text = %q
    forbid_reclamation = %q
}
`, name, text, forbidReclamation)
}

func testAccRecordTxtName(name, text string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_txt" "test_name" {
	name = %q
	text = %q
}
`, name, text)
}

func testAccRecordTxtText(name, text string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_txt" "test_text" {
	name = %q
	text = %q
}
`, name, text)
}

func testAccRecordTxtTtl(name, text string, ttl int32, useTTL string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_txt" "test_ttl" {
	name = %q
	text = %q
    ttl = %d
	use_ttl = %q
}
`, name, text, ttl, useTTL)
}

func testAccRecordTxtUseTtl(name, text, useTtl string, ttl int) string {
	return fmt.Sprintf(`
resource "nios_dns_record_txt" "test_use_ttl" {
	name = %q
	text = %q
    use_ttl = %q
	ttl = %d
}
`, name, text, useTtl, ttl)
}
