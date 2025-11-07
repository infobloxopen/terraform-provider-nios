package threatinsight_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/threatinsight"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForThreatinsightAllowlist = "comment,disable,fqdn,type"

func TestAccThreatinsightAllowlistResource_basic(t *testing.T) {
	var resourceName = "nios_threatinsight_allowlist.test"
	var v threatinsight.ThreatinsightAllowlist

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatinsightAllowlistBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatinsightAllowlistExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatinsightAllowlistResource_disappears(t *testing.T) {
	resourceName := "nios_threatinsight_allowlist.test"
	var v threatinsight.ThreatinsightAllowlist

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckThreatinsightAllowlistDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccThreatinsightAllowlistBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatinsightAllowlistExists(context.Background(), resourceName, &v),
					testAccCheckThreatinsightAllowlistDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccThreatinsightAllowlistResource_Ref(t *testing.T) {
	var resourceName = "nios_threatinsight_allowlist.test_ref"
	var v threatinsight.ThreatinsightAllowlist

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatinsightAllowlistRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatinsightAllowlistExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatinsightAllowlistRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatinsightAllowlistExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatinsightAllowlistResource_Comment(t *testing.T) {
	var resourceName = "nios_threatinsight_allowlist.test_comment"
	var v threatinsight.ThreatinsightAllowlist

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatinsightAllowlistComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatinsightAllowlistExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatinsightAllowlistComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatinsightAllowlistExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatinsightAllowlistResource_Disable(t *testing.T) {
	var resourceName = "nios_threatinsight_allowlist.test_disable"
	var v threatinsight.ThreatinsightAllowlist

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatinsightAllowlistDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatinsightAllowlistExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatinsightAllowlistDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatinsightAllowlistExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatinsightAllowlistResource_Fqdn(t *testing.T) {
	var resourceName = "nios_threatinsight_allowlist.test_fqdn"
	var v threatinsight.ThreatinsightAllowlist

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatinsightAllowlistFqdn("FQDN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatinsightAllowlistExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "FQDN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatinsightAllowlistFqdn("FQDN_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatinsightAllowlistExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "FQDN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckThreatinsightAllowlistExists(ctx context.Context, resourceName string, v *threatinsight.ThreatinsightAllowlist) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.ThreatInsightAPI.
			ThreatinsightAllowlistAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForThreatinsightAllowlist).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetThreatinsightAllowlistResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetThreatinsightAllowlistResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckThreatinsightAllowlistDestroy(ctx context.Context, v *threatinsight.ThreatinsightAllowlist) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.ThreatInsightAPI.
			ThreatinsightAllowlistAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForThreatinsightAllowlist).
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

func testAccCheckThreatinsightAllowlistDisappears(ctx context.Context, v *threatinsight.ThreatinsightAllowlist) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.ThreatInsightAPI.
			ThreatinsightAllowlistAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccThreatinsightAllowlistBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_threatinsight_allowlist" "test" {
}
`)
}

func testAccThreatinsightAllowlistRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_threatinsight_allowlist" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccThreatinsightAllowlistComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_threatinsight_allowlist" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccThreatinsightAllowlistDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_threatinsight_allowlist" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccThreatinsightAllowlistFqdn(fqdn string) string {
	return fmt.Sprintf(`
resource "nios_threatinsight_allowlist" "test_fqdn" {
    fqdn = %q
}
`, fqdn)
}
