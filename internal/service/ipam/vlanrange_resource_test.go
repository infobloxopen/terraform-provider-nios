package ipam_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
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
	vlanRange := acctest.RandomNameWithPrefix("vlan-range")
	vlanView := acctest.RandomNameWithPrefix("vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanrangeBasicConfig(71, vlanRange, 61, vlanView),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "end_vlan_id", "71"),
					resource.TestCheckResourceAttr(resourceName, "name", vlanRange),
					resource.TestCheckResourceAttr(resourceName, "start_vlan_id", "61"),
					resource.TestCheckResourceAttrPair(resourceName, "vlan_view", "nios_ipam_vlanview."+vlanView, "ref"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanrangeResource_disappears(t *testing.T) {
	resourceName := "nios_ipam_vlanrange.test"
	var v ipam.Vlanrange
	vlanRange := acctest.RandomNameWithPrefix("vlan-range")
	vlanView := acctest.RandomNameWithPrefix("vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVlanrangeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccVlanrangeBasicConfig(71, vlanRange, 61, vlanView),
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
	vlanRange := acctest.RandomNameWithPrefix("vlan-range")
	vlanView := acctest.RandomNameWithPrefix("vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanrangeBasicConfig(71, vlanRange, 61, vlanView),
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
	vlanRange := acctest.RandomNameWithPrefix("vlan-range")
	vlanView := acctest.RandomNameWithPrefix("vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanrangeComment(71, vlanRange, 61, vlanView, "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanrangeComment(71, vlanRange, 61, vlanView, "Updated comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the object"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanrangeResource_EndVlanId(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test_end_vlan_id"
	var v ipam.Vlanrange
	vlanRange := acctest.RandomNameWithPrefix("vlan-range")
	vlanView := acctest.RandomNameWithPrefix("vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanrangeEndVlanId(99, vlanRange, 51, vlanView),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "end_vlan_id", "99"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanrangeEndVlanId(51, vlanRange, 51, vlanView),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "end_vlan_id", "51"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccVlanrangeResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test_extattrs"
	var v ipam.Vlanrange
	vlanRange := acctest.RandomNameWithPrefix("vlan-range")
	vlanView := acctest.RandomNameWithPrefix("vlan-view")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanrangeExtAttrs(71, vlanRange, 61, vlanView, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccVlanrangeExtAttrs(71, vlanRange, 61, vlanView, map[string]string{
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
	vlanRange := acctest.RandomNameWithPrefix("vlan-range")
	vlanRangeUpdated := acctest.RandomNameWithPrefix("vlan-range")
	vlanView := acctest.RandomNameWithPrefix("vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanrangeName(71, vlanRange, 61, vlanView),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", vlanRange),
				),
			},
			// Update and Read
			{
				Config: testAccVlanrangeName(71, vlanRangeUpdated, 61, vlanView),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", vlanRangeUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanrangeResource_PreCreateVlan(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test_pre_create_vlan"
	var v ipam.Vlanrange
	vlanRange := acctest.RandomNameWithPrefix("vlan-range")
	vlanView := acctest.RandomNameWithPrefix("vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanrangePreCreateVlan(71, vlanRange, 61, vlanView, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pre_create_vlan", "true"),
				),
			},
			// Update Not Possible
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanrangeResource_StartVlanId(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test_start_vlan_id"
	var v ipam.Vlanrange
	vlanRange := acctest.RandomNameWithPrefix("vlan-range")
	vlanView := acctest.RandomNameWithPrefix("vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanrangeStartVlanId(71, vlanRange, 61, vlanView),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_vlan_id", "61"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanrangeStartVlanId(71, vlanRange, 51, vlanView),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "start_vlan_id", "51"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccVlanrangeResource_VlanNamePrefix(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test_vlan_name_prefix"
	var v ipam.Vlanrange
	vlanRange := acctest.RandomNameWithPrefix("vlan-range")
	vlanView := acctest.RandomNameWithPrefix("vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanrangeVlanNamePrefix(71, vlanRange, 61, vlanView, "prefixCaseInsensitive"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vlan_name_prefix", "prefixCaseInsensitive"),
				),
			},
			// Update Not Possible
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanrangeResource_VlanView(t *testing.T) {
	var resourceName = "nios_ipam_vlanrange.test_vlan_view"
	var v ipam.Vlanrange
	vlanRange := acctest.RandomNameWithPrefix("vlan-range")
	vlanView1 := acctest.RandomNameWithPrefix("vlan-view")
	vlanView2 := acctest.RandomNameWithPrefix("vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanrangeVlanView(71, vlanRange, 60, vlanView1, vlanView2, "one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "vlan_view", "nios_ipam_vlanview.one", "ref"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanrangeVlanView(71, vlanRange, 60, vlanView1, vlanView2, "two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanrangeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "vlan_view", "nios_ipam_vlanview.two", "ref"),
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

func testAccVlanrangeBasicConfig(endVlanId int, name string, startVlanId int, vlanView string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_vlanrange" "test" {
    end_vlan_id = %d
    name = %q
    start_vlan_id = %d
    vlan_view = nios_ipam_vlanview.%s.ref
}
`, endVlanId, name, startVlanId, vlanView)
	return strings.Join([]string{testAccBaseWithVlanView(vlanView), config}, "")
}

func testAccVlanrangeComment(endVlanId int, name string, startVlanId int, vlanView string, comment string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_vlanrange" "test_comment" {
    end_vlan_id = %d
    name = %q
    start_vlan_id = %d
    vlan_view = nios_ipam_vlanview.%s.ref
    comment = %q
}
`, endVlanId, name, startVlanId, vlanView, comment)
	return strings.Join([]string{testAccBaseWithVlanView(vlanView), config}, "")
}

func testAccVlanrangeEndVlanId(endVlanId int, name string, startVlanId int, vlanView string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_vlanrange" "test_end_vlan_id" {
    end_vlan_id = %d
    name = %q
    start_vlan_id = %d
    vlan_view = nios_ipam_vlanview.%s.ref
}
`, endVlanId, name, startVlanId, vlanView)
	return strings.Join([]string{testAccBaseWithVlanView(vlanView), config}, "")
}

func testAccVlanrangeExtAttrs(endVlanId int, name string, startVlanId int, vlanView string, extAttrs map[string]string) string {
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	config := fmt.Sprintf(`
resource "nios_ipam_vlanrange" "test_extattrs" {
    end_vlan_id = %d
    name = %q
    start_vlan_id = %d
    vlan_view = nios_ipam_vlanview.%s.ref
    extattrs = %s
}
`, endVlanId, name, startVlanId, vlanView, extAttrsStr)
	return strings.Join([]string{testAccBaseWithVlanView(vlanView), config}, "")
}

func testAccVlanrangeName(endVlanId int, name string, startVlanId int, vlanView string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_vlanrange" "test_name" {
    end_vlan_id = %d
    name = %q
    start_vlan_id = %d
    vlan_view = nios_ipam_vlanview.%s.ref
}
`, endVlanId, name, startVlanId, vlanView)
	return strings.Join([]string{testAccBaseWithVlanView(vlanView), config}, "")
}

func testAccVlanrangePreCreateVlan(endVlanId int, name string, startVlanId int, vlanView string, preCreateVlan string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_vlanrange" "test_pre_create_vlan" {
    end_vlan_id = %d
    name = %q
    start_vlan_id = %d
    vlan_view = nios_ipam_vlanview.%s.ref
    pre_create_vlan = %q
}
`, endVlanId, name, startVlanId, vlanView, preCreateVlan)
	return strings.Join([]string{testAccBaseWithVlanView(vlanView), config}, "")
}

func testAccVlanrangeStartVlanId(endVlanId int, name string, startVlanId int, vlanView string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_vlanrange" "test_start_vlan_id" {
    end_vlan_id = %d
    name = %q
    start_vlan_id = %d
    vlan_view = nios_ipam_vlanview.%s.ref
}
`, endVlanId, name, startVlanId, vlanView)
	return strings.Join([]string{testAccBaseWithVlanView(vlanView), config}, "")
}

func testAccVlanrangeVlanNamePrefix(endVlanId int, name string, startVlanId int, vlanView string, vlanNamePrefix string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_vlanrange" "test_vlan_name_prefix" {
    end_vlan_id = %d
    name = %q
    start_vlan_id = %d
    vlan_view = nios_ipam_vlanview.%s.ref
    vlan_name_prefix = %q
	pre_create_vlan = true
}
`, endVlanId, name, startVlanId, vlanView, vlanNamePrefix)
	return strings.Join([]string{testAccBaseWithVlanView(vlanView), config}, "")
}

func testAccVlanrangeVlanView(endVlanId int, name string, startVlanId int, vlanView1, vlanView2, parent string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_vlanrange" "test_vlan_view" {
    end_vlan_id = %d
    name = %q
    start_vlan_id = %d
    vlan_view = nios_ipam_vlanview.%s.ref
	depends_on = [nios_ipam_vlanview.one, nios_ipam_vlanview.two]
}
`, endVlanId, name, startVlanId, parent)
	return strings.Join([]string{testAccBaseWithTwoVlanViews(vlanView1, vlanView2), config}, "")
}
