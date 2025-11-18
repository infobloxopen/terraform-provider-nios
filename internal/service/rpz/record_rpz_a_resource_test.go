package rpz_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/rpz"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForRecordRpzA = "comment,disable,extattrs,ipv4addr,name,rp_zone,ttl,use_ttl,view,zone"

func TestAccRecordRpzAResource_basic(t *testing.T) {
	var resourceName = "nios_rpz_record_rpz_a.test"
	var v rpz.RecordRpzA
	name := acctest.RandomName() + ".rpz.example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzABasicConfig(name, "10.10.0.1", "rpz.example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "10.10.0.1"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", "rpz.example.com"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAResource_disappears(t *testing.T) {
	resourceName := "nios_rpz_record_rpz_a.test"
	var v rpz.RecordRpzA
	name := acctest.RandomName() + ".rpz.example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzADestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzABasicConfig(name, "10.10.0.1", "rpz.example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					testAccCheckRecordRpzADisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordRpzAResource_Comment(t *testing.T) {
	var resourceName = "nios_rpz_record_rpz_a.test_comment"
	var v rpz.RecordRpzA
	name := acctest.RandomName() + ".rpz.example.com"
	comment1 := "This is a new rpz a record"
	comment2 := "This is a updated rpz a record"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAComment(name, "10.10.0.1", "rpz.example.com", comment1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAComment(name, "10.10.0.1", "rpz.example.com", comment2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAResource_Disable(t *testing.T) {
	var resourceName = "nios_rpz_record_rpz_a.test_disable"
	var v rpz.RecordRpzA
	name := acctest.RandomName() + ".rpz.example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzADisable(name, "10.10.0.1", "rpz.example.com", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzADisable(name, "10.10.0.1", "rpz.example.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_rpz_record_rpz_a.test_extattrs"
	var v rpz.RecordRpzA
	name := acctest.RandomName() + ".rpz.example.com"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAExtAttrs(name, "10.10.0.1", "rpz.example.com", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAExtAttrs(name, "10.10.0.1", "rpz.example.com", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAResource_Ipv4addr(t *testing.T) {
	var resourceName = "nios_rpz_record_rpz_a.test_ipv4addr"
	var v rpz.RecordRpzA
	name := acctest.RandomName() + ".rpz.example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAIpv4addr(name, "10.10.0.1", "rpz.example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "10.10.0.1"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAIpv4addr(name, "10.10.0.2", "rpz.example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "10.10.0.2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAResource_Name(t *testing.T) {
	var resourceName = "nios_rpz_record_rpz_a.test_name"
	var v rpz.RecordRpzA
	name1 := acctest.RandomName() + ".rpz.example.com"
	name2 := acctest.RandomName() + ".rpz.example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAName(name1, "10.10.0.1", "rpz.example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAName(name2, "10.10.0.1", "rpz.example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAResource_RpZone(t *testing.T) {
	var resourceName = "nios_rpz_record_rpz_a.test_rp_zone"
	var v rpz.RecordRpzA
	name := acctest.RandomName() + ".rpz.example.com"
	rpZone := "rpz.example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzARpZone(name, "10.10.0.1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
				),
			},
			// Can't update rp_zone as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAResource_Ttl(t *testing.T) {
	var resourceName = "nios_rpz_record_rpz_a.test_ttl"
	var v rpz.RecordRpzA
	name := acctest.RandomName() + ".rpz.example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzATtl(name, "10.10.0.1", "rpz.example.com", "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzATtl(name, "10.10.0.1", "rpz.example.com", "true", 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAResource_UseTtl(t *testing.T) {
	var resourceName = "nios_rpz_record_rpz_a.test_use_ttl"
	var v rpz.RecordRpzA
	name := acctest.RandomName() + ".rpz.example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAUseTtl(name, "10.10.0.1", "rpz.example.com", "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzAUseTtl(name, "10.10.0.1", "rpz.example.com", "false", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzAResource_View(t *testing.T) {
	var resourceName = "nios_rpz_record_rpz_a.test_view"
	var v rpz.RecordRpzA
	name := acctest.RandomName() + ".rpz.example.com"
	view := "custom_view_1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzAView(name, "10.10.0.1", "rpz.example.com", view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", view),
				),
			},
			// Can't update view as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordRpzAExists(ctx context.Context, resourceName string, v *rpz.RecordRpzA) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.RPZAPI.
			RecordRpzAAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordRpzA).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordRpzAResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordRpzAResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordRpzADestroy(ctx context.Context, v *rpz.RecordRpzA) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.RPZAPI.
			RecordRpzAAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordRpzA).
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

func testAccCheckRecordRpzADisappears(ctx context.Context, v *rpz.RecordRpzA) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.RPZAPI.
			RecordRpzAAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordRpzABasicConfig(name, ipV4Addr, rpZone string) string {
	return fmt.Sprintf(`
resource "nios_rpz_record_rpz_a" "test" {
	name = %q
	ipv4addr = %q
	rp_zone = %q
}
`, name, ipV4Addr, rpZone)
}

func testAccRecordRpzAComment(name, ipV4Addr, rpZone, comment string) string {
	return fmt.Sprintf(`
resource "nios_rpz_record_rpz_a" "test_comment" {
    name = %q
	ipv4addr = %q
	rp_zone = %q
	comment = %q
}
`, name, ipV4Addr, rpZone, comment)
}

func testAccRecordRpzADisable(name, ipV4Addr, rpZone, disable string) string {
	return fmt.Sprintf(`
resource "nios_rpz_record_rpz_a" "test_disable" {
	name = %q
	ipv4addr = %q
	rp_zone = %q
    disable = %q
}
`, name, ipV4Addr, rpZone, disable)
}

func testAccRecordRpzAExtAttrs(name, ipV4Addr, rpZone string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
	%s = %q
	`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_rpz_record_rpz_a" "test_extattrs" {
	name = %q
	ipv4addr = %q
	rp_zone = %q
    extattrs = %s
}
`, name, ipV4Addr, rpZone, extattrsStr)
}

func testAccRecordRpzAIpv4addr(name, ipV4Addr, rpZone string) string {
	return fmt.Sprintf(`
resource "nios_rpz_record_rpz_a" "test_ipv4addr" {
    name = %q
	ipv4addr = %q
	rp_zone = %q
}
`, name, ipV4Addr, rpZone)
}

func testAccRecordRpzAName(name, ipV4Addr, rpZone string) string {
	return fmt.Sprintf(`
resource "nios_rpz_record_rpz_a" "test_name" {
    name = %q
	ipv4addr = %q
	rp_zone = %q
}
`, name, ipV4Addr, rpZone)
}

func testAccRecordRpzARpZone(name, ipV4Addr, rpZone string) string {
	return fmt.Sprintf(`
resource "nios_rpz_record_rpz_a" "test_rp_zone" {
    name = %q
	ipv4addr = %q
	rp_zone = %q
}
`, name, ipV4Addr, rpZone)
}

func testAccRecordRpzATtl(name, ipV4Addr, rpZone string, use_ttl string, ttl int32) string {
	return fmt.Sprintf(`
resource "nios_rpz_record_rpz_a" "test_ttl" {
    name = %q
	ipv4addr = %q
	rp_zone = %q
	ttl = %d
	use_ttl = %q
}
`, name, ipV4Addr, rpZone, ttl, use_ttl)
}

func testAccRecordRpzAUseTtl(name, ipV4Addr, rpZone string, use_ttl string, ttl int32) string {
	return fmt.Sprintf(`
resource "nios_rpz_record_rpz_a" "test_use_ttl" {
    name = %q
	ipv4addr = %q
	rp_zone = %q
	ttl = %d
	use_ttl = %q
}
`, name, ipV4Addr, rpZone, ttl, use_ttl)
}

func testAccRecordRpzAView(name, ipV4Addr, rpZone, view string) string {
	return fmt.Sprintf(`
resource "nios_rpz_record_rpz_a" "test_view" {
    name = %q
	ipv4addr = %q
	rp_zone = %q
	view = %q
}
`, name, ipV4Addr, rpZone, view)
}
