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

var readableAttributesForNsgroupForwardingmember = "comment,extattrs,forwarding_servers,name"

func TestAccNsgroupForwardingmemberResource_basic(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardingmember.test"
	var v dns.NsgroupForwardingmember

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupForwardingmemberBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupForwardingmemberResource_disappears(t *testing.T) {
	resourceName := "nios_dns_nsgroup_forwardingmember.test"
	var v dns.NsgroupForwardingmember

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupForwardingmemberDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupForwardingmemberBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					testAccCheckNsgroupForwardingmemberDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNsgroupForwardingmemberResource_Ref(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardingmember.test_ref"
	var v dns.NsgroupForwardingmember

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupForwardingmemberRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupForwardingmemberRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupForwardingmemberResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardingmember.test_comment"
	var v dns.NsgroupForwardingmember

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupForwardingmemberComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupForwardingmemberComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupForwardingmemberResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardingmember.test_extattrs"
	var v dns.NsgroupForwardingmember

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupForwardingmemberExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupForwardingmemberExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupForwardingmemberResource_ForwardingServers(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardingmember.test_forwarding_servers"
	var v dns.NsgroupForwardingmember

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupForwardingmemberForwardingServers("FORWARDING_SERVERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarding_servers", "FORWARDING_SERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupForwardingmemberForwardingServers("FORWARDING_SERVERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarding_servers", "FORWARDING_SERVERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupForwardingmemberResource_Name(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardingmember.test_name"
	var v dns.NsgroupForwardingmember

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupForwardingmemberName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupForwardingmemberName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckNsgroupForwardingmemberExists(ctx context.Context, resourceName string, v *dns.NsgroupForwardingmember) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			NsgroupForwardingmemberAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForNsgroupForwardingmember).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetNsgroupForwardingmemberResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetNsgroupForwardingmemberResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckNsgroupForwardingmemberDestroy(ctx context.Context, v *dns.NsgroupForwardingmember) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			NsgroupForwardingmemberAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForNsgroupForwardingmember).
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

func testAccCheckNsgroupForwardingmemberDisappears(ctx context.Context, v *dns.NsgroupForwardingmember) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			NsgroupForwardingmemberAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccNsgroupForwardingmemberBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardingmember" "test" {
}
`)
}

func testAccNsgroupForwardingmemberRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardingmember" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccNsgroupForwardingmemberComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardingmember" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccNsgroupForwardingmemberExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardingmember" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccNsgroupForwardingmemberForwardingServers(forwardingServers string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardingmember" "test_forwarding_servers" {
    forwarding_servers = %q
}
`, forwardingServers)
}

func testAccNsgroupForwardingmemberName(name string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardingmember" "test_name" {
    name = %q
}
`, name)
}
