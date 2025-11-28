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

var readableAttributesForNetworktemplate = "allow_any_netmask,authority,auto_create_reversezone,bootfile,bootserver,cloud_api_compatible,comment,ddns_domainname,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,ddns_update_fixed_addresses,ddns_use_option81,delegated_member,deny_bootp,email_list,enable_ddns,enable_dhcp_thresholds,enable_email_warnings,enable_pxe_lease_time,enable_snmp_warnings,extattrs,fixed_address_templates,high_water_mark,high_water_mark_reset,ignore_dhcp_option_list_request,ipam_email_addresses,ipam_threshold_settings,ipam_trap_settings,lease_scavenge_time,logic_filter_rules,low_water_mark,low_water_mark_reset,members,name,netmask,nextserver,options,pxe_lease_time,range_templates,recycle_leases,rir,rir_organization,rir_registration_action,rir_registration_status,send_rir_request,update_dns_on_lease_renewal,use_authority,use_bootfile,use_bootserver,use_ddns_domainname,use_ddns_generate_hostname,use_ddns_ttl,use_ddns_update_fixed_addresses,use_ddns_use_option81,use_deny_bootp,use_email_list,use_enable_ddns,use_enable_dhcp_thresholds,use_ignore_dhcp_option_list_request,use_ipam_email_addresses,use_ipam_threshold_settings,use_ipam_trap_settings,use_lease_scavenge_time,use_logic_filter_rules,use_nextserver,use_options,use_pxe_lease_time,use_recycle_leases,use_update_dns_on_lease_renewal"

