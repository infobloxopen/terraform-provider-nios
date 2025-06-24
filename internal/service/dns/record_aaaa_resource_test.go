package dns_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/acctest"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/utils"
)

// TODO : Add readable attributes for the resource
var readableAttributesForRecordAaaa = ""

func TestAccRecordAaaaResource_basic(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test"
	var v dns.RecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_disappears(t *testing.T) {
	resourceName := "nios_dns_record_aaaa.test"
	var v dns.RecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordAaaaDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordAaaaBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					testAccCheckRecordAaaaDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRecordAaaaResource_Ref(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test__ref"
	var v dns.RecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "_ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "_ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_AwsRte53RecordInfo(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_aws_rte53_record_info"
	var v dns.RecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaAwsRte53RecordInfo("AWS_RTE53_RECORD_INFO_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "aws_rte53_record_info", "AWS_RTE53_RECORD_INFO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaAwsRte53RecordInfo("AWS_RTE53_RECORD_INFO_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "aws_rte53_record_info", "AWS_RTE53_RECORD_INFO_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_CloudInfo(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_cloud_info"
	var v dns.RecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaCloudInfo("CLOUD_INFO_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_info", "CLOUD_INFO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaCloudInfo("CLOUD_INFO_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_info", "CLOUD_INFO_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_comment"
	var v dns.RecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_Creator(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_creator"
	var v dns.RecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaCreator("CREATOR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "CREATOR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaCreator("CREATOR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "CREATOR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_DdnsPrincipal(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_ddns_principal"
	var v dns.RecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaDdnsPrincipal("DDNS_PRINCIPAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "DDNS_PRINCIPAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaDdnsPrincipal("DDNS_PRINCIPAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "DDNS_PRINCIPAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_DdnsProtected(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_ddns_protected"
	var v dns.RecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaDdnsProtected("DDNS_PROTECTED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "DDNS_PROTECTED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaDdnsProtected("DDNS_PROTECTED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "DDNS_PROTECTED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_disable"
	var v dns.RecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_DiscoveredData(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_discovered_data"
	var v dns.RecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaDiscoveredData("DISCOVERED_DATA_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovered_data", "DISCOVERED_DATA_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaDiscoveredData("DISCOVERED_DATA_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "discovered_data", "DISCOVERED_DATA_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_extattrs"
	var v dns.RecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_ForbidReclamation(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_forbid_reclamation"
	var v dns.RecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaForbidReclamation("FORBID_RECLAMATION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "FORBID_RECLAMATION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaForbidReclamation("FORBID_RECLAMATION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "FORBID_RECLAMATION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_Ipv6addr(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_ipv6addr"
	var v dns.RecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaIpv6addr("IPV6ADDR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", "IPV6ADDR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaIpv6addr("IPV6ADDR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6addr", "IPV6ADDR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_FuncCall(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_func_call"
	var v dns.RecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaFuncCall("FUNC_CALL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "func_call", "FUNC_CALL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaFuncCall("FUNC_CALL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "func_call", "FUNC_CALL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_MsAdUserData(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_ms_ad_user_data"
	var v dns.RecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaMsAdUserData("MS_AD_USER_DATA_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_user_data", "MS_AD_USER_DATA_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaMsAdUserData("MS_AD_USER_DATA_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_user_data", "MS_AD_USER_DATA_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_Name(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_name"
	var v dns.RecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_RemoveAssociatedPtr(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_remove_associated_ptr"
	var v dns.RecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaRemoveAssociatedPtr("REMOVE_ASSOCIATED_PTR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remove_associated_ptr", "REMOVE_ASSOCIATED_PTR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaRemoveAssociatedPtr("REMOVE_ASSOCIATED_PTR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remove_associated_ptr", "REMOVE_ASSOCIATED_PTR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_ttl"
	var v dns.RecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaTtl("TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaTtl("TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_use_ttl"
	var v dns.RecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaUseTtl("USE_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "USE_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaUseTtl("USE_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "USE_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordAaaaResource_View(t *testing.T) {
	var resourceName = "nios_dns_record_aaaa.test_view"
	var v dns.RecordAaaa

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAaaaView("VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAaaaView("VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordAaaaExists(ctx context.Context, resourceName string, v *dns.RecordAaaa) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			RecordAaaaAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordAaaa).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordAaaaResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordAaaaResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordAaaaDestroy(ctx context.Context, v *dns.RecordAaaa) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			RecordAaaaAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordAaaa).
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

func testAccCheckRecordAaaaDisappears(ctx context.Context, v *dns.RecordAaaa) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			RecordAaaaAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordAaaaBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test" {
}
`)
}

func testAccRecordAaaaRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test__ref" {
    _ref = %q
}
`, ref)
}

func testAccRecordAaaaAwsRte53RecordInfo(awsRte53RecordInfo string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_aws_rte53_record_info" {
    aws_rte53_record_info = %q
}
`, awsRte53RecordInfo)
}

func testAccRecordAaaaCloudInfo(cloudInfo string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_cloud_info" {
    cloud_info = %q
}
`, cloudInfo)
}

func testAccRecordAaaaComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccRecordAaaaCreator(creator string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_creator" {
    creator = %q
}
`, creator)
}

func testAccRecordAaaaDdnsPrincipal(ddnsPrincipal string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_ddns_principal" {
    ddns_principal = %q
}
`, ddnsPrincipal)
}

func testAccRecordAaaaDdnsProtected(ddnsProtected string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_ddns_protected" {
    ddns_protected = %q
}
`, ddnsProtected)
}

func testAccRecordAaaaDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccRecordAaaaDiscoveredData(discoveredData string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_discovered_data" {
    discovered_data = %q
}
`, discoveredData)
}

func testAccRecordAaaaExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccRecordAaaaForbidReclamation(forbidReclamation string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_forbid_reclamation" {
    forbid_reclamation = %q
}
`, forbidReclamation)
}

func testAccRecordAaaaIpv6addr(ipv6addr string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_ipv6addr" {
    ipv6addr = %q
}
`, ipv6addr)
}

func testAccRecordAaaaFuncCall(funcCall string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_func_call" {
    func_call = %q
}
`, funcCall)
}

func testAccRecordAaaaMsAdUserData(msAdUserData string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_ms_ad_user_data" {
    ms_ad_user_data = %q
}
`, msAdUserData)
}

func testAccRecordAaaaName(name string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_name" {
    name = %q
}
`, name)
}

func testAccRecordAaaaRemoveAssociatedPtr(removeAssociatedPtr string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_remove_associated_ptr" {
    remove_associated_ptr = %q
}
`, removeAssociatedPtr)
}

func testAccRecordAaaaTtl(ttl string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_ttl" {
    ttl = %q
}
`, ttl)
}

func testAccRecordAaaaUseTtl(useTtl string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_use_ttl" {
    use_ttl = %q
}
`, useTtl)
}

func testAccRecordAaaaView(view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test_view" {
    view = %q
}
`, view)
}
