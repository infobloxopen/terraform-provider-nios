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
// - IPV4 reverse mapping zone - 92.168.10.0/24
// - IPV6 reverse mapping zone - 2001::/64
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
					resource.TestCheckResourceAttr(resourceName, "name", "22.10.168.192.in-addr.arpa"),
					// Test fields with default value
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
				Config: testAccRecordPtrDdnsPrincipal("25.10.168.192.in-addr.arpa", "ptr.example.com", "default", "DYNAMIC", "host/myhost.example.com@EXAMPLE.COM"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "host/myhost.example.com@EXAMPLE.COM"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrDdnsPrincipal("25.10.168.192.in-addr.arpa", "ptr.example.com", "default", "DYNAMIC", "host/otherhost.example.net@EXAMPLE.NET"),
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
				Config: testAccRecordPtrDdnsProtected("26.10.168.192.in-addr.arpa", "ptr.example.com", "default", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrDdnsProtected("26.10.168.192.in-addr.arpa", "ptr.example.com", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "true"),
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
				Config: testAccRecordPtrDisable("27.10.168.192.in-addr.arpa", "ptr.example.com", "default", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrDisable("27.10.168.192.in-addr.arpa", "ptr.example.com", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordPtrResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_record_ptr.test_extattrs"
	var v dns.RecordPtr
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordPtrExtAttrs("28.10.168.192.in-addr.arpa", "ptr.example.com", "default", map[string]string{"Site": extAttrValue1}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrExtAttrs("28.10.168.192.in-addr.arpa", "ptr.example.com", "default", map[string]string{"Site": extAttrValue2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
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
				Config: testAccRecordPtrForbidReclamation("29.10.168.192.in-addr.arpa", "ptr.example.com", "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrForbidReclamation("29.10.168.192.in-addr.arpa", "ptr.example.com", "default", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "false"),
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
				Config: testAccRecordPtrIpv4addr("192.168.10.30", "ptr.example.com", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "192.168.10.30"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrIpv4addr("192.168.10.31", "ptr.example.com", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "192.168.10.31"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// TestAccRecordPtrResource_FuncCallIpv4Addr tests the "func_call" attribute functionality
// which allocates IP addresses using next_available_ip. Since func_call attribute can't be
// updated, the comment is updated to demonstrate an update to the resource
func TestAccRecordPtrResource_FuncCallIpv4Addr(t *testing.T) {
	var resourceName = "nios_dns_record_ptr.test_func_call"
	var v dns.RecordPtr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordPtrFuncCallIpv4Addr("192.168.10.0/24", "ptr.example.com", "default", "Created with func_call"),
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
				Config: testAccRecordPtrFuncCallIpv4Addr("192.168.10.0/24", "ptr2.example.com", "default", "Updated with func_call"),
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
				Config: testAccRecordPtrName("32.10.168.192.in-addr.arpa", "ptr.example.com", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "32.10.168.192.in-addr.arpa"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrName("33.10.168.192.in-addr.arpa", "ptr.example.com", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "33.10.168.192.in-addr.arpa"),
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
				Config: testAccRecordPtrPtrdname("ptr.example.com", "192.168.10.34", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ptrdname", "ptr.example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrPtrdname("updated.example.com", "192.168.10.34", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ptrdname", "updated.example.com"),
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
				Config: testAccRecordPtrTtl("2001::26", "ptr.example.com", "default", 300, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "300"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrTtl("2001::26", "ptr.example.com", "default", 600, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "600"),
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
				Config: testAccRecordPtrUseTtl("2001::27", "ptr.example.com", "default", "true", 300),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPtrUseTtl("2001::27", "ptr.example.com", "default", "false", 300),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordPtrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
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

func testAccRecordPtrDdnsPrincipal(Name, Ptrdname, View, Creator, DdnsPrincipal string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test_ddns_principal" {
	name = %q
	ptrdname = %q
	view = %q
	creator = %q
    ddns_principal = %q
}
`, Name, Ptrdname, View, Creator, DdnsPrincipal)
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

func testAccRecordPtrExtAttrs(Name, Ptrdname, View string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test_extattrs" {
	name = %q
	ptrdname = %q
	view = %q
    extattrs = %s
}
`, Name, Ptrdname, View, extattrsStr)
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

func testAccRecordPtrFuncCallIpv4Addr(network, ptrdname, view, comment string) string {
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

func testAccRecordPtrName(name, ptrdname, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test_name" {
    name = %q
	ptrdname = %q
	view = %q
}
`, name, ptrdname, view)
}

func testAccRecordPtrPtrdname(ptrdname, ipv4addr, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test_ptrdname" {
    ptrdname = %q
	ipv4addr = %q
	view = %q
}
`, ptrdname, ipv4addr, view)
}

func testAccRecordPtrTtl(ipv6addr, ptrdname, view string, ttl int32, useTtl string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test_ttl" {
	ipv6addr = %q
	ptrdname = %q
	view = %q
    ttl = %d
	use_ttl = %q
}
`, ipv6addr, ptrdname, view, ttl, useTtl)
}

func testAccRecordPtrUseTtl(ipv6addr, ptrdname, view, useTtl string, ttl int32) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ptr" "test_use_ttl" {
	ipv6addr = %q
	ptrdname = %q
	view = %q
    use_ttl = %q
	ttl = %d
}
`, ipv6addr, ptrdname, view, useTtl, ttl)
}
