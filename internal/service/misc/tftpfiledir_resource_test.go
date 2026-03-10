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
	name := acctest.RandomNameWithPrefix("tftpfiledir")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTftpfiledirBasicConfig(name, "DIRECTORY"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "type", "DIRECTORY"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "directory", "/"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTftpfiledirResource_disappears(t *testing.T) {
	resourceName := "nios_misc_tftpfiledir.test"
	var v misc.Tftpfiledir
	name := acctest.RandomNameWithPrefix("tftpfiledir")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTftpfiledirDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccTftpfiledirBasicConfig(name, "DIRECTORY"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
					testAccCheckTftpfiledirDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccTftpfiledirResource_Import(t *testing.T) {
	var resourceName = "nios_misc_tftpfiledir.test"
	var v misc.Tftpfiledir
	name := acctest.RandomNameWithPrefix("tftpfiledir")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTftpfiledirBasicConfig(name, "DIRECTORY"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
				),
			},
			// Import with PlanOnly to detect differences
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccTftpfiledirImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "ref",
				PlanOnly:                             true,
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccTftpfiledirImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTftpfiledirResource_Directory(t *testing.T) {
	var resourceName = "nios_misc_tftpfiledir.test_directory"
	var v misc.Tftpfiledir
	name := acctest.RandomNameWithPrefix("tftpfiledir")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTftpfiledirDirectory(name, "DIRECTORY", "/filedir"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "directory", "/filedir"),
				),
			},
			// Skip Update testing as this field cannot be updated
		},
	})
}

func TestAccTftpfiledirResource_Name(t *testing.T) {
	var resourceName = "nios_misc_tftpfiledir.test_name"
	var v misc.Tftpfiledir
	name := acctest.RandomNameWithPrefix("tftpfiledir")
	name2 := acctest.RandomNameWithPrefix("tftpfiledir")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTftpfiledirName(name, "DIRECTORY"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccTftpfiledirName(name2, "DIRECTORY"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTftpfiledirResource_Type(t *testing.T) {
	var resourceName = "nios_misc_tftpfiledir.test_type"
	var v misc.Tftpfiledir
	name := acctest.RandomNameWithPrefix("tftpfiledir")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTftpfiledirType(name, "FILE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "type", "FILE"),
				),
			},
			// Type is immutable, hence Update testing is not applicable
		},
	})
}

func TestAccTftpfiledirResource_VtftpDirMembers(t *testing.T) {
	var resourceName = "nios_misc_tftpfiledir.test_vtftp_dir_members"
	var v misc.Tftpfiledir
	name := acctest.RandomNameWithPrefix("tftpfiledir")
	vtftpDirMembersVal := []map[string]any{
		{
			"member":  "infoblox.localdomain",
			"ip_type": "ADDRESS",
			"address": "192.168.1.1",
		},
		{
			"member":  "infoblox.localdomain",
			"ip_type": "NETWORK",
			"network": "192.168.1.1",
			"cidr":    24,
		},
		{
			"member":        "infoblox.localdomain",
			"ip_type":       "RANGE",
			"start_address": "192.168.1.1",
			"end_address":   "192.168.1.10",
		},
	}
	vtftpDirMembersValUpdated := []map[string]any{
		{
			"member":  "infoblox.localdomain",
			"ip_type": "ADDRESS",
			//"address": "192.168.1.1",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccTftpfiledirVtftpDirMembers(name, "DIRECTORY", vtftpDirMembersVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vtftp_dir_members.0.member", "infoblox.localdomain"),
					resource.TestCheckResourceAttr(resourceName, "vtftp_dir_members.0.ip_type", "ADDRESS"),
					resource.TestCheckResourceAttr(resourceName, "vtftp_dir_members.0.address", "192.168.1.1"),
					resource.TestCheckResourceAttr(resourceName, "vtftp_dir_members.1.member", "infoblox.localdomain"),
					resource.TestCheckResourceAttr(resourceName, "vtftp_dir_members.1.ip_type", "NETWORK"),
					resource.TestCheckResourceAttr(resourceName, "vtftp_dir_members.1.network", "192.168.1.1"),
					resource.TestCheckResourceAttr(resourceName, "vtftp_dir_members.1.cidr", "24"),
					resource.TestCheckResourceAttr(resourceName, "vtftp_dir_members.2.member", "infoblox.localdomain"),
					resource.TestCheckResourceAttr(resourceName, "vtftp_dir_members.2.ip_type", "RANGE"),
					resource.TestCheckResourceAttr(resourceName, "vtftp_dir_members.2.start_address", "192.168.1.1"),
					resource.TestCheckResourceAttr(resourceName, "vtftp_dir_members.2.end_address", "192.168.1.10"),
				),
			},
			// Update and Read
			{
				Config: testAccTftpfiledirVtftpDirMembers(name, "DIRECTORY", vtftpDirMembersValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTftpfiledirExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vtftp_dir_members.0.member", "infoblox.localdomain"),
					resource.TestCheckResourceAttr(resourceName, "vtftp_dir_members.0.ip_type", "ADDRESS")),
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

func testAccTftpfiledirImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccTftpfiledirBasicConfig(name, type_ string) string {
	return fmt.Sprintf(`
resource "nios_misc_tftpfiledir" "test" {
    name = %q
    type = %q
}
`, name, type_)
}

func testAccTftpfiledirDirectory(name string, type_ string, directory string) string {
	return fmt.Sprintf(`
resource "nios_misc_tftpfiledir" "test_directory" {
    name = %q
    type = %q
    directory = %q
}
`, name, type_, directory)
}

func testAccTftpfiledirName(name string, type_ string) string {
	return fmt.Sprintf(`
resource "nios_misc_tftpfiledir" "test_name" {
    name = %q
    type = %q
}
`, name, type_)
}

func testAccTftpfiledirType(name string, type_ string) string {
	return fmt.Sprintf(`
resource "nios_misc_tftpfiledir" "test_type" {
    name = %q
    type = %q
}
`, name, type_)
}

func testAccTftpfiledirVtftpDirMembers(name string, type_ string, vtftpDirMembers []map[string]any) string {
	vtftpDirMembersStr := utils.ConvertSliceOfMapsToHCL(vtftpDirMembers)
	return fmt.Sprintf(`
resource "nios_misc_tftpfiledir" "test_vtftp_dir_members" {
    name = %q
    type = %q
    vtftp_dir_members = %s
}
`, name, type_, vtftpDirMembersStr)
}
