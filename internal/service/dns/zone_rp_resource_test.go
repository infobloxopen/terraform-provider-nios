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

//TODO : Objects to be created in the grid for testing
// - Record Name Policy - example-policy , example-policy-update

var readableAttributesForZoneRp = "address,comment,disable,display_domain,dns_soa_email,extattrs,external_primaries,external_secondaries,fireeye_rule_mapping,fqdn,grid_primary,grid_secondaries,locked,locked_by,log_rpz,mask_prefix,member_soa_mnames,member_soa_serials,network_view,ns_group,parent,prefix,primary_type,record_name_policy,rpz_drop_ip_rule_enabled,rpz_drop_ip_rule_min_prefix_length_ipv4,rpz_drop_ip_rule_min_prefix_length_ipv6,rpz_last_updated_time,rpz_policy,rpz_priority,rpz_priority_end,rpz_severity,rpz_type,soa_default_ttl,soa_email,soa_expire,soa_negative_ttl,soa_refresh,soa_retry,soa_serial_number,substitute_name,use_external_primary,use_grid_zone_timer,use_log_rpz,use_record_name_policy,use_rpz_drop_ip_rule,use_soa_email,view"

func TestAccZoneRpResource_basic(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpBasicConfig(zoneFqdn, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", zoneFqdn),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "locked", "false"),
					resource.TestCheckResourceAttr(resourceName, "log_rpz", "true"),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv4", "29"),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv6", "112"),
					resource.TestCheckResourceAttr(resourceName, "rpz_severity", "MAJOR"),
					resource.TestCheckResourceAttr(resourceName, "rpz_type", "LOCAL"),
					resource.TestCheckResourceAttr(resourceName, "use_log_rpz", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_external_primary", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_record_name_policy", "false"),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_disappears(t *testing.T) {
	resourceName := "nios_dns_zone_rp.test"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckZoneRpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneRpBasicConfig(zoneFqdn, "default"),
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
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpComment(zoneFqdn, "default", "Comment for ZONE RP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Comment for ZONE RP"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpComment(zoneFqdn, "default", "Updated Comment for ZONE RP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Updated Comment for ZONE RP"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_disable"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpDisable(zoneFqdn, "default", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpDisable(zoneFqdn, "default", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_extattrs"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpExtAttrs(zoneFqdn, "default", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpExtAttrs(zoneFqdn, "default", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_ExternalPrimaries(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_external_primaries"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"
	tsigKeyName := acctest.RandomName()
	externalPrimaries := []map[string]any{
		{
			"address":           "10.0.0.0",
			"name":              "example-server",
			"tsig_key_alg":      "HMAC-SHA256",
			"tsig_key":          "X4oRe92t54I+T98NdQpV2w==",
			"use_tsig_key_name": true,
			"tsig_key_name":     tsigKeyName,
		},
	}
	updatedExternalPrimaries := []map[string]any{
		{
			"address": "10.0.0.2",
			"name":    "example-updated-server",
		},
	}
	gridSecondary := []map[string]any{
		{
			"name": "member.com",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpExternalPrimaries(zoneFqdn, "default", externalPrimaries, gridSecondary, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.address", "10.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.name", "example-server"),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.use_tsig_key_name", "true"),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.tsig_key_name", tsigKeyName),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.tsig_key", "X4oRe92t54I+T98NdQpV2w=="),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.tsig_key_alg", "HMAC-SHA256"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpExternalPrimaries(zoneFqdn, "default", updatedExternalPrimaries, gridSecondary, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.address", "10.0.0.2"),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.name", "example-updated-server")),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_ExternalSecondaries(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_external_secondaries"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"
	externalSecondaries := []map[string]any{
		{
			"address":           "10.0.0.0",
			"name":              "example.com",
			"tsig_key_alg":      "HMAC-SHA256",
			"tsig_key":          "X4oRe92t54I+T98NdQpV2w==",
			"use_tsig_key_name": false,
			"tsig_key_name":     acctest.RandomName(),
		},
	}
	gridPrimary := []map[string]any{
		{
			"name": "member.com",
		},
	}
	updatedExternalSecondaries := []map[string]any{
		{
			"address": "10.0.0.2",
			"name":    "updated-example.com",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpExternalSecondaries(zoneFqdn, "default", externalSecondaries, gridPrimary),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.address", "10.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.name", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.use_tsig_key_name", "false"),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.tsig_key", "X4oRe92t54I+T98NdQpV2w=="),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.tsig_key_alg", "HMAC-SHA256"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpExternalSecondaries(zoneFqdn, "default", updatedExternalSecondaries, gridPrimary),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.address", "10.0.0.2"),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.name", "updated-example.com")),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_FireeyeRuleMapping(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_fireeye_rule_mapping"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	fireEyeRuleMapping := map[string]any{
		"apt_override": "PASSTHRU",
		"fireeye_alert_mapping": []map[string]any{
			{
				"alert_type": "DOMAIN_MATCH",
				"lifetime":   "86400",
				"rpz_rule":   "NODATA",
			},
		},
	}
	fireEyeRuleMappingUpdate1 := map[string]any{
		"apt_override": "NOOVERRIDE",
		"fireeye_alert_mapping": []map[string]any{
			{
				"alert_type": "INFECTION_MATCH",
				"lifetime":   "0",
				"rpz_rule":   "NONE",
			},
		},
	}

	fireEyeRuleMappingUpdate2 := map[string]any{
		"apt_override": "NXDOMAIN",
		"fireeye_alert_mapping": []map[string]any{
			{
				"alert_type": "MALWARE_CALLBACK",
				"lifetime":   "172800",
				"rpz_rule":   "NXDOMAIN",
			},
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpFireeyeRuleMapping(zoneFqdn, "default", fireEyeRuleMapping),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fireeye_rule_mapping.apt_override", "NODATA"),
					resource.TestCheckResourceAttr(resourceName, "fireeye_rule_mapping.fireeye_alert_mapping.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "fireeye_rule_mapping.fireeye_alert_mapping.0.alert_type", "DOMAIN_MATCH"),
					resource.TestCheckResourceAttr(resourceName, "fireeye_rule_mapping.fireeye_alert_mapping.0.lifetime", "5"),
					resource.TestCheckResourceAttr(resourceName, "fireeye_rule_mapping.fireeye_alert_mapping.0.rpz_rule", "NODATA"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpFireeyeRuleMapping(zoneFqdn, "default", fireEyeRuleMappingUpdate1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fireeye_rule_mapping.apt_override", "NOOVERRIDE"),
					resource.TestCheckResourceAttr(resourceName, "fireeye_rule_mapping.fireeye_alert_mapping.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "fireeye_rule_mapping.fireeye_alert_mapping.0.alert_type", "INFECTION_MATCH"),
					resource.TestCheckResourceAttr(resourceName, "fireeye_rule_mapping.fireeye_alert_mapping.0.lifetime", "0"),
					resource.TestCheckResourceAttr(resourceName, "fireeye_rule_mapping.fireeye_alert_mapping.0.rpz_rule", "NONE"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpFireeyeRuleMapping(zoneFqdn, "default", fireEyeRuleMappingUpdate2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fireeye_rule_mapping.apt_override", "NXDOMAIN"),
					resource.TestCheckResourceAttr(resourceName, "fireeye_rule_mapping.fireeye_alert_mapping.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "fireeye_rule_mapping.fireeye_alert_mapping.0.alert_type", "MALWARE_CALLBACK"),
					resource.TestCheckResourceAttr(resourceName, "fireeye_rule_mapping.fireeye_alert_mapping.0.lifetime", "500"),
					resource.TestCheckResourceAttr(resourceName, "fireeye_rule_mapping.fireeye_alert_mapping.0.rpz_rule", "NXDOMAIN"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_GridPrimary(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_grid_primary"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"
	gridPrimary := []map[string]any{
		{
			"name":    "member.com",
			"stealth": false,
		},
	}
	gridPrimaryUpdated := []map[string]any{
		{
			"name":    "infoblox.localdomain",
			"stealth": true,
		},
		{
			"name":    "member.com",
			"stealth": false,
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpGridPrimary(zoneFqdn, "default", gridPrimary),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_primary.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "grid_primary.0.name", "member.com")),
			},
			// Update and Read
			{
				Config: testAccZoneRpGridPrimary(zoneFqdn, "default", gridPrimaryUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_primary.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "grid_primary.0.name", "infoblox.localdomain"),
					resource.TestCheckResourceAttr(resourceName, "grid_primary.0.stealth", "true"),
					resource.TestCheckResourceAttr(resourceName, "grid_primary.1.name", "member.com"),
					resource.TestCheckResourceAttr(resourceName, "grid_primary.1.stealth", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_GridSecondaries(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_grid_secondaries"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},
	}
	gridSecondary := []map[string]any{
		{
			"name":                       "member.com",
			"stealth":                    false,
			"grid_replicate":             true,
			"lead":                       false,
			"enable_preferred_primaries": false,
		},
	}
	updatedgridPrimary := []map[string]any{
		{
			"name": "member.com",
		},
	}
	updatedGridSecondary := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpGridSecondaries(zoneFqdn, "default", gridPrimary, gridSecondary),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries.0.name", "member.com"),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries.0.stealth", "false"),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries.0.grid_replicate", "true"),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries.0.lead", "false"),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries.0.enable_preferred_primaries", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpGridSecondaries(zoneFqdn, "default", updatedgridPrimary, updatedGridSecondary),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries.0.name", "infoblox.localdomain")),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_Locked(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_locked"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpLocked(zoneFqdn, "default", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "locked", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpLocked(zoneFqdn, "default", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "locked", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// Logging RPZ requires enabling rpz logging on the GRID or member level
func TestAccZoneRpResource_LogRpz(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_log_rpz"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpLogRpz(zoneFqdn, "default", false, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_rpz", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpLogRpz(zoneFqdn, "default", true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_rpz", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_MemberSoaMnames(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_member_soa_mnames"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},
	}
	memberSoaMnames := []map[string]any{
		{
			"grid_primary": "infoblox.localdomain",
			"mname":        "example.com",
		},
	}
	updatedMemberSoaMnames := []map[string]any{
		{
			"mname": "example.com",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpMemberSoaMnames(zoneFqdn, "default", gridPrimary, memberSoaMnames),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member_soa_mnames.0.grid_primary", "infoblox.localdomain"),
					resource.TestCheckResourceAttr(resourceName, "member_soa_mnames.0.mname", "example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpMemberSoaMnames(zoneFqdn, "default", gridPrimary, updatedMemberSoaMnames),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member_soa_mnames.0.mname", "example.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_NsGroup(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_ns_group"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"
	gridPrimary := []map[string]any{
		{
			"name": "member.com",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpNsGroup(zoneFqdn, "default", "nsgroup1", gridPrimary),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ns_group", "nsgroup1"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpNsGroup(zoneFqdn, "default", "nsgroup2", gridPrimary),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ns_group", "nsgroup2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_Prefix(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_prefix"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpPrefix(zoneFqdn, "default", "STUB-b"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", "STUB-b"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpPrefix(zoneFqdn, "default", "121/26"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", "121/26"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_RecordNamePolicy(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_record_name_policy"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpRecordNamePolicy(zoneFqdn, "default", "example-policy", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "record_name_policy", "example-policy"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpRecordNamePolicy(zoneFqdn, "default", "example-policy-update", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "record_name_policy", "example-policy-update"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_RpzDropIpRuleEnabled(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_rpz_drop_ip_rule_enabled"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpRpzDropIpRuleEnabled(zoneFqdn, "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpRpzDropIpRuleEnabled(zoneFqdn, "default", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_RpzDropIpRuleMinPrefixLengthIpv4(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_rpz_drop_ip_rule_min_prefix_length_ipv4"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpRpzDropIpRuleMinPrefixLengthIpv4(zoneFqdn, "default", "20"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv4", "20"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpRpzDropIpRuleMinPrefixLengthIpv4(zoneFqdn, "default", "30"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv4", "30"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_RpzDropIpRuleMinPrefixLengthIpv6(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_rpz_drop_ip_rule_min_prefix_length_ipv6"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpRpzDropIpRuleMinPrefixLengthIpv6(zoneFqdn, "default", "40"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv6", "40"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpRpzDropIpRuleMinPrefixLengthIpv6(zoneFqdn, "default", "50"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv6", "50"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_RpzPolicy(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_rpz_policy"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpRpzPolicy(zoneFqdn, "default", "DISABLED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_policy", "DISABLED"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpRpzPolicy(zoneFqdn, "default", "NODATA"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_policy", "NODATA"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpRpzPolicy(zoneFqdn, "default", "PASSTHRU"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_policy", "PASSTHRU"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpRpzPolicy(zoneFqdn, "default", "SUBSTITUTE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_policy", "SUBSTITUTE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_RpzSeverity(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_rpz_severity"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpRpzSeverity(zoneFqdn, "default", "CRITICAL"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_severity", "CRITICAL"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpRpzSeverity(zoneFqdn, "default", "INFORMATIONAL"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_severity", "INFORMATIONAL"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpRpzSeverity(zoneFqdn, "default", "WARNING"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_severity", "WARNING"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_RpzType(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_rpz_type"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"
	fireEyeRuleMapping := map[string]any{
		"apt_override": "NODATA",
		"fireeye_alert_mapping": []map[string]any{
			{
				"alert_type": "DOMAIN_MATCH",
				"lifetime":   "0",
				"rpz_rule":   "NODATA",
			},
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpRpzType(zoneFqdn, "default", "FIREEYE", fireEyeRuleMapping),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_type", "FIREEYE"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpRpzType(zoneFqdn, "default", "FEED", fireEyeRuleMapping),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_type", "FEED"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_SetSoaSerialNumber(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_set_soa_serial_number"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpSetSoaSerialNumber(zoneFqdn, "default", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "set_soa_serial_number", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpSetSoaSerialNumber(zoneFqdn, "default", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "set_soa_serial_number", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_SoaDefaultTtl(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_soa_default_ttl"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpSoaDefaultTtl(zoneFqdn, "default", gridPrimary, 8, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_default_ttl", "8"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpSoaDefaultTtl(zoneFqdn, "default", gridPrimary, 10, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_default_ttl", "10"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_SoaEmail(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_soa_email"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"
	gridPrimary := []map[string]any{
		{
			"name": "member.com",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpSoaEmail(zoneFqdn, "default", gridPrimary, "user1@example.com", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_email", "user1@example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpSoaEmail(zoneFqdn, "default", gridPrimary, "user2@example.com", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_email", "user2@example.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_SoaExpire(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_soa_expire"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpSoaExpire(zoneFqdn, "default", gridPrimary, 24192, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_expire", "24192"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpSoaExpire(zoneFqdn, "default", gridPrimary, 24100, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_expire", "24100"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_SoaNegativeTtl(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_soa_negative_ttl"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpSoaNegativeTtl(zoneFqdn, "default", gridPrimary, 800, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_negative_ttl", "800"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpSoaNegativeTtl(zoneFqdn, "default", gridPrimary, 900, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_negative_ttl", "900"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_SoaRefresh(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_soa_refresh"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpSoaRefresh(zoneFqdn, "default", gridPrimary, 800, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_refresh", "800"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpSoaRefresh(zoneFqdn, "default", gridPrimary, 900, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_refresh", "900"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_SoaRetry(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_soa_retry"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpSoaRetry(zoneFqdn, "default", gridPrimary, 1600, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_retry", "1600"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpSoaRetry(zoneFqdn, "default", gridPrimary, 1700, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_retry", "1700"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_SoaSerialNumber(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_soa_serial_number"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpSoaSerialNumber(zoneFqdn, "default", gridPrimary, 10, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_serial_number", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpSoaSerialNumber(zoneFqdn, "default", gridPrimary, 20, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_serial_number", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_SubstituteName(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_substitute_name"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpSubstituteName(zoneFqdn, "default", "alternate.fqdn", "SUBSTITUTE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "substitute_name", "alternate.fqdn"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpSubstituteName(zoneFqdn, "default", "updated-Alternate.fqdn", "SUBSTITUTE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "substitute_name", "updated-Alternate.fqdn"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_UseExternalPrimary(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_use_external_primary"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"
	externalPrimaries := []map[string]any{
		{
			"address": "10.0.0.0",
			"name":    "example-server",
		},
	}
	gridSecondaries := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpUseExternalPrimary(zoneFqdn, "default", externalPrimaries, gridSecondaries, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_external_primary", "true"),
				),
			},
			// Update and Read
			{
				Config: estAccZoneRpUseExternalPrimaryUpdate(zoneFqdn, "default", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_external_primary", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_UseGridZoneTimer(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_use_grid_zone_timer"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpUseGridZoneTimer(zoneFqdn, "default", gridPrimary, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_grid_zone_timer", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpUseGridZoneTimerUpdate(zoneFqdn, "default", gridPrimary, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_grid_zone_timer", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_UseLogRpz(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_use_log_rpz"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpUseLogRpz(zoneFqdn, "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_log_rpz", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpUseLogRpz(zoneFqdn, "default", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_log_rpz", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_UseRecordNamePolicy(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_use_record_name_policy"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpUseRecordNamePolicy(zoneFqdn, "default", "example-policy", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_record_name_policy", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpUseRecordNamePolicy(zoneFqdn, "default", "example-policy", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_record_name_policy", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_UseRpzDropIpRule(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_use_rpz_drop_ip_rule"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpUseRpzDropIpRule(zoneFqdn, "default", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_rpz_drop_ip_rule", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpUseRpzDropIpRule(zoneFqdn, "default", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_rpz_drop_ip_rule", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_UseSoaEmail(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_use_soa_email"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpUseSoaEmail(zoneFqdn, "default", gridPrimary, "user@example.com", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_soa_email", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneRpUseSoaEmail(zoneFqdn, "default", gridPrimary, "user@example.com", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_soa_email", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneRpResource_View(t *testing.T) {
	var resourceName = "nios_dns_zone_rp.test_view"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneRpView(zoneFqdn, "test_view"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "test_view"),
				),
			},
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

func testAccZoneRpBasicConfig(zoneFqdn, view string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test" {
    fqdn = %q
    view = %q
}
`, zoneFqdn, view)
}

func testAccZoneRpComment(zoneFqdn, view, comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_comment" {
    fqdn = %q
    view = %q
    comment = %q
}
`, zoneFqdn, view, comment)
}

func testAccZoneRpDisable(zoneFqdn, view string, disable bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_disable" {
    fqdn = %q
    view = %q
    disable = %t
}
`, zoneFqdn, view, disable)
}

func testAccZoneRpExtAttrs(zoneFqdn, view string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_extattrs" {
    fqdn = %q
    view = %q
    extattrs = %s
}
`, zoneFqdn, view, extattrsStr)
}

func testAccZoneRpExternalPrimaries(zoneFqdn, view string, externalPrimaries, gridSecondary []map[string]any, useExternalPrimary bool) string {
	externalPrimariesHCL := utils.ConvertSliceOfMapsToHCL(externalPrimaries)
	gridSecondaryHCL := utils.ConvertSliceOfMapsToHCL(gridSecondary)
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_external_primaries" {
    fqdn = %q
    view = %q
    external_primaries = %s
    grid_secondaries = %s
    use_external_primary = %t
}
`, zoneFqdn, view, externalPrimariesHCL, gridSecondaryHCL, useExternalPrimary)
}

func testAccZoneRpExternalSecondaries(zoneFqdn, view string, externalSecondaries, gridPrimary []map[string]any) string {
	externalSecondariesHCL := utils.ConvertSliceOfMapsToHCL(externalSecondaries)
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_external_secondaries" {
    fqdn = %q
    view = %q
    grid_primary = %s
    external_secondaries = %s
}
`, zoneFqdn, view, gridPrimaryHCL, externalSecondariesHCL)
}

func testAccZoneRpFireeyeRuleMapping(zoneFqdn, view string, fireeyeRuleMapping map[string]any) string {
	fireeyeRuleMappingHCL := utils.ConvertMapToHCL(fireeyeRuleMapping)
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_fireeye_rule_mapping" {
    fqdn = %q
    view = %q
    fireeye_rule_mapping = %s
	rpz_type = "FIREEYE"
}
`, zoneFqdn, view, fireeyeRuleMappingHCL)
}

func testAccZoneRpGridPrimary(zoneFqdn, view string, gridPrimary []map[string]any) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_grid_primary" {
    fqdn = %q
    view = %q
    grid_primary = %s
}
`, zoneFqdn, view, gridPrimaryHCL)
}

func testAccZoneRpGridSecondaries(zoneFqdn, view string, gridPrimary, gridSecondary []map[string]any) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	gridSecondaryHCL := utils.ConvertSliceOfMapsToHCL(gridSecondary)
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_grid_secondaries" {
    fqdn = %q
    view = %q
	grid_primary = %s
    grid_secondaries = %s
}
`, zoneFqdn, view, gridPrimaryHCL, gridSecondaryHCL)
}

func testAccZoneRpLocked(zoneFqdn, view string, locked bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_locked" {
    fqdn = %q
    view = %q
    locked = %t
}
`, zoneFqdn, view, locked)
}

func testAccZoneRpLogRpz(zoneFqdn, view string, logRpz, useLogRpz bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_log_rpz" {
    fqdn = %q
    view = %q
    log_rpz = %t
	use_log_rpz = %t
}
`, zoneFqdn, view, logRpz, useLogRpz)
}

func testAccZoneRpMemberSoaMnames(zoneFqdn, view string, gridPrimary, memberSoaMnames []map[string]any) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	memberSoaMnamesHCL := utils.ConvertSliceOfMapsToHCL(memberSoaMnames)
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_member_soa_mnames" {
    fqdn = %q
    view = %q
    grid_primary = %s
    member_soa_mnames = %s
}
`, zoneFqdn, view, gridPrimaryHCL, memberSoaMnamesHCL)
}

func testAccZoneRpNsGroup(zoneFqdn, view string, nsGroupName string, gridPrimary []map[string]any) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test" {
  name = %q
  grid_primary = %s
}

resource "nios_dns_zone_rp" "test_ns_group" {
    fqdn = %q
    view = %q
    ns_group = nios_dns_nsgroup.test.name
}
`, nsGroupName, gridPrimaryHCL, zoneFqdn, view)
}

func testAccZoneRpPrefix(zoneFqdn, view, prefix string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_prefix" {
    fqdn = %q
    view = %q
    prefix = %q
}
`, zoneFqdn, view, prefix)
}

func testAccZoneRpRecordNamePolicy(zoneFqdn, view, recordNamePolicy string, useRecordNamePolicy bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_record_name_policy" {
    fqdn = %q
    view = %q
    record_name_policy = %q
	use_record_name_policy = %t
}
`, zoneFqdn, view, recordNamePolicy, useRecordNamePolicy)
}

func testAccZoneRpRpzDropIpRuleEnabled(zoneFqdn, view, rpzDropIpRuleEnabled string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_rpz_drop_ip_rule_enabled" {
    fqdn = %q
    view = %q
    rpz_drop_ip_rule_enabled = %q
    use_rpz_drop_ip_rule = true
}
`, zoneFqdn, view, rpzDropIpRuleEnabled)
}

func testAccZoneRpRpzDropIpRuleMinPrefixLengthIpv4(zoneFqdn, view, rpzDropIpRuleMinPrefixLengthIpv4 string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_rpz_drop_ip_rule_min_prefix_length_ipv4" {
    fqdn = %q
    view = %q
    rpz_drop_ip_rule_min_prefix_length_ipv4 = %q
    use_rpz_drop_ip_rule = true
}
`, zoneFqdn, view, rpzDropIpRuleMinPrefixLengthIpv4)
}

func testAccZoneRpRpzDropIpRuleMinPrefixLengthIpv6(zoneFqdn, view, rpzDropIpRuleMinPrefixLengthIpv6 string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_rpz_drop_ip_rule_min_prefix_length_ipv6" {
    fqdn = %q
    view = %q
    rpz_drop_ip_rule_min_prefix_length_ipv6 = %q
    use_rpz_drop_ip_rule = true
}
`, zoneFqdn, view, rpzDropIpRuleMinPrefixLengthIpv6)
}

func testAccZoneRpRpzPolicy(zoneFqdn, view, rpzPolicy string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_rpz_policy" {
    fqdn = %q
    view = %q
    rpz_policy = %q
    substitute_name = "substitute.fqdn"
}
`, zoneFqdn, view, rpzPolicy)
}

func testAccZoneRpRpzSeverity(zoneFqdn, view, rpzSeverity string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_rpz_severity" {
    fqdn = %q
    view = %q
    rpz_severity = %q
}
`, zoneFqdn, view, rpzSeverity)
}

func testAccZoneRpRpzType(zoneFqdn, view, rpzType string, fireEyeRuleMapping map[string]any) string {
	fireeyeRuleMappingHCL := utils.ConvertMapToHCL(fireEyeRuleMapping)
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_rpz_type" {
    fqdn = %q
    view = %q
    rpz_type = %q
    fireeye_rule_mapping = %s
}
`, zoneFqdn, view, rpzType, fireeyeRuleMappingHCL)
}

func testAccZoneRpSetSoaSerialNumber(zoneFqdn, view string, setSoaSerialNumber bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_set_soa_serial_number" {
    fqdn = %q
    view = %q
    set_soa_serial_number = %t
}
`, zoneFqdn, view, setSoaSerialNumber)
}

func testAccZoneRpSoaDefaultTtl(zoneFqdn, view string, gridPrimary []map[string]any, soaDefaultTtl int64, useGridZoneTimer bool) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_soa_default_ttl" {
    fqdn = %q
    view = %q
    grid_primary = %s
    soa_default_ttl = %d
	soa_expire = 2419200
    soa_negative_ttl = 900
    soa_refresh = 10800
    soa_retry = 3600
    use_grid_zone_timer = %t
}
`, zoneFqdn, view, gridPrimaryHCL, soaDefaultTtl, useGridZoneTimer)
}

func testAccZoneRpSoaEmail(zoneFqdn, view string, gridPrimary []map[string]any, soaEmail string, useSoaEmail bool) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_soa_email" {
    fqdn = %q
    view = %q
    grid_primary = %s
    soa_email = %q
    use_soa_email = %t
}
`, zoneFqdn, view, gridPrimaryHCL, soaEmail, useSoaEmail)
}

func testAccZoneRpSoaExpire(zoneFqdn, view string, gridPrimary []map[string]any, soaExpire int64, useGridZoneTimer bool) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_soa_expire" {
    fqdn = %q
    view = %q
    grid_primary = %s
    soa_expire = %d
	soa_default_ttl = 28800
	soa_negative_ttl = 900
    soa_refresh = 10800
    soa_retry = 3600
    use_grid_zone_timer = %t
}
`, zoneFqdn, view, gridPrimaryHCL, soaExpire, useGridZoneTimer)
}

func testAccZoneRpSoaNegativeTtl(zoneFqdn, view string, gridPrimary []map[string]any, soaNegativeTtl int64, useGridZoneTimer bool) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_soa_negative_ttl" {
    fqdn = %q
    view = %q
    grid_primary = %s
    soa_negative_ttl = %d
	soa_expire = 2419200
	soa_default_ttl = 28800
    soa_refresh = 10800
    soa_retry = 3600

    use_grid_zone_timer = %t
}
`, zoneFqdn, view, gridPrimaryHCL, soaNegativeTtl, useGridZoneTimer)
}

func testAccZoneRpSoaRefresh(zoneFqdn, view string, gridPrimary []map[string]any, soaRefresh int64, useGridZoneTimer bool) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_soa_refresh" {
    fqdn = %q
    view = %q
    grid_primary = %s
    soa_refresh = %d
    soa_negative_ttl = 900
	soa_expire = 2419200
	soa_default_ttl = 28800
    soa_retry = 3600  
	use_grid_zone_timer = %t
}
`, zoneFqdn, view, gridPrimaryHCL, soaRefresh, useGridZoneTimer)
}

func testAccZoneRpSoaRetry(zoneFqdn, view string, gridPrimary []map[string]any, soaRetry int64, useGridZoneTimer bool) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_soa_retry" {
    fqdn = %q
    view = %q
    grid_primary = %s
    soa_retry = %d
    soa_negative_ttl = 900
	soa_expire = 2419200
	soa_default_ttl = 28800
    soa_refresh = 10800
    use_grid_zone_timer = %t
}
`, zoneFqdn, view, gridPrimaryHCL, soaRetry, useGridZoneTimer)
}

func testAccZoneRpSoaSerialNumber(zoneFqdn, view string, gridPrimary []map[string]any, soaSerialNumber int64, SetSoaSerialNumber bool) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_soa_serial_number" {
    fqdn = %q
    view = %q
    grid_primary = %s
    soa_serial_number = %d
	set_soa_serial_number = %t
	soa_retry = 3600
    soa_negative_ttl = 900
	soa_expire = 2419200
	soa_default_ttl = 28800
    soa_refresh = 10800
	use_grid_zone_timer = true

}`, zoneFqdn, view, gridPrimaryHCL, soaSerialNumber, SetSoaSerialNumber)
}

func testAccZoneRpSubstituteName(zoneFqdn, view, substituteName, rpzPolicy string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_substitute_name" {
    fqdn = %q
    view = %q
    substitute_name = %q
	rpz_policy = %q
}
`, zoneFqdn, view, substituteName, rpzPolicy)
}

func testAccZoneRpUseExternalPrimary(zoneFqdn, view string, externalPrimaries, gridSecondaries []map[string]any, useExternalPrimary bool) string {
	externalPrimariesHCL := utils.ConvertSliceOfMapsToHCL(externalPrimaries)
	msSecondariesHCL := utils.ConvertSliceOfMapsToHCL(gridSecondaries)
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_use_external_primary" {
    fqdn = %q
    view = %q
    external_primaries = %s
    grid_secondaries = %s
    use_external_primary = %t
}
`, zoneFqdn, view, externalPrimariesHCL, msSecondariesHCL, useExternalPrimary)
}

func estAccZoneRpUseExternalPrimaryUpdate(zoneFqdn, view string, useExternalPrimary bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_use_external_primary" {
    fqdn = %q
    view = %q
    use_external_primary = %t
}
`, zoneFqdn, view, useExternalPrimary)
}

func testAccZoneRpUseGridZoneTimer(zoneFqdn, view string, gridPrimary []map[string]any, useGridZoneTimer bool) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_use_grid_zone_timer" {
    fqdn = %q
    view = %q
    grid_primary = %s
	soa_default_ttl = 37000
	soa_expire = 2419200
    soa_negative_ttl = 900
    soa_refresh = 10800
    soa_retry = 3600
    use_grid_zone_timer = %t
}
`, zoneFqdn, view, gridPrimaryHCL, useGridZoneTimer)
}

func testAccZoneRpUseGridZoneTimerUpdate(zoneFqdn, view string, gridPrimary []map[string]any, useGridZoneTimer bool) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_use_grid_zone_timer" {
    fqdn = %q
    view = %q
    grid_primary = %s
    use_grid_zone_timer = %t
}
`, zoneFqdn, view, gridPrimaryHCL, useGridZoneTimer)
}

func testAccZoneRpUseLogRpz(zoneFqdn, view, useLogRpz string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_use_log_rpz" {
    fqdn = %q
    view = %q
    use_log_rpz = %q
    log_rpz = true
}
`, zoneFqdn, view, useLogRpz)
}

func testAccZoneRpUseRecordNamePolicy(zoneFqdn, view, recordNamePolicy string, useRecordNamePolicy bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_use_record_name_policy" {
    fqdn = %q
    view = %q
    record_name_policy = %q
    use_record_name_policy = %t
}
`, zoneFqdn, view, recordNamePolicy, useRecordNamePolicy)
}

func testAccZoneRpUseRpzDropIpRule(zoneFqdn, view, useRpzDropIpRule string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_use_rpz_drop_ip_rule" {
    fqdn = %q
    view = %q
    use_rpz_drop_ip_rule = %q
    rpz_drop_ip_rule_min_prefix_length_ipv4 = 20
}
`, zoneFqdn, view, useRpzDropIpRule)
}

func testAccZoneRpUseSoaEmail(zoneFqdn, view string, gridPrimary []map[string]any, soaEmail string, useSoaEmail bool) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test_use_soa_email" {
    fqdn = %q
    view = %q
    grid_primary = %s
    soa_email = %q
    use_soa_email = %t
}
`, zoneFqdn, view, gridPrimaryHCL, soaEmail, useSoaEmail)
}

func testAccZoneRpView(zoneFqdn, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_view" {
  name = "test_view"
}

resource "nios_dns_zone_rp" "test_view" {
	fqdn = %q
	view = %q
	depends_on = [nios_dns_view.test_view]
}
`, zoneFqdn, view)
}
