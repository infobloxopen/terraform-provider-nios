package dhcp_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForFilterrelayagent = "circuit_id_name,circuit_id_substring_length,circuit_id_substring_offset,comment,extattrs,is_circuit_id,is_circuit_id_substring,is_remote_id,is_remote_id_substring,name,remote_id_name,remote_id_substring_length,remote_id_substring_offset"

func TestAccFilterrelayagentResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_filterrelayagent.test"
	var v dhcp.Filterrelayagent
	name := "TEST_FILTERRELAYAGENT_RESOURCE_BASIC"
	isCircuitID := "NOT_SET"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilterrelayagentBasicConfig(name, isCircuitID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "is_circuit_id", isCircuitID),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "is_remote_id", "ANY"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilterrelayagentResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_filterrelayagent.test"
	var v dhcp.Filterrelayagent
	name := "TEST_FILTERRELAYAGENT_RESOURCE_DISAPPEARS"
	isCircuitID := "NOT_SET"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFilterrelayagentDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFilterrelayagentBasicConfig(name, isCircuitID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					testAccCheckFilterrelayagentDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccFilterrelayagentResource_CircuitIdName(t *testing.T) {
	var resourceName = "nios_dhcp_filterrelayagent.test_circuit_id_name"
	var v dhcp.Filterrelayagent
	name := acctest.RandomNameWithPrefix("filterrelayagent")
	cicuitIdName := "CIRCUIT_ID_NAME_TEST"
	isCircuitId := "MATCHES_VALUE"
	updateCircuitIdName := "CIRCUIT_ID_NAME_TEST_UPDATED"
	isCircuitIdSubstring := "false"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilterrelayagentCircuitIdName(name, cicuitIdName, isCircuitId, isCircuitIdSubstring),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "circuit_id_name", cicuitIdName),
				),
			},
			// Update and Read
			{
				Config: testAccFilterrelayagentCircuitIdName(name, updateCircuitIdName, isCircuitId, isCircuitIdSubstring),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "circuit_id_name", updateCircuitIdName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilterrelayagentResource_CircuitIdSubstringLength(t *testing.T) {
	var resourceName = "nios_dhcp_filterrelayagent.test_circuit_id_substring_length"
	var v dhcp.Filterrelayagent
	name := acctest.RandomNameWithPrefix("filterrelayagent")
	isCircuitID := "MATCHES_VALUE"
	isCircuitIdSubstring := "true"
	circuitIdSubstringOffset := "0"
	circuitIdName := "CIRCUIT_ID_NAME_TEST"
	circuitIdSubstringLength := fmt.Sprintf("%d", len(circuitIdName))
	updatedCircuitIDName := "CIRCUIT_ID_NAME_TEST_UPDATED"
	updatedCircuitIdSubstringLength := fmt.Sprintf("%d", len(updatedCircuitIDName))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilterrelayagentCircuitIdSubstringLength(name, isCircuitID, circuitIdName, circuitIdSubstringLength, circuitIdSubstringOffset, isCircuitIdSubstring),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "circuit_id_substring_length", circuitIdSubstringLength),
				),
			},
			// Update and Read
			{
				Config: testAccFilterrelayagentCircuitIdSubstringLength(name, isCircuitID, updatedCircuitIDName, updatedCircuitIdSubstringLength, circuitIdSubstringOffset, isCircuitIdSubstring),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "circuit_id_substring_length", updatedCircuitIdSubstringLength),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilterrelayagentResource_CircuitIdSubstringOffset(t *testing.T) {
	var resourceName = "nios_dhcp_filterrelayagent.test_circuit_id_substring_offset"
	var v dhcp.Filterrelayagent
	name := acctest.RandomNameWithPrefix("filterrelayagent")
	isCircuitID := "MATCHES_VALUE"
	isCircuitIdSubstring := "true"
	circuitIdSubstringOffset := "0"
	circuitIdName := "CIRCUIT_ID_NAME_TEST"
	isCircuitIdSubstringLength := fmt.Sprintf("%d", len(circuitIdName))
	updatedCircuitIdSubstringOffset := "1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilterrelayagentCircuitIdSubstringOffset(name, isCircuitID, circuitIdName, isCircuitIdSubstringLength, circuitIdSubstringOffset, isCircuitIdSubstring),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "circuit_id_substring_offset", "0"),
				),
			},
			// Update and Read
			{
				Config: testAccFilterrelayagentCircuitIdSubstringOffset(name, isCircuitID, circuitIdName, isCircuitIdSubstringLength, updatedCircuitIdSubstringOffset, isCircuitIdSubstring),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "circuit_id_substring_offset", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilterrelayagentResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_filterrelayagent.test_comment"
	var v dhcp.Filterrelayagent
	name := acctest.RandomNameWithPrefix("filterrelayagent")
	isCircuitID := "NOT_SET"
	comment := "COMMENT_TEST"
	updatedComment := "COMMENT_TEST_UPDATED"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilterrelayagentComment(name, isCircuitID, comment),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_TEST"),
				),
			},
			// Update and Read
			{
				Config: testAccFilterrelayagentComment(name, isCircuitID, updatedComment),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_TEST_UPDATED"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilterrelayagentResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dhcp_filterrelayagent.test_extattrs"
	var v dhcp.Filterrelayagent
	name := acctest.RandomNameWithPrefix("filterrelayagent")
	isCircuitId := "NOT_SET"
	siteExtAttr := acctest.RandomName()
	updatedSiteExtAttr := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilterrelayagentExtAttrs(name, isCircuitId, map[string]string{"Site": siteExtAttr}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", siteExtAttr),
				),
			},
			// Update and Read
			{
				Config: testAccFilterrelayagentExtAttrs(name, isCircuitId, map[string]string{"Site": updatedSiteExtAttr}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", updatedSiteExtAttr),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilterrelayagentResource_IsCircuitId(t *testing.T) {
	var resourceName = "nios_dhcp_filterrelayagent.test_is_circuit_id"
	var v dhcp.Filterrelayagent
	name := acctest.RandomNameWithPrefix("filterrelayagent")
	isCircuitId := "NOT_SET"
	updatedIsCircuitId := "MATCHES_VALUE"
	circuitIDName := "CIRCUIT_ID_NAME_TEST"
	isCircuitIdSubstring := "false"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilterrelayagentIsCircuitId(name, isCircuitId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "is_circuit_id", "NOT_SET"),
				),
			},
			// Update and Read
			{
				Config: testAccFilterrelayagentIsCircuitIdUpdated(name, updatedIsCircuitId, circuitIDName, isCircuitIdSubstring),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "is_circuit_id", "MATCHES_VALUE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilterrelayagentResource_IsCircuitIdSubstring(t *testing.T) {
	var resourceName = "nios_dhcp_filterrelayagent.test_is_circuit_id_substring"
	var v dhcp.Filterrelayagent
	name := acctest.RandomNameWithPrefix("filterrelayagent")
	isCircuitId := "MATCHES_VALUE"
	circuitIDName := "CIRCUIT_ID_NAME_TEST"
	isCircuitIdSubstring := "true"
	substringoffset := "0"
	isCircuitIdSubstringUpdate := "false"
	substringlength := fmt.Sprintf("%d", len(circuitIDName))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilterrelayagentIsCircuitIdSubstring(name, isCircuitId, circuitIDName, substringlength, substringoffset, isCircuitIdSubstring),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "is_circuit_id_substring", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFilterrelayagentIsCircuitIdSubstringUpdated(name, isCircuitId, circuitIDName, isCircuitIdSubstringUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "is_circuit_id_substring", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilterrelayagentResource_IsRemoteId(t *testing.T) {
	var resourceName = "nios_dhcp_filterrelayagent.test_is_remote_id"
	var v dhcp.Filterrelayagent
	isRemoteID := "MATCHES_VALUE"
	isRemoteIDUpdate := "NOT_SET"
	name := acctest.RandomNameWithPrefix("filterrelayagent")
	remoteName := "REMOTE_ID_NAME_TEST"
	isRemoteIdSubstring := "false"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilterrelayagentIsRemoteId(name, isRemoteID, remoteName, isRemoteIdSubstring),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "is_remote_id", "MATCHES_VALUE"),
				),
			},
			// Update and Read
			{
				Config: testAccFilterrelayagentIsRemoteIdUpdate(name, isRemoteIDUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "is_remote_id", "NOT_SET"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilterrelayagentResource_IsRemoteIdSubstring(t *testing.T) {
	var resourceName = "nios_dhcp_filterrelayagent.test_is_remote_id_substring"
	var v dhcp.Filterrelayagent
	name := acctest.RandomNameWithPrefix("filterrelayagent")
	isRemoteIdSubstring := "true"
	updatedIsRemoteIdSubstring := "false"
	isRemoteId := "MATCHES_VALUE"
	remoteName := "REMOTE_ID_NAME_TEST"
	remoteIdSubstringLength := fmt.Sprintf("%d", len(remoteName))
	remoteOffset := "0"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilterrelayagentIsRemoteIdSubstring(name, isRemoteId, remoteName, remoteIdSubstringLength, remoteOffset, isRemoteIdSubstring),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "is_remote_id_substring", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFilterrelayagentIsRemoteIdSubstringUpdated(name, isRemoteId, remoteName, updatedIsRemoteIdSubstring),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "is_remote_id_substring", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilterrelayagentResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_filterrelayagent.test_name"
	var v dhcp.Filterrelayagent
	name := "NAME_REPLACE_ME"
	updateName := "NAME_UPDATE_REPLACE_ME"
	isCircuitId := "NOT_SET"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilterrelayagentName(name, isCircuitId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFilterrelayagentName(updateName, isCircuitId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilterrelayagentResource_RemoteIdName(t *testing.T) {
	var resourceName = "nios_dhcp_filterrelayagent.test_remote_id_name"
	var v dhcp.Filterrelayagent
	name := acctest.RandomNameWithPrefix("filterrelayagent")
	isRemoteId := "MATCHES_VALUE"
	isRemoteIdSubstring := "false"
	remoteIdName := "REMOTE_ID_NAME_REPLACE_ME"
	updateRemoteIdName := "REMOTE_ID_NAME_UPDATE_REPLACE_ME"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilterrelayagentRemoteIdName(name, isRemoteId, remoteIdName, isRemoteIdSubstring),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_id_name", "REMOTE_ID_NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFilterrelayagentRemoteIdName(name, isRemoteId, updateRemoteIdName, isRemoteIdSubstring),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_id_name", "REMOTE_ID_NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilterrelayagentResource_RemoteIdSubstringLength(t *testing.T) {
	var resourceName = "nios_dhcp_filterrelayagent.test_remote_id_substring_length"
	var v dhcp.Filterrelayagent
	remoteIdName := "REMOTE_ID_NAME_TEST"
	name := acctest.RandomNameWithPrefix("filterrelayagent")
	isRemoteID := "MATCHES_VALUE"
	isRemoteIdSubstring := "true"
	updateRemoteIdName := "REMOTE_ID_NAME_TEST_UPDATED"
	remoteIdSubstringLength := fmt.Sprintf("%d", len(remoteIdName))
	updatedRemoteIdSubstringLength := fmt.Sprintf("%d", len(updateRemoteIdName))
	remoteIdSubstringOffset := "0"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilterrelayagentRemoteIdSubstringLength(name, isRemoteID, remoteIdName, remoteIdSubstringLength, remoteIdSubstringOffset, isRemoteIdSubstring),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_id_substring_length", remoteIdSubstringLength),
				),
			},
			// Update and Read
			{
				Config: testAccFilterrelayagentRemoteIdSubstringLength(name, isRemoteID, updateRemoteIdName, updatedRemoteIdSubstringLength, remoteIdSubstringOffset, isRemoteIdSubstring),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_id_substring_length", updatedRemoteIdSubstringLength),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFilterrelayagentResource_RemoteIdSubstringOffset(t *testing.T) {
	var resourceName = "nios_dhcp_filterrelayagent.test_remote_id_substring_offset"
	var v dhcp.Filterrelayagent
	name := acctest.RandomNameWithPrefix("filterrelayagent")
	isRemoteID := "MATCHES_VALUE"
	remoteIdName := "REMOTE_ID_NAME_TEST"
	remoteIdSubstringLength := fmt.Sprintf("%d", len(remoteIdName))
	isRemoteIdSubstring := "true"
	remoteIdSubstringOffset := "0"
	updatedRemoteIdSubstringOffset := "1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFilterrelayagentRemoteIdSubstringOffset(name, isRemoteID, remoteIdName, remoteIdSubstringLength, remoteIdSubstringOffset, isRemoteIdSubstring),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_id_substring_offset", remoteIdSubstringOffset),
				),
			},
			// Update and Read
			{
				Config: testAccFilterrelayagentRemoteIdSubstringOffset(name, isRemoteID, remoteIdName, remoteIdSubstringLength, updatedRemoteIdSubstringOffset, isRemoteIdSubstring),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterrelayagentExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_id_substring_offset", updatedRemoteIdSubstringOffset),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckFilterrelayagentExists(ctx context.Context, resourceName string, v *dhcp.Filterrelayagent) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			FilterrelayagentAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForFilterrelayagent).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetFilterrelayagentResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetFilterrelayagentResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckFilterrelayagentDestroy(ctx context.Context, v *dhcp.Filterrelayagent) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			FilterrelayagentAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForFilterrelayagent).
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

func testAccCheckFilterrelayagentDisappears(ctx context.Context, v *dhcp.Filterrelayagent) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			FilterrelayagentAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccFilterrelayagentBasicConfig(name string, isCircuitID string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dhcp_filterrelayagent" "test" {
    name = %q
	is_circuit_id = %q
}
`, name, isCircuitID)
}

func testAccFilterrelayagentCircuitIdName(name string, circuitIdName string, isCircuitId string, isCircuitIdSubstring string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filterrelayagent" "test_circuit_id_name" {
    name = %q
    circuit_id_name = %q
    is_circuit_id = %q
	is_circuit_id_substring = %q
}
`, name, circuitIdName, isCircuitId, isCircuitIdSubstring)
}

func testAccFilterrelayagentCircuitIdSubstringLength(name, isCircuitID, circuitIdName string, circuitIdSubstringLength string, circuit_id_substring_offset string, isCircuitIdSubstring string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filterrelayagent" "test_circuit_id_substring_length" {
    name = %q
	is_circuit_id = %q
    circuit_id_name = %q
    circuit_id_substring_length = %q
	circuit_id_substring_offset = %q
    is_circuit_id_substring = %q
}
`, name, isCircuitID, circuitIdName, circuitIdSubstringLength, circuit_id_substring_offset, isCircuitIdSubstring)
}

func testAccFilterrelayagentCircuitIdSubstringOffset(name, isCircuitID, circuitIdName string, circuitIdSubstringLength string, circuit_id_substring_offset string, isCircuitIdSubstring string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filterrelayagent" "test_circuit_id_substring_offset" {
	name = %q
	is_circuit_id = %q
	circuit_id_name = %q
	circuit_id_substring_length = %q
	circuit_id_substring_offset = %q
	is_circuit_id_substring = %q
}
`, name, isCircuitID, circuitIdName, circuitIdSubstringLength, circuit_id_substring_offset, isCircuitIdSubstring)
}

