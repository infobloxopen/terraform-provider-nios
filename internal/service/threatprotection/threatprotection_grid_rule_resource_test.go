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
				Config: testAccThreatprotectionGridRuleBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionGridRuleExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
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
				Config: testAccThreatprotectionGridRuleBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionGridRuleExists(context.Background(), resourceName, &v),
					testAccCheckThreatprotectionGridRuleDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccThreatprotectionGridRuleResource_Ref(t *testing.T) {
	var resourceName = "nios_threatprotection_grid_rule.test_ref"
	var v threatprotection.ThreatprotectionGridRule

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionGridRuleRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionGridRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatprotectionGridRuleRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionGridRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
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
				Config: testAccThreatprotectionGridRuleComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionGridRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatprotectionGridRuleComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionGridRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatprotectionGridRuleResource_Config(t *testing.T) {
	var resourceName = "nios_threatprotection_grid_rule.test_config"
	var v threatprotection.ThreatprotectionGridRule

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionGridRuleConfig("CONFIG_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionGridRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "config", "CONFIG_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatprotectionGridRuleConfig("CONFIG_UPDATE_REPLACE_ME"),
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
				Config: testAccThreatprotectionGridRuleDisabled("DISABLED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionGridRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "DISABLED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatprotectionGridRuleDisabled("DISABLED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionGridRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "DISABLED_UPDATE_REPLACE_ME"),
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
				Config: testAccThreatprotectionGridRuleTemplate("TEMPLATE_UPDATE_REPLACE_ME"),
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

func testAccThreatprotectionGridRuleBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_threatprotection_grid_rule" "test" {
}
`)
}

func testAccThreatprotectionGridRuleRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_grid_rule" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccThreatprotectionGridRuleComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_grid_rule" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccThreatprotectionGridRuleConfig(config string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_grid_rule" "test_config" {
    config = %q
}
`, config)
}

func testAccThreatprotectionGridRuleDisabled(disabled string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_grid_rule" "test_disabled" {
    disabled = %q
}
`, disabled)
}

func testAccThreatprotectionGridRuleTemplate(template string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_grid_rule" "test_template" {
    template = %q
}
`, template)
}
