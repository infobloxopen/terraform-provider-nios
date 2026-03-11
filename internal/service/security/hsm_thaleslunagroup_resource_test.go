package security_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// NOTE: Complete the following steps before running the tests:
// 1. Generate client certificate on the NIOS Grid:
//    Grid -> Members -> Certificates -> Client Cert -> Generate Client Certificate (RSA SHA256)
// 2. Login as root to NIOS box and copy the client certificate to the Luna HSM box:
//    cd /storage/safenet-hsm/lunasa_7/cert/client/
//    scp <NIOS_IP>.pem admin@<Luna_HSM_IP>:
// 3. Login to Luna HSM box and register the NIOS client and assign partition:
//    ssh admin@<Luna_HSM_IP>
//    client register -c <username> -i <NIOS_IP>
//    client assignPartition -c <username> -P <partition_name>
// 4. Download server.pem from the Luna HSM box to your workstation and place it at:
//    internal/testdata/nios_security_hsm_thaleslunagroup/server.pem
// 5. Get the partition serial number using "partition list" on the Luna HSM box CLI

// Replace the placeholder values with actual values before running the tests

var readableAttributesForHsmThaleslunagroup = "comment,group_sn,hsm_version,name,status,thalesluna"

var (
	hsmThalesLunaVersion         = "Luna_7_CPL"
	hsmThalesLunaPassPhrase      = "< Enter Password for the Luna HSM >"
	hsmThalesLunaPartitionSerial = "< Enter Partition Serial Number >"
	hsmThalesLuna1               = "< Enter Name for Luna HSM Device 1>"
	hsmThalesLuna2               = "< Enter Name for Luna HSM Device 2>"

	hsmThalesLunaServerFilePath = filepath.Join(getServerCertPath(), "server.pem")

	hsmThalesLunaGroup = []map[string]any{
		{
			"name":                    hsmThalesLuna1,
			"partition_serial_number": hsmThalesLunaPartitionSerial,
			"server_cert_file_path":   hsmThalesLunaServerFilePath,
		},
	}

	hsmThalesLunaGroupUpdated = []map[string]any{
		{
			"name":                    hsmThalesLuna1,
			"partition_serial_number": hsmThalesLunaPartitionSerial,
			"server_cert_file_path":   hsmThalesLunaServerFilePath,
		},
		{
			"name":                    hsmThalesLuna2,
			"partition_serial_number": hsmThalesLunaPartitionSerial,
			"server_cert_file_path":   hsmThalesLunaServerFilePath,
			"disable":                 true,
		},
	}

	hsmThalesLunaGroup_HCL        = FormatThaleslunaToHCL(hsmThalesLunaGroup)
	hsmThalesLunaGroupUpdated_HCL = FormatThaleslunaToHCL(hsmThalesLunaGroupUpdated)
)

