package ipam_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForBulkhostnametemplate = "is_grid_default,pre_defined,template_format,template_name"

func TestAccBulkhostnametemplateResource_basic(t *testing.T) {
	var resourceName = "nios_ipam_bulkhostnametemplate.test"
	var v ipam.Bulkhostnametemplate
	templateName := acctest.RandomNameWithPrefix("test-template")
	templateFormat := "host-$4"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccBulkhostnametemplateBasicConfig(templateName, templateFormat),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBulkhostnametemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_name", templateName),
					resource.TestCheckResourceAttr(resourceName, "template_format", templateFormat),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccBulkhostnametemplateResource_disappears(t *testing.T) {
	resourceName := "nios_ipam_bulkhostnametemplate.test"
	var v ipam.Bulkhostnametemplate
	templateName := acctest.RandomNameWithPrefix("test-template")
	templateFormat := "host-$4"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBulkhostnametemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccBulkhostnametemplateBasicConfig(templateName, templateFormat),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBulkhostnametemplateExists(context.Background(), resourceName, &v),
					testAccCheckBulkhostnametemplateDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccBulkhostnametemplateResource_TemplateFormat(t *testing.T) {
	var resourceName = "nios_ipam_bulkhostnametemplate.test"
	var v ipam.Bulkhostnametemplate
	templateName := acctest.RandomNameWithPrefix("test-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccBulkhostnametemplateTemplateFormat(templateName, "server-$4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBulkhostnametemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_format", "server-$4"),
					resource.TestCheckResourceAttr(resourceName, "template_name", templateName),
				),
			},
			// Update and Read
			{
				Config: testAccBulkhostnametemplateTemplateFormat(templateName, "server-$3-$4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBulkhostnametemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_format", "server-$3-$4"),
					resource.TestCheckResourceAttr(resourceName, "template_name", templateName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccBulkhostnametemplateResource_TemplateName(t *testing.T) {
	var resourceName = "nios_ipam_bulkhostnametemplate.test"
	var v ipam.Bulkhostnametemplate
	templateName := acctest.RandomNameWithPrefix("test-template")
	templateNameUpdated := acctest.RandomNameWithPrefix("updated-template")
	templateFormat := "server-$4"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccBulkhostnametemplateTemplateName(templateName, templateFormat),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBulkhostnametemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_format", templateFormat),
					resource.TestCheckResourceAttr(resourceName, "template_name", templateName),
				),
			},
			// Update and Read
			{
				Config: testAccBulkhostnametemplateTemplateName(templateNameUpdated, templateFormat),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBulkhostnametemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_format", templateFormat),
					resource.TestCheckResourceAttr(resourceName, "template_name", templateNameUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckBulkhostnametemplateExists(ctx context.Context, resourceName string, v *ipam.Bulkhostnametemplate) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.IPAMAPI.
			BulkhostnametemplateAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForBulkhostnametemplate).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetBulkhostnametemplateResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetBulkhostnametemplateResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckBulkhostnametemplateDestroy(ctx context.Context, v *ipam.Bulkhostnametemplate) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.IPAMAPI.
			BulkhostnametemplateAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForBulkhostnametemplate).
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

func testAccCheckBulkhostnametemplateDisappears(ctx context.Context, v *ipam.Bulkhostnametemplate) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.IPAMAPI.
			BulkhostnametemplateAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccBulkhostnametemplateBasicConfig(templateName, templateFormat string) string {
	return fmt.Sprintf(`
resource "nios_ipam_bulkhostnametemplate" "test" {
    template_name   = %q
    template_format = %q
}
`, templateName, templateFormat)
}

func testAccBulkhostnametemplateTemplateFormat(templateName, templateFormat string) string {
	return fmt.Sprintf(`
resource "nios_ipam_bulkhostnametemplate" "test" {
    template_name   = %q
    template_format = %q
}
`, templateName, templateFormat)
}

func testAccBulkhostnametemplateTemplateName(templateName, templateFormat string) string {
	return fmt.Sprintf(`
resource "nios_ipam_bulkhostnametemplate" "test" {
    template_name   = %q
    template_format = %q
}
`, templateName, templateFormat)
}
