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

var readableAttributesForVlanview = "allow_range_overlapping,comment,end_vlan_id,extattrs,name,pre_create_vlan,start_vlan_id,vlan_name_prefix"

func TestAccVlanviewResource_basic(t *testing.T) {
	var resourceName = "nios_ipam_vlanview.test"
	var v ipam.Vlanview
	name := acctest.RandomNameWithPrefix("vlan_view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanviewBasicConfig(15, name, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "end_vlan_id", "15"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "start_vlan_id", "10"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "allow_range_overlapping", "false"),
					resource.TestCheckResourceAttr(resourceName, "pre_create_vlan", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanviewResource_disappears(t *testing.T) {
	resourceName := "nios_ipam_vlanview.test"
	var v ipam.Vlanview
	name := acctest.RandomNameWithPrefix("vlan_view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVlanviewDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccVlanviewBasicConfig(15, name, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					testAccCheckVlanviewDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccVlanviewResource_Import(t *testing.T) {
	var resourceName = "nios_ipam_vlanview.test"
	var v ipam.Vlanview
	name := acctest.RandomNameWithPrefix("vlan_view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanviewBasicConfig(15, name, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
				),
			},
			// Import with PlanOnly to detect differences
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccVlanviewImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "ref",
				PlanOnly:                             true,
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccVlanviewImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanviewResource_AllowRangeOverlapping(t *testing.T) {
	var resourceName = "nios_ipam_vlanview.test_allow_range_overlapping"
	var v ipam.Vlanview
	name := acctest.RandomNameWithPrefix("vlan_view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanviewAllowRangeOverlapping(15, name, 10, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_range_overlapping", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanviewAllowRangeOverlapping(15, name, 10, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_range_overlapping", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccVlanviewResource_Comment(t *testing.T) {
	var resourceName = "nios_ipam_vlanview.test_comment"
	var v ipam.Vlanview
	name := acctest.RandomNameWithPrefix("vlan_view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanviewComment(15, name, 10, "Comment for the Vlan view object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the Vlan view object"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanviewComment(15, name, 10, "Updated comment for the Vlan view object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the Vlan view object"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanviewResource_EndVlanId(t *testing.T) {
	var resourceName = "nios_ipam_vlanview.test_end_vlan_id"
	var v ipam.Vlanview
	name := acctest.RandomNameWithPrefix("vlan_view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanviewEndVlanId(4094, name, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "end_vlan_id", "4094"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanviewEndVlanId(1, name, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "end_vlan_id", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccVlanviewResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_ipam_vlanview.test_extattrs"
	var v ipam.Vlanview
	name := acctest.RandomNameWithPrefix("vlan_view")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanviewExtAttrs(15, name, 10, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccVlanviewExtAttrs(15, name, 10, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccVlanviewResource_Name(t *testing.T) {
	var resourceName = "nios_ipam_vlanview.test_name"
	var v ipam.Vlanview
	name := acctest.RandomNameWithPrefix("vlan_view")
	name2 := acctest.RandomNameWithPrefix("vlan_view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanviewName(15, name, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccVlanviewName(15, name2, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccVlanviewResource_PreCreateVlan(t *testing.T) {
	t.Skip("Skipping test as WAPI does not support updating pre_create_vlan attribute")
	var resourceName = "nios_ipam_vlanview.test_pre_create_vlan"
	var v ipam.Vlanview
	name := acctest.RandomNameWithPrefix("vlan_view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanviewPreCreateVlan(15, name, 10, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pre_create_vlan", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccVlanviewResource_StartVlanId(t *testing.T) {
	var resourceName = "nios_ipam_vlanview.test_start_vlan_id"
	var v ipam.Vlanview
	name := acctest.RandomNameWithPrefix("vlan_view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanviewStartVlanId(4094, name, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_vlan_id", "1"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanviewStartVlanId(4094, name, 4094),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_vlan_id", "4094"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccVlanviewResource_VlanNamePrefix(t *testing.T) {
	t.Skip("WAPI Doesnt allow updating vlan_name_prefix field")
	var resourceName = "nios_ipam_vlanview.test_vlan_name_prefix"
	var v ipam.Vlanview
	name := acctest.RandomNameWithPrefix("vlan_view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanviewVlanNamePrefix(15, name, 10, "prefixCaseInsensitive"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vlan_name_prefix", "prefixCaseInsensitive"),
				),
			},
			// Update and check for empty string
			{
				Config: testAccVlanviewVlanNamePrefix(15, name, 10, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vlan_name_prefix", ""),
				),
			},
			// Update and Read
			{
				Config: testAccVlanviewVlanNamePrefix(15, name, 10, "VLAN_NAME_PREFIX_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vlan_name_prefix", "VLAN_NAME_PREFIX_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckVlanviewExists(ctx context.Context, resourceName string, v *ipam.Vlanview) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.IPAMAPI.
			VlanviewAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForVlanview).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetVlanviewResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetVlanviewResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckVlanviewDestroy(ctx context.Context, v *ipam.Vlanview) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.IPAMAPI.
			VlanviewAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForVlanview).
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

func testAccCheckVlanviewDisappears(ctx context.Context, v *ipam.Vlanview) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.IPAMAPI.
			VlanviewAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccVlanviewImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccVlanviewBasicConfig(endVlanId int, name string, startVlanId int) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "test" {
    end_vlan_id = %d
    name = %q
    start_vlan_id = %d
}
`, endVlanId, name, startVlanId)
}

func testAccVlanviewAllowRangeOverlapping(endVlanId int, name string, startVlanId int, allowRangeOverlapping string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "test_allow_range_overlapping" {
    end_vlan_id = %d
    name = %q
    start_vlan_id = %d
    allow_range_overlapping = %s
}
`, endVlanId, name, startVlanId, allowRangeOverlapping)
}

func testAccVlanviewComment(endVlanId int, name string, startVlanId int, comment string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "test_comment" {
    end_vlan_id = %d
    name = %q
    start_vlan_id = %d
    comment = %q
}
`, endVlanId, name, startVlanId, comment)
}

func testAccVlanviewEndVlanId(endVlanId int, name string, startVlanId int) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "test_end_vlan_id" {
    end_vlan_id = %d
    name = %q
    start_vlan_id = %d
}
`, endVlanId, name, startVlanId)
}

func testAccVlanviewExtAttrs(endVlanId int, name string, startVlanId int, extAttrs map[string]string) string {
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "test_extattrs" {
    end_vlan_id = %d
    name = %q
    start_vlan_id = %d
    extattrs = %s
}
`, endVlanId, name, startVlanId, extAttrsStr)
}

func testAccVlanviewName(endVlanId int, name string, startVlanId int) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "test_name" {
    end_vlan_id = %d
    name = %q
    start_vlan_id = %d
}
`, endVlanId, name, startVlanId)
}

func testAccVlanviewPreCreateVlan(endVlanId int, name string, startVlanId int, preCreateVlan string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "test_pre_create_vlan" {
    end_vlan_id = %d
    name = %q
    start_vlan_id = %d
    pre_create_vlan = %s
}
`, endVlanId, name, startVlanId, preCreateVlan)
}

func testAccVlanviewStartVlanId(endVlanId int, name string, startVlanId int) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "test_start_vlan_id" {
    end_vlan_id = %d
    name = %q
    start_vlan_id = %d
}
`, endVlanId, name, startVlanId)
}

func testAccVlanviewVlanNamePrefix(endVlanId int, name string, startVlanId int, vlanNamePrefix string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "test_vlan_name_prefix" {
    end_vlan_id = %d
    name = %q
    start_vlan_id = %d
    vlan_name_prefix = %q
	pre_create_vlan = true
}
`, endVlanId, name, startVlanId, vlanNamePrefix)
}
