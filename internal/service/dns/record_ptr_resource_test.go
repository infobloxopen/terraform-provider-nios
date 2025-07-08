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

// TODO : Add readable attributes for the resource
var readableAttributesForRecordPtr = "aws_rte53_record_info,cloud_info,comment,creation_time,creator,ddns_principal,ddns_protected,disable,discovered_data,dns_name,dns_ptrdname,extattrs,forbid_reclamation,ipv4addr,ipv6addr,last_queried,ms_ad_user_data,name,ptrdname,reclaimable,shared_record_group,ttl,use_ttl,view,zone"

func TestAccRecordPtrResource_basic(t *testing.T) {
	var resourceName = "nios_dns_record_ptr.test"
	var v dns.RecordPtr

	Ipv4addr := "192.168.10.22"
	Ptrdname := "ptr.example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordPtrBasicConfig(Ipv4addr, Ptrdname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", Ipv4addr),
					resource.TestCheckResourceAttr(resourceName, "ptrdname", Ptrdname),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
					resource.TestCheckResourceAttr(resourceName, "name", "23.10.168.192.in-addr.arpa"),
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "false"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
					resource.TestCheckResourceAttr(resourceName, "zone", "10.168.192.in-addr.arpa"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordPtrResource_disappears(t *testing.T) {
	resourceName := "nios_dns_record_ptr.test"
	var v dns.RecordPtr

	Ipv4addr := "192.168.10.22"
	Ptrdname := "ptr.example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordPtrDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordPtrBasicConfig(Ipv4addr, Ptrdname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					testAccCheckRecordPtrDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordPtrResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_record_ptr.test_comment"
	var v dns.RecordPtr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordPtrComment("23.10.168.192.in-addr.arpa", "ptr.example.com", "default", "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrComment("23.10.168.192.in-addr.arpa", "ptr.example.com", "default", "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordPtrResource_Creator(t *testing.T) {
	var resourceName = "nios_dns_record_ptr.test_creator"
	var v dns.RecordPtr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordPtrCreator("24.10.168.192.in-addr.arpa", "ptr.example.com", "default", "STATIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrCreator("24.10.168.192.in-addr.arpa", "ptr.example.com", "default", "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "DYNAMIC"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordPtrResource_DdnsPrincipal(t *testing.T) {
	var resourceName = "nios_dns_record_ptr.test_ddns_principal"
	var v dns.RecordPtr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordPtrDdnsPrincipal("25.10.168.192.in-addr.arpa", "ptr.example.com", "default", "host/myhost.example.com@EXAMPLE.COM"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "host/myhost.example.com@EXAMPLE.COM"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrDdnsPrincipal("25.10.168.192.in-addr.arpa", "ptr.example.com", "default", "host/otherhost.example.net@EXAMPLE.NET"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "host/otherhost.example.net@EXAMPLE.NET"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordPtrResource_DdnsProtected(t *testing.T) {
	var resourceName = "nios_dns_record_ptr.test_ddns_protected"
	var v dns.RecordPtr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordPtrDdnsProtected("DDNS_PROTECTED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "DDNS_PROTECTED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrDdnsProtected("DDNS_PROTECTED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "DDNS_PROTECTED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordPtrResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_record_ptr.test_disable"
	var v dns.RecordPtr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordPtrDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordPtrResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_record_ptr.test_extattrs"
	var v dns.RecordPtr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordPtrExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordPtrResource_ForbidReclamation(t *testing.T) {
	var resourceName = "nios_dns_record_ptr.test_forbid_reclamation"
	var v dns.RecordPtr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordPtrForbidReclamation("FORBID_RECLAMATION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "FORBID_RECLAMATION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrForbidReclamation("FORBID_RECLAMATION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "FORBID_RECLAMATION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordPtrResource_Ipv4addr(t *testing.T) {
	var resourceName = "nios_dns_record_ptr.test_ipv4addr"
	var v dns.RecordPtr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordPtrIpv4addr("IPV4ADDR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "IPV4ADDR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrIpv4addr("IPV4ADDR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "IPV4ADDR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordPtrResource_FuncCall(t *testing.T) {
	var resourceName = "nios_dns_record_ptr.test_func_call"
	var v dns.RecordPtr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordPtrFuncCall("192.168.10.0/24", "ptr.example.com", "default", "Created with func_call"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "192.168.10.1"),
					resource.TestCheckResourceAttr(resourceName, "ptrdname", "ptr.example.com"),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
					resource.TestCheckResourceAttr(resourceName, "comment", "Created with func_call"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrFuncCall("192.168.10.0/24", "ptr2.example.com", "default", "Updated with func_call"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "192.168.10.1"),
					resource.TestCheckResourceAttr(resourceName, "ptrdname", "ptr2.example.com"),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated with func_call"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordPtrResource_Ipv6addr(t *testing.T) {
	var resourceName = "nios_dns_record_ptr.test_ipv6addr"
	var v dns.RecordPtr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordPtrIpv6addr("2001::24", "test.example.com", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", "2001::24"),
					resource.TestCheckResourceAttr(resourceName, "ptrdname", "test.example.com"),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrIpv6addr("2001::25", "test.example.com", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", "2001::25"),
					resource.TestCheckResourceAttr(resourceName, "ptrdname", "test.example.com"),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordPtrResource_Name(t *testing.T) {
	var resourceName = "nios_dns_record_ptr.test_name"
	var v dns.RecordPtr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordPtrName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordPtrResource_Ptrdname(t *testing.T) {
	var resourceName = "nios_dns_record_ptr.test_ptrdname"
	var v dns.RecordPtr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordPtrPtrdname("PTRDNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ptrdname", "PTRDNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrPtrdname("PTRDNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ptrdname", "PTRDNAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordPtrResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_record_ptr.test_ttl"
	var v dns.RecordPtr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordPtrTtl("TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrTtl("TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordPtrResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_record_ptr.test_use_ttl"
	var v dns.RecordPtr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordPtrUseTtl("USE_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "USE_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrUseTtl("USE_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "USE_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordPtrResource_View(t *testing.T) {
	var resourceName = "nios_dns_record_ptr.test_view"
	var v dns.RecordPtr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordPtrView("VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrView("VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordPtrExists(ctx context.Context, resourceName string, v *dns.RecordPtr) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			RecordPtrAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordPtr).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordPtrResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordPtrResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordPtrDestroy(ctx context.Context, v *dns.RecordPtr) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			RecordPtrAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordPtr).
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

func testAccCheckRecordPtrDisappears(ctx context.Context, v *dns.RecordPtr) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			RecordPtrAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordPtrBasicConfig(Ipv4addr, Ptrdname string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test" {
	ipv4addr = %q
	ptrdname = %q
}
`, Ipv4addr, Ptrdname)
}

func testAccRecordPtrComment(Name, Ptrdname, View, Comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test_comment" {
	name = %q
	ptrdname = %q
	view = %q
    comment = %q
}
`, Name, Ptrdname, View, Comment)
}

func testAccRecordPtrCreator(Name, Ptrdname, View, Creator string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test_creator" {
	name = %q
	ptrdname = %q
	view = %q
    creator = %q
}
`, Name, Ptrdname, View, Creator)
}

func testAccRecordPtrDdnsPrincipal(Name, Ptrdname, View, DdnsPrincipal string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test_ddns_principal" {
	name = %q
	ptrdname = %q
	view = %q
    ddns_principal = %q
}
`, Name, Ptrdname, View, DdnsPrincipal)
}

func testAccRecordPtrDdnsProtected(Name, Ptrdname, View, DdnsProtected string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test_ddns_protected" {
	name = %q
	ptrdname = %q
	view = %q
    ddns_protected = %q
}
`, Name, Ptrdname, View, DdnsProtected)
}

func testAccRecordPtrDisable(Name, Ptrdname, View, disable string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test_disable" {
	name = %q
	ptrdname = %q
	view = %q
    disable = %q
}
`, Name, Ptrdname, View, disable)
}

func testAccRecordPtrExtAttrs(Name, Ptrdname, View, extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test_extattrs" {
	name = %q
	ptrdname = %q
	view = %q
    extattrs = %q
}
`, Name, Ptrdname, View, extAttrs)
}

func testAccRecordPtrForbidReclamation(Name, Ptrdname, View, forbidReclamation string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test_forbid_reclamation" {
	name = %q
	ptrdname = %q
	view = %q
    forbid_reclamation = %q
}
`, Name, Ptrdname, View, forbidReclamation)
}

func testAccRecordPtrIpv4addr(ipv4addr, Ptrdname, View string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test_ipv4addr" {
    ipv4addr = %q
	ptrdname = %q
	view = %q
}
`, ipv4addr, Ptrdname, View)
}

func testAccRecordPtrFuncCall(network, ptrdname, view, comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test_func_call" {
	func_call = {
		"attribute_name" = "ipv4addr"
		"object_function" = "next_available_ip"
		"result_field" = "ips"
		"object" = "network"
		"object_parameters" = {
			"network" = %q
			"network_view" = "default"
	}
	ptrdname = %q
	view = %q
	comment = %q
}
`, network, ptrdname, view, comment)
}

func testAccRecordPtrIpv6addr(ipv6addr, ptrdname, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test_ipv6addr" {
    ipv6addr = %q
	ptrdname = %q
	view = %q
}
`, ipv6addr, ptrdname, view)
}

func testAccRecordPtrName(name string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test_name" {
    name = %q
}
`, name)
}

func testAccRecordPtrPtrdname(ptrdname string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test_ptrdname" {
    ptrdname = %q
}
`, ptrdname)
}

func testAccRecordPtrTtl(ttl string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test_ttl" {
    ttl = %q
}
`, ttl)
}

func testAccRecordPtrUseTtl(useTtl string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test_use_ttl" {
    use_ttl = %q
}
`, useTtl)
}

func testAccRecordPtrView(view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test_view" {
    view = %q
}
`, view)
}
