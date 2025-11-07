package parentalcontrol_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/parentalcontrol"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForParentalcontrolSubscriberrecord = "accounting_session_id,alt_ip_addr,ans0,ans1,ans2,ans3,ans4,black_list,bwflag,dynamic_category_policy,flags,ip_addr,ipsd,localid,nas_contextual,op_code,parental_control_policy,prefix,proxy_all,site,subscriber_id,subscriber_secure_policy,unknown_category_policy,white_list,wpc_category_policy"

func TestAccParentalcontrolSubscriberrecordResource_basic(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_disappears(t *testing.T) {
	resourceName := "nios_parentalcontrol_subscriberrecord.test"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckParentalcontrolSubscriberrecordDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccParentalcontrolSubscriberrecordBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					testAccCheckParentalcontrolSubscriberrecordDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Ref(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_ref"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_AccountingSessionId(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_accounting_session_id"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAccountingSessionId("ACCOUNTING_SESSION_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "accounting_session_id", "ACCOUNTING_SESSION_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAccountingSessionId("ACCOUNTING_SESSION_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "accounting_session_id", "ACCOUNTING_SESSION_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_AltIpAddr(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_alt_ip_addr"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAltIpAddr("ALT_IP_ADDR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "alt_ip_addr", "ALT_IP_ADDR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAltIpAddr("ALT_IP_ADDR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "alt_ip_addr", "ALT_IP_ADDR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Ans0(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_ans0"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAns0("ANS0_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ans0", "ANS0_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAns0("ANS0_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ans0", "ANS0_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Ans1(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_ans1"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAns1("ANS1_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ans1", "ANS1_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAns1("ANS1_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ans1", "ANS1_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Ans2(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_ans2"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAns2("ANS2_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ans2", "ANS2_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAns2("ANS2_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ans2", "ANS2_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Ans3(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_ans3"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAns3("ANS3_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ans3", "ANS3_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAns3("ANS3_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ans3", "ANS3_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Ans4(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_ans4"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAns4("ANS4_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ans4", "ANS4_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAns4("ANS4_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ans4", "ANS4_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_BlackList(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_black_list"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordBlackList("BLACK_LIST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "black_list", "BLACK_LIST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordBlackList("BLACK_LIST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "black_list", "BLACK_LIST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Bwflag(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_bwflag"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordBwflag("BWFLAG_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bwflag", "BWFLAG_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordBwflag("BWFLAG_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bwflag", "BWFLAG_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_DynamicCategoryPolicy(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_dynamic_category_policy"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordDynamicCategoryPolicy("DYNAMIC_CATEGORY_POLICY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dynamic_category_policy", "DYNAMIC_CATEGORY_POLICY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordDynamicCategoryPolicy("DYNAMIC_CATEGORY_POLICY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dynamic_category_policy", "DYNAMIC_CATEGORY_POLICY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Flags(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_flags"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordFlags("FLAGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "flags", "FLAGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordFlags("FLAGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "flags", "FLAGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_IpAddr(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_ip_addr"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordIpAddr("IP_ADDR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ip_addr", "IP_ADDR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordIpAddr("IP_ADDR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ip_addr", "IP_ADDR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Ipsd(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_ipsd"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordIpsd("IPSD_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipsd", "IPSD_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordIpsd("IPSD_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipsd", "IPSD_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Localid(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_localid"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordLocalid("LOCALID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "localid", "LOCALID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordLocalid("LOCALID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "localid", "LOCALID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_NasContextual(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_nas_contextual"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordNasContextual("NAS_CONTEXTUAL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nas_contextual", "NAS_CONTEXTUAL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordNasContextual("NAS_CONTEXTUAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nas_contextual", "NAS_CONTEXTUAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_OpCode(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_op_code"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordOpCode("OP_CODE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "op_code", "OP_CODE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordOpCode("OP_CODE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "op_code", "OP_CODE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_ParentalControlPolicy(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_parental_control_policy"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordParentalControlPolicy("PARENTAL_CONTROL_POLICY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "parental_control_policy", "PARENTAL_CONTROL_POLICY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordParentalControlPolicy("PARENTAL_CONTROL_POLICY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "parental_control_policy", "PARENTAL_CONTROL_POLICY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Prefix(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_prefix"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordPrefix("PREFIX_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", "PREFIX_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordPrefix("PREFIX_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", "PREFIX_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_ProxyAll(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_proxy_all"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordProxyAll("PROXY_ALL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "proxy_all", "PROXY_ALL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordProxyAll("PROXY_ALL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "proxy_all", "PROXY_ALL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Site(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_site"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordSite("SITE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "site", "SITE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordSite("SITE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "site", "SITE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_SubscriberId(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_subscriber_id"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordSubscriberId("SUBSCRIBER_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscriber_id", "SUBSCRIBER_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordSubscriberId("SUBSCRIBER_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscriber_id", "SUBSCRIBER_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_SubscriberSecurePolicy(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_subscriber_secure_policy"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordSubscriberSecurePolicy("SUBSCRIBER_SECURE_POLICY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscriber_secure_policy", "SUBSCRIBER_SECURE_POLICY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordSubscriberSecurePolicy("SUBSCRIBER_SECURE_POLICY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscriber_secure_policy", "SUBSCRIBER_SECURE_POLICY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_UnknownCategoryPolicy(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_unknown_category_policy"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordUnknownCategoryPolicy("UNKNOWN_CATEGORY_POLICY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "unknown_category_policy", "UNKNOWN_CATEGORY_POLICY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordUnknownCategoryPolicy("UNKNOWN_CATEGORY_POLICY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "unknown_category_policy", "UNKNOWN_CATEGORY_POLICY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_WhiteList(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_white_list"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordWhiteList("WHITE_LIST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "white_list", "WHITE_LIST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordWhiteList("WHITE_LIST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "white_list", "WHITE_LIST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_WpcCategoryPolicy(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_wpc_category_policy"
	var v parentalcontrol.ParentalcontrolSubscriberrecord

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordWpcCategoryPolicy("WPC_CATEGORY_POLICY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wpc_category_policy", "WPC_CATEGORY_POLICY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordWpcCategoryPolicy("WPC_CATEGORY_POLICY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wpc_category_policy", "WPC_CATEGORY_POLICY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckParentalcontrolSubscriberrecordExists(ctx context.Context, resourceName string, v *parentalcontrol.ParentalcontrolSubscriberrecord) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.ParentalControlAPI.
			ParentalcontrolSubscriberrecordAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForParentalcontrolSubscriberrecord).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetParentalcontrolSubscriberrecordResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetParentalcontrolSubscriberrecordResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckParentalcontrolSubscriberrecordDestroy(ctx context.Context, v *parentalcontrol.ParentalcontrolSubscriberrecord) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.ParentalControlAPI.
			ParentalcontrolSubscriberrecordAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForParentalcontrolSubscriberrecord).
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

func testAccCheckParentalcontrolSubscriberrecordDisappears(ctx context.Context, v *parentalcontrol.ParentalcontrolSubscriberrecord) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.ParentalControlAPI.
			ParentalcontrolSubscriberrecordAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccParentalcontrolSubscriberrecordBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test" {
}
`)
}

func testAccParentalcontrolSubscriberrecordRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccParentalcontrolSubscriberrecordAccountingSessionId(accountingSessionId string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_accounting_session_id" {
    accounting_session_id = %q
}
`, accountingSessionId)
}

func testAccParentalcontrolSubscriberrecordAltIpAddr(altIpAddr string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_alt_ip_addr" {
    alt_ip_addr = %q
}
`, altIpAddr)
}

func testAccParentalcontrolSubscriberrecordAns0(ans0 string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_ans0" {
    ans0 = %q
}
`, ans0)
}

func testAccParentalcontrolSubscriberrecordAns1(ans1 string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_ans1" {
    ans1 = %q
}
`, ans1)
}

func testAccParentalcontrolSubscriberrecordAns2(ans2 string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_ans2" {
    ans2 = %q
}
`, ans2)
}

func testAccParentalcontrolSubscriberrecordAns3(ans3 string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_ans3" {
    ans3 = %q
}
`, ans3)
}

func testAccParentalcontrolSubscriberrecordAns4(ans4 string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_ans4" {
    ans4 = %q
}
`, ans4)
}

func testAccParentalcontrolSubscriberrecordBlackList(blackList string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_black_list" {
    black_list = %q
}
`, blackList)
}

func testAccParentalcontrolSubscriberrecordBwflag(bwflag string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_bwflag" {
    bwflag = %q
}
`, bwflag)
}

func testAccParentalcontrolSubscriberrecordDynamicCategoryPolicy(dynamicCategoryPolicy string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_dynamic_category_policy" {
    dynamic_category_policy = %q
}
`, dynamicCategoryPolicy)
}

func testAccParentalcontrolSubscriberrecordFlags(flags string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_flags" {
    flags = %q
}
`, flags)
}

func testAccParentalcontrolSubscriberrecordIpAddr(ipAddr string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_ip_addr" {
    ip_addr = %q
}
`, ipAddr)
}

func testAccParentalcontrolSubscriberrecordIpsd(ipsd string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_ipsd" {
    ipsd = %q
}
`, ipsd)
}

func testAccParentalcontrolSubscriberrecordLocalid(localid string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_localid" {
    localid = %q
}
`, localid)
}

func testAccParentalcontrolSubscriberrecordNasContextual(nasContextual string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_nas_contextual" {
    nas_contextual = %q
}
`, nasContextual)
}

func testAccParentalcontrolSubscriberrecordOpCode(opCode string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_op_code" {
    op_code = %q
}
`, opCode)
}

func testAccParentalcontrolSubscriberrecordParentalControlPolicy(parentalControlPolicy string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_parental_control_policy" {
    parental_control_policy = %q
}
`, parentalControlPolicy)
}

func testAccParentalcontrolSubscriberrecordPrefix(prefix string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_prefix" {
    prefix = %q
}
`, prefix)
}

func testAccParentalcontrolSubscriberrecordProxyAll(proxyAll string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_proxy_all" {
    proxy_all = %q
}
`, proxyAll)
}

func testAccParentalcontrolSubscriberrecordSite(site string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_site" {
    site = %q
}
`, site)
}

func testAccParentalcontrolSubscriberrecordSubscriberId(subscriberId string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_subscriber_id" {
    subscriber_id = %q
}
`, subscriberId)
}

func testAccParentalcontrolSubscriberrecordSubscriberSecurePolicy(subscriberSecurePolicy string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_subscriber_secure_policy" {
    subscriber_secure_policy = %q
}
`, subscriberSecurePolicy)
}

func testAccParentalcontrolSubscriberrecordUnknownCategoryPolicy(unknownCategoryPolicy string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_unknown_category_policy" {
    unknown_category_policy = %q
}
`, unknownCategoryPolicy)
}

func testAccParentalcontrolSubscriberrecordWhiteList(whiteList string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_white_list" {
    white_list = %q
}
`, whiteList)
}

func testAccParentalcontrolSubscriberrecordWpcCategoryPolicy(wpcCategoryPolicy string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_wpc_category_policy" {
    wpc_category_policy = %q
}
`, wpcCategoryPolicy)
}
