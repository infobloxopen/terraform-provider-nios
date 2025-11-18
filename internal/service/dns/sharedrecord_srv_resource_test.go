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

// TODO: OBJECTS TO BE PRESENT IN THE GRID FOR TESTS
// sharedrecordgroup: "example-sharedrecordgroup"

var readableAttributesForSharedrecordSrv = "comment,disable,dns_name,dns_target,extattrs,name,port,priority,shared_record_group,target,ttl,use_ttl,weight"

func TestAccSharedrecordSrvResource_basic(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_srv.test"
	var v dns.SharedrecordSrv
	name := acctest.RandomNameWithPrefix("sharedrecord-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordSrvBasicConfig(name, 10, 80, sharedRecordGroup, target, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "port", "10"),
					resource.TestCheckResourceAttr(resourceName, "priority", "80"),
					resource.TestCheckResourceAttr(resourceName, "shared_record_group", sharedRecordGroup),
					resource.TestCheckResourceAttr(resourceName, "target", target),
					resource.TestCheckResourceAttr(resourceName, "weight", "10"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordSrvResource_disappears(t *testing.T) {
	resourceName := "nios_dns_sharedrecord_srv.test"
	var v dns.SharedrecordSrv
	name := acctest.RandomNameWithPrefix("sharedrecord-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharedrecordSrvDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharedrecordSrvBasicConfig(name, 10, 80, sharedRecordGroup, target, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					testAccCheckSharedrecordSrvDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSharedrecordSrvResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_srv.test_comment"
	var v dns.SharedrecordSrv
	name := acctest.RandomNameWithPrefix("sharedrecord-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordSrvComment(name, 80, 10, sharedRecordGroup, target, 10, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordSrvComment(name, 80, 10, sharedRecordGroup, target, 10, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordSrvResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_srv.test_disable"
	var v dns.SharedrecordSrv
	name := acctest.RandomNameWithPrefix("sharedrecord-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordSrvDisable(name, 80, 10, sharedRecordGroup, target, 10, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordSrvDisable(name, 80, 10, sharedRecordGroup, target, 10, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordSrvResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_srv.test_extattrs"
	var v dns.SharedrecordSrv
	name := acctest.RandomNameWithPrefix("sharedrecord-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordSrvExtAttrs(name, 80, 10, sharedRecordGroup, target, 10, map[string]any{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordSrvExtAttrs(name, 80, 10, sharedRecordGroup, target, 10, map[string]any{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordSrvResource_Name(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_srv.test_name"
	var v dns.SharedrecordSrv
	name1 := acctest.RandomNameWithPrefix("sharedrecord-srv") + ".example.com"
	name2 := acctest.RandomNameWithPrefix("sharedrecord-srv-updated") + ".example.com"
	target := acctest.RandomName() + ".target.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordSrvName(name1, 80, 10, sharedRecordGroup, target, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordSrvName(name2, 80, 10, sharedRecordGroup, target, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordSrvResource_Port(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_srv.test_port"
	var v dns.SharedrecordSrv
	name := acctest.RandomNameWithPrefix("sharedrecord-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordSrvPort(name, 80, 10, sharedRecordGroup, target, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "80"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordSrvPort(name, 443, 10, sharedRecordGroup, target, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "443"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordSrvResource_Priority(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_srv.test_priority"
	var v dns.SharedrecordSrv
	name := acctest.RandomNameWithPrefix("sharedrecord-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordSrvPriority(name, 80, 10, sharedRecordGroup, target, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "priority", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordSrvPriority(name, 80, 20, sharedRecordGroup, target, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "priority", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordSrvResource_SharedRecordGroup(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_srv.test_shared_record_group"
	var v dns.SharedrecordSrv
	name := acctest.RandomNameWithPrefix("sharedrecord-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordSrvSharedRecordGroup(name, 80, 10, sharedRecordGroup, target, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "shared_record_group", sharedRecordGroup),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordSrvResource_Target(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_srv.test_target"
	var v dns.SharedrecordSrv
	name := acctest.RandomNameWithPrefix("sharedrecord-srv") + ".example.com"
	target1 := acctest.RandomName() + ".target1.com"
	target2 := acctest.RandomName() + ".target2.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordSrvTarget(name, 80, 10, sharedRecordGroup, target1, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "target", target1),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordSrvTarget(name, 80, 10, sharedRecordGroup, target2, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "target", target2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordSrvResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_srv.test_ttl"
	var v dns.SharedrecordSrv
	name := acctest.RandomNameWithPrefix("sharedrecord-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordSrvTtl(name, 80, 10, sharedRecordGroup, target, 10, 3600, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "3600"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordSrvTtl(name, 80, 10, sharedRecordGroup, target, 10, 7200, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "7200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordSrvResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_srv.test_use_ttl"
	var v dns.SharedrecordSrv
	name := acctest.RandomNameWithPrefix("sharedrecord-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordSrvUseTtl(name, 80, 10, sharedRecordGroup, target, 10, 3600, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordSrvUseTtl(name, 80, 10, sharedRecordGroup, target, 10, 3600, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordSrvResource_Weight(t *testing.T) {
	var resourceName = "nios_dns_sharedrecord_srv.test_weight"
	var v dns.SharedrecordSrv
	name := acctest.RandomNameWithPrefix("sharedrecord-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordSrvWeight(name, 80, 10, sharedRecordGroup, target, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "weight", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordSrvWeight(name, 80, 10, sharedRecordGroup, target, 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "weight", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckSharedrecordSrvExists(ctx context.Context, resourceName string, v *dns.SharedrecordSrv) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			SharedrecordSrvAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForSharedrecordSrv).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetSharedrecordSrvResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetSharedrecordSrvResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckSharedrecordSrvDestroy(ctx context.Context, v *dns.SharedrecordSrv) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			SharedrecordSrvAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForSharedrecordSrv).
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

func testAccCheckSharedrecordSrvDisappears(ctx context.Context, v *dns.SharedrecordSrv) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			SharedrecordSrvAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}
func testAccSharedrecordSrvBasicConfig(name string, port, priority int, sharedRecordGroup, target string, weight int) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_srv" "test" {
  name                = %q
  port                = %d
  priority            = %d
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
  target              = %q
  weight              = %d
}
`, name, port, priority, target, weight)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordSrvComment(name string, port, priority int, sharedRecordGroup, target string, weight int, comment string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_srv" "test_comment" {
  name                = %q
  port                = %d
  priority            = %d
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
  target              = %q
  weight              = %d
  comment             = %q
}
`, name, port, priority, target, weight, comment)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordSrvDisable(name string, port, priority int, sharedRecordGroup, target string, weight int, disable bool) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_srv" "test_disable" {
  name                = %q
  port                = %d
  priority            = %d
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
  target              = %q
  weight              = %d
  disable             = %t
}
`, name, port, priority, target, weight, disable)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordSrvExtAttrs(name string, port, priority int, sharedRecordGroup, target string, weight int, extAttrs map[string]any) string {
	extAttrsStr := utils.ConvertMapToHCL(extAttrs)
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_srv" "test_extattrs" {
  name                = %q
  port                = %d
  priority            = %d
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
  target              = %q
  weight              = %d
  extattrs            = %s
}
`, name, port, priority, target, weight, extAttrsStr)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordSrvName(name string, port, priority int, sharedRecordGroup, target string, weight int) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_srv" "test_name" {
  name                = %q
  port                = %d
  priority            = %d
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
  target              = %q
  weight              = %d
}
`, name, port, priority, target, weight)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordSrvPort(name string, port, priority int, sharedRecordGroup, target string, weight int) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_srv" "test_port" {
  name                = %q
  port                = %d
  priority            = %d
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
  target              = %q
  weight              = %d
}
`, name, port, priority, target, weight)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordSrvPriority(name string, port, priority int, sharedRecordGroup, target string, weight int) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_srv" "test_priority" {
  name                = %q
  port                = %d
  priority            = %d
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
  target              = %q
  weight              = %d
}
`, name, port, priority, target, weight)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordSrvSharedRecordGroup(name string, port, priority int, sharedRecordGroup, target string, weight int) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_srv" "test_shared_record_group" {
  name                = %q
  port                = %d
  priority            = %d
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
  target              = %q
  weight              = %d
}
`, name, port, priority, target, weight)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordSrvTarget(name string, port, priority int, sharedRecordGroup, target string, weight int) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_srv" "test_target" {
  name                = %q
  port                = %d
  priority            = %d
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
  target              = %q
  weight              = %d
}
`, name, port, priority, target, weight)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordSrvTtl(name string, port, priority int, sharedRecordGroup, target string, weight int, ttl int, useTtl bool) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_srv" "test_ttl" {
  name                = %q
  port                = %d
  priority            = %d
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
  target              = %q
  weight              = %d
  ttl                 = %d
  use_ttl             = %t
}
`, name, port, priority, target, weight, ttl, useTtl)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordSrvUseTtl(name string, port, priority int, sharedRecordGroup, target string, weight int, ttl int, useTtl bool) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_srv" "test_use_ttl" {
  name                = %q
  port                = %d
  priority            = %d
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
  target              = %q
  weight              = %d
  ttl                 = %d
  use_ttl             = %t
}
`, name, port, priority, target, weight, ttl, useTtl)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordSrvWeight(name string, port, priority int, sharedRecordGroup, target string, weight int) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_srv" "test_weight" {
  name                = %q
  port                = %d
  priority            = %d
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
  target              = %q
  weight              = %d
}
`, name, port, priority, target, weight)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccBaseSharedRecordGroup(name string) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecordgroup" "parent_sharedrecord_group" {
  name = %q
}
`, name)
}
