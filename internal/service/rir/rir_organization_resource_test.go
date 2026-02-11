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
	id := fmt.Sprintf("ORG-CB%d-IBTEST", acctest.RandomNumber(9999))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationBasicConfig(map[string]string{
					"RIPE Admin Contact":     "ib-contact",
					"RIPE Country":           "United Kingdom (GB)",
					"RIPE Technical Contact": "TEST123-IB",
					"RIPE Email":             "support@infoblox.com",
				},
					id, "infoblox", name, "test-pass", "RIPE", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.RIPE Admin Contact", "ib-contact"),
					resource.TestCheckResourceAttr(resourceName, "extattrs.RIPE Country", "United Kingdom (GB)"),
					resource.TestCheckResourceAttr(resourceName, "extattrs.RIPE Technical Contact", "TEST123-IB"),
					resource.TestCheckResourceAttr(resourceName, "extattrs.RIPE Email", "support@infoblox.com"),
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "maintainer", "infoblox"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "password", "test-pass"),
					resource.TestCheckResourceAttr(resourceName, "sender_email", "support@infoblox.com"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "rir", "RIPE"),
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
	id := fmt.Sprintf("ORG-CB%d-IBTEST", acctest.RandomNumber(9999))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRirOrganizationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRirOrganizationBasicConfig(map[string]string{
					"RIPE Admin Contact":     "ib-contact",
					"RIPE Country":           "United Kingdom (GB)",
					"RIPE Technical Contact": "TEST123-IB",
					"RIPE Email":             "support@infoblox.com",
				},
					id, "infoblox", name, "test-pass", "RIPE", "support@infoblox.com"),
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
	id := fmt.Sprintf("ORG-CB%d-IBTEST", acctest.RandomNumber(9999))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationBasicConfig(map[string]string{
					"RIPE Admin Contact":     "ib-contact",
					"RIPE Country":           "United Kingdom (GB)",
					"RIPE Technical Contact": "TEST123-IB",
					"RIPE Email":             "support@infoblox.com",
				},
					id, "infoblox", name, "test-pass", "RIPE", "support@infoblox.com"),
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
				ImportStateVerifyIgnore:              []string{"password"},
				PlanOnly:                             true,
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccRirOrganizationImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"password"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRirOrganizationResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_rir_organization.test_extattrs"
	var v rir.RirOrganization
	name := acctest.RandomNameWithPrefix("rir-org")
	id := fmt.Sprintf("ORG-CB%d-IBTEST", acctest.RandomNumber(9999))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationExtAttrs(map[string]string{
					"RIPE Admin Contact":     "ib-contact",
					"RIPE Country":           "United Kingdom (GB)",
					"RIPE Technical Contact": "TEST123-IB",
					"RIPE Email":             "support@infoblox.com",
				},
					id, "infoblox", name, "test-pass", "RIPE", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "maintainer", "infoblox"),
				),
			},
			// Update and Read
			{
				Config: testAccRirOrganizationExtAttrs(map[string]string{
					"RIPE Admin Contact":     "ib-contact",
					"RIPE Country":           "United Kingdom (GB)",
					"RIPE Technical Contact": "TEST123-IB",
					"RIPE Email":             "support@infoblox.com",
				},
					id, "nios-support", name, "test-pass", "RIPE", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "maintainer", "nios-support"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRirOrganizationResource_Id(t *testing.T) {
	var resourceName = "nios_rir_organization.test_id"
	var v rir.RirOrganization
	name := acctest.RandomNameWithPrefix("rir-org")
	id := fmt.Sprintf("ORG-CB%d-IBTEST", acctest.RandomNumber(9999))
	id2 := fmt.Sprintf("ORG-CB%d-IBTEST", acctest.RandomNumber(9999))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationId(map[string]string{
					"RIPE Admin Contact":     "ib-contact",
					"RIPE Country":           "United Kingdom (GB)",
					"RIPE Technical Contact": "TEST123-IB",
					"RIPE Email":             "support@infoblox.com",
				},
					id, "infoblox", name, "test-pass", "RIPE", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "id", id),
				),
			},
			// Update and Read
			{
				Config: testAccRirOrganizationId(map[string]string{
					"RIPE Admin Contact":     "ib-contact",
					"RIPE Country":           "United Kingdom (GB)",
					"RIPE Technical Contact": "TEST123-IB",
					"RIPE Email":             "support@infoblox.com",
				},
					id2, "infoblox", name, "test-pass", "RIPE", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "id", id2),
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
	id := fmt.Sprintf("ORG-CB%d-IBTEST", acctest.RandomNumber(9999))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationMaintainer(map[string]string{
					"RIPE Admin Contact":     "ib-contact",
					"RIPE Country":           "United Kingdom (GB)",
					"RIPE Technical Contact": "TEST123-IB",
					"RIPE Email":             "support@infoblox.com",
				},
					id, "infoblox", name, "test-pass", "RIPE", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "maintainer", "infoblox"),
				),
			},
			// Update and Read
			{
				Config: testAccRirOrganizationMaintainer(map[string]string{
					"RIPE Admin Contact":     "ib-contact",
					"RIPE Country":           "United Kingdom (GB)",
					"RIPE Technical Contact": "TEST123-IB",
					"RIPE Email":             "support@infoblox.com",
				},
					id, "nios-support", name, "test-pass", "RIPE", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "maintainer", "nios-support"),
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
	name2 := acctest.RandomNameWithPrefix("rir-org")
	id := fmt.Sprintf("ORG-CB%d-IBTEST", acctest.RandomNumber(9999))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationName(map[string]string{
					"RIPE Admin Contact":     "ib-contact",
					"RIPE Country":           "United Kingdom (GB)",
					"RIPE Technical Contact": "TEST123-IB",
					"RIPE Email":             "support@infoblox.com",
				},
					id, "infoblox", name, "test-pass", "RIPE", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccRirOrganizationName(map[string]string{
					"RIPE Admin Contact":     "ib-contact",
					"RIPE Country":           "United Kingdom (GB)",
					"RIPE Technical Contact": "TEST123-IB",
					"RIPE Email":             "support@infoblox.com",
				},
					id, "infoblox", name2, "test-pass", "RIPE", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
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
	id := fmt.Sprintf("ORG-CB%d-IBTEST", acctest.RandomNumber(9999))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationPassword(map[string]string{
					"RIPE Admin Contact":     "ib-contact",
					"RIPE Country":           "United Kingdom (GB)",
					"RIPE Technical Contact": "TEST123-IB",
					"RIPE Email":             "support@infoblox.com",
				},
					id, "infoblox", name, "test-pass", "RIPE", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "password", "test-pass"),
				),
			},
			// Update and Read
			{
				Config: testAccRirOrganizationPassword(map[string]string{
					"RIPE Admin Contact":     "ib-contact",
					"RIPE Country":           "United Kingdom (GB)",
					"RIPE Technical Contact": "TEST123-IB",
					"RIPE Email":             "support@infoblox.com",
				},
					id, "infoblox", name, "test-pass2", "RIPE", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "password", "test-pass2"),
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
	id := fmt.Sprintf("ORG-CB%d-IBTEST", acctest.RandomNumber(9999))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationRir(map[string]string{
					"RIPE Admin Contact":     "ib-contact",
					"RIPE Country":           "United Kingdom (GB)",
					"RIPE Technical Contact": "TEST123-IB",
					"RIPE Email":             "support@infoblox.com",
				},
					id, "infoblox", name, "test-pass", "RIPE", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir", "RIPE"),
				),
			},
			// Update unavailable as RIPE is the only supported RIR type
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRirOrganizationResource_SenderEmail(t *testing.T) {
	var resourceName = "nios_rir_organization.test_sender_email"
	var v rir.RirOrganization
	name := acctest.RandomNameWithPrefix("rir-org")
	id := fmt.Sprintf("ORG-CB%d-IBTEST", acctest.RandomNumber(9999))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRirOrganizationSenderEmail(map[string]string{
					"RIPE Admin Contact":     "ib-contact",
					"RIPE Country":           "United Kingdom (GB)",
					"RIPE Technical Contact": "TEST123-IB",
					"RIPE Email":             "support@infoblox.com",
				},
					id, "infoblox", name, "test-pass", "RIPE", "support@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sender_email", "support@infoblox.com"),
				),
			},
			// Update and Read
			{
				Config: testAccRirOrganizationSenderEmail(map[string]string{
					"RIPE Admin Contact":     "ib-contact",
					"RIPE Country":           "United Kingdom (GB)",
					"RIPE Technical Contact": "TEST123-IB",
					"RIPE Email":             "support@infoblox.com",
				},
					id, "infoblox", name, "test-pass", "RIPE", "support2@infoblox.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRirOrganizationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sender_email", "support2@infoblox.com"),
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

func formatExtAttrs(extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %q = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return extattrsStr
}

func testAccRirOrganizationBasicConfig(extAttrs map[string]string, id, maintainer, name, password, rir, senderEmail string) string {
	extattrsStr := formatExtAttrs(extAttrs)
	return fmt.Sprintf(`
resource "nios_rir_organization" "test" {
    extattrs = %s
    id = %q
    maintainer = %q
    name = %q
    password = %q
    rir = %q
    sender_email = %q
}
`, extattrsStr, id, maintainer, name, password, rir, senderEmail)
}

func testAccRirOrganizationExtAttrs(extAttrs map[string]string, id string, maintainer string, name string, password string, rir string, senderEmail string) string {
	extattrsStr := formatExtAttrs(extAttrs)
	return fmt.Sprintf(`
resource "nios_rir_organization" "test_extattrs" {
    extattrs = %s
    id = %q
    maintainer = %q
    name = %q
    password = %q
    rir = %q
    sender_email = %q
}
`, extattrsStr, id, maintainer, name, password, rir, senderEmail)
}

func testAccRirOrganizationId(extAttrs map[string]string, id string, maintainer string, name string, password string, rir string, senderEmail string) string {
	extattrsStr := formatExtAttrs(extAttrs)
	return fmt.Sprintf(`
resource "nios_rir_organization" "test_id" {
    extattrs = %s
    id = %q
    maintainer = %q
    name = %q
    password = %q
    rir = %q
    sender_email = %q
}
`, extattrsStr, id, maintainer, name, password, rir, senderEmail)
}

func testAccRirOrganizationMaintainer(extAttrs map[string]string, id string, maintainer string, name string, password string, rir string, senderEmail string) string {
	extattrsStr := formatExtAttrs(extAttrs)
	return fmt.Sprintf(`
resource "nios_rir_organization" "test_maintainer" {
    extattrs = %s
    id = %q
    maintainer = %q
    name = %q
    password = %q
    rir = %q
    sender_email = %q
}
`, extattrsStr, id, maintainer, name, password, rir, senderEmail)
}

func testAccRirOrganizationName(extAttrs map[string]string, id string, maintainer string, name string, password string, rir string, senderEmail string) string {
	extattrsStr := formatExtAttrs(extAttrs)
	return fmt.Sprintf(`
resource "nios_rir_organization" "test_name" {
    extattrs = %s
    id = %q
    maintainer = %q
    name = %q
    password = %q
    rir = %q
    sender_email = %q
}
`, extattrsStr, id, maintainer, name, password, rir, senderEmail)
}

func testAccRirOrganizationPassword(extAttrs map[string]string, id string, maintainer string, name string, password string, rir string, senderEmail string) string {
	extattrsStr := formatExtAttrs(extAttrs)
	return fmt.Sprintf(`
resource "nios_rir_organization" "test_password" {
    extattrs = %s
    id = %q
    maintainer = %q
    name = %q
    password = %q
    rir = %q
    sender_email = %q
}
`, extattrsStr, id, maintainer, name, password, rir, senderEmail)
}

func testAccRirOrganizationRir(extAttrs map[string]string, id string, maintainer string, name string, password string, rir string, senderEmail string) string {
	extattrsStr := formatExtAttrs(extAttrs)
	return fmt.Sprintf(`
resource "nios_rir_organization" "test_rir" {
    extattrs = %s
    id = %q
    maintainer = %q
    name = %q
    password = %q
    rir = %q
    sender_email = %q
}
`, extattrsStr, id, maintainer, name, password, rir, senderEmail)
}

func testAccRirOrganizationSenderEmail(extAttrs map[string]string, id string, maintainer string, name string, password string, rir string, senderEmail string) string {
	extattrsStr := formatExtAttrs(extAttrs)
	return fmt.Sprintf(`
resource "nios_rir_organization" "test_sender_email" {
    extattrs = %s
    id = %q
    maintainer = %q
    name = %q
    password = %q
    rir = %q
    sender_email = %q
}
`, extattrsStr, id, maintainer, name, password, rir, senderEmail)
}
