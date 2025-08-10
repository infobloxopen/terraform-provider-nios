package misc_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

//TODO:
// To be able to view the NXDOMAIN ruleset, install the Add Query Redirection license on the NIOS Grid

var readableAttributesForRuleset = "comment,disabled,name,nxdomain_rules,type"

func TestAccRulesetResource_basic(t *testing.T) {
	var resourceName = "nios_misc_ruleset.test"
	var v misc.Ruleset

	name := acctest.RandomNameWithPrefix("example_ruleset")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRulesetBasicConfig(name, "NXDOMAIN"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRulesetExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "type", "NXDOMAIN"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
					// resource.TestCheckResourceAttr(resourceName, "nxdomain_rules.0.action", "PASS"),
					// resource.TestCheckResourceAttr(resourceName, "nxdomain_rules.0.pattern", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRulesetResource_disappears(t *testing.T) {
	//t.Skip("Skipping test as it throws Expected a non-empty plan, but got an empty refresh plan")

	resourceName := "nios_misc_ruleset.test"
	var v misc.Ruleset

	name := acctest.RandomNameWithPrefix("example_ruleset")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRulesetDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRulesetBasicConfig(name, "NXDOMAIN"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRulesetExists(context.Background(), resourceName, &v),
					testAccCheckRulesetDisappears(context.Background(), &v),
				),
				//ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRulesetResource_Comment(t *testing.T) {
	var resourceName = "nios_misc_ruleset.test_comment"
	var v misc.Ruleset

	name := acctest.RandomNameWithPrefix("example_ruleset")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRulesetComment(name, "NXDOMAIN", "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRulesetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccRulesetComment(name, "NXDOMAIN", "Updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRulesetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRulesetResource_Disabled(t *testing.T) {
	var resourceName = "nios_misc_ruleset.test_disabled"
	var v misc.Ruleset

	name := acctest.RandomNameWithPrefix("example_ruleset")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRulesetDisabled(name, "NXDOMAIN", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRulesetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRulesetDisabled(name, "NXDOMAIN", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRulesetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRulesetResource_Name(t *testing.T) {
	var resourceName = "nios_misc_ruleset.test_name"
	var v misc.Ruleset

	name1 := acctest.RandomNameWithPrefix("example_ruleset")
	name2 := acctest.RandomNameWithPrefix("example_ruleset_updated")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRulesetName(name1, "NXDOMAIN"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRulesetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccRulesetName(name2, "NXDOMAIN"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRulesetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRulesetResource_NxdomainRules(t *testing.T) {
	var resourceName = "nios_misc_ruleset.test_nxdomain_rules"
	var v misc.Ruleset

	name := acctest.RandomNameWithPrefix("example_ruleset")

	nxDomainRules1 := []map[string]any{
		{
			"action":  "PASS",
			"pattern": "example.com",
		},
	}

	nxDomainRules2 := []map[string]any{
		{
			"action":  "MODIFY",
			"pattern": "test.com",
		},
	}

	nxdomainRulesHCL1 := FormatNxdomainRulesToHCL(nxDomainRules1)
	nxdomainRulesHCL2 := FormatNxdomainRulesToHCL(nxDomainRules2)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRulesetNxdomainRules(name, "NXDOMAIN", nxdomainRulesHCL1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRulesetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_rules.0.action", "PASS"),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_rules.0.pattern", "example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccRulesetNxdomainRules(name, "NXDOMAIN", nxdomainRulesHCL2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRulesetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_rules.0.action", "MODIFY"),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_rules.0.pattern", "test.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

//"Changing the type of Ruleset ruleset_2 is not allowed"
// func TestAccRulesetResource_Type(t *testing.T) {
// 	var resourceName = "nios_misc_ruleset.test_type"
// 	var v misc.Ruleset

// 	name := acctest.RandomNameWithPrefix("example_ruleset")

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccRulesetType(name, "BLACKLIST"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckRulesetExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "type", "BLACKLIST"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccRulesetType(name, "NXDOMAIN"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckRulesetExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "type", "NXDOMAIN"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

func testAccCheckRulesetExists(ctx context.Context, resourceName string, v *misc.Ruleset) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.MiscAPI.
			RulesetAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRuleset).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRulesetResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRulesetResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRulesetDestroy(ctx context.Context, v *misc.Ruleset) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.MiscAPI.
			RulesetAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRuleset).
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

func testAccCheckRulesetDisappears(ctx context.Context, v *misc.Ruleset) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.MiscAPI.
			RulesetAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRulesetBasicConfig(name, ruleset_type string) string {
	return fmt.Sprintf(`
resource "nios_misc_ruleset" "test" {
	name = %q
	type = %q
}
`, name, ruleset_type)
}

func testAccRulesetComment(name, ruleset_type, comment string) string {
	return fmt.Sprintf(`
resource "nios_misc_ruleset" "test_comment" {
	name = %q
	type = %q
    comment = %q
}
`, name, ruleset_type, comment)
}

func testAccRulesetDisabled(name, ruleset_type, disabled string) string {
	return fmt.Sprintf(`
resource "nios_misc_ruleset" "test_disabled" {
	name = %q
	type = %q
    disabled = %q
}
`, name, ruleset_type, disabled)
}

func testAccRulesetName(name, ruleset_type string) string {
	return fmt.Sprintf(`
resource "nios_misc_ruleset" "test_name" {
    name = %q
	type = %q
}
`, name, ruleset_type)
}

func testAccRulesetNxdomainRules(name, ruleset_type, nxdomainRules string) string {
	return fmt.Sprintf(`
resource "nios_misc_ruleset" "test_nxdomain_rules" {
	name = %q
	type = %q
    nxdomain_rules = %s
}
`, name, ruleset_type, nxdomainRules)
}

// func testAccRulesetType(name, ruleset_type string) string {
// 	return fmt.Sprintf(`
// resource "nios_misc_ruleset" "test_type" {
// 	name = %q
//     type = %q
// }
// `, name, ruleset_type)
// }

func FormatNxdomainRulesToHCL(rules []map[string]any) string {
	var ruleBlocks []string

	for _, rule := range rules {
		block := fmt.Sprintf(`    {
      action = %q
      pattern = %q
    }`, rule["action"], rule["pattern"])
		ruleBlocks = append(ruleBlocks, block)
	}

	return fmt.Sprintf(`[
%s
  ]`, strings.Join(ruleBlocks, ",\n"))
}
