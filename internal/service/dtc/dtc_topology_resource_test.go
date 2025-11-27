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
	serverName := acctest.RandomNameWithPrefix("dtc-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcTopologyRulesWithServer(name, serverName, "2.2.2.2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rules.0.dest_type", "SERVER"),
				),
			},
			// Update server host and verify topology rule still works
			{
				Config: testAccDtcTopologyRulesWithServer(name, serverName, "3.3.3.3"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rules.0.dest_type", "SERVER"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcTopologyResource_RulesWithPool(t *testing.T) {
	var resourceName = "nios_dtc_topology.test_rules_pool"
	var v dtc.DtcTopology
	name := acctest.RandomNameWithPrefix("dtc-topology")
	poolName := acctest.RandomNameWithPrefix("dtc-topology")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcTopologyRulesWithPool(name, poolName, "ROUND_ROBIN"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rules.0.dest_type", "POOL"),
				),
			},
			// Update pool lb_preferred_method and verify topology rule still works
			{
				Config: testAccDtcTopologyRulesWithPool(name, poolName, "GLOBAL_AVAILABILITY"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rules.0.dest_type", "POOL"),
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

func testAccDtcTopologyRulesWithServer(topologyName, serverName, serverHost string) string {
	return fmt.Sprintf(`
resource "nios_dtc_server" "test_server" {
    name = "%s"
    host = "%s"
}

resource "nios_dtc_topology" "test_rules" {
    name = "%s"
    rules = [
        {
            dest_type = "SERVER"
            destination_link = nios_dtc_server.test_server.ref
        }
    ]
}
`, serverName, serverHost, topologyName)
}

func testAccDtcTopologyRulesWithPool(topologyName, poolName, lbMethod string) string {
	return fmt.Sprintf(`
resource "nios_dtc_server" "test_server_for_pool" {
    name = "%s-server"
    host = "2.3.3.4"
}

resource "nios_dtc_pool" "test_pool" {
    name                = "%s"
    lb_preferred_method = "%s"
    servers = [
        {
            server = nios_dtc_server.test_server_for_pool.ref
            ratio  = 1
        }
    ]
}

resource "nios_dtc_topology" "test_rules_pool" {
    name = "%s"
    rules = [
        {
            dest_type = "POOL"
            destination_link = nios_dtc_pool.test_pool.ref
        }
    ]
}
`, poolName, poolName, lbMethod, topologyName)
}
