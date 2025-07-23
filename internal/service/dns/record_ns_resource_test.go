package dns_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// TODO: OBJECTS TO BE PRESENT IN GRID FOR TESTS
// -> Parent Zone: example.com (in default view)

var readableAttributesForRecordNs = "addresses,cloud_info,creator,dns_name,last_queried,ms_delegation_name,name,nameserver,policy,view,zone"

func TestAccRecordNsResource_basic(t *testing.T) {
	var resourceName = "nios_dns_record_ns.test"
	var v dns.RecordNs
	name := "example.com"
	nameserver := acctest.RandomNameWithPrefix("nameserver") + ".example.com"
	addresses := []map[string]any{
		{
			"address":         "20.0.0.0",
			"auto_create_ptr": false,
		},
	}
	addressesHCL := FormatZoneNameServersToHCL(addresses)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNsBasicConfig(name, nameserver, addressesHCL, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "nameserver", nameserver),
					resource.TestCheckResourceAttr(resourceName, "addresses.0.address", "20.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "addresses.0.auto_create_ptr", "false"),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
					resource.TestCheckResourceAttr(resourceName, "ms_delegation_name", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordNsResource_disappears(t *testing.T) {
	resourceName := "nios_dns_record_ns.test"
	var v dns.RecordNs
	name := "example.com"
	nameserver := acctest.RandomNameWithPrefix("nameserver") + ".example.com"
	addresses := []map[string]any{
		{
			"address":         "20.0.0.0",
			"auto_create_ptr": false,
		},
	}
	addressesHCL := FormatZoneNameServersToHCL(addresses)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordNsDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordNsBasicConfig(name, nameserver, addressesHCL, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNsExists(context.Background(), resourceName, &v),
					testAccCheckRecordNsDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordNsResource_Addresses(t *testing.T) {
	var resourceName = "nios_dns_record_ns.test_addresses"
	var v dns.RecordNs
	name := "example.com"
	nameserver := acctest.RandomNameWithPrefix("nameserver") + ".example.com"
	addresses1 := []map[string]any{
		{
			"address":         "20.0.0.0",
			"auto_create_ptr": false,
		},
	}
	addresses2 := []map[string]any{
		{
			"address":         "40.0.0.0",
			"auto_create_ptr": false,
		},
	}
	addresses3 := []map[string]any{
		{
			"address":         "40.0.0.0",
			"auto_create_ptr": true,
		},
	}

	addressesHCL1 := FormatZoneNameServersToHCL(addresses1)
	addressesHCL2 := FormatZoneNameServersToHCL(addresses2)
	addressesHCL3 := FormatZoneNameServersToHCL(addresses3)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNsAddresses(name, nameserver, addressesHCL1, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "addresses.0.address", "20.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "addresses.0.auto_create_ptr", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNsAddresses(name, nameserver, addressesHCL2, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "addresses.0.address", "40.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "addresses.0.auto_create_ptr", "false"),
				),
			},
			// Update field auto_create_ptr and Read
			{
				Config: testAccRecordNsAddresses(name, "ns1.example.com", addressesHCL3, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "addresses.0.address", "40.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "addresses.0.auto_create_ptr", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordNsResource_Nameserver(t *testing.T) {
	var resourceName = "nios_dns_record_ns.test_nameserver"
	var v dns.RecordNs
	name := "example.com"
	nameserver1 := acctest.RandomNameWithPrefix("nameserver") + ".example.com"
	nameserver2 := acctest.RandomNameWithPrefix("nameserver_updated") + ".example.com"
	addresses := []map[string]any{
		{
			"address":         "20.0.0.0",
			"auto_create_ptr": false,
		},
	}
	addressesHCL := FormatZoneNameServersToHCL(addresses)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNsNameserver(name, nameserver1, addressesHCL, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nameserver", nameserver1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNsNameserver(name, nameserver2, addressesHCL, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordNsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nameserver", nameserver2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordNsExists(ctx context.Context, resourceName string, v *dns.RecordNs) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			RecordNsAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordNs).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordNsResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordNsResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordNsDestroy(ctx context.Context, v *dns.RecordNs) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			RecordNsAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordNs).
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

func testAccCheckRecordNsDisappears(ctx context.Context, v *dns.RecordNs) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			RecordNsAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordNsBasicConfig(name, nameserver, addresses, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ns" "test" {
    name        = %q
    nameserver  = %q
    addresses   = %s
    view        = %q
}

`, name, nameserver, addresses, view)
}

func testAccRecordNsAddresses(name, nameserver, addresses, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ns" "test_addresses" {
    name        = %q
    nameserver  = %q
    addresses   = %s
    view        = %q
}
`, name, nameserver, addresses, view)
}

func testAccRecordNsNameserver(name, nameserver, addresses, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_ns" "test_nameserver" {
    name        = %q
    nameserver  = %q
    addresses   = %s
    view        = %q
}
`, name, nameserver, addresses, view)
}

func FormatZoneNameServersToHCL(servers []map[string]any) string {
	var serverBlocks []string

	for _, server := range servers {
		block := fmt.Sprintf(`    {
      address = %q
      auto_create_ptr = %t
    }`, server["address"], server["auto_create_ptr"])
		serverBlocks = append(serverBlocks, block)
	}

	return fmt.Sprintf(`[
%s
  ]`, strings.Join(serverBlocks, ",\n"))
}
