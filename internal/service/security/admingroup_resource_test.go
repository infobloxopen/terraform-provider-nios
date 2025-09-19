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
	cloudSetCmds := "set_cloud_services_portal:true"
	cloudSetCmds1 := "set_cloud_services_portal:false"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupCloudSetCommands(name, cloudSetCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_set_commands.set_cloud_services_portal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupCloudSetCommands(name, cloudSetCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
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
	name := acctest.RandomNameWithPrefix("admin-group")
	cloudShowCmds := "show_cloud_services_portal:true"
	cloudShowCmds1 := "show_cloud_services_portal:false"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupCloudShowCommands(name, cloudShowCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_show_commands.show_cloud_services_portal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupCloudShowCommands(name, cloudShowCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_show_commands.show_cloud_services_portal", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_Comment(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_comment"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupComment(name, "admin group comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "admin group comment"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupComment(name, "admin group comment updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "admin group comment updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_DatabaseSetCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_database_set_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	databaseSetCmds := "set_db_rollover:true,set_database_transfer:true"
	databaseSetCmds1 := "set_db_rollover:true,set_database_transfer:false"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDatabaseSetCommands(name, databaseSetCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "database_set_commands.set_db_rollover", "true"),
					resource.TestCheckResourceAttr(resourceName, "database_set_commands.set_database_transfer", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDatabaseSetCommands(name, databaseSetCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "database_set_commands.set_db_rollover", "true"),
					resource.TestCheckResourceAttr(resourceName, "database_set_commands.set_database_transfer", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_DatabaseShowCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_database_show_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	databaseShowCmds := "show_backup:true,show_dbsize:true"
	databaseShowCmds1 := "show_backup:false,show_dbsize:false"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDatabaseShowCommands(name, databaseShowCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "database_show_commands.show_backup", "true"),
					resource.TestCheckResourceAttr(resourceName, "database_show_commands.show_dbsize", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDatabaseShowCommands(name, databaseShowCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "database_show_commands.show_backup", "false"),
					resource.TestCheckResourceAttr(resourceName, "database_show_commands.show_dbsize", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_DhcpSetCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_dhcp_set_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	dhcpSetCmds := "set_log_txn_id:true,set_overload_bootp:true"
	dhcpSetCmds1 := "set_log_txn_id:false,set_overload_bootp:false"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDhcpSetCommands(name, dhcpSetCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_set_commands.set_log_txn_id", "true"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_set_commands.set_overload_bootp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDhcpSetCommands(name, dhcpSetCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_set_commands.set_log_txn_id", "false"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_set_commands.set_overload_bootp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_DhcpShowCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_dhcp_show_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	dhcpShowCmds := "show_overload_bootp:true,show_log_txn_id:true"
	dhcpShowCmds1 := "show_overload_bootp:false,show_log_txn_id:false"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDhcpShowCommands(name, dhcpShowCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_show_commands.show_overload_bootp", "true"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_show_commands.show_log_txn_id", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDhcpShowCommands(name, dhcpShowCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_show_commands.show_overload_bootp", "false"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_show_commands.show_log_txn_id", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_Disable(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_disable"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDisable(name, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDisable(name, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_DisableConcurrentLogin(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_disable_concurrent_login"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDisableConcurrentLogin(name, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_concurrent_login", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDisableConcurrentLogin(name, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_concurrent_login", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_DnsSetCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_dns_set_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	dnsSetCmds := "set_allow_query_domain:true,set_dns_accel_debug:true"
	dnsSetCmds1 := "set_allow_query_domain:false,set_dns_accel_debug:false"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDnsSetCommands(name, dnsSetCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_set_commands.set_allow_query_domain", "true"),
					resource.TestCheckResourceAttr(resourceName, "dns_set_commands.set_dns_accel_debug", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDnsSetCommands(name, dnsSetCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_set_commands.set_allow_query_domain", "false"),
					resource.TestCheckResourceAttr(resourceName, "dns_set_commands.set_dns_accel_debug", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_DnsShowCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_dns_show_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	dndShowCmds := `show_allow_query_domain:true,show_dns:true`
	dnsShowCmds1 := `show_allow_query_domain:false,show_dns:false`
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDnsShowCommands(name, dndShowCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_show_commands.show_allow_query_domain", "true"),
					resource.TestCheckResourceAttr(resourceName, "dns_show_commands.show_dns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDnsShowCommands(name, dnsShowCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_show_commands.show_allow_query_domain", "false"),
					resource.TestCheckResourceAttr(resourceName, "dns_show_commands.show_dns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_DnsToplevelCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_dns_toplevel_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	dnsTopLevelCmds := `ddns_add:true,ddns_delete:true`
	dnsTopLevelCmds1 := `ddns_add:false,ddns_delete:false`
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDnsToplevelCommands(name, dnsTopLevelCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_toplevel_commands.ddns_add", "true"),
					resource.TestCheckResourceAttr(resourceName, "dns_toplevel_commands.ddns_delete", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDnsToplevelCommands(name, dnsTopLevelCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_toplevel_commands.ddns_add", "false"),
					resource.TestCheckResourceAttr(resourceName, "dns_toplevel_commands.ddns_delete", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_DockerSetCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_docker_set_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	dockerSetCmds := "set_docker_bridge:true"
	dockerSetCmds1 := "set_docker_bridge:false"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDockerSetCommands(name, dockerSetCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "docker_set_commands.set_docker_bridge", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDockerSetCommands(name, dockerSetCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "docker_set_commands.set_docker_bridge", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_DockerShowCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_docker_show_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	dockerShowCmds := "show_docker_bridge:true"
	dockerShowCmds1 := "show_docker_bridge:false"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupDockerShowCommands(name, dockerShowCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "docker_show_commands.show_docker_bridge", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupDockerShowCommands(name, dockerShowCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "docker_show_commands.show_docker_bridge", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_EmailAddresses(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_email_addresses"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	emailAddresses := `["abc@info.com","xyz@example.com"]`
	emailAddresses1 := `["abc@info1.com","xyz@example1.com"]`
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupEmailAddresses(name, emailAddresses),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_addresses.0", "abc@info.com"),
					resource.TestCheckResourceAttr(resourceName, "email_addresses.1", "xyz@example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupEmailAddresses(name, emailAddresses1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_addresses.0", "abc@info1.com"),
					resource.TestCheckResourceAttr(resourceName, "email_addresses.1", "xyz@example1.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_EnableRestrictedUserAccess(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_enable_restricted_user_access"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupEnableRestrictedUserAccess(name, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_restricted_user_access", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupEnableRestrictedUserAccess(name, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_restricted_user_access", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_extattrs"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupExtAttrs(name, map[string]string{"Site": extAttrValue1}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupExtAttrs(name, map[string]string{"Site": extAttrValue2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_GridSetCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_grid_set_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	gridSetCmds := "set_membership:true,set_dscp:true"
	gridSetCmds1 := "set_membership:false,set_dscp:false"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupGridSetCommands(name, gridSetCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_set_commands.set_membership", "true"),
					resource.TestCheckResourceAttr(resourceName, "grid_set_commands.set_dscp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupGridSetCommands(name, gridSetCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_set_commands.set_membership", "false"),
					resource.TestCheckResourceAttr(resourceName, "grid_set_commands.set_dscp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_GridShowCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_grid_show_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	gridShowCmds := "show_token:true,show_dscp:true"
	gridShowCmds1 := "show_token:false,show_dscp:false"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupGridShowCommands(name, gridShowCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_show_commands.show_token", "true"),
					resource.TestCheckResourceAttr(resourceName, "grid_show_commands.show_dscp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupGridShowCommands(name, gridShowCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_show_commands.show_token", "false"),
					resource.TestCheckResourceAttr(resourceName, "grid_show_commands.show_dscp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_InactivityLockoutSetting(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_inactivity_lockout_setting"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	inactivityLockoutSetting := `account_inactivity_lockout_enable:true,inactive_days:50,reminder_days:15`
	inactivityLockoutSetting1 := `account_inactivity_lockout_enable:false,inactive_days:20,reminder_days:10`
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupInactivityLockoutSetting(name, inactivityLockoutSetting),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inactivity_lockout_setting.account_inactivity_lockout_enable", "true"),
					resource.TestCheckResourceAttr(resourceName, "inactivity_lockout_setting.inactive_days", "50"),
					resource.TestCheckResourceAttr(resourceName, "inactivity_lockout_setting.reminder_days", "15"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupInactivityLockoutSetting(name, inactivityLockoutSetting1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inactivity_lockout_setting.account_inactivity_lockout_enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "inactivity_lockout_setting.inactive_days", "20"),
					resource.TestCheckResourceAttr(resourceName, "inactivity_lockout_setting.reminder_days", "10"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_LicensingSetCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_licensing_set_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	licensingSetCmds := "set_temp_license:true,set_license:true"
	licensingSetCmds1 := "set_temp_license:false,set_license:false"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupLicensingSetCommands(name, licensingSetCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "licensing_set_commands.set_temp_license", "true"),
					resource.TestCheckResourceAttr(resourceName, "licensing_set_commands.set_license", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupLicensingSetCommands(name, licensingSetCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "licensing_set_commands.set_temp_license", "false"),
					resource.TestCheckResourceAttr(resourceName, "licensing_set_commands.set_license", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_LicensingShowCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_licensing_show_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	licensingShowCmds := "show_license_uid:true,show_license:true"
	licensingShowCmds1 := "show_license_uid:false,show_license:false"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupLicensingShowCommands(name, licensingShowCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "licensing_show_commands.show_license_uid", "true"),
					resource.TestCheckResourceAttr(resourceName, "licensing_show_commands.show_license", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupLicensingShowCommands(name, licensingShowCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "licensing_show_commands.show_license_uid", "false"),
					resource.TestCheckResourceAttr(resourceName, "licensing_show_commands.show_license", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_LockoutSetting(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_lockout_setting"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	lockoutSetting := `failed_lockout_duration:20,never_unlock_user:true`
	lockoutSetting1 := `failed_lockout_duration:30,never_unlock_user:false`
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupLockoutSetting(name, lockoutSetting),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lockout_setting.failed_lockout_duration", "20"),
					resource.TestCheckResourceAttr(resourceName, "lockout_setting.never_unlock_user", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupLockoutSetting(name, lockoutSetting1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lockout_setting.failed_lockout_duration", "30"),
					resource.TestCheckResourceAttr(resourceName, "lockout_setting.never_unlock_user", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_MachineControlToplevelCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_machine_control_toplevel_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	mcTopLevelCmds := `reboot:true,reset:true`
	mcTopLevelCmds1 := `reboot:false,reset:false`
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupMachineControlToplevelCommands(name, mcTopLevelCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "machine_control_toplevel_commands.reboot", "true"),
					resource.TestCheckResourceAttr(resourceName, "machine_control_toplevel_commands.reset", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupMachineControlToplevelCommands(name, mcTopLevelCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "machine_control_toplevel_commands.reboot", "false"),
					resource.TestCheckResourceAttr(resourceName, "machine_control_toplevel_commands.reset", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_Name(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_name"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	name1 := acctest.RandomNameWithPrefix("admin-group1")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupName(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_NetworkingSetCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_networking_set_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	networkingSetCmds := `set_connection_limit:true,set_prompt:true`
	networkingSetCmds1 := `set_connection_limit:false,set_prompt:false`
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupNetworkingSetCommands(name, networkingSetCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "networking_set_commands.set_connection_limit", "true"),
					resource.TestCheckResourceAttr(resourceName, "networking_set_commands.set_prompt", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupNetworkingSetCommands(name, networkingSetCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "networking_set_commands.set_connection_limit", "false"),
					resource.TestCheckResourceAttr(resourceName, "networking_set_commands.set_prompt", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_NetworkingShowCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_networking_show_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	networkingShowCmds := `show_connection_limit:true,show_connections:true`
	networkingShowCmds1 := `show_connection_limit:false,show_connections:false`
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupNetworkingShowCommands(name, networkingShowCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "networking_show_commands.show_connection_limit", "true"),
					resource.TestCheckResourceAttr(resourceName, "networking_show_commands.show_connections", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupNetworkingShowCommands(name, networkingShowCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "networking_show_commands.show_connection_limit", "false"),
					resource.TestCheckResourceAttr(resourceName, "networking_show_commands.show_connections", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_PasswordSetting(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_password_setting"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	passwordSetting := "expire_days:20,expire_enable:true"
	passwordSetting1 := "expire_days:30,expire_enable:true"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupPasswordSetting(name, passwordSetting),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "password_setting.expire_days", "20"),
					resource.TestCheckResourceAttr(resourceName, "password_setting.expire_enable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupPasswordSetting(name, passwordSetting1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "password_setting.expire_days", "30"),
					resource.TestCheckResourceAttr(resourceName, "password_setting.expire_enable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_Roles(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_roles"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	roles := `["DNS Admin","SAML Admin","DHCP Admin"]`
	roles1 := `["DHCP Admin"]`
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupRoles(name, roles),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "roles.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "roles.0", "DNS Admin"),
					resource.TestCheckResourceAttr(resourceName, "roles.1", "SAML Admin"),
					resource.TestCheckResourceAttr(resourceName, "roles.2", "DHCP Admin"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupRoles(name, roles1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "roles.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "roles.0", "DHCP Admin"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_SamlSetting(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_saml_setting"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	samlSetting := "auto_create_user:true,persist_auto_created_user:true"
	samlSetting1 := "auto_create_user:false,persist_auto_created_user:false"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupSamlSetting(name, samlSetting),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "saml_setting.auto_create_user", "true"),
					resource.TestCheckResourceAttr(resourceName, "saml_setting.persist_auto_created_user", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupSamlSetting(name, samlSetting1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "saml_setting.auto_create_user", "false"),
					resource.TestCheckResourceAttr(resourceName, "saml_setting.persist_auto_created_user", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_SecuritySetCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_security_set_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	securitySetCmds := `set_adp:true,set_cc_mode:true`
	securitySetCmds1 := `set_adp:false,set_cc_mode:false`
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupSecuritySetCommands(name, securitySetCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "security_set_commands.set_adp", "true"),
					resource.TestCheckResourceAttr(resourceName, "security_set_commands.set_cc_mode", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupSecuritySetCommands(name, securitySetCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "security_set_commands.set_adp", "false"),
					resource.TestCheckResourceAttr(resourceName, "security_set_commands.set_cc_mode", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_SecurityShowCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_security_show_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	secutiryShowCmds := `show_security:true,show_cc_mode:true`
	secutiryShowCmds1 := `show_security:false,show_cc_mode:false`
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupSecurityShowCommands(name, secutiryShowCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "security_show_commands.show_security", "true"),
					resource.TestCheckResourceAttr(resourceName, "security_show_commands.show_cc_mode", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupSecurityShowCommands(name, secutiryShowCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "security_show_commands.show_security", "false"),
					resource.TestCheckResourceAttr(resourceName, "security_show_commands.show_cc_mode", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_Superuser(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_superuser"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupSuperuser(name, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "superuser", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupSuperuser(name, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "superuser", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_TroubleShootingToplevelCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_trouble_shooting_toplevel_commands"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	troubleShootingTopLevelCmds := `console:true,ping:true,dig:true`
	troubleShootingTopLevelCmds1 := `console:false,ping:false,dig:false`
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupTroubleShootingToplevelCommands(name, troubleShootingTopLevelCmds),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "trouble_shooting_toplevel_commands.console", "true"),
					resource.TestCheckResourceAttr(resourceName, "trouble_shooting_toplevel_commands.ping", "true"),
					resource.TestCheckResourceAttr(resourceName, "trouble_shooting_toplevel_commands.dig", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupTroubleShootingToplevelCommands(name, troubleShootingTopLevelCmds1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "trouble_shooting_toplevel_commands.console", "false"),
					resource.TestCheckResourceAttr(resourceName, "trouble_shooting_toplevel_commands.ping", "false"),
					resource.TestCheckResourceAttr(resourceName, "trouble_shooting_toplevel_commands.dig", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_UseAccountInactivityLockoutEnable(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_use_account_inactivity_lockout_enable"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupUseAccountInactivityLockoutEnable(name, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_account_inactivity_lockout_enable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupUseAccountInactivityLockoutEnable(name, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_account_inactivity_lockout_enable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_UseDisableConcurrentLogin(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_use_disable_concurrent_login"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupUseDisableConcurrentLogin(name, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_disable_concurrent_login", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupUseDisableConcurrentLogin(name, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_disable_concurrent_login", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_UseLockoutSetting(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_use_lockout_setting"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupUseLockoutSetting(name, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_lockout_setting", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupUseLockoutSetting(name, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_lockout_setting", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_UsePasswordSetting(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_use_password_setting"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupUsePasswordSetting(name, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_password_setting", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupUsePasswordSetting(name, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_password_setting", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_UserAccess(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_user_access"
	var v security.Admingroup
	name := acctest.RandomNameWithPrefix("admin-group")
	userAccess := `[{"address":"12.12.1.11","permission":"ALLOW"}]`
	userAccess1 := `[{"address":"12.12.1.12","permission":"ALLOW"}]`
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupUserAccess(name, userAccess),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "user_access.0.address", "12.12.1.11"),
					resource.TestCheckResourceAttr(resourceName, "user_access.0.permission", "ALLOW"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupUserAccess(name, userAccess1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "user_access.0.address", "12.12.1.12"),
					resource.TestCheckResourceAttr(resourceName, "user_access.0.permission", "ALLOW"),
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
    use_disable_concurrent_login = true
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
    email_addresses = %s
}
`, name, emailAddresses)
}

func testAccAdmingroupEnableRestrictedUserAccess(name string, enableRestrictedUserAccess bool) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_enable_restricted_user_access" {
	name = %q
    enable_restricted_user_access = %t
}
`, name, enableRestrictedUserAccess)
}

func testAccAdmingroupExtAttrs(name string, extAttrs map[string]string) string {
	extattrsStr := "{"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`%s = %q`, k, v)
	}
	extattrsStr += "}"
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_extattrs" {
    name = %q
    extattrs = %s
}
`, name, extattrsStr)
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
	name = %q
    grid_show_commands = {%s}
}
`, name, gridShowCommands)
}

func testAccAdmingroupInactivityLockoutSetting(name, inactivityLockoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_inactivity_lockout_setting" {
	name = %q
    inactivity_lockout_setting = {%s}
    use_account_inactivity_lockout_enable = true
}
`, name, inactivityLockoutSetting)
}

func testAccAdmingroupLicensingSetCommands(name, licensingSetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_licensing_set_commands" {
    name = %q
    licensing_set_commands = {%s}
}
`, name, licensingSetCommands)
}

func testAccAdmingroupLicensingShowCommands(name, licensingShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_licensing_show_commands" {
	name = %q
    licensing_show_commands = {%s}
}
`, name, licensingShowCommands)
}

func testAccAdmingroupLockoutSetting(name, lockoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_lockout_setting" {
	name = %q
    lockout_setting = {%s}
}
`, name, lockoutSetting)
}

func testAccAdmingroupMachineControlToplevelCommands(name, machineControlToplevelCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_machine_control_toplevel_commands" {
    name = %q
    machine_control_toplevel_commands = {%s}
}
`, name, machineControlToplevelCommands)
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
    name = %q
    networking_set_commands = {%s}
}
`, name, networkingSetCommands)
}

func testAccAdmingroupNetworkingShowCommands(name, networkingShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_networking_show_commands" {
    name = %q
    networking_show_commands = {%s}
}
`, name, networkingShowCommands)
}

func testAccAdmingroupPasswordSetting(name, passwordSetting string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_password_setting" {
	name = %q
    password_setting = {%s}
    use_password_setting = true
}
`, name, passwordSetting)
}

func testAccAdmingroupRoles(name, roles string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_roles" {
	name = %q
    roles = %s
}
`, name, roles)
}

func testAccAdmingroupSamlSetting(name, samlSetting string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_saml_setting" {
	name = %q
    saml_setting = {%s}
}
`, name, samlSetting)
}

func testAccAdmingroupSecuritySetCommands(name, securitySetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_security_set_commands" {
	name = %q
    security_set_commands = {%s}
}
`, name, securitySetCommands)
}

func testAccAdmingroupSecurityShowCommands(name, securityShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_security_show_commands" {
    name = %q
    security_show_commands = {%s}
}
`, name, securityShowCommands)
}

func testAccAdmingroupSuperuser(name string, superuser bool) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_superuser" {
	name = %q
    superuser = %t
}
`, name, superuser)
}

func testAccAdmingroupTroubleShootingToplevelCommands(name, troubleShootingToplevelCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_trouble_shooting_toplevel_commands" {
    name = %q
    trouble_shooting_toplevel_commands = {%s}
}
`, name, troubleShootingToplevelCommands)
}

func testAccAdmingroupUseAccountInactivityLockoutEnable(name string, useAccountInactivityLockoutEnable bool) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_use_account_inactivity_lockout_enable" {
    name = %q
    use_account_inactivity_lockout_enable = %t
}
`, name, useAccountInactivityLockoutEnable)
}

func testAccAdmingroupUseDisableConcurrentLogin(name string, useDisableConcurrentLogin bool) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_use_disable_concurrent_login" {
    name = %q
    use_disable_concurrent_login = %t
}
`, name, useDisableConcurrentLogin)
}

func testAccAdmingroupUseLockoutSetting(name string, useLockoutSetting bool) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_use_lockout_setting" {
    name = %q
    use_lockout_setting = %t
}
`, name, useLockoutSetting)
}

func testAccAdmingroupUsePasswordSetting(name string, usePasswordSetting bool) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_use_password_setting" {
    name = %q
    use_password_setting = %t
}
`, name, usePasswordSetting)
}

func testAccAdmingroupUserAccess(name, userAccess string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_user_access" {
    name = %q
    user_access = %s
    enable_restricted_user_access = true
}
`, name, userAccess)
}