func TestAccNetworktemplateResource_basic(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateBasicConfig(name, 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "netmask", "24"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "allow_any_netmask", "false"),
					resource.TestCheckResourceAttr(resourceName, "authority", "false"),
					resource.TestCheckResourceAttr(resourceName, "auto_create_reversezone", "false"),
					resource.TestCheckResourceAttr(resourceName, "cloud_api_compatible", "true"),
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
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "-1"),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "0"),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "10"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetworktemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworktemplateBasicConfig(name, 24),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateBasicConfig(name, 24),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateAllowAnyNetmask(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_any_netmask", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateAllowAnyNetmask(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateAuthority(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "authority", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateAuthority(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateAutoCreateReversezone(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_create_reversezone", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateAutoCreateReversezone(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateBootfile(name, 24, "bootfile.txt"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "bootfile.txt"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateBootfile(name, 24, "bootfile_UPDATED.txt"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootfile", "bootfile_UPDATED.txt"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_Bootserver(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_bootserver"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateBootserver(name, 24, "test_bootserver"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "test_bootserver"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateBootserver(name, 24, "updated_test_bootserver"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "bootserver", "updated_test_bootserver"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// The testcase will fail, as this is a known issue
// If the user is a cloud-user, then they need Terraform internal ID with cloud permission and enable cloud delegation for the user to create a range template.
// if the user is a non cloud-user, they need to have  Terraform internal ID without cloud permission.
func TestAccNetworktemplateResource_CloudApiCompatible(t *testing.T) {
	t.Skip("Skipping this test as it is a known issue.")
	var resourceName = "nios_ipam_networktemplate.test_cloud_api_compatible"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateCloudApiCompatible(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_api_compatible", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateCloudApiCompatible(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateComment(name, 24, "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateComment(name, 24, "Updated comment for the object"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateDdnsDomainname(name, 24, "ddns_domain.name"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "ddns_domain.name"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateDdnsDomainname(name, 24, "UPDATED_ddns_domain.name"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "UPDATED_ddns_domain.name"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_DdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_ddns_generate_hostname"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateDdnsGenerateHostname(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateDdnsGenerateHostname(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateDdnsServerAlwaysUpdates(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateDdnsServerAlwaysUpdates(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateDdnsTtl(name, 24, 1000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "1000"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateDdnsTtl(name, 24, 1000000000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "1000000000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworktemplateResource_DdnsUpdateFixedAddresses(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_ddns_update_fixed_addresses"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateDdnsUpdateFixedAddresses(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_update_fixed_addresses", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateDdnsUpdateFixedAddresses(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateDdnsUseOption81(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_option81", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateDdnsUseOption81(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")
	delegatedMemberVal := map[string]any{
		"name":     "infoblox.172_28_82_250",
		"ipv4addr": "172.28.82.250",
	}
	delegatedMemberValUpdated := map[string]any{
		"name":     "infoblox.localdomain",
		"ipv4addr": "1.1.1.1",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateDelegatedMember(name, 24, delegatedMemberVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegated_member.name", "infoblox.172_28_82_250"),
					resource.TestCheckResourceAttr(resourceName, "delegated_member.ipv4addr", "172.28.82.250"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateDelegatedMember(name, 24, delegatedMemberValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegated_member.name", "infoblox.localdomain"),
					resource.TestCheckResourceAttr(resourceName, "delegated_member.ipv4addr", "1.1.1.1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworktemplateResource_DenyBootp(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_deny_bootp"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateDenyBootp(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "deny_bootp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateDenyBootp(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")
	emailListVal := []string{"bbb@info.com", "aaa@wapi.com"}
	emailListValUpdated := []string{"abc@info.com", "xyz@wapi.com"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateEmailList(name, 24, emailListVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_list.0", "bbb@info.com"),
					resource.TestCheckResourceAttr(resourceName, "email_list.1", "aaa@wapi.com"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateEmailList(name, 24, emailListValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "email_list.0", "abc@info.com"),
					resource.TestCheckResourceAttr(resourceName, "email_list.1", "xyz@wapi.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_EnableDdns(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_enable_ddns"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateEnableDdns(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateEnableDdns(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateEnableDhcpThresholds(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_dhcp_thresholds", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateEnableDhcpThresholds(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateEnableEmailWarnings(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_email_warnings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateEnableEmailWarnings(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateEnablePxeLeaseTime(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_pxe_lease_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateEnablePxeLeaseTime(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateEnableSnmpWarnings(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_snmp_warnings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateEnableSnmpWarnings(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateExtAttrs(name, 24, map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateExtAttrs(name, 24, map[string]string{
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
	name := acctest.RandomNameWithPrefix("network-template")
	fixedAddressTemplatesVal := []string{"FIXED_ADDRESS_TEMPLATES_REPLACE_ME1", "FIXED_ADDRESS_TEMPLATES_REPLACE_ME2"}
	fixedAddressTemplatesValUpdated := []string{"FIXED_ADDRESS_TEMPLATES_REPLACE_ME1", "FIXED_ADDRESS_TEMPLATES_REPLACE_ME2"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateFixedAddressTemplates(name, 24, fixedAddressTemplatesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fixed_address_templates", "FIXED_ADDRESS_TEMPLATES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateFixedAddressTemplates(name, 24, fixedAddressTemplatesValUpdated),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateHighWaterMark(name, 24, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "1"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateHighWaterMark(name, 24, 99),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark", "99"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworktemplateResource_HighWaterMarkReset(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_high_water_mark_reset"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateHighWaterMarkReset(name, 24, 50),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "50"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateHighWaterMarkReset(name, 24, 100),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "high_water_mark_reset", "100"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworktemplateResource_IgnoreDhcpOptionListRequest(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_ignore_dhcp_option_list_request"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateIgnoreDhcpOptionListRequest(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ignore_dhcp_option_list_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateIgnoreDhcpOptionListRequest(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")
	ipamEmailAddressesVal := []string{"bbb@info.com", "aaa@wapi.com"}
	ipamEmailAddressesValUpdated := []string{"abc@info.com", "xyz@wapi.com"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateIpamEmailAddresses(name, 24, ipamEmailAddressesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_email_addresses.0", "bbb@info.com"),
					resource.TestCheckResourceAttr(resourceName, "ipam_email_addresses.1", "aaa@wapi.com"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateIpamEmailAddresses(name, 24, ipamEmailAddressesValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_email_addresses.0", "abc@info.com"),
					resource.TestCheckResourceAttr(resourceName, "ipam_email_addresses.1", "xyz@wapi.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_IpamThresholdSettings(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_ipam_threshold_settings"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")
	ipamThresholdSettingsVal := map[string]any{
		"trigger_value": "100",
		"reset_value":   "10",
	}
	ipamThresholdSettingsValUpdated := map[string]any{
		"trigger_value": "9",
		"reset_value":   "1",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateIpamThresholdSettings(name, 24, ipamThresholdSettingsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_threshold_settings", "IPAM_THRESHOLD_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateIpamThresholdSettings(name, 24, ipamThresholdSettingsValUpdated),
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
	name := acctest.RandomNameWithPrefix("network-template")
	ipamTrapSettingsVal := map[string]any{
		"enable_email_warnings": "true",
		"enable_snmp_warnings":  "true",
	}
	ipamTrapSettingsValUpdated := map[string]any{
		"enable_email_warnings": "false",
		"enable_snmp_warnings":  "false",
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateIpamTrapSettings(name, 24, ipamTrapSettingsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_trap_settings.enable_email_warnings", "true"),
					resource.TestCheckResourceAttr(resourceName, "ipam_trap_settings.enable_snmp_warnings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateIpamTrapSettings(name, 24, ipamTrapSettingsValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipam_trap_settings.enable_email_warnings", "false"),
					resource.TestCheckResourceAttr(resourceName, "ipam_trap_settings.enable_snmp_warnings", "false")),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworktemplateResource_LeaseScavengeTime(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_lease_scavenge_time"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateLeaseScavengeTime(name, 24, 2147471999),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "2147471999"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateLeaseScavengeTime(name, 24, 86401),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lease_scavenge_time", "86401"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworktemplateResource_LogicFilterRules(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_logic_filter_rules"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")
	logicFilterRulesVal := []map[string]any{
		{
			"filter": "ipv6_option_filter",
			"type":   "Option",
		},
	}
	logicFilterRulesValUpdated := []map[string]any{
		{
			"filter": "ipv6_option_filter1",
			"type":   "Option",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateLogicFilterRules(name, 24, logicFilterRulesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.filter", "ipv6_option_filter"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.type", "Option"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateLogicFilterRules(name, 24, logicFilterRulesValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.filter", "ipv6_option_filter1"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.type", "Option"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_LowWaterMark(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_low_water_mark"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateLowWaterMark(name, 24, 100),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "100"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateLowWaterMark(name, 24, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworktemplateResource_LowWaterMarkReset(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_low_water_mark_reset"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateLowWaterMarkReset(name, 24, "30"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "30"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateLowWaterMarkReset(name, 24, "0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "low_water_mark_reset", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworktemplateResource_Members(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_members"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")
	membersVal := []map[string]any{}
	membersValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateMembers(name, 24, membersVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "members", "MEMBERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateMembers(name, 24, membersValUpdated),
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
	name := acctest.RandomNameWithPrefix("network-template")
	nameUpdated := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateName(name, 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateName(nameUpdated, 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_Netmask(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_netmask"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateNetmask(name, 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "netmask", "24"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateNetmask(name, 42),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "netmask", "42"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworktemplateResource_Nextserver(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_nextserver"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateNextserver(name, 24, "NEXTSERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nextserver", "NEXTSERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateNextserver(name, 24, "NEXTSERVER_UPDATE_REPLACE_ME"),
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
	name := acctest.RandomNameWithPrefix("network-template")
	optionsVal := []map[string]any{
		{
			"name":  "time-offset",
			"num":   2,
			"value": "50",
		},
		{
			"name":  "subnet-mask",
			"value": "1.1.1.1",
		},
	}

	optionsValUpdated := []map[string]any{
		{
			"num":   51,
			"value": "7200",
		},
		{
			"name":  "subnet-mask",
			"value": "1.1.1.1",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateOptions(name, 24, optionsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "time-offset"),
					resource.TestCheckResourceAttr(resourceName, "options.0.num", "2"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "50"),
					resource.TestCheckResourceAttr(resourceName, "options.1.name", "subnet-mask"),
					resource.TestCheckResourceAttr(resourceName, "options.1.value", "1.1.1.1")),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateOptions(name, 24, optionsValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "dhcp-lease-time"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "7200"),
					resource.TestCheckResourceAttr(resourceName, "options.1.name", "subnet-mask"),
					resource.TestCheckResourceAttr(resourceName, "options.1.value", "1.1.1.1")),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworktemplateResource_PxeLeaseTime(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_pxe_lease_time"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplatePxeLeaseTime(name, 24, "1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "1"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplatePxeLeaseTime(name, 24, "1000"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "pxe_lease_time", "1000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccNetworktemplateResource_RangeTemplates(t *testing.T) {
	var resourceName = "nios_ipam_networktemplate.test_range_templates"
	var v ipam.Networktemplate
	name := acctest.RandomNameWithPrefix("network-template")
	rangeTemplatesVal := []string{"RANGE_TEMPLATES_REPLACE_ME1", "RANGE_TEMPLATES_REPLACE_ME2"}
	rangeTemplatesValUpdated := []string{"RANGE_TEMPLATES_REPLACE_ME1", "RANGE_TEMPLATES_REPLACE_ME2"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateRangeTemplates(name, 24, rangeTemplatesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "range_templates", "RANGE_TEMPLATES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateRangeTemplates(name, 24, rangeTemplatesValUpdated),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateRecycleLeases(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateRecycleLeases(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateRirOrganization(name, 24, "RIR_ORGANIZATION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_organization", "RIR_ORGANIZATION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateRirOrganization(name, 24, "RIR_ORGANIZATION_UPDATE_REPLACE_ME"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateRirRegistrationAction(name, 24, "CREATE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_action", "CREATE"),
				),
			},
			{
				Config: testAccNetworktemplateRirRegistrationAction(name, 24, "NONE"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateRirRegistrationStatus(name, 24, "NOT_REGISTERED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_status", "NOT_REGISTERED"),
				),
			},
			{
				Config: testAccNetworktemplateRirRegistrationStatus(name, 24, "REGISTERED"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateSendRirRequest(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "send_rir_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateSendRirRequest(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUpdateDnsOnLeaseRenewal(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUpdateDnsOnLeaseRenewal(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseAuthority(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_authority", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseAuthority(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseBootfile(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootfile", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseBootfile(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseBootserver(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_bootserver", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseBootserver(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseDdnsDomainname(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseDdnsDomainname(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseDdnsGenerateHostname(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseDdnsGenerateHostname(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseDdnsTtl(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseDdnsTtl(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseDdnsUpdateFixedAddresses(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_update_fixed_addresses", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseDdnsUpdateFixedAddresses(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseDdnsUseOption81(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_use_option81", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseDdnsUseOption81(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseDenyBootp(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_deny_bootp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseDenyBootp(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseEmailList(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_email_list", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseEmailList(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseEnableDdns(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseEnableDdns(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseEnableDhcpThresholds(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_dhcp_thresholds", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseEnableDhcpThresholds(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseIgnoreDhcpOptionListRequest(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ignore_dhcp_option_list_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseIgnoreDhcpOptionListRequest(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseIpamEmailAddresses(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_email_addresses", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseIpamEmailAddresses(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseIpamThresholdSettings(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_threshold_settings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseIpamThresholdSettings(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseIpamTrapSettings(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ipam_trap_settings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseIpamTrapSettings(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseLeaseScavengeTime(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_lease_scavenge_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseLeaseScavengeTime(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseLogicFilterRules(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseLogicFilterRules(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseNextserver(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nextserver", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseNextserver(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseOptions(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseOptions(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUsePxeLeaseTime(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_pxe_lease_time", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUsePxeLeaseTime(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseRecycleLeases(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseRecycleLeases(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworktemplateUseUpdateDnsOnLeaseRenewal(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworktemplateUseUpdateDnsOnLeaseRenewal(name, 24, "false"),
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

func testAccNetworktemplateBasicConfig(name string, netmask int) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test" {
    name = %q
    netmask = %d
}
`, name, netmask)
}

func testAccNetworktemplateAllowAnyNetmask(name string, netmask int, allowAnyNetmask string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_allow_any_netmask" {
    name = %q
    netmask = %d
    allow_any_netmask = %q
}
`, name, netmask, allowAnyNetmask)
}

func testAccNetworktemplateAuthority(name string, netmask int, authority string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_authority" {
    name = %q
    netmask = %d
    authority = %q
    use_authority = true
}
`, name, netmask, authority)
}

func testAccNetworktemplateAutoCreateReversezone(name string, netmask int, autoCreateReversezone string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_auto_create_reversezone" {
    name = %q
    netmask = %d
    auto_create_reversezone = %q
}
`, name, netmask, autoCreateReversezone)
}

func testAccNetworktemplateBootfile(name string, netmask int, bootfile string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_bootfile" {
    name = %q
    netmask = %d
    bootfile = %q
    use_bootfile = true
}
`, name, netmask, bootfile)
}

func testAccNetworktemplateBootserver(name string, netmask int, bootserver string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_bootserver" {
    name = %q
    netmask = %d
    bootserver = %q
    use_bootserver = true
}
`, name, netmask, bootserver)
}

func testAccNetworktemplateCloudApiCompatible(name string, netmask int, cloudApiCompatible string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_cloud_api_compatible" {
    name = %q
    netmask = %d
    cloud_api_compatible = %q
}
`, name, netmask, cloudApiCompatible)
}

func testAccNetworktemplateComment(name string, netmask int, comment string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_comment" {
    name = %q
    netmask = %d
    comment = %q
}
`, name, netmask, comment)
}

func testAccNetworktemplateDdnsDomainname(name string, netmask int, ddnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_ddns_domainname" {
    name = %q
    netmask = %d
    ddns_domainname = %q
    use_ddns_domainname = true
}
`, name, netmask, ddnsDomainname)
}

func testAccNetworktemplateDdnsGenerateHostname(name string, netmask int, ddnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_ddns_generate_hostname" {
    name = %q
    netmask = %d
    ddns_generate_hostname = %q
    use_ddns_generate_hostname = true
}
`, name, netmask, ddnsGenerateHostname)
}

func testAccNetworktemplateDdnsServerAlwaysUpdates(name string, netmask int, ddnsServerAlwaysUpdates string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_ddns_server_always_updates" {
    name = %q
    netmask = %d
    ddns_server_always_updates = %q
    ddns_use_option81 = true
    use_ddns_use_option81 = true
}
`, name, netmask, ddnsServerAlwaysUpdates)
}

func testAccNetworktemplateDdnsTtl(name string, netmask int, ddnsTtl int) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_ddns_ttl" {
    name = %q
    netmask = %d
    ddns_ttl = %d
    use_ddns_ttl = true
}
`, name, netmask, ddnsTtl)
}

func testAccNetworktemplateDdnsUpdateFixedAddresses(name string, netmask int, ddnsUpdateFixedAddresses string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_ddns_update_fixed_addresses" {
    name = %q
    netmask = %d
    ddns_update_fixed_addresses = %q
    use_ddns_update_fixed_addresses = true
}
`, name, netmask, ddnsUpdateFixedAddresses)
}

func testAccNetworktemplateDdnsUseOption81(name string, netmask int, ddnsUseOption81 string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_ddns_use_option81" {
    name = %q
    netmask = %d
    ddns_use_option81 = %q
    use_ddns_use_option81 = true
}
`, name, netmask, ddnsUseOption81)
}

func testAccNetworktemplateDelegatedMember(name string, netmask int, delegatedMember map[string]any) string {
	delegatedMemberStr := utils.ConvertMapToHCL(delegatedMember)
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_delegated_member" {
    name = %q
    netmask = %d
    delegated_member = %s
}
`, name, netmask, delegatedMemberStr)
}

func testAccNetworktemplateDenyBootp(name string, netmask int, denyBootp string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_deny_bootp" {
    name = %q
    netmask = %d
    deny_bootp = %q
    use_deny_bootp = true
}
`, name, netmask, denyBootp)
}

func testAccNetworktemplateEmailList(name string, netmask int, emailList []string) string {
	emailListStr := utils.ConvertStringSliceToHCL(emailList)
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_email_list" {
    name = %q
    netmask = %d
    email_list = %s
    use_email_list = true
}
`, name, netmask, emailListStr)
}

func testAccNetworktemplateEnableDdns(name string, netmask int, enableDdns string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_enable_ddns" {
    name = %q
    netmask = %d
    enable_ddns = %q
    use_enable_ddns = true
}
`, name, netmask, enableDdns)
}

func testAccNetworktemplateEnableDhcpThresholds(name string, netmask int, enableDhcpThresholds string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_enable_dhcp_thresholds" {
    name = %q
    netmask = %d
    enable_dhcp_thresholds = %q
    use_enable_dhcp_thresholds = true
}
`, name, netmask, enableDhcpThresholds)
}

func testAccNetworktemplateEnableEmailWarnings(name string, netmask int, enableEmailWarnings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_enable_email_warnings" {
    name = %q
    netmask = %d
    enable_email_warnings = %q
}
`, name, netmask, enableEmailWarnings)
}

func testAccNetworktemplateEnablePxeLeaseTime(name string, netmask int, enablePxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_enable_pxe_lease_time" {
    name = %q
    netmask = %d
    enable_pxe_lease_time = %q
    pxe_lease_time = 600
    use_pxe_lease_time = true
}
`, name, netmask, enablePxeLeaseTime)
}

func testAccNetworktemplateEnableSnmpWarnings(name string, netmask int, enableSnmpWarnings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_enable_snmp_warnings" {
    name = %q
    netmask = %d
    enable_snmp_warnings = %q
}
`, name, netmask, enableSnmpWarnings)
}

func testAccNetworktemplateExtAttrs(name string, netmask int, extAttrs map[string]string) string {
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_extattrs" {
    name = %q
    netmask = %d
    extattrs = %s
}
`, name, netmask, extAttrsStr)
}

func testAccNetworktemplateFixedAddressTemplates(name string, netmask int, fixedAddressTemplates []string) string {
	fixedAddressTemplatesStr := utils.ConvertStringSliceToHCL(fixedAddressTemplates)
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_fixed_address_templates" {
    name = %q
    netmask = %d
    fixed_address_templates = %s
}
`, name, netmask, fixedAddressTemplatesStr)
}

func testAccNetworktemplateHighWaterMark(name string, netmask int, highWaterMark int) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_high_water_mark" {
    name = %q
    netmask = %d
    high_water_mark = %d
}
`, name, netmask, highWaterMark)
}

func testAccNetworktemplateHighWaterMarkReset(name string, netmask int, highWaterMarkReset int) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_high_water_mark_reset" {
    name = %q
    netmask = %d
    high_water_mark_reset = %d
}
`, name, netmask, highWaterMarkReset)
}

func testAccNetworktemplateIgnoreDhcpOptionListRequest(name string, netmask int, ignoreDhcpOptionListRequest string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_ignore_dhcp_option_list_request" {
    name = %q
    netmask = %d
    ignore_dhcp_option_list_request = %q
    use_ignore_dhcp_option_list_request = true
}
`, name, netmask, ignoreDhcpOptionListRequest)
}

func testAccNetworktemplateIpamEmailAddresses(name string, netmask int, ipamEmailAddresses []string) string {
	ipamEmailAddressesStr := utils.ConvertStringSliceToHCL(ipamEmailAddresses)
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_ipam_email_addresses" {
    name = %q
    netmask = %d
    ipam_email_addresses = %s
    use_ipam_email_addresses = true
}
`, name, netmask, ipamEmailAddressesStr)
}

func testAccNetworktemplateIpamThresholdSettings(name string, netmask int, ipamThresholdSettings map[string]any) string {
	ipamThresholdSettingsStr := utils.ConvertMapToHCL(ipamThresholdSettings)
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_ipam_threshold_settings" {
    name = %q
    netmask = %d
    ipam_threshold_settings = %s
    use_ipam_threshold_settings = true
}
`, name, netmask, ipamThresholdSettingsStr)
}

func testAccNetworktemplateIpamTrapSettings(name string, netmask int, ipamTrapSettings map[string]any) string {
	ipamTrapSettingsStr := utils.ConvertMapToHCL(ipamTrapSettings)
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_ipam_trap_settings" {
    name = %q
    netmask = %d
    ipam_trap_settings = %s
    use_ipam_trap_settings = true
}
`, name, netmask, ipamTrapSettingsStr)
}

func testAccNetworktemplateLeaseScavengeTime(name string, netmask int, leaseScavengeTime int) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_lease_scavenge_time" {
    name = %q
    netmask = %d
    lease_scavenge_time = %d
    use_lease_scavenge_time = true
}
`, name, netmask, leaseScavengeTime)
}

func testAccNetworktemplateLogicFilterRules(name string, netmask int, logicFilterRules []map[string]any) string {
	logicFilterRulesStr := utils.ConvertSliceOfMapsToHCL(logicFilterRules)
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_logic_filter_rules" {
    name = %q
    netmask = %d
    logic_filter_rules = %s
    use_logic_filter_rules = true
}
`, name, netmask, logicFilterRulesStr)
}

func testAccNetworktemplateLowWaterMark(name string, netmask int, lowWaterMark int) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_low_water_mark" {
    name = %q
    netmask = %d
    low_water_mark = %d
}
`, name, netmask, lowWaterMark)
}

func testAccNetworktemplateLowWaterMarkReset(name string, netmask int, lowWaterMarkReset string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_low_water_mark_reset" {
    name = %q
    netmask = %d
    low_water_mark_reset = %q
}
`, name, netmask, lowWaterMarkReset)
}

func testAccNetworktemplateMembers(name string, netmask int, members []map[string]any) string {
	membersStr := utils.ConvertSliceOfMapsToHCL(members)
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_members" {
    name = %q
    netmask = %d
    members = %s
}
`, name, netmask, membersStr)
}

func testAccNetworktemplateName(name string, netmask int) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_name" {
    name = %q
    netmask = %d
}
`, name, netmask)
}

func testAccNetworktemplateNetmask(name string, netmask int) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_netmask" {
    name = %q
    netmask = %d
}
`, name, netmask)
}

func testAccNetworktemplateNextserver(name string, netmask int, nextserver string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_nextserver" {
    name = %q
    netmask = %d
    nextserver = %q
    use_nextserver = true
}
`, name, netmask, nextserver)
}

func testAccNetworktemplateOptions(name string, netmask int, options []map[string]any) string {
	optionsStr := utils.ConvertSliceOfMapsToHCL(options)
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_options" {
    name = %q
    netmask = %d
    options = %s
    use_options = true
}
`, name, netmask, optionsStr)
}

func testAccNetworktemplatePxeLeaseTime(name string, netmask int, pxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_pxe_lease_time" {
    name = %q
    netmask = %d
    pxe_lease_time = %q
    use_pxe_lease_time = true
}
`, name, netmask, pxeLeaseTime)
}

func testAccNetworktemplateRangeTemplates(name string, netmask int, rangeTemplates []string) string {
	rangeTemplatesStr := utils.ConvertStringSliceToHCL(rangeTemplates)
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_range_templates" {
    name = %q
    netmask = %d
    range_templates = %q
}
`, name, netmask, rangeTemplatesStr)
}

func testAccNetworktemplateRecycleLeases(name string, netmask int, recycleLeases string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_recycle_leases" {
    name = %q
    netmask = %d
    recycle_leases = %q
    use_recycle_leases = true
}
`, name, netmask, recycleLeases)
}

func testAccNetworktemplateRirOrganization(name string, netmask int, rirOrganization string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_rir_organization" {
    name = %q
    netmask = %d
    rir_organization = %q
}
`, name, netmask, rirOrganization)
}

func testAccNetworktemplateRirRegistrationAction(name string, netmask int, rirRegistrationAction string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_rir_registration_action" {
    name = %q
    netmask = %d
    rir_registration_action = %q
    rir_organization = "Test Organization"
}
`, name, netmask, rirRegistrationAction)
}

func testAccNetworktemplateRirRegistrationStatus(name string, netmask int, rirRegistrationStatus string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_rir_registration_status" {
    name = %q
    netmask = %d
    rir_registration_status = %q
}
`, name, netmask, rirRegistrationStatus)
}

func testAccNetworktemplateSendRirRequest(name string, netmask int, sendRirRequest string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_send_rir_request" {
    name = %q
    netmask = %d
    send_rir_request = %q
}
`, name, netmask, sendRirRequest)
}

func testAccNetworktemplateUpdateDnsOnLeaseRenewal(name string, netmask int, updateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_update_dns_on_lease_renewal" {
    name = %q
    netmask = %d
    update_dns_on_lease_renewal = %q
    use_update_dns_on_lease_renewal = true
}
`, name, netmask, updateDnsOnLeaseRenewal)
}

func testAccNetworktemplateUseAuthority(name string, netmask int, useAuthority string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_authority" {
    name = %q
    netmask = %d
    use_authority = %q
}
`, name, netmask, useAuthority)
}

func testAccNetworktemplateUseBootfile(name string, netmask int, useBootfile string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_bootfile" {
    name = %q
    netmask = %d
    use_bootfile = %q
}
`, name, netmask, useBootfile)
}

func testAccNetworktemplateUseBootserver(name string, netmask int, useBootserver string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_bootserver" {
    name = %q
    netmask = %d
    use_bootserver = %q
}
`, name, netmask, useBootserver)
}

func testAccNetworktemplateUseDdnsDomainname(name string, netmask int, useDdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_ddns_domainname" {
    name = %q
    netmask = %d
    use_ddns_domainname = %q
}
`, name, netmask, useDdnsDomainname)
}

func testAccNetworktemplateUseDdnsGenerateHostname(name string, netmask int, useDdnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_ddns_generate_hostname" {
    name = %q
    netmask = %d
    use_ddns_generate_hostname = %q
}
`, name, netmask, useDdnsGenerateHostname)
}

func testAccNetworktemplateUseDdnsTtl(name string, netmask int, useDdnsTtl string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_ddns_ttl" {
    name = %q
    netmask = %d
    use_ddns_ttl = %q
}
`, name, netmask, useDdnsTtl)
}

func testAccNetworktemplateUseDdnsUpdateFixedAddresses(name string, netmask int, useDdnsUpdateFixedAddresses string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_ddns_update_fixed_addresses" {
    name = %q
    netmask = %d
    use_ddns_update_fixed_addresses = %q
}
`, name, netmask, useDdnsUpdateFixedAddresses)
}

func testAccNetworktemplateUseDdnsUseOption81(name string, netmask int, useDdnsUseOption81 string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_ddns_use_option81" {
    name = %q
    netmask = %d
    use_ddns_use_option81 = %q
}
`, name, netmask, useDdnsUseOption81)
}

func testAccNetworktemplateUseDenyBootp(name string, netmask int, useDenyBootp string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_deny_bootp" {
    name = %q
    netmask = %d
    use_deny_bootp = %q
}
`, name, netmask, useDenyBootp)
}

func testAccNetworktemplateUseEmailList(name string, netmask int, useEmailList string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_email_list" {
    name = %q
    netmask = %d
    use_email_list = %q
}
`, name, netmask, useEmailList)
}

func testAccNetworktemplateUseEnableDdns(name string, netmask int, useEnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_enable_ddns" {
    name = %q
    netmask = %d
    use_enable_ddns = %q
}
`, name, netmask, useEnableDdns)
}

func testAccNetworktemplateUseEnableDhcpThresholds(name string, netmask int, useEnableDhcpThresholds string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_enable_dhcp_thresholds" {
    name = %q
    netmask = %d
    use_enable_dhcp_thresholds = %q
}
`, name, netmask, useEnableDhcpThresholds)
}

func testAccNetworktemplateUseIgnoreDhcpOptionListRequest(name string, netmask int, useIgnoreDhcpOptionListRequest string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_ignore_dhcp_option_list_request" {
    name = %q
    netmask = %d
    use_ignore_dhcp_option_list_request = %q
}
`, name, netmask, useIgnoreDhcpOptionListRequest)
}

func testAccNetworktemplateUseIpamEmailAddresses(name string, netmask int, useIpamEmailAddresses string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_ipam_email_addresses" {
    name = %q
    netmask = %d
    use_ipam_email_addresses = %q
}
`, name, netmask, useIpamEmailAddresses)
}

func testAccNetworktemplateUseIpamThresholdSettings(name string, netmask int, useIpamThresholdSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_ipam_threshold_settings" {
    name = %q
    netmask = %d
    use_ipam_threshold_settings = %q
}
`, name, netmask, useIpamThresholdSettings)
}

func testAccNetworktemplateUseIpamTrapSettings(name string, netmask int, useIpamTrapSettings string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_ipam_trap_settings" {
    name = %q
    netmask = %d
    use_ipam_trap_settings = %q
}
`, name, netmask, useIpamTrapSettings)
}

func testAccNetworktemplateUseLeaseScavengeTime(name string, netmask int, useLeaseScavengeTime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_lease_scavenge_time" {
    name = %q
    netmask = %d
    use_lease_scavenge_time = %q
}
`, name, netmask, useLeaseScavengeTime)
}

func testAccNetworktemplateUseLogicFilterRules(name string, netmask int, useLogicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_logic_filter_rules" {
    name = %q
    netmask = %d
    use_logic_filter_rules = %q
}
`, name, netmask, useLogicFilterRules)
}

func testAccNetworktemplateUseNextserver(name string, netmask int, useNextserver string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_nextserver" {
    name = %q
    netmask = %d
    use_nextserver = %q
}
`, name, netmask, useNextserver)
}

func testAccNetworktemplateUseOptions(name string, netmask int, useOptions string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_options" {
    name = %q
    netmask = %d
    use_options = %q
}
`, name, netmask, useOptions)
}

func testAccNetworktemplateUsePxeLeaseTime(name string, netmask int, usePxeLeaseTime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_pxe_lease_time" {
    name = %q
    netmask = %d
    use_pxe_lease_time = %q
}
`, name, netmask, usePxeLeaseTime)
}

func testAccNetworktemplateUseRecycleLeases(name string, netmask int, useRecycleLeases string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_recycle_leases" {
    name = %q
    netmask = %d
    use_recycle_leases = %q
}
`, name, netmask, useRecycleLeases)
}

func testAccNetworktemplateUseUpdateDnsOnLeaseRenewal(name string, netmask int, useUpdateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networktemplate" "test_use_update_dns_on_lease_renewal" {
    name = %q
    netmask = %d
    use_update_dns_on_lease_renewal = %q
}
`, name, netmask, useUpdateDnsOnLeaseRenewal)
}
