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

var readableAttributesForDtcMonitorSnmp = "comment,community,context,engine_id,extattrs,interval,name,oids,port,retry_down,retry_up,timeout,user,version"

func TestAccDtcMonitorSnmpResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test"
	var v dtc.DtcMonitorSnmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-snmp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "community", "public"),
					resource.TestCheckResourceAttr(resourceName, "interval", "5"),
					resource.TestCheckResourceAttr(resourceName, "port", "161"),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "1"),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "1"),
					resource.TestCheckResourceAttr(resourceName, "timeout", "15"),
					resource.TestCheckResourceAttr(resourceName, "version", "V2C"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_monitor_snmp.test"
	var v dtc.DtcMonitorSnmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-snmp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcMonitorSnmpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcMonitorSnmpBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					testAccCheckDtcMonitorSnmpDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcMonitorSnmpResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_comment"
	var v dtc.DtcMonitorSnmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-snmp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpComment(name, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpComment(name, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_Community(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_community"
	var v dtc.DtcMonitorSnmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-snmp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpCommunity(name, "private"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "community", "private"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpCommunity(name, "trapuser"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "community", "trapuser"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_Context(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_context"
	var v dtc.DtcMonitorSnmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-snmp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpContext(name, "text"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "context", "text"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpContext(name, "update context"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "context", "update context"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_EngineId(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_engine_id"
	var v dtc.DtcMonitorSnmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-snmp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpEngineId(name, "66356e6574776f726B73"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "engine_id", "66356e6574776f726B73"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpEngineId(name, "800007DB03000C754120"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "engine_id", "800007DB03000C754120"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_extattrs"
	var v dtc.DtcMonitorSnmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-snmp")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpExtAttrs(name, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpExtAttrs(name, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_Interval(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_interval"
	var v dtc.DtcMonitorSnmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-snmp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpInterval(name, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interval", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpInterval(name, 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interval", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_Name(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_name"
	var v dtc.DtcMonitorSnmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-snmp")
	nameUpdate := acctest.RandomNameWithPrefix("dtc-monitor-snmp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpName(nameUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_Oids(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_oids"
	var v dtc.DtcMonitorSnmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-snmp")
	oids := []map[string]any{
		{
			"oid":       ".2",
			"condition": "EXACT",
			"first":     "10",
		},
		{
			"oid":       ".02",
			"condition": "RANGE",
			"first":     "2",
			"last":      "4",
			"type":      "INTEGER",
		},
		{
			"oid":       ".1",
			"condition": "EXACT",
			"first":     "20",
		},
	}
	oidsUpdate := []map[string]any{
		{
			"oid":       ".2",
			"condition": "LEQ",
			"first":     "10",
		},
		{
			"oid":       ".01",
			"condition": "GEQ",
			"first":     "25",
		},
		{
			"oid":       ".1",
			"condition": "LEQ",
			"first":     "20",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpOids(name, oids),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "oids.0.oid", ".2"),
					resource.TestCheckResourceAttr(resourceName, "oids.0.condition", "EXACT"),
					resource.TestCheckResourceAttr(resourceName, "oids.0.first", "10"),
					resource.TestCheckResourceAttr(resourceName, "oids.1.oid", ".02"),
					resource.TestCheckResourceAttr(resourceName, "oids.1.condition", "RANGE"),
					resource.TestCheckResourceAttr(resourceName, "oids.1.first", "2"),
					resource.TestCheckResourceAttr(resourceName, "oids.1.last", "4"),
					resource.TestCheckResourceAttr(resourceName, "oids.1.type", "INTEGER"),
					resource.TestCheckResourceAttr(resourceName, "oids.2.oid", ".1"),
					resource.TestCheckResourceAttr(resourceName, "oids.2.condition", "EXACT"),
					resource.TestCheckResourceAttr(resourceName, "oids.2.first", "20"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpOids(name, oidsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "oids.0.oid", ".2"),
					resource.TestCheckResourceAttr(resourceName, "oids.0.condition", "LEQ"),
					resource.TestCheckResourceAttr(resourceName, "oids.0.first", "10"),
					resource.TestCheckResourceAttr(resourceName, "oids.1.oid", ".01"),
					resource.TestCheckResourceAttr(resourceName, "oids.1.condition", "GEQ"),
					resource.TestCheckResourceAttr(resourceName, "oids.1.first", "25"),
					resource.TestCheckResourceAttr(resourceName, "oids.2.oid", ".1"),
					resource.TestCheckResourceAttr(resourceName, "oids.2.condition", "LEQ"),
					resource.TestCheckResourceAttr(resourceName, "oids.2.first", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_Port(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_port"
	var v dtc.DtcMonitorSnmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-snmp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpPort(name, 10161),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "10161"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpPort(name, 10162),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "10162"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_RetryDown(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_retry_down"
	var v dtc.DtcMonitorSnmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-snmp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpRetryDown(name, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "5"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpRetryDown(name, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_down", "10"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_RetryUp(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_retry_up"
	var v dtc.DtcMonitorSnmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-snmp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpRetryUp(name, 3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "3"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpRetryUp(name, 7),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retry_up", "7"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_Timeout(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_timeout"
	var v dtc.DtcMonitorSnmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-snmp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpTimeout(name, 30),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "30"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpTimeout(name, 45),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "45"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_User(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_user"
	var v dtc.DtcMonitorSnmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-snmp")
	snmpUser1 := "nios_security_snmp_user.snmpuser_parent"
	snmpUser2 := "nios_security_snmp_user.snmpuser_parent1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpUser(name, snmpUser1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "user", snmpUser1, "name"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpUser(name, snmpUser2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "user", snmpUser2, "name"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcMonitorSnmpResource_Version(t *testing.T) {
	var resourceName = "nios_dtc_monitor_snmp.test_version"
	var v dtc.DtcMonitorSnmp
	name := acctest.RandomNameWithPrefix("dtc-monitor-snmp")
	snmpUser := "nios_security_snmp_user.snmpuser_parent1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcMonitorSnmpVersion(name, "V1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "version", "V1"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcMonitorSnmpVersion(name, "V2C"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "version", "V2C"),
				),
			},
			{
				Config: testAccDtcMonitorSnmpVersionV3(name, "V3", snmpUser),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcMonitorSnmpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "version", "V3"),
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

func testAccDtcMonitorSnmpBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test" {
    name = %q
}
`, name)
}

func testAccDtcMonitorSnmpComment(name, comment string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_comment" {
	name = %q
    comment = %q
}
`, name, comment)
}

func testAccDtcMonitorSnmpCommunity(name, community string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_community" {
	name = %q
    community = %q
}
`, name, community)
}

func testAccDtcMonitorSnmpContext(name, context string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_context" {
	name = %q
    context = %q
}
`, name, context)
}

func testAccDtcMonitorSnmpEngineId(name, engineId string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_engine_id" {
	name = %q
    engine_id = %q
}
`, name, engineId)
}

func testAccDtcMonitorSnmpExtAttrs(name string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_extattrs" {
	name = %q
    extattrs = %s
}
`, name, extattrsStr)
}

func testAccDtcMonitorSnmpInterval(name string, interval int64) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_interval" {
	name = %q
    interval = %d
}
`, name, interval)
}

func testAccDtcMonitorSnmpName(name string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_name" {
    name = %q
}
`, name)
}

func testAccDtcMonitorSnmpOids(name string, oids []map[string]any) string {
	oidsStr := utils.ConvertSliceOfMapsToHCL(oids)
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_oids" {
	name = %q
    oids = %s
}
`, name, oidsStr)
}

func testAccDtcMonitorSnmpPort(name string, port int64) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_port" {
	name = %q
    port = %d
}
`, name, port)
}

func testAccDtcMonitorSnmpRetryDown(name string, retryDown int64) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_retry_down" {
	name = %q
    retry_down = %d
}
`, name, retryDown)
}

func testAccDtcMonitorSnmpRetryUp(name string, retryUp int64) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_retry_up" {
	name = %q
    retry_up = %d
}
`, name, retryUp)
}

func testAccDtcMonitorSnmpTimeout(name string, timeout int64) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_timeout" {
	name = %q
    timeout = %d
}
`, name, timeout)
}

func testAccDtcMonitorSnmpUser(name, user string) string {
	config := fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_user" {
    name = %q
    version = "V3"
    user = %s.name
}
`, name, user)
	return strings.Join([]string{testAccBaseWithSnmpUsers(), config}, "")
}

func testAccBaseWithSnmpUsers() string {
	snmpUser1 := acctest.RandomNameWithPrefix("snmpuser")
	snmpUser2 := acctest.RandomNameWithPrefix("snmpuser")

	return fmt.Sprintf(`
resource "nios_security_snmp_user" "snmpuser_parent" {
    name                    = %q
    authentication_protocol = "NONE"
    privacy_protocol        = "NONE"
}

resource "nios_security_snmp_user" "snmpuser_parent1" {
    name                    = %q
    authentication_protocol = "NONE"
    privacy_protocol        = "NONE"
}
`, snmpUser1, snmpUser2)
}

func testAccDtcMonitorSnmpVersion(name, version string) string {
	return fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_version" {
	name = %q
    version = %q
}
`, name, version)
}

func testAccDtcMonitorSnmpVersionV3(name, version, snmpUser string) string {
	config := fmt.Sprintf(`
resource "nios_dtc_monitor_snmp" "test_version" {
    name = %q
    version = %q
    user = %s.name
}
`, name, version, snmpUser)
	return strings.Join([]string{testAccBaseWithSnmpUsers(), config}, "")
}
