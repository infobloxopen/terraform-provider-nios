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

var readableAttributesForVlan = "assigned_to,comment,contact,department,description,extattrs,id,name,parent,reserved,status"

func TestAccVlanResource_basic(t *testing.T) {
	var resourceName = "nios_ipam_vlan.test"
	var v ipam.Vlan
	name := acctest.RandomNameWithPrefix("vlan")
	view := acctest.RandomNameWithPrefix("example-vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanBasicConfig(51, name, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "id", "51"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrSet(resourceName, "parent"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "reserved", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanResource_disappears(t *testing.T) {
	resourceName := "nios_ipam_vlan.test"
	var v ipam.Vlan
	name := acctest.RandomNameWithPrefix("vlan")
	view := acctest.RandomNameWithPrefix("example-vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVlanDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccVlanBasicConfig(52, name, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					testAccCheckVlanDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccVlanResource_Import(t *testing.T) {
	var resourceName = "nios_ipam_vlan.test"
	var v ipam.Vlan
	name := acctest.RandomNameWithPrefix("vlan")
	view := acctest.RandomNameWithPrefix("example-vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanBasicConfig(53, name, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
				),
			},
			// Import with PlanOnly to detect differences
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccVlanImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "ref",
				PlanOnly:                             true,
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccVlanImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanResource_Comment(t *testing.T) {
	var resourceName = "nios_ipam_vlan.test_comment"
	var v ipam.Vlan
	name := acctest.RandomNameWithPrefix("vlan")
	view := acctest.RandomNameWithPrefix("example-vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanComment(54, name, view, "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanComment(54, name, view, "Updated comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the object"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanResource_Contact(t *testing.T) {
	var resourceName = "nios_ipam_vlan.test_contact"
	var v ipam.Vlan
	name := acctest.RandomNameWithPrefix("vlan")
	view := acctest.RandomNameWithPrefix("example-vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanContact(55, name, view, "contact_FIRST"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "contact", "contact_FIRST"),
				),
			},
			// Update and check for Empty String
			{
				Config: testAccVlanContact(55, name, view, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "contact", ""),
				),
			},
			// Update and Read
			{
				Config: testAccVlanContact(51, name, view, "CONTACT_2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "contact", "CONTACT_2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccVlanResource_Department(t *testing.T) {
	var resourceName = "nios_ipam_vlan.test_department"
	var v ipam.Vlan
	name := acctest.RandomNameWithPrefix("vlan")
	view := acctest.RandomNameWithPrefix("example-vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanDepartment(56, name, view, "DEPARTMENT"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "department", "DEPARTMENT"),
				),
			},
			// Update and Check for Empty String
			{
				Config: testAccVlanDepartment(56, name, view, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "department", ""),
				),
			},
			// Update and Read
			{
				Config: testAccVlanDepartment(56, name, view, "department_UPDATE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "department", "department_UPDATE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccVlanResource_Description(t *testing.T) {
	var resourceName = "nios_ipam_vlan.test_description"
	var v ipam.Vlan
	name := acctest.RandomNameWithPrefix("vlan")
	view := acctest.RandomNameWithPrefix("example-vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanDescription(57, name, view, "description_INITIAL"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "description_INITIAL"),
				),
			},
			// Update and Check for Empty String
			{
				Config: testAccVlanDescription(57, name, view, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
				),
			},
			// Update and Read
			{
				Config: testAccVlanDescription(51, name, view, "DESCRIPTION_UPDATE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "DESCRIPTION_UPDATE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccVlanResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_ipam_vlan.test_extattrs"
	var v ipam.Vlan
	name := acctest.RandomNameWithPrefix("vlan")
	view := acctest.RandomNameWithPrefix("example-vlan-view")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanExtAttrs(58, name, view, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccVlanExtAttrs(58, name, view, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVlanResource_Id(t *testing.T) {
	var resourceName = "nios_ipam_vlan.test_id"
	var v ipam.Vlan
	name := acctest.RandomNameWithPrefix("vlan")
	view := acctest.RandomNameWithPrefix("example-vlan-view")
	name2 := acctest.RandomNameWithPrefix("vlan")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanId(51, name, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "id", "51"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanId(59, name2, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "id", "59"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccVlanResource_Name(t *testing.T) {
	var resourceName = "nios_ipam_vlan.test_name"
	var v ipam.Vlan
	name := acctest.RandomNameWithPrefix("vlan")
	view := acctest.RandomNameWithPrefix("example-vlan-view")
	name2 := acctest.RandomNameWithPrefix("vlan")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanName(60, name, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccVlanName(60, name2, view),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccVlanResource_Parent(t *testing.T) {
	var resourceName = "nios_ipam_vlan.test_parent"
	var v ipam.Vlan
	name := acctest.RandomNameWithPrefix("vlan")
	view := acctest.RandomNameWithPrefix("example-vlan-view")
	view2 := acctest.RandomNameWithPrefix("example-vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanParent(61, name, view, view2, "one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrSet(resourceName, "parent"),
					resource.TestCheckResourceAttrPair(resourceName, "parent", "nios_ipam_vlanview.one", "ref"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanParent(61, name, view, view2, "two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "parent", "nios_ipam_vlanview.two", "ref"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccVlanResource_Reserved(t *testing.T) {
	var resourceName = "nios_ipam_vlan.test_reserved"
	var v ipam.Vlan
	name := acctest.RandomNameWithPrefix("vlan")
	view := acctest.RandomNameWithPrefix("example-vlan-view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVlanReserved(62, name, view, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "reserved", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccVlanReserved(62, name, view, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "reserved", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckVlanExists(ctx context.Context, resourceName string, v *ipam.Vlan) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.IPAMAPI.
			VlanAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForVlan).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetVlanResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetVlanResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckVlanDestroy(ctx context.Context, v *ipam.Vlan) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.IPAMAPI.
			VlanAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForVlan).
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

func testAccCheckVlanDisappears(ctx context.Context, v *ipam.Vlan) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.IPAMAPI.
			VlanAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccVlanImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccBaseWithVlanView(vlanViewName string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "%[1]s" {
  start_vlan_id = 50
  end_vlan_id   = 100
  name          = %[1]q
}
`, vlanViewName)
}

func testAccBaseWithTwoVlanViews(vlanViewName, vlanViewName2 string) string {
	return fmt.Sprintf(`
resource "nios_ipam_vlanview" "one" {
  start_vlan_id = 50
  end_vlan_id   = 100
  name          = %q
}

resource "nios_ipam_vlanview" "two" {
  start_vlan_id = 51
  end_vlan_id   = 101
  name          = %q
}
`, vlanViewName, vlanViewName2)
}

func testAccVlanBasicConfig(id int, name, parent string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_vlan" "test" {
    id = %d
    name = %q
    parent = nios_ipam_vlanview.%s.ref
}
`, id, name, parent)
	return strings.Join([]string{testAccBaseWithVlanView(parent), config}, "")
}

func testAccVlanComment(id int, name string, parent string, comment string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_vlan" "test_comment" {
    id = %d
    name = %q
    parent = nios_ipam_vlanview.%s.ref
    comment = %q
}
`, id, name, parent, comment)
	return strings.Join([]string{testAccBaseWithVlanView(parent), config}, "")
}

func testAccVlanContact(id int, name string, parent string, contact string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_vlan" "test_contact" {
    id = %d
    name = %q
    parent = nios_ipam_vlanview.%s.ref
    contact = %q
}
`, id, name, parent, contact)
	return strings.Join([]string{testAccBaseWithVlanView(parent), config}, "")
}

func testAccVlanDepartment(id int, name string, parent string, department string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_vlan" "test_department" {
    id = %d
    name = %q
    parent = nios_ipam_vlanview.%s.ref
    department = %q
}
`, id, name, parent, department)
	return strings.Join([]string{testAccBaseWithVlanView(parent), config}, "")
}

func testAccVlanDescription(id int, name string, parent string, description string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_vlan" "test_description" {
    id = %d
    name = %q
    parent = nios_ipam_vlanview.%s.ref
    description = %q
}
`, id, name, parent, description)
	return strings.Join([]string{testAccBaseWithVlanView(parent), config}, "")
}

func testAccVlanExtAttrs(id int, name string, parent string, extAttrs map[string]string) string {
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	config := fmt.Sprintf(`
resource "nios_ipam_vlan" "test_extattrs" {
    id = %d
    name = %q
    parent = nios_ipam_vlanview.%s.ref
    extattrs = %s
}
`, id, name, parent, extAttrsStr)
	return strings.Join([]string{testAccBaseWithVlanView(parent), config}, "")
}

func testAccVlanId(id int, name string, parent string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_vlan" "test_id" {
    id = %d
    name = %q
    parent = nios_ipam_vlanview.%s.ref
}
`, id, name, parent)
	return strings.Join([]string{testAccBaseWithVlanView(parent), config}, "")
}

func testAccVlanName(id int, name string, parent string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_vlan" "test_name" {
    id = %d
    name = %q
    parent = nios_ipam_vlanview.%s.ref
}
`, id, name, parent)
	return strings.Join([]string{testAccBaseWithVlanView(parent), config}, "")
}

func testAccVlanParent(id int, name string, vlanName1, vlanName2, parent string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_vlan" "test_parent" {
    id = %d
    name = %q
    parent = nios_ipam_vlanview.%s.ref
	depends_on = [nios_ipam_vlanview.one, nios_ipam_vlanview.two]
}
`, id, name, parent)
	return strings.Join([]string{testAccBaseWithTwoVlanViews(vlanName1, vlanName2), config}, "")
}

func testAccVlanReserved(id int, name string, parent string, reserved string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_vlan" "test_reserved" {
    id = %d
    name = %q
    parent = nios_ipam_vlanview.%s.ref
    reserved = %q
}
`, id, name, parent, reserved)
	return strings.Join([]string{testAccBaseWithVlanView(parent), config}, "")
}
