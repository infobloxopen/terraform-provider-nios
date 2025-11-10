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

var readableAttributesForRadiusAuthservice = "acct_retries,acct_timeout,auth_retries,auth_timeout,cache_ttl,comment,disable,enable_cache,mode,name,recovery_interval,servers"

func TestAccRadiusAuthserviceResource_basic(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test"
	var v security.RadiusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_disappears(t *testing.T) {
	resourceName := "nios_security_radius_authservice.test"
	var v security.RadiusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRadiusAuthserviceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRadiusAuthserviceBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					testAccCheckRadiusAuthserviceDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRadiusAuthserviceResource_Ref(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_ref"
	var v security.RadiusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_AcctRetries(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_acct_retries"
	var v security.RadiusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceAcctRetries("ACCT_RETRIES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "acct_retries", "ACCT_RETRIES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceAcctRetries("ACCT_RETRIES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "acct_retries", "ACCT_RETRIES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_AcctTimeout(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_acct_timeout"
	var v security.RadiusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceAcctTimeout("ACCT_TIMEOUT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "acct_timeout", "ACCT_TIMEOUT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceAcctTimeout("ACCT_TIMEOUT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "acct_timeout", "ACCT_TIMEOUT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_AuthRetries(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_auth_retries"
	var v security.RadiusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceAuthRetries("AUTH_RETRIES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_retries", "AUTH_RETRIES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceAuthRetries("AUTH_RETRIES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_retries", "AUTH_RETRIES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_AuthTimeout(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_auth_timeout"
	var v security.RadiusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceAuthTimeout("AUTH_TIMEOUT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_timeout", "AUTH_TIMEOUT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceAuthTimeout("AUTH_TIMEOUT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_timeout", "AUTH_TIMEOUT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_CacheTtl(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_cache_ttl"
	var v security.RadiusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceCacheTtl("CACHE_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cache_ttl", "CACHE_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceCacheTtl("CACHE_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cache_ttl", "CACHE_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_Comment(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_comment"
	var v security.RadiusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_Disable(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_disable"
	var v security.RadiusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_EnableCache(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_enable_cache"
	var v security.RadiusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceEnableCache("ENABLE_CACHE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_cache", "ENABLE_CACHE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceEnableCache("ENABLE_CACHE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_cache", "ENABLE_CACHE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_Mode(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_mode"
	var v security.RadiusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceMode("MODE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mode", "MODE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceMode("MODE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mode", "MODE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_Name(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_name"
	var v security.RadiusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_RecoveryInterval(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_recovery_interval"
	var v security.RadiusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceRecoveryInterval("RECOVERY_INTERVAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recovery_interval", "RECOVERY_INTERVAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceRecoveryInterval("RECOVERY_INTERVAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recovery_interval", "RECOVERY_INTERVAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_Servers(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_servers"
	var v security.RadiusAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceServers("SERVERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "servers", "SERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceServers("SERVERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "servers", "SERVERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRadiusAuthserviceExists(ctx context.Context, resourceName string, v *security.RadiusAuthservice) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SecurityAPI.
			RadiusAuthserviceAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRadiusAuthservice).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRadiusAuthserviceResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRadiusAuthserviceResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRadiusAuthserviceDestroy(ctx context.Context, v *security.RadiusAuthservice) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SecurityAPI.
			RadiusAuthserviceAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRadiusAuthservice).
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

func testAccCheckRadiusAuthserviceDisappears(ctx context.Context, v *security.RadiusAuthservice) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SecurityAPI.
			RadiusAuthserviceAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRadiusAuthserviceBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return `
resource "nios_security_radius_authservice" "test" {
}
`
}

func testAccRadiusAuthserviceRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccRadiusAuthserviceAcctRetries(acctRetries string) string {
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_acct_retries" {
    acct_retries = %q
}
`, acctRetries)
}

func testAccRadiusAuthserviceAcctTimeout(acctTimeout string) string {
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_acct_timeout" {
    acct_timeout = %q
}
`, acctTimeout)
}

func testAccRadiusAuthserviceAuthRetries(authRetries string) string {
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_auth_retries" {
    auth_retries = %q
}
`, authRetries)
}

func testAccRadiusAuthserviceAuthTimeout(authTimeout string) string {
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_auth_timeout" {
    auth_timeout = %q
}
`, authTimeout)
}

func testAccRadiusAuthserviceCacheTtl(cacheTtl string) string {
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_cache_ttl" {
    cache_ttl = %q
}
`, cacheTtl)
}

func testAccRadiusAuthserviceComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccRadiusAuthserviceDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccRadiusAuthserviceEnableCache(enableCache string) string {
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_enable_cache" {
    enable_cache = %q
}
`, enableCache)
}

func testAccRadiusAuthserviceMode(mode string) string {
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_mode" {
    mode = %q
}
`, mode)
}

func testAccRadiusAuthserviceName(name string) string {
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_name" {
    name = %q
}
`, name)
}

func testAccRadiusAuthserviceRecoveryInterval(recoveryInterval string) string {
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_recovery_interval" {
    recovery_interval = %q
}
`, recoveryInterval)
}

func testAccRadiusAuthserviceServers(servers string) string {
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_servers" {
    servers = %q
}
`, servers)
}
