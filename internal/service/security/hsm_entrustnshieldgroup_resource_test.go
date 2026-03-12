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

// TODO:
//Register your Grid Master IP as a new client in the RFS interface.

// NOTE: The values keyServerIp, remoteIp, and keyhash are placeholders and must be replaced with real values for actual implementation.
// Remove 'disable = true' when running tests against an active HSM group.

var readableAttributesForHsmEntrustnshieldgroup = "card_name,comment,entrustnshield_hsm,key_server_ip,key_server_port,name,protection,status"

var keyServerIp = "10.10.10.10"
var keyServerPort = 9004
var keyhash = "keyhash-for-testing"
var remoteIp = "10.11.10.10"
var remotePort = 9004

var entrustnshieldHSM = []map[string]any{
	{
		"keyhash":     keyhash,
		"remote_ip":   remoteIp,
		"remote_port": remotePort,
		"disable":     true, //Remove this line when running tests against an active HSM group
	},
}

func TestAccHsmEntrustnshieldgroupResource_basic(t *testing.T) {
	var resourceName = "nios_security_hsm_entrustnshieldgroup.test"
	var v security.HsmEntrustnshieldgroup

	name := acctest.RandomNameWithPrefix("entrustnshieldgroup-hsm-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmEntrustnshieldgroupBasicConfig(name, keyServerIp, entrustnshieldHSM),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "key_server_ip", keyServerIp),
					resource.TestCheckResourceAttr(resourceName, "entrustnshield_hsm.0.remote_ip", remoteIp),
					resource.TestCheckResourceAttr(resourceName, "entrustnshield_hsm.0.remote_port", fmt.Sprintf("%d", remotePort)),
					resource.TestCheckResourceAttr(resourceName, "entrustnshield_hsm.0.keyhash", keyhash),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "protection", "MODULE"),
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "key_server_port", "9004"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmEntrustnshieldgroupResource_disappears(t *testing.T) {
	resourceName := "nios_security_hsm_entrustnshieldgroup.test"
	var v security.HsmEntrustnshieldgroup

	name := acctest.RandomNameWithPrefix("entrustnshieldgroup-hsm-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHsmEntrustnshieldgroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccHsmEntrustnshieldgroupBasicConfig(name, keyServerIp, entrustnshieldHSM),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					testAccCheckHsmEntrustnshieldgroupDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccHsmEntrustnshieldgroupResource_CardName(t *testing.T) {
	var resourceName = "nios_security_hsm_entrustnshieldgroup.test_card_name"
	var v security.HsmEntrustnshieldgroup

	name := acctest.RandomNameWithPrefix("entrustnshieldgroup-hsm-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmEntrustnshieldgroupCardName(name, "SOFTCARD", "example-softcard", "examplepassphrase@123", keyServerIp, keyServerPort, entrustnshieldHSM),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "card_name", "example-softcard"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmEntrustnshieldgroupCardName(name, "SOFTCARD", "example-softcard2", "updatedexamplepassphrase@123", keyServerIp, keyServerPort, entrustnshieldHSM),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "card_name", "example-softcard2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmEntrustnshieldgroupResource_Comment(t *testing.T) {
	var resourceName = "nios_security_hsm_entrustnshieldgroup.test_comment"
	var v security.HsmEntrustnshieldgroup

	name := acctest.RandomNameWithPrefix("entrustnshieldgroup-hsm-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmEntrustnshieldgroupComment(name, keyServerIp, entrustnshieldHSM, "sample comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "sample comment"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmEntrustnshieldgroupComment(name, keyServerIp, entrustnshieldHSM, "updated sample comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "updated sample comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmEntrustnshieldgroupResource_EntrustnshieldHsm(t *testing.T) {
	var resourceName = "nios_security_hsm_entrustnshieldgroup.test_entrustnshield_hsm"
	var v security.HsmEntrustnshieldgroup

	name := acctest.RandomNameWithPrefix("entrustnshieldgroup-hsm-")
	updatedentrustnshieldHSM := []map[string]any{
		{
			"keyhash":     keyhash,
			"remote_ip":   remoteIp,
			"remote_port": "9005",
			"disable":     true,
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmEntrustnshieldgroupEntrustnshieldHsm(name, keyServerIp, entrustnshieldHSM),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "entrustnshield_hsm.0.remote_port", "9004"),
					resource.TestCheckResourceAttr(resourceName, "entrustnshield_hsm.0.disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmEntrustnshieldgroupEntrustnshieldHsm(name, keyServerIp, updatedentrustnshieldHSM),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "entrustnshield_hsm.0.remote_port", "9005"),
					resource.TestCheckResourceAttr(resourceName, "entrustnshield_hsm.0.disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmEntrustnshieldgroupResource_KeyServerIp(t *testing.T) {
	var resourceName = "nios_security_hsm_entrustnshieldgroup.test_key_server_ip"
	var v security.HsmEntrustnshieldgroup

	name := acctest.RandomNameWithPrefix("entrustnshieldgroup-hsm-")
	updatedkeyServerIp := "10.10.10.15"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmEntrustnshieldgroupKeyServerIp(name, keyServerIp, entrustnshieldHSM),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "key_server_ip", keyServerIp),
				),
			},
			// Update and Read
			{
				Config: testAccHsmEntrustnshieldgroupKeyServerIp(name, updatedkeyServerIp, entrustnshieldHSM),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "key_server_ip", updatedkeyServerIp),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmEntrustnshieldgroupResource_KeyServerPort(t *testing.T) {
	var resourceName = "nios_security_hsm_entrustnshieldgroup.test_key_server_port"
	var v security.HsmEntrustnshieldgroup

	name := acctest.RandomNameWithPrefix("entrustnshieldgroup-hsm-")
	updatedKeyServerPort := 9005

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmEntrustnshieldgroupKeyServerPort(name, keyServerIp, keyServerPort, entrustnshieldHSM),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "key_server_port", "9004"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmEntrustnshieldgroupKeyServerPort(name, keyServerIp, updatedKeyServerPort, entrustnshieldHSM),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "key_server_port", "9005"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmEntrustnshieldgroupResource_Name(t *testing.T) {
	var resourceName = "nios_security_hsm_entrustnshieldgroup.test_name"
	var v security.HsmEntrustnshieldgroup

	name := acctest.RandomNameWithPrefix("entrustnshieldgroup-hsm-")
	updatedName := acctest.RandomNameWithPrefix("entrustnshieldgroup-hsm-1")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmEntrustnshieldgroupName(name, keyServerIp, entrustnshieldHSM),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccHsmEntrustnshieldgroupName(updatedName, keyServerIp, entrustnshieldHSM),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmEntrustnshieldgroupResource_PassPhrase(t *testing.T) {
	var resourceName = "nios_security_hsm_entrustnshieldgroup.test_pass_phrase"
	var v security.HsmEntrustnshieldgroup

	name := acctest.RandomNameWithPrefix("entrustnshieldgroup-hsm-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmEntrustnshieldgroupPassPhrase(name, "SOFTCARD", "example-softcard", "examplepassphrase@123", keyServerIp, 9004, entrustnshieldHSM),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pass_phrase", "examplepassphrase@123"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmEntrustnshieldgroupPassPhrase(name, "SOFTCARD", "example-softcard", "updatedpassphrase@123", keyServerIp, 9004, entrustnshieldHSM),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pass_phrase", "updatedpassphrase@123"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmEntrustnshieldgroupResource_Protection(t *testing.T) {
	var resourceName = "nios_security_hsm_entrustnshieldgroup.test_protection"
	var v security.HsmEntrustnshieldgroup

	name := acctest.RandomNameWithPrefix("entrustnshieldgroup-hsm-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmEntrustnshieldgroupProtectionSoftcard(name, "SOFTCARD", "example-softcard", "examplepassphrase@123", keyServerIp, keyServerPort, entrustnshieldHSM),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "protection", "SOFTCARD"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmEntrustnshieldgroupProtectionModule(name, "MODULE", keyServerIp, keyServerPort, entrustnshieldHSM),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmEntrustnshieldgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "protection", "MODULE"),
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

func testAccHsmEntrustnshieldgroupBasicConfig(name, keyServerIp string, entrustnshieldHSM []map[string]any) string {
	entrustnshieldHSM_HCL := utils.ConvertSliceOfMapsToHCL(entrustnshieldHSM)

	return fmt.Sprintf(`
resource "nios_security_hsm_entrustnshieldgroup" "test" {
    name = %q
    key_server_ip = %q
	entrustnshield_hsm = %s
}
`, name, keyServerIp, entrustnshieldHSM_HCL)
}

func testAccHsmEntrustnshieldgroupCardName(name, protection, cardName, passwordPhrase, keyServerIp string, keyServerPort int, entrustnshieldHSM []map[string]any) string {
	entrustnshieldHSM_HCL := utils.ConvertSliceOfMapsToHCL(entrustnshieldHSM)

	return fmt.Sprintf(`
resource "nios_security_hsm_entrustnshieldgroup" "test_card_name" {
	name = %q
	protection = %q
	card_name = %q
	pass_phrase = %q
    key_server_ip = %q
	key_server_port = %d
	entrustnshield_hsm = %s
}
`, name, protection, cardName, passwordPhrase, keyServerIp, keyServerPort, entrustnshieldHSM_HCL)
}

func testAccHsmEntrustnshieldgroupComment(name, keyServerIp string, entrustnshieldHSM []map[string]any, comment string) string {
	entrustnshieldHSM_HCL := utils.ConvertSliceOfMapsToHCL(entrustnshieldHSM)

	return fmt.Sprintf(`
resource "nios_security_hsm_entrustnshieldgroup" "test_comment" {
	name = %q
    key_server_ip = %q
	entrustnshield_hsm = %s
    comment = %q
}
`, name, keyServerIp, entrustnshieldHSM_HCL, comment)
}

func testAccHsmEntrustnshieldgroupEntrustnshieldHsm(name, keyServerIp string, entrustnshieldHsm []map[string]any) string {
	entrustnshieldHSM_HCL := utils.ConvertSliceOfMapsToHCL(entrustnshieldHsm)

	return fmt.Sprintf(`
resource "nios_security_hsm_entrustnshieldgroup" "test_entrustnshield_hsm" {
    name = %q
    key_server_ip = %q
    entrustnshield_hsm = %s
}
`, name, keyServerIp, entrustnshieldHSM_HCL)
}

func testAccHsmEntrustnshieldgroupKeyServerIp(name, keyServerIp string, entrustnshieldHsm []map[string]any) string {
	entrustnshieldHSM_HCL := utils.ConvertSliceOfMapsToHCL(entrustnshieldHsm)

	return fmt.Sprintf(`
resource "nios_security_hsm_entrustnshieldgroup" "test_key_server_ip" {
    name = %q
    key_server_ip = %q
    entrustnshield_hsm = %s
}
`, name, keyServerIp, entrustnshieldHSM_HCL)
}

func testAccHsmEntrustnshieldgroupKeyServerPort(name, keyServerIp string, keyServerPort int, entrustnshieldHsm []map[string]any) string {
	entrustnshieldHSM_HCL := utils.ConvertSliceOfMapsToHCL(entrustnshieldHsm)

	return fmt.Sprintf(`
resource "nios_security_hsm_entrustnshieldgroup" "test_key_server_port" {
    name = %q
    key_server_ip = %q
	key_server_port = %d
    entrustnshield_hsm = %s
}
`, name, keyServerIp, keyServerPort, entrustnshieldHSM_HCL)
}

func testAccHsmEntrustnshieldgroupName(name, keyServerIp string, entrustnshieldHsm []map[string]any) string {
	entrustnshieldHSM_HCL := utils.ConvertSliceOfMapsToHCL(entrustnshieldHsm)

	return fmt.Sprintf(`
resource "nios_security_hsm_entrustnshieldgroup" "test_name" {
    name = %q
    key_server_ip = %q
    entrustnshield_hsm = %s
}
`, name, keyServerIp, entrustnshieldHSM_HCL)
}

func testAccHsmEntrustnshieldgroupPassPhrase(name, protection, cardName, passwordPhrase, keyServerIp string, keyServerPort int, entrustnshieldHSM []map[string]any) string {
	entrustnshieldHSM_HCL := utils.ConvertSliceOfMapsToHCL(entrustnshieldHSM)

	return fmt.Sprintf(`
resource "nios_security_hsm_entrustnshieldgroup" "test_pass_phrase" {
   	name = %q
	protection = %q
	card_name = %q
	pass_phrase = %q
    key_server_ip = %q
	key_server_port = %d
	entrustnshield_hsm = %s
}
`, name, protection, cardName, passwordPhrase, keyServerIp, keyServerPort, entrustnshieldHSM_HCL)
}

func testAccHsmEntrustnshieldgroupProtectionModule(name, protection, keyServerIp string, keyServerPort int, entrustnshieldHSM []map[string]any) string {
	entrustnshieldHSM_HCL := utils.ConvertSliceOfMapsToHCL(entrustnshieldHSM)

	return fmt.Sprintf(`
resource "nios_security_hsm_entrustnshieldgroup" "test_protection" {
	name = %q
	protection = %q
    key_server_ip = %q
	key_server_port = %d
    entrustnshield_hsm = %s
}
`, name, protection, keyServerIp, keyServerPort, entrustnshieldHSM_HCL)
}

func testAccHsmEntrustnshieldgroupProtectionSoftcard(name, protection, cardName, passwordPhrase, keyServerIp string, keyServerPort int, entrustnshieldHSM []map[string]any) string {
	entrustnshieldHSM_HCL := utils.ConvertSliceOfMapsToHCL(entrustnshieldHSM)

	return fmt.Sprintf(`
resource "nios_security_hsm_entrustnshieldgroup" "test_protection" {
    name = %q
	protection = %q
	card_name = %q
	pass_phrase = %q
    key_server_ip = %q
	key_server_port = %d
	entrustnshield_hsm = %s
}
`, name, protection, cardName, passwordPhrase, keyServerIp, keyServerPort, entrustnshieldHSM_HCL)
}
