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
	t.Skip("skipping test as it needs file operation, adding this under TODO")
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
	taskGroupName := acctest.RandomNameWithPrefix("test-taskgroup")
	gridMember := "infoblox.localdomain"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupComment(taskGroupName, gridMember, "COMMENT"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsrte53taskgroupComment(taskGroupName, gridMember, "COMMENT_UPDATE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsrte53taskgroupResource_ConsolidateZones(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_consolidate_zones"
	var v cloud.Awsrte53taskgroup
	taskGroupName := acctest.RandomNameWithPrefix("test-taskgroup")
	gridMember := "infoblox.localdomain"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// // Create and Read
			{
				Config: testAccAwsrte53taskgroupConsolidateZones(taskGroupName, gridMember, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "consolidate_zones", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsrte53taskgroupResource_ConsolidatedView(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_consolidated_view"
	var v cloud.Awsrte53taskgroup
	taskGroupName := acctest.RandomNameWithPrefix("test-taskgroup")
	gridMember := "infoblox.localdomain"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupConsolidatedView(taskGroupName, gridMember, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "consolidated_view", "default"),
				),
			},
		},
	})
}

func TestAccAwsrte53taskgroupResource_Disabled(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_disabled"
	var v cloud.Awsrte53taskgroup
	taskGroupName := acctest.RandomNameWithPrefix("test-taskgroup")
	gridMember := "infoblox.localdomain"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupDisabled(taskGroupName, gridMember, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsrte53taskgroupDisabled(taskGroupName, gridMember, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsrte53taskgroupResource_GridMember(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_grid_member"
	var v cloud.Awsrte53taskgroup
	taskGroupName := acctest.RandomNameWithPrefix("test-taskgroup")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupGridMember("infoblox.localdomain", taskGroupName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_member", "infoblox.localdomain"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsrte53taskgroupGridMember("member.com", taskGroupName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_member", "member.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsrte53taskgroupResource_MultipleAccountsSyncPolicy(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_multiple_accounts_sync_policy"
	var v cloud.Awsrte53taskgroup
	taskGroupName := acctest.RandomNameWithPrefix("test-taskgroup")
	gridMember := "infoblox.localdomain"
	roleArn := "arn:aws:iam:::role/Role-name"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupMultipleAccountsSyncPolicy(taskGroupName, gridMember, "DISCOVER_CHILDREN", roleArn),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "multiple_accounts_sync_policy", "DISCOVER_CHILDREN"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsrte53taskgroupResource_Name(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_name"
	var v cloud.Awsrte53taskgroup
	taskGroupName := acctest.RandomNameWithPrefix("test-taskgroup")
	gridMember := "infoblox.localdomain"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupName(taskGroupName, gridMember),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", taskGroupName),
				),
			},
			// Update and Read
			{
				Config: testAccAwsrte53taskgroupName(taskGroupName+"-updated", gridMember),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", taskGroupName+"-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsrte53taskgroupResource_NetworkView(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_network_view"
	var v cloud.Awsrte53taskgroup
	taskGroupName := acctest.RandomNameWithPrefix("test-taskgroup")
	gridMember := "infoblox.localdomain"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupNetworkView(taskGroupName, gridMember, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "default"),
				),
			},
		},
	})
}

func TestAccAwsrte53taskgroupResource_NetworkViewMappingPolicy(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_network_view_mapping_policy"
	var v cloud.Awsrte53taskgroup
	taskGroupName := acctest.RandomNameWithPrefix("test-taskgroup")
	gridMember := "infoblox.localdomain"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupNetworkViewMappingPolicy(taskGroupName, gridMember, "AUTO_CREATE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view_mapping_policy", "AUTO_CREATE"),
				),
			},
		},
	})
}

