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

var readableAttributesForRecordNaptr = "cloud_info,comment,creation_time,creator,ddns_principal,ddns_protected,disable,dns_name,dns_replacement,extattrs,flags,forbid_reclamation,last_queried,name,order,preference,reclaimable,regexp,replacement,services,ttl,use_ttl,view,zone"

func TestAccRecordNaptrResource_basic(t *testing.T) {
	var resourceName = "nios_dns_record_naptr.test"
	var v dns.RecordNaptr
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrBasicConfig(zoneFqdn, 10, 10, "."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", zoneFqdn),
					resource.TestCheckResourceAttr(resourceName, "order", "10"),
					resource.TestCheckResourceAttr(resourceName, "preference", "10"),
					resource.TestCheckResourceAttr(resourceName, "replacement", "."),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "false"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "regexp", ""),
					resource.TestCheckResourceAttr(resourceName, "services", ""),
					resource.TestCheckResourceAttr(resourceName, "flags", ""),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordNaptrResource_disappears(t *testing.T) {
	resourceName := "nios_dns_record_naptr.test"
	var v dns.RecordNaptr
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordNaptrDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordNaptrBasicConfig(zoneFqdn, 10, 10, "."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					testAccCheckRecordNaptrDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordNaptrResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_record_naptr.test_comment"
	var v dns.RecordNaptr
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrComment(zoneFqdn, 10, 10, ".", "comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "comment"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrComment(zoneFqdn, 10, 10, ".", "updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordNaptrResource_Creator(t *testing.T) {
	var resourceName = "nios_dns_record_naptr.test_creator"
	var v dns.RecordNaptr
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrCreator(zoneFqdn, 10, 10, ".", "STATIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrCreator(zoneFqdn, 10, 10, ".", "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "DYNAMIC"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordNaptrResource_DdnsPrincipal(t *testing.T) {
	var resourceName = "nios_dns_record_naptr.test_ddns_principal"
	var v dns.RecordNaptr
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrDdnsPrincipal(zoneFqdn, 10, 10, ".", "ddns_principal", "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "ddns_principal"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrDdnsPrincipal(zoneFqdn, 10, 10, ".", "updated_ddns_principal", "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "updated_ddns_principal"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordNaptrResource_DdnsProtected(t *testing.T) {
	var resourceName = "nios_dns_record_naptr.test_ddns_protected"
	var v dns.RecordNaptr
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrDdnsProtected(zoneFqdn, 10, 10, ".", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrDdnsProtected(zoneFqdn, 10, 10, ".", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordNaptrResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_record_naptr.test_disable"
	var v dns.RecordNaptr
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrDisable(zoneFqdn, 10, 10, ".", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrDisable(zoneFqdn, 10, 10, ".", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordNaptrResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_record_naptr.test_extattrs"
	var v dns.RecordNaptr
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrExtAttrs(zoneFqdn, 10, 10, ".", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrExtAttrs(zoneFqdn, 10, 10, ".", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordNaptrResource_Flags(t *testing.T) {
	var resourceName = "nios_dns_record_naptr.test_flags"
	var v dns.RecordNaptr
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrFlags(zoneFqdn, 10, 10, ".", "U"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "flags", "U"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrFlags(zoneFqdn, 10, 10, ".", "S"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "flags", "S"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordNaptrResource_ForbidReclamation(t *testing.T) {
	var resourceName = "nios_dns_record_naptr.test_forbid_reclamation"
	var v dns.RecordNaptr
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrForbidReclamation(zoneFqdn, 10, 10, ".", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrForbidReclamation(zoneFqdn, 10, 10, ".", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordNaptrResource_Name(t *testing.T) {
	var resourceName = "nios_dns_record_naptr.test_name"
	var v dns.RecordNaptr
	zoneFqdn1 := acctest.RandomNameWithPrefix("test-zone") + ".com"
	zoneFqdn2 := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrName(zoneFqdn1, zoneFqdn2, 10, 10, "."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", zoneFqdn1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrNameUpdate(zoneFqdn1, zoneFqdn2, 10, 10, "."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", zoneFqdn2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordNaptrResource_Order(t *testing.T) {
	var resourceName = "nios_dns_record_naptr.test_order"
	var v dns.RecordNaptr
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrOrder(zoneFqdn, 10, 10, "."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "order", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrOrder(zoneFqdn, 20, 10, "."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "order", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordNaptrResource_Preference(t *testing.T) {
	var resourceName = "nios_dns_record_naptr.test_preference"
	var v dns.RecordNaptr
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrPreference(zoneFqdn, 10, 10, "."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preference", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrPreference(zoneFqdn, 10, 20, "."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preference", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordNaptrResource_Regexp(t *testing.T) {
	var resourceName = "nios_dns_record_naptr.test_regexp"
	var v dns.RecordNaptr
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrRegexp(zoneFqdn, 10, 10, ".", ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "regexp", ""),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrRegexp(zoneFqdn, 10, 10, ".", "!^.*$!sip:jdoe@corpxyz.com!"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "regexp", "!^.*$!sip:jdoe@corpxyz.com!"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordNaptrResource_Replacement(t *testing.T) {
	var resourceName = "nios_dns_record_naptr.test_replacement"
	var v dns.RecordNaptr
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrReplacement(zoneFqdn, 10, 10, "."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "replacement", "."),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrReplacement(zoneFqdn, 10, 20, "test.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "replacement", "test.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordNaptrResource_Services(t *testing.T) {
	var resourceName = "nios_dns_record_naptr.test_services"
	var v dns.RecordNaptr
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrServices(zoneFqdn, 10, 10, ".", "http+E2U"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "services", "http+E2U"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrServices(zoneFqdn, 10, 20, ".", "SIPS+D2T"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "services", "SIPS+D2T"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordNaptrResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_record_naptr.test_ttl"
	var v dns.RecordNaptr
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrTtl(zoneFqdn, 10, 10, ".", 10, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrTtl(zoneFqdn, 10, 20, ".", 20, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordNaptrResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_record_naptr.test_use_ttl"
	var v dns.RecordNaptr
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrUseTtl(zoneFqdn, 10, 10, ".", 10, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrUseTtl(zoneFqdn, 10, 10, ".", 10, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordNaptrResource_View(t *testing.T) {
	var resourceName = "nios_dns_record_naptr.test_view"
	var v dns.RecordNaptr
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrView(zoneFqdn, 10, 10, "."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordNaptrExists(ctx context.Context, resourceName string, v *dns.RecordNaptr) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			RecordNaptrAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordNaptr).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordNaptrResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordNaptrResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordNaptrDestroy(ctx context.Context, v *dns.RecordNaptr) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			RecordNaptrAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordNaptr).
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

func testAccCheckRecordNaptrDisappears(ctx context.Context, v *dns.RecordNaptr) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			RecordNaptrAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordNaptrBasicConfig(zoneFqdn string, order, preference int, replacement string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test" {
    name = nios_dns_zone_auth.test.fqdn
    order = %d
    preference = %d
    replacement = %q
}
`, order, preference, replacement)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrComment(zoneFqdn string, order, preference int, replacement string, comment string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_comment" {
    name = nios_dns_zone_auth.test.fqdn
    order = %d
    preference = %d
    replacement = %q
    comment = %q
}
`, order, preference, replacement, comment)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrCreator(zoneFqdn string, order, preference int, replacement string, creator string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_creator" {
    name = nios_dns_zone_auth.test.fqdn
    order = %d
    preference = %d
    replacement = %q
    creator = %q
}
`, order, preference, replacement, creator)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrDdnsPrincipal(zoneFqdn string, order, preference int, replacement, ddnsPrincipal, creator string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_ddns_principal" {
    name = nios_dns_zone_auth.test.fqdn
    order = %d
    preference = %d
    replacement = %q
    ddns_principal = %q
	creator = %q
}
`, order, preference, replacement, ddnsPrincipal, creator)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrDdnsProtected(zoneFqdn string, order, preference int, replacement, ddnsProtected string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_ddns_protected" {
    name = nios_dns_zone_auth.test.fqdn
    order = %d
    preference = %d
    replacement = %q
    ddns_protected = %q
}
`, order, preference, replacement, ddnsProtected)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrDisable(zoneFqdn string, order, preference int, replacement, disable string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_disable" {
    name = nios_dns_zone_auth.test.fqdn
    order = %d
    preference = %d
    replacement = %q
    disable = %q
}
`, order, preference, replacement, disable)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrExtAttrs(zoneFqdn string, order, preference int, replacement string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_extattrs" {
    name = nios_dns_zone_auth.test.fqdn
    order = %d
    preference = %d
    replacement = %q
    extattrs = %s
}
`, order, preference, replacement, extattrsStr)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrFlags(zoneFqdn string, order, preference int, replacement string, flags string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_flags" {
    name = nios_dns_zone_auth.test.fqdn
    order = %d
    preference = %d
    replacement = %q
    flags = %q
}
`, order, preference, replacement, flags)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrForbidReclamation(zoneFqdn string, order, preference int, replacement string, forbidReclamation string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_forbid_reclamation" {
    name = nios_dns_zone_auth.test.fqdn
    order = %d
    preference = %d
    replacement = %q
    forbid_reclamation = %q
}
`, order, preference, replacement, forbidReclamation)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrName(zoneFqdn1, zoneFqdn2 string, order int, preference int, replacement string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_name" {
    name = nios_dns_zone_auth.test.fqdn
    order = %d
    preference = %d
    replacement = %q
}
`, order, preference, replacement)
	return strings.Join([]string{testAccBaseWithTwoZones(zoneFqdn1, zoneFqdn2), config}, "")
}

func testAccRecordNaptrNameUpdate(zoneFqdn1, zoneFqdn2 string, order int, preference int, replacement string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_name" {
    name = nios_dns_zone_auth.updated_zone.fqdn
    order = %d
    preference = %d
    replacement = %q
}
`, order, preference, replacement)
	return strings.Join([]string{testAccBaseWithTwoZones(zoneFqdn1, zoneFqdn2), config}, "")
}

func testAccRecordNaptrOrder(zoneFqdn string, order, preference int, replacement string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_order" {
    name = nios_dns_zone_auth.test.fqdn
    order = %d
    preference = %d
    replacement = %q
}
`, order, preference, replacement)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrPreference(zoneFqdn string, order, preference int, replacement string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_preference" {
    name = nios_dns_zone_auth.test.fqdn
    order = %d
    preference = %d
    replacement = %q
}
`, order, preference, replacement)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrRegexp(zoneFqdn string, order, preference int, replacement, regexp string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_regexp" {
    name = nios_dns_zone_auth.test.fqdn
    order = %d
    preference = %d
    replacement = %q
    regexp = %q
}
`, order, preference, replacement, regexp)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrReplacement(zoneFqdn string, order, preference int, replacement string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_replacement" {
    name = nios_dns_zone_auth.test.fqdn
    order = %d
    preference = %d
    replacement = %q
}
`, order, preference, replacement)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrServices(zoneFqdn string, order, preference int, replacement, services string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_services" {
    name = nios_dns_zone_auth.test.fqdn
    order = %d
    preference = %d
    replacement = %q
    services = %q
}
`, order, preference, replacement, services)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrTtl(zoneFqdn string, order, preference int, replacement string, ttl int, useTtl string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_ttl" {
    name = nios_dns_zone_auth.test.fqdn
    order = %d
    preference = %d
    replacement = %q
    ttl = %d
    use_ttl = %q
}
`, order, preference, replacement, ttl, useTtl)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrUseTtl(zoneFqdn string, order, preference int, replacement string, ttl int, useTtl string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_use_ttl" {
    name = nios_dns_zone_auth.test.fqdn
    order = %d
    preference = %d
    replacement = %q
    ttl = %d
    use_ttl = %q
}
`, order, preference, replacement, ttl, useTtl)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrView(zoneFqdn string, order, preference int, replacement string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_view" {
    name = nios_dns_zone_auth.test.fqdn
    order = %d
    preference = %d
    replacement = %q
}
`, order, preference, replacement)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccBaseWithZone(zoneFqdn string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test" {
    fqdn = %q
}
`, zoneFqdn)
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
