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

var readableAttributesForDtcRecordAaaa = "auto_created,comment,disable,dtc_server,ipv6addr,ttl,use_ttl"

func TestAccDtcRecordAaaaResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_record_aaaa.test"
	var v dtc.DtcRecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordAaaaBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAaaaResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_record_aaaa.test"
	var v dtc.DtcRecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcRecordAaaaDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcRecordAaaaBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					testAccCheckDtcRecordAaaaDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcRecordAaaaResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_record_aaaa.test_comment"
	var v dtc.DtcRecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordAaaaComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordAaaaComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAaaaResource_Disable(t *testing.T) {
	var resourceName = "nios_dtc_record_aaaa.test_disable"
	var v dtc.DtcRecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordAaaaDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordAaaaDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAaaaResource_DtcServer(t *testing.T) {
	var resourceName = "nios_dtc_record_aaaa.test_dtc_server"
	var v dtc.DtcRecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordAaaaDtcServer("DTC_SERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dtc_server", "DTC_SERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordAaaaDtcServer("DTC_SERVER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dtc_server", "DTC_SERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAaaaResource_Ipv6addr(t *testing.T) {
	var resourceName = "nios_dtc_record_aaaa.test_ipv6addr"
	var v dtc.DtcRecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordAaaaIpv6addr("IPV6ADDR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", "IPV6ADDR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordAaaaIpv6addr("IPV6ADDR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", "IPV6ADDR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAaaaResource_Ttl(t *testing.T) {
	var resourceName = "nios_dtc_record_aaaa.test_ttl"
	var v dtc.DtcRecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordAaaaTtl("TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordAaaaTtl("TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAaaaResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dtc_record_aaaa.test_use_ttl"
	var v dtc.DtcRecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordAaaaUseTtl("USE_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "USE_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordAaaaUseTtl("USE_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "USE_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDtcRecordAaaaExists(ctx context.Context, resourceName string, v *dtc.DtcRecordAaaa) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DTCAPI.
			DtcRecordAaaaAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDtcRecordAaaa).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDtcRecordAaaaResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDtcRecordAaaaResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDtcRecordAaaaDestroy(ctx context.Context, v *dtc.DtcRecordAaaa) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DTCAPI.
			DtcRecordAaaaAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDtcRecordAaaa).
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

func testAccCheckDtcRecordAaaaDisappears(ctx context.Context, v *dtc.DtcRecordAaaa) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DTCAPI.
			DtcRecordAaaaAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDtcRecordAaaaBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dtc_record_aaaa" "test" {
}
`)
}

func testAccDtcRecordAaaaComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_aaaa" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccDtcRecordAaaaDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_aaaa" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccDtcRecordAaaaDtcServer(dtcServer string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_aaaa" "test_dtc_server" {
    dtc_server = %q
}
`, dtcServer)
}

func testAccDtcRecordAaaaIpv6addr(ipv6addr string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_aaaa" "test_ipv6addr" {
    ipv6addr = %q
}
`, ipv6addr)
}

func testAccDtcRecordAaaaTtl(ttl string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_aaaa" "test_ttl" {
    ttl = %q
}
`, ttl)
}

func testAccDtcRecordAaaaUseTtl(useTtl string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_aaaa" "test_use_ttl" {
    use_ttl = %q
}
`, useTtl)
}
