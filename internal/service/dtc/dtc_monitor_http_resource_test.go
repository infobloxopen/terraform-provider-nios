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

var readableAttributesForDtcMonitorHttp = "ciphers,client_cert,comment,content_check,content_check_input,content_check_op,content_check_regex,content_extract_group,content_extract_type,content_extract_value,enable_sni,extattrs,interval,name,port,request,result,result_code,retry_down,retry_up,secure,timeout,validate_cert"

func TestAccDtcMonitorHttpResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "content_check", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "content_check_input", "ALL"),
					resource.TestCheckResourceAttr(resourceName, "content_extract_group", "0"),
					resource.TestCheckResourceAttr(resourceName, "content_extract_type", "STRING"),
					resource.TestCheckResourceAttr(resourceName, "enable_sni", "false"),
					resource.TestCheckResourceAttr(resourceName, "interval", "5"),
					resource.TestCheckResourceAttr(resourceName, "port", "80"),
					resource.TestCheckResourceAttr(resourceName, "result", "ANY"),
					resource.TestCheckResourceAttr(resourceName, "result_code", "200"),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "1"),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "1"),
					resource.TestCheckResourceAttr(resourceName, "secure", "false"),
					resource.TestCheckResourceAttr(resourceName, "timeout", "15"),
					resource.TestCheckResourceAttr(resourceName, "validate_cert", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_monitor_http.test"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcMonitorHttpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcMonitorHttpBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					testAccCheckDtcMonitorHttpDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcMonitorHttpResource_Ciphers(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_ciphers"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpCiphers(name, "DHE-RSA-AES256-SHA"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ciphers", "DHE-RSA-AES256-SHA"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpCiphers(name, "DEFAULT"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ciphers", "DEFAULT"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ClientCert(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_client_cert"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")
	certificate1 := "dtc:certificate/ZG5zLmlkbnNfY2VydGlmaWNhdGUkYjNiY2FjYjM0YjYzNWM4NzkwMGMxODUxN2MzODA0NTIyMGI1MTEyZTQwNzMzZTM1YzNkNTdmZjk3MGMwYTliY2NkN2Q4NzhkODBmYjE0MmQzZmYzYTRkNjI3Zjk5NDlmZmYwOWMwMDdkZGEzMmQxMDhiMWZkMTlmYjRkOThiNTI:7e78dced4c4715b761a25427848e7153671c496d"
	certificate2 := "dtc:certificate/ZG5zLmlkbnNfY2VydGlmaWNhdGUkYmJiNmU0NDBlZGY5NDA3ODQ0NzAzMDA3MzgxN2RmYWU4OTg3YzQwYjg5N2E1NmQyZDMwMWQ5YjlkY2NkZjFiYmM3MmJmZmNlODg1ZTJkODdmMzQ5YjczYzI0OTY3YmYyMGFlNDQ5MTI2Zjg5OGMzMDk2MzJiYzI4YWFlMDM0NzI:6595a8c26f84ac126cec62e770dff3058118e658"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpClientCert(name, certificate1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_cert", certificate1),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpClientCert(name, certificate2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_cert", certificate2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_comment"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpComment(name, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpComment(name, "This comment is updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This comment is updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ContentCheck(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_content_check"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")
	contentCheckOp := "EQ"
	contentCheckRegex := "The current load is ([0-9]+)"
	contentExtractType := "STRING"
	contentExtractValue := "default value"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpContentCheck(name, "EXTRACT", contentCheckOp, contentCheckRegex, contentExtractType, contentExtractValue),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_check", "EXTRACT"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpContentCheck(name, "MATCH", contentCheckOp, contentCheckRegex, contentExtractType, contentExtractValue),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_check", "MATCH"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ContentCheckInput(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_content_check_input"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpContentCheckInput(name, "BODY"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_check_input", "BODY"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpContentCheckInput(name, "HEADERS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_check_input", "HEADERS"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ContentCheckOp(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_content_check_op"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpContentCheckOp(name, "GEQ"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_check_op", "GEQ"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpContentCheckOp(name, "LEQ"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_check_op", "LEQ"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpContentCheckOp(name, "EQ"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_check_op", "EQ"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpContentCheckOp(name, "NEQ"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_check_op", "NEQ"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ContentCheckRegex(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_content_check_regex"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpContentCheckRegex(name, "HTTP/1\\.[01] (200|201|204)"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_check_regex", "HTTP/1\\.[01] (200|201|204)"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpContentCheckRegex(name, "Status: (2[0-9]{2})"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_check_regex", "Status: (2[0-9]{2})"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ContentExtractGroup(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_content_extract_group"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpContentExtractGroup(name, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_extract_group", "5"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpContentExtractGroup(name, 8),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_extract_group", "8"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ContentExtractType(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_content_extract_type"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpContentExtractType(name, "STRING"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_extract_type", "STRING"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpContentExtractType(name, "INTEGER"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_extract_type", "INTEGER"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ContentExtractValue(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_content_extract_value"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpContentExtractValue(name, "SUCCESS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_extract_value", "SUCCESS"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpContentExtractValue(name, "ACTIVE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_extract_value", "ACTIVE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_EnableSni(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_enable_sni"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpEnableSni(name, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_sni", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpEnableSni(name, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_sni", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_extattrs"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpExtAttrs(name, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpExtAttrs(name, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_Interval(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_interval"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpInterval(name, 4),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interval", "4"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpInterval(name, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interval", "10"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_Name(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_name"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")
	nameUpdate := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpName(nameUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_Port(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_port"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpPort(name, 80),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "80"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpPort(name, 8080),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "8080"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_Request(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_request"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpRequest(name, "GET /api/health HTTP/1.1\nHost: example.com\nUser-Agent: NIOS-Monitor"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "request", "GET /api/health HTTP/1.1\nHost: example.com\nUser-Agent: NIOS-Monitor"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpRequest(name, "HEAD /resource HTTP/1.1\nHost: example.com\nAccept: */*"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "request", "HEAD /resource HTTP/1.1\nHost: example.com\nAccept: */*"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_Result(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_result"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpResult(name, "CODE_IS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "result", "CODE_IS"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpResult(name, "CODE_IS_NOT"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "result", "CODE_IS_NOT"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ResultCode(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_result_code"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpResultCode(name, 200),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "result_code", "200"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpResultCode(name, 404),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "result_code", "404"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_RetryDown(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_retry_down"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpRetryDown(name, 4),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "4"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpRetryDown(name, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "5"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_RetryUp(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_retry_up"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpRetryUp(name, 4),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "4"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpRetryUp(name, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "5"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_Secure(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_secure"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpSecure(name, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "secure", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpSecure(name, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "secure", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_Timeout(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_timeout"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpTimeout(name, 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "20"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpTimeout(name, 30),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "30"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ValidateCert(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_validate_cert"
	var v dtc.DtcMonitorHttp
	name := acctest.RandomNameWithPrefix("dtc-monitor-http")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpValidateCert(name, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "validate_cert", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpValidateCert("name", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "validate_cert", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDtcMonitorHttpExists(ctx context.Context, resourceName string, v *dtc.DtcMonitorHttp) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DTCAPI.
			DtcMonitorHttpAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDtcMonitorHttp).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDtcMonitorHttpResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDtcMonitorHttpResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDtcMonitorHttpDestroy(ctx context.Context, v *dtc.DtcMonitorHttp) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DTCAPI.
			DtcMonitorHttpAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDtcMonitorHttp).
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

func testAccCheckDtcMonitorHttpDisappears(ctx context.Context, v *dtc.DtcMonitorHttp) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DTCAPI.
			DtcMonitorHttpAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDtcMonitorHttpBasicConfig(name string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test" {
    name = %q
}
`, name)
}


func testAccDtcMonitorHttpCiphers(name, ciphers string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_ciphers" {
	name    = %q
    ciphers = %q
}
`, name, ciphers)
}

func testAccDtcMonitorHttpClientCert(name, clientCert string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_client_cert" {
	name    = %q
    client_cert = %q
}
`, name, clientCert)
}

func testAccDtcMonitorHttpComment(name, comment string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_comment" {
	name    = %q
    comment = %q
}
`, name, comment)
}

func testAccDtcMonitorHttpContentCheck(name, contentCheck, contentCheckOp, contentCheckRegex, contentExtractType, contentExtractValue string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_content_check" {
    name         = %q
	content_check = %q
    content_check_op = %q
    content_check_regex = %q
    content_extract_type = %q
    content_extract_value = %q
}
`, name, contentCheck, contentCheckOp, contentCheckRegex, contentExtractType, contentExtractValue)
}

func testAccDtcMonitorHttpContentCheckInput(name, contentCheckInput string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_content_check_input" {
	name    = %q
    content_check_input = %q
}
`, name, contentCheckInput)
}

func testAccDtcMonitorHttpContentCheckOp(name, contentCheckOp string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_content_check_op" {
	name    = %q
    content_check_op = %q
}
`, name, contentCheckOp)
}

func testAccDtcMonitorHttpContentCheckRegex(name, contentCheckRegex string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_content_check_regex" {
	name    = %q
    content_check_regex = %q
}
`, name, contentCheckRegex)
}

func testAccDtcMonitorHttpContentExtractGroup(name string, contentExtractGroup int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_content_extract_group" {
	name    = %q
    content_extract_group = %d
}
`, name, contentExtractGroup)
}

func testAccDtcMonitorHttpContentExtractType(name, contentExtractType string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_content_extract_type" {
	name    = %q
    content_extract_type = %q
}
`, name, contentExtractType)
}

func testAccDtcMonitorHttpContentExtractValue(name, contentExtractValue string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_content_extract_value" {
	name    = %q
    content_extract_value = %q
}
`, name, contentExtractValue)
}

func testAccDtcMonitorHttpEnableSni(name string, enableSni bool) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_enable_sni" {
	name    = %q
    enable_sni = %t
}
`, name, enableSni)
}

func testAccDtcMonitorHttpExtAttrs(name string, extAttrs map[string]string) string {
	extattrsStr := "{"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`%s = %q`, k, v)
	}
	extattrsStr += "}"
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_extattrs" {
	name    = %q
    extattrs = %s
}
`, name, extattrsStr)
}

func testAccDtcMonitorHttpInterval(name string, interval int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_interval" {
	name    = %q
    interval = %d
}
`, name, interval)
}

func testAccDtcMonitorHttpName(name string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_name" {
    name = %q
}
`, name)
}

func testAccDtcMonitorHttpPort(name string, port int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_port" {
    name = %q
    port = %d
}
`, name, port)
}

func testAccDtcMonitorHttpRequest(name, request string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_request" {
	name    = %q
    request = %q
}
`, name, request)
}

func testAccDtcMonitorHttpResult(name, result string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_result" {
	name    = %q
    result = %q
}
`, name, result)
}

func testAccDtcMonitorHttpResultCode(name string, resultCode int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_result_code" {
	name    = %q
    result_code = %d
}
`, name, resultCode)
}

func testAccDtcMonitorHttpRetryDown(name string, retryDown int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_retry_down" {
	name = %q	
    retry_down = %d
}
`, name, retryDown)
}

func testAccDtcMonitorHttpRetryUp(name string, RetryUp int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_retry_up" {
	name = %q
    retry_up = %d
}
`, name, RetryUp)
}

func testAccDtcMonitorHttpSecure(name string, secure bool) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_secure" {
	name    = %q
    secure = %t
}
`, name, secure)
}

func testAccDtcMonitorHttpTimeout(name string, timeout int) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_timeout" {
    name    = %q
    timeout = %d
}
`, name, timeout)
}

func testAccDtcMonitorHttpValidateCert(name string, validateCert bool) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_validate_cert" {
    name = %q
    validate_cert = %t
}
`, name, validateCert)
}
