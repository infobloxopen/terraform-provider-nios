package cloud_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/cloud"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForAwsrte53taskgroup = "account_id,comment,consolidate_zones,consolidated_view,disabled,grid_member,name,network_view,network_view_mapping_policy,role_arn,sync_child_accounts,sync_status,task_list"

func TestAccAwsrte53taskgroupResource_basic(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test"
	var v cloud.Awsrte53taskgroup
	taskGroupName := acctest.RandomNameWithPrefix("test-taskgroup")
	gridMember := "infoblox.localdomain"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupBasicConfig(taskGroupName, gridMember),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", taskGroupName),
					resource.TestCheckResourceAttr(resourceName, "grid_member", gridMember),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "sync_child_accounts", "false"),
					resource.TestCheckResourceAttr(resourceName, "network_view_mapping_policy", "AUTO_CREATE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsrte53taskgroupResource_disappears(t *testing.T) {
	resourceName := "nios_cloud_awsrte53taskgroup.test"
	var v cloud.Awsrte53taskgroup
	taskGroupName := acctest.RandomNameWithPrefix("test-taskgroup")
	gridMember := "infoblox.localdomain"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAwsrte53taskgroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAwsrte53taskgroupBasicConfig(taskGroupName, gridMember),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					testAccCheckAwsrte53taskgroupDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAwsrte53taskgroupResource_AwsAccountIdsFileToken(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_aws_account_ids_file_token"
	var v cloud.Awsrte53taskgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupAwsAccountIdsFileToken("AWS_ACCOUNT_IDS_FILE_TOKEN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "aws_account_ids_file_token", "AWS_ACCOUNT_IDS_FILE_TOKEN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsrte53taskgroupAwsAccountIdsFileToken("AWS_ACCOUNT_IDS_FILE_TOKEN_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "aws_account_ids_file_token", "AWS_ACCOUNT_IDS_FILE_TOKEN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsrte53taskgroupResource_Comment(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_comment"
	var v cloud.Awsrte53taskgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsrte53taskgroupComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsrte53taskgroupResource_ConsolidateZones(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_consolidate_zones"
	var v cloud.Awsrte53taskgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupConsolidateZones("CONSOLIDATE_ZONES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "consolidate_zones", "CONSOLIDATE_ZONES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsrte53taskgroupConsolidateZones("CONSOLIDATE_ZONES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "consolidate_zones", "CONSOLIDATE_ZONES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsrte53taskgroupResource_ConsolidatedView(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_consolidated_view"
	var v cloud.Awsrte53taskgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupConsolidatedView("CONSOLIDATED_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "consolidated_view", "CONSOLIDATED_VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsrte53taskgroupConsolidatedView("CONSOLIDATED_VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "consolidated_view", "CONSOLIDATED_VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsrte53taskgroupResource_Disabled(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_disabled"
	var v cloud.Awsrte53taskgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupDisabled("DISABLED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "DISABLED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsrte53taskgroupDisabled("DISABLED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "DISABLED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsrte53taskgroupResource_GridMember(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_grid_member"
	var v cloud.Awsrte53taskgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupGridMember("GRID_MEMBER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_member", "GRID_MEMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsrte53taskgroupGridMember("GRID_MEMBER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_member", "GRID_MEMBER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsrte53taskgroupResource_MultipleAccountsSyncPolicy(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_multiple_accounts_sync_policy"
	var v cloud.Awsrte53taskgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupMultipleAccountsSyncPolicy("MULTIPLE_ACCOUNTS_SYNC_POLICY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "multiple_accounts_sync_policy", "MULTIPLE_ACCOUNTS_SYNC_POLICY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsrte53taskgroupMultipleAccountsSyncPolicy("MULTIPLE_ACCOUNTS_SYNC_POLICY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "multiple_accounts_sync_policy", "MULTIPLE_ACCOUNTS_SYNC_POLICY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsrte53taskgroupResource_Name(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_name"
	var v cloud.Awsrte53taskgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsrte53taskgroupName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsrte53taskgroupResource_NetworkView(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_network_view"
	var v cloud.Awsrte53taskgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupNetworkView("NETWORK_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsrte53taskgroupNetworkView("NETWORK_VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsrte53taskgroupResource_NetworkViewMappingPolicy(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_network_view_mapping_policy"
	var v cloud.Awsrte53taskgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupNetworkViewMappingPolicy("NETWORK_VIEW_MAPPING_POLICY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view_mapping_policy", "NETWORK_VIEW_MAPPING_POLICY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsrte53taskgroupNetworkViewMappingPolicy("NETWORK_VIEW_MAPPING_POLICY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view_mapping_policy", "NETWORK_VIEW_MAPPING_POLICY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsrte53taskgroupResource_RoleArn(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_role_arn"
	var v cloud.Awsrte53taskgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupRoleArn("ROLE_ARN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "role_arn", "ROLE_ARN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsrte53taskgroupRoleArn("ROLE_ARN_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "role_arn", "ROLE_ARN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsrte53taskgroupResource_SyncChildAccounts(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_sync_child_accounts"
	var v cloud.Awsrte53taskgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupSyncChildAccounts("SYNC_CHILD_ACCOUNTS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sync_child_accounts", "SYNC_CHILD_ACCOUNTS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsrte53taskgroupSyncChildAccounts("SYNC_CHILD_ACCOUNTS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sync_child_accounts", "SYNC_CHILD_ACCOUNTS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsrte53taskgroupResource_TaskList(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_task_list"
	var v cloud.Awsrte53taskgroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupTaskList("TASK_LIST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "task_list", "TASK_LIST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsrte53taskgroupTaskList("TASK_LIST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "task_list", "TASK_LIST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckAwsrte53taskgroupExists(ctx context.Context, resourceName string, v *cloud.Awsrte53taskgroup) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.CloudAPI.
			Awsrte53taskgroupAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForAwsrte53taskgroup).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetAwsrte53taskgroupResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetAwsrte53taskgroupResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckAwsrte53taskgroupDestroy(ctx context.Context, v *cloud.Awsrte53taskgroup) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.CloudAPI.
			Awsrte53taskgroupAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForAwsrte53taskgroup).
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

func testAccCheckAwsrte53taskgroupDisappears(ctx context.Context, v *cloud.Awsrte53taskgroup) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.CloudAPI.
			Awsrte53taskgroupAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccAwsrte53taskgroupBasicConfig(name, gridMember string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test" {
    name                         = %q
    grid_member                  = %q
    disabled                     = false
    sync_child_accounts          = false
    network_view_mapping_policy  = "AUTO_CREATE"
}
`, name, gridMember)
}

func testAccAwsrte53taskgroupRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccAwsrte53taskgroupAwsAccountIdsFileToken(awsAccountIdsFileToken string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_aws_account_ids_file_token" {
    aws_account_ids_file_token = %q
}
`, awsAccountIdsFileToken)
}

func testAccAwsrte53taskgroupComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccAwsrte53taskgroupConsolidateZones(consolidateZones string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_consolidate_zones" {
    consolidate_zones = %q
}
`, consolidateZones)
}

func testAccAwsrte53taskgroupConsolidatedView(consolidatedView string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_consolidated_view" {
    consolidated_view = %q
}
`, consolidatedView)
}

func testAccAwsrte53taskgroupDisabled(disabled string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_disabled" {
    disabled = %q
}
`, disabled)
}

func testAccAwsrte53taskgroupGridMember(gridMember string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_grid_member" {
    grid_member = %q
}
`, gridMember)
}

func testAccAwsrte53taskgroupMultipleAccountsSyncPolicy(multipleAccountsSyncPolicy string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_multiple_accounts_sync_policy" {
    multiple_accounts_sync_policy = %q
}
`, multipleAccountsSyncPolicy)
}

func testAccAwsrte53taskgroupName(name string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_name" {
    name = %q
}
`, name)
}

func testAccAwsrte53taskgroupNetworkView(networkView string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_network_view" {
    network_view = %q
}
`, networkView)
}

func testAccAwsrte53taskgroupNetworkViewMappingPolicy(networkViewMappingPolicy string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_network_view_mapping_policy" {
    network_view_mapping_policy = %q
}
`, networkViewMappingPolicy)
}

func testAccAwsrte53taskgroupRoleArn(roleArn string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_role_arn" {
    role_arn = %q
}
`, roleArn)
}

func testAccAwsrte53taskgroupSyncChildAccounts(syncChildAccounts string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_sync_child_accounts" {
    sync_child_accounts = %q
}
`, syncChildAccounts)
}

func testAccAwsrte53taskgroupTaskList(taskList string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_task_list" {
    task_list = %q
}
`, taskList)
}
