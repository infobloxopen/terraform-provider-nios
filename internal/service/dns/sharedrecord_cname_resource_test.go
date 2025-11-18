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

var readableAttributesForSharedrecordCname = "canonical,comment,disable,dns_canonical,dns_name,extattrs,name,shared_record_group,ttl,use_ttl"

func TestAccSharedrecordCnameResource_basic(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_cname.test"
	var v dns.SharedrecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordCnameBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordCnameResource_disappears(t *testing.T) {
	resourceName := "nios_dns_sharedrecord_cname.test"
	var v dns.SharedrecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharedrecordCnameDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharedrecordCnameBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					testAccCheckSharedrecordCnameDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSharedrecordCnameResource_Ref(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_cname.test_ref"
	var v dns.SharedrecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordCnameRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordCnameRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordCnameResource_Canonical(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_cname.test_canonical"
	var v dns.SharedrecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordCnameCanonical("CANONICAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", "CANONICAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordCnameCanonical("CANONICAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", "CANONICAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordCnameResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_cname.test_comment"
	var v dns.SharedrecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordCnameComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordCnameComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordCnameResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_cname.test_disable"
	var v dns.SharedrecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordCnameDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordCnameDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordCnameResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_cname.test_extattrs"
	var v dns.SharedrecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordCnameExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordCnameExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordCnameResource_Name(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_cname.test_name"
	var v dns.SharedrecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordCnameName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordCnameName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordCnameResource_SharedRecordGroup(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_cname.test_shared_record_group"
	var v dns.SharedrecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordCnameSharedRecordGroup("SHARED_RECORD_GROUP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "shared_record_group", "SHARED_RECORD_GROUP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordCnameSharedRecordGroup("SHARED_RECORD_GROUP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "shared_record_group", "SHARED_RECORD_GROUP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordCnameResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_cname.test_ttl"
	var v dns.SharedrecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordCnameTtl("TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordCnameTtl("TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordCnameResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_cname.test_use_ttl"
	var v dns.SharedrecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordCnameUseTtl("USE_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "USE_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordCnameUseTtl("USE_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "USE_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckSharedrecordCnameExists(ctx context.Context, resourceName string, v *dns.SharedrecordCname) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			SharedrecordCnameAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForSharedrecordCname).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetSharedrecordCnameResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetSharedrecordCnameResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckSharedrecordCnameDestroy(ctx context.Context, v *dns.SharedrecordCname) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			SharedrecordCnameAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForSharedrecordCname).
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

func testAccCheckSharedrecordCnameDisappears(ctx context.Context, v *dns.SharedrecordCname) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			SharedrecordCnameAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccSharedrecordCnameBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dns_sharedrecord_cname" "test" {
}
`)
}

func testAccSharedrecordCnameRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecord_cname" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccSharedrecordCnameCanonical(canonical string) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecord_cname" "test_canonical" {
    canonical = %q
}
`, canonical)
}

func testAccSharedrecordCnameComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecord_cname" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccSharedrecordCnameDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecord_cname" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccSharedrecordCnameExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecord_cname" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccSharedrecordCnameName(name string) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecord_cname" "test_name" {
    name = %q
}
`, name)
}

func testAccSharedrecordCnameSharedRecordGroup(sharedRecordGroup string) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecord_cname" "test_shared_record_group" {
    shared_record_group = %q
}
`, sharedRecordGroup)
}

func testAccSharedrecordCnameTtl(ttl string) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecord_cname" "test_ttl" {
    ttl = %q
}
`, ttl)
}

func testAccSharedrecordCnameUseTtl(useTtl string) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecord_cname" "test_use_ttl" {
    use_ttl = %q
}
`, useTtl)
}
