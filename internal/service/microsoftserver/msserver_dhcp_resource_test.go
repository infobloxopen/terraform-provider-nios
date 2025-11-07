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

var readableAttributesForMsserverDhcp = "address,comment,dhcp_utilization,dhcp_utilization_status,dynamic_hosts,last_sync_ts,login_name,network_view,next_sync_control,read_only,server_name,static_hosts,status,status_detail,status_last_updated,supports_failover,synchronization_interval,total_hosts,use_login,use_synchronization_interval"

func TestAccMsserverDhcpResource_basic(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_dhcp.test"
	var v microsoftserver.MsserverDhcp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDhcpBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDhcpExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverDhcpResource_disappears(t *testing.T) {
	resourceName := "nios_microsoftserver_msserver_dhcp.test"
	var v microsoftserver.MsserverDhcp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMsserverDhcpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMsserverDhcpBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDhcpExists(context.Background(), resourceName, &v),
					testAccCheckMsserverDhcpDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccMsserverDhcpResource_Ref(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_dhcp.test_ref"
	var v microsoftserver.MsserverDhcp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDhcpRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDhcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDhcpRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDhcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverDhcpResource_LoginName(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_dhcp.test_login_name"
	var v microsoftserver.MsserverDhcp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDhcpLoginName("LOGIN_NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDhcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "login_name", "LOGIN_NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDhcpLoginName("LOGIN_NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDhcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "login_name", "LOGIN_NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverDhcpResource_LoginPassword(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_dhcp.test_login_password"
	var v microsoftserver.MsserverDhcp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDhcpLoginPassword("LOGIN_PASSWORD_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDhcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "login_password", "LOGIN_PASSWORD_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDhcpLoginPassword("LOGIN_PASSWORD_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDhcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "login_password", "LOGIN_PASSWORD_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverDhcpResource_NextSyncControl(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_dhcp.test_next_sync_control"
	var v microsoftserver.MsserverDhcp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDhcpNextSyncControl("NEXT_SYNC_CONTROL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDhcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "next_sync_control", "NEXT_SYNC_CONTROL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDhcpNextSyncControl("NEXT_SYNC_CONTROL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDhcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "next_sync_control", "NEXT_SYNC_CONTROL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverDhcpResource_Status(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_dhcp.test_status"
	var v microsoftserver.MsserverDhcp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDhcpStatus("STATUS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDhcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "status", "STATUS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDhcpStatus("STATUS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDhcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "status", "STATUS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverDhcpResource_SynchronizationInterval(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_dhcp.test_synchronization_interval"
	var v microsoftserver.MsserverDhcp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDhcpSynchronizationInterval("SYNCHRONIZATION_INTERVAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDhcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "synchronization_interval", "SYNCHRONIZATION_INTERVAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDhcpSynchronizationInterval("SYNCHRONIZATION_INTERVAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDhcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "synchronization_interval", "SYNCHRONIZATION_INTERVAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverDhcpResource_UseLogin(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_dhcp.test_use_login"
	var v microsoftserver.MsserverDhcp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDhcpUseLogin("USE_LOGIN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDhcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_login", "USE_LOGIN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDhcpUseLogin("USE_LOGIN_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDhcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_login", "USE_LOGIN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverDhcpResource_UseSynchronizationInterval(t *testing.T) {
	var resourceName = "nios_microsoftserver_msserver_dhcp.test_use_synchronization_interval"
	var v microsoftserver.MsserverDhcp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDhcpUseSynchronizationInterval("USE_SYNCHRONIZATION_INTERVAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDhcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_synchronization_interval", "USE_SYNCHRONIZATION_INTERVAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDhcpUseSynchronizationInterval("USE_SYNCHRONIZATION_INTERVAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverDhcpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_synchronization_interval", "USE_SYNCHRONIZATION_INTERVAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckMsserverDhcpExists(ctx context.Context, resourceName string, v *microsoftserver.MsserverDhcp) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.MicrosoftServerAPI.
			MsserverDhcpAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForMsserverDhcp).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetMsserverDhcpResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetMsserverDhcpResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckMsserverDhcpDestroy(ctx context.Context, v *microsoftserver.MsserverDhcp) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.MicrosoftServerAPI.
			MsserverDhcpAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForMsserverDhcp).
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

func testAccCheckMsserverDhcpDisappears(ctx context.Context, v *microsoftserver.MsserverDhcp) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.MicrosoftServerAPI.
			MsserverDhcpAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccMsserverDhcpBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_dhcp" "test" {
}
`)
}

func testAccMsserverDhcpRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_dhcp" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccMsserverDhcpLoginName(loginName string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_dhcp" "test_login_name" {
    login_name = %q
}
`, loginName)
}

func testAccMsserverDhcpLoginPassword(loginPassword string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_dhcp" "test_login_password" {
    login_password = %q
}
`, loginPassword)
}

func testAccMsserverDhcpNextSyncControl(nextSyncControl string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_dhcp" "test_next_sync_control" {
    next_sync_control = %q
}
`, nextSyncControl)
}

func testAccMsserverDhcpStatus(status string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_dhcp" "test_status" {
    status = %q
}
`, status)
}

func testAccMsserverDhcpSynchronizationInterval(synchronizationInterval string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_dhcp" "test_synchronization_interval" {
    synchronization_interval = %q
}
`, synchronizationInterval)
}

func testAccMsserverDhcpUseLogin(useLogin string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_dhcp" "test_use_login" {
    use_login = %q
}
`, useLogin)
}

func testAccMsserverDhcpUseSynchronizationInterval(useSynchronizationInterval string) string {
	return fmt.Sprintf(`
resource "nios_microsoftserver_msserver_dhcp" "test_use_synchronization_interval" {
    use_synchronization_interval = %q
}
`, useSynchronizationInterval)
}