func testAccFilterrelayagentComment(name, isCircuitID, comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filterrelayagent" "test_comment" {
    name = %q
	is_circuit_id = %q
    comment = %q
}
`, name, isCircuitID, comment)
}

func testAccFilterrelayagentExtAttrs(name, isCircuitId string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf("  %s = %q\n", k, v)
	}
	extattrsStr += "}"
	return fmt.Sprintf(`
resource "nios_dhcp_filterrelayagent" "test_extattrs" {
	name = %q
	is_circuit_id = %q
    extattrs = %s
}
`, name, isCircuitId, extattrsStr)
}

func testAccFilterrelayagentIsCircuitId(name, isCircuitId string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filterrelayagent" "test_is_circuit_id" {
    name = %q
    is_circuit_id = %q
}
`, name, isCircuitId)
}

func testAccFilterrelayagentIsCircuitIdUpdated(name, isCircuitId string, circuitIDName string, isCircuitIdSubstring string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filterrelayagent" "test_is_circuit_id" {
    name = %q
    is_circuit_id = %q
	circuit_id_name = %q
	is_circuit_id_substring = %q
}
`, name, isCircuitId, circuitIDName, isCircuitIdSubstring)
}

func testAccFilterrelayagentIsCircuitIdSubstring(name string, isCircuitID string, circuitName string, substringLength string, substringOffset string, isCircuitIdSubstring string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filterrelayagent" "test_is_circuit_id_substring" {
    name = %q
	is_circuit_id = %q
	circuit_id_name = %q
	circuit_id_substring_length = %q
	circuit_id_substring_offset = %q
    is_circuit_id_substring = %q
}
`, name, isCircuitID, circuitName, substringLength, substringOffset, isCircuitIdSubstring)
}

