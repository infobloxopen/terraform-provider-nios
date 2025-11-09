package dns_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"testing"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// TODO: OBJECTS TO BE PRESENT IN GRID FOR TESTS
// - Parent Zone: example.com (in default view)
// - IPv4 Network: 85.85.0.0/16 (for func call tests)
func TestAccRecordAResource_basic(t *testing.T) {
	var resourceName = "nios_dns_record_a.test"
	var v dns.RecordA
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordABasicConfig(name, "10.0.0.20", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "10.0.0.20"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
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

func TestAccRecordAResource_disappears(t *testing.T) {
	resourceName := "nios_dns_record_a.test"
	var v dns.RecordA
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordADestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordABasicConfig(name, "10.0.0.20", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					testAccCheckRecordADisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordAResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_record_a.test_comment"
	var v dns.RecordA
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAComment(name, "10.0.0.20", "default", "This is a new record"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a new record"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAComment(name, "10.0.0.20", "default", "This is an updated record"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated record"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAResource_Creator(t *testing.T) {
	var resourceName = "nios_dns_record_a.test_creator"
	var v dns.RecordA
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordACreator(name, "10.0.0.20", "default", "STATIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordACreator(name, "10.0.0.20", "default", "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "DYNAMIC"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAResource_DdnsPrincipal(t *testing.T) {
	var resourceName = "nios_dns_record_a.test_ddns_principal"
	var v dns.RecordA
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordADdnsPrincipal(name, "10.0.0.20", "default", "DDNS_PRINCIPAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "DDNS_PRINCIPAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordADdnsPrincipal(name, "10.0.0.20", "default", "DDNS_PRINCIPAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "DDNS_PRINCIPAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAResource_DdnsProtected(t *testing.T) {
	var resourceName = "nios_dns_record_a.test_ddns_protected"
	var v dns.RecordA
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordADdnsProtected(name, "10.0.0.20", "default", "false"),
				Check: resource.ComposeTestCheckFunc(
					//testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordADdnsProtected(name, "10.0.0.20", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_record_a.test_disable"
	var v dns.RecordA
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordADisable(name, "10.0.0.20", "default", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordADisable(name, "10.0.0.20", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAResource_Extattrs(t *testing.T) {
	var resourceName = "nios_dns_record_a.test_extattrs"
	var v dns.RecordA
	name := acctest.RandomName() + ".example.com"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAExtattrs(name, "10.0.0.20", "default", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAExtattrs(name, "10.0.0.20", "default", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAResource_ForbidReclamation(t *testing.T) {
	var resourceName = "nios_dns_record_a.test_forbid_reclamation"
	var v dns.RecordA
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAForbidReclamation(name, "10.0.0.20", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAForbidReclamation(name, "10.0.0.20", "default", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// TestAccRecordAResource_FuncCall tests the "func_call" attribute functionality
// which allocates IPv4 addresses using next_available_ip. Since func_call attribute can't be
// updated, the comment is updated to demonstrate an update to the resource
func TestAccRecordAResource_FuncCall(t *testing.T) {
	var resourceName = "nios_dns_record_a.test_func_call"
	var v dns.RecordA
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAFuncCall(name, "default", "ipv4addr", "next_available_ip", "", "ips", "network", "85.85.0.0/16", "Original Function Call"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAFuncCall(name, "default", "comment", "next_available_ip", "", "ips", "network", "85.85.0.0/16", "Function Call with Update"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
				),
			},
		},
	})
}

func TestAccRecordAResource_Ipv4addr(t *testing.T) {
	var resourceName = "nios_dns_record_a.test_ipv4addr"
	var v dns.RecordA
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAIpv4addr(name, "10.0.0.20", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "10.0.0.20"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAIpv4addr(name, "10.1.0.20", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "10.1.0.20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAResource_Name(t *testing.T) {
	var resourceName = "nios_dns_record_a.test_name"
	var v dns.RecordA
	name1 := acctest.RandomName() + ".example.com"
	name2 := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAName(name1, "10.0.0.20", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAName(name2, "10.0.0.20", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_record_a.test_ttl"
	var v dns.RecordA
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordATtl(name, "10.0.0.20", "default", 10, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordATtl(name, "10.0.0.20", "default", 0, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_record_a.test_use_ttl"
	var v dns.RecordA
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAUseTtl(name, "10.0.0.20", "default", "true", 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAUseTtl(name, "10.0.0.20", "default", "false", 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordAExists(ctx context.Context, resourceName string, v *dns.RecordA) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	var readableAttributes = "aws_rte53_record_info,cloud_info,comment,creation_time,creator,ddns_principal,ddns_protected,disable,discovered_data,dns_name,extattrs,forbid_reclamation,ipv4addr,last_queried,ms_ad_user_data,name,reclaimable,shared_record_group,ttl,use_ttl,view,zone"
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		state.RootModule().Resources[resourceName].Primary.ID = utils.ExtractResourceRef(rs.Primary.Attributes["ref"])
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			RecordAAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributes).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordAResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordAResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordADestroy(ctx context.Context, v *dns.RecordA) resource.TestCheckFunc {
	// Verify the resource was destroyed
	var readableAttributes = "aws_rte53_record_info,cloud_info,comment,creation_time,creator,ddns_principal,ddns_protected,disable,discovered_data,dns_name,extattrs,forbid_reclamation,ipv4addr,last_queried,ms_ad_user_data,name,reclaimable,shared_record_group,ttl,use_ttl,view,zone"
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			RecordAAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributes).
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

func testAccCheckRecordADisappears(ctx context.Context, v *dns.RecordA) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			RecordAAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordABasicConfig(name, ipV4Addr, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_a" "test" {
	name = %q
	ipv4addr = %q
	view = %q
}
`, name, ipV4Addr, view)
}

func testAccRecordAComment(name, ipV4Addr, view, comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_a" "test_comment" {
	name = %q
	ipv4addr = %q
	view = %q
	comment = %q
}
`, name, ipV4Addr, view, comment)
}

func testAccRecordACreator(name, ipV4Addr, view, creator string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_a" "test_creator" {
	name = %q
	ipv4addr = %q
	view = %q  
	creator = %q
}
`, name, ipV4Addr, view, creator)
}

func testAccRecordADdnsPrincipal(name, ipV4Addr, view, ddnsPrincipal string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_a" "test_ddns_principal" {
	name = %q
	ipv4addr = %q
	view = %q
	creator = "DYNAMIC"
	ddns_principal = %q
}
`, name, ipV4Addr, view, ddnsPrincipal)
}

func testAccRecordADdnsProtected(name, ipV4Addr, view, ddnsProtected string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_a" "test_ddns_protected" {
    name = %q
	ipv4addr = %q
	view = %q
	ddns_protected = %q
}
`, name, ipV4Addr, view, ddnsProtected)
}

func testAccRecordADisable(name, ipV4Addr, view, disable string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_a" "test_disable" {
    name = %q
	ipv4addr = %q
	view = %q
	disable = %q
}
`, name, ipV4Addr, view, disable)
}

func testAccRecordAExtattrs(name, ipV4Addr, view string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_dns_record_a" "test_extattrs" {
    name = %q
 	ipv4addr = %q
 	view = %q
 	extattrs = %s
}
`, name, ipV4Addr, view, extattrsStr)
}

func testAccRecordAForbidReclamation(name, ipV4Addr, view, forbidReclamation string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_a" "test_forbid_reclamation" {
    name = %q
	ipv4addr = %q
	view = %q
	forbid_reclamation = %q
}
`, name, ipV4Addr, view, forbidReclamation)
}

func testAccRecordAFuncCall(name, view, attributeName, objFunc, parameters, resultField, object, objectParameters, comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_a" "test_func_call" {
	name = %q
	view = %q
	func_call = {
		"attribute_name" = %q
		"object_function" = %q
		"result_field" = %q
		"object" = %q
		"object_parameters" = {
			"network" = %q
			"network_view" = "default"
		}
	}
	comment = %q
}
`, name, view, attributeName, objFunc, resultField, object, objectParameters, comment)
}

func testAccRecordAIpv4addr(name, ipV4addr, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_a" "test_ipv4addr" {
	name = %q
	ipv4addr = %q
	view = %q
}
`, name, ipV4addr, view)
}

func testAccRecordAName(name, ipV4addr, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_a" "test_name" {
	name = %q
	ipv4addr = %q
	view = %q
}
`, name, ipV4addr, view)
}

func testAccRecordATtl(name, ipV4Addr, view string, ttl int32, use_ttl string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_a" "test_ttl" {
    name = %q
	ipv4addr = %q
	view = %q
	ttl = %d
	use_ttl = %q
}
`, name, ipV4Addr, view, ttl, use_ttl)
}

func testAccRecordAUseTtl(name, ipV4Addr, view, useTtl string, ttl int32) string {
	return fmt.Sprintf(`
resource "nios_dns_record_a" "test_use_ttl" {
    name = %q
	ipv4addr = %q
	view = %q
	use_ttl = %q
	ttl = %d
}
`, name, ipV4Addr, view, useTtl, ttl)
}
