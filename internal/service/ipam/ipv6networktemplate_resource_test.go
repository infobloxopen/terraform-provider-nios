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
// Manage ipam Ipv6networktemplate with Basic Fields
resource "nios_ipam_ipv6networktemplate" "ipam_ipv6networktemplate_basic" {
    name = "NAME_REPLACE_ME"
}

// Manage ipam Ipv6networktemplate with Additional Fields
resource "nios_ipam_ipv6networktemplate" "ipam_ipv6networktemplate_with_additional_fields" {
    name = "NAME_REPLACE_ME"

// TODO : Add additional optional fields below

    //Extensible Attributes
    extattrs = {
        Site = "location-1"
    }
}
*/

var readableAttributesForIpv6networktemplate = "allow_any_netmask,auto_create_reversezone,cidr,cloud_api_compatible,comment,ddns_domainname,ddns_enable_option_fqdn,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,delegated_member,domain_name,domain_name_servers,enable_ddns,extattrs,fixed_address_templates,ipv6prefix,logic_filter_rules,members,name,options,preferred_lifetime,range_templates,recycle_leases,rir,rir_organization,rir_registration_action,rir_registration_status,send_rir_request,update_dns_on_lease_renewal,use_ddns_domainname,use_ddns_enable_option_fqdn,use_ddns_generate_hostname,use_ddns_ttl,use_domain_name,use_domain_name_servers,use_enable_ddns,use_logic_filter_rules,use_options,use_preferred_lifetime,use_recycle_leases,use_update_dns_on_lease_renewal,use_valid_lifetime,valid_lifetime"

