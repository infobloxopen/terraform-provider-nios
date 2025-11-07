package threatprotection_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/threatprotection"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForThreatprotectionProfile = "comment,current_ruleset,disable_multiple_dns_tcp_request,events_per_second_per_rule,extattrs,members,name,use_current_ruleset,use_disable_multiple_dns_tcp_request,use_events_per_second_per_rule"

func TestAccThreatprotectionProfileResource_basic(t *testing.T) {
	var resourceName = "nios_threatprotection_profile.test"
	var v threatprotection.ThreatprotectionProfile

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionProfileBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatprotectionProfileResource_disappears(t *testing.T) {
	resourceName := "nios_threatprotection_profile.test"
	var v threatprotection.ThreatprotectionProfile

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckThreatprotectionProfileDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccThreatprotectionProfileBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					testAccCheckThreatprotectionProfileDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccThreatprotectionProfileResource_Ref(t *testing.T) {
	var resourceName = "nios_threatprotection_profile.test_ref"
	var v threatprotection.ThreatprotectionProfile

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionProfileRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatprotectionProfileRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatprotectionProfileResource_Comment(t *testing.T) {
	var resourceName = "nios_threatprotection_profile.test_comment"
	var v threatprotection.ThreatprotectionProfile

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionProfileComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatprotectionProfileComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatprotectionProfileResource_CurrentRuleset(t *testing.T) {
	var resourceName = "nios_threatprotection_profile.test_current_ruleset"
	var v threatprotection.ThreatprotectionProfile

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionProfileCurrentRuleset("CURRENT_RULESET_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "current_ruleset", "CURRENT_RULESET_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatprotectionProfileCurrentRuleset("CURRENT_RULESET_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "current_ruleset", "CURRENT_RULESET_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatprotectionProfileResource_DisableMultipleDnsTcpRequest(t *testing.T) {
	var resourceName = "nios_threatprotection_profile.test_disable_multiple_dns_tcp_request"
	var v threatprotection.ThreatprotectionProfile

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionProfileDisableMultipleDnsTcpRequest("DISABLE_MULTIPLE_DNS_TCP_REQUEST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_multiple_dns_tcp_request", "DISABLE_MULTIPLE_DNS_TCP_REQUEST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatprotectionProfileDisableMultipleDnsTcpRequest("DISABLE_MULTIPLE_DNS_TCP_REQUEST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_multiple_dns_tcp_request", "DISABLE_MULTIPLE_DNS_TCP_REQUEST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatprotectionProfileResource_EventsPerSecondPerRule(t *testing.T) {
	var resourceName = "nios_threatprotection_profile.test_events_per_second_per_rule"
	var v threatprotection.ThreatprotectionProfile

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionProfileEventsPerSecondPerRule("EVENTS_PER_SECOND_PER_RULE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "events_per_second_per_rule", "EVENTS_PER_SECOND_PER_RULE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatprotectionProfileEventsPerSecondPerRule("EVENTS_PER_SECOND_PER_RULE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "events_per_second_per_rule", "EVENTS_PER_SECOND_PER_RULE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatprotectionProfileResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_threatprotection_profile.test_extattrs"
	var v threatprotection.ThreatprotectionProfile

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionProfileExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatprotectionProfileExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatprotectionProfileResource_Members(t *testing.T) {
	var resourceName = "nios_threatprotection_profile.test_members"
	var v threatprotection.ThreatprotectionProfile

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionProfileMembers("MEMBERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "members", "MEMBERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatprotectionProfileMembers("MEMBERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "members", "MEMBERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatprotectionProfileResource_Name(t *testing.T) {
	var resourceName = "nios_threatprotection_profile.test_name"
	var v threatprotection.ThreatprotectionProfile

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionProfileName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatprotectionProfileName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatprotectionProfileResource_SourceMember(t *testing.T) {
	var resourceName = "nios_threatprotection_profile.test_source_member"
	var v threatprotection.ThreatprotectionProfile

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionProfileSourceMember("SOURCE_MEMBER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "source_member", "SOURCE_MEMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatprotectionProfileSourceMember("SOURCE_MEMBER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "source_member", "SOURCE_MEMBER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatprotectionProfileResource_SourceProfile(t *testing.T) {
	var resourceName = "nios_threatprotection_profile.test_source_profile"
	var v threatprotection.ThreatprotectionProfile

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionProfileSourceProfile("SOURCE_PROFILE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "source_profile", "SOURCE_PROFILE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatprotectionProfileSourceProfile("SOURCE_PROFILE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "source_profile", "SOURCE_PROFILE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatprotectionProfileResource_UseCurrentRuleset(t *testing.T) {
	var resourceName = "nios_threatprotection_profile.test_use_current_ruleset"
	var v threatprotection.ThreatprotectionProfile

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionProfileUseCurrentRuleset("USE_CURRENT_RULESET_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_current_ruleset", "USE_CURRENT_RULESET_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatprotectionProfileUseCurrentRuleset("USE_CURRENT_RULESET_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_current_ruleset", "USE_CURRENT_RULESET_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatprotectionProfileResource_UseDisableMultipleDnsTcpRequest(t *testing.T) {
	var resourceName = "nios_threatprotection_profile.test_use_disable_multiple_dns_tcp_request"
	var v threatprotection.ThreatprotectionProfile

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionProfileUseDisableMultipleDnsTcpRequest("USE_DISABLE_MULTIPLE_DNS_TCP_REQUEST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_disable_multiple_dns_tcp_request", "USE_DISABLE_MULTIPLE_DNS_TCP_REQUEST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatprotectionProfileUseDisableMultipleDnsTcpRequest("USE_DISABLE_MULTIPLE_DNS_TCP_REQUEST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_disable_multiple_dns_tcp_request", "USE_DISABLE_MULTIPLE_DNS_TCP_REQUEST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccThreatprotectionProfileResource_UseEventsPerSecondPerRule(t *testing.T) {
	var resourceName = "nios_threatprotection_profile.test_use_events_per_second_per_rule"
	var v threatprotection.ThreatprotectionProfile

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccThreatprotectionProfileUseEventsPerSecondPerRule("USE_EVENTS_PER_SECOND_PER_RULE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_events_per_second_per_rule", "USE_EVENTS_PER_SECOND_PER_RULE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccThreatprotectionProfileUseEventsPerSecondPerRule("USE_EVENTS_PER_SECOND_PER_RULE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckThreatprotectionProfileExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_events_per_second_per_rule", "USE_EVENTS_PER_SECOND_PER_RULE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckThreatprotectionProfileExists(ctx context.Context, resourceName string, v *threatprotection.ThreatprotectionProfile) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.ThreatProtectionAPI.
			ThreatprotectionProfileAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForThreatprotectionProfile).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetThreatprotectionProfileResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetThreatprotectionProfileResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckThreatprotectionProfileDestroy(ctx context.Context, v *threatprotection.ThreatprotectionProfile) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.ThreatProtectionAPI.
			ThreatprotectionProfileAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForThreatprotectionProfile).
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

func testAccCheckThreatprotectionProfileDisappears(ctx context.Context, v *threatprotection.ThreatprotectionProfile) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.ThreatProtectionAPI.
			ThreatprotectionProfileAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccThreatprotectionProfileBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_threatprotection_profile" "test" {
}
`)
}

func testAccThreatprotectionProfileRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_profile" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccThreatprotectionProfileComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_profile" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccThreatprotectionProfileCurrentRuleset(currentRuleset string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_profile" "test_current_ruleset" {
    current_ruleset = %q
}
`, currentRuleset)
}

func testAccThreatprotectionProfileDisableMultipleDnsTcpRequest(disableMultipleDnsTcpRequest string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_profile" "test_disable_multiple_dns_tcp_request" {
    disable_multiple_dns_tcp_request = %q
}
`, disableMultipleDnsTcpRequest)
}

func testAccThreatprotectionProfileEventsPerSecondPerRule(eventsPerSecondPerRule string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_profile" "test_events_per_second_per_rule" {
    events_per_second_per_rule = %q
}
`, eventsPerSecondPerRule)
}

func testAccThreatprotectionProfileExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_profile" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccThreatprotectionProfileMembers(members string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_profile" "test_members" {
    members = %q
}
`, members)
}

func testAccThreatprotectionProfileName(name string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_profile" "test_name" {
    name = %q
}
`, name)
}

func testAccThreatprotectionProfileSourceMember(sourceMember string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_profile" "test_source_member" {
    source_member = %q
}
`, sourceMember)
}

func testAccThreatprotectionProfileSourceProfile(sourceProfile string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_profile" "test_source_profile" {
    source_profile = %q
}
`, sourceProfile)
}

func testAccThreatprotectionProfileUseCurrentRuleset(useCurrentRuleset string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_profile" "test_use_current_ruleset" {
    use_current_ruleset = %q
}
`, useCurrentRuleset)
}

func testAccThreatprotectionProfileUseDisableMultipleDnsTcpRequest(useDisableMultipleDnsTcpRequest string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_profile" "test_use_disable_multiple_dns_tcp_request" {
    use_disable_multiple_dns_tcp_request = %q
}
`, useDisableMultipleDnsTcpRequest)
}

func testAccThreatprotectionProfileUseEventsPerSecondPerRule(useEventsPerSecondPerRule string) string {
	return fmt.Sprintf(`
resource "nios_threatprotection_profile" "test_use_events_per_second_per_rule" {
    use_events_per_second_per_rule = %q
}
`, useEventsPerSecondPerRule)
}
