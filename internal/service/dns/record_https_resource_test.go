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

var readableAttributesForRecordHttps = ""

func TestAccRecordHttpsResource_basic(t *testing.T) {
	var resourceName = "nios_dns_record_https.test"
	var v dns.RecordHttps
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-https")
	priority := "10"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordHttpsBasicConfig(zoneFqdn, name, priority),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordHttpsResource_disappears(t *testing.T) {
	resourceName := "nios_dns_record_https.test"
	var v dns.RecordHttps
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-https")
	priority := "10"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordHttpsDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordHttpsBasicConfig(zoneFqdn, name, priority),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					testAccCheckRecordHttpsDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordHttpsResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_record_https.test_comment"
	var v dns.RecordHttps
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-https")
	priority := "10"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordHttpsComment(zoneFqdn, name, priority, "Example HTTPS Record"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Example HTTPS Record"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordHttpsComment(zoneFqdn, name, priority, "Example HTTPS Record Updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Example HTTPS Record Updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordHttpsResource_Creator(t *testing.T) {
	var resourceName = "nios_dns_record_https.test_creator"
	var v dns.RecordHttps
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-https")
	priority := "10"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordHttpsCreator(zoneFqdn, name, priority, "STATIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordHttpsCreator(zoneFqdn, name, priority, "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "DYNAMIC"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordHttpsResource_DdnsPrincipal(t *testing.T) {
	var resourceName = "nios_dns_record_https.test_ddns_principal"
	var v dns.RecordHttps
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-https")
	priority := "10"
	ddnsPrincipal1 := acctest.RandomNameWithPrefix("test-ddns-principal")
	ddnsPrincipal2 := acctest.RandomNameWithPrefix("test-ddns-principal")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordHttpsDdnsPrincipal(zoneFqdn, name, priority, ddnsPrincipal1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", ddnsPrincipal1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordHttpsDdnsPrincipal(zoneFqdn, name, priority, ddnsPrincipal2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", ddnsPrincipal2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordHttpsResource_DdnsProtected(t *testing.T) {
	var resourceName = "nios_dns_record_https.test_ddns_protected"
	var v dns.RecordHttps
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-https")
	priority := "10"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordHttpsDdnsProtected(zoneFqdn, name, priority, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordHttpsDdnsProtected(zoneFqdn, name, priority, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordHttpsResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_record_https.test_disable"
	var v dns.RecordHttps
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-https")
	priority := "10"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordHttpsDisable(zoneFqdn, name, priority, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordHttpsDisable(zoneFqdn, name, priority, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordHttpsResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_record_https.test_extattrs"
	var v dns.RecordHttps
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-https")
	priority := "10"
	extAttrs1 := acctest.RandomName()
	extAttrs2 := acctest.RandomName()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordHttpsExtAttrs(zoneFqdn, name, priority, map[string]any{
					"Site": extAttrs1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrs1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordHttpsExtAttrs(zoneFqdn, name, priority, map[string]any{
					"Site": extAttrs2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrs2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordHttpsResource_ForbidReclamation(t *testing.T) {
	var resourceName = "nios_dns_record_https.test_forbid_reclamation"
	var v dns.RecordHttps
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-https")
	priority := "10"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordHttpsForbidReclamation(zoneFqdn, name, priority, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordHttpsForbidReclamation(zoneFqdn, name, priority, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordHttpsResource_Name(t *testing.T) {
	var resourceName = "nios_dns_record_https.test_name"
	var v dns.RecordHttps
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name1 := acctest.RandomNameWithPrefix("record-https")
	name2 := acctest.RandomNameWithPrefix("record-https")
	priority := "10"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordHttpsName(zoneFqdn, name1, priority),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1+"."+zoneFqdn),
				),
			},
			// Update and Read
			{
				Config: testAccRecordHttpsName(zoneFqdn, name2, priority),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2+"."+zoneFqdn),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordHttpsResource_Priority(t *testing.T) {
	var resourceName = "nios_dns_record_https.test_priority"
	var v dns.RecordHttps
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-https")
	priority1 := "10"
	priority2 := "20"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordHttpsPriority(zoneFqdn, name, priority1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "priority", priority1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordHttpsPriority(zoneFqdn, name, priority2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "priority", priority2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordHttpsResource_SvcParameters(t *testing.T) {
	var resourceName = "nios_dns_record_https.test_svc_parameters"
	var v dns.RecordHttps
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-https")
	priority := "10"
	svcParameters1 := []map[string]any{
		{
			"mandatory": true,
			"svc_key":   "port",
			"svc_value": []string{"110"},
		},
		{
			"mandatory": false,
			"svc_key":   "ipv4hint",
			"svc_value": []string{"11.11.1.1"},
		},
		{
			"mandatory": false,
			"svc_key":   "ipv6hint",
			"svc_value": []string{"123::99:0"},
		},
	}
	svcParameters2 := []map[string]any{
		{
			"mandatory": true,
			"svc_key":   "port",
			"svc_value": []string{"11"},
		},
		{
			"mandatory": false,
			"svc_key":   "ipv6hint",
			"svc_value": []string{"124::92:0"},
		},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordHttpsSvcParameters(zoneFqdn, name, priority, svcParameters1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.0.mandatory", "true"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.0.svc_key", "port"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.0.svc_value.0", "110"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.1.mandatory", "false"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.1.svc_key", "ipv4hint"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.1.svc_value.0", "11.11.1.1"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.2.svc_key", "ipv6hint"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.2.svc_value.0", "123::99:0"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordHttpsSvcParameters(zoneFqdn, name, priority, svcParameters2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.0.mandatory", "true"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.0.svc_key", "port"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.0.svc_value.0", "11"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.1.svc_key", "ipv6hint"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.1.svc_value.0", "124::92:0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordHttpsResource_TargetName(t *testing.T) {
	var resourceName = "nios_dns_record_https.test_target_name"
	var v dns.RecordHttps
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-https")
	priority := "10"
	targetName1 := "example." + zoneFqdn
	targetName2 := "example1." + zoneFqdn
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordHttpsTargetName(zoneFqdn, name, priority, targetName1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "target_name", targetName1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordHttpsTargetName(zoneFqdn, name, priority, targetName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "target_name", targetName2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordHttpsResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_record_https.test_ttl"
	var v dns.RecordHttps
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-https")
	priority := "10"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordHttpsTtl(zoneFqdn, name, priority, "200", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "200"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordHttpsTtl(zoneFqdn, name, priority, "300", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "300"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordHttpsResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_record_https.test_use_ttl"
	var v dns.RecordHttps
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-https")
	priority := "10"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordHttpsUseTtl(zoneFqdn, name, priority, "100", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordHttpsUseTtl(zoneFqdn, name, priority, "100", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordHttpsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordHttpsExists(ctx context.Context, resourceName string, v *dns.RecordHttps) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			RecordHttpsAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordHttps).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordHttpsResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordHttpsResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordHttpsDestroy(ctx context.Context, v *dns.RecordHttps) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			RecordHttpsAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordHttps).
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

func testAccCheckRecordHttpsDisappears(ctx context.Context, v *dns.RecordHttps) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			RecordHttpsAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordHttpsBasicConfig(zoneFqdn, name, priority string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_https" "test" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
}
`, name, zoneFqdn, priority)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordHttpsComment(zoneFqdn, name, priority, comment string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_https" "test_comment" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
   comment = %q
}
`, name, zoneFqdn, priority, comment)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordHttpsCreator(zoneFqdn, name, priority, creator string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_https" "test_creator" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
   creator = %q
}
`, name, zoneFqdn, priority, creator)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordHttpsDdnsPrincipal(zoneFqdn, name, priority, ddnsPrincipal string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_https" "test_ddns_principal" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
   ddns_principal = %q
}
`, name, zoneFqdn, priority, ddnsPrincipal)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordHttpsDdnsProtected(zoneFqdn, name, priority, ddnsProtected string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_https" "test_ddns_protected" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
   ddns_protected = %q
}
`, name, zoneFqdn, priority, ddnsProtected)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordHttpsDisable(zoneFqdn, name, priority, disable string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_https" "test_disable" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
   disable = %q
}
`, name, zoneFqdn, priority, disable)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordHttpsExtAttrs(zoneFqdn, name, priority string, extAttrs map[string]any) string {
	extAttrsStr := utils.ConvertMapToHCL(extAttrs)
	config := fmt.Sprintf(`
resource "nios_dns_record_https" "test_extattrs" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
   extattrs = %s
}
`, name, zoneFqdn, priority, extAttrsStr)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordHttpsForbidReclamation(zoneFqdn, name, priority, forbidReclamation string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_https" "test_forbid_reclamation" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
   forbid_reclamation = %q
}
`, name, zoneFqdn, priority, forbidReclamation)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordHttpsName(zoneFqdn, name, priority string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_https" "test_name" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
}
`, name, zoneFqdn, priority)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordHttpsPriority(zoneFqdn, name, priority string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_https" "test_priority" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
}
`, name, zoneFqdn, priority)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordHttpsSvcParameters(zoneFqdn, name, priority string, svcParameters []map[string]any) string {
	svcParamsStr := utils.ConvertSliceOfMapsToHCL(svcParameters)
	config := fmt.Sprintf(`
resource "nios_dns_record_https" "test_svc_parameters" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
   svc_parameters = %s
}
`, name, zoneFqdn, priority, svcParamsStr)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordHttpsTargetName(zoneFqdn, name, priority, targetName string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_https" "test_target_name" {
   name = "%s.%s"
   priority = %q
   target_name = %q
   depends_on = [nios_dns_zone_auth.test]
}
`, name, zoneFqdn, priority, targetName)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordHttpsTtl(zoneFqdn, name, priority, ttl, useTtl string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_https" "test_ttl" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
   ttl = %q
   use_ttl = %q
}
`, name, zoneFqdn, priority, ttl, useTtl)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordHttpsUseTtl(zoneFqdn, name, priority, ttl, useTtl string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_https" "test_use_ttl" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
   ttl = %q
   use_ttl = %q
}
`, name, zoneFqdn, priority, ttl, useTtl)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}
