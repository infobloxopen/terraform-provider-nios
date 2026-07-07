package security_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
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
	name := acctest.RandomNameWithPrefix("saml_authservice")
	testDataPath := getSamlTestDataPath()
	idp := map[string]any{
		"idp_type":           "AZURE_SSO",
		"metadata_file_path": filepath.Join(testDataPath, "metadata.xml"),
		"sso_redirect_url":   "2.2.2.2",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSamlAuthserviceBasicConfig(name, idp),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "session_timeout", "1800"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSamlAuthserviceResource_disappears(t *testing.T) {
	resourceName := "nios_security_saml_authservice.test"
	var v security.SamlAuthservice
	name := acctest.RandomNameWithPrefix("saml_authservice")
	testDataPath := getSamlTestDataPath()
	idp := map[string]any{
		"idp_type":           "AZURE_SSO",
		"metadata_file_path": filepath.Join(testDataPath, "metadata.xml"),
		"sso_redirect_url":   "2.2.2.2",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSamlAuthserviceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSamlAuthserviceBasicConfig(name, idp),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					testAccCheckSamlAuthserviceDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSamlAuthserviceResource_Comment(t *testing.T) {
	var resourceName = "nios_security_saml_authservice.test_comment"
	var v security.SamlAuthservice
	name := acctest.RandomNameWithPrefix("saml_authservice")
	testDataPath := getSamlTestDataPath()
	idp := map[string]any{
		"idp_type":           "AZURE_SSO",
		"metadata_file_path": filepath.Join(testDataPath, "metadata.xml"),
		"sso_redirect_url":   "2.2.2.2",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSamlAuthserviceComment(name, "This is a comment", idp),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccSamlAuthserviceComment(name, "This comment is updated", idp),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This comment is updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSamlAuthserviceResource_Idp(t *testing.T) {
	var resourceName = "nios_security_saml_authservice.test_idp"
	var v security.SamlAuthservice
	name := acctest.RandomNameWithPrefix("saml_authservice")
	testDataPath := getSamlTestDataPath()
	idp := map[string]any{
		"idp_type":           "AZURE_SSO",
		"metadata_file_path": filepath.Join(testDataPath, "metadata.xml"),
		"sso_redirect_url":   "2.2.2.2",
	}
	idpUpdate := map[string]any{
		"idp_type":         "OKTA",
		"metadata_url":     "https://idp.example.com",
		"sso_redirect_url": "2.2.2.1",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSamlAuthserviceIdp(name, idp),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "idp.idp_type", "AZURE_SSO"),
					resource.TestCheckResourceAttr(resourceName, "idp.sso_redirect_url", "2.2.2.2"),
				),
			},
			// Update and Read
			{
				Config: testAccSamlAuthserviceIdp(name, idpUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "idp.idp_type", "OKTA"),
					resource.TestCheckResourceAttr(resourceName, "idp.metadata_url", "https://idp.example.com"),
					resource.TestCheckResourceAttr(resourceName, "idp.sso_redirect_url", "2.2.2.1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSamlAuthserviceResource_Name(t *testing.T) {
	var resourceName = "nios_security_saml_authservice.test_name"
	var v security.SamlAuthservice
	name := acctest.RandomNameWithPrefix("saml_authservice")
	nameUpdate := acctest.RandomNameWithPrefix("saml_authservice")
	testDataPath := getSamlTestDataPath()
	idp := map[string]any{
		"idp_type":           "AZURE_SSO",
		"metadata_file_path": filepath.Join(testDataPath, "metadata.xml"),
		"sso_redirect_url":   "2.2.2.2",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSamlAuthserviceName(name, idp),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccSamlAuthserviceName(nameUpdate, idp),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSamlAuthserviceResource_SessionTimeout(t *testing.T) {
	var resourceName = "nios_security_saml_authservice.test_session_timeout"
	var v security.SamlAuthservice
	name := acctest.RandomNameWithPrefix("saml_authservice")
	testDataPath := getSamlTestDataPath()
	idp := map[string]any{
		"idp_type":           "AZURE_SSO",
		"metadata_file_path": filepath.Join(testDataPath, "metadata.xml"),
		"sso_redirect_url":   "2.2.2.2",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSamlAuthserviceSessionTimeout(name, idp, 2700),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "session_timeout", "2700"),
				),
			},
			// Update and Read
			{
				Config: testAccSamlAuthserviceSessionTimeout(name, idp, 3600),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSamlAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "session_timeout", "3600"),
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

func testAccSamlAuthserviceBasicConfig(name string, idp map[string]any) string {
	idpString := utils.ConvertMapToHCL(idp)
	return fmt.Sprintf(`
resource "nios_security_saml_authservice" "test" {
	name = %q
	idp = %s
}
`, name, idpString)
}

func testAccSamlAuthserviceComment(name, comment string, idp map[string]any) string {
	idpString := utils.ConvertMapToHCL(idp)
	return fmt.Sprintf(`
resource "nios_security_saml_authservice" "test_comment" {
	name = %q
	idp = %s
    comment = %q
}
`, name, idpString, comment)
}

func testAccSamlAuthserviceIdp(name string, idp map[string]any) string {
	idpString := utils.ConvertMapToHCL(idp)
	return fmt.Sprintf(`
resource "nios_security_saml_authservice" "test_idp" {
	name = %q
    idp = %s
}
`, name, idpString)
}

func testAccSamlAuthserviceName(name string, idp map[string]any) string {
	idpString := utils.ConvertMapToHCL(idp)
	return fmt.Sprintf(`
resource "nios_security_saml_authservice" "test_name" {
    name = %q
	idp = %s
}
`, name, idpString)
}

func testAccSamlAuthserviceSessionTimeout(name string, idp map[string]any, sessionTimeout int) string {
	idpString := utils.ConvertMapToHCL(idp)
	return fmt.Sprintf(`
resource "nios_security_saml_authservice" "test_session_timeout" {
	name = %q
	idp = %s
    session_timeout = %d
}
`, name, idpString, sessionTimeout)
}

func getSamlTestDataPath() string {
	wd, err := os.Getwd()
	if err != nil {
		return "../../testdata/nios_security_saml_authservice"
	}
	return filepath.Join(wd, "../../testdata/nios_security_saml_authservice")
}
