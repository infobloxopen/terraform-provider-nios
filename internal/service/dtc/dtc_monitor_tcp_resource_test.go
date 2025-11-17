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
	name := acctest.RandomNameWithPrefix("dtc-monitor-tcp")
	port := 49152

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorTcpBasicConfig(name, port),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "port", fmt.Sprintf("%d", port)),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "interval", "5"),
					resource.TestCheckResourceAttr(resourceName, "timeout", "15"),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "1"),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorTcpResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_monitor_tcp.test"
	var v dtc.DtcMonitorTcp
	name := acctest.RandomNameWithPrefix("dtc-monitor-tcp")
	port := 49152

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcMonitorTcpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcMonitorTcpBasicConfig(name, port),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					testAccCheckDtcMonitorTcpDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcMonitorTcpResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_monitor_tcp.test_comment"
	var v dtc.DtcMonitorTcp
	name := acctest.RandomNameWithPrefix("dtc-monitor-tcp")
	port := 49152

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorTcpComment(name, port, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorTcpComment(name, port, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorTcpResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dtc_monitor_tcp.test_extattrs"
	var v dtc.DtcMonitorTcp
	name := acctest.RandomNameWithPrefix("dtc-monitor-tcp")
	port := 49152
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorTcpExtAttrs(name , port , map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorTcpExtAttrs(name , port , map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorTcpResource_Interval(t *testing.T) {
	var resourceName = "nios_dtc_monitor_tcp.test_interval"
	var v dtc.DtcMonitorTcp
	name := acctest.RandomNameWithPrefix("dtc-monitor-tcp")
	port := 49152

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorTcpInterval(name , port , 4),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interval", "4"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorTcpInterval(name , port , 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interval", "10"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorTcpResource_Name(t *testing.T) {
	var resourceName = "nios_dtc_monitor_tcp.test_name"
	var v dtc.DtcMonitorTcp
	name := acctest.RandomNameWithPrefix("dtc-monitor-tcp")
	nameUpdate := acctest.RandomNameWithPrefix("dtc-monitor-tcp")
	port := 49152

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorTcpName(name, port),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorTcpName(nameUpdate , port ),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorTcpResource_Port(t *testing.T) {
	var resourceName = "nios_dtc_monitor_tcp.test_port"
	var v dtc.DtcMonitorTcp
	name := acctest.RandomNameWithPrefix("dtc-monitor-tcp")
	port := 49152
	portUpdate := 49153

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorTcpPort(name , port),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", fmt.Sprintf("%d", port)),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorTcpPort(name , portUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", fmt.Sprintf("%d", portUpdate)),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorTcpResource_RetryDown(t *testing.T) {
	var resourceName = "nios_dtc_monitor_tcp.test_retry_down"
	var v dtc.DtcMonitorTcp
	name := acctest.RandomNameWithPrefix("dtc-monitor-tcp")
	port := 49152

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorTcpRetryDown(name , port , 3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "3"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorTcpRetryDown(name , port , 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "5"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorTcpResource_RetryUp(t *testing.T) {
	var resourceName = "nios_dtc_monitor_tcp.test_retry_up"
	var v dtc.DtcMonitorTcp
	name := acctest.RandomNameWithPrefix("dtc-monitor-tcp")
	port := 49152

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorTcpRetryUp(name , port , 4),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "4"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorTcpRetryUp(name , port , 6),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "6"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorTcpResource_Timeout(t *testing.T) {
	var resourceName = "nios_dtc_monitor_tcp.test_timeout"
	var v dtc.DtcMonitorTcp
	name := acctest.RandomNameWithPrefix("dtc-monitor-tcp")
	port := 49152

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorTcpTimeout(name , port , 30),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "30"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorTcpTimeout(name , port , 40),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorTcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "40"),
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

func testAccDtcMonitorTcpBasicConfig(name string, port int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test" {
    name  = %q
    port  = %d
}
`, name, port)
}

func testAccDtcMonitorTcpRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccDtcMonitorTcpComment(name string , port int, comment string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test_comment" {
	name  = %q
	port  = %d
    comment = %q
}
`, name, port, comment)
}

func testAccDtcMonitorTcpExtAttrs(name string , port int , extAttrs map[string]string) string {
	extattrsStr := "{"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`%s = %q`, k, v)
	}
	extattrsStr += "}"
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test_extattrs" {
    name     = %q
    port     = %d
    extattrs = %s
}
`, name, port, extattrsStr)
}

func testAccDtcMonitorTcpInterval(name string , port int , interval int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test_interval" {
	name     = %q
	port     = %d
    interval = %d
}
`, name, port, interval)
}

func testAccDtcMonitorTcpName(name string , port int ) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test_name" {
    name = %q
    port = %d
}
`, name, port)
}

func testAccDtcMonitorTcpPort(name string , port int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test_port" {
    name = %q
    port = %d
}
`, name, port)
}

func testAccDtcMonitorTcpRetryDown(name string , port int , retryDown int ) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test_retry_down" {
	name = %q
	port = %d
    retry_down = %d
}
`, name, port, retryDown)
}

func testAccDtcMonitorTcpRetryUp(name string , port int , retryUp int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test_retry_up" {
    name = %q
    port = %d
    retry_up = %d
}
`, name, port, retryUp)
}

func testAccDtcMonitorTcpTimeout(name string , port int , timeout int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_tcp" "test_timeout" {
	name = %q
	port = %d
    timeout = %d
}
`, name, port, timeout)
}
