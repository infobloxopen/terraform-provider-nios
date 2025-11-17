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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_monitor_http.test"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcMonitorHttpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcMonitorHttpBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					testAccCheckDtcMonitorHttpDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcMonitorHttpResource_Ref(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_ref"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_Ciphers(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_ciphers"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpCiphers("CIPHERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ciphers", "CIPHERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpCiphers("CIPHERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ciphers", "CIPHERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ClientCert(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_client_cert"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpClientCert("CLIENT_CERT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_cert", "CLIENT_CERT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpClientCert("CLIENT_CERT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_cert", "CLIENT_CERT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_comment"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ContentCheck(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_content_check"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpContentCheck("CONTENT_CHECK_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_check", "CONTENT_CHECK_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpContentCheck("CONTENT_CHECK_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_check", "CONTENT_CHECK_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ContentCheckInput(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_content_check_input"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpContentCheckInput("CONTENT_CHECK_INPUT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_check_input", "CONTENT_CHECK_INPUT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpContentCheckInput("CONTENT_CHECK_INPUT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_check_input", "CONTENT_CHECK_INPUT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ContentCheckOp(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_content_check_op"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpContentCheckOp("CONTENT_CHECK_OP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_check_op", "CONTENT_CHECK_OP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpContentCheckOp("CONTENT_CHECK_OP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_check_op", "CONTENT_CHECK_OP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ContentCheckRegex(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_content_check_regex"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpContentCheckRegex("CONTENT_CHECK_REGEX_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_check_regex", "CONTENT_CHECK_REGEX_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpContentCheckRegex("CONTENT_CHECK_REGEX_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_check_regex", "CONTENT_CHECK_REGEX_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ContentExtractGroup(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_content_extract_group"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpContentExtractGroup("CONTENT_EXTRACT_GROUP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_extract_group", "CONTENT_EXTRACT_GROUP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpContentExtractGroup("CONTENT_EXTRACT_GROUP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_extract_group", "CONTENT_EXTRACT_GROUP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ContentExtractType(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_content_extract_type"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpContentExtractType("CONTENT_EXTRACT_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_extract_type", "CONTENT_EXTRACT_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpContentExtractType("CONTENT_EXTRACT_TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_extract_type", "CONTENT_EXTRACT_TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ContentExtractValue(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_content_extract_value"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpContentExtractValue("CONTENT_EXTRACT_VALUE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_extract_value", "CONTENT_EXTRACT_VALUE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpContentExtractValue("CONTENT_EXTRACT_VALUE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "content_extract_value", "CONTENT_EXTRACT_VALUE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_EnableSni(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_enable_sni"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpEnableSni("ENABLE_SNI_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_sni", "ENABLE_SNI_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpEnableSni("ENABLE_SNI_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_sni", "ENABLE_SNI_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_extattrs"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_Interval(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_interval"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpInterval("INTERVAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interval", "INTERVAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpInterval("INTERVAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interval", "INTERVAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_Name(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_name"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_Port(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_port"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpPort("PORT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "PORT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpPort("PORT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "PORT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_Request(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_request"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpRequest("REQUEST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "request", "REQUEST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpRequest("REQUEST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "request", "REQUEST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_Result(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_result"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpResult("RESULT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "result", "RESULT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpResult("RESULT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "result", "RESULT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ResultCode(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_result_code"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpResultCode("RESULT_CODE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "result_code", "RESULT_CODE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpResultCode("RESULT_CODE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "result_code", "RESULT_CODE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_RetryDown(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_retry_down"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpRetryDown("RETRY_DOWN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "RETRY_DOWN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpRetryDown("RETRY_DOWN_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "RETRY_DOWN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_RetryUp(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_retry_up"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpRetryUp("RETRY_UP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "RETRY_UP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpRetryUp("RETRY_UP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "RETRY_UP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_Secure(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_secure"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpSecure("SECURE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "secure", "SECURE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpSecure("SECURE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "secure", "SECURE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_Timeout(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_timeout"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpTimeout("TIMEOUT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "TIMEOUT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpTimeout("TIMEOUT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "TIMEOUT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorHttpResource_ValidateCert(t *testing.T) {
	var resourceName = "nios_dtc_monitor_http.test_validate_cert"
	var v dtc.DtcMonitorHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorHttpValidateCert("VALIDATE_CERT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "validate_cert", "VALIDATE_CERT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorHttpValidateCert("VALIDATE_CERT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorHttpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "validate_cert", "VALIDATE_CERT_UPDATE_REPLACE_ME"),
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

func testAccDtcMonitorHttpBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test" {
}
`)
}

func testAccDtcMonitorHttpRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccDtcMonitorHttpCiphers(ciphers string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_ciphers" {
    ciphers = %q
}
`, ciphers)
}

func testAccDtcMonitorHttpClientCert(clientCert string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_client_cert" {
    client_cert = %q
}
`, clientCert)
}

func testAccDtcMonitorHttpComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccDtcMonitorHttpContentCheck(contentCheck string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_content_check" {
    content_check = %q
}
`, contentCheck)
}

func testAccDtcMonitorHttpContentCheckInput(contentCheckInput string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_content_check_input" {
    content_check_input = %q
}
`, contentCheckInput)
}

func testAccDtcMonitorHttpContentCheckOp(contentCheckOp string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_content_check_op" {
    content_check_op = %q
}
`, contentCheckOp)
}

func testAccDtcMonitorHttpContentCheckRegex(contentCheckRegex string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_content_check_regex" {
    content_check_regex = %q
}
`, contentCheckRegex)
}

func testAccDtcMonitorHttpContentExtractGroup(contentExtractGroup string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_content_extract_group" {
    content_extract_group = %q
}
`, contentExtractGroup)
}

func testAccDtcMonitorHttpContentExtractType(contentExtractType string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_content_extract_type" {
    content_extract_type = %q
}
`, contentExtractType)
}

func testAccDtcMonitorHttpContentExtractValue(contentExtractValue string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_content_extract_value" {
    content_extract_value = %q
}
`, contentExtractValue)
}

func testAccDtcMonitorHttpEnableSni(enableSni string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_enable_sni" {
    enable_sni = %q
}
`, enableSni)
}

func testAccDtcMonitorHttpExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccDtcMonitorHttpInterval(interval string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_interval" {
    interval = %q
}
`, interval)
}

func testAccDtcMonitorHttpName(name string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_name" {
    name = %q
}
`, name)
}

func testAccDtcMonitorHttpPort(port string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_port" {
    port = %q
}
`, port)
}

func testAccDtcMonitorHttpRequest(request string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_request" {
    request = %q
}
`, request)
}

func testAccDtcMonitorHttpResult(result string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_result" {
    result = %q
}
`, result)
}

func testAccDtcMonitorHttpResultCode(resultCode string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_result_code" {
    result_code = %q
}
`, resultCode)
}

func testAccDtcMonitorHttpRetryDown(retryDown string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_retry_down" {
    retry_down = %q
}
`, retryDown)
}

func testAccDtcMonitorHttpRetryUp(retryUp string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_retry_up" {
    retry_up = %q
}
`, retryUp)
}

func testAccDtcMonitorHttpSecure(secure string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_secure" {
    secure = %q
}
`, secure)
}

func testAccDtcMonitorHttpTimeout(timeout string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_timeout" {
    timeout = %q
}
`, timeout)
}

func testAccDtcMonitorHttpValidateCert(validateCert string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_http" "test_validate_cert" {
    validate_cert = %q
}
`, validateCert)
}
