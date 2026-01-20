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

var readableAttributesForDtcMonitorPdp = "comment,extattrs,interval,name,port,retry_down,retry_up,timeout"

func TestAccDtcMonitorPdpResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_monitor_pdp.test"
	var v dtc.DtcMonitorPdp
	name := acctest.RandomNameWithPrefix("dtc-monitor-pdp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorPdpBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorPdpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "port", "2123"),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "1"),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "1"),
					resource.TestCheckResourceAttr(resourceName, "timeout", "15"),
					resource.TestCheckResourceAttr(resourceName,"interval", "5"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorPdpResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_monitor_pdp.test"
	var v dtc.DtcMonitorPdp
	name := acctest.RandomNameWithPrefix("dtc-monitor-pdp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcMonitorPdpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcMonitorPdpBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorPdpExists(context.Background(), resourceName, &v),
					testAccCheckDtcMonitorPdpDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcMonitorPdpResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_monitor_pdp.test_comment"
	var v dtc.DtcMonitorPdp
	name := acctest.RandomNameWithPrefix("dtc-monitor-pdp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorPdpComment(name , "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorPdpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorPdpComment(name , "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorPdpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorPdpResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dtc_monitor_pdp.test_extattrs"
	var v dtc.DtcMonitorPdp
	name := acctest.RandomNameWithPrefix("dtc-monitor-pdp")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorPdpExtAttrs(name , map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorPdpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorPdpExtAttrs(name , map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorPdpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorPdpResource_Interval(t *testing.T) {
	var resourceName = "nios_dtc_monitor_pdp.test_interval"
	var v dtc.DtcMonitorPdp
	name := acctest.RandomNameWithPrefix("dtc-monitor-pdp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorPdpInterval(name, 4),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorPdpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interval", "4"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorPdpInterval(name, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorPdpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interval", "10"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorPdpResource_Name(t *testing.T) {
	var resourceName = "nios_dtc_monitor_pdp.test_name"
	var v dtc.DtcMonitorPdp
	name := acctest.RandomNameWithPrefix("dtc-monitor-pdp")
	nameUpdate := acctest.RandomNameWithPrefix("dtc-monitor-pdp-updated")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorPdpName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorPdpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorPdpName(nameUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorPdpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorPdpResource_Port(t *testing.T) {
	var resourceName = "nios_dtc_monitor_pdp.test_port"
	var v dtc.DtcMonitorPdp
	name := acctest.RandomNameWithPrefix("dtc-monitor-pdp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorPdpPort(name , 2314),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorPdpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "2314"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorPdpPort(name , 4321),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorPdpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "4321"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorPdpResource_RetryDown(t *testing.T) {
	var resourceName = "nios_dtc_monitor_pdp.test_retry_down"
	var v dtc.DtcMonitorPdp
	name := acctest.RandomNameWithPrefix("dtc-monitor-pdp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorPdpRetryDown(name , 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorPdpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "5"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorPdpRetryDown(name , 3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorPdpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorPdpResource_RetryUp(t *testing.T) {
	var resourceName = "nios_dtc_monitor_pdp.test_retry_up"
	var v dtc.DtcMonitorPdp
	name := acctest.RandomNameWithPrefix("dtc-monitor-pdp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorPdpRetryUp(name , 2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorPdpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "2"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorPdpRetryUp(name , 4),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorPdpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "4"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorPdpResource_Timeout(t *testing.T) {
	var resourceName = "nios_dtc_monitor_pdp.test_timeout"
	var v dtc.DtcMonitorPdp
	name := acctest.RandomNameWithPrefix("dtc-monitor-pdp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorPdpTimeout(name , 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorPdpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "20"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorPdpTimeout(name , 25),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorPdpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "25"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDtcMonitorPdpExists(ctx context.Context, resourceName string, v *dtc.DtcMonitorPdp) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DTCAPI.
			DtcMonitorPdpAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDtcMonitorPdp).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDtcMonitorPdpResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDtcMonitorPdpResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDtcMonitorPdpDestroy(ctx context.Context, v *dtc.DtcMonitorPdp) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DTCAPI.
			DtcMonitorPdpAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDtcMonitorPdp).
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

func testAccCheckDtcMonitorPdpDisappears(ctx context.Context, v *dtc.DtcMonitorPdp) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DTCAPI.
			DtcMonitorPdpAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDtcMonitorPdpBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_pdp" "test" {
    name = %q
}
`, name)
}

func testAccDtcMonitorPdpComment(name, comment string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_pdp" "test_comment" {
	name = %q
    comment = %q
}
`, name, comment)
}

func testAccDtcMonitorPdpExtAttrs(name string , extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_dtc_monitor_pdp" "test_extattrs" {
	name = %q
    extattrs = %s
}
`, name, extattrsStr)
}

func testAccDtcMonitorPdpInterval(name string , interval int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_pdp" "test_interval" {
	name = %q
    interval = %d
}
`, name, interval)
}

func testAccDtcMonitorPdpName(name string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_pdp" "test_name" {
    name = %q
}
`, name)
}

func testAccDtcMonitorPdpPort(name string, port int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_pdp" "test_port" {
	name = %q
    port = %d
}
`, name, port)
}

func testAccDtcMonitorPdpRetryDown(name string , retryDown int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_pdp" "test_retry_down" {
	name = %q
    retry_down = %d
}
`, name, retryDown)
}

func testAccDtcMonitorPdpRetryUp(name string , retryUp int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_pdp" "test_retry_up" {
	name = %q
    retry_up = %d
}
`, name, retryUp)
}

func testAccDtcMonitorPdpTimeout(name string , timeout int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_pdp" "test_timeout" {
	name = %q
    timeout = %d
}
`, name, timeout)
}
