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

func TestAccDtcTopologyRuleDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dtc_topology_rule.test"
	resourceName := "nios_dtc_topology_rule.test"
	var v dtc.DtcTopologyRule

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcTopologyRuleDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcTopologyRuleDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDtcTopologyRuleExists(context.Background(), resourceName, &v),
					}, testAccCheckDtcTopologyRuleResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDtcTopologyRuleResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "dest_type", dataSourceName, "result.0.dest_type"),
		resource.TestCheckResourceAttrPair(resourceName, "destination_link", dataSourceName, "result.0.destination_link"),
		resource.TestCheckResourceAttrPair(resourceName, "return_type", dataSourceName, "result.0.return_type"),
		resource.TestCheckResourceAttrPair(resourceName, "sources", dataSourceName, "result.0.sources"),
		resource.TestCheckResourceAttrPair(resourceName, "topology", dataSourceName, "result.0.topology"),
		resource.TestCheckResourceAttrPair(resourceName, "valid", dataSourceName, "result.0.valid"),
	}
}

func testAccDtcTopologyRuleDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_dtc_topology_rule" "test" {
}

data "nios_dtc_topology_rule" "test" {
  filters = {
	 = nios_dtc_topology_rule.test.
  }
}
`)
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