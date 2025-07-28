package dhcp_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForRangetemplate = "bootfile,bootserver,cloud_api_compatible,comment,ddns_domainname,ddns_generate_hostname,delegated_member,deny_all_clients,deny_bootp,email_list,enable_ddns,enable_dhcp_thresholds,enable_email_warnings,enable_pxe_lease_time,enable_snmp_warnings,exclude,extattrs,failover_association,fingerprint_filter_rules,high_water_mark,high_water_mark_reset,ignore_dhcp_option_list_request,known_clients,lease_scavenge_time,logic_filter_rules,low_water_mark,low_water_mark_reset,mac_filter_rules,member,ms_options,ms_server,nac_filter_rules,name,nextserver,number_of_addresses,offset,option_filter_rules,options,pxe_lease_time,recycle_leases,relay_agent_filter_rules,server_association_type,unknown_clients,update_dns_on_lease_renewal,use_bootfile,use_bootserver,use_ddns_domainname,use_ddns_generate_hostname,use_deny_bootp,use_email_list,use_enable_ddns,use_enable_dhcp_thresholds,use_ignore_dhcp_option_list_request,use_known_clients,use_lease_scavenge_time,use_logic_filter_rules,use_ms_options,use_nextserver,use_options,use_pxe_lease_time,use_recycle_leases,use_unknown_clients,use_update_dns_on_lease_renewal"

