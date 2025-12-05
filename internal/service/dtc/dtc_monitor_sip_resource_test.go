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

var readableAttributesForDtcMonitorSip = "ciphers,client_cert,comment,extattrs,interval,name,port,request,result,result_code,retry_down,retry_up,timeout,transport,validate_cert"

func TestAccDtcMonitorSipResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_monitor_sip.test"
	var v dtc.DtcMonitorSip
	name := acctest.RandomNameWithPrefix("dtc-monitor-sip")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSipBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "validate_cert", "true"),
					resource.TestCheckResourceAttr(resourceName, "interval", "5"),
					resource.TestCheckResourceAttr(resourceName, "port", "5060"),
					resource.TestCheckResourceAttr(resourceName, "timeout", "15"),
					resource.TestCheckResourceAttr(resourceName, "transport", "TCP"),
					resource.TestCheckResourceAttr(resourceName, "result_code", "200"),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "1"),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "1"),
					resource.TestCheckResourceAttr(resourceName, "result", "CODE_IS"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSipResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_monitor_sip.test"
	var v dtc.DtcMonitorSip
	name := acctest.RandomNameWithPrefix("dtc-monitor-sip")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcMonitorSipDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcMonitorSipBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					testAccCheckDtcMonitorSipDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcMonitorSipResource_Ciphers(t *testing.T) {
	var resourceName = "nios_dtc_monitor_sip.test_ciphers"
	var v dtc.DtcMonitorSip
	name := acctest.RandomNameWithPrefix("dtc-monitor-sip")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSipCiphers(name, "DHE-RSA-AES256-SHA"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ciphers", "DHE-RSA-AES256-SHA"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSipCiphers(name, "ECDHE-ECDSA-AES256-GCM-SHA384"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ciphers", "ECDHE-ECDSA-AES256-GCM-SHA384"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSipResource_ClientCert(t *testing.T) {
	var resourceName = "nios_dtc_monitor_sip.test_client_cert"
	var v dtc.DtcMonitorSip
	name := acctest.RandomNameWithPrefix("dtc-monitor-sip")
	certificate1 := "dtc:certificate/ZG5zLmlkbnNfY2VydGlmaWNhdGUkYjNiY2FjYjM0YjYzNWM4NzkwMGMxODUxN2MzODA0NTIyMGI1MTEyZTQwNzMzZTM1YzNkNTdmZjk3MGMwYTliY2NkN2Q4NzhkODBmYjE0MmQzZmYzYTRkNjI3Zjk5NDlmZmYwOWMwMDdkZGEzMmQxMDhiMWZkMTlmYjRkOThiNTI:7e78dced4c4715b761a25427848e7153671c496d"
	certificate2 := "dtc:certificate/ZG5zLmlkbnNfY2VydGlmaWNhdGUkYmJiNmU0NDBlZGY5NDA3ODQ0NzAzMDA3MzgxN2RmYWU4OTg3YzQwYjg5N2E1NmQyZDMwMWQ5YjlkY2NkZjFiYmM3MmJmZmNlODg1ZTJkODdmMzQ5YjczYzI0OTY3YmYyMGFlNDQ5MTI2Zjg5OGMzMDk2MzJiYzI4YWFlMDM0NzI:6595a8c26f84ac126cec62e770dff3058118e658"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSipClientCert(name, certificate1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_cert", certificate1),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSipClientCert(name, certificate2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_cert", certificate2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSipResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_monitor_sip.test_comment"
	var v dtc.DtcMonitorSip
	name := acctest.RandomNameWithPrefix("dtc-monitor-sip")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSipComment(name, "This is a sip monitor comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSipComment(name, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSipResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dtc_monitor_sip.test_extattrs"
	var v dtc.DtcMonitorSip
	name := acctest.RandomNameWithPrefix("dtc-monitor-sip")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSipExtAttrs(name, map[string]string{"Site": extAttrValue1}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSipExtAttrs(name, map[string]string{"Site": extAttrValue2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSipResource_Interval(t *testing.T) {
	var resourceName = "nios_dtc_monitor_sip.test_interval"
	var v dtc.DtcMonitorSip
	name := acctest.RandomNameWithPrefix("dtc-monitor-sip")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSipInterval(name, 3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interval", "3"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSipInterval(name, 7),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interval", "7"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSipResource_Name(t *testing.T) {
	var resourceName = "nios_dtc_monitor_sip.test_name"
	var v dtc.DtcMonitorSip
	name := acctest.RandomNameWithPrefix("dtc-monitor-sip")
	nameUpdate := acctest.RandomNameWithPrefix("dtc-monitor-sip")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSipName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSipName(nameUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSipResource_Port(t *testing.T) {
	var resourceName = "nios_dtc_monitor_sip.test_port"
	var v dtc.DtcMonitorSip
	name := acctest.RandomNameWithPrefix("dtc-monitor-sip")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSipPort(name, 4),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "4"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSipPort(name, 8),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "8"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSipResource_Request(t *testing.T) {
	var resourceName = "nios_dtc_monitor_sip.test_request"
	var v dtc.DtcMonitorSip
	name := acctest.RandomNameWithPrefix("dtc-monitor-sip")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSipRequest(name, "OPTIONS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "request", "OPTIONS"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSipRequest(name, "REGISTER"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "request", "REGISTER"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSipResource_Result(t *testing.T) {
	var resourceName = "nios_dtc_monitor_sip.test_result"
	var v dtc.DtcMonitorSip
	name := acctest.RandomNameWithPrefix("dtc-monitor-sip")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSipResult(name, "CODE_IS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "result", "CODE_IS"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSipResult(name, "CODE_IS_NOT"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "result", "CODE_IS_NOT"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSipResource_ResultCode(t *testing.T) {
	var resourceName = "nios_dtc_monitor_sip.test_result_code"
	var v dtc.DtcMonitorSip
	name := acctest.RandomNameWithPrefix("dtc-monitor-sip")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSipResultCode(name, 400),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "result_code", "400"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSipResultCode(name, 486),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "result_code", "486"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSipResource_RetryDown(t *testing.T) {
	var resourceName = "nios_dtc_monitor_sip.test_retry_down"
	var v dtc.DtcMonitorSip
	name := acctest.RandomNameWithPrefix("dtc-monitor-sip")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSipRetryDown(name, 3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "3"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSipRetryDown(name, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "5"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSipResource_RetryUp(t *testing.T) {
	var resourceName = "nios_dtc_monitor_sip.test_retry_up"
	var v dtc.DtcMonitorSip
	name := acctest.RandomNameWithPrefix("dtc-monitor-sip")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSipRetryUp(name, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "5"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSipRetryUp(name, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "10"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSipResource_Timeout(t *testing.T) {
	var resourceName = "nios_dtc_monitor_sip.test_timeout"
	var v dtc.DtcMonitorSip
	name := acctest.RandomNameWithPrefix("dtc-monitor-sip")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSipTimeout(name, 30),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "30"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSipTimeout(name, 45),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "45"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSipResource_Transport(t *testing.T) {
	var resourceName = "nios_dtc_monitor_sip.test_transport"
	var v dtc.DtcMonitorSip
	name := acctest.RandomNameWithPrefix("dtc-monitor-sip")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSipTransport(name, "UDP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transport", "UDP"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSipTransport(name, "TLS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transport", "TLS"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSipResource_ValidateCert(t *testing.T) {
	var resourceName = "nios_dtc_monitor_sip.test_validate_cert"
	var v dtc.DtcMonitorSip
	name := acctest.RandomNameWithPrefix("dtc-monitor-sip")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSipValidateCert(name, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "validate_cert", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSipValidateCert(name, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSipExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "validate_cert", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDtcMonitorSipExists(ctx context.Context, resourceName string, v *dtc.DtcMonitorSip) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DTCAPI.
			DtcMonitorSipAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDtcMonitorSip).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDtcMonitorSipResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDtcMonitorSipResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDtcMonitorSipDestroy(ctx context.Context, v *dtc.DtcMonitorSip) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DTCAPI.
			DtcMonitorSipAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDtcMonitorSip).
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

func testAccCheckDtcMonitorSipDisappears(ctx context.Context, v *dtc.DtcMonitorSip) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DTCAPI.
			DtcMonitorSipAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDtcMonitorSipBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_sip" "test" {
	name = %q
}
`, name)
}

func testAccDtcMonitorSipCiphers(name, ciphers string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_sip" "test_ciphers" {
	name = %q
    ciphers = %q
}
`, name, ciphers)
}

func testAccDtcMonitorSipClientCert(name, clientCert string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_sip" "test_client_cert" {
	name = %q
	client_cert = %q
}
`, name, clientCert)
}

func testAccDtcMonitorSipComment(name, comment string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_sip" "test_comment" {
    name    = %q
    comment = %q
}
`, name, comment)
}

func testAccDtcMonitorSipExtAttrs(name string, extAttrs map[string]string) string {
	extattrsStr := "{"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`%s = %q`, k, v)
	}
	extattrsStr += "}"
	return fmt.Sprintf(`
resource "nios_dtc_monitor_sip" "test_extattrs" {
	name     = %q
    extattrs = %s  
}
`, name, extattrsStr)
}

func testAccDtcMonitorSipInterval(name string, interval int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_sip" "test_interval" {
    name     = %q
    interval = %d
}
`, name, interval)
}

func testAccDtcMonitorSipName(name string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_sip" "test_name" {
    name = %q
}
`, name)
}

func testAccDtcMonitorSipPort(name string, port int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_sip" "test_port" {
	name = %q
    port = %d
}
`, name, port)
}

func testAccDtcMonitorSipRequest(name, request string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_sip" "test_request" {
	name    = %q
    request = %q
}
`, name, request)
}

func testAccDtcMonitorSipResult(name, result string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_sip" "test_result" {
	name   = %q
    result = %q
}
`, name, result)
}

func testAccDtcMonitorSipResultCode(name string, resultCode int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_sip" "test_result_code" {
    name        = %q
    result_code = %d
}
`, name, resultCode)
}

func testAccDtcMonitorSipRetryDown(name string, retryDown int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_sip" "test_retry_down" {
	name       = %q
    retry_down = %d
}
`, name, retryDown)
}

func testAccDtcMonitorSipRetryUp(name string, retryUp int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_sip" "test_retry_up" {
	name     = %q
    retry_up = %d
}
`, name, retryUp)
}

func testAccDtcMonitorSipTimeout(name string, timeout int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_sip" "test_timeout" {
    name    = %q
    timeout = %d
}
`, name, timeout)
}

func testAccDtcMonitorSipTransport(name, transport string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_sip" "test_transport" {
	name      = %q
    transport = %q
}
`, name, transport)
}

func testAccDtcMonitorSipValidateCert(name string, validateCert bool) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_sip" "test_validate_cert" {
	name          = %q
    validate_cert = %t
}
`, name, validateCert)
}
