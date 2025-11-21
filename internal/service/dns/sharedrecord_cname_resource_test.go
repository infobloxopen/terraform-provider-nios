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

var readableAttributesForSharedrecordCname = "canonical,comment,disable,dns_canonical,dns_name,extattrs,name,shared_record_group,ttl,use_ttl"

func TestAccSharedrecordCnameResource_basic(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_cname.test"
	var v dns.SharedrecordCname
	name := acctest.RandomNameWithPrefix("sharedrecord-cname-")
	canonical := acctest.RandomNameWithPrefix("canonical-name") + ".com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordCnameBasicConfig(name, canonical, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "canonical", canonical),
					resource.TestCheckResourceAttr(resourceName, "shared_record_group", sharedRecordGroup),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordCnameResource_disappears(t *testing.T) {
	resourceName := "nios_dns_sharedrecord_cname.test"
	var v dns.SharedrecordCname
	name := acctest.RandomNameWithPrefix("sharedrecord-cname-")
	canonical := acctest.RandomNameWithPrefix("canonical-name") + ".com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharedrecordCnameDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharedrecordCnameBasicConfig(name, canonical, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					testAccCheckSharedrecordCnameDisappears(context.Background(), &v),
				),
			},
		},
	})
}

func TestAccSharedrecordCnameResource_Canonical(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_cname.test_canonical"
	var v dns.SharedrecordCname
	name := acctest.RandomNameWithPrefix("sharedrecord-cname-")
	canonical1 := acctest.RandomNameWithPrefix("canonical-name") + ".com"
	canonical2 := acctest.RandomNameWithPrefix("canonical-name") + ".com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordCnameCanonical(name, canonical1, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", canonical1),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordCnameCanonical(name, canonical2, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", canonical2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordCnameResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_cname.test_comment"
	var v dns.SharedrecordCname
	name := acctest.RandomNameWithPrefix("sharedrecord-cname-")
	canonical := acctest.RandomNameWithPrefix("canonical-name") + ".com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordCnameComment(name, canonical, sharedRecordGroup, "Example Shared CNAME Record Comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Example Shared CNAME Record Comment"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordCnameComment(name, canonical, sharedRecordGroup, "Example Shared CNAME Record Comment Updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Example Shared CNAME Record Comment Updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordCnameResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_cname.test_disable"
	var v dns.SharedrecordCname
	name := acctest.RandomNameWithPrefix("sharedrecord-cname-")
	canonical := acctest.RandomNameWithPrefix("canonical-name") + ".com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordCnameDisable(name, canonical, sharedRecordGroup, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordCnameDisable(name, canonical, sharedRecordGroup, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordCnameResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_cname.test_extattrs"
	var v dns.SharedrecordCname
	name := acctest.RandomNameWithPrefix("sharedrecord-cname-")
	canonical := acctest.RandomNameWithPrefix("canonical-name") + ".com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")
	extAttrs1 := acctest.RandomName()
	extAttrs2 := acctest.RandomName()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordCnameExtAttrs(name, canonical, sharedRecordGroup, map[string]any{
					"Site": extAttrs1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrs1),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordCnameExtAttrs(name, canonical, sharedRecordGroup, map[string]any{
					"Site": extAttrs2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrs2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordCnameResource_Name(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_cname.test_name"
	var v dns.SharedrecordCname
	name1 := acctest.RandomNameWithPrefix("sharedrecord-cname-")
	name2 := acctest.RandomNameWithPrefix("sharedrecord-cname-")
	canonical := acctest.RandomNameWithPrefix("canonical-name") + ".com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordCnameName(name1, canonical, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordCnameName(name2, canonical, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordCnameResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_cname.test_ttl"
	var v dns.SharedrecordCname
	name := acctest.RandomNameWithPrefix("sharedrecord-cname-")
	canonical := acctest.RandomNameWithPrefix("canonical-name") + ".com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordCnameTtl(name, canonical, sharedRecordGroup, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordCnameTtl(name, canonical, sharedRecordGroup, 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordCnameResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_cname.test_use_ttl"
	var v dns.SharedrecordCname
	name := acctest.RandomNameWithPrefix("sharedrecord-cname-")
	canonical := acctest.RandomNameWithPrefix("canonical-name") + ".com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordCnameUseTtl(name, canonical, sharedRecordGroup, 10, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordCnameUseTtl(name, canonical, sharedRecordGroup, 10, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
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

func testAccSharedrecordCnameBasicConfig(name, canonical, sharedRecordGroup string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_cname" "test" {
	name = %q
	shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
	canonical = %q
}
`, name, canonical)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordCnameCanonical(name, canonical, sharedRecordGroup string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_cname" "test_canonical" {
	name = %q
	shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
	canonical = %q
}
`, name, canonical)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordCnameComment(name, canonical, sharedRecordGroup, comment string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_cname" "test_comment" {
	name = %q
	shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
	canonical = %q
	comment = %q
}
`, name, canonical, comment)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordCnameDisable(name, canonical, sharedRecordGroup string, disable bool) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_cname" "test_disable" {
	name = %q
	shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
	canonical = %q
	disable = %t
}
`, name, canonical, disable)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordCnameExtAttrs(name, canonical, sharedRecordGroup string, extAttrs map[string]any) string {
	extattrsStr := utils.ConvertMapToHCL(extAttrs)
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_cname" "test_extattrs" {
	name = %q
	shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
	canonical = %q
	extattrs = %s
}
`, name, canonical, extattrsStr)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordCnameName(name, canonical, sharedRecordGroup string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_cname" "test_name" {
	name = %q
	shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
	canonical = %q
}
`, name, canonical)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordCnameTtl(name, canonical, sharedRecordGroup string, ttl int32) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_cname" "test_ttl" {
	name = %q
	shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
	canonical = %q
	ttl = %d
    use_ttl = true
}
`, name, canonical, ttl)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordCnameUseTtl(name, canonical, sharedRecordGroup string, ttl int32, useTtl bool) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_cname" "test_use_ttl" {
	name = %q
	shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
	canonical = %q
    ttl = %d
	use_ttl = %t
}
`, name, canonical, ttl, useTtl)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}
