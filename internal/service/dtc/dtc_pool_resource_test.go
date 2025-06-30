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

func TestAccDtcPoolResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_pool.test"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "ROUND_ROBIN"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcPoolBasicConfig(name, lbPreferredMethod),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "lb_preferred_method", lbPreferredMethod),
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

func TestAccDtcPoolResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_pool.test"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "ROUND_ROBIN"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcPoolDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcPoolBasicConfig(name, lbPreferredMethod),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					testAccCheckDtcPoolDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcPoolResource_Ref(t *testing.T) {
	var resourceName = "nios_dtc_pool.test__ref"
	var v dtc.DtcPool

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcPoolRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "_ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcPoolRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "_ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcPoolResource_AutoConsolidatedMonitors(t *testing.T) {
	var resourceName = "nios_dtc_pool.test_auto_consolidated_monitors"
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
				Config: testAccDtcPoolAutoConsolidatedMonitors(name, lbPreferredMethod, autoConsolidatedMonitors),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_consolidated_monitors", autoConsolidatedMonitors),
				),
			},
			// Update and Read
			{
				Config: testAccDtcPoolAutoConsolidatedMonitors(name, lbPreferredMethod, autoConsolidatedMonitorsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_consolidated_monitors", autoConsolidatedMonitorsUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcPoolResource_Availability(t *testing.T) {
	var resourceName = "nios_dtc_pool.test_availability"
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
				Config: testAccDtcPoolAvailability(name, lbPreferredMethod, initialAvailability),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "availability", initialAvailability),
					resource.TestCheckResourceAttr(resourceName, "lb_preferred_method", lbPreferredMethod),
				),
			},
			// Update and Read
			{
				Config: testAccDtcPoolAvailability(name, lbPreferredMethod, updatedAvailability),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "availability", updatedAvailability),
					resource.TestCheckResourceAttr(resourceName, "lb_preferred_method", lbPreferredMethod),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcPoolResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_pool.test_comment"
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
				Config: testAccDtcPoolComment(name, lbPreferredMethod, initialComment),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", initialComment),
				),
			},
			// Update and Read
			{
				Config: testAccDtcPoolComment(name, lbPreferredMethod, updatedComment),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", updatedComment),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcPoolResource_ConsolidatedMonitors(t *testing.T) {
	var resourceName = "nios_dtc_pool.test_consolidated_monitors"
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
				Config: testAccDtcPoolConsolidatedMonitors(name, lbPreferredMethod, monitors, consolidatedMonitors),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "consolidated_monitors.0.monitor", "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http"),
					resource.TestCheckResourceAttr(resourceName, "consolidated_monitors.0.availability", "ANY"),
					resource.TestCheckResourceAttr(resourceName, "consolidated_monitors.0.full_health_communication", "false"),
					resource.TestCheckResourceAttr(resourceName, "consolidated_monitors.0.members.0", "infoblox.localdomain"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcPoolConsolidatedMonitors(name, lbPreferredMethod, monitors, consolidatedMonitorsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
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

func TestAccDtcPoolResource_Disable(t *testing.T) {
	var resourceName = "nios_dtc_pool.test_disable"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "ROUND_ROBIN"
	initialDisable := "false"
	updatedDisable := "true"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcPoolDisable(name, lbPreferredMethod, initialDisable),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", initialDisable),
				),
			},
			// Update and Read
			{
				Config: testAccDtcPoolDisable(name, lbPreferredMethod, updatedDisable),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", updatedDisable),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcPoolResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dtc_pool.test_extattrs"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "ROUND_ROBIN"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcPoolExtAttrs(name, lbPreferredMethod, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccDtcPoolExtAttrs(name, lbPreferredMethod, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcPoolResource_LbAlternateMethod(t *testing.T) {
	var resourceName = "nios_dtc_pool.test_lb_alternate_method"
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
				Config: testAccDtcPoolLbAlternateMethod(name, lbPreferredMethod, lbPreferredToplogyMethod, lbAlternateMethod, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
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
				Config: testAccDtcPoolLbAlternateMethod(name, lbPreferredMethod, lbPreferredToplogyMethod, lbAlternateMethodUpdate, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
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

func TestAccDtcPoolResource_LbAlternateTopology(t *testing.T) {
	var resourceName = "nios_dtc_pool.test_lb_alternate_topology"
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
				Config: testAccDtcPoolLbAlternateTopology(name, lbPreferredMethod, lbPreferredTopology, lbAlternateMethod, lbAlternateTopology, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_alternate_topology", lbAlternateTopology),
				),
			},
			// Update and Read
			{
				Config: testAccDtcPoolLbAlternateTopology(name, lbPreferredMethod, lbPreferredTopologyUpdate, lbAlternateMethod, lbAlternateTopologyUpdate, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_alternate_topology", lbAlternateTopologyUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcPoolResource_LbDynamicRatioAlternate(t *testing.T) {
	var resourceName = "nios_dtc_pool.test_lb_dynamic_ratio_alternate"
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
				Config: testAccDtcPoolLbDynamicRatioAlternate(name, lbPreferredMethod, lbPreferredTopology, lbAlternateMethod, lbDynamicRatioAlternate, servers, monitors),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_alternate.method", "ROUND_TRIP_DELAY"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_alternate.monitor", "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_alternate.monitor_metric", ".0"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_alternate.monitor_weighing", "RATIO"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_alternate.invert_monitor_metric", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcPoolLbDynamicRatioAlternate(name, lbPreferredMethod, lbPreferredTopology, lbAlternateMethod, lbDynamicRatioAlternateUpdate, servers, monitors),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
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

func TestAccDtcPoolResource_LbDynamicRatioPreferred(t *testing.T) {
	var resourceName = "nios_dtc_pool.test_lb_dynamic_ratio_preferred"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "DYNAMIC_RATIO"
	monitors := []string{
		"dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http", "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp",
	}
	lbDynamicRatioPreferred := map[string]interface{}{
		"method":                "ROUND_TRIP_DELAY",
		"monitor":               "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http",
		"monitor_metric":        ".0",
		"monitor_weighing":      "RATIO",
		"invert_monitor_metric": false,
	}
	lbDynamicRatioPreferredUpdate := map[string]interface{}{
		"method":                "ROUND_TRIP_DELAY",
		"monitor":               "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp",
		"monitor_metric":        ".2",
		"monitor_weighing":      "RATIO",
		"invert_monitor_metric": true,
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcPoolLbDynamicRatioPreferred(name, lbPreferredMethod, lbDynamicRatioPreferred, monitors),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.method", "ROUND_TRIP_DELAY"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.monitor", monitors[0]),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.monitor_metric", ".0"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.monitor_weighing", "RATIO"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.invert_monitor_metric", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcPoolLbDynamicRatioPreferred(name, lbPreferredMethod, lbDynamicRatioPreferredUpdate, monitors),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.method", "ROUND_TRIP_DELAY"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.monitor", monitors[1]),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.monitor_metric", ".2"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.monitor_weighing", "RATIO"),
					resource.TestCheckResourceAttr(resourceName, "lb_dynamic_ratio_preferred.invert_monitor_metric", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcPoolResource_LbPreferredMethod(t *testing.T) {
	var resourceName = "nios_dtc_pool.test_lb_preferred_method"
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
				Config: testAccDtcPoolLbPreferredMethod(name, lbPreferredMethod),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_preferred_method", lbPreferredMethod),
				),
			},
			// Update and Read
			{
				Config: testAccDtcPoolLbPreferredMethod(name, lbpreferredMethodUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_preferred_method", lbpreferredMethodUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcPoolResource_LbPreferredTopology(t *testing.T) {
	var resourceName = "nios_dtc_pool.test_lb_preferred_topology"
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
				Config: testAccDtcPoolLbPreferredTopology(name, lbPreferredMethod, lbPreferredToplogy, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_preferred_topology", lbPreferredToplogy),
				),
			},
			// Update and Read
			{
				Config: testAccDtcPoolLbPreferredTopology(name, lbPreferredMethod, lbPreferredToplogyUpdate, serversUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_preferred_topology", lbPreferredToplogyUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcPoolResource_Monitors(t *testing.T) {
	var resourceName = "nios_dtc_pool.test_monitors"
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
				Config: testAccDtcPoolMonitors(name, lbPreferredMethod, monitors),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "monitors.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "monitors.0", monitors[0]),
					resource.TestCheckResourceAttr(resourceName, "monitors.1", monitors[1]),
				),
			},
			// Update and Read
			{
				Config: testAccDtcPoolMonitors(name, lbPreferredMethod, monitorsUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "monitors.0", monitorsUpdated[0]),
					resource.TestCheckResourceAttr(resourceName, "monitors.1", monitorsUpdated[1]),
					resource.TestCheckResourceAttr(resourceName, "monitors.2", monitorsUpdated[2]),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcPoolResource_Name(t *testing.T) {
	var resourceName = "nios_dtc_pool.test_name"
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
				Config: testAccDtcPoolName(name, lbPreferredMethod),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDtcPoolName(updateName, lbPreferredMethod),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", updateName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcPoolResource_Quorum(t *testing.T) {
	var resourceName = "nios_dtc_pool.test_quorum"
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
				Config: testAccDtcPoolQuorum(name, lbPreferredMethod, quorum),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "quorum", quorum),
				),
			},
			// Update and Read
			{
				Config: testAccDtcPoolQuorum(name, lbPreferredMethod, quorumUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "quorum", quorumUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcPoolResource_Servers(t *testing.T) {
	var resourceName = "nios_dtc_pool.test_servers"
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
				Config: testAccDtcPoolServers(name, lbPreferredMethod, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.server", "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Qtc2VydmVyLmNvbQ:test-server.com"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.ratio", "100"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcPoolServers(name, lbPreferredMethod, updatedServers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
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

func TestAccDtcPoolResource_Ttl(t *testing.T) {
	var resourceName = "nios_dtc_pool.test_ttl"
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
				Config: testAccDtcPoolTtl(name, lbPreferredMethod, ttl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", ttl),
				),
			},
			// Update and Read
			{
				Config: testAccDtcPoolTtl(name, lbPreferredMethod, updateTtl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", updateTtl),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcPoolResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dtc_pool.test_use_ttl"
	var v dtc.DtcPool

	name := acctest.RandomName()
	lbPreferredMethod := "ROUND_ROBIN"
	useTtl := "true"
	ttl := "24"
	updateUseTtl := "false"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcPoolUseTtl(name, lbPreferredMethod, ttl, useTtl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", useTtl),
				),
			},
			// Update and Read
			{
				Config: testAccDtcPoolUseTtl(name, lbPreferredMethod, ttl, updateUseTtl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcPoolExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", updateUseTtl),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDtcPoolExists(ctx context.Context, resourceName string, v *dtc.DtcPool) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DTCAPI.
			DtcPoolAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDtcPool).
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

func testAccCheckDtcPoolDestroy(ctx context.Context, v *dtc.DtcPool) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DTCAPI.
			DtcPoolAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDtcPool).
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

func testAccCheckDtcPoolDisappears(ctx context.Context, v *dtc.DtcPool) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DTCAPI.
			DtcPoolAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDtcPoolBasicConfig(name, lbPreferredMethod string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dtc_pool" "test" {
	name = %q
	lb_preferred_method = %q
}
`, name, lbPreferredMethod)
}

func testAccDtcPoolRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dtc_pool" "test__ref" {
    _ref = %q
}
`, ref)
}

func testAccDtcPoolAutoConsolidatedMonitors(name, lbPreferredMethod string, autoConsolidatedMonitors string) string {
	return fmt.Sprintf(`
resource "nios_dtc_pool" "test_auto_consolidated_monitors" {
name = %q
	lb_preferred_method = %q
    auto_consolidated_monitors = %q
}
`, name, lbPreferredMethod, autoConsolidatedMonitors)
}

func testAccDtcPoolAvailability(name, lbPreferredMethod, availability string) string {
	return fmt.Sprintf(`
resource "nios_dtc_pool" "test_availability" {
    name = %q
    lb_preferred_method = %q
    availability = %q
}
`, name, lbPreferredMethod, availability)
}

func testAccDtcPoolComment(name, lbPreferredMethod, comment string) string {
	return fmt.Sprintf(`
resource "nios_dtc_pool" "test_comment" {
    name = %q
    lb_preferred_method = %q
    comment = %q
}
`, name, lbPreferredMethod, comment)
}

func testAccDtcPoolConsolidatedMonitors(name, lbPreferredMethod string, monitors []string, consolidatedMonitors []map[string]interface{}) string {
	monitorsHCL := formatMonitorsToHCL(monitors)
	consolidatedMonitorsHCL := formatConsolidatedMonitorsToHCL(consolidatedMonitors)
	return fmt.Sprintf(`
resource "nios_dtc_pool" "test_consolidated_monitors" {
    consolidated_monitors = %s
	disable = true
	name = %q
	lb_preferred_method = %q
	monitors = %s
}
`, consolidatedMonitorsHCL, name, lbPreferredMethod, monitorsHCL)
}

func testAccDtcPoolDisable(name, lbPreferredMethod, disable string) string {
	return fmt.Sprintf(`
	resource "nios_dtc_pool" "test_disable" {
		name = %q
		lb_preferred_method = %q
		disable = %q
	}
`, name, lbPreferredMethod, disable)
}

func testAccDtcPoolExtAttrs(name, lbPreferredMethod string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_dtc_pool" "test_extattrs" {
    name = %q
 	lb_preferred_method = %q
 	extattrs = %s
}
`, name, lbPreferredMethod, extattrsStr)
}


func testAccDtcPoolLbAlternateMethod(name, lbPreferredMethod, lbPreferredTopology, lbAlternateMethod string, servers []map[string]interface{}) string {
	serversHCL := formatServersToHCL(servers)
	return fmt.Sprintf(`
resource "nios_dtc_pool" "test_lb_alternate_method" {
   name = %q
    lb_preferred_method = %q
    lb_preferred_topology = %q
    lb_alternate_method = %q
    servers = %s
}
`, name, lbPreferredMethod, lbPreferredTopology, lbAlternateMethod, serversHCL)
}

func testAccDtcPoolLbAlternateTopology(name, lbPreferredMethod, lbPreferredTopology, lbAlternateMethod, lbAlternateTopology string, servers []map[string]interface{}) string {
	serversHCL := formatServersToHCL(servers)
	return fmt.Sprintf(`
resource "nios_dtc_pool" "test_lb_alternate_topology" {
   name = %q
    lb_preferred_method = %q
    lb_preferred_topology = %q
    lb_alternate_method = %q
    lb_alternate_topology = %q
    servers = %s
}
`, name, lbPreferredMethod, lbPreferredTopology, lbAlternateMethod, lbAlternateTopology, serversHCL)
}

func testAccDtcPoolLbDynamicRatioAlternate(name, lbPreferredMethod, lbPreferredTopology string, lbAlternateMethod string, lbDynamicRatioAlternate map[string]interface{}, servers []map[string]interface{}, monitors []string) string {
	monitorsHCL := formatMonitorsToHCL(monitors)
	serversHCL := formatServersToHCL(servers)
	lbDynamicRatioAlternateStr := formatLbDynamicRatioToHCL(lbDynamicRatioAlternate)
	return fmt.Sprintf(`
resource "nios_dtc_pool" "test_lb_dynamic_ratio_alternate" {
name = %q
	lb_preferred_method = %q
	lb_preferred_topology = %q
	lb_alternate_method = %q
    lb_dynamic_ratio_alternate = %s
    servers = %s
	monitors = %s
}
`, name, lbPreferredMethod, lbPreferredTopology, lbAlternateMethod, lbDynamicRatioAlternateStr, serversHCL, monitorsHCL)
}

func testAccDtcPoolLbDynamicRatioPreferred(name, lbPreferredMethod string, lbDynamicRatioAlternate map[string]interface{}, monitors []string) string {
	monitorsHCL := formatMonitorsToHCL(monitors)
	lbDynamicRatioPreferredStr := formatLbDynamicRatioToHCL(lbDynamicRatioAlternate)
	return fmt.Sprintf(`
resource "nios_dtc_pool" "test_lb_dynamic_ratio_preferred" {
	name = %q
	lb_preferred_method = %q
    lb_dynamic_ratio_preferred = %s
	monitors = %s
}
`, name, lbPreferredMethod, lbDynamicRatioPreferredStr, monitorsHCL)
}

func testAccDtcPoolLbPreferredMethod(name, lbPreferredMethod string) string {
	return fmt.Sprintf(`
resource "nios_dtc_pool" "test_lb_preferred_method" {
    lb_preferred_method = %q
    name = %q
}
`, lbPreferredMethod, name)
}

func testAccDtcPoolLbPreferredTopology(name, lbPreferredMethod, lbPreferredTopology string, servers []map[string]interface{}) string {
	serversHCL := formatServersToHCL(servers)
	return fmt.Sprintf(`
resource "nios_dtc_pool" "test_lb_preferred_topology" {
   name = %q
	lb_preferred_method = %q
    lb_preferred_topology = %q
    servers = %s
}
`, name, lbPreferredMethod, lbPreferredTopology, serversHCL)
}

func testAccDtcPoolMonitors(name, lbPreferredMethod string, monitors []string) string {
	monitorsList := formatMonitorsToHCL(monitors)
	return fmt.Sprintf(`
resource "nios_dtc_pool" "test_monitors" {
   name = %q
    lb_preferred_method = %q
    monitors = %s
}
`, name, lbPreferredMethod, monitorsList)
}

func testAccDtcPoolName(name string, lbPreferredMethod string) string {
	return fmt.Sprintf(`
resource "nios_dtc_pool" "test_name" {
    name = %q
	lb_preferred_method = %q
}
`, name, lbPreferredMethod)
}

func testAccDtcPoolQuorum(name, lbPreferredMethod, quorum string) string {
	return fmt.Sprintf(`
resource "nios_dtc_pool" "test_quorum" {
    name = %q
	lb_preferred_method = %q
    quorum = %q
}
`, name, lbPreferredMethod, quorum)
}

func testAccDtcPoolServers(name , lbPreferredMethod string , servers []map[string]interface{}) string {
	serversHCL := formatServersToHCL(servers)
	return fmt.Sprintf(`
resource "nios_dtc_pool" "test_servers" {
name = %q
    lb_preferred_method = %q
    servers = %s
}
`, name, lbPreferredMethod, serversHCL)
}

func testAccDtcPoolTtl(name , lbPreferredMethod, ttl string) string {
	return fmt.Sprintf(`
resource "nios_dtc_pool" "test_ttl" {
    name = %q
    lb_preferred_method = %q
	use_ttl = true
    ttl = %q
}
`, name, lbPreferredMethod, ttl)
}

func testAccDtcPoolUseTtl(name, lbPreferredMethod, ttl, useTtl string) string {
	return fmt.Sprintf(`
resource "nios_dtc_pool" "test_use_ttl" {
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
