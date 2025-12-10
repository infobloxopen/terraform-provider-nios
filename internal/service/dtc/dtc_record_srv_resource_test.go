package dtc_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForDtcRecordSrv = "comment,disable,dtc_server,name,port,priority,target,ttl,use_ttl,weight"

func TestAccDtcRecordSrvResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test"
	var v dtc.DtcRecordSrv
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvBasicConfig(21, 10, "infoblox.com", 3, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "port", "21"),
					resource.TestCheckResourceAttr(resourceName, "priority", "10"),
					resource.TestCheckResourceAttr(resourceName, "target", "infoblox.com"),
					resource.TestCheckResourceAttr(resourceName, "weight", "3"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordSrvResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_record_srv.test"
	var v dtc.DtcRecordSrv
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcRecordSrvDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcRecordSrvBasicConfig(21, 10, "infoblox.com", 3, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					testAccCheckDtcRecordSrvDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcRecordSrvResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test_comment"
	var v dtc.DtcRecordSrv
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvComment(21, 10, "infoblox.com", 3, serverName, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordSrvComment(21, 10, "infoblox.com", 3, serverName, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordSrvResource_Disable(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test_disable"
	var v dtc.DtcRecordSrv
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvDisable(21, 10, "infoblox.com", 3, false, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordSrvDisable(21, 10, "infoblox.com", 3, true, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordSrvResource_DtcServer(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test_dtc_server"
	var v dtc.DtcRecordSrv
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvDtcServer(21, 10, "infoblox.com", 3, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dtc_server", serverName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordSrvResource_Name(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test_name"
	var v dtc.DtcRecordSrv
	serverName := acctest.RandomNameWithPrefix("dtc-server")
	name := acctest.RandomNameWithPrefix("srv-record")
	nameUpdate := acctest.RandomNameWithPrefix("srv-record-updated")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvName(21, 10, "infoblox.com", 3, name, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordSrvName(21, 10, "infoblox.com", 3, nameUpdate, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordSrvResource_Port(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test_port"
	var v dtc.DtcRecordSrv
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvPort(24, 10, "infoblox.com", 3, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "24"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordSrvPort(21, 10, "infoblox.com", 3, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "21"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordSrvResource_Priority(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test_priority"
	var v dtc.DtcRecordSrv
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvPriority(24, 10, "infoblox.com", 3, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "priority", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordSrvPriority(24, 20, "infoblox.com", 3, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "priority", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordSrvResource_Target(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test_target"
	var v dtc.DtcRecordSrv
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvTarget(24, 10, "infoblox.com", 3, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "target", "infoblox.com"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordSrvTarget(24, 10, "uddiinfoblox.com", 3, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "target", "uddiinfoblox.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordSrvResource_Ttl(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test_ttl"
	var v dtc.DtcRecordSrv
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvTtl(24, 10, "infoblox.com", 3, serverName, 4),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "4"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordSrvTtl(24, 10, "infoblox.com", 3, serverName, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "5"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordSrvResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test_use_ttl"
	var v dtc.DtcRecordSrv
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvUseTtl(24, 10, "infoblox.com", 3, serverName, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordSrvUseTtl(24, 10, "infoblox.com", 3, serverName, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordSrvResource_Weight(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test_weight"
	var v dtc.DtcRecordSrv
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvWeight(24, 10, "infoblox.com", 30, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "weight", "30"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordSrvWeight(24, 10, "infoblox.com", 34, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "weight", "34"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDtcRecordSrvExists(ctx context.Context, resourceName string, v *dtc.DtcRecordSrv) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DTCAPI.
			DtcRecordSrvAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDtcRecordSrv).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDtcRecordSrvResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDtcRecordSrvResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDtcRecordSrvDestroy(ctx context.Context, v *dtc.DtcRecordSrv) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DTCAPI.
			DtcRecordSrvAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDtcRecordSrv).
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

func testAccCheckDtcRecordSrvDisappears(ctx context.Context, v *dtc.DtcRecordSrv) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DTCAPI.
			DtcRecordSrvAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDtcRecordSrvBasicConfig(port, priority int, target string, weight int, serverName string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_srv" "test" {
		port     = %d
		priority = %d
		target   = %q
		weight   = %d
		dtc_server = nios_dtc_server.test.name
	}		
	`, port, priority, target, weight)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.2.2.2"), config}, "")
}
func testAccBaseWithDtcServer(name, host string) string {
	return fmt.Sprintf(`
resource "nios_dtc_server" "test" {
  name = %q
  host = %q
}
`, name, host)
}

func testAccDtcRecordSrvComment(port, priority int, target string, weight int, serverName, comment string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_srv" "test_comment" {
		port     = %d
		priority = %d
		target   = %q
		weight   = %d
		dtc_server = nios_dtc_server.test.name
    comment = %q
}
	`, port, priority, target, weight, comment)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.2.2.2"), config}, "")
}

func testAccDtcRecordSrvDisable(port, priority int, target string, weight int, disable bool, serverName string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_srv" "test_disable" {
		port     = %d
		priority = %d
		target   = %q
		weight   = %d
		dtc_server = nios_dtc_server.test.name
    	disable = %t
	}
	`, port, priority, target, weight, disable)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.2.2.2"), config}, "")
}

func testAccDtcRecordSrvDtcServer(port, priority int, target string, weight int, serverName string) string {
	config := fmt.Sprintf(`
resource "nios_dtc_record_srv" "test_dtc_server" {
		port     = %d
		priority = %d
		target   = %q
		weight   = %d
		dtc_server = nios_dtc_server.test.name
}
`, port, priority, target, weight)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.2.2.2"), config}, "")
}

func testAccDtcRecordSrvName(port, priority int, target string, weight int, name string, serverName string) string {
	config := fmt.Sprintf(`
resource "nios_dtc_record_srv" "test_name" {
    name = %q
    port = %d
    priority = %d
    target = %q
    weight = %d
    dtc_server = nios_dtc_server.test.name
}
`, name, port, priority, target, weight)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.2.2.2"), config}, "")
}

func testAccDtcRecordSrvPort(port, priority int, target string, weight int, serverName string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_srv" "test_port" {
    	port     = %d
		priority = %d
		target   = %q
		weight   = %d
		dtc_server = nios_dtc_server.test.name
}
`, port, priority, target, weight)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.2.2.2"), config}, "")
}

func testAccDtcRecordSrvPriority(port, priority int, target string, weight int, serverName string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_srv" "test_priority" {
   port     = %d
		priority = %d
		target   = %q
		weight   = %d
		dtc_server = nios_dtc_server.test.name
	}
	`, port, priority, target, weight)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.2.2.2"), config}, "")
}

func testAccDtcRecordSrvTarget(port, priority int, target string, weight int, serverName string) string {
	config := fmt.Sprintf(`
resource "nios_dtc_record_srv" "test_target" {
    port     = %d
    priority = %d
    target   = %q
    weight   = %d
	dtc_server = nios_dtc_server.test.name
}
`, port, priority, target, weight)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.2.2.2"), config}, "")
}

func testAccDtcRecordSrvTtl(port, priority int, target string, weight int, serverName string, ttl int) string {
	config := fmt.Sprintf(`
resource "nios_dtc_record_srv" "test_ttl" {
    port     = %d
    priority = %d
    target   = %q
    weight   = %d
	ttl      = %d
	use_ttl  = true
	dtc_server = nios_dtc_server.test.name
}
`, port, priority, target, weight, ttl)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.2.2.2"), config}, "")
}

func testAccDtcRecordSrvUseTtl(port, priority int, target string, weight int, serverName string, useTtl bool) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_srv" "test_use_ttl" {
	 port     = %d
    priority = %d
    target   = %q
    weight   = %d
	ttl = 20
    use_ttl = %t
	dtc_server = nios_dtc_server.test.name
}`, port, priority, target, weight, useTtl)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.2.2.2"), config}, "")
}

func testAccDtcRecordSrvWeight(port, priority int, target string, weight int, serverName string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_srv" "test_weight" {
     port     = %d
    priority = %d
    target   = %q
    weight   = %d
	dtc_server = nios_dtc_server.test.name
	}
	`, port, priority, target, weight)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.2.2.2"), config}, "")
}
