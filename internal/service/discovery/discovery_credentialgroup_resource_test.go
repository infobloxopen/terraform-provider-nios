package discovery_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/discovery"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForDiscoveryCredentialgroup = "name"

func TestAccDiscoveryCredentialgroupResource_basic(t *testing.T) {
	var resourceName = "nios_discovery_credentialgroup.test"
	var v discovery.DiscoveryCredentialgroup
	name := acctest.RandomNameWithPrefix("example-discovery-credentialgroup")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDiscoveryCredentialgroupBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDiscoveryCredentialgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDiscoveryCredentialgroupResource_disappears(t *testing.T) {
	resourceName := "nios_discovery_credentialgroup.test"
	var v discovery.DiscoveryCredentialgroup
	name := acctest.RandomNameWithPrefix("example-discovery-credentialgroup")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDiscoveryCredentialgroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDiscoveryCredentialgroupBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDiscoveryCredentialgroupExists(context.Background(), resourceName, &v),
					testAccCheckDiscoveryCredentialgroupDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDiscoveryCredentialgroupResource_Name(t *testing.T) {
	var resourceName = "nios_discovery_credentialgroup.test_name"
	name := acctest.RandomNameWithPrefix("example-discovery-credentialgroup")
	updatedName := acctest.RandomNameWithPrefix("example-discovery-credentialgroup")
	var v discovery.DiscoveryCredentialgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDiscoveryCredentialgroupName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDiscoveryCredentialgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDiscoveryCredentialgroupName(updatedName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDiscoveryCredentialgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDiscoveryCredentialgroupExists(ctx context.Context, resourceName string, v *discovery.DiscoveryCredentialgroup) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DiscoveryAPI.
			DiscoveryCredentialgroupAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDiscoveryCredentialgroup).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDiscoveryCredentialgroupResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDiscoveryCredentialgroupResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDiscoveryCredentialgroupDestroy(ctx context.Context, v *discovery.DiscoveryCredentialgroup) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DiscoveryAPI.
			DiscoveryCredentialgroupAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDiscoveryCredentialgroup).
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

func testAccCheckDiscoveryCredentialgroupDisappears(ctx context.Context, v *discovery.DiscoveryCredentialgroup) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DiscoveryAPI.
			DiscoveryCredentialgroupAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDiscoveryCredentialgroupBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_discovery_credentialgroup" "test" {
	name = %q
}
`, name)
}

func testAccDiscoveryCredentialgroupName(name string) string {
	return fmt.Sprintf(`
resource "nios_discovery_credentialgroup" "test_name" {
    name = %q
}
`, name)
}
