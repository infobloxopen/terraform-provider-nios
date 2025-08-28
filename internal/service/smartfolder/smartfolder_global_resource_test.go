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

var readableAttributesForSmartfolderGlobal = "comment,group_bys,name,query_items"

func TestAccSmartfolderGlobalResource_basic(t *testing.T) {
	var resourceName = "nios_smartfolder_global.test"
	var v smartfolder.SmartfolderGlobal

	name := acctest.RandomNameWithPrefix("example-smartfolder-global-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSmartfolderGlobalBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderGlobalExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					//Test default values
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSmartfolderGlobalResource_disappears(t *testing.T) {
	resourceName := "nios_smartfolder_global.test"
	var v smartfolder.SmartfolderGlobal

	name := acctest.RandomNameWithPrefix("example-smartfolder-global-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSmartfolderGlobalDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSmartfolderGlobalBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderGlobalExists(context.Background(), resourceName, &v),
					testAccCheckSmartfolderGlobalDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSmartfolderGlobalResource_Comment(t *testing.T) {
	var resourceName = "nios_smartfolder_global.test_comment"
	var v smartfolder.SmartfolderGlobal

	name := acctest.RandomNameWithPrefix("example-smartfolder-global-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSmartfolderGlobalComment(name, "This is a comment."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderGlobalExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment."),
				),
			},
			// Update and Read
			{
				Config: testAccSmartfolderGlobalComment(name, "This is an updated comment."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderGlobalExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSmartfolderGlobalResource_GroupBys(t *testing.T) {
	var resourceName = "nios_smartfolder_global.test_group_bys"
	var v smartfolder.SmartfolderGlobal

	name := acctest.RandomNameWithPrefix("example-smartfolder-global-")

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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSmartfolderGlobalGroupBys(name, groupBys1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderGlobalExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "group_bys.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "group_bys.0.enable_grouping", "true"),
					resource.TestCheckResourceAttr(resourceName, "group_bys.0.value", "Availability zone"),
					resource.TestCheckResourceAttr(resourceName, "group_bys.0.value_type", "NORMAL"),
				),
			},
			{
				Config: testAccSmartfolderGlobalGroupBys(name, groupBys2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderGlobalExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "group_bys.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "group_bys.0.enable_grouping", "true"),
					resource.TestCheckResourceAttr(resourceName, "group_bys.0.value", "Site"),
					resource.TestCheckResourceAttr(resourceName, "group_bys.0.value_type", "EXTATTR"),
				),
			},
			{
				Config: testAccSmartfolderGlobalGroupBys(name, groupBys3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderGlobalExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "group_bys.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "group_bys.0.enable_grouping", "false"),
					resource.TestCheckResourceAttr(resourceName, "group_bys.0.value", "Site"),
					resource.TestCheckResourceAttr(resourceName, "group_bys.0.value_type", "EXTATTR"),
				),
			},
		},
	})
}

func TestAccSmartfolderGlobalResource_Name(t *testing.T) {
	var resourceName = "nios_smartfolder_global.test_name"
	var v smartfolder.SmartfolderGlobal

	name1 := acctest.RandomNameWithPrefix("example-smartfolder-global-")
	name2 := acctest.RandomNameWithPrefix("updated-example-smartfolder-global-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSmartfolderGlobalName(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderGlobalExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccSmartfolderGlobalName(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderGlobalExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSmartfolderGlobalResource_QueryItems(t *testing.T) {
	var resourceName = "nios_smartfolder_global.test_query_items"
	var v smartfolder.SmartfolderGlobal

	name := acctest.RandomNameWithPrefix("example-smartfolder-global-")

	queryItems1 := []map[string]any{
		{
			"field_type": "NORMAL",
			"name":       "type",
			"op_match":   true,
			"operator":   "EQ",
			"value": map[string]any{
				"value_string": "Host",
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSmartfolderGlobalQueryItems(name, queryItems1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderGlobalExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "query_items.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.field_type", "NORMAL"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.name", "type"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.op_match", "true"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.operator", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.value.value_string", "Host"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.value_type", "ENUM"),
				),
			},
			{
				Config: testAccSmartfolderGlobalQueryItems(name, queryItems2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmartfolderGlobalExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "query_items.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.field_type", "NORMAL"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.name", "type"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.op_match", "true"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.operator", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.value.value_string", "Zone"),
					resource.TestCheckResourceAttr(resourceName, "query_items.0.value_type", "ENUM"),
				),
			},
		},
	})
}

func testAccCheckSmartfolderGlobalExists(ctx context.Context, resourceName string, v *smartfolder.SmartfolderGlobal) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SmartFolderAPI.
			SmartfolderGlobalAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForSmartfolderGlobal).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetSmartfolderGlobalResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetSmartfolderGlobalResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckSmartfolderGlobalDestroy(ctx context.Context, v *smartfolder.SmartfolderGlobal) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SmartFolderAPI.
			SmartfolderGlobalAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForSmartfolderGlobal).
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

func testAccCheckSmartfolderGlobalDisappears(ctx context.Context, v *smartfolder.SmartfolderGlobal) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SmartFolderAPI.
			SmartfolderGlobalAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccSmartfolderGlobalBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_smartfolder_global" "test" {
  name = %q
}
`, name)
}

func testAccSmartfolderGlobalComment(name, comment string) string {
	return fmt.Sprintf(`
resource "nios_smartfolder_global" "test_comment" {
    name    = %q
    comment = %q
}
`, name, comment)
}

func testAccSmartfolderGlobalGroupBys(name string, groupBys []map[string]any) string {
	groupBysHCL := utils.ConvertSliceOfMapsToHCL(groupBys)
	return fmt.Sprintf(`
resource "nios_smartfolder_global" "test_group_bys" {
    name      = %q
    group_bys = %s
}
`, name, groupBysHCL)
}

func testAccSmartfolderGlobalName(name string) string {
	return fmt.Sprintf(`
resource "nios_smartfolder_global" "test_name" {
    name = %q
}
`, name)
}

func testAccSmartfolderGlobalQueryItems(name string, queryItems []map[string]any) string {
	queryItemsHCL := utils.ConvertSliceOfMapsToHCL(queryItems)
	return fmt.Sprintf(`
resource "nios_smartfolder_global" "test_query_items" {
    name        = %q
    query_items = %s
}
`, name, queryItemsHCL)
}
