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

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/acctest"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/utils"
)

// TODO : Add readable attributes for the resource
var readableAttributesForDtcPool = "extattrs,lb_preferred_method,auto_consolidated_monitors,availability,comment,consolidated_monitors,disable,health,lb_alternate_method,lb_alternate_topology,lb_dynamic_ratio_alternate,lb_dynamic_ratio_preferred,lb_preferred_topology,name,quorum,servers,ttl,use_ttl,monitors"

func TestAccDtcpoolResource_basic(t *testing.T) {
	var resourceName = "nios_resource_nios_DtcPool.test"
	var v dtc.DtcPool
	name := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcpoolBasicConfig(name, "ROUND_ROBIN"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "lb_preferred_method", "ROUND_ROBIN"),
					resource.TestCheckResourceAttr(resourceName, "auto_consolidated_monitors", "false"),
					resource.TestCheckResourceAttr(resourceName, "availability", "ALL"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
					resource.TestCheckResourceAttr(resourceName, "lb_alternate_method", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.method", "MONITOR"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.invert_monitor_metric", "false"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.monitor_weighing", "RATIO"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_alternate.method", "MONITOR"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_alternate.invert_monitor_metric", "false"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_alternate.monitor_weighing", "RATIO"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcpoolResource_disappears(t *testing.T) {
	resourceName := "nios_resource_nios_DtcPool.test"
	var v dtc.DtcPool
	name := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcpoolDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcpoolBasicConfig(name, "ROUND_ROBIN"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					testAccCheckDtcpoolDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcpoolResource_Ref(t *testing.T) {
	var resourceName = "nios_resource_nios_DtcPool.test__ref"
	var v dtc.DtcPool

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcpoolRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "_ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcpoolRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "_ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcpoolResource_AutoConsolidatedMonitors(t *testing.T) {
	var resourceName = "nios_resource_nios_DtcPool.test_auto_consolidated_monitors"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "ROUND_ROBIN"
	autoConsolidatedMonitors := "true"
	autoConsolidatedMonitorsUpdate := "false"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcpoolAutoConsolidatedMonitors(name, lbPreferredMethod, autoConsolidatedMonitors),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_consolidated_monitors", autoConsolidatedMonitors),
				),
			},
			// Update and Read
			{
				Config: testAccDtcpoolAutoConsolidatedMonitors(name, lbPreferredMethod, autoConsolidatedMonitorsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_consolidated_monitors", autoConsolidatedMonitorsUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcpoolResource_Availability(t *testing.T) {
	var resourceName = "nios_resource_nios_DtcPool.test_availability"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "ROUND_ROBIN"
	initialAvailability := "ANY"
	updatedAvailability := "ALL"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcpoolAvailability(name, lbPreferredMethod, initialAvailability),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "availability", initialAvailability),
					resource.TestCheckResourceAttr(resourceName, "lb_preferred_method", lbPreferredMethod),
				),
			},
			// Update and Read
			{
				Config: testAccDtcpoolAvailability(name, lbPreferredMethod, updatedAvailability),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "availability", updatedAvailability),
					resource.TestCheckResourceAttr(resourceName, "lb_preferred_method", lbPreferredMethod),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcpoolResource_Comment(t *testing.T) {
	var resourceName = "nios_resource_nios_DtcPool.test_comment"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "ROUND_ROBIN"
	initialComment := "pool testing"
	updatedComment := "updated pool comment"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcpoolComment(name, lbPreferredMethod, initialComment),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "pool testing"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcpoolComment(name, lbPreferredMethod, updatedComment),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "updated pool comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcpoolResource_ConsolidatedMonitors(t *testing.T) {
	var resourceName = "nios_resource_nios_DtcPool.test_consolidated_monitors"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "ROUND_ROBIN"
	consolidatedMonitors := []map[string]interface{}{
		{
			"monitor":                   "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http",
			"availability":              "ANY",
			"full_health_communication": false,
			"members":                   []string{"infoblox.localdomain"},
		},
	}
	monitors := []string{"dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http", "dtc:monitor:icmp/ZG5zLmlkbnNfbW9uaXRvcl9pY21wJGljbXA:icmp"}
	consolidatedMonitorsUpdate := []map[string]interface{}{
		{
			"monitor":                   "dtc:monitor:icmp/ZG5zLmlkbnNfbW9uaXRvcl9pY21wJGljbXA:icmp",
			"availability":              "ALL",
			"full_health_communication": false,
			"members":                   []string{"infoblox.localdomain"},
		},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcpoolConsolidatedMonitors(name, lbPreferredMethod, monitors, consolidatedMonitors),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "consolidated_monitors.0.monitor", "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http"),
					resource.TestCheckResourceAttr(resourceName, "consolidated_monitors.0.availability", "ANY"),
					resource.TestCheckResourceAttr(resourceName, "consolidated_monitors.0.full_health_communication", "false"),
					resource.TestCheckResourceAttr(resourceName, "consolidated_monitors.0.members.0", "infoblox.localdomain"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcpoolConsolidatedMonitors(name, lbPreferredMethod, monitors, consolidatedMonitorsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "consolidated_monitors.0.monitor", "dtc:monitor:icmp/ZG5zLmlkbnNfbW9uaXRvcl9pY21wJGljbXA:icmp"),
					resource.TestCheckResourceAttr(resourceName, "consolidated_monitors.0.availability", "ALL"),
					resource.TestCheckResourceAttr(resourceName, "consolidated_monitors.0.full_health_communication", "false"),
					resource.TestCheckResourceAttr(resourceName, "consolidated_monitors.0.members.0", "infoblox.localdomain"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcpoolResource_Disable(t *testing.T) {
	var resourceName = "nios_resource_nios_DtcPool.test_disable"
	var v dtc.DtcPool

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcpoolDisable("true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcpoolDisable("false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcpoolResource_Extattrs(t *testing.T) {
	var resourceName = "nios_resource_nios_DtcPool.test_extattrs"
	var v dtc.DtcPool

	extAttrValue1 := acctest.RandomName()
	name := acctest.RandomName()
	lbPreferredMethod := "ROUND_ROBIN"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcpoolExtattrs(name, lbPreferredMethod, map[string]struct{ value string }{
					"Site": {
						value: extAttrValue1,
					},
					"mystrung": {
						value: "myvalue",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site.value", extAttrValue1),
					resource.TestCheckResourceAttr(resourceName, "extattrs.mystrung.value", "myvalue"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcpoolExtattrs(name, lbPreferredMethod, map[string]struct{ value string }{
					"Site": {
						value: "updatedSiteValue",
					},
					"mystrung": {
						value: "updatedMyValue",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site.value", "updatedSiteValue"),
					resource.TestCheckResourceAttr(resourceName, "extattrs.mystrung.value", "updatedMyValue"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

//health field is not writable
// func TestAccDtcpoolResource_Health(t *testing.T) {
// 	var resourceName = "nios_resource_nios_DtcPool.test_health"
// 	var v dtc.DtcPool

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:                 func() { acctest.PreCheck(t) },
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			// Create and Read
// 			{
// 				Config: testAccDtcpoolHealth("HEALTH_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "health", "HEALTH_REPLACE_ME"),
// 				),
// 			},
// 			// Update and Read
// 			{
// 				Config: testAccDtcpoolHealth("HEALTH_UPDATE_REPLACE_ME"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
// 					resource.TestCheckResourceAttr(resourceName, "health", "HEALTH_UPDATE_REPLACE_ME"),
// 				),
// 			},
// 			// Delete testing automatically occurs in TestCase
// 		},
// 	})
// }

func TestAccDtcpoolResource_LbAlternateMethod(t *testing.T) {
	var resourceName = "nios_resource_nios_DtcPool.test_lb_alternate_method"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "TOPOLOGY"
	lbPreferredToplogyMethod := "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdG9wb2xvZ3lfcnVsZXNldA:topology_ruleset"
	lbAlternateMethod := "ALL_AVAILABLE"
	lbAlternateMethodUpdate := "GLOBAL_AVAILABILITY"
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
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcpoolLbAlternateMethod(name, lbPreferredMethod, lbPreferredToplogyMethod, lbAlternateMethod, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_alternate_method", lbAlternateMethod),
					resource.TestCheckResourceAttr(resourceName, "lb_preferred_method", lbPreferredMethod),
					resource.TestCheckResourceAttr(resourceName, "lb_preferred_topology", lbPreferredToplogyMethod),
					resource.TestCheckResourceAttr(resourceName, "servers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.server", "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Qtc2VydmVyLmNvbQ:test-server.com"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.ratio", "100"),
					resource.TestCheckResourceAttr(resourceName, "servers.1.server", "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Qtc2VydmVyMi5jb20:test-server2.com"),
					resource.TestCheckResourceAttr(resourceName, "servers.1.ratio", "50"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcpoolLbAlternateMethod(name, lbPreferredMethod, lbPreferredToplogyMethod, lbAlternateMethodUpdate, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_alternate_method", lbAlternateMethodUpdate),
					resource.TestCheckResourceAttr(resourceName, "lb_preferred_method", lbPreferredMethod),
					resource.TestCheckResourceAttr(resourceName, "lb_preferred_topology", lbPreferredToplogyMethod),
					resource.TestCheckResourceAttr(resourceName, "servers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.server", "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Qtc2VydmVyLmNvbQ:test-server.com"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.ratio", "100"),
					resource.TestCheckResourceAttr(resourceName, "servers.1.server", "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Qtc2VydmVyMi5jb20:test-server2.com"),
					resource.TestCheckResourceAttr(resourceName, "servers.1.ratio", "50"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcpoolResource_LbAlternateTopology(t *testing.T) {
	var resourceName = "nios_resource_nios_DtcPool.test_lb_alternate_topology"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "TOPOLOGY"
	lbPreferredTopology := "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdG9wb2xvZ3lfcnVsZXNldA:topology_ruleset"
	lbAlternateMethod := "TOPOLOGY"
	lbAlternateTopology := "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdG9wb2xvZ3lfcnVsZXNldDI:topology_ruleset2"
	lbPreferredTopologyUpdate := "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdG9wb2xvZ3lfcnVsZXNldDI:topology_ruleset2"
	lbAlternateTopologyUpdate := "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdG9wb2xvZ3lfcnVsZXNldA:topology_ruleset"
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
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcpoolLbAlternateTopology(name, lbPreferredMethod, lbPreferredTopology, lbAlternateMethod, lbAlternateTopology, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_alternate_topology", lbAlternateTopology),
				),
			},
			// Update and Read
			{
				Config: testAccDtcpoolLbAlternateTopology(name, lbPreferredMethod, lbPreferredTopologyUpdate, lbAlternateMethod, lbAlternateTopologyUpdate, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_alternate_topology", lbAlternateTopologyUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcpoolResource_LbDynamicRatioAlternate(t *testing.T) {
	var resourceName = "nios_resource_nios_DtcPool.test_lb_dynamic_ratio_alternate"
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
	lbDynamicRatioAlternateUpdate := map[string]interface{}{
		"method":                "MONITOR",
		"monitor":               "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp",
		"monitor_metric":        ".2",
		"monitor_weighing":      "RATIO",
		"invert_monitor_metric": false,
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcpoolLbDynamicRatioAlternate(name, lbPreferredMethod, lbPreferredTopology, lbAlternateMethod, lbDynamicRatioAlternate, servers , monitors),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_alternate.method", "ROUND_TRIP_DELAY"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_alternate.monitor", "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_alternate.monitor_metric", ".0"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_alternate.monitor_weighing", "RATIO"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_alternate.invert_monitor_metric", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcpoolLbDynamicRatioAlternate(name, lbPreferredMethod, lbPreferredTopology, lbAlternateMethod, lbDynamicRatioAlternateUpdate, servers , monitors),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_alternate.method", "MONITOR"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_alternate.monitor", "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_alternate.monitor_metric", ".2"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_alternate.monitor_weighing", "RATIO"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_alternate.invert_monitor_metric", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcpoolResource_LbDynamicRatioPreferred(t *testing.T) {
	var resourceName = "nios_resource_nios_DtcPool.test_lb_dynamic_ratio_preferred"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "DYNAMIC_RATIO"
	method := "ROUND_TRIP_DELAY"
	monitorWeighing := "RATIO"
	invertMonitorMetric := false
	monitor := "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http"
	monitorUpdated := "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp"
	monitorMetric := ".0"
	monitors := []string{monitor, monitorUpdated}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcpoolLbDynamicRatioPreferred(name, lbPreferredMethod, method, monitorWeighing, monitor, monitorMetric, invertMonitorMetric, monitors),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.method", "ROUND_TRIP_DELAY"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.monitor", monitor),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.monitor_metric", ".0"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.monitor_weighing", "RATIO"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.invert_monitor_metric", fmt.Sprintf("%t", invertMonitorMetric)),
				),
			},
			// Update and Read
			{
				Config: testAccDtcpoolLbDynamicRatioPreferred(name, lbPreferredMethod, method, monitorWeighing, monitorUpdated, monitorMetric, invertMonitorMetric, monitors),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.method", "ROUND_TRIP_DELAY"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.monitor", monitorUpdated),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.monitor_metric", ".0"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.monitor_weighing", "RATIO"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.invert_monitor_metric", fmt.Sprintf("%t", invertMonitorMetric)),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcpoolResource_LbPreferredMethod(t *testing.T) {
	var resourceName = "nios_resource_nios_DtcPool.test_lb_preferred_method"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "ROUND_ROBIN"
	lbpreferredMethodUpdate := "ALL_AVAILABLE"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcpoolLbPreferredMethod(name, lbPreferredMethod),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_preferred_method", lbPreferredMethod),
				),
			},
			// Update and Read
			{
				Config: testAccDtcpoolLbPreferredMethod(name, lbpreferredMethodUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_preferred_method", lbpreferredMethodUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcpoolResource_LbPreferredTopology(t *testing.T) {
	var resourceName = "nios_resource_nios_DtcPool.test_lb_preferred_topology"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "TOPOLOGY"
	lbPreferredToplogy := "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdG9wb2xvZ3lfcnVsZXNldA:topology_ruleset"
	servers := []map[string]interface{}{
		{
			"server": "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Qtc2VydmVyLmNvbQ:test-server.com",
			"ratio":  100,
		},
	}
	serversUpdated := []map[string]interface{}{
		{
			"server": "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Qtc2VydmVyMi5jb20:test-server2.com",
			"ratio":  50,
		},
	}
	lbPreferredToplogyUpdate := "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdG9wb2xvZ3lfcnVsZXNldDI:topology_ruleset2"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcpoolLbPreferredTopology(name, lbPreferredMethod, lbPreferredToplogy, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_preferred_topology", lbPreferredToplogy),
				),
			},
			// Update and Read
			{
				Config: testAccDtcpoolLbPreferredTopology(name, lbPreferredMethod, lbPreferredToplogyUpdate, serversUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_preferred_topology", lbPreferredToplogyUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcpoolResource_Monitors(t *testing.T) {
	var resourceName = "nios_resource_nios_DtcPool.test_monitors"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "ROUND_ROBIN"
	monitors := []string{"dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http", "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp"}
	monitorsUpdated := []string{"dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp", "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http", "dtc:monitor:icmp/ZG5zLmlkbnNfbW9uaXRvcl9pY21wJGljbXA:icmp"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcpoolMonitors(name, lbPreferredMethod, monitors),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "monitors.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "monitors.0", monitors[0]),
					resource.TestCheckResourceAttr(resourceName, "monitors.1", monitors[1]),
				),
			},
			// Update and Read
			{
				Config: testAccDtcpoolMonitors(name, lbPreferredMethod, monitorsUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "monitors.#", fmt.Sprintf("%d", len(monitorsUpdated))),
					resource.TestCheckResourceAttr(resourceName, "monitors.0", monitorsUpdated[0]),
					resource.TestCheckResourceAttr(resourceName, "monitors.1", monitorsUpdated[1]),
					resource.TestCheckResourceAttr(resourceName, "monitors.2", monitorsUpdated[2]),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcpoolResource_Name(t *testing.T) {
	var resourceName = "nios_resource_nios_DtcPool.test_name"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "ROUND_ROBIN"
	updateName := acctest.RandomName()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcpoolName(name, lbPreferredMethod),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDtcpoolName(updateName, lbPreferredMethod),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcpoolResource_Quorum(t *testing.T) {
	var resourceName = "nios_resource_nios_DtcPool.test_quorum"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "ROUND_ROBIN"
	quorum := "3"
	quorumUpdate := "5"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcpoolQuorum(name, lbPreferredMethod, quorum),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "quorum", quorum),
				),
			},
			// Update and Read
			{
				Config: testAccDtcpoolQuorum(name, lbPreferredMethod, quorumUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "quorum", quorumUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcpoolResource_Servers(t *testing.T) {
	var resourceName = "nios_resource_nios_DtcPool.test_servers"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "ROUND_ROBIN"
	servers := []map[string]interface{}{
		{
			"server": "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Qtc2VydmVyLmNvbQ:test-server.com",
			"ratio":  100,
		},
	}
	updatedServers := []map[string]interface{}{
		{
			"server": "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Qtc2VydmVyLmNvbQ:test-server.com",
			"ratio":  100,
		},
		{

			"server": "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Qtc2VydmVyMi5jb20:test-server2.com",
			"ratio":  50,
		}}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcpoolServers(name, lbPreferredMethod, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.server", "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Qtc2VydmVyLmNvbQ:test-server.com"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.ratio", "100"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcpoolServers(name, lbPreferredMethod, updatedServers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "servers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.server", "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Qtc2VydmVyLmNvbQ:test-server.com"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.ratio", "100"),
					resource.TestCheckResourceAttr(resourceName, "servers.1.server", "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Qtc2VydmVyMi5jb20:test-server2.com"),
					resource.TestCheckResourceAttr(resourceName, "servers.1.ratio", "50"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcpoolResource_Ttl(t *testing.T) {
	var resourceName = "nios_resource_nios_DtcPool.test_ttl"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "ROUND_ROBIN"
	ttl := "24"
	updateTtl := "34"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcpoolTtl(name, lbPreferredMethod, ttl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", ttl),
				),
			},
			// Update and Read
			{
				Config: testAccDtcpoolTtl(name, lbPreferredMethod, updateTtl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", updateTtl),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcpoolResource_UseTtl(t *testing.T) {
	var resourceName = "nios_resource_nios_DtcPool.test_use_ttl"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "ROUND_ROBIN"
	use_ttl := "true"
	ttl := "24"
	update_use_ttl := "false"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcpoolUseTtl(name, lbPreferredMethod, ttl, use_ttl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", use_ttl),
				),
			},
			// Update and Read
			{
				Config: testAccDtcpoolUseTtl(name, lbPreferredMethod, ttl, update_use_ttl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcpoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", update_use_ttl),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDtcpoolExists(ctx context.Context, resourceName string, v *dtc.DtcPool) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DTCAPI.
			DtcpoolAPI.
			ReferenceGet(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFields2(readableAttributesForDtcPool).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDtcPoolResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDtcPoolResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDtcpoolDestroy(ctx context.Context, v *dtc.DtcPool) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DTCAPI.
			DtcpoolAPI.
			ReferenceGet(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFields2(readableAttributesForDtcPool).
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

func testAccCheckDtcpoolDisappears(ctx context.Context, v *dtc.DtcPool) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DTCAPI.
			DtcpoolAPI.
			ReferenceDelete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDtcpoolBasicConfig(name, lb_preferred_method string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test" {
				name = %q
				lb_preferred_method = %q
}
`, name, lb_preferred_method)
}

func testAccDtcpoolRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test__ref" {
    _ref = %q
}
`, ref)
}

func testAccDtcpoolAutoConsolidatedMonitors(name, lbPreferredMethod string, autoConsolidatedMonitors string) string {
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test_auto_consolidated_monitors" {
    name = %q
    lb_preferred_method = %q
    auto_consolidated_monitors = %q
}
`, name, lbPreferredMethod, autoConsolidatedMonitors)
}

func testAccDtcpoolAvailability(name, lbPreferredMethod, availability string) string {
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test_availability" {
	name = %q
	lb_preferred_method = %q
    availability = %q
}
`, name, lbPreferredMethod, availability)
}

func testAccDtcpoolComment(name, lbPreferredMethod, comment string) string {
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test_comment" {
	name = %q
	lb_preferred_method = %q
    comment = %q
}
`, name, lbPreferredMethod, comment)
}

func testAccDtcpoolConsolidatedMonitors(name, lbPreferredMethod string, monitors []string, consolidatedMonitors []map[string]interface{}) string {
	monitorsHCL := formatMonitorsToHCL(monitors)
	consolidatedMonitorsHCL := formatConsolidatedMonitorsToHCL(consolidatedMonitors)
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test_consolidated_monitors" {
    consolidated_monitors = %s
	disable = true
	name = %q
	lb_preferred_method = %q
	monitors = %s
}
`, consolidatedMonitorsHCL, name, lbPreferredMethod, monitorsHCL)
}

func testAccDtcpoolDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test_disable" {
	name = "test_dtc_pool"
	lb_preferred_method = "ROUND_ROBIN"
    disable = %q
}
`, disable)
}

func testAccDtcpoolExtattrs(name, lbPreferredMethod string, extAttrs map[string]struct{ value string }) string {
	valueStr := ""
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		valueStr += fmt.Sprintf(`%q = {`, k)
		valueStr += fmt.Sprintf(`
					value = %q,
		`, v.value)
		valueStr += "\t},"
		extattrsStr += valueStr
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test_extattrs" {
	name = %q
	lb_preferred_method = %q
    extattrs = %s
}
`, name, lbPreferredMethod, extattrsStr)
}

// func testAccDtcpoolHealth(health string) string {
// 	return fmt.Sprintf(`
// resource "nios_resource_nios_DtcPool" "test_health" {
//     health = %q
// }
// `, health)
// }

func testAccDtcpoolLbAlternateMethod(name, lbPreferredMethod, lbPreferredTopology, lbAlternateMethod string, servers []map[string]interface{}) string {

	serversHCL := formatServersToHCL(servers)
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test_lb_alternate_method" {
    name = %q
    lb_preferred_method = %q
    lb_preferred_topology = %q
    lb_alternate_method = %q
    servers = %s
}
`, name, lbPreferredMethod, lbPreferredTopology, lbAlternateMethod, serversHCL)
}

func testAccDtcpoolLbAlternateTopology(name, lbPreferredMethod, lbPreferredTopology, lbAlternateMethod, lbAlternateTopology string, servers []map[string]interface{}) string {
	serversHCL := formatServersToHCL(servers)
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test_lb_alternate_topology" {
    name = %q
    lb_preferred_method = %q
    lb_preferred_topology = %q
    lb_alternate_method = %q
    lb_alternate_topology = %q
    servers = %s
}
`, name, lbPreferredMethod, lbPreferredTopology, lbAlternateMethod, lbAlternateTopology, serversHCL)
}

func testAccDtcpoolLbDynamicRatioAlternate(name, lbPreferredMethod, lbPreferredTopology string, lbAlternateMethod string , lbDynamicRatioAlternate map[string]interface{}, servers []map[string]interface{}, monitors []string) string {
	monitorsHCL := formatMonitorsToHCL(monitors)
	serversHCL := formatServersToHCL(servers)
	lbDynamicRatioAlternateStr := formatLbDynamicRatioToHCL(lbDynamicRatioAlternate)
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test_lb_dynamic_ratio_alternate" {
	name = %q
	lb_preferred_method = %q
	lb_preferred_topology = %q
	lb_alternate_method = %q
    lb_dynamic_ratio_alternate = %s
    servers = %s
	monitors = %s
}
`, name, lbPreferredMethod, lbPreferredTopology, lbAlternateMethod, lbDynamicRatioAlternateStr, serversHCL , monitorsHCL)
}

func testAccDtcpoolLbDynamicRatioPreferred(name, lbPreferredMethod, method, monitorWeighing, monitor, monitorMetric string, invertMonitorMetric bool, monitors []string) string {
	monitorsList := formatMonitorsToHCL(monitors)
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test_lb_dynamic_ratio_preferred" {
	name = %q
	lb_preferred_method = %q
	lb_dynamic_ratio_preferred = {
		monitor_weighing = %q,
		invert_monitor_metric = %t,
		monitor = %q,
		monitor_metric = %q,
		method = %q
	}
	monitors = %s
}
`, name, lbPreferredMethod, monitorWeighing, invertMonitorMetric, monitor, monitorMetric, method, monitorsList)
}

func testAccDtcpoolLbPreferredMethod(name, lbPreferredMethod string) string {
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test_lb_preferred_method" {
    lb_preferred_method = %q
    name = %q
}
`, lbPreferredMethod, name)
}

func testAccDtcpoolLbPreferredTopology(name, lbPreferredMethod, lbPreferredTopology string, servers []map[string]interface{}) string {
	serversHCL := formatServersToHCL(servers)
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test_lb_preferred_topology" {
	name = %q
	lb_preferred_method = %q
    lb_preferred_topology = %q
    servers = %s
}
`, name, lbPreferredMethod, lbPreferredTopology, serversHCL)
}

func testAccDtcpoolMonitors(name, lbPreferredMethod string, monitors []string) string {
	monitorsList := make([]string, len(monitors))
	for i, m := range monitors {
		monitorsList[i] = fmt.Sprintf("%q", m)
	}
	monitorsStr := fmt.Sprintf("[%s]", strings.Join(monitorsList, ", "))
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test_monitors" {
    name = %q
    lb_preferred_method = %q
    monitors = %s
}
`, name, lbPreferredMethod, monitorsStr)
}

func testAccDtcpoolName(name string, lbPreferredMethod string) string {
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test_name" {
    name = %q
    lb_preferred_method = %q
}
`, name, lbPreferredMethod)
}

func testAccDtcpoolQuorum(name, lbPreferredMethod, quorum string) string {
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test_quorum" {
	name = %q
	lb_preferred_method = %q
    quorum = %q
}
`, name, lbPreferredMethod, quorum)
}

func testAccDtcpoolServers(name, lbPreferredMethod string, servers []map[string]interface{}) string {
	serversHCL := formatServersToHCL(servers)
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test_servers" {
    name = %q
    lb_preferred_method = %q
    servers = %s
}
`, name, lbPreferredMethod, serversHCL)
}

func testAccDtcpoolTtl(name, lbPreferredMethod, ttl string) string {
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test_ttl" {
    ttl = %q
    name = %q
    lb_preferred_method = %q
}
`, ttl, name, lbPreferredMethod)
}

func testAccDtcpoolUseTtl(name, lbPreferredMethod, ttl, useTtl string) string {
	return fmt.Sprintf(`
resource "nios_resource_nios_DtcPool" "test_use_ttl" {
    use_ttl = %q
    name = %q
    lb_preferred_method = %q
	ttl = %q
}
`, useTtl, name, lbPreferredMethod, ttl)
}

func formatMonitorsToHCL(monitors []string) string {
	monitorsList := make([]string, len(monitors))
	for i, m := range monitors {
		monitorsList[i] = fmt.Sprintf("%q", m)
	}
	return fmt.Sprintf("[%s]", strings.Join(monitorsList, ", "))
}

func formatServersToHCL(servers []map[string]interface{}) string {
	var serverBlocks []string

	for _, server := range servers {
		serverBlock := fmt.Sprintf(`    {
      server = %q
      ratio = %d
    }`, server["server"], server["ratio"])
		serverBlocks = append(serverBlocks, serverBlock)
	}

	return fmt.Sprintf(`[
%s
  ]`, strings.Join(serverBlocks, ",\n"))
}

func formatConsolidatedMonitorsToHCL(monitors []map[string]interface{}) string {
	var monitorBlocks []string

	for _, monitor := range monitors {
		// Convert members slice to HCL string format
		members := monitor["members"].([]string)
		membersStr := make([]string, len(members))
		for i, m := range members {
			membersStr[i] = fmt.Sprintf(`%q`, m)
		}

		monitorBlock := fmt.Sprintf(`{
      monitor = %q
      availability = %q
	  full_health_communication = %t
      members = [%s]
    }`,
			monitor["monitor"], monitor["availability"], monitor["full_health_communication"], strings.Join(membersStr, ", "))

		monitorBlocks = append(monitorBlocks, monitorBlock)
	}

	return fmt.Sprintf(`[
%s
  ]`, strings.Join(monitorBlocks, ",\n"))
}

func formatLbDynamicRatioToHCL(ratio map[string]interface{}) string {
	return fmt.Sprintf(`{
      method = %q
      monitor = %q
      monitor_metric = %q
      monitor_weighing = %q
      invert_monitor_metric = %t
    }`,
		ratio["method"],
		ratio["monitor"],
		ratio["monitor_metric"],
		ratio["monitor_weighing"],
		ratio["invert_monitor_metric"])
}
