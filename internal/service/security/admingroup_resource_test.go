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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_disappears(t *testing.T) {
	resourceName := "nios_security_admin_group.test"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAdmingroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAdmingroupBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					testAccCheckAdmingroupDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAdmingroupResource_Ref(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_ref"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_AccessMethod(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_access_method"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupAccessMethod("ACCESS_METHOD_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "access_method", "ACCESS_METHOD_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupAccessMethod("ACCESS_METHOD_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "access_method", "ACCESS_METHOD_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_AdminSetCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_admin_set_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupAdminSetCommands("ADMIN_SET_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "admin_set_commands", "ADMIN_SET_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupAdminSetCommands("ADMIN_SET_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "admin_set_commands", "ADMIN_SET_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_AdminShowCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_admin_show_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupAdminShowCommands("ADMIN_SHOW_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "admin_show_commands", "ADMIN_SHOW_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupAdminShowCommands("ADMIN_SHOW_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "admin_show_commands", "ADMIN_SHOW_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_AdminToplevelCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_admin_toplevel_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupAdminToplevelCommands("ADMIN_TOPLEVEL_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "admin_toplevel_commands", "ADMIN_TOPLEVEL_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupAdminToplevelCommands("ADMIN_TOPLEVEL_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "admin_toplevel_commands", "ADMIN_TOPLEVEL_COMMANDS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdmingroupResource_CloudSetCommands(t *testing.T) {
	var resourceName = "nios_security_admin_group.test_cloud_set_commands"
	var v security.Admingroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdmingroupCloudSetCommands("CLOUD_SET_COMMANDS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_set_commands", "CLOUD_SET_COMMANDS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAdmingroupCloudSetCommands("CLOUD_SET_COMMANDS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmingroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_set_commands", "CLOUD_SET_COMMANDS_UPDATE_REPLACE_ME"),
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

func testAccAdmingroupBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test" {
}
`)
}

func testAccAdmingroupRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccAdmingroupAccessMethod(accessMethod string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_access_method" {
    access_method = %q
}
`, accessMethod)
}

func testAccAdmingroupAdminSetCommands(adminSetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_admin_set_commands" {
    admin_set_commands = %q
}
`, adminSetCommands)
}

func testAccAdmingroupAdminShowCommands(adminShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_admin_show_commands" {
    admin_show_commands = %q
}
`, adminShowCommands)
}

func testAccAdmingroupAdminToplevelCommands(adminToplevelCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_admin_toplevel_commands" {
    admin_toplevel_commands = %q
}
`, adminToplevelCommands)
}

func testAccAdmingroupCloudSetCommands(cloudSetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_cloud_set_commands" {
    cloud_set_commands = %q
}
`, cloudSetCommands)
}

func testAccAdmingroupCloudShowCommands(cloudShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_cloud_show_commands" {
    cloud_show_commands = %q
}
`, cloudShowCommands)
}

func testAccAdmingroupComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccAdmingroupDatabaseSetCommands(databaseSetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_database_set_commands" {
    database_set_commands = %q
}
`, databaseSetCommands)
}

func testAccAdmingroupDatabaseShowCommands(databaseShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_database_show_commands" {
    database_show_commands = %q
}
`, databaseShowCommands)
}

func testAccAdmingroupDhcpSetCommands(dhcpSetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_dhcp_set_commands" {
    dhcp_set_commands = %q
}
`, dhcpSetCommands)
}

func testAccAdmingroupDhcpShowCommands(dhcpShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_dhcp_show_commands" {
    dhcp_show_commands = %q
}
`, dhcpShowCommands)
}

func testAccAdmingroupDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccAdmingroupDisableConcurrentLogin(disableConcurrentLogin string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_disable_concurrent_login" {
    disable_concurrent_login = %q
}
`, disableConcurrentLogin)
}

func testAccAdmingroupDnsSetCommands(dnsSetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_dns_set_commands" {
    dns_set_commands = %q
}
`, dnsSetCommands)
}

func testAccAdmingroupDnsShowCommands(dnsShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_dns_show_commands" {
    dns_show_commands = %q
}
`, dnsShowCommands)
}

