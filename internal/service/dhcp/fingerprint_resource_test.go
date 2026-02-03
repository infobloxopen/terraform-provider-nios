package dhcp_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForFingerprint = "comment,device_class,disable,extattrs,ipv6_option_sequence,name,option_sequence,type,vendor_id"

func TestAccFingerprintResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_fingerprint.test"
	var v dhcp.Fingerprint
	name := acctest.RandomNameWithPrefix("fingerprint")
	optionSequence := []string{"1,2,3,4,5,99"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFingerprintBasicConfig("Windows OS", name, optionSequence),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_class", "Windows OS"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "type", "CUSTOM"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFingerprintResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_fingerprint.test"
	var v dhcp.Fingerprint
	name := acctest.RandomNameWithPrefix("fingerprint")
	optionSequence := []string{"1,2,3,4,5,6,99"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFingerprintDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFingerprintBasicConfig("Windows OS", name, optionSequence),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFingerprintExists(context.Background(), resourceName, &v),
					testAccCheckFingerprintDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccFingerprintResource_Import(t *testing.T) {
	var resourceName = "nios_dhcp_fingerprint.test"
	var v dhcp.Fingerprint
	name := acctest.RandomNameWithPrefix("fingerprint")
	optionSequence := []string{"1,2,3,4,5,6,7,99"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFingerprintBasicConfig("Windows OS", name, optionSequence),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFingerprintExists(context.Background(), resourceName, &v),
				),
			},
			// Import with PlanOnly to detect differences
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccFingerprintImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "ref",
				PlanOnly:                             true,
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccFingerprintImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFingerprintResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_fingerprint.test_comment"
	var v dhcp.Fingerprint
	name := acctest.RandomNameWithPrefix("fingerprint")
	optionSequence := []string{"1,2,3,4,5,6,7,8,99"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFingerprintComment("Windows OS", name, "Comment for the object", optionSequence),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccFingerprintComment("Windows OS", name, "Updated comment for the object", optionSequence),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the object"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFingerprintResource_DeviceClass(t *testing.T) {
	var resourceName = "nios_dhcp_fingerprint.test_device_class"
	var v dhcp.Fingerprint
	name := acctest.RandomNameWithPrefix("fingerprint")
	optionSequence := []string{"1,2,3,4,5,6,7,8,9,99"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFingerprintDeviceClass("Windows OS", name, optionSequence),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_class", "Windows OS"),
				),
			},
			// Update and Read
			{
				Config: testAccFingerprintDeviceClass("Mac OS", name, optionSequence),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "device_class", "Mac OS"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFingerprintResource_Disable(t *testing.T) {
	var resourceName = "nios_dhcp_fingerprint.test_disable"
	var v dhcp.Fingerprint
	name := acctest.RandomNameWithPrefix("fingerprint")
	optionSequence := []string{"1,2,3,4,5,6,7,8,9,10,99"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFingerprintDisable("Windows OS", name, "true", optionSequence),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFingerprintDisable("Windows OS", name, "false", optionSequence),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFingerprintResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dhcp_fingerprint.test_extattrs"
	var v dhcp.Fingerprint
	name := acctest.RandomNameWithPrefix("fingerprint")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()
	optionSequence := []string{"1,2,3,4,5,6,7,8,9,10,99,100"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFingerprintExtAttrs("Windows OS", name, map[string]string{
					"Site": extAttrValue1,
				}, optionSequence),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccFingerprintExtAttrs("Windows OS", name, map[string]string{
					"Site": extAttrValue2,
				}, optionSequence),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFingerprintResource_Ipv6OptionSequence(t *testing.T) {
	var resourceName = "nios_dhcp_fingerprint.test_ipv6_option_sequence"
	var v dhcp.Fingerprint
	name := acctest.RandomNameWithPrefix("fingerprint")
	ipv6OptionSequenceVal := []string{
		"1,2,3,4,5,6,7,8,9,10,99,100",
		"1,2,3,4,5,6,7,8,9,10,99,100,101,102",
		"1,2,3,4,5,6,7,8,9,10,99,100,101,102,103",
	}
	ipv6OptionSequenceValUpdated := []string{
		"1,2,3,4,5,6,7,8,9,10,99,100",
		"1,2,3,4,5,6,7,8,9,10,99,100,101,102",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFingerprintIpv6OptionSequence("Windows OS", name, ipv6OptionSequenceVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_option_sequence.0", "1,2,3,4,5,6,7,8,9,10,99,100"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_option_sequence.1", "1,2,3,4,5,6,7,8,9,10,99,100,101,102"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_option_sequence.2", "1,2,3,4,5,6,7,8,9,10,99,100,101,102,103"),
				),
			},
			// Update and Read
			{
				Config: testAccFingerprintIpv6OptionSequence("Windows OS", name, ipv6OptionSequenceValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6_option_sequence.0", "1,2,3,4,5,6,7,8,9,10,99,100"),
					resource.TestCheckResourceAttr(resourceName, "ipv6_option_sequence.1", "1,2,3,4,5,6,7,8,9,10,99,100,101,102"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFingerprintResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_fingerprint.test_name"
	var v dhcp.Fingerprint
	name := acctest.RandomNameWithPrefix("fingerprint")
	nameUpdated := acctest.RandomNameWithPrefix("fingerprint")
	optionSequence := []string{"1,2,3,4,5,6,7,8,9,10,11,99,100"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFingerprintName("Windows OS", name, optionSequence),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccFingerprintName("Windows OS", nameUpdated, optionSequence),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFingerprintResource_OptionSequence(t *testing.T) {
	var resourceName = "nios_dhcp_fingerprint.test_option_sequence"
	var v dhcp.Fingerprint
	name := acctest.RandomNameWithPrefix("fingerprint")
	optionSequenceVal := []string{
		"1,2,3,5,6,7,8,9,10,99,100",
		"1,2,3,5,6,7,8,9,10,99,100,101,102",
		"1,2,3,5,6,7,8,9,10,99,100,101,102,103",
	}
	optionSequenceValUpdated := []string{
		"1,2,3,5,6,7,8,9,10,99,100",
		"1,2,3,5,6,7,8,9,10,99,100,101,102",
		"1,2,3,5,6,7,8,9,10,99,100,101,102,103,104",
		"1,2,3,5,6,7,8,9,10,99,100,101,102,103,104,105",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFingerprintOptionSequence("Windows OS", name, optionSequenceVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_sequence.0", "1,2,3,5,6,7,8,9,10,99,100"),
					resource.TestCheckResourceAttr(resourceName, "option_sequence.1", "1,2,3,5,6,7,8,9,10,99,100,101,102"),
					resource.TestCheckResourceAttr(resourceName, "option_sequence.2", "1,2,3,5,6,7,8,9,10,99,100,101,102,103"),
				),
			},
			// Update and Read
			{
				Config: testAccFingerprintOptionSequence("Windows OS", name, optionSequenceValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_sequence.0", "1,2,3,5,6,7,8,9,10,99,100"),
					resource.TestCheckResourceAttr(resourceName, "option_sequence.1", "1,2,3,5,6,7,8,9,10,99,100,101,102"),
					resource.TestCheckResourceAttr(resourceName, "option_sequence.2", "1,2,3,5,6,7,8,9,10,99,100,101,102,103,104"),
					resource.TestCheckResourceAttr(resourceName, "option_sequence.3", "1,2,3,5,6,7,8,9,10,99,100,101,102,103,104,105"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFingerprintResource_Type(t *testing.T) {
	var resourceName = "nios_dhcp_fingerprint.test_type"
	var v dhcp.Fingerprint
	name := acctest.RandomNameWithPrefix("fingerprint")
	optionSequence := []string{"1,2,3,4,5,6,7,8,9,10,11,99,100,199"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFingerprintType("Windows OS", name, optionSequence, "CUSTOM"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "CUSTOM"),
				),
			},
			// Cannot be Updated as Type is Immutable
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFingerprintResource_VendorId(t *testing.T) {
	var resourceName = "nios_dhcp_fingerprint.test_vendor_id"
	var v dhcp.Fingerprint
	vendorIdVal := []string{"vendor-id-1", "vendor-id-2"}
	vendorIdValUpdated := []string{"vendor-id-1", "vendor-id-2-updated"}
	name := acctest.RandomNameWithPrefix("fingerprint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFingerprintVendorId("Windows OS", name, vendorIdVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_id.0", "vendor-id-1"),
					resource.TestCheckResourceAttr(resourceName, "vendor_id.1", "vendor-id-2"),
				),
			},
			// Update and Read
			{
				Config: testAccFingerprintVendorId("Windows OS", name, vendorIdValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFingerprintExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_id.0", "vendor-id-1"),
					resource.TestCheckResourceAttr(resourceName, "vendor_id.1", "vendor-id-2-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckFingerprintExists(ctx context.Context, resourceName string, v *dhcp.Fingerprint) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			FingerprintAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForFingerprint).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetFingerprintResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetFingerprintResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckFingerprintDestroy(ctx context.Context, v *dhcp.Fingerprint) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			FingerprintAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForFingerprint).
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

func testAccCheckFingerprintDisappears(ctx context.Context, v *dhcp.Fingerprint) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			FingerprintAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccFingerprintImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccFingerprintBasicConfig(deviceClass, name string, optionSequence []string) string {
	optionSequenceStr := utils.ConvertStringSliceToHCL(optionSequence)
	return fmt.Sprintf(`
resource "nios_dhcp_fingerprint" "test" {
    device_class = %q
    name = %q
    option_sequence = %s
}
`, deviceClass, name, optionSequenceStr)
}

func testAccFingerprintComment(deviceClass string, name string, comment string, optionSequence []string) string {
	optionSequenceStr := utils.ConvertStringSliceToHCL(optionSequence)
	return fmt.Sprintf(`
resource "nios_dhcp_fingerprint" "test_comment" {
    device_class = %q
    name = %q
    comment = %q
	option_sequence = %s
}
`, deviceClass, name, comment, optionSequenceStr)
}

func testAccFingerprintDeviceClass(deviceClass string, name string, optionSequence []string) string {
	optionSequenceStr := utils.ConvertStringSliceToHCL(optionSequence)
	return fmt.Sprintf(`
resource "nios_dhcp_fingerprint" "test_device_class" {
    device_class = %q
    name = %q
    option_sequence = %s
}
`, deviceClass, name, optionSequenceStr)
}

func testAccFingerprintDisable(deviceClass string, name string, disable string, optionSequence []string) string {
	optionSequenceStr := utils.ConvertStringSliceToHCL(optionSequence)
	return fmt.Sprintf(`
resource "nios_dhcp_fingerprint" "test_disable" {
    device_class = %q
    name = %q
    disable = %q
	option_sequence = %s
}
`, deviceClass, name, disable, optionSequenceStr)
}

func testAccFingerprintExtAttrs(deviceClass string, name string, extAttrs map[string]string, optionSequence []string) string {
	optionSequenceStr := utils.ConvertStringSliceToHCL(optionSequence)
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	return fmt.Sprintf(`
resource "nios_dhcp_fingerprint" "test_extattrs" {
    device_class = %q
    name = %q
    extattrs = %s
	option_sequence = %s
}
`, deviceClass, name, extAttrsStr, optionSequenceStr)
}

func testAccFingerprintIpv6OptionSequence(deviceClass string, name string, ipv6OptionSequence []string) string {
	ipv6OptionSequenceStr := utils.ConvertStringSliceToHCL(ipv6OptionSequence)
	return fmt.Sprintf(`
resource "nios_dhcp_fingerprint" "test_ipv6_option_sequence" {
    device_class = %q
    name = %q
    ipv6_option_sequence = %s
}
`, deviceClass, name, ipv6OptionSequenceStr)
}

func testAccFingerprintName(deviceClass string, name string, optionSequence []string) string {
	optionSequenceStr := utils.ConvertStringSliceToHCL(optionSequence)
	return fmt.Sprintf(`
resource "nios_dhcp_fingerprint" "test_name" {
    device_class = %q
    name = %q
	option_sequence = %s
}
`, deviceClass, name, optionSequenceStr)
}

func testAccFingerprintOptionSequence(deviceClass string, name string, optionSequence []string) string {
	optionSequenceStr := utils.ConvertStringSliceToHCL(optionSequence)
	return fmt.Sprintf(`
resource "nios_dhcp_fingerprint" "test_option_sequence" {
    device_class = %q
    name = %q
    option_sequence = %s
}
`, deviceClass, name, optionSequenceStr)
}

func testAccFingerprintType(deviceClass string, name string, optionSequence []string, fingerprintType string) string {
	optionSequenceStr := utils.ConvertStringSliceToHCL(optionSequence)
	return fmt.Sprintf(`
resource "nios_dhcp_fingerprint" "test_type" {
    device_class = %q
    name = %q
    type = %q
	option_sequence = %s
}
`, deviceClass, name, fingerprintType, optionSequenceStr)
}

func testAccFingerprintVendorId(deviceClass string, name string, vendorId []string) string {
	vendorIdStr := utils.ConvertStringSliceToHCL(vendorId)
	return fmt.Sprintf(`
resource "nios_dhcp_fingerprint" "test_vendor_id" {
    device_class = %q
    name = %q
    vendor_id = %s
}
`, deviceClass, name, vendorIdStr)
}
