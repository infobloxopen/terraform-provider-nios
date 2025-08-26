package grid_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForExtensibleattributedef = "allowed_object_types,comment,default_value,flags,list_values,max,min,name,namespace,type"

func TestAccExtensibleattributedefResource_basic(t *testing.T) {
	var resourceName = "nios_grid_extensibleattributedef.test"
	var v grid.Extensibleattributedef
	name := "tf_test_" + acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccExtensibleattributedefBasicConfig(name, "STRING"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExtensibleattributedefExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "type", "STRING"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccExtensibleattributedefResource_disappears(t *testing.T) {
	resourceName := "nios_grid_extensibleattributedef.test"
	var v grid.Extensibleattributedef
	name := "tf_test_" + acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckExtensibleattributedefDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccExtensibleattributedefBasicConfig(name, "STRING"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExtensibleattributedefExists(context.Background(), resourceName, &v),
					testAccCheckExtensibleattributedefDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccExtensibleattributedefResource_AllowedObjectTypes(t *testing.T) {
	var resourceName = "nios_grid_extensibleattributedef.test_allowed_object_types"
	var v grid.Extensibleattributedef
	name := "tf_test_" + acctest.RandomName()

	initialObjectTypes := []string{
		"NetworkContainer",
		"IPv6NetworkContainer",
		"Network",
	}

	updatedObjectTypes := []string{
		"NetworkContainer",
		"IPv6NetworkContainer",
		"Network",
		"IPv6Network",
		"FixedAddress",
		"IPv6FixedAddress",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccExtensibleattributedefAllowedObjectTypes(name, initialObjectTypes),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExtensibleattributedefExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allowed_object_types.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "allowed_object_types.0", "NetworkContainer"),
					resource.TestCheckResourceAttr(resourceName, "allowed_object_types.1", "IPv6NetworkContainer"),
					resource.TestCheckResourceAttr(resourceName, "allowed_object_types.2", "Network"),
				),
			},
			// Update and Read
			{
				Config: testAccExtensibleattributedefAllowedObjectTypes(name, updatedObjectTypes),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExtensibleattributedefExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allowed_object_types.#", "6"),
					resource.TestCheckResourceAttr(resourceName, "allowed_object_types.0", "NetworkContainer"),
					resource.TestCheckResourceAttr(resourceName, "allowed_object_types.1", "IPv6NetworkContainer"),
					resource.TestCheckResourceAttr(resourceName, "allowed_object_types.2", "Network"),
					resource.TestCheckResourceAttr(resourceName, "allowed_object_types.3", "IPv6Network"),
					resource.TestCheckResourceAttr(resourceName, "allowed_object_types.4", "FixedAddress"),
					resource.TestCheckResourceAttr(resourceName, "allowed_object_types.5", "IPv6FixedAddress"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccExtensibleattributedefResource_Comment(t *testing.T) {
	var resourceName = "nios_grid_extensibleattributedef.test_comment"
	var v grid.Extensibleattributedef

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccExtensibleattributedefComment("EXTDEF COMMENT"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExtensibleattributedefExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "EXTDEF COMMENT"),
				),
			},
			// Update and Read
			{
				Config: testAccExtensibleattributedefComment("EXTDEF COMMENT UPDATE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExtensibleattributedefExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "EXTDEF COMMENT UPDATE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccExtensibleattributedefResource_DefaultValue(t *testing.T) {
	var resourceName = "nios_grid_extensibleattributedef.test_default_value"
	var v grid.Extensibleattributedef

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccExtensibleattributedefDefaultValue("STRING"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExtensibleattributedefExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "default_value", "STRING"),
				),
			},
			// Update and Read
			{
				Config: testAccExtensibleattributedefDefaultValue("9945"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExtensibleattributedefExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "default_value", "9945"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccExtensibleattributedefResource_Flags(t *testing.T) {
	var resourceName = "nios_grid_extensibleattributedef.test_flags"
	var v grid.Extensibleattributedef

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccExtensibleattributedefFlags("C"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExtensibleattributedefExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "flags", "C"),
				),
			},
			// Update and Read
			{
				Config: testAccExtensibleattributedefFlags("CL"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExtensibleattributedefExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "flags", "CL"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccExtensibleattributedefResource_ListValues(t *testing.T) {
	var resourceName = "nios_grid_extensibleattributedef.test_list_values"
	var v grid.Extensibleattributedef
	name := "tf_test_list_" + acctest.RandomName()

	initialListValues := []string{
		"value1",
		"value2",
		"value3",
	}

	updatedListValues := []string{
		"value1",
		"value2",
		"value3",
		"value4",
		"value5",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccExtensibleattributedefListValues(name, initialListValues),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExtensibleattributedefExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "list_values.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "list_values.0.value", "value1"),
					resource.TestCheckResourceAttr(resourceName, "list_values.1.value", "value2"),
					resource.TestCheckResourceAttr(resourceName, "list_values.2.value", "value3"),
				),
			},
			// Update and Read
			{
				Config: testAccExtensibleattributedefListValues(name, updatedListValues),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExtensibleattributedefExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "list_values.#", "5"),
					resource.TestCheckResourceAttr(resourceName, "list_values.0.value", "value1"),
					resource.TestCheckResourceAttr(resourceName, "list_values.1.value", "value2"),
					resource.TestCheckResourceAttr(resourceName, "list_values.2.value", "value3"),
					resource.TestCheckResourceAttr(resourceName, "list_values.3.value", "value4"),
					resource.TestCheckResourceAttr(resourceName, "list_values.4.value", "value5"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccExtensibleattributedefResource_Max(t *testing.T) {
	var resourceName = "nios_grid_extensibleattributedef.test_max"
	var v grid.Extensibleattributedef
	name := "tf_test_max_" + acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccExtensibleattributedefMax(name, 100),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExtensibleattributedefExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max", "100"),
				),
			},
			// Update and Read
			{
				Config: testAccExtensibleattributedefMax(name, 200),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExtensibleattributedefExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max", "200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccExtensibleattributedefResource_Min(t *testing.T) {
	var resourceName = "nios_grid_extensibleattributedef.test_min"
	var v grid.Extensibleattributedef
	name := "tf_test_min_" + acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccExtensibleattributedefMin(name, 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExtensibleattributedefExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "min", "10"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccExtensibleattributedefResource_Name(t *testing.T) {
	var resourceName = "nios_grid_extensibleattributedef.test_name"
	var v grid.Extensibleattributedef

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccExtensibleattributedefName("tf_test_name_initial"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExtensibleattributedefExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "tf_test_name_initial"),
				),
			},
			// Update and Read
			{
				Config: testAccExtensibleattributedefName("tf_test_name_updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExtensibleattributedefExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "tf_test_name_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccExtensibleattributedefResource_Type(t *testing.T) {
	var resourceName = "nios_grid_extensibleattributedef.test_type"
	var v grid.Extensibleattributedef
	name := "tf_test_type_" + acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccExtensibleattributedefType(name, "STRING"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExtensibleattributedefExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "STRING"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckExtensibleattributedefExists(ctx context.Context, resourceName string, v *grid.Extensibleattributedef) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.GridAPI.
			ExtensibleattributedefAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForExtensibleattributedef).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetExtensibleattributedefResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetExtensibleattributedefResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckExtensibleattributedefDestroy(ctx context.Context, v *grid.Extensibleattributedef) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.GridAPI.
			ExtensibleattributedefAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForExtensibleattributedef).
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

func testAccCheckExtensibleattributedefDisappears(ctx context.Context, v *grid.Extensibleattributedef) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.GridAPI.
			ExtensibleattributedefAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccExtensibleattributedefBasicConfig(name, type_ string) string {
	return fmt.Sprintf(`
resource "nios_grid_extensibleattributedef" "test" {
  name = %q
  type = %q

}
`, name, type_)
}

func testAccExtensibleattributedefAllowedObjectTypes(name string, allowedObjectTypes []string) string {
	objectTypesStr := "[\n"
	for _, objType := range allowedObjectTypes {
		objectTypesStr += fmt.Sprintf("\t\t%q,\n", objType)
	}
	objectTypesStr += "\t]"

	return fmt.Sprintf(`
resource "nios_grid_extensibleattributedef" "test_allowed_object_types" {
    name = %q
    type = "STRING"
    allowed_object_types = %s
}
`, name, objectTypesStr)
}

func testAccExtensibleattributedefComment(comment string) string {
	name := "tf_test_comment_" + acctest.RandomName()
	return fmt.Sprintf(`
resource "nios_grid_extensibleattributedef" "test_comment" {
    name = %q
    type = "STRING"
    comment = %q
}
`, name, comment)
}

func testAccExtensibleattributedefDefaultValue(defaultValue string) string {
	name := "tf_test_defval_" + acctest.RandomName()
	return fmt.Sprintf(`
resource "nios_grid_extensibleattributedef" "test_default_value" {
    name = %q
    type = "STRING"
    default_value = %q
}
`, name, defaultValue)
}

func testAccExtensibleattributedefFlags(flags string) string {
	name := "tf_test_flags_" + acctest.RandomName()
	return fmt.Sprintf(`
resource "nios_grid_extensibleattributedef" "test_flags" {
    name = %q
    type = "STRING"
    flags = %q
}
`, name, flags)
}

func testAccExtensibleattributedefListValues(name string, listValues []string) string {
	valuesStr := "[\n"
	for _, val := range listValues {
		valuesStr += fmt.Sprintf("\t{\n\t\tvalue = %q\n\t},\n", val)
	}
	valuesStr += "]"

	return fmt.Sprintf(`
resource "nios_grid_extensibleattributedef" "test_list_values" {
    name = %q
    type = "ENUM"
    flags = "L"
    list_values = %s
}
`, name, valuesStr)
}

func testAccExtensibleattributedefMax(name string, maxValue int) string {
	return fmt.Sprintf(`
resource "nios_grid_extensibleattributedef" "test_max" {
    name = %q
    type = "INTEGER"
    max = %d
}
`, name, maxValue)
}

func testAccExtensibleattributedefMin(name string, minValue int) string {
	return fmt.Sprintf(`
resource "nios_grid_extensibleattributedef" "test_min" {
    name = %q
    type = "INTEGER"
    min = %d
}
`, name, minValue)
}

func testAccExtensibleattributedefName(name string) string {
	return fmt.Sprintf(`
resource "nios_grid_extensibleattributedef" "test_name" {
    name = %q
	type = "STRING"
}
`, name)
}

func testAccExtensibleattributedefType(name string, type_ string) string {
	return fmt.Sprintf(`
resource "nios_grid_extensibleattributedef" "test_type" {
    name = %q
    type = %q
}
`, name, type_)
}
