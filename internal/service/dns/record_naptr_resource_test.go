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
	name := acctest.RandomNameWithPrefix("test-naptr")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrBasicConfig(zoneFqdn, name, 10, 10, "."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("%s.%s", name, zoneFqdn)),
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
	name := acctest.RandomNameWithPrefix("test-naptr")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordNaptrDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordNaptrBasicConfig(zoneFqdn, name, 10, 10, "."),
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
	name := acctest.RandomNameWithPrefix("test-naptr")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrComment(zoneFqdn, name, 10, 10, ".", "comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "comment"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrComment(zoneFqdn, name, 10, 10, ".", "updated comment"),
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
	name := acctest.RandomNameWithPrefix("test-naptr")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrCreator(zoneFqdn, name, 10, 10, ".", "STATIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrCreator(zoneFqdn, name, 10, 10, ".", "DYNAMIC"),
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
	name := acctest.RandomNameWithPrefix("test-naptr")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrDdnsPrincipal(zoneFqdn, name, 10, 10, ".", "ddns_principal", "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "ddns_principal"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrDdnsPrincipal(zoneFqdn, name, 10, 10, ".", "updated_ddns_principal", "DYNAMIC"),
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
	name := acctest.RandomNameWithPrefix("test-naptr")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrDdnsProtected(zoneFqdn, name, 10, 10, ".", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrDdnsProtected(zoneFqdn, name, 10, 10, ".", "true"),
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
	name := acctest.RandomNameWithPrefix("test-naptr")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrDisable(zoneFqdn, name, 10, 10, ".", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrDisable(zoneFqdn, name, 10, 10, ".", "true"),
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
	name := acctest.RandomNameWithPrefix("test-naptr")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrExtAttrs(zoneFqdn, name, 10, 10, ".", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrExtAttrs(zoneFqdn, name, 10, 10, ".", map[string]string{
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
	name := acctest.RandomNameWithPrefix("test-naptr")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read with "U" flag
			{
				Config: testAccRecordNaptrFlags(zoneFqdn, name, 10, 10, ".", "U"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "flags", "U"),
				),
			},
			// Update to "S" flag
			{
				Config: testAccRecordNaptrFlags(zoneFqdn, name, 10, 10, ".", "S"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "flags", "S"),
				),
			},
			// Update to "A" flag
			{
				Config: testAccRecordNaptrFlags(zoneFqdn, name, 10, 10, ".", "A"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "flags", "A"),
				),
			},
			// Update to "P" flag
			{
				Config: testAccRecordNaptrFlags(zoneFqdn, name, 10, 10, ".", "P"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "flags", "P"),
				),
			},
			// Update to empty flag
			{
				Config: testAccRecordNaptrFlags(zoneFqdn, name, 10, 10, ".", ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "flags", ""),
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
	name := acctest.RandomNameWithPrefix("test-naptr")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrForbidReclamation(zoneFqdn, name, 10, 10, ".", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrForbidReclamation(zoneFqdn, name, 10, 10, ".", "true"),
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
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name1 := acctest.RandomNameWithPrefix("test-naptr")
	name2 := acctest.RandomNameWithPrefix("test-naptr")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrName(zoneFqdn, name1, 10, 10, "."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("%s.%s", name1, zoneFqdn)),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrName(zoneFqdn, name2, 10, 10, "."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("%s.%s", name2, zoneFqdn)),
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
	name := acctest.RandomNameWithPrefix("test-naptr")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrOrder(zoneFqdn, name, 10, 10, "."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "order", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrOrder(zoneFqdn, name, 20, 10, "."),
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
	name := acctest.RandomNameWithPrefix("test-naptr")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrPreference(zoneFqdn, name, 10, 10, "."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preference", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrPreference(zoneFqdn, name, 10, 20, "."),
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
	name := acctest.RandomNameWithPrefix("test-naptr")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrRegexp(zoneFqdn, name, 10, 10, ".", ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "regexp", ""),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrRegexp(zoneFqdn, name, 10, 10, ".", "!^.*$!sip:jdoe@corpxyz.com!"),
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
	name := acctest.RandomNameWithPrefix("test-naptr")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrReplacement(zoneFqdn, name, 10, 10, "."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "replacement", "."),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrReplacement(zoneFqdn, name, 10, 20, "test.com"),
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
	name := acctest.RandomNameWithPrefix("test-naptr")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrServices(zoneFqdn, name, 10, 10, ".", "http+E2U"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "services", "http+E2U"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrServices(zoneFqdn, name, 10, 20, ".", "SIPS+D2T"),
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
	name := acctest.RandomNameWithPrefix("test-naptr")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrTtl(zoneFqdn, name, 10, 10, ".", 10, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrTtl(zoneFqdn, name, 10, 20, ".", 20, "true"),
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
	name := acctest.RandomNameWithPrefix("test-naptr")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNaptrUseTtl(zoneFqdn, name, 10, 10, ".", 10, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNaptrUseTtl(zoneFqdn, name, 10, 10, ".", 10, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
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

func testAccRecordNaptrBasicConfig(zoneFqdn, name string, order, preference int, replacement string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test" {
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    order = %d
    preference = %d
    replacement = %q
}
`, name, order, preference, replacement)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrComment(zoneFqdn, name string, order, preference int, replacement string, comment string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_comment" {
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    order = %d
    preference = %d
    replacement = %q
    comment = %q
}
`, name, order, preference, replacement, comment)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrCreator(zoneFqdn, name string, order, preference int, replacement string, creator string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_creator" {
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    order = %d
    preference = %d
    replacement = %q
    creator = %q
}
`, name, order, preference, replacement, creator)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrDdnsPrincipal(zoneFqdn, name string, order, preference int, replacement, ddnsPrincipal, creator string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_ddns_principal" {
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    order = %d
    preference = %d
    replacement = %q
    ddns_principal = %q
	creator = %q
}
`, name, order, preference, replacement, ddnsPrincipal, creator)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrDdnsProtected(zoneFqdn, name string, order, preference int, replacement, ddnsProtected string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_ddns_protected" {
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    order = %d
    preference = %d
    replacement = %q
    ddns_protected = %q
}
`, name, order, preference, replacement, ddnsProtected)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrDisable(zoneFqdn, name string, order, preference int, replacement, disable string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_disable" {
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    order = %d
    preference = %d
    replacement = %q
    disable = %q
}
`, name, order, preference, replacement, disable)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrExtAttrs(zoneFqdn, name string, order, preference int, replacement string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_extattrs" {
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    order = %d
    preference = %d
    replacement = %q
    extattrs = %s
}
`, name, order, preference, replacement, extattrsStr)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrFlags(zoneFqdn, name string, order, preference int, replacement string, flags string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_flags" {
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    order = %d
    preference = %d
    replacement = %q
    flags = %q
}
`, name, order, preference, replacement, flags)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrForbidReclamation(zoneFqdn, name string, order, preference int, replacement string, forbidReclamation string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_forbid_reclamation" {
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    order = %d
    preference = %d
    replacement = %q
    forbid_reclamation = %q
}
`, name, order, preference, replacement, forbidReclamation)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrName(zoneFqdn, name string, order int, preference int, replacement string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_name" {
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    order = %d
    preference = %d
    replacement = %q
}
`, name, order, preference, replacement)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrOrder(zoneFqdn, name string, order, preference int, replacement string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_order" {
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    order = %d
    preference = %d
    replacement = %q
}
`, name, order, preference, replacement)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrPreference(zoneFqdn, name string, order, preference int, replacement string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_preference" {
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    order = %d
    preference = %d
    replacement = %q
}
`, name, order, preference, replacement)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrRegexp(zoneFqdn, name string, order, preference int, replacement, regexp string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_regexp" {
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    order = %d
    preference = %d
    replacement = %q
    regexp = %q
}
`, name, order, preference, replacement, regexp)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrReplacement(zoneFqdn, name string, order, preference int, replacement string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_replacement" {
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    order = %d
    preference = %d
    replacement = %q
}
`, name, order, preference, replacement)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrServices(zoneFqdn, name string, order, preference int, replacement, services string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_services" {
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    order = %d
    preference = %d
    replacement = %q
    services = %q
}
`, name, order, preference, replacement, services)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrTtl(zoneFqdn, name string, order, preference int, replacement string, ttl int, useTtl string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_ttl" {
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    order = %d
    preference = %d
    replacement = %q
    ttl = %d
    use_ttl = %q
}
`, name, order, preference, replacement, ttl, useTtl)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNaptrUseTtl(zoneFqdn, name string, order, preference int, replacement string, ttl int, useTtl string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_naptr" "test_use_ttl" {
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    order = %d
    preference = %d
    replacement = %q
    ttl = %d
    use_ttl = %q
}
`, name, order, preference, replacement, ttl, useTtl)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccBaseWithZone(zoneFqdn string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test" {
    fqdn = %q
}
`, zoneFqdn)
}
