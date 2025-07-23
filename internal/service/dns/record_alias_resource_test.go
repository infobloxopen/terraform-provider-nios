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
// -> Parent Zone: example.com (in default view)

var readableAttributesForRecordAlias = "aws_rte53_record_info,cloud_info,comment,creator,disable,dns_name,dns_target_name,extattrs,last_queried,name,target_name,target_type,ttl,use_ttl,view,zone"

func TestAccRecordAliasResource_basic(t *testing.T) {
	var resourceName = "nios_dns_record_alias.test"
	var v dns.RecordAlias

	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAliasBasicConfig(name, "server.example.com", "A", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "target_name", "server.example.com"),
					resource.TestCheckResourceAttr(resourceName, "target_type", "A"),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAliasResource_disappears(t *testing.T) {
	resourceName := "nios_dns_record_alias.test"
	var v dns.RecordAlias
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordAliasDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordAliasBasicConfig(name, "server.example.com", "A", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					testAccCheckRecordAliasDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordAliasResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_record_alias.test_comment"
	var v dns.RecordAlias
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAliasComment(name, "server.example.com", "A", "default", "This is a sample comment."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a sample comment."),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAliasComment(name, "server.example.com", "A", "default", "This is an updated comment."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAliasResource_Creator(t *testing.T) {
	var resourceName = "nios_dns_record_alias.test_creator"
	var v dns.RecordAlias
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAliasCreator(name, "server.example.com", "A", "default", "STATIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAliasCreator(name, "server.example.com", "A", "default", "STATIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAliasResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_record_alias.test_disable"
	var v dns.RecordAlias
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAliasDisable(name, "server.example.com", "A", "default", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAliasDisable(name, "server.example.com", "A", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAliasResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_record_alias.test_extattrs"
	var v dns.RecordAlias
	name := acctest.RandomName() + ".example.com"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAliasExtAttrs(name, "server.example.com", "A", "default", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAliasExtAttrs(name, "server.example.com", "A", "default", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAliasResource_Name(t *testing.T) {
	var resourceName = "nios_dns_record_alias.test_name"
	var v dns.RecordAlias
	name1 := acctest.RandomName() + ".example.com"
	name2 := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAliasName(name1, "server.example.com", "A", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAliasName(name2, "server.example.com", "A", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAliasResource_TargetName(t *testing.T) {
	var resourceName = "nios_dns_record_alias.test_target_name"
	var v dns.RecordAlias
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAliasTargetName(name, "server.example.com", "A", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "target_name", "server.example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAliasTargetName(name, "updated-server.example.com", "A", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "target_name", "updated-server.example.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAliasResource_TargetType(t *testing.T) {
	var resourceName = "nios_dns_record_alias.test_target_type"
	var v dns.RecordAlias
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAliasTargetType(name, "server.example.com", "A", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "target_type", "A"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAliasTargetType(name, "server.example.com", "AAAA", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "target_type", "AAAA"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAliasResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_record_alias.test_ttl"
	var v dns.RecordAlias
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAliasTtl(name, "server.example.com", "A", "default", 10, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAliasTtl(name, "server.example.com", "A", "default", 0, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAliasResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_record_alias.test_use_ttl"
	var v dns.RecordAlias
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAliasUseTtl(name, "server.example.com", "A", "default", "true", 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAliasUseTtl(name, "server.example.com", "A", "default", "false", 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAliasResource_View(t *testing.T) {
	var resourceName = "nios_dns_record_alias.test_view"
	var v dns.RecordAlias
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAliasView(name, "server.example.com", "A", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAliasView(name, "server.example.com", "A", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordAliasExists(ctx context.Context, resourceName string, v *dns.RecordAlias) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		state.RootModule().Resources[resourceName].Primary.ID = utils.ExtractResourceRef(rs.Primary.Attributes["ref"])
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			RecordAliasAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordAlias).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordAliasResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordAliasResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordAliasDestroy(ctx context.Context, v *dns.RecordAlias) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			RecordAliasAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordAlias).
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

func testAccCheckRecordAliasDisappears(ctx context.Context, v *dns.RecordAlias) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			RecordAliasAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordAliasBasicConfig(name, target_name, target_type, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_alias" "test" {
	name 		= %q
	target_name = %q
	target_type = %q
	view 		= %q
}
`, name, target_name, target_type, view)
}

func testAccRecordAliasComment(name, target_name, target_type, view, comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_alias" "test_comment" {
	name		= %q
	target_name = %q
	target_type = %q
	view 		= %q
    comment 	= %q
}
`, name, target_name, target_type, view, comment)
}

func testAccRecordAliasCreator(name, target_name, target_type, view, creator string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_alias" "test_creator" {
	name 		= %q
	target_name = %q
	target_type = %q
	view 		= %q
    creator 	= %q
}
`, name, target_name, target_type, view, creator)
}

func testAccRecordAliasDisable(name, target_name, target_type, view, disable string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_alias" "test_disable" {
	name 		= %q
	target_name = %q
	target_type = %q
	view 		= %q
    disable 	= %q
}
`, name, target_name, target_type, view, disable)
}

func testAccRecordAliasExtAttrs(name, target_name, target_type, view string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_dns_record_alias" "test_extattrs" {
	name 		= %q
	target_name = %q
	target_type = %q
	view 		= %q
    extattrs 	= %s
}
`, name, target_name, target_type, view, extattrsStr)
}

func testAccRecordAliasName(name, target_name, target_type, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_alias" "test_name" {
    name 		= %q
	target_name = %q
	target_type = %q
	view 		= %q
}
`, name, target_name, target_type, view)
}

func testAccRecordAliasTargetName(name, target_name, target_type, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_alias" "test_target_name" {
    name 		= %q
	target_name = %q
	target_type = %q
	view 		= %q
}
`, name, target_name, target_type, view)
}

func testAccRecordAliasTargetType(name, target_name, target_type, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_alias" "test_target_type" {
    name 		= %q
	target_name = %q
	target_type = %q
	view 		= %q
}
`, name, target_name, target_type, view)
}

func testAccRecordAliasTtl(name, target_name, target_type, view string, ttl int32, use_ttl string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_alias" "test_ttl" {
	name 		= %q
	target_name = %q
	target_type = %q
	view 		= %q
    ttl 		= %d
	use_ttl 	= %q
}
`, name, target_name, target_type, view, ttl, use_ttl)
}

func testAccRecordAliasUseTtl(name, target_name, target_type, view, useTtl string, ttl int32) string {
	return fmt.Sprintf(`
resource "nios_dns_record_alias" "test_use_ttl" {
	name 		= %q
	target_name = %q
	target_type = %q
	view 		= %q
    use_ttl 	= %q
	ttl 		= %d
}
`, name, target_name, target_type, view, useTtl, ttl)
}

func testAccRecordAliasView(name, target_name, target_type, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_alias" "test_view" {
    name 		= %q
	target_name = %q
	target_type = %q
	view 		= %q
}
`, name, target_name, target_type, view)
}
