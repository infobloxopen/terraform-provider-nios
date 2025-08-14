package security_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForAdminuser = "admin_groups,auth_method,auth_type,ca_certificate_issuer,client_certificate_serial_number,comment,disable,email,enable_certificate_authentication,extattrs,name,ssh_keys,status,time_zone,use_ssh_keys,use_time_zone"

func TestAccAdminuserResource_basic(t *testing.T) {
	var resourceName = "nios_security_admin_user.test"
	var v security.Adminuser
	name := acctest.RandomNameWithPrefix("admin-user")
	password := "Example-Admin123!"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserBasicConfig(name, password, "admin-group"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "auth_method", "KEYPAIR"),
					resource.TestCheckResourceAttr(resourceName, "auth_type", "LOCAL"),
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "email", ""),
					resource.TestCheckResourceAttr(resourceName, "enable_certificate_authentication", "false"),
					resource.TestCheckResourceAttr(resourceName, "time_zone", "UTC"),
					resource.TestCheckResourceAttr(resourceName, "use_time_zone", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ssh_keys", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_disappears(t *testing.T) {
	resourceName := "nios_security_admin_user.test"
	var v security.Adminuser
	name := acctest.RandomNameWithPrefix("admin-user")
	password := "Example-Admin123!"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAdminuserDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAdminuserBasicConfig(name, password, "admin-group"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					testAccCheckAdminuserDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAdminuserResource_AdminGroups(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_admin_groups"
	var v security.Adminuser
	name := acctest.RandomNameWithPrefix("admin-user")
	password := "Example-Admin123!"
	adminGroups := "admin-group"
	adminGroups1 := "opa-group"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserAdminGroups(name, password, adminGroups),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "admin_groups.0", "admin-group"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserAdminGroups(name, password, adminGroups1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "admin_groups.0", "opa-group"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_AuthMethod(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_auth_method"
	var v security.Adminuser
	name := acctest.RandomNameWithPrefix("admin-user")
	password := "Example-Admin123!"
	authMethod := "KEYPAIR"
	authMethod1 := "KEYPAIR_PASSWORD"
	sshkeys, err := filepath.Abs(filepath.Join("..", "..", "..", "examples", "resources", "nios_security_admin_user", "sample_key.pub"))
	if err != nil {
		t.Errorf("Failed to construct absolute path for SSH key: %v", err)
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserAuthMethod(name, password, "admin-group", authMethod, false, nil, nil, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_method", authMethod),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserAuthMethod(name, password, "admin-group", authMethod1, true, utils.Ptr("sample-key"), utils.Ptr("RSA"), utils.Ptr(sshkeys)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_method", authMethod1),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_AuthType(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_auth_type"
	var v security.Adminuser
	name := acctest.RandomNameWithPrefix("admin-user")
	password := "Example-Admin123!"
	authType := "LOCAL"
	authype1 := "REMOTE"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserAuthType(name, password, "admin-group", authType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_type", authType),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserAuthType(name, password, "admin-group", authype1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auth_type", authype1),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_CaCertificateIssuer(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_ca_certificate_issuer"
	var v security.Adminuser
	name := acctest.RandomNameWithPrefix("admin-user")
	password := "Example-Admin123!"
	caCertificateIssuer := "cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22"
	caCertificateIssuer1 := "cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuZGM2MTlhMWYyYmI0NGYwYjUzMWFiNzcwZjk1ZDQ0MDRhNWY2ODQxZGQxOTQ3Y2Q0YjcxMjU1YWU1MjY5MzM1MTRhMDljNWI5OTMwNmNhYzRiMjczY2JhN2NhODYwOWQ5ODY2YWYxYzU3NDdkNTVmNTFjZjM0ZGY4NzRmYTFjZWU:CN%3D%22ib-root-ca%22"
	clientCertificateSerialNumber := "4e7c675cd972ecd2e5b895ad6cb4e38e6d77b4b4"
	clientCertificateSerialNumber1 := "682f792e09c242093c192cee888d44658e88abfb"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserCaCertificateIssuer(name, password, "admin-group", caCertificateIssuer, clientCertificateSerialNumber, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_certificate_issuer", caCertificateIssuer),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserCaCertificateIssuer(name, password, "admin-group", caCertificateIssuer1, clientCertificateSerialNumber1, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_certificate_issuer", caCertificateIssuer1),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_ClientCertificateSerialNumber(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_client_certificate_serial_number"
	var v security.Adminuser
	name := acctest.RandomNameWithPrefix("admin-user")
	password := "Example-Admin123!"
	caCertificateIssuer := "cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22"
	caCertificateIssuer1 := "cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuZGM2MTlhMWYyYmI0NGYwYjUzMWFiNzcwZjk1ZDQ0MDRhNWY2ODQxZGQxOTQ3Y2Q0YjcxMjU1YWU1MjY5MzM1MTRhMDljNWI5OTMwNmNhYzRiMjczY2JhN2NhODYwOWQ5ODY2YWYxYzU3NDdkNTVmNTFjZjM0ZGY4NzRmYTFjZWU:CN%3D%22ib-root-ca%22"
	clientCertificateSerialNumber := "4e7c675cd972ecd2e5b895ad6cb4e38e6d77b4b4"
	clientCertificateSerialNumber1 := "682f792e09c242093c192cee888d44658e88abfb"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserClientCertificateSerialNumber(name, password, "admin-group", caCertificateIssuer, clientCertificateSerialNumber, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_certificate_serial_number", clientCertificateSerialNumber),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserClientCertificateSerialNumber(name, password, "admin-group", caCertificateIssuer1, clientCertificateSerialNumber1, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_certificate_serial_number", clientCertificateSerialNumber1),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_Comment(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_comment"
	var v security.Adminuser
	name := acctest.RandomNameWithPrefix("admin-user")
	password := "Example-Admin123!"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserComment(name, password, "admin-group", "example admin user"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "example admin user"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserComment(name, password, "admin-group", "example admin user updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "example admin user updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_Disable(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_disable"
	var v security.Adminuser
	name := acctest.RandomNameWithPrefix("admin-user")
	password := "Example-Admin123!"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserDisable(name, password, "admin-group", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserDisable(name, password, "admin-group", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_Email(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_email"
	var v security.Adminuser
	name := acctest.RandomNameWithPrefix("admin-user")
	password := "Example-Admin123!"
	email := "abc@example.com"
	email1 := "xyz@sample.com"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserEmail(name, password, "admin-group", email),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email", email),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserEmail(name, password, "admin-group", email1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email", email1),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_EnableCertificateAuthentication(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_enable_certificate_authentication"
	var v security.Adminuser
	name := acctest.RandomNameWithPrefix("admin-user")
	password := "Example-Admin123!"
	caCertificateIssuer := "cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22"
	clientCertificateSerialNumber := "4e7c675cd972ecd2e5b895ad6cb4e38e6d77b4b4"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserEnableCertificateAuthentication(name, password, "admin-group", &caCertificateIssuer, &clientCertificateSerialNumber, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_certificate_authentication", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserEnableCertificateAuthentication(name, password, "admin-group", nil, nil, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_certificate_authentication", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_extattrs"
	var v security.Adminuser
	name := acctest.RandomNameWithPrefix("admin-user")
	password := "Example-Admin123!"
	extAttrValue := acctest.RandomName()
	extAttrValue1 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserExtAttrs(name, password, "admin-group", map[string]string{"Site": extAttrValue}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserExtAttrs(name, password, "admin-group", map[string]string{"Site": extAttrValue1}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_Name(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_name"
	var v security.Adminuser
	name := acctest.RandomNameWithPrefix("admin-user")
	password := "Example-Admin123!"
	name1 := acctest.RandomNameWithPrefix("admin-user-update")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserName(name, password, "admin-group"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserName(name1, password, "admin-group"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_Password(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_password"
	var v security.Adminuser
	name := acctest.RandomNameWithPrefix("admin-user")
	password := "Example-Admin123!"
	password1 := "Example-password123!"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserPassword(name, password, "admin-group"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "password", password),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserPassword(name, password1, "admin-group"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "password", password1),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_SshKeys(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_ssh_keys"
	var v security.Adminuser
	name := acctest.RandomNameWithPrefix("admin-user")
	password := "Example-Admin123!"

	keyFilePath, err := filepath.Abs(filepath.Join("..", "..", "..", "examples", "resources", "nios_security_admin_user", "sample_key.pub"))
	if err != nil {
		t.Errorf("Failed to construct absolute path for SSH key: %v", err)
	}
	keyContent, err := os.ReadFile(keyFilePath)
	if err != nil {
		t.Errorf("Failed to read SSH key file: %v", err)
	}

	// Replace newline characters in the file content
	keyValue := strings.ReplaceAll(string(keyContent), "\n", "")

	keyFilePath1, err := filepath.Abs(filepath.Join("..", "..", "..", "examples", "resources", "nios_security_admin_user", "sample_key.pub"))
	if err != nil {
		t.Errorf("Failed to construct absolute path for SSH key: %v", err)
	}
	keyContent1, err := os.ReadFile(keyFilePath1)
	if err != nil {
		t.Errorf("Failed to read SSH key file: %v", err)
	}

	// Replace newline characters in the file content
	keyValue1 := strings.ReplaceAll(string(keyContent1), "\n", "")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserSshKeys(name, password, "admin-group", "sample-key", "RSA", keyFilePath),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ssh_keys.0.key_name", "sample-key"),
					resource.TestCheckResourceAttr(resourceName, "ssh_keys.0.key_type", "RSA"),
					resource.TestCheckResourceAttr(resourceName, "ssh_keys.0.key_value", keyValue),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserSshKeys(name, password, "admin-group", "sample-key", "RSA", keyFilePath1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ssh_keys.0.key_name", "sample-key"),
					resource.TestCheckResourceAttr(resourceName, "ssh_keys.0.key_type", "RSA"),
					resource.TestCheckResourceAttr(resourceName, "ssh_keys.0.key_value", keyValue1),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_TimeZone(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_time_zone"
	var v security.Adminuser
	name := acctest.RandomNameWithPrefix("admin-user")
	password := "Example-Admin123!"
	timeZone := "UTC"
	timeZone1 := "Singapore"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserTimeZone(name, password, "admin-group", timeZone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "time_zone", timeZone),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserTimeZone(name, password, "admin-group", timeZone1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "time_zone", timeZone1),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_UseSshKeys(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_use_ssh_keys"
	var v security.Adminuser
	name := acctest.RandomNameWithPrefix("admin-user")
	password := "Example-Admin123!"
	sshkeys, err := filepath.Abs(filepath.Join("..", "..", "..", "examples", "resources", "nios_security_admin_user", "sample_key.pub"))
	if err != nil {
		t.Errorf("Failed to construct absolute path for SSH key: %v", err)
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserUseSshKeys(name, password, "admin-group", true, utils.Ptr("sample-key"), utils.Ptr("RSA"), &sshkeys),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ssh_keys", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserUseSshKeys(name, password, "admin-group", false, nil, nil, nil),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ssh_keys", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAdminuserResource_UseTimeZone(t *testing.T) {
	var resourceName = "nios_security_admin_user.test_use_time_zone"
	var v security.Adminuser
	name := acctest.RandomNameWithPrefix("admin-user")
	password := "Example-Admin123!"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAdminuserUseTimeZone(name, password, "admin-group", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_time_zone", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAdminuserUseTimeZone(name, password, "admin-group", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdminuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_time_zone", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckAdminuserExists(ctx context.Context, resourceName string, v *security.Adminuser) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SecurityAPI.
			AdminuserAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForAdminuser).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetAdminuserResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetAdminuserResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckAdminuserDestroy(ctx context.Context, v *security.Adminuser) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SecurityAPI.
			AdminuserAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForAdminuser).
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

func testAccCheckAdminuserDisappears(ctx context.Context, v *security.Adminuser) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SecurityAPI.
			AdminuserAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccAdminuserBasicConfig(name, password, adminGroups string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test" {
  name = %q
  password = %q
  admin_groups = [%q]
}
`, name, password, adminGroups)
}

func testAccAdminuserAdminGroups(name, password, adminGroups string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_admin_groups" {
   name = %q
   password = %q
   admin_groups = [%q]
}
`, name, password, adminGroups)
}

func testAccAdminuserAuthMethod(name, password, authGroups, authMethod string, useSshKeys bool, keyName, keyType, keyValue *string) string {
	var extraConfig string
	if keyName != nil && keyType != nil && keyValue != nil {
		extraConfig = fmt.Sprintf(`
		ssh_keys = [{
          key_name  = %q
          key_type = %q
          key_value = replace(file(%q), "\n", "")
		}]
		`, *keyName, *keyType, *keyValue)
	}
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_auth_method" {
   name = %q
   password = %q
   admin_groups = [%q]
   auth_method = %q
   use_ssh_keys = %t
   %s
}
`, name, password, authGroups, authMethod, useSshKeys, extraConfig)
}

func testAccAdminuserAuthType(name, password, authGroups, authType string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_auth_type" {
   name = %q
   password = %q
   admin_groups = [%q]
   auth_type = %q
}
`, name, password, authGroups, authType)
}

func testAccAdminuserCaCertificateIssuer(name, password, authGroups, caCertificateIssuer, clientCertificateSerialNumber string, enableCertAuth bool) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_ca_certificate_issuer" {
   name = %q
   password = %q
   admin_groups = [%q]
   ca_certificate_issuer = %q
   client_certificate_serial_number = %q 
   enable_certificate_authentication = %t
}
`, name, password, authGroups, caCertificateIssuer, clientCertificateSerialNumber, enableCertAuth)
}

func testAccAdminuserClientCertificateSerialNumber(name, password, authGroups, caCertificateIssuer, clientCertificateSerialNumber string, enableCertAuth bool) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_client_certificate_serial_number" {
   name = %q
   password = %q
   admin_groups = [%q]
   ca_certificate_issuer = %q
   client_certificate_serial_number = %q 
   enable_certificate_authentication = %t
}
`, name, password, authGroups, caCertificateIssuer, clientCertificateSerialNumber, enableCertAuth)
}

func testAccAdminuserComment(name, password, adminGroups, comment string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_comment" {
   name = %q
   password = %q
   admin_groups = [%q]
   comment = %q
}
`, name, password, adminGroups, comment)
}

func testAccAdminuserDisable(name, password, adminGroups string, disable bool) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_disable" {
   name = %q
   password = %q
   admin_groups = [%q]
   disable = %t
}
`, name, password, adminGroups, disable)
}

func testAccAdminuserEmail(name, password, adminGroups, email string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_email" {
   name = %q
   password = %q
   admin_groups = [%q]
   email = %q
}
`, name, password, adminGroups, email)
}

func testAccAdminuserEnableCertificateAuthentication(name, password, adminGroups string, caCertificateIssuer, clientCertificateSerialNumber *string, enableCertificateAuthentication bool) string {
	var extraConfig string
	if caCertificateIssuer != nil && clientCertificateSerialNumber != nil {
		extraConfig = fmt.Sprintf(`
		ca_certificate_issuer = %q
		client_certificate_serial_number = %q
		`, *caCertificateIssuer, *clientCertificateSerialNumber)
	}
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_enable_certificate_authentication" {
   name = %q
   password = %q
   admin_groups = [%q]
   enable_certificate_authentication = %t
   %s
}
`, name, password, adminGroups, enableCertificateAuthentication, extraConfig)
}

func testAccAdminuserExtAttrs(name, password, adminGroups string, extAttrs map[string]string) string {
	extattrsStr := "{"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`%s = %q`, k, v)
	}
	extattrsStr += "}"
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_extattrs" {
   name = %q
   password = %q
   admin_groups = [%q]
   extattrs = %s
}
`, name, password, adminGroups, extattrsStr)
}

func testAccAdminuserName(name, password, adminGroups string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_name" {
   name = %q
   password = %q
   admin_groups = [%q]
}
`, name, password, adminGroups)
}

func testAccAdminuserPassword(name, password, adminGroups string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_password" {
   name = %q
   password = %q
   admin_groups = [%q]
}
`, name, password, adminGroups)
}

func testAccAdminuserSshKeys(name, password, adminGroups, keyName, keyType, keyValue string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_ssh_keys" {
   name = %q
   password = %q
   admin_groups = [%q]
   use_ssh_keys = true
   ssh_keys = [
     {
       key_name  = %q
       key_type  = %q
       key_value = replace(file(%q), "\n", "")
     }
   ]
}
`, name, password, adminGroups, keyName, keyType, keyValue)
}

func testAccAdminuserTimeZone(name, password, adminGroups, timeZone string) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_time_zone" {
   name = %q
   password = %q
   admin_groups = [%q]
   use_time_zone = true
   time_zone = %q
}
`, name, password, adminGroups, timeZone)
}

func testAccAdminuserUseSshKeys(name, password, adminGroups string, useSshKeys bool, keyName, keyType, keyValue *string) string {
	var extraConfig string
	if keyName != nil && keyType != nil && keyValue != nil {
		extraConfig = fmt.Sprintf(`
		ssh_keys = [{
          key_name  = %q
          key_type = %q
          key_value = replace(file(%q), "\n", "")
		}]
		`, *keyName, *keyType, *keyValue)
	}
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_use_ssh_keys" {
   name = %q
   password = %q
   admin_groups = [%q]
   use_ssh_keys = %t
   %s
}
`, name, password, adminGroups, useSshKeys, extraConfig)
}

func testAccAdminuserUseTimeZone(name, password, adminGroups string, useTimeZone bool) string {
	return fmt.Sprintf(`
resource "nios_security_admin_user" "test_use_time_zone" {
   name = %q
   password = %q
   admin_groups = [%q]
   use_time_zone = %t
}
`, name, password, adminGroups, useTimeZone)
}
