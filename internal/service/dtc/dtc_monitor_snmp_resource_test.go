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

var readableAttributesForDtcMonitorSnmp = "comment,community,context,engine_id,extattrs,interval,name,oids,port,retry_down,retry_up,timeout,user,version"

func TestAccDtcMonitorSnmpResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test"
	var v dtc.DtcMonitorSnmp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_monitor_snmp.test"
	var v dtc.DtcMonitorSnmp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcMonitorSnmpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcMonitorSnmpBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					testAccCheckDtcMonitorSnmpDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcMonitorSnmpResource_Ref(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_ref"
	var v dtc.DtcMonitorSnmp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_comment"
	var v dtc.DtcMonitorSnmp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_Community(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_community"
	var v dtc.DtcMonitorSnmp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpCommunity("COMMUNITY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "community", "COMMUNITY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpCommunity("COMMUNITY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "community", "COMMUNITY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_Context(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_context"
	var v dtc.DtcMonitorSnmp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpContext("CONTEXT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "context", "CONTEXT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpContext("CONTEXT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "context", "CONTEXT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_EngineId(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_engine_id"
	var v dtc.DtcMonitorSnmp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpEngineId("ENGINE_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "engine_id", "ENGINE_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpEngineId("ENGINE_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "engine_id", "ENGINE_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_extattrs"
	var v dtc.DtcMonitorSnmp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_Interval(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_interval"
	var v dtc.DtcMonitorSnmp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpInterval("INTERVAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interval", "INTERVAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpInterval("INTERVAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interval", "INTERVAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_Name(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_name"
	var v dtc.DtcMonitorSnmp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_Oids(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_oids"
	var v dtc.DtcMonitorSnmp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpOids("OIDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "oids", "OIDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpOids("OIDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "oids", "OIDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_Port(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_port"
	var v dtc.DtcMonitorSnmp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpPort("PORT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "PORT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpPort("PORT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "PORT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_RetryDown(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_retry_down"
	var v dtc.DtcMonitorSnmp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpRetryDown("RETRY_DOWN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "RETRY_DOWN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpRetryDown("RETRY_DOWN_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "RETRY_DOWN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_RetryUp(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_retry_up"
	var v dtc.DtcMonitorSnmp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpRetryUp("RETRY_UP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "RETRY_UP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpRetryUp("RETRY_UP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "RETRY_UP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_Timeout(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_timeout"
	var v dtc.DtcMonitorSnmp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpTimeout("TIMEOUT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "TIMEOUT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpTimeout("TIMEOUT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "TIMEOUT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_User(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_user"
	var v dtc.DtcMonitorSnmp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpUser("USER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "user", "USER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpUser("USER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "user", "USER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_Version(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_version"
	var v dtc.DtcMonitorSnmp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpVersion("VERSION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "version", "VERSION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpVersion("VERSION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "version", "VERSION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDtcMonitorSnmpExists(ctx context.Context, resourceName string, v *dtc.DtcMonitorSnmp) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DTCAPI.
			DtcMonitorSnmpAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDtcMonitorSnmp).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDtcMonitorSnmpResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDtcMonitorSnmpResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDtcMonitorSnmpDestroy(ctx context.Context, v *dtc.DtcMonitorSnmp) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DTCAPI.
			DtcMonitorSnmpAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDtcMonitorSnmp).
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

func testAccCheckDtcMonitorSnmpDisappears(ctx context.Context, v *dtc.DtcMonitorSnmp) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DTCAPI.
			DtcMonitorSnmpAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDtcMonitorSnmpBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test" {
}
`)
}

func testAccDtcMonitorSnmpRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccDtcMonitorSnmpComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccDtcMonitorSnmpCommunity(community string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_community" {
    community = %q
}
`, community)
}

func testAccDtcMonitorSnmpContext(context string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_context" {
    context = %q
}
`, context)
}

func testAccDtcMonitorSnmpEngineId(engineId string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_engine_id" {
    engine_id = %q
}
`, engineId)
}

func testAccDtcMonitorSnmpExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccDtcMonitorSnmpInterval(interval string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_interval" {
    interval = %q
}
`, interval)
}

func testAccDtcMonitorSnmpName(name string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_name" {
    name = %q
}
`, name)
}

func testAccDtcMonitorSnmpOids(oids string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_oids" {
    oids = %q
}
`, oids)
}

func testAccDtcMonitorSnmpPort(port string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_port" {
    port = %q
}
`, port)
}

func testAccDtcMonitorSnmpRetryDown(retryDown string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_retry_down" {
    retry_down = %q
}
`, retryDown)
}

func testAccDtcMonitorSnmpRetryUp(retryUp string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_retry_up" {
    retry_up = %q
}
`, retryUp)
}

func testAccDtcMonitorSnmpTimeout(timeout string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_timeout" {
    timeout = %q
}
`, timeout)
}

func testAccDtcMonitorSnmpUser(user string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_user" {
    user = %q
}
`, user)
}

func testAccDtcMonitorSnmpVersion(version string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_version" {
    version = %q
}
`, version)
}
