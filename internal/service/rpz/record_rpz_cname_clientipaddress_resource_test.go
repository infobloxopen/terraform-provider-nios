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

var readableAttributesForRecordRpzCnameClientipaddress = "canonical,comment,disable,extattrs,is_ipv4,name,rp_zone,ttl,use_ttl,view,zone"

func TestAccRecordRpzCnameClientipaddressResource_basic(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_clientipaddress.test"
	var v rpz.RecordRpzCnameClientipaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP := "12.0.0.1"
	name := nameIP + "." + rpZone
	canonical := "rpz-passthru"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameClientipaddressBasicConfig(nameIP, canonical, rpZone, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "canonical", canonical),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
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

func TestAccRecordRpzCnameClientipaddressResource_disappears(t *testing.T) {
	resourceName := "nios_rpz_record_cname_clientipaddress.test"
	var v rpz.RecordRpzCnameClientipaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	canonical := ""
	nameIP := "12.0.0.2"
	view := acctest.RandomNameWithPrefix("view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzCnameClientipaddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzCnameClientipaddressBasicConfig(nameIP, canonical, rpZone, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressExists(context.Background(), resourceName, &v),
					testAccCheckRecordRpzCnameClientipaddressDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordRpzCnameClientipaddressResource_Canonical(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_clientipaddress.test_canonical"
	var v rpz.RecordRpzCnameClientipaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP := "12.0.0.3"
	view := acctest.RandomNameWithPrefix("view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Block IP Address (No Such Domain) Rule
			{
				Config: testAccRecordRpzCnameClientipaddressCanonical(nameIP, "", rpZone, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", ""),
				),
			},
			// Update to No Data rule
			{
				Config: testAccRecordRpzCnameClientipaddressCanonical(nameIP, "*", rpZone, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", "*"),
				),
			},
			// Update to Passthru rule
			{
				Config: testAccRecordRpzCnameClientipaddressCanonical(nameIP, "rpz-passthru", rpZone, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", "rpz-passthru"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameClientipaddressResource_Comment(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_clientipaddress.test_comment"
	var v rpz.RecordRpzCnameClientipaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP := "12.0.0.4"
	canonical := ""
	comment1 := "This is a new rpz cname client IP address record"
	comment2 := "This is an updated rpz cname client IP address record"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameClientipaddressComment(nameIP, canonical, rpZone, comment1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameClientipaddressComment(nameIP, canonical, rpZone, comment2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameClientipaddressResource_Disable(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_clientipaddress.test_disable"
	var v rpz.RecordRpzCnameClientipaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP := "12.0.0.5"
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameClientipaddressDisable(nameIP, canonical, rpZone, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameClientipaddressDisable(nameIP, canonical, rpZone, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameClientipaddressResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_clientipaddress.test_extattrs"
	var v rpz.RecordRpzCnameClientipaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP := "12.0.0.6"
	canonical := ""
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameClientipaddressExtAttrs(nameIP, canonical, rpZone, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameClientipaddressExtAttrs(nameIP, canonical, rpZone, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameClientipaddressResource_Name(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_clientipaddress.test_name"
	var v rpz.RecordRpzCnameClientipaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP1 := "12.0.0.7"
	nameIP2 := "12.0.0.8"
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameClientipaddressName(nameIP1, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameIP1+"."+rpZone),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameClientipaddressName(nameIP2, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameIP2+"."+rpZone),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameClientipaddressResource_RpZone(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_clientipaddress.test_rp_zone"
	var v rpz.RecordRpzCnameClientipaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP := "12.0.0.9"
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameClientipaddressRpZone(nameIP, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameClientipaddressResource_Ttl(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_clientipaddress.test_ttl"
	var v rpz.RecordRpzCnameClientipaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP := "12.0.0.10"
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameClientipaddressTtl(nameIP, canonical, rpZone, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameClientipaddressTtl(nameIP, canonical, rpZone, "true", 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameClientipaddressResource_UseTtl(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_clientipaddress.test_use_ttl"
	var v rpz.RecordRpzCnameClientipaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP := "12.0.0.11"
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameClientipaddressUseTtl(nameIP, canonical, rpZone, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameClientipaddressUseTtl(nameIP, canonical, rpZone, "false", 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameClientipaddressResource_View(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_clientipaddress.test_view"
	var v rpz.RecordRpzCnameClientipaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP := "12.0.0.12"
	canonical := ""
	view := acctest.RandomNameWithPrefix("view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameClientipaddressView(nameIP, canonical, rpZone, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameClientipaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", view),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordRpzCnameClientipaddressExists(ctx context.Context, resourceName string, v *rpz.RecordRpzCnameClientipaddress) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.RPZAPI.
			RecordRpzCnameClientipaddressAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordRpzCnameClientipaddress).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordRpzCnameClientipaddressResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordRpzCnameClientipaddressResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordRpzCnameClientipaddressDestroy(ctx context.Context, v *rpz.RecordRpzCnameClientipaddress) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.RPZAPI.
			RecordRpzCnameClientipaddressAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordRpzCnameClientipaddress).
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

func testAccCheckRecordRpzCnameClientipaddressDisappears(ctx context.Context, v *rpz.RecordRpzCnameClientipaddress) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.RPZAPI.
			RecordRpzCnameClientipaddressAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordRpzCnameClientipaddressBasicConfig(nameIP, canonical, rpZone, view string) string {
	// create basic resource with required fields
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddress" "test" {
	name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
}
`, nameIP, canonical)
	if view != "" {
		config = fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddress" "test" {
	name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	view = nios_dns_view.custom_view.name
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
}
`, nameIP, canonical)
		return strings.Join([]string{testAccBaseWithView(view), testAccBaseWithZoneRPNetwork(rpZone, "nios_dns_view.custom_view.name"), config}, "")
	}
	return strings.Join([]string{testAccBaseWithZoneRPNetwork(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameClientipaddressCanonical(nameIP, canonical, rpZone, view string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddress" "test_canonical" {
	name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	view = nios_dns_view.custom_view.name
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
}
`, nameIP, canonical)
	return strings.Join([]string{testAccBaseWithView(view), testAccBaseWithZoneRPNetwork(rpZone, "nios_dns_view.custom_view.name"), config}, "")

}

func testAccRecordRpzCnameClientipaddressComment(nameIp, canonical, rpZone, comment string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddress" "test_comment" {
	name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
    view = "default"
    comment = %q
}
`, nameIp, canonical, comment)
	return strings.Join([]string{testAccBaseWithZoneRPNetwork(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameClientipaddressDisable(nameIP, canonical, rpZone, disable string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddress" "test_disable" {
    name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
	view = "default"
    disable = %q
}
`, nameIP, canonical, disable)
	return strings.Join([]string{testAccBaseWithZoneRPNetwork(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameClientipaddressExtAttrs(nameIP, canonical, rpZone string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
	%s = %q
	`, k, v)
	}
	extattrsStr += "\t}"

	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddress" "test_extattrs" {
name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
	view = "default"
    extattrs = %s
}
`, nameIP, canonical, extattrsStr)
	return strings.Join([]string{testAccBaseWithZoneRPNetwork(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameClientipaddressName(nameIP, canonical, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddress" "test_name" {
    name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
	view = "default"
}
`, nameIP, canonical)
	return strings.Join([]string{testAccBaseWithZoneRPNetwork(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameClientipaddressRpZone(nameIP, canonical, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddress" "test_rp_zone" {
    name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
	view = "default"
}
`, nameIP, canonical)
	return strings.Join([]string{testAccBaseWithZoneRPNetwork(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameClientipaddressTtl(nameIP, canonical, rpZone, useTtl string, ttl int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddress" "test_ttl" {
    name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
	view = "default"
	ttl = %d
	use_ttl = %q
}
`, nameIP, canonical, ttl, useTtl)
	return strings.Join([]string{testAccBaseWithZoneRPNetwork(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameClientipaddressUseTtl(nameIP, canonical, rpZone, useTtl string, ttl int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddress" "test_use_ttl" {
name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
	view = "default"
    use_ttl = %q
	ttl = %d
}
`, nameIP, canonical, useTtl, ttl)
	return strings.Join([]string{testAccBaseWithZoneRPNetwork(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameClientipaddressView(nameIP, canonical, rpZone, view string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_clientipaddress" "test_view" {
    name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
    view = nios_dns_view.custom_view.name
}
`, nameIP, canonical)
	return strings.Join([]string{testAccBaseWithView(view), testAccBaseWithZoneRPNetwork(rpZone, "nios_dns_view.custom_view.name"), config}, "")
}
