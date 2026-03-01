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
	syncMinDelay := "2"

	adsitesLoginName := acctest.RandomName()
	adsitesSyncMinDelay := "5"
	adsites := map[string]any{
		"login_name":                    adsitesLoginName,
		"use_login":                     true,
		"synchronization_min_delay":     adsitesSyncMinDelay,
		"use_synchronization_min_delay": true,
	}

	updatedAdsitesLoginName := acctest.RandomName()
	updatedAdSyncMinDelay := "10"
	updatedAdsites := map[string]any{
		"login_name":                    updatedAdsitesLoginName,
		"use_login":                     true,
		"synchronization_min_delay":     updatedAdSyncMinDelay,
		"use_synchronization_min_delay": true,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverAdSites(address, loginName, syncMinDelay, adsites),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ad_sites.login_name", adsitesLoginName),
					resource.TestCheckResourceAttr(resourceName, "ad_sites.synchronization_min_delay", adsitesSyncMinDelay),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverAdSites(address, loginName, updatedAdSyncMinDelay, updatedAdsites),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ad_sites.login_name", updatedAdsitesLoginName),
					resource.TestCheckResourceAttr(resourceName, "ad_sites.synchronization_min_delay", updatedAdSyncMinDelay),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_AdUser(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_ad_user"
	var v microsoft.Msserver

	address := "10.10.0.1"
	loginName := acctest.RandomName()
	syncMinDelay := "2"

	aduserLoginName := acctest.RandomName()
	aduserSyncMinDelay := "5"
	aduser := map[string]any{
		"login_name":                    aduserLoginName,
		"use_login":                     true,
		"synchronization_interval":      aduserSyncMinDelay,
		"use_synchronization_interval":  true,
		"use_synchronization_min_delay": true,
	}

	updatedAduserLoginName := acctest.RandomName()
	updatedAduserSyncMinDelay := "10"
	updatedAduser := map[string]any{
		"login_name":                    updatedAduserLoginName,
		"use_login":                     true,
		"synchronization_interval":      updatedAduserSyncMinDelay,
		"use_synchronization_interval":  true,
		"use_synchronization_min_delay": true,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverAdUser(address, loginName, syncMinDelay, aduser),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ad_user.login_name", aduserLoginName),
					resource.TestCheckResourceAttr(resourceName, "ad_user.synchronization_interval", aduserSyncMinDelay),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverAdUser(address, loginName, syncMinDelay, updatedAduser),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ad_user.login_name", updatedAduserLoginName),
					resource.TestCheckResourceAttr(resourceName, "ad_user.synchronization_interval", updatedAduserSyncMinDelay),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_Address(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_address"
	var v microsoft.Msserver

	address1 := "10.10.0.1"
	address2 := "10.10.0.2"
	loginName := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverAddress(address1, loginName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", address1),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverAddress(address2, loginName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", address2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_Comment(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_comment"
	var v microsoft.Msserver
	address := "10.10.0.1"
	loginName := acctest.RandomName()
	comment1 := "This is a new msserver"
	comment2 := "This is an updated msserver"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverComment(address, loginName, comment1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment1),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverComment(address, loginName, comment2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_DhcpServer(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_dhcp_server"
	var v microsoft.Msserver

	address := "10.10.0.1"
	loginName := acctest.RandomName()
	syncMinDelay := "2"

	dhcpServerLoginName := acctest.RandomName()
	dhcpServerSyncMinDelay := "5"
	dhcp := map[string]any{
		"login_name":                    dhcpServerLoginName,
		"use_login":                     true,
		"synchronization_min_delay":     dhcpServerSyncMinDelay,
		"use_synchronization_min_delay": true,
	}

	updatedDhcpServerLoginName := acctest.RandomName()
	updatedDhcpServerSyncMinDelay := "10"
	updatedDhcp := map[string]any{
		"login_name":                    updatedDhcpServerLoginName,
		"use_login":                     true,
		"synchronization_min_delay":     updatedDhcpServerSyncMinDelay,
		"use_synchronization_min_delay": true,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDhcpServer(address, loginName, syncMinDelay, dhcp),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_server.login_name", dhcpServerLoginName),
					resource.TestCheckResourceAttr(resourceName, "dhcp_server.synchronization_min_delay", dhcpServerSyncMinDelay),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDhcpServer(address, loginName, syncMinDelay, updatedDhcp),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_server.login_name", updatedDhcpServerLoginName),
					resource.TestCheckResourceAttr(resourceName, "dhcp_server.synchronization_min_delay", updatedDhcpServerSyncMinDelay),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_Disabled(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_disabled"
	var v microsoft.Msserver

	address := "10.10.0.1"
	loginName := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDisabled(address, loginName, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDisabled(address, loginName, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_DnsServer(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_dns_server"
	var v microsoft.Msserver

	address := "10.10.0.1"
	loginName := acctest.RandomName()
	syncMinDelay := "2"

	dnsServerLoginName := acctest.RandomName()
	dnsServerSyncMinDelay := "5"
	dns := map[string]any{
		"login_name":                    dnsServerLoginName,
		"use_login":                     true,
		"synchronization_min_delay":     dnsServerSyncMinDelay,
		"use_synchronization_min_delay": true,
	}

	updatedDnsServerLoginName := acctest.RandomName()
	updatedDnsServerSyncMinDelay := "10"
	updatedDns := map[string]any{
		"login_name":                    updatedDnsServerLoginName,
		"use_login":                     true,
		"synchronization_min_delay":     updatedDnsServerSyncMinDelay,
		"use_synchronization_min_delay": true,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDnsServer(address, loginName, syncMinDelay, dns),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_server.login_name", dnsServerLoginName),
					resource.TestCheckResourceAttr(resourceName, "dns_server.synchronization_min_delay", dnsServerSyncMinDelay),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverDnsServer(address, loginName, syncMinDelay, updatedDns),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_server.login_name", updatedDnsServerLoginName),
					resource.TestCheckResourceAttr(resourceName, "dns_server.synchronization_min_delay", updatedDnsServerSyncMinDelay),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_DnsView(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_dns_view"
	var v microsoft.Msserver

	address := "10.10.0.1"
	loginName := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverDnsView(address, loginName, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_view", "default"),
				),
			},
			// Can't update view as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_extattrs"
	var v microsoft.Msserver

	address := "10.10.0.1"
	loginName := acctest.RandomName()
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverExtAttrs(address, loginName, map[string]any{"Site": extAttrValue1}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverExtAttrs(address, loginName, map[string]any{"Site": extAttrValue2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_GridMember(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_grid_member"
	var v microsoft.Msserver

	address := "10.10.0.1"
	loginName := acctest.RandomName()
	member1 := "infoblox.localdomain"
	member2 := "member.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverGridMember(address, loginName, member1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_member", member1),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverGridMember(address, loginName, member2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_member", member2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_LogDestination(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_log_destination"
	var v microsoft.Msserver

	address := "10.10.0.1"
	loginName := acctest.RandomName()
	member1 := "MSLOG"
	member2 := "SYSLOG"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverLogDestination(address, loginName, member1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_destination", member1),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverLogDestination(address, loginName, member2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_destination", member2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_LogLevel(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_log_level"
	var v microsoft.Msserver

	address := "10.10.0.1"
	loginName := acctest.RandomName()
	member1 := "ADVANCED"
	member2 := "FULL"
	member3 := "MINIMUM"
	member4 := "NORMAL"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverLogLevel(address, loginName, member1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", member1),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverLogLevel(address, loginName, member2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", member2),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverLogLevel(address, loginName, member3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", member3),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverLogLevel(address, loginName, member4),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", member4),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_LoginName(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_login_name"
	var v microsoft.Msserver

	address := "10.10.0.1"
	loginName1 := acctest.RandomName()
	loginName2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverLoginName(address, loginName1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "login_name", loginName1),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverLoginName(address, loginName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "login_name", loginName2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_LoginPassword(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_login_password"
	var v microsoft.Msserver
	address := "10.10.0.1"
	loginName := acctest.RandomName()
	loginPassword1 := "Password@123"
	loginPassword2 := "Password@456"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverLoginPassword(address, loginName, loginPassword1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "login_password", loginPassword1),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverLoginPassword(address, loginName, loginPassword2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "login_password", loginPassword2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_MsMaxConnection(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_ms_max_connection"
	var v microsoft.Msserver

	address := "10.10.0.1"
	loginName := acctest.RandomName()
	maxConnection := "10"
	updatedMaxConnection := "20"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverMsMaxConnection(address, loginName, maxConnection),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_max_connection", maxConnection),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverMsMaxConnection(address, loginName, updatedMaxConnection),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_max_connection", updatedMaxConnection),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_MsRpcTimeoutInSeconds(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_ms_rpc_timeout_in_seconds"
	var v microsoft.Msserver

	address := "10.10.0.1"
	loginName := acctest.RandomName()
	msRpcTimeoutInSeconds := "30"
	updatedMsRpcTimeoutInSeconds := "60"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverMsRpcTimeoutInSeconds(address, loginName, msRpcTimeoutInSeconds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_rpc_timeout_in_seconds", msRpcTimeoutInSeconds),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverMsRpcTimeoutInSeconds(address, loginName, updatedMsRpcTimeoutInSeconds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_rpc_timeout_in_seconds", updatedMsRpcTimeoutInSeconds),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_NetworkView(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_network_view"
	var v microsoft.Msserver

	address := "10.10.0.1"
	loginName := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverNetworkView(address, loginName, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "default"),
				),
			},
			// Can't update view as it is immutable

			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_ReadOnly(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_read_only"
	var v microsoft.Msserver

	address := "10.10.0.1"
	loginName := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverReadOnly(address, loginName, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "read_only", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverReadOnly(address, loginName, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "read_only", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_SynchronizationMinDelay(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_synchronization_min_delay"
	var v microsoft.Msserver

	address := "10.10.0.1"
	loginName := acctest.RandomName()
	synchronizationMinDelay := "2"
	updatedSynchronizationMinDelay := "5"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverSynchronizationMinDelay(address, loginName, synchronizationMinDelay),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "synchronization_min_delay", synchronizationMinDelay),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverSynchronizationMinDelay(address, loginName, updatedSynchronizationMinDelay),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "synchronization_min_delay", updatedSynchronizationMinDelay),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_UseLogDestination(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_use_log_destination"
	var v microsoft.Msserver

	address := "10.10.0.1"
	loginName := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverUseLogDestination(address, loginName, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_log_destination", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverUseLogDestination(address, loginName, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_log_destination", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_UseMsMaxConnection(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_use_ms_max_connection"
	var v microsoft.Msserver

	address := "10.10.0.1"
	loginName := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverUseMsMaxConnection(address, loginName, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ms_max_connection", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverUseMsMaxConnection(address, loginName, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ms_max_connection", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMsserverResource_UseMsRpcTimeoutInSeconds(t *testing.T) {
	var resourceName = "nios_microsoft_msserver.test_use_ms_rpc_timeout_in_seconds"
	var v microsoft.Msserver

	address := "10.10.0.1"
	loginName := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMsserverUseMsRpcTimeoutInSeconds(address, loginName, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ms_rpc_timeout_in_seconds", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccMsserverUseMsRpcTimeoutInSeconds(address, loginName, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMsserverExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ms_rpc_timeout_in_seconds", "true"),
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

func testAccMsserverAdSites(address, loginName, syncMinDelay string, adSites map[string]any) string {
	adSitesHCL := utils.ConvertMapToHCL(adSites)

	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_ad_sites" {
    address = %q
	login_name = %q
	synchronization_min_delay = %q
	ad_sites = %s
}
`, address, loginName, syncMinDelay, adSitesHCL)
}

func testAccMsserverAdUser(address, loginName, syncMinDelay string, adUser map[string]any) string {
	adUserHCL := utils.ConvertMapToHCL(adUser)

	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_ad_user" {
    address = %q
	login_name = %q
	synchronization_min_delay = %q
	ad_user = %s
}
`, address, loginName, syncMinDelay, adUserHCL)
}

func testAccMsserverAddress(address, loginName string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_address" {
    address = %q
	login_name = %q
}
`, address, loginName)
}

func testAccMsserverComment(address, loginName, comment string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_comment" {
    address = %q
	login_name = %q
    comment = %q
}
`, address, loginName, comment)
}

func testAccMsserverDhcpServer(address, loginName, syncMinDelay string, dhcpServer map[string]any) string {
	dhcpServerHCL := utils.ConvertMapToHCL(dhcpServer)

	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_dhcp_server" {
    address = %q
	login_name = %q
	synchronization_min_delay = %q
	dhcp_server = %s
}
`, address, loginName, syncMinDelay, dhcpServerHCL)
}

func testAccMsserverDisabled(address, loginName, disabled string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_disabled" {
    address = %q
	login_name = %q
    disabled = %q
}
`, address, loginName, disabled)
}

func testAccMsserverDnsServer(address, loginName, syncMinDelay string, dnsServer map[string]any) string {
	dnsServerHCL := utils.ConvertMapToHCL(dnsServer)

	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_dns_server" {
    address = %q
	login_name = %q
    synchronization_min_delay = %q
    dns_server = %s
}
`, address, loginName, syncMinDelay, dnsServerHCL)
}

func testAccMsserverDnsView(address, loginName, dnsView string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_dns_view" {
    address = %q
	login_name = %q
    dns_view = %q
}
`, address, loginName, dnsView)
}

func testAccMsserverExtAttrs(address, loginName string, extAttrs map[string]any) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
	%s = %q
	`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_extattrs" {
 	address = %q
	login_name = %q
    extattrs = %s
}
`, address, loginName, extattrsStr)
}

func testAccMsserverGridMember(address, loginName, gridMember string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_grid_member" {
    address = %q
	login_name = %q
    grid_member = %q
}
`, address, loginName, gridMember)
}

func testAccMsserverLogDestination(address, loginName, logDestination string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_log_destination" {
    address = %q
	login_name = %q
    log_destination = %q
}
`, address, loginName, logDestination)
}

func testAccMsserverLogLevel(address, loginName, logLevel string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_log_level" {
    address = %q
	login_name = %q
    log_level = %q
}
`, address, loginName, logLevel)
}

func testAccMsserverLoginName(address, loginName string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_login_name" {
    address = %q
	login_name = %q
}
`, address, loginName)
}

func testAccMsserverLoginPassword(address, loginName, loginPassword string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_login_password" {
    address = %q
	login_name = %q
    login_password = %q
}
`, address, loginName, loginPassword)
}

func testAccMsserverMsMaxConnection(address, loginName, msMaxConnection string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_ms_max_connection" {
    address = %q
	login_name = %q
    ms_max_connection = %q
}
`, address, loginName, msMaxConnection)
}

func testAccMsserverMsRpcTimeoutInSeconds(address, loginName, msRpcTimeoutInSeconds string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_ms_rpc_timeout_in_seconds" {
	address = %q
	login_name = %q
    ms_rpc_timeout_in_seconds = %q
}
`, address, loginName, msRpcTimeoutInSeconds)
}

func testAccMsserverNetworkView(address, loginName, networkView string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_network_view" {
	address = %q
	login_name = %q
    network_view = %q
}
`, address, loginName, networkView)
}

func testAccMsserverReadOnly(address, loginName string, readOnly bool) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_read_only" {
    address = %q
	login_name = %q
    read_only = %t
}
`, address, loginName, readOnly)
}

func testAccMsserverSynchronizationMinDelay(address, loginName, synchronizationMinDelay string) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_synchronization_min_delay" {
    address = %q
	login_name = %q
    synchronization_min_delay = %q
}
`, address, loginName, synchronizationMinDelay)
}

func testAccMsserverUseLogDestination(address, loginName string, useLogDestination bool) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_use_log_destination" {
    address = %q
	login_name = %q
    use_log_destination = %t
}
`, address, loginName, useLogDestination)
}

func testAccMsserverUseMsMaxConnection(address, loginName string, useMsMaxConnection bool) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_use_ms_max_connection" {
    address = %q
	login_name = %q
    use_ms_max_connection = %t
}
`, address, loginName, useMsMaxConnection)
}

func testAccMsserverUseMsRpcTimeoutInSeconds(address, loginName string, useMsRpcTimeoutInSeconds bool) string {
	return fmt.Sprintf(`
resource "nios_microsoft_msserver" "test_use_ms_rpc_timeout_in_seconds" {
	address = %q
	login_name = %q
    use_ms_rpc_timeout_in_seconds = %t
}
`, address, loginName, useMsRpcTimeoutInSeconds)
}
