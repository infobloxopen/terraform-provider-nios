package misc_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForTftpfiledir = "directory,is_synced_to_gm,last_modify,name,type,vtftp_dir_members"

func TestAccTftpfiledirResource_basic(t *testing.T) {
	var resourceName = "nios_misc_tftpfiledir.test"
	var v misc.Tftpfiledir

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTftpfiledirBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTftpfiledirResource_disappears(t *testing.T) {
	resourceName := "nios_misc_tftpfiledir.test"
	var v misc.Tftpfiledir

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTftpfiledirDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccTftpfiledirBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
					testAccCheckTftpfiledirDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccTftpfiledirResource_Ref(t *testing.T) {
	var resourceName = "nios_misc_tftpfiledir.test_ref"
	var v misc.Tftpfiledir

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTftpfiledirRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccTftpfiledirRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTftpfiledirResource_Directory(t *testing.T) {
	var resourceName = "nios_misc_tftpfiledir.test_directory"
	var v misc.Tftpfiledir

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTftpfiledirDirectory("DIRECTORY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "directory", "DIRECTORY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccTftpfiledirDirectory("DIRECTORY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "directory", "DIRECTORY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTftpfiledirResource_Name(t *testing.T) {
	var resourceName = "nios_misc_tftpfiledir.test_name"
	var v misc.Tftpfiledir

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTftpfiledirName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccTftpfiledirName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTftpfiledirResource_Type(t *testing.T) {
	var resourceName = "nios_misc_tftpfiledir.test_type"
	var v misc.Tftpfiledir

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTftpfiledirType("TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccTftpfiledirType("TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTftpfiledirResource_VtftpDirMembers(t *testing.T) {
	var resourceName = "nios_misc_tftpfiledir.test_vtftp_dir_members"
	var v misc.Tftpfiledir

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTftpfiledirVtftpDirMembers("VTFTP_DIR_MEMBERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vtftp_dir_members", "VTFTP_DIR_MEMBERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccTftpfiledirVtftpDirMembers("VTFTP_DIR_MEMBERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vtftp_dir_members", "VTFTP_DIR_MEMBERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckTftpfiledirExists(ctx context.Context, resourceName string, v *misc.Tftpfiledir) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.MiscAPI.
			TftpfiledirAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForTftpfiledir).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetTftpfiledirResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetTftpfiledirResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckTftpfiledirDestroy(ctx context.Context, v *misc.Tftpfiledir) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.MiscAPI.
			TftpfiledirAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForTftpfiledir).
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

func testAccCheckTftpfiledirDisappears(ctx context.Context, v *misc.Tftpfiledir) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.MiscAPI.
			TftpfiledirAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccTftpfiledirBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return `
resource "nios_misc_tftpfiledir" "test" {
}
`
}

func testAccTftpfiledirRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_misc_tftpfiledir" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccTftpfiledirDirectory(directory string) string {
	return fmt.Sprintf(`
resource "nios_misc_tftpfiledir" "test_directory" {
    directory = %q
}
`, directory)
}

func testAccTftpfiledirName(name string) string {
	return fmt.Sprintf(`
resource "nios_misc_tftpfiledir" "test_name" {
    name = %q
}
`, name)
}

func testAccTftpfiledirType(type_ string) string {
	return fmt.Sprintf(`
resource "nios_misc_tftpfiledir" "test_type" {
    type = %q
}
`, type_)
}

func testAccTftpfiledirVtftpDirMembers(vtftpDirMembers string) string {
	return fmt.Sprintf(`
resource "nios_misc_tftpfiledir" "test_vtftp_dir_members" {
    vtftp_dir_members = %q
}
`, vtftpDirMembers)
}
