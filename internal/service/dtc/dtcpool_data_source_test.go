
package dtc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/acctest"
)

func TestAccDtcpoolDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_datasource_nios_DtcPool.test"
	resourceName := "nios_resource_nios_DtcPool.test"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "ROUND_ROBIN"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcpoolDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcpoolDataSourceConfigFilters(name , lbPreferredMethod),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
						}, testAccCheckDtcpoolResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccDtcpoolDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.nios_datasource_nios_DtcPool.test"
	resourceName := "nios_resource_nios_DtcPool.test"
	var v dtc.DtcPool
	name := acctest.RandomName()
	lbPreferredMethod := "TOPOLOGY"
	lbPreferredTopology := "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdG9wb2xvZ3lfcnVsZXNldA:topology_ruleset"
	monitors := []string{
		"dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http", "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp",
	}
	servers := []map[string]interface{}{
		{
			"server": "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Qtc2VydmVyLmNvbQ:test-server.com",
			"ratio":  100,
		},
		{
			"server": "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Qtc2VydmVyMi5jb20:test-server2.com",
			"ratio":  50,
		},
	}
	lbAlternateMethod := "DYNAMIC_RATIO"
	lbDynamicRatioAlternate := map[string]interface{}{
		"method":                "ROUND_TRIP_DELAY",
		"monitor":               "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http",
		"monitor_metric":        ".0",
		"monitor_weighing":      "RATIO",
		"invert_monitor_metric": false,
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcpoolDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcpoolDataSourceConfigExtAttrFilters(name , lbPreferredMethod, lbPreferredTopology, lbAlternateMethod, lbDynamicRatioAlternate, servers, monitors, "Blr"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
						}, testAccCheckDtcpoolResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDtcpoolResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
    return []resource.TestCheckFunc{
        resource.TestCheckResourceAttrPair(resourceName, "_ref", dataSourceName, "result.0._ref"),
        resource.TestCheckResourceAttrPair(resourceName, "auto_consolidated_monitors", dataSourceName, "result.0.auto_consolidated_monitors"),
        resource.TestCheckResourceAttrPair(resourceName, "availability", dataSourceName, "result.0.availability"),
        resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
        resource.TestCheckResourceAttrPair(resourceName, "consolidated_monitors", dataSourceName, "result.0.consolidated_monitors"),
        resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
        resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
        resource.TestCheckResourceAttrPair(resourceName, "health", dataSourceName, "result.0.health"),
        resource.TestCheckResourceAttrPair(resourceName, "lb_alternate_method", dataSourceName, "result.0.lb_alternate_method"),
        resource.TestCheckResourceAttrPair(resourceName, "lb_alternate_topology", dataSourceName, "result.0.lb_alternate_topology"),
        resource.TestCheckResourceAttrPair(resourceName, "lb_dynamic_ratio_alternate", dataSourceName, "result.0.lb_dynamic_ratio_alternate"),
        resource.TestCheckResourceAttrPair(resourceName, "lb_dynamic_ratio_preferred", dataSourceName, "result.0.lb_dynamic_ratio_preferred"),
        resource.TestCheckResourceAttrPair(resourceName, "lb_preferred_method", dataSourceName, "result.0.lb_preferred_method"),
        resource.TestCheckResourceAttrPair(resourceName, "lb_preferred_topology", dataSourceName, "result.0.lb_preferred_topology"),
        resource.TestCheckResourceAttrPair(resourceName, "monitors", dataSourceName, "result.0.monitors"),
        resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
        resource.TestCheckResourceAttrPair(resourceName, "quorum", dataSourceName, "result.0.quorum"),
        resource.TestCheckResourceAttrPair(resourceName, "servers", dataSourceName, "result.0.servers"),
        resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
        resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
    }
}

func testAccDtcpoolDataSourceConfigFilters(name , lbPreferredMethod string) string {
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test" {
  name = %q
  lb_preferred_method = %q
}

data "nios_datasource_nios_DtcPool" "test" {
  filters = {
	 name = nios_resource_nios_DtcPool.test.name
  }
}
`, name, lbPreferredMethod)
}

func testAccDtcpoolDataSourceConfigExtAttrFilters(name, lbPreferredMethod, lbPreferredTopology string, lbAlternateMethod string , lbDynamicRatioAlternate map[string]interface{}, servers []map[string]interface{}, monitors []string, extAttrsValue string) string {
	monitorsHCL := formatMonitorsToHCL(monitors)
	serversHCL := formatServersToHCL(servers)
	lbDynamicRatioAlternateStr := formatLbDynamicRatioToHCL(lbDynamicRatioAlternate)
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test" {
name = %q
	lb_preferred_method = %q
	lb_preferred_topology = %q
	lb_alternate_method = %q
    lb_dynamic_ratio_alternate = %s
    servers = %s
	monitors = %s
  extattrs = {
    Site = {
        value = %q
    }
  	}
}

data "nios_datasource_nios_DtcPool" "test" {
  extattrfilters = {
	"Site" = nios_resource_nios_DtcPool.test.extattrs.Site.value
  }
}
`, name, lbPreferredMethod, lbPreferredTopology, lbAlternateMethod, lbDynamicRatioAlternateStr, serversHCL , monitorsHCL, extAttrsValue)
}

