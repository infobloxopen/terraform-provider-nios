package dtc_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForDtcMonitorTcp = "comment,extattrs,interval,name,port,retry_down,retry_up,timeout"

func TestAccDtcMonitorTcpResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_monitor_tcp.test"
	var v dtc.DtcMonitorTcp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorTcpBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorTcpResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_monitor_tcp.test"
	var v dtc.DtcMonitorTcp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcMonitorTcpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcMonitorTcpBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					testAccCheckDtcMonitorTcpDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcMonitorTcpResource_Ref(t *testing.T) {
	var resourceName = "nios_dtc_monitor_tcp.test_ref"
	var v dtc.DtcMonitorTcp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorTcpRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorTcpRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorTcpResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_monitor_tcp.test_comment"
	var v dtc.DtcMonitorTcp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorTcpComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorTcpComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorTcpResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dtc_monitor_tcp.test_extattrs"
	var v dtc.DtcMonitorTcp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorTcpExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorTcpExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorTcpResource_Interval(t *testing.T) {
	var resourceName = "nios_dtc_monitor_tcp.test_interval"
	var v dtc.DtcMonitorTcp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorTcpInterval("INTERVAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interval", "INTERVAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorTcpInterval("INTERVAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interval", "INTERVAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorTcpResource_Name(t *testing.T) {
	var resourceName = "nios_dtc_monitor_tcp.test_name"
	var v dtc.DtcMonitorTcp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorTcpName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorTcpName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorTcpResource_Port(t *testing.T) {
	var resourceName = "nios_dtc_monitor_tcp.test_port"
	var v dtc.DtcMonitorTcp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorTcpPort("PORT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "PORT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorTcpPort("PORT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "PORT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorTcpResource_RetryDown(t *testing.T) {
	var resourceName = "nios_dtc_monitor_tcp.test_retry_down"
	var v dtc.DtcMonitorTcp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorTcpRetryDown("RETRY_DOWN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "RETRY_DOWN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorTcpRetryDown("RETRY_DOWN_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "RETRY_DOWN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorTcpResource_RetryUp(t *testing.T) {
	var resourceName = "nios_dtc_monitor_tcp.test_retry_up"
	var v dtc.DtcMonitorTcp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorTcpRetryUp("RETRY_UP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "RETRY_UP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorTcpRetryUp("RETRY_UP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "RETRY_UP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorTcpResource_Timeout(t *testing.T) {
	var resourceName = "nios_dtc_monitor_tcp.test_timeout"
	var v dtc.DtcMonitorTcp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorTcpTimeout("TIMEOUT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "TIMEOUT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorTcpTimeout("TIMEOUT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "TIMEOUT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDtcMonitorTcpExists(ctx context.Context, resourceName string, v *dtc.DtcMonitorTcp) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DTCAPI.
			DtcMonitorTcpAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDtcMonitorTcp).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDtcMonitorTcpResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDtcMonitorTcpResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDtcMonitorTcpDestroy(ctx context.Context, v *dtc.DtcMonitorTcp) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DTCAPI.
			DtcMonitorTcpAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDtcMonitorTcp).
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

func testAccCheckDtcMonitorTcpDisappears(ctx context.Context, v *dtc.DtcMonitorTcp) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DTCAPI.
			DtcMonitorTcpAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDtcMonitorTcpBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test" {
}
`)
}

func testAccDtcMonitorTcpRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccDtcMonitorTcpComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccDtcMonitorTcpExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccDtcMonitorTcpInterval(interval string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test_interval" {
    interval = %q
}
`, interval)
}

func testAccDtcMonitorTcpName(name string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test_name" {
    name = %q
}
`, name)
}

func testAccDtcMonitorTcpPort(port string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test_port" {
    port = %q
}
`, port)
}

func testAccDtcMonitorTcpRetryDown(retryDown string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test_retry_down" {
    retry_down = %q
}
`, retryDown)
}

func testAccDtcMonitorTcpRetryUp(retryUp string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test_retry_up" {
    retry_up = %q
}
`, retryUp)
}

func testAccDtcMonitorTcpTimeout(timeout string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test_timeout" {
    timeout = %q
}
`, timeout)
}
