package dns_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// TODO: OBJECTS TO BE PRESENT IN GRID FOR TESTS
// A NIOS grid with 2 members - infoblox.172_28_82_110 , infoblox.172_28_83_25
var readableAttributesForNsgroup = "comment,extattrs,external_primaries,external_secondaries,grid_primary,grid_secondaries,is_grid_default,is_multimaster,name,use_external_primary"

func TestAccNsgroupResource_basic(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test"
	var v dns.Nsgroup
	name := acctest.RandomNameWithPrefix("ns-group")
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.172_28_83_208",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupBasicConfig(name, gridPrimary),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "grid_primary.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "grid_primary.0.name", "infoblox.172_28_83_208"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "is_grid_default", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_multimaster", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_external_primary", "false"),
					resource.TestCheckResourceAttr(resourceName, "grid_primary.0.stealth", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupResource_disappears(t *testing.T) {
	resourceName := "nios_dns_nsgroup.test"
	var v dns.Nsgroup
	name := acctest.RandomNameWithPrefix("ns-group")
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.172_28_83_208",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupBasicConfig(name, gridPrimary),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					testAccCheckNsgroupDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNsgroupResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test_comment"
	var v dns.Nsgroup
	name := acctest.RandomNameWithPrefix("ns-group")
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.172_28_83_208",
		},
	}
	comment := "This is a test comment"
	commentUpdate := "This is an updated test comment"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupComment(name, gridPrimary, comment),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupComment(name, gridPrimary, commentUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", commentUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test_extattrs"
	var v dns.Nsgroup
	name := acctest.RandomNameWithPrefix("ns-group")
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.172_28_83_208",
		},
	}
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupExtAttrs(name, gridPrimary, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupExtAttrs(name, gridPrimary, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupResource_ExternalPrimaries(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test_external_primaries"
	var v dns.Nsgroup
	name := acctest.RandomNameWithPrefix("ns-group")
	externalPrimaries := []map[string]any{
		{
			"name":    "external.primary.1",
			"address": "2.3.4.5",
		},
	}
	gridSecondaries := []map[string]any{
		{
			"name": "infoblox.172_28_83_208",
		},
	}
	externalPrimariesUpdate := []map[string]any{
		{
			"name":    "external.primary.2",
			"address": "20.1.12.23",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupExternalPrimaries(name, externalPrimaries, gridSecondaries),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.name", "external.primary.1"),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.address", "2.3.4.5"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupExternalPrimaries(name, externalPrimariesUpdate, gridSecondaries),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.name", "external.primary.2"),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.address", "20.1.12.23"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupResource_ExternalSecondaries(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test_external_secondaries"
	var v dns.Nsgroup
	name := acctest.RandomNameWithPrefix("ns-group")
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.172_28_83_208",
		},
	}
	externalSecondaries := []map[string]any{
		{
			"name":    "external.secondary.1",
			"address": "2.3.3.3",
		},
	}
	externalSecondariesUpdate := []map[string]any{
		{
			"name":    "external.secondary.2",
			"address": "20.3.32.3",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupExternalSecondaries(name, externalSecondaries, gridPrimary),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.name", "external.secondary.1"),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.address", "2.3.3.3"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupExternalSecondaries(name, externalSecondariesUpdate, gridPrimary),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.name", "external.secondary.2"),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.address", "20.3.32.3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupResource_GridPrimary(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test_grid_primary"
	var v dns.Nsgroup
	name := acctest.RandomNameWithPrefix("ns-group")
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.172_28_83_208",
		},
	}
	gridPrimaryUpdate := []map[string]any{
		{
			"name": "infoblox.172_28_83_25",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupGridPrimary(name, gridPrimary),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_primary.0.name", "infoblox.172_28_83_208"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupGridPrimary(name, gridPrimaryUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_primary.0.name", "infoblox.172_28_83_25"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupResource_GridSecondaries(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test_grid_secondaries"
	var v dns.Nsgroup
	name := acctest.RandomNameWithPrefix("ns-group")
	gridSecondaries := []map[string]any{
		{
			"name": "infoblox.172_28_83_208",
		},
	}
	externalPrimaries := []map[string]any{
		{
			"name": "external,primaries",
			"address":"2.3.3.4",
		},
	}
	gridSecondariesUpdate := []map[string]any{
		{
			"name": "infoblox.172_28_83_25",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupGridSecondaries(name, externalPrimaries, gridSecondaries),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries.0.name", "infoblox.localdomain"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupGridSecondaries(name, externalPrimaries, gridSecondariesUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries.0.name", "infoblox.172_28_83_25"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupResource_IsGridDefault(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test_is_grid_default"
	var v dns.Nsgroup
	name := acctest.RandomNameWithPrefix("ns-group")
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.172_28_83_208",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupIsGridDefault(name, gridPrimary, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "is_grid_default", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupIsGridDefault(name, gridPrimary, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "is_grid_default", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupResource_IsMultimaster(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test_is_multimaster"
	var v dns.Nsgroup
	name := acctest.RandomNameWithPrefix("ns-group")
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.172_28_83_208",
		},
		{
			"name": "infoblox.172_28_83_25",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupIsMultimaster(name, gridPrimary, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "is_multimaster", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupResource_Name(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test_name"
	var v dns.Nsgroup
	name := acctest.RandomNameWithPrefix("ns-group")
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.172_28_83_208",
		},
	}
	nameUpdate := acctest.RandomNameWithPrefix("ns-group")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupName(name, gridPrimary),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupName(nameUpdate, gridPrimary),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupResource_UseExternalPrimary(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test_use_external_primary"
	var v dns.Nsgroup
	name := acctest.RandomNameWithPrefix("ns-group")
	gridSecondaries := []map[string]any{
		{
			"name": "infoblox.172_28_83_208",
		},
	}
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.172_28_83_25",
		},
	}
	externalPrimaries := []map[string]any{
		{
			"name":    "external.primary.1",
			"address": "2.3.4.5",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupUseExternalPrimary(name, gridPrimary, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_external_primary", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupUseExternalPrimaryUpdate(name, gridSecondaries, externalPrimaries, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_external_primary", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckNsgroupExists(ctx context.Context, resourceName string, v *dns.Nsgroup) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			NsgroupAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForNsgroup).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetNsgroupResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetNsgroupResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckNsgroupDestroy(ctx context.Context, v *dns.Nsgroup) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			NsgroupAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForNsgroup).
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

func testAccCheckNsgroupDisappears(ctx context.Context, v *dns.Nsgroup) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			NsgroupAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccNsgroupBasicConfig(name string, gridPrimary []map[string]any) string {
	gridPrimaryStr := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test" {
	name = %q
	grid_primary = %s
}
`, name, gridPrimaryStr)
}

func testAccNsgroupComment(name string, gridPrimary []map[string]any, comment string) string {
	gridPrimaryStr := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_comment" {
    name = %q
    grid_primary = %s
    comment = %q
}
`, name, gridPrimaryStr, comment)
}

func testAccNsgroupExtAttrs(name string, gridPrimary []map[string]any, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	gridPrimaryStr := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_extattrs" {
    name = %q
    grid_primary = %s
    extattrs = %s
}
`, name, gridPrimaryStr, extattrsStr)
}

func testAccNsgroupExternalPrimaries(name string, externalPrimaries, gridSecondaries []map[string]any) string {
	externalPrimariesStr := utils.ConvertSliceOfMapsToHCL(externalPrimaries)
	gridSecondariesStr := utils.ConvertSliceOfMapsToHCL(gridSecondaries)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_external_primaries" {
	name = %q
    external_primaries = %s
    grid_secondaries = %s
	use_external_primary = true
}
`, name, externalPrimariesStr, gridSecondariesStr)
}

func testAccNsgroupExternalSecondaries(name string, externalSecondaries, gridPrimary []map[string]any) string {
	externalSecondariesStr := utils.ConvertSliceOfMapsToHCL(externalSecondaries)
	gridPrimaryStr := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_external_secondaries" {
	name = %q
	grid_primary = %s
    external_secondaries = %s
}
`, name, gridPrimaryStr, externalSecondariesStr)
}

func testAccNsgroupGridPrimary(name string, gridPrimary []map[string]any) string {
	gridPrimaryStr := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_grid_primary" {
	name = %q
    grid_primary = %s
}
`, name, gridPrimaryStr)
}

func testAccNsgroupGridSecondaries(name string, externalPrimaries []map[string]any, gridSecondaries []map[string]any) string {
	externalPrimariesStr := utils.ConvertSliceOfMapsToHCL(externalPrimaries)
	gridSecondariesStr := utils.ConvertSliceOfMapsToHCL(gridSecondaries)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_grid_secondaries" {
	name = %q
    grid_secondaries = %s
	external_primaries = %s
	use_external_primary = true 
}
`, name, gridSecondariesStr, externalPrimariesStr)
}

func testAccNsgroupIsGridDefault(name string, gridPrimary []map[string]any, isGridDefault bool) string {
	gridPrimaryStr := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_is_grid_default" {
    name = %q
    grid_primary = %s
    is_grid_default = %t
}
`, name, gridPrimaryStr, isGridDefault)
}

func testAccNsgroupIsMultimaster(name string, gridPrimary []map[string]any, isMultimaster bool) string {
	gridPrimaryStr := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_is_multimaster" {
	name = %q
    grid_primary = %s
    is_multimaster = %t
}
`, name, gridPrimaryStr, isMultimaster)
}

func testAccNsgroupName(name string, gridPrimary []map[string]any) string {
	gridPrimaryStr := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_name" {
    name = %q
    grid_primary = %s
}
`, name, gridPrimaryStr)
}

func testAccNsgroupUseExternalPrimary(name string ,  gridPrimary []map[string]any, useExternalPrimary bool) string {
	gridPrimaryStr := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_use_external_primary" {
    name = %q
    grid_primary = %s
    use_external_primary = %t
}
`, name, gridPrimaryStr, useExternalPrimary)
}
func testAccNsgroupUseExternalPrimaryUpdate(name string, gridSecondaries, externalPrimaries []map[string]any, useExternalPrimary bool) string {
	gridSecondariesStr := utils.ConvertSliceOfMapsToHCL(gridSecondaries)
	externalPrimariesStr := utils.ConvertSliceOfMapsToHCL(externalPrimaries)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_use_external_primary" {
    name = %q
    grid_secondaries = %s
    external_primaries = %s
    use_external_primary = %t
}
`, name, gridSecondariesStr, externalPrimariesStr, useExternalPrimary)
}
