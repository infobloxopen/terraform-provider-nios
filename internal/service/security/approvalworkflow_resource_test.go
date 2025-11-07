package security_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForApprovalworkflow = "approval_group,approval_notify_to,approved_notify_to,approver_comment,enable_approval_notify,enable_approved_notify,enable_failed_notify,enable_notify_group,enable_notify_user,enable_rejected_notify,enable_rescheduled_notify,enable_succeeded_notify,extattrs,failed_notify_to,rejected_notify_to,rescheduled_notify_to,submitter_comment,submitter_group,succeeded_notify_to,ticket_number"

func TestAccApprovalworkflowResource_basic(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowBasicConfig("REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApprovalworkflowResource_disappears(t *testing.T) {
	resourceName := "nios_security_approvalworkflow.test"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckApprovalworkflowDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccApprovalworkflowBasicConfig("REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					testAccCheckApprovalworkflowDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccApprovalworkflowResource_Ref(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test_ref"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccApprovalworkflowRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApprovalworkflowResource_ApprovalGroup(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test_approval_group"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowApprovalGroup("APPROVAL_GROUP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "approval_group", "APPROVAL_GROUP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccApprovalworkflowApprovalGroup("APPROVAL_GROUP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "approval_group", "APPROVAL_GROUP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApprovalworkflowResource_ApprovalNotifyTo(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test_approval_notify_to"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowApprovalNotifyTo("APPROVAL_NOTIFY_TO_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "approval_notify_to", "APPROVAL_NOTIFY_TO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccApprovalworkflowApprovalNotifyTo("APPROVAL_NOTIFY_TO_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "approval_notify_to", "APPROVAL_NOTIFY_TO_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApprovalworkflowResource_ApprovedNotifyTo(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test_approved_notify_to"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowApprovedNotifyTo("APPROVED_NOTIFY_TO_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "approved_notify_to", "APPROVED_NOTIFY_TO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccApprovalworkflowApprovedNotifyTo("APPROVED_NOTIFY_TO_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "approved_notify_to", "APPROVED_NOTIFY_TO_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApprovalworkflowResource_ApproverComment(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test_approver_comment"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowApproverComment("APPROVER_COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "approver_comment", "APPROVER_COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccApprovalworkflowApproverComment("APPROVER_COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "approver_comment", "APPROVER_COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApprovalworkflowResource_EnableApprovalNotify(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test_enable_approval_notify"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowEnableApprovalNotify("ENABLE_APPROVAL_NOTIFY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_approval_notify", "ENABLE_APPROVAL_NOTIFY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccApprovalworkflowEnableApprovalNotify("ENABLE_APPROVAL_NOTIFY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_approval_notify", "ENABLE_APPROVAL_NOTIFY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApprovalworkflowResource_EnableApprovedNotify(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test_enable_approved_notify"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowEnableApprovedNotify("ENABLE_APPROVED_NOTIFY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_approved_notify", "ENABLE_APPROVED_NOTIFY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccApprovalworkflowEnableApprovedNotify("ENABLE_APPROVED_NOTIFY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_approved_notify", "ENABLE_APPROVED_NOTIFY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApprovalworkflowResource_EnableFailedNotify(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test_enable_failed_notify"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowEnableFailedNotify("ENABLE_FAILED_NOTIFY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_failed_notify", "ENABLE_FAILED_NOTIFY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccApprovalworkflowEnableFailedNotify("ENABLE_FAILED_NOTIFY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_failed_notify", "ENABLE_FAILED_NOTIFY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApprovalworkflowResource_EnableNotifyGroup(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test_enable_notify_group"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowEnableNotifyGroup("ENABLE_NOTIFY_GROUP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_notify_group", "ENABLE_NOTIFY_GROUP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccApprovalworkflowEnableNotifyGroup("ENABLE_NOTIFY_GROUP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_notify_group", "ENABLE_NOTIFY_GROUP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApprovalworkflowResource_EnableNotifyUser(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test_enable_notify_user"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowEnableNotifyUser("ENABLE_NOTIFY_USER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_notify_user", "ENABLE_NOTIFY_USER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccApprovalworkflowEnableNotifyUser("ENABLE_NOTIFY_USER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_notify_user", "ENABLE_NOTIFY_USER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApprovalworkflowResource_EnableRejectedNotify(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test_enable_rejected_notify"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowEnableRejectedNotify("ENABLE_REJECTED_NOTIFY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_rejected_notify", "ENABLE_REJECTED_NOTIFY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccApprovalworkflowEnableRejectedNotify("ENABLE_REJECTED_NOTIFY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_rejected_notify", "ENABLE_REJECTED_NOTIFY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApprovalworkflowResource_EnableRescheduledNotify(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test_enable_rescheduled_notify"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowEnableRescheduledNotify("ENABLE_RESCHEDULED_NOTIFY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_rescheduled_notify", "ENABLE_RESCHEDULED_NOTIFY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccApprovalworkflowEnableRescheduledNotify("ENABLE_RESCHEDULED_NOTIFY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_rescheduled_notify", "ENABLE_RESCHEDULED_NOTIFY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApprovalworkflowResource_EnableSucceededNotify(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test_enable_succeeded_notify"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowEnableSucceededNotify("ENABLE_SUCCEEDED_NOTIFY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_succeeded_notify", "ENABLE_SUCCEEDED_NOTIFY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccApprovalworkflowEnableSucceededNotify("ENABLE_SUCCEEDED_NOTIFY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_succeeded_notify", "ENABLE_SUCCEEDED_NOTIFY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApprovalworkflowResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test_extattrs"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccApprovalworkflowExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApprovalworkflowResource_FailedNotifyTo(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test_failed_notify_to"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowFailedNotifyTo("FAILED_NOTIFY_TO_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "failed_notify_to", "FAILED_NOTIFY_TO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccApprovalworkflowFailedNotifyTo("FAILED_NOTIFY_TO_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "failed_notify_to", "FAILED_NOTIFY_TO_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApprovalworkflowResource_RejectedNotifyTo(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test_rejected_notify_to"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowRejectedNotifyTo("REJECTED_NOTIFY_TO_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rejected_notify_to", "REJECTED_NOTIFY_TO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccApprovalworkflowRejectedNotifyTo("REJECTED_NOTIFY_TO_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rejected_notify_to", "REJECTED_NOTIFY_TO_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApprovalworkflowResource_RescheduledNotifyTo(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test_rescheduled_notify_to"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowRescheduledNotifyTo("RESCHEDULED_NOTIFY_TO_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rescheduled_notify_to", "RESCHEDULED_NOTIFY_TO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccApprovalworkflowRescheduledNotifyTo("RESCHEDULED_NOTIFY_TO_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rescheduled_notify_to", "RESCHEDULED_NOTIFY_TO_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApprovalworkflowResource_SubmitterComment(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test_submitter_comment"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowSubmitterComment("SUBMITTER_COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "submitter_comment", "SUBMITTER_COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccApprovalworkflowSubmitterComment("SUBMITTER_COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "submitter_comment", "SUBMITTER_COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApprovalworkflowResource_SubmitterGroup(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test_submitter_group"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowSubmitterGroup("SUBMITTER_GROUP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "submitter_group", "SUBMITTER_GROUP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccApprovalworkflowSubmitterGroup("SUBMITTER_GROUP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "submitter_group", "SUBMITTER_GROUP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApprovalworkflowResource_SucceededNotifyTo(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test_succeeded_notify_to"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowSucceededNotifyTo("SUCCEEDED_NOTIFY_TO_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "succeeded_notify_to", "SUCCEEDED_NOTIFY_TO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccApprovalworkflowSucceededNotifyTo("SUCCEEDED_NOTIFY_TO_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "succeeded_notify_to", "SUCCEEDED_NOTIFY_TO_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApprovalworkflowResource_TicketNumber(t *testing.T) {
	var resourceName = "nios_security_approvalworkflow.test_ticket_number"
	var v security.Approvalworkflow

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApprovalworkflowTicketNumber("TICKET_NUMBER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ticket_number", "TICKET_NUMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccApprovalworkflowTicketNumber("TICKET_NUMBER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApprovalworkflowExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ticket_number", "TICKET_NUMBER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckApprovalworkflowExists(ctx context.Context, resourceName string, v *security.Approvalworkflow) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.SecurityAPI.
			ApprovalworkflowAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForApprovalworkflow).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetApprovalworkflowResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetApprovalworkflowResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckApprovalworkflowDestroy(ctx context.Context, v *security.Approvalworkflow) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.SecurityAPI.
			ApprovalworkflowAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForApprovalworkflow).
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

func testAccCheckApprovalworkflowDisappears(ctx context.Context, v *security.Approvalworkflow) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.SecurityAPI.
			ApprovalworkflowAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccApprovalworkflowBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test" {
}
`)
}

func testAccApprovalworkflowRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccApprovalworkflowApprovalGroup(approvalGroup string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test_approval_group" {
    approval_group = %q
}
`, approvalGroup)
}

func testAccApprovalworkflowApprovalNotifyTo(approvalNotifyTo string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test_approval_notify_to" {
    approval_notify_to = %q
}
`, approvalNotifyTo)
}

func testAccApprovalworkflowApprovedNotifyTo(approvedNotifyTo string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test_approved_notify_to" {
    approved_notify_to = %q
}
`, approvedNotifyTo)
}

func testAccApprovalworkflowApproverComment(approverComment string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test_approver_comment" {
    approver_comment = %q
}
`, approverComment)
}

func testAccApprovalworkflowEnableApprovalNotify(enableApprovalNotify string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test_enable_approval_notify" {
    enable_approval_notify = %q
}
`, enableApprovalNotify)
}

func testAccApprovalworkflowEnableApprovedNotify(enableApprovedNotify string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test_enable_approved_notify" {
    enable_approved_notify = %q
}
`, enableApprovedNotify)
}

func testAccApprovalworkflowEnableFailedNotify(enableFailedNotify string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test_enable_failed_notify" {
    enable_failed_notify = %q
}
`, enableFailedNotify)
}

func testAccApprovalworkflowEnableNotifyGroup(enableNotifyGroup string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test_enable_notify_group" {
    enable_notify_group = %q
}
`, enableNotifyGroup)
}

func testAccApprovalworkflowEnableNotifyUser(enableNotifyUser string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test_enable_notify_user" {
    enable_notify_user = %q
}
`, enableNotifyUser)
}

func testAccApprovalworkflowEnableRejectedNotify(enableRejectedNotify string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test_enable_rejected_notify" {
    enable_rejected_notify = %q
}
`, enableRejectedNotify)
}

func testAccApprovalworkflowEnableRescheduledNotify(enableRescheduledNotify string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test_enable_rescheduled_notify" {
    enable_rescheduled_notify = %q
}
`, enableRescheduledNotify)
}

func testAccApprovalworkflowEnableSucceededNotify(enableSucceededNotify string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test_enable_succeeded_notify" {
    enable_succeeded_notify = %q
}
`, enableSucceededNotify)
}

func testAccApprovalworkflowExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccApprovalworkflowFailedNotifyTo(failedNotifyTo string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test_failed_notify_to" {
    failed_notify_to = %q
}
`, failedNotifyTo)
}

func testAccApprovalworkflowRejectedNotifyTo(rejectedNotifyTo string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test_rejected_notify_to" {
    rejected_notify_to = %q
}
`, rejectedNotifyTo)
}

func testAccApprovalworkflowRescheduledNotifyTo(rescheduledNotifyTo string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test_rescheduled_notify_to" {
    rescheduled_notify_to = %q
}
`, rescheduledNotifyTo)
}

func testAccApprovalworkflowSubmitterComment(submitterComment string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test_submitter_comment" {
    submitter_comment = %q
}
`, submitterComment)
}

func testAccApprovalworkflowSubmitterGroup(submitterGroup string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test_submitter_group" {
    submitter_group = %q
}
`, submitterGroup)
}

func testAccApprovalworkflowSucceededNotifyTo(succeededNotifyTo string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test_succeeded_notify_to" {
    succeeded_notify_to = %q
}
`, succeededNotifyTo)
}

func testAccApprovalworkflowTicketNumber(ticketNumber string) string {
	return fmt.Sprintf(`
resource "nios_security_approvalworkflow" "test_ticket_number" {
    ticket_number = %q
}
`, ticketNumber)
}
