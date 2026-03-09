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
					resource.TestCheckResourceAttr(resourceName, "global_allow_list_rpz", ""),
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
				Config: testAccParentalcontrolSubscribersiteGlobalAllowListRpz(name, "4567"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "global_allow_list_rpz", "4567"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteGlobalAllowListRpz(name, "7890"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "global_allow_list_rpz", "7890"),
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
	name1 := acctest.RandomNameWithPrefix("subscriber-site")
	name2 := acctest.RandomNameWithPrefix("subscriber-site")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteName(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteName(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscribersiteResource_NasGateways(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_nas_gateways"
	var v parentalcontrol.ParentalcontrolSubscribersite
	//name := acctest.RandomNameWithPrefix("subscriber-site")
	//nasGateways1 := []map[string]string{{
	//	"ip_address": "12.1.1.1",
	//	"name":       "test-nas-gateway",
	//	"shared_secret":     "secret123",
	//}}
	//nasGateways2 := []map[string]string{{
	//	"ip_address": "12.1.1.1",
	//	"name":       "test-nas-gateway",
	//	"shared_secret":     "secret123",
	//}}
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
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteProxyRpzPassthru(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "proxy_rpz_passthru", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteProxyRpzPassthru(name, "false"),
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

func TestAccParentalcontrolSubscribersiteResource_SubscriberCollectionType(t *testing.T) {
	var resourceName = "nios_parentalcontrol_subscribersite.test_subscriber_collection_type"
	var v parentalcontrol.ParentalcontrolSubscribersite
	name := acctest.RandomNameWithPrefix("subscriber-site")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscribersiteSubscriberCollectionType(name, "API"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscriber_collection_type", "API"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscribersiteSubscriberCollectionType(name, "RADIUS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscribersiteExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscriber_collection_type", "RADIUS"),
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
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test" {
    name = %q
}
`, name)
}

func testAccParentalcontrolSubscribersiteAbss(name, abss string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_abss" {
    name = %q
    abss = %q
}
`, name, abss)
}

func testAccParentalcontrolSubscribersiteApiMembers(name, apiMembers string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_api_members" {
    name = %q
    api_members = %q
}
`, name, apiMembers)
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
    extattrs = %q
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

func testAccParentalcontrolSubscribersiteMembers(name, members string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_members" {
    name = %q
    members = %q
}
`, name, members)
}

func testAccParentalcontrolSubscribersiteMsps(name, msps string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_msps" {
    name = %q
    msps = %q
}
`, name, msps)
}

func testAccParentalcontrolSubscribersiteName(name string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_name" {
    name = %q
}
`, name)
}

func testAccParentalcontrolSubscribersiteNasGateways(name, nasGateways string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_nas_gateways" {
    name = %q
    nas_gateways = %q
}
`, name, nasGateways)
}

func testAccParentalcontrolSubscribersiteNasPort(name, nasPort string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_nas_port" {
    name = %q
    nas_port = %q
}
`, name, nasPort)
}

func testAccParentalcontrolSubscribersiteProxyRpzPassthru(name, proxyRpzPassthru string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_proxy_rpz_passthru" {
    name = %q
    proxy_rpz_passthru = %q
}
`, name, proxyRpzPassthru)
}

func testAccParentalcontrolSubscribersiteSpms(name, spms string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_spms" {
    name = %q
    spms = %q
}
`, name, spms)
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

func testAccParentalcontrolSubscribersiteSubscriberCollectionType(name, subscriberCollectionType string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscribersite" "test_subscriber_collection_type" {
    name = %q
    subscriber_collection_type = %q
}
`, name, subscriberCollectionType)
}
