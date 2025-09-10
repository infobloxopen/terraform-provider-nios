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

var readableAttributesForAdmingroup = "access_method,admin_set_commands,admin_show_commands,admin_toplevel_commands,cloud_set_commands,cloud_show_commands,comment,database_set_commands,database_show_commands,dhcp_set_commands,dhcp_show_commands,disable,disable_concurrent_login,dns_set_commands,dns_show_commands,dns_toplevel_commands,docker_set_commands,docker_show_commands,email_addresses,enable_restricted_user_access,extattrs,grid_set_commands,grid_show_commands,inactivity_lockout_setting,licensing_set_commands,licensing_show_commands,lockout_setting,machine_control_toplevel_commands,name,networking_set_commands,networking_show_commands,password_setting,roles,saml_setting,security_set_commands,security_show_commands,superuser,trouble_shooting_toplevel_commands,use_account_inactivity_lockout_enable,use_disable_concurrent_login,use_lockout_setting,use_password_setting,user_access"

func TestAccAdmingroupResource_basic(t *testing.T) {
	var resourceName = "nios_security_admin_group.test"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "access_method.#", "4"),
					resource.TestCheckResourceAttr(resourceName, "access_method.0", "GUI"),
					resource.TestCheckResourceAttr(resourceName, "access_method.1", "API"),
					resource.TestCheckResourceAttr(resourceName, "access_method.2", "TAXII"),
					resource.TestCheckResourceAttr(resourceName, "access_method.3", "CLI"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "disable_concurrent_login", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_restricted_user_access", "false"),
					// Test inactivity_lockout_setting default values
					resource.TestCheckResourceAttr(resourceName, "inactivity_lockout_setting.account_inactivity_lockout_enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "inactivity_lockout_setting.inactive_days", "30"),
					resource.TestCheckResourceAttr(resourceName, "inactivity_lockout_setting.reactivate_via_remote_console_enable", "true"),
					resource.TestCheckResourceAttr(resourceName, "inactivity_lockout_setting.reactivate_via_serial_console_enable", "true"),
					resource.TestCheckResourceAttr(resourceName, "inactivity_lockout_setting.reminder_days", "15"),
					// Test lockout_setting default values
					resource.TestCheckResourceAttr(resourceName, "lockout_setting.enable_sequential_failed_login_attempts_lockout", "false"),
					resource.TestCheckResourceAttr(resourceName, "lockout_setting.failed_lockout_duration", "5"),
					resource.TestCheckResourceAttr(resourceName, "lockout_setting.never_unlock_user", "false"),
					resource.TestCheckResourceAttr(resourceName, "lockout_setting.sequential_attempts", "5"),
					// Test password_setting default values
					//resource.TestCheckResourceAttr(resourceName, "password_setting.expire_days", "30"),
					//resource.TestCheckResourceAttr(resourceName, "password_setting.expire_enable", "false"),
					//resource.TestCheckResourceAttr(resourceName, "password_setting.reminder_days", "15"),

					resource.TestCheckResourceAttr(resourceName, "superuser", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_account_inactivity_lockout_enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_disable_concurrent_login", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_lockout_setting", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_password_setting", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_disappears(t *testing.T) {
	resourceName := "nios_security_admin_group.test"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAdmingroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAdmingroupBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					testAccCheckAdmingroupDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAdmingroupResource_AccessMethod(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_access_method"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	accessMethod := `["CLI", "GUI"]`
	accessMethod1 := `["CLI","GUI","API"]`
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupAccessMethod(name, accessMethod),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "access_method.0", "CLI"),
					resource.TestCheckResourceAttr(resourceName, "access_method.1", "GUI"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupAccessMethod(name, accessMethod1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "access_method.0", "CLI"),
					resource.TestCheckResourceAttr(resourceName, "access_method.1", "GUI"),
					resource.TestCheckResourceAttr(resourceName, "access_method.2", "API"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_AdminSetCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_admin_set_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	adminSetCmds := "set_collect_old_logs:true,set_bfd:true,set_bgp:true"
	adminSetCmds1 := "set_collect_old_logs:false,set_bfd:true,set_bgp:false,set_core_files_quota:true"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupAdminSetCommands(name, adminSetCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "admin_set_commands.set_collect_old_logs", "true"),
					resource.TestCheckResourceAttr(resourceName, "admin_set_commands.set_bfd", "true"),
					resource.TestCheckResourceAttr(resourceName, "admin_set_commands.set_bgp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupAdminSetCommands(name, adminSetCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "admin_set_commands.set_collect_old_logs", "false"),
					resource.TestCheckResourceAttr(resourceName, "admin_set_commands.set_bfd", "true"),
					resource.TestCheckResourceAttr(resourceName, "admin_set_commands.set_bgp", "false"),
					resource.TestCheckResourceAttr(resourceName, "admin_set_commands.set_core_files_quota", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_AdminShowCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_admin_show_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	adminShowCmds := "show_arp:true,show_bfd:true,show_bgp:true"
	adminShowCmds1 := "show_arp:false,show_bfd:true,show_bgp:false,show_capacity:true"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupAdminShowCommands(name, adminShowCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "admin_show_commands.show_arp", "true"),
					resource.TestCheckResourceAttr(resourceName, "admin_show_commands.show_bfd", "true"),
					resource.TestCheckResourceAttr(resourceName, "admin_show_commands.show_bgp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupAdminShowCommands(name, adminShowCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "admin_show_commands.show_arp", "false"),
					resource.TestCheckResourceAttr(resourceName, "admin_show_commands.show_bfd", "true"),
					resource.TestCheckResourceAttr(resourceName, "admin_show_commands.show_bgp", "false"),
					resource.TestCheckResourceAttr(resourceName, "admin_show_commands.show_capacity", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_AdminToplevelCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_admin_toplevel_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	adminTopLevelCmds := "iostat:true,netstat:true,ps:true"
	adminTopLevelCmds1 := "iostat:false,netstat:true,ps:false,restart_product:true"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupAdminToplevelCommands(name, adminTopLevelCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "admin_toplevel_commands.iostat", "true"),
					resource.TestCheckResourceAttr(resourceName, "admin_toplevel_commands.netstat", "true"),
					resource.TestCheckResourceAttr(resourceName, "admin_toplevel_commands.ps", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupAdminToplevelCommands(name, adminTopLevelCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "admin_toplevel_commands.iostat", "false"),
					resource.TestCheckResourceAttr(resourceName, "admin_toplevel_commands.netstat", "true"),
					resource.TestCheckResourceAttr(resourceName, "admin_toplevel_commands.ps", "false"),
					resource.TestCheckResourceAttr(resourceName, "admin_toplevel_commands.restart_product", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_CloudSetCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_cloud_set_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	cloudSetCmds := "disable_all:true,set_cloud_services_portal:true"
	cloudSetCmds1 := "enable_all:true,set_cloud_services_portal:false"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupCloudSetCommands(name, cloudSetCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_set_commands.disable_all", "true"),
					resource.TestCheckResourceAttr(resourceName, "cloud_set_commands.set_cloud_services_portal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupCloudSetCommands(name, cloudSetCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_set_commands.enable_all", "true"),
					resource.TestCheckResourceAttr(resourceName, "cloud_set_commands.set_cloud_services_portal", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_CloudShowCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_cloud_show_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupCloudShowCommands("CLOUD_SHOW_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_show_commands", "CLOUD_SHOW_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupCloudShowCommands("CLOUD_SHOW_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_show_commands", "CLOUD_SHOW_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_Comment(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_comment"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_DatabaseSetCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_database_set_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDatabaseSetCommands("DATABASE_SET_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "database_set_commands", "DATABASE_SET_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDatabaseSetCommands("DATABASE_SET_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "database_set_commands", "DATABASE_SET_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_DatabaseShowCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_database_show_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDatabaseShowCommands("DATABASE_SHOW_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "database_show_commands", "DATABASE_SHOW_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDatabaseShowCommands("DATABASE_SHOW_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "database_show_commands", "DATABASE_SHOW_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_DhcpSetCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_dhcp_set_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDhcpSetCommands("DHCP_SET_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_set_commands", "DHCP_SET_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDhcpSetCommands("DHCP_SET_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_set_commands", "DHCP_SET_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_DhcpShowCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_dhcp_show_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDhcpShowCommands("DHCP_SHOW_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_show_commands", "DHCP_SHOW_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDhcpShowCommands("DHCP_SHOW_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_show_commands", "DHCP_SHOW_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_Disable(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_disable"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_DisableConcurrentLogin(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_disable_concurrent_login"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDisableConcurrentLogin("DISABLE_CONCURRENT_LOGIN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_concurrent_login", "DISABLE_CONCURRENT_LOGIN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDisableConcurrentLogin("DISABLE_CONCURRENT_LOGIN_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_concurrent_login", "DISABLE_CONCURRENT_LOGIN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_DnsSetCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_dns_set_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDnsSetCommands("DNS_SET_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_set_commands", "DNS_SET_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDnsSetCommands("DNS_SET_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_set_commands", "DNS_SET_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_DnsShowCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_dns_show_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDnsShowCommands("DNS_SHOW_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_show_commands", "DNS_SHOW_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDnsShowCommands("DNS_SHOW_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_show_commands", "DNS_SHOW_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_DnsToplevelCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_dns_toplevel_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDnsToplevelCommands("DNS_TOPLEVEL_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_toplevel_commands", "DNS_TOPLEVEL_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDnsToplevelCommands("DNS_TOPLEVEL_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_toplevel_commands", "DNS_TOPLEVEL_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_DockerSetCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_docker_set_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDockerSetCommands("DOCKER_SET_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "docker_set_commands", "DOCKER_SET_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDockerSetCommands("DOCKER_SET_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "docker_set_commands", "DOCKER_SET_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_DockerShowCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_docker_show_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDockerShowCommands("DOCKER_SHOW_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "docker_show_commands", "DOCKER_SHOW_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDockerShowCommands("DOCKER_SHOW_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "docker_show_commands", "DOCKER_SHOW_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_EmailAddresses(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_email_addresses"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupEmailAddresses("EMAIL_ADDRESSES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_addresses", "EMAIL_ADDRESSES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupEmailAddresses("EMAIL_ADDRESSES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_addresses", "EMAIL_ADDRESSES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_EnableRestrictedUserAccess(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_enable_restricted_user_access"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupEnableRestrictedUserAccess("ENABLE_RESTRICTED_USER_ACCESS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_restricted_user_access", "ENABLE_RESTRICTED_USER_ACCESS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupEnableRestrictedUserAccess("ENABLE_RESTRICTED_USER_ACCESS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_restricted_user_access", "ENABLE_RESTRICTED_USER_ACCESS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_extattrs"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_GridSetCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_grid_set_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupGridSetCommands("GRID_SET_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_set_commands", "GRID_SET_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupGridSetCommands("GRID_SET_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_set_commands", "GRID_SET_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_GridShowCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_grid_show_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupGridShowCommands("GRID_SHOW_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_show_commands", "GRID_SHOW_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupGridShowCommands("GRID_SHOW_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_show_commands", "GRID_SHOW_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_InactivityLockoutSetting(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_inactivity_lockout_setting"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupInactivityLockoutSetting("INACTIVITY_LOCKOUT_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inactivity_lockout_setting", "INACTIVITY_LOCKOUT_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupInactivityLockoutSetting("INACTIVITY_LOCKOUT_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inactivity_lockout_setting", "INACTIVITY_LOCKOUT_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_LicensingSetCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_licensing_set_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupLicensingSetCommands("LICENSING_SET_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "licensing_set_commands", "LICENSING_SET_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupLicensingSetCommands("LICENSING_SET_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "licensing_set_commands", "LICENSING_SET_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_LicensingShowCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_licensing_show_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupLicensingShowCommands("LICENSING_SHOW_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "licensing_show_commands", "LICENSING_SHOW_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupLicensingShowCommands("LICENSING_SHOW_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "licensing_show_commands", "LICENSING_SHOW_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_LockoutSetting(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_lockout_setting"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupLockoutSetting("LOCKOUT_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lockout_setting", "LOCKOUT_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupLockoutSetting("LOCKOUT_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lockout_setting", "LOCKOUT_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_MachineControlToplevelCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_machine_control_toplevel_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupMachineControlToplevelCommands("MACHINE_CONTROL_TOPLEVEL_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "machine_control_toplevel_commands", "MACHINE_CONTROL_TOPLEVEL_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupMachineControlToplevelCommands("MACHINE_CONTROL_TOPLEVEL_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "machine_control_toplevel_commands", "MACHINE_CONTROL_TOPLEVEL_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_Name(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_name"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_NetworkingSetCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_networking_set_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupNetworkingSetCommands("NETWORKING_SET_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "networking_set_commands", "NETWORKING_SET_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupNetworkingSetCommands("NETWORKING_SET_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "networking_set_commands", "NETWORKING_SET_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_NetworkingShowCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_networking_show_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupNetworkingShowCommands("NETWORKING_SHOW_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "networking_show_commands", "NETWORKING_SHOW_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupNetworkingShowCommands("NETWORKING_SHOW_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "networking_show_commands", "NETWORKING_SHOW_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_PasswordSetting(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_password_setting"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupPasswordSetting("PASSWORD_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "password_setting", "PASSWORD_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupPasswordSetting("PASSWORD_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "password_setting", "PASSWORD_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_Roles(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_roles"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupRoles("ROLES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "roles", "ROLES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupRoles("ROLES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "roles", "ROLES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_SamlSetting(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_saml_setting"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupSamlSetting("SAML_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "saml_setting", "SAML_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupSamlSetting("SAML_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "saml_setting", "SAML_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_SecuritySetCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_security_set_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupSecuritySetCommands("SECURITY_SET_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "security_set_commands", "SECURITY_SET_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupSecuritySetCommands("SECURITY_SET_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "security_set_commands", "SECURITY_SET_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_SecurityShowCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_security_show_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupSecurityShowCommands("SECURITY_SHOW_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "security_show_commands", "SECURITY_SHOW_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupSecurityShowCommands("SECURITY_SHOW_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "security_show_commands", "SECURITY_SHOW_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_Superuser(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_superuser"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupSuperuser("SUPERUSER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "superuser", "SUPERUSER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupSuperuser("SUPERUSER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "superuser", "SUPERUSER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_TroubleShootingToplevelCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_trouble_shooting_toplevel_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupTroubleShootingToplevelCommands("TROUBLE_SHOOTING_TOPLEVEL_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "trouble_shooting_toplevel_commands", "TROUBLE_SHOOTING_TOPLEVEL_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupTroubleShootingToplevelCommands("TROUBLE_SHOOTING_TOPLEVEL_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "trouble_shooting_toplevel_commands", "TROUBLE_SHOOTING_TOPLEVEL_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_UseAccountInactivityLockoutEnable(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_use_account_inactivity_lockout_enable"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupUseAccountInactivityLockoutEnable("USE_ACCOUNT_INACTIVITY_LOCKOUT_ENABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_account_inactivity_lockout_enable", "USE_ACCOUNT_INACTIVITY_LOCKOUT_ENABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupUseAccountInactivityLockoutEnable("USE_ACCOUNT_INACTIVITY_LOCKOUT_ENABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_account_inactivity_lockout_enable", "USE_ACCOUNT_INACTIVITY_LOCKOUT_ENABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_UseDisableConcurrentLogin(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_use_disable_concurrent_login"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupUseDisableConcurrentLogin("USE_DISABLE_CONCURRENT_LOGIN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_disable_concurrent_login", "USE_DISABLE_CONCURRENT_LOGIN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupUseDisableConcurrentLogin("USE_DISABLE_CONCURRENT_LOGIN_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_disable_concurrent_login", "USE_DISABLE_CONCURRENT_LOGIN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_UseLockoutSetting(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_use_lockout_setting"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupUseLockoutSetting("USE_LOCKOUT_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_lockout_setting", "USE_LOCKOUT_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupUseLockoutSetting("USE_LOCKOUT_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_lockout_setting", "USE_LOCKOUT_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_UsePasswordSetting(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_use_password_setting"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupUsePasswordSetting("USE_PASSWORD_SETTING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_password_setting", "USE_PASSWORD_SETTING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupUsePasswordSetting("USE_PASSWORD_SETTING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_password_setting", "USE_PASSWORD_SETTING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_UserAccess(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_user_access"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupUserAccess("USER_ACCESS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "user_access", "USER_ACCESS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupUserAccess("USER_ACCESS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "user_access", "USER_ACCESS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckAdmingroupExists(ctx context.Context, resourceName string, v *security.Admingroup) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SecurityAPI.
			AdmingroupAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForAdmingroup).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetAdmingroupResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetAdmingroupResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckAdmingroupDestroy(ctx context.Context, v *security.Admingroup) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SecurityAPI.
			AdmingroupAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForAdmingroup).
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

func testAccCheckAdmingroupDisappears(ctx context.Context, v *security.Admingroup) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SecurityAPI.
			AdmingroupAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccAdmingroupBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test" {
	name = %q
}
`, name)
}

func testAccAdmingroupAccessMethod(name, accessMethod string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_access_method" {
    name = %q
    access_method = %s
}
`, name, accessMethod)
}

func testAccAdmingroupAdminSetCommands(name, adminSetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_admin_set_commands" {
	name = %q
    admin_set_commands = {%s}
}
`, name, adminSetCommands)
}

func testAccAdmingroupAdminShowCommands(name, adminShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_admin_show_commands" {
	name = %q
    admin_show_commands = {%s}
}
`, name, adminShowCommands)
}

func testAccAdmingroupAdminToplevelCommands(name, adminToplevelCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_admin_toplevel_commands" {
	name = %q
    admin_toplevel_commands = {%s}
}
`, name, adminToplevelCommands)
}

func testAccAdmingroupCloudSetCommands(name, cloudSetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_cloud_set_commands" {
	name = %q
    cloud_set_commands = {%s}
}
`, name, cloudSetCommands)
}

func testAccAdmingroupCloudShowCommands(name, cloudShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_cloud_show_commands" {
	name = %q
    cloud_show_commands = {%s}
}
`, name, cloudShowCommands)
}

func testAccAdmingroupComment(name, comment string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccAdmingroupDatabaseSetCommands(name, databaseSetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_database_set_commands" {
    name = %q
    database_set_commands = {%s}
}
`, name, databaseSetCommands)
}

func testAccAdmingroupDatabaseShowCommands(name, databaseShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_database_show_commands" {
    name = %q
    database_show_commands = {%s}
}
`, name, databaseShowCommands)
}

func testAccAdmingroupDhcpSetCommands(name, dhcpSetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_dhcp_set_commands" {
	name = %q
    dhcp_set_commands = {%s}
}
`, name, dhcpSetCommands)
}

func testAccAdmingroupDhcpShowCommands(name, dhcpShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_dhcp_show_commands" {
	name = %q
    dhcp_show_commands = {%s}
}
`, name, dhcpShowCommands)
}

func testAccAdmingroupDisable(name string, disable bool) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_disable" {
	name = %q
    disable = %t
}
`, name, disable)
}

func testAccAdmingroupDisableConcurrentLogin(name string, disableConcurrentLogin bool) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_disable_concurrent_login" {
	name = %q
    disable_concurrent_login = %t
}
`, name, disableConcurrentLogin)
}

func testAccAdmingroupDnsSetCommands(name, dnsSetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_dns_set_commands" {
	name = %q
    dns_set_commands = {%s}
}
`, name, dnsSetCommands)
}

func testAccAdmingroupDnsShowCommands(name, dnsShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_dns_show_commands" {
	name = %q
    dns_show_commands = {%s}
}
`, name, dnsShowCommands)
}

func testAccAdmingroupDnsToplevelCommands(name, dnsToplevelCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_dns_toplevel_commands" {
	name = %q
    dns_toplevel_commands = {%s}
}
`, name, dnsToplevelCommands)
}

func testAccAdmingroupDockerSetCommands(name, dockerSetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_docker_set_commands" {
	name = %q
    docker_set_commands = {%s}
}
`, name, dockerSetCommands)
}

func testAccAdmingroupDockerShowCommands(name, dockerShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_docker_show_commands" {
	name = %q
    docker_show_commands = {%s}
}
`, name, dockerShowCommands)
}

func testAccAdmingroupEmailAddresses(name, emailAddresses string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_email_addresses" {
	name = %q
    email_addresses = %q
}
`, name, emailAddresses)
}

func testAccAdmingroupEnableRestrictedUserAccess(name, enableRestrictedUserAccess bool) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_enable_restricted_user_access" {
	name = %q
    enable_restricted_user_access = %t
}
`, name, enableRestrictedUserAccess)
}

func testAccAdmingroupExtAttrs(name, extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_extattrs" {
	name = %q
    extattrs = %q
}
`, name, extAttrs)
}

func testAccAdmingroupGridSetCommands(name, gridSetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_grid_set_commands" {
	name = %q
    grid_set_commands = {%s}
}
`, name, gridSetCommands)
}

func testAccAdmingroupGridShowCommands(name, gridShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_grid_show_commands" {
    grid_show_commands = %q
}
`, gridShowCommands)
}

func testAccAdmingroupInactivityLockoutSetting(name, inactivityLockoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_inactivity_lockout_setting" {
    inactivity_lockout_setting = %q
}
`, inactivityLockoutSetting)
}

func testAccAdmingroupLicensingSetCommands(name, licensingSetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_licensing_set_commands" {
    licensing_set_commands = %q
}
`, licensingSetCommands)
}

func testAccAdmingroupLicensingShowCommands(name, licensingShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_licensing_show_commands" {
    licensing_show_commands = %q
}
`, licensingShowCommands)
}

func testAccAdmingroupLockoutSetting(name, lockoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_lockout_setting" {
    lockout_setting = %q
}
`, lockoutSetting)
}

func testAccAdmingroupMachineControlToplevelCommands(name, machineControlToplevelCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_machine_control_toplevel_commands" {
    machine_control_toplevel_commands = %q
}
`, machineControlToplevelCommands)
}

func testAccAdmingroupName(name string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_name" {
    name = %q
}
`, name)
}

func testAccAdmingroupNetworkingSetCommands(name, networkingSetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_networking_set_commands" {
    networking_set_commands = %q
}
`, networkingSetCommands)
}

func testAccAdmingroupNetworkingShowCommands(name, networkingShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_networking_show_commands" {
    networking_show_commands = %q
}
`, networkingShowCommands)
}

func testAccAdmingroupPasswordSetting(name, passwordSetting string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_password_setting" {
    password_setting = %q
}
`, passwordSetting)
}

func testAccAdmingroupRoles(name, roles string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_roles" {
    roles = %q
}
`, roles)
}

func testAccAdmingroupSamlSetting(name, samlSetting string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_saml_setting" {
    saml_setting = %q
}
`, samlSetting)
}

func testAccAdmingroupSecuritySetCommands(name, securitySetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_security_set_commands" {
    security_set_commands = %q
}
`, securitySetCommands)
}

func testAccAdmingroupSecurityShowCommands(name, securityShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_security_show_commands" {
    security_show_commands = %q
}
`, securityShowCommands)
}

func testAccAdmingroupSuperuser(name, superuser string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_superuser" {
    superuser = %q
}
`, superuser)
}

func testAccAdmingroupTroubleShootingToplevelCommands(name, troubleShootingToplevelCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_trouble_shooting_toplevel_commands" {
    trouble_shooting_toplevel_commands = %q
}
`, troubleShootingToplevelCommands)
}

func testAccAdmingroupUseAccountInactivityLockoutEnable(name, useAccountInactivityLockoutEnable string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_use_account_inactivity_lockout_enable" {
    use_account_inactivity_lockout_enable = %q
}
`, useAccountInactivityLockoutEnable)
}

func testAccAdmingroupUseDisableConcurrentLogin(name, useDisableConcurrentLogin string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_use_disable_concurrent_login" {
    use_disable_concurrent_login = %q
}
`, useDisableConcurrentLogin)
}

func testAccAdmingroupUseLockoutSetting(name, useLockoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_use_lockout_setting" {
    use_lockout_setting = %q
}
`, useLockoutSetting)
}

func testAccAdmingroupUsePasswordSetting(name, usePasswordSetting string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_use_password_setting" {
    use_password_setting = %q
}
`, usePasswordSetting)
}

func testAccAdmingroupUserAccess(name, userAccess string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_user_access" {
    user_access = %q
}
`, userAccess)
}
