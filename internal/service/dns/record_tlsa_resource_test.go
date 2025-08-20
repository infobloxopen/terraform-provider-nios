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

var readableAttributesForRecordTlsa = "certificate_data,certificate_usage,cloud_info,comment,creator,disable,dns_name,extattrs,last_queried,matched_type,name,selector,ttl,use_ttl,view,zone"

func TestAccRecordTlsaResource_basic(t *testing.T) {
	var resourceName = "nios_dns_record_tlsa.test"
	var v dns.RecordTlsa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-tlsa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTlsaBasicConfig(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "certificate_data", "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971"),
					resource.TestCheckResourceAttr(resourceName, "certificate_usage", "2"),
					resource.TestCheckResourceAttr(resourceName, "matched_type", "0"),
					resource.TestCheckResourceAttr(resourceName, "selector", "0"),
					// Test fields with default values
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTlsaResource_disappears(t *testing.T) {
	resourceName := "nios_dns_record_tlsa.test"
	var v dns.RecordTlsa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-tlsa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordTlsaDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordTlsaBasicConfig(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					testAccCheckRecordTlsaDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordTlsaResource_CertificateData(t *testing.T) {
	var resourceName = "nios_dns_record_tlsa.test_certificate_data"
	var v dns.RecordTlsa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-tlsa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTlsaCertificateData(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "certificate_data", "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTlsaCertificateData(zoneFqdn, name, "E2ABDE240D7CD3EE6B4B28C54DF134B97983A1D16E8A410E4561CB106618E972", 2, 0, 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "certificate_data", "E2ABDE240D7CD3EE6B4B28C54DF134B97983A1D16E8A410E4561CB106618E972"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTlsaResource_CertificateUsage(t *testing.T) {
	var resourceName = "nios_dns_record_tlsa.test_certificate_usage"
	var v dns.RecordTlsa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-tlsa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTlsaCertificateUsage(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "certificate_usage", "2"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTlsaCertificateUsage(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 3, 0, 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "certificate_usage", "3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTlsaResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_record_tlsa.test_comment"
	var v dns.RecordTlsa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-tlsa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTlsaComment(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTlsaComment(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTlsaResource_Creator(t *testing.T) {
	var resourceName = "nios_dns_record_tlsa.test_creator"
	var v dns.RecordTlsa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-tlsa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTlsaCreator(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0, "STATIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTlsaCreator(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0, "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "DYNAMIC"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTlsaResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_record_tlsa.test_disable"
	var v dns.RecordTlsa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-tlsa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTlsaDisable(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTlsaDisable(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTlsaResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_record_tlsa.test_extattrs"
	var v dns.RecordTlsa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-tlsa")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTlsaExtAttrs(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTlsaExtAttrs(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTlsaResource_MatchedType(t *testing.T) {
	var resourceName = "nios_dns_record_tlsa.test_matched_type"
	var v dns.RecordTlsa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-tlsa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTlsaMatchedType(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "matched_type", "0"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTlsaMatchedType(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 1, 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "matched_type", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTlsaResource_Name(t *testing.T) {
	var resourceName = "nios_dns_record_tlsa.test_name"
	var v dns.RecordTlsa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name1 := acctest.RandomNameWithPrefix("record-tlsa")
	name2 := acctest.RandomNameWithPrefix("record-tlsa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTlsaName(zoneFqdn, name1, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("%s.%s", name1, zoneFqdn)),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTlsaName(zoneFqdn, name2, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("%s.%s", name2, zoneFqdn)),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTlsaResource_Selector(t *testing.T) {
	var resourceName = "nios_dns_record_tlsa.test_selector"
	var v dns.RecordTlsa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-tlsa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTlsaSelector(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "selector", "0"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTlsaSelector(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "selector", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTlsaResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_record_tlsa.test_ttl"
	var v dns.RecordTlsa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-tlsa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTlsaTtl(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0, 10, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTlsaTtl(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0, 20, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTlsaResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_record_tlsa.test_use_ttl"
	var v dns.RecordTlsa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-tlsa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTlsaUseTtl(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0, "false", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTlsaUseTtl(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0, "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordTlsaResource_View(t *testing.T) {
	var resourceName = "nios_dns_record_tlsa.test_view"
	var v dns.RecordTlsa
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-tlsa")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTlsaView(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0, "custom_view_1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "custom_view_1"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTlsaView(zoneFqdn, name, "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971", 2, 0, 0, "custom_view_2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordTlsaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "custom_view_2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordTlsaExists(ctx context.Context, resourceName string, v *dns.RecordTlsa) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			RecordTlsaAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordTlsa).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordTlsaResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordTlsaResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordTlsaDestroy(ctx context.Context, v *dns.RecordTlsa) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			RecordTlsaAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordTlsa).
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

func testAccCheckRecordTlsaDisappears(ctx context.Context, v *dns.RecordTlsa) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			RecordTlsaAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordTlsaBasicConfig(zoneFqdn, name, certificateData string, certificateUsage, matchedType, selector int) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_tlsa" "test" {
    certificate_data = %q
    certificate_usage = %d
    matched_type = %d
    selector = %d
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
}
`, certificateData, certificateUsage, matchedType, selector, name)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordTlsaCertificateData(zoneFqdn, name, certificateData string, certificateUsage, matchedType, selector int) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_tlsa" "test_certificate_data" {
    certificate_data = %q
    certificate_usage = %d
    matched_type = %d
    selector = %d
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
}
`, certificateData, certificateUsage, matchedType, selector, name)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordTlsaCertificateUsage(zoneFqdn, name, certificateData string, certificateUsage, matchedType, selector int) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_tlsa" "test_certificate_usage" {
    certificate_data = %q
    certificate_usage = %d
    matched_type = %d
    selector = %d
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
}
`, certificateData, certificateUsage, matchedType, selector, name)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordTlsaComment(zoneFqdn, name, certificateData string, certificateUsage, matchedType, selector int, comment string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_tlsa" "test_comment" {
    certificate_data = %q
    certificate_usage = %d
    matched_type = %d
    selector = %d
    comment = %q
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
}
`, certificateData, certificateUsage, matchedType, selector, comment, name)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordTlsaCreator(zoneFqdn, name, certificateData string, certificateUsage, matchedType, selector int, creator string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_tlsa" "test_creator" {
    certificate_data = %q
    certificate_usage = %d
    matched_type = %d
    selector = %d
    creator = %q
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
}
`, certificateData, certificateUsage, matchedType, selector, creator, name)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordTlsaDisable(zoneFqdn, name, certificateData string, certificateUsage, matchedType, selector int, disable string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_tlsa" "test_disable" {
    certificate_data = %q
    certificate_usage = %d
    matched_type = %d
    selector = %d
    disable = %q
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
}
`, certificateData, certificateUsage, matchedType, selector, disable, name)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordTlsaExtAttrs(zoneFqdn, name, certificateData string, certificateUsage, matchedType, selector int, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	config := fmt.Sprintf(`
resource "nios_dns_record_tlsa" "test_extattrs" {
    certificate_data = %q
    certificate_usage = %d
    matched_type = %d
    selector = %d
    extattrs = %s
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
}
`, certificateData, certificateUsage, matchedType, selector, extattrsStr, name)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordTlsaMatchedType(zoneFqdn, name, certificateData string, certificateUsage, matchedType, selector int) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_tlsa" "test_matched_type" {
    certificate_data = %q
    certificate_usage = %d
    matched_type = %d
    selector = %d
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
}
`, certificateData, certificateUsage, matchedType, selector, name)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordTlsaName(zoneFqdn, name, certificateData string, certificateUsage, matchedType, selector int) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_tlsa" "test_name" {
    certificate_data = %q
    certificate_usage = %d
    matched_type = %d
    selector = %d
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
}
`, certificateData, certificateUsage, matchedType, selector, name)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordTlsaSelector(zoneFqdn, name, certificateData string, certificateUsage, matchedType, selector int) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_tlsa" "test_selector" {
    certificate_data = %q
    certificate_usage = %d
    matched_type = %d
    selector = %d
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
}
`, certificateData, certificateUsage, matchedType, selector, name)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordTlsaTtl(zoneFqdn, name, certificateData string, certificateUsage, matchedType, selector int, ttl int, useTtl string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_tlsa" "test_ttl" {
    certificate_data = %q
    certificate_usage = %d
    matched_type = %d
    selector = %d
    ttl = %d
	use_ttl = %q
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
}
`, certificateData, certificateUsage, matchedType, selector, ttl, useTtl, name)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordTlsaUseTtl(zoneFqdn, name, certificateData string, certificateUsage, matchedType, selector int, useTtl string, ttl int) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_tlsa" "test_use_ttl" {
    certificate_data = %q
    certificate_usage = %d
    matched_type = %d
    selector = %d
    use_ttl = %q
	ttl = %d
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
}
`, certificateData, certificateUsage, matchedType, selector, useTtl, ttl, name)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordTlsaView(zoneFqdn, name, certificateData string, certificateUsage int, matchedType, selector int, view string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_tlsa" "test_view" {
    certificate_data = %q
    certificate_usage = %d
    matched_type = %d
    selector = %d
    view = %q
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
}
`, certificateData, certificateUsage, matchedType, selector, view, name)
	return strings.Join([]string{testAccBaseWithZoneandView(zoneFqdn, view), config}, "")
}

func testAccBaseWithZone(zoneFqdn string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test" {
    fqdn = %q
}
`, zoneFqdn)
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
