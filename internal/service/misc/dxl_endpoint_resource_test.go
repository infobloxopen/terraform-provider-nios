package misc_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// OBJECTS TO BE PRESENT IN GRID FOR TESTS
// DXL Template : "Version5_DXL_Session_Template", "Version5_DXL_Session_Template_2"
// Grid Master Candidate : "infoblox.member, infoblox.grid_master_candidate2"

var readableAttributesForDxlEndpoint = "brokers,client_certificate_subject,client_certificate_valid_from,client_certificate_valid_to,comment,disable,extattrs,log_level,name,outbound_member_type,outbound_members,template_instance,timeout,topics,vendor_identifier,wapi_user_name"

var (
	testDataPath          = getTestDataPath()
	clientCertificateFile = filepath.Join(testDataPath, "client.pem")
	broker                = []map[string]any{
		{
			"host_name": "example.com",
		},
	}
)

func TestAccDxlEndpointResource_basic(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test"
	var v misc.DxlEndpoint
	name := acctest.RandomNameWithPrefix("dxl-endpoint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointBasicConfig(clientCertificateFile, name, "GM", broker),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", "GM"),
					resource.TestCheckResourceAttr(resourceName, "brokers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "brokers.0.host_name", "example.com"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "log_level", "WARNING"),
					resource.TestCheckResourceAttr(resourceName, "timeout", "30"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_disappears(t *testing.T) {
	resourceName := "nios_misc_dxl_endpoint.test"
	var v misc.DxlEndpoint
	name := acctest.RandomNameWithPrefix("dxl-endpoint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDxlEndpointDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDxlEndpointBasicConfig(clientCertificateFile, name, "GM", broker),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					testAccCheckDxlEndpointDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDxlEndpointResource_Brokers(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_brokers"
	var v misc.DxlEndpoint
	name := acctest.RandomNameWithPrefix("dxl-endpoint")
	brokersUpdated := []map[string]any{
		{
			"host_name": "example.com",
			"port":      1234,
		},
		{
			"host_name": "example2.com",
			"port":      5678,
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointBrokers(clientCertificateFile, name, "GM", broker),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "brokers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "brokers.0.host_name", "example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointBrokers(clientCertificateFile, name, "GM", brokersUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "brokers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "brokers.0.host_name", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "brokers.0.port", "1234"),
					resource.TestCheckResourceAttr(resourceName, "brokers.1.host_name", "example2.com"),
					resource.TestCheckResourceAttr(resourceName, "brokers.1.port", "5678"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_BrokersImportToken(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_brokers_import_token"
	var v misc.DxlEndpoint
	name := acctest.RandomNameWithPrefix("dxl-endpoint")
	testDataPath := getTestDataPath()
	brokerPropertiesFile := filepath.Join(testDataPath, "brokerlist.properties")
	brokerPropertiesFileUpdated := filepath.Join(testDataPath, "brokerlist_updated.properties")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointBrokersImportToken(clientCertificateFile, name, "GM", brokerPropertiesFile),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointBrokersImportToken(clientCertificateFile, name, "GM", brokerPropertiesFileUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_ClientCertificateToken(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_client_certificate_token"
	var v misc.DxlEndpoint
	name := acctest.RandomNameWithPrefix("dxl-endpoint")
	testDataPath := getTestDataPath()
	clientCertificateFile := filepath.Join(testDataPath, "client.pem")
	updatedClientCertificateFile := filepath.Join(testDataPath, "client_updated.pem")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointClientCertificateToken(clientCertificateFile, name, "GM", broker),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointClientCertificateToken(updatedClientCertificateFile, name, "GM", broker),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_Comment(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_comment"
	var v misc.DxlEndpoint
	name := acctest.RandomNameWithPrefix("dxl-endpoint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointComment(clientCertificateFile, broker, name, "GM", "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointComment(clientCertificateFile, broker, name, "GM", "Updated comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the object"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_Disable(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_disable"
	var v misc.DxlEndpoint
	name := acctest.RandomNameWithPrefix("dxl-endpoint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointDisable(clientCertificateFile, broker, name, "GM", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointDisable(clientCertificateFile, broker, name, "GM", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_extattrs"
	var v misc.DxlEndpoint
	name := acctest.RandomNameWithPrefix("dxl-endpoint")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointExtAttrs(clientCertificateFile, broker, name, "GM", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointExtAttrs(clientCertificateFile, broker, name, "GM", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_LogLevel(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_log_level"
	var v misc.DxlEndpoint
	name := acctest.RandomNameWithPrefix("dxl-endpoint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointLogLevelDefault(clientCertificateFile, broker, name, "GM"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "WARNING"),
				),
			},
			{
				Config: testAccDxlEndpointLogLevel(clientCertificateFile, broker, name, "GM", "DEBUG"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "DEBUG"),
				),
			},
			{
				Config: testAccDxlEndpointLogLevel(clientCertificateFile, broker, name, "GM", "ERROR"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "ERROR"),
				),
			},
			{
				Config: testAccDxlEndpointLogLevel(clientCertificateFile, broker, name, "GM", "INFO"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_level", "INFO"),
				),
			},
		},
	})
}

func TestAccDxlEndpointResource_Name(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_name"
	var v misc.DxlEndpoint
	name := acctest.RandomNameWithPrefix("dxl-endpoint")
	nameUpdated := acctest.RandomNameWithPrefix("dxl-endpoint-updated")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointName(clientCertificateFile, broker, name, "GM"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointName(clientCertificateFile, broker, nameUpdated, "GM"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_OutboundMemberType(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_outbound_member_type"
	var v misc.DxlEndpoint
	name := acctest.RandomNameWithPrefix("dxl-endpoint")
	memberUpdatedName := "infoblox.member2"
	outboundMembers := []string{memberUpdatedName}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointOutboundMemberType(clientCertificateFile, broker, name, "GM"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", "GM"),
				),
			},
			{
				Config: testAccDxlEndpointOutboundMemberTypeUpdate(clientCertificateFile, broker, name, "MEMBER", outboundMembers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_member_type", "MEMBER"),
				),
			},
		},
	})
}

func TestAccDxlEndpointResource_OutboundMembers(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_outbound_members"
	var v misc.DxlEndpoint
	name := acctest.RandomNameWithPrefix("dxl-endpoint")
	memberUpdatedName := "infoblox.member2"
	outboundMembersVal := []string{memberUpdatedName}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointOutboundMembers(clientCertificateFile, broker, name, "MEMBER", outboundMembersVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "outbound_members.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "outbound_members.0", memberUpdatedName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_TemplateInstance(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_template_instance"
	var v misc.DxlEndpoint
	name := acctest.RandomNameWithPrefix("dxl-endpoint")
	templateInstanceVal := map[string]any{
		"template": "Version5_DXL_Session_Template",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointTemplateInstance(clientCertificateFile, broker, name, "GM", templateInstanceVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "template_instance.template", "Version5_DXL_Session_Template"),
				),
			},
			// TODO : Update Testing Requires More Templates
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccDxlEndpointResource_Timeout(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_timeout"
	var v misc.DxlEndpoint
	name := acctest.RandomNameWithPrefix("dxl-endpoint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointTimeout(clientCertificateFile, broker, name, "GM", "60"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "60"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointTimeout(clientCertificateFile, broker, name, "GM", "120"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "timeout", "120"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccDxlEndpointResource_Topics(t *testing.T) {
	t.Skip("Additional config is required to run this test")
	var resourceName = "nios_misc_dxl_endpoint.test_topics"
	var v misc.DxlEndpoint
	name := acctest.RandomNameWithPrefix("dxl-endpoint")
	topicsVal := []string{"/outbound/session", "/infoblox/outbound/LEASE"}
	topicsValUpdated := []string{"/outbound/session/updated", "/outbound/session"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointTopics(clientCertificateFile, broker, name, "GM", topicsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "topics.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "topics.0", "/outbound/session"),
					resource.TestCheckResourceAttr(resourceName, "topics.1", "/infoblox/outbound/LEASE"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointTopics(clientCertificateFile, broker, name, "GM", topicsValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "topics.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "topics.0", "/outbound/session/updated"),
					resource.TestCheckResourceAttr(resourceName, "topics.1", "/outbound/session"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_VendorIdentifier(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_vendor_identifier"
	var v misc.DxlEndpoint
	name := acctest.RandomNameWithPrefix("dxl-endpoint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointVendorIdentifier(clientCertificateFile, broker, name, "GM", "McAfee"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_identifier", "McAfee"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointVendorIdentifier(clientCertificateFile, broker, name, "GM", "testing123"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_identifier", "testing123"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDxlEndpointResource_WapiUserName(t *testing.T) {
	var resourceName = "nios_misc_dxl_endpoint.test_wapi_user_name"
	var v misc.DxlEndpoint
	name := acctest.RandomNameWithPrefix("dxl-endpoint")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDxlEndpointWapiUserName(clientCertificateFile, broker, name, "GM", "admin", "password"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_name", "admin"),
				),
			},
			// Update and Read
			{
				Config: testAccDxlEndpointWapiUserName(clientCertificateFile, broker, name, "GM", "admin_updated", "password"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDxlEndpointExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wapi_user_name", "admin_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDxlEndpointExists(ctx context.Context, resourceName string, v *misc.DxlEndpoint) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.MiscAPI.
			DxlEndpointAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForDxlEndpoint).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetDxlEndpointResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetDxlEndpointResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckDxlEndpointDestroy(ctx context.Context, v *misc.DxlEndpoint) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.MiscAPI.
			DxlEndpointAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForDxlEndpoint).
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

func testAccCheckDxlEndpointDisappears(ctx context.Context, v *misc.DxlEndpoint) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.MiscAPI.
			DxlEndpointAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDxlEndpointBasicConfig(clientCertificateToken, name, outboundMemberType string, broker []map[string]any) string {
	brokerStr := utils.ConvertSliceOfMapsToHCL(broker)
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
	brokers = %s
}
`, clientCertificateToken, name, outboundMemberType, brokerStr)
}

func testAccDxlEndpointBrokers(clientCertificateToken string, name string, outboundMemberType string, brokers []map[string]any) string {
	brokersStr := utils.ConvertSliceOfMapsToHCL(brokers)
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_brokers" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    brokers = %s
}
`, clientCertificateToken, name, outboundMemberType, brokersStr)
}

func testAccDxlEndpointBrokersImportToken(clientCertificateToken string, name string, outboundMemberType string, brokersImportFile string) string {
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_brokers_import_token" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    brokers_import_file = %q
}
`, clientCertificateToken, name, outboundMemberType, brokersImportFile)
}

func testAccDxlEndpointClientCertificateToken(clientCertificateToken string, name string, outboundMemberType string, broker []map[string]any) string {
	brokerStr := utils.ConvertSliceOfMapsToHCL(broker)
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_client_certificate_token" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    brokers = %s
}
`, clientCertificateToken, name, outboundMemberType, brokerStr)
}

func testAccDxlEndpointComment(clientCertificateToken string, broker []map[string]any, name string, outboundMemberType string, comment string) string {
	brokerStr := utils.ConvertSliceOfMapsToHCL(broker)
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_comment" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    comment = %q
	brokers = %s
}
`, clientCertificateToken, name, outboundMemberType, comment, brokerStr)
}

func testAccDxlEndpointDisable(clientCertificateToken string, broker []map[string]any, name string, outboundMemberType string, disable string) string {
	brokerStr := utils.ConvertSliceOfMapsToHCL(broker)
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_disable" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    disable = %q
	brokers = %s
}
`, clientCertificateToken, name, outboundMemberType, disable, brokerStr)
}

func testAccDxlEndpointExtAttrs(clientCertificateToken string, broker []map[string]any, name string, outboundMemberType string, extAttrs map[string]string) string {
	brokerStr := utils.ConvertSliceOfMapsToHCL(broker)
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_extattrs" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    extattrs = %s
	brokers = %s
}
`, clientCertificateToken, name, outboundMemberType, extAttrsStr, brokerStr)
}

func testAccDxlEndpointLogLevelDefault(clientCertificateToken string, broker []map[string]any, name string, outboundMemberType string) string {
	brokerStr := utils.ConvertSliceOfMapsToHCL(broker)
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_log_level" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
	brokers = %s
}
`, clientCertificateToken, name, outboundMemberType, brokerStr)
}

func testAccDxlEndpointLogLevel(clientCertificateToken string, broker []map[string]any, name string, outboundMemberType string, logLevel string) string {
	brokerStr := utils.ConvertSliceOfMapsToHCL(broker)
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_log_level" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    log_level = %q
	brokers = %s
}
`, clientCertificateToken, name, outboundMemberType, logLevel, brokerStr)
}

func testAccDxlEndpointName(clientCertificateToken string, broker []map[string]any, name string, outboundMemberType string) string {
	brokerStr := utils.ConvertSliceOfMapsToHCL(broker)
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_name" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
	brokers = %s
}
`, clientCertificateToken, name, outboundMemberType, brokerStr)
}

func testAccDxlEndpointOutboundMemberType(clientCertificateToken string, broker []map[string]any, name string, outboundMemberType string) string {
	brokerStr := utils.ConvertSliceOfMapsToHCL(broker)
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_outbound_member_type" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
	brokers = %s
}
`, clientCertificateToken, name, outboundMemberType, brokerStr)
}

func testAccDxlEndpointOutboundMemberTypeUpdate(clientCertificateToken string, broker []map[string]any, name string, outboundMemberType string, outboundMembers []string) string {
	brokerStr := utils.ConvertSliceOfMapsToHCL(broker)
	outboundMembersStr := utils.ConvertStringSliceToHCL(outboundMembers)
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_outbound_member_type" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
	brokers = %s
	outbound_members = %s
}
`, clientCertificateToken, name, outboundMemberType, brokerStr, outboundMembersStr)
}

func testAccDxlEndpointOutboundMembers(clientCertificateToken string, broker []map[string]any, name string, outboundMemberType string, outboundMembers []string) string {
	brokerStr := utils.ConvertSliceOfMapsToHCL(broker)
	outboundMembersStr := utils.ConvertStringSliceToHCL(outboundMembers)
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_outbound_members" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    outbound_members = %s
	brokers = %s
}
`, clientCertificateToken, name, outboundMemberType, outboundMembersStr, brokerStr)
}

func testAccDxlEndpointTemplateInstance(clientCertificateToken string, broker []map[string]any, name string, outboundMemberType string, templateInstance map[string]any) string {
	brokerStr := utils.ConvertSliceOfMapsToHCL(broker)
	templateInstanceStr := utils.ConvertMapToHCL(templateInstance)
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_template_instance" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    template_instance = %s
	brokers = %s
}
`, clientCertificateToken, name, outboundMemberType, templateInstanceStr, brokerStr)
}

func testAccDxlEndpointTimeout(clientCertificateToken string, broker []map[string]any, name string, outboundMemberType string, timeout string) string {
	brokerStr := utils.ConvertSliceOfMapsToHCL(broker)
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_timeout" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    timeout = %q
	brokers = %s
}
`, clientCertificateToken, name, outboundMemberType, timeout, brokerStr)
}

func testAccDxlEndpointTopics(clientCertificateToken string, broker []map[string]any, name string, outboundMemberType string, topics []string) string {
	brokerStr := utils.ConvertSliceOfMapsToHCL(broker)
	topicsStr := utils.ConvertStringSliceToHCL(topics)
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_topics" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    topics = %s
	brokers = %s
}
`, clientCertificateToken, name, outboundMemberType, topicsStr, brokerStr)
}

func testAccDxlEndpointVendorIdentifier(clientCertificateToken string, broker []map[string]any, name string, outboundMemberType string, vendorIdentifier string) string {
	brokerStr := utils.ConvertSliceOfMapsToHCL(broker)
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_vendor_identifier" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    vendor_identifier = %q
	brokers = %s
}
`, clientCertificateToken, name, outboundMemberType, vendorIdentifier, brokerStr)
}

func testAccDxlEndpointWapiUserName(clientCertificateToken string, broker []map[string]any, name string, outboundMemberType string, wapiUserName, password string) string {
	brokerStr := utils.ConvertSliceOfMapsToHCL(broker)
	return fmt.Sprintf(`
resource "nios_misc_dxl_endpoint" "test_wapi_user_name" {
    client_certificate_file = %q
    name = %q
    outbound_member_type = %q
    wapi_user_name = %q
    wapi_user_password = %q
	brokers = %s
}
`, clientCertificateToken, name, outboundMemberType, wapiUserName, password, brokerStr)
}

func getTestDataPath() string {
	wd, err := os.Getwd()
	if err != nil {
		return "../../testdata/nios_misc_dxl_endpoint"
	}
	return filepath.Join(wd, "../../testdata/nios_misc_dxl_endpoint")
}
