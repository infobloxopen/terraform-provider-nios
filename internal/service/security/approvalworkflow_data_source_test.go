
package security_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccApprovalworkflowDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_security_approvalworkflow.test"
	resourceName := "nios_security_approvalworkflow.test"
	var v security.Approvalworkflow

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckApprovalworkflowDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccApprovalworkflowDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
						}, testAccCheckApprovalworkflowResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccApprovalworkflowDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_security_approvalworkflow.test"
	resourceName := "nios_security_approvalworkflow.test"
	var v security.Approvalworkflow
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckApprovalworkflowDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccApprovalworkflowDataSourceConfigExtAttrFilters( acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
						}, testAccCheckApprovalworkflowResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckApprovalworkflowResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
    return []resource.TestCheckFunc{
        resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
        resource.TestCheckResourceAttrPair(resourceName, "approval_group", dataSourceName, "result.0.approval_group"),
        resource.TestCheckResourceAttrPair(resourceName, "approval_notify_to", dataSourceName, "result.0.approval_notify_to"),
        resource.TestCheckResourceAttrPair(resourceName, "approved_notify_to", dataSourceName, "result.0.approved_notify_to"),
        resource.TestCheckResourceAttrPair(resourceName, "approver_comment", dataSourceName, "result.0.approver_comment"),
        resource.TestCheckResourceAttrPair(resourceName, "enable_approval_notify", dataSourceName, "result.0.enable_approval_notify"),
        resource.TestCheckResourceAttrPair(resourceName, "enable_approved_notify", dataSourceName, "result.0.enable_approved_notify"),
        resource.TestCheckResourceAttrPair(resourceName, "enable_failed_notify", dataSourceName, "result.0.enable_failed_notify"),
        resource.TestCheckResourceAttrPair(resourceName, "enable_notify_group", dataSourceName, "result.0.enable_notify_group"),
        resource.TestCheckResourceAttrPair(resourceName, "enable_notify_user", dataSourceName, "result.0.enable_notify_user"),
        resource.TestCheckResourceAttrPair(resourceName, "enable_rejected_notify", dataSourceName, "result.0.enable_rejected_notify"),
        resource.TestCheckResourceAttrPair(resourceName, "enable_rescheduled_notify", dataSourceName, "result.0.enable_rescheduled_notify"),
        resource.TestCheckResourceAttrPair(resourceName, "enable_succeeded_notify", dataSourceName, "result.0.enable_succeeded_notify"),
        resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
        resource.TestCheckResourceAttrPair(resourceName, "failed_notify_to", dataSourceName, "result.0.failed_notify_to"),
        resource.TestCheckResourceAttrPair(resourceName, "rejected_notify_to", dataSourceName, "result.0.rejected_notify_to"),
        resource.TestCheckResourceAttrPair(resourceName, "rescheduled_notify_to", dataSourceName, "result.0.rescheduled_notify_to"),
        resource.TestCheckResourceAttrPair(resourceName, "submitter_comment", dataSourceName, "result.0.submitter_comment"),
        resource.TestCheckResourceAttrPair(resourceName, "submitter_group", dataSourceName, "result.0.submitter_group"),
        resource.TestCheckResourceAttrPair(resourceName, "succeeded_notify_to", dataSourceName, "result.0.succeeded_notify_to"),
        resource.TestCheckResourceAttrPair(resourceName, "ticket_number", dataSourceName, "result.0.ticket_number"),
    }
}

func testAccApprovalworkflowDataSourceConfigFilters() string {
	return `
resource "nios_security_approvalworkflow" "test" {
}

data "nios_security_approvalworkflow" "test" {
  filters = {
	 = nios_security_approvalworkflow.test.
  }
}
`
}

func testAccApprovalworkflowDataSourceConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test" {
  extattrs = {
    Site = %q
  } 
}

data "nios_security_approvalworkflow" "test" {
  extattrfilters = {
	Site = nios_security_approvalworkflow.test.extattrs.Site
  }
}
`,extAttrsValue)
}

