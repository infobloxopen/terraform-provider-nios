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

var readableAttributesForDtcRecordAaaa = "auto_created,comment,disable,dtc_server,ipv6addr,ttl,use_ttl"

func TestAccDtcRecordAaaaResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_record_aaaa.test"
	var v dtc.DtcRecordAaaa
	ipv6Addr := "2001:db8:85a3::8a2e:370:7335"
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordAaaaBasicConfig(ipv6Addr, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", ipv6Addr),
					resource.TestCheckResourceAttr(resourceName, "dtc_server", serverName),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAaaaResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_record_aaaa.test"
	var v dtc.DtcRecordAaaa
	ipv6Addr := "2001:db8:85a3::8a2e:370:7335"
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcRecordAaaaDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcRecordAaaaBasicConfig(ipv6Addr, serverName),
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
	ipv6Addr := "2001:db8:85a3::8a2e:370:7335"
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordAaaaComment(ipv6Addr, serverName, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordAaaaComment(ipv6Addr, serverName, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAaaaResource_Disable(t *testing.T) {
	var resourceName = "nios_dtc_record_aaaa.test_disable"
	var v dtc.DtcRecordAaaa
	ipv6Addr := "2001:db8:85a3::8a2e:370:7335"
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordAaaaDisable(ipv6Addr, serverName, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordAaaaDisable(ipv6Addr, serverName, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAaaaResource_DtcServer(t *testing.T) {
	var resourceName = "nios_dtc_record_aaaa.test_dtc_server"
	var v dtc.DtcRecordAaaa
	ipv6Addr := "2001:db8:85a3::8a2e:370:7335"
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordAaaaDtcServer(ipv6Addr, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dtc_server", serverName),
				),
			},
			// Removed update step as dtc_server is immutable
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAaaaResource_Ipv6addr(t *testing.T) {
	var resourceName = "nios_dtc_record_aaaa.test_ipv6addr"
	var v dtc.DtcRecordAaaa
	ipv6Addr := "2001:db8:85a3::8a2e:370:7335"
	ipv6AddrUpdate := "2001:db8:85a3::8a2e:370:7336"
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordAaaaIpv6addr(ipv6Addr, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", ipv6Addr),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordAaaaIpv6addr(ipv6AddrUpdate, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", ipv6AddrUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAaaaResource_Ttl(t *testing.T) {
	var resourceName = "nios_dtc_record_aaaa.test_ttl"
	var v dtc.DtcRecordAaaa
	ipv6Addr := "2001:db8:85a3::8a2e:370:7335"
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordAaaaTtl(ipv6Addr, serverName, 30),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "30"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordAaaaTtl(ipv6Addr, serverName, 60),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "60"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordAaaaResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dtc_record_aaaa.test_use_ttl"
	var v dtc.DtcRecordAaaa
	ipv6Addr := "2001:db8:85a3::8a2e:370:7335"
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordAaaaUseTtl(ipv6Addr, serverName, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordAaaaUseTtl(ipv6Addr, serverName, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
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

func testAccDtcRecordAaaaBasicConfig(ipv6Addr, serverName string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_aaaa" "test" {
		ipv6addr = %q
		dtc_server = nios_dtc_server.test.name
	}
	`, ipv6Addr)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.2.2.2"), config}, "")
}

func testAccDtcRecordAaaaComment(ipv6Addr, serverName, comment string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_aaaa" "test_comment" {
		ipv6addr = %q
		comment = %q
		dtc_server = nios_dtc_server.test.name
	}
	`, ipv6Addr, comment)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.2.2.2"), config}, "")
}

func testAccDtcRecordAaaaDisable(ipv6Addr, serverName string, disable bool) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_aaaa" "test_disable" {
		ipv6addr = %q
		disable = %t
		dtc_server = nios_dtc_server.test.name
	}
	`, ipv6Addr, disable)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.2.2.2"), config}, "")
}

func testAccDtcRecordAaaaDtcServer(ipv6Addr, serverName string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_aaaa" "test_dtc_server" {
		ipv6addr = %q
		dtc_server = nios_dtc_server.test.name
	}
	`, ipv6Addr)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.2.2.2"), config}, "")
}

func testAccDtcRecordAaaaIpv6addr(ipv6Addr, serverName string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_aaaa" "test_ipv6addr" {
		ipv6addr = %q
		dtc_server = nios_dtc_server.test.name
	}
	`, ipv6Addr)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.2.2.2"), config}, "")
}

func testAccDtcRecordAaaaTtl(ipv6Addr, serverName string, ttl int) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_aaaa" "test_ttl" {
		ipv6addr = %q
		dtc_server = nios_dtc_server.test.name
		ttl = %d
		use_ttl = true
	}
	`, ipv6Addr, ttl)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.2.2.2"), config}, "")
}

func testAccDtcRecordAaaaUseTtl(ipv4addr, serverName string, useTtl bool) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_aaaa" "test_use_ttl" {
		ipv6addr = %q
		dtc_server = nios_dtc_server.test.name
		use_ttl = %t
		ttl = %d
	}
	`, ipv4addr, useTtl, 30)
	return strings.Join([]string{testAccBaseWithDtcServer(serverName, "2.2.2.2"), config}, "")
}

