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

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/acctest"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/utils"
)

var readableAttributesForDtcServer = "extattrs,auto_create_host_record,disable,comment,disable,health,host,monitors,name,sni_hostname,use_sni_hostname"

func TestAccDtcServerResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_server.test"
	var v dtc.DtcServer

	name := acctest.RandomName()
	host := acctest.RandomIP()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcServerBasicConfig(name, host),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "host", host),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "auto_create_host_record", "true"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "sni_hostname", ""),
					resource.TestCheckResourceAttr(resourceName, "use_sni_hostname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcServerResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_server.test"
	var v dtc.DtcServer

	name := acctest.RandomName()
	host := acctest.RandomIP()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcServerDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcServerBasicConfig(name, host),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					testAccCheckDtcServerDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcServerResource_AutoCreateHostRecord(t *testing.T) {
	var resourceName = "nios_dtc_server.test_auto_create_host_record"
	var v dtc.DtcServer

	name := acctest.RandomName()
	host := acctest.RandomIP()
	autoCreateHostRecord := "false"
	autoCreateHostRecordUpdate := "true"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcServerAutoCreateHostRecord(name, host, autoCreateHostRecord),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_create_host_record", autoCreateHostRecord),
				),
			},
			// Update and Read
			{
				Config: testAccDtcServerAutoCreateHostRecord(name, host, autoCreateHostRecordUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_create_host_record", autoCreateHostRecordUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcServerResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_server.test_comment"
	var v dtc.DtcServer

	name := acctest.RandomName()
	host := acctest.RandomIP()
	comment := "initial comment"
	commentUpdate := "updated comment"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcServerComment(name, host, comment),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment),
				),
			},
			// Update and Read
			{
				Config: testAccDtcServerComment(name, host, commentUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", commentUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcServerResource_Disable(t *testing.T) {
	var resourceName = "nios_dtc_server.test_disable"
	var v dtc.DtcServer

	name := acctest.RandomName()
	host := acctest.RandomIP()
	disable := "true"
	disableUpdate := "false"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcServerDisable(name, host, disable),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", disable),
				),
			},
			// Update and Read
			{
				Config: testAccDtcServerDisable(name, host, disableUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", disableUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcServerResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dtc_server.test_extattrs"
	var v dtc.DtcServer

	name := acctest.RandomName()
	host := acctest.RandomIP()
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcServerExtAttrs(name, host, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccDtcServerExtAttrs(name, host, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcServerResource_Host(t *testing.T) {
	var resourceName = "nios_dtc_server.test_host"
	var v dtc.DtcServer

	name := acctest.RandomName()
	host := acctest.RandomIP()
	hostUpdate := acctest.RandomIP() // Use a different IP for update to test replace behavior
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcServerHost(name, host),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "host", host),
				),
			},
			// Update and Read
			{
				Config: testAccDtcServerHost(name, hostUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "host", hostUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcServerResource_Monitors(t *testing.T) {
	var resourceName = "nios_dtc_server.test_monitors"
	var v dtc.DtcServer

	name := acctest.RandomName()
	host := acctest.RandomIP()
	monitors := []map[string]any{
		{
			"host":    "3.2.2.2",
			"monitor": "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHBz:https",
		},
		{
			"host":    "3.231.2.2",
			"monitor": "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http",
		},
	}
	monitorsUpdate := []map[string]any{
		{
			"host":    "3.2.2.2",
			"monitor": "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHBz:https",
		},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcServerMonitors(name, host, monitors),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "monitors.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "monitors.0.host", "3.2.2.2"),
					resource.TestCheckResourceAttr(resourceName, "monitors.0.monitor", "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHBz:https"),
					resource.TestCheckResourceAttr(resourceName, "monitors.1.host", "3.231.2.2"),
					resource.TestCheckResourceAttr(resourceName, "monitors.1.monitor", "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcServerMonitors(name, host, monitorsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "monitors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "monitors.0.host", "3.2.2.2"),
					resource.TestCheckResourceAttr(resourceName, "monitors.0.monitor", "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHBz:https"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcServerResource_Name(t *testing.T) {
	var resourceName = "nios_dtc_server.test_name"
	var v dtc.DtcServer

	name := acctest.RandomName()
	updateName := acctest.RandomName()
	host := acctest.RandomIP()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcServerName(name, host),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDtcServerName(updateName, host),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", updateName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcServerResource_SniHostname(t *testing.T) {
	var resourceName = "nios_dtc_server.test_sni_hostname"
	var v dtc.DtcServer

	name := acctest.RandomName()
	host := acctest.RandomIP()
	sniHostName := acctest.RandomName()
	sniHostNameUpdate := acctest.RandomName() + "-update"
	useSniHostName := "true"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcServerSniHostname(name, host, sniHostName, useSniHostName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sni_hostname", sniHostName),
				),
			},
			// Update and Read
			{
				Config: testAccDtcServerSniHostname(name, host, sniHostNameUpdate, useSniHostName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sni_hostname", sniHostNameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcServerResource_UseSniHostname(t *testing.T) {
	var resourceName = "nios_dtc_server.test_use_sni_hostname"
	var v dtc.DtcServer

	name := acctest.RandomName()
	host := acctest.RandomIP()
	sniHostName := acctest.RandomName()
	useSniHostName := "true"
	useSniHostNameUpdate := "false"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcServerUseSniHostname(name, host, sniHostName, useSniHostName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_sni_hostname", useSniHostName),
				),
			},
			// Update and Read
			{
				Config: testAccDtcServerUseSniHostname(name, host, sniHostName, useSniHostNameUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_sni_hostname", useSniHostNameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDtcServerExists(ctx context.Context, resourceName string, v *dtc.DtcServer) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DTCAPI.
			DtcServerAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDtcServer).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDtcServerResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDtcServerResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDtcServerDestroy(ctx context.Context, v *dtc.DtcServer) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DTCAPI.
			DtcServerAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDtcServer).
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

func testAccCheckDtcServerDisappears(ctx context.Context, v *dtc.DtcServer) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DTCAPI.
			DtcServerAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDtcServerBasicConfig(name, host string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dtc_server" "test" {
	name = %q
	host = %q
}
`, name, host)
}

func testAccDtcServerAutoCreateHostRecord(name, host, autoCreateHostRecord string) string {
	return fmt.Sprintf(`
resource "nios_dtc_server" "test_auto_create_host_record" {
	name = %q
	host = %q
    auto_create_host_record = %q
}
`, name, host, autoCreateHostRecord)
}

func testAccDtcServerComment(name, host, comment string) string {
	return fmt.Sprintf(`
resource "nios_dtc_server" "test_comment" {
    name = %q
    host = %q
    comment = %q
}
`, name, host, comment)
}

func testAccDtcServerDisable(name, host, disable string) string {
	return fmt.Sprintf(`
resource "nios_dtc_server" "test_disable" {
	name = %q
	host = %q		
    disable = %q
}
`, name, host, disable)
}

func testAccDtcServerExtAttrs(name, host string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_dtc_server" "test_extattrs" {
	name = %q
	host = %q
    extattrs = %s
}
`, name, host, extattrsStr)
}

func testAccDtcServerHost(name, host string) string {
	return fmt.Sprintf(`
resource "nios_dtc_server" "test_host" {
	name = %q
    host = %q
}
`, name, host)
}

func testAccDtcServerMonitors(name, host string, monitors []map[string]any) string {
	monitorsHCL := formatMonitorsInterfaceToHCL(monitors)
	return fmt.Sprintf(`
resource "nios_dtc_server" "test_monitors" {
	name = %q
	host = %q
    monitors = %s
}
`, name, host, monitorsHCL)
}

func testAccDtcServerName(name, host string) string {
	return fmt.Sprintf(`
resource "nios_dtc_server" "test_name" {
    name = %q
	host = %q
}
`, name, host)
}

func testAccDtcServerSniHostname(name, host, sniHostname, useSniHostName string) string {
	return fmt.Sprintf(`
resource "nios_dtc_server" "test_sni_hostname" {
	name = %q
	host = %q
    sni_hostname = %q
	use_sni_hostname = %q
}
`, name, host, sniHostname, useSniHostName)
}

func testAccDtcServerUseSniHostname(name, host, sniHostname, useSniHostname string) string {
	return fmt.Sprintf(`
resource "nios_dtc_server" "test_use_sni_hostname" {
    name = %q
    host = %q
    sni_hostname = %q
    use_sni_hostname = %q
}
`, name, host, sniHostname, useSniHostname)
}

func formatMonitorsInterfaceToHCL(monitors []map[string]any) string {
	var monitorBlocks []string

	for _, monitor := range monitors {
		monitorBlock := fmt.Sprintf(`    {
	      host = %q
	      monitor = %q
	    }`, monitor["host"], monitor["monitor"])
		monitorBlocks = append(monitorBlocks, monitorBlock)
	}

	return fmt.Sprintf(`[
	%s
	  ]`, strings.Join(monitorBlocks, ",\n"))
}
