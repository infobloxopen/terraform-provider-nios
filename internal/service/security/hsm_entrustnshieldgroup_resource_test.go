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

var readableAttributesForHsmEntrustnshieldgroup = "card_name,comment,entrustnshield_hsm,key_server_ip,key_server_port,name,protection,status"

func TestAccHsmEntrustnshieldgroupResource_basic(t *testing.T) {
	var resourceName = "nios_security_hsm_entrustnshieldgroup.test"
	var v security.HsmEntrustnshieldgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmEntrustnshieldgroupBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmEntrustnshieldgroupResource_disappears(t *testing.T) {
	resourceName := "nios_security_hsm_entrustnshieldgroup.test"
	var v security.HsmEntrustnshieldgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHsmEntrustnshieldgroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccHsmEntrustnshieldgroupBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					testAccCheckHsmEntrustnshieldgroupDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccHsmEntrustnshieldgroupResource_Ref(t *testing.T) {
	var resourceName = "nios_security_hsm_entrustnshieldgroup.test_ref"
	var v security.HsmEntrustnshieldgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmEntrustnshieldgroupRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmEntrustnshieldgroupRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmEntrustnshieldgroupResource_CardName(t *testing.T) {
	var resourceName = "nios_security_hsm_entrustnshieldgroup.test_card_name"
	var v security.HsmEntrustnshieldgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmEntrustnshieldgroupCardName("CARD_NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "card_name", "CARD_NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmEntrustnshieldgroupCardName("CARD_NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "card_name", "CARD_NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmEntrustnshieldgroupResource_Comment(t *testing.T) {
	var resourceName = "nios_security_hsm_entrustnshieldgroup.test_comment"
	var v security.HsmEntrustnshieldgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmEntrustnshieldgroupComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmEntrustnshieldgroupComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmEntrustnshieldgroupResource_EntrustnshieldHsm(t *testing.T) {
	var resourceName = "nios_security_hsm_entrustnshieldgroup.test_entrustnshield_hsm"
	var v security.HsmEntrustnshieldgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmEntrustnshieldgroupEntrustnshieldHsm("ENTRUSTNSHIELD_HSM_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "entrustnshield_hsm", "ENTRUSTNSHIELD_HSM_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmEntrustnshieldgroupEntrustnshieldHsm("ENTRUSTNSHIELD_HSM_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "entrustnshield_hsm", "ENTRUSTNSHIELD_HSM_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmEntrustnshieldgroupResource_KeyServerIp(t *testing.T) {
	var resourceName = "nios_security_hsm_entrustnshieldgroup.test_key_server_ip"
	var v security.HsmEntrustnshieldgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmEntrustnshieldgroupKeyServerIp("KEY_SERVER_IP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "key_server_ip", "KEY_SERVER_IP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmEntrustnshieldgroupKeyServerIp("KEY_SERVER_IP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "key_server_ip", "KEY_SERVER_IP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmEntrustnshieldgroupResource_KeyServerPort(t *testing.T) {
	var resourceName = "nios_security_hsm_entrustnshieldgroup.test_key_server_port"
	var v security.HsmEntrustnshieldgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmEntrustnshieldgroupKeyServerPort("KEY_SERVER_PORT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "key_server_port", "KEY_SERVER_PORT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmEntrustnshieldgroupKeyServerPort("KEY_SERVER_PORT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "key_server_port", "KEY_SERVER_PORT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmEntrustnshieldgroupResource_Name(t *testing.T) {
	var resourceName = "nios_security_hsm_entrustnshieldgroup.test_name"
	var v security.HsmEntrustnshieldgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmEntrustnshieldgroupName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmEntrustnshieldgroupName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmEntrustnshieldgroupResource_PassPhrase(t *testing.T) {
	var resourceName = "nios_security_hsm_entrustnshieldgroup.test_pass_phrase"
	var v security.HsmEntrustnshieldgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmEntrustnshieldgroupPassPhrase("PASS_PHRASE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pass_phrase", "PASS_PHRASE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmEntrustnshieldgroupPassPhrase("PASS_PHRASE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pass_phrase", "PASS_PHRASE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmEntrustnshieldgroupResource_Protection(t *testing.T) {
	var resourceName = "nios_security_hsm_entrustnshieldgroup.test_protection"
	var v security.HsmEntrustnshieldgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmEntrustnshieldgroupProtection("PROTECTION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "protection", "PROTECTION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmEntrustnshieldgroupProtection("PROTECTION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "protection", "PROTECTION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckHsmEntrustnshieldgroupExists(ctx context.Context, resourceName string, v *security.HsmEntrustnshieldgroup) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SecurityAPI.
			HsmEntrustnshieldgroupAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForHsmEntrustnshieldgroup).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetHsmEntrustnshieldgroupResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetHsmEntrustnshieldgroupResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckHsmEntrustnshieldgroupDestroy(ctx context.Context, v *security.HsmEntrustnshieldgroup) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SecurityAPI.
			HsmEntrustnshieldgroupAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForHsmEntrustnshieldgroup).
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

func testAccCheckHsmEntrustnshieldgroupDisappears(ctx context.Context, v *security.HsmEntrustnshieldgroup) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SecurityAPI.
			HsmEntrustnshieldgroupAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccHsmEntrustnshieldgroupBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return `
resource "nios_security_hsm_entrustnshieldgroup" "test" {
}
`
}

func testAccHsmEntrustnshieldgroupRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_security_hsm_entrustnshieldgroup" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccHsmEntrustnshieldgroupCardName(cardName string) string {
	return fmt.Sprintf(`
resource "nios_security_hsm_entrustnshieldgroup" "test_card_name" {
    card_name = %q
}
`, cardName)
}

func testAccHsmEntrustnshieldgroupComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_security_hsm_entrustnshieldgroup" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccHsmEntrustnshieldgroupEntrustnshieldHsm(entrustnshieldHsm string) string {
	return fmt.Sprintf(`
resource "nios_security_hsm_entrustnshieldgroup" "test_entrustnshield_hsm" {
    entrustnshield_hsm = %q
}
`, entrustnshieldHsm)
}

func testAccHsmEntrustnshieldgroupKeyServerIp(keyServerIp string) string {
	return fmt.Sprintf(`
resource "nios_security_hsm_entrustnshieldgroup" "test_key_server_ip" {
    key_server_ip = %q
}
`, keyServerIp)
}

func testAccHsmEntrustnshieldgroupKeyServerPort(keyServerPort string) string {
	return fmt.Sprintf(`
resource "nios_security_hsm_entrustnshieldgroup" "test_key_server_port" {
    key_server_port = %q
}
`, keyServerPort)
}

func testAccHsmEntrustnshieldgroupName(name string) string {
	return fmt.Sprintf(`
resource "nios_security_hsm_entrustnshieldgroup" "test_name" {
    name = %q
}
`, name)
}

func testAccHsmEntrustnshieldgroupPassPhrase(passPhrase string) string {
	return fmt.Sprintf(`
resource "nios_security_hsm_entrustnshieldgroup" "test_pass_phrase" {
    pass_phrase = %q
}
`, passPhrase)
}

func testAccHsmEntrustnshieldgroupProtection(protection string) string {
	return fmt.Sprintf(`
resource "nios_security_hsm_entrustnshieldgroup" "test_protection" {
    protection = %q
}
`, protection)
}
