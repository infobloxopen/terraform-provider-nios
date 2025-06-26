package dns_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/acctest"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/utils"
)

func TestAccRecordAaaaResource_basic(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test"
	var v dns.RecordAaaa
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaBasicConfig(name, "2002:1111::1401", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", "2002:1111::1401"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_disappears(t *testing.T) {
	t.Skip("Skipping test for disappears")
	resourceName := "nios_dns_record_aaaa.test"
	var v dns.RecordAaaa
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordAaaaDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordAaaaBasicConfig(name, "2002:1111::1401", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					testAccCheckRecordAaaaDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordAaaaResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_comment"
	var v dns.RecordAaaa
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaComment(name, "2002:1111::1401", "default", "This is a new record"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a new record"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaComment(name, "2002:1111::1401", "default", "This is an updated record"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated record"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_Creator(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_creator"
	var v dns.RecordAaaa
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaCreator(name, "2002:1111::1401", "default", "STATIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaCreator(name, "2002:1111::1401", "default", "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "DYNAMIC"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_DdnsPrincipal(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_ddns_principal"
	var v dns.RecordAaaa
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaDdnsPrincipal(name, "2002:1111::1401", "default", "DYNAMIC", "ddns_principal"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "ddns_principal"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaDdnsPrincipal(name, "2002:1111::1401", "default", "DYNAMIC", "updated_ddns_principal"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "updated_ddns_principal"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_DdnsProtected(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_ddns_protected"
	var v dns.RecordAaaa
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaDdnsProtected(name, "2002:1111::1401", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaDdnsProtected(name, "2002:1111::1401", "default", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_disable"
	var v dns.RecordAaaa
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaDisable(name, "2002:1111::1401", "default", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaDisable(name, "2002:1111::1401", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_extattrs"
	var v dns.RecordAaaa
	name := acctest.RandomName() + ".example.com"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaExtAttrs(name, "2002:1111::1401", "default", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaExtAttrs(name, "2002:1111::1401", "default", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_ForbidReclamation(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_forbid_reclamation"
	var v dns.RecordAaaa
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaForbidReclamation(name, "2002:1111::1401", "default", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaForbidReclamation(name, "2002:1111::1401", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_Ipv6addr(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_ipv6addr"
	var v dns.RecordAaaa
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaIpv6addr(name, "default", "2002:1111::1401"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", "2002:1111::1401"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaIpv6addr(name, "default", "2002:1111::1402"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", "2002:1111::1402"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_FuncCall(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_func_call"
	var v dns.RecordAaaa
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaFuncCall(name, "default", "ipv6addr", "next_available_ip", "ips", "ipv6network", "2001:db8:abcd:12::/64", "Original Function Call"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaFuncCall(name, "default", "ipv6addr", "next_available_ip", "ips", "ipv6network", "2001:db8:abcd:12::/64", "Updated Function Call"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_Name(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_name"
	var v dns.RecordAaaa
	name1 := acctest.RandomName() + ".example.com"
	name2 := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaName(name1, "2002:1111::1401", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaName(name2, "2002:1111::1402", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_ttl"
	var v dns.RecordAaaa
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaTtl(name, "2002:1111::1401", "default", 10, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaTtl(name, "2002:1111::1401", "default", 0, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_use_ttl"
	var v dns.RecordAaaa
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaUseTtl(name, "2002:1111::1401", "default", 10, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaUseTtl(name, "2002:1111::1401", "default", 10, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_View(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_view"
	var v dns.RecordAaaa
	name := acctest.RandomName() + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaView(name, "2002:1111::1401", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaView(name, "2002:1111::1401", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordAaaaExists(ctx context.Context, resourceName string, v *dns.RecordAaaa) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	var readableAttributes = "aws_rte53_record_info,cloud_info,comment,creation_time,creator,ddns_principal,ddns_protected,disable,discovered_data,dns_name,extattrs,forbid_reclamation,ipv6addr,last_queried,ms_ad_user_data,name,reclaimable,shared_record_group,ttl,use_ttl,view,zone"

	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			RecordAaaaAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributes).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordAaaaResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordAaaaResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordAaaaDestroy(ctx context.Context, v *dns.RecordAaaa) resource.TestCheckFunc {

	var readableAttributes = "aws_rte53_record_info,cloud_info,comment,creation_time,creator,ddns_principal,ddns_protected,disable,discovered_data,dns_name,extattrs,forbid_reclamation,ipv6addr,last_queried,ms_ad_user_data,name,reclaimable,shared_record_group,ttl,use_ttl,view,zone"
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			RecordAaaaAPI.
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

func testAccCheckRecordAaaaDisappears(ctx context.Context, v *dns.RecordAaaa) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			RecordAaaaAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordAaaaBasicConfig(name, ipV6Addr, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test" {
    name     = %q
    ipv6addr = %q
    view     = %q
}
`, name, ipV6Addr, view)
}

func testAccRecordAaaaComment(name, ipV6Addr, view, comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_comment" {
	name     = %q
	ipv6addr = %q
	view     = %q
    comment  = %q
}
`, name, ipV6Addr, view, comment)
}

func testAccRecordAaaaCreator(name, ipV6Addr, view, creator string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_creator" {
    name     = %q
    ipv6addr = %q
    view     = %q
    creator  = %q
}
`, name, ipV6Addr, view, creator)
}

func testAccRecordAaaaDdnsPrincipal(name, ipV6Addr, view, creator, ddnsPrincipal string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_ddns_principal" {
    name            = %q
	ipv6addr        = %q
	view            = %q
	creator         = %q
    ddns_principal  = %q
}
`, name, ipV6Addr, view, creator, ddnsPrincipal)
}

func testAccRecordAaaaDdnsProtected(name, ipV6Addr, view, ddnsProtected string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_ddns_protected" {
    name            = %q
	ipv6addr        = %q
	view            = %q
    ddns_protected  = %q
}
`, name, ipV6Addr, view, ddnsProtected)
}

func testAccRecordAaaaDisable(name, ipV6Addr, view, disable string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_disable" {
	name     = %q
	ipv6addr = %q
	view     = %q
    disable  = %q
}
`, name, ipV6Addr, view, disable)
}

func testAccRecordAaaaExtAttrs(name, ipV6Addr, view string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf("    %q = %q\n", k, v)
	}
	extattrsStr += "  }"
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_extattrs" {
    name     = %q
    ipv6addr = %q
    view     = %q
    extattrs = %s
}
`, name, ipV6Addr, view, extattrsStr)
}
func testAccRecordAaaaForbidReclamation(name, ipV6Addr, view, forbidReclamation string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_forbid_reclamation" {
    name              = %q
    ipv6addr          = %q
    view              = %q
    forbid_reclamation = %q
}
`, name, ipV6Addr, view, forbidReclamation)
}

func testAccRecordAaaaIpv6addr(name, view, ipv6addr string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_ipv6addr" {
    name     = %q
    view     = %q
    ipv6addr = %q
}
`, name, view, ipv6addr)
}

func testAccRecordAaaaFuncCall(name, view, attributeName, objFunc, resultField, object, objectParameters, comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_func_call" {
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

func testAccRecordAaaaName(name, ipV6Addr, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_name" {
    name     = %q
    ipv6addr = %q
    view     = %q
}
`, name, ipV6Addr, view)
}

func testAccRecordAaaaTtl(name, ipV6Addr, view string, ttl int32, use_ttl string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_ttl" {
	name     = %q
	ipv6addr = %q
	view     = %q
    ttl      = %d
	use_ttl  = %q
}
`, name, ipV6Addr, view, ttl, use_ttl)
}

func testAccRecordAaaaUseTtl(name, ipV6Addr, view string, ttl int32, useTtl string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_use_ttl" {
    name     = %q
    ipv6addr = %q
    view     = %q
    ttl      = %d
    use_ttl  = %q
}
`, name, ipV6Addr, view, ttl, useTtl)
}

func testAccRecordAaaaView(name, ipV6Addr, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_view" {
	name     = %q
	ipv6addr = %q
	view     = %q
}
`, name, ipV6Addr, view)
}
