package parentalcontrol_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/parentalcontrol"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForParentalcontrolAvp = "comment,domain_types,is_restricted,name,type,user_defined,value_type,vendor_id,vendor_type"

func TestAccParentalcontrolAvpResource_basic(t *testing.T) {
	var resourceName = "nios_parentalcontrol_avp.test"
	var v parentalcontrol.ParentalcontrolAvp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolAvpBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolAvpResource_disappears(t *testing.T) {
	resourceName := "nios_parentalcontrol_avp.test"
	var v parentalcontrol.ParentalcontrolAvp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckParentalcontrolAvpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccParentalcontrolAvpBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					testAccCheckParentalcontrolAvpDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccParentalcontrolAvpResource_Ref(t *testing.T) {
	var resourceName = "nios_parentalcontrol_avp.test_ref"
	var v parentalcontrol.ParentalcontrolAvp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolAvpRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolAvpRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolAvpResource_Comment(t *testing.T) {
	var resourceName = "nios_parentalcontrol_avp.test_comment"
	var v parentalcontrol.ParentalcontrolAvp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolAvpComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolAvpComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolAvpResource_DomainTypes(t *testing.T) {
	var resourceName = "nios_parentalcontrol_avp.test_domain_types"
	var v parentalcontrol.ParentalcontrolAvp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolAvpDomainTypes("DOMAIN_TYPES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_types", "DOMAIN_TYPES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolAvpDomainTypes("DOMAIN_TYPES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_types", "DOMAIN_TYPES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolAvpResource_IsRestricted(t *testing.T) {
	var resourceName = "nios_parentalcontrol_avp.test_is_restricted"
	var v parentalcontrol.ParentalcontrolAvp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolAvpIsRestricted("IS_RESTRICTED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "is_restricted", "IS_RESTRICTED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolAvpIsRestricted("IS_RESTRICTED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "is_restricted", "IS_RESTRICTED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolAvpResource_Name(t *testing.T) {
	var resourceName = "nios_parentalcontrol_avp.test_name"
	var v parentalcontrol.ParentalcontrolAvp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolAvpName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolAvpName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolAvpResource_Type(t *testing.T) {
	var resourceName = "nios_parentalcontrol_avp.test_type"
	var v parentalcontrol.ParentalcontrolAvp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolAvpType("TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolAvpType("TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolAvpResource_ValueType(t *testing.T) {
	var resourceName = "nios_parentalcontrol_avp.test_value_type"
	var v parentalcontrol.ParentalcontrolAvp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolAvpValueType("VALUE_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "value_type", "VALUE_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolAvpValueType("VALUE_TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "value_type", "VALUE_TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolAvpResource_VendorId(t *testing.T) {
	var resourceName = "nios_parentalcontrol_avp.test_vendor_id"
	var v parentalcontrol.ParentalcontrolAvp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolAvpVendorId("VENDOR_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_id", "VENDOR_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolAvpVendorId("VENDOR_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_id", "VENDOR_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolAvpResource_VendorType(t *testing.T) {
	var resourceName = "nios_parentalcontrol_avp.test_vendor_type"
	var v parentalcontrol.ParentalcontrolAvp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolAvpVendorType("VENDOR_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_type", "VENDOR_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolAvpVendorType("VENDOR_TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_type", "VENDOR_TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckParentalcontrolAvpExists(ctx context.Context, resourceName string, v *parentalcontrol.ParentalcontrolAvp) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.ParentalControlAPI.
			ParentalcontrolAvpAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForParentalcontrolAvp).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetParentalcontrolAvpResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetParentalcontrolAvpResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckParentalcontrolAvpDestroy(ctx context.Context, v *parentalcontrol.ParentalcontrolAvp) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.ParentalControlAPI.
			ParentalcontrolAvpAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForParentalcontrolAvp).
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

func testAccCheckParentalcontrolAvpDisappears(ctx context.Context, v *parentalcontrol.ParentalcontrolAvp) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.ParentalControlAPI.
			ParentalcontrolAvpAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccParentalcontrolAvpBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return `
resource "nios_parentalcontrol_avp" "test" {
}
`
}

func testAccParentalcontrolAvpRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_avp" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccParentalcontrolAvpComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_avp" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccParentalcontrolAvpDomainTypes(domainTypes string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_avp" "test_domain_types" {
    domain_types = %q
}
`, domainTypes)
}

func testAccParentalcontrolAvpIsRestricted(isRestricted string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_avp" "test_is_restricted" {
    is_restricted = %q
}
`, isRestricted)
}

func testAccParentalcontrolAvpName(name string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_avp" "test_name" {
    name = %q
}
`, name)
}

func testAccParentalcontrolAvpType(type_ string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_avp" "test_type" {
    type = %q
}
`, type_)
}

func testAccParentalcontrolAvpValueType(valueType string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_avp" "test_value_type" {
    value_type = %q
}
`, valueType)
}

func testAccParentalcontrolAvpVendorId(vendorId string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_avp" "test_vendor_id" {
    vendor_id = %q
}
`, vendorId)
}

func testAccParentalcontrolAvpVendorType(vendorType string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_avp" "test_vendor_type" {
    vendor_type = %q
}
`, vendorType)
}