func TestAccIpv6networktemplateResource_basic(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateBasicConfig("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "allow_any_netmask", "false"),
					resource.TestCheckResourceAttr(resourceName, "auto_create_reversezone", "false"),
					resource.TestCheckResourceAttr(resourceName, "cloud_api_compatible", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_enable_option_fqdn", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "true"),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "0"),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
					// TODO : Add validation for default value for field logic_filter_rules if applicable
					// TODO : Add validation for default value for field members if applicable
					// TODO : Add validation for default value for field options if applicable
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "27000"),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "true"),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_action", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_status", "NOT_REGISTERED"),
					resource.TestCheckResourceAttr(resourceName, "send_rir_request", "false"),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_enable_option_fqdn", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "false"),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "43200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_disappears(t *testing.T) {
	resourceName := "nios_ipam_ipv6networktemplate.test"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6networktemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6networktemplateBasicConfig("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					testAccCheckIpv6networktemplateDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccIpv6networktemplateResource_Import(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateBasicConfig("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
				),
			},
			// Import with PlanOnly to detect differences
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccIpv6networktemplateImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "ref",
				PlanOnly:                             true,
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccIpv6networktemplateImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_AllowAnyNetmask(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_allow_any_netmask"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateAllowAnyNetmask("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_any_netmask", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateAllowAnyNetmask("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_any_netmask", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_AutoCreateReversezone(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_auto_create_reversezone"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateAutoCreateReversezone("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_create_reversezone", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateAutoCreateReversezone("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_create_reversezone", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_Cidr(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_cidr"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateCidr("NAME_REPLACE_ME", "CIDR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cidr", "CIDR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateCidr("NAME_REPLACE_ME", "CIDR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cidr", "CIDR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccIpv6networktemplateResource_CloudApiCompatible(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_cloud_api_compatible"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateCloudApiCompatible("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_api_compatible", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateCloudApiCompatible("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_api_compatible", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_Comment(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_comment"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateComment("NAME_REPLACE_ME", "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateComment("NAME_REPLACE_ME", "Updated comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated comment for the object"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_DdnsDomainname(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_ddns_domainname"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateDdnsDomainname("NAME_REPLACE_ME", "DDNS_DOMAINNAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "DDNS_DOMAINNAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateDdnsDomainname("NAME_REPLACE_ME", "DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "DDNS_DOMAINNAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_DdnsEnableOptionFqdn(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_ddns_enable_option_fqdn"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateDdnsEnableOptionFqdn("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_enable_option_fqdn", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateDdnsEnableOptionFqdn("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_enable_option_fqdn", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_DdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_ddns_generate_hostname"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateDdnsGenerateHostname("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateDdnsGenerateHostname("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_DdnsServerAlwaysUpdates(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_ddns_server_always_updates"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateDdnsServerAlwaysUpdates("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateDdnsServerAlwaysUpdates("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_DdnsTtl(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_ddns_ttl"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateDdnsTtl("NAME_REPLACE_ME", "DDNS_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "DDNS_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateDdnsTtl("NAME_REPLACE_ME", "DDNS_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "DDNS_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccIpv6networktemplateResource_DelegatedMember(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_delegated_member"
	var v ipam.Ipv6networktemplate
	delegatedMemberVal := map[string]any{}
	delegatedMemberValUpdated := map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateDelegatedMember("NAME_REPLACE_ME", delegatedMemberVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegated_member", "DELEGATED_MEMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateDelegatedMember("NAME_REPLACE_ME", delegatedMemberValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegated_member", "DELEGATED_MEMBER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccIpv6networktemplateResource_DomainName(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_domain_name"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateDomainName("NAME_REPLACE_ME", "DOMAIN_NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name", "DOMAIN_NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateDomainName("NAME_REPLACE_ME", "DOMAIN_NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name", "DOMAIN_NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_DomainNameServers(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_domain_name_servers"
	var v ipam.Ipv6networktemplate
	domainNameServersVal := []string{"DOMAIN_NAME_SERVERS_REPLACE_ME1", "DOMAIN_NAME_SERVERS_REPLACE_ME2"}
	domainNameServersValUpdated := []string{"DOMAIN_NAME_SERVERS_REPLACE_ME1", "DOMAIN_NAME_SERVERS_REPLACE_ME2"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateDomainNameServers("NAME_REPLACE_ME", domainNameServersVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers", "DOMAIN_NAME_SERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateDomainNameServers("NAME_REPLACE_ME", domainNameServersValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers", "DOMAIN_NAME_SERVERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_EnableDdns(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_enable_ddns"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateEnableDdns("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateEnableDdns("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_extattrs"
	var v ipam.Ipv6networktemplate
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateExtAttrs("NAME_REPLACE_ME", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateExtAttrs("NAME_REPLACE_ME", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_FixedAddressTemplates(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_fixed_address_templates"
	var v ipam.Ipv6networktemplate
	fixedAddressTemplatesVal := []string{"FIXED_ADDRESS_TEMPLATES_REPLACE_ME1", "FIXED_ADDRESS_TEMPLATES_REPLACE_ME2"}
	fixedAddressTemplatesValUpdated := []string{"FIXED_ADDRESS_TEMPLATES_REPLACE_ME1", "FIXED_ADDRESS_TEMPLATES_REPLACE_ME2"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateFixedAddressTemplates("NAME_REPLACE_ME", fixedAddressTemplatesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fixed_address_templates", "FIXED_ADDRESS_TEMPLATES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateFixedAddressTemplates("NAME_REPLACE_ME", fixedAddressTemplatesValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fixed_address_templates", "FIXED_ADDRESS_TEMPLATES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_Ipv6prefix(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_ipv6prefix"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateIpv6prefix("NAME_REPLACE_ME", "IPV6PREFIX_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6prefix", "IPV6PREFIX_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateIpv6prefix("NAME_REPLACE_ME", "IPV6PREFIX_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6prefix", "IPV6PREFIX_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_LogicFilterRules(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_logic_filter_rules"
	var v ipam.Ipv6networktemplate
	logicFilterRulesVal := []map[string]any{}
	logicFilterRulesValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateLogicFilterRules("NAME_REPLACE_ME", logicFilterRulesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateLogicFilterRules("NAME_REPLACE_ME", logicFilterRulesValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules", "LOGIC_FILTER_RULES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_Members(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_members"
	var v ipam.Ipv6networktemplate
	membersVal := []map[string]any{}
	membersValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateMembers("NAME_REPLACE_ME", membersVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "members", "MEMBERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateMembers("NAME_REPLACE_ME", membersValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "members", "MEMBERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_Name(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_name"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_Options(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_options"
	var v ipam.Ipv6networktemplate
	optionsVal := []map[string]any{}
	optionsValUpdated := []map[string]any{}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateOptions("NAME_REPLACE_ME", optionsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateOptions("NAME_REPLACE_ME", optionsValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options", "OPTIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_PreferredLifetime(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_preferred_lifetime"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplatePreferredLifetime("NAME_REPLACE_ME", "PREFERRED_LIFETIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "PREFERRED_LIFETIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplatePreferredLifetime("NAME_REPLACE_ME", "PREFERRED_LIFETIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "PREFERRED_LIFETIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccIpv6networktemplateResource_RangeTemplates(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_range_templates"
	var v ipam.Ipv6networktemplate
	rangeTemplatesVal := []string{"RANGE_TEMPLATES_REPLACE_ME1", "RANGE_TEMPLATES_REPLACE_ME2"}
	rangeTemplatesValUpdated := []string{"RANGE_TEMPLATES_REPLACE_ME1", "RANGE_TEMPLATES_REPLACE_ME2"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateRangeTemplates("NAME_REPLACE_ME", rangeTemplatesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "range_templates", "RANGE_TEMPLATES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateRangeTemplates("NAME_REPLACE_ME", rangeTemplatesValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "range_templates", "RANGE_TEMPLATES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_RecycleLeases(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_recycle_leases"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateRecycleLeases("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateRecycleLeases("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_RirOrganization(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_rir_organization"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateRirOrganization("NAME_REPLACE_ME", "RIR_ORGANIZATION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_organization", "RIR_ORGANIZATION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateRirOrganization("NAME_REPLACE_ME", "RIR_ORGANIZATION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_organization", "RIR_ORGANIZATION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_RirRegistrationAction(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_rir_registration_action"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateRirRegistrationAction("RIR_REGISTRATION_ACTION_REPLACE_ME", "CREATE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_action", "CREATE"),
				),
			},
			{
				Config: testAccIpv6networktemplateRirRegistrationAction("RIR_REGISTRATION_ACTION_REPLACE_ME", "NONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_action", "NONE"),
				),
			},
		},
	})
}

func TestAccIpv6networktemplateResource_RirRegistrationStatus(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_rir_registration_status"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateRirRegistrationStatus("RIR_REGISTRATION_STATUS_REPLACE_ME", "NOT_REGISTERED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_status", "NOT_REGISTERED"),
				),
			},
			{
				Config: testAccIpv6networktemplateRirRegistrationStatus("RIR_REGISTRATION_STATUS_REPLACE_ME", "REGISTERED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_status", "REGISTERED"),
				),
			},
		},
	})
}

func TestAccIpv6networktemplateResource_SendRirRequest(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_send_rir_request"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateSendRirRequest("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "send_rir_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateSendRirRequest("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "send_rir_request", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_UpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_update_dns_on_lease_renewal"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUpdateDnsOnLeaseRenewal("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUpdateDnsOnLeaseRenewal("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_UseDdnsDomainname(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_use_ddns_domainname"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseDdnsDomainname("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseDdnsDomainname("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_UseDdnsEnableOptionFqdn(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_use_ddns_enable_option_fqdn"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseDdnsEnableOptionFqdn("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_enable_option_fqdn", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseDdnsEnableOptionFqdn("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_enable_option_fqdn", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_UseDdnsGenerateHostname(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_use_ddns_generate_hostname"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseDdnsGenerateHostname("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseDdnsGenerateHostname("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_UseDdnsTtl(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_use_ddns_ttl"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseDdnsTtl("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseDdnsTtl("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_UseDomainName(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_use_domain_name"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseDomainName("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseDomainName("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_UseDomainNameServers(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_use_domain_name_servers"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseDomainNameServers("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseDomainNameServers("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_UseEnableDdns(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_use_enable_ddns"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseEnableDdns("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseEnableDdns("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_UseLogicFilterRules(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_use_logic_filter_rules"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseLogicFilterRules("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseLogicFilterRules("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_UseOptions(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_use_options"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseOptions("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseOptions("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_UsePreferredLifetime(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_use_preferred_lifetime"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUsePreferredLifetime("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUsePreferredLifetime("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_UseRecycleLeases(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_use_recycle_leases"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseRecycleLeases("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseRecycleLeases("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_UseUpdateDnsOnLeaseRenewal(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_use_update_dns_on_lease_renewal"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseUpdateDnsOnLeaseRenewal("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseUpdateDnsOnLeaseRenewal("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_UseValidLifetime(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_use_valid_lifetime"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseValidLifetime("NAME_REPLACE_ME", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseValidLifetime("NAME_REPLACE_ME", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_ValidLifetime(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_valid_lifetime"
	var v ipam.Ipv6networktemplate

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateValidLifetime("NAME_REPLACE_ME", "VALID_LIFETIME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "VALID_LIFETIME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateValidLifetime("NAME_REPLACE_ME", "VALID_LIFETIME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "VALID_LIFETIME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckIpv6networktemplateExists(ctx context.Context, resourceName string, v *ipam.Ipv6networktemplate) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.IPAMAPI.
			Ipv6networktemplateAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForIpv6networktemplate).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetIpv6networktemplateResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetIpv6networktemplateResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckIpv6networktemplateDestroy(ctx context.Context, v *ipam.Ipv6networktemplate) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.IPAMAPI.
			Ipv6networktemplateAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForIpv6networktemplate).
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

func testAccCheckIpv6networktemplateDisappears(ctx context.Context, v *ipam.Ipv6networktemplate) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.IPAMAPI.
			Ipv6networktemplateAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccIpv6networktemplateImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
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

func testAccIpv6networktemplateBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test" {
    name = %q
}
`, name)
}

func testAccIpv6networktemplateAllowAnyNetmask(name string, allowAnyNetmask string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_allow_any_netmask" {
    name = %q
    allow_any_netmask = %q
}
`, name, allowAnyNetmask)
}

func testAccIpv6networktemplateAutoCreateReversezone(name string, autoCreateReversezone string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_auto_create_reversezone" {
    name = %q
    auto_create_reversezone = %q
}
`, name, autoCreateReversezone)
}

func testAccIpv6networktemplateCidr(name string, cidr string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_cidr" {
    name = %q
    cidr = %q
}
`, name, cidr)
}

func testAccIpv6networktemplateCloudApiCompatible(name string, cloudApiCompatible string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_cloud_api_compatible" {
    name = %q
    cloud_api_compatible = %q
}
`, name, cloudApiCompatible)
}

func testAccIpv6networktemplateComment(name string, comment string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccIpv6networktemplateDdnsDomainname(name string, ddnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_ddns_domainname" {
    name = %q
    ddns_domainname = %q
    use_ddns_domainname = true
}
`, name, ddnsDomainname)
}

func testAccIpv6networktemplateDdnsEnableOptionFqdn(name string, ddnsEnableOptionFqdn string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_ddns_enable_option_fqdn" {
    name = %q
    ddns_enable_option_fqdn = %q
    use_ddns_enable_option_fqdn = true
}
`, name, ddnsEnableOptionFqdn)
}

func testAccIpv6networktemplateDdnsGenerateHostname(name string, ddnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_ddns_generate_hostname" {
    name = %q
    ddns_generate_hostname = %q
    use_ddns_generate_hostname = true
}
`, name, ddnsGenerateHostname)
}

func testAccIpv6networktemplateDdnsServerAlwaysUpdates(name string, ddnsServerAlwaysUpdates string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_ddns_server_always_updates" {
    name = %q
    ddns_server_always_updates = %q
}
`, name, ddnsServerAlwaysUpdates)
}

func testAccIpv6networktemplateDdnsTtl(name string, ddnsTtl string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_ddns_ttl" {
    name = %q
    ddns_ttl = %q
    use_ddns_ttl = true
}
`, name, ddnsTtl)
}

func testAccIpv6networktemplateDelegatedMember(name string, delegatedMember map[string]string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_delegated_member" {
    name = %q
    delegated_member = %s
}
`, name, delegatedMember)
}

func testAccIpv6networktemplateDomainName(name string, domainName string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_domain_name" {
    name = %q
    domain_name = %q
    use_domain_name = true
}
`, name, domainName)
}

func testAccIpv6networktemplateDomainNameServers(name string, domainNameServers []string) string {
	domainNameServersStr := utils.ConvertStringSliceToHCL(domainNameServers)
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_domain_name_servers" {
    name = %q
    domain_name_servers = %s
    use_domain_name_servers = true
}
`, name, domainNameServersStr)
}

func testAccIpv6networktemplateEnableDdns(name string, enableDdns string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_enable_ddns" {
    name = %q
    enable_ddns = %q
    use_enable_ddns = true
}
`, name, enableDdns)
}

func testAccIpv6networktemplateExtAttrs(name string, extAttrs map[string]string) string {
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		extAttrsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	extAttrsStr += "  }"
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_extattrs" {
    name = %q
    extattrs = %s
}
`, name, extAttrsStr)
}

func testAccIpv6networktemplateFixedAddressTemplates(name string, fixedAddressTemplates []string) string {
	fixedAddressTemplatesStr := utils.ConvertStringSliceToHCL(fixedAddressTemplates)
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_fixed_address_templates" {
    name = %q
    fixed_address_templates = %s
}
`, name, fixedAddressTemplatesStr)
}

func testAccIpv6networktemplateIpv6prefix(name string, ipv6prefix string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_ipv6prefix" {
    name = %q
    ipv6prefix = %q
}
`, name, ipv6prefix)
}

func testAccIpv6networktemplateLogicFilterRules(name string, logicFilterRules []map[string]any) string {
	logicFilterRulesStr := utils.ConvertSliceOfMapsToHCL(logicFilterRules)
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_logic_filter_rules" {
    name = %q
    logic_filter_rules = %s
    use_logic_filter_rules = true
}
`, name, logicFilterRulesStr)
}

func testAccIpv6networktemplateMembers(name string, members []map[string]any) string {
	membersStr := utils.ConvertSliceOfMapsToHCL(members)
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_members" {
    name = %q
    members = %s
}
`, name, membersStr)
}

func testAccIpv6networktemplateName(name string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_name" {
    name = %q
}
`, name)
}

func testAccIpv6networktemplateOptions(name string, options []map[string]any) string {
	optionsStr := utils.ConvertSliceOfMapsToHCL(options)
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_options" {
    name = %q
    options = %s
    use_options = true
}
`, name, optionsStr)
}

func testAccIpv6networktemplatePreferredLifetime(name string, preferredLifetime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_preferred_lifetime" {
    name = %q
    preferred_lifetime = %q
    use_preferred_lifetime = true
}
`, name, preferredLifetime)
}

func testAccIpv6networktemplateRangeTemplates(name string, rangeTemplates []string) string {
	rangeTemplatesStr := utils.ConvertStringSliceToHCL(rangeTemplates)
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_range_templates" {
    name = %q
    range_templates = %s
}
`, name, rangeTemplatesStr)
}

func testAccIpv6networktemplateRecycleLeases(name string, recycleLeases string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_recycle_leases" {
    name = %q
    recycle_leases = %q
    use_recycle_leases = true
}
`, name, recycleLeases)
}

func testAccIpv6networktemplateRirOrganization(name string, rirOrganization string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_rir_organization" {
    name = %q
    rir_organization = %q
}
`, name, rirOrganization)
}

func testAccIpv6networktemplateRirRegistrationAction(name string, rirRegistrationAction string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_rir_registration_action" {
    name = %q
    rir_registration_action = %q
}
`, name, rirRegistrationAction)
}

func testAccIpv6networktemplateRirRegistrationStatus(name string, rirRegistrationStatus string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_rir_registration_status" {
    name = %q
    rir_registration_status = %q
}
`, name, rirRegistrationStatus)
}

func testAccIpv6networktemplateSendRirRequest(name string, sendRirRequest string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_send_rir_request" {
    name = %q
    send_rir_request = %q
}
`, name, sendRirRequest)
}

func testAccIpv6networktemplateUpdateDnsOnLeaseRenewal(name string, updateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_update_dns_on_lease_renewal" {
    name = %q
    update_dns_on_lease_renewal = %q
    use_update_dns_on_lease_renewal = true
}
`, name, updateDnsOnLeaseRenewal)
}

func testAccIpv6networktemplateUseDdnsDomainname(name string, useDdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_ddns_domainname" {
    name = %q
    use_ddns_domainname = %q
}
`, name, useDdnsDomainname)
}

func testAccIpv6networktemplateUseDdnsEnableOptionFqdn(name string, useDdnsEnableOptionFqdn string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_ddns_enable_option_fqdn" {
    name = %q
    use_ddns_enable_option_fqdn = %q
}
`, name, useDdnsEnableOptionFqdn)
}

func testAccIpv6networktemplateUseDdnsGenerateHostname(name string, useDdnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_ddns_generate_hostname" {
    name = %q
    use_ddns_generate_hostname = %q
}
`, name, useDdnsGenerateHostname)
}

func testAccIpv6networktemplateUseDdnsTtl(name string, useDdnsTtl string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_ddns_ttl" {
    name = %q
    use_ddns_ttl = %q
}
`, name, useDdnsTtl)
}

func testAccIpv6networktemplateUseDomainName(name string, useDomainName string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_domain_name" {
    name = %q
    use_domain_name = %q
}
`, name, useDomainName)
}

func testAccIpv6networktemplateUseDomainNameServers(name string, useDomainNameServers string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_domain_name_servers" {
    name = %q
    use_domain_name_servers = %q
}
`, name, useDomainNameServers)
}

func testAccIpv6networktemplateUseEnableDdns(name string, useEnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_enable_ddns" {
    name = %q
    use_enable_ddns = %q
}
`, name, useEnableDdns)
}

func testAccIpv6networktemplateUseLogicFilterRules(name string, useLogicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_logic_filter_rules" {
    name = %q
    use_logic_filter_rules = %q
}
`, name, useLogicFilterRules)
}

func testAccIpv6networktemplateUseOptions(name string, useOptions string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_options" {
    name = %q
    use_options = %q
}
`, name, useOptions)
}

func testAccIpv6networktemplateUsePreferredLifetime(name string, usePreferredLifetime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_preferred_lifetime" {
    name = %q
    use_preferred_lifetime = %q
}
`, name, usePreferredLifetime)
}

func testAccIpv6networktemplateUseRecycleLeases(name string, useRecycleLeases string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_recycle_leases" {
    name = %q
    use_recycle_leases = %q
}
`, name, useRecycleLeases)
}

func testAccIpv6networktemplateUseUpdateDnsOnLeaseRenewal(name string, useUpdateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_update_dns_on_lease_renewal" {
    name = %q
    use_update_dns_on_lease_renewal = %q
}
`, name, useUpdateDnsOnLeaseRenewal)
}

func testAccIpv6networktemplateUseValidLifetime(name string, useValidLifetime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_valid_lifetime" {
    name = %q
    use_valid_lifetime = %q
}
`, name, useValidLifetime)
}

func testAccIpv6networktemplateValidLifetime(name string, validLifetime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_valid_lifetime" {
    name = %q
    valid_lifetime = %q
    use_valid_lifetime = true
}
`, name, validLifetime)
}
