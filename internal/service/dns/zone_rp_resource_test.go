package dns_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForZoneRp = "address,comment,disable,display_domain,dns_soa_email,extattrs,external_primaries,external_secondaries,fireeye_rule_mapping,fqdn,grid_primary,grid_secondaries,locked,locked_by,log_rpz,mask_prefix,member_soa_mnames,member_soa_serials,network_view,ns_group,parent,prefix,primary_type,record_name_policy,rpz_drop_ip_rule_enabled,rpz_drop_ip_rule_min_prefix_length_ipv4,rpz_drop_ip_rule_min_prefix_length_ipv6,rpz_last_updated_time,rpz_policy,rpz_priority,rpz_priority_end,rpz_severity,rpz_type,soa_default_ttl,soa_email,soa_expire,soa_negative_ttl,soa_refresh,soa_retry,soa_serial_number,substitute_name,use_external_primary,use_grid_zone_timer,use_log_rpz,use_record_name_policy,use_rpz_drop_ip_rule,use_soa_email,view"

func TestAccZoneRpResource_basic(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_disappears(t *testing.T) {
	resourceName := "nios_dns_zone_rp.test"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckZoneRpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneRpBasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					testAccCheckZoneRpDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccZoneRpResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_comment"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_disable"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_extattrs"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_ExternalPrimaries(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_external_primaries"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpExternalPrimaries("EXTERNAL_PRIMARIES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_primaries", "EXTERNAL_PRIMARIES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpExternalPrimaries("EXTERNAL_PRIMARIES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_primaries", "EXTERNAL_PRIMARIES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_ExternalSecondaries(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_external_secondaries"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpExternalSecondaries("EXTERNAL_SECONDARIES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries", "EXTERNAL_SECONDARIES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpExternalSecondaries("EXTERNAL_SECONDARIES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries", "EXTERNAL_SECONDARIES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_FireeyeRuleMapping(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_fireeye_rule_mapping"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpFireeyeRuleMapping("FIREEYE_RULE_MAPPING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fireeye_rule_mapping", "FIREEYE_RULE_MAPPING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpFireeyeRuleMapping("FIREEYE_RULE_MAPPING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fireeye_rule_mapping", "FIREEYE_RULE_MAPPING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_Fqdn(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_fqdn"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpFqdn("FQDN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "FQDN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpFqdn("FQDN_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "FQDN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_GridPrimary(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_grid_primary"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpGridPrimary("GRID_PRIMARY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_primary", "GRID_PRIMARY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpGridPrimary("GRID_PRIMARY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_primary", "GRID_PRIMARY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_GridSecondaries(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_grid_secondaries"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpGridSecondaries("GRID_SECONDARIES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries", "GRID_SECONDARIES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpGridSecondaries("GRID_SECONDARIES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries", "GRID_SECONDARIES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_Locked(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_locked"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpLocked("LOCKED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "locked", "LOCKED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpLocked("LOCKED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "locked", "LOCKED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_LogRpz(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_log_rpz"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpLogRpz("LOG_RPZ_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_rpz", "LOG_RPZ_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpLogRpz("LOG_RPZ_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_rpz", "LOG_RPZ_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_MemberSoaMnames(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_member_soa_mnames"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpMemberSoaMnames("MEMBER_SOA_MNAMES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member_soa_mnames", "MEMBER_SOA_MNAMES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpMemberSoaMnames("MEMBER_SOA_MNAMES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member_soa_mnames", "MEMBER_SOA_MNAMES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_NsGroup(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_ns_group"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpNsGroup("NS_GROUP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ns_group", "NS_GROUP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpNsGroup("NS_GROUP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ns_group", "NS_GROUP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_Prefix(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_prefix"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpPrefix("PREFIX_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", "PREFIX_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpPrefix("PREFIX_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", "PREFIX_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_RecordNamePolicy(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_record_name_policy"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpRecordNamePolicy("RECORD_NAME_POLICY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "record_name_policy", "RECORD_NAME_POLICY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpRecordNamePolicy("RECORD_NAME_POLICY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "record_name_policy", "RECORD_NAME_POLICY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_RpzDropIpRuleEnabled(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_rpz_drop_ip_rule_enabled"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpRpzDropIpRuleEnabled("RPZ_DROP_IP_RULE_ENABLED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_enabled", "RPZ_DROP_IP_RULE_ENABLED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpRpzDropIpRuleEnabled("RPZ_DROP_IP_RULE_ENABLED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_enabled", "RPZ_DROP_IP_RULE_ENABLED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_RpzDropIpRuleMinPrefixLengthIpv4(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_rpz_drop_ip_rule_min_prefix_length_ipv4"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpRpzDropIpRuleMinPrefixLengthIpv4("RPZ_DROP_IP_RULE_MIN_PREFIX_LENGTH_IPV4_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv4", "RPZ_DROP_IP_RULE_MIN_PREFIX_LENGTH_IPV4_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpRpzDropIpRuleMinPrefixLengthIpv4("RPZ_DROP_IP_RULE_MIN_PREFIX_LENGTH_IPV4_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv4", "RPZ_DROP_IP_RULE_MIN_PREFIX_LENGTH_IPV4_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_RpzDropIpRuleMinPrefixLengthIpv6(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_rpz_drop_ip_rule_min_prefix_length_ipv6"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpRpzDropIpRuleMinPrefixLengthIpv6("RPZ_DROP_IP_RULE_MIN_PREFIX_LENGTH_IPV6_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv6", "RPZ_DROP_IP_RULE_MIN_PREFIX_LENGTH_IPV6_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpRpzDropIpRuleMinPrefixLengthIpv6("RPZ_DROP_IP_RULE_MIN_PREFIX_LENGTH_IPV6_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv6", "RPZ_DROP_IP_RULE_MIN_PREFIX_LENGTH_IPV6_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_RpzPolicy(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_rpz_policy"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpRpzPolicy("RPZ_POLICY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_policy", "RPZ_POLICY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpRpzPolicy("RPZ_POLICY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_policy", "RPZ_POLICY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_RpzSeverity(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_rpz_severity"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpRpzSeverity("RPZ_SEVERITY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_severity", "RPZ_SEVERITY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpRpzSeverity("RPZ_SEVERITY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_severity", "RPZ_SEVERITY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_RpzType(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_rpz_type"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpRpzType("RPZ_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_type", "RPZ_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpRpzType("RPZ_TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_type", "RPZ_TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_SetSoaSerialNumber(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_set_soa_serial_number"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpSetSoaSerialNumber("SET_SOA_SERIAL_NUMBER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "set_soa_serial_number", "SET_SOA_SERIAL_NUMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpSetSoaSerialNumber("SET_SOA_SERIAL_NUMBER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "set_soa_serial_number", "SET_SOA_SERIAL_NUMBER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_SoaDefaultTtl(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_soa_default_ttl"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpSoaDefaultTtl("SOA_DEFAULT_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_default_ttl", "SOA_DEFAULT_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpSoaDefaultTtl("SOA_DEFAULT_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_default_ttl", "SOA_DEFAULT_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_SoaEmail(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_soa_email"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpSoaEmail("SOA_EMAIL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_email", "SOA_EMAIL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpSoaEmail("SOA_EMAIL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_email", "SOA_EMAIL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_SoaExpire(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_soa_expire"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpSoaExpire("SOA_EXPIRE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_expire", "SOA_EXPIRE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpSoaExpire("SOA_EXPIRE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_expire", "SOA_EXPIRE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_SoaNegativeTtl(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_soa_negative_ttl"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpSoaNegativeTtl("SOA_NEGATIVE_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_negative_ttl", "SOA_NEGATIVE_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpSoaNegativeTtl("SOA_NEGATIVE_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_negative_ttl", "SOA_NEGATIVE_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_SoaRefresh(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_soa_refresh"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpSoaRefresh("SOA_REFRESH_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_refresh", "SOA_REFRESH_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpSoaRefresh("SOA_REFRESH_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_refresh", "SOA_REFRESH_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_SoaRetry(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_soa_retry"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpSoaRetry("SOA_RETRY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_retry", "SOA_RETRY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpSoaRetry("SOA_RETRY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_retry", "SOA_RETRY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_SoaSerialNumber(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_soa_serial_number"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpSoaSerialNumber("SOA_SERIAL_NUMBER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_serial_number", "SOA_SERIAL_NUMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpSoaSerialNumber("SOA_SERIAL_NUMBER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_serial_number", "SOA_SERIAL_NUMBER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_SubstituteName(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_substitute_name"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpSubstituteName("SUBSTITUTE_NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "substitute_name", "SUBSTITUTE_NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpSubstituteName("SUBSTITUTE_NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "substitute_name", "SUBSTITUTE_NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_UseExternalPrimary(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_use_external_primary"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpUseExternalPrimary("USE_EXTERNAL_PRIMARY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_external_primary", "USE_EXTERNAL_PRIMARY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpUseExternalPrimary("USE_EXTERNAL_PRIMARY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_external_primary", "USE_EXTERNAL_PRIMARY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_UseGridZoneTimer(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_use_grid_zone_timer"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpUseGridZoneTimer("USE_GRID_ZONE_TIMER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_grid_zone_timer", "USE_GRID_ZONE_TIMER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpUseGridZoneTimer("USE_GRID_ZONE_TIMER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_grid_zone_timer", "USE_GRID_ZONE_TIMER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_UseLogRpz(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_use_log_rpz"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpUseLogRpz("USE_LOG_RPZ_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_log_rpz", "USE_LOG_RPZ_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpUseLogRpz("USE_LOG_RPZ_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_log_rpz", "USE_LOG_RPZ_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_UseRecordNamePolicy(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_use_record_name_policy"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpUseRecordNamePolicy("USE_RECORD_NAME_POLICY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_record_name_policy", "USE_RECORD_NAME_POLICY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpUseRecordNamePolicy("USE_RECORD_NAME_POLICY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_record_name_policy", "USE_RECORD_NAME_POLICY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_UseRpzDropIpRule(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_use_rpz_drop_ip_rule"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpUseRpzDropIpRule("USE_RPZ_DROP_IP_RULE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_rpz_drop_ip_rule", "USE_RPZ_DROP_IP_RULE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpUseRpzDropIpRule("USE_RPZ_DROP_IP_RULE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_rpz_drop_ip_rule", "USE_RPZ_DROP_IP_RULE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_UseSoaEmail(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_use_soa_email"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpUseSoaEmail("USE_SOA_EMAIL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_soa_email", "USE_SOA_EMAIL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpUseSoaEmail("USE_SOA_EMAIL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_soa_email", "USE_SOA_EMAIL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_View(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_view"
	var v dns.ZoneRp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpView("VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpView("VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckZoneRpExists(ctx context.Context, resourceName string, v *dns.ZoneRp) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			ZoneRpAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForZoneRp).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetZoneRpResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetZoneRpResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckZoneRpDestroy(ctx context.Context, v *dns.ZoneRp) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			ZoneRpAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForZoneRp).
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

func testAccCheckZoneRpDisappears(ctx context.Context, v *dns.ZoneRp) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			ZoneRpAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccZoneRpBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test" {
}
`)
}

func testAccZoneRpRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccZoneRpComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccZoneRpDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccZoneRpExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccZoneRpExternalPrimaries(externalPrimaries string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_external_primaries" {
    external_primaries = %q
}
`, externalPrimaries)
}

func testAccZoneRpExternalSecondaries(externalSecondaries string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_external_secondaries" {
    external_secondaries = %q
}
`, externalSecondaries)
}

func testAccZoneRpFireeyeRuleMapping(fireeyeRuleMapping string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_fireeye_rule_mapping" {
    fireeye_rule_mapping = %q
}
`, fireeyeRuleMapping)
}

func testAccZoneRpFqdn(fqdn string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_fqdn" {
    fqdn = %q
}
`, fqdn)
}

func testAccZoneRpGridPrimary(gridPrimary string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_grid_primary" {
    grid_primary = %q
}
`, gridPrimary)
}

func testAccZoneRpGridSecondaries(gridSecondaries string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_grid_secondaries" {
    grid_secondaries = %q
}
`, gridSecondaries)
}

func testAccZoneRpLocked(locked string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_locked" {
    locked = %q
}
`, locked)
}

func testAccZoneRpLogRpz(logRpz string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_log_rpz" {
    log_rpz = %q
}
`, logRpz)
}

func testAccZoneRpMemberSoaMnames(memberSoaMnames string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_member_soa_mnames" {
    member_soa_mnames = %q
}
`, memberSoaMnames)
}

func testAccZoneRpNsGroup(nsGroup string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_ns_group" {
    ns_group = %q
}
`, nsGroup)
}

func testAccZoneRpPrefix(prefix string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_prefix" {
    prefix = %q
}
`, prefix)
}

func testAccZoneRpRecordNamePolicy(recordNamePolicy string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_record_name_policy" {
    record_name_policy = %q
}
`, recordNamePolicy)
}

func testAccZoneRpRpzDropIpRuleEnabled(rpzDropIpRuleEnabled string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_rpz_drop_ip_rule_enabled" {
    rpz_drop_ip_rule_enabled = %q
}
`, rpzDropIpRuleEnabled)
}

func testAccZoneRpRpzDropIpRuleMinPrefixLengthIpv4(rpzDropIpRuleMinPrefixLengthIpv4 string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_rpz_drop_ip_rule_min_prefix_length_ipv4" {
    rpz_drop_ip_rule_min_prefix_length_ipv4 = %q
}
`, rpzDropIpRuleMinPrefixLengthIpv4)
}

func testAccZoneRpRpzDropIpRuleMinPrefixLengthIpv6(rpzDropIpRuleMinPrefixLengthIpv6 string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_rpz_drop_ip_rule_min_prefix_length_ipv6" {
    rpz_drop_ip_rule_min_prefix_length_ipv6 = %q
}
`, rpzDropIpRuleMinPrefixLengthIpv6)
}

func testAccZoneRpRpzPolicy(rpzPolicy string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_rpz_policy" {
    rpz_policy = %q
}
`, rpzPolicy)
}

func testAccZoneRpRpzSeverity(rpzSeverity string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_rpz_severity" {
    rpz_severity = %q
}
`, rpzSeverity)
}

func testAccZoneRpRpzType(rpzType string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_rpz_type" {
    rpz_type = %q
}
`, rpzType)
}

func testAccZoneRpSetSoaSerialNumber(setSoaSerialNumber string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_set_soa_serial_number" {
    set_soa_serial_number = %q
}
`, setSoaSerialNumber)
}

func testAccZoneRpSoaDefaultTtl(soaDefaultTtl string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_soa_default_ttl" {
    soa_default_ttl = %q
}
`, soaDefaultTtl)
}

func testAccZoneRpSoaEmail(soaEmail string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_soa_email" {
    soa_email = %q
}
`, soaEmail)
}

func testAccZoneRpSoaExpire(soaExpire string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_soa_expire" {
    soa_expire = %q
}
`, soaExpire)
}

func testAccZoneRpSoaNegativeTtl(soaNegativeTtl string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_soa_negative_ttl" {
    soa_negative_ttl = %q
}
`, soaNegativeTtl)
}

func testAccZoneRpSoaRefresh(soaRefresh string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_soa_refresh" {
    soa_refresh = %q
}
`, soaRefresh)
}

func testAccZoneRpSoaRetry(soaRetry string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_soa_retry" {
    soa_retry = %q
}
`, soaRetry)
}

func testAccZoneRpSoaSerialNumber(soaSerialNumber string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_soa_serial_number" {
    soa_serial_number = %q
}
`, soaSerialNumber)
}

func testAccZoneRpSubstituteName(substituteName string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_substitute_name" {
    substitute_name = %q
}
`, substituteName)
}

func testAccZoneRpUseExternalPrimary(useExternalPrimary string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_use_external_primary" {
    use_external_primary = %q
}
`, useExternalPrimary)
}

func testAccZoneRpUseGridZoneTimer(useGridZoneTimer string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_use_grid_zone_timer" {
    use_grid_zone_timer = %q
}
`, useGridZoneTimer)
}

func testAccZoneRpUseLogRpz(useLogRpz string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_use_log_rpz" {
    use_log_rpz = %q
}
`, useLogRpz)
}

func testAccZoneRpUseRecordNamePolicy(useRecordNamePolicy string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_use_record_name_policy" {
    use_record_name_policy = %q
}
`, useRecordNamePolicy)
}

func testAccZoneRpUseRpzDropIpRule(useRpzDropIpRule string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_use_rpz_drop_ip_rule" {
    use_rpz_drop_ip_rule = %q
}
`, useRpzDropIpRule)
}

func testAccZoneRpUseSoaEmail(useSoaEmail string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_use_soa_email" {
    use_soa_email = %q
}
`, useSoaEmail)
}

func testAccZoneRpView(view string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_view" {
    view = %q
}
`, view)
}
