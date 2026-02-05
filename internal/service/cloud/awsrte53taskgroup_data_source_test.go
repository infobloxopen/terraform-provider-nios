package cloud_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/cloud"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccAwsrte53taskgroupDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_cloud_aws_route53_task_group.test"
	resourceName := "nios_cloud_aws_route53_task_group.test"
	var v cloud.Awsrte53taskgroup
	taskGroupName := acctest.RandomNameWithPrefix("test-taskgroup")
	gridMember := "infoblox.localdomain"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAwsrte53taskgroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAwsrte53taskgroupDataSourceConfigFilters(taskGroupName, gridMember),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckAwsrte53taskgroupExists(context.Background(), resourceName, &v),
					}, testAccCheckAwsrte53taskgroupResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckAwsrte53taskgroupResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "account_id", dataSourceName, "result.0.account_id"),
		resource.TestCheckResourceAttrPair(resourceName, "accounts_list", dataSourceName, "result.0.accounts_list"),
		resource.TestCheckResourceAttrPair(resourceName, "aws_account_ids_file_token", dataSourceName, "result.0.aws_account_ids_file_token"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "consolidate_zones", dataSourceName, "result.0.consolidate_zones"),
		resource.TestCheckResourceAttrPair(resourceName, "consolidated_view", dataSourceName, "result.0.consolidated_view"),
		resource.TestCheckResourceAttrPair(resourceName, "disabled", dataSourceName, "result.0.disabled"),
		resource.TestCheckResourceAttrPair(resourceName, "grid_member", dataSourceName, "result.0.grid_member"),
		resource.TestCheckResourceAttrPair(resourceName, "multiple_accounts_sync_policy", dataSourceName, "result.0.multiple_accounts_sync_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "network_view", dataSourceName, "result.0.network_view"),
		resource.TestCheckResourceAttrPair(resourceName, "network_view_mapping_policy", dataSourceName, "result.0.network_view_mapping_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "role_arn", dataSourceName, "result.0.role_arn"),
		resource.TestCheckResourceAttrPair(resourceName, "sync_child_accounts", dataSourceName, "result.0.sync_child_accounts"),
		resource.TestCheckResourceAttrPair(resourceName, "sync_status", dataSourceName, "result.0.sync_status"),
		resource.TestCheckResourceAttrPair(resourceName, "task_list", dataSourceName, "result.0.task_list"),
	}
}

func testAccAwsrte53taskgroupDataSourceConfigFilters(taskGroupName, gridMember string) string {
	return fmt.Sprintf(`
resource "nios_cloud_aws_route53_task_group" "test" {
    name                         = %q
    grid_member                  = %q
    disabled                     = false
    sync_child_accounts          = false
    network_view_mapping_policy  = "AUTO_CREATE"
}

data "nios_cloud_aws_route53_task_group" "test" {
  filters = {
    name = nios_cloud_aws_route53_task_group.test.name
    grid_member = nios_cloud_aws_route53_task_group.test.grid_member
  }
}
`, taskGroupName, gridMember)
}
