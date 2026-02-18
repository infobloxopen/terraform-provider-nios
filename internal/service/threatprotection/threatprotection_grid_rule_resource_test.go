package threatprotection_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/threatprotection"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForThreatprotectionGridRule = "allowed_actions,category,comment,config,description,disabled,is_factory_reset_enabled,name,ruleset,sid,template,type"

func TestAccThreatprotectionGridRuleResource_basic(t *testing.T) {
	var resourceName = "nios_threatprotection_grid_rule.test"
	var v threatprotection.ThreatprotectionGridRule

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionGridRuleBasicConfig("TEMPLATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionGridRuleExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "template", "TEMPLATE_REPLACE_ME"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatprotectionGridRuleResource_disappears(t *testing.T) {
	resourceName := "nios_threatprotection_grid_rule.test"
	var v threatprotection.ThreatprotectionGridRule

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckThreatprotectionGridRuleDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccThreatprotectionGridRuleBasicConfig("TEMPLATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionGridRuleExists(context.Background(), resourceName, &v),
					testAccCheckThreatprotectionGridRuleDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccThreatprotectionGridRuleResource_Import(t *testing.T) {
	var resourceName = "nios_threatprotection_grid_rule.test"
	var v threatprotection.ThreatprotectionGridRule

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionGridRuleBasicConfig("TEMPLATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionGridRuleExists(context.Background(), resourceName, &v),
				),
			},
			// Import with PlanOnly to detect differences
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccThreatprotectionGridRuleImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "ref",
				PlanOnly:                             true,
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccThreatprotectionGridRuleImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatprotectionGridRuleResource_Comment(t *testing.T) {
	var resourceName = "nios_threatprotection_grid_rule.test_comment"
	var v threatprotection.ThreatprotectionGridRule

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionGridRuleComment("TEMPLATE_REPLACE_ME", "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionGridRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatprotectionGridRuleComment("TEMPLATE_REPLACE_ME", "Updated comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionGridRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the object"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatprotectionGridRuleResource_Config(t *testing.T) {
	var resourceName = "nios_threatprotection_grid_rule.test_config"
	var v threatprotection.ThreatprotectionGridRule
	configVal := map[string]any{}
	configValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionGridRuleConfig("TEMPLATE_REPLACE_ME", configVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionGridRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "config", "CONFIG_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatprotectionGridRuleConfig("TEMPLATE_REPLACE_ME", configValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionGridRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "config", "CONFIG_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccThreatprotectionGridRuleResource_Disabled(t *testing.T) {
	var resourceName = "nios_threatprotection_grid_rule.test_disabled"
	var v threatprotection.ThreatprotectionGridRule

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionGridRuleDisabled("TEMPLATE_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionGridRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatprotectionGridRuleDisabled("TEMPLATE_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionGridRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatprotectionGridRuleResource_Template(t *testing.T) {
	var resourceName = "nios_threatprotection_grid_rule.test_template"
	var v threatprotection.ThreatprotectionGridRule

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionGridRuleTemplate("TEMPLATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionGridRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template", "TEMPLATE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatprotectionGridRuleTemplate("TEMPLATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionGridRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template", "TEMPLATE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckThreatprotectionGridRuleExists(ctx context.Context, resourceName string, v *threatprotection.ThreatprotectionGridRule) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.ThreatProtectionAPI.
			ThreatprotectionGridRuleAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForThreatprotectionGridRule).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetThreatprotectionGridRuleResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetThreatprotectionGridRuleResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckThreatprotectionGridRuleDestroy(ctx context.Context, v *threatprotection.ThreatprotectionGridRule) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.ThreatProtectionAPI.
			ThreatprotectionGridRuleAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForThreatprotectionGridRule).
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

func testAccCheckThreatprotectionGridRuleDisappears(ctx context.Context, v *threatprotection.ThreatprotectionGridRule) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.ThreatProtectionAPI.
			ThreatprotectionGridRuleAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccThreatprotectionGridRuleImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccThreatprotectionGridRuleBasicConfig(template string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_grid_rule" "test" {
    template = %q
}
`, template)
}

func testAccThreatprotectionGridRuleComment(template string, comment string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_grid_rule" "test_comment" {
    template = %q
    comment = %q
}
`, template, comment)
}

func testAccThreatprotectionGridRuleConfig(template string, config map[string]any) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_grid_rule" "test_config" {
    template = %q
    config = %s
}
`, template, config)
}

func testAccThreatprotectionGridRuleDisabled(template string, disabled string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_grid_rule" "test_disabled" {
    template = %q
    disabled = %q
}
`, template, disabled)
}

func testAccThreatprotectionGridRuleTemplate(template string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_grid_rule" "test_template" {
    template = %q
}
`, template)
}
