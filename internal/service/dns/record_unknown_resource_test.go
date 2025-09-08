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

var readableAttributesForRecordUnknown = "cloud_info,comment,creator,disable,display_rdata,dns_name,enable_host_name_policy,extattrs,last_queried,name,policy,record_type,subfield_values,ttl,use_ttl,view,zone"

func TestAccRecordUnknownResource_basic(t *testing.T) {
	var resourceName = "nios_dns_record_unknown.test"
	var v dns.RecordUnknown
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-unknown") + ".com"
	subfieldValues := []map[string]any{
		{
			"field_type":     "T",
			"field_value":    "example-text",
			"include_length": "8_BIT",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordUnknownBasicConfig(zoneFqdn, name, "SPF", subfieldValues),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "record_type", "SPF"),
					resource.TestCheckResourceAttr(resourceName, "subfield_values.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subfield_values.0.field_type", "T"),
					resource.TestCheckResourceAttr(resourceName, "subfield_values.0.field_value", "example-text"),
					resource.TestCheckResourceAttr(resourceName, "subfield_values.0.include_length", "8_BIT"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
					resource.TestCheckResourceAttr(resourceName, "enable_host_name_policy", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordUnknownResource_disappears(t *testing.T) {
	resourceName := "nios_dns_record_unknown.test"
	var v dns.RecordUnknown
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-unknown") + ".com"
	subfieldValues := []map[string]any{
		{
			"field_type":     "T",
			"field_value":    "example-text",
			"include_length": "8_BIT",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordUnknownDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordUnknownBasicConfig(zoneFqdn, name, "SPF", subfieldValues),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					testAccCheckRecordUnknownDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordUnknownResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_record_unknown.test_comment"
	var v dns.RecordUnknown
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-unknown") + ".com"
	subfieldValues := []map[string]any{
		{
			"field_type":     "T",
			"field_value":    "example-text",
			"include_length": "8_BIT",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordUnknownComment(zoneFqdn, name, "SPF", subfieldValues, "This is a comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a comment"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordUnknownComment(zoneFqdn, name, "SPF", subfieldValues, "This is an updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordUnknownResource_Creator(t *testing.T) {
	var resourceName = "nios_dns_record_unknown.test_creator"
	var v dns.RecordUnknown
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-unknown") + ".com"
	subfieldValues := []map[string]any{
		{
			"field_type":     "T",
			"field_value":    "example-text",
			"include_length": "8_BIT",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordUnknownCreator(zoneFqdn, name, "SPF", subfieldValues, "STATIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordUnknownCreator(zoneFqdn, name, "SPF", subfieldValues, "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "DYNAMIC"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordUnknownResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_record_unknown.test_disable"
	var v dns.RecordUnknown
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-unknown") + ".com"
	subfieldValues := []map[string]any{
		{
			"field_type":     "T",
			"field_value":    "example-text",
			"include_length": "8_BIT",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordUnknownDisable(zoneFqdn, name, "SPF", subfieldValues, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordUnknownDisable(zoneFqdn, name, "SPF", subfieldValues, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordUnknownResource_EnableHostNamePolicy(t *testing.T) {
	var resourceName = "nios_dns_record_unknown.test_enable_host_name_policy"
	var v dns.RecordUnknown
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-unknown") + ".com"
	subfieldValues := []map[string]any{
		{
			"field_type":     "T",
			"field_value":    "example-text",
			"include_length": "8_BIT",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordUnknownEnableHostNamePolicy(zoneFqdn, name, "SPF", subfieldValues, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_host_name_policy", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordUnknownEnableHostNamePolicy(zoneFqdn, name, "SPF", subfieldValues, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_host_name_policy", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordUnknownResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_record_unknown.test_extattrs"
	var v dns.RecordUnknown
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-unknown") + ".com"
	subfieldValues := []map[string]any{
		{
			"field_type":     "T",
			"field_value":    "example-text",
			"include_length": "8_BIT",
		},
	}
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordUnknownExtAttrs(zoneFqdn, name, "SPF", subfieldValues, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordUnknownExtAttrs(zoneFqdn, name, "SPF", subfieldValues, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordUnknownResource_Name(t *testing.T) {
	var resourceName = "nios_dns_record_unknown.test_name"
	var v dns.RecordUnknown
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-unknown") + ".com"
	updatedName := acctest.RandomNameWithPrefix("record-unknown") + ".com"
	subfieldValues := []map[string]any{
		{
			"field_type":     "T",
			"field_value":    "example-text",
			"include_length": "8_BIT",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordUnknownName(zoneFqdn, name, "SPF", subfieldValues),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("%s.%s", name, zoneFqdn)),
				),
			},
			// Update and Read
			{
				Config: testAccRecordUnknownName(zoneFqdn, updatedName, "SPF", subfieldValues),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("%s.%s", updatedName, zoneFqdn)),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordUnknownResource_RecordType(t *testing.T) {
	var resourceName = "nios_dns_record_unknown.test_record_type"
	var v dns.RecordUnknown
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-unknown") + ".com"
	subfieldValues1 := []map[string]any{
		{
			"field_type":     "T",
			"field_value":    "example-text",
			"include_length": "8_BIT",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordUnknownRecordType(zoneFqdn, name, "SPF", subfieldValues1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "record_type", "SPF"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordUnknownResource_SubfieldValues(t *testing.T) {
	var resourceName = "nios_dns_record_unknown.test_subfield_values"
	var v dns.RecordUnknown
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-unknown") + ".com"
	subfieldValues := []map[string]any{
		{
			"field_type":     "T",
			"field_value":    "example-text",
			"include_length": "8_BIT",
		},
	}
	updatedSubfieldValues := []map[string]any{
		{
			"field_type":     "T",
			"field_value":    "updated-text",
			"include_length": "16_BIT",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordUnknownSubfieldValues(zoneFqdn, name, "SPF", subfieldValues),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subfield_values.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subfield_values.0.field_type", "T"),
					resource.TestCheckResourceAttr(resourceName, "subfield_values.0.field_value", "example-text"),
					resource.TestCheckResourceAttr(resourceName, "subfield_values.0.include_length", "8_BIT"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordUnknownSubfieldValues(zoneFqdn, name, "SPF", updatedSubfieldValues),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subfield_values.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subfield_values.0.field_type", "T"),
					resource.TestCheckResourceAttr(resourceName, "subfield_values.0.field_value", "updated-text"),
					resource.TestCheckResourceAttr(resourceName, "subfield_values.0.include_length", "16_BIT"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordUnknownResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_record_unknown.test_ttl"
	var v dns.RecordUnknown
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-unknown") + ".com"
	subfieldValues := []map[string]any{
		{
			"field_type":     "T",
			"field_value":    "example-text",
			"include_length": "8_BIT",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordUnknownTtl(zoneFqdn, name, "SPF", subfieldValues, 10, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordUnknownTtl(zoneFqdn, name, "SPF", subfieldValues, 20, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordUnknownResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_record_unknown.test_use_ttl"
	var v dns.RecordUnknown
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-unknown") + ".com"
	subfieldValues := []map[string]any{
		{
			"field_type":     "T",
			"field_value":    "example-text",
			"include_length": "8_BIT",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordUnknownUseTtl(zoneFqdn, name, "SPF", subfieldValues, 10, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordUnknownUseTtl(zoneFqdn, name, "SPF", subfieldValues, 10, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordUnknownResource_View(t *testing.T) {
	var resourceName = "nios_dns_record_unknown.test_view"
	var v dns.RecordUnknown
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record-unknown") + ".com"
	subfieldValues := []map[string]any{
		{
			"field_type":     "T",
			"field_value":    "example-text",
			"include_length": "8_BIT",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordUnknownView(zoneFqdn, name, "SPF", subfieldValues, "custom_view_1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "custom_view_1"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordUnknownView(zoneFqdn, name, "SPF", subfieldValues, "custom_view_2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordUnknownExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "custom_view_2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordUnknownExists(ctx context.Context, resourceName string, v *dns.RecordUnknown) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			RecordUnknownAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordUnknown).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordUnknownResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordUnknownResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordUnknownDestroy(ctx context.Context, v *dns.RecordUnknown) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			RecordUnknownAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordUnknown).
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

func testAccCheckRecordUnknownDisappears(ctx context.Context, v *dns.RecordUnknown) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			RecordUnknownAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordUnknownBasicConfig(zoneFqdn, name, recordType string, subfieldValues []map[string]any) string {
	subfieldValuesHCL := utils.ConvertSliceOfMapsToHCL(subfieldValues)
	config := fmt.Sprintf(`
resource "nios_dns_record_unknown" "test" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    record_type = %q
    subfield_values = %s
}
`, name, recordType, subfieldValuesHCL)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordUnknownComment(zoneFqdn, name, recordType string, subfieldValues []map[string]any, comment string) string {
	subfieldValuesHCL := utils.ConvertSliceOfMapsToHCL(subfieldValues)
	config := fmt.Sprintf(`
resource "nios_dns_record_unknown" "test_comment" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    record_type = %q
    subfield_values = %s
    comment = %q
}
`, name, recordType, subfieldValuesHCL, comment)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordUnknownCreator(zoneFqdn, name, recordType string, subfieldValues []map[string]any, creator string) string {
	subfieldValuesHCL := utils.ConvertSliceOfMapsToHCL(subfieldValues)
	config := fmt.Sprintf(`
resource "nios_dns_record_unknown" "test_creator" {
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    record_type = %q
    subfield_values = %s
    creator = %q
}
`, name, recordType, subfieldValuesHCL, creator)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordUnknownDisable(zoneFqdn, name, recordType string, subfieldValues []map[string]any, disable string) string {
	subfieldValuesHCL := utils.ConvertSliceOfMapsToHCL(subfieldValues)
	config := fmt.Sprintf(`
resource "nios_dns_record_unknown" "test_disable" {
    name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    record_type = %q
    subfield_values = %s
    disable = %q
}
`, name, recordType, subfieldValuesHCL, disable)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordUnknownEnableHostNamePolicy(zoneFqdn, name, recordType string, subfieldValues []map[string]any, enableHostNamePolicy string) string {
	subfieldValuesHCL := utils.ConvertSliceOfMapsToHCL(subfieldValues)
	config := fmt.Sprintf(`
resource "nios_dns_record_unknown" "test_enable_host_name_policy" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    record_type = %q
    subfield_values = %s
    enable_host_name_policy = %q
}
`, name, recordType, subfieldValuesHCL, enableHostNamePolicy)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordUnknownExtAttrs(zoneFqdn, name, recordType string, subfieldValues []map[string]any, extAttrs map[string]string) string {
	subfieldValuesHCL := utils.ConvertSliceOfMapsToHCL(subfieldValues)
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	config := fmt.Sprintf(`
resource "nios_dns_record_unknown" "test_extattrs" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    record_type = %q
    subfield_values = %s
    extattrs = %s
}
`, name, recordType, subfieldValuesHCL, extattrsStr)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordUnknownName(zoneFqdn, name, recordType string, subfieldValues []map[string]any) string {
	subfieldValuesHCL := utils.ConvertSliceOfMapsToHCL(subfieldValues)
	config := fmt.Sprintf(`
resource "nios_dns_record_unknown" "test_name" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    record_type = %q
    subfield_values = %s
}
`, name, recordType, subfieldValuesHCL)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordUnknownRecordType(zoneFqdn, name, recordType string, subfieldValues []map[string]any) string {
	subfieldValuesHCL := utils.ConvertSliceOfMapsToHCL(subfieldValues)
	config := fmt.Sprintf(`
resource "nios_dns_record_unknown" "test_record_type" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    record_type = %q
    subfield_values = %s
}
`, name, recordType, subfieldValuesHCL)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordUnknownSubfieldValues(zoneFqdn, name, recordType string, subfieldValues []map[string]any) string {
	subfieldValuesHCL := utils.ConvertSliceOfMapsToHCL(subfieldValues)
	config := fmt.Sprintf(`
resource "nios_dns_record_unknown" "test_subfield_values" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    record_type = %q
    subfield_values = %s
}
`, name, recordType, subfieldValuesHCL)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordUnknownTtl(zoneFqdn, name, recordType string, subfieldValues []map[string]any, ttl int, useTtl string) string {
	subfieldValuesHCL := utils.ConvertSliceOfMapsToHCL(subfieldValues)
	config := fmt.Sprintf(`
resource "nios_dns_record_unknown" "test_ttl" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    record_type = %q
    subfield_values = %s
    ttl = %d
	use_ttl = %q
}
`, name, recordType, subfieldValuesHCL, ttl, useTtl)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordUnknownUseTtl(zoneFqdn, name, recordType string, subfieldValues []map[string]any, ttl int, useTtl string) string {
	subfieldValuesHCL := utils.ConvertSliceOfMapsToHCL(subfieldValues)
	config := fmt.Sprintf(`
resource "nios_dns_record_unknown" "test_use_ttl" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    record_type = %q
    subfield_values = %s
	ttl = %d
    use_ttl = %q
}
`, name, recordType, subfieldValuesHCL, ttl, useTtl)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordUnknownView(zoneFqdn, name, recordType string, subfieldValues []map[string]any, view string) string {
	subfieldValuesHCL := utils.ConvertSliceOfMapsToHCL(subfieldValues)
	config := fmt.Sprintf(`
resource "nios_dns_record_unknown" "test_view" {
	name = "${%q}.${nios_dns_zone_auth.test.fqdn}"
    record_type = %q
    subfield_values = %s
    view = %q
}
`, name, recordType, subfieldValuesHCL, view)
	return strings.Join([]string{testAccBaseWithZoneandView(zoneFqdn, view), config}, "")
}
