package microsoft_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/microsoft"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForMssuperscope = "comment,dhcp_utilization,dhcp_utilization_status,disable,dynamic_hosts,extattrs,high_water_mark,high_water_mark_reset,low_water_mark,low_water_mark_reset,name,network_view,ranges,static_hosts,total_hosts"

func TestAccMssuperscopeResource_basic(t *testing.T) {
	var resourceName = "nios_microsoft_mssuperscope.test"
	var v microsoft.Mssuperscope
	name := acctest.RandomNameWithPrefix("mssuperscope")
	startAddrRange := "117.0.0.10"
	endAddrRange := "117.0.0.15"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeBasicConfig(name, startAddrRange, endAddrRange),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrPair(resourceName, "ranges.0", "nios_dhcp_range.test", "ref"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMssuperscopeResource_disappears(t *testing.T) {
	resourceName := "nios_microsoft_mssuperscope.test"
	var v microsoft.Mssuperscope
	name := acctest.RandomNameWithPrefix("mssuperscope")
	startAddrRange := "117.0.0.16"
	endAddrRange := "117.0.0.20"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMssuperscopeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMssuperscopeBasicConfig(name, startAddrRange, endAddrRange),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					testAccCheckMssuperscopeDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccMssuperscopeResource_Import(t *testing.T) {
	var resourceName = "nios_microsoft_mssuperscope.test"
	var v microsoft.Mssuperscope
	name := acctest.RandomNameWithPrefix("mssuperscope")
	startAddrRange := "117.0.0.21"
	endAddrRange := "117.0.0.25"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeBasicConfig(name, startAddrRange, endAddrRange),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
				),
			},
			// Import with PlanOnly to detect differences
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccMssuperscopeImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "uuid",
				PlanOnly:                             true,
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccMssuperscopeImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all"},
				ImportStateVerifyIdentifierAttribute: "uuid",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMssuperscopeResource_Comment(t *testing.T) {
	var resourceName = "nios_microsoft_mssuperscope.test_comment"
	var v microsoft.Mssuperscope
	name := acctest.RandomNameWithPrefix("mssuperscope")
	startAddrRange := "117.0.0.26"
	endAddrRange := "117.0.0.30"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeComment(name, startAddrRange, endAddrRange, "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccMssuperscopeComment(name, startAddrRange, endAddrRange, "Updated comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the object"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMssuperscopeResource_Disable(t *testing.T) {
	var resourceName = "nios_microsoft_mssuperscope.test_disable"
	var v microsoft.Mssuperscope
	name := acctest.RandomNameWithPrefix("mssuperscope")
	startAddrRange := "117.0.0.31"
	endAddrRange := "117.0.0.35"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeDisable(name, startAddrRange, endAddrRange, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccMssuperscopeDisable(name, startAddrRange, endAddrRange, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMssuperscopeResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_microsoft_mssuperscope.test_extattrs"
	var v microsoft.Mssuperscope
	name := acctest.RandomNameWithPrefix("mssuperscope")
	startAddrRange := "117.0.0.36"
	endAddrRange := "117.0.0.40"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeExtAttrs(name, startAddrRange, endAddrRange, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccMssuperscopeExtAttrs(name, startAddrRange, endAddrRange, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMssuperscopeResource_Name(t *testing.T) {
	var resourceName = "nios_microsoft_mssuperscope.test_name"
	var v microsoft.Mssuperscope
	name1 := acctest.RandomNameWithPrefix("mssuperscope")
	name2 := acctest.RandomNameWithPrefix("mssuperscope")
	startAddrRange := "117.0.0.41"
	endAddrRange := "117.0.0.45"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeName(name1, startAddrRange, endAddrRange),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccMssuperscopeName(name2, startAddrRange, endAddrRange),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMssuperscopeResource_NetworkView(t *testing.T) {
	var resourceName = "nios_microsoft_mssuperscope.test_network_view"
	var v microsoft.Mssuperscope
	name := acctest.RandomNameWithPrefix("mssuperscope")
	startAddrRange1 := "117.0.0.46"
	endAddrRange1 := "117.0.0.50"
	startAddrRange2 := "117.0.0.61"
	endAddrRange2 := "117.0.0.65"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeNetworkView(name, startAddrRange1, endAddrRange1, startAddrRange2, endAddrRange2, "test", "ms_server"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "ms_server"),
				),
			},
			// Update and Read
			{
				Config: testAccMssuperscopeNetworkView(name, startAddrRange1, endAddrRange1, startAddrRange2, endAddrRange2, "test2", "ms_server2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "ms_server2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccMssuperscopeResource_Ranges(t *testing.T) {
	var resourceName = "nios_microsoft_mssuperscope.test_ranges"
	var v microsoft.Mssuperscope
	name := acctest.RandomNameWithPrefix("mssuperscope")
	startAddrRange1 := "117.0.0.51"
	endAddrRange1 := "117.0.0.55"
	startAddrRange2 := "117.0.0.56"
	endAddrRange2 := "117.0.0.60"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccMssuperscopeRanges(name, startAddrRange1, endAddrRange1, startAddrRange2, endAddrRange2, "test", "ms_server"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "ranges.0", "nios_dhcp_range.test", "ref"),
				),
			},
			// Update and Read
			{
				Config: testAccMssuperscopeRanges(name, startAddrRange1, endAddrRange1, startAddrRange2, endAddrRange2, "test2", "ms_server2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "ranges.0", "nios_dhcp_range.test2", "ref"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckMssuperscopeExists(ctx context.Context, resourceName string, v *microsoft.Mssuperscope) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.MicrosoftAPI.
			MssuperscopeAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForMssuperscope).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetMssuperscopeResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetMssuperscopeResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckMssuperscopeDestroy(ctx context.Context, v *microsoft.Mssuperscope) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.MicrosoftAPI.
			MssuperscopeAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForMssuperscope).
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

func testAccCheckMssuperscopeDisappears(ctx context.Context, v *microsoft.Mssuperscope) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.MicrosoftAPI.
			MssuperscopeAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccMssuperscopeImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.Attributes["uuid"] == "" {
			return "", fmt.Errorf("uuid is not set")
		}
		return rs.Primary.Attributes["uuid"], nil
	}
}

func testAccBaseWithRanges(startAddr, endAddr string) string {
	return fmt.Sprintf(`
resource "nios_ipam_network" "example_network" {
	network      = "117.0.0.0/24"
	network_view = "ms_server"
	members = [
		{
			struct = "msdhcpserver"
			ipv4addr = "ms_example_server"
		}
	]
}

resource "nios_dhcp_range" "test" {
	start_addr = %q
	end_addr   = %q
	server_association_type = "MS_SERVER"
	ms_server = {ipv4addr="ms_example_server"}
	network_view = "ms_server"
	depends_on = [nios_ipam_network.example_network]
}
`, startAddr, endAddr)
}

func testAccMssuperscopeBasicConfig(name, startAddr string, endAddr string) string {
	config := fmt.Sprintf(`
resource "nios_microsoft_mssuperscope" "test" {
    name = %q
    ranges = [nios_dhcp_range.test.ref]
	network_view = "ms_server"
}
`, name)
	return strings.Join([]string{testAccBaseWithRanges(startAddr, endAddr), config}, "")
}

func testAccMssuperscopeComment(name string, startAddr string, endAddr string, comment string) string {
	config := fmt.Sprintf(`
resource "nios_microsoft_mssuperscope" "test_comment" {
    name = %q
    ranges = [nios_dhcp_range.test.ref]
    comment = %q
    network_view = "ms_server"
}
`, name, comment)
	return strings.Join([]string{testAccBaseWithRanges(startAddr, endAddr), config}, "")
}

func testAccMssuperscopeDisable(name string, startAddr string, endAddr string, disable string) string {
	config := fmt.Sprintf(`
resource "nios_microsoft_mssuperscope" "test_disable" {
    name = %q
    ranges = [nios_dhcp_range.test.ref]
    disable = %q
    network_view = "ms_server"
}
`, name, disable)
	return strings.Join([]string{testAccBaseWithRanges(startAddr, endAddr), config}, "")
}

func testAccMssuperscopeExtAttrs(name string, startAddr string, endAddr string, extAttrs map[string]string) string {
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	config := fmt.Sprintf(`
resource "nios_microsoft_mssuperscope" "test_extattrs" {
    name = %q
    ranges = [nios_dhcp_range.test.ref]
    extattrs = %s
    network_view = "ms_server"
}
`, name, extAttrsStr)
	return strings.Join([]string{testAccBaseWithRanges(startAddr, endAddr), config}, "")
}

func testAccMssuperscopeName(name string, startAddr string, endAddr string) string {
	config := fmt.Sprintf(`
resource "nios_microsoft_mssuperscope" "test_name" {
    name = %q
    ranges = [nios_dhcp_range.test.ref]
    network_view = "ms_server"
}
`, name)
	return strings.Join([]string{testAccBaseWithRanges(startAddr, endAddr), config}, "")
}

func testAccMssuperscopeNetworkView(name string, startAddrRange1, endAddrRange1, startAddrRange2, endAddrRange2, rangeResource, msServer string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_network" "example_network2" {
	network      = "117.0.0.0/24"
	network_view = "ms_server2"
	members = [
		{
			struct = "msdhcpserver"
			ipv4addr = "ms_example_server2"
		}
	]
}

resource "nios_dhcp_range" "test2" {
	start_addr = %[1]q
	end_addr   = %[2]q
	server_association_type = "MS_SERVER"
	ms_server = {ipv4addr="ms_example_server2"}
	network_view = "ms_server2"
	depends_on = [nios_ipam_network.example_network2]
}

resource "nios_microsoft_mssuperscope" "test_network_view" {
    name = %[3]q
    ranges = [nios_dhcp_range.%[4]s.ref]
    network_view = "%[5]s"
}
`, startAddrRange2, endAddrRange2, name, rangeResource, msServer)
	return strings.Join([]string{testAccBaseWithRanges(startAddrRange1, endAddrRange1), config}, "")
}

func testAccMssuperscopeRanges(name string, startAddrRange1, endAddrRange1, startAddrRange2, endAddrRange2, rangeResource, msServer string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_network" "example_network2" {
	network      = "117.0.0.0/24"
	network_view = "ms_server2"
	members = [
		{
			struct = "msdhcpserver"
			ipv4addr = "ms_example_server2"
		}
	]
}

resource "nios_dhcp_range" "test2" {
	start_addr = %[1]q
	end_addr   = %[2]q
	server_association_type = "MS_SERVER"
	ms_server = {ipv4addr="ms_example_server2"}
	network_view = "ms_server2"
	depends_on = [nios_ipam_network.example_network2]
}

resource "nios_microsoft_mssuperscope" "test_ranges" {
    name = %[3]q
    ranges = [nios_dhcp_range.%[4]s.ref]
    network_view = "%[5]s"
}
`, startAddrRange2, endAddrRange2, name, rangeResource, msServer)
	return strings.Join([]string{testAccBaseWithRanges(startAddrRange1, endAddrRange1), config}, "")
}