func testAccFilterrelayagentIsCircuitIdSubstringUpdated(name string, isCircuitID string, circuitName string, isCircuitIdSubstring string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filterrelayagent" "test_is_circuit_id_substring" {
    name = %q
	is_circuit_id = %q
	circuit_id_name = %q
	is_circuit_id_substring = %q
}
`, name, isCircuitID, circuitName, isCircuitIdSubstring)
}

func testAccFilterrelayagentIsRemoteId(name string, isRemoteId string, remoteName string, isRemoteIdSubstring string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filterrelayagent" "test_is_remote_id" {
    name = %q
	is_remote_id = %q
	remote_id_name = %q
	is_remote_id_substring = %q
}
`, name, isRemoteId, remoteName, isRemoteIdSubstring)
}

func testAccFilterrelayagentIsRemoteIdUpdate(name string, isRemoteId string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filterrelayagent" "test_is_remote_id" {
	name = %q
    is_remote_id = %q
}
`, name, isRemoteId)
}

func testAccFilterrelayagentIsRemoteIdSubstring(name string, isRemoteId string, remoteName string, remoteIdSubstringLength string, remoteOffset string, isRemoteIDSubstring string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filterrelayagent" "test_is_remote_id_substring" {
	name = %q
	is_remote_id = %q
	remote_id_name = %q
	remote_id_substring_length = %q
	remote_id_substring_offset = %q
	is_remote_id_substring = %q
}
`, name, isRemoteId, remoteName, remoteIdSubstringLength, remoteOffset, isRemoteIDSubstring)
}

