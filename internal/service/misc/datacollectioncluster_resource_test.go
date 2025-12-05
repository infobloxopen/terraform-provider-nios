package misc_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForDatacollectioncluster = "enable_registration,name"

func TestAccDatacollectionclusterResource_basic(t *testing.T) {
	var resourceName = "nios_misc_datacollectioncluster.test"
	var v misc.Datacollectioncluster

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDatacollectionclusterBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatacollectionclusterExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDatacollectionclusterResource_disappears(t *testing.T) {
	resourceName := "nios_misc_datacollectioncluster.test"
	var v misc.Datacollectioncluster

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDatacollectionclusterDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDatacollectionclusterBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatacollectionclusterExists(context.Background(), resourceName, &v),
					testAccCheckDatacollectionclusterDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDatacollectionclusterResource_Ref(t *testing.T) {
	var resourceName = "nios_misc_datacollectioncluster.test_ref"
	var v misc.Datacollectioncluster

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDatacollectionclusterRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatacollectionclusterExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDatacollectionclusterRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatacollectionclusterExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDatacollectionclusterResource_EnableRegistration(t *testing.T) {
	var resourceName = "nios_misc_datacollectioncluster.test_enable_registration"
	var v misc.Datacollectioncluster

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDatacollectionclusterEnableRegistration("ENABLE_REGISTRATION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatacollectionclusterExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_registration", "ENABLE_REGISTRATION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDatacollectionclusterEnableRegistration("ENABLE_REGISTRATION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatacollectionclusterExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_registration", "ENABLE_REGISTRATION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDatacollectionclusterExists(ctx context.Context, resourceName string, v *misc.Datacollectioncluster) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.MiscAPI.
			DatacollectionclusterAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDatacollectioncluster).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDatacollectionclusterResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDatacollectionclusterResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDatacollectionclusterDestroy(ctx context.Context, v *misc.Datacollectioncluster) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.MiscAPI.
			DatacollectionclusterAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDatacollectioncluster).
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

func testAccCheckDatacollectionclusterDisappears(ctx context.Context, v *misc.Datacollectioncluster) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.MiscAPI.
			DatacollectionclusterAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDatacollectionclusterBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return `
resource "nios_misc_datacollectioncluster" "test" {
}
`
}

func testAccDatacollectionclusterRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_misc_datacollectioncluster" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccDatacollectionclusterEnableRegistration(enableRegistration string) string {
	return fmt.Sprintf(`
resource "nios_misc_datacollectioncluster" "test_enable_registration" {
    enable_registration = %q
}
`, enableRegistration)
}
