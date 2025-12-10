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

var readableAttributesForDtcRecordCname = "auto_created,canonical,comment,disable,dns_canonical,dtc_server,ttl,use_ttl"

func TestAccDtcRecordCnameResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_record_cname.test"
	var v dtc.DtcRecordCname
	name := acctest.RandomNameWithPrefix("dtc-cname")
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordCnameBasicConfig(name, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordCnameResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_record_cname.test"
	var v dtc.DtcRecordCname
	name := acctest.RandomNameWithPrefix("dtc-cname")
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcRecordCnameDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcRecordCnameBasicConfig(name, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordCnameExists(context.Background(), resourceName, &v),
					testAccCheckDtcRecordCnameDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcRecordCnameResource_Canonical(t *testing.T) {
	var resourceName = "nios_dtc_record_cname.test_canonical"
	var v dtc.DtcRecordCname
	name := acctest.RandomNameWithPrefix("dtc-cname")
	nameUpdate := acctest.RandomNameWithPrefix("dtc-cname-upd")
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordCnameCanonical(name, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", name),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordCnameCanonical(nameUpdate, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", nameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordCnameResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_record_cname.test_comment"
	var v dtc.DtcRecordCname
	name := acctest.RandomNameWithPrefix("dtc-cname")
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordCnameComment(name, serverName, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordCnameComment(name, serverName, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordCnameResource_Disable(t *testing.T) {
	var resourceName = "nios_dtc_record_cname.test_disable"
	var v dtc.DtcRecordCname
	name := acctest.RandomNameWithPrefix("dtc-cname")
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordCnameDisable(name, serverName, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordCnameDisable(name, serverName, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordCnameResource_DtcServer(t *testing.T) {
	var resourceName = "nios_dtc_record_cname.test_dtc_server"
	var v dtc.DtcRecordCname
	name := acctest.RandomNameWithPrefix("dtc-cname")
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordCnameDtcServer(name, serverName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dtc_server", serverName),
				),
			},
			// remove Update step as dtc_server is immutable
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordCnameResource_Ttl(t *testing.T) {
	var resourceName = "nios_dtc_record_cname.test_ttl"
	var v dtc.DtcRecordCname
	name := acctest.RandomNameWithPrefix("dtc-cname")
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordCnameTtl(name, serverName, 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "20"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordCnameTtl(name, serverName, 30),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "30"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcRecordCnameResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dtc_record_cname.test_use_ttl"
	var v dtc.DtcRecordCname
	name := acctest.RandomNameWithPrefix("dtc-cname")
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcRecordCnameUseTtl(name, serverName, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcRecordCnameUseTtl(name, serverName, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDtcRecordCnameExists(ctx context.Context, resourceName string, v *dtc.DtcRecordCname) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DTCAPI.
			DtcRecordCnameAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDtcRecordCname).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDtcRecordCnameResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDtcRecordCnameResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDtcRecordCnameDestroy(ctx context.Context, v *dtc.DtcRecordCname) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DTCAPI.
			DtcRecordCnameAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDtcRecordCname).
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

func testAccCheckDtcRecordCnameDisappears(ctx context.Context, v *dtc.DtcRecordCname) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DTCAPI.
			DtcRecordCnameAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDtcRecordCnameBasicConfig(name, serverName string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_cname" "test" {
	canonical = %q
	dtc_server = nios_dtc_server.test.name
}
	`, name)
	return strings.Join([]string{testAccBaseWithDtcServerDisable(serverName, "2.2.2.2"), config}, "")
}

func testAccBaseWithDtcServerDisable(name, host string) string {
	return fmt.Sprintf(`
resource "nios_dtc_server" "test" {
  name = %q
  host = %q
  auto_create_host_record = false
}
`, name, host)
}

func testAccDtcRecordCnameCanonical(canonical, serverName string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_cname" "test_canonical" {
    canonical = %q
    dtc_server = nios_dtc_server.test.name
}
	`, canonical)
	return strings.Join([]string{testAccBaseWithDtcServerDisable(serverName, "2.2.2.2"), config}, "")
}

func testAccDtcRecordCnameComment(canonical, serverName, comment string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_cname" "test_comment" {
	canonical = %q
	dtc_server = nios_dtc_server.test.name
    comment = %q
}
`, canonical, comment)
	return strings.Join([]string{testAccBaseWithDtcServerDisable(serverName, "2.2.2.2"), config}, "")
}

func testAccDtcRecordCnameDisable(canonical, serverName string, disable bool) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_cname" "test_disable" {
	canonical = %q
	dtc_server = nios_dtc_server.test.name
    disable = %t
}
`, canonical, disable)
	return strings.Join([]string{testAccBaseWithDtcServerDisable(serverName, "2.2.2.2"), config}, "")
}

func testAccDtcRecordCnameDtcServer(canonical, serverName string) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_cname" "test_dtc_server" {
	canonical = %q
	dtc_server = nios_dtc_server.test.name
	}
	`, canonical)
	return strings.Join([]string{testAccBaseWithDtcServerDisable(serverName, "2.2.2.2"), config}, "")
}

func testAccDtcRecordCnameTtl(canonical, serverName string, ttl int) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_cname" "test_ttl" {
	canonical = %q
	dtc_server = nios_dtc_server.test.name
    ttl = %d
	use_ttl = true 
}
	`, canonical, ttl)
	return strings.Join([]string{testAccBaseWithDtcServerDisable(serverName, "2.2.2.2"), config}, "")
}

func testAccDtcRecordCnameUseTtl(canonical, serverName string, useTtl bool) string {
	config := fmt.Sprintf(`
	resource "nios_dtc_record_cname" "test_use_ttl" {
	canonical = %q
	dtc_server = nios_dtc_server.test.name
    use_ttl = %t
	ttl = 25
}
	`, canonical, useTtl)
	return strings.Join([]string{testAccBaseWithDtcServerDisable(serverName, "2.2.2.2"), config}, "")
}
