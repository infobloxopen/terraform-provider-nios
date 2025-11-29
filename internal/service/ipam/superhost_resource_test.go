package ipam_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

/*
// Manage ipam Superhost with Basic Fields
resource "nios_ipam_superhost" "ipam_superhost_basic" {
    name = "NAME_REPLACE_ME"
}

// Manage ipam Superhost with Additional Fields
resource "nios_ipam_superhost" "ipam_superhost_with_additional_fields" {
    name = "NAME_REPLACE_ME"

// TODO : Add additional optional fields below

    //Extensible Attributes
    extattrs = {
        Site = "location-1"
    }
}
*/

var readableAttributesForSuperhost = "comment,dhcp_associated_objects,disabled,dns_associated_objects,extattrs,name"

func TestAccSuperhostResource_basic(t *testing.T) {
	var resourceName = "nios_ipam_superhost.test"
	var v ipam.Superhost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSuperhostBasicConfig("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "delete_associated_objects", "false"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSuperhostResource_disappears(t *testing.T) {
	resourceName := "nios_ipam_superhost.test"
	var v ipam.Superhost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSuperhostDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSuperhostBasicConfig("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					testAccCheckSuperhostDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSuperhostResource_Import(t *testing.T) {
	var resourceName = "nios_ipam_superhost.test"
	var v ipam.Superhost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSuperhostBasicConfig("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
				),
			},
			// Import with PlanOnly to detect differences
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccSuperhostImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "ref",
				PlanOnly:                             true,
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccSuperhostImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSuperhostResource_Comment(t *testing.T) {
	var resourceName = "nios_ipam_superhost.test_comment"
	var v ipam.Superhost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSuperhostComment("NAME_REPLACE_ME", "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccSuperhostComment("NAME_REPLACE_ME", "Updated comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the object"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSuperhostResource_DeleteAssociatedObjects(t *testing.T) {
	var resourceName = "nios_ipam_superhost.test_delete_associated_objects"
	var v ipam.Superhost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSuperhostDeleteAssociatedObjects("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delete_associated_objects", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSuperhostDeleteAssociatedObjects("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delete_associated_objects", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSuperhostResource_DhcpAssociatedObjects(t *testing.T) {
	var resourceName = "nios_ipam_superhost.test_dhcp_associated_objects"
	var v ipam.Superhost
	dhcpAssociatedObjectsVal := []string{"DHCP_ASSOCIATED_OBJECTS_REPLACE_ME1", "DHCP_ASSOCIATED_OBJECTS_REPLACE_ME2"}
	dhcpAssociatedObjectsValUpdated := []string{"DHCP_ASSOCIATED_OBJECTS_REPLACE_ME1", "DHCP_ASSOCIATED_OBJECTS_REPLACE_ME2"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSuperhostDhcpAssociatedObjects("NAME_REPLACE_ME", dhcpAssociatedObjectsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_associated_objects", "DHCP_ASSOCIATED_OBJECTS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSuperhostDhcpAssociatedObjects("NAME_REPLACE_ME", dhcpAssociatedObjectsValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_associated_objects", "DHCP_ASSOCIATED_OBJECTS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSuperhostResource_Disabled(t *testing.T) {
	var resourceName = "nios_ipam_superhost.test_disabled"
	var v ipam.Superhost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSuperhostDisabled("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSuperhostDisabled("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSuperhostResource_DnsAssociatedObjects(t *testing.T) {
	var resourceName = "nios_ipam_superhost.test_dns_associated_objects"
	var v ipam.Superhost
	dnsAssociatedObjectsVal := []string{"DNS_ASSOCIATED_OBJECTS_REPLACE_ME1", "DNS_ASSOCIATED_OBJECTS_REPLACE_ME2"}
	dnsAssociatedObjectsValUpdated := []string{"DNS_ASSOCIATED_OBJECTS_REPLACE_ME1", "DNS_ASSOCIATED_OBJECTS_REPLACE_ME2"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSuperhostDnsAssociatedObjects("NAME_REPLACE_ME", dnsAssociatedObjectsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_associated_objects", "DNS_ASSOCIATED_OBJECTS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSuperhostDnsAssociatedObjects("NAME_REPLACE_ME", dnsAssociatedObjectsValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_associated_objects", "DNS_ASSOCIATED_OBJECTS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSuperhostResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_ipam_superhost.test_extattrs"
	var v ipam.Superhost
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSuperhostExtAttrs("NAME_REPLACE_ME", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccSuperhostExtAttrs("NAME_REPLACE_ME", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSuperhostResource_Name(t *testing.T) {
	var resourceName = "nios_ipam_superhost.test_name"
	var v ipam.Superhost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSuperhostName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSuperhostName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckSuperhostExists(ctx context.Context, resourceName string, v *ipam.Superhost) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.IPAMAPI.
			SuperhostAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForSuperhost).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetSuperhostResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetSuperhostResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckSuperhostDestroy(ctx context.Context, v *ipam.Superhost) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.IPAMAPI.
			SuperhostAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForSuperhost).
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

func testAccCheckSuperhostDisappears(ctx context.Context, v *ipam.Superhost) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.IPAMAPI.
			SuperhostAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccSuperhostImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccSuperhostBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_ipam_superhost" "test" {
    name = %q
}
`, name)
}

func testAccSuperhostComment(name string, comment string) string {
	return fmt.Sprintf(`
resource "nios_ipam_superhost" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccSuperhostDeleteAssociatedObjects(name string, deleteAssociatedObjects string) string {
	return fmt.Sprintf(`
resource "nios_ipam_superhost" "test_delete_associated_objects" {
    name = %q
    delete_associated_objects = %q
}
`, name, deleteAssociatedObjects)
}

func testAccSuperhostDhcpAssociatedObjects(name string, dhcpAssociatedObjects []string) string {
	dhcpAssociatedObjectsStr := utils.ConvertStringSliceToHCL(dhcpAssociatedObjects)
	return fmt.Sprintf(`
resource "nios_ipam_superhost" "test_dhcp_associated_objects" {
    name = %q
    dhcp_associated_objects = %s
}
`, name, dhcpAssociatedObjectsStr)
}

func testAccSuperhostDisabled(name string, disabled string) string {
	return fmt.Sprintf(`
resource "nios_ipam_superhost" "test_disabled" {
    name = %q
    disabled = %q
}
`, name, disabled)
}

func testAccSuperhostDnsAssociatedObjects(name string, dnsAssociatedObjects []string) string {
	dnsAssociatedObjectsStr := utils.ConvertStringSliceToHCL(dnsAssociatedObjects)
	return fmt.Sprintf(`
resource "nios_ipam_superhost" "test_dns_associated_objects" {
    name = %q
    dns_associated_objects = %s
}
`, name, dnsAssociatedObjectsStr)
}

func testAccSuperhostExtAttrs(name string, extAttrs map[string]string) string {
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	return fmt.Sprintf(`
resource "nios_ipam_superhost" "test_extattrs" {
    name = %q
    extattrs = %s
}
`, name, extAttrsStr)
}

func testAccSuperhostName(name string) string {
	return fmt.Sprintf(`
resource "nios_ipam_superhost" "test_name" {
    name = %q
}
`, name)
}
