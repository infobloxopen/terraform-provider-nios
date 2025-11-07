package security_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForHsmThaleslunagroup = "comment,group_sn,hsm_version,name,status,thalesluna"

func TestAccHsmThaleslunagroupResource_basic(t *testing.T) {
	var resourceName = "nios_security_hsm_thaleslunagroup.test"
	var v security.HsmThaleslunagroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmThaleslunagroupBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmThaleslunagroupResource_disappears(t *testing.T) {
	resourceName := "nios_security_hsm_thaleslunagroup.test"
	var v security.HsmThaleslunagroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHsmThaleslunagroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccHsmThaleslunagroupBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					testAccCheckHsmThaleslunagroupDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccHsmThaleslunagroupResource_Ref(t *testing.T) {
	var resourceName = "nios_security_hsm_thaleslunagroup.test_ref"
	var v security.HsmThaleslunagroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmThaleslunagroupRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmThaleslunagroupRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmThaleslunagroupResource_Comment(t *testing.T) {
	var resourceName = "nios_security_hsm_thaleslunagroup.test_comment"
	var v security.HsmThaleslunagroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmThaleslunagroupComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmThaleslunagroupComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmThaleslunagroupResource_HsmVersion(t *testing.T) {
	var resourceName = "nios_security_hsm_thaleslunagroup.test_hsm_version"
	var v security.HsmThaleslunagroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmThaleslunagroupHsmVersion("HSM_VERSION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hsm_version", "HSM_VERSION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmThaleslunagroupHsmVersion("HSM_VERSION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hsm_version", "HSM_VERSION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmThaleslunagroupResource_Name(t *testing.T) {
	var resourceName = "nios_security_hsm_thaleslunagroup.test_name"
	var v security.HsmThaleslunagroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmThaleslunagroupName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmThaleslunagroupName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmThaleslunagroupResource_PassPhrase(t *testing.T) {
	var resourceName = "nios_security_hsm_thaleslunagroup.test_pass_phrase"
	var v security.HsmThaleslunagroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmThaleslunagroupPassPhrase("PASS_PHRASE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pass_phrase", "PASS_PHRASE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmThaleslunagroupPassPhrase("PASS_PHRASE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pass_phrase", "PASS_PHRASE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmThaleslunagroupResource_Thalesluna(t *testing.T) {
	var resourceName = "nios_security_hsm_thaleslunagroup.test_thalesluna"
	var v security.HsmThaleslunagroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmThaleslunagroupThalesluna("THALESLUNA_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "thalesluna", "THALESLUNA_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmThaleslunagroupThalesluna("THALESLUNA_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "thalesluna", "THALESLUNA_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckHsmThaleslunagroupExists(ctx context.Context, resourceName string, v *security.HsmThaleslunagroup) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SecurityAPI.
			HsmThaleslunagroupAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForHsmThaleslunagroup).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetHsmThaleslunagroupResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetHsmThaleslunagroupResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckHsmThaleslunagroupDestroy(ctx context.Context, v *security.HsmThaleslunagroup) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SecurityAPI.
			HsmThaleslunagroupAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForHsmThaleslunagroup).
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

func testAccCheckHsmThaleslunagroupDisappears(ctx context.Context, v *security.HsmThaleslunagroup) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SecurityAPI.
			HsmThaleslunagroupAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccHsmThaleslunagroupBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_security_hsm_thaleslunagroup" "test" {
}
`)
}

func testAccHsmThaleslunagroupRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_security_hsm_thaleslunagroup" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccHsmThaleslunagroupComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_security_hsm_thaleslunagroup" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccHsmThaleslunagroupHsmVersion(hsmVersion string) string {
	return fmt.Sprintf(`
resource "nios_security_hsm_thaleslunagroup" "test_hsm_version" {
    hsm_version = %q
}
`, hsmVersion)
}

func testAccHsmThaleslunagroupName(name string) string {
	return fmt.Sprintf(`
resource "nios_security_hsm_thaleslunagroup" "test_name" {
    name = %q
}
`, name)
}

func testAccHsmThaleslunagroupPassPhrase(passPhrase string) string {
	return fmt.Sprintf(`
resource "nios_security_hsm_thaleslunagroup" "test_pass_phrase" {
    pass_phrase = %q
}
`, passPhrase)
}

func testAccHsmThaleslunagroupThalesluna(thalesluna string) string {
	return fmt.Sprintf(`
resource "nios_security_hsm_thaleslunagroup" "test_thalesluna" {
    thalesluna = %q
}
`, thalesluna)
}
