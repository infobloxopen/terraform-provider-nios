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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanviewBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanviewResource_disappears(t *testing.T) {
	resourceName := "nios_ipam_vlanview.test"
	var v ipam.Vlanview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVlanviewDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccVlanviewBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					testAccCheckVlanviewDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccVlanviewResource_Ref(t *testing.T) {
	var resourceName = "nios_ipam_vlanview.test_ref"
	var v ipam.Vlanview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanviewRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanviewRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanviewResource_AllowRangeOverlapping(t *testing.T) {
	var resourceName = "nios_ipam_vlanview.test_allow_range_overlapping"
	var v ipam.Vlanview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanviewAllowRangeOverlapping("ALLOW_RANGE_OVERLAPPING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_range_overlapping", "ALLOW_RANGE_OVERLAPPING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanviewAllowRangeOverlapping("ALLOW_RANGE_OVERLAPPING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_range_overlapping", "ALLOW_RANGE_OVERLAPPING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanviewResource_Comment(t *testing.T) {
	var resourceName = "nios_ipam_vlanview.test_comment"
	var v ipam.Vlanview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanviewComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanviewComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanviewResource_EndVlanId(t *testing.T) {
	var resourceName = "nios_ipam_vlanview.test_end_vlan_id"
	var v ipam.Vlanview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanviewEndVlanId("END_VLAN_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "end_vlan_id", "END_VLAN_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanviewEndVlanId("END_VLAN_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "end_vlan_id", "END_VLAN_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanviewResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_ipam_vlanview.test_extattrs"
	var v ipam.Vlanview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanviewExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanviewExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanviewResource_Name(t *testing.T) {
	var resourceName = "nios_ipam_vlanview.test_name"
	var v ipam.Vlanview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanviewName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanviewName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanviewResource_PreCreateVlan(t *testing.T) {
	var resourceName = "nios_ipam_vlanview.test_pre_create_vlan"
	var v ipam.Vlanview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanviewPreCreateVlan("PRE_CREATE_VLAN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pre_create_vlan", "PRE_CREATE_VLAN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanviewPreCreateVlan("PRE_CREATE_VLAN_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pre_create_vlan", "PRE_CREATE_VLAN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanviewResource_StartVlanId(t *testing.T) {
	var resourceName = "nios_ipam_vlanview.test_start_vlan_id"
	var v ipam.Vlanview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanviewStartVlanId("START_VLAN_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_vlan_id", "START_VLAN_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanviewStartVlanId("START_VLAN_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_vlan_id", "START_VLAN_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanviewResource_VlanNamePrefix(t *testing.T) {
	var resourceName = "nios_ipam_vlanview.test_vlan_name_prefix"
	var v ipam.Vlanview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanviewVlanNamePrefix("VLAN_NAME_PREFIX_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vlan_name_prefix", "VLAN_NAME_PREFIX_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanviewVlanNamePrefix("VLAN_NAME_PREFIX_UPDATE_REPLACE_ME"),
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

func testAccVlanviewBasicConfig() string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "test" {
}
`)
}

func testAccVlanviewRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccVlanviewAllowRangeOverlapping(allowRangeOverlapping string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "test_allow_range_overlapping" {
    allow_range_overlapping = %q
}
`, allowRangeOverlapping)
}

func testAccVlanviewComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccVlanviewEndVlanId(endVlanId string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "test_end_vlan_id" {
    end_vlan_id = %q
}
`, endVlanId)
}

func testAccVlanviewExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccVlanviewName(name string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "test_name" {
    name = %q
}
`, name)
}

func testAccVlanviewPreCreateVlan(preCreateVlan string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "test_pre_create_vlan" {
    pre_create_vlan = %q
}
`, preCreateVlan)
}

func testAccVlanviewStartVlanId(startVlanId string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "test_start_vlan_id" {
    start_vlan_id = %q
}
`, startVlanId)
}

func testAccVlanviewVlanNamePrefix(vlanNamePrefix string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "test_vlan_name_prefix" {
    vlan_name_prefix = %q
}
`, vlanNamePrefix)
}
