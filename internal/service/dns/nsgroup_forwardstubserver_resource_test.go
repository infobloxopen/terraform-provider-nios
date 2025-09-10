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
	name := acctest.RandomNameWithPrefix("ns-group-forwardstubserver")
	externalServers := []map[string]any{
		{
			"name":    "example.com",
			"address": "2.3.3.4",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupForwardstubserverBasicConfig(name, externalServers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "external_servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "external_servers.0.name", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "external_servers.0.address", "2.3.3.4"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupForwardstubserverResource_disappears(t *testing.T) {
	resourceName := "nios_dns_nsgroup_forwardstubserver.test"
	var v dns.NsgroupForwardstubserver
	name := acctest.RandomNameWithPrefix("ns-group-forwardstubserver")
	externalServers := []map[string]any{
		{
			"name":    "example.com",
			"address": "2.3.3.4",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupForwardstubserverDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupForwardstubserverBasicConfig(name, externalServers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					testAccCheckNsgroupForwardstubserverDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNsgroupForwardstubserverResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardstubserver.test_comment"
	var v dns.NsgroupForwardstubserver
	name := acctest.RandomNameWithPrefix("ns-group-forwardstubserver")
	externalServers := []map[string]any{
		{
			"name":    "example.com",
			"address": "2.3.3.4",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupForwardstubserverComment(name, externalServers, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupForwardstubserverComment(name, externalServers, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupForwardstubserverResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardstubserver.test_extattrs"
	var v dns.NsgroupForwardstubserver
	name := acctest.RandomNameWithPrefix("ns-group-forwardstubserver")
	externalServers := []map[string]any{
		{
			"name":    "example.com",
			"address": "2.3.3.4",
		},
	}
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupForwardstubserverExtAttrs(name, externalServers, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupForwardstubserverExtAttrs(name, externalServers, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupForwardstubserverResource_ExternalServers(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardstubserver.test_external_servers"
	var v dns.NsgroupForwardstubserver
	name := acctest.RandomNameWithPrefix("ns-group-forwardstubserver")
	externalServers := []map[string]any{
		{
			"name":    "example.com",
			"address": "2.3.3.4",
		},
	}
	externalServersUpdate := []map[string]any{
		{
			"name":    "example1.com",
			"address": "2.3.4.4",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupForwardstubserverExternalServers(name, externalServers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_servers.0.name", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "external_servers.0.address", "2.3.3.4"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupForwardstubserverExternalServers(name, externalServersUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_servers.0.name", "example1.com"),
					resource.TestCheckResourceAttr(resourceName, "external_servers.0.address", "2.3.4.4"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupForwardstubserverResource_Name(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardstubserver.test_name"
	var v dns.NsgroupForwardstubserver
	name := acctest.RandomNameWithPrefix("ns-group-forwardstubserver")
	nameUpdate := acctest.RandomNameWithPrefix("ns-group-forwardstubserver")
	externalServers := []map[string]any{
		{
			"name":    "example.com",
			"address": "2.3.3.4",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupForwardstubserverName(name, externalServers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupForwardstubserverName(nameUpdate, externalServers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardstubserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
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

func testAccNsgroupForwardstubserverBasicConfig(name string, externalServers []map[string]any) string {
	externalServersStr := utils.ConvertSliceOfMapsToHCL(externalServers)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardstubserver" "test" {
    name = %q
    external_servers = %s
}
`, name, externalServersStr)
}

func testAccNsgroupForwardstubserverComment(name string, externalServer []map[string]any, comment string) string {
	externalServersStr := utils.ConvertSliceOfMapsToHCL(externalServer)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardstubserver" "test_comment" {
    name = %q
    external_servers = %s
    comment = %q
}
`, name, externalServersStr, comment)
}

func testAccNsgroupForwardstubserverExtAttrs(name string, externalServer []map[string]any, extAttrs map[string]string) string {
	externalServersStr := utils.ConvertSliceOfMapsToHCL(externalServer)
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardstubserver" "test_extattrs" {
    name = %q
    external_servers = %s
    extattrs = %s
}
`, name, externalServersStr, extattrsStr)
}

func testAccNsgroupForwardstubserverExternalServers(name string, externalServers []map[string]any) string {
	externalServersStr := utils.ConvertSliceOfMapsToHCL(externalServers)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardstubserver" "test_external_servers" {
    name = %q
    external_servers = %s
}
`, name, externalServersStr)
}

func testAccNsgroupForwardstubserverName(name string, externalServers []map[string]any) string {
	externalServersStr := utils.ConvertSliceOfMapsToHCL(externalServers)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardstubserver" "test_name" {
    name = %q
    external_servers = %s
}
`, name, externalServersStr)
}
