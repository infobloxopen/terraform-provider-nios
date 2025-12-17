package ipam_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForIpv6networktemplate = "allow_any_netmask,auto_create_reversezone,cidr,cloud_api_compatible,comment,ddns_domainname,ddns_enable_option_fqdn,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,delegated_member,domain_name,domain_name_servers,enable_ddns,extattrs,fixed_address_templates,ipv6prefix,logic_filter_rules,members,name,options,preferred_lifetime,range_templates,recycle_leases,rir,rir_organization,rir_registration_action,rir_registration_status,send_rir_request,update_dns_on_lease_renewal,use_ddns_domainname,use_ddns_enable_option_fqdn,use_ddns_generate_hostname,use_ddns_ttl,use_domain_name,use_domain_name_servers,use_enable_ddns,use_logic_filter_rules,use_options,use_preferred_lifetime,use_recycle_leases,use_update_dns_on_lease_renewal,use_valid_lifetime,valid_lifetime"

func TestAccIpv6networktemplateResource_basic(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test"
	var v ipam.Ipv6networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateBasicConfig(name, 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "allow_any_netmask", "false"),
					resource.TestCheckResourceAttr(resourceName, "auto_create_reversezone", "false"),
					resource.TestCheckResourceAttr(resourceName, "cloud_api_compatible", "true"),
					resource.TestCheckResourceAttr(resourceName, "ddns_enable_option_fqdn", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "true"),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "0"),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "false"),
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
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_disappears(t *testing.T) {
	resourceName := "nios_ipam_ipv6networktemplate.test"
	var v ipam.Ipv6networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6networktemplateDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6networktemplateBasicConfig(name, 24),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateBasicConfig(name, 24),
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
				ImportStateVerifyIgnore:              []string{"options"},
				PlanOnly:                             true,
			},
			// Import and Verify
			{
				ResourceName:                         resourceName,
				ImportState:                          true,
				ImportStateIdFunc:                    testAccIpv6networktemplateImportStateIdFunc(resourceName),
				ImportStateVerify:                    true,
				ImportStateVerifyIgnore:              []string{"extattrs_all", "options"},
				ImportStateVerifyIdentifierAttribute: "ref",
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_AllowAnyNetmask(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_allow_any_netmask"
	var v ipam.Ipv6networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateAllowAnyNetmask(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_any_netmask", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateAllowAnyNetmask(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateAutoCreateReversezone(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_create_reversezone", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateAutoCreateReversezone(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateCidr(name, 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cidr", "24"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateCidr(name, 32),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cidr", "32"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_CloudApiCompatible(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_cloud_api_compatible"
	var v ipam.Ipv6networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateCloudApiCompatible(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_api_compatible", "true"),
				),
			},
			// Unable to set to false as we use "Terraform Internal ID" as an EA which is cloud_compatible
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_Comment(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_comment"
	var v ipam.Ipv6networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateComment(name, 24, "Comment for the object"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for the object"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateComment(name, 24, "Updated comment for the object"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateDdnsDomainname(name, 24, "example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateDdnsDomainname(name, 24, "updated-example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domainname", "updated-example.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_DdnsEnableOptionFqdn(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_ddns_enable_option_fqdn"
	var v ipam.Ipv6networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateDdnsEnableOptionFqdn(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_enable_option_fqdn", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateDdnsEnableOptionFqdn(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateDdnsGenerateHostname(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateDdnsGenerateHostname(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateDdnsServerAlwaysUpdates(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_server_always_updates", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateDdnsServerAlwaysUpdates(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateDdnsTtl(name, 24, "1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "1"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateDdnsTtl(name, 24, "1000"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl", "1000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccIpv6networktemplateResource_DelegatedMember(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_delegated_member"
	var v ipam.Ipv6networktemplate
	name := acctest.RandomNameWithPrefix("network-template")
	delegatedMemberVal := map[string]any{
		"name":     "infoblox.member",
		"ipv4addr": "6.6.6.6",
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
				Config: testAccIpv6networktemplateDelegatedMember(name, 24, delegatedMemberVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegated_member.name", "infoblox.member"),
					resource.TestCheckResourceAttr(resourceName, "delegated_member.ipv4addr", "6.6.6.6"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateDelegatedMember(name, 24, delegatedMemberValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegated_member.name", "infoblox.localdomain"),
					resource.TestCheckResourceAttr(resourceName, "delegated_member.ipv4addr", "1.1.1.1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccIpv6networktemplateResource_DomainName(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_domain_name"
	var v ipam.Ipv6networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateDomainName(name, 24, "ddns_domain.name"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name", "ddns_domain.name"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateDomainName(name, 24, "updated_ddns_domain.name"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name", "updated_ddns_domain.name"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_DomainNameServers(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_domain_name_servers"
	var v ipam.Ipv6networktemplate
	name := acctest.RandomNameWithPrefix("network-template")
	domainNameServersVal := []string{"2001:4860:4860::8888", "2001:4860:4860::9999", "2001:4860:4860::8899"}
	domainNameServersValUpdated := []string{"2001:4860:4860::8888", "2001:4860:4860::8844"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateDomainNameServers(name, 24, domainNameServersVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.0", "2001:4860:4860::8888"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.1", "2001:4860:4860::9999"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.2", "2001:4860:4860::8899"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateDomainNameServers(name, 24, domainNameServersValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.0", "2001:4860:4860::8888"),
					resource.TestCheckResourceAttr(resourceName, "domain_name_servers.1", "2001:4860:4860::8844"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_EnableDdns(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_enable_ddns"
	var v ipam.Ipv6networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateEnableDdns(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateEnableDdns(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateExtAttrs(name, 24, map[string]string{
					"Tenant ID": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Tenant ID", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateExtAttrs(name, 24, map[string]string{
					"Tenant ID": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Tenant ID", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_FixedAddressTemplates(t *testing.T) {
	t.Skip("FA Template cannot be cloud compatible. Skipping test for cloud users as we use a EA with cloud compatible enabled.")
	var resourceName = "nios_ipam_ipv6networktemplate.test_fixed_address_templates"
	var v ipam.Ipv6networktemplate
	name := acctest.RandomNameWithPrefix("network-template")
	var fixedAddressTemplatesVal []string
	var fixedAddressTemplatesValUpdated []string

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateFixedAddressTemplates(name, 24, fixedAddressTemplatesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fixed_address_templates", "FIXED_ADDRESS_TEMPLATES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateFixedAddressTemplates(name, 24, fixedAddressTemplatesValUpdated),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateIpv6prefix(name, 24, "2001:db8:abcd:12::/64"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6prefix", "2001:db8:abcd:12::/64"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateIpv6prefix(name, 24, "2001:db8:cdef:12::/64"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ipv6prefix", "2001:db8:cdef:12::/64"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_LogicFilterRules(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_logic_filter_rules"
	var v ipam.Ipv6networktemplate
	name := acctest.RandomNameWithPrefix("network-template")
	logicFilterRulesVal := []map[string]any{
		{
			"filter": "ipv6_nac_filter",
			"type":   "NAC",
		},
	}
	logicFilterRulesValUpdated := []map[string]any{
		{
			"filter": "ipv6_option_filter",
			"type":   "Option",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateLogicFilterRules(name, 24, logicFilterRulesVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.filter", "ipv6_nac_filter"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.type", "NAC"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateLogicFilterRules(name, 24, logicFilterRulesValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.filter", "ipv6_option_filter"),
					resource.TestCheckResourceAttr(resourceName, "logic_filter_rules.0.type", "Option"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_Members(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_members"
	var v ipam.Ipv6networktemplate
	name := acctest.RandomNameWithPrefix("network-template")
	membersVal := []map[string]any{
		{
			"struct": "dhcpmember",
			"name":   "infoblox.localdomain",
		},
	}
	membersValUpdated := []map[string]any{
		{
			"struct": "dhcpmember",
			"name":   "infoblox.member",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateMembers(name, 24, membersVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "members.struct", "dhcpmember"),
					resource.TestCheckResourceAttr(resourceName, "members.name", "infoblox.localdomain"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateMembers(name, 24, membersValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "members.struct", "dhcpmember"),
					resource.TestCheckResourceAttr(resourceName, "members.name", "infoblox.member"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_Name(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_name"
	var v ipam.Ipv6networktemplate
	name1 := acctest.RandomNameWithPrefix("network-template")
	name2 := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateName(name1, 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateName(name2, 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_Options(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_options"
	var v ipam.Ipv6networktemplate
	name := acctest.RandomNameWithPrefix("network-template")
	optionsVal := []map[string]any{
		{
			"name":  "domain-name",
			"num":   "15",
			"value": "example.com",
		},
		{
			"num":          "37",
			"value":        "remote-id",
			"vendor_class": "DHCPv6",
		},
		{
			"name":         "dhcp6.subscriber-id",
			"value":        "subscriber-id",
			"vendor_class": "DHCPv6",
		},
	}
	optionsValUpdated := []map[string]any{
		{
			"name":  "domain-name",
			"num":   "15",
			"value": "example.org",
		},
		{
			"num":          "37",
			"value":        "remote-id-updated",
			"vendor_class": "DHCPv6",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateOptions(name, 24, optionsVal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "domain-name"),
					resource.TestCheckResourceAttr(resourceName, "options.0.num", "15"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "options.1.num", "37"),
					resource.TestCheckResourceAttr(resourceName, "options.1.value", "remote-id"),
					resource.TestCheckResourceAttr(resourceName, "options.1.vendor_class", "DHCPv6"),
					resource.TestCheckResourceAttr(resourceName, "options.2.name", "dhcp6.subscriber-id"),
					resource.TestCheckResourceAttr(resourceName, "options.2.value", "subscriber-id"),
					resource.TestCheckResourceAttr(resourceName, "options.2.vendor_class", "DHCPv6"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateOptions(name, 24, optionsValUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "options.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "options.0.name", "domain-name"),
					resource.TestCheckResourceAttr(resourceName, "options.0.num", "15"),
					resource.TestCheckResourceAttr(resourceName, "options.0.value", "example.org"),
					resource.TestCheckResourceAttr(resourceName, "options.1.num", "37"),
					resource.TestCheckResourceAttr(resourceName, "options.1.value", "remote-id-updated"),
					resource.TestCheckResourceAttr(resourceName, "options.1.vendor_class", "DHCPv6"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_PreferredLifetime(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_preferred_lifetime"
	var v ipam.Ipv6networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplatePreferredLifetime(name, 24, "200"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "200"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplatePreferredLifetime(name, 24, "2000"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "preferred_lifetime", "2000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccIpv6networktemplateResource_RangeTemplates(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_range_templates"
	var v ipam.Ipv6networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateRangeTemplates(name, 24, "one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "range_templates.0", "nios_dhcp_ipv6_range_template.one", "name"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateRangeTemplates(name, 24, "two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "range_templates.0", "nios_dhcp_ipv6_range_template.two", "name"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_RecycleLeases(t *testing.T) {
	var resourceName = "nios_ipam_ipv6networktemplate.test_recycle_leases"
	var v ipam.Ipv6networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateRecycleLeases(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recycle_leases", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateRecycleLeases(name, 24, "false"),
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
	t.Skip("Cloud-compatible templates cannot set RIR registration action")
	var resourceName = "nios_ipam_ipv6networktemplate.test_rir_organization"
	var v ipam.Ipv6networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateRirOrganization(name, 24, "test"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_organization", "test1"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateRirOrganization(name, 24, "test"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_organization", "test1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpv6networktemplateResource_RirRegistrationAction(t *testing.T) {
	t.Skip("Cloud-compatible templates cannot set RIR registration action")
	var resourceName = "nios_ipam_ipv6networktemplate.test_rir_registration_action"
	var v ipam.Ipv6networktemplate
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateRirRegistrationAction(name, 24, "CREATE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_action", "CREATE"),
				),
			},
			{
				Config: testAccIpv6networktemplateRirRegistrationAction(name, 24, "NONE"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateRirRegistrationStatus(name, 24, "NOT_REGISTERED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rir_registration_status", "NOT_REGISTERED"),
				),
			},
			{
				Config: testAccIpv6networktemplateRirRegistrationStatus(name, 24, "REGISTERED"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateSendRirRequest(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "send_rir_request", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateSendRirRequest(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUpdateDnsOnLeaseRenewal(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_on_lease_renewal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUpdateDnsOnLeaseRenewal(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseDdnsDomainname(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_domainname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseDdnsDomainname(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseDdnsEnableOptionFqdn(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_enable_option_fqdn", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseDdnsEnableOptionFqdn(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseDdnsGenerateHostname(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_generate_hostname", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseDdnsGenerateHostname(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseDdnsTtl(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseDdnsTtl(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseDomainName(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseDomainName(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseDomainNameServers(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_domain_name_servers", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseDomainNameServers(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseEnableDdns(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_enable_ddns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseEnableDdns(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseLogicFilterRules(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_logic_filter_rules", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseLogicFilterRules(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseOptions(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_options", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseOptions(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUsePreferredLifetime(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_preferred_lifetime", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUsePreferredLifetime(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseRecycleLeases(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recycle_leases", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseRecycleLeases(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseUpdateDnsOnLeaseRenewal(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_update_dns_on_lease_renewal", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseUpdateDnsOnLeaseRenewal(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateUseValidLifetime(name, 24, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_valid_lifetime", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateUseValidLifetime(name, 24, "false"),
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
	name := acctest.RandomNameWithPrefix("network-template")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpv6networktemplateValidLifetime(name, 24, "200"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "200"),
				),
			},
			// Update and Read
			{
				Config: testAccIpv6networktemplateValidLifetime(name, 24, "400"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6networktemplateExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "valid_lifetime", "400"),
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

func testAccBaseWithIPv6RangeTemplates(template1, template2 string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6_range_template" "one" {
    name = %q
    number_of_addresses = 100
    offset = 50
    cloud_api_compatible = true
}

resource "nios_dhcp_ipv6_range_template" "two" {
    name = %q
    number_of_addresses = 100
    offset = 50
    cloud_api_compatible = true
}
`, template1, template2)
}

func testAccIpv6networktemplateBasicConfig(name string, cidr int) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test" {
    name = %q
    cidr = %d
}
`, name, cidr)
}

func testAccIpv6networktemplateAllowAnyNetmask(name string, cidr int, allowAnyNetmask string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_allow_any_netmask" {
    name = %q
    cidr = %d
    allow_any_netmask = %q
}
`, name, cidr, allowAnyNetmask)
}

func testAccIpv6networktemplateAutoCreateReversezone(name string, cidr int, autoCreateReversezone string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_auto_create_reversezone" {
    name = %q
    cidr = %d
    auto_create_reversezone = %q
}
`, name, cidr, autoCreateReversezone)
}

func testAccIpv6networktemplateCidr(name string, cidr int) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_cidr" {
    name = %q
    cidr = %d
}
`, name, cidr)
}

func testAccIpv6networktemplateCloudApiCompatible(name string, cidr int, cloudApiCompatible string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_cloud_api_compatible" {
    name = %q
    cidr = %d
    cloud_api_compatible = %q
}
`, name, cidr, cloudApiCompatible)
}

func testAccIpv6networktemplateComment(name string, cidr int, comment string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_comment" {
    name = %q
    cidr = %d
    comment = %q
}
`, name, cidr, comment)
}

func testAccIpv6networktemplateDdnsDomainname(name string, cidr int, ddnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_ddns_domainname" {
    name = %q
    cidr = %d
    ddns_domainname = %q
    use_ddns_domainname = true
}
`, name, cidr, ddnsDomainname)
}

func testAccIpv6networktemplateDdnsEnableOptionFqdn(name string, cidr int, ddnsEnableOptionFqdn string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_ddns_enable_option_fqdn" {
    name = %q
    cidr = %d
    ddns_enable_option_fqdn = %q
    use_ddns_enable_option_fqdn = true
}
`, name, cidr, ddnsEnableOptionFqdn)
}

func testAccIpv6networktemplateDdnsGenerateHostname(name string, cidr int, ddnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_ddns_generate_hostname" {
    name = %q
    cidr = %d
    ddns_generate_hostname = %q
    use_ddns_generate_hostname = true
}
`, name, cidr, ddnsGenerateHostname)
}

func testAccIpv6networktemplateDdnsServerAlwaysUpdates(name string, cidr int, ddnsServerAlwaysUpdates string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_ddns_server_always_updates" {
    name = %q
    cidr = %d
    ddns_server_always_updates = %q
    ddns_enable_option_fqdn = true
    use_ddns_enable_option_fqdn = true
}
`, name, cidr, ddnsServerAlwaysUpdates)
}

func testAccIpv6networktemplateDdnsTtl(name string, cidr int, ddnsTtl string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_ddns_ttl" {
    name = %q
    cidr = %d
    ddns_ttl = %q
    use_ddns_ttl = true
}
`, name, cidr, ddnsTtl)
}

func testAccIpv6networktemplateDelegatedMember(name string, cidr int, delegatedMember map[string]any) string {
	delegatedMemberStr := utils.ConvertMapToHCL(delegatedMember)
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_delegated_member" {
    name = %q
    cidr = %d
    delegated_member = %s
}
`, name, cidr, delegatedMemberStr)
}

func testAccIpv6networktemplateDomainName(name string, cidr int, domainName string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_domain_name" {
    name = %q
    cidr = %d
    domain_name = %q
    use_domain_name = true
}
`, name, cidr, domainName)
}

func testAccIpv6networktemplateDomainNameServers(name string, cidr int, domainNameServers []string) string {
	domainNameServersStr := utils.ConvertStringSliceToHCL(domainNameServers)
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_domain_name_servers" {
    name = %q
    cidr = %d
    domain_name_servers = %s
    use_domain_name_servers = true
}
`, name, cidr, domainNameServersStr)
}

func testAccIpv6networktemplateEnableDdns(name string, cidr int, enableDdns string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_enable_ddns" {
    name = %q
    cidr = %d
    enable_ddns = %q
    use_enable_ddns = true
}
`, name, cidr, enableDdns)
}

func testAccIpv6networktemplateExtAttrs(name string, cidr int, extAttrs map[string]string) string {
	extAttrsStr := "{\n"
	for k, v := range extAttrs {
		// Quote keys that contain spaces
		key := k
		if strings.Contains(k, " ") {
			key = fmt.Sprintf("%q", k)
		}
		extAttrsStr += fmt.Sprintf("    %s = %q\n", key, v)
	}
	extAttrsStr += "  }"
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_extattrs" {
    name = %q
    cidr = %d
    extattrs = %s
}
`, name, cidr, extAttrsStr)
}

func testAccIpv6networktemplateFixedAddressTemplates(name string, cidr int, fixedAddressTemplates []string) string {
	fixedAddressTemplatesStr := utils.ConvertStringSliceToHCL(fixedAddressTemplates)
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_fixed_address_templates" {
    name = %q
    cidr = %d
    fixed_address_templates = %s
}
`, name, cidr, fixedAddressTemplatesStr)
}

func testAccIpv6networktemplateIpv6prefix(name string, cidr int, ipv6prefix string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_ipv6prefix" {
    name = %q
    cidr = %d
    ipv6prefix = %q
}
`, name, cidr, ipv6prefix)
}

func testAccIpv6networktemplateLogicFilterRules(name string, cidr int, logicFilterRules []map[string]any) string {
	logicFilterRulesStr := utils.ConvertSliceOfMapsToHCL(logicFilterRules)
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_logic_filter_rules" {
    name = %q
    cidr = %d
    logic_filter_rules = %s
    use_logic_filter_rules = true
}
`, name, cidr, logicFilterRulesStr)
}

func testAccIpv6networktemplateMembers(name string, cidr int, members []map[string]any) string {
	membersStr := utils.ConvertSliceOfMapsToHCL(members)
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_members" {
    name = %q
    cidr = %d
    members = %s
}
`, name, cidr, membersStr)
}

func testAccIpv6networktemplateName(name string, cidr int) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_name" {
    name = %q
    cidr = %d
}
`, name, cidr)
}

func testAccIpv6networktemplateOptions(name string, cidr int, options []map[string]any) string {
	optionsStr := utils.ConvertSliceOfMapsToHCL(options)
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_options" {
    name = %q
    cidr = %d
    options = %s
    use_options = true
}
`, name, cidr, optionsStr)
}

func testAccIpv6networktemplatePreferredLifetime(name string, cidr int, preferredLifetime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_preferred_lifetime" {
    name = %q
    cidr = %d
    preferred_lifetime = %q
    use_preferred_lifetime = true
}
`, name, cidr, preferredLifetime)
}

func testAccIpv6networktemplateRangeTemplates(name string, cidr int, rangeTemplates string) string {
	config := fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_range_templates" {
    name = %q
    cidr = %d
    range_templates = [nios_dhcp_ipv6_range_template.%s.name]
    depends_on = [nios_dhcp_ipv6_range_template.one, nios_dhcp_ipv6_range_template.two]
}
`, name, cidr, rangeTemplates)
	return strings.Join([]string{testAccBaseWithIPv6RangeTemplates(acctest.RandomName(), acctest.RandomName()), config}, "")

}

func testAccIpv6networktemplateRecycleLeases(name string, cidr int, recycleLeases string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_recycle_leases" {
    name = %q
    cidr = %d
    recycle_leases = %q
    use_recycle_leases = true
}
`, name, cidr, recycleLeases)
}

func testAccIpv6networktemplateRirOrganization(name string, cidr int, rirOrganization string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_rir_organization" {
    name = %q
    cidr = %d
    rir_organization = %q
}
`, name, cidr, rirOrganization)
}

func testAccIpv6networktemplateRirRegistrationAction(name string, cidr int, rirRegistrationAction string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_rir_registration_action" {
    name = %q
    cidr = %d
    rir_registration_action = %q
    rir_organization = "test"
	cloud_api_compatible = false
}
`, name, cidr, rirRegistrationAction)
}

func testAccIpv6networktemplateRirRegistrationStatus(name string, cidr int, rirRegistrationStatus string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_rir_registration_status" {
    name = %q
    cidr = %d
    rir_registration_status = %q
}
`, name, cidr, rirRegistrationStatus)
}

func testAccIpv6networktemplateSendRirRequest(name string, cidr int, sendRirRequest string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_send_rir_request" {
    name = %q
    cidr = %d
    send_rir_request = %q
}
`, name, cidr, sendRirRequest)
}

func testAccIpv6networktemplateUpdateDnsOnLeaseRenewal(name string, cidr int, updateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_update_dns_on_lease_renewal" {
    name = %q
    cidr = %d
    update_dns_on_lease_renewal = %q
    use_update_dns_on_lease_renewal = true
}
`, name, cidr, updateDnsOnLeaseRenewal)
}

func testAccIpv6networktemplateUseDdnsDomainname(name string, cidr int, useDdnsDomainname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_ddns_domainname" {
    name = %q
    cidr = %d
    use_ddns_domainname = %q
}
`, name, cidr, useDdnsDomainname)
}

func testAccIpv6networktemplateUseDdnsEnableOptionFqdn(name string, cidr int, useDdnsEnableOptionFqdn string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_ddns_enable_option_fqdn" {
    name = %q
    cidr = %d
    use_ddns_enable_option_fqdn = %q
}
`, name, cidr, useDdnsEnableOptionFqdn)
}

func testAccIpv6networktemplateUseDdnsGenerateHostname(name string, cidr int, useDdnsGenerateHostname string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_ddns_generate_hostname" {
    name = %q
    cidr = %d
    use_ddns_generate_hostname = %q
}
`, name, cidr, useDdnsGenerateHostname)
}

func testAccIpv6networktemplateUseDdnsTtl(name string, cidr int, useDdnsTtl string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_ddns_ttl" {
    name = %q
    cidr = %d
    use_ddns_ttl = %q
}
`, name, cidr, useDdnsTtl)
}

func testAccIpv6networktemplateUseDomainName(name string, cidr int, useDomainName string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_domain_name" {
    name = %q
    cidr = %d
    use_domain_name = %q
    domain_name = "example.com"
}
`, name, cidr, useDomainName)
}

func testAccIpv6networktemplateUseDomainNameServers(name string, cidr int, useDomainNameServers string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_domain_name_servers" {
    name = %q
    cidr = %d
    use_domain_name_servers = %q
}
`, name, cidr, useDomainNameServers)
}

func testAccIpv6networktemplateUseEnableDdns(name string, cidr int, useEnableDdns string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_enable_ddns" {
    name = %q
    cidr = %d
    use_enable_ddns = %q
}
`, name, cidr, useEnableDdns)
}

func testAccIpv6networktemplateUseLogicFilterRules(name string, cidr int, useLogicFilterRules string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_logic_filter_rules" {
    name = %q
    cidr = %d
    use_logic_filter_rules = %q
}
`, name, cidr, useLogicFilterRules)
}

func testAccIpv6networktemplateUseOptions(name string, cidr int, useOptions string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_options" {
    name = %q
    cidr = %d
    use_options = %q
}
`, name, cidr, useOptions)
}

func testAccIpv6networktemplateUsePreferredLifetime(name string, cidr int, usePreferredLifetime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_preferred_lifetime" {
    name = %q
    cidr = %d
    use_preferred_lifetime = %q
}
`, name, cidr, usePreferredLifetime)
}

func testAccIpv6networktemplateUseRecycleLeases(name string, cidr int, useRecycleLeases string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_recycle_leases" {
    name = %q
    cidr = %d
    use_recycle_leases = %q
}
`, name, cidr, useRecycleLeases)
}

func testAccIpv6networktemplateUseUpdateDnsOnLeaseRenewal(name string, cidr int, useUpdateDnsOnLeaseRenewal string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_update_dns_on_lease_renewal" {
    name = %q
    cidr = %d
    use_update_dns_on_lease_renewal = %q
}
`, name, cidr, useUpdateDnsOnLeaseRenewal)
}

func testAccIpv6networktemplateUseValidLifetime(name string, cidr int, useValidLifetime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_use_valid_lifetime" {
    name = %q
    cidr = %d
    use_valid_lifetime = %q
}
`, name, cidr, useValidLifetime)
}

func testAccIpv6networktemplateValidLifetime(name string, cidr int, validLifetime string) string {
	return fmt.Sprintf(`
resource "nios_ipam_ipv6networktemplate" "test_valid_lifetime" {
    name = %q
    cidr = %d
    valid_lifetime = %q
    use_valid_lifetime = true
}
`, name, cidr, validLifetime)
}
