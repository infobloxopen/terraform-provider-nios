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

var readableAttributesForDtcMonitorIcmp = "comment,extattrs,interval,name,retry_down,retry_up,timeout"

func TestAccDtcMonitorIcmpResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_monitor_icmp.test"
	var v dtc.DtcMonitorIcmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-icmp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorIcmpBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorIcmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "interval", "5"),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "1"),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "1"),
					resource.TestCheckResourceAttr(resourceName, "timeout", "15"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorIcmpResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_monitor_icmp.test"
	var v dtc.DtcMonitorIcmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-icmp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcMonitorIcmpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcMonitorIcmpBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorIcmpExists(context.Background(), resourceName, &v),
					testAccCheckDtcMonitorIcmpDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcMonitorIcmpResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_monitor_icmp.test_comment"
	var v dtc.DtcMonitorIcmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-icmp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorIcmpComment(name, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorIcmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorIcmpComment(name, "Updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorIcmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorIcmpResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dtc_monitor_icmp.test_extattrs"
	var v dtc.DtcMonitorIcmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-icmp")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorIcmpExtAttrs(name, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorIcmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorIcmpExtAttrs(name, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorIcmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorIcmpResource_Interval(t *testing.T) {
	var resourceName = "nios_dtc_monitor_icmp.test_interval"
	var v dtc.DtcMonitorIcmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-icmp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorIcmpInterval(name, 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorIcmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interval", "20"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorIcmpInterval(name, 30),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorIcmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interval", "30"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorIcmpResource_Name(t *testing.T) {
	var resourceName = "nios_dtc_monitor_icmp.test_name"
	var v dtc.DtcMonitorIcmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-icmp")
	nameUpdate := acctest.RandomNameWithPrefix("dtc-monitor-icmp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorIcmpName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorIcmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorIcmpName(nameUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorIcmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorIcmpResource_RetryDown(t *testing.T) {
	var resourceName = "nios_dtc_monitor_icmp.test_retry_down"
	var v dtc.DtcMonitorIcmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-icmp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorIcmpRetryDown(name, 2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorIcmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "2"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorIcmpRetryDown(name, 3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorIcmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorIcmpResource_RetryUp(t *testing.T) {
	var resourceName = "nios_dtc_monitor_icmp.test_retry_up"
	var v dtc.DtcMonitorIcmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-icmp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorIcmpRetryUp(name, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorIcmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "5"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorIcmpRetryUp(name, 4),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorIcmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "4"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorIcmpResource_Timeout(t *testing.T) {
	var resourceName = "nios_dtc_monitor_icmp.test_timeout"
	var v dtc.DtcMonitorIcmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-icmp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorIcmpTimeout(name, 30),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorIcmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "30"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorIcmpTimeout(name, 45),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorIcmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "45"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDtcMonitorIcmpExists(ctx context.Context, resourceName string, v *dtc.DtcMonitorIcmp) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DTCAPI.
			DtcMonitorIcmpAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDtcMonitorIcmp).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDtcMonitorIcmpResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDtcMonitorIcmpResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDtcMonitorIcmpDestroy(ctx context.Context, v *dtc.DtcMonitorIcmp) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DTCAPI.
			DtcMonitorIcmpAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDtcMonitorIcmp).
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

func testAccCheckDtcMonitorIcmpDisappears(ctx context.Context, v *dtc.DtcMonitorIcmp) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DTCAPI.
			DtcMonitorIcmpAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDtcMonitorIcmpBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_icmp" "test" {
	name = %q
}
`, name)
}

func testAccDtcMonitorIcmpComment(name, comment string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_icmp" "test_comment" {
	name = %q
	comment = %q
}
`, name, comment)
}

func testAccDtcMonitorIcmpExtAttrs(name string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_dtc_monitor_icmp" "test_extattrs" {
	name = %q
    extattrs = %s
}
`, name, extattrsStr)
}

func testAccDtcMonitorIcmpInterval(name string, interval int64) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_icmp" "test_interval" {
	name = %q
    interval = %d
}
`, name, interval)
}

func testAccDtcMonitorIcmpName(name string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_icmp" "test_name" {
	name = %q
}
`, name)
}

func testAccDtcMonitorIcmpRetryDown(name string, retryDown int64) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_icmp" "test_retry_down" {
	name = %q
    retry_down = %d
}
`, name, retryDown)
}

func testAccDtcMonitorIcmpRetryUp(name string, retryUp int64) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_icmp" "test_retry_up" {
	name = %q
    retry_up = %d
}
`, name, retryUp)
}

func testAccDtcMonitorIcmpTimeout(name string, timeout int64) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_icmp" "test_timeout" {
    name = %q
    timeout = %d
}
`, name, timeout)
}
