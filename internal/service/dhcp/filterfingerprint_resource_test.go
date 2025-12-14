package dhcp_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

/*
// Manage dhcp Filterfingerprint with Basic Fields
resource "nios_dhcp_filterfingerprint" "dhcp_filterfingerprint_basic" {
    fingerprint = "FINGERPRINT_REPLACE_ME"
    name = name
}

// Manage dhcp Filterfingerprint with Additional Fields
resource "nios_dhcp_filterfingerprint" "dhcp_filterfingerprint_with_additional_fields" {
    fingerprint = "FINGERPRINT_REPLACE_ME"
    name = name

// TODO : Add additional optional fields below

    //Extensible Attributes
    extattrs = {
        Site = "location-1"
    }
}
*/

var readableAttributesForFilterfingerprint = "comment,extattrs,fingerprint,name"

func TestAccFilterfingerprintResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_filterfingerprint.test"
	var v dhcp.Filterfingerprint
	name := acctest.RandomNameWithPrefix("filter-fingerprint")
	fingerprint1 := acctest.RandomNameWithPrefix("fingerprint")
	fingerprint2 := acctest.RandomNameWithPrefix("fingerprint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilterfingerprintBasicConfig(fingerprint1, fingerprint2, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterfingerprintExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "fingerprint.0", fingerprint1),
					resource.TestCheckResourceAttr(resourceName, "fingerprint.1", fingerprint2),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilterfingerprintResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_filterfingerprint.test"
	var v dhcp.Filterfingerprint
	name := acctest.RandomNameWithPrefix("filter-fingerprint")
	fingerprint1 := acctest.RandomNameWithPrefix("fingerprint")
	fingerprint2 := acctest.RandomNameWithPrefix("fingerprint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFilterfingerprintDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFilterfingerprintBasicConfig(fingerprint1, fingerprint2, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterfingerprintExists(context.Background(), resourceName, &v),
					testAccCheckFilterfingerprintDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccFilterfingerprintResource_Import(t *testing.T) {
	var resourceName = "nios_dhcp_filterfingerprint.test"
	var v dhcp.Filterfingerprint
	name := acctest.RandomNameWithPrefix("filter-fingerprint")
	fingerprint1 := acctest.RandomNameWithPrefix("fingerprint")
	fingerprint2 := acctest.RandomNameWithPrefix("fingerprint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilterfingerprintBasicConfig(fingerprint1, fingerprint2, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterfingerprintExists(context.Background(), resourceName, &v),
				),
			},
			// Import with PlanOnly to detect differences
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccFilterfingerprintImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "ref",
				PlanOnly:                             true,
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccFilterfingerprintImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilterfingerprintResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_filterfingerprint.test_comment"
	var v dhcp.Filterfingerprint
	name := acctest.RandomNameWithPrefix("filter-fingerprint")
	fingerprint1 := acctest.RandomNameWithPrefix("fingerprint")
	fingerprint2 := acctest.RandomNameWithPrefix("fingerprint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilterfingerprintComment(fingerprint1, fingerprint2, name, "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterfingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccFilterfingerprintComment(fingerprint1, fingerprint2, name, "Updated comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterfingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the object"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilterfingerprintResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dhcp_filterfingerprint.test_extattrs"
	var v dhcp.Filterfingerprint
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()
	name := acctest.RandomNameWithPrefix("filter-fingerprint")
	fingerprint1 := acctest.RandomNameWithPrefix("fingerprint")
	fingerprint2 := acctest.RandomNameWithPrefix("fingerprint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilterfingerprintExtAttrs(fingerprint1, fingerprint2, name, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterfingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccFilterfingerprintExtAttrs(fingerprint1, fingerprint2, name, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterfingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilterfingerprintResource_Fingerprint(t *testing.T) {
	var resourceName = "nios_dhcp_filterfingerprint.test_fingerprint"
	var v dhcp.Filterfingerprint
	name := acctest.RandomNameWithPrefix("filter-fingerprint")
	fingerprint1 := acctest.RandomNameWithPrefix("fingerprint")
	fingerprint2 := acctest.RandomNameWithPrefix("fingerprint")
	fingerprint3 := acctest.RandomNameWithPrefix("fingerprint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilterfingerprintFingerprint(fingerprint1, fingerprint2, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterfingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fingerprint.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "fingerprint.0", fingerprint1),
					resource.TestCheckResourceAttr(resourceName, "fingerprint.1", fingerprint2),
				),
			},
			// Update and Read
			{
				Config: testAccFilterfingerprintFingerprintUpdate(fingerprint3, fingerprint2, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterfingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fingerprint.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "fingerprint.0", fingerprint3),
					resource.TestCheckResourceAttr(resourceName, "fingerprint.1", fingerprint2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilterfingerprintResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_filterfingerprint.test_name"
	var v dhcp.Filterfingerprint
	name := acctest.RandomNameWithPrefix("filter-fingerprint")
	nameUpdated := acctest.RandomNameWithPrefix("filter-fingerprint-updated")
	fingerprint1 := acctest.RandomNameWithPrefix("fingerprint")
	fingerprint2 := acctest.RandomNameWithPrefix("fingerprint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilterfingerprintName(fingerprint1, fingerprint2, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterfingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccFilterfingerprintName(fingerprint1, fingerprint2, nameUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterfingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckFilterfingerprintExists(ctx context.Context, resourceName string, v *dhcp.Filterfingerprint) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			FilterfingerprintAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForFilterfingerprint).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetFilterfingerprintResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetFilterfingerprintResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckFilterfingerprintDestroy(ctx context.Context, v *dhcp.Filterfingerprint) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			FilterfingerprintAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForFilterfingerprint).
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

func testAccCheckFilterfingerprintDisappears(ctx context.Context, v *dhcp.Filterfingerprint) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			FilterfingerprintAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccFilterfingerprintImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.Attributes["ref"] == "" {
			return "", fmt.Errorf("ref is not set")
		}
		return rs.Primary.Attributes["ref"], nil
	}
}

func testAccFilterfingerprintBasicConfig(fingerprint1, fingerprint2, name string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_filterfingerprint" "test" {
    fingerprint = [
		nios_dhcp_fingerprint.test1.name,
		nios_dhcp_fingerprint.test2.name,
	]
    name = %q
}
`, name)
	return strings.Join([]string{testAccBaseWithFingerprint(fingerprint1, fingerprint2), config}, "")
}

func testAccFilterfingerprintComment(fingerprint1, fingerprint2, name string, comment string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_filterfingerprint" "test_comment" {
    fingerprint = [
		nios_dhcp_fingerprint.test1.name,
		nios_dhcp_fingerprint.test2.name,
	]
    name = %q
    comment = %q
}
`, name, comment)
	return strings.Join([]string{testAccBaseWithFingerprint(fingerprint1, fingerprint2), config}, "")
}

func testAccFilterfingerprintExtAttrs(fingerprint1, fingerprint2 string, name string, extAttrs map[string]string) string {
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	config := fmt.Sprintf(`
resource "nios_dhcp_filterfingerprint" "test_extattrs" {
    fingerprint = [
		nios_dhcp_fingerprint.test1.name,
		nios_dhcp_fingerprint.test2.name,
	]
    name = %q
    extattrs = %s
}
`, name, extAttrsStr)
	return strings.Join([]string{testAccBaseWithFingerprint(fingerprint1, fingerprint2), config}, "")
}

func testAccFilterfingerprintFingerprint(fingerprint1, fingerprint2, name string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_filterfingerprint" "test_fingerprint" {
    fingerprint = [
        nios_dhcp_fingerprint.test1.name,
        nios_dhcp_fingerprint.test2.name,
    ]
    name = %q
}
`, name)
	return strings.Join([]string{testAccBaseWithFingerprint(fingerprint1, fingerprint2), config}, "")
}

func testAccFilterfingerprintFingerprintUpdate(fingerprint1, fingerprint2, name string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_filterfingerprint" "test_fingerprint" {
    fingerprint = [
        nios_dhcp_fingerprint.test1.name,
		nios_dhcp_fingerprint.test2.name,
    ]
    name = %q
}
`, name)
	return strings.Join([]string{testAccBaseWithFingerprint(fingerprint1, fingerprint2), config}, "")
}

func testAccFilterfingerprintName(fingerprint1, fingerprint2, name string) string {
	config := fmt.Sprintf(`
resource "nios_dhcp_filterfingerprint" "test_name" {
    fingerprint = [
		nios_dhcp_fingerprint.test1.name,
		nios_dhcp_fingerprint.test2.name,
	]
    name = %q
}
`, name)
	return strings.Join([]string{testAccBaseWithFingerprint(fingerprint1, fingerprint2), config}, "")
}

func testAccBaseWithFingerprint(name1, name2 string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_fingerprint" "test1" {
	name = %q
	device_class = "Windows OS"
	vendor_id = ["MSFT 5.0"]
}
resource "nios_dhcp_fingerprint" "test2" {
	name = %q
	device_class = "Windows OS"
	vendor_id = ["MSFT 5.0"]
}
`, name1, name2)
}
