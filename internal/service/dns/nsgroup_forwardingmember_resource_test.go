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
	name := acctest.RandomNameWithPrefix("ns-group-forwardingMember")
	forwardingServers := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},	
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupForwardingmemberBasicConfig(name,forwardingServers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName,"name",name),
					resource.TestCheckResourceAttr(resourceName,"forwarding_servers.0.name", "infoblox.localdomain"), 
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName,"comment",""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupForwardingmemberResource_disappears(t *testing.T) {
	resourceName := "nios_dns_nsgroup_forwardingmember.test"
	var v dns.NsgroupForwardingmember
	name := acctest.RandomNameWithPrefix("ns-group-forwardingMember")
	forwardingServers := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},	
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupForwardingmemberDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupForwardingmemberBasicConfig(name,forwardingServers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					testAccCheckNsgroupForwardingmemberDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNsgroupForwardingmemberResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardingmember.test_comment"
	var v dns.NsgroupForwardingmember
	name := acctest.RandomNameWithPrefix("ns-group-forwardingMember")
	forwardingServers := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},	
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupForwardingmemberComment(name , forwardingServers , "this is an comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "this is an comment"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupForwardingmemberComment(name , forwardingServers , "this is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "this is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupForwardingmemberResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardingmember.test_extattrs"
	var v dns.NsgroupForwardingmember
	name := acctest.RandomNameWithPrefix("ns-group-forwardingMember")
	forwardingServers := []map[string]any{
		{
			"name": "infoblox.localdomain",
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
				Config: testAccNsgroupForwardingmemberExtAttrs(name , forwardingServers ,  map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupForwardingmemberExtAttrs(name , forwardingServers ,  map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupForwardingmemberResource_ForwardingServers(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardingmember.test_forwarding_servers"
	var v dns.NsgroupForwardingmember
	name := acctest.RandomNameWithPrefix("ns-group-forwardingMember")
	forwardingServers := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},	
	}
	forwardingServersUpdate := []map[string]any{
		{
			"name": "member.com",
			"use_override_forwarders": true , 
			"forward_to":[]map[string]any{
				{
					"name":"forwarder.com",
					"address":"2.3.4.5",
				},
			},
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupForwardingmemberForwardingServers(name , forwardingServers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarding_servers.0.name", "infoblox.localdomain"),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupForwardingmemberForwardingServers(name , forwardingServersUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarding_servers.0.name", "member.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNsgroupForwardingmemberResource_Name(t *testing.T) {
	var resourceName = "nios_dns_nsgroup_forwardingmember.test_name"
	var v dns.NsgroupForwardingmember
	name := acctest.RandomNameWithPrefix("ns-group-forwardingMember")
	nameUpdate := acctest.RandomNameWithPrefix("ns-group-forwardingMember")
	forwardingServers := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},	
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNsgroupForwardingmemberName(name , forwardingServers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccNsgroupForwardingmemberName(nameUpdate,forwardingServers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsgroupForwardingmemberExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
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

func testAccNsgroupForwardingmemberBasicConfig(name string, forwardingServer []map[string]any) string {
	forwardingServerStr := utils.ConvertSliceOfMapsToHCL(forwardingServer)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardingmember" "test" {
	name = %q
	forwarding_servers = %s
}
`, name, forwardingServerStr)
}


func testAccNsgroupForwardingmemberComment(name string , forwardingServer []map[string]any, comment string) string {
	forwardingServerStr := utils.ConvertSliceOfMapsToHCL(forwardingServer)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardingmember" "test_comment" {
    name = %q
    forwarding_servers = %s
    comment = %q
}
`, name, forwardingServerStr, comment)
}

func testAccNsgroupForwardingmemberExtAttrs(name string , forwardingServer []map[string]any , extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	forwardingServerStr := utils.ConvertSliceOfMapsToHCL(forwardingServer)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardingmember" "test_extattrs" {
	name = %q
	forwarding_servers = %s
    extattrs = %s
}
`, name , forwardingServerStr , extattrsStr)
}

func testAccNsgroupForwardingmemberForwardingServers(name string , forwardingServers []map[string]any) string {
	forwardingServersStr := utils.ConvertSliceOfMapsToHCL(forwardingServers)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardingmember" "test_forwarding_servers" {
    name = %q
    forwarding_servers = %s
}
`, name, forwardingServersStr)
}

func testAccNsgroupForwardingmemberName(name string, forwardingServers []map[string]any) string {
	forwardingServersStr := utils.ConvertSliceOfMapsToHCL(forwardingServers)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup_forwardingmember" "test_name" {
    name = %q
    forwarding_servers = %s
}
`, name, forwardingServersStr)
}
