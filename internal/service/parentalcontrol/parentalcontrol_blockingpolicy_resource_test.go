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
	name := acctest.RandomNameWithPrefix("blockingpolicy-")
	value := "0000000abcd00000000000000000aaa1"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolBlockingpolicyBasicConfig(name, value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolBlockingpolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "value", value),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolBlockingpolicyResource_disappears(t *testing.T) {
	resourceName := "nios_parentalcontrol_blockingpolicy.test"
	var v parentalcontrol.ParentalcontrolBlockingpolicy
	name := acctest.RandomNameWithPrefix("blockingpolicy-")
	value := "0000000abcd00000000000000000bbb1"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckParentalcontrolBlockingpolicyDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccParentalcontrolBlockingpolicyBasicConfig(name, value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolBlockingpolicyExists(context.Background(), resourceName, &v),
					testAccCheckParentalcontrolBlockingpolicyDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccParentalcontrolBlockingpolicyResource_Name(t *testing.T) {
	var resourceName = "nios_parentalcontrol_blockingpolicy.test_name"
	var v parentalcontrol.ParentalcontrolBlockingpolicy
	name1 := acctest.RandomNameWithPrefix("blockingpolicy-")
	name2 := acctest.RandomNameWithPrefix("blockingpolicy-")
	value := "0000000abcd000000000000000aba001"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolBlockingpolicyName(name1, value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolBlockingpolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolBlockingpolicyName(name2, value),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolBlockingpolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolBlockingpolicyResource_Value(t *testing.T) {
	var resourceName = "nios_parentalcontrol_blockingpolicy.test_value"
	var v parentalcontrol.ParentalcontrolBlockingpolicy
	name := acctest.RandomNameWithPrefix("blockingpolicy-")
	value1 := "0000000abcd000000000000000078901"
	value2 := "0000000abcd0000000045600000efef1"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolBlockingpolicyValue(name, value1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolBlockingpolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "value", value1),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolBlockingpolicyValue(name, value2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolBlockingpolicyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "value", value2),
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

func testAccParentalcontrolBlockingpolicyBasicConfig(name, value string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_blockingpolicy" "test" {
    name = %q
    value = %q
}
`, name, value)
}

func testAccParentalcontrolBlockingpolicyName(name, value string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_blockingpolicy" "test_name" {
    name = %q
    value = %q
}
`, name, value)
}

func testAccParentalcontrolBlockingpolicyValue(name, value string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_blockingpolicy" "test_value" {
    name = %q
    value = %q
}
`, name, value)
}