func testAccFilterrelayagentIsRemoteIdSubstringUpdated(name string, isRemoteId string, remoteName string, isRemoteIDSubstring string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filterrelayagent" "test_is_remote_id_substring" {
	name = %q
	is_remote_id = %q
	remote_id_name = %q
	is_remote_id_substring = %q
}
`, name, isRemoteId, remoteName, isRemoteIDSubstring)
}

func testAccFilterrelayagentName(name string, isCircuitId string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filterrelayagent" "test_name" {
    name = %q
	is_circuit_id = %q
}
`, name, isCircuitId)
}

func testAccFilterrelayagentRemoteIdName(name string, isRemoteId string, remoteIdName string, isRemoteIdSubstring string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filterrelayagent" "test_remote_id_name" {
    name = %q
	is_remote_id = %q
    remote_id_name = %q
	is_remote_id_substring = %q
}
`, name, isRemoteId, remoteIdName, isRemoteIdSubstring)
}

func testAccFilterrelayagentRemoteIdSubstringLength(name string, isRemoteID string, remoteIdName string, remoteIdSubstringLength string, remoteIdSubstringOffset string, isRemoteIdSubstring string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filterrelayagent" "test_remote_id_substring_length" {
    name = %q
	is_remote_id = %q
	remote_id_name = %q
	remote_id_substring_length = %q
	remote_id_substring_offset = %q
	is_remote_id_substring = %q
}
`, name, isRemoteID, remoteIdName, remoteIdSubstringLength, remoteIdSubstringOffset, isRemoteIdSubstring)
}

func testAccFilterrelayagentRemoteIdSubstringOffset(name string, isRemoteID string, remoteIdName string, remoteIdSubstringLength string, remoteIdSubstringOffset string, isRemoteIdSubstring string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filterrelayagent" "test_remote_id_substring_offset" {
    name = %q
	is_remote_id = %q
	remote_id_name = %q
	remote_id_substring_length = %q
	remote_id_substring_offset = %q
	is_remote_id_substring = %q
}
`, name, isRemoteID, remoteIdName, remoteIdSubstringLength, remoteIdSubstringOffset, isRemoteIdSubstring)
}
