package dtc_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
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
	name := acctest.RandomNameWithPrefix("dtc-topology")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcTopologyBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcTopologyResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_topology.test"
	var v dtc.DtcTopology
	name := acctest.RandomNameWithPrefix("dtc-topology")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcTopologyDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcTopologyBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					testAccCheckDtcTopologyDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcTopologyResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_topology.test_comment"
	var v dtc.DtcTopology
	name := acctest.RandomNameWithPrefix("dtc-topology")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcTopologyComment(name, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcTopologyComment(name, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcTopologyResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dtc_topology.test_extattrs"
	var v dtc.DtcTopology
	name := acctest.RandomNameWithPrefix("dtc-topology")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcTopologyExtAttrs(name, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccDtcTopologyExtAttrs(name, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcTopologyResource_Name(t *testing.T) {
	var resourceName = "nios_dtc_topology.test_name"
	var v dtc.DtcTopology
	name := acctest.RandomNameWithPrefix("dtc-topology")
	nameUpdate := acctest.RandomNameWithPrefix("dtc-topology")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcTopologyName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDtcTopologyName(nameUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcTopologyResource_Rules(t *testing.T) {
	var resourceName = "nios_dtc_topology.test_rules"
	var v dtc.DtcTopology
	name := acctest.RandomNameWithPrefix("dtc-topology")
	rules1 := []map[string]interface{}{
		{
			"dest_type": "SERVER",
			"name":      "example-server1",
			"host":      "2.2.2.2",
		},
	}
	rules2 := []map[string]interface{}{
		{
			"dest_type": "SERVER",
			"name":      "example-server1",
			"host":      "2.2.2.2",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcTopologyRules(name, rules1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rules.0.dest_type", "SERVER"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcTopologyRules(name, rules2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rules.0.dest_type", "SERVER"),
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

func testAccDtcTopologyBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_dtc_topology" "test" {
	name = "%s"
}
`, name)
}

func testAccDtcTopologyComment(name, comment string) string {
	return fmt.Sprintf(`
resource "nios_dtc_topology" "test_comment" {
	name = %q
    comment = %q
}
`, name, comment)
}

func testAccDtcTopologyExtAttrs(name string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_dtc_topology" "test_extattrs" {
	name = %q
    extattrs = %s
}
`, name, extattrsStr)
}

func testAccDtcTopologyName(name string) string {
	return fmt.Sprintf(`
resource "nios_dtc_topology" "test_name" {
    name = %q
}
`, name)
}

// func testAccDtcTopologyRules(name string , rules []map[string]any) string {
// 	return fmt.Sprintf(`
// resource "nios_dtc_topology" "test_rules" {
//     rules = %q
// }
// `, rules)
// }

func testAccDtcServer(resourceName, name, host string) string {
	return fmt.Sprintf(`
resource "nios_dtc_server" "%s" {
    name = "%s"
    host = "%s"
}
`, resourceName, name, host)
}

func testAccDtcTopologyRules(topologyName string, rules []map[string]interface{}) string {
	var serverConfigs []string
	var ruleConfigs []string

	// First, create all the server resources
	for i, rule := range rules {
		if rule["dest_type"] == "SERVER" {
			serverResourceName := fmt.Sprintf("server_%d", i)
			serverName := rule["name"].(string)
			serverHost := rule["host"].(string)

			// Add server configuration - now passing serverName correctly
			serverConfigs = append(serverConfigs, testAccDtcServer(serverResourceName, serverName, serverHost))

			// Build rule configuration with reference to the server
			ruleConfig := fmt.Sprintf(`
        {
            dest_type = "SERVER"
            destination_link = nios_dtc_server.server_%d.ref
        }`, i)
			ruleConfigs = append(ruleConfigs, ruleConfig)
		}
	}

	// Join all server configs
	serversConfig := strings.Join(serverConfigs, "\n")

	// Join all rule configs
	rulesConfig := strings.Join(ruleConfigs, ",")

	// Build the complete configuration
	return fmt.Sprintf(`
%s

resource "nios_dtc_topology" "test_rules" {
    name = "%s"
    rules = [%s
    ]
}
`, serversConfig, topologyName, rulesConfig)
}
