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

var readableAttributesForTacacsplusAuthservice = "acct_retries,acct_timeout,auth_retries,auth_timeout,comment,disable,name,servers"

func TestAccTacacsplusAuthserviceResource_basic(t *testing.T) {
	var resourceName = "nios_security_tacacsplus_authservice.test"
	var v security.TacacsplusAuthservice
	name := acctest.RandomNameWithPrefix("tacacsplus_authservice")
	servers := []map[string]any{
		{
			"address":        "2.2.3.3",
			"auth_type":      "CHAP",
			"disable":        false,
			"port":           49,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTacacsplusAuthserviceBasicConfig(name, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.address", "2.2.3.3"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.auth_type", "CHAP"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.port", "49"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.shared_secret", "test"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "acct_retries", "0"),
					resource.TestCheckResourceAttr(resourceName, "acct_timeout", "1000"),
					resource.TestCheckResourceAttr(resourceName, "auth_retries", "0"),
					resource.TestCheckResourceAttr(resourceName, "auth_timeout", "5000"),
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTacacsplusAuthserviceResource_disappears(t *testing.T) {
	resourceName := "nios_security_tacacsplus_authservice.test"
	var v security.TacacsplusAuthservice
	name := acctest.RandomNameWithPrefix("tacacsplus_authservice")
	servers := []map[string]any{
		{
			"address":        "2.2.3.3",
			"auth_type":      "CHAP",
			"disable":        false,
			"port":           49,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTacacsplusAuthserviceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccTacacsplusAuthserviceBasicConfig(name, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					testAccCheckTacacsplusAuthserviceDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccTacacsplusAuthserviceResource_AcctRetries(t *testing.T) {
	var resourceName = "nios_security_tacacsplus_authservice.test_acct_retries"
	var v security.TacacsplusAuthservice
	name := acctest.RandomNameWithPrefix("tacacsplus_authservice")
	servers := []map[string]any{
		{
			"address":        "2.2.3.3",
			"auth_type":      "CHAP",
			"disable":        false,
			"port":           49,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTacacsplusAuthserviceAcctRetries(name, servers, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "acct_retries", "1"),
				),
			},
			// Update and Read
			{
				Config: testAccTacacsplusAuthserviceAcctRetries(name, servers, 3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "acct_retries", "3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTacacsplusAuthserviceResource_AcctTimeout(t *testing.T) {
	var resourceName = "nios_security_tacacsplus_authservice.test_acct_timeout"
	var v security.TacacsplusAuthservice
	name := acctest.RandomNameWithPrefix("tacacsplus_authservice")
	servers := []map[string]any{
		{
			"address":        "2.2.3.3",
			"auth_type":      "CHAP",
			"disable":        false,
			"port":           49,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTacacsplusAuthserviceAcctTimeout(name, servers, 3600),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "acct_timeout", "3600"),
				),
			},
			// Update and Read
			{
				Config: testAccTacacsplusAuthserviceAcctTimeout(name, servers, 7200),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "acct_timeout", "7200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTacacsplusAuthserviceResource_AuthRetries(t *testing.T) {
	var resourceName = "nios_security_tacacsplus_authservice.test_auth_retries"
	var v security.TacacsplusAuthservice
	name := acctest.RandomNameWithPrefix("tacacsplus_authservice")
	servers := []map[string]any{
		{
			"address":        "2.2.3.3",
			"auth_type":      "CHAP",
			"disable":        false,
			"port":           49,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTacacsplusAuthserviceAuthRetries(name, servers, 2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_retries", "2"),
				),
			},
			// Update and Read
			{
				Config: testAccTacacsplusAuthserviceAuthRetries(name, servers, 4),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_retries", "4"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTacacsplusAuthserviceResource_AuthTimeout(t *testing.T) {
	var resourceName = "nios_security_tacacsplus_authservice.test_auth_timeout"
	var v security.TacacsplusAuthservice
	name := acctest.RandomNameWithPrefix("tacacsplus_authservice")
	servers := []map[string]any{
		{
			"address":        "2.2.3.3",
			"auth_type":      "CHAP",
			"disable":        false,
			"port":           49,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTacacsplusAuthserviceAuthTimeout(name, servers, 7000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_timeout", "7000"),
				),
			},
			// Update and Read
			{
				Config: testAccTacacsplusAuthserviceAuthTimeout(name, servers, 6000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_timeout", "6000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTacacsplusAuthserviceResource_Comment(t *testing.T) {
	var resourceName = "nios_security_tacacsplus_authservice.test_comment"
	var v security.TacacsplusAuthservice
	name := acctest.RandomNameWithPrefix("tacacsplus_authservice")
	servers := []map[string]any{
		{
			"address":        "2.2.3.3",
			"auth_type":      "CHAP",
			"disable":        false,
			"port":           49,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTacacsplusAuthserviceComment(name, servers, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccTacacsplusAuthserviceComment(name, servers, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTacacsplusAuthserviceResource_Disable(t *testing.T) {
	var resourceName = "nios_security_tacacsplus_authservice.test_disable"
	var v security.TacacsplusAuthservice
	name := acctest.RandomNameWithPrefix("tacacsplus_authservice")
	servers := []map[string]any{
		{
			"address":        "2.2.3.3",
			"auth_type":      "CHAP",
			"disable":        false,
			"port":           49,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTacacsplusAuthserviceDisable(name, servers, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccTacacsplusAuthserviceDisable(name, servers, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTacacsplusAuthserviceResource_Name(t *testing.T) {
	var resourceName = "nios_security_tacacsplus_authservice.test_name"
	var v security.TacacsplusAuthservice
	name := acctest.RandomNameWithPrefix("tacacsplus_authservice")
	nameUpdate := acctest.RandomNameWithPrefix("tacacsplus_authservice")
	servers := []map[string]any{
		{
			"address":        "2.2.3.3",
			"auth_type":      "CHAP",
			"disable":        false,
			"port":           49,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTacacsplusAuthserviceName(name, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccTacacsplusAuthserviceName(nameUpdate, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTacacsplusAuthserviceResource_Servers(t *testing.T) {
	var resourceName = "nios_security_tacacsplus_authservice.test_servers"
	var v security.TacacsplusAuthservice
	name := acctest.RandomNameWithPrefix("tacacsplus_authservice")
	servers := []map[string]any{
		{
			"address":        "2.2.3.3",
			"auth_type":      "CHAP",
			"disable":        false,
			"port":           49,
			"use_accounting": false,
			"use_mgmt_port":  false,
			"shared_secret":  "test",
		},
	}
	serversUpdate := []map[string]any{
		{
			"address":        "2.2.1.3",
			"auth_type":      "PAP",
			"disable":        true,
			"port":           49,
			"use_accounting": false,
			"use_mgmt_port":  true,
			"shared_secret":  "testing_key",
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTacacsplusAuthserviceServers(name, servers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.address", "2.2.3.3"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.auth_type", "CHAP"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.port", "49"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.use_accounting", "false"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.use_mgmt_port", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccTacacsplusAuthserviceServers(name, serversUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTacacsplusAuthserviceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.address", "2.2.1.3"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.auth_type", "PAP"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.disable", "true"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.port", "49"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.use_accounting", "false"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.use_mgmt_port", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckTacacsplusAuthserviceExists(ctx context.Context, resourceName string, v *security.TacacsplusAuthservice) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SecurityAPI.
			TacacsplusAuthserviceAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForTacacsplusAuthservice).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetTacacsplusAuthserviceResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetTacacsplusAuthserviceResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckTacacsplusAuthserviceDestroy(ctx context.Context, v *security.TacacsplusAuthservice) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SecurityAPI.
			TacacsplusAuthserviceAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForTacacsplusAuthservice).
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

func testAccCheckTacacsplusAuthserviceDisappears(ctx context.Context, v *security.TacacsplusAuthservice) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SecurityAPI.
			TacacsplusAuthserviceAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccTacacsplusAuthserviceBasicConfig(name string, servers []map[string]any) string {
	serversStr := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_tacacsplus_authservice" "test" {
	name = %q
	servers = %s
}
`, name, serversStr)
}

func testAccTacacsplusAuthserviceAcctRetries(name string, servers []map[string]any, acctRetries int) string {
	serversStr := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_tacacsplus_authservice" "test_acct_retries" {
	name = %q
	servers = %s
    acct_retries = %d
}
`, name, serversStr, acctRetries)
}

func testAccTacacsplusAuthserviceAcctTimeout(name string, servers []map[string]any, acctTimeout int) string {
	serversStr := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_tacacsplus_authservice" "test_acct_timeout" {
	name = %q
	servers = %s
    acct_timeout = %d
}
`, name, serversStr, acctTimeout)
}

func testAccTacacsplusAuthserviceAuthRetries(name string, servers []map[string]any, authRetries int) string {
	serversStr := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_tacacsplus_authservice" "test_auth_retries" {
	name = %q
	servers = %s
    auth_retries = %d
}
`, name, serversStr, authRetries)
}

func testAccTacacsplusAuthserviceAuthTimeout(name string, servers []map[string]any, authTimeout int) string {
	serversStr := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_tacacsplus_authservice" "test_auth_timeout" {
	name = %q
	servers = %s
    auth_timeout = %d
}
`, name, serversStr, authTimeout)
}

func testAccTacacsplusAuthserviceComment(name string, servers []map[string]any, comment string) string {
	serversStr := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_tacacsplus_authservice" "test_comment" {
	name = %q
	servers = %s
    comment = %q
}
`, name, serversStr, comment)
}

func testAccTacacsplusAuthserviceDisable(name string, servers []map[string]any, disable bool) string {
	serversStr := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_tacacsplus_authservice" "test_disable" {
	name = %q
	servers = %s
    disable = %t
}
`, name, serversStr, disable)
}

func testAccTacacsplusAuthserviceName(name string, servers []map[string]any) string {
	serversStr := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_tacacsplus_authservice" "test_name" {
    name = %q
	servers = %s
}
`, name, serversStr)
}

func testAccTacacsplusAuthserviceServers(name string, servers []map[string]any) string {
	serversStr := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_tacacsplus_authservice" "test_servers" {
	name = %q
    servers = %s
}
`, name, serversStr)
}
