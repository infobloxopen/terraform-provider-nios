package rpz_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/rpz"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForRecordRpzSrv = "comment,disable,extattrs,name,port,priority,rp_zone,target,ttl,use_ttl,view,weight,zone"

func TestAccRecordRpzSrvResource_basic(t *testing.T) {
	var resourceName = "nios_rpz_record_srv.test"
	var v rpz.RecordRpzSrv
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzSrvBasicConfig(name, rpZone, target, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
					resource.TestCheckResourceAttr(resourceName, "port", "80"),
					resource.TestCheckResourceAttr(resourceName, "priority", "10"),
					resource.TestCheckResourceAttr(resourceName, "weight", "360"),
					resource.TestCheckResourceAttr(resourceName, "target", target),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzSrvResource_disappears(t *testing.T) {
	resourceName := "nios_rpz_record_srv.test"
	var v rpz.RecordRpzSrv
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzSrvDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzSrvBasicConfig(name, rpZone, target, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					testAccCheckRecordRpzSrvDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordRpzSrvResource_Comment(t *testing.T) {
	var resourceName = "nios_rpz_record_srv.test_comment"
	var v rpz.RecordRpzSrv
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	target := acctest.RandomName() + ".target.com"
	comment1 := "This is a new rpz srv record"
	comment2 := "This is an updated rpz srv record"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzSrvComment(name, rpZone, target, comment1, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzSrvComment(name, rpZone, target, comment2, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzSrvResource_Disable(t *testing.T) {
	var resourceName = "nios_rpz_record_srv.test_disable"
	var v rpz.RecordRpzSrv
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzSrvDisable(name, rpZone, target, "false", 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzSrvDisable(name, rpZone, target, "true", 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzSrvResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_rpz_record_srv.test_extattrs"
	var v rpz.RecordRpzSrv
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	target := acctest.RandomName() + ".target.com"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzSrvExtAttrs(name, rpZone, target, 80, 10, 360, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzSrvExtAttrs(name, rpZone, target, 80, 10, 360, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzSrvResource_Name(t *testing.T) {
	var resourceName = "nios_rpz_record_srv.test_name"
	var v rpz.RecordRpzSrv
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	target := acctest.RandomName() + ".target.com"
	name1 := acctest.RandomName() + "." + rpZone
	name2 := acctest.RandomName() + "." + rpZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzSrvName(name1, rpZone, target, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzSrvName(name2, rpZone, target, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzSrvResource_Port(t *testing.T) {
	var resourceName = "nios_rpz_record_srv.test_port"
	var v rpz.RecordRpzSrv
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzSrvPort(name, rpZone, target, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "80"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzSrvPort(name, rpZone, target, 8080, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "8080"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzSrvResource_Priority(t *testing.T) {
	var resourceName = "nios_rpz_record_srv.test_priority"
	var v rpz.RecordRpzSrv
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzSrvPriority(name, rpZone, target, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "priority", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzSrvPriority(name, rpZone, target, 80, 1, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "priority", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzSrvResource_RpZone(t *testing.T) {
	var resourceName = "nios_rpz_record_srv.test_rp_zone"
	var v rpz.RecordRpzSrv
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzSrvRpZone(name, rpZone, target, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rp_zone", rpZone),
				),
			},
			// Can't update rp_zone as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzSrvResource_Target(t *testing.T) {
	var resourceName = "nios_rpz_record_srv.test_target"
	var v rpz.RecordRpzSrv
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	target1 := acctest.RandomName() + ".target.com"
	target2 := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzSrvTarget(name, rpZone, target1, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "target", target1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzSrvTarget(name, rpZone, target2, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "target", target2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzSrvResource_Ttl(t *testing.T) {
	var resourceName = "nios_rpz_record_srv.test_ttl"
	var v rpz.RecordRpzSrv
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzSrvTtl(name, rpZone, target, "true", 80, 10, 360, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzSrvTtl(name, rpZone, target, "true", 80, 10, 360, 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzSrvResource_UseTtl(t *testing.T) {
	var resourceName = "nios_rpz_record_srv.test_use_ttl"
	var v rpz.RecordRpzSrv
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzSrvUseTtl(name, rpZone, target, "true", 80, 10, 360, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzSrvUseTtl(name, rpZone, target, "false", 80, 10, 360, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzSrvResource_View(t *testing.T) {
	var resourceName = "nios_rpz_record_srv.test_view"
	var v rpz.RecordRpzSrv
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	target := acctest.RandomName() + ".target.com"
	view := acctest.RandomNameWithPrefix("test-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzSrvView(name, rpZone, target, view, 80, 10, 365),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", view),
				),
			},
			// Can't update view as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordRpzSrvResource_Weight(t *testing.T) {
	var resourceName = "nios_rpz_record_srv.test_weight"
	var v rpz.RecordRpzSrv
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordRpzSrvWeight(name, rpZone, target, 80, 10, 365),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "weight", "365"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordRpzSrvWeight(name, rpZone, target, 80, 10, 720),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "weight", "720"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordRpzSrvExists(ctx context.Context, resourceName string, v *rpz.RecordRpzSrv) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.RPZAPI.
			RecordRpzSrvAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordRpzSrv).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordRpzSrvResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordRpzSrvResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordRpzSrvDestroy(ctx context.Context, v *rpz.RecordRpzSrv) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.RPZAPI.
			RecordRpzSrvAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordRpzSrv).
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

func testAccCheckRecordRpzSrvDisappears(ctx context.Context, v *rpz.RecordRpzSrv) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.RPZAPI.
			RecordRpzSrvAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordRpzSrvBasicConfig(name, rpZone, target string, port, priority, weight int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_srv" "test" {
	name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	port = %d
	priority = %d
	weight = %d
	target = %q
}
`, name, port, priority, weight, target)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzSrvComment(name, rpZone, target, comment string, port, priority, weight int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_srv" "test_comment" {
	name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	port = %d
	priority = %d
	weight = %d
	target = %q
    comment = %q
}
`, name, port, priority, weight, target, comment)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzSrvDisable(name, rpZone, target, disable string, port, priority, weight int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_srv" "test_disable" {
	name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	port = %d
	priority = %d
	weight = %d
	target = %q
    disable = %q
}
`, name, port, priority, weight, target, disable)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzSrvExtAttrs(name, rpZone, target string, port, priority, weight int, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
	%s = %q
	`, k, v)
	}
	extattrsStr += "\t}"

	config := fmt.Sprintf(`
resource "nios_rpz_record_srv" "test_extattrs" {
	name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	port = %d
	priority = %d
	weight = %d
	target = %q
    extattrs = %s
}
`, name, port, priority, weight, target, extattrsStr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzSrvName(name, rpZone, target string, port, priority, weight int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_srv" "test_name" {
    name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	port = %d
	priority = %d
	weight = %d
	target = %q
}
`, name, port, priority, weight, target)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzSrvPort(name, rpZone, target string, port, priority, weight int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_srv" "test_port" {
    name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	port = %d
	priority = %d
	weight = %d
	target = %q
}
`, name, port, priority, weight, target)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzSrvPriority(name, rpZone, target string, port, priority, weight int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_srv" "test_priority" {
    name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	port = %d
	priority = %d
	weight = %d
	target = %q
}
`, name, port, priority, weight, target)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzSrvRpZone(name, rpZone, target string, port, priority, weight int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_srv" "test_rp_zone" {
    name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	port = %d
	priority = %d
	weight = %d
	target = %q
}
`, name, port, priority, weight, target)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzSrvTarget(name, rpZone, target string, port, priority, weight int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_srv" "test_target" {
    name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	port = %d
	priority = %d
	weight = %d
	target = %q
}
`, name, port, priority, weight, target)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzSrvTtl(name, rpZone, target, use_ttl string, port, priority, weight, ttl int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_srv" "test_ttl" {
 	name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	port = %d
	priority = %d
	weight = %d
	target = %q
    ttl = %d
	use_ttl = %q
}
`, name, port, priority, weight, target, ttl, use_ttl)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzSrvUseTtl(name, rpZone, target, use_ttl string, port, priority, weight, ttl int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_srv" "test_use_ttl" {
   	name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	port = %d
	priority = %d
	weight = %d
	target = %q
    ttl = %d
	use_ttl = %q
}
`, name, port, priority, weight, target, ttl, use_ttl)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzSrvView(name, rpZone, target, view string, port, priority, weight int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_srv" "test_view" {
	name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	port = %d
	priority = %d
	weight = %d
	target = %q
	view = nios_dns_view.custom_view.name
}
`, name, port, priority, weight, target)

	return strings.Join([]string{testAccBaseWithView(view), testAccBaseWithZone(rpZone, "nios_dns_view.custom_view.name"), config}, "")
}

func testAccRecordRpzSrvWeight(name, rpZone, target string, port, priority, weight int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_srv" "test_weight" {
	name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	port = %d
	priority = %d
	weight = %d
	target = %q
}
`, name, port, priority, weight, target)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}
