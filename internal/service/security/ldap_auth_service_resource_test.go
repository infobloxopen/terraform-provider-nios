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

var readableAttributesForLdapAuthService = "comment,disable,ea_mapping,ldap_group_attribute,ldap_group_authentication_type,ldap_user_attribute,mode,name,recovery_interval,retries,search_scope,servers,timeout"

func TestAccLdapAuthServiceResource_basic(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test"
	var v security.LdapAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_disappears(t *testing.T) {
	resourceName := "nios_security_ldap_auth_service.test"
	var v security.LdapAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLdapAuthServiceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccLdapAuthServiceBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					testAccCheckLdapAuthServiceDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccLdapAuthServiceResource_Ref(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_ref"
	var v security.LdapAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_Comment(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_comment"
	var v security.LdapAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_Disable(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_disable"
	var v security.LdapAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_EaMapping(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_ea_mapping"
	var v security.LdapAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceEaMapping("EA_MAPPING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ea_mapping", "EA_MAPPING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceEaMapping("EA_MAPPING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ea_mapping", "EA_MAPPING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_LdapGroupAttribute(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_ldap_group_attribute"
	var v security.LdapAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceLdapGroupAttribute("LDAP_GROUP_ATTRIBUTE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ldap_group_attribute", "LDAP_GROUP_ATTRIBUTE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceLdapGroupAttribute("LDAP_GROUP_ATTRIBUTE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ldap_group_attribute", "LDAP_GROUP_ATTRIBUTE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_LdapGroupAuthenticationType(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_ldap_group_authentication_type"
	var v security.LdapAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceLdapGroupAuthenticationType("LDAP_GROUP_AUTHENTICATION_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ldap_group_authentication_type", "LDAP_GROUP_AUTHENTICATION_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceLdapGroupAuthenticationType("LDAP_GROUP_AUTHENTICATION_TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ldap_group_authentication_type", "LDAP_GROUP_AUTHENTICATION_TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_LdapUserAttribute(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_ldap_user_attribute"
	var v security.LdapAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceLdapUserAttribute("LDAP_USER_ATTRIBUTE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ldap_user_attribute", "LDAP_USER_ATTRIBUTE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceLdapUserAttribute("LDAP_USER_ATTRIBUTE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ldap_user_attribute", "LDAP_USER_ATTRIBUTE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_Mode(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_mode"
	var v security.LdapAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceMode("MODE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mode", "MODE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceMode("MODE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mode", "MODE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_Name(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_name"
	var v security.LdapAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_RecoveryInterval(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_recovery_interval"
	var v security.LdapAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceRecoveryInterval("RECOVERY_INTERVAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recovery_interval", "RECOVERY_INTERVAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceRecoveryInterval("RECOVERY_INTERVAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recovery_interval", "RECOVERY_INTERVAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_Retries(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_retries"
	var v security.LdapAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceRetries("RETRIES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retries", "RETRIES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceRetries("RETRIES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retries", "RETRIES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_SearchScope(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_search_scope"
	var v security.LdapAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceSearchScope("SEARCH_SCOPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "search_scope", "SEARCH_SCOPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceSearchScope("SEARCH_SCOPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "search_scope", "SEARCH_SCOPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_Servers(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_servers"
	var v security.LdapAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceServers("SERVERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "servers", "SERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceServers("SERVERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "servers", "SERVERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_Timeout(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_timeout"
	var v security.LdapAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceTimeout("TIMEOUT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "TIMEOUT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceTimeout("TIMEOUT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "TIMEOUT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckLdapAuthServiceExists(ctx context.Context, resourceName string, v *security.LdapAuthService) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SecurityAPI.
			LdapAuthServiceAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForLdapAuthService).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetLdapAuthServiceResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetLdapAuthServiceResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckLdapAuthServiceDestroy(ctx context.Context, v *security.LdapAuthService) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SecurityAPI.
			LdapAuthServiceAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForLdapAuthService).
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

func testAccCheckLdapAuthServiceDisappears(ctx context.Context, v *security.LdapAuthService) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SecurityAPI.
			LdapAuthServiceAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccLdapAuthServiceBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test" {
}
`)
}

func testAccLdapAuthServiceRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccLdapAuthServiceComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccLdapAuthServiceDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccLdapAuthServiceEaMapping(eaMapping string) string {
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_ea_mapping" {
    ea_mapping = %q
}
`, eaMapping)
}

func testAccLdapAuthServiceLdapGroupAttribute(ldapGroupAttribute string) string {
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_ldap_group_attribute" {
    ldap_group_attribute = %q
}
`, ldapGroupAttribute)
}

func testAccLdapAuthServiceLdapGroupAuthenticationType(ldapGroupAuthenticationType string) string {
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_ldap_group_authentication_type" {
    ldap_group_authentication_type = %q
}
`, ldapGroupAuthenticationType)
}

func testAccLdapAuthServiceLdapUserAttribute(ldapUserAttribute string) string {
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_ldap_user_attribute" {
    ldap_user_attribute = %q
}
`, ldapUserAttribute)
}

func testAccLdapAuthServiceMode(mode string) string {
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_mode" {
    mode = %q
}
`, mode)
}

func testAccLdapAuthServiceName(name string) string {
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_name" {
    name = %q
}
`, name)
}

func testAccLdapAuthServiceRecoveryInterval(recoveryInterval string) string {
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_recovery_interval" {
    recovery_interval = %q
}
`, recoveryInterval)
}

func testAccLdapAuthServiceRetries(retries string) string {
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_retries" {
    retries = %q
}
`, retries)
}

func testAccLdapAuthServiceSearchScope(searchScope string) string {
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_search_scope" {
    search_scope = %q
}
`, searchScope)
}

func testAccLdapAuthServiceServers(servers string) string {
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_servers" {
    servers = %q
}
`, servers)
}

func testAccLdapAuthServiceTimeout(timeout string) string {
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_timeout" {
    timeout = %q
}
`, timeout)
}
