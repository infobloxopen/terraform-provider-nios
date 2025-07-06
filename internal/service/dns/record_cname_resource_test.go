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
var readableAttributesForRecordCname = "aws_rte53_record_info,canonical,cloud_info,comment,creation_time,creator,ddns_principal,ddns_protected,disable,dns_canonical,dns_name,extattrs,forbid_reclamation,last_queried,name,reclaimable,shared_record_group,ttl,use_ttl,view,zone"

func TestAccRecordCnameResource_basic(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test"
	var v dns.RecordCname
	canonical := acctest.RandomNameWithPrefix("test-cname-") + ".example.com"
	//name := acctest.RandomNameWithPrefix("test-alias-") + ".example.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameBasicConfig(canonical, "example_record.example.com", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

//func TestAccRecordCnameResource_disappears(t *testing.T) {
//	resourceName := "nios_dns_record_cname.test"
//	var v dns.RecordCname
//
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:                 func() { acctest.PreCheck(t) },
//		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
//		CheckDestroy:             testAccCheckRecordCnameDestroy(context.Background(), &v),
//		Steps: []resource.TestStep{
//			{
//				Config: testAccRecordCnameBasicConfig(""),
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
//					testAccCheckRecordCnameDisappears(context.Background(), &v),
//				),
//				ExpectNonEmptyPlan: true,
//			},
//		},
//	})
//}

func TestAccRecordCnameResource_Ref(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test__ref"
	var v dns.RecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "_ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "_ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_AwsRte53RecordInfo(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_aws_rte53_record_info"
	var v dns.RecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameAwsRte53RecordInfo("AWS_RTE53_RECORD_INFO_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "aws_rte53_record_info", "AWS_RTE53_RECORD_INFO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameAwsRte53RecordInfo("AWS_RTE53_RECORD_INFO_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "aws_rte53_record_info", "AWS_RTE53_RECORD_INFO_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_Canonical(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_canonical"
	var v dns.RecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameCanonical("CANONICAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", "CANONICAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameCanonical("CANONICAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "canonical", "CANONICAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_CloudInfo(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_cloud_info"
	var v dns.RecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameCloudInfo("CLOUD_INFO_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_info", "CLOUD_INFO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameCloudInfo("CLOUD_INFO_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_info", "CLOUD_INFO_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_comment"
	var v dns.RecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_Creator(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_creator"
	var v dns.RecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameCreator("CREATOR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "CREATOR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameCreator("CREATOR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "creator", "CREATOR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_DdnsPrincipal(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_ddns_principal"
	var v dns.RecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameDdnsPrincipal("DDNS_PRINCIPAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "DDNS_PRINCIPAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameDdnsPrincipal("DDNS_PRINCIPAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal", "DDNS_PRINCIPAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_DdnsProtected(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_ddns_protected"
	var v dns.RecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameDdnsProtected("DDNS_PROTECTED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "DDNS_PROTECTED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameDdnsProtected("DDNS_PROTECTED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_protected", "DDNS_PROTECTED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_disable"
	var v dns.RecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_Extattrs(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_extattrs"
	var v dns.RecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameExtattrs("EXTATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXTATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameExtattrs("EXTATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXTATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_ForbidReclamation(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_forbid_reclamation"
	var v dns.RecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameForbidReclamation("FORBID_RECLAMATION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "FORBID_RECLAMATION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameForbidReclamation("FORBID_RECLAMATION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forbid_reclamation", "FORBID_RECLAMATION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_Name(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_name"
	var v dns.RecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_Ttl(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_ttl"
	var v dns.RecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameTtl("TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameTtl("TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_UseTtl(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_use_ttl"
	var v dns.RecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameUseTtl("USE_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "USE_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameUseTtl("USE_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "USE_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCnameResource_View(t *testing.T) {
	var resourceName = "nios_dns_record_cname.test_view"
	var v dns.RecordCname

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCnameView("VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCnameView("VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRecordCnameExists(ctx context.Context, resourceName string, v *dns.RecordCname) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			RecordCnameAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRecordCname).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRecordCnameResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRecordCnameResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRecordCnameDestroy(ctx context.Context, v *dns.RecordCname) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			RecordCnameAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRecordCname).
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

func testAccCheckRecordCnameDisappears(ctx context.Context, v *dns.RecordCname) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			RecordCnameAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRecordCnameBasicConfig(name, ipV4Addr, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test" {
	canonical = %q
	name = %q
	view = %q
}
`, name, ipV4Addr, view)
}

func testAccRecordCnameRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test__ref" {
    _ref = %q
}
`, ref)
}

func testAccRecordCnameAwsRte53RecordInfo(awsRte53RecordInfo string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_aws_rte53_record_info" {
    aws_rte53_record_info = %q
}
`, awsRte53RecordInfo)
}

func testAccRecordCnameCanonical(canonical string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_canonical" {
    canonical = %q
}
`, canonical)
}

func testAccRecordCnameCloudInfo(cloudInfo string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_cloud_info" {
    cloud_info = %q
}
`, cloudInfo)
}

func testAccRecordCnameComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccRecordCnameCreator(creator string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_creator" {
    creator = %q
}
`, creator)
}

func testAccRecordCnameDdnsPrincipal(ddnsPrincipal string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_ddns_principal" {
    ddns_principal = %q
}
`, ddnsPrincipal)
}

func testAccRecordCnameDdnsProtected(ddnsProtected string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_ddns_protected" {
    ddns_protected = %q
}
`, ddnsProtected)
}

func testAccRecordCnameDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccRecordCnameExtattrs(extattrs string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_extattrs" {
    extattrs = %q
}
`, extattrs)
}

func testAccRecordCnameForbidReclamation(forbidReclamation string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_forbid_reclamation" {
    forbid_reclamation = %q
}
`, forbidReclamation)
}

func testAccRecordCnameName(name string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_name" {
    name = %q
}
`, name)
}

func testAccRecordCnameTtl(ttl string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_ttl" {
    ttl = %q
}
`, ttl)
}

func testAccRecordCnameUseTtl(useTtl string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_use_ttl" {
    use_ttl = %q
}
`, useTtl)
}

func testAccRecordCnameView(view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test_view" {
    view = %q
}
`, view)
}
