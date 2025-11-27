package ipam_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

/*
// Manage ipam Networktemplate with Basic Fields
resource "nios_ipam_networktemplate" "ipam_networktemplate_basic" {
    name = "NAME_REPLACE_ME"
}

// Manage ipam Networktemplate with Additional Fields
resource "nios_ipam_networktemplate" "ipam_networktemplate_with_additional_fields" {
    name = "NAME_REPLACE_ME"

// TODO : Add additional optional fields below

    //Extensible Attributes
    extattrs = {
        Site = "location-1"
    }
}
*/

var readableAttributesForNetworktemplate = "allow_any_netmask,authority,auto_create_reversezone,bootfile,bootserver,cloud_api_compatible,comment,ddns_domainname,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,ddns_update_fixed_addresses,ddns_use_option81,delegated_member,deny_bootp,email_list,enable_ddns,enable_dhcp_thresholds,enable_email_warnings,enable_pxe_lease_time,enable_snmp_warnings,extattrs,fixed_address_templates,high_water_mark,high_water_mark_reset,ignore_dhcp_option_list_request,ipam_email_addresses,ipam_threshold_settings,ipam_trap_settings,lease_scavenge_time,logic_filter_rules,low_water_mark,low_water_mark_reset,members,name,netmask,nextserver,options,pxe_lease_time,range_templates,recycle_leases,rir,rir_organization,rir_registration_action,rir_registration_status,send_rir_request,update_dns_on_lease_renewal,use_authority,use_bootfile,use_bootserver,use_ddns_domainname,use_ddns_generate_hostname,use_ddns_ttl,use_ddns_update_fixed_addresses,use_ddns_use_option81,use_deny_bootp,use_email_list,use_enable_ddns,use_enable_dhcp_thresholds,use_ignore_dhcp_option_list_request,use_ipam_email_addresses,use_ipam_threshold_settings,use_ipam_trap_settings,use_lease_scavenge_time,use_logic_filter_rules,use_nextserver,use_options,use_pxe_lease_time,use_recycle_leases,use_update_dns_on_lease_renewal"