func TestAccAwsrte53taskgroupResource_RoleArn(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_role_arn"
	var v cloud.Awsrte53taskgroup
	taskGroupName := acctest.RandomNameWithPrefix("test-taskgroup")
	gridMember := "infoblox.localdomain"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupRoleArn(taskGroupName, gridMember, "arn:aws:iam:::role/Role-name"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "role_arn", "arn:aws:iam:::role/Role-name"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsrte53taskgroupRoleArn(taskGroupName, gridMember, "arn:aws:iam::1:role/Role-name"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "role_arn", "arn:aws:iam::1:role/Role-name"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsrte53taskgroupResource_SyncChildAccounts(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_sync_child_accounts"
	var v cloud.Awsrte53taskgroup
	taskGroupName := acctest.RandomNameWithPrefix("test-taskgroup")
	gridMember := "infoblox.localdomain"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupSyncChildAccounts(taskGroupName, gridMember, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sync_child_accounts", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsrte53taskgroupSyncChildAccounts(taskGroupName, gridMember, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sync_child_accounts", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsrte53taskgroupResource_TaskList(t *testing.T) {
	var resourceName = "nios_cloud_awsrte53taskgroup.test_task_list"
	var v cloud.Awsrte53taskgroup
	taskGroupName := acctest.RandomNameWithPrefix("test-taskgroup")
	gridMember := "infoblox.localdomain"
	task_list := []map[string]any{
		{
			"aws_user":           "awsuser/b25lLmF3c191c2VyJEFLSUFVNUpHWlRURVRTWEwyVEU0:AKIAU5JGZTTETSXL2TE4",
			"credentials_type":   "DIRECT",
			"disabled":           false,
			"filter":             "*",
			"name":               "test1",
			"schedule_interval":  5,
			"schedule_units":     "DAYS",
			"sync_private_zones": true,
			"sync_public_zones":  true,
		},
		{
			"aws_user":           "awsuser/b25lLmF3c191c2VyJEFLSUFVNUpHWlRURVRTWEwyVEU0:AKIAU5JGZTTETSXL2TE4",
			"credentials_type":   "DIRECT",
			"disabled":           false,
			"filter":             "*",
			"name":               "test2",
			"schedule_interval":  10,
			"schedule_units":     "DAYS",
			"sync_private_zones": true,
			"sync_public_zones":  true,
		},
	}
	task_list_update := []map[string]any{
		{
			"aws_user":           "awsuser/b25lLmF3c191c2VyJEFLSUFhYWFhYWFhYWFhYWFhYWFh:AKIAaaaaaaaaaaaaaaaa",
			"credentials_type":   "DIRECT",
			"disabled":           false,
			"filter":             "*",
			"name":               "test1",
			"schedule_interval":  1,
			"schedule_units":     "MINS",
			"sync_private_zones": true,
			"sync_public_zones":  true,
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsrte53taskgroupTaskList(taskGroupName, gridMember, task_list),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "task_list.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "task_list.0.aws_user", "awsuser/b25lLmF3c191c2VyJEFLSUFVNUpHWlRURVRTWEwyVEU0:AKIAU5JGZTTETSXL2TE4"),
					resource.TestCheckResourceAttr(resourceName, "task_list.0.credentials_type", "DIRECT"),
					resource.TestCheckResourceAttr(resourceName, "task_list.0.disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "task_list.0.filter", "*"),
					resource.TestCheckResourceAttr(resourceName, "task_list.0.name", "test1"),
					resource.TestCheckResourceAttr(resourceName, "task_list.0.schedule_interval", "5"),
					resource.TestCheckResourceAttr(resourceName, "task_list.0.schedule_units", "DAYS"),
					resource.TestCheckResourceAttr(resourceName, "task_list.0.sync_private_zones", "true"),
					resource.TestCheckResourceAttr(resourceName, "task_list.0.sync_public_zones", "true"),
					resource.TestCheckResourceAttr(resourceName, "task_list.1.name", "test2"),
					resource.TestCheckResourceAttr(resourceName, "task_list.1.schedule_interval", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsrte53taskgroupTaskList(taskGroupName, gridMember, task_list_update),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "task_list.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "task_list.0.aws_user", "awsuser/b25lLmF3c191c2VyJEFLSUFhYWFhYWFhYWFhYWFhYWFh:AKIAaaaaaaaaaaaaaaaa"),
					resource.TestCheckResourceAttr(resourceName, "task_list.0.credentials_type", "DIRECT"),
					resource.TestCheckResourceAttr(resourceName, "task_list.0.disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "task_list.0.filter", "*"),
					resource.TestCheckResourceAttr(resourceName, "task_list.0.name", "test1"),
					resource.TestCheckResourceAttr(resourceName, "task_list.0.schedule_interval", "1"),
					resource.TestCheckResourceAttr(resourceName, "task_list.0.schedule_units", "MINS"),
					resource.TestCheckResourceAttr(resourceName, "task_list.0.sync_private_zones", "true"),
					resource.TestCheckResourceAttr(resourceName, "task_list.0.sync_public_zones", "true"),
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

func testAccAwsrte53taskgroupAwsAccountIdsFileToken(awsAccountIdsFileToken string) string {
	taskGroupName := acctest.RandomNameWithPrefix("test-taskgroup")
	gridMember := "infoblox.localdomain"

	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_aws_account_ids_file_token" {
    name                        = %q
    grid_member                 = %q
    aws_account_ids_file_token  = %q
    disabled                    = false
    sync_child_accounts         = false
    network_view_mapping_policy = "AUTO_CREATE"
}
`, taskGroupName, gridMember, awsAccountIdsFileToken)
}

func testAccAwsrte53taskgroupComment(taskGroupName, gridMember, comment string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_comment" {
    name    = %q
    grid_member = %q
    comment = %q
    disabled = false
    sync_child_accounts = false
	network_view_mapping_policy = "AUTO_CREATE"
}
`, taskGroupName, gridMember, comment)
}

func testAccAwsrte53taskgroupConsolidateZones(taskGroupName, gridMember, consolidateZones string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_consolidate_zones" {
    name    = %q
    grid_member = %q    
    consolidate_zones = %q
	network_view_mapping_policy = "AUTO_CREATE"
}
`, taskGroupName, gridMember, consolidateZones)
}

func testAccAwsrte53taskgroupConsolidatedView(taskGroupName, gridMember, consolidatedView string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_consolidated_view" {
    name    = %q
    grid_member = %q      
    consolidated_view = %q
	consolidate_zones = false
	network_view_mapping_policy = "DIRECT"
	network_view= "default"
}
`, taskGroupName, gridMember, consolidatedView)
}

func testAccAwsrte53taskgroupDisabled(taskGroupName, gridMember, disabled string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_disabled" {
    name    = %q
    grid_member = %q
    disabled = %q
	network_view_mapping_policy = "AUTO_CREATE"
}
`, taskGroupName, gridMember, disabled)
}

func testAccAwsrte53taskgroupGridMember(gridMember, taskGroupName string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_grid_member" {
	grid_member = %q
	name    = %q 
}
`, gridMember, taskGroupName)
}

func testAccAwsrte53taskgroupMultipleAccountsSyncPolicy(taskGroupName, gridMember, multipleAccountsSyncPolicy, roleArn string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_multiple_accounts_sync_policy" {
    name                        = %q
    grid_member                 = %q
    multiple_accounts_sync_policy = %q
    role_arn                    = %q
	network_view_mapping_policy = "AUTO_CREATE"
}
`, taskGroupName, gridMember, multipleAccountsSyncPolicy, roleArn)
}

func testAccAwsrte53taskgroupName(taskGroupName, gridMember string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_name" {
    name    = %q
    grid_member = %q
    disabled = false
	network_view_mapping_policy = "AUTO_CREATE"
}
`, taskGroupName, gridMember)
}

func testAccAwsrte53taskgroupNetworkView(taskGroupName, gridMember, networkView string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_network_view" {
    name    = %q
    grid_member = %q      
    consolidated_view = "default"
	consolidate_zones = false
	network_view_mapping_policy = "DIRECT"
	network_view= %q
}
`, taskGroupName, gridMember, networkView)
}

func testAccAwsrte53taskgroupNetworkViewMappingPolicy(taskGroupName, gridMember, networkViewMappingPolicy string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_network_view_mapping_policy" {
    name                        = %q
    grid_member                 = %q
    network_view_mapping_policy = %q
}
`, taskGroupName, gridMember, networkViewMappingPolicy)
}

func testAccAwsrte53taskgroupRoleArn(taskGroupName, gridMember, roleArn string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_role_arn" {
    name                        = %q
    grid_member                 = %q	   
	role_arn = %q
	network_view_mapping_policy = "AUTO_CREATE"
}
`, taskGroupName, gridMember, roleArn)
}

func testAccAwsrte53taskgroupSyncChildAccounts(taskGroupName, gridMember, syncChildAccounts string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_sync_child_accounts" {
    name                        = %q
    grid_member                 = %q	    
	sync_child_accounts = %q
	network_view_mapping_policy = "AUTO_CREATE"
	role_arn = "arn:aws:iam::123456789012:role/Role-name"
}
`, taskGroupName, gridMember, syncChildAccounts)
}

func testAccAwsrte53taskgroupTaskList(taskGroupName, gridMember string, taskList []map[string]any) string {
	taskListHCL := utils.ConvertSliceOfMapsToHCL(taskList)
	return fmt.Sprintf(`
resource "nios_cloud_awsrte53taskgroup" "test_task_list" {
    name                        = %q
    grid_member                 = %q    
	task_list = %s
	network_view_mapping_policy = "DIRECT"
	network_view= "default"
	consolidated_view = "default"
	role_arn = "arn:aws:iam::123456789012:role/Role-name"
}
`, taskGroupName, gridMember, taskListHCL)
}
