package smartfolder_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/smartfolder"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForSmartfolderPersonal = "comment,group_bys,is_shortcut,name,query_items"

func TestAccSmartfolderPersonalResource_basic(t *testing.T) {
	var resourceName = "nios_smartfolder_personal.test"
	var v smartfolder.SmartfolderPersonal

	name := acctest.RandomNameWithPrefix("hi-example-smartfolder-personal-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSmartfolderPersonalBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderPersonalExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSmartfolderPersonalResource_disappears(t *testing.T) {
	resourceName := "nios_smartfolder_personal.test"
	var v smartfolder.SmartfolderPersonal

	name := acctest.RandomNameWithPrefix("example-smartfolder-personal-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSmartfolderPersonalDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSmartfolderPersonalBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderPersonalExists(context.Background(), resourceName, &v),
					testAccCheckSmartfolderPersonalDisappears(context.Background(), &v),
				),
				//ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSmartfolderPersonalResource_Comment(t *testing.T) {
	var resourceName = "nios_smartfolder_personal.test_comment"
	var v smartfolder.SmartfolderPersonal

	name := acctest.RandomNameWithPrefix("example-smartfolder-personal-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSmartfolderPersonalComment(name, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderPersonalExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccSmartfolderPersonalComment(name, "Updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderPersonalExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSmartfolderPersonalResource_GroupBys(t *testing.T) {
	var resourceName = "nios_smartfolder_personal.test_group_bys"
	var v smartfolder.SmartfolderPersonal

	name := acctest.RandomNameWithPrefix("example-smartfolder-personal-")
	groupBys1 := []map[string]any{
		{
			"enable_grouping": true,
			"value":           "Availability zone",
			"value_type":      "NORMAL",
		},
	}
	groupBys2 := []map[string]any{
		{
			"enable_grouping": true,
			"value":           "Site",
			"value_type":      "EXTATTR",
		},
	}
	groupBys3 := []map[string]any{
		{
			"enable_grouping": false,
			"value":           "Site",
			"value_type":      "EXTATTR",
		},
	}

	groupBysHCL1 := utils.ConvertSliceOfMapsToHCL(groupBys1)
	groupBysHCL2 := utils.ConvertSliceOfMapsToHCL(groupBys2)
	groupBysHCL3 := utils.ConvertSliceOfMapsToHCL(groupBys3)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSmartfolderPersonalGroupBys(name, groupBysHCL1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderPersonalExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "group_bys.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "group_bys.0.enable_grouping", "true"),
					resource.TestCheckResourceAttr(resourceName, "group_bys.0.value", "Availability zone"),
					resource.TestCheckResourceAttr(resourceName, "group_bys.0.value_type", "NORMAL"),
				),
			},
			// Update and Read
			{
				Config: testAccSmartfolderPersonalGroupBys(name, groupBysHCL2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderPersonalExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "group_bys.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "group_bys.0.enable_grouping", "true"),
					resource.TestCheckResourceAttr(resourceName, "group_bys.0.value", "Site"),
					resource.TestCheckResourceAttr(resourceName, "group_bys.0.value_type", "EXTATTR"),
				),
			},
			// Update and Read
			{
				Config: testAccSmartfolderPersonalGroupBys(name, groupBysHCL3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderPersonalExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "group_bys.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "group_bys.0.enable_grouping", "false"),
					resource.TestCheckResourceAttr(resourceName, "group_bys.0.value", "Site"),
					resource.TestCheckResourceAttr(resourceName, "group_bys.0.value_type", "EXTATTR"),
				),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSmartfolderPersonalResource_Name(t *testing.T) {
	var resourceName = "nios_smartfolder_personal.test_name"
	var v smartfolder.SmartfolderPersonal

	name1 := acctest.RandomNameWithPrefix("example-smartfolder-personal-")
	name2 := acctest.RandomNameWithPrefix("updated-example-smartfolder-personal-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSmartfolderPersonalName(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderPersonalExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccSmartfolderPersonalName(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderPersonalExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSmartfolderPersonalResource_QueryItems(t *testing.T) {
	var resourceName = "nios_smartfolder_personal.test_query_items"
	var v smartfolder.SmartfolderPersonal

	name := acctest.RandomNameWithPrefix("example-smartfolder-personal-")

	queryItems1 := []map[string]any{
		{
			"field_type": "NORMAL",
			"name":       "type",
			"op_match":   true,
			"operator":   "EQ",
			"value": map[string]any{
				"value_string": "Network",
			},
			"value_type": "ENUM",
		},
	}
	queryItems2 := []map[string]any{
		{
			"field_type": "NORMAL",
			"name":       "type",
			"op_match":   true,
			"operator":   "EQ",
			"value": map[string]any{
				"value_string": "Zone",
			},
			"value_type": "ENUM",
		},
	}

	queryItemsHCL1 := utils.ConvertSliceOfMapsToHCL(queryItems1)
	queryItemsHCL2 := utils.ConvertSliceOfMapsToHCL(queryItems2)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSmartfolderPersonalQueryItems(name, queryItemsHCL1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderPersonalExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "query_items.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.field_type", "NORMAL"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.name", "type"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.op_match", "true"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.operator", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.value.value_string", "Network"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.value_type", "ENUM"),
				),
			},
			// Update and Read
			{
				Config: testAccSmartfolderPersonalQueryItems(name, queryItemsHCL2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderPersonalExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "query_items.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.field_type", "NORMAL"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.name", "type"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.op_match", "true"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.operator", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.value.value_string", "Zone"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.value_type", "ENUM"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckSmartfolderPersonalExists(ctx context.Context, resourceName string, v *smartfolder.SmartfolderPersonal) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SmartFolderAPI.
			SmartfolderPersonalAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForSmartfolderPersonal).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetSmartfolderPersonalResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetSmartfolderPersonalResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckSmartfolderPersonalDestroy(ctx context.Context, v *smartfolder.SmartfolderPersonal) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SmartFolderAPI.
			SmartfolderPersonalAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForSmartfolderPersonal).
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

func testAccCheckSmartfolderPersonalDisappears(ctx context.Context, v *smartfolder.SmartfolderPersonal) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SmartFolderAPI.
			SmartfolderPersonalAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccSmartfolderPersonalBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_smartfolder_personal" "test" {
	name = %q
}
`, name)
}

func testAccSmartfolderPersonalComment(name, comment string) string {
	return fmt.Sprintf(`
resource "nios_smartfolder_personal" "test_comment" {
    name    = %q
    comment = %q
}
`, name, comment)
}

func testAccSmartfolderPersonalGroupBys(name, groupBys string) string {
	return fmt.Sprintf(`
resource "nios_smartfolder_personal" "test_group_bys" {
    name      = %q
    group_bys = %s
}
`, name, groupBys)
}

func testAccSmartfolderPersonalName(name string) string {
	return fmt.Sprintf(`
resource "nios_smartfolder_personal" "test_name" {
    name = %q
}
`, name)
}

func testAccSmartfolderPersonalQueryItems(name, queryItems string) string {
	return fmt.Sprintf(`
resource "nios_smartfolder_personal" "test_query_items" {
    name        = %q
    query_items = %s
}
`, name, queryItems)
}
