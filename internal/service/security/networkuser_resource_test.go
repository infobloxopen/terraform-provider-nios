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

var readableAttributesForNetworkuser = "address,address_object,data_source,data_source_ip,domainname,first_seen_time,guid,last_seen_time,last_updated_time,logon_id,logout_time,name,network,network_view,user_status"

func TestAccNetworkuserResource_basic(t *testing.T) {
	var resourceName = "nios_security_networkuser.test"
	var v security.Networkuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkuserBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkuserResource_disappears(t *testing.T) {
	resourceName := "nios_security_networkuser.test"
	var v security.Networkuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetworkuserDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkuserBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					testAccCheckNetworkuserDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNetworkuserResource_Ref(t *testing.T) {
	var resourceName = "nios_security_networkuser.test_ref"
	var v security.Networkuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkuserRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkuserRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkuserResource_Address(t *testing.T) {
	var resourceName = "nios_security_networkuser.test_address"
	var v security.Networkuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkuserAddress("ADDRESS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "ADDRESS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkuserAddress("ADDRESS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "ADDRESS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkuserResource_Domainname(t *testing.T) {
	var resourceName = "nios_security_networkuser.test_domainname"
	var v security.Networkuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkuserDomainname("DOMAINNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domainname", "DOMAINNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkuserDomainname("DOMAINNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domainname", "DOMAINNAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkuserResource_FirstSeenTime(t *testing.T) {
	var resourceName = "nios_security_networkuser.test_first_seen_time"
	var v security.Networkuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkuserFirstSeenTime("FIRST_SEEN_TIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "first_seen_time", "FIRST_SEEN_TIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkuserFirstSeenTime("FIRST_SEEN_TIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "first_seen_time", "FIRST_SEEN_TIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkuserResource_Guid(t *testing.T) {
	var resourceName = "nios_security_networkuser.test_guid"
	var v security.Networkuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkuserGuid("GUID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "guid", "GUID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkuserGuid("GUID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "guid", "GUID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkuserResource_LastSeenTime(t *testing.T) {
	var resourceName = "nios_security_networkuser.test_last_seen_time"
	var v security.Networkuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkuserLastSeenTime("LAST_SEEN_TIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "last_seen_time", "LAST_SEEN_TIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkuserLastSeenTime("LAST_SEEN_TIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "last_seen_time", "LAST_SEEN_TIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkuserResource_LastUpdatedTime(t *testing.T) {
	var resourceName = "nios_security_networkuser.test_last_updated_time"
	var v security.Networkuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkuserLastUpdatedTime("LAST_UPDATED_TIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "last_updated_time", "LAST_UPDATED_TIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkuserLastUpdatedTime("LAST_UPDATED_TIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "last_updated_time", "LAST_UPDATED_TIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkuserResource_LogonId(t *testing.T) {
	var resourceName = "nios_security_networkuser.test_logon_id"
	var v security.Networkuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkuserLogonId("LOGON_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logon_id", "LOGON_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkuserLogonId("LOGON_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logon_id", "LOGON_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkuserResource_LogoutTime(t *testing.T) {
	var resourceName = "nios_security_networkuser.test_logout_time"
	var v security.Networkuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkuserLogoutTime("LOGOUT_TIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logout_time", "LOGOUT_TIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkuserLogoutTime("LOGOUT_TIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logout_time", "LOGOUT_TIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkuserResource_Name(t *testing.T) {
	var resourceName = "nios_security_networkuser.test_name"
	var v security.Networkuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkuserName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkuserName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkuserResource_NetworkView(t *testing.T) {
	var resourceName = "nios_security_networkuser.test_network_view"
	var v security.Networkuser

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkuserNetworkView("NETWORK_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkuserNetworkView("NETWORK_VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckNetworkuserExists(ctx context.Context, resourceName string, v *security.Networkuser) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SecurityAPI.
			NetworkuserAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForNetworkuser).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetNetworkuserResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetNetworkuserResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckNetworkuserDestroy(ctx context.Context, v *security.Networkuser) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SecurityAPI.
			NetworkuserAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForNetworkuser).
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

func testAccCheckNetworkuserDisappears(ctx context.Context, v *security.Networkuser) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SecurityAPI.
			NetworkuserAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccNetworkuserBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_security_networkuser" "test" {
}
`)
}

func testAccNetworkuserRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_security_networkuser" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccNetworkuserAddress(address string) string {
	return fmt.Sprintf(`
resource "nios_security_networkuser" "test_address" {
    address = %q
}
`, address)
}

func testAccNetworkuserDomainname(domainname string) string {
	return fmt.Sprintf(`
resource "nios_security_networkuser" "test_domainname" {
    domainname = %q
}
`, domainname)
}

func testAccNetworkuserFirstSeenTime(firstSeenTime string) string {
	return fmt.Sprintf(`
resource "nios_security_networkuser" "test_first_seen_time" {
    first_seen_time = %q
}
`, firstSeenTime)
}

func testAccNetworkuserGuid(guid string) string {
	return fmt.Sprintf(`
resource "nios_security_networkuser" "test_guid" {
    guid = %q
}
`, guid)
}

func testAccNetworkuserLastSeenTime(lastSeenTime string) string {
	return fmt.Sprintf(`
resource "nios_security_networkuser" "test_last_seen_time" {
    last_seen_time = %q
}
`, lastSeenTime)
}

func testAccNetworkuserLastUpdatedTime(lastUpdatedTime string) string {
	return fmt.Sprintf(`
resource "nios_security_networkuser" "test_last_updated_time" {
    last_updated_time = %q
}
`, lastUpdatedTime)
}

func testAccNetworkuserLogonId(logonId string) string {
	return fmt.Sprintf(`
resource "nios_security_networkuser" "test_logon_id" {
    logon_id = %q
}
`, logonId)
}

func testAccNetworkuserLogoutTime(logoutTime string) string {
	return fmt.Sprintf(`
resource "nios_security_networkuser" "test_logout_time" {
    logout_time = %q
}
`, logoutTime)
}

func testAccNetworkuserName(name string) string {
	return fmt.Sprintf(`
resource "nios_security_networkuser" "test_name" {
    name = %q
}
`, name)
}

func testAccNetworkuserNetworkView(networkView string) string {
	return fmt.Sprintf(`
resource "nios_security_networkuser" "test_network_view" {
    network_view = %q
}
`, networkView)
}
