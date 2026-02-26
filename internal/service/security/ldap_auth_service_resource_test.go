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

var readableAttributesForLdapAuthService = "comment,disable,ea_mapping,ldap_group_attribute,ldap_group_authentication_type,ldap_user_attribute,mode,name,recovery_interval,retries,search_scope,servers,timeout"

func TestAccLdapAuthServiceResource_basic(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test"
	var v security.LdapAuthService
	name := acctest.RandomNameWithPrefix("ldap-auth-service")
	servers := []map[string]any{
		{
			"address":             "2.2.2.2",
			"authentication_type": "ANONYMOUS",
			"base_dn":             "ou=People,dc=example,dc=com",
			"disable":             false,
			"encryption":          "SSL",
			"port":                636,
			"use_mgmt_port":       false,
			"version":             "V3",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceBasicConfig(name, servers, "adminID", 30, 5, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.address", "2.2.2.2"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.authentication_type", "ANONYMOUS"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.base_dn", "ou=People,dc=example,dc=com"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.encryption", "SSL"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.port", "636"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.use_mgmt_port", "false"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.version", "V3"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "search_scope", "ONELEVEL"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "ea_mapping.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "ldap_group_attribute", "memberOf"),
					resource.TestCheckResourceAttr(resourceName, "ldap_group_authentication_type", "GROUP_ATTRIBUTE"),
					resource.TestCheckResourceAttr(resourceName, "mode", "ORDERED_LIST"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_disappears(t *testing.T) {
	resourceName := "nios_security_ldap_auth_service.test"
	var v security.LdapAuthService
	name := acctest.RandomNameWithPrefix("ldap-auth-service")
	servers := []map[string]any{
		{
			"address":             "2.2.2.2",
			"authentication_type": "ANONYMOUS",
			"base_dn":             "ou=People,dc=example,dc=com",
			"disable":             false,
			"encryption":          "SSL",
			"port":                636,
			"use_mgmt_port":       false,
			"version":             "V3",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLdapAuthServiceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccLdapAuthServiceBasicConfig(name, servers, "adminID", 30, 5, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					testAccCheckLdapAuthServiceDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccLdapAuthServiceResource_Comment(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_comment"
	var v security.LdapAuthService
	name := acctest.RandomNameWithPrefix("ldap-auth-service")
	servers := []map[string]any{
		{
			"address":             "2.2.2.2",
			"authentication_type": "ANONYMOUS",
			"base_dn":             "ou=People,dc=example,dc=com",
			"disable":             false,
			"encryption":          "SSL",
			"port":                636,
			"use_mgmt_port":       false,
			"version":             "V3",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceComment(name, servers, "adminID", 30, 5, 5, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceComment(name, servers, "adminID", 30, 5, 5, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_Disable(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_disable"
	var v security.LdapAuthService
	name := acctest.RandomNameWithPrefix("ldap-auth-service")
	servers := []map[string]any{
		{
			"address":             "2.2.2.2",
			"authentication_type": "ANONYMOUS",
			"base_dn":             "ou=People,dc=example,dc=com",
			"disable":             false,
			"encryption":          "SSL",
			"port":                636,
			"use_mgmt_port":       false,
			"version":             "V3",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceDisable(name, servers, "adminID", 30, 5, 5, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceDisable(name, servers, "adminID", 30, 5, 5, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_EaMapping(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_ea_mapping"
	var v security.LdapAuthService
	name := acctest.RandomNameWithPrefix("ldap-auth-service")
	servers := []map[string]any{
		{
			"address":             "2.2.2.2",
			"authentication_type": "ANONYMOUS",
			"base_dn":             "ou=People,dc=example,dc=com",
			"disable":             false,
			"encryption":          "SSL",
			"port":                636,
			"use_mgmt_port":       false,
			"version":             "V3",
		},
	}
	eaMapping := []map[string]any{
		{
			"mapped_ea": "Availability zone",
			"name":      "ldapfield",
		},
	}
	eaMappingUpdate := []map[string]any{
		{
			"mapped_ea": "Subnet Name",
			"name":      "ldapfield12",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceEaMapping(name, servers, "adminID", 30, 5, 5, eaMapping),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ea_mapping.0.mapped_ea", "Availability zone"),
					resource.TestCheckResourceAttr(resourceName, "ea_mapping.0.name", "ldapfield"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceEaMapping(name, servers, "adminID", 30, 5, 5, eaMappingUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ea_mapping.0.mapped_ea", "Subnet Name"),
					resource.TestCheckResourceAttr(resourceName, "ea_mapping.0.name", "ldapfield12"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_LdapGroupAttribute(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_ldap_group_attribute"
	var v security.LdapAuthService
	name := acctest.RandomNameWithPrefix("ldap-auth-service")
	servers := []map[string]any{
		{
			"address":             "2.2.2.2",
			"authentication_type": "ANONYMOUS",
			"base_dn":             "ou=People,dc=example,dc=com",
			"disable":             false,
			"encryption":          "SSL",
			"port":                636,
			"use_mgmt_port":       false,
			"version":             "V3",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceLdapGroupAttribute(name, servers, "adminID", 30, 5, 5, "namecn"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ldap_group_attribute", "namecn"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceLdapGroupAttribute(name, servers, "adminID", 30, 5, 5, "namecnid"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ldap_group_attribute", "namecnid"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_LdapGroupAuthenticationType(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_ldap_group_authentication_type"
	var v security.LdapAuthService
	name := acctest.RandomNameWithPrefix("ldap-auth-service")
	servers := []map[string]any{
		{
			"address":             "2.2.2.2",
			"authentication_type": "ANONYMOUS",
			"base_dn":             "ou=People,dc=example,dc=com",
			"disable":             false,
			"encryption":          "SSL",
			"port":                636,
			"use_mgmt_port":       false,
			"version":             "V3",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceLdapGroupAuthenticationType(name, servers, "adminID", 30, 5, 5, "POSIX_GROUP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ldap_group_authentication_type", "POSIX_GROUP"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceLdapGroupAuthenticationType(name, servers, "adminID", 30, 5, 5, "GROUP_ATTRIBUTE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ldap_group_authentication_type", "GROUP_ATTRIBUTE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_LdapUserAttribute(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_ldap_user_attribute"
	var v security.LdapAuthService
	name := acctest.RandomNameWithPrefix("ldap-auth-service")
	servers := []map[string]any{
		{
			"address":             "2.2.2.2",
			"authentication_type": "ANONYMOUS",
			"base_dn":             "ou=People,dc=example,dc=com",
			"disable":             false,
			"encryption":          "SSL",
			"port":                636,
			"use_mgmt_port":       false,
			"version":             "V3",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceLdapUserAttribute(name, servers, "adminID", 30, 5, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ldap_user_attribute", "adminID"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceLdapUserAttribute(name, servers, "adminID12", 30, 5, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ldap_user_attribute", "adminID12"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_Mode(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_mode"
	var v security.LdapAuthService
	name := acctest.RandomNameWithPrefix("ldap-auth-service")
	servers := []map[string]any{
		{
			"address":             "2.2.2.2",
			"authentication_type": "ANONYMOUS",
			"base_dn":             "ou=People,dc=example,dc=com",
			"disable":             false,
			"encryption":          "SSL",
			"port":                636,
			"use_mgmt_port":       false,
			"version":             "V3",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceMode(name, servers, "adminID", 30, 5, 5, "ORDERED_LIST"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mode", "ORDERED_LIST"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceMode(name, servers, "adminID", 30, 5, 5, "ROUND_ROBIN"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mode", "ROUND_ROBIN"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_Name(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_name"
	var v security.LdapAuthService
	name := acctest.RandomNameWithPrefix("ldap-auth-service")
	nameUpdate := acctest.RandomNameWithPrefix("ldap-auth-service")
	servers := []map[string]any{
		{
			"address":             "2.2.2.2",
			"authentication_type": "ANONYMOUS",
			"base_dn":             "ou=People,dc=example,dc=com",
			"disable":             false,
			"encryption":          "SSL",
			"port":                636,
			"use_mgmt_port":       false,
			"version":             "V3",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceName(name, servers, "adminID", 30, 5, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceName(nameUpdate, servers, "adminID", 30, 5, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_RecoveryInterval(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_recovery_interval"
	var v security.LdapAuthService
	name := acctest.RandomNameWithPrefix("ldap-auth-service")
	servers := []map[string]any{
		{
			"address":             "2.2.2.2",
			"authentication_type": "ANONYMOUS",
			"base_dn":             "ou=People,dc=example,dc=com",
			"disable":             false,
			"encryption":          "SSL",
			"port":                636,
			"use_mgmt_port":       false,
			"version":             "V3",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceRecoveryInterval(name, servers, "adminID", 30, 5, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recovery_interval", "30"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceRecoveryInterval(name, servers, "adminID", 60, 5, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recovery_interval", "60"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_Retries(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_retries"
	var v security.LdapAuthService
	name := acctest.RandomNameWithPrefix("ldap-auth-service")
	servers := []map[string]any{
		{
			"address":             "2.2.2.2",
			"authentication_type": "ANONYMOUS",
			"base_dn":             "ou=People,dc=example,dc=com",
			"disable":             false,
			"encryption":          "SSL",
			"port":                636,
			"use_mgmt_port":       false,
			"version":             "V3",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceRetries(name, servers, "adminID", 30, 2, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retries", "2"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceRetries(name, servers, "adminID", 30, 5, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "retries", "5"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_SearchScope(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_search_scope"
	var v security.LdapAuthService
	name := acctest.RandomNameWithPrefix("ldap-auth-service")
	servers := []map[string]any{
		{
			"address":             "2.2.2.2",
			"authentication_type": "ANONYMOUS",
			"base_dn":             "ou=People,dc=example,dc=com",
			"disable":             false,
			"encryption":          "SSL",
			"port":                636,
			"use_mgmt_port":       false,
			"version":             "V3",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceSearchScope(name, servers, "adminID", 30, 5, 5, "BASE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "search_scope", "BASE"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceSearchScope(name, servers, "adminID", 30, 5, 5, "SUBTREE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "search_scope", "SUBTREE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_Servers(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_servers"
	var v security.LdapAuthService
	name := acctest.RandomNameWithPrefix("ldap-auth-service")
	servers := []map[string]any{
		{
			"address":             "2.2.2.2",
			"authentication_type": "ANONYMOUS",
			"base_dn":             "ou=People,dc=example,dc=com",
			"disable":             false,
			"encryption":          "SSL",
			"port":                636,
			"use_mgmt_port":       false,
			"version":             "V3",
		},
	}
	serversUpdate := []map[string]any{
		{
			"address":             "2.2.2.4",
			"authentication_type": "ANONYMOUS",
			"base_dn":             "ou=People,dc=example1,dc=com",
			"disable":             true,
			"encryption":          "SSL",
			"port":                631,
			"use_mgmt_port":       false,
			"version":             "V2",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceServers(name, servers, "adminID", 30, 5, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "servers.0.address", "2.2.2.2"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.authentication_type", "ANONYMOUS"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.base_dn", "ou=People,dc=example,dc=com"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.encryption", "SSL"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.port", "636"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.use_mgmt_port", "false"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.version", "V3"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceServers(name, serversUpdate, "adminID", 30, 5, 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "servers.0.address", "2.2.2.4"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.authentication_type", "ANONYMOUS"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.base_dn", "ou=People,dc=example1,dc=com"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.disable", "true"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.encryption", "SSL"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.port", "631"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.use_mgmt_port", "false"),
					resource.TestCheckResourceAttr(resourceName, "servers.0.version", "V2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccLdapAuthServiceResource_Timeout(t *testing.T) {
	var resourceName = "nios_security_ldap_auth_service.test_timeout"
	var v security.LdapAuthService
	name := acctest.RandomNameWithPrefix("ldap-auth-service")
	servers := []map[string]any{
		{
			"address":             "2.2.2.2",
			"authentication_type": "ANONYMOUS",
			"base_dn":             "ou=People,dc=example,dc=com",
			"disable":             false,
			"encryption":          "SSL",
			"port":                636,
			"use_mgmt_port":       false,
			"version":             "V3",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccLdapAuthServiceTimeout(name, servers, "adminID", 30, 5, 15),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "15"),
				),
			},
			// Update and Read
			{
				Config: testAccLdapAuthServiceTimeout(name, servers, "adminID", 30, 5, 25),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapAuthServiceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "25"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckLdapAuthServiceExists(ctx context.Context, resourceName string, v *security.LdapAuthService) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SecurityAPI.
			LdapAuthServiceAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForLdapAuthService).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetLdapAuthServiceResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetLdapAuthServiceResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckLdapAuthServiceDestroy(ctx context.Context, v *security.LdapAuthService) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SecurityAPI.
			LdapAuthServiceAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForLdapAuthService).
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

func testAccCheckLdapAuthServiceDisappears(ctx context.Context, v *security.LdapAuthService) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SecurityAPI.
			LdapAuthServiceAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccLdapAuthServiceBasicConfig(name string, servers []map[string]any, ldapUserAttribute string, recoveryInterval, retries, timeout int) string {
	serverString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test" {
	name = %q
ldap_user_attribute = %q
recovery_interval = %d
retries = %d
timeout = %d
servers = %s
}
`, name, ldapUserAttribute, recoveryInterval, retries, timeout, serverString)
}

func testAccLdapAuthServiceComment(name string, servers []map[string]any, ldapUserAttribute string, recoveryInterval, retries, timeout int, comment string) string {
	serverString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_comment" {
	name = %q
ldap_user_attribute = %q
recovery_interval = %d
retries = %d
timeout = %d
servers = %s
    comment = %q
}
`, name, ldapUserAttribute, recoveryInterval, retries, timeout, serverString, comment)
}

func testAccLdapAuthServiceDisable(name string, servers []map[string]any, ldapUserAttribute string, recoveryInterval, retries, timeout int, disable bool) string {
	serverString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_disable" {
	name = %q
ldap_user_attribute = %q
recovery_interval = %d
retries = %d
timeout = %d
servers = %s
    disable = %t
}
`, name, ldapUserAttribute, recoveryInterval, retries, timeout, serverString, disable)
}

func testAccLdapAuthServiceEaMapping(name string, servers []map[string]any, ldapUserAttribute string, recoveryInterval, retries, timeout int, eaMapping []map[string]any) string {
	serverString := utils.ConvertSliceOfMapsToHCL(servers)
	eaMappingString := utils.ConvertSliceOfMapsToHCL(eaMapping)
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_ea_mapping" {
	name = %q
ldap_user_attribute = %q
recovery_interval = %d
retries = %d
timeout = %d
servers = %s
    ea_mapping = %s
}
`, name, ldapUserAttribute, recoveryInterval, retries, timeout, serverString, eaMappingString)
}

func testAccLdapAuthServiceLdapGroupAttribute(name string, servers []map[string]any, ldapUserAttribute string, recoveryInterval, retries, timeout int, ldapGroupAttribute string) string {
	serverString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_ldap_group_attribute" {
	name = %q
ldap_user_attribute = %q
recovery_interval = %d
retries = %d
timeout = %d
servers = %s
    ldap_group_attribute = %q
}
`, name, ldapUserAttribute, recoveryInterval, retries, timeout, serverString, ldapGroupAttribute)
}

func testAccLdapAuthServiceLdapGroupAuthenticationType(name string, servers []map[string]any, ldapUserAttribute string, recoveryInterval, retries, timeout int, ldapGroupAuthenticationType string) string {
	serverString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_ldap_group_authentication_type" {
	name = %q
ldap_user_attribute = %q
recovery_interval = %d
retries = %d
timeout = %d
servers = %s
    ldap_group_authentication_type = %q
}
`, name, ldapUserAttribute, recoveryInterval, retries, timeout, serverString, ldapGroupAuthenticationType)
}

func testAccLdapAuthServiceLdapUserAttribute(name string, servers []map[string]any, ldapUserAttribute string, recoveryInterval, retries, timeout int) string {
	serverString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_ldap_user_attribute" {
	name = %q
recovery_interval = %d
retries = %d
timeout = %d
servers = %s
    ldap_user_attribute = %q
}
`, name, recoveryInterval, retries, timeout, serverString, ldapUserAttribute)
}

func testAccLdapAuthServiceMode(name string, servers []map[string]any, ldapUserAttribute string, recoveryInterval, retries, timeout int, mode string) string {
	serverString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_mode" {
	name = %q
ldap_user_attribute = %q
recovery_interval = %d
retries = %d
timeout = %d
servers = %s
    mode = %q
}
`, name, ldapUserAttribute, recoveryInterval, retries, timeout, serverString, mode)
}

func testAccLdapAuthServiceName(name string, servers []map[string]any, ldapUserAttribute string, recoveryInterval, retries, timeout int) string {
	serverString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_name" {
    name = %q
ldap_user_attribute = %q
recovery_interval = %d
retries = %d
timeout = %d
servers = %s
}
`, name, ldapUserAttribute, recoveryInterval, retries, timeout, serverString)
}

func testAccLdapAuthServiceRecoveryInterval(name string, servers []map[string]any, ldapUserAttribute string, recoveryInterval, retries, timeout int) string {
	serverString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_recovery_interval" {
	name = %q
ldap_user_attribute = %q
    recovery_interval = %d
retries = %d
timeout = %d
servers = %s
}
`, name, ldapUserAttribute, recoveryInterval, retries, timeout, serverString)
}

func testAccLdapAuthServiceRetries(name string, servers []map[string]any, ldapUserAttribute string, recoveryInterval, retries, timeout int) string {
	serverString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_retries" {
	name = %q
ldap_user_attribute = %q
recovery_interval = %d
timeout = %d
servers = %s
    retries = %d
}
`, name, ldapUserAttribute, recoveryInterval, timeout, serverString, retries)
}

func testAccLdapAuthServiceSearchScope(name string, servers []map[string]any, ldapUserAttribute string, recoveryInterval, retries, timeout int, searchScope string) string {
	serverString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_search_scope" {
	name = %q
ldap_user_attribute = %q
recovery_interval = %d
retries = %d
timeout = %d
servers = %s
    search_scope = %q
}
`, name, ldapUserAttribute, recoveryInterval, retries, timeout, serverString, searchScope)
}

func testAccLdapAuthServiceServers(name string, servers []map[string]any, ldapUserAttribute string, recoveryInterval, retries, timeout int) string {
	serversString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_servers" {
	name = %q
ldap_user_attribute = %q
recovery_interval = %d
retries = %d
timeout = %d
    servers = %s
}
`, name, ldapUserAttribute, recoveryInterval, retries, timeout, serversString)
}

func testAccLdapAuthServiceTimeout(name string, servers []map[string]any, ldapUserAttribute string, recoveryInterval, retries, timeout int) string {
	serversString := utils.ConvertSliceOfMapsToHCL(servers)
	return fmt.Sprintf(`
resource "nios_security_ldap_auth_service" "test_timeout" {
	name = %q
ldap_user_attribute = %q
recovery_interval = %d
retries = %d
servers = %s
    timeout = %d
}
`, name, ldapUserAttribute, recoveryInterval, retries, serversString, timeout)
}
