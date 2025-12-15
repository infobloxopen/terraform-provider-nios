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

var readableAttributesForRecordRpzPtr = "comment,disable,extattrs,ipv4addr,ipv6addr,name,ptrdname,rp_zone,ttl,use_ttl,view,zone"

func TestAccRecordRpzPtrResource_basic(t *testing.T) {
	var resourceName = "nios_rpz_record_ptr.test"
	var v rpz.RecordRpzPtr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	ptrDName := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzPtrBasicConfig(ptrDName, "10.10.0.1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "10.10.0.1"),
					resource.TestCheckResourceAttr(resourceName, "ptrdname", ptrDName),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
					resource.TestCheckResourceAttr(resourceName, "name", "1.0.10.10.in-addr.arpa."+rpZone),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzPtrResource_disappears(t *testing.T) {
	resourceName := "nios_rpz_record_ptr.test"
	var v rpz.RecordRpzPtr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	ptrDName := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzPtrDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzPtrBasicConfig(ptrDName, "10.10.0.1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					testAccCheckRecordRpzPtrDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordRpzPtrResource_Comment(t *testing.T) {
	var resourceName = "nios_rpz_record_ptr.test_comment"
	var v rpz.RecordRpzPtr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	ptrDName := acctest.RandomName() + "." + rpZone
	comment1 := "This is a new rpz ptr record"
	comment2 := "This is a updated rpz ptr record"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzPtrComment(ptrDName, "10.10.0.1", rpZone, comment1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzPtrComment(ptrDName, "10.10.0.1", rpZone, comment2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzPtrResource_Disable(t *testing.T) {
	var resourceName = "nios_rpz_record_ptr.test_disable"
	var v rpz.RecordRpzPtr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	ptrDName := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzPtrDisable(ptrDName, "10.10.0.1", rpZone, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzPtrDisable(ptrDName, "10.10.0.1", rpZone, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzPtrResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_rpz_record_ptr.test_extattrs"
	var v rpz.RecordRpzPtr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	ptrDName := acctest.RandomName() + "." + rpZone
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzPtrExtAttrs(ptrDName, "10.10.0.1", rpZone, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzPtrExtAttrs(ptrDName, "10.10.0.1", rpZone, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzPtrResource_Ipv4addr(t *testing.T) {
	var resourceName = "nios_rpz_record_ptr.test_ipv4addr"
	var v rpz.RecordRpzPtr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	ptrDName := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzPtrIpv4addr(ptrDName, "10.10.0.1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "10.10.0.1"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzPtrIpv4addr(ptrDName, "10.10.0.2", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "10.10.0.2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzPtrResource_Ipv6addr(t *testing.T) {
	var resourceName = "nios_rpz_record_ptr.test_ipv6addr"
	var v rpz.RecordRpzPtr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	ptrDName := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzPtrIpv6addr(ptrDName, "2002:1f93::12:1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", "2002:1f93::12:1"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzPtrIpv6addr(ptrDName, "2002:1f93::12:2", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", "2002:1f93::12:2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzPtrResource_Name(t *testing.T) {
	var resourceName = "nios_rpz_record_ptr.test_name"
	var v rpz.RecordRpzPtr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	ptrDName := acctest.RandomName() + "." + rpZone
	name1 := "1.0.10.10.in-addr.arpa." + rpZone
	name2 := "2.0.10.10.in-addr.arpa." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzPtrName(ptrDName, name1, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzPtrName(ptrDName, name2, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzPtrResource_Ptrdname(t *testing.T) {
	var resourceName = "nios_rpz_record_ptr.test_ptrdname"
	var v rpz.RecordRpzPtr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	ptrDName1 := acctest.RandomName() + "." + rpZone
	ptrDName2 := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzPtrPtrdname(ptrDName1, "10.10.0.1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ptrdname", ptrDName1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzPtrPtrdname(ptrDName2, "10.10.0.1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ptrdname", ptrDName2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzPtrResource_RpZone(t *testing.T) {
	var resourceName = "nios_rpz_record_ptr.test_rp_zone"
	var v rpz.RecordRpzPtr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	ptrDName := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzPtrRpZone(ptrDName, "10.10.0.1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
				),
			},
			// Can't update rp_zone as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzPtrResource_Ttl(t *testing.T) {
	var resourceName = "nios_rpz_record_ptr.test_ttl"
	var v rpz.RecordRpzPtr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	ptrDName := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzPtrTtl(ptrDName, "10.10.0.1", rpZone, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzPtrTtl(ptrDName, "10.10.0.1", rpZone, "true", 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzPtrResource_UseTtl(t *testing.T) {
	var resourceName = "nios_rpz_record_ptr.test_use_ttl"
	var v rpz.RecordRpzPtr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	ptrDName := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzPtrUseTtl(ptrDName, "10.10.0.1", rpZone, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzPtrUseTtl(ptrDName, "10.10.0.1", rpZone, "false", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzPtrResource_View(t *testing.T) {
	var resourceName = "nios_rpz_record_ptr.test_view"
	var v rpz.RecordRpzPtr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	ptrDName := acctest.RandomName() + "." + rpZone
	view := acctest.RandomNameWithPrefix("test-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzPtrView(ptrDName, "10.10.0.1", rpZone, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", view),
				),
			},
			// Can't update view as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordRpzPtrExists(ctx context.Context, resourceName string, v *rpz.RecordRpzPtr) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.RPZAPI.
			RecordRpzPtrAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordRpzPtr).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordRpzPtrResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordRpzPtrResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordRpzPtrDestroy(ctx context.Context, v *rpz.RecordRpzPtr) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.RPZAPI.
			RecordRpzPtrAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordRpzPtr).
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

func testAccCheckRecordRpzPtrDisappears(ctx context.Context, v *rpz.RecordRpzPtr) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.RPZAPI.
			RecordRpzPtrAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordRpzPtrBasicConfig(ptrdname, ipV4Addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_ptr" "test" {
	ptrdname = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, ptrdname, ipV4Addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzPtrComment(ptrdname, ipV4Addr, rpZone, comment string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_ptr" "test_comment" {
	ptrdname = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    comment = %q
}
`, ptrdname, ipV4Addr, comment)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzPtrDisable(ptrdname, ipV4Addr, rpZone, disable string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_ptr" "test_disable" {
	ptrdname = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    disable = %q
}
`, ptrdname, ipV4Addr, disable)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzPtrExtAttrs(ptrdname, ipV4Addr, rpZone string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
	%s = %q
	`, k, v)
	}
	extattrsStr += "\t}"
	config := fmt.Sprintf(`
resource "nios_rpz_record_ptr" "test_extattrs" {
	ptrdname = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
    extattrs = %s
}
`, ptrdname, ipV4Addr, extattrsStr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzPtrIpv4addr(ptrdname, ipV4Addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_ptr" "test_ipv4addr" {
    ptrdname = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, ptrdname, ipV4Addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzPtrIpv6addr(ptrdname, ipV6Addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_ptr" "test_ipv6addr" {
    ptrdname = %q
    ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, ptrdname, ipV6Addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzPtrName(ptrdname, name, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_ptr" "test_name" {
 	ptrdname = %q
    name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, ptrdname, name)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzPtrPtrdname(ptrdname, ipV4Addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_ptr" "test_ptrdname" {
    ptrdname = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, ptrdname, ipV4Addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzPtrRpZone(ptrdname, ipV4Addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_ptr" "test_rp_zone" {
    ptrdname = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}
`, ptrdname, ipV4Addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzPtrTtl(ptrdname, ipV4Addr, rpZone, use_ttl string, ttl int32) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_ptr" "test_ttl" {
 	ptrdname = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	ttl = %d
	use_ttl = %q
}
`, ptrdname, ipV4Addr, ttl, use_ttl)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzPtrUseTtl(ptrdname, ipV4Addr, rpZone, use_ttl string, ttl int32) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_ptr" "test_use_ttl" {
    ptrdname = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	ttl = %d
	use_ttl = %q
}
`, ptrdname, ipV4Addr, ttl, use_ttl)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzPtrView(ptrDName, ipV4Addr, rpZone, view string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_ptr" "test_view" {
    ptrdname = %q
	ipv4addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	view = nios_dns_view.custom_view.name
}
`, ptrDName, ipV4Addr)

	return strings.Join([]string{testAccBaseWithView(view), testAccBaseWithZone(rpZone, "nios_dns_view.custom_view.name"), config}, "")
}
