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

var readableAttributesForSamlAuthservice = "comment,idp,name,session_timeout"

func TestAccSamlAuthserviceResource_basic(t *testing.T) {
	var resourceName = "nios_security_saml_authservice.test"
	var v security.SamlAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSamlAuthserviceBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSamlAuthserviceResource_disappears(t *testing.T) {
	resourceName := "nios_security_saml_authservice.test"
	var v security.SamlAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSamlAuthserviceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSamlAuthserviceBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					testAccCheckSamlAuthserviceDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSamlAuthserviceResource_Ref(t *testing.T) {
	var resourceName = "nios_security_saml_authservice.test_ref"
	var v security.SamlAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSamlAuthserviceRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSamlAuthserviceRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSamlAuthserviceResource_Comment(t *testing.T) {
	var resourceName = "nios_security_saml_authservice.test_comment"
	var v security.SamlAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSamlAuthserviceComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSamlAuthserviceComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSamlAuthserviceResource_Idp(t *testing.T) {
	var resourceName = "nios_security_saml_authservice.test_idp"
	var v security.SamlAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSamlAuthserviceIdp("IDP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "idp", "IDP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSamlAuthserviceIdp("IDP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "idp", "IDP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSamlAuthserviceResource_Name(t *testing.T) {
	var resourceName = "nios_security_saml_authservice.test_name"
	var v security.SamlAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSamlAuthserviceName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSamlAuthserviceName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSamlAuthserviceResource_SessionTimeout(t *testing.T) {
	var resourceName = "nios_security_saml_authservice.test_session_timeout"
	var v security.SamlAuthservice

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSamlAuthserviceSessionTimeout("SESSION_TIMEOUT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "session_timeout", "SESSION_TIMEOUT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSamlAuthserviceSessionTimeout("SESSION_TIMEOUT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "session_timeout", "SESSION_TIMEOUT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckSamlAuthserviceExists(ctx context.Context, resourceName string, v *security.SamlAuthservice) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SecurityAPI.
			SamlAuthserviceAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForSamlAuthservice).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetSamlAuthserviceResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetSamlAuthserviceResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckSamlAuthserviceDestroy(ctx context.Context, v *security.SamlAuthservice) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SecurityAPI.
			SamlAuthserviceAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForSamlAuthservice).
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

func testAccCheckSamlAuthserviceDisappears(ctx context.Context, v *security.SamlAuthservice) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SecurityAPI.
			SamlAuthserviceAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccSamlAuthserviceBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_security_saml_authservice" "test" {
}
`)
}

func testAccSamlAuthserviceRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_security_saml_authservice" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccSamlAuthserviceComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_security_saml_authservice" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccSamlAuthserviceIdp(idp string) string {
	return fmt.Sprintf(`
resource "nios_security_saml_authservice" "test_idp" {
    idp = %q
}
`, idp)
}

func testAccSamlAuthserviceName(name string) string {
	return fmt.Sprintf(`
resource "nios_security_saml_authservice" "test_name" {
    name = %q
}
`, name)
}

func testAccSamlAuthserviceSessionTimeout(sessionTimeout string) string {
	return fmt.Sprintf(`
resource "nios_security_saml_authservice" "test_session_timeout" {
    session_timeout = %q
}
`, sessionTimeout)
}
