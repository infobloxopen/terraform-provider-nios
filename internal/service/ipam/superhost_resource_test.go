package ipam_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForSuperhost = "comment,dhcp_associated_objects,disabled,dns_associated_objects,extattrs,name"

func TestAccSuperhostResource_basic(t *testing.T) {
	var resourceName = "nios_ipam_superhost.test"
	var v ipam.Superhost
	name := acctest.RandomNameWithPrefix("super-host")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSuperhostBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
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
	name := acctest.RandomNameWithPrefix("super-host")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSuperhostDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSuperhostBasicConfig(name),
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
	name := acctest.RandomNameWithPrefix("super-host")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSuperhostBasicConfig(name),
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
				ImportStateVerifyIgnore:              []string{"delete_associated_objects"},
				PlanOnly:                             true,
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccSuperhostImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all", "delete_associated_objects"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSuperhostResource_Comment(t *testing.T) {
	var resourceName = "nios_ipam_superhost.test_comment"
	var v ipam.Superhost
	name := acctest.RandomNameWithPrefix("super-host")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSuperhostComment(name, "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccSuperhostComment(name, "Updated comment for the object"),
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
	name := acctest.RandomNameWithPrefix("super-host")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSuperhostDeleteAssociatedObjects(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delete_associated_objects", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSuperhostDeleteAssociatedObjects(name, "false"),
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
	name := acctest.RandomNameWithPrefix("super-host")
	dhcpAssociatedObjectsVal := []string{"${nios_dhcp_fixed_address.fixed_address.ref}"}
	dhcpAssociatedObjectsValUpdated := []string{"${nios_dhcp_fixed_address.fixed_address2.ref}"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSuperhostDhcpAssociatedObjects(name, dhcpAssociatedObjectsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "dhcp_associated_objects.0", "nios_dhcp_fixed_address.fixed_address", "ref"),
				),
			},
			// Update and Read
			{
				Config: testAccSuperhostDhcpAssociatedObjects(name, dhcpAssociatedObjectsValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "dhcp_associated_objects.0", "nios_dhcp_fixed_address.fixed_address2", "ref"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSuperhostResource_Disabled(t *testing.T) {
	var resourceName = "nios_ipam_superhost.test_disabled"
	var v ipam.Superhost
	name := acctest.RandomNameWithPrefix("super-host")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSuperhostDisabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSuperhostDisabled(name, "false"),
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
	name := acctest.RandomNameWithPrefix("super-host")
	dnsAssociatedObjectsVal := []string{"${nios_dns_record_a.record_a.ref}", "${nios_dns_record_aaaa.record_aaaa.ref}"}
	dnsAssociatedObjectsValUpdated := []string{"${nios_dns_record_ptr.record_ptr.ref}", "${nios_ip_association.association.ref}"}
	authZoneFQDN := acctest.RandomNameWithPrefix("auth-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSuperhostDnsAssociatedObjects(name, dnsAssociatedObjectsVal, authZoneFQDN),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "dns_associated_objects.0", "nios_dns_record_a.record_a", "ref"),
					resource.TestCheckResourceAttrPair(resourceName, "dns_associated_objects.1", "nios_dns_record_aaaa.record_aaaa", "ref"),
				),
			},
			// Update and Read
			{
				Config: testAccSuperhostDnsAssociatedObjects(name, dnsAssociatedObjectsValUpdated, authZoneFQDN),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "dns_associated_objects.0", "nios_dns_record_ptr.record_ptr", "ref"),
					resource.TestCheckResourceAttrPair(resourceName, "dns_associated_objects.1", "nios_ip_association.association", "ref"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSuperhostResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_ipam_superhost.test_extattrs"
	var v ipam.Superhost
	name := acctest.RandomNameWithPrefix("super-host")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSuperhostExtAttrs(name, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccSuperhostExtAttrs(name, map[string]string{
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
	name := acctest.RandomNameWithPrefix("super-host")
	name2 := acctest.RandomNameWithPrefix("super-host")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSuperhostName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccSuperhostName(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSuperhostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
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

func testAccBaseWithDNSObjects(fqdn string) string {
	ipv4addr := []map[string]any{
		{
			"ipv4addr": "192.168.1.10",
		},
	}
	ipv4addrHCL := utils.ConvertSliceOfMapsToHCL(ipv4addr)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "parent_auth_zone" {
  fqdn        = %q
  view        = "default"
}

resource "nios_dns_record_a" "record_a" {
	name = "parent-record_a.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
	ipv4addr = "10.0.0.20"
	view = "default"

	depends_on = [nios_dns_zone_auth.parent_auth_zone]
}

resource "nios_dns_record_aaaa" "record_aaaa" {
   name     = "parent-record_aaaa.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
   ipv6addr = "2002:1111::1401"
   view     = "default"
}

resource "nios_dns_record_ptr" "record_ptr" {
	name = %q
	ptrdname = "test.example.com"
	view = "default"
}

resource "nios_ip_allocation" "allocation" {
	name = "parent-record_host.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
	ipv4addrs = %s
	view = "default"
}

resource "nios_ip_association" "association" {
	ref = nios_ip_allocation.allocation.ref
	mac = %q
	configure_for_dhcp = %q
}
`, fqdn, "23.10.168.192.in-addr.arpa", ipv4addrHCL, "12:00:43:fe:9a:8c", "true")
}

func testAccBaseWithDHCPObjects() string {
	return `
resource "nios_ipam_network" "parent_network" {
  network      = "22.0.0.0/24"
  network_view = "default"
  comment      = "Parent network for DHCP fixed addresses"
}

resource "nios_dhcp_fixed_address" "fixed_address" {
	ipv4addr = "22.0.0.20"
	match_client = "CIRCUIT_ID"
	agent_circuit_id = 23

	depends_on = [nios_ipam_network.parent_network]
}

resource "nios_dhcp_fixed_address" "fixed_address2" {
	ipv4addr = "22.0.0.21"
	match_client = "CIRCUIT_ID"
	agent_circuit_id = 24

	depends_on = [nios_ipam_network.parent_network]
}
`
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
	config := fmt.Sprintf(`
resource "nios_ipam_superhost" "test_dhcp_associated_objects" {
    name = %q
    dhcp_associated_objects = %s
}
`, name, dhcpAssociatedObjectsStr)
	return strings.Join([]string{testAccBaseWithDHCPObjects(), config}, "")
}

func testAccSuperhostDisabled(name string, disabled string) string {
	return fmt.Sprintf(`
resource "nios_ipam_superhost" "test_disabled" {
    name = %q
    disabled = %q
}
`, name, disabled)
}

func testAccSuperhostDnsAssociatedObjects(name string, dnsAssociatedObjects []string, authZoneFQDN string) string {
	dnsAssociatedObjectsStr := utils.ConvertStringSliceToHCL(dnsAssociatedObjects)
	config := fmt.Sprintf(`
resource "nios_ipam_superhost" "test_dns_associated_objects" {
    name = %q
    dns_associated_objects = %s
}
`, name, dnsAssociatedObjectsStr)
	return strings.Join([]string{testAccBaseWithDNSObjects(authZoneFQDN), config}, "")

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
