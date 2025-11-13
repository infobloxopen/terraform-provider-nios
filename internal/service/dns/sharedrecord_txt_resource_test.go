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

var readableAttributesForSharedrecordTxt = "comment,disable,dns_name,extattrs,name,shared_record_group,text,ttl,use_ttl"

func TestAccSharedrecordTxtResource_basic(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_txt.test"
	var v dns.SharedrecordTxt
	name := acctest.RandomNameWithPrefix("sharedrecord-txt-")
	text := "This is a shared record TXT record"
	sharedRecordGroup := "sr1"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordTxtBasicConfig(name, sharedRecordGroup, text),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordTxtExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordTxtResource_disappears(t *testing.T) {
	resourceName := "nios_dns_sharedrecord_txt.test"
	var v dns.SharedrecordTxt
	name := acctest.RandomNameWithPrefix("sharedrecord-txt-")
	text := "This is a shared record TXT record"
	sharedRecordGroup := "sr1"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharedrecordTxtDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharedrecordTxtBasicConfig(name, sharedRecordGroup, text),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordTxtExists(context.Background(), resourceName, &v),
					testAccCheckSharedrecordTxtDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSharedrecordTxtResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_txt.test_comment"
	var v dns.SharedrecordTxt
	name := acctest.RandomNameWithPrefix("sharedrecord-txt-")
	text := "This is a shared record TXT record"
	sharedRecordGroup := "sr1"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordTxtComment(name, sharedRecordGroup, text, "Shared TXT Record Comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Shared TXT Record Comment"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordTxtComment(name, sharedRecordGroup, text, "Shared TXT Record Comment Updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Shared TXT Record Comment Updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordTxtResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_txt.test_disable"
	var v dns.SharedrecordTxt
	name := acctest.RandomNameWithPrefix("sharedrecord-txt-")
	text := "This is a shared record TXT record"
	sharedRecordGroup := "sr1"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordTxtDisable(name, sharedRecordGroup, text, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordTxtDisable(name, sharedRecordGroup, text, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordTxtResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_txt.test_extattrs"
	var v dns.SharedrecordTxt
	name := acctest.RandomNameWithPrefix("sharedrecord-txt-")
	text := "This is a shared record TXT record"
	sharedRecordGroup := "sr1"
	extAttrs1 := acctest.RandomName()
	extAttrs2 := acctest.RandomName()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordTxtExtAttrs(name, sharedRecordGroup, text, map[string]any{"Site": extAttrs1}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrs1),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordTxtExtAttrs(name, sharedRecordGroup, text, map[string]any{"Site": extAttrs2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrs2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordTxtResource_Name(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_txt.test_name"
	var v dns.SharedrecordTxt
	name1 := acctest.RandomNameWithPrefix("sharedrecord-txt-")
	name2 := acctest.RandomNameWithPrefix("sharedrecord-txt-")
	text := "This is a shared record TXT record"
	sharedRecordGroup := "sr1"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordTxtName(name1, sharedRecordGroup, text),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordTxtName(name2, sharedRecordGroup, text),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordTxtResource_Text(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_txt.test_text"
	var v dns.SharedrecordTxt
	name := acctest.RandomNameWithPrefix("sharedrecord-txt-")
	text1 := acctest.RandomName()
	text2 := acctest.RandomName()
	sharedRecordGroup := "sr1"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordTxtText(name, sharedRecordGroup, text1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "text", text1),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordTxtText(name, sharedRecordGroup, text2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "text", text2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordTxtResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_txt.test_ttl"
	var v dns.SharedrecordTxt
	name := acctest.RandomNameWithPrefix("sharedrecord-txt-")
	text := "This is a shared record TXT record"
	sharedRecordGroup := "sr1"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordTxtTtl(name, sharedRecordGroup, text, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordTxtTtl(name, sharedRecordGroup, text, 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordTxtResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_txt.test_use_ttl"
	var v dns.SharedrecordTxt
	name := acctest.RandomNameWithPrefix("sharedrecord-txt-")
	text := "This is a shared record TXT record"
	sharedRecordGroup := "sr1"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordTxtUseTtl(name, sharedRecordGroup, text, 300, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordTxtUseTtl(name, sharedRecordGroup, text, 300, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordTxtExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckSharedrecordTxtExists(ctx context.Context, resourceName string, v *dns.SharedrecordTxt) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			SharedrecordTxtAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForSharedrecordTxt).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetSharedrecordTxtResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetSharedrecordTxtResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckSharedrecordTxtDestroy(ctx context.Context, v *dns.SharedrecordTxt) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			SharedrecordTxtAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForSharedrecordTxt).
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

func testAccCheckSharedrecordTxtDisappears(ctx context.Context, v *dns.SharedrecordTxt) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			SharedrecordTxtAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccSharedrecordTxtBasicConfig(name, sharedRecordGroup, text string) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecord_txt" "test" {
	name = %q
	shared_record_group = %q
	text = %q
}
`, name, sharedRecordGroup, text)
}

func testAccSharedrecordTxtComment(name, sharedRecordGroup, text, comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecord_txt" "test_comment" {
    name = %q
	shared_record_group = %q
	text = %q
    comment = %q
}
`, name, sharedRecordGroup, text, comment)
}

func testAccSharedrecordTxtDisable(name, sharedRecordGroup, text string, disable bool) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecord_txt" "test_disable" {
    name = %q
	shared_record_group = %q
	text = %q
    disable = %t
}
`, name, sharedRecordGroup, text, disable)
}

func testAccSharedrecordTxtExtAttrs(name, sharedRecordGroup, text string, extAttrs map[string]any) string {
	extAttrsStr := utils.ConvertMapToHCL(extAttrs)
	return fmt.Sprintf(`
resource "nios_dns_sharedrecord_txt" "test_extattrs" {
    name = %q
	shared_record_group = %q
	text = %q
    extattrs = %s
}
`, name, sharedRecordGroup, text, extAttrsStr)
}

func testAccSharedrecordTxtName(name, sharedRecordGroup, text string) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecord_txt" "test_name" {
    name = %q
	shared_record_group = %q
	text = %q
}
`, name, sharedRecordGroup, text)
}

func testAccSharedrecordTxtText(name, sharedRecordGroup, text string) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecord_txt" "test_text" {
    name = %q
	shared_record_group = %q
	text = %q
}
`, name, sharedRecordGroup, text)
}

func testAccSharedrecordTxtTtl(name, sharedRecordGroup, text string, ttl int32) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecord_txt" "test_ttl" {
    name = %q
	shared_record_group = %q
	text = %q
    ttl = %d
    use_ttl = true
}
`, name, sharedRecordGroup, text, ttl)
}

func testAccSharedrecordTxtUseTtl(name, sharedRecordGroup, text string, ttl int32, useTtl bool) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecord_txt" "test_use_ttl" {
    name = %q
	shared_record_group = %q
	text = %q
	ttl = %d
    use_ttl = %t
}
`, name, sharedRecordGroup, text, ttl, useTtl)
}