func TestAccNetworktemplateResource_basic(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateBasicConfig("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "allow_any_netmask", "false"),
					resource.TestCheckResourceAttr(resourceName, "authority", "false"),
					resource.TestCheckResourceAttr(resourceName, "auto_create_reversezone", "false"),
					resource.TestCheckResourceAttr(resourceName, "cloud_api_compatible", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "true"),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "0"),
					resource.TestCheckResourceAttr(resourceName, "ddns_update_fixed_addresses", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "false"),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_dhcp_thresholds", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_email_warnings", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_snmp_warnings", "false"),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "95"),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "85"),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "false"),
					// TODO : Add validation for default value for field ipam_threshold_settings if applicable
					// TODO : Add validation for default value for field ipam_trap_settings if applicable
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "-1"),
					// TODO : Add validation for default value for field logic_filter_rules if applicable
					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "0"),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "10"),
					// TODO : Add validation for default value for field members if applicable
					// TODO : Add validation for default value for field options if applicable
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "true"),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_action", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_status", "NOT_REGISTERED"),
					resource.TestCheckResourceAttr(resourceName, "send_rir_request", "false"),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_authority", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_update_fixed_addresses", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_email_list", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_email_addresses", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_threshold_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_trap_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_disappears(t *testing.T) {
	resourceName := "nios_ipam_networktemplate.test"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetworktemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworktemplateBasicConfig("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					testAccCheckNetworktemplateDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNetworktemplateResource_Import(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateBasicConfig("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
				),
			},
			// Import with PlanOnly to detect differences
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccNetworktemplateImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "ref",
				PlanOnly:                             true,
				//ExpectError:                          regexp.MustCompile(`ImportStateVerify attributes not equivalent`),
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccNetworktemplateImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_AllowAnyNetmask(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_allow_any_netmask"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateAllowAnyNetmask("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_any_netmask", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateAllowAnyNetmask("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_any_netmask", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_Authority(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_authority"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateAuthority("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "authority", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateAuthority("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "authority", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_AutoCreateReversezone(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_auto_create_reversezone"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateAutoCreateReversezone("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_create_reversezone", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateAutoCreateReversezone("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_create_reversezone", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_Bootfile(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_bootfile"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateBootfile("NAME_REPLACE_ME", "BOOTFILE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "BOOTFILE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateBootfile("NAME_REPLACE_ME", "BOOTFILE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "BOOTFILE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_Bootserver(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_bootserver"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateBootserver("NAME_REPLACE_ME", "BOOTSERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "BOOTSERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateBootserver("NAME_REPLACE_ME", "BOOTSERVER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "BOOTSERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_CloudApiCompatible(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_cloud_api_compatible"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateCloudApiCompatible("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_api_compatible", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateCloudApiCompatible("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_api_compatible", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_Comment(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_comment"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateComment("NAME_REPLACE_ME", "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateComment("NAME_REPLACE_ME", "Updated comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the object"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_DdnsDomainname(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_ddns_domainname"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateDdnsDomainname("NAME_REPLACE_ME", "DDNS_DOMAINNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "DDNS_DOMAINNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateDdnsDomainname("NAME_REPLACE_ME", "DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_DdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_ddns_generate_hostname"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateDdnsGenerateHostname("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateDdnsGenerateHostname("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_DdnsServerAlwaysUpdates(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_ddns_server_always_updates"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateDdnsServerAlwaysUpdates("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateDdnsServerAlwaysUpdates("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_DdnsTtl(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_ddns_ttl"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateDdnsTtl("NAME_REPLACE_ME", "DDNS_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "DDNS_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateDdnsTtl("NAME_REPLACE_ME", "DDNS_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "DDNS_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworktemplateResource_DdnsUpdateFixedAddresses(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_ddns_update_fixed_addresses"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateDdnsUpdateFixedAddresses("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_update_fixed_addresses", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateDdnsUpdateFixedAddresses("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_update_fixed_addresses", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_DdnsUseOption81(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_ddns_use_option81"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateDdnsUseOption81("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateDdnsUseOption81("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_DelegatedMember(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_delegated_member"
	var v ipam.Networktemplate
	delegatedMemberVal := map[string]string{}
	delegatedMemberValUpdated := map[string]string{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateDelegatedMember("NAME_REPLACE_ME", delegatedMemberVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegated_member", "DELEGATED_MEMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateDelegatedMember("NAME_REPLACE_ME", delegatedMemberValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegated_member", "DELEGATED_MEMBER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworktemplateResource_DenyBootp(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_deny_bootp"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateDenyBootp("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateDenyBootp("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_EmailList(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_email_list"
	var v ipam.Networktemplate
	emailListVal := []string{"EMAIL_LIST_REPLACE_ME1", "EMAIL_LIST_REPLACE_ME2"}
	emailListValUpdated := []string{"EMAIL_LIST_REPLACE_ME1", "EMAIL_LIST_REPLACE_ME2"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateEmailList("NAME_REPLACE_ME", emailListVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_list", "EMAIL_LIST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateEmailList("NAME_REPLACE_ME", emailListValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_list", "EMAIL_LIST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_EnableDdns(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_enable_ddns"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateEnableDdns("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateEnableDdns("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_EnableDhcpThresholds(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_enable_dhcp_thresholds"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateEnableDhcpThresholds("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_dhcp_thresholds", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateEnableDhcpThresholds("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_dhcp_thresholds", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_EnableEmailWarnings(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_enable_email_warnings"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateEnableEmailWarnings("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_email_warnings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateEnableEmailWarnings("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_email_warnings", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_EnablePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_enable_pxe_lease_time"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateEnablePxeLeaseTime("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateEnablePxeLeaseTime("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_EnableSnmpWarnings(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_enable_snmp_warnings"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateEnableSnmpWarnings("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_snmp_warnings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateEnableSnmpWarnings("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_snmp_warnings", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_extattrs"
	var v ipam.Networktemplate
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateExtAttrs("NAME_REPLACE_ME", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateExtAttrs("NAME_REPLACE_ME", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_FixedAddressTemplates(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_fixed_address_templates"
	var v ipam.Networktemplate
	fixedAddressTemplatesVal := []string{"FIXED_ADDRESS_TEMPLATES_REPLACE_ME1", "FIXED_ADDRESS_TEMPLATES_REPLACE_ME2"}
	fixedAddressTemplatesValUpdated := []string{"FIXED_ADDRESS_TEMPLATES_REPLACE_ME1", "FIXED_ADDRESS_TEMPLATES_REPLACE_ME2"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateFixedAddressTemplates("NAME_REPLACE_ME", fixedAddressTemplatesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fixed_address_templates", "FIXED_ADDRESS_TEMPLATES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateFixedAddressTemplates("NAME_REPLACE_ME", fixedAddressTemplatesValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fixed_address_templates", "FIXED_ADDRESS_TEMPLATES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_HighWaterMark(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_high_water_mark"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateHighWaterMark("NAME_REPLACE_ME", "HIGH_WATER_MARK_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "HIGH_WATER_MARK_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateHighWaterMark("NAME_REPLACE_ME", "HIGH_WATER_MARK_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "HIGH_WATER_MARK_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworktemplateResource_HighWaterMarkReset(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_high_water_mark_reset"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateHighWaterMarkReset("NAME_REPLACE_ME", "HIGH_WATER_MARK_RESET_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "HIGH_WATER_MARK_RESET_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateHighWaterMarkReset("NAME_REPLACE_ME", "HIGH_WATER_MARK_RESET_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "HIGH_WATER_MARK_RESET_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworktemplateResource_IgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_ignore_dhcp_option_list_request"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateIgnoreDhcpOptionListRequest("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateIgnoreDhcpOptionListRequest("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_IpamEmailAddresses(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_ipam_email_addresses"
	var v ipam.Networktemplate
	ipamEmailAddressesVal := []string{"IPAM_EMAIL_ADDRESSES_REPLACE_ME1", "IPAM_EMAIL_ADDRESSES_REPLACE_ME2"}
	ipamEmailAddressesValUpdated := []string{"IPAM_EMAIL_ADDRESSES_REPLACE_ME1", "IPAM_EMAIL_ADDRESSES_REPLACE_ME2"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateIpamEmailAddresses("NAME_REPLACE_ME", ipamEmailAddressesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_email_addresses", "IPAM_EMAIL_ADDRESSES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateIpamEmailAddresses("NAME_REPLACE_ME", ipamEmailAddressesValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_email_addresses", "IPAM_EMAIL_ADDRESSES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_IpamThresholdSettings(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_ipam_threshold_settings"
	var v ipam.Networktemplate
	ipamThresholdSettingsVal := map[string]string{}
	ipamThresholdSettingsValUpdated := map[string]string{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateIpamThresholdSettings("NAME_REPLACE_ME", ipamThresholdSettingsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_threshold_settings", "IPAM_THRESHOLD_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateIpamThresholdSettings("NAME_REPLACE_ME", ipamThresholdSettingsValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_threshold_settings", "IPAM_THRESHOLD_SETTINGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworktemplateResource_IpamTrapSettings(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_ipam_trap_settings"
	var v ipam.Networktemplate
	ipamTrapSettingsVal := map[string]string{}
	ipamTrapSettingsValUpdated := map[string]string{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateIpamTrapSettings("NAME_REPLACE_ME", ipamTrapSettingsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_trap_settings", "IPAM_TRAP_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateIpamTrapSettings("NAME_REPLACE_ME", ipamTrapSettingsValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_trap_settings", "IPAM_TRAP_SETTINGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworktemplateResource_LeaseScavengeTime(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_lease_scavenge_time"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateLeaseScavengeTime("NAME_REPLACE_ME", "LEASE_SCAVENGE_TIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "LEASE_SCAVENGE_TIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateLeaseScavengeTime("NAME_REPLACE_ME", "LEASE_SCAVENGE_TIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "LEASE_SCAVENGE_TIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworktemplateResource_LogicFilterRules(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_logic_filter_rules"
	var v ipam.Networktemplate
	logicFilterRulesVal := []map[string]any{}
	logicFilterRulesValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateLogicFilterRules("NAME_REPLACE_ME", logicFilterRulesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateLogicFilterRules("NAME_REPLACE_ME", logicFilterRulesValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_LowWaterMark(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_low_water_mark"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateLowWaterMark("NAME_REPLACE_ME", "LOW_WATER_MARK_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "LOW_WATER_MARK_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateLowWaterMark("NAME_REPLACE_ME", "LOW_WATER_MARK_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "LOW_WATER_MARK_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworktemplateResource_LowWaterMarkReset(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_low_water_mark_reset"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateLowWaterMarkReset("NAME_REPLACE_ME", "LOW_WATER_MARK_RESET_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "LOW_WATER_MARK_RESET_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateLowWaterMarkReset("NAME_REPLACE_ME", "LOW_WATER_MARK_RESET_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "LOW_WATER_MARK_RESET_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworktemplateResource_Members(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_members"
	var v ipam.Networktemplate
	membersVal := []map[string]any{}
	membersValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateMembers("NAME_REPLACE_ME", membersVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "members", "MEMBERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateMembers("NAME_REPLACE_ME", membersValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "members", "MEMBERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_Name(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_name"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_Netmask(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_netmask"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateNetmask("NAME_REPLACE_ME", "NETMASK_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "netmask", "NETMASK_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateNetmask("NAME_REPLACE_ME", "NETMASK_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "netmask", "NETMASK_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworktemplateResource_Nextserver(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_nextserver"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateNextserver("NAME_REPLACE_ME", "NEXTSERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", "NEXTSERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateNextserver("NAME_REPLACE_ME", "NEXTSERVER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", "NEXTSERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_Options(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_options"
	var v ipam.Networktemplate
	optionsVal := []map[string]any{}
	optionsValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateOptions("NAME_REPLACE_ME", optionsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateOptions("NAME_REPLACE_ME", optionsValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_PxeLeaseTime(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_pxe_lease_time"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplatePxeLeaseTime("NAME_REPLACE_ME", "PXE_LEASE_TIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "PXE_LEASE_TIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplatePxeLeaseTime("NAME_REPLACE_ME", "PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "PXE_LEASE_TIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworktemplateResource_RangeTemplates(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_range_templates"
	var v ipam.Networktemplate
	rangeTemplatesVal := []string{"RANGE_TEMPLATES_REPLACE_ME1", "RANGE_TEMPLATES_REPLACE_ME2"}
	rangeTemplatesValUpdated := []string{"RANGE_TEMPLATES_REPLACE_ME1", "RANGE_TEMPLATES_REPLACE_ME2"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateRangeTemplates("NAME_REPLACE_ME", rangeTemplatesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "range_templates", "RANGE_TEMPLATES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateRangeTemplates("NAME_REPLACE_ME", rangeTemplatesValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "range_templates", "RANGE_TEMPLATES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_RecycleLeases(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_recycle_leases"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateRecycleLeases("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateRecycleLeases("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_RirOrganization(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_rir_organization"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateRirOrganization("NAME_REPLACE_ME", "RIR_ORGANIZATION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_organization", "RIR_ORGANIZATION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateRirOrganization("NAME_REPLACE_ME", "RIR_ORGANIZATION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_organization", "RIR_ORGANIZATION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_RirRegistrationAction(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_rir_registration_action"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateRirRegistrationAction("RIR_REGISTRATION_ACTION_REPLACE_ME", "CREATE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_action", "CREATE"),
				),
			},
			{
				Config: testAccNetworktemplateRirRegistrationAction("RIR_REGISTRATION_ACTION_REPLACE_ME", "NONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_action", "NONE"),
				),
			},
		},
	})
}

func TestAccNetworktemplateResource_RirRegistrationStatus(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_rir_registration_status"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateRirRegistrationStatus("RIR_REGISTRATION_STATUS_REPLACE_ME", "NOT_REGISTERED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_status", "NOT_REGISTERED"),
				),
			},
			{
				Config: testAccNetworktemplateRirRegistrationStatus("RIR_REGISTRATION_STATUS_REPLACE_ME", "REGISTERED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_status", "REGISTERED"),
				),
			},
		},
	})
}

func TestAccNetworktemplateResource_SendRirRequest(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_send_rir_request"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateSendRirRequest("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "send_rir_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateSendRirRequest("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "send_rir_request", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_update_dns_on_lease_renewal"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUpdateDnsOnLeaseRenewal("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUpdateDnsOnLeaseRenewal("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseAuthority(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_authority"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseAuthority("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_authority", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseAuthority("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_authority", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseBootfile(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_bootfile"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseBootfile("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseBootfile("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseBootserver(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_bootserver"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseBootserver("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseBootserver("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseDdnsDomainname(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_ddns_domainname"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseDdnsDomainname("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseDdnsDomainname("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseDdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_ddns_generate_hostname"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseDdnsGenerateHostname("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseDdnsGenerateHostname("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseDdnsTtl(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_ddns_ttl"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseDdnsTtl("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseDdnsTtl("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseDdnsUpdateFixedAddresses(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_ddns_update_fixed_addresses"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseDdnsUpdateFixedAddresses("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_update_fixed_addresses", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseDdnsUpdateFixedAddresses("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_update_fixed_addresses", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseDdnsUseOption81(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_ddns_use_option81"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseDdnsUseOption81("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseDdnsUseOption81("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseDenyBootp(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_deny_bootp"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseDenyBootp("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseDenyBootp("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseEmailList(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_email_list"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseEmailList("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_email_list", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseEmailList("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_email_list", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseEnableDdns(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_enable_ddns"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseEnableDdns("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseEnableDdns("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseEnableDhcpThresholds(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_enable_dhcp_thresholds"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseEnableDhcpThresholds("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseEnableDhcpThresholds("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseIgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_ignore_dhcp_option_list_request"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseIgnoreDhcpOptionListRequest("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseIgnoreDhcpOptionListRequest("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseIpamEmailAddresses(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_ipam_email_addresses"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseIpamEmailAddresses("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_email_addresses", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseIpamEmailAddresses("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_email_addresses", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseIpamThresholdSettings(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_ipam_threshold_settings"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseIpamThresholdSettings("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_threshold_settings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseIpamThresholdSettings("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_threshold_settings", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseIpamTrapSettings(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_ipam_trap_settings"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseIpamTrapSettings("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_trap_settings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseIpamTrapSettings("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_trap_settings", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseLeaseScavengeTime(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_lease_scavenge_time"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseLeaseScavengeTime("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseLeaseScavengeTime("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseLogicFilterRules(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_logic_filter_rules"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseLogicFilterRules("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseLogicFilterRules("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseNextserver(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_nextserver"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseNextserver("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseNextserver("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseOptions(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_options"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseOptions("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseOptions("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UsePxeLeaseTime(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_pxe_lease_time"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUsePxeLeaseTime("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUsePxeLeaseTime("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseRecycleLeases(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_recycle_leases"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseRecycleLeases("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseRecycleLeases("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_UseUpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_use_update_dns_on_lease_renewal"
	var v ipam.Networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseUpdateDnsOnLeaseRenewal("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseUpdateDnsOnLeaseRenewal("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckNetworktemplateExists(ctx context.Context, resourceName string, v *ipam.Networktemplate) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.IPAMAPI.
			NetworktemplateAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForNetworktemplate).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetNetworktemplateResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetNetworktemplateResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckNetworktemplateDestroy(ctx context.Context, v *ipam.Networktemplate) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.IPAMAPI.
			NetworktemplateAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForNetworktemplate).
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

func testAccCheckNetworktemplateDisappears(ctx context.Context, v *ipam.Networktemplate) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.IPAMAPI.
			NetworktemplateAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccNetworktemplateImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.Attributes["ref"] == "" {
			return "", fmt.Errorf("ref is not set")
		}
		return rs.Primary.Attributes["ref"], nil
	}
}

func testAccNetworktemplateBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test" {
    name = %q
}
`, name)
}

func testAccNetworktemplateAllowAnyNetmask(name string, allowAnyNetmask string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_allow_any_netmask" {
    name = %q
    allow_any_netmask = %q
}
`, name, allowAnyNetmask)
}

func testAccNetworktemplateAuthority(name string, authority string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_authority" {
    name = %q
    authority = %q
}
`, name, authority)
}

func testAccNetworktemplateAutoCreateReversezone(name string, autoCreateReversezone string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_auto_create_reversezone" {
    name = %q
    auto_create_reversezone = %q
}
`, name, autoCreateReversezone)
}

func testAccNetworktemplateBootfile(name string, bootfile string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_bootfile" {
    name = %q
    bootfile = %q
}
`, name, bootfile)
}

func testAccNetworktemplateBootserver(name string, bootserver string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_bootserver" {
    name = %q
    bootserver = %q
}
`, name, bootserver)
}

func testAccNetworktemplateCloudApiCompatible(name string, cloudApiCompatible string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_cloud_api_compatible" {
    name = %q
    cloud_api_compatible = %q
}
`, name, cloudApiCompatible)
}

func testAccNetworktemplateComment(name string, comment string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccNetworktemplateDdnsDomainname(name string, ddnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_ddns_domainname" {
    name = %q
    ddns_domainname = %q
}
`, name, ddnsDomainname)
}

func testAccNetworktemplateDdnsGenerateHostname(name string, ddnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_ddns_generate_hostname" {
    name = %q
    ddns_generate_hostname = %q
}
`, name, ddnsGenerateHostname)
}

func testAccNetworktemplateDdnsServerAlwaysUpdates(name string, ddnsServerAlwaysUpdates string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_ddns_server_always_updates" {
    name = %q
    ddns_server_always_updates = %q
}
`, name, ddnsServerAlwaysUpdates)
}

func testAccNetworktemplateDdnsTtl(name string, ddnsTtl string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_ddns_ttl" {
    name = %q
    ddns_ttl = %q
}
`, name, ddnsTtl)
}

func testAccNetworktemplateDdnsUpdateFixedAddresses(name string, ddnsUpdateFixedAddresses string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_ddns_update_fixed_addresses" {
    name = %q
    ddns_update_fixed_addresses = %q
}
`, name, ddnsUpdateFixedAddresses)
}

func testAccNetworktemplateDdnsUseOption81(name string, ddnsUseOption81 string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_ddns_use_option81" {
    name = %q
    ddns_use_option81 = %q
}
`, name, ddnsUseOption81)
}

func testAccNetworktemplateDelegatedMember(name string, delegatedMember map[string]string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_delegated_member" {
    name = %q
    delegated_member = %s
}
`, name, delegatedMember)
}

func testAccNetworktemplateDenyBootp(name string, denyBootp string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_deny_bootp" {
    name = %q
    deny_bootp = %q
}
`, name, denyBootp)
}

func testAccNetworktemplateEmailList(name string, emailList []string) string {
	emailListStr := utils.ConvertStringSliceToHCL(emailList)
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_email_list" {
    name = %q
    email_list = %q
}
`, name, emailListStr)
}

func testAccNetworktemplateEnableDdns(name string, enableDdns string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_enable_ddns" {
    name = %q
    enable_ddns = %q
}
`, name, enableDdns)
}

func testAccNetworktemplateEnableDhcpThresholds(name string, enableDhcpThresholds string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_enable_dhcp_thresholds" {
    name = %q
    enable_dhcp_thresholds = %q
}
`, name, enableDhcpThresholds)
}

func testAccNetworktemplateEnableEmailWarnings(name string, enableEmailWarnings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_enable_email_warnings" {
    name = %q
    enable_email_warnings = %q
}
`, name, enableEmailWarnings)
}

func testAccNetworktemplateEnablePxeLeaseTime(name string, enablePxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_enable_pxe_lease_time" {
    name = %q
    enable_pxe_lease_time = %q
}
`, name, enablePxeLeaseTime)
}

func testAccNetworktemplateEnableSnmpWarnings(name string, enableSnmpWarnings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_enable_snmp_warnings" {
    name = %q
    enable_snmp_warnings = %q
}
`, name, enableSnmpWarnings)
}

func testAccNetworktemplateExtAttrs(name string, extAttrs map[string]string) string {
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_extattrs" {
    name = %q
    extattrs = %s
}
`, name, extAttrsStr)
}

func testAccNetworktemplateFixedAddressTemplates(name string, fixedAddressTemplates []string) string {
	fixedAddressTemplatesStr := utils.ConvertStringSliceToHCL(fixedAddressTemplates)
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_fixed_address_templates" {
    name = %q
    fixed_address_templates = %q
}
`, name, fixedAddressTemplatesStr)
}

func testAccNetworktemplateHighWaterMark(name string, highWaterMark string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_high_water_mark" {
    name = %q
    high_water_mark = %q
}
`, name, highWaterMark)
}

func testAccNetworktemplateHighWaterMarkReset(name string, highWaterMarkReset string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_high_water_mark_reset" {
    name = %q
    high_water_mark_reset = %q
}
`, name, highWaterMarkReset)
}

func testAccNetworktemplateIgnoreDhcpOptionListRequest(name string, ignoreDhcpOptionListRequest string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_ignore_dhcp_option_list_request" {
    name = %q
    ignore_dhcp_option_list_request = %q
}
`, name, ignoreDhcpOptionListRequest)
}

func testAccNetworktemplateIpamEmailAddresses(name string, ipamEmailAddresses []string) string {
	ipamEmailAddressesStr := utils.ConvertStringSliceToHCL(ipamEmailAddresses)
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_ipam_email_addresses" {
    name = %q
    ipam_email_addresses = %q
}
`, name, ipamEmailAddressesStr)
}

func testAccNetworktemplateIpamThresholdSettings(name string, ipamThresholdSettings map[string]string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_ipam_threshold_settings" {
    name = %q
    ipam_threshold_settings = %s
}
`, name, ipamThresholdSettings)
}

func testAccNetworktemplateIpamTrapSettings(name string, ipamTrapSettings map[string]string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_ipam_trap_settings" {
    name = %q
    ipam_trap_settings = %s
}
`, name, ipamTrapSettings)
}

func testAccNetworktemplateLeaseScavengeTime(name string, leaseScavengeTime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_lease_scavenge_time" {
    name = %q
    lease_scavenge_time = %q
}
`, name, leaseScavengeTime)
}

func testAccNetworktemplateLogicFilterRules(name string, logicFilterRules []map[string]any) string {
	logicFilterRulesStr := utils.ConvertSliceOfMapsToHCL(logicFilterRules)
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_logic_filter_rules" {
    name = %q
    logic_filter_rules = %s
}
`, name, logicFilterRulesStr)
}

func testAccNetworktemplateLowWaterMark(name string, lowWaterMark string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_low_water_mark" {
    name = %q
    low_water_mark = %q
}
`, name, lowWaterMark)
}

func testAccNetworktemplateLowWaterMarkReset(name string, lowWaterMarkReset string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_low_water_mark_reset" {
    name = %q
    low_water_mark_reset = %q
}
`, name, lowWaterMarkReset)
}

func testAccNetworktemplateMembers(name string, members []map[string]any) string {
	membersStr := utils.ConvertSliceOfMapsToHCL(members)
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_members" {
    name = %q
    members = %s
}
`, name, membersStr)
}

func testAccNetworktemplateName(name string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_name" {
    name = %q
}
`, name)
}

func testAccNetworktemplateNetmask(name string, netmask string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_netmask" {
    name = %q
    netmask = %q
}
`, name, netmask)
}

func testAccNetworktemplateNextserver(name string, nextserver string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_nextserver" {
    name = %q
    nextserver = %q
}
`, name, nextserver)
}

func testAccNetworktemplateOptions(name string, options []map[string]any) string {
	optionsStr := utils.ConvertSliceOfMapsToHCL(options)
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_options" {
    name = %q
    options = %s
}
`, name, optionsStr)
}

func testAccNetworktemplatePxeLeaseTime(name string, pxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_pxe_lease_time" {
    name = %q
    pxe_lease_time = %q
}
`, name, pxeLeaseTime)
}

func testAccNetworktemplateRangeTemplates(name string, rangeTemplates []string) string {
	rangeTemplatesStr := utils.ConvertStringSliceToHCL(rangeTemplates)
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_range_templates" {
    name = %q
    range_templates = %q
}
`, name, rangeTemplatesStr)
}

func testAccNetworktemplateRecycleLeases(name string, recycleLeases string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_recycle_leases" {
    name = %q
    recycle_leases = %q
}
`, name, recycleLeases)
}

func testAccNetworktemplateRirOrganization(name string, rirOrganization string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_rir_organization" {
    name = %q
    rir_organization = %q
}
`, name, rirOrganization)
}

func testAccNetworktemplateRirRegistrationAction(name string, rirRegistrationAction string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_rir_registration_action" {
    name = %q
    rir_registration_action = %q
}
`, name, rirRegistrationAction)
}

func testAccNetworktemplateRirRegistrationStatus(name string, rirRegistrationStatus string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_rir_registration_status" {
    name = %q
    rir_registration_status = %q
}
`, name, rirRegistrationStatus)
}

func testAccNetworktemplateSendRirRequest(name string, sendRirRequest string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_send_rir_request" {
    name = %q
    send_rir_request = %q
}
`, name, sendRirRequest)
}

func testAccNetworktemplateUpdateDnsOnLeaseRenewal(name string, updateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_update_dns_on_lease_renewal" {
    name = %q
    update_dns_on_lease_renewal = %q
}
`, name, updateDnsOnLeaseRenewal)
}

func testAccNetworktemplateUseAuthority(name string, useAuthority string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_authority" {
    name = %q
    use_authority = %q
}
`, name, useAuthority)
}

func testAccNetworktemplateUseBootfile(name string, useBootfile string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_bootfile" {
    name = %q
    use_bootfile = %q
}
`, name, useBootfile)
}

func testAccNetworktemplateUseBootserver(name string, useBootserver string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_bootserver" {
    name = %q
    use_bootserver = %q
}
`, name, useBootserver)
}

func testAccNetworktemplateUseDdnsDomainname(name string, useDdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_ddns_domainname" {
    name = %q
    use_ddns_domainname = %q
}
`, name, useDdnsDomainname)
}

func testAccNetworktemplateUseDdnsGenerateHostname(name string, useDdnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_ddns_generate_hostname" {
    name = %q
    use_ddns_generate_hostname = %q
}
`, name, useDdnsGenerateHostname)
}

func testAccNetworktemplateUseDdnsTtl(name string, useDdnsTtl string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_ddns_ttl" {
    name = %q
    use_ddns_ttl = %q
}
`, name, useDdnsTtl)
}

func testAccNetworktemplateUseDdnsUpdateFixedAddresses(name string, useDdnsUpdateFixedAddresses string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_ddns_update_fixed_addresses" {
    name = %q
    use_ddns_update_fixed_addresses = %q
}
`, name, useDdnsUpdateFixedAddresses)
}

func testAccNetworktemplateUseDdnsUseOption81(name string, useDdnsUseOption81 string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_ddns_use_option81" {
    name = %q
    use_ddns_use_option81 = %q
}
`, name, useDdnsUseOption81)
}

func testAccNetworktemplateUseDenyBootp(name string, useDenyBootp string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_deny_bootp" {
    name = %q
    use_deny_bootp = %q
}
`, name, useDenyBootp)
}

func testAccNetworktemplateUseEmailList(name string, useEmailList string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_email_list" {
    name = %q
    use_email_list = %q
}
`, name, useEmailList)
}

func testAccNetworktemplateUseEnableDdns(name string, useEnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_enable_ddns" {
    name = %q
    use_enable_ddns = %q
}
`, name, useEnableDdns)
}

func testAccNetworktemplateUseEnableDhcpThresholds(name string, useEnableDhcpThresholds string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_enable_dhcp_thresholds" {
    name = %q
    use_enable_dhcp_thresholds = %q
}
`, name, useEnableDhcpThresholds)
}

func testAccNetworktemplateUseIgnoreDhcpOptionListRequest(name string, useIgnoreDhcpOptionListRequest string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_ignore_dhcp_option_list_request" {
    name = %q
    use_ignore_dhcp_option_list_request = %q
}
`, name, useIgnoreDhcpOptionListRequest)
}

func testAccNetworktemplateUseIpamEmailAddresses(name string, useIpamEmailAddresses string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_ipam_email_addresses" {
    name = %q
    use_ipam_email_addresses = %q
}
`, name, useIpamEmailAddresses)
}

func testAccNetworktemplateUseIpamThresholdSettings(name string, useIpamThresholdSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_ipam_threshold_settings" {
    name = %q
    use_ipam_threshold_settings = %q
}
`, name, useIpamThresholdSettings)
}

func testAccNetworktemplateUseIpamTrapSettings(name string, useIpamTrapSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_ipam_trap_settings" {
    name = %q
    use_ipam_trap_settings = %q
}
`, name, useIpamTrapSettings)
}

func testAccNetworktemplateUseLeaseScavengeTime(name string, useLeaseScavengeTime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_lease_scavenge_time" {
    name = %q
    use_lease_scavenge_time = %q
}
`, name, useLeaseScavengeTime)
}

func testAccNetworktemplateUseLogicFilterRules(name string, useLogicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_logic_filter_rules" {
    name = %q
    use_logic_filter_rules = %q
}
`, name, useLogicFilterRules)
}

func testAccNetworktemplateUseNextserver(name string, useNextserver string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_nextserver" {
    name = %q
    use_nextserver = %q
}
`, name, useNextserver)
}

func testAccNetworktemplateUseOptions(name string, useOptions string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_options" {
    name = %q
    use_options = %q
}
`, name, useOptions)
}

func testAccNetworktemplateUsePxeLeaseTime(name string, usePxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_pxe_lease_time" {
    name = %q
    use_pxe_lease_time = %q
}
`, name, usePxeLeaseTime)
}

func testAccNetworktemplateUseRecycleLeases(name string, useRecycleLeases string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_recycle_leases" {
    name = %q
    use_recycle_leases = %q
}
`, name, useRecycleLeases)
}

func testAccNetworktemplateUseUpdateDnsOnLeaseRenewal(name string, useUpdateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_update_dns_on_lease_renewal" {
    name = %q
    use_update_dns_on_lease_renewal = %q
}
`, name, useUpdateDnsOnLeaseRenewal)
}
