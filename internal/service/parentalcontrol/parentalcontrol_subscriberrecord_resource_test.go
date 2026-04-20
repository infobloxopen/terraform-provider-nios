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
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordBasicConfig(ipAddr, ipsd, localId, prefix, site, subscriberId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ip_addr", ipAddr),
					resource.TestCheckResourceAttr(resourceName, "ipsd", ipsd),
					resource.TestCheckResourceAttr(resourceName, "localid", localId),
					resource.TestCheckResourceAttr(resourceName, "prefix", prefix),
					resource.TestCheckResourceAttr(resourceName, "site", site),
					resource.TestCheckResourceAttr(resourceName, "subscriber_id", subscriberId),
					// Test fields with default value
					// All the default values are undefined
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_disappears(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	resourceName := "nios_parentalcontrol_subscriberrecord.test"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckParentalcontrolSubscriberrecordDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccParentalcontrolSubscriberrecordBasicConfig(ipAddr, ipsd, localId, prefix, site, subscriberId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					testAccCheckParentalcontrolSubscriberrecordDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_AccountingSessionId(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_accounting_session_id"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	accountingSessionId1 := "Acct-Session-Id=9999732d-34590346"
	accountingSessionId2 := "Acct-Session-Id=9999732d-34590357"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAccountingSessionId(ipAddr, ipsd, localId, prefix, site, subscriberId, accountingSessionId1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "accounting_session_id", accountingSessionId1),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAccountingSessionId(ipAddr, ipsd, localId, prefix, site, subscriberId, accountingSessionId2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "accounting_session_id", accountingSessionId2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_AltIpAddr(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_alt_ip_addr"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	altIp1 := "2123:345:287::6727:22"
	altIp2 := "2123:345:287::6727:221"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAltIpAddr(ipAddr, ipsd, localId, prefix, site, subscriberId, altIp1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "alt_ip_addr", altIp1),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAltIpAddr(ipAddr, ipsd, localId, prefix, site, subscriberId, altIp2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "alt_ip_addr", altIp2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Ans0(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_ans0"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAns0(ipAddr, ipsd, localId, prefix, site, subscriberId, "User-Name=JOHN"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ans0", "User-Name=JOHN"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAns0(ipAddr, ipsd, localId, prefix, site, subscriberId, "User-Name=JOHN_UPDATED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ans0", "User-Name=JOHN_UPDATED"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Ans1(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_ans1"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAns1(ipAddr, ipsd, localId, prefix, site, subscriberId, "IMEI=1234567890"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ans1", "IMEI=1234567890"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAns1(ipAddr, ipsd, localId, prefix, site, subscriberId, "IMEI=1234567890_UPDATED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ans1", "IMEI=1234567890_UPDATED"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Ans2(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_ans2"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAns2(ipAddr, ipsd, localId, prefix, site, subscriberId, "IMSI=12345"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ans2", "IMSI=12345"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAns2(ipAddr, ipsd, localId, prefix, site, subscriberId, "IMSI=12345_UPDATED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ans2", "IMSI=12345_UPDATED"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Ans3(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_ans3"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAns3(ipAddr, ipsd, localId, prefix, site, subscriberId, "LocalId=11011"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ans3", "LocalId=11011"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAns3(ipAddr, ipsd, localId, prefix, site, subscriberId, "LocalId=11010101"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ans3", "LocalId=11010101"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Ans4(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_ans4"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAns4(ipAddr, ipsd, localId, prefix, site, subscriberId, "MSISDN=12345"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ans4", "MSISDN=12345"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordAns4(ipAddr, ipsd, localId, prefix, site, subscriberId, "MSISDN=67890"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ans4", "MSISDN=67890"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_BlackList(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_black_list"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordBlackList(ipAddr, ipsd, localId, prefix, site, subscriberId, "facebook.com,bad.com,verybad.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "black_list", "facebook.com,bad.com,verybad.com"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordBlackList(ipAddr, ipsd, localId, prefix, site, subscriberId, "wapi.com,info.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "black_list", "wapi.com,info.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Bwflag(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_bwflag"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordBwflag(ipAddr, ipsd, localId, prefix, site, subscriberId, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bwflag", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordBwflag(ipAddr, ipsd, localId, prefix, site, subscriberId, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bwflag", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_DynamicCategoryPolicy(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_dynamic_category_policy"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordDynamicCategoryPolicy(ipAddr, ipsd, localId, prefix, site, subscriberId, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dynamic_category_policy", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordDynamicCategoryPolicy(ipAddr, ipsd, localId, prefix, site, subscriberId, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dynamic_category_policy", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Flags(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_flags"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordFlags(ipAddr, ipsd, localId, prefix, site, subscriberId, "SB"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "flags", "SB"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordFlags(ipAddr, ipsd, localId, prefix, site, subscriberId, "B"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "flags", "B"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_IpAddr(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_ip_addr"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr1 := acctest.RandomIP()
	ipAddr2 := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordIpAddr(ipAddr1, ipsd, localId, prefix, site, subscriberId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ip_addr", ipAddr1),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordIpAddr(ipAddr2, ipsd, localId, prefix, site, subscriberId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ip_addr", ipAddr2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Ipsd(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_ipsd"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd1 := "N/A"
	ipsd2 := "N/A"
	localId := "N/A"
	prefix := "32"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordIpsd(ipAddr, ipsd1, localId, prefix, site, subscriberId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipsd", ipsd1),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordIpsd(ipAddr, ipsd2, localId, prefix, site, subscriberId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipsd", ipsd2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Localid(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_localid"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId1 := "N/A"
	localId2 := "N/A"
	prefix := "32"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordLocalid(ipAddr, ipsd, localId1, prefix, site, subscriberId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "localid", localId1),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordLocalid(ipAddr, ipsd, localId2, prefix, site, subscriberId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "localid", localId2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_NasContextual(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_nas_contextual"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	nasContextual1 := "NAS-PORT=1813"
	nasContextual2 := "NAS-PORT=1814"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordNasContextual(ipAddr, ipsd, localId, prefix, site, subscriberId, nasContextual1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nas_contextual", nasContextual1),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordNasContextual(ipAddr, ipsd, localId, prefix, site, subscriberId, nasContextual2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nas_contextual", nasContextual2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_OpCode(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_op_code"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordOpCode(ipAddr, ipsd, localId, prefix, site, subscriberId, "101"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "op_code", "101"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordOpCode(ipAddr, ipsd, localId, prefix, site, subscriberId, "110"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "op_code", "110"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_ParentalControlPolicy(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_parental_control_policy"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	parentalControlPolicy1 := "104"
	parentalControlPolicy2 := "101"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordParentalControlPolicy(ipAddr, ipsd, localId, prefix, site, subscriberId, parentalControlPolicy1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "parental_control_policy", parentalControlPolicy1),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordParentalControlPolicy(ipAddr, ipsd, localId, prefix, site, subscriberId, parentalControlPolicy2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "parental_control_policy", parentalControlPolicy2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Prefix(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_prefix"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix1 := "32"
	prefix2 := "31"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordPrefix(ipAddr, ipsd, localId, prefix1, site, subscriberId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", prefix1),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordPrefix(ipAddr, ipsd, localId, prefix2, site, subscriberId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", prefix2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_ProxyAll(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_proxy_all"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordProxyAll(ipAddr, ipsd, localId, prefix, site, subscriberId, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "proxy_all", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordProxyAll(ipAddr, ipsd, localId, prefix, site, subscriberId, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "proxy_all", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_Site(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_site"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site1 := "site1"
	site2 := "site2"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := int32(32)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordSite(ipAddr, ipsd, localId, prefix, site1, subscriberId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "site", site1),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordSite(ipAddr, ipsd, localId, prefix, site2, subscriberId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "site", site2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_SubscriberId(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_subscriber_id"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId1 := "IMSI=12345"
	subscriberId2 := "IMSI=67890"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordSubscriberId(ipAddr, ipsd, localId, prefix, site, subscriberId1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscriber_id", subscriberId1),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordSubscriberId(ipAddr, ipsd, localId, prefix, site, subscriberId2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscriber_id", subscriberId2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_SubscriberSecurePolicy(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_subscriber_secure_policy"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	subscriberSecurePolicy1 := "FF"
	subscriberSecurePolicy2 := "FE"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordSubscriberSecurePolicy(ipAddr, ipsd, localId, prefix, site, subscriberId, subscriberSecurePolicy1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscriber_secure_policy", subscriberSecurePolicy1),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordSubscriberSecurePolicy(ipAddr, ipsd, localId, prefix, site, subscriberId, subscriberSecurePolicy2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "subscriber_secure_policy", subscriberSecurePolicy2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_UnknownCategoryPolicy(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_unknown_category_policy"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordUnknownCategoryPolicy(ipAddr, ipsd, localId, prefix, site, subscriberId, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "unknown_category_policy", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordUnknownCategoryPolicy(ipAddr, ipsd, localId, prefix, site, subscriberId, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "unknown_category_policy", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_WhiteList(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_white_list"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordWhiteList(ipAddr, ipsd, localId, prefix, site, subscriberId, "info.com,good.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "white_list", "info.com,good.com"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordWhiteList(ipAddr, ipsd, localId, prefix, site, subscriberId, "wapi.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "white_list", "wapi.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccParentalcontrolSubscriberrecordResource_WpcCategoryPolicy(t *testing.T) {
	t.Skip("TODO - TO BE FIXED IN FUTURE RELEASES FOR INTEGRATION TESTS")
	var resourceName = "nios_parentalcontrol_subscriberrecord.test_wpc_category_policy"
	var v parentalcontrol.ParentalcontrolSubscriberrecord
	ipAddr := acctest.RandomIP()
	site := "site1"
	subscriberId := "IMSI=12345"
	ipsd := "N/A"
	localId := "N/A"
	prefix := "32"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccParentalcontrolSubscriberrecordWpcCategoryPolicy(ipAddr, ipsd, localId, prefix, site, subscriberId, "1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wpc_category_policy", "1"),
				),
			},
			// Update and Read
			{
				Config: testAccParentalcontrolSubscriberrecordWpcCategoryPolicy(ipAddr, ipsd, localId, prefix, site, subscriberId, "2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckParentalcontrolSubscriberrecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "wpc_category_policy", "2"),
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

func testAccParentalcontrolSubscriberrecordBasicConfig(ipAddr, ipsd, localId, prefix, site, subscriberId string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test" {
	ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId)
}

func testAccParentalcontrolSubscriberrecordAccountingSessionId(ipAddr, ipsd, localId, prefix, site, subscriberId, accountingSessionId string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_accounting_session_id" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
	accounting_session_id = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId, accountingSessionId)
}

func testAccParentalcontrolSubscriberrecordAltIpAddr(ipAddr, ipsd, localId, prefix, site, subscriberId, altIpAddr string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_alt_ip_addr" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
	alt_ip_addr = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId, altIpAddr)
}

func testAccParentalcontrolSubscriberrecordAns0(ipAddr, ipsd, localId, prefix, site, subscriberId, ans0 string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_ans0" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
	ans0 = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId, ans0)
}

func testAccParentalcontrolSubscriberrecordAns1(ipAddr, ipsd, localId, prefix, site, subscriberId, ans1 string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_ans1" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
	ans1 = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId, ans1)
}

func testAccParentalcontrolSubscriberrecordAns2(ipAddr, ipsd, localId, prefix, site, subscriberId, ans2 string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_ans2" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
	ans2 = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId, ans2)
}

func testAccParentalcontrolSubscriberrecordAns3(ipAddr, ipsd, localId, prefix, site, subscriberId, ans3 string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_ans3" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
	ans3 = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId, ans3)
}

func testAccParentalcontrolSubscriberrecordAns4(ipAddr, ipsd, localId, prefix, site, subscriberId, ans4 string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_ans4" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
	ans4 = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId, ans4)
}

func testAccParentalcontrolSubscriberrecordBlackList(ipAddr, ipsd, localId, prefix, site, subscriberId, blackList string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_black_list" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
	black_list = %q
	bwflag = true
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId, blackList)
}

func testAccParentalcontrolSubscriberrecordBwflag(ipAddr, ipsd, localId, prefix, site, subscriberId, bwflag string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_bwflag" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
	bwflag = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId, bwflag)
}

func testAccParentalcontrolSubscriberrecordDynamicCategoryPolicy(ipAddr, ipsd, localId, prefix, site, subscriberId, dynamicCategoryPolicy string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_dynamic_category_policy" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
	dynamic_category_policy = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId, dynamicCategoryPolicy)
}

func testAccParentalcontrolSubscriberrecordFlags(ipAddr, ipsd, localId, prefix, site, subscriberId, flags string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_flags" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
	flags = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId, flags)
}

func testAccParentalcontrolSubscriberrecordIpAddr(ipAddr, ipsd, localId, prefix, site, subscriberId string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_ip_addr" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId)
}

func testAccParentalcontrolSubscriberrecordIpsd(ipAddr, ipsd, localId, prefix, site, subscriberId string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_ipsd" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId)
}

func testAccParentalcontrolSubscriberrecordLocalid(ipAddr, ipsd, localId, prefix, site, subscriberId string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_localid" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId)
}

func testAccParentalcontrolSubscriberrecordNasContextual(ipAddr, ipsd, localId, prefix, site, subscriberId, nasContextual string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_nas_contextual" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
	nas_contextual = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId, nasContextual)
}

func testAccParentalcontrolSubscriberrecordOpCode(ipAddr, ipsd, localId, prefix, site, subscriberId, opCode string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_op_code" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
	op_code = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId, opCode)
}

func testAccParentalcontrolSubscriberrecordParentalControlPolicy(ipAddr, ipsd, localId, prefix, site, subscriberId, parentalControlPolicy string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_parental_control_policy" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
	parental_control_policy = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId, parentalControlPolicy)
}

func testAccParentalcontrolSubscriberrecordPrefix(ipAddr, ipsd, localId, prefix, site, subscriberId string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_prefix" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId)
}

func testAccParentalcontrolSubscriberrecordProxyAll(ipAddr, ipsd, localId, prefix, site, subscriberId, proxyAll string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_proxy_all" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
	proxy_all = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId, proxyAll)
}

func testAccParentalcontrolSubscriberrecordSite(ipAddr, ipsd, localId string, prefix int32, site, subscriberId string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_site" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %d
	site = %q
	subscriber_id = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId)
}

func testAccParentalcontrolSubscriberrecordSubscriberId(ipAddr, ipsd, localId, prefix, site, subscriberId string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_subscriber_id" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId)
}

func testAccParentalcontrolSubscriberrecordSubscriberSecurePolicy(ipAddr, ipsd, localId, prefix, site, subscriberId, subscriberSecurePolicy string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_subscriber_secure_policy" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
	subscriber_secure_policy = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId, subscriberSecurePolicy)
}

func testAccParentalcontrolSubscriberrecordUnknownCategoryPolicy(ipAddr, ipsd, localId, prefix, site, subscriberId, unknownCategoryPolicy string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_unknown_category_policy" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
	unknown_category_policy = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId, unknownCategoryPolicy)
}

func testAccParentalcontrolSubscriberrecordWhiteList(ipAddr, ipsd, localId, prefix, site, subscriberId, whiteList string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_white_list" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
	white_list = %q
	bwflag = true
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId, whiteList)
}

func testAccParentalcontrolSubscriberrecordWpcCategoryPolicy(ipAddr, ipsd, localId, prefix, site, subscriberId, wpcCategoryPolicy string) string {
	return fmt.Sprintf(`
resource "nios_parentalcontrol_subscriberrecord" "test_wpc_category_policy" {
    ip_addr = %q
	ipsd = %q
	localid = %q
	prefix = %q
	site = %q
	subscriber_id = %q
	wpc_category_policy = %q
}
`, ipAddr, ipsd, localId, prefix, site, subscriberId, wpcCategoryPolicy)
}
