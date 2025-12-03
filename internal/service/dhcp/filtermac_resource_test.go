package dhcp_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForFiltermac = "comment,default_mac_address_expiration,disable,enforce_expiration_times,extattrs,lease_time,name,never_expires,options,reserved_for_infoblox"

func TestAccFiltermacResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_filtermac.test"
	var v dhcp.Filtermac
	name := acctest.RandomNameWithPrefix("mac_filter")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFiltermacBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "enforce_expiration_times", "true"),
					resource.TestCheckResourceAttr(resourceName, "never_expires", "true"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFiltermacResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_filtermac.test"
	name := acctest.RandomNameWithPrefix("mac_filter")
	var v dhcp.Filtermac
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFiltermacDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFiltermacBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					testAccCheckFiltermacDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccFiltermacResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_filtermac.test_comment"
	var v dhcp.Filtermac
	macFilterName := acctest.RandomNameWithPrefix("mac_filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFiltermacComment(macFilterName, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccFiltermacComment(macFilterName, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFiltermacResource_DefaultMacAddressExpiration(t *testing.T) {
	var resourceName = "nios_dhcp_filtermac.test_default_mac_address_expiration"
	var v dhcp.Filtermac
	macFilterName := acctest.RandomNameWithPrefix("mac_filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFiltermacDefaultMacAddressExpiration(macFilterName, "1200"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "default_mac_address_expiration", "1200"),
				),
			},
			// Update and Read
			{
				Config: testAccFiltermacDefaultMacAddressExpiration(macFilterName, "2400"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "default_mac_address_expiration", "2400"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFiltermacResource_Disable(t *testing.T) {
	var resourceName = "nios_dhcp_filtermac.test_disable"
	var v dhcp.Filtermac
	macFilterName := acctest.RandomNameWithPrefix("mac_filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFiltermacDisable(macFilterName, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFiltermacDisable(macFilterName, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFiltermacResource_EnforceExpirationTimes(t *testing.T) {
	var resourceName = "nios_dhcp_filtermac.test_enforce_expiration_times"
	var v dhcp.Filtermac
	macFilterName := acctest.RandomNameWithPrefix("mac_filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFiltermacEnforceExpirationTimes(macFilterName, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enforce_expiration_times", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFiltermacEnforceExpirationTimes(macFilterName, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enforce_expiration_times", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFiltermacResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dhcp_filtermac.test_extattrs"
	var v dhcp.Filtermac
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()
	macFilterName := acctest.RandomNameWithPrefix("mac_filter")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFiltermacExtAttrs(macFilterName, map[string]string{"Site": extAttrValue1}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccFiltermacExtAttrs(macFilterName, map[string]string{"Site": extAttrValue2}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFiltermacResource_LeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_filtermac.test_lease_time"
	var v dhcp.Filtermac
	lease_time := "7200"
	updated_lease_time := "3600"
	macFilterName := acctest.RandomNameWithPrefix("mac_filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFiltermacLeaseTime(macFilterName, lease_time),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_time", lease_time),
				),
			},
			// Update and Read
			{
				Config: testAccFiltermacLeaseTime(macFilterName, updated_lease_time),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_time", updated_lease_time),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFiltermacResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_filtermac.test_name"
	var v dhcp.Filtermac
	macFilterName := acctest.RandomNameWithPrefix("mac_filter")
	macFilterNameUpdate := acctest.RandomNameWithPrefix("mac_filter_upd")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFiltermacName(macFilterName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", macFilterName),
				),
			},
			// Update and Read
			{
				Config: testAccFiltermacName(macFilterNameUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", macFilterNameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFiltermacResource_NeverExpires(t *testing.T) {
	var resourceName = "nios_dhcp_filtermac.test_never_expires"
	var v dhcp.Filtermac
	macFilterName := acctest.RandomNameWithPrefix("mac_filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFiltermacNeverExpires(macFilterName, "true", ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "never_expires", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccFiltermacNeverExpires(macFilterName, "false", "3600"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "never_expires", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFiltermacResource_Options(t *testing.T) {
	var resourceName = "nios_dhcp_filtermac.test_options"
	var v dhcp.Filtermac
	macFilterName := acctest.RandomNameWithPrefix("mac_filter")
	options := []map[string]any{
		{
			"name":  "dhcp-lease-time",
			"num":   "51",
			"value": "1200",
		},
		{
			"name":  "time-offset",
			"num":   2,
			"value": "3600",
		},
	}
	updatedOptions := []map[string]any{
		{
			"name":  "dhcp-lease-time",
			"num":   "51",
			"value": "1800",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFiltermacOptions(macFilterName, options),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr(resourceName, "options.0.num", "51"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "1200"),
					resource.TestCheckResourceAttr(resourceName, "options.1.name", "time-offset"),
					resource.TestCheckResourceAttr(resourceName, "options.1.num", "2"),
					resource.TestCheckResourceAttr(resourceName, "options.1.value", "3600"),
				),
			},
			// Update and Read
			{
				Config: testAccFiltermacOptions(macFilterName, updatedOptions),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr(resourceName, "options.0.num", "51"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "1800"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFiltermacResource_ReservedForInfoblox(t *testing.T) {
	var resourceName = "nios_dhcp_filtermac.test_reserved_for_infoblox"
	var v dhcp.Filtermac
	reserved_for_infoblox := acctest.RandomName()
	reserved_for_infoblox_update := acctest.RandomName()
	macFilterName := acctest.RandomNameWithPrefix("mac_filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFiltermacReservedForInfoblox(macFilterName, reserved_for_infoblox),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "reserved_for_infoblox", reserved_for_infoblox),
				),
			},
			// Update and Read
			{
				Config: testAccFiltermacReservedForInfoblox(macFilterName, reserved_for_infoblox_update),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "reserved_for_infoblox", reserved_for_infoblox_update),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckFiltermacExists(ctx context.Context, resourceName string, v *dhcp.Filtermac) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			FiltermacAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForFiltermac).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetFiltermacResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetFiltermacResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckFiltermacDestroy(ctx context.Context, v *dhcp.Filtermac) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			FiltermacAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForFiltermac).
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

func testAccCheckFiltermacDisappears(ctx context.Context, v *dhcp.Filtermac) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			FiltermacAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccFiltermacBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filtermac" "test" {
    name = %q
}
`, name)
}

func testAccFiltermacComment(name string, comment string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filtermac" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccFiltermacDefaultMacAddressExpiration(name string, defaultMacAddressExpiration string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filtermac" "test_default_mac_address_expiration" {
    name = %q
    default_mac_address_expiration = %q
}
`, name, defaultMacAddressExpiration)
}

func testAccFiltermacDisable(name string, disable string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filtermac" "test_disable" {
    name = %q
    disable = %q
}
`, name, disable)
}

func testAccFiltermacEnforceExpirationTimes(name string, enforceExpirationTimes string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filtermac" "test_enforce_expiration_times" {
    name = %q
    enforce_expiration_times = %q
}
`, name, enforceExpirationTimes)
}
func testAccFiltermacExtAttrs(name string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf("  %s = %q\n", k, v)
	}
	extattrsStr += "}"
	return fmt.Sprintf(`
resource "nios_dhcp_filtermac" "test_extattrs" {
	name = %q
	extattrs = %s
}
`, name, extattrsStr)
}

func testAccFiltermacLeaseTime(name string, leaseTime string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filtermac" "test_lease_time" {
    name = %q
    lease_time = %q
}
`, name, leaseTime)
}

func testAccFiltermacName(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filtermac" "test_name" {
    name = %q
}
`, name)
}

func testAccFiltermacNeverExpires(name string, neverExpires string, macAddressExpiration string) string {
	expirationConfig := ""
	if macAddressExpiration != "" {
		expirationConfig = fmt.Sprintf("\n    default_mac_address_expiration = %s", macAddressExpiration)
	}
	return fmt.Sprintf(`
resource "nios_dhcp_filtermac" "test_never_expires" {
    name = %q
    never_expires = %q%s
}
`, name, neverExpires, expirationConfig)
}

func testAccFiltermacOptions(name string, options []map[string]any) string {
	optionsStr := utils.ConvertSliceOfMapsToHCL(options)
	return fmt.Sprintf(`
resource "nios_dhcp_filtermac" "test_options" {
    name = %q
    options = %s
}
`, name, optionsStr)
}
func testAccFiltermacReservedForInfoblox(name string, reservedForInfoblox string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filtermac" "test_reserved_for_infoblox" {
    name = %q
    reserved_for_infoblox = %q
}
`, name, reservedForInfoblox)
}
