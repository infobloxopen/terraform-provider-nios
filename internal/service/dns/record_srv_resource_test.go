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

var readableAttributesForRecordSrv = "aws_rte53_record_info,cloud_info,comment,creation_time,creator,ddns_principal,ddns_protected,disable,dns_name,dns_target,extattrs,forbid_reclamation,last_queried,name,port,priority,reclaimable,shared_record_group,target,ttl,use_ttl,view,weight,zone"

func TestAccRecordSrvResource_basic(t *testing.T) {
	var resourceName = "nios_dns_record_srv.test"
	var v dns.RecordSrv
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSrvBasicConfig(name, target, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "target", target),
					resource.TestCheckResourceAttr(resourceName, "port", "80"),
					resource.TestCheckResourceAttr(resourceName, "priority", "10"),
					resource.TestCheckResourceAttr(resourceName, "weight", "360"),
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

func TestAccRecordSrvResource_disappears(t *testing.T) {
	resourceName := "nios_dns_record_srv.test"
	var v dns.RecordSrv
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordSrvDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordSrvBasicConfig(name, target, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					testAccCheckRecordSrvDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordSrvResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_record_srv.test_comment"
	var v dns.RecordSrv
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSrvComment(name, target, 80, 10, 360, "This is a new record"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a new record"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSrvComment(name, target, 80, 10, 360, "This is a updated record"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a updated record"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSrvResource_Creator(t *testing.T) {
	var resourceName = "nios_dns_record_srv.test_creator"
	var v dns.RecordSrv
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSrvCreator(name, target, 80, 10, 360, "STATIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSrvCreator(name, target, 80, 10, 360, "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "DYNAMIC"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSrvResource_DdnsPrincipal(t *testing.T) {
	var resourceName = "nios_dns_record_srv.test_ddns_principal"
	var v dns.RecordSrv
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSrvDdnsPrincipal(name, target, 80, 10, 360, "dhcp/server1@CORP.LOCAL", "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "dhcp/server1@CORP.LOCAL"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSrvDdnsPrincipal(name, target, 80, 10, 360, "dhcp/server2@CORP.LOCAL", "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "dhcp/server2@CORP.LOCAL"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSrvResource_DdnsProtected(t *testing.T) {
	var resourceName = "nios_dns_record_srv.test_ddns_protected"
	var v dns.RecordSrv
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSrvDdnsProtected(name, target, 80, 10, 360, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSrvDdnsProtected(name, target, 80, 10, 360, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSrvResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_record_srv.test_disable"
	var v dns.RecordSrv
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSrvDisable(name, target, 80, 10, 360, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSrvDisable(name, target, 80, 10, 360, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSrvResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_record_srv.test_extattrs"
	var v dns.RecordSrv
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSrvExtAttrs(name, target, 80, 10, 360, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSrvExtAttrs(name, target, 80, 10, 360, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSrvResource_ForbidReclamation(t *testing.T) {
	var resourceName = "nios_dns_record_srv.test_forbid_reclamation"
	var v dns.RecordSrv
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSrvForbidReclamation(name, target, 80, 10, 360, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSrvForbidReclamation(name, target, 80, 10, 360, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSrvResource_Name(t *testing.T) {
	var resourceName = "nios_dns_record_srv.test_name"
	var v dns.RecordSrv
	name1 := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	name2 := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSrvName(name1, target, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSrvName(name2, target, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSrvResource_Port(t *testing.T) {
	var resourceName = "nios_dns_record_srv.test_port"
	var v dns.RecordSrv
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSrvPort(name, target, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "80"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSrvPort(name, target, 8080, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "8080"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSrvResource_Priority(t *testing.T) {
	var resourceName = "nios_dns_record_srv.test_priority"
	var v dns.RecordSrv
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSrvPriority(name, target, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "priority", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSrvPriority(name, target, 80, 1, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "priority", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSrvResource_Target(t *testing.T) {
	var resourceName = "nios_dns_record_srv.test_target"
	var v dns.RecordSrv
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target1 := acctest.RandomName() + ".target.com"
	target2 := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSrvTarget(name, target1, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "target", target1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSrvTarget(name, target2, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "target", target2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSrvResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_record_srv.test_ttl"
	var v dns.RecordSrv
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSrvTtl(name, target, 80, 10, 360, 10, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSrvTtl(name, target, 80, 10, 360, 1000, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "1000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSrvResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_record_srv.test_use_ttl"
	var v dns.RecordSrv
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSrvUseTtl(name, target, 80, 10, 360, "false", 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSrvUseTtl(name, target, 80, 10, 360, "true", 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSrvResource_View(t *testing.T) {
	var resourceName = "nios_dns_record_srv.test_view"
	var v dns.RecordSrv
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSrvView(name, target, 80, 10, 360, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSrvView(name, target, 80, 10, 360, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSrvResource_Weight(t *testing.T) {
	var resourceName = "nios_dns_record_srv.test_weight"
	var v dns.RecordSrv
	name := acctest.RandomNameWithPrefix("record-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSrvWeight(name, target, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "weight", "360"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSrvWeight(name, target, 80, 10, 720),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "weight", "720"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordSrvExists(ctx context.Context, resourceName string, v *dns.RecordSrv) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			RecordSrvAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordSrv).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordSrvResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordSrvResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordSrvDestroy(ctx context.Context, v *dns.RecordSrv) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			RecordSrvAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordSrv).
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

func testAccCheckRecordSrvDisappears(ctx context.Context, v *dns.RecordSrv) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			RecordSrvAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordSrvBasicConfig(name, target string, port, priority, weight int) string {
	return fmt.Sprintf(`
resource "nios_dns_record_srv" "test" {
	name = %q
	target = %q
	port = %d
	priority = %d
	weight = %d
}
`, name, target, port, priority, weight)
}

func testAccRecordSrvComment(name, target string, port, priority, weight int, comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_srv" "test_comment" {
	name = %q
	target = %q
	port = %d
	priority = %d
	weight = %d
    comment = %q
}
`, name, target, port, priority, weight, comment)
}

func testAccRecordSrvCreator(name, target string, port, priority, weight int, creator string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_srv" "test_creator" {
	name = %q
	target = %q
	port = %d
	priority = %d
	weight = %d
    creator = %q
}
`, name, target, port, priority, weight, creator)
}

func testAccRecordSrvDdnsPrincipal(name, target string, port, priority, weight int, ddnsPrincipal, creator string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_srv" "test_ddns_principal" {
	name = %q
	target = %q
	port = %d
	priority = %d
	weight = %d
    ddns_principal = %q
	creator = %q
}
`, name, target, port, priority, weight, ddnsPrincipal, creator)
}

func testAccRecordSrvDdnsProtected(name, target string, port, priority, weight int, ddnsProtected string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_srv" "test_ddns_protected" {
	name = %q
	target = %q
	port = %d
	priority = %d
	weight = %d
    ddns_protected = %q
}
`, name, target, port, priority, weight, ddnsProtected)
}

func testAccRecordSrvDisable(name, target string, port, priority, weight int, disable string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_srv" "test_disable" {
	name = %q
	target = %q
	port = %d
	priority = %d
	weight = %d
    disable = %q
}
`, name, target, port, priority, weight, disable)
}

func testAccRecordSrvExtAttrs(name, target string, port, priority, weight int, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_dns_record_srv" "test_extattrs" {
	name = %q
	target = %q
	port = %d
	priority = %d
	weight = %d
    extattrs = %s
}
`, name, target, port, priority, weight, extattrsStr)
}

func testAccRecordSrvForbidReclamation(name, target string, port, priority, weight int, forbidReclamation string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_srv" "test_forbid_reclamation" {
	name = %q
	target = %q
	port = %d
	priority = %d
	weight = %d
    forbid_reclamation = %q
}
`, name, target, port, priority, weight, forbidReclamation)
}

func testAccRecordSrvName(name, target string, port, priority, weight int) string {
	return fmt.Sprintf(`
resource "nios_dns_record_srv" "test_name" {
	name = %q
	target = %q
	port = %d
	priority = %d
	weight = %d
}
`, name, target, port, priority, weight)
}

func testAccRecordSrvPort(name, target string, port, priority, weight int) string {
	return fmt.Sprintf(`
resource "nios_dns_record_srv" "test_port" {
	name = %q
	target = %q
	port = %d
	priority = %d
	weight = %d
}
`, name, target, port, priority, weight)
}

func testAccRecordSrvPriority(name, target string, port, priority, weight int) string {
	return fmt.Sprintf(`
resource "nios_dns_record_srv" "test_priority" {
	name = %q
	target = %q
	port = %d
	priority = %d
	weight = %d
}
`, name, target, port, priority, weight)
}

func testAccRecordSrvTarget(name, target string, port, priority, weight int) string {
	return fmt.Sprintf(`
resource "nios_dns_record_srv" "test_target" {
	name = %q
	target = %q
	port = %d
	priority = %d
	weight = %d
}
`, name, target, port, priority, weight)
}

func testAccRecordSrvTtl(name, target string, port, priority, weight, ttl int, useTTL string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_srv" "test_ttl" {
	name = %q
	target = %q
	port = %d
	priority = %d
	weight = %d
    ttl = %d
	use_ttl = %q
}
`, name, target, port, priority, weight, ttl, useTTL)
}

func testAccRecordSrvUseTtl(name, target string, port, priority, weight int, useTtl string, ttl int) string {
	return fmt.Sprintf(`
resource "nios_dns_record_srv" "test_use_ttl" {
	name = %q
	target = %q
	port = %d
	priority = %d
	weight = %d
    use_ttl = %q
	ttl = %d
}
`, name, target, port, priority, weight, useTtl, ttl)
}

func testAccRecordSrvView(name, target string, port, priority, weight int, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_srv" "test_view" {
	name = %q
	target = %q
	port = %d
	priority = %d
	weight = %d
    view = %q
}
`, name, target, port, priority, weight, view)
}

func testAccRecordSrvWeight(name, target string, port, priority, weight int) string {
	return fmt.Sprintf(`
resource "nios_dns_record_srv" "test_weight" {
	name = %q
	target = %q
	port = %d
	priority = %d
	weight = %d
}
`, name, target, port, priority, weight)
}
