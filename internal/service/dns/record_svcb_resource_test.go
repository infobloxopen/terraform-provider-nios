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

var readableAttributesForRecordSvcb = "aws_rte53_record_info,cloud_info,comment,creation_time,creator,ddns_principal,ddns_protected,disable,extattrs,forbid_reclamation,last_queried,name,priority,reclaimable,svc_parameters,target_name,ttl,use_ttl,view,zone"

func TestAccRecordSvcbResource_basic(t *testing.T) {
	var resourceName = "nios_dns_record_svcb.test"
	var v dns.RecordSvcb
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record_svcb")
	priority := "10"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSvcbBasicConfig(zoneFqdn, name, priority),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name+"."+zoneFqdn),
					resource.TestCheckResourceAttr(resourceName, "priority", priority),
					resource.TestCheckResourceAttr(resourceName, "target_name", zoneFqdn),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", ""),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "false"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSvcbResource_disappears(t *testing.T) {
	resourceName := "nios_dns_record_svcb.test"
	var v dns.RecordSvcb
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record_svcb")
	priority := "10"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordSvcbDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordSvcbBasicConfig(zoneFqdn, name, priority),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					testAccCheckRecordSvcbDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordSvcbResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_record_svcb.test_comment"
	var v dns.RecordSvcb
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record_svcb")
	priority := "10"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSvcbComment(zoneFqdn, name, priority, "Example SVCB Record"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Example SVCB Record"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSvcbComment(zoneFqdn, name, priority, "Example SVCB Record Updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Example SVCB Record Updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSvcbResource_Creator(t *testing.T) {
	var resourceName = "nios_dns_record_svcb.test_creator"
	var v dns.RecordSvcb
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record_svcb")
	priority := "10"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSvcbCreator(zoneFqdn, name, priority, "STATIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "STATIC"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSvcbCreator(zoneFqdn, name, priority, "DYNAMIC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "DYNAMIC"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSvcbResource_DdnsPrincipal(t *testing.T) {
	var resourceName = "nios_dns_record_svcb.test_ddns_principal"
	var v dns.RecordSvcb
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record_svcb")
	priority := "10"
	ddnsPrincipal1 := acctest.RandomNameWithPrefix("ddns_principal")
	ddnsPrincipal2 := acctest.RandomNameWithPrefix("ddns_principal")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSvcbDdnsPrincipal(zoneFqdn, name, priority, ddnsPrincipal1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", ddnsPrincipal1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSvcbDdnsPrincipal(zoneFqdn, name, priority, ddnsPrincipal2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", ddnsPrincipal2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSvcbResource_DdnsProtected(t *testing.T) {
	var resourceName = "nios_dns_record_svcb.test_ddns_protected"
	var v dns.RecordSvcb
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record_svcb")
	priority := "10"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSvcbDdnsProtected(zoneFqdn, name, priority, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSvcbDdnsProtected(zoneFqdn, name, priority, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSvcbResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_record_svcb.test_disable"
	var v dns.RecordSvcb
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record_svcb")
	priority := "10"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSvcbDisable(zoneFqdn, name, priority, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSvcbDisable(zoneFqdn, name, priority, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSvcbResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_record_svcb.test_extattrs"
	var v dns.RecordSvcb
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record_svcb")
	priority := "10"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSvcbExtAttrs(zoneFqdn, name, priority, map[string]any{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSvcbExtAttrs(zoneFqdn, name, priority, map[string]any{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSvcbResource_ForbidReclamation(t *testing.T) {
	var resourceName = "nios_dns_record_svcb.test_forbid_reclamation"
	var v dns.RecordSvcb
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record_svcb")
	priority := "10"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSvcbForbidReclamation(zoneFqdn, name, priority, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSvcbForbidReclamation(zoneFqdn, name, priority, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSvcbResource_Name(t *testing.T) {
	var resourceName = "nios_dns_record_svcb.test_name"
	var v dns.RecordSvcb
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name1 := acctest.RandomNameWithPrefix("record_svcb")
	name2 := acctest.RandomNameWithPrefix("record_svcb")
	priority := "10"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSvcbName(zoneFqdn, name1, priority),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1+"."+zoneFqdn),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSvcbName(zoneFqdn, name2, priority),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2+"."+zoneFqdn),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSvcbResource_Priority(t *testing.T) {
	var resourceName = "nios_dns_record_svcb.test_priority"
	var v dns.RecordSvcb
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record_svcb")
	priority1 := "10"
	priority2 := "20"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSvcbPriority(zoneFqdn, name, priority1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "priority", priority1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSvcbPriority(zoneFqdn, name, priority2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "priority", priority2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSvcbResource_SvcParameters(t *testing.T) {
	var resourceName = "nios_dns_record_svcb.test_svc_parameters"
	var v dns.RecordSvcb
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record_svcb")
	priority := "10"
	svcParameters1 := []map[string]any{
		{
			"mandatory": true,
			"svc_key":   "port",
			"svc_value": []string{"110"},
		},
		{
			"mandatory": false,
			"svc_key":   "ipv4hint",
			"svc_value": []string{"11.11.1.1"},
		},
		{
			"mandatory": false,
			"svc_key":   "ipv6hint",
			"svc_value": []string{"123::99:0"},
		},
	}
	svcParameters2 := []map[string]any{
		{
			"mandatory": true,
			"svc_key":   "port",
			"svc_value": []string{"11"},
		},
		{
			"mandatory": false,
			"svc_key":   "ipv6hint",
			"svc_value": []string{"124::92:0"},
		},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSvcbSvcParameters(zoneFqdn, name, priority, svcParameters1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.0.mandatory", "true"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.0.svc_key", "port"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.0.svc_value.0", "110"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.1.mandatory", "false"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.1.svc_key", "ipv4hint"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.1.svc_value.0", "11.11.1.1"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.2.svc_key", "ipv6hint"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.2.svc_value.0", "123::99:0"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSvcbSvcParameters(zoneFqdn, name, priority, svcParameters2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.0.mandatory", "true"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.0.svc_key", "port"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.0.svc_value.0", "11"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.1.svc_key", "ipv6hint"),
					resource.TestCheckResourceAttr(resourceName, "svc_parameters.1.svc_value.0", "124::92:0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSvcbResource_TargetName(t *testing.T) {
	var resourceName = "nios_dns_record_svcb.test_target_name"
	var v dns.RecordSvcb
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record_svcb")
	priority := "10"
	targetName1 := "example." + zoneFqdn
	targetName2 := "example1." + zoneFqdn
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSvcbTargetName(zoneFqdn, name, priority, targetName1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "target_name", targetName1),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSvcbTargetName(zoneFqdn, name, priority, targetName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "target_name", targetName2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSvcbResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_record_svcb.test_ttl"
	var v dns.RecordSvcb
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record_svcb")
	priority := "10"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSvcbTtl(zoneFqdn, name, priority, "10", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSvcbTtl(zoneFqdn, name, priority, "20", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordSvcbResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_record_svcb.test_use_ttl"
	var v dns.RecordSvcb
	zoneFqdn := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomNameWithPrefix("record_svcb")
	priority := "10"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSvcbUseTtl(zoneFqdn, name, priority, "30", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSvcbUseTtl(zoneFqdn, name, priority, "30", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordSvcbExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordSvcbExists(ctx context.Context, resourceName string, v *dns.RecordSvcb) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			RecordSvcbAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordSvcb).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordSvcbResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordSvcbResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordSvcbDestroy(ctx context.Context, v *dns.RecordSvcb) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			RecordSvcbAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordSvcb).
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

func testAccCheckRecordSvcbDisappears(ctx context.Context, v *dns.RecordSvcb) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			RecordSvcbAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordSvcbBasicConfig(zoneFqdn, name, priority string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_svcb" "test" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
}
`, name, zoneFqdn, priority)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordSvcbComment(zoneFqdn, name, priority, comment string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_svcb" "test_comment" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
   comment = %q
}
`, name, zoneFqdn, priority, comment)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordSvcbCreator(zoneFqdn, name, priority, creator string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_svcb" "test_creator" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
   creator = %q
}
`, name, zoneFqdn, priority, creator)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordSvcbDdnsPrincipal(zoneFqdn, name, priority, ddnsPrincipal string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_svcb" "test_ddns_principal" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
   ddns_principal = %q
}
`, name, zoneFqdn, priority, ddnsPrincipal)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordSvcbDdnsProtected(zoneFqdn, name, priority, ddnsProtected string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_svcb" "test_ddns_protected" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
   ddns_protected = %q
}
`, name, zoneFqdn, priority, ddnsProtected)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordSvcbDisable(zoneFqdn, name, priority, disable string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_svcb" "test_disable" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
   disable = %q
}
`, name, zoneFqdn, priority, disable)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordSvcbExtAttrs(zoneFqdn, name, priority string, extAttrs map[string]any) string {
	extAttrsStr := utils.ConvertMapToHCL(extAttrs)
	config := fmt.Sprintf(`
resource "nios_dns_record_svcb" "test_extattrs" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
   extattrs = %s
}
`, name, zoneFqdn, priority, extAttrsStr)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordSvcbForbidReclamation(zoneFqdn, name, priority, forbidReclamation string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_svcb" "test_forbid_reclamation" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
   forbid_reclamation = %q
}
`, name, zoneFqdn, priority, forbidReclamation)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordSvcbName(zoneFqdn, name, priority string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_svcb" "test_name" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
}
`, name, zoneFqdn, priority)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordSvcbPriority(zoneFqdn, name, priority string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_svcb" "test_priority" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
}
`, name, zoneFqdn, priority)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordSvcbSvcParameters(zoneFqdn, name, priority string, svcParameters []map[string]any) string {
	svcParamsStr := utils.ConvertSliceOfMapsToHCL(svcParameters)
	config := fmt.Sprintf(`
resource "nios_dns_record_svcb" "test_svc_parameters" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
   svc_parameters = %s
}
`, name, zoneFqdn, priority, svcParamsStr)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordSvcbTargetName(zoneFqdn, name, priority, targetName string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_svcb" "test_target_name" {
   name = "%s.%s"
   priority = %q
   target_name = %q
   depends_on = [nios_dns_zone_auth.test]
}
`, name, zoneFqdn, priority, targetName)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordSvcbTtl(zoneFqdn, name, priority, ttl, useTtl string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_svcb" "test_ttl" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
   ttl = %q
   use_ttl = %q
}
`, name, zoneFqdn, priority, ttl, useTtl)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordSvcbUseTtl(zoneFqdn, name, priority, ttl, useTtl string) string {
	config := fmt.Sprintf(`
resource "nios_dns_record_svcb" "test_use_ttl" {
   name = "%s.%s"
   priority = %q
   target_name = nios_dns_zone_auth.test.fqdn
   ttl = %q
   use_ttl = %q
}
`, name, zoneFqdn, priority, ttl, useTtl)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}
