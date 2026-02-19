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
	name := acctest.RandomNameWithPrefix("parentalcontrol-avp")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolAvpBasicConfig(name, "BYTE", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "is_restricted", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolAvpResource_disappears(t *testing.T) {
	resourceName := "nios_parentalcontrol_avp.test"
	var v parentalcontrol.ParentalcontrolAvp
	name := acctest.RandomNameWithPrefix("parentalcontrol-avp")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckParentalcontrolAvpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccParentalcontrolAvpBasicConfig(name, "BYTE", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					testAccCheckParentalcontrolAvpDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccParentalcontrolAvpResource_Comment(t *testing.T) {
	var resourceName = "nios_parentalcontrol_avp.test_comment"
	var v parentalcontrol.ParentalcontrolAvp
	name := acctest.RandomNameWithPrefix("parentalcontrol-avp")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolAvpComment(name, "BYTE", "Parental control AVP", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Parental control AVP"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolAvpComment(name, "BYTE", "Parental control AVP updated", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Parental control AVP updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolAvpResource_DomainTypes(t *testing.T) {
	var resourceName = "nios_parentalcontrol_avp.test_domain_types"
	var v parentalcontrol.ParentalcontrolAvp
	name := acctest.RandomNameWithPrefix("parentalcontrol-avp")
	domainTypes1 := []string{"ANCILLARY"}
	domainTypes2 := []string{"LOCAL_ID", "IP_SPACE_DIS"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolAvpDomainTypes(name, "BYTE", domainTypes1, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_types.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "domain_types.0", "ANCILLARY"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolAvpDomainTypes(name, "BYTE", domainTypes2, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_types.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "domain_types.0", "LOCAL_ID"),
					resource.TestCheckResourceAttr(resourceName, "domain_types.1", "IP_SPACE_DIS"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolAvpResource_IsRestricted(t *testing.T) {
	var resourceName = "nios_parentalcontrol_avp.test_is_restricted"
	var v parentalcontrol.ParentalcontrolAvp
	name := acctest.RandomNameWithPrefix("parentalcontrol-avp")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolAvpIsRestricted(name, "BYTE", "true", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "is_restricted", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolAvpIsRestricted(name, "BYTE", "false", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "is_restricted", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolAvpResource_Name(t *testing.T) {
	var resourceName = "nios_parentalcontrol_avp.test_name"
	var v parentalcontrol.ParentalcontrolAvp
	name1 := acctest.RandomNameWithPrefix("parentalcontrol-avp")
	name2 := acctest.RandomNameWithPrefix("parentalcontrol-avp")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolAvpName(name1, "BYTE", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolAvpName(name2, "BYTE", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolAvpResource_Type(t *testing.T) {
	var resourceName = "nios_parentalcontrol_avp.test_type"
	var v parentalcontrol.ParentalcontrolAvp
	name := acctest.RandomNameWithPrefix("parentalcontrol-avp")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolAvpType(name, "BYTE", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolAvpType(name, "BYTE", 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolAvpResource_ValueType(t *testing.T) {
	var resourceName = "nios_parentalcontrol_avp.test_value_type"
	var v parentalcontrol.ParentalcontrolAvp
	name := acctest.RandomNameWithPrefix("parentalcontrol-avp")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolAvpValueType(name, "BYTE", 212),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "value_type", "BYTE"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolAvpValueType(name, "INTEGER", 213),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "value_type", "INTEGER"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolAvpResource_VendorId(t *testing.T) {
	var resourceName = "nios_parentalcontrol_avp.test_vendor_id"
	var v parentalcontrol.ParentalcontrolAvp
	name := acctest.RandomNameWithPrefix("parentalcontrol-avp")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolAvpVendorId(name, "BYTE", 26, 123, 122),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_id", "123"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolAvpVendorId(name, "BYTE", 26, 232, 122),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_id", "232"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolAvpResource_VendorType(t *testing.T) {
	var resourceName = "nios_parentalcontrol_avp.test_vendor_type"
	var v parentalcontrol.ParentalcontrolAvp
	name := acctest.RandomNameWithPrefix("parentalcontrol-avp")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolAvpVendorType(name, "BYTE", 26, 123, 122),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_type", "122"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolAvpVendorType(name, "BYTE", 26, 123, 132),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolAvpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_type", "132"),
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

func testAccParentalcontrolAvpBasicConfig(name, valueType string, type_ int32) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_avp" "test" {
    name = %q
	value_type = %q
	type = %d
}
`, name, valueType, type_)
}

func testAccParentalcontrolAvpComment(name, valueType, comment string, type_ int32) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_avp" "test_comment" {
    name = %q
	value_type = %q
	type = %d
    comment = %q
}
`, name, valueType, type_, comment)
}

func testAccParentalcontrolAvpDomainTypes(name, valueType string, domainTypes []string, type_ int32) string {
	domainTypesStr := utils.ConvertStringSliceToHCL(domainTypes)
	return fmt.Sprintf(`
resource "nios_parentalcontrol_avp" "test_domain_types" {
    name = %q
	value_type = %q
	type = %d
    domain_types = %s
}
`, name, valueType, type_, domainTypesStr)
}

func testAccParentalcontrolAvpIsRestricted(name, valueType, isRestricted string, type_ int32) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_avp" "test_is_restricted" {
    name = %q
	value_type = %q
	type = %d
    is_restricted = %q
}
`, name, valueType, type_, isRestricted)
}

func testAccParentalcontrolAvpName(name, valueType string, type_ int32) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_avp" "test_name" {
    name = %q
	value_type = %q
	type = %d
}
`, name, valueType, type_)
}

func testAccParentalcontrolAvpType(name, valueType string, type_ int32) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_avp" "test_type" {
    name = %q
	value_type = %q
	type = %d
}
`, name, valueType, type_)
}

func testAccParentalcontrolAvpValueType(name, valueType string, type_ int32) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_avp" "test_value_type" {
    name = %q
	value_type = %q
	type = %d
}
`, name, valueType, type_)
}

func testAccParentalcontrolAvpVendorId(name, valueType string, type_, vendorId, vendorType int32) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_avp" "test_vendor_id" {
    name = %q
	value_type = %q
	type = %d
    vendor_id = %d
    vendor_type = %d
}
`, name, valueType, type_, vendorId, vendorType)
}

func testAccParentalcontrolAvpVendorType(name, valueType string, type_, vendorId, vendorType int32) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_avp" "test_vendor_type" {
    name = %q
	value_type = %q
	type = %d
    vendor_id = %d
    vendor_type = %d
}
`, name, valueType, type_, vendorId, vendorType)
}
