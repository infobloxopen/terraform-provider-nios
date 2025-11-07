package microsoftserver_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/microsoftserver"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForMsserverAdsitesSite = "domain,name,networks"

func TestAccMsserverAdsitesSiteResource_basic(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_adsites_site.test"
	var v microsoftserver.MsserverAdsitesSite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverAdsitesSiteBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverAdsitesSiteExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverAdsitesSiteResource_disappears(t *testing.T) {
	resourceName := "nios_microsoftserver_msserver_adsites_site.test"
	var v microsoftserver.MsserverAdsitesSite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMsserverAdsitesSiteDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMsserverAdsitesSiteBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverAdsitesSiteExists(context.Background(), resourceName, &v),
					testAccCheckMsserverAdsitesSiteDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccMsserverAdsitesSiteResource_Ref(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_adsites_site.test_ref"
	var v microsoftserver.MsserverAdsitesSite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverAdsitesSiteRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverAdsitesSiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverAdsitesSiteRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverAdsitesSiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverAdsitesSiteResource_Domain(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_adsites_site.test_domain"
	var v microsoftserver.MsserverAdsitesSite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverAdsitesSiteDomain("DOMAIN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverAdsitesSiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain", "DOMAIN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverAdsitesSiteDomain("DOMAIN_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverAdsitesSiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain", "DOMAIN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverAdsitesSiteResource_Name(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_adsites_site.test_name"
	var v microsoftserver.MsserverAdsitesSite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverAdsitesSiteName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverAdsitesSiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverAdsitesSiteName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverAdsitesSiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverAdsitesSiteResource_Networks(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_adsites_site.test_networks"
	var v microsoftserver.MsserverAdsitesSite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverAdsitesSiteNetworks("NETWORKS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverAdsitesSiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "networks", "NETWORKS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverAdsitesSiteNetworks("NETWORKS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverAdsitesSiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "networks", "NETWORKS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckMsserverAdsitesSiteExists(ctx context.Context, resourceName string, v *microsoftserver.MsserverAdsitesSite) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.MicrosoftServerAPI.
			MsserverAdsitesSiteAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForMsserverAdsitesSite).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetMsserverAdsitesSiteResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetMsserverAdsitesSiteResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckMsserverAdsitesSiteDestroy(ctx context.Context, v *microsoftserver.MsserverAdsitesSite) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.MicrosoftServerAPI.
			MsserverAdsitesSiteAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForMsserverAdsitesSite).
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

func testAccCheckMsserverAdsitesSiteDisappears(ctx context.Context, v *microsoftserver.MsserverAdsitesSite) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.MicrosoftServerAPI.
			MsserverAdsitesSiteAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccMsserverAdsitesSiteBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_adsites_site" "test" {
}
`)
}

func testAccMsserverAdsitesSiteRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_adsites_site" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccMsserverAdsitesSiteDomain(domain string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_adsites_site" "test_domain" {
    domain = %q
}
`, domain)
}

func testAccMsserverAdsitesSiteName(name string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_adsites_site" "test_name" {
    name = %q
}
`, name)
}

func testAccMsserverAdsitesSiteNetworks(networks string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_adsites_site" "test_networks" {
    networks = %q
}
`, networks)
}
