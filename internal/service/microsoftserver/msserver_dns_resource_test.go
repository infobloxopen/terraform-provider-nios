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

var readableAttributesForMsserverDns = "address,enable_dns_reports_sync,login_name,synchronization_interval,use_enable_dns_reports_sync,use_login,use_synchronization_interval"

func TestAccMsserverDnsResource_basic(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_dns.test"
	var v microsoftserver.MsserverDns

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDnsBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDnsExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverDnsResource_disappears(t *testing.T) {
	resourceName := "nios_microsoftserver_msserver_dns.test"
	var v microsoftserver.MsserverDns

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMsserverDnsDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMsserverDnsBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDnsExists(context.Background(), resourceName, &v),
					testAccCheckMsserverDnsDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccMsserverDnsResource_Ref(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_dns.test_ref"
	var v microsoftserver.MsserverDns

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDnsRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDnsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDnsRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDnsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverDnsResource_EnableDnsReportsSync(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_dns.test_enable_dns_reports_sync"
	var v microsoftserver.MsserverDns

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDnsEnableDnsReportsSync("ENABLE_DNS_REPORTS_SYNC_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDnsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_dns_reports_sync", "ENABLE_DNS_REPORTS_SYNC_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDnsEnableDnsReportsSync("ENABLE_DNS_REPORTS_SYNC_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDnsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_dns_reports_sync", "ENABLE_DNS_REPORTS_SYNC_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverDnsResource_LoginName(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_dns.test_login_name"
	var v microsoftserver.MsserverDns

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDnsLoginName("LOGIN_NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDnsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "login_name", "LOGIN_NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDnsLoginName("LOGIN_NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDnsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "login_name", "LOGIN_NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverDnsResource_LoginPassword(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_dns.test_login_password"
	var v microsoftserver.MsserverDns

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDnsLoginPassword("LOGIN_PASSWORD_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDnsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "login_password", "LOGIN_PASSWORD_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDnsLoginPassword("LOGIN_PASSWORD_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDnsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "login_password", "LOGIN_PASSWORD_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverDnsResource_SynchronizationInterval(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_dns.test_synchronization_interval"
	var v microsoftserver.MsserverDns

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDnsSynchronizationInterval("SYNCHRONIZATION_INTERVAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDnsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "synchronization_interval", "SYNCHRONIZATION_INTERVAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDnsSynchronizationInterval("SYNCHRONIZATION_INTERVAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDnsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "synchronization_interval", "SYNCHRONIZATION_INTERVAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverDnsResource_UseEnableDnsReportsSync(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_dns.test_use_enable_dns_reports_sync"
	var v microsoftserver.MsserverDns

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDnsUseEnableDnsReportsSync("USE_ENABLE_DNS_REPORTS_SYNC_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDnsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_dns_reports_sync", "USE_ENABLE_DNS_REPORTS_SYNC_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDnsUseEnableDnsReportsSync("USE_ENABLE_DNS_REPORTS_SYNC_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDnsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_dns_reports_sync", "USE_ENABLE_DNS_REPORTS_SYNC_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverDnsResource_UseLogin(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_dns.test_use_login"
	var v microsoftserver.MsserverDns

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDnsUseLogin("USE_LOGIN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDnsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_login", "USE_LOGIN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDnsUseLogin("USE_LOGIN_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDnsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_login", "USE_LOGIN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverDnsResource_UseSynchronizationInterval(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_dns.test_use_synchronization_interval"
	var v microsoftserver.MsserverDns

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDnsUseSynchronizationInterval("USE_SYNCHRONIZATION_INTERVAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDnsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_synchronization_interval", "USE_SYNCHRONIZATION_INTERVAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDnsUseSynchronizationInterval("USE_SYNCHRONIZATION_INTERVAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDnsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_synchronization_interval", "USE_SYNCHRONIZATION_INTERVAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckMsserverDnsExists(ctx context.Context, resourceName string, v *microsoftserver.MsserverDns) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.MicrosoftServerAPI.
			MsserverDnsAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForMsserverDns).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetMsserverDnsResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetMsserverDnsResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckMsserverDnsDestroy(ctx context.Context, v *microsoftserver.MsserverDns) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.MicrosoftServerAPI.
			MsserverDnsAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForMsserverDns).
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

func testAccCheckMsserverDnsDisappears(ctx context.Context, v *microsoftserver.MsserverDns) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.MicrosoftServerAPI.
			MsserverDnsAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccMsserverDnsBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_dns" "test" {
}
`)
}

func testAccMsserverDnsRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_dns" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccMsserverDnsEnableDnsReportsSync(enableDnsReportsSync string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_dns" "test_enable_dns_reports_sync" {
    enable_dns_reports_sync = %q
}
`, enableDnsReportsSync)
}

func testAccMsserverDnsLoginName(loginName string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_dns" "test_login_name" {
    login_name = %q
}
`, loginName)
}

func testAccMsserverDnsLoginPassword(loginPassword string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_dns" "test_login_password" {
    login_password = %q
}
`, loginPassword)
}

func testAccMsserverDnsSynchronizationInterval(synchronizationInterval string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_dns" "test_synchronization_interval" {
    synchronization_interval = %q
}
`, synchronizationInterval)
}

func testAccMsserverDnsUseEnableDnsReportsSync(useEnableDnsReportsSync string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_dns" "test_use_enable_dns_reports_sync" {
    use_enable_dns_reports_sync = %q
}
`, useEnableDnsReportsSync)
}

func testAccMsserverDnsUseLogin(useLogin string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_dns" "test_use_login" {
    use_login = %q
}
`, useLogin)
}

func testAccMsserverDnsUseSynchronizationInterval(useSynchronizationInterval string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_dns" "test_use_synchronization_interval" {
    use_synchronization_interval = %q
}
`, useSynchronizationInterval)
}
