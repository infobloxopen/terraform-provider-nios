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

var readableAttributesForDtcRecordNaptr = "comment,disable,dtc_server,flags,order,preference,regexp,replacement,services,ttl,use_ttl"

func TestAccDtcRecordNaptrResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test"
	var v dtc.DtcRecordNaptr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_record_naptr.test"
	var v dtc.DtcRecordNaptr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcRecordNaptrDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcRecordNaptrBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					testAccCheckDtcRecordNaptrDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcRecordNaptrResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test_comment"
	var v dtc.DtcRecordNaptr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordNaptrComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_Disable(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test_disable"
	var v dtc.DtcRecordNaptr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordNaptrDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_DtcServer(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test_dtc_server"
	var v dtc.DtcRecordNaptr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrDtcServer("DTC_SERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dtc_server", "DTC_SERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordNaptrDtcServer("DTC_SERVER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dtc_server", "DTC_SERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_Flags(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test_flags"
	var v dtc.DtcRecordNaptr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrFlags("FLAGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "flags", "FLAGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordNaptrFlags("FLAGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "flags", "FLAGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_Order(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test_order"
	var v dtc.DtcRecordNaptr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrOrder("ORDER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "order", "ORDER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordNaptrOrder("ORDER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "order", "ORDER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_Preference(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test_preference"
	var v dtc.DtcRecordNaptr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrPreference("PREFERENCE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preference", "PREFERENCE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordNaptrPreference("PREFERENCE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preference", "PREFERENCE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_Regexp(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test_regexp"
	var v dtc.DtcRecordNaptr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrRegexp("REGEXP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "regexp", "REGEXP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordNaptrRegexp("REGEXP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "regexp", "REGEXP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_Replacement(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test_replacement"
	var v dtc.DtcRecordNaptr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrReplacement("REPLACEMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "replacement", "REPLACEMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordNaptrReplacement("REPLACEMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "replacement", "REPLACEMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_Services(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test_services"
	var v dtc.DtcRecordNaptr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrServices("SERVICES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "services", "SERVICES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordNaptrServices("SERVICES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "services", "SERVICES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_Ttl(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test_ttl"
	var v dtc.DtcRecordNaptr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrTtl("TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordNaptrTtl("TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test_use_ttl"
	var v dtc.DtcRecordNaptr

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrUseTtl("USE_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "USE_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordNaptrUseTtl("USE_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "USE_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDtcRecordNaptrExists(ctx context.Context, resourceName string, v *dtc.DtcRecordNaptr) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DTCAPI.
			DtcRecordNaptrAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDtcRecordNaptr).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDtcRecordNaptrResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDtcRecordNaptrResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDtcRecordNaptrDestroy(ctx context.Context, v *dtc.DtcRecordNaptr) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DTCAPI.
			DtcRecordNaptrAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDtcRecordNaptr).
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

func testAccCheckDtcRecordNaptrDisappears(ctx context.Context, v *dtc.DtcRecordNaptr) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DTCAPI.
			DtcRecordNaptrAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDtcRecordNaptrBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dtc_record_naptr" "test" {
}
`)
}

func testAccDtcRecordNaptrComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_naptr" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccDtcRecordNaptrDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_naptr" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccDtcRecordNaptrDtcServer(dtcServer string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_naptr" "test_dtc_server" {
    dtc_server = %q
}
`, dtcServer)
}

func testAccDtcRecordNaptrFlags(flags string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_naptr" "test_flags" {
    flags = %q
}
`, flags)
}

func testAccDtcRecordNaptrOrder(order string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_naptr" "test_order" {
    order = %q
}
`, order)
}

func testAccDtcRecordNaptrPreference(preference string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_naptr" "test_preference" {
    preference = %q
}
`, preference)
}

func testAccDtcRecordNaptrRegexp(regexp string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_naptr" "test_regexp" {
    regexp = %q
}
`, regexp)
}

func testAccDtcRecordNaptrReplacement(replacement string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_naptr" "test_replacement" {
    replacement = %q
}
`, replacement)
}

func testAccDtcRecordNaptrServices(services string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_naptr" "test_services" {
    services = %q
}
`, services)
}

func testAccDtcRecordNaptrTtl(ttl string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_naptr" "test_ttl" {
    ttl = %q
}
`, ttl)
}

func testAccDtcRecordNaptrUseTtl(useTtl string) string {
	return fmt.Sprintf(`
resource "nios_dtc_record_naptr" "test_use_ttl" {
    use_ttl = %q
}
`, useTtl)
}
