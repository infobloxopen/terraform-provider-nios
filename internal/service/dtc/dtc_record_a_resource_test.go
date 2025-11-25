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

var readableAttributesForDtcRecordA = "auto_created,comment,disable,dtc_server,ipv4addr,ttl,use_ttl"

func TestAccDtcRecordAResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_record_a.test"
	var v dtc.DtcRecordA

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordABasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_record_a.test"
	var v dtc.DtcRecordA

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcRecordADestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcRecordABasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					testAccCheckDtcRecordADisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcRecordAResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_record_a.test_comment"
	var v dtc.DtcRecordA

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordAComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordAComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAResource_Disable(t *testing.T) {
	var resourceName = "nios_dtc_record_a.test_disable"
	var v dtc.DtcRecordA

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordADisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordADisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAResource_DtcServer(t *testing.T) {
	var resourceName = "nios_dtc_record_a.test_dtc_server"
	var v dtc.DtcRecordA

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordADtcServer("DTC_SERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dtc_server", "DTC_SERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordADtcServer("DTC_SERVER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dtc_server", "DTC_SERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAResource_Ipv4addr(t *testing.T) {
	var resourceName = "nios_dtc_record_a.test_ipv4addr"
	var v dtc.DtcRecordA

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordAIpv4addr("IPV4ADDR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "IPV4ADDR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordAIpv4addr("IPV4ADDR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "IPV4ADDR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAResource_Ttl(t *testing.T) {
	var resourceName = "nios_dtc_record_a.test_ttl"
	var v dtc.DtcRecordA

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordATtl("TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordATtl("TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dtc_record_a.test_use_ttl"
	var v dtc.DtcRecordA

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordAUseTtl("USE_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "USE_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordAUseTtl("USE_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "USE_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDtcRecordAExists(ctx context.Context, resourceName string, v *dtc.DtcRecordA) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DTCAPI.
			DtcRecordAAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDtcRecordA).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDtcRecordAResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDtcRecordAResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDtcRecordADestroy(ctx context.Context, v *dtc.DtcRecordA) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DTCAPI.
			DtcRecordAAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDtcRecordA).
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

func testAccCheckDtcRecordADisappears(ctx context.Context, v *dtc.DtcRecordA) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DTCAPI.
			DtcRecordAAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDtcRecordABasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dtc_record_a" "test" {
}
`)
}

func testAccDtcRecordAComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_a" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccDtcRecordADisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_a" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccDtcRecordADtcServer(dtcServer string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_a" "test_dtc_server" {
    dtc_server = %q
}
`, dtcServer)
}

func testAccDtcRecordAIpv4addr(ipv4addr string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_a" "test_ipv4addr" {
    ipv4addr = %q
}
`, ipv4addr)
}

func testAccDtcRecordATtl(ttl string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_a" "test_ttl" {
    ttl = %q
}
`, ttl)
}

func testAccDtcRecordAUseTtl(useTtl string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_a" "test_use_ttl" {
    use_ttl = %q
}
`, useTtl)
}