func testAccAdmingroupDnsToplevelCommands(dnsToplevelCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_dns_toplevel_commands" {
    dns_toplevel_commands = %q
}
`, dnsToplevelCommands)
}

func testAccAdmingroupDockerSetCommands(dockerSetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_docker_set_commands" {
    docker_set_commands = %q
}
`, dockerSetCommands)
}

func testAccAdmingroupDockerShowCommands(dockerShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_docker_show_commands" {
    docker_show_commands = %q
}
`, dockerShowCommands)
}

func testAccAdmingroupEmailAddresses(emailAddresses string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_email_addresses" {
    email_addresses = %q
}
`, emailAddresses)
}

func testAccAdmingroupEnableRestrictedUserAccess(enableRestrictedUserAccess string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_enable_restricted_user_access" {
    enable_restricted_user_access = %q
}
`, enableRestrictedUserAccess)
}

func testAccAdmingroupExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccAdmingroupGridSetCommands(gridSetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_grid_set_commands" {
    grid_set_commands = %q
}
`, gridSetCommands)
}

func testAccAdmingroupGridShowCommands(gridShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_grid_show_commands" {
    grid_show_commands = %q
}
`, gridShowCommands)
}

func testAccAdmingroupInactivityLockoutSetting(inactivityLockoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_inactivity_lockout_setting" {
    inactivity_lockout_setting = %q
}
`, inactivityLockoutSetting)
}

func testAccAdmingroupLicensingSetCommands(licensingSetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_licensing_set_commands" {
    licensing_set_commands = %q
}
`, licensingSetCommands)
}

func testAccAdmingroupLicensingShowCommands(licensingShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_licensing_show_commands" {
    licensing_show_commands = %q
}
`, licensingShowCommands)
}

func testAccAdmingroupLockoutSetting(lockoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_lockout_setting" {
    lockout_setting = %q
}
`, lockoutSetting)
}

func testAccAdmingroupMachineControlToplevelCommands(machineControlToplevelCommands string) string {
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

func testAccAdmingroupNetworkingSetCommands(networkingSetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_networking_set_commands" {
    networking_set_commands = %q
}
`, networkingSetCommands)
}

func testAccAdmingroupNetworkingShowCommands(networkingShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_networking_show_commands" {
    networking_show_commands = %q
}
`, networkingShowCommands)
}

func testAccAdmingroupPasswordSetting(passwordSetting string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_password_setting" {
    password_setting = %q
}
`, passwordSetting)
}

func testAccAdmingroupRoles(roles string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_roles" {
    roles = %q
}
`, roles)
}

func testAccAdmingroupSamlSetting(samlSetting string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_saml_setting" {
    saml_setting = %q
}
`, samlSetting)
}

func testAccAdmingroupSecuritySetCommands(securitySetCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_security_set_commands" {
    security_set_commands = %q
}
`, securitySetCommands)
}

func testAccAdmingroupSecurityShowCommands(securityShowCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_security_show_commands" {
    security_show_commands = %q
}
`, securityShowCommands)
}

func testAccAdmingroupSuperuser(superuser string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_superuser" {
    superuser = %q
}
`, superuser)
}

func testAccAdmingroupTroubleShootingToplevelCommands(troubleShootingToplevelCommands string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_trouble_shooting_toplevel_commands" {
    trouble_shooting_toplevel_commands = %q
}
`, troubleShootingToplevelCommands)
}

func testAccAdmingroupUseAccountInactivityLockoutEnable(useAccountInactivityLockoutEnable string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_use_account_inactivity_lockout_enable" {
    use_account_inactivity_lockout_enable = %q
}
`, useAccountInactivityLockoutEnable)
}

func testAccAdmingroupUseDisableConcurrentLogin(useDisableConcurrentLogin string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_use_disable_concurrent_login" {
    use_disable_concurrent_login = %q
}
`, useDisableConcurrentLogin)
}

func testAccAdmingroupUseLockoutSetting(useLockoutSetting string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_use_lockout_setting" {
    use_lockout_setting = %q
}
`, useLockoutSetting)
}

func testAccAdmingroupUsePasswordSetting(usePasswordSetting string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_use_password_setting" {
    use_password_setting = %q
}
`, usePasswordSetting)
}

func testAccAdmingroupUserAccess(userAccess string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_group" "test_user_access" {
    user_access = %q
}
`, userAccess)
}
