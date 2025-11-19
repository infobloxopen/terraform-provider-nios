package ipam_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForVlanrange = "comment,end_vlan_id,extattrs,name,pre_create_vlan,start_vlan_id,vlan_name_prefix,vlan_view"

func TestAccVlanrangeResource_basic(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test"
	var v ipam.Vlanrange

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanrangeBasicConfig("END_VLAN_ID_REPLACE_ME", "NAME_REPLACE_ME", "START_VLAN_ID_REPLACE_ME", "VLAN_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "end_vlan_id", "END_VLAN_ID_REPLACE_ME"),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
					resource.TestCheckResourceAttr(resourceName, "start_vlan_id", "START_VLAN_ID_REPLACE_ME"),
					resource.TestCheckResourceAttr(resourceName, "vlan_view", "VLAN_VIEW_REPLACE_ME"),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanrangeResource_disappears(t *testing.T) {
	resourceName := "nios_ipam_vlanrange.test"
	var v ipam.Vlanrange

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVlanrangeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccVlanrangeBasicConfig("END_VLAN_ID_REPLACE_ME", "NAME_REPLACE_ME", "START_VLAN_ID_REPLACE_ME", "VLAN_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					testAccCheckVlanrangeDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccVlanrangeResource_Import(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test"
	var v ipam.Vlanrange

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanrangeBasicConfig("END_VLAN_ID_REPLACE_ME", "NAME_REPLACE_ME", "START_VLAN_ID_REPLACE_ME", "VLAN_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
				),
			},
			// Import with PlanOnly to detect differences
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccVlanrangeImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "ref",
				PlanOnly:                             true,
				//ExpectError:                          regexp.MustCompile(`ImportStateVerify attributes not equivalent`),
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccVlanrangeImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanrangeResource_Comment(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test_comment"
	var v ipam.Vlanrange

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanrangeComment("END_VLAN_ID_REPLACE_ME", "NAME_REPLACE_ME", "START_VLAN_ID_REPLACE_ME", "VLAN_VIEW_REPLACE_ME", "Comment for the Vlanrange object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanrangeComment("END_VLAN_ID_REPLACE_ME", "NAME_REPLACE_ME", "START_VLAN_ID_REPLACE_ME", "VLAN_VIEW_REPLACE_ME", "Updated comment for the Vlanrange object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the object"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanrangeResource_DeleteVlans(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test_delete_vlans"
	var v ipam.Vlanrange

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanrangeDeleteVlans("END_VLAN_ID_REPLACE_ME", "NAME_REPLACE_ME", "START_VLAN_ID_REPLACE_ME", "VLAN_VIEW_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delete_vlans", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanrangeDeleteVlans("END_VLAN_ID_REPLACE_ME", "NAME_REPLACE_ME", "START_VLAN_ID_REPLACE_ME", "VLAN_VIEW_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delete_vlans", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanrangeResource_EndVlanId(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test_end_vlan_id"
	var v ipam.Vlanrange

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanrangeEndVlanId("END_VLAN_ID_REPLACE_ME", "NAME_REPLACE_ME", "START_VLAN_ID_REPLACE_ME", "VLAN_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "end_vlan_id", "END_VLAN_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanrangeEndVlanId("END_VLAN_ID_REPLACE_ME", "NAME_REPLACE_ME", "START_VLAN_ID_REPLACE_ME", "VLAN_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "end_vlan_id", "END_VLAN_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccVlanrangeResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test_extattrs"
	var v ipam.Vlanrange
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanrangeExtAttrs("END_VLAN_ID_REPLACE_ME", "NAME_REPLACE_ME", "START_VLAN_ID_REPLACE_ME", "VLAN_VIEW_REPLACE_ME", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccVlanrangeExtAttrs("END_VLAN_ID_REPLACE_ME", "NAME_REPLACE_ME", "START_VLAN_ID_REPLACE_ME", "VLAN_VIEW_REPLACE_ME", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanrangeResource_Name(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test_name"
	var v ipam.Vlanrange

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanrangeName("END_VLAN_ID_REPLACE_ME", "NAME_REPLACE_ME", "START_VLAN_ID_REPLACE_ME", "VLAN_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanrangeName("END_VLAN_ID_REPLACE_ME", "NAME_REPLACE_ME", "START_VLAN_ID_REPLACE_ME", "VLAN_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanrangeResource_PreCreateVlan(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test_pre_create_vlan"
	var v ipam.Vlanrange

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanrangePreCreateVlan("END_VLAN_ID_REPLACE_ME", "NAME_REPLACE_ME", "START_VLAN_ID_REPLACE_ME", "VLAN_VIEW_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pre_create_vlan", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanrangePreCreateVlan("END_VLAN_ID_REPLACE_ME", "NAME_REPLACE_ME", "START_VLAN_ID_REPLACE_ME", "VLAN_VIEW_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pre_create_vlan", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanrangeResource_StartVlanId(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test_start_vlan_id"
	var v ipam.Vlanrange

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanrangeStartVlanId("END_VLAN_ID_REPLACE_ME", "NAME_REPLACE_ME", "START_VLAN_ID_REPLACE_ME", "VLAN_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_vlan_id", "START_VLAN_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanrangeStartVlanId("END_VLAN_ID_REPLACE_ME", "NAME_REPLACE_ME", "START_VLAN_ID_REPLACE_ME", "VLAN_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_vlan_id", "START_VLAN_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccVlanrangeResource_VlanNamePrefix(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test_vlan_name_prefix"
	var v ipam.Vlanrange

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanrangeVlanNamePrefix("END_VLAN_ID_REPLACE_ME", "NAME_REPLACE_ME", "START_VLAN_ID_REPLACE_ME", "VLAN_VIEW_REPLACE_ME", "VLAN_NAME_PREFIX_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vlan_name_prefix", "VLAN_NAME_PREFIX_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanrangeVlanNamePrefix("END_VLAN_ID_REPLACE_ME", "NAME_REPLACE_ME", "START_VLAN_ID_REPLACE_ME", "VLAN_VIEW_REPLACE_ME", "VLAN_NAME_PREFIX_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vlan_name_prefix", "VLAN_NAME_PREFIX_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanrangeResource_VlanView(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test_vlan_view"
	var v ipam.Vlanrange

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanrangeVlanView("END_VLAN_ID_REPLACE_ME", "NAME_REPLACE_ME", "START_VLAN_ID_REPLACE_ME", "VLAN_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vlan_view", "VLAN_VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanrangeVlanView("END_VLAN_ID_REPLACE_ME", "NAME_REPLACE_ME", "START_VLAN_ID_REPLACE_ME", "VLAN_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vlan_view", "VLAN_VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckVlanrangeExists(ctx context.Context, resourceName string, v *ipam.Vlanrange) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.IPAMAPI.
			VlanrangeAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForVlanrange).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetVlanrangeResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetVlanrangeResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckVlanrangeDestroy(ctx context.Context, v *ipam.Vlanrange) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.IPAMAPI.
			VlanrangeAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForVlanrange).
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

func testAccCheckVlanrangeDisappears(ctx context.Context, v *ipam.Vlanrange) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.IPAMAPI.
			VlanrangeAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccVlanrangeImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.Attributes["ref"] == "" {
			return "", fmt.Errorf("ref is not set")
		}
		return rs.Primary.Attributes["ref"], nil
	}
}

func testAccVlanrangeBasicConfig(endVlanId, name, startVlanId, vlanView string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanrange" "test" {
    end_vlan_id = %q
    name = %q
    start_vlan_id = %q
    vlan_view = %q
}
`, endVlanId, name, startVlanId, vlanView)
}

func testAccVlanrangeComment(endVlanId string, name string, startVlanId string, vlanView string, comment string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanrange" "test_comment" {
    end_vlan_id = %q
    name = %q
    start_vlan_id = %q
    vlan_view = %q
    comment = %q
}
`, endVlanId, name, startVlanId, vlanView, comment)
}

func testAccVlanrangeDeleteVlans(endVlanId string, name string, startVlanId string, vlanView string, deleteVlans string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanrange" "test_delete_vlans" {
    end_vlan_id = %q
    name = %q
    start_vlan_id = %q
    vlan_view = %q
    delete_vlans = %q
}
`, endVlanId, name, startVlanId, vlanView, deleteVlans)
}

func testAccVlanrangeEndVlanId(endVlanId string, name string, startVlanId string, vlanView string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanrange" "test_end_vlan_id" {
    end_vlan_id = %q
    name = %q
    start_vlan_id = %q
    vlan_view = %q
}
`, endVlanId, name, startVlanId, vlanView)
}

func testAccVlanrangeExtAttrs(endVlanId string, name string, startVlanId string, vlanView string, extAttrs map[string]string) string {
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	return fmt.Sprintf(`
resource "nios_ipam_vlanrange" "test_extattrs" {
    end_vlan_id = %q
    name = %q
    start_vlan_id = %q
    vlan_view = %q
    extattrs = %q
}
`, endVlanId, name, startVlanId, vlanView, extAttrsStr)
}

func testAccVlanrangeName(endVlanId string, name string, startVlanId string, vlanView string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanrange" "test_name" {
    end_vlan_id = %q
    name = %q
    start_vlan_id = %q
    vlan_view = %q
}
`, endVlanId, name, startVlanId, vlanView)
}

func testAccVlanrangePreCreateVlan(endVlanId string, name string, startVlanId string, vlanView string, preCreateVlan string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanrange" "test_pre_create_vlan" {
    end_vlan_id = %q
    name = %q
    start_vlan_id = %q
    vlan_view = %q
    pre_create_vlan = %q
}
`, endVlanId, name, startVlanId, vlanView, preCreateVlan)
}

func testAccVlanrangeStartVlanId(endVlanId string, name string, startVlanId string, vlanView string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanrange" "test_start_vlan_id" {
    end_vlan_id = %q
    name = %q
    start_vlan_id = %q
    vlan_view = %q
}
`, endVlanId, name, startVlanId, vlanView)
}

func testAccVlanrangeVlanNamePrefix(endVlanId string, name string, startVlanId string, vlanView string, vlanNamePrefix string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanrange" "test_vlan_name_prefix" {
    end_vlan_id = %q
    name = %q
    start_vlan_id = %q
    vlan_view = %q
    vlan_name_prefix = %q
}
`, endVlanId, name, startVlanId, vlanView, vlanNamePrefix)
}

func testAccVlanrangeVlanView(endVlanId string, name string, startVlanId string, vlanView string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanrange" "test_vlan_view" {
    end_vlan_id = %q
    name = %q
    start_vlan_id = %q
    vlan_view = %q
}
`, endVlanId, name, startVlanId, vlanView)
}
