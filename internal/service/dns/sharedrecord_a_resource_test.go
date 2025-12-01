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

var readableAttributesForSharedrecordA = "comment,disable,dns_name,extattrs,ipv4addr,name,shared_record_group,ttl,use_ttl"

func TestAccSharedrecordAResource_basic(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_a.test"
	var v dns.SharedrecordA
	name := acctest.RandomNameWithPrefix("sharedrecord-a")
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordABasicConfig(name, "10.0.0.0", sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "10.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "shared_record_group", sharedRecordGroup),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordAResource_disappears(t *testing.T) {
	resourceName := "nios_dns_sharedrecord_a.test"
	var v dns.SharedrecordA
	name := acctest.RandomNameWithPrefix("sharedrecord-a")
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharedrecordADestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharedrecordABasicConfig(name, "10.0.0.0", sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAExists(context.Background(), resourceName, &v),
					testAccCheckSharedrecordADisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSharedrecordAResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_a.test_comment"
	var v dns.SharedrecordA
	name := acctest.RandomNameWithPrefix("sharedrecord-a")
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordAComment(name, "10.0.0.0", sharedRecordGroup, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordAComment(name, "10.0.0.0", sharedRecordGroup, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordAResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_a.test_disable"
	var v dns.SharedrecordA
	name := acctest.RandomNameWithPrefix("sharedrecord-a")
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordADisable(name, "10.0.0.0", sharedRecordGroup, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordADisable(name, "10.0.0.0", sharedRecordGroup, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordAResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_a.test_extattrs"
	var v dns.SharedrecordA
	name := acctest.RandomNameWithPrefix("sharedrecord-a") + ".example.com"
	ipv4addr := "10.0.0.1"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordAExtAttrs(name, ipv4addr, sharedRecordGroup, map[string]any{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordAExtAttrs(name, ipv4addr, sharedRecordGroup, map[string]any{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordAResource_Ipv4addr(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_a.test_ipv4addr"
	var v dns.SharedrecordA
	name := acctest.RandomNameWithPrefix("sharedrecord-a")
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordAIpv4addr(name, "10.0.0.0", sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "10.0.0.0"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordAIpv4addr(name, "10.0.0.1", sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "10.0.0.1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordAResource_Name(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_a.test_name"
	var v dns.SharedrecordA
	name1 := acctest.RandomNameWithPrefix("sharedrecord-a")
	name2 := acctest.RandomNameWithPrefix("updatedsharedrecord-a")
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordAName(name1, "10.0.0.0", sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordAName(name2, "10.0.0.0", sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordAResource_SharedRecordGroup(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_a.test_shared_record_group"
	var v dns.SharedrecordA
	name := acctest.RandomNameWithPrefix("sharedrecord-a")
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordASharedRecordGroup(name, "10.0.0.0", sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "shared_record_group", sharedRecordGroup),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordAResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_a.test_ttl"
	var v dns.SharedrecordA
	name := acctest.RandomNameWithPrefix("sharedrecord-a")
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordATtl(name, "10.0.0.0", sharedRecordGroup, 3600, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "3600"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordATtl(name, "10.0.0.0", sharedRecordGroup, 7200, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "7200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordAResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_a.test_use_ttl"
	var v dns.SharedrecordA
	name := acctest.RandomNameWithPrefix("sharedrecord-a")
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordAUseTtl(name, "10.0.0.0", sharedRecordGroup, 3600, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordAUseTtl(name, "10.0.0.0", sharedRecordGroup, 7200, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckSharedrecordAExists(ctx context.Context, resourceName string, v *dns.SharedrecordA) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			SharedrecordAAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForSharedrecordA).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetSharedrecordAResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetSharedrecordAResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckSharedrecordADestroy(ctx context.Context, v *dns.SharedrecordA) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			SharedrecordAAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForSharedrecordA).
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

func testAccCheckSharedrecordADisappears(ctx context.Context, v *dns.SharedrecordA) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			SharedrecordAAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccSharedrecordABasicConfig(name, ipv4addr, sharedRecordGroup string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_a" "test" {
    name                = %q
    ipv4addr            = %q
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
}
`, name, ipv4addr)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordAComment(name, ipv4addr, sharedRecordGroup, comment string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_a" "test_comment" {
    name                = %q
    ipv4addr            = %q
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
    comment             = %q
}
`, name, ipv4addr, comment)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordADisable(name, ipv4addr, sharedRecordGroup string, disable bool) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_a" "test_disable" {
    name                = %q
    ipv4addr            = %q
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
    disable             = %t
}
`, name, ipv4addr, disable)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordAExtAttrs(name, ipv4addr, sharedRecordGroup string, extAttrs map[string]any) string {
	extAttrsStr := utils.ConvertMapToHCL(extAttrs)
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_a" "test_extattrs" {
    name                = %q
    ipv4addr            = %q
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
    extattrs            = %s
}
`, name, ipv4addr, extAttrsStr)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordAIpv4addr(name, ipv4addr, sharedRecordGroup string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_a" "test_ipv4addr" {
    name                = %q
    ipv4addr            = %q
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
}
`, name, ipv4addr)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordAName(name, ipv4addr, sharedRecordGroup string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_a" "test_name" {
    name                = %q
    ipv4addr            = %q
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
}
`, name, ipv4addr)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordASharedRecordGroup(name, ipv4addr, sharedRecordGroup string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_a" "test_shared_record_group" {
    name                = %q
    ipv4addr            = %q
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
}
`, name, ipv4addr)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordATtl(name, ipv4addr, sharedRecordGroup string, ttl int, useTtl bool) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_a" "test_ttl" {
    name                = %q
    ipv4addr            = %q
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
    ttl                 = %d
    use_ttl             = %t
}
`, name, ipv4addr, ttl, useTtl)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordAUseTtl(name, ipv4addr, sharedRecordGroup string, ttl int, useTtl bool) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_a" "test_use_ttl" {
    name                = %q
    ipv4addr            = %q
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
    ttl                 = %d
    use_ttl             = %t
}
`, name, ipv4addr, ttl, useTtl)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}