func TestAccRangetemplateResource_basic(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateBasicConfig(name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "number_of_addresses", fmt.Sprintf("%d", numberOfAdresses)),
					resource.TestCheckResourceAttr(resourceName, "offset", fmt.Sprintf("%d", offset)),

					resource.TestCheckResourceAttr(resourceName, "cloud_api_compatible", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "deny_all_clients", "false"),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_dhcp_thresholds", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_email_warnings", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_snmp_warnings", "false"),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "95"),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "85"),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "false"),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "-1"),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "0"),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "10"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "43200"),
					resource.TestCheckResourceAttr(resourceName, "options.0.num", "51"),
					resource.TestCheckResourceAttr(resourceName, "options.0.use_option", "false"),
					resource.TestCheckResourceAttr(resourceName, "options.0.vendor_class", "DHCP"),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "true"),
					resource.TestCheckResourceAttr(resourceName, "server_association_type", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "false"),

					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_email_list", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_known_clients", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ms_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_unknown_clients", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_disappears(t *testing.T) {
	resourceName := "nios_dhcp_rangetemplate.test"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRangetemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRangetemplateBasicConfig(name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					testAccCheckRangetemplateDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRangetemplateResource_Bootfile(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_bootfile"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	bootFile := "bootfile.txt"
	useBootFile := true

	bootFileUpdated := "bootfile12.txt"
	useBootFileUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateBootfile(useBootFile, bootFile, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "bootfile.txt"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateBootfile(useBootFileUpdated, bootFileUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "bootfile12.txt"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_Bootserver(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_bootserver"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	bootServer := "bootServer"
	useBootServer := true

	bootServerUpdated := "bootServer3"
	useBootServerUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateBootserver(useBootServer, bootServer, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", bootServer),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateBootserver(useBootServerUpdated, bootServerUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", bootServerUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_CloudApiCompatible(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_cloud_api_compatible"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	useCloudApiCompatible := true
	useCloudApiCompatibleUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateCloudApiCompatible(useCloudApiCompatible, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_api_compatible", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateCloudApiCompatible(useCloudApiCompatibleUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_api_compatible", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_Comment(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_comment"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	comment := "comment for range template"
	commentUpdated := "comment for range template updated"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateComment(comment, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateComment(commentUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", commentUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_DdnsDomainname(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_ddns_domainname"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	ddnsDomainName := "aa.bb.com"
	ddnsDomainNameUpdated := "qq.ww.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateDdnsDomainname(ddnsDomainName, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", ddnsDomainName),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateDdnsDomainname(ddnsDomainNameUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", ddnsDomainNameUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_DdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_ddns_generate_hostname"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	ddnsGenerateHostname := true
	ddnsGenerateHostnameUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateDdnsGenerateHostname(ddnsGenerateHostname, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateDdnsGenerateHostname(ddnsGenerateHostnameUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_DelegatedMember(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_delegated_member"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	delegatedMember := ""
	delegatedMemberUpdated := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateDelegatedMember(delegatedMember, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegated_member", ""),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateDelegatedMember(delegatedMemberUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegated_member", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_DenyAllClients(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_deny_all_clients"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	denyAllClients := true
	denyAllClientsUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateDenyAllClients(denyAllClients, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_all_clients", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateDenyAllClients(denyAllClientsUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_all_clients", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_DenyBootp(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_deny_bootp"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	denyBootp := true
	denyBootpUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateDenyBootp(denyBootp, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateDenyBootp(denyBootpUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_EmailList(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_email_list"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)

	emailList := []string{"bbb@info.com", "aaa@wapi.com"}
	emailListUpdated := []string{"abc@info.com", "xyz@wapi.com"}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateEmailList(emailList, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_list.0", "bbb@info.com"),
					resource.TestCheckResourceAttr(resourceName, "email_list.1", "aaa@wapi.com"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateEmailList(emailListUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_list.0", "abc@info.com"),
					resource.TestCheckResourceAttr(resourceName, "email_list.1", "xyz@wapi.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_EnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_enable_ddns"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)

	enableDdns := true
	enableDdnsUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateEnableDdns(enableDdns, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateEnableDdns(enableDdnsUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_EnableDhcpThresholds(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_enable_dhcp_thresholds"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)

	enableDhcpThresholds := true
	enableDhcpThresholdsUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateEnableDhcpThresholds(enableDhcpThresholds, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_dhcp_thresholds", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateEnableDhcpThresholds(enableDhcpThresholdsUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_dhcp_thresholds", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_EnableEmailWarnings(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_enable_email_warnings"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)

	enableEmailWarnings := true
	enableEmailWarningsUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateEnableEmailWarnings(enableEmailWarnings, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_email_warnings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateEnableEmailWarnings(enableEmailWarningsUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_email_warnings", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_EnablePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_enable_pxe_lease_time"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)

	enablePxeLeaseTime := true
	enablePxeLeaseTimeUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateEnablePxeLeaseTime(enablePxeLeaseTime, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateEnablePxeLeaseTime(enablePxeLeaseTimeUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_EnableSnmpWarnings(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_enable_snmp_warnings"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)

	enableSnmpWarnings := true
	enableSnmpWarningsUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateEnableSnmpWarnings(enableSnmpWarnings, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_snmp_warnings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateEnableSnmpWarnings(enableSnmpWarningsUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_snmp_warnings", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

//func TestAccRangetemplateResource_Exclude(t *testing.T) {
//	var resourceName = "nios_dhcp_rangetemplate.test_exclude"
//	var v dhcp.Rangetemplate
//
//	name := acctest.RandomNameWithPrefix("range-template")
//	numberOfAdresses := int64(100)
//	offset := int64(50)
//
//	exclude := map[string]interface{}{}
//	excludeUpdated := map[string]interface{}{}
//
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:                 func() { acctest.PreCheck(t) },
//		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
//		Steps: []resource.TestStep{
//			// Create and Read
//			{
//				Config: testAccRangetemplateExclude(exclude, name, numberOfAdresses, offset),
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
//					resource.TestCheckResourceAttr(resourceName, "exclude", "EXCLUDE_REPLACE_ME"),
//				),
//			},
//			// Update and Read
//			{
//				Config: testAccRangetemplateExclude(excludeUpdated, name, numberOfAdresses, offset),
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
//					resource.TestCheckResourceAttr(resourceName, "exclude", "EXCLUDE_UPDATE_REPLACE_ME"),
//				),
//			},
//			// Delete testing automatically occurs in TestCase
//		},
//	})
//}

func TestAccRangetemplateResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_extattrs"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)

	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateExtAttrs(map[string]string{"Site": extAttrValue1}, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateExtAttrs(map[string]string{"Site": extAttrValue2}, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_FailoverAssociation(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_failover_association"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	failoverAssociation := "FALOVER"
	failoverAssociationUpdated := "FAILOVER_MS"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateFailoverAssociation(failoverAssociation, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "failover_association", "FAILOVER"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateFailoverAssociation(failoverAssociationUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "failover_association", "FAILOVER_MS"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_FingerprintFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_fingerprint_filter_rules"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	fingerPrintRules := ""
	fingerPrintRulesUpdated := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateFingerprintFilterRules(fingerPrintRules, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fingerprint_filter_rules", ""),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateFingerprintFilterRules(fingerPrintRulesUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fingerprint_filter_rules", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_HighWaterMark(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_high_water_mark"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)

	highWaterMark := int64(1000)
	highWaterMarkUpdated := int64(2000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateHighWaterMark(highWaterMark, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "1000"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateHighWaterMark(highWaterMarkUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "2000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_HighWaterMarkReset(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_high_water_mark_reset"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)

	highWaterMarkReset := int64(1000)
	highWaterMarkResetUpdated := int64(2000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateHighWaterMarkReset(highWaterMarkReset, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "1000"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateHighWaterMarkReset(highWaterMarkResetUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "2000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_IgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_ignore_dhcp_option_list_request"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)

	ignoreDhcpOptionListRequest := true
	ignoreDhcpOptionListRequestUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateIgnoreDhcpOptionListRequest(ignoreDhcpOptionListRequest, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateIgnoreDhcpOptionListRequest(ignoreDhcpOptionListRequestUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_KnownClients(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_known_clients"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)

	knownClients := "Allow"
	knownClientsUpdated := "Deny"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateKnownClients(knownClients, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "known_clients", "Allow"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateKnownClients(knownClientsUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "known_clients", "Deny"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_LeaseScavengeTime(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_lease_scavenge_time"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)

	leastScavngeTime := int64(3600)
	leaseScavengeTimeUpdated := int64(7200)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateLeaseScavengeTime(leastScavngeTime, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "3600"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateLeaseScavengeTime(leaseScavengeTimeUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "7200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_LogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_logic_filter_rules"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	logicFilterRules := ""
	logicFilterRulesUpdated := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateLogicFilterRules(logicFilterRules, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateLogicFilterRules(logicFilterRulesUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_LowWaterMark(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_low_water_mark"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)

	lowWaterMark := int64(7200)
	lowWaterMarkUpdated := int64(14400)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateLowWaterMark(lowWaterMark, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "7200"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateLowWaterMark(lowWaterMarkUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "14400"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_LowWaterMarkReset(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_low_water_mark_reset"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)

	lowMaterkReset := int64(36000)
	lowMaterkResetUpdated := int64(144000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateLowWaterMarkReset(lowMaterkReset, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "36000"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateLowWaterMarkReset(lowMaterkResetUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "144000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_MacFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_mac_filter_rules"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)

	macFilterRules := ""
	macFilterRulesUpdated := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateMacFilterRules(macFilterRules, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mac_filter_rules", "MAC_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateMacFilterRules(macFilterRulesUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mac_filter_rules", "MAC_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_Member(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_member"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)

	member := ""
	memberUpdated := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateMember(member, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member", "MEMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateMember(memberUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member", "MEMBER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_MsOptions(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_ms_options"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	msOptions := ""
	msOptionsUpdated := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateMsOptions(msOptions, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_options", "MS_OPTIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateMsOptions(msOptionsUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_options", "MS_OPTIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_MsServer(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_ms_server"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)

	msServer := "ms_server1"
	msServerUpdated := "ms_server2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateMsServer(msServer, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_server", "ms_server1"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateMsServer(msServerUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_server", "ms_server2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_NacFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_nac_filter_rules"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)

	nacFilterRules := ""
	nacFilterRulesUpdated := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateNacFilterRules(nacFilterRules, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nac_filter_rules", "NAC_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateNacFilterRules(nacFilterRulesUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nac_filter_rules", "NAC_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_Name(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_name"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)

	nameUpated := acctest.RandomNameWithPrefix("range-template-updated")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateName(name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateName(nameUpated, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_Nextserver(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_nextserver"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	nextServer := "next-server-1"
	nextServerUpdated := "next-server-2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateNextserver(nextServer, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", nextServer),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateNextserver(nextServerUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", nextServerUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_NumberOfAddresses(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_number_of_addresses"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	numberOfAdressesUpdated := int64(500)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateNumberOfAddresses(name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "number_of_addresses", "50"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateNumberOfAddresses(name, numberOfAdressesUpdated, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "number_of_addresses", "500"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_Offset(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_offset"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	offsetUpdated := int64(2000)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateOffset(name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "offset", "50"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateOffset(name, numberOfAdresses, offsetUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "offset", "2000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_OptionFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_option_filter_rules"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	optionFilterRules := ""
	optionFilterRulesUpdated := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateOptionFilterRules(optionFilterRules, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_filter_rules", "OPTION_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateOptionFilterRules(optionFilterRulesUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "option_filter_rules", "OPTION_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_Options(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_options"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	options := ""
	optionsUpdated := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateOptions(options, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateOptions(optionsUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_PxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_pxe_lease_time"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	pxeLeaseTime := int64(3600)
	pxeLeaseTimeUpdated := int64(7200)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplatePxeLeaseTime(pxeLeaseTime, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "3600"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplatePxeLeaseTime(pxeLeaseTimeUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "7200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_RecycleLeases(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_recycle_leases"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	recycleLeases := false
	recycleLeasesUpdated := true

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateRecycleLeases(recycleLeases, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateRecycleLeases(recycleLeasesUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_RelayAgentFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_relay_agent_filter_rules"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	relayAgentFilterRules := ""
	relayAgentFilterRulesUpdated := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateRelayAgentFilterRules(relayAgentFilterRules, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "relay_agent_filter_rules", "RELAY_AGENT_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateRelayAgentFilterRules(relayAgentFilterRulesUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "relay_agent_filter_rules", "RELAY_AGENT_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_ServerAssociationType(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_server_association_type"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	servserAssociationType := "FAILOVER"
	servserAssociationTypeUpdated := "NONE"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateServerAssociationType(servserAssociationType, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "server_association_type", "FAILOVER"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateServerAssociationType(servserAssociationTypeUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "server_association_type", "NONE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_UnknownClients(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_unknown_clients"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	unknownClients := "Deny"
	unknownClientsUpdated := "Allow"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateUnknownClients(unknownClients, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "unknown_clients", "Deny"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateUnknownClients(unknownClientsUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "unknown_clients", "Allow"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_UpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_update_dns_on_lease_renewal"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	updateDnsOnLeaseRenewal := true
	updateDnsOnLeaseRenewalUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateUpdateDnsOnLeaseRenewal(updateDnsOnLeaseRenewal, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateUpdateDnsOnLeaseRenewal(updateDnsOnLeaseRenewalUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_UseBootfile(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_use_bootfile"
	var v dhcp.Rangetemplate

	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	useBootFile := true
	useBootFileUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateUseBootfile(useBootFile, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateUseBootfile(useBootFileUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_UseBootserver(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_use_bootserver"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	useBootServer := true
	useBootServerUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateUseBootserver(useBootServer, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateUseBootserver(useBootServerUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_UseDdnsDomainname(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_use_ddns_domainname"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	useDdnsDomainName := true
	useDdnsDomainNameUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateUseDdnsDomainname(useDdnsDomainName, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateUseDdnsDomainname(useDdnsDomainNameUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_UseDdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_use_ddns_generate_hostname"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	useDdnsGenerateHostName := true
	useDdnsGenerateHostNameUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateUseDdnsGenerateHostname(useDdnsGenerateHostName, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateUseDdnsGenerateHostname(useDdnsGenerateHostNameUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_UseDenyBootp(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_use_deny_bootp"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	useDenyBootp := true
	useDenyBootpUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateUseDenyBootp(useDenyBootp, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateUseDenyBootp(useDenyBootpUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_UseEmailList(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_use_email_list"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	useEmailList := true
	useEmailListUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateUseEmailList(useEmailList, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_email_list", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateUseEmailList(useEmailListUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_email_list", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_UseEnableDdns(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_use_enable_ddns"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	useEnableDns := true
	useEnableDnsUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateUseEnableDdns(useEnableDns, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateUseEnableDdns(useEnableDnsUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_UseEnableDhcpThresholds(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_use_enable_dhcp_thresholds"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	useEnableDhcpThreshold := true
	useEnableDhcpThresholdUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateUseEnableDhcpThresholds(useEnableDhcpThreshold, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateUseEnableDhcpThresholds(useEnableDhcpThresholdUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_UseIgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_use_ignore_dhcp_option_list_request"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	useIgnoreDhcpOptionListRequest := true
	useIgnoreDhcpOptionListRequestUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateUseIgnoreDhcpOptionListRequest(useIgnoreDhcpOptionListRequest, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateUseIgnoreDhcpOptionListRequest(useIgnoreDhcpOptionListRequestUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_UseKnownClients(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_use_known_clients"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	useKnownClients := true
	useKnownClientsUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateUseKnownClients(useKnownClients, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_known_clients", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateUseKnownClients(useKnownClientsUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_known_clients", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_UseLeaseScavengeTime(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_use_lease_scavenge_time"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	useLeaseScavngeTime := true
	useLeaseScavngeTimeUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateUseLeaseScavengeTime(useLeaseScavngeTime, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateUseLeaseScavengeTime(useLeaseScavngeTimeUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_UseLogicFilterRules(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_use_logic_filter_rules"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	useLogicFilterRules := true
	useLogicFilterRulesUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateUseLogicFilterRules(useLogicFilterRules, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateUseLogicFilterRules(useLogicFilterRulesUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_UseMsOptions(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_use_ms_options"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	useMsOptions := true
	useMsOptionsUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateUseMsOptions(useMsOptions, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ms_options", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateUseMsOptions(useMsOptionsUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ms_options", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_UseNextserver(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_use_nextserver"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	useNextServer := true
	useNextServerUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateUseNextserver(useNextServer, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateUseNextserver(useNextServerUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_UseOptions(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_use_options"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	useOptions := true
	useOptionsUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateUseOptions(useOptions, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateUseOptions(useOptionsUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_UsePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_use_pxe_lease_time"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	usePxeLeaseTime := true
	usePxeLeaseTimeUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateUsePxeLeaseTime(usePxeLeaseTime, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateUsePxeLeaseTime(usePxeLeaseTimeUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_UseRecycleLeases(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_use_recycle_leases"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	useRecycleLeases := true
	useRecycleLeasesUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateUseRecycleLeases(useRecycleLeases, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateUseRecycleLeases(useRecycleLeasesUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_UseUnknownClients(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_use_unknown_clients"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	useUnknownClients := true
	useUnknownClientsUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateUseUnknownClients(useUnknownClients, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_unknown_clients", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateUseUnknownClients(useUnknownClientsUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_unknown_clients", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRangetemplateResource_UseUpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_dhcp_rangetemplate.test_use_update_dns_on_lease_renewal"
	var v dhcp.Rangetemplate
	name := acctest.RandomNameWithPrefix("range-template")
	numberOfAdresses := int64(100)
	offset := int64(50)
	useUpdateDnsOnLeaseRenewal := true
	useUpdateDnsOnLeaseRenewalUpdated := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRangetemplateUseUpdateDnsOnLeaseRenewal(useUpdateDnsOnLeaseRenewal, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccRangetemplateUseUpdateDnsOnLeaseRenewal(useUpdateDnsOnLeaseRenewalUpdated, name, numberOfAdresses, offset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangetemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckRangetemplateExists(ctx context.Context, resourceName string, v *dhcp.Rangetemplate) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DHCPAPI.
			RangetemplateAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForRangetemplate).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetRangetemplateResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetRangetemplateResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckRangetemplateDestroy(ctx context.Context, v *dhcp.Rangetemplate) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DHCPAPI.
			RangetemplateAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForRangetemplate).
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

func testAccCheckRangetemplateDisappears(ctx context.Context, v *dhcp.Rangetemplate) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DHCPAPI.
			RangetemplateAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccRangetemplateBasicConfig(name string, numberOfAddresses, offset int64) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test" {
	name = %q
	number_of_addresses = %d
	offset = %d
}
`, name, numberOfAddresses, offset)
}

func testAccRangetemplateBootfile(useBootFile bool, bootfile, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_bootfile" {
    use_bootfile = %t
    bootfile = %q
	name = %q
	number_of_addresses = %d
	offset = %d
}
`, useBootFile, bootfile, name, numberOfAddresses, offset)
}

func testAccRangetemplateBootserver(useBootServer bool, bootServer, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_bootserver" {
    use_bootserver = %t
    bootserver = %q
	name = %q
	number_of_addresses = %d
	offset = %d
}
`, useBootServer, bootServer, name, numberOfAddresses, offset)
}

func testAccRangetemplateCloudApiCompatible(cloudApiCompatible bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_cloud_api_compatible" {
   cloud_api_compatible = %t
	name = %q
	number_of_addresses = %d
	offset = %d
}
`, cloudApiCompatible, name, numberOfAddresses, offset)
}

func testAccRangetemplateComment(comment, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_comment" {
   comment = %q
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, comment, name, numberOfAddresses, offset)
}

func testAccRangetemplateDdnsDomainname(ddnsDomainname, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_ddns_domainname" {
   ddns_domainname = %q
	name = %q
	number_of_addresses = %d
	offset = %d
}
`, ddnsDomainname, name, numberOfAddresses, offset)
}

func testAccRangetemplateDdnsGenerateHostname(ddnsGenerateHostname bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_ddns_generate_hostname" {
   ddns_generate_hostname = %t
	name = %q
	number_of_addresses = %d
	offset = %d
}
`, ddnsGenerateHostname, name, numberOfAddresses, offset)
}

func testAccRangetemplateDelegatedMember(delegatedMember, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_delegated_member" {
   delegated_member = %q
	name = %q
	number_of_addresses = %d
	offset = %d
}
`, delegatedMember, name, numberOfAddresses, offset)
}

func testAccRangetemplateDenyAllClients(denyAllClients bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_deny_all_clients" {
   deny_all_clients = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, denyAllClients, name, numberOfAddresses, offset)
}

func testAccRangetemplateDenyBootp(denyBootp bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_deny_bootp" {
   deny_bootp = %t
	name = %q
	number_of_addresses = %d
	offset = %d
}
`, denyBootp, name, numberOfAddresses, offset)
}

func testAccRangetemplateEmailList(emailList []string, name string, numberOfAddresses, offset int64) string {
	emailListStr := `{"` + strings.Join(emailList, `", "`) + `"}`
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_email_list" {
   email_list = %q
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, emailListStr, name, numberOfAddresses, offset)
}

func testAccRangetemplateEnableDdns(enableDdns bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_enable_ddns" {
   enable_ddns = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, enableDdns, name, numberOfAddresses, offset)
}

func testAccRangetemplateEnableDhcpThresholds(enableDhcpThresholds bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_enable_dhcp_thresholds" {
   enable_dhcp_thresholds = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, enableDhcpThresholds, name, numberOfAddresses, offset)
}

func testAccRangetemplateEnableEmailWarnings(enableEmailWarnings bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_enable_email_warnings" {
   enable_email_warnings = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, enableEmailWarnings, name, numberOfAddresses, offset)
}

func testAccRangetemplateEnablePxeLeaseTime(enablePxeLeaseTime bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_enable_pxe_lease_time" {
   enable_pxe_lease_time = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, enablePxeLeaseTime, name, numberOfAddresses, offset)
}

func testAccRangetemplateEnableSnmpWarnings(enableSnmpWarnings bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_enable_snmp_warnings" {
   enable_snmp_warnings = %t
   name = %q
	number_of_addresses = %d
	offset = %d
}
`, enableSnmpWarnings, name, numberOfAddresses, offset)
}

func testAccRangetemplateExclude(exclude, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_exclude" {
   exclude = %q
   name = %q
	number_of_addresses = %d
	offset = %d
}
`, exclude, name, numberOfAddresses, offset)
}

func testAccRangetemplateExtAttrs(extAttrs map[string]string, name string, numberOfAddresses, offset int64) string {
	extattrsStr := "{"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`%s = %q`, k, v)
	}
	extattrsStr += "}"
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_extattrs" {
   extattrs = %q
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, extattrsStr, name, numberOfAddresses, offset)
}

func testAccRangetemplateFailoverAssociation(failoverAssociation, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_failover_association" {
   failover_association = %q
   name = %q
	number_of_addresses = %d
	offset = %d
}
`, failoverAssociation, name, numberOfAddresses, offset)
}

func testAccRangetemplateFingerprintFilterRules(fingerprintFilterRules, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_fingerprint_filter_rules" {
   fingerprint_filter_rules = %q
   name = %q
	number_of_addresses = %d
	offset = %d
}
`, fingerprintFilterRules, name, numberOfAddresses, offset)
}

func testAccRangetemplateHighWaterMark(highWaterMark int64, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_high_water_mark" {
   high_water_mark = %d
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, highWaterMark, name, numberOfAddresses, offset)
}

func testAccRangetemplateHighWaterMarkReset(highWaterMarkReset int64, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_high_water_mark_reset" {
   high_water_mark_reset = %d
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, highWaterMarkReset, name, numberOfAddresses, offset)
}

func testAccRangetemplateIgnoreDhcpOptionListRequest(ignoreDhcpOptionListRequest bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_ignore_dhcp_option_list_request" {
   ignore_dhcp_option_list_request = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, ignoreDhcpOptionListRequest, name, numberOfAddresses, offset)
}

func testAccRangetemplateKnownClients(knownClients, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_known_clients" {
   known_clients = %q
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, knownClients, name, numberOfAddresses, offset)
}

func testAccRangetemplateLeaseScavengeTime(leaseScavengeTime int64, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_lease_scavenge_time" {
   lease_scavenge_time = %d
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, leaseScavengeTime, name, numberOfAddresses, offset)
}

func testAccRangetemplateLogicFilterRules(logicFilterRules, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_logic_filter_rules" {
   logic_filter_rules = %q
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, logicFilterRules, name, numberOfAddresses, offset)
}

func testAccRangetemplateLowWaterMark(lowWaterMark int64, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_low_water_mark" {
   low_water_mark = %d
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, lowWaterMark, name, numberOfAddresses, offset)
}

func testAccRangetemplateLowWaterMarkReset(lowWaterMarkReset int64, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_low_water_mark_reset" {
   low_water_mark_reset = %d
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, lowWaterMarkReset, name, numberOfAddresses, offset)
}

func testAccRangetemplateMacFilterRules(macFilterRules, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_mac_filter_rules" {
   mac_filter_rules = %q
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, macFilterRules, name, numberOfAddresses, offset)
}

func testAccRangetemplateMember(member, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_member" {
   member = %q
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, member, name, numberOfAddresses, offset)
}

func testAccRangetemplateMsOptions(msOptions, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_ms_options" {
   ms_options = %q
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, msOptions, name, numberOfAddresses, offset)
}

func testAccRangetemplateMsServer(msServer, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_ms_server" {
   ms_server = %q
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, msServer, name, numberOfAddresses, offset)
}

func testAccRangetemplateNacFilterRules(nacFilterRules, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_nac_filter_rules" {
   nac_filter_rules = %q
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, nacFilterRules, name, numberOfAddresses, offset)
}

func testAccRangetemplateName(name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_name" {
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, name, numberOfAddresses, offset)
}

func testAccRangetemplateNextserver(nextserver, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_nextserver" {
   nextserver = %q
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, nextserver, name, numberOfAddresses, offset)
}

func testAccRangetemplateNumberOfAddresses(name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_number_of_addresses" {
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, name, numberOfAddresses, offset)
}

func testAccRangetemplateOffset(name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_offset" {
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, name, numberOfAddresses, offset)
}

func testAccRangetemplateOptionFilterRules(optionFilterRules, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_option_filter_rules" {
   option_filter_rules = %q
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, optionFilterRules, name, numberOfAddresses, offset)
}

func testAccRangetemplateOptions(options, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_options" {
   options = %q
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, options, name, numberOfAddresses, offset)
}

func testAccRangetemplatePxeLeaseTime(pxeLeaseTime int64, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_pxe_lease_time" {
   pxe_lease_time = %d
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, pxeLeaseTime, name, numberOfAddresses, offset)
}

func testAccRangetemplateRecycleLeases(recycleLeases bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_recycle_leases" {
   recycle_leases = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, recycleLeases, name, numberOfAddresses, offset)
}

func testAccRangetemplateRelayAgentFilterRules(relayAgentFilterRules, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_relay_agent_filter_rules" {
   relay_agent_filter_rules = %q
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, relayAgentFilterRules, name, numberOfAddresses, offset)
}

func testAccRangetemplateServerAssociationType(serverAssociationType, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_server_association_type" {
   server_association_type = %q
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, serverAssociationType, name, numberOfAddresses, offset)
}

func testAccRangetemplateUnknownClients(unknownClients, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_unknown_clients" {
   unknown_clients = %q
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, unknownClients, name, numberOfAddresses, offset)
}

func testAccRangetemplateUpdateDnsOnLeaseRenewal(updateDnsOnLeaseRenewal bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_update_dns_on_lease_renewal" {
   update_dns_on_lease_renewal = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, updateDnsOnLeaseRenewal, name, numberOfAddresses, offset)
}

func testAccRangetemplateUseBootfile(useBootfile bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_use_bootfile" {
   use_bootfile = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, useBootfile, name, numberOfAddresses, offset)
}

func testAccRangetemplateUseBootserver(useBootserver bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_use_bootserver" {
   use_bootserver = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, useBootserver, name, numberOfAddresses, offset)
}

func testAccRangetemplateUseDdnsDomainname(useDdnsDomainname bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_use_ddns_domainname" {
   use_ddns_domainname = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, useDdnsDomainname, name, numberOfAddresses, offset)
}

func testAccRangetemplateUseDdnsGenerateHostname(useDdnsGenerateHostname bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_use_ddns_generate_hostname" {
   use_ddns_generate_hostname = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, useDdnsGenerateHostname, name, numberOfAddresses, offset)
}

func testAccRangetemplateUseDenyBootp(useDenyBootp bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_use_deny_bootp" {
   use_deny_bootp = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, useDenyBootp, name, numberOfAddresses, offset)
}

func testAccRangetemplateUseEmailList(useEmailList bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_use_email_list" {
   use_email_list = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, useEmailList, name, numberOfAddresses, offset)
}

func testAccRangetemplateUseEnableDdns(useEnableDdns bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_use_enable_ddns" {
   use_enable_ddns = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, useEnableDdns, name, numberOfAddresses, offset)
}

func testAccRangetemplateUseEnableDhcpThresholds(useEnableDhcpThresholds bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_use_enable_dhcp_thresholds" {
   use_enable_dhcp_thresholds = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, useEnableDhcpThresholds, name, numberOfAddresses, offset)
}

func testAccRangetemplateUseIgnoreDhcpOptionListRequest(useIgnoreDhcpOptionListRequest bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_use_ignore_dhcp_option_list_request" {
   use_ignore_dhcp_option_list_request = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, useIgnoreDhcpOptionListRequest, name, numberOfAddresses, offset)
}

func testAccRangetemplateUseKnownClients(useKnownClients bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_use_known_clients" {
   use_known_clients = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, useKnownClients, name, numberOfAddresses, offset)
}

func testAccRangetemplateUseLeaseScavengeTime(useLeaseScavengeTime bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_use_lease_scavenge_time" {
   use_lease_scavenge_time = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, useLeaseScavengeTime, name, numberOfAddresses, offset)
}

func testAccRangetemplateUseLogicFilterRules(useLogicFilterRules bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_use_logic_filter_rules" {
   use_logic_filter_rules = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, useLogicFilterRules, name, numberOfAddresses, offset)
}

func testAccRangetemplateUseMsOptions(useMsOptions bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_use_ms_options" {
   use_ms_options = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, useMsOptions, name, numberOfAddresses, offset)
}

func testAccRangetemplateUseNextserver(useNextserver bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_use_nextserver" {
   use_nextserver = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, useNextserver, name, numberOfAddresses, offset)
}

func testAccRangetemplateUseOptions(useOptions bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_use_options" {
   use_options = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, useOptions, name, numberOfAddresses, offset)
}

func testAccRangetemplateUsePxeLeaseTime(usePxeLeaseTime bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_use_pxe_lease_time" {
   use_pxe_lease_time = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, usePxeLeaseTime, name, numberOfAddresses, offset)
}

func testAccRangetemplateUseRecycleLeases(useRecycleLeases bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_use_recycle_leases" {
   use_recycle_leases = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, useRecycleLeases, name, numberOfAddresses, offset)
}

func testAccRangetemplateUseUnknownClients(useUnknownClients bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_use_unknown_clients" {
   use_unknown_clients = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, useUnknownClients, name, numberOfAddresses, offset)
}

func testAccRangetemplateUseUpdateDnsOnLeaseRenewal(useUpdateDnsOnLeaseRenewal bool, name string, numberOfAddresses, offset int64) string {
	return fmt.Sprintf(`
resource "nios_dhcp_rangetemplate" "test_use_update_dns_on_lease_renewal" {
   use_update_dns_on_lease_renewal = %t
   name = %q
   number_of_addresses = %d
   offset = %d
}
`, useUpdateDnsOnLeaseRenewal, name, numberOfAddresses, offset)
}
