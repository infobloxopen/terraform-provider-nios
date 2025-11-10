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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRirOrganizationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRirOrganizationBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					testAccCheckRirOrganizationDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRirOrganizationResource_Ref(t *testing.T) {
	var resourceName = "nios_rir_organization.test_ref"
	var v rir.RirOrganization

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRirOrganizationRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRirOrganizationResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_rir_organization.test_extattrs"
	var v rir.RirOrganization

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRirOrganizationExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRirOrganizationResource_Id(t *testing.T) {
	var resourceName = "nios_rir_organization.test_id"
	var v rir.RirOrganization

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationId("ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "id", "ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRirOrganizationId("ID_UPDATE_REPLACE_ME"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationMaintainer("MAINTAINER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "maintainer", "MAINTAINER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRirOrganizationMaintainer("MAINTAINER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "maintainer", "MAINTAINER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRirOrganizationResource_Name(t *testing.T) {
	var resourceName = "nios_rir_organization.test_name"
	var v rir.RirOrganization

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRirOrganizationName("NAME_UPDATE_REPLACE_ME"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationPassword("PASSWORD_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "password", "PASSWORD_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRirOrganizationPassword("PASSWORD_UPDATE_REPLACE_ME"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationRir("RIR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir", "RIR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRirOrganizationRir("RIR_UPDATE_REPLACE_ME"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationSenderEmail("SENDER_EMAIL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sender_email", "SENDER_EMAIL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRirOrganizationSenderEmail("SENDER_EMAIL_UPDATE_REPLACE_ME"),
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

func testAccRirOrganizationBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_rir_organization" "test" {
}
`)
}

func testAccRirOrganizationRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_rir_organization" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccRirOrganizationExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_rir_organization" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccRirOrganizationId(id string) string {
	return fmt.Sprintf(`
resource "nios_rir_organization" "test_id" {
    id = %q
}
`, id)
}

func testAccRirOrganizationMaintainer(maintainer string) string {
	return fmt.Sprintf(`
resource "nios_rir_organization" "test_maintainer" {
    maintainer = %q
}
`, maintainer)
}

func testAccRirOrganizationName(name string) string {
	return fmt.Sprintf(`
resource "nios_rir_organization" "test_name" {
    name = %q
}
`, name)
}

func testAccRirOrganizationPassword(password string) string {
	return fmt.Sprintf(`
resource "nios_rir_organization" "test_password" {
    password = %q
}
`, password)
}

func testAccRirOrganizationRir(rir string) string {
	return fmt.Sprintf(`
resource "nios_rir_organization" "test_rir" {
    rir = %q
}
`, rir)
}

func testAccRirOrganizationSenderEmail(senderEmail string) string {
	return fmt.Sprintf(`
resource "nios_rir_organization" "test_sender_email" {
    sender_email = %q
}
`, senderEmail)
}
