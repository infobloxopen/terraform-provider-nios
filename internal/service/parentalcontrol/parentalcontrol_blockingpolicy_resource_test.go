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

var readableAttributesForParentalcontrolBlockingpolicy = "name,value"

func TestAccParentalcontrolBlockingpolicyResource_basic(t *testing.T) {
	var resourceName = "nios_parentalcontrol_blockingpolicy.test"
	var v parentalcontrol.ParentalcontrolBlockingpolicy

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolBlockingpolicyBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolBlockingpolicyExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolBlockingpolicyResource_disappears(t *testing.T) {
	resourceName := "nios_parentalcontrol_blockingpolicy.test"
	var v parentalcontrol.ParentalcontrolBlockingpolicy

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckParentalcontrolBlockingpolicyDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccParentalcontrolBlockingpolicyBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolBlockingpolicyExists(context.Background(), resourceName, &v),
					testAccCheckParentalcontrolBlockingpolicyDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccParentalcontrolBlockingpolicyResource_Ref(t *testing.T) {
	var resourceName = "nios_parentalcontrol_blockingpolicy.test_ref"
	var v parentalcontrol.ParentalcontrolBlockingpolicy

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolBlockingpolicyRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolBlockingpolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolBlockingpolicyRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolBlockingpolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolBlockingpolicyResource_Name(t *testing.T) {
	var resourceName = "nios_parentalcontrol_blockingpolicy.test_name"
	var v parentalcontrol.ParentalcontrolBlockingpolicy

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolBlockingpolicyName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolBlockingpolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolBlockingpolicyName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolBlockingpolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolBlockingpolicyResource_Value(t *testing.T) {
	var resourceName = "nios_parentalcontrol_blockingpolicy.test_value"
	var v parentalcontrol.ParentalcontrolBlockingpolicy

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolBlockingpolicyValue("VALUE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolBlockingpolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "value", "VALUE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolBlockingpolicyValue("VALUE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolBlockingpolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "value", "VALUE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckParentalcontrolBlockingpolicyExists(ctx context.Context, resourceName string, v *parentalcontrol.ParentalcontrolBlockingpolicy) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.ParentalControlAPI.
			ParentalcontrolBlockingpolicyAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForParentalcontrolBlockingpolicy).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetParentalcontrolBlockingpolicyResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetParentalcontrolBlockingpolicyResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckParentalcontrolBlockingpolicyDestroy(ctx context.Context, v *parentalcontrol.ParentalcontrolBlockingpolicy) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.ParentalControlAPI.
			ParentalcontrolBlockingpolicyAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForParentalcontrolBlockingpolicy).
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

func testAccCheckParentalcontrolBlockingpolicyDisappears(ctx context.Context, v *parentalcontrol.ParentalcontrolBlockingpolicy) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.ParentalControlAPI.
			ParentalcontrolBlockingpolicyAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccParentalcontrolBlockingpolicyBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return `
resource "nios_parentalcontrol_blockingpolicy" "test" {
}
`
}

func testAccParentalcontrolBlockingpolicyRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_blockingpolicy" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccParentalcontrolBlockingpolicyName(name string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_blockingpolicy" "test_name" {
    name = %q
}
`, name)
}

func testAccParentalcontrolBlockingpolicyValue(value string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_blockingpolicy" "test_value" {
    value = %q
}
`, value)
}
