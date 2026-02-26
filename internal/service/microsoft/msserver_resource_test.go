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

var readableAttributesForMsserver = "ad_domain,ad_sites,ad_user,address,comment,connection_status,connection_status_detail,dhcp_server,disabled,dns_server,dns_view,extattrs,grid_member,last_seen,log_destination,log_level,login_name,managing_member,ms_max_connection,ms_rpc_timeout_in_seconds,network_view,read_only,root_ad_domain,server_name,synchronization_min_delay,synchronization_status,synchronization_status_detail,use_log_destination,use_ms_max_connection,use_ms_rpc_timeout_in_seconds,version"

func TestAccMsserverResource_basic(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test"
	var v microsoft.Msserver

	address := "10.10.0.1"
	loginName := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverBasicConfig(address, loginName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", address),
					resource.TestCheckResourceAttr(resourceName, "login_name", loginName),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "ms_max_connection", "10"),
					resource.TestCheckResourceAttr(resourceName, "ms_rpc_timeout_in_seconds", "60"),
					resource.TestCheckResourceAttr(resourceName, "read_only", "false"),
					resource.TestCheckResourceAttr(resourceName, "synchronization_min_delay", "2"),
					resource.TestCheckResourceAttr(resourceName, "use_log_destination", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_disappears(t *testing.T) {
	resourceName := "nios_microsoft_msserver.test"
	var v microsoft.Msserver

	address := "10.10.0.1"
	loginName := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMsserverDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMsserverBasicConfig(address, loginName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					testAccCheckMsserverDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccMsserverResource_AdSites(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_ad_sites"
	var v microsoft.Msserver

	address := "10.10.0.1"
	loginName := acctest.RandomName()
	adsitesLoginName := acctest.RandomName()
	adsites := map[string]any{
		"login_name": adsitesLoginName,
	}

	updatedAdsitesLoginName := acctest.RandomName()
	updatedAdsites := map[string]any{
		"login_name": updatedAdsitesLoginName,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverAdSites(address, loginName, adsites),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ad_sites.login_name", adsitesLoginName),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverAdSites(address, loginName, updatedAdsites),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ad_sites.login_name", updatedAdsitesLoginName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_AdUser(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_ad_user"
	var v microsoft.Msserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverAdUser("AD_USER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ad_user", "AD_USER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverAdUser("AD_USER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ad_user", "AD_USER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_Address(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_address"
	var v microsoft.Msserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverAddress("ADDRESS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "ADDRESS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverAddress("ADDRESS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "ADDRESS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_Comment(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_comment"
	var v microsoft.Msserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_DhcpServer(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_dhcp_server"
	var v microsoft.Msserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDhcpServer("DHCP_SERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_server", "DHCP_SERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDhcpServer("DHCP_SERVER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_server", "DHCP_SERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_Disabled(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_disabled"
	var v microsoft.Msserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDisabled("DISABLED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "DISABLED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDisabled("DISABLED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "DISABLED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_DnsServer(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_dns_server"
	var v microsoft.Msserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDnsServer("DNS_SERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_server", "DNS_SERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDnsServer("DNS_SERVER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_server", "DNS_SERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_DnsView(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_dns_view"
	var v microsoft.Msserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDnsView("DNS_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_view", "DNS_VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDnsView("DNS_VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_view", "DNS_VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_extattrs"
	var v microsoft.Msserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_GridMember(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_grid_member"
	var v microsoft.Msserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverGridMember("GRID_MEMBER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_member", "GRID_MEMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverGridMember("GRID_MEMBER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_member", "GRID_MEMBER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_LogDestination(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_log_destination"
	var v microsoft.Msserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverLogDestination("LOG_DESTINATION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_destination", "LOG_DESTINATION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverLogDestination("LOG_DESTINATION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_destination", "LOG_DESTINATION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_LogLevel(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_log_level"
	var v microsoft.Msserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverLogLevel("LOG_LEVEL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "LOG_LEVEL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverLogLevel("LOG_LEVEL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "LOG_LEVEL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_LoginName(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_login_name"
	var v microsoft.Msserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverLoginName("LOGIN_NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "login_name", "LOGIN_NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverLoginName("LOGIN_NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "login_name", "LOGIN_NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_LoginPassword(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_login_password"
	var v microsoft.Msserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverLoginPassword("LOGIN_PASSWORD_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "login_password", "LOGIN_PASSWORD_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverLoginPassword("LOGIN_PASSWORD_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "login_password", "LOGIN_PASSWORD_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_MsMaxConnection(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_ms_max_connection"
	var v microsoft.Msserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverMsMaxConnection("MS_MAX_CONNECTION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_max_connection", "MS_MAX_CONNECTION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverMsMaxConnection("MS_MAX_CONNECTION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_max_connection", "MS_MAX_CONNECTION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_MsRpcTimeoutInSeconds(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_ms_rpc_timeout_in_seconds"
	var v microsoft.Msserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverMsRpcTimeoutInSeconds("MS_RPC_TIMEOUT_IN_SECONDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_rpc_timeout_in_seconds", "MS_RPC_TIMEOUT_IN_SECONDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverMsRpcTimeoutInSeconds("MS_RPC_TIMEOUT_IN_SECONDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_rpc_timeout_in_seconds", "MS_RPC_TIMEOUT_IN_SECONDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_NetworkView(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_network_view"
	var v microsoft.Msserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverNetworkView("NETWORK_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverNetworkView("NETWORK_VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_ReadOnly(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_read_only"
	var v microsoft.Msserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverReadOnly("READ_ONLY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "read_only", "READ_ONLY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverReadOnly("READ_ONLY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "read_only", "READ_ONLY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_SynchronizationMinDelay(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_synchronization_min_delay"
	var v microsoft.Msserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverSynchronizationMinDelay("SYNCHRONIZATION_MIN_DELAY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "synchronization_min_delay", "SYNCHRONIZATION_MIN_DELAY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverSynchronizationMinDelay("SYNCHRONIZATION_MIN_DELAY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "synchronization_min_delay", "SYNCHRONIZATION_MIN_DELAY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_UseLogDestination(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_use_log_destination"
	var v microsoft.Msserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverUseLogDestination("USE_LOG_DESTINATION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_log_destination", "USE_LOG_DESTINATION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverUseLogDestination("USE_LOG_DESTINATION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_log_destination", "USE_LOG_DESTINATION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_UseMsMaxConnection(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_use_ms_max_connection"
	var v microsoft.Msserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverUseMsMaxConnection("USE_MS_MAX_CONNECTION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ms_max_connection", "USE_MS_MAX_CONNECTION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverUseMsMaxConnection("USE_MS_MAX_CONNECTION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ms_max_connection", "USE_MS_MAX_CONNECTION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_UseMsRpcTimeoutInSeconds(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_use_ms_rpc_timeout_in_seconds"
	var v microsoft.Msserver

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverUseMsRpcTimeoutInSeconds("USE_MS_RPC_TIMEOUT_IN_SECONDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ms_rpc_timeout_in_seconds", "USE_MS_RPC_TIMEOUT_IN_SECONDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverUseMsRpcTimeoutInSeconds("USE_MS_RPC_TIMEOUT_IN_SECONDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ms_rpc_timeout_in_seconds", "USE_MS_RPC_TIMEOUT_IN_SECONDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckMsserverExists(ctx context.Context, resourceName string, v *microsoft.Msserver) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.MicrosoftAPI.
			MsserverAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForMsserver).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetMsserverResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetMsserverResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckMsserverDestroy(ctx context.Context, v *microsoft.Msserver) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.MicrosoftAPI.
			MsserverAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForMsserver).
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

func testAccCheckMsserverDisappears(ctx context.Context, v *microsoft.Msserver) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.MicrosoftAPI.
			MsserverAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccMsserverBasicConfig(address, loginName string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test" {
	address = %q
	login_name = %q
}
`, address, loginName)
}

func testAccMsserverAdSites(address, loginName string, adSites map[string]any) string {
	adSitesHCL := utils.ConvertMapToHCL(adSites)

	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_ad_sites" {
    address = %q
	login_name = %q
	ad_sites = %s
}
`, address, loginName, adSitesHCL)
}

func testAccMsserverAdUser(adUser string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_ad_user" {
    ad_user = %q
}
`, adUser)
}

func testAccMsserverAddress(address string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_address" {
    address = %q
}
`, address)
}

func testAccMsserverComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccMsserverDhcpServer(dhcpServer string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_dhcp_server" {
    dhcp_server = %q
}
`, dhcpServer)
}

func testAccMsserverDisabled(disabled string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_disabled" {
    disabled = %q
}
`, disabled)
}

func testAccMsserverDnsServer(dnsServer string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_dns_server" {
    dns_server = %q
}
`, dnsServer)
}

func testAccMsserverDnsView(dnsView string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_dns_view" {
    dns_view = %q
}
`, dnsView)
}

func testAccMsserverExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccMsserverGridMember(gridMember string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_grid_member" {
    grid_member = %q
}
`, gridMember)
}

func testAccMsserverLogDestination(logDestination string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_log_destination" {
    log_destination = %q
}
`, logDestination)
}

func testAccMsserverLogLevel(logLevel string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_log_level" {
    log_level = %q
}
`, logLevel)
}

func testAccMsserverLoginName(loginName string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_login_name" {
    login_name = %q
}
`, loginName)
}

func testAccMsserverLoginPassword(loginPassword string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_login_password" {
    login_password = %q
}
`, loginPassword)
}

func testAccMsserverMsMaxConnection(msMaxConnection string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_ms_max_connection" {
    ms_max_connection = %q
}
`, msMaxConnection)
}

func testAccMsserverMsRpcTimeoutInSeconds(msRpcTimeoutInSeconds string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_ms_rpc_timeout_in_seconds" {
    ms_rpc_timeout_in_seconds = %q
}
`, msRpcTimeoutInSeconds)
}

func testAccMsserverNetworkView(networkView string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_network_view" {
    network_view = %q
}
`, networkView)
}

func testAccMsserverReadOnly(readOnly string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_read_only" {
    read_only = %q
}
`, readOnly)
}

func testAccMsserverSynchronizationMinDelay(synchronizationMinDelay string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_synchronization_min_delay" {
    synchronization_min_delay = %q
}
`, synchronizationMinDelay)
}

func testAccMsserverUseLogDestination(useLogDestination string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_use_log_destination" {
    use_log_destination = %q
}
`, useLogDestination)
}

func testAccMsserverUseMsMaxConnection(useMsMaxConnection string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_use_ms_max_connection" {
    use_ms_max_connection = %q
}
`, useMsMaxConnection)
}

func testAccMsserverUseMsRpcTimeoutInSeconds(useMsRpcTimeoutInSeconds string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_use_ms_rpc_timeout_in_seconds" {
    use_ms_rpc_timeout_in_seconds = %q
}
`, useMsRpcTimeoutInSeconds)
}
