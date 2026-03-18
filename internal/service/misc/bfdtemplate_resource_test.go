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

var readableAttributesForBfdtemplate = "detection_multiplier,min_rx_interval,min_tx_interval,name"

func TestAccBfdtemplateResource_basic(t *testing.T) {
	var resourceName = "nios_misc_bfdtemplate.test"
	var v misc.Bfdtemplate
	name := acctest.RandomNameWithPrefix("tf_test_bfd_")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccBfdtemplateBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBfdtemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "detection_multiplier", "3"),
					resource.TestCheckResourceAttr(resourceName, "min_rx_interval", "100"),
					resource.TestCheckResourceAttr(resourceName, "min_tx_interval", "100"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccBfdtemplateResource_disappears(t *testing.T) {
	resourceName := "nios_misc_bfdtemplate.test"
	var v misc.Bfdtemplate
	name := acctest.RandomNameWithPrefix("tf_test_bfd_")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBfdtemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccBfdtemplateBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBfdtemplateExists(context.Background(), resourceName, &v),
					testAccCheckBfdtemplateDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccBfdtemplateResource_MinRxInterval(t *testing.T) {
	var resourceName = "nios_misc_bfdtemplate.test_min_rx_interval"
	var v misc.Bfdtemplate
	name := acctest.RandomNameWithPrefix("tf_test_bfd_")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccBfdtemplateMinRxInterval(name, "200"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBfdtemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "min_rx_interval", "200"),
				),
			},
			// Update and Read
			{
				Config: testAccBfdtemplateMinRxInterval(name, "300"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBfdtemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "min_rx_interval", "300"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccBfdtemplateResource_MinTxInterval(t *testing.T) {
	var resourceName = "nios_misc_bfdtemplate.test_min_tx_interval"
	var v misc.Bfdtemplate
	name := acctest.RandomNameWithPrefix("tf_test_bfd_")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccBfdtemplateMinTxInterval(name, "200"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBfdtemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "min_tx_interval", "200"),
				),
			},
			// Update and Read
			{
				Config: testAccBfdtemplateMinTxInterval(name, "300"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBfdtemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "min_tx_interval", "300"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccBfdtemplateResource_Name(t *testing.T) {
	var resourceName = "nios_misc_bfdtemplate.test_name"
	var v misc.Bfdtemplate
	name := acctest.RandomNameWithPrefix("tf_test_bfd_")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccBfdtemplateName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBfdtemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccBfdtemplateName(name + "updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBfdtemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name+"updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckBfdtemplateExists(ctx context.Context, resourceName string, v *misc.Bfdtemplate) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.MiscAPI.
			BfdtemplateAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForBfdtemplate).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetBfdtemplateResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetBfdtemplateResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckBfdtemplateDestroy(ctx context.Context, v *misc.Bfdtemplate) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.MiscAPI.
			BfdtemplateAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForBfdtemplate).
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

func testAccCheckBfdtemplateDisappears(ctx context.Context, v *misc.Bfdtemplate) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.MiscAPI.
			BfdtemplateAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccBfdtemplateBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_misc_bfdtemplate" "test" {
    name = %q
}
`, name)
}

func testAccBfdtemplateMinRxInterval(name, minRxInterval string) string {
	return fmt.Sprintf(`
resource "nios_misc_bfdtemplate" "test_min_rx_interval" {
    name = %q
    min_rx_interval = %q
}
`, name, minRxInterval)
}

func testAccBfdtemplateMinTxInterval(name, minTxInterval string) string {
	return fmt.Sprintf(`
resource "nios_misc_bfdtemplate" "test_min_tx_interval" {
    name = %q
    min_tx_interval = %q
}
`, name, minTxInterval)
}

func testAccBfdtemplateName(name string) string {
	return fmt.Sprintf(`
resource "nios_misc_bfdtemplate" "test_name" {
    name = %q
}
`, name)
}
