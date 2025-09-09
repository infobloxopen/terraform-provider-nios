package dns_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForNsgroupForwardstubserver = "comment,extattrs,external_servers,name"

func TestAccNsgroupForwardstubserverResource_basic(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardstubserver.test"
	var v dns.NsgroupForwardstubserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupForwardstubserverBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupForwardstubserverResource_disappears(t *testing.T) {
	resourceName := "nios_dns_nsgroup_forwardstubserver.test"
	var v dns.NsgroupForwardstubserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupForwardstubserverDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupForwardstubserverBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					testAccCheckNsgroupForwardstubserverDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNsgroupForwardstubserverResource_Ref(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardstubserver.test_ref"
	var v dns.NsgroupForwardstubserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupForwardstubserverRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupForwardstubserverRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupForwardstubserverResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardstubserver.test_comment"
	var v dns.NsgroupForwardstubserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupForwardstubserverComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupForwardstubserverComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupForwardstubserverResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardstubserver.test_extattrs"
	var v dns.NsgroupForwardstubserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupForwardstubserverExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupForwardstubserverExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupForwardstubserverResource_ExternalServers(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardstubserver.test_external_servers"
	var v dns.NsgroupForwardstubserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupForwardstubserverExternalServers("EXTERNAL_SERVERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_servers", "EXTERNAL_SERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupForwardstubserverExternalServers("EXTERNAL_SERVERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_servers", "EXTERNAL_SERVERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupForwardstubserverResource_Name(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardstubserver.test_name"
	var v dns.NsgroupForwardstubserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupForwardstubserverName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupForwardstubserverName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckNsgroupForwardstubserverExists(ctx context.Context, resourceName string, v *dns.NsgroupForwardstubserver) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			NsgroupForwardstubserverAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForNsgroupForwardstubserver).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetNsgroupForwardstubserverResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetNsgroupForwardstubserverResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckNsgroupForwardstubserverDestroy(ctx context.Context, v *dns.NsgroupForwardstubserver) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			NsgroupForwardstubserverAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForNsgroupForwardstubserver).
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

func testAccCheckNsgroupForwardstubserverDisappears(ctx context.Context, v *dns.NsgroupForwardstubserver) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			NsgroupForwardstubserverAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccNsgroupForwardstubserverBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardstubserver" "test" {
}
`)
}

func testAccNsgroupForwardstubserverRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardstubserver" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccNsgroupForwardstubserverComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardstubserver" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccNsgroupForwardstubserverExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardstubserver" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccNsgroupForwardstubserverExternalServers(externalServers string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardstubserver" "test_external_servers" {
    external_servers = %q
}
`, externalServers)
}

func testAccNsgroupForwardstubserverName(name string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardstubserver" "test_name" {
    name = %q
}
`, name)
}