func TestAccHsmThaleslunagroupResource_basic(t *testing.T) {
	t.Skip("Skipping acceptance test as it requires a setup")
	var resourceName = "nios_security_hsm_thaleslunagroup.test"
	var v security.HsmThaleslunagroup

	name := acctest.RandomNameWithPrefix("thalesluna-hsm-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmThaleslunagroupBasicConfig(name, hsmThalesLunaVersion, hsmThalesLunaPassPhrase, hsmThalesLunaGroup_HCL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "hsm_version", hsmThalesLunaVersion),
					resource.TestCheckResourceAttr(resourceName, "pass_phrase", hsmThalesLunaPassPhrase),
					resource.TestCheckResourceAttr(resourceName, "thalesluna.0.name", hsmThalesLuna1),
					resource.TestCheckResourceAttr(resourceName, "thalesluna.0.partition_serial_number", hsmThalesLunaPartitionSerial),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "thalesluna.0.disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmThaleslunagroupResource_disappears(t *testing.T) {
	t.Skip("Skipping acceptance test as it requires a setup")
	resourceName := "nios_security_hsm_thaleslunagroup.test"
	var v security.HsmThaleslunagroup

	name := acctest.RandomNameWithPrefix("thalesluna-hsm-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHsmThaleslunagroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccHsmThaleslunagroupBasicConfig(name, hsmThalesLunaVersion, hsmThalesLunaPassPhrase, hsmThalesLunaGroup_HCL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					testAccCheckHsmThaleslunagroupDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccHsmThaleslunagroupResource_Comment(t *testing.T) {
	t.Skip("Skipping acceptance test as it requires a setup")
	var resourceName = "nios_security_hsm_thaleslunagroup.test_comment"
	var v security.HsmThaleslunagroup

	name := acctest.RandomNameWithPrefix("thalesluna-hsm-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmThaleslunagroupComment(name, hsmThalesLunaVersion, hsmThalesLunaPassPhrase, hsmThalesLunaGroup_HCL, "sample comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "sample comment"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmThaleslunagroupComment(name, hsmThalesLunaVersion, hsmThalesLunaPassPhrase, hsmThalesLunaGroup_HCL, "updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmThaleslunagroupResource_HsmVersion(t *testing.T) {
	t.Skip("Skipping acceptance test as it requires a setup")
	var resourceName = "nios_security_hsm_thaleslunagroup.test_hsm_version"
	var v security.HsmThaleslunagroup

	name := acctest.RandomNameWithPrefix("thalesluna-hsm-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmThaleslunagroupHsmVersion(name, hsmThalesLunaVersion, hsmThalesLunaPassPhrase, hsmThalesLunaGroup_HCL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hsm_version", hsmThalesLunaVersion),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmThaleslunagroupResource_Name(t *testing.T) {
	t.Skip("Skipping acceptance test as it requires a setup")
	var resourceName = "nios_security_hsm_thaleslunagroup.test_name"
	var v security.HsmThaleslunagroup

	name1 := acctest.RandomNameWithPrefix("thalesluna-hsm-")
	name2 := acctest.RandomNameWithPrefix("thalesluna-hsm-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmThaleslunagroupName(name1, hsmThalesLunaVersion, hsmThalesLunaPassPhrase, hsmThalesLunaGroup_HCL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccHsmThaleslunagroupName(name2, hsmThalesLunaVersion, hsmThalesLunaPassPhrase, hsmThalesLunaGroup_HCL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmThaleslunagroupResource_PassPhrase(t *testing.T) {
	t.Skip("Skipping acceptance test as it requires a setup")
	var resourceName = "nios_security_hsm_thaleslunagroup.test_pass_phrase"
	var v security.HsmThaleslunagroup

	name := acctest.RandomNameWithPrefix("thalesluna-hsm-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmThaleslunagroupPassPhrase(name, hsmThalesLunaVersion, hsmThalesLunaPassPhrase, hsmThalesLunaGroup_HCL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pass_phrase", hsmThalesLunaPassPhrase),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHsmThaleslunagroupResource_Thalesluna(t *testing.T) {
	t.Skip("Skipping acceptance test as it requires a setup")
	var resourceName = "nios_security_hsm_thaleslunagroup.test_thalesluna"
	var v security.HsmThaleslunagroup

	name := acctest.RandomNameWithPrefix("thalesluna-hsm-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHsmThaleslunagroupThalesluna(name, hsmThalesLunaVersion, hsmThalesLunaPassPhrase, hsmThalesLunaGroup_HCL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "thalesluna.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "thalesluna.0.name", hsmThalesLuna1),
					resource.TestCheckResourceAttr(resourceName, "thalesluna.0.disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccHsmThaleslunagroupThalesluna(name, hsmThalesLunaVersion, hsmThalesLunaPassPhrase, hsmThalesLunaGroupUpdated_HCL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHsmThaleslunagroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "thalesluna.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "thalesluna.0.name", hsmThalesLuna1),
					resource.TestCheckResourceAttr(resourceName, "thalesluna.0.disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "thalesluna.1.name", hsmThalesLuna2),
					resource.TestCheckResourceAttr(resourceName, "thalesluna.1.disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckHsmThaleslunagroupExists(ctx context.Context, resourceName string, v *security.HsmThaleslunagroup) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SecurityAPI.
			HsmThaleslunagroupAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForHsmThaleslunagroup).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetHsmThaleslunagroupResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetHsmThaleslunagroupResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckHsmThaleslunagroupDestroy(ctx context.Context, v *security.HsmThaleslunagroup) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SecurityAPI.
			HsmThaleslunagroupAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForHsmThaleslunagroup).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				return nil
			}
			return err
		}
		return errors.New("expected to be deleted")
	}
}

func testAccCheckHsmThaleslunagroupDisappears(ctx context.Context, v *security.HsmThaleslunagroup) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SecurityAPI.
			HsmThaleslunagroupAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		return err
	}
}

func testAccHsmThaleslunagroupBasicConfig(name, hsmVersion, passPhrase, thalesLuna string) string {
	return fmt.Sprintf(`
resource "nios_security_hsm_thaleslunagroup" "test" {
    name        = %q
    hsm_version = %q
    pass_phrase = %q
    thalesluna  = %s
}
`, name, hsmVersion, passPhrase, thalesLuna)
}

func testAccHsmThaleslunagroupComment(name, hsmVersion, passPhrase, thalesLuna, comment string) string {
	return fmt.Sprintf(`
resource "nios_security_hsm_thaleslunagroup" "test_comment" {
    name        = %q
    hsm_version = %q
    pass_phrase = %q
    thalesluna  = %s
    comment     = %q
}
`, name, hsmVersion, passPhrase, thalesLuna, comment)
}

func testAccHsmThaleslunagroupHsmVersion(name, hsmVersion, passPhrase, thalesLuna string) string {
	return fmt.Sprintf(`
resource "nios_security_hsm_thaleslunagroup" "test_hsm_version" {
    name        = %q
    hsm_version = %q
    pass_phrase = %q
    thalesluna  = %s
}
`, name, hsmVersion, passPhrase, thalesLuna)
}

func testAccHsmThaleslunagroupName(name, hsmVersion, passPhrase, thalesLuna string) string {
	return fmt.Sprintf(`
resource "nios_security_hsm_thaleslunagroup" "test_name" {
    name        = %q
    hsm_version = %q
    pass_phrase = %q
    thalesluna  = %s
}
`, name, hsmVersion, passPhrase, thalesLuna)
}

func testAccHsmThaleslunagroupPassPhrase(name, hsmVersion, passPhrase, thalesLuna string) string {
	return fmt.Sprintf(`
resource "nios_security_hsm_thaleslunagroup" "test_pass_phrase" {
	name        = %q
	hsm_version = %q
	pass_phrase = %q
	thalesluna  = %s
}
`, name, hsmVersion, passPhrase, thalesLuna)
}

func testAccHsmThaleslunagroupThalesluna(name, hsmVersion, passPhrase, thalesLuna string) string {
	return fmt.Sprintf(`
resource "nios_security_hsm_thaleslunagroup" "test_thalesluna" {
	name        = %q
	hsm_version = %q
	pass_phrase = %q
	thalesluna  = %s
}
`, name, hsmVersion, passPhrase, thalesLuna)
}

func FormatThaleslunaToHCL(thaleslunaList []map[string]any) string {
	var thaleslunaBlocks []string

	for _, thalesluna := range thaleslunaList {
		disable := false
		if val, ok := thalesluna["disable"]; ok {
			disable = val.(bool)
		}
		block := fmt.Sprintf(`    {
	  name                    = %q
      partition_serial_number = %q
      server_cert_file_path   = %q
      disable                 = %t
    }`,
			thalesluna["name"],
			thalesluna["partition_serial_number"],
			thalesluna["server_cert_file_path"],
			disable,
		)
		thaleslunaBlocks = append(thaleslunaBlocks, block)
	}

	return fmt.Sprintf(`[
%s
  ]`, strings.Join(thaleslunaBlocks, ",\n"))
}

func getServerCertPath() string {
	_, filename, _, _ := runtime.Caller(0)
	testDir := filepath.Dir(filename)
	return filepath.Join(testDir, "..", "..", "testdata", "nios_security_hsm_thaleslunagroup")
}
