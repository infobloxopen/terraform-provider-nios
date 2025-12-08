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

var readableAttributesForDtcRecordA = "auto_created,comment,disable,dtc_server,ipv4addr,ttl,use_ttl"

func TestAccDtcRecordAResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_record_a.test"
	var v dtc.DtcRecordA
	ipv4addr := acctest.RandomIP()
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordABasicConfig(serverName, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", ipv4addr),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_record_a.test"
	var v dtc.DtcRecordA
	ipv4addr := acctest.RandomIP()
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcRecordADestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcRecordABasicConfig(serverName, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					testAccCheckDtcRecordADisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcRecordAResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_record_a.test_comment"
	var v dtc.DtcRecordA
	ipv4addr := acctest.RandomIP()
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordAComment(serverName, ipv4addr, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordAComment(serverName, ipv4addr, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAResource_Disable(t *testing.T) {
	var resourceName = "nios_dtc_record_a.test_disable"
	var v dtc.DtcRecordA
	ipv4addr := acctest.RandomIP()
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordADisable(serverName, ipv4addr, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordADisable(serverName, ipv4addr, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAResource_DtcServer(t *testing.T) {
	var resourceName = "nios_dtc_record_a.test_dtc_server"
	var v dtc.DtcRecordA
	ipv4addr := acctest.RandomIP()
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordADtcServer(ipv4addr, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dtc_server", serverName),
				),
			},
			// Removed update call since the field can't be updated
		},
	})
}

func TestAccDtcRecordAResource_Ipv4addr(t *testing.T) {
	var resourceName = "nios_dtc_record_a.test_ipv4addr"
	var v dtc.DtcRecordA
	ipv4addr := acctest.RandomIP()
	ipv4addrUpdate := acctest.RandomIP()
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordAIpv4addr(serverName, ipv4addr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", ipv4addr),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordAIpv4addr(serverName, ipv4addrUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", ipv4addrUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAResource_Ttl(t *testing.T) {
	var resourceName = "nios_dtc_record_a.test_ttl"
	var v dtc.DtcRecordA
	ipv4addr := acctest.RandomIP()
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordATtl(ipv4addr, serverName, 30),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "30"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordATtl(ipv4addr, serverName, 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dtc_record_a.test_use_ttl"
	var v dtc.DtcRecordA
	ipv4addr := acctest.RandomIP()
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordAUseTtl(ipv4addr, serverName, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordAUseTtl(ipv4addr, serverName, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDtcRecordAExists(ctx context.Context, resourceName string, v *dtc.DtcRecordA) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DTCAPI.
			DtcRecordAAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDtcRecordA).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDtcRecordAResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDtcRecordAResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDtcRecordADestroy(ctx context.Context, v *dtc.DtcRecordA) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DTCAPI.
			DtcRecordAAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDtcRecordA).
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

func testAccCheckDtcRecordADisappears(ctx context.Context, v *dtc.DtcRecordA) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DTCAPI.
			DtcRecordAAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDtcRecordABasicConfig(serverName, ipv4addr string) string {
	config := fmt.Sprintf(`
resource "nios_dtc_record_a" "test" {
  dtc_server = nios_dtc_server.test.name
  ipv4addr   = %q
}
`, ipv4addr)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.3.3.4"), config}, "")
}
func testAccBaseWithDtcServer(name, host string) string {
	return fmt.Sprintf(`
resource "nios_dtc_server" "test" {
  name = %q
  host = %q
}
`, name, host)
}

func testAccDtcRecordAComment(serverName, ipv4addr, comment string) string {
	config := fmt.Sprintf(`
resource "nios_dtc_record_a" "test_comment" {
    dtc_server = nios_dtc_server.test.name
    ipv4addr   = %q
    comment = %q
}
`, ipv4addr, comment)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.3.3.4"), config}, "")
}

func testAccDtcRecordADisable(serverName, ipv4addr string, disable bool) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_a" "test_disable" {
    disable = %t
    dtc_server = nios_dtc_server.test.name
    ipv4addr   = %q
}
`, disable, ipv4addr)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.3.3.2"), config}, "")
}

func testAccDtcRecordADtcServer(ipv4addr, dtcServer string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_a" "test_dtc_server" {
    dtc_server = nios_dtc_server.test.name
    ipv4addr   = %q
}
`, ipv4addr)
	return strings.Join([]string{testAccBaseWithDtcServer(dtcServer, "2.2.2.2"), config}, "")
}

func testAccDtcRecordAIpv4addr(serverName, ipv4addr string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_a" "test_ipv4addr" {
    ipv4addr = %q
	dtc_server = nios_dtc_server.test.name
}
	`, ipv4addr)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.3.3.4"), config}, "")
}

func testAccDtcRecordATtl(ipv4addr, serverName string, ttl int) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_a" "test_ttl" {
		ipv4addr = %q
		dtc_server = nios_dtc_server.test.name
		ttl = %d
		use_ttl = true 
}`, ipv4addr, ttl)

	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.3.3.4"), config}, "")
}

func testAccDtcRecordAUseTtl(ipv4addr, serverName string, useTtl bool) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_a" "test_use_ttl" {
    ipv4addr = %q
    dtc_server = nios_dtc_server.test.name
	ttl = 20
    use_ttl = %t
}
	`, ipv4addr, useTtl)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.2.2.2"), config}, "")
}
