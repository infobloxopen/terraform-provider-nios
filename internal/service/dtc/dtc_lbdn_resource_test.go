package dtc_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/acctest"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/utils"
)

// TODO: Required parents for the execution of tests
var readableAttributesForDtcLbdn = "extattrs,disable,auth_zones,auto_consolidated_monitors,lb_method,patterns,persistence,pools,priority,topology,types,health,ttl,use_ttl"

func TestAccDtcLbdnResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test"
	var v dtc.DtcLbdn
	name := "dtc-lbdn-" + acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnBasicConfig(name, "ROUND_ROBIN"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "lb_method", "ROUND_ROBIN"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "auto_consolidated_monitors", "false"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "persistence", "0"),
					resource.TestCheckResourceAttr(resourceName, "priority", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_lbdn.test"
	var v dtc.DtcLbdn
	name := "dtc-lbdn-" + acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcLbdnDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcLbdnBasicConfig(name, "ROUND_ROBIN"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					testAccCheckDtcLbdnDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcLbdnResource_AuthZones(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_auth_zones"
	var v dtc.DtcLbdn
	name := "dtc-lbdn-" + acctest.RandomName()
	authZones := []string{"zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS50ZXN0:test.com/default", "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5yZWNvcmRfdGVzdA:record_test.com/default"}
	authZonesUpdated := []string{"zone_auth/ZG5zLnpvbmUkLjEuY29tLnRlc3Q:test.com/default.custom_view"}
	pools := []map[string]interface{}{
		{"pool": "dtc:pool/ZG5zLmlkbnNfcG9vbCRkdGNfcG9vbC5jb20:dtc_pool.com", "ratio": 2},
	}
	patterns := []string{"*.test.com", "*.record_test.com"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnAuthZones(name, "SOURCE_IP_HASH", authZones, pools, patterns),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_zones.#", "2"),
					resource.TestCheckTypeSetElemAttr(resourceName, "auth_zones.*", "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS50ZXN0:test.com/default"),
					resource.TestCheckTypeSetElemAttr(resourceName, "auth_zones.*", "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5yZWNvcmRfdGVzdA:record_test.com/default"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnAuthZones("lbdn-resource-13", "SOURCE_IP_HASH", authZonesUpdated, pools, patterns),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_zones.#", "1"), // Check element count
					resource.TestCheckTypeSetElemAttr(resourceName, "auth_zones.*", "zone_auth/ZG5zLnpvbmUkLjEuY29tLnRlc3Q:test.com/default.custom_view"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_AutoConsolidatedMonitors(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_auto_consolidated_monitors"
	var v dtc.DtcLbdn
	name := "dtc-lbdn-" + acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnAutoConsolidatedMonitors(name, "RATIO", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_consolidated_monitors", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnAutoConsolidatedMonitors(name, "RATIO", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_consolidated_monitors", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_comment"
	var v dtc.DtcLbdn
	name := "dtc-lbdn-" + acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnComment(name, "GLOBAL_AVAILABILITY", "resource-comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "resource-comment"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnComment(name, "GLOBAL_AVAILABILITY", "resource-comment-update"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "resource-comment-update"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_Disable(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_disable"
	var v dtc.DtcLbdn
	name := "dtc-lbdn-" + acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnDisable(name, "RATIO", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnDisable(name, "RATIO", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_extattrs"
	var v dtc.DtcLbdn
	name := "dtc-lbdn-" + acctest.RandomName()
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnExtattrs(name, "ROUND_ROBIN", map[string]string{"Site": extAttrValue1}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnExtattrs(name, "ROUND_ROBIN", map[string]string{"Site": extAttrValue2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_LbMethod(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_lb_method"
	var v dtc.DtcLbdn
	name := "dtc-lbdn-" + acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnLbMethod(name, "GLOBAL_AVAILABILITY", nil),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_method", "GLOBAL_AVAILABILITY"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnLbMethod(name, "RATIO", nil),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_method", "RATIO"),
				),
			},
			{
				Config: testAccDtcLbdnLbMethod(name, "ROUND_ROBIN", nil),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_method", "ROUND_ROBIN"),
				),
			},
			{
				Config: testAccDtcLbdnLbMethod(name, "SOURCE_IP_HASH", nil),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_method", "SOURCE_IP_HASH"),
				),
			},
			{
				Config: testAccDtcLbdnLbMethod(name, "TOPOLOGY", utils.Ptr("dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdG9wbzE:topo1")),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_method", "TOPOLOGY"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_Name(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_name"
	var v dtc.DtcLbdn
	name1 := "dtc-lbdn-" + acctest.RandomName()
	name2 := "dtc-lbdn-" + acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnName(name1, "SOURCE_IP_HASH"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnName(name2, "SOURCE_IP_HASH"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_Patterns(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_patterns"
	var v dtc.DtcLbdn
	name := "dtc-lbdn-" + acctest.RandomName()
	patterns := []string{"*.test.com", "*.info.com"}
	patternsUpdated := []string{"*.test123.com", "*.info*.com"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnPatterns(name, "SOURCE_IP_HASH", patterns),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "patterns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "patterns.0", "*.test.com"),
					resource.TestCheckResourceAttr(resourceName, "patterns.1", "*.info.com"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnPatterns(name, "SOURCE_IP_HASH", patternsUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "patterns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "patterns.0", "*.test123.com"),
					resource.TestCheckResourceAttr(resourceName, "patterns.1", "*.info*.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_Persistence(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_persistence"
	var v dtc.DtcLbdn
	name := "dtc-lbdn-" + acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnPersistence(name, "ROUND_ROBIN", "3"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "persistence", "3"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnPersistence(name, "ROUND_ROBIN", "8"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "persistence", "8"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_Pools(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_pools"
	var v dtc.DtcLbdn
	name := "dtc-lbdn-" + acctest.RandomName()
	pools := []map[string]interface{}{
		{"pool": "dtc:pool/ZG5zLmlkbnNfcG9vbCRkdGNfcG9vbF9fbmV3MzAy:dtc_pool__new302", "ratio": 2},
		{"pool": "dtc:pool/ZG5zLmlkbnNfcG9vbCRkdGNfcG9vbF9fbmV3NDg3:dtc_pool__new487", "ratio": 3},
	}
	poolsUpdated := []map[string]interface{}{
		{"pool": "dtc:pool/ZG5zLmlkbnNfcG9vbCRkdGNfcG9vbF9fbmV3NDg5:dtc_pool__new489", "ratio": 2},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnPools(name, "ROUND_ROBIN", pools),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pools.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "pools.0.pool", "dtc:pool/ZG5zLmlkbnNfcG9vbCRkdGNfcG9vbF9fbmV3MzAy:dtc_pool__new302"),
					resource.TestCheckResourceAttr(resourceName, "pools.0.ratio", "2"),
					resource.TestCheckResourceAttr(resourceName, "pools.1.pool", "dtc:pool/ZG5zLmlkbnNfcG9vbCRkdGNfcG9vbF9fbmV3NDg3:dtc_pool__new487"),
					resource.TestCheckResourceAttr(resourceName, "pools.1.ratio", "3"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnPools(name, "ROUND_ROBIN", poolsUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pools.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "pools.0.pool", "dtc:pool/ZG5zLmlkbnNfcG9vbCRkdGNfcG9vbF9fbmV3NDg5:dtc_pool__new489"),
					resource.TestCheckResourceAttr(resourceName, "pools.0.ratio", "2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_Priority(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_priority"
	var v dtc.DtcLbdn
	name := "dtc-lbdn-" + acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnPriority(name, "RATIO", "1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "priority", "1"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnPriority(name, "RATIO", "3"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "priority", "3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_Topology(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_topology"
	var v dtc.DtcLbdn
	name := "dtc-lbdn-" + acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnTopology(name, "TOPOLOGY", "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdG9wbzE:topo1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "topology", "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdG9wbzE:topo1"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnTopology(name, "TOPOLOGY", "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdG9wbzI:topo2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "topology", "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdG9wbzI:topo2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_Ttl(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_ttl"
	var v dtc.DtcLbdn
	name := "dtc-lbdn-" + acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnTtl(name, "GLOBAL_AVAILABILITY", 260, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "260"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnTtl(name, "GLOBAL_AVAILABILITY", 480, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "480"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_Types(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_types"
	var v dtc.DtcLbdn
	name := "dtc-lbdn-" + acctest.RandomName()
	types := []string{"A", "AAAA", "CNAME"}
	typesUpdated := []string{"A", "AAAA", "CNAME", "SRV"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnTypes(name, "ROUND_ROBIN", types),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "types.#", "3"),
					resource.TestCheckTypeSetElemAttr(resourceName, "types.*", "A"),
					resource.TestCheckTypeSetElemAttr(resourceName, "types.*", "AAAA"),
					resource.TestCheckTypeSetElemAttr(resourceName, "types.*", "CNAME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnTypes(name, "ROUND_ROBIN", typesUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "types.#", "4"),
					resource.TestCheckTypeSetElemAttr(resourceName, "types.*", "A"),
					resource.TestCheckTypeSetElemAttr(resourceName, "types.*", "AAAA"),
					resource.TestCheckTypeSetElemAttr(resourceName, "types.*", "CNAME"),
					resource.TestCheckTypeSetElemAttr(resourceName, "types.*", "SRV"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_use_ttl"
	var v dtc.DtcLbdn
	name := "dtc-lbdn-" + acctest.RandomName()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnUseTtl(name, "SOURCE_IP_HASH", true, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnUseTtl(name, "SOURCE_IP_HASH", false, 120),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDtcLbdnExists(ctx context.Context, resourceName string, v *dtc.DtcLbdn) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DTCAPI.
			DtcLbdnAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDtcLbdn).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDtcLbdnResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDtcLbdnResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDtcLbdnDestroy(ctx context.Context, v *dtc.DtcLbdn) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DTCAPI.
			DtcLbdnAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDtcLbdn).
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

func testAccCheckDtcLbdnDisappears(ctx context.Context, v *dtc.DtcLbdn) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DTCAPI.
			DtcLbdnAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDtcLbdnBasicConfig(name, lbMethod string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test" {
	name = %q
	lb_method = %q
}
`, name, lbMethod)
}

func testAccDtcLbdnAuthZones(name, lbMethod string, authZones []string, pools []map[string]interface{}, patterns []string) string {
	authZonesStr := "[\n"
	for _, zone := range authZones {
		authZonesStr += fmt.Sprintf("\t%q,\n", zone)
	}
	authZonesStr += "]"
	poolsStr := "[\n"
	for _, pool := range pools {
		poolsStr += fmt.Sprintf("\t{\n\t\tpool = %q,\n\t\tratio = %d\n\t},\n", pool["pool"], pool["ratio"])
	}
	poolsStr += "]"
	patternsStr := "[\n"
	for _, pattern := range patterns {
		patternsStr += fmt.Sprintf("\t%q,\n", pattern)
	}
	patternsStr += "]"
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_auth_zones" {
    name = %q
    lb_method = %q
    auth_zones = %s
    pools = %s
    patterns = %s
}
`, name, lbMethod, authZonesStr, poolsStr, patternsStr)
}

func testAccDtcLbdnAutoConsolidatedMonitors(name, lbMethod, autoConsolidatedMonitors string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_auto_consolidated_monitors" {
	name = %q
	lb_method = %q
    auto_consolidated_monitors = %q
}
`, name, lbMethod, autoConsolidatedMonitors)
}

func testAccDtcLbdnComment(name, lbMethod, comment string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_comment" {
    name = %q
    lb_method = %q
    comment = %q
}
`, name, lbMethod, comment)
}

func testAccDtcLbdnDisable(name, lbMethod, disable string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_disable" {
    name = %q
    lb_method = %q
    disable = %q
}
`, name, lbMethod, disable)
}

func testAccDtcLbdnExtattrs(name, lbMethod string, extAttrs map[string]string) string {
	extattrsStr := "{"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`%s = %q`, k, v)
	}
	extattrsStr += "}"
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_extattrs" {
    name = %q
    lb_method = %q
    extattrs = %s
}
`, name, lbMethod, extattrsStr)
}

func testAccDtcLbdnLbMethod(name, lbMethod string, topology *string) string {
	var extraConfig string
	if lbMethod == "TOPOLOGY" && topology != nil {
		extraConfig = fmt.Sprintf(`topology = %q`, *topology)
	}

	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_lb_method" {
  name      = %q
  lb_method = %q
  %s
}
`, name, lbMethod, extraConfig)
}

func testAccDtcLbdnName(name, lbMethod string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_name" {
    name = %q
    lb_method = %q
}
`, name, lbMethod)
}

func testAccDtcLbdnPatterns(name, lbMethod string, patterns []string) string {
	patternsStr := "[\n"
	for _, pattern := range patterns {
		patternsStr += fmt.Sprintf("\t%q,\n", pattern)
	}
	patternsStr += "]"
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_patterns" {
    name = %q
    lb_method = %q
    patterns = %s
}
`, name, lbMethod, patternsStr)
}

func testAccDtcLbdnPersistence(name, lbMethod, persistence string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_persistence" {
	name = %q
	lb_method = %q
    persistence = %q
}
`, name, lbMethod, persistence)
}

func testAccDtcLbdnPools(name, lbMethod string, pools []map[string]interface{}) string {
	poolsStr := "[\n"
	for _, pool := range pools {
		poolsStr += fmt.Sprintf("\t{\n\t\tpool = %q,\n\t\tratio = %d\n\t},\n", pool["pool"], pool["ratio"])
	}
	poolsStr += "]"
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_pools" {
	name = %q
	lb_method = %q
    pools = %s
}
`, name, lbMethod, poolsStr)
}

func testAccDtcLbdnPriority(name, lbMethod, priority string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_priority" {
	name = %q
	lb_method = %q
    priority = %q
}
`, name, lbMethod, priority)
}

func testAccDtcLbdnTopology(name, lbMethod, topology string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_topology" {
	name = %q
	lb_method = %q
    topology = %q
}
`, name, lbMethod, topology)
}

func testAccDtcLbdnTtl(name, lbMethod string, ttl uint32, useTtl bool) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_ttl" {
	name = %q
	lb_method = %q
    ttl = %d
	use_ttl = %t
}
`, name, lbMethod, ttl, useTtl)
}

func testAccDtcLbdnTypes(name, lbMethod string, types []string) string {
	typesStr := "[\n"
	for _, lbdnType := range types {
		typesStr += fmt.Sprintf("\t%q,\n", lbdnType)
	}
	typesStr += "]"
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_types" {
	name = %q
	lb_method = %q
    types = %s
}
`, name, lbMethod, typesStr)
}

func testAccDtcLbdnUseTtl(name, lbMethod string, useTtl bool, ttl uint32) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_use_ttl" {
	name = %q
	lb_method = %q
    use_ttl = %t
	ttl = %d
}
`, name, lbMethod, useTtl, ttl)
}
