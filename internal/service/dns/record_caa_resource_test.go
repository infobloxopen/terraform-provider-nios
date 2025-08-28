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

var readableAttributesForRecordCaa = "ca_flag,ca_tag,ca_value,cloud_info,comment,creation_time,creator,ddns_principal,ddns_protected,disable,dns_name,extattrs,forbid_reclamation,last_queried,name,reclaimable,ttl,use_ttl,view,zone"

func TestAccRecordCaaResource_basic(t *testing.T) {
	var resourceName = "nios_dns_record_caa.test"
	var v dns.RecordCaa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-caa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCaaBasicConfig(zoneFqdn, name, 0, "issue", "digicert.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_flag", "0"),
					resource.TestCheckResourceAttr(resourceName, "ca_tag", "issue"),
					resource.TestCheckResourceAttr(resourceName, "ca_value", "digicert.com"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "false"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCaaResource_disappears(t *testing.T) {
	resourceName := "nios_dns_record_caa.test"
	var v dns.RecordCaa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-caa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordCaaDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordCaaBasicConfig(zoneFqdn, name, 0, "issue", "digicert.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					testAccCheckRecordCaaDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordCaaResource_CaFlag(t *testing.T) {
	var resourceName = "nios_dns_record_caa.test_ca_flag"
	var v dns.RecordCaa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-caa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCaaCaFlag(zoneFqdn, name, 0, "issue", "digicert.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_flag", "0"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCaaCaFlag(zoneFqdn, name, 1, "issue", "digicert.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_flag", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCaaResource_CaTag(t *testing.T) {
	var resourceName = "nios_dns_record_caa.test_ca_tag"
	var v dns.RecordCaa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-caa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCaaCaTag(zoneFqdn, name, 0, "issue", "digicert.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_tag", "issue"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCaaCaTag(zoneFqdn, name, 0, "issuewild", "digicert.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_tag", "issuewild"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCaaResource_CaValue(t *testing.T) {
	var resourceName = "nios_dns_record_caa.test_ca_value"
	var v dns.RecordCaa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-caa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCaaCaValue(zoneFqdn, name, 0, "issue", "digicert.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_value", "digicert.com"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCaaCaValue(zoneFqdn, name, 0, "issue", "letsencrypt.org"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_value", "letsencrypt.org"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCaaResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_record_caa.test_comment"
	var v dns.RecordCaa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-caa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCaaComment(zoneFqdn, name, 0, "issue", "digicert.com", "comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "comment"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCaaComment(zoneFqdn, name, 0, "issue", "digicert.com", "updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCaaResource_Creator(t *testing.T) {
	var resourceName = "nios_dns_record_caa.test_creator"
	var v dns.RecordCaa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-caa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCaaCreator(zoneFqdn, name, 0, "issue", "digicert.com", "STATIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCaaCreator(zoneFqdn, name, 0, "issue", "digicert.com", "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "DYNAMIC"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCaaResource_DdnsPrincipal(t *testing.T) {
	var resourceName = "nios_dns_record_caa.test_ddns_principal"
	var v dns.RecordCaa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-caa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCaaDdnsPrincipal(zoneFqdn, name, 0, "issue", "digicert.com", "ddns_principal", "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "ddns_principal"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCaaDdnsPrincipal(zoneFqdn, name, 0, "issue", "digicert.com", "updated_ddns_principal", "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "updated_ddns_principal"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCaaResource_DdnsProtected(t *testing.T) {
	var resourceName = "nios_dns_record_caa.test_ddns_protected"
	var v dns.RecordCaa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-caa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCaaDdnsProtected(zoneFqdn, name, 0, "issue", "digicert.com", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCaaDdnsProtected(zoneFqdn, name, 0, "issue", "digicert.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCaaResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_record_caa.test_disable"
	var v dns.RecordCaa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-caa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCaaDisable(zoneFqdn, name, 0, "issue", "digicert.com", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCaaDisable(zoneFqdn, name, 0, "issue", "digicert.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCaaResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_record_caa.test_extattrs"
	var v dns.RecordCaa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-caa")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCaaExtAttrs(zoneFqdn, name, 0, "issue", "digicert.com", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCaaExtAttrs(zoneFqdn, name, 0, "issue", "digicert.com", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCaaResource_ForbidReclamation(t *testing.T) {
	var resourceName = "nios_dns_record_caa.test_forbid_reclamation"
	var v dns.RecordCaa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-caa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCaaForbidReclamation(zoneFqdn, name, 0, "issue", "digicert.com", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCaaForbidReclamation(zoneFqdn, name, 0, "issue", "digicert.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCaaResource_Name(t *testing.T) {
	var resourceName = "nios_dns_record_caa.test_name"
	var v dns.RecordCaa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name1 := acctest.RandomNameWithPrefix("record-caa")
	name2 := acctest.RandomNameWithPrefix("record-caa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCaaName(zoneFqdn, name1, 0, "issue", "digicert.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("%s.%s", name1, zoneFqdn)),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCaaName(zoneFqdn, name2, 0, "issue", "digicert.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("%s.%s", name2, zoneFqdn)),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCaaResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_record_caa.test_ttl"
	var v dns.RecordCaa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-caa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCaaTtl(zoneFqdn, name, 0, "issue", "digicert.com", 10, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCaaTtl(zoneFqdn, name, 0, "issue", "digicert.com", 20, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCaaResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_record_caa.test_use_ttl"
	var v dns.RecordCaa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-caa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCaaUseTtl(zoneFqdn, name, 0, "issue", "digicert.com", 10, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCaaUseTtl(zoneFqdn, name, 0, "issue", "digicert.com", 10, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCaaResource_View(t *testing.T) {
	var resourceName = "nios_dns_record_caa.test_view"
	var v dns.RecordCaa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-caa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCaaView(zoneFqdn, name, 0, "issue", "digicert.com", "custom_view_1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "custom_view_1"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCaaView(zoneFqdn, name, 0, "issue", "digicert.com", "custom_view_2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "custom_view_2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordCaaExists(ctx context.Context, resourceName string, v *dns.RecordCaa) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			RecordCaaAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordCaa).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordCaaResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordCaaResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordCaaDestroy(ctx context.Context, v *dns.RecordCaa) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			RecordCaaAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordCaa).
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

func testAccCheckRecordCaaDisappears(ctx context.Context, v *dns.RecordCaa) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			RecordCaaAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordCaaBasicConfig(zoneFqdn, name string, caFlag int, caTag, caValue string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_caa" "test" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    ca_flag = %d
    ca_tag = %q
    ca_value = %q
}
`, name, caFlag, caTag, caValue)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordCaaCaFlag(zoneFqdn, name string, caFlag int, caTag string, caValue string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_caa" "test_ca_flag" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    ca_flag = %d
    ca_tag = %q
    ca_value = %q
}
`, name, caFlag, caTag, caValue)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordCaaCaTag(zoneFqdn, name string, caFlag int, caTag string, caValue string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_caa" "test_ca_tag" {
 	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"	
    ca_flag = %d
    ca_tag = %q
    ca_value = %q
}
`, name, caFlag, caTag, caValue)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordCaaCaValue(zoneFqdn, name string, caFlag int, caTag string, caValue string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_caa" "test_ca_value" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    ca_flag = %d
    ca_tag = %q
    ca_value = %q
}
`, name, caFlag, caTag, caValue)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordCaaComment(zoneFqdn, name string, caFlag int, caTag string, caValue string, comment string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_caa" "test_comment" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    ca_flag = %d
    ca_tag = %q
    ca_value = %q
    comment = %q
}
`, name, caFlag, caTag, caValue, comment)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordCaaCreator(zoneFqdn, name string, caFlag int, caTag string, caValue string, creator string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_caa" "test_creator" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    ca_flag = %d
    ca_tag = %q
    ca_value = %q
    creator = %q
}
`, name, caFlag, caTag, caValue, creator)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordCaaDdnsPrincipal(zoneFqdn, name string, caFlag int, caTag string, caValue string, ddnsPrincipal, creator string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_caa" "test_ddns_principal" {
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    ca_flag = %d
    ca_tag = %q
    ca_value = %q
    ddns_principal = %q
    creator = %q
}
`, name, caFlag, caTag, caValue, ddnsPrincipal, creator)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordCaaDdnsProtected(zoneFqdn, name string, caFlag int, caTag string, caValue string, ddnsProtected string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_caa" "test_ddns_protected" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    ca_flag = %d
    ca_tag = %q
    ca_value = %q
    ddns_protected = %q
}
`, name, caFlag, caTag, caValue, ddnsProtected)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordCaaDisable(zoneFqdn, name string, caFlag int, caTag string, caValue string, disable string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_caa" "test_disable" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    ca_flag = %d
    ca_tag = %q
    ca_value = %q
    disable = %q
}
`, name, caFlag, caTag, caValue, disable)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordCaaExtAttrs(zoneFqdn, name string, caFlag int, caTag string, caValue string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	config := fmt.Sprintf(`
resource "nios_dns_record_caa" "test_extattrs" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    ca_flag = %d
    ca_tag = %q
    ca_value = %q
    extattrs = %s
}
`, name, caFlag, caTag, caValue, extattrsStr)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordCaaForbidReclamation(zoneFqdn, name string, caFlag int, caTag string, caValue string, forbidReclamation string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_caa" "test_forbid_reclamation" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    ca_flag = %d
    ca_tag = %q
    ca_value = %q
    forbid_reclamation = %q
}
`, name, caFlag, caTag, caValue, forbidReclamation)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordCaaName(zoneFqdn, name string, caFlag int, caTag string, caValue string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_caa" "test_name" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    ca_flag = %d
    ca_tag = %q
    ca_value = %q
}
`, name, caFlag, caTag, caValue)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordCaaTtl(zoneFqdn, name string, caFlag int, caTag string, caValue string, ttl int, useTtl string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_caa" "test_ttl" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    ca_flag = %d
    ca_tag = %q
    ca_value = %q
    ttl = %d
    use_ttl = %q
}
`, name, caFlag, caTag, caValue, ttl, useTtl)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordCaaUseTtl(zoneFqdn, name string, caFlag int, caTag string, caValue string, ttl int, useTtl string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_caa" "test_use_ttl" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    ca_flag = %d
    ca_tag = %q
    ca_value = %q
    ttl = %d
    use_ttl = %q
}
`, name, caFlag, caTag, caValue, ttl, useTtl)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordCaaView(zoneFqdn, name string, caFlag int, caTag string, caValue string, view string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_caa" "test_view" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    ca_flag = %d
    ca_tag = %q
    ca_value = %q
    view = %q
}
`, name, caFlag, caTag, caValue, view)
	return strings.Join([]string{testAccBaseWithZoneandView(zoneFqdn, view), config}, "")
}

func testAccBaseWithZoneandView(zoneFqdn, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test" {
	name = %q
}
resource "nios_dns_zone_auth" "test" {
    fqdn = %q
    view = nios_dns_view.test.name
}
`, view, zoneFqdn)
}
