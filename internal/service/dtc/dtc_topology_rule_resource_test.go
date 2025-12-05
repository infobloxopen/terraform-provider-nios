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
				Config: testAccDtcTopologyRuleBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyRuleExists(context.Background(), resourceName, &v),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "dest_type", "SERVER"),
					resource.TestCheckResourceAttr(resourceName, "destination_link", "dtc:server/ZG5zLmlkbnNfc2VydmVyJGR0Y19TZXJ2ZXI:dtc_Server"),
					resource.TestCheckResourceAttr(resourceName, "return_type", "REGULAR"),
					resource.TestCheckResourceAttr(resourceName, "sources.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "sources.0.source_op", "IS"),
					resource.TestCheckResourceAttr(resourceName, "sources.0.source_type", "SUBNET"),
					resource.TestCheckResourceAttr(resourceName, "sources.0.source_value", "10.10.0.0/24"),
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
	t.Skip("Skipping as dest_type cannot be updated via API since the destination type must match the topology rule type")
}

func TestAccDtcTopologyRuleResource_DestinationLink(t *testing.T) {
	t.Skip("Skipping as detination_link cannot be unlinked from the DTC Topology Rule via API")
}

func TestAccDtcTopologyRuleResource_ReturnType(t *testing.T) {
	var resourceName = "nios_dtc_topology_rule.test_return_type"
	var v dtc.DtcTopologyRule

	sources := []map[string]any{
		{
			"source_op":    "IS",
			"source_type":  "SUBNET",
			"source_value": "10.10.0.0/24",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcTopologyRuleReturnType("NXDOMAIN", sources),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "return_type", "NXDOMAIN"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcTopologyRuleReturnType("NOERR", sources),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "return_type", "NOERR"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcTopologyRuleResource_Sources(t *testing.T) {
	var resourceName = "nios_dtc_topology_rule.test_sources"
	var v dtc.DtcTopologyRule

	sources := []map[string]any{
		{
			"source_op":    "IS",
			"source_type":  "SUBNET",
			"source_value": "10.10.0.0/24",
		},
	}
	sourceUpdate := []map[string]any{
		{
			
			"source_op": "IS",
			"source_type": "CONTINENT",
			"source_value": "Africa",
		},
		{
			"source_op": "IS",
			"source_type": "COUNTRY",
			"source_value": "Algeria",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcTopologyRuleSources(sources),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sources.0.source_op", "IS"),
					resource.TestCheckResourceAttr(resourceName, "sources.0.source_type", "SUBNET"),
					resource.TestCheckResourceAttr(resourceName, "sources.0.source_value", "10.10.0.0/24"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcTopologyRuleSources(sourceUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcTopologyRuleExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sources.0.source_op", "IS"),
					resource.TestCheckResourceAttr(resourceName, "sources.0.source_type", "CONTINENT"),
					resource.TestCheckResourceAttr(resourceName, "sources.0.source_value", "Africa"),
					resource.TestCheckResourceAttr(resourceName, "sources.1.source_op", "IS"),
					resource.TestCheckResourceAttr(resourceName, "sources.1.source_type", "COUNTRY"),
					resource.TestCheckResourceAttr(resourceName, "sources.1.source_value", "Algeria"),
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

func testAccDtcTopologyRuleBasicConfig() string {
	return fmt.Sprintf(`
	import {
  id = "dtc:topology:rule/ZG5zLmlkbnNfdG9wb2xvZ3lfcnVsZSRkdGNfdG9wb2xvZ3lfcnVsZS42NGRkMDFhNi03MTUwLTRjY2YtYmFkOS04YzkxOTRmZWVmMTI:dtc_topology_rule/dtc_Server"
  to = nios_dtc_topology_rule.test
}
resource "nios_dtc_topology_rule" "test" {
dest_type        = "SERVER"
}
`)
}

func testAccDtcTopologyRuleDestType(destType string) string {
	return fmt.Sprintf(`
import {
  id = "dtc:topology:rule/ZG5zLmlkbnNfdG9wb2xvZ3lfcnVsZSRkdGNfdG9wb2xvZ3lfcnVsZS42NGRkMDFhNi03MTUwLTRjY2YtYmFkOS04YzkxOTRmZWVmMTI:dtc_topology_rule/dtc_Server"
  to = nios_dtc_topology_rule.test_dest_type
}
resource "nios_dtc_topology_rule" "test_dest_type" {
    dest_type = %q
}
`, destType)
}

func testAccDtcTopologyRuleReturnType(returnType string , sources []map[string]any) string {
	sourcesStr := utils.ConvertSliceOfMapsToHCL(sources)
	return fmt.Sprintf(`
	import {
  id = "dtc:topology:rule/ZG5zLmlkbnNfdG9wb2xvZ3lfcnVsZSR0b3BvbG9neV90ZXN0LmQwZDA2Zjk4LWY2YmMtNDk2OS1hOWFhLTA0NzcyMDgzNGZiMw:topology_test/NXDOMAIN/1"
  to = nios_dtc_topology_rule.test_return_type
}
resource "nios_dtc_topology_rule" "test_return_type" {
    return_type = %q
    sources = %s
}
`, returnType, sourcesStr)
}

func testAccDtcTopologyRuleSources(sources []map[string]any) string {
	sourcesStr := utils.ConvertSliceOfMapsToHCL(sources)
	return fmt.Sprintf(`
import {
  id = "dtc:topology:rule/ZG5zLmlkbnNfdG9wb2xvZ3lfcnVsZSR0b3BvbG9neV90ZXN0LmQwZDA2Zjk4LWY2YmMtNDk2OS1hOWFhLTA0NzcyMDgzNGZiMw:topology_test/NXDOMAIN/1"
  to = nios_dtc_topology_rule.test_sources
}
resource "nios_dtc_topology_rule" "test_sources" {
	return_type = "NXDOMAIN"
    sources = %s
}
`, sourcesStr)
}


func testAccResourceDTCTopologyRuleImportBasic() string {
	return `
resource "nios_dtc_topology_rule" "dtc_topology_rule1" {
  dest_type = "POOL"
  return_type = "NOERR"
}
`
}

func testAccResourceDTCTopologyRuleImportUpdate() string {
	return `
resource "nios_dtc_topology_rule" "dtc_topology_rule1" {
  dest_type = "POOL"
  return_type = "NOERR"
  sources = [
    {
      source_op    = "IS"
      source_type  = "CONTINENT"
      source_value = "Antarctica"
    },
    {
      source_op    = "IS"
      source_type  = "COUNTRY"
      source_value = "Antarctica"
    }
  ]
}
`
}
