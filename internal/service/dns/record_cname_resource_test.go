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

var readableAttributesForRecordCname = "aws_rte53_record_info,canonical,cloud_info,comment,creation_time,creator,ddns_principal,ddns_protected,disable,dns_canonical,dns_name,extattrs,forbid_reclamation,last_queried,name,reclaimable,shared_record_group,ttl,use_ttl,view,zone"

func TestAccRecordCnameResource_basic(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test"
	var v dns.RecordCname
	canonical := acctest.RandomNameWithPrefix("test-cname") + ".example.com"
	name := acctest.RandomNameWithPrefix("test-cname") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameBasicConfig(canonical, name, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_disappears(t *testing.T) {
	resourceName := "nios_dns_record_cname.test"
	var v dns.RecordCname
	canonical := acctest.RandomNameWithPrefix("test-cname") + ".example.com"
	name := acctest.RandomNameWithPrefix("test-cname") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordCnameDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordCnameBasicConfig(canonical, name, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					testAccCheckRecordCnameDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordCnameResource_Canonical(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_canonical"
	var v dns.RecordCname
	canonical1 := acctest.RandomNameWithPrefix("test-cname") + ".example.com"
	canonical2 := acctest.RandomNameWithPrefix("test-cname") + ".example.com"
	name := acctest.RandomNameWithPrefix("test-cname") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameCanonical(canonical1, name, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", canonical1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameCanonical(canonical2, name, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", canonical2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_comment"
	var v dns.RecordCname
	canonical := acctest.RandomNameWithPrefix("test-cname") + ".example.com"
	name := acctest.RandomNameWithPrefix("test-cname") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameComment(canonical, name, "default", "This is a new record"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a new record"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameComment(canonical, name, "default", "This is an updated record"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated record"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_Creator(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_creator"
	var v dns.RecordCname
	canonical := acctest.RandomNameWithPrefix("test-cname") + ".example.com"
	name := acctest.RandomNameWithPrefix("test-cname") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameCreator(canonical, name, "default", "STATIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameCreator(canonical, name, "default", "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "DYNAMIC"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_DdnsPrincipal(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_ddns_principal"
	var v dns.RecordCname
	canonical := acctest.RandomNameWithPrefix("test-cname") + ".example.com"
	name := acctest.RandomNameWithPrefix("test-cname") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameDdnsPrincipal(canonical, name, "default", "DDNS_PRINCIPAL_REPLACE_ME", "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "DDNS_PRINCIPAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameDdnsPrincipal(canonical, name, "default", "DDNS_PRINCIPAL_UPDATE_REPLACE_ME", "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "DDNS_PRINCIPAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_DdnsProtected(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_ddns_protected"
	var v dns.RecordCname
	canonical := acctest.RandomNameWithPrefix("test-cname") + ".example.com"
	name := acctest.RandomNameWithPrefix("test-cname") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameDdnsProtected(canonical, name, "default", "false"),
				Check: resource.ComposeTestCheckFunc(
					//testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameDdnsProtected(canonical, name, "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_disable"
	var v dns.RecordCname
	canonical := acctest.RandomNameWithPrefix("test-cname") + ".example.com"
	name := acctest.RandomNameWithPrefix("test-cname") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameDisable(canonical, name, "default", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameDisable(canonical, name, "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_Extattrs(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_extattrs"
	var v dns.RecordCname
	canonical := acctest.RandomNameWithPrefix("test-cname") + ".example.com"
	name := acctest.RandomNameWithPrefix("test-cname") + ".example.com"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameExtAttrs(canonical, name, "default", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameExtAttrs(canonical, name, "default", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_ForbidReclamation(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_forbid_reclamation"
	var v dns.RecordCname
	canonical := acctest.RandomNameWithPrefix("test-cname") + ".example.com"
	name := acctest.RandomNameWithPrefix("test-cname") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameForbidReclamation(canonical, name, "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameForbidReclamation(canonical, name, "default", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_Name(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_name"
	var v dns.RecordCname
	name1 := acctest.RandomName() + ".example.com"
	name2 := acctest.RandomName() + ".example.com"
	canonical := acctest.RandomNameWithPrefix("test-cname") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameName(canonical, name1, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameName(canonical, name2, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_ttl"
	var v dns.RecordCname
	canonical := acctest.RandomNameWithPrefix("test-cname") + ".example.com"
	name := acctest.RandomNameWithPrefix("test-cname") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameTtl(canonical, name, "default", 10, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameTtl(canonical, name, "default", 0, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_use_ttl"
	var v dns.RecordCname
	canonical := acctest.RandomNameWithPrefix("test-cname") + ".example.com"
	name := acctest.RandomNameWithPrefix("test-cname") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameUseTtl(canonical, name, "default", "true", 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameUseTtl(canonical, name, "default", "false", 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_View(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_view"
	var v dns.RecordCname
	canonical := acctest.RandomNameWithPrefix("test-cname") + ".example.com"
	name := acctest.RandomNameWithPrefix("test-cname") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameView(canonical, name, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameView(canonical, name, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordCnameExists(ctx context.Context, resourceName string, v *dns.RecordCname) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			RecordCnameAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordCname).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordCnameResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordCnameResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordCnameDestroy(ctx context.Context, v *dns.RecordCname) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			RecordCnameAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordCname).
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

func testAccCheckRecordCnameDisappears(ctx context.Context, v *dns.RecordCname) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			RecordCnameAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordCnameBasicConfig(name, ipV4Addr, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test" {
	canonical = %q
	name = %q
	view = %q
}
`, name, ipV4Addr, view)
}

func testAccRecordCnameCanonical(canonical, name, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_canonical" {
	canonical = %q
	name      = %q
	view      = %q
}
`, canonical, name, view)
}

func testAccRecordCnameComment(canonical, name, view, comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_comment" {
	canonical = %q
	name      = %q
	view      = %q
	comment   = %q
}
`, canonical, name, view, comment)
}

func testAccRecordCnameCreator(canonical, name, view, creator string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_creator" {
	canonical = %q
	name      = %q
	view      = %q
	creator   = %q
}
`, canonical, name, view, creator)
}

func testAccRecordCnameDdnsPrincipal(canonical, name, view, ddnsPrincipal, creator string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_ddns_principal" {
	canonical      = %q
	name           = %q
	view           = %q
	ddns_principal = %q
	creator		   = %q
}
`, canonical, name, view, ddnsPrincipal, creator)
}

func testAccRecordCnameDdnsProtected(canonical, name, view, ddnsProtected string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_ddns_protected" {
	canonical       = %q
	name            = %q
	view            = %q
	ddns_protected  = %q
}
`, canonical, name, view, ddnsProtected)
}

func testAccRecordCnameDisable(canonical, name, view, disable string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_disable" {
	canonical = %q
	name      = %q
	view      = %q
	disable   = %q
}
`, canonical, name, view, disable)
}

func testAccRecordCnameExtAttrs(canonical, name, view string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_extattrs" {
	canonical = %q
	name      = %q
	view      = %q
	extattrs  = %s
}
`, canonical, name, view, extattrsStr)
}

func testAccRecordCnameForbidReclamation(canonical, name, view, forbidReclamation string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_forbid_reclamation" {
	canonical          = %q
	name               = %q
	view               = %q
	forbid_reclamation = %q
}
`, canonical, name, view, forbidReclamation)
}

func testAccRecordCnameName(canonical, name, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_name" {
	canonical = %q
	name      = %q
	view      = %q
}
`, canonical, name, view)
}

func testAccRecordCnameTtl(canonical, name, view string, ttl int32, useTTL string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_ttl" {
	canonical = %q
	name      = %q
	view      = %q
	ttl       = %d
	use_ttl   = %q

}
`, canonical, name, view, ttl, useTTL)
}

func testAccRecordCnameUseTtl(canonical, name, view, useTtl string, ttl int32) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_use_ttl" {
	canonical = %q
	name      = %q
	view      = %q
	use_ttl   = %q
	ttl 	  = %d
}
`, canonical, name, view, useTtl, ttl)
}

func testAccRecordCnameView(canonical, name, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_view" {
	canonical = %q
	name      = %q
	view      = %q
}
`, canonical, name, view)
}
