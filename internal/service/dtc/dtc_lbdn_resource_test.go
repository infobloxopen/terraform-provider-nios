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

// TODO : Add readable attributes for the resource
var readableAttributesForDtcLbdn = ""

func TestAccDtcLbdnResource_basic(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test"
	var v dtc.DtcLbdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_disappears(t *testing.T) {
	resourceName := "nios_dtc_lbdn.test"
	var v dtc.DtcLbdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDtcLbdnDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDtcLbdnBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					testAccCheckDtcLbdnDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDtcLbdnResource_Ref(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test__ref"
	var v dtc.DtcLbdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "_ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "_ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_AuthZones(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_auth_zones"
	var v dtc.DtcLbdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnAuthZones("AUTH_ZONES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_zones", "AUTH_ZONES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnAuthZones("AUTH_ZONES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_zones", "AUTH_ZONES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_AutoConsolidatedMonitors(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_auto_consolidated_monitors"
	var v dtc.DtcLbdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnAutoConsolidatedMonitors("AUTO_CONSOLIDATED_MONITORS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_consolidated_monitors", "AUTO_CONSOLIDATED_MONITORS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnAutoConsolidatedMonitors("AUTO_CONSOLIDATED_MONITORS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_consolidated_monitors", "AUTO_CONSOLIDATED_MONITORS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_Comment(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_comment"
	var v dtc.DtcLbdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_Disable(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_disable"
	var v dtc.DtcLbdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_extattrs"
	var v dtc.DtcLbdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_Health(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_health"
	var v dtc.DtcLbdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnHealth("HEALTH_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "health", "HEALTH_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnHealth("HEALTH_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "health", "HEALTH_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_LbMethod(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_lb_method"
	var v dtc.DtcLbdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnLbMethod("LB_METHOD_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_method", "LB_METHOD_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnLbMethod("LB_METHOD_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lb_method", "LB_METHOD_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_Name(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_name"
	var v dtc.DtcLbdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_Patterns(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_patterns"
	var v dtc.DtcLbdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnPatterns("PATTERNS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "patterns", "PATTERNS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnPatterns("PATTERNS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "patterns", "PATTERNS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_Persistence(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_persistence"
	var v dtc.DtcLbdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnPersistence("PERSISTENCE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "persistence", "PERSISTENCE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnPersistence("PERSISTENCE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "persistence", "PERSISTENCE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_Pools(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_pools"
	var v dtc.DtcLbdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnPools("POOLS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pools", "POOLS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnPools("POOLS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pools", "POOLS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_Priority(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_priority"
	var v dtc.DtcLbdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnPriority("PRIORITY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "priority", "PRIORITY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnPriority("PRIORITY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "priority", "PRIORITY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_Topology(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_topology"
	var v dtc.DtcLbdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnTopology("TOPOLOGY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "topology", "TOPOLOGY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnTopology("TOPOLOGY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "topology", "TOPOLOGY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_Ttl(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_ttl"
	var v dtc.DtcLbdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnTtl("TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnTtl("TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_Types(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_types"
	var v dtc.DtcLbdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnTypes("TYPES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "types", "TYPES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnTypes("TYPES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "types", "TYPES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDtcLbdnResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dtc_lbdn.test_use_ttl"
	var v dtc.DtcLbdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDtcLbdnUseTtl("USE_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "USE_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDtcLbdnUseTtl("USE_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDtcLbdnExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "USE_TTL_UPDATE_REPLACE_ME"),
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

func testAccDtcLbdnBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test" {
}
`)
}

func testAccDtcLbdnRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test__ref" {
    _ref = %q
}
`, ref)
}

func testAccDtcLbdnAuthZones(authZones string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_auth_zones" {
    auth_zones = %q
}
`, authZones)
}

func testAccDtcLbdnAutoConsolidatedMonitors(autoConsolidatedMonitors string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_auto_consolidated_monitors" {
    auto_consolidated_monitors = %q
}
`, autoConsolidatedMonitors)
}

func testAccDtcLbdnComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccDtcLbdnDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccDtcLbdnExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccDtcLbdnHealth(health string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_health" {
    health = %q
}
`, health)
}

func testAccDtcLbdnLbMethod(lbMethod string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_lb_method" {
    lb_method = %q
}
`, lbMethod)
}

func testAccDtcLbdnName(name string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_name" {
    name = %q
}
`, name)
}

func testAccDtcLbdnPatterns(patterns string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_patterns" {
    patterns = %q
}
`, patterns)
}

func testAccDtcLbdnPersistence(persistence string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_persistence" {
    persistence = %q
}
`, persistence)
}

func testAccDtcLbdnPools(pools string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_pools" {
    pools = %q
}
`, pools)
}

func testAccDtcLbdnPriority(priority string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_priority" {
    priority = %q
}
`, priority)
}

func testAccDtcLbdnTopology(topology string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_topology" {
    topology = %q
}
`, topology)
}

func testAccDtcLbdnTtl(ttl string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_ttl" {
    ttl = %q
}
`, ttl)
}

func testAccDtcLbdnTypes(types string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_types" {
    types = %q
}
`, types)
}

func testAccDtcLbdnUseTtl(useTtl string) string {
	return fmt.Sprintf(`
resource "nios_dtc_lbdn" "test_use_ttl" {
    use_ttl = %q
}
`, useTtl)
}
