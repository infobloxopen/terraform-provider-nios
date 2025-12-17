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

var readableAttributesForDtcRecordNaptr = "comment,disable,dtc_server,flags,order,preference,regexp,replacement,services,ttl,use_ttl"

func TestAccDtcRecordNaptrResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test"
	var v dtc.DtcRecordNaptr
	serverName := acctest.RandomNameWithPrefix("dtc-server")
	serverIp := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrBasicConfig(serverName, serverIp, 2, 5, "example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "order", "2"),
					resource.TestCheckResourceAttr(resourceName, "preference", "5"),
					resource.TestCheckResourceAttr(resourceName, "replacement", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "dtc_server", serverName),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_record_naptr.test"
	var v dtc.DtcRecordNaptr
	serverName := acctest.RandomNameWithPrefix("dtc-server")
	serverIp := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcRecordNaptrDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcRecordNaptrBasicConfig(serverName, serverIp, 2, 5, "example.com"),
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
	serverName := acctest.RandomNameWithPrefix("dtc-server")
	serverIp := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrComment(serverName, serverIp, 2, 5, "example.com", "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordNaptrComment(serverName, serverIp, 2, 5, "example.com", "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_Disable(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test_disable"
	var v dtc.DtcRecordNaptr
	serverName := acctest.RandomNameWithPrefix("dtc-server")
	serverIp := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrDisable(serverName, serverIp, 2, 5, "example.com", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordNaptrDisable(serverName, serverIp, 2, 5, "example.com", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_DtcServer(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test_dtc_server"
	var v dtc.DtcRecordNaptr
	serverName := acctest.RandomNameWithPrefix("dtc-server")
	serverIp := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrDtcServer(serverName, serverIp, 2, 5, "example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dtc_server", serverName),
				),
			},
			// Removed Update and Read step as dtc_server is required and cannot be updated
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_Flags(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test_flags"
	var v dtc.DtcRecordNaptr
	serverName := acctest.RandomNameWithPrefix("dtc-server")
	serverIp := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrFlags(serverName, serverIp, 2, 5, "example.com", "U"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "flags", "U"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordNaptrFlags(serverName, serverIp, 2, 5, "example.com", "S"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "flags", "S"),
				),
			},
			{
				Config: testAccDtcRecordNaptrFlags(serverName, serverIp, 2, 5, "example.com", "P"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "flags", "P"),
				),
			},
			{
				Config: testAccDtcRecordNaptrFlags(serverName, serverIp, 2, 5, "example.com", "A"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "flags", "A"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_Order(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test_order"
	var v dtc.DtcRecordNaptr
	serverName := acctest.RandomNameWithPrefix("dtc-server")
	serverIp := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrOrder(serverName, serverIp, 2, 5, "example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "order", "2"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordNaptrOrder(serverName, serverIp, 10, 5, "example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "order", "10"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_Preference(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test_preference"
	var v dtc.DtcRecordNaptr
	serverName := acctest.RandomNameWithPrefix("dtc-server")
	serverIp := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrPreference(serverName, serverIp, 2, 5, "example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preference", "5"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordNaptrPreference(serverName, serverIp, 2, 10, "example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preference", "10"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_Regexp(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test_regexp"
	var v dtc.DtcRecordNaptr
	serverName := acctest.RandomNameWithPrefix("dtc-server")
	serverIp := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrRegexp(serverName, serverIp, 2, 5, "example.com", "!^.*$!sip:info@example.com!"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "regexp", "!^.*$!sip:info@example.com!"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordNaptrRegexp(serverName, serverIp, 2, 5, "example.com", "!^.*$!sip:updated@example.com!"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "regexp", "!^.*$!sip:updated@example.com!"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_Replacement(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test_replacement"
	var v dtc.DtcRecordNaptr
	serverName := acctest.RandomNameWithPrefix("dtc-server")
	serverIp := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrReplacement(serverName, serverIp, 2, 5, "example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "replacement", "example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordNaptrReplacement(serverName, serverIp, 2, 5, "infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "replacement", "infoblox.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_Services(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test_services"
	var v dtc.DtcRecordNaptr
	serverName := acctest.RandomNameWithPrefix("dtc-server")
	serverIp := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrServices(serverName, serverIp, 2, 5, "example.com", "E2U+email"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "services", "E2U+email"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordNaptrServices(serverName, serverIp, 2, 5, "example.com", "E2U+sip"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "services", "E2U+sip"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_Ttl(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test_ttl"
	var v dtc.DtcRecordNaptr
	serverName := acctest.RandomNameWithPrefix("dtc-server")
	serverIp := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrTtl(serverName, serverIp, 2, 5, "example.com", "30"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "30"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordNaptrTtl(serverName, serverIp, 2, 5, "example.com", "60"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "60"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordNaptrResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dtc_record_naptr.test_use_ttl"
	var v dtc.DtcRecordNaptr
	serverName := acctest.RandomNameWithPrefix("dtc-server")
	serverIp := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordNaptrUseTtl(serverName, serverIp, 2, 5, "example.com", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordNaptrUseTtl(serverName, serverIp, 2, 5, "example.com", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordNaptrExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
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

func testAccDtcRecordNaptrBasicConfig(serverName, serverIP string, order, preference int, replacement string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_naptr" "test" {
		dtc_server = nios_dtc_server.test.name
		order      = %d
		preference = %d
		replacement = %q
	}
	`, order, preference, replacement)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, serverIP), config}, "\n")
}

func testAccDtcRecordNaptrComment(serverName, serverIP string, order, preference int, replacement, comment string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_naptr" "test_comment" {
		dtc_server = nios_dtc_server.test.name
		order      = %d
		preference = %d
		replacement = %q
		comment = %q
	}
	`, order, preference, replacement, comment)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, serverIP), config}, "\n")
}

func testAccDtcRecordNaptrDisable(serverName, serverIP string, order, preference int, replacement string, disable bool) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_naptr" "test_disable" {
		dtc_server = nios_dtc_server.test.name
		order      = %d
		preference = %d
		replacement = %q
		disable = %t
	}
	`, order, preference, replacement, disable)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, serverIP), config}, "\n")
}

func testAccDtcRecordNaptrDtcServer(serverName, serverIP string, order, preference int, replacement string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_naptr" "test_dtc_server" {
    	dtc_server = nios_dtc_server.test.name
		order      = %d
		preference = %d
		replacement = %q
	}
	`, order, preference, replacement)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, serverIP), config}, "\n")
}

func testAccDtcRecordNaptrFlags(serverName, serverIP string, order, preference int, replacement, flags string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_naptr" "test_flags" {
		dtc_server = nios_dtc_server.test.name
		order      = %d
		preference = %d
		replacement = %q
		flags = %q
	}
	`, order, preference, replacement, flags)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, serverIP), config}, "\n")
}
func testAccDtcRecordNaptrOrder(serverName, serverIP string, order, preference int, replacement string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_naptr" "test_order" {
		dtc_server = nios_dtc_server.test.name
		order      = %d
		preference = %d
		replacement = %q
	}
	`, order, preference, replacement)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, serverIP), config}, "\n")
}

func testAccDtcRecordNaptrPreference(serverName, serverIP string, order, preference int, replacement string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_naptr" "test_preference" {
		dtc_server = nios_dtc_server.test.name
		order      = %d
		preference = %d
		replacement = %q
	}
	`, order, preference, replacement)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, serverIP), config}, "\n")
}

func testAccDtcRecordNaptrRegexp(serverName, serverIP string, order, preference int, replacement, regexp string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_naptr" "test_regexp" {
		dtc_server = nios_dtc_server.test.name
		order      = %d
		preference = %d
		replacement = %q
		regexp = %q
	}
	`, order, preference, replacement, regexp)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, serverIP), config}, "\n")
}

func testAccDtcRecordNaptrReplacement(serverName, serverIP string, order, preference int, replacement string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_naptr" "test_replacement" {
		dtc_server = nios_dtc_server.test.name
		order      = %d
		preference = %d
		replacement = %q
	}
	`, order, preference, replacement)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, serverIP), config}, "\n")
}

func testAccDtcRecordNaptrServices(serverName, serverIP string, order, preference int, replacement, services string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_naptr" "test_services" {
		dtc_server = nios_dtc_server.test.name
		order      = %d
		preference = %d
		replacement = %q
		services = %q
	}
	`, order, preference, replacement, services)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, serverIP), config}, "\n")
}

func testAccDtcRecordNaptrTtl(serverName, serverIP string, order, preference int, replacement, ttl string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_naptr" "test_ttl" {
		dtc_server = nios_dtc_server.test.name
		order      = %d
		preference = %d
		replacement = %q
		ttl = %q
		use_ttl = false
	}
	`, order, preference, replacement, ttl)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, serverIP), config}, "\n")
}

func testAccDtcRecordNaptrUseTtl(serverName, serverIP string, order, preference int, replacement string, useTtl bool) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_naptr" "test_use_ttl" {
		dtc_server = nios_dtc_server.test.name
		order      = %d
		preference = %d
		replacement = %q
		use_ttl = %t
		ttl = 30
	}
	`, order, preference, replacement, useTtl)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, serverIP), config}, "\n")
}
