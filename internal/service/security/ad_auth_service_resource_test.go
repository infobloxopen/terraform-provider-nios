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

var readableAttributesForAdAuthService = "ad_domain,comment,disabled,domain_controllers,name,nested_group_querying,timeout"

func TestAccAdAuthServiceResource_basic(t *testing.T) {
	var resourceName = "nios_security_ad_auth_service.test"
	var v security.AdAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdAuthServiceBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdAuthServiceExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdAuthServiceResource_disappears(t *testing.T) {
	resourceName := "nios_security_ad_auth_service.test"
	var v security.AdAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAdAuthServiceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAdAuthServiceBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdAuthServiceExists(context.Background(), resourceName, &v),
					testAccCheckAdAuthServiceDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAdAuthServiceResource_Ref(t *testing.T) {
	var resourceName = "nios_security_ad_auth_service.test_ref"
	var v security.AdAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdAuthServiceRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdAuthServiceRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdAuthServiceResource_AdDomain(t *testing.T) {
	var resourceName = "nios_security_ad_auth_service.test_ad_domain"
	var v security.AdAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdAuthServiceAdDomain("AD_DOMAIN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ad_domain", "AD_DOMAIN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdAuthServiceAdDomain("AD_DOMAIN_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ad_domain", "AD_DOMAIN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdAuthServiceResource_Comment(t *testing.T) {
	var resourceName = "nios_security_ad_auth_service.test_comment"
	var v security.AdAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdAuthServiceComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdAuthServiceComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdAuthServiceResource_Disabled(t *testing.T) {
	var resourceName = "nios_security_ad_auth_service.test_disabled"
	var v security.AdAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdAuthServiceDisabled("DISABLED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "DISABLED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdAuthServiceDisabled("DISABLED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "DISABLED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdAuthServiceResource_DomainControllers(t *testing.T) {
	var resourceName = "nios_security_ad_auth_service.test_domain_controllers"
	var v security.AdAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdAuthServiceDomainControllers("DOMAIN_CONTROLLERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_controllers", "DOMAIN_CONTROLLERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdAuthServiceDomainControllers("DOMAIN_CONTROLLERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_controllers", "DOMAIN_CONTROLLERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdAuthServiceResource_Name(t *testing.T) {
	var resourceName = "nios_security_ad_auth_service.test_name"
	var v security.AdAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdAuthServiceName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdAuthServiceName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdAuthServiceResource_NestedGroupQuerying(t *testing.T) {
	var resourceName = "nios_security_ad_auth_service.test_nested_group_querying"
	var v security.AdAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdAuthServiceNestedGroupQuerying("NESTED_GROUP_QUERYING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nested_group_querying", "NESTED_GROUP_QUERYING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdAuthServiceNestedGroupQuerying("NESTED_GROUP_QUERYING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nested_group_querying", "NESTED_GROUP_QUERYING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdAuthServiceResource_Timeout(t *testing.T) {
	var resourceName = "nios_security_ad_auth_service.test_timeout"
	var v security.AdAuthService

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdAuthServiceTimeout("TIMEOUT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "TIMEOUT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdAuthServiceTimeout("TIMEOUT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "TIMEOUT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckAdAuthServiceExists(ctx context.Context, resourceName string, v *security.AdAuthService) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SecurityAPI.
			AdAuthServiceAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForAdAuthService).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetAdAuthServiceResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetAdAuthServiceResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckAdAuthServiceDestroy(ctx context.Context, v *security.AdAuthService) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SecurityAPI.
			AdAuthServiceAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForAdAuthService).
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

func testAccCheckAdAuthServiceDisappears(ctx context.Context, v *security.AdAuthService) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SecurityAPI.
			AdAuthServiceAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccAdAuthServiceBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return `
resource "nios_security_ad_auth_service" "test" {
}
`
}

func testAccAdAuthServiceRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_security_ad_auth_service" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccAdAuthServiceAdDomain(adDomain string) string {
	return fmt.Sprintf(`
resource "nios_security_ad_auth_service" "test_ad_domain" {
    ad_domain = %q
}
`, adDomain)
}

func testAccAdAuthServiceComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_security_ad_auth_service" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccAdAuthServiceDisabled(disabled string) string {
	return fmt.Sprintf(`
resource "nios_security_ad_auth_service" "test_disabled" {
    disabled = %q
}
`, disabled)
}

func testAccAdAuthServiceDomainControllers(domainControllers string) string {
	return fmt.Sprintf(`
resource "nios_security_ad_auth_service" "test_domain_controllers" {
    domain_controllers = %q
}
`, domainControllers)
}

func testAccAdAuthServiceName(name string) string {
	return fmt.Sprintf(`
resource "nios_security_ad_auth_service" "test_name" {
    name = %q
}
`, name)
}

func testAccAdAuthServiceNestedGroupQuerying(nestedGroupQuerying string) string {
	return fmt.Sprintf(`
resource "nios_security_ad_auth_service" "test_nested_group_querying" {
    nested_group_querying = %q
}
`, nestedGroupQuerying)
}

func testAccAdAuthServiceTimeout(timeout string) string {
	return fmt.Sprintf(`
resource "nios_security_ad_auth_service" "test_timeout" {
    timeout = %q
}
`, timeout)
}
