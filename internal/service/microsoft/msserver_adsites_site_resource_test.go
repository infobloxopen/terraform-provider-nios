package microsoft_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/microsoft"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForMsserverAdsitesSite = "domain,name,networks"

func TestAccMsserverAdsitesSiteResource_basic(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_microsoft_msserver_adsites_site.test"
	var v microsoft.MsserverAdsitesSite

	name := acctest.RandomName()
	domain := "msserver:adsites:domain/ZG5zLm1zX2FkX3NpdGVzX2RvbWFpbiQwLkFELTE3MA:example.local/default"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverAdsitesSiteBasicConfig(domain, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverAdsitesSiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain", domain),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverAdsitesSiteResource_disappears(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	resourceName := "nios_microsoft_msserver_adsites_site.test"
	var v microsoft.MsserverAdsitesSite

	name := acctest.RandomName()
	domain := "msserver:adsites:domain/ZG5zLm1zX2FkX3NpdGVzX2RvbWFpbiQwLkFELTE3MA:example.local/default"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMsserverAdsitesSiteDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMsserverAdsitesSiteBasicConfig(domain, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverAdsitesSiteExists(context.Background(), resourceName, &v),
					testAccCheckMsserverAdsitesSiteDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccMsserverAdsitesSiteResource_Domain(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_microsoft_msserver_adsites_site.test_domain"
	var v microsoft.MsserverAdsitesSite

	name := acctest.RandomName()
	domain1 := "msserver:adsites:domain/ZG5zLm1zX2FkX3NpdGVzX2RvbWFpbiQwLkFELTE3MA:example1.local/default"
	domain2 := "msserver:adsites:domain/ZG5zLm1zX2FkX3NpdGVzX2RvbWFpbiQwLkFELTE4MB:example2.local/default"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverAdsitesSiteDomain(domain1, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverAdsitesSiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain", domain1),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverAdsitesSiteDomain(domain2, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverAdsitesSiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain", domain2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverAdsitesSiteResource_Name(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_microsoft_msserver_adsites_site.test_name"
	var v microsoft.MsserverAdsitesSite

	name1 := acctest.RandomName()
	name2 := acctest.RandomName()
	domain := "msserver:adsites:domain/ZG5zLm1zX2FkX3NpdGVzX2RvbWFpbiQwLkFELTE3MA:example.local/default"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverAdsitesSiteName(domain, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverAdsitesSiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverAdsitesSiteName(domain, name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverAdsitesSiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverAdsitesSiteResource_Networks(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_microsoft_msserver_adsites_site.test_networks"
	var v microsoft.MsserverAdsitesSite

	name := acctest.RandomName()
	domain := "msserver:adsites:domain/ZG5zLm1zX2FkX3NpdGVzX2RvbWFpbiQwLkFELTE3MA:example.local/default"
	network1 := "network/ZG5zLm5ldHdvcmskMTAuMTAuMC4wLzE2LzA:10.10.0.0/16/default"
	network2 := "network/ZG5zLm5ldHdvcmskMTEuMC4wLjAvMjQvMA:11.0.0.0/24/default"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverAdsitesSiteNetworks(domain, name, []string{network1}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverAdsitesSiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "networks.0", network1),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverAdsitesSiteNetworks(domain, name, []string{network2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverAdsitesSiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "networks.0", network2),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverAdsitesSiteNetworks(domain, name, []string{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverAdsitesSiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "networks.#", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckMsserverAdsitesSiteExists(ctx context.Context, resourceName string, v *microsoft.MsserverAdsitesSite) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.MicrosoftAPI.
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

func testAccCheckMsserverAdsitesSiteDestroy(ctx context.Context, v *microsoft.MsserverAdsitesSite) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.MicrosoftAPI.
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

func testAccCheckMsserverAdsitesSiteDisappears(ctx context.Context, v *microsoft.MsserverAdsitesSite) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.MicrosoftAPI.
			MsserverAdsitesSiteAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccMsserverAdsitesSiteBasicConfig(domain, name string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver_adsites_site" "test" {
	domain = %q
	name = %q
}
`, domain, name)
}

func testAccMsserverAdsitesSiteDomain(domain, name string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver_adsites_site" "test_domain" {
    domain = %q
	name = %q
}
`, domain, name)
}

func testAccMsserverAdsitesSiteName(domain, name string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver_adsites_site" "test_name" {
    domain = %q
	name = %q
}
`, domain, name)
}

func testAccMsserverAdsitesSiteNetworks(domain, name string, networks []string) string {
	networksStr := `[]`
	if len(networks) > 0 {
		networksStr = utils.ConvertStringSliceToHCL(networks)
	}

	return fmt.Sprintf(`
resource "nios_microsoft_msserver_adsites_site" "test_networks" {
    domain = %q
	name = %q
    networks = %s
}
`, domain, name, networksStr)
}
