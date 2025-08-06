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

var readableAttributesForNsgroup = "comment,extattrs,external_primaries,external_secondaries,grid_primary,grid_secondaries,is_grid_default,is_multimaster,name,use_external_primary"

func TestAccNsgroupResource_basic(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test"
	var v dns.Nsgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupResource_disappears(t *testing.T) {
	resourceName := "nios_dns_nsgroup.test"
	var v dns.Nsgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					testAccCheckNsgroupDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNsgroupResource_Ref(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test_ref"
	var v dns.Nsgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test_comment"
	var v dns.Nsgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test_extattrs"
	var v dns.Nsgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupResource_ExternalPrimaries(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test_external_primaries"
	var v dns.Nsgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupExternalPrimaries("EXTERNAL_PRIMARIES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_primaries", "EXTERNAL_PRIMARIES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupExternalPrimaries("EXTERNAL_PRIMARIES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_primaries", "EXTERNAL_PRIMARIES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupResource_ExternalSecondaries(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test_external_secondaries"
	var v dns.Nsgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupExternalSecondaries("EXTERNAL_SECONDARIES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries", "EXTERNAL_SECONDARIES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupExternalSecondaries("EXTERNAL_SECONDARIES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries", "EXTERNAL_SECONDARIES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupResource_GridPrimary(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test_grid_primary"
	var v dns.Nsgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupGridPrimary("GRID_PRIMARY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_primary", "GRID_PRIMARY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupGridPrimary("GRID_PRIMARY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_primary", "GRID_PRIMARY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupResource_GridSecondaries(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test_grid_secondaries"
	var v dns.Nsgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupGridSecondaries("GRID_SECONDARIES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries", "GRID_SECONDARIES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupGridSecondaries("GRID_SECONDARIES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries", "GRID_SECONDARIES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupResource_IsGridDefault(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test_is_grid_default"
	var v dns.Nsgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupIsGridDefault("IS_GRID_DEFAULT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "is_grid_default", "IS_GRID_DEFAULT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupIsGridDefault("IS_GRID_DEFAULT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "is_grid_default", "IS_GRID_DEFAULT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupResource_IsMultimaster(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test_is_multimaster"
	var v dns.Nsgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupIsMultimaster("IS_MULTIMASTER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "is_multimaster", "IS_MULTIMASTER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupIsMultimaster("IS_MULTIMASTER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "is_multimaster", "IS_MULTIMASTER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupResource_Name(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test_name"
	var v dns.Nsgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupResource_UseExternalPrimary(t *testing.T) {
	var resourceName = "nios_dns_nsgroup.test_use_external_primary"
	var v dns.Nsgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupUseExternalPrimary("USE_EXTERNAL_PRIMARY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_external_primary", "USE_EXTERNAL_PRIMARY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupUseExternalPrimary("USE_EXTERNAL_PRIMARY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_external_primary", "USE_EXTERNAL_PRIMARY_UPDATE_REPLACE_ME"),
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

func testAccNsgroupBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test" {
}
`)
}

func testAccNsgroupRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccNsgroupComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccNsgroupExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccNsgroupExternalPrimaries(externalPrimaries string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_external_primaries" {
    external_primaries = %q
}
`, externalPrimaries)
}

func testAccNsgroupExternalSecondaries(externalSecondaries string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_external_secondaries" {
    external_secondaries = %q
}
`, externalSecondaries)
}

func testAccNsgroupGridPrimary(gridPrimary string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_grid_primary" {
    grid_primary = %q
}
`, gridPrimary)
}

func testAccNsgroupGridSecondaries(gridSecondaries string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_grid_secondaries" {
    grid_secondaries = %q
}
`, gridSecondaries)
}

func testAccNsgroupIsGridDefault(isGridDefault string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_is_grid_default" {
    is_grid_default = %q
}
`, isGridDefault)
}

func testAccNsgroupIsMultimaster(isMultimaster string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_is_multimaster" {
    is_multimaster = %q
}
`, isMultimaster)
}

func testAccNsgroupName(name string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_name" {
    name = %q
}
`, name)
}

func testAccNsgroupUseExternalPrimary(useExternalPrimary string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test_use_external_primary" {
    use_external_primary = %q
}
`, useExternalPrimary)
}
