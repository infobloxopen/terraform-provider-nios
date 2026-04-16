package parentalcontrol_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/parentalcontrol"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForParentalcontrolSubscribersite = "abss,api_members,api_port,block_size,blocking_ipv4_vip1,blocking_ipv4_vip2,blocking_ipv6_vip1,blocking_ipv6_vip2,comment,dca_sub_bw_list,dca_sub_query_count,enable_global_allow_list_rpz,enable_rpz_filtering_bypass,extattrs,first_port,global_allow_list_rpz,maximum_subscribers,members,msps,name,nas_gateways,nas_port,proxy_rpz_passthru,spms,stop_anycast,strict_nat,subscriber_collection_type"

// TODO : Required parents for the execution of tests - members (infoblox.localdomain, infoblox.member1)

func TestAccParentalcontrolSubscribersiteResource_basic(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "block_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "blocking_ipv4_vip1", ""),
					resource.TestCheckResourceAttr(resourceName, "blocking_ipv4_vip2", ""),
					resource.TestCheckResourceAttr(resourceName, "blocking_ipv6_vip1", ""),
					resource.TestCheckResourceAttr(resourceName, "blocking_ipv6_vip2", ""),
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
					resource.TestCheckResourceAttr(resourceName, "dca_sub_bw_list", "false"),
					resource.TestCheckResourceAttr(resourceName, "dca_sub_query_count", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_global_allow_list_rpz", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_rpz_filtering_bypass", "false"),
					resource.TestCheckResourceAttr(resourceName, "first_port", "1024"),
					resource.TestCheckResourceAttr(resourceName, "maximum_subscribers", "1000000"),
					resource.TestCheckResourceAttr(resourceName, "nas_port", "1813"),
					resource.TestCheckResourceAttr(resourceName, "proxy_rpz_passthru", "false"),
					resource.TestCheckResourceAttr(resourceName, "stop_anycast", "true"),
					resource.TestCheckResourceAttr(resourceName, "strict_nat", "true"),
					resource.TestCheckResourceAttr(resourceName, "subscriber_collection_type", "RADIUS"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_disappears(t *testing.T) {
	resourceName := "nios_parentalcontrol_subscribersite.test"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckParentalcontrolSubscribersiteDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccParentalcontrolSubscribersiteBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					testAccCheckParentalcontrolSubscribersiteDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_Abss(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_abss"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	blockingPolicy1 := acctest.RandomNameWithPrefix("blocking-policy-")
	blockingPolicy2 := acctest.RandomNameWithPrefix("blocking-policy-")
	value1 := acctest.Random32Hexadecimal()
	value2 := acctest.Random32Hexadecimal()
	abss1 := []map[string]any{
		{
			"ip_address":      "12.12.1.1",
			"blocking_policy": blockingPolicy1,
		},
	}
	abss2 := []map[string]any{
		{
			"ip_address":      "12.12.10.10",
			"blocking_policy": blockingPolicy2,
		},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteAbss(name, abss1, blockingPolicy1, value1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "abss.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "abss.0.ip_address", "12.12.1.1"),
					resource.TestCheckResourceAttr(resourceName, "abss.0.blocking_policy", blockingPolicy1),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteAbss(name, abss2, blockingPolicy2, value2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "abss.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "abss.0.ip_address", "12.12.10.10"),
					resource.TestCheckResourceAttr(resourceName, "abss.0.blocking_policy", blockingPolicy2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_ApiMembers(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_api_members"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	memberName := utils.GetNIOSGridMasterHostName()
	memberUpdateName := utils.GetNIOSGridMemberHostName()
	apiMembers1 := []map[string]any{
		{"name": memberName},
	}
	apiMembers2 := []map[string]any{
		{"name": memberUpdateName},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteApiMembers(name, apiMembers1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "api_members.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "api_members.0.name", memberName),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteApiMembers(name, apiMembers2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "api_members.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "api_members.0.name", memberUpdateName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_BlockSize(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_block_size"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteBlockSize(name, "200"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "block_size", "200"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteBlockSize(name, "100"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "block_size", "100"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_BlockingIpv4Vip1(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_blocking_ipv4_vip1"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteBlockingIpv4Vip1(name, "12.2.1.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blocking_ipv4_vip1", "12.2.1.1"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteBlockingIpv4Vip1(name, "12.2.1.2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blocking_ipv4_vip1", "12.2.1.2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_BlockingIpv4Vip2(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_blocking_ipv4_vip2"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteBlockingIpv4Vip2(name, "20.20.1.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blocking_ipv4_vip2", "20.20.1.1"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteBlockingIpv4Vip2(name, "20.20.1.2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blocking_ipv4_vip2", "20.20.1.2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_BlockingIpv6Vip1(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_blocking_ipv6_vip1"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteBlockingIpv6Vip1(name, "2002:1f93::12:1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blocking_ipv6_vip1", "2002:1f93::12:1"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteBlockingIpv6Vip1(name, "2002:1f93::12:2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blocking_ipv6_vip1", "2002:1f93::12:2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_BlockingIpv6Vip2(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_blocking_ipv6_vip2"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteBlockingIpv6Vip2(name, "2002:1f93::13:1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blocking_ipv6_vip2", "2002:1f93::13:1"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteBlockingIpv6Vip2(name, "2002:1f93::13:2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blocking_ipv6_vip2", "2002:1f93::13:2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_Comment(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_comment"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteComment(name, "Example Parental Control Subscriber Site Comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Example Parental Control Subscriber Site Comment"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteComment(name, "Example Parental Control Subscriber Site Comment Updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Example Parental Control Subscriber Site Comment Updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_DcaSubBwList(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_dca_sub_bw_list"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteDcaSubBwList(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dca_sub_bw_list", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteDcaSubBwList(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dca_sub_bw_list", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_DcaSubQueryCount(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_dca_sub_query_count"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteDcaSubQueryCount(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dca_sub_query_count", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteDcaSubQueryCount(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dca_sub_query_count", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_EnableGlobalAllowListRpz(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_enable_global_allow_list_rpz"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteEnableGlobalAllowListRpz(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_global_allow_list_rpz", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteEnableGlobalAllowListRpz(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_global_allow_list_rpz", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_EnableRpzFilteringBypass(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_enable_rpz_filtering_bypass"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteEnableRpzFilteringBypass(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_rpz_filtering_bypass", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteEnableRpzFilteringBypass(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_rpz_filtering_bypass", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_extattrs"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteExtAttrs(name, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteExtAttrs(name, map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_FirstPort(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_first_port"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteFirstPort(name, "1234"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "first_port", "1234"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteFirstPort(name, "2345"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "first_port", "2345"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_GlobalAllowListRpz(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_global_allow_list_rpz"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteGlobalAllowListRpz(name, "34"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "global_allow_list_rpz", "34"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteGlobalAllowListRpz(name, "45"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "global_allow_list_rpz", "45"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_MaximumSubscribers(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_maximum_subscribers"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteMaximumSubscribers(name, "20000"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "maximum_subscribers", "20000"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteMaximumSubscribers(name, "30000"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "maximum_subscribers", "30000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_Members(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_members"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	memberName := utils.GetNIOSGridMasterHostName()
	memberUpdateName := utils.GetNIOSGridMemberHostName()
	members1 := []map[string]any{
		{"name": memberName},
	}
	members2 := []map[string]any{
		{"name": memberUpdateName},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteMembers(name, members1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "members.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "members.0.name", memberName),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteMembers(name, members2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "members.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "members.0.name", memberUpdateName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_Msps(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_msps"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	msps1 := []map[string]any{
		{"ip_address": "12.13.15.15"},
	}
	msps2 := []map[string]any{
		{"ip_address": "12.13.15.16"},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteMsps(name, msps1, "12.13.155.15"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "msps.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "msps.0.ip_address", "12.13.15.15"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteMsps(name, msps2, "12.13.155.15"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "msps.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "msps.0.ip_address", "12.13.15.16"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_NasGateways(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_nas_gateways"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	nasGateways1 := []map[string]any{{
		"ip_address":    "12.1.1.1",
		"name":          "test-nas-gateway11",
		"shared_secret": "secret123",
	}}
	nasGateways2 := []map[string]any{{
		"ip_address":    "12.12.11.11",
		"name":          "test-nas-gateway12",
		"shared_secret": "secret123",
	}}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteNasGateways(name, nasGateways1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nas_gateways.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "nas_gateways.0.ip_address", "12.1.1.1"),
					resource.TestCheckResourceAttr(resourceName, "nas_gateways.0.name", "test-nas-gateway11"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteNasGateways(name, nasGateways2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nas_gateways.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "nas_gateways.0.ip_address", "12.12.11.11"),
					resource.TestCheckResourceAttr(resourceName, "nas_gateways.0.name", "test-nas-gateway12"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_NasPort(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_nas_port"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteNasPort(name, "2000"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nas_port", "2000"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteNasPort(name, "2100"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nas_port", "2100"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_ProxyRpzPassthru(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_proxy_rpz_passthru"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	msps := []map[string]any{
		{"ip_address": "12.13.15.15"},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteProxyRpzPassthru(name, "true", "12.13.14.18", msps),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "proxy_rpz_passthru", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteProxyRpzPassthru(name, "false", "12.13.14.19", msps),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "proxy_rpz_passthru", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_Spms(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_spms"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	spms1 := []map[string]any{
		{"ip_address": "12.13.14.15"},
	}
	spms2 := []map[string]any{
		{"ip_address": "12.13.14.16"},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteSpms(name, spms1, "12.12.122.12"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "spms.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spms.0.ip_address", "12.13.14.15"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteSpms(name, spms2, "12.12.122.12"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "spms.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spms.0.ip_address", "12.13.14.16"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_StopAnycast(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_stop_anycast"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteStopAnycast(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "stop_anycast", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteStopAnycast(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "stop_anycast", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_StrictNat(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_strict_nat"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteStrictNat(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "strict_nat", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteStrictNat(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "strict_nat", "false"),
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

func testAccParentalcontrolSubscribersiteBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test" {
    name = %q
}
`, name)
}

func testAccParentalcontrolSubscribersiteAbss(name string, abss []map[string]any, blockingPolicy, value string) string {
	abssStr := utils.ConvertSliceOfMapsToHCL(abss)
	config := fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_abss" {
    name = %q
    abss = %s
	depends_on = [nios_parentalcontrol_blockingpolicy.test_blocking_policy]
}
`, name, abssStr)
	return strings.Join([]string{testAccParentBlockingPolicy(blockingPolicy, value), config}, "")
}

func testAccParentalcontrolSubscribersiteApiMembers(name string, apiMembers []map[string]any) string {
	apiMembersStr := utils.ConvertSliceOfMapsToHCL(apiMembers)
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_api_members" {
    name = %q
	members = %s
    api_members = %s
}
`, name, apiMembersStr, apiMembersStr)
}

func testAccParentalcontrolSubscribersiteBlockSize(name, blockSize string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_block_size" {
    name = %q
    block_size = %q
}
`, name, blockSize)
}

func testAccParentalcontrolSubscribersiteBlockingIpv4Vip1(name, blockingIpv4Vip1 string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_blocking_ipv4_vip1" {
    name = %q
    blocking_ipv4_vip1 = %q
}
`, name, blockingIpv4Vip1)
}

func testAccParentalcontrolSubscribersiteBlockingIpv4Vip2(name, blockingIpv4Vip2 string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_blocking_ipv4_vip2" {
    name = %q
    blocking_ipv4_vip2 = %q
}
`, name, blockingIpv4Vip2)
}

func testAccParentalcontrolSubscribersiteBlockingIpv6Vip1(name, blockingIpv6Vip1 string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_blocking_ipv6_vip1" {
    name = %q
    blocking_ipv6_vip1 = %q
}
`, name, blockingIpv6Vip1)
}

func testAccParentalcontrolSubscribersiteBlockingIpv6Vip2(name, blockingIpv6Vip2 string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_blocking_ipv6_vip2" {
    name = %q
    blocking_ipv6_vip2 = %q
}
`, name, blockingIpv6Vip2)
}

func testAccParentalcontrolSubscribersiteComment(name, comment string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccParentalcontrolSubscribersiteDcaSubBwList(name, dcaSubBwList string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_dca_sub_bw_list" {
    name = %q
    dca_sub_bw_list = %q
}
`, name, dcaSubBwList)
}

func testAccParentalcontrolSubscribersiteDcaSubQueryCount(name, dcaSubQueryCount string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_dca_sub_query_count" {
    name = %q
    dca_sub_query_count = %q
}
`, name, dcaSubQueryCount)
}

func testAccParentalcontrolSubscribersiteEnableGlobalAllowListRpz(name, enableGlobalAllowListRpz string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_enable_global_allow_list_rpz" {
    name = %q
    enable_global_allow_list_rpz = %q
}
`, name, enableGlobalAllowListRpz)
}

func testAccParentalcontrolSubscribersiteEnableRpzFilteringBypass(name, enableRpzFilteringBypass string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_enable_rpz_filtering_bypass" {
    name = %q
    enable_rpz_filtering_bypass = %q
}
`, name, enableRpzFilteringBypass)
}

func testAccParentalcontrolSubscribersiteExtAttrs(name string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_extattrs" {
    name = %q
    extattrs = %s
}
`, name, extattrsStr)
}

func testAccParentalcontrolSubscribersiteFirstPort(name, firstPort string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_first_port" {
    name = %q
    first_port = %q
}
`, name, firstPort)
}

func testAccParentalcontrolSubscribersiteGlobalAllowListRpz(name, globalAllowListRpz string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_global_allow_list_rpz" {
    name = %q
    global_allow_list_rpz = %q
}
`, name, globalAllowListRpz)
}

func testAccParentalcontrolSubscribersiteMaximumSubscribers(name, maximumSubscribers string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_maximum_subscribers" {
    name = %q
    maximum_subscribers = %q
}
`, name, maximumSubscribers)
}

func testAccParentalcontrolSubscribersiteMembers(name string, members []map[string]any) string {
	membersStr := utils.ConvertSliceOfMapsToHCL(members)
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_members" {
    name = %q
    members = %s
}
`, name, membersStr)
}

func testAccParentalcontrolSubscribersiteMsps(name string, msps []map[string]any, blockingIpv4Vip2 string) string {
	mspsStr := utils.ConvertSliceOfMapsToHCL(msps)
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_msps" {
    name = %q
    msps = %s
    blocking_ipv4_vip2 = %q
}
`, name, mspsStr, blockingIpv4Vip2)
}

func testAccParentalcontrolSubscribersiteNasGateways(name string, nasGateways []map[string]any) string {
	nasGatewaysStr := utils.ConvertSliceOfMapsToHCL(nasGateways)
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_nas_gateways" {
    name = %q
    nas_gateways = %s
}
`, name, nasGatewaysStr)
}

func testAccParentalcontrolSubscribersiteNasPort(name, nasPort string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_nas_port" {
    name = %q
    nas_port = %q
}
`, name, nasPort)
}

func testAccParentalcontrolSubscribersiteProxyRpzPassthru(name, proxyRpzPassthru, blockingIpv4Vip2 string, msps []map[string]any) string {
	mspsStr := utils.ConvertSliceOfMapsToHCL(msps)
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_proxy_rpz_passthru" {
    name = %q
    proxy_rpz_passthru = %q
    blocking_ipv4_vip2 = %q
	msps = %s
}
`, name, proxyRpzPassthru, blockingIpv4Vip2, mspsStr)
}

func testAccParentalcontrolSubscribersiteSpms(name string, spms []map[string]any, blockingIpv4Vip2 string) string {
	spmsStr := utils.ConvertSliceOfMapsToHCL(spms)
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_spms" {
    name = %q
    spms = %s
    blocking_ipv4_vip2 = %q
}
`, name, spmsStr, blockingIpv4Vip2)
}

func testAccParentalcontrolSubscribersiteStopAnycast(name, stopAnycast string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_stop_anycast" {
    name = %q
    stop_anycast = %q
}
`, name, stopAnycast)
}

func testAccParentalcontrolSubscribersiteStrictNat(name, strictNat string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_strict_nat" {
    name = %q
    strict_nat = %q
}
`, name, strictNat)
}

func testAccParentBlockingPolicy(name, value string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_blockingpolicy" "test_blocking_policy" {
	name = %q
	value = %q
}
`, name, value)
}
