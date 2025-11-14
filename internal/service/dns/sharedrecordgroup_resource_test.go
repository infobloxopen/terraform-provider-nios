package dns_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForSharedrecordgroup = "comment,extattrs,name,record_name_policy,use_record_name_policy,zone_associations"

func TestAccSharedrecordgroupResource_basic(t *testing.T) {
	var resourceName = "nios_dns_sharedrecordgroup.test"
	var v dns.Sharedrecordgroup
	name := acctest.RandomNameWithPrefix("sharedrecordgroup")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordgroupBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "record_name_policy", ""),
					resource.TestCheckResourceAttr(resourceName, "use_record_name_policy", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordgroupResource_disappears(t *testing.T) {
	resourceName := "nios_dns_sharedrecordgroup.test"
	var v dns.Sharedrecordgroup
	name := acctest.RandomNameWithPrefix("sharedrecordgroup")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharedrecordgroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharedrecordgroupBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordgroupExists(context.Background(), resourceName, &v),
					testAccCheckSharedrecordgroupDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSharedrecordgroupResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_sharedrecordgroup.test_comment"
	var v dns.Sharedrecordgroup
	name := acctest.RandomNameWithPrefix("sharedrecordgroup")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordgroupComment(name, "shared record group comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "shared record group comment"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordgroupComment(name, "shared record group comment updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "shared record group comment updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordgroupResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_sharedrecordgroup.test_extattrs"
	var v dns.Sharedrecordgroup
	name := acctest.RandomNameWithPrefix("sharedrecordgroup")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordgroupExtAttrs(name, map[string]any{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordgroupExtAttrs(name, map[string]any{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordgroupResource_Name(t *testing.T) {
	var resourceName = "nios_dns_sharedrecordgroup.test_name"
	var v dns.Sharedrecordgroup
	name1 := acctest.RandomNameWithPrefix("sharedrecordgroup")
	name2 := acctest.RandomNameWithPrefix("sharedrecordgroup")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordgroupName(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordgroupName(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordgroupResource_RecordNamePolicy(t *testing.T) {
	var resourceName = "nios_dns_sharedrecordgroup.test_record_name_policy"
	var v dns.Sharedrecordgroup
	name := acctest.RandomNameWithPrefix("sharedrecordgroup")
	recordNamePolicy1 := "Allow Underscore"
	recordNamePolicy2 := "Allow Any"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordgroupRecordNamePolicy(name, recordNamePolicy1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "record_name_policy", recordNamePolicy1),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordgroupRecordNamePolicy(name, recordNamePolicy2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "record_name_policy", recordNamePolicy2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordgroupResource_UseRecordNamePolicy(t *testing.T) {
	var resourceName = "nios_dns_sharedrecordgroup.test_use_record_name_policy"
	var v dns.Sharedrecordgroup
	name := acctest.RandomNameWithPrefix("sharedrecordgroup")
	recordNamePolicy := "Allow Any"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordgroupUseRecordNamePolicy(name, recordNamePolicy, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_record_name_policy", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordgroupUseRecordNamePolicy(name, recordNamePolicy, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_record_name_policy", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSharedrecordgroupResource_ZoneAssociations(t *testing.T) {
	var resourceName = "nios_dns_sharedrecordgroup.test_zone_associations"
	var v dns.Sharedrecordgroup
	name := acctest.RandomNameWithPrefix("sharedrecordgroup")
	zoneFqdn1 := acctest.RandomNameWithPrefix("test-zone") + ".com"
	zoneFqdn2 := acctest.RandomNameWithPrefix("test-zone") + ".com"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSharedrecordgroupZoneAssociations(name, zoneFqdn1, "default", "test1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone_associations.0.fqdn", zoneFqdn1),
					resource.TestCheckResourceAttr(resourceName, "zone_associations.0.view", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordgroupZoneAssociations(name, zoneFqdn2, "default", "test2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone_associations.0.fqdn", zoneFqdn2),
					resource.TestCheckResourceAttr(resourceName, "zone_associations.0.view", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccSharedrecordgroupZoneAssociations(name, "", "", ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSharedrecordgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone_associations.#", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckSharedrecordgroupExists(ctx context.Context, resourceName string, v *dns.Sharedrecordgroup) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			SharedrecordgroupAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForSharedrecordgroup).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetSharedrecordgroupResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetSharedrecordgroupResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckSharedrecordgroupDestroy(ctx context.Context, v *dns.Sharedrecordgroup) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			SharedrecordgroupAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForSharedrecordgroup).
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

func testAccCheckSharedrecordgroupDisappears(ctx context.Context, v *dns.Sharedrecordgroup) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			SharedrecordgroupAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccSharedrecordgroupBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecordgroup" "test" {
	name = %q
}
`, name)
}

func testAccSharedrecordgroupComment(name, comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecordgroup" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccSharedrecordgroupExtAttrs(name string, extAttrs map[string]any) string {
	extAttrsStr := utils.ConvertMapToHCL(extAttrs)
	return fmt.Sprintf(`
resource "nios_dns_sharedrecordgroup" "test_extattrs" {
	name = %q
    extattrs = %s
}
`, name, extAttrsStr)
}

func testAccSharedrecordgroupName(name string) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecordgroup" "test_name" {
    name = %q
}
`, name)
}

func testAccSharedrecordgroupRecordNamePolicy(name, recordNamePolicy string) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecordgroup" "test_record_name_policy" {
    name = %q
    record_name_policy = %q
    use_record_name_policy = true
}
`, name, recordNamePolicy)
}

func testAccSharedrecordgroupUseRecordNamePolicy(name, recordNamePolicy string, useRecordNamePolicy bool) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecordgroup" "test_use_record_name_policy" {
    name = %q
    record_name_policy = %q
    use_record_name_policy = %t
}
`, name, recordNamePolicy, useRecordNamePolicy)
}

func testAccSharedrecordgroupZoneAssociations(name string, zone, view, parentZoneAuthResource string) string {
	var zoneAssociationsConfig, parentZoneAuthConfig string

	if zone != "" && view != "" && parentZoneAuthResource != "" {
		zoneAssociationsConfig = fmt.Sprintf(`[
            {
                fqdn = nios_dns_zone_auth.%s.fqdn
                view = nios_dns_zone_auth.%s.view
            }
        ]`, parentZoneAuthResource, parentZoneAuthResource)
		parentZoneAuthConfig = testAccParentZoneAuth(zone, view, parentZoneAuthResource)
	} else {
		// Explicitly unset zone_associations
		zoneAssociationsConfig = "null"
	}

	return fmt.Sprintf(`
%s
resource "nios_dns_sharedrecordgroup" "test_zone_associations" {
    name = %q
    zone_associations = %s
}
`, parentZoneAuthConfig, name, zoneAssociationsConfig)
}

func testAccParentZoneAuth(zone, view, testZone string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" %q {
  fqdn = %q
  view = %q
}
`, testZone, zone, view)
}
