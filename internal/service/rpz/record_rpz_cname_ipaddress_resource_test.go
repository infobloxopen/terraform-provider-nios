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

var readableAttributesForRecordRpzCnameIpaddress = "canonical,comment,disable,extattrs,is_ipv4,name,rp_zone,ttl,use_ttl,view,zone"

func TestAccRecordRpzCnameIpaddressResource_basic(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_ipaddress.test"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP := "11.0.0.1"
	name := nameIP + "." + rpZone
	canonical := "11.0.0.1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameIpaddressBasicConfig(nameIP, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
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

func TestAccRecordRpzCnameIpaddressResource_disappears(t *testing.T) {
	resourceName := "nios_rpz_record_cname_ipaddress.test"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	canonical := ""
	nameIP := "11.0.0.2"
	view := acctest.RandomNameWithPrefix("view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzCnameIpaddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzCnameIpaddressConfig(nameIP, canonical, rpZone, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					testAccCheckRecordRpzCnameIpaddressDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordRpzCnameIpaddressResource_Canonical(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_ipaddress.test"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP := "11.0.0.3"
	view := acctest.RandomNameWithPrefix("view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameIpaddressConfig(nameIP, nameIP, rpZone, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", nameIP),
				),
			},
			// Update to No Data rule
			{
				Config: testAccRecordRpzCnameIpaddressConfig(nameIP, "*", rpZone, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", "*"),
				),
			},
			// Update to Passthru Domain Name Rule
			{
				Config: testAccRecordRpzCnameIpaddressConfig(nameIP, "", rpZone, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", ""),
				),
			},
			// update to Substitution rule
			{
				Config: testAccRecordRpzCnameIpaddressConfig(nameIP, nameIP, rpZone, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", nameIP),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameIpaddressResource_Comment(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_ipaddress.test_comment"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP := "11.0.0.4"
	canonical := ""
	comment1 := "This is a new rpz cname record"
	comment2 := "This is an updated rpz cname record"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameIpaddressComment(nameIP, canonical, rpZone, comment1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameIpaddressComment(nameIP, canonical, rpZone, comment2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameIpaddressResource_Disable(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_ipaddress.test_disable"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP := "11.0.0.5"
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameIpaddressDisable(nameIP, canonical, rpZone, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameIpaddressDisable(nameIP, canonical, rpZone, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameIpaddressResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_ipaddress.test_extattrs"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP := "11.0.0.6"
	canonical := ""
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameIpaddressExtAttrs(nameIP, canonical, rpZone, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameIpaddressExtAttrs(nameIP, canonical, rpZone, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameIpaddressResource_Name(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_ipaddress.test_name"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP1 := "11.0.0.7"
	nameIP2 := "11.0.0.8"
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameIpaddressName(nameIP1, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameIP1+"."+rpZone),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameIpaddressName(nameIP2, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameIP2+"."+rpZone),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameIpaddressResource_RpZone(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_ipaddress.test_rp_zone"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP := "11.0.0.9"
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameIpaddressRpZone(nameIP, canonical, rpZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
				),
			},
			// Can't update rp_zone as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameIpaddressResource_Ttl(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_ipaddress.test_ttl"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP := "11.0.0.10"
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameIpaddressTtl(nameIP, canonical, rpZone, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameIpaddressTtl(nameIP, canonical, rpZone, "true", 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameIpaddressResource_UseTtl(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_ipaddress.test_use_ttl"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP := "11.0.0.11"
	canonical := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameIpaddressUseTtl(nameIP, canonical, rpZone, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzCnameIpaddressUseTtl(nameIP, canonical, rpZone, "false", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzCnameIpaddressResource_View(t *testing.T) {
	var resourceName = "nios_rpz_record_cname_ipaddress.test_view"
	var v rpz.RecordRpzCnameIpaddress
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	nameIP := "11.0.0.12"
	canonical := ""
	view := acctest.RandomNameWithPrefix("view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzCnameIpaddressView(nameIP, canonical, rpZone, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzCnameIpaddressExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", view),
				),
			},
			// Can't update view as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordRpzCnameIpaddressExists(ctx context.Context, resourceName string, v *rpz.RecordRpzCnameIpaddress) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.RPZAPI.
			RecordRpzCnameIpaddressAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordRpzCnameIpaddress).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordRpzCnameIpaddressResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordRpzCnameIpaddressResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordRpzCnameIpaddressDestroy(ctx context.Context, v *rpz.RecordRpzCnameIpaddress) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.RPZAPI.
			RecordRpzCnameIpaddressAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordRpzCnameIpaddress).
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

func testAccCheckRecordRpzCnameIpaddressDisappears(ctx context.Context, v *rpz.RecordRpzCnameIpaddress) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.RPZAPI.
			RecordRpzCnameIpaddressAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordRpzCnameIpaddressBasicConfig(nameIP, canonical, rpZone string) string {
	// create basic resource with required fields
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_ipaddress" "test" {
	name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
}
`, nameIP, canonical)
	return strings.Join([]string{testAccBaseWithZoneRPNetwork(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameIpaddressConfig(nameIP, canonical, rpZone, view string) string {
	// create resource with required fields
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_ipaddress" "test" {
	name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	view = nios_dns_view.custom_view.name
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
}
`, nameIP, canonical)
	return strings.Join([]string{testAccBaseWithView(view), testAccBaseWithZoneRPNetwork(rpZone, "nios_dns_view.custom_view.name"), config}, "")
}

func testAccBaseWithZoneRPNetwork(rpZone, view string) string {
	if view == "" {
		view = `"default"`
	}
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_zone" {
    fqdn = %q
	view = %s
}
`, rpZone, view)
}

func testAccRecordRpzCnameIpaddressComment(nameIp, canonical, rpZone, comment string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_ipaddress" "test_comment" {
	name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
    view = "default"
    comment = %q
}
`, nameIp, canonical, comment)
	return strings.Join([]string{testAccBaseWithZoneRPNetwork(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameIpaddressDisable(nameIP, canonical, rpZone, disable string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_ipaddress" "test_disable" {
    name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
	view = "default"
    disable = %q
}
`, nameIP, canonical, disable)
	return strings.Join([]string{testAccBaseWithZoneRPNetwork(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameIpaddressExtAttrs(nameIP, canonical, rpZone string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
	%s = %q
	`, k, v)
	}
	extattrsStr += "\t}"

	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_ipaddress" "test_extattrs" {
    name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
	view = "default"
    extattrs = %s
}
`, nameIP, canonical, extattrsStr)
	return strings.Join([]string{testAccBaseWithZoneRPNetwork(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameIpaddressName(nameIP, canonical, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_ipaddress" "test_name" {
    name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
	view = "default"
}
`, nameIP, canonical)
	return strings.Join([]string{testAccBaseWithZoneRPNetwork(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameIpaddressRpZone(nameIP, canonical, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_ipaddress" "test_rp_zone" {
	name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
	view = "default"
}
`, nameIP, canonical)
	return strings.Join([]string{testAccBaseWithZoneRPNetwork(rpZone, ""), config}, "")
}

func testAccRecordRpzCnameIpaddressTtl(nameIP, canonical, rpZone, useTtl string, ttl int32) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_ipaddress" "test_ttl" {
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

func testAccRecordRpzCnameIpaddressUseTtl(nameIP, canonical, rpZone, useTtl string, ttl int32) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_ipaddress" "test_use_ttl" {
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

func testAccRecordRpzCnameIpaddressView(nameIP, canonical, rpZone, view string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_cname_ipaddress" "test_view" {
	name = "%s.${nios_dns_zone_rp.test_zone.fqdn}"
	canonical = %q
	rp_zone = nios_dns_zone_rp.test_zone.fqdn
    view = nios_dns_view.custom_view.name
}
`, nameIP, canonical)
	return strings.Join([]string{testAccBaseWithView(view), testAccBaseWithZoneRPNetwork(rpZone, "nios_dns_view.custom_view.name"), config}, "")
}
