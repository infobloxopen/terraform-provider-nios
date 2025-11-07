package dtc_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForDtcTopology = "comment,extattrs,name,rules"

func TestAccDtcTopologyResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_topology.test"
	var v dtc.DtcTopology

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcTopologyBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcTopologyResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_topology.test"
	var v dtc.DtcTopology

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcTopologyDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcTopologyBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					testAccCheckDtcTopologyDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcTopologyResource_Ref(t *testing.T) {
	var resourceName = "nios_dtc_topology.test_ref"
	var v dtc.DtcTopology

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcTopologyRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcTopologyRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcTopologyResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_topology.test_comment"
	var v dtc.DtcTopology

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcTopologyComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcTopologyComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcTopologyResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dtc_topology.test_extattrs"
	var v dtc.DtcTopology

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcTopologyExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcTopologyExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcTopologyResource_Name(t *testing.T) {
	var resourceName = "nios_dtc_topology.test_name"
	var v dtc.DtcTopology

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcTopologyName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcTopologyName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcTopologyResource_Rules(t *testing.T) {
	var resourceName = "nios_dtc_topology.test_rules"
	var v dtc.DtcTopology

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcTopologyRules("RULES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rules", "RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcTopologyRules("RULES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rules", "RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDtcTopologyExists(ctx context.Context, resourceName string, v *dtc.DtcTopology) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DTCAPI.
			DtcTopologyAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDtcTopology).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDtcTopologyResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDtcTopologyResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDtcTopologyDestroy(ctx context.Context, v *dtc.DtcTopology) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DTCAPI.
			DtcTopologyAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDtcTopology).
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

func testAccCheckDtcTopologyDisappears(ctx context.Context, v *dtc.DtcTopology) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DTCAPI.
			DtcTopologyAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDtcTopologyBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dtc_topology" "test" {
}
`)
}

func testAccDtcTopologyRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dtc_topology" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccDtcTopologyComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dtc_topology" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccDtcTopologyExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_dtc_topology" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccDtcTopologyName(name string) string {
	return fmt.Sprintf(`
resource "nios_dtc_topology" "test_name" {
    name = %q
}
`, name)
}

func testAccDtcTopologyRules(rules string) string {
	return fmt.Sprintf(`
resource "nios_dtc_topology" "test_rules" {
    rules = %q
}
`, rules)
}
