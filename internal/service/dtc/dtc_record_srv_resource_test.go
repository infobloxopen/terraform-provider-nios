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

var readableAttributesForDtcRecordSrv = "comment,disable,dtc_server,name,port,priority,target,ttl,use_ttl,weight"

func TestAccDtcRecordSrvResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test"
	var v dtc.DtcRecordSrv

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvBasicConfig(21, 10, "infoblox.com", 3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordSrvResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_record_srv.test"
	var v dtc.DtcRecordSrv

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcRecordSrvDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcRecordSrvBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					testAccCheckDtcRecordSrvDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcRecordSrvResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test_comment"
	var v dtc.DtcRecordSrv

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordSrvComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordSrvResource_Disable(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test_disable"
	var v dtc.DtcRecordSrv

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordSrvDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordSrvResource_DtcServer(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test_dtc_server"
	var v dtc.DtcRecordSrv

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvDtcServer("DTC_SERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dtc_server", "DTC_SERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordSrvDtcServer("DTC_SERVER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dtc_server", "DTC_SERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordSrvResource_Name(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test_name"
	var v dtc.DtcRecordSrv

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordSrvName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordSrvResource_Port(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test_port"
	var v dtc.DtcRecordSrv

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvPort("PORT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "PORT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordSrvPort("PORT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "PORT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordSrvResource_Priority(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test_priority"
	var v dtc.DtcRecordSrv

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvPriority("PRIORITY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "priority", "PRIORITY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordSrvPriority("PRIORITY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "priority", "PRIORITY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordSrvResource_Target(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test_target"
	var v dtc.DtcRecordSrv

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvTarget("TARGET_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "target", "TARGET_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordSrvTarget("TARGET_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "target", "TARGET_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordSrvResource_Ttl(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test_ttl"
	var v dtc.DtcRecordSrv

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvTtl("TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordSrvTtl("TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordSrvResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test_use_ttl"
	var v dtc.DtcRecordSrv

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvUseTtl("USE_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "USE_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordSrvUseTtl("USE_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "USE_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordSrvResource_Weight(t *testing.T) {
	var resourceName = "nios_dtc_record_srv.test_weight"
	var v dtc.DtcRecordSrv

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordSrvWeight("WEIGHT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "weight", "WEIGHT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordSrvWeight("WEIGHT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordSrvExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "weight", "WEIGHT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDtcRecordSrvExists(ctx context.Context, resourceName string, v *dtc.DtcRecordSrv) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DTCAPI.
			DtcRecordSrvAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDtcRecordSrv).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDtcRecordSrvResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDtcRecordSrvResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDtcRecordSrvDestroy(ctx context.Context, v *dtc.DtcRecordSrv) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DTCAPI.
			DtcRecordSrvAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDtcRecordSrv).
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

func testAccCheckDtcRecordSrvDisappears(ctx context.Context, v *dtc.DtcRecordSrv) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DTCAPI.
			DtcRecordSrvAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDtcRecordSrvBasicConfig(string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_srv" "test" {
}
`)
}

func testAccDtcRecordSrvComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_srv" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccDtcRecordSrvDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_srv" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccDtcRecordSrvDtcServer(dtcServer string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_srv" "test_dtc_server" {
    dtc_server = %q
}
`, dtcServer)
}

func testAccDtcRecordSrvName(name string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_srv" "test_name" {
    name = %q
}
`, name)
}

func testAccDtcRecordSrvPort(port string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_srv" "test_port" {
    port = %q
}
`, port)
}

func testAccDtcRecordSrvPriority(priority string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_srv" "test_priority" {
    priority = %q
}
`, priority)
}

func testAccDtcRecordSrvTarget(target string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_srv" "test_target" {
    target = %q
}
`, target)
}

func testAccDtcRecordSrvTtl(ttl string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_srv" "test_ttl" {
    ttl = %q
}
`, ttl)
}

func testAccDtcRecordSrvUseTtl(useTtl string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_srv" "test_use_ttl" {
    use_ttl = %q
}
`, useTtl)
}

func testAccDtcRecordSrvWeight(weight string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_srv" "test_weight" {
    weight = %q
}
`, weight)
}
