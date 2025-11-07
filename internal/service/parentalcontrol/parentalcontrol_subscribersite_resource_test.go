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

var readableAttributesForParentalcontrolSubscribersite = "abss,api_members,api_port,block_size,blocking_ipv4_vip1,blocking_ipv4_vip2,blocking_ipv6_vip1,blocking_ipv6_vip2,comment,dca_sub_bw_list,dca_sub_query_count,enable_global_allow_list_rpz,enable_rpz_filtering_bypass,extattrs,first_port,global_allow_list_rpz,maximum_subscribers,members,msps,name,nas_gateways,nas_port,proxy_rpz_passthru,spms,stop_anycast,strict_nat,subscriber_collection_type"

func TestAccParentalcontrolSubscribersiteResource_basic(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_disappears(t *testing.T) {
	resourceName := "nios_parentalcontrol_subscribersite.test"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckParentalcontrolSubscribersiteDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccParentalcontrolSubscribersiteBasicConfig("replace_me"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					testAccCheckParentalcontrolSubscribersiteDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_Ref(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_ref"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_Abss(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_abss"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteAbss("ABSS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "abss", "ABSS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteAbss("ABSS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "abss", "ABSS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_ApiMembers(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_api_members"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteApiMembers("API_MEMBERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "api_members", "API_MEMBERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteApiMembers("API_MEMBERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "api_members", "API_MEMBERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_BlockSize(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_block_size"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteBlockSize("BLOCK_SIZE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "block_size", "BLOCK_SIZE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteBlockSize("BLOCK_SIZE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "block_size", "BLOCK_SIZE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_BlockingIpv4Vip1(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_blocking_ipv4_vip1"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteBlockingIpv4Vip1("BLOCKING_IPV4_VIP1_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blocking_ipv4_vip1", "BLOCKING_IPV4_VIP1_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteBlockingIpv4Vip1("BLOCKING_IPV4_VIP1_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blocking_ipv4_vip1", "BLOCKING_IPV4_VIP1_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_BlockingIpv4Vip2(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_blocking_ipv4_vip2"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteBlockingIpv4Vip2("BLOCKING_IPV4_VIP2_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blocking_ipv4_vip2", "BLOCKING_IPV4_VIP2_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteBlockingIpv4Vip2("BLOCKING_IPV4_VIP2_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blocking_ipv4_vip2", "BLOCKING_IPV4_VIP2_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_BlockingIpv6Vip1(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_blocking_ipv6_vip1"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteBlockingIpv6Vip1("BLOCKING_IPV6_VIP1_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blocking_ipv6_vip1", "BLOCKING_IPV6_VIP1_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteBlockingIpv6Vip1("BLOCKING_IPV6_VIP1_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blocking_ipv6_vip1", "BLOCKING_IPV6_VIP1_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_BlockingIpv6Vip2(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_blocking_ipv6_vip2"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteBlockingIpv6Vip2("BLOCKING_IPV6_VIP2_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blocking_ipv6_vip2", "BLOCKING_IPV6_VIP2_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteBlockingIpv6Vip2("BLOCKING_IPV6_VIP2_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blocking_ipv6_vip2", "BLOCKING_IPV6_VIP2_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_Comment(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_comment"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_DcaSubBwList(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_dca_sub_bw_list"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteDcaSubBwList("DCA_SUB_BW_LIST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dca_sub_bw_list", "DCA_SUB_BW_LIST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteDcaSubBwList("DCA_SUB_BW_LIST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dca_sub_bw_list", "DCA_SUB_BW_LIST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_DcaSubQueryCount(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_dca_sub_query_count"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteDcaSubQueryCount("DCA_SUB_QUERY_COUNT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dca_sub_query_count", "DCA_SUB_QUERY_COUNT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteDcaSubQueryCount("DCA_SUB_QUERY_COUNT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dca_sub_query_count", "DCA_SUB_QUERY_COUNT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_EnableGlobalAllowListRpz(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_enable_global_allow_list_rpz"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteEnableGlobalAllowListRpz("ENABLE_GLOBAL_ALLOW_LIST_RPZ_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_global_allow_list_rpz", "ENABLE_GLOBAL_ALLOW_LIST_RPZ_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteEnableGlobalAllowListRpz("ENABLE_GLOBAL_ALLOW_LIST_RPZ_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_global_allow_list_rpz", "ENABLE_GLOBAL_ALLOW_LIST_RPZ_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_EnableRpzFilteringBypass(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_enable_rpz_filtering_bypass"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteEnableRpzFilteringBypass("ENABLE_RPZ_FILTERING_BYPASS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_rpz_filtering_bypass", "ENABLE_RPZ_FILTERING_BYPASS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteEnableRpzFilteringBypass("ENABLE_RPZ_FILTERING_BYPASS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_rpz_filtering_bypass", "ENABLE_RPZ_FILTERING_BYPASS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_extattrs"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_FirstPort(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_first_port"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteFirstPort("FIRST_PORT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "first_port", "FIRST_PORT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteFirstPort("FIRST_PORT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "first_port", "FIRST_PORT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_GlobalAllowListRpz(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_global_allow_list_rpz"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteGlobalAllowListRpz("GLOBAL_ALLOW_LIST_RPZ_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "global_allow_list_rpz", "GLOBAL_ALLOW_LIST_RPZ_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteGlobalAllowListRpz("GLOBAL_ALLOW_LIST_RPZ_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "global_allow_list_rpz", "GLOBAL_ALLOW_LIST_RPZ_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_MaximumSubscribers(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_maximum_subscribers"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteMaximumSubscribers("MAXIMUM_SUBSCRIBERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "maximum_subscribers", "MAXIMUM_SUBSCRIBERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteMaximumSubscribers("MAXIMUM_SUBSCRIBERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "maximum_subscribers", "MAXIMUM_SUBSCRIBERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_Members(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_members"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteMembers("MEMBERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "members", "MEMBERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteMembers("MEMBERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "members", "MEMBERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_Msps(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_msps"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteMsps("MSPS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "msps", "MSPS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteMsps("MSPS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "msps", "MSPS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_Name(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_name"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_NasGateways(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_nas_gateways"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteNasGateways("NAS_GATEWAYS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nas_gateways", "NAS_GATEWAYS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteNasGateways("NAS_GATEWAYS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nas_gateways", "NAS_GATEWAYS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_NasPort(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_nas_port"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteNasPort("NAS_PORT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nas_port", "NAS_PORT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteNasPort("NAS_PORT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nas_port", "NAS_PORT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_ProxyRpzPassthru(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_proxy_rpz_passthru"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteProxyRpzPassthru("PROXY_RPZ_PASSTHRU_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "proxy_rpz_passthru", "PROXY_RPZ_PASSTHRU_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteProxyRpzPassthru("PROXY_RPZ_PASSTHRU_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "proxy_rpz_passthru", "PROXY_RPZ_PASSTHRU_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_Spms(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_spms"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteSpms("SPMS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "spms", "SPMS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteSpms("SPMS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "spms", "SPMS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_StopAnycast(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_stop_anycast"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteStopAnycast("STOP_ANYCAST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "stop_anycast", "STOP_ANYCAST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteStopAnycast("STOP_ANYCAST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "stop_anycast", "STOP_ANYCAST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_StrictNat(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_strict_nat"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteStrictNat("STRICT_NAT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "strict_nat", "STRICT_NAT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteStrictNat("STRICT_NAT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "strict_nat", "STRICT_NAT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_SubscriberCollectionType(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_subscriber_collection_type"
	var v parentalcontrol.ParentalcontrolSubscribersite

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteSubscriberCollectionType("SUBSCRIBER_COLLECTION_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscriber_collection_type", "SUBSCRIBER_COLLECTION_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteSubscriberCollectionType("SUBSCRIBER_COLLECTION_TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscriber_collection_type", "SUBSCRIBER_COLLECTION_TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckParentalcontrolSubscribersiteExists(ctx context.Context, resourceName string, v *parentalcontrol.ParentalcontrolSubscribersite) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.ParentalControlAPI.
			ParentalcontrolSubscribersiteAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForParentalcontrolSubscribersite).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetParentalcontrolSubscribersiteResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetParentalcontrolSubscribersiteResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckParentalcontrolSubscribersiteDestroy(ctx context.Context, v *parentalcontrol.ParentalcontrolSubscribersite) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.ParentalControlAPI.
			ParentalcontrolSubscribersiteAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForParentalcontrolSubscribersite).
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

func testAccCheckParentalcontrolSubscribersiteDisappears(ctx context.Context, v *parentalcontrol.ParentalcontrolSubscribersite) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.ParentalControlAPI.
			ParentalcontrolSubscribersiteAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccParentalcontrolSubscribersiteBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test" {
}
`)
}

func testAccParentalcontrolSubscribersiteRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccParentalcontrolSubscribersiteAbss(abss string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_abss" {
    abss = %q
}
`, abss)
}

func testAccParentalcontrolSubscribersiteApiMembers(apiMembers string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_api_members" {
    api_members = %q
}
`, apiMembers)
}

func testAccParentalcontrolSubscribersiteBlockSize(blockSize string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_block_size" {
    block_size = %q
}
`, blockSize)
}

func testAccParentalcontrolSubscribersiteBlockingIpv4Vip1(blockingIpv4Vip1 string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_blocking_ipv4_vip1" {
    blocking_ipv4_vip1 = %q
}
`, blockingIpv4Vip1)
}

func testAccParentalcontrolSubscribersiteBlockingIpv4Vip2(blockingIpv4Vip2 string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_blocking_ipv4_vip2" {
    blocking_ipv4_vip2 = %q
}
`, blockingIpv4Vip2)
}

func testAccParentalcontrolSubscribersiteBlockingIpv6Vip1(blockingIpv6Vip1 string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_blocking_ipv6_vip1" {
    blocking_ipv6_vip1 = %q
}
`, blockingIpv6Vip1)
}

func testAccParentalcontrolSubscribersiteBlockingIpv6Vip2(blockingIpv6Vip2 string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_blocking_ipv6_vip2" {
    blocking_ipv6_vip2 = %q
}
`, blockingIpv6Vip2)
}

func testAccParentalcontrolSubscribersiteComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccParentalcontrolSubscribersiteDcaSubBwList(dcaSubBwList string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_dca_sub_bw_list" {
    dca_sub_bw_list = %q
}
`, dcaSubBwList)
}

func testAccParentalcontrolSubscribersiteDcaSubQueryCount(dcaSubQueryCount string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_dca_sub_query_count" {
    dca_sub_query_count = %q
}
`, dcaSubQueryCount)
}

func testAccParentalcontrolSubscribersiteEnableGlobalAllowListRpz(enableGlobalAllowListRpz string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_enable_global_allow_list_rpz" {
    enable_global_allow_list_rpz = %q
}
`, enableGlobalAllowListRpz)
}

func testAccParentalcontrolSubscribersiteEnableRpzFilteringBypass(enableRpzFilteringBypass string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_enable_rpz_filtering_bypass" {
    enable_rpz_filtering_bypass = %q
}
`, enableRpzFilteringBypass)
}

func testAccParentalcontrolSubscribersiteExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccParentalcontrolSubscribersiteFirstPort(firstPort string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_first_port" {
    first_port = %q
}
`, firstPort)
}

func testAccParentalcontrolSubscribersiteGlobalAllowListRpz(globalAllowListRpz string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_global_allow_list_rpz" {
    global_allow_list_rpz = %q
}
`, globalAllowListRpz)
}

func testAccParentalcontrolSubscribersiteMaximumSubscribers(maximumSubscribers string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_maximum_subscribers" {
    maximum_subscribers = %q
}
`, maximumSubscribers)
}

func testAccParentalcontrolSubscribersiteMembers(members string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_members" {
    members = %q
}
`, members)
}

func testAccParentalcontrolSubscribersiteMsps(msps string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_msps" {
    msps = %q
}
`, msps)
}

func testAccParentalcontrolSubscribersiteName(name string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_name" {
    name = %q
}
`, name)
}

func testAccParentalcontrolSubscribersiteNasGateways(nasGateways string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_nas_gateways" {
    nas_gateways = %q
}
`, nasGateways)
}

func testAccParentalcontrolSubscribersiteNasPort(nasPort string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_nas_port" {
    nas_port = %q
}
`, nasPort)
}

func testAccParentalcontrolSubscribersiteProxyRpzPassthru(proxyRpzPassthru string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_proxy_rpz_passthru" {
    proxy_rpz_passthru = %q
}
`, proxyRpzPassthru)
}

func testAccParentalcontrolSubscribersiteSpms(spms string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_spms" {
    spms = %q
}
`, spms)
}

func testAccParentalcontrolSubscribersiteStopAnycast(stopAnycast string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_stop_anycast" {
    stop_anycast = %q
}
`, stopAnycast)
}

func testAccParentalcontrolSubscribersiteStrictNat(strictNat string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_strict_nat" {
    strict_nat = %q
}
`, strictNat)
}

func testAccParentalcontrolSubscribersiteSubscriberCollectionType(subscriberCollectionType string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_subscriber_collection_type" {
    subscriber_collection_type = %q
}
`, subscriberCollectionType)
}
