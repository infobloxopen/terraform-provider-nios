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

var readableAttributesForRecordDname = "cloud_info,comment,creation_time,creator,ddns_principal,ddns_protected,disable,dns_name,dns_target,extattrs,forbid_reclamation,last_queried,name,reclaimable,shared_record_group,target,ttl,use_ttl,view,zone"

func TestAccRecordDnameResource_basic(t *testing.T) {
	var resourceName = "nios_dns_record_dname.test"
	var v dns.RecordDname
	target := acctest.RandomNameWithPrefix("test-dname") + ".com"
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordDnameBasicConfig(target, zoneFqdn),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "target", target),
					resource.TestCheckResourceAttr(resourceName, "name", zoneFqdn),
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

func TestAccRecordDnameResource_disappears(t *testing.T) {
	resourceName := "nios_dns_record_dname.test"
	var v dns.RecordDname
	target := acctest.RandomNameWithPrefix("test-dname-target") + ".example.com"
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordDnameDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordDnameBasicConfig(target, zoneFqdn),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					testAccCheckRecordDnameDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordDnameResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_record_dname.test_comment"
	var v dns.RecordDname
	target := acctest.RandomNameWithPrefix("test-dname") + ".com"
	view := acctest.RandomNameWithPrefix("test-view")
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordDnameComment(target, view, zoneFqdn, "comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "comment"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordDnameComment(target, view, zoneFqdn, "updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordDnameResource_Creator(t *testing.T) {
	var resourceName = "nios_dns_record_dname.test_creator"
	var v dns.RecordDname
	target := acctest.RandomNameWithPrefix("test-dname") + ".com"
	view := acctest.RandomNameWithPrefix("test-view")
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordDnameCreator(target, view, zoneFqdn, "STATIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordDnameCreator(target, view, zoneFqdn, "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "DYNAMIC"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordDnameResource_DdnsPrincipal(t *testing.T) {
	var resourceName = "nios_dns_record_dname.test_ddns_principal"
	var v dns.RecordDname
	target := acctest.RandomNameWithPrefix("test-dname") + ".com"
	view := acctest.RandomNameWithPrefix("test-view")
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordDnameDdnsPrincipal(target, view, zoneFqdn, "ddns_principal", "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "ddns_principal"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordDnameDdnsPrincipal(target, view, zoneFqdn, "updated_ddns_principal", "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "updated_ddns_principal"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordDnameResource_DdnsProtected(t *testing.T) {
	var resourceName = "nios_dns_record_dname.test_ddns_protected"
	var v dns.RecordDname
	target := acctest.RandomNameWithPrefix("test-dname") + ".com"
	view := acctest.RandomNameWithPrefix("test-view")
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordDnameDdnsProtected(target, view, zoneFqdn, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordDnameDdnsProtected(target, view, zoneFqdn, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordDnameResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_record_dname.test_disable"
	var v dns.RecordDname
	target := acctest.RandomNameWithPrefix("test-dname") + ".com"
	view := acctest.RandomNameWithPrefix("test-view")
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordDnameDisable(target, view, zoneFqdn, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordDnameDisable(target, view, zoneFqdn, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordDnameResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_record_dname.test_extattrs"
	var v dns.RecordDname
	target := acctest.RandomNameWithPrefix("test-dname") + ".com"
	view := acctest.RandomNameWithPrefix("test-view")
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordDnameExtAttrs(target, view, zoneFqdn, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordDnameExtAttrs(target, view, zoneFqdn, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordDnameResource_ForbidReclamation(t *testing.T) {
	var resourceName = "nios_dns_record_dname.test_forbid_reclamation"
	var v dns.RecordDname
	target := acctest.RandomNameWithPrefix("test-dname") + ".com"
	view := acctest.RandomNameWithPrefix("test-view")
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordDnameForbidReclamation(target, view, zoneFqdn, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordDnameForbidReclamation(target, view, zoneFqdn, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordDnameResource_Name(t *testing.T) {
	var resourceName = "nios_dns_record_dname.test_name"
	var v dns.RecordDname
	target := acctest.RandomNameWithPrefix("test-dname") + ".com"
	view := acctest.RandomNameWithPrefix("test-view")
	zoneFqdn1 := acctest.RandomNameWithPrefix("test-zone") + ".com"
	zoneFqdn2 := acctest.RandomNameWithPrefix("test-zone-update") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordDnameName(target, view, zoneFqdn1, zoneFqdn2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", zoneFqdn1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordDnameNameUpdate(target, view, zoneFqdn1, zoneFqdn2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", zoneFqdn2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordDnameResource_Target(t *testing.T) {
	var resourceName = "nios_dns_record_dname.test_target"
	var v dns.RecordDname
	target1 := acctest.RandomNameWithPrefix("test-dname") + ".com"
	target2 := acctest.RandomNameWithPrefix("test-dname-update") + ".com"
	view := acctest.RandomNameWithPrefix("test-view")
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordDnameTarget(target1, view, zoneFqdn),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "target", target1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordDnameTarget(target2, view, zoneFqdn),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "target", target2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordDnameResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_record_dname.test_ttl"
	var v dns.RecordDname
	target := acctest.RandomNameWithPrefix("test-dname") + ".com"
	view := acctest.RandomNameWithPrefix("test-view")
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordDnameTtl(target, view, zoneFqdn, 10, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordDnameTtl(target, view, zoneFqdn, 20, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordDnameResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_record_dname.test_use_ttl"
	var v dns.RecordDname
	target := acctest.RandomNameWithPrefix("test-dname") + ".com"
	view := acctest.RandomNameWithPrefix("test-view")
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordDnameUseTtl(target, view, zoneFqdn, 10, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordDnameUseTtl(target, view, zoneFqdn, 10, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordDnameResource_View(t *testing.T) {
	var resourceName = "nios_dns_record_dname.test_view"
	var v dns.RecordDname
	target := acctest.RandomNameWithPrefix("test-dname") + ".com"
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordDnameView(target, "default", zoneFqdn),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordDnameExists(ctx context.Context, resourceName string, v *dns.RecordDname) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			RecordDnameAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordDname).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordDnameResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordDnameResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordDnameDestroy(ctx context.Context, v *dns.RecordDname) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			RecordDnameAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordDname).
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

func testAccCheckRecordDnameDisappears(ctx context.Context, v *dns.RecordDname) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			RecordDnameAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordDnameBasicConfig(target, zoneFqdn string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_dname" "test" {
    name   = nios_dns_zone_auth.test.fqdn
    target = %q
	view = nios_dns_zone_auth.test.view
}
`, target)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordDnameComment(target, view, zoneFqdn, comment string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_dname" "test_comment" {
    name   = nios_dns_zone_auth.test.fqdn
    target = %q
	comment = %q
	view = nios_dns_zone_auth.test.view
}
`, target, comment)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordDnameCreator(target, view, zoneFqdn, creator string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_dname" "test_creator" {
    name   = nios_dns_zone_auth.test.fqdn
    target = %q
	creator = %q
	view = nios_dns_zone_auth.test.view
}
`, target, creator)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordDnameDdnsPrincipal(target, view, zoneFqdn, ddnsPrincipal, creator string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_dname" "test_ddns_principal" {
    name   = nios_dns_zone_auth.test.fqdn
    target = %q
	ddns_principal = %q
	creator = %q
	view = nios_dns_zone_auth.test.view
}
`, target, ddnsPrincipal, creator)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordDnameDdnsProtected(target, view, zoneFqdn, ddnsProtected string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_dname" "test_ddns_protected" {
    name   = nios_dns_zone_auth.test.fqdn
    target = %q
	ddns_protected = %q
	view = nios_dns_zone_auth.test.view
}
`, target, ddnsProtected)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordDnameDisable(target, view, zoneFqdn, disable string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_dname" "test_disable" {
    name   = nios_dns_zone_auth.test.fqdn
    target = %q
	disable = %q
	view = nios_dns_zone_auth.test.view
}
`, target, disable)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordDnameExtAttrs(target, view, zoneFqdn string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	config := fmt.Sprintf(`
resource "nios_dns_record_dname" "test_extattrs" {
    name   = nios_dns_zone_auth.test.fqdn
    target = %q
    extattrs = %s
	view = nios_dns_zone_auth.test.view
}
`, target, extattrsStr)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordDnameForbidReclamation(target, view, zoneFqdn, forbidReclamation string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_dname" "test_forbid_reclamation" {
    name   = nios_dns_zone_auth.test.fqdn
    target = %q
    forbid_reclamation = %q
	view = nios_dns_zone_auth.test.view
}
`, target, forbidReclamation)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordDnameName(target, view, zoneFqdn1, zoneFqdn2 string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_dname" "test_name" {
    name   = nios_dns_zone_auth.test.fqdn
    target = %q
	view = nios_dns_zone_auth.test.view
}
`, target)
	return strings.Join([]string{testAccBaseWithTwoZones(zoneFqdn1, zoneFqdn2), config}, "")
}

func testAccRecordDnameNameUpdate(target, view, zoneFqdn1, zoneFqdn2 string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_dname" "test_name" {
    name   = nios_dns_zone_auth.updated_zone.fqdn
    target = %q
	view = nios_dns_zone_auth.test.view
}
`, target)
	return strings.Join([]string{testAccBaseWithTwoZones(zoneFqdn1, zoneFqdn2), config}, "")
}

func testAccRecordDnameTarget(target, view, zoneFqdn string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_dname" "test_target" {
	name   = nios_dns_zone_auth.test.fqdn
    target = %q
	view = nios_dns_zone_auth.test.view
}
`, target)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordDnameTtl(target, view, zoneFqdn string, ttl int, useTtl string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_dname" "test_ttl" {
	name   = nios_dns_zone_auth.test.fqdn
	target = %q
    ttl = %d
	use_ttl = %q
	view = nios_dns_zone_auth.test.view
}
`, target, ttl, useTtl)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordDnameUseTtl(target, view, zoneFqdn string, ttl int, useTtl string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_dname" "test_use_ttl" {
    name   = nios_dns_zone_auth.test.fqdn
	target = %q
    ttl = %d
	use_ttl = %q
	view = nios_dns_zone_auth.test.view

}
`, target, ttl, useTtl)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordDnameView(target, view, zoneFqdn string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_dname" "test_view" {
    name = nios_dns_zone_auth.test.fqdn
	target = %q
}	
`, target)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccBaseWithTwoZones(zoneFqdn1, zoneFqdn2 string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test" {
    fqdn = %q
}

resource "nios_dns_zone_auth" "updated_zone" {
    fqdn = %q
}
`, zoneFqdn1, zoneFqdn2)
}
