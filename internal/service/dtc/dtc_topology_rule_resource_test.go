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

var readableAttributesForDtcTopologyRule = "dest_type,destination_link,return_type,sources,topology,valid"

func TestAccDtcTopologyRuleResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_topology_rule.test"
	var v dtc.DtcTopologyRule

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcTopologyRuleBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyRuleExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcTopologyRuleResource_disappears(t *testing.T) {
	t.Skip("Skipping as DtcTopologyRule resource cannot be deleted via API")
}

func TestAccDtcTopologyRuleResource_DestType(t *testing.T) {
	var resourceName = "nios_dtc_topology_rule.test_dest_type"
	var v dtc.DtcTopologyRule

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcTopologyRuleDestType("DEST_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dest_type", "DEST_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcTopologyRuleDestType("DEST_TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dest_type", "DEST_TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcTopologyRuleResource_DestinationLink(t *testing.T) {
	var resourceName = "nios_dtc_topology_rule.test_destination_link"
	var v dtc.DtcTopologyRule

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcTopologyRuleDestinationLink("DESTINATION_LINK_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "destination_link", "DESTINATION_LINK_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcTopologyRuleDestinationLink("DESTINATION_LINK_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "destination_link", "DESTINATION_LINK_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcTopologyRuleResource_ReturnType(t *testing.T) {
	var resourceName = "nios_dtc_topology_rule.test_return_type"
	var v dtc.DtcTopologyRule

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcTopologyRuleReturnType("RETURN_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "return_type", "RETURN_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcTopologyRuleReturnType("RETURN_TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "return_type", "RETURN_TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcTopologyRuleResource_Sources(t *testing.T) {
	var resourceName = "nios_dtc_topology_rule.test_sources"
	var v dtc.DtcTopologyRule

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcTopologyRuleSources("SOURCES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sources", "SOURCES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcTopologyRuleSources("SOURCES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sources", "SOURCES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDtcTopologyRuleExists(ctx context.Context, resourceName string, v *dtc.DtcTopologyRule) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DTCAPI.
			DtcTopologyRuleAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDtcTopologyRule).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDtcTopologyRuleResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDtcTopologyRuleResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDtcTopologyRuleDestroy(ctx context.Context, v *dtc.DtcTopologyRule) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DTCAPI.
			DtcTopologyRuleAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDtcTopologyRule).
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

// func testAccCheckDtcTopologyRuleDisappears(ctx context.Context, v *dtc.DtcTopologyRule) resource.TestCheckFunc {
// 	// Delete the resource externally to verify disappears test
// 	return func(state *terraform.State) error {
// 		_, err := acctest.NIOSClient.DTCAPI.
// 			DtcTopologyRuleAPI.
// 			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
// 			Execute()
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	}
// }

func testAccDtcTopologyRuleBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dtc_topology_rule" "test" {
}
`)
}

func testAccDtcTopologyRuleDestType(destType string) string {
	return fmt.Sprintf(`
resource "nios_dtc_topology_rule" "test_dest_type" {
    dest_type = %q
}
`, destType)
}

func testAccDtcTopologyRuleDestinationLink(destinationLink string) string {
	return fmt.Sprintf(`
resource "nios_dtc_topology_rule" "test_destination_link" {
    destination_link = %q
}
`, destinationLink)
}

func testAccDtcTopologyRuleReturnType(returnType string) string {
	return fmt.Sprintf(`
resource "nios_dtc_topology_rule" "test_return_type" {
    return_type = %q
}
`, returnType)
}

func testAccDtcTopologyRuleSources(sources string) string {
	return fmt.Sprintf(`
resource "nios_dtc_topology_rule" "test_sources" {
    sources = %q
}
`, sources)
}
