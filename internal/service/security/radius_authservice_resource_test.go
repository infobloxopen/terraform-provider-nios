package security_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForRadiusAuthservice = "acct_retries,acct_timeout,auth_retries,auth_timeout,cache_ttl,comment,disable,enable_cache,mode,name,recovery_interval,servers"

func TestAccRadiusAuthserviceResource_basic(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test"
	var v security.RadiusAuthservice
	name := acctest.RandomNameWithPrefix("radius-authservice")
	servers := []map[string]interface{}{
		{
			"acct_port":      1813,
			"address":        "2.2.3.1",
			"auth_port":      1812,
			"auth_type":      "PAP",
			"disable":        false,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceBasicConfig(name, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "acct_retries", "1000"),
					resource.TestCheckResourceAttr(resourceName, "acct_timeout", "5000"),
					resource.TestCheckResourceAttr(resourceName, "auth_retries", "6"),
					resource.TestCheckResourceAttr(resourceName, "auth_timeout", "5000"),
					resource.TestCheckResourceAttr(resourceName, "cache_ttl", "3600"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_cache", "false"),
					resource.TestCheckResourceAttr(resourceName, "mode", "HUNT_GROUP"),
					resource.TestCheckResourceAttr(resourceName, "recovery_interval", "30"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_disappears(t *testing.T) {
	resourceName := "nios_security_radius_authservice.test"
	var v security.RadiusAuthservice
	name := acctest.RandomNameWithPrefix("radius-authservice")
	servers := []map[string]interface{}{
		{
			"acct_port":      1813,
			"address":        "2.2.3.1",
			"auth_port":      1812,
			"auth_type":      "PAP",
			"disable":        false,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRadiusAuthserviceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRadiusAuthserviceBasicConfig(name, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					testAccCheckRadiusAuthserviceDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRadiusAuthserviceResource_AcctRetries(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_acct_retries"
	var v security.RadiusAuthservice
	name := acctest.RandomNameWithPrefix("radius-authservice")
	servers := []map[string]interface{}{
		{
			"acct_port":      1813,
			"address":        "2.2.3.1",
			"auth_port":      1812,
			"auth_type":      "PAP",
			"disable":        false,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceAcctRetries(name, servers, 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "acct_retries", "20"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceAcctRetries(name, servers, 30),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "acct_retries", "30"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_AcctTimeout(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_acct_timeout"
	var v security.RadiusAuthservice
	name := acctest.RandomNameWithPrefix("radius-authservice")
	servers := []map[string]interface{}{
		{
			"acct_port":      1813,
			"address":        "2.2.3.1",
			"auth_port":      1812,
			"auth_type":      "PAP",
			"disable":        false,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceAcctTimeout(name, servers, 3600),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "acct_timeout", "3600"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceAcctTimeout(name, servers, 7200),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "acct_timeout", "7200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_AuthRetries(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_auth_retries"
	var v security.RadiusAuthservice
	name := acctest.RandomNameWithPrefix("radius-authservice")
	servers := []map[string]interface{}{
		{
			"acct_port":      1813,
			"address":        "2.2.3.1",
			"auth_port":      1812,
			"auth_type":      "PAP",
			"disable":        false,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceAuthRetries(name, servers, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_retries", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceAuthRetries(name, servers, 7),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_retries", "7"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_AuthTimeout(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_auth_timeout"
	var v security.RadiusAuthservice
	name := acctest.RandomNameWithPrefix("radius-authservice")
	servers := []map[string]interface{}{
		{
			"acct_port":      1813,
			"address":        "2.2.3.1",
			"auth_port":      1812,
			"auth_type":      "PAP",
			"disable":        false,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceAuthTimeout(name, servers, 4000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_timeout", "4000"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceAuthTimeout(name, servers, 4500),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_timeout", "4500"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_CacheTtl(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_cache_ttl"
	var v security.RadiusAuthservice
	name := acctest.RandomNameWithPrefix("radius-authservice")
	servers := []map[string]interface{}{
		{
			"acct_port":      1813,
			"address":        "2.2.3.1",
			"auth_port":      1812,
			"auth_type":      "PAP",
			"disable":        false,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceCacheTtl(name, servers, 4000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cache_ttl", "4000"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceCacheTtl(name, servers, 4500),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cache_ttl", "4500"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_Comment(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_comment"
	var v security.RadiusAuthservice
	name := acctest.RandomNameWithPrefix("radius-authservice")
	servers := []map[string]interface{}{
		{
			"acct_port":      1813,
			"address":        "2.2.3.1",
			"auth_port":      1812,
			"auth_type":      "PAP",
			"disable":        false,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceComment(name, servers, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceComment(name, servers, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_Disable(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_disable"
	var v security.RadiusAuthservice
	name := acctest.RandomNameWithPrefix("radius-authservice")
	servers := []map[string]interface{}{
		{
			"acct_port":      1813,
			"address":        "2.2.3.1",
			"auth_port":      1812,
			"auth_type":      "PAP",
			"disable":        false,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceDisable(name, servers, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceDisable(name, servers, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_EnableCache(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_enable_cache"
	var v security.RadiusAuthservice
	name := acctest.RandomNameWithPrefix("radius-authservice")
	servers := []map[string]interface{}{
		{
			"acct_port":      1813,
			"address":        "2.2.3.1",
			"auth_port":      1812,
			"auth_type":      "PAP",
			"disable":        false,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceEnableCache(name, servers, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_cache", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceEnableCache(name, servers, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_cache", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_Mode(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_mode"
	var v security.RadiusAuthservice
	name := acctest.RandomNameWithPrefix("radius-authservice")
	servers := []map[string]interface{}{
		{
			"acct_port":      1813,
			"address":        "2.2.3.1",
			"auth_port":      1812,
			"auth_type":      "PAP",
			"disable":        false,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceMode(name, servers, "ROUND_ROBIN"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mode", "ROUND_ROBIN"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceMode(name, servers, "HUNT_GROUP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mode", "HUNT_GROUP"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_Name(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_name"
	var v security.RadiusAuthservice
	name := acctest.RandomNameWithPrefix("radius-authservice")
	nameUpdate := acctest.RandomNameWithPrefix("radius-authservice")
	servers := []map[string]interface{}{
		{
			"acct_port":      1813,
			"address":        "2.2.3.1",
			"auth_port":      1812,
			"auth_type":      "PAP",
			"disable":        false,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceName(name, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceName(nameUpdate, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_RecoveryInterval(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_recovery_interval"
	var v security.RadiusAuthservice
	name := acctest.RandomNameWithPrefix("radius-authservice")
	servers := []map[string]interface{}{
		{
			"acct_port":      1813,
			"address":        "2.2.3.1",
			"auth_port":      1812,
			"auth_type":      "PAP",
			"disable":        false,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceRecoveryInterval(name, servers, 45),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recovery_interval", "45"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceRecoveryInterval(name, servers, 60),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recovery_interval", "60"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRadiusAuthserviceResource_Servers(t *testing.T) {
	var resourceName = "nios_security_radius_authservice.test_servers"
	var v security.RadiusAuthservice
	name := acctest.RandomNameWithPrefix("radius-authservice")
	servers := []map[string]interface{}{
		{
			"acct_port":      1813,
			"address":        "2.2.3.1",
			"auth_port":      1812,
			"auth_type":      "PAP",
			"disable":        false,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	serversUpdate := []map[string]interface{}{
		{
			"acct_port":      1813,
			"address":        "2.2.3.2",
			"auth_port":      1812,
			"auth_type":      "CHAP",
			"disable":        false,
			"use_accounting": true,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRadiusAuthserviceServers(name, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.address", "2.2.3.1"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.auth_port", "1812"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.acct_port", "1813"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.auth_type", "PAP"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.use_accounting", "false"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.use_mgmt_port", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRadiusAuthserviceServers(name, serversUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.address", "2.2.3.2"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.auth_port", "1812"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.acct_port", "1813"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.auth_type", "CHAP"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.use_accounting", "true"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.use_mgmt_port", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRadiusAuthserviceExists(ctx context.Context, resourceName string, v *security.RadiusAuthservice) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SecurityAPI.
			RadiusAuthserviceAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRadiusAuthservice).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRadiusAuthserviceResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRadiusAuthserviceResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRadiusAuthserviceDestroy(ctx context.Context, v *security.RadiusAuthservice) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SecurityAPI.
			RadiusAuthserviceAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRadiusAuthservice).
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

func testAccCheckRadiusAuthserviceDisappears(ctx context.Context, v *security.RadiusAuthservice) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SecurityAPI.
			RadiusAuthserviceAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRadiusAuthserviceBasicConfig(name string, servers []map[string]any) string {
	serversString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test" {
	name = %q
	servers = %s
}
`, name, serversString)
}

func testAccRadiusAuthserviceAcctRetries(name string, servers []map[string]any, acctRetries int) string {
	serversString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_acct_retries" {
	name = %q
	servers = %s
    acct_retries = %d
}
`, name, serversString, acctRetries)
}

func testAccRadiusAuthserviceAcctTimeout(name string, servers []map[string]any, acctTimeout int) string {
	serversString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_acct_timeout" {
	name = %q
	servers = %s
    acct_timeout = %d
}
`, name, serversString, acctTimeout)
}

func testAccRadiusAuthserviceAuthRetries(name string, servers []map[string]any, authRetries int) string {
	serversString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_auth_retries" {
	name = %q
	servers = %s
    auth_retries = %d
}
`, name, serversString, authRetries)
}

func testAccRadiusAuthserviceAuthTimeout(name string, servers []map[string]any, authTimeout int) string {
	serversString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_auth_timeout" {
	name = %q
	servers = %s
    auth_timeout = %d
}
`, name, serversString, authTimeout)
}

func testAccRadiusAuthserviceCacheTtl(name string, servers []map[string]any, cacheTtl int) string {
	serversString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_cache_ttl" {
	name = %q
	servers = %s
    cache_ttl = %d
}
`, name, serversString, cacheTtl)
}

func testAccRadiusAuthserviceComment(name string, servers []map[string]any, comment string) string {
	serversString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_comment" {
	name = %q
	servers = %s
    comment = %q
}
`, name, serversString, comment)
}

func testAccRadiusAuthserviceDisable(name string, servers []map[string]any, disable bool) string {
	serversString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_disable" {
	name = %q
	servers = %s
    disable = %t
}
`, name, serversString, disable)
}

func testAccRadiusAuthserviceEnableCache(name string, servers []map[string]any, enableCache bool) string {
	serversString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_enable_cache" {
	name = %q
	servers = %s
    enable_cache = %t
}
`, name, serversString, enableCache)
}

func testAccRadiusAuthserviceMode(name string, servers []map[string]any, mode string) string {
	serversString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_mode" {
	name = %q
	servers = %s
    mode = %q
}
`, name, serversString, mode)
}

func testAccRadiusAuthserviceName(name string, servers []map[string]any) string {
	serversString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_name" {
    name = %q
	servers = %s
}
`, name, serversString)
}

func testAccRadiusAuthserviceRecoveryInterval(name string, servers []map[string]any, recoveryInterval int) string {
	serversString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_recovery_interval" {
	name = %q
	servers = %s
    recovery_interval = %d
}
`, name, serversString, recoveryInterval)
}

func testAccRadiusAuthserviceServers(name string, servers []map[string]any) string {
	serversString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_radius_authservice" "test_servers" {
	name = %q
    servers = %s
}
`, name, serversString)
}
