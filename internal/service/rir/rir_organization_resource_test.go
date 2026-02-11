package rir_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/rir"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForRirOrganization = "extattrs,id,maintainer,name,rir,sender_email"

func TestAccRirOrganizationResource_basic(t *testing.T) {
	var resourceName = "nios_rir_organization.test"
	var v rir.RirOrganization
	name := acctest.RandomNameWithPrefix("rir-org")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationBasicConfig("ID_REPLACE_ME", "infoblox", name, "PASSWORD_REPLACE_ME", "RIR", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "id", "ID_REPLACE_ME"),
					resource.TestCheckResourceAttr(resourceName, "maintainer", "infoblox"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "password", "PASSWORD_REPLACE_ME"),
					resource.TestCheckResourceAttr(resourceName, "rir", "RIR"),
					resource.TestCheckResourceAttr(resourceName, "sender_email", "support@infoblox.com"),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRirOrganizationResource_disappears(t *testing.T) {
	resourceName := "nios_rir_organization.test"
	var v rir.RirOrganization
	name := acctest.RandomNameWithPrefix("rir-org")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRirOrganizationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRirOrganizationBasicConfig("ID_REPLACE_ME", "infoblox", name, "PASSWORD_REPLACE_ME", "RIR", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					testAccCheckRirOrganizationDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRirOrganizationResource_Import(t *testing.T) {
	var resourceName = "nios_rir_organization.test"
	var v rir.RirOrganization
	name := acctest.RandomNameWithPrefix("rir-org")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationBasicConfig("ID_REPLACE_ME", "infoblox", name, "PASSWORD_REPLACE_ME", "RIR", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
				),
			},
			// Import with PlanOnly to detect differences
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccRirOrganizationImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "ref",
				PlanOnly:                             true,
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccRirOrganizationImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRirOrganizationResource_Id(t *testing.T) {
	var resourceName = "nios_rir_organization.test_id"
	var v rir.RirOrganization
	name := acctest.RandomNameWithPrefix("rir-org")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationId("ID_REPLACE_ME", "infoblox", name, "PASSWORD_REPLACE_ME", "RIR", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "id", "ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRirOrganizationId("ID_REPLACE_ME", "infoblox", name, "PASSWORD_REPLACE_ME", "RIR", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "id", "ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRirOrganizationResource_Maintainer(t *testing.T) {
	var resourceName = "nios_rir_organization.test_maintainer"
	var v rir.RirOrganization
	name := acctest.RandomNameWithPrefix("rir-org")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationMaintainer("ID_REPLACE_ME", "infoblox", name, "PASSWORD_REPLACE_ME", "RIR", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "maintainer", "infoblox"),
				),
			},
			// Update and Read
			{
				Config: testAccRirOrganizationMaintainer("ID_REPLACE_ME", "NIOS", name, "PASSWORD_REPLACE_ME", "RIR", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "maintainer", "NIOS"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRirOrganizationResource_Name(t *testing.T) {
	var resourceName = "nios_rir_organization.test_name"
	var v rir.RirOrganization
	name := acctest.RandomNameWithPrefix("rir-org")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationName("ID_REPLACE_ME", "infoblox", name, "PASSWORD_REPLACE_ME", "RIR", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccRirOrganizationName("ID_REPLACE_ME", "infoblox", name, "PASSWORD_REPLACE_ME", "RIR", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRirOrganizationResource_Password(t *testing.T) {
	var resourceName = "nios_rir_organization.test_password"
	var v rir.RirOrganization
	name := acctest.RandomNameWithPrefix("rir-org")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationPassword("ID_REPLACE_ME", "infoblox", name, "PASSWORD_REPLACE_ME", "RIR", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "password", "PASSWORD_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRirOrganizationPassword("ID_REPLACE_ME", "infoblox", name, "PASSWORD_REPLACE_ME", "RIR", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "password", "PASSWORD_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRirOrganizationResource_Rir(t *testing.T) {
	var resourceName = "nios_rir_organization.test_rir"
	var v rir.RirOrganization
	name := acctest.RandomNameWithPrefix("rir-org")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationRir("ID_REPLACE_ME", "infoblox", name, "PASSWORD_REPLACE_ME", "RIR", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir", "RIR"),
				),
			},
			// Update and Read
			{
				Config: testAccRirOrganizationRir("ID_REPLACE_ME", "infoblox", name, "PASSWORD_REPLACE_ME", "RIR", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir", "RIR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRirOrganizationResource_SenderEmail(t *testing.T) {
	var resourceName = "nios_rir_organization.test_sender_email"
	var v rir.RirOrganization
	name := acctest.RandomNameWithPrefix("rir-org")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationSenderEmail("ID_REPLACE_ME", "infoblox", name, "PASSWORD_REPLACE_ME", "RIR", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sender_email", "support@infoblox.com"),
				),
			},
			// Update and Read
			{
				Config: testAccRirOrganizationSenderEmail("ID_REPLACE_ME", "infoblox", name, "PASSWORD_REPLACE_ME", "RIR", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sender_email", "SENDER_EMAIL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRirOrganizationExists(ctx context.Context, resourceName string, v *rir.RirOrganization) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.RIRAPI.
			RirOrganizationAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRirOrganization).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRirOrganizationResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRirOrganizationResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRirOrganizationDestroy(ctx context.Context, v *rir.RirOrganization) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.RIRAPI.
			RirOrganizationAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRirOrganization).
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

func testAccCheckRirOrganizationDisappears(ctx context.Context, v *rir.RirOrganization) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.RIRAPI.
			RirOrganizationAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRirOrganizationImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.Attributes["ref"] == "" {
			return "", fmt.Errorf("ref is not set")
		}
		return rs.Primary.Attributes["ref"], nil
	}
}

func testAccRirOrganizationBasicConfig(id, maintainer, name, password, rir, senderEmail string) string {
	return fmt.Sprintf(`
resource "nios_rir_organization" "test" {
    id = %q
    maintainer = %q
    name = %q
    password = %q
    rir = %q
    sender_email = %q
}
`, id, maintainer, name, password, rir, senderEmail)
}

func testAccRirOrganizationId(id string, maintainer string, name string, password string, rir string, senderEmail string) string {
	return fmt.Sprintf(`
resource "nios_rir_organization" "test_id" {
    id = %q
    maintainer = %q
    name = %q
    password = %q
    rir = %q
    sender_email = %q
}
`, id, maintainer, name, password, rir, senderEmail)
}

func testAccRirOrganizationMaintainer(id string, maintainer string, name string, password string, rir string, senderEmail string) string {
	return fmt.Sprintf(`
resource "nios_rir_organization" "test_maintainer" {
    id = %q
    maintainer = %q
    name = %q
    password = %q
    rir = %q
    sender_email = %q
}
`, id, maintainer, name, password, rir, senderEmail)
}

func testAccRirOrganizationName(id string, maintainer string, name string, password string, rir string, senderEmail string) string {
	return fmt.Sprintf(`
resource "nios_rir_organization" "test_name" {
    id = %q
    maintainer = %q
    name = %q
    password = %q
    rir = %q
    sender_email = %q
}
`, id, maintainer, name, password, rir, senderEmail)
}

func testAccRirOrganizationPassword(id string, maintainer string, name string, password string, rir string, senderEmail string) string {
	return fmt.Sprintf(`
resource "nios_rir_organization" "test_password" {
    id = %q
    maintainer = %q
    name = %q
    password = %q
    rir = %q
    sender_email = %q
}
`, id, maintainer, name, password, rir, senderEmail)
}

func testAccRirOrganizationRir(id string, maintainer string, name string, password string, rir string, senderEmail string) string {
	return fmt.Sprintf(`
resource "nios_rir_organization" "test_rir" {
    id = %q
    maintainer = %q
    name = %q
    password = %q
    rir = %q
    sender_email = %q
}
`, id, maintainer, name, password, rir, senderEmail)
}

func testAccRirOrganizationSenderEmail(id string, maintainer string, name string, password string, rir string, senderEmail string) string {
	return fmt.Sprintf(`
resource "nios_rir_organization" "test_sender_email" {
    id = %q
    maintainer = %q
    name = %q
    password = %q
    rir = %q
    sender_email = %q
}
`, id, maintainer, name, password, rir, senderEmail)
}
